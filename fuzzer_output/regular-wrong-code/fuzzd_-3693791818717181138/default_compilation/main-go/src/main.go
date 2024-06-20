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
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false, !(false), false)).Cardinality()), _dafny.IntOfInt64(-825), (_dafny.Zero).Minus(_dafny.IntOfInt64(-343)), _dafny.IntOfInt64(185))).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("piggfikm")).Cardinality())) > 0 {
		return _dafny.UnicodeSeqOfUtf8Bytes("rehg")
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hm"), _dafny.UnicodeSeqOfUtf8Bytes("clvp"))
	}
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("bxb"), _dafny.UnicodeSeqOfUtf8Bytes("p")) {
		return _dafny.CodePoint('t')
	} else {
		return _dafny.CodePoint('i')
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kivrkb"), _dafny.UnicodeSeqOfUtf8Bytes("fqfafs"))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Map, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC11_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(231), true)).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-259), !(true))).Cardinality(), _dafny.IntOfInt64(54), _dafny.IntOfInt64(207), _dafny.IntOfInt64(841), (_dafny.SetOf(_dafny.IntOfInt64(774))).Cardinality())).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-259), !(true))).Cardinality(), _dafny.IntOfInt64(54), _dafny.IntOfInt64(207), _dafny.IntOfInt64(841), (_dafny.SetOf(_dafny.IntOfInt64(774))).Cardinality()), _0_v0) {
				_coll0.Add((_0_v0).Minus(_dafny.IntOfInt64(998)), !(!(!(!(false)))))
			}
		}
		return _coll0.ToMap()
	}())).Cardinality(), (Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ekmhsfs"))).Dtor_cf10())
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 D3, globalState *GlobalState) _dafny.Sequence {
	return (Companion_D16_.Create_DC41_(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tfnuqeu")))).Dtor_cf72()
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(240))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjdopi")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), (_dafny.SetOf(_dafny.IntOfInt64(100))).Cardinality())).Cardinality()), _dafny.IntOfInt64(-881), _dafny.IntOfInt64(473), (_dafny.SetOf(true)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC4_(true)
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() D4 {
		if true {
			return Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("gwsfpeakx"))
		}
		return Companion_D4_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		})))
	})())
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) D4 {
	return Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ysmgae"))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yriwev")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(908)), !(true))).Cardinality(), _dafny.IntOfInt64(640))).Intersection(_dafny.MultiSetOf(_dafny.IntOfInt64(182), _dafny.IntOfUint32((_dafny.SeqOf(true, !(!(true)), false, true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(false, true)).Difference(_dafny.MultiSetOf(true))).Difference(_dafny.MultiSetOf(true, !(false), !(true), false, false))
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(589), _dafny.IntOfInt64(859))); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_1).(_dafny.Int)
			if ((_dafny.IntOfInt64(589)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(859)) < 0) {
				_coll1.Add((_2_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqxdyp")).Cardinality())), true)
			}
		}
		return _coll1.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-882))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _3_v0 _dafny.Map
			_3_v0 = interface{}(_compr_2).(_dafny.Map)
			if (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(-882))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))).Contains(_3_v0) {
				_coll2.Add(_3_v0, (Companion_D3_.Create_DC6_(!(true), _dafny.IntOfInt64(563), _dafny.IntOfInt64(882))).Dtor_cf8())
			}
		}
		return _coll2.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge((Companion_D10_.Create_DC24_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Dtor_cf42())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	if !(!(!(!(true)))) {
		return (func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.IntOfInt64(152))).Cardinality())).Cardinality(), _dafny.IntOfInt64(114), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(322))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()))).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _5_v0 _dafny.Int
				_5_v0 = interface{}(_compr_3).(_dafny.Int)
				if (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.IntOfInt64(152))).Cardinality())).Cardinality(), _dafny.IntOfInt64(114), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(322))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg2 _dafny.Int) interface{} {
						return coer2(arg2)
					}
				}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}))).Cardinality()))).Contains(_5_v0) {
					_coll3.Add((_5_v0).Minus((_dafny.SetOf(!(true))).Cardinality()))
				}
			}
			return _coll3.ToSet()
		}()).Intersection(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chjhmkgm")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfInt64(875))).Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _6_v1 _dafny.Int
				_6_v1 = interface{}(_compr_4).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("chjhmkgm")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality(), _dafny.IntOfInt64(875)), _6_v1) {
					_coll4.Add((_6_v1).Plus(_dafny.IntOfInt64(-795)))
				}
			}
			return _coll4.ToSet()
		}())
	} else {
		return _dafny.SetOf((func() _dafny.Set {
			var _coll5 = _dafny.NewBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(887), _dafny.IntOfInt64(374))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _7_v2 _dafny.Int
				_7_v2 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(887)).Cmp(_7_v2) <= 0) && ((_7_v2).Cmp(_dafny.IntOfInt64(374)) < 0) {
					_coll5.Add((_7_v2).Times((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.MultiSetOf(false)).Cardinality()))).Cardinality()))
				}
			}
			return _coll5.ToSet()
		}()).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC16_(Companion_D7_.Create_DC15_())
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D12 {
	return Companion_D12_.Create_DC30_(_dafny.IntOfInt64(378), (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(true)))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((func() bool {
		if !(true) {
			return true
		}
		return false
	})())
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	var _source0 D8 = Companion_D8_.Create_DC18_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.IntOfInt64(798))).Cardinality()), _dafny.IntOfInt64(95))).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(702))).Cardinality()))).Cardinality())), true, _dafny.IntOfInt64(138))
	_ = _source0
	if _source0.Is_DC18() {
		var _8___mcc_h0 _dafny.Map = _source0.Get_().(D8_DC18).Cf28
		_ = _8___mcc_h0
		var _9___mcc_h1 bool = _source0.Get_().(D8_DC18).Cf29
		_ = _9___mcc_h1
		var _10___mcc_h2 _dafny.Int = _source0.Get_().(D8_DC18).Cf30
		_ = _10___mcc_h2
		var _11_cf30 _dafny.Int = _10___mcc_h2
		_ = _11_cf30
		var _12_cf29 bool = _9___mcc_h1
		_ = _12_cf29
		var _13_cf28 _dafny.Map = _8___mcc_h0
		_ = _13_cf28
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_12_cf29), _dafny.SeqOf(_12_cf29, _12_cf29, _12_cf29))
	} else if _source0.Is_DC19() {
		var _14___mcc_h3 bool = _source0.Get_().(D8_DC19).Cf31
		_ = _14___mcc_h3
		var _15_cf31 bool = _14___mcc_h3
		_ = _15_cf31
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_15_cf31), _dafny.SeqOf(_15_cf31, _15_cf31))
	} else {
		var _16___mcc_h4 _dafny.Set = _source0.Get_().(D8_DC17).Cf27
		_ = _16___mcc_h4
		var _17_cf27 _dafny.Set = _16___mcc_h4
		_ = _17_cf27
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, !(true)), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) M4(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) (_dafny.Int, bool) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var _18_v0 _dafny.MultiSet
	_ = _18_v0
	_18_v0 = _dafny.MultiSetOf(p0)
	var _19_v1 _dafny.Map
	_ = _19_v1
	_19_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
	var _20_v2 T1
	_ = _20_v2
	var _nw0 *C3 = New_C3_()
	_ = _nw0
	_nw0.Ctor__((func() bool {
		if (_19_v1).Contains(p1) {
			return (_19_v1).Get(p1).(bool)
		}
		return p1
	})(), p1)
	_20_v2 = _nw0
	var _21_v3 _dafny.Set
	_ = _21_v3
	_21_v3 = _dafny.SetOf(_20_v2)
	var _22_v4 _dafny.Sequence
	_ = _22_v4
	_22_v4 = _dafny.SeqOf(_20_v2.F8())
	var _23_v5 _dafny.Sequence
	_ = _23_v5
	_23_v5 = _dafny.SeqOf(_dafny.IntOfUint32((_22_v4).Cardinality()))
	var _24_v6 D0
	_ = _24_v6
	_24_v6 = Companion_D0_.Create_DC1_(p1, _dafny.SeqOf(p1, (_22_v4).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_22_v4).Cardinality()))).Uint32()).(bool)), false)
	var _25_v7 T2
	_ = _25_v7
	var _nw1 *C3 = New_C3_()
	_ = _nw1
	_nw1.Ctor__(p1, p1)
	_25_v7 = _nw1
	var _26_v8 _dafny.Map
	_ = _26_v8
	_26_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_25_v7, p0)
	var _27_v9 D9
	_ = _27_v9
	_27_v9 = Companion_D9_.Create_DC22_(p0, _dafny.IntOfUint32((_23_v5).Cardinality()), p1, (_dafny.Zero).Minus((_20_v2).Fm2(_24_v6, globalState)), _26_v8)
	var _28_v10 _dafny.MultiSet
	_ = _28_v10
	_28_v10 = _dafny.MultiSetOf((func() _dafny.Int {
		if (_18_v0).Contains((_21_v3).Cardinality()) {
			return (_18_v0).Multiplicity((_21_v3).Cardinality())
		}
		return _dafny.IntOfInt64(-700)
	})(), (_23_v5).Select((Companion_Default___.SafeIndex((_27_v9).Dtor_cf35(), _dafny.IntOfUint32((_23_v5).Cardinality()))).Uint32()).(_dafny.Int))
	var _29_v11 _dafny.Array
	_ = _29_v11
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
	_ = _nw2
	_29_v11 = _nw2
	var _30_v12 _dafny.Set
	_ = _30_v12
	_30_v12 = _dafny.SetOf(p0)
	var _31_v13 _dafny.Sequence
	_ = _31_v13
	_31_v13 = _dafny.SeqOf(_30_v12, _30_v12, _30_v12, _30_v12)
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_29_v11), 0))
	_ = _index0
	(_29_v11).ArraySet1((_31_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(545), _dafny.IntOfUint32((_31_v13).Cardinality()))).Uint32()).(_dafny.Set), (_index0).Int())
	var _32_v14 _dafny.Map
	_ = _32_v14
	_32_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_25_v7).F9(), p0)
	var _33_v15 D5
	_ = _33_v15
	_33_v15 = Companion_D5_.Create_DC11_((_32_v14).Cardinality(), p2)
	var _34_v16 _dafny.CodePoint
	_ = _34_v16
	_34_v16 = _dafny.CodePoint('m')
	var _pat_let_tv0 = p0
	_ = _pat_let_tv0
	var _pat_let_tv1 = _34_v16
	_ = _pat_let_tv1
	var _35_v17 _dafny.Array
	_ = _35_v17
	var _nwElement0_0 _dafny.Int = p0
	_ = _nwElement0_0
	var _nw3 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(20))
	_ = _nw3
	(_nw3).ArraySet1(_nwElement0_0, 0)
	(_nw3).ArraySet1(p0, 1)
	(_nw3).ArraySet1((_dafny.IntOfInt64(772)).Plus(p0), 2)
	(_nw3).ArraySet1(p0, 3)
	(_nw3).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("en")).Cardinality()), 4)
	(_nw3).ArraySet1((_dafny.Zero).Minus((p0).Minus(p0)), 5)
	(_nw3).ArraySet1(p0, 6)
	(_nw3).ArraySet1(p0, 7)
	(_nw3).ArraySet1((p0).Plus(p0), 8)
	(_nw3).ArraySet1(p0, 9)
	(_nw3).ArraySet1(_dafny.IntOfInt64(141), 10)
	(_nw3).ArraySet1((_20_v2).Fm2(_24_v6, globalState), 11)
	(_nw3).ArraySet1((_23_v5).Select((Companion_Default___.SafeIndex((func(_pat_let0_0 D5) D5 {
		return func(_36_dt__update__tmp_h0 D5) D5 {
			return func(_pat_let1_0 _dafny.Sequence) D5 {
				return func(_39_dt__update_hcf19_h0 _dafny.Sequence) D5 {
					return Companion_D5_.Create_DC11_((_36_dt__update__tmp_h0).Dtor_cf18(), _39_dt__update_hcf19_h0)
				}(_pat_let1_0)
			}(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-596))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg3 _dafny.Int) interface{} {
					return coer3(arg3)
				}
			}(func(_37_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), (Companion_Default___.SafeIndex(_pat_let_tv0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-596))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))).Cardinality()))).Uint32(), _pat_let_tv1))
		}(_pat_let0_0)
	}(_33_v15)).Dtor_cf18(), _dafny.IntOfUint32((_23_v5).Cardinality()))).Uint32()).(_dafny.Int), 12)
	(_nw3).ArraySet1(p0, 13)
	(_nw3).ArraySet1(p0, 14)
	(_nw3).ArraySet1(p0, 15)
	(_nw3).ArraySet1((p0).Plus(p0), 16)
	(_nw3).ArraySet1(p0, 17)
	(_nw3).ArraySet1(p0, 18)
	(_nw3).ArraySet1((p0).Times(p0), 19)
	_35_v17 = _nw3
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_29_v11), 0))
	_ = _index1
	var _rhs0 _dafny.MultiSet = _dafny.MultiSetOf(p0)
	_ = _rhs0
	var _rhs1 _dafny.Set = (_30_v12).Intersection(_30_v12)
	_ = _rhs1
	var _rhs2 _dafny.Array = _35_v17
	_ = _rhs2
	var _rhs3 _dafny.Int = Companion_Default___.SafeModuloInt((_23_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_23_v5).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_25_v7).F9()), _22_v4)).Cardinality()))
	_ = _rhs3
	var _lhs0 _dafny.Array = _29_v11
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_29_v11), 0))
	_ = _lhs1
	_18_v0 = _rhs0
	(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
	_35_v17 = _rhs2
	r0 = _rhs3
	for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_35_v17), 0))); ; {
		_guard_loop_0, _ok6 := _iter6()
		if !_ok6 {
			break
		}
		var _40_i1 _dafny.Int
		_40_i1 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_40_i1).Sign() != -1) && ((_40_i1).Cmp(_dafny.ArrayLen((_35_v17), 0)) < 0)) {
			(_35_v17).ArraySet1((_40_i1).Times(p0), (_40_i1).Int())
		}
	}
	var _pat_let_tv2 = _25_v7
	_ = _pat_let_tv2
	(_20_v2).F8_set_(func(_source1 D0) bool {
		if _source1.Is_DC0() {
			return (_pat_let_tv2).F9()
		} else {
			var _41___mcc_h0 bool = _source1.Get_().(D0_DC1).Cf0
			_ = _41___mcc_h0
			var _42___mcc_h1 _dafny.Sequence = _source1.Get_().(D0_DC1).Cf1
			_ = _42___mcc_h1
			var _43___mcc_h2 bool = _source1.Get_().(D0_DC1).Cf2
			_ = _43___mcc_h2
			var _44_cf2 bool = _43___mcc_h2
			_ = _44_cf2
			var _45_cf1 _dafny.Sequence = _42___mcc_h1
			_ = _45_cf1
			var _46_cf0 bool = _41___mcc_h0
			_ = _46_cf0
			return !(false)
		}
	}(_24_v6))
	var _hi0 _dafny.Int = (p0).Times(p0)
	_ = _hi0
	for _47_i2 := p0; _47_i2.Cmp(_hi0) < 0; _47_i2 = _47_i2.Plus(_dafny.One) {
		var _48_v18 _dafny.Sequence
		_ = _48_v18
		var _49_v19 _dafny.Int
		_ = _49_v19
		var _50_v20 bool
		_ = _50_v20
		var _out0 _dafny.Sequence
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 bool
		_ = _out2
		_out0, _out1, _out2 = (_25_v7).M1(_dafny.Companion_Sequence_.Concatenate(p2, p2), (_dafny.MultiSetOf(_20_v2.F8(), p1)).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf((_20_v2).Fm0(globalState), _20_v2.F8()))), globalState)
		_48_v18 = _out0
		_49_v19 = _out1
		_50_v20 = _out2
		var _51_v21 _dafny.Map
		_ = _51_v21
		_51_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v18, _20_v2.F8())
		r1 = ((_51_v21).Cardinality()).Cmp(p0) > 0
		var _52_v22 D4
		_ = _52_v22
		_52_v22 = Companion_D4_.Create_DC7_(Companion_Default___.Fm9((_25_v7).F9(), _20_v2.F8(), globalState))
		var _53_v23 _dafny.Sequence
		_ = _53_v23
		_53_v23 = _dafny.SeqOf(_52_v22, _52_v22, _52_v22, _52_v22, _52_v22)
		var _54_v24 *C1
		_ = _54_v24
		var _nw4 *C1 = New_C1_()
		_ = _nw4
		_nw4.Ctor__(_34_v16, _dafny.SetOf(_53_v23, _53_v23), _dafny.Companion_Sequence_.Contains(_23_v5, _47_i2), _20_v2.F8())
		_54_v24 = _nw4
		var _55_v25 *C3
		_ = _55_v25
		var _nw5 *C3 = New_C3_()
		_ = _nw5
		_nw5.Ctor__(_50_v20, _20_v2.F8())
		_55_v25 = _nw5
		var _nw6 *C3 = New_C3_()
		_ = _nw6
		_nw6.Ctor__(!(_20_v2.F8()) || (_50_v20), (func() bool {
			if (_25_v7).F9() {
				return _50_v20
			}
			return _20_v2.F8()
		})())
		_55_v25 = _nw6
	}
	var _56_v26 _dafny.Array
	_ = _56_v26
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(2)
	_ = _len0_0
	var _nw7 _dafny.Array
	_ = _nw7
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw7 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Set = (func(_57_v2 T1, _58_v7 T2) func(_dafny.Int) _dafny.Set {
			return func(_59_i4 _dafny.Int) _dafny.Set {
				return _dafny.SetOf(_57_v2.F8(), !(_57_v2.F8()), (_58_v7).F9())
			}
		})(_20_v2, _25_v7)
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
	_56_v26 = _nw7
	for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_56_v26), 0))); ; {
		_guard_loop_1, _ok7 := _iter7()
		if !_ok7 {
			break
		}
		var _60_i3 _dafny.Int
		_60_i3 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_60_i3).Sign() != -1) && ((_60_i3).Cmp(_dafny.ArrayLen((_56_v26), 0)) < 0)) {
			(_56_v26).ArraySet1(_dafny.SetOf((_25_v7).F9(), _20_v2.F8()), (_60_i3).Int())
		}
	}
	var _hi1 _dafny.Int = (_dafny.Zero).Minus((p0).Plus(p0))
	_ = _hi1
	for _61_i5 := (p0).Times(p0); _61_i5.Cmp(_hi1) < 0; _61_i5 = _61_i5.Plus(_dafny.One) {
		r0 = (_dafny.IntOfInt64(672)).Times(_dafny.IntOfUint32((p2).Cardinality()))
		var _62_v27 _dafny.MultiSet
		_ = _62_v27
		_62_v27 = _dafny.MultiSetOf(_20_v2.F8(), true)
		var _63_v28 _dafny.Map
		_ = _63_v28
		_63_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _62_v27)
		r0 = Companion_Default___.SafeModuloInt((_63_v28).Cardinality(), (_20_v2).Fm2(_24_v6, globalState))
		var _64_v29 D5
		_ = _64_v29
		_64_v29 = Companion_D5_.Create_DC10_(_35_v17)
		var _65_v30 _dafny.MultiSet
		_ = _65_v30
		_65_v30 = _dafny.MultiSetOf((_64_v29).Dtor_cf17())
		_65_v30 = (_65_v30).Union((_dafny.MultiSetOf(_35_v17, _35_v17)).Intersection(_65_v30))
		var _66_v31 _dafny.Array
		_ = _66_v31
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_1
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_67_v2 T1) func(_dafny.Int) bool {
				return func(_68_i6 _dafny.Int) bool {
					return !(_67_v2.F8())
				}
			})(_20_v2)
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
		_66_v31 = _nw8
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_66_v31), 0))
		_ = _index2
		(_66_v31).ArraySet1((_25_v7).F9(), (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_66_v31), 0))
		_ = _index3
		(_66_v31).ArraySet1(_20_v2.F8(), (_index3).Int())
	}
	r0 = p0
	var _69_v32 D12
	_ = _69_v32
	_69_v32 = Companion_D12_.Create_DC30_(p0, p0)
	var _70_v33 _dafny.Map
	_ = _70_v33
	_70_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D12_.Create_DC32_(_69_v32), p2)
	var _71_v34 _dafny.Set
	_ = _71_v34
	_71_v34 = _dafny.SetOf((func() _dafny.Sequence {
		if (_70_v33).Contains(Companion_D12_.Create_DC32_(_69_v32)) {
			return (_70_v33).Get(Companion_D12_.Create_DC32_(_69_v32)).(_dafny.Sequence)
		}
		return p2
	})())
	r1 = (func() bool {
		if _20_v2.F8() {
			return _20_v2.F8()
		}
		return (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("kl"))).IsDisjointFrom(_71_v34)
	})()
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _72_v0 _dafny.Sequence
	_ = _72_v0
	_72_v0 = _dafny.UnicodeSeqOfUtf8Bytes("xl")
	var _73_v1 _dafny.Array
	_ = _73_v1
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
	_ = _len0_2
	var _nw9 _dafny.Array
	_ = _nw9
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw9 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) bool = func(_74_i0 _dafny.Int) bool {
			return true
		}
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw9 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw9).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw9).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_73_v1 = _nw9
	var _75_globalState *GlobalState
	_ = _75_globalState
	var _nw10 *GlobalState = New_GlobalState_()
	_ = _nw10
	_nw10.Ctor__(_dafny.CodePoint('i'), _dafny.IntOfInt64(114), _dafny.IntOfInt64(-993), _dafny.IntOfInt64(495), _72_v0, _dafny.IntOfInt64(-616), _73_v1, _dafny.IntOfInt64(-825))
	_75_globalState = _nw10
	var _76_v2 bool
	_ = _76_v2
	_76_v2 = false
	var _77_v3 _dafny.Sequence
	_ = _77_v3
	_77_v3 = _dafny.SeqOf(_76_v2, _76_v2, _76_v2, _76_v2)
	var _pat_let_tv3 = _77_v3
	_ = _pat_let_tv3
	var _source2 D0 = func(_source3 D0) D0 {
		if _source3.Is_DC0() {
			return Companion_D0_.Create_DC0_()
		} else {
			var _78___mcc_h3 bool = _source3.Get_().(D0_DC1).Cf0
			_ = _78___mcc_h3
			var _79___mcc_h4 _dafny.Sequence = _source3.Get_().(D0_DC1).Cf1
			_ = _79___mcc_h4
			var _80___mcc_h5 bool = _source3.Get_().(D0_DC1).Cf2
			_ = _80___mcc_h5
			var _81_cf2 bool = _80___mcc_h5
			_ = _81_cf2
			var _82_cf1 _dafny.Sequence = _79___mcc_h4
			_ = _82_cf1
			var _83_cf0 bool = _78___mcc_h3
			_ = _83_cf0
			return Companion_D0_.Create_DC0_()
		}
	}(func(_pat_let2_0 D0) D0 {
		return func(_84_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let3_0 _dafny.Sequence) D0 {
				return func(_85_dt__update_hcf1_h0 _dafny.Sequence) D0 {
					return Companion_D0_.Create_DC1_((_84_dt__update__tmp_h0).Dtor_cf0(), _85_dt__update_hcf1_h0, (_84_dt__update__tmp_h0).Dtor_cf2())
				}(_pat_let3_0)
			}(_pat_let_tv3)
		}(_pat_let2_0)
	}(Companion_D0_.Create_DC1_(_76_v2, _77_v3, _76_v2)))
	_ = _source2
	if _source2.Is_DC0() {
		if _76_v2 {
			var _86_v4 *C3
			_ = _86_v4
			var _nw11 *C3 = New_C3_()
			_ = _nw11
			_nw11.Ctor__(_76_v2, _76_v2)
			_86_v4 = _nw11
			var _87_v5 _dafny.CodePoint
			_ = _87_v5
			_87_v5 = _dafny.CodePoint('p')
			var _rhs4 _dafny.CodePoint = _dafny.CodePoint('c')
			_ = _rhs4
			var _rhs5 bool = _76_v2
			_ = _rhs5
			var _rhs6 _dafny.Sequence = _77_v3
			_ = _rhs6
			var _rhs7 _dafny.CodePoint = _87_v5
			_ = _rhs7
			_87_v5 = _rhs4
			_76_v2 = _rhs5
			_77_v3 = _rhs6
			_87_v5 = _rhs7
			var _88_v6 _dafny.Array
			_ = _88_v6
			var _nwElement0_1 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_72_v0, _72_v0)
			_ = _nwElement0_1
			var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(6))
			_ = _nw12
			(_nw12).ArraySet1(_nwElement0_1, 0)
			(_nw12).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fblhxtr"), 1)
			(_nw12).ArraySet1(_72_v0, 2)
			(_nw12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_72_v0, _72_v0), 3)
			(_nw12).ArraySet1(_72_v0, 4)
			(_nw12).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("whybwpehq"), 5)
			_88_v6 = _nw12
			var _89_v7 _dafny.Map
			_ = _89_v7
			_89_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _72_v0)
			var _90_v8 _dafny.Int
			_ = _90_v8
			_90_v8 = _dafny.IntOfInt64(873)
			var _91_v9 _dafny.Sequence
			_ = _91_v9
			_91_v9 = _dafny.SeqOf(_90_v8)
			var _92_v10 _dafny.Map
			_ = _92_v10
			_92_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_91_v9).Cardinality()), _72_v0)
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_88_v6), 0))
			_ = _index4
			(_88_v6).ArraySet1((func() _dafny.Sequence {
				if (_89_v7).Contains(_76_v2) {
					return (_89_v7).Get(_76_v2).(_dafny.Sequence)
				}
				return (func() _dafny.Sequence {
					if (_92_v10).Contains(_90_v8) {
						return (_92_v10).Get(_90_v8).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(621))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg5 _dafny.Int) interface{} {
							return coer5(arg5)
						}
					}((func(_93_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_94_i1 _dafny.Int) _dafny.CodePoint {
							return _93_v5
						}
					})(_87_v5)))
				})()
			})(), (_index4).Int())
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_88_v6), 0))
			_ = _index5
			(_88_v6).ArraySet1(_72_v0, (_index5).Int())
			var _95_v11 D9
			_ = _95_v11
			_95_v11 = Companion_D9_.Create_DC21_(_90_v8, _90_v8)
			var _96_v12 _dafny.Map
			_ = _96_v12
			_96_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_95_v11, _90_v8)
			var _97_v13 D0
			_ = _97_v13
			_97_v13 = Companion_D0_.Create_DC1_(_76_v2, _77_v3, _76_v2)
			var _98_v14 _dafny.Array
			_ = _98_v14
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw13
			_98_v14 = _nw13
			var _99_v15 T0
			_ = _99_v15
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_98_v14)
			_99_v15 = _nw14
			var _100_v16 _dafny.Map
			_ = _100_v16
			_100_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_72_v0).Cardinality()), _99_v15)
			var _101_v17 _dafny.Sequence
			_ = _101_v17
			_101_v17 = _dafny.SeqOf(_100_v16)
			var _102_v18 _dafny.Set
			_ = _102_v18
			_102_v18 = _dafny.SetOf(_73_v1, _73_v1)
			var _103_v19 _dafny.Map
			_ = _103_v19
			_103_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _90_v8)
			var _104_v20 _dafny.Map
			_ = _104_v20
			_104_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v8, (_86_v4).Fm2(_97_v13, _75_globalState))
			var _105_v21 _dafny.Map
			_ = _105_v21
			_105_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_90_v8, true)
			var _106_v22 _dafny.Map
			_ = _106_v22
			_106_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, !(_76_v2))
			var _107_v23 _dafny.MultiSet
			_ = _107_v23
			_107_v23 = _dafny.MultiSetOf(_106_v22, _106_v22)
			var _108_v24 D5
			_ = _108_v24
			_108_v24 = Companion_D5_.Create_DC11_((_107_v23).Cardinality(), (_88_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_88_v6), 0))).Int()).(_dafny.Sequence))
			var _109_v25 _dafny.Array
			_ = _109_v25
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_3
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_110_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_111_i2 _dafny.Int) _dafny.CodePoint {
						return _110_v5
					}
				})(_87_v5)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw15 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw15).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw15).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_109_v25 = _nw15
			var _112_v26 _dafny.Map
			_ = _112_v26
			_112_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_109_v25, _76_v2)
			var _pat_let_tv4 = _90_v8
			_ = _pat_let_tv4
			var _pat_let_tv5 = _90_v8
			_ = _pat_let_tv5
			var _113_v27 _dafny.Array
			_ = _113_v27
			var _nwElement0_2 _dafny.Int = _90_v8
			_ = _nwElement0_2
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(29))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_2, 0)
			(_nw16).ArraySet1(_90_v8, 1)
			(_nw16).ArraySet1(_90_v8, 2)
			(_nw16).ArraySet1(_90_v8, 3)
			(_nw16).ArraySet1(_90_v8, 4)
			(_nw16).ArraySet1(_90_v8, 5)
			(_nw16).ArraySet1((func() _dafny.Int {
				if (_96_v12).Contains(func(_pat_let4_0 D9) D9 {
					return func(_114_dt__update__tmp_h1 D9) D9 {
						return func(_pat_let5_0 _dafny.Int) D9 {
							return func(_115_dt__update_hcf33_h0 _dafny.Int) D9 {
								return Companion_D9_.Create_DC21_(_115_dt__update_hcf33_h0, (_114_dt__update__tmp_h1).Dtor_cf34())
							}(_pat_let5_0)
						}(_pat_let_tv4)
					}(_pat_let4_0)
				}(_95_v11)) {
					return (_96_v12).Get(func(_pat_let6_0 D9) D9 {
						return func(_116_dt__update__tmp_h2 D9) D9 {
							return func(_pat_let7_0 _dafny.Int) D9 {
								return func(_117_dt__update_hcf33_h1 _dafny.Int) D9 {
									return Companion_D9_.Create_DC21_(_117_dt__update_hcf33_h1, (_116_dt__update__tmp_h2).Dtor_cf34())
								}(_pat_let7_0)
							}(_pat_let_tv5)
						}(_pat_let6_0)
					}(_95_v11)).(_dafny.Int)
				}
				return _90_v8
			})(), 6)
			(_nw16).ArraySet1(Companion_Default___.SafeDivisionInt((_86_v4).Fm2(_97_v13, _75_globalState), _90_v8), 7)
			(_nw16).ArraySet1(Companion_Default___.SafeDivisionInt(_90_v8, (_86_v4).Fm2(Companion_D0_.Create_DC1_(_76_v2, _77_v3, _76_v2), _75_globalState)), 8)
			(_nw16).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_101_v17, _101_v17)).Cardinality()), 9)
			(_nw16).ArraySet1((_dafny.Zero).Minus((_102_v18).Cardinality()), 10)
			(_nw16).ArraySet1(_90_v8, 11)
			(_nw16).ArraySet1((func() _dafny.Int {
				if _76_v2 {
					return _90_v8
				}
				return _90_v8
			})(), 12)
			(_nw16).ArraySet1(_90_v8, 13)
			(_nw16).ArraySet1(_90_v8, 14)
			(_nw16).ArraySet1((Companion_Default___.Fm21((_103_v19).Cardinality(), _75_globalState)).Cardinality(), 15)
			(_nw16).ArraySet1((_90_v8).Plus(_90_v8), 16)
			(_nw16).ArraySet1((_86_v4).Fm2(_97_v13, _75_globalState), 17)
			(_nw16).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_91_v9, (Companion_Default___.SafeIndex(_90_v8, _dafny.IntOfUint32((_91_v9).Cardinality()))).Uint32(), _90_v8)).Cardinality()), 18)
			(_nw16).ArraySet1(_90_v8, 19)
			(_nw16).ArraySet1((func() _dafny.Int {
				if (_104_v20).Contains(_dafny.IntOfUint32((_72_v0).Cardinality())) {
					return (_104_v20).Get(_dafny.IntOfUint32((_72_v0).Cardinality())).(_dafny.Int)
				}
				return _90_v8
			})(), 20)
			(_nw16).ArraySet1((_dafny.Zero).Minus(((_105_v21).Cardinality()).Plus(_dafny.IntOfUint32((_77_v3).Cardinality()))), 21)
			(_nw16).ArraySet1((_dafny.IntOfInt64(-632)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _90_v8)).Cardinality()), 22)
			(_nw16).ArraySet1(_dafny.IntOfInt64(-302), 23)
			(_nw16).ArraySet1(_90_v8, 24)
			(_nw16).ArraySet1(_90_v8, 25)
			(_nw16).ArraySet1((func() _dafny.Int {
				if _76_v2 {
					return (_108_v24).Dtor_cf18()
				}
				return (_112_v26).Cardinality()
			})(), 26)
			(_nw16).ArraySet1((_dafny.Zero).Minus((_90_v8).Times(_90_v8)), 27)
			(_nw16).ArraySet1(_90_v8, 28)
			_113_v27 = _nw16
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_113_v27), 0))
			_ = _index6
			(_113_v27).ArraySet1((_dafny.Zero).Minus(_90_v8), (_index6).Int())
			var _118_v28 _dafny.MultiSet
			_ = _118_v28
			_118_v28 = _dafny.MultiSetOf((_86_v4).Fm2(_97_v13, _75_globalState), _dafny.IntOfInt64(334), _90_v8, Companion_Default___.SafeModuloInt(_90_v8, _90_v8), _90_v8)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_113_v27), 0))
			_ = _index7
			(_113_v27).ArraySet1((func() _dafny.Int {
				if (_118_v28).Contains((_86_v4).Fm2(Companion_D0_.Create_DC1_(_76_v2, _dafny.SeqOf(_76_v2), _76_v2), _75_globalState)) {
					return (_118_v28).Multiplicity((_86_v4).Fm2(Companion_D0_.Create_DC1_(_76_v2, _dafny.SeqOf(_76_v2), _76_v2), _75_globalState))
				}
				return _dafny.IntOfInt64(-339)
			})(), (_index7).Int())
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_73_v1), 0))
			_ = _index8
			(_73_v1).ArraySet1(_76_v2, (_index8).Int())
			var _119_v29 _dafny.MultiSet
			_ = _119_v29
			_119_v29 = _dafny.MultiSetOf((_99_v15).Fm0(_75_globalState))
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_73_v1), 0))
			_ = _index9
			(_73_v1).ArraySet1(((_113_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_113_v27), 0))).Int()).(_dafny.Int)).Cmp((func() _dafny.Int {
				if (_119_v29).Contains(_76_v2) {
					return (_119_v29).Multiplicity(_76_v2)
				}
				return _90_v8
			})()) >= 0, (_index9).Int())
		} else {
			var _120_v30 _dafny.Int
			_ = _120_v30
			_120_v30 = _dafny.IntOfInt64(-507)
			var _121_v31 _dafny.Sequence
			_ = _121_v31
			_121_v31 = _dafny.SeqOf((_dafny.Zero).Minus(_120_v30))
			var _122_v32 _dafny.CodePoint
			_ = _122_v32
			_122_v32 = _dafny.CodePoint('e')
			_72_v0 = _dafny.Companion_Sequence_.Update(_72_v0, (Companion_Default___.SafeIndex((_121_v31).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.IntOfUint32((_121_v31).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_72_v0).Cardinality()))).Uint32(), _122_v32)
			var _123_v33 _dafny.Map
			_ = _123_v33
			_123_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _73_v1)
			_73_v1 = (func() _dafny.Array {
				if (_123_v33).Contains(_76_v2) {
					return (_123_v33).Get(_76_v2).(_dafny.Array)
				}
				return _73_v1
			})()
			var _124_v34 _dafny.Set
			_ = _124_v34
			_124_v34 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(532))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}((func(_125_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_126_i3 _dafny.Int) _dafny.CodePoint {
					return _125_v32
				}
			})(_122_v32))), _72_v0)
			_124_v34 = _124_v34
			var _127_v35 _dafny.Map
			_ = _127_v35
			_127_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _dafny.IntOfUint32((_72_v0).Cardinality()))
			_127_v35 = (_127_v35).Update(_76_v2, _120_v30)
			_120_v30 = (_120_v30).Times(_120_v30)
		}
		var _128_v36 _dafny.Int
		_ = _128_v36
		_128_v36 = _dafny.IntOfInt64(115)
		_128_v36 = _128_v36
		var _129_v37 _dafny.Map
		_ = _129_v37
		_129_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(44), _76_v2)
		_129_v37 = (_129_v37).Update(_128_v36, (func() bool {
			if _76_v2 {
				return false
			}
			return _76_v2
		})())
		var _130_v38 _dafny.CodePoint
		_ = _130_v38
		_130_v38 = _dafny.CodePoint('c')
		var _131_v40 D4
		_ = _131_v40
		_131_v40 = Companion_D4_.Create_DC7_(_72_v0)
		var _132_v41 _dafny.Sequence
		_ = _132_v41
		_132_v41 = _dafny.SeqOf(_131_v40, _131_v40, _131_v40, _131_v40, _131_v40)
		var _133_v42 _dafny.Set
		_ = _133_v42
		_133_v42 = _dafny.SetOf(_132_v41)
		var _134_v43 *C1
		_ = _134_v43
		var _nw17 *C1 = New_C1_()
		_ = _nw17
		_nw17.Ctor__(_130_v38, (func() _dafny.Set {
			var _coll6 = _dafny.NewBuilder()
			_ = _coll6
			for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-383))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg7 _dafny.Int) interface{} {
					return coer7(arg7)
				}
			}((func(_135_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_136_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(Companion_D4_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg8 _dafny.Int) interface{} {
							return coer8(arg8)
						}
					}((func(_137_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_138_i5 _dafny.Int) _dafny.CodePoint {
							return _137_v38
						}
					})(_135_v38)))))
				}
			})(_130_v38)))).Elements()); ; {
				_compr_6, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _139_v39 _dafny.Sequence
				_139_v39 = interface{}(_compr_6).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-383))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}((func(_140_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_136_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(Companion_D4_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg10 _dafny.Int) interface{} {
								return coer10(arg10)
							}
						}((func(_141_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_138_i5 _dafny.Int) _dafny.CodePoint {
								return _141_v38
							}
						})(_140_v38)))))
					}
				})(_130_v38))), _139_v39) {
					_coll6.Add(_139_v39)
				}
			}
			return _coll6.ToSet()
		}()).Union(_133_v42), _76_v2, _76_v2)
		_134_v43 = _nw17
		_134_v43 = _134_v43
	} else {
		var _142___mcc_h0 bool = _source2.Get_().(D0_DC1).Cf0
		_ = _142___mcc_h0
		var _143___mcc_h1 _dafny.Sequence = _source2.Get_().(D0_DC1).Cf1
		_ = _143___mcc_h1
		var _144___mcc_h2 bool = _source2.Get_().(D0_DC1).Cf2
		_ = _144___mcc_h2
		var _145_cf2 bool = _144___mcc_h2
		_ = _145_cf2
		var _146_cf1 _dafny.Sequence = _143___mcc_h1
		_ = _146_cf1
		var _147_cf0 bool = _142___mcc_h0
		_ = _147_cf0
		var _148_v44 _dafny.Int
		_ = _148_v44
		_148_v44 = _dafny.IntOfInt64(948)
		_147_cf0 = ((_dafny.IntOfUint32((_72_v0).Cardinality())).Cmp(_148_v44) <= 0) && (_145_cf2)
		var _149_v45 D4
		_ = _149_v45
		_149_v45 = Companion_D4_.Create_DC8_(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_148_v44), _148_v44), _145_cf2)
		var _source4 D4 = _149_v45
		_ = _source4
		if _source4.Is_DC8() {
			var _150___mcc_h6 _dafny.Int = _source4.Get_().(D4_DC8).Cf11
			_ = _150___mcc_h6
			var _151___mcc_h7 bool = _source4.Get_().(D4_DC8).Cf12
			_ = _151___mcc_h7
			var _152_cf12 bool = _151___mcc_h7
			_ = _152_cf12
			var _153_cf11 _dafny.Int = _150___mcc_h6
			_ = _153_cf11
			var _154_v46 _dafny.Map
			_ = _154_v46
			_154_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_cf2, _145_cf2)
			var _155_v47 _dafny.Sequence
			_ = _155_v47
			_155_v47 = _dafny.SeqOf(_154_v46)
			var _156_v48 _dafny.MultiSet
			_ = _156_v48
			_156_v48 = _dafny.MultiSetOf(_147_cf0)
			var _157_v49 D3
			_ = _157_v49
			_157_v49 = Companion_D3_.Create_DC6_(_147_cf0, _153_cf11, (_156_v48).Cardinality())
			var _158_v50 _dafny.Array
			_ = _158_v50
			var _nwElement0_3 _dafny.Map = (_154_v46).Merge(_154_v46)
			_ = _nwElement0_3
			var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(13))
			_ = _nw18
			(_nw18).ArraySet1(_nwElement0_3, 0)
			(_nw18).ArraySet1((_154_v46).Merge(_154_v46), 1)
			(_nw18).ArraySet1(_154_v46, 2)
			(_nw18).ArraySet1((_155_v47).Select((Companion_Default___.SafeIndex(_148_v44, _dafny.IntOfUint32((_155_v47).Cardinality()))).Uint32()).(_dafny.Map), 3)
			(_nw18).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, (_157_v49).Dtor_cf7()), 4)
			(_nw18).ArraySet1((_154_v46).Merge(_154_v46), 5)
			(_nw18).ArraySet1(_154_v46, 6)
			(_nw18).ArraySet1(((_154_v46).Update(_76_v2, _145_cf2)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, !(_147_cf0))), 7)
			(_nw18).ArraySet1(_154_v46, 8)
			(_nw18).ArraySet1((_154_v46).Merge(_154_v46), 9)
			(_nw18).ArraySet1((_154_v46).Merge(_154_v46), 10)
			(_nw18).ArraySet1(Companion_Default___.Fm22(_153_cf11, _dafny.IntOfUint32((_72_v0).Cardinality()), _dafny.IntOfInt64(240), false, _75_globalState), 11)
			(_nw18).ArraySet1((_154_v46).Update(_152_cf12, _152_cf12), 12)
			_158_v50 = _nw18
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_158_v50), 0))
			_ = _index10
			(_158_v50).ArraySet1(_154_v46, (_index10).Int())
			var _159_v51 _dafny.Array
			_ = _159_v51
			var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw19
			_159_v51 = _nw19
			var _160_v52 D10
			_ = _160_v52
			_160_v52 = Companion_D10_.Create_DC24_(_154_v46)
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_158_v50), 0))
			_ = _index11
			var _rhs8 _dafny.Map = (_160_v52).Dtor_cf42()
			_ = _rhs8
			var _rhs9 _dafny.Array = _159_v51
			_ = _rhs9
			var _lhs2 _dafny.Array = _158_v50
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_158_v50), 0))
			_ = _lhs3
			(_lhs2).ArraySet1(_rhs8, (_lhs3).Int())
			_159_v51 = _rhs9
			var _161_v53 _dafny.Sequence
			_ = _161_v53
			_161_v53 = _dafny.SeqOf(_77_v3, _77_v3, _dafny.SeqOf(_147_cf0), _77_v3, _dafny.SeqOf(true))
			var _162_v54 D11
			_ = _162_v54
			_162_v54 = Companion_D11_.Create_DC26_(_161_v53)
			_148_v44 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_148_v44, (_dafny.Zero).Minus(_153_cf11)), _dafny.IntOfUint32(((_162_v54).Dtor_cf44()).Cardinality()))
			var _163_v55 _dafny.Set
			_ = _163_v55
			_163_v55 = _dafny.SetOf(_dafny.IntOfUint32((_72_v0).Cardinality()), _dafny.IntOfUint32((_77_v3).Cardinality()))
			_148_v44 = (_153_cf11).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _163_v55)).Cardinality()))
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_159_v51), 0))
			_ = _index12
			(_159_v51).ArraySet1(_148_v44, (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(66), _dafny.ArrayLen((_159_v51), 0))
			_ = _index13
			(_159_v51).ArraySet1(((func() _dafny.Int {
				if _145_cf2 {
					return _148_v44
				}
				return _148_v44
			})()).Plus(_148_v44), (_index13).Int())
		} else if _source4.Is_DC9() {
			var _164___mcc_h8 bool = _source4.Get_().(D4_DC9).Cf13
			_ = _164___mcc_h8
			var _165___mcc_h9 _dafny.Int = _source4.Get_().(D4_DC9).Cf14
			_ = _165___mcc_h9
			var _166___mcc_h10 _dafny.Array = _source4.Get_().(D4_DC9).Cf15
			_ = _166___mcc_h10
			var _167___mcc_h11 _dafny.Int = _source4.Get_().(D4_DC9).Cf16
			_ = _167___mcc_h11
			var _168_cf16 _dafny.Int = _167___mcc_h11
			_ = _168_cf16
			var _169_cf15 _dafny.Array = _166___mcc_h10
			_ = _169_cf15
			var _170_cf14 _dafny.Int = _165___mcc_h9
			_ = _170_cf14
			var _171_cf13 bool = _164___mcc_h8
			_ = _171_cf13
			_171_cf13 = _171_cf13
			var _172_v56 _dafny.Set
			_ = _172_v56
			_172_v56 = _dafny.SetOf(_147_cf0, _76_v2)
			var _173_v57 _dafny.MultiSet
			_ = _173_v57
			_173_v57 = _dafny.MultiSetOf((_172_v56).Cardinality())
			var _174_v58 _dafny.MultiSet
			_ = _174_v58
			_174_v58 = _dafny.MultiSetOf(_173_v57)
			var _175_v59 _dafny.Set
			_ = _175_v59
			_175_v59 = _dafny.SetOf((_174_v58).Cardinality())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_73_v1), 0))
			_ = _index14
			(_73_v1).ArraySet1((func() bool {
				if _171_cf13 {
					return _76_v2
				}
				return false
			})(), (_index14).Int())
			var _176_v60 D2
			_ = _176_v60
			_176_v60 = Companion_D2_.Create_DC3_(_170_cf14)
			var _177_v61 *C2
			_ = _177_v61
			var _nw20 *C2 = New_C2_()
			_ = _nw20
			_nw20.Ctor__((_168_cf16).Cmp((_176_v60).Dtor_cf4()) == 0, _76_v2)
			_177_v61 = _nw20
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_73_v1), 0))
			_ = _index15
			var _rhs10 _dafny.Set = Companion_Default___.Fm23(_171_cf13, _dafny.SetOf(_72_v0), _168_cf16, _75_globalState)
			_ = _rhs10
			var _rhs11 bool = (_dafny.IntOfUint32((_72_v0).Cardinality())).Cmp(_148_v44) >= 0
			_ = _rhs11
			var _rhs12 bool = _171_cf13
			_ = _rhs12
			var _rhs13 _dafny.Int = _170_cf14
			_ = _rhs13
			var _rhs14 *C2 = _177_v61
			_ = _rhs14
			var _lhs4 _dafny.Array = _73_v1
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_73_v1), 0))
			_ = _lhs5
			_175_v59 = _rhs10
			_76_v2 = _rhs11
			(_lhs4).ArraySet1(_rhs12, (_lhs5).Int())
			_170_cf14 = _rhs13
			_177_v61 = _rhs14
			var _178_v62 *C2
			_ = _178_v62
			var _nw21 *C2 = New_C2_()
			_ = _nw21
			_nw21.Ctor__(!((_173_v57).IsDisjointFrom(_dafny.MultiSetOf(_dafny.IntOfUint32((Companion_Default___.Fm14(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}(func(_179_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()), _72_v0, _76_v2, _75_globalState)).Cardinality())))), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
			_178_v62 = _nw21
			var _180_v63 *C2
			_ = _180_v63
			var _nw22 *C2 = New_C2_()
			_ = _nw22
			_nw22.Ctor__((func() bool {
				if _76_v2 {
					return _147_cf0
				}
				return (_178_v62).Fm0(_75_globalState)
			})(), _76_v2)
			_180_v63 = _nw22
		} else {
			var _181___mcc_h12 _dafny.Sequence = _source4.Get_().(D4_DC7).Cf10
			_ = _181___mcc_h12
			var _182_cf10 _dafny.Sequence = _181___mcc_h12
			_ = _182_cf10
			_148_v44 = ((_148_v44).Minus(_dafny.IntOfInt64(-862))).Minus(_148_v44)
			_145_cf2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(338))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_183_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), _dafny.UnicodeSeqOfUtf8Bytes("hduwoi"))
			var _184_v64 _dafny.Map
			_ = _184_v64
			_184_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v44, (_dafny.IntOfUint32((_146_cf1).Cardinality())).Plus(_dafny.IntOfInt64(859)))
			_148_v44 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_184_v64).Contains(_dafny.IntOfUint32((_182_cf10).Cardinality())) {
					return (_184_v64).Get(_dafny.IntOfUint32((_182_cf10).Cardinality())).(_dafny.Int)
				}
				return _148_v44
			})())
			var _185_v65 _dafny.Array
			_ = _185_v65
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw23
			_185_v65 = _nw23
			var _186_v66 *C3
			_ = _186_v66
			var _nw24 *C3 = New_C3_()
			_ = _nw24
			_nw24.Ctor__(_76_v2, _145_cf2)
			_186_v66 = _nw24
			var _187_v67 D12
			_ = _187_v67
			_187_v67 = Companion_D12_.Create_DC29_(_186_v66)
			var _188_v68 _dafny.Array
			_ = _188_v68
			var _nwElement0_4 *C3 = _186_v66
			_ = _nwElement0_4
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_4, 0)
			(_nw25).ArraySet1(_186_v66, 1)
			(_nw25).ArraySet1(_186_v66, 2)
			(_nw25).ArraySet1(_186_v66, 3)
			(_nw25).ArraySet1((_187_v67).Dtor_cf49(), 4)
			(_nw25).ArraySet1(_186_v66, 5)
			(_nw25).ArraySet1(_186_v66, 6)
			(_nw25).ArraySet1(_186_v66, 7)
			(_nw25).ArraySet1(_186_v66, 8)
			(_nw25).ArraySet1(_186_v66, 9)
			(_nw25).ArraySet1(_186_v66, 10)
			(_nw25).ArraySet1(_186_v66, 11)
			(_nw25).ArraySet1(_186_v66, 12)
			(_nw25).ArraySet1(_186_v66, 13)
			(_nw25).ArraySet1(_186_v66, 14)
			(_nw25).ArraySet1(_186_v66, 15)
			(_nw25).ArraySet1(_186_v66, 16)
			(_nw25).ArraySet1(_186_v66, 17)
			(_nw25).ArraySet1(_186_v66, 18)
			(_nw25).ArraySet1(_186_v66, 19)
			(_nw25).ArraySet1(_186_v66, 20)
			(_nw25).ArraySet1(_186_v66, 21)
			(_nw25).ArraySet1(_186_v66, 22)
			(_nw25).ArraySet1((Companion_D12_.Create_DC29_(_186_v66)).Dtor_cf49(), 23)
			(_nw25).ArraySet1(_186_v66, 24)
			(_nw25).ArraySet1(_186_v66, 25)
			(_nw25).ArraySet1(_186_v66, 26)
			(_nw25).ArraySet1(_186_v66, 27)
			_188_v68 = _nw25
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_185_v65), 0))
			_ = _index16
			(_185_v65).ArraySet1(_188_v68, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(675), _dafny.ArrayLen((_185_v65), 0))
			_ = _index17
			(_185_v65).ArraySet1(_188_v68, (_index17).Int())
		}
		_148_v44 = ((_148_v44).Minus(_148_v44)).Times(Companion_Default___.SafeModuloInt(_148_v44, _148_v44))
		if (func() bool {
			if _76_v2 {
				return _145_cf2
			}
			return _145_cf2
		})() {
			var _189_v69 _dafny.Array
			_ = _189_v69
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
			_ = _nw26
			_189_v69 = _nw26
			var _190_v70 *C0
			_ = _190_v70
			var _nw27 *C0 = New_C0_()
			_ = _nw27
			_nw27.Ctor__(_189_v69)
			_190_v70 = _nw27
			var _191_v71 *C0
			_ = _191_v71
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__(_189_v69)
			_191_v71 = _nw28
			var _192_v72 _dafny.Array
			_ = _192_v72
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw29
			_192_v72 = _nw29
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_192_v72), 0))
			_ = _index18
			(_192_v72).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_77_v3).Cardinality())), (_index18).Int())
			var _193_v73 _dafny.Set
			_ = _193_v73
			_193_v73 = _dafny.SetOf(_dafny.IntOfInt64(426))
			var _194_v74 _dafny.Sequence
			_ = _194_v74
			_194_v74 = _dafny.SeqOf(_148_v44, _148_v44)
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_192_v72), 0))
			_ = _index19
			var _rhs15 bool = (_193_v73).IsSubsetOf(_193_v73)
			_ = _rhs15
			var _rhs16 _dafny.Int = Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(_148_v44)).Plus(_148_v44), ((_dafny.Zero).Minus(_dafny.IntOfUint32((_194_v74).Cardinality()))).Minus(_148_v44))
			_ = _rhs16
			var _lhs6 _dafny.Array = _192_v72
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen((_192_v72), 0))
			_ = _lhs7
			_145_cf2 = _rhs15
			(_lhs6).ArraySet1(_rhs16, (_lhs7).Int())
			var _195_v75 _dafny.Int
			_ = _195_v75
			var _196_v76 bool
			_ = _196_v76
			var _out3 _dafny.Int
			_ = _out3
			var _out4 bool
			_ = _out4
			_out3, _out4 = Companion_Default___.M4(_148_v44, (func() bool {
				if _145_cf2 {
					return _147_cf0
				}
				return _76_v2
			})(), _72_v0, _75_globalState)
			_195_v75 = _out3
			_196_v76 = _out4
			var _197_v77 _dafny.Set
			_ = _197_v77
			_197_v77 = _dafny.SetOf(_76_v2)
			_147_cf0 = ((_197_v77).Intersection(_197_v77)).IsSubsetOf(_197_v77)
		} else {
			var _198_v78 *C3
			_ = _198_v78
			var _nw30 *C3 = New_C3_()
			_ = _nw30
			_nw30.Ctor__(_147_cf0, _145_cf2)
			_198_v78 = _nw30
			var _199_v79 D12
			_ = _199_v79
			_199_v79 = Companion_D12_.Create_DC29_(_198_v78)
			var _pat_let_tv6 = _198_v78
			_ = _pat_let_tv6
			var _200_v80 _dafny.Array
			_ = _200_v80
			var _nwElement0_5 D12 = _199_v79
			_ = _nwElement0_5
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(19))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_5, 0)
			(_nw31).ArraySet1(_199_v79, 1)
			(_nw31).ArraySet1(_199_v79, 2)
			(_nw31).ArraySet1((func() D12 {
				if _76_v2 {
					return Companion_D12_.Create_DC29_(_198_v78)
				}
				return _199_v79
			})(), 3)
			(_nw31).ArraySet1((func() D12 {
				if _145_cf2 {
					return _199_v79
				}
				return Companion_D12_.Create_DC29_(_198_v78)
			})(), 4)
			(_nw31).ArraySet1(_199_v79, 5)
			(_nw31).ArraySet1(_199_v79, 6)
			(_nw31).ArraySet1(_199_v79, 7)
			(_nw31).ArraySet1(_199_v79, 8)
			(_nw31).ArraySet1(_199_v79, 9)
			(_nw31).ArraySet1(func(_pat_let8_0 D12) D12 {
				return func(_201_dt__update__tmp_h3 D12) D12 {
					return func(_pat_let9_0 *C3) D12 {
						return func(_202_dt__update_hcf49_h0 *C3) D12 {
							return Companion_D12_.Create_DC29_(_202_dt__update_hcf49_h0)
						}(_pat_let9_0)
					}(_pat_let_tv6)
				}(_pat_let8_0)
			}(_199_v79), 10)
			(_nw31).ArraySet1(_199_v79, 11)
			(_nw31).ArraySet1(_199_v79, 12)
			(_nw31).ArraySet1(_199_v79, 13)
			(_nw31).ArraySet1(_199_v79, 14)
			(_nw31).ArraySet1(_199_v79, 15)
			(_nw31).ArraySet1(_199_v79, 16)
			(_nw31).ArraySet1(_199_v79, 17)
			(_nw31).ArraySet1(_199_v79, 18)
			_200_v80 = _nw31
			_200_v80 = _200_v80
			var _203_v81 _dafny.Map
			_ = _203_v81
			_203_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_148_v44, Companion_Default___.Fm24(_75_globalState))
			var _204_v82 _dafny.Set
			_ = _204_v82
			_204_v82 = _dafny.SetOf(_145_cf2)
			var _205_v83 D7
			_ = _205_v83
			_205_v83 = Companion_D7_.Create_DC14_(_204_v82)
			_203_v81 = (_203_v81).Update((_dafny.Zero).Minus(_148_v44), Companion_D7_.Create_DC16_(_205_v83))
			var _206_v84 _dafny.Int
			_ = _206_v84
			var _207_v85 bool
			_ = _207_v85
			var _out5 _dafny.Int
			_ = _out5
			var _out6 bool
			_ = _out6
			_out5, _out6 = Companion_Default___.M4(Companion_Default___.SafeDivisionInt(_148_v44, _148_v44), false, _72_v0, _75_globalState)
			_206_v84 = _out5
			_207_v85 = _out6
			var _208_v86 _dafny.Array
			_ = _208_v86
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(2))
			_ = _nw32
			_208_v86 = _nw32
			_208_v86 = _208_v86
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_73_v1), 0))
			_ = _index20
			(_73_v1).ArraySet1(_207_v85, (_index20).Int())
			var _209_v87 _dafny.Set
			_ = _209_v87
			_209_v87 = _dafny.SetOf(_206_v84, _148_v44)
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_73_v1), 0))
			_ = _index21
			(_73_v1).ArraySet1((_209_v87).IsProperSubsetOf(func() _dafny.Set {
				var _coll7 = _dafny.NewBuilder()
				_ = _coll7
				for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-626), _dafny.IntOfInt64(343))); ; {
					_compr_7, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _210_v88 _dafny.Int
					_210_v88 = interface{}(_compr_7).(_dafny.Int)
					if ((_dafny.IntOfInt64(-626)).Cmp(_210_v88) <= 0) && ((_210_v88).Cmp(_dafny.IntOfInt64(343)) < 0) {
						_coll7.Add(Companion_Default___.SafeModuloInt(_210_v88, _148_v44))
					}
				}
				return _coll7.ToSet()
			}()), (_index21).Int())
		}
	}
	var _211_v89 _dafny.Int
	_ = _211_v89
	_211_v89 = _dafny.IntOfInt64(-166)
	var _212_v90 _dafny.Int
	_ = _212_v90
	var _213_v91 bool
	_ = _213_v91
	var _out7 _dafny.Int
	_ = _out7
	var _out8 bool
	_ = _out8
	_out7, _out8 = Companion_Default___.M4(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(827), _211_v89), _76_v2, Companion_Default___.Fm4(_211_v89, _dafny.IntOfInt64(610), _75_globalState), _75_globalState)
	_212_v90 = _out7
	_213_v91 = _out8
	var _214_v92 _dafny.Array
	_ = _214_v92
	var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
	_ = _nw33
	_214_v92 = _nw33
	var _215_v93 *C0
	_ = _215_v93
	var _nw34 *C0 = New_C0_()
	_ = _nw34
	_nw34.Ctor__(_214_v92)
	_215_v93 = _nw34
	var _216_v94 _dafny.Map
	_ = _216_v94
	_216_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _211_v89)
	var _217_v95 _dafny.MultiSet
	_ = _217_v95
	_217_v95 = _dafny.MultiSetOf((_216_v94).Cardinality())
	var _218_v96 _dafny.Map
	_ = _218_v96
	_218_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, (_217_v95).Cardinality())
	var _219_v97 _dafny.Map
	_ = _219_v97
	_219_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_216_v94).Merge(_216_v94)).Cardinality(), (_211_v89).Times(_212_v90))
	var _220_v98 *C3
	_ = _220_v98
	var _nw35 *C3 = New_C3_()
	_ = _nw35
	_nw35.Ctor__(_76_v2, _76_v2)
	_220_v98 = _nw35
	var _221_v99 D12
	_ = _221_v99
	_221_v99 = Companion_D12_.Create_DC29_(_220_v98)
	var _222_v100 D12
	_ = _222_v100
	_222_v100 = Companion_D12_.Create_DC32_(_221_v99)
	var _223_v101 _dafny.Array
	_ = _223_v101
	var _nwElement0_6 D12 = _222_v100
	_ = _nwElement0_6
	var _nw36 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(8))
	_ = _nw36
	(_nw36).ArraySet1(_nwElement0_6, 0)
	(_nw36).ArraySet1(_222_v100, 1)
	(_nw36).ArraySet1(_222_v100, 2)
	(_nw36).ArraySet1(_222_v100, 3)
	(_nw36).ArraySet1(_222_v100, 4)
	(_nw36).ArraySet1(Companion_D12_.Create_DC32_(Companion_D12_.Create_DC29_(_220_v98)), 5)
	(_nw36).ArraySet1(_222_v100, 6)
	(_nw36).ArraySet1(Companion_D12_.Create_DC32_(_221_v99), 7)
	_223_v101 = _nw36
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_223_v101), 0))
	_ = _index22
	(_223_v101).ArraySet1(_222_v100, (_index22).Int())
	var _224_v102 _dafny.Array
	_ = _224_v102
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(9)
	_ = _len0_4
	var _nw37 _dafny.Array
	_ = _nw37
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw37 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) _dafny.Int = (func(_225_v90 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_226_i8 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_226_i8, _225_v90)
			}
		})(_212_v90)
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
	_224_v102 = _nw37
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_223_v101), 0))
	_ = _index23
	var _rhs17 _dafny.Map = _219_v97
	_ = _rhs17
	var _rhs18 D12 = _222_v100
	_ = _rhs18
	var _rhs19 _dafny.Array = _224_v102
	_ = _rhs19
	var _rhs20 bool = !(((func() _dafny.Int {
		if _213_v91 {
			return _212_v90
		}
		return _211_v89
	})()).Cmp(_212_v90) <= 0)
	_ = _rhs20
	var _rhs21 _dafny.Int = (_dafny.Zero).Minus(_212_v90)
	_ = _rhs21
	var _lhs8 _dafny.Array = _223_v101
	_ = _lhs8
	var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_223_v101), 0))
	_ = _lhs9
	_219_v97 = _rhs17
	(_lhs8).ArraySet1(_rhs18, (_lhs9).Int())
	_224_v102 = _rhs19
	_213_v91 = _rhs20
	_212_v90 = _rhs21
	var _227_v103 _dafny.Sequence
	_ = _227_v103
	_227_v103 = _dafny.SeqOf(_212_v90)
	_227_v103 = _227_v103
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
	_ = _index24
	(_73_v1).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_72_v0, _dafny.UnicodeSeqOfUtf8Bytes("rggpcxlop")), (_index24).Int())
	var _228_v104 _dafny.CodePoint
	_ = _228_v104
	_228_v104 = _dafny.CodePoint('q')
	var _229_v105 _dafny.Array
	_ = _229_v105
	var _nwElement0_7 _dafny.Sequence = Companion_Default___.Fm4(_212_v90, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-737))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}((func(_230_v91 bool, _231_v2 bool) func(_dafny.Int) _dafny.Sequence {
		return func(_232_i9 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(_230_v91, _231_v2, _230_v91, _230_v91)
		}
	})(_213_v91, _76_v2)))).Cardinality()), _75_globalState)
	_ = _nwElement0_7
	var _nw38 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(8))
	_ = _nw38
	(_nw38).ArraySet1(_nwElement0_7, 0)
	(_nw38).ArraySet1(_72_v0, 1)
	(_nw38).ArraySet1((func() _dafny.Sequence {
		if _213_v91 {
			return _72_v0
		}
		return _72_v0
	})(), 2)
	(_nw38).ArraySet1(_72_v0, 3)
	(_nw38).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hw"), 4)
	(_nw38).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
		if _213_v91 {
			return _72_v0
		}
		return _72_v0
	})(), (Companion_Default___.SafeIndex(_211_v89, _dafny.IntOfUint32(((func() _dafny.Sequence {
		if _213_v91 {
			return _72_v0
		}
		return _72_v0
	})()).Cardinality()))).Uint32(), _228_v104), 5)
	(_nw38).ArraySet1(_72_v0, 6)
	(_nw38).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(589))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_233_i10 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('j')
	})), 7)
	_229_v105 = _nw38
	var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
	_ = _index25
	(_229_v105).ArraySet1(_72_v0, (_index25).Int())
	var _234_v106 _dafny.Map
	_ = _234_v106
	_234_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v1, _76_v2)
	var _235_v107 _dafny.Map
	_ = _235_v107
	_235_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v91, _76_v2)
	var _pat_let_tv7 = _211_v89
	_ = _pat_let_tv7
	var _pat_let_tv8 = _235_v107
	_ = _pat_let_tv8
	var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
	_ = _index26
	var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
	_ = _index27
	var _rhs22 bool = ((_dafny.Zero).Minus((func() _dafny.Int {
		if _213_v91 {
			return (_dafny.Zero).Minus(_211_v89)
		}
		return _211_v89
	})())).Cmp((func() _dafny.Int {
		if (_219_v97).Contains(_211_v89) {
			return (_219_v97).Get(_211_v89).(_dafny.Int)
		}
		return (_dafny.Zero).Minus((_234_v106).Cardinality())
	})()) >= 0
	_ = _rhs22
	var _rhs23 _dafny.Int = (func() _dafny.Int {
		if _76_v2 {
			return _211_v89
		}
		return ((Companion_Default___.Fm25(_212_v90, _211_v89, _dafny.IntOfInt64(781), _75_globalState)).Dtor_cf51()).Plus(_212_v90)
	})()
	_ = _rhs23
	var _rhs24 bool = func(_source5 D10) bool {
		if _source5.Is_DC25() {
			var _236___mcc_h13 _dafny.Sequence = _source5.Get_().(D10_DC25).Cf43
			_ = _236___mcc_h13
			var _237_cf43 _dafny.Sequence = _236___mcc_h13
			_ = _237_cf43
			return false
		} else {
			var _238___mcc_h14 _dafny.Map = _source5.Get_().(D10_DC24).Cf42
			_ = _238___mcc_h14
			var _239_cf42 _dafny.Map = _238___mcc_h14
			_ = _239_cf42
			return (_pat_let_tv7).Cmp(_dafny.IntOfInt64(833)) == 0
		}
	}(func(_pat_let10_0 D10) D10 {
		return func(_240_dt__update__tmp_h4 D10) D10 {
			return func(_pat_let11_0 _dafny.Map) D10 {
				return func(_241_dt__update_hcf42_h0 _dafny.Map) D10 {
					return Companion_D10_.Create_DC24_(_241_dt__update_hcf42_h0)
				}(_pat_let11_0)
			}(_pat_let_tv8)
		}(_pat_let10_0)
	}(Companion_D10_.Create_DC24_(_235_v107)))
	_ = _rhs24
	var _rhs25 _dafny.Sequence = _72_v0
	_ = _rhs25
	var _lhs10 _dafny.Array = _73_v1
	_ = _lhs10
	var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
	_ = _lhs11
	var _lhs12 _dafny.Array = _229_v105
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
	_ = _lhs13
	_213_v91 = _rhs22
	_212_v90 = _rhs23
	(_lhs10).ArraySet1(_rhs24, (_lhs11).Int())
	(_lhs12).ArraySet1(_rhs25, (_lhs13).Int())
	var _242_v108 D4
	_ = _242_v108
	_242_v108 = Companion_D4_.Create_DC7_((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence))
	var _243_v109 _dafny.Set
	_ = _243_v109
	_243_v109 = _dafny.SetOf(_dafny.SeqOf(_242_v108, _242_v108, _242_v108, _242_v108))
	var _244_v110 _dafny.Sequence
	_ = _244_v110
	_244_v110 = _dafny.SeqOf(_220_v98)
	var _245_v111 *C1
	_ = _245_v111
	var _nw39 *C1 = New_C1_()
	_ = _nw39
	_nw39.Ctor__(_228_v104, (_243_v109).Union(_243_v109), _dafny.Companion_Sequence_.Equal(_244_v110, _244_v110), _213_v91)
	_245_v111 = _nw39
	var _246_i11 _dafny.Int
	_ = _246_i11
	_246_i11 = _dafny.Zero
	{
		for _213_v91 {
			{
				if (_246_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_246_i11 = (_246_i11).Plus(_dafny.One)
				var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
				_ = _index28
				(_229_v105).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _72_v0), (_index28).Int())
				if (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool) {
					var _247_v112 _dafny.MultiSet
					_ = _247_v112
					_247_v112 = _dafny.MultiSetOf(!(!(_76_v2)), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
					var _248_v113 D0
					_ = _248_v113
					_248_v113 = Companion_D0_.Create_DC0_()
					(_220_v98).M0(_247_v112, _248_v113, _76_v2, _75_globalState)
					_76_v2 = _76_v2
					var _249_v114 _dafny.Map
					_ = _249_v114
					_249_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(637), (_245_v111).F10())
					var _250_v115 _dafny.Map
					_ = _250_v115
					_250_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, _249_v114)
					_213_v91 = (_250_v115).Contains(_211_v89)
					var _251_v116 *C0
					_ = _251_v116
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__((_215_v93).F12())
					_251_v116 = _nw40
					var _252_v117 _dafny.Set
					_ = _252_v117
					_252_v117 = _dafny.SetOf((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _213_v91)
					var _253_v118 _dafny.Sequence
					_ = _253_v118
					_253_v118 = _dafny.SeqOf(_252_v117)
					var _254_v119 D7
					_ = _254_v119
					_254_v119 = Companion_D7_.Create_DC14_((_253_v118).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(214))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg15 _dafny.Int) interface{} {
							return coer15(arg15)
						}
					}((func(_255_v89 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_256_i12 _dafny.Int) _dafny.Int {
							return _255_v89
						}
					})(_211_v89)))).Cardinality()), _dafny.IntOfUint32((_253_v118).Cardinality()))).Uint32()).(_dafny.Set))
					var _pat_let_tv9 = _252_v117
					_ = _pat_let_tv9
					_254_v119 = func(_pat_let12_0 D7) D7 {
						return func(_257_dt__update__tmp_h5 D7) D7 {
							return func(_pat_let13_0 _dafny.Set) D7 {
								return func(_258_dt__update_hcf25_h0 _dafny.Set) D7 {
									return Companion_D7_.Create_DC14_(_258_dt__update_hcf25_h0)
								}(_pat_let13_0)
							}(_pat_let_tv9)
						}(_pat_let12_0)
					}(_254_v119)
				} else {
					var _259_v120 D7
					_ = _259_v120
					_259_v120 = Companion_D7_.Create_DC14_(_dafny.SetOf((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)))
					var _260_v121 D7
					_ = _260_v121
					_260_v121 = Companion_D7_.Create_DC16_(_259_v120)
					_260_v121 = Companion_D7_.Create_DC16_(_259_v120)
					var _261_v122 _dafny.Array
					_ = _261_v122
					var _nw41 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(5))
					_ = _nw41
					_261_v122 = _nw41
					var _262_v123 D3
					_ = _262_v123
					_262_v123 = Companion_D3_.Create_DC6_(_76_v2, _dafny.IntOfInt64(839), _212_v90)
					var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_261_v122), 0))
					_ = _index29
					(_261_v122).ArraySet1(_262_v123, (_index29).Int())
					var _263_v124 _dafny.Set
					_ = _263_v124
					_263_v124 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(835))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}(func(_264_i13 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('x')
					}))).Cardinality()), _211_v89)
					var _265_v125 D0
					_ = _265_v125
					_265_v125 = Companion_D0_.Create_DC1_(!(!(_263_v124).Equals(_263_v124)), _77_v3, (_213_v91) || (true))
					var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_261_v122), 0))
					_ = _index30
					var _rhs26 D3 = _262_v123
					_ = _rhs26
					var _rhs27 D0 = _265_v125
					_ = _rhs27
					var _rhs28 _dafny.CodePoint = _228_v104
					_ = _rhs28
					var _lhs14 _dafny.Array = _261_v122
					_ = _lhs14
					var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(224), _dafny.ArrayLen((_261_v122), 0))
					_ = _lhs15
					(_lhs14).ArraySet1(_rhs26, (_lhs15).Int())
					_265_v125 = _rhs27
					_228_v104 = _rhs28
					var _266_v126 _dafny.Set
					_ = _266_v126
					_266_v126 = _dafny.SetOf(_76_v2, (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
					var _267_v127 D2
					_ = _267_v127
					_267_v127 = Companion_D2_.Create_DC3_((func() _dafny.Int {
						if (_217_v95).Contains((_266_v126).Cardinality()) {
							return (_217_v95).Multiplicity((_266_v126).Cardinality())
						}
						return _212_v90
					})())
					_211_v89 = (_267_v127).Dtor_cf4()
					_211_v89 = _212_v90
					_211_v89 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm9((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _75_globalState), Companion_Default___.Fm4(_212_v90, _dafny.IntOfInt64(921), _75_globalState))).Cardinality())
				}
				var _268_v128 _dafny.Map
				_ = _268_v128
				_268_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(247))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}((func(_269_v90 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_270_i14 _dafny.Int) _dafny.Int {
						return _269_v90
					}
				})(_212_v90)))).Cardinality()), _217_v95)
				var _271_v129 D11
				_ = _271_v129
				_271_v129 = Companion_D11_.Create_DC27_(true, _268_v128)
				_213_v91 = !(!((_271_v129).Dtor_cf45()))
				var _272_v130 _dafny.MultiSet
				_ = _272_v130
				_272_v130 = _dafny.MultiSetOf(_220_v98, _220_v98, _220_v98, _220_v98)
				var _273_v131 D12
				_ = _273_v131
				_273_v131 = Companion_D12_.Create_DC30_(_dafny.IntOfInt64(885), (_dafny.Zero).Minus((_272_v130).Cardinality()))
				_216_v94 = (_216_v94).Update((_213_v91) == ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)), ((_273_v131).Dtor_cf51()).Plus(_211_v89))
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _hi2 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_72_v0).Cardinality()))).Plus(_dafny.IntOfInt64(203))
	_ = _hi2
	for _274_i15 := _dafny.IntOfInt64(929); _274_i15.Cmp(_hi2) < 0; _274_i15 = _274_i15.Plus(_dafny.One) {
		var _275_v133 _dafny.Set
		_ = _275_v133
		_275_v133 = _dafny.SetOf(_dafny.IntOfUint32((_72_v0).Cardinality()), (func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(868), _dafny.IntOfInt64(615))); ; {
				_compr_8, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _276_v132 _dafny.Int
				_276_v132 = interface{}(_compr_8).(_dafny.Int)
				if ((_dafny.IntOfInt64(868)).Cmp(_276_v132) <= 0) && ((_276_v132).Cmp(_dafny.IntOfInt64(615)) < 0) {
					_coll8.Add(Companion_Default___.SafeDivisionInt(_276_v132, _211_v89), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_77_v3).Cardinality()))))).Cardinality())
				}
			}
			return _coll8.ToMap()
		}()).Cardinality())
		if (_dafny.SetOf(_212_v90)).IsProperSubsetOf(_275_v133) {
			var _277_v134 T1
			_ = _277_v134
			var _nw42 *C2 = New_C2_()
			_ = _nw42
			_nw42.Ctor__((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _76_v2)
			_277_v134 = _nw42
			var _278_v135 _dafny.Array
			_ = _278_v135
			var _nwElement0_8 T1 = _277_v134
			_ = _nwElement0_8
			var _nw43 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
			_ = _nw43
			(_nw43).ArraySet1(_nwElement0_8, 0)
			(_nw43).ArraySet1(_277_v134, 1)
			(_nw43).ArraySet1(_277_v134, 2)
			(_nw43).ArraySet1(_277_v134, 3)
			_278_v135 = _nw43
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_278_v135), 0))
			_ = _index31
			(_278_v135).ArraySet1(_277_v134, (_index31).Int())
			var _279_v136 _dafny.Map
			_ = _279_v136
			_279_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_76_v2, _277_v134)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_278_v135), 0))
			_ = _index32
			(_278_v135).ArraySet1((func() T1 {
				if (_279_v136).Contains(false) {
					return (_279_v136).Get(false).(T1)
				}
				return _277_v134
			})(), (_index32).Int())
			_228_v104 = (_245_v111).F10()
			_212_v90 = (_211_v89).Minus(Companion_Default___.SafeDivisionInt(_274_i15, _274_i15))
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index33
			(_73_v1).ArraySet1((_215_v93).Fm0(_75_globalState), (_index33).Int())
			var _280_v137 _dafny.Int
			_ = _280_v137
			var _281_v138 bool
			_ = _281_v138
			var _out9 _dafny.Int
			_ = _out9
			var _out10 bool
			_ = _out10
			_out9, _out10 = (_220_v98).M2(_277_v134.F8(), _75_globalState)
			_280_v137 = _out9
			_281_v138 = _out10
		} else {
			_216_v94 = (_216_v94).Update((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _274_i15)
			_228_v104 = _dafny.CodePoint('g')
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
			_ = _index34
			(_229_v105).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cubgc"), (_index34).Int())
			var _282_v139 _dafny.Array
			_ = _282_v139
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_5
			var _nw44 _dafny.Array
			_ = _nw44
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw44 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_283_v2 bool, _284_v3 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
					return func(_285_i16 _dafny.Int) _dafny.MultiSet {
						return (_dafny.MultiSetOf(!(_283_v2))).Intersection(_dafny.MultiSetFromSeq(_284_v3))
					}
				})(_76_v2, _77_v3)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw44 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw44).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw44).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_282_v139 = _nw44
			var _286_v140 _dafny.Set
			_ = _286_v140
			_286_v140 = _dafny.SetOf(_213_v91)
			var _287_v141 _dafny.Map
			_ = _287_v141
			_287_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_219_v97).Cardinality()), _286_v140)
			var _rhs29 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ovh")
			_ = _rhs29
			var _rhs30 _dafny.Int = ((_287_v141).Update(_275_v133, (_286_v140).Union(_286_v140))).Cardinality()
			_ = _rhs30
			var _rhs31 _dafny.Array = _282_v139
			_ = _rhs31
			_72_v0 = _rhs29
			_212_v90 = _rhs30
			_282_v139 = _rhs31
			var _288_v142 *C0
			_ = _288_v142
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__((_215_v93).F12())
			_288_v142 = _nw45
		}
		var _289_v143 D0
		_ = _289_v143
		_289_v143 = Companion_D0_.Create_DC1_(_76_v2, _77_v3, false)
		var _290_v144 _dafny.Int
		_ = _290_v144
		var _291_v145 bool
		_ = _291_v145
		var _out11 _dafny.Int
		_ = _out11
		var _out12 bool
		_ = _out12
		_out11, _out12 = (_220_v98).M2((_220_v98).Fm1((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _289_v143, _75_globalState), _75_globalState)
		_290_v144 = _out11
		_291_v145 = _out12
		_291_v145 = _76_v2
		_291_v145 = (_290_v144).Cmp(_212_v90) >= 0
	}
	_76_v2 = (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
	var _source6 D11 = Companion_D11_.Create_DC27_(_76_v2, func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(421), _dafny.IntOfInt64(465))); ; {
			_compr_9, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _292_v146 _dafny.Int
			_292_v146 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(421)).Cmp(_292_v146) <= 0) && ((_292_v146).Cmp(_dafny.IntOfInt64(465)) < 0) {
				_coll9.Add((_292_v146).Minus(_212_v90), _217_v95)
			}
		}
		return _coll9.ToMap()
	}())
	_ = _source6
	if _source6.Is_DC27() {
		var _293___mcc_h15 bool = _source6.Get_().(D11_DC27).Cf45
		_ = _293___mcc_h15
		var _294___mcc_h16 _dafny.Map = _source6.Get_().(D11_DC27).Cf46
		_ = _294___mcc_h16
		var _295_cf46 _dafny.Map = _294___mcc_h16
		_ = _295_cf46
		var _296_cf45 bool = _293___mcc_h15
		_ = _296_cf45
		if _76_v2 {
			_212_v90 = _211_v89
			_242_v108 = _242_v108
			_220_v98 = _220_v98
			_211_v89 = _212_v90
			var _297_v147 *C2
			_ = _297_v147
			var _nw46 *C2 = New_C2_()
			_ = _nw46
			_nw46.Ctor__(_213_v91, _76_v2)
			_297_v147 = _nw46
			var _298_v148 _dafny.MultiSet
			_ = _298_v148
			_298_v148 = _dafny.MultiSetOf((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence))
			var _299_v149 D0
			_ = _299_v149
			_299_v149 = Companion_D0_.Create_DC1_(_76_v2, _77_v3, (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
			var _300_v150 _dafny.Map
			_ = _300_v150
			_300_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_245_v111).F10(), _211_v89)
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index35
			var _rhs32 bool = (_298_v148).IsProperSubsetOf(_298_v148)
			_ = _rhs32
			var _rhs33 *C2 = _297_v147
			_ = _rhs33
			var _rhs34 _dafny.Int = (_245_v111).Fm2(_299_v149, _75_globalState)
			_ = _rhs34
			var _rhs35 bool = (_77_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_300_v150).Contains((_245_v111).F10()) {
					return (_300_v150).Get((_245_v111).F10()).(_dafny.Int)
				}
				return _212_v90
			})(), _dafny.IntOfUint32((_77_v3).Cardinality()))).Uint32()).(bool)
			_ = _rhs35
			var _rhs36 _dafny.Int = (_297_v147).Fm2(_299_v149, _75_globalState)
			_ = _rhs36
			var _lhs16 _dafny.Array = _73_v1
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _lhs17
			(_lhs16).ArraySet1(_rhs32, (_lhs17).Int())
			_297_v147 = _rhs33
			_211_v89 = _rhs34
			_213_v91 = _rhs35
			_211_v89 = _rhs36
		} else {
			var _301_v151 _dafny.Sequence
			_ = _301_v151
			_301_v151 = _dafny.SeqOf(_224_v102, _224_v102, _224_v102, _224_v102, _224_v102)
			_296_cf45 = _dafny.Companion_Sequence_.IsProperPrefixOf(_301_v151, _301_v151)
			_76_v2 = !(_213_v91)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_224_v102), 0))
			_ = _index36
			(_224_v102).ArraySet1(((_dafny.MultiSetFromSeq(_227_v103)).Union(_217_v95)).Cardinality(), (_index36).Int())
			var _302_v152 _dafny.Sequence
			_ = _302_v152
			_302_v152 = _dafny.SeqOf(_242_v108)
			var _303_v153 T2
			_ = _303_v153
			var _nw47 *C1 = New_C1_()
			_ = _nw47
			_nw47.Ctor__(_228_v104, _dafny.SetOf(_302_v152, _302_v152, _302_v152), !_dafny.Companion_Sequence_.Contains(_77_v3, true), _76_v2)
			_303_v153 = _nw47
			var _304_v154 _dafny.Sequence
			_ = _304_v154
			_304_v154 = _dafny.SeqOf(_303_v153)
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_224_v102), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index38
			var _rhs37 _dafny.Int = _dafny.IntOfInt64(296)
			_ = _rhs37
			var _rhs38 T2 = (func() T2 {
				if _213_v91 {
					return (_304_v154).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xj")).Cardinality()), _dafny.IntOfUint32((_304_v154).Cardinality()))).Uint32()).(T2)
				}
				return _303_v153
			})()
			_ = _rhs38
			var _rhs39 bool = (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
			_ = _rhs39
			var _lhs18 _dafny.Array = _224_v102
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen((_224_v102), 0))
			_ = _lhs19
			var _lhs20 _dafny.Array = _73_v1
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _lhs21
			(_lhs18).ArraySet1(_rhs37, (_lhs19).Int())
			_303_v153 = _rhs38
			(_lhs20).ArraySet1(_rhs39, (_lhs21).Int())
			var _305_v155 _dafny.MultiSet
			_ = _305_v155
			_305_v155 = _dafny.MultiSetOf(_228_v104, (_245_v111).F10())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index39
			(_73_v1).ArraySet1((_305_v155).IsProperSubsetOf(_305_v155), (_index39).Int())
			var _306_v156 _dafny.Int
			_ = _306_v156
			var _307_v157 bool
			_ = _307_v157
			var _out13 _dafny.Int
			_ = _out13
			var _out14 bool
			_ = _out14
			_out13, _out14 = Companion_Default___.M4(_211_v89, _296_cf45, (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _75_globalState)
			_306_v156 = _out13
			_307_v157 = _out14
		}
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index40
		(_73_v1).ArraySet1((_245_v111).Fm0(_75_globalState), (_index40).Int())
		var _308_v158 _dafny.CodePoint
		_ = _308_v158
		_308_v158 = _228_v104
		var _309_v159 _dafny.Map
		_ = _309_v159
		_309_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, (_308_v158))
		var _rhs40 _dafny.Int = _211_v89
		_ = _rhs40
		var _rhs41 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_211_v89, (_309_v159).Cardinality()), (_212_v90).Plus(_211_v89))
		_ = _rhs41
		_211_v89 = _rhs40
		_211_v89 = _rhs41
		if _213_v91 {
			var _310_v160 *C0
			_ = _310_v160
			var _nw48 *C0 = New_C0_()
			_ = _nw48
			_nw48.Ctor__((_215_v93).F12())
			_310_v160 = _nw48
			var _311_v161 *C1
			_ = _311_v161
			var _nw49 *C1 = New_C1_()
			_ = _nw49
			_nw49.Ctor__((_245_v111).F10(), (_245_v111).F11(), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _296_cf45)
			_311_v161 = _nw49
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index41
			(_73_v1).ArraySet1((func() bool {
				if _213_v91 {
					return _76_v2
				}
				return (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
			})(), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index42
			(_73_v1).ArraySet1(!(_213_v91), (_index42).Int())
			_76_v2 = false
		} else {
			var _312_v162 _dafny.Int
			_ = _312_v162
			var _313_v163 bool
			_ = _313_v163
			var _out15 _dafny.Int
			_ = _out15
			var _out16 bool
			_ = _out16
			_out15, _out16 = (_220_v98).M2(!(((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)) == ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))), _75_globalState)
			_312_v162 = _out15
			_313_v163 = _out16
			_76_v2 = true
			var _314_v164 _dafny.MultiSet
			_ = _314_v164
			_314_v164 = _dafny.MultiSetOf(_313_v163, _213_v91)
			var _315_v165 D0
			_ = _315_v165
			_315_v165 = Companion_D0_.Create_DC0_()
			(_245_v111).M0(_314_v164, _315_v165, (func() bool {
				if _76_v2 {
					return _313_v163
				}
				return _296_cf45
			})(), _75_globalState)
			var _316_v166 _dafny.Array
			_ = _316_v166
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_6
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) D8 = (func(_317_v1 _dafny.Array) func(_dafny.Int) D8 {
					return func(_318_i17 _dafny.Int) D8 {
						return Companion_D8_.Create_DC19_((_317_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_317_v1), 0))).Int()).(bool))
					}
				})(_73_v1)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw50 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw50).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw50).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_316_v166 = _nw50
			var _319_v167 D8
			_ = _319_v167
			_319_v167 = Companion_D8_.Create_DC19_(_296_cf45)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_316_v166), 0))
			_ = _index43
			(_316_v166).ArraySet1(_319_v167, (_index43).Int())
			var _pat_let_tv10 = _213_v91
			_ = _pat_let_tv10
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_316_v166), 0))
			_ = _index44
			(_316_v166).ArraySet1(func(_pat_let14_0 D8) D8 {
				return func(_320_dt__update__tmp_h6 D8) D8 {
					return func(_pat_let15_0 bool) D8 {
						return func(_321_dt__update_hcf31_h0 bool) D8 {
							return Companion_D8_.Create_DC19_(_321_dt__update_hcf31_h0)
						}(_pat_let15_0)
					}(_pat_let_tv10)
				}(_pat_let14_0)
			}(_319_v167), (_index44).Int())
			_211_v89 = _dafny.IntOfUint32(((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)).Cardinality())
		}
	} else if _source6.Is_DC28() {
		var _322___mcc_h17 _dafny.Int = _source6.Get_().(D11_DC28).Cf47
		_ = _322___mcc_h17
		var _323___mcc_h18 _dafny.Array = _source6.Get_().(D11_DC28).Cf48
		_ = _323___mcc_h18
		var _324_cf48 _dafny.Array = _323___mcc_h18
		_ = _324_cf48
		var _325_cf47 _dafny.Int = _322___mcc_h17
		_ = _325_cf47
		var _326_v168 _dafny.Array
		_ = _326_v168
		var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
		_ = _nw51
		_326_v168 = _nw51
		var _327_v169 _dafny.Array
		_ = _327_v169
		var _nwElement0_9 _dafny.Array = (func() _dafny.Array {
			if _213_v91 {
				return _326_v168
			}
			return _326_v168
		})()
		_ = _nwElement0_9
		var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(11))
		_ = _nw52
		(_nw52).ArraySet1(_nwElement0_9, 0)
		(_nw52).ArraySet1(_326_v168, 1)
		(_nw52).ArraySet1(_326_v168, 2)
		(_nw52).ArraySet1(_326_v168, 3)
		(_nw52).ArraySet1(_326_v168, 4)
		(_nw52).ArraySet1(_326_v168, 5)
		(_nw52).ArraySet1(_326_v168, 6)
		(_nw52).ArraySet1(_326_v168, 7)
		(_nw52).ArraySet1((func() _dafny.Array {
			if _213_v91 {
				return _326_v168
			}
			return _326_v168
		})(), 8)
		(_nw52).ArraySet1(_326_v168, 9)
		(_nw52).ArraySet1(_326_v168, 10)
		_327_v169 = _nw52
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_327_v169), 0))
		_ = _index45
		(_327_v169).ArraySet1(_326_v168, (_index45).Int())
		var _328_v170 D13
		_ = _328_v170
		_328_v170 = Companion_D13_.Create_DC33_(_326_v168)
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_327_v169), 0))
		_ = _index46
		(_327_v169).ArraySet1((_328_v170).Dtor_cf56(), (_index46).Int())
		var _329_v171 _dafny.Sequence
		_ = _329_v171
		_329_v171 = _dafny.SeqOf(_dafny.SeqOf(_213_v91))
		var _330_v172 D11
		_ = _330_v172
		_330_v172 = Companion_D11_.Create_DC26_(_329_v171)
		var _331_v173 *C2
		_ = _331_v173
		var _nw53 *C2 = New_C2_()
		_ = _nw53
		_nw53.Ctor__((Companion_D13_.Create_DC34_(_330_v172, (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))).Dtor_cf59(), (_211_v89).Cmp(_212_v90) > 0)
		_331_v173 = _nw53
		_219_v97 = (_219_v97).Update(_325_cf47, (_dafny.Zero).Minus(_212_v90))
		var _332_v174 D0
		_ = _332_v174
		_332_v174 = Companion_D0_.Create_DC1_(_76_v2, _77_v3, _213_v91)
		var _333_v175 *C2
		_ = _333_v175
		var _nw54 *C2 = New_C2_()
		_ = _nw54
		_nw54.Ctor__(_76_v2, (_77_v3).Select((Companion_Default___.SafeIndex((_220_v98).Fm2(_332_v174, _75_globalState), _dafny.IntOfUint32((_77_v3).Cardinality()))).Uint32()).(bool))
		_333_v175 = _nw54
	} else {
		var _334___mcc_h19 _dafny.Sequence = _source6.Get_().(D11_DC26).Cf44
		_ = _334___mcc_h19
		var _335_cf44 _dafny.Sequence = _334___mcc_h19
		_ = _335_cf44
		var _rhs42 _dafny.Int = (_dafny.Zero).Minus(_212_v90)
		_ = _rhs42
		var _rhs43 _dafny.Int = _211_v89
		_ = _rhs43
		var _rhs44 _dafny.Int = _dafny.IntOfInt64(373)
		_ = _rhs44
		_211_v89 = _rhs42
		_211_v89 = _rhs43
		_211_v89 = _rhs44
		var _336_v176 _dafny.Set
		_ = _336_v176
		_336_v176 = _dafny.SetOf((_211_v89).Minus((_dafny.Zero).Minus(_212_v90)))
		var _337_v177 T2
		_ = _337_v177
		var _nw55 *C3 = New_C3_()
		_ = _nw55
		_nw55.Ctor__(_76_v2, _213_v91)
		_337_v177 = _nw55
		var _338_v178 D9
		_ = _338_v178
		_338_v178 = Companion_D9_.Create_DC22_((_216_v94).Cardinality(), _212_v90, _76_v2, _211_v89, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D14_.Create_DC36_(_337_v177)).Dtor_cf61(), _211_v89))
		var _pat_let_tv11 = _212_v90
		_ = _pat_let_tv11
		var _339_v179 _dafny.Sequence
		_ = _339_v179
		_339_v179 = _dafny.SeqOf(_338_v178, _338_v178, func(_pat_let16_0 D9) D9 {
			return func(_340_dt__update__tmp_h7 D9) D9 {
				return func(_pat_let17_0 _dafny.Int) D9 {
					return func(_341_dt__update_hcf36_h0 _dafny.Int) D9 {
						return Companion_D9_.Create_DC22_((_340_dt__update__tmp_h7).Dtor_cf35(), _341_dt__update_hcf36_h0, (_340_dt__update__tmp_h7).Dtor_cf37(), (_340_dt__update__tmp_h7).Dtor_cf38(), (_340_dt__update__tmp_h7).Dtor_cf39())
					}(_pat_let17_0)
				}(_pat_let_tv11)
			}(_pat_let16_0)
		}(_338_v178), _338_v178)
		var _342_v180 _dafny.Set
		_ = _342_v180
		_342_v180 = _dafny.SetOf((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence))
		var _343_v182 _dafny.Map
		_ = _343_v182
		_343_v182 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v177, _76_v2)
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
		_ = _index47
		var _rhs45 _dafny.Set = Companion_Default___.Fm23(_76_v2, func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter12 := _dafny.Iterate((_342_v180).Elements()); ; {
				_compr_10, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _344_v181 _dafny.Sequence
				_344_v181 = interface{}(_compr_10).(_dafny.Sequence)
				if (_342_v180).Contains(_344_v181) {
					_coll10.Add(_344_v181)
				}
			}
			return _coll10.ToSet()
		}(), (_211_v89).Minus((_343_v182).Cardinality()), _75_globalState)
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _339_v179
		_ = _rhs46
		var _rhs47 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_77_v3).Cardinality()), (_dafny.Zero).Minus(_211_v89))
		_ = _rhs47
		var _rhs48 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qwk"), (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence))
		_ = _rhs48
		var _lhs22 _dafny.Array = _229_v105
		_ = _lhs22
		var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
		_ = _lhs23
		_336_v176 = _rhs45
		_339_v179 = _rhs46
		_211_v89 = _rhs47
		(_lhs22).ArraySet1(_rhs48, (_lhs23).Int())
		var _345_v183 *C2
		_ = _345_v183
		var _nw56 *C2 = New_C2_()
		_ = _nw56
		_nw56.Ctor__((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _76_v2)
		_345_v183 = _nw56
		var _346_v184 _dafny.MultiSet
		_ = _346_v184
		_346_v184 = _dafny.MultiSetOf(_345_v183)
		var _347_v185 D0
		_ = _347_v185
		_347_v185 = Companion_D0_.Create_DC1_((_346_v184).IsSubsetOf(_346_v184), _77_v3, (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
		var _348_v186 _dafny.MultiSet
		_ = _348_v186
		_348_v186 = _dafny.MultiSetOf(_228_v104)
		var _349_v187 _dafny.Sequence
		_ = _349_v187
		_349_v187 = _dafny.SeqOf(_345_v183, _345_v183, _345_v183, _345_v183, _345_v183)
		var _pat_let_tv12 = _76_v2
		_ = _pat_let_tv12
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index48
		var _rhs49 bool = true
		_ = _rhs49
		var _rhs50 bool = !(!((_337_v177).F9()) || ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))) || ((_348_v186).IsProperSubsetOf(_dafny.MultiSetOf(Companion_Default___.Fm7(_211_v89, (_337_v177).F9(), _212_v90, _75_globalState))))
		_ = _rhs50
		var _rhs51 D0 = func(_pat_let18_0 D0) D0 {
			return func(_350_dt__update__tmp_h8 D0) D0 {
				return func(_pat_let19_0 bool) D0 {
					return func(_351_dt__update_hcf0_h0 bool) D0 {
						return Companion_D0_.Create_DC1_(_351_dt__update_hcf0_h0, (_350_dt__update__tmp_h8).Dtor_cf1(), (_350_dt__update__tmp_h8).Dtor_cf2())
					}(_pat_let19_0)
				}(_pat_let_tv12)
			}(_pat_let18_0)
		}(_347_v185)
		_ = _rhs51
		var _rhs52 _dafny.Int = _211_v89
		_ = _rhs52
		var _rhs53 *C2 = (_349_v187).Select((Companion_Default___.SafeIndex(_211_v89, _dafny.IntOfUint32((_349_v187).Cardinality()))).Uint32()).(*C2)
		_ = _rhs53
		var _lhs24 _dafny.Array = _73_v1
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs25
		(_lhs24).ArraySet1(_rhs49, (_lhs25).Int())
		_213_v91 = _rhs50
		_347_v185 = _rhs51
		_212_v90 = _rhs52
		_345_v183 = _rhs53
		var _352_v188 _dafny.Array
		_ = _352_v188
		var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw57
		_352_v188 = _nw57
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_352_v188), 0))
		_ = _index49
		(_352_v188).ArraySet1(_229_v105, (_index49).Int())
		var _353_v189 _dafny.Sequence
		_ = _353_v189
		_353_v189 = _dafny.SeqOf(_224_v102)
		var _354_v190 _dafny.Map
		_ = _354_v190
		_354_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_73_v1, _229_v105)
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_352_v188), 0))
		_ = _index50
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index51
		var _rhs54 _dafny.Array = (func() _dafny.Array {
			if (_354_v190).Contains(_73_v1) {
				return (_354_v190).Get(_73_v1).(_dafny.Array)
			}
			return _229_v105
		})()
		_ = _rhs54
		var _rhs55 bool = _76_v2
		_ = _rhs55
		var _rhs56 bool = (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
		_ = _rhs56
		var _rhs57 _dafny.Sequence = _353_v189
		_ = _rhs57
		var _lhs26 _dafny.Array = _352_v188
		_ = _lhs26
		var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_352_v188), 0))
		_ = _lhs27
		var _lhs28 _dafny.Array = _73_v1
		_ = _lhs28
		var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs29
		(_lhs26).ArraySet1(_rhs54, (_lhs27).Int())
		(_lhs28).ArraySet1(_rhs55, (_lhs29).Int())
		_213_v91 = _rhs56
		_353_v189 = _rhs57
	}
	var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
	_ = _index52
	(_73_v1).ArraySet1(_213_v91, (_index52).Int())
	var _355_v191 D6
	_ = _355_v191
	_355_v191 = Companion_D6_.Create_DC13_(_213_v91, true, (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _213_v91)
	var _pat_let_tv13 = _73_v1
	_ = _pat_let_tv13
	var _pat_let_tv14 = _73_v1
	_ = _pat_let_tv14
	var _pat_let_tv15 = _213_v91
	_ = _pat_let_tv15
	var _source7 D6 = func(_pat_let20_0 D6) D6 {
		return func(_356_dt__update__tmp_h9 D6) D6 {
			return func(_pat_let21_0 bool) D6 {
				return func(_357_dt__update_hcf22_h0 bool) D6 {
					return func(_pat_let22_0 bool) D6 {
						return func(_358_dt__update_hcf24_h0 bool) D6 {
							return Companion_D6_.Create_DC13_((_356_dt__update__tmp_h9).Dtor_cf21(), _357_dt__update_hcf22_h0, (_356_dt__update__tmp_h9).Dtor_cf23(), _358_dt__update_hcf24_h0)
						}(_pat_let22_0)
					}(_pat_let_tv15)
				}(_pat_let21_0)
			}((_pat_let_tv14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_pat_let_tv13), 0))).Int()).(bool))
		}(_pat_let20_0)
	}(_355_v191)
	_ = _source7
	if _source7.Is_DC13() {
		var _359___mcc_h20 bool = _source7.Get_().(D6_DC13).Cf21
		_ = _359___mcc_h20
		var _360___mcc_h21 bool = _source7.Get_().(D6_DC13).Cf22
		_ = _360___mcc_h21
		var _361___mcc_h22 _dafny.Sequence = _source7.Get_().(D6_DC13).Cf23
		_ = _361___mcc_h22
		var _362___mcc_h23 bool = _source7.Get_().(D6_DC13).Cf24
		_ = _362___mcc_h23
		var _363_cf24 bool = _362___mcc_h23
		_ = _363_cf24
		var _364_cf23 _dafny.Sequence = _361___mcc_h22
		_ = _364_cf23
		var _365_cf22 bool = _360___mcc_h21
		_ = _365_cf22
		var _366_cf21 bool = _359___mcc_h20
		_ = _366_cf21
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index53
		var _rhs58 bool = !(_366_cf21)
		_ = _rhs58
		var _rhs59 bool = !(_76_v2)
		_ = _rhs59
		var _rhs60 _dafny.Int = (_220_v98).Fm2(Companion_D0_.Create_DC1_(!((_77_v3).Select((Companion_Default___.SafeIndex(_212_v90, _dafny.IntOfUint32((_77_v3).Cardinality()))).Uint32()).(bool)), _77_v3, _76_v2), _75_globalState)
		_ = _rhs60
		var _lhs30 _dafny.Array = _73_v1
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs31
		(_lhs30).ArraySet1(_rhs58, (_lhs31).Int())
		_365_cf22 = _rhs59
		_211_v89 = _rhs60
		var _367_v192 T1
		_ = _367_v192
		var _nw58 *C2 = New_C2_()
		_ = _nw58
		_nw58.Ctor__(_76_v2, _366_cf21)
		_367_v192 = _nw58
		_367_v192 = (func() T1 {
			if _dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-465))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}((func(_368_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_369_i18 _dafny.Int) _dafny.Sequence {
					return _368_v3
				}
			})(_77_v3))), _dafny.SeqOf(_77_v3, _77_v3)) {
				return _367_v192
			}
			return _367_v192
		})()
		_213_v91 = _365_cf22
		_365_cf22 = !(_366_cf21) || ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
	} else {
		var _370___mcc_h24 _dafny.Array = _source7.Get_().(D6_DC12).Cf20
		_ = _370___mcc_h24
		var _371_cf20 _dafny.Array = _370___mcc_h24
		_ = _371_cf20
		_212_v90 = (_dafny.Zero).Minus(_212_v90)
		_76_v2 = !(!((_220_v98).Fm0(_75_globalState)))
		var _372_v193 _dafny.Array
		_ = _372_v193
		var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(11))
		_ = _nw59
		_372_v193 = _nw59
		var _373_v194 _dafny.Sequence
		_ = _373_v194
		_373_v194 = _dafny.SeqOf(_372_v193)
		_372_v193 = (_373_v194).Select((Companion_Default___.SafeIndex(_212_v90, _dafny.IntOfUint32((_373_v194).Cardinality()))).Uint32()).(_dafny.Array)
		_212_v90 = _212_v90
	}
	var _374_v195 D6
	_ = _374_v195
	_374_v195 = Companion_D6_.Create_DC12_((_215_v93).F12())
	_374_v195 = _374_v195
	var _375_v196 _dafny.MultiSet
	_ = _375_v196
	_375_v196 = _dafny.MultiSetOf(_213_v91, _76_v2, _213_v91)
	var _376_v197 _dafny.Sequence
	_ = _376_v197
	_376_v197 = _dafny.SeqOf(_73_v1, _73_v1, _73_v1)
	var _377_v198 _dafny.Set
	_ = _377_v198
	_377_v198 = _dafny.SetOf(_dafny.IntOfInt64(155), _212_v90, (_375_v196).Cardinality(), _dafny.IntOfInt64(436), _dafny.IntOfUint32((_376_v197).Cardinality()))
	var _378_v199 D8
	_ = _378_v199
	_378_v199 = Companion_D8_.Create_DC17_(_377_v198)
	var _pat_let_tv16 = _76_v2
	_ = _pat_let_tv16
	var _pat_let_tv17 = _217_v95
	_ = _pat_let_tv17
	var _pat_let_tv18 = _217_v95
	_ = _pat_let_tv18
	var _pat_let_tv19 = _219_v97
	_ = _pat_let_tv19
	var _pat_let_tv20 = _235_v107
	_ = _pat_let_tv20
	if func(_source8 D8) bool {
		if _source8.Is_DC18() {
			var _379___mcc_h29 _dafny.Map = _source8.Get_().(D8_DC18).Cf28
			_ = _379___mcc_h29
			var _380___mcc_h30 bool = _source8.Get_().(D8_DC18).Cf29
			_ = _380___mcc_h30
			var _381___mcc_h31 _dafny.Int = _source8.Get_().(D8_DC18).Cf30
			_ = _381___mcc_h31
			var _382_cf30 _dafny.Int = _381___mcc_h31
			_ = _382_cf30
			var _383_cf29 bool = _380___mcc_h30
			_ = _383_cf29
			var _384_cf28 _dafny.Map = _379___mcc_h29
			_ = _384_cf28
			return _pat_let_tv16
		} else if _source8.Is_DC19() {
			var _385___mcc_h32 bool = _source8.Get_().(D8_DC19).Cf31
			_ = _385___mcc_h32
			var _386_cf31 bool = _385___mcc_h32
			_ = _386_cf31
			return (_pat_let_tv17).IsDisjointFrom(_pat_let_tv18)
		} else {
			var _387___mcc_h33 _dafny.Set = _source8.Get_().(D8_DC17).Cf27
			_ = _387___mcc_h33
			var _388_cf27 _dafny.Set = _387___mcc_h33
			_ = _388_cf27
			return !(!(((_pat_let_tv19).Cardinality()).Cmp(((Companion_D10_.Create_DC24_(_pat_let_tv20)).Dtor_cf42()).Cardinality()) <= 0))
		}
	}((func() D8 {
		if (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool) {
			return _378_v199
		}
		return _378_v199
	})()) {
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
		_ = _index54
		var _rhs61 _dafny.Sequence = _72_v0
		_ = _rhs61
		var _rhs62 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_227_v103, _227_v103)).Cardinality())).Times(_211_v89)
		_ = _rhs62
		var _lhs32 _dafny.Array = _229_v105
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
		_ = _lhs33
		(_lhs32).ArraySet1(_rhs61, (_lhs33).Int())
		_212_v90 = _rhs62
		var _389_v200 _dafny.Map
		_ = _389_v200
		_389_v200 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-452), _213_v91)
		var _390_v201 _dafny.MultiSet
		_ = _390_v201
		_390_v201 = _dafny.MultiSetOf(_389_v200)
		var _rhs63 _dafny.Int = _212_v90
		_ = _rhs63
		var _rhs64 _dafny.Int = (_211_v89).Plus(_dafny.IntOfInt64(-447))
		_ = _rhs64
		var _rhs65 _dafny.Sequence = Companion_Default___.Fm14(_212_v90, (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _76_v2, _75_globalState)
		_ = _rhs65
		var _rhs66 _dafny.Int = (_dafny.Zero).Minus((((func() _dafny.MultiSet {
			if _213_v91 {
				return _390_v201
			}
			return _390_v201
		})()).Difference(_390_v201)).Cardinality())
		_ = _rhs66
		_211_v89 = _rhs63
		_212_v90 = _rhs64
		_227_v103 = _rhs65
		_212_v90 = _rhs66
		var _391_v202 _dafny.Map
		_ = _391_v202
		_391_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v111, _213_v91)
		_391_v202 = (_391_v202).Update(_245_v111, (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
		_73_v1 = _73_v1
		_211_v89 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm4(_211_v89, _212_v90, _75_globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-753))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}((func(_392_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
			return func(_393_i19 _dafny.Int) _dafny.CodePoint {
				return (_392_v111).F10()
			}
		})(_245_v111))), _72_v0))).Cardinality()))
	} else {
		var _source9 D3 = Companion_D3_.Create_DC6_((Companion_D6_.Create_DC13_(false, _213_v91, _72_v0, _213_v91)).Dtor_cf22(), (_dafny.Zero).Minus((_216_v94).Cardinality()), _212_v90)
		_ = _source9
		if _source9.Is_DC6() {
			var _394___mcc_h25 bool = _source9.Get_().(D3_DC6).Cf7
			_ = _394___mcc_h25
			var _395___mcc_h26 _dafny.Int = _source9.Get_().(D3_DC6).Cf8
			_ = _395___mcc_h26
			var _396___mcc_h27 _dafny.Int = _source9.Get_().(D3_DC6).Cf9
			_ = _396___mcc_h27
			var _397_cf9 _dafny.Int = _396___mcc_h27
			_ = _397_cf9
			var _398_cf8 _dafny.Int = _395___mcc_h26
			_ = _398_cf8
			var _399_cf7 bool = _394___mcc_h25
			_ = _399_cf7
			var _400_v203 *C3
			_ = _400_v203
			var _nw60 *C3 = New_C3_()
			_ = _nw60
			_nw60.Ctor__(_76_v2, (_77_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_77_v3).Cardinality()), _dafny.IntOfUint32((_77_v3).Cardinality()))).Uint32()).(bool))
			_400_v203 = _nw60
			var _401_v204 _dafny.Sequence
			_ = _401_v204
			var _402_v205 _dafny.Int
			_ = _402_v205
			var _403_v206 bool
			_ = _403_v206
			var _out17 _dafny.Sequence
			_ = _out17
			var _out18 _dafny.Int
			_ = _out18
			var _out19 bool
			_ = _out19
			_out17, _out18, _out19 = (_220_v98).M1(_72_v0, _399_cf7, _75_globalState)
			_401_v204 = _out17
			_402_v205 = _out18
			_403_v206 = _out19
			_76_v2 = (_212_v90).Cmp((_397_cf9).Plus(_397_cf9)) >= 0
			var _404_v207 *C2
			_ = _404_v207
			var _nw61 *C2 = New_C2_()
			_ = _nw61
			_nw61.Ctor__((_403_v206) || (true), _403_v206)
			_404_v207 = _nw61
		} else {
			var _405___mcc_h28 _dafny.Sequence = _source9.Get_().(D3_DC5).Cf6
			_ = _405___mcc_h28
			var _406_cf6 _dafny.Sequence = _405___mcc_h28
			_ = _406_cf6
			var _407_v208 T2
			_ = _407_v208
			var _nw62 *C3 = New_C3_()
			_ = _nw62
			_nw62.Ctor__((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _213_v91)
			_407_v208 = _nw62
			var _408_v209 D14
			_ = _408_v209
			_408_v209 = Companion_D14_.Create_DC36_(_407_v208)
			var _409_v210 _dafny.Map
			_ = _409_v210
			_409_v210 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_211_v89, _211_v89), _408_v209)
			_409_v210 = (_409_v210).Update((func() _dafny.Int {
				if (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool) {
					return _212_v90
				}
				return _212_v90
			})(), _408_v209)
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_224_v102), 0))
			_ = _index55
			(_224_v102).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_410_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_411_i20 _dafny.Int) _dafny.CodePoint {
					return (_410_v111).F10()
				}
			})(_245_v111))), (Companion_Default___.SafeIndex(_212_v90, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}((func(_412_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_413_i20 _dafny.Int) _dafny.CodePoint {
					return (_412_v111).F10()
				}
			})(_245_v111)))).Cardinality()))).Uint32(), _228_v104)).Cardinality()), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_224_v102), 0))
			_ = _index56
			(_224_v102).ArraySet1(_211_v89, (_index56).Int())
			_213_v91 = ((_212_v90).Times(_dafny.IntOfInt64(929))).Cmp((_dafny.Zero).Minus(_211_v89)) >= 0
			var _414_v211 _dafny.Map
			_ = _414_v211
			_414_v211 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_407_v208).F9(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-383))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}((func(_415_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_416_i21 _dafny.Int) _dafny.CodePoint {
					return (_415_v111).F10()
				}
			})(_245_v111))))
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_224_v102), 0))
			_ = _index57
			(_224_v102).ArraySet1(((_414_v211).Update(!_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(413))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_417_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_418_i22 _dafny.Int) _dafny.CodePoint {
					return (_417_v111).F10()
				}
			})(_245_v111))), (_245_v111).F10()), _dafny.Companion_Sequence_.Concatenate((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _72_v0))).Cardinality(), (_index57).Int())
		}
		var _419_v212 _dafny.Sequence
		_ = _419_v212
		var _420_v213 _dafny.Int
		_ = _420_v213
		var _421_v214 bool
		_ = _421_v214
		var _out20 _dafny.Sequence
		_ = _out20
		var _out21 _dafny.Int
		_ = _out21
		var _out22 bool
		_ = _out22
		_out20, _out21, _out22 = (_245_v111).M1((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _76_v2, _75_globalState)
		_419_v212 = _out20
		_420_v213 = _out21
		_421_v214 = _out22
		var _422_v215 D12
		_ = _422_v215
		_422_v215 = Companion_D12_.Create_DC31_(_76_v2, (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _212_v90)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index58
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index59
		var _rhs67 _dafny.Int = Companion_Default___.SafeDivisionInt(_420_v213, _212_v90)
		_ = _rhs67
		var _rhs68 _dafny.Array = _214_v92
		_ = _rhs68
		var _rhs69 bool = (_422_v215).Dtor_cf53()
		_ = _rhs69
		var _rhs70 bool = _421_v214
		_ = _rhs70
		var _lhs34 _dafny.Array = _73_v1
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs35
		var _lhs36 _dafny.Array = _73_v1
		_ = _lhs36
		var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs37
		_211_v89 = _rhs67
		_214_v92 = _rhs68
		(_lhs34).ArraySet1(_rhs69, (_lhs35).Int())
		(_lhs36).ArraySet1(_rhs70, (_lhs37).Int())
		var _423_v216 D0
		_ = _423_v216
		_423_v216 = Companion_D0_.Create_DC0_()
		(_220_v98).M0(_375_v196, _423_v216, _76_v2, _75_globalState)
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_224_v102), 0))
		_ = _index60
		(_224_v102).ArraySet1(_212_v90, (_index60).Int())
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(131), _dafny.ArrayLen((_224_v102), 0))
		_ = _index61
		(_224_v102).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(390), _420_v213), _211_v89), (_index61).Int())
	}
	var _424_v217 _dafny.Map
	_ = _424_v217
	_424_v217 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, _217_v95)
	var _425_v218 D11
	_ = _425_v218
	_425_v218 = Companion_D11_.Create_DC27_((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _424_v217)
	var _source10 D11 = _425_v218
	_ = _source10
	if _source10.Is_DC27() {
		var _426___mcc_h34 bool = _source10.Get_().(D11_DC27).Cf45
		_ = _426___mcc_h34
		var _427___mcc_h35 _dafny.Map = _source10.Get_().(D11_DC27).Cf46
		_ = _427___mcc_h35
		var _428_cf46 _dafny.Map = _427___mcc_h35
		_ = _428_cf46
		var _429_cf45 bool = _426___mcc_h34
		_ = _429_cf45
		_212_v90 = _212_v90
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index62
		(_73_v1).ArraySet1(_429_cf45, (_index62).Int())
		var _430_v219 D12
		_ = _430_v219
		_430_v219 = Companion_D12_.Create_DC31_(_429_cf45, _429_cf45, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)), _dafny.SeqOf((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _76_v2))).Cardinality()))
		var _rhs71 bool = (_212_v90).Cmp(_212_v90) == 0
		_ = _rhs71
		var _rhs72 _dafny.Int = (_dafny.Zero).Minus((_211_v89).Minus(Companion_Default___.SafeModuloInt(_212_v90, (_219_v97).Cardinality())))
		_ = _rhs72
		var _rhs73 D12 = _430_v219
		_ = _rhs73
		_429_cf45 = _rhs71
		_212_v90 = _rhs72
		_430_v219 = _rhs73
		_235_v107 = (_235_v107).Update((_213_v91) && ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)), (_dafny.MultiSetOf(_dafny.IntOfInt64(-973))).IsDisjointFrom(_dafny.MultiSetOf(_211_v89)))
	} else if _source10.Is_DC28() {
		var _431___mcc_h36 _dafny.Int = _source10.Get_().(D11_DC28).Cf47
		_ = _431___mcc_h36
		var _432___mcc_h37 _dafny.Array = _source10.Get_().(D11_DC28).Cf48
		_ = _432___mcc_h37
		var _433_cf48 _dafny.Array = _432___mcc_h37
		_ = _433_cf48
		var _434_cf47 _dafny.Int = _431___mcc_h36
		_ = _434_cf47
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))
		_ = _index63
		(_224_v102).ArraySet1((_dafny.Zero).Minus((_211_v89).Times(_211_v89)), (_index63).Int())
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index64
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))
		_ = _index65
		var _rhs74 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_211_v89, _434_cf47), _227_v103), _227_v103)
		_ = _rhs74
		var _rhs75 _dafny.Int = (func() _dafny.Int {
			if _76_v2 {
				return _211_v89
			}
			return _211_v89
		})()
		_ = _rhs75
		var _rhs76 bool = (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
		_ = _rhs76
		var _lhs38 _dafny.Array = _73_v1
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs39
		var _lhs40 _dafny.Array = _224_v102
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))
		_ = _lhs41
		(_lhs38).ArraySet1(_rhs74, (_lhs39).Int())
		(_lhs40).ArraySet1(_rhs75, (_lhs41).Int())
		_76_v2 = _rhs76
		var _435_v220 D12
		_ = _435_v220
		_435_v220 = Companion_D12_.Create_DC31_(_76_v2, _76_v2, _dafny.IntOfInt64(624))
		var _source11 D12 = _435_v220
		_ = _source11
		if _source11.Is_DC30() {
			var _436___mcc_h39 _dafny.Int = _source11.Get_().(D12_DC30).Cf50
			_ = _436___mcc_h39
			var _437___mcc_h40 _dafny.Int = _source11.Get_().(D12_DC30).Cf51
			_ = _437___mcc_h40
			var _438_cf51 _dafny.Int = _437___mcc_h40
			_ = _438_cf51
			var _439_cf50 _dafny.Int = _436___mcc_h39
			_ = _439_cf50
			var _440_v221 *C2
			_ = _440_v221
			var _nw63 *C2 = New_C2_()
			_ = _nw63
			_nw63.Ctor__(_76_v2, (func() bool {
				if (_235_v107).Contains(_213_v91) {
					return (_235_v107).Get(_213_v91).(bool)
				}
				return _76_v2
			})())
			_440_v221 = _nw63
			var _441_v222 _dafny.Map
			_ = _441_v222
			_441_v222 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_440_v221, _433_cf48)
			_441_v222 = (_441_v222).Update(_440_v221, _433_cf48)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index66
			(_73_v1).ArraySet1(true, (_index66).Int())
			_438_cf51 = _438_cf51
			var _442_v223 *C3
			_ = _442_v223
			var _nw64 *C3 = New_C3_()
			_ = _nw64
			_nw64.Ctor__(true, _76_v2)
			_442_v223 = _nw64
		} else if _source11.Is_DC31() {
			var _443___mcc_h41 bool = _source11.Get_().(D12_DC31).Cf52
			_ = _443___mcc_h41
			var _444___mcc_h42 bool = _source11.Get_().(D12_DC31).Cf53
			_ = _444___mcc_h42
			var _445___mcc_h43 _dafny.Int = _source11.Get_().(D12_DC31).Cf54
			_ = _445___mcc_h43
			var _446_cf54 _dafny.Int = _445___mcc_h43
			_ = _446_cf54
			var _447_cf53 bool = _444___mcc_h42
			_ = _447_cf53
			var _448_cf52 bool = _443___mcc_h41
			_ = _448_cf52
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index67
			(_73_v1).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32(((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)).Cardinality()))).IsSubsetOf((_217_v95).Union(_dafny.MultiSetFromSeq(_227_v103))), (_index67).Int())
			var _449_v224 _dafny.Int
			_ = _449_v224
			var _450_v225 bool
			_ = _450_v225
			var _out23 _dafny.Int
			_ = _out23
			var _out24 bool
			_ = _out24
			_out23, _out24 = Companion_Default___.M4(_dafny.IntOfUint32((_dafny.SeqOf(_211_v89)).Cardinality()), _213_v91, _dafny.UnicodeSeqOfUtf8Bytes("xexs"), _75_globalState)
			_449_v224 = _out23
			_450_v225 = _out24
			var _451_v226 _dafny.Int
			_ = _451_v226
			var _452_v227 bool
			_ = _452_v227
			var _out25 _dafny.Int
			_ = _out25
			var _out26 bool
			_ = _out26
			_out25, _out26 = Companion_Default___.M4(_212_v90, ((_dafny.Zero).Minus(_211_v89)).Cmp(_449_v224) < 0, (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _75_globalState)
			_451_v226 = _out25
			_452_v227 = _out26
			_216_v94 = (_216_v94).Update(_213_v91, (_449_v224).Minus((_dafny.Zero).Minus(_434_cf47)))
		} else if _source11.Is_DC29() {
			var _453___mcc_h44 *C3 = _source11.Get_().(D12_DC29).Cf49
			_ = _453___mcc_h44
			var _454_cf49 *C3 = _453___mcc_h44
			_ = _454_cf49
			var _455_v228 T1
			_ = _455_v228
			var _nw65 *C3 = New_C3_()
			_ = _nw65
			_nw65.Ctor__(_76_v2, _76_v2)
			_455_v228 = _nw65
			_455_v228 = _455_v228
			var _456_v229 D9
			_ = _456_v229
			_456_v229 = Companion_D9_.Create_DC21_(_dafny.IntOfInt64(808), (_dafny.Zero).Minus((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)))
			var _457_v230 _dafny.Array
			_ = _457_v230
			var _nwElement0_10 _dafny.Int = _211_v89
			_ = _nwElement0_10
			var _nw66 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(26))
			_ = _nw66
			(_nw66).ArraySet1(_nwElement0_10, 0)
			(_nw66).ArraySet1(_dafny.IntOfInt64(585), 1)
			(_nw66).ArraySet1((_212_v90).Plus((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)), 2)
			(_nw66).ArraySet1(_212_v90, 3)
			(_nw66).ArraySet1((func() _dafny.Int {
				if (_219_v97).Contains(_211_v89) {
					return (_219_v97).Get(_211_v89).(_dafny.Int)
				}
				return _211_v89
			})(), 4)
			(_nw66).ArraySet1(_434_cf47, 5)
			(_nw66).ArraySet1((_211_v89).Plus(_211_v89), 6)
			(_nw66).ArraySet1(_dafny.IntOfUint32(((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)).Cardinality()), 7)
			(_nw66).ArraySet1(((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)).Plus(_211_v89), 8)
			(_nw66).ArraySet1(_434_cf47, 9)
			(_nw66).ArraySet1(_212_v90, 10)
			(_nw66).ArraySet1(_dafny.IntOfUint32((_72_v0).Cardinality()), 11)
			(_nw66).ArraySet1(((_dafny.Zero).Minus(_212_v90)).Minus(_212_v90), 12)
			(_nw66).ArraySet1((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int), 13)
			(_nw66).ArraySet1((_dafny.Zero).Minus((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)), 14)
			(_nw66).ArraySet1(_dafny.IntOfUint32((_77_v3).Cardinality()), 15)
			(_nw66).ArraySet1(_212_v90, 16)
			(_nw66).ArraySet1(_434_cf47, 17)
			(_nw66).ArraySet1(_434_cf47, 18)
			(_nw66).ArraySet1(((_456_v229).Dtor_cf33()).Times(_211_v89), 19)
			(_nw66).ArraySet1((_211_v89).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(421))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_458_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_459_i23 _dafny.Int) _dafny.CodePoint {
					return (_458_v111).F10()
				}
			})(_245_v111)))).Cardinality())), 20)
			(_nw66).ArraySet1(((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)).Minus(_212_v90), 21)
			(_nw66).ArraySet1(_dafny.IntOfInt64(-503), 22)
			(_nw66).ArraySet1((_dafny.Zero).Minus((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)), 23)
			(_nw66).ArraySet1(_434_cf47, 24)
			(_nw66).ArraySet1((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int), 25)
			_457_v230 = _nw66
			var _rhs77 _dafny.Array = _457_v230
			_ = _rhs77
			var _rhs78 _dafny.Int = _212_v90
			_ = _rhs78
			_457_v230 = _rhs77
			_211_v89 = _rhs78
			var _460_v231 T2
			_ = _460_v231
			var _nw67 *C3 = New_C3_()
			_ = _nw67
			_nw67.Ctor__(_455_v228.F8(), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
			_460_v231 = _nw67
			var _461_v232 _dafny.Sequence
			_ = _461_v232
			_461_v232 = _dafny.SeqOf(_242_v108)
			var _nw68 *C1 = New_C1_()
			_ = _nw68
			_nw68.Ctor__((_245_v111).F10(), ((_245_v111).F11()).Difference(_dafny.SetOf(_461_v232, _461_v232)), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), (_460_v231).F9())
			_460_v231 = _nw68
			var _462_v233 D0
			_ = _462_v233
			_462_v233 = Companion_D0_.Create_DC1_(_76_v2, _77_v3, _455_v228.F8())
			_76_v2 = ((!(_455_v228.F8())) || ((_245_v111).Fm1((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _462_v233, _75_globalState))) == ((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
		} else {
			var _463___mcc_h45 D12 = _source11.Get_().(D12_DC32).Cf55
			_ = _463___mcc_h45
			var _464_cf55 D12 = _463___mcc_h45
			_ = _464_cf55
			var _465_v234 _dafny.Array
			_ = _465_v234
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(12))
			_ = _nw69
			_465_v234 = _nw69
			var _466_v235 T1
			_ = _466_v235
			var _nw70 *C3 = New_C3_()
			_ = _nw70
			_nw70.Ctor__((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), _213_v91)
			_466_v235 = _nw70
			var _467_v236 D9
			_ = _467_v236
			_467_v236 = Companion_D9_.Create_DC23_(_466_v235, _76_v2)
			var _pat_let_tv21 = _76_v2
			_ = _pat_let_tv21
			var _468_v237 _dafny.Array
			_ = _468_v237
			var _nwElement0_11 D9 = _467_v236
			_ = _nwElement0_11
			var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(15))
			_ = _nw71
			(_nw71).ArraySet1(_nwElement0_11, 0)
			(_nw71).ArraySet1(_467_v236, 1)
			(_nw71).ArraySet1(_467_v236, 2)
			(_nw71).ArraySet1(_467_v236, 3)
			(_nw71).ArraySet1(Companion_D9_.Create_DC23_(_466_v235, !(_76_v2)), 4)
			(_nw71).ArraySet1(func(_pat_let23_0 D9) D9 {
				return func(_469_dt__update__tmp_h10 D9) D9 {
					return func(_pat_let24_0 bool) D9 {
						return func(_470_dt__update_hcf41_h0 bool) D9 {
							return Companion_D9_.Create_DC23_((_469_dt__update__tmp_h10).Dtor_cf40(), _470_dt__update_hcf41_h0)
						}(_pat_let24_0)
					}(_pat_let_tv21)
				}(_pat_let23_0)
			}(_467_v236), 5)
			(_nw71).ArraySet1(_467_v236, 6)
			(_nw71).ArraySet1(_467_v236, 7)
			(_nw71).ArraySet1(_467_v236, 8)
			(_nw71).ArraySet1(_467_v236, 9)
			(_nw71).ArraySet1(_467_v236, 10)
			(_nw71).ArraySet1(_467_v236, 11)
			(_nw71).ArraySet1(_467_v236, 12)
			(_nw71).ArraySet1(Companion_D9_.Create_DC23_(_466_v235, !((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))), 13)
			(_nw71).ArraySet1(_467_v236, 14)
			_468_v237 = _nw71
			var _471_v238 _dafny.Map
			_ = _471_v238
			_471_v238 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_212_v90, _220_v98)
			var _472_v239 _dafny.Map
			_ = _472_v239
			_472_v239 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v237, ((_471_v238).Update(_dafny.IntOfUint32((_77_v3).Cardinality()), _220_v98)).Cardinality())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_465_v234), 0))
			_ = _index68
			(_465_v234).ArraySet1(_472_v239, (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(962), _dafny.ArrayLen((_465_v234), 0))
			_ = _index69
			(_465_v234).ArraySet1((_472_v239).Merge(_472_v239), (_index69).Int())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))
			_ = _index70
			(_224_v102).ArraySet1((_227_v103).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(330), _434_cf47), _dafny.IntOfUint32((_227_v103).Cardinality()))).Uint32()).(_dafny.Int), (_index70).Int())
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))
			_ = _index71
			(_224_v102).ArraySet1(((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int)).Plus((_dafny.Zero).Minus((_224_v102).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_224_v102), 0))).Int()).(_dafny.Int))), (_index71).Int())
			var _473_v240 D0
			_ = _473_v240
			_473_v240 = Companion_D0_.Create_DC1_(_466_v235.F8(), _77_v3, _213_v91)
			var _474_v241 _dafny.Sequence
			_ = _474_v241
			var _475_v242 _dafny.Int
			_ = _475_v242
			var _476_v243 bool
			_ = _476_v243
			var _out27 _dafny.Sequence
			_ = _out27
			var _out28 _dafny.Int
			_ = _out28
			var _out29 bool
			_ = _out29
			_out27, _out28, _out29 = (_220_v98).M1((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), (_220_v98).Fm1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(275))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_477_v111 *C1) func(_dafny.Int) _dafny.CodePoint {
				return func(_478_i24 _dafny.Int) _dafny.CodePoint {
					return (_477_v111).F10()
				}
			})(_245_v111))), _473_v240, _75_globalState), _75_globalState)
			_474_v241 = _out27
			_475_v242 = _out28
			_476_v243 = _out29
		}
		var _479_v244 *C2
		_ = _479_v244
		var _nw72 *C2 = New_C2_()
		_ = _nw72
		_nw72.Ctor__((!((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))) && (_213_v91), _76_v2)
		_479_v244 = _nw72
		var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index72
		(_73_v1).ArraySet1(_76_v2, (_index72).Int())
	} else {
		var _480___mcc_h38 _dafny.Sequence = _source10.Get_().(D11_DC26).Cf44
		_ = _480___mcc_h38
		var _481_cf44 _dafny.Sequence = _480___mcc_h38
		_ = _481_cf44
		var _482_v245 T0
		_ = _482_v245
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(_214_v92)
		_482_v245 = _nw73
		var _483_v246 _dafny.Map
		_ = _483_v246
		_483_v246 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_72_v0, _72_v0)).Cardinality()), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _index73
		var _rhs79 bool = _76_v2
		_ = _rhs79
		var _rhs80 T0 = _482_v245
		_ = _rhs80
		var _rhs81 bool = (func() bool {
			if (_483_v246).Contains(_211_v89) {
				return (_483_v246).Get(_211_v89).(bool)
			}
			return (_213_v91) == (_213_v91)
		})()
		_ = _rhs81
		var _rhs82 _dafny.Int = (_212_v90).Plus((Companion_Default___.Fm26(_216_v94, _212_v90, _75_globalState)).Cardinality())
		_ = _rhs82
		var _lhs42 _dafny.Array = _73_v1
		_ = _lhs42
		var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
		_ = _lhs43
		_213_v91 = _rhs79
		_482_v245 = _rhs80
		(_lhs42).ArraySet1(_rhs81, (_lhs43).Int())
		_212_v90 = _rhs82
		var _484_v247 D0
		_ = _484_v247
		_484_v247 = Companion_D0_.Create_DC1_((_245_v111).Fm0(_75_globalState), Companion_Default___.Fm27(_75_globalState), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
		if (_215_v93).Fm1(_dafny.Companion_Sequence_.Concatenate(_72_v0, _72_v0), _484_v247, _75_globalState) {
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index74
			(_73_v1).ArraySet1((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool), (_index74).Int())
			_212_v90 = _dafny.IntOfInt64(135)
			_212_v90 = (_211_v89).Times(_dafny.IntOfUint32(((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			var _485_v248 T2
			_ = _485_v248
			var _nw74 *C1 = New_C1_()
			_ = _nw74
			_nw74.Ctor__((_245_v111).F10(), (_245_v111).F11(), _213_v91, true)
			_485_v248 = _nw74
			var _486_v249 _dafny.Sequence
			_ = _486_v249
			_486_v249 = _dafny.SeqOf(_485_v248)
			var _487_v250 _dafny.Map
			_ = _487_v250
			_487_v250 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_486_v249).Select((Companion_Default___.SafeIndex(_211_v89, _dafny.IntOfUint32((_486_v249).Cardinality()))).Uint32()).(T2), !(false))
			_487_v250 = (_487_v250).Update(_485_v248, false)
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_224_v102), 0))
			_ = _index75
			(_224_v102).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mnwoqpjhn")).Cardinality()), (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_224_v102), 0))
			_ = _index76
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index77
			var _rhs83 _dafny.Int = Companion_Default___.SafeDivisionInt(_212_v90, _211_v89)
			_ = _rhs83
			var _rhs84 bool = (_485_v248).F9()
			_ = _rhs84
			var _rhs85 _dafny.Int = (_211_v89).Times(_211_v89)
			_ = _rhs85
			var _lhs44 _dafny.Array = _224_v102
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_224_v102), 0))
			_ = _lhs45
			var _lhs46 _dafny.Array = _73_v1
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _lhs47
			(_lhs44).ArraySet1(_rhs83, (_lhs45).Int())
			(_lhs46).ArraySet1(_rhs84, (_lhs47).Int())
			_211_v89 = _rhs85
		} else {
			_212_v90 = _211_v89
			_213_v91 = _213_v91
			var _488_v251 _dafny.Array
			_ = _488_v251
			var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(26))
			_ = _nw75
			_488_v251 = _nw75
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_488_v251), 0))
			_ = _index78
			(_488_v251).ArraySet1(_375_v196, (_index78).Int())
			var _489_v252 _dafny.CodePoint
			_ = _489_v252
			_489_v252 = (_245_v111).F10()
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_488_v251), 0))
			_ = _index79
			(_488_v251).ArraySet1(Companion_Default___.Fm19(_76_v2, (_245_v111).F10(), (_dafny.IntOfInt64(589)).Plus(_dafny.IntOfUint32(((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)).Cardinality())), _75_globalState), (_index79).Int())
			var _490_v253 _dafny.Map
			_ = _490_v253
			_490_v253 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm4(_211_v89, _211_v89, _75_globalState), (Companion_Default___.SafeIndex(_212_v90, _dafny.IntOfUint32((Companion_Default___.Fm4(_211_v89, _211_v89, _75_globalState)).Cardinality()))).Uint32(), (_245_v111).F10()), (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool))
			var _491_v254 _dafny.Map
			_ = _491_v254
			_491_v254 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_211_v89).Times(_dafny.IntOfInt64(420)), _490_v253)
			_491_v254 = (_491_v254).Update(Companion_Default___.SafeModuloInt(_212_v90, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tbnxyymh")).Cardinality())), (func() _dafny.Map {
				if (_491_v254).Contains((_219_v97).Cardinality()) {
					return (_491_v254).Get((_219_v97).Cardinality()).(_dafny.Map)
				}
				return _490_v253
			})())
			var _492_v255 _dafny.Sequence
			_ = _492_v255
			var _493_v256 _dafny.Int
			_ = _493_v256
			var _494_v257 bool
			_ = _494_v257
			var _out30 _dafny.Sequence
			_ = _out30
			var _out31 _dafny.Int
			_ = _out31
			var _out32 bool
			_ = _out32
			_out30, _out31, _out32 = (_220_v98).M1(_dafny.Companion_Sequence_.Concatenate((_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence), _72_v0), !((_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)) || (false), _75_globalState)
			_492_v255 = _out30
			_493_v256 = _out31
			_494_v257 = _out32
		}
		var _495_v258 _dafny.MultiSet
		_ = _495_v258
		_495_v258 = _dafny.MultiSetOf(_224_v102)
		var _496_v259 T1
		_ = _496_v259
		var _nw76 *C1 = New_C1_()
		_ = _nw76
		_nw76.Ctor__((_245_v111).F10(), (_245_v111).F11(), !((_dafny.MultiSetOf(_224_v102)).IsDisjointFrom(_495_v258)), true)
		_496_v259 = _nw76
		var _497_v260 _dafny.Map
		_ = _497_v260
		_497_v260 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, _243_v109)
		var _nw77 *C1 = New_C1_()
		_ = _nw77
		_nw77.Ctor__((_245_v111).F10(), ((func() _dafny.Set {
			if (_497_v260).Contains(_212_v90) {
				return (_497_v260).Get(_212_v90).(_dafny.Set)
			}
			return (func() _dafny.Set {
				if (_497_v260).Contains(_212_v90) {
					return (_497_v260).Get(_212_v90).(_dafny.Set)
				}
				return (_245_v111).F11()
			})()
		})()).Intersection(_243_v109), _76_v2, _213_v91)
		_496_v259 = _nw77
		if _76_v2 {
			var _498_v261 _dafny.Set
			_ = _498_v261
			_498_v261 = _dafny.SetOf(_224_v102)
			var _499_v262 *C2
			_ = _499_v262
			var _nw78 *C2 = New_C2_()
			_ = _nw78
			_nw78.Ctor__(_213_v91, (_498_v261).IsSubsetOf(_498_v261))
			_499_v262 = _nw78
			var _500_v263 _dafny.MultiSet
			_ = _500_v263
			_500_v263 = _dafny.MultiSetOf(_72_v0)
			var _501_v264 D15
			_ = _501_v264
			_501_v264 = Companion_D15_.Create_DC38_(_499_v262)
			var _502_v265 _dafny.Map
			_ = _502_v265
			_502_v265 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_377_v198).Cardinality()).Cmp(_dafny.IntOfInt64(775)) != 0, (_501_v264).Dtor_cf66())
			var _503_v266 _dafny.Map
			_ = _503_v266
			_503_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_224_v102, _72_v0)
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index80
			var _rhs86 bool = (_500_v263).IsDisjointFrom(_500_v263)
			_ = _rhs86
			var _rhs87 *C2 = (func() *C2 {
				if (_502_v265).Contains((func() bool {
					if _496_v259.F8() {
						return _213_v91
					}
					return (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
				})()) {
					return (_502_v265).Get((func() bool {
						if _496_v259.F8() {
							return _213_v91
						}
						return (_73_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))).Int()).(bool)
					})()).(*C2)
				}
				return _499_v262
			})()
			_ = _rhs87
			var _rhs88 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_72_v0, (func() _dafny.Sequence {
				if (_503_v266).Contains(_224_v102) {
					return (_503_v266).Get(_224_v102).(_dafny.Sequence)
				}
				return (_229_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))).Int()).(_dafny.Sequence)
			})())
			_ = _rhs88
			var _lhs48 _dafny.Array = _73_v1
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _lhs49
			(_lhs48).ArraySet1(_rhs86, (_lhs49).Int())
			_499_v262 = _rhs87
			_72_v0 = _rhs88
			var _504_v267 _dafny.Sequence
			_ = _504_v267
			_504_v267 = _dafny.SeqOf(_217_v95)
			var _505_v268 _dafny.Set
			_ = _505_v268
			_505_v268 = _dafny.SetOf(((_504_v267).Select((Companion_Default___.SafeIndex(_211_v89, _dafny.IntOfUint32((_504_v267).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsSubsetOf(_217_v95), _76_v2)
			_505_v268 = _505_v268
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
			_ = _index81
			var _rhs89 _dafny.Int = _211_v89
			_ = _rhs89
			var _rhs90 _dafny.Sequence = _72_v0
			_ = _rhs90
			var _lhs50 _dafny.Array = _229_v105
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
			_ = _lhs51
			_211_v89 = _rhs89
			(_lhs50).ArraySet1(_rhs90, (_lhs51).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(762), _dafny.ArrayLen((_229_v105), 0))
			_ = _index82
			(_229_v105).ArraySet1(_72_v0, (_index82).Int())
			var _506_v269 _dafny.Map
			_ = _506_v269
			_506_v269 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, _245_v111)
			_506_v269 = _506_v269
		} else {
			var _507_v270 _dafny.Set
			_ = _507_v270
			_507_v270 = _dafny.SetOf(_496_v259.F8())
			var _508_v271 *C1
			_ = _508_v271
			var _nw79 *C1 = New_C1_()
			_ = _nw79
			_nw79.Ctor__(Companion_Default___.Fm7((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_245_v111, (_377_v198).Cardinality())).Cardinality(), _76_v2, _211_v89, _75_globalState), ((_245_v111).F11()).Union((_245_v111).F11()), (_507_v270).IsProperSubsetOf(_507_v270), _213_v91)
			_508_v271 = _nw79
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen(((_215_v93).F12()), 0))
			_ = _index83
			((_215_v93).F12()).ArraySet1((Companion_D5_.Create_DC10_(_224_v102)).Dtor_cf17(), (_index83).Int())
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(370), _dafny.ArrayLen(((_215_v93).F12()), 0))
			_ = _index84
			((_215_v93).F12()).ArraySet1(_224_v102, (_index84).Int())
			var _509_v272 _dafny.Map
			_ = _509_v272
			_509_v272 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v89, _377_v198)
			_483_v246 = (_483_v246).Update(_dafny.IntOfInt64(929), (_377_v198).IsProperSubsetOf((func() _dafny.Set {
				if (_509_v272).Contains(_212_v90) {
					return (_509_v272).Get(_212_v90).(_dafny.Set)
				}
				return _dafny.SetOf(_212_v90, (_377_v198).Cardinality(), _212_v90)
			})()))
			(_496_v259).F8_set_(_76_v2)
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_73_v1), 0))
			_ = _index85
			(_73_v1).ArraySet1(!(!(_496_v259.F8())) || (_496_v259.F8()), (_index85).Int())
		}
	}
	_dafny.Print(_72_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_73_v1).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_75_globalState).F6()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_75_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_76_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_77_v3, _dafny.SeqOf(false, false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_211_v89)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_212_v90)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_213_v91)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v94).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(14))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_217_v95).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v96).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v97).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(-27058))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_221_v99).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_221_v99).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_222_v100).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_222_v100).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.Zero).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.Zero).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.One).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.One).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_223_v101).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D12)).Dtor_cf55()).Dtor_cf49().F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_223_v101).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(D12)).Dtor_cf55()).Dtor_cf49()).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_224_v102).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_227_v103, _dafny.SeqOf(_dafny.One, _dafny.IntOfInt64(-881), _dafny.IntOfInt64(473), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_v104)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v105).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_229_v105).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v106).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v107).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_242_v108).Dtor_cf10().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_243_v109).Equals(_dafny.SetOf(_dafny.SeqOf(Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe"))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_244_v110).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_245_v111).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_245_v111).F11()).Equals(_dafny.SetOf(_dafny.SeqOf(Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe")), Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("xe"))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_246_i11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v191).Dtor_cf21())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v191).Dtor_cf22())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v191).Dtor_cf23().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_355_v191).Dtor_cf24())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_375_v196).Equals(_dafny.MultiSetOf(true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_376_v197).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_377_v198).Equals(_dafny.SetOf(_dafny.IntOfInt64(155), _dafny.IntOfInt64(-161), _dafny.IntOfInt64(3), _dafny.IntOfInt64(436))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_378_v199).Dtor_cf27()).Equals(_dafny.SetOf(_dafny.IntOfInt64(155), _dafny.IntOfInt64(-161), _dafny.IntOfInt64(3), _dafny.IntOfInt64(436))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_424_v217).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-761), _dafny.MultiSetOf(_dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_425_v218).Dtor_cf45())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_425_v218).Dtor_cf46()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-761), _dafny.MultiSetOf(_dafny.One))))
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
	Cf0 bool
	Cf1 _dafny.Sequence
	Cf2 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf0 bool, Cf1 _dafny.Sequence, Cf2 bool) D0 {
	return D0{D0_DC1{Cf0, Cf1, Cf2}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_()
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC1).Cf0
}

func (_this D0) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
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
			return "D0.DC1" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
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
			return ok && data1.Cf0 == data2.Cf0 && data1.Cf1.Equals(data2.Cf1) && data1.Cf2 == data2.Cf2
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
	Cf3 _dafny.CodePoint
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf3 _dafny.CodePoint) D1 {
	return D1{D1_DC2{Cf3}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() _dafny.CodePoint {
	return _dafny.CodePoint('D')
}

func (_this D1) Dtor_cf3() _dafny.CodePoint {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf3) + ")"
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

type D2_DC4 struct {
	Cf5 bool
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf5 bool) D2 {
	return D2{D2_DC4{Cf5}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC3 struct {
	Cf4 _dafny.Int
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf4 _dafny.Int) D2 {
	return D2{D2_DC3{Cf4}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(false)
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC4).Cf5
}

func (_this D2) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D2_DC3).Cf4
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf5) + ")"
		}
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
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf5 == data2.Cf5
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf7 bool
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf7 bool, Cf8 _dafny.Int, Cf9 _dafny.Int) D3 {
	return D3{D3_DC6{Cf7, Cf8, Cf9}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC5 struct {
	Cf6 _dafny.Sequence
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf6 _dafny.Sequence) D3 {
	return D3{D3_DC5{Cf6}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D3) Dtor_cf7() bool {
	return _this.Get_().(D3_DC6).Cf7
}

func (_this D3) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf8
}

func (_this D3) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf9
}

func (_this D3) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D3_DC5).Cf6
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf6.Equals(data2.Cf6)
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

type D4_DC8 struct {
	Cf11 _dafny.Int
	Cf12 bool
}

func (D4_DC8) isD4() {}

func (CompanionStruct_D4_) Create_DC8_(Cf11 _dafny.Int, Cf12 bool) D4 {
	return D4{D4_DC8{Cf11, Cf12}}
}

func (_this D4) Is_DC8() bool {
	_, ok := _this.Get_().(D4_DC8)
	return ok
}

type D4_DC9 struct {
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 _dafny.Array
	Cf16 _dafny.Int
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf13 bool, Cf14 _dafny.Int, Cf15 _dafny.Array, Cf16 _dafny.Int) D4 {
	return D4{D4_DC9{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

type D4_DC7 struct {
	Cf10 _dafny.Sequence
}

func (D4_DC7) isD4() {}

func (CompanionStruct_D4_) Create_DC7_(Cf10 _dafny.Sequence) D4 {
	return D4{D4_DC7{Cf10}}
}

func (_this D4) Is_DC7() bool {
	_, ok := _this.Get_().(D4_DC7)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC8_(_dafny.Zero, false)
}

func (_this D4) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D4_DC8).Cf11
}

func (_this D4) Dtor_cf12() bool {
	return _this.Get_().(D4_DC8).Cf12
}

func (_this D4) Dtor_cf13() bool {
	return _this.Get_().(D4_DC9).Cf13
}

func (_this D4) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf14
}

func (_this D4) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D4_DC9).Cf16
}

func (_this D4) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D4_DC7).Cf10
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC8:
		{
			return "D4.DC8" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC7:
		{
			return "D4.DC7" + "(" + data.Cf10.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC8:
		{
			data2, ok := other.Get_().(D4_DC8)
			return ok && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12
		}
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D4_DC7:
		{
			data2, ok := other.Get_().(D4_DC7)
			return ok && data1.Cf10.Equals(data2.Cf10)
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
	Cf18 _dafny.Int
	Cf19 _dafny.Sequence
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf18 _dafny.Int, Cf19 _dafny.Sequence) D5 {
	return D5{D5_DC11{Cf18, Cf19}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC10 struct {
	Cf17 _dafny.Array
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf17 _dafny.Array) D5 {
	return D5{D5_DC10{Cf17}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC11_(_dafny.Zero, _dafny.EmptySeq)
}

func (_this D5) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D5_DC11).Cf18
}

func (_this D5) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D5_DC11).Cf19
}

func (_this D5) Dtor_cf17() _dafny.Array {
	return _this.Get_().(D5_DC10).Cf17
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf18) + ", " + data.Cf19.VerbatimString(true) + ")"
		}
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19.Equals(data2.Cf19)
		}
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf17 == data2.Cf17
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
	Cf21 bool
	Cf22 bool
	Cf23 _dafny.Sequence
	Cf24 bool
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf21 bool, Cf22 bool, Cf23 _dafny.Sequence, Cf24 bool) D6 {
	return D6{D6_DC13{Cf21, Cf22, Cf23, Cf24}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC12 struct {
	Cf20 _dafny.Array
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf20 _dafny.Array) D6 {
	return D6{D6_DC12{Cf20}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC13_(false, false, _dafny.EmptySeq, false)
}

func (_this D6) Dtor_cf21() bool {
	return _this.Get_().(D6_DC13).Cf21
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D6_DC13).Cf23
}

func (_this D6) Dtor_cf24() bool {
	return _this.Get_().(D6_DC13).Cf24
}

func (_this D6) Dtor_cf20() _dafny.Array {
	return _this.Get_().(D6_DC12).Cf20
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + data.Cf23.VerbatimString(true) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23) && data1.Cf24 == data2.Cf24
		}
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf20 == data2.Cf20
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
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_() D7 {
	return D7{D7_DC15{}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC14 struct {
	Cf25 _dafny.Set
}

func (D7_DC14) isD7() {}

func (CompanionStruct_D7_) Create_DC14_(Cf25 _dafny.Set) D7 {
	return D7{D7_DC14{Cf25}}
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
	return Companion_D7_.Create_DC15_()
}

func (_this D7) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D7_DC14).Cf25
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
			return "D7.DC15"
		}
	case D7_DC14:
		{
			return "D7.DC14" + "(" + _dafny.String(data.Cf25) + ")"
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
			_, ok := other.Get_().(D7_DC15)
			return ok
		}
	case D7_DC14:
		{
			data2, ok := other.Get_().(D7_DC14)
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf28 _dafny.Map
	Cf29 bool
	Cf30 _dafny.Int
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf28 _dafny.Map, Cf29 bool, Cf30 _dafny.Int) D8 {
	return D8{D8_DC18{Cf28, Cf29, Cf30}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

type D8_DC19 struct {
	Cf31 bool
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf31 bool) D8 {
	return D8{D8_DC19{Cf31}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC17 struct {
	Cf27 _dafny.Set
}

func (D8_DC17) isD8() {}

func (CompanionStruct_D8_) Create_DC17_(Cf27 _dafny.Set) D8 {
	return D8{D8_DC17{Cf27}}
}

func (_this D8) Is_DC17() bool {
	_, ok := _this.Get_().(D8_DC17)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC18_(_dafny.EmptyMap, false, _dafny.Zero)
}

func (_this D8) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D8_DC18).Cf28
}

func (_this D8) Dtor_cf29() bool {
	return _this.Get_().(D8_DC18).Cf29
}

func (_this D8) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D8_DC18).Cf30
}

func (_this D8) Dtor_cf31() bool {
	return _this.Get_().(D8_DC19).Cf31
}

func (_this D8) Dtor_cf27() _dafny.Set {
	return _this.Get_().(D8_DC17).Cf27
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ")"
		}
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf31) + ")"
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
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf28.Equals(data2.Cf28) && data1.Cf29 == data2.Cf29 && data1.Cf30.Cmp(data2.Cf30) == 0
		}
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf31 == data2.Cf31
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

type D9_DC21 struct {
	Cf33 _dafny.Int
	Cf34 _dafny.Int
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_(Cf33 _dafny.Int, Cf34 _dafny.Int) D9 {
	return D9{D9_DC21{Cf33, Cf34}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

type D9_DC22 struct {
	Cf35 _dafny.Int
	Cf36 _dafny.Int
	Cf37 bool
	Cf38 _dafny.Int
	Cf39 _dafny.Map
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf35 _dafny.Int, Cf36 _dafny.Int, Cf37 bool, Cf38 _dafny.Int, Cf39 _dafny.Map) D9 {
	return D9{D9_DC22{Cf35, Cf36, Cf37, Cf38, Cf39}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC23 struct {
	Cf40 T1
	Cf41 bool
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf40 T1, Cf41 bool) D9 {
	return D9{D9_DC23{Cf40, Cf41}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC20 struct {
	Cf32 _dafny.MultiSet
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf32 _dafny.MultiSet) D9 {
	return D9{D9_DC20{Cf32}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC21_(_dafny.Zero, _dafny.Zero)
}

func (_this D9) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D9_DC21).Cf33
}

func (_this D9) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D9_DC21).Cf34
}

func (_this D9) Dtor_cf35() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf35
}

func (_this D9) Dtor_cf36() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf36
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC22).Cf37
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf38
}

func (_this D9) Dtor_cf39() _dafny.Map {
	return _this.Get_().(D9_DC22).Cf39
}

func (_this D9) Dtor_cf40() T1 {
	return _this.Get_().(D9_DC23).Cf40
}

func (_this D9) Dtor_cf41() bool {
	return _this.Get_().(D9_DC23).Cf41
}

func (_this D9) Dtor_cf32() _dafny.MultiSet {
	return _this.Get_().(D9_DC20).Cf32
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC21:
		{
			return "D9.DC21" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC21:
		{
			data2, ok := other.Get_().(D9_DC21)
			return ok && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf35.Cmp(data2.Cf35) == 0 && data1.Cf36.Cmp(data2.Cf36) == 0 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39.Equals(data2.Cf39)
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && _dafny.AreEqual(data1.Cf40, data2.Cf40) && data1.Cf41 == data2.Cf41
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D10_DC25 struct {
	Cf43 _dafny.Sequence
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_(Cf43 _dafny.Sequence) D10 {
	return D10{D10_DC25{Cf43}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

type D10_DC24 struct {
	Cf42 _dafny.Map
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf42 _dafny.Map) D10 {
	return D10{D10_DC24{Cf42}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC25_(_dafny.EmptySeq)
}

func (_this D10) Dtor_cf43() _dafny.Sequence {
	return _this.Get_().(D10_DC25).Cf43
}

func (_this D10) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D10_DC24).Cf42
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC25:
		{
			return "D10.DC25" + "(" + data.Cf43.VerbatimString(true) + ")"
		}
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf42) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC25:
		{
			data2, ok := other.Get_().(D10_DC25)
			return ok && data1.Cf43.Equals(data2.Cf43)
		}
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
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

type D11_DC27 struct {
	Cf45 bool
	Cf46 _dafny.Map
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf45 bool, Cf46 _dafny.Map) D11 {
	return D11{D11_DC27{Cf45, Cf46}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC28 struct {
	Cf47 _dafny.Int
	Cf48 _dafny.Array
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf47 _dafny.Int, Cf48 _dafny.Array) D11 {
	return D11{D11_DC28{Cf47, Cf48}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC26 struct {
	Cf44 _dafny.Sequence
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf44 _dafny.Sequence) D11 {
	return D11{D11_DC26{Cf44}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC27_(false, _dafny.EmptyMap)
}

func (_this D11) Dtor_cf45() bool {
	return _this.Get_().(D11_DC27).Cf45
}

func (_this D11) Dtor_cf46() _dafny.Map {
	return _this.Get_().(D11_DC27).Cf46
}

func (_this D11) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D11_DC28).Cf47
}

func (_this D11) Dtor_cf48() _dafny.Array {
	return _this.Get_().(D11_DC28).Cf48
}

func (_this D11) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D11_DC26).Cf44
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf44) + ")"
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
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46.Equals(data2.Cf46)
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf44.Equals(data2.Cf44)
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
	Cf50 _dafny.Int
	Cf51 _dafny.Int
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf50 _dafny.Int, Cf51 _dafny.Int) D12 {
	return D12{D12_DC30{Cf50, Cf51}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

type D12_DC31 struct {
	Cf52 bool
	Cf53 bool
	Cf54 _dafny.Int
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf52 bool, Cf53 bool, Cf54 _dafny.Int) D12 {
	return D12{D12_DC31{Cf52, Cf53, Cf54}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC29 struct {
	Cf49 *C3
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf49 *C3) D12 {
	return D12{D12_DC29{Cf49}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC32 struct {
	Cf55 D12
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf55 D12) D12 {
	return D12{D12_DC32{Cf55}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC30_(_dafny.Zero, _dafny.Zero)
}

func (_this D12) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D12_DC30).Cf50
}

func (_this D12) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D12_DC30).Cf51
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

func (_this D12) Dtor_cf49() *C3 {
	return _this.Get_().(D12_DC29).Cf49
}

func (_this D12) Dtor_cf55() D12 {
	return _this.Get_().(D12_DC32).Cf55
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf55) + ")"
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
			return ok && data1.Cf50.Cmp(data2.Cf50) == 0 && data1.Cf51.Cmp(data2.Cf51) == 0
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf49 == data2.Cf49
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
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

type D13_DC34 struct {
	Cf57 D11
	Cf58 _dafny.Sequence
	Cf59 bool
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf57 D11, Cf58 _dafny.Sequence, Cf59 bool) D13 {
	return D13{D13_DC34{Cf57, Cf58, Cf59}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

type D13_DC33 struct {
	Cf56 _dafny.Array
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf56 _dafny.Array) D13 {
	return D13{D13_DC33{Cf56}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

type D13_DC35 struct {
	Cf60 D13
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf60 D13) D13 {
	return D13{D13_DC35{Cf60}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC34_(Companion_D11_.Default(), _dafny.EmptySeq, false)
}

func (_this D13) Dtor_cf57() D11 {
	return _this.Get_().(D13_DC34).Cf57
}

func (_this D13) Dtor_cf58() _dafny.Sequence {
	return _this.Get_().(D13_DC34).Cf58
}

func (_this D13) Dtor_cf59() bool {
	return _this.Get_().(D13_DC34).Cf59
}

func (_this D13) Dtor_cf56() _dafny.Array {
	return _this.Get_().(D13_DC33).Cf56
}

func (_this D13) Dtor_cf60() D13 {
	return _this.Get_().(D13_DC35).Cf60
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf57) + ", " + data.Cf58.VerbatimString(true) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf56) + ")"
		}
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf57.Equals(data2.Cf57) && data1.Cf58.Equals(data2.Cf58) && data1.Cf59 == data2.Cf59
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf56 == data2.Cf56
		}
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

type D14_DC37 struct {
	Cf62 _dafny.CodePoint
	Cf63 bool
	Cf64 _dafny.Int
	Cf65 T1
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf62 _dafny.CodePoint, Cf63 bool, Cf64 _dafny.Int, Cf65 T1) D14 {
	return D14{D14_DC37{Cf62, Cf63, Cf64, Cf65}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

type D14_DC36 struct {
	Cf61 T2
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf61 T2) D14 {
	return D14{D14_DC36{Cf61}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC37_(_dafny.CodePoint('D'), false, _dafny.Zero, (T1)(nil))
}

func (_this D14) Dtor_cf62() _dafny.CodePoint {
	return _this.Get_().(D14_DC37).Cf62
}

func (_this D14) Dtor_cf63() bool {
	return _this.Get_().(D14_DC37).Cf63
}

func (_this D14) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D14_DC37).Cf64
}

func (_this D14) Dtor_cf65() T1 {
	return _this.Get_().(D14_DC37).Cf65
}

func (_this D14) Dtor_cf61() T2 {
	return _this.Get_().(D14_DC36).Cf61
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC37:
		{
			return "D14.DC37" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC37:
		{
			data2, ok := other.Get_().(D14_DC37)
			return ok && data1.Cf62 == data2.Cf62 && data1.Cf63 == data2.Cf63 && data1.Cf64.Cmp(data2.Cf64) == 0 && _dafny.AreEqual(data1.Cf65, data2.Cf65)
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && _dafny.AreEqual(data1.Cf61, data2.Cf61)
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
	Cf67 bool
	Cf68 _dafny.Int
	Cf69 bool
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf67 bool, Cf68 _dafny.Int, Cf69 bool) D15 {
	return D15{D15_DC39{Cf67, Cf68, Cf69}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC40 struct {
	Cf70 bool
	Cf71 _dafny.Int
}

func (D15_DC40) isD15() {}

func (CompanionStruct_D15_) Create_DC40_(Cf70 bool, Cf71 _dafny.Int) D15 {
	return D15{D15_DC40{Cf70, Cf71}}
}

func (_this D15) Is_DC40() bool {
	_, ok := _this.Get_().(D15_DC40)
	return ok
}

type D15_DC38 struct {
	Cf66 *C2
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf66 *C2) D15 {
	return D15{D15_DC38{Cf66}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC39_(false, _dafny.Zero, false)
}

func (_this D15) Dtor_cf67() bool {
	return _this.Get_().(D15_DC39).Cf67
}

func (_this D15) Dtor_cf68() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf68
}

func (_this D15) Dtor_cf69() bool {
	return _this.Get_().(D15_DC39).Cf69
}

func (_this D15) Dtor_cf70() bool {
	return _this.Get_().(D15_DC40).Cf70
}

func (_this D15) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D15_DC40).Cf71
}

func (_this D15) Dtor_cf66() *C2 {
	return _this.Get_().(D15_DC38).Cf66
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D15_DC40:
		{
			return "D15.DC40" + "(" + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf66) + ")"
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
			return ok && data1.Cf67 == data2.Cf67 && data1.Cf68.Cmp(data2.Cf68) == 0 && data1.Cf69 == data2.Cf69
		}
	case D15_DC40:
		{
			data2, ok := other.Get_().(D15_DC40)
			return ok && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf66 == data2.Cf66
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
	Cf73 _dafny.Sequence
	Cf74 _dafny.Sequence
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf73 _dafny.Sequence, Cf74 _dafny.Sequence) D16 {
	return D16{D16_DC42{Cf73, Cf74}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC41 struct {
	Cf72 _dafny.Sequence
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf72 _dafny.Sequence) D16 {
	return D16{D16_DC41{Cf72}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC42_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D16) Dtor_cf73() _dafny.Sequence {
	return _this.Get_().(D16_DC42).Cf73
}

func (_this D16) Dtor_cf74() _dafny.Sequence {
	return _this.Get_().(D16_DC42).Cf74
}

func (_this D16) Dtor_cf72() _dafny.Sequence {
	return _this.Get_().(D16_DC41).Cf72
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf72) + ")"
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
			return ok && data1.Cf73.Equals(data2.Cf73) && data1.Cf74.Equals(data2.Cf74)
		}
	case D16_DC41:
		{
			data2, ok := other.Get_().(D16_DC41)
			return ok && data1.Cf72.Equals(data2.Cf72)
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
	Fm0(globalState *GlobalState) bool
	Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool
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
	Fm0(globalState *GlobalState) bool
	Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool
	F8() bool
	F8_set_(value bool)
	Fm2(p0 D0, globalState *GlobalState) _dafny.Int
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

// Definition of trait T2
type T2 interface {
	String() string
	M0(p0 _dafny.MultiSet, p1 D0, p2 bool, globalState *GlobalState)
	M1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool)
	F9() bool
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	_f0 _dafny.CodePoint
	_f1 _dafny.Int
	_f2 _dafny.Int
	_f3 _dafny.Int
	_f4 _dafny.Sequence
	_f5 _dafny.Int
	_f6 _dafny.Array
	_f7 _dafny.Int
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this._f0 = _dafny.CodePoint('D')
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.EmptySeq
	_this._f5 = _dafny.Zero
	_this._f6 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f7 = _dafny.Zero
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

func (_this *GlobalState) Ctor__(f0 _dafny.CodePoint, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Int, f4 _dafny.Sequence, f5 _dafny.Int, f6 _dafny.Array, f7 _dafny.Int) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *GlobalState) F0() _dafny.CodePoint {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Int {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Int {
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

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 _dafny.Array
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C0) Ctor__(f12 _dafny.Array) {
	{
		(_this)._f12 = f12
	}
}
func (_this *C0) Fm0(globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(-807)).Cmp(_dafny.IntOfInt64(402)) > 0
	}
}
func (_this *C0) Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eacmud"), _dafny.UnicodeSeqOfUtf8Bytes("h")), (func() _dafny.Sequence {
			if !(false) {
				return _dafny.UnicodeSeqOfUtf8Bytes("gdjemsj")
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}(func(_510_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))
		})())
	}
}
func (_this *C0) Fm10(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	{
		return ((func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(846), _dafny.IntOfInt64(41))); ; {
				_compr_11, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _511_v0 _dafny.Int
				_511_v0 = interface{}(_compr_11).(_dafny.Int)
				if ((_dafny.IntOfInt64(846)).Cmp(_511_v0) <= 0) && ((_511_v0).Cmp(_dafny.IntOfInt64(41)) < 0) {
					_coll11.Add((_511_v0).Minus(_dafny.IntOfInt64(450)), _dafny.IntOfInt64(550))
				}
			}
			return _coll11.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-609), (_dafny.Zero).Minus(_dafny.IntOfInt64(-626))))).Merge((func() _dafny.Map {
			if true {
				return func() _dafny.Map {
					var _coll12 = _dafny.NewMapBuilder()
					_ = _coll12
					for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(942), _dafny.IntOfInt64(845))); ; {
						_compr_12, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _512_v1 _dafny.Int
						_512_v1 = interface{}(_compr_12).(_dafny.Int)
						if ((_dafny.IntOfInt64(942)).Cmp(_512_v1) <= 0) && ((_512_v1).Cmp(_dafny.IntOfInt64(845)) < 0) {
							_coll12.Add(Companion_Default___.SafeDivisionInt(_512_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(275))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
								return func(arg27 _dafny.Int) interface{} {
									return coer27(arg27)
								}
							}(func(_513_i0 _dafny.Int) _dafny.Int {
								return _dafny.IntOfInt64(939)
							}))).Cardinality())), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(528), _dafny.IntOfInt64(570)))).Cardinality())
						}
					}
					return _coll12.ToMap()
				}()
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(299), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(849))).Cardinality()))
		})())
	}
}
func (_this *C0) F12() _dafny.Array {
	{
		return _this._f12
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f8  bool
	_f9  bool
	_f10 _dafny.CodePoint
	_f11 _dafny.Set
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f8 = false
	_this._f9 = false
	_this._f10 = _dafny.CodePoint('D')
	_this._f11 = _dafny.EmptySet
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T2 = &C1{}
var _ T0 = &C1{}
var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F8() bool {
	return _this._f8
}
func (_this *C1) F8_set_(value bool) {
	_this._f8 = value
}
func (_this *C1) F9() bool {
	return _this._f9
}
func (_this *C1) Ctor__(f10 _dafny.CodePoint, f11 _dafny.Set, f9 bool, f8 bool) {
	{
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f9 = f9
		(_this)._f8 = f8
	}
}
func (_this *C1) Fm0(globalState *GlobalState) bool {
	{
		return _this.F8()
	}
}
func (_this *C1) Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool {
	{
		return (func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter15 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("awbrbsnwd"))).Elements()); ; {
				_compr_13, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _514_v0 _dafny.Sequence
				_514_v0 = interface{}(_compr_13).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("awbrbsnwd")), _514_v0) {
					_coll13.Add(_514_v0, (_this).F9())
				}
			}
			return _coll13.ToMap()
		}()).Contains(_dafny.UnicodeSeqOfUtf8Bytes("sapwe"))
	}
}
func (_this *C1) Fm2(p0 D0, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(348)).Plus(_dafny.IntOfInt64(959)), ((_dafny.MultiSetOf((_this).F9())).Cardinality()).Plus((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_515_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-831)
		})))).Cardinality()))
	}
}
func (_this *C1) Fm8(globalState *GlobalState) bool {
	{
		return _this.F8()
	}
}
func (_this *C1) M0(p0 _dafny.MultiSet, p1 D0, p2 bool, globalState *GlobalState) {
	{
		var _516_v0 _dafny.Array
		_ = _516_v0
		var _nw80 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(16))
		_ = _nw80
		_516_v0 = _nw80
		var _517_v1 _dafny.Sequence
		_ = _517_v1
		_517_v1 = _dafny.SeqOf(p2, p2, (_this).F9(), _this.F8(), true)
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_516_v0), 0))
		_ = _index86
		(_516_v0).ArraySet1(_517_v1, (_index86).Int())
		var _518_v2 D2
		_ = _518_v2
		_518_v2 = Companion_D2_.Create_DC4_(false)
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_516_v0), 0))
		_ = _index87
		(_516_v0).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_517_v1, _dafny.SeqOf(_this.F8(), (_518_v2).Dtor_cf5())), (_index87).Int())
		var _519_v3 _dafny.Int
		_ = _519_v3
		_519_v3 = _dafny.IntOfInt64(-928)
		var _520_v4 D0
		_ = _520_v4
		_520_v4 = Companion_D0_.Create_DC1_(_this.F8(), _517_v1, p2)
		var _521_v5 _dafny.Sequence
		_ = _521_v5
		_521_v5 = _dafny.UnicodeSeqOfUtf8Bytes("riwfnhqa")
		var _hi3 _dafny.Int = _dafny.IntOfUint32((_521_v5).Cardinality())
		_ = _hi3
		for _522_i0 := (_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if p2 {
				return _519_v3
			}
			return (_this).Fm2(_520_v4, globalState)
		})())); _522_i0.Cmp(_hi3) < 0; _522_i0 = _522_i0.Plus(_dafny.One) {
			var _523_v6 _dafny.Map
			_ = _523_v6
			_523_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v3, _521_v5)
			var _524_v7 _dafny.Sequence
			_ = _524_v7
			_524_v7 = _dafny.SeqOf(_521_v5)
			var _525_v8 _dafny.Array
			_ = _525_v8
			var _nwElement0_12 _dafny.Sequence = _521_v5
			_ = _nwElement0_12
			var _nw81 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(25))
			_ = _nw81
			(_nw81).ArraySet1(_nwElement0_12, 0)
			(_nw81).ArraySet1(Companion_Default___.Fm9(!(_this.F8()), _this.F8(), globalState), 1)
			(_nw81).ArraySet1(_521_v5, 2)
			(_nw81).ArraySet1(_521_v5, 3)
			(_nw81).ArraySet1(_521_v5, 4)
			(_nw81).ArraySet1(_521_v5, 5)
			(_nw81).ArraySet1(_521_v5, 6)
			(_nw81).ArraySet1(_521_v5, 7)
			(_nw81).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_523_v6).Contains((_this).Fm2(_520_v4, globalState)) {
					return (_523_v6).Get((_this).Fm2(_520_v4, globalState)).(_dafny.Sequence)
				}
				return _521_v5
			})(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(430))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_526_i1 _dafny.Int) _dafny.CodePoint {
				return (_this).F10()
			}))), 8)
			(_nw81).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("fltsycqgp"), 9)
			(_nw81).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_521_v5, _521_v5), 10)
			(_nw81).ArraySet1(_521_v5, 11)
			(_nw81).ArraySet1(_521_v5, 12)
			(_nw81).ArraySet1(_521_v5, 13)
			(_nw81).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("rboifjmne"), 14)
			(_nw81).ArraySet1(_521_v5, 15)
			(_nw81).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_521_v5, _dafny.UnicodeSeqOfUtf8Bytes("mk")), 16)
			(_nw81).ArraySet1(_521_v5, 17)
			(_nw81).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("nvmycqm"), (Companion_Default___.SafeIndex(_522_i0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nvmycqm")).Cardinality()))).Uint32(), (_this).F10()), 18)
			(_nw81).ArraySet1(_521_v5, 19)
			(_nw81).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tjel"), 20)
			(_nw81).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_521_v5, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(793))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_527_i2 _dafny.Int) _dafny.CodePoint {
				return (_this).F10()
			}))), 21)
			(_nw81).ArraySet1((_524_v7).Select((Companion_Default___.SafeIndex((_this).Fm2(_520_v4, globalState), _dafny.IntOfUint32((_524_v7).Cardinality()))).Uint32()).(_dafny.Sequence), 22)
			(_nw81).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(339))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_528_i3 _dafny.Int) _dafny.CodePoint {
				return (_this).F10()
			})), 23)
			(_nw81).ArraySet1(_521_v5, 24)
			_525_v8 = _nw81
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_525_v8), 0))
			_ = _index88
			(_525_v8).ArraySet1(_521_v5, (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_525_v8), 0))
			_ = _index89
			(_525_v8).ArraySet1(_dafny.Companion_Sequence_.Update(_521_v5, (Companion_Default___.SafeIndex(_522_i0, _dafny.IntOfUint32((_521_v5).Cardinality()))).Uint32(), _dafny.CodePoint('n')), (_index89).Int())
			(_this).F8_set_(!(p2))
			(_this).F8_set_(!((_this).F9()))
			(_this).F8_set_((_this).Fm8(globalState))
		}
		var _529_v9 _dafny.Set
		_ = _529_v9
		_529_v9 = _dafny.SetOf(_519_v3)
		(_this).F8_set_((_529_v9).IsProperSubsetOf(_529_v9))
		var _530_i4 _dafny.Int
		_ = _530_i4
		_530_i4 = _dafny.Zero
		{
			for (_519_v3).Cmp(_519_v3) <= 0 {
				{
					if (_530_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_530_i4 = (_530_i4).Plus(_dafny.One)
					var _531_v10 _dafny.Set
					_ = _531_v10
					_531_v10 = _dafny.SetOf(_this.F8(), false)
					var _532_v11 _dafny.Sequence
					_ = _532_v11
					_532_v11 = _dafny.SeqOf(_531_v10)
					var _533_v12 _dafny.Map
					_ = _533_v12
					_533_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("yial"), _521_v5)), _532_v11)
					_533_v12 = (_533_v12).Update(_this.F8(), _532_v11)
					var _534_v13 _dafny.Array
					_ = _534_v13
					var _nw82 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
					_ = _nw82
					_534_v13 = _nw82
					var _535_v14 _dafny.Map
					_ = _535_v14
					_535_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
					var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_534_v13), 0))
					_ = _index90
					(_534_v13).ArraySet1((func() bool {
						if (_535_v14).Contains(false) {
							return (_535_v14).Get(false).(bool)
						}
						return (_this).Fm8(globalState)
					})(), (_index90).Int())
					var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_534_v13), 0))
					_ = _index91
					(_534_v13).ArraySet1((_this).Fm0(globalState), (_index91).Int())
					var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_534_v13), 0))
					_ = _index92
					(_534_v13).ArraySet1((_this).F9(), (_index92).Int())
					(_this).F8_set_(_this.F8())
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		if _this.F8() {
			var _536_v15 _dafny.Array
			_ = _536_v15
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_7
			var _nw83 _dafny.Array
			_ = _nw83
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw83 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = func(_537_i5 _dafny.Int) bool {
					return _this.F8()
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw83 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw83).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw83).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_536_v15 = _nw83
			_536_v15 = _536_v15
			var _538_v16 _dafny.Array
			_ = _538_v16
			var _nwElement0_13 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("atblghtma")
			_ = _nwElement0_13
			var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(8))
			_ = _nw84
			(_nw84).ArraySet1(_nwElement0_13, 0)
			(_nw84).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(133))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg32 _dafny.Int) interface{} {
					return coer32(arg32)
				}
			}(func(_539_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			})), 1)
			(_nw84).ArraySet1(_521_v5, 2)
			(_nw84).ArraySet1(_521_v5, 3)
			(_nw84).ArraySet1(_521_v5, 4)
			(_nw84).ArraySet1(_521_v5, 5)
			(_nw84).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("hvtuvijkn"), 6)
			(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_521_v5, _521_v5), 7)
			_538_v16 = _nw84
			var _540_v17 _dafny.Sequence
			_ = _540_v17
			_540_v17 = _dafny.SeqOf(_521_v5)
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))
			_ = _index93
			(_538_v16).ArraySet1((_540_v17).Select((Companion_Default___.SafeIndex(((_dafny.MultiSetOf(_this.F8(), !(p2), _this.F8())).Update(!(_this.F8()), Companion_Default___.Abs(_519_v3))).Cardinality(), _dafny.IntOfUint32((_540_v17).Cardinality()))).Uint32()).(_dafny.Sequence), (_index93).Int())
			var _541_v18 _dafny.Map
			_ = _541_v18
			_541_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_519_v3, _519_v3)
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))
			_ = _index94
			(_538_v16).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if p2 {
					return (_540_v17).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_541_v18).Contains(_519_v3) {
							return (_541_v18).Get(_519_v3).(_dafny.Int)
						}
						return _519_v3
					})(), _dafny.IntOfUint32((_540_v17).Cardinality()))).Uint32()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_521_v5, _521_v5)
			})(), (Companion_Default___.SafeIndex((_dafny.IntOfInt64(754)).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(429))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg33 _dafny.Int) interface{} {
					return coer33(arg33)
				}
			}(func(_542_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			}))).Cardinality())), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p2 {
					return (_540_v17).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_541_v18).Contains(_519_v3) {
							return (_541_v18).Get(_519_v3).(_dafny.Int)
						}
						return _519_v3
					})(), _dafny.IntOfUint32((_540_v17).Cardinality()))).Uint32()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_521_v5, _521_v5)
			})()).Cardinality()))).Uint32(), (_this).F10()), (_index94).Int())
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))
			_ = _index95
			(_538_v16).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence), (_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence)), (_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence)), (_index95).Int())
			var _543_v19 _dafny.Map
			_ = _543_v19
			_543_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F8())
			var _544_v20 _dafny.MultiSet
			_ = _544_v20
			_544_v20 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()), _519_v3, (p0).Cardinality(), _519_v3)
			_543_v19 = (_543_v19).Update(_this.F8(), !(_544_v20).Contains((_543_v19).Cardinality()))
			if !(!((_this).F9())) || (_this.F8()) {
				(_this).F8_set_((Companion_Default___.SafeDivisionInt(_519_v3, _519_v3)).Cmp((_dafny.IntOfInt64(766)).Plus(_dafny.IntOfUint32(((_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence)).Cardinality()))) > 0)
				var _545_v21 _dafny.Sequence
				_ = _545_v21
				_545_v21 = _dafny.SeqOf(_519_v3, _519_v3, _519_v3)
				var _546_v22 D4
				_ = _546_v22
				_546_v22 = Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("arjmj"))
				var _547_v23 _dafny.Map
				_ = _547_v23
				_547_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_546_v22, _519_v3)
				var _548_v24 _dafny.Array
				_ = _548_v24
				var _nwElement0_14 _dafny.Int = _519_v3
				_ = _nwElement0_14
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(9))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_14, 0)
				(_nw85).ArraySet1((_dafny.Zero).Minus(_519_v3), 1)
				(_nw85).ArraySet1(_dafny.IntOfUint32((_545_v21).Cardinality()), 2)
				(_nw85).ArraySet1((_519_v3).Minus(_dafny.IntOfInt64(898)), 3)
				(_nw85).ArraySet1((_this).Fm2(_520_v4, globalState), 4)
				(_nw85).ArraySet1(Companion_Default___.SafeDivisionInt(_519_v3, _dafny.IntOfUint32((_dafny.SeqOf(_519_v3, (_this).Fm2(_520_v4, globalState))).Cardinality())), 5)
				(_nw85).ArraySet1(_519_v3, 6)
				(_nw85).ArraySet1(_519_v3, 7)
				(_nw85).ArraySet1((_519_v3).Minus((func() _dafny.Int {
					if (_547_v23).Contains(_546_v22) {
						return (_547_v23).Get(_546_v22).(_dafny.Int)
					}
					return (_this).Fm2(Companion_D0_.Create_DC1_(_this.F8(), (_516_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_516_v0), 0))).Int()).(_dafny.Sequence), !(_this.F8())), globalState)
				})()), 8)
				_548_v24 = _nw85
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_548_v24), 0))
				_ = _index96
				(_548_v24).ArraySet1(_519_v3, (_index96).Int())
				var _549_v25 _dafny.Set
				_ = _549_v25
				_549_v25 = _dafny.SetOf(_this.F8())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_548_v24), 0))
				_ = _index97
				(_548_v24).ArraySet1((_549_v25).Cardinality(), (_index97).Int())
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_536_v15), 0))
				_ = _index98
				(_536_v15).ArraySet1(true, (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_536_v15), 0))
				_ = _index99
				(_536_v15).ArraySet1((_this).F9(), (_index99).Int())
				var _550_v26 _dafny.Sequence
				_ = _550_v26
				_550_v26 = _dafny.SeqOf(_545_v21, _545_v21)
				_550_v26 = _550_v26
				var _551_v27 _dafny.Array
				_ = _551_v27
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(4))
				_ = _nw86
				_551_v27 = _nw86
				var _552_v28 _dafny.Map
				_ = _552_v28
				_552_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_548_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_548_v24), 0))).Int()).(_dafny.Int))
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_551_v27), 0))
				_ = _index100
				(_551_v27).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _552_v28), (_index100).Int())
				var _553_v30 _dafny.Map
				_ = _553_v30
				_553_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _552_v28)
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_551_v27), 0))
				_ = _index101
				(_551_v27).ArraySet1((func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter16 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg34 _dafny.Int) interface{} {
							return coer34(arg34)
						}
					}(func(_554_i8 _dafny.Int) _dafny.CodePoint {
						return (_this).F10()
					}))).Elements()); ; {
						_compr_14, _ok16 := _iter16()
						if !_ok16 {
							break
						}
						var _555_v29 _dafny.CodePoint
						_555_v29 = interface{}(_compr_14).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-42))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg35 _dafny.Int) interface{} {
								return coer35(arg35)
							}
						}(func(_554_i8 _dafny.Int) _dafny.CodePoint {
							return (_this).F10()
						})), _555_v29) {
							_coll14.Add(_555_v29, _552_v28)
						}
					}
					return _coll14.ToMap()
				}()).Merge(_553_v30), (_index101).Int())
			} else {
				var _556_v31 _dafny.Map
				_ = _556_v31
				_556_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_521_v5, _519_v3)
				_556_v31 = (_556_v31).Update((_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence), _519_v3)
				_518_v2 = _518_v2
				var _557_v32 _dafny.Array
				_ = _557_v32
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw87
				_557_v32 = _nw87
				var _558_v33 *C0
				_ = _558_v33
				var _nw88 *C0 = New_C0_()
				_ = _nw88
				_nw88.Ctor__(_557_v32)
				_558_v33 = _nw88
				var _559_v34 _dafny.Array
				_ = _559_v34
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_8
				var _nw89 _dafny.Array
				_ = _nw89
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw89 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.CodePoint = func(_560_i9 _dafny.Int) _dafny.CodePoint {
						return (_this).F10()
					}
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw89 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw89).ArraySet1CodePoint(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw89).ArraySet1CodePoint(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_559_v34 = _nw89
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_559_v34), 0))
				_ = _index102
				(_559_v34).ArraySet1CodePoint((_this).F10(), (_index102).Int())
				var _561_v35 _dafny.Map
				_ = _561_v35
				_561_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_541_v18).Cardinality(), _521_v5)
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_559_v34), 0))
				_ = _index103
				var _rhs91 _dafny.CodePoint = (_this).F10()
				_ = _rhs91
				var _rhs92 _dafny.Array = _536_v15
				_ = _rhs92
				var _rhs93 bool = (_519_v3).Cmp(_519_v3) > 0
				_ = _rhs93
				var _rhs94 bool = (_558_v33).Fm1(_dafny.Companion_Sequence_.Concatenate((_538_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(788), _dafny.ArrayLen((_538_v16), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_561_v35).Contains((_dafny.Zero).Minus(_519_v3)) {
						return (_561_v35).Get((_dafny.Zero).Minus(_519_v3)).(_dafny.Sequence)
					}
					return (_540_v17).Select((Companion_Default___.SafeIndex(_519_v3, _dafny.IntOfUint32((_540_v17).Cardinality()))).Uint32()).(_dafny.Sequence)
				})(), (Companion_Default___.SafeIndex((_this).Fm2(Companion_D0_.Create_DC1_(!(!(p2)), (_516_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_516_v0), 0))).Int()).(_dafny.Sequence), _this.F8()), globalState), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_561_v35).Contains((_dafny.Zero).Minus(_519_v3)) {
						return (_561_v35).Get((_dafny.Zero).Minus(_519_v3)).(_dafny.Sequence)
					}
					return (_540_v17).Select((Companion_Default___.SafeIndex(_519_v3, _dafny.IntOfUint32((_540_v17).Cardinality()))).Uint32()).(_dafny.Sequence)
				})()).Cardinality()))).Uint32(), (_this).F10())), _520_v4, globalState)
				_ = _rhs94
				var _lhs52 _dafny.Array = _559_v34
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(42), _dafny.ArrayLen((_559_v34), 0))
				_ = _lhs53
				var _lhs54 *C1 = _this
				_ = _lhs54
				var _lhs55 *C1 = _this
				_ = _lhs55
				(_lhs52).ArraySet1CodePoint(_rhs91, (_lhs53).Int())
				_536_v15 = _rhs92
				_lhs54.F8_set_(_rhs93)
				_lhs55.F8_set_(_rhs94)
				(_this).F8_set_(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf(_this.F8(), (_this).F9()), (_516_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(467), _dafny.ArrayLen((_516_v0), 0))).Int()).(_dafny.Sequence)))
			}
		} else {
			var _562_v36 _dafny.Array
			_ = _562_v36
			var _nwElement0_15 _dafny.Int = (_519_v3).Plus(_519_v3)
			_ = _nwElement0_15
			var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(18))
			_ = _nw90
			(_nw90).ArraySet1(_nwElement0_15, 0)
			(_nw90).ArraySet1(_519_v3, 1)
			(_nw90).ArraySet1(_519_v3, 2)
			(_nw90).ArraySet1(_519_v3, 3)
			(_nw90).ArraySet1(_519_v3, 4)
			(_nw90).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(330), _519_v3), 5)
			(_nw90).ArraySet1(_dafny.IntOfInt64(408), 6)
			(_nw90).ArraySet1(_519_v3, 7)
			(_nw90).ArraySet1((_this).Fm2(Companion_D0_.Create_DC1_(false, _517_v1, (_this).F9()), globalState), 8)
			(_nw90).ArraySet1(_519_v3, 9)
			(_nw90).ArraySet1(_519_v3, 10)
			(_nw90).ArraySet1((_dafny.IntOfInt64(392)).Minus(_519_v3), 11)
			(_nw90).ArraySet1(_519_v3, 12)
			(_nw90).ArraySet1(_519_v3, 13)
			(_nw90).ArraySet1(_dafny.IntOfUint32((_521_v5).Cardinality()), 14)
			(_nw90).ArraySet1(Companion_Default___.SafeDivisionInt(_519_v3, _519_v3), 15)
			(_nw90).ArraySet1(_519_v3, 16)
			(_nw90).ArraySet1(_519_v3, 17)
			_562_v36 = _nw90
			_562_v36 = _562_v36
			_519_v3 = _519_v3
			var _563_v37 D5
			_ = _563_v37
			_563_v37 = Companion_D5_.Create_DC10_(_562_v36)
			var _564_v38 _dafny.Sequence
			_ = _564_v38
			_564_v38 = _dafny.SeqOf(_562_v36)
			var _565_v39 _dafny.Array
			_ = _565_v39
			var _nwElement0_16 _dafny.Array = _562_v36
			_ = _nwElement0_16
			var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(19))
			_ = _nw91
			(_nw91).ArraySet1(_nwElement0_16, 0)
			(_nw91).ArraySet1(_562_v36, 1)
			(_nw91).ArraySet1((_563_v37).Dtor_cf17(), 2)
			(_nw91).ArraySet1(_562_v36, 3)
			(_nw91).ArraySet1(_562_v36, 4)
			(_nw91).ArraySet1(_562_v36, 5)
			(_nw91).ArraySet1(_562_v36, 6)
			(_nw91).ArraySet1(_562_v36, 7)
			(_nw91).ArraySet1(_562_v36, 8)
			(_nw91).ArraySet1(_562_v36, 9)
			(_nw91).ArraySet1(_562_v36, 10)
			(_nw91).ArraySet1(_562_v36, 11)
			(_nw91).ArraySet1(_562_v36, 12)
			(_nw91).ArraySet1(_562_v36, 13)
			(_nw91).ArraySet1(_562_v36, 14)
			(_nw91).ArraySet1((_564_v38).Select((Companion_Default___.SafeIndex(_519_v3, _dafny.IntOfUint32((_564_v38).Cardinality()))).Uint32()).(_dafny.Array), 15)
			(_nw91).ArraySet1(_562_v36, 16)
			(_nw91).ArraySet1(_562_v36, 17)
			(_nw91).ArraySet1(_562_v36, 18)
			_565_v39 = _nw91
			var _566_v40 *C0
			_ = _566_v40
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__(_565_v39)
			_566_v40 = _nw92
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_562_v36), 0))
			_ = _index104
			(_562_v36).ArraySet1(_519_v3, (_index104).Int())
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_562_v36), 0))
			_ = _index105
			var _rhs95 _dafny.Int = (((_this).Fm2(_520_v4, globalState)).Plus(_519_v3)).Plus(Companion_Default___.SafeDivisionInt(_519_v3, _519_v3))
			_ = _rhs95
			var _rhs96 _dafny.Int = _519_v3
			_ = _rhs96
			var _rhs97 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gucbasaxj")).Cardinality()), _519_v3)
			_ = _rhs97
			var _rhs98 _dafny.Sequence = _521_v5
			_ = _rhs98
			var _lhs56 _dafny.Array = _562_v36
			_ = _lhs56
			var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(49), _dafny.ArrayLen((_562_v36), 0))
			_ = _lhs57
			(_lhs56).ArraySet1(_rhs95, (_lhs57).Int())
			_519_v3 = _rhs96
			_519_v3 = _rhs97
			_521_v5 = _rhs98
			var _567_v41 *C0
			_ = _567_v41
			var _nw93 *C0 = New_C0_()
			_ = _nw93
			_nw93.Ctor__((_566_v40).F12())
			_567_v41 = _nw93
			_567_v41 = _566_v40
		}
		var _568_v42 D3
		_ = _568_v42
		_568_v42 = Companion_D3_.Create_DC6_((_this).F9(), _519_v3, _dafny.IntOfUint32((_521_v5).Cardinality()))
		var _569_v43 _dafny.Map
		_ = _569_v43
		_569_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _568_v42)
		var _570_v44 _dafny.MultiSet
		_ = _570_v44
		_570_v44 = _dafny.MultiSetOf(_519_v3)
		_569_v43 = (_569_v43).Update((func() bool {
			if p2 {
				return p2
			}
			return _this.F8()
		})(), Companion_D3_.Create_DC6_((_this).F9(), ((_570_v44).Update(_519_v3, Companion_Default___.Abs(_519_v3))).Cardinality(), _dafny.IntOfInt64(307)))
	}
}
func (_this *C1) M1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _571_v0 _dafny.Array
		_ = _571_v0
		var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(17))
		_ = _nw94
		_571_v0 = _nw94
		var _572_v1 _dafny.Int
		_ = _572_v1
		_572_v1 = _dafny.IntOfInt64(802)
		var _573_v2 _dafny.Set
		_ = _573_v2
		_573_v2 = _dafny.SetOf(_572_v1)
		var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_571_v0), 0))
		_ = _index106
		(_571_v0).ArraySet1(_573_v2, (_index106).Int())
		var _574_v3 _dafny.Set
		_ = _574_v3
		_574_v3 = _dafny.SetOf((_this).F10(), (_this).F10())
		var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(524), _dafny.ArrayLen((_571_v0), 0))
		_ = _index107
		(_571_v0).ArraySet1(_dafny.SetOf((_572_v1).Plus((_574_v3).Cardinality())), (_index107).Int())
		var _hi4 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(337), (_dafny.Zero).Minus(_572_v1))).Cardinality())
		_ = _hi4
		for _575_i0 := _572_v1; _575_i0.Cmp(_hi4) < 0; _575_i0 = _575_i0.Plus(_dafny.One) {
			if (_this).Fm0(globalState) {
				var _576_v4 D5
				_ = _576_v4
				_576_v4 = Companion_D5_.Create_DC11_(_575_i0, p0)
				var _577_v5 _dafny.Map
				_ = _577_v5
				_577_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F10())
				_576_v4 = Companion_Default___.Fm11((_577_v5).Merge(_577_v5), (_this).F9(), !((func() bool {
					if false {
						return p1
					}
					return _this.F8()
				})()), _this.F8(), globalState)
				var _578_v6 _dafny.Array
				_ = _578_v6
				var _nw95 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
				_ = _nw95
				_578_v6 = _nw95
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))
				_ = _index108
				(_578_v6).ArraySet1((_this).F9(), (_index108).Int())
				var _579_v7 _dafny.Sequence
				_ = _579_v7
				_579_v7 = _dafny.SeqOf(_577_v5)
				var _580_v8 _dafny.MultiSet
				_ = _580_v8
				_580_v8 = _dafny.MultiSetOf(p1)
				var _581_v9 _dafny.Array
				_ = _581_v9
				var _nwElement0_17 _dafny.Int = _575_i0
				_ = _nwElement0_17
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.One)
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_17, 0)
				_581_v9 = _nw96
				var _582_v10 _dafny.MultiSet
				_ = _582_v10
				_582_v10 = _dafny.MultiSetOf(_581_v9, _581_v9)
				var _583_v11 D5
				_ = _583_v11
				_583_v11 = Companion_D5_.Create_DC10_(_581_v9)
				var _584_v12 _dafny.Sequence
				_ = _584_v12
				_584_v12 = _dafny.SeqOf(_582_v10, _dafny.MultiSetOf(_581_v9, _581_v9, (_583_v11).Dtor_cf17(), _581_v9, _581_v9), _582_v10)
				var _585_v13 _dafny.Sequence
				_ = _585_v13
				_585_v13 = _dafny.SeqOf(p1)
				var _586_v14 _dafny.Sequence
				_ = _586_v14
				_586_v14 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_585_v13, (Companion_Default___.SafeIndex(_575_i0, _dafny.IntOfUint32((_585_v13).Cardinality()))).Uint32(), p1)).Cardinality()), _572_v1, _575_i0)
				var _587_v15 _dafny.Map
				_ = _587_v15
				_587_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_575_i0, p1)
				var _588_v16 D0
				_ = _588_v16
				_588_v16 = Companion_D0_.Create_DC1_(p1, _585_v13, _this.F8())
				var _589_v17 _dafny.MultiSet
				_ = _589_v17
				_589_v17 = _dafny.MultiSetOf((_this).F10())
				var _590_v18 _dafny.Array
				_ = _590_v18
				var _nwElement0_18 _dafny.Int = ((_579_v7).Select((Companion_Default___.SafeIndex(_572_v1, _dafny.IntOfUint32((_579_v7).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()
				_ = _nwElement0_18
				var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(29))
				_ = _nw97
				(_nw97).ArraySet1(_nwElement0_18, 0)
				(_nw97).ArraySet1((_575_i0).Plus(_572_v1), 1)
				(_nw97).ArraySet1((func() _dafny.Int {
					if (_580_v8).Contains(p1) {
						return (_580_v8).Multiplicity(p1)
					}
					return _572_v1
				})(), 2)
				(_nw97).ArraySet1(_572_v1, 3)
				(_nw97).ArraySet1(_575_i0, 4)
				(_nw97).ArraySet1(_575_i0, 5)
				(_nw97).ArraySet1(_572_v1, 6)
				(_nw97).ArraySet1(_572_v1, 7)
				(_nw97).ArraySet1(_572_v1, 8)
				(_nw97).ArraySet1(_575_i0, 9)
				(_nw97).ArraySet1(_575_i0, 10)
				(_nw97).ArraySet1(_572_v1, 11)
				(_nw97).ArraySet1(_575_i0, 12)
				(_nw97).ArraySet1(((_584_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_586_v14).Cardinality()), _dafny.IntOfUint32((_584_v12).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 13)
				(_nw97).ArraySet1(_572_v1, 14)
				(_nw97).ArraySet1(_572_v1, 15)
				(_nw97).ArraySet1(_dafny.IntOfInt64(-560), 16)
				(_nw97).ArraySet1(Companion_Default___.SafeDivisionInt((_587_v15).Cardinality(), _575_i0), 17)
				(_nw97).ArraySet1(_575_i0, 18)
				(_nw97).ArraySet1(_dafny.IntOfUint32((_586_v14).Cardinality()), 19)
				(_nw97).ArraySet1(_575_i0, 20)
				(_nw97).ArraySet1(_572_v1, 21)
				(_nw97).ArraySet1(_575_i0, 22)
				(_nw97).ArraySet1(_572_v1, 23)
				(_nw97).ArraySet1((_dafny.Zero).Minus((_572_v1).Plus(_572_v1)), 24)
				(_nw97).ArraySet1(Companion_Default___.SafeModuloInt(_575_i0, _575_i0), 25)
				(_nw97).ArraySet1((_this).Fm2(func(_pat_let25_0 D0) D0 {
					return func(_591_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let26_0 bool) D0 {
							return func(_592_dt__update_hcf0_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_592_dt__update_hcf0_h0, (_591_dt__update__tmp_h0).Dtor_cf1(), (_591_dt__update__tmp_h0).Dtor_cf2())
							}(_pat_let26_0)
						}((_this).F9())
					}(_pat_let25_0)
				}(_588_v16), globalState), 26)
				(_nw97).ArraySet1((_575_i0).Minus((func() _dafny.Int {
					if (_589_v17).Contains((_this).F10()) {
						return (_589_v17).Multiplicity((_this).F10())
					}
					return _575_i0
				})()), 27)
				(_nw97).ArraySet1(_dafny.IntOfInt64(-577), 28)
				_590_v18 = _nw97
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_581_v9), 0))
				_ = _index109
				(_581_v9).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_586_v14, _586_v14)).Cardinality()), (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))
				_ = _index110
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_581_v9), 0))
				_ = _index111
				var _rhs99 _dafny.Sequence = Companion_Default___.Fm9((_575_i0).Cmp(_575_i0) == 0, true, globalState)
				_ = _rhs99
				var _rhs100 bool = p1
				_ = _rhs100
				var _rhs101 _dafny.Int = (_575_i0).Minus(_575_i0)
				_ = _rhs101
				var _lhs58 _dafny.Array = _578_v6
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))
				_ = _lhs59
				var _lhs60 _dafny.Array = _581_v9
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_581_v9), 0))
				_ = _lhs61
				r0 = _rhs99
				(_lhs58).ArraySet1(_rhs100, (_lhs59).Int())
				(_lhs60).ArraySet1(_rhs101, (_lhs61).Int())
				var _593_v19 D4
				_ = _593_v19
				_593_v19 = Companion_D4_.Create_DC8_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}(func(_594_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))).Cardinality()), (_this).F9())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))
				_ = _index112
				(_578_v6).ArraySet1((_593_v19).Dtor_cf12(), (_index112).Int())
				var _595_v20 _dafny.Sequence
				_ = _595_v20
				_595_v20 = _dafny.SeqOf(p0, p0, p0)
				var _596_v21 _dafny.Map
				_ = _596_v21
				_596_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_595_v20, (_578_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))).Int()).(bool))
				var _597_v22 D3
				_ = _597_v22
				_597_v22 = Companion_D3_.Create_DC5_(_586_v14)
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_578_v6), 0))
				_ = _index113
				(_578_v6).ArraySet1((func() bool {
					if (_596_v21).Contains(Companion_Default___.Fm12(p1, _597_v22, globalState)) {
						return (_596_v21).Get(Companion_Default___.Fm12(p1, _597_v22, globalState)).(bool)
					}
					return (_585_v13).Select((Companion_Default___.SafeIndex((_this).Fm2(_588_v16, globalState), _dafny.IntOfUint32((_585_v13).Cardinality()))).Uint32()).(bool)
				})(), (_index113).Int())
				r2 = ((_581_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_581_v9), 0))).Int()).(_dafny.Int)).Cmp((_581_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(872), _dafny.ArrayLen((_581_v9), 0))).Int()).(_dafny.Int)) > 0
			} else {
				var _598_v23 _dafny.MultiSet
				_ = _598_v23
				_598_v23 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(p0, p0), _dafny.Companion_Sequence_.Concatenate(p0, p0))
				_598_v23 = (_598_v23).Difference((func() _dafny.MultiSet {
					if p1 {
						return _dafny.MultiSetOf(p0, p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(376))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg37 _dafny.Int) interface{} {
								return coer37(arg37)
							}
						}(func(_599_i2 _dafny.Int) _dafny.CodePoint {
							return (_this).F10()
						})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(989))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}(func(_600_i3 _dafny.Int) _dafny.CodePoint {
							return (_this).F10()
						})))
					}
					return _dafny.MultiSetOf(p0, p0, p0)
				})())
				var _601_v24 _dafny.Array
				_ = _601_v24
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw98
				_601_v24 = _nw98
				var _602_v25 *C0
				_ = _602_v25
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__(_601_v24)
				_602_v25 = _nw99
				var _603_v26 _dafny.Sequence
				_ = _603_v26
				_603_v26 = _dafny.SeqOf(_602_v25, _602_v25, _602_v25, _602_v25, _602_v25)
				var _604_v27 _dafny.Map
				_ = _604_v27
				_604_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(-991))
				var _605_v28 _dafny.Sequence
				_ = _605_v28
				_605_v28 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), _572_v1), _604_v27)
				var _606_v29 _dafny.Sequence
				_ = _606_v29
				_606_v29 = _dafny.SeqOf(((_dafny.Zero).Minus(_dafny.IntOfUint32((_603_v26).Cardinality()))).Times(_572_v1), _575_i0, Companion_Default___.SafeModuloInt(_575_i0, _572_v1), _575_i0, (_575_i0).Times(((_605_v28).Select((Companion_Default___.SafeIndex(_575_i0, _dafny.IntOfUint32((_605_v28).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()))
				var _rhs102 _dafny.Int = _dafny.IntOfUint32((_606_v29).Cardinality())
				_ = _rhs102
				var _rhs103 bool = p1
				_ = _rhs103
				_572_v1 = _rhs102
				r2 = _rhs103
				var _607_v30 *C0
				_ = _607_v30
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(_601_v24)
				_607_v30 = _nw100
				var _608_v31 _dafny.MultiSet
				_ = _608_v31
				_608_v31 = _dafny.MultiSetOf(_572_v1)
				_608_v31 = (_608_v31).Intersection(_608_v31)
				var _609_v32 _dafny.Array
				_ = _609_v32
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw101
				_609_v32 = _nw101
				var _610_v33 _dafny.Map
				_ = _610_v33
				_610_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _609_v32)
				_610_v33 = (_610_v33).Update(_dafny.CodePoint('h'), _609_v32)
			}
			var _611_v34 _dafny.Map
			_ = _611_v34
			_611_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _572_v1)
			var _612_v35 _dafny.Sequence
			_ = _612_v35
			_612_v35 = _dafny.SeqOf(_611_v34, _611_v34, Companion_Default___.Fm13(globalState))
			_612_v35 = _612_v35
			var _613_v36 _dafny.Array
			_ = _613_v36
			var _nw102 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw102
			_613_v36 = _nw102
			var _614_v37 _dafny.Array
			_ = _614_v37
			var _nwElement0_19 _dafny.Array = _613_v36
			_ = _nwElement0_19
			var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(4))
			_ = _nw103
			(_nw103).ArraySet1(_nwElement0_19, 0)
			(_nw103).ArraySet1(_613_v36, 1)
			(_nw103).ArraySet1(_613_v36, 2)
			(_nw103).ArraySet1(_613_v36, 3)
			_614_v37 = _nw103
			var _615_v38 D6
			_ = _615_v38
			_615_v38 = Companion_D6_.Create_DC12_(_614_v37)
			var _616_v39 *C0
			_ = _616_v39
			var _nw104 *C0 = New_C0_()
			_ = _nw104
			_nw104.Ctor__((_615_v38).Dtor_cf20())
			_616_v39 = _nw104
			var _rhs104 bool = !(!(_this.F8()))
			_ = _rhs104
			var _rhs105 bool = p1
			_ = _rhs105
			var _lhs62 *C1 = _this
			_ = _lhs62
			var _lhs63 *C1 = _this
			_ = _lhs63
			_lhs62.F8_set_(_rhs104)
			_lhs63.F8_set_(_rhs105)
		}
		var _617_v40 _dafny.Array
		_ = _617_v40
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_9
		var _nw105 _dafny.Array
		_ = _nw105
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw105 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = (func(_618_v1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_619_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate((Companion_D3_.Create_DC5_(_dafny.SeqOf(_618_v1, (_dafny.SetOf(_this.F8())).Cardinality()))).Dtor_cf6(), (Companion_D3_.Create_DC5_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(869))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg39 _dafny.Int) interface{} {
							return coer39(arg39)
						}
					}((func(_620_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_621_i5 _dafny.Int) _dafny.Int {
							return _620_v1
						}
					})(_618_v1))))).Dtor_cf6())
				}
			})(_572_v1)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw105 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw105).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw105).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_617_v40 = _nw105
		var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_617_v40), 0))
		_ = _index114
		(_617_v40).ArraySet1(Companion_Default___.Fm14(_572_v1, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}(func(_622_i6 _dafny.Int) _dafny.CodePoint {
			return (_this).F10()
		})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}(func(_623_i6 _dafny.Int) _dafny.CodePoint {
			return (_this).F10()
		}))).Cardinality()))).Uint32(), (_this).F10()), _this.F8(), globalState), (_index114).Int())
		var _624_v41 _dafny.MultiSet
		_ = _624_v41
		_624_v41 = _dafny.MultiSetOf(false)
		var _625_v42 _dafny.Map
		_ = _625_v42
		_625_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_572_v1, _624_v41)
		var _626_v43 _dafny.Array
		_ = _626_v43
		var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw106
		_626_v43 = _nw106
		var _627_v44 D5
		_ = _627_v44
		_627_v44 = Companion_D5_.Create_DC10_(_626_v43)
		var _628_v45 _dafny.MultiSet
		_ = _628_v45
		_628_v45 = _dafny.MultiSetOf(_627_v44, Companion_D5_.Create_DC10_(_626_v43))
		var _629_v46 _dafny.Sequence
		_ = _629_v46
		_629_v46 = _dafny.SeqOf((_572_v1).Plus(_dafny.IntOfInt64(235)), _572_v1, _dafny.IntOfInt64(-143), (_572_v1).Minus(((func() _dafny.MultiSet {
			if (_625_v42).Contains((_628_v45).Cardinality()) {
				return (_625_v42).Get((_628_v45).Cardinality()).(_dafny.MultiSet)
			}
			return _624_v41
		})()).Cardinality()), _dafny.IntOfInt64(388))
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_617_v40), 0))
		_ = _index115
		(_617_v40).ArraySet1(_629_v46, (_index115).Int())
		var _630_v47 _dafny.Array
		_ = _630_v47
		var _nwElement0_20 bool = p1
		_ = _nwElement0_20
		var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(11))
		_ = _nw107
		(_nw107).ArraySet1(_nwElement0_20, 0)
		(_nw107).ArraySet1(p1, 1)
		(_nw107).ArraySet1(_this.F8(), 2)
		(_nw107).ArraySet1(false, 3)
		(_nw107).ArraySet1(p1, 4)
		(_nw107).ArraySet1(_this.F8(), 5)
		(_nw107).ArraySet1(_this.F8(), 6)
		(_nw107).ArraySet1(_this.F8(), 7)
		(_nw107).ArraySet1(!_dafny.Companion_Sequence_.Equal((_617_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_617_v40), 0))).Int()).(_dafny.Sequence), (_617_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_617_v40), 0))).Int()).(_dafny.Sequence)), 8)
		(_nw107).ArraySet1((true) == ((_this).F9()), 9)
		(_nw107).ArraySet1(p1, 10)
		_630_v47 = _nw107
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))
		_ = _index116
		(_630_v47).ArraySet1(!_dafny.Companion_Sequence_.Contains((_617_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_617_v40), 0))).Int()).(_dafny.Sequence), _572_v1), (_index116).Int())
		var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))
		_ = _index117
		(_630_v47).ArraySet1(_dafny.Companion_Sequence_.Contains(p0, (_this).F10()), (_index117).Int())
		var _hi5 _dafny.Int = _572_v1
		_ = _hi5
		for _631_i7 := _572_v1; _631_i7.Cmp(_hi5) < 0; _631_i7 = _631_i7.Plus(_dafny.One) {
			var _632_v48 _dafny.Sequence
			_ = _632_v48
			_632_v48 = _dafny.SeqOf((_630_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))).Int()).(bool), p1, !((_630_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))).Int()).(bool)), (_630_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))).Int()).(bool), (_630_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))).Int()).(bool))
			var _633_v49 _dafny.Map
			_ = _633_v49
			_633_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_632_v48).Cardinality())).Cmp(_572_v1) <= 0, _631_i7)
			var _634_v50 D0
			_ = _634_v50
			_634_v50 = Companion_D0_.Create_DC1_((_this).F9(), _632_v48, true)
			_633_v49 = (_633_v49).Update((_this).F9(), Companion_Default___.SafeDivisionInt((_this).Fm2(_634_v50, globalState), _dafny.IntOfInt64(59)))
			var _635_v51 _dafny.Array
			_ = _635_v51
			var _nwElement0_21 _dafny.Array = _626_v43
			_ = _nwElement0_21
			var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(20))
			_ = _nw108
			(_nw108).ArraySet1(_nwElement0_21, 0)
			(_nw108).ArraySet1(_626_v43, 1)
			(_nw108).ArraySet1(_626_v43, 2)
			(_nw108).ArraySet1(_626_v43, 3)
			(_nw108).ArraySet1(_626_v43, 4)
			(_nw108).ArraySet1(_626_v43, 5)
			(_nw108).ArraySet1(_626_v43, 6)
			(_nw108).ArraySet1(_626_v43, 7)
			(_nw108).ArraySet1((_627_v44).Dtor_cf17(), 8)
			(_nw108).ArraySet1(_626_v43, 9)
			(_nw108).ArraySet1(_626_v43, 10)
			(_nw108).ArraySet1((_627_v44).Dtor_cf17(), 11)
			(_nw108).ArraySet1(_626_v43, 12)
			(_nw108).ArraySet1(_626_v43, 13)
			(_nw108).ArraySet1(_626_v43, 14)
			(_nw108).ArraySet1(_626_v43, 15)
			(_nw108).ArraySet1(_626_v43, 16)
			(_nw108).ArraySet1(_626_v43, 17)
			(_nw108).ArraySet1(_626_v43, 18)
			(_nw108).ArraySet1(_626_v43, 19)
			_635_v51 = _nw108
			var _636_v52 *C0
			_ = _636_v52
			var _nw109 *C0 = New_C0_()
			_ = _nw109
			_nw109.Ctor__(_635_v51)
			_636_v52 = _nw109
			var _637_v53 _dafny.Array
			_ = _637_v53
			var _nw110 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
			_ = _nw110
			_637_v53 = _nw110
			var _638_v54 _dafny.Map
			_ = _638_v54
			_638_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_631_i7, _this.F8())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_637_v53), 0))
			_ = _index118
			(_637_v53).ArraySet1(_638_v54, (_index118).Int())
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(167), _dafny.ArrayLen((_637_v53), 0))
			_ = _index119
			(_637_v53).ArraySet1(func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(796), _dafny.IntOfInt64(249))); ; {
					_compr_15, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _639_v55 _dafny.Int
					_639_v55 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(796)).Cmp(_639_v55) <= 0) && ((_639_v55).Cmp(_dafny.IntOfInt64(249)) < 0) {
						_coll15.Add((_639_v55).Times(_572_v1), (_this).F9())
					}
				}
				return _coll15.ToMap()
			}(), (_index119).Int())
			r2 = _this.F8()
		}
		_572_v1 = _dafny.IntOfInt64(421)
		r0 = _dafny.Companion_Sequence_.Concatenate(p0, p0)
		r1 = _572_v1
		r2 = (_630_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_630_v47), 0))).Int()).(bool)
		return r0, r1, r2
	}
}
func (_this *C1) F10() _dafny.CodePoint {
	{
		return _this._f10
	}
}
func (_this *C1) F11() _dafny.Set {
	{
		return _this._f11
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f8 bool
	_f9 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f8 = false
	_this._f9 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C2{}
var _ T2 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F8() bool {
	return _this._f8
}
func (_this *C2) F8_set_(value bool) {
	_this._f8 = value
}
func (_this *C2) F9() bool {
	return _this._f9
}
func (_this *C2) Ctor__(f8 bool, f9 bool) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C2) Fm2(p0 D0, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((((func() D3 {
			if _this.F8() {
				return Companion_D3_.Create_DC5_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("efhmg")).Cardinality()), (_dafny.SetOf(!((_this).F9()), _this.F8())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jpesa")).Cardinality())))
			}
			return Companion_D3_.Create_DC5_(_dafny.SeqOf(_dafny.IntOfInt64(989), _dafny.IntOfInt64(490)))
		})()).Dtor_cf6()).Cardinality())
	}
}
func (_this *C2) Fm0(globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(775)).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-788), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), _dafny.IntOfInt64(718))).Cardinality())), _this.F8())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(319))).Cardinality()))).Cardinality()))) != 0
	}
}
func (_this *C2) Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool {
	{
		return (_this).F9()
	}
}
func (_this *C2) Fm5(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf((_this).F9(), (_this).F9(), (_this).F9(), (_this).F9())
	}
}
func (_this *C2) Fm6(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(470), _dafny.IntOfInt64(464))), _dafny.IntOfInt64(-7))
	}
}
func (_this *C2) M0(p0 _dafny.MultiSet, p1 D0, p2 bool, globalState *GlobalState) {
	{
		var _640_v0 _dafny.Int
		_ = _640_v0
		_640_v0 = _dafny.IntOfInt64(195)
		var _641_v1 _dafny.Map
		_ = _641_v1
		_641_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(194), Companion_Default___.Fm7(_640_v0, _this.F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(217))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg42 _dafny.Int) interface{} {
				return coer42(arg42)
			}
		}((func(_642_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_643_i0 _dafny.Int) _dafny.Int {
				return _642_v0
			}
		})(_640_v0)))).Cardinality()), globalState))
		var _644_v2 _dafny.CodePoint
		_ = _644_v2
		_644_v2 = _dafny.CodePoint('w')
		_641_v1 = (_641_v1).Update(_640_v0, _644_v2)
		_640_v0 = (_640_v0).Minus((_640_v0).Plus(_640_v0))
		var _645_v3 _dafny.Map
		_ = _645_v3
		_645_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_v0, _640_v0)
		var _646_v4 D3
		_ = _646_v4
		_646_v4 = Companion_D3_.Create_DC6_((_this).F9(), ((_645_v3).Cardinality()).Times(_640_v0), _640_v0)
		var _source12 D3 = _646_v4
		_ = _source12
		if _source12.Is_DC6() {
			var _647___mcc_h0 bool = _source12.Get_().(D3_DC6).Cf7
			_ = _647___mcc_h0
			var _648___mcc_h1 _dafny.Int = _source12.Get_().(D3_DC6).Cf8
			_ = _648___mcc_h1
			var _649___mcc_h2 _dafny.Int = _source12.Get_().(D3_DC6).Cf9
			_ = _649___mcc_h2
			var _650_cf9 _dafny.Int = _649___mcc_h2
			_ = _650_cf9
			var _651_cf8 _dafny.Int = _648___mcc_h1
			_ = _651_cf8
			var _652_cf7 bool = _647___mcc_h0
			_ = _652_cf7
			_644_v2 = (func() _dafny.CodePoint {
				if (_this).F9() {
					return _644_v2
				}
				return _644_v2
			})()
			var _653_v5 _dafny.Array
			_ = _653_v5
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_10
			var _nw111 _dafny.Array
			_ = _nw111
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw111 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) bool = (func(_654_cf8 _dafny.Int) func(_dafny.Int) bool {
					return func(_655_i1 _dafny.Int) bool {
						return (_654_cf8).Cmp(_654_cf8) != 0
					}
				})(_651_cf8)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw111 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw111).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw111).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_653_v5 = _nw111
			var _rhs106 bool = (_this).F9()
			_ = _rhs106
			var _rhs107 _dafny.Int = (Companion_Default___.SafeDivisionInt(_640_v0, _650_cf9)).Times(_650_cf9)
			_ = _rhs107
			var _rhs108 _dafny.Array = _653_v5
			_ = _rhs108
			_652_cf7 = _rhs106
			_640_v0 = _rhs107
			_653_v5 = _rhs108
			_650_cf9 = (_640_v0).Plus((_650_cf9).Minus(_640_v0))
			_653_v5 = _653_v5
		} else {
			var _656___mcc_h3 _dafny.Sequence = _source12.Get_().(D3_DC5).Cf6
			_ = _656___mcc_h3
			var _657_cf6 _dafny.Sequence = _656___mcc_h3
			_ = _657_cf6
			var _658_v6 _dafny.Array
			_ = _658_v6
			var _nwElement0_22 bool = (_640_v0).Cmp(_640_v0) != 0
			_ = _nwElement0_22
			var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(6))
			_ = _nw112
			(_nw112).ArraySet1(_nwElement0_22, 0)
			(_nw112).ArraySet1((_this).F9(), 1)
			(_nw112).ArraySet1((_this.F8()) == (!(_this.F8())), 2)
			(_nw112).ArraySet1(true, 3)
			(_nw112).ArraySet1((_this).F9(), 4)
			(_nw112).ArraySet1(true, 5)
			_658_v6 = _nw112
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_658_v6), 0))
			_ = _index120
			(_658_v6).ArraySet1(p2, (_index120).Int())
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(389), _dafny.ArrayLen((_658_v6), 0))
			_ = _index121
			(_658_v6).ArraySet1(p2, (_index121).Int())
			_640_v0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_640_v0), _640_v0)
			var _659_v7 _dafny.Sequence
			_ = _659_v7
			_659_v7 = _dafny.UnicodeSeqOfUtf8Bytes("tfsokpds")
			var _660_v8 D4
			_ = _660_v8
			_660_v8 = Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ussmtod"))
			var _661_v9 _dafny.Sequence
			_ = _661_v9
			_661_v9 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ypcmbt"), _dafny.Companion_Sequence_.Concatenate(_659_v7, (_660_v8).Dtor_cf10()))
			_659_v7 = (_661_v9).Select((Companion_Default___.SafeIndex(_640_v0, _dafny.IntOfUint32((_661_v9).Cardinality()))).Uint32()).(_dafny.Sequence)
			var _662_v10 _dafny.Sequence
			_ = _662_v10
			_662_v10 = _dafny.SeqOf(_660_v8, Companion_D4_.Create_DC7_(_659_v7), _660_v8)
			var _663_v11 _dafny.Set
			_ = _663_v11
			_663_v11 = _dafny.SetOf(_662_v10)
			var _664_v12 *C1
			_ = _664_v12
			var _nw113 *C1 = New_C1_()
			_ = _nw113
			_nw113.Ctor__(_644_v2, _663_v11, (_this).F9(), (_this).F9())
			_664_v12 = _nw113
		}
		(_this).F8_set_((_this).F9())
		var _665_v13 _dafny.Map
		_ = _665_v13
		_665_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_640_v0, p2)
		var _666_v14 _dafny.Sequence
		_ = _666_v14
		_666_v14 = _dafny.SeqOf(_665_v13, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F9())).Cardinality(), p2))
		(_this).F8_set_(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_666_v14, _666_v14), _666_v14))
		var _667_v15 _dafny.Array
		_ = _667_v15
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_11
		var _nw114 _dafny.Array
		_ = _nw114
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw114 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Int = func(_668_i2 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_668_i2, _dafny.IntOfInt64(540))
			}
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw114 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw114).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw114).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_667_v15 = _nw114
		var _source13 D5 = Companion_D5_.Create_DC10_(_667_v15)
		_ = _source13
		if _source13.Is_DC11() {
			var _669___mcc_h4 _dafny.Int = _source13.Get_().(D5_DC11).Cf18
			_ = _669___mcc_h4
			var _670___mcc_h5 _dafny.Sequence = _source13.Get_().(D5_DC11).Cf19
			_ = _670___mcc_h5
			var _671_cf19 _dafny.Sequence = _670___mcc_h5
			_ = _671_cf19
			var _672_cf18 _dafny.Int = _669___mcc_h4
			_ = _672_cf18
			var _673_v16 _dafny.Array
			_ = _673_v16
			var _nwElement0_23 bool = p2
			_ = _nwElement0_23
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(29))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_23, 0)
			(_nw115).ArraySet1(!(!(_this.F8()) || ((_this).F9())), 1)
			(_nw115).ArraySet1(p2, 2)
			(_nw115).ArraySet1(p2, 3)
			(_nw115).ArraySet1(p2, 4)
			(_nw115).ArraySet1(p2, 5)
			(_nw115).ArraySet1((_this).F9(), 6)
			(_nw115).ArraySet1(_this.F8(), 7)
			(_nw115).ArraySet1(p2, 8)
			(_nw115).ArraySet1(_this.F8(), 9)
			(_nw115).ArraySet1(_this.F8(), 10)
			(_nw115).ArraySet1(((_dafny.Zero).Minus(_640_v0)).Cmp((_dafny.Zero).Minus(_672_cf18)) == 0, 11)
			(_nw115).ArraySet1(_this.F8(), 12)
			(_nw115).ArraySet1((_this).F9(), 13)
			(_nw115).ArraySet1((_this).F9(), 14)
			(_nw115).ArraySet1(_this.F8(), 15)
			(_nw115).ArraySet1((_672_cf18).Cmp((_this).Fm6(_672_cf18, _672_cf18, globalState)) >= 0, 16)
			(_nw115).ArraySet1(p2, 17)
			(_nw115).ArraySet1(!((_this).F9()) || (_this.F8()), 18)
			(_nw115).ArraySet1(p2, 19)
			(_nw115).ArraySet1(_this.F8(), 20)
			(_nw115).ArraySet1(p2, 21)
			(_nw115).ArraySet1(_this.F8(), 22)
			(_nw115).ArraySet1(!(_this.F8()), 23)
			(_nw115).ArraySet1(p2, 24)
			(_nw115).ArraySet1(p2, 25)
			(_nw115).ArraySet1(_this.F8(), 26)
			(_nw115).ArraySet1((_672_cf18).Cmp(_dafny.IntOfInt64(552)) > 0, 27)
			(_nw115).ArraySet1(false, 28)
			_673_v16 = _nw115
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))
			_ = _index122
			(_673_v16).ArraySet1(p2, (_index122).Int())
			var _674_v17 D4
			_ = _674_v17
			_674_v17 = Companion_D4_.Create_DC8_(_640_v0, (_this).F9())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))
			_ = _index123
			(_673_v16).ArraySet1((_674_v17).Dtor_cf12(), (_index123).Int())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))
			_ = _index124
			var _rhs109 _dafny.Int = _640_v0
			_ = _rhs109
			var _rhs110 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_672_cf18).Times(_640_v0), (func() bool {
				if p2 {
					return (_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool)
				}
				return _this.F8()
			})())).Cardinality()
			_ = _rhs110
			var _rhs111 bool = (_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool)
			_ = _rhs111
			var _lhs64 _dafny.Array = _673_v16
			_ = _lhs64
			var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))
			_ = _lhs65
			_672_cf18 = _rhs109
			_640_v0 = _rhs110
			(_lhs64).ArraySet1(_rhs111, (_lhs65).Int())
			var _675_v18 _dafny.MultiSet
			_ = _675_v18
			_675_v18 = _dafny.MultiSetOf(_640_v0, (_dafny.Zero).Minus(_672_cf18))
			var _676_v19 _dafny.Map
			_ = _676_v19
			_676_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_cf18, _675_v18)
			var _677_v20 _dafny.Map
			_ = _677_v20
			_677_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), !((_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool)))
			var _678_v21 _dafny.Map
			_ = _678_v21
			_678_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_677_v20, _644_v2)
			var _679_v22 _dafny.Map
			_ = _679_v22
			_679_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_675_v18).Union((func() _dafny.MultiSet {
				if (_676_v19).Contains((_678_v21).Cardinality()) {
					return (_676_v19).Get((_678_v21).Cardinality()).(_dafny.MultiSet)
				}
				return _675_v18
			})()), _667_v15)
			_679_v22 = (_679_v22).Update(_675_v18, _667_v15)
			var _680_v23 _dafny.Set
			_ = _680_v23
			_680_v23 = _dafny.SetOf((_640_v0).Plus(_640_v0), _672_cf18, _640_v0, _672_cf18)
			var _681_v24 _dafny.Sequence
			_ = _681_v24
			_681_v24 = _dafny.SeqOf(_this.F8())
			var _682_v25 _dafny.Map
			_ = _682_v25
			_682_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm15(globalState)).Dtor_cf5(), _681_v24)
			var _683_v26 _dafny.Sequence
			_ = _683_v26
			_683_v26 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(_672_cf18, _dafny.IntOfInt64(894)))
			var _684_v27 _dafny.Map
			_ = _684_v27
			_684_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool), _680_v23)
			var _685_v28 _dafny.Map
			_ = _685_v28
			_685_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_644_v2, _672_cf18)
			var _686_v29 _dafny.Map
			_ = _686_v29
			_686_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_cf18, _685_v28)
			var _687_v30 _dafny.Sequence
			_ = _687_v30
			_687_v30 = _dafny.SeqOf(_682_v25)
			var _rhs112 _dafny.Set = ((func() _dafny.Set {
				if (_684_v27).Contains((_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool)) {
					return (_684_v27).Get((_673_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_673_v16), 0))).Int()).(bool)).(_dafny.Set)
				}
				return _dafny.SetOf((_686_v29).Cardinality())
			})()).Intersection(_680_v23)
			_ = _rhs112
			var _rhs113 _dafny.Map = (_687_v30).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_640_v0), _dafny.IntOfUint32((_687_v30).Cardinality()))).Uint32()).(_dafny.Map)
			_ = _rhs113
			var _rhs114 _dafny.Sequence = _683_v26
			_ = _rhs114
			var _rhs115 _dafny.Int = Companion_Default___.SafeModuloInt(_640_v0, _640_v0)
			_ = _rhs115
			_680_v23 = _rhs112
			_682_v25 = _rhs113
			_683_v26 = _rhs114
			_672_cf18 = _rhs115
		} else {
			var _688___mcc_h6 _dafny.Array = _source13.Get_().(D5_DC10).Cf17
			_ = _688___mcc_h6
			var _689_cf17 _dafny.Array = _688___mcc_h6
			_ = _689_cf17
			var _690_v31 _dafny.MultiSet
			_ = _690_v31
			_690_v31 = _dafny.MultiSetOf(_640_v0)
			var _691_v32 _dafny.MultiSet
			_ = _691_v32
			_691_v32 = _dafny.MultiSetOf((p2) == (_this.F8()), ((_this).F9()) == ((_this).F9()), p2, true, !(((_690_v31).Cardinality()).Cmp(_640_v0) == 0))
			var _692_v33 _dafny.Set
			_ = _692_v33
			_692_v33 = _dafny.SetOf(p2)
			var _693_v34 D7
			_ = _693_v34
			_693_v34 = Companion_D7_.Create_DC14_(_692_v33)
			var _694_v35 _dafny.Sequence
			_ = _694_v35
			_694_v35 = _dafny.SeqOf(((_693_v34).Dtor_cf25()).IsDisjointFrom(_692_v33))
			var _rhs116 _dafny.Int = ((_690_v31).Update(_640_v0, Companion_Default___.Abs(_640_v0))).Cardinality()
			_ = _rhs116
			var _rhs117 _dafny.MultiSet = _dafny.MultiSetFromSeq(_694_v35)
			_ = _rhs117
			_640_v0 = _rhs116
			_691_v32 = _rhs117
			var _695_v36 _dafny.Sequence
			_ = _695_v36
			_695_v36 = _dafny.UnicodeSeqOfUtf8Bytes("tpt")
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))
			_ = _index125
			(_667_v15).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_695_v36, _695_v36)).Cardinality()), (_index125).Int())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))
			_ = _index126
			var _rhs118 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if false {
					return _640_v0
				}
				return _640_v0
			})(), Companion_Default___.SafeModuloInt(_640_v0, _640_v0))
			_ = _rhs118
			var _rhs119 _dafny.Int = _640_v0
			_ = _rhs119
			var _lhs66 _dafny.Array = _667_v15
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))
			_ = _lhs67
			_640_v0 = _rhs118
			(_lhs66).ArraySet1(_rhs119, (_lhs67).Int())
			var _696_v37 D7
			_ = _696_v37
			_696_v37 = Companion_D7_.Create_DC14_(_692_v33)
			var _697_v38 D7
			_ = _697_v38
			_697_v38 = Companion_D7_.Create_DC16_(_696_v37)
			var _698_v39 _dafny.Sequence
			_ = _698_v39
			_698_v39 = _dafny.SeqOf(_696_v37)
			var _699_v40 D7
			_ = _699_v40
			_699_v40 = Companion_D7_.Create_DC16_((_698_v39).Select((Companion_Default___.SafeIndex((_dafny.SetOf((_this).F9(), _this.F8())).Cardinality(), _dafny.IntOfUint32((_698_v39).Cardinality()))).Uint32()).(D7))
			_699_v40 = _699_v40
			if false {
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))
				_ = _index127
				(_667_v15).ArraySet1((_640_v0).Times(_640_v0), (_index127).Int())
				var _700_v41 _dafny.Map
				_ = _700_v41
				_700_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_667_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))).Int()).(_dafny.Int)).Cmp(_640_v0) <= 0, _689_cf17)
				_700_v41 = (_700_v41).Update(p2, _689_cf17)
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))
				_ = _index128
				(_667_v15).ArraySet1((_667_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))).Int()).(_dafny.Int), (_index128).Int())
				(_this).F8_set_((_this).F9())
				var _701_v42 _dafny.Sequence
				_ = _701_v42
				_701_v42 = _dafny.SeqOf(_689_cf17)
				var _702_v43 _dafny.Map
				_ = _702_v43
				_702_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v36, _689_cf17)
				var _703_v44 _dafny.Array
				_ = _703_v44
				var _nwElement0_24 _dafny.Array = _689_cf17
				_ = _nwElement0_24
				var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(18))
				_ = _nw116
				(_nw116).ArraySet1(_nwElement0_24, 0)
				(_nw116).ArraySet1(_689_cf17, 1)
				(_nw116).ArraySet1(_689_cf17, 2)
				(_nw116).ArraySet1(_689_cf17, 3)
				(_nw116).ArraySet1(_689_cf17, 4)
				(_nw116).ArraySet1(_689_cf17, 5)
				(_nw116).ArraySet1(_689_cf17, 6)
				(_nw116).ArraySet1(_689_cf17, 7)
				(_nw116).ArraySet1(_689_cf17, 8)
				(_nw116).ArraySet1(_689_cf17, 9)
				(_nw116).ArraySet1((_701_v42).Select((Companion_Default___.SafeIndex((_667_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_701_v42).Cardinality()))).Uint32()).(_dafny.Array), 10)
				(_nw116).ArraySet1(_689_cf17, 11)
				(_nw116).ArraySet1(_689_cf17, 12)
				(_nw116).ArraySet1(_689_cf17, 13)
				(_nw116).ArraySet1((func() _dafny.Array {
					if (_702_v43).Contains(_695_v36) {
						return (_702_v43).Get(_695_v36).(_dafny.Array)
					}
					return _689_cf17
				})(), 14)
				(_nw116).ArraySet1(_689_cf17, 15)
				(_nw116).ArraySet1(_689_cf17, 16)
				(_nw116).ArraySet1(_689_cf17, 17)
				_703_v44 = _nw116
				var _704_v45 *C0
				_ = _704_v45
				var _nw117 *C0 = New_C0_()
				_ = _nw117
				_nw117.Ctor__(_703_v44)
				_704_v45 = _nw117
			} else {
				_640_v0 = (_667_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))).Int()).(_dafny.Int)
				var _705_v46 _dafny.Sequence
				_ = _705_v46
				_705_v46 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("u"))
				var _706_v47 D0
				_ = _706_v47
				_706_v47 = Companion_D0_.Create_DC1_(_this.F8(), _694_v35, p2)
				(_this).F8_set_((_this).Fm1(_dafny.Companion_Sequence_.Concatenate(_695_v36, (_705_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(568), _dafny.IntOfUint32((_705_v46).Cardinality()))).Uint32()).(_dafny.Sequence)), func(_pat_let27_0 D0) D0 {
					return func(_707_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let28_0 bool) D0 {
							return func(_708_dt__update_hcf2_h0 bool) D0 {
								return Companion_D0_.Create_DC1_((_707_dt__update__tmp_h0).Dtor_cf0(), (_707_dt__update__tmp_h0).Dtor_cf1(), _708_dt__update_hcf2_h0)
							}(_pat_let28_0)
						}((_this).F9())
					}(_pat_let27_0)
				}(_706_v47), globalState))
				_690_v31 = _690_v31
				var _709_v48 _dafny.Array
				_ = _709_v48
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_12
				var _nw118 _dafny.Array
				_ = _nw118
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw118 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) bool = func(_710_i3 _dafny.Int) bool {
						return _this.F8()
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw118 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw118).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw118).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_709_v48 = _nw118
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_709_v48), 0))
				_ = _index129
				(_709_v48).ArraySet1(_this.F8(), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(805), _dafny.ArrayLen((_709_v48), 0))
				_ = _index130
				(_709_v48).ArraySet1((_this).F9(), (_index130).Int())
				var _711_v49 _dafny.Sequence
				_ = _711_v49
				_711_v49 = _dafny.SeqOf(Companion_D4_.Create_DC7_(_695_v36))
				var _712_v50 _dafny.Set
				_ = _712_v50
				_712_v50 = _dafny.SetOf(_711_v49, _711_v49)
				var _713_v51 T2
				_ = _713_v51
				var _nw119 *C1 = New_C1_()
				_ = _nw119
				_nw119.Ctor__(_644_v2, _712_v50, false, !(p2))
				_713_v51 = _nw119
				var _714_v52 _dafny.Array
				_ = _714_v52
				var _nwElement0_25 T2 = _713_v51
				_ = _nwElement0_25
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(14))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_25, 0)
				(_nw120).ArraySet1(_713_v51, 1)
				(_nw120).ArraySet1(_713_v51, 2)
				(_nw120).ArraySet1(_713_v51, 3)
				(_nw120).ArraySet1(_713_v51, 4)
				(_nw120).ArraySet1(_713_v51, 5)
				(_nw120).ArraySet1(_713_v51, 6)
				(_nw120).ArraySet1(_713_v51, 7)
				(_nw120).ArraySet1(_713_v51, 8)
				(_nw120).ArraySet1(_713_v51, 9)
				(_nw120).ArraySet1(_713_v51, 10)
				(_nw120).ArraySet1(_713_v51, 11)
				(_nw120).ArraySet1(_713_v51, 12)
				(_nw120).ArraySet1(_713_v51, 13)
				_714_v52 = _nw120
				var _715_v53 _dafny.Map
				_ = _715_v53
				_715_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v36, _714_v52)
				var _716_v54 _dafny.Map
				_ = _716_v54
				_716_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.Companion_Sequence_.Update(_695_v36, (Companion_Default___.SafeIndex((_667_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(72), _dafny.ArrayLen((_667_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_695_v36).Cardinality()))).Uint32(), _644_v2))
				_715_v53 = (_715_v53).Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_716_v54).Contains(p2) {
						return (_716_v54).Get(p2).(_dafny.Sequence)
					}
					return _695_v36
				})(), _695_v36), _714_v52)
			}
		}
	}
}
func (_this *C2) M1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _717_v0 _dafny.Array
		_ = _717_v0
		var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
		_ = _nw121
		_717_v0 = _nw121
		var _718_v1 _dafny.Int
		_ = _718_v1
		_718_v1 = _dafny.IntOfInt64(345)
		var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))
		_ = _index131
		(_717_v0).ArraySet1(_718_v1, (_index131).Int())
		var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))
		_ = _index132
		(_717_v0).ArraySet1((_this).Fm6(_718_v1, (_718_v1).Plus(_718_v1), globalState), (_index132).Int())
		var _719_v2 D5
		_ = _719_v2
		_719_v2 = Companion_D5_.Create_DC10_(_717_v0)
		var _source14 D5 = _719_v2
		_ = _source14
		if _source14.Is_DC11() {
			var _720___mcc_h0 _dafny.Int = _source14.Get_().(D5_DC11).Cf18
			_ = _720___mcc_h0
			var _721___mcc_h1 _dafny.Sequence = _source14.Get_().(D5_DC11).Cf19
			_ = _721___mcc_h1
			var _722_cf19 _dafny.Sequence = _721___mcc_h1
			_ = _722_cf19
			var _723_cf18 _dafny.Int = _720___mcc_h0
			_ = _723_cf18
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))
			_ = _index133
			(_717_v0).ArraySet1(_718_v1, (_index133).Int())
			r1 = _dafny.IntOfInt64(590)
			var _724_v3 _dafny.Array
			_ = _724_v3
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
			_ = _nw122
			_724_v3 = _nw122
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_724_v3), 0))
			_ = _index134
			(_724_v3).ArraySet1((func() bool {
				if _this.F8() {
					return _this.F8()
				}
				return p1
			})(), (_index134).Int())
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(487), _dafny.ArrayLen((_724_v3), 0))
			_ = _index135
			(_724_v3).ArraySet1((_this).F9(), (_index135).Int())
			_717_v0 = _717_v0
		} else {
			var _725___mcc_h2 _dafny.Array = _source14.Get_().(D5_DC10).Cf17
			_ = _725___mcc_h2
			var _726_cf17 _dafny.Array = _725___mcc_h2
			_ = _726_cf17
			var _727_v4 _dafny.CodePoint
			_ = _727_v4
			_727_v4 = _dafny.CodePoint('o')
			var _728_v5 D4
			_ = _728_v5
			_728_v5 = Companion_D4_.Create_DC7_(p0)
			var _729_v6 _dafny.Sequence
			_ = _729_v6
			_729_v6 = _dafny.SeqOf(Companion_D4_.Create_DC7_(p0), _728_v5, _728_v5, _728_v5, _728_v5)
			var _730_v7 _dafny.Array
			_ = _730_v7
			var _nwElement0_26 _dafny.Array = _726_cf17
			_ = _nwElement0_26
			var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(28))
			_ = _nw123
			(_nw123).ArraySet1(_nwElement0_26, 0)
			(_nw123).ArraySet1(_717_v0, 1)
			(_nw123).ArraySet1(_717_v0, 2)
			(_nw123).ArraySet1(_726_cf17, 3)
			(_nw123).ArraySet1(_726_cf17, 4)
			(_nw123).ArraySet1(_717_v0, 5)
			(_nw123).ArraySet1(_726_cf17, 6)
			(_nw123).ArraySet1(_717_v0, 7)
			(_nw123).ArraySet1(_717_v0, 8)
			(_nw123).ArraySet1(_726_cf17, 9)
			(_nw123).ArraySet1(_726_cf17, 10)
			(_nw123).ArraySet1(_726_cf17, 11)
			(_nw123).ArraySet1(_726_cf17, 12)
			(_nw123).ArraySet1(_717_v0, 13)
			(_nw123).ArraySet1(_726_cf17, 14)
			(_nw123).ArraySet1(_717_v0, 15)
			(_nw123).ArraySet1(_726_cf17, 16)
			(_nw123).ArraySet1(_717_v0, 17)
			(_nw123).ArraySet1(_726_cf17, 18)
			(_nw123).ArraySet1(_726_cf17, 19)
			(_nw123).ArraySet1(_717_v0, 20)
			(_nw123).ArraySet1(_726_cf17, 21)
			(_nw123).ArraySet1(_717_v0, 22)
			(_nw123).ArraySet1(_726_cf17, 23)
			(_nw123).ArraySet1(_726_cf17, 24)
			(_nw123).ArraySet1(_717_v0, 25)
			(_nw123).ArraySet1(_717_v0, 26)
			(_nw123).ArraySet1(_726_cf17, 27)
			_730_v7 = _nw123
			var _731_v8 *C0
			_ = _731_v8
			var _nw124 *C0 = New_C0_()
			_ = _nw124
			_nw124.Ctor__(_730_v7)
			_731_v8 = _nw124
			var _732_v9 _dafny.Map
			_ = _732_v9
			_732_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_731_v8, _this.F8())
			var _733_v10 _dafny.Set
			_ = _733_v10
			_733_v10 = _dafny.SetOf(_dafny.SeqOf(_728_v5), _729_v6, Companion_Default___.Fm16(false, _718_v1, (_732_v9).Cardinality(), globalState), _729_v6)
			var _734_v11 *C1
			_ = _734_v11
			var _nw125 *C1 = New_C1_()
			_ = _nw125
			_nw125.Ctor__(_727_v4, _733_v10, ((_this).F9()) && (_this.F8()), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-972))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_735_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_736_i0 _dafny.Int) _dafny.CodePoint {
					return _735_v4
				}
			})(_727_v4))), p0))
			_734_v11 = _nw125
			(_this).F8_set_(true)
			_726_cf17 = _726_cf17
			r1 = (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)
		}
		var _737_v12 _dafny.Sequence
		_ = _737_v12
		_737_v12 = _dafny.SeqOf(_717_v0, _717_v0, _717_v0, _717_v0, _717_v0)
		var _738_v13 _dafny.Map
		_ = _738_v13
		_738_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _717_v0)
		var _739_v14 _dafny.Array
		_ = _739_v14
		var _nwElement0_27 _dafny.Array = _717_v0
		_ = _nwElement0_27
		var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
		_ = _nw126
		(_nw126).ArraySet1(_nwElement0_27, 0)
		(_nw126).ArraySet1(_717_v0, 1)
		(_nw126).ArraySet1(_717_v0, 2)
		(_nw126).ArraySet1(_717_v0, 3)
		(_nw126).ArraySet1(_717_v0, 4)
		(_nw126).ArraySet1(_717_v0, 5)
		(_nw126).ArraySet1((_737_v12).Select((Companion_Default___.SafeIndex(_718_v1, _dafny.IntOfUint32((_737_v12).Cardinality()))).Uint32()).(_dafny.Array), 6)
		(_nw126).ArraySet1(_717_v0, 7)
		(_nw126).ArraySet1(_717_v0, 8)
		(_nw126).ArraySet1(_717_v0, 9)
		(_nw126).ArraySet1(_717_v0, 10)
		(_nw126).ArraySet1((func() _dafny.Array {
			if (_738_v13).Contains(_this.F8()) {
				return (_738_v13).Get(_this.F8()).(_dafny.Array)
			}
			return (_719_v2).Dtor_cf17()
		})(), 11)
		(_nw126).ArraySet1(_717_v0, 12)
		(_nw126).ArraySet1(_717_v0, 13)
		_739_v14 = _nw126
		var _740_v15 *C0
		_ = _740_v15
		var _nw127 *C0 = New_C0_()
		_ = _nw127
		_nw127.Ctor__(_739_v14)
		_740_v15 = _nw127
		var _741_v16 _dafny.MultiSet
		_ = _741_v16
		_741_v16 = _dafny.MultiSetOf(p1)
		var _742_v17 _dafny.Sequence
		_ = _742_v17
		_742_v17 = _dafny.SeqOf((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))
		var _743_v18 _dafny.Set
		_ = _743_v18
		_743_v18 = _dafny.SetOf(_this.F8(), (_this).F9(), (_this).F9(), _this.F8(), p1)
		var _744_v19 D7
		_ = _744_v19
		_744_v19 = Companion_D7_.Create_DC14_(_743_v18)
		var _745_v20 _dafny.Map
		_ = _745_v20
		_745_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), p1)
		var _746_v21 _dafny.Sequence
		_ = _746_v21
		_746_v21 = _dafny.SeqOf((_this).F9())
		var _747_v22 D0
		_ = _747_v22
		_747_v22 = Companion_D0_.Create_DC1_(p1, _746_v21, (_746_v21).Select((Companion_Default___.SafeIndex((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_746_v21).Cardinality()))).Uint32()).(bool))
		var _748_v23 _dafny.Array
		_ = _748_v23
		var _nwElement0_28 bool = (_740_v15).Fm0(globalState)
		_ = _nwElement0_28
		var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(27))
		_ = _nw128
		(_nw128).ArraySet1(_nwElement0_28, 0)
		(_nw128).ArraySet1(p1, 1)
		(_nw128).ArraySet1((_this).F9(), 2)
		(_nw128).ArraySet1(p1, 3)
		(_nw128).ArraySet1(true, 4)
		(_nw128).ArraySet1(_this.F8(), 5)
		(_nw128).ArraySet1(_this.F8(), 6)
		(_nw128).ArraySet1(p1, 7)
		(_nw128).ArraySet1(p1, 8)
		(_nw128).ArraySet1(p1, 9)
		(_nw128).ArraySet1((_this).F9(), 10)
		(_nw128).ArraySet1((_this).F9(), 11)
		(_nw128).ArraySet1((_this).F9(), 12)
		(_nw128).ArraySet1((_this).F9(), 13)
		(_nw128).ArraySet1(p1, 14)
		(_nw128).ArraySet1((_this).F9(), 15)
		(_nw128).ArraySet1((_this).F9(), 16)
		(_nw128).ArraySet1(_this.F8(), 17)
		(_nw128).ArraySet1(false, 18)
		(_nw128).ArraySet1((_this).F9(), 19)
		(_nw128).ArraySet1((_this).F9(), 20)
		(_nw128).ArraySet1(_this.F8(), 21)
		(_nw128).ArraySet1(p1, 22)
		(_nw128).ArraySet1(p1, 23)
		(_nw128).ArraySet1((_this).F9(), 24)
		(_nw128).ArraySet1(p1, 25)
		(_nw128).ArraySet1(p1, 26)
		_748_v23 = _nw128
		var _749_v24 _dafny.Sequence
		_ = _749_v24
		_749_v24 = _dafny.SeqOf(_748_v23, _748_v23)
		var _750_v25 _dafny.Map
		_ = _750_v25
		_750_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))
		var _751_v26 _dafny.Array
		_ = _751_v26
		var _nwElement0_29 bool = (_718_v1).Cmp((func() _dafny.Int {
			if (_741_v16).Contains(_this.F8()) {
				return (_741_v16).Multiplicity(_this.F8())
			}
			return (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)
		})()) < 0
		_ = _nwElement0_29
		var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(26))
		_ = _nw129
		(_nw129).ArraySet1(_nwElement0_29, 0)
		(_nw129).ArraySet1(_this.F8(), 1)
		(_nw129).ArraySet1((_this).Fm0(globalState), 2)
		(_nw129).ArraySet1((_this).F9(), 3)
		(_nw129).ArraySet1((_dafny.IntOfInt64(558)).Cmp(_718_v1) <= 0, 4)
		(_nw129).ArraySet1((_this).F9(), 5)
		(_nw129).ArraySet1((_this).F9(), 6)
		(_nw129).ArraySet1((_dafny.IntOfUint32((_742_v17).Cardinality())).Cmp((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)) >= 0, 7)
		(_nw129).ArraySet1(_this.F8(), 8)
		(_nw129).ArraySet1((_dafny.SetOf(p1, (_this).F9())).IsDisjointFrom((_744_v19).Dtor_cf25()), 9)
		(_nw129).ArraySet1(_this.F8(), 10)
		(_nw129).ArraySet1(_this.F8(), 11)
		(_nw129).ArraySet1(_this.F8(), 12)
		(_nw129).ArraySet1(p1, 13)
		(_nw129).ArraySet1(false, 14)
		(_nw129).ArraySet1(p1, 15)
		(_nw129).ArraySet1(!(((_745_v20).Cardinality()).Cmp(_718_v1) > 0), 16)
		(_nw129).ArraySet1((_740_v15).Fm1(p0, _747_v22, globalState), 17)
		(_nw129).ArraySet1(_this.F8(), 18)
		(_nw129).ArraySet1(!((_750_v25).Contains(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_749_v24).Cardinality()), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))).Cardinality()))), 19)
		(_nw129).ArraySet1((_this).F9(), 20)
		(_nw129).ArraySet1(p1, 21)
		(_nw129).ArraySet1(p1, 22)
		(_nw129).ArraySet1(!(_this.F8()) || ((_this).F9()), 23)
		(_nw129).ArraySet1(_this.F8(), 24)
		(_nw129).ArraySet1(p1, 25)
		_751_v26 = _nw129
		var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))
		_ = _index136
		(_748_v23).ArraySet1(p1, (_index136).Int())
		var _752_v27 _dafny.Set
		_ = _752_v27
		_752_v27 = _dafny.SetOf(_718_v1)
		var _753_v28 D8
		_ = _753_v28
		_753_v28 = Companion_D8_.Create_DC17_(_752_v27)
		var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))
		_ = _index137
		var _rhs120 bool = (_dafny.SetOf(_dafny.IntOfInt64(120))).IsProperSubsetOf(((_753_v28).Dtor_cf27()).Difference(_752_v27))
		_ = _rhs120
		var _rhs121 *C0 = _740_v15
		_ = _rhs121
		var _rhs122 _dafny.Int = (_dafny.Zero).Minus((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))
		_ = _rhs122
		var _rhs123 bool = p1
		_ = _rhs123
		var _rhs124 bool = !((_dafny.IntOfUint32((_742_v17).Cardinality())).Cmp(Companion_Default___.SafeModuloInt(_718_v1, (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int))) >= 0)
		_ = _rhs124
		var _lhs68 _dafny.Array = _748_v23
		_ = _lhs68
		var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))
		_ = _lhs69
		r2 = _rhs120
		_740_v15 = _rhs121
		_718_v1 = _rhs122
		(_lhs68).ArraySet1(_rhs123, (_lhs69).Int())
		r2 = _rhs124
		var _754_v29 D2
		_ = _754_v29
		_754_v29 = Companion_D2_.Create_DC4_((_this).F9())
		var _755_i1 _dafny.Int
		_ = _755_i1
		_755_i1 = _dafny.Zero
		{
			var _pat_let_tv22 = _748_v23
			_ = _pat_let_tv22
			var _pat_let_tv23 = _748_v23
			_ = _pat_let_tv23
			var _pat_let_tv24 = _748_v23
			_ = _pat_let_tv24
			var _pat_let_tv25 = _748_v23
			_ = _pat_let_tv25
			var _pat_let_tv26 = _748_v23
			_ = _pat_let_tv26
			var _pat_let_tv27 = _748_v23
			_ = _pat_let_tv27
			var _pat_let_tv28 = p1
			_ = _pat_let_tv28
			var _pat_let_tv29 = p1
			_ = _pat_let_tv29
			var _pat_let_tv30 = _748_v23
			_ = _pat_let_tv30
			var _pat_let_tv31 = _748_v23
			_ = _pat_let_tv31
			for func(_source15 D2) bool {
				if _source15.Is_DC4() {
					var _768___mcc_h3 bool = _source15.Get_().(D2_DC4).Cf5
					_ = _768___mcc_h3
					var _769_cf5 bool = _768___mcc_h3
					_ = _769_cf5
					if (_pat_let_tv23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_pat_let_tv22), 0))).Int()).(bool) {
						return (Companion_D8_.Create_DC19_((_this).F9())).Dtor_cf31()
					} else {
						return (Companion_D6_.Create_DC13_((_pat_let_tv25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_pat_let_tv24), 0))).Int()).(bool), (_this).F9(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-939))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg48 _dafny.Int) interface{} {
								return coer48(arg48)
							}
						}(func(_770_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						})), _769_cf5)).Dtor_cf21()
					}
				} else {
					var _771___mcc_h4 _dafny.Int = _source15.Get_().(D2_DC3).Cf4
					_ = _771___mcc_h4
					var _772_cf4 _dafny.Int = _771___mcc_h4
					_ = _772_cf4
					if !((_pat_let_tv27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_pat_let_tv26), 0))).Int()).(bool)) {
						return _pat_let_tv28
					} else {
						return _pat_let_tv29
					}
				}
			}(func(_pat_let29_0 D2) D2 {
				return func(_773_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let30_0 bool) D2 {
						return func(_774_dt__update_hcf5_h0 bool) D2 {
							return Companion_D2_.Create_DC4_(_774_dt__update_hcf5_h0)
						}(_pat_let30_0)
					}(!((_pat_let_tv31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_pat_let_tv30), 0))).Int()).(bool)))
				}(_pat_let29_0)
			}(_754_v29)) {
				{
					if (_755_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_755_i1 = (_755_i1).Plus(_dafny.One)
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))
					_ = _index138
					(_717_v0).ArraySet1(_dafny.IntOfInt64(-765), (_index138).Int())
					(_this).F8_set_(p1)
					var _756_v30 _dafny.CodePoint
					_ = _756_v30
					_756_v30 = _dafny.CodePoint('m')
					var _757_v31 D4
					_ = _757_v31
					_757_v31 = Companion_D4_.Create_DC7_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(887))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_758_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_759_i4 _dafny.Int) _dafny.CodePoint {
							return _758_v30
						}
					})(_756_v30))))
					var _760_v32 _dafny.Set
					_ = _760_v32
					_760_v32 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer45 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}((func(_761_p0 _dafny.Sequence) func(_dafny.Int) D4 {
						return func(_762_i3 _dafny.Int) D4 {
							return Companion_D4_.Create_DC7_(_761_p0)
						}
					})(p0))), (Companion_Default___.SafeIndex((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer46 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_763_p0 _dafny.Sequence) func(_dafny.Int) D4 {
						return func(_764_i3 _dafny.Int) D4 {
							return Companion_D4_.Create_DC7_(_763_p0)
						}
					})(p0)))).Cardinality()))).Uint32(), _757_v31))
					var _765_v33 *C1
					_ = _765_v33
					var _nw130 *C1 = New_C1_()
					_ = _nw130
					_nw130.Ctor__(_756_v30, (_760_v32).Difference(_760_v32), (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool), (_this).F9())
					_765_v33 = _nw130
					_765_v33 = _765_v33
					_756_v30 = (func() _dafny.CodePoint {
						if !(p1) || ((_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)) {
							return Companion_Default___.Fm7((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_this).F9(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg47 _dafny.Int) interface{} {
									return coer47(arg47)
								}
							}((func(_766_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_767_i5 _dafny.Int) _dafny.CodePoint {
									return _766_v30
								}
							})(_756_v30)))).Cardinality()), globalState)
						}
						return (func() _dafny.CodePoint {
							if true {
								return _756_v30
							}
							return Companion_Default___.Fm7((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool), (_752_v27).Cardinality(), globalState)
						})()
					})()
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _775_i6 _dafny.Int
		_ = _775_i6
		_775_i6 = _dafny.Zero
		{
			for !(p1) {
				{
					if (_775_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_775_i6 = (_775_i6).Plus(_dafny.One)
					var _776_v34 _dafny.Sequence
					_ = _776_v34
					_776_v34 = _dafny.SeqOf(_746_v21)
					var _777_v35 D4
					_ = _777_v35
					_777_v35 = Companion_D4_.Create_DC7_(p0)
					var _778_v36 _dafny.Map
					_ = _778_v36
					_778_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_776_v34, _777_v35)
					_778_v36 = (_778_v36).Update(_776_v34, Companion_Default___.Fm17((_this).F9(), _this.F8(), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), globalState))
					var _779_v37 _dafny.Array
					_ = _779_v37
					var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
					_ = _nw131
					_779_v37 = _nw131
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_779_v37), 0))
					_ = _index139
					(_779_v37).ArraySet1(_751_v26, (_index139).Int())
					var _780_v38 _dafny.Array
					_ = _780_v38
					var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
					_ = _nw132
					_780_v38 = _nw132
					var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))
					_ = _index140
					(_780_v38).ArraySet1(_dafny.Companion_Sequence_.Concatenate(p0, _dafny.UnicodeSeqOfUtf8Bytes("iudnfs")), (_index140).Int())
					var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_779_v37), 0))
					_ = _index141
					var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))
					_ = _index142
					var _rhs125 bool = (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)
					_ = _rhs125
					var _rhs126 _dafny.Array = _748_v23
					_ = _rhs126
					var _rhs127 _dafny.Sequence = p0
					_ = _rhs127
					var _rhs128 _dafny.Sequence = p0
					_ = _rhs128
					var _lhs70 _dafny.Array = _779_v37
					_ = _lhs70
					var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(925), _dafny.ArrayLen((_779_v37), 0))
					_ = _lhs71
					var _lhs72 _dafny.Array = _780_v38
					_ = _lhs72
					var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))
					_ = _lhs73
					r2 = _rhs125
					(_lhs70).ArraySet1(_rhs126, (_lhs71).Int())
					(_lhs72).ArraySet1(_rhs127, (_lhs73).Int())
					r0 = _rhs128
					if (func() bool {
						if (func() bool {
							if p1 {
								return true
							}
							return (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)
						})() {
							return (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)
						}
						return (_718_v1).Cmp((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)) != 0
					})() {
						var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))
						_ = _index143
						(_717_v0).ArraySet1(((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)).Minus((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int)), (_index143).Int())
						var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))
						_ = _index144
						(_748_v23).ArraySet1((_this).Fm0(globalState), (_index144).Int())
						var _781_v39 _dafny.Set
						_ = _781_v39
						_781_v39 = _dafny.SetOf(p0)
						_781_v39 = _781_v39
						r1 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_742_v17).Cardinality()), (_743_v18).Cardinality())), _dafny.IntOfInt64(15))
						var _782_v40 *C0
						_ = _782_v40
						var _nw133 *C0 = New_C0_()
						_ = _nw133
						_nw133.Ctor__(_739_v14)
						_782_v40 = _nw133
					} else {
						(_this).F8_set_(((func() _dafny.Set {
							if (_this).F9() {
								return _743_v18
							}
							return _743_v18
						})()).IsSubsetOf(_743_v18))
						var _783_v41 _dafny.Set
						_ = _783_v41
						_783_v41 = _dafny.SetOf(_742_v17, _dafny.Companion_Sequence_.Concatenate(_742_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg49 _dafny.Int) interface{} {
								return coer49(arg49)
							}
						}((func(_784_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_785_i7 _dafny.Int) _dafny.Int {
								return _784_v1
							}
						})(_718_v1)))), _dafny.SeqOf((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _718_v1))
						var _786_v42 _dafny.CodePoint
						_ = _786_v42
						_786_v42 = _dafny.CodePoint('h')
						var _787_v43 _dafny.CodePoint
						_ = _787_v43
						_787_v43 = _786_v42
						var _788_v44 _dafny.Array
						_ = _788_v44
						var _nwElement0_30 _dafny.CodePoint = _dafny.CodePoint('j')
						_ = _nwElement0_30
						var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(29))
						_ = _nw134
						(_nw134).ArraySet1CodePoint(_nwElement0_30, 0)
						(_nw134).ArraySet1CodePoint((func() _dafny.CodePoint {
							if (_this).F9() {
								return _786_v42
							}
							return (_787_v43)
						})(), 1)
						(_nw134).ArraySet1CodePoint(_786_v42, 2)
						(_nw134).ArraySet1CodePoint(_786_v42, 3)
						(_nw134).ArraySet1CodePoint(_786_v42, 4)
						(_nw134).ArraySet1CodePoint(_786_v42, 5)
						(_nw134).ArraySet1CodePoint(_786_v42, 6)
						(_nw134).ArraySet1CodePoint(_786_v42, 7)
						(_nw134).ArraySet1CodePoint(_786_v42, 8)
						(_nw134).ArraySet1CodePoint(_dafny.CodePoint('n'), 9)
						(_nw134).ArraySet1CodePoint(_786_v42, 10)
						(_nw134).ArraySet1CodePoint(((_780_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.IntOfUint32(((_780_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint), 11)
						(_nw134).ArraySet1CodePoint(_786_v42, 12)
						(_nw134).ArraySet1CodePoint(_786_v42, 13)
						(_nw134).ArraySet1CodePoint(_786_v42, 14)
						(_nw134).ArraySet1CodePoint(_786_v42, 15)
						(_nw134).ArraySet1CodePoint(_786_v42, 16)
						(_nw134).ArraySet1CodePoint(_786_v42, 17)
						(_nw134).ArraySet1CodePoint(Companion_Default___.Fm7((_this).Fm6(_718_v1, (Companion_Default___.Fm18(_this.F8(), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), (_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), globalState)).Cardinality(), globalState), (_this).Fm1((_780_v38).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_780_v38), 0))).Int()).(_dafny.Sequence), _747_v22, globalState), _718_v1, globalState), 18)
						(_nw134).ArraySet1CodePoint(_786_v42, 19)
						(_nw134).ArraySet1CodePoint(_786_v42, 20)
						(_nw134).ArraySet1CodePoint(Companion_Default___.Fm7((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), p1, (_this).Fm2(Companion_D0_.Create_DC1_(false, _dafny.SeqOf((_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)), _this.F8()), globalState), globalState), 21)
						(_nw134).ArraySet1CodePoint((_786_v42), 22)
						(_nw134).ArraySet1CodePoint(_786_v42, 23)
						(_nw134).ArraySet1CodePoint((_787_v43), 24)
						(_nw134).ArraySet1CodePoint(_786_v42, 25)
						(_nw134).ArraySet1CodePoint(_dafny.CodePoint('q'), 26)
						(_nw134).ArraySet1CodePoint(Companion_Default___.Fm7(_dafny.IntOfInt64(254), p1, (_dafny.Zero).Minus(_718_v1), globalState), 27)
						(_nw134).ArraySet1CodePoint((func() _dafny.CodePoint {
							if p1 {
								return _786_v42
							}
							return _786_v42
						})(), 28)
						_788_v44 = _nw134
						var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_788_v44), 0))
						_ = _index145
						(_788_v44).ArraySet1CodePoint(_786_v42, (_index145).Int())
						var _789_v45 D8
						_ = _789_v45
						_789_v45 = Companion_D8_.Create_DC19_(p1)
						var _790_v46 _dafny.Array
						_ = _790_v46
						var _nw135 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(10))
						_ = _nw135
						_790_v46 = _nw135
						var _791_v47 _dafny.Map
						_ = _791_v47
						_791_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_789_v45, _790_v46)
						var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_788_v44), 0))
						_ = _index146
						var _rhs129 bool = !(true)
						_ = _rhs129
						var _rhs130 _dafny.Set = ((_dafny.SetOf(_742_v17, _dafny.Companion_Sequence_.Update(_742_v17, (Companion_Default___.SafeIndex((_717_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_717_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_742_v17).Cardinality()))).Uint32(), _dafny.IntOfInt64(858)), _742_v17)).Difference(_783_v41)).Difference((_783_v41).Intersection(_783_v41))
						_ = _rhs130
						var _rhs131 _dafny.CodePoint = _dafny.CodePoint('t')
						_ = _rhs131
						var _rhs132 _dafny.Map = (_791_v47).Update(func(_pat_let31_0 D8) D8 {
							return func(_792_dt__update__tmp_h1 D8) D8 {
								return func(_pat_let32_0 bool) D8 {
									return func(_793_dt__update_hcf31_h0 bool) D8 {
										return Companion_D8_.Create_DC19_(_793_dt__update_hcf31_h0)
									}(_pat_let32_0)
								}((_this).F9())
							}(_pat_let31_0)
						}(_789_v45), _790_v46)
						_ = _rhs132
						var _lhs74 *C2 = _this
						_ = _lhs74
						var _lhs75 _dafny.Array = _788_v44
						_ = _lhs75
						var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_788_v44), 0))
						_ = _lhs76
						_lhs74.F8_set_(_rhs129)
						_783_v41 = _rhs130
						(_lhs75).ArraySet1CodePoint(_rhs131, (_lhs76).Int())
						_791_v47 = _rhs132
						var _794_v48 *C0
						_ = _794_v48
						var _nw136 *C0 = New_C0_()
						_ = _nw136
						_nw136.Ctor__((_740_v15).F12())
						_794_v48 = _nw136
						(_this).F8_set_((_752_v27).IsSubsetOf(_752_v27))
						_718_v1 = (_this).Fm2(_747_v22, globalState)
					}
					_717_v0 = _717_v0
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _795_v49 *C0
		_ = _795_v49
		var _nw137 *C0 = New_C0_()
		_ = _nw137
		_nw137.Ctor__((_740_v15).F12())
		_795_v49 = _nw137
		r0 = p0
		r1 = _dafny.IntOfUint32((_746_v21).Cardinality())
		r2 = (_748_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_748_v23), 0))).Int()).(bool)
		return r0, r1, r2
	}
}
func (_this *C2) M3(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _796_v0 _dafny.Map
		_ = _796_v0
		_796_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), (_this).F9())
		var _797_v1 D5
		_ = _797_v1
		_797_v1 = Companion_D5_.Create_DC11_((_796_v0).Cardinality(), p0)
		var _798_v2 D5
		_ = _798_v2
		_798_v2 = Companion_D5_.Create_DC11_(p1, _dafny.Companion_Sequence_.Update((_797_v1).Dtor_cf19(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32(((_797_v1).Dtor_cf19()).Cardinality()))).Uint32(), Companion_Default___.Fm7((_dafny.Zero).Minus(p1), _this.F8(), p1, globalState)))
		_797_v1 = _798_v2
		var _799_v3 _dafny.MultiSet
		_ = _799_v3
		_799_v3 = _dafny.MultiSetOf(_dafny.MultiSetOf((_this).F9(), _this.F8()))
		var _hi6 _dafny.Int = (func() _dafny.Int {
			if (_this).F9() {
				return p1
			}
			return p1
		})()
		_ = _hi6
		for _800_i0 := (_799_v3).Cardinality(); _800_i0.Cmp(_hi6) < 0; _800_i0 = _800_i0.Plus(_dafny.One) {
			var _801_v4 _dafny.MultiSet
			_ = _801_v4
			_801_v4 = _dafny.MultiSetOf(p0, p0)
			var _802_v5 _dafny.Array
			_ = _802_v5
			var _nwElement0_31 bool = (_dafny.MultiSetFromSeq(_dafny.SeqOf(p0, p0, p0, Companion_Default___.Fm9(_this.F8(), !((_this).F9()), globalState)))).IsDisjointFrom(_801_v4)
			_ = _nwElement0_31
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(9))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_31, 0)
			(_nw138).ArraySet1((_this).F9(), 1)
			(_nw138).ArraySet1((_this).F9(), 2)
			(_nw138).ArraySet1((false) || (_this.F8()), 3)
			(_nw138).ArraySet1((_this).F9(), 4)
			(_nw138).ArraySet1(_this.F8(), 5)
			(_nw138).ArraySet1(_this.F8(), 6)
			(_nw138).ArraySet1(_this.F8(), 7)
			(_nw138).ArraySet1((_this).F9(), 8)
			_802_v5 = _nw138
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))
			_ = _index147
			(_802_v5).ArraySet1(_this.F8(), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))
			_ = _index148
			(_802_v5).ArraySet1(!_dafny.Companion_Sequence_.Equal(p0, _dafny.UnicodeSeqOfUtf8Bytes("kggm")), (_index148).Int())
			var _803_v6 _dafny.MultiSet
			_ = _803_v6
			_803_v6 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.IsPrefixOf(p0, p0), _this.F8(), _this.F8())
			var _804_v7 _dafny.Sequence
			_ = _804_v7
			_804_v7 = _dafny.SeqOf((_802_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))).Int()).(bool))
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))
			_ = _index149
			var _rhs133 _dafny.Int = _800_i0
			_ = _rhs133
			var _rhs134 _dafny.Int = (_803_v6).Cardinality()
			_ = _rhs134
			var _rhs135 bool = !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_802_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))).Int()).(bool)), _dafny.Companion_Sequence_.Concatenate(_804_v7, _804_v7))
			_ = _rhs135
			var _lhs77 _dafny.Array = _802_v5
			_ = _lhs77
			var _lhs78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))
			_ = _lhs78
			r0 = _rhs133
			r0 = _rhs134
			(_lhs77).ArraySet1(_rhs135, (_lhs78).Int())
			var _805_v8 _dafny.CodePoint
			_ = _805_v8
			_805_v8 = _dafny.CodePoint('o')
			_805_v8 = _805_v8
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_802_v5), 0))
			_ = _index150
			(_802_v5).ArraySet1((_this).F9(), (_index150).Int())
		}
		var _806_v9 _dafny.MultiSet
		_ = _806_v9
		_806_v9 = _dafny.MultiSetOf(p1)
		var _807_i1 _dafny.Int
		_ = _807_i1
		_807_i1 = _dafny.Zero
		{
			for ((_806_v9).Difference(_806_v9)).IsSubsetOf((_806_v9).Update(p1, Companion_Default___.Abs(p1))) {
				{
					if (_807_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_807_i1 = (_807_i1).Plus(_dafny.One)
					r1 = !(_this.F8())
					var _808_v10 _dafny.Sequence
					_ = _808_v10
					_808_v10 = _dafny.UnicodeSeqOfUtf8Bytes("xkd")
					var _809_v11 _dafny.Sequence
					_ = _809_v11
					_809_v11 = _dafny.SeqOf(p1)
					var _810_v12 D3
					_ = _810_v12
					_810_v12 = Companion_D3_.Create_DC5_(_809_v11)
					var _811_v13 _dafny.Array
					_ = _811_v13
					var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
					_ = _nw139
					_811_v13 = _nw139
					var _812_v14 _dafny.Array
					_ = _812_v14
					var _len0_13 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_13
					var _nw140 _dafny.Array
					_ = _nw140
					if _len0_13.Cmp(_dafny.Zero) == 0 {
						_nw140 = _dafny.NewArray(_len0_13)
					} else {
						var _init13 func(_dafny.Int) _dafny.Int = func(_813_i2 _dafny.Int) _dafny.Int {
							return (_813_i2).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), _this.F8())).Cardinality()))
						}
						_ = _init13
						var _element0_13 = _init13(_dafny.Zero)
						_ = _element0_13
						_nw140 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
						(_nw140).ArraySet1(_element0_13, 0)
						var _nativeLen0_13 = (_len0_13).Int()
						_ = _nativeLen0_13
						for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
							(_nw140).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
						}
					}
					_812_v14 = _nw140
					var _814_v15 _dafny.Sequence
					_ = _814_v15
					_814_v15 = _dafny.SeqOf(_812_v14, _812_v14)
					var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_811_v13), 0))
					_ = _index151
					(_811_v13).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_814_v15, _814_v15), (_index151).Int())
					var _815_v16 _dafny.CodePoint
					_ = _815_v16
					_815_v16 = _dafny.CodePoint('g')
					var _816_v17 _dafny.Map
					_ = _816_v17
					_816_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _815_v16)
					var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_811_v13), 0))
					_ = _index152
					var _rhs136 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("rdheeg"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdheeg")).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
						if (_816_v17).Contains(_this.F8()) {
							return (_816_v17).Get(_this.F8()).(_dafny.CodePoint)
						}
						return _dafny.CodePoint('c')
					})())
					_ = _rhs136
					var _rhs137 D3 = _810_v12
					_ = _rhs137
					var _rhs138 _dafny.Sequence = _814_v15
					_ = _rhs138
					var _rhs139 _dafny.Int = p1
					_ = _rhs139
					var _rhs140 bool = (_this).Fm0(globalState)
					_ = _rhs140
					var _lhs79 _dafny.Array = _811_v13
					_ = _lhs79
					var _lhs80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_811_v13), 0))
					_ = _lhs80
					var _lhs81 *C2 = _this
					_ = _lhs81
					_808_v10 = _rhs136
					_810_v12 = _rhs137
					(_lhs79).ArraySet1(_rhs138, (_lhs80).Int())
					r0 = _rhs139
					_lhs81.F8_set_(_rhs140)
					var _817_v18 _dafny.Array
					_ = _817_v18
					var _len0_14 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_14
					var _nw141 _dafny.Array
					_ = _nw141
					if _len0_14.Cmp(_dafny.Zero) == 0 {
						_nw141 = _dafny.NewArray(_len0_14)
					} else {
						var _init14 func(_dafny.Int) _dafny.Map = (func(_818_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_819_i3 _dafny.Int) _dafny.Map {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_818_p1, _dafny.IntOfInt64(215))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_818_p1, _818_p1))
							}
						})(p1)
						_ = _init14
						var _element0_14 = _init14(_dafny.Zero)
						_ = _element0_14
						_nw141 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
						(_nw141).ArraySet1(_element0_14, 0)
						var _nativeLen0_14 = (_len0_14).Int()
						_ = _nativeLen0_14
						for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
							(_nw141).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
						}
					}
					_817_v18 = _nw141
					_817_v18 = _817_v18
					var _820_v19 D4
					_ = _820_v19
					_820_v19 = Companion_D4_.Create_DC7_(_808_v10)
					var _821_v20 _dafny.Sequence
					_ = _821_v20
					_821_v20 = _dafny.SeqOf(_820_v19, _820_v19)
					var _822_v21 _dafny.Set
					_ = _822_v21
					_822_v21 = _dafny.SetOf(_821_v20, _821_v20)
					var _823_v22 T2
					_ = _823_v22
					var _nw142 *C1 = New_C1_()
					_ = _nw142
					_nw142.Ctor__(_815_v16, _822_v21, _this.F8(), (_this).F9())
					_823_v22 = _nw142
					var _824_v23 _dafny.Map
					_ = _824_v23
					_824_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v22, p1)
					var _825_v24 _dafny.Sequence
					_ = _825_v24
					_825_v24 = _dafny.SeqOf((_824_v23).Merge(_824_v23), (_824_v23).Update(_823_v22, p1), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v22, p1), _824_v23, _824_v23)
					_824_v23 = (_825_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg50 _dafny.Int) interface{} {
							return coer50(arg50)
						}
					}((func(_826_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_827_i4 _dafny.Int) _dafny.Int {
							return _826_p1
						}
					})(p1)))).Cardinality()), _dafny.IntOfUint32((_825_v24).Cardinality()))).Uint32()).(_dafny.Map)
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _828_v25 _dafny.Array
		_ = _828_v25
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_15
		var _nw143 _dafny.Array
		_ = _nw143
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw143 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) _dafny.CodePoint = func(_829_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			}
			_ = _init15
			var _element0_15 = _init15(_dafny.Zero)
			_ = _element0_15
			_nw143 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
			(_nw143).ArraySet1CodePoint(_element0_15, 0)
			var _nativeLen0_15 = (_len0_15).Int()
			_ = _nativeLen0_15
			for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
				(_nw143).ArraySet1CodePoint(_init15(_dafny.IntOf(_i0_15)), _i0_15)
			}
		}
		_828_v25 = _nw143
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_828_v25), 0))); ; {
			_guard_loop_2, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _830_i5 _dafny.Int
			_830_i5 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_830_i5).Sign() != -1) && ((_830_i5).Cmp(_dafny.ArrayLen((_828_v25), 0)) < 0)) {
				(_828_v25).ArraySet1CodePoint(_dafny.CodePoint('c'), (_830_i5).Int())
			}
		}
		var _831_v26 _dafny.Array
		_ = _831_v26
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw144
		_831_v26 = _nw144
		var _832_v27 *C0
		_ = _832_v27
		var _nw145 *C0 = New_C0_()
		_ = _nw145
		_nw145.Ctor__(_831_v26)
		_832_v27 = _nw145
		var _833_v28 _dafny.Map
		_ = _833_v28
		_833_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_832_v27, _832_v27, _832_v27, _832_v27, _832_v27)).Cardinality()), (_this).F9())
		_833_v28 = (_833_v28).Update(Companion_Default___.SafeModuloInt(p1, _dafny.IntOfInt64(499)), (_this).F9())
		r0 = p1
		r0 = p1
		r1 = true
		return r0, r1
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f8 bool
	_f9 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f8 = false
	_this._f9 = false
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T2 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F8() bool {
	return _this._f8
}
func (_this *C3) F8_set_(value bool) {
	_this._f8 = value
}
func (_this *C3) F9() bool {
	return _this._f9
}
func (_this *C3) Ctor__(f8 bool, f9 bool) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C3) Fm2(p0 D0, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(-710)).Times((_dafny.SetOf(true)).Cardinality())).Times(_dafny.IntOfInt64(433))
	}
}
func (_this *C3) Fm0(globalState *GlobalState) bool {
	{
		return (func() _dafny.Set {
			var _coll16 = _dafny.NewBuilder()
			_ = _coll16
			for _iter19 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-88)), _dafny.IntOfUint32((_dafny.SeqOf(_this.F8(), _this.F8())).Cardinality()))).Elements()); ; {
				_compr_16, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _834_v0 _dafny.Int
				_834_v0 = interface{}(_compr_16).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-88)), _dafny.IntOfUint32((_dafny.SeqOf(_this.F8(), _this.F8())).Cardinality())), _834_v0) {
					_coll16.Add(Companion_Default___.SafeModuloInt(_834_v0, _dafny.IntOfInt64(7)))
				}
			}
			return _coll16.ToSet()
		}()).IsProperSubsetOf(_dafny.SetOf((_dafny.MultiSetOf((_this).F9(), _this.F8())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), (func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(57), _dafny.IntOfInt64(809))); ; {
				_compr_17, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _835_v1 _dafny.Int
				_835_v1 = interface{}(_compr_17).(_dafny.Int)
				if ((_dafny.IntOfInt64(57)).Cmp(_835_v1) <= 0) && ((_835_v1).Cmp(_dafny.IntOfInt64(809)) < 0) {
					_coll17.Add((_835_v1).Times(_dafny.IntOfInt64(-161)), _this.F8())
				}
			}
			return _coll17.ToMap()
		}()).Cardinality())).Cardinality(), _dafny.IntOfInt64(944), _dafny.IntOfInt64(-645)))
	}
}
func (_this *C3) Fm1(p0 _dafny.Sequence, p1 D0, globalState *GlobalState) bool {
	{
		return (_this).F9()
	}
}
func (_this *C3) Fm3(globalState *GlobalState) _dafny.MultiSet {
	{
		return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(976), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ceyl")).Cardinality()), _dafny.IntOfInt64(563)))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_836_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(201)
		}))).Cardinality())))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(443), _dafny.IntOfInt64(200), (func() _dafny.Set {
			var _coll18 = _dafny.NewBuilder()
			_ = _coll18
			for _iter21 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(748), _dafny.IntOfInt64(142))).Elements()); ; {
				_compr_18, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _837_v0 _dafny.Int
				_837_v0 = interface{}(_compr_18).(_dafny.Int)
				if (_dafny.MultiSetOf(_dafny.IntOfInt64(748), _dafny.IntOfInt64(142))).Contains(_837_v0) {
					_coll18.Add((_837_v0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.CodePoint('o')))).Cardinality())))
				}
			}
			return _coll18.ToSet()
		}()).Cardinality())).Cardinality())))
	}
}
func (_this *C3) M0(p0 _dafny.MultiSet, p1 D0, p2 bool, globalState *GlobalState) {
	{
		var _838_v0 _dafny.Array
		_ = _838_v0
		var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
		_ = _nw146
		_838_v0 = _nw146
		var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_838_v0), 0))
		_ = _index153
		(_838_v0).ArraySet1(p0, (_index153).Int())
		var _839_v1 _dafny.Sequence
		_ = _839_v1
		_839_v1 = _dafny.SeqOf(p0, _dafny.MultiSetOf(_this.F8(), p2, p2, p2, (_this).F9()), p0, p0, p0)
		var _840_v2 _dafny.Int
		_ = _840_v2
		_840_v2 = _dafny.IntOfInt64(475)
		var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen((_838_v0), 0))
		_ = _index154
		(_838_v0).ArraySet1(((func() _dafny.MultiSet {
			if (_this).F9() {
				return p0
			}
			return _dafny.MultiSetOf(false, _this.F8())
		})()).Intersection((_839_v1).Select((Companion_Default___.SafeIndex(_840_v2, _dafny.IntOfUint32((_839_v1).Cardinality()))).Uint32()).(_dafny.MultiSet)), (_index154).Int())
		var _841_v3 _dafny.Sequence
		_ = _841_v3
		_841_v3 = _dafny.SeqOf(_840_v2)
		(_this).F8_set_(!(!((_dafny.IntOfUint32((_841_v3).Cardinality())).Cmp(_840_v2) == 0)))
		var _842_v4 _dafny.Sequence
		_ = _842_v4
		_842_v4 = _dafny.UnicodeSeqOfUtf8Bytes("ks")
		var _843_v5 D2
		_ = _843_v5
		_843_v5 = Companion_D2_.Create_DC3_((_841_v3).Select((Companion_Default___.SafeIndex(_840_v2, _dafny.IntOfUint32((_841_v3).Cardinality()))).Uint32()).(_dafny.Int))
		var _rhs141 _dafny.Sequence = Companion_Default___.Fm4(_840_v2, _840_v2, globalState)
		_ = _rhs141
		var _rhs142 _dafny.Int = (_dafny.Zero).Minus(((_843_v5).Dtor_cf4()).Times((_dafny.Zero).Minus(_840_v2)))
		_ = _rhs142
		var _rhs143 bool = _this.F8()
		_ = _rhs143
		var _lhs82 *C3 = _this
		_ = _lhs82
		_842_v4 = _rhs141
		_840_v2 = _rhs142
		_lhs82.F8_set_(_rhs143)
		var _844_v6 _dafny.Array
		_ = _844_v6
		var _len0_16 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_16
		var _nw147 _dafny.Array
		_ = _nw147
		if _len0_16.Cmp(_dafny.Zero) == 0 {
			_nw147 = _dafny.NewArray(_len0_16)
		} else {
			var _init16 func(_dafny.Int) bool = func(_845_i0 _dafny.Int) bool {
				return false
			}
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
		_844_v6 = _nw147
		_844_v6 = _844_v6
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_844_v6), 0))
		_ = _index155
		(_844_v6).ArraySet1(p2, (_index155).Int())
		var _846_v7 _dafny.Sequence
		_ = _846_v7
		_846_v7 = _dafny.SeqOf(p2)
		var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_844_v6), 0))
		_ = _index156
		(_844_v6).ArraySet1((_846_v7).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_840_v2).Plus(_840_v2)), _dafny.IntOfUint32((_846_v7).Cardinality()))).Uint32()).(bool), (_index156).Int())
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_844_v6), 0))
		_ = _index157
		var _rhs144 bool = (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fvnhbyxkp")).Cardinality())).Cmp(_840_v2) == 0
		_ = _rhs144
		var _rhs145 _dafny.Int = (_840_v2).Plus(_dafny.IntOfInt64(637))
		_ = _rhs145
		var _lhs83 _dafny.Array = _844_v6
		_ = _lhs83
		var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_844_v6), 0))
		_ = _lhs84
		(_lhs83).ArraySet1(_rhs144, (_lhs84).Int())
		_840_v2 = _rhs145
	}
}
func (_this *C3) M1(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) (_dafny.Sequence, _dafny.Int, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _847_v0 T2
		_ = _847_v0
		var _nw148 *C2 = New_C2_()
		_ = _nw148
		_nw148.Ctor__((_this).F9(), !(p1))
		_847_v0 = _nw148
		var _848_v1 _dafny.Map
		_ = _848_v1
		_848_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_v0, !(p1))
		var _849_v2 _dafny.Map
		_ = _849_v2
		_849_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_847_v0).F9(), _this.F8())
		var _850_v3 _dafny.Int
		_ = _850_v3
		_850_v3 = _dafny.IntOfInt64(995)
		var _851_v4 _dafny.MultiSet
		_ = _851_v4
		_851_v4 = _dafny.MultiSetOf((_849_v2).Cardinality(), _850_v3)
		var _852_v5 _dafny.Sequence
		_ = _852_v5
		_852_v5 = _dafny.SeqOf(_851_v4)
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_852_v5).Cardinality())
		_ = _hi7
		for _853_i0 := (_848_v1).Cardinality(); _853_i0.Cmp(_hi7) < 0; _853_i0 = _853_i0.Plus(_dafny.One) {
			var _854_v6 _dafny.Sequence
			_ = _854_v6
			_854_v6 = _dafny.SeqOf(_this.F8(), true, false)
			var _855_v7 _dafny.Array
			_ = _855_v7
			var _nwElement0_32 bool = (_this).F9()
			_ = _nwElement0_32
			var _nw149 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(9))
			_ = _nw149
			(_nw149).ArraySet1(_nwElement0_32, 0)
			(_nw149).ArraySet1((_this).F9(), 1)
			(_nw149).ArraySet1(false, 2)
			(_nw149).ArraySet1((_847_v0).F9(), 3)
			(_nw149).ArraySet1((_854_v6).Select((Companion_Default___.SafeIndex(_850_v3, _dafny.IntOfUint32((_854_v6).Cardinality()))).Uint32()).(bool), 4)
			(_nw149).ArraySet1((_this).Fm0(globalState), 5)
			(_nw149).ArraySet1((func() bool {
				if (_849_v2).Contains((_847_v0).F9()) {
					return (_849_v2).Get((_847_v0).F9()).(bool)
				}
				return p1
			})(), 6)
			(_nw149).ArraySet1(p1, 7)
			(_nw149).ArraySet1(true, 8)
			_855_v7 = _nw149
			var _856_v8 _dafny.Map
			_ = _856_v8
			_856_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_855_v7, p0)
			_856_v8 = (_856_v8).Update(_855_v7, p0)
			r1 = (_853_i0).Times(_853_i0)
			var _857_v9 _dafny.Sequence
			_ = _857_v9
			_857_v9 = _dafny.SeqOf(_dafny.IntOfInt64(875), _850_v3)
			var _858_v10 D3
			_ = _858_v10
			_858_v10 = Companion_D3_.Create_DC5_(_857_v9)
			var _source16 D3 = _858_v10
			_ = _source16
			if _source16.Is_DC6() {
				var _859___mcc_h0 bool = _source16.Get_().(D3_DC6).Cf7
				_ = _859___mcc_h0
				var _860___mcc_h1 _dafny.Int = _source16.Get_().(D3_DC6).Cf8
				_ = _860___mcc_h1
				var _861___mcc_h2 _dafny.Int = _source16.Get_().(D3_DC6).Cf9
				_ = _861___mcc_h2
				var _862_cf9 _dafny.Int = _861___mcc_h2
				_ = _862_cf9
				var _863_cf8 _dafny.Int = _860___mcc_h1
				_ = _863_cf8
				var _864_cf7 bool = _859___mcc_h0
				_ = _864_cf7
				var _865_v11 _dafny.Array
				_ = _865_v11
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_17
				var _nw150 _dafny.Array
				_ = _nw150
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw150 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = func(_866_i1 _dafny.Int) _dafny.Int {
						return (_866_i1).Plus(_dafny.IntOfInt64(-678))
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw150 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw150).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw150).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_865_v11 = _nw150
				_865_v11 = _865_v11
				var _867_v12 _dafny.CodePoint
				_ = _867_v12
				_867_v12 = _dafny.CodePoint('u')
				var _868_v13 _dafny.Map
				_ = _868_v13
				_868_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(532), _867_v12)
				_868_v13 = (_868_v13).Update(_850_v3, _867_v12)
				var _869_v14 _dafny.Map
				_ = _869_v14
				_869_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_863_cf8, !(_this.F8()))
				var _870_v15 D0
				_ = _870_v15
				_870_v15 = Companion_D0_.Create_DC1_((func() bool {
					if (_869_v14).Contains(_dafny.IntOfUint32((p0).Cardinality())) {
						return (_869_v14).Get(_dafny.IntOfUint32((p0).Cardinality())).(bool)
					}
					return p1
				})(), _854_v6, (_847_v0).F9())
				_863_cf8 = (_this).Fm2(_870_v15, globalState)
				var _871_v16 _dafny.MultiSet
				_ = _871_v16
				_871_v16 = _dafny.MultiSetOf((_847_v0).F9(), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(287))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_872_v2 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_873_i2 _dafny.Int) _dafny.Map {
						return _872_v2
					}
				})(_849_v2)))).Cardinality())).Cmp(_dafny.IntOfUint32((p0).Cardinality())) == 0)
				_871_v16 = _871_v16
			} else {
				var _874___mcc_h3 _dafny.Sequence = _source16.Get_().(D3_DC5).Cf6
				_ = _874___mcc_h3
				var _875_cf6 _dafny.Sequence = _874___mcc_h3
				_ = _875_cf6
				var _876_v17 _dafny.Map
				_ = _876_v17
				_876_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_v0, _857_v9)
				r2 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_875_cf6, (func() _dafny.Sequence {
					if (_876_v17).Contains(_847_v0) {
						return (_876_v17).Get(_847_v0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-723))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg53 _dafny.Int) interface{} {
							return coer53(arg53)
						}
					}((func(_877_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_878_i3 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_877_i0)
						}
					})(_853_i0)))
				})()), _857_v9)
				_850_v3 = Companion_Default___.SafeModuloInt(((_dafny.MultiSetOf(_853_i0)).Union(_851_v4)).Cardinality(), _853_i0)
				var _879_v18 _dafny.MultiSet
				_ = _879_v18
				_879_v18 = _dafny.MultiSetOf(p1, _this.F8(), _this.F8(), (_this).F9(), _this.F8())
				r1 = (_dafny.Zero).Minus((Companion_D2_.Create_DC3_((_857_v9).Select((Companion_Default___.SafeIndex((_879_v18).Cardinality(), _dafny.IntOfUint32((_857_v9).Cardinality()))).Uint32()).(_dafny.Int))).Dtor_cf4())
				var _880_v19 _dafny.Set
				_ = _880_v19
				_880_v19 = _dafny.SetOf(true, (_847_v0).F9())
				var _881_v20 D7
				_ = _881_v20
				_881_v20 = Companion_D7_.Create_DC14_(_880_v19)
				var _882_v21 D7
				_ = _882_v21
				_882_v21 = Companion_D7_.Create_DC16_(_881_v20)
				_882_v21 = _882_v21
			}
			r2 = (_this).F9()
		}
		var _883_v22 _dafny.Array
		_ = _883_v22
		var _len0_18 _dafny.Int = _dafny.One
		_ = _len0_18
		var _nw151 _dafny.Array
		_ = _nw151
		if _len0_18.Cmp(_dafny.Zero) == 0 {
			_nw151 = _dafny.NewArray(_len0_18)
		} else {
			var _init18 func(_dafny.Int) _dafny.Int = (func(_884_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_885_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_885_i4, _884_v3)
				}
			})(_850_v3)
			_ = _init18
			var _element0_18 = _init18(_dafny.Zero)
			_ = _element0_18
			_nw151 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
			(_nw151).ArraySet1(_element0_18, 0)
			var _nativeLen0_18 = (_len0_18).Int()
			_ = _nativeLen0_18
			for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
				(_nw151).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
			}
		}
		_883_v22 = _nw151
		var _886_v23 _dafny.Sequence
		_ = _886_v23
		_886_v23 = _dafny.SeqOf(_883_v22, _883_v22)
		var _887_v24 _dafny.Sequence
		_ = _887_v24
		_887_v24 = _dafny.SeqOf(_886_v23, _886_v23)
		_886_v23 = (_887_v24).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if _this.F8() {
				return _850_v3
			}
			return _850_v3
		})(), _dafny.IntOfUint32((_887_v24).Cardinality()))).Uint32()).(_dafny.Sequence)
		r1 = _dafny.IntOfInt64(292)
		r1 = (_dafny.Zero).Minus(_850_v3)
		var _888_v25 _dafny.Sequence
		_ = _888_v25
		_888_v25 = _dafny.SeqOf((_this).F9(), (_847_v0).F9())
		var _889_v26 D0
		_ = _889_v26
		_889_v26 = Companion_D0_.Create_DC1_(_this.F8(), _888_v25, p1)
		var _890_v27 _dafny.Array
		_ = _890_v27
		var _nwElement0_33 _dafny.Array = _883_v22
		_ = _nwElement0_33
		var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(17))
		_ = _nw152
		(_nw152).ArraySet1(_nwElement0_33, 0)
		(_nw152).ArraySet1(_883_v22, 1)
		(_nw152).ArraySet1(_883_v22, 2)
		(_nw152).ArraySet1(_883_v22, 3)
		(_nw152).ArraySet1(_883_v22, 4)
		(_nw152).ArraySet1(_883_v22, 5)
		(_nw152).ArraySet1(_883_v22, 6)
		(_nw152).ArraySet1(_883_v22, 7)
		(_nw152).ArraySet1(_883_v22, 8)
		(_nw152).ArraySet1(_883_v22, 9)
		(_nw152).ArraySet1(_883_v22, 10)
		(_nw152).ArraySet1(_883_v22, 11)
		(_nw152).ArraySet1(_883_v22, 12)
		(_nw152).ArraySet1(_883_v22, 13)
		(_nw152).ArraySet1(_883_v22, 14)
		(_nw152).ArraySet1(_883_v22, 15)
		(_nw152).ArraySet1(_883_v22, 16)
		_890_v27 = _nw152
		var _891_v28 *C0
		_ = _891_v28
		var _nw153 *C0 = New_C0_()
		_ = _nw153
		_nw153.Ctor__(_890_v27)
		_891_v28 = _nw153
		var _892_v29 _dafny.MultiSet
		_ = _892_v29
		_892_v29 = _dafny.MultiSetOf(_891_v28)
		var _893_v30 D2
		_ = _893_v30
		_893_v30 = Companion_D2_.Create_DC4_(p1)
		var _894_v31 _dafny.MultiSet
		_ = _894_v31
		_894_v31 = _dafny.MultiSetOf(Companion_D5_.Create_DC10_(_883_v22))
		var _895_v32 _dafny.Array
		_ = _895_v32
		var _nwElement0_34 bool = _this.F8()
		_ = _nwElement0_34
		var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(11))
		_ = _nw154
		(_nw154).ArraySet1(_nwElement0_34, 0)
		(_nw154).ArraySet1((func() bool {
			if (_this).F9() {
				return _this.F8()
			}
			return _this.F8()
		})(), 1)
		(_nw154).ArraySet1((_847_v0).F9(), 2)
		(_nw154).ArraySet1((func() bool {
			if (_this).Fm1(_dafny.UnicodeSeqOfUtf8Bytes("uk"), _889_v26, globalState) {
				return !((_this).F9())
			}
			return (_this).Fm1(p0, _889_v26, globalState)
		})(), 3)
		(_nw154).ArraySet1((_847_v0).F9(), 4)
		(_nw154).ArraySet1(_this.F8(), 5)
		(_nw154).ArraySet1((_892_v29).IsSubsetOf(_892_v29), 6)
		(_nw154).ArraySet1((_this).F9(), 7)
		(_nw154).ArraySet1(!((_893_v30).Dtor_cf5()), 8)
		(_nw154).ArraySet1(_this.F8(), 9)
		(_nw154).ArraySet1(((_894_v31).Cardinality()).Cmp(_dafny.IntOfInt64(653)) != 0, 10)
		_895_v32 = _nw154
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_895_v32), 0))
		_ = _index158
		(_895_v32).ArraySet1(_this.F8(), (_index158).Int())
		var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_895_v32), 0))
		_ = _index159
		(_895_v32).ArraySet1(true, (_index159).Int())
		(_this).F8_set_((_this).F9())
		r0 = p0
		var _896_v33 _dafny.MultiSet
		_ = _896_v33
		_896_v33 = _dafny.MultiSetOf(_this.F8())
		r1 = (_850_v3).Minus((func() _dafny.Int {
			if (_896_v33).Contains(p1) {
				return (_896_v33).Multiplicity(p1)
			}
			return _850_v3
		})())
		r2 = (((_this).F9()) == ((_891_v28).Fm0(globalState))) == ((_850_v3).Cmp(_dafny.IntOfInt64(154)) >= 0)
		return r0, r1, r2
	}
}
func (_this *C3) M2(p0 bool, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _897_v0 _dafny.Int
		_ = _897_v0
		_897_v0 = _dafny.IntOfInt64(-16)
		var _hi8 _dafny.Int = _897_v0
		_ = _hi8
		for _898_i0 := _897_v0; _898_i0.Cmp(_hi8) < 0; _898_i0 = _898_i0.Plus(_dafny.One) {
			var _899_v1 _dafny.Array
			_ = _899_v1
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
			_ = _nw155
			_899_v1 = _nw155
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_899_v1), 0))
			_ = _index160
			(_899_v1).ArraySet1((_this).F9(), (_index160).Int())
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_899_v1), 0))
			_ = _index161
			(_899_v1).ArraySet1(p0, (_index161).Int())
			var _900_v2 _dafny.Sequence
			_ = _900_v2
			_900_v2 = _dafny.UnicodeSeqOfUtf8Bytes("ikpfgysl")
			var _901_v3 _dafny.Array
			_ = _901_v3
			var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
			_ = _nw156
			_901_v3 = _nw156
			var _902_v4 _dafny.Map
			_ = _902_v4
			_902_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_900_v2, _901_v3)
			_902_v4 = (_902_v4).Update(_900_v2, _901_v3)
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_899_v1), 0))
			_ = _index162
			(_899_v1).ArraySet1(p0, (_index162).Int())
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_899_v1), 0))
			_ = _index163
			(_899_v1).ArraySet1((_899_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_899_v1), 0))).Int()).(bool), (_index163).Int())
		}
		if (_this).F9() {
			var _903_v5 _dafny.Array
			_ = _903_v5
			var _nw157 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw157
			_903_v5 = _nw157
			var _904_v6 _dafny.Array
			_ = _904_v6
			var _nwElement0_35 bool = p0
			_ = _nwElement0_35
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(21))
			_ = _nw158
			(_nw158).ArraySet1(_nwElement0_35, 0)
			(_nw158).ArraySet1(!(p0), 1)
			(_nw158).ArraySet1(_this.F8(), 2)
			(_nw158).ArraySet1(false, 3)
			(_nw158).ArraySet1((_this).F9(), 4)
			(_nw158).ArraySet1((_this).F9(), 5)
			(_nw158).ArraySet1(_this.F8(), 6)
			(_nw158).ArraySet1(p0, 7)
			(_nw158).ArraySet1(_this.F8(), 8)
			(_nw158).ArraySet1(false, 9)
			(_nw158).ArraySet1(false, 10)
			(_nw158).ArraySet1(p0, 11)
			(_nw158).ArraySet1((_this).F9(), 12)
			(_nw158).ArraySet1(!((_this).F9()), 13)
			(_nw158).ArraySet1(true, 14)
			(_nw158).ArraySet1(_this.F8(), 15)
			(_nw158).ArraySet1((_this).F9(), 16)
			(_nw158).ArraySet1(!(true), 17)
			(_nw158).ArraySet1(_this.F8(), 18)
			(_nw158).ArraySet1(_this.F8(), 19)
			(_nw158).ArraySet1(_this.F8(), 20)
			_904_v6 = _nw158
			var _905_v7 _dafny.Sequence
			_ = _905_v7
			_905_v7 = _dafny.SeqOf(_904_v6)
			var _906_v8 _dafny.Map
			_ = _906_v8
			_906_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v5, _dafny.Companion_Sequence_.Contains(_905_v7, _904_v6))
			var _907_v9 _dafny.Set
			_ = _907_v9
			_907_v9 = _dafny.SetOf(p0, p0)
			var _908_v10 D4
			_ = _908_v10
			_908_v10 = Companion_D4_.Create_DC8_(_897_v0, (_this).F9())
			var _909_v11 _dafny.Sequence
			_ = _909_v11
			_909_v11 = _dafny.SeqOf(_908_v10, _908_v10, _908_v10, _908_v10)
			_906_v8 = (_906_v8).Update(_903_v5, ((_907_v9).Cardinality()).Cmp(_dafny.IntOfUint32((_909_v11).Cardinality())) != 0)
			var _910_v12 _dafny.Map
			_ = _910_v12
			_910_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_897_v0, _897_v0)
			var _911_v13 _dafny.Map
			_ = _911_v13
			_911_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _897_v0)
			var _912_v14 _dafny.Set
			_ = _912_v14
			_912_v14 = _dafny.SetOf((func() _dafny.Int {
				if (_910_v12).Contains(_897_v0) {
					return (_910_v12).Get(_897_v0).(_dafny.Int)
				}
				return (_911_v13).Cardinality()
			})(), _dafny.IntOfInt64(-441), _897_v0)
			if (_dafny.SetOf(_dafny.IntOfInt64(498))).IsSubsetOf(_912_v14) {
				_897_v0 = (_dafny.Zero).Minus(_897_v0)
				var _913_v15 _dafny.Sequence
				_ = _913_v15
				_913_v15 = _dafny.UnicodeSeqOfUtf8Bytes("niqye")
				var _914_v16 _dafny.Sequence
				_ = _914_v16
				_914_v16 = _dafny.SeqOf(p0, (_this).F9())
				var _915_v17 D0
				_ = _915_v17
				_915_v17 = Companion_D0_.Create_DC1_((_this).F9(), _914_v16, true)
				var _916_v18 *C2
				_ = _916_v18
				var _nw159 *C2 = New_C2_()
				_ = _nw159
				_nw159.Ctor__(((_this).F9()) || ((_this).Fm1(_913_v15, _915_v17, globalState)), false)
				_916_v18 = _nw159
				_916_v18 = _916_v18
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_903_v5), 0))
				_ = _index164
				(_903_v5).ArraySet1(_897_v0, (_index164).Int())
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_903_v5), 0))
				_ = _index165
				(_903_v5).ArraySet1(_897_v0, (_index165).Int())
				var _917_v19 _dafny.Sequence
				_ = _917_v19
				_917_v19 = _dafny.SeqOf((_903_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_903_v5), 0))).Int()).(_dafny.Int))
				_917_v19 = _917_v19
				r0 = (_903_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_903_v5), 0))).Int()).(_dafny.Int)
			} else {
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_903_v5), 0))
				_ = _index166
				(_903_v5).ArraySet1(_897_v0, (_index166).Int())
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_903_v5), 0))
				_ = _index167
				(_903_v5).ArraySet1(_897_v0, (_index167).Int())
				var _918_v20 _dafny.CodePoint
				_ = _918_v20
				_918_v20 = _dafny.CodePoint('m')
				var _919_v21 _dafny.Sequence
				_ = _919_v21
				_919_v21 = _dafny.SeqOf(_918_v20, _918_v20)
				var _920_v23 *C1
				_ = _920_v23
				var _nw160 *C1 = New_C1_()
				_ = _nw160
				_nw160.Ctor__((_919_v21).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_897_v0), _dafny.IntOfUint32((_919_v21).Cardinality()))).Uint32()).(_dafny.CodePoint), func() _dafny.Set {
					var _coll19 = _dafny.NewBuilder()
					_ = _coll19
					for _iter22 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer54 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}(func(_921_i1 _dafny.Int) D4 {
						return Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ymfwvwo"))
					})))).Elements()); ; {
						_compr_19, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _922_v22 _dafny.Sequence
						_922_v22 = interface{}(_compr_19).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer55 func(_dafny.Int) D4) func(_dafny.Int) interface{} {
							return func(arg55 _dafny.Int) interface{} {
								return coer55(arg55)
							}
						}(func(_921_i1 _dafny.Int) D4 {
							return Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ymfwvwo"))
						})))).Contains(_922_v22) {
							_coll19.Add(_922_v22)
						}
					}
					return _coll19.ToSet()
				}(), p0, !(!((_this).F9()) || (_this.F8())))
				_920_v23 = _nw160
				var _923_v24 _dafny.Array
				_ = _923_v24
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_19
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Int = (func(_924_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_925_i2 _dafny.Int) _dafny.Int {
							return (_925_i2).Minus((_924_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_924_v5), 0))).Int()).(_dafny.Int))
						}
					})(_903_v5)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw161 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw161).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw161).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_923_v24 = _nw161
				var _926_v25 _dafny.MultiSet
				_ = _926_v25
				_926_v25 = _dafny.MultiSetOf((_this).F9(), (_this).F9(), (_this).F9(), !((func() bool {
					if (_906_v8).Contains(_923_v24) {
						return (_906_v8).Get(_923_v24).(bool)
					}
					return p0
				})()), _this.F8())
				(_920_v23).M0(_926_v25, Companion_D0_.Create_DC0_(), (func() bool {
					if (_this).F9() {
						return p0
					}
					return p0
				})(), globalState)
				r1 = (_897_v0).Cmp((_903_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_903_v5), 0))).Int()).(_dafny.Int)) >= 0
				var _927_v26 _dafny.Array
				_ = _927_v26
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
				_ = _nw162
				_927_v26 = _nw162
				var _928_v27 _dafny.Array
				_ = _928_v27
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_20
				var _nw163 _dafny.Array
				_ = _nw163
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw163 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.CodePoint = (func(_929_v23 *C1) func(_dafny.Int) _dafny.CodePoint {
						return func(_930_i3 _dafny.Int) _dafny.CodePoint {
							return (_929_v23).F10()
						}
					})(_920_v23)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw163 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw163).ArraySet1CodePoint(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw163).ArraySet1CodePoint(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_928_v27 = _nw163
				var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_927_v26), 0))
				_ = _index168
				(_927_v26).ArraySet1(_928_v27, (_index168).Int())
				var _931_v28 _dafny.Sequence
				_ = _931_v28
				_931_v28 = _dafny.SeqOf(_this.F8())
				var _932_v29 D0
				_ = _932_v29
				_932_v29 = Companion_D0_.Create_DC1_((_this).F9(), _931_v28, p0)
				var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_927_v26), 0))
				_ = _index169
				(_927_v26).ArraySet1((func() _dafny.Array {
					if (_920_v23).Fm1(_919_v21, _932_v29, globalState) {
						return _928_v27
					}
					return _928_v27
				})(), (_index169).Int())
			}
			var _933_v30 _dafny.Array
			_ = _933_v30
			var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
			_ = _nw164
			_933_v30 = _nw164
			var _934_v31 D4
			_ = _934_v31
			_934_v31 = Companion_D4_.Create_DC9_(p0, _dafny.IntOfInt64(-727), _904_v6, _897_v0)
			var _935_v32 _dafny.Map
			_ = _935_v32
			_935_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), _904_v6)
			var _936_v33 _dafny.Array
			_ = _936_v33
			var _nwElement0_36 _dafny.Array = _904_v6
			_ = _nwElement0_36
			var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(29))
			_ = _nw165
			(_nw165).ArraySet1(_nwElement0_36, 0)
			(_nw165).ArraySet1(_904_v6, 1)
			(_nw165).ArraySet1(_904_v6, 2)
			(_nw165).ArraySet1((_934_v31).Dtor_cf15(), 3)
			(_nw165).ArraySet1((func() _dafny.Array {
				if (_935_v32).Contains((_this).F9()) {
					return (_935_v32).Get((_this).F9()).(_dafny.Array)
				}
				return _904_v6
			})(), 4)
			(_nw165).ArraySet1(_904_v6, 5)
			(_nw165).ArraySet1(_904_v6, 6)
			(_nw165).ArraySet1(_904_v6, 7)
			(_nw165).ArraySet1(_904_v6, 8)
			(_nw165).ArraySet1(_904_v6, 9)
			(_nw165).ArraySet1(_904_v6, 10)
			(_nw165).ArraySet1(_904_v6, 11)
			(_nw165).ArraySet1(_904_v6, 12)
			(_nw165).ArraySet1(_904_v6, 13)
			(_nw165).ArraySet1(_904_v6, 14)
			(_nw165).ArraySet1(_904_v6, 15)
			(_nw165).ArraySet1(_904_v6, 16)
			(_nw165).ArraySet1(_904_v6, 17)
			(_nw165).ArraySet1(_904_v6, 18)
			(_nw165).ArraySet1(_904_v6, 19)
			(_nw165).ArraySet1(_904_v6, 20)
			(_nw165).ArraySet1(_904_v6, 21)
			(_nw165).ArraySet1(_904_v6, 22)
			(_nw165).ArraySet1(_904_v6, 23)
			(_nw165).ArraySet1(_904_v6, 24)
			(_nw165).ArraySet1(_904_v6, 25)
			(_nw165).ArraySet1(_904_v6, 26)
			(_nw165).ArraySet1(_904_v6, 27)
			(_nw165).ArraySet1(_904_v6, 28)
			_936_v33 = _nw165
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_933_v30), 0))
			_ = _index170
			(_933_v30).ArraySet1(_936_v33, (_index170).Int())
			var _937_v34 _dafny.Sequence
			_ = _937_v34
			_937_v34 = _dafny.SeqOf(p0)
			var _938_v35 _dafny.MultiSet
			_ = _938_v35
			_938_v35 = _dafny.MultiSetOf(_897_v0, _897_v0)
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_933_v30), 0))
			_ = _index171
			var _rhs146 _dafny.Array = _936_v33
			_ = _rhs146
			var _rhs147 _dafny.Int = _dafny.IntOfUint32((_937_v34).Cardinality())
			_ = _rhs147
			var _rhs148 bool = ((_938_v35).Union(_938_v35)).IsProperSubsetOf((_dafny.MultiSetOf(_897_v0)).Intersection(_938_v35))
			_ = _rhs148
			var _lhs85 _dafny.Array = _933_v30
			_ = _lhs85
			var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(47), _dafny.ArrayLen((_933_v30), 0))
			_ = _lhs86
			var _lhs87 *C3 = _this
			_ = _lhs87
			(_lhs85).ArraySet1(_rhs146, (_lhs86).Int())
			r0 = _rhs147
			_lhs87.F8_set_(_rhs148)
			r0 = _897_v0
			r1 = (_this).F9()
		} else {
			var _939_v36 _dafny.CodePoint
			_ = _939_v36
			_939_v36 = _dafny.CodePoint('h')
			var _940_v37 _dafny.Sequence
			_ = _940_v37
			_940_v37 = _dafny.UnicodeSeqOfUtf8Bytes("piygtl")
			var _941_v38 _dafny.Map
			_ = _941_v38
			_941_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v36, _940_v37)
			_941_v38 = (_941_v38).Update(_939_v36, _dafny.Companion_Sequence_.Concatenate(_940_v37, _940_v37))
			var _942_v39 _dafny.Sequence
			_ = _942_v39
			_942_v39 = _dafny.SeqOf(_dafny.IntOfInt64(-155))
			var _943_v40 _dafny.Sequence
			_ = _943_v40
			_943_v40 = _dafny.SeqOf(p0)
			var _944_v41 _dafny.Map
			_ = _944_v41
			_944_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_942_v39).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_942_v39).Cardinality()))).Uint32()).(_dafny.Int), Companion_D4_.Create_DC8_(_dafny.IntOfUint32((_943_v40).Cardinality()), _this.F8()))
			var _945_v42 _dafny.MultiSet
			_ = _945_v42
			_945_v42 = _dafny.MultiSetOf(p0)
			var _946_v43 D4
			_ = _946_v43
			_946_v43 = Companion_D4_.Create_DC8_((_945_v42).Cardinality(), p0)
			_944_v41 = (_944_v41).Update(_dafny.IntOfUint32((_943_v40).Cardinality()), _946_v43)
			_897_v0 = _897_v0
			var _947_v44 _dafny.Array
			_ = _947_v44
			var _nw166 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw166
			_947_v44 = _nw166
			_947_v44 = _947_v44
			var _948_v45 _dafny.CodePoint
			_ = _948_v45
			_948_v45 = _939_v36
			_945_v42 = (_945_v42).Difference(Companion_Default___.Fm19((_this).F9(), _948_v45, (_942_v39).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_942_v39).Cardinality()))).Uint32()).(_dafny.Int), globalState))
		}
		var _949_v46 _dafny.Set
		_ = _949_v46
		_949_v46 = _dafny.SetOf(_897_v0, _897_v0)
		var _950_v47 _dafny.Sequence
		_ = _950_v47
		_950_v47 = _dafny.UnicodeSeqOfUtf8Bytes("ev")
		var _951_v48 _dafny.Sequence
		_ = _951_v48
		_951_v48 = _dafny.SeqOf(_949_v46, _dafny.SetOf(_dafny.IntOfUint32((_950_v47).Cardinality()), _897_v0), _949_v46)
		var _952_v49 D8
		_ = _952_v49
		_952_v49 = Companion_D8_.Create_DC17_((_951_v48).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_951_v48).Cardinality()))).Uint32()).(_dafny.Set))
		var _source17 D8 = _952_v49
		_ = _source17
		if _source17.Is_DC18() {
			var _953___mcc_h0 _dafny.Map = _source17.Get_().(D8_DC18).Cf28
			_ = _953___mcc_h0
			var _954___mcc_h1 bool = _source17.Get_().(D8_DC18).Cf29
			_ = _954___mcc_h1
			var _955___mcc_h2 _dafny.Int = _source17.Get_().(D8_DC18).Cf30
			_ = _955___mcc_h2
			var _956_cf30 _dafny.Int = _955___mcc_h2
			_ = _956_cf30
			var _957_cf29 bool = _954___mcc_h1
			_ = _957_cf29
			var _958_cf28 _dafny.Map = _953___mcc_h0
			_ = _958_cf28
			var _959_v50 _dafny.Map
			_ = _959_v50
			_959_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_897_v0, (_this).F9())
			var _960_v51 *C2
			_ = _960_v51
			var _nw167 *C2 = New_C2_()
			_ = _nw167
			_nw167.Ctor__(_957_cf29, _957_cf29)
			_960_v51 = _nw167
			var _961_v52 _dafny.Set
			_ = _961_v52
			_961_v52 = _dafny.SetOf(_950_v47, _950_v47)
			var _962_v53 _dafny.MultiSet
			_ = _962_v53
			_962_v53 = _dafny.MultiSetOf(_956_cf30)
			var _963_v54 D9
			_ = _963_v54
			_963_v54 = Companion_D9_.Create_DC20_(_962_v53)
			var _964_v55 D9
			_ = _964_v55
			_964_v55 = Companion_D9_.Create_DC20_(((_963_v54).Dtor_cf32()).Update(_956_cf30, Companion_Default___.Abs(_dafny.IntOfInt64(726))))
			var _rhs149 bool = !((((Companion_D9_.Create_DC20_((_962_v53).Update((_962_v53).Cardinality(), Companion_Default___.Abs(_956_cf30)))).Dtor_cf32()).Union((_964_v55).Dtor_cf32())).IsDisjointFrom((_962_v53).Intersection(_962_v53)))
			_ = _rhs149
			var _rhs150 _dafny.Map = (_959_v50).Merge((_959_v50).Update(_897_v0, (_this).F9()))
			_ = _rhs150
			var _rhs151 *C2 = (func() *C2 {
				if _957_cf29 {
					return _960_v51
				}
				return _960_v51
			})()
			_ = _rhs151
			var _rhs152 *C2 = _960_v51
			_ = _rhs152
			var _rhs153 _dafny.Set = (_961_v52).Difference(_961_v52)
			_ = _rhs153
			var _lhs88 *C3 = _this
			_ = _lhs88
			_lhs88.F8_set_(_rhs149)
			_959_v50 = _rhs150
			_960_v51 = _rhs151
			_960_v51 = _rhs152
			_961_v52 = _rhs153
			var _965_v56 _dafny.CodePoint
			_ = _965_v56
			_965_v56 = _dafny.CodePoint('c')
			var _966_v57 D4
			_ = _966_v57
			_966_v57 = Companion_D4_.Create_DC7_(_950_v47)
			var _967_v58 _dafny.Sequence
			_ = _967_v58
			_967_v58 = _dafny.SeqOf(_966_v57)
			var _968_v59 _dafny.Map
			_ = _968_v59
			_968_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_962_v53, _967_v58)
			var _pat_let_tv32 = _965_v56
			_ = _pat_let_tv32
			var _969_v60 _dafny.Set
			_ = _969_v60
			_969_v60 = _dafny.SetOf(_967_v58, (func() _dafny.Sequence {
				if (_968_v59).Contains(_962_v53) {
					return (_968_v59).Get(_962_v53).(_dafny.Sequence)
				}
				return _dafny.SeqOf(Companion_D4_.Create_DC7_(_950_v47), _966_v57, _966_v57, _966_v57, func(_pat_let33_0 D4) D4 {
					return func(_970_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let34_0 _dafny.Sequence) D4 {
							return func(_973_dt__update_hcf10_h0 _dafny.Sequence) D4 {
								return Companion_D4_.Create_DC7_(_973_dt__update_hcf10_h0)
							}(_pat_let34_0)
						}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(754))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg56 _dafny.Int) interface{} {
								return coer56(arg56)
							}
						}((func(_971_v56 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_972_i4 _dafny.Int) _dafny.CodePoint {
								return _971_v56
							}
						})(_pat_let_tv32))))
					}(_pat_let33_0)
				}(Companion_D4_.Create_DC7_(_950_v47)))
			})())
			var _974_v61 D4
			_ = _974_v61
			_974_v61 = Companion_D4_.Create_DC8_(_897_v0, true)
			var _975_v62 *C1
			_ = _975_v62
			var _nw168 *C1 = New_C1_()
			_ = _nw168
			_nw168.Ctor__(_965_v56, _969_v60, (_this).F9(), (_974_v61).Dtor_cf12())
			_975_v62 = _nw168
			_956_cf30 = _897_v0
			var _976_v63 _dafny.Array
			_ = _976_v63
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(Companion_D8_.Default(), _dafny.IntOfInt64(15))
			_ = _nw169
			_976_v63 = _nw169
			var _977_v64 _dafny.Array
			_ = _977_v64
			var _len0_21 _dafny.Int = _dafny.One
			_ = _len0_21
			var _nw170 _dafny.Array
			_ = _nw170
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw170 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Sequence = (func(_978_v0 _dafny.Int, _979_v53 _dafny.MultiSet, _980_cf30 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_981_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_978_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), (_979_v53).Cardinality())).Cardinality(), _980_cf30, _980_cf30), _dafny.SeqOf(_978_v0))
					}
				})(_897_v0, _962_v53, _956_cf30)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw170 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw170).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw170).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_977_v64 = _nw170
			var _982_v65 _dafny.Sequence
			_ = _982_v65
			_982_v65 = _dafny.SeqOf(_960_v51)
			var _rhs154 _dafny.Array = (func() _dafny.Array {
				if _957_cf29 {
					return _976_v63
				}
				return _976_v63
			})()
			_ = _rhs154
			var _rhs155 _dafny.Array = _977_v64
			_ = _rhs155
			var _rhs156 _dafny.Int = Companion_Default___.SafeDivisionInt(_897_v0, ((_962_v53).Intersection(_962_v53)).Cardinality())
			_ = _rhs156
			var _rhs157 _dafny.Int = _956_cf30
			_ = _rhs157
			var _rhs158 _dafny.Int = _dafny.IntOfUint32((_982_v65).Cardinality())
			_ = _rhs158
			_976_v63 = _rhs154
			_977_v64 = _rhs155
			_897_v0 = _rhs156
			_897_v0 = _rhs157
			_956_cf30 = _rhs158
		} else if _source17.Is_DC19() {
			var _983___mcc_h3 bool = _source17.Get_().(D8_DC19).Cf31
			_ = _983___mcc_h3
			var _984_cf31 bool = _983___mcc_h3
			_ = _984_cf31
			var _985_v66 _dafny.Sequence
			_ = _985_v66
			_985_v66 = _dafny.SeqOf(_dafny.IntOfUint32((_950_v47).Cardinality()), _897_v0, _dafny.IntOfUint32((_950_v47).Cardinality()), _897_v0, _dafny.IntOfInt64(26))
			var _986_v68 _dafny.Sequence
			_ = _986_v68
			_986_v68 = _dafny.SeqOf(p0, true)
			var _987_v70 _dafny.Array
			_ = _987_v70
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_22
			var _nw171 _dafny.Array
			_ = _nw171
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw171 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.Int = (func(_988_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_989_i6 _dafny.Int) _dafny.Int {
						return (_989_i6).Plus(_988_v0)
					}
				})(_897_v0)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw171 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw171).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw171).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_987_v70 = _nw171
			var _990_v71 _dafny.Map
			_ = _990_v71
			_990_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F8()), _dafny.SetOf(_897_v0))
			var _991_v72 _dafny.Array
			_ = _991_v72
			var _nwElement0_37 _dafny.Set = (func() _dafny.Set {
				var _coll20 = _dafny.NewBuilder()
				_ = _coll20
				for _iter23 := _dafny.Iterate((_985_v66).Elements()); ; {
					_compr_20, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _992_v67 _dafny.Int
					_992_v67 = interface{}(_compr_20).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_985_v66, _992_v67) {
						_coll20.Add(Companion_Default___.SafeDivisionInt(_992_v67, (Companion_D2_.Create_DC3_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ieuojgwcd")).Cardinality()))).Dtor_cf4()))
					}
				}
				return _coll20.ToSet()
			}()).Intersection(_949_v46)
			_ = _nwElement0_37
			var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(23))
			_ = _nw172
			(_nw172).ArraySet1(_nwElement0_37, 0)
			(_nw172).ArraySet1(_949_v46, 1)
			(_nw172).ArraySet1(_949_v46, 2)
			(_nw172).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32((_986_v68).Cardinality()), _dafny.IntOfUint32((_950_v47).Cardinality())), 3)
			(_nw172).ArraySet1(_dafny.SetOf(_897_v0), 4)
			(_nw172).ArraySet1(_949_v46, 5)
			(_nw172).ArraySet1(_949_v46, 6)
			(_nw172).ArraySet1(_949_v46, 7)
			(_nw172).ArraySet1((_949_v46).Difference(_949_v46), 8)
			(_nw172).ArraySet1(((_951_v48).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_951_v48).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_949_v46), 9)
			(_nw172).ArraySet1(_949_v46, 10)
			(_nw172).ArraySet1(func() _dafny.Set {
				var _coll21 = _dafny.NewBuilder()
				_ = _coll21
				for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(622), _dafny.IntOfInt64(777))); ; {
					_compr_21, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _993_v69 _dafny.Int
					_993_v69 = interface{}(_compr_21).(_dafny.Int)
					if ((_dafny.IntOfInt64(622)).Cmp(_993_v69) <= 0) && ((_993_v69).Cmp(_dafny.IntOfInt64(777)) < 0) {
						_coll21.Add((_993_v69).Plus((Companion_D3_.Create_DC6_(p0, (_dafny.MultiSetOf(p0, (_this).F9())).Cardinality(), _897_v0)).Dtor_cf8()))
					}
				}
				return _coll21.ToSet()
			}(), 11)
			(_nw172).ArraySet1((_949_v46).Union(_949_v46), 12)
			(_nw172).ArraySet1(_949_v46, 13)
			(_nw172).ArraySet1(_949_v46, 14)
			(_nw172).ArraySet1((_dafny.SetOf(_897_v0)).Union(_dafny.SetOf(_897_v0, _897_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_987_v70, (_dafny.Zero).Minus(_897_v0))).Cardinality())), 15)
			(_nw172).ArraySet1(_949_v46, 16)
			(_nw172).ArraySet1((func() _dafny.Set {
				if (_990_v71).Contains(_984_cf31) {
					return (_990_v71).Get(_984_cf31).(_dafny.Set)
				}
				return _949_v46
			})(), 17)
			(_nw172).ArraySet1((_949_v46).Intersection(_949_v46), 18)
			(_nw172).ArraySet1((_949_v46).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-300))), 19)
			(_nw172).ArraySet1(_949_v46, 20)
			(_nw172).ArraySet1(_949_v46, 21)
			(_nw172).ArraySet1(_949_v46, 22)
			_991_v72 = _nw172
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_991_v72), 0))
			_ = _index172
			(_991_v72).ArraySet1(_dafny.SetOf(_897_v0), (_index172).Int())
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_991_v72), 0))
			_ = _index173
			(_991_v72).ArraySet1(_dafny.SetOf((_897_v0).Minus(_897_v0)), (_index173).Int())
			var _994_v73 _dafny.Array
			_ = _994_v73
			var _nw173 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw173
			_994_v73 = _nw173
			_994_v73 = _994_v73
			_994_v73 = _994_v73
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_987_v70), 0))
			_ = _index174
			(_987_v70).ArraySet1(_897_v0, (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_987_v70), 0))
			_ = _index175
			(_987_v70).ArraySet1(_897_v0, (_index175).Int())
		} else {
			var _995___mcc_h4 _dafny.Set = _source17.Get_().(D8_DC17).Cf27
			_ = _995___mcc_h4
			var _996_cf27 _dafny.Set = _995___mcc_h4
			_ = _996_cf27
			var _997_v74 _dafny.Array
			_ = _997_v74
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_23
			var _nw174 _dafny.Array
			_ = _nw174
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw174 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) bool = (func(_998_v0 _dafny.Int) func(_dafny.Int) bool {
					return func(_999_i7 _dafny.Int) bool {
						return ((Companion_D4_.Create_DC8_(_998_v0, (_this).F9())).Dtor_cf11()).Cmp(_998_v0) != 0
					}
				})(_897_v0)
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw174 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw174).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw174).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_997_v74 = _nw174
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_997_v74), 0))
			_ = _index176
			(_997_v74).ArraySet1(_this.F8(), (_index176).Int())
			var _1000_v75 _dafny.Map
			_ = _1000_v75
			_1000_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm0(globalState), (_this).F9())
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_997_v74), 0))
			_ = _index177
			(_997_v74).ArraySet1((func() bool {
				if (_1000_v75).Contains((_this).F9()) {
					return (_1000_v75).Get((_this).F9()).(bool)
				}
				return _this.F8()
			})(), (_index177).Int())
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_997_v74), 0))
			_ = _index178
			(_997_v74).ArraySet1(p0, (_index178).Int())
			var _rhs159 _dafny.Int = (_dafny.Zero).Minus(_897_v0)
			_ = _rhs159
			var _rhs160 _dafny.Int = _897_v0
			_ = _rhs160
			_897_v0 = _rhs159
			_897_v0 = _rhs160
			var _1001_v76 _dafny.Array
			_ = _1001_v76
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(Companion_D9_.Default(), _dafny.IntOfInt64(16))
			_ = _nw175
			_1001_v76 = _nw175
			var _1002_v77 D4
			_ = _1002_v77
			_1002_v77 = Companion_D4_.Create_DC7_(_950_v47)
			var _1003_v78 _dafny.Sequence
			_ = _1003_v78
			_1003_v78 = _dafny.SeqOf(_1002_v77, func(_pat_let35_0 D4) D4 {
				return func(_1004_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let36_0 _dafny.Sequence) D4 {
						return func(_1006_dt__update_hcf10_h1 _dafny.Sequence) D4 {
							return Companion_D4_.Create_DC7_(_1006_dt__update_hcf10_h1)
						}(_pat_let36_0)
					}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(219))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg57 _dafny.Int) interface{} {
							return coer57(arg57)
						}
					}(func(_1005_i8 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('i')
					})))
				}(_pat_let35_0)
			}(_1002_v77))
			var _1007_v79 _dafny.Set
			_ = _1007_v79
			_1007_v79 = _dafny.SetOf(_1003_v78)
			var _1008_v80 T2
			_ = _1008_v80
			var _nw176 *C1 = New_C1_()
			_ = _nw176
			_nw176.Ctor__((_950_v47).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_950_v47).Cardinality()))).Uint32()).(_dafny.CodePoint), _1007_v79, _this.F8(), (_997_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen((_997_v74), 0))).Int()).(bool))
			_1008_v80 = _nw176
			var _1009_v81 _dafny.Map
			_ = _1009_v81
			_1009_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1008_v80, _897_v0)
			var _1010_v82 D9
			_ = _1010_v82
			_1010_v82 = Companion_D9_.Create_DC22_(_897_v0, _897_v0, p0, _897_v0, _1009_v81)
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1001_v76), 0))
			_ = _index179
			(_1001_v76).ArraySet1(_1010_v82, (_index179).Int())
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((_1001_v76), 0))
			_ = _index180
			(_1001_v76).ArraySet1(Companion_D9_.Create_DC22_(_897_v0, (_897_v0).Times(_897_v0), p0, (Companion_Default___.Fm20(globalState)).Cardinality(), (_1009_v81).Merge((_1009_v81).Update(_1008_v80, _897_v0))), (_index180).Int())
		}
		var _1011_v83 D8
		_ = _1011_v83
		_1011_v83 = Companion_D8_.Create_DC19_(_this.F8())
		var _pat_let_tv33 = _950_v47
		_ = _pat_let_tv33
		_897_v0 = func(_source18 D8) _dafny.Int {
			if _source18.Is_DC18() {
				var _1012___mcc_h5 _dafny.Map = _source18.Get_().(D8_DC18).Cf28
				_ = _1012___mcc_h5
				var _1013___mcc_h6 bool = _source18.Get_().(D8_DC18).Cf29
				_ = _1013___mcc_h6
				var _1014___mcc_h7 _dafny.Int = _source18.Get_().(D8_DC18).Cf30
				_ = _1014___mcc_h7
				var _1015_cf30 _dafny.Int = _1014___mcc_h7
				_ = _1015_cf30
				var _1016_cf29 bool = _1013___mcc_h6
				_ = _1016_cf29
				var _1017_cf28 _dafny.Map = _1012___mcc_h5
				_ = _1017_cf28
				return (_1015_cf30).Times(_dafny.IntOfUint32((_dafny.SeqOf(_1015_cf30, _1015_cf30)).Cardinality()))
			} else if _source18.Is_DC19() {
				var _1018___mcc_h8 bool = _source18.Get_().(D8_DC19).Cf31
				_ = _1018___mcc_h8
				var _1019_cf31 bool = _1018___mcc_h8
				_ = _1019_cf31
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_pat_let_tv33).Cardinality()))
			} else {
				var _1020___mcc_h9 _dafny.Set = _source18.Get_().(D8_DC17).Cf27
				_ = _1020___mcc_h9
				var _1021_cf27 _dafny.Set = _1020___mcc_h9
				_ = _1021_cf27
				return _dafny.IntOfInt64(100)
			}
		}(_1011_v83)
		var _1022_v84 _dafny.Map
		_ = _1022_v84
		_1022_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _897_v0)
		var _1023_v85 _dafny.Sequence
		_ = _1023_v85
		_1023_v85 = _dafny.SeqOf(_1022_v84, Companion_Default___.Fm13(globalState), _1022_v84)
		var _1024_i9 _dafny.Int
		_ = _1024_i9
		_1024_i9 = _dafny.Zero
		{
			for (_dafny.IntOfUint32((_1023_v85).Cardinality())).Cmp(_897_v0) == 0 {
				{
					if (_1024_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1024_i9 = (_1024_i9).Plus(_dafny.One)
					var _1025_v86 _dafny.MultiSet
					_ = _1025_v86
					_1025_v86 = _dafny.MultiSetOf((_dafny.Zero).Minus(_897_v0), _897_v0)
					var _1026_v87 _dafny.Set
					_ = _1026_v87
					_1026_v87 = _dafny.SetOf(_this.F8(), (_1025_v86).IsSubsetOf(_1025_v86), (_this).F9(), p0)
					var _1027_v88 D0
					_ = _1027_v88
					_1027_v88 = Companion_D0_.Create_DC1_(p0, _dafny.SeqOf((_this).F9(), (_this).F9()), !(p0))
					var _rhs161 _dafny.Int = _897_v0
					_ = _rhs161
					var _rhs162 _dafny.Sequence = _950_v47
					_ = _rhs162
					var _rhs163 _dafny.Int = _897_v0
					_ = _rhs163
					var _rhs164 _dafny.Int = (_this).Fm2((func() D0 {
						if p0 {
							return _1027_v88
						}
						return func(_pat_let37_0 D0) D0 {
							return func(_1028_dt__update__tmp_h2 D0) D0 {
								return func(_pat_let38_0 bool) D0 {
									return func(_1029_dt__update_hcf0_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_1029_dt__update_hcf0_h0, (_1028_dt__update__tmp_h2).Dtor_cf1(), (_1028_dt__update__tmp_h2).Dtor_cf2())
									}(_pat_let38_0)
								}((_this).F9())
							}(_pat_let37_0)
						}(_1027_v88)
					})(), globalState)
					_ = _rhs164
					var _rhs165 _dafny.Set = _1026_v87
					_ = _rhs165
					r0 = _rhs161
					_950_v47 = _rhs162
					_897_v0 = _rhs163
					_897_v0 = _rhs164
					_1026_v87 = _rhs165
					if true {
						var _1030_v89 _dafny.Sequence
						_ = _1030_v89
						_1030_v89 = _dafny.SeqOf(_this.F8())
						r1 = (_1030_v89).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_897_v0, (_dafny.Zero).Minus(_897_v0)), _dafny.IntOfUint32((_1030_v89).Cardinality()))).Uint32()).(bool)
						(_this).F8_set_((_this).F9())
						var _1031_v90 _dafny.Array
						_ = _1031_v90
						var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
						_ = _nw177
						_1031_v90 = _nw177
						var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
						_ = _nw178
						_1031_v90 = _nw178
						var _1032_v91 _dafny.Map
						_ = _1032_v91
						_1032_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), Companion_Default___.Fm7(_897_v0, (_this).F9(), (_this).Fm2(_1027_v88, globalState), globalState))
						var _1033_v92 _dafny.Set
						_ = _1033_v92
						_1033_v92 = _dafny.SetOf(_1032_v91)
						var _1034_v93 _dafny.Sequence
						_ = _1034_v93
						_1034_v93 = _dafny.SeqOf((_this).Fm2(_1027_v88, globalState))
						var _1035_v94 _dafny.Map
						_ = _1035_v94
						_1035_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1033_v92, (func() _dafny.Int {
							if (_this).F9() {
								return _dafny.IntOfUint32((_1034_v93).Cardinality())
							}
							return _897_v0
						})())
						var _1036_v95 _dafny.Map
						_ = _1036_v95
						_1036_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), true)
						_1035_v94 = (_1035_v94).Update(_dafny.SetOf(_1032_v91, _1032_v91, _1032_v91), (_1036_v95).Cardinality())
						var _1037_v96 *C0
						_ = _1037_v96
						var _nw179 *C0 = New_C0_()
						_ = _nw179
						_nw179.Ctor__(_1031_v90)
						_1037_v96 = _nw179
					} else {
						var _1038_v97 D4
						_ = _1038_v97
						_1038_v97 = Companion_D4_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("ki"))
						var _pat_let_tv34 = _950_v47
						_ = _pat_let_tv34
						var _1039_v98 _dafny.Sequence
						_ = _1039_v98
						_1039_v98 = _dafny.SeqOf(_1038_v97, _1038_v97, func(_pat_let39_0 D4) D4 {
							return func(_1040_dt__update__tmp_h3 D4) D4 {
								return func(_pat_let40_0 _dafny.Sequence) D4 {
									return func(_1041_dt__update_hcf10_h2 _dafny.Sequence) D4 {
										return Companion_D4_.Create_DC7_(_1041_dt__update_hcf10_h2)
									}(_pat_let40_0)
								}(_pat_let_tv34)
							}(_pat_let39_0)
						}(_1038_v97), _1038_v97, _1038_v97)
						var _1042_v99 *C1
						_ = _1042_v99
						var _nw180 *C1 = New_C1_()
						_ = _nw180
						_nw180.Ctor__(_dafny.CodePoint('q'), _dafny.SetOf(_1039_v98), (_1026_v87).Equals(_1026_v87), !(_this.F8()) || (p0))
						_1042_v99 = _nw180
						var _1043_v100 _dafny.Array
						_ = _1043_v100
						var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
						_ = _nw181
						_1043_v100 = _nw181
						var _1044_v101 _dafny.Array
						_ = _1044_v101
						var _nwElement0_38 bool = false
						_ = _nwElement0_38
						var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(17))
						_ = _nw182
						(_nw182).ArraySet1(_nwElement0_38, 0)
						(_nw182).ArraySet1(_this.F8(), 1)
						(_nw182).ArraySet1((_this).F9(), 2)
						(_nw182).ArraySet1(false, 3)
						(_nw182).ArraySet1(false, 4)
						(_nw182).ArraySet1(p0, 5)
						(_nw182).ArraySet1((_this).F9(), 6)
						(_nw182).ArraySet1(_this.F8(), 7)
						(_nw182).ArraySet1((_this).F9(), 8)
						(_nw182).ArraySet1(_this.F8(), 9)
						(_nw182).ArraySet1(p0, 10)
						(_nw182).ArraySet1(_this.F8(), 11)
						(_nw182).ArraySet1(true, 12)
						(_nw182).ArraySet1(_this.F8(), 13)
						(_nw182).ArraySet1((_this).F9(), 14)
						(_nw182).ArraySet1(p0, 15)
						(_nw182).ArraySet1((_this).F9(), 16)
						_1044_v101 = _nw182
						var _1045_v102 _dafny.Sequence
						_ = _1045_v102
						_1045_v102 = _dafny.SeqOf(_1044_v101, _1044_v101, _1044_v101)
						var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1043_v100), 0))
						_ = _index181
						(_1043_v100).ArraySet1((_1045_v102).Select((Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_1045_v102).Cardinality()))).Uint32()).(_dafny.Array), (_index181).Int())
						var _1046_v103 _dafny.Sequence
						_ = _1046_v103
						_1046_v103 = _dafny.SeqOf(_897_v0)
						var _1047_v104 _dafny.Array
						_ = _1047_v104
						var _len0_24 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_24
						var _nw183 _dafny.Array
						_ = _nw183
						if _len0_24.Cmp(_dafny.Zero) == 0 {
							_nw183 = _dafny.NewArray(_len0_24)
						} else {
							var _init24 func(_dafny.Int) _dafny.Int = (func(_1048_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
								return func(_1049_i10 _dafny.Int) _dafny.Int {
									return (_1049_i10).Times(_1048_v0)
								}
							})(_897_v0)
							_ = _init24
							var _element0_24 = _init24(_dafny.Zero)
							_ = _element0_24
							_nw183 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
							(_nw183).ArraySet1(_element0_24, 0)
							var _nativeLen0_24 = (_len0_24).Int()
							_ = _nativeLen0_24
							for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
								(_nw183).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
							}
						}
						_1047_v104 = _nw183
						var _1050_v105 _dafny.Array
						_ = _1050_v105
						var _nwElement0_39 _dafny.Array = _1047_v104
						_ = _nwElement0_39
						var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(2))
						_ = _nw184
						(_nw184).ArraySet1(_nwElement0_39, 0)
						(_nw184).ArraySet1(_1047_v104, 1)
						_1050_v105 = _nw184
						var _1051_v106 *C0
						_ = _1051_v106
						var _nw185 *C0 = New_C0_()
						_ = _nw185
						_nw185.Ctor__(_1050_v105)
						_1051_v106 = _nw185
						var _1052_v107 _dafny.Map
						_ = _1052_v107
						_1052_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v106, (_this).F9())
						var _1053_v108 D4
						_ = _1053_v108
						_1053_v108 = Companion_D4_.Create_DC9_(p0, (_1052_v107).Cardinality(), _1044_v101, _897_v0)
						var _1054_v109 _dafny.Map
						_ = _1054_v109
						_1054_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_897_v0, (_1053_v108).Dtor_cf15())
						var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1043_v100), 0))
						_ = _index182
						var _rhs166 _dafny.Array = (func() _dafny.Array {
							if (_1054_v109).Contains((_897_v0).Plus(_897_v0)) {
								return (_1054_v109).Get((_897_v0).Plus(_897_v0)).(_dafny.Array)
							}
							return _1044_v101
						})()
						_ = _rhs166
						var _rhs167 _dafny.Array = _1044_v101
						_ = _rhs167
						var _rhs168 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-424))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg58 _dafny.Int) interface{} {
								return coer58(arg58)
							}
						}((func(_1055_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1056_i11 _dafny.Int) _dafny.Int {
								return (_dafny.Zero).Minus(_1055_v0)
							}
						})(_897_v0))), _1046_v103), _dafny.SeqOf((_dafny.SetOf(_897_v0)).Cardinality()))
						_ = _rhs168
						var _rhs169 _dafny.Sequence = _950_v47
						_ = _rhs169
						var _lhs89 _dafny.Array = _1043_v100
						_ = _lhs89
						var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(286), _dafny.ArrayLen((_1043_v100), 0))
						_ = _lhs90
						(_lhs89).ArraySet1(_rhs166, (_lhs90).Int())
						_1044_v101 = _rhs167
						_1046_v103 = _rhs168
						_950_v47 = _rhs169
						var _1057_v110 *C1
						_ = _1057_v110
						var _nw186 *C1 = New_C1_()
						_ = _nw186
						_nw186.Ctor__((_1042_v99).F10(), (_1042_v99).F11(), (p0) && (true), !(false))
						_1057_v110 = _nw186
						var _1058_v111 _dafny.CodePoint
						_ = _1058_v111
						_1058_v111 = _dafny.CodePoint('f')
						var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1047_v104), 0))
						_ = _index183
						(_1047_v104).ArraySet1((_1042_v99).Fm2(_1027_v88, globalState), (_index183).Int())
						var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1047_v104), 0))
						_ = _index184
						var _rhs170 _dafny.Array = _1047_v104
						_ = _rhs170
						var _rhs171 _dafny.CodePoint = _1058_v111
						_ = _rhs171
						var _rhs172 _dafny.Int = (_897_v0).Minus(_897_v0)
						_ = _rhs172
						var _lhs91 _dafny.Array = _1047_v104
						_ = _lhs91
						var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1047_v104), 0))
						_ = _lhs92
						_1047_v104 = _rhs170
						_1058_v111 = _rhs171
						(_lhs91).ArraySet1(_rhs172, (_lhs92).Int())
						var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1047_v104), 0))
						_ = _index185
						(_1047_v104).ArraySet1(Companion_Default___.SafeModuloInt((_1042_v99).Fm2(_1027_v88, globalState), (_1047_v104).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1047_v104), 0))).Int()).(_dafny.Int)), (_index185).Int())
					}
					var _1059_v112 _dafny.Array
					_ = _1059_v112
					var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
					_ = _nw187
					_1059_v112 = _nw187
					var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
					_ = _nw188
					_1059_v112 = _nw188
					var _1060_v113 _dafny.CodePoint
					_ = _1060_v113
					_1060_v113 = _dafny.CodePoint('e')
					var _1061_v114 _dafny.Sequence
					_ = _1061_v114
					_1061_v114 = _dafny.SeqOf((_this).F9(), p0, (_this).F9())
					r1 = (p0) && ((_this).Fm1(_dafny.Companion_Sequence_.Update(_950_v47, (Companion_Default___.SafeIndex(_897_v0, _dafny.IntOfUint32((_950_v47).Cardinality()))).Uint32(), _1060_v113), Companion_D0_.Create_DC1_(false, _1061_v114, _this.F8()), globalState))
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1062_v115 _dafny.MultiSet
		_ = _1062_v115
		_1062_v115 = _dafny.MultiSetOf(_1022_v84, _1022_v84)
		if (_1062_v115).IsProperSubsetOf(_1062_v115) {
			(_this).F8_set_(!(_this.F8()))
			var _1063_v116 _dafny.Sequence
			_ = _1063_v116
			_1063_v116 = _dafny.SeqOf(p0, _this.F8(), (_this).F9())
			var _1064_v117 D0
			_ = _1064_v117
			_1064_v117 = Companion_D0_.Create_DC1_(true, _1063_v116, p0)
			r0 = (_dafny.Zero).Minus((_this).Fm2(_1064_v117, globalState))
			var _1065_v118 _dafny.Array
			_ = _1065_v118
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw189
			_1065_v118 = _nw189
			var _1066_v119 _dafny.Sequence
			_ = _1066_v119
			_1066_v119 = _dafny.SeqOf(_1065_v118, _1065_v118, _1065_v118, _1065_v118, _1065_v118)
			var _1067_v120 _dafny.Array
			_ = _1067_v120
			var _nwElement0_40 _dafny.Array = _1065_v118
			_ = _nwElement0_40
			var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(17))
			_ = _nw190
			(_nw190).ArraySet1(_nwElement0_40, 0)
			(_nw190).ArraySet1(_1065_v118, 1)
			(_nw190).ArraySet1(_1065_v118, 2)
			(_nw190).ArraySet1(_1065_v118, 3)
			(_nw190).ArraySet1(_1065_v118, 4)
			(_nw190).ArraySet1(_1065_v118, 5)
			(_nw190).ArraySet1(_1065_v118, 6)
			(_nw190).ArraySet1(_1065_v118, 7)
			(_nw190).ArraySet1(_1065_v118, 8)
			(_nw190).ArraySet1(_1065_v118, 9)
			(_nw190).ArraySet1(_1065_v118, 10)
			(_nw190).ArraySet1(_1065_v118, 11)
			(_nw190).ArraySet1(_1065_v118, 12)
			(_nw190).ArraySet1((_1066_v119).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_897_v0), _dafny.IntOfUint32((_1066_v119).Cardinality()))).Uint32()).(_dafny.Array), 13)
			(_nw190).ArraySet1(_1065_v118, 14)
			(_nw190).ArraySet1(_1065_v118, 15)
			(_nw190).ArraySet1(_1065_v118, 16)
			_1067_v120 = _nw190
			var _1068_v121 *C0
			_ = _1068_v121
			var _nw191 *C0 = New_C0_()
			_ = _nw191
			_nw191.Ctor__(_1067_v120)
			_1068_v121 = _nw191
			var _1069_v122 _dafny.Map
			_ = _1069_v122
			_1069_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F8(), (_this).F9()), _this.F8())
			_1069_v122 = _1069_v122
			var _1070_v123 _dafny.CodePoint
			_ = _1070_v123
			_1070_v123 = _dafny.CodePoint('n')
			var _rhs173 _dafny.Int = (_this).Fm2(_1064_v117, globalState)
			_ = _rhs173
			var _rhs174 _dafny.CodePoint = (func() _dafny.CodePoint {
				if p0 {
					return _1070_v123
				}
				return _1070_v123
			})()
			_ = _rhs174
			var _rhs175 bool = !(_this.F8()) || (_this.F8())
			_ = _rhs175
			var _rhs176 _dafny.Int = (func() _dafny.Int {
				if (_897_v0).Cmp(_897_v0) == 0 {
					return _897_v0
				}
				return _897_v0
			})()
			_ = _rhs176
			var _rhs177 bool = (_this).F9()
			_ = _rhs177
			var _lhs93 *C3 = _this
			_ = _lhs93
			_897_v0 = _rhs173
			_1070_v123 = _rhs174
			r1 = _rhs175
			r0 = _rhs176
			_lhs93.F8_set_(_rhs177)
		} else {
			var _1071_v124 _dafny.Array
			_ = _1071_v124
			var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw192
			_1071_v124 = _nw192
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_1071_v124), 0))
			_ = _index186
			(_1071_v124).ArraySet1(_897_v0, (_index186).Int())
			var _1072_v125 _dafny.Sequence
			_ = _1072_v125
			_1072_v125 = _dafny.SeqOf((_this).F9(), (_this).F9())
			var _1073_v126 D0
			_ = _1073_v126
			_1073_v126 = Companion_D0_.Create_DC1_(_this.F8(), _1072_v125, (_this).F9())
			var _1074_v127 D0
			_ = _1074_v127
			_1074_v127 = Companion_D0_.Create_DC1_((_this).Fm1(_dafny.UnicodeSeqOfUtf8Bytes("k"), _1073_v126, globalState), _1072_v125, p0)
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_1071_v124), 0))
			_ = _index187
			(_1071_v124).ArraySet1((_this).Fm2(func(_pat_let41_0 D0) D0 {
				return func(_1075_dt__update__tmp_h4 D0) D0 {
					return func(_pat_let42_0 bool) D0 {
						return func(_1076_dt__update_hcf2_h0 bool) D0 {
							return Companion_D0_.Create_DC1_((_1075_dt__update__tmp_h4).Dtor_cf0(), (_1075_dt__update__tmp_h4).Dtor_cf1(), _1076_dt__update_hcf2_h0)
						}(_pat_let42_0)
					}((_this).F9())
				}(_pat_let41_0)
			}(_1074_v127), globalState), (_index187).Int())
			var _1077_v128 _dafny.Array
			_ = _1077_v128
			var _nw193 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
			_ = _nw193
			_1077_v128 = _nw193
			var _1078_v129 *C0
			_ = _1078_v129
			var _nw194 *C0 = New_C0_()
			_ = _nw194
			_nw194.Ctor__(_1077_v128)
			_1078_v129 = _nw194
			var _1079_v130 _dafny.Map
			_ = _1079_v130
			_1079_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1078_v129)
			_1079_v130 = (_1079_v130).Update(p0, _1078_v129)
			(_this).F8_set_(_this.F8())
			r0 = (_1071_v124).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_1071_v124), 0))).Int()).(_dafny.Int)
			var _1080_v131 _dafny.Set
			_ = _1080_v131
			_1080_v131 = _dafny.SetOf((_this).F9())
			_897_v0 = ((_1080_v131).Union(_1080_v131)).Cardinality()
		}
		var _1081_v132 _dafny.CodePoint
		_ = _1081_v132
		_1081_v132 = _dafny.CodePoint('h')
		r0 = Companion_Default___.SafeDivisionInt(_897_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(967))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}(func(_1082_i12 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-265), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(967))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg60 _dafny.Int) interface{} {
				return coer60(arg60)
			}
		}(func(_1083_i12 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()))).Uint32(), _1081_v132)).Cardinality()))
		r1 = _this.F8()
		return r0, r1
	}
}

// End of class C3
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
