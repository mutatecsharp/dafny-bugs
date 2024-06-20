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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(280), (_dafny.IntOfInt64(-786)).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-585))))
}
func (_static *CompanionStruct_Default___) Fm7(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(true)
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Map, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Sequence {
	if ((_dafny.Zero).Minus(_dafny.IntOfInt64(-465))).Cmp(_dafny.IntOfInt64(701)) <= 0 {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(377))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))
	} else {
		return _dafny.UnicodeSeqOfUtf8Bytes("xyrvtep")
	}
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('e')
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 bool, globalState *GlobalState) bool {
	return !((_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('l'))).Cardinality(), !(true)))).IsDisjointFrom((func() _dafny.MultiSet {
		if false {
			return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), true)))
		}
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vkne")).Cardinality()), false), func() _dafny.Map {
			var _coll0 = _dafny.NewMapBuilder()
			_ = _coll0
			for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(828), _dafny.IntOfInt64(917))); ; {
				_compr_0, _ok0 := _iter0()
				if !_ok0 {
					break
				}
				var _1_v0 _dafny.Int
				_1_v0 = interface{}(_compr_0).(_dafny.Int)
				if ((_dafny.IntOfInt64(828)).Cmp(_1_v0) <= 0) && ((_1_v0).Cmp(_dafny.IntOfInt64(917)) < 0) {
					_coll0.Add((_1_v0).Minus(_dafny.IntOfInt64(890)), !(false))
				}
			}
			return _coll0.ToMap()
		}(), func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), true)).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _2_v1 _dafny.Int
				_2_v1 = interface{}(_compr_1).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), true)).Contains(_2_v1) {
					_coll1.Add((_2_v1).Times(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), false)
				}
			}
			return _coll1.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iwq")).Cardinality()))).Cardinality(), true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(562), false)))
	})()))
}
func (_static *CompanionStruct_Default___) Fm14(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(606))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_3_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("gcrwhf")
	}))).Cardinality()), (_dafny.IntOfInt64(83)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(518), _dafny.CodePoint('k'))).Cardinality()), ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(641), false)).Cardinality())).Minus((_dafny.Zero).Minus((func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(518), _dafny.IntOfInt64(111))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _4_v0 _dafny.Int
			_4_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(518)).Cmp(_4_v0) <= 0) && ((_4_v0).Cmp(_dafny.IntOfInt64(111)) < 0) {
				_coll2.Add((_4_v0).Plus(_dafny.IntOfInt64(828)))
			}
		}
		return _coll2.ToSet()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC10_(_dafny.MultiSetOf(false, true))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(490))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_5_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(757), _dafny.IntOfInt64(86))).Cardinality()
	})))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(488)))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(880))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_6_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-421))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_7_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality())
	}))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(914), !(true))).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _8_v0 _dafny.Int
			_8_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Contains(_8_v0) {
				_coll3.Add((_8_v0).Times((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wnrwxf")).Cardinality()))).Cardinality())).Cardinality()), true)
			}
		}
		return _coll3.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!(true), false, true, false)).Cardinality())))).Cardinality(), _dafny.IntOfInt64(671)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(996))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_9_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-936)
	})))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(125))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(443), _dafny.IntOfInt64(719))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(931)), _dafny.SeqOf(_dafny.IntOfInt64(856))))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, false, !(true)), _dafny.SeqOf(!(true)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 D2, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source0 D12 = Companion_D12_.Create_DC30_(_dafny.UnicodeSeqOfUtf8Bytes("oausrgoo"), _dafny.IntOfInt64(442), _dafny.CodePoint('l'))
	_ = _source0
	if _source0.Is_DC29() {
		var _10___mcc_h0 _dafny.Sequence = _source0.Get_().(D12_DC29).Cf45
		_ = _10___mcc_h0
		var _11___mcc_h1 _dafny.Int = _source0.Get_().(D12_DC29).Cf46
		_ = _11___mcc_h1
		var _12___mcc_h2 bool = _source0.Get_().(D12_DC29).Cf47
		_ = _12___mcc_h2
		var _13_cf47 bool = _12___mcc_h2
		_ = _13_cf47
		var _14_cf46 _dafny.Int = _11___mcc_h1
		_ = _14_cf46
		var _15_cf45 _dafny.Sequence = _10___mcc_h0
		_ = _15_cf45
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(494))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg6 _dafny.Int) interface{} {
				return coer6(arg6)
			}
		}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('p')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(539))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_17_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		})))
	} else if _source0.Is_DC30() {
		var _18___mcc_h3 _dafny.Sequence = _source0.Get_().(D12_DC30).Cf48
		_ = _18___mcc_h3
		var _19___mcc_h4 _dafny.Int = _source0.Get_().(D12_DC30).Cf49
		_ = _19___mcc_h4
		var _20___mcc_h5 _dafny.CodePoint = _source0.Get_().(D12_DC30).Cf50
		_ = _20___mcc_h5
		var _21_cf50 _dafny.CodePoint = _20___mcc_h5
		_ = _21_cf50
		var _22_cf49 _dafny.Int = _19___mcc_h4
		_ = _22_cf49
		var _23_cf48 _dafny.Sequence = _18___mcc_h3
		_ = _23_cf48
		return _dafny.Companion_Sequence_.Concatenate(_23_cf48, _dafny.Companion_Sequence_.Update(_23_cf48, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-209), _dafny.IntOfUint32((_23_cf48).Cardinality()))).Uint32(), _dafny.CodePoint('o')))
	} else if _source0.Is_DC31() {
		var _24___mcc_h6 _dafny.Int = _source0.Get_().(D12_DC31).Cf51
		_ = _24___mcc_h6
		var _25___mcc_h7 bool = _source0.Get_().(D12_DC31).Cf52
		_ = _25___mcc_h7
		var _26___mcc_h8 bool = _source0.Get_().(D12_DC31).Cf53
		_ = _26___mcc_h8
		var _27___mcc_h9 _dafny.Int = _source0.Get_().(D12_DC31).Cf54
		_ = _27___mcc_h9
		var _28_cf54 _dafny.Int = _27___mcc_h9
		_ = _28_cf54
		var _29_cf53 bool = _26___mcc_h8
		_ = _29_cf53
		var _30_cf52 bool = _25___mcc_h7
		_ = _30_cf52
		var _31_cf51 _dafny.Int = _24___mcc_h6
		_ = _31_cf51
		return _dafny.UnicodeSeqOfUtf8Bytes("xdvd")
	} else {
		var _32___mcc_h10 _dafny.MultiSet = _source0.Get_().(D12_DC28).Cf44
		_ = _32___mcc_h10
		var _33_cf44 _dafny.MultiSet = _32___mcc_h10
		_ = _33_cf44
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_34_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))
	}
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false, !(!(false)), false, true)).Difference((_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ptgi"))).Union(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(474))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_35_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))))).Cardinality(), false, !(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_dafny.MultiSetOf(false, true)), _dafny.SeqOf(_dafny.MultiSetOf(false, !(true), true, true, !(false)), _dafny.MultiSetOf(!(true), true)))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(77)), _dafny.SeqOf((_dafny.SetOf(true, true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(479), false)).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(538), _dafny.IntOfInt64(754), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality()))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(949))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_36_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('h'))).Cardinality()))).Cardinality())
	}))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(230), _dafny.IntOfInt64(-506))))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(105))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_37_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})), _dafny.SeqOf(_dafny.CodePoint('a'))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_38_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(660))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_39_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))), _dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('p'), _dafny.CodePoint('n')))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true)).Difference(_dafny.SetOf(!(false), true))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC6_(_dafny.IntOfInt64(538), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(!(false), !(true), false, false, false), _dafny.SeqOf(true)), _dafny.SeqOf(_dafny.SeqOf(false))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-211), true)).Cardinality(), false)
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.CodePoint, p1 bool, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq((func() _dafny.Sequence {
		if true {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(974))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}(func(_40_i0 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(507))).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xwenmsl")).Cardinality()))).Cardinality())).Cardinality())).Cardinality()))
			})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-309))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}(func(_41_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf((_dafny.SetOf(false)).Cardinality())
			})))
		}
		return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(904))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg16 _dafny.Int) interface{} {
				return coer16(arg16)
			}
		}(func(_42_i2 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-134)
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-210))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_43_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqmyop")).Cardinality())
		})), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((Companion_D15_.Create_DC37_(func() _dafny.Map {
			var _coll4 = _dafny.NewMapBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('g'))).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _44_v1 _dafny.CodePoint
					_44_v1 = interface{}(_compr_5).(_dafny.CodePoint)
					if (_dafny.SetOf(_dafny.CodePoint('g'))).Contains(_44_v1) {
						_coll5.Add(_44_v1, _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))
					}
				}
				return _coll5.ToMap()
			}()).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _45_v0 _dafny.CodePoint
				_45_v0 = interface{}(_compr_4).(_dafny.CodePoint)
				if (func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('g'))).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _44_v1 _dafny.CodePoint
						_44_v1 = interface{}(_compr_6).(_dafny.CodePoint)
						if (_dafny.SetOf(_dafny.CodePoint('g'))).Contains(_44_v1) {
							_coll6.Add(_44_v1, _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))
						}
					}
					return _coll6.ToMap()
				}()).Contains(_45_v0) {
					_coll4.Add(_45_v0, !(!(true)))
				}
			}
			return _coll4.ToMap()
		}())).Dtor_cf64()).Cardinality(), true)).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(805), _dafny.IntOfInt64(-464))).Merge(func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(301))).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _46_v0 _dafny.Int
			_46_v0 = interface{}(_compr_7).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(9), _dafny.IntOfInt64(301))).Contains(_46_v0) {
				_coll7.Add((_46_v0).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D12_.Create_DC28_(_dafny.MultiSetOf(_dafny.IntOfInt64(-568), _dafny.IntOfInt64(314)))).Dtor_cf44(), _dafny.IntOfInt64(341))).Cardinality()), _dafny.IntOfInt64(503))
			}
		}
		return _coll7.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false, _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(507))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ghvpdnr")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), true)), (_dafny.MultiSetOf(true)).Cardinality())).Cardinality()), _dafny.IntOfInt64(179))).Cardinality())), _dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(504))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_47_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(94), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(449), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(985), _dafny.IntOfInt64(165))).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true)).Cardinality(), _dafny.IntOfInt64(425))).Cardinality()))).Cardinality()))).Cardinality())
	}))), _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(188), _dafny.IntOfInt64(-128))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ahqc")).Cardinality()))), _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(942)))))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("cxtyubjr")
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.CodePoint, p1 D3, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(758))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg19 _dafny.Int) interface{} {
			return coer19(arg19)
		}
	}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))).Cardinality()))), _dafny.SeqOf(func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('x'), _dafny.CodePoint('f'), _dafny.CodePoint('q'))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _49_v0 _dafny.CodePoint
			_49_v0 = interface{}(_compr_8).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('x'), _dafny.CodePoint('f'), _dafny.CodePoint('q')), _49_v0) {
				_coll8.Add(_49_v0, _dafny.IntOfInt64(386))
			}
		}
		return _coll8.ToMap()
	}())), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.IntOfInt64(538)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D4_.Create_DC13_(_dafny.CodePoint('e'))).Dtor_cf24(), (_dafny.MultiSetOf(_dafny.IntOfInt64(762), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("quneatwcc")).Cardinality()))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D7 {
	var _source1 D15 = Companion_D15_.Create_DC38_()
	_ = _source1
	if _source1.Is_DC38() {
		return Companion_D7_.Create_DC18_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), _dafny.IntOfInt64(-356))))
	} else {
		var _50___mcc_h0 _dafny.Map = _source1.Get_().(D15_DC37).Cf64
		_ = _50___mcc_h0
		var _51_cf64 _dafny.Map = _50___mcc_h0
		_ = _51_cf64
		return Companion_D7_.Create_DC18_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfInt64(217))))
	}
}
func (_static *CompanionStruct_Default___) Fm38(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Intersection(_dafny.SetOf(!(false)))).Difference(_dafny.SetOf(false, false, false, false))
}
func (_static *CompanionStruct_Default___) Fm39(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(426), _dafny.IntOfInt64(976), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rkh")).Cardinality())))).Difference(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-561))).Uint32(), func(coer20 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg20 _dafny.Int) interface{} {
			return coer20(arg20)
		}
	}(func(_52_i0 _dafny.Int) _dafny.Int {
		return (func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(562), _dafny.IntOfInt64(427))); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _53_v0 _dafny.Int
				_53_v0 = interface{}(_compr_9).(_dafny.Int)
				if ((_dafny.IntOfInt64(562)).Cmp(_53_v0) <= 0) && ((_53_v0).Cmp(_dafny.IntOfInt64(427)) < 0) {
					_coll9.Add(Companion_Default___.SafeDivisionInt(_53_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(357), (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Cardinality())))
				}
			}
			return _coll9.ToSet()
		}()).Cardinality()
	}))))).Intersection((_dafny.SetOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ddtk")).Cardinality()))).Cardinality()))))).Intersection(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(80)), _dafny.SeqOf(_dafny.IntOfInt64(798)), _dafny.SeqOf(_dafny.IntOfInt64(-939)))))
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC11_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(560), true)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, p3 _dafny.MultiSet, globalState *GlobalState) D4 {
	var _source2 D12 = Companion_D12_.Create_DC31_(_dafny.IntOfUint32((_dafny.SeqOf(false, (Companion_D12_.Create_DC29_(_dafny.SeqOf(_dafny.IntOfInt64(887), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(355))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_54_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality())), (func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(273), _dafny.IntOfInt64(623))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _55_v0 _dafny.Int
			_55_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(273)).Cmp(_55_v0) <= 0) && ((_55_v0).Cmp(_dafny.IntOfInt64(623)) < 0) {
				_coll10.Add((_55_v0).Plus(_dafny.IntOfInt64(518)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dbeopaxvs")).Cardinality()))
			}
		}
		return _coll10.ToMap()
	}()).Cardinality(), false)).Dtor_cf47())).Cardinality()), false, false, (func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(478), _dafny.IntOfInt64(125))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _56_v1 _dafny.Int
			_56_v1 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(478)).Cmp(_56_v1) <= 0) && ((_56_v1).Cmp(_dafny.IntOfInt64(125)) < 0) {
				_coll11.Add(Companion_Default___.SafeDivisionInt(_56_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-611))).Cardinality()), _dafny.CodePoint('g'))
			}
		}
		return _coll11.ToMap()
	}()).Cardinality())
	_ = _source2
	if _source2.Is_DC29() {
		var _57___mcc_h0 _dafny.Sequence = _source2.Get_().(D12_DC29).Cf45
		_ = _57___mcc_h0
		var _58___mcc_h1 _dafny.Int = _source2.Get_().(D12_DC29).Cf46
		_ = _58___mcc_h1
		var _59___mcc_h2 bool = _source2.Get_().(D12_DC29).Cf47
		_ = _59___mcc_h2
		var _60_cf47 bool = _59___mcc_h2
		_ = _60_cf47
		var _61_cf46 _dafny.Int = _58___mcc_h1
		_ = _61_cf46
		var _62_cf45 _dafny.Sequence = _57___mcc_h0
		_ = _62_cf45
		return Companion_D4_.Create_DC13_(_dafny.CodePoint('i'))
	} else if _source2.Is_DC30() {
		var _63___mcc_h3 _dafny.Sequence = _source2.Get_().(D12_DC30).Cf48
		_ = _63___mcc_h3
		var _64___mcc_h4 _dafny.Int = _source2.Get_().(D12_DC30).Cf49
		_ = _64___mcc_h4
		var _65___mcc_h5 _dafny.CodePoint = _source2.Get_().(D12_DC30).Cf50
		_ = _65___mcc_h5
		var _66_cf50 _dafny.CodePoint = _65___mcc_h5
		_ = _66_cf50
		var _67_cf49 _dafny.Int = _64___mcc_h4
		_ = _67_cf49
		var _68_cf48 _dafny.Sequence = _63___mcc_h3
		_ = _68_cf48
		return Companion_D4_.Create_DC13_(_66_cf50)
	} else if _source2.Is_DC31() {
		var _69___mcc_h6 _dafny.Int = _source2.Get_().(D12_DC31).Cf51
		_ = _69___mcc_h6
		var _70___mcc_h7 bool = _source2.Get_().(D12_DC31).Cf52
		_ = _70___mcc_h7
		var _71___mcc_h8 bool = _source2.Get_().(D12_DC31).Cf53
		_ = _71___mcc_h8
		var _72___mcc_h9 _dafny.Int = _source2.Get_().(D12_DC31).Cf54
		_ = _72___mcc_h9
		var _73_cf54 _dafny.Int = _72___mcc_h9
		_ = _73_cf54
		var _74_cf53 bool = _71___mcc_h8
		_ = _74_cf53
		var _75_cf52 bool = _70___mcc_h7
		_ = _75_cf52
		var _76_cf51 _dafny.Int = _69___mcc_h6
		_ = _76_cf51
		return Companion_D4_.Create_DC13_(_dafny.CodePoint('t'))
	} else {
		var _77___mcc_h10 _dafny.MultiSet = _source2.Get_().(D12_DC28).Cf44
		_ = _77___mcc_h10
		var _78_cf44 _dafny.MultiSet = _77___mcc_h10
		_ = _78_cf44
		return Companion_D4_.Create_DC13_(_dafny.CodePoint('c'))
	}
}
func (_static *CompanionStruct_Default___) Fm42(p0 _dafny.Int, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfInt64(130), (_dafny.IntOfInt64(483)).Times(_dafny.IntOfInt64(460)), (_dafny.IntOfInt64(954)).Minus((func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iatc")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(322), _dafny.IntOfInt64(477))); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _79_v1 _dafny.Int
				_79_v1 = interface{}(_compr_13).(_dafny.Int)
				if ((_dafny.IntOfInt64(322)).Cmp(_79_v1) <= 0) && ((_79_v1).Cmp(_dafny.IntOfInt64(477)) < 0) {
					_coll13.Add((_79_v1).Minus(_dafny.IntOfInt64(26)), false)
				}
			}
			return _coll13.ToMap()
		}()).Cardinality(), true)).Cardinality(), (_dafny.MultiSetOf(true)).Cardinality())).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _80_v0 _dafny.Int
			_80_v0 = interface{}(_compr_12).(_dafny.Int)
			if (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iatc")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll14 = _dafny.NewMapBuilder()
				_ = _coll14
				for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(322), _dafny.IntOfInt64(477))); ; {
					_compr_14, _ok14 := _iter14()
					if !_ok14 {
						break
					}
					var _79_v1 _dafny.Int
					_79_v1 = interface{}(_compr_14).(_dafny.Int)
					if ((_dafny.IntOfInt64(322)).Cmp(_79_v1) <= 0) && ((_79_v1).Cmp(_dafny.IntOfInt64(477)) < 0) {
						_coll14.Add((_79_v1).Minus(_dafny.IntOfInt64(26)), false)
					}
				}
				return _coll14.ToMap()
			}()).Cardinality(), true)).Cardinality(), (_dafny.MultiSetOf(true)).Cardinality())).Contains(_80_v0) {
				_coll12.Add((_80_v0).Plus(_dafny.IntOfInt64(905)), false)
			}
		}
		return _coll12.ToMap()
	}()).Cardinality()), _dafny.IntOfInt64(997), ((_dafny.Zero).Minus(_dafny.IntOfInt64(-292))).Times((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(253))).Cardinality()), _dafny.IntOfInt64(-790)), _dafny.IntOfInt64(334))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.CodePoint, p1 _dafny.Map, globalState *GlobalState) _dafny.Map {
	if (_dafny.IntOfInt64(908)).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qpimlci")).Cardinality())) != 0 {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(668), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-988))))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.CodePoint('x'))).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(956))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_81_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(251)
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm44(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("nlcjilh"), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("cec"), false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("q"), true))
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(480))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(554))))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()).Cmp(_dafny.IntOfInt64(-344)) <= 0 {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-546))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_82_i0 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(102))
		})), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true)).Cardinality())))
	} else {
		return _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(-804)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(625)))
	}
}
func (_static *CompanionStruct_Default___) M8(globalState *GlobalState) {
	var _83_v0 bool
	_ = _83_v0
	_83_v0 = false
	var _84_v1 _dafny.Sequence
	_ = _84_v1
	_84_v1 = _dafny.SeqOf(_83_v0, true, _83_v0, _83_v0, _83_v0)
	var _85_v2 _dafny.Int
	_ = _85_v2
	_85_v2 = _dafny.IntOfInt64(667)
	var _86_v3 _dafny.Sequence
	_ = _86_v3
	_86_v3 = _dafny.UnicodeSeqOfUtf8Bytes("rcmkx")
	(globalState).F13 = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_84_v1).Cardinality()))).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_85_v2, _dafny.IntOfUint32((_86_v3).Cardinality())))) < 0
	var _87_v4 _dafny.Array
	_ = _87_v4
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
	_ = _nw0
	_87_v4 = _nw0
	var _88_v5 _dafny.Array
	_ = _88_v5
	_88_v5 = _87_v4
	var _source3 _dafny.Array = _88_v5
	_ = _source3
	var _89___mcc_h0 _dafny.Array = _source3
	_ = _89___mcc_h0
	var _90_cf42 _dafny.Array = _89___mcc_h0
	_ = _90_cf42
	var _91_v6 _dafny.Array
	_ = _91_v6
	var _nw1 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
	_ = _nw1
	_91_v6 = _nw1
	var _92_v7 D4
	_ = _92_v7
	_92_v7 = Companion_D4_.Create_DC14_(_83_v0)
	var _93_v8 D4
	_ = _93_v8
	_93_v8 = Companion_D4_.Create_DC15_(_92_v7)
	var _94_v9 _dafny.MultiSet
	_ = _94_v9
	_94_v9 = _dafny.MultiSetOf(_85_v2)
	var _95_v10 _dafny.Array
	_ = _95_v10
	var _nwElement0_0 _dafny.Int = _dafny.IntOfInt64(-911)
	_ = _nwElement0_0
	var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
	_ = _nw2
	(_nw2).ArraySet1(_nwElement0_0, 0)
	(_nw2).ArraySet1((_dafny.Zero).Minus(_85_v2), 1)
	(_nw2).ArraySet1((_94_v9).Cardinality(), 2)
	(_nw2).ArraySet1(_dafny.IntOfInt64(994), 3)
	(_nw2).ArraySet1(_85_v2, 4)
	(_nw2).ArraySet1(_85_v2, 5)
	(_nw2).ArraySet1(_85_v2, 6)
	(_nw2).ArraySet1(_85_v2, 7)
	(_nw2).ArraySet1(_85_v2, 8)
	(_nw2).ArraySet1(_85_v2, 9)
	_95_v10 = _nw2
	var _96_v11 _dafny.Array
	_ = _96_v11
	_96_v11 = _95_v10
	var _97_v12 _dafny.Array
	_ = _97_v12
	var _nwElement0_1 _dafny.Array = _95_v10
	_ = _nwElement0_1
	var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(14))
	_ = _nw3
	(_nw3).ArraySet1(_nwElement0_1, 0)
	(_nw3).ArraySet1(_95_v10, 1)
	(_nw3).ArraySet1(_95_v10, 2)
	(_nw3).ArraySet1((_96_v11), 3)
	(_nw3).ArraySet1(_95_v10, 4)
	(_nw3).ArraySet1(_95_v10, 5)
	(_nw3).ArraySet1(_95_v10, 6)
	(_nw3).ArraySet1(_95_v10, 7)
	(_nw3).ArraySet1(_95_v10, 8)
	(_nw3).ArraySet1(_95_v10, 9)
	(_nw3).ArraySet1(_95_v10, 10)
	(_nw3).ArraySet1(_95_v10, 11)
	(_nw3).ArraySet1(_95_v10, 12)
	(_nw3).ArraySet1(_95_v10, 13)
	_97_v12 = _nw3
	var _98_v13 T1
	_ = _98_v13
	var _nw4 *C5 = New_C5_()
	_ = _nw4
	_nw4.Ctor__(_97_v12, _85_v2)
	_98_v13 = _nw4
	var _99_v14 _dafny.Sequence
	_ = _99_v14
	_99_v14 = _dafny.SeqOf(_90_cf42)
	var _100_v15 _dafny.Map
	_ = _100_v15
	_100_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v2, _85_v2)
	var _101_v16 _dafny.CodePoint
	_ = _101_v16
	_101_v16 = _dafny.CodePoint('i')
	var _102_v17 _dafny.Array
	_ = _102_v17
	var _nwElement0_2 _dafny.Int = _dafny.IntOfInt64(670)
	_ = _nwElement0_2
	var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(28))
	_ = _nw5
	(_nw5).ArraySet1(_nwElement0_2, 0)
	(_nw5).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfUint32((_86_v3).Cardinality()), globalState), 1)
	(_nw5).ArraySet1((func() _dafny.Int {
		if _83_v0 {
			return _dafny.IntOfInt64(203)
		}
		return _85_v2
	})(), 2)
	(_nw5).ArraySet1(_85_v2, 3)
	(_nw5).ArraySet1(_dafny.IntOfInt64(-330), 4)
	(_nw5).ArraySet1(_85_v2, 5)
	(_nw5).ArraySet1(_85_v2, 6)
	(_nw5).ArraySet1((Companion_D7_.Create_DC20_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}((func(_103_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_104_i0 _dafny.Int) _dafny.Int {
			return _103_v2
		}
	})(_85_v2))), _91_v6, _93_v8, _83_v0, _85_v2)).Dtor_cf34(), 7)
	(_nw5).ArraySet1((_dafny.Zero).Minus(_85_v2), 8)
	(_nw5).ArraySet1(_dafny.IntOfUint32((_84_v1).Cardinality()), 9)
	(_nw5).ArraySet1(_85_v2, 10)
	(_nw5).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v13, _91_v6)).Cardinality(), 11)
	(_nw5).ArraySet1((_dafny.Zero).Minus(_85_v2), 12)
	(_nw5).ArraySet1(_85_v2, 13)
	(_nw5).ArraySet1(_85_v2, 14)
	(_nw5).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_83_v0, Companion_Default___.Fm12(_83_v0, _83_v0, globalState), _83_v0, !(false)), (Companion_Default___.SafeIndex(_85_v2, _dafny.IntOfUint32((_dafny.SeqOf(_83_v0, Companion_Default___.Fm12(_83_v0, _83_v0, globalState), _83_v0, !(false))).Cardinality()))).Uint32(), _83_v0)).Cardinality()), 15)
	(_nw5).ArraySet1(_85_v2, 16)
	(_nw5).ArraySet1(_85_v2, 17)
	(_nw5).ArraySet1(_85_v2, 18)
	(_nw5).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_99_v14, _99_v14)).Cardinality()), 19)
	(_nw5).ArraySet1(_85_v2, 20)
	(_nw5).ArraySet1(_dafny.IntOfInt64(115), 21)
	(_nw5).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm44(_85_v2, _dafny.IntOfInt64(361), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_86_v3, (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_100_v15).Cardinality(), globalState), _dafny.IntOfUint32((_86_v3).Cardinality()))).Uint32(), _101_v16), (Companion_Default___.SafeIndex((_98_v13).Fm2(globalState), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_86_v3, (Companion_Default___.SafeIndex(Companion_Default___.Fm0((_100_v15).Cardinality(), globalState), _dafny.IntOfUint32((_86_v3).Cardinality()))).Uint32(), _101_v16)).Cardinality()))).Uint32(), _101_v16), globalState)).Cardinality()), 22)
	(_nw5).ArraySet1(_dafny.IntOfInt64(-60), 23)
	(_nw5).ArraySet1(_85_v2, 24)
	(_nw5).ArraySet1(Companion_Default___.SafeModuloInt(_85_v2, _dafny.IntOfUint32((_86_v3).Cardinality())), 25)
	(_nw5).ArraySet1(_85_v2, 26)
	(_nw5).ArraySet1(_85_v2, 27)
	_102_v17 = _nw5
	var _105_v18 _dafny.Set
	_ = _105_v18
	_105_v18 = _dafny.SetOf(_85_v2)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))
	_ = _index0
	(_95_v10).ArraySet1(Companion_Default___.SafeModuloInt((_105_v18).Cardinality(), _85_v2), (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))
	_ = _index1
	(_102_v17).ArraySet1(_85_v2, (_index1).Int())
	var _106_v19 _dafny.Array
	_ = _106_v19
	var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
	_ = _nw6
	_106_v19 = _nw6
	var _107_v20 _dafny.Array
	_ = _107_v20
	var _nw7 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
	_ = _nw7
	_107_v20 = _nw7
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_106_v19), 0))
	_ = _index2
	(_106_v19).ArraySet1(_107_v20, (_index2).Int())
	var _108_v21 *C5
	_ = _108_v21
	var _nw8 *C5 = New_C5_()
	_ = _nw8
	_nw8.Ctor__(_97_v12, (_dafny.Zero).Minus(_85_v2))
	_108_v21 = _nw8
	var _109_v22 _dafny.Set
	_ = _109_v22
	_109_v22 = _dafny.SetOf(_108_v21, _108_v21)
	var _110_v23 _dafny.Map
	_ = _110_v23
	_110_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_83_v0, _83_v0)
	var _111_v24 _dafny.MultiSet
	_ = _111_v24
	_111_v24 = _dafny.MultiSetOf(_83_v0, _83_v0)
	var _112_v25 _dafny.Map
	_ = _112_v25
	_112_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_101_v16, (_111_v24).IsDisjointFrom(Companion_Default___.Fm24(_85_v2, _83_v0, globalState)))
	var _113_v26 _dafny.Map
	_ = _113_v26
	_113_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(143), _107_v20)
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))
	_ = _index3
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))
	_ = _index4
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_106_v19), 0))
	_ = _index5
	var _rhs0 bool = !((func() bool {
		if (_109_v22).IsSubsetOf(_109_v22) {
			return _83_v0
		}
		return (func() bool {
			if (_110_v23).Contains(_83_v0) {
				return (_110_v23).Get(_83_v0).(bool)
			}
			return _83_v0
		})()
	})())
	_ = _rhs0
	var _rhs1 _dafny.Int = _85_v2
	_ = _rhs1
	var _rhs2 bool = (func() bool {
		if (_112_v25).Contains(_dafny.CodePoint('r')) {
			return (_112_v25).Get(_dafny.CodePoint('r')).(bool)
		}
		return !((_83_v0) || (_83_v0))
	})()
	_ = _rhs2
	var _rhs3 _dafny.Int = ((func() _dafny.Int {
		if _83_v0 {
			return _108_v21.F30
		}
		return _85_v2
	})()).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_86_v3, _86_v3)).Cardinality()))
	_ = _rhs3
	var _rhs4 _dafny.Array = (func() _dafny.Array {
		if _83_v0 {
			return (func() _dafny.Array {
				if (_113_v26).Contains(_108_v21.F30) {
					return (_113_v26).Get(_108_v21.F30).(_dafny.Array)
				}
				return _107_v20
			})()
		}
		return _107_v20
	})()
	_ = _rhs4
	var _lhs0 _dafny.Array = _95_v10
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))
	_ = _lhs1
	var _lhs2 *GlobalState = globalState
	_ = _lhs2
	var _lhs3 _dafny.Array = _102_v17
	_ = _lhs3
	var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))
	_ = _lhs4
	var _lhs5 _dafny.Array = _106_v19
	_ = _lhs5
	var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_106_v19), 0))
	_ = _lhs6
	_83_v0 = _rhs0
	(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
	_lhs2.F13 = _rhs2
	(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
	(_lhs5).ArraySet1(_rhs4, (_lhs6).Int())
	(globalState).F24 = _105_v18
	if (func() bool {
		if _83_v0 {
			return (_84_v1).Select((Companion_Default___.SafeIndex((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_84_v1).Cardinality()))).Uint32()).(bool)
		}
		return !(_83_v0)
	})() {
		var _114_v27 *C0
		_ = _114_v27
		var _nw9 *C0 = New_C0_()
		_ = _nw9
		_nw9.Ctor__(!(false))
		_114_v27 = _nw9
		var _rhs5 *C0 = _114_v27
		_ = _rhs5
		var _rhs6 _dafny.Array = _91_v6
		_ = _rhs6
		var _rhs7 _dafny.Int = _85_v2
		_ = _rhs7
		var _lhs7 *C5 = _108_v21
		_ = _lhs7
		_114_v27 = _rhs5
		_91_v6 = _rhs6
		_lhs7.F30 = _rhs7
		var _115_v28 _dafny.Array
		_ = _115_v28
		var _nw10 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
		_ = _nw10
		_115_v28 = _nw10
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_115_v28), 0))
		_ = _index6
		(_115_v28).ArraySet1(_98_v13, (_index6).Int())
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.ArrayLen((_115_v28), 0))
		_ = _index7
		(_115_v28).ArraySet1(_98_v13, (_index7).Int())
		var _116_v29 _dafny.Map
		_ = _116_v29
		_116_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(211), _91_v6)
		var _117_v30 _dafny.Set
		_ = _117_v30
		_117_v30 = _dafny.SetOf(_83_v0)
		var _118_v31 _dafny.Sequence
		_ = _118_v31
		_118_v31 = _dafny.SeqOf(_105_v18, _105_v18, Companion_Default___.Fm42((_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))).Int()).(_dafny.Int), _117_v30, (_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))).Int()).(_dafny.Int), globalState))
		var _119_v32 _dafny.Map
		_ = _119_v32
		_119_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_108_v21.F30), (_114_v27).F32())
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))
		_ = _index8
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))
		_ = _index9
		var _rhs8 _dafny.Map = _116_v29
		_ = _rhs8
		var _rhs9 bool = ((_105_v18).Difference(_105_v18)).IsSubsetOf((_105_v18).Union((_118_v31).Select((Companion_Default___.SafeIndex((_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_118_v31).Cardinality()))).Uint32()).(_dafny.Set)))
		_ = _rhs9
		var _rhs10 _dafny.Int = Companion_Default___.SafeModuloInt((_110_v23).Cardinality(), _85_v2)
		_ = _rhs10
		var _rhs11 _dafny.Int = (func() _dafny.Int {
			if !((_84_v1).Select((Companion_Default___.SafeIndex((_94_v9).Cardinality(), _dafny.IntOfUint32((_84_v1).Cardinality()))).Uint32()).(bool)) {
				return (_dafny.Zero).Minus(_108_v21.F30)
			}
			return (func() _dafny.Int {
				if _83_v0 {
					return _dafny.IntOfUint32((_dafny.SeqOf((func() bool {
						if (_119_v32).Contains((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int)) {
							return (_119_v32).Get((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int)).(bool)
						}
						return _83_v0
					})(), false)).Cardinality())
				}
				return _108_v21.F30
			})()
		})()
		_ = _rhs11
		var _rhs12 _dafny.Int = (_dafny.Zero).Minus((_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))).Int()).(_dafny.Int))
		_ = _rhs12
		var _lhs8 *GlobalState = globalState
		_ = _lhs8
		var _lhs9 *GlobalState = globalState
		_ = _lhs9
		var _lhs10 _dafny.Array = _95_v10
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))
		_ = _lhs11
		var _lhs12 _dafny.Array = _102_v17
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))
		_ = _lhs13
		_116_v29 = _rhs8
		_lhs8.F22 = _rhs9
		_lhs9.F10 = _rhs10
		(_lhs10).ArraySet1(_rhs11, (_lhs11).Int())
		(_lhs12).ArraySet1(_rhs12, (_lhs13).Int())
		var _120_v33 _dafny.Map
		_ = _120_v33
		_120_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_114_v27).F32(), _108_v21.F30)
		var _121_v34 _dafny.MultiSet
		_ = _121_v34
		_121_v34 = _dafny.MultiSetOf(_120_v33)
		var _122_v35 _dafny.Sequence
		_ = _122_v35
		_122_v35 = _dafny.SeqOf(_117_v30, _117_v30)
		var _123_v36 _dafny.Array
		_ = _123_v36
		var _nwElement0_3 _dafny.MultiSet = (_121_v34).Intersection(_dafny.MultiSetOf(_120_v33))
		_ = _nwElement0_3
		var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(25))
		_ = _nw11
		(_nw11).ArraySet1(_nwElement0_3, 0)
		(_nw11).ArraySet1(Companion_Default___.Fm45((_dafny.Zero).Minus((_95_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_95_v10), 0))).Int()).(_dafny.Int)), globalState), 1)
		(_nw11).ArraySet1(_dafny.MultiSetOf(_120_v33, _120_v33, _120_v33, _120_v33, _120_v33), 2)
		(_nw11).ArraySet1(_121_v34, 3)
		(_nw11).ArraySet1(_121_v34, 4)
		(_nw11).ArraySet1(_121_v34, 5)
		(_nw11).ArraySet1(_121_v34, 6)
		(_nw11).ArraySet1(_121_v34, 7)
		(_nw11).ArraySet1((_121_v34).Update(_120_v33, Companion_Default___.Abs((_98_v13).Fm2(globalState))), 8)
		(_nw11).ArraySet1(_121_v34, 9)
		(_nw11).ArraySet1(_dafny.MultiSetFromSeq(Companion_Default___.Fm46((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int), _85_v2, _108_v21.F30, globalState)), 10)
		(_nw11).ArraySet1((_121_v34).Intersection(_121_v34), 11)
		(_nw11).ArraySet1(_121_v34, 12)
		(_nw11).ArraySet1((_121_v34).Intersection(_121_v34), 13)
		(_nw11).ArraySet1(Companion_Default___.Fm45((_108_v21).Fm3((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int), (_114_v27).F32(), _dafny.IntOfInt64(-344), Companion_D0_.Create_DC2_((_102_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_102_v17), 0))).Int()).(_dafny.Int), true), globalState), globalState), 14)
		(_nw11).ArraySet1(_121_v34, 15)
		(_nw11).ArraySet1(_121_v34, 16)
		(_nw11).ArraySet1(_121_v34, 17)
		(_nw11).ArraySet1(_121_v34, 18)
		(_nw11).ArraySet1(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_114_v27).F32(), _dafny.IntOfUint32((_122_v35).Cardinality()))), 19)
		(_nw11).ArraySet1((_121_v34).Intersection(_121_v34), 20)
		(_nw11).ArraySet1(_121_v34, 21)
		(_nw11).ArraySet1(_121_v34, 22)
		(_nw11).ArraySet1(_121_v34, 23)
		(_nw11).ArraySet1(((_121_v34).Update(_120_v33, Companion_Default___.Abs(_108_v21.F30))).Union(_121_v34), 24)
		_123_v36 = _nw11
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_123_v36), 0))
		_ = _index10
		(_123_v36).ArraySet1(_121_v34, (_index10).Int())
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_123_v36), 0))
		_ = _index11
		(_123_v36).ArraySet1(_121_v34, (_index11).Int())
		(globalState).F19 = _101_v16
	} else {
		(globalState).F7 = _83_v0
		(globalState).F13 = _83_v0
		var _124_v37 *C1
		_ = _124_v37
		var _nw12 *C1 = New_C1_()
		_ = _nw12
		_nw12.Ctor__()
		_124_v37 = _nw12
		_124_v37 = _124_v37
		_108_v21 = _108_v21
		(_108_v21).F30 = _dafny.IntOfInt64(667)
	}
	_84_v1 = _84_v1
	var _125_v38 D14
	_ = _125_v38
	_125_v38 = Companion_D14_.Create_DC34_(true)
	var _126_v39 _dafny.Array
	_ = _126_v39
	var _nwElement0_4 bool = _83_v0
	_ = _nwElement0_4
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(9))
	_ = _nw13
	(_nw13).ArraySet1(_nwElement0_4, 0)
	(_nw13).ArraySet1(_83_v0, 1)
	(_nw13).ArraySet1(_dafny.Companion_Sequence_.Equal(_84_v1, _84_v1), 2)
	(_nw13).ArraySet1(_83_v0, 3)
	(_nw13).ArraySet1((_83_v0) == (_83_v0), 4)
	(_nw13).ArraySet1(_83_v0, 5)
	(_nw13).ArraySet1(_83_v0, 6)
	(_nw13).ArraySet1(_83_v0, 7)
	(_nw13).ArraySet1(false, 8)
	_126_v39 = _nw13
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))
	_ = _index12
	(_126_v39).ArraySet1(_83_v0, (_index12).Int())
	var _127_v40 _dafny.CodePoint
	_ = _127_v40
	_127_v40 = _dafny.CodePoint('v')
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))
	_ = _index13
	var _rhs13 D14 = Companion_D14_.Create_DC34_(_83_v0)
	_ = _rhs13
	var _rhs14 _dafny.Int = _85_v2
	_ = _rhs14
	var _rhs15 bool = _83_v0
	_ = _rhs15
	var _rhs16 _dafny.Int = _85_v2
	_ = _rhs16
	var _rhs17 _dafny.Sequence = Companion_Default___.Fm10(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v2, _85_v2), _127_v40, _83_v0, globalState)
	_ = _rhs17
	var _lhs14 *GlobalState = globalState
	_ = _lhs14
	var _lhs15 _dafny.Array = _126_v39
	_ = _lhs15
	var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))
	_ = _lhs16
	_125_v38 = _rhs13
	_lhs14.F10 = _rhs14
	(_lhs15).ArraySet1(_rhs15, (_lhs16).Int())
	_85_v2 = _rhs16
	_86_v3 = _rhs17
	(globalState).F10 = _dafny.IntOfUint32((_84_v1).Cardinality())
	(globalState).F7 = (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool)
	var _128_v41 _dafny.Map
	_ = _128_v41
	_128_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool), (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool))
	var _129_v42 _dafny.Map
	_ = _129_v42
	_129_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_85_v2), _85_v2)
	var _130_v43 _dafny.Sequence
	_ = _130_v43
	_130_v43 = _dafny.SeqOf((_128_v41).Cardinality(), _dafny.IntOfInt64(81), (_129_v42).Cardinality())
	var _hi0 _dafny.Int = _dafny.IntOfUint32((_130_v43).Cardinality())
	_ = _hi0
	for _131_i1 := _85_v2; _131_i1.Cmp(_hi0) < 0; _131_i1 = _131_i1.Plus(_dafny.One) {
		var _132_v44 D4
		_ = _132_v44
		_132_v44 = Companion_D4_.Create_DC14_(_83_v0)
		var _133_v45 _dafny.Sequence
		_ = _133_v45
		_133_v45 = _dafny.SeqOf(_132_v44, _132_v44, _132_v44)
		_133_v45 = _133_v45
		var _134_v46 _dafny.MultiSet
		_ = _134_v46
		_134_v46 = _dafny.MultiSetOf(_83_v0, _83_v0)
		(globalState).F10 = ((func() _dafny.Int {
			if (_134_v46).Contains(_83_v0) {
				return (_134_v46).Multiplicity(_83_v0)
			}
			return _85_v2
		})()).Plus(_85_v2)
		var _135_v47 _dafny.Set
		_ = _135_v47
		_135_v47 = _dafny.SetOf(_131_i1, _131_i1, _85_v2)
		var _rhs18 _dafny.Array = _87_v4
		_ = _rhs18
		var _rhs19 _dafny.Array = _126_v39
		_ = _rhs19
		var _rhs20 bool = (_128_v41).Equals(_128_v41)
		_ = _rhs20
		var _rhs21 bool = (func() bool {
			if !(_135_v47).Contains((_dafny.Zero).Minus(_131_i1)) {
				return (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool)
			}
			return (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool)
		})()
		_ = _rhs21
		var _lhs17 *GlobalState = globalState
		_ = _lhs17
		var _lhs18 *GlobalState = globalState
		_ = _lhs18
		_87_v4 = _rhs18
		_126_v39 = _rhs19
		_lhs17.F13 = _rhs20
		_lhs18.F7 = _rhs21
		var _136_v48 D7
		_ = _136_v48
		_136_v48 = Companion_D7_.Create_DC19_()
		var _source4 D7 = _136_v48
		_ = _source4
		if _source4.Is_DC19() {
			_126_v39 = _126_v39
			(globalState).F13 = true
			var _137_v49 _dafny.Array
			_ = _137_v49
			var _nwElement0_5 _dafny.Sequence = _86_v3
			_ = _nwElement0_5
			var _nw14 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(7))
			_ = _nw14
			(_nw14).ArraySet1(_nwElement0_5, 0)
			(_nw14).ArraySet1(_86_v3, 1)
			(_nw14).ArraySet1(_86_v3, 2)
			(_nw14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_127_v40), _86_v3), 3)
			(_nw14).ArraySet1(_86_v3, 4)
			(_nw14).ArraySet1(_86_v3, 5)
			(_nw14).ArraySet1(_86_v3, 6)
			_137_v49 = _nw14
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_137_v49), 0))
			_ = _index14
			(_137_v49).ArraySet1(_86_v3, (_index14).Int())
			var _138_v50 *C0
			_ = _138_v50
			var _nw15 *C0 = New_C0_()
			_ = _nw15
			_nw15.Ctor__(false)
			_138_v50 = _nw15
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_137_v49), 0))
			_ = _index15
			var _rhs22 _dafny.Sequence = _86_v3
			_ = _rhs22
			var _rhs23 *C0 = _138_v50
			_ = _rhs23
			var _lhs19 _dafny.Array = _137_v49
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_137_v49), 0))
			_ = _lhs20
			(_lhs19).ArraySet1(_rhs22, (_lhs20).Int())
			_138_v50 = _rhs23
			var _139_v51 _dafny.Map
			_ = _139_v51
			_139_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v40, _85_v2)
			var _140_v52 _dafny.Sequence
			_ = _140_v52
			_140_v52 = _dafny.SeqOf(_139_v51, _139_v51, _139_v51, _139_v51, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_127_v40, _131_i1))
			var _141_v53 D7
			_ = _141_v53
			_141_v53 = Companion_D7_.Create_DC18_(_140_v52)
			_141_v53 = _141_v53
		} else if _source4.Is_DC20() {
			var _142___mcc_h1 _dafny.Sequence = _source4.Get_().(D7_DC20).Cf30
			_ = _142___mcc_h1
			var _143___mcc_h2 _dafny.Array = _source4.Get_().(D7_DC20).Cf31
			_ = _143___mcc_h2
			var _144___mcc_h3 D4 = _source4.Get_().(D7_DC20).Cf32
			_ = _144___mcc_h3
			var _145___mcc_h4 bool = _source4.Get_().(D7_DC20).Cf33
			_ = _145___mcc_h4
			var _146___mcc_h5 _dafny.Int = _source4.Get_().(D7_DC20).Cf34
			_ = _146___mcc_h5
			var _147_cf34 _dafny.Int = _146___mcc_h5
			_ = _147_cf34
			var _148_cf33 bool = _145___mcc_h4
			_ = _148_cf33
			var _149_cf32 D4 = _144___mcc_h3
			_ = _149_cf32
			var _150_cf31 _dafny.Array = _143___mcc_h2
			_ = _150_cf31
			var _151_cf30 _dafny.Sequence = _142___mcc_h1
			_ = _151_cf30
			(globalState).F22 = _83_v0
			var _152_v54 _dafny.Map
			_ = _152_v54
			_152_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_85_v2, !(true) || (_148_cf33))
			_152_v54 = (_152_v54).Update(Companion_Default___.SafeModuloInt(_147_cf34, _147_cf34), _83_v0)
			_130_v43 = _dafny.SeqOf(Companion_Default___.Fm0(_85_v2, globalState), _147_cf34, _131_i1)
			var _153_v55 _dafny.Array
			_ = _153_v55
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_0
			var _nw16 _dafny.Array
			_ = _nw16
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw16 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Sequence = (func(_154_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_155_i2 _dafny.Int) _dafny.Sequence {
						return _154_v3
					}
				})(_86_v3)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw16 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw16).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw16).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_153_v55 = _nw16
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_153_v55), 0))
			_ = _index16
			(_153_v55).ArraySet1(_86_v3, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_153_v55), 0))
			_ = _index17
			(_153_v55).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if _148_cf33 {
					return _86_v3
				}
				return _86_v3
			})(), (Companion_Default___.SafeIndex(_131_i1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _148_cf33 {
					return _86_v3
				}
				return _86_v3
			})()).Cardinality()))).Uint32(), _127_v40), _dafny.UnicodeSeqOfUtf8Bytes("uyvafgoga")), (_index17).Int())
		} else {
			var _156___mcc_h6 _dafny.Sequence = _source4.Get_().(D7_DC18).Cf29
			_ = _156___mcc_h6
			var _157_cf29 _dafny.Sequence = _156___mcc_h6
			_ = _157_cf29
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))
			_ = _index18
			var _rhs24 bool = (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool)
			_ = _rhs24
			var _rhs25 bool = (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool)
			_ = _rhs25
			var _rhs26 _dafny.Int = (_dafny.Zero).Minus(_85_v2)
			_ = _rhs26
			var _lhs21 *GlobalState = globalState
			_ = _lhs21
			var _lhs22 _dafny.Array = _126_v39
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))
			_ = _lhs23
			var _lhs24 *GlobalState = globalState
			_ = _lhs24
			_lhs21.F22 = _rhs24
			(_lhs22).ArraySet1(_rhs25, (_lhs23).Int())
			_lhs24.F10 = _rhs26
			var _158_v56 _dafny.Sequence
			_ = _158_v56
			_158_v56 = _dafny.SeqOf(_84_v1)
			var _159_v57 _dafny.Array
			_ = _159_v57
			var _nwElement0_6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_158_v56).Select((Companion_Default___.SafeIndex(_131_i1, _dafny.IntOfUint32((_158_v56).Cardinality()))).Uint32()).(_dafny.Sequence), _84_v1)
			_ = _nwElement0_6
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(2))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_6, 0)
			(_nw17).ArraySet1(_dafny.SeqOf((_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool), _83_v0, (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool), (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool), _83_v0), 1)
			_159_v57 = _nw17
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_159_v57), 0))
			_ = _index19
			(_159_v57).ArraySet1(_84_v1, (_index19).Int())
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_159_v57), 0))
			_ = _index20
			(_159_v57).ArraySet1((func() _dafny.Sequence {
				if _83_v0 {
					return _dafny.SeqOf((_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool), (_126_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(546), _dafny.ArrayLen((_126_v39), 0))).Int()).(bool))
				}
				return _84_v1
			})(), (_index20).Int())
			(globalState).F10 = (_dafny.Zero).Minus(_85_v2)
			_85_v2 = _dafny.IntOfInt64(889)
		}
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _160_v0 _dafny.Array
	_ = _160_v0
	var _nw18 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
	_ = _nw18
	_160_v0 = _nw18
	var _161_v1 _dafny.MultiSet
	_ = _161_v1
	_161_v1 = _dafny.MultiSetOf(_160_v0, _160_v0)
	var _162_v2 bool
	_ = _162_v2
	_162_v2 = false
	var _163_v3 _dafny.Sequence
	_ = _163_v3
	_163_v3 = _dafny.SeqOf(_162_v2)
	var _164_v4 _dafny.Sequence
	_ = _164_v4
	_164_v4 = _dafny.SeqOf((func() _dafny.Int {
		if (_161_v1).Contains(_160_v0) {
			return (_161_v1).Multiplicity(_160_v0)
		}
		return _dafny.IntOfUint32((_163_v3).Cardinality())
	})())
	var _165_v5 _dafny.Map
	_ = _165_v5
	_165_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v4, _160_v0)
	var _166_v6 _dafny.Int
	_ = _166_v6
	_166_v6 = _dafny.IntOfInt64(-551)
	var _167_v7 _dafny.Array
	_ = _167_v7
	var _nwElement0_7 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_163_v3, (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_163_v3).Cardinality()))).Uint32(), _162_v2)
	_ = _nwElement0_7
	var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(2))
	_ = _nw19
	(_nw19).ArraySet1(_nwElement0_7, 0)
	(_nw19).ArraySet1(_163_v3, 1)
	_167_v7 = _nw19
	var _168_v8 _dafny.Sequence
	_ = _168_v8
	_168_v8 = _dafny.UnicodeSeqOfUtf8Bytes("wrohtasav")
	var _169_v9 _dafny.MultiSet
	_ = _169_v9
	_169_v9 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}((func(_170_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_171_i0 _dafny.Int) _dafny.Int {
			return _170_v6
		}
	})(_166_v6)))).Cardinality()))
	var _172_v10 _dafny.Map
	_ = _172_v10
	_172_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_168_v8).Select((Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32()).(_dafny.CodePoint), _169_v9)
	var _173_v11 _dafny.Array
	_ = _173_v11
	var _nwElement0_8 _dafny.MultiSet = (func() _dafny.MultiSet {
		if (_172_v10).Contains(_dafny.CodePoint('n')) {
			return (_172_v10).Get(_dafny.CodePoint('n')).(_dafny.MultiSet)
		}
		return _169_v9
	})()
	_ = _nwElement0_8
	var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
	_ = _nw20
	(_nw20).ArraySet1(_nwElement0_8, 0)
	(_nw20).ArraySet1(_169_v9, 1)
	(_nw20).ArraySet1(_169_v9, 2)
	(_nw20).ArraySet1(_dafny.MultiSetOf(_166_v6, _dafny.IntOfInt64(985)), 3)
	_173_v11 = _nw20
	var _174_v12 _dafny.Array
	_ = _174_v12
	var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
	_ = _nw21
	_174_v12 = _nw21
	var _175_v13 _dafny.CodePoint
	_ = _175_v13
	_175_v13 = _dafny.CodePoint('s')
	var _176_v14 _dafny.Map
	_ = _176_v14
	_176_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, _162_v2)
	var _177_v15 _dafny.Map
	_ = _177_v15
	_177_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_176_v14, _162_v2)
	var _178_v16 _dafny.Map
	_ = _178_v16
	_178_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v0, _162_v2)
	var _179_v17 _dafny.Sequence
	_ = _179_v17
	_179_v17 = _dafny.SeqOf(_168_v8)
	var _180_v18 _dafny.Array
	_ = _180_v18
	var _nwElement0_9 _dafny.Sequence = _168_v8
	_ = _nwElement0_9
	var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(18))
	_ = _nw22
	(_nw22).ArraySet1(_nwElement0_9, 0)
	(_nw22).ArraySet1(_168_v8, 1)
	(_nw22).ArraySet1(_168_v8, 2)
	(_nw22).ArraySet1(_168_v8, 3)
	(_nw22).ArraySet1(_168_v8, 4)
	(_nw22).ArraySet1(_dafny.Companion_Sequence_.Update(_168_v8, (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32(), _175_v13), 5)
	(_nw22).ArraySet1(_168_v8, 6)
	(_nw22).ArraySet1((_179_v17).Select((Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_179_v17).Cardinality()))).Uint32()).(_dafny.Sequence), 7)
	(_nw22).ArraySet1(_168_v8, 8)
	(_nw22).ArraySet1(_168_v8, 9)
	(_nw22).ArraySet1((Companion_D0_.Create_DC0_(_168_v8)).Dtor_cf0(), 10)
	(_nw22).ArraySet1(_168_v8, 11)
	(_nw22).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("vfacyd"), (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfacyd")).Cardinality()))).Uint32(), (_168_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_163_v3).Cardinality()), _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)), 12)
	(_nw22).ArraySet1(_168_v8, 13)
	(_nw22).ArraySet1(_168_v8, 14)
	(_nw22).ArraySet1(_168_v8, 15)
	(_nw22).ArraySet1(_168_v8, 16)
	(_nw22).ArraySet1(_dafny.Companion_Sequence_.Update(_168_v8, (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32(), _175_v13), 17)
	_180_v18 = _nw22
	var _181_v19 _dafny.Set
	_ = _181_v19
	_181_v19 = _dafny.SetOf(_dafny.IntOfUint32((_163_v3).Cardinality()))
	var _182_globalState *GlobalState
	_ = _182_globalState
	var _nw23 *GlobalState = New_GlobalState_()
	_ = _nw23
	_nw23.Ctor__((_165_v5).Merge(_165_v5), _167_v7, true, _dafny.CodePoint('h'), false, _173_v11, _174_v12, false, _169_v9, _dafny.IntOfInt64(974), _dafny.IntOfInt64(-767), _177_v15, _178_v16, true, true, true, false, _dafny.IntOfInt64(811), true, _dafny.CodePoint('d'), true, _180_v18, true, true, _181_v19, _dafny.Companion_Sequence_.Concatenate(_163_v3, _163_v3), true, _dafny.IntOfInt64(80), _dafny.CodePoint('x'))
	_182_globalState = _nw23
	var _183_v20 _dafny.Map
	_ = _183_v20
	_183_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_164_v4).Cardinality()), false)
	_183_v20 = _183_v20
	var _184_v21 D0
	_ = _184_v21
	_184_v21 = Companion_D0_.Create_DC1_(Companion_Default___.Fm0((_dafny.MultiSetFromSeq(_164_v4)).Cardinality(), _182_globalState), _166_v6, _166_v6, _160_v0)
	var _pat_let_tv0 = _184_v21
	_ = _pat_let_tv0
	var _source5 D0 = func(_pat_let0_0 D0) D0 {
		return func(_185_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 D0) D0 {
				return func(_186_dt__update_hcf7_h0 D0) D0 {
					return Companion_D0_.Create_DC3_(_186_dt__update_hcf7_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(Companion_D0_.Create_DC3_(_184_v21))
	_ = _source5
	if _source5.Is_DC1() {
		var _187___mcc_h0 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
		_ = _187___mcc_h0
		var _188___mcc_h1 _dafny.Int = _source5.Get_().(D0_DC1).Cf2
		_ = _188___mcc_h1
		var _189___mcc_h2 _dafny.Int = _source5.Get_().(D0_DC1).Cf3
		_ = _189___mcc_h2
		var _190___mcc_h3 _dafny.Array = _source5.Get_().(D0_DC1).Cf4
		_ = _190___mcc_h3
		var _191_cf4 _dafny.Array = _190___mcc_h3
		_ = _191_cf4
		var _192_cf3 _dafny.Int = _189___mcc_h2
		_ = _192_cf3
		var _193_cf2 _dafny.Int = _188___mcc_h1
		_ = _193_cf2
		var _194_cf1 _dafny.Int = _187___mcc_h0
		_ = _194_cf1
		(_182_globalState).F7 = (_192_cf3).Cmp(_192_cf3) >= 0
		var _195_v22 D0
		_ = _195_v22
		_195_v22 = Companion_D0_.Create_DC2_(_193_cf2, _162_v2)
		_195_v22 = _195_v22
		var _196_v23 *C1
		_ = _196_v23
		var _nw24 *C1 = New_C1_()
		_ = _nw24
		_nw24.Ctor__()
		_196_v23 = _nw24
		var _197_v24 D0
		_ = _197_v24
		var _198_v25 bool
		_ = _198_v25
		var _out0 D0
		_ = _out0
		var _out1 bool
		_ = _out1
		_out0, _out1 = (_196_v23).M1(_dafny.Companion_Sequence_.Update(_164_v4, (Companion_Default___.SafeIndex(_194_cf1, _dafny.IntOfUint32((_164_v4).Cardinality()))).Uint32(), _194_cf1), (func() _dafny.Int {
			if _162_v2 {
				return _194_cf1
			}
			return _194_cf1
		})(), _162_v2, _182_globalState)
		_197_v24 = _out0
		_198_v25 = _out1
	} else if _source5.Is_DC2() {
		var _199___mcc_h4 _dafny.Int = _source5.Get_().(D0_DC2).Cf5
		_ = _199___mcc_h4
		var _200___mcc_h5 bool = _source5.Get_().(D0_DC2).Cf6
		_ = _200___mcc_h5
		var _201_cf6 bool = _200___mcc_h5
		_ = _201_cf6
		var _202_cf5 _dafny.Int = _199___mcc_h4
		_ = _202_cf5
		var _203_v26 _dafny.Map
		_ = _203_v26
		_203_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _201_cf6)
		_183_v20 = (_183_v20).Update(Companion_Default___.SafeModuloInt((_183_v20).Cardinality(), _202_cf5), !((_202_cf5).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(!((func() bool {
			if (_203_v26).Contains(_162_v2) {
				return (_203_v26).Get(_162_v2).(bool)
			}
			return _162_v2
		})()), _201_cf6), (Companion_Default___.SafeIndex(_202_cf5, _dafny.IntOfUint32((_dafny.SeqOf(!((func() bool {
			if (_203_v26).Contains(_162_v2) {
				return (_203_v26).Get(_162_v2).(bool)
			}
			return _162_v2
		})()), _201_cf6)).Cardinality()))).Uint32(), _201_cf6), (Companion_Default___.SafeIndex(_202_cf5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(!((func() bool {
			if (_203_v26).Contains(_162_v2) {
				return (_203_v26).Get(_162_v2).(bool)
			}
			return _162_v2
		})()), _201_cf6), (Companion_Default___.SafeIndex(_202_cf5, _dafny.IntOfUint32((_dafny.SeqOf(!((func() bool {
			if (_203_v26).Contains(_162_v2) {
				return (_203_v26).Get(_162_v2).(bool)
			}
			return _162_v2
		})()), _201_cf6)).Cardinality()))).Uint32(), _201_cf6)).Cardinality()))).Uint32(), _162_v2)).Cardinality())) >= 0))
		Companion_Default___.M8(_182_globalState)
		var _204_v27 _dafny.Array
		_ = _204_v27
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_1
		var _nw25 _dafny.Array
		_ = _nw25
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw25 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) _dafny.Int = func(_205_i1 _dafny.Int) _dafny.Int {
				return (_205_i1).Times(_dafny.IntOfInt64(-936))
			}
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw25 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw25).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw25).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_204_v27 = _nw25
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_204_v27), 0))
		_ = _index21
		(_204_v27).ArraySet1(_166_v6, (_index21).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_204_v27), 0))
		_ = _index22
		(_204_v27).ArraySet1(_202_cf5, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_160_v0), 0))
		_ = _index23
		(_160_v0).ArraySet1((_163_v3).Select((Companion_Default___.SafeIndex((_204_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_204_v27), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_163_v3).Cardinality()))).Uint32()).(bool), (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_160_v0), 0))
		_ = _index24
		(_160_v0).ArraySet1(_201_cf6, (_index24).Int())
	} else if _source5.Is_DC0() {
		var _206___mcc_h6 _dafny.Sequence = _source5.Get_().(D0_DC0).Cf0
		_ = _206___mcc_h6
		var _207_cf0 _dafny.Sequence = _206___mcc_h6
		_ = _207_cf0
		Companion_Default___.M8(_182_globalState)
		var _208_v28 _dafny.Array
		_ = _208_v28
		var _nwElement0_10 _dafny.CodePoint = _175_v13
		_ = _nwElement0_10
		var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(21))
		_ = _nw26
		(_nw26).ArraySet1CodePoint(_nwElement0_10, 0)
		(_nw26).ArraySet1CodePoint(_175_v13, 1)
		(_nw26).ArraySet1CodePoint(_dafny.CodePoint('j'), 2)
		(_nw26).ArraySet1CodePoint(_175_v13, 3)
		(_nw26).ArraySet1CodePoint(_175_v13, 4)
		(_nw26).ArraySet1CodePoint((_207_cf0).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_166_v6), _dafny.IntOfUint32((_207_cf0).Cardinality()))).Uint32()).(_dafny.CodePoint), 5)
		(_nw26).ArraySet1CodePoint(_dafny.CodePoint('d'), 6)
		(_nw26).ArraySet1CodePoint(_175_v13, 7)
		(_nw26).ArraySet1CodePoint(_175_v13, 8)
		(_nw26).ArraySet1CodePoint(_175_v13, 9)
		(_nw26).ArraySet1CodePoint(_dafny.CodePoint('p'), 10)
		(_nw26).ArraySet1CodePoint((_168_v8).Select((Companion_Default___.SafeIndex((_169_v9).Cardinality(), _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32()).(_dafny.CodePoint), 11)
		(_nw26).ArraySet1CodePoint(_175_v13, 12)
		(_nw26).ArraySet1CodePoint(_175_v13, 13)
		(_nw26).ArraySet1CodePoint(_175_v13, 14)
		(_nw26).ArraySet1CodePoint(_175_v13, 15)
		(_nw26).ArraySet1CodePoint(_175_v13, 16)
		(_nw26).ArraySet1CodePoint(_175_v13, 17)
		(_nw26).ArraySet1CodePoint(_175_v13, 18)
		(_nw26).ArraySet1CodePoint(_175_v13, 19)
		(_nw26).ArraySet1CodePoint(_175_v13, 20)
		_208_v28 = _nw26
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_208_v28), 0))
		_ = _index25
		(_208_v28).ArraySet1CodePoint(_dafny.CodePoint('l'), (_index25).Int())
		var _209_v29 _dafny.Map
		_ = _209_v29
		_209_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_160_v0, _175_v13)
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(472), _dafny.ArrayLen((_208_v28), 0))
		_ = _index26
		(_208_v28).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _162_v2 {
				return _175_v13
			}
			return (func() _dafny.CodePoint {
				if (_209_v29).Contains(_160_v0) {
					return (_209_v29).Get(_160_v0).(_dafny.CodePoint)
				}
				return _175_v13
			})()
		})(), (_index26).Int())
		(_182_globalState).F10 = (_166_v6).Times((_dafny.Zero).Minus((_166_v6).Plus(_166_v6)))
		var _210_v30 D0
		_ = _210_v30
		_210_v30 = Companion_D0_.Create_DC3_((func() D0 {
			if _162_v2 {
				return _184_v21
			}
			return _184_v21
		})())
		var _source6 D0 = _210_v30
		_ = _source6
		if _source6.Is_DC1() {
			var _211___mcc_h8 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
			_ = _211___mcc_h8
			var _212___mcc_h9 _dafny.Int = _source6.Get_().(D0_DC1).Cf2
			_ = _212___mcc_h9
			var _213___mcc_h10 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
			_ = _213___mcc_h10
			var _214___mcc_h11 _dafny.Array = _source6.Get_().(D0_DC1).Cf4
			_ = _214___mcc_h11
			var _215_cf4 _dafny.Array = _214___mcc_h11
			_ = _215_cf4
			var _216_cf3 _dafny.Int = _213___mcc_h10
			_ = _216_cf3
			var _217_cf2 _dafny.Int = _212___mcc_h9
			_ = _217_cf2
			var _218_cf1 _dafny.Int = _211___mcc_h8
			_ = _218_cf1
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_215_cf4), 0))
			_ = _index27
			(_215_cf4).ArraySet1(Companion_Default___.Fm12(_162_v2, _162_v2, _182_globalState), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_215_cf4), 0))
			_ = _index28
			(_215_cf4).ArraySet1(_162_v2, (_index28).Int())
			_216_cf3 = _216_cf3
			Companion_Default___.M8(_182_globalState)
			_207_cf0 = _207_cf0
		} else if _source6.Is_DC2() {
			var _219___mcc_h12 _dafny.Int = _source6.Get_().(D0_DC2).Cf5
			_ = _219___mcc_h12
			var _220___mcc_h13 bool = _source6.Get_().(D0_DC2).Cf6
			_ = _220___mcc_h13
			var _221_cf6 bool = _220___mcc_h13
			_ = _221_cf6
			var _222_cf5 _dafny.Int = _219___mcc_h12
			_ = _222_cf5
			var _223_v31 _dafny.Array
			_ = _223_v31
			var _nw27 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw27
			_223_v31 = _nw27
			var _224_v32 _dafny.Array
			_ = _224_v32
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw28
			_224_v32 = _nw28
			var _225_v33 T1
			_ = _225_v33
			var _nw29 *C4 = New_C4_()
			_ = _nw29
			_nw29.Ctor__(_224_v32)
			_225_v33 = _nw29
			var _226_v34 _dafny.Map
			_ = _226_v34
			_226_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_223_v31, _225_v33)
			_226_v34 = (_226_v34).Update(_223_v31, _225_v33)
			var _227_v35 _dafny.Array
			_ = _227_v35
			var _nwElement0_11 _dafny.Int = _222_cf5
			_ = _nwElement0_11
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(4))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_11, 0)
			(_nw30).ArraySet1(_166_v6, 1)
			(_nw30).ArraySet1((_dafny.Zero).Minus(_222_cf5), 2)
			(_nw30).ArraySet1(_166_v6, 3)
			_227_v35 = _nw30
			_227_v35 = (func() _dafny.Array {
				if _221_cf6 {
					return _227_v35
				}
				return _227_v35
			})()
			var _228_v36 _dafny.Array
			_ = _228_v36
			_228_v36 = _208_v28
			var _rhs27 _dafny.Array = _227_v35
			_ = _rhs27
			var _rhs28 bool = (_166_v6).Cmp(_222_cf5) != 0
			_ = _rhs28
			var _rhs29 _dafny.Array = _208_v28
			_ = _rhs29
			var _lhs25 *GlobalState = _182_globalState
			_ = _lhs25
			_227_v35 = _rhs27
			_lhs25.F7 = _rhs28
			_228_v36 = _rhs29
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_227_v35), 0))
			_ = _index29
			(_227_v35).ArraySet1(_222_cf5, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(289), _dafny.ArrayLen((_227_v35), 0))
			_ = _index30
			(_227_v35).ArraySet1(_222_cf5, (_index30).Int())
		} else if _source6.Is_DC0() {
			var _229___mcc_h14 _dafny.Sequence = _source6.Get_().(D0_DC0).Cf0
			_ = _229___mcc_h14
			var _230_cf0 _dafny.Sequence = _229___mcc_h14
			_ = _230_cf0
			var _231_v37 _dafny.Sequence
			_ = _231_v37
			_231_v37 = _dafny.SeqOf(_183_v20)
			_231_v37 = (func() _dafny.Sequence {
				if _162_v2 {
					return _dafny.Companion_Sequence_.Concatenate(_231_v37, _dafny.SeqOf(_183_v20))
				}
				return _dafny.SeqOf(_183_v20)
			})()
			var _232_v38 _dafny.Array
			_ = _232_v38
			var _nw31 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
			_ = _nw31
			_232_v38 = _nw31
			var _233_v39 _dafny.Sequence
			_ = _233_v39
			_233_v39 = _dafny.SeqOf(_232_v38, _232_v38, _232_v38, _232_v38, _232_v38)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_232_v38), 0))
			_ = _index31
			(_232_v38).ArraySet1(_dafny.IntOfUint32((_233_v39).Cardinality()), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_232_v38), 0))
			_ = _index32
			(_232_v38).ArraySet1((_164_v4).Select((Companion_Default___.SafeIndex((_dafny.IntOfInt64(640)).Minus((_181_v19).Cardinality()), _dafny.IntOfUint32((_164_v4).Cardinality()))).Uint32()).(_dafny.Int), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_232_v38), 0))
			_ = _index33
			(_232_v38).ArraySet1(_166_v6, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_160_v0), 0))
			_ = _index34
			(_160_v0).ArraySet1(Companion_Default___.Fm12(false, false, _182_globalState), (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_160_v0), 0))
			_ = _index35
			(_160_v0).ArraySet1(!(!_dafny.Companion_Sequence_.Contains(_179_v17, _dafny.UnicodeSeqOfUtf8Bytes("gr"))), (_index35).Int())
		} else {
			var _234___mcc_h15 D0 = _source6.Get_().(D0_DC3).Cf7
			_ = _234___mcc_h15
			var _235_cf7 D0 = _234___mcc_h15
			_ = _235_cf7
			_166_v6 = _166_v6
			(_182_globalState).F10 = (_164_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_166_v6), _dafny.IntOfUint32((_164_v4).Cardinality()))).Uint32()).(_dafny.Int)
			Companion_Default___.M8(_182_globalState)
			var _236_v40 _dafny.Map
			_ = _236_v40
			_236_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _208_v28)
			_236_v40 = (_236_v40).Update(_162_v2, _208_v28)
		}
	} else {
		var _237___mcc_h7 D0 = _source5.Get_().(D0_DC3).Cf7
		_ = _237___mcc_h7
		var _238_cf7 D0 = _237___mcc_h7
		_ = _238_cf7
		_167_v7 = _167_v7
		if (_169_v9).IsSubsetOf(_dafny.MultiSetOf(_166_v6)) {
			var _239_v41 *C0
			_ = _239_v41
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(!(!(Companion_Default___.Fm12(_162_v2, _162_v2, _182_globalState)) || (_162_v2)))
			_239_v41 = _nw32
			(_182_globalState).F10 = Companion_Default___.Fm0(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_166_v6), _166_v6), _182_globalState)
			(_182_globalState).F7 = (_239_v41).F32()
			var _240_v42 _dafny.Array
			_ = _240_v42
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_2
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) _dafny.Int = (func(_241_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_242_i2 _dafny.Int) _dafny.Int {
						return (_242_i2).Times(_241_v6)
					}
				})(_166_v6)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw33 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw33).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw33).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_240_v42 = _nw33
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_240_v42), 0))
			_ = _index36
			(_240_v42).ArraySet1(_166_v6, (_index36).Int())
			var _243_v44 _dafny.Sequence
			_ = _243_v44
			_243_v44 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_166_v6, _182_globalState), true), func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(599), _dafny.IntOfInt64(-29))); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _244_v43 _dafny.Int
					_244_v43 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(599)).Cmp(_244_v43) <= 0) && ((_244_v43).Cmp(_dafny.IntOfInt64(-29)) < 0) {
						_coll15.Add(Companion_Default___.SafeDivisionInt(_244_v43, _dafny.IntOfUint32((_168_v8).Cardinality())), _162_v2)
					}
				}
				return _coll15.ToMap()
			}())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_240_v42), 0))
			_ = _index37
			(_240_v42).ArraySet1((Companion_Default___.SafeDivisionInt(_166_v6, _166_v6)).Plus((_166_v6).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_243_v44).Cardinality()))).Cardinality())))), (_index37).Int())
			(_182_globalState).F22 = (_239_v41).F32()
		} else {
			var _245_v45 _dafny.Array
			_ = _245_v45
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw34
			_245_v45 = _nw34
			_245_v45 = _245_v45
			var _246_v46 *C1
			_ = _246_v46
			var _nw35 *C1 = New_C1_()
			_ = _nw35
			_nw35.Ctor__()
			_246_v46 = _nw35
			var _247_v47 _dafny.MultiSet
			_ = _247_v47
			_247_v47 = _dafny.MultiSetOf(_162_v2)
			var _248_v48 _dafny.Map
			_ = _248_v48
			_248_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_247_v47, _246_v46)
			var _249_v49 _dafny.Map
			_ = _249_v49
			_249_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_168_v8).Cardinality()), _246_v46)
			var _250_v50 _dafny.Array
			_ = _250_v50
			var _nwElement0_12 *C1 = _246_v46
			_ = _nwElement0_12
			var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(26))
			_ = _nw36
			(_nw36).ArraySet1(_nwElement0_12, 0)
			(_nw36).ArraySet1(_246_v46, 1)
			(_nw36).ArraySet1(_246_v46, 2)
			(_nw36).ArraySet1(_246_v46, 3)
			(_nw36).ArraySet1(_246_v46, 4)
			(_nw36).ArraySet1(_246_v46, 5)
			(_nw36).ArraySet1(_246_v46, 6)
			(_nw36).ArraySet1(_246_v46, 7)
			(_nw36).ArraySet1(_246_v46, 8)
			(_nw36).ArraySet1(_246_v46, 9)
			(_nw36).ArraySet1(_246_v46, 10)
			(_nw36).ArraySet1((func() *C1 {
				if (_248_v48).Contains(_247_v47) {
					return (_248_v48).Get(_247_v47).(*C1)
				}
				return _246_v46
			})(), 11)
			(_nw36).ArraySet1(_246_v46, 12)
			(_nw36).ArraySet1(_246_v46, 13)
			(_nw36).ArraySet1(_246_v46, 14)
			(_nw36).ArraySet1(_246_v46, 15)
			(_nw36).ArraySet1(_246_v46, 16)
			(_nw36).ArraySet1(_246_v46, 17)
			(_nw36).ArraySet1(_246_v46, 18)
			(_nw36).ArraySet1(_246_v46, 19)
			(_nw36).ArraySet1(_246_v46, 20)
			(_nw36).ArraySet1(_246_v46, 21)
			(_nw36).ArraySet1(_246_v46, 22)
			(_nw36).ArraySet1(_246_v46, 23)
			(_nw36).ArraySet1((func() *C1 {
				if (_249_v49).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_162_v2, _162_v2, _162_v2)).Cardinality())) {
					return (_249_v49).Get(_dafny.IntOfUint32((_dafny.SeqOf(_162_v2, _162_v2, _162_v2)).Cardinality())).(*C1)
				}
				return _246_v46
			})(), 24)
			(_nw36).ArraySet1(_246_v46, 25)
			_250_v50 = _nw36
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_250_v50), 0))
			_ = _index38
			(_250_v50).ArraySet1(_246_v46, (_index38).Int())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_250_v50), 0))
			_ = _index39
			(_250_v50).ArraySet1(_246_v46, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_250_v50), 0))
			_ = _index40
			(_250_v50).ArraySet1((func() *C1 {
				if _162_v2 {
					return _246_v46
				}
				return _246_v46
			})(), (_index40).Int())
			Companion_Default___.M8(_182_globalState)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_160_v0), 0))
			_ = _index41
			(_160_v0).ArraySet1(_162_v2, (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_160_v0), 0))
			_ = _index42
			(_160_v0).ArraySet1(!(false), (_index42).Int())
		}
		var _251_v51 _dafny.Array
		_ = _251_v51
		var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw37
		_251_v51 = _nw37
		var _252_v52 _dafny.Map
		_ = _252_v52
		_252_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, _251_v51)
		var _253_v53 T1
		_ = _253_v53
		var _nw38 *C4 = New_C4_()
		_ = _nw38
		_nw38.Ctor__((func() _dafny.Array {
			if (_252_v52).Contains(_166_v6) {
				return (_252_v52).Get(_166_v6).(_dafny.Array)
			}
			return _251_v51
		})())
		_253_v53 = _nw38
		var _rhs30 bool = !(!(!_dafny.Companion_Sequence_.Equal(_179_v17, _179_v17))) || (_162_v2)
		_ = _rhs30
		var _rhs31 T1 = _253_v53
		_ = _rhs31
		var _lhs26 *GlobalState = _182_globalState
		_ = _lhs26
		_lhs26.F7 = _rhs30
		_253_v53 = _rhs31
		var _rhs32 bool = (_162_v2) == (_162_v2)
		_ = _rhs32
		var _rhs33 _dafny.CodePoint = _175_v13
		_ = _rhs33
		var _lhs27 *GlobalState = _182_globalState
		_ = _lhs27
		var _lhs28 *GlobalState = _182_globalState
		_ = _lhs28
		_lhs27.F13 = _rhs32
		_lhs28.F19 = _rhs33
	}
	var _254_v54 _dafny.Set
	_ = _254_v54
	_254_v54 = _dafny.SetOf(_162_v2)
	var _source7 D8 = Companion_D8_.Create_DC22_((_254_v54).Cardinality(), _162_v2, _175_v13)
	_ = _source7
	if _source7.Is_DC22() {
		var _255___mcc_h16 _dafny.Int = _source7.Get_().(D8_DC22).Cf36
		_ = _255___mcc_h16
		var _256___mcc_h17 bool = _source7.Get_().(D8_DC22).Cf37
		_ = _256___mcc_h17
		var _257___mcc_h18 _dafny.CodePoint = _source7.Get_().(D8_DC22).Cf38
		_ = _257___mcc_h18
		var _258_cf38 _dafny.CodePoint = _257___mcc_h18
		_ = _258_cf38
		var _259_cf37 bool = _256___mcc_h17
		_ = _259_cf37
		var _260_cf36 _dafny.Int = _255___mcc_h16
		_ = _260_cf36
		var _261_v55 _dafny.Array
		_ = _261_v55
		var _nw39 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw39
		_261_v55 = _nw39
		var _262_v56 _dafny.Sequence
		_ = _262_v56
		_262_v56 = _dafny.SeqOf(_261_v55)
		(_182_globalState).F22 = (_166_v6).Cmp(_dafny.IntOfUint32((_262_v56).Cardinality())) > 0
		_259_cf37 = _dafny.Companion_Sequence_.IsPrefixOf((func() _dafny.Sequence {
			if !(_259_cf37) {
				return _164_v4
			}
			return _164_v4
		})(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(251))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_263_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_264_i3 _dafny.Int) _dafny.Int {
				return _263_v6
			}
		})(_166_v6))), _164_v4))
		var _265_v57 *C1
		_ = _265_v57
		var _nw40 *C1 = New_C1_()
		_ = _nw40
		_nw40.Ctor__()
		_265_v57 = _nw40
		(_182_globalState).F10 = Companion_Default___.Fm0(_166_v6, _182_globalState)
	} else if _source7.Is_DC23() {
		var _266___mcc_h19 bool = _source7.Get_().(D8_DC23).Cf39
		_ = _266___mcc_h19
		var _267_cf39 bool = _266___mcc_h19
		_ = _267_cf39
		var _268_v58 _dafny.Sequence
		_ = _268_v58
		_268_v58 = _dafny.SeqOf(_160_v0)
		var _269_v59 _dafny.Array
		_ = _269_v59
		var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw41
		_269_v59 = _nw41
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_269_v59), 0))
		_ = _index43
		(_269_v59).ArraySet1(_166_v6, (_index43).Int())
		var _270_v60 _dafny.Array
		_ = _270_v60
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
		_ = _nw42
		_270_v60 = _nw42
		var _271_v61 T0
		_ = _271_v61
		var _nw43 *C4 = New_C4_()
		_ = _nw43
		_nw43.Ctor__(_270_v60)
		_271_v61 = _nw43
		var _272_v62 _dafny.Sequence
		_ = _272_v62
		_272_v62 = _dafny.SeqOf(_271_v61)
		var _273_v63 _dafny.Map
		_ = _273_v63
		_273_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_166_v6, (_dafny.Zero).Minus(_166_v6)), (_272_v62).Select((Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_272_v62).Cardinality()))).Uint32()).(T0))
		var _274_v64 _dafny.Map
		_ = _274_v64
		_274_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _162_v2)
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_269_v59), 0))
		_ = _index44
		var _rhs34 _dafny.Sequence = _268_v58
		_ = _rhs34
		var _rhs35 _dafny.Int = (_166_v6).Times(((_274_v64).Cardinality()).Times(_166_v6))
		_ = _rhs35
		var _rhs36 _dafny.Map = (_273_v63).Merge((_273_v63).Merge(_273_v63))
		_ = _rhs36
		var _lhs29 _dafny.Array = _269_v59
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_269_v59), 0))
		_ = _lhs30
		_268_v58 = _rhs34
		(_lhs29).ArraySet1(_rhs35, (_lhs30).Int())
		_273_v63 = _rhs36
		var _275_v65 _dafny.Map
		_ = _275_v65
		_275_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_269_v59, _162_v2)
		var _276_v66 _dafny.Map
		_ = _276_v66
		_276_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _175_v13)
		var _277_v67 _dafny.Array
		_ = _277_v67
		var _nwElement0_13 _dafny.CodePoint = _175_v13
		_ = _nwElement0_13
		var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(16))
		_ = _nw44
		(_nw44).ArraySet1CodePoint(_nwElement0_13, 0)
		(_nw44).ArraySet1CodePoint(_175_v13, 1)
		(_nw44).ArraySet1CodePoint(_175_v13, 2)
		(_nw44).ArraySet1CodePoint(_175_v13, 3)
		(_nw44).ArraySet1CodePoint(_175_v13, 4)
		(_nw44).ArraySet1CodePoint(_dafny.CodePoint('l'), 5)
		(_nw44).ArraySet1CodePoint(_175_v13, 6)
		(_nw44).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _267_cf39 {
				return _175_v13
			}
			return _175_v13
		})(), 7)
		(_nw44).ArraySet1CodePoint(_175_v13, 8)
		(_nw44).ArraySet1CodePoint(_dafny.CodePoint('n'), 9)
		(_nw44).ArraySet1CodePoint(_175_v13, 10)
		(_nw44).ArraySet1CodePoint(Companion_Default___.Fm11(((_275_v65).Update(_269_v59, _267_cf39)).Cardinality(), _182_globalState), 11)
		(_nw44).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_276_v66).Contains(_267_cf39) {
				return (_276_v66).Get(_267_cf39).(_dafny.CodePoint)
			}
			return _175_v13
		})(), 12)
		(_nw44).ArraySet1CodePoint(_175_v13, 13)
		(_nw44).ArraySet1CodePoint(_175_v13, 14)
		(_nw44).ArraySet1CodePoint(_175_v13, 15)
		_277_v67 = _nw44
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_277_v67), 0))
		_ = _index45
		(_277_v67).ArraySet1CodePoint(_175_v13, (_index45).Int())
		var _278_v68 D8
		_ = _278_v68
		_278_v68 = Companion_D8_.Create_DC22_(_dafny.IntOfInt64(245), _162_v2, _175_v13)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_277_v67), 0))
		_ = _index46
		(_277_v67).ArraySet1CodePoint((func() _dafny.CodePoint {
			if _267_cf39 {
				return _175_v13
			}
			return (_278_v68).Dtor_cf38()
		})(), (_index46).Int())
		var _279_v69 _dafny.Map
		_ = _279_v69
		_279_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_269_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_269_v59), 0))).Int()).(_dafny.Int), _164_v4)
		_279_v69 = (_279_v69).Update(_166_v6, _164_v4)
		var _280_v70 *C3
		_ = _280_v70
		var _nw45 *C3 = New_C3_()
		_ = _nw45
		_nw45.Ctor__()
		_280_v70 = _nw45
		var _281_v71 _dafny.Sequence
		_ = _281_v71
		_281_v71 = _dafny.SeqOf(_280_v70)
		_166_v6 = (_dafny.Zero).Minus((_166_v6).Times((_dafny.IntOfUint32((_281_v71).Cardinality())).Minus(_166_v6)))
	} else if _source7.Is_DC21() {
		var _282___mcc_h20 _dafny.Sequence = _source7.Get_().(D8_DC21).Cf35
		_ = _282___mcc_h20
		var _283_cf35 _dafny.Sequence = _282___mcc_h20
		_ = _283_cf35
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_160_v0), 0))
		_ = _index47
		(_160_v0).ArraySet1(_162_v2, (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_160_v0), 0))
		_ = _index48
		(_160_v0).ArraySet1(!(_162_v2), (_index48).Int())
		_254_v54 = _254_v54
		(_182_globalState).F10 = Companion_Default___.SafeModuloInt(_166_v6, (_166_v6).Minus(_166_v6))
		var _284_v72 _dafny.Map
		_ = _284_v72
		_284_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, _166_v6)
		var _285_v73 _dafny.MultiSet
		_ = _285_v73
		_285_v73 = _dafny.MultiSetOf(_175_v13)
		(_182_globalState).F10 = (func() _dafny.Int {
			if (_284_v72).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus((_285_v73).Cardinality()))) {
				return (_284_v72).Get((_dafny.Zero).Minus((_dafny.Zero).Minus((_285_v73).Cardinality()))).(_dafny.Int)
			}
			return _166_v6
		})()
	} else {
		var _286___mcc_h21 D8 = _source7.Get_().(D8_DC24).Cf40
		_ = _286___mcc_h21
		var _287_cf40 D8 = _286___mcc_h21
		_ = _287_cf40
		var _288_v74 _dafny.Array
		_ = _288_v74
		var _nwElement0_14 _dafny.CodePoint = _175_v13
		_ = _nwElement0_14
		var _nw46 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(23))
		_ = _nw46
		(_nw46).ArraySet1CodePoint(_nwElement0_14, 0)
		(_nw46).ArraySet1CodePoint(_175_v13, 1)
		(_nw46).ArraySet1CodePoint(Companion_Default___.Fm11(_166_v6, _182_globalState), 2)
		(_nw46).ArraySet1CodePoint(_175_v13, 3)
		(_nw46).ArraySet1CodePoint(_175_v13, 4)
		(_nw46).ArraySet1CodePoint(_175_v13, 5)
		(_nw46).ArraySet1CodePoint(_175_v13, 6)
		(_nw46).ArraySet1CodePoint(_175_v13, 7)
		(_nw46).ArraySet1CodePoint(_175_v13, 8)
		(_nw46).ArraySet1CodePoint(_175_v13, 9)
		(_nw46).ArraySet1CodePoint(_175_v13, 10)
		(_nw46).ArraySet1CodePoint(_175_v13, 11)
		(_nw46).ArraySet1CodePoint(_175_v13, 12)
		(_nw46).ArraySet1CodePoint(_175_v13, 13)
		(_nw46).ArraySet1CodePoint(_175_v13, 14)
		(_nw46).ArraySet1CodePoint(_175_v13, 15)
		(_nw46).ArraySet1CodePoint(_175_v13, 16)
		(_nw46).ArraySet1CodePoint(_175_v13, 17)
		(_nw46).ArraySet1CodePoint(_175_v13, 18)
		(_nw46).ArraySet1CodePoint(_175_v13, 19)
		(_nw46).ArraySet1CodePoint(_175_v13, 20)
		(_nw46).ArraySet1CodePoint(_175_v13, 21)
		(_nw46).ArraySet1CodePoint(_175_v13, 22)
		_288_v74 = _nw46
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))
		_ = _index49
		(_288_v74).ArraySet1CodePoint(_175_v13, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))
		_ = _index50
		(_288_v74).ArraySet1CodePoint(_dafny.CodePoint('n'), (_index50).Int())
		var _289_v75 D4
		_ = _289_v75
		_289_v75 = Companion_D4_.Create_DC14_(_162_v2)
		var _290_v76 _dafny.Sequence
		_ = _290_v76
		_290_v76 = _dafny.SeqOf(_164_v4, _164_v4)
		var _291_v77 D0
		_ = _291_v77
		_291_v77 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_290_v76).Cardinality()), _162_v2)
		var _source8 D0 = (func() D0 {
			if (_289_v75).Dtor_cf25() {
				return Companion_D0_.Create_DC2_(_dafny.IntOfInt64(334), !(_162_v2))
			}
			return _291_v77
		})()
		_ = _source8
		if _source8.Is_DC1() {
			var _292___mcc_h22 _dafny.Int = _source8.Get_().(D0_DC1).Cf1
			_ = _292___mcc_h22
			var _293___mcc_h23 _dafny.Int = _source8.Get_().(D0_DC1).Cf2
			_ = _293___mcc_h23
			var _294___mcc_h24 _dafny.Int = _source8.Get_().(D0_DC1).Cf3
			_ = _294___mcc_h24
			var _295___mcc_h25 _dafny.Array = _source8.Get_().(D0_DC1).Cf4
			_ = _295___mcc_h25
			var _296_cf4 _dafny.Array = _295___mcc_h25
			_ = _296_cf4
			var _297_cf3 _dafny.Int = _294___mcc_h24
			_ = _297_cf3
			var _298_cf2 _dafny.Int = _293___mcc_h23
			_ = _298_cf2
			var _299_cf1 _dafny.Int = _292___mcc_h22
			_ = _299_cf1
			(_182_globalState).F13 = !((Companion_Default___.Fm12(false, _162_v2, _182_globalState)) == (!(_162_v2)))
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))
			_ = _index51
			(_288_v74).ArraySet1CodePoint((_288_v74).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))).Int()), (_index51).Int())
			Companion_Default___.M8(_182_globalState)
			Companion_Default___.M8(_182_globalState)
		} else if _source8.Is_DC2() {
			var _300___mcc_h26 _dafny.Int = _source8.Get_().(D0_DC2).Cf5
			_ = _300___mcc_h26
			var _301___mcc_h27 bool = _source8.Get_().(D0_DC2).Cf6
			_ = _301___mcc_h27
			var _302_cf6 bool = _301___mcc_h27
			_ = _302_cf6
			var _303_cf5 _dafny.Int = _300___mcc_h26
			_ = _303_cf5
			Companion_Default___.M8(_182_globalState)
			(_182_globalState).F7 = _162_v2
			var _304_v78 *C1
			_ = _304_v78
			var _nw47 *C1 = New_C1_()
			_ = _nw47
			_nw47.Ctor__()
			_304_v78 = _nw47
			_304_v78 = _304_v78
			var _305_v79 _dafny.Array
			_ = _305_v79
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(15))
			_ = _nw48
			_305_v79 = _nw48
			var _306_v80 _dafny.MultiSet
			_ = _306_v80
			_306_v80 = _dafny.MultiSetOf(_302_cf6)
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_305_v79), 0))
			_ = _index52
			(_305_v79).ArraySet1(Companion_D3_.Create_DC10_(_306_v80), (_index52).Int())
			var _307_v81 D3
			_ = _307_v81
			_307_v81 = Companion_D3_.Create_DC10_(_306_v80)
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_305_v79), 0))
			_ = _index53
			(_305_v79).ArraySet1(_307_v81, (_index53).Int())
		} else if _source8.Is_DC0() {
			var _308___mcc_h28 _dafny.Sequence = _source8.Get_().(D0_DC0).Cf0
			_ = _308___mcc_h28
			var _309_cf0 _dafny.Sequence = _308___mcc_h28
			_ = _309_cf0
			var _310_v82 *C0
			_ = _310_v82
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__(false)
			_310_v82 = _nw49
			var _311_v83 _dafny.Map
			_ = _311_v83
			_311_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_310_v82, _166_v6)
			_311_v83 = (_311_v83).Update(_310_v82, _166_v6)
			var _312_v84 _dafny.Map
			_ = _312_v84
			_312_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _162_v2)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_160_v0), 0))
			_ = _index54
			(_160_v0).ArraySet1((func() bool {
				if (_312_v84).Contains(Companion_Default___.Fm12((_310_v82).F32(), (func() bool {
					if (_183_v20).Contains(_166_v6) {
						return (_183_v20).Get(_166_v6).(bool)
					}
					return _162_v2
				})(), _182_globalState)) {
					return (_312_v84).Get(Companion_Default___.Fm12((_310_v82).F32(), (func() bool {
						if (_183_v20).Contains(_166_v6) {
							return (_183_v20).Get(_166_v6).(bool)
						}
						return _162_v2
					})(), _182_globalState)).(bool)
				}
				return !(_162_v2)
			})(), (_index54).Int())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_160_v0), 0))
			_ = _index55
			(_160_v0).ArraySet1(_162_v2, (_index55).Int())
			_169_v9 = _169_v9
			(_182_globalState).F10 = Companion_Default___.SafeDivisionInt(_166_v6, (_166_v6).Times((_dafny.Zero).Minus((_dafny.SetOf(_166_v6)).Cardinality())))
		} else {
			var _313___mcc_h29 D0 = _source8.Get_().(D0_DC3).Cf7
			_ = _313___mcc_h29
			var _314_cf7 D0 = _313___mcc_h29
			_ = _314_cf7
			var _315_v85 _dafny.Array
			_ = _315_v85
			var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
			_ = _nw50
			_315_v85 = _nw50
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_315_v85), 0))
			_ = _index56
			(_315_v85).ArraySet1(_160_v0, (_index56).Int())
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))
			_ = _index57
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_315_v85), 0))
			_ = _index58
			var _rhs37 _dafny.CodePoint = Companion_Default___.Fm11(_166_v6, _182_globalState)
			_ = _rhs37
			var _rhs38 _dafny.Array = _160_v0
			_ = _rhs38
			var _rhs39 bool = !(_162_v2)
			_ = _rhs39
			var _lhs31 _dafny.Array = _288_v74
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_288_v74), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _315_v85
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_315_v85), 0))
			_ = _lhs34
			(_lhs31).ArraySet1CodePoint(_rhs37, (_lhs32).Int())
			(_lhs33).ArraySet1(_rhs38, (_lhs34).Int())
			_162_v2 = _rhs39
			_166_v6 = _166_v6
			(_182_globalState).F10 = _166_v6
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_173_v11), 0))
			_ = _index59
			(_173_v11).ArraySet1(_169_v9, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_173_v11), 0))
			_ = _index60
			var _rhs40 _dafny.Int = _166_v6
			_ = _rhs40
			var _rhs41 _dafny.Int = (_dafny.Zero).Minus(_166_v6)
			_ = _rhs41
			var _rhs42 _dafny.MultiSet = _dafny.MultiSetOf(_166_v6, _166_v6)
			_ = _rhs42
			var _lhs35 *GlobalState = _182_globalState
			_ = _lhs35
			var _lhs36 _dafny.Array = _173_v11
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_173_v11), 0))
			_ = _lhs37
			_lhs35.F10 = _rhs40
			_166_v6 = _rhs41
			(_lhs36).ArraySet1(_rhs42, (_lhs37).Int())
		}
		var _316_v86 _dafny.Map
		_ = _316_v86
		_316_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, Companion_Default___.SafeDivisionInt((_254_v54).Cardinality(), _dafny.IntOfUint32((_168_v8).Cardinality())))
		var _317_v87 *C1
		_ = _317_v87
		var _nw51 *C1 = New_C1_()
		_ = _nw51
		_nw51.Ctor__()
		_317_v87 = _nw51
		var _318_v88 _dafny.Sequence
		_ = _318_v88
		_318_v88 = _dafny.SeqOf(_317_v87)
		_316_v86 = (_316_v86).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_318_v88, (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_169_v9).Contains(_166_v6) {
				return (_169_v9).Multiplicity(_166_v6)
			}
			return _166_v6
		})(), _dafny.IntOfUint32((_318_v88).Cardinality()))).Uint32(), _317_v87), _162_v2)).Cardinality(), _166_v6)
		var _319_v89 _dafny.Array
		_ = _319_v89
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
		_ = _nw52
		_319_v89 = _nw52
		var _320_v90 T1
		_ = _320_v90
		var _nw53 *C5 = New_C5_()
		_ = _nw53
		_nw53.Ctor__(_319_v89, _dafny.IntOfInt64(-624))
		_320_v90 = _nw53
		_320_v90 = (func() T1 {
			if _162_v2 {
				return _320_v90
			}
			return _320_v90
		})()
	}
	(_182_globalState).F13 = _162_v2
	var _321_i4 _dafny.Int
	_ = _321_i4
	_321_i4 = _dafny.Zero
	{
		for _162_v2 {
			{
				if (_321_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_321_i4 = (_321_i4).Plus(_dafny.One)
				var _322_v91 _dafny.Map
				_ = _322_v91
				_322_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_163_v3, _166_v6)
				var _323_v92 _dafny.Map
				_ = _323_v92
				_323_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, (_dafny.Zero).Minus((_322_v91).Cardinality()))
				var _324_v93 *C0
				_ = _324_v93
				var _nw54 *C0 = New_C0_()
				_ = _nw54
				_nw54.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(Companion_Default___.Fm10(_323_v92, _175_v13, true, _182_globalState), _168_v8))
				_324_v93 = _nw54
				var _325_v94 _dafny.Array
				_ = _325_v94
				var _len0_3 _dafny.Int = _dafny.One
				_ = _len0_3
				var _nw55 _dafny.Array
				_ = _nw55
				if _len0_3.Cmp(_dafny.Zero) == 0 {
					_nw55 = _dafny.NewArray(_len0_3)
				} else {
					var _init3 func(_dafny.Int) _dafny.Int = (func(_326_v6 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_327_i5 _dafny.Int) _dafny.Int {
							return (_327_i5).Minus(_326_v6)
						}
					})(_166_v6)
					_ = _init3
					var _element0_3 = _init3(_dafny.Zero)
					_ = _element0_3
					_nw55 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
					(_nw55).ArraySet1(_element0_3, 0)
					var _nativeLen0_3 = (_len0_3).Int()
					_ = _nativeLen0_3
					for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
						(_nw55).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
					}
				}
				_325_v94 = _nw55
				var _328_v95 _dafny.Map
				_ = _328_v95
				_328_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_325_v94, _166_v6)
				(_182_globalState).F22 = Companion_Default___.Fm12(_162_v2, !(_328_v95).Equals(_328_v95), _182_globalState)
				var _329_v96 *C3
				_ = _329_v96
				var _nw56 *C3 = New_C3_()
				_ = _nw56
				_nw56.Ctor__()
				_329_v96 = _nw56
				_329_v96 = _329_v96
				var _330_v97 _dafny.Array
				_ = _330_v97
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
				_ = _nw57
				_330_v97 = _nw57
				var _331_v98 *C4
				_ = _331_v98
				var _nw58 *C4 = New_C4_()
				_ = _nw58
				_nw58.Ctor__(_330_v97)
				_331_v98 = _nw58
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_325_v94), 0))
				_ = _index61
				(_325_v94).ArraySet1(_166_v6, (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_325_v94), 0))
				_ = _index62
				var _rhs43 bool = !_dafny.Companion_Sequence_.Contains(_168_v8, _175_v13)
				_ = _rhs43
				var _rhs44 *C4 = _331_v98
				_ = _rhs44
				var _rhs45 _dafny.Sequence = _168_v8
				_ = _rhs45
				var _rhs46 _dafny.Int = _166_v6
				_ = _rhs46
				var _lhs38 *GlobalState = _182_globalState
				_ = _lhs38
				var _lhs39 _dafny.Array = _325_v94
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_325_v94), 0))
				_ = _lhs40
				_lhs38.F7 = _rhs43
				_331_v98 = _rhs44
				_168_v8 = _rhs45
				(_lhs39).ArraySet1(_rhs46, (_lhs40).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _332_v100 _dafny.Map
	_ = _332_v100
	_332_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_168_v8).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _333_v99 _dafny.CodePoint
			_333_v99 = interface{}(_compr_16).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_168_v8, _333_v99) {
				_coll16.Add(_333_v99, _166_v6)
			}
		}
		return _coll16.ToMap()
	}())
	(_182_globalState).F10 = (func() _dafny.Int {
		if Companion_Default___.Fm12(!(_162_v2), _162_v2, _182_globalState) {
			return ((_332_v100).Cardinality()).Times(_166_v6)
		}
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_168_v8).Cardinality()), _166_v6)
	})()
	var _334_v101 D12
	_ = _334_v101
	_334_v101 = Companion_D12_.Create_DC28_(_dafny.MultiSetFromSeq(_164_v4))
	(_182_globalState).F13 = ((_169_v9).Intersection((_334_v101).Dtor_cf44())).IsProperSubsetOf(_169_v9)
	var _335_v102 _dafny.Array
	_ = _335_v102
	var _nwElement0_15 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_162_v2), _163_v3)).Cardinality())
	_ = _nwElement0_15
	var _nw59 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(3))
	_ = _nw59
	(_nw59).ArraySet1(_nwElement0_15, 0)
	(_nw59).ArraySet1(_166_v6, 1)
	(_nw59).ArraySet1(_dafny.IntOfUint32((_168_v8).Cardinality()), 2)
	_335_v102 = _nw59
	for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_335_v102), 0))); ; {
		_guard_loop_0, _ok17 := _iter17()
		if !_ok17 {
			break
		}
		var _336_i6 _dafny.Int
		_336_i6 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_336_i6).Sign() != -1) && ((_336_i6).Cmp(_dafny.ArrayLen((_335_v102), 0)) < 0)) {
			(_335_v102).ArraySet1((_336_i6).Minus((_166_v6).Minus(_166_v6)), (_336_i6).Int())
		}
	}
	if !(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_163_v3, (Companion_D8_.Create_DC21_(_163_v3)).Dtor_cf35()), _163_v3)) {
		var _337_v103 _dafny.Map
		_ = _337_v103
		_337_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, (_166_v6).Plus(_166_v6))
		_337_v103 = (_337_v103).Update(_175_v13, _dafny.IntOfInt64(526))
		var _338_v104 D3
		_ = _338_v104
		_338_v104 = Companion_D3_.Create_DC10_(_dafny.MultiSetFromSeq(_163_v3))
		var _pat_let_tv1 = _162_v2
		_ = _pat_let_tv1
		var _pat_let_tv2 = _162_v2
		_ = _pat_let_tv2
		var _pat_let_tv3 = _182_globalState
		_ = _pat_let_tv3
		var _pat_let_tv4 = _163_v3
		_ = _pat_let_tv4
		var _source9 D3 = func(_pat_let2_0 D3) D3 {
			return func(_339_dt__update__tmp_h2 D3) D3 {
				return func(_pat_let3_0 _dafny.MultiSet) D3 {
					return func(_340_dt__update_hcf21_h0 _dafny.MultiSet) D3 {
						return Companion_D3_.Create_DC10_(_340_dt__update_hcf21_h0)
					}(_pat_let3_0)
				}((_dafny.MultiSetOf(Companion_Default___.Fm12(_pat_let_tv1, _pat_let_tv2, _pat_let_tv3))).Difference(_dafny.MultiSetFromSeq(_pat_let_tv4)))
			}(_pat_let2_0)
		}(_338_v104)
		_ = _source9
		if _source9.Is_DC11() {
			var _341___mcc_h30 _dafny.Int = _source9.Get_().(D3_DC11).Cf22
			_ = _341___mcc_h30
			var _342_cf22 _dafny.Int = _341___mcc_h30
			_ = _342_cf22
			var _343_v105 *C3
			_ = _343_v105
			var _nw60 *C3 = New_C3_()
			_ = _nw60
			_nw60.Ctor__()
			_343_v105 = _nw60
			var _344_v106 _dafny.MultiSet
			_ = _344_v106
			_344_v106 = _dafny.MultiSetOf(_343_v105)
			(_182_globalState).F22 = !(((func() _dafny.Int {
				if (_344_v106).Contains(_343_v105) {
					return (_344_v106).Multiplicity(_343_v105)
				}
				return (_343_v105).Fm8(_342_cf22, _166_v6, _342_cf22, _182_globalState)
			})()).Cmp(((_169_v9).Cardinality()).Times(_342_cf22)) > 0)
			var _345_v107 D4
			_ = _345_v107
			_345_v107 = Companion_D4_.Create_DC13_(_175_v13)
			var _346_v108 D4
			_ = _346_v108
			_346_v108 = Companion_D4_.Create_DC15_(_345_v107)
			var _347_v109 D7
			_ = _347_v109
			_347_v109 = Companion_D7_.Create_DC20_(_dafny.SeqOf(_342_cf22), _160_v0, _346_v108, _162_v2, (func() _dafny.Int {
				if _162_v2 {
					return _166_v6
				}
				return _342_cf22
			})())
			_347_v109 = _347_v109
			var _348_v110 _dafny.Array
			_ = _348_v110
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw61
			_348_v110 = _nw61
			var _349_v111 *C4
			_ = _349_v111
			var _nw62 *C4 = New_C4_()
			_ = _nw62
			_nw62.Ctor__(_348_v110)
			_349_v111 = _nw62
			var _350_v112 T1
			_ = _350_v112
			var _nw63 *C4 = New_C4_()
			_ = _nw63
			_nw63.Ctor__((_349_v111).F31())
			_350_v112 = _nw63
			var _351_v113 _dafny.Array
			_ = _351_v113
			_351_v113 = _335_v102
			var _352_v114 _dafny.Array
			_ = _352_v114
			var _nwElement0_16 _dafny.Array = _335_v102
			_ = _nwElement0_16
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(12))
			_ = _nw64
			(_nw64).ArraySet1(_nwElement0_16, 0)
			(_nw64).ArraySet1((_351_v113), 1)
			(_nw64).ArraySet1(_335_v102, 2)
			(_nw64).ArraySet1(_335_v102, 3)
			(_nw64).ArraySet1(_335_v102, 4)
			(_nw64).ArraySet1(_335_v102, 5)
			(_nw64).ArraySet1(_335_v102, 6)
			(_nw64).ArraySet1(_335_v102, 7)
			(_nw64).ArraySet1(_335_v102, 8)
			(_nw64).ArraySet1(_335_v102, 9)
			(_nw64).ArraySet1(_335_v102, 10)
			(_nw64).ArraySet1(_335_v102, 11)
			_352_v114 = _nw64
			var _nw65 *C5 = New_C5_()
			_ = _nw65
			_nw65.Ctor__(_352_v114, ((_dafny.Zero).Minus(_166_v6)).Plus(_166_v6))
			_350_v112 = _nw65
		} else if _source9.Is_DC12() {
			var _353___mcc_h31 bool = _source9.Get_().(D3_DC12).Cf23
			_ = _353___mcc_h31
			var _354_cf23 bool = _353___mcc_h31
			_ = _354_cf23
			(_182_globalState).F10 = _166_v6
			var _355_v115 _dafny.Map
			_ = _355_v115
			_355_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, _254_v54)
			var _356_v116 _dafny.Array
			_ = _356_v116
			var _nwElement0_17 _dafny.Set = _254_v54
			_ = _nwElement0_17
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(13))
			_ = _nw66
			(_nw66).ArraySet1(_nwElement0_17, 0)
			(_nw66).ArraySet1((_dafny.SetOf(_354_cf23)).Union(_254_v54), 1)
			(_nw66).ArraySet1(_254_v54, 2)
			(_nw66).ArraySet1(_254_v54, 3)
			(_nw66).ArraySet1(_254_v54, 4)
			(_nw66).ArraySet1(_dafny.SetOf(_162_v2, _354_cf23), 5)
			(_nw66).ArraySet1(_254_v54, 6)
			(_nw66).ArraySet1((func() _dafny.Set {
				if _354_cf23 {
					return (func() _dafny.Set {
						if (_355_v115).Contains(_175_v13) {
							return (_355_v115).Get(_175_v13).(_dafny.Set)
						}
						return _254_v54
					})()
				}
				return _dafny.SetOf(_162_v2, _354_cf23)
			})(), 7)
			(_nw66).ArraySet1(_dafny.SetOf(_354_cf23, (_163_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-580), _dafny.IntOfUint32((_163_v3).Cardinality()))).Uint32()).(bool), _354_cf23, Companion_Default___.Fm12(!(_354_cf23), false, _182_globalState)), 8)
			(_nw66).ArraySet1(_254_v54, 9)
			(_nw66).ArraySet1(_254_v54, 10)
			(_nw66).ArraySet1(_254_v54, 11)
			(_nw66).ArraySet1(_254_v54, 12)
			_356_v116 = _nw66
			var _357_v117 _dafny.Map
			_ = _357_v117
			_357_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, !(true))
			var _358_v118 _dafny.Map
			_ = _358_v118
			_358_v118 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _354_cf23)).Update(false, _354_cf23)
			var _359_v119 D12
			_ = _359_v119
			_359_v119 = Companion_D12_.Create_DC29_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(993))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_360_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(351)
			})), _166_v6, _162_v2)
			var _361_v120 _dafny.Array
			_ = _361_v120
			var _nwElement0_18 _dafny.Map = (func() _dafny.Map {
				if _162_v2 {
					return _357_v117
				}
				return _358_v118
			})()
			_ = _nwElement0_18
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(26))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_18, 0)
			(_nw67).ArraySet1(_358_v118, 1)
			(_nw67).ArraySet1(_358_v118, 2)
			(_nw67).ArraySet1(_357_v117, 3)
			(_nw67).ArraySet1(_357_v117, 4)
			(_nw67).ArraySet1(_358_v118, 5)
			(_nw67).ArraySet1(_358_v118, 6)
			(_nw67).ArraySet1(_358_v118, 7)
			(_nw67).ArraySet1(_357_v117, 8)
			(_nw67).ArraySet1(_358_v118, 9)
			(_nw67).ArraySet1(_358_v118, 10)
			(_nw67).ArraySet1(_358_v118, 11)
			(_nw67).ArraySet1(_358_v118, 12)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _354_cf23), 13)
			(_nw67).ArraySet1(_358_v118, 14)
			(_nw67).ArraySet1(_358_v118, 15)
			(_nw67).ArraySet1((_357_v117).Update(_162_v2, !(_354_cf23)), 16)
			(_nw67).ArraySet1(_357_v117, 17)
			(_nw67).ArraySet1(_358_v118, 18)
			(_nw67).ArraySet1(_358_v118, 19)
			(_nw67).ArraySet1(_358_v118, 20)
			(_nw67).ArraySet1(_358_v118, 21)
			(_nw67).ArraySet1(_358_v118, 22)
			(_nw67).ArraySet1(_358_v118, 23)
			(_nw67).ArraySet1(_358_v118, 24)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, (_359_v119).Dtor_cf47()), 25)
			_361_v120 = _nw67
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_361_v120), 0))
			_ = _index63
			(_361_v120).ArraySet1((_357_v117).Update(_354_cf23, _162_v2), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))
			_ = _index64
			(_335_v102).ArraySet1(Companion_Default___.Fm0(_dafny.IntOfInt64(504), _182_globalState), (_index64).Int())
			var _362_v121 _dafny.Set
			_ = _362_v121
			_362_v121 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_168_v8, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_166_v6), _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32(), (_168_v8).Select((Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_168_v8).Cardinality()))).Uint32()).(_dafny.CodePoint)))
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_361_v120), 0))
			_ = _index65
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))
			_ = _index66
			var _rhs47 _dafny.Array = _356_v116
			_ = _rhs47
			var _rhs48 _dafny.Int = (_166_v6).Plus(_dafny.IntOfUint32((_168_v8).Cardinality()))
			_ = _rhs48
			var _rhs49 _dafny.Map = _358_v118
			_ = _rhs49
			var _rhs50 _dafny.Int = Companion_Default___.SafeDivisionInt(_166_v6, (_362_v121).Cardinality())
			_ = _rhs50
			var _rhs51 _dafny.Int = _dafny.IntOfInt64(951)
			_ = _rhs51
			var _lhs41 _dafny.Array = _361_v120
			_ = _lhs41
			var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(180), _dafny.ArrayLen((_361_v120), 0))
			_ = _lhs42
			var _lhs43 _dafny.Array = _335_v102
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))
			_ = _lhs44
			_356_v116 = _rhs47
			_166_v6 = _rhs48
			(_lhs41).ArraySet1(_rhs49, (_lhs42).Int())
			(_lhs43).ArraySet1(_rhs50, (_lhs44).Int())
			_166_v6 = _rhs51
			(_182_globalState).F10 = (_335_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))).Int()).(_dafny.Int)
			var _363_v122 _dafny.Map
			_ = _363_v122
			_363_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v8, (_183_v20).Equals(_183_v20))
			var _364_v123 _dafny.Array
			_ = _364_v123
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw68
			_364_v123 = _nw68
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_364_v123), 0))
			_ = _index67
			(_364_v123).ArraySet1(_180_v18, (_index67).Int())
			var _365_v124 _dafny.Array
			_ = _365_v124
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw69
			_365_v124 = _nw69
			var _366_v125 *C4
			_ = _366_v125
			var _nw70 *C4 = New_C4_()
			_ = _nw70
			_nw70.Ctor__(_365_v124)
			_366_v125 = _nw70
			var _367_v126 _dafny.MultiSet
			_ = _367_v126
			_367_v126 = _dafny.MultiSetOf(_366_v125)
			var _368_v127 _dafny.Sequence
			_ = _368_v127
			_368_v127 = _dafny.SeqOf(_367_v126)
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_160_v0), 0))
			_ = _index68
			(_160_v0).ArraySet1(_dafny.Companion_Sequence_.Contains(_368_v127, _dafny.MultiSetOf(_366_v125)), (_index68).Int())
			var _369_v128 _dafny.Map
			_ = _369_v128
			_369_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(387), _168_v8)
			var _370_v129 _dafny.Map
			_ = _370_v129
			_370_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_369_v128).Contains(_dafny.IntOfUint32((_168_v8).Cardinality())) {
					return (_369_v128).Get(_dafny.IntOfUint32((_168_v8).Cardinality())).(_dafny.Sequence)
				}
				return _168_v8
			})()).Cardinality()))
			var _371_v130 _dafny.Array
			_ = _371_v130
			var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw71
			_371_v130 = _nw71
			var _372_v131 *C5
			_ = _372_v131
			var _nw72 *C5 = New_C5_()
			_ = _nw72
			_nw72.Ctor__(_371_v130, _dafny.IntOfInt64(-280))
			_372_v131 = _nw72
			var _373_v132 _dafny.Map
			_ = _373_v132
			_373_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_169_v9, _372_v131)
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))
			_ = _index69
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_364_v123), 0))
			_ = _index70
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_160_v0), 0))
			_ = _index71
			var _rhs52 _dafny.Int = (_166_v6).Times((func() _dafny.Int {
				if (_370_v129).Contains(_162_v2) {
					return (_370_v129).Get(_162_v2).(_dafny.Int)
				}
				return _166_v6
			})())
			_ = _rhs52
			var _rhs53 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v8, _354_cf23)).Merge(_363_v122)
			_ = _rhs53
			var _rhs54 _dafny.Array = _180_v18
			_ = _rhs54
			var _rhs55 bool = (Companion_Default___.Fm0((_373_v132).Cardinality(), _182_globalState)).Cmp(_dafny.IntOfInt64(-786)) >= 0
			_ = _rhs55
			var _rhs56 bool = _162_v2
			_ = _rhs56
			var _lhs45 _dafny.Array = _335_v102
			_ = _lhs45
			var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_335_v102), 0))
			_ = _lhs46
			var _lhs47 _dafny.Array = _364_v123
			_ = _lhs47
			var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_364_v123), 0))
			_ = _lhs48
			var _lhs49 *GlobalState = _182_globalState
			_ = _lhs49
			var _lhs50 _dafny.Array = _160_v0
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_160_v0), 0))
			_ = _lhs51
			(_lhs45).ArraySet1(_rhs52, (_lhs46).Int())
			_363_v122 = _rhs53
			(_lhs47).ArraySet1(_rhs54, (_lhs48).Int())
			_lhs49.F7 = _rhs55
			(_lhs50).ArraySet1(_rhs56, (_lhs51).Int())
		} else {
			var _374___mcc_h32 _dafny.MultiSet = _source9.Get_().(D3_DC10).Cf21
			_ = _374___mcc_h32
			var _375_cf21 _dafny.MultiSet = _374___mcc_h32
			_ = _375_cf21
			var _rhs57 bool = Companion_Default___.Fm12(_162_v2, _162_v2, _182_globalState)
			_ = _rhs57
			var _rhs58 bool = true
			_ = _rhs58
			var _rhs59 _dafny.Sequence = _dafny.SeqOf(_166_v6, _166_v6, _166_v6, _166_v6, _166_v6)
			_ = _rhs59
			var _lhs52 *GlobalState = _182_globalState
			_ = _lhs52
			var _lhs53 *GlobalState = _182_globalState
			_ = _lhs53
			_lhs52.F7 = _rhs57
			_lhs53.F7 = _rhs58
			_164_v4 = _rhs59
			_168_v8 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qmpwft"), _168_v8), _168_v8), (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qmpwft"), _168_v8), _168_v8)).Cardinality()))).Uint32(), _175_v13)
			var _376_v133 _dafny.Array
			_ = _376_v133
			var _nw73 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
			_ = _nw73
			_376_v133 = _nw73
			var _377_v134 *C3
			_ = _377_v134
			var _nw74 *C3 = New_C3_()
			_ = _nw74
			_nw74.Ctor__()
			_377_v134 = _nw74
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_376_v133), 0))
			_ = _index72
			(_376_v133).ArraySet1(_377_v134, (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(127), _dafny.ArrayLen((_376_v133), 0))
			_ = _index73
			(_376_v133).ArraySet1(_377_v134, (_index73).Int())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_160_v0), 0))
			_ = _index74
			(_160_v0).ArraySet1(Companion_Default___.Fm12(!(_162_v2), _162_v2, _182_globalState), (_index74).Int())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_160_v0), 0))
			_ = _index75
			var _rhs60 _dafny.Int = _166_v6
			_ = _rhs60
			var _rhs61 bool = Companion_Default___.Fm12(_162_v2, _162_v2, _182_globalState)
			_ = _rhs61
			var _rhs62 _dafny.Int = (func() _dafny.Int {
				if true {
					return (_377_v134).Fm8(_166_v6, _dafny.IntOfInt64(184), _166_v6, _182_globalState)
				}
				return (_377_v134).Fm8((_181_v19).Cardinality(), _166_v6, _dafny.IntOfUint32((_163_v3).Cardinality()), _182_globalState)
			})()
			_ = _rhs62
			var _lhs54 _dafny.Array = _160_v0
			_ = _lhs54
			var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_160_v0), 0))
			_ = _lhs55
			_166_v6 = _rhs60
			(_lhs54).ArraySet1(_rhs61, (_lhs55).Int())
			_166_v6 = _rhs62
		}
		var _378_v135 _dafny.Array
		_ = _378_v135
		var _nwElement0_19 _dafny.Array = _335_v102
		_ = _nwElement0_19
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(5))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_19, 0)
		(_nw75).ArraySet1(_335_v102, 1)
		(_nw75).ArraySet1(_335_v102, 2)
		(_nw75).ArraySet1(_335_v102, 3)
		(_nw75).ArraySet1(_335_v102, 4)
		_378_v135 = _nw75
		var _379_v136 _dafny.Map
		_ = _379_v136
		_379_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_164_v4), _166_v6)
		var _380_v137 _dafny.Map
		_ = _380_v137
		_380_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_379_v136, _dafny.SetOf(_162_v2))
		var _381_v139 _dafny.Map
		_ = _381_v139
		_381_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
			if (_380_v137).Contains(_379_v136) {
				return (_380_v137).Get(_379_v136).(_dafny.Set)
			}
			return _254_v54
		})(), (_dafny.Zero).Minus((_dafny.MultiSetOf(func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter18 := _dafny.Iterate((_181_v19).Elements()); ; {
				_compr_17, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _382_v138 _dafny.Int
				_382_v138 = interface{}(_compr_17).(_dafny.Int)
				if (_181_v19).Contains(_382_v138) {
					_coll17.Add((_382_v138).Plus(_dafny.IntOfInt64(-583)), _166_v6)
				}
			}
			return _coll17.ToMap()
		}())).Cardinality()))
		var _383_v140 T1
		_ = _383_v140
		var _nw76 *C5 = New_C5_()
		_ = _nw76
		_nw76.Ctor__(_378_v135, (_dafny.Zero).Minus((func() _dafny.Int {
			if (_381_v139).Contains(_254_v54) {
				return (_381_v139).Get(_254_v54).(_dafny.Int)
			}
			return _166_v6
		})()))
		_383_v140 = _nw76
		var _384_v141 _dafny.Array
		_ = _384_v141
		var _nwElement0_20 T1 = _383_v140
		_ = _nwElement0_20
		var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(2))
		_ = _nw77
		(_nw77).ArraySet1(_nwElement0_20, 0)
		(_nw77).ArraySet1(_383_v140, 1)
		_384_v141 = _nw77
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_384_v141), 0))
		_ = _index76
		(_384_v141).ArraySet1(_383_v140, (_index76).Int())
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_384_v141), 0))
		_ = _index77
		var _rhs63 _dafny.Int = Companion_Default___.SafeModuloInt(_166_v6, (_dafny.Zero).Minus((_dafny.Zero).Minus(_166_v6)))
		_ = _rhs63
		var _rhs64 T1 = _383_v140
		_ = _rhs64
		var _rhs65 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((func() bool {
			if (_183_v20).Contains((_181_v19).Cardinality()) {
				return (_183_v20).Get((_181_v19).Cardinality()).(bool)
			}
			return _162_v2
		})())).Cardinality()))
		_ = _rhs65
		var _lhs56 *GlobalState = _182_globalState
		_ = _lhs56
		var _lhs57 _dafny.Array = _384_v141
		_ = _lhs57
		var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_384_v141), 0))
		_ = _lhs58
		var _lhs59 *GlobalState = _182_globalState
		_ = _lhs59
		_lhs56.F10 = _rhs63
		(_lhs57).ArraySet1(_rhs64, (_lhs58).Int())
		_lhs59.F10 = _rhs65
		var _385_v142 D12
		_ = _385_v142
		_385_v142 = Companion_D12_.Create_DC31_(_166_v6, !(false), _162_v2, (_169_v9).Cardinality())
		var _386_v143 D8
		_ = _386_v143
		_386_v143 = Companion_D8_.Create_DC23_((_385_v142).Dtor_cf52())
		_386_v143 = _386_v143
		var _387_v144 *C0
		_ = _387_v144
		var _nw78 *C0 = New_C0_()
		_ = _nw78
		_nw78.Ctor__(_162_v2)
		_387_v144 = _nw78
	} else {
		Companion_Default___.M8(_182_globalState)
		(_182_globalState).F10 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_166_v6, _182_globalState), _166_v6))
		(_182_globalState).F7 = !(!_dafny.Companion_Sequence_.Equal(_168_v8, _dafny.UnicodeSeqOfUtf8Bytes("glmboqi")))
		var _388_v145 *C2
		_ = _388_v145
		var _nw79 *C2 = New_C2_()
		_ = _nw79
		_nw79.Ctor__((Companion_D0_.Create_DC0_(_168_v8)).Dtor_cf0())
		_388_v145 = _nw79
		(_388_v145).F33 = _388_v145.F33
	}
	_335_v102 = _335_v102
	if _162_v2 {
		_166_v6 = _166_v6
		var _389_v146 _dafny.Map
		_ = _389_v146
		_389_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _166_v6)
		_389_v146 = (_389_v146).Update((func() bool {
			if _162_v2 {
				return _162_v2
			}
			return _162_v2
		})(), _166_v6)
		var _nwElement0_21 bool = !(!(_162_v2) || (_162_v2))
		_ = _nwElement0_21
		var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(2))
		_ = _nw80
		(_nw80).ArraySet1(_nwElement0_21, 0)
		(_nw80).ArraySet1(_162_v2, 1)
		_160_v0 = _nw80
		var _390_v147 _dafny.Array
		_ = _390_v147
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(18))
		_ = _nw81
		_390_v147 = _nw81
		var _391_v148 D8
		_ = _391_v148
		_391_v148 = Companion_D8_.Create_DC22_(_166_v6, _162_v2, _dafny.CodePoint('v'))
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_390_v147), 0))
		_ = _index78
		(_390_v147).ArraySet1(_391_v148, (_index78).Int())
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_390_v147), 0))
		_ = _index79
		(_390_v147).ArraySet1(_391_v148, (_index79).Int())
		var _392_v149 _dafny.Array
		_ = _392_v149
		var _nw82 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw82
		_392_v149 = _nw82
		var _393_v150 _dafny.Array
		_ = _393_v150
		_393_v150 = _392_v149
		var _394_v151 *C4
		_ = _394_v151
		var _nw83 *C4 = New_C4_()
		_ = _nw83
		_nw83.Ctor__((_393_v150))
		_394_v151 = _nw83
	} else {
		_168_v8 = _dafny.UnicodeSeqOfUtf8Bytes("faliulvks")
		var _395_v152 _dafny.Array
		_ = _395_v152
		var _nwElement0_22 _dafny.Array = _335_v102
		_ = _nwElement0_22
		var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(22))
		_ = _nw84
		(_nw84).ArraySet1(_nwElement0_22, 0)
		(_nw84).ArraySet1(_335_v102, 1)
		(_nw84).ArraySet1(_335_v102, 2)
		(_nw84).ArraySet1(_335_v102, 3)
		(_nw84).ArraySet1(_335_v102, 4)
		(_nw84).ArraySet1(_335_v102, 5)
		(_nw84).ArraySet1(_335_v102, 6)
		(_nw84).ArraySet1(_335_v102, 7)
		(_nw84).ArraySet1(_335_v102, 8)
		(_nw84).ArraySet1(_335_v102, 9)
		(_nw84).ArraySet1(_335_v102, 10)
		(_nw84).ArraySet1(_335_v102, 11)
		(_nw84).ArraySet1(_335_v102, 12)
		(_nw84).ArraySet1(_335_v102, 13)
		(_nw84).ArraySet1(_335_v102, 14)
		(_nw84).ArraySet1(_335_v102, 15)
		(_nw84).ArraySet1(_335_v102, 16)
		(_nw84).ArraySet1(_335_v102, 17)
		(_nw84).ArraySet1(_335_v102, 18)
		(_nw84).ArraySet1(_335_v102, 19)
		(_nw84).ArraySet1(_335_v102, 20)
		(_nw84).ArraySet1(_335_v102, 21)
		_395_v152 = _nw84
		var _396_v153 *C5
		_ = _396_v153
		var _nw85 *C5 = New_C5_()
		_ = _nw85
		_nw85.Ctor__(_395_v152, _166_v6)
		_396_v153 = _nw85
		var _397_v154 _dafny.MultiSet
		_ = _397_v154
		_397_v154 = _dafny.MultiSetOf(_335_v102, _335_v102, _335_v102)
		var _rhs66 *C5 = _396_v153
		_ = _rhs66
		var _rhs67 bool = (_254_v54).Contains(_162_v2)
		_ = _rhs67
		var _rhs68 _dafny.Array = _335_v102
		_ = _rhs68
		var _rhs69 _dafny.MultiSet = _dafny.MultiSetOf(_335_v102, _335_v102, _335_v102)
		_ = _rhs69
		var _rhs70 _dafny.Array = _335_v102
		_ = _rhs70
		var _lhs60 *GlobalState = _182_globalState
		_ = _lhs60
		_396_v153 = _rhs66
		_lhs60.F13 = _rhs67
		_335_v102 = _rhs68
		_397_v154 = _rhs69
		_335_v102 = _rhs70
		var _398_v155 _dafny.Array
		_ = _398_v155
		var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
		_ = _nw86
		_398_v155 = _nw86
		var _399_v156 T1
		_ = _399_v156
		var _nw87 *C4 = New_C4_()
		_ = _nw87
		_nw87.Ctor__(_398_v155)
		_399_v156 = _nw87
		var _400_v157 _dafny.Set
		_ = _400_v157
		_400_v157 = _dafny.SetOf((Companion_D14_.Create_DC33_(_399_v156)).Dtor_cf56(), _399_v156, _399_v156)
		var _401_v158 _dafny.Map
		_ = _401_v158
		_401_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_400_v157, (_179_v17).Select((Companion_Default___.SafeIndex(_396_v153.F30, _dafny.IntOfUint32((_179_v17).Cardinality()))).Uint32()).(_dafny.Sequence))
		_401_v158 = (_401_v158).Update(_400_v157, _dafny.Companion_Sequence_.Concatenate(_168_v8, _168_v8))
		var _402_v159 _dafny.Array
		_ = _402_v159
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_4
		var _nw88 _dafny.Array
		_ = _nw88
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw88 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Map = (func(_403_v2 bool, _404_v3 _dafny.Sequence, _405_v6 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_406_i8 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v2, (_404_v3).Select((Companion_Default___.SafeIndex(_405_v6, _dafny.IntOfUint32((_404_v3).Cardinality()))).Uint32()).(bool))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v2, _403_v2))
				}
			})(_162_v2, _163_v3, _166_v6)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw88 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw88).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw88).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_402_v159 = _nw88
		var _407_v160 _dafny.Map
		_ = _407_v160
		_407_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _162_v2)
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_402_v159), 0))
		_ = _index80
		(_402_v159).ArraySet1((_407_v160).Update(_162_v2, (_396_v153).Fm4(_168_v8, _166_v6, _182_globalState)), (_index80).Int())
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_402_v159), 0))
		_ = _index81
		(_402_v159).ArraySet1(((_407_v160).Merge(_407_v160)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _162_v2)), (_index81).Int())
		(_396_v153).F30 = _396_v153.F30
	}
	var _408_v161 D4
	_ = _408_v161
	_408_v161 = Companion_D4_.Create_DC15_(Companion_D4_.Create_DC14_(_162_v2))
	(_182_globalState).F10 = (Companion_D7_.Create_DC20_(_dafny.Companion_Sequence_.Update(_164_v4, (Companion_Default___.SafeIndex(_166_v6, _dafny.IntOfUint32((_164_v4).Cardinality()))).Uint32(), _166_v6), _160_v0, _408_v161, _162_v2, _166_v6)).Dtor_cf34()
	if _162_v2 {
		var _409_v162 D7
		_ = _409_v162
		_409_v162 = Companion_D7_.Create_DC19_()
		var _rhs71 bool = _162_v2
		_ = _rhs71
		var _rhs72 bool = Companion_Default___.Fm12(_162_v2, _162_v2, _182_globalState)
		_ = _rhs72
		var _rhs73 D7 = Companion_D7_.Create_DC19_()
		_ = _rhs73
		var _lhs61 *GlobalState = _182_globalState
		_ = _lhs61
		var _lhs62 *GlobalState = _182_globalState
		_ = _lhs62
		_lhs61.F13 = _rhs71
		_lhs62.F13 = _rhs72
		_409_v162 = _rhs73
		(_182_globalState).F13 = _162_v2
		if (_163_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.IntOfUint32((_163_v3).Cardinality()))).Uint32()).(bool) {
			(_182_globalState).F10 = ((_dafny.Zero).Minus(_166_v6)).Times(_166_v6)
			var _410_v163 _dafny.Map
			_ = _410_v163
			_410_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_164_v4, _162_v2)
			_410_v163 = _410_v163
			_168_v8 = _168_v8
			_164_v4 = _dafny.Companion_Sequence_.Concatenate(_164_v4, _dafny.SeqOf(_166_v6))
			var _411_v164 D4
			_ = _411_v164
			_411_v164 = Companion_D4_.Create_DC13_(_175_v13)
			var _pat_let_tv5 = _175_v13
			_ = _pat_let_tv5
			var _412_v165 _dafny.Map
			_ = _412_v165
			_412_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let4_0 D4) D4 {
				return func(_413_dt__update__tmp_h5 D4) D4 {
					return func(_pat_let5_0 _dafny.CodePoint) D4 {
						return func(_414_dt__update_hcf24_h0 _dafny.CodePoint) D4 {
							return Companion_D4_.Create_DC13_(_414_dt__update_hcf24_h0)
						}(_pat_let5_0)
					}(_pat_let_tv5)
				}(_pat_let4_0)
			}(_411_v164), (func() _dafny.Int {
				if _162_v2 {
					return _166_v6
				}
				return _166_v6
			})())
			(_182_globalState).F10 = (func() _dafny.Int {
				if (_412_v165).Contains(_411_v164) {
					return (_412_v165).Get(_411_v164).(_dafny.Int)
				}
				return _166_v6
			})()
		} else {
			var _415_v166 _dafny.Map
			_ = _415_v166
			_415_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_162_v2), _162_v2)
			var _416_v167 _dafny.Map
			_ = _416_v167
			_416_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_415_v166).Update(false, _162_v2), _166_v6)
			var _rhs74 _dafny.Int = (_dafny.IntOfInt64(802)).Times(_166_v6)
			_ = _rhs74
			var _rhs75 _dafny.Array = (func() _dafny.Array {
				if _162_v2 {
					return _160_v0
				}
				return _160_v0
			})()
			_ = _rhs75
			var _rhs76 _dafny.Map = _416_v167
			_ = _rhs76
			var _rhs77 bool = !(!(_183_v20).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_166_v6, _162_v2)))
			_ = _rhs77
			var _rhs78 bool = _162_v2
			_ = _rhs78
			var _lhs63 *GlobalState = _182_globalState
			_ = _lhs63
			var _lhs64 *GlobalState = _182_globalState
			_ = _lhs64
			var _lhs65 *GlobalState = _182_globalState
			_ = _lhs65
			_lhs63.F10 = _rhs74
			_160_v0 = _rhs75
			_416_v167 = _rhs76
			_lhs64.F22 = _rhs77
			_lhs65.F7 = _rhs78
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_160_v0), 0))
			_ = _index82
			(_160_v0).ArraySet1((func() bool {
				if (_415_v166).Contains(!(_162_v2)) {
					return (_415_v166).Get(!(_162_v2)).(bool)
				}
				return _162_v2
			})(), (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_160_v0), 0))
			_ = _index83
			(_160_v0).ArraySet1(_162_v2, (_index83).Int())
			var _417_v168 _dafny.Array
			_ = _417_v168
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_5
			var _nw89 _dafny.Array
			_ = _nw89
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw89 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) bool = (func(_418_v0 _dafny.Array) func(_dafny.Int) bool {
					return func(_419_i9 _dafny.Int) bool {
						return (_418_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_418_v0), 0))).Int()).(bool)
					}
				})(_160_v0)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw89 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw89).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw89).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_417_v168 = _nw89
			var _420_v169 _dafny.Map
			_ = _420_v169
			_420_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_417_v168, _335_v102)
			_420_v169 = (_420_v169).Update(_417_v168, _335_v102)
			_183_v20 = (_183_v20).Update(_dafny.IntOfUint32((_168_v8).Cardinality()), (_160_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_160_v0), 0))).Int()).(bool))
			(_182_globalState).F10 = _166_v6
		}
		var _421_v170 *C2
		_ = _421_v170
		var _nw90 *C2 = New_C2_()
		_ = _nw90
		_nw90.Ctor__(_168_v8)
		_421_v170 = _nw90
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_335_v102), 0))
		_ = _index84
		(_335_v102).ArraySet1(_dafny.IntOfUint32((_168_v8).Cardinality()), (_index84).Int())
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_335_v102), 0))
		_ = _index85
		(_335_v102).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_163_v3, (Companion_Default___.SafeIndex(((_181_v19).Intersection(_181_v19)).Cardinality(), _dafny.IntOfUint32((_163_v3).Cardinality()))).Uint32(), !(_162_v2))).Cardinality()), (_index85).Int())
	} else {
		var _422_v171 _dafny.Map
		_ = _422_v171
		_422_v171 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-338), _335_v102)
		var _423_v172 _dafny.Array
		_ = _423_v172
		_423_v172 = _335_v102
		var _424_v173 _dafny.Array
		_ = _424_v173
		var _nwElement0_23 _dafny.Array = _335_v102
		_ = _nwElement0_23
		var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(23))
		_ = _nw91
		(_nw91).ArraySet1(_nwElement0_23, 0)
		(_nw91).ArraySet1(_335_v102, 1)
		(_nw91).ArraySet1(_335_v102, 2)
		(_nw91).ArraySet1(_335_v102, 3)
		(_nw91).ArraySet1((func() _dafny.Array {
			if (_422_v171).Contains(_166_v6) {
				return (_422_v171).Get(_166_v6).(_dafny.Array)
			}
			return _335_v102
		})(), 4)
		(_nw91).ArraySet1(_335_v102, 5)
		(_nw91).ArraySet1(_335_v102, 6)
		(_nw91).ArraySet1(_335_v102, 7)
		(_nw91).ArraySet1(_335_v102, 8)
		(_nw91).ArraySet1(_335_v102, 9)
		(_nw91).ArraySet1(_335_v102, 10)
		(_nw91).ArraySet1(_335_v102, 11)
		(_nw91).ArraySet1((_423_v172), 12)
		(_nw91).ArraySet1(_335_v102, 13)
		(_nw91).ArraySet1(_335_v102, 14)
		(_nw91).ArraySet1(_335_v102, 15)
		(_nw91).ArraySet1(_335_v102, 16)
		(_nw91).ArraySet1(_335_v102, 17)
		(_nw91).ArraySet1(_335_v102, 18)
		(_nw91).ArraySet1(_335_v102, 19)
		(_nw91).ArraySet1(_335_v102, 20)
		(_nw91).ArraySet1(_335_v102, 21)
		(_nw91).ArraySet1(_335_v102, 22)
		_424_v173 = _nw91
		var _425_v174 *C5
		_ = _425_v174
		var _nw92 *C5 = New_C5_()
		_ = _nw92
		_nw92.Ctor__(_424_v173, _166_v6)
		_425_v174 = _nw92
		(_182_globalState).F13 = _162_v2
		var _426_v175 _dafny.Set
		_ = _426_v175
		_426_v175 = _dafny.SetOf((_169_v9).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_164_v4).Cardinality()))), _169_v9, (_169_v9).Union(_169_v9))
		var _rhs79 _dafny.Set = (_426_v175).Intersection(_426_v175)
		_ = _rhs79
		var _rhs80 bool = _162_v2
		_ = _rhs80
		var _rhs81 _dafny.Int = (Companion_Default___.Fm43((func() _dafny.CodePoint {
			if _162_v2 {
				return _175_v13
			}
			return _175_v13
		})(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v2, _175_v13), _182_globalState)).Cardinality()
		_ = _rhs81
		var _lhs66 *GlobalState = _182_globalState
		_ = _lhs66
		var _lhs67 *GlobalState = _182_globalState
		_ = _lhs67
		_426_v175 = _rhs79
		_lhs66.F7 = _rhs80
		_lhs67.F10 = _rhs81
		(_182_globalState).F10 = _dafny.IntOfUint32((_168_v8).Cardinality())
		var _427_v176 _dafny.Map
		_ = _427_v176
		_427_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, _160_v0)
		var _428_v177 _dafny.Map
		_ = _428_v177
		_428_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_425_v174, _162_v2)
		var _429_v178 _dafny.Sequence
		_ = _429_v178
		_429_v178 = _dafny.SeqOf(_160_v0, _160_v0)
		var _430_v179 _dafny.Map
		_ = _430_v179
		_430_v179 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_428_v177).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), (_429_v178).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.IntOfUint32((_429_v178).Cardinality()))).Uint32()).(_dafny.Array)))
		var _431_v180 _dafny.Array
		_ = _431_v180
		var _nwElement0_24 _dafny.Map = _427_v176
		_ = _nwElement0_24
		var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(16))
		_ = _nw93
		(_nw93).ArraySet1(_nwElement0_24, 0)
		(_nw93).ArraySet1((_427_v176).Merge(_427_v176), 1)
		(_nw93).ArraySet1(_427_v176, 2)
		(_nw93).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_166_v6, _182_globalState), _160_v0)).Merge(_427_v176), 3)
		(_nw93).ArraySet1(_427_v176, 4)
		(_nw93).ArraySet1((func() _dafny.Map {
			if (_430_v179).Contains(_425_v174.F30) {
				return (_430_v179).Get(_425_v174.F30).(_dafny.Map)
			}
			return _427_v176
		})(), 5)
		(_nw93).ArraySet1((_427_v176).Merge(_427_v176), 6)
		(_nw93).ArraySet1((_427_v176).Merge(_427_v176), 7)
		(_nw93).ArraySet1(_427_v176, 8)
		(_nw93).ArraySet1(_427_v176, 9)
		(_nw93).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, _160_v0), 10)
		(_nw93).ArraySet1((_427_v176).Update(_dafny.CodePoint('d'), _160_v0), 11)
		(_nw93).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v13, _160_v0), 12)
		(_nw93).ArraySet1(_427_v176, 13)
		(_nw93).ArraySet1((_427_v176).Update(_175_v13, _160_v0), 14)
		(_nw93).ArraySet1((_427_v176).Merge(_427_v176), 15)
		_431_v180 = _nw93
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_431_v180), 0))
		_ = _index86
		(_431_v180).ArraySet1(_427_v176, (_index86).Int())
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(522), _dafny.ArrayLen((_431_v180), 0))
		_ = _index87
		(_431_v180).ArraySet1(((_427_v176).Merge((_427_v176).Update(_175_v13, _160_v0))).Merge(_427_v176), (_index87).Int())
	}
	_168_v8 = _dafny.Companion_Sequence_.Concatenate(_168_v8, _168_v8)
	Companion_Default___.M8(_182_globalState)
	Companion_Default___.M8(_182_globalState)
	_dafny.Print((_160_v0).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_160_v0).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_161_v1).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_162_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_163_v3, _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_164_v4, _dafny.SeqOf(_dafny.IntOfInt64(-551), _dafny.IntOfInt64(-551), _dafny.IntOfInt64(-551), _dafny.IntOfInt64(-551), _dafny.IntOfInt64(-551))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_165_v5).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_166_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_167_v7).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_167_v7).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_168_v8.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_169_v9).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_172_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.MultiSetOf(_dafny.IntOfInt64(177)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_173_v11).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_173_v11).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_173_v11).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_173_v11).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-551), _dafny.IntOfInt64(985))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_175_v13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_176_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_177_v15).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_178_v16).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_179_v17, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("wrohtasav"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_180_v18).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_181_v19).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F0()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_182_globalState).F1()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_182_globalState).F1()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_182_globalState).F5()).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_182_globalState).F5()).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_182_globalState).F5()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_182_globalState).F5()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-551), _dafny.IntOfInt64(985))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F8()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(177))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState.F11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F12()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F13)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F17())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F18())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F19)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F20())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_182_globalState).F21()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_182_globalState.F22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F23())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState.F24).Equals(_dafny.SetOf(_dafny.IntOfInt64(667))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_182_globalState).F25(), _dafny.SeqOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F26())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F27())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_182_globalState).F28())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_183_v20).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v21).Dtor_cf1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v21).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_184_v21).Dtor_cf3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v21).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_184_v21).Dtor_cf4()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_254_v54).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_321_i4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_332_v100).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-551), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('r'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('o'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('h'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('a'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('s'), _dafny.IntOfInt64(-551)).UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(-551)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_334_v101).Dtor_cf44()).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v102).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v102).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_335_v102).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_408_v161).Dtor_cf26()).Dtor_cf25())
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
	Cf4 _dafny.Array
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 _dafny.Array) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.Int
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.Int, Cf6 bool) D0 {
	return D0{D0_DC2{Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Sequence
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Sequence) D0 {
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
	return Companion_D0_.Create_DC1_(_dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
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

func (_this D0) Dtor_cf4() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf0() _dafny.Sequence {
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + data.Cf0.VerbatimString(true) + ")"
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
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
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
	Cf9  _dafny.Int
	Cf10 bool
	Cf11 bool
	Cf12 _dafny.Map
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 _dafny.Int, Cf10 bool, Cf11 bool, Cf12 _dafny.Map) D1 {
	return D1{D1_DC5{Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC6 struct {
	Cf13 _dafny.Int
	Cf14 _dafny.Sequence
	Cf15 _dafny.Int
	Cf16 bool
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf13 _dafny.Int, Cf14 _dafny.Sequence, Cf15 _dafny.Int, Cf16 bool) D1 {
	return D1{D1_DC6{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

type D1_DC4 struct {
	Cf8 _dafny.Map
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Map) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(_dafny.Zero, false, false, _dafny.EmptyMap)
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf10() bool {
	return _this.Get_().(D1_DC5).Cf10
}

func (_this D1) Dtor_cf11() bool {
	return _this.Get_().(D1_DC5).Cf11
}

func (_this D1) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D1_DC5).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf13
}

func (_this D1) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D1_DC6).Cf14
}

func (_this D1) Dtor_cf15() _dafny.Int {
	return _this.Get_().(D1_DC6).Cf15
}

func (_this D1) Dtor_cf16() bool {
	return _this.Get_().(D1_DC6).Cf16
}

func (_this D1) Dtor_cf8() _dafny.Map {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11 == data2.Cf11 && data1.Cf12.Equals(data2.Cf12)
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0 && data1.Cf14.Equals(data2.Cf14) && data1.Cf15.Cmp(data2.Cf15) == 0 && data1.Cf16 == data2.Cf16
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf18 bool
	Cf19 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf18 bool, Cf19 bool) D2 {
	return D2{D2_DC8{Cf18, Cf19}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC7 struct {
	Cf17 _dafny.Map
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf17 _dafny.Map) D2 {
	return D2{D2_DC7{Cf17}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC9 struct {
	Cf20 D2
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf20 D2) D2 {
	return D2{D2_DC9{Cf20}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(false, false)
}

func (_this D2) Dtor_cf18() bool {
	return _this.Get_().(D2_DC8).Cf18
}

func (_this D2) Dtor_cf19() bool {
	return _this.Get_().(D2_DC8).Cf19
}

func (_this D2) Dtor_cf17() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf17
}

func (_this D2) Dtor_cf20() D2 {
	return _this.Get_().(D2_DC9).Cf20
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19 == data2.Cf19
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf20.Equals(data2.Cf20)
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
	Cf22 _dafny.Int
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf22 _dafny.Int) D3 {
	return D3{D3_DC11{Cf22}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC12 struct {
	Cf23 bool
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf23 bool) D3 {
	return D3{D3_DC12{Cf23}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

type D3_DC10 struct {
	Cf21 _dafny.MultiSet
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf21 _dafny.MultiSet) D3 {
	return D3{D3_DC10{Cf21}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(_dafny.Zero)
}

func (_this D3) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) Dtor_cf23() bool {
	return _this.Get_().(D3_DC12).Cf23
}

func (_this D3) Dtor_cf21() _dafny.MultiSet {
	return _this.Get_().(D3_DC10).Cf21
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf23) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf21) + ")"
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
			return ok && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf23 == data2.Cf23
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf25 bool
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf25 bool) D4 {
	return D4{D4_DC14{Cf25}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC13 struct {
	Cf24 _dafny.CodePoint
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf24 _dafny.CodePoint) D4 {
	return D4{D4_DC13{Cf24}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC15 struct {
	Cf26 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf26 D4) D4 {
	return D4{D4_DC15{Cf26}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC14_(false)
}

func (_this D4) Dtor_cf25() bool {
	return _this.Get_().(D4_DC14).Cf25
}

func (_this D4) Dtor_cf24() _dafny.CodePoint {
	return _this.Get_().(D4_DC13).Cf24
}

func (_this D4) Dtor_cf26() D4 {
	return _this.Get_().(D4_DC15).Cf26
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf24) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf25 == data2.Cf25
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf24 == data2.Cf24
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
			return ok && data1.Cf26.Equals(data2.Cf26)
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

type D5_DC16 struct {
	Cf27 _dafny.Map
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf27 _dafny.Map) D5 {
	return D5{D5_DC16{Cf27}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D5) Dtor_cf27() _dafny.Map {
	return _this.Get_().(D5_DC16).Cf27
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf27.Equals(data2.Cf27)
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
	Cf28 _dafny.Array
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf28 _dafny.Array) D6 {
	return D6{D6_DC17{Cf28}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D6) Dtor_cf28() _dafny.Array {
	return _this.Get_().(D6_DC17).Cf28
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf28 == data2.Cf28
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
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_() D7 {
	return D7{D7_DC19{}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
	Cf30 _dafny.Sequence
	Cf31 _dafny.Array
	Cf32 D4
	Cf33 bool
	Cf34 _dafny.Int
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf30 _dafny.Sequence, Cf31 _dafny.Array, Cf32 D4, Cf33 bool, Cf34 _dafny.Int) D7 {
	return D7{D7_DC20{Cf30, Cf31, Cf32, Cf33, Cf34}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC18 struct {
	Cf29 _dafny.Sequence
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf29 _dafny.Sequence) D7 {
	return D7{D7_DC18{Cf29}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_()
}

func (_this D7) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D7_DC20).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D7_DC20).Cf31
}

func (_this D7) Dtor_cf32() D4 {
	return _this.Get_().(D7_DC20).Cf32
}

func (_this D7) Dtor_cf33() bool {
	return _this.Get_().(D7_DC20).Cf33
}

func (_this D7) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D7_DC20).Cf34
}

func (_this D7) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D7_DC18).Cf29
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19"
		}
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf29) + ")"
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
			_, ok := other.Get_().(D7_DC19)
			return ok
		}
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf30.Equals(data2.Cf30) && data1.Cf31 == data2.Cf31 && data1.Cf32.Equals(data2.Cf32) && data1.Cf33 == data2.Cf33 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf29.Equals(data2.Cf29)
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
	Cf36 _dafny.Int
	Cf37 bool
	Cf38 _dafny.CodePoint
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf36 _dafny.Int, Cf37 bool, Cf38 _dafny.CodePoint) D8 {
	return D8{D8_DC22{Cf36, Cf37, Cf38}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC23 struct {
	Cf39 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf39 bool) D8 {
	return D8{D8_DC23{Cf39}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC21 struct {
	Cf35 _dafny.Sequence
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf35 _dafny.Sequence) D8 {
	return D8{D8_DC21{Cf35}}
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
	return Companion_D8_.Create_DC22_(_dafny.Zero, false, _dafny.CodePoint('D'))
}

func (_this D8) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf36
}

func (_this D8) Dtor_cf37() bool {
	return _this.Get_().(D8_DC22).Cf37
}

func (_this D8) Dtor_cf38() _dafny.CodePoint {
	return _this.Get_().(D8_DC22).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC23).Cf39
}

func (_this D8) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D8_DC21).Cf35
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
			return "D8.DC22" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf39) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf39 == data2.Cf39
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf35.Equals(data2.Cf35)
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

type D9_DC25 struct {
	Cf41 _dafny.Map
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf41 _dafny.Map) D9 {
	return D9{D9_DC25{Cf41}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf41() _dafny.Map {
	return _this.Get_().(D9_DC25).Cf41
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
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
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf41.Equals(data2.Cf41)
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
	Cf42 _dafny.Array
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf42 _dafny.Array) D10 {
	return D10{D10_DC26{Cf42}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf42() _dafny.Array {
	return _this.Get_().(D10_DC26).Cf42
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf42 == data2.Cf42
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

type D11_DC27 struct {
	Cf43 _dafny.Array
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf43 _dafny.Array) D11 {
	return D11{D11_DC27{Cf43}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

func (CompanionStruct_D11_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D11) Dtor_cf43() _dafny.Array {
	return _this.Get_().(D11_DC27).Cf43
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf43) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf43 == data2.Cf43
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
	Cf45 _dafny.Sequence
	Cf46 _dafny.Int
	Cf47 bool
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf45 _dafny.Sequence, Cf46 _dafny.Int, Cf47 bool) D12 {
	return D12{D12_DC29{Cf45, Cf46, Cf47}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC30 struct {
	Cf48 _dafny.Sequence
	Cf49 _dafny.Int
	Cf50 _dafny.CodePoint
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf48 _dafny.Sequence, Cf49 _dafny.Int, Cf50 _dafny.CodePoint) D12 {
	return D12{D12_DC30{Cf48, Cf49, Cf50}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

type D12_DC31 struct {
	Cf51 _dafny.Int
	Cf52 bool
	Cf53 bool
	Cf54 _dafny.Int
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf51 _dafny.Int, Cf52 bool, Cf53 bool, Cf54 _dafny.Int) D12 {
	return D12{D12_DC31{Cf51, Cf52, Cf53, Cf54}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC28 struct {
	Cf44 _dafny.MultiSet
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf44 _dafny.MultiSet) D12 {
	return D12{D12_DC28{Cf44}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC29_(_dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D12) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D12_DC29).Cf45
}

func (_this D12) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D12_DC29).Cf46
}

func (_this D12) Dtor_cf47() bool {
	return _this.Get_().(D12_DC29).Cf47
}

func (_this D12) Dtor_cf48() _dafny.Sequence {
	return _this.Get_().(D12_DC30).Cf48
}

func (_this D12) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D12_DC30).Cf49
}

func (_this D12) Dtor_cf50() _dafny.CodePoint {
	return _this.Get_().(D12_DC30).Cf50
}

func (_this D12) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D12_DC31).Cf51
}

func (_this D12) Dtor_cf52() bool {
	return _this.Get_().(D12_DC31).Cf52
}

func (_this D12) Dtor_cf53() bool {
	return _this.Get_().(D12_DC31).Cf53
}

func (_this D12) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D12_DC31).Cf54
}

func (_this D12) Dtor_cf44() _dafny.MultiSet {
	return _this.Get_().(D12_DC28).Cf44
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D12_DC30:
		{
			return "D12.DC30" + "(" + data.Cf48.VerbatimString(true) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf44) + ")"
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
			return ok && data1.Cf45.Equals(data2.Cf45) && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47 == data2.Cf47
		}
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf48.Equals(data2.Cf48) && data1.Cf49.Cmp(data2.Cf49) == 0 && data1.Cf50 == data2.Cf50
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
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

type D13_DC32 struct {
	Cf55 _dafny.Array
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf55 _dafny.Array) D13 {
	return D13{D13_DC32{Cf55}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D13) Dtor_cf55() _dafny.Array {
	return _this.Get_().(D13_DC32).Cf55
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf55) + ")"
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
			return ok && data1.Cf55 == data2.Cf55
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

type D14_DC34 struct {
	Cf57 bool
}

func (D14_DC34) isD14() {}

func (CompanionStruct_D14_) Create_DC34_(Cf57 bool) D14 {
	return D14{D14_DC34{Cf57}}
}

func (_this D14) Is_DC34() bool {
	_, ok := _this.Get_().(D14_DC34)
	return ok
}

type D14_DC35 struct {
	Cf58 bool
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf58 bool) D14 {
	return D14{D14_DC35{Cf58}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

type D14_DC36 struct {
	Cf59 _dafny.Int
	Cf60 bool
	Cf61 bool
	Cf62 bool
	Cf63 _dafny.Int
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf59 _dafny.Int, Cf60 bool, Cf61 bool, Cf62 bool, Cf63 _dafny.Int) D14 {
	return D14{D14_DC36{Cf59, Cf60, Cf61, Cf62, Cf63}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC33 struct {
	Cf56 T1
}

func (D14_DC33) isD14() {}

func (CompanionStruct_D14_) Create_DC33_(Cf56 T1) D14 {
	return D14{D14_DC33{Cf56}}
}

func (_this D14) Is_DC33() bool {
	_, ok := _this.Get_().(D14_DC33)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC34_(false)
}

func (_this D14) Dtor_cf57() bool {
	return _this.Get_().(D14_DC34).Cf57
}

func (_this D14) Dtor_cf58() bool {
	return _this.Get_().(D14_DC35).Cf58
}

func (_this D14) Dtor_cf59() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf59
}

func (_this D14) Dtor_cf60() bool {
	return _this.Get_().(D14_DC36).Cf60
}

func (_this D14) Dtor_cf61() bool {
	return _this.Get_().(D14_DC36).Cf61
}

func (_this D14) Dtor_cf62() bool {
	return _this.Get_().(D14_DC36).Cf62
}

func (_this D14) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf63
}

func (_this D14) Dtor_cf56() T1 {
	return _this.Get_().(D14_DC33).Cf56
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC34:
		{
			return "D14.DC34" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D14_DC33:
		{
			return "D14.DC33" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC34:
		{
			data2, ok := other.Get_().(D14_DC34)
			return ok && data1.Cf57 == data2.Cf57
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf58 == data2.Cf58
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf59.Cmp(data2.Cf59) == 0 && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62 && data1.Cf63.Cmp(data2.Cf63) == 0
		}
	case D14_DC33:
		{
			data2, ok := other.Get_().(D14_DC33)
			return ok && _dafny.AreEqual(data1.Cf56, data2.Cf56)
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

type D15_DC38 struct {
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_() D15 {
	return D15{D15_DC38{}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

type D15_DC37 struct {
	Cf64 _dafny.Map
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf64 _dafny.Map) D15 {
	return D15{D15_DC37{Cf64}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC38_()
}

func (_this D15) Dtor_cf64() _dafny.Map {
	return _this.Get_().(D15_DC37).Cf64
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC38:
		{
			return "D15.DC38"
		}
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf64) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC38:
		{
			_, ok := other.Get_().(D15_DC38)
			return ok
		}
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf64.Equals(data2.Cf64)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map
	Fm2(globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState)
	M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool)
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
	Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map
	Fm2(globalState *GlobalState) _dafny.Int
	M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState)
	M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool)
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
	F6   _dafny.Array
	F7   bool
	F10  _dafny.Int
	F11  _dafny.Map
	F13  bool
	F19  _dafny.CodePoint
	F22  bool
	F24  _dafny.Set
	_f0  _dafny.Map
	_f1  _dafny.Array
	_f2  bool
	_f3  _dafny.CodePoint
	_f4  bool
	_f5  _dafny.Array
	_f8  _dafny.MultiSet
	_f9  _dafny.Int
	_f12 _dafny.Map
	_f14 bool
	_f15 bool
	_f16 bool
	_f17 _dafny.Int
	_f18 bool
	_f20 bool
	_f21 _dafny.Array
	_f23 bool
	_f25 _dafny.Sequence
	_f26 bool
	_f27 _dafny.Int
	_f28 _dafny.CodePoint
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F7 = false
	_this.F10 = _dafny.Zero
	_this.F11 = _dafny.EmptyMap
	_this.F13 = false
	_this.F19 = _dafny.CodePoint('D')
	_this.F22 = false
	_this.F24 = _dafny.EmptySet
	_this._f0 = _dafny.EmptyMap
	_this._f1 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f2 = false
	_this._f3 = _dafny.CodePoint('D')
	_this._f4 = false
	_this._f5 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptyMultiSet
	_this._f9 = _dafny.Zero
	_this._f12 = _dafny.EmptyMap
	_this._f14 = false
	_this._f15 = false
	_this._f16 = false
	_this._f17 = _dafny.Zero
	_this._f18 = false
	_this._f20 = false
	_this._f21 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f23 = false
	_this._f25 = _dafny.EmptySeq
	_this._f26 = false
	_this._f27 = _dafny.Zero
	_this._f28 = _dafny.CodePoint('D')
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

func (_this *GlobalState) Ctor__(f0 _dafny.Map, f1 _dafny.Array, f2 bool, f3 _dafny.CodePoint, f4 bool, f5 _dafny.Array, f6 _dafny.Array, f7 bool, f8 _dafny.MultiSet, f9 _dafny.Int, f10 _dafny.Int, f11 _dafny.Map, f12 _dafny.Map, f13 bool, f14 bool, f15 bool, f16 bool, f17 _dafny.Int, f18 bool, f19 _dafny.CodePoint, f20 bool, f21 _dafny.Array, f22 bool, f23 bool, f24 _dafny.Set, f25 _dafny.Sequence, f26 bool, f27 _dafny.Int, f28 _dafny.CodePoint) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f18 = f18
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this).F22 = f22
		(_this)._f23 = f23
		(_this).F24 = f24
		(_this)._f25 = f25
		(_this)._f26 = f26
		(_this)._f27 = f27
		(_this)._f28 = f28
	}
}
func (_this *GlobalState) F0() _dafny.Map {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Array {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.CodePoint {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Array {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F8() _dafny.MultiSet {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F12() _dafny.Map {
	{
		return _this._f12
	}
}
func (_this *GlobalState) F14() bool {
	{
		return _this._f14
	}
}
func (_this *GlobalState) F15() bool {
	{
		return _this._f15
	}
}
func (_this *GlobalState) F16() bool {
	{
		return _this._f16
	}
}
func (_this *GlobalState) F17() _dafny.Int {
	{
		return _this._f17
	}
}
func (_this *GlobalState) F18() bool {
	{
		return _this._f18
	}
}
func (_this *GlobalState) F20() bool {
	{
		return _this._f20
	}
}
func (_this *GlobalState) F21() _dafny.Array {
	{
		return _this._f21
	}
}
func (_this *GlobalState) F23() bool {
	{
		return _this._f23
	}
}
func (_this *GlobalState) F25() _dafny.Sequence {
	{
		return _this._f25
	}
}
func (_this *GlobalState) F26() bool {
	{
		return _this._f26
	}
}
func (_this *GlobalState) F27() _dafny.Int {
	{
		return _this._f27
	}
}
func (_this *GlobalState) F28() _dafny.CodePoint {
	{
		return _this._f28
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f32 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f32 = false
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

func (_this *C0) Ctor__(f32 bool) {
	{
		(_this)._f32 = f32
	}
}
func (_this *C0) Fm13(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(func(_source10 D1) _dafny.Int {
			if _source10.Is_DC5() {
				var _432___mcc_h0 _dafny.Int = _source10.Get_().(D1_DC5).Cf9
				_ = _432___mcc_h0
				var _433___mcc_h1 bool = _source10.Get_().(D1_DC5).Cf10
				_ = _433___mcc_h1
				var _434___mcc_h2 bool = _source10.Get_().(D1_DC5).Cf11
				_ = _434___mcc_h2
				var _435___mcc_h3 _dafny.Map = _source10.Get_().(D1_DC5).Cf12
				_ = _435___mcc_h3
				var _436_cf12 _dafny.Map = _435___mcc_h3
				_ = _436_cf12
				var _437_cf11 bool = _434___mcc_h2
				_ = _437_cf11
				var _438_cf10 bool = _433___mcc_h1
				_ = _438_cf10
				var _439_cf9 _dafny.Int = _432___mcc_h0
				_ = _439_cf9
				return _dafny.IntOfInt64(732)
			} else if _source10.Is_DC6() {
				var _440___mcc_h4 _dafny.Int = _source10.Get_().(D1_DC6).Cf13
				_ = _440___mcc_h4
				var _441___mcc_h5 _dafny.Sequence = _source10.Get_().(D1_DC6).Cf14
				_ = _441___mcc_h5
				var _442___mcc_h6 _dafny.Int = _source10.Get_().(D1_DC6).Cf15
				_ = _442___mcc_h6
				var _443___mcc_h7 bool = _source10.Get_().(D1_DC6).Cf16
				_ = _443___mcc_h7
				var _444_cf16 bool = _443___mcc_h7
				_ = _444_cf16
				var _445_cf15 _dafny.Int = _442___mcc_h6
				_ = _445_cf15
				var _446_cf14 _dafny.Sequence = _441___mcc_h5
				_ = _446_cf14
				var _447_cf13 _dafny.Int = _440___mcc_h4
				_ = _447_cf13
				return (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_445_cf15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_444_cf16, _447_cf13)).Cardinality()))
			} else {
				var _448___mcc_h8 _dafny.Map = _source10.Get_().(D1_DC4).Cf8
				_ = _448___mcc_h8
				var _449_cf8 _dafny.Map = _448___mcc_h8
				_ = _449_cf8
				return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(632), (_dafny.SetOf(_dafny.IntOfInt64(339))).Cardinality())
			}
		}(Companion_D1_.Create_DC6_((Companion_D0_.Create_DC2_(_dafny.IntOfInt64(891), (_this).F32())).Dtor_cf5(), _dafny.SeqOf(_dafny.SeqOf((_this).F32()), _dafny.SeqOf(!((_this).F32()), (_this).F32())), _dafny.IntOfInt64(124), (_this).F32())))
	}
}
func (_this *C0) F32() bool {
	{
		return _this._f32
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
func (_this *C1) Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false)), !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_450_i0 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true, false, false, true, true)).Cardinality())).Cardinality(), false)
		})), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-535), !(true))))
	}
}
func (_this *C1) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(778))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_451_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-566), _dafny.IntOfInt64(-96)), (_dafny.IntOfInt64(964)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality())), (_dafny.IntOfInt64(68)).Times(_dafny.IntOfInt64(-443)))).Cardinality()))
	}
}
func (_this *C1) Fm21(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-402))).Cardinality(), _dafny.IntOfInt64(-703), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.MultiSetOf(true, true)).Cardinality())).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thqpyy")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))).Cardinality()))
	}
}
func (_this *C1) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		if p1 {
			var _452_v0 _dafny.Array
			_ = _452_v0
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_6
			var _nw94 _dafny.Array
			_ = _nw94
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw94 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Int = func(_453_i0 _dafny.Int) _dafny.Int {
					return (_453_i0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ecpxkbg")).Cardinality()))
				}
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw94 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw94).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw94).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_452_v0 = _nw94
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _index88
			(_452_v0).ArraySet1(p0, (_index88).Int())
			var _454_v1 _dafny.MultiSet
			_ = _454_v1
			_454_v1 = _dafny.MultiSetOf(_dafny.IntOfInt64(-622), p2, p0, p2)
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_452_v0), 0))
			_ = _index89
			(_452_v0).ArraySet1((func() _dafny.Int {
				if (_454_v1).Contains(p0) {
					return (_454_v1).Multiplicity(p0)
				}
				return p0
			})(), (_index89).Int())
			var _455_v2 _dafny.Array
			_ = _455_v2
			var _nwElement0_25 bool = (p1) == (p1)
			_ = _nwElement0_25
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(2))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_25, 0)
			(_nw95).ArraySet1(p1, 1)
			_455_v2 = _nw95
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))
			_ = _index90
			(_455_v2).ArraySet1(p3, (_index90).Int())
			var _456_v3 _dafny.Sequence
			_ = _456_v3
			_456_v3 = _dafny.SeqOf(p0, p0)
			var _457_v4 _dafny.Set
			_ = _457_v4
			_457_v4 = _dafny.SetOf(_dafny.IntOfUint32((_456_v3).Cardinality()), p2)
			var _458_v5 _dafny.Sequence
			_ = _458_v5
			_458_v5 = _dafny.SeqOf(p0, p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, false)).Cardinality(), (_dafny.Zero).Minus((_457_v4).Cardinality()))
			var _459_v6 _dafny.Set
			_ = _459_v6
			_459_v6 = _dafny.SetOf(p1, !(false), p1, p3)
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _index91
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_452_v0), 0))
			_ = _index92
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))
			_ = _index93
			var _rhs82 _dafny.Int = p0
			_ = _rhs82
			var _rhs83 _dafny.Int = Companion_Default___.Fm0(p0, globalState)
			_ = _rhs83
			var _rhs84 _dafny.Int = (func() _dafny.Int {
				if p1 {
					return p0
				}
				return (_456_v3).Select((Companion_Default___.SafeIndex((_459_v6).Cardinality(), _dafny.IntOfUint32((_456_v3).Cardinality()))).Uint32()).(_dafny.Int)
			})()
			_ = _rhs84
			var _rhs85 bool = p1
			_ = _rhs85
			var _lhs68 *GlobalState = globalState
			_ = _lhs68
			var _lhs69 _dafny.Array = _452_v0
			_ = _lhs69
			var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _lhs70
			var _lhs71 _dafny.Array = _452_v0
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_452_v0), 0))
			_ = _lhs72
			var _lhs73 _dafny.Array = _455_v2
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))
			_ = _lhs74
			_lhs68.F10 = _rhs82
			(_lhs69).ArraySet1(_rhs83, (_lhs70).Int())
			(_lhs71).ArraySet1(_rhs84, (_lhs72).Int())
			(_lhs73).ArraySet1(_rhs85, (_lhs74).Int())
			var _460_v7 _dafny.Map
			_ = _460_v7
			_460_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_455_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))).Int()).(bool), p3)
			var _461_v8 _dafny.Sequence
			_ = _461_v8
			_461_v8 = _dafny.SeqOf((func() bool {
				if (_460_v7).Contains(p3) {
					return (_460_v7).Get(p3).(bool)
				}
				return true
			})(), !(p1) || (p3), p1)
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _index94
			var _rhs86 _dafny.Int = (Companion_Default___.SafeDivisionInt(p0, Companion_Default___.Fm0(p0, globalState))).Plus(p0)
			_ = _rhs86
			var _rhs87 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_461_v8, _dafny.Companion_Sequence_.Concatenate(_461_v8, _461_v8))
			_ = _rhs87
			var _rhs88 _dafny.Array = (func() _dafny.Array {
				if false {
					return _452_v0
				}
				return _452_v0
			})()
			_ = _rhs88
			var _lhs75 _dafny.Array = _452_v0
			_ = _lhs75
			var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _lhs76
			(_lhs75).ArraySet1(_rhs86, (_lhs76).Int())
			_461_v8 = _rhs87
			_452_v0 = _rhs88
			if (_this).Fm21(p0, globalState) {
				var _462_v9 _dafny.CodePoint
				_ = _462_v9
				_462_v9 = _dafny.CodePoint('w')
				(globalState).F19 = _462_v9
				(globalState).F13 = true
				var _463_v10 *C0
				_ = _463_v10
				var _nw96 *C0 = New_C0_()
				_ = _nw96
				_nw96.Ctor__(p3)
				_463_v10 = _nw96
				(globalState).F22 = p3
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))
				_ = _index95
				(_455_v2).ArraySet1((_this).Fm21((_dafny.Zero).Minus(p2), globalState), (_index95).Int())
			} else {
				_461_v8 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, (_this).Fm21((_dafny.Zero).Minus((_452_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))).Int()).(_dafny.Int)), globalState), p1), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm22((_455_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_455_v2), 0))).Int()).(bool), globalState), _461_v8))
				var _464_v11 _dafny.Sequence
				_ = _464_v11
				_464_v11 = _dafny.UnicodeSeqOfUtf8Bytes("xfsnymi")
				var _465_v12 D0
				_ = _465_v12
				_465_v12 = Companion_D0_.Create_DC0_(_464_v11)
				(globalState).F10 = _dafny.IntOfUint32(((_465_v12).Dtor_cf0()).Cardinality())
				(globalState).F13 = p1
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
				_ = _nw97
				_452_v0 = _nw97
				(globalState).F7 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_461_v8, _461_v8), p1)
			}
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_452_v0), 0))
			_ = _index96
			(_452_v0).ArraySet1(((_dafny.Zero).Minus(p0)).Minus(p2), (_index96).Int())
			(globalState).F13 = p1
		} else {
			var _466_v13 _dafny.MultiSet
			_ = _466_v13
			_466_v13 = _dafny.MultiSetOf(p1, p1, p3, p1, p1)
			var _467_v14 _dafny.Sequence
			_ = _467_v14
			_467_v14 = _dafny.SeqOf(false, (_this).Fm21(p0, globalState), p1, p1)
			var _468_v15 _dafny.Set
			_ = _468_v15
			_468_v15 = _dafny.SetOf(p1)
			var _469_v16 _dafny.Map
			_ = _469_v16
			_469_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _470_v17 _dafny.Sequence
			_ = _470_v17
			_470_v17 = _dafny.SeqOf(_467_v14)
			var _471_v18 D1
			_ = _471_v18
			_471_v18 = Companion_D1_.Create_DC6_(p2, _470_v17, p2, p3)
			var _472_v19 _dafny.Sequence
			_ = _472_v19
			_472_v19 = _dafny.UnicodeSeqOfUtf8Bytes("sj")
			var _473_v20 _dafny.Array
			_ = _473_v20
			var _nwElement0_26 _dafny.MultiSet = _466_v13
			_ = _nwElement0_26
			var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(24))
			_ = _nw98
			(_nw98).ArraySet1(_nwElement0_26, 0)
			(_nw98).ArraySet1(_466_v13, 1)
			(_nw98).ArraySet1(_466_v13, 2)
			(_nw98).ArraySet1((func() _dafny.MultiSet {
				if (_467_v14).Select((Companion_Default___.SafeIndex((_468_v15).Cardinality(), _dafny.IntOfUint32((_467_v14).Cardinality()))).Uint32()).(bool) {
					return _466_v13
				}
				return _dafny.MultiSetOf((_467_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_469_v16), p0, (_471_v18).Dtor_cf15(), globalState)).Cardinality()), _dafny.IntOfUint32((_467_v14).Cardinality()))).Uint32()).(bool), p3, false)
			})(), 3)
			(_nw98).ArraySet1(_466_v13, 4)
			(_nw98).ArraySet1((func() _dafny.MultiSet {
				if p3 {
					return _466_v13
				}
				return _466_v13
			})(), 5)
			(_nw98).ArraySet1(_dafny.MultiSetFromSeq(_467_v14), 6)
			(_nw98).ArraySet1((_466_v13).Union(_466_v13), 7)
			(_nw98).ArraySet1(Companion_Default___.Fm24(_dafny.IntOfUint32((_472_v19).Cardinality()), p3, globalState), 8)
			(_nw98).ArraySet1((_466_v13).Union(_466_v13), 9)
			(_nw98).ArraySet1(_466_v13, 10)
			(_nw98).ArraySet1(_466_v13, 11)
			(_nw98).ArraySet1((_dafny.MultiSetFromSeq(_467_v14)).Difference(_466_v13), 12)
			(_nw98).ArraySet1(_466_v13, 13)
			(_nw98).ArraySet1(_466_v13, 14)
			(_nw98).ArraySet1(_466_v13, 15)
			(_nw98).ArraySet1(_466_v13, 16)
			(_nw98).ArraySet1(_466_v13, 17)
			(_nw98).ArraySet1(_dafny.MultiSetOf(false, p3), 18)
			(_nw98).ArraySet1(_466_v13, 19)
			(_nw98).ArraySet1(_466_v13, 20)
			(_nw98).ArraySet1((_466_v13).Difference(_466_v13), 21)
			(_nw98).ArraySet1(_466_v13, 22)
			(_nw98).ArraySet1(_466_v13, 23)
			_473_v20 = _nw98
			_473_v20 = _473_v20
			var _474_v21 _dafny.Sequence
			_ = _474_v21
			_474_v21 = _dafny.SeqOf(_dafny.IntOfInt64(228))
			_474_v21 = _dafny.Companion_Sequence_.Concatenate(_474_v21, _474_v21)
			var _475_v22 _dafny.CodePoint
			_ = _475_v22
			_475_v22 = _dafny.CodePoint('n')
			(globalState).F19 = _475_v22
			if (p0).Cmp(_dafny.IntOfUint32((_474_v21).Cardinality())) == 0 {
				var _476_v23 _dafny.Array
				_ = _476_v23
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw99
				_476_v23 = _nw99
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))
				_ = _index97
				(_476_v23).ArraySet1(false, (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))
				_ = _index98
				var _rhs89 bool = p3
				_ = _rhs89
				var _rhs90 _dafny.CodePoint = _475_v22
				_ = _rhs90
				var _rhs91 bool = ((func() bool {
					if p1 {
						return p1
					}
					return p1
				})()) == (p1)
				_ = _rhs91
				var _lhs77 _dafny.Array = _476_v23
				_ = _lhs77
				var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))
				_ = _lhs78
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				(_lhs77).ArraySet1(_rhs89, (_lhs78).Int())
				_475_v22 = _rhs90
				_lhs79.F22 = _rhs91
				var _477_v24 D2
				_ = _477_v24
				_477_v24 = Companion_D2_.Create_DC7_(_469_v16)
				var _478_v25 _dafny.Map
				_ = _478_v25
				_478_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_476_v23, _472_v19)
				var _479_v26 _dafny.Array
				_ = _479_v26
				var _nwElement0_27 _dafny.Sequence = Companion_Default___.Fm23(_477_v24, (_dafny.Zero).Minus(p2), p2, globalState)
				_ = _nwElement0_27
				var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(16))
				_ = _nw100
				(_nw100).ArraySet1(_nwElement0_27, 0)
				(_nw100).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(400))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_480_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_481_i1 _dafny.Int) _dafny.CodePoint {
						return _480_v22
					}
				})(_475_v22))), 1)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_472_v19, _472_v19), 2)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_472_v19, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(945))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}((func(_482_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_483_i2 _dafny.Int) _dafny.CodePoint {
						return _482_v22
					}
				})(_475_v22)))), 3)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_472_v19, _472_v19), 4)
				(_nw100).ArraySet1((func() _dafny.Sequence {
					if (_478_v25).Contains(_476_v23) {
						return (_478_v25).Get(_476_v23).(_dafny.Sequence)
					}
					return _472_v19
				})(), 5)
				(_nw100).ArraySet1(_472_v19, 6)
				(_nw100).ArraySet1(_472_v19, 7)
				(_nw100).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-190))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}((func(_484_v19 _dafny.Sequence) func(_dafny.Int) _dafny.CodePoint {
					return func(_485_i3 _dafny.Int) _dafny.CodePoint {
						return (_484_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_484_v19).Cardinality()), _dafny.IntOfUint32((_484_v19).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_472_v19))), 8)
				(_nw100).ArraySet1(_472_v19, 9)
				(_nw100).ArraySet1(_472_v19, 10)
				(_nw100).ArraySet1(Companion_Default___.Fm23(_477_v24, _dafny.IntOfInt64(-323), p2, globalState), 11)
				(_nw100).ArraySet1(_dafny.Companion_Sequence_.Update(_472_v19, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_474_v21).Cardinality()), _dafny.IntOfUint32((_472_v19).Cardinality()))).Uint32(), _475_v22), 12)
				(_nw100).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ubj"), 13)
				(_nw100).ArraySet1(_472_v19, 14)
				(_nw100).ArraySet1(_472_v19, 15)
				_479_v26 = _nw100
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))
				_ = _index99
				(_479_v26).ArraySet1(_472_v19, (_index99).Int())
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))
				_ = _index100
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))
				_ = _index101
				var _rhs92 bool = (_this).Fm21(_dafny.IntOfInt64(532), globalState)
				_ = _rhs92
				var _rhs93 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fye"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("uerjb"), _472_v19))
				_ = _rhs93
				var _rhs94 bool = false
				_ = _rhs94
				var _lhs80 _dafny.Array = _476_v23
				_ = _lhs80
				var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))
				_ = _lhs81
				var _lhs82 _dafny.Array = _479_v26
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))
				_ = _lhs83
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				(_lhs80).ArraySet1(_rhs92, (_lhs81).Int())
				(_lhs82).ArraySet1(_rhs93, (_lhs83).Int())
				_lhs84.F22 = _rhs94
				(globalState).F10 = p0
				var _486_v27 _dafny.Array
				_ = _486_v27
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(17))
				_ = _nw101
				_486_v27 = _nw101
				var _487_v28 _dafny.Map
				_ = _487_v28
				_487_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_476_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))).Int()).(bool))
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_486_v27), 0))
				_ = _index102
				(_486_v27).ArraySet1(_487_v28, (_index102).Int())
				var _488_v29 _dafny.Array
				_ = _488_v29
				var _nw102 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
				_ = _nw102
				_488_v29 = _nw102
				var _489_v30 *C0
				_ = _489_v30
				var _nw103 *C0 = New_C0_()
				_ = _nw103
				_nw103.Ctor__((_476_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))).Int()).(bool))
				_489_v30 = _nw103
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_488_v29), 0))
				_ = _index103
				(_488_v29).ArraySet1(_489_v30, (_index103).Int())
				var _490_v31 _dafny.Array
				_ = _490_v31
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_7
				var _nw104 _dafny.Array
				_ = _nw104
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw104 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.Int = (func(_491_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_492_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_492_i4, _491_p0)
						}
					})(p0)
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw104 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw104).ArraySet1(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw104).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_490_v31 = _nw104
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_490_v31), 0))
				_ = _index104
				(_490_v31).ArraySet1(Companion_Default___.SafeDivisionInt(p2, p0), (_index104).Int())
				var _493_v32 D3
				_ = _493_v32
				_493_v32 = Companion_D3_.Create_DC10_(_dafny.MultiSetFromSeq(_467_v14))
				var _494_v33 _dafny.Map
				_ = _494_v33
				_494_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_479_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))).Int()).(_dafny.Sequence), _490_v31)
				var _495_v34 _dafny.Array
				_ = _495_v34
				_495_v34 = _490_v31
				var _496_v35 _dafny.Sequence
				_ = _496_v35
				_496_v35 = _dafny.SeqOf(_490_v31)
				var _497_v37 _dafny.Array
				_ = _497_v37
				var _nwElement0_28 _dafny.Array = _490_v31
				_ = _nwElement0_28
				var _nw105 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(14))
				_ = _nw105
				(_nw105).ArraySet1(_nwElement0_28, 0)
				(_nw105).ArraySet1(_490_v31, 1)
				(_nw105).ArraySet1(_490_v31, 2)
				(_nw105).ArraySet1(_490_v31, 3)
				(_nw105).ArraySet1(_490_v31, 4)
				(_nw105).ArraySet1((func() _dafny.Array {
					if (_this).Fm21(p2, globalState) {
						return (func() _dafny.Array {
							if (_494_v33).Contains(_472_v19) {
								return (_494_v33).Get(_472_v19).(_dafny.Array)
							}
							return _490_v31
						})()
					}
					return _490_v31
				})(), 5)
				(_nw105).ArraySet1(_490_v31, 6)
				(_nw105).ArraySet1(_490_v31, 7)
				(_nw105).ArraySet1(_490_v31, 8)
				(_nw105).ArraySet1((_495_v34), 9)
				(_nw105).ArraySet1((_496_v35).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter19 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(-571))).Elements()); ; {
						_compr_18, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _498_v36 _dafny.Int
						_498_v36 = interface{}(_compr_18).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(-571)), _498_v36) {
							_coll18.Add((_498_v36).Times((_474_v21).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_474_v21).Cardinality()))).Uint32()).(_dafny.Int)), p0)
						}
					}
					return _coll18.ToMap()
				}()).Cardinality()), _dafny.IntOfUint32((_496_v35).Cardinality()))).Uint32()).(_dafny.Array), 10)
				(_nw105).ArraySet1(_490_v31, 11)
				(_nw105).ArraySet1(_490_v31, 12)
				(_nw105).ArraySet1(_490_v31, 13)
				_497_v37 = _nw105
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_486_v27), 0))
				_ = _index105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_488_v29), 0))
				_ = _index106
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_490_v31), 0))
				_ = _index107
				var _rhs95 _dafny.Map = func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(207), _dafny.IntOfInt64(900))); ; {
						_compr_19, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _499_v38 _dafny.Int
						_499_v38 = interface{}(_compr_19).(_dafny.Int)
						if ((_dafny.IntOfInt64(207)).Cmp(_499_v38) <= 0) && ((_499_v38).Cmp(_dafny.IntOfInt64(900)) < 0) {
							_coll19.Add((_499_v38).Plus(p0), (p0).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_474_v21).Cardinality())))) == 0)
						}
					}
					return _coll19.ToMap()
				}()
				_ = _rhs95
				var _rhs96 *C0 = _489_v30
				_ = _rhs96
				var _rhs97 _dafny.Int = (p2).Plus(p2)
				_ = _rhs97
				var _rhs98 D3 = _493_v32
				_ = _rhs98
				var _rhs99 _dafny.Array = (func() _dafny.Array {
					if (p3) == (!((_476_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))).Int()).(bool))) {
						return _497_v37
					}
					return _497_v37
				})()
				_ = _rhs99
				var _lhs85 _dafny.Array = _486_v27
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.ArrayLen((_486_v27), 0))
				_ = _lhs86
				var _lhs87 _dafny.Array = _488_v29
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_488_v29), 0))
				_ = _lhs88
				var _lhs89 _dafny.Array = _490_v31
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_490_v31), 0))
				_ = _lhs90
				(_lhs85).ArraySet1(_rhs95, (_lhs86).Int())
				(_lhs87).ArraySet1(_rhs96, (_lhs88).Int())
				(_lhs89).ArraySet1(_rhs97, (_lhs90).Int())
				_493_v32 = _rhs98
				_497_v37 = _rhs99
				var _500_v39 _dafny.Map
				_ = _500_v39
				_500_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_476_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))).Int()).(bool), (_490_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(648), _dafny.ArrayLen((_490_v31), 0))).Int()).(_dafny.Int))
				var _501_v40 _dafny.Map
				_ = _501_v40
				_501_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_476_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_476_v23), 0))).Int()).(bool), _500_v39)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))
				_ = _index108
				(_479_v26).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_472_v19, _dafny.Companion_Sequence_.Concatenate((_479_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))).Int()).(_dafny.Sequence), (_479_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))).Int()).(_dafny.Sequence))), (Companion_Default___.SafeIndex(((_501_v40).Merge(_501_v40)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_472_v19, _dafny.Companion_Sequence_.Concatenate((_479_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))).Int()).(_dafny.Sequence), (_479_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(923), _dafny.ArrayLen((_479_v26), 0))).Int()).(_dafny.Sequence)))).Cardinality()))).Uint32(), _475_v22), (_index108).Int())
			} else {
				var _502_v41 _dafny.Array
				_ = _502_v41
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_8
				var _nw106 _dafny.Array
				_ = _nw106
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw106 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) D1 = (func(_503_v13 _dafny.MultiSet, _504_p1 bool) func(_dafny.Int) D1 {
						return func(_505_i5 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_503_v13, _504_p1))
						}
					})(_466_v13, p1)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw106 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw106).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw106).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_502_v41 = _nw106
				var _506_v43 _dafny.Map
				_ = _506_v43
				_506_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_466_v13, p0)
				var _507_v44 D1
				_ = _507_v44
				_507_v44 = Companion_D1_.Create_DC4_(func() _dafny.Map {
					var _coll20 = _dafny.NewMapBuilder()
					_ = _coll20
					for _iter21 := _dafny.Iterate((_506_v43).Keys().Elements()); ; {
						_compr_20, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _508_v42 _dafny.MultiSet
						_508_v42 = interface{}(_compr_20).(_dafny.MultiSet)
						if (_506_v43).Contains(_508_v42) {
							_coll20.Add(_508_v42, p3)
						}
					}
					return _coll20.ToMap()
				}())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(81), _dafny.ArrayLen((_502_v41), 0))
				_ = _index109
				(_502_v41).ArraySet1(_507_v44, (_index109).Int())
				var _509_v45 _dafny.Map
				_ = _509_v45
				_509_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_466_v13, p3)
				var _pat_let_tv6 = p1
				_ = _pat_let_tv6
				var _pat_let_tv7 = p1
				_ = _pat_let_tv7
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(81), _dafny.ArrayLen((_502_v41), 0))
				_ = _index110
				(_502_v41).ArraySet1(func(_pat_let6_0 D1) D1 {
					return func(_510_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let7_0 _dafny.Map) D1 {
							return func(_511_dt__update_hcf8_h0 _dafny.Map) D1 {
								return Companion_D1_.Create_DC4_(_511_dt__update_hcf8_h0)
							}(_pat_let7_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_pat_let_tv6), _pat_let_tv7))
					}(_pat_let6_0)
				}(Companion_D1_.Create_DC4_(_509_v45)), (_index110).Int())
				var _512_v46 _dafny.Array
				_ = _512_v46
				var _len0_9 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_9
				var _nw107 _dafny.Array
				_ = _nw107
				if _len0_9.Cmp(_dafny.Zero) == 0 {
					_nw107 = _dafny.NewArray(_len0_9)
				} else {
					var _init9 func(_dafny.Int) _dafny.Int = func(_513_i6 _dafny.Int) _dafny.Int {
						return (_513_i6).Times(_dafny.IntOfInt64(-650))
					}
					_ = _init9
					var _element0_9 = _init9(_dafny.Zero)
					_ = _element0_9
					_nw107 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
					(_nw107).ArraySet1(_element0_9, 0)
					var _nativeLen0_9 = (_len0_9).Int()
					_ = _nativeLen0_9
					for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
						(_nw107).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
					}
				}
				_512_v46 = _nw107
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _index111
				(_512_v46).ArraySet1(p2, (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _index112
				var _rhs100 _dafny.Int = p2
				_ = _rhs100
				var _rhs101 _dafny.CodePoint = _475_v22
				_ = _rhs101
				var _rhs102 _dafny.Int = (_dafny.Zero).Minus(p2)
				_ = _rhs102
				var _lhs91 _dafny.Array = _512_v46
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				var _lhs94 *GlobalState = globalState
				_ = _lhs94
				(_lhs91).ArraySet1(_rhs100, (_lhs92).Int())
				_lhs93.F19 = _rhs101
				_lhs94.F10 = _rhs102
				_471_v18 = (func() D1 {
					if false {
						return _471_v18
					}
					return _471_v18
				})()
				var _514_v47 _dafny.Array
				_ = _514_v47
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_10
				var _nw108 _dafny.Array
				_ = _nw108
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw108 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Map = (func(_515_p1 bool, _516_v19 _dafny.Sequence, _517_p3 bool) func(_dafny.Int) _dafny.Map {
						return func(_518_i7 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(20), _515_p1), _515_p1)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_516_v19).Cardinality()), false), _517_p3))
						}
					})(p1, _472_v19, p3)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw108 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw108).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw108).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_514_v47 = _nw108
				var _519_v48 _dafny.Map
				_ = _519_v48
				_519_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
				var _520_v49 _dafny.Map
				_ = _520_v49
				_520_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v48, (_this).Fm21((_512_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))).Int()).(_dafny.Int), globalState))
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_514_v47), 0))
				_ = _index113
				(_514_v47).ArraySet1(_520_v49, (_index113).Int())
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _index114
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_514_v47), 0))
				_ = _index115
				var _rhs103 _dafny.Int = p0
				_ = _rhs103
				var _rhs104 _dafny.Map = _520_v49
				_ = _rhs104
				var _rhs105 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_474_v21, _474_v21), _dafny.SeqOf(_dafny.IntOfInt64(208)))
				_ = _rhs105
				var _rhs106 bool = p1
				_ = _rhs106
				var _lhs95 _dafny.Array = _512_v46
				_ = _lhs95
				var _lhs96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _lhs96
				var _lhs97 _dafny.Array = _514_v47
				_ = _lhs97
				var _lhs98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_514_v47), 0))
				_ = _lhs98
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				(_lhs95).ArraySet1(_rhs103, (_lhs96).Int())
				(_lhs97).ArraySet1(_rhs104, (_lhs98).Int())
				_474_v21 = _rhs105
				_lhs99.F13 = _rhs106
				var _521_v50 D0
				_ = _521_v50
				_521_v50 = Companion_D0_.Create_DC2_(Companion_Default___.Fm0(p2, globalState), p1)
				var _522_v51 _dafny.Array
				_ = _522_v51
				var _nwElement0_29 bool = p3
				_ = _nwElement0_29
				var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(15))
				_ = _nw109
				(_nw109).ArraySet1(_nwElement0_29, 0)
				(_nw109).ArraySet1(p3, 1)
				(_nw109).ArraySet1(p3, 2)
				(_nw109).ArraySet1(!(p1), 3)
				(_nw109).ArraySet1(!(!(!(_468_v15).Contains(p3))), 4)
				(_nw109).ArraySet1((_dafny.IntOfInt64(867)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-93))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}((func(_523_p3 bool, _524_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
					return func(_525_i8 _dafny.Int) _dafny.Int {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_p3, _524_v22)).Cardinality()
					}
				})(p3, _475_v22)))).Cardinality())) <= 0, 5)
				(_nw109).ArraySet1((p0).Cmp(p0) >= 0, 6)
				(_nw109).ArraySet1(p1, 7)
				(_nw109).ArraySet1(p1, 8)
				(_nw109).ArraySet1(p3, 9)
				(_nw109).ArraySet1((!(p3)) || (p1), 10)
				(_nw109).ArraySet1(p1, 11)
				(_nw109).ArraySet1(p1, 12)
				(_nw109).ArraySet1(p1, 13)
				(_nw109).ArraySet1((p0).Cmp((_521_v50).Dtor_cf5()) < 0, 14)
				_522_v51 = _nw109
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_522_v51), 0))
				_ = _index116
				(_522_v51).ArraySet1(p1, (_index116).Int())
				var _526_v52 D0
				_ = _526_v52
				_526_v52 = Companion_D0_.Create_DC1_(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)).Cardinality(), p2, _522_v51)
				var _527_v53 D0
				_ = _527_v53
				_527_v53 = Companion_D0_.Create_DC3_(_526_v52)
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_522_v51), 0))
				_ = _index117
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _index118
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _index119
				var _rhs107 bool = p1
				_ = _rhs107
				var _rhs108 D0 = (func() D0 {
					if p3 {
						return _527_v53
					}
					return _527_v53
				})()
				_ = _rhs108
				var _rhs109 _dafny.Int = ((p2).Times((_512_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))).Int()).(_dafny.Int))).Times((_512_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))).Int()).(_dafny.Int))
				_ = _rhs109
				var _rhs110 _dafny.Int = p0
				_ = _rhs110
				var _lhs100 _dafny.Array = _522_v51
				_ = _lhs100
				var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_522_v51), 0))
				_ = _lhs101
				var _lhs102 _dafny.Array = _512_v46
				_ = _lhs102
				var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _lhs103
				var _lhs104 _dafny.Array = _512_v46
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_512_v46), 0))
				_ = _lhs105
				(_lhs100).ArraySet1(_rhs107, (_lhs101).Int())
				_527_v53 = _rhs108
				(_lhs102).ArraySet1(_rhs109, (_lhs103).Int())
				(_lhs104).ArraySet1(_rhs110, (_lhs105).Int())
			}
			var _528_v54 _dafny.Array
			_ = _528_v54
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_11
			var _nw110 _dafny.Array
			_ = _nw110
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw110 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_529_p3 bool) func(_dafny.Int) bool {
					return func(_530_i9 _dafny.Int) bool {
						return _529_p3
					}
				})(p3)
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
			_528_v54 = _nw110
			var _531_v55 D0
			_ = _531_v55
			_531_v55 = Companion_D0_.Create_DC1_(p2, _dafny.IntOfInt64(933), _dafny.IntOfInt64(-728), _528_v54)
			var _532_v56 _dafny.Sequence
			_ = _532_v56
			_532_v56 = _dafny.SeqOf(_531_v55, _531_v55)
			var _533_v57 _dafny.Map
			_ = _533_v57
			_533_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _532_v56)
			var _534_v58 _dafny.Set
			_ = _534_v58
			_534_v58 = _dafny.SetOf(_dafny.SeqOf(_531_v55, _531_v55, _531_v55), _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_533_v57).Contains(p3) {
					return (_533_v57).Get(p3).(_dafny.Sequence)
				}
				return _532_v56
			})(), _532_v56))
			var _535_v59 _dafny.Array
			_ = _535_v59
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw111
			_535_v59 = _nw111
			var _536_v60 _dafny.Array
			_ = _536_v60
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(12))
			_ = _nw112
			_536_v60 = _nw112
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_535_v59), 0))
			_ = _index120
			(_535_v59).ArraySet1(_536_v60, (_index120).Int())
			var _537_v61 _dafny.Array
			_ = _537_v61
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_12
			var _nw113 _dafny.Array
			_ = _nw113
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw113 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.CodePoint = (func(_538_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_539_i10 _dafny.Int) _dafny.CodePoint {
						return _538_v22
					}
				})(_475_v22)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw113 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw113).ArraySet1CodePoint(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw113).ArraySet1CodePoint(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_537_v61 = _nw113
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_537_v61), 0))
			_ = _index121
			(_537_v61).ArraySet1CodePoint(_dafny.CodePoint('e'), (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_535_v59), 0))
			_ = _index122
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_537_v61), 0))
			_ = _index123
			var _rhs111 _dafny.Set = _534_v58
			_ = _rhs111
			var _rhs112 _dafny.Array = _536_v60
			_ = _rhs112
			var _rhs113 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_467_v14, _dafny.Companion_Sequence_.Update(_467_v14, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_472_v19).Cardinality()), _dafny.IntOfUint32((_467_v14).Cardinality()))).Uint32(), p1)), p3)
			_ = _rhs113
			var _rhs114 _dafny.Int = p2
			_ = _rhs114
			var _rhs115 _dafny.CodePoint = _475_v22
			_ = _rhs115
			var _lhs106 _dafny.Array = _535_v59
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_535_v59), 0))
			_ = _lhs107
			var _lhs108 *GlobalState = globalState
			_ = _lhs108
			var _lhs109 *GlobalState = globalState
			_ = _lhs109
			var _lhs110 _dafny.Array = _537_v61
			_ = _lhs110
			var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_537_v61), 0))
			_ = _lhs111
			_534_v58 = _rhs111
			(_lhs106).ArraySet1(_rhs112, (_lhs107).Int())
			_lhs108.F13 = _rhs113
			_lhs109.F10 = _rhs114
			(_lhs110).ArraySet1CodePoint(_rhs115, (_lhs111).Int())
		}
		var _hi1 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(61), _dafny.IntOfInt64(131))
		_ = _hi1
		for _540_i11 := _dafny.IntOfInt64(461); _540_i11.Cmp(_hi1) < 0; _540_i11 = _540_i11.Plus(_dafny.One) {
			var _541_v62 _dafny.Map
			_ = _541_v62
			_541_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			_541_v62 = _541_v62
			var _542_v63 _dafny.Sequence
			_ = _542_v63
			_542_v63 = _dafny.UnicodeSeqOfUtf8Bytes("hrq")
			var _543_v64 _dafny.Sequence
			_ = _543_v64
			_543_v64 = _dafny.SeqOf(_dafny.IntOfUint32((_542_v63).Cardinality()))
			var _544_v65 _dafny.Set
			_ = _544_v65
			_544_v65 = _dafny.SetOf(_dafny.IntOfUint32((_543_v64).Cardinality()))
			var _545_v66 *C0
			_ = _545_v66
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__(!(p1))
			_545_v66 = _nw114
			var _546_v67 _dafny.Sequence
			_ = _546_v67
			_546_v67 = _dafny.SeqOf(_545_v66)
			var _547_v68 _dafny.CodePoint
			_ = _547_v68
			_547_v68 = _dafny.CodePoint('v')
			var _548_v69 _dafny.Set
			_ = _548_v69
			_548_v69 = _dafny.SetOf(_dafny.CodePoint('f'), _547_v68, _547_v68, _547_v68)
			var _549_v70 _dafny.MultiSet
			_ = _549_v70
			_549_v70 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_546_v67, (Companion_Default___.SafeIndex((_548_v69).Cardinality(), _dafny.IntOfUint32((_546_v67).Cardinality()))).Uint32(), _545_v66), _546_v67)
			var _550_v71 _dafny.MultiSet
			_ = _550_v71
			_550_v71 = _dafny.MultiSetOf(p2)
			var _551_v72 _dafny.Array
			_ = _551_v72
			var _nwElement0_30 _dafny.Int = p0
			_ = _nwElement0_30
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(24))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_30, 0)
			(_nw115).ArraySet1(p0, 1)
			(_nw115).ArraySet1(_540_i11, 2)
			(_nw115).ArraySet1(p0, 3)
			(_nw115).ArraySet1((p0).Times((_544_v65).Cardinality()), 4)
			(_nw115).ArraySet1(p2, 5)
			(_nw115).ArraySet1(_540_i11, 6)
			(_nw115).ArraySet1(_dafny.IntOfUint32((_542_v63).Cardinality()), 7)
			(_nw115).ArraySet1(_540_i11, 8)
			(_nw115).ArraySet1(_dafny.IntOfUint32((_542_v63).Cardinality()), 9)
			(_nw115).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm22(p3, globalState)).Cardinality()), 10)
			(_nw115).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_542_v63).Cardinality()), _540_i11), 11)
			(_nw115).ArraySet1(_540_i11, 12)
			(_nw115).ArraySet1((_this).Fm2(globalState), 13)
			(_nw115).ArraySet1((func() _dafny.Int {
				if (_549_v70).Contains(_546_v67) {
					return (_549_v70).Multiplicity(_546_v67)
				}
				return Companion_Default___.Fm0((_550_v71).Cardinality(), globalState)
			})(), 14)
			(_nw115).ArraySet1(p2, 15)
			(_nw115).ArraySet1(p0, 16)
			(_nw115).ArraySet1(p2, 17)
			(_nw115).ArraySet1(_540_i11, 18)
			(_nw115).ArraySet1(p0, 19)
			(_nw115).ArraySet1(Companion_Default___.SafeModuloInt((_545_v66).Fm13(p2, globalState), _dafny.IntOfUint32((_dafny.SeqOf(_547_v68, _547_v68, _547_v68)).Cardinality())), 20)
			(_nw115).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_545_v66).F32(), _547_v68)).Cardinality(), 21)
			(_nw115).ArraySet1(Companion_Default___.SafeDivisionInt((_545_v66).Fm13(p0, globalState), Companion_Default___.Fm0(_540_i11, globalState)), 22)
			(_nw115).ArraySet1(p0, 23)
			_551_v72 = _nw115
			var _552_v73 _dafny.Set
			_ = _552_v73
			_552_v73 = _dafny.SetOf((_545_v66).F32())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _index124
			(_551_v72).ArraySet1(Companion_Default___.SafeModuloInt((_552_v73).Cardinality(), p2), (_index124).Int())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _index125
			(_551_v72).ArraySet1(p2, (_index125).Int())
			var _553_v74 _dafny.Array
			_ = _553_v74
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_13
			var _nw116 _dafny.Array
			_ = _nw116
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw116 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) bool = (func(_554_v66 *C0, _555_i11 _dafny.Int, _556_v72 _dafny.Array) func(_dafny.Int) bool {
					return func(_557_i12 _dafny.Int) bool {
						return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_554_v66).F32(), Companion_D2_.Create_DC7_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_i11, (_556_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_556_v72), 0))).Int()).(_dafny.Int))))).Cardinality()).Cmp(_555_i11) > 0
					}
				})(_545_v66, _540_i11, _551_v72)
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
			_553_v74 = _nw116
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_553_v74), 0))
			_ = _index126
			(_553_v74).ArraySet1(p1, (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _index127
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _index128
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_553_v74), 0))
			_ = _index129
			var _rhs116 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_543_v64).Select((Companion_Default___.SafeIndex(_540_i11, _dafny.IntOfUint32((_543_v64).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(_540_i11), _dafny.IntOfUint32((_542_v63).Cardinality()))).Cardinality()), p2))
			_ = _rhs116
			var _rhs117 _dafny.Int = _540_i11
			_ = _rhs117
			var _rhs118 bool = (_540_i11).Cmp(_540_i11) >= 0
			_ = _rhs118
			var _lhs112 _dafny.Array = _551_v72
			_ = _lhs112
			var _lhs113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _lhs113
			var _lhs114 _dafny.Array = _551_v72
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_551_v72), 0))
			_ = _lhs115
			var _lhs116 _dafny.Array = _553_v74
			_ = _lhs116
			var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_553_v74), 0))
			_ = _lhs117
			(_lhs112).ArraySet1(_rhs116, (_lhs113).Int())
			(_lhs114).ArraySet1(_rhs117, (_lhs115).Int())
			(_lhs116).ArraySet1(_rhs118, (_lhs117).Int())
			var _558_v75 *C0
			_ = _558_v75
			var _nw117 *C0 = New_C0_()
			_ = _nw117
			_nw117.Ctor__((_553_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_553_v74), 0))).Int()).(bool))
			_558_v75 = _nw117
		}
		var _559_v76 _dafny.Array
		_ = _559_v76
		var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
		_ = _nw118
		_559_v76 = _nw118
		var _560_v77 _dafny.Sequence
		_ = _560_v77
		_560_v77 = _dafny.SeqOf(p2)
		var _561_v78 _dafny.Sequence
		_ = _561_v78
		_561_v78 = _dafny.SeqOf((_this).Fm21(p0, globalState))
		var _562_v79 _dafny.MultiSet
		_ = _562_v79
		_562_v79 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(28))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_563_i13 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('l')
		}))).Cardinality()))
		var _564_v80 _dafny.Sequence
		_ = _564_v80
		_564_v80 = _dafny.UnicodeSeqOfUtf8Bytes("cxrrfwi")
		var _565_v81 _dafny.Map
		_ = _565_v81
		_565_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(globalState), _564_v80)
		var _566_v82 _dafny.Array
		_ = _566_v82
		var _nwElement0_31 _dafny.Int = p2
		_ = _nwElement0_31
		var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(10))
		_ = _nw119
		(_nw119).ArraySet1(_nwElement0_31, 0)
		(_nw119).ArraySet1(p2, 1)
		(_nw119).ArraySet1(_dafny.IntOfUint32((_560_v77).Cardinality()), 2)
		(_nw119).ArraySet1(_dafny.IntOfInt64(380), 3)
		(_nw119).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_561_v78).Cardinality()))).Cardinality(), 4)
		(_nw119).ArraySet1(Companion_Default___.Fm0(p0, globalState), 5)
		(_nw119).ArraySet1(p2, 6)
		(_nw119).ArraySet1(_dafny.IntOfInt64(411), 7)
		(_nw119).ArraySet1((func() _dafny.Int {
			if (_562_v79).Contains(p0) {
				return (_562_v79).Multiplicity(p0)
			}
			return _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_565_v81).Contains(p0) {
					return (_565_v81).Get(p0).(_dafny.Sequence)
				}
				return _564_v80
			})()).Cardinality())
		})(), 8)
		(_nw119).ArraySet1(p0, 9)
		_566_v82 = _nw119
		var _567_v83 _dafny.Array
		_ = _567_v83
		var _nwElement0_32 _dafny.Array = _566_v82
		_ = _nwElement0_32
		var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(3))
		_ = _nw120
		(_nw120).ArraySet1(_nwElement0_32, 0)
		(_nw120).ArraySet1(_566_v82, 1)
		(_nw120).ArraySet1(_566_v82, 2)
		_567_v83 = _nw120
		var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_559_v76), 0))
		_ = _index130
		(_559_v76).ArraySet1(_567_v83, (_index130).Int())
		var _568_v84 _dafny.Map
		_ = _568_v84
		_568_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _567_v83)
		var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_559_v76), 0))
		_ = _index131
		(_559_v76).ArraySet1((func() _dafny.Array {
			if (_568_v84).Contains(p1) {
				return (_568_v84).Get(p1).(_dafny.Array)
			}
			return _567_v83
		})(), (_index131).Int())
		var _569_v85 _dafny.Sequence
		_ = _569_v85
		_569_v85 = _dafny.SeqOf(_561_v78, _561_v78)
		var _570_v86 D1
		_ = _570_v86
		_570_v86 = Companion_D1_.Create_DC6_(p2, _569_v85, p0, p3)
		var _571_v87 _dafny.Map
		_ = _571_v87
		_571_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p3)
		var _572_v88 _dafny.Map
		_ = _572_v88
		_572_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_571_v87).Contains(p3) {
				return (_571_v87).Get(p3).(bool)
			}
			return p1
		})(), p1)
		var _573_v89 D0
		_ = _573_v89
		_573_v89 = Companion_D0_.Create_DC2_(p0, p1)
		var _574_v90 _dafny.Array
		_ = _574_v90
		var _nwElement0_33 bool = p3
		_ = _nwElement0_33
		var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(19))
		_ = _nw121
		(_nw121).ArraySet1(_nwElement0_33, 0)
		(_nw121).ArraySet1(p1, 1)
		(_nw121).ArraySet1(p1, 2)
		(_nw121).ArraySet1(p1, 3)
		(_nw121).ArraySet1((_570_v86).Dtor_cf16(), 4)
		(_nw121).ArraySet1(p3, 5)
		(_nw121).ArraySet1(p3, 6)
		(_nw121).ArraySet1((func() bool {
			if (_571_v87).Contains((func() bool {
				if (_572_v88).Contains(p3) {
					return (_572_v88).Get(p3).(bool)
				}
				return p1
			})()) {
				return (_571_v87).Get((func() bool {
					if (_572_v88).Contains(p3) {
						return (_572_v88).Get(p3).(bool)
					}
					return p1
				})()).(bool)
			}
			return p3
		})(), 7)
		(_nw121).ArraySet1(p1, 8)
		(_nw121).ArraySet1(p3, 9)
		(_nw121).ArraySet1(p1, 10)
		(_nw121).ArraySet1((_573_v89).Dtor_cf6(), 11)
		(_nw121).ArraySet1((_561_v78).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_561_v78).Cardinality()))).Uint32()).(bool), 12)
		(_nw121).ArraySet1(p1, 13)
		(_nw121).ArraySet1(p1, 14)
		(_nw121).ArraySet1(p3, 15)
		(_nw121).ArraySet1(p3, 16)
		(_nw121).ArraySet1(!(p1), 17)
		(_nw121).ArraySet1(false, 18)
		_574_v90 = _nw121
		var _575_v91 _dafny.Set
		_ = _575_v91
		_575_v91 = _dafny.SetOf(_574_v90)
		var _576_v92 _dafny.Set
		_ = _576_v92
		_576_v92 = _dafny.SetOf((_560_v77).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cqty")).Cardinality()), _dafny.IntOfUint32((_560_v77).Cardinality()))).Uint32()).(_dafny.Int))
		var _577_v93 _dafny.Sequence
		_ = _577_v93
		_577_v93 = _dafny.SeqOf(_dafny.SetOf(_574_v90, _574_v90, _574_v90), _575_v91)
		var _rhs119 _dafny.Int = (((_576_v92).Union(_dafny.SetOf((_dafny.Zero).Minus(p2)))).Intersection((func() _dafny.Set {
			if p1 {
				return _576_v92
			}
			return _576_v92
		})())).Cardinality()
		_ = _rhs119
		var _rhs120 _dafny.Set = (_577_v93).Select((Companion_Default___.SafeIndex((p0).Times(p2), _dafny.IntOfUint32((_577_v93).Cardinality()))).Uint32()).(_dafny.Set)
		_ = _rhs120
		var _rhs121 _dafny.Int = p0
		_ = _rhs121
		var _lhs118 *GlobalState = globalState
		_ = _lhs118
		var _lhs119 *GlobalState = globalState
		_ = _lhs119
		_lhs118.F10 = _rhs119
		_575_v91 = _rhs120
		_lhs119.F10 = _rhs121
		(globalState).F13 = p3
		var _578_v94 _dafny.CodePoint
		_ = _578_v94
		_578_v94 = _dafny.CodePoint('y')
		var _579_v95 D4
		_ = _579_v95
		_579_v95 = Companion_D4_.Create_DC13_(_578_v94)
		var _pat_let_tv8 = _578_v94
		_ = _pat_let_tv8
		var _pat_let_tv9 = _578_v94
		_ = _pat_let_tv9
		var _pat_let_tv10 = _578_v94
		_ = _pat_let_tv10
		var _rhs122 _dafny.Int = p0
		_ = _rhs122
		var _rhs123 _dafny.Int = p2
		_ = _rhs123
		var _rhs124 _dafny.CodePoint = (func() _dafny.CodePoint {
			if p1 {
				return _578_v94
			}
			return _578_v94
		})()
		_ = _rhs124
		var _rhs125 _dafny.CodePoint = func(_source11 D4) _dafny.CodePoint {
			if _source11.Is_DC14() {
				var _580___mcc_h0 bool = _source11.Get_().(D4_DC14).Cf25
				_ = _580___mcc_h0
				var _581_cf25 bool = _580___mcc_h0
				_ = _581_cf25
				return _pat_let_tv8
			} else if _source11.Is_DC13() {
				var _582___mcc_h1 _dafny.CodePoint = _source11.Get_().(D4_DC13).Cf24
				_ = _582___mcc_h1
				var _583_cf24 _dafny.CodePoint = _582___mcc_h1
				_ = _583_cf24
				return _pat_let_tv9
			} else {
				var _584___mcc_h2 D4 = _source11.Get_().(D4_DC15).Cf26
				_ = _584___mcc_h2
				var _585_cf26 D4 = _584___mcc_h2
				_ = _585_cf26
				return _pat_let_tv10
			}
		}(_579_v95)
		_ = _rhs125
		var _lhs120 *GlobalState = globalState
		_ = _lhs120
		var _lhs121 *GlobalState = globalState
		_ = _lhs121
		var _lhs122 *GlobalState = globalState
		_ = _lhs122
		var _lhs123 *GlobalState = globalState
		_ = _lhs123
		_lhs120.F10 = _rhs122
		_lhs121.F10 = _rhs123
		_lhs122.F19 = _rhs124
		_lhs123.F19 = _rhs125
	}
}
func (_this *C1) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 bool = false
		_ = r1
		var _586_v0 _dafny.Sequence
		_ = _586_v0
		_586_v0 = _dafny.UnicodeSeqOfUtf8Bytes("btqf")
		_586_v0 = _586_v0
		if p2 {
			(globalState).F13 = !(p2)
			var _587_v1 _dafny.Map
			_ = _587_v1
			_587_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(408), !(p2))
			var _588_v2 _dafny.Array
			_ = _588_v2
			var _nwElement0_34 _dafny.Int = (_dafny.Zero).Minus(p1)
			_ = _nwElement0_34
			var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(17))
			_ = _nw122
			(_nw122).ArraySet1(_nwElement0_34, 0)
			(_nw122).ArraySet1((p1).Plus(p1), 1)
			(_nw122).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_586_v0, _586_v0)).Cardinality()), 2)
			(_nw122).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((p1).Times(_dafny.IntOfUint32((_586_v0).Cardinality()))))), 3)
			(_nw122).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), 4)
			(_nw122).ArraySet1((_dafny.Zero).Minus((p1).Times(p1)), 5)
			(_nw122).ArraySet1(p1, 6)
			(_nw122).ArraySet1(p1, 7)
			(_nw122).ArraySet1(p1, 8)
			(_nw122).ArraySet1(p1, 9)
			(_nw122).ArraySet1(p1, 10)
			(_nw122).ArraySet1(p1, 11)
			(_nw122).ArraySet1((_587_v1).Cardinality(), 12)
			(_nw122).ArraySet1(p1, 13)
			(_nw122).ArraySet1(p1, 14)
			(_nw122).ArraySet1((_dafny.Zero).Minus(p1), 15)
			(_nw122).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_586_v0, _586_v0)).Cardinality()), 16)
			_588_v2 = _nw122
			_588_v2 = (func() _dafny.Array {
				if !(p2) {
					return _588_v2
				}
				return _588_v2
			})()
			var _589_v3 _dafny.Array
			_ = _589_v3
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw123
			_589_v3 = _nw123
			var _590_v4 _dafny.Sequence
			_ = _590_v4
			_590_v4 = _dafny.SeqOf(p2, p2, p2)
			var _rhs126 _dafny.Int = (func() _dafny.Int {
				if !(p2) {
					return p1
				}
				return (func() _dafny.Int {
					if (_590_v4).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_590_v4).Cardinality()))).Uint32()).(bool) {
						return p1
					}
					return _dafny.IntOfInt64(-576)
				})()
			})()
			_ = _rhs126
			var _rhs127 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ocl")
			_ = _rhs127
			var _rhs128 _dafny.Array = _589_v3
			_ = _rhs128
			var _rhs129 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(548), p1)
			_ = _rhs129
			var _rhs130 bool = (_this).Fm21((_dafny.Zero).Minus(p1), globalState)
			_ = _rhs130
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			var _lhs125 *GlobalState = globalState
			_ = _lhs125
			var _lhs126 *GlobalState = globalState
			_ = _lhs126
			_lhs124.F10 = _rhs126
			_586_v0 = _rhs127
			_589_v3 = _rhs128
			_lhs125.F10 = _rhs129
			_lhs126.F13 = _rhs130
			if p2 {
				var _591_v5 _dafny.Sequence
				_ = _591_v5
				_591_v5 = _dafny.SeqOf(p1, p1)
				_591_v5 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_591_v5, _dafny.SeqOf(p1, (_dafny.Zero).Minus(p1), p1)), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_591_v5, _dafny.SeqOf(p1, (_dafny.Zero).Minus(p1), p1))).Cardinality()))).Uint32(), (func() _dafny.Int {
					if p2 {
						return p1
					}
					return p1
				})())
				(globalState).F7 = p2
				var _592_v6 _dafny.MultiSet
				_ = _592_v6
				_592_v6 = _dafny.MultiSetOf(p1, _dafny.IntOfUint32((_586_v0).Cardinality()))
				var _593_v7 _dafny.Map
				_ = _593_v7
				_593_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v6, _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), (_dafny.MultiSetOf(true)).Cardinality()))
				var _594_v8 D1
				_ = _594_v8
				_594_v8 = Companion_D1_.Create_DC5_(p1, !(p2), p2, _593_v7)
				var _595_v9 _dafny.Array
				_ = _595_v9
				var _nwElement0_35 D1 = Companion_Default___.Fm25(p2, !(p2), p1, p2, globalState)
				_ = _nwElement0_35
				var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(10))
				_ = _nw124
				(_nw124).ArraySet1(_nwElement0_35, 0)
				(_nw124).ArraySet1(_594_v8, 1)
				(_nw124).ArraySet1(_594_v8, 2)
				(_nw124).ArraySet1(_594_v8, 3)
				(_nw124).ArraySet1(_594_v8, 4)
				(_nw124).ArraySet1(_594_v8, 5)
				(_nw124).ArraySet1(_594_v8, 6)
				(_nw124).ArraySet1(_594_v8, 7)
				(_nw124).ArraySet1(Companion_D1_.Create_DC5_(p1, p2, true, _593_v7), 8)
				(_nw124).ArraySet1(_594_v8, 9)
				_595_v9 = _nw124
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_595_v9), 0))
				_ = _index132
				(_595_v9).ArraySet1(_594_v8, (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_595_v9), 0))
				_ = _index133
				(_595_v9).ArraySet1(_594_v8, (_index133).Int())
				var _596_v10 _dafny.CodePoint
				_ = _596_v10
				var _out2 _dafny.CodePoint
				_ = _out2
				_out2 = (_this).M6((_dafny.IntOfUint32((_591_v5).Cardinality())).Minus(p1), (_dafny.IntOfUint32((_586_v0).Cardinality())).Minus(p1), (_dafny.IntOfUint32((_590_v4).Cardinality())).Times(p1), Companion_Default___.SafeDivisionInt(p1, p1), globalState)
				_596_v10 = _out2
				var _597_v11 _dafny.Set
				_ = _597_v11
				_597_v11 = _dafny.SetOf(true, p2, p2)
				(globalState).F13 = !(!((_597_v11).Contains(p2))) || (p2)
			} else {
				var _598_v12 _dafny.Map
				_ = _598_v12
				_598_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), p1), p0))
				var _599_v13 _dafny.MultiSet
				_ = _599_v13
				_599_v13 = _dafny.MultiSetOf(p2)
				var _600_v14 _dafny.Map
				_ = _600_v14
				_600_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_599_v13, true)
				var _601_v15 D1
				_ = _601_v15
				_601_v15 = Companion_D1_.Create_DC4_(_600_v14)
				var _602_v16 _dafny.Sequence
				_ = _602_v16
				_602_v16 = _dafny.SeqOf(_601_v15, _601_v15)
				var _603_v17 _dafny.CodePoint
				_ = _603_v17
				_603_v17 = _dafny.CodePoint('k')
				var _604_v18 _dafny.Map
				_ = _604_v18
				_604_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_588_v2, _603_v17)
				var _605_v19 _dafny.Set
				_ = _605_v19
				_605_v19 = _dafny.SetOf(p2, p2)
				var _606_v20 _dafny.MultiSet
				_ = _606_v20
				_606_v20 = _dafny.MultiSetOf(_603_v17)
				var _607_v21 _dafny.Map
				_ = _607_v21
				_607_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_606_v20, p2)
				var _608_v22 _dafny.MultiSet
				_ = _608_v22
				_608_v22 = _dafny.MultiSetOf(_dafny.IntOfUint32((_602_v16).Cardinality()), ((_604_v18).Cardinality()).Times(p1), p1, Companion_Default___.SafeDivisionInt((_605_v19).Cardinality(), (_607_v21).Cardinality()), p1)
				var _609_v23 D0
				_ = _609_v23
				_609_v23 = Companion_D0_.Create_DC2_(p1, p2)
				var _610_v24 _dafny.Map
				_ = _610_v24
				_610_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_609_v23).Dtor_cf6(), p2)
				var _rhs131 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if p2 {
						return _590_v4
					}
					return _dafny.Companion_Sequence_.Concatenate(_590_v4, Companion_Default___.Fm22(p2, globalState))
				})()).Cardinality())
				_ = _rhs131
				var _rhs132 _dafny.Map = Companion_Default___.Fm26(p1, (func() bool {
					if (_610_v24).Contains(p2) {
						return (_610_v24).Get(p2).(bool)
					}
					return p2
				})(), (_586_v0).Select((Companion_Default___.SafeIndex((_608_v22).Cardinality(), _dafny.IntOfUint32((_586_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState)
				_ = _rhs132
				var _rhs133 _dafny.MultiSet = _608_v22
				_ = _rhs133
				var _rhs134 bool = true
				_ = _rhs134
				var _rhs135 _dafny.Int = p1
				_ = _rhs135
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 *GlobalState = globalState
				_ = _lhs128
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				_lhs127.F10 = _rhs131
				_598_v12 = _rhs132
				_608_v22 = _rhs133
				_lhs128.F13 = _rhs134
				_lhs129.F10 = _rhs135
				var _611_v25 _dafny.Set
				_ = _611_v25
				_611_v25 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_603_v17), _dafny.Companion_Sequence_.Update(_586_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_586_v0).Cardinality()))).Uint32(), _603_v17)))
				var _612_v26 _dafny.Map
				_ = _612_v26
				_612_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v19, p2)
				var _rhs136 _dafny.Set = Companion_Default___.Fm27(_612_v26, globalState)
				_ = _rhs136
				var _rhs137 bool = p2
				_ = _rhs137
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_611_v25 = _rhs136
				_lhs130.F22 = _rhs137
				var _613_v27 _dafny.Array
				_ = _613_v27
				var _nwElement0_36 bool = p2
				_ = _nwElement0_36
				var _nw125 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(27))
				_ = _nw125
				(_nw125).ArraySet1(_nwElement0_36, 0)
				(_nw125).ArraySet1(p2, 1)
				(_nw125).ArraySet1(p2, 2)
				(_nw125).ArraySet1(p2, 3)
				(_nw125).ArraySet1(p2, 4)
				(_nw125).ArraySet1(p2, 5)
				(_nw125).ArraySet1(p2, 6)
				(_nw125).ArraySet1(p2, 7)
				(_nw125).ArraySet1(p2, 8)
				(_nw125).ArraySet1(false, 9)
				(_nw125).ArraySet1(p2, 10)
				(_nw125).ArraySet1((_this).Fm21(p1, globalState), 11)
				(_nw125).ArraySet1(p2, 12)
				(_nw125).ArraySet1(p2, 13)
				(_nw125).ArraySet1(false, 14)
				(_nw125).ArraySet1(p2, 15)
				(_nw125).ArraySet1(p2, 16)
				(_nw125).ArraySet1(p2, 17)
				(_nw125).ArraySet1(p2, 18)
				(_nw125).ArraySet1(!(p2), 19)
				(_nw125).ArraySet1(p2, 20)
				(_nw125).ArraySet1(!(false), 21)
				(_nw125).ArraySet1(p2, 22)
				(_nw125).ArraySet1(p2, 23)
				(_nw125).ArraySet1(p2, 24)
				(_nw125).ArraySet1((_590_v4).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_590_v4).Cardinality()))).Uint32()).(bool), 25)
				(_nw125).ArraySet1(p2, 26)
				_613_v27 = _nw125
				var _614_v28 _dafny.Sequence
				_ = _614_v28
				_614_v28 = _dafny.SeqOf(_613_v27, _613_v27)
				var _615_v29 _dafny.Map
				_ = _615_v29
				_615_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_614_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_614_v28).Cardinality()))).Uint32()).(_dafny.Array), _dafny.IntOfUint32((_586_v0).Cardinality()))
				var _616_v30 _dafny.Map
				_ = _616_v30
				_616_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v29, _dafny.Companion_Sequence_.Contains(_586_v0, _603_v17))
				_616_v30 = (_616_v30).Update(_615_v29, p2)
				(globalState).F7 = p2
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_589_v3), 0))
				_ = _index134
				(_589_v3).ArraySet1(_613_v27, (_index134).Int())
				var _617_v31 D4
				_ = _617_v31
				_617_v31 = Companion_D4_.Create_DC14_(p2)
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_589_v3), 0))
				_ = _index135
				var _rhs138 _dafny.Array = _613_v27
				_ = _rhs138
				var _rhs139 bool = p2
				_ = _rhs139
				var _rhs140 D4 = _617_v31
				_ = _rhs140
				var _lhs131 _dafny.Array = _589_v3
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_589_v3), 0))
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				(_lhs131).ArraySet1(_rhs138, (_lhs132).Int())
				_lhs133.F22 = _rhs139
				_617_v31 = _rhs140
			}
			(globalState).F10 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_590_v4, Companion_Default___.Fm22(p2, globalState)), _590_v4)).Cardinality())
		} else {
			var _618_v32 _dafny.Set
			_ = _618_v32
			_618_v32 = _dafny.SetOf(_586_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(303))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}(func(_619_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})))
			(globalState).F7 = (_this).Fm21((_618_v32).Cardinality(), globalState)
			_586_v0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-309))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_620_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('s')
			}))
			var _621_v33 _dafny.Set
			_ = _621_v33
			_621_v33 = _dafny.SetOf(p1, p1, p1, p1)
			var _622_v34 _dafny.Array
			_ = _622_v34
			var _nwElement0_37 bool = false
			_ = _nwElement0_37
			var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(19))
			_ = _nw126
			(_nw126).ArraySet1(_nwElement0_37, 0)
			(_nw126).ArraySet1(false, 1)
			(_nw126).ArraySet1((_this).Fm21(_dafny.IntOfInt64(-719), globalState), 2)
			(_nw126).ArraySet1((_621_v33).IsSubsetOf(_621_v33), 3)
			(_nw126).ArraySet1(p2, 4)
			(_nw126).ArraySet1(p2, 5)
			(_nw126).ArraySet1(false, 6)
			(_nw126).ArraySet1(((Companion_Default___.Fm28(p2, globalState)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqstee")).Cardinality())) > 0, 7)
			(_nw126).ArraySet1(!(true), 8)
			(_nw126).ArraySet1(p2, 9)
			(_nw126).ArraySet1(p2, 10)
			(_nw126).ArraySet1((p2) == (p2), 11)
			(_nw126).ArraySet1(true, 12)
			(_nw126).ArraySet1(p2, 13)
			(_nw126).ArraySet1(p2, 14)
			(_nw126).ArraySet1(!(false), 15)
			(_nw126).ArraySet1(p2, 16)
			(_nw126).ArraySet1(p2, 17)
			(_nw126).ArraySet1(p2, 18)
			_622_v34 = _nw126
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))
			_ = _index136
			(_622_v34).ArraySet1((p1).Cmp(p1) <= 0, (_index136).Int())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))
			_ = _index137
			(_622_v34).ArraySet1(p2, (_index137).Int())
			(globalState).F13 = (_622_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))).Int()).(bool)
			var _623_v35 _dafny.Set
			_ = _623_v35
			_623_v35 = _dafny.SetOf((_622_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))).Int()).(bool))
			var _624_v36 _dafny.Sequence
			_ = _624_v36
			_624_v36 = _dafny.SeqOf((_622_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))).Int()).(bool))
			var _625_v37 _dafny.MultiSet
			_ = _625_v37
			_625_v37 = _dafny.MultiSetOf(p1)
			var _626_v38 _dafny.Set
			_ = _626_v38
			_626_v38 = _dafny.SetOf(!(p2), (_623_v35).IsProperSubsetOf(_dafny.SetOf(p2)), (func() bool {
				if !(p2) {
					return (_624_v36).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_624_v36).Cardinality()))).Uint32()).(bool)
				}
				return !(p2)
			})(), (_622_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(839), _dafny.ArrayLen((_622_v34), 0))).Int()).(bool), (_625_v37).IsProperSubsetOf(_625_v37))
			_623_v35 = _626_v38
		}
		var _627_v39 _dafny.Array
		_ = _627_v39
		var _nw127 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw127
		_627_v39 = _nw127
		var _628_v40 _dafny.Sequence
		_ = _628_v40
		_628_v40 = _dafny.SeqOf(p2)
		var _629_v41 D1
		_ = _629_v41
		_629_v41 = Companion_D1_.Create_DC6_(p1, _dafny.SeqOf(_628_v40), p1, p2)
		var _rhs141 _dafny.Array = _627_v39
		_ = _rhs141
		var _rhs142 _dafny.Int = p1
		_ = _rhs142
		var _rhs143 D1 = Companion_Default___.Fm29(globalState)
		_ = _rhs143
		var _lhs134 *GlobalState = globalState
		_ = _lhs134
		_627_v39 = _rhs141
		_lhs134.F10 = _rhs142
		_629_v41 = _rhs143
		var _630_v42 _dafny.Array
		_ = _630_v42
		var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw128
		_630_v42 = _nw128
		_630_v42 = _630_v42
		var _631_v43 _dafny.Array
		_ = _631_v43
		var _nw129 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
		_ = _nw129
		_631_v43 = _nw129
		var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_631_v43), 0))
		_ = _index138
		(_631_v43).ArraySet1(_627_v39, (_index138).Int())
		var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_631_v43), 0))
		_ = _index139
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw130
		(_631_v43).ArraySet1(_nw130, (_index139).Int())
		(globalState).F10 = p1
		var _632_v44 D0
		_ = _632_v44
		_632_v44 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(914), p1, p1, _dafny.ArrayCastTo((_631_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(960), _dafny.ArrayLen((_631_v43), 0))).Int())))
		r0 = (func() D0 {
			if !((p2) == (p2)) {
				return _632_v44
			}
			return _632_v44
		})()
		r1 = (_628_v40).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_628_v40).Cardinality()))).Uint32()).(bool)
		return r0, r1
	}
}
func (_this *C1) M6(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var _633_v0 _dafny.Sequence
		_ = _633_v0
		_633_v0 = _dafny.SeqOf(p2)
		_633_v0 = _633_v0
		var _634_v1 _dafny.Array
		_ = _634_v1
		var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
		_ = _nw131
		_634_v1 = _nw131
		var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
		_ = _index140
		(_634_v1).ArraySet1(p2, (_index140).Int())
		var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
		_ = _index141
		(_634_v1).ArraySet1(p1, (_index141).Int())
		var _635_v2 bool
		_ = _635_v2
		_635_v2 = false
		var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
		_ = _index142
		(_634_v1).ArraySet1((func() _dafny.Int {
			if !(_635_v2) {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(835))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_636_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_637_i0 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_636_p1)
					}
				})(p1)))).Cardinality())
			}
			return p2
		})(), (_index142).Int())
		var _638_v3 _dafny.Sequence
		_ = _638_v3
		_638_v3 = _dafny.SeqOf(_635_v2)
		(globalState).F10 = ((_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_638_v3, (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_638_v3).Cardinality()))).Uint32(), _635_v2)).Cardinality())))
		if (_635_v2) == (_635_v2) {
			var _639_v4 _dafny.Sequence
			_ = _639_v4
			_639_v4 = _dafny.UnicodeSeqOfUtf8Bytes("ffxpkyq")
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _index143
			var _rhs144 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(_635_v2)).Cardinality())).Times(p0)), (_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int))
			_ = _rhs144
			var _rhs145 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_633_v0).Cardinality()), p3))).Plus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_639_v4).Cardinality()), (_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int)))
			_ = _rhs145
			var _rhs146 _dafny.Int = p2
			_ = _rhs146
			var _lhs135 *GlobalState = globalState
			_ = _lhs135
			var _lhs136 *GlobalState = globalState
			_ = _lhs136
			var _lhs137 _dafny.Array = _634_v1
			_ = _lhs137
			var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _lhs138
			_lhs135.F10 = _rhs144
			_lhs136.F10 = _rhs145
			(_lhs137).ArraySet1(_rhs146, (_lhs138).Int())
			var _640_v5 *C0
			_ = _640_v5
			var _nw132 *C0 = New_C0_()
			_ = _nw132
			_nw132.Ctor__(!_dafny.Companion_Sequence_.Equal(_639_v4, _639_v4))
			_640_v5 = _nw132
			var _641_v6 D3
			_ = _641_v6
			_641_v6 = Companion_D3_.Create_DC12_((_640_v5).F32())
			var _642_v7 _dafny.MultiSet
			_ = _642_v7
			_642_v7 = _dafny.MultiSetOf(p0)
			var _643_v8 D1
			_ = _643_v8
			_643_v8 = Companion_D1_.Create_DC5_(p1, (Companion_D3_.Create_DC12_((_640_v5).F32())).Dtor_cf23(), !((_641_v6).Dtor_cf23()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_642_v7, _633_v0))
			_643_v8 = _643_v8
			var _644_v9 _dafny.Map
			_ = _644_v9
			_644_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-298), _640_v5)
			(globalState).F10 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_644_v9).Cardinality(), (_dafny.Zero).Minus(p3)), (p3).Plus(p2))
			var _645_v10 _dafny.CodePoint
			_ = _645_v10
			_645_v10 = _dafny.CodePoint('n')
			(globalState).F19 = _645_v10
		} else {
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _index144
			var _rhs147 bool = (_this).Fm21((_dafny.Zero).Minus(p3), globalState)
			_ = _rhs147
			var _rhs148 _dafny.Int = p1
			_ = _rhs148
			var _rhs149 bool = !((p3).Cmp(p0) > 0)
			_ = _rhs149
			var _lhs139 *GlobalState = globalState
			_ = _lhs139
			var _lhs140 _dafny.Array = _634_v1
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _lhs141
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			_lhs139.F22 = _rhs147
			(_lhs140).ArraySet1(_rhs148, (_lhs141).Int())
			_lhs142.F7 = _rhs149
			var _646_v11 *C0
			_ = _646_v11
			var _nw133 *C0 = New_C0_()
			_ = _nw133
			_nw133.Ctor__((_this).Fm21(p2, globalState))
			_646_v11 = _nw133
			var _647_v12 _dafny.MultiSet
			_ = _647_v12
			_647_v12 = _dafny.MultiSetOf((_646_v11).F32())
			var _648_v13 _dafny.Map
			_ = _648_v13
			_648_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_646_v11).F32())), !(!((_646_v11).F32())))
			(globalState).F10 = (func() _dafny.Int {
				if (_646_v11).F32() {
					return (func() _dafny.Int {
						if (_647_v12).Contains(false) {
							return (_647_v12).Multiplicity(false)
						}
						return (_648_v13).Cardinality()
					})()
				}
				return (_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int)
			})()
			var _649_v14 *C0
			_ = _649_v14
			var _nw134 *C0 = New_C0_()
			_ = _nw134
			_nw134.Ctor__(_635_v2)
			_649_v14 = _nw134
			var _650_v15 _dafny.Sequence
			_ = _650_v15
			_650_v15 = _dafny.UnicodeSeqOfUtf8Bytes("lroq")
			var _651_v16 _dafny.Map
			_ = _651_v16
			_651_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int), p3)
			var _652_v17 D2
			_ = _652_v17
			_652_v17 = Companion_D2_.Create_DC7_(_651_v16)
			var _653_v18 _dafny.CodePoint
			_ = _653_v18
			_653_v18 = _dafny.CodePoint('x')
			var _654_v19 D3
			_ = _654_v19
			_654_v19 = Companion_D3_.Create_DC11_(p3)
			var _655_v20 _dafny.Sequence
			_ = _655_v20
			_655_v20 = _dafny.SeqOf(_650_v15)
			var _656_v21 _dafny.Array
			_ = _656_v21
			var _nwElement0_38 _dafny.Sequence = _650_v15
			_ = _nwElement0_38
			var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(28))
			_ = _nw135
			(_nw135).ArraySet1(_nwElement0_38, 0)
			(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(663))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}(func(_657_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), 1)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23(_652_v17, p2, (_dafny.Zero).Minus((_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int)), globalState), _dafny.UnicodeSeqOfUtf8Bytes("dtnxh")), 2)
			(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(620))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}(func(_658_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), 3)
			(_nw135).ArraySet1(_650_v15, 4)
			(_nw135).ArraySet1(_650_v15, 5)
			(_nw135).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vhcdvwyj"), 6)
			(_nw135).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("chvp"), 7)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_650_v15, _650_v15), 8)
			(_nw135).ArraySet1(_650_v15, 9)
			(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(44))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_659_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			})), 10)
			(_nw135).ArraySet1(_650_v15, 11)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Update(_650_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(781), _dafny.IntOfUint32((_650_v15).Cardinality()))).Uint32(), _653_v18), 12)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_650_v15, _dafny.UnicodeSeqOfUtf8Bytes("xjxlm")), 13)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kxppprism"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_660_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_661_i4 _dafny.Int) _dafny.CodePoint {
					return _660_v18
				}
			})(_653_v18)))), 14)
			(_nw135).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qghd"), 15)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ajvtl"), _650_v15), 16)
			(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(137))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_662_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_663_i5 _dafny.Int) _dafny.CodePoint {
					return _662_v18
				}
			})(_653_v18))), 17)
			(_nw135).ArraySet1(_650_v15, 18)
			(_nw135).ArraySet1(Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_651_v16), _dafny.IntOfInt64(422), (_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int), globalState), 19)
			(_nw135).ArraySet1(_650_v15, 20)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(785))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_664_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_665_i6 _dafny.Int) _dafny.CodePoint {
					return _664_v18
				}
			})(_653_v18))), _650_v15), 21)
			(_nw135).ArraySet1(_650_v15, 22)
			(_nw135).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-926))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg44 _dafny.Int) interface{} {
					return coer44(arg44)
				}
			}((func(_666_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_667_i7 _dafny.Int) _dafny.CodePoint {
					return _666_v18
				}
			})(_653_v18))), 23)
			(_nw135).ArraySet1(_650_v15, 24)
			(_nw135).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ytkqvx"), 25)
			(_nw135).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_650_v15, (Companion_Default___.SafeIndex((_654_v19).Dtor_cf22(), _dafny.IntOfUint32((_650_v15).Cardinality()))).Uint32(), _653_v18), _650_v15), 26)
			(_nw135).ArraySet1((_655_v20).Select((Companion_Default___.SafeIndex((_634_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_655_v20).Cardinality()))).Uint32()).(_dafny.Sequence), 27)
			_656_v21 = _nw135
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_656_v21), 0))
			_ = _index145
			(_656_v21).ArraySet1(_650_v15, (_index145).Int())
			var _668_v22 _dafny.Array
			_ = _668_v22
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(28))
			_ = _nw136
			_668_v22 = _nw136
			var _669_v23 D4
			_ = _669_v23
			_669_v23 = Companion_D4_.Create_DC13_(_653_v18)
			var _670_v24 D4
			_ = _670_v24
			_670_v24 = Companion_D4_.Create_DC15_(_669_v23)
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_668_v22), 0))
			_ = _index146
			(_668_v22).ArraySet1(_670_v24, (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_656_v21), 0))
			_ = _index147
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _index148
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_668_v22), 0))
			_ = _index149
			var _rhs150 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("akub")
			_ = _rhs150
			var _rhs151 _dafny.Int = _dafny.IntOfUint32((_638_v3).Cardinality())
			_ = _rhs151
			var _rhs152 D4 = _670_v24
			_ = _rhs152
			var _rhs153 bool = (_646_v11).F32()
			_ = _rhs153
			var _lhs143 _dafny.Array = _656_v21
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(835), _dafny.ArrayLen((_656_v21), 0))
			_ = _lhs144
			var _lhs145 _dafny.Array = _634_v1
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_634_v1), 0))
			_ = _lhs146
			var _lhs147 _dafny.Array = _668_v22
			_ = _lhs147
			var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_668_v22), 0))
			_ = _lhs148
			var _lhs149 *GlobalState = globalState
			_ = _lhs149
			(_lhs143).ArraySet1(_rhs150, (_lhs144).Int())
			(_lhs145).ArraySet1(_rhs151, (_lhs146).Int())
			(_lhs147).ArraySet1(_rhs152, (_lhs148).Int())
			_lhs149.F7 = _rhs153
		}
		var _671_v25 _dafny.Array
		_ = _671_v25
		var _nw137 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw137
		_671_v25 = _nw137
		var _672_v26 _dafny.Sequence
		_ = _672_v26
		_672_v26 = _dafny.UnicodeSeqOfUtf8Bytes("cfgaepsi")
		var _rhs154 _dafny.Array = _671_v25
		_ = _rhs154
		var _rhs155 _dafny.Sequence = _672_v26
		_ = _rhs155
		var _rhs156 _dafny.Sequence = _633_v0
		_ = _rhs156
		var _rhs157 _dafny.Int = p3
		_ = _rhs157
		var _lhs150 *GlobalState = globalState
		_ = _lhs150
		_671_v25 = _rhs154
		_672_v26 = _rhs155
		_633_v0 = _rhs156
		_lhs150.F10 = _rhs157
		var _673_v27 _dafny.CodePoint
		_ = _673_v27
		_673_v27 = _dafny.CodePoint('m')
		r0 = _673_v27
		return r0
	}
}
func (_this *C1) M7(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) (_dafny.Array, bool) {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var r1 bool = false
		_ = r1
		var _674_v0 _dafny.Set
		_ = _674_v0
		_674_v0 = _dafny.SetOf(false)
		var _675_v1 bool
		_ = _675_v1
		_675_v1 = true
		var _676_v2 _dafny.Array
		_ = _676_v2
		var _nwElement0_39 bool = _675_v1
		_ = _nwElement0_39
		var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(6))
		_ = _nw138
		(_nw138).ArraySet1(_nwElement0_39, 0)
		(_nw138).ArraySet1(_675_v1, 1)
		(_nw138).ArraySet1(_675_v1, 2)
		(_nw138).ArraySet1(_675_v1, 3)
		(_nw138).ArraySet1(_675_v1, 4)
		(_nw138).ArraySet1(_675_v1, 5)
		_676_v2 = _nw138
		var _pat_let_tv11 = _675_v1
		_ = _pat_let_tv11
		var _pat_let_tv12 = _676_v2
		_ = _pat_let_tv12
		var _pat_let_tv13 = _676_v2
		_ = _pat_let_tv13
		var _source12 D0 = func(_pat_let8_0 D0) D0 {
			return func(_677_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let9_0 _dafny.Array) D0 {
					return func(_678_dt__update_hcf4_h0 _dafny.Array) D0 {
						return Companion_D0_.Create_DC1_((_677_dt__update__tmp_h0).Dtor_cf1(), (_677_dt__update__tmp_h0).Dtor_cf2(), (_677_dt__update__tmp_h0).Dtor_cf3(), _678_dt__update_hcf4_h0)
					}(_pat_let9_0)
				}((func() _dafny.Array {
					if _pat_let_tv11 {
						return _pat_let_tv12
					}
					return _pat_let_tv13
				})())
			}(_pat_let8_0)
		}(Companion_D0_.Create_DC1_((_674_v0).Cardinality(), _dafny.IntOfInt64(-556), p0, _676_v2))
		_ = _source12
		if _source12.Is_DC1() {
			var _679___mcc_h0 _dafny.Int = _source12.Get_().(D0_DC1).Cf1
			_ = _679___mcc_h0
			var _680___mcc_h1 _dafny.Int = _source12.Get_().(D0_DC1).Cf2
			_ = _680___mcc_h1
			var _681___mcc_h2 _dafny.Int = _source12.Get_().(D0_DC1).Cf3
			_ = _681___mcc_h2
			var _682___mcc_h3 _dafny.Array = _source12.Get_().(D0_DC1).Cf4
			_ = _682___mcc_h3
			var _683_cf4 _dafny.Array = _682___mcc_h3
			_ = _683_cf4
			var _684_cf3 _dafny.Int = _681___mcc_h2
			_ = _684_cf3
			var _685_cf2 _dafny.Int = _680___mcc_h1
			_ = _685_cf2
			var _686_cf1 _dafny.Int = _679___mcc_h0
			_ = _686_cf1
			var _687_v3 _dafny.Array
			_ = _687_v3
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_14
			var _nw139 _dafny.Array
			_ = _nw139
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw139 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Set = (func(_688_v0 _dafny.Set) func(_dafny.Int) _dafny.Set {
					return func(_689_i0 _dafny.Int) _dafny.Set {
						return (_688_v0).Union(_688_v0)
					}
				})(_674_v0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw139 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw139).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw139).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_687_v3 = _nw139
			_687_v3 = _687_v3
			var _690_v4 _dafny.Sequence
			_ = _690_v4
			_690_v4 = _dafny.UnicodeSeqOfUtf8Bytes("gteyicgu")
			var _691_v5 D0
			_ = _691_v5
			_691_v5 = Companion_D0_.Create_DC0_(_690_v4)
			var _692_v6 _dafny.Set
			_ = _692_v6
			_692_v6 = _dafny.SetOf(_691_v5)
			_685_cf2 = (_692_v6).Cardinality()
			var _693_v7 _dafny.MultiSet
			_ = _693_v7
			_693_v7 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_690_v4, _690_v4))
			_693_v7 = _dafny.MultiSetOf(_690_v4, _690_v4)
			var _694_v8 _dafny.Array
			_ = _694_v8
			var _nwElement0_40 _dafny.Int = _686_cf1
			_ = _nwElement0_40
			var _nw140 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(4))
			_ = _nw140
			(_nw140).ArraySet1(_nwElement0_40, 0)
			(_nw140).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(p0, _684_cf3, _dafny.IntOfUint32((_690_v4).Cardinality()), _dafny.IntOfInt64(21))).Cardinality())).Plus(_686_cf1)), 1)
			(_nw140).ArraySet1(_685_cf2, 2)
			(_nw140).ArraySet1((_dafny.IntOfInt64(500)).Plus(_686_cf1), 3)
			_694_v8 = _nw140
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_694_v8), 0))
			_ = _index150
			(_694_v8).ArraySet1((_dafny.Zero).Minus(p1), (_index150).Int())
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_694_v8), 0))
			_ = _index151
			(_694_v8).ArraySet1(_686_cf1, (_index151).Int())
		} else if _source12.Is_DC2() {
			var _695___mcc_h4 _dafny.Int = _source12.Get_().(D0_DC2).Cf5
			_ = _695___mcc_h4
			var _696___mcc_h5 bool = _source12.Get_().(D0_DC2).Cf6
			_ = _696___mcc_h5
			var _697_cf6 bool = _696___mcc_h5
			_ = _697_cf6
			var _698_cf5 _dafny.Int = _695___mcc_h4
			_ = _698_cf5
			var _699_v9 _dafny.Sequence
			_ = _699_v9
			_699_v9 = _dafny.UnicodeSeqOfUtf8Bytes("rpgfswiru")
			(globalState).F13 = !(_dafny.Companion_Sequence_.IsPrefixOf(_699_v9, _699_v9)) || ((_698_cf5).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm22(_697_cf6, globalState)).Cardinality())) == 0)
			var _700_v10 _dafny.Array
			_ = _700_v10
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_15
			var _nw141 _dafny.Array
			_ = _nw141
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw141 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Int = (func(_701_cf5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_702_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_702_i1, _701_cf5)
					}
				})(_698_cf5)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw141 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw141).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw141).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_700_v10 = _nw141
			var _703_v11 _dafny.Sequence
			_ = _703_v11
			_703_v11 = _dafny.SeqOf(_698_cf5)
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_700_v10), 0))
			_ = _index152
			(_700_v10).ArraySet1((_703_v11).Select((Companion_Default___.SafeIndex((_this).Fm2(globalState), _dafny.IntOfUint32((_703_v11).Cardinality()))).Uint32()).(_dafny.Int), (_index152).Int())
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(735), _dafny.ArrayLen((_700_v10), 0))
			_ = _index153
			(_700_v10).ArraySet1(p1, (_index153).Int())
			r0 = _700_v10
			_698_cf5 = _698_cf5
		} else if _source12.Is_DC0() {
			var _704___mcc_h6 _dafny.Sequence = _source12.Get_().(D0_DC0).Cf0
			_ = _704___mcc_h6
			var _705_cf0 _dafny.Sequence = _704___mcc_h6
			_ = _705_cf0
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_676_v2), 0))
			_ = _index154
			(_676_v2).ArraySet1(!(_675_v1), (_index154).Int())
			var _706_v12 D0
			_ = _706_v12
			_706_v12 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(527), _675_v1)
			var _pat_let_tv14 = _675_v1
			_ = _pat_let_tv14
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_676_v2), 0))
			_ = _index155
			(_676_v2).ArraySet1(((_675_v1) || ((func(_pat_let10_0 D0) D0 {
				return func(_707_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let11_0 bool) D0 {
						return func(_708_dt__update_hcf6_h0 bool) D0 {
							return Companion_D0_.Create_DC2_((_707_dt__update__tmp_h1).Dtor_cf5(), _708_dt__update_hcf6_h0)
						}(_pat_let11_0)
					}(_pat_let_tv14)
				}(_pat_let10_0)
			}(_706_v12)).Dtor_cf6())) || (false), (_index155).Int())
			var _709_v13 _dafny.Array
			_ = _709_v13
			var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
			_ = _nw142
			_709_v13 = _nw142
			r0 = _709_v13
			var _710_v14 D3
			_ = _710_v14
			_710_v14 = Companion_D3_.Create_DC11_(p1)
			(globalState).F10 = (_710_v14).Dtor_cf22()
			var _711_v15 _dafny.Map
			_ = _711_v15
			_711_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_710_v14, !(_675_v1))
			var _712_v16 _dafny.Sequence
			_ = _712_v16
			_712_v16 = _dafny.SeqOf(_711_v15)
			_712_v16 = _712_v16
		} else {
			var _713___mcc_h7 D0 = _source12.Get_().(D0_DC3).Cf7
			_ = _713___mcc_h7
			var _714_cf7 D0 = _713___mcc_h7
			_ = _714_cf7
			var _715_v17 _dafny.CodePoint
			_ = _715_v17
			_715_v17 = _dafny.CodePoint('a')
			var _716_v18 _dafny.Sequence
			_ = _716_v18
			_716_v18 = _dafny.SeqOf(_715_v17, _715_v17, _dafny.CodePoint('v'), _715_v17, _dafny.CodePoint('g'))
			_716_v18 = _dafny.Companion_Sequence_.Concatenate(_716_v18, _dafny.SeqOf(_715_v17, _715_v17, _715_v17))
			var _717_v19 _dafny.Sequence
			_ = _717_v19
			_717_v19 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg45 _dafny.Int) interface{} {
					return coer45(arg45)
				}
			}((func(_718_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_719_i2 _dafny.Int) _dafny.CodePoint {
					return _718_v17
				}
			})(_715_v17))))
			_716_v18 = (_717_v19).Select((Companion_Default___.SafeIndex(((func() _dafny.Set {
				if _675_v1 {
					return _674_v0
				}
				return _674_v0
			})()).Cardinality(), _dafny.IntOfUint32((_717_v19).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _720_v20 D0
			_ = _720_v20
			_720_v20 = Companion_D0_.Create_DC0_(_716_v18)
			var _pat_let_tv15 = _716_v18
			_ = _pat_let_tv15
			var _pat_let_tv16 = _716_v18
			_ = _pat_let_tv16
			var _source13 D0 = func(_pat_let12_0 D0) D0 {
				return func(_721_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let13_0 _dafny.Sequence) D0 {
						return func(_722_dt__update_hcf0_h0 _dafny.Sequence) D0 {
							return Companion_D0_.Create_DC0_(_722_dt__update_hcf0_h0)
						}(_pat_let13_0)
					}(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv15, _pat_let_tv16))
				}(_pat_let12_0)
			}(_720_v20)
			_ = _source13
			if _source13.Is_DC1() {
				var _723___mcc_h8 _dafny.Int = _source13.Get_().(D0_DC1).Cf1
				_ = _723___mcc_h8
				var _724___mcc_h9 _dafny.Int = _source13.Get_().(D0_DC1).Cf2
				_ = _724___mcc_h9
				var _725___mcc_h10 _dafny.Int = _source13.Get_().(D0_DC1).Cf3
				_ = _725___mcc_h10
				var _726___mcc_h11 _dafny.Array = _source13.Get_().(D0_DC1).Cf4
				_ = _726___mcc_h11
				var _727_cf4 _dafny.Array = _726___mcc_h11
				_ = _727_cf4
				var _728_cf3 _dafny.Int = _725___mcc_h10
				_ = _728_cf3
				var _729_cf2 _dafny.Int = _724___mcc_h9
				_ = _729_cf2
				var _730_cf1 _dafny.Int = _723___mcc_h8
				_ = _730_cf1
				var _731_v21 _dafny.Array
				_ = _731_v21
				var _nw143 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
				_ = _nw143
				_731_v21 = _nw143
				var _732_v22 _dafny.Set
				_ = _732_v22
				_732_v22 = _dafny.SetOf(_728_cf3, _729_cf2)
				var _rhs158 _dafny.Sequence = (func() _dafny.Sequence {
					if _675_v1 {
						return _dafny.UnicodeSeqOfUtf8Bytes("l")
					}
					return _716_v18
				})()
				_ = _rhs158
				var _rhs159 bool = _675_v1
				_ = _rhs159
				var _rhs160 _dafny.Array = _731_v21
				_ = _rhs160
				var _rhs161 bool = (((_732_v22).Cardinality()).Cmp(_728_cf3) > 0) == (_675_v1)
				_ = _rhs161
				var _lhs151 *GlobalState = globalState
				_ = _lhs151
				_716_v18 = _rhs158
				_675_v1 = _rhs159
				_731_v21 = _rhs160
				_lhs151.F7 = _rhs161
				var _733_v23 D3
				_ = _733_v23
				_733_v23 = Companion_D3_.Create_DC11_(p1)
				_733_v23 = _733_v23
				var _734_v24 D0
				_ = _734_v24
				_734_v24 = Companion_D0_.Create_DC2_(p1, _675_v1)
				var _735_v25 D0
				_ = _735_v25
				_735_v25 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC3_(_734_v24))
				var _736_v26 _dafny.Array
				_ = _736_v26
				var _nwElement0_41 D0 = _735_v25
				_ = _nwElement0_41
				var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(2))
				_ = _nw144
				(_nw144).ArraySet1(_nwElement0_41, 0)
				(_nw144).ArraySet1((func() D0 {
					if _675_v1 {
						return _735_v25
					}
					return _735_v25
				})(), 1)
				_736_v26 = _nw144
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_736_v26), 0))
				_ = _index156
				(_736_v26).ArraySet1((func() D0 {
					if _675_v1 {
						return _735_v25
					}
					return _735_v25
				})(), (_index156).Int())
				var _737_v27 _dafny.Map
				_ = _737_v27
				_737_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_729_cf2, _675_v1), _675_v1)
				var _738_v28 D0
				_ = _738_v28
				_738_v28 = Companion_D0_.Create_DC2_(p0, _675_v1)
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_736_v26), 0))
				_ = _index157
				var _rhs162 _dafny.Sequence = (_717_v19).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_728_cf3), _dafny.IntOfUint32((_717_v19).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs162
				var _rhs163 D0 = _735_v25
				_ = _rhs163
				var _rhs164 bool = (func() bool {
					if _675_v1 {
						return _675_v1
					}
					return (_738_v28).Dtor_cf6()
				})()
				_ = _rhs164
				var _rhs165 _dafny.Map = _737_v27
				_ = _rhs165
				var _lhs152 _dafny.Array = _736_v26
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_736_v26), 0))
				_ = _lhs153
				var _lhs154 *GlobalState = globalState
				_ = _lhs154
				_716_v18 = _rhs162
				(_lhs152).ArraySet1(_rhs163, (_lhs153).Int())
				_lhs154.F13 = _rhs164
				_737_v27 = _rhs165
				var _739_v29 _dafny.Array
				_ = _739_v29
				var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(29))
				_ = _nw145
				_739_v29 = _nw145
				var _740_v30 _dafny.Array
				_ = _740_v30
				_740_v30 = _739_v29
				var _741_v31 _dafny.Sequence
				_ = _741_v31
				_741_v31 = _dafny.SeqOf((_740_v30), _739_v29, _739_v29)
				var _rhs166 _dafny.Array = (_741_v31).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_741_v31).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs166
				var _rhs167 _dafny.Array = _676_v2
				_ = _rhs167
				var _rhs168 bool = _675_v1
				_ = _rhs168
				var _rhs169 _dafny.Int = _730_cf1
				_ = _rhs169
				r0 = _rhs166
				_676_v2 = _rhs167
				r1 = _rhs168
				_729_cf2 = _rhs169
			} else if _source13.Is_DC2() {
				var _742___mcc_h12 _dafny.Int = _source13.Get_().(D0_DC2).Cf5
				_ = _742___mcc_h12
				var _743___mcc_h13 bool = _source13.Get_().(D0_DC2).Cf6
				_ = _743___mcc_h13
				var _744_cf6 bool = _743___mcc_h13
				_ = _744_cf6
				var _745_cf5 _dafny.Int = _742___mcc_h12
				_ = _745_cf5
				(globalState).F13 = _744_cf6
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_676_v2), 0))
				_ = _index158
				(_676_v2).ArraySet1(!(_dafny.Companion_Sequence_.Equal(_716_v18, _716_v18)), (_index158).Int())
				var _746_v32 _dafny.Map
				_ = _746_v32
				_746_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v1, !(_675_v1))
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_676_v2), 0))
				_ = _index159
				var _rhs170 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
				_ = _rhs170
				var _rhs171 bool = _744_cf6
				_ = _rhs171
				var _rhs172 _dafny.Int = (func() _dafny.Int {
					if (Companion_Default___.Fm30(true, globalState)).Equals(_746_v32) {
						return Companion_Default___.SafeDivisionInt(p1, p0)
					}
					return (_dafny.Zero).Minus(_745_cf5)
				})()
				_ = _rhs172
				var _lhs155 *GlobalState = globalState
				_ = _lhs155
				var _lhs156 _dafny.Array = _676_v2
				_ = _lhs156
				var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_676_v2), 0))
				_ = _lhs157
				var _lhs158 *GlobalState = globalState
				_ = _lhs158
				_lhs155.F10 = _rhs170
				(_lhs156).ArraySet1(_rhs171, (_lhs157).Int())
				_lhs158.F10 = _rhs172
				var _747_v33 _dafny.Sequence
				_ = _747_v33
				_747_v33 = _dafny.SeqOf((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				var _748_v34 _dafny.Sequence
				_ = _748_v34
				_748_v34 = _dafny.SeqOf((_747_v33).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_747_v33).Cardinality()))).Uint32()).(bool), (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				_748_v34 = _747_v33
				var _749_v35 _dafny.Array
				_ = _749_v35
				var _nw146 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw146
				_749_v35 = _nw146
				var _750_v36 _dafny.Sequence
				_ = _750_v36
				_750_v36 = _dafny.SeqOf(_749_v35)
				r1 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_750_v36, _750_v36), _750_v36)
			} else if _source13.Is_DC0() {
				var _751___mcc_h14 _dafny.Sequence = _source13.Get_().(D0_DC0).Cf0
				_ = _751___mcc_h14
				var _752_cf0 _dafny.Sequence = _751___mcc_h14
				_ = _752_cf0
				var _753_v37 _dafny.MultiSet
				_ = _753_v37
				_753_v37 = _dafny.MultiSetOf(_675_v1)
				var _754_v38 _dafny.Map
				_ = _754_v38
				_754_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _675_v1))
				var _755_v39 _dafny.Sequence
				_ = _755_v39
				_755_v39 = _dafny.SeqOf(Companion_Default___.Fm0(p1, globalState), (_this).Fm2(globalState), (_753_v37).Cardinality(), p0, (_754_v38).Cardinality())
				var _756_v40 _dafny.CodePoint
				_ = _756_v40
				var _out3 _dafny.CodePoint
				_ = _out3
				_out3 = (_this).M6(Companion_Default___.Fm0(p0, globalState), (p0).Times(p1), (_755_v39).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_755_v39).Cardinality()))).Uint32()).(_dafny.Int), p0, globalState)
				_756_v40 = _out3
				(globalState).F22 = _675_v1
				var _757_v41 _dafny.Array
				_ = _757_v41
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_16
				var _nw147 _dafny.Array
				_ = _nw147
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw147 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_758_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_759_i3 _dafny.Int) _dafny.Int {
							return (_759_i3).Times(_758_p1)
						}
					})(p1)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw147 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw147).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw147).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_757_v41 = _nw147
				var _760_v42 _dafny.Map
				_ = _760_v42
				_760_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v1, _757_v41)
				_760_v42 = (_760_v42).Update(_675_v1, _757_v41)
				(globalState).F13 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg46 _dafny.Int) interface{} {
						return coer46(arg46)
					}
				}((func(_761_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_762_i4 _dafny.Int) _dafny.CodePoint {
						return _761_v17
					}
				})(_715_v17))), _716_v18), (_720_v20).Dtor_cf0())
			} else {
				var _763___mcc_h15 D0 = _source13.Get_().(D0_DC3).Cf7
				_ = _763___mcc_h15
				var _764_cf7 D0 = _763___mcc_h15
				_ = _764_cf7
				var _765_v43 _dafny.Array
				_ = _765_v43
				var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
				_ = _nw148
				_765_v43 = _nw148
				var _766_v44 _dafny.Sequence
				_ = _766_v44
				_766_v44 = _dafny.SeqOf(_765_v43)
				var _rhs173 _dafny.CodePoint = _dafny.CodePoint('p')
				_ = _rhs173
				var _rhs174 bool = _675_v1
				_ = _rhs174
				var _rhs175 _dafny.Array = (_766_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(89), _dafny.IntOfUint32((_766_v44).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs175
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				_715_v17 = _rhs173
				_lhs159.F22 = _rhs174
				_765_v43 = _rhs175
				_676_v2 = _676_v2
				(globalState).F10 = Companion_Default___.SafeDivisionInt(p0, (_dafny.Zero).Minus(p1))
				(globalState).F22 = _675_v1
			}
			var _767_v45 _dafny.Array
			_ = _767_v45
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_17
			var _nw149 _dafny.Array
			_ = _nw149
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw149 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) D4 = (func(_768_v1 bool) func(_dafny.Int) D4 {
					return func(_769_i5 _dafny.Int) D4 {
						return (func() D4 {
							if (Companion_D2_.Create_DC8_(_768_v1, _768_v1)).Dtor_cf19() {
								return Companion_D4_.Create_DC13_(_dafny.CodePoint('i'))
							}
							return Companion_D4_.Create_DC13_(_dafny.CodePoint('o'))
						})()
					}
				})(_675_v1)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw149 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw149).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw149).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_767_v45 = _nw149
			var _770_v46 D4
			_ = _770_v46
			_770_v46 = Companion_D4_.Create_DC13_(_715_v17)
			var _pat_let_tv17 = _715_v17
			_ = _pat_let_tv17
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_767_v45), 0))
			_ = _index160
			(_767_v45).ArraySet1(func(_pat_let14_0 D4) D4 {
				return func(_771_dt__update__tmp_h3 D4) D4 {
					return func(_pat_let15_0 _dafny.CodePoint) D4 {
						return func(_772_dt__update_hcf24_h0 _dafny.CodePoint) D4 {
							return Companion_D4_.Create_DC13_(_772_dt__update_hcf24_h0)
						}(_pat_let15_0)
					}(_pat_let_tv17)
				}(_pat_let14_0)
			}(_770_v46), (_index160).Int())
			var _pat_let_tv18 = _715_v17
			_ = _pat_let_tv18
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_767_v45), 0))
			_ = _index161
			(_767_v45).ArraySet1(func(_pat_let16_0 D4) D4 {
				return func(_773_dt__update__tmp_h4 D4) D4 {
					return func(_pat_let17_0 _dafny.CodePoint) D4 {
						return func(_774_dt__update_hcf24_h1 _dafny.CodePoint) D4 {
							return Companion_D4_.Create_DC13_(_774_dt__update_hcf24_h1)
						}(_pat_let17_0)
					}(_pat_let_tv18)
				}(_pat_let16_0)
			}(_770_v46), (_index161).Int())
		}
		var _775_v47 _dafny.Map
		_ = _775_v47
		_775_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm21(p1, globalState)), _676_v2)
		_775_v47 = (_775_v47).Update((func() bool {
			if (_this).Fm21(p0, globalState) {
				return _675_v1
			}
			return _675_v1
		})(), _676_v2)
		var _776_v48 _dafny.Sequence
		_ = _776_v48
		_776_v48 = _dafny.SeqOf(_675_v1)
		var _hi2 _dafny.Int = p1
		_ = _hi2
		for _777_i6 := _dafny.IntOfUint32((_776_v48).Cardinality()); _777_i6.Cmp(_hi2) < 0; _777_i6 = _777_i6.Plus(_dafny.One) {
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))
			_ = _index162
			(_676_v2).ArraySet1(true, (_index162).Int())
			var _778_v49 _dafny.Map
			_ = _778_v49
			_778_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_675_v1, p0)
			var _779_v50 _dafny.Sequence
			_ = _779_v50
			_779_v50 = _dafny.SeqOf(_777_i6)
			var _780_v51 _dafny.Sequence
			_ = _780_v51
			_780_v51 = _dafny.SeqOf((func() _dafny.Int {
				if (_778_v49).Contains(!(_675_v1)) {
					return (_778_v49).Get(!(_675_v1)).(_dafny.Int)
				}
				return p1
			})(), _777_i6, (_779_v50).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_779_v50).Cardinality()))).Uint32()).(_dafny.Int), _777_i6, _777_i6)
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))
			_ = _index163
			(_676_v2).ArraySet1(((_780_v51).Select((Companion_Default___.SafeIndex(_777_i6, _dafny.IntOfUint32((_780_v51).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_777_i6) > 0, (_index163).Int())
			(globalState).F22 = !((_776_v48).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_776_v48).Cardinality()))).Uint32()).(bool))
			var _781_v52 D3
			_ = _781_v52
			_781_v52 = Companion_D3_.Create_DC11_(p1)
			var _source14 D3 = _781_v52
			_ = _source14
			if _source14.Is_DC11() {
				var _782___mcc_h16 _dafny.Int = _source14.Get_().(D3_DC11).Cf22
				_ = _782___mcc_h16
				var _783_cf22 _dafny.Int = _782___mcc_h16
				_ = _783_cf22
				var _784_v53 _dafny.Sequence
				_ = _784_v53
				_784_v53 = _dafny.UnicodeSeqOfUtf8Bytes("lkekff")
				_784_v53 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_784_v53, _784_v53), _784_v53)
				var _785_v54 _dafny.Array
				_ = _785_v54
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
				_ = _nw150
				_785_v54 = _nw150
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_785_v54), 0))
				_ = _index164
				(_785_v54).ArraySet1CodePoint(_dafny.CodePoint('e'), (_index164).Int())
				var _786_v55 _dafny.Set
				_ = _786_v55
				_786_v55 = _dafny.SetOf(_783_cf22)
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(428), _dafny.ArrayLen((_785_v54), 0))
				_ = _index165
				(_785_v54).ArraySet1CodePoint((_784_v53).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_786_v55).Intersection(_786_v55)).Cardinality()), _dafny.IntOfUint32((_784_v53).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index165).Int())
				var _787_v56 _dafny.Map
				_ = _787_v56
				_787_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), _675_v1), (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				var _788_v57 D1
				_ = _788_v57
				_788_v57 = Companion_D1_.Create_DC4_(_787_v56)
				_788_v57 = _788_v57
				var _789_v58 _dafny.MultiSet
				_ = _789_v58
				_789_v58 = _dafny.MultiSetOf((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), _675_v1)
				var _790_v59 _dafny.CodePoint
				_ = _790_v59
				var _out4 _dafny.CodePoint
				_ = _out4
				_out4 = (_this).M6(_777_i6, _783_cf22, _777_i6, Companion_Default___.SafeModuloInt(_783_cf22, (func() _dafny.Int {
					if (_789_v58).Contains(_675_v1) {
						return (_789_v58).Multiplicity(_675_v1)
					}
					return (_dafny.Zero).Minus(_dafny.IntOfInt64(-473))
				})()), globalState)
				_790_v59 = _out4
			} else if _source14.Is_DC12() {
				var _791___mcc_h17 bool = _source14.Get_().(D3_DC12).Cf23
				_ = _791___mcc_h17
				var _792_cf23 bool = _791___mcc_h17
				_ = _792_cf23
				var _793_v60 _dafny.Sequence
				_ = _793_v60
				_793_v60 = _dafny.UnicodeSeqOfUtf8Bytes("bxjrgwaol")
				var _794_v61 _dafny.Array
				_ = _794_v61
				var _nwElement0_42 bool = _792_cf23
				_ = _nwElement0_42
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(6))
				_ = _nw151
				(_nw151).ArraySet1(_nwElement0_42, 0)
				(_nw151).ArraySet1(_675_v1, 1)
				(_nw151).ArraySet1(_675_v1, 2)
				(_nw151).ArraySet1(_792_cf23, 3)
				(_nw151).ArraySet1((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), 4)
				(_nw151).ArraySet1(_792_cf23, 5)
				_794_v61 = _nw151
				var _795_v62 _dafny.MultiSet
				_ = _795_v62
				_795_v62 = _dafny.MultiSetOf(false)
				var _796_v63 _dafny.Map
				_ = _796_v63
				_796_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_781_v52, (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				var _797_v64 _dafny.Sequence
				_ = _797_v64
				_797_v64 = _dafny.SeqOf(_796_v63)
				var _798_v65 _dafny.Map
				_ = _798_v65
				_798_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_792_cf23, (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				var _799_v66 D4
				_ = _799_v66
				_799_v66 = Companion_D4_.Create_DC14_(_675_v1)
				var _800_v67 D4
				_ = _800_v67
				_800_v67 = Companion_D4_.Create_DC15_(_799_v66)
				var _801_v68 D4
				_ = _801_v68
				_801_v68 = Companion_D4_.Create_DC15_(_800_v67)
				var _802_v69 D7
				_ = _802_v69
				_802_v69 = Companion_D7_.Create_DC20_(_dafny.SeqOf(_dafny.IntOfInt64(590)), _794_v61, _801_v68, (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), p1)
				var _803_v70 _dafny.Map
				_ = _803_v70
				_803_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(414), (_802_v69).Dtor_cf34())
				var _804_v71 _dafny.Array
				_ = _804_v71
				var _nwElement0_43 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((Companion_D7_.Create_DC18_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_805_v48 _dafny.Sequence) func(_dafny.Int) _dafny.Map {
					return func(_806_i7 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _dafny.IntOfUint32((_805_v48).Cardinality()))
					}
				})(_776_v48))))).Dtor_cf29()).Cardinality()), _dafny.IntOfUint32((_793_v60).Cardinality()))
				_ = _nwElement0_43
				var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(25))
				_ = _nw152
				(_nw152).ArraySet1(_nwElement0_43, 0)
				(_nw152).ArraySet1(p1, 1)
				(_nw152).ArraySet1(_dafny.IntOfInt64(698), 2)
				(_nw152).ArraySet1(p0, 3)
				(_nw152).ArraySet1(p0, 4)
				(_nw152).ArraySet1(p1, 5)
				(_nw152).ArraySet1(_dafny.IntOfInt64(249), 6)
				(_nw152).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_776_v48).Cardinality()))), 7)
				(_nw152).ArraySet1(_dafny.IntOfInt64(504), 8)
				(_nw152).ArraySet1(p0, 9)
				(_nw152).ArraySet1(p0, 10)
				(_nw152).ArraySet1(((_dafny.MultiSetOf(_794_v61)).Cardinality()).Minus(_dafny.IntOfInt64(-328)), 11)
				(_nw152).ArraySet1((_795_v62).Cardinality(), 12)
				(_nw152).ArraySet1(Companion_Default___.SafeDivisionInt(_777_i6, (_dafny.Zero).Minus(_777_i6)), 13)
				(_nw152).ArraySet1(((_797_v64).Select((Companion_Default___.SafeIndex((_795_v62).Cardinality(), _dafny.IntOfUint32((_797_v64).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), 14)
				(_nw152).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_793_v60, _dafny.UnicodeSeqOfUtf8Bytes("gicvf"))).Cardinality()), 15)
				(_nw152).ArraySet1(p0, 16)
				(_nw152).ArraySet1(_dafny.IntOfInt64(264), 17)
				(_nw152).ArraySet1(p0, 18)
				(_nw152).ArraySet1(_dafny.IntOfInt64(-87), 19)
				(_nw152).ArraySet1((func() _dafny.Int {
					if _675_v1 {
						return _777_i6
					}
					return _777_i6
				})(), 20)
				(_nw152).ArraySet1((_777_i6).Plus(p0), 21)
				(_nw152).ArraySet1(_dafny.IntOfInt64(487), 22)
				(_nw152).ArraySet1(Companion_Default___.SafeModuloInt((_798_v65).Cardinality(), (_803_v70).Cardinality()), 23)
				(_nw152).ArraySet1(_777_i6, 24)
				_804_v71 = _nw152
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_804_v71), 0))
				_ = _index166
				(_804_v71).ArraySet1(_dafny.IntOfUint32((_793_v60).Cardinality()), (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_804_v71), 0))
				_ = _index167
				var _rhs176 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_776_v48, _776_v48)
				_ = _rhs176
				var _rhs177 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(912))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_807_i6 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_808_i8 _dafny.Int) _dafny.Int {
						return _807_i6
					}
				})(_777_i6))), _779_v50)
				_ = _rhs177
				var _rhs178 _dafny.Int = p0
				_ = _rhs178
				var _lhs160 _dafny.Array = _804_v71
				_ = _lhs160
				var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_804_v71), 0))
				_ = _lhs161
				_776_v48 = _rhs176
				_780_v51 = _rhs177
				(_lhs160).ArraySet1(_rhs178, (_lhs161).Int())
				_798_v65 = (_798_v65).Update((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool), !(_792_cf23))
				(globalState).F7 = (_776_v48).Select((Companion_Default___.SafeIndex(_777_i6, _dafny.IntOfUint32((_776_v48).Cardinality()))).Uint32()).(bool)
				_798_v65 = (func() _dafny.Map {
					if _792_cf23 {
						return _798_v65
					}
					return (_798_v65).Merge(_798_v65)
				})()
			} else {
				var _809___mcc_h18 _dafny.MultiSet = _source14.Get_().(D3_DC10).Cf21
				_ = _809___mcc_h18
				var _810_cf21 _dafny.MultiSet = _809___mcc_h18
				_ = _810_cf21
				(globalState).F7 = (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool)
				(globalState).F13 = (_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool)
				var _811_v72 _dafny.Map
				_ = _811_v72
				_811_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_777_i6, _776_v48)
				var _812_v73 *C0
				_ = _812_v73
				var _nw153 *C0 = New_C0_()
				_ = _nw153
				_nw153.Ctor__((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool))
				_812_v73 = _nw153
				var _813_v74 _dafny.Sequence
				_ = _813_v74
				_813_v74 = _dafny.SeqOf(_812_v73)
				var _814_v75 _dafny.Sequence
				_ = _814_v75
				_814_v75 = _dafny.SeqOf(_776_v48, (func() _dafny.Sequence {
					if (_811_v72).Contains(_dafny.IntOfUint32((_813_v74).Cardinality())) {
						return (_811_v72).Get(_dafny.IntOfUint32((_813_v74).Cardinality())).(_dafny.Sequence)
					}
					return _776_v48
				})())
				var _815_v76 D1
				_ = _815_v76
				_815_v76 = Companion_D1_.Create_DC6_(_777_i6, _814_v75, _777_i6, (_812_v73).F32())
				var _816_v77 _dafny.Set
				_ = _816_v77
				_816_v77 = _dafny.SetOf(_815_v76)
				var _817_v78 _dafny.MultiSet
				_ = _817_v78
				_817_v78 = _dafny.MultiSetOf(_816_v77)
				var _818_v79 D3
				_ = _818_v79
				_818_v79 = Companion_D3_.Create_DC12_((_812_v73).F32())
				var _819_v80 _dafny.Map
				_ = _819_v80
				_819_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_818_v79, _817_v78)
				(globalState).F22 = ((_dafny.MultiSetOf(_816_v77)).Difference((func() _dafny.MultiSet {
					if (_819_v80).Contains(_818_v79) {
						return (_819_v80).Get(_818_v79).(_dafny.MultiSet)
					}
					return _817_v78
				})())).IsSubsetOf(_817_v78)
				var _820_v81 _dafny.Sequence
				_ = _820_v81
				_820_v81 = _dafny.UnicodeSeqOfUtf8Bytes("epv")
				var _821_v82 *C0
				_ = _821_v82
				var _nw154 *C0 = New_C0_()
				_ = _nw154
				_nw154.Ctor__((_dafny.IntOfUint32((_820_v81).Cardinality())).Cmp(p0) > 0)
				_821_v82 = _nw154
			}
			var _822_v83 *C0
			_ = _822_v83
			var _nw155 *C0 = New_C0_()
			_ = _nw155
			_nw155.Ctor__(!(_675_v1) || ((_676_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_676_v2), 0))).Int()).(bool)))
			_822_v83 = _nw155
		}
		var _823_v84 D4
		_ = _823_v84
		_823_v84 = Companion_D4_.Create_DC15_(Companion_D4_.Create_DC14_(_675_v1))
		var _824_v85 _dafny.Map
		_ = _824_v85
		_824_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_674_v0).IsProperSubsetOf(_674_v0), _823_v84)
		_824_v85 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _823_v84)).Merge(_824_v85)
		var _825_v86 _dafny.Sequence
		_ = _825_v86
		_825_v86 = _dafny.SeqOf(_676_v2, _676_v2, _676_v2, (func() _dafny.Array {
			if true {
				return _676_v2
			}
			return _676_v2
		})(), _676_v2)
		_676_v2 = (_825_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_825_v86).Cardinality()))).Uint32()).(_dafny.Array)
		var _826_v88 _dafny.Sequence
		_ = _826_v88
		_826_v88 = _dafny.SeqOf(p1)
		var _827_v89 _dafny.Map
		_ = _827_v89
		_827_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _675_v1)
		(globalState).F13 = (func() bool {
			if (func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter22 := _dafny.Iterate((_826_v88).Elements()); ; {
					_compr_21, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _828_v87 _dafny.Int
					_828_v87 = interface{}(_compr_21).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_826_v88, _828_v87) {
						_coll21.Add((_828_v87).Minus(p0), _675_v1)
					}
				}
				return _coll21.ToMap()
			}()).Equals(_827_v89) {
				return _675_v1
			}
			return _675_v1
		})()
		var _len0_18 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_18
		var _nw156 _dafny.Array
		_ = _nw156
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw156 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.Int = (func(_829_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_830_i9 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_830_i9, _829_p1)
				}
			})(p1)
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw156 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw156).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw156).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		r0 = _nw156
		r1 = _675_v1
		return r0, r1
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F33 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F33 = _dafny.EmptySeq
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

func (_this *C2) Ctor__(f33 _dafny.Sequence) {
	{
		(_this).F33 = f33
	}
}
func (_this *C2) Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter23 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.MultiSetOf(true, !(true)))).Elements()); ; {
				_compr_22, _ok23 := _iter23()
				if !_ok23 {
					break
				}
				var _831_v0 _dafny.MultiSet
				_831_v0 = interface{}(_compr_22).(_dafny.MultiSet)
				if (_dafny.MultiSetOf(_dafny.MultiSetOf(true, !(true)))).Contains(_831_v0) {
					_coll22.Add(_831_v0, false)
				}
			}
			return _coll22.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(true)), !(false)))
	}
}
func (_this *C2) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(999)).Minus((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-668)), _dafny.IntOfInt64(-602))).Cardinality())).Plus((func() _dafny.Int {
			if true {
				return (_dafny.SetOf(_dafny.IntOfInt64(-570), _dafny.IntOfUint32((_this.F33).Cardinality()))).Cardinality()
			}
			return (_dafny.SetOf(!(false), true, false)).Cardinality()
		})())
	}
}
func (_this *C2) Fm31(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !((func() bool {
			if !(true) {
				return true
			}
			return false
		})()) || (!_dafny.Companion_Sequence_.Contains(_this.F33, _dafny.CodePoint('n')))
	}
}
func (_this *C2) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _832_v0 _dafny.Array
		_ = _832_v0
		var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw157
		_832_v0 = _nw157
		var _833_v1 _dafny.MultiSet
		_ = _833_v1
		_833_v1 = _dafny.MultiSetOf(p1, p1)
		var _834_v2 _dafny.Array
		_ = _834_v2
		var _nwElement0_44 bool = false
		_ = _nwElement0_44
		var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(8))
		_ = _nw158
		(_nw158).ArraySet1(_nwElement0_44, 0)
		(_nw158).ArraySet1(p3, 1)
		(_nw158).ArraySet1(true, 2)
		(_nw158).ArraySet1(p1, 3)
		(_nw158).ArraySet1(p3, 4)
		(_nw158).ArraySet1(p3, 5)
		(_nw158).ArraySet1(p3, 6)
		(_nw158).ArraySet1(p1, 7)
		_834_v2 = _nw158
		var _835_v3 _dafny.MultiSet
		_ = _835_v3
		_835_v3 = _dafny.MultiSetOf(_834_v2)
		var _836_v4 _dafny.MultiSet
		_ = _836_v4
		_836_v4 = _dafny.MultiSetOf((_833_v1).Cardinality(), _dafny.IntOfUint32((_this.F33).Cardinality()), (_835_v3).Cardinality())
		var _837_v5 *C0
		_ = _837_v5
		var _nw159 *C0 = New_C0_()
		_ = _nw159
		_nw159.Ctor__(p1)
		_837_v5 = _nw159
		var _838_v6 _dafny.Set
		_ = _838_v6
		_838_v6 = _dafny.SetOf(_837_v5, _837_v5)
		var _839_v7 _dafny.Map
		_ = _839_v7
		_839_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_v4, (_838_v6).Cardinality())
		var _840_v8 _dafny.Set
		_ = _840_v8
		_840_v8 = _dafny.SetOf(p1)
		var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
		_ = _index168
		(_832_v0).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(855), (func() _dafny.Int {
			if (_839_v7).Contains(_836_v4) {
				return (_839_v7).Get(_836_v4).(_dafny.Int)
			}
			return (_840_v8).Cardinality()
		})()), (_index168).Int())
		var _841_v9 D3
		_ = _841_v9
		_841_v9 = Companion_D3_.Create_DC10_(_833_v1)
		var _842_v10 _dafny.Set
		_ = _842_v10
		_842_v10 = _dafny.SetOf(_841_v9, Companion_D3_.Create_DC10_(_833_v1), _841_v9, _841_v9)
		var _843_v11 _dafny.Map
		_ = _843_v11
		_843_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm31(_836_v4, globalState), _842_v10)
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
		_ = _index169
		(_832_v0).ArraySet1(((_843_v11).Update((_this).Fm31(_836_v4, globalState), (_842_v10).Intersection(_842_v10))).Cardinality(), (_index169).Int())
		var _844_v12 _dafny.Array
		_ = _844_v12
		var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw160
		_844_v12 = _nw160
		_844_v12 = _844_v12
		var _845_v13 _dafny.Array
		_ = _845_v13
		var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
		_ = _nw161
		_845_v13 = _nw161
		var _846_v14 _dafny.CodePoint
		_ = _846_v14
		_846_v14 = _dafny.CodePoint('h')
		var _847_v15 _dafny.Sequence
		_ = _847_v15
		_847_v15 = _dafny.SeqOf(Companion_Default___.Fm0(p2, globalState))
		var _848_v16 _dafny.MultiSet
		_ = _848_v16
		_848_v16 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_847_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.IntOfUint32((_847_v15).Cardinality()))).Uint32(), p2), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_847_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.IntOfUint32((_847_v15).Cardinality()))).Uint32(), p2)).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_837_v5).F32(), (_dafny.Zero).Minus(p2))).Cardinality()), (Companion_Default___.SafeIndex((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_847_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.IntOfUint32((_847_v15).Cardinality()))).Uint32(), p2), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_847_v15, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.IntOfUint32((_847_v15).Cardinality()))).Uint32(), p2)).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_837_v5).F32(), (_dafny.Zero).Minus(p2))).Cardinality())).Cardinality()))).Uint32(), p0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(261))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_849_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-686)
		})))
		var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_845_v13), 0))
		_ = _index170
		(_845_v13).ArraySet1((Companion_Default___.Fm32(_846_v14, (_837_v5).F32(), _846_v14, globalState)).Difference(_848_v16), (_index170).Int())
		var _850_v17 _dafny.Sequence
		_ = _850_v17
		_850_v17 = _dafny.SeqOf(p1)
		var _851_v18 _dafny.Map
		_ = _851_v18
		_851_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_850_v17).Cardinality()))
		var _852_v19 _dafny.Sequence
		_ = _852_v19
		_852_v19 = _dafny.SeqOf(_851_v18)
		var _853_v20 _dafny.Map
		_ = _853_v20
		_853_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_852_v19, _dafny.MultiSetOf(_847_v15))
		var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(710), _dafny.ArrayLen((_845_v13), 0))
		_ = _index171
		(_845_v13).ArraySet1((func() _dafny.MultiSet {
			if (_853_v20).Contains(_852_v19) {
				return (_853_v20).Get(_852_v19).(_dafny.MultiSet)
			}
			return _848_v16
		})(), (_index171).Int())
		if p1 {
			(_this).F33 = _this.F33
			if !(p1) {
				var _854_v21 _dafny.Map
				_ = _854_v21
				_854_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_837_v5).F32()), (_837_v5).F32())
				_854_v21 = (_854_v21).Merge(_854_v21)
				var _rhs179 _dafny.Array = _832_v0
				_ = _rhs179
				var _rhs180 _dafny.Int = Companion_Default___.Fm0(p0, globalState)
				_ = _rhs180
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				_832_v0 = _rhs179
				_lhs162.F10 = _rhs180
				var _855_v22 *C0
				_ = _855_v22
				var _nw162 *C0 = New_C0_()
				_ = _nw162
				_nw162.Ctor__((p1) || (p3))
				_855_v22 = _nw162
				(globalState).F10 = (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), p2)).Times(((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)).Times(p2)))
				(_this).F33 = _this.F33
			} else {
				(globalState).F10 = p0
				(globalState).F10 = Companion_Default___.SafeDivisionInt((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_850_v17).Cardinality()))
				(globalState).F10 = Companion_Default___.Fm0((_this).Fm2(globalState), globalState)
				(globalState).F7 = (func() bool {
					if (_this).Fm31(_836_v4, globalState) {
						return true
					}
					return (p0).Cmp(_dafny.IntOfUint32((_this.F33).Cardinality())) != 0
				})()
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
				_ = _index172
				(_832_v0).ArraySet1((_this).Fm2(globalState), (_index172).Int())
			}
			var _856_v23 D2
			_ = _856_v23
			_856_v23 = Companion_D2_.Create_DC9_(Companion_D2_.Create_DC9_(Companion_D2_.Create_DC7_(_851_v18)))
			var _857_v24 _dafny.Map
			_ = _857_v24
			_857_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC9_(_856_v23), (_837_v5).F32())
			var _858_v25 D2
			_ = _858_v25
			_858_v25 = Companion_D2_.Create_DC9_(_856_v23)
			_857_v24 = (_857_v24).Update(_858_v25, (_this).Fm31(_dafny.MultiSetFromSeq(_847_v15), globalState))
			var _859_v26 _dafny.Set
			_ = _859_v26
			_859_v26 = _dafny.SetOf((_840_v8).Cardinality())
			var _860_v27 _dafny.Map
			_ = _860_v27
			_860_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-76), (_859_v26).Cardinality()))
			_860_v27 = (_860_v27).Update(p1, (_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int))
			var _861_v28 *C1
			_ = _861_v28
			var _nw163 *C1 = New_C1_()
			_ = _nw163
			_nw163.Ctor__()
			_861_v28 = _nw163
		} else {
			var _862_v29 _dafny.Map
			_ = _862_v29
			_862_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p3) || (false), p2)
			(globalState).F10 = (func() _dafny.Int {
				if (_862_v29).Contains((_837_v5).F32()) {
					return (_862_v29).Get((_837_v5).F32()).(_dafny.Int)
				}
				return (_this).Fm2(globalState)
			})()
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_834_v2), 0))
			_ = _index173
			(_834_v2).ArraySet1(true, (_index173).Int())
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_834_v2), 0))
			_ = _index174
			var _rhs181 bool = !(_851_v18).Contains((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int))
			_ = _rhs181
			var _rhs182 _dafny.Int = (func() _dafny.Int {
				if p3 {
					return (func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(416), _dafny.IntOfInt64(822))); ; {
							_compr_23, _ok24 := _iter24()
							if !_ok24 {
								break
							}
							var _863_v30 _dafny.Int
							_863_v30 = interface{}(_compr_23).(_dafny.Int)
							if ((_dafny.IntOfInt64(416)).Cmp(_863_v30) <= 0) && ((_863_v30).Cmp(_dafny.IntOfInt64(822)) < 0) {
								_coll23.Add(Companion_Default___.SafeDivisionInt(_863_v30, p0), p3)
							}
						}
						return _coll23.ToMap()
					}()).Cardinality()
				}
				return Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_847_v15).Cardinality()))
			})()
			_ = _rhs182
			var _lhs163 _dafny.Array = _834_v2
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_834_v2), 0))
			_ = _lhs164
			var _lhs165 *GlobalState = globalState
			_ = _lhs165
			(_lhs163).ArraySet1(_rhs181, (_lhs164).Int())
			_lhs165.F10 = _rhs182
			(globalState).F13 = p3
			var _864_v31 D8
			_ = _864_v31
			_864_v31 = Companion_D8_.Create_DC22_((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), (_834_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_834_v2), 0))).Int()).(bool), _846_v14)
			var _865_v32 _dafny.Map
			_ = _865_v32
			_865_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_864_v31).Dtor_cf37()), (_837_v5).F32())
			var _source15 _dafny.Map = _865_v32
			_ = _source15
			var _866___mcc_h0 _dafny.Map = _source15
			_ = _866___mcc_h0
			var _867_cf27 _dafny.Map = _866___mcc_h0
			_ = _867_cf27
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
			_ = _index175
			var _rhs183 bool = false
			_ = _rhs183
			var _rhs184 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_847_v15, _847_v15), _847_v15)
			_ = _rhs184
			var _rhs185 _dafny.Int = p2
			_ = _rhs185
			var _lhs166 *GlobalState = globalState
			_ = _lhs166
			var _lhs167 *GlobalState = globalState
			_ = _lhs167
			var _lhs168 _dafny.Array = _832_v0
			_ = _lhs168
			var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
			_ = _lhs169
			_lhs166.F13 = _rhs183
			_lhs167.F13 = _rhs184
			(_lhs168).ArraySet1(_rhs185, (_lhs169).Int())
			(globalState).F7 = true
			var _868_v33 _dafny.Array
			_ = _868_v33
			var _nwElement0_45 bool = (_dafny.IntOfInt64(675)).Cmp((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)) >= 0
			_ = _nwElement0_45
			var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(9))
			_ = _nw164
			(_nw164).ArraySet1(_nwElement0_45, 0)
			(_nw164).ArraySet1(!((_836_v4).IsDisjointFrom(_836_v4)), 1)
			(_nw164).ArraySet1((_this).Fm31(_836_v4, globalState), 2)
			(_nw164).ArraySet1((_837_v5).F32(), 3)
			(_nw164).ArraySet1((_837_v5).F32(), 4)
			(_nw164).ArraySet1((_837_v5).F32(), 5)
			(_nw164).ArraySet1(false, 6)
			(_nw164).ArraySet1(((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)).Cmp(p0) > 0, 7)
			(_nw164).ArraySet1((_837_v5).F32(), 8)
			_868_v33 = _nw164
			var _rhs186 _dafny.Array = _868_v33
			_ = _rhs186
			var _rhs187 _dafny.Int = p2
			_ = _rhs187
			var _rhs188 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_this.F33, _this.F33)
			_ = _rhs188
			var _rhs189 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm0(p2, globalState)).Plus(Companion_Default___.Fm0((_dafny.Zero).Minus(p2), globalState)))
			_ = _rhs189
			var _lhs170 *GlobalState = globalState
			_ = _lhs170
			var _lhs171 *C2 = _this
			_ = _lhs171
			var _lhs172 *GlobalState = globalState
			_ = _lhs172
			_868_v33 = _rhs186
			_lhs170.F10 = _rhs187
			_lhs171.F33 = _rhs188
			_lhs172.F10 = _rhs189
			(globalState).F22 = !(!((_837_v5).F32()) || ((_this).Fm31(_836_v4, globalState)))
			var _869_v34 _dafny.Array
			_ = _869_v34
			var _nw165 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw165
			_869_v34 = _nw165
			var _870_v35 D0
			_ = _870_v35
			_870_v35 = Companion_D0_.Create_DC1_(p0, p2, (_dafny.Zero).Minus(p2), _869_v34)
			var _871_v36 _dafny.Map
			_ = _871_v36
			_871_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_850_v17).Cardinality()), (_834_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(168), _dafny.ArrayLen((_834_v2), 0))).Int()).(bool))
			_870_v35 = Companion_D0_.Create_DC1_((_871_v36).Cardinality(), (_dafny.IntOfUint32((_this.F33).Cardinality())).Plus((Companion_Default___.Fm33((_837_v5).F32(), (_837_v5).F32(), p3, (_837_v5).F32(), globalState)).Cardinality()), (func() _dafny.Int {
				if (_851_v18).Contains((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)) {
					return (_851_v18).Get((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return (_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int)
			})(), _869_v34)
		}
		if p1 {
			if false {
				(_this).F33 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_this.F33, _dafny.Companion_Sequence_.Update(_this.F33, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32(), _dafny.CodePoint('u'))), _this.F33)
				_850_v17 = _dafny.Companion_Sequence_.Concatenate(_850_v17, _dafny.Companion_Sequence_.Concatenate(_850_v17, _850_v17))
				(globalState).F13 = (func() bool {
					if true {
						return (_837_v5).F32()
					}
					return (_837_v5).F32()
				})()
				(globalState).F13 = p3
				(globalState).F10 = (_dafny.Zero).Minus(p0)
			} else {
				var _872_v37 _dafny.Sequence
				_ = _872_v37
				_872_v37 = _dafny.SeqOf(_832_v0)
				_872_v37 = _872_v37
				var _873_v38 *C1
				_ = _873_v38
				var _nw166 *C1 = New_C1_()
				_ = _nw166
				_nw166.Ctor__()
				_873_v38 = _nw166
				(globalState).F10 = Companion_Default___.Fm0(p0, globalState)
				(globalState).F22 = (_873_v38).Fm21((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), globalState)
				var _874_v39 _dafny.Array
				_ = _874_v39
				var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
				_ = _nw167
				_874_v39 = _nw167
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_874_v39), 0))
				_ = _index176
				(_874_v39).ArraySet1(_833_v1, (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(967), _dafny.ArrayLen((_874_v39), 0))
				_ = _index177
				(_874_v39).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_850_v17, _dafny.Companion_Sequence_.Concatenate(_850_v17, _850_v17))), (_index177).Int())
			}
			var _875_v40 *C0
			_ = _875_v40
			var _nw168 *C0 = New_C0_()
			_ = _nw168
			_nw168.Ctor__(!((_837_v5).F32()) || (!(p3)))
			_875_v40 = _nw168
			var _876_v41 _dafny.Array
			_ = _876_v41
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw169
			_876_v41 = _nw169
			_876_v41 = (func() _dafny.Array {
				if (p0).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm34(p1, (_this.F33).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32()).(_dafny.CodePoint), (_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), globalState)).Cardinality())) < 0 {
					return _876_v41
				}
				return _876_v41
			})()
			var _877_v42 _dafny.Sequence
			_ = _877_v42
			_877_v42 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_850_v17, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_850_v17).Cardinality()))).Uint32(), p1), _850_v17), _dafny.SeqOf((_837_v5).F32(), false, (_875_v40).F32(), (_this).Fm31(_dafny.MultiSetOf(p2), globalState)))
			var _878_v43 _dafny.Map
			_ = _878_v43
			_878_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_875_v40).F32(), _834_v2)
			var _879_v44 _dafny.Set
			_ = _879_v44
			_879_v44 = _dafny.SetOf((func() _dafny.Array {
				if (_878_v43).Contains(p3) {
					return (_878_v43).Get(p3).(_dafny.Array)
				}
				return _834_v2
			})())
			var _880_v45 D4
			_ = _880_v45
			_880_v45 = Companion_D4_.Create_DC13_(_846_v14)
			var _881_v46 D4
			_ = _881_v46
			_881_v46 = Companion_D4_.Create_DC15_(_880_v45)
			var _882_v47 D7
			_ = _882_v47
			_882_v47 = Companion_D7_.Create_DC20_(_847_v15, _834_v2, _881_v46, (_837_v5).F32(), (_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int))
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
			_ = _index178
			var _rhs190 _dafny.Sequence = _dafny.SeqOf(_850_v17, _850_v17, _dafny.SeqOf(!(p3)), _850_v17)
			_ = _rhs190
			var _rhs191 _dafny.Int = Companion_Default___.SafeModuloInt((p0).Times(p0), p0)
			_ = _rhs191
			var _rhs192 bool = (_879_v44).IsDisjointFrom(_dafny.SetOf(_834_v2, _834_v2, (_882_v47).Dtor_cf31(), _834_v2))
			_ = _rhs192
			var _lhs173 _dafny.Array = _832_v0
			_ = _lhs173
			var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			_877_v42 = _rhs190
			(_lhs173).ArraySet1(_rhs191, (_lhs174).Int())
			_lhs175.F22 = _rhs192
			(globalState).F7 = (_837_v5).F32()
		} else {
			var _883_v48 _dafny.Map
			_ = _883_v48
			_883_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_837_v5).F32(), (_837_v5).F32())
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
			_ = _index179
			(_832_v0).ArraySet1((_dafny.Zero).Minus((p0).Plus(((_883_v48).Update(p1, (_837_v5).F32())).Cardinality())), (_index179).Int())
			var _884_v49 _dafny.Map
			_ = _884_v49
			_884_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F33, p1)
			_883_v48 = (_883_v48).Update((_837_v5).F32(), (func() bool {
				if (_884_v49).Contains(_this.F33) {
					return (_884_v49).Get(_this.F33).(bool)
				}
				return p1
			})())
			var _885_v50 *C1
			_ = _885_v50
			var _nw170 *C1 = New_C1_()
			_ = _nw170
			_nw170.Ctor__()
			_885_v50 = _nw170
			var _886_v51 _dafny.Map
			_ = _886_v51
			_886_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('t'), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_847_v15, _dafny.SeqOf(_dafny.IntOfUint32((_this.F33).Cardinality()), p0)), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F33).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_847_v15, _dafny.SeqOf(_dafny.IntOfUint32((_this.F33).Cardinality()), p0))).Cardinality()))).Uint32(), _847_v15)).Cardinality()))
			var _887_v52 _dafny.Map
			_ = _887_v52
			_887_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_886_v51, Companion_Default___.Fm35((_832_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))).Int()).(_dafny.Int), p3, !((_837_v5).F32()), globalState))
			_887_v52 = (func() _dafny.Map {
				if _dafny.Companion_Sequence_.IsProperPrefixOf(_this.F33, _dafny.UnicodeSeqOfUtf8Bytes("nos")) {
					return (_887_v52).Update(_886_v51, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-70))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_888_v14 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_889_i1 _dafny.Int) _dafny.CodePoint {
							return _888_v14
						}
					})(_846_v14))))
				}
				return (_887_v52).Merge(func() _dafny.Map {
					var _coll24 = _dafny.NewMapBuilder()
					_ = _coll24
					for _iter25 := _dafny.Iterate((Companion_Default___.Fm36(_846_v14, Companion_D3_.Create_DC10_(_833_v1), p1, globalState)).Elements()); ; {
						_compr_24, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _890_v53 _dafny.Map
						_890_v53 = interface{}(_compr_24).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm36(_846_v14, Companion_D3_.Create_DC10_(_833_v1), p1, globalState), _890_v53) {
							_coll24.Add(_890_v53, _dafny.UnicodeSeqOfUtf8Bytes("wqaidjcb"))
						}
					}
					return _coll24.ToMap()
				}())
			})()
			(globalState).F7 = p1
		}
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(146), _dafny.ArrayLen((_832_v0), 0))
		_ = _index180
		(_832_v0).ArraySet1(_dafny.IntOfInt64(921), (_index180).Int())
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 bool = false
		_ = r1
		var _891_v0 _dafny.Array
		_ = _891_v0
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_19
		var _nw171 _dafny.Array
		_ = _nw171
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw171 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) D7 = (func(_892_p2 bool, _893_p1 _dafny.Int) func(_dafny.Int) D7 {
				return func(_894_i0 _dafny.Int) D7 {
					return Companion_D7_.Create_DC18_(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_892_p2, _892_p2)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _893_p1)))
				}
			})(p2, p1)
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw171 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw171).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw171).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_891_v0 = _nw171
		var _895_v1 _dafny.CodePoint
		_ = _895_v1
		_895_v1 = _dafny.CodePoint('k')
		var _896_v2 _dafny.Sequence
		_ = _896_v2
		_896_v2 = _dafny.SeqOf(p2, p2, p2, p2, p2)
		var _897_v3 _dafny.Map
		_ = _897_v3
		_897_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_895_v1, _dafny.IntOfUint32((_896_v2).Cardinality()))
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_891_v0), 0))
		_ = _index181
		(_891_v0).ArraySet1(Companion_D7_.Create_DC18_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_897_v3), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(_897_v3)).Cardinality()))).Uint32(), _897_v3)), (_index181).Int())
		var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_891_v0), 0))
		_ = _index182
		(_891_v0).ArraySet1(Companion_Default___.Fm37(p1, p1, p1, globalState), (_index182).Int())
		(_this).F33 = _dafny.UnicodeSeqOfUtf8Bytes("xcb")
		var _898_v4 _dafny.Array
		_ = _898_v4
		var _len0_20 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_20
		var _nw172 _dafny.Array
		_ = _nw172
		if _len0_20.Cmp(_dafny.Zero) == 0 {
			_nw172 = _dafny.NewArray(_len0_20)
		} else {
			var _init20 func(_dafny.Int) _dafny.Int = (func(_899_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_900_i1 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_900_i1, _899_p1)
				}
			})(p1)
			_ = _init20
			var _element0_20 = _init20(_dafny.Zero)
			_ = _element0_20
			_nw172 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
			(_nw172).ArraySet1(_element0_20, 0)
			var _nativeLen0_20 = (_len0_20).Int()
			_ = _nativeLen0_20
			for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
				(_nw172).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
			}
		}
		_898_v4 = _nw172
		var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))
		_ = _index183
		(_898_v4).ArraySet1(p1, (_index183).Int())
		var _901_v5 _dafny.Set
		_ = _901_v5
		_901_v5 = _dafny.SetOf(p1, p1, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))
		var _902_v6 _dafny.MultiSet
		_ = _902_v6
		_902_v6 = _dafny.MultiSetOf(Companion_Default___.Fm0((_901_v5).Cardinality(), globalState), p1)
		var _903_v7 D8
		_ = _903_v7
		_903_v7 = Companion_D8_.Create_DC22_(p1, p2, _895_v1)
		var _904_v8 _dafny.Map
		_ = _904_v8
		_904_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		var _905_v10 D1
		_ = _905_v10
		_905_v10 = Companion_D1_.Create_DC6_(p1, _dafny.SeqOf(_896_v2, _dafny.SeqOf(p2)), p1, p2)
		var _906_v11 _dafny.Array
		_ = _906_v11
		var _nwElement0_46 bool = ((func() _dafny.Int {
			if (_902_v6).Contains((_903_v7).Dtor_cf36()) {
				return (_902_v6).Multiplicity((_903_v7).Dtor_cf36())
			}
			return p1
		})()).Cmp(_dafny.IntOfUint32((_this.F33).Cardinality())) >= 0
		_ = _nwElement0_46
		var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(3))
		_ = _nw173
		(_nw173).ArraySet1(_nwElement0_46, 0)
		(_nw173).ArraySet1((func() _dafny.Set {
			var _coll25 = _dafny.NewBuilder()
			_ = _coll25
			for _iter26 := _dafny.Iterate((_904_v8).Keys().Elements()); ; {
				_compr_25, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _907_v9 _dafny.Int
				_907_v9 = interface{}(_compr_25).(_dafny.Int)
				if (_904_v8).Contains(_907_v9) {
					_coll25.Add((_907_v9).Times(_dafny.IntOfInt64(-326)))
				}
			}
			return _coll25.ToSet()
		}()).IsProperSubsetOf(_901_v5), 1)
		(_nw173).ArraySet1((_905_v10).Dtor_cf16(), 2)
		_906_v11 = _nw173
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_906_v11), 0))
		_ = _index184
		(_906_v11).ArraySet1(!(p2), (_index184).Int())
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_898_v4), 0))
		_ = _index185
		(_898_v4).ArraySet1(p1, (_index185).Int())
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))
		_ = _index186
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_906_v11), 0))
		_ = _index187
		var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_898_v4), 0))
		_ = _index188
		var _rhs193 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_896_v2, _dafny.Companion_Sequence_.Update(_896_v2, (Companion_Default___.SafeIndex((_this).Fm2(globalState), _dafny.IntOfUint32((_896_v2).Cardinality()))).Uint32(), (_this).Fm31(_902_v6, globalState)))).Cardinality())
		_ = _rhs193
		var _rhs194 bool = !(!((_dafny.SetOf(p2, p2)).IsDisjointFrom(Companion_Default___.Fm38(_dafny.IntOfUint32((p0).Cardinality()), true, globalState)))) || (p2)
		_ = _rhs194
		var _rhs195 _dafny.Int = (p1).Times(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-561), p1))
		_ = _rhs195
		var _rhs196 _dafny.Int = (_902_v6).Cardinality()
		_ = _rhs196
		var _rhs197 _dafny.Int = p1
		_ = _rhs197
		var _lhs176 _dafny.Array = _898_v4
		_ = _lhs176
		var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))
		_ = _lhs177
		var _lhs178 _dafny.Array = _906_v11
		_ = _lhs178
		var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_906_v11), 0))
		_ = _lhs179
		var _lhs180 *GlobalState = globalState
		_ = _lhs180
		var _lhs181 _dafny.Array = _898_v4
		_ = _lhs181
		var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(405), _dafny.ArrayLen((_898_v4), 0))
		_ = _lhs182
		var _lhs183 *GlobalState = globalState
		_ = _lhs183
		(_lhs176).ArraySet1(_rhs193, (_lhs177).Int())
		(_lhs178).ArraySet1(_rhs194, (_lhs179).Int())
		_lhs180.F10 = _rhs195
		(_lhs181).ArraySet1(_rhs196, (_lhs182).Int())
		_lhs183.F10 = _rhs197
		var _ingredients0 = _dafny.NewBuilder()
		_ = _ingredients0
		for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_898_v4), 0))); ; {
			_guard_loop_1, _ok27 := _iter27()
			if !_ok27 {
				break
			}
			var _908_i2 _dafny.Int
			_908_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_908_i2).Sign() != -1) && ((_908_i2).Cmp(_dafny.ArrayLen((_898_v4), 0)) < 0)) {
				_ingredients0.Add(_dafny.TupleOf(_898_v4, (_908_i2).Int(), Companion_Default___.SafeModuloInt(_908_i2, (_898_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))).Int()).(_dafny.Int))))
			}
		}
		for _iter28 := _dafny.Iterate(_ingredients0); ; {
			_tup0, _ok28 := _iter28()
			if !_ok28 {
				break
			}
			(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
		}
		var _909_v12 D0
		_ = _909_v12
		_909_v12 = Companion_D0_.Create_DC0_(_this.F33)
		var _910_v13 _dafny.Map
		_ = _910_v13
		_910_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_898_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))).Int()).(_dafny.Int), _this.F33)
		var _911_v14 _dafny.Array
		_ = _911_v14
		var _nwElement0_47 _dafny.Sequence = (_909_v12).Dtor_cf0()
		_ = _nwElement0_47
		var _nw174 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(20))
		_ = _nw174
		(_nw174).ArraySet1(_nwElement0_47, 0)
		(_nw174).ArraySet1((Companion_D0_.Create_DC0_(_this.F33)).Dtor_cf0(), 1)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("gbvposwab"), 2)
		(_nw174).ArraySet1(_dafny.Companion_Sequence_.Update(_this.F33, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32(), _895_v1), 3)
		(_nw174).ArraySet1(_this.F33, 4)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iwtorenfq"), 5)
		(_nw174).ArraySet1(_this.F33, 6)
		(_nw174).ArraySet1((func() _dafny.Sequence {
			if (_910_v13).Contains(p1) {
				return (_910_v13).Get(p1).(_dafny.Sequence)
			}
			return _this.F33
		})(), 7)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("scc"), 8)
		(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qltrd"), _this.F33), 9)
		(_nw174).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(30))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}((func(_912_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_913_i3 _dafny.Int) _dafny.CodePoint {
				return _912_v1
			}
		})(_895_v1))), 10)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ssiokybl"), 11)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("atla"), 12)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("j"), 13)
		(_nw174).ArraySet1(_dafny.Companion_Sequence_.Update(_this.F33, (Companion_Default___.SafeIndex((_904_v8).Cardinality(), _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32(), _895_v1), 14)
		(_nw174).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_this.F33, Companion_Default___.Fm35((_898_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_898_v4), 0))).Int()).(_dafny.Int), !(p2), (_906_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_906_v11), 0))).Int()).(bool), globalState)), 15)
		(_nw174).ArraySet1(_this.F33, 16)
		(_nw174).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uokfrfjh"), 17)
		(_nw174).ArraySet1(_this.F33, 18)
		(_nw174).ArraySet1(Companion_Default___.Fm35(_dafny.IntOfInt64(308), p2, false, globalState), 19)
		_911_v14 = _nw174
		var _914_v15 _dafny.Map
		_ = _914_v15
		_914_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_906_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(951), _dafny.ArrayLen((_906_v11), 0))).Int()).(bool))
		var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_911_v14), 0))
		_ = _index189
		(_911_v14).ArraySet1(_dafny.Companion_Sequence_.Update(_this.F33, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_this.F33).Cardinality()), _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32(), (_this.F33).Select((Companion_Default___.SafeIndex((_914_v15).Cardinality(), _dafny.IntOfUint32((_this.F33).Cardinality()))).Uint32()).(_dafny.CodePoint)), (_index189).Int())
		var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_911_v14), 0))
		_ = _index190
		(_911_v14).ArraySet1(_this.F33, (_index190).Int())
		var _915_v16 _dafny.Set
		_ = _915_v16
		_915_v16 = _dafny.SetOf(p2, (_dafny.MultiSetOf(_dafny.IntOfUint32((_896_v2).Cardinality()))).IsProperSubsetOf(_902_v6))
		var _916_v17 _dafny.Sequence
		_ = _916_v17
		_916_v17 = _dafny.SeqOf(_915_v16)
		var _rhs198 _dafny.Set = ((Companion_Default___.Fm38(p1, p2, globalState)).Intersection(_915_v16)).Difference((_916_v17).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_916_v17).Cardinality()))).Uint32()).(_dafny.Set))
		_ = _rhs198
		var _rhs199 _dafny.CodePoint = _895_v1
		_ = _rhs199
		_915_v16 = _rhs198
		_895_v1 = _rhs199
		r0 = Companion_D0_.Create_DC1_(p1, Companion_Default___.SafeModuloInt(p1, p1), p1, _906_v11)
		r1 = p2
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	dummy byte
}

func New_C3_() *C3 {
	_this := C3{}

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

func (_this *C3) Ctor__() {
	{
	}
}
func (_this *C3) Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(false), true, false), !(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, !(true)), !(true)))
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Int {
	{
		var _source16 D1 = Companion_D1_.Create_DC6_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(24), _dafny.IntOfInt64(182))); ; {
				_compr_26, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _917_v0 _dafny.Int
				_917_v0 = interface{}(_compr_26).(_dafny.Int)
				if ((_dafny.IntOfInt64(24)).Cmp(_917_v0) <= 0) && ((_917_v0).Cmp(_dafny.IntOfInt64(182)) < 0) {
					_coll26.Add(Companion_Default___.SafeModuloInt(_917_v0, _dafny.IntOfInt64(21)))
				}
			}
			return _coll26.ToSet()
		}()).Cardinality())).Cardinality(), _dafny.SeqOf(_dafny.SeqOf(false, false), _dafny.SeqOf(false)), _dafny.IntOfInt64(423), !(true))
		_ = _source16
		if _source16.Is_DC5() {
			var _918___mcc_h0 _dafny.Int = _source16.Get_().(D1_DC5).Cf9
			_ = _918___mcc_h0
			var _919___mcc_h1 bool = _source16.Get_().(D1_DC5).Cf10
			_ = _919___mcc_h1
			var _920___mcc_h2 bool = _source16.Get_().(D1_DC5).Cf11
			_ = _920___mcc_h2
			var _921___mcc_h3 _dafny.Map = _source16.Get_().(D1_DC5).Cf12
			_ = _921___mcc_h3
			var _922_cf12 _dafny.Map = _921___mcc_h3
			_ = _922_cf12
			var _923_cf11 bool = _920___mcc_h2
			_ = _923_cf11
			var _924_cf10 bool = _919___mcc_h1
			_ = _924_cf10
			var _925_cf9 _dafny.Int = _918___mcc_h0
			_ = _925_cf9
			return Companion_Default___.SafeDivisionInt(_925_cf9, _925_cf9)
		} else if _source16.Is_DC6() {
			var _926___mcc_h4 _dafny.Int = _source16.Get_().(D1_DC6).Cf13
			_ = _926___mcc_h4
			var _927___mcc_h5 _dafny.Sequence = _source16.Get_().(D1_DC6).Cf14
			_ = _927___mcc_h5
			var _928___mcc_h6 _dafny.Int = _source16.Get_().(D1_DC6).Cf15
			_ = _928___mcc_h6
			var _929___mcc_h7 bool = _source16.Get_().(D1_DC6).Cf16
			_ = _929___mcc_h7
			var _930_cf16 bool = _929___mcc_h7
			_ = _930_cf16
			var _931_cf15 _dafny.Int = _928___mcc_h6
			_ = _931_cf15
			var _932_cf14 _dafny.Sequence = _927___mcc_h5
			_ = _932_cf14
			var _933_cf13 _dafny.Int = _926___mcc_h4
			_ = _933_cf13
			return _933_cf13
		} else {
			var _934___mcc_h8 _dafny.Map = _source16.Get_().(D1_DC4).Cf8
			_ = _934___mcc_h8
			var _935_cf8 _dafny.Map = _934___mcc_h8
			_ = _935_cf8
			return Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vpngnpve")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(872), _dafny.IntOfInt64(-382))).Cardinality())
		}
	}
}
func (_this *C3) Fm8(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yxgl")).Cardinality())).Times(_dafny.IntOfInt64(-516)), (_dafny.IntOfInt64(183)).Times(_dafny.IntOfUint32(((Companion_D0_.Create_DC0_(_dafny.UnicodeSeqOfUtf8Bytes("ulgojwmoa"))).Dtor_cf0()).Cardinality())))
	}
}
func (_this *C3) Fm9(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_(((Companion_D1_.Create_DC4_(func() _dafny.Map {
			var _coll27 = _dafny.NewMapBuilder()
			_ = _coll27
			for _iter30 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, !(false), true, false, false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-510), (func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(521), _dafny.IntOfInt64(997))); ; {
					_compr_28, _ok31 := _iter31()
					if !_ok31 {
						break
					}
					var _936_v1 _dafny.Int
					_936_v1 = interface{}(_compr_28).(_dafny.Int)
					if ((_dafny.IntOfInt64(521)).Cmp(_936_v1) <= 0) && ((_936_v1).Cmp(_dafny.IntOfInt64(997)) < 0) {
						_coll28.Add(Companion_Default___.SafeDivisionInt(_936_v1, _dafny.IntOfInt64(433)), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
					}
				}
				return _coll28.ToMap()
			}()).Cardinality())).Cardinality())).Keys().Elements()); ; {
				_compr_27, _ok30 := _iter30()
				if !_ok30 {
					break
				}
				var _937_v0 _dafny.MultiSet
				_937_v0 = interface{}(_compr_27).(_dafny.MultiSet)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, !(false), true, false, false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-510), (func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(521), _dafny.IntOfInt64(997))); ; {
						_compr_29, _ok32 := _iter32()
						if !_ok32 {
							break
						}
						var _936_v1 _dafny.Int
						_936_v1 = interface{}(_compr_29).(_dafny.Int)
						if ((_dafny.IntOfInt64(521)).Cmp(_936_v1) <= 0) && ((_936_v1).Cmp(_dafny.IntOfInt64(997)) < 0) {
							_coll29.Add(Companion_Default___.SafeDivisionInt(_936_v1, _dafny.IntOfInt64(433)), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
						}
					}
					return _coll29.ToMap()
				}()).Cardinality())).Cardinality())).Contains(_937_v0) {
					_coll27.Add(_937_v0, false)
				}
			}
			return _coll27.ToMap()
		}())).Dtor_cf8()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, false, false, false), false)))
	}
}
func (_this *C3) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _938_v0 _dafny.Map
		_ = _938_v0
		_938_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(604))
		var _939_v1 _dafny.Map
		_ = _939_v1
		_939_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_938_v0).Cardinality())
		var _940_v2 _dafny.Sequence
		_ = _940_v2
		_940_v2 = _dafny.SeqOf(_938_v0)
		(globalState).F22 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_940_v2, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_940_v2).Cardinality()))).Uint32(), (Companion_D2_.Create_DC7_(_939_v1)).Dtor_cf17()), _dafny.Companion_Sequence_.Concatenate(_940_v2, _940_v2))
		var _941_v3 _dafny.Array
		_ = _941_v3
		var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
		_ = _nw175
		_941_v3 = _nw175
		var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
		_ = _index191
		(_941_v3).ArraySet1(p0, (_index191).Int())
		var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
		_ = _index192
		(_941_v3).ArraySet1(Companion_Default___.SafeModuloInt(p2, p0), (_index192).Int())
		var _942_v4 _dafny.Sequence
		_ = _942_v4
		_942_v4 = _dafny.UnicodeSeqOfUtf8Bytes("iiynvtexy")
		var _943_v5 D0
		_ = _943_v5
		_943_v5 = Companion_D0_.Create_DC0_(_942_v4)
		var _944_v6 _dafny.CodePoint
		_ = _944_v6
		_944_v6 = _dafny.CodePoint('j')
		var _945_v7 _dafny.Map
		_ = _945_v7
		_945_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _942_v4)
		var _pat_let_tv19 = _942_v4
		_ = _pat_let_tv19
		var _pat_let_tv20 = _942_v4
		_ = _pat_let_tv20
		var _pat_let_tv21 = _943_v5
		_ = _pat_let_tv21
		var _pat_let_tv22 = _942_v4
		_ = _pat_let_tv22
		var _946_v8 _dafny.Array
		_ = _946_v8
		var _nwElement0_48 D0 = _943_v5
		_ = _nwElement0_48
		var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(16))
		_ = _nw176
		(_nw176).ArraySet1(_nwElement0_48, 0)
		(_nw176).ArraySet1(_943_v5, 1)
		(_nw176).ArraySet1(Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Update(_942_v4, (Companion_Default___.SafeIndex((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_942_v4).Cardinality()))).Uint32(), _944_v6)), 2)
		(_nw176).ArraySet1(_943_v5, 3)
		(_nw176).ArraySet1(_943_v5, 4)
		(_nw176).ArraySet1(func(_pat_let18_0 D0) D0 {
			return func(_949_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let21_0 _dafny.Sequence) D0 {
					return func(_950_dt__update_hcf0_h1 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_950_dt__update_hcf0_h1)
					}(_pat_let21_0)
				}(_pat_let_tv20)
			}(_pat_let18_0)
		}(func(_pat_let19_0 D0) D0 {
			return func(_947_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let20_0 _dafny.Sequence) D0 {
					return func(_948_dt__update_hcf0_h0 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_948_dt__update_hcf0_h0)
					}(_pat_let20_0)
				}(_pat_let_tv19)
			}(_pat_let19_0)
		}(_943_v5)), 5)
		(_nw176).ArraySet1(Companion_D0_.Create_DC0_(Companion_Default___.Fm10(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), p0), Companion_Default___.Fm11((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), globalState), p3, globalState)), 6)
		(_nw176).ArraySet1(_943_v5, 7)
		(_nw176).ArraySet1(_943_v5, 8)
		(_nw176).ArraySet1(_943_v5, 9)
		(_nw176).ArraySet1(_943_v5, 10)
		(_nw176).ArraySet1(func(_pat_let22_0 D0) D0 {
			return func(_951_dt__update__tmp_h2 D0) D0 {
				return func(_pat_let23_0 _dafny.Sequence) D0 {
					return func(_952_dt__update_hcf0_h2 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_952_dt__update_hcf0_h2)
					}(_pat_let23_0)
				}((_pat_let_tv21).Dtor_cf0())
			}(_pat_let22_0)
		}(_943_v5), 11)
		(_nw176).ArraySet1(_943_v5, 12)
		(_nw176).ArraySet1(func(_pat_let24_0 D0) D0 {
			return func(_953_dt__update__tmp_h3 D0) D0 {
				return func(_pat_let25_0 _dafny.Sequence) D0 {
					return func(_954_dt__update_hcf0_h3 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_954_dt__update_hcf0_h3)
					}(_pat_let25_0)
				}(_pat_let_tv22)
			}(_pat_let24_0)
		}(Companion_D0_.Create_DC0_((func() _dafny.Sequence {
			if (_945_v7).Contains(p3) {
				return (_945_v7).Get(p3).(_dafny.Sequence)
			}
			return _942_v4
		})())), 13)
		(_nw176).ArraySet1(_943_v5, 14)
		(_nw176).ArraySet1(_943_v5, 15)
		_946_v8 = _nw176
		var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_946_v8), 0))
		_ = _index193
		(_946_v8).ArraySet1(_943_v5, (_index193).Int())
		var _955_v9 _dafny.Sequence
		_ = _955_v9
		_955_v9 = _dafny.SeqOf(true, p1, !(p1))
		var _pat_let_tv23 = _942_v4
		_ = _pat_let_tv23
		var _pat_let_tv24 = p2
		_ = _pat_let_tv24
		var _pat_let_tv25 = _942_v4
		_ = _pat_let_tv25
		var _pat_let_tv26 = _944_v6
		_ = _pat_let_tv26
		var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_946_v8), 0))
		_ = _index194
		(_946_v8).ArraySet1(func(_pat_let26_0 D0) D0 {
			return func(_956_dt__update__tmp_h4 D0) D0 {
				return func(_pat_let27_0 _dafny.Sequence) D0 {
					return func(_957_dt__update_hcf0_h4 _dafny.Sequence) D0 {
						return Companion_D0_.Create_DC0_(_957_dt__update_hcf0_h4)
					}(_pat_let27_0)
				}(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_pat_let_tv23, _dafny.UnicodeSeqOfUtf8Bytes("q")), (Companion_Default___.SafeIndex(_pat_let_tv24, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_pat_let_tv25, _dafny.UnicodeSeqOfUtf8Bytes("q"))).Cardinality()))).Uint32(), _pat_let_tv26))
			}(_pat_let26_0)
		}((func() D0 {
			if !((_955_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_955_v9).Cardinality()))).Uint32()).(bool)) {
				return _943_v5
			}
			return Companion_D0_.Create_DC0_((_943_v5).Dtor_cf0())
		})()), (_index194).Int())
		var _958_v10 _dafny.Map
		_ = _958_v10
		_958_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_955_v9).Cardinality()), p3)
		var _959_v11 _dafny.Map
		_ = _959_v11
		_959_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_958_v10).Update((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), p1)).Cardinality(), _941_v3)
		_959_v11 = (_959_v11).Update(p0, _941_v3)
		if (p3) == (!(p1)) {
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
			_ = _index195
			(_941_v3).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm10((_938_v0).Merge(_939_v1), _dafny.CodePoint('s'), p3, globalState)).Cardinality()), (_index195).Int())
			var _960_v12 _dafny.MultiSet
			_ = _960_v12
			_960_v12 = _dafny.MultiSetOf(p2)
			var _961_v13 _dafny.Map
			_ = _961_v13
			_961_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v12, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}((func(_962_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_963_i0 _dafny.Int) _dafny.Int {
					return _962_p0
				}
			})(p0))), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_964_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_965_i0 _dafny.Int) _dafny.Int {
					return _964_p0
				}
			})(p0)))).Cardinality()))).Uint32(), p0))
			var _966_v14 D1
			_ = _966_v14
			_966_v14 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_955_v9).Cardinality()), p1, p1, _961_v13)
			var _pat_let_tv27 = p3
			_ = _pat_let_tv27
			var _pat_let_tv28 = p2
			_ = _pat_let_tv28
			var _pat_let_tv29 = _961_v13
			_ = _pat_let_tv29
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
			_ = _index196
			(_941_v3).ArraySet1((_dafny.Zero).Minus(((_958_v10).Cardinality()).Minus((func(_pat_let28_0 D1) D1 {
				return func(_967_dt__update__tmp_h5 D1) D1 {
					return func(_pat_let29_0 bool) D1 {
						return func(_968_dt__update_hcf11_h0 bool) D1 {
							return func(_pat_let30_0 _dafny.Int) D1 {
								return func(_969_dt__update_hcf9_h0 _dafny.Int) D1 {
									return func(_pat_let31_0 _dafny.Map) D1 {
										return func(_970_dt__update_hcf12_h0 _dafny.Map) D1 {
											return Companion_D1_.Create_DC5_(_969_dt__update_hcf9_h0, (_967_dt__update__tmp_h5).Dtor_cf10(), _968_dt__update_hcf11_h0, _970_dt__update_hcf12_h0)
										}(_pat_let31_0)
									}(_pat_let_tv29)
								}(_pat_let30_0)
							}(_pat_let_tv28)
						}(_pat_let29_0)
					}(_pat_let_tv27)
				}(_pat_let28_0)
			}(_966_v14)).Dtor_cf9())), (_index196).Int())
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
			_ = _index197
			(_941_v3).ArraySet1(p2, (_index197).Int())
			(globalState).F22 = Companion_Default___.Fm12(p1, p1, globalState)
			var _971_v15 _dafny.Map
			_ = _971_v15
			_971_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_942_v4), p1)
			var _972_v16 _dafny.Array
			_ = _972_v16
			var _nwElement0_49 bool = p1
			_ = _nwElement0_49
			var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(21))
			_ = _nw177
			(_nw177).ArraySet1(_nwElement0_49, 0)
			(_nw177).ArraySet1(p1, 1)
			(_nw177).ArraySet1((func() bool {
				if (_971_v15).Contains((_946_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_946_v8), 0))).Int()).(D0)) {
					return (_971_v15).Get((_946_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_946_v8), 0))).Int()).(D0)).(bool)
				}
				return p1
			})(), 2)
			(_nw177).ArraySet1(p1, 3)
			(_nw177).ArraySet1(false, 4)
			(_nw177).ArraySet1(p3, 5)
			(_nw177).ArraySet1(true, 6)
			(_nw177).ArraySet1(p3, 7)
			(_nw177).ArraySet1(false, 8)
			(_nw177).ArraySet1(p3, 9)
			(_nw177).ArraySet1(p1, 10)
			(_nw177).ArraySet1(p3, 11)
			(_nw177).ArraySet1(p3, 12)
			(_nw177).ArraySet1(p1, 13)
			(_nw177).ArraySet1(p1, 14)
			(_nw177).ArraySet1(p1, 15)
			(_nw177).ArraySet1(true, 16)
			(_nw177).ArraySet1(p1, 17)
			(_nw177).ArraySet1(p3, 18)
			(_nw177).ArraySet1(p1, 19)
			(_nw177).ArraySet1(false, 20)
			_972_v16 = _nw177
			var _973_v17 _dafny.Map
			_ = _973_v17
			_973_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _972_v16)
			var _974_v18 _dafny.Sequence
			_ = _974_v18
			_974_v18 = _dafny.SeqOf(_972_v16, _972_v16, _972_v16)
			_973_v17 = (_973_v17).Update(p0, (_974_v18).Select((Companion_Default___.SafeIndex((_this).Fm2(globalState), _dafny.IntOfUint32((_974_v18).Cardinality()))).Uint32()).(_dafny.Array))
		} else {
			var _975_v19 *C0
			_ = _975_v19
			var _nw178 *C0 = New_C0_()
			_ = _nw178
			_nw178.Ctor__(p3)
			_975_v19 = _nw178
			_942_v4 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_976_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_977_i1 _dafny.Int) _dafny.CodePoint {
					return _976_v6
				}
			})(_944_v6)))
			(globalState).F10 = _dafny.IntOfInt64(718)
			if Companion_Default___.Fm12(!(p3), (_975_v19).F32(), globalState) {
				var _978_v20 _dafny.Set
				_ = _978_v20
				_978_v20 = _dafny.SetOf(p3, p3, (_975_v19).F32(), true, p3)
				var _979_v21 _dafny.Sequence
				_ = _979_v21
				_979_v21 = _dafny.SeqOf(p0)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))
				_ = _index198
				(_941_v3).ArraySet1((((_945_v7).Cardinality()).Times((_978_v20).Cardinality())).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_979_v21, Companion_Default___.Fm14(false, _dafny.IntOfInt64(-248), p2, globalState))).Cardinality())), (_index198).Int())
				var _980_v22 bool
				_ = _980_v22
				var _981_v23 _dafny.Int
				_ = _981_v23
				var _982_v24 T0
				_ = _982_v24
				var _out5 bool
				_ = _out5
				var _out6 _dafny.Int
				_ = _out6
				var _out7 T0
				_ = _out7
				_out5, _out6, _out7 = (_this).M4(_942_v4, _dafny.IntOfUint32((_979_v21).Cardinality()), globalState)
				_980_v22 = _out5
				_981_v23 = _out6
				_982_v24 = _out7
				(globalState).F22 = (_955_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_955_v9).Cardinality()))).Uint32()).(bool)
				var _983_v25 _dafny.Map
				_ = _983_v25
				_983_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_941_v3, _dafny.CodePoint('u'))
				_983_v25 = (_983_v25).Update(_941_v3, _944_v6)
				var _984_v26 *C0
				_ = _984_v26
				var _nw179 *C0 = New_C0_()
				_ = _nw179
				_nw179.Ctor__(p1)
				_984_v26 = _nw179
			} else {
				var _985_v27 *C0
				_ = _985_v27
				var _nw180 *C0 = New_C0_()
				_ = _nw180
				_nw180.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), _955_v9)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_955_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-175), _dafny.IntOfUint32((_955_v9).Cardinality()))).Uint32()).(bool)), _955_v9)))
				_985_v27 = _nw180
				(globalState).F22 = !((_985_v27).F32())
				(globalState).F13 = p1
				_958_v10 = (_958_v10).Update(p0, !((_975_v19).F32()))
				(globalState).F10 = Companion_Default___.SafeDivisionInt(((_958_v10).Cardinality()).Plus(_dafny.IntOfUint32((_942_v4).Cardinality())), ((_dafny.Zero).Minus(p2)).Plus((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int)))
			}
			(globalState).F19 = (func() _dafny.CodePoint {
				if (_975_v19).F32() {
					return _944_v6
				}
				return (func() _dafny.CodePoint {
					if (_975_v19).F32() {
						return _944_v6
					}
					return _944_v6
				})()
			})()
		}
		var _986_v28 _dafny.Array
		_ = _986_v28
		var _nwElement0_50 _dafny.Map = _958_v10
		_ = _nwElement0_50
		var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(14))
		_ = _nw181
		(_nw181).ArraySet1(_nwElement0_50, 0)
		(_nw181).ArraySet1((_958_v10).Merge((_958_v10).Update((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), true)), 1)
		(_nw181).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), p3), 2)
		(_nw181).ArraySet1((_958_v10).Merge(_958_v10), 3)
		(_nw181).ArraySet1(_958_v10, 4)
		(_nw181).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), p1), 5)
		(_nw181).ArraySet1(_958_v10, 6)
		(_nw181).ArraySet1((func() _dafny.Map {
			if p1 {
				return _958_v10
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), p1)
		})(), 7)
		(_nw181).ArraySet1(_958_v10, 8)
		(_nw181).ArraySet1(_958_v10, 9)
		(_nw181).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), !(p1))).Merge(_958_v10), 10)
		(_nw181).ArraySet1((_958_v10).Merge((_958_v10).Update(p2, p1)), 11)
		(_nw181).ArraySet1((_958_v10).Merge((_958_v10).Update((_941_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(663), _dafny.ArrayLen((_941_v3), 0))).Int()).(_dafny.Int), !(p3))), 12)
		(_nw181).ArraySet1(_958_v10, 13)
		_986_v28 = _nw181
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_986_v28), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _987_i2 _dafny.Int
			_987_i2 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_987_i2).Sign() != -1) && ((_987_i2).Cmp(_dafny.ArrayLen((_986_v28), 0)) < 0)) {
				(_986_v28).ArraySet1(_958_v10, (_987_i2).Int())
			}
		}
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 bool = false
		_ = r1
		var _988_v0 _dafny.Array
		_ = _988_v0
		var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
		_ = _nw182
		_988_v0 = _nw182
		var _989_v1 _dafny.Array
		_ = _989_v1
		var _nw183 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
		_ = _nw183
		_989_v1 = _nw183
		var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_988_v0), 0))
		_ = _index199
		(_988_v0).ArraySet1(_989_v1, (_index199).Int())
		var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_988_v0), 0))
		_ = _index200
		(_988_v0).ArraySet1((func() _dafny.Array {
			if Companion_Default___.Fm12(p2, p2, globalState) {
				return _989_v1
			}
			return _989_v1
		})(), (_index200).Int())
		var _990_v2 _dafny.Sequence
		_ = _990_v2
		_990_v2 = _dafny.UnicodeSeqOfUtf8Bytes("nnaijwfws")
		(globalState).F10 = _dafny.IntOfUint32((_990_v2).Cardinality())
		var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))
		_ = _index201
		(_989_v1).ArraySet1(p2, (_index201).Int())
		var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))
		_ = _index202
		(_989_v1).ArraySet1((Companion_D1_.Create_DC6_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ydkxlx")).Cardinality()), _dafny.SeqOf(_dafny.SeqOf(p2, p2, p2, p2, p2)), p1, !(p2))).Dtor_cf16(), (_index202).Int())
		var _991_v3 *C0
		_ = _991_v3
		var _nw184 *C0 = New_C0_()
		_ = _nw184
		_nw184.Ctor__(true)
		_991_v3 = _nw184
		var _992_i0 _dafny.Int
		_ = _992_i0
		_992_i0 = _dafny.Zero
		{
			for p2 {
				{
					if (_992_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_992_i0 = (_992_i0).Plus(_dafny.One)
					(globalState).F13 = (_991_v3).F32()
					var _993_v4 *C0
					_ = _993_v4
					var _nw185 *C0 = New_C0_()
					_ = _nw185
					_nw185.Ctor__(p2)
					_993_v4 = _nw185
					var _994_v5 _dafny.CodePoint
					_ = _994_v5
					_994_v5 = _dafny.CodePoint('y')
					var _995_v6 _dafny.MultiSet
					_ = _995_v6
					_995_v6 = _dafny.MultiSetOf((_993_v4).F32(), !((_991_v3).F32()))
					var _996_v7 *C0
					_ = _996_v7
					var _nw186 *C0 = New_C0_()
					_ = _nw186
					_nw186.Ctor__(!(!((Companion_Default___.Fm15(_994_v5, p1, Companion_Default___.Fm12((_989_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))).Int()).(bool), p2, globalState), globalState)).Dtor_cf21()).Equals(_995_v6)))
					_996_v7 = _nw186
					var _997_v8 _dafny.Map
					_ = _997_v8
					_997_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_994_v5, Companion_D0_.Create_DC2_(p1, false))).Cardinality())
					var _998_v9 _dafny.Sequence
					_ = _998_v9
					_998_v9 = _dafny.SeqOf(p1, ((_dafny.Zero).Minus(p1)).Times(p1), p1, (_dafny.Zero).Minus((func() _dafny.Int {
						if (_997_v8).Contains(p1) {
							return (_997_v8).Get(p1).(_dafny.Int)
						}
						return p1
					})()), _dafny.IntOfInt64(-191))
					var _999_v10 _dafny.Sequence
					_ = _999_v10
					_999_v10 = _dafny.SeqOf((_989_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))).Int()).(bool))
					var _1000_v11 _dafny.Set
					_ = _1000_v11
					_1000_v11 = _dafny.SetOf((_989_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))).Int()).(bool), (_989_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.ArrayLen((_989_v1), 0))).Int()).(bool), p2)
					var _rhs200 bool = (func() bool {
						if p2 {
							return p2
						}
						return _dafny.Companion_Sequence_.Equal(_999_v10, Companion_Default___.Fm16(globalState))
					})()
					_ = _rhs200
					var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_998_v9, _dafny.Companion_Sequence_.Concatenate(_998_v9, _998_v9))
					_ = _rhs201
					var _rhs202 _dafny.Int = p1
					_ = _rhs202
					var _rhs203 bool = p2
					_ = _rhs203
					var _rhs204 _dafny.Int = ((_1000_v11).Intersection((_dafny.SetOf(true)).Difference(_1000_v11))).Cardinality()
					_ = _rhs204
					var _lhs184 *GlobalState = globalState
					_ = _lhs184
					var _lhs185 *GlobalState = globalState
					_ = _lhs185
					var _lhs186 *GlobalState = globalState
					_ = _lhs186
					var _lhs187 *GlobalState = globalState
					_ = _lhs187
					_lhs184.F7 = _rhs200
					_998_v9 = _rhs201
					_lhs185.F10 = _rhs202
					_lhs186.F22 = _rhs203
					_lhs187.F10 = _rhs204
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _1001_v12 _dafny.Array
		_ = _1001_v12
		var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
		_ = _nw187
		_1001_v12 = _nw187
		var _1002_v13 _dafny.CodePoint
		_ = _1002_v13
		_1002_v13 = _dafny.CodePoint('h')
		var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1001_v12), 0))
		_ = _index203
		(_1001_v12).ArraySet1CodePoint(_1002_v13, (_index203).Int())
		var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1001_v12), 0))
		_ = _index204
		(_1001_v12).ArraySet1CodePoint(_1002_v13, (_index204).Int())
		var _1003_v14 D0
		_ = _1003_v14
		_1003_v14 = Companion_D0_.Create_DC1_(p1, p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_990_v2).Cardinality())), _dafny.ArrayCastTo((_988_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_988_v0), 0))).Int())))
		var _1004_v15 _dafny.Sequence
		_ = _1004_v15
		_1004_v15 = _dafny.SeqOf(_991_v3)
		var _1005_v16 _dafny.Set
		_ = _1005_v16
		_1005_v16 = _dafny.SetOf(_1004_v15)
		var _pat_let_tv30 = p2
		_ = _pat_let_tv30
		var _pat_let_tv31 = _1005_v16
		_ = _pat_let_tv31
		var _pat_let_tv32 = p1
		_ = _pat_let_tv32
		var _pat_let_tv33 = _989_v1
		_ = _pat_let_tv33
		r0 = func(_pat_let32_0 D0) D0 {
			return func(_1006_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let33_0 _dafny.Int) D0 {
					return func(_1007_dt__update_hcf3_h0 _dafny.Int) D0 {
						return func(_pat_let34_0 _dafny.Int) D0 {
							return func(_1008_dt__update_hcf1_h0 _dafny.Int) D0 {
								return func(_pat_let35_0 _dafny.Array) D0 {
									return func(_1009_dt__update_hcf4_h0 _dafny.Array) D0 {
										return Companion_D0_.Create_DC1_(_1008_dt__update_hcf1_h0, (_1006_dt__update__tmp_h0).Dtor_cf2(), _1007_dt__update_hcf3_h0, _1009_dt__update_hcf4_h0)
									}(_pat_let35_0)
								}(_pat_let_tv33)
							}(_pat_let34_0)
						}((func() _dafny.Int {
							if _pat_let_tv30 {
								return (_dafny.Zero).Minus((_pat_let_tv31).Cardinality())
							}
							return _pat_let_tv32
						})())
					}(_pat_let33_0)
				}(_dafny.IntOfInt64(-614))
			}(_pat_let32_0)
		}(_1003_v14)
		var _1010_v17 _dafny.MultiSet
		_ = _1010_v17
		_1010_v17 = _dafny.MultiSetOf(p1, p1, p1)
		var _1011_v18 _dafny.MultiSet
		_ = _1011_v18
		_1011_v18 = _dafny.MultiSetOf(_1010_v17)
		r1 = ((_1011_v18).Update(_1010_v17, Companion_Default___.Abs(p1))).IsDisjointFrom(_1011_v18)
		return r0, r1
	}
}
func (_this *C3) M4(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) (bool, _dafny.Int, T0) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 T0 = (T0)(nil)
		_ = r2
		var _1012_v0 bool
		_ = _1012_v0
		_1012_v0 = false
		var _1013_i0 _dafny.Int
		_ = _1013_i0
		_1013_i0 = _dafny.Zero
		{
			for _1012_v0 {
				{
					if (_1013_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_1013_i0 = (_1013_i0).Plus(_dafny.One)
					var _1014_v1 _dafny.Array
					_ = _1014_v1
					var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
					_ = _nw188
					_1014_v1 = _nw188
					var _1015_v2 _dafny.MultiSet
					_ = _1015_v2
					_1015_v2 = _dafny.MultiSetOf(_1014_v1, _1014_v1)
					var _1016_v3 _dafny.Sequence
					_ = _1016_v3
					_1016_v3 = _dafny.SeqOf(_1012_v0, !(_1015_v2).Equals(_1015_v2), _1012_v0, !((func() bool {
						if Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState) {
							return _1012_v0
						}
						return _1012_v0
					})()))
					_1016_v3 = _1016_v3
					var _1017_v4 _dafny.Sequence
					_ = _1017_v4
					_1017_v4 = _dafny.SeqOf(_1016_v3, _1016_v3, _1016_v3, Companion_Default___.Fm16(globalState))
					var _1018_v5 D1
					_ = _1018_v5
					_1018_v5 = Companion_D1_.Create_DC6_(p1, _1017_v4, p1, _1012_v0)
					var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1014_v1), 0))
					_ = _index205
					(_1014_v1).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), (_index205).Int())
					var _1019_v6 _dafny.Sequence
					_ = _1019_v6
					_1019_v6 = _dafny.SeqOf(p1)
					var _1020_v7 _dafny.Set
					_ = _1020_v7
					_1020_v7 = _dafny.SetOf(p1, p1)
					var _1021_v8 _dafny.CodePoint
					_ = _1021_v8
					_1021_v8 = _dafny.CodePoint('a')
					var _1022_v9 _dafny.Map
					_ = _1022_v9
					_1022_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm17((_dafny.Zero).Minus(p1), _1012_v0, _1021_v8, globalState), _1019_v6)
					var _1023_v10 D1
					_ = _1023_v10
					_1023_v10 = Companion_D1_.Create_DC5_(p1, !(_1012_v0), _1012_v0, _1022_v9)
					var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1014_v1), 0))
					_ = _index206
					var _rhs205 bool = _1012_v0
					_ = _rhs205
					var _rhs206 bool = !(_1020_v7).Contains((_1019_v6).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1019_v6).Cardinality()))).Uint32()).(_dafny.Int))
					_ = _rhs206
					var _rhs207 bool = Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState)
					_ = _rhs207
					var _rhs208 D1 = _1018_v5
					_ = _rhs208
					var _rhs209 _dafny.Int = (_1023_v10).Dtor_cf9()
					_ = _rhs209
					var _lhs188 *GlobalState = globalState
					_ = _lhs188
					var _lhs189 _dafny.Array = _1014_v1
					_ = _lhs189
					var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1014_v1), 0))
					_ = _lhs190
					_lhs188.F7 = _rhs205
					r0 = _rhs206
					_1012_v0 = _rhs207
					_1018_v5 = _rhs208
					(_lhs189).ArraySet1(_rhs209, (_lhs190).Int())
					var _1024_v11 _dafny.Array
					_ = _1024_v11
					var _nw189 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
					_ = _nw189
					_1024_v11 = _nw189
					var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1024_v11), 0))
					_ = _index207
					(_1024_v11).ArraySet1(true, (_index207).Int())
					var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_1024_v11), 0))
					_ = _index208
					(_1024_v11).ArraySet1(_1012_v0, (_index208).Int())
					var _1025_v12 D3
					_ = _1025_v12
					_1025_v12 = Companion_D3_.Create_DC11_((_1014_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1014_v1), 0))).Int()).(_dafny.Int))
					var _1026_v13 _dafny.Set
					_ = _1026_v13
					_1026_v13 = _dafny.SetOf(_1014_v1, _1014_v1)
					var _pat_let_tv34 = _1014_v1
					_ = _pat_let_tv34
					var _pat_let_tv35 = _1014_v1
					_ = _pat_let_tv35
					var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_1014_v1), 0))
					_ = _index209
					(_1014_v1).ArraySet1((_dafny.Zero).Minus(((func(_pat_let36_0 D3) D3 {
						return func(_1027_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let37_0 _dafny.Int) D3 {
								return func(_1028_dt__update_hcf22_h0 _dafny.Int) D3 {
									return Companion_D3_.Create_DC11_(_1028_dt__update_hcf22_h0)
								}(_pat_let37_0)
							}((_dafny.Zero).Minus((_pat_let_tv35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_pat_let_tv34), 0))).Int()).(_dafny.Int)))
						}(_pat_let36_0)
					}(_1025_v12)).Dtor_cf22()).Plus((p1).Minus((_1026_v13).Cardinality()))), (_index209).Int())
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _hi3 _dafny.Int = p1
		_ = _hi3
		for _1029_i1 := p1; _1029_i1.Cmp(_hi3) < 0; _1029_i1 = _1029_i1.Plus(_dafny.One) {
			var _1030_v14 _dafny.Map
			_ = _1030_v14
			_1030_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1012_v0)
			var _1031_v15 _dafny.MultiSet
			_ = _1031_v15
			_1031_v15 = _dafny.MultiSetOf((_dafny.Zero).Minus((_1030_v14).Cardinality()))
			var _1032_v16 _dafny.Sequence
			_ = _1032_v16
			_1032_v16 = _dafny.SeqOf(_1012_v0, _1012_v0)
			var _1033_v17 _dafny.Set
			_ = _1033_v17
			_1033_v17 = _dafny.SetOf(_1012_v0)
			var _1034_v18 _dafny.Array
			_ = _1034_v18
			var _nwElement0_51 bool = (_1031_v15).IsDisjointFrom(_dafny.MultiSetOf(Companion_Default___.Fm0(p1, globalState), _1029_i1))
			_ = _nwElement0_51
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(28))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_51, 0)
			(_nw190).ArraySet1(_1012_v0, 1)
			(_nw190).ArraySet1((Companion_Default___.Fm0(p1, globalState)).Cmp(_1029_i1) < 0, 2)
			(_nw190).ArraySet1(_1012_v0, 3)
			(_nw190).ArraySet1(_1012_v0, 4)
			(_nw190).ArraySet1(_1012_v0, 5)
			(_nw190).ArraySet1(_1012_v0, 6)
			(_nw190).ArraySet1(_1012_v0, 7)
			(_nw190).ArraySet1(_1012_v0, 8)
			(_nw190).ArraySet1((_1029_i1).Cmp(p1) <= 0, 9)
			(_nw190).ArraySet1(_1012_v0, 10)
			(_nw190).ArraySet1((p1).Cmp(_1029_i1) >= 0, 11)
			(_nw190).ArraySet1(!(_1012_v0), 12)
			(_nw190).ArraySet1(_dafny.Companion_Sequence_.Contains(_1032_v16, _1012_v0), 13)
			(_nw190).ArraySet1(_1012_v0, 14)
			(_nw190).ArraySet1(Companion_Default___.Fm12(!(_1012_v0), _1012_v0, globalState), 15)
			(_nw190).ArraySet1(_1012_v0, 16)
			(_nw190).ArraySet1(_1012_v0, 17)
			(_nw190).ArraySet1(!(_1012_v0) || (_1012_v0), 18)
			(_nw190).ArraySet1(_1012_v0, 19)
			(_nw190).ArraySet1(_1012_v0, 20)
			(_nw190).ArraySet1(_1012_v0, 21)
			(_nw190).ArraySet1(_1012_v0, 22)
			(_nw190).ArraySet1((_1012_v0) == (_1012_v0), 23)
			(_nw190).ArraySet1(!(true), 24)
			(_nw190).ArraySet1((!(_1012_v0)) || (_1012_v0), 25)
			(_nw190).ArraySet1(_1012_v0, 26)
			(_nw190).ArraySet1((_1033_v17).IsSubsetOf(_1033_v17), 27)
			_1034_v18 = _nw190
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1034_v18), 0))
			_ = _index210
			(_1034_v18).ArraySet1(false, (_index210).Int())
			var _1035_v19 _dafny.Set
			_ = _1035_v19
			_1035_v19 = _dafny.SetOf(p1)
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1034_v18), 0))
			_ = _index211
			var _rhs210 bool = Companion_Default___.Fm12(_1012_v0, (_1029_i1).Cmp((_1035_v19).Cardinality()) < 0, globalState)
			_ = _rhs210
			var _rhs211 bool = false
			_ = _rhs211
			var _rhs212 _dafny.Int = p1
			_ = _rhs212
			var _lhs191 _dafny.Array = _1034_v18
			_ = _lhs191
			var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1034_v18), 0))
			_ = _lhs192
			var _lhs193 *GlobalState = globalState
			_ = _lhs193
			r0 = _rhs210
			(_lhs191).ArraySet1(_rhs211, (_lhs192).Int())
			_lhs193.F10 = _rhs212
			var _1036_v20 _dafny.Array
			_ = _1036_v20
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_21
			var _nw191 _dafny.Array
			_ = _nw191
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw191 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Map = (func(_1037_v14 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1038_i2 _dafny.Int) _dafny.Map {
						return _1037_v14
					}
				})(_1030_v14)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw191 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw191).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw191).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_1036_v20 = _nw191
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1036_v20), 0))
			_ = _index212
			(_1036_v20).ArraySet1(Companion_Default___.Fm18(p1, _1012_v0, globalState), (_index212).Int())
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(420), _dafny.ArrayLen((_1036_v20), 0))
			_ = _index213
			(_1036_v20).ArraySet1(_1030_v14, (_index213).Int())
			var _1039_v21 _dafny.Array
			_ = _1039_v21
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_22
			var _nw192 _dafny.Array
			_ = _nw192
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw192 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Sequence = (func(_1040_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1041_i3 _dafny.Int) _dafny.Sequence {
						return _1040_p0
					}
				})(p0)
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
			_1039_v21 = _nw192
			var _1042_v22 _dafny.Map
			_ = _1042_v22
			_1042_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-100), p0)
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1039_v21), 0))
			_ = _index214
			(_1039_v21).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1042_v22).Contains(_1029_i1) {
					return (_1042_v22).Get(_1029_i1).(_dafny.Sequence)
				}
				return p0
			})(), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1042_v22).Contains(_1029_i1) {
					return (_1042_v22).Get(_1029_i1).(_dafny.Sequence)
				}
				return p0
			})()).Cardinality()))).Uint32(), Companion_Default___.Fm11(p1, globalState)), (_index214).Int())
			var _1043_v23 _dafny.CodePoint
			_ = _1043_v23
			_1043_v23 = _dafny.CodePoint('t')
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1039_v21), 0))
			_ = _index215
			(_1039_v21).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1043_v23), p0), (_index215).Int())
			var _1044_v24 *C0
			_ = _1044_v24
			var _nw193 *C0 = New_C0_()
			_ = _nw193
			_nw193.Ctor__(Companion_Default___.Fm12((_1034_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1034_v18), 0))).Int()).(bool), (_1034_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(212), _dafny.ArrayLen((_1034_v18), 0))).Int()).(bool), globalState))
			_1044_v24 = _nw193
		}
		var _hi4 _dafny.Int = p1
		_ = _hi4
		for _1045_i4 := _dafny.IntOfUint32((p0).Cardinality()); _1045_i4.Cmp(_hi4) < 0; _1045_i4 = _1045_i4.Plus(_dafny.One) {
			var _1046_v25 _dafny.Map
			_ = _1046_v25
			_1046_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, _1012_v0)
			_1046_v25 = (_1046_v25).Update(_1012_v0, false)
			var _1047_v26 _dafny.CodePoint
			_ = _1047_v26
			_1047_v26 = _dafny.CodePoint('q')
			var _1048_v27 _dafny.MultiSet
			_ = _1048_v27
			_1048_v27 = _dafny.MultiSetOf(_1047_v26, _1047_v26, _1047_v26)
			r1 = (func() _dafny.Int {
				if _1012_v0 {
					return (_1048_v27).Cardinality()
				}
				return _dafny.IntOfInt64(331)
			})()
			var _1049_v28 _dafny.Array
			_ = _1049_v28
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw194
			_1049_v28 = _nw194
			var _1050_v29 _dafny.Map
			_ = _1050_v29
			_1050_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1049_v28, true)
			var _1051_v30 _dafny.MultiSet
			_ = _1051_v30
			_1051_v30 = _dafny.MultiSetOf(_1012_v0)
			_1050_v29 = (_1050_v29).Update(_1049_v28, (_1051_v30).IsDisjointFrom(_1051_v30))
			var _1052_v31 _dafny.Sequence
			_ = _1052_v31
			_1052_v31 = _dafny.SeqOf(_1012_v0)
			if (_1052_v31).Select((Companion_Default___.SafeIndex(_1045_i4, _dafny.IntOfUint32((_1052_v31).Cardinality()))).Uint32()).(bool) {
				(globalState).F10 = (p1).Plus(Companion_Default___.Fm0(_1045_i4, globalState))
				var _1053_v32 _dafny.Array
				_ = _1053_v32
				var _nw195 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
				_ = _nw195
				_1053_v32 = _nw195
				var _1054_v33 _dafny.Map
				_ = _1054_v33
				_1054_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1053_v32, (p1).Times(_1045_i4))
				(globalState).F10 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_1054_v33).Contains(_1053_v32) {
						return (_1054_v33).Get(_1053_v32).(_dafny.Int)
					}
					return (_1045_i4).Plus(_1045_i4)
				})())
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1049_v28), 0))
				_ = _index216
				(_1049_v28).ArraySet1(p1, (_index216).Int())
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1049_v28), 0))
				_ = _index217
				(_1049_v28).ArraySet1(p1, (_index217).Int())
				(globalState).F22 = !(Companion_Default___.Fm12(!(_1012_v0), _1012_v0, globalState))
				var _1055_v34 _dafny.Array
				_ = _1055_v34
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_23
				var _nw196 _dafny.Array
				_ = _nw196
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw196 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Map = (func(_1056_v25 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_1057_i5 _dafny.Int) _dafny.Map {
							return _1056_v25
						}
					})(_1046_v25)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw196 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw196).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw196).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1055_v34 = _nw196
				var _1058_v35 _dafny.Sequence
				_ = _1058_v35
				_1058_v35 = _dafny.SeqOf(_1046_v25, _1046_v25, _1046_v25, _1046_v25, _1046_v25)
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1055_v34), 0))
				_ = _index218
				(_1055_v34).ArraySet1((_1058_v35).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1058_v35).Cardinality()))).Uint32()).(_dafny.Map), (_index218).Int())
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1055_v34), 0))
				_ = _index219
				(_1055_v34).ArraySet1((func() _dafny.Map {
					if Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState) {
						return _1046_v25
					}
					return _1046_v25
				})(), (_index219).Int())
			} else {
				var _1059_v36 _dafny.Array
				_ = _1059_v36
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_24
				var _nw197 _dafny.Array
				_ = _nw197
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw197 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) bool = (func(_1060_v0 bool) func(_dafny.Int) bool {
						return func(_1061_i6 _dafny.Int) bool {
							return _1060_v0
						}
					})(_1012_v0)
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw197 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw197).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw197).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1059_v36 = _nw197
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1059_v36), 0))
				_ = _index220
				(_1059_v36).ArraySet1(_1012_v0, (_index220).Int())
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1059_v36), 0))
				_ = _index221
				(_1059_v36).ArraySet1(true, (_index221).Int())
				_1047_v26 = _1047_v26
				var _1062_v37 _dafny.Map
				_ = _1062_v37
				_1062_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1047_v26, _1012_v0)
				var _1063_v38 _dafny.Map
				_ = _1063_v38
				_1063_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1045_i4, true)
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1059_v36), 0))
				_ = _index222
				(_1059_v36).ArraySet1((Companion_Default___.Fm12((_1059_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1059_v36), 0))).Int()).(bool), (func() bool {
					if (_1062_v37).Contains(_1047_v26) {
						return (_1062_v37).Get(_1047_v26).(bool)
					}
					return (func() bool {
						if (_1063_v38).Contains(p1) {
							return (_1063_v38).Get(p1).(bool)
						}
						return true
					})()
				})(), globalState)) || (_1012_v0), (_index222).Int())
				_1012_v0 = true
				var _1064_v39 _dafny.Array
				_ = _1064_v39
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
				_ = _nw198
				_1064_v39 = _nw198
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1064_v39), 0))
				_ = _index223
				(_1064_v39).ArraySet1CodePoint(_1047_v26, (_index223).Int())
				var _1065_v40 D4
				_ = _1065_v40
				_1065_v40 = Companion_D4_.Create_DC13_((p0).Select((Companion_Default___.SafeIndex(_1045_i4, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1064_v39), 0))
				_ = _index224
				(_1064_v39).ArraySet1CodePoint((_1065_v40).Dtor_cf24(), (_index224).Int())
			}
		}
		(globalState).F13 = !(_1012_v0)
		var _1066_v41 _dafny.Sequence
		_ = _1066_v41
		_1066_v41 = _dafny.SeqOf(false)
		if ((_1066_v41).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1066_v41).Cardinality()))).Uint32()).(bool)) == (Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState)) {
			var _1067_v42 D0
			_ = _1067_v42
			_1067_v42 = Companion_D0_.Create_DC2_(p1, _1012_v0)
			var _1068_v43 _dafny.MultiSet
			_ = _1068_v43
			_1068_v43 = _dafny.MultiSetOf(_1067_v42)
			(globalState).F22 = (_1068_v43).IsProperSubsetOf(_1068_v43)
			(globalState).F10 = (func() _dafny.Int {
				if _1012_v0 {
					return p1
				}
				return p1
			})()
			var _1069_v44 D3
			_ = _1069_v44
			_1069_v44 = Companion_D3_.Create_DC11_((p1).Plus(p1))
			var _source17 D3 = _1069_v44
			_ = _source17
			if _source17.Is_DC11() {
				var _1070___mcc_h0 _dafny.Int = _source17.Get_().(D3_DC11).Cf22
				_ = _1070___mcc_h0
				var _1071_cf22 _dafny.Int = _1070___mcc_h0
				_ = _1071_cf22
				var _1072_v45 _dafny.Map
				_ = _1072_v45
				_1072_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm11(_1071_cf22, globalState), _1066_v41)
				_1072_v45 = _1072_v45
				var _1073_v46 _dafny.Map
				_ = _1073_v46
				_1073_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
				var _1074_v47 _dafny.Map
				_ = _1074_v47
				_1074_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_1075_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1076_i7 _dafny.Int) _dafny.Int {
						return _1075_p1
					}
				})(p1))), Companion_Default___.Fm14(_1012_v0, _1071_cf22, Companion_Default___.Fm0(_1071_cf22, globalState), globalState)), _1073_v46)
				var _1077_v48 _dafny.Sequence
				_ = _1077_v48
				_1077_v48 = _dafny.SeqOf(p1)
				_1074_v47 = (_1074_v47).Update(_1077_v48, _1073_v46)
				var _1078_v49 *C0
				_ = _1078_v49
				var _nw199 *C0 = New_C0_()
				_ = _nw199
				_nw199.Ctor__(_1012_v0)
				_1078_v49 = _nw199
				var _1079_v50 _dafny.Map
				_ = _1079_v50
				_1079_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if _1012_v0 {
						return Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState)
					}
					return _1012_v0
				})(), (func() _dafny.Int {
					if _1012_v0 {
						return _dafny.IntOfUint32((_dafny.SeqOf(_1078_v49, _1078_v49)).Cardinality())
					}
					return (_this).Fm8(p1, _1071_cf22, p1, globalState)
				})())
				_1079_v50 = (_1079_v50).Update(!(_1012_v0), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality()))
				var _1080_v51 _dafny.Map
				_ = _1080_v51
				_1080_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, (_1078_v49).F32())
				var _1081_v52 _dafny.Map
				_ = _1081_v52
				_1081_v52 = _1080_v51
				var _1082_v53 _dafny.MultiSet
				_ = _1082_v53
				_1082_v53 = _dafny.MultiSetOf(_1071_cf22, _1071_cf22)
				var _1083_v54 _dafny.Array
				_ = _1083_v54
				var _nwElement0_52 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1012_v0)
				_ = _nwElement0_52
				var _nw200 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(14))
				_ = _nw200
				(_nw200).ArraySet1(_nwElement0_52, 0)
				(_nw200).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, (_1078_v49).F32()), 1)
				(_nw200).ArraySet1(_1080_v51, 2)
				(_nw200).ArraySet1(_1080_v51, 3)
				(_nw200).ArraySet1(_1080_v51, 4)
				(_nw200).ArraySet1(_1080_v51, 5)
				(_nw200).ArraySet1(_1080_v51, 6)
				(_nw200).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1078_v49).F32(), (_1078_v49).F32())).Merge(_1080_v51), 7)
				(_nw200).ArraySet1(_1080_v51, 8)
				(_nw200).ArraySet1(_1080_v51, 9)
				(_nw200).ArraySet1((_1081_v52), 10)
				(_nw200).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1078_v49).F32(), (_1078_v49).F32())).Update((_1078_v49).F32(), _1012_v0), 11)
				(_nw200).ArraySet1(Companion_Default___.Fm19(_1082_v53, p1, globalState), 12)
				(_nw200).ArraySet1((_1080_v51).Merge(_1080_v51), 13)
				_1083_v54 = _nw200
				_1083_v54 = _1083_v54
			} else if _source17.Is_DC12() {
				var _1084___mcc_h1 bool = _source17.Get_().(D3_DC12).Cf23
				_ = _1084___mcc_h1
				var _1085_cf23 bool = _1084___mcc_h1
				_ = _1085_cf23
				var _1086_v55 _dafny.Array
				_ = _1086_v55
				var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw201
				_1086_v55 = _nw201
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1086_v55), 0))
				_ = _index225
				(_1086_v55).ArraySet1(p1, (_index225).Int())
				var _1087_v56 _dafny.MultiSet
				_ = _1087_v56
				_1087_v56 = _dafny.MultiSetOf(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(702))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_1088_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				})))
				var _1089_v57 *C0
				_ = _1089_v57
				var _nw202 *C0 = New_C0_()
				_ = _nw202
				_nw202.Ctor__(false)
				_1089_v57 = _nw202
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1086_v55), 0))
				_ = _index226
				var _rhs213 bool = (_1066_v41).Select((Companion_Default___.SafeIndex(((_1087_v56).Union(_1087_v56)).Cardinality(), _dafny.IntOfUint32((_1066_v41).Cardinality()))).Uint32()).(bool)
				_ = _rhs213
				var _rhs214 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqOf(_1089_v57)).Cardinality())).Plus(p1)
				_ = _rhs214
				var _rhs215 _dafny.Int = (_dafny.Zero).Minus(p1)
				_ = _rhs215
				var _rhs216 _dafny.Int = p1
				_ = _rhs216
				var _rhs217 _dafny.Int = p1
				_ = _rhs217
				var _lhs194 *GlobalState = globalState
				_ = _lhs194
				var _lhs195 *GlobalState = globalState
				_ = _lhs195
				var _lhs196 *GlobalState = globalState
				_ = _lhs196
				var _lhs197 _dafny.Array = _1086_v55
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1086_v55), 0))
				_ = _lhs198
				_lhs194.F13 = _rhs213
				_lhs195.F10 = _rhs214
				r1 = _rhs215
				_lhs196.F10 = _rhs216
				(_lhs197).ArraySet1(_rhs217, (_lhs198).Int())
				(globalState).F10 = p1
				var _1090_v58 _dafny.Map
				_ = _1090_v58
				_1090_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1086_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1086_v55), 0))).Int()).(_dafny.Int))
				var _1091_v59 _dafny.Map
				_ = _1091_v59
				_1091_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (func() _dafny.Int {
					if (_1090_v58).Contains(true) {
						return (_1090_v58).Get(true).(_dafny.Int)
					}
					return _dafny.IntOfInt64(891)
				})())
				_1091_v59 = (_1091_v59).Update(_dafny.IntOfUint32((p0).Cardinality()), (_1086_v55).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1086_v55), 0))).Int()).(_dafny.Int))
				(globalState).F7 = _1012_v0
			} else {
				var _1092___mcc_h2 _dafny.MultiSet = _source17.Get_().(D3_DC10).Cf21
				_ = _1092___mcc_h2
				var _1093_cf21 _dafny.MultiSet = _1092___mcc_h2
				_ = _1093_cf21
				var _1094_v60 _dafny.Set
				_ = _1094_v60
				_1094_v60 = _dafny.SetOf(_dafny.IntOfUint32((p0).Cardinality()))
				var _1095_v61 _dafny.Sequence
				_ = _1095_v61
				_1095_v61 = _dafny.SeqOf((_dafny.Zero).Minus((_this).Fm8(p1, _dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfInt64(790), globalState)), _dafny.IntOfInt64(100), (_1094_v60).Cardinality())
				var _1096_v62 _dafny.Map
				_ = _1096_v62
				_1096_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1095_v61), _1095_v61)
				var _1097_v63 D1
				_ = _1097_v63
				_1097_v63 = Companion_D1_.Create_DC5_(p1, _1012_v0, _1012_v0, Companion_Default___.Fm20(p1, _1012_v0, (_dafny.Zero).Minus(p1), globalState))
				var _1098_v64 _dafny.Sequence
				_ = _1098_v64
				_1098_v64 = _dafny.SeqOf(Companion_D1_.Create_DC5_(_dafny.IntOfUint32((p0).Cardinality()), _1012_v0, _1012_v0, _1096_v62), _1097_v63, Companion_D1_.Create_DC5_(p1, !(_1012_v0), !(_1012_v0), _1096_v62), _1097_v63)
				_1098_v64 = _1098_v64
				var _1099_v65 _dafny.Array
				_ = _1099_v65
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(25))
				_ = _nw203
				_1099_v65 = _nw203
				var _1100_v66 _dafny.Map
				_ = _1100_v66
				_1100_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(p1))
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_1099_v65), 0))
				_ = _index227
				(_1099_v65).ArraySet1((_1094_v60).Union(func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter34 := _dafny.Iterate((_1100_v66).Keys().Elements()); ; {
						_compr_30, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _1101_v67 _dafny.Int
						_1101_v67 = interface{}(_compr_30).(_dafny.Int)
						if (_1100_v66).Contains(_1101_v67) {
							_coll30.Add((_1101_v67).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mbi")).Cardinality())))
						}
					}
					return _coll30.ToSet()
				}()), (_index227).Int())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_1099_v65), 0))
				_ = _index228
				(_1099_v65).ArraySet1(_dafny.SetOf(p1), (_index228).Int())
				var _1102_v68 _dafny.CodePoint
				_ = _1102_v68
				_1102_v68 = _dafny.CodePoint('d')
				var _1103_v69 D4
				_ = _1103_v69
				_1103_v69 = Companion_D4_.Create_DC13_(_1102_v68)
				var _1104_v70 _dafny.Sequence
				_ = _1104_v70
				_1104_v70 = _dafny.SeqOf(_1103_v69)
				r1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1104_v70, _1104_v70), _dafny.SeqOf(Companion_D4_.Create_DC13_(_1102_v68)))).Cardinality())
				var _1105_v71 _dafny.Set
				_ = _1105_v71
				_1105_v71 = _dafny.SetOf(_1102_v68, _1102_v68)
				r0 = (_1105_v71).IsSubsetOf(_1105_v71)
			}
			(globalState).F22 = (_1012_v0) || (_1012_v0)
			(globalState).F10 = p1
		} else {
			(globalState).F7 = _1012_v0
			var _1106_v72 _dafny.Array
			_ = _1106_v72
			var _nwElement0_53 bool = _1012_v0
			_ = _nwElement0_53
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(8))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_53, 0)
			(_nw204).ArraySet1(Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState), 1)
			(_nw204).ArraySet1(_1012_v0, 2)
			(_nw204).ArraySet1((_1066_v41).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1066_v41).Cardinality()))).Uint32()).(bool), 3)
			(_nw204).ArraySet1((func() bool {
				if (_1066_v41).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(p1, globalState), _dafny.IntOfUint32((_1066_v41).Cardinality()))).Uint32()).(bool) {
					return _1012_v0
				}
				return _1012_v0
			})(), 4)
			(_nw204).ArraySet1(_1012_v0, 5)
			(_nw204).ArraySet1((!(_1012_v0)) && (_1012_v0), 6)
			(_nw204).ArraySet1(_1012_v0, 7)
			_1106_v72 = _nw204
			var _1107_v73 _dafny.Map
			_ = _1107_v73
			_1107_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, _1012_v0)
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1106_v72), 0))
			_ = _index229
			(_1106_v72).ArraySet1((_1107_v73).Equals(_1107_v73), (_index229).Int())
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1106_v72), 0))
			_ = _index230
			(_1106_v72).ArraySet1(_1012_v0, (_index230).Int())
			var _1108_v74 _dafny.Sequence
			_ = _1108_v74
			_1108_v74 = _dafny.UnicodeSeqOfUtf8Bytes("rlvkgkw")
			var _rhs218 _dafny.Sequence = _1108_v74
			_ = _rhs218
			var _rhs219 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("opjwmmjub"))
			_ = _rhs219
			_1108_v74 = _rhs218
			_1108_v74 = _rhs219
			var _1109_v75 *C0
			_ = _1109_v75
			var _nw205 *C0 = New_C0_()
			_ = _nw205
			_nw205.Ctor__((_1106_v72).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(557), _dafny.ArrayLen((_1106_v72), 0))).Int()).(bool))
			_1109_v75 = _nw205
			var _1110_v76 T0
			_ = _1110_v76
			var _nw206 *C1 = New_C1_()
			_ = _nw206
			_nw206.Ctor__()
			_1110_v76 = _nw206
			var _1111_v77 _dafny.CodePoint
			_ = _1111_v77
			_1111_v77 = _dafny.CodePoint('h')
			var _1112_v78 _dafny.Map
			_ = _1112_v78
			_1112_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1110_v76, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm10(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1), _1111_v77, (_1109_v75).F32(), globalState), p0)).Cardinality()))
			_1112_v78 = (_1112_v78).Update(_1110_v76, p1)
		}
		var _1113_v79 _dafny.MultiSet
		_ = _1113_v79
		_1113_v79 = _dafny.MultiSetOf(p1)
		var _1114_v80 _dafny.Map
		_ = _1114_v80
		_1114_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1066_v41).Cardinality()), _1012_v0)
		var _1115_v81 _dafny.Map
		_ = _1115_v81
		_1115_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1113_v79).Update((_dafny.Zero).Minus((_1114_v80).Cardinality()), Companion_Default___.Abs(_dafny.IntOfInt64(-908))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(687))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}((func(_1116_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1117_i9 _dafny.Int) _dafny.Int {
				return _1116_p1
			}
		})(p1))))
		var _1118_v82 D1
		_ = _1118_v82
		_1118_v82 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_1066_v41).Cardinality()), (p1).Cmp(p1) <= 0, _1012_v0, _1115_v81)
		var _source18 D1 = _1118_v82
		_ = _source18
		if _source18.Is_DC5() {
			var _1119___mcc_h3 _dafny.Int = _source18.Get_().(D1_DC5).Cf9
			_ = _1119___mcc_h3
			var _1120___mcc_h4 bool = _source18.Get_().(D1_DC5).Cf10
			_ = _1120___mcc_h4
			var _1121___mcc_h5 bool = _source18.Get_().(D1_DC5).Cf11
			_ = _1121___mcc_h5
			var _1122___mcc_h6 _dafny.Map = _source18.Get_().(D1_DC5).Cf12
			_ = _1122___mcc_h6
			var _1123_cf12 _dafny.Map = _1122___mcc_h6
			_ = _1123_cf12
			var _1124_cf11 bool = _1121___mcc_h5
			_ = _1124_cf11
			var _1125_cf10 bool = _1120___mcc_h4
			_ = _1125_cf10
			var _1126_cf9 _dafny.Int = _1119___mcc_h3
			_ = _1126_cf9
			if _1125_cf10 {
				var _1127_v83 _dafny.MultiSet
				_ = _1127_v83
				_1127_v83 = _dafny.MultiSetOf(_1124_cf11, _1125_cf10)
				var _1128_v84 _dafny.Sequence
				_ = _1128_v84
				_1128_v84 = _dafny.SeqOf(_1127_v83, _1127_v83)
				var _1129_v85 _dafny.Array
				_ = _1129_v85
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw207
				_1129_v85 = _nw207
				var _rhs220 _dafny.Int = p1
				_ = _rhs220
				var _rhs221 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1128_v84, _1128_v84), (Companion_Default___.SafeIndex((_dafny.MultiSetOf(_1129_v85, _1129_v85, _1129_v85)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1128_v84, _1128_v84)).Cardinality()))).Uint32(), _dafny.MultiSetOf(_1012_v0, _1012_v0))
				_ = _rhs221
				_1126_cf9 = _rhs220
				_1128_v84 = _rhs221
				var _1130_v86 _dafny.Map
				_ = _1130_v86
				_1130_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, _1124_cf11)
				var _1131_v87 _dafny.CodePoint
				_ = _1131_v87
				_1131_v87 = _dafny.CodePoint('p')
				var _1132_v88 _dafny.Sequence
				_ = _1132_v88
				_1132_v88 = _dafny.SeqOf((_1130_v86).Cardinality(), _1126_cf9, (_dafny.SetOf(_1131_v87)).Cardinality(), _1126_cf9)
				var _1133_v89 _dafny.Sequence
				_ = _1133_v89
				_1133_v89 = _dafny.SeqOf(_1132_v88, _1132_v88)
				var _1134_v90 _dafny.Map
				_ = _1134_v90
				_1134_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
					if _1124_cf11 {
						return (_1133_v89).Select((Companion_Default___.SafeIndex(_1126_cf9, _dafny.IntOfUint32((_1133_v89).Cardinality()))).Uint32()).(_dafny.Sequence)
					}
					return _1132_v88
				})(), _1124_cf11)
				_1134_v90 = _1134_v90
				var _1135_v91 _dafny.MultiSet
				_ = _1135_v91
				_1135_v91 = _dafny.MultiSetOf(_1131_v87, _1131_v87)
				var _1136_v92 _dafny.Set
				_ = _1136_v92
				_1136_v92 = _dafny.SetOf(_1126_cf9, (_1135_v91).Cardinality())
				var _1137_v93 D0
				_ = _1137_v93
				_1137_v93 = Companion_D0_.Create_DC2_(p1, false)
				var _1138_v94 _dafny.Map
				_ = _1138_v94
				_1138_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1125_cf10), p1)
				var _1139_v95 _dafny.Array
				_ = _1139_v95
				var _nwElement0_54 bool = _1125_cf10
				_ = _nwElement0_54
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(22))
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_54, 0)
				(_nw208).ArraySet1(_1124_cf11, 1)
				(_nw208).ArraySet1((_dafny.MultiSetOf(Companion_Default___.Fm11(_dafny.IntOfUint32((_1066_v41).Cardinality()), globalState))).IsProperSubsetOf(_dafny.MultiSetOf(_1131_v87)), 2)
				(_nw208).ArraySet1(true, 3)
				(_nw208).ArraySet1(_1125_cf10, 4)
				(_nw208).ArraySet1(((_1136_v92).Cardinality()).Cmp(p1) < 0, 5)
				(_nw208).ArraySet1(_1125_cf10, 6)
				(_nw208).ArraySet1((func() bool {
					if (_1130_v86).Contains(false) {
						return (_1130_v86).Get(false).(bool)
					}
					return _1125_cf10
				})(), 7)
				(_nw208).ArraySet1(_1124_cf11, 8)
				(_nw208).ArraySet1((_1137_v93).Dtor_cf6(), 9)
				(_nw208).ArraySet1(_1012_v0, 10)
				(_nw208).ArraySet1(Companion_Default___.Fm12(_1012_v0, _1012_v0, globalState), 11)
				(_nw208).ArraySet1(Companion_Default___.Fm12(_1012_v0, false, globalState), 12)
				(_nw208).ArraySet1(Companion_Default___.Fm12(_1124_cf11, !(_1124_cf11), globalState), 13)
				(_nw208).ArraySet1(_1124_cf11, 14)
				(_nw208).ArraySet1(_1012_v0, 15)
				(_nw208).ArraySet1((_1138_v94).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, _1126_cf9)), 16)
				(_nw208).ArraySet1(_1012_v0, 17)
				(_nw208).ArraySet1(!(_1012_v0), 18)
				(_nw208).ArraySet1(_1012_v0, 19)
				(_nw208).ArraySet1(!(_1113_v79).Equals(_1113_v79), 20)
				(_nw208).ArraySet1(_1124_cf11, 21)
				_1139_v95 = _nw208
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1139_v95), 0))
				_ = _index231
				(_1139_v95).ArraySet1(_1125_cf10, (_index231).Int())
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((_1139_v95), 0))
				_ = _index232
				(_1139_v95).ArraySet1(!((_1127_v83).IsDisjointFrom(_1127_v83)), (_index232).Int())
				var _1140_v96 _dafny.Set
				_ = _1140_v96
				_1140_v96 = _dafny.SetOf(_1012_v0, _1012_v0)
				var _1141_v97 _dafny.Map
				_ = _1141_v97
				_1141_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1139_v95, _1140_v96)
				var _1142_v98 D8
				_ = _1142_v98
				_1142_v98 = Companion_D8_.Create_DC21_(_1066_v41)
				var _1143_v99 _dafny.Map
				_ = _1143_v99
				_1143_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1141_v97, _dafny.Companion_Sequence_.Concatenate((_1142_v98).Dtor_cf35(), _1066_v41))
				_1143_v99 = (_1143_v99).Update(_1141_v97, _1066_v41)
				(globalState).F24 = (_1136_v92).Intersection(_1136_v92)
			} else {
				var _1144_v100 T1
				_ = _1144_v100
				var _nw209 *C2 = New_C2_()
				_ = _nw209
				_nw209.Ctor__(_dafny.Companion_Sequence_.Concatenate(p0, p0))
				_1144_v100 = _nw209
				var _1145_v101 _dafny.Array
				_ = _1145_v101
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(3))
				_ = _nw210
				_1145_v101 = _nw210
				var _1146_v102 _dafny.CodePoint
				_ = _1146_v102
				_1146_v102 = _dafny.CodePoint('i')
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_1145_v101), 0))
				_ = _index233
				(_1145_v101).ArraySet1CodePoint(_1146_v102, (_index233).Int())
				var _1147_v103 _dafny.Array
				_ = _1147_v103
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
				_ = _nw211
				_1147_v103 = _nw211
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_1145_v101), 0))
				_ = _index234
				var _rhs222 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p0).Cardinality()), p1)
				_ = _rhs222
				var _rhs223 T1 = _1144_v100
				_ = _rhs223
				var _rhs224 _dafny.CodePoint = _1146_v102
				_ = _rhs224
				var _rhs225 _dafny.Array = _1147_v103
				_ = _rhs225
				var _lhs199 *GlobalState = globalState
				_ = _lhs199
				var _lhs200 _dafny.Array = _1145_v101
				_ = _lhs200
				var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_1145_v101), 0))
				_ = _lhs201
				_lhs199.F10 = _rhs222
				_1144_v100 = _rhs223
				(_lhs200).ArraySet1CodePoint(_rhs224, (_lhs201).Int())
				_1147_v103 = _rhs225
				var _1148_v104 _dafny.Sequence
				_ = _1148_v104
				_1148_v104 = _dafny.SeqOf(_1126_cf9, p1)
				(globalState).F10 = ((_dafny.Zero).Minus((_1148_v104).Select((Companion_Default___.SafeIndex(_1126_cf9, _dafny.IntOfUint32((_1148_v104).Cardinality()))).Uint32()).(_dafny.Int))).Times(_1126_cf9)
				r1 = (_dafny.Zero).Minus(p1)
				(globalState).F10 = _1126_cf9
				var _1149_v105 _dafny.Map
				_ = _1149_v105
				_1149_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1124_cf11, _1146_v102)
				var _1150_v106 _dafny.Array
				_ = _1150_v106
				var _nwElement0_55 _dafny.Map = (_1149_v105).Merge(_1149_v105)
				_ = _nwElement0_55
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(3))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_55, 0)
				(_nw212).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1124_cf11, (_1145_v101).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_1145_v101), 0))).Int())), 1)
				(_nw212).ArraySet1(_1149_v105, 2)
				_1150_v106 = _nw212
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_1150_v106), 0))
				_ = _index235
				(_1150_v106).ArraySet1(_1149_v105, (_index235).Int())
				var _1151_v107 _dafny.Map
				_ = _1151_v107
				_1151_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1125_cf10)
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(669), _dafny.ArrayLen((_1150_v106), 0))
				_ = _index236
				(_1150_v106).ArraySet1((_1149_v105).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if (_1151_v107).Contains(_1012_v0) {
						return (_1151_v107).Get(_1012_v0).(bool)
					}
					return _1125_cf10
				})(), _1146_v102)), (_index236).Int())
			}
			var _1152_v108 _dafny.Array
			_ = _1152_v108
			var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw213
			_1152_v108 = _nw213
			var _rhs226 _dafny.Array = _1152_v108
			_ = _rhs226
			var _rhs227 bool = (_1126_cf9).Cmp(p1) >= 0
			_ = _rhs227
			var _rhs228 _dafny.Int = p1
			_ = _rhs228
			var _rhs229 bool = _1125_cf10
			_ = _rhs229
			var _rhs230 _dafny.Int = _1126_cf9
			_ = _rhs230
			var _lhs202 *GlobalState = globalState
			_ = _lhs202
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 *GlobalState = globalState
			_ = _lhs204
			var _lhs205 *GlobalState = globalState
			_ = _lhs205
			_1152_v108 = _rhs226
			_lhs202.F13 = _rhs227
			_lhs203.F10 = _rhs228
			_lhs204.F13 = _rhs229
			_lhs205.F10 = _rhs230
			_1114_v80 = (_1114_v80).Update(_dafny.IntOfInt64(235), (_1124_cf11) == (_1125_cf10))
			var _1153_v109 _dafny.Array
			_ = _1153_v109
			_1153_v109 = _1152_v108
			var _source19 _dafny.Array = _1152_v108
			_ = _source19
			var _1154___mcc_h12 _dafny.Array = _source19
			_ = _1154___mcc_h12
			var _1155_cf28 _dafny.Array = _1154___mcc_h12
			_ = _1155_cf28
			(globalState).F10 = p1
			var _1156_v110 _dafny.Map
			_ = _1156_v110
			_1156_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_1152_v108)).Cardinality(), (_dafny.Zero).Minus(_1126_cf9))
			var _1157_v112 _dafny.MultiSet
			_ = _1157_v112
			_1157_v112 = _dafny.MultiSetOf(_1012_v0, _1124_cf11)
			var _1158_v113 _dafny.Sequence
			_ = _1158_v113
			_1158_v113 = _dafny.SeqOf(_1157_v112, _dafny.MultiSetOf(true), _dafny.MultiSetFromSeq(_1066_v41))
			_1156_v110 = (_1156_v110).Update(Companion_Default___.SafeDivisionInt((_this).Fm8(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("isvwjc")).Cardinality()), _dafny.IntOfInt64(-69), _1126_cf9, globalState), p1), (func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter35 := _dafny.Iterate((_1158_v113).Elements()); ; {
					_compr_31, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _1159_v111 _dafny.MultiSet
					_1159_v111 = interface{}(_compr_31).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_1158_v113, _1159_v111) {
						_coll31.Add(_1159_v111, _1126_cf9)
					}
				}
				return _coll31.ToMap()
			}()).Cardinality())
			var _1160_v114 _dafny.Array
			_ = _1160_v114
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_25
			var _nw214 _dafny.Array
			_ = _nw214
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw214 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_1161_cf11 bool) func(_dafny.Int) bool {
					return func(_1162_i10 _dafny.Int) bool {
						return _1161_cf11
					}
				})(_1124_cf11)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw214 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw214).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw214).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1160_v114 = _nw214
			var _1163_v115 _dafny.Map
			_ = _1163_v115
			_1163_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1126_cf9, _1160_v114)
			var _1164_v116 _dafny.Sequence
			_ = _1164_v116
			_1164_v116 = _dafny.SeqOf(_1126_cf9, _1126_cf9)
			_1163_v115 = (_1163_v115).Update(_dafny.IntOfUint32((_1164_v116).Cardinality()), _1160_v114)
			(globalState).F10 = _1126_cf9
		} else if _source18.Is_DC6() {
			var _1165___mcc_h7 _dafny.Int = _source18.Get_().(D1_DC6).Cf13
			_ = _1165___mcc_h7
			var _1166___mcc_h8 _dafny.Sequence = _source18.Get_().(D1_DC6).Cf14
			_ = _1166___mcc_h8
			var _1167___mcc_h9 _dafny.Int = _source18.Get_().(D1_DC6).Cf15
			_ = _1167___mcc_h9
			var _1168___mcc_h10 bool = _source18.Get_().(D1_DC6).Cf16
			_ = _1168___mcc_h10
			var _1169_cf16 bool = _1168___mcc_h10
			_ = _1169_cf16
			var _1170_cf15 _dafny.Int = _1167___mcc_h9
			_ = _1170_cf15
			var _1171_cf14 _dafny.Sequence = _1166___mcc_h8
			_ = _1171_cf14
			var _1172_cf13 _dafny.Int = _1165___mcc_h7
			_ = _1172_cf13
			var _1173_v117 _dafny.Sequence
			_ = _1173_v117
			_1173_v117 = _dafny.UnicodeSeqOfUtf8Bytes("esqyiagdp")
			var _1174_v118 _dafny.CodePoint
			_ = _1174_v118
			_1174_v118 = _dafny.CodePoint('w')
			var _1175_v119 _dafny.Map
			_ = _1175_v119
			_1175_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1012_v0, _dafny.Companion_Sequence_.Update(_1173_v117, (Companion_Default___.SafeIndex(_1172_cf13, _dafny.IntOfUint32((_1173_v117).Cardinality()))).Uint32(), _1174_v118))
			_1173_v117 = _dafny.Companion_Sequence_.Update(_1173_v117, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if !(_1012_v0) {
					return _1173_v117
				}
				return (func() _dafny.Sequence {
					if (_1175_v119).Contains(_1012_v0) {
						return (_1175_v119).Get(_1012_v0).(_dafny.Sequence)
					}
					return p0
				})()
			})()).Cardinality()), _dafny.IntOfUint32((_1173_v117).Cardinality()))).Uint32(), _1174_v118)
			(globalState).F13 = (func() bool {
				if _1012_v0 {
					return _1169_cf16
				}
				return (Companion_Default___.Fm0(_1172_cf13, globalState)).Cmp(p1) < 0
			})()
			var _1176_v120 _dafny.Map
			_ = _1176_v120
			_1176_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1066_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(271), _dafny.IntOfUint32((_1066_v41).Cardinality()))).Uint32()).(bool)), true)
			var _1177_v121 _dafny.Set
			_ = _1177_v121
			_1177_v121 = _dafny.SetOf(_1169_cf16, _1012_v0)
			(globalState).F13 = (func() bool {
				if (_1176_v120).Contains((_1177_v121).IsSubsetOf(_1177_v121)) {
					return (_1176_v120).Get((_1177_v121).IsSubsetOf(_1177_v121)).(bool)
				}
				return _1012_v0
			})()
			var _1178_v122 _dafny.Array
			_ = _1178_v122
			var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw215
			_1178_v122 = _nw215
			var _1179_v123 _dafny.Map
			_ = _1179_v123
			_1179_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(824), _1172_cf13), _1178_v122)
			_1179_v123 = (_1179_v123).Update((_1176_v120).Cardinality(), _1178_v122)
		} else {
			var _1180___mcc_h11 _dafny.Map = _source18.Get_().(D1_DC4).Cf8
			_ = _1180___mcc_h11
			var _1181_cf8 _dafny.Map = _1180___mcc_h11
			_ = _1181_cf8
			var _1182_v124 _dafny.MultiSet
			_ = _1182_v124
			_1182_v124 = _dafny.MultiSetOf(_1012_v0)
			var _1183_v125 _dafny.Sequence
			_ = _1183_v125
			_1183_v125 = _dafny.SeqOf(p1)
			var _1184_v126 _dafny.Map
			_ = _1184_v126
			_1184_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1182_v124).IsProperSubsetOf(_1182_v124)), _1183_v125)
			_1184_v126 = _1184_v126
			var _1185_v127 _dafny.Array
			_ = _1185_v127
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_26
			var _nw216 _dafny.Array
			_ = _nw216
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw216 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) bool = (func(_1186_v0 bool) func(_dafny.Int) bool {
					return func(_1187_i11 _dafny.Int) bool {
						return _1186_v0
					}
				})(_1012_v0)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw216 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw216).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw216).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1185_v127 = _nw216
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1185_v127), 0))
			_ = _index237
			(_1185_v127).ArraySet1((p1).Cmp((_1183_v125).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1183_v125).Cardinality()))).Uint32()).(_dafny.Int)) >= 0, (_index237).Int())
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1185_v127), 0))
			_ = _index238
			(_1185_v127).ArraySet1(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_1012_v0, _1012_v0), (func() _dafny.Sequence {
				if _1012_v0 {
					return _dafny.SeqOf(_1012_v0)
				}
				return _1066_v41
			})()), (_index238).Int())
			(globalState).F13 = (_1185_v127).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(345), _dafny.ArrayLen((_1185_v127), 0))).Int()).(bool)
			var _1188_v128 _dafny.Array
			_ = _1188_v128
			var _nw217 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw217
			_1188_v128 = _nw217
			var _1189_v129 _dafny.Set
			_ = _1189_v129
			_1189_v129 = _dafny.SetOf(true)
			var _1190_v130 _dafny.Array
			_ = _1190_v130
			var _nwElement0_56 _dafny.MultiSet = _1113_v79
			_ = _nwElement0_56
			var _nw218 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(9))
			_ = _nw218
			(_nw218).ArraySet1(_nwElement0_56, 0)
			(_nw218).ArraySet1(_1113_v79, 1)
			(_nw218).ArraySet1(_1113_v79, 2)
			(_nw218).ArraySet1(_dafny.MultiSetOf(p1, p1), 3)
			(_nw218).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(30))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}(func(_1191_i12 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(961)
			}))), 4)
			(_nw218).ArraySet1(_1113_v79, 5)
			(_nw218).ArraySet1(_1113_v79, 6)
			(_nw218).ArraySet1(_1113_v79, 7)
			(_nw218).ArraySet1(_dafny.MultiSetOf((_1189_v129).Cardinality(), p1), 8)
			_1190_v130 = _nw218
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_1188_v128), 0))
			_ = _index239
			(_1188_v128).ArraySet1(_1190_v130, (_index239).Int())
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(19), _dafny.ArrayLen((_1188_v128), 0))
			_ = _index240
			var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw219
			(_1188_v128).ArraySet1(_nw219, (_index240).Int())
		}
		var _1192_v131 _dafny.Array
		_ = _1192_v131
		var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
		_ = _nw220
		_1192_v131 = _nw220
		var _1193_v132 _dafny.MultiSet
		_ = _1193_v132
		_1193_v132 = _dafny.MultiSetOf(_1192_v131, _1192_v131)
		r0 = (_1193_v132).IsProperSubsetOf((_1193_v132).Union(_1193_v132))
		r1 = (p1).Minus(_dafny.IntOfInt64(367))
		var _1194_v133 T0
		_ = _1194_v133
		var _nw221 *C1 = New_C1_()
		_ = _nw221
		_nw221.Ctor__()
		_1194_v133 = _nw221
		r2 = _1194_v133
		return r0, r1, r2
	}
}
func (_this *C3) M5(p0 _dafny.Int, p1 bool, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		(globalState).F22 = false
		(globalState).F22 = p3
		var _1195_v0 _dafny.Set
		_ = _1195_v0
		_1195_v0 = _dafny.SetOf(p0, (_dafny.Zero).Minus(p0), p0)
		var _hi5 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(212))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}(func(_1196_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))).Cardinality())
		_ = _hi5
		for _1197_i0 := ((_1195_v0).Cardinality()).Minus(p0); _1197_i0.Cmp(_hi5) < 0; _1197_i0 = _1197_i0.Plus(_dafny.One) {
			var _1198_v1 _dafny.Sequence
			_ = _1198_v1
			_1198_v1 = _dafny.UnicodeSeqOfUtf8Bytes("rfyol")
			_1198_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("arqqntbg"), _1198_v1), _1198_v1)
			var _1199_v2 _dafny.Array
			_ = _1199_v2
			var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
			_ = _nw222
			_1199_v2 = _nw222
			var _1200_v3 _dafny.Map
			_ = _1200_v3
			_1200_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12(p3, !(p2), globalState), _1199_v2)
			_1200_v3 = (_1200_v3).Update(p3, _1199_v2)
			var _rhs231 bool = p2
			_ = _rhs231
			var _rhs232 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1198_v1, _1198_v1)).Cardinality()), _1197_i0)
			_ = _rhs232
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			var _lhs207 *GlobalState = globalState
			_ = _lhs207
			_lhs206.F13 = _rhs231
			_lhs207.F10 = _rhs232
			var _1201_v4 _dafny.Array
			_ = _1201_v4
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_27
			var _nw223 _dafny.Array
			_ = _nw223
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw223 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) bool = (func(_1202_p3 bool) func(_dafny.Int) bool {
					return func(_1203_i2 _dafny.Int) bool {
						return _1202_p3
					}
				})(p3)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw223 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw223).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw223).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1201_v4 = _nw223
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1201_v4), 0))
			_ = _index241
			(_1201_v4).ArraySet1((_1195_v0).IsSubsetOf(_1195_v0), (_index241).Int())
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((_1201_v4), 0))
			_ = _index242
			(_1201_v4).ArraySet1(p3, (_index242).Int())
		}
		var _1204_v5 _dafny.CodePoint
		_ = _1204_v5
		_1204_v5 = _dafny.CodePoint('s')
		var _1205_v7 _dafny.MultiSet
		_ = _1205_v7
		_1205_v7 = _dafny.MultiSetOf(p3)
		var _1206_v8 _dafny.Sequence
		_ = _1206_v8
		_1206_v8 = _dafny.SeqOf(_1205_v7)
		var _1207_v9 D1
		_ = _1207_v9
		_1207_v9 = Companion_D1_.Create_DC4_(func() _dafny.Map {
			var _coll32 = _dafny.NewMapBuilder()
			_ = _coll32
			for _iter36 := _dafny.Iterate((_1206_v8).Elements()); ; {
				_compr_32, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _1208_v6 _dafny.MultiSet
				_1208_v6 = interface{}(_compr_32).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_1206_v8, _1208_v6) {
					_coll32.Add(_1208_v6, p3)
				}
			}
			return _coll32.ToMap()
		}())
		var _1209_v10 _dafny.Map
		_ = _1209_v10
		_1209_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("bk"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bk")).Cardinality()))).Uint32(), _1204_v5), _1207_v9)
		var _1210_v11 _dafny.Map
		_ = _1210_v11
		_1210_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1209_v10)
		var _1211_v12 _dafny.Sequence
		_ = _1211_v12
		_1211_v12 = _dafny.UnicodeSeqOfUtf8Bytes("rwbrc")
		_1210_v11 = (_1210_v11).Update(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Update(_1211_v12, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1211_v12).Cardinality()))).Uint32(), _1204_v5), _dafny.UnicodeSeqOfUtf8Bytes("wouynh"))), (_1209_v10).Update(_1211_v12, _1207_v9))
		var _1212_i3 _dafny.Int
		_ = _1212_i3
		_1212_i3 = _dafny.Zero
		{
			for p1 {
				{
					if (_1212_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_1212_i3 = (_1212_i3).Plus(_dafny.One)
					var _1213_v13 _dafny.Map
					_ = _1213_v13
					_1213_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1211_v12).Cardinality()))
					var _1214_v14 _dafny.Sequence
					_ = _1214_v14
					_1214_v14 = _dafny.SeqOf(p1, p1)
					_1213_v13 = Companion_Default___.Fm33((p3) && (true), (p2) == (p2), (_1214_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1214_v14).Cardinality()))).Uint32()).(bool), p1, globalState)
					var _1215_v15 *C0
					_ = _1215_v15
					var _nw224 *C0 = New_C0_()
					_ = _nw224
					_nw224.Ctor__(p2)
					_1215_v15 = _nw224
					var _1216_v16 *C0
					_ = _1216_v16
					var _nw225 *C0 = New_C0_()
					_ = _nw225
					_nw225.Ctor__((p0).Cmp(p0) != 0)
					_1216_v16 = _nw225
					var _1217_v17 _dafny.MultiSet
					_ = _1217_v17
					_1217_v17 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_1211_v12).Cardinality()))
					var _1218_v18 _dafny.Map
					_ = _1218_v18
					_1218_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_1214_v14, _1214_v14), _1217_v17)
					var _1219_v19 *C2
					_ = _1219_v19
					var _nw226 *C2 = New_C2_()
					_ = _nw226
					_nw226.Ctor__(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(45))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg60 _dafny.Int) interface{} {
							return coer60(arg60)
						}
					}((func(_1220_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1221_i4 _dafny.Int) _dafny.CodePoint {
							return _1220_v5
						}
					})(_1204_v5))))
					_1219_v19 = _nw226
					var _1222_v20 _dafny.Map
					_ = _1222_v20
					_1222_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1219_v19, p0)
					var _1223_v21 _dafny.Map
					_ = _1223_v21
					_1223_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_1219_v19.F33).Cardinality()))
					var _rhs233 _dafny.Int = (_1218_v18).Cardinality()
					_ = _rhs233
					var _rhs234 _dafny.Int = (func() _dafny.Int {
						if (_1222_v20).Contains(_1219_v19) {
							return (_1222_v20).Get(_1219_v19).(_dafny.Int)
						}
						return p0
					})()
					_ = _rhs234
					var _rhs235 _dafny.Int = (_dafny.Zero).Minus((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.Zero).Minus(p0))).Merge(_1223_v21)).Update((_1216_v16).F32(), p0)).Cardinality())
					_ = _rhs235
					var _lhs208 *GlobalState = globalState
					_ = _lhs208
					var _lhs209 *GlobalState = globalState
					_ = _lhs209
					var _lhs210 *GlobalState = globalState
					_ = _lhs210
					_lhs208.F10 = _rhs233
					_lhs209.F10 = _rhs234
					_lhs210.F10 = _rhs235
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _1224_v22 _dafny.Array
		_ = _1224_v22
		var _nw227 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
		_ = _nw227
		_1224_v22 = _nw227
		var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1224_v22), 0))
		_ = _index243
		(_1224_v22).ArraySet1(p0, (_index243).Int())
		var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1224_v22), 0))
		_ = _index244
		(_1224_v22).ArraySet1(p0, (_index244).Int())
		r0 = p0
		r1 = (_1224_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(537), _dafny.ArrayLen((_1224_v22), 0))).Int()).(_dafny.Int)
		return r0, r1
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f31 _dafny.Array
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f31 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C4) Ctor__(f31 _dafny.Array) {
	{
		(_this)._f31 = f31
	}
}
func (_this *C4) Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		if (func() bool {
			if false {
				return false
			}
			return !(true)
		})() {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(!(true)), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true, true), true))
		} else {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), false))
		}
	}
}
func (_this *C4) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((func() _dafny.Int {
			if (func() bool {
				if false {
					return false
				}
				return false
			})() {
				return (_dafny.IntOfInt64(973)).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(631))).Cardinality())))
			}
			return _dafny.IntOfUint32(((func() _dafny.Sequence {
				if true {
					return _dafny.SeqOf(_dafny.CodePoint('t'))
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(357))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}(func(_1225_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('a')
				}))
			})()).Cardinality())
		})())
	}
}
func (_this *C4) Fm5(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) bool {
	{
		if !(!(_dafny.MultiSetOf(_dafny.IntOfInt64(-735))).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(669)))) {
			return true
		} else if !(true) {
			return false
		} else {
			return false
		}
	}
}
func (_this *C4) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-665)
	}
}
func (_this *C4) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _rhs236 bool = p3
		_ = _rhs236
		var _rhs237 bool = p1
		_ = _rhs237
		var _rhs238 _dafny.Int = p2
		_ = _rhs238
		var _lhs211 *GlobalState = globalState
		_ = _lhs211
		var _lhs212 *GlobalState = globalState
		_ = _lhs212
		var _lhs213 *GlobalState = globalState
		_ = _lhs213
		_lhs211.F13 = _rhs236
		_lhs212.F22 = _rhs237
		_lhs213.F10 = _rhs238
		var _1226_i0 _dafny.Int
		_ = _1226_i0
		_1226_i0 = _dafny.Zero
		{
			for (p2).Cmp((func() _dafny.Int {
				if p1 {
					return _dafny.IntOfInt64(363)
				}
				return p2
			})()) <= 0 {
				{
					if (_1226_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1226_i0 = (_1226_i0).Plus(_dafny.One)
					var _1227_v0 _dafny.Sequence
					_ = _1227_v0
					_1227_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gjwrovu")
					var _1228_v1 _dafny.CodePoint
					_ = _1228_v1
					_1228_v1 = _dafny.CodePoint('o')
					var _1229_v2 _dafny.Map
					_ = _1229_v2
					_1229_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1227_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(993))).Cardinality()), _dafny.IntOfUint32((_1227_v0).Cardinality()))).Uint32(), _1228_v1)).Cardinality()), _dafny.MultiSetOf(p2))
					_1229_v2 = (_1229_v2).Update(p2, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(979), p0, p0))).Update(p0, Companion_Default___.Abs(p2)))
					var _1230_v3 _dafny.Set
					_ = _1230_v3
					_1230_v3 = _dafny.SetOf(p3, p3)
					var _1231_v4 T0
					_ = _1231_v4
					var _nw228 *C3 = New_C3_()
					_ = _nw228
					_nw228.Ctor__()
					_1231_v4 = _nw228
					var _1232_v5 _dafny.Sequence
					_ = _1232_v5
					_1232_v5 = _dafny.SeqOf(_1231_v4)
					var _1233_v6 _dafny.MultiSet
					_ = _1233_v6
					_1233_v6 = _dafny.MultiSetOf(_1230_v3, _1230_v3)
					if ((_dafny.MultiSetOf(Companion_Default___.Fm7(p1, p3, globalState), _1230_v3, _dafny.SetOf(p1))).Update(_dafny.SetOf(p3), Companion_Default___.Abs(_dafny.IntOfUint32((_1232_v5).Cardinality())))).IsSubsetOf(_1233_v6) {
						var _1234_v7 D2
						_ = _1234_v7
						_1234_v7 = Companion_D2_.Create_DC8_(p3, p1)
						var _1235_v8 D2
						_ = _1235_v8
						_1235_v8 = Companion_D2_.Create_DC9_(_1234_v7)
						var _1236_v9 _dafny.Map
						_ = _1236_v9
						_1236_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v1, Companion_Default___.SafeDivisionInt(p0, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1235_v8)).Update(p2, _1235_v8)).Cardinality()))
						_1236_v9 = (_1236_v9).Update(_1228_v1, p2)
						var _1237_v10 _dafny.Array
						_ = _1237_v10
						var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
						_ = _nw229
						_1237_v10 = _nw229
						var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1237_v10), 0))
						_ = _index245
						(_1237_v10).ArraySet1(p0, (_index245).Int())
						var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1237_v10), 0))
						_ = _index246
						(_1237_v10).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0), (_index246).Int())
						_1237_v10 = _1237_v10
						var _1238_v11 _dafny.Sequence
						_ = _1238_v11
						_1238_v11 = _dafny.SeqOf(_dafny.IntOfInt64(742), (_this).Fm2(globalState), (_1237_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1237_v10), 0))).Int()).(_dafny.Int))
						var _1239_v12 _dafny.Map
						_ = _1239_v12
						_1239_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_1238_v11).Cardinality()))
						var _1240_v13 _dafny.Set
						_ = _1240_v13
						_1240_v13 = _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p2))
						(globalState).F7 = !((_dafny.SetOf(_1239_v12)).Difference(_1240_v13)).Equals(_dafny.SetOf(_1239_v12, _1239_v12, _1239_v12, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(208)), _1239_v12))
						var _1241_v14 _dafny.Array
						_ = _1241_v14
						var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
						_ = _nw230
						_1241_v14 = _nw230
						var _1242_v15 _dafny.Array
						_ = _1242_v15
						var _nw231 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
						_ = _nw231
						_1242_v15 = _nw231
						var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1241_v14), 0))
						_ = _index247
						(_1241_v14).ArraySet1(_1242_v15, (_index247).Int())
						var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1241_v14), 0))
						_ = _index248
						(_1241_v14).ArraySet1(_1242_v15, (_index248).Int())
					} else {
						var _1243_v16 _dafny.Array
						_ = _1243_v16
						var _len0_28 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_28
						var _nw232 _dafny.Array
						_ = _nw232
						if _len0_28.Cmp(_dafny.Zero) == 0 {
							_nw232 = _dafny.NewArray(_len0_28)
						} else {
							var _init28 func(_dafny.Int) bool = (func(_1244_p1 bool) func(_dafny.Int) bool {
								return func(_1245_i1 _dafny.Int) bool {
									return _1244_p1
								}
							})(p1)
							_ = _init28
							var _element0_28 = _init28(_dafny.Zero)
							_ = _element0_28
							_nw232 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
							(_nw232).ArraySet1(_element0_28, 0)
							var _nativeLen0_28 = (_len0_28).Int()
							_ = _nativeLen0_28
							for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
								(_nw232).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
							}
						}
						_1243_v16 = _nw232
						var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1243_v16), 0))
						_ = _index249
						(_1243_v16).ArraySet1(p3, (_index249).Int())
						var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1243_v16), 0))
						_ = _index250
						(_1243_v16).ArraySet1(true, (_index250).Int())
						var _1246_v17 _dafny.Array
						_ = _1246_v17
						var _len0_29 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_29
						var _nw233 _dafny.Array
						_ = _nw233
						if _len0_29.Cmp(_dafny.Zero) == 0 {
							_nw233 = _dafny.NewArray(_len0_29)
						} else {
							var _init29 func(_dafny.Int) _dafny.Int = (func(_1247_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
								return func(_1248_i2 _dafny.Int) _dafny.Int {
									return (_1248_i2).Plus(_dafny.IntOfUint32((_1247_v0).Cardinality()))
								}
							})(_1227_v0)
							_ = _init29
							var _element0_29 = _init29(_dafny.Zero)
							_ = _element0_29
							_nw233 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
							(_nw233).ArraySet1(_element0_29, 0)
							var _nativeLen0_29 = (_len0_29).Int()
							_ = _nativeLen0_29
							for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
								(_nw233).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
							}
						}
						_1246_v17 = _nw233
						var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_1246_v17), 0))
						_ = _index251
						(_1246_v17).ArraySet1(p2, (_index251).Int())
						var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_1246_v17), 0))
						_ = _index252
						var _rhs239 _dafny.Int = p0
						_ = _rhs239
						var _rhs240 bool = !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("xnwqqkdl"), _1228_v1)
						_ = _rhs240
						var _rhs241 _dafny.Int = p2
						_ = _rhs241
						var _lhs214 *GlobalState = globalState
						_ = _lhs214
						var _lhs215 *GlobalState = globalState
						_ = _lhs215
						var _lhs216 _dafny.Array = _1246_v17
						_ = _lhs216
						var _lhs217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_1246_v17), 0))
						_ = _lhs217
						_lhs214.F10 = _rhs239
						_lhs215.F13 = _rhs240
						(_lhs216).ArraySet1(_rhs241, (_lhs217).Int())
						var _1249_v18 *C0
						_ = _1249_v18
						var _nw234 *C0 = New_C0_()
						_ = _nw234
						_nw234.Ctor__(false)
						_1249_v18 = _nw234
						var _1250_v19 _dafny.Map
						_ = _1250_v19
						_1250_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1228_v1)
						var _1251_v20 _dafny.Sequence
						_ = _1251_v20
						_1251_v20 = _dafny.SeqOf(p1)
						var _1252_v21 _dafny.Sequence
						_ = _1252_v21
						_1252_v21 = _dafny.SeqOf((_1251_v20).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1251_v20).Cardinality()))).Uint32()).(bool))
						_1250_v19 = (_1250_v19).Update(!_dafny.Companion_Sequence_.Equal(_1252_v21, _1252_v21), _1228_v1)
						(globalState).F10 = (_dafny.Zero).Minus(p0)
					}
					var _1253_v23 _dafny.Map
					_ = _1253_v23
					_1253_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1228_v1, p2)
					var _1254_v24 _dafny.MultiSet
					_ = _1254_v24
					_1254_v24 = _dafny.MultiSetOf(p1, p1)
					var _1255_v25 _dafny.Map
					_ = _1255_v25
					_1255_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1254_v24, p3)
					var _1256_v26 D1
					_ = _1256_v26
					_1256_v26 = Companion_D1_.Create_DC4_(_1255_v25)
					var _pat_let_tv36 = _1255_v25
					_ = _pat_let_tv36
					var _1257_v27 _dafny.Map
					_ = _1257_v27
					_1257_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
						var _coll33 = _dafny.NewMapBuilder()
						_ = _coll33
						for _iter37 := _dafny.Iterate((_1253_v23).Keys().Elements()); ; {
							_compr_33, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _1258_v22 _dafny.CodePoint
							_1258_v22 = interface{}(_compr_33).(_dafny.CodePoint)
							if (_1253_v23).Contains(_1258_v22) {
								_coll33.Add(_1258_v22, p2)
							}
						}
						return _coll33.ToMap()
					}()).Cardinality(), func(_pat_let38_0 D1) D1 {
						return func(_1259_dt__update__tmp_h0 D1) D1 {
							return func(_pat_let39_0 _dafny.Map) D1 {
								return func(_1260_dt__update_hcf8_h0 _dafny.Map) D1 {
									return Companion_D1_.Create_DC4_(_1260_dt__update_hcf8_h0)
								}(_pat_let39_0)
							}(_pat_let_tv36)
						}(_pat_let38_0)
					}(_1256_v26))
					var _1261_v28 _dafny.MultiSet
					_ = _1261_v28
					_1261_v28 = _dafny.MultiSetOf(_1257_v27, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1256_v26), _1257_v27)
					var _1262_v29 _dafny.Map
					_ = _1262_v29
					_1262_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1261_v28).Union(_1261_v28))
					var _1263_v30 _dafny.Set
					_ = _1263_v30
					_1263_v30 = _dafny.SetOf(_1227_v0)
					var _1264_v31 _dafny.Sequence
					_ = _1264_v31
					_1264_v31 = _dafny.SeqOf(_1257_v27)
					_1262_v29 = (_1262_v29).Update(!((_1263_v30).IsProperSubsetOf(_1263_v30)), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1264_v31, _1264_v31)))
					if p3 {
						(globalState).F10 = p2
						var _1265_v32 _dafny.Array
						_ = _1265_v32
						var _len0_30 _dafny.Int = _dafny.IntOfInt64(8)
						_ = _len0_30
						var _nw235 _dafny.Array
						_ = _nw235
						if _len0_30.Cmp(_dafny.Zero) == 0 {
							_nw235 = _dafny.NewArray(_len0_30)
						} else {
							var _init30 func(_dafny.Int) bool = func(_1266_i3 _dafny.Int) bool {
								return true
							}
							_ = _init30
							var _element0_30 = _init30(_dafny.Zero)
							_ = _element0_30
							_nw235 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
							(_nw235).ArraySet1(_element0_30, 0)
							var _nativeLen0_30 = (_len0_30).Int()
							_ = _nativeLen0_30
							for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
								(_nw235).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
							}
						}
						_1265_v32 = _nw235
						var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1265_v32), 0))
						_ = _index253
						(_1265_v32).ArraySet1(p3, (_index253).Int())
						var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1265_v32), 0))
						_ = _index254
						(_1265_v32).ArraySet1((p3) && (p3), (_index254).Int())
						var _1267_v33 *C0
						_ = _1267_v33
						var _nw236 *C0 = New_C0_()
						_ = _nw236
						_nw236.Ctor__(p1)
						_1267_v33 = _nw236
						(globalState).F10 = _dafny.IntOfInt64(383)
						var _1268_v34 _dafny.Sequence
						_ = _1268_v34
						_1268_v34 = _dafny.SeqOf(((_dafny.Zero).Minus((_dafny.Zero).Minus((_this).Fm2(globalState)))).Times(p2), _dafny.IntOfUint32(((func() _dafny.Sequence {
							if !((_1265_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_1265_v32), 0))).Int()).(bool)) {
								return _1227_v0
							}
							return _1227_v0
						})()).Cardinality()), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), p2))
						_1268_v34 = _1268_v34
					} else {
						var _1269_v35 *C3
						_ = _1269_v35
						var _nw237 *C3 = New_C3_()
						_ = _nw237
						_nw237.Ctor__()
						_1269_v35 = _nw237
						var _1270_v36 _dafny.Sequence
						_ = _1270_v36
						_1270_v36 = _dafny.SeqOf(_1227_v0)
						var _1271_v37 _dafny.Map
						_ = _1271_v37
						_1271_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
						var _1272_v38 _dafny.Array
						_ = _1272_v38
						var _nwElement0_57 _dafny.Sequence = Companion_Default___.Fm10(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1227_v0).Cardinality()), (_1230_v3).Cardinality()), _1228_v1, p1, globalState)
						_ = _nwElement0_57
						var _nw238 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(16))
						_ = _nw238
						(_nw238).ArraySet1(_nwElement0_57, 0)
						(_nw238).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1227_v0, _dafny.UnicodeSeqOfUtf8Bytes("gdhniefl")), 1)
						(_nw238).ArraySet1(_1227_v0, 2)
						(_nw238).ArraySet1(_1227_v0, 3)
						(_nw238).ArraySet1(_1227_v0, 4)
						(_nw238).ArraySet1(_1227_v0, 5)
						(_nw238).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("arq"), 6)
						(_nw238).ArraySet1(_1227_v0, 7)
						(_nw238).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(777))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg62 _dafny.Int) interface{} {
								return coer62(arg62)
							}
						}((func(_1273_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1274_i4 _dafny.Int) _dafny.CodePoint {
								return _1273_v1
							}
						})(_1228_v1))), _1227_v0), 8)
						(_nw238).ArraySet1(_1227_v0, 9)
						(_nw238).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1270_v36).Select((Companion_Default___.SafeIndex((_1271_v37).Cardinality(), _dafny.IntOfUint32((_1270_v36).Cardinality()))).Uint32()).(_dafny.Sequence), _1227_v0), 10)
						(_nw238).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pl"), 11)
						(_nw238).ArraySet1(_1227_v0, 12)
						(_nw238).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("abm"), 13)
						(_nw238).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(381))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg63 _dafny.Int) interface{} {
								return coer63(arg63)
							}
						}((func(_1275_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1276_i5 _dafny.Int) _dafny.CodePoint {
								return _1275_v1
							}
						})(_1228_v1))), _1227_v0), 14)
						(_nw238).ArraySet1(_1227_v0, 15)
						_1272_v38 = _nw238
						var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1272_v38), 0))
						_ = _index255
						(_1272_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1227_v0, _1227_v0), (_index255).Int())
						var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(692), _dafny.ArrayLen((_1272_v38), 0))
						_ = _index256
						(_1272_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(914))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg64 _dafny.Int) interface{} {
								return coer64(arg64)
							}
						}((func(_1277_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1278_i6 _dafny.Int) _dafny.CodePoint {
								return _1277_v1
							}
						})(_1228_v1))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(577))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}(func(_1279_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('x')
						})), _1227_v0)), (_index256).Int())
						(globalState).F10 = p0
						var _1280_v39 _dafny.Map
						_ = _1280_v39
						_1280_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
						(globalState).F10 = ((_1280_v39).Update(_dafny.IntOfInt64(591), p1)).Cardinality()
						var _1281_v40 *C1
						_ = _1281_v40
						var _nw239 *C1 = New_C1_()
						_ = _nw239
						_nw239.Ctor__()
						_1281_v40 = _nw239
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _hi6 _dafny.Int = _dafny.IntOfInt64(-238)
		_ = _hi6
		for _1282_i8 := p0; _1282_i8.Cmp(_hi6) < 0; _1282_i8 = _1282_i8.Plus(_dafny.One) {
			var _1283_v41 _dafny.MultiSet
			_ = _1283_v41
			_1283_v41 = _dafny.MultiSetOf(p1, p3)
			var _1284_v42 D3
			_ = _1284_v42
			_1284_v42 = Companion_D3_.Create_DC10_(_1283_v41)
			_1284_v42 = (func() D3 {
				if p3 {
					return _1284_v42
				}
				return _1284_v42
			})()
			var _1285_v43 T0
			_ = _1285_v43
			var _nw240 *C1 = New_C1_()
			_ = _nw240
			_nw240.Ctor__()
			_1285_v43 = _nw240
			_1285_v43 = _1285_v43
			var _1286_v44 _dafny.Sequence
			_ = _1286_v44
			_1286_v44 = _dafny.UnicodeSeqOfUtf8Bytes("v")
			var _1287_v45 _dafny.Map
			_ = _1287_v45
			_1287_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1286_v44, _dafny.IntOfUint32((_1286_v44).Cardinality()))
			var _1288_v46 _dafny.Map
			_ = _1288_v46
			_1288_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
			var _1289_v47 _dafny.Map
			_ = _1289_v47
			_1289_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_i8, p2)
			var _1290_v48 _dafny.MultiSet
			_ = _1290_v48
			_1290_v48 = _dafny.MultiSetOf(Companion_D2_.Create_DC7_(_1289_v47))
			var _1291_v49 _dafny.Sequence
			_ = _1291_v49
			_1291_v49 = _dafny.SeqOf(_1290_v48)
			var _1292_v50 _dafny.Sequence
			_ = _1292_v50
			_1292_v50 = _dafny.SeqOf((_1288_v46).Cardinality(), ((_1291_v49).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1291_v49).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality())
			var _1293_v51 _dafny.Array
			_ = _1293_v51
			var _nwElement0_58 _dafny.Int = p0
			_ = _nwElement0_58
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(15))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_58, 0)
			(_nw241).ArraySet1((func() _dafny.Int {
				if (_1287_v45).Contains(_1286_v44) {
					return (_1287_v45).Get(_1286_v44).(_dafny.Int)
				}
				return p2
			})(), 1)
			(_nw241).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(682), _dafny.IntOfInt64(-113)), 2)
			(_nw241).ArraySet1(_1282_i8, 3)
			(_nw241).ArraySet1(Companion_Default___.Fm0(_1282_i8, globalState), 4)
			(_nw241).ArraySet1(((_1283_v41).Cardinality()).Times(p2), 5)
			(_nw241).ArraySet1(_dafny.IntOfInt64(584), 6)
			(_nw241).ArraySet1(p0, 7)
			(_nw241).ArraySet1(p0, 8)
			(_nw241).ArraySet1((_1282_i8).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-7))), 9)
			(_nw241).ArraySet1(p2, 10)
			(_nw241).ArraySet1(p2, 11)
			(_nw241).ArraySet1((_1285_v43).Fm2(globalState), 12)
			(_nw241).ArraySet1((_1292_v50).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1292_v50).Cardinality()))).Uint32()).(_dafny.Int), 13)
			(_nw241).ArraySet1(p0, 14)
			_1293_v51 = _nw241
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1293_v51), 0))
			_ = _index257
			(_1293_v51).ArraySet1((func() _dafny.Int {
				if p1 {
					return p2
				}
				return p2
			})(), (_index257).Int())
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1293_v51), 0))
			_ = _index258
			var _rhs242 _dafny.Int = (_1282_i8).Plus(p0)
			_ = _rhs242
			var _rhs243 _dafny.Int = p2
			_ = _rhs243
			var _lhs218 *GlobalState = globalState
			_ = _lhs218
			var _lhs219 _dafny.Array = _1293_v51
			_ = _lhs219
			var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1293_v51), 0))
			_ = _lhs220
			_lhs218.F10 = _rhs242
			(_lhs219).ArraySet1(_rhs243, (_lhs220).Int())
			if !(!(p1)) {
				(globalState).F10 = Companion_Default___.SafeDivisionInt((_1293_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1293_v51), 0))).Int()).(_dafny.Int), _1282_i8)
				(globalState).F7 = !(p1)
				var _1294_v52 _dafny.CodePoint
				_ = _1294_v52
				_1294_v52 = _dafny.CodePoint('l')
				var _1295_v53 D4
				_ = _1295_v53
				_1295_v53 = Companion_D4_.Create_DC13_(_1294_v52)
				var _1296_v54 D4
				_ = _1296_v54
				_1296_v54 = Companion_D4_.Create_DC15_(_1295_v53)
				var _1297_v55 D4
				_ = _1297_v55
				_1297_v55 = Companion_D4_.Create_DC15_(_1295_v53)
				var _1298_v56 _dafny.Map
				_ = _1298_v56
				_1298_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_i8, _1297_v55)
				_1298_v56 = _1298_v56
				var _1299_v57 _dafny.Array
				_ = _1299_v57
				var _nw242 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
				_ = _nw242
				_1299_v57 = _nw242
				var _1300_v58 _dafny.Set
				_ = _1300_v58
				_1300_v58 = _dafny.SetOf(_1299_v57, _1299_v57, _1299_v57)
				(globalState).F22 = !(_1300_v58).Contains(_1299_v57)
				_1286_v44 = _1286_v44
			} else {
				var _1301_v59 _dafny.Set
				_ = _1301_v59
				_1301_v59 = _dafny.SetOf(_1282_i8)
				var _1302_v60 D8
				_ = _1302_v60
				_1302_v60 = Companion_D8_.Create_DC23_((_1301_v59).IsProperSubsetOf(_1301_v59))
				_1302_v60 = _1302_v60
				(globalState).F13 = p1
				var _1303_v61 *C2
				_ = _1303_v61
				var _nw243 *C2 = New_C2_()
				_ = _nw243
				_nw243.Ctor__(_1286_v44)
				_1303_v61 = _nw243
				(globalState).F13 = p1
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_1293_v51), 0))
				_ = _index259
				(_1293_v51).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg66 _dafny.Int) interface{} {
						return coer66(arg66)
					}
				}(func(_1304_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				}))).Cardinality()), (_index259).Int())
			}
		}
		var _1305_v62 _dafny.Array
		_ = _1305_v62
		var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
		_ = _nw244
		_1305_v62 = _nw244
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1305_v62), 0))); ; {
			_guard_loop_3, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _1306_i10 _dafny.Int
			_1306_i10 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1306_i10).Sign() != -1) && ((_1306_i10).Cmp(_dafny.ArrayLen((_1305_v62), 0)) < 0)) {
				(_1305_v62).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(704))).Uint32(), func(coer67 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}(func(_1307_i11 _dafny.Int) D4 {
					return Companion_D4_.Create_DC13_(_dafny.CodePoint('n'))
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(83))).Uint32(), func(coer68 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_1308_i12 _dafny.Int) D4 {
					return Companion_D4_.Create_DC13_(_dafny.CodePoint('q'))
				}))), (_1306_i10).Int())
			}
		}
		if (func() bool {
			if !((p0).Cmp(p2) != 0) {
				return !(p3) || (p1)
			}
			return p3
		})() {
			(globalState).F10 = p2
			var _1309_v63 *C0
			_ = _1309_v63
			var _nw245 *C0 = New_C0_()
			_ = _nw245
			_nw245.Ctor__(p1)
			_1309_v63 = _nw245
			var _1310_v64 _dafny.Array
			_ = _1310_v64
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_31
			var _nw246 _dafny.Array
			_ = _nw246
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw246 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Sequence = func(_1311_i13 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(false)
				}
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw246 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw246).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw246).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1310_v64 = _nw246
			var _1312_v65 _dafny.Sequence
			_ = _1312_v65
			_1312_v65 = _dafny.SeqOf(_1310_v64)
			var _1313_v66 _dafny.CodePoint
			_ = _1313_v66
			_1313_v66 = _dafny.CodePoint('a')
			var _1314_v67 _dafny.Map
			_ = _1314_v67
			_1314_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1312_v65).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1312_v65).Cardinality()))).Uint32()).(_dafny.Array), _1313_v66)
			_1314_v67 = (_1314_v67).Update(_1310_v64, _1313_v66)
			(globalState).F7 = p3
			var _1315_v68 _dafny.MultiSet
			_ = _1315_v68
			_1315_v68 = _dafny.MultiSetOf(p3, true, p3, p1, p3)
			var _1316_v69 _dafny.Sequence
			_ = _1316_v69
			_1316_v69 = _dafny.SeqOf((func() _dafny.MultiSet {
				if (_1309_v63).F32() {
					return _1315_v68
				}
				return _1315_v68
			})())
			_1316_v69 = _1316_v69
		} else {
			var _1317_v70 _dafny.Map
			_ = _1317_v70
			_1317_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
			var _1318_v71 _dafny.MultiSet
			_ = _1318_v71
			_1318_v71 = _dafny.MultiSetOf(p3, p1, p1, (func() bool {
				if (_1317_v70).Contains(p2) {
					return (_1317_v70).Get(p2).(bool)
				}
				return p3
			})(), p3)
			var _1319_v72 _dafny.Set
			_ = _1319_v72
			_1319_v72 = _dafny.SetOf(_1318_v71, _1318_v71, _dafny.MultiSetOf(!(p3), p1))
			_1319_v72 = _1319_v72
			var _pat_let_tv37 = _1318_v71
			_ = _pat_let_tv37
			var _pat_let_tv38 = p1
			_ = _pat_let_tv38
			var _pat_let_tv39 = p2
			_ = _pat_let_tv39
			var _source20 D3 = func(_pat_let40_0 D3) D3 {
				return func(_1320_dt__update__tmp_h1 D3) D3 {
					return func(_pat_let41_0 _dafny.MultiSet) D3 {
						return func(_1321_dt__update_hcf21_h0 _dafny.MultiSet) D3 {
							return Companion_D3_.Create_DC10_(_1321_dt__update_hcf21_h0)
						}(_pat_let41_0)
					}((_pat_let_tv37).Update(_pat_let_tv38, Companion_Default___.Abs((_dafny.Zero).Minus(_pat_let_tv39))))
				}(_pat_let40_0)
			}(Companion_D3_.Create_DC10_(Companion_Default___.Fm24(p2, p1, globalState)))
			_ = _source20
			if _source20.Is_DC11() {
				var _1322___mcc_h0 _dafny.Int = _source20.Get_().(D3_DC11).Cf22
				_ = _1322___mcc_h0
				var _1323_cf22 _dafny.Int = _1322___mcc_h0
				_ = _1323_cf22
				var _1324_v73 _dafny.Array
				_ = _1324_v73
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_32
				var _nw247 _dafny.Array
				_ = _nw247
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw247 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) bool = (func(_1325_p1 bool) func(_dafny.Int) bool {
						return func(_1326_i14 _dafny.Int) bool {
							return _1325_p1
						}
					})(p1)
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw247 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw247).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw247).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1324_v73 = _nw247
				var _1327_v74 _dafny.Set
				_ = _1327_v74
				_1327_v74 = _dafny.SetOf(_1324_v73, _1324_v73)
				var _1328_v75 _dafny.Sequence
				_ = _1328_v75
				_1328_v75 = _dafny.SeqOf(_1327_v74, _1327_v74, _1327_v74)
				(globalState).F10 = ((_1328_v75).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1328_v75).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
				var _1329_v76 *C0
				_ = _1329_v76
				var _nw248 *C0 = New_C0_()
				_ = _nw248
				_nw248.Ctor__(p3)
				_1329_v76 = _nw248
				_1329_v76 = _1329_v76
				(globalState).F19 = _dafny.CodePoint('f')
				var _1330_v77 *C1
				_ = _1330_v77
				var _nw249 *C1 = New_C1_()
				_ = _nw249
				_nw249.Ctor__()
				_1330_v77 = _nw249
				var _1331_v78 _dafny.Set
				_ = _1331_v78
				_1331_v78 = _dafny.SetOf(_1323_cf22)
				var _1332_v79 _dafny.Sequence
				_ = _1332_v79
				_1332_v79 = _dafny.SeqOf((_1331_v78).Equals(_1331_v78))
				var _1333_v80 _dafny.Sequence
				_ = _1333_v80
				_1333_v80 = _dafny.SeqOf(_1323_cf22, _1323_cf22)
				var _1334_v81 _dafny.Sequence
				_ = _1334_v81
				_1334_v81 = _dafny.UnicodeSeqOfUtf8Bytes("pwm")
				var _rhs244 *C1 = _1330_v77
				_ = _rhs244
				var _rhs245 _dafny.Int = (_1333_v80).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1333_v80).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs245
				var _rhs246 bool = ((_this).Fm2(globalState)).Cmp(_dafny.IntOfUint32((_1334_v81).Cardinality())) != 0
				_ = _rhs246
				var _rhs247 _dafny.Sequence = (func() _dafny.Sequence {
					if false {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1332_v79, (Companion_Default___.SafeIndex(_1323_cf22, _dafny.IntOfUint32((_1332_v79).Cardinality()))).Uint32(), (_1329_v76).F32()), _1332_v79)
					}
					return _1332_v79
				})()
				_ = _rhs247
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				var _lhs222 *GlobalState = globalState
				_ = _lhs222
				_1330_v77 = _rhs244
				_lhs221.F10 = _rhs245
				_lhs222.F22 = _rhs246
				_1332_v79 = _rhs247
			} else if _source20.Is_DC12() {
				var _1335___mcc_h1 bool = _source20.Get_().(D3_DC12).Cf23
				_ = _1335___mcc_h1
				var _1336_cf23 bool = _1335___mcc_h1
				_ = _1336_cf23
				var _1337_v82 _dafny.CodePoint
				_ = _1337_v82
				_1337_v82 = _dafny.CodePoint('j')
				var _1338_v83 D8
				_ = _1338_v83
				_1338_v83 = Companion_D8_.Create_DC22_(_dafny.IntOfInt64(652), p1, _1337_v82)
				_1336_cf23 = ((_dafny.Zero).Minus((_1338_v83).Dtor_cf36())).Cmp(p2) != 0
				var _1339_v84 _dafny.Set
				_ = _1339_v84
				_1339_v84 = _dafny.SetOf(_dafny.IntOfInt64(708))
				var _1340_v85 _dafny.Sequence
				_ = _1340_v85
				_1340_v85 = _dafny.UnicodeSeqOfUtf8Bytes("tmjx")
				var _1341_v86 _dafny.MultiSet
				_ = _1341_v86
				_1341_v86 = _dafny.MultiSetOf(p2, (_1339_v84).Cardinality(), Companion_Default___.SafeModuloInt((Companion_Default___.Fm39(globalState)).Cardinality(), p2), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ywd"), _1340_v85)).Cardinality())))
				var _1342_v87 _dafny.Sequence
				_ = _1342_v87
				_1342_v87 = _dafny.SeqOf(_1340_v85, _dafny.UnicodeSeqOfUtf8Bytes("vwntchfj"))
				var _1343_v88 _dafny.Sequence
				_ = _1343_v88
				_1343_v88 = _dafny.SeqOf((_this).Fm5(p3, _1342_v87, false, globalState), _1336_cf23)
				var _1344_v89 _dafny.Sequence
				_ = _1344_v89
				_1344_v89 = _dafny.SeqOf(_dafny.IntOfUint32((_1343_v88).Cardinality()), _dafny.IntOfInt64(184))
				_1341_v86 = (_1341_v86).Union(_dafny.MultiSetFromSeq(_1344_v89))
				(globalState).F10 = Companion_Default___.SafeDivisionInt((p0).Minus(p0), _dafny.IntOfInt64(774))
				var _1345_v90 _dafny.Array
				_ = _1345_v90
				var _nwElement0_59 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(433))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg69 _dafny.Int) interface{} {
						return coer69(arg69)
					}
				}((func(_1346_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1347_i15 _dafny.Int) _dafny.CodePoint {
						return _1346_v82
					}
				})(_1337_v82)))
				_ = _nwElement0_59
				var _nw250 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(11))
				_ = _nw250
				(_nw250).ArraySet1(_nwElement0_59, 0)
				(_nw250).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1340_v85, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1340_v85).Cardinality()))).Uint32(), _1337_v82), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(465))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1348_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1349_i16 _dafny.Int) _dafny.CodePoint {
						return _1348_v82
					}
				})(_1337_v82))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1340_v85).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(465))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1350_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1351_i16 _dafny.Int) _dafny.CodePoint {
						return _1350_v82
					}
				})(_1337_v82)))).Cardinality()))).Uint32(), _1337_v82)).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1340_v85, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1340_v85).Cardinality()))).Uint32(), _1337_v82)).Cardinality()))).Uint32(), _dafny.CodePoint('h')), 1)
				(_nw250).ArraySet1(_1340_v85, 2)
				(_nw250).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1352_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1353_i17 _dafny.Int) _dafny.CodePoint {
						return _1352_v82
					}
				})(_1337_v82))), 3)
				(_nw250).ArraySet1(_1340_v85, 4)
				(_nw250).ArraySet1(_1340_v85, 5)
				(_nw250).ArraySet1(Companion_Default___.Fm10(Companion_Default___.Fm33(p3, p3, true, true, globalState), _1337_v82, p3, globalState), 6)
				(_nw250).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("pyctsgo"), 7)
				(_nw250).ArraySet1(_dafny.Companion_Sequence_.Update(_1340_v85, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1340_v85).Cardinality()), _dafny.IntOfUint32((_1340_v85).Cardinality()))).Uint32(), _1337_v82), 8)
				(_nw250).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("qdgk"), 9)
				(_nw250).ArraySet1(_1340_v85, 10)
				_1345_v90 = _nw250
				var _1354_v91 _dafny.Map
				_ = _1354_v91
				_1354_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1345_v90, _1337_v82)
				_1354_v91 = (_1354_v91).Update(_1345_v90, _dafny.CodePoint('u'))
			} else {
				var _1355___mcc_h2 _dafny.MultiSet = _source20.Get_().(D3_DC10).Cf21
				_ = _1355___mcc_h2
				var _1356_cf21 _dafny.MultiSet = _1355___mcc_h2
				_ = _1356_cf21
				var _1357_v92 _dafny.Array
				_ = _1357_v92
				var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
				_ = _nw251
				_1357_v92 = _nw251
				var _1358_v93 _dafny.Map
				_ = _1358_v93
				_1358_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _1359_v94 _dafny.Map
				_ = _1359_v94
				_1359_v94 = (_1358_v93).Update(p1, p1)
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1357_v92), 0))
				_ = _index260
				(_1357_v92).ArraySet1(_1359_v94, (_index260).Int())
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1357_v92), 0))
				_ = _index261
				(_1357_v92).ArraySet1(_1358_v93, (_index261).Int())
				var _1360_v95 _dafny.Sequence
				_ = _1360_v95
				_1360_v95 = _dafny.UnicodeSeqOfUtf8Bytes("tdhv")
				var _1361_v96 *C2
				_ = _1361_v96
				var _nw252 *C2 = New_C2_()
				_ = _nw252
				_nw252.Ctor__(_1360_v95)
				_1361_v96 = _nw252
				var _1362_v97 D8
				_ = _1362_v97
				_1362_v97 = Companion_D8_.Create_DC21_(_dafny.SeqOf(false, p1, !(p3)))
				_1362_v97 = _1362_v97
				var _1363_v98 *C1
				_ = _1363_v98
				var _nw253 *C1 = New_C1_()
				_ = _nw253
				_nw253.Ctor__()
				_1363_v98 = _nw253
			}
			if p3 {
				var _1364_v99 _dafny.Sequence
				_ = _1364_v99
				_1364_v99 = _dafny.UnicodeSeqOfUtf8Bytes("qvt")
				var _1365_v100 *C2
				_ = _1365_v100
				var _nw254 *C2 = New_C2_()
				_ = _nw254
				_nw254.Ctor__(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rlamqeya"), _1364_v99))
				_1365_v100 = _nw254
				var _1366_v101 _dafny.Array
				_ = _1366_v101
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw255
				_1366_v101 = _nw255
				var _1367_v102 _dafny.Sequence
				_ = _1367_v102
				_1367_v102 = _dafny.SeqOf(p1, p3)
				var _1368_v103 _dafny.Sequence
				_ = _1368_v103
				_1368_v103 = _dafny.SeqOf(_1365_v100.F33)
				var _1369_v104 _dafny.Sequence
				_ = _1369_v104
				_1369_v104 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_1368_v103).Cardinality()))
				var _1370_v105 _dafny.MultiSet
				_ = _1370_v105
				_1370_v105 = _dafny.MultiSetOf(p2)
				var _1371_v106 _dafny.Map
				_ = _1371_v106
				_1371_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1367_v102).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1367_v102).Cardinality()))).Uint32()).(bool), (_1369_v104).Select((Companion_Default___.SafeIndex((_1370_v105).Cardinality(), _dafny.IntOfUint32((_1369_v104).Cardinality()))).Uint32()).(_dafny.Int))
				var _1372_v107 _dafny.Map
				_ = _1372_v107
				_1372_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1371_v106).Equals(_1371_v106), _1366_v101)
				_1366_v101 = (func() _dafny.Array {
					if (_1372_v107).Contains((_1367_v102).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1367_v102).Cardinality()))).Uint32()).(bool)) {
						return (_1372_v107).Get((_1367_v102).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1367_v102).Cardinality()))).Uint32()).(bool)).(_dafny.Array)
					}
					return _1366_v101
				})()
				var _1373_v108 _dafny.Set
				_ = _1373_v108
				_1373_v108 = _dafny.SetOf(p1)
				var _1374_v109 D8
				_ = _1374_v109
				_1374_v109 = Companion_D8_.Create_DC23_(p1)
				var _1375_v110 D8
				_ = _1375_v110
				_1375_v110 = Companion_D8_.Create_DC24_(_1374_v109)
				var _1376_v111 _dafny.CodePoint
				_ = _1376_v111
				_1376_v111 = _dafny.CodePoint('r')
				var _rhs248 _dafny.Set = (_1373_v108).Difference(_1373_v108)
				_ = _rhs248
				var _rhs249 _dafny.CodePoint = _1376_v111
				_ = _rhs249
				var _rhs250 bool = p3
				_ = _rhs250
				var _rhs251 D8 = _1375_v110
				_ = _rhs251
				var _rhs252 bool = !(!(p1))
				_ = _rhs252
				var _lhs223 *GlobalState = globalState
				_ = _lhs223
				var _lhs224 *GlobalState = globalState
				_ = _lhs224
				var _lhs225 *GlobalState = globalState
				_ = _lhs225
				_1373_v108 = _rhs248
				_lhs223.F19 = _rhs249
				_lhs224.F13 = _rhs250
				_1375_v110 = _rhs251
				_lhs225.F7 = _rhs252
				(globalState).F19 = _1376_v111
				var _1377_v112 _dafny.Map
				_ = _1377_v112
				_1377_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1365_v100.F33, p2)
				(globalState).F10 = (_dafny.Zero).Minus(((_1377_v112).Cardinality()).Minus(p0))
			} else {
				var _1378_v113 _dafny.Array
				_ = _1378_v113
				var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
				_ = _nw256
				_1378_v113 = _nw256
				var _1379_v114 _dafny.Map
				_ = _1379_v114
				_1379_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index262
				(_1378_v113).ArraySet1(((_1379_v114).Cardinality()).Times(p2), (_index262).Int())
				var _1380_v115 _dafny.Map
				_ = _1380_v115
				_1380_v115 = _1379_v114
				var _1381_v116 _dafny.Sequence
				_ = _1381_v116
				_1381_v116 = _dafny.UnicodeSeqOfUtf8Bytes("esrblhn")
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index263
				(_1378_v113).ArraySet1(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfUint32((_1381_v116).Cardinality())), (_index263).Int())
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index264
				(_1378_v113).ArraySet1(p2, (_index264).Int())
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index265
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index266
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1378_v113), 0))
				_ = _index267
				var _rhs253 _dafny.Int = ((p0).Times(p2)).Times(_dafny.IntOfInt64(373))
				_ = _rhs253
				var _rhs254 bool = (_this).Fm5(p3, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ochwdrjm"), _1381_v116), p1, globalState)
				_ = _rhs254
				var _rhs255 _dafny.Map = _1380_v115
				_ = _rhs255
				var _rhs256 _dafny.Int = (_dafny.Zero).Minus(((_this).Fm6(globalState)).Times(((_1317_v70).Merge(_1317_v70)).Cardinality()))
				_ = _rhs256
				var _rhs257 _dafny.Int = _dafny.IntOfUint32((_1381_v116).Cardinality())
				_ = _rhs257
				var _lhs226 _dafny.Array = _1378_v113
				_ = _lhs226
				var _lhs227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_1378_v113), 0))
				_ = _lhs227
				var _lhs228 *GlobalState = globalState
				_ = _lhs228
				var _lhs229 _dafny.Array = _1378_v113
				_ = _lhs229
				var _lhs230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(626), _dafny.ArrayLen((_1378_v113), 0))
				_ = _lhs230
				var _lhs231 _dafny.Array = _1378_v113
				_ = _lhs231
				var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen((_1378_v113), 0))
				_ = _lhs232
				(_lhs226).ArraySet1(_rhs253, (_lhs227).Int())
				_lhs228.F22 = _rhs254
				_1380_v115 = _rhs255
				(_lhs229).ArraySet1(_rhs256, (_lhs230).Int())
				(_lhs231).ArraySet1(_rhs257, (_lhs232).Int())
				var _1382_v117 _dafny.Array
				_ = _1382_v117
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_33
				var _nw257 _dafny.Array
				_ = _nw257
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw257 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = (func(_1383_p1 bool) func(_dafny.Int) bool {
						return func(_1384_i18 _dafny.Int) bool {
							return _1383_p1
						}
					})(p1)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw257 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw257).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw257).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1382_v117 = _nw257
				var _1385_v118 _dafny.Array
				_ = _1385_v118
				var _nwElement0_60 _dafny.Array = _1382_v117
				_ = _nwElement0_60
				var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(6))
				_ = _nw258
				(_nw258).ArraySet1(_nwElement0_60, 0)
				(_nw258).ArraySet1(_1382_v117, 1)
				(_nw258).ArraySet1(_1382_v117, 2)
				(_nw258).ArraySet1(_1382_v117, 3)
				(_nw258).ArraySet1(_1382_v117, 4)
				(_nw258).ArraySet1((func() _dafny.Array {
					if p1 {
						return _1382_v117
					}
					return _1382_v117
				})(), 5)
				_1385_v118 = _nw258
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))
				_ = _index268
				(_1385_v118).ArraySet1(_1382_v117, (_index268).Int())
				var _1386_v119 *C2
				_ = _1386_v119
				var _nw259 *C2 = New_C2_()
				_ = _nw259
				_nw259.Ctor__(_1381_v116)
				_1386_v119 = _nw259
				var _1387_v120 _dafny.Set
				_ = _1387_v120
				_1387_v120 = _dafny.SetOf(p1)
				var _1388_v121 _dafny.Map
				_ = _1388_v121
				_1388_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _1386_v119)
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))
				_ = _index269
				var _rhs258 bool = !(_1387_v120).Equals((func() _dafny.Set {
					if p3 {
						return _1387_v120
					}
					return _1387_v120
				})())
				_ = _rhs258
				var _rhs259 _dafny.Array = _1382_v117
				_ = _rhs259
				var _rhs260 *C2 = (func() *C2 {
					if (_1388_v121).Contains(p3) {
						return (_1388_v121).Get(p3).(*C2)
					}
					return _1386_v119
				})()
				_ = _rhs260
				var _rhs261 bool = p3
				_ = _rhs261
				var _lhs233 *GlobalState = globalState
				_ = _lhs233
				var _lhs234 _dafny.Array = _1385_v118
				_ = _lhs234
				var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))
				_ = _lhs235
				var _lhs236 *GlobalState = globalState
				_ = _lhs236
				_lhs233.F13 = _rhs258
				(_lhs234).ArraySet1(_rhs259, (_lhs235).Int())
				_1386_v119 = _rhs260
				_lhs236.F22 = _rhs261
				(globalState).F10 = Companion_Default___.SafeModuloInt((p2).Times(p0), p2)
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1385_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))).Int()))
				_ = _arr0
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_dafny.ArrayCastTo((_1385_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))).Int()))), 0))
				_ = _index270
				(_arr0).ArraySet1(p3, (_index270).Int())
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1385_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))).Int()))
				_ = _arr1
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(716), _dafny.ArrayLen((_dafny.ArrayCastTo((_1385_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_1385_v118), 0))).Int()))), 0))
				_ = _index271
				(_arr1).ArraySet1(p3, (_index271).Int())
				(globalState).F10 = p0
			}
			var _rhs262 bool = p1
			_ = _rhs262
			var _rhs263 bool = p1
			_ = _rhs263
			var _lhs237 *GlobalState = globalState
			_ = _lhs237
			var _lhs238 *GlobalState = globalState
			_ = _lhs238
			_lhs237.F13 = _rhs262
			_lhs238.F13 = _rhs263
			var _1389_v123 _dafny.Sequence
			_ = _1389_v123
			_1389_v123 = _dafny.UnicodeSeqOfUtf8Bytes("idmp")
			var _1390_v124 _dafny.MultiSet
			_ = _1390_v124
			_1390_v124 = _dafny.MultiSetOf(_1389_v123)
			(globalState).F10 = (func() _dafny.Map {
				var _coll34 = _dafny.NewMapBuilder()
				_ = _coll34
				for _iter39 := _dafny.Iterate((_1390_v124).Elements()); ; {
					_compr_34, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _1391_v122 _dafny.Sequence
					_1391_v122 = interface{}(_compr_34).(_dafny.Sequence)
					if (_1390_v124).Contains(_1391_v122) {
						_coll34.Add(_1391_v122, p1)
					}
				}
				return _coll34.ToMap()
			}()).Cardinality()
		}
		var _1392_v125 _dafny.Array
		_ = _1392_v125
		var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw260
		_1392_v125 = _nw260
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1392_v125), 0))); ; {
			_guard_loop_4, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _1393_i19 _dafny.Int
			_1393_i19 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1393_i19).Sign() != -1) && ((_1393_i19).Cmp(_dafny.ArrayLen((_1392_v125), 0)) < 0)) {
				(_1392_v125).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("roywi"), _dafny.UnicodeSeqOfUtf8Bytes("fllqnoebg")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1394_i20 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))), (_1393_i19).Int())
			}
		}
	}
}
func (_this *C4) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 bool = false
		_ = r1
		r1 = p2
		var _1395_i0 _dafny.Int
		_ = _1395_i0
		_1395_i0 = _dafny.Zero
		{
			for p2 {
				{
					if (_1395_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1395_i0 = (_1395_i0).Plus(_dafny.One)
					var _1396_v0 _dafny.Sequence
					_ = _1396_v0
					_1396_v0 = _dafny.UnicodeSeqOfUtf8Bytes("oekpnmql")
					var _1397_v1 _dafny.Map
					_ = _1397_v1
					_1397_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(776))
					var _1398_v2 _dafny.Map
					_ = _1398_v2
					_1398_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.UnicodeSeqOfUtf8Bytes("i"))
					var _1399_v3 _dafny.Set
					_ = _1399_v3
					_1399_v3 = _dafny.SetOf(p2, !(p2), p2)
					var _1400_v4 _dafny.CodePoint
					_ = _1400_v4
					_1400_v4 = _dafny.CodePoint('k')
					var _rhs264 _dafny.Sequence = (func() _dafny.Sequence {
						if (_1398_v2).Contains(_dafny.IntOfInt64(-300)) {
							return (_1398_v2).Get(_dafny.IntOfInt64(-300)).(_dafny.Sequence)
						}
						return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("phjwav"), _1396_v0), (Companion_Default___.SafeIndex((_1399_v3).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("phjwav"), _1396_v0)).Cardinality()))).Uint32(), _1400_v4)
					})()
					_ = _rhs264
					var _rhs265 _dafny.Map = (func() _dafny.Map {
						if p2 {
							return _1397_v1
						}
						return _1397_v1
					})()
					_ = _rhs265
					_1396_v0 = _rhs264
					_1397_v1 = _rhs265
					var _1401_v5 _dafny.Set
					_ = _1401_v5
					_1401_v5 = _dafny.SetOf(_1396_v0, _1396_v0, _1396_v0)
					var _1402_v6 _dafny.Map
					_ = _1402_v6
					_1402_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).Plus((_dafny.Zero).Minus((_this).Fm2(globalState))), _1401_v5)
					var _1403_v7 _dafny.Map
					_ = _1403_v7
					_1403_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(990), (_this).Fm2(globalState))
					_1402_v6 = (_1402_v6).Update(_dafny.IntOfUint32((Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_1403_v7), _dafny.IntOfInt64(728), p1, globalState)).Cardinality()), (_1401_v5).Difference(_1401_v5))
					(globalState).F7 = p2
					var _1404_v8 _dafny.Array
					_ = _1404_v8
					var _out8 _dafny.Array
					_ = _out8
					_out8 = (_this).M2(globalState)
					_1404_v8 = _out8
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		(globalState).F10 = p1
		var _1405_v9 _dafny.MultiSet
		_ = _1405_v9
		_1405_v9 = _dafny.MultiSetOf(p2)
		var _1406_v10 D3
		_ = _1406_v10
		_1406_v10 = Companion_D3_.Create_DC10_(_1405_v9)
		if ((_1406_v10).Dtor_cf21()).IsDisjointFrom((_1405_v9).Union(_1405_v9)) {
			var _1407_v11 _dafny.Array
			_ = _1407_v11
			var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
			_ = _nw261
			_1407_v11 = _nw261
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1407_v11), 0))
			_ = _index272
			(_1407_v11).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xbbmhsxm"), (_index272).Int())
			var _1408_v12 _dafny.Sequence
			_ = _1408_v12
			_1408_v12 = _dafny.UnicodeSeqOfUtf8Bytes("oobix")
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_1407_v11), 0))
			_ = _index273
			(_1407_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1408_v12, _1408_v12), (_index273).Int())
			var _1409_v14 D0
			_ = _1409_v14
			_1409_v14 = Companion_D0_.Create_DC2_(p1, !(p2))
			var _1410_v15 _dafny.Array
			_ = _1410_v15
			var _nwElement0_61 bool = true
			_ = _nwElement0_61
			var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(19))
			_ = _nw262
			(_nw262).ArraySet1(_nwElement0_61, 0)
			(_nw262).ArraySet1(p2, 1)
			(_nw262).ArraySet1(p2, 2)
			(_nw262).ArraySet1(p2, 3)
			(_nw262).ArraySet1(false, 4)
			(_nw262).ArraySet1(p2, 5)
			(_nw262).ArraySet1(p2, 6)
			(_nw262).ArraySet1(p2, 7)
			(_nw262).ArraySet1(false, 8)
			(_nw262).ArraySet1(p2, 9)
			(_nw262).ArraySet1(p2, 10)
			(_nw262).ArraySet1(p2, 11)
			(_nw262).ArraySet1(p2, 12)
			(_nw262).ArraySet1(p2, 13)
			(_nw262).ArraySet1((_1409_v14).Dtor_cf6(), 14)
			(_nw262).ArraySet1(p2, 15)
			(_nw262).ArraySet1(p2, 16)
			(_nw262).ArraySet1(p2, 17)
			(_nw262).ArraySet1(p2, 18)
			_1410_v15 = _nw262
			var _1411_v16 D4
			_ = _1411_v16
			_1411_v16 = Companion_D4_.Create_DC14_(p2)
			var _1412_v17 D4
			_ = _1412_v17
			_1412_v17 = Companion_D4_.Create_DC15_(_1411_v16)
			var _1413_v18 D7
			_ = _1413_v18
			_1413_v18 = Companion_D7_.Create_DC20_(p0, _1410_v15, _1412_v17, p2, p1)
			var _1414_v19 _dafny.Sequence
			_ = _1414_v19
			_1414_v19 = _dafny.SeqOf(p2)
			var _1415_v20 _dafny.Map
			_ = _1415_v20
			_1415_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(-839))
			var _1416_v22 _dafny.Map
			_ = _1416_v22
			_1416_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			var _1417_v23 _dafny.Map
			_ = _1417_v23
			_1417_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(91))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}((func(_1418_v12 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1419_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1418_v12).Cardinality())
				}
			})(_1408_v12)))).Cardinality()), (func() _dafny.Map {
				var _coll35 = _dafny.NewMapBuilder()
				_ = _coll35
				for _iter41 := _dafny.Iterate((_1416_v22).Keys().Elements()); ; {
					_compr_35, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _1420_v21 _dafny.Int
					_1420_v21 = interface{}(_compr_35).(_dafny.Int)
					if (_1416_v22).Contains(_1420_v21) {
						_coll35.Add(Companion_Default___.SafeDivisionInt(_1420_v21, _dafny.IntOfInt64(-486)), p2)
					}
				}
				return _coll35.ToMap()
			}()).Cardinality())).Cardinality(), (_this).Fm5(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(884))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg75 _dafny.Int) interface{} {
					return coer75(arg75)
				}
			}((func(_1421_v12 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1422_i2 _dafny.Int) _dafny.Sequence {
					return _1421_v12
				}
			})(_1408_v12))), p2, globalState))
			var _1423_v24 _dafny.Array
			_ = _1423_v24
			var _nwElement0_62 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("axtltt")).Cardinality()))
			_ = _nwElement0_62
			var _nw263 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(22))
			_ = _nw263
			(_nw263).ArraySet1(_nwElement0_62, 0)
			(_nw263).ArraySet1(p1, 1)
			(_nw263).ArraySet1(p1, 2)
			(_nw263).ArraySet1(p1, 3)
			(_nw263).ArraySet1(p1, 4)
			(_nw263).ArraySet1(p1, 5)
			(_nw263).ArraySet1(p1, 6)
			(_nw263).ArraySet1(p1, 7)
			(_nw263).ArraySet1(p1, 8)
			(_nw263).ArraySet1(p1, 9)
			(_nw263).ArraySet1((func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter42 := _dafny.Iterate(((_1413_v18).Dtor_cf30()).Elements()); ; {
					_compr_36, _ok42 := _iter42()
					if !_ok42 {
						break
					}
					var _1424_v13 _dafny.Int
					_1424_v13 = interface{}(_compr_36).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((_1413_v18).Dtor_cf30(), _1424_v13) {
						_coll36.Add((_1424_v13).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf(p1))).Cardinality()), p1)
					}
				}
				return _coll36.ToMap()
			}()).Cardinality(), 10)
			(_nw263).ArraySet1(p1, 11)
			(_nw263).ArraySet1((_dafny.Zero).Minus(p1), 12)
			(_nw263).ArraySet1(p1, 13)
			(_nw263).ArraySet1(_dafny.IntOfUint32((_1414_v19).Cardinality()), 14)
			(_nw263).ArraySet1((_1415_v20).Cardinality(), 15)
			(_nw263).ArraySet1(p1, 16)
			(_nw263).ArraySet1(p1, 17)
			(_nw263).ArraySet1(_dafny.IntOfInt64(218), 18)
			(_nw263).ArraySet1((_1417_v23).Cardinality(), 19)
			(_nw263).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p2, !(false))).Cardinality()), 20)
			(_nw263).ArraySet1(p1, 21)
			_1423_v24 = _nw263
			var _1425_v25 _dafny.MultiSet
			_ = _1425_v25
			_1425_v25 = _dafny.MultiSetOf(_1423_v24)
			var _1426_v26 _dafny.MultiSet
			_ = _1426_v26
			_1426_v26 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(p1, (_1425_v25).Cardinality()), p1, p1, ((p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)).Minus(p1))
			_1426_v26 = (func() _dafny.MultiSet {
				if p2 {
					return _1426_v26
				}
				return _1426_v26
			})()
			var _1427_v27 _dafny.Set
			_ = _1427_v27
			_1427_v27 = _dafny.SetOf(p1)
			(globalState).F24 = _1427_v27
			(globalState).F7 = (p2) == (p2)
			_1414_v19 = _dafny.SeqOf(p2, p2)
		} else {
			var _1428_v28 _dafny.Array
			_ = _1428_v28
			var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
			_ = _nw264
			_1428_v28 = _nw264
			var _1429_v29 _dafny.Array
			_ = _1429_v29
			var _nw265 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw265
			_1429_v29 = _nw265
			var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1428_v28), 0))
			_ = _index274
			(_1428_v28).ArraySet1(_dafny.MultiSetOf(_1429_v29), (_index274).Int())
			var _1430_v30 _dafny.MultiSet
			_ = _1430_v30
			_1430_v30 = _dafny.MultiSetOf(_1429_v29, _1429_v29)
			var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_1428_v28), 0))
			_ = _index275
			(_1428_v28).ArraySet1(_1430_v30, (_index275).Int())
			(globalState).F19 = _dafny.CodePoint('b')
			_1429_v29 = _1429_v29
			var _1431_v31 _dafny.CodePoint
			_ = _1431_v31
			_1431_v31 = _dafny.CodePoint('d')
			var _1432_v32 _dafny.Sequence
			_ = _1432_v32
			_1432_v32 = _dafny.UnicodeSeqOfUtf8Bytes("kxbccg")
			var _1433_v33 _dafny.Map
			_ = _1433_v33
			_1433_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1431_v31, _1432_v32)
			(globalState).F10 = Companion_Default___.SafeModuloInt((_1433_v33).Cardinality(), p1)
			var _1434_v34 *C1
			_ = _1434_v34
			var _nw266 *C1 = New_C1_()
			_ = _nw266
			_nw266.Ctor__()
			_1434_v34 = _nw266
		}
		if p2 {
			var _1435_v35 _dafny.Sequence
			_ = _1435_v35
			_1435_v35 = _dafny.UnicodeSeqOfUtf8Bytes("g")
			var _1436_v36 *C2
			_ = _1436_v36
			var _nw267 *C2 = New_C2_()
			_ = _nw267
			_nw267.Ctor__(_dafny.Companion_Sequence_.Concatenate(_1435_v35, _dafny.UnicodeSeqOfUtf8Bytes("enwdan")))
			_1436_v36 = _nw267
			(globalState).F10 = p1
			var _1437_v37 D0
			_ = _1437_v37
			_1437_v37 = Companion_D0_.Create_DC2_(Companion_Default___.Fm0(p1, globalState), false)
			var _pat_let_tv40 = p1
			_ = _pat_let_tv40
			if (func(_pat_let42_0 D0) D0 {
				return func(_1438_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let43_0 _dafny.Int) D0 {
						return func(_1439_dt__update_hcf5_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC2_(_1439_dt__update_hcf5_h0, (_1438_dt__update__tmp_h0).Dtor_cf6())
						}(_pat_let43_0)
					}(_pat_let_tv40)
				}(_pat_let42_0)
			}(_1437_v37)).Dtor_cf6() {
				(globalState).F13 = p2
				var _1440_v40 _dafny.Map
				_ = _1440_v40
				_1440_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1441_v41 _dafny.Sequence
				_ = _1441_v41
				_1441_v41 = _dafny.SeqOf((func() bool {
					if (_1440_v40).Contains(p2) {
						return (_1440_v40).Get(p2).(bool)
					}
					return p2
				})())
				var _1442_v42 _dafny.Set
				_ = _1442_v42
				_1442_v42 = _dafny.SetOf(_1441_v41)
				var _rhs266 _dafny.Int = (func() _dafny.Int {
					if !(p2) {
						return (_this).Fm2(globalState)
					}
					return (func() _dafny.Map {
						var _coll37 = _dafny.NewMapBuilder()
						_ = _coll37
						for _iter43 := _dafny.Iterate((func() _dafny.Map {
							var _coll38 = _dafny.NewMapBuilder()
							_ = _coll38
							for _iter44 := _dafny.Iterate((_1442_v42).Elements()); ; {
								_compr_38, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _1443_v39 _dafny.Sequence
								_1443_v39 = interface{}(_compr_38).(_dafny.Sequence)
								if (_1442_v42).Contains(_1443_v39) {
									_coll38.Add(_1443_v39, p2)
								}
							}
							return _coll38.ToMap()
						}()).Keys().Elements()); ; {
							_compr_37, _ok43 := _iter43()
							if !_ok43 {
								break
							}
							var _1444_v38 _dafny.Sequence
							_1444_v38 = interface{}(_compr_37).(_dafny.Sequence)
							if (func() _dafny.Map {
								var _coll39 = _dafny.NewMapBuilder()
								_ = _coll39
								for _iter45 := _dafny.Iterate((_1442_v42).Elements()); ; {
									_compr_39, _ok45 := _iter45()
									if !_ok45 {
										break
									}
									var _1443_v39 _dafny.Sequence
									_1443_v39 = interface{}(_compr_39).(_dafny.Sequence)
									if (_1442_v42).Contains(_1443_v39) {
										_coll39.Add(_1443_v39, p2)
									}
								}
								return _coll39.ToMap()
							}()).Contains(_1444_v38) {
								_coll37.Add(_1444_v38, (_1405_v9).Cardinality())
							}
						}
						return _coll37.ToMap()
					}()).Cardinality()
				})()
				_ = _rhs266
				var _rhs267 _dafny.MultiSet = _1405_v9
				_ = _rhs267
				var _lhs239 *GlobalState = globalState
				_ = _lhs239
				_lhs239.F10 = _rhs266
				_1405_v9 = _rhs267
				var _1445_v43 _dafny.CodePoint
				_ = _1445_v43
				_1445_v43 = _dafny.CodePoint('c')
				var _1446_v44 _dafny.Map
				_ = _1446_v44
				_1446_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _1447_v45 _dafny.Map
				_ = _1447_v45
				_1447_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_v43, (_1446_v44).Cardinality())
				var _1448_v46 _dafny.Map
				_ = _1448_v46
				_1448_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex((_1447_v45).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), p1))
				(_1436_v36).M0(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1448_v46).Contains(p1) {
						return (_1448_v46).Get(p1).(_dafny.Sequence)
					}
					return p0
				})()).Cardinality()), p2, p1, !(p2), globalState)
				(globalState).F13 = ((_dafny.Zero).Minus(p1)).Cmp(p1) < 0
				var _1449_v47 _dafny.Map
				_ = _1449_v47
				_1449_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)
				_1449_v47 = (_1449_v47).Update((p1).Cmp(p1) <= 0, p1)
			} else {
				(globalState).F13 = p2
				(globalState).F10 = (p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
				var _1450_v48 _dafny.Array
				_ = _1450_v48
				var _nw268 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw268
				_1450_v48 = _nw268
				(globalState).F10 = (_dafny.MultiSetOf(_1450_v48, _1450_v48)).Cardinality()
				var _1451_v49 _dafny.CodePoint
				_ = _1451_v49
				_1451_v49 = _dafny.CodePoint('h')
				(globalState).F19 = _1451_v49
				var _1452_v50 _dafny.Map
				_ = _1452_v50
				_1452_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1436_v36.F33)
				var _1453_v51 _dafny.Sequence
				_ = _1453_v51
				_1453_v51 = _dafny.SeqOf(_1436_v36.F33, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_1452_v50).Contains(p2) {
						return (_1452_v50).Get(p2).(_dafny.Sequence)
					}
					return _1436_v36.F33
				})(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1452_v50).Contains(p2) {
						return (_1452_v50).Get(p2).(_dafny.Sequence)
					}
					return _1436_v36.F33
				})()).Cardinality()))).Uint32(), _1451_v49), _dafny.Companion_Sequence_.Concatenate(_1435_v35, _1435_v35))
				_1453_v51 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1453_v51, _1453_v51), _1453_v51)
			}
			(_1436_v36).M0(_dafny.IntOfUint32((_1435_v35).Cardinality()), p2, (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, p1)), true, globalState)
			(globalState).F10 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((p1).Minus(p1), p1))
		} else {
			(globalState).F13 = p2
			(globalState).F10 = ((_1405_v9).Cardinality()).Plus((_dafny.IntOfInt64(146)).Plus(p1))
			var _1454_v52 _dafny.Map
			_ = _1454_v52
			_1454_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_this).Fm6(globalState), p1), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, p0)).Cardinality()).Times(p1))
			_1454_v52 = (_1454_v52).Update(p1, _dafny.IntOfInt64(841))
			(globalState).F10 = _dafny.IntOfInt64(622)
			var _1455_v53 _dafny.Map
			_ = _1455_v53
			_1455_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality()))
			var _1456_v54 _dafny.MultiSet
			_ = _1456_v54
			_1456_v54 = _dafny.MultiSetOf(_dafny.IntOfInt64(844))
			var _1457_v55 _dafny.Sequence
			_ = _1457_v55
			_1457_v55 = _dafny.SeqOf(_1456_v54)
			var _1458_v56 _dafny.Sequence
			_ = _1458_v56
			_1458_v56 = _dafny.UnicodeSeqOfUtf8Bytes("gyebnjvkd")
			var _1459_v57 _dafny.Map
			_ = _1459_v57
			_1459_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			var _1460_v58 _dafny.Array
			_ = _1460_v58
			var _nwElement0_63 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(p1), p2)).Cardinality()
			_ = _nwElement0_63
			var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(26))
			_ = _nw269
			(_nw269).ArraySet1(_nwElement0_63, 0)
			(_nw269).ArraySet1(p1, 1)
			(_nw269).ArraySet1(p1, 2)
			(_nw269).ArraySet1(p1, 3)
			(_nw269).ArraySet1(p1, 4)
			(_nw269).ArraySet1(p1, 5)
			(_nw269).ArraySet1((_1455_v53).Cardinality(), 6)
			(_nw269).ArraySet1(p1, 7)
			(_nw269).ArraySet1(_dafny.IntOfUint32((_1457_v55).Cardinality()), 8)
			(_nw269).ArraySet1(p1, 9)
			(_nw269).ArraySet1(p1, 10)
			(_nw269).ArraySet1(p1, 11)
			(_nw269).ArraySet1(p1, 12)
			(_nw269).ArraySet1(p1, 13)
			(_nw269).ArraySet1((func() _dafny.Int {
				if (_1456_v54).Contains(p1) {
					return (_1456_v54).Multiplicity(p1)
				}
				return p1
			})(), 14)
			(_nw269).ArraySet1(((_1456_v54).Update(p1, Companion_Default___.Abs(p1))).Cardinality(), 15)
			(_nw269).ArraySet1(_dafny.IntOfUint32((_1458_v56).Cardinality()), 16)
			(_nw269).ArraySet1(p1, 17)
			(_nw269).ArraySet1(p1, 18)
			(_nw269).ArraySet1(_dafny.IntOfUint32((_1458_v56).Cardinality()), 19)
			(_nw269).ArraySet1(p1, 20)
			(_nw269).ArraySet1((_1459_v57).Cardinality(), 21)
			(_nw269).ArraySet1(p1, 22)
			(_nw269).ArraySet1(p1, 23)
			(_nw269).ArraySet1(p1, 24)
			(_nw269).ArraySet1(p1, 25)
			_1460_v58 = _nw269
			var _1461_v59 _dafny.Map
			_ = _1461_v59
			_1461_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1460_v58, p2)
			(globalState).F10 = (_dafny.IntOfInt64(447)).Plus((_1461_v59).Cardinality())
		}
		var _1462_v60 _dafny.Map
		_ = _1462_v60
		_1462_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p1), _dafny.IntOfInt64(-889)), (_dafny.Zero).Minus((func() _dafny.Int {
			if p2 {
				return p1
			}
			return p1
		})()))
		_1462_v60 = (_1462_v60).Update(p1, _dafny.IntOfInt64(565))
		var _1463_v61 _dafny.Sequence
		_ = _1463_v61
		_1463_v61 = _dafny.SeqOf(p2)
		var _1464_v62 _dafny.Map
		_ = _1464_v62
		_1464_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _1465_v63 _dafny.MultiSet
		_ = _1465_v63
		_1465_v63 = _dafny.MultiSetOf(p1)
		var _1466_v64 _dafny.Map
		_ = _1466_v64
		_1466_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_v63, p0)
		var _1467_v65 D1
		_ = _1467_v65
		_1467_v65 = Companion_D1_.Create_DC5_(p1, p2, (func() bool {
			if (_1464_v62).Contains(false) {
				return (_1464_v62).Get(false).(bool)
			}
			return !(p2)
		})(), _1466_v64)
		var _1468_v66 _dafny.Array
		_ = _1468_v66
		var _nw270 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
		_ = _nw270
		_1468_v66 = _nw270
		var _1469_v67 D0
		_ = _1469_v67
		_1469_v67 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1463_v61, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1463_v61).Cardinality()))).Uint32(), p2)).Cardinality()), p1, (func() _dafny.Int {
			if (_1467_v65).Dtor_cf10() {
				return (_dafny.Zero).Minus(p1)
			}
			return _dafny.IntOfUint32((_1463_v61).Cardinality())
		})(), _1468_v66)
		r0 = _1469_v67
		r1 = (_1465_v63).IsSubsetOf((func() _dafny.MultiSet {
			if true {
				return _1465_v63
			}
			return _1465_v63
		})())
		return r0, r1
	}
}
func (_this *C4) M2(globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _1470_v0 _dafny.Array
		_ = _1470_v0
		var _nw271 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
		_ = _nw271
		_1470_v0 = _nw271
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1470_v0), 0))); ; {
			_guard_loop_5, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _1471_i0 _dafny.Int
			_1471_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1471_i0).Sign() != -1) && ((_1471_i0).Cmp(_dafny.ArrayLen((_1470_v0), 0)) < 0)) {
				(_1470_v0).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("gkgylxbum"), _dafny.UnicodeSeqOfUtf8Bytes("o")), _dafny.UnicodeSeqOfUtf8Bytes("txvt")), (_1471_i0).Int())
			}
		}
		var _1472_v1 _dafny.Int
		_ = _1472_v1
		_1472_v1 = _dafny.IntOfInt64(109)
		var _1473_v2 bool
		_ = _1473_v2
		_1473_v2 = false
		var _1474_v3 D0
		_ = _1474_v3
		_1474_v3 = Companion_D0_.Create_DC2_(_1472_v1, _1473_v2)
		var _1475_v4 _dafny.Sequence
		_ = _1475_v4
		_1475_v4 = _dafny.SeqOf(_1474_v3, _1474_v3, _1474_v3)
		var _1476_v6 _dafny.MultiSet
		_ = _1476_v6
		_1476_v6 = _dafny.MultiSetOf(func() _dafny.Map {
			var _coll40 = _dafny.NewMapBuilder()
			_ = _coll40
			for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-15), _dafny.IntOfInt64(-440))); ; {
				_compr_40, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _1477_v5 _dafny.Int
				_1477_v5 = interface{}(_compr_40).(_dafny.Int)
				if ((_dafny.IntOfInt64(-15)).Cmp(_1477_v5) <= 0) && ((_1477_v5).Cmp(_dafny.IntOfInt64(-440)) < 0) {
					_coll40.Add((_1477_v5).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lem")).Cardinality()), _1472_v1)).Cardinality()), _dafny.IntOfInt64(518))
				}
			}
			return _coll40.ToMap()
		}())
		var _1478_v8 D0
		_ = _1478_v8
		_1478_v8 = Companion_D0_.Create_DC3_((_1475_v4).Select((Companion_Default___.SafeIndex((func() _dafny.Set {
			var _coll41 = _dafny.NewBuilder()
			_ = _coll41
			for _iter48 := _dafny.Iterate((_1476_v6).Elements()); ; {
				_compr_41, _ok48 := _iter48()
				if !_ok48 {
					break
				}
				var _1479_v7 _dafny.Map
				_1479_v7 = interface{}(_compr_41).(_dafny.Map)
				if (_1476_v6).Contains(_1479_v7) {
					_coll41.Add(_1479_v7)
				}
			}
			return _coll41.ToSet()
		}()).Cardinality(), _dafny.IntOfUint32((_1475_v4).Cardinality()))).Uint32()).(D0))
		var _1480_v9 _dafny.Map
		_ = _1480_v9
		_1480_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1472_v1, _1478_v8)
		var _1481_v10 _dafny.Sequence
		_ = _1481_v10
		_1481_v10 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _1482_v11 D3
		_ = _1482_v11
		_1482_v11 = Companion_D3_.Create_DC11_(_1472_v1)
		var _1483_v12 _dafny.MultiSet
		_ = _1483_v12
		_1483_v12 = _dafny.MultiSetOf(_1472_v1)
		var _1484_v13 _dafny.Map
		_ = _1484_v13
		_1484_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1483_v12, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(808))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg76 _dafny.Int) interface{} {
				return coer76(arg76)
			}
		}(func(_1485_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-553)
		})))
		var _1486_v14 D1
		_ = _1486_v14
		_1486_v14 = Companion_D1_.Create_DC5_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("csvrecknv")).Cardinality()), _1473_v2, _1473_v2, _1484_v13)
		var _1487_v15 D1
		_ = _1487_v15
		_1487_v15 = Companion_D1_.Create_DC5_(_1472_v1, false, _1473_v2, (_1486_v14).Dtor_cf12())
		var _1488_v16 _dafny.Array
		_ = _1488_v16
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_34
		var _nw272 _dafny.Array
		_ = _nw272
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw272 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Int = (func(_1489_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1490_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1490_i2, _dafny.IntOfUint32((_1489_v10).Cardinality()))
				}
			})(_1481_v10)
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw272 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw272).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw272).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1488_v16 = _nw272
		var _1491_v17 _dafny.Map
		_ = _1491_v17
		_1491_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1487_v15, _1488_v16)
		var _1492_v18 _dafny.Array
		_ = _1492_v18
		var _nwElement0_64 _dafny.Int = _dafny.IntOfInt64(-908)
		_ = _nwElement0_64
		var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(12))
		_ = _nw273
		(_nw273).ArraySet1(_nwElement0_64, 0)
		(_nw273).ArraySet1(_1472_v1, 1)
		(_nw273).ArraySet1((_1480_v9).Cardinality(), 2)
		(_nw273).ArraySet1((_dafny.IntOfUint32((_1481_v10).Cardinality())).Plus((_dafny.Zero).Minus(_1472_v1)), 3)
		(_nw273).ArraySet1(_1472_v1, 4)
		(_nw273).ArraySet1((_1482_v11).Dtor_cf22(), 5)
		(_nw273).ArraySet1((_dafny.Zero).Minus(_1472_v1), 6)
		(_nw273).ArraySet1(_1472_v1, 7)
		(_nw273).ArraySet1(_1472_v1, 8)
		(_nw273).ArraySet1((_dafny.Zero).Minus(((_1491_v17).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25(_1473_v2, _1473_v2, _1472_v1, _1473_v2, globalState), _1488_v16))).Cardinality()), 9)
		(_nw273).ArraySet1(_1472_v1, 10)
		(_nw273).ArraySet1(_1472_v1, 11)
		_1492_v18 = _nw273
		var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1492_v18), 0))
		_ = _index276
		(_1492_v18).ArraySet1(_1472_v1, (_index276).Int())
		var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1492_v18), 0))
		_ = _index277
		(_1492_v18).ArraySet1(_1472_v1, (_index277).Int())
		var _1493_v19 _dafny.Map
		_ = _1493_v19
		_1493_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1492_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1492_v18), 0))).Int()).(_dafny.Int)).Cmp(_1472_v1) <= 0, _dafny.SetOf(_1488_v16, _1488_v16, _1492_v18))
		var _1494_v20 _dafny.Set
		_ = _1494_v20
		_1494_v20 = _dafny.SetOf(_1492_v18)
		_1493_v19 = (_1493_v19).Update(_1473_v2, _1494_v20)
		var _1495_v21 D2
		_ = _1495_v21
		_1495_v21 = Companion_D2_.Create_DC8_(_1473_v2, _1473_v2)
		var _pat_let_tv41 = _1472_v1
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1472_v1
		_ = _pat_let_tv42
		var _pat_let_tv43 = _1473_v2
		_ = _pat_let_tv43
		_1472_v1 = func(_source21 D2) _dafny.Int {
			if _source21.Is_DC8() {
				var _1496___mcc_h0 bool = _source21.Get_().(D2_DC8).Cf18
				_ = _1496___mcc_h0
				var _1497___mcc_h1 bool = _source21.Get_().(D2_DC8).Cf19
				_ = _1497___mcc_h1
				var _1498_cf19 bool = _1497___mcc_h1
				_ = _1498_cf19
				var _1499_cf18 bool = _1496___mcc_h0
				_ = _1499_cf18
				return _pat_let_tv41
			} else if _source21.Is_DC7() {
				var _1500___mcc_h2 _dafny.Map = _source21.Get_().(D2_DC7).Cf17
				_ = _1500___mcc_h2
				var _1501_cf17 _dafny.Map = _1500___mcc_h2
				_ = _1501_cf17
				return _pat_let_tv42
			} else {
				var _1502___mcc_h3 D2 = _source21.Get_().(D2_DC9).Cf20
				_ = _1502___mcc_h3
				var _1503_cf20 D2 = _1502___mcc_h3
				_ = _1503_cf20
				return _dafny.IntOfInt64(866)
			}
		}(func(_pat_let44_0 D2) D2 {
			return func(_1504_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let45_0 bool) D2 {
					return func(_1505_dt__update_hcf18_h0 bool) D2 {
						return Companion_D2_.Create_DC8_(_1505_dt__update_hcf18_h0, (_1504_dt__update__tmp_h0).Dtor_cf19())
					}(_pat_let45_0)
				}(_pat_let_tv43)
			}(_pat_let44_0)
		}(_1495_v21))
		var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1492_v18), 0))
		_ = _index278
		(_1492_v18).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_1492_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1492_v18), 0))).Int()).(_dafny.Int))), (_index278).Int())
		var _1506_v22 D1
		_ = _1506_v22
		var _1507_v23 _dafny.Int
		_ = _1507_v23
		var _1508_v24 _dafny.Sequence
		_ = _1508_v24
		var _out9 D1
		_ = _out9
		var _out10 _dafny.Int
		_ = _out10
		var _out11 _dafny.Sequence
		_ = _out11
		_out9, _out10, _out11 = (_this).M3((func() bool {
			if _1473_v2 {
				return _1473_v2
			}
			return _1473_v2
		})(), _1470_v0, globalState)
		_1506_v22 = _out9
		_1507_v23 = _out10
		_1508_v24 = _out11
		r0 = _1488_v16
		return r0
	}
}
func (_this *C4) M3(p0 bool, p1 _dafny.Array, globalState *GlobalState) (D1, _dafny.Int, _dafny.Sequence) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var _1509_v0 _dafny.Sequence
		_ = _1509_v0
		_1509_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ri")
		var _1510_v1 _dafny.Int
		_ = _1510_v1
		_1510_v1 = _dafny.IntOfInt64(212)
		var _1511_v2 _dafny.Map
		_ = _1511_v2
		_1511_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v0, _dafny.SetOf((_dafny.Zero).Minus(_1510_v1), _1510_v1))
		var _1512_v3 _dafny.Set
		_ = _1512_v3
		_1512_v3 = _dafny.SetOf(_1510_v1)
		_1511_v2 = (_1511_v2).Update(_1509_v0, _1512_v3)
		(globalState).F13 = p0
		var _1513_v4 _dafny.Sequence
		_ = _1513_v4
		_1513_v4 = _dafny.SeqOf(p0)
		var _1514_v5 _dafny.Sequence
		_ = _1514_v5
		_1514_v5 = _dafny.SeqOf((_1513_v4).Select((Companion_Default___.SafeIndex(_1510_v1, _dafny.IntOfUint32((_1513_v4).Cardinality()))).Uint32()).(bool))
		var _1515_v6 _dafny.Map
		_ = _1515_v6
		_1515_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1513_v4).Cardinality()), _1510_v1)
		_1515_v6 = (_1515_v6).Update(_1510_v1, _1510_v1)
		var _1516_v7 *C0
		_ = _1516_v7
		var _nw274 *C0 = New_C0_()
		_ = _nw274
		_nw274.Ctor__(p0)
		_1516_v7 = _nw274
		_1516_v7 = _1516_v7
		(globalState).F13 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1509_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(668))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg77 _dafny.Int) interface{} {
				return coer77(arg77)
			}
		}(func(_1517_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))), Companion_Default___.Fm11(_1510_v1, globalState))
		var _1518_v8 *C3
		_ = _1518_v8
		var _nw275 *C3 = New_C3_()
		_ = _nw275
		_nw275.Ctor__()
		_1518_v8 = _nw275
		r0 = Companion_Default___.Fm29(globalState)
		var _1519_v9 _dafny.Map
		_ = _1519_v9
		_1519_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1520_v10 _dafny.Set
		_ = _1520_v10
		_1520_v10 = _dafny.SetOf(Companion_Default___.Fm30((_1516_v7).F32(), globalState), _1519_v9)
		r1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_1510_v1, (_1520_v10).Cardinality())), (_this).Fm6(globalState))
		r2 = Companion_Default___.Fm35(_1510_v1, p0, (func() bool {
			if (_1516_v7).F32() {
				return (_1516_v7).F32()
			}
			return p0
		})(), globalState)
		return r0, r1, r2
	}
}
func (_this *C4) F31() _dafny.Array {
	{
		return _this._f31
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	F30  _dafny.Int
	_f29 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this.F30 = _dafny.Zero
	_this._f29 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f29 _dafny.Array, f30 _dafny.Int) {
	{
		(_this)._f29 = f29
		(_this).F30 = f30
	}
}
func (_this *C5) Fm1(p0 D0, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return (Companion_D1_.Create_DC4_(func() _dafny.Map {
			var _coll42 = _dafny.NewMapBuilder()
			_ = _coll42
			for _iter49 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, true, false), _this.F30)).Keys().Elements()); ; {
				_compr_42, _ok49 := _iter49()
				if !_ok49 {
					break
				}
				var _1521_v0 _dafny.MultiSet
				_1521_v0 = interface{}(_compr_42).(_dafny.MultiSet)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(false, true, false), _this.F30)).Contains(_1521_v0) {
					_coll42.Add(_1521_v0, true)
				}
			}
			return _coll42.ToMap()
		}())).Dtor_cf8()
	}
}
func (_this *C5) Fm2(globalState *GlobalState) _dafny.Int {
	{
		return _this.F30
	}
}
func (_this *C5) Fm3(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 D0, globalState *GlobalState) _dafny.Int {
	{
		return (_this.F30).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, !(false))).Cardinality())
	}
}
func (_this *C5) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C5) M0(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		if p1 {
			var _1522_v0 _dafny.Array
			_ = _1522_v0
			var _nw276 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(11))
			_ = _nw276
			_1522_v0 = _nw276
			var _1523_v1 _dafny.Array
			_ = _1523_v1
			var _len0_35 _dafny.Int = _dafny.One
			_ = _len0_35
			var _nw277 _dafny.Array
			_ = _nw277
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw277 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) bool = (func(_1524_p3 bool) func(_dafny.Int) bool {
					return func(_1525_i0 _dafny.Int) bool {
						return _1524_p3
					}
				})(p3)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw277 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw277).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw277).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1523_v1 = _nw277
			var _1526_v2 _dafny.MultiSet
			_ = _1526_v2
			_1526_v2 = _dafny.MultiSetOf(_1523_v1)
			var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1522_v0), 0))
			_ = _index279
			(_1522_v0).ArraySet1(_1526_v2, (_index279).Int())
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1522_v0), 0))
			_ = _index280
			var _rhs268 _dafny.MultiSet = (_1526_v2).Difference(_1526_v2)
			_ = _rhs268
			var _rhs269 bool = (p1) == (p1)
			_ = _rhs269
			var _lhs240 _dafny.Array = _1522_v0
			_ = _lhs240
			var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1522_v0), 0))
			_ = _lhs241
			var _lhs242 *GlobalState = globalState
			_ = _lhs242
			(_lhs240).ArraySet1(_rhs268, (_lhs241).Int())
			_lhs242.F22 = _rhs269
			var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1523_v1), 0))
			_ = _index281
			(_1523_v1).ArraySet1(p3, (_index281).Int())
			var _1527_v3 _dafny.Sequence
			_ = _1527_v3
			_1527_v3 = _dafny.SeqOf(true)
			var _1528_v4 _dafny.Set
			_ = _1528_v4
			_1528_v4 = _dafny.SetOf(p1, p3, (_1527_v3).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1527_v3).Cardinality()))).Uint32()).(bool))
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1523_v1), 0))
			_ = _index282
			(_1523_v1).ArraySet1(((_dafny.SetOf(false)).Difference(_dafny.SetOf(!(p1), p1, p1, true, p1))).IsProperSubsetOf(_1528_v4), (_index282).Int())
			var _1529_v5 _dafny.Array
			_ = _1529_v5
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_36
			var _nw278 _dafny.Array
			_ = _nw278
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw278 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Sequence = (func(_1530_p2 _dafny.Int, _1531_p0 _dafny.Int, _1532_p1 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1533_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer78 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg78 _dafny.Int) interface{} {
								return coer78(arg78)
							}
						}((func(_1534_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1535_i2 _dafny.Int) _dafny.Int {
								return _1534_p2
							}
						})(_1530_p2))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(604))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg79 _dafny.Int) interface{} {
								return coer79(arg79)
							}
						}((func(_1536_p0 _dafny.Int, _1537_p1 bool) func(_dafny.Int) _dafny.Int {
							return func(_1538_i3 _dafny.Int) _dafny.Int {
								return (Companion_D0_.Create_DC2_(_1536_p0, _1537_p1)).Dtor_cf5()
							}
						})(_1531_p0, _1532_p1))))
					}
				})(p2, p0, p1)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw278 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw278).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw278).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1529_v5 = _nw278
			var _1539_v6 _dafny.MultiSet
			_ = _1539_v6
			_1539_v6 = _dafny.MultiSetOf(p0, _dafny.IntOfInt64(259))
			var _rhs270 _dafny.Array = _1529_v5
			_ = _rhs270
			var _rhs271 _dafny.MultiSet = _1539_v6
			_ = _rhs271
			var _rhs272 bool = ((_1523_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1523_v1), 0))).Int()).(bool)) && (!((_1523_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1523_v1), 0))).Int()).(bool)))
			_ = _rhs272
			var _lhs243 *GlobalState = globalState
			_ = _lhs243
			_1529_v5 = _rhs270
			_1539_v6 = _rhs271
			_lhs243.F7 = _rhs272
			var _1540_v7 _dafny.Sequence
			_ = _1540_v7
			_1540_v7 = _dafny.UnicodeSeqOfUtf8Bytes("spjomng")
			var _1541_v8 D0
			_ = _1541_v8
			_1541_v8 = Companion_D0_.Create_DC0_(_dafny.Companion_Sequence_.Concatenate(_1540_v7, _1540_v7))
			_1541_v8 = _1541_v8
			var _1542_v9 _dafny.Map
			_ = _1542_v9
			_1542_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1528_v4).IsProperSubsetOf(_1528_v4), p3)
			var _rhs273 bool = (func() bool {
				if (_1542_v9).Contains(!(p3)) {
					return (_1542_v9).Get(!(p3)).(bool)
				}
				return p1
			})()
			_ = _rhs273
			var _rhs274 _dafny.Int = _this.F30
			_ = _rhs274
			var _rhs275 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((p2).Minus(p0)))
			_ = _rhs275
			var _rhs276 _dafny.Sequence = _1540_v7
			_ = _rhs276
			var _rhs277 _dafny.Int = (_this).Fm2(globalState)
			_ = _rhs277
			var _lhs244 *GlobalState = globalState
			_ = _lhs244
			var _lhs245 *GlobalState = globalState
			_ = _lhs245
			var _lhs246 *C5 = _this
			_ = _lhs246
			var _lhs247 *C5 = _this
			_ = _lhs247
			_lhs244.F7 = _rhs273
			_lhs245.F10 = _rhs274
			_lhs246.F30 = _rhs275
			_1540_v7 = _rhs276
			_lhs247.F30 = _rhs277
		} else {
			(globalState).F7 = p1
			var _1543_v10 _dafny.Array
			_ = _1543_v10
			var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw279
			_1543_v10 = _nw279
			var _1544_v11 T1
			_ = _1544_v11
			var _nw280 *C4 = New_C4_()
			_ = _nw280
			_nw280.Ctor__(_1543_v10)
			_1544_v11 = _nw280
			var _1545_v12 _dafny.Map
			_ = _1545_v12
			_1545_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1544_v11, p0)
			var _1546_v13 _dafny.Array
			_ = _1546_v13
			var _nwElement0_65 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1544_v11, p0)
			_ = _nwElement0_65
			var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(3))
			_ = _nw281
			(_nw281).ArraySet1(_nwElement0_65, 0)
			(_nw281).ArraySet1(_1545_v12, 1)
			(_nw281).ArraySet1(_1545_v12, 2)
			_1546_v13 = _nw281
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_1546_v13), 0))
			_ = _index283
			(_1546_v13).ArraySet1(_1545_v12, (_index283).Int())
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(587), _dafny.ArrayLen((_1546_v13), 0))
			_ = _index284
			(_1546_v13).ArraySet1((_1545_v12).Merge(_1545_v12), (_index284).Int())
			var _1547_v14 _dafny.Array
			_ = _1547_v14
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_37
			var _nw282 _dafny.Array
			_ = _nw282
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw282 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Sequence = func(_1548_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("qj")
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw282 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw282).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw282).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1547_v14 = _nw282
			var _1549_v15 _dafny.Sequence
			_ = _1549_v15
			_1549_v15 = _dafny.UnicodeSeqOfUtf8Bytes("vkarjbt")
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_1547_v14), 0))
			_ = _index285
			(_1547_v14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1549_v15, _1549_v15), (_index285).Int())
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_1547_v14), 0))
			_ = _index286
			(_1547_v14).ArraySet1(Companion_Default___.Fm35(_dafny.IntOfInt64(603), p1, (_this.F30).Cmp(_dafny.IntOfInt64(162)) != 0, globalState), (_index286).Int())
			var _1550_v17 _dafny.MultiSet
			_ = _1550_v17
			_1550_v17 = _dafny.MultiSetOf(func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(469), _dafny.IntOfInt64(878))); ; {
					_compr_43, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _1551_v16 _dafny.Int
					_1551_v16 = interface{}(_compr_43).(_dafny.Int)
					if ((_dafny.IntOfInt64(469)).Cmp(_1551_v16) <= 0) && ((_1551_v16).Cmp(_dafny.IntOfInt64(878)) < 0) {
						_coll43.Add((_1551_v16).Plus(_this.F30), p1)
					}
				}
				return _coll43.ToMap()
			}())
			var _1552_v18 _dafny.Sequence
			_ = _1552_v18
			_1552_v18 = _dafny.SeqOf(_1550_v17)
			var _1553_v19 _dafny.Map
			_ = _1553_v19
			_1553_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F30)
			var _1554_v20 _dafny.Array
			_ = _1554_v20
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_38
			var _nw283 _dafny.Array
			_ = _nw283
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw283 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) bool = (func(_1555_p1 bool) func(_dafny.Int) bool {
					return func(_1556_i5 _dafny.Int) bool {
						return _1555_p1
					}
				})(p1)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw283 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw283).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw283).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1554_v20 = _nw283
			var _1557_v21 _dafny.Map
			_ = _1557_v21
			_1557_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1553_v19, _1554_v20)
			var _1558_v22 _dafny.Map
			_ = _1558_v22
			_1558_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _1559_v23 _dafny.Map
			_ = _1559_v23
			_1559_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1558_v22, p3)
			var _1560_v24 _dafny.Map
			_ = _1560_v24
			_1560_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
			var _1561_v25 D2
			_ = _1561_v25
			_1561_v25 = Companion_D2_.Create_DC7_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_1553_v19).Contains(_this.F30) {
					return (_1553_v19).Get(_this.F30).(_dafny.Int)
				}
				return p2
			})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())))
			var _1562_v26 _dafny.Map
			_ = _1562_v26
			_1562_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1561_v25, p0)
			var _1563_v27 _dafny.Sequence
			_ = _1563_v27
			_1563_v27 = _dafny.SeqOf(_1562_v26, _1562_v26, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC7_(_1553_v19), p0), _1562_v26, _1562_v26)
			var _1564_v28 _dafny.Map
			_ = _1564_v28
			_1564_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if p1 {
					return (_1552_v18).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1557_v21).Cardinality(), p0)).Cardinality(), _dafny.IntOfUint32((_1552_v18).Cardinality()))).Uint32()).(_dafny.MultiSet)
				}
				return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1558_v22).Contains(p1) {
						return (_1558_v22).Get(p1).(_dafny.Int)
					}
					return (_1559_v23).Cardinality()
				})(), p1), _1560_v24)).Update(_1560_v24, Companion_Default___.Abs(p2))
			})(), _1563_v27)
			_1564_v28 = (_1564_v28).Update(_1550_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(11))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}((func(_1565_v26 _dafny.Map) func(_dafny.Int) _dafny.Map {
				return func(_1566_i6 _dafny.Int) _dafny.Map {
					return _1565_v26
				}
			})(_1562_v26))))
			var _1567_v29 _dafny.CodePoint
			_ = _1567_v29
			_1567_v29 = _dafny.CodePoint('c')
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_1547_v14), 0))
			_ = _index287
			(_1547_v14).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_1553_v19), (_dafny.Zero).Minus((func() _dafny.Int {
				if (_1558_v22).Contains(p1) {
					return (_1558_v22).Get(p1).(_dafny.Int)
				}
				return p2
			})()), (_dafny.Zero).Minus(p0), globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_1553_v19), (_dafny.Zero).Minus((func() _dafny.Int {
				if (_1558_v22).Contains(p1) {
					return (_1558_v22).Get(p1).(_dafny.Int)
				}
				return p2
			})()), (_dafny.Zero).Minus(p0), globalState)).Cardinality()))).Uint32(), _1567_v29), (_index287).Int())
		}
		(globalState).F10 = _this.F30
		var _1568_v30 _dafny.Array
		_ = _1568_v30
		var _nw284 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
		_ = _nw284
		_1568_v30 = _nw284
		var _1569_v31 _dafny.CodePoint
		_ = _1569_v31
		_1569_v31 = _dafny.CodePoint('a')
		var _1570_v32 _dafny.MultiSet
		_ = _1570_v32
		_1570_v32 = _dafny.MultiSetOf(_1569_v31, _1569_v31, _1569_v31)
		var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
		_ = _index288
		(_1568_v30).ArraySet1((_1570_v32).Cardinality(), (_index288).Int())
		var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_1568_v30), 0))
		_ = _index289
		(_1568_v30).ArraySet1(Companion_Default___.SafeModuloInt(p2, _this.F30), (_index289).Int())
		var _1571_v33 _dafny.Map
		_ = _1571_v33
		_1571_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p3), _dafny.IntOfInt64(554))
		var _1572_v34 _dafny.MultiSet
		_ = _1572_v34
		_1572_v34 = _dafny.MultiSetOf((_1571_v33).Cardinality())
		var _1573_v36 D3
		_ = _1573_v36
		_1573_v36 = Companion_D3_.Create_DC11_(p0)
		var _1574_v37 _dafny.Sequence
		_ = _1574_v37
		_1574_v37 = _dafny.SeqOf((_1572_v34).Cardinality(), (func() _dafny.Map {
			var _coll44 = _dafny.NewMapBuilder()
			_ = _coll44
			for _iter51 := _dafny.Iterate((Companion_Default___.Fm14(p1, p2, (_1573_v36).Dtor_cf22(), globalState)).Elements()); ; {
				_compr_44, _ok51 := _iter51()
				if !_ok51 {
					break
				}
				var _1575_v35 _dafny.Int
				_1575_v35 = interface{}(_compr_44).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm14(p1, p2, (_1573_v36).Dtor_cf22(), globalState), _1575_v35) {
					_coll44.Add(Companion_Default___.SafeModuloInt(_1575_v35, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1572_v34)).Cardinality()), _1569_v31)
				}
			}
			return _coll44.ToMap()
		}()).Cardinality(), (_this.F30).Plus(Companion_Default___.Fm0(p0, globalState)))
		var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
		_ = _index290
		var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_1568_v30), 0))
		_ = _index291
		var _rhs278 _dafny.Int = (_dafny.Zero).Minus((p0).Minus(p2))
		_ = _rhs278
		var _rhs279 _dafny.Int = p2
		_ = _rhs279
		var _rhs280 _dafny.Int = (_1574_v37).Select((Companion_Default___.SafeIndex((p0).Times(p0), _dafny.IntOfUint32((_1574_v37).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs280
		var _rhs281 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(-948)), _1574_v37)).Cardinality())).Times((p2).Plus(_this.F30))
		_ = _rhs281
		var _rhs282 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqxo")).Cardinality())
		_ = _rhs282
		var _lhs248 _dafny.Array = _1568_v30
		_ = _lhs248
		var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
		_ = _lhs249
		var _lhs250 *C5 = _this
		_ = _lhs250
		var _lhs251 _dafny.Array = _1568_v30
		_ = _lhs251
		var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_1568_v30), 0))
		_ = _lhs252
		var _lhs253 *C5 = _this
		_ = _lhs253
		var _lhs254 *GlobalState = globalState
		_ = _lhs254
		(_lhs248).ArraySet1(_rhs278, (_lhs249).Int())
		_lhs250.F30 = _rhs279
		(_lhs251).ArraySet1(_rhs280, (_lhs252).Int())
		_lhs253.F30 = _rhs281
		_lhs254.F10 = _rhs282
		var _1576_v38 D0
		_ = _1576_v38
		_1576_v38 = Companion_D0_.Create_DC2_((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int), !(p3))
		_1576_v38 = _1576_v38
		var _1577_v39 *C3
		_ = _1577_v39
		var _nw285 *C3 = New_C3_()
		_ = _nw285
		_nw285.Ctor__()
		_1577_v39 = _nw285
		var _1578_v40 _dafny.MultiSet
		_ = _1578_v40
		_1578_v40 = _dafny.MultiSetOf(_1577_v39)
		var _1579_v41 _dafny.Map
		_ = _1579_v41
		_1579_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1578_v40).Cardinality(), (_1577_v39).Fm8(p2, _dafny.IntOfInt64(-44), p0, globalState))
		_1579_v41 = (_1579_v41).Update(_dafny.IntOfInt64(744), (_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int))
		var _1580_i7 _dafny.Int
		_ = _1580_i7
		_1580_i7 = _dafny.Zero
		{
			for p3 {
				{
					if (_1580_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1580_i7 = (_1580_i7).Plus(_dafny.One)
					if p3 {
						var _1581_v42 _dafny.Array
						_ = _1581_v42
						var _len0_39 _dafny.Int = _dafny.IntOfInt64(11)
						_ = _len0_39
						var _nw286 _dafny.Array
						_ = _nw286
						if _len0_39.Cmp(_dafny.Zero) == 0 {
							_nw286 = _dafny.NewArray(_len0_39)
						} else {
							var _init39 func(_dafny.Int) bool = (func(_1582_p3 bool, _1583_p1 bool) func(_dafny.Int) bool {
								return func(_1584_i8 _dafny.Int) bool {
									return (func() bool {
										if _1582_p3 {
											return _1583_p1
										}
										return _1582_p3
									})()
								}
							})(p3, p1)
							_ = _init39
							var _element0_39 = _init39(_dafny.Zero)
							_ = _element0_39
							_nw286 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
							(_nw286).ArraySet1(_element0_39, 0)
							var _nativeLen0_39 = (_len0_39).Int()
							_ = _nativeLen0_39
							for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
								(_nw286).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
							}
						}
						_1581_v42 = _nw286
						var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))
						_ = _index292
						(_1581_v42).ArraySet1(p1, (_index292).Int())
						var _1585_v43 _dafny.Sequence
						_ = _1585_v43
						_1585_v43 = _dafny.UnicodeSeqOfUtf8Bytes("fxrsbmfcb")
						var _1586_v44 D4
						_ = _1586_v44
						_1586_v44 = Companion_D4_.Create_DC14_(p1)
						var _1587_v45 D4
						_ = _1587_v45
						_1587_v45 = Companion_D4_.Create_DC15_(_1586_v44)
						var _1588_v46 D7
						_ = _1588_v46
						_1588_v46 = Companion_D7_.Create_DC20_(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F30), (Companion_Default___.SafeIndex((_1579_v41).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F30)).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1585_v43).Cardinality())), _1581_v42, _1587_v45, p3, _this.F30)
						var _1589_v47 D0
						_ = _1589_v47
						_1589_v47 = Companion_D0_.Create_DC0_(_1585_v43)
						var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))
						_ = _index293
						var _rhs283 bool = (_1588_v46).Dtor_cf33()
						_ = _rhs283
						var _rhs284 bool = p3
						_ = _rhs284
						var _rhs285 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1585_v43, (_1589_v47).Dtor_cf0()), _dafny.Companion_Sequence_.Concatenate(_1585_v43, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(936))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg81 _dafny.Int) interface{} {
								return coer81(arg81)
							}
						}((func(_1590_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1591_i9 _dafny.Int) _dafny.CodePoint {
								return _1590_v31
							}
						})(_1569_v31)))))
						_ = _rhs285
						var _rhs286 _dafny.Int = p2
						_ = _rhs286
						var _lhs255 _dafny.Array = _1581_v42
						_ = _lhs255
						var _lhs256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))
						_ = _lhs256
						var _lhs257 *GlobalState = globalState
						_ = _lhs257
						var _lhs258 *GlobalState = globalState
						_ = _lhs258
						var _lhs259 *GlobalState = globalState
						_ = _lhs259
						(_lhs255).ArraySet1(_rhs283, (_lhs256).Int())
						_lhs257.F7 = _rhs284
						_lhs258.F13 = _rhs285
						_lhs259.F10 = _rhs286
						var _1592_v49 _dafny.Sequence
						_ = _1592_v49
						_1592_v49 = _dafny.SeqOf((_1581_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))).Int()).(bool))
						var _1593_v50 _dafny.Sequence
						_ = _1593_v50
						_1593_v50 = _dafny.SeqOf(_1592_v49)
						var _pat_let_tv44 = _1574_v37
						_ = _pat_let_tv44
						var _pat_let_tv45 = _1593_v50
						_ = _pat_let_tv45
						var _pat_let_tv46 = _1593_v50
						_ = _pat_let_tv46
						var _pat_let_tv47 = _1581_v42
						_ = _pat_let_tv47
						var _pat_let_tv48 = _1581_v42
						_ = _pat_let_tv48
						var _pat_let_tv49 = _1568_v30
						_ = _pat_let_tv49
						var _pat_let_tv50 = _1568_v30
						_ = _pat_let_tv50
						_1588_v46 = func(_pat_let46_0 D7) D7 {
							return func(_1594_dt__update__tmp_h0 D7) D7 {
								return func(_pat_let47_0 _dafny.Sequence) D7 {
									return func(_1595_dt__update_hcf30_h0 _dafny.Sequence) D7 {
										return func(_pat_let48_0 _dafny.Int) D7 {
											return func(_1597_dt__update_hcf34_h0 _dafny.Int) D7 {
												return Companion_D7_.Create_DC20_(_1595_dt__update_hcf30_h0, (_1594_dt__update__tmp_h0).Dtor_cf31(), (_1594_dt__update__tmp_h0).Dtor_cf32(), (_1594_dt__update__tmp_h0).Dtor_cf33(), _1597_dt__update_hcf34_h0)
											}(_pat_let48_0)
										}(Companion_Default___.SafeDivisionInt((func() _dafny.Map {
											var _coll45 = _dafny.NewMapBuilder()
											_ = _coll45
											for _iter52 := _dafny.Iterate((_dafny.MultiSetFromSeq(_pat_let_tv45)).Elements()); ; {
												_compr_45, _ok52 := _iter52()
												if !_ok52 {
													break
												}
												var _1596_v48 _dafny.Sequence
												_1596_v48 = interface{}(_compr_45).(_dafny.Sequence)
												if (_dafny.MultiSetFromSeq(_pat_let_tv46)).Contains(_1596_v48) {
													_coll45.Add(_1596_v48, (_pat_let_tv48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_pat_let_tv47), 0))).Int()).(bool))
												}
											}
											return _coll45.ToMap()
										}()).Cardinality(), (_pat_let_tv50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_pat_let_tv49), 0))).Int()).(_dafny.Int)))
									}(_pat_let47_0)
								}(_pat_let_tv44)
							}(_pat_let46_0)
						}(_1588_v46)
						var _1598_v51 _dafny.Array
						_ = _1598_v51
						var _len0_40 _dafny.Int = _dafny.IntOfInt64(25)
						_ = _len0_40
						var _nw287 _dafny.Array
						_ = _nw287
						if _len0_40.Cmp(_dafny.Zero) == 0 {
							_nw287 = _dafny.NewArray(_len0_40)
						} else {
							var _init40 func(_dafny.Int) _dafny.Sequence = (func(_1599_v49 _dafny.Sequence, _1600_v42 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
								return func(_1601_i10 _dafny.Int) _dafny.Sequence {
									return _dafny.Companion_Sequence_.Concatenate((Companion_D8_.Create_DC21_(_1599_v49)).Dtor_cf35(), _dafny.SeqOf((_1600_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1600_v42), 0))).Int()).(bool)))
								}
							})(_1592_v49, _1581_v42)
							_ = _init40
							var _element0_40 = _init40(_dafny.Zero)
							_ = _element0_40
							_nw287 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
							(_nw287).ArraySet1(_element0_40, 0)
							var _nativeLen0_40 = (_len0_40).Int()
							_ = _nativeLen0_40
							for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
								(_nw287).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
							}
						}
						_1598_v51 = _nw287
						_1598_v51 = _1598_v51
						var _1602_v52 _dafny.Array
						_ = _1602_v52
						var _nw288 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(14))
						_ = _nw288
						_1602_v52 = _nw288
						var _1603_v53 _dafny.Map
						_ = _1603_v53
						_1603_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1581_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))).Int()).(bool), _1602_v52)
						var _1604_v54 _dafny.Sequence
						_ = _1604_v54
						_1604_v54 = _dafny.SeqOf((func() _dafny.Array {
							if (_1603_v53).Contains(false) {
								return (_1603_v53).Get(false).(_dafny.Array)
							}
							return _1602_v52
						})(), _1602_v52, _1602_v52, _1602_v52)
						var _1605_v55 _dafny.Array
						_ = _1605_v55
						var _nwElement0_66 _dafny.Array = _1602_v52
						_ = _nwElement0_66
						var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(6))
						_ = _nw289
						(_nw289).ArraySet1(_nwElement0_66, 0)
						(_nw289).ArraySet1((_1604_v54).Select((Companion_Default___.SafeIndex((_1574_v37).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1574_v37).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1604_v54).Cardinality()))).Uint32()).(_dafny.Array), 1)
						(_nw289).ArraySet1(_1602_v52, 2)
						(_nw289).ArraySet1(_1602_v52, 3)
						(_nw289).ArraySet1(_1602_v52, 4)
						(_nw289).ArraySet1(_1602_v52, 5)
						_1605_v55 = _nw289
						var _1606_v56 T1
						_ = _1606_v56
						var _nw290 *C4 = New_C4_()
						_ = _nw290
						_nw290.Ctor__(_1605_v55)
						_1606_v56 = _nw290
						_1606_v56 = _1606_v56
						var _1607_v57 _dafny.Array
						_ = _1607_v57
						var _nw291 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
						_ = _nw291
						_1607_v57 = _nw291
						var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1607_v57), 0))
						_ = _index294
						(_1607_v57).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(427))).Uint32(), func(coer82 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg82 _dafny.Int) interface{} {
								return coer82(arg82)
							}
						}((func(_1608_v43 _dafny.Sequence, _1609_v31 _dafny.CodePoint, _1610_p2 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
							return func(_1611_i11 _dafny.Int) _dafny.CodePoint {
								return (_1608_v43).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1609_v31, _1610_p2)).Cardinality(), _dafny.IntOfUint32((_1608_v43).Cardinality()))).Uint32()).(_dafny.CodePoint)
							}
						})(_1585_v43, _1569_v31, p2))), (_index294).Int())
						var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1607_v57), 0))
						_ = _index295
						var _rhs287 _dafny.Sequence = Companion_Default___.Fm34(p1, _1569_v31, (func() _dafny.Int {
							if (_1581_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1581_v42), 0))).Int()).(bool) {
								return (_1577_v39).Fm2(globalState)
							}
							return (_1571_v33).Cardinality()
						})(), globalState)
						_ = _rhs287
						var _rhs288 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-404))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg83 _dafny.Int) interface{} {
								return coer83(arg83)
							}
						}((func(_1612_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1613_i12 _dafny.Int) _dafny.CodePoint {
								return _1612_v31
							}
						})(_1569_v31))), _1585_v43)
						_ = _rhs288
						var _lhs260 _dafny.Array = _1607_v57
						_ = _lhs260
						var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_1607_v57), 0))
						_ = _lhs261
						_1592_v49 = _rhs287
						(_lhs260).ArraySet1(_rhs288, (_lhs261).Int())
					} else {
						var _1614_v58 _dafny.Array
						_ = _1614_v58
						var _len0_41 _dafny.Int = _dafny.One
						_ = _len0_41
						var _nw292 _dafny.Array
						_ = _nw292
						if _len0_41.Cmp(_dafny.Zero) == 0 {
							_nw292 = _dafny.NewArray(_len0_41)
						} else {
							var _init41 func(_dafny.Int) D3 = (func(_1615_v36 D3) func(_dafny.Int) D3 {
								return func(_1616_i13 _dafny.Int) D3 {
									return _1615_v36
								}
							})(_1573_v36)
							_ = _init41
							var _element0_41 = _init41(_dafny.Zero)
							_ = _element0_41
							_nw292 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
							(_nw292).ArraySet1(_element0_41, 0)
							var _nativeLen0_41 = (_len0_41).Int()
							_ = _nativeLen0_41
							for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
								(_nw292).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
							}
						}
						_1614_v58 = _nw292
						var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_1614_v58), 0))
						_ = _index296
						(_1614_v58).ArraySet1(_1573_v36, (_index296).Int())
						var _pat_let_tv51 = _1574_v37
						_ = _pat_let_tv51
						var _pat_let_tv52 = _1574_v37
						_ = _pat_let_tv52
						var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_1614_v58), 0))
						_ = _index297
						(_1614_v58).ArraySet1(func(_pat_let49_0 D3) D3 {
							return func(_1617_dt__update__tmp_h1 D3) D3 {
								return func(_pat_let50_0 _dafny.Int) D3 {
									return func(_1618_dt__update_hcf22_h0 _dafny.Int) D3 {
										return Companion_D3_.Create_DC11_(_1618_dt__update_hcf22_h0)
									}(_pat_let50_0)
								}((_pat_let_tv51).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(323), _dafny.IntOfUint32((_pat_let_tv52).Cardinality()))).Uint32()).(_dafny.Int))
							}(_pat_let49_0)
						}(Companion_Default___.Fm40((_1572_v34).Cardinality(), false, p3, globalState)), (_index297).Int())
						var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
						_ = _index298
						var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
						_ = _index299
						var _rhs289 _dafny.Int = (_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)
						_ = _rhs289
						var _rhs290 _dafny.Int = p0
						_ = _rhs290
						var _rhs291 _dafny.Int = _this.F30
						_ = _rhs291
						var _lhs262 _dafny.Array = _1568_v30
						_ = _lhs262
						var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
						_ = _lhs263
						var _lhs264 _dafny.Array = _1568_v30
						_ = _lhs264
						var _lhs265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
						_ = _lhs265
						var _lhs266 *GlobalState = globalState
						_ = _lhs266
						(_lhs262).ArraySet1(_rhs289, (_lhs263).Int())
						(_lhs264).ArraySet1(_rhs290, (_lhs265).Int())
						_lhs266.F10 = _rhs291
						(globalState).F7 = p3
						var _1619_v59 *C1
						_ = _1619_v59
						var _nw293 *C1 = New_C1_()
						_ = _nw293
						_nw293.Ctor__()
						_1619_v59 = _nw293
						var _1620_v60 _dafny.Map
						_ = _1620_v60
						_1620_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1619_v59, p1)
						var _1621_v61 _dafny.Map
						_ = _1621_v61
						_1621_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1569_v31, p1)
						_1620_v60 = (_1620_v60).Update(_1619_v59, (func() bool {
							if (_1621_v61).Contains(_1569_v31) {
								return (_1621_v61).Get(_1569_v31).(bool)
							}
							return true
						})())
						var _1622_v62 _dafny.Map
						_ = _1622_v62
						_1622_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(478), _1571_v33)
						var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))
						_ = _index300
						(_1568_v30).ArraySet1(((func() _dafny.Int {
							if p1 {
								return (_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)
							}
							return (_1622_v62).Cardinality()
						})()).Times(p2), (_index300).Int())
					}
					var _1623_v63 _dafny.Sequence
					_ = _1623_v63
					_1623_v63 = _dafny.UnicodeSeqOfUtf8Bytes("v")
					var _rhs292 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg84 _dafny.Int) interface{} {
							return coer84(arg84)
						}
					}((func(_1624_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1625_i14 _dafny.Int) _dafny.CodePoint {
							return _1624_v31
						}
					})(_1569_v31))), (Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(46))).Uint32(), func(coer85 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg85 _dafny.Int) interface{} {
							return coer85(arg85)
						}
					}((func(_1626_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1627_i14 _dafny.Int) _dafny.CodePoint {
							return _1626_v31
						}
					})(_1569_v31)))).Cardinality()))).Uint32(), _1569_v31)
					_ = _rhs292
					var _rhs293 _dafny.Int = ((_1572_v34).Difference(_dafny.MultiSetFromSeq(_1574_v37))).Cardinality()
					_ = _rhs293
					var _lhs267 *GlobalState = globalState
					_ = _lhs267
					_1623_v63 = _rhs292
					_lhs267.F10 = _rhs293
					if p3 {
						var _1628_v64 _dafny.Array
						_ = _1628_v64
						var _nw294 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.One)
						_ = _nw294
						_1628_v64 = _nw294
						var _1629_v65 _dafny.Array
						_ = _1629_v65
						var _nwElement0_67 _dafny.Array = _1628_v64
						_ = _nwElement0_67
						var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(10))
						_ = _nw295
						(_nw295).ArraySet1(_nwElement0_67, 0)
						(_nw295).ArraySet1(_1628_v64, 1)
						(_nw295).ArraySet1(_1628_v64, 2)
						(_nw295).ArraySet1(_1628_v64, 3)
						(_nw295).ArraySet1(_1628_v64, 4)
						(_nw295).ArraySet1(_1628_v64, 5)
						(_nw295).ArraySet1(_1628_v64, 6)
						(_nw295).ArraySet1(_1628_v64, 7)
						(_nw295).ArraySet1(_1628_v64, 8)
						(_nw295).ArraySet1(_1628_v64, 9)
						_1629_v65 = _nw295
						var _1630_v66 *C4
						_ = _1630_v66
						var _nw296 *C4 = New_C4_()
						_ = _nw296
						_nw296.Ctor__(_1629_v65)
						_1630_v66 = _nw296
						var _1631_v67 _dafny.MultiSet
						_ = _1631_v67
						_1631_v67 = _dafny.MultiSetOf(_1568_v30, _1568_v30)
						(globalState).F10 = (func() _dafny.Int {
							if (_1631_v67).Contains(_1568_v30) {
								return (_1631_v67).Multiplicity(_1568_v30)
							}
							return _this.F30
						})()
						(globalState).F10 = _dafny.IntOfInt64(-41)
						(globalState).F13 = !(((p0).Minus(p0)).Cmp(_this.F30) < 0)
						var _1632_v68 _dafny.Sequence
						_ = _1632_v68
						_1632_v68 = _dafny.SeqOf(p3)
						_1571_v33 = (_1571_v33).Update(!(p3) || (p3), (_dafny.IntOfUint32((_1632_v68).Cardinality())).Minus(_this.F30))
					} else {
						var _1633_v69 _dafny.Map
						_ = _1633_v69
						_1633_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1568_v30, (_this).Fm4(_1623_v63, _this.F30, globalState))
						var _1634_v70 _dafny.Sequence
						_ = _1634_v70
						_1634_v70 = _dafny.SeqOf(_1633_v69, _1633_v69, _1633_v69, _1633_v69, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1568_v30, p3))
						var _1635_v71 _dafny.Sequence
						_ = _1635_v71
						_1635_v71 = _dafny.SeqOf(_1568_v30, _1568_v30)
						_1633_v69 = ((_1634_v70).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(881))).Cardinality()), _dafny.IntOfUint32((_1634_v70).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1635_v71).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1635_v71).Cardinality()))).Uint32()).(_dafny.Array), true))
						var _1636_v72 _dafny.Sequence
						_ = _1636_v72
						_1636_v72 = _dafny.SeqOf(Companion_D4_.Create_DC14_(p1))
						var _1637_v73 _dafny.Map
						_ = _1637_v73
						_1637_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1636_v72, _this.F30)
						var _rhs294 _dafny.Map = _1637_v73
						_ = _rhs294
						var _rhs295 _dafny.Sequence = _1623_v63
						_ = _rhs295
						_1637_v73 = _rhs294
						_1623_v63 = _rhs295
						var _1638_v74 _dafny.Array
						_ = _1638_v74
						var _nw297 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(22))
						_ = _nw297
						_1638_v74 = _nw297
						var _1639_v75 _dafny.Map
						_ = _1639_v75
						_1639_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)
						var _1640_v76 _dafny.Sequence
						_ = _1640_v76
						_1640_v76 = _dafny.SeqOf(_1639_v75)
						var _1641_v77 _dafny.Array
						_ = _1641_v77
						var _nwElement0_68 bool = p3
						_ = _nwElement0_68
						var _nw298 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(26))
						_ = _nw298
						(_nw298).ArraySet1(_nwElement0_68, 0)
						(_nw298).ArraySet1(true, 1)
						(_nw298).ArraySet1(!(p1), 2)
						(_nw298).ArraySet1(p1, 3)
						(_nw298).ArraySet1(false, 4)
						(_nw298).ArraySet1(p3, 5)
						(_nw298).ArraySet1(p1, 6)
						(_nw298).ArraySet1(!(true), 7)
						(_nw298).ArraySet1(p3, 8)
						(_nw298).ArraySet1(p3, 9)
						(_nw298).ArraySet1(p3, 10)
						(_nw298).ArraySet1(p3, 11)
						(_nw298).ArraySet1((_this).Fm4(_1623_v63, _dafny.IntOfUint32((_1640_v76).Cardinality()), globalState), 12)
						(_nw298).ArraySet1(p1, 13)
						(_nw298).ArraySet1(p1, 14)
						(_nw298).ArraySet1(p1, 15)
						(_nw298).ArraySet1(p3, 16)
						(_nw298).ArraySet1(!(false), 17)
						(_nw298).ArraySet1(false, 18)
						(_nw298).ArraySet1(p3, 19)
						(_nw298).ArraySet1(p1, 20)
						(_nw298).ArraySet1((_this).Fm4(_dafny.UnicodeSeqOfUtf8Bytes("whc"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(138))).Cardinality(), globalState), 21)
						(_nw298).ArraySet1(p1, 22)
						(_nw298).ArraySet1(p3, 23)
						(_nw298).ArraySet1(p3, 24)
						(_nw298).ArraySet1(p3, 25)
						_1641_v77 = _nw298
						var _1642_v78 D4
						_ = _1642_v78
						_1642_v78 = Companion_D4_.Create_DC14_(p3)
						var _1643_v79 D4
						_ = _1643_v79
						_1643_v79 = Companion_D4_.Create_DC15_(_1642_v78)
						var _1644_v80 D7
						_ = _1644_v80
						_1644_v80 = Companion_D7_.Create_DC20_(_1574_v37, _1641_v77, _1643_v79, p3, p0)
						var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_1638_v74), 0))
						_ = _index301
						(_1638_v74).ArraySet1(_1644_v80, (_index301).Int())
						var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(516), _dafny.ArrayLen((_1638_v74), 0))
						_ = _index302
						(_1638_v74).ArraySet1(Companion_D7_.Create_DC20_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(312))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg86 _dafny.Int) interface{} {
								return coer86(arg86)
							}
						}(func(_1645_i15 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("flxdys")).Cardinality())
						})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-914))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg87 _dafny.Int) interface{} {
								return coer87(arg87)
							}
						}((func(_1646_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1647_i16 _dafny.Int) _dafny.Int {
								return _1646_p0
							}
						})(p0)))), _1641_v77, _1643_v79, p3, p0), (_index302).Int())
						var _1648_v81 _dafny.Map
						_ = _1648_v81
						_1648_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1623_v63, Companion_Default___.Fm10(_1579_v41, _1569_v31, p3, globalState)), _dafny.SeqOf((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)))
						_1648_v81 = (_1648_v81).Update(_1623_v63, _1574_v37)
						var _1649_v82 _dafny.Map
						_ = _1649_v82
						_1649_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(680), _1641_v77)
						_1641_v77 = (func() _dafny.Array {
							if (_1649_v82).Contains((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)) {
								return (_1649_v82).Get((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)).(_dafny.Array)
							}
							return _1641_v77
						})()
					}
					var _1650_v83 _dafny.Array
					_ = _1650_v83
					var _nw299 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
					_ = _nw299
					_1650_v83 = _nw299
					var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1650_v83), 0))
					_ = _index303
					(_1650_v83).ArraySet1(((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)).Cmp((_1568_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_1568_v30), 0))).Int()).(_dafny.Int)) >= 0, (_index303).Int())
					var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1650_v83), 0))
					_ = _index304
					(_1650_v83).ArraySet1((_this).Fm4(_1623_v63, _dafny.IntOfUint32((_1623_v63).Cardinality()), globalState), (_index304).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
	}
}
func (_this *C5) M1(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, globalState *GlobalState) (D0, bool) {
	{
		var r0 D0 = Companion_D0_.Default()
		_ = r0
		var r1 bool = false
		_ = r1
		var _1651_v0 _dafny.Map
		_ = _1651_v0
		_1651_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, p1)
		var _hi7 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, (_1651_v0).Cardinality())
		_ = _hi7
		for _1652_i0 := p1; _1652_i0.Cmp(_hi7) < 0; _1652_i0 = _1652_i0.Plus(_dafny.One) {
			var _1653_v1 _dafny.Sequence
			_ = _1653_v1
			_1653_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ctawxugn")
			var _1654_v2 _dafny.Sequence
			_ = _1654_v2
			_1654_v2 = _dafny.SeqOf(p2, p2)
			var _1655_v3 _dafny.Map
			_ = _1655_v3
			_1655_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(_1653_v1, (_dafny.MultiSetFromSeq(_1654_v2)).Cardinality(), globalState), _dafny.IntOfInt64(-828))
			var _1656_v4 _dafny.MultiSet
			_ = _1656_v4
			_1656_v4 = _dafny.MultiSetOf(_dafny.IntOfInt64(311))
			var _1657_v5 D2
			_ = _1657_v5
			_1657_v5 = Companion_D2_.Create_DC7_(_1651_v0)
			var _1658_v6 *C2
			_ = _1658_v6
			var _nw300 *C2 = New_C2_()
			_ = _nw300
			_nw300.Ctor__(_1653_v1)
			_1658_v6 = _nw300
			var _1659_v7 _dafny.MultiSet
			_ = _1659_v7
			_1659_v7 = _dafny.MultiSetOf(_1658_v6)
			var _1660_v8 _dafny.Map
			_ = _1660_v8
			_1660_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _1659_v7)
			var _1661_v9 _dafny.Array
			_ = _1661_v9
			var _nwElement0_69 bool = !(true)
			_ = _nwElement0_69
			var _nw301 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(27))
			_ = _nw301
			(_nw301).ArraySet1(_nwElement0_69, 0)
			(_nw301).ArraySet1(p2, 1)
			(_nw301).ArraySet1((_1652_i0).Cmp(p1) == 0, 2)
			(_nw301).ArraySet1(!(p2), 3)
			(_nw301).ArraySet1(p2, 4)
			(_nw301).ArraySet1(p2, 5)
			(_nw301).ArraySet1((_this).Fm4(_1653_v1, _this.F30, globalState), 6)
			(_nw301).ArraySet1(p2, 7)
			(_nw301).ArraySet1(p2, 8)
			(_nw301).ArraySet1(p2, 9)
			(_nw301).ArraySet1(((_1651_v0).Cardinality()).Cmp(_dafny.IntOfInt64(232)) < 0, 10)
			(_nw301).ArraySet1((func() bool {
				if !(p2) {
					return p2
				}
				return p2
			})(), 11)
			(_nw301).ArraySet1(!(p2) || (p2), 12)
			(_nw301).ArraySet1(p2, 13)
			(_nw301).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)).Cardinality()).Cmp((_1655_v3).Cardinality()) >= 0, 14)
			(_nw301).ArraySet1((_dafny.IntOfInt64(801)).Cmp((_dafny.MultiSetOf(_dafny.IntOfUint32((_1654_v2).Cardinality()), _dafny.IntOfInt64(657), p1, _dafny.IntOfInt64(56))).Cardinality()) >= 0, 15)
			(_nw301).ArraySet1((_1656_v4).IsProperSubsetOf(_dafny.MultiSetOf(_1652_i0, (_dafny.Zero).Minus(_1652_i0), _dafny.IntOfUint32((_1653_v1).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm23(_1657_v5, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_1653_v1).Cardinality()))).Cardinality()), _this.F30, globalState)).Cardinality()), _dafny.IntOfUint32((_1654_v2).Cardinality()))), 16)
			(_nw301).ArraySet1(p2, 17)
			(_nw301).ArraySet1((_this.F30).Cmp(p1) > 0, 18)
			(_nw301).ArraySet1(p2, 19)
			(_nw301).ArraySet1(p2, 20)
			(_nw301).ArraySet1(p2, 21)
			(_nw301).ArraySet1(p2, 22)
			(_nw301).ArraySet1(p2, 23)
			(_nw301).ArraySet1((_1660_v8).Contains(_1652_i0), 24)
			(_nw301).ArraySet1(((func() _dafny.Int {
				if (_1655_v3).Contains(!(!(!(p2)))) {
					return (_1655_v3).Get(!(!(!(p2)))).(_dafny.Int)
				}
				return _1652_i0
			})()).Cmp(_this.F30) >= 0, 25)
			(_nw301).ArraySet1(p2, 26)
			_1661_v9 = _nw301
			var _rhs296 _dafny.Int = (_1652_i0).Minus(_dafny.IntOfInt64(376))
			_ = _rhs296
			var _rhs297 _dafny.Array = _1661_v9
			_ = _rhs297
			var _rhs298 bool = (_1654_v2).Select((Companion_Default___.SafeIndex(_1652_i0, _dafny.IntOfUint32((_1654_v2).Cardinality()))).Uint32()).(bool)
			_ = _rhs298
			var _lhs268 *GlobalState = globalState
			_ = _lhs268
			var _lhs269 *GlobalState = globalState
			_ = _lhs269
			_lhs268.F10 = _rhs296
			_1661_v9 = _rhs297
			_lhs269.F7 = _rhs298
			var _1662_v10 _dafny.CodePoint
			_ = _1662_v10
			_1662_v10 = _dafny.CodePoint('b')
			var _rhs299 _dafny.Int = _1652_i0
			_ = _rhs299
			var _rhs300 _dafny.CodePoint = _1662_v10
			_ = _rhs300
			var _rhs301 _dafny.Int = p1
			_ = _rhs301
			var _lhs270 *GlobalState = globalState
			_ = _lhs270
			var _lhs271 *GlobalState = globalState
			_ = _lhs271
			var _lhs272 *C5 = _this
			_ = _lhs272
			_lhs270.F10 = _rhs299
			_lhs271.F19 = _rhs300
			_lhs272.F30 = _rhs301
			var _1663_v11 *C3
			_ = _1663_v11
			var _nw302 *C3 = New_C3_()
			_ = _nw302
			_nw302.Ctor__()
			_1663_v11 = _nw302
			var _1664_v12 D0
			_ = _1664_v12
			_1664_v12 = Companion_D0_.Create_DC2_(p1, p2)
			var _1665_v13 D0
			_ = _1665_v13
			_1665_v13 = Companion_D0_.Create_DC3_(_1664_v12)
			var _1666_v14 D0
			_ = _1666_v14
			_1666_v14 = Companion_D0_.Create_DC3_(_1664_v12)
			var _pat_let_tv53 = _1665_v13
			_ = _pat_let_tv53
			var _pat_let_tv54 = _1664_v12
			_ = _pat_let_tv54
			var _1667_v15 _dafny.Array
			_ = _1667_v15
			var _nwElement0_70 D0 = _1666_v14
			_ = _nwElement0_70
			var _nw303 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(20))
			_ = _nw303
			(_nw303).ArraySet1(_nwElement0_70, 0)
			(_nw303).ArraySet1(_1666_v14, 1)
			(_nw303).ArraySet1(_1666_v14, 2)
			(_nw303).ArraySet1(_1666_v14, 3)
			(_nw303).ArraySet1(Companion_D0_.Create_DC3_(_1665_v13), 4)
			(_nw303).ArraySet1(func(_pat_let51_0 D0) D0 {
				return func(_1668_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let52_0 D0) D0 {
						return func(_1669_dt__update_hcf7_h0 D0) D0 {
							return Companion_D0_.Create_DC3_(_1669_dt__update_hcf7_h0)
						}(_pat_let52_0)
					}(_pat_let_tv53)
				}(_pat_let51_0)
			}(_1666_v14), 5)
			(_nw303).ArraySet1(func(_pat_let53_0 D0) D0 {
				return func(_1670_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let54_0 D0) D0 {
						return func(_1671_dt__update_hcf7_h1 D0) D0 {
							return Companion_D0_.Create_DC3_(_1671_dt__update_hcf7_h1)
						}(_pat_let54_0)
					}(_pat_let_tv54)
				}(_pat_let53_0)
			}(_1666_v14), 6)
			(_nw303).ArraySet1(_1666_v14, 7)
			(_nw303).ArraySet1(_1666_v14, 8)
			(_nw303).ArraySet1(_1666_v14, 9)
			(_nw303).ArraySet1(Companion_D0_.Create_DC3_(_1665_v13), 10)
			(_nw303).ArraySet1(_1666_v14, 11)
			(_nw303).ArraySet1(Companion_D0_.Create_DC3_(_1664_v12), 12)
			(_nw303).ArraySet1(Companion_D0_.Create_DC3_(_1664_v12), 13)
			(_nw303).ArraySet1(_1666_v14, 14)
			(_nw303).ArraySet1(_1666_v14, 15)
			(_nw303).ArraySet1(_1666_v14, 16)
			(_nw303).ArraySet1(_1666_v14, 17)
			(_nw303).ArraySet1(_1666_v14, 18)
			(_nw303).ArraySet1(Companion_D0_.Create_DC3_(_1665_v13), 19)
			_1667_v15 = _nw303
			var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1667_v15), 0))
			_ = _index305
			(_1667_v15).ArraySet1(_1666_v14, (_index305).Int())
			var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1667_v15), 0))
			_ = _index306
			(_1667_v15).ArraySet1(Companion_D0_.Create_DC3_(_1665_v13), (_index306).Int())
		}
		var _1672_v17 _dafny.MultiSet
		_ = _1672_v17
		_1672_v17 = _dafny.MultiSetOf(_this.F30)
		var _1673_v18 D2
		_ = _1673_v18
		_1673_v18 = Companion_D2_.Create_DC7_(func() _dafny.Map {
			var _coll46 = _dafny.NewMapBuilder()
			_ = _coll46
			for _iter53 := _dafny.Iterate((_1672_v17).Elements()); ; {
				_compr_46, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _1674_v16 _dafny.Int
				_1674_v16 = interface{}(_compr_46).(_dafny.Int)
				if (_1672_v17).Contains(_1674_v16) {
					_coll46.Add((_1674_v16).Plus(_dafny.IntOfInt64(645)), p1)
				}
			}
			return _coll46.ToMap()
		}())
		var _1675_v19 _dafny.Map
		_ = _1675_v19
		_1675_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
		var _1676_v20 D3
		_ = _1676_v20
		_1676_v20 = Companion_D3_.Create_DC12_((_this).Fm4(Companion_Default___.Fm23(_1673_v18, _this.F30, (_1675_v19).Cardinality(), globalState), p1, globalState))
		_1676_v20 = _1676_v20
		var _1677_v21 _dafny.Array
		_ = _1677_v21
		var _len0_42 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_42
		var _nw304 _dafny.Array
		_ = _nw304
		if _len0_42.Cmp(_dafny.Zero) == 0 {
			_nw304 = _dafny.NewArray(_len0_42)
		} else {
			var _init42 func(_dafny.Int) _dafny.Int = (func(_1678_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1679_i1 _dafny.Int) _dafny.Int {
					return (_1679_i1).Plus(_1678_p1)
				}
			})(p1)
			_ = _init42
			var _element0_42 = _init42(_dafny.Zero)
			_ = _element0_42
			_nw304 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
			(_nw304).ArraySet1(_element0_42, 0)
			var _nativeLen0_42 = (_len0_42).Int()
			_ = _nativeLen0_42
			for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
				(_nw304).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
			}
		}
		_1677_v21 = _nw304
		if (_dafny.MultiSetOf(_1677_v21, _1677_v21)).IsDisjointFrom(_dafny.MultiSetOf(_1677_v21)) {
			var _1680_v22 _dafny.MultiSet
			_ = _1680_v22
			_1680_v22 = _dafny.MultiSetOf(p2, true, p2, Companion_Default___.Fm12(p2, p2, globalState), p2)
			var _1681_v23 D1
			_ = _1681_v23
			_1681_v23 = Companion_D1_.Create_DC4_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1680_v22, p2))
			var _source22 D1 = _1681_v23
			_ = _source22
			if _source22.Is_DC5() {
				var _1682___mcc_h0 _dafny.Int = _source22.Get_().(D1_DC5).Cf9
				_ = _1682___mcc_h0
				var _1683___mcc_h1 bool = _source22.Get_().(D1_DC5).Cf10
				_ = _1683___mcc_h1
				var _1684___mcc_h2 bool = _source22.Get_().(D1_DC5).Cf11
				_ = _1684___mcc_h2
				var _1685___mcc_h3 _dafny.Map = _source22.Get_().(D1_DC5).Cf12
				_ = _1685___mcc_h3
				var _1686_cf12 _dafny.Map = _1685___mcc_h3
				_ = _1686_cf12
				var _1687_cf11 bool = _1684___mcc_h2
				_ = _1687_cf11
				var _1688_cf10 bool = _1683___mcc_h1
				_ = _1688_cf10
				var _1689_cf9 _dafny.Int = _1682___mcc_h0
				_ = _1689_cf9
				_1672_v17 = _1672_v17
				var _1690_v24 D8
				_ = _1690_v24
				_1690_v24 = Companion_D8_.Create_DC22_(_dafny.IntOfInt64(98), _1687_cf11, _dafny.CodePoint('b'))
				var _1691_v25 _dafny.CodePoint
				_ = _1691_v25
				_1691_v25 = _dafny.CodePoint('o')
				var _1692_v26 _dafny.Sequence
				_ = _1692_v26
				_1692_v26 = _dafny.SeqOf(_1690_v24, _1690_v24, _1690_v24, Companion_D8_.Create_DC22_(_1689_cf9, p2, _1691_v25))
				_1692_v26 = _1692_v26
				var _1693_v27 _dafny.Array
				_ = _1693_v27
				var _nw305 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
				_ = _nw305
				_1693_v27 = _nw305
				var _1694_v28 T1
				_ = _1694_v28
				var _nw306 *C4 = New_C4_()
				_ = _nw306
				_nw306.Ctor__(_1693_v27)
				_1694_v28 = _nw306
				var _1695_v29 _dafny.Sequence
				_ = _1695_v29
				_1695_v29 = _dafny.SeqOf(_1687_cf11)
				var _1696_v30 _dafny.Sequence
				_ = _1696_v30
				_1696_v30 = _dafny.UnicodeSeqOfUtf8Bytes("ra")
				var _1697_v32 D8
				_ = _1697_v32
				_1697_v32 = Companion_D8_.Create_DC21_(_1695_v29)
				var _1698_v33 _dafny.Sequence
				_ = _1698_v33
				_1698_v33 = _dafny.SeqOf(_1697_v32)
				var _1699_v34 _dafny.Map
				_ = _1699_v34
				_1699_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1697_v32, _dafny.IntOfUint32((_1696_v30).Cardinality()))
				var _1700_v35 _dafny.Array
				_ = _1700_v35
				var _nwElement0_71 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC21_(_1695_v29), _dafny.IntOfUint32((_1696_v30).Cardinality()))).Merge(func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter54 := _dafny.Iterate((_1698_v33).Elements()); ; {
						_compr_47, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1701_v31 D8
						_1701_v31 = interface{}(_compr_47).(D8)
						if _dafny.Companion_Sequence_.Contains(_1698_v33, _1701_v31) {
							_coll47.Add(_1701_v31, (_dafny.SetOf((_1672_v17).Cardinality())).Cardinality())
						}
					}
					return _coll47.ToMap()
				}())
				_ = _nwElement0_71
				var _nw307 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(5))
				_ = _nw307
				(_nw307).ArraySet1(_nwElement0_71, 0)
				(_nw307).ArraySet1((_1699_v34).Update(Companion_D8_.Create_DC21_(_dafny.Companion_Sequence_.Update(_1695_v29, (Companion_Default___.SafeIndex(_1689_cf9, _dafny.IntOfUint32((_1695_v29).Cardinality()))).Uint32(), p2)), _dafny.IntOfUint32((_1696_v30).Cardinality())), 1)
				(_nw307).ArraySet1(_1699_v34, 2)
				(_nw307).ArraySet1(_1699_v34, 3)
				(_nw307).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1697_v32, _this.F30)).Merge(_1699_v34), 4)
				_1700_v35 = _nw307
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1700_v35), 0))
				_ = _index307
				(_1700_v35).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1697_v32, _this.F30), (_index307).Int())
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1700_v35), 0))
				_ = _index308
				var _rhs302 bool = _1688_cf10
				_ = _rhs302
				var _rhs303 T1 = _1694_v28
				_ = _rhs303
				var _rhs304 _dafny.Map = _1699_v34
				_ = _rhs304
				var _lhs273 *GlobalState = globalState
				_ = _lhs273
				var _lhs274 _dafny.Array = _1700_v35
				_ = _lhs274
				var _lhs275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_1700_v35), 0))
				_ = _lhs275
				_lhs273.F22 = _rhs302
				_1694_v28 = _rhs303
				(_lhs274).ArraySet1(_rhs304, (_lhs275).Int())
				_1689_cf9 = (Companion_Default___.SafeDivisionInt(p1, _this.F30)).Times((func() _dafny.Int {
					if _1688_cf10 {
						return (_dafny.Zero).Minus(_this.F30)
					}
					return _dafny.IntOfUint32((_1696_v30).Cardinality())
				})())
			} else if _source22.Is_DC6() {
				var _1702___mcc_h4 _dafny.Int = _source22.Get_().(D1_DC6).Cf13
				_ = _1702___mcc_h4
				var _1703___mcc_h5 _dafny.Sequence = _source22.Get_().(D1_DC6).Cf14
				_ = _1703___mcc_h5
				var _1704___mcc_h6 _dafny.Int = _source22.Get_().(D1_DC6).Cf15
				_ = _1704___mcc_h6
				var _1705___mcc_h7 bool = _source22.Get_().(D1_DC6).Cf16
				_ = _1705___mcc_h7
				var _1706_cf16 bool = _1705___mcc_h7
				_ = _1706_cf16
				var _1707_cf15 _dafny.Int = _1704___mcc_h6
				_ = _1707_cf15
				var _1708_cf14 _dafny.Sequence = _1703___mcc_h5
				_ = _1708_cf14
				var _1709_cf13 _dafny.Int = _1702___mcc_h4
				_ = _1709_cf13
				var _1710_v36 _dafny.Sequence
				_ = _1710_v36
				_1710_v36 = _dafny.UnicodeSeqOfUtf8Bytes("rmnqe")
				var _1711_v37 _dafny.Map
				_ = _1711_v37
				_1711_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1706_cf16, _1710_v36)
				var _1712_v38 _dafny.Set
				_ = _1712_v38
				_1712_v38 = _dafny.SetOf(true)
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index309
				(_1677_v21).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if _1706_cf16 {
						return (_dafny.MultiSetOf((_1711_v37).Cardinality(), _1709_cf13, _dafny.IntOfInt64(287))).Cardinality()
					}
					return (_1712_v38).Cardinality()
				})()), (_index309).Int())
				var _1713_v39 _dafny.Array
				_ = _1713_v39
				var _nwElement0_72 _dafny.CodePoint = _dafny.CodePoint('h')
				_ = _nwElement0_72
				var _nw308 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.One)
				_ = _nw308
				(_nw308).ArraySet1CodePoint(_nwElement0_72, 0)
				_1713_v39 = _nw308
				var _1714_v40 _dafny.CodePoint
				_ = _1714_v40
				_1714_v40 = _dafny.CodePoint('t')
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1713_v39), 0))
				_ = _index310
				(_1713_v39).ArraySet1CodePoint(_1714_v40, (_index310).Int())
				var _1715_v41 D8
				_ = _1715_v41
				_1715_v41 = Companion_D8_.Create_DC23_(_1706_cf16)
				var _1716_v42 _dafny.Array
				_ = _1716_v42
				var _nwElement0_73 bool = !(p2)
				_ = _nwElement0_73
				var _nw309 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(17))
				_ = _nw309
				(_nw309).ArraySet1(_nwElement0_73, 0)
				(_nw309).ArraySet1(p2, 1)
				(_nw309).ArraySet1(p2, 2)
				(_nw309).ArraySet1(p2, 3)
				(_nw309).ArraySet1(p2, 4)
				(_nw309).ArraySet1(p2, 5)
				(_nw309).ArraySet1(p2, 6)
				(_nw309).ArraySet1(Companion_Default___.Fm12(_1706_cf16, p2, globalState), 7)
				(_nw309).ArraySet1(p2, 8)
				(_nw309).ArraySet1(_1706_cf16, 9)
				(_nw309).ArraySet1(_1706_cf16, 10)
				(_nw309).ArraySet1(p2, 11)
				(_nw309).ArraySet1(p2, 12)
				(_nw309).ArraySet1(p2, 13)
				(_nw309).ArraySet1(_1706_cf16, 14)
				(_nw309).ArraySet1(_1706_cf16, 15)
				(_nw309).ArraySet1(false, 16)
				_1716_v42 = _nw309
				var _1717_v43 D4
				_ = _1717_v43
				_1717_v43 = Companion_D4_.Create_DC13_(_1714_v40)
				var _1718_v44 D4
				_ = _1718_v44
				_1718_v44 = Companion_D4_.Create_DC15_(_1717_v43)
				var _1719_v45 _dafny.Map
				_ = _1719_v45
				_1719_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v17, _dafny.SeqOf(_1709_cf13, _this.F30))
				var _1720_v46 D1
				_ = _1720_v46
				_1720_v46 = Companion_D1_.Create_DC5_(_1709_cf13, _1706_cf16, _1706_cf16, _1719_v45)
				var _1721_v47 _dafny.Map
				_ = _1721_v47
				_1721_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, _1720_v46)
				var _1722_v49 D7
				_ = _1722_v49
				_1722_v49 = Companion_D7_.Create_DC20_(p0, _1716_v42, _1718_v44, !(_1706_cf16), (p0).Select((Companion_Default___.SafeIndex((func() _dafny.Set {
					var _coll48 = _dafny.NewBuilder()
					_ = _coll48
					for _iter55 := _dafny.Iterate((_1721_v47).Keys().Elements()); ; {
						_compr_48, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1723_v48 _dafny.Int
						_1723_v48 = interface{}(_compr_48).(_dafny.Int)
						if (_1721_v47).Contains(_1723_v48) {
							_coll48.Add(Companion_Default___.SafeModuloInt(_1723_v48, _dafny.IntOfInt64(330)))
						}
					}
					return _coll48.ToSet()
				}()).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index311
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1713_v39), 0))
				_ = _index312
				var _rhs305 _dafny.Int = (_dafny.Zero).Minus((p0).Select((Companion_Default___.SafeIndex(_1709_cf13, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs305
				var _rhs306 bool = (_1715_v41).Dtor_cf39()
				_ = _rhs306
				var _rhs307 _dafny.Int = (_1722_v49).Dtor_cf34()
				_ = _rhs307
				var _rhs308 _dafny.Int = p1
				_ = _rhs308
				var _rhs309 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (Companion_Default___.Fm0(_1709_cf13, globalState)).Cmp(_1709_cf13) <= 0 {
						return _1714_v40
					}
					return _1714_v40
				})()
				_ = _rhs309
				var _lhs276 _dafny.Array = _1677_v21
				_ = _lhs276
				var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(647), _dafny.ArrayLen((_1677_v21), 0))
				_ = _lhs277
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				var _lhs280 _dafny.Array = _1713_v39
				_ = _lhs280
				var _lhs281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((_1713_v39), 0))
				_ = _lhs281
				(_lhs276).ArraySet1(_rhs305, (_lhs277).Int())
				_lhs278.F22 = _rhs306
				_lhs279.F10 = _rhs307
				_1709_cf13 = _rhs308
				(_lhs280).ArraySet1CodePoint(_rhs309, (_lhs281).Int())
				var _1724_v50 _dafny.MultiSet
				_ = _1724_v50
				_1724_v50 = _dafny.MultiSetOf(_1710_v36, _1710_v36, _dafny.UnicodeSeqOfUtf8Bytes("ybyitirvg"))
				_1724_v50 = _1724_v50
				var _1725_v51 *C0
				_ = _1725_v51
				var _nw310 *C0 = New_C0_()
				_ = _nw310
				_nw310.Ctor__(p2)
				_1725_v51 = _nw310
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_1716_v42), 0))
				_ = _index313
				(_1716_v42).ArraySet1(p2, (_index313).Int())
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_1716_v42), 0))
				_ = _index314
				(_1716_v42).ArraySet1((_1707_cf15).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer88 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}(func(_1726_i2 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(401)
				}))).Cardinality()))) < 0, (_index314).Int())
				var _1727_v52 _dafny.Sequence
				_ = _1727_v52
				_1727_v52 = _dafny.SeqOf(_1712_v38, _1712_v38)
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_1716_v42), 0))
				_ = _index315
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_1716_v42), 0))
				_ = _index316
				var _rhs310 _dafny.Int = (_1725_v51).Fm13(_1709_cf13, globalState)
				_ = _rhs310
				var _rhs311 bool = !(_1706_cf16)
				_ = _rhs311
				var _rhs312 bool = !(!(!(p2)))
				_ = _rhs312
				var _rhs313 bool = _dafny.Companion_Sequence_.IsPrefixOf(_1727_v52, _1727_v52)
				_ = _rhs313
				var _lhs282 *GlobalState = globalState
				_ = _lhs282
				var _lhs283 _dafny.Array = _1716_v42
				_ = _lhs283
				var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(631), _dafny.ArrayLen((_1716_v42), 0))
				_ = _lhs284
				var _lhs285 _dafny.Array = _1716_v42
				_ = _lhs285
				var _lhs286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_1716_v42), 0))
				_ = _lhs286
				var _lhs287 *GlobalState = globalState
				_ = _lhs287
				_lhs282.F10 = _rhs310
				(_lhs283).ArraySet1(_rhs311, (_lhs284).Int())
				(_lhs285).ArraySet1(_rhs312, (_lhs286).Int())
				_lhs287.F13 = _rhs313
			} else {
				var _1728___mcc_h8 _dafny.Map = _source22.Get_().(D1_DC4).Cf8
				_ = _1728___mcc_h8
				var _1729_cf8 _dafny.Map = _1728___mcc_h8
				_ = _1729_cf8
				var _1730_v53 _dafny.Sequence
				_ = _1730_v53
				_1730_v53 = _dafny.SeqOf(p2, p2, p2)
				var _1731_v54 _dafny.Sequence
				_ = _1731_v54
				_1731_v54 = _dafny.SeqOf(_1730_v53)
				var _1732_v55 D1
				_ = _1732_v55
				_1732_v55 = Companion_D1_.Create_DC6_(p1, _1731_v54, p1, p2)
				var _1733_v56 _dafny.Map
				_ = _1733_v56
				_1733_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1732_v55, _1680_v22)
				_1733_v56 = (_1733_v56).Update(_1732_v55, Companion_Default___.Fm24(_dafny.IntOfInt64(199), p2, globalState))
				var _1734_v57 _dafny.Map
				_ = _1734_v57
				_1734_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(728), p1), !(true) || (p2))
				var _1735_v58 _dafny.Sequence
				_ = _1735_v58
				_1735_v58 = _dafny.UnicodeSeqOfUtf8Bytes("nvdnrsmts")
				_1734_v57 = (_1734_v57).Update(p1, !((_this).Fm4(_1735_v58, _this.F30, globalState)))
				var _1736_v59 D8
				_ = _1736_v59
				_1736_v59 = Companion_D8_.Create_DC21_(_1730_v53)
				var _1737_v60 _dafny.Map
				_ = _1737_v60
				_1737_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1736_v59, _1677_v21)
				var _1738_v61 _dafny.Map
				_ = _1738_v61
				_1738_v61 = _1737_v60
				var _1739_v62 _dafny.CodePoint
				_ = _1739_v62
				_1739_v62 = _dafny.CodePoint('u')
				var _rhs314 _dafny.Map = (_1738_v61)
				_ = _rhs314
				var _rhs315 _dafny.CodePoint = _1739_v62
				_ = _rhs315
				var _rhs316 bool = !(p2) || ((p1).Cmp(_this.F30) != 0)
				_ = _rhs316
				var _lhs288 *GlobalState = globalState
				_ = _lhs288
				var _lhs289 *GlobalState = globalState
				_ = _lhs289
				_1737_v60 = _rhs314
				_lhs288.F19 = _rhs315
				_lhs289.F7 = _rhs316
				var _1740_v63 _dafny.Map
				_ = _1740_v63
				_1740_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, _dafny.Companion_Sequence_.Concatenate(_1735_v58, _1735_v58))
				_1740_v63 = (_1740_v63).Update((_this.F30).Minus(p1), _1735_v58)
			}
			(globalState).F10 = (p1).Minus(_dafny.IntOfInt64(627))
			var _1741_v64 _dafny.Sequence
			_ = _1741_v64
			_1741_v64 = _dafny.UnicodeSeqOfUtf8Bytes("vvk")
			var _1742_v65 _dafny.Array
			_ = _1742_v65
			var _nw311 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
			_ = _nw311
			_1742_v65 = _nw311
			var _1743_v66 _dafny.Map
			_ = _1743_v66
			_1743_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v64, _1742_v65)
			_1743_v66 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v64, _1742_v65)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1741_v64, _1742_v65)).Merge(_1743_v66))
			_1742_v65 = _1742_v65
			var _1744_v67 *C0
			_ = _1744_v67
			var _nw312 *C0 = New_C0_()
			_ = _nw312
			_nw312.Ctor__(!((p1).Cmp(_this.F30) <= 0))
			_1744_v67 = _nw312
		} else {
			var _1745_v68 *C3
			_ = _1745_v68
			var _nw313 *C3 = New_C3_()
			_ = _nw313
			_nw313.Ctor__()
			_1745_v68 = _nw313
			var _1746_v69 _dafny.Map
			_ = _1746_v69
			_1746_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1745_v68)
			var _rhs317 bool = p2
			_ = _rhs317
			var _rhs318 _dafny.Map = _1746_v69
			_ = _rhs318
			var _lhs290 *GlobalState = globalState
			_ = _lhs290
			_lhs290.F13 = _rhs317
			_1746_v69 = _rhs318
			(globalState).F7 = p2
			if p2 {
				var _1747_v70 _dafny.CodePoint
				_ = _1747_v70
				_1747_v70 = _dafny.CodePoint('y')
				(globalState).F19 = _1747_v70
				var _1748_v71 _dafny.MultiSet
				_ = _1748_v71
				_1748_v71 = _dafny.MultiSetOf(p2, !(p2))
				var _1749_v72 _dafny.Map
				_ = _1749_v72
				_1749_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1677_v21)
				var _1750_v73 _dafny.Map
				_ = _1750_v73
				_1750_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1748_v71, (func() _dafny.Array {
					if (_1749_v72).Contains(!(true)) {
						return (_1749_v72).Get(!(true)).(_dafny.Array)
					}
					return _1677_v21
				})())
				_1750_v73 = (_1750_v73).Update(_1748_v71, _1677_v21)
				(globalState).F10 = ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf(p2)).Cardinality(), p1))).Plus(_this.F30)
				var _1751_v74 _dafny.Sequence
				_ = _1751_v74
				_1751_v74 = _dafny.UnicodeSeqOfUtf8Bytes("lcy")
				var _1752_v75 _dafny.Map
				_ = _1752_v75
				_1752_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_1751_v74).Cardinality()))
				(_this).F30 = ((_dafny.Zero).Minus((p1).Minus((_1752_v75).Cardinality()))).Times(_this.F30)
				var _1753_v76 _dafny.Array
				_ = _1753_v76
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(7)
				_ = _len0_43
				var _nw314 _dafny.Array
				_ = _nw314
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw314 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) bool = func(_1754_i3 _dafny.Int) bool {
						return true
					}
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw314 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw314).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw314).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1753_v76 = _nw314
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_1753_v76), 0))
				_ = _index317
				(_1753_v76).ArraySet1(p2, (_index317).Int())
				var _1755_v77 _dafny.Sequence
				_ = _1755_v77
				_1755_v77 = _dafny.SeqOf(p2, false)
				var _1756_v78 _dafny.Set
				_ = _1756_v78
				_1756_v78 = _dafny.SetOf(_dafny.SeqOf(p1), p0, p0, p0)
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_1753_v76), 0))
				_ = _index318
				var _rhs319 bool = p2
				_ = _rhs319
				var _rhs320 _dafny.Sequence = Companion_Default___.Fm22((_dafny.SetOf(p0, p0, _dafny.SeqOf(_dafny.IntOfUint32((p0).Cardinality())))).IsSubsetOf(_1756_v78), globalState)
				_ = _rhs320
				var _rhs321 bool = p2
				_ = _rhs321
				var _lhs291 _dafny.Array = _1753_v76
				_ = _lhs291
				var _lhs292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(79), _dafny.ArrayLen((_1753_v76), 0))
				_ = _lhs292
				var _lhs293 *GlobalState = globalState
				_ = _lhs293
				(_lhs291).ArraySet1(_rhs319, (_lhs292).Int())
				_1755_v77 = _rhs320
				_lhs293.F13 = _rhs321
			} else {
				var _1757_v79 _dafny.Sequence
				_ = _1757_v79
				_1757_v79 = _dafny.SeqOf(Companion_Default___.Fm12(true, p2, globalState))
				var _1758_v80 _dafny.Sequence
				_ = _1758_v80
				_1758_v80 = _dafny.UnicodeSeqOfUtf8Bytes("shrji")
				var _1759_v81 *C0
				_ = _1759_v81
				var _nw315 *C0 = New_C0_()
				_ = _nw315
				_nw315.Ctor__((_1757_v79).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1758_v80).Cardinality()), _dafny.IntOfUint32((_1757_v79).Cardinality()))).Uint32()).(bool))
				_1759_v81 = _nw315
				var _1760_v82 _dafny.Sequence
				_ = _1760_v82
				_1760_v82 = _dafny.SeqOf(_1759_v81)
				var _1761_v83 _dafny.Array
				_ = _1761_v83
				var _nwElement0_74 *C0 = _1759_v81
				_ = _nwElement0_74
				var _nw316 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(23))
				_ = _nw316
				(_nw316).ArraySet1(_nwElement0_74, 0)
				(_nw316).ArraySet1(_1759_v81, 1)
				(_nw316).ArraySet1(_1759_v81, 2)
				(_nw316).ArraySet1(_1759_v81, 3)
				(_nw316).ArraySet1(_1759_v81, 4)
				(_nw316).ArraySet1(_1759_v81, 5)
				(_nw316).ArraySet1((_1760_v82).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1760_v82).Cardinality()))).Uint32()).(*C0), 6)
				(_nw316).ArraySet1(_1759_v81, 7)
				(_nw316).ArraySet1(_1759_v81, 8)
				(_nw316).ArraySet1(_1759_v81, 9)
				(_nw316).ArraySet1(_1759_v81, 10)
				(_nw316).ArraySet1(_1759_v81, 11)
				(_nw316).ArraySet1((func() *C0 {
					if p2 {
						return _1759_v81
					}
					return _1759_v81
				})(), 12)
				(_nw316).ArraySet1(_1759_v81, 13)
				(_nw316).ArraySet1(_1759_v81, 14)
				(_nw316).ArraySet1(_1759_v81, 15)
				(_nw316).ArraySet1((func() *C0 {
					if (_1759_v81).F32() {
						return _1759_v81
					}
					return _1759_v81
				})(), 16)
				(_nw316).ArraySet1(_1759_v81, 17)
				(_nw316).ArraySet1(_1759_v81, 18)
				(_nw316).ArraySet1(_1759_v81, 19)
				(_nw316).ArraySet1(_1759_v81, 20)
				(_nw316).ArraySet1(_1759_v81, 21)
				(_nw316).ArraySet1(_1759_v81, 22)
				_1761_v83 = _nw316
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1761_v83), 0))
				_ = _index319
				(_1761_v83).ArraySet1(_1759_v81, (_index319).Int())
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(455), _dafny.ArrayLen((_1761_v83), 0))
				_ = _index320
				(_1761_v83).ArraySet1(_1759_v81, (_index320).Int())
				(globalState).F10 = _dafny.IntOfUint32((_1758_v80).Cardinality())
				var _1762_v84 _dafny.Map
				_ = _1762_v84
				_1762_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p0).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(345))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg89 _dafny.Int) interface{} {
						return coer89(arg89)
					}
				}(func(_1763_i4 _dafny.Int) _dafny.Int {
					return _this.F30
				})))
				var _1764_v85 _dafny.Array
				_ = _1764_v85
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_44
				var _nw317 _dafny.Array
				_ = _nw317
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw317 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) bool = (func(_1765_p2 bool) func(_dafny.Int) bool {
						return func(_1766_i5 _dafny.Int) bool {
							return !(_1765_p2)
						}
					})(p2)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw317 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw317).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw317).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1764_v85 = _nw317
				var _1767_v86 _dafny.CodePoint
				_ = _1767_v86
				_1767_v86 = _dafny.CodePoint('x')
				var _1768_v87 D4
				_ = _1768_v87
				_1768_v87 = Companion_D4_.Create_DC13_(_1767_v86)
				var _1769_v88 D4
				_ = _1769_v88
				_1769_v88 = Companion_D4_.Create_DC15_(_1768_v87)
				var _1770_v89 D7
				_ = _1770_v89
				_1770_v89 = Companion_D7_.Create_DC20_((func() _dafny.Sequence {
					if (_1762_v84).Contains(_dafny.IntOfInt64(897)) {
						return (_1762_v84).Get(_dafny.IntOfInt64(897)).(_dafny.Sequence)
					}
					return p0
				})(), _1764_v85, _1769_v88, p2, (_this).Fm2(globalState))
				var _rhs322 _dafny.Int = _this.F30
				_ = _rhs322
				var _rhs323 D7 = (func() D7 {
					if (p1).Cmp(_this.F30) <= 0 {
						return _1770_v89
					}
					return _1770_v89
				})()
				_ = _rhs323
				var _lhs294 *C5 = _this
				_ = _lhs294
				_lhs294.F30 = _rhs322
				_1770_v89 = _rhs323
				(globalState).F10 = (_dafny.Zero).Minus(_this.F30)
				var _1771_v90 _dafny.Map
				_ = _1771_v90
				_1771_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _1772_v91 _dafny.Map
				_ = _1772_v91
				_1772_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), (_1759_v81).F32())
				var _1773_v92 _dafny.Array
				_ = _1773_v92
				var _nw318 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(11))
				_ = _nw318
				_1773_v92 = _nw318
				var _1774_v93 _dafny.Sequence
				_ = _1774_v93
				_1774_v93 = _dafny.SeqOf(_1758_v80, _dafny.UnicodeSeqOfUtf8Bytes("e"))
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1773_v92), 0))
				_ = _index321
				(_1773_v92).ArraySet1(Companion_Default___.Fm41(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1774_v93, (Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1774_v93).Cardinality()))).Uint32(), _1758_v80)).Cardinality()), p2, Companion_Default___.Fm24(_dafny.IntOfInt64(93), p2, globalState), _1672_v17, globalState), (_index321).Int())
				var _1775_v95 D4
				_ = _1775_v95
				_1775_v95 = Companion_D4_.Create_DC13_(Companion_Default___.Fm11(p1, globalState))
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1773_v92), 0))
				_ = _index322
				var _rhs324 _dafny.Map = (func() _dafny.Map {
					var _coll49 = _dafny.NewMapBuilder()
					_ = _coll49
					for _iter56 := _dafny.Iterate((_1651_v0).Keys().Elements()); ; {
						_compr_49, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1776_v94 _dafny.Int
						_1776_v94 = interface{}(_compr_49).(_dafny.Int)
						if (_1651_v0).Contains(_1776_v94) {
							_coll49.Add(Companion_Default___.SafeDivisionInt(_1776_v94, (_1672_v17).Cardinality()), (_1759_v81).F32())
						}
					}
					return _coll49.ToMap()
				}()).Merge(_1771_v90)
				_ = _rhs324
				var _rhs325 _dafny.Map = _1772_v91
				_ = _rhs325
				var _rhs326 D4 = _1775_v95
				_ = _rhs326
				var _lhs295 _dafny.Array = _1773_v92
				_ = _lhs295
				var _lhs296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(114), _dafny.ArrayLen((_1773_v92), 0))
				_ = _lhs296
				_1771_v90 = _rhs324
				_1772_v91 = _rhs325
				(_lhs295).ArraySet1(_rhs326, (_lhs296).Int())
			}
			if !(p2) || (false) {
				var _1777_v96 _dafny.CodePoint
				_ = _1777_v96
				_1777_v96 = _dafny.CodePoint('i')
				var _1778_v97 _dafny.MultiSet
				_ = _1778_v97
				_1778_v97 = _dafny.MultiSetOf(_1777_v96)
				var _rhs327 _dafny.CodePoint = (func() _dafny.CodePoint {
					if p2 {
						return _dafny.CodePoint('q')
					}
					return _1777_v96
				})()
				_ = _rhs327
				var _rhs328 _dafny.Int = p1
				_ = _rhs328
				var _rhs329 _dafny.Int = ((p0).Select((Companion_Default___.SafeIndex(((_1778_v97).Update(_1777_v96, Companion_Default___.Abs(p1))).Cardinality(), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_this.F30)
				_ = _rhs329
				var _rhs330 _dafny.Int = _this.F30
				_ = _rhs330
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				var _lhs299 *C5 = _this
				_ = _lhs299
				var _lhs300 *C5 = _this
				_ = _lhs300
				_lhs297.F19 = _rhs327
				_lhs298.F10 = _rhs328
				_lhs299.F30 = _rhs329
				_lhs300.F30 = _rhs330
				var _1779_v98 _dafny.Set
				_ = _1779_v98
				_1779_v98 = _dafny.SetOf(p2, p2, p2, p2)
				(globalState).F13 = !((_1779_v98).Equals(_1779_v98))
				(_this).F30 = (Companion_Default___.SafeDivisionInt(_this.F30, _this.F30)).Minus(_this.F30)
				(globalState).F13 = !(!(true)) || (true)
				(globalState).F10 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p0, p0)).Cardinality())
			} else {
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index323
				(_1677_v21).ArraySet1((_dafny.Zero).Minus(_this.F30), (_index323).Int())
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index324
				(_1677_v21).ArraySet1((p1).Times(((_1745_v68).Fm2(globalState)).Times(p1)), (_index324).Int())
				(globalState).F10 = (func() _dafny.Int {
					if !(_1675_v19).Equals(_1675_v19) {
						return p1
					}
					return _this.F30
				})()
				var _1780_v99 _dafny.Sequence
				_ = _1780_v99
				_1780_v99 = _dafny.SeqOf(p2)
				var _1781_v100 _dafny.MultiSet
				_ = _1781_v100
				_1781_v100 = _dafny.MultiSetOf(_1780_v99, _1780_v99)
				var _1782_v101 _dafny.MultiSet
				_ = _1782_v101
				_1782_v101 = _dafny.MultiSetOf(p2)
				(globalState).F10 = (func() _dafny.Int {
					if !(p2) {
						return (_1781_v100).Cardinality()
					}
					return (_dafny.Zero).Minus(((_1677_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1677_v21), 0))).Int()).(_dafny.Int)).Plus((_dafny.MultiSetOf((_1677_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1677_v21), 0))).Int()).(_dafny.Int), (_1782_v101).Cardinality(), _this.F30)).Cardinality()))
				})()
				var _1783_v102 _dafny.Map
				_ = _1783_v102
				_1783_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1677_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(567), _dafny.ArrayLen((_1677_v21), 0))).Int()).(_dafny.Int))
				var _1784_v103 _dafny.Sequence
				_ = _1784_v103
				_1784_v103 = _dafny.UnicodeSeqOfUtf8Bytes("rbconal")
				_1783_v102 = (_1783_v102).Update(p2, _dafny.IntOfUint32((_1784_v103).Cardinality()))
				(globalState).F13 = p2
			}
			var _1785_v104 _dafny.Sequence
			_ = _1785_v104
			_1785_v104 = _dafny.UnicodeSeqOfUtf8Bytes("llhcw")
			var _1786_v105 D0
			_ = _1786_v105
			_1786_v105 = Companion_D0_.Create_DC2_((func() _dafny.Int {
				if !(p2) {
					return (func() _dafny.Int {
						if (_1672_v17).Contains(_dafny.IntOfInt64(714)) {
							return (_1672_v17).Multiplicity(_dafny.IntOfInt64(714))
						}
						return _this.F30
					})()
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_1785_v104).Cardinality()))
			})(), !(p2))
			_1786_v105 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F30, (_dafny.Zero).Minus(p1))), p2)
		}
		var _rhs331 _dafny.Int = (func() _dafny.Int {
			if (_this.F30).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(p2)).Cardinality())) < 0 {
				return _this.F30
			}
			return _this.F30
		})()
		_ = _rhs331
		var _rhs332 bool = p2
		_ = _rhs332
		var _rhs333 _dafny.Int = (func() _dafny.Int {
			if p2 {
				return (p0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p0).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)
			}
			return (_dafny.SetOf(p2)).Cardinality()
		})()
		_ = _rhs333
		var _lhs301 *GlobalState = globalState
		_ = _lhs301
		var _lhs302 *C5 = _this
		_ = _lhs302
		_lhs301.F10 = _rhs331
		r1 = _rhs332
		_lhs302.F30 = _rhs333
		if p2 {
			var _1787_v106 _dafny.Array
			_ = _1787_v106
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_45
			var _nw319 _dafny.Array
			_ = _nw319
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw319 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) bool = (func(_1788_p2 bool) func(_dafny.Int) bool {
					return func(_1789_i6 _dafny.Int) bool {
						return _1788_p2
					}
				})(p2)
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw319 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw319).ArraySet1(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw319).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1787_v106 = _nw319
			var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))
			_ = _index325
			(_1787_v106).ArraySet1(p2, (_index325).Int())
			var _1790_v107 _dafny.Array
			_ = _1790_v107
			var _len0_46 _dafny.Int = _dafny.One
			_ = _len0_46
			var _nw320 _dafny.Array
			_ = _nw320
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw320 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) _dafny.CodePoint = func(_1791_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw320 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw320).ArraySet1CodePoint(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw320).ArraySet1CodePoint(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_1790_v107 = _nw320
			var _1792_v108 _dafny.Array
			_ = _1792_v108
			_1792_v108 = _1790_v107
			var _1793_v109 _dafny.Sequence
			_ = _1793_v109
			_1793_v109 = _dafny.SeqOf(_1790_v107)
			var _1794_v110 _dafny.Array
			_ = _1794_v110
			var _nwElement0_75 _dafny.Array = _1790_v107
			_ = _nwElement0_75
			var _nw321 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(17))
			_ = _nw321
			(_nw321).ArraySet1(_nwElement0_75, 0)
			(_nw321).ArraySet1(_1790_v107, 1)
			(_nw321).ArraySet1(_1790_v107, 2)
			(_nw321).ArraySet1(_1790_v107, 3)
			(_nw321).ArraySet1(_1790_v107, 4)
			(_nw321).ArraySet1(_1790_v107, 5)
			(_nw321).ArraySet1(_1790_v107, 6)
			(_nw321).ArraySet1(_1790_v107, 7)
			(_nw321).ArraySet1(_1790_v107, 8)
			(_nw321).ArraySet1(_1790_v107, 9)
			(_nw321).ArraySet1(_1790_v107, 10)
			(_nw321).ArraySet1(_1790_v107, 11)
			(_nw321).ArraySet1(_1790_v107, 12)
			(_nw321).ArraySet1((_1792_v108), 13)
			(_nw321).ArraySet1(_1790_v107, 14)
			(_nw321).ArraySet1(_1790_v107, 15)
			(_nw321).ArraySet1((_1793_v109).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1793_v109).Cardinality()))).Uint32()).(_dafny.Array), 16)
			_1794_v110 = _nw321
			var _1795_v111 _dafny.Sequence
			_ = _1795_v111
			_1795_v111 = _dafny.SeqOf(_1794_v110)
			var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))
			_ = _index326
			var _rhs334 bool = p2
			_ = _rhs334
			var _rhs335 bool = p2
			_ = _rhs335
			var _rhs336 _dafny.Array = (_1795_v111).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1795_v111).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs336
			var _lhs303 _dafny.Array = _1787_v106
			_ = _lhs303
			var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))
			_ = _lhs304
			var _lhs305 *GlobalState = globalState
			_ = _lhs305
			var _lhs306 *GlobalState = globalState
			_ = _lhs306
			(_lhs303).ArraySet1(_rhs334, (_lhs304).Int())
			_lhs305.F22 = _rhs335
			_lhs306.F6 = _rhs336
			var _1796_v112 _dafny.MultiSet
			_ = _1796_v112
			_1796_v112 = _dafny.MultiSetOf(_dafny.SetOf(_this.F30))
			if (_1796_v112).IsProperSubsetOf(_1796_v112) {
				var _1797_v113 _dafny.Sequence
				_ = _1797_v113
				_1797_v113 = _dafny.SeqOf((_1672_v17).IsSubsetOf(_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)).Cardinality(), p1, p1, _this.F30)), p2)
				_1797_v113 = Companion_Default___.Fm22(!((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool)), globalState)
				var _1798_v114 _dafny.Sequence
				_ = _1798_v114
				_1798_v114 = _dafny.UnicodeSeqOfUtf8Bytes("bvnrgdjr")
				var _1799_v115 _dafny.Sequence
				_ = _1799_v115
				_1799_v115 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg90 _dafny.Int) interface{} {
						return coer90(arg90)
					}
				}(func(_1800_i9 _dafny.Int) _dafny.Int {
					return _this.F30
				})))
				var _1801_v116 _dafny.CodePoint
				_ = _1801_v116
				_1801_v116 = _dafny.CodePoint('m')
				var _1802_v117 D0
				_ = _1802_v117
				_1802_v117 = Companion_D0_.Create_DC0_(_1798_v114)
				var _1803_v118 _dafny.Array
				_ = _1803_v118
				var _nwElement0_76 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1798_v114, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg91 _dafny.Int) interface{} {
						return coer91(arg91)
					}
				}(func(_1804_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1799_v115).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1798_v114, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(177))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}(func(_1805_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})))).Cardinality()))).Uint32(), _1801_v116)
				_ = _nwElement0_76
				var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(11))
				_ = _nw322
				(_nw322).ArraySet1(_nwElement0_76, 0)
				(_nw322).ArraySet1(_1798_v114, 1)
				(_nw322).ArraySet1(_1798_v114, 2)
				(_nw322).ArraySet1(_1798_v114, 3)
				(_nw322).ArraySet1(_1798_v114, 4)
				(_nw322).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-829))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg93 _dafny.Int) interface{} {
						return coer93(arg93)
					}
				}((func(_1806_v116 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1807_i10 _dafny.Int) _dafny.CodePoint {
						return _1806_v116
					}
				})(_1801_v116))), 5)
				(_nw322).ArraySet1(_1798_v114, 6)
				(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg94 _dafny.Int) interface{} {
						return coer94(arg94)
					}
				}((func(_1808_v116 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1809_i11 _dafny.Int) _dafny.CodePoint {
						return _1808_v116
					}
				})(_1801_v116))), _dafny.UnicodeSeqOfUtf8Bytes("aiqhwkuhp")), 7)
				(_nw322).ArraySet1((_1802_v117).Dtor_cf0(), 8)
				(_nw322).ArraySet1(_1798_v114, 9)
				(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1798_v114, _1798_v114), 10)
				_1803_v118 = _nw322
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))
				_ = _index327
				(_1803_v118).ArraySet1(_1798_v114, (_index327).Int())
				var _1810_v119 _dafny.Map
				_ = _1810_v119
				_1810_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-220))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg95 _dafny.Int) interface{} {
						return coer95(arg95)
					}
				}((func(_1811_v116 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1812_i12 _dafny.Int) _dafny.CodePoint {
						return _1811_v116
					}
				})(_1801_v116))))
				var _1813_v120 _dafny.Map
				_ = _1813_v120
				_1813_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1787_v106, p1)
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))
				_ = _index328
				(_1803_v118).ArraySet1((func() _dafny.Sequence {
					if (_1810_v119).Contains((p2) && ((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool))) {
						return (_1810_v119).Get((p2) && ((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool))).(_dafny.Sequence)
					}
					return Companion_Default___.Fm23(_1673_v18, p1, ((_1813_v120).Update(_1787_v106, p1)).Cardinality(), globalState)
				})(), (_index328).Int())
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))
				_ = _index329
				(_1803_v118).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_1803_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))).Int()).(_dafny.Sequence), _1798_v114), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("d")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_1803_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))).Int()).(_dafny.Sequence), _1798_v114)).Cardinality()))).Uint32(), _1801_v116), (_index329).Int())
				(globalState).F22 = (_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool)
				var _1814_v121 _dafny.Sequence
				_ = _1814_v121
				_1814_v121 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1801_v116, _dafny.IntOfInt64(467)))
				var _1815_v122 D7
				_ = _1815_v122
				_1815_v122 = Companion_D7_.Create_DC18_(_1814_v121)
				var _pat_let_tv55 = _1814_v121
				_ = _pat_let_tv55
				var _1816_v123 _dafny.Sequence
				_ = _1816_v123
				_1816_v123 = _dafny.SeqOf(_1815_v122, func(_pat_let55_0 D7) D7 {
					return func(_1817_dt__update__tmp_h2 D7) D7 {
						return func(_pat_let56_0 _dafny.Sequence) D7 {
							return func(_1818_dt__update_hcf29_h0 _dafny.Sequence) D7 {
								return Companion_D7_.Create_DC18_(_1818_dt__update_hcf29_h0)
							}(_pat_let56_0)
						}(_pat_let_tv55)
					}(_pat_let55_0)
				}(_1815_v122))
				(globalState).F13 = (_this).Fm4((_1803_v118).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_1803_v118), 0))).Int()).(_dafny.Sequence), _dafny.IntOfUint32((_1816_v123).Cardinality()), globalState)
			} else {
				(globalState).F10 = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_dafny.IntOfInt64(-529), globalState), _this.F30)
				var _1819_v124 _dafny.MultiSet
				_ = _1819_v124
				_1819_v124 = _dafny.MultiSetOf(p2)
				var _1820_v125 _dafny.Sequence
				_ = _1820_v125
				_1820_v125 = _dafny.SeqOf((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool))
				var _1821_v126 _dafny.Set
				_ = _1821_v126
				_1821_v126 = _dafny.SetOf(_1787_v106, _1787_v106)
				var _rhs337 _dafny.MultiSet = _dafny.MultiSetOf(p2, !(_dafny.SetOf(_dafny.IntOfUint32((_1820_v125).Cardinality()))).Equals(_dafny.SetOf(p1)), !((_dafny.SetOf(_1787_v106)).Equals(_1821_v126)), (_dafny.MultiSetOf(_this.F30)).IsProperSubsetOf(_dafny.MultiSetOf(p1)))
				_ = _rhs337
				var _rhs338 bool = p2
				_ = _rhs338
				var _rhs339 _dafny.Int = p1
				_ = _rhs339
				var _lhs307 *GlobalState = globalState
				_ = _lhs307
				var _lhs308 *C5 = _this
				_ = _lhs308
				_1819_v124 = _rhs337
				_lhs307.F7 = _rhs338
				_lhs308.F30 = _rhs339
				var _1822_v127 _dafny.Array
				_ = _1822_v127
				var _nw323 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(19))
				_ = _nw323
				_1822_v127 = _nw323
				var _1823_v128 _dafny.Map
				_ = _1823_v128
				_1823_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetFromSeq(_1820_v125), p2)
				var _1824_v129 D1
				_ = _1824_v129
				_1824_v129 = Companion_D1_.Create_DC4_(_1823_v128)
				var _1825_v130 _dafny.Map
				_ = _1825_v130
				_1825_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1824_v129, _1822_v127)
				var _1826_v131 _dafny.Map
				_ = _1826_v131
				_1826_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1822_v127)
				var _1827_v132 _dafny.Array
				_ = _1827_v132
				var _nwElement0_77 _dafny.Array = _1822_v127
				_ = _nwElement0_77
				var _nw324 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(27))
				_ = _nw324
				(_nw324).ArraySet1(_nwElement0_77, 0)
				(_nw324).ArraySet1(_1822_v127, 1)
				(_nw324).ArraySet1(_1822_v127, 2)
				(_nw324).ArraySet1(_1822_v127, 3)
				(_nw324).ArraySet1(_1822_v127, 4)
				(_nw324).ArraySet1(_1822_v127, 5)
				(_nw324).ArraySet1(_1822_v127, 6)
				(_nw324).ArraySet1(_1822_v127, 7)
				(_nw324).ArraySet1(_1822_v127, 8)
				(_nw324).ArraySet1(_1822_v127, 9)
				(_nw324).ArraySet1(_1822_v127, 10)
				(_nw324).ArraySet1(_1822_v127, 11)
				(_nw324).ArraySet1(_1822_v127, 12)
				(_nw324).ArraySet1(_1822_v127, 13)
				(_nw324).ArraySet1(_1822_v127, 14)
				(_nw324).ArraySet1(_1822_v127, 15)
				(_nw324).ArraySet1(_1822_v127, 16)
				(_nw324).ArraySet1(_1822_v127, 17)
				(_nw324).ArraySet1(_1822_v127, 18)
				(_nw324).ArraySet1(_1822_v127, 19)
				(_nw324).ArraySet1(_1822_v127, 20)
				(_nw324).ArraySet1((func() _dafny.Array {
					if (_1825_v130).Contains(_1824_v129) {
						return (_1825_v130).Get(_1824_v129).(_dafny.Array)
					}
					return _1822_v127
				})(), 21)
				(_nw324).ArraySet1(_1822_v127, 22)
				(_nw324).ArraySet1(_1822_v127, 23)
				(_nw324).ArraySet1(_1822_v127, 24)
				(_nw324).ArraySet1(_1822_v127, 25)
				(_nw324).ArraySet1((func() _dafny.Array {
					if (_1826_v131).Contains(p2) {
						return (_1826_v131).Get(p2).(_dafny.Array)
					}
					return _1822_v127
				})(), 26)
				_1827_v132 = _nw324
				var _1828_v133 *C4
				_ = _1828_v133
				var _nw325 *C4 = New_C4_()
				_ = _nw325
				_nw325.Ctor__(_1827_v132)
				_1828_v133 = _nw325
				var _1829_v134 _dafny.Map
				_ = _1829_v134
				_1829_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() *C4 {
					if p2 {
						return _1828_v133
					}
					return _1828_v133
				})())
				_1829_v134 = (_1829_v134).Update((_1787_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))).Int()).(bool), _1828_v133)
				var _1830_v135 _dafny.Set
				_ = _1830_v135
				_1830_v135 = _dafny.SetOf(((_dafny.Zero).Minus(_this.F30)).Plus(_this.F30), _this.F30)
				(globalState).F24 = _1830_v135
				var _1831_v136 *C4
				_ = _1831_v136
				var _nw326 *C4 = New_C4_()
				_ = _nw326
				_nw326.Ctor__((_1828_v133).F31())
				_1831_v136 = _nw326
			}
			(globalState).F13 = ((_dafny.Zero).Minus(_this.F30)).Cmp(Companion_Default___.SafeDivisionInt(_this.F30, p1)) >= 0
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(119), _dafny.ArrayLen((_1787_v106), 0))
			_ = _index330
			(_1787_v106).ArraySet1(p2, (_index330).Int())
			(globalState).F13 = p2
		} else {
			var _1832_v137 _dafny.Array
			_ = _1832_v137
			var _nw327 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw327
			_1832_v137 = _nw327
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen(((_this).F29()), 0))
			_ = _index331
			((_this).F29()).ArraySet1(_1677_v21, (_index331).Int())
			var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen(((_this).F29()), 0))
			_ = _index332
			var _rhs340 _dafny.Array = _1832_v137
			_ = _rhs340
			var _rhs341 _dafny.Int = _this.F30
			_ = _rhs341
			var _rhs342 _dafny.Array = _1677_v21
			_ = _rhs342
			var _rhs343 _dafny.Int = _this.F30
			_ = _rhs343
			var _lhs309 *GlobalState = globalState
			_ = _lhs309
			var _lhs310 _dafny.Array = (_this).F29()
			_ = _lhs310
			var _lhs311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen(((_this).F29()), 0))
			_ = _lhs311
			var _lhs312 *GlobalState = globalState
			_ = _lhs312
			_1832_v137 = _rhs340
			_lhs309.F10 = _rhs341
			(_lhs310).ArraySet1(_rhs342, (_lhs311).Int())
			_lhs312.F10 = _rhs343
			if p2 {
				var _1833_v138 _dafny.MultiSet
				_ = _1833_v138
				_1833_v138 = _dafny.MultiSetOf(p2)
				(globalState).F10 = (Companion_Default___.SafeDivisionInt(p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ekgpmyva")).Cardinality())))).Plus(((_1833_v138).Union(_dafny.MultiSetOf(p2))).Cardinality())
				(globalState).F10 = Companion_Default___.SafeModuloInt(p1, Companion_Default___.SafeDivisionInt(_this.F30, _this.F30))
				(globalState).F13 = p2
				(globalState).F7 = p2
				(globalState).F22 = false
			} else {
				(globalState).F13 = p2
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index333
				(_1677_v21).ArraySet1(p1, (_index333).Int())
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(999), _dafny.ArrayLen((_1677_v21), 0))
				_ = _index334
				(_1677_v21).ArraySet1(_this.F30, (_index334).Int())
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen(((_this).F29()), 0))
				_ = _index335
				((_this).F29()).ArraySet1(_dafny.ArrayCastTo(((_this).F29()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen(((_this).F29()), 0))).Int())), (_index335).Int())
				var _1834_v139 _dafny.Sequence
				_ = _1834_v139
				_1834_v139 = _dafny.UnicodeSeqOfUtf8Bytes("lhnsen")
				_1834_v139 = _1834_v139
				var _1835_v140 _dafny.Array
				_ = _1835_v140
				var _nwElement0_78 bool = (p2) == (p2)
				_ = _nwElement0_78
				var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(10))
				_ = _nw328
				(_nw328).ArraySet1(_nwElement0_78, 0)
				(_nw328).ArraySet1(p2, 1)
				(_nw328).ArraySet1(p2, 2)
				(_nw328).ArraySet1(p2, 3)
				(_nw328).ArraySet1(p2, 4)
				(_nw328).ArraySet1((func() bool {
					if p2 {
						return p2
					}
					return p2
				})(), 5)
				(_nw328).ArraySet1(p2, 6)
				(_nw328).ArraySet1(p2, 7)
				(_nw328).ArraySet1(true, 8)
				(_nw328).ArraySet1(p2, 9)
				_1835_v140 = _nw328
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1835_v140), 0))
				_ = _index336
				(_1835_v140).ArraySet1(!(p2), (_index336).Int())
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1835_v140), 0))
				_ = _index337
				(_1835_v140).ArraySet1(p2, (_index337).Int())
			}
			var _1836_v141 _dafny.Map
			_ = _1836_v141
			_1836_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F30), p2)
			var _1837_v142 _dafny.Map
			_ = _1837_v142
			_1837_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1672_v17, p0)
			var _1838_v143 D1
			_ = _1838_v143
			_1838_v143 = Companion_D1_.Create_DC5_((_1836_v141).Cardinality(), true, p2, _1837_v142)
			var _1839_v144 _dafny.Sequence
			_ = _1839_v144
			_1839_v144 = _dafny.UnicodeSeqOfUtf8Bytes("ipg")
			var _1840_v145 _dafny.Array
			_ = _1840_v145
			var _nwElement0_79 bool = p2
			_ = _nwElement0_79
			var _nw329 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(27))
			_ = _nw329
			(_nw329).ArraySet1(_nwElement0_79, 0)
			(_nw329).ArraySet1(p2, 1)
			(_nw329).ArraySet1(p2, 2)
			(_nw329).ArraySet1(p2, 3)
			(_nw329).ArraySet1(p2, 4)
			(_nw329).ArraySet1(p2, 5)
			(_nw329).ArraySet1(!(p2), 6)
			(_nw329).ArraySet1(p2, 7)
			(_nw329).ArraySet1((_1838_v143).Dtor_cf11(), 8)
			(_nw329).ArraySet1(p2, 9)
			(_nw329).ArraySet1(p2, 10)
			(_nw329).ArraySet1(p2, 11)
			(_nw329).ArraySet1(p2, 12)
			(_nw329).ArraySet1(p2, 13)
			(_nw329).ArraySet1(p2, 14)
			(_nw329).ArraySet1(!(true), 15)
			(_nw329).ArraySet1(p2, 16)
			(_nw329).ArraySet1(true, 17)
			(_nw329).ArraySet1(p2, 18)
			(_nw329).ArraySet1(p2, 19)
			(_nw329).ArraySet1(p2, 20)
			(_nw329).ArraySet1((_this).Fm4(_1839_v144, p1, globalState), 21)
			(_nw329).ArraySet1(p2, 22)
			(_nw329).ArraySet1(true, 23)
			(_nw329).ArraySet1(p2, 24)
			(_nw329).ArraySet1(p2, 25)
			(_nw329).ArraySet1(p2, 26)
			_1840_v145 = _nw329
			var _1841_v146 D4
			_ = _1841_v146
			_1841_v146 = Companion_D4_.Create_DC14_(p2)
			var _1842_v147 D4
			_ = _1842_v147
			_1842_v147 = Companion_D4_.Create_DC15_(_1841_v146)
			var _1843_v148 D7
			_ = _1843_v148
			_1843_v148 = Companion_D7_.Create_DC20_(p0, _1840_v145, Companion_D4_.Create_DC15_(_1841_v146), !_dafny.Companion_Sequence_.Contains(p0, _this.F30), p1)
			var _1844_v149 _dafny.Sequence
			_ = _1844_v149
			_1844_v149 = _dafny.SeqOf(p0, _dafny.SeqOf(_this.F30, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(937))).Uint32(), func(coer96 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg96 _dafny.Int) interface{} {
					return coer96(arg96)
				}
			}(func(_1845_i13 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_this.F30)
			}))).Cardinality())), _this.F30), p0, p0, p0)
			var _1846_v150 D4
			_ = _1846_v150
			_1846_v150 = Companion_D4_.Create_DC15_(_1841_v146)
			_1843_v148 = (func() D7 {
				if p2 {
					return Companion_D7_.Create_DC20_(_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1844_v149).Select((Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1844_v149).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), p1), _1840_v145, _1846_v150, p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, (_dafny.Zero).Minus(p1))).Cardinality())
				}
				return _1843_v148
			})()
			var _1847_v151 *C1
			_ = _1847_v151
			var _nw330 *C1 = New_C1_()
			_ = _nw330
			_nw330.Ctor__()
			_1847_v151 = _nw330
			if ((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, (_this).Fm4(_1839_v144, _this.F30, globalState))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p0, (Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((p0).Cardinality()))).Uint32(), _dafny.IntOfUint32((p0).Cardinality()))).Cardinality())) >= 0 {
				var _1848_v152 D0
				_ = _1848_v152
				_1848_v152 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus(p1), p2)
				var _1849_v153 _dafny.Sequence
				_ = _1849_v153
				_1849_v153 = _dafny.SeqOf(_1848_v152)
				var _1850_v154 _dafny.Sequence
				_ = _1850_v154
				_1850_v154 = _dafny.SeqOf(_1849_v153, _1849_v153, _dafny.Companion_Sequence_.Update(_1849_v153, (Companion_Default___.SafeIndex(_this.F30, _dafny.IntOfUint32((_1849_v153).Cardinality()))).Uint32(), _1848_v152), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer97 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg97 _dafny.Int) interface{} {
						return coer97(arg97)
					}
				}((func(_1851_v152 D0) func(_dafny.Int) D0 {
					return func(_1852_i14 _dafny.Int) D0 {
						return _1851_v152
					}
				})(_1848_v152))), _1849_v153)
				var _1853_v155 _dafny.Map
				_ = _1853_v155
				_1853_v155 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.Companion_Sequence_.Update(_1850_v154, (Companion_Default___.SafeIndex((_1651_v0).Cardinality(), _dafny.IntOfUint32((_1850_v154).Cardinality()))).Uint32(), _1849_v153))
				_1850_v154 = (func() _dafny.Sequence {
					if p2 {
						return _1850_v154
					}
					return (func() _dafny.Sequence {
						if (_1853_v155).Contains(false) {
							return (_1853_v155).Get(false).(_dafny.Sequence)
						}
						return _1850_v154
					})()
				})()
				var _1854_v156 *C3
				_ = _1854_v156
				var _nw331 *C3 = New_C3_()
				_ = _nw331
				_nw331.Ctor__()
				_1854_v156 = _nw331
				(globalState).F10 = p1
				var _1855_v157 _dafny.Map
				_ = _1855_v157
				_1855_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1847_v151).Fm21(p1, globalState), (_dafny.IntOfUint32((_1839_v144).Cardinality())).Plus(_this.F30))
				(globalState).F10 = (_1855_v157).Cardinality()
				var _1856_v158 _dafny.Array
				_ = _1856_v158
				var _nw332 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(22))
				_ = _nw332
				_1856_v158 = _nw332
				var _1857_v159 _dafny.Array
				_ = _1857_v159
				_1857_v159 = _1856_v158
				var _1858_v160 _dafny.Map
				_ = _1858_v160
				_1858_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1856_v158)
				var _1859_v161 _dafny.Array
				_ = _1859_v161
				var _nwElement0_80 _dafny.Array = _1856_v158
				_ = _nwElement0_80
				var _nw333 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(12))
				_ = _nw333
				(_nw333).ArraySet1(_nwElement0_80, 0)
				(_nw333).ArraySet1(_1856_v158, 1)
				(_nw333).ArraySet1(_1856_v158, 2)
				(_nw333).ArraySet1(_1856_v158, 3)
				(_nw333).ArraySet1(_1856_v158, 4)
				(_nw333).ArraySet1(_1856_v158, 5)
				(_nw333).ArraySet1(_1856_v158, 6)
				(_nw333).ArraySet1(_1856_v158, 7)
				(_nw333).ArraySet1((_1857_v159), 8)
				(_nw333).ArraySet1((func() _dafny.Array {
					if (_1858_v160).Contains(p2) {
						return (_1858_v160).Get(p2).(_dafny.Array)
					}
					return _1856_v158
				})(), 9)
				(_nw333).ArraySet1(_1856_v158, 10)
				(_nw333).ArraySet1(_1856_v158, 11)
				_1859_v161 = _nw333
				var _1860_v162 T1
				_ = _1860_v162
				var _nw334 *C4 = New_C4_()
				_ = _nw334
				_nw334.Ctor__(_1859_v161)
				_1860_v162 = _nw334
				_1860_v162 = _1860_v162
			} else {
				(globalState).F7 = p2
				var _1861_v163 *C1
				_ = _1861_v163
				var _nw335 *C1 = New_C1_()
				_ = _nw335
				_nw335.Ctor__()
				_1861_v163 = _nw335
				(globalState).F22 = p2
				var _1862_v164 _dafny.Sequence
				_ = _1862_v164
				_1862_v164 = _dafny.SeqOf(p2)
				var _1863_v165 _dafny.CodePoint
				_ = _1863_v165
				_1863_v165 = _dafny.CodePoint('y')
				_1672_v17 = (Companion_Default___.Fm17(_this.F30, (_1862_v164).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1862_v164).Cardinality()))).Uint32()).(bool), _1863_v165, globalState)).Intersection(_1672_v17)
				var _1864_v166 *C0
				_ = _1864_v166
				var _nw336 *C0 = New_C0_()
				_ = _nw336
				_nw336.Ctor__((func() bool {
					if true {
						return p2
					}
					return p2
				})())
				_1864_v166 = _nw336
			}
		}
		var _1865_i15 _dafny.Int
		_ = _1865_i15
		_1865_i15 = _dafny.Zero
		{
			for p2 {
				{
					if (_1865_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1865_i15 = (_1865_i15).Plus(_dafny.One)
					(globalState).F22 = (func() bool {
						if (_1675_v19).Contains(p2) {
							return (_1675_v19).Get(p2).(bool)
						}
						return false
					})()
					(globalState).F22 = p2
					if p2 {
						(globalState).F10 = ((p0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p0).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_this.F30)
						(globalState).F10 = Companion_Default___.SafeModuloInt(p1, (_dafny.Zero).Minus(_this.F30))
						var _rhs344 _dafny.Array = _1677_v21
						_ = _rhs344
						var _rhs345 _dafny.Int = Companion_Default___.Fm0((p1).Plus(p1), globalState)
						_ = _rhs345
						var _lhs313 *GlobalState = globalState
						_ = _lhs313
						_1677_v21 = _rhs344
						_lhs313.F10 = _rhs345
						var _1866_v167 *C3
						_ = _1866_v167
						var _nw337 *C3 = New_C3_()
						_ = _nw337
						_nw337.Ctor__()
						_1866_v167 = _nw337
						var _1867_v168 _dafny.Sequence
						_ = _1867_v168
						_1867_v168 = _dafny.UnicodeSeqOfUtf8Bytes("kdqjxd")
						var _1868_v169 *C2
						_ = _1868_v169
						var _nw338 *C2 = New_C2_()
						_ = _nw338
						_nw338.Ctor__(_1867_v168)
						_1868_v169 = _nw338
					} else {
						_1677_v21 = _1677_v21
						var _1869_v170 _dafny.Array
						_ = _1869_v170
						var _nwElement0_81 _dafny.Array = _1677_v21
						_ = _nwElement0_81
						var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(10))
						_ = _nw339
						(_nw339).ArraySet1(_nwElement0_81, 0)
						(_nw339).ArraySet1((func() _dafny.Array {
							if p2 {
								return _1677_v21
							}
							return _1677_v21
						})(), 1)
						(_nw339).ArraySet1(_1677_v21, 2)
						(_nw339).ArraySet1(_1677_v21, 3)
						(_nw339).ArraySet1(_1677_v21, 4)
						(_nw339).ArraySet1((func() _dafny.Array {
							if p2 {
								return _1677_v21
							}
							return _1677_v21
						})(), 5)
						(_nw339).ArraySet1(_1677_v21, 6)
						(_nw339).ArraySet1(_1677_v21, 7)
						(_nw339).ArraySet1(_1677_v21, 8)
						(_nw339).ArraySet1(_1677_v21, 9)
						_1869_v170 = _nw339
						_1869_v170 = _1869_v170
						var _1870_v171 _dafny.Array
						_ = _1870_v171
						var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.One)
						_ = _nw340
						_1870_v171 = _nw340
						_1870_v171 = _1870_v171
						var _1871_v172 _dafny.CodePoint
						_ = _1871_v172
						_1871_v172 = _dafny.CodePoint('o')
						var _1872_v173 D8
						_ = _1872_v173
						_1872_v173 = Companion_D8_.Create_DC22_(p1, !(p2), _1871_v172)
						var _1873_v174 _dafny.Set
						_ = _1873_v174
						_1873_v174 = _dafny.SetOf(_1872_v173)
						var _1874_v175 _dafny.Sequence
						_ = _1874_v175
						_1874_v175 = _dafny.UnicodeSeqOfUtf8Bytes("lfkqt")
						_1873_v174 = ((func() _dafny.Set {
							if (_this).Fm4(_1874_v175, p1, globalState) {
								return _1873_v174
							}
							return _1873_v174
						})()).Union(_1873_v174)
						_1651_v0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_this.F30), p1)).Merge(_1651_v0)
					}
					var _1875_v176 *C1
					_ = _1875_v176
					var _nw341 *C1 = New_C1_()
					_ = _nw341
					_nw341.Ctor__()
					_1875_v176 = _nw341
					var _1876_v177 _dafny.Map
					_ = _1876_v177
					_1876_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1875_v176)
					if ((_1876_v177).Merge(_1876_v177)).Contains(p2) {
						var _1877_v178 _dafny.Map
						_ = _1877_v178
						_1877_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.UnicodeSeqOfUtf8Bytes("q"))
						var _1878_v179 _dafny.Sequence
						_ = _1878_v179
						_1878_v179 = _dafny.UnicodeSeqOfUtf8Bytes("djlw")
						_1877_v178 = (_1877_v178).Update(p2, _dafny.Companion_Sequence_.Concatenate(_1878_v179, _1878_v179))
						r1 = !(p2)
						var _1879_v180 _dafny.Array
						_ = _1879_v180
						var _nwElement0_82 _dafny.Sequence = _1878_v179
						_ = _nwElement0_82
						var _nw342 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(7))
						_ = _nw342
						(_nw342).ArraySet1(_nwElement0_82, 0)
						(_nw342).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1878_v179, _1878_v179), 1)
						(_nw342).ArraySet1(_1878_v179, 2)
						(_nw342).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("vxtcqx"), 3)
						(_nw342).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1878_v179, _1878_v179), 4)
						(_nw342).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1878_v179, _1878_v179), 5)
						(_nw342).ArraySet1(Companion_Default___.Fm23(Companion_D2_.Create_DC7_(_1651_v0), _this.F30, (_dafny.Zero).Minus(_this.F30), globalState), 6)
						_1879_v180 = _nw342
						var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))
						_ = _index338
						(_1879_v180).ArraySet1(_1878_v179, (_index338).Int())
						var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))
						_ = _index339
						(_1879_v180).ArraySet1(_1878_v179, (_index339).Int())
						var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1677_v21), 0))
						_ = _index340
						(_1677_v21).ArraySet1(Companion_Default___.SafeModuloInt(_this.F30, _dafny.IntOfInt64(-430)), (_index340).Int())
						var _1880_v181 _dafny.Map
						_ = _1880_v181
						_1880_v181 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1875_v176, _1677_v21)
						var _1881_v182 _dafny.Array
						_ = _1881_v182
						var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
						_ = _nw343
						_1881_v182 = _nw343
						var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1881_v182), 0))
						_ = _index341
						(_1881_v182).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1878_v179, _1878_v179, (_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence))), (_index341).Int())
						var _1882_v183 *C0
						_ = _1882_v183
						var _nw344 *C0 = New_C0_()
						_ = _nw344
						_nw344.Ctor__(true)
						_1882_v183 = _nw344
						var _1883_v184 _dafny.Sequence
						_ = _1883_v184
						_1883_v184 = _dafny.SeqOf((_1882_v183).F32(), (_1882_v183).F32())
						var _1884_v185 D8
						_ = _1884_v185
						_1884_v185 = Companion_D8_.Create_DC21_(_1883_v184)
						var _1885_v186 D8
						_ = _1885_v186
						_1885_v186 = Companion_D8_.Create_DC24_(_1884_v185)
						var _1886_v187 _dafny.Map
						_ = _1886_v187
						_1886_v187 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_v183, _1885_v186)
						var _1887_v188 _dafny.Map
						_ = _1887_v188
						_1887_v188 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
						var _1888_v189 _dafny.Map
						_ = _1888_v189
						_1888_v189 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1886_v187).Merge(_1886_v187), _dafny.MultiSetOf(Companion_Default___.Fm23(_1673_v18, (_dafny.Zero).Minus((_1887_v188).Cardinality()), p1, globalState), (_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence), (_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence)))
						var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1677_v21), 0))
						_ = _index342
						var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1881_v182), 0))
						_ = _index343
						var _rhs346 _dafny.Int = p1
						_ = _rhs346
						var _rhs347 bool = (_dafny.SetOf(p1)).IsDisjointFrom(Companion_Default___.Fm42((_dafny.Zero).Minus(_this.F30), _dafny.SetOf(p2, !(p2)), p1, globalState))
						_ = _rhs347
						var _rhs348 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
							if p2 {
								return (func() _dafny.Sequence {
									if (_1877_v178).Contains(p2) {
										return (_1877_v178).Get(p2).(_dafny.Sequence)
									}
									return Companion_Default___.Fm23(_1673_v18, p1, _this.F30, globalState)
								})()
							}
							return (_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence)
						})()).Cardinality())
						_ = _rhs348
						var _rhs349 _dafny.Map = _1880_v181
						_ = _rhs349
						var _rhs350 _dafny.MultiSet = (func() _dafny.MultiSet {
							if (_1888_v189).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_v183, _1885_v186)).Merge(_1886_v187)) {
								return (_1888_v189).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_v183, _1885_v186)).Merge(_1886_v187)).(_dafny.MultiSet)
							}
							return (_dafny.MultiSetOf((_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence), (_1879_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1879_v180), 0))).Int()).(_dafny.Sequence))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(883))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg98 _dafny.Int) interface{} {
									return coer98(arg98)
								}
							}((func(_1889_v180 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
								return func(_1890_i16 _dafny.Int) _dafny.Sequence {
									return (_1889_v180).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_1889_v180), 0))).Int()).(_dafny.Sequence)
								}
							})(_1879_v180)))))
						})()
						_ = _rhs350
						var _lhs314 *GlobalState = globalState
						_ = _lhs314
						var _lhs315 *GlobalState = globalState
						_ = _lhs315
						var _lhs316 _dafny.Array = _1677_v21
						_ = _lhs316
						var _lhs317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(475), _dafny.ArrayLen((_1677_v21), 0))
						_ = _lhs317
						var _lhs318 _dafny.Array = _1881_v182
						_ = _lhs318
						var _lhs319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(260), _dafny.ArrayLen((_1881_v182), 0))
						_ = _lhs319
						_lhs314.F10 = _rhs346
						_lhs315.F7 = _rhs347
						(_lhs316).ArraySet1(_rhs348, (_lhs317).Int())
						_1880_v181 = _rhs349
						(_lhs318).ArraySet1(_rhs350, (_lhs319).Int())
						(globalState).F22 = p2
					} else {
						(globalState).F13 = p2
						var _1891_v190 _dafny.Map
						_ = _1891_v190
						_1891_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F30, p2)
						var _1892_v191 D2
						_ = _1892_v191
						_1892_v191 = Companion_D2_.Create_DC8_(p2, !((func() bool {
							if (_1891_v190).Contains(p1) {
								return (_1891_v190).Get(p1).(bool)
							}
							return true
						})()))
						_1892_v191 = _1892_v191
						var _1893_v192 _dafny.Sequence
						_ = _1893_v192
						_1893_v192 = _dafny.SeqOf(Companion_Default___.Fm12(p2, p2, globalState), p2)
						var _1894_v193 _dafny.Sequence
						_ = _1894_v193
						_1894_v193 = _dafny.SeqOf(_1651_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1893_v192)).Cardinality(), _this.F30))
						var _1895_v194 _dafny.Map
						_ = _1895_v194
						_1895_v194 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1894_v193)
						var _1896_v195 _dafny.Map
						_ = _1896_v195
						_1896_v195 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F30)
						_1895_v194 = (_1895_v194).Update((func() _dafny.Int {
							if (_1896_v195).Contains(!(false)) {
								return (_1896_v195).Get(!(false)).(_dafny.Int)
							}
							return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqoaxevgx")).Cardinality()))).Cardinality())
						})(), _1894_v193)
						(_1875_v176).M0(_dafny.IntOfInt64(436), _dafny.Companion_Sequence_.Contains(p0, p1), _this.F30, p2, globalState)
						(globalState).F10 = _dafny.IntOfInt64(-389)
					}
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1897_v196 _dafny.Sequence
		_ = _1897_v196
		_1897_v196 = _dafny.SeqOf(!(!(p2)), p2, p2)
		var _1898_v197 _dafny.Sequence
		_ = _1898_v197
		_1898_v197 = _dafny.UnicodeSeqOfUtf8Bytes("htgli")
		var _1899_v198 D0
		_ = _1899_v198
		_1899_v198 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((p0).Cardinality()), p2)
		var _1900_v199 _dafny.Array
		_ = _1900_v199
		var _nw345 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw345
		_1900_v199 = _nw345
		var _1901_v200 D0
		_ = _1901_v200
		_1901_v200 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1897_v196, _1897_v196)).Cardinality()), p1, (_this).Fm3(_dafny.IntOfUint32((_1898_v197).Cardinality()), p2, (_this).Fm2(globalState), _1899_v198, globalState), _1900_v199)
		r0 = _1901_v200
		r1 = p2
		return r0, r1
	}
}
func (_this *C5) F29() _dafny.Array {
	{
		return _this._f29
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
