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
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("t"), _dafny.IntOfInt64(-80))
}
func (_static *CompanionStruct_Default___) Fm5(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qt")).Cardinality())))).Cardinality(), (_dafny.SetOf(true)).Cardinality())).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfInt64(-958))), !(!(false)) || (false), !(false))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('j')
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D1 = Companion_D1_.Create_DC4_(_dafny.SetOf(_dafny.IntOfInt64(934)))
	_ = _source0
	if _source0.Is_DC5() {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(207))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(626))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg1 _dafny.Int) interface{} {
				return coer1(arg1)
			}
		}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})))
	} else if _source0.Is_DC6() {
		var _2___mcc_h0 _dafny.Sequence = _source0.Get_().(D1_DC6).Cf5
		_ = _2___mcc_h0
		var _3___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC6).Cf6
		_ = _3___mcc_h1
		var _4_cf6 _dafny.Int = _3___mcc_h1
		_ = _4_cf6
		var _5_cf5 _dafny.Sequence = _2___mcc_h0
		_ = _5_cf5
		return _5_cf5
	} else {
		var _6___mcc_h2 _dafny.Set = _source0.Get_().(D1_DC4).Cf4
		_ = _6___mcc_h2
		var _7_cf4 _dafny.Set = _6___mcc_h2
		_ = _7_cf4
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(997))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}(func(_8_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true) || (false), ((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-626))).Cardinality())).Cardinality()).Cmp(_dafny.IntOfInt64(984)) >= 0)
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("je"), _dafny.UnicodeSeqOfUtf8Bytes("rfsuhwou")), (func() _dafny.Sequence {
		if true {
			return _dafny.UnicodeSeqOfUtf8Bytes("a")
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(168))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(780), _dafny.IntOfInt64(411), _dafny.IntOfInt64(574), (_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(362)), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality()), _dafny.IntOfInt64(986), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality())), false)).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(187)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(161))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_10_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(false, true)).Cardinality()
	}))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(338))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_11_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))).Cardinality()), _dafny.IntOfInt64(623), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-883)))).Cardinality()))).Cardinality())).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(940)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(221))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_12_i2 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), false)).Cardinality()
	}))))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(710))).Cardinality()), _dafny.IntOfInt64(-724)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-115), _dafny.IntOfInt64(120)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Intersection((_dafny.MultiSetOf(!(true), true)).Intersection(_dafny.MultiSetOf(false, false, false)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 D0, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC3_(false)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-390))).Uint32(), func(coer7 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_13_i0 _dafny.Int) D0 {
			return Companion_D0_.Create_DC3_(false)
		}))))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _14_v0 D0
			_14_v0 = interface{}(_compr_0).(D0)
			if (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_D0_.Create_DC3_(false)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-390))).Uint32(), func(coer8 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_13_i0 _dafny.Int) D0 {
				return Companion_D0_.Create_DC3_(false)
			}))))).Contains(_14_v0) {
				_coll0.Add(_14_v0, !((_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Cmp(_dafny.IntOfInt64(713)) == 0))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('y')
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jx")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-653), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pbdbqo")).Cardinality()), _dafny.IntOfInt64(-821), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("u")).Cardinality()))))).Cardinality())).Intersection(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(387))).Cardinality()))).Cardinality()))).Union(_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D10_.Create_DC30_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wrnolacmv")).Cardinality())), _dafny.IntOfInt64(291))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Set, p1 bool, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_()
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, !(false), false, false)).Intersection((_dafny.SetOf(true)).Intersection(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Set, p1 bool, globalState *GlobalState) D3 {
	var _source1 _dafny.MultiSet = _dafny.MultiSetOf(false)
	_ = _source1
	var _15___mcc_h0 _dafny.MultiSet = _source1
	_ = _15___mcc_h0
	var _16_cf64 _dafny.MultiSet = _15___mcc_h0
	_ = _16_cf64
	return Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false))
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true))), false)).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), false)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)), true)
	})())
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC21_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)), true)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_((true) || (true), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).IsProperSubsetOf(_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(792)))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.CodePoint, p2 D4, globalState *GlobalState) _dafny.Map {
	var _source2 D4 = Companion_D4_.Create_DC14_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(315))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_17_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()))).Cardinality()), !(!(true)), (_dafny.MultiSetOf((_dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(807))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfInt64(865))).Cardinality())).Cardinality(), !(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), true)).Cardinality())
	_ = _source2
	if _source2.Is_DC12() {
		var _18___mcc_h0 bool = _source2.Get_().(D4_DC12).Cf14
		_ = _18___mcc_h0
		var _19___mcc_h1 _dafny.Int = _source2.Get_().(D4_DC12).Cf15
		_ = _19___mcc_h1
		var _20___mcc_h2 bool = _source2.Get_().(D4_DC12).Cf16
		_ = _20___mcc_h2
		var _21_cf16 bool = _20___mcc_h2
		_ = _21_cf16
		var _22_cf15 _dafny.Int = _19___mcc_h1
		_ = _22_cf15
		var _23_cf14 bool = _18___mcc_h0
		_ = _23_cf14
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_21_cf16, _22_cf15), _22_cf15)
	} else if _source2.Is_DC13() {
		var _24___mcc_h3 bool = _source2.Get_().(D4_DC13).Cf17
		_ = _24___mcc_h3
		var _25___mcc_h4 _dafny.Int = _source2.Get_().(D4_DC13).Cf18
		_ = _25___mcc_h4
		var _26___mcc_h5 bool = _source2.Get_().(D4_DC13).Cf19
		_ = _26___mcc_h5
		var _27_cf19 bool = _26___mcc_h5
		_ = _27_cf19
		var _28_cf18 _dafny.Int = _25___mcc_h4
		_ = _28_cf18
		var _29_cf17 bool = _24___mcc_h3
		_ = _29_cf17
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_29_cf17), (_dafny.SetOf(_27_cf19)).Cardinality()), _28_cf18)
	} else if _source2.Is_DC14() {
		var _30___mcc_h6 _dafny.Int = _source2.Get_().(D4_DC14).Cf20
		_ = _30___mcc_h6
		var _31___mcc_h7 bool = _source2.Get_().(D4_DC14).Cf21
		_ = _31___mcc_h7
		var _32___mcc_h8 _dafny.Int = _source2.Get_().(D4_DC14).Cf22
		_ = _32___mcc_h8
		var _33___mcc_h9 bool = _source2.Get_().(D4_DC14).Cf23
		_ = _33___mcc_h9
		var _34___mcc_h10 _dafny.Int = _source2.Get_().(D4_DC14).Cf24
		_ = _34___mcc_h10
		var _35_cf24 _dafny.Int = _34___mcc_h10
		_ = _35_cf24
		var _36_cf23 bool = _33___mcc_h9
		_ = _36_cf23
		var _37_cf22 _dafny.Int = _32___mcc_h8
		_ = _37_cf22
		var _38_cf21 bool = _31___mcc_h7
		_ = _38_cf21
		var _39_cf20 _dafny.Int = _30___mcc_h6
		_ = _39_cf20
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_36_cf23, (_dafny.SetOf(_35_cf24)).Cardinality()), _39_cf20)).Merge(func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_37_cf22)), !(_38_cf21))).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _40_v0 _dafny.Map
				_40_v0 = interface{}(_compr_1).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_37_cf22)), !(_38_cf21))).Contains(_40_v0) {
					_coll1.Add(_40_v0, _39_cf20)
				}
			}
			return _coll1.ToMap()
		}())
	} else if _source2.Is_DC11() {
		var _41___mcc_h11 _dafny.Sequence = _source2.Get_().(D4_DC11).Cf13
		_ = _41___mcc_h11
		var _42_cf13 _dafny.Sequence = _41___mcc_h11
		_ = _42_cf13
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(438)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(90), true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(345)), _dafny.IntOfInt64(705)))
	} else {
		var _43___mcc_h12 D4 = _source2.Get_().(D4_DC15).Cf25
		_ = _43___mcc_h12
		var _44_cf25 D4 = _43___mcc_h12
		_ = _44_cf25
		if true {
			return func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(322)))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _45_v1 _dafny.Map
					_45_v1 = interface{}(_compr_2).(_dafny.Map)
					if (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(322)))).Contains(_45_v1) {
						_coll2.Add(_45_v1, _dafny.IntOfInt64(288))
					}
				}
				return _coll2.ToMap()
			}()
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cjrwfbgy")).Cardinality()), _dafny.IntOfInt64(-569))).Cardinality())), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(486), _dafny.IntOfInt64(-839))).Cardinality()))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC13_(!(true) || (true), _dafny.IntOfInt64(915), (func() bool {
		if true {
			return false
		}
		return false
	})())
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(((_dafny.MultiSetOf(_dafny.IntOfInt64(475))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(619), _dafny.IntOfInt64(473), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())))).Cardinality(), _dafny.IntOfInt64(547))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(611))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_46_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), false)
	})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), !(true))))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("qqwr"))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(832), true)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(828), !(false))).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-711), _dafny.IntOfInt64(941))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _47_v0 _dafny.Int
			_47_v0 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(-711)).Cmp(_47_v0) <= 0) && ((_47_v0).Cmp(_dafny.IntOfInt64(941)) < 0) {
				_coll3.Add(Companion_Default___.SafeModuloInt(_47_v0, _dafny.IntOfInt64(433)), true)
			}
		}
		return _coll3.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))))
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((true) || (false), true)
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate((func() _dafny.Map {
				var _coll6 = _dafny.NewMapBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Keys().Elements()); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _48_v1 _dafny.Int
					_48_v1 = interface{}(_compr_6).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Contains(_48_v1) {
						_coll6.Add(Companion_Default___.SafeDivisionInt(_48_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajo")).Cardinality())), _dafny.IntOfInt64(229))
					}
				}
				return _coll6.ToMap()
			}()).Keys().Elements()); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _49_v0 _dafny.Int
				_49_v0 = interface{}(_compr_5).(_dafny.Int)
				if (func() _dafny.Map {
					var _coll7 = _dafny.NewMapBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Keys().Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _48_v1 _dafny.Int
						_48_v1 = interface{}(_compr_7).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Contains(_48_v1) {
							_coll7.Add(Companion_Default___.SafeDivisionInt(_48_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajo")).Cardinality())), _dafny.IntOfInt64(229))
						}
					}
					return _coll7.ToMap()
				}()).Contains(_49_v0) {
					_coll5.Add((_49_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("jvslhv"))
				}
			}
			return _coll5.ToMap()
		}()).Keys().Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _50_v2 _dafny.Int
			_50_v2 = interface{}(_compr_4).(_dafny.Int)
			if (func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate((func() _dafny.Map {
					var _coll9 = _dafny.NewMapBuilder()
					_ = _coll9
					for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Keys().Elements()); ; {
						_compr_9, _ok9 := _iter9()
						if !_ok9 {
							break
						}
						var _48_v1 _dafny.Int
						_48_v1 = interface{}(_compr_9).(_dafny.Int)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Contains(_48_v1) {
							_coll9.Add(Companion_Default___.SafeDivisionInt(_48_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajo")).Cardinality())), _dafny.IntOfInt64(229))
						}
					}
					return _coll9.ToMap()
				}()).Keys().Elements()); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _49_v0 _dafny.Int
					_49_v0 = interface{}(_compr_8).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll10 = _dafny.NewMapBuilder()
						_ = _coll10
						for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Keys().Elements()); ; {
							_compr_10, _ok10 := _iter10()
							if !_ok10 {
								break
							}
							var _48_v1 _dafny.Int
							_48_v1 = interface{}(_compr_10).(_dafny.Int)
							if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), false)).Contains(_48_v1) {
								_coll10.Add(Companion_Default___.SafeDivisionInt(_48_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajo")).Cardinality())), _dafny.IntOfInt64(229))
							}
						}
						return _coll10.ToMap()
					}()).Contains(_49_v0) {
						_coll8.Add((_49_v0).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("jvslhv"))
					}
				}
				return _coll8.ToMap()
			}()).Contains(_50_v2) {
				_coll4.Add((_50_v2).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-61))))
			}
		}
		return _coll4.ToSet()
	}()).Union(func() _dafny.Set {
		var _coll11 = _dafny.NewBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(567), _dafny.IntOfInt64(438))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _51_v3 _dafny.Int
			_51_v3 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(567)).Cmp(_51_v3) <= 0) && ((_51_v3).Cmp(_dafny.IntOfInt64(438)) < 0) {
				_coll11.Add((_51_v3).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-116))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg11 _dafny.Int) interface{} {
						return coer11(arg11)
					}
				}(func(_52_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				}))).Cardinality())))
			}
		}
		return _coll11.ToSet()
	}())).Union(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xjjpql")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(404), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(508), _dafny.CodePoint('y'))).Cardinality())).Cardinality(), _dafny.IntOfInt64(944)))
}
func (_static *CompanionStruct_Default___) Fm35(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll12 = _dafny.NewBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(570))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_53_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.UnicodeSeqOfUtf8Bytes("a")
		}))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _54_v0 _dafny.Sequence
			_54_v0 = interface{}(_compr_12).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(570))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_53_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("a")
			})), _54_v0) {
				_coll12.Add(_54_v0)
			}
		}
		return _coll12.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("kejowhwr")))).Union((func() _dafny.Set {
		var _coll13 = _dafny.NewBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("iofop"), _dafny.SeqOf(true, false))).Keys().Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _55_v1 _dafny.Sequence
			_55_v1 = interface{}(_compr_13).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("iofop"), _dafny.SeqOf(true, false))).Contains(_55_v1) {
				_coll13.Add(_55_v1)
			}
		}
		return _coll13.ToSet()
	}()).Intersection(func() _dafny.Set {
		var _coll14 = _dafny.NewBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("thbw"))).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _56_v2 _dafny.Sequence
			_56_v2 = interface{}(_compr_14).(_dafny.Sequence)
			if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("thbw"))).Contains(_56_v2) {
				_coll14.Add(_56_v2)
			}
		}
		return _coll14.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm36(p0 D5, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC9_((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(949)))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(-759), _dafny.IntOfInt64(320))), Companion_D0_.Create_DC2_(), _dafny.IntOfInt64(829))
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Int {
	return (Companion_D3_.Create_DC9_(_dafny.MultiSetOf(_dafny.IntOfInt64(198), (func() _dafny.Set {
		var _coll15 = _dafny.NewBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter16 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('h'))).Elements()); ; {
				_compr_16, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _57_v0 _dafny.CodePoint
				_57_v0 = interface{}(_compr_16).(_dafny.CodePoint)
				if (_dafny.MultiSetOf(_dafny.CodePoint('h'))).Contains(_57_v0) {
					_coll16.Add(_57_v0, false)
				}
			}
			return _coll16.ToMap()
		}()).Keys().Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _58_v1 _dafny.CodePoint
			_58_v1 = interface{}(_compr_15).(_dafny.CodePoint)
			if (func() _dafny.Map {
				var _coll17 = _dafny.NewMapBuilder()
				_ = _coll17
				for _iter17 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('h'))).Elements()); ; {
					_compr_17, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _57_v0 _dafny.CodePoint
					_57_v0 = interface{}(_compr_17).(_dafny.CodePoint)
					if (_dafny.MultiSetOf(_dafny.CodePoint('h'))).Contains(_57_v0) {
						_coll17.Add(_57_v0, false)
					}
				}
				return _coll17.ToMap()
			}()).Contains(_58_v1) {
				_coll15.Add(_58_v1)
			}
		}
		return _coll15.ToSet()
	}()).Cardinality()), Companion_D0_.Create_DC2_(), _dafny.IntOfInt64(690))).Dtor_cf11()
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Map, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	return (Companion_D12_.Create_DC35_(true, _dafny.IntOfInt64(723), _dafny.IntOfInt64(135))).Dtor_cf57()
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Map {
	if !((_dafny.IntOfInt64(-773)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(11))).Cardinality()), (_dafny.SetOf(true)).Cardinality())).Cardinality())) == 0) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gnr")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(643)))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(824))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pmlngaghv")).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.CodePoint, globalState *GlobalState) (_dafny.Array, bool, _dafny.MultiSet, _dafny.Int) {
	var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	var _59_v0 bool
	_ = _59_v0
	_59_v0 = true
	if (_59_v0) || ((_59_v0) || (_59_v0)) {
		var _60_v1 _dafny.Int
		_ = _60_v1
		_60_v1 = _dafny.IntOfInt64(125)
		var _61_v2 _dafny.Sequence
		_ = _61_v2
		_61_v2 = _dafny.UnicodeSeqOfUtf8Bytes("or")
		var _62_v3 _dafny.MultiSet
		_ = _62_v3
		_62_v3 = _dafny.MultiSetOf(p0)
		var _63_v4 D9
		_ = _63_v4
		_63_v4 = Companion_D9_.Create_DC27_(_62_v3)
		var _64_v5 _dafny.Sequence
		_ = _64_v5
		_64_v5 = _dafny.SeqOf(_63_v4)
		var _65_v6 _dafny.MultiSet
		_ = _65_v6
		_65_v6 = _dafny.MultiSetOf(_59_v0)
		var _66_v7 _dafny.Map
		_ = _66_v7
		_66_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _65_v6)
		var _67_v8 _dafny.Map
		_ = _67_v8
		_67_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_66_v7).Cardinality(), _60_v1)
		var _nwElement0_0 _dafny.Int = _60_v1
		_ = _nwElement0_0
		var _nw0 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(14))
		_ = _nw0
		(_nw0).ArraySet1(_nwElement0_0, 0)
		(_nw0).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_61_v2, _61_v2)).Cardinality()), 1)
		(_nw0).ArraySet1(_60_v1, 2)
		(_nw0).ArraySet1(_60_v1, 3)
		(_nw0).ArraySet1(_dafny.IntOfInt64(-822), 4)
		(_nw0).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_64_v5).Cardinality())), 5)
		(_nw0).ArraySet1(_60_v1, 6)
		(_nw0).ArraySet1(_60_v1, 7)
		(_nw0).ArraySet1(_60_v1, 8)
		(_nw0).ArraySet1(_60_v1, 9)
		(_nw0).ArraySet1(_dafny.IntOfInt64(-244), 10)
		(_nw0).ArraySet1(_60_v1, 11)
		(_nw0).ArraySet1(_60_v1, 12)
		(_nw0).ArraySet1((_67_v8).Cardinality(), 13)
		r0 = _nw0
		var _68_v9 _dafny.Array
		_ = _68_v9
		var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw1
		_68_v9 = _nw1
		var _69_v10 _dafny.Array
		_ = _69_v10
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_0
		var _nw2 _dafny.Array
		_ = _nw2
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw2 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_70_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_71_i0 _dafny.Int) _dafny.Int {
					return (_71_i0).Times(_70_v1)
				}
			})(_60_v1)
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
		_69_v10 = _nw2
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_68_v9), 0))
		_ = _index0
		(_68_v9).ArraySet1(_69_v10, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_68_v9), 0))
		_ = _index1
		var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw3
		(_68_v9).ArraySet1(_nw3, (_index1).Int())
		r3 = Companion_Default___.SafeDivisionInt(_60_v1, (func() _dafny.Int {
			if _59_v0 {
				return _60_v1
			}
			return _60_v1
		})())
		var _72_v11 _dafny.Sequence
		_ = _72_v11
		_72_v11 = _dafny.SeqOf(_60_v1)
		var _73_v12 _dafny.Map
		_ = _73_v12
		_73_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v0, (_72_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.IntOfUint32((_72_v11).Cardinality()))).Uint32()).(_dafny.Int))
		var _74_v13 _dafny.MultiSet
		_ = _74_v13
		_74_v13 = _dafny.MultiSetOf(_dafny.IntOfInt64(-256), _dafny.IntOfInt64(656))
		var _rhs0 bool = (_74_v13).IsProperSubsetOf(_74_v13)
		_ = _rhs0
		var _rhs1 _dafny.Map = (_73_v12).Merge(_73_v12)
		_ = _rhs1
		var _rhs2 bool = _59_v0
		_ = _rhs2
		var _rhs3 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-981))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}((func(_75_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_76_i1 _dafny.Int) _dafny.CodePoint {
				return _75_p0
			}
		})(p0)))
		_ = _rhs3
		var _rhs4 _dafny.Int = _60_v1
		_ = _rhs4
		r1 = _rhs0
		_73_v12 = _rhs1
		r1 = _rhs2
		_61_v2 = _rhs3
		_60_v1 = _rhs4
		r1 = _59_v0
	} else {
		var _77_v14 _dafny.Int
		_ = _77_v14
		_77_v14 = _dafny.One
		var _78_v15 _dafny.Sequence
		_ = _78_v15
		_78_v15 = _dafny.SeqOf(_59_v0)
		var _rhs5 _dafny.Int = (_dafny.Zero).Minus(_77_v14)
		_ = _rhs5
		var _rhs6 _dafny.Int = _dafny.IntOfInt64(844)
		_ = _rhs6
		var _rhs7 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_78_v15, _78_v15)).Cardinality())).Plus(_77_v14))
		_ = _rhs7
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		_lhs0.F2 = _rhs5
		_lhs1.F2 = _rhs6
		_lhs2.F2 = _rhs7
		_77_v14 = _dafny.IntOfInt64(148)
		var _79_v16 _dafny.Sequence
		_ = _79_v16
		_79_v16 = _dafny.UnicodeSeqOfUtf8Bytes("tuyexyypg")
		var _80_v17 *C2
		_ = _80_v17
		var _nw4 *C2 = New_C2_()
		_ = _nw4
		_nw4.Ctor__(_dafny.IntOfUint32((_79_v16).Cardinality()), _77_v14)
		_80_v17 = _nw4
		(globalState).F2 = _77_v14
		(globalState).F2 = (Companion_Default___.SafeDivisionInt(Companion_Default___.Fm37(_59_v0, _dafny.IntOfInt64(805), !(_59_v0), globalState), _77_v14)).Minus(_dafny.IntOfInt64(491))
	}
	var _81_v18 _dafny.Int
	_ = _81_v18
	_81_v18 = _dafny.IntOfInt64(-649)
	var _hi0 _dafny.Int = _81_v18
	_ = _hi0
	for _82_i2 := _81_v18; _82_i2.Cmp(_hi0) < 0; _82_i2 = _82_i2.Plus(_dafny.One) {
		var _83_v19 T1
		_ = _83_v19
		var _nw5 *C3 = New_C3_()
		_ = _nw5
		_nw5.Ctor__(false, (_82_i2).Times(_81_v18), Companion_Default___.SafeModuloInt(_81_v18, _82_i2))
		_83_v19 = _nw5
		var _nw6 *C1 = New_C1_()
		_ = _nw6
		_nw6.Ctor__(_82_i2, (_82_i2).Times(_dafny.IntOfInt64(-405)))
		_83_v19 = _nw6
		var _84_v20 *C2
		_ = _84_v20
		var _nw7 *C2 = New_C2_()
		_ = _nw7
		_nw7.Ctor__(_81_v18, _83_v19.F6())
		_84_v20 = _nw7
		var _85_v21 _dafny.Sequence
		_ = _85_v21
		_85_v21 = _dafny.UnicodeSeqOfUtf8Bytes("krfj")
		_85_v21 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7(p0, true, _59_v0, _81_v18, globalState), _dafny.Companion_Sequence_.Concatenate(_85_v21, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("bwr"), (Companion_Default___.SafeIndex(_82_i2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bwr")).Cardinality()))).Uint32(), p0))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm7(p0, true, _59_v0, _81_v18, globalState), _dafny.Companion_Sequence_.Concatenate(_85_v21, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("bwr"), (Companion_Default___.SafeIndex(_82_i2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bwr")).Cardinality()))).Uint32(), p0)))).Cardinality()))).Uint32(), (_85_v21).Select((Companion_Default___.SafeIndex(_82_i2, _dafny.IntOfUint32((_85_v21).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _86_v22 _dafny.Array
		_ = _86_v22
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_1
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_87_v0 bool) func(_dafny.Int) bool {
				return func(_88_i3 _dafny.Int) bool {
					return _87_v0
				}
			})(_59_v0)
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
		_86_v22 = _nw8
		var _89_v23 D7
		_ = _89_v23
		_89_v23 = Companion_D7_.Create_DC22_(_86_v22, _85_v21)
		_59_v0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("vuhemdl"), (_89_v23).Dtor_cf39())
	}
	var _90_v24 _dafny.Array
	_ = _90_v24
	var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
	_ = _nw9
	_90_v24 = _nw9
	var _91_v25 D5
	_ = _91_v25
	_91_v25 = Companion_D5_.Create_DC16_(_90_v24)
	var _92_v26 _dafny.Map
	_ = _92_v26
	_92_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v25, _81_v18)
	var _93_v27 _dafny.MultiSet
	_ = _93_v27
	_93_v27 = _dafny.MultiSetOf(_92_v26, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v25, _81_v18), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D5_.Create_DC16_(_90_v24), _81_v18))
	_93_v27 = (_93_v27).Intersection(_93_v27)
	var _94_v28 _dafny.Sequence
	_ = _94_v28
	_94_v28 = _dafny.SeqOf(!(_59_v0))
	var _95_v29 _dafny.Set
	_ = _95_v29
	_95_v29 = _dafny.SetOf(_94_v28)
	var _96_v30 _dafny.MultiSet
	_ = _96_v30
	_96_v30 = _dafny.MultiSetOf(_59_v0)
	var _hi1 _dafny.Int = (_96_v30).Cardinality()
	_ = _hi1
	for _97_i4 := ((_95_v29).Union(_95_v29)).Cardinality(); _97_i4.Cmp(_hi1) < 0; _97_i4 = _97_i4.Plus(_dafny.One) {
		var _98_v31 D4
		_ = _98_v31
		_98_v31 = Companion_D4_.Create_DC14_(_81_v18, _59_v0, _97_i4, _59_v0, _97_i4)
		_96_v30 = _dafny.MultiSetOf(!((func(_pat_let0_0 D4) D4 {
			return func(_99_dt__update__tmp_h0 D4) D4 {
				return func(_pat_let1_0 _dafny.Int) D4 {
					return func(_100_dt__update_hcf20_h0 _dafny.Int) D4 {
						return func(_pat_let2_0 _dafny.Int) D4 {
							return func(_101_dt__update_hcf24_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC14_(_100_dt__update_hcf20_h0, (_99_dt__update__tmp_h0).Dtor_cf21(), (_99_dt__update__tmp_h0).Dtor_cf22(), (_99_dt__update__tmp_h0).Dtor_cf23(), _101_dt__update_hcf24_h0)
							}(_pat_let2_0)
						}(_97_i4)
					}(_pat_let1_0)
				}(_97_i4)
			}(_pat_let0_0)
		}(_98_v31)).Dtor_cf21()))
		(globalState).F2 = _81_v18
		var _102_v32 _dafny.Map
		_ = _102_v32
		_102_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_97_i4, _59_v0)
		var _103_v33 _dafny.Sequence
		_ = _103_v33
		_103_v33 = _dafny.SeqOf(_97_i4, _81_v18)
		r1 = Companion_Default___.Fm38(_102_v32, _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_103_v33).Cardinality())), _97_i4), true, (_81_v18).Cmp(_97_i4) > 0, globalState)
		var _104_v34 *C1
		_ = _104_v34
		var _nw10 *C1 = New_C1_()
		_ = _nw10
		_nw10.Ctor__(_dafny.IntOfInt64(-663), Companion_Default___.Fm37(_59_v0, _81_v18, _59_v0, globalState))
		_104_v34 = _nw10
	}
	var _105_v35 _dafny.Sequence
	_ = _105_v35
	_105_v35 = _dafny.UnicodeSeqOfUtf8Bytes("oyt")
	_105_v35 = _dafny.Companion_Sequence_.Concatenate(_105_v35, _105_v35)
	var _106_v36 D0
	_ = _106_v36
	_106_v36 = Companion_D0_.Create_DC0_(_59_v0)
	var _107_v37 D10
	_ = _107_v37
	_107_v37 = Companion_D10_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_81_v18, _dafny.IntOfUint32((_105_v35).Cardinality())))
	var _rhs8 D0 = _106_v36
	_ = _rhs8
	var _rhs9 D10 = _107_v37
	_ = _rhs9
	var _rhs10 bool = _59_v0
	_ = _rhs10
	var _rhs11 bool = _59_v0
	_ = _rhs11
	_106_v36 = _rhs8
	_107_v37 = _rhs9
	r1 = _rhs10
	_59_v0 = _rhs11
	var _108_v38 _dafny.Map
	_ = _108_v38
	_108_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_105_v35, _105_v35), _90_v24)
	r0 = (func() _dafny.Array {
		if (_108_v38).Contains(_dafny.Companion_Sequence_.Concatenate(_105_v35, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("kcykgfojv"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-354), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kcykgfojv")).Cardinality()))).Uint32(), _dafny.CodePoint('r')))) {
			return (_108_v38).Get(_dafny.Companion_Sequence_.Concatenate(_105_v35, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("kcykgfojv"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-354), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kcykgfojv")).Cardinality()))).Uint32(), _dafny.CodePoint('r')))).(_dafny.Array)
		}
		return _90_v24
	})()
	r1 = _59_v0
	var _109_v39 _dafny.MultiSet
	_ = _109_v39
	_109_v39 = _dafny.MultiSetOf(_105_v35, _105_v35)
	r2 = _109_v39
	var _110_v40 _dafny.MultiSet
	_ = _110_v40
	_110_v40 = _dafny.MultiSetOf(_81_v18)
	var _111_v41 _dafny.Sequence
	_ = _111_v41
	_111_v41 = _dafny.SeqOf(_dafny.IntOfInt64(747), (_dafny.Zero).Minus((func() _dafny.Int {
		if (_110_v40).Contains(_81_v18) {
			return (_110_v40).Multiplicity(_81_v18)
		}
		return _81_v18
	})()))
	r3 = _dafny.IntOfUint32((_111_v41).Cardinality())
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _112_v0 _dafny.Int
	_ = _112_v0
	_112_v0 = _dafny.IntOfInt64(-723)
	var _113_v1 _dafny.CodePoint
	_ = _113_v1
	_113_v1 = _dafny.CodePoint('i')
	var _114_v2 bool
	_ = _114_v2
	_114_v2 = true
	var _115_v3 _dafny.Map
	_ = _115_v3
	_115_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v1, _114_v2)
	var _116_v4 _dafny.MultiSet
	_ = _116_v4
	_116_v4 = _dafny.MultiSetOf(_112_v0, _112_v0, _112_v0)
	var _117_v5 _dafny.Map
	_ = _117_v5
	_117_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v1, _116_v4)
	var _118_globalState *GlobalState
	_ = _118_globalState
	var _nw11 *GlobalState = New_GlobalState_()
	_ = _nw11
	_nw11.Ctor__(_dafny.IntOfInt64(275), _dafny.IntOfInt64(628), _dafny.IntOfInt64(-930), _dafny.IntOfInt64(254), (_dafny.MultiSetOf(_112_v0, (_dafny.Zero).Minus((_115_v3).Cardinality()), _112_v0, _112_v0, _112_v0)).Union((func() _dafny.MultiSet {
		if (_117_v5).Contains(_113_v1) {
			return (_117_v5).Get(_113_v1).(_dafny.MultiSet)
		}
		return _116_v4
	})()))
	_118_globalState = _nw11
	if _114_v2 {
		var _119_v6 _dafny.Array
		_ = _119_v6
		var _120_v7 bool
		_ = _120_v7
		var _121_v8 _dafny.MultiSet
		_ = _121_v8
		var _122_v9 _dafny.Int
		_ = _122_v9
		var _out0 _dafny.Array
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.MultiSet
		_ = _out2
		var _out3 _dafny.Int
		_ = _out3
		_out0, _out1, _out2, _out3 = Companion_Default___.M0(_113_v1, _118_globalState)
		_119_v6 = _out0
		_120_v7 = _out1
		_121_v8 = _out2
		_122_v9 = _out3
		var _123_v10 *C3
		_ = _123_v10
		var _nw12 *C3 = New_C3_()
		_ = _nw12
		_nw12.Ctor__(_120_v7, _112_v0, _122_v9)
		_123_v10 = _nw12
		var _124_v11 *C1
		_ = _124_v11
		var _nw13 *C1 = New_C1_()
		_ = _nw13
		_nw13.Ctor__(Companion_Default___.SafeModuloInt(_122_v9, _122_v9), _112_v0)
		_124_v11 = _nw13
		var _nw14 *C1 = New_C1_()
		_ = _nw14
		_nw14.Ctor__(Companion_Default___.SafeDivisionInt(_122_v9, (_dafny.Zero).Minus(_112_v0)), _112_v0)
		_124_v11 = _nw14
		var _125_v12 _dafny.Sequence
		_ = _125_v12
		_125_v12 = _dafny.UnicodeSeqOfUtf8Bytes("a")
		var _126_v13 _dafny.Map
		_ = _126_v13
		_126_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.IsProperPrefixOf(_125_v12, _125_v12), _112_v0)
		var _127_v14 _dafny.MultiSet
		_ = _127_v14
		_127_v14 = _dafny.MultiSetOf(_124_v11)
		_126_v13 = (_126_v13).Update((!((_123_v10).F7())) || (_114_v2), (func() _dafny.Int {
			if (_127_v14).Contains(_124_v11) {
				return (_127_v14).Multiplicity(_124_v11)
			}
			return _112_v0
		})())
		_120_v7 = !(_120_v7)
	} else {
		var _128_v15 _dafny.Sequence
		_ = _128_v15
		_128_v15 = _dafny.UnicodeSeqOfUtf8Bytes("mw")
		var _129_v16 _dafny.Map
		_ = _129_v16
		_129_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(933), _112_v0)
		var _130_v17 _dafny.Sequence
		_ = _130_v17
		_130_v17 = _dafny.SeqOf(_114_v2)
		var _131_v18 _dafny.MultiSet
		_ = _131_v18
		_131_v18 = _dafny.MultiSetOf(_114_v2, _114_v2)
		var _132_v19 _dafny.MultiSet
		_ = _132_v19
		_132_v19 = _131_v18
		var _133_v20 _dafny.Array
		_ = _133_v20
		var _nwElement0_1 _dafny.Int = _dafny.IntOfUint32((_128_v15).Cardinality())
		_ = _nwElement0_1
		var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
		_ = _nw15
		(_nw15).ArraySet1(_nwElement0_1, 0)
		(_nw15).ArraySet1((_129_v16).Cardinality(), 1)
		(_nw15).ArraySet1(Companion_Default___.SafeModuloInt(_112_v0, _112_v0), 2)
		(_nw15).ArraySet1(Companion_Default___.SafeDivisionInt(_112_v0, (_dafny.SetOf(_112_v0)).Cardinality()), 3)
		(_nw15).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_130_v17, _130_v17)).Cardinality()), 4)
		(_nw15).ArraySet1(_112_v0, 5)
		(_nw15).ArraySet1(_dafny.IntOfInt64(-903), 6)
		(_nw15).ArraySet1((_112_v0).Plus((_116_v4).Cardinality()), 7)
		(_nw15).ArraySet1(_112_v0, 8)
		(_nw15).ArraySet1(_112_v0, 9)
		(_nw15).ArraySet1(_112_v0, 10)
		(_nw15).ArraySet1(_dafny.IntOfUint32((_128_v15).Cardinality()), 11)
		(_nw15).ArraySet1(_112_v0, 12)
		(_nw15).ArraySet1((Companion_Default___.Fm35(true, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(813))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}((func(_134_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_135_i0 _dafny.Int) _dafny.Int {
				return _134_v0
			}
		})(_112_v0)))).Cardinality()), _112_v0, _112_v0, _118_globalState)).Cardinality(), 13)
		(_nw15).ArraySet1((_132_v19).Cardinality(), 14)
		(_nw15).ArraySet1(_dafny.IntOfInt64(281), 15)
		(_nw15).ArraySet1(_112_v0, 16)
		(_nw15).ArraySet1(_112_v0, 17)
		(_nw15).ArraySet1(_dafny.IntOfInt64(-225), 18)
		(_nw15).ArraySet1(_dafny.IntOfUint32((_130_v17).Cardinality()), 19)
		(_nw15).ArraySet1(_112_v0, 20)
		_133_v20 = _nw15
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
		_ = _index2
		(_133_v20).ArraySet1(_112_v0, (_index2).Int())
		var _136_v21 D0
		_ = _136_v21
		_136_v21 = Companion_D0_.Create_DC2_()
		var _137_v22 D3
		_ = _137_v22
		_137_v22 = Companion_D3_.Create_DC9_(_116_v4, _136_v21, _112_v0)
		var _138_v23 _dafny.Array
		_ = _138_v23
		var _nwElement0_2 D3 = _137_v22
		_ = _nwElement0_2
		var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.One)
		_ = _nw16
		(_nw16).ArraySet1(_nwElement0_2, 0)
		_138_v23 = _nw16
		var _pat_let_tv0 = _136_v21
		_ = _pat_let_tv0
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_138_v23), 0))
		_ = _index3
		(_138_v23).ArraySet1(func(_pat_let3_0 D3) D3 {
			return func(_139_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let4_0 D0) D3 {
					return func(_140_dt__update_hcf10_h0 D0) D3 {
						return Companion_D3_.Create_DC9_((_139_dt__update__tmp_h0).Dtor_cf9(), _140_dt__update_hcf10_h0, (_139_dt__update__tmp_h0).Dtor_cf11())
					}(_pat_let4_0)
				}(_pat_let_tv0)
			}(_pat_let3_0)
		}(Companion_Default___.Fm36(Companion_D5_.Create_DC19_(_112_v0, _113_v1, _112_v0), _113_v1, _114_v2, _118_globalState)), (_index3).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
		_ = _index4
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_138_v23), 0))
		_ = _index5
		var _rhs12 bool = !(_114_v2)
		_ = _rhs12
		var _rhs13 _dafny.Int = ((_dafny.IntOfInt64(524)).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_128_v15, (Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_128_v15).Cardinality()))).Uint32(), _113_v1)).Cardinality()))).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_130_v17).Cardinality()))))
		_ = _rhs13
		var _rhs14 D3 = _137_v22
		_ = _rhs14
		var _lhs3 _dafny.Array = _133_v20
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
		_ = _lhs4
		var _lhs5 _dafny.Array = _138_v23
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_138_v23), 0))
		_ = _lhs6
		_114_v2 = _rhs12
		(_lhs3).ArraySet1(_rhs13, (_lhs4).Int())
		(_lhs5).ArraySet1(_rhs14, (_lhs6).Int())
		var _141_v24 D7
		_ = _141_v24
		_141_v24 = Companion_D7_.Create_DC23_((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), _128_v15, (_129_v16).Cardinality())
		var _142_v25 _dafny.MultiSet
		_ = _142_v25
		_142_v25 = _dafny.MultiSetOf(_128_v15, (_141_v24).Dtor_cf41(), _128_v15)
		if (_142_v25).Equals(_142_v25) {
			var _143_v26 _dafny.Array
			_ = _143_v26
			var _144_v27 bool
			_ = _144_v27
			var _145_v28 _dafny.MultiSet
			_ = _145_v28
			var _146_v29 _dafny.Int
			_ = _146_v29
			var _out4 _dafny.Array
			_ = _out4
			var _out5 bool
			_ = _out5
			var _out6 _dafny.MultiSet
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0(_113_v1, _118_globalState)
			_143_v26 = _out4
			_144_v27 = _out5
			_145_v28 = _out6
			_146_v29 = _out7
			var _147_v30 *C3
			_ = _147_v30
			var _nw17 *C3 = New_C3_()
			_ = _nw17
			_nw17.Ctor__(_144_v27, _112_v0, _112_v0)
			_147_v30 = _nw17
			var _148_v31 _dafny.Map
			_ = _148_v31
			_148_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(608)), _147_v30)
			var _149_v32 _dafny.Array
			_ = _149_v32
			var _nwElement0_3 bool = (_147_v30).F7()
			_ = _nwElement0_3
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(3))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_3, 0)
			(_nw18).ArraySet1((func() bool {
				if (_115_v3).Contains(_113_v1) {
					return (_115_v3).Get(_113_v1).(bool)
				}
				return (_147_v30).F7()
			})(), 1)
			(_nw18).ArraySet1(_144_v27, 2)
			_149_v32 = _nw18
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_149_v32), 0))
			_ = _index6
			(_149_v32).ArraySet1(!(_114_v2), (_index6).Int())
			var _150_v33 _dafny.Map
			_ = _150_v33
			_150_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v1, _128_v15)
			var _151_v34 _dafny.Map
			_ = _151_v34
			_151_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_150_v33).Contains(_113_v1) {
					return (_150_v33).Get(_113_v1).(_dafny.Sequence)
				}
				return _128_v15
			})()).Cardinality()), (_147_v30).F7())
			var _152_v35 _dafny.MultiSet
			_ = _152_v35
			_152_v35 = _dafny.MultiSetOf(_151_v34, _151_v34)
			var _153_v36 _dafny.Set
			_ = _153_v36
			_153_v36 = _dafny.SetOf(_114_v2)
			var _154_v37 _dafny.Sequence
			_ = _154_v37
			_154_v37 = _dafny.SeqOf((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), (_dafny.SetOf(true, !(_114_v2))).Cardinality())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_149_v32), 0))
			_ = _index7
			var _rhs15 _dafny.Map = _148_v31
			_ = _rhs15
			var _rhs16 bool = (_147_v30).Fm0((_152_v35).Difference(_152_v35), _114_v2, _128_v15, _118_globalState)
			_ = _rhs16
			var _rhs17 bool = ((_147_v30).Fm2(_dafny.IntOfInt64(209), _153_v36, _118_globalState)).Cmp(_dafny.IntOfUint32((_154_v37).Cardinality())) >= 0
			_ = _rhs17
			var _rhs18 _dafny.Int = _146_v29
			_ = _rhs18
			var _lhs7 _dafny.Array = _149_v32
			_ = _lhs7
			var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen((_149_v32), 0))
			_ = _lhs8
			var _lhs9 *GlobalState = _118_globalState
			_ = _lhs9
			_148_v31 = _rhs15
			(_lhs7).ArraySet1(_rhs16, (_lhs8).Int())
			_144_v27 = _rhs17
			_lhs9.F2 = _rhs18
			var _155_v38 *C3
			_ = _155_v38
			var _nw19 *C3 = New_C3_()
			_ = _nw19
			_nw19.Ctor__(_114_v2, (_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), _112_v0)
			_155_v38 = _nw19
			var _156_v39 *C1
			_ = _156_v39
			var _nw20 *C1 = New_C1_()
			_ = _nw20
			_nw20.Ctor__((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)))
			_156_v39 = _nw20
			var _157_v40 _dafny.Sequence
			_ = _157_v40
			_157_v40 = _dafny.SeqOf(_149_v32, _149_v32, _149_v32, _149_v32)
			var _158_v41 D7
			_ = _158_v41
			_158_v41 = Companion_D7_.Create_DC22_(_149_v32, _128_v15)
			var _159_v42 _dafny.Array
			_ = _159_v42
			var _nwElement0_4 _dafny.Array = (_157_v40).Select((Companion_Default___.SafeIndex((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_157_v40).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _nwElement0_4
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(13))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_4, 0)
			(_nw21).ArraySet1(_149_v32, 1)
			(_nw21).ArraySet1(_149_v32, 2)
			(_nw21).ArraySet1(_149_v32, 3)
			(_nw21).ArraySet1((_157_v40).Select((Companion_Default___.SafeIndex((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_157_v40).Cardinality()))).Uint32()).(_dafny.Array), 4)
			(_nw21).ArraySet1(_149_v32, 5)
			(_nw21).ArraySet1((_158_v41).Dtor_cf38(), 6)
			(_nw21).ArraySet1(_149_v32, 7)
			(_nw21).ArraySet1(_149_v32, 8)
			(_nw21).ArraySet1((_158_v41).Dtor_cf38(), 9)
			(_nw21).ArraySet1(_149_v32, 10)
			(_nw21).ArraySet1(_149_v32, 11)
			(_nw21).ArraySet1(_149_v32, 12)
			_159_v42 = _nw21
			var _160_v43 _dafny.Sequence
			_ = _160_v43
			_160_v43 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(167))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}((func(_161_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_162_i1 _dafny.Int) _dafny.CodePoint {
					return _161_v1
				}
			})(_113_v1))))
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
			_ = _index8
			var _rhs19 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_146_v29, _dafny.IntOfUint32((_dafny.SeqOf(_114_v2, (_155_v38).F7(), _114_v2, (_155_v38).F7(), true)).Cardinality())))
			_ = _rhs19
			var _rhs20 _dafny.Int = (_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)
			_ = _rhs20
			var _rhs21 _dafny.Array = _159_v42
			_ = _rhs21
			var _rhs22 _dafny.Int = _dafny.IntOfUint32(((_160_v43).Select((Companion_Default___.SafeIndex((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_160_v43).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
			_ = _rhs22
			var _lhs10 _dafny.Array = _133_v20
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
			_ = _lhs11
			var _lhs12 *GlobalState = _118_globalState
			_ = _lhs12
			(_lhs10).ArraySet1(_rhs19, (_lhs11).Int())
			_lhs12.F2 = _rhs20
			_159_v42 = _rhs21
			_112_v0 = _rhs22
		} else {
			(_118_globalState).F2 = _112_v0
			var _163_v44 *C3
			_ = _163_v44
			var _nw22 *C3 = New_C3_()
			_ = _nw22
			_nw22.Ctor__(_114_v2, ((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfInt64(11)), (_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int))
			_163_v44 = _nw22
			var _164_v45 *C3
			_ = _164_v45
			_164_v45 = _163_v44
			_163_v44 = (_164_v45)
			var _165_v46 _dafny.Array
			_ = _165_v46
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
			_ = _nw23
			_165_v46 = _nw23
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_165_v46), 0))
			_ = _index9
			(_165_v46).ArraySet1(_131_v18, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_165_v46), 0))
			_ = _index10
			var _rhs23 _dafny.Int = (_dafny.Zero).Minus((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int))
			_ = _rhs23
			var _rhs24 _dafny.MultiSet = (_131_v18).Union(_131_v18)
			_ = _rhs24
			var _lhs13 _dafny.Array = _165_v46
			_ = _lhs13
			var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_165_v46), 0))
			_ = _lhs14
			_112_v0 = _rhs23
			(_lhs13).ArraySet1(_rhs24, (_lhs14).Int())
			var _166_v47 _dafny.Array
			_ = _166_v47
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_2
			var _nw24 _dafny.Array
			_ = _nw24
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw24 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_167_v44 *C3) func(_dafny.Int) bool {
					return func(_168_i2 _dafny.Int) bool {
						return (_167_v44).F7()
					}
				})(_163_v44)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw24 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw24).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw24).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_166_v47 = _nw24
			var _169_v48 _dafny.Map
			_ = _169_v48
			_169_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_113_v1, _166_v47)
			var _170_v49 _dafny.Map
			_ = _170_v49
			_170_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_169_v48).Contains(_113_v1) {
					return (_169_v48).Get(_113_v1).(_dafny.Array)
				}
				return _166_v47
			})(), (_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int))
			_170_v49 = (_170_v49).Merge((_170_v49).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v47, _dafny.IntOfUint32((_128_v15).Cardinality()))))
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))
			_ = _index11
			(_133_v20).ArraySet1((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int), (_index11).Int())
		}
		var _171_v50 *C0
		_ = _171_v50
		var _nw25 *C0 = New_C0_()
		_ = _nw25
		_nw25.Ctor__(_dafny.IntOfInt64(-984))
		_171_v50 = _nw25
		var _172_v51 _dafny.MultiSet
		_ = _172_v51
		_172_v51 = _dafny.MultiSetOf(_171_v50)
		var _173_v52 T1
		_ = _173_v52
		var _nw26 *C1 = New_C1_()
		_ = _nw26
		_nw26.Ctor__(((func() _dafny.Int {
			if (_172_v51).Contains(_171_v50) {
				return (_172_v51).Multiplicity(_171_v50)
			}
			return _dafny.IntOfUint32((_128_v15).Cardinality())
		})()).Times((_171_v50).F8()), ((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)).Plus(_112_v0))
		_173_v52 = _nw26
		_173_v52 = _173_v52
		var _174_v53 _dafny.Sequence
		_ = _174_v53
		_174_v53 = _dafny.SeqOf(_112_v0, (func() _dafny.Int {
			if (_129_v16).Contains((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)) {
				return (_129_v16).Get((_133_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_133_v20), 0))).Int()).(_dafny.Int)).(_dafny.Int)
			}
			return (_173_v52).F5()
		})(), (_173_v52).F5())
		var _175_v54 *C3
		_ = _175_v54
		var _nw27 *C3 = New_C3_()
		_ = _nw27
		_nw27.Ctor__(_114_v2, _dafny.IntOfUint32((_174_v53).Cardinality()), (_171_v50).F8())
		_175_v54 = _nw27
		_128_v15 = _128_v15
	}
	var _176_i3 _dafny.Int
	_ = _176_i3
	_176_i3 = _dafny.Zero
	{
		for _114_v2 {
			{
				if (_176_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_176_i3 = (_176_i3).Plus(_dafny.One)
				var _rhs25 bool = (func() bool {
					if _114_v2 {
						return _114_v2
					}
					return !((_116_v4).IsProperSubsetOf(_116_v4))
				})()
				_ = _rhs25
				var _rhs26 bool = _114_v2
				_ = _rhs26
				var _rhs27 bool = _114_v2
				_ = _rhs27
				_114_v2 = _rhs25
				_114_v2 = _rhs26
				_114_v2 = _rhs27
				var _177_v55 _dafny.MultiSet
				_ = _177_v55
				_177_v55 = _dafny.MultiSetOf(_114_v2, _114_v2)
				var _178_v56 _dafny.Map
				_ = _178_v56
				_178_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v2, _112_v0)
				(_118_globalState).F2 = (func() _dafny.Int {
					if (_177_v55).Contains((func() bool {
						if (_115_v3).Contains(_113_v1) {
							return (_115_v3).Get(_113_v1).(bool)
						}
						return _114_v2
					})()) {
						return (_177_v55).Multiplicity((func() bool {
							if (_115_v3).Contains(_113_v1) {
								return (_115_v3).Get(_113_v1).(bool)
							}
							return _114_v2
						})())
					}
					return ((func() _dafny.Int {
						if (_178_v56).Contains(_114_v2) {
							return (_178_v56).Get(_114_v2).(_dafny.Int)
						}
						return _112_v0
					})()).Plus(_112_v0)
				})()
				_114_v2 = !(_114_v2)
				var _179_v57 _dafny.Sequence
				_ = _179_v57
				_179_v57 = _dafny.UnicodeSeqOfUtf8Bytes("lqy")
				var _180_v58 _dafny.Map
				_ = _180_v58
				_180_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v2, _113_v1)
				var _181_v59 _dafny.Sequence
				_ = _181_v59
				_181_v59 = _dafny.SeqOf(_114_v2)
				_179_v57 = Companion_Default___.Fm7((func() _dafny.CodePoint {
					if (_180_v58).Contains(_114_v2) {
						return (_180_v58).Get(_114_v2).(_dafny.CodePoint)
					}
					return _113_v1
				})(), _114_v2, (_dafny.IntOfUint32((_181_v59).Cardinality())).Cmp((_dafny.Zero).Minus(_112_v0)) == 0, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_112_v0), _112_v0), _118_globalState)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _182_v60 _dafny.Sequence
	_ = _182_v60
	_182_v60 = _dafny.SeqOf(false)
	var _183_v61 _dafny.Set
	_ = _183_v61
	_183_v61 = _dafny.SetOf(!((_182_v60).Select((Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_182_v60).Cardinality()))).Uint32()).(bool)))
	var _184_v62 _dafny.Sequence
	_ = _184_v62
	_184_v62 = _dafny.UnicodeSeqOfUtf8Bytes("uapcj")
	var _185_v63 _dafny.Array
	_ = _185_v63
	var _nwElement0_5 _dafny.Int = (func() _dafny.Int {
		if _114_v2 {
			return _112_v0
		}
		return (_dafny.Zero).Minus(_112_v0)
	})()
	_ = _nwElement0_5
	var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(9))
	_ = _nw28
	(_nw28).ArraySet1(_nwElement0_5, 0)
	(_nw28).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus(_112_v0)).Times(_112_v0)), 1)
	(_nw28).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(804))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}((func(_186_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_187_i5 _dafny.Int) _dafny.CodePoint {
			return _186_v1
		}
	})(_113_v1)))).Cardinality()), 2)
	(_nw28).ArraySet1(_112_v0, 3)
	(_nw28).ArraySet1((_183_v61).Cardinality(), 4)
	(_nw28).ArraySet1(_112_v0, 5)
	(_nw28).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_184_v62, _dafny.UnicodeSeqOfUtf8Bytes("qmhnk"))).Cardinality()), 6)
	(_nw28).ArraySet1(((_183_v61).Intersection(_183_v61)).Cardinality(), 7)
	(_nw28).ArraySet1(_112_v0, 8)
	_185_v63 = _nw28
	for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_185_v63), 0))); ; {
		_guard_loop_0, _ok18 := _iter18()
		if !_ok18 {
			break
		}
		var _188_i4 _dafny.Int
		_188_i4 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_188_i4).Sign() != -1) && ((_188_i4).Cmp(_dafny.ArrayLen((_185_v63), 0)) < 0)) {
			(_185_v63).ArraySet1((_188_i4).Plus(_dafny.IntOfUint32((_182_v60).Cardinality())), (_188_i4).Int())
		}
	}
	var _189_v64 D1
	_ = _189_v64
	_189_v64 = Companion_D1_.Create_DC6_(_184_v62, (_dafny.Zero).Minus(((func() _dafny.Set {
		if _114_v2 {
			return _183_v61
		}
		return _183_v61
	})()).Cardinality()))
	var _source3 D1 = _189_v64
	_ = _source3
	if _source3.Is_DC5() {
		var _rhs28 _dafny.CodePoint = _113_v1
		_ = _rhs28
		var _rhs29 bool = _114_v2
		_ = _rhs29
		_113_v1 = _rhs28
		_114_v2 = _rhs29
		_112_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-266), (_112_v0).Plus(_112_v0)))
		var _190_v65 _dafny.Array
		_ = _190_v65
		var _191_v66 bool
		_ = _191_v66
		var _192_v67 _dafny.MultiSet
		_ = _192_v67
		var _193_v68 _dafny.Int
		_ = _193_v68
		var _out8 _dafny.Array
		_ = _out8
		var _out9 bool
		_ = _out9
		var _out10 _dafny.MultiSet
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out8, _out9, _out10, _out11 = Companion_Default___.M0(_113_v1, _118_globalState)
		_190_v65 = _out8
		_191_v66 = _out9
		_192_v67 = _out10
		_193_v68 = _out11
		_191_v66 = true
	} else if _source3.Is_DC6() {
		var _194___mcc_h0 _dafny.Sequence = _source3.Get_().(D1_DC6).Cf5
		_ = _194___mcc_h0
		var _195___mcc_h1 _dafny.Int = _source3.Get_().(D1_DC6).Cf6
		_ = _195___mcc_h1
		var _196_cf6 _dafny.Int = _195___mcc_h1
		_ = _196_cf6
		var _197_cf5 _dafny.Sequence = _194___mcc_h0
		_ = _197_cf5
		_114_v2 = (func() bool {
			if _114_v2 {
				return (_183_v61).IsSubsetOf(_183_v61)
			}
			return false
		})()
		var _198_v69 *C1
		_ = _198_v69
		var _nw29 *C1 = New_C1_()
		_ = _nw29
		_nw29.Ctor__(((_dafny.SetOf(Companion_Default___.Fm37(_114_v2, _112_v0, _114_v2, _118_globalState))).Difference(_dafny.SetOf(_112_v0, _112_v0, _112_v0, _dafny.IntOfInt64(678), _196_cf6))).Cardinality(), _196_cf6)
		_198_v69 = _nw29
		_198_v69 = _198_v69
		var _199_v70 D9
		_ = _199_v70
		_199_v70 = Companion_D9_.Create_DC28_(_114_v2, _196_cf6, _112_v0)
		_114_v2 = !((_199_v70).Dtor_cf46())
		var _200_v71 _dafny.Array
		_ = _200_v71
		var _nw30 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
		_ = _nw30
		_200_v71 = _nw30
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_200_v71), 0))
		_ = _index12
		(_200_v71).ArraySet1(_114_v2, (_index12).Int())
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_200_v71), 0))
		_ = _index13
		(_200_v71).ArraySet1(_114_v2, (_index13).Int())
	} else {
		var _201___mcc_h2 _dafny.Set = _source3.Get_().(D1_DC4).Cf4
		_ = _201___mcc_h2
		var _202_cf4 _dafny.Set = _201___mcc_h2
		_ = _202_cf4
		if (func() bool {
			if _114_v2 {
				return (_112_v0).Cmp(_112_v0) <= 0
			}
			return _114_v2
		})() {
			var _203_v72 _dafny.Sequence
			_ = _203_v72
			_203_v72 = _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm7(_113_v1, _114_v2, _114_v2, _dafny.IntOfInt64(373), _118_globalState)).Cardinality()))
			_203_v72 = _203_v72
			var _204_v73 _dafny.Map
			_ = _204_v73
			_204_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v62, _114_v2)
			_204_v73 = (_204_v73).Update(_dafny.UnicodeSeqOfUtf8Bytes("xajwiulrn"), _114_v2)
			_114_v2 = false
			(_118_globalState).F2 = _112_v0
			var _205_v74 _dafny.Array
			_ = _205_v74
			var _nw31 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
			_ = _nw31
			_205_v74 = _nw31
			var _206_v75 _dafny.Map
			_ = _206_v75
			_206_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v0, true)
			var _207_v76 _dafny.MultiSet
			_ = _207_v76
			_207_v76 = _dafny.MultiSetOf((func() bool {
				if (_206_v75).Contains((_202_cf4).Cardinality()) {
					return (_206_v75).Get((_202_cf4).Cardinality()).(bool)
				}
				return true
			})())
			var _208_v77 _dafny.Map
			_ = _208_v77
			_208_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v2, _114_v2)
			var _209_v78 _dafny.Map
			_ = _209_v78
			_209_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Int {
				if (_207_v76).Contains(_114_v2) {
					return (_207_v76).Multiplicity(_114_v2)
				}
				return (_208_v77).Cardinality()
			})())
			var _210_v79 _dafny.Map
			_ = _210_v79
			_210_v79 = (_209_v78).Merge(_209_v78)
			var _rhs30 _dafny.Array = _205_v74
			_ = _rhs30
			var _rhs31 _dafny.Int = (_dafny.Zero).Minus(_112_v0)
			_ = _rhs31
			var _rhs32 _dafny.Map = _209_v78
			_ = _rhs32
			var _rhs33 bool = (Companion_Default___.SafeDivisionInt(_112_v0, _112_v0)).Cmp(_112_v0) == 0
			_ = _rhs33
			var _lhs15 *GlobalState = _118_globalState
			_ = _lhs15
			_205_v74 = _rhs30
			_lhs15.F2 = _rhs31
			_210_v79 = _rhs32
			_114_v2 = _rhs33
		} else {
			_114_v2 = _114_v2
			var _211_v80 _dafny.Sequence
			_ = _211_v80
			_211_v80 = _dafny.SeqOf(_112_v0)
			var _212_v81 _dafny.Map
			_ = _212_v81
			_212_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v80, _112_v0)
			var _213_v82 _dafny.Map
			_ = _213_v82
			_213_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v62, _112_v0)
			_112_v0 = ((_212_v81).Update(_dafny.Companion_Sequence_.Concatenate(_211_v80, _dafny.SeqOf(_112_v0)), (func() _dafny.Int {
				if (_213_v82).Contains(_dafny.Companion_Sequence_.Update(_184_v62, (Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_184_v62).Cardinality()))).Uint32(), _113_v1)) {
					return (_213_v82).Get(_dafny.Companion_Sequence_.Update(_184_v62, (Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_184_v62).Cardinality()))).Uint32(), _113_v1)).(_dafny.Int)
				}
				return _112_v0
			})())).Cardinality()
			_182_v60 = _182_v60
			var _214_v83 *C1
			_ = _214_v83
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__(_dafny.IntOfInt64(896), _dafny.IntOfInt64(501))
			_214_v83 = _nw32
			var _215_v84 _dafny.Map
			_ = _215_v84
			_215_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v0, _214_v83)
			var _216_v85 _dafny.Array
			_ = _216_v85
			var _nwElement0_6 *C1 = _214_v83
			_ = _nwElement0_6
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(6))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_6, 0)
			(_nw33).ArraySet1((func() *C1 {
				if (_215_v84).Contains(_112_v0) {
					return (_215_v84).Get(_112_v0).(*C1)
				}
				return _214_v83
			})(), 1)
			(_nw33).ArraySet1(_214_v83, 2)
			(_nw33).ArraySet1(_214_v83, 3)
			(_nw33).ArraySet1(_214_v83, 4)
			(_nw33).ArraySet1(_214_v83, 5)
			_216_v85 = _nw33
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_216_v85), 0))
			_ = _index14
			(_216_v85).ArraySet1(_214_v83, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_216_v85), 0))
			_ = _index15
			(_216_v85).ArraySet1(_214_v83, (_index15).Int())
			var _217_v86 _dafny.Map
			_ = _217_v86
			_217_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(602), _114_v2)
			var _218_v87 _dafny.MultiSet
			_ = _218_v87
			_218_v87 = _dafny.MultiSetOf(_217_v86)
			(_118_globalState).F2 = (Companion_Default___.SafeModuloInt(_112_v0, _112_v0)).Plus(Companion_Default___.Fm37(_114_v2, _112_v0, (_214_v83).Fm0(_218_v87, _114_v2, _184_v62, _118_globalState), _118_globalState))
		}
		var _219_v88 _dafny.Array
		_ = _219_v88
		var _220_v89 bool
		_ = _220_v89
		var _221_v90 _dafny.MultiSet
		_ = _221_v90
		var _222_v91 _dafny.Int
		_ = _222_v91
		var _out12 _dafny.Array
		_ = _out12
		var _out13 bool
		_ = _out13
		var _out14 _dafny.MultiSet
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out12, _out13, _out14, _out15 = Companion_Default___.M0(_113_v1, _118_globalState)
		_219_v88 = _out12
		_220_v89 = _out13
		_221_v90 = _out14
		_222_v91 = _out15
		var _223_v92 _dafny.Array
		_ = _223_v92
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_3
		var _nw34 _dafny.Array
		_ = _nw34
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw34 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_224_v89 bool) func(_dafny.Int) bool {
				return func(_225_i6 _dafny.Int) bool {
					return _224_v89
				}
			})(_220_v89)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw34 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw34).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw34).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_223_v92 = _nw34
		var _226_v93 D7
		_ = _226_v93
		_226_v93 = Companion_D7_.Create_DC22_(_223_v92, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_227_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_228_i7 _dafny.Int) _dafny.CodePoint {
				return _227_v1
			}
		})(_113_v1))))
		var _229_v94 _dafny.Sequence
		_ = _229_v94
		_229_v94 = _dafny.SeqOf(_223_v92, (_226_v93).Dtor_cf38(), _223_v92)
		(_118_globalState).F2 = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _114_v2 {
				return _229_v94
			}
			return _229_v94
		})()).Cardinality())
		var _230_v95 *C1
		_ = _230_v95
		var _nw35 *C1 = New_C1_()
		_ = _nw35
		_nw35.Ctor__(_dafny.IntOfInt64(220), _222_v91)
		_230_v95 = _nw35
		_230_v95 = _230_v95
	}
	_112_v0 = ((_dafny.Zero).Minus(_112_v0)).Plus(_112_v0)
	var _231_v96 _dafny.Map
	_ = _231_v96
	_231_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_114_v2), _114_v2)
	var _232_v97 _dafny.Sequence
	_ = _232_v97
	_232_v97 = _dafny.SeqOf((_116_v4).Cardinality())
	var _233_v98 _dafny.Map
	_ = _233_v98
	_233_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_232_v97).Select((Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_232_v97).Cardinality()))).Uint32()).(_dafny.Int), _114_v2)
	var _234_v99 _dafny.MultiSet
	_ = _234_v99
	_234_v99 = _dafny.MultiSetOf(_114_v2)
	var _235_i8 _dafny.Int
	_ = _235_i8
	_235_i8 = _dafny.Zero
	{
		var _pat_let_tv1 = _114_v2
		_ = _pat_let_tv1
		var _pat_let_tv2 = _114_v2
		_ = _pat_let_tv2
		var _pat_let_tv3 = _114_v2
		_ = _pat_let_tv3
		var _pat_let_tv4 = _114_v2
		_ = _pat_let_tv4
		for func(_source4 D1) bool {
			if _source4.Is_DC5() {
				return _pat_let_tv1
			} else if _source4.Is_DC6() {
				var _241___mcc_h3 _dafny.Sequence = _source4.Get_().(D1_DC6).Cf5
				_ = _241___mcc_h3
				var _242___mcc_h4 _dafny.Int = _source4.Get_().(D1_DC6).Cf6
				_ = _242___mcc_h4
				var _243_cf6 _dafny.Int = _242___mcc_h4
				_ = _243_cf6
				var _244_cf5 _dafny.Sequence = _241___mcc_h3
				_ = _244_cf5
				return _pat_let_tv2
			} else {
				var _245___mcc_h5 _dafny.Set = _source4.Get_().(D1_DC4).Cf4
				_ = _245___mcc_h5
				var _246_cf4 _dafny.Set = _245___mcc_h5
				_ = _246_cf4
				return !(_pat_let_tv3) || (_pat_let_tv4)
			}
		}(Companion_Default___.Fm3(Companion_Default___.Fm37(false, (_231_v96).Cardinality(), Companion_Default___.Fm38(_233_v98, _114_v2, _114_v2, _114_v2, _118_globalState), _118_globalState), _dafny.IntOfInt64(377), _dafny.SeqOf(_234_v99), _dafny.Companion_Sequence_.Update(_232_v97, (Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_232_v97).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("daff")).Cardinality())), _118_globalState)) {
			{
				if (_235_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_235_i8 = (_235_i8).Plus(_dafny.One)
				var _236_v100 _dafny.Array
				_ = _236_v100
				var _237_v101 bool
				_ = _237_v101
				var _238_v102 _dafny.MultiSet
				_ = _238_v102
				var _239_v103 _dafny.Int
				_ = _239_v103
				var _out16 _dafny.Array
				_ = _out16
				var _out17 bool
				_ = _out17
				var _out18 _dafny.MultiSet
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				_out16, _out17, _out18, _out19 = Companion_Default___.M0(_113_v1, _118_globalState)
				_236_v100 = _out16
				_237_v101 = _out17
				_238_v102 = _out18
				_239_v103 = _out19
				(_118_globalState).F2 = (func() _dafny.Int {
					if (_239_v103).Cmp(_239_v103) < 0 {
						return _239_v103
					}
					return (_234_v99).Cardinality()
				})()
				var _240_v104 *C1
				_ = _240_v104
				var _nw36 *C1 = New_C1_()
				_ = _nw36
				_nw36.Ctor__(_dafny.IntOfInt64(623), (_112_v0).Times(_112_v0))
				_240_v104 = _nw36
				_185_v63 = _236_v100
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _247_v105 _dafny.Sequence
	_ = _247_v105
	_247_v105 = _dafny.SeqOf(_116_v4)
	var _hi2 _dafny.Int = _dafny.IntOfUint32((_184_v62).Cardinality())
	_ = _hi2
	for _248_i9 := _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_247_v105, _247_v105)).Cardinality()); _248_i9.Cmp(_hi2) < 0; _248_i9 = _248_i9.Plus(_dafny.One) {
		_114_v2 = _114_v2
		var _249_v107 _dafny.Set
		_ = _249_v107
		_249_v107 = _dafny.SetOf(Companion_Default___.Fm37(_114_v2, _dafny.IntOfUint32((_184_v62).Cardinality()), _114_v2, _118_globalState))
		_233_v98 = (func() _dafny.Map {
			if _114_v2 {
				return (func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate((_249_v107).Elements()); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _250_v106 _dafny.Int
						_250_v106 = interface{}(_compr_18).(_dafny.Int)
						if (_249_v107).Contains(_250_v106) {
							_coll18.Add(Companion_Default___.SafeDivisionInt(_250_v106, _112_v0), _114_v2)
						}
					}
					return _coll18.ToMap()
				}()).Merge(_233_v98)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v0, _114_v2)
		})()
		var _251_v108 _dafny.Map
		_ = _251_v108
		_251_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v60, _112_v0)
		_251_v108 = (_251_v108).Update(_182_v60, _112_v0)
		_114_v2 = _114_v2
	}
	var _252_v109 _dafny.Array
	_ = _252_v109
	var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
	_ = _nw37
	_252_v109 = _nw37
	var _253_v110 _dafny.Map
	_ = _253_v110
	_253_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_112_v0), Companion_Default___.Fm37(_114_v2, (_231_v96).Cardinality(), _114_v2, _118_globalState))
	var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_252_v109), 0))
	_ = _index16
	(_252_v109).ArraySet1((_253_v110).Update(_112_v0, _112_v0), (_index16).Int())
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(135), _dafny.ArrayLen((_252_v109), 0))
	_ = _index17
	(_252_v109).ArraySet1((func() _dafny.Map {
		if !(_114_v2) || (_114_v2) {
			return (_253_v110).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(536), _112_v0))
		}
		return _253_v110
	})(), (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
	_ = _index18
	(_185_v63).ArraySet1((_dafny.IntOfInt64(-644)).Plus(_112_v0), (_index18).Int())
	var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
	_ = _index19
	var _rhs34 bool = !(true)
	_ = _rhs34
	var _rhs35 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_112_v0, _112_v0))).Times((_112_v0).Plus(_dafny.IntOfUint32((_184_v62).Cardinality())))
	_ = _rhs35
	var _rhs36 _dafny.Int = (_dafny.Zero).Minus((_112_v0).Minus(_112_v0))
	_ = _rhs36
	var _rhs37 _dafny.Int = _112_v0
	_ = _rhs37
	var _lhs16 *GlobalState = _118_globalState
	_ = _lhs16
	var _lhs17 _dafny.Array = _185_v63
	_ = _lhs17
	var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
	_ = _lhs18
	var _lhs19 *GlobalState = _118_globalState
	_ = _lhs19
	_114_v2 = _rhs34
	_lhs16.F2 = _rhs35
	(_lhs17).ArraySet1(_rhs36, (_lhs18).Int())
	_lhs19.F2 = _rhs37
	var _254_v111 _dafny.Map
	_ = _254_v111
	_254_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v2, (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int))
	_254_v111 = (_254_v111).Update(!(_114_v2), _112_v0)
	var _255_v112 D3
	_ = _255_v112
	_255_v112 = Companion_D3_.Create_DC8_(_231_v96)
	var _256_v113 _dafny.Sequence
	_ = _256_v113
	_256_v113 = _dafny.SeqOf(_255_v112, _255_v112)
	var _257_v114 D16
	_ = _257_v114
	_257_v114 = Companion_D16_.Create_DC40_(_256_v113)
	_256_v113 = _dafny.Companion_Sequence_.Concatenate(_256_v113, _dafny.Companion_Sequence_.Concatenate(_256_v113, (_257_v114).Dtor_cf66()))
	var _258_v117 _dafny.Map
	_ = _258_v117
	_258_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _233_v98)
	var _259_v119 _dafny.Set
	_ = _259_v119
	_259_v119 = _dafny.SetOf(_112_v0)
	var _260_v121 _dafny.Map
	_ = _260_v121
	_260_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_233_v98, _233_v98)
	var _261_v122 _dafny.Array
	_ = _261_v122
	var _nwElement0_7 _dafny.Map = _233_v98
	_ = _nwElement0_7
	var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(23))
	_ = _nw38
	(_nw38).ArraySet1(_nwElement0_7, 0)
	(_nw38).ArraySet1(_233_v98, 1)
	(_nw38).ArraySet1(_233_v98, 2)
	(_nw38).ArraySet1(_233_v98, 3)
	(_nw38).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(284), false), 4)
	(_nw38).ArraySet1((_233_v98).Merge(_233_v98), 5)
	(_nw38).ArraySet1(_233_v98, 6)
	(_nw38).ArraySet1(_233_v98, 7)
	(_nw38).ArraySet1((func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter20 := _dafny.Iterate((_116_v4).Elements()); ; {
			_compr_19, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _262_v115 _dafny.Int
			_262_v115 = interface{}(_compr_19).(_dafny.Int)
			if (_116_v4).Contains(_262_v115) {
				_coll19.Add(Companion_Default___.SafeDivisionInt(_262_v115, _dafny.IntOfInt64(-416)), _114_v2)
			}
		}
		return _coll19.ToMap()
	}()).Merge(_233_v98), 8)
	(_nw38).ArraySet1((_233_v98).Merge(_233_v98), 9)
	(_nw38).ArraySet1((_233_v98).Merge(_233_v98), 10)
	(_nw38).ArraySet1(_233_v98, 11)
	(_nw38).ArraySet1(((_233_v98).Update((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _114_v2)).Merge(_233_v98), 12)
	(_nw38).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(197), _114_v2), 13)
	(_nw38).ArraySet1((func() _dafny.Map {
		if _114_v2 {
			return _233_v98
		}
		return func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter21 := _dafny.Iterate((_116_v4).Elements()); ; {
				_compr_20, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _263_v116 _dafny.Int
				_263_v116 = interface{}(_compr_20).(_dafny.Int)
				if (_116_v4).Contains(_263_v116) {
					_coll20.Add(Companion_Default___.SafeDivisionInt(_263_v116, (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)), _114_v2)
				}
			}
			return _coll20.ToMap()
		}()
	})(), 14)
	(_nw38).ArraySet1((func() _dafny.Map {
		if (_258_v117).Contains(_dafny.IntOfUint32((_184_v62).Cardinality())) {
			return (_258_v117).Get(_dafny.IntOfUint32((_184_v62).Cardinality())).(_dafny.Map)
		}
		return _233_v98
	})(), 15)
	(_nw38).ArraySet1(func() _dafny.Map {
		var _coll21 = _dafny.NewMapBuilder()
		_ = _coll21
		for _iter22 := _dafny.Iterate((_259_v119).Elements()); ; {
			_compr_21, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _264_v118 _dafny.Int
			_264_v118 = interface{}(_compr_21).(_dafny.Int)
			if (_259_v119).Contains(_264_v118) {
				_coll21.Add(Companion_Default___.SafeModuloInt(_264_v118, _112_v0), _114_v2)
			}
		}
		return _coll21.ToMap()
	}(), 16)
	(_nw38).ArraySet1(func() _dafny.Map {
		var _coll22 = _dafny.NewMapBuilder()
		_ = _coll22
		for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(831), _dafny.IntOfInt64(188))); ; {
			_compr_22, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _265_v120 _dafny.Int
			_265_v120 = interface{}(_compr_22).(_dafny.Int)
			if ((_dafny.IntOfInt64(831)).Cmp(_265_v120) <= 0) && ((_265_v120).Cmp(_dafny.IntOfInt64(188)) < 0) {
				_coll22.Add((_265_v120).Plus((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)), _114_v2)
			}
		}
		return _coll22.ToMap()
	}(), 17)
	(_nw38).ArraySet1(_233_v98, 18)
	(_nw38).ArraySet1(_233_v98, 19)
	(_nw38).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _114_v2)).Merge(_233_v98), 20)
	(_nw38).ArraySet1((func() _dafny.Map {
		if (_260_v121).Contains(_233_v98) {
			return (_260_v121).Get(_233_v98).(_dafny.Map)
		}
		return _233_v98
	})(), 21)
	(_nw38).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), !(_114_v2)), 22)
	_261_v122 = _nw38
	var _266_v123 *C2
	_ = _266_v123
	var _nw39 *C2 = New_C2_()
	_ = _nw39
	_nw39.Ctor__(_112_v0, Companion_Default___.Fm37(_114_v2, _dafny.IntOfInt64(799), Companion_Default___.Fm38(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_112_v0, _114_v2), _114_v2, _114_v2, _114_v2, _118_globalState), _118_globalState))
	_266_v123 = _nw39
	var _267_v124 _dafny.MultiSet
	_ = _267_v124
	_267_v124 = _dafny.MultiSetOf(_266_v123)
	var _rhs38 _dafny.Array = _261_v122
	_ = _rhs38
	var _rhs39 bool = (_267_v124).IsSubsetOf((func() _dafny.MultiSet {
		if _114_v2 {
			return _dafny.MultiSetOf(_266_v123)
		}
		return _267_v124
	})())
	_ = _rhs39
	_261_v122 = _rhs38
	_114_v2 = _rhs39
	var _268_v125 _dafny.Int
	_ = _268_v125
	var _269_v126 bool
	_ = _269_v126
	var _out20 _dafny.Int
	_ = _out20
	var _out21 bool
	_ = _out21
	_out20, _out21 = (_266_v123).M1(_118_globalState)
	_268_v125 = _out20
	_269_v126 = _out21
	var _hi3 _dafny.Int = _268_v125
	_ = _hi3
	for _270_i10 := _268_v125; _270_i10.Cmp(_hi3) < 0; _270_i10 = _270_i10.Plus(_dafny.One) {
		if _114_v2 {
			_231_v96 = (_231_v96).Update(!(!(_269_v126)), (_dafny.IntOfInt64(106)).Cmp((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)) > 0)
			_114_v2 = ((func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(816), _dafny.IntOfInt64(77))); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _271_v128 _dafny.Int
					_271_v128 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(816)).Cmp(_271_v128) <= 0) && ((_271_v128).Cmp(_dafny.IntOfInt64(77)) < 0) {
						_coll23.Add(Companion_Default___.SafeDivisionInt(_271_v128, _270_i10))
					}
				}
				return _coll23.ToSet()
			}()).Union(_259_v119)).IsSubsetOf((func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(276), _dafny.IntOfInt64(504))); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _272_v127 _dafny.Int
					_272_v127 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(276)).Cmp(_272_v127) <= 0) && ((_272_v127).Cmp(_dafny.IntOfInt64(504)) < 0) {
						_coll24.Add(Companion_Default___.SafeModuloInt(_272_v127, (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll24.ToSet()
			}()).Union(_dafny.SetOf((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), (_dafny.MultiSetOf(_112_v0, _268_v125)).Cardinality())))
			_114_v2 = ((_116_v4).Cardinality()).Cmp((_266_v123).Fm2((_266_v123).Fm2((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _183_v61, _118_globalState), _dafny.SetOf(true), _118_globalState)) > 0
			var _273_v129 _dafny.Array
			_ = _273_v129
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
			_ = _nw40
			_273_v129 = _nw40
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_273_v129), 0))
			_ = _index20
			(_273_v129).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_182_v60, _182_v60), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_273_v129), 0))
			_ = _index21
			(_273_v129).ArraySet1(_dafny.SeqOf(_269_v126, !(false) || (_114_v2)), (_index21).Int())
			var _274_v130 _dafny.Array
			_ = _274_v130
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_4
			var _nw41 _dafny.Array
			_ = _nw41
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw41 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_275_v2 bool) func(_dafny.Int) bool {
					return func(_276_i11 _dafny.Int) bool {
						return (!(false)) || (!(_275_v2))
					}
				})(_114_v2)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw41 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw41).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw41).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_274_v130 = _nw41
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_274_v130), 0))
			_ = _index22
			(_274_v130).ArraySet1(_269_v126, (_index22).Int())
			var _277_v131 D16
			_ = _277_v131
			_277_v131 = Companion_D16_.Create_DC42_((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_184_v62).Cardinality()), _113_v1)
			var _278_v132 _dafny.Map
			_ = _278_v132
			_278_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_273_v129).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(361), _dafny.ArrayLen((_273_v129), 0))).Int()).(_dafny.Sequence), _dafny.IntOfUint32((_184_v62).Cardinality()))
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_274_v130), 0))
			_ = _index23
			var _rhs40 _dafny.CodePoint = (_277_v131).Dtor_cf70()
			_ = _rhs40
			var _rhs41 bool = _114_v2
			_ = _rhs41
			var _rhs42 bool = !(((_278_v132).Cardinality()).Cmp(_268_v125) < 0)
			_ = _rhs42
			var _lhs20 _dafny.Array = _274_v130
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_274_v130), 0))
			_ = _lhs21
			_113_v1 = _rhs40
			_114_v2 = _rhs41
			(_lhs20).ArraySet1(_rhs42, (_lhs21).Int())
		} else {
			var _279_v133 _dafny.Map
			_ = _279_v133
			_279_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_232_v97, _dafny.Companion_Sequence_.Update(_232_v97, (Companion_Default___.SafeIndex((_231_v96).Cardinality(), _dafny.IntOfUint32((_232_v97).Cardinality()))).Uint32(), _268_v125)), _268_v125)
			_279_v133 = (_279_v133).Update(_232_v97, _dafny.IntOfUint32((_232_v97).Cardinality()))
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _index24
			(_185_v63).ArraySet1(_270_i10, (_index24).Int())
			(_266_v123).M2(_268_v125, _118_globalState)
			var _280_v134 D7
			_ = _280_v134
			_280_v134 = Companion_D7_.Create_DC23_(Companion_Default___.Fm37(_114_v2, _112_v0, _269_v126, _118_globalState), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_281_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_282_i12 _dafny.Int) _dafny.CodePoint {
					return _281_v1
				}
			})(_113_v1))), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nltodsjko")).Cardinality()), (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)))
			_280_v134 = _280_v134
			_185_v63 = _185_v63
		}
		if _114_v2 {
			var _283_v135 _dafny.Map
			_ = _283_v135
			_283_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xmmvfaq")), _114_v2)
			var _284_v136 _dafny.MultiSet
			_ = _284_v136
			_284_v136 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xsltdk"), _184_v62, _dafny.UnicodeSeqOfUtf8Bytes("etaeaj"))
			var _285_v137 _dafny.Map
			_ = _285_v137
			_285_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _184_v62)
			_283_v135 = (_283_v135).Update((_284_v136).Update((func() _dafny.Sequence {
				if (_285_v137).Contains(_114_v2) {
					return (_285_v137).Get(_114_v2).(_dafny.Sequence)
				}
				return _184_v62
			})(), Companion_Default___.Abs(_270_i10)), _dafny.Companion_Sequence_.IsPrefixOf(_184_v62, _184_v62))
			var _286_v138 *C0
			_ = _286_v138
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__(_270_i10)
			_286_v138 = _nw42
			var _287_v139 *C1
			_ = _287_v139
			var _nw43 *C1 = New_C1_()
			_ = _nw43
			_nw43.Ctor__(_dafny.IntOfUint32((_184_v62).Cardinality()), (_270_i10).Minus(_268_v125))
			_287_v139 = _nw43
			var _288_v140 _dafny.Array
			_ = _288_v140
			var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw44
			_288_v140 = _nw44
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v140), 0))
			_ = _index25
			(_288_v140).ArraySet1(_dafny.SeqOf(_270_i10), (_index25).Int())
			var _289_v141 _dafny.Map
			_ = _289_v141
			_289_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_287_v139, _112_v0)
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v140), 0))
			_ = _index26
			var _rhs43 *C1 = _287_v139
			_ = _rhs43
			var _rhs44 _dafny.Sequence = _232_v97
			_ = _rhs44
			var _rhs45 bool = (_114_v2) == (_114_v2)
			_ = _rhs45
			var _rhs46 _dafny.Int = (func() _dafny.Int {
				if _114_v2 {
					return (_232_v97).Select((Companion_Default___.SafeIndex((_289_v141).Cardinality(), _dafny.IntOfUint32((_232_v97).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return _112_v0
			})()
			_ = _rhs46
			var _lhs22 _dafny.Array = _288_v140
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v140), 0))
			_ = _lhs23
			var _lhs24 *GlobalState = _118_globalState
			_ = _lhs24
			_287_v139 = _rhs43
			(_lhs22).ArraySet1(_rhs44, (_lhs23).Int())
			_114_v2 = _rhs45
			_lhs24.F2 = _rhs46
			_114_v2 = _269_v126
			_113_v1 = _113_v1
		} else {
			(_266_v123).M4((_270_i10).Cmp((_dafny.MultiSetOf(_269_v126)).Cardinality()) >= 0, _118_globalState)
			var _290_v142 _dafny.Array
			_ = _290_v142
			var _nwElement0_8 _dafny.Sequence = _184_v62
			_ = _nwElement0_8
			var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(2))
			_ = _nw45
			(_nw45).ArraySet1(_nwElement0_8, 0)
			(_nw45).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_184_v62, _dafny.UnicodeSeqOfUtf8Bytes("dmsgbdkrd")), 1)
			_290_v142 = _nw45
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_290_v142), 0))
			_ = _index27
			(_290_v142).ArraySet1(Companion_Default___.Fm11(_dafny.IntOfInt64(74), (_259_v119).Cardinality(), _112_v0, _112_v0, _118_globalState), (_index27).Int())
			var _291_v143 _dafny.Sequence
			_ = _291_v143
			_291_v143 = _dafny.SeqOf(_184_v62)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(903), _dafny.ArrayLen((_290_v142), 0))
			_ = _index28
			(_290_v142).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_184_v62, _dafny.UnicodeSeqOfUtf8Bytes("laqun")), _dafny.Companion_Sequence_.Concatenate(_184_v62, (_291_v143).Select((Companion_Default___.SafeIndex(_270_i10, _dafny.IntOfUint32((_291_v143).Cardinality()))).Uint32()).(_dafny.Sequence))), (_index28).Int())
			var _292_v144 _dafny.Array
			_ = _292_v144
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(28))
			_ = _nw46
			_292_v144 = _nw46
			_292_v144 = _292_v144
			var _293_v145 T0
			_ = _293_v145
			var _nw47 *C1 = New_C1_()
			_ = _nw47
			_nw47.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _268_v125)).Merge(_254_v111)).Cardinality(), _dafny.IntOfInt64(296))
			_293_v145 = _nw47
			var _rhs47 bool = _269_v126
			_ = _rhs47
			var _rhs48 _dafny.Sequence = Companion_Default___.Fm28(_268_v125, _118_globalState)
			_ = _rhs48
			var _rhs49 bool = _269_v126
			_ = _rhs49
			var _rhs50 T0 = _293_v145
			_ = _rhs50
			_114_v2 = _rhs47
			_232_v97 = _rhs48
			_269_v126 = _rhs49
			_293_v145 = _rhs50
			var _294_v146 _dafny.Map
			_ = _294_v146
			_294_v146 = (_254_v111).Update(_114_v2, _268_v125)
			var _295_v147 _dafny.Array
			_ = _295_v147
			var _nwElement0_9 _dafny.Map = _294_v146
			_ = _nwElement0_9
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(26))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_9, 0)
			(_nw48).ArraySet1(_294_v146, 1)
			(_nw48).ArraySet1(_294_v146, 2)
			(_nw48).ArraySet1(_294_v146, 3)
			(_nw48).ArraySet1(_294_v146, 4)
			(_nw48).ArraySet1(_254_v111, 5)
			(_nw48).ArraySet1(_254_v111, 6)
			(_nw48).ArraySet1(_294_v146, 7)
			(_nw48).ArraySet1(_254_v111, 8)
			(_nw48).ArraySet1(_254_v111, 9)
			(_nw48).ArraySet1(_294_v146, 10)
			(_nw48).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _dafny.IntOfInt64(-772)), 11)
			(_nw48).ArraySet1(_294_v146, 12)
			(_nw48).ArraySet1(_294_v146, 13)
			(_nw48).ArraySet1(_294_v146, 14)
			(_nw48).ArraySet1((func() _dafny.Map {
				if _269_v126 {
					return _294_v146
				}
				return _294_v146
			})(), 15)
			(_nw48).ArraySet1(_294_v146, 16)
			(_nw48).ArraySet1(_254_v111, 17)
			(_nw48).ArraySet1(_294_v146, 18)
			(_nw48).ArraySet1(_294_v146, 19)
			(_nw48).ArraySet1(_294_v146, 20)
			(_nw48).ArraySet1(_294_v146, 21)
			(_nw48).ArraySet1(_294_v146, 22)
			(_nw48).ArraySet1(_294_v146, 23)
			(_nw48).ArraySet1(_294_v146, 24)
			(_nw48).ArraySet1(_294_v146, 25)
			_295_v147 = _nw48
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_295_v147), 0))
			_ = _index29
			(_295_v147).ArraySet1(_254_v111, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_295_v147), 0))
			_ = _index30
			(_295_v147).ArraySet1(_294_v146, (_index30).Int())
		}
		var _296_v148 _dafny.Map
		_ = _296_v148
		_296_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _dafny.SetOf(_114_v2, _269_v126, _269_v126))
		var _297_v149 *C1
		_ = _297_v149
		var _nw49 *C1 = New_C1_()
		_ = _nw49
		_nw49.Ctor__(((func() _dafny.Set {
			if (_296_v148).Contains(_114_v2) {
				return (_296_v148).Get(_114_v2).(_dafny.Set)
			}
			return _dafny.SetOf(_114_v2)
		})()).Cardinality(), _112_v0)
		_297_v149 = _nw49
		var _298_v150 _dafny.Map
		_ = _298_v150
		_298_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(640), _114_v2)).Cardinality())), _297_v149)
		var _299_v151 D11
		_ = _299_v151
		_299_v151 = Companion_D11_.Create_DC31_((func() *C1 {
			if (_298_v150).Contains((_dafny.Zero).Minus(_112_v0)) {
				return (_298_v150).Get((_dafny.Zero).Minus(_112_v0)).(*C1)
			}
			return _297_v149
		})())
		var _pat_let_tv5 = _297_v149
		_ = _pat_let_tv5
		var _source5 D11 = func(_pat_let5_0 D11) D11 {
			return func(_300_dt__update__tmp_h2 D11) D11 {
				return func(_pat_let6_0 *C1) D11 {
					return func(_301_dt__update_hcf51_h0 *C1) D11 {
						return Companion_D11_.Create_DC31_(_301_dt__update_hcf51_h0)
					}(_pat_let6_0)
				}(_pat_let_tv5)
			}(_pat_let5_0)
		}(_299_v151)
		_ = _source5
		if _source5.Is_DC32() {
			var _302___mcc_h6 bool = _source5.Get_().(D11_DC32).Cf52
			_ = _302___mcc_h6
			var _303___mcc_h7 bool = _source5.Get_().(D11_DC32).Cf53
			_ = _303___mcc_h7
			var _304___mcc_h8 bool = _source5.Get_().(D11_DC32).Cf54
			_ = _304___mcc_h8
			var _305_cf54 bool = _304___mcc_h8
			_ = _305_cf54
			var _306_cf53 bool = _303___mcc_h7
			_ = _306_cf53
			var _307_cf52 bool = _302___mcc_h6
			_ = _307_cf52
			var _308_v152 _dafny.Array
			_ = _308_v152
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw50
			_308_v152 = _nw50
			var _309_v153 D1
			_ = _309_v153
			_309_v153 = Companion_D1_.Create_DC5_()
			var _rhs51 bool = (_266_v123).Fm9(_309_v153, _114_v2, _118_globalState)
			_ = _rhs51
			var _rhs52 bool = _114_v2
			_ = _rhs52
			var _rhs53 _dafny.Int = _112_v0
			_ = _rhs53
			var _rhs54 _dafny.Set = ((_259_v119).Union(_259_v119)).Difference(_259_v119)
			_ = _rhs54
			var _rhs55 _dafny.Array = _308_v152
			_ = _rhs55
			var _lhs25 *GlobalState = _118_globalState
			_ = _lhs25
			_305_cf54 = _rhs51
			_269_v126 = _rhs52
			_lhs25.F2 = _rhs53
			_259_v119 = _rhs54
			_308_v152 = _rhs55
			_306_cf53 = (_268_v125).Cmp(_270_i10) <= 0
			var _310_v154 *C2
			_ = _310_v154
			var _nw51 *C2 = New_C2_()
			_ = _nw51
			_nw51.Ctor__(_dafny.IntOfInt64(320), (_dafny.Zero).Minus(_112_v0))
			_310_v154 = _nw51
			var _311_v155 _dafny.Array
			_ = _311_v155
			var _nw52 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw52
			_311_v155 = _nw52
			var _312_v156 D7
			_ = _312_v156
			_312_v156 = Companion_D7_.Create_DC22_(_311_v155, _184_v62)
			_311_v155 = (_312_v156).Dtor_cf38()
		} else {
			var _313___mcc_h9 *C1 = _source5.Get_().(D11_DC31).Cf51
			_ = _313___mcc_h9
			var _314_cf51 *C1 = _313___mcc_h9
			_ = _314_cf51
			_114_v2 = (_270_i10).Cmp(_268_v125) != 0
			_114_v2 = (_234_v99).Equals(Companion_Default___.Fm5(_232_v97, !(_269_v126), (_314_cf51).Fm13(_270_i10, _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sbeh")), _118_globalState), _118_globalState))
			var _315_v157 T1
			_ = _315_v157
			var _nw53 *C3 = New_C3_()
			_ = _nw53
			_nw53.Ctor__(_114_v2, (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _270_i10)
			_315_v157 = _nw53
			var _316_v158 _dafny.Map
			_ = _316_v158
			_316_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v62, _315_v157)
			var _317_v159 T0
			_ = _317_v159
			var _nw54 *C2 = New_C2_()
			_ = _nw54
			_nw54.Ctor__(_268_v125, (func() _dafny.Int {
				if (_254_v111).Contains(_114_v2) {
					return (_254_v111).Get(_114_v2).(_dafny.Int)
				}
				return ((_316_v158).Update(_184_v62, _315_v157)).Cardinality()
			})())
			_317_v159 = _nw54
			var _318_v160 T0
			_ = _318_v160
			_318_v160 = _317_v159
			var _rhs56 bool = false
			_ = _rhs56
			var _rhs57 bool = _269_v126
			_ = _rhs57
			var _rhs58 T0 = _318_v160
			_ = _rhs58
			_269_v126 = _rhs56
			_114_v2 = _rhs57
			_318_v160 = _rhs58
			var _319_v161 _dafny.MultiSet
			_ = _319_v161
			_319_v161 = _dafny.MultiSetOf(_184_v62, _184_v62)
			var _320_v162 *C0
			_ = _320_v162
			var _nw55 *C0 = New_C0_()
			_ = _nw55
			_nw55.Ctor__((_297_v149).Fm13((_314_cf51).Fm13(_315_v157.F6(), _319_v161, _118_globalState), _319_v161, _118_globalState))
			_320_v162 = _nw55
			var _321_v163 _dafny.Sequence
			_ = _321_v163
			_321_v163 = _dafny.SeqOf(_320_v162, _320_v162, _320_v162)
			var _322_v164 _dafny.Map
			_ = _322_v164
			_322_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_321_v163, _321_v163), _112_v0)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _index31
			var _rhs59 _dafny.Int = (_dafny.Zero).Minus(((_320_v162).F8()).Times((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)))
			_ = _rhs59
			var _rhs60 _dafny.Map = (_322_v164).Update(_dafny.SeqOf(_320_v162), (_dafny.Zero).Minus((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)))
			_ = _rhs60
			var _lhs26 _dafny.Array = _185_v63
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _lhs27
			(_lhs26).ArraySet1(_rhs59, (_lhs27).Int())
			_322_v164 = _rhs60
		}
		var _323_v165 T1
		_ = _323_v165
		var _nw56 *C3 = New_C3_()
		_ = _nw56
		_nw56.Ctor__((_183_v61).IsDisjointFrom(_dafny.SetOf(_114_v2)), (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _112_v0)
		_323_v165 = _nw56
		_323_v165 = _323_v165
	}
	_112_v0 = Companion_Default___.Fm37(Companion_Default___.Fm38(_233_v98, _114_v2, _114_v2, _269_v126, _118_globalState), (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _114_v2, _118_globalState)
	if (_259_v119).Equals(_259_v119) {
		var _324_v166 _dafny.Array
		_ = _324_v166
		var _nwElement0_10 _dafny.Array = _185_v63
		_ = _nwElement0_10
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(13))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_10, 0)
		(_nw57).ArraySet1(_185_v63, 1)
		(_nw57).ArraySet1(_185_v63, 2)
		(_nw57).ArraySet1(_185_v63, 3)
		(_nw57).ArraySet1(_185_v63, 4)
		(_nw57).ArraySet1(_185_v63, 5)
		(_nw57).ArraySet1(_185_v63, 6)
		(_nw57).ArraySet1(_185_v63, 7)
		(_nw57).ArraySet1(_185_v63, 8)
		(_nw57).ArraySet1(_185_v63, 9)
		(_nw57).ArraySet1(_185_v63, 10)
		(_nw57).ArraySet1(_185_v63, 11)
		(_nw57).ArraySet1(_185_v63, 12)
		_324_v166 = _nw57
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_324_v166), 0))
		_ = _index32
		(_324_v166).ArraySet1(_185_v63, (_index32).Int())
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_324_v166), 0))
		_ = _index33
		(_324_v166).ArraySet1(_185_v63, (_index33).Int())
		var _325_v167 _dafny.Map
		_ = _325_v167
		_325_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_183_v61).Cardinality(), _185_v63)
		var _326_v168 _dafny.Map
		_ = _326_v168
		_326_v168 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_325_v167).Merge(_325_v167), _269_v126)
		if (func() bool {
			if (_326_v168).Contains(_325_v167) {
				return (_326_v168).Get(_325_v167).(bool)
			}
			return _269_v126
		})() {
			var _327_v169 D4
			_ = _327_v169
			_327_v169 = Companion_D4_.Create_DC14_((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), _114_v2, _dafny.IntOfUint32((_184_v62).Cardinality()), !(!(_269_v126)), _112_v0)
			var _328_v170 _dafny.Map
			_ = _328_v170
			_328_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_327_v169, _dafny.IntOfInt64(8))
			var _329_v171 _dafny.Map
			_ = _329_v171
			_329_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v2, _328_v170)
			var _330_v172 T0
			_ = _330_v172
			var _nw58 *C1 = New_C1_()
			_ = _nw58
			_nw58.Ctor__(_268_v125, (_329_v171).Cardinality())
			_330_v172 = _nw58
			var _331_v173 _dafny.Map
			_ = _331_v173
			_331_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v125, _231_v96)
			var _332_v174 _dafny.Array
			_ = _332_v174
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_5
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_333_v2 bool) func(_dafny.Int) bool {
					return func(_334_i13 _dafny.Int) bool {
						return _333_v2
					}
				})(_114_v2)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw59 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw59).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw59).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_332_v174 = _nw59
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _index34
			(_332_v174).ArraySet1((_112_v0).Cmp(_268_v125) == 0, (_index34).Int())
			var _335_v175 T1
			_ = _335_v175
			var _nw60 *C1 = New_C1_()
			_ = _nw60
			_nw60.Ctor__(_112_v0, _112_v0)
			_335_v175 = _nw60
			var _336_v176 D9
			_ = _336_v176
			_336_v176 = Companion_D9_.Create_DC28_(_269_v126, _dafny.IntOfUint32((_dafny.SeqOf(_335_v175, _335_v175, _335_v175, _335_v175)).Cardinality()), _112_v0)
			var _337_v177 _dafny.Sequence
			_ = _337_v177
			_337_v177 = _dafny.SeqOf(_336_v176)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _index35
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _index36
			var _rhs61 T0 = _330_v172
			_ = _rhs61
			var _rhs62 _dafny.Map = _331_v173
			_ = _rhs62
			var _rhs63 bool = !(_114_v2)
			_ = _rhs63
			var _rhs64 _dafny.Int = (func() _dafny.Int {
				if _114_v2 {
					return (_dafny.IntOfUint32((_337_v177).Cardinality())).Minus(_dafny.IntOfUint32((_182_v60).Cardinality()))
				}
				return _335_v175.F6()
			})()
			_ = _rhs64
			var _lhs28 _dafny.Array = _332_v174
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _lhs29
			var _lhs30 _dafny.Array = _185_v63
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _lhs31
			_330_v172 = _rhs61
			_331_v173 = _rhs62
			(_lhs28).ArraySet1(_rhs63, (_lhs29).Int())
			(_lhs30).ArraySet1(_rhs64, (_lhs31).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _index37
			(_332_v174).ArraySet1(_269_v126, (_index37).Int())
			var _338_v178 _dafny.Int
			_ = _338_v178
			var _339_v179 bool
			_ = _339_v179
			var _out22 _dafny.Int
			_ = _out22
			var _out23 bool
			_ = _out23
			_out22, _out23 = (_266_v123).M1(_118_globalState)
			_338_v178 = _out22
			_339_v179 = _out23
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _index38
			var _rhs65 _dafny.MultiSet = ((_267_v124).Union(_267_v124)).Update(_266_v123, Companion_Default___.Abs(_dafny.IntOfInt64(328)))
			_ = _rhs65
			var _rhs66 bool = (_332_v174).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))).Int()).(bool)
			_ = _rhs66
			var _rhs67 *C2 = _266_v123
			_ = _rhs67
			var _rhs68 bool = _114_v2
			_ = _rhs68
			var _rhs69 bool = (_182_v60).Select((Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_335_v175).F5())).Cardinality()), _113_v1)).Update((_dafny.Zero).Minus(_335_v175.F6()), _113_v1)).Cardinality(), _dafny.IntOfUint32((_182_v60).Cardinality()))).Uint32()).(bool)
			_ = _rhs69
			var _lhs32 _dafny.Array = _332_v174
			_ = _lhs32
			var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_332_v174), 0))
			_ = _lhs33
			_267_v124 = _rhs65
			(_lhs32).ArraySet1(_rhs66, (_lhs33).Int())
			_266_v123 = _rhs67
			_114_v2 = _rhs68
			_339_v179 = _rhs69
			_114_v2 = (_dafny.IntOfUint32((_184_v62).Cardinality())).Cmp(_dafny.IntOfInt64(998)) != 0
		} else {
			(_266_v123).M4(!(true), _118_globalState)
			var _340_v180 _dafny.Array
			_ = _340_v180
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw61
			_340_v180 = _nw61
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_340_v180), 0))
			_ = _index39
			(_340_v180).ArraySet1(!(!(_114_v2) || (true)), (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_340_v180), 0))
			_ = _index40
			(_340_v180).ArraySet1(_114_v2, (_index40).Int())
			(_118_globalState).F2 = _112_v0
			var _341_v181 _dafny.Int
			_ = _341_v181
			var _342_v182 bool
			_ = _342_v182
			var _out24 _dafny.Int
			_ = _out24
			var _out25 bool
			_ = _out25
			_out24, _out25 = (_266_v123).M1(_118_globalState)
			_341_v181 = _out24
			_342_v182 = _out25
			_116_v4 = ((_116_v4).Update(_dafny.IntOfUint32((_184_v62).Cardinality()), Companion_Default___.Abs(_341_v181))).Difference(_116_v4)
		}
		var _343_v183 _dafny.Array
		_ = _343_v183
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw62
		_343_v183 = _nw62
		var _344_v184 _dafny.Array
		_ = _344_v184
		var _nw63 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw63
		_344_v184 = _nw63
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_343_v183), 0))
		_ = _index41
		(_343_v183).ArraySet1(_344_v184, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_343_v183), 0))
		_ = _index42
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
		_ = _index43
		var _rhs70 _dafny.Int = _268_v125
		_ = _rhs70
		var _rhs71 _dafny.Array = _344_v184
		_ = _rhs71
		var _rhs72 _dafny.Int = _268_v125
		_ = _rhs72
		var _rhs73 _dafny.Sequence = _184_v62
		_ = _rhs73
		var _rhs74 _dafny.Sequence = (func() _dafny.Sequence {
			if _114_v2 {
				return _184_v62
			}
			return _184_v62
		})()
		_ = _rhs74
		var _lhs34 *GlobalState = _118_globalState
		_ = _lhs34
		var _lhs35 _dafny.Array = _343_v183
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_343_v183), 0))
		_ = _lhs36
		var _lhs37 _dafny.Array = _185_v63
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
		_ = _lhs38
		_lhs34.F2 = _rhs70
		(_lhs35).ArraySet1(_rhs71, (_lhs36).Int())
		(_lhs37).ArraySet1(_rhs72, (_lhs38).Int())
		_184_v62 = _rhs73
		_184_v62 = _rhs74
		var _345_v185 _dafny.Array
		_ = _345_v185
		var _nw64 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(2))
		_ = _nw64
		_345_v185 = _nw64
		var _346_v186 *C1
		_ = _346_v186
		var _nw65 *C1 = New_C1_()
		_ = _nw65
		_nw65.Ctor__((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int), (_253_v110).Cardinality())
		_346_v186 = _nw65
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_345_v185), 0))
		_ = _index44
		(_345_v185).ArraySet1(_346_v186, (_index44).Int())
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_345_v185), 0))
		_ = _index45
		(_345_v185).ArraySet1(_346_v186, (_index45).Int())
		var _rhs75 _dafny.Int = (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)
		_ = _rhs75
		var _rhs76 _dafny.CodePoint = _113_v1
		_ = _rhs76
		var _lhs39 *GlobalState = _118_globalState
		_ = _lhs39
		_lhs39.F2 = _rhs75
		_113_v1 = _rhs76
	} else {
		_184_v62 = _dafny.UnicodeSeqOfUtf8Bytes("qcqcxq")
		(_266_v123).M2(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_184_v62).Cardinality()), (_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int)), _118_globalState)
		_269_v126 = !(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("miaw"), _dafny.CodePoint('a'))) || ((func() bool {
			if (_233_v98).Contains(_dafny.IntOfInt64(461)) {
				return (_233_v98).Get(_dafny.IntOfInt64(461)).(bool)
			}
			return _114_v2
		})())
		if (_269_v126) || (!(_114_v2)) {
			var _347_v187 *C0
			_ = _347_v187
			var _nw66 *C0 = New_C0_()
			_ = _nw66
			_nw66.Ctor__((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int))
			_347_v187 = _nw66
			var _348_v188 _dafny.Sequence
			_ = _348_v188
			_348_v188 = _dafny.SeqOf(_254_v111, _254_v111)
			var _349_v189 _dafny.Set
			_ = _349_v189
			_349_v189 = _dafny.SetOf(_113_v1)
			var _350_v190 _dafny.Array
			_ = _350_v190
			var _nwElement0_11 _dafny.Map = _254_v111
			_ = _nwElement0_11
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(11))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_11, 0)
			(_nw67).ArraySet1((_254_v111).Merge(_254_v111), 1)
			(_nw67).ArraySet1((_254_v111).Merge(_254_v111), 2)
			(_nw67).ArraySet1(_254_v111, 3)
			(_nw67).ArraySet1((_348_v188).Select((Companion_Default___.SafeIndex((_347_v187).F8(), _dafny.IntOfUint32((_348_v188).Cardinality()))).Uint32()).(_dafny.Map), 4)
			(_nw67).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _268_v125)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, (_349_v189).Cardinality())), 5)
			(_nw67).ArraySet1(Companion_Default___.Fm39(_118_globalState), 6)
			(_nw67).ArraySet1(_254_v111, 7)
			(_nw67).ArraySet1(_254_v111, 8)
			(_nw67).ArraySet1(_254_v111, 9)
			(_nw67).ArraySet1(_254_v111, 10)
			_350_v190 = _nw67
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_350_v190), 0))
			_ = _index46
			(_350_v190).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_269_v126), _268_v125)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v126, _dafny.IntOfInt64(193))), (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_350_v190), 0))
			_ = _index47
			var _rhs77 _dafny.Int = Companion_Default___.SafeModuloInt(_268_v125, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("awmjrtbl")).Cardinality()))
			_ = _rhs77
			var _rhs78 _dafny.Sequence = _232_v97
			_ = _rhs78
			var _rhs79 _dafny.Map = (_348_v188).Select((Companion_Default___.SafeIndex(_268_v125, _dafny.IntOfUint32((_348_v188).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs79
			var _rhs80 _dafny.Int = (_266_v123).Fm2((_dafny.Zero).Minus((_dafny.Zero).Minus((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int))), _183_v61, _118_globalState)
			_ = _rhs80
			var _rhs81 bool = _269_v126
			_ = _rhs81
			var _lhs40 *GlobalState = _118_globalState
			_ = _lhs40
			var _lhs41 _dafny.Array = _350_v190
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_350_v190), 0))
			_ = _lhs42
			var _lhs43 *GlobalState = _118_globalState
			_ = _lhs43
			_lhs40.F2 = _rhs77
			_232_v97 = _rhs78
			(_lhs41).ArraySet1(_rhs79, (_lhs42).Int())
			_lhs43.F2 = _rhs80
			_114_v2 = _rhs81
			var _rhs82 _dafny.Int = (_112_v0).Plus((_185_v63).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))).Int()).(_dafny.Int))
			_ = _rhs82
			var _rhs83 _dafny.CodePoint = _113_v1
			_ = _rhs83
			var _rhs84 _dafny.Map = func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(243), _dafny.IntOfInt64(405))); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _351_v191 _dafny.Int
					_351_v191 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(243)).Cmp(_351_v191) <= 0) && ((_351_v191).Cmp(_dafny.IntOfInt64(405)) < 0) {
						_coll25.Add(Companion_Default___.SafeModuloInt(_351_v191, ((_116_v4).Update((_183_v61).Cardinality(), Companion_Default___.Abs((_231_v96).Cardinality()))).Cardinality()), (_347_v187).F8())
					}
				}
				return _coll25.ToMap()
			}()
			_ = _rhs84
			var _lhs44 *GlobalState = _118_globalState
			_ = _lhs44
			_lhs44.F2 = _rhs82
			_113_v1 = _rhs83
			_253_v110 = _rhs84
			_184_v62 = _dafny.Companion_Sequence_.Update(_184_v62, (Companion_Default___.SafeIndex(_268_v125, _dafny.IntOfUint32((_184_v62).Cardinality()))).Uint32(), _113_v1)
			var _352_v192 _dafny.Set
			_ = _352_v192
			_352_v192 = (_183_v61).Difference(_183_v61)
			_352_v192 = _183_v61
		} else {
			var _353_v193 _dafny.Int
			_ = _353_v193
			var _354_v194 bool
			_ = _354_v194
			var _out26 _dafny.Int
			_ = _out26
			var _out27 bool
			_ = _out27
			_out26, _out27 = (_266_v123).M1(_118_globalState)
			_353_v193 = _out26
			_354_v194 = _out27
			var _355_v195 *C2
			_ = _355_v195
			var _nw68 *C2 = New_C2_()
			_ = _nw68
			_nw68.Ctor__((_232_v97).Select((Companion_Default___.SafeIndex(_112_v0, _dafny.IntOfUint32((_232_v97).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(136))
			_355_v195 = _nw68
			var _356_v196 _dafny.Sequence
			_ = _356_v196
			_356_v196 = _dafny.SeqOf(_184_v62)
			var _357_v197 _dafny.Sequence
			_ = _357_v197
			_357_v197 = _dafny.SeqOf(_259_v119)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _index48
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _index49
			var _rhs85 bool = (func() bool {
				if (_233_v98).Contains(_dafny.IntOfUint32(((_356_v196).Select((Companion_Default___.SafeIndex(((_357_v197).Select((Companion_Default___.SafeIndex(_353_v193, _dafny.IntOfUint32((_357_v197).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_356_v196).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())) {
					return (_233_v98).Get(_dafny.IntOfUint32(((_356_v196).Select((Companion_Default___.SafeIndex(((_357_v197).Select((Companion_Default___.SafeIndex(_353_v193, _dafny.IntOfUint32((_357_v197).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_356_v196).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).(bool)
				}
				return _354_v194
			})()
			_ = _rhs85
			var _rhs86 _dafny.Int = (_112_v0).Plus((_353_v193).Times(_dafny.IntOfInt64(955)))
			_ = _rhs86
			var _rhs87 _dafny.Int = _268_v125
			_ = _rhs87
			var _lhs45 _dafny.Array = _185_v63
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _lhs46
			var _lhs47 _dafny.Array = _185_v63
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_185_v63), 0))
			_ = _lhs48
			_354_v194 = _rhs85
			(_lhs45).ArraySet1(_rhs86, (_lhs46).Int())
			(_lhs47).ArraySet1(_rhs87, (_lhs48).Int())
			_268_v125 = _268_v125
			(_118_globalState).F2 = Companion_Default___.SafeModuloInt(_112_v0, _268_v125)
		}
		_114_v2 = !(!(_114_v2)) || (_269_v126)
	}
	_dafny.Print(_112_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_113_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_115_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v4).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_117_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.MultiSetOf(_dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_118_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_118_globalState).F4()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-1))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_176_i3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_182_v60, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v61).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_184_v62.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_185_v63).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v64).Dtor_cf5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_189_v64).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_231_v96).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_232_v97, _dafny.SeqOf(_dafny.IntOfInt64(3))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v98).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v99).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_235_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_247_v105, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723), _dafny.IntOfInt64(-723)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_252_v109).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(690)).UpdateUnsafe(_dafny.IntOfInt64(536), _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_253_v110).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.IntOfInt64(690))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v111).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.Zero).UpdateUnsafe(true, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_255_v112).Dtor_cf8()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_256_v113, _dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_257_v114).Dtor_cf66(), _dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)), Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v117).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v119).Equals(_dafny.SetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v121).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(284), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), false).UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true).UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(197), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-723), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false).UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), true).UpdateUnsafe(_dafny.IntOfInt64(3), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v122).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_267_v124).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_268_v125)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_269_v126)
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
	Cf2 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_() D0 {
	return D0{D0_DC2{}}
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
	return Companion_D0_.Create_DC1_(false, false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC3).Cf3
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2
		}
	case D0_DC2:
		{
			_, ok := other.Get_().(D0_DC2)
			return ok
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf3 == data2.Cf3
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

type D1_DC5 struct {
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_() D1 {
	return D1{D1_DC5{}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf5 _dafny.Sequence
	Cf6 _dafny.Int
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf5 _dafny.Sequence, Cf6 _dafny.Int) D1 {
	return D1{D1_DC6{Cf5, Cf6}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC4 struct {
	Cf4 _dafny.Set
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Set) D1 {
	return D1{D1_DC4{Cf4}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_()
}

func (_this D1) Dtor_cf5() _dafny.Sequence {
	return _this.Get_().(D1_DC6).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf6
}

func (_this D1) Dtor_cf4() _dafny.Set {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + data.Cf5.VerbatimString(true) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ")"
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
			_, ok := other.Get_().(D1_DC5)
			return ok
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf5.Equals(data2.Cf5) && data1.Cf6.Cmp(data2.Cf6) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
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
	Cf7 _dafny.Map
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf7 _dafny.Map) D2 {
	return D2{D2_DC7{Cf7}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D2) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf9  _dafny.MultiSet
	Cf10 D0
	Cf11 _dafny.Int
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf9 _dafny.MultiSet, Cf10 D0, Cf11 _dafny.Int) D3 {
	return D3{D3_DC9{Cf9, Cf10, Cf11}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf8 _dafny.Map
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf8 _dafny.Map) D3 {
	return D3{D3_DC8{Cf8}}
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
	return Companion_D3_.Create_DC9_(_dafny.EmptyMultiSet, Companion_D0_.Default(), _dafny.Zero)
}

func (_this D3) Dtor_cf9() _dafny.MultiSet {
	return _this.Get_().(D3_DC9).Cf9
}

func (_this D3) Dtor_cf10() D0 {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D3_DC9).Cf11
}

func (_this D3) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D3_DC8).Cf8
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
			return "D3.DC9" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf9.Equals(data2.Cf9) && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Cmp(data2.Cf11) == 0
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf14 bool
	Cf15 _dafny.Int
	Cf16 bool
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf14 bool, Cf15 _dafny.Int, Cf16 bool) D4 {
	return D4{D4_DC12{Cf14, Cf15, Cf16}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC13 struct {
	Cf17 bool
	Cf18 _dafny.Int
	Cf19 bool
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf17 bool, Cf18 _dafny.Int, Cf19 bool) D4 {
	return D4{D4_DC13{Cf17, Cf18, Cf19}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf20 _dafny.Int
	Cf21 bool
	Cf22 _dafny.Int
	Cf23 bool
	Cf24 _dafny.Int
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf20 _dafny.Int, Cf21 bool, Cf22 _dafny.Int, Cf23 bool, Cf24 _dafny.Int) D4 {
	return D4{D4_DC14{Cf20, Cf21, Cf22, Cf23, Cf24}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC11 struct {
	Cf13 _dafny.Sequence
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf13 _dafny.Sequence) D4 {
	return D4{D4_DC11{Cf13}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
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
	return Companion_D4_.Create_DC12_(false, _dafny.Zero, false)
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC12).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf15
}

func (_this D4) Dtor_cf16() bool {
	return _this.Get_().(D4_DC12).Cf16
}

func (_this D4) Dtor_cf17() bool {
	return _this.Get_().(D4_DC13).Cf17
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf18
}

func (_this D4) Dtor_cf19() bool {
	return _this.Get_().(D4_DC13).Cf19
}

func (_this D4) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf20
}

func (_this D4) Dtor_cf21() bool {
	return _this.Get_().(D4_DC14).Cf21
}

func (_this D4) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf22
}

func (_this D4) Dtor_cf23() bool {
	return _this.Get_().(D4_DC14).Cf23
}

func (_this D4) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D4_DC14).Cf24
}

func (_this D4) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D4_DC11).Cf13
}

func (_this D4) Dtor_cf25() D4 {
	return _this.Get_().(D4_DC15).Cf25
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf13) + ")"
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
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf14 == data2.Cf14 && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf17 == data2.Cf17 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf13.Equals(data2.Cf13)
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
	Cf28 _dafny.Array
	Cf29 _dafny.Map
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf27 _dafny.Int, Cf28 _dafny.Array, Cf29 _dafny.Map) D5 {
	return D5{D5_DC17{Cf27, Cf28, Cf29}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
	Cf30 bool
	Cf31 _dafny.Int
	Cf32 bool
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf30 bool, Cf31 _dafny.Int, Cf32 bool) D5 {
	return D5{D5_DC18{Cf30, Cf31, Cf32}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC19 struct {
	Cf33 _dafny.Int
	Cf34 _dafny.CodePoint
	Cf35 _dafny.Int
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf33 _dafny.Int, Cf34 _dafny.CodePoint, Cf35 _dafny.Int) D5 {
	return D5{D5_DC19{Cf33, Cf34, Cf35}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

type D5_DC16 struct {
	Cf26 _dafny.Array
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf26 _dafny.Array) D5 {
	return D5{D5_DC16{Cf26}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptyMap)
}

func (_this D5) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D5_DC17).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D5_DC17).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Map {
	return _this.Get_().(D5_DC17).Cf29
}

func (_this D5) Dtor_cf30() bool {
	return _this.Get_().(D5_DC18).Cf30
}

func (_this D5) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D5_DC18).Cf31
}

func (_this D5) Dtor_cf32() bool {
	return _this.Get_().(D5_DC18).Cf32
}

func (_this D5) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D5_DC19).Cf33
}

func (_this D5) Dtor_cf34() _dafny.CodePoint {
	return _this.Get_().(D5_DC19).Cf34
}

func (_this D5) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D5_DC19).Cf35
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC16).Cf26
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28 == data2.Cf28 && data1.Cf29.Equals(data2.Cf29)
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
			return ok && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34 == data2.Cf34 && data1.Cf35.Cmp(data2.Cf35) == 0
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf26 == data2.Cf26
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
	Cf36 T0
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf36 T0) D6 {
	return D6{D6_DC20{Cf36}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() T0 {
	return (T0)(nil)
}

func (_this D6) Dtor_cf36() T0 {
	return _this.Get_().(D6_DC20).Cf36
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && _dafny.AreEqual(data1.Cf36, data2.Cf36)
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

type D7_DC22 struct {
	Cf38 _dafny.Array
	Cf39 _dafny.Sequence
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf38 _dafny.Array, Cf39 _dafny.Sequence) D7 {
	return D7{D7_DC22{Cf38, Cf39}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC23 struct {
	Cf40 _dafny.Int
	Cf41 _dafny.Sequence
	Cf42 _dafny.Int
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf40 _dafny.Int, Cf41 _dafny.Sequence, Cf42 _dafny.Int) D7 {
	return D7{D7_DC23{Cf40, Cf41, Cf42}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC24 struct {
}

func (D7_DC24) isD7() {}

func (CompanionStruct_D7_) Create_DC24_() D7 {
	return D7{D7_DC24{}}
}

func (_this D7) Is_DC24() bool {
	_, ok := _this.Get_().(D7_DC24)
	return ok
}

type D7_DC21 struct {
	Cf37 _dafny.Map
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf37 _dafny.Map) D7 {
	return D7{D7_DC21{Cf37}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC25 struct {
	Cf43 D7
}

func (D7_DC25) isD7() {}

func (CompanionStruct_D7_) Create_DC25_(Cf43 D7) D7 {
	return D7{D7_DC25{Cf43}}
}

func (_this D7) Is_DC25() bool {
	_, ok := _this.Get_().(D7_DC25)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC22_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq)
}

func (_this D7) Dtor_cf38() _dafny.Array {
	return _this.Get_().(D7_DC22).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D7_DC22).Cf39
}

func (_this D7) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf40
}

func (_this D7) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D7_DC23).Cf41
}

func (_this D7) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf42
}

func (_this D7) Dtor_cf37() _dafny.Map {
	return _this.Get_().(D7_DC21).Cf37
}

func (_this D7) Dtor_cf43() D7 {
	return _this.Get_().(D7_DC25).Cf43
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf38) + ", " + data.Cf39.VerbatimString(true) + ")"
		}
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf40) + ", " + data.Cf41.VerbatimString(true) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D7_DC24:
		{
			return "D7.DC24"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf37) + ")"
		}
	case D7_DC25:
		{
			return "D7.DC25" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf38 == data2.Cf38 && data1.Cf39.Equals(data2.Cf39)
		}
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0 && data1.Cf41.Equals(data2.Cf41) && data1.Cf42.Cmp(data2.Cf42) == 0
		}
	case D7_DC24:
		{
			_, ok := other.Get_().(D7_DC24)
			return ok
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf37.Equals(data2.Cf37)
		}
	case D7_DC25:
		{
			data2, ok := other.Get_().(D7_DC25)
			return ok && data1.Cf43.Equals(data2.Cf43)
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
	Cf44 _dafny.Map
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf44 _dafny.Map) D8 {
	return D8{D8_DC26{Cf44}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D8) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D8_DC26).Cf44
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf44) + ")"
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
			return ok && data1.Cf44.Equals(data2.Cf44)
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
	Cf46 bool
	Cf47 _dafny.Int
	Cf48 _dafny.Int
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf46 bool, Cf47 _dafny.Int, Cf48 _dafny.Int) D9 {
	return D9{D9_DC28{Cf46, Cf47, Cf48}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

type D9_DC27 struct {
	Cf45 _dafny.MultiSet
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf45 _dafny.MultiSet) D9 {
	return D9{D9_DC27{Cf45}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC28_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf46() bool {
	return _this.Get_().(D9_DC28).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf47
}

func (_this D9) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D9_DC28).Cf48
}

func (_this D9) Dtor_cf45() _dafny.MultiSet {
	return _this.Get_().(D9_DC27).Cf45
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf45) + ")"
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
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D10_DC30 struct {
	Cf50 _dafny.Int
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf50 _dafny.Int) D10 {
	return D10{D10_DC30{Cf50}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

type D10_DC29 struct {
	Cf49 _dafny.Map
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf49 _dafny.Map) D10 {
	return D10{D10_DC29{Cf49}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC30_(_dafny.Zero)
}

func (_this D10) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D10_DC30).Cf50
}

func (_this D10) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D10_DC29).Cf49
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf49.Equals(data2.Cf49)
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

type D11_DC32 struct {
	Cf52 bool
	Cf53 bool
	Cf54 bool
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf52 bool, Cf53 bool, Cf54 bool) D11 {
	return D11{D11_DC32{Cf52, Cf53, Cf54}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

type D11_DC31 struct {
	Cf51 *C1
}

func (D11_DC31) isD11() {}

func (CompanionStruct_D11_) Create_DC31_(Cf51 *C1) D11 {
	return D11{D11_DC31{Cf51}}
}

func (_this D11) Is_DC31() bool {
	_, ok := _this.Get_().(D11_DC31)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC32_(false, false, false)
}

func (_this D11) Dtor_cf52() bool {
	return _this.Get_().(D11_DC32).Cf52
}

func (_this D11) Dtor_cf53() bool {
	return _this.Get_().(D11_DC32).Cf53
}

func (_this D11) Dtor_cf54() bool {
	return _this.Get_().(D11_DC32).Cf54
}

func (_this D11) Dtor_cf51() *C1 {
	return _this.Get_().(D11_DC31).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D11_DC31:
		{
			return "D11.DC31" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54
		}
	case D11_DC31:
		{
			data2, ok := other.Get_().(D11_DC31)
			return ok && data1.Cf51 == data2.Cf51
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
	Cf56 bool
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf56 bool) D12 {
	return D12{D12_DC34{Cf56}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC35 struct {
	Cf57 bool
	Cf58 _dafny.Int
	Cf59 _dafny.Int
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf57 bool, Cf58 _dafny.Int, Cf59 _dafny.Int) D12 {
	return D12{D12_DC35{Cf57, Cf58, Cf59}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
	Cf60 bool
	Cf61 bool
	Cf62 _dafny.Sequence
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf60 bool, Cf61 bool, Cf62 _dafny.Sequence) D12 {
	return D12{D12_DC36{Cf60, Cf61, Cf62}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC33 struct {
	Cf55 _dafny.MultiSet
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf55 _dafny.MultiSet) D12 {
	return D12{D12_DC33{Cf55}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC34_(false)
}

func (_this D12) Dtor_cf56() bool {
	return _this.Get_().(D12_DC34).Cf56
}

func (_this D12) Dtor_cf57() bool {
	return _this.Get_().(D12_DC35).Cf57
}

func (_this D12) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D12_DC35).Cf58
}

func (_this D12) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D12_DC35).Cf59
}

func (_this D12) Dtor_cf60() bool {
	return _this.Get_().(D12_DC36).Cf60
}

func (_this D12) Dtor_cf61() bool {
	return _this.Get_().(D12_DC36).Cf61
}

func (_this D12) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D12_DC36).Cf62
}

func (_this D12) Dtor_cf55() _dafny.MultiSet {
	return _this.Get_().(D12_DC33).Cf55
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + data.Cf62.VerbatimString(true) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf55) + ")"
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
			return ok && data1.Cf56 == data2.Cf56
		}
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0 && data1.Cf59.Cmp(data2.Cf59) == 0
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61 && data1.Cf62.Equals(data2.Cf62)
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D13_DC37 struct {
	Cf63 _dafny.Set
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf63 _dafny.Set) D13 {
	return D13{D13_DC37{Cf63}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D13) Dtor_cf63() _dafny.Set {
	return _this.Get_().(D13_DC37).Cf63
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf63) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf63.Equals(data2.Cf63)
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

type D14_DC38 struct {
	Cf64 _dafny.MultiSet
}

func (D14_DC38) isD14() {}

func (CompanionStruct_D14_) Create_DC38_(Cf64 _dafny.MultiSet) D14 {
	return D14{D14_DC38{Cf64}}
}

func (_this D14) Is_DC38() bool {
	_, ok := _this.Get_().(D14_DC38)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D14) Dtor_cf64() _dafny.MultiSet {
	return _this.Get_().(D14_DC38).Cf64
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC38:
		{
			return "D14.DC38" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC38:
		{
			data2, ok := other.Get_().(D14_DC38)
			return ok && data1.Cf64.Equals(data2.Cf64)
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
	Cf65 *C3
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf65 *C3) D15 {
	return D15{D15_DC39{Cf65}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

func (CompanionStruct_D15_) Default() *C3 {
	return (*C3)(nil)
}

func (_this D15) Dtor_cf65() *C3 {
	return _this.Get_().(D15_DC39).Cf65
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf65) + ")"
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
			return ok && data1.Cf65 == data2.Cf65
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

type D16_DC41 struct {
	Cf67 _dafny.Sequence
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf67 _dafny.Sequence) D16 {
	return D16{D16_DC41{Cf67}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC42 struct {
	Cf68 _dafny.Int
	Cf69 _dafny.Int
	Cf70 _dafny.CodePoint
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf68 _dafny.Int, Cf69 _dafny.Int, Cf70 _dafny.CodePoint) D16 {
	return D16{D16_DC42{Cf68, Cf69, Cf70}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC40 struct {
	Cf66 _dafny.Sequence
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf66 _dafny.Sequence) D16 {
	return D16{D16_DC40{Cf66}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC43 struct {
	Cf71 D16
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf71 D16) D16 {
	return D16{D16_DC43{Cf71}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC41_(_dafny.EmptySeq)
}

func (_this D16) Dtor_cf67() _dafny.Sequence {
	return _this.Get_().(D16_DC41).Cf67
}

func (_this D16) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf68
}

func (_this D16) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D16_DC42).Cf69
}

func (_this D16) Dtor_cf70() _dafny.CodePoint {
	return _this.Get_().(D16_DC42).Cf70
}

func (_this D16) Dtor_cf66() _dafny.Sequence {
	return _this.Get_().(D16_DC40).Cf66
}

func (_this D16) Dtor_cf71() D16 {
	return _this.Get_().(D16_DC43).Cf71
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69.Cmp(data2.Cf69) == 0 && data1.Cf70 == data2.Cf70
		}
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf71.Equals(data2.Cf71)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm0(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool
	Fm1(globalState *GlobalState) bool
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
	Fm0(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool
	Fm1(globalState *GlobalState) bool
	F6() _dafny.Int
	F6_set_(value _dafny.Int)
	Fm2(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int
	M1(globalState *GlobalState) (_dafny.Int, bool)
	M2(p0 _dafny.Int, globalState *GlobalState)
	F5() _dafny.Int
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
	_f0 _dafny.Int
	_f1 _dafny.Int
	_f3 _dafny.Int
	_f4 _dafny.MultiSet
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.Zero
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.EmptyMultiSet
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Int, f4 _dafny.MultiSet) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.MultiSet {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f8 _dafny.Int
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f8 = _dafny.Zero
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

func (_this *C0) Ctor__(f8 _dafny.Int) {
	{
		(_this)._f8 = f8
	}
}
func (_this *C0) Fm4(p0 _dafny.CodePoint, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C0) F8() _dafny.Int {
	{
		return _this._f8
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f6 _dafny.Int
	_f5 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f6 = _dafny.Zero
	_this._f5 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C1{}
var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F6() _dafny.Int {
	return _this._f6
}
func (_this *C1) F6_set_(value _dafny.Int) {
	_this._f6 = value
}
func (_this *C1) F5() _dafny.Int {
	return _this._f5
}
func (_this *C1) Ctor__(f5 _dafny.Int, f6 _dafny.Int) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C1) Fm0(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this.F6()).Cmp(_dafny.IntOfInt64(188)) > 0
	}
}
func (_this *C1) Fm1(globalState *GlobalState) bool {
	{
		return ((func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(110), _dafny.IntOfInt64(681))); ; {
				_compr_26, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _358_v0 _dafny.Int
				_358_v0 = interface{}(_compr_26).(_dafny.Int)
				if ((_dafny.IntOfInt64(110)).Cmp(_358_v0) <= 0) && ((_358_v0).Cmp(_dafny.IntOfInt64(681)) < 0) {
					_coll26.Add((_358_v0).Times(_this.F6()))
				}
			}
			return _coll26.ToSet()
		}()).Difference(func() _dafny.Set {
			var _coll27 = _dafny.NewBuilder()
			_ = _coll27
			for _iter28 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}(func(_359_i0 _dafny.Int) _dafny.Int {
				return _this.F6()
			}))).Cardinality()), _this.F6())).Elements()); ; {
				_compr_27, _ok28 := _iter28()
				if !_ok28 {
					break
				}
				var _360_v1 _dafny.Int
				_360_v1 = interface{}(_compr_27).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}(func(_359_i0 _dafny.Int) _dafny.Int {
					return _this.F6()
				}))).Cardinality()), _this.F6()), _360_v1) {
					_coll27.Add(Companion_Default___.SafeDivisionInt(_360_v1, _dafny.IntOfInt64(499)))
				}
			}
			return _coll27.ToSet()
		}())).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(231)))
	}
}
func (_this *C1) Fm2(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return _this.F6()
	}
}
func (_this *C1) Fm12(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.SetOf((_this).F5())).Intersection(_dafny.SetOf((_this).F5()))).Cardinality()
	}
}
func (_this *C1) Fm13(p0 _dafny.Int, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F5()
	}
}
func (_this *C1) M1(globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _361_v0 _dafny.Array
		_ = _361_v0
		var _len0_6 _dafny.Int = _dafny.One
		_ = _len0_6
		var _nw69 _dafny.Array
		_ = _nw69
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw69 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Sequence = func(_362_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("yjwhec")
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw69 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw69).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw69).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_361_v0 = _nw69
		var _363_v1 _dafny.Sequence
		_ = _363_v1
		_363_v1 = _dafny.UnicodeSeqOfUtf8Bytes("oll")
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))
		_ = _index50
		(_361_v0).ArraySet1(_363_v1, (_index50).Int())
		var _364_v2 _dafny.CodePoint
		_ = _364_v2
		_364_v2 = _dafny.CodePoint('r')
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))
		_ = _index51
		var _rhs88 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("alijtn")
		_ = _rhs88
		var _rhs89 _dafny.CodePoint = _364_v2
		_ = _rhs89
		var _lhs49 _dafny.Array = _361_v0
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))
		_ = _lhs50
		(_lhs49).ArraySet1(_rhs88, (_lhs50).Int())
		_364_v2 = _rhs89
		var _365_v3 _dafny.Array
		_ = _365_v3
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw70
		_365_v3 = _nw70
		for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_365_v3), 0))); ; {
			_guard_loop_1, _ok29 := _iter29()
			if !_ok29 {
				break
			}
			var _366_i1 _dafny.Int
			_366_i1 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_366_i1).Sign() != -1) && ((_366_i1).Cmp(_dafny.ArrayLen((_365_v3), 0)) < 0)) {
				(_365_v3).ArraySet1(!(!(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-399))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}((func(_367_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_368_i2 _dafny.Int) _dafny.CodePoint {
						return _367_v2
					}
				})(_364_v2))), _dafny.UnicodeSeqOfUtf8Bytes("djet")))) || (false), (_366_i1).Int())
			}
		}
		var _369_v4 _dafny.Array
		_ = _369_v4
		var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw71
		_369_v4 = _nw71
		var _370_v5 _dafny.Sequence
		_ = _370_v5
		_370_v5 = _dafny.SeqOf(_369_v4)
		var _371_v6 _dafny.Set
		_ = _371_v6
		_371_v6 = _dafny.SetOf((_370_v5).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_370_v5).Cardinality()))).Uint32()).(_dafny.Array))
		_371_v6 = (_dafny.SetOf(_369_v4, _369_v4)).Intersection(_371_v6)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
		_ = _index52
		(_369_v4).ArraySet1((_dafny.IntOfInt64(85)).Times((_this).F5()), (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
		_ = _index53
		(_369_v4).ArraySet1((_this).F5(), (_index53).Int())
		var _372_v7 _dafny.Sequence
		_ = _372_v7
		_372_v7 = _dafny.SeqOf((_this).F5())
		var _373_v8 _dafny.Set
		_ = _373_v8
		_373_v8 = _dafny.SetOf(_372_v7)
		var _374_v9 bool
		_ = _374_v9
		_374_v9 = true
		var _375_i3 _dafny.Int
		_ = _375_i3
		_375_i3 = _dafny.Zero
		{
			for ((_373_v8).Intersection(_373_v8)).IsProperSubsetOf(Companion_Default___.Fm14(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iwogcvy")).Cardinality()), _374_v9, globalState)) {
				{
					if (_375_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_375_i3 = (_375_i3).Plus(_dafny.One)
					(_this).F6_set_(_dafny.IntOfInt64(131))
					r1 = !(!(_374_v9) || (_374_v9))
					var _376_v10 _dafny.Map
					_ = _376_v10
					_376_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _this.F6())
					var _377_v11 _dafny.Map
					_ = _377_v11
					_377_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("tyctsial"), (_this).F5())).Dtor_cf6(), _dafny.SeqOf(_376_v10, _376_v10))
					_377_v11 = (_377_v11).Update((_this).F5(), Companion_Default___.Fm15(_364_v2, globalState))
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_365_v3), 0))
					_ = _index54
					(_365_v3).ArraySet1(_374_v9, (_index54).Int())
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_365_v3), 0))
					_ = _index55
					(_365_v3).ArraySet1(false, (_index55).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		if _dafny.Companion_Sequence_.Equal(_372_v7, _dafny.SeqOf(_this.F6())) {
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
			_ = _index56
			(_365_v3).ArraySet1(!(_374_v9), (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
			_ = _index57
			(_365_v3).ArraySet1(_374_v9, (_index57).Int())
			var _378_v12 D0
			_ = _378_v12
			_378_v12 = Companion_D0_.Create_DC0_(!(false))
			var _source6 D0 = _378_v12
			_ = _source6
			if _source6.Is_DC1() {
				var _379___mcc_h0 bool = _source6.Get_().(D0_DC1).Cf1
				_ = _379___mcc_h0
				var _380___mcc_h1 bool = _source6.Get_().(D0_DC1).Cf2
				_ = _380___mcc_h1
				var _381_cf2 bool = _380___mcc_h1
				_ = _381_cf2
				var _382_cf1 bool = _379___mcc_h0
				_ = _382_cf1
				(_this).F6_set_((_this).F5())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
				_ = _index58
				(_365_v3).ArraySet1(_374_v9, (_index58).Int())
				var _383_v13 *C0
				_ = _383_v13
				var _nw72 *C0 = New_C0_()
				_ = _nw72
				_nw72.Ctor__(_dafny.IntOfInt64(85))
				_383_v13 = _nw72
				var _384_v14 *C0
				_ = _384_v14
				var _nw73 *C0 = New_C0_()
				_ = _nw73
				_nw73.Ctor__((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				_384_v14 = _nw73
			} else if _source6.Is_DC2() {
				_363_v1 = (_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)
				var _385_v15 _dafny.Sequence
				_ = _385_v15
				_385_v15 = _dafny.SeqOf(_374_v9)
				var _386_v16 _dafny.Map
				_ = _386_v16
				_386_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-348), !((_385_v15).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F5()), _dafny.IntOfUint32((_385_v15).Cardinality()))).Uint32()).(bool)))
				var _387_v18 *C0
				_ = _387_v18
				var _nw74 *C0 = New_C0_()
				_ = _nw74
				_nw74.Ctor__(_this.F6())
				_387_v18 = _nw74
				var _388_v19 _dafny.Map
				_ = _388_v19
				_388_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_387_v18, (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				var _389_v20 _dafny.Set
				_ = _389_v20
				_389_v20 = _dafny.SetOf((_388_v19).Cardinality())
				var _390_v21 _dafny.MultiSet
				_ = _390_v21
				_390_v21 = _dafny.MultiSetOf(_374_v9, _374_v9, (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool))
				_374_v9 = ((_dafny.MultiSetOf(false)).Union(_390_v21)).IsProperSubsetOf(Companion_Default___.Fm16(_386_v16, (_this).F5(), func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter30 := _dafny.Iterate((_389_v20).Elements()); ; {
						_compr_28, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _391_v17 _dafny.Int
						_391_v17 = interface{}(_compr_28).(_dafny.Int)
						if (_389_v20).Contains(_391_v17) {
							_coll28.Add(Companion_Default___.SafeModuloInt(_391_v17, (_387_v18).F8()), false)
						}
					}
					return _coll28.ToMap()
				}(), globalState))
				var _392_v22 _dafny.Array
				_ = _392_v22
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_7
				var _nw75 _dafny.Array
				_ = _nw75
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw75 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) bool = (func(_393_v3 _dafny.Array) func(_dafny.Int) bool {
						return func(_394_i4 _dafny.Int) bool {
							return (_393_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_393_v3), 0))).Int()).(bool)
						}
					})(_365_v3)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw75 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw75).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw75).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_392_v22 = _nw75
				var _395_v23 _dafny.Map
				_ = _395_v23
				_395_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_387_v18, _392_v22)
				(globalState).F2 = Companion_Default___.SafeDivisionInt(_this.F6(), (_this).Fm12(_374_v9, (_395_v23).Cardinality(), (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), (Companion_Default___.Fm17((_387_v18).F8(), Companion_D0_.Create_DC0_(_374_v9), _this.F6(), true, globalState)).Cardinality(), globalState))
				var _396_v24 _dafny.Array
				_ = _396_v24
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw76
				_396_v24 = _nw76
				var _397_v25 D0
				_ = _397_v25
				_397_v25 = Companion_D0_.Create_DC1_(false, (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool))
				var _pat_let_tv6 = _374_v9
				_ = _pat_let_tv6
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
				_ = _index59
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
				_ = _index60
				var _rhs90 bool = (Companion_Default___.SafeModuloInt((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), (_this).F5())).Cmp(_this.F6()) >= 0
				_ = _rhs90
				var _rhs91 _dafny.Array = (func() _dafny.Array {
					if (func(_pat_let7_0 D0) D0 {
						return func(_398_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let8_0 bool) D0 {
								return func(_399_dt__update_hcf2_h0 bool) D0 {
									return Companion_D0_.Create_DC1_((_398_dt__update__tmp_h0).Dtor_cf1(), _399_dt__update_hcf2_h0)
								}(_pat_let8_0)
							}(_pat_let_tv6)
						}(_pat_let7_0)
					}(_397_v25)).Dtor_cf2() {
						return _396_v24
					}
					return _396_v24
				})()
				_ = _rhs91
				var _rhs92 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int)), _this.F6())
				_ = _rhs92
				var _rhs93 _dafny.Array = _392_v22
				_ = _rhs93
				var _rhs94 _dafny.Sequence = _372_v7
				_ = _rhs94
				var _lhs51 _dafny.Array = _365_v3
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
				_ = _lhs52
				var _lhs53 _dafny.Array = _369_v4
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
				_ = _lhs54
				(_lhs51).ArraySet1(_rhs90, (_lhs52).Int())
				_396_v24 = _rhs91
				(_lhs53).ArraySet1(_rhs92, (_lhs54).Int())
				_392_v22 = _rhs93
				_372_v7 = _rhs94
			} else if _source6.Is_DC3() {
				var _400___mcc_h2 bool = _source6.Get_().(D0_DC3).Cf3
				_ = _400___mcc_h2
				var _401_cf3 bool = _400___mcc_h2
				_ = _401_cf3
				_401_cf3 = (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)
				_363_v1 = (_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_8
				var _nw77 _dafny.Array
				_ = _nw77
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw77 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = func(_402_i5 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_402_i5, (_this).F5())
					}
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw77 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw77).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw77).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_369_v4 = _nw77
				var _403_v26 _dafny.Array
				_ = _403_v26
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_9
				var _nw78 _dafny.Array
				_ = _nw78
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw78 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) bool = (func(_404_v9 bool) func(_dafny.Int) bool {
						return func(_405_i6 _dafny.Int) bool {
							return _404_v9
						}
					})(_374_v9)
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw78 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw78).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw78).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_403_v26 = _nw78
				var _406_v27 _dafny.Map
				_ = _406_v27
				_406_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_cf3, (_this).F5())
				var _407_v28 _dafny.Map
				_ = _407_v28
				_407_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v26, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_cf3, (_this).F5()))
				var _408_v29 _dafny.Array
				_ = _408_v29
				var _nwElement0_12 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v26, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F6()))).Update(_403_v26, _406_v27)
				_ = _nwElement0_12
				var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(3))
				_ = _nw79
				(_nw79).ArraySet1(_nwElement0_12, 0)
				(_nw79).ArraySet1(_407_v28, 1)
				(_nw79).ArraySet1((_407_v28).Merge(_407_v28), 2)
				_408_v29 = _nw79
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_408_v29), 0))
				_ = _index61
				(_408_v29).ArraySet1(_407_v28, (_index61).Int())
				var _409_v30 _dafny.Map
				_ = _409_v30
				_409_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v9, _374_v9)
				var _410_v31 D3
				_ = _410_v31
				_410_v31 = Companion_D3_.Create_DC8_(_409_v30)
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_408_v29), 0))
				_ = _index62
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
				_ = _index63
				var _rhs95 _dafny.Int = ((_410_v31).Dtor_cf8()).Cardinality()
				_ = _rhs95
				var _rhs96 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)).Cardinality()), _this.F6()), (_dafny.MultiSetOf((_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence), _363_v1)).Cardinality()))
				_ = _rhs96
				var _rhs97 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v26, _406_v27)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v26, _406_v27))
				_ = _rhs97
				var _rhs98 bool = ((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int)).Cmp(_this.F6()) >= 0
				_ = _rhs98
				var _lhs55 *C1 = _this
				_ = _lhs55
				var _lhs56 *GlobalState = globalState
				_ = _lhs56
				var _lhs57 _dafny.Array = _408_v29
				_ = _lhs57
				var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(61), _dafny.ArrayLen((_408_v29), 0))
				_ = _lhs58
				var _lhs59 _dafny.Array = _365_v3
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))
				_ = _lhs60
				_lhs55.F6_set_(_rhs95)
				_lhs56.F2 = _rhs96
				(_lhs57).ArraySet1(_rhs97, (_lhs58).Int())
				(_lhs59).ArraySet1(_rhs98, (_lhs60).Int())
			} else {
				var _411___mcc_h3 bool = _source6.Get_().(D0_DC0).Cf0
				_ = _411___mcc_h3
				var _412_cf0 bool = _411___mcc_h3
				_ = _412_cf0
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
				_ = _index64
				(_369_v4).ArraySet1(((_dafny.Zero).Minus(_this.F6())).Minus(_this.F6()), (_index64).Int())
				var _413_v32 _dafny.MultiSet
				_ = _413_v32
				_413_v32 = _dafny.MultiSetOf(_372_v7)
				_413_v32 = _dafny.MultiSetOf(_dafny.SeqOf((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int)), _dafny.SeqOf((_this).Fm12(_374_v9, (_this).F5(), _374_v9, (_this).Fm12((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), _374_v9, _this.F6(), globalState), globalState)), (Companion_D4_.Create_DC11_(_372_v7)).Dtor_cf13(), _372_v7)
				var _414_v33 *C0
				_ = _414_v33
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__((_this).F5())
				_414_v33 = _nw80
				var _415_v34 _dafny.Sequence
				_ = _415_v34
				_415_v34 = _dafny.SeqOf(_412_cf0)
				var _416_v35 _dafny.MultiSet
				_ = _416_v35
				_416_v35 = _dafny.MultiSetOf(_dafny.IntOfUint32((_415_v34).Cardinality()), _dafny.IntOfInt64(979), _dafny.IntOfInt64(490))
				_416_v35 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_372_v7, _dafny.Companion_Sequence_.Concatenate(_372_v7, _372_v7)))
			}
			var _417_v36 _dafny.Map
			_ = _417_v36
			_417_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(134), _374_v9)
			var _418_v37 _dafny.MultiSet
			_ = _418_v37
			_418_v37 = _dafny.MultiSetOf(_417_v36)
			var _419_v38 _dafny.Map
			_ = _419_v38
			_419_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool))
			var _420_v39 _dafny.Set
			_ = _420_v39
			_420_v39 = _dafny.SetOf((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool))
			var _421_v40 _dafny.Array
			_ = _421_v40
			var _nwElement0_13 bool = _374_v9
			_ = _nwElement0_13
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(21))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_13, 0)
			(_nw81).ArraySet1(_374_v9, 1)
			(_nw81).ArraySet1(!((_this).Fm0(_418_v37, true, _363_v1, globalState)), 2)
			(_nw81).ArraySet1(false, 3)
			(_nw81).ArraySet1(!(true), 4)
			(_nw81).ArraySet1((_this).Fm1(globalState), 5)
			(_nw81).ArraySet1(!(_419_v38).Contains(_374_v9), 6)
			(_nw81).ArraySet1(_374_v9, 7)
			(_nw81).ArraySet1(_374_v9, 8)
			(_nw81).ArraySet1((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), 9)
			(_nw81).ArraySet1(_374_v9, 10)
			(_nw81).ArraySet1((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), 11)
			(_nw81).ArraySet1(_374_v9, 12)
			(_nw81).ArraySet1((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), 13)
			(_nw81).ArraySet1(_374_v9, 14)
			(_nw81).ArraySet1((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), 15)
			(_nw81).ArraySet1((_420_v39).IsSubsetOf(_420_v39), 16)
			(_nw81).ArraySet1((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), 17)
			(_nw81).ArraySet1(false, 18)
			(_nw81).ArraySet1(_374_v9, 19)
			(_nw81).ArraySet1(_374_v9, 20)
			_421_v40 = _nw81
			_421_v40 = _421_v40
			var _422_v41 _dafny.Array
			_ = _422_v41
			var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw82
			_422_v41 = _nw82
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_422_v41), 0))
			_ = _index65
			(_422_v41).ArraySet1(_421_v40, (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_422_v41), 0))
			_ = _index66
			(_422_v41).ArraySet1(_421_v40, (_index66).Int())
			var _423_v42 _dafny.Map
			_ = _423_v42
			_423_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_369_v4, (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
			var _424_v43 _dafny.MultiSet
			_ = _424_v43
			_424_v43 = _dafny.MultiSetOf(_this.F6())
			var _425_v44 _dafny.Map
			_ = _425_v44
			_425_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v42, _424_v43)
			var _426_v45 D3
			_ = _426_v45
			_426_v45 = Companion_D3_.Create_DC9_((func() _dafny.MultiSet {
				if (_425_v44).Contains((_423_v42).Update(_369_v4, _dafny.IntOfInt64(-664))) {
					return (_425_v44).Get((_423_v42).Update(_369_v4, _dafny.IntOfInt64(-664))).(_dafny.MultiSet)
				}
				return _424_v43
			})(), Companion_D0_.Create_DC2_(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(29))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_427_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_428_i7 _dafny.Int) _dafny.CodePoint {
					return _427_v2
				}
			})(_364_v2)))).Cardinality()))
			var _source7 D3 = _426_v45
			_ = _source7
			if _source7.Is_DC9() {
				var _429___mcc_h4 _dafny.MultiSet = _source7.Get_().(D3_DC9).Cf9
				_ = _429___mcc_h4
				var _430___mcc_h5 D0 = _source7.Get_().(D3_DC9).Cf10
				_ = _430___mcc_h5
				var _431___mcc_h6 _dafny.Int = _source7.Get_().(D3_DC9).Cf11
				_ = _431___mcc_h6
				var _432_cf11 _dafny.Int = _431___mcc_h6
				_ = _432_cf11
				var _433_cf10 D0 = _430___mcc_h5
				_ = _433_cf10
				var _434_cf9 _dafny.MultiSet = _429___mcc_h4
				_ = _434_cf9
				_369_v4 = _369_v4
				var _435_v46 _dafny.Map
				_ = _435_v46
				_435_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v9, _372_v7)
				var _436_v47 *C0
				_ = _436_v47
				var _nw83 *C0 = New_C0_()
				_ = _nw83
				_nw83.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg24 _dafny.Int) interface{} {
						return coer24(arg24)
					}
				}((func(_437_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_438_i8 _dafny.Int) _dafny.CodePoint {
						return _437_v2
					}
				})(_364_v2))), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_435_v46).Contains((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)) {
						return (_435_v46).Get((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg25 _dafny.Int) interface{} {
							return coer25(arg25)
						}
					}(func(_439_i9 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_this.F6())
					}))
				})()).Cardinality()))).Cardinality())
				_436_v47 = _nw83
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))
				_ = _index67
				(_361_v0).ArraySet1(_363_v1, (_index67).Int())
				(globalState).F2 = (_436_v47).F8()
			} else if _source7.Is_DC8() {
				var _440___mcc_h7 _dafny.Map = _source7.Get_().(D3_DC8).Cf8
				_ = _440___mcc_h7
				var _441_cf8 _dafny.Map = _440___mcc_h7
				_ = _441_cf8
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
				_ = _index68
				(_369_v4).ArraySet1((_dafny.Zero).Minus((_this).Fm12(true, (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), _this.F6(), globalState)), (_index68).Int())
				_374_v9 = (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)
				var _442_v48 *C0
				_ = _442_v48
				var _nw84 *C0 = New_C0_()
				_ = _nw84
				_nw84.Ctor__((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				_442_v48 = _nw84
				var _443_v49 _dafny.Map
				_ = _443_v49
				_443_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_374_v9) || ((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)), Companion_Default___.SafeDivisionInt((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), (_372_v7).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_372_v7).Cardinality()))).Uint32()).(_dafny.Int)))
				var _rhs99 _dafny.Int = (_dafny.Zero).Minus(_this.F6())
				_ = _rhs99
				var _rhs100 _dafny.Map = _443_v49
				_ = _rhs100
				var _rhs101 _dafny.CodePoint = _364_v2
				_ = _rhs101
				var _lhs61 *GlobalState = globalState
				_ = _lhs61
				_lhs61.F2 = _rhs99
				_443_v49 = _rhs100
				_364_v2 = _rhs101
			} else {
				var _444___mcc_h8 D3 = _source7.Get_().(D3_DC10).Cf12
				_ = _444___mcc_h8
				var _445_cf12 D3 = _444___mcc_h8
				_ = _445_cf12
				var _446_v50 _dafny.Sequence
				_ = _446_v50
				_446_v50 = _dafny.SeqOf(_374_v9)
				_446_v50 = _446_v50
				var _rhs102 _dafny.Array = _369_v4
				_ = _rhs102
				var _rhs103 bool = !(_424_v43).Contains((_dafny.Zero).Minus((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int)))
				_ = _rhs103
				_369_v4 = _rhs102
				_374_v9 = _rhs103
				var _447_v51 _dafny.Array
				_ = _447_v51
				var _nwElement0_14 _dafny.CodePoint = _364_v2
				_ = _nwElement0_14
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(23))
				_ = _nw85
				(_nw85).ArraySet1CodePoint(_nwElement0_14, 0)
				(_nw85).ArraySet1CodePoint(_364_v2, 1)
				(_nw85).ArraySet1CodePoint(_364_v2, 2)
				(_nw85).ArraySet1CodePoint(_364_v2, 3)
				(_nw85).ArraySet1CodePoint(_364_v2, 4)
				(_nw85).ArraySet1CodePoint(Companion_Default___.Fm18((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), _this.F6(), (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), globalState), 5)
				(_nw85).ArraySet1CodePoint(_364_v2, 6)
				(_nw85).ArraySet1CodePoint(_364_v2, 7)
				(_nw85).ArraySet1CodePoint(_364_v2, 8)
				(_nw85).ArraySet1CodePoint(_364_v2, 9)
				(_nw85).ArraySet1CodePoint((func() _dafny.CodePoint {
					if (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool) {
						return _364_v2
					}
					return Companion_Default___.Fm18((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), _this.F6(), _dafny.IntOfInt64(-788), _374_v9, globalState)
				})(), 10)
				(_nw85).ArraySet1CodePoint(_364_v2, 11)
				(_nw85).ArraySet1CodePoint(_364_v2, 12)
				(_nw85).ArraySet1CodePoint(_364_v2, 13)
				(_nw85).ArraySet1CodePoint(_364_v2, 14)
				(_nw85).ArraySet1CodePoint(_364_v2, 15)
				(_nw85).ArraySet1CodePoint(_364_v2, 16)
				(_nw85).ArraySet1CodePoint(_364_v2, 17)
				(_nw85).ArraySet1CodePoint(_364_v2, 18)
				(_nw85).ArraySet1CodePoint(_364_v2, 19)
				(_nw85).ArraySet1CodePoint(_364_v2, 20)
				(_nw85).ArraySet1CodePoint(_364_v2, 21)
				(_nw85).ArraySet1CodePoint(_364_v2, 22)
				_447_v51 = _nw85
				_447_v51 = _447_v51
				var _448_v52 D0
				_ = _448_v52
				_448_v52 = Companion_D0_.Create_DC3_((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool))
				var _449_v53 _dafny.Array
				_ = _449_v53
				var _nwElement0_15 D0 = _448_v52
				_ = _nwElement0_15
				var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(10))
				_ = _nw86
				(_nw86).ArraySet1(_nwElement0_15, 0)
				(_nw86).ArraySet1(_448_v52, 1)
				(_nw86).ArraySet1(_448_v52, 2)
				(_nw86).ArraySet1(Companion_D0_.Create_DC3_(_374_v9), 3)
				(_nw86).ArraySet1(_448_v52, 4)
				(_nw86).ArraySet1(_448_v52, 5)
				(_nw86).ArraySet1(Companion_D0_.Create_DC3_((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool)), 6)
				(_nw86).ArraySet1(_448_v52, 7)
				(_nw86).ArraySet1(_448_v52, 8)
				(_nw86).ArraySet1(_448_v52, 9)
				_449_v53 = _nw86
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_449_v53), 0))
				_ = _index69
				(_449_v53).ArraySet1(_448_v52, (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_449_v53), 0))
				_ = _index70
				(_449_v53).ArraySet1(_448_v52, (_index70).Int())
			}
		} else {
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))
			_ = _index71
			(_361_v0).ArraySet1((_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence), (_index71).Int())
			(globalState).F2 = Companion_Default___.SafeModuloInt((_this).Fm13(_this.F6(), _dafny.MultiSetOf(_363_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-505))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_450_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_451_i10 _dafny.Int) _dafny.CodePoint {
					return _450_v2
				}
			})(_364_v2))), _363_v1, (_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)), globalState), _dafny.IntOfInt64(393))
			if _374_v9 {
				var _452_v54 _dafny.Sequence
				_ = _452_v54
				_452_v54 = _dafny.SeqOf((func() _dafny.Sequence {
					if _374_v9 {
						return _363_v1
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("ltkjnhdi")
				})())
				var _453_v55 _dafny.Map
				_ = _453_v55
				_453_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v9, _374_v9)
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))
				_ = _index72
				(_369_v4).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_452_v54, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_453_v55).Cardinality(), _dafny.IntOfInt64(835)), _dafny.IntOfUint32((_452_v54).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}((func(_454_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_455_i11 _dafny.Int) _dafny.CodePoint {
						return _454_v2
					}
				})(_364_v2))), (_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)))).Cardinality()), (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))
				_ = _index73
				(_365_v3).ArraySet1(((_this).F5()).Cmp((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int)) != 0, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))
				_ = _index74
				(_365_v3).ArraySet1((_this).Fm1(globalState), (_index74).Int())
				var _456_v56 *C0
				_ = _456_v56
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				_456_v56 = _nw87
				var _457_v57 _dafny.MultiSet
				_ = _457_v57
				_457_v57 = _dafny.MultiSetOf((_456_v56).F8(), (_this).F5(), _this.F6(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_363_v1).Cardinality())), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				var _458_v58 _dafny.Map
				_ = _458_v58
				_458_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_363_v1, (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				var _459_v59 _dafny.Map
				_ = _459_v59
				_459_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), (_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				var _460_v60 _dafny.Map
				_ = _460_v60
				_460_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), _dafny.MultiSetFromSeq(_372_v7))
				var _461_v61 _dafny.Array
				_ = _461_v61
				var _nwElement0_16 _dafny.MultiSet = _457_v57
				_ = _nwElement0_16
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(29))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_16, 0)
				(_nw88).ArraySet1(_dafny.MultiSetOf((_this).F5()), 1)
				(_nw88).ArraySet1(_457_v57, 2)
				(_nw88).ArraySet1(_457_v57, 3)
				(_nw88).ArraySet1(_457_v57, 4)
				(_nw88).ArraySet1((_dafny.MultiSetFromSeq(_372_v7)).Difference(_457_v57), 5)
				(_nw88).ArraySet1(_457_v57, 6)
				(_nw88).ArraySet1((_dafny.MultiSetFromSeq(_372_v7)).Union(_457_v57), 7)
				(_nw88).ArraySet1(_457_v57, 8)
				(_nw88).ArraySet1((Companion_Default___.Fm19((_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), (_456_v56).F8(), (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool), globalState)).Update(_this.F6(), Companion_Default___.Abs((func() _dafny.Int {
					if (_458_v58).Contains((_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)) {
						return (_458_v58).Get((_361_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_361_v0), 0))).Int()).(_dafny.Sequence)).(_dafny.Int)
					}
					return (_this).F5()
				})())), 9)
				(_nw88).ArraySet1(_457_v57, 10)
				(_nw88).ArraySet1((_dafny.MultiSetOf((_this).F5())).Update((func() _dafny.Int {
					if (_459_v59).Contains((_this).F5()) {
						return (_459_v59).Get((_this).F5()).(_dafny.Int)
					}
					return _this.F6()
				})(), Companion_Default___.Abs(_dafny.IntOfInt64(690))), 11)
				(_nw88).ArraySet1(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_363_v1).Cardinality())), (_457_v57).Cardinality()), 12)
				(_nw88).ArraySet1(_457_v57, 13)
				(_nw88).ArraySet1(_457_v57, 14)
				(_nw88).ArraySet1(_457_v57, 15)
				(_nw88).ArraySet1((_457_v57).Union((func() _dafny.MultiSet {
					if (_460_v60).Contains(false) {
						return (_460_v60).Get(false).(_dafny.MultiSet)
					}
					return _457_v57
				})()), 16)
				(_nw88).ArraySet1(_457_v57, 17)
				(_nw88).ArraySet1(_457_v57, 18)
				(_nw88).ArraySet1(_457_v57, 19)
				(_nw88).ArraySet1(_457_v57, 20)
				(_nw88).ArraySet1(_457_v57, 21)
				(_nw88).ArraySet1((_457_v57).Union(_457_v57), 22)
				(_nw88).ArraySet1(_457_v57, 23)
				(_nw88).ArraySet1(_dafny.MultiSetOf((_this).F5()), 24)
				(_nw88).ArraySet1(_457_v57, 25)
				(_nw88).ArraySet1(_457_v57, 26)
				(_nw88).ArraySet1((func() _dafny.MultiSet {
					if (_365_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))).Int()).(bool) {
						return _457_v57
					}
					return _dafny.MultiSetFromSeq(_372_v7)
				})(), 27)
				(_nw88).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_372_v7, _dafny.SeqOf((_456_v56).F8()))), 28)
				_461_v61 = _nw88
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_461_v61), 0))
				_ = _index75
				(_461_v61).ArraySet1((_457_v57).Intersection(_457_v57), (_index75).Int())
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_461_v61), 0))
				_ = _index76
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))
				_ = _index77
				var _rhs104 _dafny.MultiSet = _457_v57
				_ = _rhs104
				var _rhs105 bool = !(!(_374_v9)) || (!(_374_v9) || (!(!(_374_v9))))
				_ = _rhs105
				var _rhs106 _dafny.Sequence = _dafny.SeqOf(_363_v1)
				_ = _rhs106
				var _lhs62 _dafny.Array = _461_v61
				_ = _lhs62
				var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(293), _dafny.ArrayLen((_461_v61), 0))
				_ = _lhs63
				var _lhs64 _dafny.Array = _365_v3
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_365_v3), 0))
				_ = _lhs65
				(_lhs62).ArraySet1(_rhs104, (_lhs63).Int())
				(_lhs64).ArraySet1(_rhs105, (_lhs65).Int())
				_452_v54 = _rhs106
				var _462_v62 _dafny.Array
				_ = _462_v62
				var _463_v63 bool
				_ = _463_v63
				var _464_v64 _dafny.MultiSet
				_ = _464_v64
				var _465_v65 _dafny.Int
				_ = _465_v65
				var _out28 _dafny.Array
				_ = _out28
				var _out29 bool
				_ = _out29
				var _out30 _dafny.MultiSet
				_ = _out30
				var _out31 _dafny.Int
				_ = _out31
				_out28, _out29, _out30, _out31 = Companion_Default___.M0(_364_v2, globalState)
				_462_v62 = _out28
				_463_v63 = _out29
				_464_v64 = _out30
				_465_v65 = _out31
			} else {
				var _466_v66 *C0
				_ = _466_v66
				var _nw89 *C0 = New_C0_()
				_ = _nw89
				_nw89.Ctor__((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int))
				_466_v66 = _nw89
				var _467_v67 D4
				_ = _467_v67
				_467_v67 = Companion_D4_.Create_DC14_((_369_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_369_v4), 0))).Int()).(_dafny.Int), false, (_this).F5(), _374_v9, (_466_v66).F8())
				(globalState).F2 = (_467_v67).Dtor_cf22()
				var _468_v68 _dafny.Array
				_ = _468_v68
				var _469_v69 bool
				_ = _469_v69
				var _470_v70 _dafny.MultiSet
				_ = _470_v70
				var _471_v71 _dafny.Int
				_ = _471_v71
				var _out32 _dafny.Array
				_ = _out32
				var _out33 bool
				_ = _out33
				var _out34 _dafny.MultiSet
				_ = _out34
				var _out35 _dafny.Int
				_ = _out35
				_out32, _out33, _out34, _out35 = Companion_Default___.M0(_dafny.CodePoint('p'), globalState)
				_468_v68 = _out32
				_469_v69 = _out33
				_470_v70 = _out34
				_471_v71 = _out35
				r1 = _469_v69
				r1 = _374_v9
			}
			r1 = _374_v9
			_374_v9 = !(_374_v9)
		}
		r0 = Companion_Default___.SafeModuloInt((_this).F5(), (_this).F5())
		r1 = _374_v9
		return r0, r1
	}
}
func (_this *C1) M2(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _472_v0 *C0
		_ = _472_v0
		var _nw90 *C0 = New_C0_()
		_ = _nw90
		_nw90.Ctor__(_this.F6())
		_472_v0 = _nw90
		var _473_v1 _dafny.Array
		_ = _473_v1
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_10
		var _nw91 _dafny.Array
		_ = _nw91
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw91 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Int = (func(_474_v0 *C0) func(_dafny.Int) _dafny.Int {
				return func(_475_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_475_i1, (_474_v0).F8())
				}
			})(_472_v0)
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw91 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw91).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw91).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_473_v1 = _nw91
		for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_473_v1), 0))); ; {
			_guard_loop_2, _ok31 := _iter31()
			if !_ok31 {
				break
			}
			var _476_i0 _dafny.Int
			_476_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_476_i0).Sign() != -1) && ((_476_i0).Cmp(_dafny.ArrayLen((_473_v1), 0)) < 0)) {
				(_473_v1).ArraySet1((_476_i0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-155))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_477_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality())), (_476_i0).Int())
			}
		}
		var _478_v2 bool
		_ = _478_v2
		_478_v2 = true
		(_this).F6_set_((func() _dafny.Int {
			if _478_v2 {
				return (_472_v0).F8()
			}
			return _this.F6()
		})())
		var _479_v3 _dafny.Array
		_ = _479_v3
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_11
		var _nw92 _dafny.Array
		_ = _nw92
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw92 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) bool = (func(_480_v0 *C0) func(_dafny.Int) bool {
				return func(_481_i4 _dafny.Int) bool {
					return ((_480_v0).F8()).Cmp(_this.F6()) == 0
				}
			})(_472_v0)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw92 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw92).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw92).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_479_v3 = _nw92
		for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_479_v3), 0))); ; {
			_guard_loop_3, _ok32 := _iter32()
			if !_ok32 {
				break
			}
			var _482_i3 _dafny.Int
			_482_i3 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_482_i3).Sign() != -1) && ((_482_i3).Cmp(_dafny.ArrayLen((_479_v3), 0)) < 0)) {
				(_479_v3).ArraySet1(_478_v2, (_482_i3).Int())
			}
		}
		var _483_v4 _dafny.Array
		_ = _483_v4
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_12
		var _nw93 _dafny.Array
		_ = _nw93
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw93 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Map = (func(_484_v2 bool, _485_v0 *C0) func(_dafny.Int) _dafny.Map {
				return func(_486_i6 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_484_v2, (_485_v0).F8())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_484_v2), (_485_v0).F8()))
				}
			})(_478_v2, _472_v0)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw93 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw93).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw93).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_483_v4 = _nw93
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_483_v4), 0))); ; {
			_guard_loop_4, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _487_i5 _dafny.Int
			_487_i5 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_487_i5).Sign() != -1) && ((_487_i5).Cmp(_dafny.ArrayLen((_483_v4), 0)) < 0)) {
				(_483_v4).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), ((func() _dafny.Map {
					if !(_478_v2) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfUint32((_dafny.SeqOf(Companion_D1_.Create_DC4_(_dafny.SetOf((_dafny.Zero).Minus((_this).F5()))), Companion_D1_.Create_DC4_(_dafny.SetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bof")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_478_v2, _478_v2)).Cardinality(), (_dafny.Zero).Minus((_472_v0).F8()), (_472_v0).F8(), _this.F6())))).Cardinality()))
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), (_dafny.Zero).Minus((_472_v0).F8()))
				})()).Cardinality()), (_487_i5).Int())
			}
		}
		var _488_v5 _dafny.CodePoint
		_ = _488_v5
		_488_v5 = _dafny.CodePoint('m')
		var _489_v6 _dafny.Set
		_ = _489_v6
		_489_v6 = _dafny.SetOf(_478_v2, _478_v2, true, _478_v2, _478_v2)
		var _rhs107 _dafny.Array = _479_v3
		_ = _rhs107
		var _rhs108 _dafny.CodePoint = (func() _dafny.CodePoint {
			if true {
				return _488_v5
			}
			return _488_v5
		})()
		_ = _rhs108
		var _rhs109 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
			if _478_v2 {
				return (_this).Fm2((_472_v0).F8(), _489_v6, globalState)
			}
			return (_472_v0).F8()
		})())
		_ = _rhs109
		var _rhs110 _dafny.Int = (_472_v0).F8()
		_ = _rhs110
		var _rhs111 _dafny.Int = (_472_v0).F8()
		_ = _rhs111
		var _lhs66 *GlobalState = globalState
		_ = _lhs66
		var _lhs67 *C1 = _this
		_ = _lhs67
		var _lhs68 *C1 = _this
		_ = _lhs68
		_479_v3 = _rhs107
		_488_v5 = _rhs108
		_lhs66.F2 = _rhs109
		_lhs67.F6_set_(_rhs110)
		_lhs68.F6_set_(_rhs111)
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f6 _dafny.Int
	_f5 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f6 = _dafny.Zero
	_this._f5 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C2{}
var _ T1 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F6() _dafny.Int {
	return _this._f6
}
func (_this *C2) F6_set_(value _dafny.Int) {
	_this._f6 = value
}
func (_this *C2) F5() _dafny.Int {
	return _this._f5
}
func (_this *C2) Ctor__(f5 _dafny.Int, f6 _dafny.Int) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C2) Fm0(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf(_dafny.SeqOf(_dafny.IntOfInt64(-363), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(true, true)).Cardinality())).Cardinality()))).IsDisjointFrom((_dafny.MultiSetOf(_dafny.SeqOf(_this.F6()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(567))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_490_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((_dafny.SetOf(false, false)).Cardinality(), _dafny.IntOfInt64(-401))
		})))))
	}
}
func (_this *C2) Fm1(globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_491_i0 _dafny.Int) _dafny.Int {
			return (_this).F5()
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg31 _dafny.Int) interface{} {
				return coer31(arg31)
			}
		}(func(_492_i1 _dafny.Int) _dafny.Int {
			return (_this).F5()
		}))), (_this).F5())
	}
}
func (_this *C2) Fm2(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-302)
	}
}
func (_this *C2) Fm8(p0 _dafny.Sequence, p1 bool, p2 bool, p3 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus((_this).F5())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("j")).Cardinality()))) > 0
	}
}
func (_this *C2) Fm9(p0 D1, p1 bool, globalState *GlobalState) bool {
	{
		return true
	}
}
func (_this *C2) M1(globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _493_v0 bool
		_ = _493_v0
		_493_v0 = true
		r1 = _493_v0
		var _494_v1 _dafny.CodePoint
		_ = _494_v1
		_494_v1 = _dafny.CodePoint('q')
		var _495_v2 _dafny.Array
		_ = _495_v2
		var _496_v3 bool
		_ = _496_v3
		var _497_v4 _dafny.MultiSet
		_ = _497_v4
		var _498_v5 _dafny.Int
		_ = _498_v5
		var _out36 _dafny.Array
		_ = _out36
		var _out37 bool
		_ = _out37
		var _out38 _dafny.MultiSet
		_ = _out38
		var _out39 _dafny.Int
		_ = _out39
		_out36, _out37, _out38, _out39 = Companion_Default___.M0(_494_v1, globalState)
		_495_v2 = _out36
		_496_v3 = _out37
		_497_v4 = _out38
		_498_v5 = _out39
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))
		_ = _index78
		(_495_v2).ArraySet1(_498_v5, (_index78).Int())
		var _499_v6 D1
		_ = _499_v6
		_499_v6 = Companion_D1_.Create_DC5_()
		var _500_v7 _dafny.MultiSet
		_ = _500_v7
		_500_v7 = _dafny.MultiSetOf(_496_v3)
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))
		_ = _index79
		(_495_v2).ArraySet1((func() _dafny.Int {
			if (_this).Fm9(_499_v6, _493_v0, globalState) {
				return (_500_v7).Cardinality()
			}
			return (_this).F5()
		})(), (_index79).Int())
		var _hi4 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(247), _498_v5)
		_ = _hi4
		for _501_i0 := (_dafny.IntOfInt64(282)).Plus((_dafny.Zero).Minus(_498_v5)); _501_i0.Cmp(_hi4) < 0; _501_i0 = _501_i0.Plus(_dafny.One) {
			var _502_v8 _dafny.Sequence
			_ = _502_v8
			_502_v8 = _dafny.SeqOf((_this.F6()).Cmp(_501_i0) != 0)
			if (_502_v8).Select((Companion_Default___.SafeIndex((_this.F6()).Minus(_501_i0), _dafny.IntOfUint32((_502_v8).Cardinality()))).Uint32()).(bool) {
				var _503_v9 _dafny.MultiSet
				_ = _503_v9
				_503_v9 = _dafny.MultiSetOf(_501_i0)
				var _504_v10 *C0
				_ = _504_v10
				var _nw94 *C0 = New_C0_()
				_ = _nw94
				_nw94.Ctor__((_503_v9).Cardinality())
				_504_v10 = _nw94
				var _505_v11 _dafny.Sequence
				_ = _505_v11
				_505_v11 = _dafny.UnicodeSeqOfUtf8Bytes("ij")
				var _506_v12 _dafny.Map
				_ = _506_v12
				_506_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_493_v0, _505_v11)
				_505_v11 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_506_v12).Contains(_496_v3) {
						return (_506_v12).Get(_496_v3).(_dafny.Sequence)
					}
					return _505_v11
				})(), (Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_506_v12).Contains(_496_v3) {
						return (_506_v12).Get(_496_v3).(_dafny.Sequence)
					}
					return _505_v11
				})()).Cardinality()))).Uint32(), _494_v1), _505_v11)
				var _507_v13 _dafny.Map
				_ = _507_v13
				_507_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_498_v5), _493_v0)
				var _508_v14 _dafny.MultiSet
				_ = _508_v14
				_508_v14 = _dafny.MultiSetOf(_507_v13, _507_v13, _507_v13)
				var _509_v15 _dafny.Array
				_ = _509_v15
				var _nwElement0_17 bool = !((func() bool {
					if (_507_v13).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_502_v8, (Companion_Default___.SafeIndex((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_502_v8).Cardinality()))).Uint32(), _496_v3)).Cardinality())) {
						return (_507_v13).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_502_v8, (Companion_Default___.SafeIndex((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_502_v8).Cardinality()))).Uint32(), _496_v3)).Cardinality())).(bool)
					}
					return true
				})()) || (_496_v3)
				_ = _nwElement0_17
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(8))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_17, 0)
				(_nw95).ArraySet1(_496_v3, 1)
				(_nw95).ArraySet1((func() bool {
					if _496_v3 {
						return false
					}
					return _493_v0
				})(), 2)
				(_nw95).ArraySet1(_496_v3, 3)
				(_nw95).ArraySet1((_this).Fm8(_505_v11, _496_v3, (_this).Fm0(_508_v14, _496_v3, _505_v11, globalState), _493_v0, globalState), 4)
				(_nw95).ArraySet1(_496_v3, 5)
				(_nw95).ArraySet1(false, 6)
				(_nw95).ArraySet1(_496_v3, 7)
				_509_v15 = _nw95
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_509_v15), 0))
				_ = _index80
				(_509_v15).ArraySet1(_493_v0, (_index80).Int())
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_509_v15), 0))
				_ = _index81
				(_509_v15).ArraySet1(!((_503_v9).Union(_503_v9)).Equals(_503_v9), (_index81).Int())
				var _510_v16 _dafny.Set
				_ = _510_v16
				_510_v16 = _dafny.SetOf(_496_v3, _493_v0, _493_v0)
				var _511_v17 *C0
				_ = _511_v17
				var _nw96 *C0 = New_C0_()
				_ = _nw96
				_nw96.Ctor__((func() _dafny.Int {
					if _493_v0 {
						return (_510_v16).Cardinality()
					}
					return _501_i0
				})())
				_511_v17 = _nw96
				_496_v3 = (_509_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_509_v15), 0))).Int()).(bool)
			} else {
				var _512_v18 _dafny.Array
				_ = _512_v18
				var _513_v19 bool
				_ = _513_v19
				var _514_v20 _dafny.MultiSet
				_ = _514_v20
				var _515_v21 _dafny.Int
				_ = _515_v21
				var _out40 _dafny.Array
				_ = _out40
				var _out41 bool
				_ = _out41
				var _out42 _dafny.MultiSet
				_ = _out42
				var _out43 _dafny.Int
				_ = _out43
				_out40, _out41, _out42, _out43 = Companion_Default___.M0(_494_v1, globalState)
				_512_v18 = _out40
				_513_v19 = _out41
				_514_v20 = _out42
				_515_v21 = _out43
				r1 = _493_v0
				var _516_v22 *C0
				_ = _516_v22
				var _nw97 *C0 = New_C0_()
				_ = _nw97
				_nw97.Ctor__((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int))
				_516_v22 = _nw97
				var _517_v23 _dafny.Sequence
				_ = _517_v23
				_517_v23 = _dafny.SeqOf((_516_v22).F8(), _501_i0)
				var _518_v24 _dafny.Sequence
				_ = _518_v24
				_518_v24 = _dafny.UnicodeSeqOfUtf8Bytes("gygdbkdr")
				var _519_v25 _dafny.Set
				_ = _519_v25
				_519_v25 = _dafny.SetOf(_501_i0)
				var _520_v26 _dafny.Array
				_ = _520_v26
				var _nwElement0_18 bool = true
				_ = _nwElement0_18
				var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(12))
				_ = _nw98
				(_nw98).ArraySet1(_nwElement0_18, 0)
				(_nw98).ArraySet1(((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int)).Cmp((_517_v23).Select((Companion_Default___.SafeIndex(_501_i0, _dafny.IntOfUint32((_517_v23).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, 1)
				(_nw98).ArraySet1((false) && (_496_v3), 2)
				(_nw98).ArraySet1((_this).Fm8(_518_v24, (_502_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-830))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_521_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(163)
				}))).Cardinality()), _dafny.IntOfUint32((_502_v8).Cardinality()))).Uint32()).(bool), _513_v19, _496_v3, globalState), 3)
				(_nw98).ArraySet1((_dafny.SetOf((_500_v7).Cardinality(), _dafny.IntOfUint32((_502_v8).Cardinality()))).IsDisjointFrom(_519_v25), 4)
				(_nw98).ArraySet1(_493_v0, 5)
				(_nw98).ArraySet1((_this).Fm9(_499_v6, _513_v19, globalState), 6)
				(_nw98).ArraySet1((!(_496_v3)) == (_496_v3), 7)
				(_nw98).ArraySet1(!(_493_v0), 8)
				(_nw98).ArraySet1(_493_v0, 9)
				(_nw98).ArraySet1(_513_v19, 10)
				(_nw98).ArraySet1(!(_513_v19), 11)
				_520_v26 = _nw98
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw99
				_520_v26 = _nw99
				_496_v3 = !(true)
			}
			var _522_v27 _dafny.MultiSet
			_ = _522_v27
			_522_v27 = _dafny.MultiSetOf(_501_i0, _dafny.IntOfUint32((_502_v8).Cardinality()))
			var _523_v28 _dafny.Set
			_ = _523_v28
			_523_v28 = _dafny.SetOf(_522_v27)
			var _524_v29 _dafny.Map
			_ = _524_v29
			_524_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_498_v5, _523_v28)
			var _525_v30 _dafny.Sequence
			_ = _525_v30
			_525_v30 = _dafny.SeqOf(_522_v27)
			_524_v29 = (_524_v29).Update((_this).F5(), func() _dafny.Set {
				var _coll29 = _dafny.NewBuilder()
				_ = _coll29
				for _iter34 := _dafny.Iterate((_525_v30).Elements()); ; {
					_compr_29, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _526_v31 _dafny.MultiSet
					_526_v31 = interface{}(_compr_29).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_525_v30, _526_v31) {
						_coll29.Add(_526_v31)
					}
				}
				return _coll29.ToSet()
			}())
			var _527_v32 _dafny.Map
			_ = _527_v32
			_527_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(516), _493_v0)
			var _528_v33 _dafny.Sequence
			_ = _528_v33
			_528_v33 = _dafny.SeqOf(_527_v32)
			var _529_v34 _dafny.Sequence
			_ = _529_v34
			_529_v34 = _dafny.SeqOf(_528_v33)
			var _530_v35 _dafny.Sequence
			_ = _530_v35
			_530_v35 = _dafny.UnicodeSeqOfUtf8Bytes("cldrgrfef")
			var _531_v36 _dafny.Array
			_ = _531_v36
			var _nwElement0_19 bool = _496_v3
			_ = _nwElement0_19
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(12))
			_ = _nw100
			(_nw100).ArraySet1(_nwElement0_19, 0)
			(_nw100).ArraySet1((func() bool {
				if _496_v3 {
					return _493_v0
				}
				return (_502_v8).Select((Companion_Default___.SafeIndex((_500_v7).Cardinality(), _dafny.IntOfUint32((_502_v8).Cardinality()))).Uint32()).(bool)
			})(), 1)
			(_nw100).ArraySet1(_496_v3, 2)
			(_nw100).ArraySet1((_this).Fm0(_dafny.MultiSetFromSeq((_529_v34).Select((Companion_Default___.SafeIndex((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_529_v34).Cardinality()))).Uint32()).(_dafny.Sequence)), false, _530_v35, globalState), 3)
			(_nw100).ArraySet1((func() bool {
				if _493_v0 {
					return false
				}
				return _493_v0
			})(), 4)
			(_nw100).ArraySet1((_493_v0) && (_496_v3), 5)
			(_nw100).ArraySet1(_493_v0, 6)
			(_nw100).ArraySet1(_493_v0, 7)
			(_nw100).ArraySet1((_498_v5).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).F5()))) > 0, 8)
			(_nw100).ArraySet1(_496_v3, 9)
			(_nw100).ArraySet1(_496_v3, 10)
			(_nw100).ArraySet1((func() bool {
				if _493_v0 {
					return _493_v0
				}
				return _496_v3
			})(), 11)
			_531_v36 = _nw100
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_531_v36), 0))
			_ = _index82
			(_531_v36).ArraySet1(_496_v3, (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_531_v36), 0))
			_ = _index83
			(_531_v36).ArraySet1(((_500_v7).Intersection(_500_v7)).IsSubsetOf((_dafny.MultiSetFromSeq(_502_v8)).Union(_500_v7)), (_index83).Int())
			var _532_v37 _dafny.Array
			_ = _532_v37
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_13
			var _nw101 _dafny.Array
			_ = _nw101
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw101 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Map = (func(_533_v36 _dafny.Array, _534_v0 bool) func(_dafny.Int) _dafny.Map {
					return func(_535_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_533_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(486), _dafny.ArrayLen((_533_v36), 0))).Int()).(bool), _534_v0)
					}
				})(_531_v36, _493_v0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw101 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw101).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw101).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_532_v37 = _nw101
			var _536_v38 _dafny.Map
			_ = _536_v38
			_536_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F5()), (_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int))
			var _537_v39 _dafny.Map
			_ = _537_v39
			_537_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_493_v0, true)
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_532_v37), 0))
			_ = _index84
			(_532_v37).ArraySet1((Companion_Default___.Fm10((_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int), _536_v38, !(_496_v3), (_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int), globalState)).Merge(_537_v39), (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_532_v37), 0))
			_ = _index85
			(_532_v37).ArraySet1(_537_v39, (_index85).Int())
		}
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))
		_ = _index86
		(_495_v2).ArraySet1(_this.F6(), (_index86).Int())
		var _538_v40 _dafny.Sequence
		_ = _538_v40
		_538_v40 = _dafny.SeqOf(_dafny.IntOfInt64(777), (_495_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))).Int()).(_dafny.Int))
		var _539_v41 _dafny.Map
		_ = _539_v41
		_539_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_493_v0, _496_v3)
		var _540_v42 _dafny.Set
		_ = _540_v42
		_540_v42 = _dafny.SetOf(_493_v0)
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_495_v2), 0))
		_ = _index87
		(_495_v2).ArraySet1((_498_v5).Minus(((_538_v40).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_538_v40).Cardinality()))).Uint32()).(_dafny.Int)).Plus((_this).Fm2((_539_v41).Cardinality(), _540_v42, globalState))), (_index87).Int())
		r0 = (_538_v40).Select((Companion_Default___.SafeIndex(_498_v5, _dafny.IntOfUint32((_538_v40).Cardinality()))).Uint32()).(_dafny.Int)
		r1 = _493_v0
		return r0, r1
	}
}
func (_this *C2) M2(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _541_v0 bool
		_ = _541_v0
		_541_v0 = false
		if _541_v0 {
			var _542_v1 _dafny.Sequence
			_ = _542_v1
			_542_v1 = _dafny.SeqOf(p0, _this.F6())
			var _543_v2 _dafny.CodePoint
			_ = _543_v2
			_543_v2 = _dafny.CodePoint('q')
			var _544_v3 _dafny.Map
			_ = _544_v3
			_544_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v2, (_dafny.Zero).Minus((_this).F5()))
			var _545_v4 _dafny.Sequence
			_ = _545_v4
			_545_v4 = _dafny.UnicodeSeqOfUtf8Bytes("rwfr")
			var _rhs112 _dafny.Sequence = _542_v1
			_ = _rhs112
			var _rhs113 bool = _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm11((_dafny.Zero).Minus((_dafny.Zero).Minus(p0)), (_this).F5(), _this.F6(), _this.F6(), globalState), Companion_Default___.Fm11((_this).F5(), ((_544_v3).Update((_545_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_545_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _dafny.IntOfInt64(996))).Cardinality(), p0, _dafny.IntOfUint32((_545_v4).Cardinality()), globalState))
			_ = _rhs113
			_542_v1 = _rhs112
			_541_v0 = _rhs113
			var _546_v5 D1
			_ = _546_v5
			_546_v5 = Companion_D1_.Create_DC6_(_545_v4, _this.F6())
			var _547_v6 D0
			_ = _547_v6
			_547_v6 = Companion_D0_.Create_DC0_(false)
			var _548_v7 _dafny.MultiSet
			_ = _548_v7
			_548_v7 = _dafny.MultiSetOf((_547_v6).Dtor_cf0(), _541_v0)
			var _549_v8 _dafny.MultiSet
			_ = _549_v8
			_549_v8 = _dafny.MultiSetOf(_548_v7)
			_541_v0 = (_this).Fm8(_dafny.Companion_Sequence_.Concatenate(_545_v4, (_546_v5).Dtor_cf5()), !(_541_v0) || (!(false)), _541_v0, (_549_v8).IsDisjointFrom(_549_v8), globalState)
			var _550_v9 _dafny.Array
			_ = _550_v9
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(3))
			_ = _nw102
			_550_v9 = _nw102
			var _551_v10 _dafny.Set
			_ = _551_v10
			_551_v10 = _dafny.SetOf(_541_v0, _541_v0, false)
			var _552_v11 _dafny.MultiSet
			_ = _552_v11
			_552_v11 = _dafny.MultiSetOf(_543_v2, _543_v2)
			var _553_v12 _dafny.MultiSet
			_ = _553_v12
			_553_v12 = _dafny.MultiSetOf((_this).F5(), (_this).Fm2((_this).F5(), _551_v10, globalState), (func() _dafny.Int {
				if (_552_v11).Contains(_543_v2) {
					return (_552_v11).Multiplicity(_543_v2)
				}
				return (func() _dafny.Int {
					if (_544_v3).Contains(_543_v2) {
						return (_544_v3).Get(_543_v2).(_dafny.Int)
					}
					return _dafny.IntOfInt64(660)
				})()
			})(), p0, (_this).F5())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_550_v9), 0))
			_ = _index88
			(_550_v9).ArraySet1(_553_v12, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_550_v9), 0))
			_ = _index89
			(_550_v9).ArraySet1(_dafny.MultiSetOf(p0), (_index89).Int())
			var _554_v13 _dafny.Set
			_ = _554_v13
			_554_v13 = _dafny.SetOf(p0)
			var _555_v14 D1
			_ = _555_v14
			_555_v14 = Companion_D1_.Create_DC4_(_554_v13)
			var _source8 D1 = _555_v14
			_ = _source8
			if _source8.Is_DC5() {
				var _556_v15 _dafny.Array
				_ = _556_v15
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw103
				_556_v15 = _nw103
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_556_v15), 0))
				_ = _index90
				(_556_v15).ArraySet1((_this).Fm2(_this.F6(), _551_v10, globalState), (_index90).Int())
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_556_v15), 0))
				_ = _index91
				var _rhs114 bool = _541_v0
				_ = _rhs114
				var _rhs115 bool = (_552_v11).IsDisjointFrom(_552_v11)
				_ = _rhs115
				var _rhs116 _dafny.Array = (func() _dafny.Array {
					if _541_v0 {
						return _556_v15
					}
					return _556_v15
				})()
				_ = _rhs116
				var _rhs117 _dafny.Int = (_dafny.Zero).Minus((_this).Fm2(_this.F6(), (_dafny.SetOf(_541_v0)).Difference(_551_v10), globalState))
				_ = _rhs117
				var _lhs69 _dafny.Array = _556_v15
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_556_v15), 0))
				_ = _lhs70
				_541_v0 = _rhs114
				_541_v0 = _rhs115
				_556_v15 = _rhs116
				(_lhs69).ArraySet1(_rhs117, (_lhs70).Int())
				(_this).F6_set_((_this).F5())
				var _557_v16 _dafny.Array
				_ = _557_v16
				var _nw104 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw104
				_557_v16 = _nw104
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_557_v16), 0))
				_ = _index92
				(_557_v16).ArraySet1(!(_541_v0), (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_557_v16), 0))
				_ = _index93
				(_557_v16).ArraySet1(_541_v0, (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_556_v15), 0))
				_ = _index94
				(_556_v15).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("npeqafbc"), _545_v4)).Cardinality()), (_index94).Int())
			} else if _source8.Is_DC6() {
				var _558___mcc_h0 _dafny.Sequence = _source8.Get_().(D1_DC6).Cf5
				_ = _558___mcc_h0
				var _559___mcc_h1 _dafny.Int = _source8.Get_().(D1_DC6).Cf6
				_ = _559___mcc_h1
				var _560_cf6 _dafny.Int = _559___mcc_h1
				_ = _560_cf6
				var _561_cf5 _dafny.Sequence = _558___mcc_h0
				_ = _561_cf5
				var _562_v17 _dafny.Array
				_ = _562_v17
				var _nwElement0_20 _dafny.Int = _560_cf6
				_ = _nwElement0_20
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(3))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_20, 0)
				(_nw105).ArraySet1(_this.F6(), 1)
				(_nw105).ArraySet1(_this.F6(), 2)
				_562_v17 = _nw105
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_562_v17), 0))
				_ = _index95
				(_562_v17).ArraySet1(_dafny.IntOfInt64(728), (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_562_v17), 0))
				_ = _index96
				(_562_v17).ArraySet1((_this).F5(), (_index96).Int())
				var _563_v18 *C0
				_ = _563_v18
				var _nw106 *C0 = New_C0_()
				_ = _nw106
				_nw106.Ctor__(p0)
				_563_v18 = _nw106
				var _564_v19 _dafny.MultiSet
				_ = _564_v19
				_564_v19 = _dafny.MultiSetOf((func() *C0 {
					if _541_v0 {
						return _563_v18
					}
					return _563_v18
				})(), _563_v18)
				_564_v19 = (_564_v19).Union(_564_v19)
				(globalState).F2 = (_this).F5()
			} else {
				var _565___mcc_h2 _dafny.Set = _source8.Get_().(D1_DC4).Cf4
				_ = _565___mcc_h2
				var _566_cf4 _dafny.Set = _565___mcc_h2
				_ = _566_cf4
				var _567_v20 D0
				_ = _567_v20
				_567_v20 = Companion_D0_.Create_DC2_()
				var _568_v21 _dafny.Array
				_ = _568_v21
				var _nwElement0_21 D0 = (func() D0 {
					if _541_v0 {
						return _567_v20
					}
					return _567_v20
				})()
				_ = _nwElement0_21
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(15))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_21, 0)
				(_nw107).ArraySet1(Companion_D0_.Create_DC2_(), 1)
				(_nw107).ArraySet1(_567_v20, 2)
				(_nw107).ArraySet1(_567_v20, 3)
				(_nw107).ArraySet1(_567_v20, 4)
				(_nw107).ArraySet1(_567_v20, 5)
				(_nw107).ArraySet1(_567_v20, 6)
				(_nw107).ArraySet1(_567_v20, 7)
				(_nw107).ArraySet1(_567_v20, 8)
				(_nw107).ArraySet1(Companion_D0_.Create_DC2_(), 9)
				(_nw107).ArraySet1(_567_v20, 10)
				(_nw107).ArraySet1(_567_v20, 11)
				(_nw107).ArraySet1(Companion_D0_.Create_DC2_(), 12)
				(_nw107).ArraySet1(Companion_D0_.Create_DC2_(), 13)
				(_nw107).ArraySet1((func() D0 {
					if _541_v0 {
						return _567_v20
					}
					return _567_v20
				})(), 14)
				_568_v21 = _nw107
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_568_v21), 0))
				_ = _index97
				(_568_v21).ArraySet1(Companion_D0_.Create_DC2_(), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_568_v21), 0))
				_ = _index98
				(_568_v21).ArraySet1(_567_v20, (_index98).Int())
				var _569_v22 _dafny.Map
				_ = _569_v22
				_569_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_v4, (((_550_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_550_v9), 0))).Int()).(_dafny.MultiSet)).Cardinality()).Times(p0))
				_569_v22 = (_569_v22).Update((func() _dafny.Sequence {
					if _541_v0 {
						return _545_v4
					}
					return _545_v4
				})(), ((_554_v13).Cardinality()).Times((_this).F5()))
				var _570_v23 _dafny.Array
				_ = _570_v23
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw108
				_570_v23 = _nw108
				var _571_v24 T1
				_ = _571_v24
				var _nw109 *C1 = New_C1_()
				_ = _nw109
				_nw109.Ctor__((_dafny.Zero).Minus(p0), p0)
				_571_v24 = _nw109
				var _572_v25 _dafny.Set
				_ = _572_v25
				_572_v25 = _dafny.SetOf(_571_v24, _571_v24, _571_v24, _571_v24, _571_v24)
				var _573_v26 _dafny.Array
				_ = _573_v26
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw110
				_573_v26 = _nw110
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_573_v26), 0))
				_ = _index99
				(_573_v26).ArraySet1(_570_v23, (_index99).Int())
				var _574_v27 _dafny.Map
				_ = _574_v27
				_574_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _541_v0)
				var _575_v28 D5
				_ = _575_v28
				_575_v28 = Companion_D5_.Create_DC17_(_this.F6(), _570_v23, _574_v27)
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_573_v26), 0))
				_ = _index100
				var _rhs118 _dafny.Array = _570_v23
				_ = _rhs118
				var _rhs119 _dafny.Set = _572_v25
				_ = _rhs119
				var _rhs120 _dafny.Array = (_575_v28).Dtor_cf28()
				_ = _rhs120
				var _rhs121 bool = ((p0).Cmp((_this).Fm2((_this).F5(), _551_v10, globalState)) < 0) && (((_this).F5()).Cmp(p0) == 0)
				_ = _rhs121
				var _rhs122 bool = ((_548_v7).Difference(_dafny.MultiSetOf(!(true)))).IsProperSubsetOf(_548_v7)
				_ = _rhs122
				var _lhs71 _dafny.Array = _573_v26
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(249), _dafny.ArrayLen((_573_v26), 0))
				_ = _lhs72
				_570_v23 = _rhs118
				_572_v25 = _rhs119
				(_lhs71).ArraySet1(_rhs120, (_lhs72).Int())
				_541_v0 = _rhs121
				_541_v0 = _rhs122
				(_this).F6_set_((_571_v24).F5())
			}
			_541_v0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_545_v4, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gyqisghb"), _545_v4))
		} else {
			var _576_v30 _dafny.Map
			_ = _576_v30
			_576_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _541_v0)
			var _577_v31 *C1
			_ = _577_v31
			var _nw111 *C1 = New_C1_()
			_ = _nw111
			_nw111.Ctor__((_this).F5(), (_this).F5())
			_577_v31 = _nw111
			var _578_v32 _dafny.Map
			_ = _578_v32
			_578_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _577_v31)
			var _579_v33 _dafny.Sequence
			_ = _579_v33
			_579_v33 = _dafny.UnicodeSeqOfUtf8Bytes("ynpxt")
			var _580_v34 _dafny.Map
			_ = _580_v34
			_580_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_579_v33, _578_v32)
			var _581_v35 _dafny.Array
			_ = _581_v35
			var _nwElement0_22 _dafny.Int = p0
			_ = _nwElement0_22
			var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(28))
			_ = _nw112
			(_nw112).ArraySet1(_nwElement0_22, 0)
			(_nw112).ArraySet1(p0, 1)
			(_nw112).ArraySet1(_this.F6(), 2)
			(_nw112).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F5(), _this.F6()), 3)
			(_nw112).ArraySet1((_dafny.Zero).Minus((_this).F5()), 4)
			(_nw112).ArraySet1((_this).F5(), 5)
			(_nw112).ArraySet1((func() _dafny.Set {
				var _coll30 = _dafny.NewBuilder()
				_ = _coll30
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(590), _dafny.IntOfInt64(716))); ; {
					_compr_30, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _582_v29 _dafny.Int
					_582_v29 = interface{}(_compr_30).(_dafny.Int)
					if ((_dafny.IntOfInt64(590)).Cmp(_582_v29) <= 0) && ((_582_v29).Cmp(_dafny.IntOfInt64(716)) < 0) {
						_coll30.Add((_582_v29).Times(_this.F6()))
					}
				}
				return _coll30.ToSet()
			}()).Cardinality(), 6)
			(_nw112).ArraySet1(((_576_v30).Cardinality()).Minus(_dafny.IntOfInt64(775)), 7)
			(_nw112).ArraySet1(_this.F6(), 8)
			(_nw112).ArraySet1(((_578_v32).Merge((func() _dafny.Map {
				if (_580_v34).Contains(_579_v33) {
					return (_580_v34).Get(_579_v33).(_dafny.Map)
				}
				return _578_v32
			})())).Cardinality(), 9)
			(_nw112).ArraySet1(p0, 10)
			(_nw112).ArraySet1(_this.F6(), 11)
			(_nw112).ArraySet1(Companion_Default___.SafeModuloInt((_this).F5(), (_this).F5()), 12)
			(_nw112).ArraySet1(((_this).F5()).Plus(_dafny.IntOfInt64(57)), 13)
			(_nw112).ArraySet1(p0, 14)
			(_nw112).ArraySet1((_this).F5(), 15)
			(_nw112).ArraySet1(p0, 16)
			(_nw112).ArraySet1((func() _dafny.Int {
				if _541_v0 {
					return _this.F6()
				}
				return (_this).F5()
			})(), 17)
			(_nw112).ArraySet1(_this.F6(), 18)
			(_nw112).ArraySet1((_this).F5(), 19)
			(_nw112).ArraySet1(p0, 20)
			(_nw112).ArraySet1((_this).F5(), 21)
			(_nw112).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if true {
					return _this.F6()
				}
				return p0
			})()), 22)
			(_nw112).ArraySet1((_this.F6()).Plus(p0), 23)
			(_nw112).ArraySet1(_this.F6(), 24)
			(_nw112).ArraySet1(_this.F6(), 25)
			(_nw112).ArraySet1((p0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fmkmdiv")).Cardinality())), 26)
			(_nw112).ArraySet1(_this.F6(), 27)
			_581_v35 = _nw112
			_581_v35 = (Companion_D5_.Create_DC17_(_this.F6(), _581_v35, _576_v30)).Dtor_cf28()
			var _583_v36 D5
			_ = _583_v36
			_583_v36 = Companion_D5_.Create_DC16_(_581_v35)
			var _584_v37 _dafny.Sequence
			_ = _584_v37
			_584_v37 = _dafny.SeqOf(_dafny.SetOf(_583_v36, _583_v36, _583_v36, _583_v36, _583_v36))
			var _rhs123 bool = _541_v0
			_ = _rhs123
			var _rhs124 bool = _541_v0
			_ = _rhs124
			var _rhs125 _dafny.Sequence = _584_v37
			_ = _rhs125
			var _rhs126 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tomjydsx"), _dafny.Companion_Sequence_.Concatenate(_579_v33, _579_v33))).Cardinality())
			_ = _rhs126
			var _rhs127 bool = !(true) || (_541_v0)
			_ = _rhs127
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			_541_v0 = _rhs123
			_541_v0 = _rhs124
			_584_v37 = _rhs125
			_lhs73.F2 = _rhs126
			_541_v0 = _rhs127
			var _585_v38 _dafny.Map
			_ = _585_v38
			_585_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _this.F6())
			var _586_v39 _dafny.Map
			_ = _586_v39
			_586_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _585_v38)
			_586_v39 = (_586_v39).Update(p0, (func() _dafny.Map {
				if !(!(_541_v0)) {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _this.F6())
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _this.F6())
			})())
			var _587_v40 _dafny.Sequence
			_ = _587_v40
			_587_v40 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(47))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_588_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			}))).Cardinality()))
			var _589_v41 _dafny.Array
			_ = _589_v41
			var _nwElement0_23 bool = _541_v0
			_ = _nwElement0_23
			var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(3))
			_ = _nw113
			(_nw113).ArraySet1(_nwElement0_23, 0)
			(_nw113).ArraySet1(!_dafny.Companion_Sequence_.Equal(_587_v40, _dafny.Companion_Sequence_.Update(_587_v40, (Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_587_v40).Cardinality()))).Uint32(), _dafny.IntOfUint32((_587_v40).Cardinality()))), 1)
			(_nw113).ArraySet1(_541_v0, 2)
			_589_v41 = _nw113
			_589_v41 = _589_v41
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_581_v35), 0))
			_ = _index101
			(_581_v35).ArraySet1(_this.F6(), (_index101).Int())
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_581_v35), 0))
			_ = _index102
			(_581_v35).ArraySet1(p0, (_index102).Int())
		}
		var _590_v42 _dafny.Array
		_ = _590_v42
		var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(14))
		_ = _nw114
		_590_v42 = _nw114
		var _591_v43 _dafny.CodePoint
		_ = _591_v43
		_591_v43 = _dafny.CodePoint('a')
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
		_ = _index103
		(_590_v42).ArraySet1CodePoint(_591_v43, (_index103).Int())
		var _592_v44 _dafny.MultiSet
		_ = _592_v44
		_592_v44 = _dafny.MultiSetOf(_541_v0, _541_v0)
		var _593_v45 _dafny.MultiSet
		_ = _593_v45
		_593_v45 = _dafny.MultiSetOf(_592_v44)
		var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
		_ = _index104
		(_590_v42).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_593_v45).IsDisjointFrom(_593_v45) {
				return _591_v43
			}
			return _591_v43
		})(), (_index104).Int())
		var _594_v46 _dafny.Sequence
		_ = _594_v46
		_594_v46 = _dafny.SeqOf(_541_v0, _541_v0)
		if (_594_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-612), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32()).(bool) {
			var _595_v47 _dafny.Array
			_ = _595_v47
			var _nw115 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw115
			_595_v47 = _nw115
			var _596_v48 _dafny.Map
			_ = _596_v48
			_596_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v44, _541_v0)
			var _597_v49 _dafny.Sequence
			_ = _597_v49
			_597_v49 = _dafny.SeqOf(p0, (_596_v48).Cardinality())
			var _598_v50 _dafny.Sequence
			_ = _598_v50
			_598_v50 = _dafny.SeqOf((_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), _591_v43)
			var _599_v51 _dafny.Sequence
			_ = _599_v51
			_599_v51 = _dafny.SeqOf((_597_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_598_v50).Cardinality()), _dafny.IntOfUint32((_597_v49).Cardinality()))).Uint32()).(_dafny.Int))
			var _600_v52 _dafny.Map
			_ = _600_v52
			_600_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_597_v49).Cardinality()), _541_v0)
			var _601_v53 D5
			_ = _601_v53
			_601_v53 = Companion_D5_.Create_DC17_((_this).F5(), _595_v47, (_600_v52).Merge(_600_v52))
			_601_v53 = _601_v53
			(globalState).F2 = (_this).F5()
			if (_541_v0) == ((_this.F6()).Cmp(p0) <= 0) {
				var _602_v54 _dafny.Set
				_ = _602_v54
				_602_v54 = _dafny.SetOf(_592_v44)
				var _603_v55 _dafny.MultiSet
				_ = _603_v55
				_603_v55 = _dafny.MultiSetOf(_602_v54)
				var _604_v57 _dafny.Map
				_ = _604_v57
				_604_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_591_v43, _541_v0)
				(_this).F6_set_(((p0).Plus((_this).F5())).Times((func() _dafny.Int {
					if (_603_v55).Contains(_602_v54) {
						return (_603_v55).Multiplicity(_602_v54)
					}
					return (func() _dafny.Map {
						var _coll31 = _dafny.NewMapBuilder()
						_ = _coll31
						for _iter36 := _dafny.Iterate((_604_v57).Keys().Elements()); ; {
							_compr_31, _ok36 := _iter36()
							if !_ok36 {
								break
							}
							var _605_v56 _dafny.CodePoint
							_605_v56 = interface{}(_compr_31).(_dafny.CodePoint)
							if (_604_v57).Contains(_605_v56) {
								_coll31.Add(_605_v56, _541_v0)
							}
						}
						return _coll31.ToMap()
					}()).Cardinality()
				})()))
				var _606_v58 _dafny.Map
				_ = _606_v58
				_606_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v47, !(_541_v0) || (_541_v0))
				var _607_v59 _dafny.Sequence
				_ = _607_v59
				_607_v59 = _dafny.SeqOf(_606_v58)
				_606_v58 = (_606_v58).Merge((_607_v59).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-172), _dafny.IntOfUint32((_607_v59).Cardinality()))).Uint32()).(_dafny.Map))
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
				_ = _index105
				(_590_v42).ArraySet1CodePoint((_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), (_index105).Int())
				_598_v50 = _598_v50
				var _608_v60 _dafny.Map
				_ = _608_v60
				_608_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), (_this).F5())
				_608_v60 = (_608_v60).Update(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_598_v50, _598_v50)).Cardinality()))
			} else {
				(globalState).F2 = (_this).F5()
				var _609_v61 _dafny.MultiSet
				_ = _609_v61
				_609_v61 = _dafny.MultiSetOf(_dafny.IntOfUint32((_598_v50).Cardinality()), _this.F6())
				var _610_v62 D0
				_ = _610_v62
				_610_v62 = Companion_D0_.Create_DC2_()
				var _611_v63 D3
				_ = _611_v63
				_611_v63 = Companion_D3_.Create_DC9_(_dafny.MultiSetOf(p0), _610_v62, (_this).F5())
				var _612_v64 _dafny.Set
				_ = _612_v64
				_612_v64 = _dafny.SetOf(_this.F6())
				var _613_v65 _dafny.MultiSet
				_ = _613_v65
				_613_v65 = _dafny.MultiSetOf((_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))
				var _614_v66 _dafny.Array
				_ = _614_v66
				var _nwElement0_24 bool = _541_v0
				_ = _nwElement0_24
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(20))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_24, 0)
				(_nw116).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_597_v49, _599_v51), 1)
				(_nw116).ArraySet1(false, 2)
				(_nw116).ArraySet1(!(false), 3)
				(_nw116).ArraySet1((_dafny.MultiSetOf((_611_v63).Dtor_cf11())).IsSubsetOf(_609_v61), 4)
				(_nw116).ArraySet1(!(_612_v64).Equals(_dafny.SetOf((_this).F5())), 5)
				(_nw116).ArraySet1(_541_v0, 6)
				(_nw116).ArraySet1(_541_v0, 7)
				(_nw116).ArraySet1((_592_v44).IsDisjointFrom(_592_v44), 8)
				(_nw116).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_598_v50, _598_v50), 9)
				(_nw116).ArraySet1((_609_v61).IsDisjointFrom(_dafny.MultiSetOf((_this).F5(), p0)), 10)
				(_nw116).ArraySet1(_541_v0, 11)
				(_nw116).ArraySet1((_this.F6()).Cmp(_dafny.IntOfInt64(888)) > 0, 12)
				(_nw116).ArraySet1(_541_v0, 13)
				(_nw116).ArraySet1(_541_v0, 14)
				(_nw116).ArraySet1(!(_541_v0), 15)
				(_nw116).ArraySet1(((_dafny.MultiSetOf(_dafny.CodePoint('b'), _591_v43)).Update(_591_v43, Companion_Default___.Abs(p0))).IsSubsetOf(_613_v65), 16)
				(_nw116).ArraySet1(_541_v0, 17)
				(_nw116).ArraySet1(_541_v0, 18)
				(_nw116).ArraySet1(_541_v0, 19)
				_614_v66 = _nw116
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_614_v66), 0))
				_ = _index106
				(_614_v66).ArraySet1(_541_v0, (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_614_v66), 0))
				_ = _index107
				(_614_v66).ArraySet1(((_this).F5()).Cmp((_this).F5()) >= 0, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_614_v66), 0))
				_ = _index108
				(_614_v66).ArraySet1(_541_v0, (_index108).Int())
				var _615_v67 _dafny.MultiSet
				_ = _615_v67
				_615_v67 = _dafny.MultiSetOf(_595_v47)
				var _616_v68 _dafny.Map
				_ = _616_v68
				_616_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v67, _597_v49)
				_616_v68 = (_616_v68).Update(_615_v67, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F5()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_599_v51).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_this).F5())).Cardinality()))).Uint32(), p0), _597_v49))
				var _617_v69 _dafny.Set
				_ = _617_v69
				_617_v69 = _dafny.SetOf(_614_v66, _614_v66)
				_617_v69 = _617_v69
			}
			var _618_v70 _dafny.Array
			_ = _618_v70
			var _619_v71 bool
			_ = _619_v71
			var _620_v72 _dafny.MultiSet
			_ = _620_v72
			var _621_v73 _dafny.Int
			_ = _621_v73
			var _out44 _dafny.Array
			_ = _out44
			var _out45 bool
			_ = _out45
			var _out46 _dafny.MultiSet
			_ = _out46
			var _out47 _dafny.Int
			_ = _out47
			_out44, _out45, _out46, _out47 = Companion_Default___.M0(_dafny.CodePoint('l'), globalState)
			_618_v70 = _out44
			_619_v71 = _out45
			_620_v72 = _out46
			_621_v73 = _out47
			var _622_v74 _dafny.Set
			_ = _622_v74
			_622_v74 = _dafny.SetOf((_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))
			_622_v74 = _622_v74
		} else {
			var _623_v75 _dafny.Map
			_ = _623_v75
			_623_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _541_v0)
			var _624_v76 _dafny.Sequence
			_ = _624_v76
			_624_v76 = _dafny.UnicodeSeqOfUtf8Bytes("ycjiany")
			_541_v0 = (func() bool {
				if (_623_v75).Contains((_dafny.IntOfUint32((_624_v76).Cardinality())).Cmp(_this.F6()) > 0) {
					return (_623_v75).Get((_dafny.IntOfUint32((_624_v76).Cardinality())).Cmp(_this.F6()) > 0).(bool)
				}
				return ((_dafny.Zero).Minus(_this.F6())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wml")).Cardinality())) > 0
			})()
			var _625_v77 _dafny.Array
			_ = _625_v77
			var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw117
			_625_v77 = _nw117
			var _626_v78 _dafny.Set
			_ = _626_v78
			_626_v78 = _dafny.SetOf((_this).F5(), p0, (_this).F5(), p0)
			var _627_v79 _dafny.Map
			_ = _627_v79
			_627_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_626_v78).Cardinality(), _dafny.IntOfInt64(952))
			var _628_v80 _dafny.Sequence
			_ = _628_v80
			_628_v80 = _dafny.SeqOf(p0)
			var _629_v81 *C0
			_ = _629_v81
			var _nw118 *C0 = New_C0_()
			_ = _nw118
			_nw118.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, (func() _dafny.Int {
				if (_627_v79).Contains(p0) {
					return (_627_v79).Get(p0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_628_v80).Cardinality())
			})())).Cardinality())
			_629_v81 = _nw118
			var _630_v82 _dafny.Sequence
			_ = _630_v82
			_630_v82 = _dafny.SeqOf(_629_v81, _629_v81)
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_625_v77), 0))
			_ = _index109
			(_625_v77).ArraySet1(_630_v82, (_index109).Int())
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_625_v77), 0))
			_ = _index110
			(_625_v77).ArraySet1(_630_v82, (_index110).Int())
			var _631_v83 *C1
			_ = _631_v83
			var _nw119 *C1 = New_C1_()
			_ = _nw119
			_nw119.Ctor__(_dafny.IntOfInt64(579), _dafny.IntOfInt64(464))
			_631_v83 = _nw119
			_541_v0 = !((func() bool {
				if _541_v0 {
					return (_594_v46).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_626_v78).Cardinality()), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32()).(bool)
				}
				return _541_v0
			})())
			var _632_v84 _dafny.Array
			_ = _632_v84
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_14
			var _nw120 _dafny.Array
			_ = _nw120
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw120 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = func(_633_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_633_i1, (_this).F5())
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
			_632_v84 = _nw120
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_632_v84), 0))
			_ = _index111
			(_632_v84).ArraySet1((_dafny.Zero).Minus((_this.F6()).Times(_this.F6())), (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(436), _dafny.ArrayLen((_632_v84), 0))
			_ = _index112
			(_632_v84).ArraySet1((_dafny.Zero).Minus((p0).Minus(_this.F6())), (_index112).Int())
		}
		var _634_i2 _dafny.Int
		_ = _634_i2
		_634_i2 = _dafny.Zero
		{
			for !(!(_541_v0)) || (_541_v0) {
				{
					if (_634_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_634_i2 = (_634_i2).Plus(_dafny.One)
					(_this).M4((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(89))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg34 _dafny.Int) interface{} {
							return coer34(arg34)
						}
					}(func(_635_i3 _dafny.Int) _dafny.Int {
						return (_this).F5()
					})))).IsProperSubsetOf(_dafny.MultiSetOf((_this).F5())), globalState)
					(_this).F6_set_(_dafny.IntOfInt64(841))
					var _636_v86 _dafny.Sequence
					_ = _636_v86
					_636_v86 = _dafny.SeqOf(func() _dafny.Set {
						var _coll32 = _dafny.NewBuilder()
						_ = _coll32
						for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(379), _dafny.IntOfInt64(-220))); ; {
							_compr_32, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _637_v85 _dafny.Int
							_637_v85 = interface{}(_compr_32).(_dafny.Int)
							if ((_dafny.IntOfInt64(379)).Cmp(_637_v85) <= 0) && ((_637_v85).Cmp(_dafny.IntOfInt64(-220)) < 0) {
								_coll32.Add(Companion_Default___.SafeModuloInt(_637_v85, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wxvn")).Cardinality())))
							}
						}
						return _coll32.ToSet()
					}())
					var _source9 D0 = Companion_Default___.Fm20((_636_v86).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_636_v86).Cardinality()))).Uint32()).(_dafny.Set), !((_this.F6()).Cmp((_this).F5()) == 0), false, globalState)
					_ = _source9
					if _source9.Is_DC1() {
						var _638___mcc_h3 bool = _source9.Get_().(D0_DC1).Cf1
						_ = _638___mcc_h3
						var _639___mcc_h4 bool = _source9.Get_().(D0_DC1).Cf2
						_ = _639___mcc_h4
						var _640_cf2 bool = _639___mcc_h4
						_ = _640_cf2
						var _641_cf1 bool = _638___mcc_h3
						_ = _641_cf1
						var _642_v87 _dafny.Sequence
						_ = _642_v87
						_642_v87 = _dafny.UnicodeSeqOfUtf8Bytes("ekgniqmts")
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
						_ = _index113
						(_590_v42).ArraySet1CodePoint((_642_v87).Select((Companion_Default___.SafeIndex((_this.F6()).Plus(_dafny.IntOfInt64(862)), _dafny.IntOfUint32((_642_v87).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index113).Int())
						_640_cf2 = _641_cf1
						var _643_v88 _dafny.Array
						_ = _643_v88
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_15
						var _nw121 _dafny.Array
						_ = _nw121
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw121 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) _dafny.MultiSet = (func(_644_p0 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
								return func(_645_i4 _dafny.Int) _dafny.MultiSet {
									return (_dafny.MultiSetOf(_dafny.IntOfInt64(729))).Union(_dafny.MultiSetOf(_644_p0, (_this).F5(), _dafny.IntOfInt64(180)))
								}
							})(p0)
							_ = _init15
							var _element0_15 = _init15(_dafny.Zero)
							_ = _element0_15
							_nw121 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
							(_nw121).ArraySet1(_element0_15, 0)
							var _nativeLen0_15 = (_len0_15).Int()
							_ = _nativeLen0_15
							for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
								(_nw121).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
							}
						}
						_643_v88 = _nw121
						_643_v88 = _643_v88
						var _646_v89 _dafny.Map
						_ = _646_v89
						_646_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_641_cf1, p0)
						var _647_v90 _dafny.Set
						_ = _647_v90
						_647_v90 = _dafny.SetOf(_641_cf1)
						var _648_v91 _dafny.Sequence
						_ = _648_v91
						_648_v91 = _dafny.SeqOf(_647_v90, _647_v90, _647_v90, _647_v90)
						var _649_v92 _dafny.Map
						_ = _649_v92
						_649_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(473), _this.F6())
						var _650_v93 _dafny.MultiSet
						_ = _650_v93
						_650_v93 = _dafny.MultiSetOf(_this.F6())
						var _651_v94 _dafny.Map
						_ = _651_v94
						_651_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_650_v93, _541_v0)
						var _652_v95 _dafny.Set
						_ = _652_v95
						_652_v95 = _dafny.SetOf((_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))
						var _653_v96 _dafny.Sequence
						_ = _653_v96
						_653_v96 = _dafny.SeqOf(_this.F6(), _dafny.IntOfUint32((_642_v87).Cardinality()))
						var _654_v98 _dafny.MultiSet
						_ = _654_v98
						_654_v98 = _dafny.MultiSetOf(_652_v95, func() _dafny.Set {
							var _coll33 = _dafny.NewBuilder()
							_ = _coll33
							for _iter38 := _dafny.Iterate((_652_v95).Elements()); ; {
								_compr_33, _ok38 := _iter38()
								if !_ok38 {
									break
								}
								var _655_v97 _dafny.CodePoint
								_655_v97 = interface{}(_compr_33).(_dafny.CodePoint)
								if (_652_v95).Contains(_655_v97) {
									_coll33.Add(_655_v97)
								}
							}
							return _coll33.ToSet()
						}())
						var _656_v99 _dafny.Array
						_ = _656_v99
						var _nwElement0_25 _dafny.Int = (_dafny.Zero).Minus((_this).Fm2((_dafny.SetOf(_641_cf1, _641_cf1)).Cardinality(), _dafny.SetOf(_541_v0), globalState))
						_ = _nwElement0_25
						var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(28))
						_ = _nw122
						(_nw122).ArraySet1(_nwElement0_25, 0)
						(_nw122).ArraySet1((_this).F5(), 1)
						(_nw122).ArraySet1((func() _dafny.Int {
							if _640_cf2 {
								return (_this).F5()
							}
							return (_this).F5()
						})(), 2)
						(_nw122).ArraySet1((_646_v89).Cardinality(), 3)
						(_nw122).ArraySet1(_this.F6(), 4)
						(_nw122).ArraySet1(_this.F6(), 5)
						(_nw122).ArraySet1((_this).F5(), 6)
						(_nw122).ArraySet1(p0, 7)
						(_nw122).ArraySet1(p0, 8)
						(_nw122).ArraySet1((_this).F5(), 9)
						(_nw122).ArraySet1((_this).F5(), 10)
						(_nw122).ArraySet1(_dafny.IntOfInt64(-309), 11)
						(_nw122).ArraySet1(_this.F6(), 12)
						(_nw122).ArraySet1((_this).Fm2(_this.F6(), _647_v90, globalState), 13)
						(_nw122).ArraySet1(_dafny.IntOfUint32((_642_v87).Cardinality()), 14)
						(_nw122).ArraySet1((func() _dafny.Int {
							if !(_640_cf2) {
								return _dafny.IntOfUint32((_648_v91).Cardinality())
							}
							return _this.F6()
						})(), 15)
						(_nw122).ArraySet1((_dafny.Zero).Minus(p0), 16)
						(_nw122).ArraySet1(_this.F6(), 17)
						(_nw122).ArraySet1(((_649_v92).Cardinality()).Times((_651_v94).Cardinality()), 18)
						(_nw122).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_594_v46, _594_v46)).Cardinality()), 19)
						(_nw122).ArraySet1((_this).F5(), 20)
						(_nw122).ArraySet1((p0).Times(_this.F6()), 21)
						(_nw122).ArraySet1((_dafny.Zero).Minus((_this).Fm2((_652_v95).Cardinality(), _647_v90, globalState)), 22)
						(_nw122).ArraySet1((_dafny.Zero).Minus((_this).F5()), 23)
						(_nw122).ArraySet1((_653_v96).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_653_v96).Cardinality()))).Uint32()).(_dafny.Int), 24)
						(_nw122).ArraySet1((_this).F5(), 25)
						(_nw122).ArraySet1(_this.F6(), 26)
						(_nw122).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F5(), (_dafny.Zero).Minus((func() _dafny.Int {
							if (_654_v98).Contains(_652_v95) {
								return (_654_v98).Multiplicity(_652_v95)
							}
							return (_this).F5()
						})())), 27)
						_656_v99 = _nw122
						var _657_v100 D4
						_ = _657_v100
						_657_v100 = Companion_D4_.Create_DC13_(!(_541_v0), (_649_v92).Cardinality(), _640_cf2)
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
						_ = _index114
						var _rhs128 _dafny.Array = _656_v99
						_ = _rhs128
						var _rhs129 _dafny.CodePoint = (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int())
						_ = _rhs129
						var _rhs130 _dafny.Array = _656_v99
						_ = _rhs130
						var _rhs131 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_653_v96, _653_v96)).Cardinality())
						_ = _rhs131
						var _rhs132 bool = (_657_v100).Dtor_cf17()
						_ = _rhs132
						var _lhs74 _dafny.Array = _590_v42
						_ = _lhs74
						var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))
						_ = _lhs75
						var _lhs76 *GlobalState = globalState
						_ = _lhs76
						_656_v99 = _rhs128
						(_lhs74).ArraySet1CodePoint(_rhs129, (_lhs75).Int())
						_656_v99 = _rhs130
						_lhs76.F2 = _rhs131
						_541_v0 = _rhs132
					} else if _source9.Is_DC2() {
						var _658_v101 _dafny.Set
						_ = _658_v101
						_658_v101 = _dafny.SetOf(_541_v0)
						_541_v0 = ((_658_v101).IsSubsetOf(_658_v101)) == (_541_v0)
						var _659_v102 _dafny.Array
						_ = _659_v102
						var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
						_ = _nw123
						_659_v102 = _nw123
						var _660_v103 _dafny.Array
						_ = _660_v103
						var _nw124 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw124
						_660_v103 = _nw124
						var _661_v104 _dafny.Sequence
						_ = _661_v104
						_661_v104 = _dafny.SeqOf(_660_v103)
						var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_659_v102), 0))
						_ = _index115
						(_659_v102).ArraySet1((_661_v104).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(390), _dafny.IntOfUint32((_661_v104).Cardinality()))).Uint32()).(_dafny.Array), (_index115).Int())
						var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_659_v102), 0))
						_ = _index116
						(_659_v102).ArraySet1(_660_v103, (_index116).Int())
						var _662_v105 _dafny.MultiSet
						_ = _662_v105
						_662_v105 = _dafny.MultiSetOf(_658_v101)
						var _663_v106 *C1
						_ = _663_v106
						var _nw125 *C1 = New_C1_()
						_ = _nw125
						_nw125.Ctor__(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}((func(_664_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_665_i5 _dafny.Int) _dafny.CodePoint {
								return _664_v43
							}
						})(_591_v43)))).Cardinality()), ((_662_v105).Update(Companion_Default___.Fm21(p0, p0, _541_v0, globalState), Companion_Default___.Abs(_this.F6()))).Cardinality())
						_663_v106 = _nw125
						(_this).F6_set_((_dafny.Zero).Minus(p0))
					} else if _source9.Is_DC3() {
						var _666___mcc_h5 bool = _source9.Get_().(D0_DC3).Cf3
						_ = _666___mcc_h5
						var _667_cf3 bool = _666___mcc_h5
						_ = _667_cf3
						var _668_v107 _dafny.Sequence
						_ = _668_v107
						_668_v107 = _dafny.UnicodeSeqOfUtf8Bytes("lsid")
						var _669_v108 _dafny.Sequence
						_ = _669_v108
						_669_v108 = _dafny.SeqOf(_668_v107)
						var _670_v109 _dafny.Map
						_ = _670_v109
						_670_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _669_v108)
						_670_v109 = (_670_v109).Update(p0, _669_v108)
						_668_v107 = _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if _667_cf3 {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-803))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg36 _dafny.Int) interface{} {
										return coer36(arg36)
									}
								}(func(_671_i6 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('r')
								}))
							}
							return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj"))).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int())), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj"))).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))).Cardinality()))).Uint32(), _591_v43)
						})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if _667_cf3 {
								return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-803))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg37 _dafny.Int) interface{} {
										return coer37(arg37)
									}
								}(func(_672_i6 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('r')
								}))
							}
							return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj"))).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int())), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj")), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_668_v107, _dafny.UnicodeSeqOfUtf8Bytes("geuucj"))).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))).Cardinality()))).Uint32(), _591_v43)
						})()).Cardinality()))).Uint32(), _591_v43)
						var _673_v110 _dafny.Sequence
						_ = _673_v110
						_673_v110 = _dafny.SeqOf(p0, _this.F6())
						var _674_v111 *C0
						_ = _674_v111
						var _nw126 *C0 = New_C0_()
						_ = _nw126
						_nw126.Ctor__((_dafny.SetOf(_673_v110)).Cardinality())
						_674_v111 = _nw126
						var _675_v112 _dafny.Array
						_ = _675_v112
						var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
						_ = _nw127
						_675_v112 = _nw127
						var _676_v113 _dafny.Map
						_ = _676_v113
						_676_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), (_594_v46).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32()).(bool))
						var _677_v114 D5
						_ = _677_v114
						_677_v114 = Companion_D5_.Create_DC17_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F6(), _dafny.IntOfUint32((_dafny.SeqOf(_667_cf3)).Cardinality()), _dafny.IntOfInt64(288))).Cardinality()), _675_v112, _676_v113)
						var _678_v115 _dafny.MultiSet
						_ = _678_v115
						_678_v115 = _dafny.MultiSetOf(Companion_D5_.Create_DC17_(_dafny.IntOfUint32((_594_v46).Cardinality()), _675_v112, _676_v113), _677_v114)
						var _679_v116 _dafny.Array
						_ = _679_v116
						var _nw128 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
						_ = _nw128
						_679_v116 = _nw128
						var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_679_v116), 0))
						_ = _index117
						(_679_v116).ArraySet1(_667_cf3, (_index117).Int())
						var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_679_v116), 0))
						_ = _index118
						(_679_v116).ArraySet1(_667_cf3, (_index118).Int())
						var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_679_v116), 0))
						_ = _index119
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_679_v116), 0))
						_ = _index120
						var _rhs133 _dafny.MultiSet = _678_v115
						_ = _rhs133
						var _rhs134 bool = _541_v0
						_ = _rhs134
						var _rhs135 bool = !(_541_v0)
						_ = _rhs135
						var _lhs77 _dafny.Array = _679_v116
						_ = _lhs77
						var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(571), _dafny.ArrayLen((_679_v116), 0))
						_ = _lhs78
						var _lhs79 _dafny.Array = _679_v116
						_ = _lhs79
						var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_679_v116), 0))
						_ = _lhs80
						_678_v115 = _rhs133
						(_lhs77).ArraySet1(_rhs134, (_lhs78).Int())
						(_lhs79).ArraySet1(_rhs135, (_lhs80).Int())
					} else {
						var _680___mcc_h6 bool = _source9.Get_().(D0_DC0).Cf0
						_ = _680___mcc_h6
						var _681_cf0 bool = _680___mcc_h6
						_ = _681_cf0
						var _682_v117 _dafny.Set
						_ = _682_v117
						_682_v117 = _dafny.SetOf(p0)
						_682_v117 = ((_682_v117).Intersection(_682_v117)).Intersection(_682_v117)
						(globalState).F2 = p0
						var _683_v118 *C1
						_ = _683_v118
						var _nw129 *C1 = New_C1_()
						_ = _nw129
						_nw129.Ctor__((_this).F5(), (_this).F5())
						_683_v118 = _nw129
						var _684_v119 T0
						_ = _684_v119
						var _nw130 *C1 = New_C1_()
						_ = _nw130
						_nw130.Ctor__(p0, (_dafny.Zero).Minus(_this.F6()))
						_684_v119 = _nw130
						var _685_v120 T0
						_ = _685_v120
						_685_v120 = _684_v119
						_684_v119 = (_685_v120)
					}
					if _541_v0 {
						_541_v0 = _541_v0
						var _686_v121 _dafny.Sequence
						_ = _686_v121
						_686_v121 = _dafny.SeqOf(_this.F6())
						var _687_v123 _dafny.Map
						_ = _687_v123
						_687_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm22(func() _dafny.Set {
							var _coll34 = _dafny.NewBuilder()
							_ = _coll34
							for _iter39 := _dafny.Iterate((_686_v121).Elements()); ; {
								_compr_34, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _688_v122 _dafny.Int
								_688_v122 = interface{}(_compr_34).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_686_v121, _688_v122) {
									_coll34.Add((_688_v122).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(562))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
										return func(arg38 _dafny.Int) interface{} {
											return coer38(arg38)
										}
									}(func(_689_i7 _dafny.Int) _dafny.Int {
										return (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(561))).Cardinality())
									}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vudi")).Cardinality()))).Cardinality()))
								}
							}
							return _coll34.ToSet()
						}(), _541_v0, globalState), _541_v0)
						var _690_v124 D7
						_ = _690_v124
						_690_v124 = Companion_D7_.Create_DC21_(_687_v123)
						var _691_v125 _dafny.Set
						_ = _691_v125
						_691_v125 = _dafny.SetOf(_541_v0)
						var _692_v126 _dafny.Map
						_ = _692_v126
						_692_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, (_this).Fm1(globalState))
						var _693_v127 D3
						_ = _693_v127
						_693_v127 = Companion_D3_.Create_DC8_(_692_v126)
						var _694_v129 _dafny.Sequence
						_ = _694_v129
						_694_v129 = _dafny.SeqOf(_693_v127)
						var _695_v131 _dafny.Sequence
						_ = _695_v131
						_695_v131 = _dafny.SeqOf(_687_v123)
						var _696_v132 _dafny.Array
						_ = _696_v132
						var _nwElement0_26 _dafny.Map = (_687_v123).Merge(_687_v123)
						_ = _nwElement0_26
						var _nw131 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(28))
						_ = _nw131
						(_nw131).ArraySet1(_nwElement0_26, 0)
						(_nw131).ArraySet1((_690_v124).Dtor_cf37(), 1)
						(_nw131).ArraySet1(_687_v123, 2)
						(_nw131).ArraySet1(_687_v123, 3)
						(_nw131).ArraySet1(_687_v123, 4)
						(_nw131).ArraySet1((_687_v123).Merge((Companion_Default___.Fm23((_this).Fm2(p0, _691_v125, globalState), globalState)).Update(_693_v127, true)), 5)
						(_nw131).ArraySet1(_687_v123, 6)
						(_nw131).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v127, _541_v0)).Merge(_687_v123), 7)
						(_nw131).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v127, _541_v0), 8)
						(_nw131).ArraySet1((_687_v123).Merge(_687_v123), 9)
						(_nw131).ArraySet1(((func() _dafny.Map {
							var _coll35 = _dafny.NewMapBuilder()
							_ = _coll35
							for _iter40 := _dafny.Iterate((_694_v129).Elements()); ; {
								_compr_35, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _697_v128 D3
								_697_v128 = interface{}(_compr_35).(D3)
								if _dafny.Companion_Sequence_.Contains(_694_v129, _697_v128) {
									_coll35.Add(_697_v128, _541_v0)
								}
							}
							return _coll35.ToMap()
						}()).Update(_693_v127, (_this).Fm1(globalState))).Merge(_687_v123), 10)
						(_nw131).ArraySet1(func() _dafny.Map {
							var _coll36 = _dafny.NewMapBuilder()
							_ = _coll36
							for _iter41 := _dafny.Iterate((_687_v123).Keys().Elements()); ; {
								_compr_36, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _698_v130 D3
								_698_v130 = interface{}(_compr_36).(D3)
								if (_687_v123).Contains(_698_v130) {
									_coll36.Add(_698_v130, _541_v0)
								}
							}
							return _coll36.ToMap()
						}(), 11)
						(_nw131).ArraySet1(_687_v123, 12)
						(_nw131).ArraySet1(_687_v123, 13)
						(_nw131).ArraySet1(_687_v123, 14)
						(_nw131).ArraySet1((func() _dafny.Map {
							if _541_v0 {
								return _687_v123
							}
							return _687_v123
						})(), 15)
						(_nw131).ArraySet1(_687_v123, 16)
						(_nw131).ArraySet1(_687_v123, 17)
						(_nw131).ArraySet1(_687_v123, 18)
						(_nw131).ArraySet1((_695_v131).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_695_v131).Cardinality()))).Uint32()).(_dafny.Map), 19)
						(_nw131).ArraySet1((_687_v123).Merge(_687_v123), 20)
						(_nw131).ArraySet1(_687_v123, 21)
						(_nw131).ArraySet1((_687_v123).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v127, _541_v0)), 22)
						(_nw131).ArraySet1((_687_v123).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v127, (_594_v46).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32()).(bool))), 23)
						(_nw131).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v127, !(_541_v0)), 24)
						(_nw131).ArraySet1(_687_v123, 25)
						(_nw131).ArraySet1(_687_v123, 26)
						(_nw131).ArraySet1(_687_v123, 27)
						_696_v132 = _nw131
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_696_v132), 0))
						_ = _index121
						(_696_v132).ArraySet1((Companion_Default___.Fm23(_this.F6(), globalState)).Merge(_687_v123), (_index121).Int())
						var _699_v133 _dafny.Set
						_ = _699_v133
						_699_v133 = _dafny.SetOf(_dafny.SeqOf((_this).F5(), (_dafny.Zero).Minus(p0), _this.F6()))
						var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_696_v132), 0))
						_ = _index122
						var _rhs136 bool = ((_699_v133).Union(_699_v133)).IsProperSubsetOf((_699_v133).Union(_699_v133))
						_ = _rhs136
						var _rhs137 bool = _541_v0
						_ = _rhs137
						var _rhs138 bool = ((_dafny.Zero).Minus(p0)).Cmp(p0) >= 0
						_ = _rhs138
						var _rhs139 _dafny.Int = (_dafny.IntOfUint32((_594_v46).Cardinality())).Times(((_this).F5()).Times((_this).F5()))
						_ = _rhs139
						var _rhs140 _dafny.Map = (_687_v123).Merge((Companion_Default___.Fm23((_dafny.Zero).Minus(_dafny.IntOfUint32((_594_v46).Cardinality())), globalState)).Merge((_687_v123).Update(_693_v127, _541_v0)))
						_ = _rhs140
						var _lhs81 *GlobalState = globalState
						_ = _lhs81
						var _lhs82 _dafny.Array = _696_v132
						_ = _lhs82
						var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(526), _dafny.ArrayLen((_696_v132), 0))
						_ = _lhs83
						_541_v0 = _rhs136
						_541_v0 = _rhs137
						_541_v0 = _rhs138
						_lhs81.F2 = _rhs139
						(_lhs82).ArraySet1(_rhs140, (_lhs83).Int())
						var _700_v134 _dafny.Array
						_ = _700_v134
						var _nwElement0_27 bool = _541_v0
						_ = _nwElement0_27
						var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(10))
						_ = _nw132
						(_nw132).ArraySet1(_nwElement0_27, 0)
						(_nw132).ArraySet1(_541_v0, 1)
						(_nw132).ArraySet1(_541_v0, 2)
						(_nw132).ArraySet1(_541_v0, 3)
						(_nw132).ArraySet1(_541_v0, 4)
						(_nw132).ArraySet1((p0).Cmp(_this.F6()) != 0, 5)
						(_nw132).ArraySet1(_541_v0, 6)
						(_nw132).ArraySet1(_541_v0, 7)
						(_nw132).ArraySet1(_541_v0, 8)
						(_nw132).ArraySet1(_541_v0, 9)
						_700_v134 = _nw132
						_700_v134 = _700_v134
						_541_v0 = !(_541_v0)
						var _701_v135 _dafny.Array
						_ = _701_v135
						var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
						_ = _nw133
						_701_v135 = _nw133
						var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_701_v135), 0))
						_ = _index123
						(_701_v135).ArraySet1((func() _dafny.Sequence {
							if _541_v0 {
								return _686_v121
							}
							return _686_v121
						})(), (_index123).Int())
						var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_701_v135), 0))
						_ = _index124
						(_701_v135).ArraySet1(_686_v121, (_index124).Int())
					} else {
						var _702_v136 _dafny.Array
						_ = _702_v136
						var _703_v137 bool
						_ = _703_v137
						var _704_v138 _dafny.MultiSet
						_ = _704_v138
						var _705_v139 _dafny.Int
						_ = _705_v139
						var _out48 _dafny.Array
						_ = _out48
						var _out49 bool
						_ = _out49
						var _out50 _dafny.MultiSet
						_ = _out50
						var _out51 _dafny.Int
						_ = _out51
						_out48, _out49, _out50, _out51 = Companion_Default___.M0(Companion_Default___.Fm18(_541_v0, _this.F6(), _this.F6(), true, globalState), globalState)
						_702_v136 = _out48
						_703_v137 = _out49
						_704_v138 = _out50
						_705_v139 = _out51
						var _706_v140 _dafny.Set
						_ = _706_v140
						_706_v140 = _dafny.SetOf(!(_541_v0))
						_705_v139 = (_this).Fm2((_this).Fm2(_this.F6(), _706_v140, globalState), _706_v140, globalState)
						var _707_v141 _dafny.Map
						_ = _707_v141
						_707_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _702_v136)
						var _708_v142 _dafny.Set
						_ = _708_v142
						_708_v142 = _dafny.SetOf(_dafny.IntOfInt64(616), _705_v139)
						var _709_v143 *C1
						_ = _709_v143
						var _nw134 *C1 = New_C1_()
						_ = _nw134
						_nw134.Ctor__(Companion_Default___.SafeDivisionInt(_this.F6(), (_707_v141).Cardinality()), (_708_v142).Cardinality())
						_709_v143 = _nw134
						var _rhs141 _dafny.CodePoint = _591_v43
						_ = _rhs141
						var _rhs142 *C1 = _709_v143
						_ = _rhs142
						var _rhs143 bool = _703_v137
						_ = _rhs143
						var _rhs144 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_705_v139), _705_v139)
						_ = _rhs144
						var _lhs84 *GlobalState = globalState
						_ = _lhs84
						_591_v43 = _rhs141
						_709_v143 = _rhs142
						_541_v0 = _rhs143
						_lhs84.F2 = _rhs144
						var _710_v144 _dafny.Sequence
						_ = _710_v144
						_710_v144 = _dafny.UnicodeSeqOfUtf8Bytes("bnmpaxw")
						var _711_v145 *C1
						_ = _711_v145
						var _nw135 *C1 = New_C1_()
						_ = _nw135
						_nw135.Ctor__((_dafny.Zero).Minus(_705_v139), (_dafny.Zero).Minus(((_this).F5()).Minus(_dafny.IntOfUint32((_710_v144).Cardinality()))))
						_711_v145 = _nw135
						var _712_v146 D0
						_ = _712_v146
						_712_v146 = Companion_D0_.Create_DC3_(_541_v0)
						var _rhs145 _dafny.Int = _this.F6()
						_ = _rhs145
						var _rhs146 D0 = _712_v146
						_ = _rhs146
						var _rhs147 bool = _703_v137
						_ = _rhs147
						var _rhs148 _dafny.Int = (_709_v143).Fm2((_this.F6()).Minus(_dafny.IntOfInt64(922)), _706_v140, globalState)
						_ = _rhs148
						var _rhs149 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_710_v144, _dafny.Companion_Sequence_.Update(_710_v144, (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_703_v137)).Cardinality(), _dafny.IntOfUint32((_710_v144).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int())))).Cardinality())).Times(p0))
						_ = _rhs149
						var _lhs85 *GlobalState = globalState
						_ = _lhs85
						var _lhs86 *GlobalState = globalState
						_ = _lhs86
						_lhs85.F2 = _rhs145
						_712_v146 = _rhs146
						_703_v137 = _rhs147
						_705_v139 = _rhs148
						_lhs86.F2 = _rhs149
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _713_v147 _dafny.Array
		_ = _713_v147
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_16
		var _nw136 _dafny.Array
		_ = _nw136
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw136 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) _dafny.Int = (func(_714_v0 bool) func(_dafny.Int) _dafny.Int {
				return func(_715_i8 _dafny.Int) _dafny.Int {
					return (_715_i8).Times((_dafny.SetOf(_714_v0, _714_v0, !(_714_v0))).Cardinality())
				}
			})(_541_v0)
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
		_713_v147 = _nw136
		var _716_v148 _dafny.Map
		_ = _716_v148
		_716_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v147, _this.F6())
		var _717_v149 _dafny.Map
		_ = _717_v149
		_717_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _716_v148)
		_717_v149 = (_717_v149).Update((func() _dafny.Int {
			if (_592_v44).Contains(_541_v0) {
				return (_592_v44).Multiplicity(_541_v0)
			}
			return (_this).F5()
		})(), _716_v148)
		if true {
			_591_v43 = _591_v43
			var _718_v150 _dafny.Set
			_ = _718_v150
			_718_v150 = _dafny.SetOf(_713_v147, _713_v147)
			_718_v150 = _718_v150
			var _719_v151 _dafny.Sequence
			_ = _719_v151
			_719_v151 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
			_ = _index125
			(_713_v147).ArraySet1(_dafny.IntOfUint32((_719_v151).Cardinality()), (_index125).Int())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
			_ = _index126
			(_713_v147).ArraySet1(_this.F6(), (_index126).Int())
			var _720_v152 D0
			_ = _720_v152
			_720_v152 = Companion_D0_.Create_DC2_()
			var _source10 D0 = _720_v152
			_ = _source10
			if _source10.Is_DC1() {
				var _721___mcc_h7 bool = _source10.Get_().(D0_DC1).Cf1
				_ = _721___mcc_h7
				var _722___mcc_h8 bool = _source10.Get_().(D0_DC1).Cf2
				_ = _722___mcc_h8
				var _723_cf2 bool = _722___mcc_h8
				_ = _723_cf2
				var _724_cf1 bool = _721___mcc_h7
				_ = _724_cf1
				var _725_v153 _dafny.Array
				_ = _725_v153
				var _nw137 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw137
				_725_v153 = _nw137
				var _726_v154 _dafny.Map
				_ = _726_v154
				_726_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_725_v153, (_this).F5())
				_726_v154 = (_726_v154).Update(_725_v153, _this.F6())
				(_this).F6_set_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(234), (_this).F5()))
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
				_ = _index127
				(_713_v147).ArraySet1(_dafny.IntOfInt64(236), (_index127).Int())
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_725_v153), 0))
				_ = _index128
				(_725_v153).ArraySet1(_541_v0, (_index128).Int())
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_725_v153), 0))
				_ = _index129
				(_725_v153).ArraySet1(_724_cf1, (_index129).Int())
			} else if _source10.Is_DC2() {
				_541_v0 = (_541_v0) || (_541_v0)
				var _727_v155 _dafny.Map
				_ = _727_v155
				_727_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _541_v0)
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
				_ = _index130
				(_713_v147).ArraySet1((((_727_v155).Merge(Companion_Default___.Fm10((_this).F5(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), p0), _541_v0, p0, globalState))).Cardinality()).Plus((p0).Minus(p0)), (_index130).Int())
				var _728_v156 _dafny.Array
				_ = _728_v156
				var _nw138 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(6))
				_ = _nw138
				_728_v156 = _nw138
				var _729_v157 D3
				_ = _729_v157
				_729_v157 = Companion_D3_.Create_DC8_(_727_v155)
				var _730_v158 _dafny.Map
				_ = _730_v158
				_730_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_729_v157, _541_v0)
				var _731_v159 _dafny.Map
				_ = _731_v159
				_731_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC21_(_730_v158), _728_v156)
				var _732_v160 _dafny.Sequence
				_ = _732_v160
				_732_v160 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _541_v0))
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
				_ = _index131
				var _rhs150 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int)), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))
				_ = _rhs150
				var _rhs151 _dafny.Array = (func() _dafny.Array {
					if (_731_v159).Contains(Companion_Default___.Fm24(_732_v160, _dafny.IntOfInt64(-597), _541_v0, globalState)) {
						return (_731_v159).Get(Companion_Default___.Fm24(_732_v160, _dafny.IntOfInt64(-597), _541_v0, globalState)).(_dafny.Array)
					}
					return _728_v156
				})()
				_ = _rhs151
				var _lhs87 _dafny.Array = _713_v147
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
				_ = _lhs88
				(_lhs87).ArraySet1(_rhs150, (_lhs88).Int())
				_728_v156 = _rhs151
				var _733_v161 _dafny.Array
				_ = _733_v161
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw139
				_733_v161 = _nw139
				var _734_v162 _dafny.Array
				_ = _734_v162
				var _nwElement0_28 bool = _541_v0
				_ = _nwElement0_28
				var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(17))
				_ = _nw140
				(_nw140).ArraySet1(_nwElement0_28, 0)
				(_nw140).ArraySet1(_541_v0, 1)
				(_nw140).ArraySet1(_541_v0, 2)
				(_nw140).ArraySet1(_541_v0, 3)
				(_nw140).ArraySet1(_541_v0, 4)
				(_nw140).ArraySet1(false, 5)
				(_nw140).ArraySet1(!(false), 6)
				(_nw140).ArraySet1(_541_v0, 7)
				(_nw140).ArraySet1(_541_v0, 8)
				(_nw140).ArraySet1(true, 9)
				(_nw140).ArraySet1(_541_v0, 10)
				(_nw140).ArraySet1(_541_v0, 11)
				(_nw140).ArraySet1(!(_541_v0), 12)
				(_nw140).ArraySet1(_541_v0, 13)
				(_nw140).ArraySet1(false, 14)
				(_nw140).ArraySet1(false, 15)
				(_nw140).ArraySet1(false, 16)
				_734_v162 = _nw140
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_733_v161), 0))
				_ = _index132
				(_733_v161).ArraySet1(_734_v162, (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_733_v161), 0))
				_ = _index133
				var _rhs152 bool = ((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))) < 0
				_ = _rhs152
				var _rhs153 _dafny.Array = _734_v162
				_ = _rhs153
				var _rhs154 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_735_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_736_i9 _dafny.Int) _dafny.CodePoint {
						return _735_v43
					}
				})(_591_v43)))
				_ = _rhs154
				var _lhs89 _dafny.Array = _733_v161
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_733_v161), 0))
				_ = _lhs90
				_541_v0 = _rhs152
				(_lhs89).ArraySet1(_rhs153, (_lhs90).Int())
				_719_v151 = _rhs154
			} else if _source10.Is_DC3() {
				var _737___mcc_h9 bool = _source10.Get_().(D0_DC3).Cf3
				_ = _737___mcc_h9
				var _738_cf3 bool = _737___mcc_h9
				_ = _738_cf3
				var _739_v163 _dafny.Set
				_ = _739_v163
				_739_v163 = _dafny.SetOf(_dafny.CodePoint('h'), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))
				var _740_v164 _dafny.MultiSet
				_ = _740_v164
				_740_v164 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_719_v151, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.IntOfUint32((_719_v151).Cardinality()))).Uint32(), (_590_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(508), _dafny.ArrayLen((_590_v42), 0))).Int()))).Cardinality()))
				var _rhs155 _dafny.Set = _739_v163
				_ = _rhs155
				var _rhs156 _dafny.Int = Companion_Default___.SafeDivisionInt(((_740_v164).Cardinality()).Minus(_this.F6()), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))
				_ = _rhs156
				var _lhs91 *GlobalState = globalState
				_ = _lhs91
				_739_v163 = _rhs155
				_lhs91.F2 = _rhs156
				var _741_v165 _dafny.Array
				_ = _741_v165
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_17
				var _nw141 _dafny.Array
				_ = _nw141
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw141 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = func(_742_i10 _dafny.Int) _dafny.Int {
						return (_742_i10).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jsc")).Cardinality()))
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw141 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw141).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw141).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_741_v165 = _nw141
				var _743_v166 _dafny.Sequence
				_ = _743_v166
				_743_v166 = _dafny.SeqOf(_741_v165, _741_v165, _741_v165)
				_741_v165 = (_743_v166).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_743_v166).Cardinality()))).Uint32()).(_dafny.Array)
				var _744_v167 D0
				_ = _744_v167
				_744_v167 = Companion_D0_.Create_DC1_(_738_cf3, _541_v0)
				var _745_v168 _dafny.Sequence
				_ = _745_v168
				_745_v168 = _dafny.SeqOf(_740_v164)
				var _746_v169 _dafny.Set
				_ = _746_v169
				_746_v169 = _dafny.SetOf(!(_541_v0))
				var _747_v170 _dafny.Map
				_ = _747_v170
				_747_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))
				var _748_v171 _dafny.Map
				_ = _748_v171
				_748_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(p0, _746_v169, globalState), (_747_v170).Update(_738_cf3, p0))
				var _749_v172 _dafny.Array
				_ = _749_v172
				var _nwElement0_29 D0 = Companion_D0_.Create_DC1_(false, _541_v0)
				_ = _nwElement0_29
				var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(8))
				_ = _nw142
				(_nw142).ArraySet1(_nwElement0_29, 0)
				(_nw142).ArraySet1(Companion_D0_.Create_DC1_(_541_v0, _541_v0), 1)
				(_nw142).ArraySet1(Companion_D0_.Create_DC1_((_this).Fm8(_dafny.UnicodeSeqOfUtf8Bytes("yotfvv"), _738_cf3, _541_v0, _541_v0, globalState), _541_v0), 2)
				(_nw142).ArraySet1(_744_v167, 3)
				(_nw142).ArraySet1(Companion_D0_.Create_DC1_(_738_cf3, _738_cf3), 4)
				(_nw142).ArraySet1(Companion_Default___.Fm25(_738_cf3, (_this).F5(), _745_v168, _748_v171, globalState), 5)
				(_nw142).ArraySet1(_744_v167, 6)
				(_nw142).ArraySet1(_744_v167, 7)
				_749_v172 = _nw142
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_749_v172), 0))
				_ = _index134
				(_749_v172).ArraySet1(Companion_Default___.Fm25(_738_cf3, _this.F6(), _745_v168, _748_v171, globalState), (_index134).Int())
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_749_v172), 0))
				_ = _index135
				(_749_v172).ArraySet1(Companion_Default___.Fm25((func() bool {
					if _541_v0 {
						return _738_cf3
					}
					return true
				})(), _this.F6(), _dafny.Companion_Sequence_.Concatenate(_745_v168, _745_v168), _748_v171, globalState), (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))
				_ = _index136
				(_713_v147).ArraySet1((_this).F5(), (_index136).Int())
			} else {
				var _750___mcc_h10 bool = _source10.Get_().(D0_DC0).Cf0
				_ = _750___mcc_h10
				var _751_cf0 bool = _750___mcc_h10
				_ = _751_cf0
				var _752_v173 D0
				_ = _752_v173
				_752_v173 = Companion_D0_.Create_DC3_(_541_v0)
				var _753_v174 _dafny.MultiSet
				_ = _753_v174
				_753_v174 = _dafny.MultiSetOf(_591_v43, _dafny.CodePoint('s'), _591_v43)
				var _754_v175 _dafny.Map
				_ = _754_v175
				_754_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_752_v173).Dtor_cf3(), (_753_v174).Cardinality())
				var _755_v176 _dafny.Sequence
				_ = _755_v176
				_755_v176 = _dafny.SeqOf(_754_v175)
				var _756_v177 _dafny.Map
				_ = _756_v177
				_756_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_755_v176).Cardinality()), _541_v0)
				var _757_v178 _dafny.Map
				_ = _757_v178
				_757_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_756_v177).Cardinality(), _751_cf0)
				var _758_v179 *C0
				_ = _758_v179
				var _nw143 *C0 = New_C0_()
				_ = _nw143
				_nw143.Ctor__((_756_v177).Cardinality())
				_758_v179 = _nw143
				var _759_v182 D1
				_ = _759_v182
				_759_v182 = Companion_D1_.Create_DC5_()
				var _760_v183 _dafny.Map
				_ = _760_v183
				_760_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v175, (_this).Fm9(_759_v182, _541_v0, globalState))
				var _761_v184 _dafny.Map
				_ = _761_v184
				_761_v184 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_594_v46).Cardinality()), _754_v175)
				var _762_v185 _dafny.Map
				_ = _762_v185
				_762_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_751_cf0, _dafny.IntOfInt64(352)), p0)
				var _763_v186 _dafny.Map
				_ = _763_v186
				_763_v186 = _762_v185
				var _764_v190 _dafny.Map
				_ = _764_v190
				_764_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v175, _719_v151)
				var _765_v191 _dafny.Array
				_ = _765_v191
				var _nwElement0_30 _dafny.Map = func() _dafny.Map {
					var _coll37 = _dafny.NewMapBuilder()
					_ = _coll37
					for _iter42 := _dafny.Iterate((func() _dafny.Map {
						var _coll38 = _dafny.NewMapBuilder()
						_ = _coll38
						for _iter43 := _dafny.Iterate((_760_v183).Keys().Elements()); ; {
							_compr_38, _ok43 := _iter43()
							if !_ok43 {
								break
							}
							var _766_v181 _dafny.Map
							_766_v181 = interface{}(_compr_38).(_dafny.Map)
							if (_760_v183).Contains(_766_v181) {
								_coll38.Add(_766_v181, (_592_v44).Cardinality())
							}
						}
						return _coll38.ToMap()
					}()).Keys().Elements()); ; {
						_compr_37, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _767_v180 _dafny.Map
						_767_v180 = interface{}(_compr_37).(_dafny.Map)
						if (func() _dafny.Map {
							var _coll39 = _dafny.NewMapBuilder()
							_ = _coll39
							for _iter44 := _dafny.Iterate((_760_v183).Keys().Elements()); ; {
								_compr_39, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _766_v181 _dafny.Map
								_766_v181 = interface{}(_compr_39).(_dafny.Map)
								if (_760_v183).Contains(_766_v181) {
									_coll39.Add(_766_v181, (_592_v44).Cardinality())
								}
							}
							return _coll39.ToMap()
						}()).Contains(_767_v180) {
							_coll37.Add(_767_v180, (_758_v179).F8())
						}
					}
					return _coll37.ToMap()
				}()
				_ = _nwElement0_30
				var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(14))
				_ = _nw144
				(_nw144).ArraySet1(_nwElement0_30, 0)
				(_nw144).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					if (_761_v184).Contains((_758_v179).F8()) {
						return (_761_v184).Get((_758_v179).F8()).(_dafny.Map)
					}
					return _754_v175
				})(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _541_v0)).Cardinality()), 1)
				(_nw144).ArraySet1(_762_v185, 2)
				(_nw144).ArraySet1((_763_v186), 3)
				(_nw144).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v175, p0), 4)
				(_nw144).ArraySet1((func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter45 := _dafny.Iterate((_760_v183).Keys().Elements()); ; {
						_compr_40, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _768_v187 _dafny.Map
						_768_v187 = interface{}(_compr_40).(_dafny.Map)
						if (_760_v183).Contains(_768_v187) {
							_coll40.Add(_768_v187, _dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Map {
								var _coll41 = _dafny.NewMapBuilder()
								_ = _coll41
								for _iter46 := _dafny.Iterate((_dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _751_cf0)))).Elements()); ; {
									_compr_41, _ok46 := _iter46()
									if !_ok46 {
										break
									}
									var _769_v188 D3
									_769_v188 = interface{}(_compr_41).(D3)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _751_cf0))), _769_v188) {
										_coll41.Add(_769_v188, (_758_v179).F8())
									}
								}
								return _coll41.ToMap()
							}())).Cardinality()))
						}
					}
					return _coll40.ToMap()
				}()).Merge(_762_v185), 5)
				(_nw144).ArraySet1(_762_v185, 6)
				(_nw144).ArraySet1(_762_v185, 7)
				(_nw144).ArraySet1(_762_v185, 8)
				(_nw144).ArraySet1((_762_v185).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v175, (_758_v179).F8())), 9)
				(_nw144).ArraySet1(func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter47 := _dafny.Iterate((_764_v190).Keys().Elements()); ; {
						_compr_42, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _770_v189 _dafny.Map
						_770_v189 = interface{}(_compr_42).(_dafny.Map)
						if (_764_v190).Contains(_770_v189) {
							_coll42.Add(_770_v189, (_this).F5())
						}
					}
					return _coll42.ToMap()
				}(), 10)
				(_nw144).ArraySet1((_762_v185).Update((_754_v175).Update(_541_v0, p0), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int)), 11)
				(_nw144).ArraySet1(_762_v185, 12)
				(_nw144).ArraySet1(_762_v185, 13)
				_765_v191 = _nw144
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_765_v191), 0))
				_ = _index137
				(_765_v191).ArraySet1(_762_v185, (_index137).Int())
				var _771_v192 D4
				_ = _771_v192
				_771_v192 = Companion_D4_.Create_DC15_(Companion_Default___.Fm27(_dafny.IntOfInt64(437), globalState))
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_765_v191), 0))
				_ = _index138
				var _rhs157 *C0 = _758_v179
				_ = _rhs157
				var _rhs158 _dafny.Map = Companion_Default___.Fm26(_541_v0, (_719_v151).Select((Companion_Default___.SafeIndex((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_719_v151).Cardinality()))).Uint32()).(_dafny.CodePoint), _771_v192, globalState)
				_ = _rhs158
				var _lhs92 _dafny.Array = _765_v191
				_ = _lhs92
				var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_765_v191), 0))
				_ = _lhs93
				_758_v179 = _rhs157
				(_lhs92).ArraySet1(_rhs158, (_lhs93).Int())
				var _772_v193 *C0
				_ = _772_v193
				var _nw145 *C0 = New_C0_()
				_ = _nw145
				_nw145.Ctor__(p0)
				_772_v193 = _nw145
				_541_v0 = _541_v0
				var _773_v194 _dafny.Set
				_ = _773_v194
				_773_v194 = _dafny.SetOf(_751_cf0)
				var _774_v195 _dafny.Sequence
				_ = _774_v195
				_774_v195 = _dafny.SeqOf(p0)
				var _775_v196 _dafny.Set
				_ = _775_v196
				_775_v196 = _dafny.SetOf(_dafny.IntOfUint32((_774_v195).Cardinality()), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))
				var _776_v197 _dafny.Sequence
				_ = _776_v197
				_776_v197 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_775_v196)).Cardinality()), (_758_v179).F8())).Cardinality(), _dafny.IntOfInt64(-505), (_dafny.Zero).Minus((_772_v193).F8()), (_this).F5())
				var _777_v198 _dafny.Sequence
				_ = _777_v198
				_777_v198 = _dafny.SeqOf(_756_v177, _756_v177, _757_v178, (_757_v178).Update(_dafny.IntOfUint32((_774_v195).Cardinality()), _751_cf0))
				var _778_v199 D5
				_ = _778_v199
				_778_v199 = Companion_D5_.Create_DC18_(_751_cf0, p0, false)
				var _pat_let_tv7 = _751_cf0
				_ = _pat_let_tv7
				var _779_v200 _dafny.Map
				_ = _779_v200
				_779_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let9_0 D5) D5 {
					return func(_780_dt__update__tmp_h0 D5) D5 {
						return func(_pat_let10_0 bool) D5 {
							return func(_781_dt__update_hcf30_h0 bool) D5 {
								return Companion_D5_.Create_DC18_(_781_dt__update_hcf30_h0, (_780_dt__update__tmp_h0).Dtor_cf31(), (_780_dt__update__tmp_h0).Dtor_cf32())
							}(_pat_let10_0)
						}(_pat_let_tv7)
					}(_pat_let9_0)
				}(_778_v199), (_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int))
				var _782_v201 _dafny.Array
				_ = _782_v201
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw146
				_782_v201 = _nw146
				var _783_v202 _dafny.Array
				_ = _783_v202
				var _nwElement0_31 _dafny.Int = (_this).F5()
				_ = _nwElement0_31
				var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(23))
				_ = _nw147
				(_nw147).ArraySet1(_nwElement0_31, 0)
				(_nw147).ArraySet1(_this.F6(), 1)
				(_nw147).ArraySet1((_this).Fm2(p0, _773_v194, globalState), 2)
				(_nw147).ArraySet1((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int), 3)
				(_nw147).ArraySet1(_dafny.IntOfInt64(195), 4)
				(_nw147).ArraySet1((_this).F5(), 5)
				(_nw147).ArraySet1(_this.F6(), 6)
				(_nw147).ArraySet1((_this).F5(), 7)
				(_nw147).ArraySet1(_dafny.IntOfUint32((_777_v198).Cardinality()), 8)
				(_nw147).ArraySet1(_dafny.IntOfUint32((_719_v151).Cardinality()), 9)
				(_nw147).ArraySet1((_772_v193).F8(), 10)
				(_nw147).ArraySet1((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int), 11)
				(_nw147).ArraySet1(_this.F6(), 12)
				(_nw147).ArraySet1(_this.F6(), 13)
				(_nw147).ArraySet1(((_779_v200).Update(_778_v199, _this.F6())).Cardinality(), 14)
				(_nw147).ArraySet1((_this).F5(), 15)
				(_nw147).ArraySet1((_this).F5(), 16)
				(_nw147).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_591_v43, _782_v201)).Cardinality(), 17)
				(_nw147).ArraySet1((_757_v178).Cardinality(), 18)
				(_nw147).ArraySet1((_713_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_713_v147), 0))).Int()).(_dafny.Int), 19)
				(_nw147).ArraySet1(_dafny.IntOfUint32((_776_v197).Cardinality()), 20)
				(_nw147).ArraySet1((_dafny.Zero).Minus((_772_v193).F8()), 21)
				(_nw147).ArraySet1(_dafny.IntOfInt64(79), 22)
				_783_v202 = _nw147
				var _784_v203 D5
				_ = _784_v203
				_784_v203 = Companion_D5_.Create_DC16_(_783_v202)
				var _pat_let_tv8 = _783_v202
				_ = _pat_let_tv8
				_784_v203 = func(_pat_let11_0 D5) D5 {
					return func(_785_dt__update__tmp_h1 D5) D5 {
						return func(_pat_let12_0 _dafny.Array) D5 {
							return func(_786_dt__update_hcf26_h0 _dafny.Array) D5 {
								return Companion_D5_.Create_DC16_(_786_dt__update_hcf26_h0)
							}(_pat_let12_0)
						}(_pat_let_tv8)
					}(_pat_let11_0)
				}(_784_v203)
			}
			var _787_v204 _dafny.Array
			_ = _787_v204
			var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw148
			_787_v204 = _nw148
			_787_v204 = _787_v204
		} else {
			_541_v0 = ((_592_v44).Union(_592_v44)).IsSubsetOf(_592_v44)
			var _788_v205 _dafny.Map
			_ = _788_v205
			_788_v205 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F6())
			_788_v205 = _788_v205
			if true {
				(globalState).F2 = _this.F6()
				var _789_v206 *C0
				_ = _789_v206
				var _nw149 *C0 = New_C0_()
				_ = _nw149
				_nw149.Ctor__(_this.F6())
				_789_v206 = _nw149
				var _790_v207 _dafny.Set
				_ = _790_v207
				_790_v207 = _dafny.SetOf(_789_v206)
				_541_v0 = !((_this.F6()).Cmp((_dafny.IntOfUint32((_594_v46).Cardinality())).Times((_790_v207).Cardinality())) >= 0)
				var _791_v208 _dafny.Set
				_ = _791_v208
				_791_v208 = _dafny.SetOf(_541_v0)
				var _792_v209 _dafny.Array
				_ = _792_v209
				var _nwElement0_32 _dafny.Sequence = _594_v46
				_ = _nwElement0_32
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_32, 0)
				(_nw150).ArraySet1(_594_v46, 1)
				(_nw150).ArraySet1(_594_v46, 2)
				(_nw150).ArraySet1((func() _dafny.Sequence {
					if _541_v0 {
						return _dafny.Companion_Sequence_.Update(_594_v46, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-17), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32(), _541_v0)
					}
					return _594_v46
				})(), 3)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_594_v46, _dafny.SeqOf(_541_v0, _541_v0)), 4)
				(_nw150).ArraySet1(_594_v46, 5)
				(_nw150).ArraySet1(_594_v46, 6)
				(_nw150).ArraySet1(_594_v46, 7)
				(_nw150).ArraySet1(_594_v46, 8)
				(_nw150).ArraySet1(_594_v46, 9)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_541_v0), _594_v46), 10)
				(_nw150).ArraySet1(_594_v46, 11)
				(_nw150).ArraySet1(_594_v46, 12)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Update(_594_v46, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32(), _541_v0), 13)
				(_nw150).ArraySet1(_594_v46, 14)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_594_v46, _594_v46), 15)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_594_v46, _594_v46), 16)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Update(_594_v46, (Companion_Default___.SafeIndex((_791_v208).Cardinality(), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32(), false), 17)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Update(_594_v46, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32(), _541_v0), 18)
				(_nw150).ArraySet1(_594_v46, 19)
				(_nw150).ArraySet1(_dafny.SeqOf(_541_v0), 20)
				(_nw150).ArraySet1(_dafny.SeqOf(_541_v0, _541_v0, _541_v0, _541_v0, true), 21)
				(_nw150).ArraySet1(_594_v46, 22)
				(_nw150).ArraySet1(_dafny.SeqOf(false, _541_v0, _541_v0), 23)
				(_nw150).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_594_v46).Select((Companion_Default___.SafeIndex((_789_v206).F8(), _dafny.IntOfUint32((_594_v46).Cardinality()))).Uint32()).(bool)), _dafny.SeqOf(_541_v0)), 24)
				_792_v209 = _nw150
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_792_v209), 0))
				_ = _index139
				(_792_v209).ArraySet1(_594_v46, (_index139).Int())
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_792_v209), 0))
				_ = _index140
				(_792_v209).ArraySet1(_594_v46, (_index140).Int())
				var _793_v210 _dafny.Array
				_ = _793_v210
				var _nw151 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw151
				_793_v210 = _nw151
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_793_v210), 0))
				_ = _index141
				(_793_v210).ArraySet1(!_dafny.Companion_Sequence_.Contains(_594_v46, _541_v0), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_793_v210), 0))
				_ = _index142
				var _rhs159 bool = _dafny.Companion_Sequence_.Contains((_792_v209).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(914), _dafny.ArrayLen((_792_v209), 0))).Int()).(_dafny.Sequence), !(_541_v0))
				_ = _rhs159
				var _rhs160 _dafny.Int = (_789_v206).F8()
				_ = _rhs160
				var _rhs161 bool = _541_v0
				_ = _rhs161
				var _rhs162 _dafny.Int = Companion_Default___.SafeModuloInt(((_789_v206).F8()).Plus((_789_v206).Fm4(_591_v43, _592_v44, _541_v0, globalState)), (_789_v206).F8())
				_ = _rhs162
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				var _lhs95 _dafny.Array = _793_v210
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(970), _dafny.ArrayLen((_793_v210), 0))
				_ = _lhs96
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				_541_v0 = _rhs159
				_lhs94.F2 = _rhs160
				(_lhs95).ArraySet1(_rhs161, (_lhs96).Int())
				_lhs97.F2 = _rhs162
				var _794_v211 _dafny.Sequence
				_ = _794_v211
				_794_v211 = _dafny.UnicodeSeqOfUtf8Bytes("nep")
				(_this).F6_set_((_dafny.Zero).Minus((_dafny.IntOfUint32((_794_v211).Cardinality())).Times(_dafny.IntOfInt64(449))))
			} else {
				_713_v147 = _713_v147
				var _795_v212 D7
				_ = _795_v212
				_795_v212 = Companion_D7_.Create_DC24_()
				_795_v212 = _795_v212
				_541_v0 = !(_541_v0)
				var _796_v213 *C1
				_ = _796_v213
				var _nw152 *C1 = New_C1_()
				_ = _nw152
				_nw152.Ctor__(p0, (_this).F5())
				_796_v213 = _nw152
				var _797_v214 _dafny.Map
				_ = _797_v214
				_797_v214 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_796_v213, _541_v0)
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_713_v147), 0))
				_ = _index143
				(_713_v147).ArraySet1((_797_v214).Cardinality(), (_index143).Int())
				var _798_v215 _dafny.Array
				_ = _798_v215
				var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw153
				_798_v215 = _nw153
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_798_v215), 0))
				_ = _index144
				(_798_v215).ArraySet1(_541_v0, (_index144).Int())
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_713_v147), 0))
				_ = _index145
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_798_v215), 0))
				_ = _index146
				var _rhs163 _dafny.Int = (_this).F5()
				_ = _rhs163
				var _rhs164 bool = _541_v0
				_ = _rhs164
				var _lhs98 _dafny.Array = _713_v147
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_713_v147), 0))
				_ = _lhs99
				var _lhs100 _dafny.Array = _798_v215
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(818), _dafny.ArrayLen((_798_v215), 0))
				_ = _lhs101
				(_lhs98).ArraySet1(_rhs163, (_lhs99).Int())
				(_lhs100).ArraySet1(_rhs164, (_lhs101).Int())
				var _799_v216 _dafny.Map
				_ = _799_v216
				_799_v216 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_798_v215, _592_v44)
				_799_v216 = (_799_v216).Update(_798_v215, _592_v44)
			}
			var _800_v217 _dafny.Array
			_ = _800_v217
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw154
			_800_v217 = _nw154
			var _801_v218 _dafny.Sequence
			_ = _801_v218
			_801_v218 = _dafny.UnicodeSeqOfUtf8Bytes("chcl")
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_800_v217), 0))
			_ = _index147
			(_800_v217).ArraySet1(_801_v218, (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_800_v217), 0))
			_ = _index148
			(_800_v217).ArraySet1(_801_v218, (_index148).Int())
			var _802_v219 _dafny.Map
			_ = _802_v219
			_802_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _this.F6())
			var _803_v220 *C0
			_ = _803_v220
			var _nw155 *C0 = New_C0_()
			_ = _nw155
			_nw155.Ctor__((func() _dafny.Int {
				if (_802_v219).Contains(_541_v0) {
					return (_802_v219).Get(_541_v0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_801_v218).Cardinality())
			})())
			_803_v220 = _nw155
		}
	}
}
func (_this *C2) M4(p0 bool, globalState *GlobalState) {
	{
		if (func() bool {
			if p0 {
				return p0
			}
			return p0
		})() {
			var _804_v0 _dafny.Set
			_ = _804_v0
			_804_v0 = _dafny.SetOf(_this.F6())
			var _805_v1 D1
			_ = _805_v1
			_805_v1 = Companion_D1_.Create_DC4_(_dafny.SetOf(_this.F6()))
			var _806_v2 _dafny.CodePoint
			_ = _806_v2
			_806_v2 = _dafny.CodePoint('a')
			var _807_v3 _dafny.Map
			_ = _807_v3
			_807_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F6())
			var _808_v4 _dafny.Map
			_ = _808_v4
			_808_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_806_v2, _807_v3)
			var _809_v5 _dafny.Array
			_ = _809_v5
			var _nwElement0_33 bool = p0
			_ = _nwElement0_33
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(21))
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_33, 0)
			(_nw156).ArraySet1((_804_v0).Contains((_this).F5()), 1)
			(_nw156).ArraySet1(p0, 2)
			(_nw156).ArraySet1(p0, 3)
			(_nw156).ArraySet1(false, 4)
			(_nw156).ArraySet1((!(p0)) && (p0), 5)
			(_nw156).ArraySet1((_804_v0).IsDisjointFrom((_805_v1).Dtor_cf4()), 6)
			(_nw156).ArraySet1(!(true) || (p0), 7)
			(_nw156).ArraySet1(p0, 8)
			(_nw156).ArraySet1(!(((_dafny.SetOf(false, false)).Cardinality()).Cmp(_this.F6()) >= 0), 9)
			(_nw156).ArraySet1(!((_808_v4).Contains(_806_v2)), 10)
			(_nw156).ArraySet1(true, 11)
			(_nw156).ArraySet1(p0, 12)
			(_nw156).ArraySet1(p0, 13)
			(_nw156).ArraySet1(true, 14)
			(_nw156).ArraySet1(p0, 15)
			(_nw156).ArraySet1(p0, 16)
			(_nw156).ArraySet1(p0, 17)
			(_nw156).ArraySet1(p0, 18)
			(_nw156).ArraySet1(p0, 19)
			(_nw156).ArraySet1(p0, 20)
			_809_v5 = _nw156
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))
			_ = _index149
			(_809_v5).ArraySet1(((_dafny.Zero).Minus((_this).F5())).Cmp((_this).F5()) <= 0, (_index149).Int())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))
			_ = _index150
			(_809_v5).ArraySet1((_this).Fm9(Companion_D1_.Create_DC5_(), ((_this).Fm1(globalState)) == (p0), globalState), (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))
			_ = _index151
			(_809_v5).ArraySet1(true, (_index151).Int())
			var _810_v6 _dafny.Set
			_ = _810_v6
			_810_v6 = _dafny.SetOf(p0, (_this.F6()).Cmp((_804_v0).Cardinality()) == 0, p0)
			_810_v6 = ((_dafny.SetOf((_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))).Difference(_810_v6)).Union(_810_v6)
			var _811_v7 _dafny.Sequence
			_ = _811_v7
			_811_v7 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F5())).Cardinality()))
			var _812_v8 _dafny.Map
			_ = _812_v8
			_812_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _dafny.IntOfUint32((_811_v7).Cardinality()))
			var _813_v9 _dafny.Sequence
			_ = _813_v9
			_813_v9 = _dafny.SeqOf(p0, (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))
			var _814_v10 _dafny.MultiSet
			_ = _814_v10
			_814_v10 = _dafny.MultiSetOf(p0, !((_813_v9).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_813_v9).Cardinality()))).Uint32()).(bool)), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))
			var _815_v11 _dafny.Array
			_ = _815_v11
			var _nwElement0_34 _dafny.Int = _this.F6()
			_ = _nwElement0_34
			var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(18))
			_ = _nw157
			(_nw157).ArraySet1(_nwElement0_34, 0)
			(_nw157).ArraySet1((_812_v8).Cardinality(), 1)
			(_nw157).ArraySet1((_this).F5(), 2)
			(_nw157).ArraySet1((_this).F5(), 3)
			(_nw157).ArraySet1((_this).F5(), 4)
			(_nw157).ArraySet1((_this).F5(), 5)
			(_nw157).ArraySet1(_this.F6(), 6)
			(_nw157).ArraySet1((_dafny.MultiSetFromSeq(_811_v7)).Cardinality(), 7)
			(_nw157).ArraySet1(_this.F6(), 8)
			(_nw157).ArraySet1(_dafny.IntOfInt64(199), 9)
			(_nw157).ArraySet1((_this).F5(), 10)
			(_nw157).ArraySet1((_this).F5(), 11)
			(_nw157).ArraySet1(_this.F6(), 12)
			(_nw157).ArraySet1((_this).F5(), 13)
			(_nw157).ArraySet1(_this.F6(), 14)
			(_nw157).ArraySet1((_this).F5(), 15)
			(_nw157).ArraySet1((_814_v10).Cardinality(), 16)
			(_nw157).ArraySet1((_this).F5(), 17)
			_815_v11 = _nw157
			var _816_v12 _dafny.Map
			_ = _816_v12
			_816_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), !((_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool)))
			var _817_v13 D5
			_ = _817_v13
			_817_v13 = Companion_D5_.Create_DC17_(_dafny.IntOfUint32((Companion_Default___.Fm11((_this).F5(), (_this).F5(), (_this).F5(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_811_v7).Cardinality())), globalState)).Cardinality()), _815_v11, _816_v12)
			var _818_v16 _dafny.Map
			_ = _818_v16
			_818_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_806_v2, (_this).F5())
			var _pat_let_tv9 = _815_v11
			_ = _pat_let_tv9
			var _819_v18 _dafny.Array
			_ = _819_v18
			var _nwElement0_35 D5 = _817_v13
			_ = _nwElement0_35
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(28))
			_ = _nw158
			(_nw158).ArraySet1(_nwElement0_35, 0)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_((_this).F5(), _815_v11, func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter48 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(171), _dafny.IntOfInt64(954))); ; {
					_compr_43, _ok48 := _iter48()
					if !_ok48 {
						break
					}
					var _820_v14 _dafny.Int
					_820_v14 = interface{}(_compr_43).(_dafny.Int)
					if ((_dafny.IntOfInt64(171)).Cmp(_820_v14) <= 0) && ((_820_v14).Cmp(_dafny.IntOfInt64(954)) < 0) {
						_coll43.Add((_820_v14).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ljpgjdur")).Cardinality())), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))
					}
				}
				return _coll43.ToMap()
			}()), 1)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_((_814_v10).Cardinality(), _815_v11, func() _dafny.Map {
				var _coll44 = _dafny.NewMapBuilder()
				_ = _coll44
				for _iter49 := _dafny.Iterate((_804_v0).Elements()); ; {
					_compr_44, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _821_v15 _dafny.Int
					_821_v15 = interface{}(_compr_44).(_dafny.Int)
					if (_804_v0).Contains(_821_v15) {
						_coll44.Add((_821_v15).Plus((_816_v12).Cardinality()), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))
					}
				}
				return _coll44.ToMap()
			}()), 2)
			(_nw158).ArraySet1((func() D5 {
				if (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool) {
					return _817_v13
				}
				return _817_v13
			})(), 3)
			(_nw158).ArraySet1(_817_v13, 4)
			(_nw158).ArraySet1(_817_v13, 5)
			(_nw158).ArraySet1(_817_v13, 6)
			(_nw158).ArraySet1(_817_v13, 7)
			(_nw158).ArraySet1(_817_v13, 8)
			(_nw158).ArraySet1(_817_v13, 9)
			(_nw158).ArraySet1(func(_pat_let13_0 D5) D5 {
				return func(_822_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let14_0 _dafny.Array) D5 {
						return func(_823_dt__update_hcf28_h0 _dafny.Array) D5 {
							return Companion_D5_.Create_DC17_((_822_dt__update__tmp_h0).Dtor_cf27(), _823_dt__update_hcf28_h0, (_822_dt__update__tmp_h0).Dtor_cf29())
						}(_pat_let14_0)
					}(_pat_let_tv9)
				}(_pat_let13_0)
			}(_817_v13), 10)
			(_nw158).ArraySet1(_817_v13, 11)
			(_nw158).ArraySet1(_817_v13, 12)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_((_dafny.Zero).Minus(_this.F6()), _815_v11, _816_v12), 13)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_((func() _dafny.Int {
				if (_818_v16).Contains(_806_v2) {
					return (_818_v16).Get(_806_v2).(_dafny.Int)
				}
				return (_this).F5()
			})(), _815_v11, func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter50 := _dafny.Iterate((_816_v12).Keys().Elements()); ; {
					_compr_45, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _824_v17 _dafny.Int
					_824_v17 = interface{}(_compr_45).(_dafny.Int)
					if (_816_v12).Contains(_824_v17) {
						_coll45.Add(Companion_Default___.SafeModuloInt(_824_v17, (_this).F5()), (_809_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_809_v5), 0))).Int()).(bool))
					}
				}
				return _coll45.ToMap()
			}()), 14)
			(_nw158).ArraySet1(_817_v13, 15)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_(_dafny.IntOfInt64(639), _815_v11, _816_v12), 16)
			(_nw158).ArraySet1(_817_v13, 17)
			(_nw158).ArraySet1(_817_v13, 18)
			(_nw158).ArraySet1(_817_v13, 19)
			(_nw158).ArraySet1(_817_v13, 20)
			(_nw158).ArraySet1(_817_v13, 21)
			(_nw158).ArraySet1(_817_v13, 22)
			(_nw158).ArraySet1(_817_v13, 23)
			(_nw158).ArraySet1(_817_v13, 24)
			(_nw158).ArraySet1(_817_v13, 25)
			(_nw158).ArraySet1(_817_v13, 26)
			(_nw158).ArraySet1(Companion_D5_.Create_DC17_(_this.F6(), _815_v11, _816_v12), 27)
			_819_v18 = _nw158
			var _nwElement0_36 D5 = _817_v13
			_ = _nwElement0_36
			var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(6))
			_ = _nw159
			(_nw159).ArraySet1(_nwElement0_36, 0)
			(_nw159).ArraySet1(_817_v13, 1)
			(_nw159).ArraySet1(_817_v13, 2)
			(_nw159).ArraySet1(_817_v13, 3)
			(_nw159).ArraySet1(_817_v13, 4)
			(_nw159).ArraySet1(_817_v13, 5)
			_819_v18 = _nw159
			_814_v10 = _814_v10
		} else {
			var _825_v19 _dafny.Map
			_ = _825_v19
			_825_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _this.F6())
			var _826_v20 _dafny.Map
			_ = _826_v20
			_826_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			(globalState).F2 = Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_825_v19).Contains((Companion_Default___.Fm10(_dafny.IntOfInt64(-670), _825_v19, p0, _this.F6(), globalState)).Cardinality()) {
					return (_825_v19).Get((Companion_Default___.Fm10(_dafny.IntOfInt64(-670), _825_v19, p0, _this.F6(), globalState)).Cardinality()).(_dafny.Int)
				}
				return _this.F6()
			})(), (_826_v20).Cardinality())
			var _827_v21 _dafny.CodePoint
			_ = _827_v21
			_827_v21 = _dafny.CodePoint('h')
			_827_v21 = _827_v21
			(globalState).F2 = _this.F6()
			var _828_v22 bool
			_ = _828_v22
			_828_v22 = false
			var _829_v23 _dafny.Sequence
			_ = _829_v23
			_829_v23 = _dafny.UnicodeSeqOfUtf8Bytes("d")
			var _830_v24 _dafny.Sequence
			_ = _830_v24
			_830_v24 = _dafny.SeqOf(_this.F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ml")).Cardinality()), (_this).F5(), _dafny.IntOfUint32((_829_v23).Cardinality()))
			var _831_v25 _dafny.Map
			_ = _831_v25
			_831_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), p0)
			var _832_v26 _dafny.Sequence
			_ = _832_v26
			_832_v26 = _dafny.SeqOf(!(!(_828_v22)), (func() bool {
				if (_831_v25).Contains((_this).F5()) {
					return (_831_v25).Get((_this).F5()).(bool)
				}
				return p0
			})())
			_828_v22 = (((_this).F5()).Minus(_dafny.IntOfUint32((_830_v24).Cardinality()))).Cmp(_dafny.IntOfUint32((_832_v26).Cardinality())) == 0
			var _833_v27 _dafny.Set
			_ = _833_v27
			_833_v27 = _dafny.SetOf(_dafny.IntOfInt64(880), _dafny.IntOfUint32((Companion_Default___.Fm11((_this).F5(), _dafny.IntOfInt64(574), _this.F6(), _this.F6(), globalState)).Cardinality()), _this.F6())
			_833_v27 = _dafny.SetOf((_this).F5(), (_dafny.Zero).Minus(_this.F6()), Companion_Default___.SafeModuloInt((_this).F5(), (_dafny.Zero).Minus(_this.F6())))
		}
		(_this).F6_set_((_this).F5())
		var _834_v28 bool
		_ = _834_v28
		_834_v28 = false
		_834_v28 = p0
		var _835_v30 _dafny.Sequence
		_ = _835_v30
		_835_v30 = _dafny.SeqOf(Companion_D3_.Create_DC8_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
		var _836_v31 D7
		_ = _836_v31
		_836_v31 = Companion_D7_.Create_DC21_(func() _dafny.Map {
			var _coll46 = _dafny.NewMapBuilder()
			_ = _coll46
			for _iter51 := _dafny.Iterate((_835_v30).Elements()); ; {
				_compr_46, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _837_v29 D3
				_837_v29 = interface{}(_compr_46).(D3)
				if _dafny.Companion_Sequence_.Contains(_835_v30, _837_v29) {
					_coll46.Add(_837_v29, false)
				}
			}
			return _coll46.ToMap()
		}())
		var _pat_let_tv10 = _834_v28
		_ = _pat_let_tv10
		var _pat_let_tv11 = _834_v28
		_ = _pat_let_tv11
		var _pat_let_tv12 = p0
		_ = _pat_let_tv12
		var _pat_let_tv13 = p0
		_ = _pat_let_tv13
		if func(_source11 D7) bool {
			if _source11.Is_DC22() {
				var _838___mcc_h1 _dafny.Array = _source11.Get_().(D7_DC22).Cf38
				_ = _838___mcc_h1
				var _839___mcc_h2 _dafny.Sequence = _source11.Get_().(D7_DC22).Cf39
				_ = _839___mcc_h2
				var _840_cf39 _dafny.Sequence = _839___mcc_h2
				_ = _840_cf39
				var _841_cf38 _dafny.Array = _838___mcc_h1
				_ = _841_cf38
				return _pat_let_tv10
			} else if _source11.Is_DC23() {
				var _842___mcc_h3 _dafny.Int = _source11.Get_().(D7_DC23).Cf40
				_ = _842___mcc_h3
				var _843___mcc_h4 _dafny.Sequence = _source11.Get_().(D7_DC23).Cf41
				_ = _843___mcc_h4
				var _844___mcc_h5 _dafny.Int = _source11.Get_().(D7_DC23).Cf42
				_ = _844___mcc_h5
				var _845_cf42 _dafny.Int = _844___mcc_h5
				_ = _845_cf42
				var _846_cf41 _dafny.Sequence = _843___mcc_h4
				_ = _846_cf41
				var _847_cf40 _dafny.Int = _842___mcc_h3
				_ = _847_cf40
				return !(_pat_let_tv11)
			} else if _source11.Is_DC24() {
				return _pat_let_tv12
			} else if _source11.Is_DC21() {
				var _848___mcc_h6 _dafny.Map = _source11.Get_().(D7_DC21).Cf37
				_ = _848___mcc_h6
				var _849_cf37 _dafny.Map = _848___mcc_h6
				_ = _849_cf37
				return true
			} else {
				var _850___mcc_h7 D7 = _source11.Get_().(D7_DC25).Cf43
				_ = _850___mcc_h7
				var _851_cf43 D7 = _850___mcc_h7
				_ = _851_cf43
				return _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(-767), (_dafny.SetOf(_pat_let_tv13)).Cardinality()), _dafny.SeqOf(_this.F6()))
			}
		}(_836_v31) {
			if p0 {
				_834_v28 = _834_v28
				var _852_v32 _dafny.CodePoint
				_ = _852_v32
				_852_v32 = _dafny.CodePoint('b')
				var _853_v33 _dafny.MultiSet
				_ = _853_v33
				_853_v33 = _dafny.MultiSetOf(_852_v32, _852_v32, _852_v32)
				var _854_v34 D9
				_ = _854_v34
				_854_v34 = Companion_D9_.Create_DC27_(_853_v33)
				_853_v33 = (_853_v33).Union((_854_v34).Dtor_cf45())
				var _855_v35 *C0
				_ = _855_v35
				var _nw160 *C0 = New_C0_()
				_ = _nw160
				_nw160.Ctor__(Companion_Default___.SafeModuloInt((_this).F5(), _dafny.IntOfInt64(541)))
				_855_v35 = _nw160
				_834_v28 = p0
				_834_v28 = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F6())).Cardinality()).Times((_this).F5())).Cmp((_855_v35).F8()) > 0
			} else {
				var _856_v36 _dafny.Sequence
				_ = _856_v36
				_856_v36 = _dafny.UnicodeSeqOfUtf8Bytes("b")
				_856_v36 = _856_v36
				var _857_v37 _dafny.Array
				_ = _857_v37
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_18
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Int = (func(_858_v36 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_859_i0 _dafny.Int) _dafny.Int {
							return (_859_i0).Minus(_dafny.IntOfUint32((_858_v36).Cardinality()))
						}
					})(_856_v36)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw161 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw161).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw161).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_857_v37 = _nw161
				_857_v37 = _857_v37
				var _860_v38 _dafny.Set
				_ = _860_v38
				_860_v38 = _dafny.SetOf(false, _834_v28)
				var _861_v39 *C1
				_ = _861_v39
				var _nw162 *C1 = New_C1_()
				_ = _nw162
				_nw162.Ctor__(((_this).Fm2(_this.F6(), _860_v38, globalState)).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F6()))), (func() _dafny.Int {
					if _834_v28 {
						return _this.F6()
					}
					return _dafny.IntOfInt64(941)
				})())
				_861_v39 = _nw162
				var _862_v40 *C1
				_ = _862_v40
				var _nw163 *C1 = New_C1_()
				_ = _nw163
				_nw163.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("srfo")).Cardinality()), (_dafny.Zero).Minus((_861_v39).Fm13((_dafny.Zero).Minus(_this.F6()), _dafny.MultiSetOf(_856_v36, _856_v36), globalState)))
				_862_v40 = _nw163
				var _863_v41 D0
				_ = _863_v41
				_863_v41 = Companion_D0_.Create_DC1_(false, _834_v28)
				var _864_v42 D9
				_ = _864_v42
				_864_v42 = Companion_D9_.Create_DC28_(_834_v28, _this.F6(), (_862_v40).Fm13(_dafny.IntOfInt64(971), _dafny.MultiSetOf(_856_v36, _856_v36), globalState))
				var _865_v43 _dafny.Sequence
				_ = _865_v43
				_865_v43 = _dafny.SeqOf(p0)
				var _866_v44 _dafny.Array
				_ = _866_v44
				var _nwElement0_37 bool = _834_v28
				_ = _nwElement0_37
				var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(26))
				_ = _nw164
				(_nw164).ArraySet1(_nwElement0_37, 0)
				(_nw164).ArraySet1(_834_v28, 1)
				(_nw164).ArraySet1(_834_v28, 2)
				(_nw164).ArraySet1(_834_v28, 3)
				(_nw164).ArraySet1(p0, 4)
				(_nw164).ArraySet1((_863_v41).Dtor_cf2(), 5)
				(_nw164).ArraySet1(_834_v28, 6)
				(_nw164).ArraySet1(p0, 7)
				(_nw164).ArraySet1(false, 8)
				(_nw164).ArraySet1(_834_v28, 9)
				(_nw164).ArraySet1(p0, 10)
				(_nw164).ArraySet1(p0, 11)
				(_nw164).ArraySet1(p0, 12)
				(_nw164).ArraySet1(_834_v28, 13)
				(_nw164).ArraySet1(p0, 14)
				(_nw164).ArraySet1(p0, 15)
				(_nw164).ArraySet1((_864_v42).Dtor_cf46(), 16)
				(_nw164).ArraySet1(_834_v28, 17)
				(_nw164).ArraySet1(true, 18)
				(_nw164).ArraySet1(p0, 19)
				(_nw164).ArraySet1(p0, 20)
				(_nw164).ArraySet1((_865_v43).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_865_v43).Cardinality()))).Uint32()).(bool), 21)
				(_nw164).ArraySet1(p0, 22)
				(_nw164).ArraySet1(p0, 23)
				(_nw164).ArraySet1(p0, 24)
				(_nw164).ArraySet1(_834_v28, 25)
				_866_v44 = _nw164
				var _867_v45 _dafny.Sequence
				_ = _867_v45
				_867_v45 = _dafny.SeqOf(_866_v44, _866_v44)
				var _868_v46 _dafny.Map
				_ = _868_v46
				_868_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F5())
				var _869_v48 _dafny.CodePoint
				_ = _869_v48
				_869_v48 = _dafny.CodePoint('b')
				var _870_v49 _dafny.Map
				_ = _870_v49
				_870_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_868_v46, _869_v48)
				var _871_v50 _dafny.Map
				_ = _871_v50
				_871_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_867_v45, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_868_v46, (_856_v36).Select((Companion_Default___.SafeIndex((func() _dafny.Set {
					var _coll47 = _dafny.NewBuilder()
					_ = _coll47
					for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(340), _dafny.IntOfInt64(549))); ; {
						_compr_47, _ok52 := _iter52()
						if !_ok52 {
							break
						}
						var _872_v47 _dafny.Int
						_872_v47 = interface{}(_compr_47).(_dafny.Int)
						if ((_dafny.IntOfInt64(340)).Cmp(_872_v47) <= 0) && ((_872_v47).Cmp(_dafny.IntOfInt64(549)) < 0) {
							_coll47.Add(Companion_Default___.SafeModuloInt(_872_v47, _dafny.IntOfUint32((_865_v43).Cardinality())))
						}
					}
					return _coll47.ToSet()
				}()).Cardinality(), _dafny.IntOfUint32((_856_v36).Cardinality()))).Uint32()).(_dafny.CodePoint))).Equals(_870_v49))
				_871_v50 = (_871_v50).Update(_867_v45, _834_v28)
			}
			_834_v28 = !(p0) || ((_dafny.MultiSetOf(_834_v28)).IsSubsetOf(_dafny.MultiSetOf(_834_v28)))
			var _873_v51 _dafny.Map
			_ = _873_v51
			_873_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F5())
			var _874_v52 _dafny.Map
			_ = _874_v52
			_874_v52 = (_873_v51).Merge(_873_v51)
			var _source12 _dafny.Map = _874_v52
			_ = _source12
			var _875___mcc_h0 _dafny.Map = _source12
			_ = _875___mcc_h0
			var _876_cf7 _dafny.Map = _875___mcc_h0
			_ = _876_cf7
			var _877_v53 _dafny.Sequence
			_ = _877_v53
			_877_v53 = _dafny.SeqOf((_dafny.Zero).Minus((_this).F5()), (_this).F5())
			var _878_v54 _dafny.Map
			_ = _878_v54
			_878_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _dafny.IntOfUint32((_877_v53).Cardinality()))
			var _879_v55 _dafny.CodePoint
			_ = _879_v55
			_879_v55 = _dafny.CodePoint('a')
			var _880_v56 _dafny.Map
			_ = _880_v56
			_880_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_878_v54, _879_v55)
			_880_v56 = _880_v56
			var _881_v58 D10
			_ = _881_v58
			_881_v58 = Companion_D10_.Create_DC29_(_878_v54)
			var _882_v59 _dafny.Sequence
			_ = _882_v59
			_882_v59 = _dafny.UnicodeSeqOfUtf8Bytes("dcfhvxpm")
			var _883_v60 _dafny.Map
			_ = _883_v60
			_883_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _882_v59)
			var _884_v61 _dafny.Array
			_ = _884_v61
			var _nwElement0_38 _dafny.Map = _878_v54
			_ = _nwElement0_38
			var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(26))
			_ = _nw165
			(_nw165).ArraySet1(_nwElement0_38, 0)
			(_nw165).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _this.F6()), 1)
			(_nw165).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_876_cf7).Cardinality(), (_this).F5()), 2)
			(_nw165).ArraySet1(func() _dafny.Map {
				var _coll48 = _dafny.NewMapBuilder()
				_ = _coll48
				for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(353), _dafny.IntOfInt64(769))); ; {
					_compr_48, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _885_v57 _dafny.Int
					_885_v57 = interface{}(_compr_48).(_dafny.Int)
					if ((_dafny.IntOfInt64(353)).Cmp(_885_v57) <= 0) && ((_885_v57).Cmp(_dafny.IntOfInt64(769)) < 0) {
						_coll48.Add((_885_v57).Plus(_this.F6()), _this.F6())
					}
				}
				return _coll48.ToMap()
			}(), 3)
			(_nw165).ArraySet1((_881_v58).Dtor_cf49(), 4)
			(_nw165).ArraySet1((_878_v54).Merge(_878_v54), 5)
			(_nw165).ArraySet1(_878_v54, 6)
			(_nw165).ArraySet1((_878_v54).Merge(_878_v54), 7)
			(_nw165).ArraySet1(_878_v54, 8)
			(_nw165).ArraySet1((_878_v54).Merge(_878_v54), 9)
			(_nw165).ArraySet1(_878_v54, 10)
			(_nw165).ArraySet1(_878_v54, 11)
			(_nw165).ArraySet1((_878_v54).Update(_this.F6(), (_883_v60).Cardinality()), 12)
			(_nw165).ArraySet1(_878_v54, 13)
			(_nw165).ArraySet1(_878_v54, 14)
			(_nw165).ArraySet1(_878_v54, 15)
			(_nw165).ArraySet1(_878_v54, 16)
			(_nw165).ArraySet1(_878_v54, 17)
			(_nw165).ArraySet1((_878_v54).Merge(_878_v54), 18)
			(_nw165).ArraySet1(_878_v54, 19)
			(_nw165).ArraySet1(_878_v54, 20)
			(_nw165).ArraySet1(_878_v54, 21)
			(_nw165).ArraySet1(_878_v54, 22)
			(_nw165).ArraySet1(_878_v54, 23)
			(_nw165).ArraySet1(_878_v54, 24)
			(_nw165).ArraySet1(_878_v54, 25)
			_884_v61 = _nw165
			var _886_v63 _dafny.Set
			_ = _886_v63
			_886_v63 = _dafny.SetOf(_this.F6())
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_884_v61), 0))
			_ = _index152
			(_884_v61).ArraySet1(func() _dafny.Map {
				var _coll49 = _dafny.NewMapBuilder()
				_ = _coll49
				for _iter54 := _dafny.Iterate((_886_v63).Elements()); ; {
					_compr_49, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _887_v62 _dafny.Int
					_887_v62 = interface{}(_compr_49).(_dafny.Int)
					if (_886_v63).Contains(_887_v62) {
						_coll49.Add(Companion_Default___.SafeDivisionInt(_887_v62, _this.F6()), (_this).F5())
					}
				}
				return _coll49.ToMap()
			}(), (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_884_v61), 0))
			_ = _index153
			(_884_v61).ArraySet1((_878_v54).Merge(func() _dafny.Map {
				var _coll50 = _dafny.NewMapBuilder()
				_ = _coll50
				for _iter55 := _dafny.Iterate((_878_v54).Keys().Elements()); ; {
					_compr_50, _ok55 := _iter55()
					if !_ok55 {
						break
					}
					var _888_v64 _dafny.Int
					_888_v64 = interface{}(_compr_50).(_dafny.Int)
					if (_878_v54).Contains(_888_v64) {
						_coll50.Add(Companion_Default___.SafeModuloInt(_888_v64, (_this).F5()), (_this).F5())
					}
				}
				return _coll50.ToMap()
			}()), (_index153).Int())
			_882_v59 = _882_v59
			(globalState).F2 = (_this).F5()
			var _889_v65 _dafny.MultiSet
			_ = _889_v65
			_889_v65 = _dafny.MultiSetOf((_this).F5())
			if (_889_v65).IsDisjointFrom(_889_v65) {
				_889_v65 = _889_v65
				var _890_v66 _dafny.Array
				_ = _890_v66
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw166
				_890_v66 = _nw166
				var _891_v67 _dafny.Set
				_ = _891_v67
				_891_v67 = _dafny.SetOf(p0, p0, _834_v28, !(p0))
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_890_v66), 0))
				_ = _index154
				(_890_v66).ArraySet1((_this).Fm2(_this.F6(), _891_v67, globalState), (_index154).Int())
				var _892_v68 _dafny.Sequence
				_ = _892_v68
				_892_v68 = _dafny.SeqOf(p0)
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_890_v66), 0))
				_ = _index155
				(_890_v66).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_892_v68), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-838))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}(func(_893_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-838))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_894_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))).Cardinality()))).Uint32(), _dafny.CodePoint('l')))).Cardinality()).Times((_this).F5()), (_index155).Int())
				var _895_v69 _dafny.CodePoint
				_ = _895_v69
				_895_v69 = _dafny.CodePoint('g')
				var _896_v70 _dafny.Map
				_ = _896_v70
				_896_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _895_v69)
				_895_v69 = (func() _dafny.CodePoint {
					if (_896_v70).Contains((func() _dafny.Int {
						if p0 {
							return (_890_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_890_v66), 0))).Int()).(_dafny.Int)
						}
						return _this.F6()
					})()) {
						return (_896_v70).Get((func() _dafny.Int {
							if p0 {
								return (_890_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen((_890_v66), 0))).Int()).(_dafny.Int)
							}
							return _this.F6()
						})()).(_dafny.CodePoint)
					}
					return _dafny.CodePoint('e')
				})()
				var _897_v71 _dafny.Array
				_ = _897_v71
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_19
				var _nw167 _dafny.Array
				_ = _nw167
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw167 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = (func(_898_p0 bool) func(_dafny.Int) bool {
						return func(_899_i2 _dafny.Int) bool {
							return _898_p0
						}
					})(p0)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw167 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw167).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw167).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_897_v71 = _nw167
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_897_v71), 0))
				_ = _index156
				(_897_v71).ArraySet1((_this).Fm1(globalState), (_index156).Int())
				var _900_v72 D1
				_ = _900_v72
				_900_v72 = Companion_D1_.Create_DC5_()
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_897_v71), 0))
				_ = _index157
				var _rhs165 bool = p0
				_ = _rhs165
				var _rhs166 _dafny.Int = (_this).F5()
				_ = _rhs166
				var _rhs167 bool = ((_this).Fm9(_900_v72, true, globalState)) && (_834_v28)
				_ = _rhs167
				var _lhs102 _dafny.Array = _897_v71
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(169), _dafny.ArrayLen((_897_v71), 0))
				_ = _lhs103
				var _lhs104 *GlobalState = globalState
				_ = _lhs104
				(_lhs102).ArraySet1(_rhs165, (_lhs103).Int())
				_lhs104.F2 = _rhs166
				_834_v28 = _rhs167
				var _901_v73 _dafny.Array
				_ = _901_v73
				var _902_v74 bool
				_ = _902_v74
				var _903_v75 _dafny.MultiSet
				_ = _903_v75
				var _904_v76 _dafny.Int
				_ = _904_v76
				var _out52 _dafny.Array
				_ = _out52
				var _out53 bool
				_ = _out53
				var _out54 _dafny.MultiSet
				_ = _out54
				var _out55 _dafny.Int
				_ = _out55
				_out52, _out53, _out54, _out55 = Companion_Default___.M0(_dafny.CodePoint('m'), globalState)
				_901_v73 = _out52
				_902_v74 = _out53
				_903_v75 = _out54
				_904_v76 = _out55
			} else {
				var _905_v77 D1
				_ = _905_v77
				_905_v77 = Companion_D1_.Create_DC5_()
				var _906_v78 _dafny.Map
				_ = _906_v78
				_906_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_905_v77, ((_this).F5()).Minus((_this).F5()))
				_906_v78 = (_906_v78).Update(_905_v77, (_dafny.Zero).Minus((_dafny.Zero).Minus(_this.F6())))
				var _907_v79 *C0
				_ = _907_v79
				var _nw168 *C0 = New_C0_()
				_ = _nw168
				_nw168.Ctor__(_dafny.IntOfInt64(780))
				_907_v79 = _nw168
				var _908_v80 _dafny.Sequence
				_ = _908_v80
				_908_v80 = _dafny.SeqOf(_834_v28, p0)
				var _909_v81 _dafny.Sequence
				_ = _909_v81
				_909_v81 = _dafny.SeqOf(_dafny.IntOfUint32((_908_v80).Cardinality()), _this.F6())
				var _910_v82 *C0
				_ = _910_v82
				var _nw169 *C0 = New_C0_()
				_ = _nw169
				_nw169.Ctor__((_this).Fm2((_this).F5(), Companion_Default___.Fm21(_dafny.IntOfInt64(-113), _dafny.IntOfUint32((_909_v81).Cardinality()), false, globalState), globalState))
				_910_v82 = _nw169
				var _911_v83 _dafny.Sequence
				_ = _911_v83
				_911_v83 = _dafny.SeqOf(_907_v79)
				_910_v82 = (_911_v83).Select((Companion_Default___.SafeIndex((_907_v79).F8(), _dafny.IntOfUint32((_911_v83).Cardinality()))).Uint32()).(*C0)
				var _912_v84 _dafny.Array
				_ = _912_v84
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw170
				_912_v84 = _nw170
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_912_v84), 0))
				_ = _index158
				(_912_v84).ArraySet1(_909_v81, (_index158).Int())
				var _913_v85 _dafny.Set
				_ = _913_v85
				_913_v85 = _dafny.SetOf(_889_v65, _889_v65)
				var _914_v86 *C1
				_ = _914_v86
				var _nw171 *C1 = New_C1_()
				_ = _nw171
				_nw171.Ctor__((_913_v85).Cardinality(), (_910_v82).F8())
				_914_v86 = _nw171
				var _915_v87 _dafny.Sequence
				_ = _915_v87
				_915_v87 = _dafny.SeqOf(_914_v86)
				var _916_v88 D11
				_ = _916_v88
				_916_v88 = Companion_D11_.Create_DC31_(_914_v86)
				var _917_v89 _dafny.Array
				_ = _917_v89
				var _nwElement0_39 *C1 = _914_v86
				_ = _nwElement0_39
				var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(23))
				_ = _nw172
				(_nw172).ArraySet1(_nwElement0_39, 0)
				(_nw172).ArraySet1(_914_v86, 1)
				(_nw172).ArraySet1(_914_v86, 2)
				(_nw172).ArraySet1(_914_v86, 3)
				(_nw172).ArraySet1(_914_v86, 4)
				(_nw172).ArraySet1(_914_v86, 5)
				(_nw172).ArraySet1(_914_v86, 6)
				(_nw172).ArraySet1((_915_v87).Select((Companion_Default___.SafeIndex((_910_v82).F8(), _dafny.IntOfUint32((_915_v87).Cardinality()))).Uint32()).(*C1), 7)
				(_nw172).ArraySet1(_914_v86, 8)
				(_nw172).ArraySet1(_914_v86, 9)
				(_nw172).ArraySet1(_914_v86, 10)
				(_nw172).ArraySet1((_916_v88).Dtor_cf51(), 11)
				(_nw172).ArraySet1(_914_v86, 12)
				(_nw172).ArraySet1(_914_v86, 13)
				(_nw172).ArraySet1(_914_v86, 14)
				(_nw172).ArraySet1(_914_v86, 15)
				(_nw172).ArraySet1(_914_v86, 16)
				(_nw172).ArraySet1(_914_v86, 17)
				(_nw172).ArraySet1(_914_v86, 18)
				(_nw172).ArraySet1(_914_v86, 19)
				(_nw172).ArraySet1(_914_v86, 20)
				(_nw172).ArraySet1(_914_v86, 21)
				(_nw172).ArraySet1(_914_v86, 22)
				_917_v89 = _nw172
				var _918_v90 _dafny.Set
				_ = _918_v90
				_918_v90 = _dafny.SetOf(_917_v89, _917_v89, _917_v89, _917_v89, _917_v89)
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_912_v84), 0))
				_ = _index159
				(_912_v84).ArraySet1(_dafny.SeqOf(((_dafny.SetOf(_917_v89)).Difference(_918_v90)).Cardinality(), (_910_v82).F8()), (_index159).Int())
				var _919_v91 _dafny.Array
				_ = _919_v91
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_20
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Int = func(_920_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_920_i3, _this.F6())
					}
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw173 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw173).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw173).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_919_v91 = _nw173
				var _921_v92 D5
				_ = _921_v92
				_921_v92 = Companion_D5_.Create_DC16_(_919_v91)
				var _922_v93 _dafny.Array
				_ = _922_v93
				var _nwElement0_40 _dafny.Array = _919_v91
				_ = _nwElement0_40
				var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(24))
				_ = _nw174
				(_nw174).ArraySet1(_nwElement0_40, 0)
				(_nw174).ArraySet1(_919_v91, 1)
				(_nw174).ArraySet1(_919_v91, 2)
				(_nw174).ArraySet1(_919_v91, 3)
				(_nw174).ArraySet1(_919_v91, 4)
				(_nw174).ArraySet1(_919_v91, 5)
				(_nw174).ArraySet1(_919_v91, 6)
				(_nw174).ArraySet1(_919_v91, 7)
				(_nw174).ArraySet1(_919_v91, 8)
				(_nw174).ArraySet1(_919_v91, 9)
				(_nw174).ArraySet1(_919_v91, 10)
				(_nw174).ArraySet1(_919_v91, 11)
				(_nw174).ArraySet1(_919_v91, 12)
				(_nw174).ArraySet1(_919_v91, 13)
				(_nw174).ArraySet1(_919_v91, 14)
				(_nw174).ArraySet1((_921_v92).Dtor_cf26(), 15)
				(_nw174).ArraySet1(_919_v91, 16)
				(_nw174).ArraySet1(_919_v91, 17)
				(_nw174).ArraySet1(_919_v91, 18)
				(_nw174).ArraySet1(_919_v91, 19)
				(_nw174).ArraySet1((_921_v92).Dtor_cf26(), 20)
				(_nw174).ArraySet1(_919_v91, 21)
				(_nw174).ArraySet1(_919_v91, 22)
				(_nw174).ArraySet1(_919_v91, 23)
				_922_v93 = _nw174
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_922_v93), 0))
				_ = _index160
				(_922_v93).ArraySet1(_919_v91, (_index160).Int())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_922_v93), 0))
				_ = _index161
				var _rhs168 _dafny.Int = (_dafny.Zero).Minus((_907_v79).F8())
				_ = _rhs168
				var _rhs169 _dafny.Array = _919_v91
				_ = _rhs169
				var _lhs105 *C2 = _this
				_ = _lhs105
				var _lhs106 _dafny.Array = _922_v93
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_922_v93), 0))
				_ = _lhs107
				_lhs105.F6_set_(_rhs168)
				(_lhs106).ArraySet1(_rhs169, (_lhs107).Int())
			}
			var _923_v94 _dafny.CodePoint
			_ = _923_v94
			_923_v94 = _dafny.CodePoint('u')
			_923_v94 = _923_v94
		} else {
			(globalState).F2 = (_this).F5()
			var _924_v95 _dafny.Array
			_ = _924_v95
			var _nwElement0_41 bool = p0
			_ = _nwElement0_41
			var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.One)
			_ = _nw175
			(_nw175).ArraySet1(_nwElement0_41, 0)
			_924_v95 = _nw175
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))
			_ = _index162
			(_924_v95).ArraySet1((_this.F6()).Cmp(_this.F6()) <= 0, (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))
			_ = _index163
			(_924_v95).ArraySet1(p0, (_index163).Int())
			_834_v28 = p0
			var _925_v96 _dafny.Sequence
			_ = _925_v96
			_925_v96 = _dafny.UnicodeSeqOfUtf8Bytes("hjndu")
			_925_v96 = _925_v96
			if (_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool) {
				_834_v28 = ((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool)) == (p0)
				var _926_v97 _dafny.MultiSet
				_ = _926_v97
				_926_v97 = _dafny.MultiSetOf((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool))
				var _927_v98 _dafny.Set
				_ = _927_v98
				_927_v98 = _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_925_v96).Cardinality())))
				var _928_v99 _dafny.MultiSet
				_ = _928_v99
				_928_v99 = _dafny.MultiSetOf((_926_v97).IsSubsetOf(_926_v97), (_927_v98).IsDisjointFrom(_927_v98))
				var _929_v100 _dafny.Sequence
				_ = _929_v100
				_929_v100 = _dafny.SeqOf((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool))
				var _930_v101 _dafny.Sequence
				_ = _930_v101
				_930_v101 = _dafny.SeqOf(_927_v98)
				_926_v97 = ((_928_v99).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_929_v100, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F6()), _dafny.IntOfUint32((_929_v100).Cardinality()))).Uint32(), p0)))).Update((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool), Companion_Default___.Abs((func() _dafny.Int {
					if false {
						return (_this).F5()
					}
					return ((_930_v101).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_930_v101).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
				})()))
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))
				_ = _index164
				(_924_v95).ArraySet1(!(_834_v28) || ((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool)), (_index164).Int())
				var _931_v102 *C1
				_ = _931_v102
				var _nw176 *C1 = New_C1_()
				_ = _nw176
				_nw176.Ctor__((_this.F6()).Plus((_this).F5()), _this.F6())
				_931_v102 = _nw176
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))
				_ = _index165
				(_924_v95).ArraySet1(_834_v28, (_index165).Int())
			} else {
				var _932_v103 *C1
				_ = _932_v103
				var _nw177 *C1 = New_C1_()
				_ = _nw177
				_nw177.Ctor__((_this).F5(), (_this).F5())
				_932_v103 = _nw177
				var _933_v104 _dafny.Map
				_ = _933_v104
				_933_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_932_v103, p0)
				var _934_v105 _dafny.Map
				_ = _934_v105
				_934_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(808))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}(func(_935_i4 _dafny.Int) _dafny.Int {
					return (_this).F5()
				})), _933_v104)
				var _936_v106 _dafny.Array
				_ = _936_v106
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_21
				var _nw178 _dafny.Array
				_ = _nw178
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw178 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Sequence = (func(_937_v95 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
						return func(_938_i5 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_937_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_937_v95), 0))).Int()).(bool), (_937_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_937_v95), 0))).Int()).(bool)), _dafny.SeqOf((_937_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_937_v95), 0))).Int()).(bool)))
						}
					})(_924_v95)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw178 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw178).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw178).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_936_v106 = _nw178
				var _939_v107 _dafny.Sequence
				_ = _939_v107
				_939_v107 = _dafny.SeqOf(_this.F6())
				var _rhs170 _dafny.Int = _this.F6()
				_ = _rhs170
				var _rhs171 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v107, _933_v104)
				_ = _rhs171
				var _rhs172 _dafny.Array = _936_v106
				_ = _rhs172
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				_lhs108.F2 = _rhs170
				_934_v105 = _rhs171
				_936_v106 = _rhs172
				_925_v96 = _dafny.UnicodeSeqOfUtf8Bytes("kobrta")
				var _940_v108 _dafny.Sequence
				_ = _940_v108
				_940_v108 = _dafny.SeqOf((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool), false)
				_940_v108 = (func() _dafny.Sequence {
					if (p0) == ((_924_v95).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))).Int()).(bool)) {
						return _dafny.Companion_Sequence_.Concatenate(_940_v108, _940_v108)
					}
					return _940_v108
				})()
				var _941_v109 D5
				_ = _941_v109
				_941_v109 = Companion_D5_.Create_DC18_(_834_v28, (_this).F5(), !(_834_v28))
				var _rhs173 _dafny.Int = (_dafny.IntOfUint32((_939_v107).Cardinality())).Minus(_this.F6())
				_ = _rhs173
				var _rhs174 D5 = _941_v109
				_ = _rhs174
				var _lhs109 *C2 = _this
				_ = _lhs109
				_lhs109.F6_set_(_rhs173)
				_941_v109 = _rhs174
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_924_v95), 0))
				_ = _index166
				(_924_v95).ArraySet1(p0, (_index166).Int())
			}
		}
		var _942_v110 _dafny.Sequence
		_ = _942_v110
		_942_v110 = _dafny.UnicodeSeqOfUtf8Bytes("gahubgf")
		var _943_v111 _dafny.Map
		_ = _943_v111
		_943_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_942_v110).Cardinality()), true)
		var _944_v112 D0
		_ = _944_v112
		_944_v112 = Companion_D0_.Create_DC1_(_834_v28, _834_v28)
		var _945_v113 _dafny.MultiSet
		_ = _945_v113
		_945_v113 = _dafny.MultiSetOf(false)
		var _946_v114 _dafny.Array
		_ = _946_v114
		var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw179
		_946_v114 = _nw179
		var _947_v115 _dafny.Map
		_ = _947_v115
		_947_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_946_v114, _834_v28)
		var _948_v116 _dafny.CodePoint
		_ = _948_v116
		_948_v116 = _dafny.CodePoint('a')
		var _949_v117 _dafny.Set
		_ = _949_v117
		_949_v117 = _dafny.SetOf(_this.F6(), _this.F6())
		var _950_v118 _dafny.Sequence
		_ = _950_v118
		_950_v118 = _dafny.SeqOf((_this).F5(), _this.F6(), (_this).F5())
		var _951_v119 _dafny.Map
		_ = _951_v119
		_951_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _943_v111)
		var _952_v120 _dafny.MultiSet
		_ = _952_v120
		_952_v120 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_950_v118).Cardinality()), _this.F6())).Cardinality(), p0), (func() _dafny.Map {
			if (_951_v119).Contains(p0) {
				return (_951_v119).Get(p0).(_dafny.Map)
			}
			return _943_v111
		})())
		var _953_v121 _dafny.Array
		_ = _953_v121
		var _nwElement0_42 bool = _834_v28
		_ = _nwElement0_42
		var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(26))
		_ = _nw180
		(_nw180).ArraySet1(_nwElement0_42, 0)
		(_nw180).ArraySet1((_this).Fm0(_dafny.MultiSetOf(_943_v111, _943_v111), p0, _942_v110, globalState), 1)
		(_nw180).ArraySet1(!(p0), 2)
		(_nw180).ArraySet1((_944_v112).Dtor_cf1(), 3)
		(_nw180).ArraySet1(_834_v28, 4)
		(_nw180).ArraySet1(!(!(_834_v28)), 5)
		(_nw180).ArraySet1((_dafny.MultiSetOf(p0)).Equals(_945_v113), 6)
		(_nw180).ArraySet1((func() bool {
			if (_947_v115).Contains(_946_v114) {
				return (_947_v115).Get(_946_v114).(bool)
			}
			return _834_v28
		})(), 7)
		(_nw180).ArraySet1(p0, 8)
		(_nw180).ArraySet1(true, 9)
		(_nw180).ArraySet1(!(p0), 10)
		(_nw180).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm11(_this.F6(), (_this).F5(), (_945_v113).Cardinality(), _this.F6(), globalState), (Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((Companion_Default___.Fm11(_this.F6(), (_this).F5(), (_945_v113).Cardinality(), _this.F6(), globalState)).Cardinality()))).Uint32(), _948_v116), _942_v110), 11)
		(_nw180).ArraySet1((_949_v117).IsSubsetOf(_949_v117), 12)
		(_nw180).ArraySet1(p0, 13)
		(_nw180).ArraySet1(p0, 14)
		(_nw180).ArraySet1(p0, 15)
		(_nw180).ArraySet1((_this).Fm0(_952_v120, p0, _942_v110, globalState), 16)
		(_nw180).ArraySet1(_834_v28, 17)
		(_nw180).ArraySet1(_834_v28, 18)
		(_nw180).ArraySet1(_834_v28, 19)
		(_nw180).ArraySet1((_945_v113).IsDisjointFrom(_945_v113), 20)
		(_nw180).ArraySet1(p0, 21)
		(_nw180).ArraySet1(p0, 22)
		(_nw180).ArraySet1((p0) || (false), 23)
		(_nw180).ArraySet1(_834_v28, 24)
		(_nw180).ArraySet1(_834_v28, 25)
		_953_v121 = _nw180
		for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_953_v121), 0))); ; {
			_guard_loop_5, _ok56 := _iter56()
			if !_ok56 {
				break
			}
			var _954_i6 _dafny.Int
			_954_i6 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_954_i6).Sign() != -1) && ((_954_i6).Cmp(_dafny.ArrayLen((_953_v121), 0)) < 0)) {
				(_953_v121).ArraySet1(((_this).F5()).Cmp(_this.F6()) == 0, (_954_i6).Int())
			}
		}
		var _955_v122 _dafny.Map
		_ = _955_v122
		_955_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _942_v110)
		var _956_v123 _dafny.Set
		_ = _956_v123
		_956_v123 = _dafny.SetOf(!(_834_v28))
		_955_v122 = (_955_v122).Update((_dafny.Zero).Minus((_this).Fm2(_this.F6(), _956_v123, globalState)), _942_v110)
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f6 _dafny.Int
	_f5 _dafny.Int
	_f7 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f6 = _dafny.Zero
	_this._f5 = _dafny.Zero
	_this._f7 = false
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

func (_this *C3) F6() _dafny.Int {
	return _this._f6
}
func (_this *C3) F6_set_(value _dafny.Int) {
	_this._f6 = value
}
func (_this *C3) F5() _dafny.Int {
	return _this._f5
}
func (_this *C3) Ctor__(f7 bool, f5 _dafny.Int, f6 _dafny.Int) {
	{
		(_this)._f7 = f7
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C3) Fm2(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return (((func() _dafny.Map {
			var _coll51 = _dafny.NewMapBuilder()
			_ = _coll51
			for _iter57 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("bbwcwhig"))).Keys().Elements()); ; {
				_compr_51, _ok57 := _iter57()
				if !_ok57 {
					break
				}
				var _957_v0 _dafny.Int
				_957_v0 = interface{}(_compr_51).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("bbwcwhig"))).Contains(_957_v0) {
					_coll51.Add((_957_v0).Minus((_this).F5()), _dafny.CodePoint('a'))
				}
			}
			return _coll51.ToMap()
		}()).Merge(func() _dafny.Map {
			var _coll52 = _dafny.NewMapBuilder()
			_ = _coll52
			for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(649), _dafny.IntOfInt64(978))); ; {
				_compr_52, _ok58 := _iter58()
				if !_ok58 {
					break
				}
				var _958_v1 _dafny.Int
				_958_v1 = interface{}(_compr_52).(_dafny.Int)
				if ((_dafny.IntOfInt64(649)).Cmp(_958_v1) <= 0) && ((_958_v1).Cmp(_dafny.IntOfInt64(978)) < 0) {
					_coll52.Add((_958_v1).Minus((_this).F5()), _dafny.CodePoint('h'))
				}
			}
			return _coll52.ToMap()
		}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F6())).Cardinality(), _dafny.CodePoint('x'))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F7())).Cardinality(), _dafny.CodePoint('u'))))).Cardinality()
	}
}
func (_this *C3) Fm0(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) bool {
	{
		if false {
			return (_this).F7()
		} else {
			return (_this).F7()
		}
	}
}
func (_this *C3) Fm1(globalState *GlobalState) bool {
	{
		return (_this).F7()
	}
}
func (_this *C3) M1(globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _959_i0 _dafny.Int
		_ = _959_i0
		_959_i0 = _dafny.Zero
		{
			for (Companion_D0_.Create_DC1_((_this).F7(), (_this).F7())).Dtor_cf2() {
				{
					if (_959_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_959_i0 = (_959_i0).Plus(_dafny.One)
					var _960_v0 _dafny.Map
					_ = _960_v0
					_960_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), false)
					r1 = (func() bool {
						if (_this).F7() {
							return (_this).F7()
						}
						return (func() bool {
							if (_960_v0).Contains((_this).F7()) {
								return (_960_v0).Get((_this).F7()).(bool)
							}
							return (_this).F7()
						})()
					})()
					var _961_v1 D0
					_ = _961_v1
					_961_v1 = Companion_D0_.Create_DC3_((_this).F7())
					r1 = (_961_v1).Dtor_cf3()
					var _962_v2 _dafny.Set
					_ = _962_v2
					_962_v2 = _dafny.SetOf(_this.F6())
					var _963_v3 _dafny.Array
					_ = _963_v3
					var _nw181 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
					_ = _nw181
					_963_v3 = _nw181
					var _964_v4 _dafny.Set
					_ = _964_v4
					_964_v4 = _dafny.SetOf((_this).F7(), (_this).F7(), (_this).F7())
					var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_963_v3), 0))
					_ = _index167
					(_963_v3).ArraySet1((_964_v4).IsSubsetOf(_964_v4), (_index167).Int())
					var _965_v5 D1
					_ = _965_v5
					_965_v5 = Companion_D1_.Create_DC4_(_dafny.SetOf(_this.F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfInt64(836))).Cardinality(), _this.F6(), (_this).F5(), _dafny.IntOfInt64(337)))
					var _966_v6 _dafny.Set
					_ = _966_v6
					_966_v6 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(853), (_this).F7()))
					var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_963_v3), 0))
					_ = _index168
					var _rhs175 _dafny.Set = (_965_v5).Dtor_cf4()
					_ = _rhs175
					var _rhs176 bool = (_this).F7()
					_ = _rhs176
					var _rhs177 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_966_v6).IsDisjointFrom(_966_v6), (_this).F7())
					_ = _rhs177
					var _lhs110 _dafny.Array = _963_v3
					_ = _lhs110
					var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_963_v3), 0))
					_ = _lhs111
					_962_v2 = _rhs175
					(_lhs110).ArraySet1(_rhs176, (_lhs111).Int())
					_960_v0 = _rhs177
					var _967_v7 _dafny.Sequence
					_ = _967_v7
					_967_v7 = _dafny.UnicodeSeqOfUtf8Bytes("eijplewq")
					var _968_v8 D1
					_ = _968_v8
					_968_v8 = Companion_D1_.Create_DC6_(_967_v7, (_this).F5())
					var _969_v9 _dafny.Sequence
					_ = _969_v9
					_969_v9 = _dafny.SeqOf((_this).F5(), (_this).F5(), (_dafny.Zero).Minus(_this.F6()), (_this).F5(), (_this).Fm2((_this).F5(), _964_v4, globalState))
					var _970_v10 _dafny.Sequence
					_ = _970_v10
					_970_v10 = _dafny.SeqOf((_dafny.Zero).Minus(_this.F6()), (_968_v8).Dtor_cf6(), (_969_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.IntOfUint32((_969_v9).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(240), _this.F6())
					r0 = (_970_v10).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(877))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg43 _dafny.Int) interface{} {
							return coer43(arg43)
						}
					}(func(_971_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("fnkhhatg")
					}))).Cardinality())), _dafny.IntOfUint32((_970_v10).Cardinality()))).Uint32()).(_dafny.Int)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r1 = (_this).F7()
		var _972_v11 _dafny.Array
		_ = _972_v11
		var _nw182 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw182
		_972_v11 = _nw182
		var _973_v12 _dafny.Map
		_ = _973_v12
		_973_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), Companion_D0_.Create_DC3_((_this).F7()))
		var _974_v13 _dafny.Map
		_ = _974_v13
		_974_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _973_v12)
		var _975_v14 _dafny.CodePoint
		_ = _975_v14
		_975_v14 = _dafny.CodePoint('q')
		var _976_v15 _dafny.Map
		_ = _976_v15
		_976_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_975_v14, _dafny.IntOfInt64(916))
		var _977_v16 _dafny.Sequence
		_ = _977_v16
		_977_v16 = _dafny.SeqOf(_976_v15)
		var _978_v17 _dafny.MultiSet
		_ = _978_v17
		_978_v17 = _dafny.MultiSetOf(_973_v12, (func() _dafny.Map {
			if (_974_v13).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_977_v16).Cardinality()))) {
				return (_974_v13).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_977_v16).Cardinality()))).(_dafny.Map)
			}
			return _973_v12
		})())
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))
		_ = _index169
		(_972_v11).ArraySet1(!((_978_v17).IsProperSubsetOf(_dafny.MultiSetOf(_973_v12))), (_index169).Int())
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))
		_ = _index170
		(_972_v11).ArraySet1((_this).F7(), (_index170).Int())
		var _979_v18 _dafny.MultiSet
		_ = _979_v18
		_979_v18 = _dafny.MultiSetOf((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool))
		var _980_i2 _dafny.Int
		_ = _980_i2
		_980_i2 = _dafny.Zero
		{
			for ((_979_v18).Cardinality()).Cmp((_dafny.Zero).Minus(_this.F6())) != 0 {
				{
					if (_980_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_980_i2 = (_980_i2).Plus(_dafny.One)
					var _981_v19 _dafny.Array
					_ = _981_v19
					var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
					_ = _nw183
					_981_v19 = _nw183
					var _982_v20 _dafny.Array
					_ = _982_v20
					var _nw184 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
					_ = _nw184
					_982_v20 = _nw184
					var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_981_v19), 0))
					_ = _index171
					(_981_v19).ArraySet1(_982_v20, (_index171).Int())
					var _983_v21 _dafny.Map
					_ = _983_v21
					_983_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F5()), _982_v20)
					var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_981_v19), 0))
					_ = _index172
					(_981_v19).ArraySet1((func() _dafny.Array {
						if (_983_v21).Contains(_this.F6()) {
							return (_983_v21).Get(_this.F6()).(_dafny.Array)
						}
						return _982_v20
					})(), (_index172).Int())
					var _984_v22 _dafny.Sequence
					_ = _984_v22
					_984_v22 = _dafny.UnicodeSeqOfUtf8Bytes("wgosnrps")
					var _985_v23 _dafny.Map
					_ = _985_v23
					_985_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(127))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_986_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_987_i3 _dafny.Int) _dafny.CodePoint {
							return _986_v14
						}
					})(_975_v14))), _984_v22), _dafny.ArrayCastTo((_981_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_981_v19), 0))).Int())))
					var _988_v24 D1
					_ = _988_v24
					_988_v24 = Companion_D1_.Create_DC6_(_984_v22, _this.F6())
					var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_981_v19), 0))
					_ = _index173
					(_981_v19).ArraySet1((func() _dafny.Array {
						if (_985_v23).Contains(_dafny.Companion_Sequence_.Concatenate((_988_v24).Dtor_cf5(), _984_v22)) {
							return (_985_v23).Get(_dafny.Companion_Sequence_.Concatenate((_988_v24).Dtor_cf5(), _984_v22)).(_dafny.Array)
						}
						return _982_v20
					})(), (_index173).Int())
					_975_v14 = _975_v14
					if (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool) {
						var _989_v25 _dafny.Sequence
						_ = _989_v25
						_989_v25 = _dafny.SeqOf((_this).F5())
						var _990_v26 _dafny.Sequence
						_ = _990_v26
						_990_v26 = _dafny.SeqOf(_dafny.SeqOf((_this).F5(), _this.F6()))
						var _991_v27 _dafny.Sequence
						_ = _991_v27
						_991_v27 = _dafny.SeqOf((_this).F7())
						r1 = (func() bool {
							if (func() bool {
								if (_this).F7() {
									return (_this).F7()
								}
								return (_this).F7()
							})() {
								return !(_dafny.Companion_Sequence_.IsProperPrefixOf(_989_v25, (_990_v26).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_dafny.IntOfUint32((_991_v27).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_990_v26).Cardinality()))).Uint32()).(_dafny.Sequence)))
							}
							return false
						})()
						var _992_v28 _dafny.Set
						_ = _992_v28
						_992_v28 = _dafny.SetOf(_972_v11, _972_v11, _972_v11)
						var _993_v29 _dafny.Array
						_ = _993_v29
						var _nwElement0_43 _dafny.Set = _992_v28
						_ = _nwElement0_43
						var _nw185 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(15))
						_ = _nw185
						(_nw185).ArraySet1(_nwElement0_43, 0)
						(_nw185).ArraySet1(_dafny.SetOf(_972_v11), 1)
						(_nw185).ArraySet1(_dafny.SetOf(_972_v11), 2)
						(_nw185).ArraySet1(_992_v28, 3)
						(_nw185).ArraySet1((_992_v28).Union(_992_v28), 4)
						(_nw185).ArraySet1(_992_v28, 5)
						(_nw185).ArraySet1(_dafny.SetOf(_972_v11, _972_v11, _972_v11, _972_v11), 6)
						(_nw185).ArraySet1((_992_v28).Difference(_992_v28), 7)
						(_nw185).ArraySet1(_992_v28, 8)
						(_nw185).ArraySet1(_dafny.SetOf(_972_v11), 9)
						(_nw185).ArraySet1(_992_v28, 10)
						(_nw185).ArraySet1(_dafny.SetOf(_972_v11), 11)
						(_nw185).ArraySet1(_992_v28, 12)
						(_nw185).ArraySet1(_992_v28, 13)
						(_nw185).ArraySet1(_992_v28, 14)
						_993_v29 = _nw185
						var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_993_v29), 0))
						_ = _index174
						(_993_v29).ArraySet1(_992_v28, (_index174).Int())
						var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_993_v29), 0))
						_ = _index175
						(_993_v29).ArraySet1(_992_v28, (_index175).Int())
						var _994_v30 _dafny.Set
						_ = _994_v30
						_994_v30 = _dafny.SetOf((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool))
						_994_v30 = ((func() _dafny.Set {
							if (_this).F7() {
								return _994_v30
							}
							return _994_v30
						})()).Intersection(_994_v30)
						var _995_v31 _dafny.Map
						_ = _995_v31
						_995_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_972_v11, (_this).F5())
						var _996_v32 _dafny.Map
						_ = _996_v32
						_996_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), (_995_v31).Equals(_995_v31))
						_996_v32 = (_996_v32).Merge(_996_v32)
						var _997_v33 _dafny.Sequence
						_ = _997_v33
						_997_v33 = _dafny.SeqOf(_984_v22, _dafny.UnicodeSeqOfUtf8Bytes("dajoybt"), _dafny.Companion_Sequence_.Update(_984_v22, (Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_984_v22).Cardinality()))).Uint32(), _975_v14), _984_v22)
						_984_v22 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_997_v33).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_997_v33).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("iugifadp")), _dafny.UnicodeSeqOfUtf8Bytes("viaetyyar"))
					} else {
						var _998_v34 _dafny.Sequence
						_ = _998_v34
						_998_v34 = _dafny.SeqOf(_this.F6(), _this.F6())
						var _999_v35 _dafny.Map
						_ = _999_v35
						_999_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F5()).Cmp((_this).F5()) != 0, (_998_v34).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_998_v34).Cardinality()))).Uint32()).(_dafny.Int))
						var _1000_v36 _dafny.Map
						_ = _1000_v36
						_1000_v36 = _999_v35
						_999_v35 = (_999_v35).Merge((_1000_v36))
						var _1001_v37 _dafny.Sequence
						_ = _1001_v37
						_1001_v37 = _dafny.SeqOf(_979_v18, _979_v18)
						(globalState).F2 = (Companion_Default___.Fm3((_this).F5(), _dafny.IntOfInt64(434), _1001_v37, _998_v34, globalState)).Dtor_cf6()
						r1 = (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool)
						var _1002_v38 *C0
						_ = _1002_v38
						var _nw186 *C0 = New_C0_()
						_ = _nw186
						_nw186.Ctor__(_dafny.IntOfUint32((_984_v22).Cardinality()))
						_1002_v38 = _nw186
						r0 = (_1002_v38).F8()
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))
		_ = _index176
		(_972_v11).ArraySet1((_this).F7(), (_index176).Int())
		if (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool) {
			var _1003_v39 *C0
			_ = _1003_v39
			var _nw187 *C0 = New_C0_()
			_ = _nw187
			_nw187.Ctor__((_this).F5())
			_1003_v39 = _nw187
			_1003_v39 = _1003_v39
			var _nw188 *C0 = New_C0_()
			_ = _nw188
			_nw188.Ctor__((_1003_v39).F8())
			_1003_v39 = _nw188
			var _1004_v40 _dafny.Sequence
			_ = _1004_v40
			_1004_v40 = _dafny.SeqOf((_this).F5())
			var _1005_v41 _dafny.Map
			_ = _1005_v41
			_1005_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v39, (_this).F7())
			var _1006_v42 _dafny.Map
			_ = _1006_v42
			_1006_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), _975_v14)
			var _1007_v43 _dafny.Array
			_ = _1007_v43
			var _nwElement0_44 _dafny.Int = ((_1003_v39).F8()).Minus(_dafny.IntOfUint32((_1004_v40).Cardinality()))
			_ = _nwElement0_44
			var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(11))
			_ = _nw189
			(_nw189).ArraySet1(_nwElement0_44, 0)
			(_nw189).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sksxj")).Cardinality()), 1)
			(_nw189).ArraySet1(((_1003_v39).F8()).Plus((_dafny.Zero).Minus((_1003_v39).F8())), 2)
			(_nw189).ArraySet1((_this).F5(), 3)
			(_nw189).ArraySet1((_1004_v40).Select((Companion_Default___.SafeIndex((_1005_v41).Cardinality(), _dafny.IntOfUint32((_1004_v40).Cardinality()))).Uint32()).(_dafny.Int), 4)
			(_nw189).ArraySet1(_this.F6(), 5)
			(_nw189).ArraySet1(_this.F6(), 6)
			(_nw189).ArraySet1((_dafny.Zero).Minus((_1003_v39).Fm4(_975_v14, Companion_Default___.Fm5(_1004_v40, (_this).F7(), _this.F6(), globalState), (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), globalState)), 7)
			(_nw189).ArraySet1((_1006_v42).Cardinality(), 8)
			(_nw189).ArraySet1(((_this).F5()).Minus(_this.F6()), 9)
			(_nw189).ArraySet1((_1003_v39).F8(), 10)
			_1007_v43 = _nw189
			var _1008_v44 _dafny.Map
			_ = _1008_v44
			_1008_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1007_v43, _this.F6())
			var _1009_v45 _dafny.Array
			_ = _1009_v45
			var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
			_ = _nw190
			_1009_v45 = _nw190
			var _1010_v46 _dafny.Sequence
			_ = _1010_v46
			_1010_v46 = _dafny.UnicodeSeqOfUtf8Bytes("gfn")
			var _1011_v47 _dafny.Map
			_ = _1011_v47
			_1011_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), _1007_v43)
			var _1012_v48 _dafny.MultiSet
			_ = _1012_v48
			_1012_v48 = _dafny.MultiSetOf(_972_v11)
			var _rhs178 _dafny.Array = (func() _dafny.Array {
				if (_1011_v47).Contains((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool)) {
					return (_1011_v47).Get((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _1007_v43
			})()
			_ = _rhs178
			var _rhs179 _dafny.Map = _1008_v44
			_ = _rhs179
			var _rhs180 _dafny.Array = _1009_v45
			_ = _rhs180
			var _rhs181 _dafny.Sequence = _1010_v46
			_ = _rhs181
			var _rhs182 _dafny.Int = (func() _dafny.Int {
				if (_1012_v48).Contains(_972_v11) {
					return (_1012_v48).Multiplicity(_972_v11)
				}
				return (_1003_v39).Fm4(_975_v14, Companion_Default___.Fm5(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(530))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_1013_v39 *C0) func(_dafny.Int) _dafny.Int {
					return func(_1014_i4 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_1013_v39).F8())).Cardinality()
					}
				})(_1003_v39))), (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), (_this).F5(), globalState), (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), globalState)
			})()
			_ = _rhs182
			_1007_v43 = _rhs178
			_1008_v44 = _rhs179
			_1009_v45 = _rhs180
			_1010_v46 = _rhs181
			r0 = _rhs182
			var _1015_v49 _dafny.Map
			_ = _1015_v49
			_1015_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), _this.F6())
			(globalState).F2 = Companion_Default___.SafeDivisionInt(((_1015_v49).Cardinality()).Minus((_1003_v39).F8()), (_1003_v39).F8())
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))
			_ = _index177
			(_972_v11).ArraySet1((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), (_index177).Int())
		} else {
			(globalState).F2 = (_this).F5()
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))
			_ = _index178
			(_972_v11).ArraySet1(true, (_index178).Int())
			r1 = ((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool)) && (false)
			var _1016_v50 _dafny.Array
			_ = _1016_v50
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(18))
			_ = _nw191
			_1016_v50 = _nw191
			var _1017_v51 _dafny.Sequence
			_ = _1017_v51
			_1017_v51 = _dafny.SeqOf(_1016_v50, _1016_v50, _1016_v50)
			_1016_v50 = (_1017_v51).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_1017_v51).Cardinality()))).Uint32()).(_dafny.Array)
			var _1018_v52 _dafny.Array
			_ = _1018_v52
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_22
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Sequence = (func(_1019_v11 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
					return func(_1020_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F7()), _dafny.SeqOf((Companion_D0_.Create_DC1_(!((_1019_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_1019_v11), 0))).Int()).(bool)), (_1019_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_1019_v11), 0))).Int()).(bool))).Dtor_cf1()))
					}
				})(_972_v11)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw192 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw192).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw192).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_1018_v52 = _nw192
			var _1021_v53 _dafny.Sequence
			_ = _1021_v53
			_1021_v53 = _dafny.SeqOf((_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool), (_972_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_972_v11), 0))).Int()).(bool))
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_1018_v52), 0))
			_ = _index179
			(_1018_v52).ArraySet1(_1021_v53, (_index179).Int())
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_1018_v52), 0))
			_ = _index180
			(_1018_v52).ArraySet1(_1021_v53, (_index180).Int())
		}
		var _1022_v54 _dafny.Array
		_ = _1022_v54
		var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw193
		_1022_v54 = _nw193
		var _1023_v55 _dafny.Map
		_ = _1023_v55
		_1023_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _1022_v54)
		r0 = ((_this).F5()).Times(Companion_Default___.SafeDivisionInt(_this.F6(), (_1023_v55).Cardinality()))
		r1 = (_this).F7()
		return r0, r1
	}
}
func (_this *C3) M2(p0 _dafny.Int, globalState *GlobalState) {
	{
		if (func() bool {
			if (_this).F7() {
				return (true) == ((_this).F7())
			}
			return (_this).F7()
		})() {
			var _1024_v0 _dafny.Array
			_ = _1024_v0
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw194
			_1024_v0 = _nw194
			var _1025_v1 _dafny.Array
			_ = _1025_v1
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_23
			var _nw195 _dafny.Array
			_ = _nw195
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw195 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Int = func(_1026_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1026_i0, _this.F6())
				}
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw195 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw195).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw195).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1025_v1 = _nw195
			var _1027_v2 _dafny.Sequence
			_ = _1027_v2
			_1027_v2 = _dafny.SeqOf(_1025_v1, _1025_v1)
			var _1028_v3 _dafny.Sequence
			_ = _1028_v3
			_1028_v3 = _dafny.SeqOf(_1025_v1, (_1027_v2).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1027_v2).Cardinality()))).Uint32()).(_dafny.Array))
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1024_v0), 0))
			_ = _index181
			(_1024_v0).ArraySet1(_1027_v2, (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen((_1024_v0), 0))
			_ = _index182
			(_1024_v0).ArraySet1(_1028_v3, (_index182).Int())
			var _1029_v4 D0
			_ = _1029_v4
			_1029_v4 = Companion_D0_.Create_DC1_((_this).F7(), (_this).F7())
			var _1030_v5 _dafny.MultiSet
			_ = _1030_v5
			_1030_v5 = _dafny.MultiSetOf(_1029_v4)
			var _1031_v6 _dafny.Sequence
			_ = _1031_v6
			_1031_v6 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1025_v1, _this.F6())).Cardinality())
			var _1032_v7 _dafny.Set
			_ = _1032_v7
			_1032_v7 = _dafny.SetOf(_1030_v5, (_1030_v5).Update(_1029_v4, Companion_Default___.Abs(_dafny.IntOfUint32((_1031_v6).Cardinality()))))
			if (_1032_v7).IsSubsetOf((func() _dafny.Set {
				if false {
					return _1032_v7
				}
				return _1032_v7
			})()) {
				var _1033_v8 _dafny.CodePoint
				_ = _1033_v8
				_1033_v8 = _dafny.CodePoint('f')
				_1033_v8 = _dafny.CodePoint('t')
				var _1034_v9 _dafny.Sequence
				_ = _1034_v9
				_1034_v9 = _dafny.UnicodeSeqOfUtf8Bytes("jxjhrkcig")
				var _1035_v10 _dafny.Map
				_ = _1035_v10
				_1035_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p0)
				var _1036_v11 _dafny.Set
				_ = _1036_v11
				_1036_v11 = _dafny.SetOf((_this).F7())
				var _1037_v12 D1
				_ = _1037_v12
				_1037_v12 = Companion_D1_.Create_DC6_(_1034_v9, (Companion_D1_.Create_DC6_(Companion_Default___.Fm7(_1033_v8, (_this).F7(), (_this).F7(), (_this).F5(), globalState), (_1036_v11).Cardinality())).Dtor_cf6())
				var _1038_v13 _dafny.Sequence
				_ = _1038_v13
				_1038_v13 = _dafny.SeqOf((_this).F7())
				var _1039_v14 _dafny.Array
				_ = _1039_v14
				var _nwElement0_45 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-602))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_1040_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1041_i1 _dafny.Int) _dafny.CodePoint {
						return _1040_v8
					}
				})(_1033_v8))), _1034_v9), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-602))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_1042_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1043_i1 _dafny.Int) _dafny.CodePoint {
						return _1042_v8
					}
				})(_1033_v8))), _1034_v9)).Cardinality()))).Uint32(), Companion_Default___.Fm6((_1035_v10).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kjqkqeiu")).Cardinality()), (_this).F7(), (_this).F7(), globalState))
				_ = _nwElement0_45
				var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(21))
				_ = _nw196
				(_nw196).ArraySet1(_nwElement0_45, 0)
				(_nw196).ArraySet1((func() _dafny.Sequence {
					if !((_this).F7()) {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(224))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg48 _dafny.Int) interface{} {
								return coer48(arg48)
							}
						}((func(_1044_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1045_i2 _dafny.Int) _dafny.CodePoint {
								return _1044_v8
							}
						})(_1033_v8)))
					}
					return _1034_v9
				})(), 1)
				(_nw196).ArraySet1(_1034_v9, 2)
				(_nw196).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("me"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("me")).Cardinality()))).Uint32(), _1033_v8), 3)
				(_nw196).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm7(_1033_v8, (_this).F7(), (_this).F7(), _this.F6(), globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm7(_1033_v8, (_this).F7(), (_this).F7(), _this.F6(), globalState)).Cardinality()))).Uint32(), _1033_v8), 4)
				(_nw196).ArraySet1((_1037_v12).Dtor_cf5(), 5)
				(_nw196).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(430))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}((func(_1046_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1047_i3 _dafny.Int) _dafny.CodePoint {
						return _1046_v8
					}
				})(_1033_v8))), 6)
				(_nw196).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-983))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_1048_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1049_i4 _dafny.Int) _dafny.CodePoint {
						return _1048_v8
					}
				})(_1033_v8))), 7)
				(_nw196).ArraySet1(_1034_v9, 8)
				(_nw196).ArraySet1(_1034_v9, 9)
				(_nw196).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hcns"), 10)
				(_nw196).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gmyra"), 11)
				(_nw196).ArraySet1(_1034_v9, 12)
				(_nw196).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dyonfg"), 13)
				(_nw196).ArraySet1(_1034_v9, 14)
				(_nw196).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1034_v9, _dafny.UnicodeSeqOfUtf8Bytes("npwl")), 15)
				(_nw196).ArraySet1(_1034_v9, 16)
				(_nw196).ArraySet1(_1034_v9, 17)
				(_nw196).ArraySet1(Companion_Default___.Fm7(_dafny.CodePoint('i'), (_1038_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1034_v9).Cardinality()), _dafny.IntOfUint32((_1038_v13).Cardinality()))).Uint32()).(bool), (_this).F7(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("biesqdee")).Cardinality()), globalState), 18)
				(_nw196).ArraySet1(_1034_v9, 19)
				(_nw196).ArraySet1(_1034_v9, 20)
				_1039_v14 = _nw196
				var _1050_v15 _dafny.MultiSet
				_ = _1050_v15
				_1050_v15 = _dafny.MultiSetOf(_this.F6(), _this.F6())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1025_v1), 0))
				_ = _index183
				(_1025_v1).ArraySet1(((_dafny.MultiSetOf(p0)).Union(_1050_v15)).Cardinality(), (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1025_v1), 0))
				_ = _index184
				var _rhs183 _dafny.Array = _1039_v14
				_ = _rhs183
				var _rhs184 _dafny.Int = (_this).Fm2(_dafny.IntOfUint32((_1031_v6).Cardinality()), _1036_v11, globalState)
				_ = _rhs184
				var _lhs112 _dafny.Array = _1025_v1
				_ = _lhs112
				var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1025_v1), 0))
				_ = _lhs113
				_1039_v14 = _rhs183
				(_lhs112).ArraySet1(_rhs184, (_lhs113).Int())
				var _1051_v16 _dafny.Map
				_ = _1051_v16
				_1051_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v8, _1033_v8)
				var _1052_v17 _dafny.Array
				_ = _1052_v17
				var _nw197 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw197
				_1052_v17 = _nw197
				var _1053_v18 _dafny.MultiSet
				_ = _1053_v18
				_1053_v18 = _dafny.MultiSetOf(_1052_v17, _1052_v17, _1052_v17)
				var _1054_v19 T0
				_ = _1054_v19
				var _nw198 *C1 = New_C1_()
				_ = _nw198
				_nw198.Ctor__(_dafny.IntOfUint32((_1038_v13).Cardinality()), (_1053_v18).Cardinality())
				_1054_v19 = _nw198
				var _1055_v20 _dafny.Map
				_ = _1055_v20
				_1055_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1054_v19, _dafny.CodePoint('n'))
				_1033_v8 = (func() _dafny.CodePoint {
					if (_1051_v16).Contains((func() _dafny.CodePoint {
						if (_1055_v20).Contains(_1054_v19) {
							return (_1055_v20).Get(_1054_v19).(_dafny.CodePoint)
						}
						return _1033_v8
					})()) {
						return (_1051_v16).Get((func() _dafny.CodePoint {
							if (_1055_v20).Contains(_1054_v19) {
								return (_1055_v20).Get(_1054_v19).(_dafny.CodePoint)
							}
							return _1033_v8
						})()).(_dafny.CodePoint)
					}
					return (_1034_v9).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1034_v9).Cardinality()))).Uint32()).(_dafny.CodePoint)
				})()
				var _1056_v21 _dafny.Array
				_ = _1056_v21
				var _nw199 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw199
				_1056_v21 = _nw199
				var _1057_v22 D4
				_ = _1057_v22
				_1057_v22 = Companion_D4_.Create_DC11_(_1031_v6)
				var _1058_v23 _dafny.Map
				_ = _1058_v23
				_1058_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1056_v21, _1057_v22)
				_1058_v23 = (_1058_v23).Update(_1056_v21, func(_pat_let15_0 D4) D4 {
					return func(_1059_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let16_0 _dafny.Sequence) D4 {
							return func(_1061_dt__update_hcf13_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC11_(_1061_dt__update_hcf13_h0)
							}(_pat_let16_0)
						}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-764))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg51 _dafny.Int) interface{} {
								return coer51(arg51)
							}
						}(func(_1060_i5 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(456)
						})))
					}(_pat_let15_0)
				}(_1057_v22))
				_1038_v13 = _1038_v13
			} else {
				var _1062_v24 T1
				_ = _1062_v24
				var _nw200 *C1 = New_C1_()
				_ = _nw200
				_nw200.Ctor__((_this).F5(), (_this).F5())
				_1062_v24 = _nw200
				var _1063_v25 _dafny.Array
				_ = _1063_v25
				var _nwElement0_46 T1 = _1062_v24
				_ = _nwElement0_46
				var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(8))
				_ = _nw201
				(_nw201).ArraySet1(_nwElement0_46, 0)
				(_nw201).ArraySet1(_1062_v24, 1)
				(_nw201).ArraySet1(_1062_v24, 2)
				(_nw201).ArraySet1(_1062_v24, 3)
				(_nw201).ArraySet1(_1062_v24, 4)
				(_nw201).ArraySet1(_1062_v24, 5)
				(_nw201).ArraySet1(_1062_v24, 6)
				(_nw201).ArraySet1(_1062_v24, 7)
				_1063_v25 = _nw201
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1063_v25), 0))
				_ = _index185
				(_1063_v25).ArraySet1(_1062_v24, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_1063_v25), 0))
				_ = _index186
				(_1063_v25).ArraySet1(_1062_v24, (_index186).Int())
				var _1064_v26 *C0
				_ = _1064_v26
				var _nw202 *C0 = New_C0_()
				_ = _nw202
				_nw202.Ctor__(_this.F6())
				_1064_v26 = _nw202
				var _1065_v27 _dafny.Int
				_ = _1065_v27
				var _1066_v28 bool
				_ = _1066_v28
				var _out56 _dafny.Int
				_ = _out56
				var _out57 bool
				_ = _out57
				_out56, _out57 = (_1062_v24).M1(globalState)
				_1065_v27 = _out56
				_1066_v28 = _out57
				var _1067_v29 _dafny.Map
				_ = _1067_v29
				_1067_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1027_v2, (_1064_v26).F8())
				(_1062_v24).F6_set_((func() _dafny.Int {
					if (_1067_v29).Contains(_dafny.SeqOf(_1025_v1, _1025_v1)) {
						return (_1067_v29).Get(_dafny.SeqOf(_1025_v1, _1025_v1)).(_dafny.Int)
					}
					return (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(351))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg52 _dafny.Int) interface{} {
							return coer52(arg52)
						}
					}(func(_1068_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('c')
					}))).Cardinality())).Minus(_dafny.IntOfInt64(40))
				})())
				var _1069_v30 _dafny.CodePoint
				_ = _1069_v30
				_1069_v30 = _dafny.CodePoint('s')
				var _1070_v31 _dafny.Sequence
				_ = _1070_v31
				_1070_v31 = _dafny.UnicodeSeqOfUtf8Bytes("yern")
				_1066_v28 = _dafny.Companion_Sequence_.Contains(_1070_v31, _1069_v30)
			}
			_1029_v4 = (func() D0 {
				if (_this).F7() {
					return _1029_v4
				}
				return func(_pat_let17_0 D0) D0 {
					return func(_1071_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let18_0 bool) D0 {
							return func(_1072_dt__update_hcf1_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_1072_dt__update_hcf1_h0, (_1071_dt__update__tmp_h1).Dtor_cf2())
							}(_pat_let18_0)
						}((_this).F7())
					}(_pat_let17_0)
				}(_1029_v4)
			})()
			var _1073_v32 _dafny.Map
			_ = _1073_v32
			_1073_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
			var _1074_v33 D3
			_ = _1074_v33
			_1074_v33 = Companion_D3_.Create_DC8_(_1073_v32)
			var _pat_let_tv14 = _1073_v32
			_ = _pat_let_tv14
			var _source13 D3 = func(_pat_let19_0 D3) D3 {
				return func(_1075_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let20_0 _dafny.Map) D3 {
						return func(_1076_dt__update_hcf8_h0 _dafny.Map) D3 {
							return Companion_D3_.Create_DC8_(_1076_dt__update_hcf8_h0)
						}(_pat_let20_0)
					}(_pat_let_tv14)
				}(_pat_let19_0)
			}(_1074_v33)
			_ = _source13
			if _source13.Is_DC9() {
				var _1077___mcc_h0 _dafny.MultiSet = _source13.Get_().(D3_DC9).Cf9
				_ = _1077___mcc_h0
				var _1078___mcc_h1 D0 = _source13.Get_().(D3_DC9).Cf10
				_ = _1078___mcc_h1
				var _1079___mcc_h2 _dafny.Int = _source13.Get_().(D3_DC9).Cf11
				_ = _1079___mcc_h2
				var _1080_cf11 _dafny.Int = _1079___mcc_h2
				_ = _1080_cf11
				var _1081_cf10 D0 = _1078___mcc_h1
				_ = _1081_cf10
				var _1082_cf9 _dafny.MultiSet = _1077___mcc_h0
				_ = _1082_cf9
				var _1083_v34 _dafny.Array
				_ = _1083_v34
				var _nw203 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
				_ = _nw203
				_1083_v34 = _nw203
				var _1084_v35 *C2
				_ = _1084_v35
				var _nw204 *C2 = New_C2_()
				_ = _nw204
				_nw204.Ctor__((_this).F5(), _this.F6())
				_1084_v35 = _nw204
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_1083_v34), 0))
				_ = _index187
				(_1083_v34).ArraySet1(_1084_v35, (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_1083_v34), 0))
				_ = _index188
				(_1083_v34).ArraySet1(_1084_v35, (_index188).Int())
				var _1085_v36 _dafny.Map
				_ = _1085_v36
				_1085_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p0)
				var _1086_v37 _dafny.Map
				_ = _1086_v37
				_1086_v37 = _1085_v36
				_1086_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _this.F6())
				var _1087_v38 *C2
				_ = _1087_v38
				var _nw205 *C2 = New_C2_()
				_ = _nw205
				_nw205.Ctor__(((_this).F5()).Plus(_this.F6()), _this.F6())
				_1087_v38 = _nw205
				var _1088_v39 _dafny.Sequence
				_ = _1088_v39
				_1088_v39 = _dafny.SeqOf((Companion_D9_.Create_DC28_((_this).F7(), (_this).F5(), (_this).F5())).Dtor_cf46())
				_1088_v39 = _dafny.SeqOf((_this).F7(), (_this).F7(), (_this).F7())
			} else if _source13.Is_DC8() {
				var _1089___mcc_h3 _dafny.Map = _source13.Get_().(D3_DC8).Cf8
				_ = _1089___mcc_h3
				var _1090_cf8 _dafny.Map = _1089___mcc_h3
				_ = _1090_cf8
				var _1091_v40 bool
				_ = _1091_v40
				_1091_v40 = false
				var _1092_v41 _dafny.Set
				_ = _1092_v41
				_1092_v41 = _dafny.SetOf(_1073_v32, _1073_v32, _1073_v32, _1090_cf8)
				_1091_v40 = !(((_1092_v41).Union(_1092_v41)).IsDisjointFrom(_1092_v41))
				var _1093_v42 _dafny.Sequence
				_ = _1093_v42
				_1093_v42 = _dafny.SeqOf(true)
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1025_v1), 0))
				_ = _index189
				(_1025_v1).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F5(), _dafny.IntOfUint32((_1093_v42).Cardinality())), (_index189).Int())
				var _1094_v43 _dafny.Map
				_ = _1094_v43
				_1094_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _1025_v1)
				var _1095_v44 _dafny.Sequence
				_ = _1095_v44
				_1095_v44 = _dafny.UnicodeSeqOfUtf8Bytes("nxsi")
				var _1096_v45 _dafny.MultiSet
				_ = _1096_v45
				_1096_v45 = _dafny.MultiSetOf(_1095_v44)
				var _1097_v46 _dafny.CodePoint
				_ = _1097_v46
				_1097_v46 = _dafny.CodePoint('d')
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1025_v1), 0))
				_ = _index190
				var _rhs185 _dafny.Int = _this.F6()
				_ = _rhs185
				var _rhs186 _dafny.Int = _dafny.IntOfInt64(-212)
				_ = _rhs186
				var _rhs187 bool = ((_dafny.IntOfUint32((_1031_v6).Cardinality())).Plus(_this.F6())).Cmp(((func() _dafny.Int {
					if (_1096_v45).Contains(_1095_v44) {
						return (_1096_v45).Multiplicity(_1095_v44)
					}
					return _dafny.IntOfInt64(291)
				})()).Minus(p0)) >= 0
				_ = _rhs187
				var _rhs188 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1097_v46, _1025_v1)
				_ = _rhs188
				var _lhs114 _dafny.Array = _1025_v1
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1025_v1), 0))
				_ = _lhs115
				var _lhs116 *C3 = _this
				_ = _lhs116
				(_lhs114).ArraySet1(_rhs185, (_lhs115).Int())
				_lhs116.F6_set_(_rhs186)
				_1091_v40 = _rhs187
				_1094_v43 = _rhs188
				var _1098_v47 _dafny.Array
				_ = _1098_v47
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_24
				var _nw206 _dafny.Array
				_ = _nw206
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw206 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) D0 = (func(_1099_v40 bool) func(_dafny.Int) D0 {
						return func(_1100_i7 _dafny.Int) D0 {
							return Companion_D0_.Create_DC0_(_1099_v40)
						}
					})(_1091_v40)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw206 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw206).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw206).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1098_v47 = _nw206
				var _1101_v48 _dafny.Sequence
				_ = _1101_v48
				_1101_v48 = _dafny.SeqOf(_1098_v47)
				var _1102_v49 D1
				_ = _1102_v49
				_1102_v49 = Companion_D1_.Create_DC5_()
				var _1103_v50 _dafny.Array
				_ = _1103_v50
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
				_ = _nw207
				_1103_v50 = _nw207
				var _rhs189 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1098_v47), _1101_v48)
				_ = _rhs189
				var _rhs190 D1 = _1102_v49
				_ = _rhs190
				var _rhs191 _dafny.Array = _1103_v50
				_ = _rhs191
				_1101_v48 = _rhs189
				_1102_v49 = _rhs190
				_1103_v50 = _rhs191
				(globalState).F2 = (_this).F5()
			} else {
				var _1104___mcc_h4 D3 = _source13.Get_().(D3_DC10).Cf12
				_ = _1104___mcc_h4
				var _1105_cf12 D3 = _1104___mcc_h4
				_ = _1105_cf12
				var _1106_v51 _dafny.Sequence
				_ = _1106_v51
				_1106_v51 = _dafny.UnicodeSeqOfUtf8Bytes("b")
				var _1107_v52 _dafny.Map
				_ = _1107_v52
				_1107_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1031_v6, p0)
				var _1108_v53 D1
				_ = _1108_v53
				_1108_v53 = Companion_D1_.Create_DC6_(_1106_v51, _this.F6())
				var _1109_v54 T1
				_ = _1109_v54
				var _nw208 *C2 = New_C2_()
				_ = _nw208
				_nw208.Ctor__(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(826))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}(func(_1110_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality()), (Companion_Default___.Fm5(_dafny.SeqOf(p0, _this.F6(), p0), (_this).F7(), _dafny.IntOfUint32(((_1108_v53).Dtor_cf5()).Cardinality()), globalState)).Cardinality())
				_1109_v54 = _nw208
				var _1111_v55 _dafny.Array
				_ = _1111_v55
				var _nwElement0_47 bool = (_this).F7()
				_ = _nwElement0_47
				var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(14))
				_ = _nw209
				(_nw209).ArraySet1(_nwElement0_47, 0)
				(_nw209).ArraySet1(((_this).F7()) || ((_this).F7()), 1)
				(_nw209).ArraySet1((_this).F7(), 2)
				(_nw209).ArraySet1((_this).F7(), 3)
				(_nw209).ArraySet1((_this).F7(), 4)
				(_nw209).ArraySet1((_this).F7(), 5)
				(_nw209).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1106_v51, _1106_v51), 6)
				(_nw209).ArraySet1((_this).F7(), 7)
				(_nw209).ArraySet1((_this).F7(), 8)
				(_nw209).ArraySet1(!(((_this).F7()) || ((_this).F7())), 9)
				(_nw209).ArraySet1((p0).Cmp(_this.F6()) < 0, 10)
				(_nw209).ArraySet1((_1107_v52).Contains(Companion_Default___.Fm28(p0, globalState)), 11)
				(_nw209).ArraySet1((_dafny.SetOf(_1109_v54, _1109_v54, _1109_v54, _1109_v54, _1109_v54)).IsProperSubsetOf(_dafny.SetOf(_1109_v54)), 12)
				(_nw209).ArraySet1(!(((_this).F7()) || (false)), 13)
				_1111_v55 = _nw209
				var _1112_v56 _dafny.Set
				_ = _1112_v56
				_1112_v56 = _dafny.SetOf(_dafny.IntOfUint32((_1106_v51).Cardinality()))
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))
				_ = _index191
				(_1111_v55).ArraySet1((_1112_v56).Contains((_this).F5()), (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))
				_ = _index192
				(_1111_v55).ArraySet1(!((_this).F7()), (_index192).Int())
				var _1113_v57 _dafny.Sequence
				_ = _1113_v57
				_1113_v57 = _dafny.SeqOf((_1111_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))).Int()).(bool))
				var _1114_v58 _dafny.MultiSet
				_ = _1114_v58
				_1114_v58 = _dafny.MultiSetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1111_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))).Int()).(bool)), _1113_v57)))
				(globalState).F2 = (_1114_v58).Cardinality()
				var _1115_v59 _dafny.Map
				_ = _1115_v59
				_1115_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1111_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ffv")).Cardinality()))
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))
				_ = _index193
				(_1111_v55).ArraySet1((_dafny.IntOfInt64(-77)).Cmp(Companion_Default___.SafeModuloInt((_this).Fm2((_1115_v59).Cardinality(), _dafny.SetOf((_this).F7(), (_1111_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))).Int()).(bool), (_this).F7()), globalState), p0)) <= 0, (_index193).Int())
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_1111_v55), 0))
				_ = _index194
				(_1111_v55).ArraySet1((_this).F7(), (_index194).Int())
			}
			var _1116_v60 *C2
			_ = _1116_v60
			var _nw210 *C2 = New_C2_()
			_ = _nw210
			_nw210.Ctor__((_1031_v6).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_1031_v6).Cardinality()))).Uint32()).(_dafny.Int), (_this).F5())
			_1116_v60 = _nw210
		} else {
			var _1117_v61 _dafny.CodePoint
			_ = _1117_v61
			_1117_v61 = _dafny.CodePoint('a')
			_1117_v61 = _1117_v61
			var _1118_v62 _dafny.Array
			_ = _1118_v62
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_25
			var _nw211 _dafny.Array
			_ = _nw211
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw211 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = func(_1119_i9 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1119_i9, (_this).F5())
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw211 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw211).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw211).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1118_v62 = _nw211
			var _1120_v63 _dafny.Map
			_ = _1120_v63
			_1120_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1117_v61, (_this).F7())
			var _1121_v64 _dafny.MultiSet
			_ = _1121_v64
			_1121_v64 = _dafny.MultiSetOf(_1120_v63, _1120_v63)
			var _1122_v65 D12
			_ = _1122_v65
			_1122_v65 = Companion_D12_.Create_DC33_(_1121_v64)
			var _1123_v66 _dafny.Sequence
			_ = _1123_v66
			_1123_v66 = _dafny.SeqOf(_1120_v63)
			var _1124_v67 _dafny.Array
			_ = _1124_v67
			var _nwElement0_48 _dafny.MultiSet = _1121_v64
			_ = _nwElement0_48
			var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(11))
			_ = _nw212
			(_nw212).ArraySet1(_nwElement0_48, 0)
			(_nw212).ArraySet1(_1121_v64, 1)
			(_nw212).ArraySet1(_1121_v64, 2)
			(_nw212).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm29(_this.F6(), globalState)), 3)
			(_nw212).ArraySet1(_1121_v64, 4)
			(_nw212).ArraySet1((_1122_v65).Dtor_cf55(), 5)
			(_nw212).ArraySet1((_1121_v64).Difference(_dafny.MultiSetFromSeq(_1123_v66)), 6)
			(_nw212).ArraySet1(_dafny.MultiSetOf(_1120_v63), 7)
			(_nw212).ArraySet1((_1121_v64).Intersection(_1121_v64), 8)
			(_nw212).ArraySet1(_1121_v64, 9)
			(_nw212).ArraySet1(_1121_v64, 10)
			_1124_v67 = _nw212
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_1124_v67), 0))
			_ = _index195
			(_1124_v67).ArraySet1(_1121_v64, (_index195).Int())
			var _1125_v68 _dafny.Sequence
			_ = _1125_v68
			_1125_v68 = _dafny.UnicodeSeqOfUtf8Bytes("tsvv")
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_1124_v67), 0))
			_ = _index196
			var _rhs192 _dafny.Array = _1118_v62
			_ = _rhs192
			var _rhs193 _dafny.CodePoint = (_1125_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1125_v68).Cardinality()), _dafny.IntOfUint32((_1125_v68).Cardinality()))).Uint32()).(_dafny.CodePoint)
			_ = _rhs193
			var _rhs194 _dafny.Int = p0
			_ = _rhs194
			var _rhs195 _dafny.MultiSet = _dafny.MultiSetOf(_1120_v63)
			_ = _rhs195
			var _lhs117 *C3 = _this
			_ = _lhs117
			var _lhs118 _dafny.Array = _1124_v67
			_ = _lhs118
			var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(713), _dafny.ArrayLen((_1124_v67), 0))
			_ = _lhs119
			_1118_v62 = _rhs192
			_1117_v61 = _rhs193
			_lhs117.F6_set_(_rhs194)
			(_lhs118).ArraySet1(_rhs195, (_lhs119).Int())
			_1125_v68 = _1125_v68
			var _1126_v69 bool
			_ = _1126_v69
			_1126_v69 = true
			var _1127_v70 _dafny.Set
			_ = _1127_v70
			_1127_v70 = _dafny.SetOf(_this.F6())
			var _rhs196 _dafny.CodePoint = _1117_v61
			_ = _rhs196
			var _rhs197 bool = (_dafny.SetOf(_this.F6())).IsProperSubsetOf((_1127_v70).Difference(_1127_v70))
			_ = _rhs197
			var _rhs198 _dafny.CodePoint = (func() _dafny.CodePoint {
				if ((_this).F5()).Cmp(_dafny.IntOfUint32((_1125_v68).Cardinality())) >= 0 {
					return _1117_v61
				}
				return _1117_v61
			})()
			_ = _rhs198
			_1117_v61 = _rhs196
			_1126_v69 = _rhs197
			_1117_v61 = _rhs198
			_1117_v61 = (func() _dafny.CodePoint {
				if (_this).F7() {
					return _1117_v61
				}
				return _1117_v61
			})()
		}
		var _1128_i10 _dafny.Int
		_ = _1128_i10
		_1128_i10 = _dafny.Zero
		{
			for !((_this).F7()) || (((_this).F7()) && ((_this).F7())) {
				{
					if (_1128_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1128_i10 = (_1128_i10).Plus(_dafny.One)
					var _1129_v71 _dafny.Array
					_ = _1129_v71
					var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
					_ = _nw213
					_1129_v71 = _nw213
					var _1130_v72 _dafny.Set
					_ = _1130_v72
					_1130_v72 = _dafny.SetOf((_this).F7())
					var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_1129_v71), 0))
					_ = _index197
					(_1129_v71).ArraySet1(_1130_v72, (_index197).Int())
					var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(775), _dafny.ArrayLen((_1129_v71), 0))
					_ = _index198
					(_1129_v71).ArraySet1(_1130_v72, (_index198).Int())
					var _1131_v73 _dafny.Map
					_ = _1131_v73
					_1131_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p0)
					var _1132_v74 _dafny.Sequence
					_ = _1132_v74
					_1132_v74 = _dafny.SeqOf(_dafny.IntOfInt64(-750), ((func() _dafny.Int {
						if (_1131_v73).Contains((_this).F7()) {
							return (_1131_v73).Get((_this).F7()).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-885)
					})()).Times(_dafny.IntOfInt64(537)))
					(globalState).F2 = (_1132_v74).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1132_v74).Cardinality()))).Uint32()).(_dafny.Int)
					var _1133_v75 *C1
					_ = _1133_v75
					var _nw214 *C1 = New_C1_()
					_ = _nw214
					_nw214.Ctor__((_dafny.Zero).Minus((func() _dafny.Int {
						if (Companion_D12_.Create_DC35_((_this).F7(), _this.F6(), _dafny.IntOfInt64(388))).Dtor_cf57() {
							return p0
						}
						return _this.F6()
					})()), p0)
					_1133_v75 = _nw214
					var _1134_v76 T1
					_ = _1134_v76
					var _nw215 *C2 = New_C2_()
					_ = _nw215
					_nw215.Ctor__(p0, _this.F6())
					_1134_v76 = _nw215
					_1134_v76 = _1134_v76
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		if !(!((_this).F7())) {
			var _1135_v77 _dafny.Set
			_ = _1135_v77
			_1135_v77 = _dafny.SetOf(!((_this).F7()))
			(globalState).F2 = ((_this).F5()).Minus(((_1135_v77).Cardinality()).Minus((_this).F5()))
			var _1136_v78 _dafny.Array
			_ = _1136_v78
			var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw216
			_1136_v78 = _nw216
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1136_v78), 0))
			_ = _index199
			(_1136_v78).ArraySet1(_this.F6(), (_index199).Int())
			var _1137_v79 _dafny.MultiSet
			_ = _1137_v79
			_1137_v79 = _dafny.MultiSetOf((_this).F7(), (_this).F7(), (_this).F7(), (_this).F7(), (_this).F7())
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1136_v78), 0))
			_ = _index200
			(_1136_v78).ArraySet1((func() _dafny.Int {
				if (_1137_v79).Contains((_this).F7()) {
					return (_1137_v79).Multiplicity((_this).F7())
				}
				return p0
			})(), (_index200).Int())
			var _1138_v80 _dafny.Array
			_ = _1138_v80
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw217
			_1138_v80 = _nw217
			var _1139_v81 _dafny.Sequence
			_ = _1139_v81
			_1139_v81 = _dafny.SeqOf((_this).F5())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1138_v80), 0))
			_ = _index201
			(_1138_v80).ArraySet1(_dafny.Companion_Sequence_.Equal(_1139_v81, _1139_v81), (_index201).Int())
			var _1140_v82 _dafny.Array
			_ = _1140_v82
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
			_ = _nw218
			_1140_v82 = _nw218
			var _1141_v83 _dafny.Sequence
			_ = _1141_v83
			_1141_v83 = _dafny.UnicodeSeqOfUtf8Bytes("thwvihdbx")
			var _1142_v84 _dafny.CodePoint
			_ = _1142_v84
			_1142_v84 = _dafny.CodePoint('m')
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1140_v82), 0))
			_ = _index202
			(_1140_v82).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1141_v83, _dafny.UnicodeSeqOfUtf8Bytes("no")), (Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1141_v83, _dafny.UnicodeSeqOfUtf8Bytes("no"))).Cardinality()))).Uint32(), _1142_v84), (_index202).Int())
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1138_v80), 0))
			_ = _index203
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1140_v82), 0))
			_ = _index204
			var _rhs199 _dafny.Int = (Companion_Default___.SafeModuloInt((_this).F5(), p0)).Plus((_this).Fm2(p0, _dafny.SetOf((_this).F7(), (_this).F7(), (_this).F7()), globalState))
			_ = _rhs199
			var _rhs200 bool = (_this).F7()
			_ = _rhs200
			var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rsn"), (Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rsn")).Cardinality()))).Uint32(), (_1141_v83).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1141_v83).Cardinality()))).Uint32()).(_dafny.CodePoint)), _dafny.Companion_Sequence_.Concatenate(_1141_v83, _1141_v83))
			_ = _rhs201
			var _lhs120 *GlobalState = globalState
			_ = _lhs120
			var _lhs121 _dafny.Array = _1138_v80
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1138_v80), 0))
			_ = _lhs122
			var _lhs123 _dafny.Array = _1140_v82
			_ = _lhs123
			var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_1140_v82), 0))
			_ = _lhs124
			_lhs120.F2 = _rhs199
			(_lhs121).ArraySet1(_rhs200, (_lhs122).Int())
			(_lhs123).ArraySet1(_rhs201, (_lhs124).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1138_v80), 0))
			_ = _index205
			(_1138_v80).ArraySet1(false, (_index205).Int())
			var _1143_v85 _dafny.Map
			_ = _1143_v85
			_1143_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1136_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1136_v78), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(525))
			var _1144_v86 _dafny.Map
			_ = _1144_v86
			_1144_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_this.F6(), (_dafny.Zero).Minus((_1136_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1136_v78), 0))).Int()).(_dafny.Int))), ((_this).Fm2(_dafny.IntOfInt64(301), _dafny.SetOf((_this).F7(), (_1138_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_1138_v80), 0))).Int()).(bool)), globalState)).Cmp((_1143_v85).Cardinality()) == 0)
			_1144_v86 = (_1144_v86).Update((_this).F5(), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_dafny.Zero).Minus((_1136_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen((_1136_v78), 0))).Int()).(_dafny.Int))), _1139_v81))
		} else {
			var _1145_v87 _dafny.Sequence
			_ = _1145_v87
			_1145_v87 = _dafny.SeqOf((_this).F7(), !((_this).F7()), (_this).F7(), (_this).F7())
			var _1146_v88 _dafny.Set
			_ = _1146_v88
			_1146_v88 = _dafny.SetOf(false)
			var _1147_v89 _dafny.Array
			_ = _1147_v89
			var _nwElement0_49 bool = (_this).F7()
			_ = _nwElement0_49
			var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(13))
			_ = _nw219
			(_nw219).ArraySet1(_nwElement0_49, 0)
			(_nw219).ArraySet1((_this).F7(), 1)
			(_nw219).ArraySet1((_this).F7(), 2)
			(_nw219).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1145_v87, _1145_v87), 3)
			(_nw219).ArraySet1((_1146_v88).IsDisjointFrom(_1146_v88), 4)
			(_nw219).ArraySet1((_this).F7(), 5)
			(_nw219).ArraySet1((_this).F7(), 6)
			(_nw219).ArraySet1(!(_1146_v88).Equals(_dafny.SetOf((_this).F7(), !((_this).F7()), (_this).F7(), (_this).F7(), (_this).F7())), 7)
			(_nw219).ArraySet1((_this).F7(), 8)
			(_nw219).ArraySet1((_this).F7(), 9)
			(_nw219).ArraySet1((func() bool {
				if (_this).F7() {
					return !(false)
				}
				return (_this).F7()
			})(), 10)
			(_nw219).ArraySet1((_this).F7(), 11)
			(_nw219).ArraySet1((_this).F7(), 12)
			_1147_v89 = _nw219
			_1147_v89 = _1147_v89
			if (_this).F7() {
				var _1148_v90 bool
				_ = _1148_v90
				_1148_v90 = false
				_1148_v90 = ((_this).F5()).Cmp((_this).Fm2(_this.F6(), _1146_v88, globalState)) >= 0
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1147_v89), 0))
				_ = _index206
				(_1147_v89).ArraySet1((_this).F7(), (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1147_v89), 0))
				_ = _index207
				(_1147_v89).ArraySet1((_this).Fm1(globalState), (_index207).Int())
				var _1149_v91 _dafny.Array
				_ = _1149_v91
				var _nwElement0_50 _dafny.Int = ((_this).F5()).Plus(_dafny.IntOfInt64(-587))
				_ = _nwElement0_50
				var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(4))
				_ = _nw220
				(_nw220).ArraySet1(_nwElement0_50, 0)
				(_nw220).ArraySet1((_this).F5(), 1)
				(_nw220).ArraySet1(p0, 2)
				(_nw220).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(35)), 3)
				_1149_v91 = _nw220
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1149_v91), 0))
				_ = _index208
				(_1149_v91).ArraySet1(_dafny.IntOfInt64(428), (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1149_v91), 0))
				_ = _index209
				(_1149_v91).ArraySet1(((func() _dafny.Int {
					if _1148_v90 {
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qiamx")).Cardinality())
					}
					return _this.F6()
				})()).Plus((p0).Times(_dafny.IntOfInt64(140))), (_index209).Int())
				var _1150_v92 D4
				_ = _1150_v92
				_1150_v92 = Companion_D4_.Create_DC14_(p0, _1148_v90, (_1149_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1149_v91), 0))).Int()).(_dafny.Int), true, _this.F6())
				var _1151_v93 _dafny.Map
				_ = _1151_v93
				_1151_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1150_v92).Dtor_cf22(), _1149_v91)
				_1151_v93 = (_1151_v93).Update((_this).F5(), _1149_v91)
				var _1152_v94 _dafny.Sequence
				_ = _1152_v94
				_1152_v94 = _dafny.SeqOf(p0)
				var _1153_v96 D1
				_ = _1153_v96
				_1153_v96 = Companion_D1_.Create_DC4_(func() _dafny.Set {
					var _coll53 = _dafny.NewBuilder()
					_ = _coll53
					for _iter59 := _dafny.Iterate((_1152_v94).Elements()); ; {
						_compr_53, _ok59 := _iter59()
						if !_ok59 {
							break
						}
						var _1154_v95 _dafny.Int
						_1154_v95 = interface{}(_compr_53).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1152_v94, _1154_v95) {
							_coll53.Add(Companion_Default___.SafeDivisionInt(_1154_v95, _dafny.IntOfUint32((_dafny.SeqOf(!(false), true)).Cardinality())))
						}
					}
					return _coll53.ToSet()
				}())
				var _1155_v97 _dafny.Set
				_ = _1155_v97
				_1155_v97 = _dafny.SetOf((_this).F5(), p0, (_1149_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1149_v91), 0))).Int()).(_dafny.Int), p0)
				var _pat_let_tv15 = _1155_v97
				_ = _pat_let_tv15
				_1153_v96 = func(_pat_let21_0 D1) D1 {
					return func(_1156_dt__update__tmp_h4 D1) D1 {
						return func(_pat_let22_0 _dafny.Set) D1 {
							return func(_1157_dt__update_hcf4_h0 _dafny.Set) D1 {
								return Companion_D1_.Create_DC4_(_1157_dt__update_hcf4_h0)
							}(_pat_let22_0)
						}(_pat_let_tv15)
					}(_pat_let21_0)
				}(_1153_v96)
			} else {
				var _1158_v98 bool
				_ = _1158_v98
				_1158_v98 = false
				var _1159_v99 _dafny.Sequence
				_ = _1159_v99
				_1159_v99 = _dafny.UnicodeSeqOfUtf8Bytes("ympy")
				var _1160_v100 T0
				_ = _1160_v100
				var _nw221 *C2 = New_C2_()
				_ = _nw221
				_nw221.Ctor__(p0, _this.F6())
				_1160_v100 = _nw221
				var _1161_v101 _dafny.MultiSet
				_ = _1161_v101
				_1161_v101 = _dafny.MultiSetOf(_1160_v100)
				var _1162_v103 _dafny.Sequence
				_ = _1162_v103
				_1162_v103 = _dafny.SeqOf(_this.F6())
				var _1163_v104 _dafny.Map
				_ = _1163_v104
				_1163_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F5())
				var _1164_v105 _dafny.CodePoint
				_ = _1164_v105
				_1164_v105 = _dafny.CodePoint('n')
				var _1165_v106 D5
				_ = _1165_v106
				_1165_v106 = Companion_D5_.Create_DC19_((_1146_v88).Cardinality(), _1164_v105, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), true)).Cardinality())
				var _rhs202 bool = !(_1161_v101).Contains(_1160_v100)
				_ = _rhs202
				var _rhs203 bool = ((func() _dafny.Map {
					var _coll54 = _dafny.NewMapBuilder()
					_ = _coll54
					for _iter60 := _dafny.Iterate((_1162_v103).Elements()); ; {
						_compr_54, _ok60 := _iter60()
						if !_ok60 {
							break
						}
						var _1166_v102 _dafny.Int
						_1166_v102 = interface{}(_compr_54).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_1162_v103, _1166_v102) {
							_coll54.Add(Companion_Default___.SafeModuloInt(_1166_v102, p0), _dafny.IntOfInt64(825))
						}
					}
					return _coll54.ToMap()
				}()).Merge(_1163_v104)).Contains((_this).Fm2((_this).F5(), _dafny.SetOf((_this).F7()), globalState))
				_ = _rhs203
				var _rhs204 bool = _1158_v98
				_ = _rhs204
				var _rhs205 _dafny.Sequence = (func() _dafny.Sequence {
					if true {
						return _dafny.Companion_Sequence_.Update(_1159_v99, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1159_v99).Cardinality()))).Uint32(), (_1165_v106).Dtor_cf34())
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(272))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}((func(_1167_v105 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1168_i11 _dafny.Int) _dafny.CodePoint {
							return _1167_v105
						}
					})(_1164_v105)))
				})()
				_ = _rhs205
				_1158_v98 = _rhs202
				_1158_v98 = _rhs203
				_1158_v98 = _rhs204
				_1159_v99 = _rhs205
				_1158_v98 = (_this).F7()
				var _1169_v107 *C2
				_ = _1169_v107
				var _nw222 *C2 = New_C2_()
				_ = _nw222
				_nw222.Ctor__(_dafny.IntOfInt64(-813), (_this).F5())
				_1169_v107 = _nw222
				var _1170_v108 _dafny.Array
				_ = _1170_v108
				var _len0_26 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_26
				var _nw223 _dafny.Array
				_ = _nw223
				if _len0_26.Cmp(_dafny.Zero) == 0 {
					_nw223 = _dafny.NewArray(_len0_26)
				} else {
					var _init26 func(_dafny.Int) _dafny.Sequence = func(_1171_i12 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("pfkrct")
					}
					_ = _init26
					var _element0_26 = _init26(_dafny.Zero)
					_ = _element0_26
					_nw223 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
					(_nw223).ArraySet1(_element0_26, 0)
					var _nativeLen0_26 = (_len0_26).Int()
					_ = _nativeLen0_26
					for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
						(_nw223).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
					}
				}
				_1170_v108 = _nw223
				_1170_v108 = (func() _dafny.Array {
					if (_1158_v98) || (false) {
						return _1170_v108
					}
					return _1170_v108
				})()
				_1158_v98 = _1158_v98
			}
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))
			_ = _index210
			(_1147_v89).ArraySet1((func() bool {
				if (_this).F7() {
					return true
				}
				return !((_this).F7())
			})(), (_index210).Int())
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))
			_ = _index211
			(_1147_v89).ArraySet1((_this).F7(), (_index211).Int())
			var _1172_v109 _dafny.Map
			_ = _1172_v109
			_1172_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1147_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))).Int()).(bool), (_1147_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))).Int()).(bool))
			_1172_v109 = (_1172_v109).Update((_1147_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))).Int()).(bool), (_this).F7())
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))
			_ = _index212
			(_1147_v89).ArraySet1((_1147_v89).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1147_v89), 0))).Int()).(bool), (_index212).Int())
		}
		var _1173_v110 _dafny.MultiSet
		_ = _1173_v110
		_1173_v110 = _dafny.MultiSetOf((_this).F5())
		(_this).F6_set_(Companion_Default___.SafeModuloInt((_this).F5(), (_1173_v110).Cardinality()))
		var _1174_v111 _dafny.Sequence
		_ = _1174_v111
		_1174_v111 = _dafny.UnicodeSeqOfUtf8Bytes("pfpqnji")
		var _1175_v112 _dafny.CodePoint
		_ = _1175_v112
		_1175_v112 = _dafny.CodePoint('s')
		var _1176_v113 _dafny.Map
		_ = _1176_v113
		_1176_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_this).F7())
		var _1177_v114 _dafny.Sequence
		_ = _1177_v114
		_1177_v114 = _dafny.SeqOf(_1176_v113)
		var _hi5 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm7(_1175_v112, (_this).Fm0(_dafny.MultiSetFromSeq(_1177_v114), (_this).F7(), _dafny.UnicodeSeqOfUtf8Bytes("pcfqhi"), globalState), (_this).F7(), (_this).F5(), globalState)).Cardinality())
		_ = _hi5
		for _1178_i13 := (p0).Minus(_dafny.IntOfUint32((_1174_v111).Cardinality())); _1178_i13.Cmp(_hi5) < 0; _1178_i13 = _1178_i13.Plus(_dafny.One) {
			var _1179_v115 D4
			_ = _1179_v115
			_1179_v115 = Companion_D4_.Create_DC12_((_this).F7(), _1178_i13, (_this).F7())
			if !((_1173_v110).Intersection(_dafny.MultiSetOf(_this.F6()))).Contains((_1179_v115).Dtor_cf15()) {
				var _1180_v116 _dafny.Sequence
				_ = _1180_v116
				_1180_v116 = _dafny.SeqOf((_this).F5(), (_this).F5())
				var _1181_v117 _dafny.Map
				_ = _1181_v117
				_1181_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1174_v111, _1180_v116)
				_1181_v117 = _1181_v117
				(globalState).F2 = _1178_i13
				var _1182_v118 _dafny.Sequence
				_ = _1182_v118
				_1182_v118 = _dafny.SeqOf((_this).F7())
				var _1183_v119 _dafny.Sequence
				_ = _1183_v119
				_1183_v119 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1182_v118, (Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_1182_v118).Cardinality()))).Uint32(), (_this).F7()))
				var _1184_v120 D12
				_ = _1184_v120
				_1184_v120 = Companion_D12_.Create_DC34_(!((_this).F7()))
				var _1185_v121 _dafny.Map
				_ = _1185_v121
				_1185_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), func(_pat_let23_0 D12) D12 {
					return func(_1186_dt__update__tmp_h5 D12) D12 {
						return func(_pat_let24_0 bool) D12 {
							return func(_1187_dt__update_hcf56_h0 bool) D12 {
								return Companion_D12_.Create_DC34_(_1187_dt__update_hcf56_h0)
							}(_pat_let24_0)
						}(!((_this).F7()))
					}(_pat_let23_0)
				}(_1184_v120))
				var _1188_v122 _dafny.Map
				_ = _1188_v122
				_1188_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1183_v119).Select((Companion_Default___.SafeIndex((_1185_v121).Cardinality(), _dafny.IntOfUint32((_1183_v119).Cardinality()))).Uint32()).(_dafny.Sequence), Companion_Default___.SafeDivisionInt(_1178_i13, _1178_i13))
				_1188_v122 = (_1188_v122).Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(true), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1174_v111).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Uint32(), true), Companion_Default___.SafeModuloInt((_1176_v113).Cardinality(), _this.F6()))
				var _1189_v123 bool
				_ = _1189_v123
				_1189_v123 = true
				var _1190_v124 _dafny.MultiSet
				_ = _1190_v124
				_1190_v124 = _dafny.MultiSetOf(_1175_v112, _1175_v112)
				_1189_v123 = (_1190_v124).IsSubsetOf(_1190_v124)
				_1189_v123 = _1189_v123
			} else {
				var _1191_v125 *C1
				_ = _1191_v125
				var _nw224 *C1 = New_C1_()
				_ = _nw224
				_nw224.Ctor__(_1178_i13, _this.F6())
				_1191_v125 = _nw224
				var _1192_v126 bool
				_ = _1192_v126
				_1192_v126 = true
				var _1193_v127 _dafny.Sequence
				_ = _1193_v127
				_1193_v127 = _dafny.SeqOf((_this).F7())
				_1192_v126 = !((_1193_v127).Select((Companion_Default___.SafeIndex(_1178_i13, _dafny.IntOfUint32((_1193_v127).Cardinality()))).Uint32()).(bool))
				var _1194_v128 _dafny.Array
				_ = _1194_v128
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_27
				var _nw225 _dafny.Array
				_ = _nw225
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw225 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = (func(_1195_v126 bool) func(_dafny.Int) bool {
						return func(_1196_i14 _dafny.Int) bool {
							return _1195_v126
						}
					})(_1192_v126)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw225 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw225).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw225).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1194_v128 = _nw225
				var _1197_v129 D7
				_ = _1197_v129
				_1197_v129 = Companion_D7_.Create_DC22_(_1194_v128, _1174_v111)
				var _1198_v130 D7
				_ = _1198_v130
				_1198_v130 = Companion_D7_.Create_DC25_(_1197_v129)
				var _1199_v131 _dafny.Set
				_ = _1199_v131
				_1199_v131 = _dafny.SetOf(_1198_v130)
				var _1200_v132 _dafny.Sequence
				_ = _1200_v132
				_1200_v132 = _dafny.SeqOf(_1199_v131, _1199_v131)
				var _1201_v133 *C2
				_ = _1201_v133
				var _nw226 *C2 = New_C2_()
				_ = _nw226
				_nw226.Ctor__(_dafny.IntOfUint32((_1174_v111).Cardinality()), ((_1199_v131).Union((_1200_v132).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.IntOfUint32((_1200_v132).Cardinality()))).Uint32()).(_dafny.Set))).Cardinality())
				_1201_v133 = _nw226
				var _1202_v134 _dafny.Array
				_ = _1202_v134
				var _nw227 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw227
				_1202_v134 = _nw227
				var _1203_v135 _dafny.Map
				_ = _1203_v135
				_1203_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v134, (_this).F5())
				_1203_v135 = (_1203_v135).Update(_1202_v134, _1178_i13)
				(globalState).F2 = (_this).F5()
			}
			var _1204_v136 D12
			_ = _1204_v136
			_1204_v136 = Companion_D12_.Create_DC35_((_this).F7(), _dafny.IntOfInt64(90), (_dafny.IntOfInt64(222)).Times(_this.F6()))
			_1204_v136 = _1204_v136
			var _1205_v137 _dafny.Set
			_ = _1205_v137
			_1205_v137 = _dafny.SetOf(true, true)
			var _1206_v138 _dafny.MultiSet
			_ = _1206_v138
			_1206_v138 = _dafny.MultiSetOf(_1205_v137, _1205_v137, _1205_v137)
			var _1207_v139 _dafny.Sequence
			_ = _1207_v139
			_1207_v139 = _dafny.SeqOf(_1205_v137, _1205_v137)
			if ((_1206_v138).Intersection(_dafny.MultiSetFromSeq(_1207_v139))).Contains((_1205_v137).Difference(_1205_v137)) {
				_1205_v137 = _1205_v137
				(globalState).F2 = Companion_Default___.SafeDivisionInt((p0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm28(_this.F6(), globalState)).Cardinality()))), (_this).F5())
				var _1208_v140 _dafny.Sequence
				_ = _1208_v140
				_1208_v140 = _dafny.SeqOf(false, (_1204_v136).Dtor_cf57())
				var _1209_v141 _dafny.Sequence
				_ = _1209_v141
				_1209_v141 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F7(), (_this).F7(), false)).Cardinality()), _1178_i13, (_this).Fm2((_this).F5(), Companion_Default___.Fm21((_dafny.MultiSetFromSeq(_1208_v140)).Cardinality(), (_dafny.Zero).Minus((_this).F5()), false, globalState), globalState), _this.F6())
				(_this).F6_set_((Companion_Default___.Fm5(_1209_v141, (_1205_v137).IsSubsetOf(_1205_v137), (_1178_i13).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(103))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1210_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1211_i15 _dafny.Int) _dafny.CodePoint {
						return _1210_v112
					}
				})(_1175_v112)))).Cardinality())), globalState)).Cardinality())
				var _1212_v142 _dafny.Array
				_ = _1212_v142
				var _nwElement0_51 bool = !((_this).F7())
				_ = _nwElement0_51
				var _nw228 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(5))
				_ = _nw228
				(_nw228).ArraySet1(_nwElement0_51, 0)
				(_nw228).ArraySet1(!(true), 1)
				(_nw228).ArraySet1((_this).F7(), 2)
				(_nw228).ArraySet1((_1178_i13).Cmp(_dafny.IntOfUint32((_1174_v111).Cardinality())) != 0, 3)
				(_nw228).ArraySet1((_this).F7(), 4)
				_1212_v142 = _nw228
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1212_v142), 0))
				_ = _index213
				(_1212_v142).ArraySet1((_this).F7(), (_index213).Int())
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1212_v142), 0))
				_ = _index214
				var _rhs206 _dafny.Sequence = _1174_v111
				_ = _rhs206
				var _rhs207 bool = true
				_ = _rhs207
				var _lhs125 _dafny.Array = _1212_v142
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(936), _dafny.ArrayLen((_1212_v142), 0))
				_ = _lhs126
				_1174_v111 = _rhs206
				(_lhs125).ArraySet1(_rhs207, (_lhs126).Int())
				_1208_v140 = _dafny.SeqOf(false)
			} else {
				(globalState).F2 = _1178_i13
				var _1213_v143 D4
				_ = _1213_v143
				_1213_v143 = Companion_D4_.Create_DC13_((_this).F7(), _1178_i13, (_this).F7())
				var _pat_let_tv16 = p0
				_ = _pat_let_tv16
				var _1214_v144 _dafny.Array
				_ = _1214_v144
				var _nwElement0_52 D4 = _1213_v143
				_ = _nwElement0_52
				var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(24))
				_ = _nw229
				(_nw229).ArraySet1(_nwElement0_52, 0)
				(_nw229).ArraySet1((func() D4 {
					if (_this).F7() {
						return func(_pat_let25_0 D4) D4 {
							return func(_1215_dt__update__tmp_h6 D4) D4 {
								return func(_pat_let26_0 bool) D4 {
									return func(_1216_dt__update_hcf17_h0 bool) D4 {
										return Companion_D4_.Create_DC13_(_1216_dt__update_hcf17_h0, (_1215_dt__update__tmp_h6).Dtor_cf18(), (_1215_dt__update__tmp_h6).Dtor_cf19())
									}(_pat_let26_0)
								}(false)
							}(_pat_let25_0)
						}(_1213_v143)
					}
					return Companion_Default___.Fm27(_this.F6(), globalState)
				})(), 1)
				(_nw229).ArraySet1(Companion_Default___.Fm27(_this.F6(), globalState), 2)
				(_nw229).ArraySet1(func(_pat_let27_0 D4) D4 {
					return func(_1217_dt__update__tmp_h7 D4) D4 {
						return func(_pat_let28_0 _dafny.Int) D4 {
							return func(_1218_dt__update_hcf18_h0 _dafny.Int) D4 {
								return Companion_D4_.Create_DC13_((_1217_dt__update__tmp_h7).Dtor_cf17(), _1218_dt__update_hcf18_h0, (_1217_dt__update__tmp_h7).Dtor_cf19())
							}(_pat_let28_0)
						}(_pat_let_tv16)
					}(_pat_let27_0)
				}(_1213_v143), 3)
				(_nw229).ArraySet1(_1213_v143, 4)
				(_nw229).ArraySet1(_1213_v143, 5)
				(_nw229).ArraySet1(_1213_v143, 6)
				(_nw229).ArraySet1(Companion_Default___.Fm27(_this.F6(), globalState), 7)
				(_nw229).ArraySet1(_1213_v143, 8)
				(_nw229).ArraySet1(Companion_D4_.Create_DC13_((_this).F7(), _this.F6(), (_this).F7()), 9)
				(_nw229).ArraySet1(_1213_v143, 10)
				(_nw229).ArraySet1(_1213_v143, 11)
				(_nw229).ArraySet1(Companion_D4_.Create_DC13_((_this).F7(), p0, (_this).F7()), 12)
				(_nw229).ArraySet1(_1213_v143, 13)
				(_nw229).ArraySet1(_1213_v143, 14)
				(_nw229).ArraySet1(Companion_Default___.Fm27(_this.F6(), globalState), 15)
				(_nw229).ArraySet1(_1213_v143, 16)
				(_nw229).ArraySet1((func() D4 {
					if (_this).F7() {
						return Companion_D4_.Create_DC13_((_this).F7(), _this.F6(), !((_this).F7()))
					}
					return Companion_D4_.Create_DC13_((_this).F7(), p0, (_this).F7())
				})(), 17)
				(_nw229).ArraySet1(Companion_D4_.Create_DC13_((_this).Fm1(globalState), _this.F6(), (_this).F7()), 18)
				(_nw229).ArraySet1(_1213_v143, 19)
				(_nw229).ArraySet1(_1213_v143, 20)
				(_nw229).ArraySet1(_1213_v143, 21)
				(_nw229).ArraySet1(_1213_v143, 22)
				(_nw229).ArraySet1(_1213_v143, 23)
				_1214_v144 = _nw229
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1214_v144), 0))
				_ = _index215
				(_1214_v144).ArraySet1((func() D4 {
					if (_this).F7() {
						return _1213_v143
					}
					return _1213_v143
				})(), (_index215).Int())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1214_v144), 0))
				_ = _index216
				(_1214_v144).ArraySet1(_1213_v143, (_index216).Int())
				var _1219_v145 _dafny.Array
				_ = _1219_v145
				var _nwElement0_53 _dafny.Int = _1178_i13
				_ = _nwElement0_53
				var _nw230 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(4))
				_ = _nw230
				(_nw230).ArraySet1(_nwElement0_53, 0)
				(_nw230).ArraySet1(p0, 1)
				(_nw230).ArraySet1(_this.F6(), 2)
				(_nw230).ArraySet1(p0, 3)
				_1219_v145 = _nw230
				var _1220_v146 _dafny.Map
				_ = _1220_v146
				_1220_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v137, _1178_i13)
				var _1221_v147 _dafny.Map
				_ = _1221_v147
				_1221_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfInt64(256))
				var _1222_v148 T1
				_ = _1222_v148
				var _nw231 *C2 = New_C2_()
				_ = _nw231
				_nw231.Ctor__(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-827))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_1223_i16 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality()), (_this).F5())
				_1222_v148 = _nw231
				var _1224_v149 _dafny.Set
				_ = _1224_v149
				_1224_v149 = _dafny.SetOf(_1222_v148, _1222_v148)
				var _nwElement0_54 _dafny.Int = p0
				_ = _nwElement0_54
				var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(15))
				_ = _nw232
				(_nw232).ArraySet1(_nwElement0_54, 0)
				(_nw232).ArraySet1(((_this).F5()).Minus(_1178_i13), 1)
				(_nw232).ArraySet1(_1178_i13, 2)
				(_nw232).ArraySet1(Companion_Default___.SafeDivisionInt(_1178_i13, _dafny.IntOfInt64(-158)), 3)
				(_nw232).ArraySet1((func() _dafny.Int {
					if (_1220_v146).Contains(_1205_v137) {
						return (_1220_v146).Get(_1205_v137).(_dafny.Int)
					}
					return (_this).F5()
				})(), 4)
				(_nw232).ArraySet1(((_1221_v147).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfInt64(627)))).Cardinality(), 5)
				(_nw232).ArraySet1((_this).F5(), 6)
				(_nw232).ArraySet1(p0, 7)
				(_nw232).ArraySet1((_1224_v149).Cardinality(), 8)
				(_nw232).ArraySet1(_1222_v148.F6(), 9)
				(_nw232).ArraySet1(_1222_v148.F6(), 10)
				(_nw232).ArraySet1((func() _dafny.Int {
					if (_1221_v147).Contains((_this).F7()) {
						return (_1221_v147).Get((_this).F7()).(_dafny.Int)
					}
					return (_1179_v115).Dtor_cf15()
				})(), 11)
				(_nw232).ArraySet1((func() _dafny.Int {
					if (_this).F7() {
						return _this.F6()
					}
					return _1178_i13
				})(), 12)
				(_nw232).ArraySet1(_1222_v148.F6(), 13)
				(_nw232).ArraySet1((_dafny.Zero).Minus((_1222_v148).Fm2(p0, _1205_v137, globalState)), 14)
				_1219_v145 = _nw232
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_1219_v145), 0))
				_ = _index217
				(_1219_v145).ArraySet1(_this.F6(), (_index217).Int())
				var _1225_v150 _dafny.MultiSet
				_ = _1225_v150
				_1225_v150 = _dafny.MultiSetOf((_this).F7())
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_1219_v145), 0))
				_ = _index218
				(_1219_v145).ArraySet1((_this).Fm2(_1222_v148.F6(), Companion_Default___.Fm21((_dafny.Zero).Minus(_dafny.IntOfUint32((_1174_v111).Cardinality())), (_1225_v150).Cardinality(), true, globalState), globalState), (_index218).Int())
				var _1226_v151 *C2
				_ = _1226_v151
				var _nw233 *C2 = New_C2_()
				_ = _nw233
				_nw233.Ctor__((_this).F5(), _1178_i13)
				_1226_v151 = _nw233
			}
			var _1227_v152 _dafny.Sequence
			_ = _1227_v152
			_1227_v152 = _dafny.SeqOf(_1178_i13)
			(_this).F6_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1227_v152, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1227_v152).Cardinality()))).Uint32(), _1178_i13)).Cardinality()))
		}
		if (_this.F6()).Cmp(_dafny.IntOfInt64(918)) <= 0 {
			var _1228_v153 bool
			_ = _1228_v153
			_1228_v153 = false
			_1228_v153 = (_this).F7()
			var _1229_v154 *C0
			_ = _1229_v154
			var _nw234 *C0 = New_C0_()
			_ = _nw234
			_nw234.Ctor__(p0)
			_1229_v154 = _nw234
			var _1230_v155 _dafny.Sequence
			_ = _1230_v155
			_1230_v155 = _dafny.SeqOf(_1229_v154)
			var _1231_v156 _dafny.Map
			_ = _1231_v156
			_1231_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1174_v111).Cardinality()), _1230_v155)
			var _1232_v157 _dafny.Set
			_ = _1232_v157
			_1232_v157 = _dafny.SetOf(_1228_v153)
			(_this).F6_set_((_this).Fm2((_1231_v156).Cardinality(), _1232_v157, globalState))
			var _1233_v158 D0
			_ = _1233_v158
			_1233_v158 = Companion_D0_.Create_DC2_()
			var _1234_v159 D3
			_ = _1234_v159
			_1234_v159 = Companion_D3_.Create_DC9_(_1173_v110, _1233_v158, _this.F6())
			var _1235_v160 D3
			_ = _1235_v160
			_1235_v160 = Companion_D3_.Create_DC10_(_1234_v159)
			_1235_v160 = Companion_D3_.Create_DC10_(_1234_v159)
			var _1236_v161 _dafny.Array
			_ = _1236_v161
			var _nw235 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw235
			_1236_v161 = _nw235
			var _1237_v162 _dafny.Set
			_ = _1237_v162
			_1237_v162 = _dafny.SetOf(_1236_v161)
			var _1238_v163 _dafny.Map
			_ = _1238_v163
			_1238_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1175_v112, _1237_v162)
			var _1239_v164 _dafny.Map
			_ = _1239_v164
			_1239_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v153, _1236_v161)
			var _1240_v165 _dafny.Array
			_ = _1240_v165
			var _nwElement0_55 _dafny.Set = _1237_v162
			_ = _nwElement0_55
			var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(15))
			_ = _nw236
			(_nw236).ArraySet1(_nwElement0_55, 0)
			(_nw236).ArraySet1(_1237_v162, 1)
			(_nw236).ArraySet1(_dafny.SetOf(_1236_v161, _1236_v161, _1236_v161, _1236_v161, _1236_v161), 2)
			(_nw236).ArraySet1(_1237_v162, 3)
			(_nw236).ArraySet1(_1237_v162, 4)
			(_nw236).ArraySet1((_1237_v162).Union((func() _dafny.Set {
				if (_1238_v163).Contains(_1175_v112) {
					return (_1238_v163).Get(_1175_v112).(_dafny.Set)
				}
				return _1237_v162
			})()), 5)
			(_nw236).ArraySet1(_1237_v162, 6)
			(_nw236).ArraySet1((_dafny.SetOf((func() _dafny.Array {
				if (_1239_v164).Contains(!(_1228_v153)) {
					return (_1239_v164).Get(!(_1228_v153)).(_dafny.Array)
				}
				return _1236_v161
			})())).Union(_1237_v162), 7)
			(_nw236).ArraySet1(_1237_v162, 8)
			(_nw236).ArraySet1(_1237_v162, 9)
			(_nw236).ArraySet1(_dafny.SetOf(_1236_v161, _1236_v161, _1236_v161), 10)
			(_nw236).ArraySet1(_1237_v162, 11)
			(_nw236).ArraySet1(_1237_v162, 12)
			(_nw236).ArraySet1(_dafny.SetOf(_1236_v161), 13)
			(_nw236).ArraySet1((_1237_v162).Union(_1237_v162), 14)
			_1240_v165 = _nw236
			var _1241_v166 _dafny.Map
			_ = _1241_v166
			_1241_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1174_v111).Cardinality()), _1237_v162)
			var _1242_v167 _dafny.Sequence
			_ = _1242_v167
			_1242_v167 = _dafny.SeqOf((func() _dafny.Set {
				if (_1241_v166).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-798))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1243_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1244_i17 _dafny.Int) _dafny.CodePoint {
						return _1243_v112
					}
				})(_1175_v112)))).Cardinality())) {
					return (_1241_v166).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-798))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg58 _dafny.Int) interface{} {
							return coer58(arg58)
						}
					}((func(_1245_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1246_i17 _dafny.Int) _dafny.CodePoint {
							return _1245_v112
						}
					})(_1175_v112)))).Cardinality())).(_dafny.Set)
				}
				return _1237_v162
			})())
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1240_v165), 0))
			_ = _index219
			(_1240_v165).ArraySet1((_1242_v167).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1242_v167).Cardinality()))).Uint32()).(_dafny.Set), (_index219).Int())
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_1240_v165), 0))
			_ = _index220
			(_1240_v165).ArraySet1(_1237_v162, (_index220).Int())
			if !(_1228_v153) {
				var _1247_v168 _dafny.Set
				_ = _1247_v168
				_1247_v168 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hkkcfpn")).Cardinality()))
				var _1248_v169 _dafny.Set
				_ = _1248_v169
				_1248_v169 = _dafny.SetOf(_dafny.IntOfInt64(-116), (_1247_v168).Cardinality(), (_1229_v154).F8(), (p0).Minus((_this).F5()))
				_1247_v168 = ((func() _dafny.Set {
					if (_this).F7() {
						return _1247_v168
					}
					return _1248_v169
				})()).Intersection(_1247_v168)
				var _1249_v170 T0
				_ = _1249_v170
				var _nw237 *C1 = New_C1_()
				_ = _nw237
				_nw237.Ctor__(_this.F6(), (func() _dafny.Int {
					if !(_1228_v153) {
						return _dafny.IntOfInt64(615)
					}
					return (_1229_v154).F8()
				})())
				_1249_v170 = _nw237
				var _1250_v171 _dafny.MultiSet
				_ = _1250_v171
				_1250_v171 = _dafny.MultiSetOf(_1228_v153, (_this).F7(), _1228_v153, _1228_v153, (_this).F7())
				var _1251_v172 _dafny.Array
				_ = _1251_v172
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw238
				_1251_v172 = _nw238
				var _1252_v173 _dafny.Map
				_ = _1252_v173
				_1252_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1251_v172, _1228_v153)
				var _1253_v174 *C1
				_ = _1253_v174
				var _nw239 *C1 = New_C1_()
				_ = _nw239
				_nw239.Ctor__(Companion_Default___.SafeDivisionInt((_1229_v154).F8(), (func() _dafny.Int {
					if (_1250_v171).Contains((_this).F7()) {
						return (_1250_v171).Multiplicity((_this).F7())
					}
					return (_1229_v154).F8()
				})()), ((_1252_v173).Cardinality()).Minus(_this.F6()))
				_1253_v174 = _nw239
				var _1254_v175 _dafny.Array
				_ = _1254_v175
				var _nw240 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(2))
				_ = _nw240
				_1254_v175 = _nw240
				var _1255_v176 _dafny.Sequence
				_ = _1255_v176
				_1255_v176 = _dafny.SeqOf(_dafny.IntOfUint32((_1174_v111).Cardinality()), _dafny.IntOfInt64(345))
				var _1256_v177 D4
				_ = _1256_v177
				_1256_v177 = Companion_D4_.Create_DC11_(_1255_v176)
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1254_v175), 0))
				_ = _index221
				(_1254_v175).ArraySet1(_1256_v177, (_index221).Int())
				var _1257_v178 _dafny.Map
				_ = _1257_v178
				_1257_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1251_v172, _1253_v174)
				var _1258_v179 _dafny.Map
				_ = _1258_v179
				_1258_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1229_v154, (_1229_v154).F8())
				var _1259_v180 _dafny.Map
				_ = _1259_v180
				_1259_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1258_v179, (_this).F7())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1254_v175), 0))
				_ = _index222
				var _rhs208 T0 = _1249_v170
				_ = _rhs208
				var _rhs209 *C1 = _1253_v174
				_ = _rhs209
				var _rhs210 D4 = _1256_v177
				_ = _rhs210
				var _rhs211 bool = (func() bool {
					if (_1259_v180).Contains((_1258_v179).Merge(_1258_v179)) {
						return (_1259_v180).Get((_1258_v179).Merge(_1258_v179)).(bool)
					}
					return _1228_v153
				})()
				_ = _rhs211
				var _rhs212 _dafny.Map = _1257_v178
				_ = _rhs212
				var _lhs127 _dafny.Array = _1254_v175
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen((_1254_v175), 0))
				_ = _lhs128
				_1249_v170 = _rhs208
				_1253_v174 = _rhs209
				(_lhs127).ArraySet1(_rhs210, (_lhs128).Int())
				_1228_v153 = _rhs211
				_1257_v178 = _rhs212
				var _1260_v181 _dafny.Sequence
				_ = _1260_v181
				_1260_v181 = _dafny.SeqOf(_1251_v172, _1251_v172, _1251_v172, _1251_v172, _1251_v172)
				_1228_v153 = (func() bool {
					if _1228_v153 {
						return !_dafny.Companion_Sequence_.Equal(_1260_v181, _1260_v181)
					}
					return (_1232_v157).IsProperSubsetOf(Companion_Default___.Fm21((_1229_v154).F8(), (_1229_v154).F8(), false, globalState))
				})()
				var _1261_v182 _dafny.MultiSet
				_ = _1261_v182
				_1261_v182 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nduue"), _1174_v111)
				var _1262_v183 *C1
				_ = _1262_v183
				var _nw241 *C1 = New_C1_()
				_ = _nw241
				_nw241.Ctor__(_this.F6(), (_1253_v174).Fm13(_dafny.IntOfUint32((_1174_v111).Cardinality()), _1261_v182, globalState))
				_1262_v183 = _nw241
				var _1263_v184 _dafny.Array
				_ = _1263_v184
				var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(7))
				_ = _nw242
				_1263_v184 = _nw242
				_1263_v184 = _1263_v184
			} else {
				var _1264_v185 _dafny.Array
				_ = _1264_v185
				var _nw243 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
				_ = _nw243
				_1264_v185 = _nw243
				_1264_v185 = _1264_v185
				var _1265_v186 _dafny.MultiSet
				_ = _1265_v186
				_1265_v186 = _dafny.MultiSetOf(_1228_v153)
				var _1266_v187 *C2
				_ = _1266_v187
				var _nw244 *C2 = New_C2_()
				_ = _nw244
				_nw244.Ctor__(p0, ((_1265_v186).Difference(_dafny.MultiSetOf(true))).Cardinality())
				_1266_v187 = _nw244
				_1266_v187 = _1266_v187
				var _1267_v188 _dafny.Sequence
				_ = _1267_v188
				_1267_v188 = _dafny.SeqOf((_this).F7())
				_1228_v153 = !(!((_1267_v188).Select((Companion_Default___.SafeIndex((_1229_v154).F8(), _dafny.IntOfUint32((_1267_v188).Cardinality()))).Uint32()).(bool)))
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_1236_v161), 0))
				_ = _index223
				(_1236_v161).ArraySet1((_this).F7(), (_index223).Int())
				var _1268_v189 _dafny.Sequence
				_ = _1268_v189
				_1268_v189 = _dafny.SeqOf((_1229_v154).F8(), (_1229_v154).F8())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_1236_v161), 0))
				_ = _index224
				(_1236_v161).ArraySet1(!(((_1229_v154).F8()).Cmp(_dafny.IntOfUint32((_1268_v189).Cardinality())) >= 0) || (_1228_v153), (_index224).Int())
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(406), _dafny.ArrayLen((_1236_v161), 0))
				_ = _index225
				(_1236_v161).ArraySet1((_this).F7(), (_index225).Int())
			}
		} else {
			if !((Companion_Default___.Fm30(globalState)).IsSubsetOf((_dafny.MultiSetOf(_1174_v111, _1174_v111)).Intersection(_dafny.MultiSetOf(_1174_v111)))) {
				var _1269_v190 *C0
				_ = _1269_v190
				var _nw245 *C0 = New_C0_()
				_ = _nw245
				_nw245.Ctor__(_dafny.IntOfInt64(389))
				_1269_v190 = _nw245
				var _1270_v191 _dafny.Sequence
				_ = _1270_v191
				_1270_v191 = _dafny.SeqOf((_1269_v190).F8())
				var _1271_v192 _dafny.MultiSet
				_ = _1271_v192
				_1271_v192 = _dafny.MultiSetOf((_this).F7())
				var _1272_v193 T0
				_ = _1272_v193
				var _nw246 *C1 = New_C1_()
				_ = _nw246
				_nw246.Ctor__((_1270_v191).Select((Companion_Default___.SafeIndex((_1269_v190).Fm4(_1175_v112, _1271_v192, (_this).F7(), globalState), _dafny.IntOfUint32((_1270_v191).Cardinality()))).Uint32()).(_dafny.Int), (_this).F5())
				_1272_v193 = _nw246
				var _1273_v194 _dafny.Map
				_ = _1273_v194
				_1273_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1272_v193, (_this).F7())
				var _1274_v195 D12
				_ = _1274_v195
				_1274_v195 = Companion_D12_.Create_DC35_((_this).F7(), ((_1273_v194).Merge(_1273_v194)).Cardinality(), _dafny.IntOfUint32((_1174_v111).Cardinality()))
				_1274_v195 = _1274_v195
				(_this).F6_set_(Companion_Default___.SafeModuloInt(p0, (_this).F5()))
				var _1275_v196 bool
				_ = _1275_v196
				_1275_v196 = false
				_1275_v196 = (_this).F7()
				var _1276_v197 _dafny.Array
				_ = _1276_v197
				var _1277_v198 bool
				_ = _1277_v198
				var _1278_v199 _dafny.MultiSet
				_ = _1278_v199
				var _1279_v200 _dafny.Int
				_ = _1279_v200
				var _out58 _dafny.Array
				_ = _out58
				var _out59 bool
				_ = _out59
				var _out60 _dafny.MultiSet
				_ = _out60
				var _out61 _dafny.Int
				_ = _out61
				_out58, _out59, _out60, _out61 = Companion_Default___.M0(_dafny.CodePoint('r'), globalState)
				_1276_v197 = _out58
				_1277_v198 = _out59
				_1278_v199 = _out60
				_1279_v200 = _out61
			} else {
				(globalState).F2 = _dafny.IntOfInt64(34)
				var _1280_v201 _dafny.Array
				_ = _1280_v201
				var _nw247 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
				_ = _nw247
				_1280_v201 = _nw247
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))
				_ = _index226
				(_1280_v201).ArraySet1((func() bool {
					if (_this).F7() {
						return (_this).F7()
					}
					return true
				})(), (_index226).Int())
				var _1281_v202 _dafny.Sequence
				_ = _1281_v202
				_1281_v202 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}((func(_1282_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1283_i18 _dafny.Int) _dafny.CodePoint {
						return _1282_v112
					}
				})(_1175_v112))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(773))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_1284_v112 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1285_i19 _dafny.Int) _dafny.CodePoint {
						return _1284_v112
					}
				})(_1175_v112))))
				var _1286_v203 _dafny.Map
				_ = _1286_v203
				_1286_v203 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32(((_1281_v202).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1281_v202).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _1287_v204 _dafny.Sequence
				_ = _1287_v204
				_1287_v204 = _dafny.SeqOf((_1286_v203).Cardinality(), p0)
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))
				_ = _index227
				(_1280_v201).ArraySet1(((_1287_v204).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality()), _dafny.IntOfUint32((_1287_v204).Cardinality()))).Uint32()).(_dafny.Int)).Cmp((_this).F5()) < 0, (_index227).Int())
				var _1288_v205 *C0
				_ = _1288_v205
				var _nw248 *C0 = New_C0_()
				_ = _nw248
				_nw248.Ctor__(_this.F6())
				_1288_v205 = _nw248
				var _1289_v206 _dafny.Array
				_ = _1289_v206
				var _nwElement0_56 *C0 = _1288_v205
				_ = _nwElement0_56
				var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(5))
				_ = _nw249
				(_nw249).ArraySet1(_nwElement0_56, 0)
				(_nw249).ArraySet1(_1288_v205, 1)
				(_nw249).ArraySet1(_1288_v205, 2)
				(_nw249).ArraySet1(_1288_v205, 3)
				(_nw249).ArraySet1(_1288_v205, 4)
				_1289_v206 = _nw249
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_1289_v206), 0))
				_ = _index228
				(_1289_v206).ArraySet1((func() *C0 {
					if (_1280_v201).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))).Int()).(bool) {
						return _1288_v205
					}
					return _1288_v205
				})(), (_index228).Int())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(966), _dafny.ArrayLen((_1289_v206), 0))
				_ = _index229
				(_1289_v206).ArraySet1(_1288_v205, (_index229).Int())
				_1286_v203 = (_1286_v203).Update((_this).F5(), Companion_Default___.SafeModuloInt(p0, _this.F6()))
				var _1290_v207 T1
				_ = _1290_v207
				var _nw250 *C1 = New_C1_()
				_ = _nw250
				_nw250.Ctor__((func() _dafny.Int {
					if (_this).F7() {
						return _this.F6()
					}
					return (_this).F5()
				})(), ((_1288_v205).F8()).Minus((_1288_v205).F8()))
				_1290_v207 = _nw250
				var _1291_v208 _dafny.Sequence
				_ = _1291_v208
				_1291_v208 = _dafny.SeqOf((_1280_v201).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))).Int()).(bool), (_this).F7(), (_this).F7())
				var _1292_v209 *C2
				_ = _1292_v209
				var _nw251 *C2 = New_C2_()
				_ = _nw251
				_nw251.Ctor__(_dafny.IntOfUint32((_1174_v111).Cardinality()), _dafny.IntOfUint32((_1291_v208).Cardinality()))
				_1292_v209 = _nw251
				var _1293_v210 _dafny.Sequence
				_ = _1293_v210
				_1293_v210 = _dafny.SeqOf(_1292_v209)
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))
				_ = _index230
				var _rhs213 T1 = _1290_v207
				_ = _rhs213
				var _rhs214 bool = !_dafny.Companion_Sequence_.Contains(_1293_v210, _1292_v209)
				_ = _rhs214
				var _rhs215 _dafny.CodePoint = _1175_v112
				_ = _rhs215
				var _rhs216 _dafny.Sequence = _1287_v204
				_ = _rhs216
				var _lhs129 _dafny.Array = _1280_v201
				_ = _lhs129
				var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen((_1280_v201), 0))
				_ = _lhs130
				_1290_v207 = _rhs213
				(_lhs129).ArraySet1(_rhs214, (_lhs130).Int())
				_1175_v112 = _rhs215
				_1287_v204 = _rhs216
			}
			var _1294_v211 _dafny.MultiSet
			_ = _1294_v211
			_1294_v211 = _dafny.MultiSetOf((_this).F7(), (_this).F7())
			var _1295_v212 *C1
			_ = _1295_v212
			var _nw252 *C1 = New_C1_()
			_ = _nw252
			_nw252.Ctor__(((_dafny.MultiSetOf((_this).F7())).Union(_1294_v211)).Cardinality(), (func() _dafny.Int {
				if (_this).F7() {
					return (_dafny.Zero).Minus((_this).F5())
				}
				return p0
			})())
			_1295_v212 = _nw252
			var _1296_v213 T1
			_ = _1296_v213
			var _nw253 *C2 = New_C2_()
			_ = _nw253
			_nw253.Ctor__((_this).F5(), _this.F6())
			_1296_v213 = _nw253
			_1296_v213 = _1296_v213
			var _1297_v214 bool
			_ = _1297_v214
			_1297_v214 = true
			_1297_v214 = (_this).Fm1(globalState)
			_1296_v213 = _1296_v213
		}
	}
}
func (_this *C3) M3(globalState *GlobalState) T0 {
	{
		var r0 T0 = (T0)(nil)
		_ = r0
		var _1298_v0 _dafny.Map
		_ = _1298_v0
		_1298_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
		_1298_v0 = (_1298_v0).Update((_this).F7(), (_this).F7())
		var _1299_v1 _dafny.Map
		_ = _1299_v1
		_1299_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _dafny.IntOfInt64(683))
		(_this).F6_set_((_dafny.IntOfInt64(-594)).Times((Companion_Default___.Fm10((_this).F5(), _1299_v1, (_this).F7(), _this.F6(), globalState)).Cardinality()))
		var _1300_v2 _dafny.MultiSet
		_ = _1300_v2
		_1300_v2 = _dafny.MultiSetOf((_this).F7())
		var _1301_v3 _dafny.Sequence
		_ = _1301_v3
		_1301_v3 = _dafny.SeqOf((_this).F7())
		var _1302_v4 _dafny.Set
		_ = _1302_v4
		_1302_v4 = _dafny.SetOf((_this).F7(), !((_this).F7()), (_this).F7())
		var _1303_v5 D9
		_ = _1303_v5
		_1303_v5 = Companion_D9_.Create_DC28_((_1300_v2).IsSubsetOf(_1300_v2), (_this).F5(), (_this).Fm2(_dafny.IntOfUint32((_1301_v3).Cardinality()), _1302_v4, globalState))
		_1303_v5 = Companion_D9_.Create_DC28_((_this).F7(), _this.F6(), _this.F6())
		var _1304_v6 _dafny.Set
		_ = _1304_v6
		_1304_v6 = _dafny.SetOf((_1298_v0).Cardinality())
		var _1305_v7 _dafny.Sequence
		_ = _1305_v7
		_1305_v7 = _dafny.UnicodeSeqOfUtf8Bytes("exw")
		var _1306_v8 _dafny.MultiSet
		_ = _1306_v8
		_1306_v8 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-957))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_1307_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), _1305_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(547))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg62 _dafny.Int) interface{} {
				return coer62(arg62)
			}
		}(func(_1308_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})), _1305_v7)
		var _1309_v9 _dafny.Map
		_ = _1309_v9
		_1309_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_this).F7())
		var _1310_v10 _dafny.Sequence
		_ = _1310_v10
		_1310_v10 = _dafny.SeqOf(_dafny.SetOf((_dafny.Zero).Minus(_this.F6())), (_1304_v6).Union(_dafny.SetOf((_1306_v8).Cardinality())), (_dafny.SetOf(_this.F6(), (_this).F5(), (_1309_v9).Cardinality())).Intersection(_1304_v6))
		var _1311_v11 bool
		_ = _1311_v11
		_1311_v11 = true
		var _1312_v12 _dafny.Set
		_ = _1312_v12
		_1312_v12 = _dafny.SetOf(Companion_Default___.Fm18(!((_this).F7()), _this.F6(), (_this).F5(), _1311_v11, globalState))
		var _1313_v13 _dafny.CodePoint
		_ = _1313_v13
		_1313_v13 = _dafny.CodePoint('e')
		var _1314_v14 T0
		_ = _1314_v14
		var _nw254 *C2 = New_C2_()
		_ = _nw254
		_nw254.Ctor__(_this.F6(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(43))).Cardinality()))
		_1314_v14 = _nw254
		var _1315_v15 _dafny.Sequence
		_ = _1315_v15
		_1315_v15 = _dafny.SeqOf(_1314_v14)
		var _1316_v16 _dafny.Map
		_ = _1316_v16
		_1316_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1313_v13, (_dafny.MultiSetFromSeq(_1315_v15)).Cardinality())
		var _rhs217 _dafny.Int = (_this).F5()
		_ = _rhs217
		var _rhs218 _dafny.Sequence = _1310_v10
		_ = _rhs218
		var _rhs219 bool = ((_this).F5()).Cmp((func() _dafny.Int {
			if (_1316_v16).Contains(_1313_v13) {
				return (_1316_v16).Get(_1313_v13).(_dafny.Int)
			}
			return _dafny.IntOfInt64(149)
		})()) < 0
		_ = _rhs219
		var _rhs220 _dafny.Set = _dafny.SetOf(_1313_v13)
		_ = _rhs220
		var _rhs221 bool = ((_this).F7()) || ((func() bool {
			if (_1298_v0).Contains((_this).F7()) {
				return (_1298_v0).Get((_this).F7()).(bool)
			}
			return (_this).F7()
		})())
		_ = _rhs221
		var _lhs131 *GlobalState = globalState
		_ = _lhs131
		_lhs131.F2 = _rhs217
		_1310_v10 = _rhs218
		_1311_v11 = _rhs219
		_1312_v12 = _rhs220
		_1311_v11 = _rhs221
		var _1317_v17 _dafny.Array
		_ = _1317_v17
		var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw255
		_1317_v17 = _nw255
		var _1318_v18 D5
		_ = _1318_v18
		_1318_v18 = Companion_D5_.Create_DC17_(_this.F6(), _1317_v17, (Companion_Default___.Fm31(_this.F6(), false, true, globalState)).Update(_this.F6(), (_this).F7()))
		var _source14 D5 = _1318_v18
		_ = _source14
		if _source14.Is_DC17() {
			var _1319___mcc_h0 _dafny.Int = _source14.Get_().(D5_DC17).Cf27
			_ = _1319___mcc_h0
			var _1320___mcc_h1 _dafny.Array = _source14.Get_().(D5_DC17).Cf28
			_ = _1320___mcc_h1
			var _1321___mcc_h2 _dafny.Map = _source14.Get_().(D5_DC17).Cf29
			_ = _1321___mcc_h2
			var _1322_cf29 _dafny.Map = _1321___mcc_h2
			_ = _1322_cf29
			var _1323_cf28 _dafny.Array = _1320___mcc_h1
			_ = _1323_cf28
			var _1324_cf27 _dafny.Int = _1319___mcc_h0
			_ = _1324_cf27
			(globalState).F2 = (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v11, _1311_v11), _1298_v0)).Cardinality()
			var _1325_v19 _dafny.Sequence
			_ = _1325_v19
			_1325_v19 = _dafny.SeqOf((_dafny.IntOfUint32((_1305_v7).Cardinality())).Plus(_1324_cf27))
			var _rhs222 bool = _1311_v11
			_ = _rhs222
			var _rhs223 _dafny.Int = (_1325_v19).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1324_cf27), _dafny.IntOfUint32((_1325_v19).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs223
			var _rhs224 _dafny.Int = _dafny.IntOfInt64(928)
			_ = _rhs224
			var _rhs225 bool = (_dafny.IntOfInt64(402)).Cmp(_dafny.IntOfUint32((_1301_v3).Cardinality())) != 0
			_ = _rhs225
			var _rhs226 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if _1311_v11 {
					return _1313_v13
				}
				return _1313_v13
			})(), ((_this).F5()).Minus(_this.F6()))
			_ = _rhs226
			var _lhs132 *C3 = _this
			_ = _lhs132
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			_1311_v11 = _rhs222
			_lhs132.F6_set_(_rhs223)
			_lhs133.F2 = _rhs224
			_1311_v11 = _rhs225
			_1316_v16 = _rhs226
			_1300_v2 = (_1300_v2).Intersection(_1300_v2)
			(globalState).F2 = (_this).F5()
		} else if _source14.Is_DC18() {
			var _1326___mcc_h3 bool = _source14.Get_().(D5_DC18).Cf30
			_ = _1326___mcc_h3
			var _1327___mcc_h4 _dafny.Int = _source14.Get_().(D5_DC18).Cf31
			_ = _1327___mcc_h4
			var _1328___mcc_h5 bool = _source14.Get_().(D5_DC18).Cf32
			_ = _1328___mcc_h5
			var _1329_cf32 bool = _1328___mcc_h5
			_ = _1329_cf32
			var _1330_cf31 _dafny.Int = _1327___mcc_h4
			_ = _1330_cf31
			var _1331_cf30 bool = _1326___mcc_h3
			_ = _1331_cf30
			(globalState).F2 = (_this).F5()
			if (_this.F6()).Cmp(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(135), _this.F6())) == 0 {
				(_this).F6_set_((_this).F5())
				var _1332_v20 *C0
				_ = _1332_v20
				var _nw256 *C0 = New_C0_()
				_ = _nw256
				_nw256.Ctor__((func() _dafny.Int {
					if (_this).F7() {
						return _this.F6()
					}
					return _this.F6()
				})())
				_1332_v20 = _nw256
				var _1333_v21 *C2
				_ = _1333_v21
				var _nw257 *C2 = New_C2_()
				_ = _nw257
				_nw257.Ctor__(_dafny.IntOfInt64(868), (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(107))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1334_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1335_i2 _dafny.Int) _dafny.CodePoint {
						return _1334_v13
					}
				})(_1313_v13))))).Cardinality())
				_1333_v21 = _nw257
				var _1336_v22 _dafny.Sequence
				_ = _1336_v22
				_1336_v22 = _dafny.SeqOf(_1333_v21, _1333_v21, _1333_v21)
				_1333_v21 = (_1336_v22).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(495), (_1332_v20).F8())), _dafny.IntOfUint32((_1336_v22).Cardinality()))).Uint32()).(*C2)
				var _1337_v23 _dafny.Array
				_ = _1337_v23
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_28
				var _nw258 _dafny.Array
				_ = _nw258
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw258 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) bool = (func(_1338_cf30 bool) func(_dafny.Int) bool {
						return func(_1339_i3 _dafny.Int) bool {
							return _1338_cf30
						}
					})(_1331_cf30)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw258 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw258).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw258).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1337_v23 = _nw258
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1337_v23), 0))
				_ = _index231
				(_1337_v23).ArraySet1(_1329_cf32, (_index231).Int())
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1337_v23), 0))
				_ = _index232
				(_1337_v23).ArraySet1((_1304_v6).IsSubsetOf(_1304_v6), (_index232).Int())
				var _1340_v24 _dafny.Map
				_ = _1340_v24
				_1340_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1337_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1337_v23), 0))).Int()).(bool)), Companion_Default___.SafeDivisionInt(_1330_cf31, (_this).Fm2((_1332_v20).F8(), _1302_v4, globalState)))
				var _1341_v25 _dafny.Array
				_ = _1341_v25
				var _nw259 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
				_ = _nw259
				_1341_v25 = _nw259
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1341_v25), 0))
				_ = _index233
				(_1341_v25).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ucngrye"), (_index233).Int())
				var _1342_v26 _dafny.Array
				_ = _1342_v26
				var _nw260 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
				_ = _nw260
				_1342_v26 = _nw260
				var _1343_v27 _dafny.Map
				_ = _1343_v27
				_1343_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _1342_v26)
				var _1344_v28 _dafny.Array
				_ = _1344_v28
				var _nwElement0_57 _dafny.Array = (func() _dafny.Array {
					if (_1343_v27).Contains(_1330_cf31) {
						return (_1343_v27).Get(_1330_cf31).(_dafny.Array)
					}
					return _1342_v26
				})()
				_ = _nwElement0_57
				var _nw261 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(4))
				_ = _nw261
				(_nw261).ArraySet1(_nwElement0_57, 0)
				(_nw261).ArraySet1(_1342_v26, 1)
				(_nw261).ArraySet1(_1342_v26, 2)
				(_nw261).ArraySet1(_1342_v26, 3)
				_1344_v28 = _nw261
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1344_v28), 0))
				_ = _index234
				(_1344_v28).ArraySet1(_1342_v26, (_index234).Int())
				var _1345_v29 D1
				_ = _1345_v29
				_1345_v29 = Companion_D1_.Create_DC5_()
				var _1346_v30 D5
				_ = _1346_v30
				_1346_v30 = Companion_D5_.Create_DC19_((_1332_v20).F8(), _1313_v13, (_this).F5())
				var _1347_v31 _dafny.MultiSet
				_ = _1347_v31
				_1347_v31 = _dafny.MultiSetOf(_this.F6(), _dafny.IntOfInt64(148), _1330_cf31)
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1337_v23), 0))
				_ = _index235
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1341_v25), 0))
				_ = _index236
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1344_v28), 0))
				_ = _index237
				var _rhs227 bool = (_1333_v21).Fm9(_1345_v29, ((_this).F5()).Cmp((_1304_v6).Cardinality()) >= 0, globalState)
				_ = _rhs227
				var _rhs228 _dafny.CodePoint = (_1346_v30).Dtor_cf34()
				_ = _rhs228
				var _rhs229 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1331_cf30, (_dafny.IntOfInt64(603)).Plus(_this.F6()))
				_ = _rhs229
				var _rhs230 _dafny.Sequence = Companion_Default___.Fm11(_1330_cf31, (func() _dafny.Int {
					if (_1347_v31).Contains((_this).F5()) {
						return (_1347_v31).Multiplicity((_this).F5())
					}
					return (_1332_v20).F8()
				})(), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_1298_v0).Cardinality()), (_1332_v20).Fm4(_1313_v13, _1300_v2, (_this).F7(), globalState)), Companion_Default___.SafeModuloInt(_1330_cf31, _this.F6()), globalState)
				_ = _rhs230
				var _rhs231 _dafny.Array = _1342_v26
				_ = _rhs231
				var _lhs134 _dafny.Array = _1337_v23
				_ = _lhs134
				var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(623), _dafny.ArrayLen((_1337_v23), 0))
				_ = _lhs135
				var _lhs136 _dafny.Array = _1341_v25
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_1341_v25), 0))
				_ = _lhs137
				var _lhs138 _dafny.Array = _1344_v28
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1344_v28), 0))
				_ = _lhs139
				(_lhs134).ArraySet1(_rhs227, (_lhs135).Int())
				_1313_v13 = _rhs228
				_1340_v24 = _rhs229
				(_lhs136).ArraySet1(_rhs230, (_lhs137).Int())
				(_lhs138).ArraySet1(_rhs231, (_lhs139).Int())
			} else {
				_1313_v13 = _1313_v13
				var _1348_v32 _dafny.Array
				_ = _1348_v32
				var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw262
				_1348_v32 = _nw262
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1348_v32), 0))
				_ = _index238
				(_1348_v32).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1305_v7, _1305_v7), (_index238).Int())
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1348_v32), 0))
				_ = _index239
				(_1348_v32).ArraySet1(_1305_v7, (_index239).Int())
				var _1349_v33 _dafny.Array
				_ = _1349_v33
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_29
				var _nw263 _dafny.Array
				_ = _nw263
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw263 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) bool = (func(_1350_v7 _dafny.Sequence, _1351_v6 _dafny.Set, _1352_cf32 bool, _1353_v2 _dafny.MultiSet, _1354_v32 _dafny.Array) func(_dafny.Int) bool {
						return func(_1355_i4 _dafny.Int) bool {
							return (_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D7_.Create_DC23_((func() _dafny.Int {
								if (_1353_v2).Contains(_1352_cf32) {
									return (_1353_v2).Multiplicity(_1352_cf32)
								}
								return (_this).F5()
							})(), (_1354_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1354_v32), 0))).Int()).(_dafny.Sequence), (_this).F5())))).IsSubsetOf(_dafny.MultiSetOf(Companion_D7_.Create_DC23_(_this.F6(), _1350_v7, (_1351_v6).Cardinality())))
						}
					})(_1305_v7, _1304_v6, _1329_cf32, _1300_v2, _1348_v32)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw263 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw263).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw263).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1349_v33 = _nw263
				_1349_v33 = _1349_v33
				var _1356_v34 _dafny.Sequence
				_ = _1356_v34
				_1356_v34 = _dafny.SeqOf((Companion_D1_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("htofcuax"), _this.F6())).Dtor_cf5(), (_1348_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1348_v32), 0))).Int()).(_dafny.Sequence), (_1348_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_1348_v32), 0))).Int()).(_dafny.Sequence))
				var _1357_v35 *C2
				_ = _1357_v35
				var _nw264 *C2 = New_C2_()
				_ = _nw264
				_nw264.Ctor__((_this).F5(), _dafny.IntOfUint32((_1356_v34).Cardinality()))
				_1357_v35 = _nw264
				var _1358_v36 *C2
				_ = _1358_v36
				var _nw265 *C2 = New_C2_()
				_ = _nw265
				_nw265.Ctor__(_dafny.IntOfInt64(-392), (_1299_v1).Cardinality())
				_1358_v36 = _nw265
			}
			var _1359_v37 _dafny.Set
			_ = _1359_v37
			_1359_v37 = _dafny.SetOf(_1300_v2)
			if (_1359_v37).IsProperSubsetOf(_1359_v37) {
				_1329_cf32 = _1311_v11
				var _1360_v38 D4
				_ = _1360_v38
				_1360_v38 = Companion_D4_.Create_DC13_(_1329_cf32, (_this).F5(), _1329_cf32)
				var _rhs232 bool = true
				_ = _rhs232
				var _rhs233 bool = _1331_cf30
				_ = _rhs233
				var _rhs234 bool = !(((_this).F5()).Cmp((_1360_v38).Dtor_cf18()) < 0)
				_ = _rhs234
				_1311_v11 = _rhs232
				_1311_v11 = _rhs233
				_1331_cf30 = _rhs234
				var _1361_v39 D12
				_ = _1361_v39
				_1361_v39 = Companion_D12_.Create_DC36_(_1311_v11, _1331_cf30, _1305_v7)
				var _1362_v40 T1
				_ = _1362_v40
				var _nw266 *C2 = New_C2_()
				_ = _nw266
				_nw266.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _1330_cf31)).Cardinality(), (_1309_v9).Cardinality())
				_1362_v40 = _nw266
				_1309_v9 = (_1309_v9).Update(Companion_Default___.SafeDivisionInt(_this.F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1361_v39, _1362_v40)).Cardinality()), (_this).F7())
				var _1363_v41 _dafny.Array
				_ = _1363_v41
				var _nw267 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw267
				_1363_v41 = _nw267
				var _1364_v42 _dafny.Map
				_ = _1364_v42
				_1364_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1362_v40, _1363_v41)
				_1364_v42 = _1364_v42
				(globalState).F2 = (_this).F5()
			} else {
				var _1365_v43 _dafny.MultiSet
				_ = _1365_v43
				_1365_v43 = _dafny.MultiSetOf((_this).F5())
				var _1366_v44 _dafny.Array
				_ = _1366_v44
				var _nwElement0_58 _dafny.MultiSet = _1365_v43
				_ = _nwElement0_58
				var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(9))
				_ = _nw268
				(_nw268).ArraySet1(_nwElement0_58, 0)
				(_nw268).ArraySet1(_1365_v43, 1)
				(_nw268).ArraySet1(_1365_v43, 2)
				(_nw268).ArraySet1((_1365_v43).Update((_1365_v43).Cardinality(), Companion_Default___.Abs((_this).F5())), 3)
				(_nw268).ArraySet1((_1365_v43).Intersection(_1365_v43), 4)
				(_nw268).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfInt64(-390), _this.F6(), (_this).F5()), 5)
				(_nw268).ArraySet1(_1365_v43, 6)
				(_nw268).ArraySet1((_1365_v43).Difference(_1365_v43), 7)
				(_nw268).ArraySet1((_1365_v43).Update((_this).F5(), Companion_Default___.Abs((_this).F5())), 8)
				_1366_v44 = _nw268
				_1366_v44 = _1366_v44
				var _1367_v45 *C0
				_ = _1367_v45
				var _nw269 *C0 = New_C0_()
				_ = _nw269
				_nw269.Ctor__((Companion_D4_.Create_DC13_(_1329_cf32, (_this).F5(), _1329_cf32)).Dtor_cf18())
				_1367_v45 = _nw269
				var _1368_v46 _dafny.Sequence
				_ = _1368_v46
				_1368_v46 = _dafny.SeqOf(_this.F6(), (_1367_v45).F8(), _1330_cf31, (_this).F5(), (_1367_v45).F8())
				_1368_v46 = _1368_v46
				var _1369_v47 *C1
				_ = _1369_v47
				var _nw270 *C1 = New_C1_()
				_ = _nw270
				_nw270.Ctor__((_1367_v45).Fm4(_1313_v13, _1300_v2, _1329_cf32, globalState), _dafny.IntOfInt64(578))
				_1369_v47 = _nw270
				_1369_v47 = _1369_v47
				var _1370_v48 _dafny.Array
				_ = _1370_v48
				var _nw271 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(18))
				_ = _nw271
				_1370_v48 = _nw271
				var _1371_v49 *C2
				_ = _1371_v49
				var _nw272 *C2 = New_C2_()
				_ = _nw272
				_nw272.Ctor__(_dafny.IntOfInt64(326), _dafny.IntOfUint32((_1301_v3).Cardinality()))
				_1371_v49 = _nw272
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1370_v48), 0))
				_ = _index240
				(_1370_v48).ArraySet1(_1371_v49, (_index240).Int())
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1370_v48), 0))
				_ = _index241
				var _rhs235 *C2 = _1371_v49
				_ = _rhs235
				var _rhs236 _dafny.Int = _this.F6()
				_ = _rhs236
				var _rhs237 *C1 = _1369_v47
				_ = _rhs237
				var _lhs140 _dafny.Array = _1370_v48
				_ = _lhs140
				var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1370_v48), 0))
				_ = _lhs141
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				(_lhs140).ArraySet1(_rhs235, (_lhs141).Int())
				_lhs142.F2 = _rhs236
				_1369_v47 = _rhs237
			}
			var _1372_v50 D12
			_ = _1372_v50
			_1372_v50 = Companion_D12_.Create_DC36_(_1311_v11, _1329_cf32, _1305_v7)
			var _1373_v51 _dafny.Sequence
			_ = _1373_v51
			_1373_v51 = _dafny.SeqOf(_1305_v7, _1305_v7, _dafny.UnicodeSeqOfUtf8Bytes("ey"))
			var _1374_v52 _dafny.Array
			_ = _1374_v52
			var _nwElement0_59 _dafny.Sequence = _1305_v7
			_ = _nwElement0_59
			var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(18))
			_ = _nw273
			(_nw273).ArraySet1(_nwElement0_59, 0)
			(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11((_this).F5(), _this.F6(), _1330_cf31, _this.F6(), globalState), _1305_v7), 1)
			(_nw273).ArraySet1((_1372_v50).Dtor_cf62(), 2)
			(_nw273).ArraySet1(_dafny.Companion_Sequence_.Update(_1305_v7, (Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_1305_v7).Cardinality()))).Uint32(), _1313_v13), 3)
			(_nw273).ArraySet1(_1305_v7, 4)
			(_nw273).ArraySet1(_1305_v7, 5)
			(_nw273).ArraySet1(_1305_v7, 6)
			(_nw273).ArraySet1(_1305_v7, 7)
			(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1305_v7, _dafny.UnicodeSeqOfUtf8Bytes("iaddku")), 8)
			(_nw273).ArraySet1(_1305_v7, 9)
			(_nw273).ArraySet1(_1305_v7, 10)
			(_nw273).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qpastef"), 11)
			(_nw273).ArraySet1(_1305_v7, 12)
			(_nw273).ArraySet1(_1305_v7, 13)
			(_nw273).ArraySet1(_1305_v7, 14)
			(_nw273).ArraySet1((_1373_v51).Select((Companion_Default___.SafeIndex(_1330_cf31, _dafny.IntOfUint32((_1373_v51).Cardinality()))).Uint32()).(_dafny.Sequence), 15)
			(_nw273).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(5))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}((func(_1375_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1376_i5 _dafny.Int) _dafny.CodePoint {
					return _1375_v13
				}
			})(_1313_v13))), 16)
			(_nw273).ArraySet1(_1305_v7, 17)
			_1374_v52 = _nw273
			var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw274
			_1374_v52 = _nw274
		} else if _source14.Is_DC19() {
			var _1377___mcc_h6 _dafny.Int = _source14.Get_().(D5_DC19).Cf33
			_ = _1377___mcc_h6
			var _1378___mcc_h7 _dafny.CodePoint = _source14.Get_().(D5_DC19).Cf34
			_ = _1378___mcc_h7
			var _1379___mcc_h8 _dafny.Int = _source14.Get_().(D5_DC19).Cf35
			_ = _1379___mcc_h8
			var _1380_cf35 _dafny.Int = _1379___mcc_h8
			_ = _1380_cf35
			var _1381_cf34 _dafny.CodePoint = _1378___mcc_h7
			_ = _1381_cf34
			var _1382_cf33 _dafny.Int = _1377___mcc_h6
			_ = _1382_cf33
			_1382_cf33 = _1382_cf33
			var _1383_v53 _dafny.Map
			_ = _1383_v53
			_1383_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v11, (_1298_v0).Merge(_1298_v0))
			_1383_v53 = Companion_Default___.Fm32((_this).F5(), !(!(true)) || ((_this).F7()), globalState)
			var _1384_v54 *C0
			_ = _1384_v54
			var _nw275 *C0 = New_C0_()
			_ = _nw275
			_nw275.Ctor__(_1382_cf33)
			_1384_v54 = _nw275
			_1384_v54 = _1384_v54
			var _1385_v55 _dafny.Array
			_ = _1385_v55
			var _len0_30 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_30
			var _nw276 _dafny.Array
			_ = _nw276
			if _len0_30.Cmp(_dafny.Zero) == 0 {
				_nw276 = _dafny.NewArray(_len0_30)
			} else {
				var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1386_v0 _dafny.Map) func(_dafny.Int) _dafny.Sequence {
					return func(_1387_i6 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf((Companion_D0_.Create_DC3_((Companion_D4_.Create_DC14_((_1386_v0).Cardinality(), (_this).F7(), _dafny.IntOfInt64(675), (_this).F7(), _this.F6())).Dtor_cf21())).Dtor_cf3())
					}
				})(_1298_v0)
				_ = _init30
				var _element0_30 = _init30(_dafny.Zero)
				_ = _element0_30
				_nw276 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
				(_nw276).ArraySet1(_element0_30, 0)
				var _nativeLen0_30 = (_len0_30).Int()
				_ = _nativeLen0_30
				for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
					(_nw276).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
				}
			}
			_1385_v55 = _nw276
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_1385_v55), 0))
			_ = _index242
			(_1385_v55).ArraySet1(_dafny.SeqOf(_1311_v11), (_index242).Int())
			var _1388_v56 _dafny.Sequence
			_ = _1388_v56
			_1388_v56 = _dafny.SeqOf(_1301_v3, _dafny.SeqOf(true, (_this).F7()), _1301_v3, _1301_v3, _1301_v3)
			var _1389_v57 _dafny.Sequence
			_ = _1389_v57
			_1389_v57 = _dafny.SeqOf((_1388_v56).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.IntOfUint32((_1388_v56).Cardinality()))).Uint32()).(_dafny.Sequence), _1301_v3)
			var _1390_v58 _dafny.Array
			_ = _1390_v58
			var _nw277 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw277
			_1390_v58 = _nw277
			var _1391_v59 _dafny.MultiSet
			_ = _1391_v59
			_1391_v59 = _dafny.MultiSetOf(_1390_v58)
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_1385_v55), 0))
			_ = _index243
			(_1385_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_1388_v56).Select((Companion_Default___.SafeIndex((_1391_v59).Cardinality(), _dafny.IntOfUint32((_1388_v56).Cardinality()))).Uint32()).(_dafny.Sequence), _1301_v3), Companion_Default___.Fm33(globalState)), (_index243).Int())
		} else {
			var _1392___mcc_h9 _dafny.Array = _source14.Get_().(D5_DC16).Cf26
			_ = _1392___mcc_h9
			var _1393_cf26 _dafny.Array = _1392___mcc_h9
			_ = _1393_cf26
			_1313_v13 = _1313_v13
			var _1394_v60 _dafny.Array
			_ = _1394_v60
			var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(18))
			_ = _nw278
			_1394_v60 = _nw278
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_1394_v60), 0))
			_ = _index244
			(_1394_v60).ArraySet1(_dafny.SeqOf((_this).F7()), (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_1394_v60), 0))
			_ = _index245
			(_1394_v60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1301_v3, Companion_Default___.Fm33(globalState)), (_index245).Int())
			var _1395_v61 _dafny.Map
			_ = _1395_v61
			_1395_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F6(), _1305_v7)
			_1311_v11 = !(_1395_v61).Equals(_1395_v61)
			var _1396_v62 _dafny.Sequence
			_ = _1396_v62
			_1396_v62 = _dafny.SeqOf((_this).F5(), _dafny.IntOfInt64(-235))
			var _1397_v63 D4
			_ = _1397_v63
			_1397_v63 = Companion_D4_.Create_DC11_(_1396_v62)
			var _pat_let_tv17 = _1396_v62
			_ = _pat_let_tv17
			var _1398_v64 _dafny.Array
			_ = _1398_v64
			var _nwElement0_60 D4 = func(_pat_let29_0 D4) D4 {
				return func(_1399_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let30_0 _dafny.Sequence) D4 {
						return func(_1400_dt__update_hcf13_h0 _dafny.Sequence) D4 {
							return Companion_D4_.Create_DC11_(_1400_dt__update_hcf13_h0)
						}(_pat_let30_0)
					}(_pat_let_tv17)
				}(_pat_let29_0)
			}(_1397_v63)
			_ = _nwElement0_60
			var _nw279 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.One)
			_ = _nw279
			(_nw279).ArraySet1(_nwElement0_60, 0)
			_1398_v64 = _nw279
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1398_v64), 0))
			_ = _index246
			(_1398_v64).ArraySet1(_1397_v63, (_index246).Int())
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_1398_v64), 0))
			_ = _index247
			(_1398_v64).ArraySet1(Companion_Default___.Fm27(_this.F6(), globalState), (_index247).Int())
		}
		if !((_1302_v4).IsDisjointFrom(_dafny.SetOf(_1311_v11, _1311_v11, (_this).F7(), !((_this).F7())))) {
			var _1401_v65 _dafny.MultiSet
			_ = _1401_v65
			_1401_v65 = _dafny.MultiSetOf(_1309_v9)
			var _1402_v66 _dafny.Sequence
			_ = _1402_v66
			_1402_v66 = _dafny.SeqOf(_1401_v65)
			if (_this).Fm0((_1402_v66).Select((Companion_Default___.SafeIndex(_this.F6(), _dafny.IntOfUint32((_1402_v66).Cardinality()))).Uint32()).(_dafny.MultiSet), (_1300_v2).IsProperSubsetOf(_1300_v2), _dafny.Companion_Sequence_.Concatenate(_1305_v7, _1305_v7), globalState) {
				var _1403_v67 _dafny.Map
				_ = _1403_v67
				_1403_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(58))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}(func(_1404_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), Companion_Default___.SafeModuloInt((_this).F5(), (_this).F5()))
				var _1405_v68 _dafny.Array
				_ = _1405_v68
				var _nw280 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw280
				_1405_v68 = _nw280
				var _1406_v69 _dafny.MultiSet
				_ = _1406_v69
				_1406_v69 = _dafny.MultiSetOf(_1405_v68)
				var _1407_v70 _dafny.Map
				_ = _1407_v70
				_1407_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v14, _1403_v67)
				var _rhs238 _dafny.Map = (func() _dafny.Map {
					if (_1407_v70).Contains(_1314_v14) {
						return (_1407_v70).Get(_1314_v14).(_dafny.Map)
					}
					return (_1403_v67).Merge(_1403_v67)
				})()
				_ = _rhs238
				var _rhs239 _dafny.MultiSet = _1406_v69
				_ = _rhs239
				_1403_v67 = _rhs238
				_1406_v69 = _rhs239
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_1405_v68), 0))
				_ = _index248
				(_1405_v68).ArraySet1(_1311_v11, (_index248).Int())
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1405_v68), 0))
				_ = _index249
				(_1405_v68).ArraySet1((_this.F6()).Cmp(_this.F6()) <= 0, (_index249).Int())
				var _1408_v71 D0
				_ = _1408_v71
				_1408_v71 = Companion_D0_.Create_DC3_((_this).F7())
				var _1409_v73 _dafny.Map
				_ = _1409_v73
				_1409_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v11, _this.F6())
				var _1410_v74 _dafny.Sequence
				_ = _1410_v74
				_1410_v74 = _dafny.SeqOf((_dafny.IntOfInt64(687)).Minus((_this).F5()), Companion_Default___.SafeDivisionInt(_this.F6(), _this.F6()), (_this.F6()).Plus(_this.F6()), (_this).F5(), (_this).Fm2(_this.F6(), _1302_v4, globalState))
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_1405_v68), 0))
				_ = _index250
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1405_v68), 0))
				_ = _index251
				var _rhs240 bool = (_1408_v71).Dtor_cf3()
				_ = _rhs240
				var _rhs241 _dafny.Map = func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-393), _dafny.IntOfInt64(743))); ; {
						_compr_55, _ok61 := _iter61()
						if !_ok61 {
							break
						}
						var _1411_v72 _dafny.Int
						_1411_v72 = interface{}(_compr_55).(_dafny.Int)
						if ((_dafny.IntOfInt64(-393)).Cmp(_1411_v72) <= 0) && ((_1411_v72).Cmp(_dafny.IntOfInt64(743)) < 0) {
							_coll55.Add(Companion_Default___.SafeModuloInt(_1411_v72, _dafny.IntOfInt64(513)), _this.F6())
						}
					}
					return _coll55.ToMap()
				}()
				_ = _rhs241
				var _rhs242 _dafny.Int = ((_this.F6()).Minus((func() _dafny.Int {
					if (_1409_v73).Contains(_1311_v11) {
						return (_1409_v73).Get(_1311_v11).(_dafny.Int)
					}
					return (_this).F5()
				})())).Times((_dafny.Zero).Minus(_this.F6()))
				_ = _rhs242
				var _rhs243 _dafny.Int = (_1410_v74).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1410_v74).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs243
				var _rhs244 bool = (_1301_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1305_v7, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("esd"), (Companion_Default___.SafeIndex((_this).Fm2(_this.F6(), _1302_v4, globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("esd")).Cardinality()))).Uint32(), _1313_v13))).Cardinality()), _dafny.IntOfUint32((_1301_v3).Cardinality()))).Uint32()).(bool)
				_ = _rhs244
				var _lhs143 _dafny.Array = _1405_v68
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_1405_v68), 0))
				_ = _lhs144
				var _lhs145 *GlobalState = globalState
				_ = _lhs145
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				var _lhs147 _dafny.Array = _1405_v68
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(44), _dafny.ArrayLen((_1405_v68), 0))
				_ = _lhs148
				(_lhs143).ArraySet1(_rhs240, (_lhs144).Int())
				_1299_v1 = _rhs241
				_lhs145.F2 = _rhs242
				_lhs146.F2 = _rhs243
				(_lhs147).ArraySet1(_rhs244, (_lhs148).Int())
				(globalState).F2 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
					if (_1405_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_1405_v68), 0))).Int()).(bool) {
						return (_this).F5()
					}
					return (_dafny.Zero).Minus(_this.F6())
				})(), (_this).Fm2((_this).F5(), _1302_v4, globalState)))
				_1309_v9 = (_1309_v9).Update((_this).F5(), (_this).Fm0(_1401_v65, (_this).F7(), _1305_v7, globalState))
				_1305_v7 = Companion_Default___.Fm7(_1313_v13, (_1405_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_1405_v68), 0))).Int()).(bool), _1311_v11, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1410_v74).Cardinality())), globalState)
			} else {
				var _1412_v75 _dafny.Array
				_ = _1412_v75
				var _nw281 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
				_ = _nw281
				_1412_v75 = _nw281
				var _1413_v76 _dafny.Array
				_ = _1413_v76
				var _nwElement0_61 _dafny.Array = _1412_v75
				_ = _nwElement0_61
				var _nw282 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(2))
				_ = _nw282
				(_nw282).ArraySet1(_nwElement0_61, 0)
				(_nw282).ArraySet1(_1412_v75, 1)
				_1413_v76 = _nw282
				var _1414_v77 _dafny.Array
				_ = _1414_v77
				var _nwElement0_62 _dafny.Array = _1413_v76
				_ = _nwElement0_62
				var _nw283 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.One)
				_ = _nw283
				(_nw283).ArraySet1(_nwElement0_62, 0)
				_1414_v77 = _nw283
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1414_v77), 0))
				_ = _index252
				(_1414_v77).ArraySet1(_1413_v76, (_index252).Int())
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_1414_v77), 0))
				_ = _index253
				var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw284
				(_1414_v77).ArraySet1(_nw284, (_index253).Int())
				(globalState).F2 = _dafny.IntOfInt64(-132)
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_1317_v17), 0))
				_ = _index254
				(_1317_v17).ArraySet1((_this.F6()).Minus(_this.F6()), (_index254).Int())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_1317_v17), 0))
				_ = _index255
				(_1317_v17).ArraySet1((_this).F5(), (_index255).Int())
				var _1415_v78 *C0
				_ = _1415_v78
				var _nw285 *C0 = New_C0_()
				_ = _nw285
				_nw285.Ctor__(((_this).F5()).Plus(_dafny.IntOfUint32((_1305_v7).Cardinality())))
				_1415_v78 = _nw285
				_1313_v13 = _1313_v13
			}
			if _1311_v11 {
				var _1416_v79 D10
				_ = _1416_v79
				_1416_v79 = Companion_D10_.Create_DC30_((_this).F5())
				(globalState).F2 = (_1416_v79).Dtor_cf50()
				_1311_v11 = (Companion_D12_.Create_DC34_((_this).F7())).Dtor_cf56()
				var _1417_v80 _dafny.Array
				_ = _1417_v80
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_31
				var _nw286 _dafny.Array
				_ = _nw286
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw286 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Map = (func(_1418_v11 bool, _1419_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
						return func(_1420_i8 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1418_v11, Companion_D7_.Create_DC23_(_this.F6(), _1419_v7, (_this).F5())), _dafny.CodePoint('c'))
						}
					})(_1311_v11, _1305_v7)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw286 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw286).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw286).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1417_v80 = _nw286
				var _1421_v81 D12
				_ = _1421_v81
				_1421_v81 = Companion_D12_.Create_DC36_((func() bool {
					if (_1298_v0).Contains(_1311_v11) {
						return (_1298_v0).Get(_1311_v11).(bool)
					}
					return (_this).F7()
				})(), _1311_v11, _1305_v7)
				var _1422_v82 D7
				_ = _1422_v82
				_1422_v82 = Companion_D7_.Create_DC23_((_this).F5(), _1305_v7, _dafny.IntOfInt64(845))
				var _1423_v83 _dafny.Map
				_ = _1423_v83
				_1423_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1421_v81).Dtor_cf60(), _1422_v82), _1313_v13)
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1417_v80), 0))
				_ = _index256
				(_1417_v80).ArraySet1((_1423_v83).Merge(_1423_v83), (_index256).Int())
				var _1424_v84 _dafny.Sequence
				_ = _1424_v84
				_1424_v84 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1301_v3, _1301_v3), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F7()), (Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality()))).Uint32(), !((_this).F7())), _1301_v3), _dafny.Companion_Sequence_.Concatenate(_1301_v3, _dafny.Companion_Sequence_.Update(_1301_v3, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bsnlo")).Cardinality()), _dafny.IntOfUint32((_1301_v3).Cardinality()))).Uint32(), false)))
				var _1425_v85 _dafny.Array
				_ = _1425_v85
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_32
				var _nw287 _dafny.Array
				_ = _nw287
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw287 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.CodePoint = (func(_1426_v13 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1427_i9 _dafny.Int) _dafny.CodePoint {
							return (func() _dafny.CodePoint {
								if (_this).F7() {
									return _1426_v13
								}
								return _dafny.CodePoint('u')
							})()
						}
					})(_1313_v13)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw287 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw287).ArraySet1CodePoint(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw287).ArraySet1CodePoint(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1425_v85 = _nw287
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_1425_v85), 0))
				_ = _index257
				(_1425_v85).ArraySet1CodePoint(_1313_v13, (_index257).Int())
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))
				_ = _index258
				(_1317_v17).ArraySet1((_this).F5(), (_index258).Int())
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1417_v80), 0))
				_ = _index259
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_1425_v85), 0))
				_ = _index260
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))
				_ = _index261
				var _rhs245 _dafny.CodePoint = _1313_v13
				_ = _rhs245
				var _rhs246 _dafny.Map = _1423_v83
				_ = _rhs246
				var _rhs247 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1424_v84, (Companion_Default___.SafeIndex((_1300_v2).Cardinality(), _dafny.IntOfUint32((_1424_v84).Cardinality()))).Uint32(), _1301_v3)
				_ = _rhs247
				var _rhs248 _dafny.CodePoint = Companion_Default___.Fm18((_1311_v11) || (_1311_v11), (_this.F6()).Plus((_dafny.Zero).Minus((_1304_v6).Cardinality())), (_this).F5(), _1311_v11, globalState)
				_ = _rhs248
				var _rhs249 _dafny.Int = Companion_Default___.SafeModuloInt((_this.F6()).Plus(_dafny.IntOfInt64(580)), Companion_Default___.SafeModuloInt((_1300_v2).Cardinality(), _this.F6()))
				_ = _rhs249
				var _lhs149 _dafny.Array = _1417_v80
				_ = _lhs149
				var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_1417_v80), 0))
				_ = _lhs150
				var _lhs151 _dafny.Array = _1425_v85
				_ = _lhs151
				var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(992), _dafny.ArrayLen((_1425_v85), 0))
				_ = _lhs152
				var _lhs153 _dafny.Array = _1317_v17
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))
				_ = _lhs154
				_1313_v13 = _rhs245
				(_lhs149).ArraySet1(_rhs246, (_lhs150).Int())
				_1424_v84 = _rhs247
				(_lhs151).ArraySet1CodePoint(_rhs248, (_lhs152).Int())
				(_lhs153).ArraySet1(_rhs249, (_lhs154).Int())
				var _1428_v86 _dafny.Sequence
				_ = _1428_v86
				_1428_v86 = _dafny.SeqOf(_1300_v2)
				var _1429_v87 D0
				_ = _1429_v87
				_1429_v87 = Companion_D0_.Create_DC1_(!(true), (_this).F7())
				var _1430_v88 *C0
				_ = _1430_v88
				var _nw288 *C0 = New_C0_()
				_ = _nw288
				_nw288.Ctor__((_1317_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))).Int()).(_dafny.Int))
				_1430_v88 = _nw288
				var _1431_v89 _dafny.Set
				_ = _1431_v89
				_1431_v89 = _dafny.SetOf(_1430_v88)
				var _1432_v90 _dafny.Sequence
				_ = _1432_v90
				_1432_v90 = _dafny.SeqOf((_1430_v88).Fm4(_dafny.CodePoint('m'), _1300_v2, (_this).F7(), globalState), (_1317_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))).Int()).(_dafny.Int))
				var _1433_v91 _dafny.Array
				_ = _1433_v91
				var _nwElement0_63 bool = (_this).F7()
				_ = _nwElement0_63
				var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(23))
				_ = _nw289
				(_nw289).ArraySet1(_nwElement0_63, 0)
				(_nw289).ArraySet1(!(!(!((_this).F7())) || ((_this).F7())), 1)
				(_nw289).ArraySet1(_1311_v11, 2)
				(_nw289).ArraySet1(!(_1311_v11), 3)
				(_nw289).ArraySet1(_1311_v11, 4)
				(_nw289).ArraySet1(_1311_v11, 5)
				(_nw289).ArraySet1((_this).F7(), 6)
				(_nw289).ArraySet1(_1311_v11, 7)
				(_nw289).ArraySet1(_1311_v11, 8)
				(_nw289).ArraySet1(_1311_v11, 9)
				(_nw289).ArraySet1(((_1428_v86).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.IntOfUint32((_1428_v86).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_1300_v2), 10)
				(_nw289).ArraySet1((_this).F7(), 11)
				(_nw289).ArraySet1((_1301_v3).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1301_v3).Cardinality()))).Uint32()).(bool), 12)
				(_nw289).ArraySet1((_this).F7(), 13)
				(_nw289).ArraySet1(((_this).F7()) || ((func() bool {
					if (_1309_v9).Contains((_this).F5()) {
						return (_1309_v9).Get((_this).F5()).(bool)
					}
					return _1311_v11
				})()), 14)
				(_nw289).ArraySet1(!(false) || (_1311_v11), 15)
				(_nw289).ArraySet1(!(true), 16)
				(_nw289).ArraySet1((_1429_v87).Dtor_cf1(), 17)
				(_nw289).ArraySet1((_1431_v89).IsProperSubsetOf(_1431_v89), 18)
				(_nw289).ArraySet1(((_1310_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(759), _dafny.IntOfUint32((_1310_v10).Cardinality()))).Uint32()).(_dafny.Set)).IsSubsetOf(Companion_Default___.Fm34(_1432_v90, globalState)), 19)
				(_nw289).ArraySet1(!(!((_this).Fm1(globalState))), 20)
				(_nw289).ArraySet1(false, 21)
				(_nw289).ArraySet1((_this).F7(), 22)
				_1433_v91 = _nw289
				_1433_v91 = _1433_v91
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(304), _dafny.ArrayLen((_1317_v17), 0))
				_ = _index262
				(_1317_v17).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(29), (_1430_v88).F8()), (_index262).Int())
			} else {
				(globalState).F2 = ((_this).Fm2((_this).F5(), _1302_v4, globalState)).Times(_this.F6())
				var _1434_v92 _dafny.Array
				_ = _1434_v92
				var _nwElement0_64 _dafny.Array = _1317_v17
				_ = _nwElement0_64
				var _nw290 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(7))
				_ = _nw290
				(_nw290).ArraySet1(_nwElement0_64, 0)
				(_nw290).ArraySet1(_1317_v17, 1)
				(_nw290).ArraySet1(_1317_v17, 2)
				(_nw290).ArraySet1(_1317_v17, 3)
				(_nw290).ArraySet1(_1317_v17, 4)
				(_nw290).ArraySet1(_1317_v17, 5)
				(_nw290).ArraySet1(_1317_v17, 6)
				_1434_v92 = _nw290
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1434_v92), 0))
				_ = _index263
				(_1434_v92).ArraySet1(_1317_v17, (_index263).Int())
				var _1435_v93 D5
				_ = _1435_v93
				_1435_v93 = Companion_D5_.Create_DC16_(_1317_v17)
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_1434_v92), 0))
				_ = _index264
				(_1434_v92).ArraySet1((_1435_v93).Dtor_cf26(), (_index264).Int())
				_1311_v11 = ((_this).F7()) && (((_this).F5()).Cmp((_this).F5()) != 0)
				var _1436_v94 _dafny.MultiSet
				_ = _1436_v94
				_1436_v94 = _dafny.MultiSetOf(_dafny.IntOfInt64(207), (_this).F5())
				var _1437_v95 D0
				_ = _1437_v95
				_1437_v95 = Companion_D0_.Create_DC2_()
				var _1438_v96 D3
				_ = _1438_v96
				_1438_v96 = Companion_D3_.Create_DC9_(_1436_v94, _1437_v95, _this.F6())
				var _1439_v97 _dafny.Map
				_ = _1439_v97
				_1439_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1311_v11, _1438_v96)
				var _1440_v98 D3
				_ = _1440_v98
				_1440_v98 = Companion_D3_.Create_DC10_((func() D3 {
					if (_1439_v97).Contains(_1311_v11) {
						return (_1439_v97).Get(_1311_v11).(D3)
					}
					return _1438_v96
				})())
				var _1441_v99 _dafny.Map
				_ = _1441_v99
				_1441_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _1438_v96)
				var _1442_v100 _dafny.Sequence
				_ = _1442_v100
				_1442_v100 = _dafny.SeqOf((_this).F5())
				var _1443_v101 _dafny.Set
				_ = _1443_v101
				_1443_v101 = _dafny.SetOf(_1440_v98, Companion_D3_.Create_DC10_((func() D3 {
					if (_1441_v99).Contains(_dafny.IntOfUint32((_1442_v100).Cardinality())) {
						return (_1441_v99).Get(_dafny.IntOfUint32((_1442_v100).Cardinality())).(D3)
					}
					return _1438_v96
				})()))
				(_this).F6_set_((_1443_v101).Cardinality())
				_1313_v13 = _1313_v13
			}
			var _1444_v102 D5
			_ = _1444_v102
			_1444_v102 = Companion_D5_.Create_DC19_(_this.F6(), _1313_v13, _this.F6())
			_1444_v102 = _1444_v102
			var _1445_v103 _dafny.Map
			_ = _1445_v103
			_1445_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _this.F6())
			_1445_v103 = (_1445_v103).Update(!(!(_1311_v11) || (_1311_v11)), (_this).F5())
			var _1446_v104 _dafny.Array
			_ = _1446_v104
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_33
			var _nw291 _dafny.Array
			_ = _nw291
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw291 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) bool = (func(_1447_v11 bool) func(_dafny.Int) bool {
					return func(_1448_i10 _dafny.Int) bool {
						return (func() bool {
							if _1447_v11 {
								return (_this).F7()
							}
							return _1447_v11
						})()
					}
				})(_1311_v11)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw291 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw291).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw291).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1446_v104 = _nw291
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1446_v104), 0))
			_ = _index265
			(_1446_v104).ArraySet1((_this).F7(), (_index265).Int())
			var _1449_v105 _dafny.Sequence
			_ = _1449_v105
			_1449_v105 = _dafny.SeqOf(_this.F6())
			var _1450_v106 *C1
			_ = _1450_v106
			var _nw292 *C1 = New_C1_()
			_ = _nw292
			_nw292.Ctor__(_dafny.IntOfUint32((_1305_v7).Cardinality()), (func() _dafny.Int {
				if (_this).F7() {
					return (_this).F5()
				}
				return _dafny.IntOfUint32((_1449_v105).Cardinality())
			})())
			_1450_v106 = _nw292
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1446_v104), 0))
			_ = _index266
			var _rhs250 bool = _1311_v11
			_ = _rhs250
			var _rhs251 *C1 = _1450_v106
			_ = _rhs251
			var _rhs252 bool = _1311_v11
			_ = _rhs252
			var _lhs155 _dafny.Array = _1446_v104
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_1446_v104), 0))
			_ = _lhs156
			(_lhs155).ArraySet1(_rhs250, (_lhs156).Int())
			_1450_v106 = _rhs251
			_1311_v11 = _rhs252
		} else {
			var _1451_v107 _dafny.Array
			_ = _1451_v107
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_34
			var _nw293 _dafny.Array
			_ = _nw293
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw293 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) bool = (func(_1452_v11 bool) func(_dafny.Int) bool {
					return func(_1453_i11 _dafny.Int) bool {
						return _1452_v11
					}
				})(_1311_v11)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw293 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw293).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw293).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1451_v107 = _nw293
			var _1454_v108 _dafny.Array
			_ = _1454_v108
			var _nwElement0_65 _dafny.Array = (func() _dafny.Array {
				if (_this).F7() {
					return _1451_v107
				}
				return _1451_v107
			})()
			_ = _nwElement0_65
			var _nw294 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(7))
			_ = _nw294
			(_nw294).ArraySet1(_nwElement0_65, 0)
			(_nw294).ArraySet1(_1451_v107, 1)
			(_nw294).ArraySet1(_1451_v107, 2)
			(_nw294).ArraySet1(_1451_v107, 3)
			(_nw294).ArraySet1(_1451_v107, 4)
			(_nw294).ArraySet1(_1451_v107, 5)
			(_nw294).ArraySet1(_1451_v107, 6)
			_1454_v108 = _nw294
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1454_v108), 0))
			_ = _index267
			(_1454_v108).ArraySet1(_1451_v107, (_index267).Int())
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(275), _dafny.ArrayLen((_1454_v108), 0))
			_ = _index268
			var _nw295 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw295
			(_1454_v108).ArraySet1(_nw295, (_index268).Int())
			var _1455_v109 _dafny.Array
			_ = _1455_v109
			var _1456_v110 bool
			_ = _1456_v110
			var _1457_v111 _dafny.MultiSet
			_ = _1457_v111
			var _1458_v112 _dafny.Int
			_ = _1458_v112
			var _out62 _dafny.Array
			_ = _out62
			var _out63 bool
			_ = _out63
			var _out64 _dafny.MultiSet
			_ = _out64
			var _out65 _dafny.Int
			_ = _out65
			_out62, _out63, _out64, _out65 = Companion_Default___.M0(_1313_v13, globalState)
			_1455_v109 = _out62
			_1456_v110 = _out63
			_1457_v111 = _out64
			_1458_v112 = _out65
			var _1459_v113 _dafny.Array
			_ = _1459_v113
			var _1460_v114 bool
			_ = _1460_v114
			var _1461_v115 _dafny.MultiSet
			_ = _1461_v115
			var _1462_v116 _dafny.Int
			_ = _1462_v116
			var _out66 _dafny.Array
			_ = _out66
			var _out67 bool
			_ = _out67
			var _out68 _dafny.MultiSet
			_ = _out68
			var _out69 _dafny.Int
			_ = _out69
			_out66, _out67, _out68, _out69 = Companion_Default___.M0(_dafny.CodePoint('y'), globalState)
			_1459_v113 = _out66
			_1460_v114 = _out67
			_1461_v115 = _out68
			_1462_v116 = _out69
			(globalState).F2 = (_this).Fm2(_1458_v112, _1302_v4, globalState)
			var _1463_v117 *C0
			_ = _1463_v117
			var _nw296 *C0 = New_C0_()
			_ = _nw296
			_nw296.Ctor__((_this).F5())
			_1463_v117 = _nw296
			var _1464_v118 _dafny.MultiSet
			_ = _1464_v118
			_1464_v118 = _dafny.MultiSetOf(_1463_v117)
			var _1465_v119 _dafny.Map
			_ = _1465_v119
			_1465_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("vwwebb"), _1313_v13), _1464_v118)
			var _1466_v120 _dafny.MultiSet
			_ = _1466_v120
			_1466_v120 = _dafny.MultiSetOf(_dafny.IntOfInt64(326), (_this).F5())
			var _1467_v121 _dafny.Sequence
			_ = _1467_v121
			_1467_v121 = _dafny.SeqOf(_1466_v120)
			_1458_v112 = ((func() _dafny.MultiSet {
				if (_1465_v119).Contains((_1466_v120).IsSubsetOf((_1467_v121).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1467_v121).Cardinality()))).Uint32()).(_dafny.MultiSet))) {
					return (_1465_v119).Get((_1466_v120).IsSubsetOf((_1467_v121).Select((Companion_Default___.SafeIndex((_this).F5(), _dafny.IntOfUint32((_1467_v121).Cardinality()))).Uint32()).(_dafny.MultiSet))).(_dafny.MultiSet)
				}
				return _1464_v118
			})()).Cardinality()
		}
		var _nw297 *C2 = New_C2_()
		_ = _nw297
		_nw297.Ctor__((_this.F6()).Plus(_dafny.IntOfInt64(159)), (_dafny.Zero).Minus((_this).F5()))
		r0 = _nw297
		return r0
	}
}
func (_this *C3) F7() bool {
	{
		return _this._f7
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
