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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, globalState *GlobalState) bool {
	if false {
		return true
	} else {
		return !(_dafny.SetOf(!(false))).Contains(false)
	}
}
func (_static *CompanionStruct_Default___) Fm1(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.CodePoint, p1 _dafny.Sequence, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC5_(Companion_D3_.Create_DC5_(Companion_D3_.Create_DC5_(Companion_D3_.Create_DC4_(_dafny.IntOfInt64(-120), true, false))))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(311), _dafny.IntOfInt64(461), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(512))).Cardinality()))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_1).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(512))).Cardinality()))).Contains(_1_v0) {
					_coll1.Add(Companion_Default___.SafeDivisionInt(_1_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpviudyfg")).Cardinality())), _dafny.IntOfInt64(446))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality())).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-299), false)).Cardinality()))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v1 _dafny.Int
			_2_v1 = interface{}(_compr_0).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfInt64(311), _dafny.IntOfInt64(461), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Map {
				var _coll2 = _dafny.NewMapBuilder()
				_ = _coll2
				for _iter2 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(512))).Cardinality()))).Elements()); ; {
					_compr_2, _ok2 := _iter2()
					if !_ok2 {
						break
					}
					var _1_v0 _dafny.Int
					_1_v0 = interface{}(_compr_2).(_dafny.Int)
					if (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(512))).Cardinality()))).Contains(_1_v0) {
						_coll2.Add(Companion_Default___.SafeDivisionInt(_1_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cpviudyfg")).Cardinality())), _dafny.IntOfInt64(446))
					}
				}
				return _coll2.ToMap()
			}()).Cardinality())).Cardinality()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-299), false)).Cardinality()))).Contains(_2_v1) {
				_coll0.Add((_2_v1).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}()).Intersection((_dafny.SetOf(_dafny.IntOfInt64(633), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), true)).Cardinality())).Intersection(func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(334))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _3_v2 _dafny.Int
			_3_v2 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(381)).Cmp(_3_v2) <= 0) && ((_3_v2).Cmp(_dafny.IntOfInt64(334)) < 0) {
				_coll3.Add((_3_v2).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(398), false), false)).Cardinality()))
			}
		}
		return _coll3.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Int {
	return (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true, !(false)))).Cardinality()))).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(978))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm6(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(15))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality())).Cmp(_dafny.IntOfInt64(865)) < 0 {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(233))
	} else {
		return _dafny.MultiSetOf(_dafny.IntOfInt64(289))
	}
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("hkhqfcu")
}
func (_static *CompanionStruct_Default___) Fm8(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('x')
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(821))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	})))).Union(_dafny.MultiSetOf(_dafny.CodePoint('e'), _dafny.CodePoint('m')))
}
func (_static *CompanionStruct_Default___) Fm10(globalState *GlobalState) _dafny.Map {
	var _source0 D7 = (func() D7 {
		if !(false) {
			return Companion_D7_.Create_DC15_(_dafny.IntOfInt64(407), false, _dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()))
		}
		return Companion_D7_.Create_DC15_((_dafny.Zero).Minus(_dafny.IntOfInt64(-917)), false, _dafny.IntOfInt64(757))
	})()
	_ = _source0
	if _source0.Is_DC14() {
		var _6___mcc_h0 _dafny.Int = _source0.Get_().(D7_DC14).Cf21
		_ = _6___mcc_h0
		var _7___mcc_h1 bool = _source0.Get_().(D7_DC14).Cf22
		_ = _7___mcc_h1
		var _8___mcc_h2 bool = _source0.Get_().(D7_DC14).Cf23
		_ = _8___mcc_h2
		var _9___mcc_h3 bool = _source0.Get_().(D7_DC14).Cf24
		_ = _9___mcc_h3
		var _10___mcc_h4 _dafny.Int = _source0.Get_().(D7_DC14).Cf25
		_ = _10___mcc_h4
		var _11_cf25 _dafny.Int = _10___mcc_h4
		_ = _11_cf25
		var _12_cf24 bool = _9___mcc_h3
		_ = _12_cf24
		var _13_cf23 bool = _8___mcc_h2
		_ = _13_cf23
		var _14_cf22 bool = _7___mcc_h1
		_ = _14_cf22
		var _15_cf21 _dafny.Int = _6___mcc_h0
		_ = _15_cf21
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ksuka"), true)
	} else if _source0.Is_DC15() {
		var _16___mcc_h5 _dafny.Int = _source0.Get_().(D7_DC15).Cf26
		_ = _16___mcc_h5
		var _17___mcc_h6 bool = _source0.Get_().(D7_DC15).Cf27
		_ = _17___mcc_h6
		var _18___mcc_h7 _dafny.Int = _source0.Get_().(D7_DC15).Cf28
		_ = _18___mcc_h7
		var _19_cf28 _dafny.Int = _18___mcc_h7
		_ = _19_cf28
		var _20_cf27 bool = _17___mcc_h6
		_ = _20_cf27
		var _21_cf26 _dafny.Int = _16___mcc_h5
		_ = _21_cf26
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})), _20_cf27)).Merge(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("enjjy"), _dafny.CodePoint('e'))).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _23_v0 _dafny.Sequence
				_23_v0 = interface{}(_compr_4).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("enjjy"), _dafny.CodePoint('e'))).Contains(_23_v0) {
					_coll4.Add(_23_v0, _20_cf27)
				}
			}
			return _coll4.ToMap()
		}())
	} else {
		var _24___mcc_h8 _dafny.Map = _source0.Get_().(D7_DC13).Cf20
		_ = _24___mcc_h8
		var _25_cf20 _dafny.Map = _24___mcc_h8
		_ = _25_cf20
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("bxdetlja"), false)
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	return (true) && (false)
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(!(!(true)) || (false), true)
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Sequence {
	var _source1 D12 = (func() D12 {
		if false {
			return Companion_D12_.Create_DC27_(false, _dafny.IntOfInt64(154))
		}
		return Companion_D12_.Create_DC27_(false, _dafny.IntOfInt64(85))
	})()
	_ = _source1
	if _source1.Is_DC27() {
		var _26___mcc_h0 bool = _source1.Get_().(D12_DC27).Cf45
		_ = _26___mcc_h0
		var _27___mcc_h1 _dafny.Int = _source1.Get_().(D12_DC27).Cf46
		_ = _27___mcc_h1
		var _28_cf46 _dafny.Int = _27___mcc_h1
		_ = _28_cf46
		var _29_cf45 bool = _26___mcc_h0
		_ = _29_cf45
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_28_cf46), _dafny.SeqOf(_28_cf46))
	} else {
		var _30___mcc_h2 _dafny.Map = _source1.Get_().(D12_DC26).Cf44
		_ = _30___mcc_h2
		var _31_cf44 _dafny.Map = _30___mcc_h2
		_ = _31_cf44
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf(!(true))).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-351))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_32_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-346)
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(91)).Times(_dafny.IntOfInt64(417)), (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('p'))).Cardinality()))).Cardinality())).Cmp((func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(341), _dafny.IntOfInt64(-240))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fwkthi")).Cardinality()))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _33_v0 _dafny.Int
			_33_v0 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(341), _dafny.IntOfInt64(-240))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fwkthi")).Cardinality())), _33_v0) {
				_coll5.Add((_33_v0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-881))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_34_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(87)
				}))).Cardinality())), _dafny.IntOfInt64(520))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality()) >= 0)
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), (_dafny.MultiSetOf(_dafny.IntOfInt64(-59))).Cardinality())).Merge((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _dafny.IntOfInt64(568))).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _35_v0 _dafny.CodePoint
			_35_v0 = interface{}(_compr_6).(_dafny.CodePoint)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), _dafny.IntOfInt64(568))).Contains(_35_v0) {
				_coll6.Add(_35_v0, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))
			}
		}
		return _coll6.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), (Companion_D8_.Create_DC17_(_dafny.CodePoint('h'), _dafny.CodePoint('a'), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Dtor_cf32())))
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("hfqueonrf"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(258))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_36_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yayviyc"))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(616))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_37_i1 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("ocnyb")
	})))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.Zero).Minus((func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(594), _dafny.IntOfInt64(946))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _38_v0 _dafny.Int
			_38_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(594)).Cmp(_38_v0) <= 0) && ((_38_v0).Cmp(_dafny.IntOfInt64(946)) < 0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_38_v0, _dafny.IntOfInt64(-345)), _dafny.IntOfInt64(331))
			}
		}
		return _coll7.ToMap()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.MultiSet, p1 D8, p2 bool, globalState *GlobalState) _dafny.Int {
	return (Companion_Default___.SafeDivisionInt((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(93))).Cardinality(), _dafny.IntOfInt64(393))).Times(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D11_.Create_DC25_(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), !(!(!(false))))).Dtor_cf43(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true)))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.Sequence {
	var _source2 D8 = Companion_D8_.Create_DC18_(Companion_D8_.Create_DC18_(Companion_D8_.Create_DC17_(_dafny.CodePoint('o'), _dafny.CodePoint('s'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ibkkhb")).Cardinality()))))
	_ = _source2
	if _source2.Is_DC17() {
		var _39___mcc_h0 _dafny.CodePoint = _source2.Get_().(D8_DC17).Cf30
		_ = _39___mcc_h0
		var _40___mcc_h1 _dafny.CodePoint = _source2.Get_().(D8_DC17).Cf31
		_ = _40___mcc_h1
		var _41___mcc_h2 _dafny.Int = _source2.Get_().(D8_DC17).Cf32
		_ = _41___mcc_h2
		var _42_cf32 _dafny.Int = _41___mcc_h2
		_ = _42_cf32
		var _43_cf31 _dafny.CodePoint = _40___mcc_h1
		_ = _43_cf31
		var _44_cf30 _dafny.CodePoint = _39___mcc_h0
		_ = _44_cf30
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), (Companion_D16_.Create_DC34_(_dafny.SeqOf(false))).Dtor_cf57())
	} else if _source2.Is_DC16() {
		var _45___mcc_h3 _dafny.Sequence = _source2.Get_().(D8_DC16).Cf29
		_ = _45___mcc_h3
		var _46_cf29 _dafny.Sequence = _45___mcc_h3
		_ = _46_cf29
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true), _dafny.SeqOf((Companion_D4_.Create_DC8_(_dafny.CodePoint('d'), true, false)).Dtor_cf14(), false, !(true)))
	} else {
		var _47___mcc_h4 D8 = _source2.Get_().(D8_DC18).Cf33
		_ = _47___mcc_h4
		var _48_cf33 D8 = _47___mcc_h4
		_ = _48_cf33
		return _dafny.SeqOf((Companion_D7_.Create_DC15_(_dafny.IntOfInt64(763), false, _dafny.IntOfInt64(866))).Dtor_cf27(), false, !((Companion_D16_.Create_DC35_(true, _dafny.IntOfInt64(348), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vbrr")).Cardinality()))).Dtor_cf58()), false)
	}
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	var _source3 D11 = Companion_D11_.Create_DC25_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(720), _dafny.IntOfInt64(-331), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(592))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_49_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-308)))).Cardinality())), true)
	_ = _source3
	if _source3.Is_DC25() {
		var _50___mcc_h0 _dafny.Int = _source3.Get_().(D11_DC25).Cf42
		_ = _50___mcc_h0
		var _51___mcc_h1 bool = _source3.Get_().(D11_DC25).Cf43
		_ = _51___mcc_h1
		var _52_cf43 bool = _51___mcc_h1
		_ = _52_cf43
		var _53_cf42 _dafny.Int = _50___mcc_h0
		_ = _53_cf42
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('n'), (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-190))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_54_cf42 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_55_i1 _dafny.Int) _dafny.Int {
				return _54_cf42
			}
		})(_53_cf42))), _dafny.SeqOf(_53_cf42, _53_cf42, _53_cf42), _dafny.SeqOf(_53_cf42, _53_cf42, (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_53_cf42)).Cardinality()))), _53_cf42, _53_cf42))).Cardinality())
	} else {
		var _56___mcc_h2 _dafny.MultiSet = _source3.Get_().(D11_DC24).Cf41
		_ = _56___mcc_h2
		var _57_cf41 _dafny.MultiSet = _56___mcc_h2
		_ = _57_cf41
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfInt64(628))
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.CodePoint, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true, false, true)).Intersection(_dafny.SetOf(false, false))).Intersection((_dafny.SetOf(false, true)).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('r')
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ybjet")).Cardinality())), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("odkdqes")).Cardinality()), _dafny.IntOfInt64(-633), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ymjr")).Cardinality())))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, globalState *GlobalState) _dafny.Sequence {
	if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(false))), _dafny.IntOfInt64(420))).Contains(false) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ggcg")).Cardinality()))), _dafny.SetOf(_dafny.IntOfInt64(836), _dafny.IntOfInt64(188))), _dafny.SeqOf(func() _dafny.Set {
			var _coll8 = _dafny.NewBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(393), _dafny.IntOfInt64(306))); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _58_v0 _dafny.Int
				_58_v0 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(393)).Cmp(_58_v0) <= 0) && ((_58_v0).Cmp(_dafny.IntOfInt64(306)) < 0) {
					_coll8.Add((_58_v0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(716))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg10 _dafny.Int) interface{} {
							return coer10(arg10)
						}
					}(func(_59_i0 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_dafny.IntOfInt64(-150))
					}))).Cardinality())))
				}
			}
			return _coll8.ToSet()
		}()))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-234))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_60_i1 _dafny.Int) _dafny.Set {
			return _dafny.SetOf(_dafny.IntOfInt64(512))
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC3_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("koggtmo")))
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) _dafny.Map {
	if !(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Equals(_dafny.MultiSetOf(!(true))) {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
				var _coll9 = _dafny.NewBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((func() _dafny.Map {
					var _coll10 = _dafny.NewMapBuilder()
					_ = _coll10
					for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(571), _dafny.IntOfInt64(-350))); ; {
						_compr_10, _ok10 := _iter10()
						if !_ok10 {
							break
						}
						var _61_v0 _dafny.Int
						_61_v0 = interface{}(_compr_10).(_dafny.Int)
						if ((_dafny.IntOfInt64(571)).Cmp(_61_v0) <= 0) && ((_61_v0).Cmp(_dafny.IntOfInt64(-350)) < 0) {
							_coll10.Add((_61_v0).Times(_dafny.IntOfInt64(278)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality()))
						}
					}
					return _coll10.ToMap()
				}()).Keys().Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _62_v1 _dafny.Int
					_62_v1 = interface{}(_compr_9).(_dafny.Int)
					if (func() _dafny.Map {
						var _coll11 = _dafny.NewMapBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(571), _dafny.IntOfInt64(-350))); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _61_v0 _dafny.Int
							_61_v0 = interface{}(_compr_11).(_dafny.Int)
							if ((_dafny.IntOfInt64(571)).Cmp(_61_v0) <= 0) && ((_61_v0).Cmp(_dafny.IntOfInt64(-350)) < 0) {
								_coll11.Add((_61_v0).Times(_dafny.IntOfInt64(278)), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality()))
							}
						}
						return _coll11.ToMap()
					}()).Contains(_62_v1) {
						_coll9.Add(Companion_Default___.SafeDivisionInt(_62_v1, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(335))).Cardinality())))
					}
				}
				return _coll9.ToSet()
			}(), _dafny.UnicodeSeqOfUtf8Bytes("esyx"))
		} else {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_63_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})))
		}
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(142)), _dafny.UnicodeSeqOfUtf8Bytes("kraaoyc"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(957)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_64_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		}))))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC17_(_dafny.CodePoint('c'), _dafny.CodePoint('c'), (_dafny.Zero).Minus(_dafny.IntOfInt64(-416)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(true) || (true))
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rknh"), _dafny.IntOfInt64(259))
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) (_dafny.Int, bool, bool, _dafny.Int) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 bool = false
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	var _65_v0 _dafny.Int
	_ = _65_v0
	_65_v0 = _dafny.IntOfInt64(772)
	var _hi0 _dafny.Int = _dafny.IntOfInt64(310)
	_ = _hi0
	for _66_i0 := _65_v0; _66_i0.Cmp(_hi0) < 0; _66_i0 = _66_i0.Plus(_dafny.One) {
		var _67_v1 _dafny.Sequence
		_ = _67_v1
		_67_v1 = _dafny.UnicodeSeqOfUtf8Bytes("cviinusi")
		var _68_v2 _dafny.MultiSet
		_ = _68_v2
		_68_v2 = _dafny.MultiSetOf(_65_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_65_v0, _dafny.IntOfInt64(374))).Cardinality(), _dafny.IntOfUint32((_67_v1).Cardinality()), _66_i0, (_dafny.Zero).Minus(_66_i0))
		_68_v2 = _68_v2
		var _69_v3 bool
		_ = _69_v3
		_69_v3 = false
		var _70_v4 bool
		_ = _70_v4
		_70_v4 = _69_v3
		var _71_v5 T1
		_ = _71_v5
		var _nw0 *C0 = New_C0_()
		_ = _nw0
		_nw0.Ctor__(_69_v3)
		_71_v5 = _nw0
		var _72_v6 _dafny.Map
		_ = _72_v6
		_72_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_71_v5, _69_v3)
		var _73_v7 _dafny.MultiSet
		_ = _73_v7
		_73_v7 = _dafny.MultiSetOf(_71_v5)
		var _rhs0 bool = _70_v4
		_ = _rhs0
		var _rhs1 _dafny.Int = ((_dafny.Zero).Minus((_65_v0).Times((_72_v6).Cardinality()))).Minus(_66_i0)
		_ = _rhs1
		var _rhs2 bool = _69_v3
		_ = _rhs2
		var _rhs3 bool = _69_v3
		_ = _rhs3
		var _rhs4 _dafny.Int = (func() _dafny.Int {
			if (_73_v7).Contains(_71_v5) {
				return (_73_v7).Multiplicity(_71_v5)
			}
			return Companion_Default___.SafeDivisionInt(_66_i0, _66_i0)
		})()
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		_70_v4 = _rhs0
		r0 = _rhs1
		_lhs0.F5 = _rhs2
		_lhs1.F5 = _rhs3
		_lhs2.F8 = _rhs4
		r1 = _69_v3
		(globalState).F8 = _65_v0
	}
	var _74_v8 bool
	_ = _74_v8
	_74_v8 = true
	(globalState).F17 = _74_v8
	var _75_v9 _dafny.Array
	_ = _75_v9
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
	_ = _nw1
	_75_v9 = _nw1
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_75_v9), 0))
	_ = _index0
	(_75_v9).ArraySet1(_65_v0, (_index0).Int())
	var _76_v10 _dafny.Sequence
	_ = _76_v10
	_76_v10 = _dafny.SeqOf(_74_v8, false)
	var _77_v11 _dafny.MultiSet
	_ = _77_v11
	_77_v11 = _dafny.MultiSetOf(_74_v8, _74_v8, Companion_Default___.Fm0(_76_v10, globalState), !(_74_v8), _74_v8)
	var _78_v12 *C0
	_ = _78_v12
	var _nw2 *C0 = New_C0_()
	_ = _nw2
	_nw2.Ctor__(_74_v8)
	_78_v12 = _nw2
	var _79_v13 _dafny.CodePoint
	_ = _79_v13
	_79_v13 = _dafny.CodePoint('c')
	var _80_v14 _dafny.Map
	_ = _80_v14
	_80_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_78_v12, _79_v13)
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_75_v9), 0))
	_ = _index1
	(_75_v9).ArraySet1(((func() _dafny.Int {
		if (_77_v11).Contains(_74_v8) {
			return (_77_v11).Multiplicity(_74_v8)
		}
		return _65_v0
	})()).Plus((_80_v14).Cardinality()), (_index1).Int())
	var _81_v15 _dafny.Array
	_ = _81_v15
	var _nw3 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
	_ = _nw3
	_81_v15 = _nw3
	var _82_v16 _dafny.Sequence
	_ = _82_v16
	_82_v16 = _dafny.UnicodeSeqOfUtf8Bytes("xnm")
	var _83_v17 _dafny.Map
	_ = _83_v17
	_83_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_65_v0), _82_v16)
	var _84_v18 T1
	_ = _84_v18
	var _nw4 *C4 = New_C4_()
	_ = _nw4
	_nw4.Ctor__(_65_v0)
	_84_v18 = _nw4
	var _85_v19 *C1
	_ = _85_v19
	var _nw5 *C1 = New_C1_()
	_ = _nw5
	_nw5.Ctor__(_83_v17, _84_v18)
	_85_v19 = _nw5
	var _86_v20 _dafny.Set
	_ = _86_v20
	_86_v20 = _dafny.SetOf(_74_v8)
	var _rhs5 _dafny.Array = _81_v15
	_ = _rhs5
	var _rhs6 *C1 = _85_v19
	_ = _rhs6
	var _rhs7 _dafny.Int = Companion_Default___.SafeModuloInt(((_86_v20).Cardinality()).Minus(_dafny.IntOfUint32((Companion_Default___.Fm13(globalState)).Cardinality())), Companion_Default___.SafeDivisionInt(_65_v0, (_75_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_75_v9), 0))).Int()).(_dafny.Int)))
	_ = _rhs7
	var _rhs8 _dafny.Int = ((_dafny.IntOfInt64(818)).Times(_65_v0)).Plus(_65_v0)
	_ = _rhs8
	var _lhs3 *GlobalState = globalState
	_ = _lhs3
	_81_v15 = _rhs5
	_85_v19 = _rhs6
	_lhs3.F8 = _rhs7
	_65_v0 = _rhs8
	var _87_v21 _dafny.Array
	_ = _87_v21
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
	_ = _nw6
	_87_v21 = _nw6
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_87_v21), 0))
	_ = _index2
	(_87_v21).ArraySet1(Companion_Default___.Fm1(globalState), (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_87_v21), 0))
	_ = _index3
	(_87_v21).ArraySet1(_82_v16, (_index3).Int())
	var _88_v22 *C3
	_ = _88_v22
	var _nw7 *C3 = New_C3_()
	_ = _nw7
	_nw7.Ctor__(_75_v9)
	_88_v22 = _nw7
	r0 = (_65_v0).Times(_dafny.IntOfInt64(-651))
	var _89_v23 D8
	_ = _89_v23
	_89_v23 = Companion_D8_.Create_DC16_((_87_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_87_v21), 0))).Int()).(_dafny.Sequence))
	r1 = _dafny.Companion_Sequence_.IsPrefixOf((_89_v23).Dtor_cf29(), _82_v16)
	r2 = true
	r3 = (_65_v0).Minus((_75_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_75_v9), 0))).Int()).(_dafny.Int))
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _90_v0 _dafny.Sequence
	_ = _90_v0
	_90_v0 = _dafny.UnicodeSeqOfUtf8Bytes("wifcawqh")
	var _91_v1 _dafny.Array
	_ = _91_v1
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(16)
	_ = _len0_0
	var _nw8 _dafny.Array
	_ = _nw8
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw8 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Int = (func(_92_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_93_i0 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_93_i0, _dafny.IntOfUint32((_92_v0).Cardinality()))
			}
		})(_90_v0)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw8 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw8).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw8).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_91_v1 = _nw8
	var _94_v2 _dafny.Array
	_ = _94_v2
	var _nwElement0_0 _dafny.Array = _91_v1
	_ = _nwElement0_0
	var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(3))
	_ = _nw9
	(_nw9).ArraySet1(_nwElement0_0, 0)
	(_nw9).ArraySet1(_91_v1, 1)
	(_nw9).ArraySet1(_91_v1, 2)
	_94_v2 = _nw9
	var _95_v3 bool
	_ = _95_v3
	_95_v3 = false
	var _96_v4 _dafny.Map
	_ = _96_v4
	_96_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v3, false)
	var _97_v5 _dafny.Set
	_ = _97_v5
	_97_v5 = _dafny.SetOf(false)
	var _98_v6 _dafny.Array
	_ = _98_v6
	var _nwElement0_1 _dafny.Set = _97_v5
	_ = _nwElement0_1
	var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.One)
	_ = _nw10
	(_nw10).ArraySet1(_nwElement0_1, 0)
	_98_v6 = _nw10
	var _99_v7 _dafny.Int
	_ = _99_v7
	_99_v7 = _dafny.IntOfInt64(637)
	var _100_v8 bool
	_ = _100_v8
	_100_v8 = _95_v3
	var _101_v9 _dafny.Map
	_ = _101_v9
	_101_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v7, (_100_v8))
	var _102_v10 _dafny.Array
	_ = _102_v10
	var _nwElement0_2 _dafny.Sequence = _90_v0
	_ = _nwElement0_2
	var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(4))
	_ = _nw11
	(_nw11).ArraySet1(_nwElement0_2, 0)
	(_nw11).ArraySet1(_90_v0, 1)
	(_nw11).ArraySet1(_90_v0, 2)
	(_nw11).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(499))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_103_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	})), 3)
	_102_v10 = _nw11
	var _104_v11 _dafny.Map
	_ = _104_v11
	_104_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_100_v8), _dafny.UnicodeSeqOfUtf8Bytes("dwmjjp"))
	var _105_globalState *GlobalState
	_ = _105_globalState
	var _nw12 *GlobalState = New_GlobalState_()
	_ = _nw12
	_nw12.Ctor__(_90_v0, _94_v2, (_96_v4).Merge(_96_v4), _dafny.CodePoint('i'), _dafny.IntOfInt64(992), false, _dafny.IntOfInt64(-915), _98_v6, _dafny.IntOfInt64(504), _101_v9, _90_v0, true, _102_v10, _94_v2, _dafny.IntOfInt64(34), _dafny.IntOfInt64(176), _104_v11, true, _dafny.IntOfInt64(649), true, false, false, true, _91_v1, _dafny.UnicodeSeqOfUtf8Bytes("l"))
	_105_globalState = _nw12
	var _106_v12 _dafny.Sequence
	_ = _106_v12
	_106_v12 = _dafny.SeqOf(_99_v7)
	var _107_v13 _dafny.Sequence
	_ = _107_v13
	_107_v13 = _dafny.SeqOf(_95_v3)
	var _108_v14 _dafny.Map
	_ = _108_v14
	_108_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_106_v12, Companion_Default___.Fm0(_107_v13, _105_globalState))
	_108_v14 = (_108_v14).Merge((_108_v14).Merge(_108_v14))
	var _109_v15 _dafny.Map
	_ = _109_v15
	_109_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_110_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	})), _95_v3)
	var _111_i2 _dafny.Int
	_ = _111_i2
	_111_i2 = _dafny.Zero
	{
		for (_109_v15).Contains(_dafny.Companion_Sequence_.Concatenate(_90_v0, _90_v0)) {
			{
				if (_111_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_111_i2 = (_111_i2).Plus(_dafny.One)
				(_105_globalState).F22 = !(_95_v3)
				_99_v7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("puigmlj"), _dafny.UnicodeSeqOfUtf8Bytes("e"))).Cardinality())
				_109_v15 = (_109_v15).Update(_90_v0, _95_v3)
				var _112_v16 _dafny.Map
				_ = _112_v16
				_112_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v8, _95_v3)
				var _113_v17 _dafny.Map
				_ = _113_v17
				_113_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v1, _112_v16)
				_112_v16 = (func() _dafny.Map {
					if (_113_v17).Contains(_91_v1) {
						return (_113_v17).Get(_91_v1).(_dafny.Map)
					}
					return _112_v16
				})()
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _114_v18 _dafny.Int
	_ = _114_v18
	var _115_v19 bool
	_ = _115_v19
	var _116_v20 bool
	_ = _116_v20
	var _117_v21 _dafny.Int
	_ = _117_v21
	var _out0 _dafny.Int
	_ = _out0
	var _out1 bool
	_ = _out1
	var _out2 bool
	_ = _out2
	var _out3 _dafny.Int
	_ = _out3
	_out0, _out1, _out2, _out3 = Companion_Default___.M0(_105_globalState)
	_114_v18 = _out0
	_115_v19 = _out1
	_116_v20 = _out2
	_117_v21 = _out3
	var _118_v22 *C0
	_ = _118_v22
	var _nw13 *C0 = New_C0_()
	_ = _nw13
	_nw13.Ctor__(_95_v3)
	_118_v22 = _nw13
	_95_v3 = _115_v19
	var _119_v23 *C3
	_ = _119_v23
	var _nw14 *C3 = New_C3_()
	_ = _nw14
	_nw14.Ctor__(_91_v1)
	_119_v23 = _nw14
	var _120_v24 D13
	_ = _120_v24
	_120_v24 = Companion_D13_.Create_DC29_(_119_v23)
	var _121_v25 D13
	_ = _121_v25
	_121_v25 = Companion_D13_.Create_DC30_(_120_v24)
	var _source4 D13 = _121_v25
	_ = _source4
	if _source4.Is_DC29() {
		var _122___mcc_h0 *C3 = _source4.Get_().(D13_DC29).Cf48
		_ = _122___mcc_h0
		var _123_cf48 *C3 = _122___mcc_h0
		_ = _123_cf48
		var _124_v26 _dafny.Int
		_ = _124_v26
		var _125_v27 bool
		_ = _125_v27
		var _126_v28 bool
		_ = _126_v28
		var _127_v29 _dafny.Int
		_ = _127_v29
		var _out4 _dafny.Int
		_ = _out4
		var _out5 bool
		_ = _out5
		var _out6 bool
		_ = _out6
		var _out7 _dafny.Int
		_ = _out7
		_out4, _out5, _out6, _out7 = Companion_Default___.M0(_105_globalState)
		_124_v26 = _out4
		_125_v27 = _out5
		_126_v28 = _out6
		_127_v29 = _out7
		var _128_v30 _dafny.Set
		_ = _128_v30
		_128_v30 = _dafny.SetOf(_114_v18)
		_117_v21 = (_128_v30).Cardinality()
		var _129_v31 T1
		_ = _129_v31
		var _nw15 *C0 = New_C0_()
		_ = _nw15
		_nw15.Ctor__(!(_116_v20))
		_129_v31 = _nw15
		var _130_v32 *C1
		_ = _130_v32
		var _nw16 *C1 = New_C1_()
		_ = _nw16
		_nw16.Ctor__(Companion_Default___.Fm26(_105_globalState), _129_v31)
		_130_v32 = _nw16
		var _nw17 *C1 = New_C1_()
		_ = _nw17
		_nw17.Ctor__((_130_v32).F28(), (_130_v32).F29())
		_130_v32 = _nw17
		var _131_v33 D6
		_ = _131_v33
		_131_v33 = Companion_D6_.Create_DC10_(_91_v1)
		var _132_v34 *C3
		_ = _132_v34
		var _nw18 *C3 = New_C3_()
		_ = _nw18
		_nw18.Ctor__((_131_v33).Dtor_cf17())
		_132_v34 = _nw18
	} else if _source4.Is_DC28() {
		var _133___mcc_h1 _dafny.Sequence = _source4.Get_().(D13_DC28).Cf47
		_ = _133___mcc_h1
		var _134_cf47 _dafny.Sequence = _133___mcc_h1
		_ = _134_cf47
		(_105_globalState).F17 = _116_v20
		var _135_v35 _dafny.MultiSet
		_ = _135_v35
		_135_v35 = _dafny.MultiSetOf(_116_v20)
		var _136_v37 _dafny.Map
		_ = _136_v37
		_136_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_100_v8, _116_v20)
		if (func() _dafny.Set {
			var _coll12 = _dafny.NewBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_135_v35).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _137_v36 bool
				_137_v36 = interface{}(_compr_12).(bool)
				if (_135_v35).Contains(_137_v36) {
					_coll12.Add(_137_v36)
				}
			}
			return _coll12.ToSet()
		}()).IsDisjointFrom(func() _dafny.Set {
			var _coll13 = _dafny.NewBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_136_v37).Keys().Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _138_v38 bool
				_138_v38 = interface{}(_compr_13).(bool)
				if (_136_v37).Contains(_138_v38) {
					_coll13.Add(_138_v38)
				}
			}
			return _coll13.ToSet()
		}()) {
			var _139_v39 _dafny.MultiSet
			_ = _139_v39
			_139_v39 = _dafny.MultiSetOf(false, (_118_v22).F27())
			var _140_v40 _dafny.Map
			_ = _140_v40
			_140_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_91_v1, Companion_Default___.Fm18(_139_v39, Companion_Default___.Fm27(_114_v18, _105_globalState), (_118_v22).F27(), _105_globalState))
			var _141_v41 _dafny.Sequence
			_ = _141_v41
			_141_v41 = _dafny.SeqOf(_119_v23, _119_v23)
			var _rhs9 _dafny.Map = _140_v40
			_ = _rhs9
			var _rhs10 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_141_v41, _141_v41)).Cardinality())
			_ = _rhs10
			var _rhs11 _dafny.Int = _117_v21
			_ = _rhs11
			var _lhs4 *GlobalState = _105_globalState
			_ = _lhs4
			var _lhs5 *GlobalState = _105_globalState
			_ = _lhs5
			_140_v40 = _rhs9
			_lhs4.F8 = _rhs10
			_lhs5.F8 = _rhs11
			var _rhs12 bool = _115_v19
			_ = _rhs12
			var _rhs13 bool = _116_v20
			_ = _rhs13
			var _lhs6 *GlobalState = _105_globalState
			_ = _lhs6
			var _lhs7 *GlobalState = _105_globalState
			_ = _lhs7
			_lhs6.F19 = _rhs12
			_lhs7.F22 = _rhs13
			(_119_v23).F26 = _91_v1
			_117_v21 = (_99_v7).Times(_99_v7)
			var _arr0 _dafny.Array = _119_v23.F26
			_ = _arr0
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_119_v23.F26), 0))
			_ = _index4
			(_arr0).ArraySet1(Companion_Default___.SafeDivisionInt(_117_v21, _117_v21), (_index4).Int())
			var _arr1 _dafny.Array = _119_v23.F26
			_ = _arr1
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_119_v23.F26), 0))
			_ = _index5
			var _rhs14 _dafny.Int = _114_v18
			_ = _rhs14
			var _rhs15 _dafny.Int = (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(296))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}(func(_142_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality())).Plus(_117_v21))
			_ = _rhs15
			var _lhs8 _dafny.Array = _119_v23.F26
			_ = _lhs8
			var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_119_v23.F26), 0))
			_ = _lhs9
			_114_v18 = _rhs14
			(_lhs8).ArraySet1(_rhs15, (_lhs9).Int())
		} else {
			_99_v7 = _114_v18
			(_105_globalState).F22 = true
			var _143_v42 _dafny.MultiSet
			_ = _143_v42
			_143_v42 = _dafny.MultiSetOf((_118_v22).F27())
			_99_v7 = Companion_Default___.SafeDivisionInt((_143_v42).Cardinality(), _117_v21)
			var _arr2 _dafny.Array = _119_v23.F26
			_ = _arr2
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_119_v23.F26), 0))
			_ = _index6
			(_arr2).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_99_v7), _114_v18), (_index6).Int())
			var _144_v43 D11
			_ = _144_v43
			_144_v43 = Companion_D11_.Create_DC24_(_dafny.MultiSetFromSeq(_107_v13))
			var _145_v44 _dafny.Map
			_ = _145_v44
			_145_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v19, _114_v18)
			var _146_v45 _dafny.Map
			_ = _146_v45
			_146_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v43, _145_v44)
			var _arr3 _dafny.Array = _119_v23.F26
			_ = _arr3
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_119_v23.F26), 0))
			_ = _index7
			(_arr3).ArraySet1(((func() _dafny.Map {
				if (_146_v45).Contains(_144_v43) {
					return (_146_v45).Get(_144_v43).(_dafny.Map)
				}
				return (_145_v44).Merge(_145_v44)
			})()).Cardinality(), (_index7).Int())
			var _147_v46 _dafny.Map
			_ = _147_v46
			_147_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_90_v0).Cardinality()), _91_v1)
			var _148_v47 _dafny.Sequence
			_ = _148_v47
			_148_v47 = _dafny.SeqOf(_91_v1)
			(_105_globalState).F23 = (func() _dafny.Array {
				if (_147_v46).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v22).F27(), (_118_v22).F27())).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_107_v13).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v22).F27(), (_118_v22).F27())).Cardinality())).Cardinality()))).Uint32(), _99_v7)).Cardinality())) {
					return (_147_v46).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v22).F27(), (_118_v22).F27())).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_107_v13).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_118_v22).F27(), (_118_v22).F27())).Cardinality())).Cardinality()))).Uint32(), _99_v7)).Cardinality())).(_dafny.Array)
				}
				return (_148_v47).Select((Companion_Default___.SafeIndex(_117_v21, _dafny.IntOfUint32((_148_v47).Cardinality()))).Uint32()).(_dafny.Array)
			})()
		}
		(_105_globalState).F17 = !(_116_v20)
		(_119_v23).M2(_105_globalState)
	} else {
		var _149___mcc_h2 D13 = _source4.Get_().(D13_DC30).Cf49
		_ = _149___mcc_h2
		var _150_cf49 D13 = _149___mcc_h2
		_ = _150_cf49
		var _151_v48 T1
		_ = _151_v48
		var _nw19 *C0 = New_C0_()
		_ = _nw19
		_nw19.Ctor__((_99_v7).Cmp(_117_v21) <= 0)
		_151_v48 = _nw19
		_151_v48 = (func() T1 {
			if (_118_v22).F27() {
				return _151_v48
			}
			return _151_v48
		})()
		_90_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(580))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_152_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})), _90_v0)
		(_119_v23).M1(_105_globalState)
		var _153_v49 _dafny.Array
		_ = _153_v49
		var _nw20 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw20
		_153_v49 = _nw20
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_153_v49), 0))
		_ = _index8
		(_153_v49).ArraySet1(_94_v2, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(60), _dafny.ArrayLen((_153_v49), 0))
		_ = _index9
		(_153_v49).ArraySet1(_94_v2, (_index9).Int())
	}
	var _154_v50 *C0
	_ = _154_v50
	var _nw21 *C0 = New_C0_()
	_ = _nw21
	_nw21.Ctor__((_107_v13).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_99_v7), _dafny.IntOfUint32((_107_v13).Cardinality()))).Uint32()).(bool))
	_154_v50 = _nw21
	var _155_v51 D11
	_ = _155_v51
	_155_v51 = Companion_D11_.Create_DC24_(_dafny.MultiSetOf(_115_v19))
	var _156_v52 _dafny.Map
	_ = _156_v52
	_156_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v7, _155_v51)
	var _157_v53 _dafny.Map
	_ = _157_v53
	_157_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_156_v52).Cardinality(), _dafny.IntOfInt64(375))
	_117_v21 = (_117_v21).Plus(((func() _dafny.Int {
		if (_157_v53).Contains(_117_v21) {
			return (_157_v53).Get(_117_v21).(_dafny.Int)
		}
		return _99_v7
	})()).Plus(_114_v18))
	var _158_v54 *C4
	_ = _158_v54
	var _nw22 *C4 = New_C4_()
	_ = _nw22
	_nw22.Ctor__(_dafny.IntOfInt64(849))
	_158_v54 = _nw22
	_158_v54 = _158_v54
	var _159_v55 *C0
	_ = _159_v55
	var _nw23 *C0 = New_C0_()
	_ = _nw23
	_nw23.Ctor__(((_154_v50).F27()) == (_116_v20))
	_159_v55 = _nw23
	var _160_v56 _dafny.Map
	_ = _160_v56
	_160_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_107_v13, false)
	var _161_v58 _dafny.Sequence
	_ = _161_v58
	_161_v58 = _dafny.SeqOf(_107_v13, _107_v13, _107_v13)
	_160_v56 = (func() _dafny.Map {
		var _coll14 = _dafny.NewMapBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate((_161_v58).Elements()); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _162_v57 _dafny.Sequence
			_162_v57 = interface{}(_compr_14).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_161_v58, _162_v57) {
				_coll14.Add(_162_v57, !((_154_v50).F27()))
			}
		}
		return _coll14.ToMap()
	}()).Update(_dafny.Companion_Sequence_.Concatenate(_107_v13, _107_v13), (_118_v22).F27())
	(_158_v54).M2(_105_globalState)
	var _163_v59 _dafny.Array
	_ = _163_v59
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_1
	var _nw24 _dafny.Array
	_ = _nw24
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw24 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.CodePoint = func(_164_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw24 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw24).ArraySet1CodePoint(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw24).ArraySet1CodePoint(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_163_v59 = _nw24
	var _165_v60 D8
	_ = _165_v60
	_165_v60 = Companion_D8_.Create_DC16_(_90_v0)
	var _166_v61 D8
	_ = _166_v61
	_166_v61 = Companion_D8_.Create_DC18_(_165_v60)
	var _167_v62 _dafny.Sequence
	_ = _167_v62
	_167_v62 = _dafny.SeqOf(_163_v59)
	var _pat_let_tv0 = _100_v8
	_ = _pat_let_tv0
	var _pat_let_tv1 = _100_v8
	_ = _pat_let_tv1
	var _pat_let_tv2 = _154_v50
	_ = _pat_let_tv2
	var _rhs16 bool = func(_source5 D8) bool {
		if _source5.Is_DC17() {
			var _168___mcc_h3 _dafny.CodePoint = _source5.Get_().(D8_DC17).Cf30
			_ = _168___mcc_h3
			var _169___mcc_h4 _dafny.CodePoint = _source5.Get_().(D8_DC17).Cf31
			_ = _169___mcc_h4
			var _170___mcc_h5 _dafny.Int = _source5.Get_().(D8_DC17).Cf32
			_ = _170___mcc_h5
			var _171_cf32 _dafny.Int = _170___mcc_h5
			_ = _171_cf32
			var _172_cf31 _dafny.CodePoint = _169___mcc_h4
			_ = _172_cf31
			var _173_cf30 _dafny.CodePoint = _168___mcc_h3
			_ = _173_cf30
			return _pat_let_tv0
		} else if _source5.Is_DC16() {
			var _174___mcc_h6 _dafny.Sequence = _source5.Get_().(D8_DC16).Cf29
			_ = _174___mcc_h6
			var _175_cf29 _dafny.Sequence = _174___mcc_h6
			_ = _175_cf29
			return _pat_let_tv1
		} else {
			var _176___mcc_h7 D8 = _source5.Get_().(D8_DC18).Cf33
			_ = _176___mcc_h7
			var _177_cf33 D8 = _176___mcc_h7
			_ = _177_cf33
			return (_pat_let_tv2).F27()
		}
	}(_166_v61)
	_ = _rhs16
	var _rhs17 _dafny.Int = _99_v7
	_ = _rhs17
	var _rhs18 _dafny.Array = (_167_v62).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_158_v54).F25(), _dafny.IntOfInt64(-610)), _dafny.IntOfUint32((_167_v62).Cardinality()))).Uint32()).(_dafny.Array)
	_ = _rhs18
	var _lhs10 *GlobalState = _105_globalState
	_ = _lhs10
	_100_v8 = _rhs16
	_lhs10.F8 = _rhs17
	_163_v59 = _rhs18
	if (_154_v50).F27() {
		var _178_v63 _dafny.CodePoint
		_ = _178_v63
		_178_v63 = _dafny.CodePoint('a')
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))
		_ = _index10
		(_163_v59).ArraySet1CodePoint(_178_v63, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))
		_ = _index11
		var _rhs19 _dafny.Int = (_dafny.Zero).Minus((_117_v21).Minus((_158_v54).F25()))
		_ = _rhs19
		var _rhs20 _dafny.CodePoint = _178_v63
		_ = _rhs20
		var _rhs21 _dafny.Int = (_106_v12).Select((Companion_Default___.SafeIndex((_158_v54).F25(), _dafny.IntOfUint32((_106_v12).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs21
		var _lhs11 *GlobalState = _105_globalState
		_ = _lhs11
		var _lhs12 _dafny.Array = _163_v59
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))
		_ = _lhs13
		_lhs11.F8 = _rhs19
		(_lhs12).ArraySet1CodePoint(_rhs20, (_lhs13).Int())
		_99_v7 = _rhs21
		var _179_v64 _dafny.MultiSet
		_ = _179_v64
		_179_v64 = _dafny.MultiSetOf(!((_118_v22).F27()))
		var _180_v65 D8
		_ = _180_v65
		_180_v65 = Companion_D8_.Create_DC17_((_163_v59).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))).Int()), (_163_v59).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))).Int()), (_158_v54).F25())
		_114_v18 = Companion_Default___.SafeModuloInt((Companion_Default___.Fm18(_179_v64, _180_v65, _95_v3, _105_globalState)).Plus((_158_v54).F25()), (_dafny.Zero).Minus(_114_v18))
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(255), _dafny.ArrayLen((_163_v59), 0))
		_ = _index12
		(_163_v59).ArraySet1CodePoint(Companion_Default___.Fm8(_105_globalState), (_index12).Int())
		var _181_v66 _dafny.Array
		_ = _181_v66
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_2
		var _nw25 _dafny.Array
		_ = _nw25
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw25 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = func(_182_i7 _dafny.Int) bool {
				return false
			}
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw25 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw25).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw25).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_181_v66 = _nw25
		var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_181_v66), 0))
		_ = _index13
		(_181_v66).ArraySet1(_95_v3, (_index13).Int())
		var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.ArrayLen((_181_v66), 0))
		_ = _index14
		(_181_v66).ArraySet1(!(_157_v53).Contains((_158_v54).F25()), (_index14).Int())
		_95_v3 = _115_v19
	} else {
		var _183_v67 _dafny.Array
		_ = _183_v67
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_3
		var _nw26 _dafny.Array
		_ = _nw26
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw26 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) bool = (func(_184_v55 *C0) func(_dafny.Int) bool {
				return func(_185_i8 _dafny.Int) bool {
					return (_184_v55).F27()
				}
			})(_159_v55)
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
		_183_v67 = _nw26
		var _186_v68 _dafny.MultiSet
		_ = _186_v68
		_186_v68 = _dafny.MultiSetOf(_114_v18)
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_183_v67), 0))
		_ = _index15
		(_183_v67).ArraySet1((_186_v68).IsDisjointFrom(_186_v68), (_index15).Int())
		var _187_v69 _dafny.Sequence
		_ = _187_v69
		_187_v69 = _dafny.SeqOf(_90_v0)
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(867), _dafny.ArrayLen((_183_v67), 0))
		_ = _index16
		(_183_v67).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
			if (_118_v22).F27() {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}(func(_188_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))
			}
			return (_187_v69).Select((Companion_Default___.SafeIndex((_158_v54).F25(), _dafny.IntOfUint32((_187_v69).Cardinality()))).Uint32()).(_dafny.Sequence)
		})(), _90_v0), (_index16).Int())
		_106_v12 = _106_v12
		var _189_v70 _dafny.Set
		_ = _189_v70
		_189_v70 = _dafny.SetOf((_dafny.Zero).Minus((_157_v53).Cardinality()))
		_189_v70 = _189_v70
		var _190_v71 *C0
		_ = _190_v71
		var _nw27 *C0 = New_C0_()
		_ = _nw27
		_nw27.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v21, true)).Contains(Companion_Default___.Fm5(_105_globalState)))
		_190_v71 = _nw27
		(_105_globalState).F5 = !(!(((_dafny.Zero).Minus(((_158_v54).F25()).Plus(_99_v7))).Cmp((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_191_i10 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality())).Times(_117_v21)) > 0))
	}
	var _192_v72 _dafny.Array
	_ = _192_v72
	var _nwElement0_3 bool = !(_116_v20) || ((_154_v50).F27())
	_ = _nwElement0_3
	var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
	_ = _nw28
	(_nw28).ArraySet1(_nwElement0_3, 0)
	(_nw28).ArraySet1(!(Companion_Default___.Fm0(_107_v13, _105_globalState)), 1)
	(_nw28).ArraySet1(false, 2)
	(_nw28).ArraySet1(true, 3)
	(_nw28).ArraySet1(_115_v19, 4)
	(_nw28).ArraySet1(_95_v3, 5)
	_192_v72 = _nw28
	for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_192_v72), 0))); ; {
		_guard_loop_0, _ok15 := _iter15()
		if !_ok15 {
			break
		}
		var _193_i11 _dafny.Int
		_193_i11 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_193_i11).Sign() != -1) && ((_193_i11).Cmp(_dafny.ArrayLen((_192_v72), 0)) < 0)) {
			(_192_v72).ArraySet1((_118_v22).F27(), (_193_i11).Int())
		}
	}
	(_105_globalState).F22 = (_159_v55).F27()
	_dafny.Print(_90_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_91_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_94_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_95_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_96_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_97_v5).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_98_v6).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_99_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_100_v8))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_101_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(637), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_102_v10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_104_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("dwmjjp"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F0().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F1).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F2).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState.F7).ArrayGet1((_dafny.Zero).Int()).(_dafny.Set)).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F9()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(637), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F10().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F12()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F12()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_105_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_105_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_105_globalState.F13).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F16).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.UnicodeSeqOfUtf8Bytes("dwmjjp"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F17)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_105_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F23).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F23).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F23).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F23).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState.F23).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_globalState).F24().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_106_v12, _dafny.SeqOf(_dafny.IntOfInt64(637))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_107_v13, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_108_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(637)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_109_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l'), _dafny.CodePoint('l')), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_111_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_114_v18)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_115_v19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_116_v20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_117_v21)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v22).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_119_v23.F26).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_120_v24).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_121_v25).Dtor_cf49()).Dtor_cf48().F26).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_154_v50).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_155_v51).Dtor_cf41()).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_156_v52).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(637), Companion_D11_.Create_DC24_(_dafny.MultiSetOf(true)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_157_v53).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(375))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_158_v54).F25())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_159_v55).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v56).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false), true).UpdateUnsafe(_dafny.SeqOf(false, false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_161_v58, _dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(false), _dafny.SeqOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(8)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(9)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(10)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(11)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(12)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(13)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(14)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(15)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(16)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_163_v59).ArrayGet1CodePoint((_dafny.IntOfInt64(17)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v60).Dtor_cf29().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_166_v61).Dtor_cf33()).Dtor_cf29().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_167_v62).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v72).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
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

func (CompanionStruct_D0_) Default() bool {
	return false
}

func (_this D0) Dtor_cf0() bool {
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

type D1_DC1 struct {
	Cf1 _dafny.Int
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Int) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.Int {
	return _dafny.Zero
}

func (_this D1) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0
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

type D2_DC2 struct {
	Cf2 _dafny.Array
}

func (D2_DC2) isD2() {}

func (CompanionStruct_D2_) Create_DC2_(Cf2 _dafny.Array) D2 {
	return D2{D2_DC2{Cf2}}
}

func (_this D2) Is_DC2() bool {
	_, ok := _this.Get_().(D2_DC2)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D2_DC2).Cf2
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC2:
		{
			return "D2.DC2" + "(" + _dafny.String(data.Cf2) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC2:
		{
			data2, ok := other.Get_().(D2_DC2)
			return ok && data1.Cf2 == data2.Cf2
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

type D3_DC4 struct {
	Cf4 _dafny.Int
	Cf5 bool
	Cf6 bool
}

func (D3_DC4) isD3() {}

func (CompanionStruct_D3_) Create_DC4_(Cf4 _dafny.Int, Cf5 bool, Cf6 bool) D3 {
	return D3{D3_DC4{Cf4, Cf5, Cf6}}
}

func (_this D3) Is_DC4() bool {
	_, ok := _this.Get_().(D3_DC4)
	return ok
}

type D3_DC3 struct {
	Cf3 _dafny.Sequence
}

func (D3_DC3) isD3() {}

func (CompanionStruct_D3_) Create_DC3_(Cf3 _dafny.Sequence) D3 {
	return D3{D3_DC3{Cf3}}
}

func (_this D3) Is_DC3() bool {
	_, ok := _this.Get_().(D3_DC3)
	return ok
}

type D3_DC5 struct {
	Cf7 D3
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf7 D3) D3 {
	return D3{D3_DC5{Cf7}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC4_(_dafny.Zero, false, false)
}

func (_this D3) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D3_DC4).Cf4
}

func (_this D3) Dtor_cf5() bool {
	return _this.Get_().(D3_DC4).Cf5
}

func (_this D3) Dtor_cf6() bool {
	return _this.Get_().(D3_DC4).Cf6
}

func (_this D3) Dtor_cf3() _dafny.Sequence {
	return _this.Get_().(D3_DC3).Cf3
}

func (_this D3) Dtor_cf7() D3 {
	return _this.Get_().(D3_DC5).Cf7
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC4:
		{
			return "D3.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D3_DC3:
		{
			return "D3.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC4:
		{
			data2, ok := other.Get_().(D3_DC4)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6 == data2.Cf6
		}
	case D3_DC3:
		{
			data2, ok := other.Get_().(D3_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf7.Equals(data2.Cf7)
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
	Cf9  _dafny.Int
	Cf10 *C0
	Cf11 bool
	Cf12 _dafny.Int
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf9 _dafny.Int, Cf10 *C0, Cf11 bool, Cf12 _dafny.Int) D4 {
	return D4{D4_DC7{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

type D4_DC8 struct {
	Cf13 _dafny.CodePoint
	Cf14 bool
	Cf15 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf13 _dafny.CodePoint, Cf14 bool, Cf15 bool) D4 {
	return D4{D4_DC8{Cf13, Cf14, Cf15}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC6 struct {
	Cf8 _dafny.Set
}

func (D4_DC6) isD4() {}

func (CompanionStruct_D4_) Create_DC6_(Cf8 _dafny.Set) D4 {
	return D4{D4_DC6{Cf8}}
}

func (_this D4) Is_DC6() bool {
	_, ok := _this.Get_().(D4_DC6)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC7_(_dafny.Zero, (*C0)(nil), false, _dafny.Zero)
}

func (_this D4) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf9
}

func (_this D4) Dtor_cf10() *C0 {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) Dtor_cf11() bool {
	return _this.Get_().(D4_DC7).Cf11
}

func (_this D4) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D4_DC7).Cf12
}

func (_this D4) Dtor_cf13() _dafny.CodePoint {
	return _this.Get_().(D4_DC8).Cf13
}

func (_this D4) Dtor_cf14() bool {
	return _this.Get_().(D4_DC8).Cf14
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC8).Cf15
}

func (_this D4) Dtor_cf8() _dafny.Set {
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
			return "D4.DC8" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14 == data2.Cf14 && data1.Cf15 == data2.Cf15
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
	Cf16 _dafny.Array
}

func (D5_DC9) isD5() {}

func (CompanionStruct_D5_) Create_DC9_(Cf16 _dafny.Array) D5 {
	return D5{D5_DC9{Cf16}}
}

func (_this D5) Is_DC9() bool {
	_, ok := _this.Get_().(D5_DC9)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf16() _dafny.Array {
	return _this.Get_().(D5_DC9).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC9:
		{
			return "D5.DC9" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf16 == data2.Cf16
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
	Cf18 bool
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf18 bool) D6 {
	return D6{D6_DC11{Cf18}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC10 struct {
	Cf17 _dafny.Array
}

func (D6_DC10) isD6() {}

func (CompanionStruct_D6_) Create_DC10_(Cf17 _dafny.Array) D6 {
	return D6{D6_DC10{Cf17}}
}

func (_this D6) Is_DC10() bool {
	_, ok := _this.Get_().(D6_DC10)
	return ok
}

type D6_DC12 struct {
	Cf19 D6
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf19 D6) D6 {
	return D6{D6_DC12{Cf19}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC11_(false)
}

func (_this D6) Dtor_cf18() bool {
	return _this.Get_().(D6_DC11).Cf18
}

func (_this D6) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D6_DC10).Cf17
}

func (_this D6) Dtor_cf19() D6 {
	return _this.Get_().(D6_DC12).Cf19
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf18) + ")"
		}
	case D6_DC10:
		{
			return "D6.DC10" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf19) + ")"
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
			return ok && data1.Cf18 == data2.Cf18
		}
	case D6_DC10:
		{
			data2, ok := other.Get_().(D6_DC10)
			return ok && data1.Cf17 == data2.Cf17
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf19.Equals(data2.Cf19)
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

type D7_DC14 struct {
	Cf21 _dafny.Int
	Cf22 bool
	Cf23 bool
	Cf24 bool
	Cf25 _dafny.Int
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf21 _dafny.Int, Cf22 bool, Cf23 bool, Cf24 bool, Cf25 _dafny.Int) D7 {
	return D7{D7_DC14{Cf21, Cf22, Cf23, Cf24, Cf25}}
}

func (_this D7) Is_DC14() bool {
	_, ok := _this.Get_().(D7_DC14)
	return ok
}

type D7_DC15 struct {
	Cf26 _dafny.Int
	Cf27 bool
	Cf28 _dafny.Int
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf26 _dafny.Int, Cf27 bool, Cf28 _dafny.Int) D7 {
	return D7{D7_DC15{Cf26, Cf27, Cf28}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC13 struct {
	Cf20 _dafny.Map
}

func (D7_DC13) isD7() {}

func (CompanionStruct_D7_) Create_DC13_(Cf20 _dafny.Map) D7 {
	return D7{D7_DC13{Cf20}}
}

func (_this D7) Is_DC13() bool {
	_, ok := _this.Get_().(D7_DC13)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC14_(_dafny.Zero, false, false, false, _dafny.Zero)
}

func (_this D7) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf21
}

func (_this D7) Dtor_cf22() bool {
	return _this.Get_().(D7_DC14).Cf22
}

func (_this D7) Dtor_cf23() bool {
	return _this.Get_().(D7_DC14).Cf23
}

func (_this D7) Dtor_cf24() bool {
	return _this.Get_().(D7_DC14).Cf24
}

func (_this D7) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D7_DC14).Cf25
}

func (_this D7) Dtor_cf26() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf26
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC15).Cf27
}

func (_this D7) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D7_DC15).Cf28
}

func (_this D7) Dtor_cf20() _dafny.Map {
	return _this.Get_().(D7_DC13).Cf20
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC13:
		{
			return "D7.DC13" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23 && data1.Cf24 == data2.Cf24 && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf26.Cmp(data2.Cf26) == 0 && data1.Cf27 == data2.Cf27 && data1.Cf28.Cmp(data2.Cf28) == 0
		}
	case D7_DC13:
		{
			data2, ok := other.Get_().(D7_DC13)
			return ok && data1.Cf20.Equals(data2.Cf20)
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

type D8_DC17 struct {
	Cf30 _dafny.CodePoint
	Cf31 _dafny.CodePoint
	Cf32 _dafny.Int
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf30 _dafny.CodePoint, Cf31 _dafny.CodePoint, Cf32 _dafny.Int) D8 {
	return D8{D8_DC17{Cf30, Cf31, Cf32}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

type D8_DC16 struct {
	Cf29 _dafny.Sequence
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf29 _dafny.Sequence) D8 {
	return D8{D8_DC16{Cf29}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

type D8_DC18 struct {
	Cf33 D8
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf33 D8) D8 {
	return D8{D8_DC18{Cf33}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC17_(_dafny.CodePoint('D'), _dafny.CodePoint('D'), _dafny.Zero)
}

func (_this D8) Dtor_cf30() _dafny.CodePoint {
	return _this.Get_().(D8_DC17).Cf30
}

func (_this D8) Dtor_cf31() _dafny.CodePoint {
	return _this.Get_().(D8_DC17).Cf31
}

func (_this D8) Dtor_cf32() _dafny.Int {
	return _this.Get_().(D8_DC17).Cf32
}

func (_this D8) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D8_DC16).Cf29
}

func (_this D8) Dtor_cf33() D8 {
	return _this.Get_().(D8_DC18).Cf33
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC17:
		{
			return "D8.DC17" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D8_DC16:
		{
			return "D8.DC16" + "(" + data.Cf29.VerbatimString(true) + ")"
		}
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC17:
		{
			data2, ok := other.Get_().(D8_DC17)
			return ok && data1.Cf30 == data2.Cf30 && data1.Cf31 == data2.Cf31 && data1.Cf32.Cmp(data2.Cf32) == 0
		}
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf29.Equals(data2.Cf29)
		}
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf33.Equals(data2.Cf33)
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
	Cf35 bool
	Cf36 bool
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf35 bool, Cf36 bool) D9 {
	return D9{D9_DC20{Cf35, Cf36}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

type D9_DC19 struct {
	Cf34 _dafny.MultiSet
}

func (D9_DC19) isD9() {}

func (CompanionStruct_D9_) Create_DC19_(Cf34 _dafny.MultiSet) D9 {
	return D9{D9_DC19{Cf34}}
}

func (_this D9) Is_DC19() bool {
	_, ok := _this.Get_().(D9_DC19)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC20_(false, false)
}

func (_this D9) Dtor_cf35() bool {
	return _this.Get_().(D9_DC20).Cf35
}

func (_this D9) Dtor_cf36() bool {
	return _this.Get_().(D9_DC20).Cf36
}

func (_this D9) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D9_DC19).Cf34
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ")"
		}
	case D9_DC19:
		{
			return "D9.DC19" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf35 == data2.Cf35 && data1.Cf36 == data2.Cf36
		}
	case D9_DC19:
		{
			data2, ok := other.Get_().(D9_DC19)
			return ok && data1.Cf34.Equals(data2.Cf34)
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

type D10_DC22 struct {
	Cf38 _dafny.Int
	Cf39 bool
}

func (D10_DC22) isD10() {}

func (CompanionStruct_D10_) Create_DC22_(Cf38 _dafny.Int, Cf39 bool) D10 {
	return D10{D10_DC22{Cf38, Cf39}}
}

func (_this D10) Is_DC22() bool {
	_, ok := _this.Get_().(D10_DC22)
	return ok
}

type D10_DC23 struct {
	Cf40 _dafny.Int
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf40 _dafny.Int) D10 {
	return D10{D10_DC23{Cf40}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

type D10_DC21 struct {
	Cf37 _dafny.Map
}

func (D10_DC21) isD10() {}

func (CompanionStruct_D10_) Create_DC21_(Cf37 _dafny.Map) D10 {
	return D10{D10_DC21{Cf37}}
}

func (_this D10) Is_DC21() bool {
	_, ok := _this.Get_().(D10_DC21)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC22_(_dafny.Zero, false)
}

func (_this D10) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D10_DC22).Cf38
}

func (_this D10) Dtor_cf39() bool {
	return _this.Get_().(D10_DC22).Cf39
}

func (_this D10) Dtor_cf40() _dafny.Int {
	return _this.Get_().(D10_DC23).Cf40
}

func (_this D10) Dtor_cf37() _dafny.Map {
	return _this.Get_().(D10_DC21).Cf37
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC22:
		{
			return "D10.DC22" + "(" + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D10_DC21:
		{
			return "D10.DC21" + "(" + _dafny.String(data.Cf37) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC22:
		{
			data2, ok := other.Get_().(D10_DC22)
			return ok && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39
		}
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && data1.Cf40.Cmp(data2.Cf40) == 0
		}
	case D10_DC21:
		{
			data2, ok := other.Get_().(D10_DC21)
			return ok && data1.Cf37.Equals(data2.Cf37)
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
	Cf42 _dafny.Int
	Cf43 bool
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf42 _dafny.Int, Cf43 bool) D11 {
	return D11{D11_DC25{Cf42, Cf43}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC24 struct {
	Cf41 _dafny.MultiSet
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf41 _dafny.MultiSet) D11 {
	return D11{D11_DC24{Cf41}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC25_(_dafny.Zero, false)
}

func (_this D11) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D11_DC25).Cf42
}

func (_this D11) Dtor_cf43() bool {
	return _this.Get_().(D11_DC25).Cf43
}

func (_this D11) Dtor_cf41() _dafny.MultiSet {
	return _this.Get_().(D11_DC24).Cf41
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf41) + ")"
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
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43 == data2.Cf43
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D12_DC27 struct {
	Cf45 bool
	Cf46 _dafny.Int
}

func (D12_DC27) isD12() {}

func (CompanionStruct_D12_) Create_DC27_(Cf45 bool, Cf46 _dafny.Int) D12 {
	return D12{D12_DC27{Cf45, Cf46}}
}

func (_this D12) Is_DC27() bool {
	_, ok := _this.Get_().(D12_DC27)
	return ok
}

type D12_DC26 struct {
	Cf44 _dafny.Map
}

func (D12_DC26) isD12() {}

func (CompanionStruct_D12_) Create_DC26_(Cf44 _dafny.Map) D12 {
	return D12{D12_DC26{Cf44}}
}

func (_this D12) Is_DC26() bool {
	_, ok := _this.Get_().(D12_DC26)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC27_(false, _dafny.Zero)
}

func (_this D12) Dtor_cf45() bool {
	return _this.Get_().(D12_DC27).Cf45
}

func (_this D12) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D12_DC27).Cf46
}

func (_this D12) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D12_DC26).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC27:
		{
			return "D12.DC27" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D12_DC26:
		{
			return "D12.DC26" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC27:
		{
			data2, ok := other.Get_().(D12_DC27)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D12_DC26:
		{
			data2, ok := other.Get_().(D12_DC26)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D13_DC29 struct {
	Cf48 *C3
}

func (D13_DC29) isD13() {}

func (CompanionStruct_D13_) Create_DC29_(Cf48 *C3) D13 {
	return D13{D13_DC29{Cf48}}
}

func (_this D13) Is_DC29() bool {
	_, ok := _this.Get_().(D13_DC29)
	return ok
}

type D13_DC28 struct {
	Cf47 _dafny.Sequence
}

func (D13_DC28) isD13() {}

func (CompanionStruct_D13_) Create_DC28_(Cf47 _dafny.Sequence) D13 {
	return D13{D13_DC28{Cf47}}
}

func (_this D13) Is_DC28() bool {
	_, ok := _this.Get_().(D13_DC28)
	return ok
}

type D13_DC30 struct {
	Cf49 D13
}

func (D13_DC30) isD13() {}

func (CompanionStruct_D13_) Create_DC30_(Cf49 D13) D13 {
	return D13{D13_DC30{Cf49}}
}

func (_this D13) Is_DC30() bool {
	_, ok := _this.Get_().(D13_DC30)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC29_((*C3)(nil))
}

func (_this D13) Dtor_cf48() *C3 {
	return _this.Get_().(D13_DC29).Cf48
}

func (_this D13) Dtor_cf47() _dafny.Sequence {
	return _this.Get_().(D13_DC28).Cf47
}

func (_this D13) Dtor_cf49() D13 {
	return _this.Get_().(D13_DC30).Cf49
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC29:
		{
			return "D13.DC29" + "(" + _dafny.String(data.Cf48) + ")"
		}
	case D13_DC28:
		{
			return "D13.DC28" + "(" + _dafny.String(data.Cf47) + ")"
		}
	case D13_DC30:
		{
			return "D13.DC30" + "(" + _dafny.String(data.Cf49) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC29:
		{
			data2, ok := other.Get_().(D13_DC29)
			return ok && data1.Cf48 == data2.Cf48
		}
	case D13_DC28:
		{
			data2, ok := other.Get_().(D13_DC28)
			return ok && data1.Cf47.Equals(data2.Cf47)
		}
	case D13_DC30:
		{
			data2, ok := other.Get_().(D13_DC30)
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

type D14_DC31 struct {
	Cf50 _dafny.Set
}

func (D14_DC31) isD14() {}

func (CompanionStruct_D14_) Create_DC31_(Cf50 _dafny.Set) D14 {
	return D14{D14_DC31{Cf50}}
}

func (_this D14) Is_DC31() bool {
	_, ok := _this.Get_().(D14_DC31)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D14) Dtor_cf50() _dafny.Set {
	return _this.Get_().(D14_DC31).Cf50
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC31:
		{
			return "D14.DC31" + "(" + _dafny.String(data.Cf50) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC31:
		{
			data2, ok := other.Get_().(D14_DC31)
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

type D15_DC33 struct {
	Cf52 bool
	Cf53 bool
	Cf54 _dafny.Int
	Cf55 *C1
	Cf56 bool
}

func (D15_DC33) isD15() {}

func (CompanionStruct_D15_) Create_DC33_(Cf52 bool, Cf53 bool, Cf54 _dafny.Int, Cf55 *C1, Cf56 bool) D15 {
	return D15{D15_DC33{Cf52, Cf53, Cf54, Cf55, Cf56}}
}

func (_this D15) Is_DC33() bool {
	_, ok := _this.Get_().(D15_DC33)
	return ok
}

type D15_DC32 struct {
	Cf51 _dafny.MultiSet
}

func (D15_DC32) isD15() {}

func (CompanionStruct_D15_) Create_DC32_(Cf51 _dafny.MultiSet) D15 {
	return D15{D15_DC32{Cf51}}
}

func (_this D15) Is_DC32() bool {
	_, ok := _this.Get_().(D15_DC32)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC33_(false, false, _dafny.Zero, (*C1)(nil), false)
}

func (_this D15) Dtor_cf52() bool {
	return _this.Get_().(D15_DC33).Cf52
}

func (_this D15) Dtor_cf53() bool {
	return _this.Get_().(D15_DC33).Cf53
}

func (_this D15) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D15_DC33).Cf54
}

func (_this D15) Dtor_cf55() *C1 {
	return _this.Get_().(D15_DC33).Cf55
}

func (_this D15) Dtor_cf56() bool {
	return _this.Get_().(D15_DC33).Cf56
}

func (_this D15) Dtor_cf51() _dafny.MultiSet {
	return _this.Get_().(D15_DC32).Cf51
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC33:
		{
			return "D15.DC33" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ")"
		}
	case D15_DC32:
		{
			return "D15.DC32" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC33:
		{
			data2, ok := other.Get_().(D15_DC33)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56
		}
	case D15_DC32:
		{
			data2, ok := other.Get_().(D15_DC32)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D16_DC35 struct {
	Cf58 bool
	Cf59 _dafny.Int
	Cf60 _dafny.Int
}

func (D16_DC35) isD16() {}

func (CompanionStruct_D16_) Create_DC35_(Cf58 bool, Cf59 _dafny.Int, Cf60 _dafny.Int) D16 {
	return D16{D16_DC35{Cf58, Cf59, Cf60}}
}

func (_this D16) Is_DC35() bool {
	_, ok := _this.Get_().(D16_DC35)
	return ok
}

type D16_DC34 struct {
	Cf57 _dafny.Sequence
}

func (D16_DC34) isD16() {}

func (CompanionStruct_D16_) Create_DC34_(Cf57 _dafny.Sequence) D16 {
	return D16{D16_DC34{Cf57}}
}

func (_this D16) Is_DC34() bool {
	_, ok := _this.Get_().(D16_DC34)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC35_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D16) Dtor_cf58() bool {
	return _this.Get_().(D16_DC35).Cf58
}

func (_this D16) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf59
}

func (_this D16) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D16_DC35).Cf60
}

func (_this D16) Dtor_cf57() _dafny.Sequence {
	return _this.Get_().(D16_DC34).Cf57
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC35:
		{
			return "D16.DC35" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D16_DC34:
		{
			return "D16.DC34" + "(" + _dafny.String(data.Cf57) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC35:
		{
			data2, ok := other.Get_().(D16_DC35)
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D16_DC34:
		{
			data2, ok := other.Get_().(D16_DC34)
			return ok && data1.Cf57.Equals(data2.Cf57)
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
	M1(globalState *GlobalState)
	M2(globalState *GlobalState)
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
	F1   _dafny.Array
	F2   _dafny.Map
	F5   bool
	F7   _dafny.Array
	F8   _dafny.Int
	F13  _dafny.Array
	F16  _dafny.Map
	F17  bool
	F19  bool
	F22  bool
	F23  _dafny.Array
	_f0  _dafny.Sequence
	_f3  _dafny.CodePoint
	_f4  _dafny.Int
	_f6  _dafny.Int
	_f9  _dafny.Map
	_f10 _dafny.Sequence
	_f11 bool
	_f12 _dafny.Array
	_f14 _dafny.Int
	_f15 _dafny.Int
	_f18 _dafny.Int
	_f20 bool
	_f21 bool
	_f24 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F2 = _dafny.EmptyMap
	_this.F5 = false
	_this.F7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F8 = _dafny.Zero
	_this.F13 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F16 = _dafny.EmptyMap
	_this.F17 = false
	_this.F19 = false
	_this.F22 = false
	_this.F23 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.EmptySeq
	_this._f3 = _dafny.CodePoint('D')
	_this._f4 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f9 = _dafny.EmptyMap
	_this._f10 = _dafny.EmptySeq
	_this._f11 = false
	_this._f12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = _dafny.Zero
	_this._f15 = _dafny.Zero
	_this._f18 = _dafny.Zero
	_this._f20 = false
	_this._f21 = false
	_this._f24 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Sequence, f1 _dafny.Array, f2 _dafny.Map, f3 _dafny.CodePoint, f4 _dafny.Int, f5 bool, f6 _dafny.Int, f7 _dafny.Array, f8 _dafny.Int, f9 _dafny.Map, f10 _dafny.Sequence, f11 bool, f12 _dafny.Array, f13 _dafny.Array, f14 _dafny.Int, f15 _dafny.Int, f16 _dafny.Map, f17 bool, f18 _dafny.Int, f19 bool, f20 bool, f21 bool, f22 bool, f23 _dafny.Array, f24 _dafny.Sequence) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this).F8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this).F16 = f16
		(_this).F17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
		(_this).F23 = f23
		(_this)._f24 = f24
	}
}
func (_this *GlobalState) F0() _dafny.Sequence {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F3() _dafny.CodePoint {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F9() _dafny.Map {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Sequence {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() bool {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F12() _dafny.Array {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() _dafny.Int {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() _dafny.Int {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F18() _dafny.Int {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F20() bool {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() bool {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F24() _dafny.Sequence {
	{
		return _this._f24
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f27 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f27 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f27 bool) {
	{
		(_this)._f27 = f27
	}
}
func (_this *C0) F27() bool {
	{
		return _this._f27
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f28 _dafny.Map
	_f29 T1
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f28 = _dafny.EmptyMap
	_this._f29 = (T1)(nil)
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

func (_this *C1) Ctor__(f28 _dafny.Map, f29 T1) {
	{
		(_this)._f28 = f28
		(_this)._f29 = f29
	}
}
func (_this *C1) M1(globalState *GlobalState) {
	{
		var _194_v0 _dafny.Int
		_ = _194_v0
		_194_v0 = _dafny.IntOfInt64(141)
		(globalState).F8 = _194_v0
		var _195_v1 bool
		_ = _195_v1
		_195_v1 = false
		var _196_v2 _dafny.Sequence
		_ = _196_v2
		_196_v2 = _dafny.SeqOf(_195_v1)
		var _197_v3 _dafny.Array
		_ = _197_v3
		var _nw29 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw29
		_197_v3 = _nw29
		var _198_v4 _dafny.Map
		_ = _198_v4
		_198_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v0, _197_v3)
		var _199_v5 _dafny.MultiSet
		_ = _199_v5
		_199_v5 = _dafny.MultiSetOf(_195_v1, (_196_v2).Select((Companion_Default___.SafeIndex((_198_v4).Cardinality(), _dafny.IntOfUint32((_196_v2).Cardinality()))).Uint32()).(bool))
		var _200_v6 _dafny.CodePoint
		_ = _200_v6
		_200_v6 = _dafny.CodePoint('f')
		var _201_v7 D8
		_ = _201_v7
		_201_v7 = Companion_D8_.Create_DC17_(_200_v6, _200_v6, _194_v0)
		var _hi1 _dafny.Int = _dafny.IntOfInt64(937)
		_ = _hi1
		for _202_i0 := Companion_Default___.Fm18(_199_v5, _201_v7, _195_v1, globalState); _202_i0.Cmp(_hi1) < 0; _202_i0 = _202_i0.Plus(_dafny.One) {
			var _203_v8 _dafny.Int
			_ = _203_v8
			_203_v8 = _202_i0
			var _204_v9 _dafny.Map
			_ = _204_v9
			_204_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v0, _203_v8)
			var _205_v10 _dafny.Sequence
			_ = _205_v10
			_205_v10 = _dafny.SeqOf(_194_v0, _202_i0, (_204_v9).Cardinality(), _202_i0, _194_v0)
			var _206_v11 _dafny.Array
			_ = _206_v11
			var _nwElement0_4 _dafny.Int = _202_i0
			_ = _nwElement0_4
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(8))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_4, 0)
			(_nw30).ArraySet1(_194_v0, 1)
			(_nw30).ArraySet1(_dafny.IntOfUint32((_205_v10).Cardinality()), 2)
			(_nw30).ArraySet1(_202_i0, 3)
			(_nw30).ArraySet1(_202_i0, 4)
			(_nw30).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_195_v1, !(Companion_Default___.Fm0(Companion_Default___.Fm19(globalState), globalState)))).Cardinality(), 5)
			(_nw30).ArraySet1((_dafny.Zero).Minus(_194_v0), 6)
			(_nw30).ArraySet1(_194_v0, 7)
			_206_v11 = _nw30
			var _207_v12 _dafny.Map
			_ = _207_v12
			_207_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_195_v1) || (_195_v1), _206_v11)
			_207_v12 = (_207_v12).Update(_195_v1, _206_v11)
			(globalState).F8 = Companion_Default___.SafeModuloInt(_194_v0, _194_v0)
			(globalState).F8 = ((_202_i0).Minus(_194_v0)).Minus((_202_i0).Times(_202_i0))
			(globalState).F19 = _195_v1
		}
		_194_v0 = _194_v0
		(globalState).F8 = _194_v0
		var _208_v13 _dafny.Map
		_ = _208_v13
		_208_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v0, _194_v0)
		_208_v13 = (_208_v13).Update(((_dafny.Zero).Minus(_194_v0)).Times(_194_v0), _194_v0)
		var _209_v14 _dafny.Map
		_ = _209_v14
		_209_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_194_v0, _196_v2)
		_209_v14 = _209_v14
	}
}
func (_this *C1) M2(globalState *GlobalState) {
	{
		var _210_v0 bool
		_ = _210_v0
		_210_v0 = false
		var _211_v1 _dafny.Sequence
		_ = _211_v1
		_211_v1 = _dafny.SeqOf(_210_v0)
		var _212_v2 _dafny.Sequence
		_ = _212_v2
		_212_v2 = _dafny.SeqOf(_211_v1, _dafny.SeqOf(_210_v0), _211_v1)
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_212_v2, _dafny.Companion_Sequence_.Concatenate(_212_v2, _212_v2)) {
			var _213_v3 _dafny.CodePoint
			_ = _213_v3
			_213_v3 = _dafny.CodePoint('a')
			var _214_v4 _dafny.Int
			_ = _214_v4
			_214_v4 = _dafny.IntOfInt64(-243)
			var _215_v5 _dafny.Sequence
			_ = _215_v5
			_215_v5 = _dafny.SeqOf(_214_v4)
			var _216_v6 _dafny.Map
			_ = _216_v6
			_216_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v3, (_215_v5).Select((Companion_Default___.SafeIndex(_214_v4, _dafny.IntOfUint32((_215_v5).Cardinality()))).Uint32()).(_dafny.Int))
			_216_v6 = Companion_Default___.Fm20(_214_v4, globalState)
			var _217_v7 _dafny.Array
			_ = _217_v7
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_4
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = (func(_218_v4 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_219_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_219_i0, _218_v4)
					}
				})(_214_v4)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw31 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw31).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw31).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_217_v7 = _nw31
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_217_v7), 0))
			_ = _index17
			(_217_v7).ArraySet1(_214_v4, (_index17).Int())
			var _220_v8 _dafny.Map
			_ = _220_v8
			_220_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_214_v4, _214_v4)
			var _221_v9 D7
			_ = _221_v9
			_221_v9 = Companion_D7_.Create_DC14_(_214_v4, _210_v0, _210_v0, _210_v0, _214_v4)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_217_v7), 0))
			_ = _index18
			var _rhs22 _dafny.Int = ((func() _dafny.Int {
				if (_220_v8).Contains(_214_v4) {
					return (_220_v8).Get(_214_v4).(_dafny.Int)
				}
				return (_221_v9).Dtor_cf25()
			})()).Times(_214_v4)
			_ = _rhs22
			var _rhs23 _dafny.Int = _214_v4
			_ = _rhs23
			var _rhs24 _dafny.Array = _217_v7
			_ = _rhs24
			var _lhs14 _dafny.Array = _217_v7
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_217_v7), 0))
			_ = _lhs15
			var _lhs16 *GlobalState = globalState
			_ = _lhs16
			_214_v4 = _rhs22
			(_lhs14).ArraySet1(_rhs23, (_lhs15).Int())
			_lhs16.F23 = _rhs24
			var _222_v10 _dafny.MultiSet
			_ = _222_v10
			_222_v10 = _dafny.MultiSetOf(_213_v3, _213_v3, _213_v3)
			(globalState).F17 = !(_222_v10).Equals((_222_v10).Union(_dafny.MultiSetOf(_213_v3)))
			var _223_v11 D4
			_ = _223_v11
			_223_v11 = Companion_D4_.Create_DC8_(_213_v3, _210_v0, _210_v0)
			var _224_v12 _dafny.Array
			_ = _224_v12
			var _nwElement0_5 bool = _210_v0
			_ = _nwElement0_5
			var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(2))
			_ = _nw32
			(_nw32).ArraySet1(_nwElement0_5, 0)
			(_nw32).ArraySet1(_210_v0, 1)
			_224_v12 = _nw32
			var _225_v13 _dafny.Array
			_ = _225_v13
			_225_v13 = _224_v12
			var _226_v14 _dafny.Sequence
			_ = _226_v14
			_226_v14 = _dafny.SeqOf(_224_v12)
			var _227_v15 _dafny.Map
			_ = _227_v15
			_227_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v0, (_226_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.IntOfUint32((_226_v14).Cardinality()))).Uint32()).(_dafny.Array))
			var _228_v16 _dafny.Array
			_ = _228_v16
			var _nwElement0_6 _dafny.Array = _224_v12
			_ = _nwElement0_6
			var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(20))
			_ = _nw33
			(_nw33).ArraySet1(_nwElement0_6, 0)
			(_nw33).ArraySet1(_224_v12, 1)
			(_nw33).ArraySet1(_224_v12, 2)
			(_nw33).ArraySet1(_224_v12, 3)
			(_nw33).ArraySet1(_224_v12, 4)
			(_nw33).ArraySet1((_225_v13), 5)
			(_nw33).ArraySet1(_224_v12, 6)
			(_nw33).ArraySet1(_224_v12, 7)
			(_nw33).ArraySet1(_224_v12, 8)
			(_nw33).ArraySet1(_224_v12, 9)
			(_nw33).ArraySet1(_224_v12, 10)
			(_nw33).ArraySet1((func() _dafny.Array {
				if (_227_v15).Contains(_210_v0) {
					return (_227_v15).Get(_210_v0).(_dafny.Array)
				}
				return _224_v12
			})(), 11)
			(_nw33).ArraySet1(_224_v12, 12)
			(_nw33).ArraySet1(_224_v12, 13)
			(_nw33).ArraySet1(_224_v12, 14)
			(_nw33).ArraySet1(_224_v12, 15)
			(_nw33).ArraySet1(_224_v12, 16)
			(_nw33).ArraySet1((_226_v14).Select((Companion_Default___.SafeIndex(_214_v4, _dafny.IntOfUint32((_226_v14).Cardinality()))).Uint32()).(_dafny.Array), 17)
			(_nw33).ArraySet1(_224_v12, 18)
			(_nw33).ArraySet1(_224_v12, 19)
			_228_v16 = _nw33
			var _229_v17 _dafny.Array
			_ = _229_v17
			_229_v17 = _228_v16
			var _source6 _dafny.Array = (func() _dafny.Array {
				if (_223_v11).Dtor_cf14() {
					return _229_v17
				}
				return _229_v17
			})()
			_ = _source6
			var _230___mcc_h0 _dafny.Array = _source6
			_ = _230___mcc_h0
			var _231_cf16 _dafny.Array = _230___mcc_h0
			_ = _231_cf16
			(globalState).F17 = _210_v0
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_224_v12), 0))
			_ = _index19
			(_224_v12).ArraySet1(_210_v0, (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_224_v12), 0))
			_ = _index20
			(_224_v12).ArraySet1(_210_v0, (_index20).Int())
			var _232_v18 _dafny.Sequence
			_ = _232_v18
			_232_v18 = _dafny.UnicodeSeqOfUtf8Bytes("rd")
			_232_v18 = _dafny.Companion_Sequence_.Concatenate(_232_v18, _232_v18)
			_211_v1 = _211_v1
			var _233_v19 _dafny.Map
			_ = _233_v19
			_233_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v0, _dafny.IntOfUint32((_215_v5).Cardinality()))
			var _234_v20 _dafny.MultiSet
			_ = _234_v20
			_234_v20 = _dafny.MultiSetOf(_210_v0)
			var _235_v21 _dafny.Map
			_ = _235_v21
			_235_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v0, _234_v20)
			var _236_v22 D8
			_ = _236_v22
			_236_v22 = Companion_D8_.Create_DC17_(_213_v3, _213_v3, (_217_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_217_v7), 0))).Int()).(_dafny.Int))
			var _237_v23 _dafny.Map
			_ = _237_v23
			_237_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_233_v19).Merge(_233_v19), ((func() _dafny.MultiSet {
				if (_235_v21).Contains(_210_v0) {
					return (_235_v21).Get(_210_v0).(_dafny.MultiSet)
				}
				return _234_v20
			})()).IsDisjointFrom((_234_v20).Update(_210_v0, Companion_Default___.Abs(Companion_Default___.Fm18(_dafny.MultiSetFromSeq(_211_v1), _236_v22, _210_v0, globalState)))))
			_237_v23 = (_237_v23).Update((func() _dafny.Map {
				if !(_210_v0) {
					return _233_v19
				}
				return _233_v19
			})(), _210_v0)
		} else {
			var _238_v25 _dafny.Array
			_ = _238_v25
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_5
			var _nw34 _dafny.Array
			_ = _nw34
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw34 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = func(_239_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_239_i1, (func() _dafny.Map {
						var _coll15 = _dafny.NewMapBuilder()
						_ = _coll15
						for _iter16 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(673))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}(func(_240_i2 _dafny.Int) _dafny.Int {
							return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(149))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg21 _dafny.Int) interface{} {
									return coer21(arg21)
								}
							}(func(_241_i3 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('i')
							})), _dafny.UnicodeSeqOfUtf8Bytes("buof"))).Cardinality()
						}))).Elements()); ; {
							_compr_15, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _242_v24 _dafny.Int
							_242_v24 = interface{}(_compr_15).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(673))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg22 _dafny.Int) interface{} {
									return coer22(arg22)
								}
							}(func(_240_i2 _dafny.Int) _dafny.Int {
								return (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(149))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg23 _dafny.Int) interface{} {
										return coer23(arg23)
									}
								}(func(_241_i3 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('i')
								})), _dafny.UnicodeSeqOfUtf8Bytes("buof"))).Cardinality()
							})), _242_v24) {
								_coll15.Add((_242_v24).Times(_dafny.IntOfInt64(-364)), _dafny.UnicodeSeqOfUtf8Bytes("u"))
							}
						}
						return _coll15.ToMap()
					}()).Cardinality())
				}
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw34 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw34).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw34).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_238_v25 = _nw34
			var _243_v26 _dafny.Int
			_ = _243_v26
			_243_v26 = _dafny.IntOfInt64(581)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_238_v25), 0))
			_ = _index21
			(_238_v25).ArraySet1(_243_v26, (_index21).Int())
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_238_v25), 0))
			_ = _index22
			(_238_v25).ArraySet1((_243_v26).Plus((_dafny.Zero).Minus(_243_v26)), (_index22).Int())
			(globalState).F19 = _210_v0
			var _244_v27 _dafny.Sequence
			_ = _244_v27
			_244_v27 = _dafny.UnicodeSeqOfUtf8Bytes("wxrqlhdfx")
			var _245_v28 _dafny.Map
			_ = _245_v28
			_245_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_238_v25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(503), _dafny.ArrayLen((_238_v25), 0))).Int()).(_dafny.Int), _244_v27)
			_245_v28 = (_245_v28).Update(_243_v26, _244_v27)
			var _246_v29 D3
			_ = _246_v29
			_246_v29 = Companion_D3_.Create_DC4_(_243_v26, _210_v0, _210_v0)
			var _247_v30 _dafny.MultiSet
			_ = _247_v30
			_247_v30 = _dafny.MultiSetOf(_246_v29)
			(globalState).F17 = !(Companion_Default___.Fm0(_211_v1, globalState)) || (!((_247_v30).IsSubsetOf(_247_v30)))
			(globalState).F19 = _210_v0
		}
		var _248_v31 *C0
		_ = _248_v31
		var _nw35 *C0 = New_C0_()
		_ = _nw35
		_nw35.Ctor__(_210_v0)
		_248_v31 = _nw35
		var _249_v32 _dafny.MultiSet
		_ = _249_v32
		_249_v32 = _dafny.MultiSetOf(_248_v31, _248_v31)
		var _250_v33 *C0
		_ = _250_v33
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__(!(_249_v32).Contains(_248_v31))
		_250_v33 = _nw36
		var _251_v34 _dafny.Int
		_ = _251_v34
		_251_v34 = _dafny.IntOfInt64(169)
		var _252_v35 _dafny.Set
		_ = _252_v35
		_252_v35 = _dafny.SetOf(_251_v34)
		var _253_v36 _dafny.Map
		_ = _253_v36
		_253_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_252_v35).Intersection(_252_v35), _dafny.IntOfInt64(-526))
		_253_v36 = ((_253_v36).Merge(func() _dafny.Map {
			var _coll16 = _dafny.NewMapBuilder()
			_ = _coll16
			for _iter17 := _dafny.Iterate((_253_v36).Keys().Elements()); ; {
				_compr_16, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _254_v37 _dafny.Set
				_254_v37 = interface{}(_compr_16).(_dafny.Set)
				if (_253_v36).Contains(_254_v37) {
					_coll16.Add(_254_v37, (_dafny.Zero).Minus(_251_v34))
				}
			}
			return _coll16.ToMap()
		}())).Merge(_253_v36)
		var _255_v38 _dafny.Set
		_ = _255_v38
		_255_v38 = _dafny.SetOf(!((_248_v31).F27()))
		var _256_v39 _dafny.Sequence
		_ = _256_v39
		_256_v39 = _dafny.UnicodeSeqOfUtf8Bytes("nfyl")
		var _257_v40 _dafny.Sequence
		_ = _257_v40
		_257_v40 = _dafny.SeqOf(_256_v39)
		var _258_v41 _dafny.Map
		_ = _258_v41
		_258_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D3_.Create_DC3_(_257_v40))
		var _259_v42 _dafny.Map
		_ = _259_v42
		_259_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_251_v34), _210_v0)
		var _260_v44 _dafny.MultiSet
		_ = _260_v44
		_260_v44 = _dafny.MultiSetOf(_dafny.IntOfUint32((_256_v39).Cardinality()))
		var _261_v45 _dafny.Map
		_ = _261_v45
		_261_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_251_v34, _dafny.IntOfInt64(812))
		var _262_v46 _dafny.MultiSet
		_ = _262_v46
		_262_v46 = _dafny.MultiSetOf((_250_v33).F27())
		var _263_v47 _dafny.Sequence
		_ = _263_v47
		_263_v47 = _dafny.SeqOf((_262_v46).Cardinality())
		var _264_v48 _dafny.Array
		_ = _264_v48
		var _nwElement0_7 bool = false
		_ = _nwElement0_7
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(27))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_7, 0)
		(_nw37).ArraySet1(!(_dafny.SetOf((_250_v33).F27(), (_250_v33).F27(), (_250_v33).F27(), (_250_v33).F27())).Equals(_255_v38), 1)
		(_nw37).ArraySet1(!(_210_v0), 2)
		(_nw37).ArraySet1((_248_v31).F27(), 3)
		(_nw37).ArraySet1(Companion_Default___.Fm0(_211_v1, globalState), 4)
		(_nw37).ArraySet1((_250_v33).F27(), 5)
		(_nw37).ArraySet1(((_258_v41).Cardinality()).Cmp(_251_v34) < 0, 6)
		(_nw37).ArraySet1(_210_v0, 7)
		(_nw37).ArraySet1((_259_v42).Equals(_259_v42), 8)
		(_nw37).ArraySet1(_210_v0, 9)
		(_nw37).ArraySet1(_210_v0, 10)
		(_nw37).ArraySet1((func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter18 := _dafny.Iterate((_260_v44).Elements()); ; {
				_compr_17, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _265_v43 _dafny.Int
				_265_v43 = interface{}(_compr_17).(_dafny.Int)
				if (_260_v44).Contains(_265_v43) {
					_coll17.Add(Companion_Default___.SafeModuloInt(_265_v43, _251_v34), _251_v34)
				}
			}
			return _coll17.ToMap()
		}()).Equals(_261_v45), 11)
		(_nw37).ArraySet1((_211_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_263_v47).Cardinality()), _dafny.IntOfUint32((_211_v1).Cardinality()))).Uint32()).(bool), 12)
		(_nw37).ArraySet1(!((_248_v31).F27()), 13)
		(_nw37).ArraySet1((_248_v31).F27(), 14)
		(_nw37).ArraySet1((_250_v33).F27(), 15)
		(_nw37).ArraySet1((_250_v33).F27(), 16)
		(_nw37).ArraySet1((_248_v31).F27(), 17)
		(_nw37).ArraySet1((_250_v33).F27(), 18)
		(_nw37).ArraySet1(_210_v0, 19)
		(_nw37).ArraySet1((_250_v33).F27(), 20)
		(_nw37).ArraySet1(_210_v0, 21)
		(_nw37).ArraySet1((_248_v31).F27(), 22)
		(_nw37).ArraySet1((_250_v33).F27(), 23)
		(_nw37).ArraySet1((func() bool {
			if (_248_v31).F27() {
				return (_248_v31).F27()
			}
			return (_248_v31).F27()
		})(), 24)
		(_nw37).ArraySet1((_248_v31).F27(), 25)
		(_nw37).ArraySet1(_210_v0, 26)
		_264_v48 = _nw37
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))
		_ = _index23
		(_264_v48).ArraySet1((_250_v33).F27(), (_index23).Int())
		var _266_v49 _dafny.CodePoint
		_ = _266_v49
		_266_v49 = _dafny.CodePoint('o')
		var _267_v50 D8
		_ = _267_v50
		_267_v50 = Companion_D8_.Create_DC17_(_266_v49, _266_v49, _dafny.IntOfUint32((_256_v39).Cardinality()))
		var _268_v51 _dafny.Map
		_ = _268_v51
		_268_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_248_v31).F27(), Companion_Default___.Fm0(_dafny.SeqOf(!(false), _210_v0, (_250_v33).F27(), (_250_v33).F27(), (_248_v31).F27()), globalState))
		var _269_v52 _dafny.Map
		_ = _269_v52
		_269_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_259_v42).Contains((func() _dafny.Int {
				if (_260_v44).Contains(_251_v34) {
					return (_260_v44).Multiplicity(_251_v34)
				}
				return Companion_Default___.Fm18(_262_v46, _267_v50, (func() bool {
					if (_268_v51).Contains(_210_v0) {
						return (_268_v51).Get(_210_v0).(bool)
					}
					return (_248_v31).F27()
				})(), globalState)
			})()) {
				return (_259_v42).Get((func() _dafny.Int {
					if (_260_v44).Contains(_251_v34) {
						return (_260_v44).Multiplicity(_251_v34)
					}
					return Companion_Default___.Fm18(_262_v46, _267_v50, (func() bool {
						if (_268_v51).Contains(_210_v0) {
							return (_268_v51).Get(_210_v0).(bool)
						}
						return (_248_v31).F27()
					})(), globalState)
				})()).(bool)
			}
			return (_250_v33).F27()
		})(), _251_v34)
		var _270_v53 D7
		_ = _270_v53
		_270_v53 = Companion_D7_.Create_DC13_(_269_v52)
		var _pat_let_tv3 = _250_v33
		_ = _pat_let_tv3
		var _pat_let_tv4 = _210_v0
		_ = _pat_let_tv4
		var _pat_let_tv5 = _210_v0
		_ = _pat_let_tv5
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))
		_ = _index24
		(_264_v48).ArraySet1(func(_source7 D7) bool {
			if _source7.Is_DC14() {
				var _271___mcc_h1 _dafny.Int = _source7.Get_().(D7_DC14).Cf21
				_ = _271___mcc_h1
				var _272___mcc_h2 bool = _source7.Get_().(D7_DC14).Cf22
				_ = _272___mcc_h2
				var _273___mcc_h3 bool = _source7.Get_().(D7_DC14).Cf23
				_ = _273___mcc_h3
				var _274___mcc_h4 bool = _source7.Get_().(D7_DC14).Cf24
				_ = _274___mcc_h4
				var _275___mcc_h5 _dafny.Int = _source7.Get_().(D7_DC14).Cf25
				_ = _275___mcc_h5
				var _276_cf25 _dafny.Int = _275___mcc_h5
				_ = _276_cf25
				var _277_cf24 bool = _274___mcc_h4
				_ = _277_cf24
				var _278_cf23 bool = _273___mcc_h3
				_ = _278_cf23
				var _279_cf22 bool = _272___mcc_h2
				_ = _279_cf22
				var _280_cf21 _dafny.Int = _271___mcc_h1
				_ = _280_cf21
				return (_pat_let_tv3).F27()
			} else if _source7.Is_DC15() {
				var _281___mcc_h6 _dafny.Int = _source7.Get_().(D7_DC15).Cf26
				_ = _281___mcc_h6
				var _282___mcc_h7 bool = _source7.Get_().(D7_DC15).Cf27
				_ = _282___mcc_h7
				var _283___mcc_h8 _dafny.Int = _source7.Get_().(D7_DC15).Cf28
				_ = _283___mcc_h8
				var _284_cf28 _dafny.Int = _283___mcc_h8
				_ = _284_cf28
				var _285_cf27 bool = _282___mcc_h7
				_ = _285_cf27
				var _286_cf26 _dafny.Int = _281___mcc_h6
				_ = _286_cf26
				return _pat_let_tv4
			} else {
				var _287___mcc_h9 _dafny.Map = _source7.Get_().(D7_DC13).Cf20
				_ = _287___mcc_h9
				var _288_cf20 _dafny.Map = _287___mcc_h9
				_ = _288_cf20
				return _pat_let_tv5
			}
		}(_270_v53), (_index24).Int())
		var _source8 D8 = _267_v50
		_ = _source8
		if _source8.Is_DC17() {
			var _289___mcc_h10 _dafny.CodePoint = _source8.Get_().(D8_DC17).Cf30
			_ = _289___mcc_h10
			var _290___mcc_h11 _dafny.CodePoint = _source8.Get_().(D8_DC17).Cf31
			_ = _290___mcc_h11
			var _291___mcc_h12 _dafny.Int = _source8.Get_().(D8_DC17).Cf32
			_ = _291___mcc_h12
			var _292_cf32 _dafny.Int = _291___mcc_h12
			_ = _292_cf32
			var _293_cf31 _dafny.CodePoint = _290___mcc_h11
			_ = _293_cf31
			var _294_cf30 _dafny.CodePoint = _289___mcc_h10
			_ = _294_cf30
			var _295_v54 _dafny.Int
			_ = _295_v54
			var _296_v55 bool
			_ = _296_v55
			var _297_v56 bool
			_ = _297_v56
			var _298_v57 _dafny.Int
			_ = _298_v57
			var _out8 _dafny.Int
			_ = _out8
			var _out9 bool
			_ = _out9
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0(globalState)
			_295_v54 = _out8
			_296_v55 = _out9
			_297_v56 = _out10
			_298_v57 = _out11
			(globalState).F8 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_295_v54, _298_v57), _292_cf32)
			if (_250_v33).F27() {
				_292_cf32 = (_dafny.Zero).Minus((_dafny.Zero).Minus(_295_v54))
				var _299_v58 _dafny.Set
				_ = _299_v58
				_299_v58 = _dafny.SetOf(Companion_Default___.Fm21(Companion_Default___.Fm22((_250_v33).F27(), globalState), _295_v54, globalState))
				var _300_v59 _dafny.Sequence
				_ = _300_v59
				_300_v59 = _dafny.SeqOf(_299_v58)
				var _301_v60 *C0
				_ = _301_v60
				var _nw38 *C0 = New_C0_()
				_ = _nw38
				_nw38.Ctor__(!(((_300_v59).Select((Companion_Default___.SafeIndex(_292_cf32, _dafny.IntOfUint32((_300_v59).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf(_299_v58)))
				_301_v60 = _nw38
				var _302_v61 _dafny.Int
				_ = _302_v61
				var _303_v62 bool
				_ = _303_v62
				var _304_v63 bool
				_ = _304_v63
				var _305_v64 _dafny.Int
				_ = _305_v64
				var _out12 _dafny.Int
				_ = _out12
				var _out13 bool
				_ = _out13
				var _out14 bool
				_ = _out14
				var _out15 _dafny.Int
				_ = _out15
				_out12, _out13, _out14, _out15 = Companion_Default___.M0(globalState)
				_302_v61 = _out12
				_303_v62 = _out13
				_304_v63 = _out14
				_305_v64 = _out15
				var _306_v65 D7
				_ = _306_v65
				_306_v65 = Companion_D7_.Create_DC15_(_305_v64, _296_v55, _305_v64)
				var _307_v66 _dafny.Array
				_ = _307_v66
				var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
				_ = _nw39
				_307_v66 = _nw39
				var _308_v67 _dafny.Sequence
				_ = _308_v67
				_308_v67 = _dafny.SeqOf(_307_v66)
				var _309_v68 _dafny.Sequence
				_ = _309_v68
				_309_v68 = _dafny.SeqOf(_308_v67)
				var _310_v69 _dafny.Array
				_ = _310_v69
				var _nwElement0_8 _dafny.Int = _305_v64
				_ = _nwElement0_8
				var _nw40 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(22))
				_ = _nw40
				(_nw40).ArraySet1(_nwElement0_8, 0)
				(_nw40).ArraySet1(((func() _dafny.Map {
					if (_306_v65).Dtor_cf27() {
						return _259_v42
					}
					return _259_v42
				})()).Cardinality(), 1)
				(_nw40).ArraySet1(_302_v61, 2)
				(_nw40).ArraySet1(_298_v57, 3)
				(_nw40).ArraySet1(_251_v34, 4)
				(_nw40).ArraySet1(_305_v64, 5)
				(_nw40).ArraySet1(Companion_Default___.SafeModuloInt(_305_v64, _295_v54), 6)
				(_nw40).ArraySet1(_302_v61, 7)
				(_nw40).ArraySet1(_302_v61, 8)
				(_nw40).ArraySet1(_292_cf32, 9)
				(_nw40).ArraySet1((_295_v54).Minus(_305_v64), 10)
				(_nw40).ArraySet1(_305_v64, 11)
				(_nw40).ArraySet1((_261_v45).Cardinality(), 12)
				(_nw40).ArraySet1(_298_v57, 13)
				(_nw40).ArraySet1((func() _dafny.Int {
					if (_250_v33).F27() {
						return _dafny.IntOfUint32((_256_v39).Cardinality())
					}
					return _dafny.IntOfInt64(909)
				})(), 14)
				(_nw40).ArraySet1(_295_v54, 15)
				(_nw40).ArraySet1(_dafny.IntOfUint32(((_309_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.IntOfUint32((_309_v68).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 16)
				(_nw40).ArraySet1((func() _dafny.Int {
					if _210_v0 {
						return _302_v61
					}
					return Companion_Default___.Fm18(_262_v46, Companion_D8_.Create_DC17_(_294_cf30, _266_v49, _dafny.IntOfInt64(435)), true, globalState)
				})(), 17)
				(_nw40).ArraySet1(_292_cf32, 18)
				(_nw40).ArraySet1(_302_v61, 19)
				(_nw40).ArraySet1(_dafny.IntOfInt64(226), 20)
				(_nw40).ArraySet1(_295_v54, 21)
				_310_v69 = _nw40
				var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_307_v66), 0))
				_ = _index25
				(_307_v66).ArraySet1(_302_v61, (_index25).Int())
				var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_307_v66), 0))
				_ = _index26
				var _rhs25 bool = (_305_v64).Cmp(_302_v61) >= 0
				_ = _rhs25
				var _rhs26 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_268_v51).Cardinality(), (func() _dafny.Int {
					if (_261_v45).Contains(_292_cf32) {
						return (_261_v45).Get(_292_cf32).(_dafny.Int)
					}
					return _298_v57
				})()), ((_dafny.Zero).Minus(_251_v34)).Times(Companion_Default___.Fm18(_262_v46, _267_v50, (_248_v31).F27(), globalState)))
				_ = _rhs26
				var _lhs17 *GlobalState = globalState
				_ = _lhs17
				var _lhs18 _dafny.Array = _307_v66
				_ = _lhs18
				var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(651), _dafny.ArrayLen((_307_v66), 0))
				_ = _lhs19
				_lhs17.F22 = _rhs25
				(_lhs18).ArraySet1(_rhs26, (_lhs19).Int())
				_302_v61 = (_dafny.Zero).Minus(_295_v54)
			} else {
				(globalState).F5 = !((_dafny.MultiSetOf((_248_v31).F27(), (_211_v1).Select((Companion_Default___.SafeIndex(_298_v57, _dafny.IntOfUint32((_211_v1).Cardinality()))).Uint32()).(bool), (_250_v33).F27(), (_264_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))).Int()).(bool), (_248_v31).F27())).IsDisjointFrom((_262_v46).Intersection(_262_v46)))
				var _311_v70 _dafny.MultiSet
				_ = _311_v70
				_311_v70 = _dafny.MultiSetOf(_269_v52)
				_251_v34 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_248_v31).F27() {
						return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_269_v52, _269_v52, _269_v52, _269_v52, _269_v52))).Union(_311_v70)).Cardinality()
					}
					return _292_cf32
				})())
				_251_v34 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-796), (_251_v34).Plus(_295_v54))
				_269_v52 = (_269_v52).Update((_250_v33).F27(), Companion_Default___.SafeModuloInt(_295_v54, (_269_v52).Cardinality()))
				var _312_v71 _dafny.Array
				_ = _312_v71
				var _nwElement0_9 _dafny.Array = _264_v48
				_ = _nwElement0_9
				var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(2))
				_ = _nw41
				(_nw41).ArraySet1(_nwElement0_9, 0)
				(_nw41).ArraySet1(_264_v48, 1)
				_312_v71 = _nw41
				var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_312_v71), 0))
				_ = _index27
				(_312_v71).ArraySet1(_264_v48, (_index27).Int())
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))
				_ = _index28
				var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_312_v71), 0))
				_ = _index29
				var _rhs27 bool = (_211_v1).Select((Companion_Default___.SafeIndex(_251_v34, _dafny.IntOfUint32((_211_v1).Cardinality()))).Uint32()).(bool)
				_ = _rhs27
				var _rhs28 _dafny.Array = _264_v48
				_ = _rhs28
				var _rhs29 bool = !((_248_v31).F27())
				_ = _rhs29
				var _rhs30 bool = _296_v55
				_ = _rhs30
				var _lhs20 _dafny.Array = _264_v48
				_ = _lhs20
				var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))
				_ = _lhs21
				var _lhs22 _dafny.Array = _312_v71
				_ = _lhs22
				var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_312_v71), 0))
				_ = _lhs23
				var _lhs24 *GlobalState = globalState
				_ = _lhs24
				(_lhs20).ArraySet1(_rhs27, (_lhs21).Int())
				(_lhs22).ArraySet1(_rhs28, (_lhs23).Int())
				_lhs24.F22 = _rhs29
				_210_v0 = _rhs30
			}
			var _313_v72 *C0
			_ = _313_v72
			var _nw42 *C0 = New_C0_()
			_ = _nw42
			_nw42.Ctor__((_248_v31).F27())
			_313_v72 = _nw42
		} else if _source8.Is_DC16() {
			var _314___mcc_h13 _dafny.Sequence = _source8.Get_().(D8_DC16).Cf29
			_ = _314___mcc_h13
			var _315_cf29 _dafny.Sequence = _314___mcc_h13
			_ = _315_cf29
			(globalState).F5 = _210_v0
			if ((func() _dafny.Int {
				if (_260_v44).Contains(_dafny.IntOfUint32((_211_v1).Cardinality())) {
					return (_260_v44).Multiplicity(_dafny.IntOfUint32((_211_v1).Cardinality()))
				}
				return (Companion_Default___.Fm23((_248_v31).F27(), _251_v34, _210_v0, globalState)).Cardinality()
			})()).Cmp(_251_v34) != 0 {
				(globalState).F8 = (func() _dafny.Int {
					if (_261_v45).Contains((_251_v34).Minus(_251_v34)) {
						return (_261_v45).Get((_251_v34).Minus(_251_v34)).(_dafny.Int)
					}
					return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_250_v33).F27(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rclqacbxk")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_264_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))).Int()).(bool)), _251_v34))).Cardinality()
				})()
				_256_v39 = _315_cf29
				var _316_v73 _dafny.MultiSet
				_ = _316_v73
				_316_v73 = _dafny.MultiSetOf(_264_v48)
				_316_v73 = _316_v73
				var _317_v74 _dafny.Array
				_ = _317_v74
				var _nw43 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(19))
				_ = _nw43
				_317_v74 = _nw43
				var _318_v75 _dafny.Sequence
				_ = _318_v75
				_318_v75 = _dafny.SeqOf(_317_v74, _317_v74)
				var _319_v76 _dafny.MultiSet
				_ = _319_v76
				_319_v76 = _dafny.MultiSetOf((_318_v75).Select((Companion_Default___.SafeIndex(_251_v34, _dafny.IntOfUint32((_318_v75).Cardinality()))).Uint32()).(_dafny.Array))
				var _320_v77 _dafny.Map
				_ = _320_v77
				_320_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_319_v76).Union(_319_v76), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_251_v34, _251_v34)))
				_320_v77 = (_320_v77).Update((_319_v76).Intersection(_dafny.MultiSetOf(_317_v74, _317_v74, _317_v74)), (func() _dafny.Int {
					if (_262_v46).Contains((_248_v31).F27()) {
						return (_262_v46).Multiplicity((_248_v31).F27())
					}
					return ((_dafny.MultiSetFromSeq(Companion_Default___.Fm24((_250_v33).F27(), globalState))).Update(func() _dafny.Set {
						var _coll18 = _dafny.NewBuilder()
						_ = _coll18
						for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(905), _dafny.IntOfInt64(5))); ; {
							_compr_18, _ok19 := _iter19()
							if !_ok19 {
								break
							}
							var _321_v78 _dafny.Int
							_321_v78 = interface{}(_compr_18).(_dafny.Int)
							if ((_dafny.IntOfInt64(905)).Cmp(_321_v78) <= 0) && ((_321_v78).Cmp(_dafny.IntOfInt64(5)) < 0) {
								_coll18.Add((_321_v78).Plus(_251_v34))
							}
						}
						return _coll18.ToSet()
					}(), Companion_Default___.Abs(_251_v34))).Cardinality()
				})())
				var _322_v79 _dafny.Array
				_ = _322_v79
				var _nw44 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
				_ = _nw44
				_322_v79 = _nw44
				var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_322_v79), 0))
				_ = _index30
				(_322_v79).ArraySet1(_250_v33, (_index30).Int())
				var _323_v80 _dafny.Array
				_ = _323_v80
				var _len0_6 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_6
				var _nw45 _dafny.Array
				_ = _nw45
				if _len0_6.Cmp(_dafny.Zero) == 0 {
					_nw45 = _dafny.NewArray(_len0_6)
				} else {
					var _init6 func(_dafny.Int) _dafny.Int = (func(_324_v34 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_325_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_325_i4, _324_v34)
						}
					})(_251_v34)
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
				_323_v80 = _nw45
				var _326_v81 _dafny.Sequence
				_ = _326_v81
				_326_v81 = _dafny.SeqOf(_323_v80)
				var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_322_v79), 0))
				_ = _index31
				var _nw46 *C0 = New_C0_()
				_ = _nw46
				_nw46.Ctor__(_dafny.Companion_Sequence_.Equal(_326_v81, _326_v81))
				(_322_v79).ArraySet1(_nw46, (_index31).Int())
			} else {
				var _327_v82 D4
				_ = _327_v82
				_327_v82 = Companion_D4_.Create_DC8_(_266_v49, (_248_v31).F27(), (_248_v31).F27())
				_266_v49 = (func() _dafny.CodePoint {
					if (func() bool {
						if (_250_v33).F27() {
							return (_250_v33).F27()
						}
						return false
					})() {
						return (_327_v82).Dtor_cf13()
					}
					return _266_v49
				})()
				(globalState).F5 = (_248_v31).F27()
				var _328_v83 _dafny.Set
				_ = _328_v83
				_328_v83 = _dafny.SetOf(_263_v47, _263_v47)
				_268_v51 = (_268_v51).Update((_250_v33).F27(), (_dafny.SetOf(_263_v47)).IsProperSubsetOf(_328_v83))
				var _329_v84 _dafny.Map
				_ = _329_v84
				_329_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v1, (_250_v33).F27())
				_329_v84 = (_329_v84).Update(_211_v1, !(!(_262_v46).Contains((_248_v31).F27())))
				var _330_v85 D3
				_ = _330_v85
				_330_v85 = Companion_D3_.Create_DC3_(_dafny.SeqOf(_315_cf29))
				var _331_v86 D3
				_ = _331_v86
				_331_v86 = Companion_D3_.Create_DC5_((func() D3 {
					if true {
						return Companion_Default___.Fm25(globalState)
					}
					return _330_v85
				})())
				var _pat_let_tv6 = _330_v85
				_ = _pat_let_tv6
				_331_v86 = func(_pat_let0_0 D3) D3 {
					return func(_332_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let1_0 D3) D3 {
							return func(_333_dt__update_hcf7_h0 D3) D3 {
								return Companion_D3_.Create_DC5_(_333_dt__update_hcf7_h0)
							}(_pat_let1_0)
						}(_pat_let_tv6)
					}(_pat_let0_0)
				}(_331_v86)
			}
			(globalState).F17 = (_211_v1).Select((Companion_Default___.SafeIndex(_251_v34, _dafny.IntOfUint32((_211_v1).Cardinality()))).Uint32()).(bool)
			var _rhs31 _dafny.Sequence = _211_v1
			_ = _rhs31
			var _rhs32 _dafny.Int = (_dafny.Zero).Minus((_268_v51).Cardinality())
			_ = _rhs32
			var _lhs25 *GlobalState = globalState
			_ = _lhs25
			_211_v1 = _rhs31
			_lhs25.F8 = _rhs32
		} else {
			var _334___mcc_h14 D8 = _source8.Get_().(D8_DC18).Cf33
			_ = _334___mcc_h14
			var _335_cf33 D8 = _334___mcc_h14
			_ = _335_cf33
			_264_v48 = _264_v48
			var _336_v87 _dafny.Array
			_ = _336_v87
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
			_ = _nw47
			_336_v87 = _nw47
			var _337_v88 _dafny.Map
			_ = _337_v88
			_337_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v31, (_this).F29())
			var _338_v89 _dafny.Array
			_ = _338_v89
			var _nwElement0_10 _dafny.Int = _251_v34
			_ = _nwElement0_10
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(15))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_10, 0)
			(_nw48).ArraySet1(_dafny.IntOfInt64(64), 1)
			(_nw48).ArraySet1(_251_v34, 2)
			(_nw48).ArraySet1((_dafny.Zero).Minus(_251_v34), 3)
			(_nw48).ArraySet1(_251_v34, 4)
			(_nw48).ArraySet1(_251_v34, 5)
			(_nw48).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_211_v1, _211_v1)).Cardinality())), 6)
			(_nw48).ArraySet1(_251_v34, 7)
			(_nw48).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_251_v34, _251_v34)), 8)
			(_nw48).ArraySet1((_dafny.IntOfInt64(-983)).Times((_337_v88).Cardinality()), 9)
			(_nw48).ArraySet1(_251_v34, 10)
			(_nw48).ArraySet1(_251_v34, 11)
			(_nw48).ArraySet1(Companion_Default___.SafeDivisionInt(_251_v34, _251_v34), 12)
			(_nw48).ArraySet1(_251_v34, 13)
			(_nw48).ArraySet1((_261_v45).Cardinality(), 14)
			_338_v89 = _nw48
			var _rhs33 _dafny.Int = _251_v34
			_ = _rhs33
			var _rhs34 _dafny.Array = _338_v89
			_ = _rhs34
			var _rhs35 _dafny.Array = _336_v87
			_ = _rhs35
			var _lhs26 *GlobalState = globalState
			_ = _lhs26
			var _lhs27 *GlobalState = globalState
			_ = _lhs27
			_lhs26.F8 = _rhs33
			_lhs27.F23 = _rhs34
			_336_v87 = _rhs35
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_338_v89), 0))
			_ = _index32
			(_338_v89).ArraySet1((_251_v34).Times(_251_v34), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_338_v89), 0))
			_ = _index33
			(_338_v89).ArraySet1((Companion_Default___.Fm18(_262_v46, _267_v50, _210_v0, globalState)).Plus(_dafny.IntOfInt64(-785)), (_index33).Int())
			_256_v39 = _256_v39
		}
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_264_v48), 0))
		_ = _index34
		(_264_v48).ArraySet1(!(Companion_Default___.Fm0(_211_v1, globalState)), (_index34).Int())
	}
}
func (_this *C1) F28() _dafny.Map {
	{
		return _this._f28
	}
}
func (_this *C1) F29() T1 {
	{
		return _this._f29
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	dummy byte
}

func New_C2_() *C2 {
	_this := C2{}

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

func (_this *C2) Ctor__() {
	{
	}
}
func (_this *C2) M1(globalState *GlobalState) {
	{
		var _339_v0 _dafny.Int
		_ = _339_v0
		_339_v0 = _dafny.IntOfInt64(420)
		var _340_v1 _dafny.Sequence
		_ = _340_v1
		_340_v1 = _dafny.SeqOf(_339_v0)
		var _341_v2 bool
		_ = _341_v2
		_341_v2 = true
		var _342_v3 D3
		_ = _342_v3
		_342_v3 = Companion_D3_.Create_DC4_(_dafny.IntOfUint32((_340_v1).Cardinality()), _341_v2, _341_v2)
		if ((_342_v3).Dtor_cf4()).Cmp((_dafny.Zero).Minus(_339_v0)) >= 0 {
			var _343_v4 _dafny.Sequence
			_ = _343_v4
			_343_v4 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			var _344_v5 _dafny.Sequence
			_ = _344_v5
			_344_v5 = _dafny.SeqOf(_343_v4)
			var _345_v6 D3
			_ = _345_v6
			_345_v6 = Companion_D3_.Create_DC3_(_344_v5)
			var _346_v7 D3
			_ = _346_v7
			_346_v7 = Companion_D3_.Create_DC5_(_345_v6)
			var _347_v8 D3
			_ = _347_v8
			_347_v8 = Companion_D3_.Create_DC5_(_345_v6)
			var _348_v9 _dafny.Map
			_ = _348_v9
			_348_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v2, _339_v0)
			var _349_v10 _dafny.Map
			_ = _349_v10
			_349_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_347_v8, ((_348_v9).Update(_341_v2, _339_v0)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v2, _339_v0)))
			var _350_v11 _dafny.CodePoint
			_ = _350_v11
			_350_v11 = _dafny.CodePoint('k')
			var _351_v12 _dafny.Map
			_ = _351_v12
			_351_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v0, _341_v2)
			_349_v10 = (_349_v10).Update(Companion_Default___.Fm3(_350_v11, _dafny.SeqOf((_351_v12).Cardinality()), globalState), _341_v2)
			var _352_v13 _dafny.MultiSet
			_ = _352_v13
			_352_v13 = _dafny.MultiSetOf(_341_v2, _341_v2, _341_v2)
			var _353_v14 _dafny.Sequence
			_ = _353_v14
			_353_v14 = _dafny.SeqOf(_352_v13, _352_v13, (_352_v13).Union(_352_v13))
			_353_v14 = _353_v14
			var _354_v15 _dafny.Map
			_ = _354_v15
			_354_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v2, _343_v4)
			var _355_v16 _dafny.Map
			_ = _355_v16
			_355_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v2, _352_v13)
			var _356_v17 _dafny.Array
			_ = _356_v17
			var _nwElement0_11 _dafny.Int = _dafny.IntOfInt64(-502)
			_ = _nwElement0_11
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(4))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_11, 0)
			(_nw49).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_339_v0, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_354_v15).Contains(true) {
					return (_354_v15).Get(true).(_dafny.Sequence)
				}
				return _343_v4
			})()).Cardinality())))), 1)
			(_nw49).ArraySet1(_dafny.IntOfInt64(124), 2)
			(_nw49).ArraySet1(((func() _dafny.Map {
				if _341_v2 {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v2, _dafny.MultiSetOf(_341_v2))
				}
				return _355_v16
			})()).Cardinality(), 3)
			_356_v17 = _nw49
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_356_v17), 0))
			_ = _index35
			(_356_v17).ArraySet1(_339_v0, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_356_v17), 0))
			_ = _index36
			(_356_v17).ArraySet1(Companion_Default___.SafeDivisionInt(_339_v0, _339_v0), (_index36).Int())
			var _357_v18 T1
			_ = _357_v18
			var _nw50 *C0 = New_C0_()
			_ = _nw50
			_nw50.Ctor__(_341_v2)
			_357_v18 = _nw50
			_357_v18 = _357_v18
			_339_v0 = (_356_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_356_v17), 0))).Int()).(_dafny.Int)
		} else {
			(globalState).F5 = (_339_v0).Cmp((_339_v0).Minus((Companion_Default___.Fm4(_341_v2, globalState)).Cardinality())) >= 0
			var _358_v19 D3
			_ = _358_v19
			_358_v19 = Companion_D3_.Create_DC4_(_339_v0, _341_v2, _341_v2)
			var _359_v20 D3
			_ = _359_v20
			_359_v20 = Companion_D3_.Create_DC5_(_358_v19)
			var _360_v21 _dafny.Map
			_ = _360_v21
			_360_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_359_v20, _341_v2)
			_360_v21 = (_360_v21).Update(Companion_D3_.Create_DC5_(_358_v19), _dafny.Companion_Sequence_.IsProperPrefixOf(_340_v1, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_361_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_362_i0 _dafny.Int) _dafny.Int {
					return _361_v0
				}
			})(_339_v0)))).Cardinality()), _339_v0)))
			var _363_v22 _dafny.Set
			_ = _363_v22
			_363_v22 = _dafny.SetOf(_339_v0)
			var _364_v23 D4
			_ = _364_v23
			_364_v23 = Companion_D4_.Create_DC6_(_363_v22)
			var _365_v24 *C0
			_ = _365_v24
			var _nw51 *C0 = New_C0_()
			_ = _nw51
			_nw51.Ctor__(_341_v2)
			_365_v24 = _nw51
			var _366_v25 _dafny.Sequence
			_ = _366_v25
			_366_v25 = _dafny.UnicodeSeqOfUtf8Bytes("qky")
			var _367_v26 _dafny.Array
			_ = _367_v26
			var _nwElement0_12 _dafny.Int = _339_v0
			_ = _nwElement0_12
			var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(5))
			_ = _nw52
			(_nw52).ArraySet1(_nwElement0_12, 0)
			(_nw52).ArraySet1((Companion_D4_.Create_DC7_(_339_v0, _365_v24, _341_v2, _339_v0)).Dtor_cf9(), 1)
			(_nw52).ArraySet1(_dafny.IntOfInt64(895), 2)
			(_nw52).ArraySet1(_dafny.IntOfUint32((_366_v25).Cardinality()), 3)
			(_nw52).ArraySet1(_339_v0, 4)
			_367_v26 = _nw52
			var _368_v27 _dafny.Map
			_ = _368_v27
			_368_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v26, (_365_v24).F27())
			var _369_v29 _dafny.Set
			_ = _369_v29
			_369_v29 = _dafny.SetOf(true)
			var _370_v30 _dafny.Map
			_ = _370_v30
			_370_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_365_v24).F27(), (_365_v24).F27())
			var _371_v32 _dafny.MultiSet
			_ = _371_v32
			_371_v32 = _dafny.MultiSetOf(!(_341_v2), (_365_v24).F27())
			var _372_v33 _dafny.Array
			_ = _372_v33
			var _nwElement0_13 _dafny.Set = _dafny.SetOf(_339_v0)
			_ = _nwElement0_13
			var _nw53 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(25))
			_ = _nw53
			(_nw53).ArraySet1(_nwElement0_13, 0)
			(_nw53).ArraySet1(_363_v22, 1)
			(_nw53).ArraySet1(_363_v22, 2)
			(_nw53).ArraySet1((_364_v23).Dtor_cf8(), 3)
			(_nw53).ArraySet1(_dafny.SetOf((_368_v27).Cardinality(), _339_v0, _339_v0), 4)
			(_nw53).ArraySet1(func() _dafny.Set {
				var _coll19 = _dafny.NewBuilder()
				_ = _coll19
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(983), _dafny.IntOfInt64(683))); ; {
					_compr_19, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _373_v28 _dafny.Int
					_373_v28 = interface{}(_compr_19).(_dafny.Int)
					if ((_dafny.IntOfInt64(983)).Cmp(_373_v28) <= 0) && ((_373_v28).Cmp(_dafny.IntOfInt64(683)) < 0) {
						_coll19.Add((_373_v28).Times(_339_v0))
					}
				}
				return _coll19.ToSet()
			}(), 5)
			(_nw53).ArraySet1(_363_v22, 6)
			(_nw53).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(22))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}(func(_374_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))).Cardinality()), _339_v0), 7)
			(_nw53).ArraySet1(_363_v22, 8)
			(_nw53).ArraySet1(_363_v22, 9)
			(_nw53).ArraySet1(_363_v22, 10)
			(_nw53).ArraySet1(_363_v22, 11)
			(_nw53).ArraySet1(_363_v22, 12)
			(_nw53).ArraySet1(_363_v22, 13)
			(_nw53).ArraySet1(_363_v22, 14)
			(_nw53).ArraySet1(_dafny.SetOf(_dafny.IntOfInt64(267), (_dafny.MultiSetFromSeq(_340_v1)).Cardinality(), Companion_Default___.Fm5(globalState), (_369_v29).Cardinality(), _339_v0), 15)
			(_nw53).ArraySet1(func() _dafny.Set {
				var _coll20 = _dafny.NewBuilder()
				_ = _coll20
				for _iter21 := _dafny.Iterate((Companion_Default___.Fm6(_340_v1, (_370_v30).Cardinality(), _dafny.IntOfInt64(655), globalState)).Elements()); ; {
					_compr_20, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _375_v31 _dafny.Int
					_375_v31 = interface{}(_compr_20).(_dafny.Int)
					if (Companion_Default___.Fm6(_340_v1, (_370_v30).Cardinality(), _dafny.IntOfInt64(655), globalState)).Contains(_375_v31) {
						_coll20.Add(Companion_Default___.SafeDivisionInt(_375_v31, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
					}
				}
				return _coll20.ToSet()
			}(), 16)
			(_nw53).ArraySet1(_363_v22, 17)
			(_nw53).ArraySet1(_363_v22, 18)
			(_nw53).ArraySet1(_363_v22, 19)
			(_nw53).ArraySet1(_dafny.SetOf((_371_v32).Cardinality()), 20)
			(_nw53).ArraySet1(_363_v22, 21)
			(_nw53).ArraySet1(_363_v22, 22)
			(_nw53).ArraySet1(_363_v22, 23)
			(_nw53).ArraySet1(_dafny.SetOf(_339_v0), 24)
			_372_v33 = _nw53
			var _376_v34 _dafny.Array
			_ = _376_v34
			var _nw54 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
			_ = _nw54
			_376_v34 = _nw54
			var _377_v35 _dafny.Map
			_ = _377_v35
			_377_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_372_v33, _376_v34)
			var _378_v36 _dafny.Sequence
			_ = _378_v36
			_378_v36 = _dafny.SeqOf(_376_v34, _376_v34, _376_v34)
			_377_v35 = (_377_v35).Update(_372_v33, (_378_v36).Select((Companion_Default___.SafeIndex(_339_v0, _dafny.IntOfUint32((_378_v36).Cardinality()))).Uint32()).(_dafny.Array))
			if !((_365_v24).F27()) || (_341_v2) {
				var _379_v37 _dafny.CodePoint
				_ = _379_v37
				_379_v37 = _dafny.CodePoint('n')
				_379_v37 = _379_v37
				var _380_v38 _dafny.Sequence
				_ = _380_v38
				_380_v38 = _dafny.SeqOf(_341_v2)
				var _381_v39 _dafny.Array
				_ = _381_v39
				var _nwElement0_14 bool = (_365_v24).F27()
				_ = _nwElement0_14
				var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(13))
				_ = _nw55
				(_nw55).ArraySet1(_nwElement0_14, 0)
				(_nw55).ArraySet1((_365_v24).F27(), 1)
				(_nw55).ArraySet1(Companion_Default___.Fm0(_380_v38, globalState), 2)
				(_nw55).ArraySet1(Companion_Default___.Fm0(_380_v38, globalState), 3)
				(_nw55).ArraySet1(_341_v2, 4)
				(_nw55).ArraySet1((_365_v24).F27(), 5)
				(_nw55).ArraySet1(_341_v2, 6)
				(_nw55).ArraySet1(((_365_v24).F27()) || ((_365_v24).F27()), 7)
				(_nw55).ArraySet1(false, 8)
				(_nw55).ArraySet1(((_365_v24).F27()), 9)
				(_nw55).ArraySet1((_365_v24).F27(), 10)
				(_nw55).ArraySet1(!_dafny.Companion_Sequence_.Equal(_380_v38, _380_v38), 11)
				(_nw55).ArraySet1((_365_v24).F27(), 12)
				_381_v39 = _nw55
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_381_v39), 0))
				_ = _index37
				(_381_v39).ArraySet1(false, (_index37).Int())
				var _382_v40 _dafny.Map
				_ = _382_v40
				_382_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v0, false)
				var _383_v41 _dafny.Map
				_ = _383_v41
				_383_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_382_v40).Contains(_339_v0) {
						return (_382_v40).Get(_339_v0).(bool)
					}
					return (_365_v24).F27()
				})(), _366_v25)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_381_v39), 0))
				_ = _index38
				var _rhs36 bool = (Companion_D4_.Create_DC8_(_dafny.CodePoint('m'), (_365_v24).F27(), _341_v2)).Dtor_cf14()
				_ = _rhs36
				var _rhs37 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_366_v25, (func() _dafny.Sequence {
					if (_383_v41).Contains(!(_341_v2)) {
						return (_383_v41).Get(!(_341_v2)).(_dafny.Sequence)
					}
					return _366_v25
				})()))
				_ = _rhs37
				var _lhs28 _dafny.Array = _381_v39
				_ = _lhs28
				var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_381_v39), 0))
				_ = _lhs29
				var _lhs30 *GlobalState = globalState
				_ = _lhs30
				(_lhs28).ArraySet1(_rhs36, (_lhs29).Int())
				_lhs30.F19 = _rhs37
				var _384_v42 _dafny.Int
				_ = _384_v42
				var _385_v43 bool
				_ = _385_v43
				var _386_v44 bool
				_ = _386_v44
				var _387_v45 _dafny.Int
				_ = _387_v45
				var _out16 _dafny.Int
				_ = _out16
				var _out17 bool
				_ = _out17
				var _out18 bool
				_ = _out18
				var _out19 _dafny.Int
				_ = _out19
				_out16, _out17, _out18, _out19 = Companion_Default___.M0(globalState)
				_384_v42 = _out16
				_385_v43 = _out17
				_386_v44 = _out18
				_387_v45 = _out19
				var _388_v46 _dafny.Map
				_ = _388_v46
				_388_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_379_v37, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("jpftdff"), Companion_Default___.Fm7(globalState)))
				_388_v46 = (_388_v46).Update(Companion_Default___.Fm8(globalState), _341_v2)
				var _389_v47 *C0
				_ = _389_v47
				var _nw56 *C0 = New_C0_()
				_ = _nw56
				_nw56.Ctor__((_381_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_381_v39), 0))).Int()).(bool))
				_389_v47 = _nw56
			} else {
				var _390_v48 _dafny.CodePoint
				_ = _390_v48
				_390_v48 = _dafny.CodePoint('o')
				var _391_v49 _dafny.MultiSet
				_ = _391_v49
				_391_v49 = _dafny.MultiSetOf(_390_v48, _390_v48, _dafny.CodePoint('t'))
				_391_v49 = Companion_Default___.Fm9(!(_341_v2) || ((_365_v24).F27()), _339_v0, _340_v1, (Companion_Default___.Fm10(globalState)).Update(_dafny.UnicodeSeqOfUtf8Bytes("cb"), (_365_v24).F27()), globalState)
				_390_v48 = _390_v48
				var _392_v50 _dafny.Array
				_ = _392_v50
				var _nwElement0_15 bool = (_365_v24).F27()
				_ = _nwElement0_15
				var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(18))
				_ = _nw57
				(_nw57).ArraySet1(_nwElement0_15, 0)
				(_nw57).ArraySet1((_365_v24).F27(), 1)
				(_nw57).ArraySet1((_365_v24).F27(), 2)
				(_nw57).ArraySet1(_341_v2, 3)
				(_nw57).ArraySet1((_365_v24).F27(), 4)
				(_nw57).ArraySet1(Companion_Default___.Fm0(_dafny.SeqOf((_365_v24).F27(), !(_341_v2)), globalState), 5)
				(_nw57).ArraySet1((_365_v24).F27(), 6)
				(_nw57).ArraySet1((_365_v24).F27(), 7)
				(_nw57).ArraySet1((_365_v24).F27(), 8)
				(_nw57).ArraySet1((_365_v24).F27(), 9)
				(_nw57).ArraySet1((_342_v3).Dtor_cf5(), 10)
				(_nw57).ArraySet1(_341_v2, 11)
				(_nw57).ArraySet1((_365_v24).F27(), 12)
				(_nw57).ArraySet1((_365_v24).F27(), 13)
				(_nw57).ArraySet1(_341_v2, 14)
				(_nw57).ArraySet1((_365_v24).F27(), 15)
				(_nw57).ArraySet1((_365_v24).F27(), 16)
				(_nw57).ArraySet1((_365_v24).F27(), 17)
				_392_v50 = _nw57
				var _393_v51 _dafny.Array
				_ = _393_v51
				_393_v51 = _392_v50
				var _394_v52 _dafny.Array
				_ = _394_v52
				var _nwElement0_16 _dafny.Array = _392_v50
				_ = _nwElement0_16
				var _nw58 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(14))
				_ = _nw58
				(_nw58).ArraySet1(_nwElement0_16, 0)
				(_nw58).ArraySet1(_392_v50, 1)
				(_nw58).ArraySet1(_392_v50, 2)
				(_nw58).ArraySet1(_392_v50, 3)
				(_nw58).ArraySet1(_392_v50, 4)
				(_nw58).ArraySet1(_392_v50, 5)
				(_nw58).ArraySet1(_392_v50, 6)
				(_nw58).ArraySet1(_392_v50, 7)
				(_nw58).ArraySet1(_392_v50, 8)
				(_nw58).ArraySet1(_392_v50, 9)
				(_nw58).ArraySet1(_392_v50, 10)
				(_nw58).ArraySet1(_392_v50, 11)
				(_nw58).ArraySet1(_392_v50, 12)
				(_nw58).ArraySet1(_392_v50, 13)
				_394_v52 = _nw58
				var _395_v53 _dafny.Map
				_ = _395_v53
				_395_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_394_v52), _339_v0)
				var _396_v54 _dafny.Map
				_ = _396_v54
				_396_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v26, (func() _dafny.Int {
					if (_395_v53).Contains(_394_v52) {
						return (_395_v53).Get(_394_v52).(_dafny.Int)
					}
					return _339_v0
				})())
				var _397_v55 D4
				_ = _397_v55
				_397_v55 = Companion_D4_.Create_DC7_((func() _dafny.Int {
					if (_396_v54).Contains(_367_v26) {
						return (_396_v54).Get(_367_v26).(_dafny.Int)
					}
					return _dafny.IntOfInt64(251)
				})(), _365_v24, (Companion_Default___.Fm11(_339_v0, _dafny.IntOfUint32((Companion_Default___.Fm12(_341_v2, _dafny.IntOfUint32((_340_v1).Cardinality()), (_365_v24).F27(), globalState)).Cardinality()), (_365_v24).F27(), _339_v0, globalState)), _339_v0)
				var _rhs38 _dafny.Array = _393_v51
				_ = _rhs38
				var _rhs39 D4 = _397_v55
				_ = _rhs39
				var _rhs40 D3 = (func() D3 {
					if (_365_v24).F27() {
						return _342_v3
					}
					return _342_v3
				})()
				_ = _rhs40
				var _rhs41 _dafny.Array = _392_v50
				_ = _rhs41
				_393_v51 = _rhs38
				_397_v55 = _rhs39
				_342_v3 = _rhs40
				_392_v50 = _rhs41
				var _398_v56 D6
				_ = _398_v56
				_398_v56 = Companion_D6_.Create_DC10_(_367_v26)
				var _399_v57 _dafny.Array
				_ = _399_v57
				var _nwElement0_17 _dafny.Array = _367_v26
				_ = _nwElement0_17
				var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(27))
				_ = _nw59
				(_nw59).ArraySet1(_nwElement0_17, 0)
				(_nw59).ArraySet1((_398_v56).Dtor_cf17(), 1)
				(_nw59).ArraySet1(_367_v26, 2)
				(_nw59).ArraySet1(_367_v26, 3)
				(_nw59).ArraySet1(_367_v26, 4)
				(_nw59).ArraySet1(_367_v26, 5)
				(_nw59).ArraySet1(_367_v26, 6)
				(_nw59).ArraySet1(_367_v26, 7)
				(_nw59).ArraySet1(_367_v26, 8)
				(_nw59).ArraySet1(_367_v26, 9)
				(_nw59).ArraySet1(_367_v26, 10)
				(_nw59).ArraySet1(_367_v26, 11)
				(_nw59).ArraySet1(_367_v26, 12)
				(_nw59).ArraySet1(_367_v26, 13)
				(_nw59).ArraySet1(_367_v26, 14)
				(_nw59).ArraySet1(_367_v26, 15)
				(_nw59).ArraySet1(_367_v26, 16)
				(_nw59).ArraySet1(_367_v26, 17)
				(_nw59).ArraySet1((func() _dafny.Array {
					if (_365_v24).F27() {
						return _367_v26
					}
					return _367_v26
				})(), 18)
				(_nw59).ArraySet1(_367_v26, 19)
				(_nw59).ArraySet1(_367_v26, 20)
				(_nw59).ArraySet1(_367_v26, 21)
				(_nw59).ArraySet1(_367_v26, 22)
				(_nw59).ArraySet1(_367_v26, 23)
				(_nw59).ArraySet1(_367_v26, 24)
				(_nw59).ArraySet1(_367_v26, 25)
				(_nw59).ArraySet1(_367_v26, 26)
				_399_v57 = _nw59
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_399_v57), 0))
				_ = _index39
				(_399_v57).ArraySet1(_367_v26, (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_399_v57), 0))
				_ = _index40
				var _rhs42 _dafny.CodePoint = _390_v48
				_ = _rhs42
				var _rhs43 _dafny.Array = _367_v26
				_ = _rhs43
				var _rhs44 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_366_v25, (Companion_Default___.SafeIndex(_339_v0, _dafny.IntOfUint32((_366_v25).Cardinality()))).Uint32(), Companion_Default___.Fm8(globalState))
				_ = _rhs44
				var _rhs45 *C0 = _365_v24
				_ = _rhs45
				var _lhs31 _dafny.Array = _399_v57
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_399_v57), 0))
				_ = _lhs32
				_390_v48 = _rhs42
				(_lhs31).ArraySet1(_rhs43, (_lhs32).Int())
				_366_v25 = _rhs44
				_365_v24 = _rhs45
				var _400_v58 _dafny.Map
				_ = _400_v58
				_400_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_365_v24).F27(), _dafny.IntOfInt64(-320))
				var _401_v59 _dafny.Sequence
				_ = _401_v59
				_401_v59 = _dafny.SeqOf(_341_v2)
				var _402_v60 D7
				_ = _402_v60
				_402_v60 = Companion_D7_.Create_DC13_(_400_v58)
				_400_v58 = ((_400_v58).Update(!(_341_v2), _dafny.IntOfUint32((_401_v59).Cardinality()))).Merge((_402_v60).Dtor_cf20())
			}
			var _403_v61 _dafny.Array
			_ = _403_v61
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_7
			var _nw60 _dafny.Array
			_ = _nw60
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw60 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = (func(_404_v24 *C0, _405_v2 bool) func(_dafny.Int) bool {
					return func(_406_i2 _dafny.Int) bool {
						return !((_404_v24).F27()) || (_405_v2)
					}
				})(_365_v24, _341_v2)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw60 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw60).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw60).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_403_v61 = _nw60
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_403_v61), 0))
			_ = _index41
			(_403_v61).ArraySet1(true, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_403_v61), 0))
			_ = _index42
			(_403_v61).ArraySet1(Companion_Default___.Fm0(_dafny.SeqOf(_341_v2), globalState), (_index42).Int())
		}
		var _407_v62 _dafny.Array
		_ = _407_v62
		var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
		_ = _nw61
		_407_v62 = _nw61
		var _408_v63 _dafny.Array
		_ = _408_v63
		var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw62
		_408_v63 = _nw62
		var _409_v64 D6
		_ = _409_v64
		_409_v64 = Companion_D6_.Create_DC10_(_408_v63)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))
		_ = _index43
		(_407_v62).ArraySet1((_409_v64).Dtor_cf17(), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))
		_ = _index44
		(_407_v62).ArraySet1(_408_v63, (_index44).Int())
		var _410_v65 _dafny.Array
		_ = _410_v65
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_8
		var _nw63 _dafny.Array
		_ = _nw63
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw63 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = (func(_411_v2 bool) func(_dafny.Int) bool {
				return func(_412_i4 _dafny.Int) bool {
					return _411_v2
				}
			})(_341_v2)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw63 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw63).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw63).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_410_v65 = _nw63
		var _413_v66 _dafny.Map
		_ = _413_v66
		_413_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v0, _410_v65)
		var _hi2 _dafny.Int = (_339_v0).Times(_339_v0)
		_ = _hi2
		for _414_i3 := (func() _dafny.Int {
			if _341_v2 {
				return _339_v0
			}
			return (_dafny.Zero).Minus((_413_v66).Cardinality())
		})(); _414_i3.Cmp(_hi2) < 0; _414_i3 = _414_i3.Plus(_dafny.One) {
			(globalState).F5 = false
			(globalState).F8 = _339_v0
			var _415_v68 _dafny.Set
			_ = _415_v68
			_415_v68 = _dafny.SetOf(_414_i3)
			var _416_v69 *C0
			_ = _416_v69
			var _nw64 *C0 = New_C0_()
			_ = _nw64
			_nw64.Ctor__((_415_v68).IsProperSubsetOf(func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(692), _dafny.IntOfInt64(714))); ; {
					_compr_21, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _417_v67 _dafny.Int
					_417_v67 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(692)).Cmp(_417_v67) <= 0) && ((_417_v67).Cmp(_dafny.IntOfInt64(714)) < 0) {
						_coll21.Add((_417_v67).Plus(_339_v0))
					}
				}
				return _coll21.ToSet()
			}()))
			_416_v69 = _nw64
			(globalState).F8 = Companion_Default___.Fm5(globalState)
		}
		if _341_v2 {
			var _418_v70 _dafny.CodePoint
			_ = _418_v70
			_418_v70 = _dafny.CodePoint('p')
			var _419_v71 _dafny.Map
			_ = _419_v71
			_419_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_418_v70, _339_v0)
			_419_v71 = (_419_v71).Update(_418_v70, _339_v0)
			var _arr4 _dafny.Array = _dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))
			_ = _arr4
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))
			_ = _index45
			(_arr4).ArraySet1(_339_v0, (_index45).Int())
			var _arr5 _dafny.Array = _dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))
			_ = _arr5
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))
			_ = _index46
			(_arr5).ArraySet1(_339_v0, (_index46).Int())
			(globalState).F5 = _341_v2
			var _arr6 _dafny.Array = _dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))
			_ = _arr6
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))
			_ = _index47
			(_arr6).ArraySet1((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))).Int()).(_dafny.Int), (_index47).Int())
			var _420_v72 _dafny.Sequence
			_ = _420_v72
			_420_v72 = _dafny.UnicodeSeqOfUtf8Bytes("xm")
			var _421_v73 _dafny.Map
			_ = _421_v73
			_421_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm13(globalState)).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_420_v72, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nbgsu"), (Companion_Default___.SafeIndex((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nbgsu")).Cardinality()))).Uint32(), _dafny.CodePoint('t'))))
			_421_v73 = (_421_v73).Update((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_dafny.ArrayCastTo((_407_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_407_v62), 0))).Int()))), 0))).Int()).(_dafny.Int), _dafny.Companion_Sequence_.Concatenate(_420_v72, _420_v72))
		} else {
			(globalState).F22 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("buaexsrxf"), _dafny.UnicodeSeqOfUtf8Bytes("pwblopusp"))
			var _422_v74 _dafny.Map
			_ = _422_v74
			_422_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_339_v0, _341_v2)
			var _423_v75 _dafny.Sequence
			_ = _423_v75
			_423_v75 = _dafny.UnicodeSeqOfUtf8Bytes("pprdiuoua")
			var _424_v76 _dafny.Sequence
			_ = _424_v76
			_424_v76 = _dafny.SeqOf(_422_v74, Companion_Default___.Fm14(_341_v2, _423_v75, globalState), _422_v74)
			var _425_v78 _dafny.Set
			_ = _425_v78
			_425_v78 = _dafny.SetOf((_dafny.Zero).Minus(_339_v0))
			var _426_v79 _dafny.Map
			_ = _426_v79
			_426_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter23 := _dafny.Iterate((_425_v78).Elements()); ; {
					_compr_22, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _427_v77 _dafny.Int
					_427_v77 = interface{}(_compr_22).(_dafny.Int)
					if (_425_v78).Contains(_427_v77) {
						_coll22.Add(Companion_Default___.SafeDivisionInt(_427_v77, _339_v0), (Companion_D4_.Create_DC8_(_dafny.CodePoint('x'), _341_v2, _341_v2)).Dtor_cf14())
					}
				}
				return _coll22.ToMap()
			}()))
			var _428_v80 _dafny.MultiSet
			_ = _428_v80
			_428_v80 = _dafny.MultiSetOf(true, _341_v2)
			var _rhs46 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_424_v76, _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_426_v79).Contains((func() bool {
					if (_422_v74).Contains(_dafny.IntOfUint32((_423_v75).Cardinality())) {
						return (_422_v74).Get(_dafny.IntOfUint32((_423_v75).Cardinality())).(bool)
					}
					return _341_v2
				})()) {
					return (_426_v79).Get((func() bool {
						if (_422_v74).Contains(_dafny.IntOfUint32((_423_v75).Cardinality())) {
							return (_422_v74).Get(_dafny.IntOfUint32((_423_v75).Cardinality())).(bool)
						}
						return _341_v2
					})()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_424_v76, (Companion_Default___.SafeIndex((_428_v80).Cardinality(), _dafny.IntOfUint32((_424_v76).Cardinality()))).Uint32(), _422_v74)
			})(), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_423_v75).Cardinality()), _341_v2))))
			_ = _rhs46
			var _rhs47 bool = _341_v2
			_ = _rhs47
			var _rhs48 bool = (_339_v0).Cmp(_dafny.IntOfUint32((_423_v75).Cardinality())) > 0
			_ = _rhs48
			var _rhs49 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_429_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_430_i5 _dafny.Int) _dafny.Int {
					return _429_v0
				}
			})(_339_v0)))).Cardinality())
			_ = _rhs49
			var _lhs33 *GlobalState = globalState
			_ = _lhs33
			var _lhs34 *GlobalState = globalState
			_ = _lhs34
			_424_v76 = _rhs46
			_lhs33.F22 = _rhs47
			_lhs34.F19 = _rhs48
			_339_v0 = _rhs49
			var _431_v81 _dafny.Sequence
			_ = _431_v81
			_431_v81 = _dafny.SeqOf(_341_v2, _341_v2, _341_v2)
			var _432_v82 _dafny.MultiSet
			_ = _432_v82
			_432_v82 = _dafny.MultiSetOf(_339_v0)
			var _433_v83 D8
			_ = _433_v83
			_433_v83 = Companion_D8_.Create_DC16_(_423_v75)
			var _434_v84 _dafny.MultiSet
			_ = _434_v84
			_434_v84 = _dafny.MultiSetOf(_dafny.IntOfInt64(165), (_339_v0).Times(_339_v0), _dafny.IntOfUint32((_431_v81).Cardinality()), (func() _dafny.Int {
				if (_432_v82).Contains(_dafny.IntOfInt64(375)) {
					return (_432_v82).Multiplicity(_dafny.IntOfInt64(375))
				}
				return _dafny.IntOfUint32(((_433_v83).Dtor_cf29()).Cardinality())
			})())
			_432_v82 = ((_434_v84).Union(_432_v82)).Union(_432_v82)
			(globalState).F17 = _341_v2
			_422_v74 = _422_v74
		}
		var _435_v86 _dafny.Map
		_ = _435_v86
		_435_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_341_v2, true)).Cardinality(), Companion_Default___.Fm15(_dafny.IntOfInt64(-198), globalState))
		var _436_v87 _dafny.CodePoint
		_ = _436_v87
		_436_v87 = _dafny.CodePoint('j')
		var _437_v88 _dafny.Map
		_ = _437_v88
		_437_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_436_v87, _339_v0)
		(globalState).F22 = (_339_v0).Cmp((_dafny.Zero).Minus((_339_v0).Plus((func() _dafny.Map {
			var _coll23 = _dafny.NewMapBuilder()
			_ = _coll23
			for _iter24 := _dafny.Iterate(((func() _dafny.Map {
				if (_435_v86).Contains(_339_v0) {
					return (_435_v86).Get(_339_v0).(_dafny.Map)
				}
				return _437_v88
			})()).Keys().Elements()); ; {
				_compr_23, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _438_v85 _dafny.CodePoint
				_438_v85 = interface{}(_compr_23).(_dafny.CodePoint)
				if ((func() _dafny.Map {
					if (_435_v86).Contains(_339_v0) {
						return (_435_v86).Get(_339_v0).(_dafny.Map)
					}
					return _437_v88
				})()).Contains(_438_v85) {
					_coll23.Add(_438_v85, _339_v0)
				}
			}
			return _coll23.ToMap()
		}()).Cardinality()))) != 0
		(globalState).F5 = _341_v2
	}
}
func (_this *C2) M2(globalState *GlobalState) {
	{
		var _439_v0 _dafny.Sequence
		_ = _439_v0
		_439_v0 = _dafny.UnicodeSeqOfUtf8Bytes("nmnolf")
		var _440_v1 D8
		_ = _440_v1
		_440_v1 = Companion_D8_.Create_DC16_(_439_v0)
		var _441_v2 _dafny.Int
		_ = _441_v2
		_441_v2 = _dafny.IntOfInt64(346)
		var _pat_let_tv7 = _439_v0
		_ = _pat_let_tv7
		var _pat_let_tv8 = _441_v2
		_ = _pat_let_tv8
		var _pat_let_tv9 = _439_v0
		_ = _pat_let_tv9
		_440_v1 = func(_pat_let2_0 D8) D8 {
			return func(_442_dt__update__tmp_h0 D8) D8 {
				return func(_pat_let3_0 _dafny.Sequence) D8 {
					return func(_443_dt__update_hcf29_h0 _dafny.Sequence) D8 {
						return Companion_D8_.Create_DC16_(_443_dt__update_hcf29_h0)
					}(_pat_let3_0)
				}(_dafny.Companion_Sequence_.Update(_pat_let_tv7, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_pat_let_tv8), _dafny.IntOfUint32((_pat_let_tv9).Cardinality()))).Uint32(), _dafny.CodePoint('k')))
			}(_pat_let2_0)
		}(_440_v1)
		var _444_v3 bool
		_ = _444_v3
		_444_v3 = false
		var _445_v4 _dafny.MultiSet
		_ = _445_v4
		_445_v4 = _dafny.MultiSetOf(_444_v3)
		var _446_v5 _dafny.Array
		_ = _446_v5
		var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
		_ = _nw65
		_446_v5 = _nw65
		var _447_v6 _dafny.CodePoint
		_ = _447_v6
		_447_v6 = _dafny.CodePoint('y')
		var _448_v7 _dafny.Sequence
		_ = _448_v7
		_448_v7 = _dafny.SeqOf(false)
		var _449_v8 _dafny.MultiSet
		_ = _449_v8
		_449_v8 = _dafny.MultiSetOf(_dafny.IntOfUint32((_448_v7).Cardinality()))
		var _450_v9 _dafny.Sequence
		_ = _450_v9
		_450_v9 = _dafny.SeqOf(_441_v2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v2, _444_v3)).Cardinality())
		var _451_v10 _dafny.Array
		_ = _451_v10
		var _nwElement0_18 _dafny.Int = _441_v2
		_ = _nwElement0_18
		var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
		_ = _nw66
		(_nw66).ArraySet1(_nwElement0_18, 0)
		(_nw66).ArraySet1((_445_v4).Cardinality(), 1)
		(_nw66).ArraySet1(_441_v2, 2)
		(_nw66).ArraySet1(_441_v2, 3)
		(_nw66).ArraySet1(_441_v2, 4)
		(_nw66).ArraySet1(_441_v2, 5)
		(_nw66).ArraySet1(_441_v2, 6)
		(_nw66).ArraySet1(_441_v2, 7)
		(_nw66).ArraySet1(_441_v2, 8)
		(_nw66).ArraySet1(_441_v2, 9)
		(_nw66).ArraySet1(_441_v2, 10)
		(_nw66).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_446_v5)).Cardinality()), 11)
		(_nw66).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dauxjhvs")).Cardinality())), 12)
		(_nw66).ArraySet1(_dafny.IntOfInt64(103), 13)
		(_nw66).ArraySet1((_dafny.SetOf(_447_v6)).Cardinality(), 14)
		(_nw66).ArraySet1(_441_v2, 15)
		(_nw66).ArraySet1(_dafny.IntOfInt64(294), 16)
		(_nw66).ArraySet1(_dafny.IntOfInt64(958), 17)
		(_nw66).ArraySet1(_dafny.IntOfInt64(990), 18)
		(_nw66).ArraySet1((_449_v8).Cardinality(), 19)
		(_nw66).ArraySet1((_dafny.Zero).Minus(_441_v2), 20)
		(_nw66).ArraySet1(_441_v2, 21)
		(_nw66).ArraySet1(Companion_Default___.Fm5(globalState), 22)
		(_nw66).ArraySet1(_dafny.IntOfUint32((_450_v9).Cardinality()), 23)
		(_nw66).ArraySet1(_441_v2, 24)
		(_nw66).ArraySet1(_dafny.IntOfUint32((_439_v0).Cardinality()), 25)
		(_nw66).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_441_v2, (_450_v9).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_450_v9).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_439_v0).Cardinality()), _441_v2)).Cardinality()), 26)
		(_nw66).ArraySet1(_441_v2, 27)
		(_nw66).ArraySet1(_441_v2, 28)
		_451_v10 = _nw66
		var _452_v11 _dafny.Sequence
		_ = _452_v11
		_452_v11 = _dafny.SeqOf(_446_v5)
		var _453_v12 _dafny.Int
		_ = _453_v12
		_453_v12 = _441_v2
		var _454_v13 _dafny.Map
		_ = _454_v13
		_454_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_453_v12, _447_v6)
		var _455_i0 _dafny.Int
		_ = _455_i0
		_455_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_451_v10, (_452_v11).Select((Companion_Default___.SafeIndex((_454_v13).Cardinality(), _dafny.IntOfUint32((_452_v11).Cardinality()))).Uint32()).(_dafny.Array), _446_v5, _451_v10), _452_v11), _dafny.Companion_Sequence_.Concatenate(_452_v11, _452_v11)) {
				{
					if (_455_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_455_i0 = (_455_i0).Plus(_dafny.One)
					if Companion_Default___.Fm0(_448_v7, globalState) {
						var _456_v14 _dafny.Array
						_ = _456_v14
						var _nw67 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
						_ = _nw67
						_456_v14 = _nw67
						var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_456_v14), 0))
						_ = _index48
						(_456_v14).ArraySet1(_444_v3, (_index48).Int())
						var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(496), _dafny.ArrayLen((_456_v14), 0))
						_ = _index49
						(_456_v14).ArraySet1(!(false), (_index49).Int())
						(globalState).F8 = Companion_Default___.Fm5(globalState)
						var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))
						_ = _index50
						(_451_v10).ArraySet1(Companion_Default___.SafeDivisionInt(_441_v2, _441_v2), (_index50).Int())
						var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))
						_ = _index51
						var _rhs50 _dafny.Sequence = _450_v9
						_ = _rhs50
						var _rhs51 _dafny.Int = _441_v2
						_ = _rhs51
						var _lhs35 _dafny.Array = _451_v10
						_ = _lhs35
						var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))
						_ = _lhs36
						_450_v9 = _rhs50
						(_lhs35).ArraySet1(_rhs51, (_lhs36).Int())
						_445_v4 = (_445_v4).Difference(_445_v4)
						var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))
						_ = _index52
						(_451_v10).ArraySet1(Companion_Default___.SafeModuloInt((_451_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))).Int()).(_dafny.Int), ((_451_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(266), _dafny.ArrayLen((_451_v10), 0))).Int()).(_dafny.Int)).Plus(_dafny.IntOfUint32((_439_v0).Cardinality()))), (_index52).Int())
					} else {
						var _457_v15 _dafny.Map
						_ = _457_v15
						_457_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v3, _444_v3)
						(globalState).F17 = ((_441_v2).Minus((_dafny.Zero).Minus((_457_v15).Cardinality()))).Cmp(_441_v2) == 0
						var _458_v16 *C0
						_ = _458_v16
						var _nw68 *C0 = New_C0_()
						_ = _nw68
						_nw68.Ctor__(_444_v3)
						_458_v16 = _nw68
						var _459_v17 _dafny.Int
						_ = _459_v17
						var _460_v18 bool
						_ = _460_v18
						var _461_v19 bool
						_ = _461_v19
						var _462_v20 _dafny.Int
						_ = _462_v20
						var _out20 _dafny.Int
						_ = _out20
						var _out21 bool
						_ = _out21
						var _out22 bool
						_ = _out22
						var _out23 _dafny.Int
						_ = _out23
						_out20, _out21, _out22, _out23 = Companion_Default___.M0(globalState)
						_459_v17 = _out20
						_460_v18 = _out21
						_461_v19 = _out22
						_462_v20 = _out23
						_447_v6 = _447_v6
						var _463_v21 _dafny.Array
						_ = _463_v21
						var _len0_9 _dafny.Int = _dafny.IntOfInt64(27)
						_ = _len0_9
						var _nw69 _dafny.Array
						_ = _nw69
						if _len0_9.Cmp(_dafny.Zero) == 0 {
							_nw69 = _dafny.NewArray(_len0_9)
						} else {
							var _init9 func(_dafny.Int) _dafny.MultiSet = (func(_464_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.MultiSet {
								return func(_465_i1 _dafny.Int) _dafny.MultiSet {
									return ((Companion_D9_.Create_DC19_(_dafny.MultiSetOf(_464_v6))).Dtor_cf34()).Union(_dafny.MultiSetOf(_dafny.CodePoint('a'), _464_v6))
								}
							})(_447_v6)
							_ = _init9
							var _element0_9 = _init9(_dafny.Zero)
							_ = _element0_9
							_nw69 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
							(_nw69).ArraySet1(_element0_9, 0)
							var _nativeLen0_9 = (_len0_9).Int()
							_ = _nativeLen0_9
							for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
								(_nw69).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
							}
						}
						_463_v21 = _nw69
						var _466_v22 _dafny.MultiSet
						_ = _466_v22
						_466_v22 = _dafny.MultiSetOf(_447_v6)
						var _467_v23 _dafny.Sequence
						_ = _467_v23
						_467_v23 = _dafny.SeqOf(_439_v0)
						var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_463_v21), 0))
						_ = _index53
						(_463_v21).ArraySet1(((_466_v22).Update(_447_v6, Companion_Default___.Abs(_dafny.IntOfUint32((_467_v23).Cardinality())))).Update(_447_v6, Companion_Default___.Abs(_459_v17)), (_index53).Int())
						var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_463_v21), 0))
						_ = _index54
						(_463_v21).ArraySet1(_466_v22, (_index54).Int())
					}
					var _468_v24 _dafny.Set
					_ = _468_v24
					_468_v24 = _dafny.SetOf(_441_v2, _441_v2, _441_v2)
					var _469_v26 _dafny.MultiSet
					_ = _469_v26
					_469_v26 = _dafny.MultiSetOf(_468_v24, _468_v24, func() _dafny.Set {
						var _coll24 = _dafny.NewBuilder()
						_ = _coll24
						for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(197), _dafny.IntOfInt64(452))); ; {
							_compr_24, _ok25 := _iter25()
							if !_ok25 {
								break
							}
							var _470_v25 _dafny.Int
							_470_v25 = interface{}(_compr_24).(_dafny.Int)
							if ((_dafny.IntOfInt64(197)).Cmp(_470_v25) <= 0) && ((_470_v25).Cmp(_dafny.IntOfInt64(452)) < 0) {
								_coll24.Add((_470_v25).Times(_441_v2))
							}
						}
						return _coll24.ToSet()
					}(), (_468_v24).Difference(_468_v24))
					var _471_v27 _dafny.Array
					_ = _471_v27
					var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
					_ = _nw70
					_471_v27 = _nw70
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_471_v27), 0))
					_ = _index55
					(_471_v27).ArraySet1(!(_444_v3), (_index55).Int())
					var _472_v28 *C0
					_ = _472_v28
					var _nw71 *C0 = New_C0_()
					_ = _nw71
					_nw71.Ctor__(!(false))
					_472_v28 = _nw71
					var _473_v29 D4
					_ = _473_v29
					_473_v29 = Companion_D4_.Create_DC7_(_441_v2, _472_v28, true, _441_v2)
					var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_471_v27), 0))
					_ = _index56
					var _rhs52 _dafny.MultiSet = (_469_v26).Intersection(_469_v26)
					_ = _rhs52
					var _rhs53 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(744))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg27 _dafny.Int) interface{} {
							return coer27(arg27)
						}
					}((func(_474_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_475_i2 _dafny.Int) _dafny.Int {
							return (_474_v2).Minus(_474_v2)
						}
					})(_441_v2)))).Cardinality())
					_ = _rhs53
					var _rhs54 _dafny.CodePoint = _447_v6
					_ = _rhs54
					var _rhs55 bool = (_473_v29).Dtor_cf11()
					_ = _rhs55
					var _lhs37 _dafny.Array = _471_v27
					_ = _lhs37
					var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_471_v27), 0))
					_ = _lhs38
					_469_v26 = _rhs52
					_441_v2 = _rhs53
					_447_v6 = _rhs54
					(_lhs37).ArraySet1(_rhs55, (_lhs38).Int())
					_441_v2 = _441_v2
					(globalState).F22 = !(_dafny.Companion_Sequence_.Equal(_450_v9, _dafny.Companion_Sequence_.Concatenate(_450_v9, _450_v9)))
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _source9 D8 = _440_v1
		_ = _source9
		if _source9.Is_DC17() {
			var _476___mcc_h0 _dafny.CodePoint = _source9.Get_().(D8_DC17).Cf30
			_ = _476___mcc_h0
			var _477___mcc_h1 _dafny.CodePoint = _source9.Get_().(D8_DC17).Cf31
			_ = _477___mcc_h1
			var _478___mcc_h2 _dafny.Int = _source9.Get_().(D8_DC17).Cf32
			_ = _478___mcc_h2
			var _479_cf32 _dafny.Int = _478___mcc_h2
			_ = _479_cf32
			var _480_cf31 _dafny.CodePoint = _477___mcc_h1
			_ = _480_cf31
			var _481_cf30 _dafny.CodePoint = _476___mcc_h0
			_ = _481_cf30
			var _482_v30 _dafny.Sequence
			_ = _482_v30
			_482_v30 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wmd"))
			var _483_v31 _dafny.Array
			_ = _483_v31
			var _nwElement0_19 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_482_v30, _482_v30)
			_ = _nwElement0_19
			var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(16))
			_ = _nw72
			(_nw72).ArraySet1(_nwElement0_19, 0)
			(_nw72).ArraySet1(_482_v30, 1)
			(_nw72).ArraySet1(_482_v30, 2)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_482_v30, Companion_Default___.Fm16(_480_cf31, _441_v2, _dafny.IntOfInt64(526), _479_cf32, globalState)), 3)
			(_nw72).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}((func(_484_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_485_i3 _dafny.Int) _dafny.Sequence {
					return _484_v0
				}
			})(_439_v0))), 4)
			(_nw72).ArraySet1(_482_v30, 5)
			(_nw72).ArraySet1(_482_v30, 6)
			(_nw72).ArraySet1(_482_v30, 7)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("sf"), _439_v0), _482_v30), 8)
			(_nw72).ArraySet1(_482_v30, 9)
			(_nw72).ArraySet1(_482_v30, 10)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_482_v30, _482_v30), 11)
			(_nw72).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_439_v0), _482_v30), 12)
			(_nw72).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(102))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_486_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_487_i4 _dafny.Int) _dafny.Sequence {
					return _486_v0
				}
			})(_439_v0))), 13)
			(_nw72).ArraySet1(_482_v30, 14)
			(_nw72).ArraySet1(_482_v30, 15)
			_483_v31 = _nw72
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_483_v31), 0))
			_ = _index57
			(_483_v31).ArraySet1(_482_v30, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_483_v31), 0))
			_ = _index58
			(_483_v31).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _444_v3 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(231))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}((func(_488_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_489_i5 _dafny.Int) _dafny.Sequence {
							return _488_v0
						}
					})(_439_v0)))
				}
				return _482_v30
			})(), _dafny.Companion_Sequence_.Concatenate(_482_v30, _dafny.SeqOf(_439_v0, _439_v0, _439_v0, _439_v0))), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_479_cf32), _479_cf32), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if _444_v3 {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(231))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg31 _dafny.Int) interface{} {
							return coer31(arg31)
						}
					}((func(_490_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_491_i5 _dafny.Int) _dafny.Sequence {
							return _490_v0
						}
					})(_439_v0)))
				}
				return _482_v30
			})(), _dafny.Companion_Sequence_.Concatenate(_482_v30, _dafny.SeqOf(_439_v0, _439_v0, _439_v0, _439_v0)))).Cardinality()))).Uint32(), _dafny.Companion_Sequence_.Concatenate(_439_v0, _439_v0)), (_index58).Int())
			var _492_v33 _dafny.Set
			_ = _492_v33
			_492_v33 = _dafny.SetOf(_479_cf32)
			var _493_v34 _dafny.Set
			_ = _493_v34
			_493_v34 = _dafny.SetOf(_441_v2, (_492_v33).Cardinality())
			var _494_v35 _dafny.Map
			_ = _494_v35
			_494_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _479_cf32)
			var _495_v36 _dafny.Map
			_ = _495_v36
			_495_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_494_v35, _444_v3)
			var _496_v37 _dafny.Set
			_ = _496_v37
			_496_v37 = _dafny.SetOf((_492_v33).Cardinality(), (_495_v36).Cardinality(), _479_cf32)
			var _497_v38 _dafny.Map
			_ = _497_v38
			_497_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v3, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_451_v10), (Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_dafny.SeqOf(_451_v10)).Cardinality()))).Uint32(), _446_v5))
			var _498_v39 _dafny.Map
			_ = _498_v39
			_498_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(731), _dafny.IntOfInt64(-903))); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _499_v32 _dafny.Int
					_499_v32 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(731)).Cmp(_499_v32) <= 0) && ((_499_v32).Cmp(_dafny.IntOfInt64(-903)) < 0) {
						_coll25.Add((_499_v32).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mxwfxlfth")).Cardinality())))
					}
				}
				return _coll25.ToSet()
			}()).Equals(_492_v33), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_497_v38).Contains(_444_v3) {
					return (_497_v38).Get(_444_v3).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_446_v5, _446_v5)
			})()).Cardinality()))
			_494_v35 = (_494_v35).Update(_444_v3, _479_cf32)
			var _500_v40 _dafny.Array
			_ = _500_v40
			var _nw73 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw73
			_500_v40 = _nw73
			var _501_v41 _dafny.Sequence
			_ = _501_v41
			_501_v41 = _dafny.SeqOf(_500_v40, _500_v40)
			var _502_v42 _dafny.Map
			_ = _502_v42
			_502_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_441_v2).Plus(_dafny.IntOfUint32((_439_v0).Cardinality())), _501_v41)
			_502_v42 = (_502_v42).Update(_441_v2, _dafny.Companion_Sequence_.Concatenate(_501_v41, _dafny.SeqOf(_500_v40, _500_v40, _500_v40, _500_v40)))
			_441_v2 = _dafny.IntOfInt64(289)
		} else if _source9.Is_DC16() {
			var _503___mcc_h3 _dafny.Sequence = _source9.Get_().(D8_DC16).Cf29
			_ = _503___mcc_h3
			var _504_cf29 _dafny.Sequence = _503___mcc_h3
			_ = _504_cf29
			var _505_v43 _dafny.Array
			_ = _505_v43
			var _nwElement0_20 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(428))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}((func(_506_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_507_i6 _dafny.Int) _dafny.CodePoint {
					return _506_v6
				}
			})(_447_v6)))
			_ = _nwElement0_20
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(2))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_20, 0)
			(_nw74).ArraySet1((func() _dafny.Sequence {
				if _444_v3 {
					return _504_cf29
				}
				return _439_v0
			})(), 1)
			_505_v43 = _nw74
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_505_v43), 0))
			_ = _index59
			(_505_v43).ArraySet1(_439_v0, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_505_v43), 0))
			_ = _index60
			var _rhs56 _dafny.Int = (_441_v2).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(936))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}((func(_508_cf29 _dafny.Sequence, _509_v2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_510_i7 _dafny.Int) _dafny.CodePoint {
					return (_508_cf29).Select((Companion_Default___.SafeIndex(_509_v2, _dafny.IntOfUint32((_508_cf29).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_504_cf29, _441_v2)))).Cardinality()))
			_ = _rhs56
			var _rhs57 _dafny.Int = Companion_Default___.SafeModuloInt(_441_v2, Companion_Default___.Fm5(globalState))
			_ = _rhs57
			var _rhs58 bool = !(_444_v3)
			_ = _rhs58
			var _rhs59 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_504_cf29, _504_cf29), _439_v0)
			_ = _rhs59
			var _lhs39 *GlobalState = globalState
			_ = _lhs39
			var _lhs40 _dafny.Array = _505_v43
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(950), _dafny.ArrayLen((_505_v43), 0))
			_ = _lhs41
			_441_v2 = _rhs56
			_441_v2 = _rhs57
			_lhs39.F17 = _rhs58
			(_lhs40).ArraySet1(_rhs59, (_lhs41).Int())
			var _511_v44 _dafny.Map
			_ = _511_v44
			_511_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_445_v4, _444_v3)
			_511_v44 = (_511_v44).Merge(_511_v44)
			_447_v6 = _447_v6
			var _512_v45 _dafny.Map
			_ = _512_v45
			_512_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v3, _dafny.IntOfInt64(479))
			_512_v45 = ((Companion_Default___.Fm17(_444_v3, _441_v2, _444_v3, _444_v3, globalState)).Merge(_512_v45)).Merge(_512_v45)
		} else {
			var _513___mcc_h4 D8 = _source9.Get_().(D8_DC18).Cf33
			_ = _513___mcc_h4
			var _514_cf33 D8 = _513___mcc_h4
			_ = _514_cf33
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_446_v5), 0))
			_ = _index61
			(_446_v5).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_448_v7).Cardinality())), (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_446_v5), 0))
			_ = _index62
			(_446_v5).ArraySet1(_441_v2, (_index62).Int())
			var _515_v46 _dafny.Array
			_ = _515_v46
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw75
			_515_v46 = _nw75
			_515_v46 = _515_v46
			var _516_v47 _dafny.Map
			_ = _516_v47
			_516_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v0, _444_v3)
			_516_v47 = ((_516_v47).Merge(_516_v47)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v0, _444_v3))
			(globalState).F22 = Companion_Default___.Fm0(_dafny.Companion_Sequence_.Concatenate(_448_v7, _dafny.SeqOf(_444_v3, !(_444_v3), !(_444_v3))), globalState)
		}
		var _hi3 _dafny.Int = _dafny.IntOfUint32((_448_v7).Cardinality())
		_ = _hi3
		for _517_i8 := (_450_v9).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_450_v9).Cardinality()))).Uint32()).(_dafny.Int); _517_i8.Cmp(_hi3) < 0; _517_i8 = _517_i8.Plus(_dafny.One) {
			var _518_v48 _dafny.Array
			_ = _518_v48
			var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
			_ = _nw76
			_518_v48 = _nw76
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_518_v48), 0))
			_ = _index63
			(_518_v48).ArraySet1(_448_v7, (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_518_v48), 0))
			_ = _index64
			(_518_v48).ArraySet1(_448_v7, (_index64).Int())
			var _519_v49 _dafny.Set
			_ = _519_v49
			_519_v49 = _dafny.SetOf(_dafny.IntOfUint32((_439_v0).Cardinality()))
			var _520_v50 _dafny.Map
			_ = _520_v50
			_520_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v49, _439_v0)
			var _521_v51 T1
			_ = _521_v51
			var _nw77 *C0 = New_C0_()
			_ = _nw77
			_nw77.Ctor__(_444_v3)
			_521_v51 = _nw77
			var _522_v52 T0
			_ = _522_v52
			var _nw78 *C1 = New_C1_()
			_ = _nw78
			_nw78.Ctor__(_520_v50, _521_v51)
			_522_v52 = _nw78
			var _523_v53 _dafny.Map
			_ = _523_v53
			_523_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_522_v52, _441_v2)
			var _rhs60 _dafny.Map = _523_v53
			_ = _rhs60
			var _rhs61 _dafny.Int = (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(329), _441_v2)).Plus(_517_i8)
			_ = _rhs61
			var _rhs62 _dafny.Int = _517_i8
			_ = _rhs62
			var _lhs42 *GlobalState = globalState
			_ = _lhs42
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			_523_v53 = _rhs60
			_lhs42.F8 = _rhs61
			_lhs43.F8 = _rhs62
			var _524_v54 D7
			_ = _524_v54
			_524_v54 = Companion_D7_.Create_DC15_((_445_v4).Cardinality(), _444_v3, _441_v2)
			var _pat_let_tv10 = _439_v0
			_ = _pat_let_tv10
			var _source10 D7 = func(_pat_let4_0 D7) D7 {
				return func(_525_dt__update__tmp_h1 D7) D7 {
					return func(_pat_let5_0 _dafny.Int) D7 {
						return func(_526_dt__update_hcf28_h0 _dafny.Int) D7 {
							return Companion_D7_.Create_DC15_((_525_dt__update__tmp_h1).Dtor_cf26(), (_525_dt__update__tmp_h1).Dtor_cf27(), _526_dt__update_hcf28_h0)
						}(_pat_let5_0)
					}(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_pat_let_tv10).Cardinality()), _dafny.IntOfInt64(422)))
				}(_pat_let4_0)
			}(_524_v54)
			_ = _source10
			if _source10.Is_DC14() {
				var _527___mcc_h5 _dafny.Int = _source10.Get_().(D7_DC14).Cf21
				_ = _527___mcc_h5
				var _528___mcc_h6 bool = _source10.Get_().(D7_DC14).Cf22
				_ = _528___mcc_h6
				var _529___mcc_h7 bool = _source10.Get_().(D7_DC14).Cf23
				_ = _529___mcc_h7
				var _530___mcc_h8 bool = _source10.Get_().(D7_DC14).Cf24
				_ = _530___mcc_h8
				var _531___mcc_h9 _dafny.Int = _source10.Get_().(D7_DC14).Cf25
				_ = _531___mcc_h9
				var _532_cf25 _dafny.Int = _531___mcc_h9
				_ = _532_cf25
				var _533_cf24 bool = _530___mcc_h8
				_ = _533_cf24
				var _534_cf23 bool = _529___mcc_h7
				_ = _534_cf23
				var _535_cf22 bool = _528___mcc_h6
				_ = _535_cf22
				var _536_cf21 _dafny.Int = _527___mcc_h5
				_ = _536_cf21
				_519_v49 = _519_v49
				var _537_v55 _dafny.Array
				_ = _537_v55
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_10
				var _nw79 _dafny.Array
				_ = _nw79
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw79 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Map = (func(_538_v6 _dafny.CodePoint, _539_cf22 bool) func(_dafny.Int) _dafny.Map {
						return func(_540_i9 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _538_v6)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_539_cf22, _538_v6))
						}
					})(_447_v6, _535_cf22)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw79 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw79).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw79).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_537_v55 = _nw79
				var _rhs63 _dafny.Array = _537_v55
				_ = _rhs63
				var _rhs64 _dafny.Array = _446_v5
				_ = _rhs64
				var _lhs44 *GlobalState = globalState
				_ = _lhs44
				_537_v55 = _rhs63
				_lhs44.F23 = _rhs64
				(globalState).F5 = _534_cf23
				var _541_v56 _dafny.Array
				_ = _541_v56
				var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
				_ = _nw80
				_541_v56 = _nw80
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_541_v56), 0))
				_ = _index65
				(_541_v56).ArraySet1CodePoint(_447_v6, (_index65).Int())
				var _542_v57 _dafny.Set
				_ = _542_v57
				_542_v57 = _dafny.SetOf(_447_v6, _447_v6, _dafny.CodePoint('l'))
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_541_v56), 0))
				_ = _index66
				var _rhs65 _dafny.CodePoint = _447_v6
				_ = _rhs65
				var _rhs66 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_439_v0, _dafny.UnicodeSeqOfUtf8Bytes("o"))
				_ = _rhs66
				var _rhs67 _dafny.Int = (func() _dafny.Int {
					if (_542_v57).IsProperSubsetOf(_542_v57) {
						return (_dafny.Zero).Minus(_517_i8)
					}
					return _517_i8
				})()
				_ = _rhs67
				var _lhs45 _dafny.Array = _541_v56
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_541_v56), 0))
				_ = _lhs46
				var _lhs47 *GlobalState = globalState
				_ = _lhs47
				(_lhs45).ArraySet1CodePoint(_rhs65, (_lhs46).Int())
				_439_v0 = _rhs66
				_lhs47.F8 = _rhs67
			} else if _source10.Is_DC15() {
				var _543___mcc_h10 _dafny.Int = _source10.Get_().(D7_DC15).Cf26
				_ = _543___mcc_h10
				var _544___mcc_h11 bool = _source10.Get_().(D7_DC15).Cf27
				_ = _544___mcc_h11
				var _545___mcc_h12 _dafny.Int = _source10.Get_().(D7_DC15).Cf28
				_ = _545___mcc_h12
				var _546_cf28 _dafny.Int = _545___mcc_h12
				_ = _546_cf28
				var _547_cf27 bool = _544___mcc_h11
				_ = _547_cf27
				var _548_cf26 _dafny.Int = _543___mcc_h10
				_ = _548_cf26
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_451_v10), 0))
				_ = _index67
				(_451_v10).ArraySet1(_517_i8, (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_451_v10), 0))
				_ = _index68
				(_451_v10).ArraySet1(_441_v2, (_index68).Int())
				var _549_v58 D6
				_ = _549_v58
				_549_v58 = Companion_D6_.Create_DC11_(_444_v3)
				var _550_v59 _dafny.Array
				_ = _550_v59
				var _nwElement0_21 bool = _444_v3
				_ = _nwElement0_21
				var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(20))
				_ = _nw81
				(_nw81).ArraySet1(_nwElement0_21, 0)
				(_nw81).ArraySet1((_dafny.IntOfInt64(383)).Cmp((_dafny.Zero).Minus(_546_cf28)) == 0, 1)
				(_nw81).ArraySet1((true) || (_444_v3), 2)
				(_nw81).ArraySet1(_444_v3, 3)
				(_nw81).ArraySet1(_444_v3, 4)
				(_nw81).ArraySet1(false, 5)
				(_nw81).ArraySet1((_549_v58).Dtor_cf18(), 6)
				(_nw81).ArraySet1(_444_v3, 7)
				(_nw81).ArraySet1(!((false) && (_444_v3)), 8)
				(_nw81).ArraySet1(_547_cf27, 9)
				(_nw81).ArraySet1(!(_547_cf27), 10)
				(_nw81).ArraySet1((_548_cf26).Cmp(Companion_Default___.Fm5(globalState)) <= 0, 11)
				(_nw81).ArraySet1((Companion_D6_.Create_DC11_(_547_cf27)).Dtor_cf18(), 12)
				(_nw81).ArraySet1(_444_v3, 13)
				(_nw81).ArraySet1(_547_cf27, 14)
				(_nw81).ArraySet1((_449_v8).IsDisjointFrom((_dafny.MultiSetOf((_dafny.Zero).Minus(_548_cf26), _546_cf28)).Update((_451_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_451_v10), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(_517_i8))), 15)
				(_nw81).ArraySet1(_547_cf27, 16)
				(_nw81).ArraySet1(_444_v3, 17)
				(_nw81).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfInt64(168))).IsDisjointFrom(_449_v8), 18)
				(_nw81).ArraySet1(_444_v3, 19)
				_550_v59 = _nw81
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_550_v59), 0))
				_ = _index69
				(_550_v59).ArraySet1(_444_v3, (_index69).Int())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_550_v59), 0))
				_ = _index70
				(_550_v59).ArraySet1(_547_cf27, (_index70).Int())
				var _551_v60 _dafny.Map
				_ = _551_v60
				_551_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Contains(_450_v9, (_450_v9).Select((Companion_Default___.SafeIndex((_451_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_451_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_450_v9).Cardinality()))).Uint32()).(_dafny.Int)), (_519_v49).Difference(_519_v49))
				var _552_v61 _dafny.Sequence
				_ = _552_v61
				_552_v61 = _dafny.SeqOf(_519_v49)
				_551_v60 = (_551_v60).Update(!_dafny.Companion_Sequence_.Contains(_439_v0, _447_v6), ((_552_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_518_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_518_v48), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_517_i8, _dafny.IntOfUint32(((_518_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(964), _dafny.ArrayLen((_518_v48), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), (_550_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_550_v59), 0))).Int()).(bool))).Cardinality()), _dafny.IntOfUint32((_552_v61).Cardinality()))).Uint32()).(_dafny.Set)).Difference(func() _dafny.Set {
					var _coll26 = _dafny.NewBuilder()
					_ = _coll26
					for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-289), _dafny.IntOfInt64(653))); ; {
						_compr_26, _ok27 := _iter27()
						if !_ok27 {
							break
						}
						var _553_v62 _dafny.Int
						_553_v62 = interface{}(_compr_26).(_dafny.Int)
						if ((_dafny.IntOfInt64(-289)).Cmp(_553_v62) <= 0) && ((_553_v62).Cmp(_dafny.IntOfInt64(653)) < 0) {
							_coll26.Add(Companion_Default___.SafeModuloInt(_553_v62, _441_v2))
						}
					}
					return _coll26.ToSet()
				}()))
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_451_v10), 0))
				_ = _index71
				(_451_v10).ArraySet1((_dafny.Zero).Minus(_517_i8), (_index71).Int())
			} else {
				var _554___mcc_h13 _dafny.Map = _source10.Get_().(D7_DC13).Cf20
				_ = _554___mcc_h13
				var _555_cf20 _dafny.Map = _554___mcc_h13
				_ = _555_cf20
				(globalState).F17 = !((!(_444_v3)) && (_dafny.Companion_Sequence_.Equal(_439_v0, _439_v0)))
				var _556_v63 _dafny.Array
				_ = _556_v63
				var _nw82 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw82
				_556_v63 = _nw82
				_556_v63 = _556_v63
				_439_v0 = _dafny.Companion_Sequence_.Concatenate(_439_v0, _439_v0)
				var _557_v64 *C1
				_ = _557_v64
				var _nw83 *C1 = New_C1_()
				_ = _nw83
				_nw83.Ctor__((_520_v50).Merge(_520_v50), _521_v51)
				_557_v64 = _nw83
			}
			var _558_v65 *C1
			_ = _558_v65
			var _nw84 *C1 = New_C1_()
			_ = _nw84
			_nw84.Ctor__(Companion_Default___.Fm26(globalState), _521_v51)
			_558_v65 = _nw84
			var _559_v66 _dafny.Array
			_ = _559_v66
			var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
			_ = _nw85
			_559_v66 = _nw85
			var _560_v67 _dafny.Map
			_ = _560_v67
			_560_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_517_i8, _dafny.IntOfInt64(670))
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_559_v66), 0))
			_ = _index72
			(_559_v66).ArraySet1(_560_v67, (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_559_v66), 0))
			_ = _index73
			var _rhs68 *C1 = _558_v65
			_ = _rhs68
			var _rhs69 _dafny.Map = (_560_v67).Update(Companion_Default___.Fm5(globalState), _517_i8)
			_ = _rhs69
			var _lhs48 _dafny.Array = _559_v66
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_559_v66), 0))
			_ = _lhs49
			_558_v65 = _rhs68
			(_lhs48).ArraySet1(_rhs69, (_lhs49).Int())
		}
		var _source11 D4 = Companion_D4_.Create_DC8_(_447_v6, _444_v3, _444_v3)
		_ = _source11
		if _source11.Is_DC7() {
			var _561___mcc_h14 _dafny.Int = _source11.Get_().(D4_DC7).Cf9
			_ = _561___mcc_h14
			var _562___mcc_h15 *C0 = _source11.Get_().(D4_DC7).Cf10
			_ = _562___mcc_h15
			var _563___mcc_h16 bool = _source11.Get_().(D4_DC7).Cf11
			_ = _563___mcc_h16
			var _564___mcc_h17 _dafny.Int = _source11.Get_().(D4_DC7).Cf12
			_ = _564___mcc_h17
			var _565_cf12 _dafny.Int = _564___mcc_h17
			_ = _565_cf12
			var _566_cf11 bool = _563___mcc_h16
			_ = _566_cf11
			var _567_cf10 *C0 = _562___mcc_h15
			_ = _567_cf10
			var _568_cf9 _dafny.Int = _561___mcc_h14
			_ = _568_cf9
			var _569_v68 _dafny.Map
			_ = _569_v68
			_569_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_444_v3), _568_cf9)
			var _570_v69 D7
			_ = _570_v69
			_570_v69 = Companion_D7_.Create_DC13_(_569_v68)
			var _571_v70 _dafny.Sequence
			_ = _571_v70
			_571_v70 = _dafny.SeqOf(_570_v69, _570_v69, Companion_D7_.Create_DC13_(_569_v68))
			(globalState).F17 = !_dafny.Companion_Sequence_.Equal(_571_v70, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_571_v70, _571_v70), (Companion_Default___.SafeIndex(_565_cf12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_571_v70, _571_v70)).Cardinality()))).Uint32(), _570_v69))
			var _572_v71 bool
			_ = _572_v71
			_572_v71 = (func() bool {
				if (_567_cf10).F27() {
					return _444_v3
				}
				return _566_cf11
			})()
			var _573_v72 _dafny.Map
			_ = _573_v72
			_573_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _566_cf11)
			var _574_v73 _dafny.Set
			_ = _574_v73
			_574_v73 = _dafny.SetOf(_dafny.IntOfInt64(-908), _565_cf12, (_573_v72).Cardinality(), _dafny.IntOfInt64(315), _441_v2)
			_572_v71 = (_574_v73).IsProperSubsetOf(_dafny.SetOf(_565_cf12))
			_439_v0 = Companion_Default___.Fm7(globalState)
			var _575_v74 D9
			_ = _575_v74
			_575_v74 = Companion_D9_.Create_DC20_((_567_cf10).F27(), false)
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_451_v10), 0))
			_ = _index74
			(_451_v10).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_575_v74).Dtor_cf35(), _444_v3)).Cardinality()), (_index74).Int())
			var _576_v75 _dafny.Map
			_ = _576_v75
			_576_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_451_v10, _574_v73)
			var _577_v76 _dafny.Set
			_ = _577_v76
			_577_v76 = _dafny.SetOf((_567_cf10).F27(), false)
			var _578_v77 _dafny.Map
			_ = _578_v77
			_578_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_576_v75, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(_568_cf9), (_dafny.Zero).Minus(_dafny.IntOfUint32((_439_v0).Cardinality())), (_577_v76).Cardinality(), _565_cf12), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}((func(_579_cf9 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_580_i10 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_579_cf9)
				}
			})(_568_cf9))))).Cardinality()))
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(123), _dafny.ArrayLen((_451_v10), 0))
			_ = _index75
			(_451_v10).ArraySet1((func() _dafny.Int {
				if (_578_v77).Contains(_576_v75) {
					return (_578_v77).Get(_576_v75).(_dafny.Int)
				}
				return _441_v2
			})(), (_index75).Int())
		} else if _source11.Is_DC8() {
			var _581___mcc_h18 _dafny.CodePoint = _source11.Get_().(D4_DC8).Cf13
			_ = _581___mcc_h18
			var _582___mcc_h19 bool = _source11.Get_().(D4_DC8).Cf14
			_ = _582___mcc_h19
			var _583___mcc_h20 bool = _source11.Get_().(D4_DC8).Cf15
			_ = _583___mcc_h20
			var _584_cf15 bool = _583___mcc_h20
			_ = _584_cf15
			var _585_cf14 bool = _582___mcc_h19
			_ = _585_cf14
			var _586_cf13 _dafny.CodePoint = _581___mcc_h18
			_ = _586_cf13
			var _587_v78 _dafny.Array
			_ = _587_v78
			var _nw86 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw86
			_587_v78 = _nw86
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_587_v78), 0))
			_ = _index76
			(_587_v78).ArraySet1(_444_v3, (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_587_v78), 0))
			_ = _index77
			(_587_v78).ArraySet1(_444_v3, (_index77).Int())
			(globalState).F19 = (_587_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_587_v78), 0))).Int()).(bool)
			var _588_v79 D6
			_ = _588_v79
			_588_v79 = Companion_D6_.Create_DC11_(!_dafny.Companion_Sequence_.Contains(_439_v0, _447_v6))
			var _source12 D6 = _588_v79
			_ = _source12
			if _source12.Is_DC11() {
				var _589___mcc_h22 bool = _source12.Get_().(D6_DC11).Cf18
				_ = _589___mcc_h22
				var _590_cf18 bool = _589___mcc_h22
				_ = _590_cf18
				_587_v78 = _587_v78
				var _591_v80 _dafny.MultiSet
				_ = _591_v80
				_591_v80 = _dafny.MultiSetOf(_439_v0)
				var _rhs70 _dafny.Int = (_dafny.IntOfUint32((_439_v0).Cardinality())).Plus(_441_v2)
				_ = _rhs70
				var _rhs71 _dafny.Int = (_591_v80).Cardinality()
				_ = _rhs71
				var _lhs50 *GlobalState = globalState
				_ = _lhs50
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				_lhs50.F8 = _rhs70
				_lhs51.F8 = _rhs71
				(globalState).F8 = _dafny.IntOfInt64(-443)
				(globalState).F5 = _444_v3
			} else if _source12.Is_DC10() {
				var _592___mcc_h23 _dafny.Array = _source12.Get_().(D6_DC10).Cf17
				_ = _592___mcc_h23
				var _593_cf17 _dafny.Array = _592___mcc_h23
				_ = _593_cf17
				_441_v2 = _441_v2
				var _rhs72 bool = (_587_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_587_v78), 0))).Int()).(bool)
				_ = _rhs72
				var _rhs73 _dafny.Int = _441_v2
				_ = _rhs73
				var _lhs52 *GlobalState = globalState
				_ = _lhs52
				_585_cf14 = _rhs72
				_lhs52.F8 = _rhs73
				_450_v9 = _450_v9
				var _594_v81 _dafny.Map
				_ = _594_v81
				_594_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _584_cf15)
				var _595_v82 _dafny.Sequence
				_ = _595_v82
				_595_v82 = _dafny.SeqOf(_594_v81, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_v3, _585_cf14))
				var _rhs74 _dafny.Sequence = _439_v0
				_ = _rhs74
				var _rhs75 _dafny.Sequence = _595_v82
				_ = _rhs75
				_439_v0 = _rhs74
				_595_v82 = _rhs75
			} else {
				var _596___mcc_h24 D6 = _source12.Get_().(D6_DC12).Cf19
				_ = _596___mcc_h24
				var _597_cf19 D6 = _596___mcc_h24
				_ = _597_cf19
				(globalState).F8 = (_dafny.IntOfInt64(-467))
				_584_cf15 = (_585_cf14) == (!(!((_587_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(991), _dafny.ArrayLen((_587_v78), 0))).Int()).(bool))))
				var _598_v83 _dafny.Array
				_ = _598_v83
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw87
				_598_v83 = _nw87
				var _599_v84 _dafny.Map
				_ = _599_v84
				_599_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_584_cf15, false)
				var _600_v85 _dafny.Map
				_ = _600_v85
				_600_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_598_v83, (_599_v84).Cardinality())
				_600_v85 = (_600_v85).Update(_598_v83, _441_v2)
				var _601_v86 _dafny.Int
				_ = _601_v86
				var _602_v87 _dafny.Array
				_ = _602_v87
				var _603_v88 bool
				_ = _603_v88
				var _out24 _dafny.Int
				_ = _out24
				var _out25 _dafny.Array
				_ = _out25
				var _out26 bool
				_ = _out26
				_out24, _out25, _out26 = (_this).M5(_444_v3, globalState)
				_601_v86 = _out24
				_602_v87 = _out25
				_603_v88 = _out26
			}
			_585_cf14 = _444_v3
		} else {
			var _604___mcc_h21 _dafny.Set = _source11.Get_().(D4_DC6).Cf8
			_ = _604___mcc_h21
			var _605_cf8 _dafny.Set = _604___mcc_h21
			_ = _605_cf8
			var _606_v89 _dafny.Set
			_ = _606_v89
			_606_v89 = _dafny.SetOf(_444_v3, _444_v3, _444_v3, _444_v3)
			if (_606_v89).IsSubsetOf(_606_v89) {
				var _607_v90 _dafny.Array
				_ = _607_v90
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_11
				var _nw88 _dafny.Array
				_ = _nw88
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw88 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) D7 = (func(_608_v3 bool, _609_v2 _dafny.Int) func(_dafny.Int) D7 {
						return func(_610_i11 _dafny.Int) D7 {
							return Companion_D7_.Create_DC13_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_608_v3, _609_v2))
						}
					})(_444_v3, _441_v2)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw88 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw88).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw88).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_607_v90 = _nw88
				var _611_v91 _dafny.Sequence
				_ = _611_v91
				_611_v91 = _dafny.SeqOf(_607_v90, _607_v90, _607_v90)
				var _612_v92 _dafny.Array
				_ = _612_v92
				var _nwElement0_22 _dafny.Array = _607_v90
				_ = _nwElement0_22
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(29))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_22, 0)
				(_nw89).ArraySet1(_607_v90, 1)
				(_nw89).ArraySet1(_607_v90, 2)
				(_nw89).ArraySet1((_611_v91).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_611_v91).Cardinality()))).Uint32()).(_dafny.Array), 3)
				(_nw89).ArraySet1(_607_v90, 4)
				(_nw89).ArraySet1(_607_v90, 5)
				(_nw89).ArraySet1((func() _dafny.Array {
					if _444_v3 {
						return _607_v90
					}
					return _607_v90
				})(), 6)
				(_nw89).ArraySet1(_607_v90, 7)
				(_nw89).ArraySet1(_607_v90, 8)
				(_nw89).ArraySet1(_607_v90, 9)
				(_nw89).ArraySet1(_607_v90, 10)
				(_nw89).ArraySet1(_607_v90, 11)
				(_nw89).ArraySet1(_607_v90, 12)
				(_nw89).ArraySet1((_611_v91).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_611_v91).Cardinality()))).Uint32()).(_dafny.Array), 13)
				(_nw89).ArraySet1(_607_v90, 14)
				(_nw89).ArraySet1(_607_v90, 15)
				(_nw89).ArraySet1(_607_v90, 16)
				(_nw89).ArraySet1(_607_v90, 17)
				(_nw89).ArraySet1(_607_v90, 18)
				(_nw89).ArraySet1(_607_v90, 19)
				(_nw89).ArraySet1(_607_v90, 20)
				(_nw89).ArraySet1(_607_v90, 21)
				(_nw89).ArraySet1(_607_v90, 22)
				(_nw89).ArraySet1(_607_v90, 23)
				(_nw89).ArraySet1(_607_v90, 24)
				(_nw89).ArraySet1(_607_v90, 25)
				(_nw89).ArraySet1(_607_v90, 26)
				(_nw89).ArraySet1(_607_v90, 27)
				(_nw89).ArraySet1(_607_v90, 28)
				_612_v92 = _nw89
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_612_v92), 0))
				_ = _index78
				(_612_v92).ArraySet1(_607_v90, (_index78).Int())
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_612_v92), 0))
				_ = _index79
				(_612_v92).ArraySet1(_607_v90, (_index79).Int())
				var _613_v93 _dafny.Sequence
				_ = _613_v93
				_613_v93 = _dafny.SeqOf(_439_v0)
				var _614_v94 _dafny.Map
				_ = _614_v94
				_614_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_439_v0).Cardinality())), (_613_v93).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_613_v93).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _615_v95 T1
				_ = _615_v95
				var _nw90 *C0 = New_C0_()
				_ = _nw90
				_nw90.Ctor__(true)
				_615_v95 = _nw90
				var _616_v96 T0
				_ = _616_v96
				var _nw91 *C1 = New_C1_()
				_ = _nw91
				_nw91.Ctor__(_614_v94, _615_v95)
				_616_v96 = _nw91
				_616_v96 = _616_v96
				var _617_v97 _dafny.Map
				_ = _617_v97
				_617_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((Companion_Default___.Fm0(_448_v7, globalState)) || (_444_v3)), (_450_v9).Select((Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_450_v9).Cardinality()))).Uint32()).(_dafny.Int))
				var _rhs76 bool = _444_v3
				_ = _rhs76
				var _rhs77 _dafny.Map = _617_v97
				_ = _rhs77
				var _rhs78 bool = (_449_v8).IsSubsetOf((_449_v8).Intersection(_449_v8))
				_ = _rhs78
				var _lhs53 *GlobalState = globalState
				_ = _lhs53
				var _lhs54 *GlobalState = globalState
				_ = _lhs54
				_lhs53.F17 = _rhs76
				_617_v97 = _rhs77
				_lhs54.F5 = _rhs78
				(globalState).F8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_452_v11, _452_v11), (func() _dafny.Sequence {
					if _444_v3 {
						return _dafny.Companion_Sequence_.Update(_452_v11, (Companion_Default___.SafeIndex((_445_v4).Cardinality(), _dafny.IntOfUint32((_452_v11).Cardinality()))).Uint32(), _451_v10)
					}
					return _452_v11
				})())).Cardinality())
				_450_v9 = _450_v9
			} else {
				var _618_v98 _dafny.MultiSet
				_ = _618_v98
				_618_v98 = _dafny.MultiSetOf(_447_v6, Companion_Default___.Fm8(globalState))
				_441_v2 = (func() _dafny.Int {
					if (_618_v98).Contains(_447_v6) {
						return (_618_v98).Multiplicity(_447_v6)
					}
					return (_606_v89).Cardinality()
				})()
				(globalState).F5 = Companion_Default___.Fm0(_448_v7, globalState)
				_439_v0 = _439_v0
				var _619_v99 *C0
				_ = _619_v99
				var _nw92 *C0 = New_C0_()
				_ = _nw92
				_nw92.Ctor__(_444_v3)
				_619_v99 = _nw92
				var _620_v100 _dafny.Set
				_ = _620_v100
				_620_v100 = _dafny.SetOf(_619_v99)
				(globalState).F5 = (_620_v100).IsProperSubsetOf(_620_v100)
				var _621_v101 _dafny.Map
				_ = _621_v101
				_621_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v2, Companion_Default___.Fm0(_448_v7, globalState))
				var _622_v102 _dafny.Set
				_ = _622_v102
				_622_v102 = _dafny.SetOf(_449_v8, _dafny.MultiSetOf(Companion_Default___.Fm18(_dafny.MultiSetFromSeq(_448_v7), Companion_Default___.Fm27(_441_v2, globalState), (_619_v99).F27(), globalState), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lgd")).Cardinality())), _449_v8, _449_v8)
				(globalState).F22 = (func() bool {
					if (_621_v101).Contains((_622_v102).Cardinality()) {
						return (_621_v101).Get((_622_v102).Cardinality()).(bool)
					}
					return !(_444_v3)
				})()
			}
			var _623_v103 _dafny.Array
			_ = _623_v103
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw93
			_623_v103 = _nw93
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_623_v103), 0))
			_ = _index80
			(_623_v103).ArraySet1(false, (_index80).Int())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_623_v103), 0))
			_ = _index81
			(_623_v103).ArraySet1(_444_v3, (_index81).Int())
			var _624_v104 _dafny.Array
			_ = _624_v104
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw94
			_624_v104 = _nw94
			var _625_v105 _dafny.Array
			_ = _625_v105
			var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
			_ = _nw95
			_625_v105 = _nw95
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_624_v104), 0))
			_ = _index82
			(_624_v104).ArraySet1(_625_v105, (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_624_v104), 0))
			_ = _index83
			(_624_v104).ArraySet1(_625_v105, (_index83).Int())
			var _626_v106 _dafny.Map
			_ = _626_v106
			_626_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_cf8, _dafny.UnicodeSeqOfUtf8Bytes("krramqlq"))
			var _627_v107 T1
			_ = _627_v107
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__((_623_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_623_v103), 0))).Int()).(bool))
			_627_v107 = _nw96
			var _628_v108 *C1
			_ = _628_v108
			var _nw97 *C1 = New_C1_()
			_ = _nw97
			_nw97.Ctor__(_626_v106, _627_v107)
			_628_v108 = _nw97
		}
		var _629_v109 _dafny.Map
		_ = _629_v109
		_629_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_439_v0, Companion_Default___.Fm5(globalState))
		var _rhs79 _dafny.Int = _441_v2
		_ = _rhs79
		var _rhs80 _dafny.Int = (func() _dafny.Int {
			if (_629_v109).Contains(_dafny.Companion_Sequence_.Concatenate(_439_v0, _439_v0)) {
				return (_629_v109).Get(_dafny.Companion_Sequence_.Concatenate(_439_v0, _439_v0)).(_dafny.Int)
			}
			return _441_v2
		})()
		_ = _rhs80
		_441_v2 = _rhs79
		_441_v2 = _rhs80
	}
}
func (_this *C2) M5(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Array, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var _630_v0 _dafny.Sequence
		_ = _630_v0
		_630_v0 = _dafny.SeqOf(p0, p0)
		(globalState).F5 = Companion_Default___.Fm0(_630_v0, globalState)
		var _631_v1 _dafny.Array
		_ = _631_v1
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_12
		var _nw98 _dafny.Array
		_ = _nw98
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw98 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) bool = (func(_632_p0 bool) func(_dafny.Int) bool {
				return func(_633_i1 _dafny.Int) bool {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rgdrpsk"), _dafny.IntOfInt64(107))).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(124))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg35 _dafny.Int) interface{} {
							return coer35(arg35)
						}
					}(func(_634_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('k')
					})), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(941), _632_p0)).Cardinality())))
				}
			})(p0)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw98 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw98).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw98).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_631_v1 = _nw98
		for _iter28 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_631_v1), 0))); ; {
			_guard_loop_1, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			var _635_i0 _dafny.Int
			_635_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_635_i0).Sign() != -1) && ((_635_i0).Cmp(_dafny.ArrayLen((_631_v1), 0)) < 0)) {
				(_631_v1).ArraySet1(p0, (_635_i0).Int())
			}
		}
		var _636_v2 _dafny.CodePoint
		_ = _636_v2
		_636_v2 = _dafny.CodePoint('g')
		var _637_v3 _dafny.Array
		_ = _637_v3
		var _nwElement0_23 _dafny.CodePoint = _636_v2
		_ = _nwElement0_23
		var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(9))
		_ = _nw99
		(_nw99).ArraySet1CodePoint(_nwElement0_23, 0)
		(_nw99).ArraySet1CodePoint(_dafny.CodePoint('i'), 1)
		(_nw99).ArraySet1CodePoint(_636_v2, 2)
		(_nw99).ArraySet1CodePoint(_636_v2, 3)
		(_nw99).ArraySet1CodePoint(_636_v2, 4)
		(_nw99).ArraySet1CodePoint(_636_v2, 5)
		(_nw99).ArraySet1CodePoint(_636_v2, 6)
		(_nw99).ArraySet1CodePoint(_636_v2, 7)
		(_nw99).ArraySet1CodePoint(_636_v2, 8)
		_637_v3 = _nw99
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))
		_ = _index84
		(_637_v3).ArraySet1CodePoint(_636_v2, (_index84).Int())
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))
		_ = _index85
		(_637_v3).ArraySet1CodePoint(_636_v2, (_index85).Int())
		for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_631_v1), 0))); ; {
			_guard_loop_2, _ok29 := _iter29()
			if !_ok29 {
				break
			}
			var _638_i3 _dafny.Int
			_638_i3 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_638_i3).Sign() != -1) && ((_638_i3).Cmp(_dafny.ArrayLen((_631_v1), 0)) < 0)) {
				(_631_v1).ArraySet1(p0, (_638_i3).Int())
			}
		}
		var _639_v4 _dafny.Int
		_ = _639_v4
		_639_v4 = _dafny.IntOfInt64(790)
		var _hi4 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf((_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()))).Cardinality())).Times(_639_v4)
		_ = _hi4
		for _640_i4 := _639_v4; _640_i4.Cmp(_hi4) < 0; _640_i4 = _640_i4.Plus(_dafny.One) {
			var _641_v5 _dafny.Set
			_ = _641_v5
			_641_v5 = _dafny.SetOf(_636_v2)
			_641_v5 = _641_v5
			var _642_v6 D7
			_ = _642_v6
			_642_v6 = Companion_D7_.Create_DC15_((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_639_v4, _639_v4)), p0, _640_i4)
			var _source13 D7 = _642_v6
			_ = _source13
			if _source13.Is_DC14() {
				var _643___mcc_h0 _dafny.Int = _source13.Get_().(D7_DC14).Cf21
				_ = _643___mcc_h0
				var _644___mcc_h1 bool = _source13.Get_().(D7_DC14).Cf22
				_ = _644___mcc_h1
				var _645___mcc_h2 bool = _source13.Get_().(D7_DC14).Cf23
				_ = _645___mcc_h2
				var _646___mcc_h3 bool = _source13.Get_().(D7_DC14).Cf24
				_ = _646___mcc_h3
				var _647___mcc_h4 _dafny.Int = _source13.Get_().(D7_DC14).Cf25
				_ = _647___mcc_h4
				var _648_cf25 _dafny.Int = _647___mcc_h4
				_ = _648_cf25
				var _649_cf24 bool = _646___mcc_h3
				_ = _649_cf24
				var _650_cf23 bool = _645___mcc_h2
				_ = _650_cf23
				var _651_cf22 bool = _644___mcc_h1
				_ = _651_cf22
				var _652_cf21 _dafny.Int = _643___mcc_h0
				_ = _652_cf21
				var _653_v7 _dafny.MultiSet
				_ = _653_v7
				_653_v7 = _dafny.MultiSetOf(_651_cf22)
				var _654_v8 *C0
				_ = _654_v8
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(_649_cf24)
				_654_v8 = _nw100
				var _655_v9 D4
				_ = _655_v9
				_655_v9 = Companion_D4_.Create_DC7_((_653_v7).Cardinality(), _654_v8, _649_cf24, _640_i4)
				_655_v9 = _655_v9
				(globalState).F19 = _650_cf23
				var _656_v10 _dafny.Map
				_ = _656_v10
				_656_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_651_cf22, !(_649_cf24))
				var _657_v11 _dafny.Sequence
				_ = _657_v11
				_657_v11 = _dafny.SeqOf(_656_v10)
				var _658_v12 D10
				_ = _658_v12
				_658_v12 = Companion_D10_.Create_DC21_(_656_v10)
				var _659_v13 _dafny.Map
				_ = _659_v13
				_659_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()), _650_cf23)
				var _660_v14 _dafny.MultiSet
				_ = _660_v14
				_660_v14 = _dafny.MultiSetOf(_dafny.IntOfInt64(906))
				var _661_v15 _dafny.Map
				_ = _661_v15
				_661_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_660_v14).Cardinality(), _656_v10)
				var _pat_let_tv11 = p0
				_ = _pat_let_tv11
				var _pat_let_tv12 = _659_v13
				_ = _pat_let_tv12
				var _pat_let_tv13 = _637_v3
				_ = _pat_let_tv13
				var _pat_let_tv14 = _637_v3
				_ = _pat_let_tv14
				var _pat_let_tv15 = _659_v13
				_ = _pat_let_tv15
				var _pat_let_tv16 = _637_v3
				_ = _pat_let_tv16
				var _pat_let_tv17 = _637_v3
				_ = _pat_let_tv17
				var _pat_let_tv18 = _654_v8
				_ = _pat_let_tv18
				var _662_v16 _dafny.Array
				_ = _662_v16
				var _nwElement0_24 _dafny.Map = (_657_v11).Select((Companion_Default___.SafeIndex(_648_cf25, _dafny.IntOfUint32((_657_v11).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _nwElement0_24
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(13))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_24, 0)
				(_nw101).ArraySet1(_656_v10, 1)
				(_nw101).ArraySet1(_656_v10, 2)
				(_nw101).ArraySet1(((func(_pat_let6_0 D10) D10 {
					return func(_663_dt__update__tmp_h0 D10) D10 {
						return func(_pat_let7_0 _dafny.Map) D10 {
							return func(_664_dt__update_hcf37_h0 _dafny.Map) D10 {
								return Companion_D10_.Create_DC21_(_664_dt__update_hcf37_h0)
							}(_pat_let7_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv11, (func() bool {
							if (_pat_let_tv12).Contains((_pat_let_tv14).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_pat_let_tv13), 0))).Int())) {
								return (_pat_let_tv15).Get((_pat_let_tv17).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_pat_let_tv16), 0))).Int())).(bool)
							}
							return (_pat_let_tv18).F27()
						})()))
					}(_pat_let6_0)
				}(_658_v12)).Dtor_cf37()).Update(p0, _650_cf23), 3)
				(_nw101).ArraySet1(_656_v10, 4)
				(_nw101).ArraySet1(_656_v10, 5)
				(_nw101).ArraySet1(_656_v10, 6)
				(_nw101).ArraySet1((func() _dafny.Map {
					if (_661_v15).Contains(_652_cf21) {
						return (_661_v15).Get(_652_cf21).(_dafny.Map)
					}
					return _656_v10
				})(), 7)
				(_nw101).ArraySet1(_656_v10, 8)
				(_nw101).ArraySet1(_656_v10, 9)
				(_nw101).ArraySet1(_656_v10, 10)
				(_nw101).ArraySet1(_656_v10, 11)
				(_nw101).ArraySet1((func() _dafny.Map {
					if Companion_Default___.Fm0(_630_v0, globalState) {
						return (_658_v12).Dtor_cf37()
					}
					return _656_v10
				})(), 12)
				_662_v16 = _nw101
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_662_v16), 0))
				_ = _index86
				(_662_v16).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _650_cf23), (_index86).Int())
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_662_v16), 0))
				_ = _index87
				(_662_v16).ArraySet1((((_656_v10).Update(_649_cf24, _651_cf22)).Merge(_656_v10)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_650_cf23, _649_cf24)).Update(p0, _649_cf24)), (_index87).Int())
				var _665_v17 _dafny.Sequence
				_ = _665_v17
				_665_v17 = _dafny.SeqOf((_dafny.Zero).Minus(_652_cf21), _640_i4, _639_v4)
				var _666_v18 _dafny.Sequence
				_ = _666_v18
				_666_v18 = _dafny.SeqOf(_665_v17, _665_v17)
				var _667_v19 _dafny.Array
				_ = _667_v19
				var _nwElement0_25 _dafny.Sequence = _665_v17
				_ = _nwElement0_25
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(26))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_25, 0)
				(_nw102).ArraySet1(_665_v17, 1)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm13(globalState), _dafny.SeqOf(_648_cf25)), 2)
				(_nw102).ArraySet1((func() _dafny.Sequence {
					if (_654_v8).F27() {
						return _665_v17
					}
					return _dafny.SeqOf(_652_cf21, _652_cf21)
				})(), 3)
				(_nw102).ArraySet1(_dafny.SeqOf(_640_i4, _648_cf25), 4)
				(_nw102).ArraySet1((_666_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(317))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_668_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_669_i5 _dafny.Int) _dafny.CodePoint {
						return _668_v2
					}
				})(_636_v2)))).Cardinality()), _dafny.IntOfUint32((_666_v18).Cardinality()))).Uint32()).(_dafny.Sequence), 5)
				(_nw102).ArraySet1(_665_v17, 6)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Update(_665_v17, (Companion_Default___.SafeIndex(_640_i4, _dafny.IntOfUint32((_665_v17).Cardinality()))).Uint32(), _dafny.IntOfInt64(-505)), 7)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(802))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_670_cf25 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_671_i6 _dafny.Int) _dafny.Int {
						return _670_cf25
					}
				})(_648_cf25))), _665_v17), 8)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(127))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg38 _dafny.Int) interface{} {
						return coer38(arg38)
					}
				}((func(_672_cf24 bool) func(_dafny.Int) _dafny.Int {
					return func(_673_i7 _dafny.Int) _dafny.Int {
						return (_dafny.SetOf(_672_cf24)).Cardinality()
					}
				})(_649_cf24))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(457))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_674_cf21 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_675_i8 _dafny.Int) _dafny.Int {
						return _674_cf21
					}
				})(_652_cf21)))), 9)
				(_nw102).ArraySet1(_665_v17, 10)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _665_v17), 11)
				(_nw102).ArraySet1(_665_v17, 12)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _665_v17), 13)
				(_nw102).ArraySet1(_665_v17, 14)
				(_nw102).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_665_v17).Cardinality()))), 15)
				(_nw102).ArraySet1(_665_v17, 16)
				(_nw102).ArraySet1(_665_v17, 17)
				(_nw102).ArraySet1(_665_v17, 18)
				(_nw102).ArraySet1(_665_v17, 19)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _dafny.SeqOf(_dafny.IntOfInt64(604), _640_i4)), 20)
				(_nw102).ArraySet1(_665_v17, 21)
				(_nw102).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _665_v17), 22)
				(_nw102).ArraySet1(_665_v17, 23)
				(_nw102).ArraySet1(_665_v17, 24)
				(_nw102).ArraySet1(_665_v17, 25)
				_667_v19 = _nw102
				var _676_v20 _dafny.Map
				_ = _676_v20
				_676_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_654_v8).F27(), _665_v17)
				var _nwElement0_26 _dafny.Sequence = _665_v17
				_ = _nwElement0_26
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(17))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_26, 0)
				(_nw103).ArraySet1(_665_v17, 1)
				(_nw103).ArraySet1((func() _dafny.Sequence {
					if (_654_v8).F27() {
						return _665_v17
					}
					return _665_v17
				})(), 2)
				(_nw103).ArraySet1(_665_v17, 3)
				(_nw103).ArraySet1(_665_v17, 4)
				(_nw103).ArraySet1(_665_v17, 5)
				(_nw103).ArraySet1(_665_v17, 6)
				(_nw103).ArraySet1(_665_v17, 7)
				(_nw103).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(184))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_677_cf25 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_678_i9 _dafny.Int) _dafny.Int {
						return _677_cf25
					}
				})(_648_cf25))), 8)
				(_nw103).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(870))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_679_i10 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(702)
				})), 9)
				(_nw103).ArraySet1((func() _dafny.Sequence {
					if (_676_v20).Contains(true) {
						return (_676_v20).Get(true).(_dafny.Sequence)
					}
					return _dafny.Companion_Sequence_.Update(_665_v17, (Companion_Default___.SafeIndex(_648_cf25, _dafny.IntOfUint32((_665_v17).Cardinality()))).Uint32(), _652_cf21)
				})(), 10)
				(_nw103).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_680_i4 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_681_i11 _dafny.Int) _dafny.Int {
						return _680_i4
					}
				})(_640_i4))), 11)
				(_nw103).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(292))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_682_v8 *C0, _683_v7 _dafny.MultiSet, _684_cf21 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_685_i12 _dafny.Int) _dafny.Int {
						return (func() _dafny.Int {
							if (_683_v7).Contains((_682_v8).F27()) {
								return (_683_v7).Multiplicity((_682_v8).F27())
							}
							return _684_cf21
						})()
					}
				})(_654_v8, _653_v7, _652_cf21)))), 12)
				(_nw103).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_665_v17, _665_v17), 13)
				(_nw103).ArraySet1(_665_v17, 14)
				(_nw103).ArraySet1(_665_v17, 15)
				(_nw103).ArraySet1(_665_v17, 16)
				_667_v19 = _nw103
			} else if _source13.Is_DC15() {
				var _686___mcc_h5 _dafny.Int = _source13.Get_().(D7_DC15).Cf26
				_ = _686___mcc_h5
				var _687___mcc_h6 bool = _source13.Get_().(D7_DC15).Cf27
				_ = _687___mcc_h6
				var _688___mcc_h7 _dafny.Int = _source13.Get_().(D7_DC15).Cf28
				_ = _688___mcc_h7
				var _689_cf28 _dafny.Int = _688___mcc_h7
				_ = _689_cf28
				var _690_cf27 bool = _687___mcc_h6
				_ = _690_cf27
				var _691_cf26 _dafny.Int = _686___mcc_h5
				_ = _691_cf26
				var _692_v21 _dafny.Map
				_ = _692_v21
				_692_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), p0)
				var _693_v22 _dafny.Sequence
				_ = _693_v22
				_693_v22 = _dafny.UnicodeSeqOfUtf8Bytes("y")
				var _694_v23 _dafny.Map
				_ = _694_v23
				_694_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_i4, _636_v2)
				var _695_v24 _dafny.Set
				_ = _695_v24
				_695_v24 = _dafny.SetOf(_690_cf27, _690_cf27)
				var _696_v25 _dafny.MultiSet
				_ = _696_v25
				_696_v25 = _dafny.MultiSetOf(p0, _690_cf27)
				var _697_v26 D11
				_ = _697_v26
				_697_v26 = Companion_D11_.Create_DC24_(_696_v25)
				var _698_v27 D8
				_ = _698_v27
				_698_v27 = Companion_D8_.Create_DC17_((_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()), _636_v2, _dafny.IntOfUint32((_693_v22).Cardinality()))
				var _699_v28 _dafny.MultiSet
				_ = _699_v28
				_699_v28 = _dafny.MultiSetOf((_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()), (_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()))
				var _700_v29 _dafny.Sequence
				_ = _700_v29
				_700_v29 = _dafny.SeqOf(_689_cf28, Companion_Default___.Fm18(_dafny.MultiSetOf(!(_690_cf27), !(_690_cf27)), _698_v27, !(p0), globalState), _691_cf26, _691_cf26)
				var _701_v30 _dafny.Array
				_ = _701_v30
				var _nwElement0_27 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_630_v0, globalState), !(p0))).Merge(_692_v21)).Cardinality()
				_ = _nwElement0_27
				var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(25))
				_ = _nw104
				(_nw104).ArraySet1(_nwElement0_27, 0)
				(_nw104).ArraySet1(_689_cf28, 1)
				(_nw104).ArraySet1((_dafny.IntOfUint32((_693_v22).Cardinality())).Times(_640_i4), 2)
				(_nw104).ArraySet1((Companion_Default___.Fm18(_dafny.MultiSetOf(p0), Companion_D8_.Create_DC17_((_637_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_637_v3), 0))).Int()), (func() _dafny.CodePoint {
					if (_694_v23).Contains((_695_v24).Cardinality()) {
						return (_694_v23).Get((_695_v24).Cardinality()).(_dafny.CodePoint)
					}
					return _636_v2
				})(), _691_cf26), _690_cf27, globalState)).Times(_691_cf26), 3)
				(_nw104).ArraySet1(Companion_Default___.Fm18((_697_v26).Dtor_cf41(), _698_v27, false, globalState), 4)
				(_nw104).ArraySet1(Companion_Default___.Fm5(globalState), 5)
				(_nw104).ArraySet1((_699_v28).Cardinality(), 6)
				(_nw104).ArraySet1((_dafny.Zero).Minus(_640_i4), 7)
				(_nw104).ArraySet1(_689_cf28, 8)
				(_nw104).ArraySet1(_dafny.IntOfInt64(517), 9)
				(_nw104).ArraySet1((_696_v25).Cardinality(), 10)
				(_nw104).ArraySet1((_dafny.IntOfUint32((_630_v0).Cardinality())).Plus(_640_i4), 11)
				(_nw104).ArraySet1((func() _dafny.Int {
					if _690_cf27 {
						return _dafny.IntOfUint32((_630_v0).Cardinality())
					}
					return _691_cf26
				})(), 12)
				(_nw104).ArraySet1(_689_cf28, 13)
				(_nw104).ArraySet1(Companion_Default___.Fm18(_696_v25, _698_v27, _690_cf27, globalState), 14)
				(_nw104).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_639_v4)), 15)
				(_nw104).ArraySet1(Companion_Default___.SafeModuloInt(_639_v4, (_dafny.Zero).Minus(_639_v4)), 16)
				(_nw104).ArraySet1((_640_i4).Times(_dafny.IntOfInt64(235)), 17)
				(_nw104).ArraySet1((_640_i4).Times(_691_cf26), 18)
				(_nw104).ArraySet1((_dafny.IntOfUint32((_700_v29).Cardinality())).Minus(_639_v4), 19)
				(_nw104).ArraySet1((_dafny.IntOfUint32((_693_v22).Cardinality())).Minus(_691_cf26), 20)
				(_nw104).ArraySet1(_dafny.IntOfInt64(690), 21)
				(_nw104).ArraySet1((_dafny.Zero).Minus(_691_cf26), 22)
				(_nw104).ArraySet1((func() _dafny.Int {
					if _690_cf27 {
						return (_dafny.Zero).Minus(_dafny.IntOfUint32((_630_v0).Cardinality()))
					}
					return _640_i4
				})(), 23)
				(_nw104).ArraySet1(_639_v4, 24)
				_701_v30 = _nw104
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_701_v30), 0))
				_ = _index88
				(_701_v30).ArraySet1(_691_cf26, (_index88).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_701_v30), 0))
				_ = _index89
				(_701_v30).ArraySet1(Companion_Default___.SafeModuloInt(((_696_v25).Update(p0, Companion_Default___.Abs((_700_v29).Select((Companion_Default___.SafeIndex(_639_v4, _dafny.IntOfUint32((_700_v29).Cardinality()))).Uint32()).(_dafny.Int)))).Cardinality(), _dafny.IntOfInt64(-466)), (_index89).Int())
				var _702_v31 _dafny.Map
				_ = _702_v31
				_702_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_636_v2, p0)
				var _703_v32 _dafny.Map
				_ = _703_v32
				_703_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_702_v31, (_630_v0).Select((Companion_Default___.SafeIndex(_689_cf28, _dafny.IntOfUint32((_630_v0).Cardinality()))).Uint32()).(bool))
				_703_v32 = (_703_v32).Update(_702_v31, p0)
				_691_cf26 = _691_cf26
				_689_cf28 = ((_700_v29).Select((Companion_Default___.SafeIndex((_701_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_701_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_700_v29).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_689_cf28)
			} else {
				var _704___mcc_h8 _dafny.Map = _source13.Get_().(D7_DC13).Cf20
				_ = _704___mcc_h8
				var _705_cf20 _dafny.Map = _704___mcc_h8
				_ = _705_cf20
				_636_v2 = _636_v2
				var _706_v33 _dafny.Sequence
				_ = _706_v33
				_706_v33 = _dafny.UnicodeSeqOfUtf8Bytes("nc")
				var _707_v34 _dafny.Map
				_ = _707_v34
				_707_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_706_v33, !(p0))
				var _708_v35 D3
				_ = _708_v35
				_708_v35 = Companion_D3_.Create_DC4_((_707_v34).Cardinality(), p0, p0)
				var _709_v36 _dafny.Map
				_ = _709_v36
				_709_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-165), _637_v3)
				var _710_v37 _dafny.Map
				_ = _710_v37
				_710_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_708_v35, (_dafny.MultiSetOf((_dafny.Zero).Minus(_639_v4), (_709_v36).Cardinality())).Cardinality())
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))
				_ = _index90
				(_631_v1).ArraySet1(p0, (_index90).Int())
				var _711_v38 _dafny.Sequence
				_ = _711_v38
				_711_v38 = _dafny.SeqOf(_640_i4, _640_i4)
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))
				_ = _index91
				var _rhs81 _dafny.Map = _710_v37
				_ = _rhs81
				var _rhs82 bool = !((p0) || (p0))
				_ = _rhs82
				var _rhs83 _dafny.Int = (_639_v4).Times(_640_i4)
				_ = _rhs83
				var _rhs84 _dafny.Sequence = (func() _dafny.Sequence {
					if p0 {
						return _630_v0
					}
					return _630_v0
				})()
				_ = _rhs84
				var _rhs85 _dafny.Int = Companion_Default___.SafeDivisionInt((_640_i4).Times(_dafny.IntOfUint32((_711_v38).Cardinality())), _dafny.IntOfUint32((_706_v33).Cardinality()))
				_ = _rhs85
				var _lhs55 _dafny.Array = _631_v1
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))
				_ = _lhs56
				var _lhs57 *GlobalState = globalState
				_ = _lhs57
				_710_v37 = _rhs81
				(_lhs55).ArraySet1(_rhs82, (_lhs56).Int())
				_639_v4 = _rhs83
				_630_v0 = _rhs84
				_lhs57.F8 = _rhs85
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))
				_ = _index92
				var _rhs86 bool = (_631_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))).Int()).(bool)
				_ = _rhs86
				var _rhs87 _dafny.Int = Companion_Default___.SafeDivisionInt(_640_i4, _640_i4)
				_ = _rhs87
				var _lhs58 _dafny.Array = _631_v1
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_631_v1), 0))
				_ = _lhs59
				var _lhs60 *GlobalState = globalState
				_ = _lhs60
				(_lhs58).ArraySet1(_rhs86, (_lhs59).Int())
				_lhs60.F8 = _rhs87
				(globalState).F8 = (_dafny.Zero).Minus(_639_v4)
			}
			var _712_v39 _dafny.Sequence
			_ = _712_v39
			_712_v39 = _dafny.SeqOf(_630_v0, _630_v0)
			var _713_v40 _dafny.Map
			_ = _713_v40
			_713_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_i4, _639_v4)
			var _714_v41 *C0
			_ = _714_v41
			var _nw105 *C0 = New_C0_()
			_ = _nw105
			_nw105.Ctor__(p0)
			_714_v41 = _nw105
			var _715_v42 D4
			_ = _715_v42
			_715_v42 = Companion_D4_.Create_DC7_(_639_v4, _714_v41, true, _639_v4)
			var _716_v43 _dafny.MultiSet
			_ = _716_v43
			_716_v43 = _dafny.MultiSetOf(_640_i4, (_715_v42).Dtor_cf12(), _dafny.IntOfInt64(557))
			var _717_v44 T1
			_ = _717_v44
			var _nw106 *C0 = New_C0_()
			_ = _nw106
			_nw106.Ctor__((_714_v41).F27())
			_717_v44 = _nw106
			var _718_v45 _dafny.Map
			_ = _718_v45
			_718_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_717_v44, _dafny.IntOfInt64(468))
			var _719_v46 _dafny.Sequence
			_ = _719_v46
			_719_v46 = _dafny.SeqOf(_639_v4)
			var _720_v47 _dafny.Array
			_ = _720_v47
			var _nwElement0_28 _dafny.Int = _639_v4
			_ = _nwElement0_28
			var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(23))
			_ = _nw107
			(_nw107).ArraySet1(_nwElement0_28, 0)
			(_nw107).ArraySet1(_640_i4, 1)
			(_nw107).ArraySet1((_dafny.Zero).Minus(_640_i4), 2)
			(_nw107).ArraySet1(_640_i4, 3)
			(_nw107).ArraySet1((_dafny.Zero).Minus(_640_i4), 4)
			(_nw107).ArraySet1((_dafny.Zero).Minus(_640_i4), 5)
			(_nw107).ArraySet1(_640_i4, 6)
			(_nw107).ArraySet1(_640_i4, 7)
			(_nw107).ArraySet1(_639_v4, 8)
			(_nw107).ArraySet1(_639_v4, 9)
			(_nw107).ArraySet1(Companion_Default___.Fm5(globalState), 10)
			(_nw107).ArraySet1(_dafny.IntOfUint32((_712_v39).Cardinality()), 11)
			(_nw107).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p0)).Cardinality(), !(p0))).Cardinality(), 12)
			(_nw107).ArraySet1(_639_v4, 13)
			(_nw107).ArraySet1(_639_v4, 14)
			(_nw107).ArraySet1((_713_v40).Cardinality(), 15)
			(_nw107).ArraySet1(((_716_v43).Update(_dafny.IntOfUint32((_dafny.SeqOf((_630_v0).Select((Companion_Default___.SafeIndex(_639_v4, _dafny.IntOfUint32((_630_v0).Cardinality()))).Uint32()).(bool))).Cardinality()), Companion_Default___.Abs(_639_v4))).Cardinality(), 16)
			(_nw107).ArraySet1((_718_v45).Cardinality(), 17)
			(_nw107).ArraySet1(_dafny.IntOfUint32((_719_v46).Cardinality()), 18)
			(_nw107).ArraySet1(_640_i4, 19)
			(_nw107).ArraySet1((_716_v43).Cardinality(), 20)
			(_nw107).ArraySet1(_640_i4, 21)
			(_nw107).ArraySet1(_dafny.IntOfUint32((_630_v0).Cardinality()), 22)
			_720_v47 = _nw107
			var _721_v48 _dafny.Set
			_ = _721_v48
			_721_v48 = _dafny.SetOf(_639_v4)
			var _722_v49 D3
			_ = _722_v49
			_722_v49 = Companion_D3_.Create_DC4_(_639_v4, (_714_v41).F27(), p0)
			var _723_v50 _dafny.Map
			_ = _723_v50
			_723_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_721_v48).Cardinality(), _dafny.SeqOf((_722_v49).Dtor_cf4(), _640_i4))
			var _724_v51 _dafny.Map
			_ = _724_v51
			_724_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_720_v47, (Companion_D12_.Create_DC26_(_723_v50)).Dtor_cf44())
			_724_v51 = (_724_v51).Update(_720_v47, func() _dafny.Map {
				var _coll27 = _dafny.NewMapBuilder()
				_ = _coll27
				for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-908), _dafny.IntOfInt64(148))); ; {
					_compr_27, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _725_v52 _dafny.Int
					_725_v52 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(-908)).Cmp(_725_v52) <= 0) && ((_725_v52).Cmp(_dafny.IntOfInt64(148)) < 0) {
						_coll27.Add(Companion_Default___.SafeModuloInt(_725_v52, _dafny.IntOfInt64(-723)), _719_v46)
					}
				}
				return _coll27.ToMap()
			}())
			var _726_v53 _dafny.Sequence
			_ = _726_v53
			_726_v53 = _dafny.UnicodeSeqOfUtf8Bytes("yq")
			_726_v53 = _726_v53
		}
		(globalState).F8 = (_dafny.Zero).Minus((_639_v4).Times(_639_v4))
		var _727_v54 _dafny.Sequence
		_ = _727_v54
		_727_v54 = _dafny.UnicodeSeqOfUtf8Bytes("uigynbjvj")
		r0 = (_dafny.IntOfInt64(252)).Plus(_dafny.IntOfUint32((_727_v54).Cardinality()))
		var _728_v55 _dafny.Array
		_ = _728_v55
		var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(16))
		_ = _nw108
		_728_v55 = _nw108
		r1 = _728_v55
		var _729_v56 D7
		_ = _729_v56
		_729_v56 = Companion_D7_.Create_DC14_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(962))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg44 _dafny.Int) interface{} {
				return coer44(arg44)
			}
		}((func(_730_v3 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_731_i13 _dafny.Int) _dafny.CodePoint {
				return (_730_v3).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_730_v3), 0))).Int())
			}
		})(_637_v3)))).Cardinality()), false, true, !(p0), _639_v4)
		var _pat_let_tv19 = _639_v4
		_ = _pat_let_tv19
		var _pat_let_tv20 = p0
		_ = _pat_let_tv20
		var _pat_let_tv21 = _637_v3
		_ = _pat_let_tv21
		var _pat_let_tv22 = _637_v3
		_ = _pat_let_tv22
		var _pat_let_tv23 = _639_v4
		_ = _pat_let_tv23
		r2 = func(_source14 D7) bool {
			if _source14.Is_DC14() {
				var _732___mcc_h9 _dafny.Int = _source14.Get_().(D7_DC14).Cf21
				_ = _732___mcc_h9
				var _733___mcc_h10 bool = _source14.Get_().(D7_DC14).Cf22
				_ = _733___mcc_h10
				var _734___mcc_h11 bool = _source14.Get_().(D7_DC14).Cf23
				_ = _734___mcc_h11
				var _735___mcc_h12 bool = _source14.Get_().(D7_DC14).Cf24
				_ = _735___mcc_h12
				var _736___mcc_h13 _dafny.Int = _source14.Get_().(D7_DC14).Cf25
				_ = _736___mcc_h13
				var _737_cf25 _dafny.Int = _736___mcc_h13
				_ = _737_cf25
				var _738_cf24 bool = _735___mcc_h12
				_ = _738_cf24
				var _739_cf23 bool = _734___mcc_h11
				_ = _739_cf23
				var _740_cf22 bool = _733___mcc_h10
				_ = _740_cf22
				var _741_cf21 _dafny.Int = _732___mcc_h9
				_ = _741_cf21
				return _740_cf22
			} else if _source14.Is_DC15() {
				var _742___mcc_h14 _dafny.Int = _source14.Get_().(D7_DC15).Cf26
				_ = _742___mcc_h14
				var _743___mcc_h15 bool = _source14.Get_().(D7_DC15).Cf27
				_ = _743___mcc_h15
				var _744___mcc_h16 _dafny.Int = _source14.Get_().(D7_DC15).Cf28
				_ = _744___mcc_h16
				var _745_cf28 _dafny.Int = _744___mcc_h16
				_ = _745_cf28
				var _746_cf27 bool = _743___mcc_h15
				_ = _746_cf27
				var _747_cf26 _dafny.Int = _742___mcc_h14
				_ = _747_cf26
				return !(false)
			} else {
				var _748___mcc_h17 _dafny.Map = _source14.Get_().(D7_DC13).Cf20
				_ = _748___mcc_h17
				var _749_cf20 _dafny.Map = _748___mcc_h17
				_ = _749_cf20
				return (Companion_D7_.Create_DC15_(_pat_let_tv19, _pat_let_tv20, (func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(688), _dafny.IntOfInt64(852))); ; {
						_compr_28, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _750_v57 _dafny.Int
						_750_v57 = interface{}(_compr_28).(_dafny.Int)
						if ((_dafny.IntOfInt64(688)).Cmp(_750_v57) <= 0) && ((_750_v57).Cmp(_dafny.IntOfInt64(852)) < 0) {
							_coll28.Add(Companion_Default___.SafeModuloInt(_750_v57, _pat_let_tv23), (_dafny.SetOf((_pat_let_tv22).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_pat_let_tv21), 0))).Int()))).Cardinality())
						}
					}
					return _coll28.ToMap()
				}()).Cardinality())).Dtor_cf27()
			}
		}(_729_v56)
		return r0, r1, r2
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	F26 _dafny.Array
}

func New_C3_() *C3 {
	_this := C3{}

	_this.F26 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C3) Ctor__(f26 _dafny.Array) {
	{
		(_this).F26 = f26
	}
}
func (_this *C3) Fm2(p0 _dafny.Int, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg45 _dafny.Int) interface{} {
				return coer45(arg45)
			}
		}(func(_751_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("glht"), _dafny.UnicodeSeqOfUtf8Bytes("wx"), _dafny.UnicodeSeqOfUtf8Bytes("rlclyu"), _dafny.UnicodeSeqOfUtf8Bytes("llfeqpt"), _dafny.UnicodeSeqOfUtf8Bytes("ao"))), _dafny.Companion_Sequence_.Concatenate((Companion_D3_.Create_DC3_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_752_i1 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(95))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_753_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}))
		})))).Dtor_cf3(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(613))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg48 _dafny.Int) interface{} {
				return coer48(arg48)
			}
		}(func(_754_i3 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_755_i4 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('h')
			}))
		}))))
	}
}
func (_this *C3) M1(globalState *GlobalState) {
	{
		var _756_v0 _dafny.Int
		_ = _756_v0
		_756_v0 = _dafny.IntOfInt64(-373)
		if (_756_v0).Cmp(_756_v0) != 0 {
			var _757_v1 bool
			_ = _757_v1
			_757_v1 = true
			(globalState).F17 = _757_v1
			(globalState).F5 = _757_v1
			var _758_v2 *C0
			_ = _758_v2
			var _nw109 *C0 = New_C0_()
			_ = _nw109
			_nw109.Ctor__(_757_v1)
			_758_v2 = _nw109
			var _759_v3 _dafny.Array
			_ = _759_v3
			var _nwElement0_29 bool = false
			_ = _nwElement0_29
			var _nw110 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(16))
			_ = _nw110
			(_nw110).ArraySet1(_nwElement0_29, 0)
			(_nw110).ArraySet1(_757_v1, 1)
			(_nw110).ArraySet1((_758_v2).F27(), 2)
			(_nw110).ArraySet1((_758_v2).F27(), 3)
			(_nw110).ArraySet1(_757_v1, 4)
			(_nw110).ArraySet1((_758_v2).F27(), 5)
			(_nw110).ArraySet1(_757_v1, 6)
			(_nw110).ArraySet1(false, 7)
			(_nw110).ArraySet1((_758_v2).F27(), 8)
			(_nw110).ArraySet1(_757_v1, 9)
			(_nw110).ArraySet1(_757_v1, 10)
			(_nw110).ArraySet1(_757_v1, 11)
			(_nw110).ArraySet1(_757_v1, 12)
			(_nw110).ArraySet1(!((_758_v2).F27()), 13)
			(_nw110).ArraySet1(_757_v1, 14)
			(_nw110).ArraySet1(!((_758_v2).F27()), 15)
			_759_v3 = _nw110
			var _760_v4 _dafny.Set
			_ = _760_v4
			_760_v4 = _dafny.SetOf(_759_v3)
			var _761_v5 _dafny.Sequence
			_ = _761_v5
			_761_v5 = _dafny.SeqOf(_760_v4)
			var _762_v6 _dafny.Sequence
			_ = _762_v6
			_762_v6 = _dafny.SeqOf(_756_v0, _756_v0)
			var _arr7 _dafny.Array = _this.F26
			_ = _arr7
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))
			_ = _index93
			(_arr7).ArraySet1(((_761_v5).Select((Companion_Default___.SafeIndex((_762_v6).Select((Companion_Default___.SafeIndex(_756_v0, _dafny.IntOfUint32((_762_v6).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_761_v5).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), (_index93).Int())
			var _763_v7 _dafny.Map
			_ = _763_v7
			_763_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_756_v0, (Companion_D4_.Create_DC7_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F26, _this.F26)).Cardinality()), _758_v2, !(false), _756_v0)).Dtor_cf11())
			var _764_v8 _dafny.Sequence
			_ = _764_v8
			_764_v8 = _dafny.SeqOf(_757_v1, _757_v1, (_758_v2).F27())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_759_v3), 0))
			_ = _index94
			(_759_v3).ArraySet1((func() bool {
				if (_763_v7).Contains(_dafny.IntOfUint32((_764_v8).Cardinality())) {
					return (_763_v7).Get(_dafny.IntOfUint32((_764_v8).Cardinality())).(bool)
				}
				return (_758_v2).F27()
			})(), (_index94).Int())
			var _765_v9 _dafny.Sequence
			_ = _765_v9
			_765_v9 = _dafny.UnicodeSeqOfUtf8Bytes("bycndgxtw")
			var _766_v10 _dafny.MultiSet
			_ = _766_v10
			_766_v10 = _dafny.MultiSetOf(_dafny.IntOfInt64(348), _dafny.IntOfUint32((_765_v9).Cardinality()))
			var _767_v11 _dafny.Set
			_ = _767_v11
			_767_v11 = _dafny.SetOf(_757_v1, (_764_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(326), _dafny.IntOfUint32((_764_v8).Cardinality()))).Uint32()).(bool))
			var _arr8 _dafny.Array = _this.F26
			_ = _arr8
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))
			_ = _index95
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_759_v3), 0))
			_ = _index96
			var _rhs88 _dafny.Int = _756_v0
			_ = _rhs88
			var _rhs89 bool = _757_v1
			_ = _rhs89
			var _rhs90 _dafny.Int = Companion_Default___.SafeDivisionInt(_756_v0, _756_v0)
			_ = _rhs90
			var _rhs91 _dafny.Map = Companion_Default___.Fm28((_756_v0).Cmp((_dafny.Zero).Minus((_766_v10).Cardinality())) < 0, (_767_v11).IsDisjointFrom(_767_v11), (_758_v2).F27(), _756_v0, globalState)
			_ = _rhs91
			var _lhs61 _dafny.Array = _this.F26
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))
			_ = _lhs62
			var _lhs63 _dafny.Array = _759_v3
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_759_v3), 0))
			_ = _lhs64
			var _lhs65 *GlobalState = globalState
			_ = _lhs65
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			(_lhs61).ArraySet1(_rhs88, (_lhs62).Int())
			(_lhs63).ArraySet1(_rhs89, (_lhs64).Int())
			_lhs65.F8 = _rhs90
			_lhs66.F2 = _rhs91
			var _768_v12 _dafny.Array
			_ = _768_v12
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(17))
			_ = _nw111
			_768_v12 = _nw111
			var _769_v13 _dafny.MultiSet
			_ = _769_v13
			_769_v13 = _dafny.MultiSetOf(false, _757_v1)
			var _770_v14 _dafny.CodePoint
			_ = _770_v14
			_770_v14 = _dafny.CodePoint('w')
			var _771_v15 D8
			_ = _771_v15
			_771_v15 = Companion_D8_.Create_DC17_(_770_v14, _770_v14, _dafny.IntOfInt64(891))
			var _772_v16 _dafny.Array
			_ = _772_v16
			var _nwElement0_30 _dafny.Int = _dafny.IntOfInt64(536)
			_ = _nwElement0_30
			var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(16))
			_ = _nw112
			(_nw112).ArraySet1(_nwElement0_30, 0)
			(_nw112).ArraySet1((_this.F26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))).Int()).(_dafny.Int), 1)
			(_nw112).ArraySet1(_dafny.IntOfInt64(195), 2)
			(_nw112).ArraySet1((_this.F26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))).Int()).(_dafny.Int), 3)
			(_nw112).ArraySet1(_756_v0, 4)
			(_nw112).ArraySet1(_dafny.IntOfInt64(740), 5)
			(_nw112).ArraySet1(_756_v0, 6)
			(_nw112).ArraySet1(Companion_Default___.Fm18(_769_v13, _771_v15, (_758_v2).F27(), globalState), 7)
			(_nw112).ArraySet1((_this.F26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_this.F26), 0))).Int()).(_dafny.Int), 8)
			(_nw112).ArraySet1(_756_v0, 9)
			(_nw112).ArraySet1(_756_v0, 10)
			(_nw112).ArraySet1((_769_v13).Cardinality(), 11)
			(_nw112).ArraySet1(_756_v0, 12)
			(_nw112).ArraySet1(_756_v0, 13)
			(_nw112).ArraySet1(_756_v0, 14)
			(_nw112).ArraySet1(_756_v0, 15)
			_772_v16 = _nw112
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_768_v12), 0))
			_ = _index97
			(_768_v12).ArraySet1(Companion_D6_.Create_DC10_(_772_v16), (_index97).Int())
			var _773_v17 D6
			_ = _773_v17
			_773_v17 = Companion_D6_.Create_DC10_((func() _dafny.Array {
				if (_759_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_759_v3), 0))).Int()).(bool) {
					return _772_v16
				}
				return _772_v16
			})())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(313), _dafny.ArrayLen((_768_v12), 0))
			_ = _index98
			(_768_v12).ArraySet1(_773_v17, (_index98).Int())
		} else {
			var _774_v18 *C2
			_ = _774_v18
			var _nw113 *C2 = New_C2_()
			_ = _nw113
			_nw113.Ctor__()
			_774_v18 = _nw113
			var _775_v19 _dafny.Sequence
			_ = _775_v19
			_775_v19 = _dafny.UnicodeSeqOfUtf8Bytes("qujdwm")
			_775_v19 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-237))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}(func(_776_i0 _dafny.Int) _dafny.CodePoint {
				return (func() _dafny.CodePoint {
					if false {
						return _dafny.CodePoint('d')
					}
					return _dafny.CodePoint('g')
				})()
			}))
			(globalState).F8 = _756_v0
			var _777_v20 bool
			_ = _777_v20
			_777_v20 = false
			var _778_v21 *C0
			_ = _778_v21
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__(_777_v20)
			_778_v21 = _nw114
			var _779_v23 _dafny.CodePoint
			_ = _779_v23
			_779_v23 = _dafny.CodePoint('n')
			var _780_v24 _dafny.MultiSet
			_ = _780_v24
			_780_v24 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("qd"), (Companion_Default___.SafeIndex(_756_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qd")).Cardinality()))).Uint32(), _779_v23), _dafny.UnicodeSeqOfUtf8Bytes("sir"))
			var _781_v25 _dafny.Map
			_ = _781_v25
			_781_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((Companion_Default___.Fm7(globalState)).Cardinality()), (Companion_D4_.Create_DC7_(_756_v0, _778_v21, _777_v20, (func() _dafny.Map {
				var _coll29 = _dafny.NewMapBuilder()
				_ = _coll29
				for _iter32 := _dafny.Iterate((_780_v24).Elements()); ; {
					_compr_29, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _782_v22 _dafny.Sequence
					_782_v22 = interface{}(_compr_29).(_dafny.Sequence)
					if (_780_v24).Contains(_782_v22) {
						_coll29.Add(_782_v22, false)
					}
				}
				return _coll29.ToMap()
			}()).Cardinality())).Dtor_cf11())
			_781_v25 = (_781_v25).Update(_756_v0, _777_v20)
			var _783_v26 _dafny.Sequence
			_ = _783_v26
			_783_v26 = _dafny.SeqOf(_777_v20, _777_v20, (func() bool {
				if (_781_v25).Contains(_756_v0) {
					return (_781_v25).Get(_756_v0).(bool)
				}
				return false
			})())
			(globalState).F17 = Companion_Default___.Fm0(_783_v26, globalState)
		}
		var _784_v27 bool
		_ = _784_v27
		_784_v27 = true
		var _785_v28 _dafny.Map
		_ = _785_v28
		_785_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if _784_v27 {
				return _756_v0
			}
			return _dafny.IntOfUint32((Companion_Default___.Fm7(globalState)).Cardinality())
		})(), Companion_Default___.Fm5(globalState))
		var _786_v29 _dafny.Sequence
		_ = _786_v29
		_786_v29 = _dafny.SeqOf(_784_v27, _784_v27)
		var _787_v30 _dafny.Sequence
		_ = _787_v30
		_787_v30 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-615))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}((func(_788_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_789_i1 _dafny.Int) _dafny.Int {
				return _788_v0
			}
		})(_756_v0)))))
		var _790_v31 _dafny.Sequence
		_ = _790_v31
		_790_v31 = _dafny.SeqOf(_dafny.IntOfUint32((_786_v29).Cardinality()), ((_787_v30).Select((Companion_Default___.SafeIndex(_756_v0, _dafny.IntOfUint32((_787_v30).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), Companion_Default___.Fm5(globalState), _dafny.IntOfInt64(173), _756_v0)
		var _791_v32 _dafny.Map
		_ = _791_v32
		_791_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_790_v31, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_756_v0, _784_v27))
		var _792_v33 _dafny.Map
		_ = _792_v33
		_792_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_756_v0, _784_v27)
		var _793_v35 _dafny.MultiSet
		_ = _793_v35
		_793_v35 = _dafny.MultiSetOf((func() _dafny.Map {
			if (_791_v32).Contains(_790_v31) {
				return (_791_v32).Get(_790_v31).(_dafny.Map)
			}
			return _792_v33
		})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-59), _784_v27), func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter33 := _dafny.Iterate((_dafny.SetOf(_756_v0)).Elements()); ; {
				_compr_30, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _794_v34 _dafny.Int
				_794_v34 = interface{}(_compr_30).(_dafny.Int)
				if (_dafny.SetOf(_756_v0)).Contains(_794_v34) {
					_coll30.Add(Companion_Default___.SafeModuloInt(_794_v34, _756_v0), _784_v27)
				}
			}
			return _coll30.ToMap()
		}())
		_785_v28 = (_785_v28).Update((func() _dafny.Int {
			if (_793_v35).Contains(_792_v33) {
				return (_793_v35).Multiplicity(_792_v33)
			}
			return _756_v0
		})(), _756_v0)
		(globalState).F8 = (_dafny.Zero).Minus(((_dafny.Zero).Minus((func() _dafny.Int {
			if _784_v27 {
				return _756_v0
			}
			return _756_v0
		})())).Minus((Companion_D10_.Create_DC22_(_756_v0, _784_v27)).Dtor_cf38()))
		var _795_v36 *C2
		_ = _795_v36
		var _nw115 *C2 = New_C2_()
		_ = _nw115
		_nw115.Ctor__()
		_795_v36 = _nw115
		var _hi5 _dafny.Int = _756_v0
		_ = _hi5
		for _796_i2 := _756_v0; _796_i2.Cmp(_hi5) < 0; _796_i2 = _796_i2.Plus(_dafny.One) {
			var _797_v37 _dafny.Sequence
			_ = _797_v37
			_797_v37 = _dafny.UnicodeSeqOfUtf8Bytes("kxl")
			var _798_v38 _dafny.Map
			_ = _798_v38
			_798_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_785_v28, _797_v37)
			_792_v33 = (_792_v33).Update((func() _dafny.Int {
				if _784_v27 {
					return (_798_v38).Cardinality()
				}
				return _796_i2
			})(), false)
			var _799_v39 _dafny.Int
			_ = _799_v39
			var _800_v40 bool
			_ = _800_v40
			var _801_v41 bool
			_ = _801_v41
			var _802_v42 _dafny.Int
			_ = _802_v42
			var _out27 _dafny.Int
			_ = _out27
			var _out28 bool
			_ = _out28
			var _out29 bool
			_ = _out29
			var _out30 _dafny.Int
			_ = _out30
			_out27, _out28, _out29, _out30 = Companion_Default___.M0(globalState)
			_799_v39 = _out27
			_800_v40 = _out28
			_801_v41 = _out29
			_802_v42 = _out30
			var _803_v43 _dafny.Array
			_ = _803_v43
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_13
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = func(_804_i3 _dafny.Int) bool {
					return false
				}
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw116 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw116).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw116).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_803_v43 = _nw116
			_803_v43 = _803_v43
			var _805_v44 *C0
			_ = _805_v44
			var _nw117 *C0 = New_C0_()
			_ = _nw117
			_nw117.Ctor__(Companion_Default___.Fm0(_786_v29, globalState))
			_805_v44 = _nw117
		}
		for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_this.F26), 0))); ; {
			_guard_loop_3, _ok34 := _iter34()
			if !_ok34 {
				break
			}
			var _806_i4 _dafny.Int
			_806_i4 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_806_i4).Sign() != -1) && ((_806_i4).Cmp(_dafny.ArrayLen((_this.F26), 0)) < 0)) {
				var _arr9 _dafny.Array = _this.F26
				_ = _arr9
				(_arr9).ArraySet1(Companion_Default___.SafeModuloInt(_806_i4, _756_v0), (_806_i4).Int())
			}
		}
	}
}
func (_this *C3) M2(globalState *GlobalState) {
	{
		var _807_v0 _dafny.Array
		_ = _807_v0
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_14
		var _nw118 _dafny.Array
		_ = _nw118
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw118 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) bool = func(_808_i0 _dafny.Int) bool {
				return false
			}
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw118 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw118).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw118).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_807_v0 = _nw118
		var _809_v1 bool
		_ = _809_v1
		_809_v1 = false
		var _810_v2 _dafny.Array
		_ = _810_v2
		var _nwElement0_31 _dafny.Array = _807_v0
		_ = _nwElement0_31
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(28))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_31, 0)
		(_nw119).ArraySet1(_807_v0, 1)
		(_nw119).ArraySet1(_807_v0, 2)
		(_nw119).ArraySet1(_807_v0, 3)
		(_nw119).ArraySet1(_807_v0, 4)
		(_nw119).ArraySet1(_807_v0, 5)
		(_nw119).ArraySet1(_807_v0, 6)
		(_nw119).ArraySet1((func() _dafny.Array {
			if _809_v1 {
				return _807_v0
			}
			return _807_v0
		})(), 7)
		(_nw119).ArraySet1(_807_v0, 8)
		(_nw119).ArraySet1(_807_v0, 9)
		(_nw119).ArraySet1(_807_v0, 10)
		(_nw119).ArraySet1(_807_v0, 11)
		(_nw119).ArraySet1(_807_v0, 12)
		(_nw119).ArraySet1(_807_v0, 13)
		(_nw119).ArraySet1(_807_v0, 14)
		(_nw119).ArraySet1(_807_v0, 15)
		(_nw119).ArraySet1(_807_v0, 16)
		(_nw119).ArraySet1(_807_v0, 17)
		(_nw119).ArraySet1(_807_v0, 18)
		(_nw119).ArraySet1(_807_v0, 19)
		(_nw119).ArraySet1(_807_v0, 20)
		(_nw119).ArraySet1((func() _dafny.Array {
			if false {
				return _807_v0
			}
			return _807_v0
		})(), 21)
		(_nw119).ArraySet1(_807_v0, 22)
		(_nw119).ArraySet1(_807_v0, 23)
		(_nw119).ArraySet1(_807_v0, 24)
		(_nw119).ArraySet1(_807_v0, 25)
		(_nw119).ArraySet1(_807_v0, 26)
		(_nw119).ArraySet1(_807_v0, 27)
		_810_v2 = _nw119
		_810_v2 = _810_v2
		(globalState).F5 = (func() bool {
			if _809_v1 {
				return _809_v1
			}
			return (_809_v1) && (_809_v1)
		})()
		var _811_v3 _dafny.Int
		_ = _811_v3
		_811_v3 = _dafny.IntOfInt64(487)
		var _812_v4 _dafny.Sequence
		_ = _812_v4
		_812_v4 = _dafny.SeqOf(_811_v3, _811_v3)
		var _813_v5 _dafny.MultiSet
		_ = _813_v5
		_813_v5 = _dafny.MultiSetOf(_811_v3, _dafny.IntOfUint32((_dafny.SeqOf(_809_v1)).Cardinality()), (_812_v4).Select((Companion_Default___.SafeIndex(_811_v3, _dafny.IntOfUint32((_812_v4).Cardinality()))).Uint32()).(_dafny.Int))
		var _814_v6 _dafny.CodePoint
		_ = _814_v6
		_814_v6 = _dafny.CodePoint('e')
		var _815_v7 _dafny.Map
		_ = _815_v7
		_815_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_809_v1, _813_v5)
		var _816_v8 _dafny.Map
		_ = _816_v8
		_816_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_811_v3, _814_v6)
		var _817_v9 _dafny.MultiSet
		_ = _817_v9
		_817_v9 = _dafny.MultiSetOf(true)
		var _818_v10 D8
		_ = _818_v10
		_818_v10 = Companion_D8_.Create_DC17_(_814_v6, _814_v6, _811_v3)
		var _819_v11 T1
		_ = _819_v11
		var _nw120 *C0 = New_C0_()
		_ = _nw120
		_nw120.Ctor__(!(_809_v1))
		_819_v11 = _nw120
		var _820_v12 _dafny.Map
		_ = _820_v12
		_820_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_811_v3, _819_v11)
		var _821_v13 _dafny.Map
		_ = _821_v13
		_821_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T1 {
			if (_820_v12).Contains(_dafny.IntOfInt64(-651)) {
				return (_820_v12).Get(_dafny.IntOfInt64(-651)).(T1)
			}
			return _819_v11
		})(), _811_v3)
		var _pat_let_tv24 = _811_v3
		_ = _pat_let_tv24
		var _822_v14 _dafny.Sequence
		_ = _822_v14
		_822_v14 = _dafny.SeqOf((_813_v5).Intersection(_813_v5), _dafny.MultiSetOf(_811_v3, _811_v3, _dafny.IntOfInt64(487), (Companion_Default___.Fm21(_814_v6, _811_v3, globalState)).Cardinality()), (func() _dafny.MultiSet {
			if (_815_v7).Contains(_809_v1) {
				return (_815_v7).Get(_809_v1).(_dafny.MultiSet)
			}
			return Companion_Default___.Fm6(_812_v4, (_816_v8).Cardinality(), _811_v3, globalState)
		})(), _dafny.MultiSetOf(_811_v3), Companion_Default___.Fm6(_dafny.SeqOf(Companion_Default___.Fm18(_817_v9, func(_pat_let8_0 D8) D8 {
			return func(_823_dt__update__tmp_h0 D8) D8 {
				return func(_pat_let9_0 _dafny.Int) D8 {
					return func(_824_dt__update_hcf32_h0 _dafny.Int) D8 {
						return Companion_D8_.Create_DC17_((_823_dt__update__tmp_h0).Dtor_cf30(), (_823_dt__update__tmp_h0).Dtor_cf31(), _824_dt__update_hcf32_h0)
					}(_pat_let9_0)
				}(_pat_let_tv24)
			}(_pat_let8_0)
		}(_818_v10), _809_v1, globalState)), (func() _dafny.Int {
			if (_821_v13).Contains(_819_v11) {
				return (_821_v13).Get(_819_v11).(_dafny.Int)
			}
			return _811_v3
		})(), _811_v3, globalState))
		_822_v14 = _822_v14
		var _825_i1 _dafny.Int
		_ = _825_i1
		_825_i1 = _dafny.Zero
		{
			for _809_v1 {
				{
					if (_825_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_825_i1 = (_825_i1).Plus(_dafny.One)
					var _826_v15 _dafny.Sequence
					_ = _826_v15
					_826_v15 = _dafny.SeqOf(_this.F26, _this.F26, _this.F26, _this.F26, _this.F26)
					var _arr10 _dafny.Array = _this.F26
					_ = _arr10
					var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_this.F26), 0))
					_ = _index99
					(_arr10).ArraySet1(Companion_Default___.SafeModuloInt(_811_v3, _811_v3), (_index99).Int())
					var _arr11 _dafny.Array = _this.F26
					_ = _arr11
					var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_this.F26), 0))
					_ = _index100
					var _rhs92 _dafny.Sequence = _826_v15
					_ = _rhs92
					var _rhs93 _dafny.Int = _811_v3
					_ = _rhs93
					var _lhs67 _dafny.Array = _this.F26
					_ = _lhs67
					var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_this.F26), 0))
					_ = _lhs68
					_826_v15 = _rhs92
					(_lhs67).ArraySet1(_rhs93, (_lhs68).Int())
					var _827_v16 _dafny.Set
					_ = _827_v16
					_827_v16 = _dafny.SetOf((_this.F26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(751), _dafny.ArrayLen((_this.F26), 0))).Int()).(_dafny.Int), _811_v3)
					var _828_v17 _dafny.Sequence
					_ = _828_v17
					_828_v17 = _dafny.UnicodeSeqOfUtf8Bytes("jqlnl")
					var _829_v18 _dafny.Map
					_ = _829_v18
					_829_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_827_v16, _828_v17)
					var _830_v19 *C1
					_ = _830_v19
					var _nw121 *C1 = New_C1_()
					_ = _nw121
					_nw121.Ctor__(_829_v18, _819_v11)
					_830_v19 = _nw121
					_807_v0 = _807_v0
					_827_v16 = (_827_v16).Difference(_827_v16)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _831_v20 _dafny.Array
		_ = _831_v20
		var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(9))
		_ = _nw122
		_831_v20 = _nw122
		_831_v20 = _831_v20
		(globalState).F8 = _811_v3
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f25 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f25 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C4{}
var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__(f25 _dafny.Int) {
	{
		(_this)._f25 = f25
	}
}
func (_this *C4) M1(globalState *GlobalState) {
	{
		var _832_v0 _dafny.Sequence
		_ = _832_v0
		_832_v0 = _dafny.SeqOf((_this).F25(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(241))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg52 _dafny.Int) interface{} {
				return coer52(arg52)
			}
		}(func(_833_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))).Cardinality()), (_this).F25(), (_this).F25(), (_this).F25())
		var _834_v1 _dafny.Int
		_ = _834_v1
		_834_v1 = _dafny.IntOfUint32((_832_v0).Cardinality())
		var _hi6 _dafny.Int = (_this).F25()
		_ = _hi6
		for _835_i0 := (_834_v1); _835_i0.Cmp(_hi6) < 0; _835_i0 = _835_i0.Plus(_dafny.One) {
			var _836_v2 bool
			_ = _836_v2
			_836_v2 = true
			var _rhs94 bool = _836_v2
			_ = _rhs94
			var _rhs95 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-854), _835_i0)
			_ = _rhs95
			var _lhs69 *GlobalState = globalState
			_ = _lhs69
			var _lhs70 *GlobalState = globalState
			_ = _lhs70
			_lhs69.F17 = _rhs94
			_lhs70.F8 = _rhs95
			var _837_v3 _dafny.Array
			_ = _837_v3
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw123
			_837_v3 = _nw123
			var _838_v4 _dafny.Sequence
			_ = _838_v4
			_838_v4 = _dafny.SeqOf(_837_v3, _837_v3)
			var _839_v5 _dafny.Array
			_ = _839_v5
			var _nwElement0_32 _dafny.Array = (func() _dafny.Array {
				if _836_v2 {
					return _837_v3
				}
				return _837_v3
			})()
			_ = _nwElement0_32
			var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(18))
			_ = _nw124
			(_nw124).ArraySet1(_nwElement0_32, 0)
			(_nw124).ArraySet1(_837_v3, 1)
			(_nw124).ArraySet1(_837_v3, 2)
			(_nw124).ArraySet1(_837_v3, 3)
			(_nw124).ArraySet1((func() _dafny.Array {
				if _836_v2 {
					return _837_v3
				}
				return _837_v3
			})(), 4)
			(_nw124).ArraySet1((func() _dafny.Array {
				if _836_v2 {
					return _837_v3
				}
				return _837_v3
			})(), 5)
			(_nw124).ArraySet1(_837_v3, 6)
			(_nw124).ArraySet1(_837_v3, 7)
			(_nw124).ArraySet1(_837_v3, 8)
			(_nw124).ArraySet1(_837_v3, 9)
			(_nw124).ArraySet1(_837_v3, 10)
			(_nw124).ArraySet1(_837_v3, 11)
			(_nw124).ArraySet1((func() _dafny.Array {
				if _836_v2 {
					return (_838_v4).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_838_v4).Cardinality()))).Uint32()).(_dafny.Array)
				}
				return _837_v3
			})(), 12)
			(_nw124).ArraySet1(_837_v3, 13)
			(_nw124).ArraySet1(_837_v3, 14)
			(_nw124).ArraySet1(_837_v3, 15)
			(_nw124).ArraySet1(_837_v3, 16)
			(_nw124).ArraySet1(_837_v3, 17)
			_839_v5 = _nw124
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_839_v5), 0))
			_ = _index101
			(_839_v5).ArraySet1(_837_v3, (_index101).Int())
			var _840_v6 _dafny.Sequence
			_ = _840_v6
			_840_v6 = _dafny.UnicodeSeqOfUtf8Bytes("upnyjowbd")
			var _841_v7 _dafny.CodePoint
			_ = _841_v7
			_841_v7 = _dafny.CodePoint('t')
			var _842_v8 _dafny.Array
			_ = _842_v8
			var _nwElement0_33 _dafny.Sequence = Companion_Default___.Fm1(globalState)
			_ = _nwElement0_33
			var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(9))
			_ = _nw125
			(_nw125).ArraySet1(_nwElement0_33, 0)
			(_nw125).ArraySet1(_840_v6, 1)
			(_nw125).ArraySet1(_dafny.Companion_Sequence_.Update(_840_v6, (Companion_Default___.SafeIndex(_835_i0, _dafny.IntOfUint32((_840_v6).Cardinality()))).Uint32(), _841_v7), 2)
			(_nw125).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_840_v6, _840_v6), 3)
			(_nw125).ArraySet1(_840_v6, 4)
			(_nw125).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-241))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_843_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_844_i2 _dafny.Int) _dafny.CodePoint {
					return _843_v7
				}
			})(_841_v7))), 5)
			(_nw125).ArraySet1(_840_v6, 6)
			(_nw125).ArraySet1(_840_v6, 7)
			(_nw125).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_840_v6, _840_v6), 8)
			_842_v8 = _nw125
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_842_v8), 0))
			_ = _index102
			(_842_v8).ArraySet1(_840_v6, (_index102).Int())
			var _845_v9 _dafny.Sequence
			_ = _845_v9
			_845_v9 = _dafny.SeqOf(!(_836_v2))
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_839_v5), 0))
			_ = _index103
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_842_v8), 0))
			_ = _index104
			var _rhs96 _dafny.Array = _837_v3
			_ = _rhs96
			var _rhs97 _dafny.Sequence = _840_v6
			_ = _rhs97
			var _rhs98 bool = Companion_Default___.Fm0(_845_v9, globalState)
			_ = _rhs98
			var _lhs71 _dafny.Array = _839_v5
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(832), _dafny.ArrayLen((_839_v5), 0))
			_ = _lhs72
			var _lhs73 _dafny.Array = _842_v8
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(57), _dafny.ArrayLen((_842_v8), 0))
			_ = _lhs74
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			(_lhs71).ArraySet1(_rhs96, (_lhs72).Int())
			(_lhs73).ArraySet1(_rhs97, (_lhs74).Int())
			_lhs75.F22 = _rhs98
			_836_v2 = _836_v2
			var _846_v10 _dafny.Map
			_ = _846_v10
			_846_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _835_i0)
			_846_v10 = (_846_v10).Update(_dafny.IntOfInt64(66), _dafny.IntOfUint32((_832_v0).Cardinality()))
		}
		var _847_v11 bool
		_ = _847_v11
		_847_v11 = false
		var _848_v12 _dafny.Set
		_ = _848_v12
		_848_v12 = _dafny.SetOf(true, false, _847_v11)
		var _849_v13 _dafny.Map
		_ = _849_v13
		_849_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), _847_v11)
		var _850_v14 _dafny.Sequence
		_ = _850_v14
		_850_v14 = _dafny.SeqOf(_847_v11, _847_v11)
		var _851_v15 _dafny.Array
		_ = _851_v15
		var _nwElement0_34 bool = _847_v11
		_ = _nwElement0_34
		var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(17))
		_ = _nw126
		(_nw126).ArraySet1(_nwElement0_34, 0)
		(_nw126).ArraySet1(_847_v11, 1)
		(_nw126).ArraySet1(_847_v11, 2)
		(_nw126).ArraySet1(_847_v11, 3)
		(_nw126).ArraySet1(_847_v11, 4)
		(_nw126).ArraySet1(true, 5)
		(_nw126).ArraySet1(_847_v11, 6)
		(_nw126).ArraySet1((func() bool {
			if (_849_v13).Contains((_dafny.Zero).Minus((_this).F25())) {
				return (_849_v13).Get((_dafny.Zero).Minus((_this).F25())).(bool)
			}
			return _847_v11
		})(), 7)
		(_nw126).ArraySet1(_847_v11, 8)
		(_nw126).ArraySet1(_847_v11, 9)
		(_nw126).ArraySet1(Companion_Default___.Fm0(_850_v14, globalState), 10)
		(_nw126).ArraySet1(_847_v11, 11)
		(_nw126).ArraySet1(_847_v11, 12)
		(_nw126).ArraySet1(_847_v11, 13)
		(_nw126).ArraySet1(false, 14)
		(_nw126).ArraySet1(false, 15)
		(_nw126).ArraySet1(_847_v11, 16)
		_851_v15 = _nw126
		var _852_v16 _dafny.Array
		_ = _852_v16
		_852_v16 = _851_v15
		var _853_v17 _dafny.Sequence
		_ = _853_v17
		_853_v17 = _dafny.UnicodeSeqOfUtf8Bytes("na")
		var _854_v18 bool
		_ = _854_v18
		var _855_v19 bool
		_ = _855_v19
		var _856_v20 bool
		_ = _856_v20
		var _out31 bool
		_ = _out31
		var _out32 bool
		_ = _out32
		var _out33 bool
		_ = _out33
		_out31, _out32, _out33 = (_this).M3(_848_v12, _847_v11, (_851_v15), _dafny.Companion_Sequence_.Concatenate(_853_v17, _853_v17), globalState)
		_854_v18 = _out31
		_855_v19 = _out32
		_856_v20 = _out33
		var _857_v21 _dafny.MultiSet
		_ = _857_v21
		_857_v21 = _dafny.MultiSetOf((_this).F25())
		var _858_v22 _dafny.Map
		_ = _858_v22
		_858_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_857_v21).Difference(_857_v21))
		_858_v22 = (_858_v22).Merge(_858_v22)
		(globalState).F19 = _847_v11
		var _859_i3 _dafny.Int
		_ = _859_i3
		_859_i3 = _dafny.Zero
		{
			for _854_v18 {
				{
					if (_859_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_859_i3 = (_859_i3).Plus(_dafny.One)
					var _860_v24 _dafny.Map
					_ = _860_v24
					_860_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
						var _coll31 = _dafny.NewBuilder()
						_ = _coll31
						for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(138), _dafny.IntOfInt64(58))); ; {
							_compr_31, _ok35 := _iter35()
							if !_ok35 {
								break
							}
							var _861_v23 _dafny.Int
							_861_v23 = interface{}(_compr_31).(_dafny.Int)
							if ((_dafny.IntOfInt64(138)).Cmp(_861_v23) <= 0) && ((_861_v23).Cmp(_dafny.IntOfInt64(58)) < 0) {
								_coll31.Add((_861_v23).Times(_dafny.IntOfUint32((_850_v14).Cardinality())))
							}
						}
						return _coll31.ToSet()
					}(), _dafny.UnicodeSeqOfUtf8Bytes("nttpwyv"))
					var _862_v25 T1
					_ = _862_v25
					var _nw127 *C0 = New_C0_()
					_ = _nw127
					_nw127.Ctor__(_855_v19)
					_862_v25 = _nw127
					var _863_v26 T0
					_ = _863_v26
					var _nw128 *C1 = New_C1_()
					_ = _nw128
					_nw128.Ctor__((_860_v24).Merge(_860_v24), _862_v25)
					_863_v26 = _nw128
					_863_v26 = _863_v26
					var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
					_ = _index105
					(_851_v15).ArraySet1(true, (_index105).Int())
					var _864_v27 _dafny.CodePoint
					_ = _864_v27
					_864_v27 = _dafny.CodePoint('v')
					var _865_v28 _dafny.Array
					_ = _865_v28
					var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
					_ = _nw129
					_865_v28 = _nw129
					var _866_v29 _dafny.MultiSet
					_ = _866_v29
					_866_v29 = _dafny.MultiSetOf(_865_v28)
					var _867_v30 D8
					_ = _867_v30
					_867_v30 = Companion_D8_.Create_DC17_(_864_v27, Companion_Default___.Fm22(false, globalState), ((_this).F25()).Minus((_866_v29).Cardinality()))
					var _868_v31 *C2
					_ = _868_v31
					var _nw130 *C2 = New_C2_()
					_ = _nw130
					_nw130.Ctor__()
					_868_v31 = _nw130
					var _869_v32 _dafny.MultiSet
					_ = _869_v32
					_869_v32 = _dafny.MultiSetOf(_868_v31, _868_v31, _868_v31)
					var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
					_ = _index106
					var _rhs99 bool = true
					_ = _rhs99
					var _rhs100 bool = Companion_Default___.Fm0(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm12(_856_v20, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(939))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}(func(_870_i4 _dafny.Int) _dafny.Int {
						return (_this).F25()
					}))).Cardinality())), !(_854_v18), globalState), _850_v14), globalState)
					_ = _rhs100
					var _rhs101 D8 = (func() D8 {
						if (_869_v32).IsProperSubsetOf(_869_v32) {
							return _867_v30
						}
						return _867_v30
					})()
					_ = _rhs101
					var _rhs102 bool = _847_v11
					_ = _rhs102
					var _lhs76 _dafny.Array = _851_v15
					_ = _lhs76
					var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
					_ = _lhs77
					var _lhs78 *GlobalState = globalState
					_ = _lhs78
					(_lhs76).ArraySet1(_rhs99, (_lhs77).Int())
					_lhs78.F17 = _rhs100
					_867_v30 = _rhs101
					_856_v20 = _rhs102
					if _854_v18 {
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _index107
						(_865_v28).ArraySet1((_this).F25(), (_index107).Int())
						var _871_v33 _dafny.Sequence
						_ = _871_v33
						_871_v33 = _dafny.SeqOf(_832_v0)
						var _872_v34 _dafny.MultiSet
						_ = _872_v34
						_872_v34 = _dafny.MultiSetOf(_dafny.SeqOf((_this).F25()))
						var _873_v35 _dafny.MultiSet
						_ = _873_v35
						_873_v35 = _dafny.MultiSetOf((_dafny.SetOf(true, true)).IsSubsetOf(_848_v12), _847_v11, _855_v19, !(_872_v34).Contains((Companion_D13_.Create_DC28_((_871_v33).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_850_v14).Cardinality()), _dafny.IntOfUint32((_871_v33).Cardinality()))).Uint32()).(_dafny.Sequence))).Dtor_cf47()))
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
						_ = _index108
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _index109
						var _rhs103 _dafny.Int = _dafny.IntOfUint32((_853_v17).Cardinality())
						_ = _rhs103
						var _rhs104 bool = _854_v18
						_ = _rhs104
						var _rhs105 _dafny.Int = (func() _dafny.Int {
							if (_873_v35).Contains(false) {
								return (_873_v35).Multiplicity(false)
							}
							return Companion_Default___.Fm18(_873_v35, _867_v30, _856_v20, globalState)
						})()
						_ = _rhs105
						var _rhs106 _dafny.Int = ((_this).F25()).Times((_this).F25())
						_ = _rhs106
						var _lhs79 *GlobalState = globalState
						_ = _lhs79
						var _lhs80 _dafny.Array = _851_v15
						_ = _lhs80
						var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
						_ = _lhs81
						var _lhs82 _dafny.Array = _865_v28
						_ = _lhs82
						var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _lhs83
						var _lhs84 *GlobalState = globalState
						_ = _lhs84
						_lhs79.F8 = _rhs103
						(_lhs80).ArraySet1(_rhs104, (_lhs81).Int())
						(_lhs82).ArraySet1(_rhs105, (_lhs83).Int())
						_lhs84.F8 = _rhs106
						var _874_v36 _dafny.Array
						_ = _874_v36
						var _len0_15 _dafny.Int = _dafny.IntOfInt64(26)
						_ = _len0_15
						var _nw131 _dafny.Array
						_ = _nw131
						if _len0_15.Cmp(_dafny.Zero) == 0 {
							_nw131 = _dafny.NewArray(_len0_15)
						} else {
							var _init15 func(_dafny.Int) _dafny.Int = func(_875_i5 _dafny.Int) _dafny.Int {
								return (_875_i5).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("trutx")).Cardinality()))
							}
							_ = _init15
							var _element0_15 = _init15(_dafny.Zero)
							_ = _element0_15
							_nw131 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
							(_nw131).ArraySet1(_element0_15, 0)
							var _nativeLen0_15 = (_len0_15).Int()
							_ = _nativeLen0_15
							for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
								(_nw131).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
							}
						}
						_874_v36 = _nw131
						var _876_v37 _dafny.Set
						_ = _876_v37
						_876_v37 = _dafny.SetOf(_874_v36, _874_v36, _874_v36)
						var _877_v38 _dafny.Set
						_ = _877_v38
						_877_v38 = _876_v37
						var _878_v39 _dafny.Map
						_ = _878_v39
						_878_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _876_v37)
						var _879_v40 _dafny.Map
						_ = _879_v40
						_879_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int), (func() _dafny.Set {
							if (_878_v39).Contains(_dafny.CodePoint('j')) {
								return (_878_v39).Get(_dafny.CodePoint('j')).(_dafny.Set)
							}
							return _876_v37
						})())
						var _880_v41 _dafny.Array
						_ = _880_v41
						var _nwElement0_35 _dafny.Set = (_876_v37).Intersection(_876_v37)
						_ = _nwElement0_35
						var _nw132 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(25))
						_ = _nw132
						(_nw132).ArraySet1(_nwElement0_35, 0)
						(_nw132).ArraySet1((_876_v37).Intersection((_877_v38)), 1)
						(_nw132).ArraySet1((_dafny.SetOf(_874_v36, _874_v36)).Intersection(_876_v37), 2)
						(_nw132).ArraySet1(_876_v37, 3)
						(_nw132).ArraySet1(_876_v37, 4)
						(_nw132).ArraySet1((_876_v37).Intersection(_dafny.SetOf(_874_v36)), 5)
						(_nw132).ArraySet1(_876_v37, 6)
						(_nw132).ArraySet1(_876_v37, 7)
						(_nw132).ArraySet1(_dafny.SetOf(_874_v36, _874_v36, _874_v36), 8)
						(_nw132).ArraySet1(((func() _dafny.Set {
							if (_879_v40).Contains((_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int)) {
								return (_879_v40).Get((_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int)).(_dafny.Set)
							}
							return _876_v37
						})()).Intersection(_dafny.SetOf(_874_v36, _874_v36, _874_v36, _874_v36, _874_v36)), 9)
						(_nw132).ArraySet1(_876_v37, 10)
						(_nw132).ArraySet1(_876_v37, 11)
						(_nw132).ArraySet1((func() _dafny.Set {
							if _854_v18 {
								return _876_v37
							}
							return _876_v37
						})(), 12)
						(_nw132).ArraySet1((_dafny.SetOf(_874_v36)).Intersection(_876_v37), 13)
						(_nw132).ArraySet1(_876_v37, 14)
						(_nw132).ArraySet1(_876_v37, 15)
						(_nw132).ArraySet1((_876_v37).Difference(_876_v37), 16)
						(_nw132).ArraySet1((_876_v37).Union(_876_v37), 17)
						(_nw132).ArraySet1(_876_v37, 18)
						(_nw132).ArraySet1(_876_v37, 19)
						(_nw132).ArraySet1(_dafny.SetOf(_874_v36, _874_v36), 20)
						(_nw132).ArraySet1(_876_v37, 21)
						(_nw132).ArraySet1(_876_v37, 22)
						(_nw132).ArraySet1((_876_v37).Intersection(_876_v37), 23)
						(_nw132).ArraySet1((func() _dafny.Set {
							if _854_v18 {
								return _876_v37
							}
							return _876_v37
						})(), 24)
						_880_v41 = _nw132
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_880_v41), 0))
						_ = _index110
						(_880_v41).ArraySet1(_876_v37, (_index110).Int())
						var _881_v42 _dafny.Array
						_ = _881_v42
						var _len0_16 _dafny.Int = _dafny.IntOfInt64(21)
						_ = _len0_16
						var _nw133 _dafny.Array
						_ = _nw133
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw133 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) bool = (func(_882_v19 bool) func(_dafny.Int) bool {
								return func(_883_i6 _dafny.Int) bool {
									return _882_v19
								}
							})(_855_v19)
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw133 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw133).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw133).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_881_v42 = _nw133
						var _884_v43 _dafny.Map
						_ = _884_v43
						_884_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v42, _847_v11)
						var _885_v44 _dafny.Array
						_ = _885_v44
						var _nwElement0_36 _dafny.Map = (func() _dafny.Map {
							if !(!(Companion_Default___.Fm0(_850_v14, globalState))) {
								return _884_v43
							}
							return (_884_v43).Update(_881_v42, _854_v18)
						})()
						_ = _nwElement0_36
						var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(8))
						_ = _nw134
						(_nw134).ArraySet1(_nwElement0_36, 0)
						(_nw134).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v42, _856_v20), 1)
						(_nw134).ArraySet1(_884_v43, 2)
						(_nw134).ArraySet1((_884_v43).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v42, _854_v18)), 3)
						(_nw134).ArraySet1(_884_v43, 4)
						(_nw134).ArraySet1(_884_v43, 5)
						(_nw134).ArraySet1(_884_v43, 6)
						(_nw134).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v42, _855_v19), 7)
						_885_v44 = _nw134
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_885_v44), 0))
						_ = _index111
						(_885_v44).ArraySet1(_884_v43, (_index111).Int())
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_880_v41), 0))
						_ = _index112
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_885_v44), 0))
						_ = _index113
						var _rhs107 _dafny.Set = (func() _dafny.Set {
							if _856_v20 {
								return _876_v37
							}
							return _876_v37
						})()
						_ = _rhs107
						var _rhs108 _dafny.Map = (_884_v43).Merge(_884_v43)
						_ = _rhs108
						var _lhs85 _dafny.Array = _880_v41
						_ = _lhs85
						var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_880_v41), 0))
						_ = _lhs86
						var _lhs87 _dafny.Array = _885_v44
						_ = _lhs87
						var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_885_v44), 0))
						_ = _lhs88
						(_lhs85).ArraySet1(_rhs107, (_lhs86).Int())
						(_lhs87).ArraySet1(_rhs108, (_lhs88).Int())
						var _886_v45 _dafny.Map
						_ = _886_v45
						_886_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_874_v36, (_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int))
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _index114
						var _rhs109 _dafny.Int = ((func() _dafny.Int {
							if (_857_v21).Contains((_this).F25()) {
								return (_857_v21).Multiplicity((_this).F25())
							}
							return (_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int)
						})()).Minus((_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int))
						_ = _rhs109
						var _rhs110 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_853_v17).Cardinality()), (_886_v45).Cardinality()), (func() _dafny.Int {
							if _855_v19 {
								return (_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int)
							}
							return _dafny.IntOfInt64(180)
						})()))
						_ = _rhs110
						var _lhs89 *GlobalState = globalState
						_ = _lhs89
						var _lhs90 _dafny.Array = _865_v28
						_ = _lhs90
						var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _lhs91
						_lhs89.F8 = _rhs109
						(_lhs90).ArraySet1(_rhs110, (_lhs91).Int())
						_864_v27 = (func() _dafny.CodePoint {
							if _847_v11 {
								return _864_v27
							}
							return _864_v27
						})()
						var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))
						_ = _index115
						(_865_v28).ArraySet1(Companion_Default___.SafeModuloInt(((_865_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(155), _dafny.ArrayLen((_865_v28), 0))).Int()).(_dafny.Int)).Times((_this).F25()), (_this).F25()), (_index115).Int())
					} else {
						var _887_v46 D6
						_ = _887_v46
						_887_v46 = Companion_D6_.Create_DC10_(_865_v28)
						(globalState).F23 = (_887_v46).Dtor_cf17()
						var _888_v47 _dafny.MultiSet
						_ = _888_v47
						_888_v47 = _dafny.MultiSetOf(true, ((_dafny.Zero).Minus((_this).F25())).Cmp((_this).F25()) == 0)
						_888_v47 = _888_v47
						var _889_v48 D3
						_ = _889_v48
						_889_v48 = Companion_D3_.Create_DC4_(_dafny.IntOfInt64(372), _855_v19, _856_v20)
						var _890_v49 _dafny.Map
						_ = _890_v49
						_890_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18(_dafny.MultiSetOf((_851_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))).Int()).(bool), _855_v19, _854_v18, !(_854_v18), _847_v11), _867_v30, true, globalState), _889_v48)
						var _891_v50 _dafny.Sequence
						_ = _891_v50
						_891_v50 = _dafny.SeqOf(_890_v49, _890_v49, _890_v49)
						var _892_v51 _dafny.Map
						_ = _892_v51
						_892_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_834_v1, (_this).F25())
						var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
						_ = _index116
						var _rhs111 _dafny.CodePoint = (func() _dafny.CodePoint {
							if (_851_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))).Int()).(bool) {
								return Companion_Default___.Fm22(_854_v18, globalState)
							}
							return _864_v27
						})()
						_ = _rhs111
						var _rhs112 bool = true
						_ = _rhs112
						var _rhs113 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_832_v0).Cardinality()), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F25()), (_this).F25()))
						_ = _rhs113
						var _rhs114 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_891_v50, (Companion_Default___.SafeIndex((func() _dafny.Int {
							if (_892_v51).Contains(_834_v1) {
								return (_892_v51).Get(_834_v1).(_dafny.Int)
							}
							return _dafny.IntOfInt64(563)
						})(), _dafny.IntOfUint32((_891_v50).Cardinality()))).Uint32(), _890_v49), _891_v50)
						_ = _rhs114
						var _lhs92 *GlobalState = globalState
						_ = _lhs92
						var _lhs93 *GlobalState = globalState
						_ = _lhs93
						var _lhs94 _dafny.Array = _851_v15
						_ = _lhs94
						var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_851_v15), 0))
						_ = _lhs95
						_864_v27 = _rhs111
						_lhs92.F17 = _rhs112
						_lhs93.F8 = _rhs113
						(_lhs94).ArraySet1(_rhs114, (_lhs95).Int())
						(globalState).F19 = !(_854_v18)
						var _893_v52 _dafny.Array
						_ = _893_v52
						var _nw135 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw135
						_893_v52 = _nw135
						var _894_v53 _dafny.Array
						_ = _894_v53
						var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(24))
						_ = _nw136
						_894_v53 = _nw136
						var _895_v54 *C1
						_ = _895_v54
						var _nw137 *C1 = New_C1_()
						_ = _nw137
						_nw137.Ctor__(_860_v24, _862_v25)
						_895_v54 = _nw137
						var _896_v55 _dafny.Sequence
						_ = _896_v55
						_896_v55 = _dafny.SeqOf(_895_v54)
						var _897_v56 _dafny.Map
						_ = _897_v56
						_897_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v18, _896_v55)
						var _rhs115 bool = !(_847_v11)
						_ = _rhs115
						var _rhs116 _dafny.Int = (_dafny.SetOf(_894_v53)).Cardinality()
						_ = _rhs116
						var _rhs117 _dafny.Array = _865_v28
						_ = _rhs117
						var _rhs118 _dafny.Array = _893_v52
						_ = _rhs118
						var _rhs119 _dafny.Int = (func() _dafny.Int {
							if _856_v20 {
								return (_897_v56).Cardinality()
							}
							return (_this).F25()
						})()
						_ = _rhs119
						var _lhs96 *GlobalState = globalState
						_ = _lhs96
						var _lhs97 *GlobalState = globalState
						_ = _lhs97
						var _lhs98 *GlobalState = globalState
						_ = _lhs98
						_lhs96.F17 = _rhs115
						_lhs97.F8 = _rhs116
						_865_v28 = _rhs117
						_893_v52 = _rhs118
						_lhs98.F8 = _rhs119
					}
					var _898_v57 _dafny.MultiSet
					_ = _898_v57
					_898_v57 = _dafny.MultiSetOf(Companion_Default___.Fm0(_850_v14, globalState), _854_v18)
					var _899_v58 _dafny.Set
					_ = _899_v58
					_899_v58 = _dafny.SetOf((func() _dafny.Int {
						if (_857_v21).Contains((_this).F25()) {
							return (_857_v21).Multiplicity((_this).F25())
						}
						return _dafny.IntOfInt64(635)
					})())
					(globalState).F8 = (func() _dafny.Int {
						if (_898_v57).Contains((_899_v58).IsDisjointFrom(_899_v58)) {
							return (_898_v57).Multiplicity((_899_v58).IsDisjointFrom(_899_v58))
						}
						return (_this).F25()
					})()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _hi7 _dafny.Int = (_this).F25()
		_ = _hi7
		for _900_i7 := (_this).F25(); _900_i7.Cmp(_hi7) < 0; _900_i7 = _900_i7.Plus(_dafny.One) {
			(globalState).F22 = _856_v20
			var _901_v59 _dafny.Array
			_ = _901_v59
			var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw138
			_901_v59 = _nw138
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_901_v59), 0))
			_ = _index117
			(_901_v59).ArraySet1((_this).F25(), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(375), _dafny.ArrayLen((_901_v59), 0))
			_ = _index118
			(_901_v59).ArraySet1((_this).F25(), (_index118).Int())
			var _902_v60 _dafny.Map
			_ = _902_v60
			_902_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_854_v18, _856_v20)
			var _903_v61 bool
			_ = _903_v61
			var _904_v62 _dafny.Array
			_ = _904_v62
			var _905_v63 bool
			_ = _905_v63
			var _906_v64 _dafny.Int
			_ = _906_v64
			var _out34 bool
			_ = _out34
			var _out35 _dafny.Array
			_ = _out35
			var _out36 bool
			_ = _out36
			var _out37 _dafny.Int
			_ = _out37
			_out34, _out35, _out36, _out37 = (_this).M4(_dafny.SeqOf((_this).F25()), _902_v60, _847_v11, globalState)
			_903_v61 = _out34
			_904_v62 = _out35
			_905_v63 = _out36
			_906_v64 = _out37
			_856_v20 = _905_v63
		}
	}
}
func (_this *C4) M2(globalState *GlobalState) {
	{
		var _907_v0 bool
		_ = _907_v0
		_907_v0 = true
		var _908_v1 _dafny.Map
		_ = _908_v1
		_908_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_907_v0, !(_907_v0))
		var _909_v2 _dafny.Map
		_ = _909_v2
		_909_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_907_v0, !(_907_v0) || ((func() bool {
			if (_908_v1).Contains(_907_v0) {
				return (_908_v1).Get(_907_v0).(bool)
			}
			return _907_v0
		})()))
		_908_v1 = (_908_v1).Update(_907_v0, !(_907_v0))
		var _910_v3 _dafny.Sequence
		_ = _910_v3
		_910_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ymupw")
		var _911_v4 _dafny.MultiSet
		_ = _911_v4
		_911_v4 = _dafny.MultiSetOf(_907_v0, _907_v0)
		var _912_v5 _dafny.Sequence
		_ = _912_v5
		_912_v5 = _dafny.SeqOf((_this).F25(), (_this).F25(), (_911_v4).Cardinality())
		var _913_v6 _dafny.Map
		_ = _913_v6
		_913_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_910_v3, _912_v5)
		var _914_v7 _dafny.CodePoint
		_ = _914_v7
		_914_v7 = _dafny.CodePoint('x')
		_913_v6 = (_913_v6).Update(_dafny.Companion_Sequence_.Concatenate(_910_v3, _dafny.Companion_Sequence_.Update(_910_v3, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_910_v3).Cardinality()))).Uint32(), _914_v7)), _912_v5)
		var _915_v8 _dafny.Sequence
		_ = _915_v8
		_915_v8 = _dafny.SeqOf(_910_v3)
		var _916_v9 D3
		_ = _916_v9
		_916_v9 = Companion_D3_.Create_DC3_(_915_v8)
		var _917_v10 D3
		_ = _917_v10
		_917_v10 = Companion_D3_.Create_DC5_(_916_v9)
		var _918_i0 _dafny.Int
		_ = _918_i0
		_918_i0 = _dafny.Zero
		{
			var _pat_let_tv25 = _907_v0
			_ = _pat_let_tv25
			for func(_source15 D3) bool {
				if _source15.Is_DC4() {
					var _940___mcc_h0 _dafny.Int = _source15.Get_().(D3_DC4).Cf4
					_ = _940___mcc_h0
					var _941___mcc_h1 bool = _source15.Get_().(D3_DC4).Cf5
					_ = _941___mcc_h1
					var _942___mcc_h2 bool = _source15.Get_().(D3_DC4).Cf6
					_ = _942___mcc_h2
					var _943_cf6 bool = _942___mcc_h2
					_ = _943_cf6
					var _944_cf5 bool = _941___mcc_h1
					_ = _944_cf5
					var _945_cf4 _dafny.Int = _940___mcc_h0
					_ = _945_cf4
					return !(_943_cf6) || (_943_cf6)
				} else if _source15.Is_DC3() {
					var _946___mcc_h3 _dafny.Sequence = _source15.Get_().(D3_DC3).Cf3
					_ = _946___mcc_h3
					var _947_cf3 _dafny.Sequence = _946___mcc_h3
					_ = _947_cf3
					return !(_pat_let_tv25)
				} else {
					var _948___mcc_h4 D3 = _source15.Get_().(D3_DC5).Cf7
					_ = _948___mcc_h4
					var _949_cf7 D3 = _948___mcc_h4
					_ = _949_cf7
					return ((_this).F25()).Cmp((_this).F25()) >= 0
				}
			}(_917_v10) {
				{
					if (_918_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_918_i0 = (_918_i0).Plus(_dafny.One)
					_909_v2 = (_909_v2).Update(!(_907_v0), !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_this).F25(), (_this).F25()), _912_v5)))
					var _919_v11 _dafny.Array
					_ = _919_v11
					var _nwElement0_37 _dafny.Sequence = Companion_Default___.Fm1(globalState)
					_ = _nwElement0_37
					var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(12))
					_ = _nw139
					(_nw139).ArraySet1(_nwElement0_37, 0)
					(_nw139).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}((func(_920_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_921_i1 _dafny.Int) _dafny.CodePoint {
							return _920_v7
						}
					})(_914_v7))), 1)
					(_nw139).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fvis"), 2)
					(_nw139).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_910_v3, _910_v3), 3)
					(_nw139).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg56 _dafny.Int) interface{} {
							return coer56(arg56)
						}
					}((func(_922_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_923_i2 _dafny.Int) _dafny.CodePoint {
							return _922_v7
						}
					})(_914_v7))), 4)
					(_nw139).ArraySet1(_dafny.Companion_Sequence_.Update(_910_v3, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_910_v3).Cardinality()))).Uint32(), _914_v7), 5)
					(_nw139).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qknenuyap"), 6)
					(_nw139).ArraySet1(_910_v3, 7)
					(_nw139).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-977))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg57 _dafny.Int) interface{} {
							return coer57(arg57)
						}
					}((func(_924_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_925_i3 _dafny.Int) _dafny.CodePoint {
							return _924_v7
						}
					})(_914_v7))), 8)
					(_nw139).ArraySet1(_910_v3, 9)
					(_nw139).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(98))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg58 _dafny.Int) interface{} {
							return coer58(arg58)
						}
					}((func(_926_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_927_i4 _dafny.Int) _dafny.CodePoint {
							return _926_v7
						}
					})(_914_v7))), 10)
					(_nw139).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("s"), 11)
					_919_v11 = _nw139
					var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_919_v11), 0))
					_ = _index119
					(_919_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(837))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg59 _dafny.Int) interface{} {
							return coer59(arg59)
						}
					}((func(_928_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_929_i5 _dafny.Int) _dafny.CodePoint {
							return _928_v7
						}
					})(_914_v7))), _910_v3), (_index119).Int())
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_919_v11), 0))
					_ = _index120
					(_919_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_910_v3, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-410))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg60 _dafny.Int) interface{} {
							return coer60(arg60)
						}
					}(func(_930_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					})), (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-410))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}(func(_931_i6 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality()))).Uint32(), _914_v7), _910_v3)), (_index120).Int())
					var _932_v12 _dafny.Array
					_ = _932_v12
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_17
					var _nw140 _dafny.Array
					_ = _nw140
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw140 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) _dafny.Set = func(_933_i7 _dafny.Int) _dafny.Set {
							return (_dafny.SetOf((_this).F25())).Union(_dafny.SetOf((_this).F25(), (_this).F25()))
						}
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw140 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw140).ArraySet1(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw140).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_932_v12 = _nw140
					var _934_v13 _dafny.Set
					_ = _934_v13
					_934_v13 = _dafny.SetOf((_this).F25(), (_this).F25(), (_this).F25())
					var _935_v14 _dafny.Sequence
					_ = _935_v14
					_935_v14 = _dafny.SeqOf(_934_v13)
					var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_932_v12), 0))
					_ = _index121
					(_932_v12).ArraySet1(((_935_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.IntOfUint32((_935_v14).Cardinality()))).Uint32()).(_dafny.Set)).Union(_dafny.SetOf((_this).F25())), (_index121).Int())
					var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_932_v12), 0))
					_ = _index122
					(_932_v12).ArraySet1(_934_v13, (_index122).Int())
					var _936_v15 _dafny.Array
					_ = _936_v15
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(16)
					_ = _len0_18
					var _nw141 _dafny.Array
					_ = _nw141
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw141 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) _dafny.Int = (func(_937_v0 bool) func(_dafny.Int) _dafny.Int {
							return func(_938_i8 _dafny.Int) _dafny.Int {
								return (_938_i8).Times(_dafny.IntOfUint32((_dafny.SeqOf(_937_v0, _937_v0)).Cardinality()))
							}
						})(_907_v0)
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
					_936_v15 = _nw141
					var _939_v16 *C3
					_ = _939_v16
					var _nw142 *C3 = New_C3_()
					_ = _nw142
					_nw142.Ctor__((func() _dafny.Array {
						if !(_907_v0) {
							return _936_v15
						}
						return _936_v15
					})())
					_939_v16 = _nw142
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		(globalState).F17 = _dafny.Companion_Sequence_.Contains(_912_v5, Companion_Default___.SafeModuloInt((_this).F25(), (_this).F25()))
		var _950_v17 _dafny.Array
		_ = _950_v17
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_19
		var _nw143 _dafny.Array
		_ = _nw143
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw143 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Int = func(_951_i9 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_951_i9, (_this).F25())
			}
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw143 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw143).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw143).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_950_v17 = _nw143
		(globalState).F23 = _950_v17
		var _952_v18 _dafny.Map
		_ = _952_v18
		_952_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
		var _953_v19 _dafny.Sequence
		_ = _953_v19
		_953_v19 = _dafny.SeqOf(_907_v0)
		var _954_v20 _dafny.Map
		_ = _954_v20
		_954_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_952_v18, _953_v19)
		var _955_v21 _dafny.Array
		_ = _955_v21
		var _nwElement0_38 bool = (true) && (true)
		_ = _nwElement0_38
		var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(5))
		_ = _nw144
		(_nw144).ArraySet1(_nwElement0_38, 0)
		(_nw144).ArraySet1(Companion_Default___.Fm0((func() _dafny.Sequence {
			if (_954_v20).Contains(_952_v18) {
				return (_954_v20).Get(_952_v18).(_dafny.Sequence)
			}
			return _953_v19
		})(), globalState), 1)
		(_nw144).ArraySet1(_907_v0, 2)
		(_nw144).ArraySet1(_907_v0, 3)
		(_nw144).ArraySet1(!(_907_v0), 4)
		_955_v21 = _nw144
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_955_v21), 0))
		_ = _index123
		(_955_v21).ArraySet1(!(_907_v0), (_index123).Int())
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_955_v21), 0))
		_ = _index124
		(_955_v21).ArraySet1(Companion_Default___.Fm0(_dafny.Companion_Sequence_.Concatenate(_953_v19, _953_v19), globalState), (_index124).Int())
	}
}
func (_this *C4) M3(p0 _dafny.Set, p1 bool, p2 _dafny.Array, p3 _dafny.Sequence, globalState *GlobalState) (bool, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _956_v0 _dafny.Sequence
		_ = _956_v0
		_956_v0 = _dafny.SeqOf(true)
		var _957_v1 _dafny.Sequence
		_ = _957_v1
		_957_v1 = _dafny.SeqOf(Companion_Default___.Fm0(_956_v0, globalState))
		(globalState).F8 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_957_v1, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, Companion_Default___.Fm0(_956_v0, globalState)), _956_v0))).Cardinality())
		var _hi8 _dafny.Int = (_this).F25()
		_ = _hi8
		for _958_i0 := (_this).F25(); _958_i0.Cmp(_hi8) < 0; _958_i0 = _958_i0.Plus(_dafny.One) {
			var _959_v2 *C0
			_ = _959_v2
			var _nw145 *C0 = New_C0_()
			_ = _nw145
			_nw145.Ctor__(p1)
			_959_v2 = _nw145
			var _960_v3 _dafny.MultiSet
			_ = _960_v3
			_960_v3 = _dafny.MultiSetOf(p1, p1, (_959_v2).F27(), (_959_v2).F27())
			_960_v3 = (_960_v3).Union(_960_v3)
			var _961_v5 _dafny.Array
			_ = _961_v5
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_20
			var _nw146 _dafny.Array
			_ = _nw146
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw146 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Set = (func(_962_i0 _dafny.Int) func(_dafny.Int) _dafny.Set {
					return func(_963_i1 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_dafny.SeqOf(_962_i0, (_dafny.Zero).Minus((_this).F25())))).Intersection(func() _dafny.Set {
							var _coll32 = _dafny.NewBuilder()
							_ = _coll32
							for _iter36 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(798))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg62 _dafny.Int) interface{} {
									return coer62(arg62)
								}
							}((func(_964_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_965_i2 _dafny.Int) _dafny.Int {
									return (_dafny.Zero).Minus(_964_i0)
								}
							})(_962_i0))), _dafny.SeqOf(_962_i0, (_this).F25()))).Elements()); ; {
								_compr_32, _ok36 := _iter36()
								if !_ok36 {
									break
								}
								var _966_v4 _dafny.Sequence
								_966_v4 = interface{}(_compr_32).(_dafny.Sequence)
								if (_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(798))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
									return func(arg63 _dafny.Int) interface{} {
										return coer63(arg63)
									}
								}((func(_967_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
									return func(_965_i2 _dafny.Int) _dafny.Int {
										return (_dafny.Zero).Minus(_967_i0)
									}
								})(_962_i0))), _dafny.SeqOf(_962_i0, (_this).F25()))).Contains(_966_v4) {
									_coll32.Add(_966_v4)
								}
							}
							return _coll32.ToSet()
						}())
					}
				})(_958_i0)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw146 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw146).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw146).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_961_v5 = _nw146
			var _968_v6 _dafny.Sequence
			_ = _968_v6
			_968_v6 = _dafny.SeqOf(_958_i0)
			var _969_v7 D13
			_ = _969_v7
			_969_v7 = Companion_D13_.Create_DC28_(_968_v6)
			var _970_v8 _dafny.Set
			_ = _970_v8
			_970_v8 = _dafny.SetOf((_969_v7).Dtor_cf47(), _dafny.SeqOf((_this).F25(), _958_i0), _dafny.Companion_Sequence_.Update(_968_v6, (Companion_Default___.SafeIndex(_958_i0, _dafny.IntOfUint32((_968_v6).Cardinality()))).Uint32(), _958_i0), _968_v6)
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_961_v5), 0))
			_ = _index125
			(_961_v5).ArraySet1((_970_v8).Difference(_dafny.SetOf(_968_v6, _968_v6)), (_index125).Int())
			var _971_v9 _dafny.Sequence
			_ = _971_v9
			_971_v9 = _dafny.SeqOf(_968_v6)
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_961_v5), 0))
			_ = _index126
			(_961_v5).ArraySet1(func() _dafny.Set {
				var _coll33 = _dafny.NewBuilder()
				_ = _coll33
				for _iter37 := _dafny.Iterate((_971_v9).Elements()); ; {
					_compr_33, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _972_v10 _dafny.Sequence
					_972_v10 = interface{}(_compr_33).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_971_v9, _972_v10) {
						_coll33.Add(_972_v10)
					}
				}
				return _coll33.ToSet()
			}(), (_index126).Int())
			var _973_v11 _dafny.CodePoint
			_ = _973_v11
			_973_v11 = _dafny.CodePoint('k')
			var _974_v12 D8
			_ = _974_v12
			_974_v12 = Companion_D8_.Create_DC17_(_dafny.CodePoint('c'), _973_v11, _dafny.IntOfUint32((p3).Cardinality()))
			var _975_v13 D12
			_ = _975_v13
			_975_v13 = Companion_D12_.Create_DC27_((_959_v2).F27(), (_this).F25())
			(globalState).F8 = Companion_Default___.Fm18(_960_v3, _974_v12, ((_975_v13).Dtor_cf46()).Cmp((_this).F25()) != 0, globalState)
		}
		var _976_v14 _dafny.Set
		_ = _976_v14
		_976_v14 = _dafny.SetOf((_this).F25())
		var _977_v15 _dafny.CodePoint
		_ = _977_v15
		_977_v15 = _dafny.CodePoint('j')
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((p2), 0))
		_ = _index127
		(p2).ArraySet1(p1, (_index127).Int())
		var _978_v16 _dafny.Sequence
		_ = _978_v16
		_978_v16 = _dafny.SeqOf(_976_v14, _976_v14)
		var _979_v17 *C0
		_ = _979_v17
		var _nw147 *C0 = New_C0_()
		_ = _nw147
		_nw147.Ctor__(p1)
		_979_v17 = _nw147
		var _980_v18 _dafny.Map
		_ = _980_v18
		_980_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F25()), _979_v17)
		var _981_v19 _dafny.Set
		_ = _981_v19
		_981_v19 = _dafny.SetOf((func() *C0 {
			if (_980_v18).Contains((_this).F25()) {
				return (_980_v18).Get((_this).F25()).(*C0)
			}
			return _979_v17
		})())
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((p2), 0))
		_ = _index128
		var _rhs120 _dafny.Set = (_978_v16).Select((Companion_Default___.SafeIndex(((_dafny.SetOf(_979_v17, _979_v17)).Union(_981_v19)).Cardinality(), _dafny.IntOfUint32((_978_v16).Cardinality()))).Uint32()).(_dafny.Set)
		_ = _rhs120
		var _rhs121 _dafny.CodePoint = Companion_Default___.Fm8(globalState)
		_ = _rhs121
		var _rhs122 bool = ((_this).F25()).Cmp(_dafny.IntOfInt64(980)) > 0
		_ = _rhs122
		var _rhs123 bool = !((_979_v17).F27())
		_ = _rhs123
		var _lhs99 *GlobalState = globalState
		_ = _lhs99
		var _lhs100 _dafny.Array = p2
		_ = _lhs100
		var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((p2), 0))
		_ = _lhs101
		_976_v14 = _rhs120
		_977_v15 = _rhs121
		_lhs99.F19 = _rhs122
		(_lhs100).ArraySet1(_rhs123, (_lhs101).Int())
		var _982_v20 _dafny.Array
		_ = _982_v20
		var _len0_21 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_21
		var _nw148 _dafny.Array
		_ = _nw148
		if _len0_21.Cmp(_dafny.Zero) == 0 {
			_nw148 = _dafny.NewArray(_len0_21)
		} else {
			var _init21 func(_dafny.Int) _dafny.Int = func(_983_i3 _dafny.Int) _dafny.Int {
				return (_983_i3).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F25())).Cardinality()))
			}
			_ = _init21
			var _element0_21 = _init21(_dafny.Zero)
			_ = _element0_21
			_nw148 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
			(_nw148).ArraySet1(_element0_21, 0)
			var _nativeLen0_21 = (_len0_21).Int()
			_ = _nativeLen0_21
			for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
				(_nw148).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
			}
		}
		_982_v20 = _nw148
		var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_982_v20), 0))
		_ = _index129
		(_982_v20).ArraySet1(_dafny.IntOfUint32((_956_v0).Cardinality()), (_index129).Int())
		var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_982_v20), 0))
		_ = _index130
		(_982_v20).ArraySet1(((_this).F25()).Plus(Companion_Default___.Fm5(globalState)), (_index130).Int())
		var _984_v21 _dafny.Map
		_ = _984_v21
		_984_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_982_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_982_v20), 0))).Int()).(_dafny.Int))
		var _hi9 _dafny.Int = (_984_v21).Cardinality()
		_ = _hi9
		for _985_i4 := (_this).F25(); _985_i4.Cmp(_hi9) < 0; _985_i4 = _985_i4.Plus(_dafny.One) {
			var _986_v22 D13
			_ = _986_v22
			_986_v22 = Companion_D13_.Create_DC28_(_dafny.SeqOf((_this).F25(), (_982_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_982_v20), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p3).Cardinality())))
			_986_v22 = _986_v22
			(globalState).F8 = (_982_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_982_v20), 0))).Int()).(_dafny.Int)
			var _987_v23 *C0
			_ = _987_v23
			var _nw149 *C0 = New_C0_()
			_ = _nw149
			_nw149.Ctor__(true)
			_987_v23 = _nw149
			var _988_v24 _dafny.Map
			_ = _988_v24
			_988_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_976_v14, _dafny.UnicodeSeqOfUtf8Bytes("p"))
			var _989_v25 T1
			_ = _989_v25
			var _nw150 *C0 = New_C0_()
			_ = _nw150
			_nw150.Ctor__(p1)
			_989_v25 = _nw150
			var _990_v26 T0
			_ = _990_v26
			var _nw151 *C1 = New_C1_()
			_ = _nw151
			_nw151.Ctor__(_988_v24, _989_v25)
			_990_v26 = _nw151
			var _991_v27 _dafny.Map
			_ = _991_v27
			_991_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm0(_956_v0, globalState))
			var _992_v28 _dafny.Map
			_ = _992_v28
			_992_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_990_v26, (func() _dafny.Map {
				if (_979_v17).F27() {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((p2), 0))).Int()).(bool)), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(145), _dafny.ArrayLen((p2), 0))).Int()).(bool))
				}
				return _991_v27
			})())
			_992_v28 = (_992_v28).Update(_990_v26, _991_v27)
		}
		var _993_v29 _dafny.Int
		_ = _993_v29
		var _994_v30 bool
		_ = _994_v30
		var _995_v31 bool
		_ = _995_v31
		var _996_v32 _dafny.Int
		_ = _996_v32
		var _out38 _dafny.Int
		_ = _out38
		var _out39 bool
		_ = _out39
		var _out40 bool
		_ = _out40
		var _out41 _dafny.Int
		_ = _out41
		_out38, _out39, _out40, _out41 = Companion_Default___.M0(globalState)
		_993_v29 = _out38
		_994_v30 = _out39
		_995_v31 = _out40
		_996_v32 = _out41
		r0 = Companion_Default___.Fm0(_956_v0, globalState)
		r1 = _995_v31
		r2 = p1
		return r0, r1, r2
	}
}
func (_this *C4) M4(p0 _dafny.Sequence, p1 _dafny.Map, p2 bool, globalState *GlobalState) (bool, _dafny.Array, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _997_v0 _dafny.Map
		_ = _997_v0
		_997_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("oapapxtbs"), p2)
		(globalState).F19 = (func() bool {
			if p2 {
				return (func() bool {
					if (_997_v0).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}(func(_998_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))) {
						return (_997_v0).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}(func(_999_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('f')
						}))).(bool)
					}
					return p2
				})()
			}
			return false
		})()
		var _1000_v1 _dafny.Sequence
		_ = _1000_v1
		_1000_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ahhoqv")
		var _rhs124 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg66 _dafny.Int) interface{} {
				return coer66(arg66)
			}
		}(func(_1001_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})), _1000_v1)
		_ = _rhs124
		var _rhs125 bool = !(true) || (p2)
		_ = _rhs125
		_1000_v1 = _rhs124
		r2 = _rhs125
		var _1002_v2 _dafny.CodePoint
		_ = _1002_v2
		_1002_v2 = _dafny.CodePoint('b')
		var _1003_v3 _dafny.Sequence
		_ = _1003_v3
		_1003_v3 = _dafny.SeqOf(p2)
		var _1004_v4 D4
		_ = _1004_v4
		_1004_v4 = Companion_D4_.Create_DC8_((func() _dafny.CodePoint {
			if !(p2) {
				return _1002_v2
			}
			return _1002_v2
		})(), Companion_Default___.Fm0(_1003_v3, globalState), p2)
		var _source16 D4 = _1004_v4
		_ = _source16
		if _source16.Is_DC7() {
			var _1005___mcc_h0 _dafny.Int = _source16.Get_().(D4_DC7).Cf9
			_ = _1005___mcc_h0
			var _1006___mcc_h1 *C0 = _source16.Get_().(D4_DC7).Cf10
			_ = _1006___mcc_h1
			var _1007___mcc_h2 bool = _source16.Get_().(D4_DC7).Cf11
			_ = _1007___mcc_h2
			var _1008___mcc_h3 _dafny.Int = _source16.Get_().(D4_DC7).Cf12
			_ = _1008___mcc_h3
			var _1009_cf12 _dafny.Int = _1008___mcc_h3
			_ = _1009_cf12
			var _1010_cf11 bool = _1007___mcc_h2
			_ = _1010_cf11
			var _1011_cf10 *C0 = _1006___mcc_h1
			_ = _1011_cf10
			var _1012_cf9 _dafny.Int = _1005___mcc_h0
			_ = _1012_cf9
			(globalState).F19 = _1010_cf11
			var _1013_v5 _dafny.Array
			_ = _1013_v5
			var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw152
			_1013_v5 = _nw152
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1013_v5), 0))
			_ = _index131
			(_1013_v5).ArraySet1(_dafny.IntOfUint32((_1000_v1).Cardinality()), (_index131).Int())
			var _1014_v6 _dafny.Map
			_ = _1014_v6
			_1014_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1002_v2)
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1013_v5), 0))
			_ = _index132
			(_1013_v5).ArraySet1((((_1014_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1002_v2))).Cardinality()).Times(_1009_cf12), (_index132).Int())
			var _1015_v7 _dafny.Int
			_ = _1015_v7
			_1015_v7 = (_1013_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_1013_v5), 0))).Int()).(_dafny.Int)
			var _1016_v8 _dafny.Map
			_ = _1016_v8
			_1016_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1010_cf11, _1015_v7)
			(globalState).F17 = !(_1016_v8).Contains(p2)
			var _1017_v9 D6
			_ = _1017_v9
			_1017_v9 = Companion_D6_.Create_DC10_(_1013_v5)
			var _1018_v10 D6
			_ = _1018_v10
			_1018_v10 = Companion_D6_.Create_DC12_(Companion_D6_.Create_DC12_(_1017_v9))
			var _1019_v11 D6
			_ = _1019_v11
			_1019_v11 = Companion_D6_.Create_DC12_(_1017_v9)
			var _1020_v12 D6
			_ = _1020_v12
			_1020_v12 = Companion_D6_.Create_DC12_(Companion_D6_.Create_DC12_(_1017_v9))
			var _1021_v13 T1
			_ = _1021_v13
			var _nw153 *C0 = New_C0_()
			_ = _nw153
			_nw153.Ctor__((_1011_cf10).F27())
			_1021_v13 = _nw153
			var _1022_v14 _dafny.Array
			_ = _1022_v14
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw154
			_1022_v14 = _nw154
			var _1023_v15 _dafny.Array
			_ = _1023_v15
			var _nwElement0_39 _dafny.Array = _1022_v14
			_ = _nwElement0_39
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(2))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_39, 0)
			(_nw155).ArraySet1(_1022_v14, 1)
			_1023_v15 = _nw155
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1023_v15), 0))
			_ = _index133
			(_1023_v15).ArraySet1(_1022_v14, (_index133).Int())
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1023_v15), 0))
			_ = _index134
			var _rhs126 D6 = Companion_D6_.Create_DC12_(_1018_v10)
			_ = _rhs126
			var _rhs127 T1 = _1021_v13
			_ = _rhs127
			var _rhs128 _dafny.Array = _1013_v5
			_ = _rhs128
			var _rhs129 _dafny.Array = _1022_v14
			_ = _rhs129
			var _rhs130 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1000_v1, _1000_v1)
			_ = _rhs130
			var _lhs102 _dafny.Array = _1023_v15
			_ = _lhs102
			var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_1023_v15), 0))
			_ = _lhs103
			_1020_v12 = _rhs126
			_1021_v13 = _rhs127
			_1013_v5 = _rhs128
			(_lhs102).ArraySet1(_rhs129, (_lhs103).Int())
			_1000_v1 = _rhs130
		} else if _source16.Is_DC8() {
			var _1024___mcc_h4 _dafny.CodePoint = _source16.Get_().(D4_DC8).Cf13
			_ = _1024___mcc_h4
			var _1025___mcc_h5 bool = _source16.Get_().(D4_DC8).Cf14
			_ = _1025___mcc_h5
			var _1026___mcc_h6 bool = _source16.Get_().(D4_DC8).Cf15
			_ = _1026___mcc_h6
			var _1027_cf15 bool = _1026___mcc_h6
			_ = _1027_cf15
			var _1028_cf14 bool = _1025___mcc_h5
			_ = _1028_cf14
			var _1029_cf13 _dafny.CodePoint = _1024___mcc_h4
			_ = _1029_cf13
			var _1030_v16 D7
			_ = _1030_v16
			_1030_v16 = Companion_D7_.Create_DC15_(_dafny.IntOfInt64(747), _1027_cf15, (_this).F25())
			var _1031_v17 _dafny.Set
			_ = _1031_v17
			_1031_v17 = _dafny.SetOf(p2)
			var _1032_v18 _dafny.Map
			_ = _1032_v18
			_1032_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1031_v17).Cardinality(), _1028_cf14)
			var _1033_v19 D12
			_ = _1033_v19
			_1033_v19 = Companion_D12_.Create_DC27_(p2, _dafny.IntOfInt64(-214))
			var _1034_v20 _dafny.MultiSet
			_ = _1034_v20
			_1034_v20 = _dafny.MultiSetOf(_1029_cf13, _1002_v2)
			var _1035_v21 _dafny.Array
			_ = _1035_v21
			var _nwElement0_40 bool = true
			_ = _nwElement0_40
			var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(27))
			_ = _nw156
			(_nw156).ArraySet1(_nwElement0_40, 0)
			(_nw156).ArraySet1((func() bool {
				if _1027_cf15 {
					return p2
				}
				return _1027_cf15
			})(), 1)
			(_nw156).ArraySet1(true, 2)
			(_nw156).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1000_v1, _1000_v1), 3)
			(_nw156).ArraySet1(((_this).F25()).Cmp((_this).F25()) <= 0, 4)
			(_nw156).ArraySet1(true, 5)
			(_nw156).ArraySet1((_1028_cf14) && (_1027_cf15), 6)
			(_nw156).ArraySet1((_1030_v16).Dtor_cf27(), 7)
			(_nw156).ArraySet1(((_this).F25()).Cmp((_1032_v18).Cardinality()) > 0, 8)
			(_nw156).ArraySet1(true, 9)
			(_nw156).ArraySet1(((_this).F25()).Cmp((_1032_v18).Cardinality()) > 0, 10)
			(_nw156).ArraySet1((func() bool {
				if _1027_cf15 {
					return _1028_cf14
				}
				return _1028_cf14
			})(), 11)
			(_nw156).ArraySet1(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_1003_v3, _1003_v3)), 12)
			(_nw156).ArraySet1(_1027_cf15, 13)
			(_nw156).ArraySet1((_1033_v19).Dtor_cf45(), 14)
			(_nw156).ArraySet1(_1028_cf14, 15)
			(_nw156).ArraySet1(_1027_cf15, 16)
			(_nw156).ArraySet1(_1027_cf15, 17)
			(_nw156).ArraySet1(true, 18)
			(_nw156).ArraySet1(false, 19)
			(_nw156).ArraySet1(p2, 20)
			(_nw156).ArraySet1(p2, 21)
			(_nw156).ArraySet1(_1027_cf15, 22)
			(_nw156).ArraySet1((_dafny.MultiSetOf(_1002_v2, _1029_cf13)).IsDisjointFrom(_1034_v20), 23)
			(_nw156).ArraySet1(_1028_cf14, 24)
			(_nw156).ArraySet1(_1028_cf14, 25)
			(_nw156).ArraySet1(false, 26)
			_1035_v21 = _nw156
			r1 = _1035_v21
			_1030_v16 = _1030_v16
			r3 = (_this).F25()
			var _1036_v23 _dafny.Array
			_ = _1036_v23
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_22
			var _nw157 _dafny.Array
			_ = _nw157
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw157 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Map = (func(_1037_p0 _dafny.Sequence, _1038_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_1039_i2 _dafny.Int) _dafny.Map {
						return (func() _dafny.Map {
							var _coll34 = _dafny.NewMapBuilder()
							_ = _coll34
							for _iter38 := _dafny.Iterate((_dafny.SeqOf(_1038_v1)).Elements()); ; {
								_compr_34, _ok38 := _iter38()
								if !_ok38 {
									break
								}
								var _1040_v22 _dafny.Sequence
								_1040_v22 = interface{}(_compr_34).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1038_v1), _1040_v22) {
									_coll34.Add(_1040_v22, (_this).F25())
								}
							}
							return _coll34.ToMap()
						}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yv"), _dafny.IntOfUint32((_1037_p0).Cardinality())))
					}
				})(p0, _1000_v1)
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
			_1036_v23 = _nw157
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_1036_v23), 0))
			_ = _index135
			(_1036_v23).ArraySet1(Companion_Default___.Fm29(false, p2, globalState), (_index135).Int())
			var _1041_v24 _dafny.Map
			_ = _1041_v24
			_1041_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v1, (_this).F25())
			var _1042_v25 _dafny.Sequence
			_ = _1042_v25
			_1042_v25 = _dafny.SeqOf(Companion_Default___.Fm1(globalState), _1000_v1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-864))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_1043_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1044_i3 _dafny.Int) _dafny.CodePoint {
					return _1043_v2
				}
			})(_1002_v2))))
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_1036_v23), 0))
			_ = _index136
			var _rhs131 _dafny.Map = _1041_v24
			_ = _rhs131
			var _rhs132 _dafny.Sequence = _dafny.SeqOf(p2)
			_ = _rhs132
			var _rhs133 _dafny.Int = (_this).F25()
			_ = _rhs133
			var _rhs134 _dafny.Int = ((func() _dafny.Map {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("k"), (_1042_v25).Select((Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1042_v25).Cardinality()))).Uint32()).(_dafny.Sequence)) {
					return p1
				}
				return p1
			})()).Cardinality()
			_ = _rhs134
			var _lhs104 _dafny.Array = _1036_v23
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(231), _dafny.ArrayLen((_1036_v23), 0))
			_ = _lhs105
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			(_lhs104).ArraySet1(_rhs131, (_lhs105).Int())
			_1003_v3 = _rhs132
			_lhs106.F8 = _rhs133
			_lhs107.F8 = _rhs134
		} else {
			var _1045___mcc_h7 _dafny.Set = _source16.Get_().(D4_DC6).Cf8
			_ = _1045___mcc_h7
			var _1046_cf8 _dafny.Set = _1045___mcc_h7
			_ = _1046_cf8
			var _1047_v26 _dafny.Map
			_ = _1047_v26
			_1047_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), ((_this).F25()).Minus((_1046_cf8).Cardinality()))
			_1047_v26 = (_1047_v26).Update((_this).F25(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1000_v1, _1000_v1)).Cardinality()))
			r3 = (_this).F25()
			r3 = (_this).F25()
			var _1048_v27 T1
			_ = _1048_v27
			var _nw158 *C0 = New_C0_()
			_ = _nw158
			_nw158.Ctor__(p2)
			_1048_v27 = _nw158
			var _1049_v28 T0
			_ = _1049_v28
			var _nw159 *C1 = New_C1_()
			_ = _nw159
			_nw159.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1046_cf8, _1000_v1), _1048_v27)
			_1049_v28 = _nw159
			var _1050_v29 _dafny.MultiSet
			_ = _1050_v29
			_1050_v29 = _dafny.MultiSetOf(_1049_v28, _1049_v28, _1049_v28)
			var _1051_v30 D15
			_ = _1051_v30
			_1051_v30 = Companion_D15_.Create_DC32_(_1050_v29)
			_1050_v29 = (_1050_v29).Difference((_1051_v30).Dtor_cf51())
		}
		var _1052_i4 _dafny.Int
		_ = _1052_i4
		_1052_i4 = _dafny.Zero
		{
			for p2 {
				{
					if (_1052_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1052_i4 = (_1052_i4).Plus(_dafny.One)
					(globalState).F19 = p2
					(globalState).F8 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_this).F25(), ((_this).F25()).Times((_this).F25())))
					var _1053_v31 _dafny.Array
					_ = _1053_v31
					var _nw160 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
					_ = _nw160
					_1053_v31 = _nw160
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1053_v31), 0))
					_ = _index137
					(_1053_v31).ArraySet1(p2, (_index137).Int())
					var _1054_v32 _dafny.Array
					_ = _1054_v32
					var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
					_ = _nw161
					_1054_v32 = _nw161
					var _1055_v33 _dafny.Array
					_ = _1055_v33
					var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
					_ = _nw162
					_1055_v33 = _nw162
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1053_v31), 0))
					_ = _index138
					var _rhs135 bool = p2
					_ = _rhs135
					var _rhs136 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1055_v33), _1055_v33)
					_ = _rhs136
					var _rhs137 bool = p2
					_ = _rhs137
					var _rhs138 _dafny.Array = _1054_v32
					_ = _rhs138
					var _lhs108 _dafny.Array = _1053_v31
					_ = _lhs108
					var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_1053_v31), 0))
					_ = _lhs109
					var _lhs110 *GlobalState = globalState
					_ = _lhs110
					(_lhs108).ArraySet1(_rhs135, (_lhs109).Int())
					_lhs110.F22 = _rhs136
					r2 = _rhs137
					_1054_v32 = _rhs138
					_997_v0 = ((_997_v0).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v1, Companion_Default___.Fm0(_dafny.Companion_Sequence_.Update(_1003_v3, (Companion_Default___.SafeIndex((_this).F25(), _dafny.IntOfUint32((_1003_v3).Cardinality()))).Uint32(), p2), globalState)))).Merge(_997_v0)
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1056_i5 _dafny.Int
		_ = _1056_i5
		_1056_i5 = _dafny.Zero
		{
			for p2 {
				{
					if (_1056_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1056_i5 = (_1056_i5).Plus(_dafny.One)
					var _1057_v34 _dafny.MultiSet
					_ = _1057_v34
					_1057_v34 = _dafny.MultiSetOf((_dafny.IntOfUint32((_1000_v1).Cardinality())).Cmp((_this).F25()) >= 0)
					var _1058_v35 D11
					_ = _1058_v35
					_1058_v35 = Companion_D11_.Create_DC24_(_1057_v34)
					_1057_v34 = (_1058_v35).Dtor_cf41()
					(globalState).F19 = (func() bool {
						if p2 {
							return p2
						}
						return p2
					})()
					var _1059_v36 _dafny.Map
					_ = _1059_v36
					_1059_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1003_v3)
					(globalState).F19 = Companion_Default___.Fm0((func() _dafny.Sequence {
						if p2 {
							return _dafny.SeqOf(p2)
						}
						return (func() _dafny.Sequence {
							if (_1059_v36).Contains(p2) {
								return (_1059_v36).Get(p2).(_dafny.Sequence)
							}
							return _1003_v3
						})()
					})(), globalState)
					(globalState).F8 = ((_this).F25()).Minus(_dafny.IntOfInt64(884))
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1060_v37 _dafny.Map
		_ = _1060_v37
		_1060_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F25(), (_this).F25())
		var _1061_v38 _dafny.Array
		_ = _1061_v38
		var _nwElement0_41 _dafny.Int = (_this).F25()
		_ = _nwElement0_41
		var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(8))
		_ = _nw163
		(_nw163).ArraySet1(_nwElement0_41, 0)
		(_nw163).ArraySet1((_this).F25(), 1)
		(_nw163).ArraySet1(_dafny.IntOfInt64(-867), 2)
		(_nw163).ArraySet1((_this).F25(), 3)
		(_nw163).ArraySet1((func() _dafny.Int {
			if (_1060_v37).Contains((_this).F25()) {
				return (_1060_v37).Get((_this).F25()).(_dafny.Int)
			}
			return (_this).F25()
		})(), 4)
		(_nw163).ArraySet1(((_this).F25()).Minus(_dafny.IntOfInt64(172)), 5)
		(_nw163).ArraySet1(((_this).F25()).Minus((_this).F25()), 6)
		(_nw163).ArraySet1((_this).F25(), 7)
		_1061_v38 = _nw163
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index139
		(_1061_v38).ArraySet1((_this).F25(), (_index139).Int())
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index140
		(_1061_v38).ArraySet1((_this).F25(), (_index140).Int())
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index141
		(_1061_v38).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((Companion_Default___.Fm5(globalState)).Times((_this).F25()))), (_index141).Int())
		var _1062_v39 _dafny.Sequence
		_ = _1062_v39
		_1062_v39 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bjxoc"), _1000_v1)
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index142
		var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index143
		var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_1061_v38), 0))
		_ = _index144
		var _rhs139 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if p2 {
				return _1062_v39
			}
			return _1062_v39
		})(), _dafny.Companion_Sequence_.Concatenate(_1062_v39, _1062_v39))).Cardinality())
		_ = _rhs139
		var _rhs140 _dafny.Int = _dafny.IntOfUint32((_1000_v1).Cardinality())
		_ = _rhs140
		var _rhs141 bool = p2
		_ = _rhs141
		var _rhs142 _dafny.Int = (((_this).F25()).Times((_this).F25())).Times((func() _dafny.Int {
			if p2 {
				return (_this).F25()
			}
			return (_this).F25()
		})())
		_ = _rhs142
		var _lhs111 _dafny.Array = _1061_v38
		_ = _lhs111
		var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1061_v38), 0))
		_ = _lhs112
		var _lhs113 _dafny.Array = _1061_v38
		_ = _lhs113
		var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_1061_v38), 0))
		_ = _lhs114
		var _lhs115 _dafny.Array = _1061_v38
		_ = _lhs115
		var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(877), _dafny.ArrayLen((_1061_v38), 0))
		_ = _lhs116
		(_lhs111).ArraySet1(_rhs139, (_lhs112).Int())
		(_lhs113).ArraySet1(_rhs140, (_lhs114).Int())
		r0 = _rhs141
		(_lhs115).ArraySet1(_rhs142, (_lhs116).Int())
		r0 = !(p2) || (_dafny.Companion_Sequence_.IsPrefixOf(_1000_v1, _dafny.UnicodeSeqOfUtf8Bytes("sjbp")))
		var _1063_v40 _dafny.Array
		_ = _1063_v40
		var _nwElement0_42 bool = p2
		_ = _nwElement0_42
		var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(27))
		_ = _nw164
		(_nw164).ArraySet1(_nwElement0_42, 0)
		(_nw164).ArraySet1(true, 1)
		(_nw164).ArraySet1(p2, 2)
		(_nw164).ArraySet1(p2, 3)
		(_nw164).ArraySet1(p2, 4)
		(_nw164).ArraySet1(p2, 5)
		(_nw164).ArraySet1(p2, 6)
		(_nw164).ArraySet1(!(p2), 7)
		(_nw164).ArraySet1(p2, 8)
		(_nw164).ArraySet1(p2, 9)
		(_nw164).ArraySet1(p2, 10)
		(_nw164).ArraySet1(p2, 11)
		(_nw164).ArraySet1(p2, 12)
		(_nw164).ArraySet1(p2, 13)
		(_nw164).ArraySet1(p2, 14)
		(_nw164).ArraySet1(p2, 15)
		(_nw164).ArraySet1(p2, 16)
		(_nw164).ArraySet1(p2, 17)
		(_nw164).ArraySet1(p2, 18)
		(_nw164).ArraySet1(p2, 19)
		(_nw164).ArraySet1(!(p2), 20)
		(_nw164).ArraySet1(p2, 21)
		(_nw164).ArraySet1(p2, 22)
		(_nw164).ArraySet1(p2, 23)
		(_nw164).ArraySet1(true, 24)
		(_nw164).ArraySet1(false, 25)
		(_nw164).ArraySet1(p2, 26)
		_1063_v40 = _nw164
		var _1064_v41 _dafny.Map
		_ = _1064_v41
		_1064_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1000_v1, _1063_v40)
		var _1065_v42 _dafny.Sequence
		_ = _1065_v42
		_1065_v42 = _dafny.SeqOf(_dafny.SeqOf(p2, p2), _1003_v3, _1003_v3)
		var _nwElement0_43 bool = !((_1064_v41).Update(_1000_v1, _1063_v40)).Equals(_1064_v41)
		_ = _nwElement0_43
		var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(5))
		_ = _nw165
		(_nw165).ArraySet1(_nwElement0_43, 0)
		(_nw165).ArraySet1(!(p2), 1)
		(_nw165).ArraySet1(Companion_Default___.Fm0((_1065_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1000_v1).Cardinality()), _dafny.IntOfUint32((_1065_v42).Cardinality()))).Uint32()).(_dafny.Sequence), globalState), 2)
		(_nw165).ArraySet1(p2, 3)
		(_nw165).ArraySet1(p2, 4)
		r1 = _nw165
		var _1066_v43 _dafny.Set
		_ = _1066_v43
		_1066_v43 = _dafny.SetOf(_dafny.IntOfInt64(-603))
		var _1067_v44 T1
		_ = _1067_v44
		var _nw166 *C0 = New_C0_()
		_ = _nw166
		_nw166.Ctor__(p2)
		_1067_v44 = _nw166
		var _1068_v45 *C1
		_ = _1068_v45
		var _nw167 *C1 = New_C1_()
		_ = _nw167
		_nw167.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1066_v43, _1000_v1), _1067_v44)
		_1068_v45 = _nw167
		var _1069_v46 D15
		_ = _1069_v46
		_1069_v46 = Companion_D15_.Create_DC33_(false, p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(171))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}(func(_1070_i6 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		}))).Cardinality()), _1068_v45, p2)
		var _1071_v47 D8
		_ = _1071_v47
		_1071_v47 = Companion_D8_.Create_DC17_(_1002_v2, _1002_v2, (_1069_v46).Dtor_cf54())
		r2 = ((_1061_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1061_v38), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.Fm18(_dafny.MultiSetOf(!(p2), p2, p2), _1071_v47, p2, globalState)) >= 0
		r3 = (_1061_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1061_v38), 0))).Int()).(_dafny.Int)
		return r0, r1, r2, r3
	}
}
func (_this *C4) F25() _dafny.Int {
	{
		return _this._f25
	}
}

// End of class C4
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
