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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, globalState *GlobalState) bool {
	if false {
		return !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("wvefrkb"), _dafny.UnicodeSeqOfUtf8Bytes("jwssl"))
	} else {
		return false
	}
}
func (_static *CompanionStruct_Default___) Fm1(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("domj"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("vg")))
}
func (_static *CompanionStruct_Default___) Fm2(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true, true, false), _dafny.SeqOf(!(false))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(!(true))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 bool, globalState *GlobalState) _dafny.Int {
	return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Times((_dafny.IntOfUint32((_dafny.SeqOf(!(false), !(false))).Cardinality())).Minus(_dafny.IntOfInt64(598)))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(882))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(519)
	})), true)
}
func (_static *CompanionStruct_Default___) Fm16(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if true {
			return _dafny.SetOf(_dafny.IntOfInt64(892), (_dafny.Zero).Minus((_dafny.SetOf(false)).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(353), _dafny.IntOfInt64(-515))).Cardinality())
		}
		return _dafny.SetOf((_dafny.SetOf(_dafny.IntOfInt64(-214))).Cardinality())
	})()).Union((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-357), _dafny.IntOfInt64(330))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-357)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(330)) < 0) {
				_coll0.Add(Companion_Default___.SafeDivisionInt(_2_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
			}
		}
		return _coll0.ToSet()
	}()).Difference(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(799), true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(52), true)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	var _source0 D1 = Companion_D1_.Create_DC4_((func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("prkc")).Cardinality()))))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _3_v0 _dafny.Map
			_3_v0 = interface{}(_compr_1).(_dafny.Map)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("prkc")).Cardinality()))))).Contains(_3_v0) {
				_coll1.Add(_3_v0)
			}
		}
		return _coll1.ToSet()
	}()).Cardinality())
	_ = _source0
	if _source0.Is_DC5() {
		var _4___mcc_h0 bool = _source0.Get_().(D1_DC5).Cf5
		_ = _4___mcc_h0
		var _5_cf5 bool = _4___mcc_h0
		_ = _5_cf5
		return _dafny.SetOf(_5_cf5)
	} else {
		var _6___mcc_h1 _dafny.Int = _source0.Get_().(D1_DC4).Cf4
		_ = _6___mcc_h1
		var _7_cf4 _dafny.Int = _6___mcc_h1
		_ = _7_cf4
		return _dafny.SetOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(946), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(461))).Cardinality()))).Cardinality())).Cardinality(), (_dafny.SetOf(true)).Cardinality()), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("begfq")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bcfe")).Cardinality()), _dafny.IntOfInt64(-714))).Cardinality()), false)
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)))
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Set, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D2 = Companion_D2_.Create_DC6_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})))
	_ = _source1
	if _source1.Is_DC7() {
		var _9___mcc_h0 _dafny.Int = _source1.Get_().(D2_DC7).Cf7
		_ = _9___mcc_h0
		var _10___mcc_h1 _dafny.Int = _source1.Get_().(D2_DC7).Cf8
		_ = _10___mcc_h1
		var _11_cf8 _dafny.Int = _10___mcc_h1
		_ = _11_cf8
		var _12_cf7 _dafny.Int = _9___mcc_h0
		_ = _12_cf7
		return _dafny.CodePoint('y')
	} else if _source1.Is_DC8() {
		var _13___mcc_h2 bool = _source1.Get_().(D2_DC8).Cf9
		_ = _13___mcc_h2
		var _14_cf9 bool = _13___mcc_h2
		_ = _14_cf9
		return _dafny.CodePoint('i')
	} else if _source1.Is_DC6() {
		var _15___mcc_h3 _dafny.Sequence = _source1.Get_().(D2_DC6).Cf6
		_ = _15___mcc_h3
		var _16_cf6 _dafny.Sequence = _15___mcc_h3
		_ = _16_cf6
		return _dafny.CodePoint('o')
	} else {
		var _17___mcc_h4 D2 = _source1.Get_().(D2_DC9).Cf10
		_ = _17___mcc_h4
		var _18_cf10 D2 = _17___mcc_h4
		_ = _18_cf10
		if true {
			return _dafny.CodePoint('e')
		} else {
			return _dafny.CodePoint('m')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm23(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_()
}
func (_static *CompanionStruct_Default___) Fm24(globalState *GlobalState) D0 {
	if _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("shj"), _dafny.UnicodeSeqOfUtf8Bytes("wfuu")) {
		return Companion_D0_.Create_DC0_(false)
	} else {
		return Companion_D0_.Create_DC0_(!(!(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, p2 D0, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Contains(false) {
		return func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-970), _dafny.IntOfInt64(-460))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _19_v0 _dafny.Int
				_19_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(-970)).Cmp(_19_v0) <= 0) && ((_19_v0).Cmp(_dafny.IntOfInt64(-460)) < 0) {
					_coll2.Add((_19_v0).Times(_dafny.IntOfInt64(802)), _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thvvevq")).Cardinality())), _dafny.IntOfInt64(-477), _dafny.IntOfInt64(803), (_dafny.MultiSetOf(false)).Cardinality()))
				}
			}
			return _coll2.ToMap()
		}()
	} else {
		return func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(488), _dafny.IntOfInt64(518))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _20_v1 _dafny.Int
				_20_v1 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(488)).Cmp(_20_v1) <= 0) && ((_20_v1).Cmp(_dafny.IntOfInt64(518)) < 0) {
					_coll3.Add(Companion_Default___.SafeDivisionInt(_20_v1, _dafny.IntOfInt64(567)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-215))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg3 _dafny.Int) interface{} {
							return coer3(arg3)
						}
					}(func(_21_i0 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-734)
					})))
				}
			}
			return _coll3.ToMap()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm26(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC5_(!_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("ero"), _dafny.CodePoint('f')))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Sequence, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return (func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(236))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_22_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-977)
		}))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _23_v0 _dafny.Int
			_23_v0 = interface{}(_compr_4).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(236))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_22_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-977)
			})), _23_v0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_23_v0, (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tpiay")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
					var _coll5 = _dafny.NewBuilder()
					_ = _coll5
					for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(14), _dafny.IntOfInt64(706))); ; {
						_compr_5, _ok5 := _iter5()
						if !_ok5 {
							break
						}
						var _24_v1 _dafny.Int
						_24_v1 = interface{}(_compr_5).(_dafny.Int)
						if ((_dafny.IntOfInt64(14)).Cmp(_24_v1) <= 0) && ((_24_v1).Cmp(_dafny.IntOfInt64(706)) < 0) {
							_coll5.Add(Companion_Default___.SafeModuloInt(_24_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg6 _dafny.Int) interface{} {
									return coer6(arg6)
								}
							}(func(_25_i1 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('x')
							}))).Cardinality())))
						}
					}
					return _coll5.ToSet()
				}()).Cardinality(), _dafny.IntOfInt64(659))).Cardinality()), _dafny.IntOfInt64(310))).Cardinality()), (_dafny.SetOf(false, false, !(true))).Cardinality())
			}
		}
		return _coll4.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-232), _dafny.IntOfInt64(13)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Map {
	var _source2 D21 = Companion_D21_.Create_DC46_(true, false, _dafny.IntOfInt64(966))
	_ = _source2
	if _source2.Is_DC46() {
		var _26___mcc_h0 bool = _source2.Get_().(D21_DC46).Cf61
		_ = _26___mcc_h0
		var _27___mcc_h1 bool = _source2.Get_().(D21_DC46).Cf62
		_ = _27___mcc_h1
		var _28___mcc_h2 _dafny.Int = _source2.Get_().(D21_DC46).Cf63
		_ = _28___mcc_h2
		var _29_cf63 _dafny.Int = _28___mcc_h2
		_ = _29_cf63
		var _30_cf62 bool = _27___mcc_h1
		_ = _30_cf62
		var _31_cf61 bool = _26___mcc_h0
		_ = _31_cf61
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _29_cf63)
	} else if _source2.Is_DC47() {
		var _32___mcc_h3 bool = _source2.Get_().(D21_DC47).Cf64
		_ = _32___mcc_h3
		var _33_cf64 bool = _32___mcc_h3
		_ = _33_cf64
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_33_cf64, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dgerly")).Cardinality()))
	} else {
		var _34___mcc_h4 _dafny.MultiSet = _source2.Get_().(D21_DC45).Cf60
		_ = _34___mcc_h4
		var _35_cf60 _dafny.MultiSet = _34___mcc_h4
		_ = _35_cf60
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(872))
	}
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfhxej")).Cardinality()))).Cardinality()), _dafny.SeqOf((_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(-270))).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-916), _dafny.IntOfInt64(191)), _dafny.SeqOf(_dafny.IntOfInt64(401), _dafny.IntOfInt64(4), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cwex")).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) _dafny.MultiSet {
	if false {
		return (_dafny.MultiSetOf(_dafny.IntOfInt64(-763))).Difference(_dafny.MultiSetOf((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("slnptiny"))).Cardinality(), _dafny.IntOfInt64(872), (_dafny.MultiSetOf(_dafny.IntOfInt64(467))).Cardinality()))
	} else {
		return _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.CodePoint('u'), _dafny.CodePoint('v'))).Cardinality(), _dafny.IntOfInt64(601))
	}
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.Map {
	var _source3 D12 = Companion_D12_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false))
	_ = _source3
	if _source3.Is_DC22() {
		var _36___mcc_h0 _dafny.Map = _source3.Get_().(D12_DC22).Cf24
		_ = _36___mcc_h0
		var _37___mcc_h1 _dafny.Sequence = _source3.Get_().(D12_DC22).Cf25
		_ = _37___mcc_h1
		var _38___mcc_h2 _dafny.CodePoint = _source3.Get_().(D12_DC22).Cf26
		_ = _38___mcc_h2
		var _39___mcc_h3 bool = _source3.Get_().(D12_DC22).Cf27
		_ = _39___mcc_h3
		var _40___mcc_h4 _dafny.Sequence = _source3.Get_().(D12_DC22).Cf28
		_ = _40___mcc_h4
		var _41_cf28 _dafny.Sequence = _40___mcc_h4
		_ = _41_cf28
		var _42_cf27 bool = _39___mcc_h3
		_ = _42_cf27
		var _43_cf26 _dafny.CodePoint = _38___mcc_h2
		_ = _43_cf26
		var _44_cf25 _dafny.Sequence = _37___mcc_h1
		_ = _44_cf25
		var _45_cf24 _dafny.Map = _36___mcc_h0
		_ = _45_cf24
		return (func() _dafny.Map {
			var _coll6 = _dafny.NewMapBuilder()
			_ = _coll6
			for _iter6 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(934))).Cardinality()), (_dafny.SetOf(_42_cf27)).Cardinality(), _dafny.IntOfInt64(-580), _dafny.IntOfUint32((_41_cf28).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-570)))).Cardinality())).Elements()); ; {
				_compr_6, _ok6 := _iter6()
				if !_ok6 {
					break
				}
				var _46_v0 _dafny.Int
				_46_v0 = interface{}(_compr_6).(_dafny.Int)
				if (_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(934))).Cardinality()), (_dafny.SetOf(_42_cf27)).Cardinality(), _dafny.IntOfInt64(-580), _dafny.IntOfUint32((_41_cf28).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-570)))).Cardinality())).Contains(_46_v0) {
					_coll6.Add(Companion_Default___.SafeDivisionInt(_46_v0, _dafny.IntOfInt64(-820)), _42_cf27)
				}
			}
			return _coll6.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_41_cf28).Cardinality()), false))
	} else if _source3.Is_DC23() {
		var _47___mcc_h5 _dafny.Array = _source3.Get_().(D12_DC23).Cf29
		_ = _47___mcc_h5
		var _48_cf29 _dafny.Array = _47___mcc_h5
		_ = _48_cf29
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-711), true)
	} else {
		var _49___mcc_h6 _dafny.Map = _source3.Get_().(D12_DC21).Cf23
		_ = _49___mcc_h6
		var _50_cf23 _dafny.Map = _49___mcc_h6
		_ = _50_cf23
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(34), true)
	}
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Sequence, p1 _dafny.MultiSet, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(379), _dafny.IntOfInt64(-580))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _51_v0 _dafny.Int
			_51_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(379)).Cmp(_51_v0) <= 0) && ((_51_v0).Cmp(_dafny.IntOfInt64(-580)) < 0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_51_v0, _dafny.IntOfInt64(624)))
			}
		}
		return _coll7.ToSet()
	}()).Intersection(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-9), _dafny.IntOfInt64(42))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _52_v1 _dafny.Int
			_52_v1 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(-9)).Cmp(_52_v1) <= 0) && ((_52_v1).Cmp(_dafny.IntOfInt64(42)) < 0) {
				_coll8.Add(Companion_Default___.SafeModuloInt(_52_v1, _dafny.IntOfInt64(-429)))
			}
		}
		return _coll8.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('h')
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 D1, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if false {
		return _dafny.SeqOf(_dafny.IntOfInt64(-517), _dafny.IntOfInt64(898), _dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(668)), _dafny.SeqOf(_dafny.IntOfInt64(452)))
	}
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-299)), _dafny.SeqOf(_dafny.IntOfInt64(-751)), _dafny.SeqOf(_dafny.IntOfInt64(854), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iwgq")).Cardinality())), false)).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(356))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(493), (_dafny.MultiSetOf(!(!(false)))).Cardinality())), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-902)))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if false {
		return _dafny.SeqOf(false)
	} else {
		return _dafny.SeqOf(false, !(false))
	}
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true, true, true), _dafny.IntOfInt64(649))
}
func (_static *CompanionStruct_Default___) Fm41(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
}
func (_static *CompanionStruct_Default___) Fm43(p0 bool, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(791))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_53_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()))))).Union((func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(853), _dafny.IntOfInt64(3))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _54_v0 _dafny.Int
			_54_v0 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(853)).Cmp(_54_v0) <= 0) && ((_54_v0).Cmp(_dafny.IntOfInt64(3)) < 0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_54_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gsah")).Cardinality())))
			}
		}
		return _coll9.ToSet()
	}()).Intersection(func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-232), _dafny.IntOfInt64(274))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _55_v1 _dafny.Int
			_55_v1 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(-232)).Cmp(_55_v1) <= 0) && ((_55_v1).Cmp(_dafny.IntOfInt64(274)) < 0) {
				_coll10.Add((_55_v1).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jynfus")).Cardinality())))
			}
		}
		return _coll10.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.CodePoint('q'), _dafny.CodePoint('s'))
}
func (_static *CompanionStruct_Default___) Fm45(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true, true)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-220))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_56_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(236))).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(971), _dafny.IntOfInt64(-640), (_dafny.MultiSetOf(_dafny.SetOf(!(true), true))).Cardinality(), (_dafny.SetOf(false, true)).Cardinality())).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-804))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_57_i1 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true, true), _dafny.CodePoint('x'))).Cardinality()
	}))))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(473), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hmbwmks")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm46(p0 bool, globalState *GlobalState) _dafny.Map {
	if !((func() bool {
		if false {
			return false
		}
		return !(false)
	})()) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
	}
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(305))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_58_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	}))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(!(true)), _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("f"), _dafny.UnicodeSeqOfUtf8Bytes("ju")), _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(false, !(true)), !(!(true))))
}
func (_static *CompanionStruct_Default___) Fm49(globalState *GlobalState) _dafny.Set {
	if true {
		return _dafny.SetOf(true)
	} else if !(true) {
		return _dafny.SetOf(true, !(false), true, true, true)
	} else {
		return _dafny.SetOf(true, true)
	}
}
func (_static *CompanionStruct_Default___) Fm50(p0 bool, p1 bool, p2 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(247), _dafny.IntOfInt64(643))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _59_v0 _dafny.Int
			_59_v0 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(247)).Cmp(_59_v0) <= 0) && ((_59_v0).Cmp(_dafny.IntOfInt64(643)) < 0) {
				_coll11.Add((_59_v0).Plus(_dafny.IntOfInt64(454)), !(true))
			}
		}
		return _coll11.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm51(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('r'))
}
func (_static *CompanionStruct_Default___) Fm52(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(398))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_60_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-816)
	})), (Companion_D6_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(772))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_61_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(247)
	})))).Dtor_cf14(), _dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-873))).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xmsk")).Cardinality()))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(513))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_62_i2 _dafny.Int) _dafny.Int {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(262), (_dafny.MultiSetOf(!(!(true)), true)).Cardinality())).Cardinality()
	})))
}
func (_static *CompanionStruct_Default___) Fm53(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('o')
}
func (_static *CompanionStruct_Default___) Fm54(globalState *GlobalState) D15 {
	return Companion_D15_.Create_DC28_(false, !(!(!((_dafny.IntOfInt64(48)).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_D19_.Create_DC41_(Companion_D19_.Create_DC38_(func() _dafny.Set {
		var _coll12 = _dafny.NewBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("jwfty"))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _63_v0 _dafny.Sequence
			_63_v0 = interface{}(_compr_12).(_dafny.Sequence)
			if (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("jwfty"))).Contains(_63_v0) {
				_coll12.Add(_63_v0)
			}
		}
		return _coll12.ToSet()
	}())), Companion_D19_.Create_DC41_(Companion_D19_.Create_DC41_(Companion_D19_.Create_DC40_(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wkltshb")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(467), _dafny.SeqOf(_dafny.IntOfInt64(111), _dafny.IntOfInt64(871))))))), true)).Cardinality()) == 0))))
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(479), _dafny.SeqOf(true))).Merge(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.SeqOf((func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.MultiSetOf(true))).Cardinality())).Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _64_v1 _dafny.Int
				_64_v1 = interface{}(_compr_14).(_dafny.Int)
				if (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.MultiSetOf(true))).Cardinality())).Contains(_64_v1) {
					_coll14.Add(Companion_Default___.SafeModuloInt(_64_v1, _dafny.IntOfInt64(67)), true)
				}
			}
			return _coll14.ToMap()
		}()).Cardinality())).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _65_v0 _dafny.Int
			_65_v0 = interface{}(_compr_13).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((func() _dafny.Map {
				var _coll15 = _dafny.NewMapBuilder()
				_ = _coll15
				for _iter15 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.MultiSetOf(true))).Cardinality())).Elements()); ; {
					_compr_15, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _64_v1 _dafny.Int
					_64_v1 = interface{}(_compr_15).(_dafny.Int)
					if (_dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.MultiSetOf(true))).Cardinality())).Contains(_64_v1) {
						_coll15.Add(Companion_Default___.SafeModuloInt(_64_v1, _dafny.IntOfInt64(67)), true)
					}
				}
				return _coll15.ToMap()
			}()).Cardinality()), _65_v0) {
				_coll13.Add((_65_v0).Minus(_dafny.IntOfInt64(520)), _dafny.SeqOf(false))
			}
		}
		return _coll13.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer14 func(_dafny.Int) D16) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_66_i0 _dafny.Int) D16 {
		return Companion_D16_.Create_DC30_(_dafny.SeqOf(_dafny.SetOf(false, false)))
	}))).Cardinality()), _dafny.SeqOf(true, true, true)))
}
func (_static *CompanionStruct_Default___) Fm56(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D12_.Create_DC21_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false))))).Cardinality()), _dafny.SeqOf((func() _dafny.Set {
		var _coll16 = _dafny.NewBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.CodePoint('v'))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _67_v0 _dafny.CodePoint
			_67_v0 = interface{}(_compr_16).(_dafny.CodePoint)
			if (_dafny.MultiSetOf(_dafny.CodePoint('v'))).Contains(_67_v0) {
				_coll16.Add(_67_v0)
			}
		}
		return _coll16.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(373))).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(829)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(361))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_68_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xibg")).Cardinality())
	}))))
}
func (_static *CompanionStruct_Default___) Fm57(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(350))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfInt64(81))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("eauvvtm"))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm58(globalState *GlobalState) D11 {
	if (Companion_D26_.Create_DC57_(true)).Dtor_cf77() {
		return Companion_D11_.Create_DC20_(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ugle")).Cardinality()), _dafny.IntOfInt64(317))).Cardinality())
	} else if !(false) {
		return Companion_D11_.Create_DC20_(false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), _dafny.MultiSetOf(_dafny.CodePoint('a'), _dafny.CodePoint('x'), _dafny.CodePoint('u')))).Cardinality(), _dafny.IntOfInt64(711))).Cardinality()))
	} else {
		return Companion_D11_.Create_DC20_(true, _dafny.IntOfInt64(-365))
	}
}
func (_static *CompanionStruct_Default___) Fm59(globalState *GlobalState) _dafny.Map {
	var _source4 D0 = Companion_D0_.Create_DC0_((Companion_D26_.Create_DC57_(!(true))).Dtor_cf77())
	_ = _source4
	if _source4.Is_DC1() {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), true)).Merge(func() _dafny.Map {
			var _coll17 = _dafny.NewMapBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('y'))).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _69_v0 _dafny.CodePoint
				_69_v0 = interface{}(_compr_17).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('y')), _69_v0) {
					_coll17.Add(_69_v0, !(true))
				}
			}
			return _coll17.ToMap()
		}())
	} else if _source4.Is_DC2() {
		var _70___mcc_h0 _dafny.Map = _source4.Get_().(D0_DC2).Cf1
		_ = _70___mcc_h0
		var _71___mcc_h1 _dafny.CodePoint = _source4.Get_().(D0_DC2).Cf2
		_ = _71___mcc_h1
		var _72_cf2 _dafny.CodePoint = _71___mcc_h1
		_ = _72_cf2
		var _73_cf1 _dafny.Map = _70___mcc_h0
		_ = _73_cf1
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_72_cf2, false))
	} else if _source4.Is_DC0() {
		var _74___mcc_h2 bool = _source4.Get_().(D0_DC0).Cf0
		_ = _74___mcc_h2
		var _75_cf0 bool = _74___mcc_h2
		_ = _75_cf0
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('u'), _75_cf0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), _75_cf0))
	} else {
		var _76___mcc_h3 D0 = _source4.Get_().(D0_DC3).Cf3
		_ = _76___mcc_h3
		var _77_cf3 D0 = _76___mcc_h3
		_ = _77_cf3
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('o'), true)
	}
}
func (_static *CompanionStruct_Default___) Fm60(globalState *GlobalState) D14 {
	return Companion_D14_.Create_DC26_((true) && (true), ((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oxkxihvnx")).Cardinality())))).Plus(_dafny.IntOfInt64(113)), _dafny.CodePoint('v'))
}
func (_static *CompanionStruct_Default___) Fm61(p0 _dafny.Map, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("jemqeaufg"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(723))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_78_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	})), _dafny.UnicodeSeqOfUtf8Bytes("oyqd")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("e"), _dafny.UnicodeSeqOfUtf8Bytes("yk")), _dafny.UnicodeSeqOfUtf8Bytes("acxl"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kqlctexte"), _dafny.UnicodeSeqOfUtf8Bytes("alohcbeea")))
}
func (_static *CompanionStruct_Default___) Fm62(p0 D15, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) D2 {
	var _source5 D11 = Companion_D11_.Create_DC19_(_dafny.CodePoint('g'))
	_ = _source5
	if _source5.Is_DC20() {
		var _79___mcc_h0 bool = _source5.Get_().(D11_DC20).Cf21
		_ = _79___mcc_h0
		var _80___mcc_h1 _dafny.Int = _source5.Get_().(D11_DC20).Cf22
		_ = _80___mcc_h1
		var _81_cf22 _dafny.Int = _80___mcc_h1
		_ = _81_cf22
		var _82_cf21 bool = _79___mcc_h0
		_ = _82_cf21
		return Companion_D2_.Create_DC8_(_82_cf21)
	} else {
		var _83___mcc_h2 _dafny.CodePoint = _source5.Get_().(D11_DC19).Cf20
		_ = _83___mcc_h2
		var _84_cf20 _dafny.CodePoint = _83___mcc_h2
		_ = _84_cf20
		return Companion_D2_.Create_DC8_(false)
	}
}
func (_static *CompanionStruct_Default___) Fm63(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(736))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_85_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})))
}
func (_static *CompanionStruct_Default___) Fm64(p0 _dafny.CodePoint, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC14_((true) == (false))
}
func (_static *CompanionStruct_Default___) Fm65(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D17 {
	return Companion_D17_.Create_DC34_(_dafny.CodePoint('r'))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Array, p1 _dafny.CodePoint, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) (D0, bool) {
	var r0 D0 = Companion_D0_.Default()
	_ = r0
	var r1 bool = false
	_ = r1
	var _86_v0 _dafny.Int
	_ = _86_v0
	_86_v0 = _dafny.IntOfInt64(842)
	var _87_v1 _dafny.Sequence
	_ = _87_v1
	_87_v1 = _dafny.SeqOf(_86_v0)
	var _88_v2 _dafny.Map
	_ = _88_v2
	_88_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_86_v0, _87_v1)
	var _89_v3 D19
	_ = _89_v3
	_89_v3 = Companion_D19_.Create_DC40_(!((!(p2)) == (p2)), _86_v0, _88_v2)
	_89_v3 = _89_v3
	(globalState).F1 = Companion_Default___.SafeModuloInt(_86_v0, _dafny.IntOfInt64(239))
	var _90_i0 _dafny.Int
	_ = _90_i0
	_90_i0 = _dafny.Zero
	{
		for Companion_Default___.Fm0(p2, globalState) {
			{
				if (_90_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_90_i0 = (_90_i0).Plus(_dafny.One)
				var _91_v4 _dafny.Array
				_ = _91_v4
				var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw0
				_91_v4 = _nw0
				var _92_v5 _dafny.Set
				_ = _92_v5
				_92_v5 = _dafny.SetOf(_91_v4)
				var _93_v6 _dafny.Map
				_ = _93_v6
				_93_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_92_v5).Cardinality(), Companion_D21_.Create_DC47_(p2))
				var _94_v7 _dafny.Map
				_ = _94_v7
				_94_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v6, p2)
				var _95_v8 _dafny.Map
				_ = _95_v8
				_95_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _93_v6)
				r1 = (func() bool {
					if (_94_v7).Contains((func() _dafny.Map {
						if (_95_v8).Contains(p3) {
							return (_95_v8).Get(p3).(_dafny.Map)
						}
						return _93_v6
					})()) {
						return (_94_v7).Get((func() _dafny.Map {
							if (_95_v8).Contains(p3) {
								return (_95_v8).Get(p3).(_dafny.Map)
							}
							return _93_v6
						})()).(bool)
					}
					return !(p2)
				})()
				var _96_v9 _dafny.Sequence
				_ = _96_v9
				_96_v9 = _dafny.UnicodeSeqOfUtf8Bytes("ioa")
				var _rhs0 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm3(p2, p2, globalState))
				_ = _rhs0
				var _rhs1 bool = p2
				_ = _rhs1
				var _rhs2 _dafny.Int = (_87_v1).Select((Companion_Default___.SafeIndex((_86_v0).Plus(_dafny.IntOfInt64(-259)), _dafny.IntOfUint32((_87_v1).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs2
				var _rhs3 bool = p2
				_ = _rhs3
				var _rhs4 _dafny.Sequence = _96_v9
				_ = _rhs4
				var _lhs0 *GlobalState = globalState
				_ = _lhs0
				var _lhs1 *GlobalState = globalState
				_ = _lhs1
				_lhs0.F1 = _rhs0
				r1 = _rhs1
				_lhs1.F1 = _rhs2
				r1 = _rhs3
				_96_v9 = _rhs4
				(globalState).F1 = Companion_Default___.SafeDivisionInt(_86_v0, _86_v0)
				r1 = p2
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	(globalState).F1 = (_dafny.Zero).Minus((_86_v0).Minus(_86_v0))
	var _97_v10 D21
	_ = _97_v10
	_97_v10 = Companion_D21_.Create_DC46_(p2, p2, _86_v0)
	(globalState).F1 = _dafny.IntOfUint32((func(_source6 D21) _dafny.Sequence {
		if _source6.Is_DC46() {
			var _98___mcc_h0 bool = _source6.Get_().(D21_DC46).Cf61
			_ = _98___mcc_h0
			var _99___mcc_h1 bool = _source6.Get_().(D21_DC46).Cf62
			_ = _99___mcc_h1
			var _100___mcc_h2 _dafny.Int = _source6.Get_().(D21_DC46).Cf63
			_ = _100___mcc_h2
			var _101_cf63 _dafny.Int = _100___mcc_h2
			_ = _101_cf63
			var _102_cf62 bool = _99___mcc_h1
			_ = _102_cf62
			var _103_cf61 bool = _98___mcc_h0
			_ = _103_cf61
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("irn"), _dafny.UnicodeSeqOfUtf8Bytes("vcyxtburc"))
		} else if _source6.Is_DC47() {
			var _104___mcc_h3 bool = _source6.Get_().(D21_DC47).Cf64
			_ = _104___mcc_h3
			var _105_cf64 bool = _104___mcc_h3
			_ = _105_cf64
			return _dafny.UnicodeSeqOfUtf8Bytes("d")
		} else {
			var _106___mcc_h4 _dafny.MultiSet = _source6.Get_().(D21_DC45).Cf60
			_ = _106___mcc_h4
			var _107_cf60 _dafny.MultiSet = _106___mcc_h4
			_ = _107_cf60
			return _dafny.UnicodeSeqOfUtf8Bytes("oesqkuq")
		}
	}(_97_v10)).Cardinality())
	var _108_i1 _dafny.Int
	_ = _108_i1
	_108_i1 = _dafny.Zero
	{
		for (_86_v0).Cmp(_86_v0) != 0 {
			{
				if (_108_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_108_i1 = (_108_i1).Plus(_dafny.One)
				var _109_v11 _dafny.Array
				_ = _109_v11
				var _nw1 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
				_ = _nw1
				_109_v11 = _nw1
				var _110_v12 _dafny.MultiSet
				_ = _110_v12
				_110_v12 = _dafny.MultiSetOf(_86_v0, _86_v0)
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))
				_ = _index0
				(_109_v11).ArraySet1((_110_v12).Union((_110_v12).Update(_dafny.IntOfInt64(145), Companion_Default___.Abs(_86_v0))), (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))
				_ = _index1
				(_109_v11).ArraySet1(_110_v12, (_index1).Int())
				if p2 {
					var _111_v13 _dafny.Sequence
					_ = _111_v13
					_111_v13 = _dafny.SeqOf(true, p2)
					var _112_v14 _dafny.Map
					_ = _112_v14
					_112_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_111_v13).Select((Companion_Default___.SafeIndex(_86_v0, _dafny.IntOfUint32((_111_v13).Cardinality()))).Uint32()).(bool), p2)
					r1 = (_111_v13).Select((Companion_Default___.SafeIndex((_112_v14).Cardinality(), _dafny.IntOfUint32((_111_v13).Cardinality()))).Uint32()).(bool)
					var _113_v15 D11
					_ = _113_v15
					_113_v15 = Companion_D11_.Create_DC20_(p2, ((_109_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))).Int()).(_dafny.MultiSet)).Cardinality())
					var _114_v16 _dafny.Map
					_ = _114_v16
					_114_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_113_v15).Dtor_cf21(), p0)
					var _115_v17 _dafny.Set
					_ = _115_v17
					_115_v17 = _dafny.SetOf(_86_v0, _86_v0)
					var _116_v18 _dafny.Map
					_ = _116_v18
					_116_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_115_v17, p0)
					var _117_v19 _dafny.Array
					_ = _117_v19
					var _nwElement0_0 _dafny.Array = p0
					_ = _nwElement0_0
					var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(25))
					_ = _nw2
					(_nw2).ArraySet1(_nwElement0_0, 0)
					(_nw2).ArraySet1(p0, 1)
					(_nw2).ArraySet1((func() _dafny.Array {
						if (_114_v16).Contains(!(p2)) {
							return (_114_v16).Get(!(p2)).(_dafny.Array)
						}
						return p0
					})(), 2)
					(_nw2).ArraySet1(p0, 3)
					(_nw2).ArraySet1(p0, 4)
					(_nw2).ArraySet1(p0, 5)
					(_nw2).ArraySet1(p0, 6)
					(_nw2).ArraySet1(p0, 7)
					(_nw2).ArraySet1(p0, 8)
					(_nw2).ArraySet1(p0, 9)
					(_nw2).ArraySet1(p0, 10)
					(_nw2).ArraySet1(p0, 11)
					(_nw2).ArraySet1(p0, 12)
					(_nw2).ArraySet1(p0, 13)
					(_nw2).ArraySet1(p0, 14)
					(_nw2).ArraySet1(p0, 15)
					(_nw2).ArraySet1(p0, 16)
					(_nw2).ArraySet1((func() _dafny.Array {
						if (_116_v18).Contains(_dafny.SetOf(_dafny.IntOfInt64(148), _86_v0)) {
							return (_116_v18).Get(_dafny.SetOf(_dafny.IntOfInt64(148), _86_v0)).(_dafny.Array)
						}
						return p0
					})(), 17)
					(_nw2).ArraySet1((func() _dafny.Array {
						if p2 {
							return p0
						}
						return p0
					})(), 18)
					(_nw2).ArraySet1(p0, 19)
					(_nw2).ArraySet1(p0, 20)
					(_nw2).ArraySet1(p0, 21)
					(_nw2).ArraySet1(p0, 22)
					(_nw2).ArraySet1(p0, 23)
					(_nw2).ArraySet1(p0, 24)
					_117_v19 = _nw2
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))
					_ = _index2
					(_117_v19).ArraySet1(p0, (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))
					_ = _index3
					(_117_v19).ArraySet1(p0, (_index3).Int())
					var _118_v20 _dafny.Array
					_ = _118_v20
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_0
					var _nw3 _dafny.Array
					_ = _nw3
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw3 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) _dafny.Int = (func(_119_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_120_i2 _dafny.Int) _dafny.Int {
								return (_120_i2).Plus(_119_v0)
							}
						})(_86_v0)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw3 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw3).ArraySet1(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw3).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_118_v20 = _nw3
					var _121_v21 _dafny.Array
					_ = _121_v21
					var _nwElement0_1 _dafny.Array = _118_v20
					_ = _nwElement0_1
					var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(22))
					_ = _nw4
					(_nw4).ArraySet1(_nwElement0_1, 0)
					(_nw4).ArraySet1(_118_v20, 1)
					(_nw4).ArraySet1(_118_v20, 2)
					(_nw4).ArraySet1(_118_v20, 3)
					(_nw4).ArraySet1(_118_v20, 4)
					(_nw4).ArraySet1(_118_v20, 5)
					(_nw4).ArraySet1(_118_v20, 6)
					(_nw4).ArraySet1(_118_v20, 7)
					(_nw4).ArraySet1(_118_v20, 8)
					(_nw4).ArraySet1(_118_v20, 9)
					(_nw4).ArraySet1(_118_v20, 10)
					(_nw4).ArraySet1(_118_v20, 11)
					(_nw4).ArraySet1(_118_v20, 12)
					(_nw4).ArraySet1((_118_v20), 13)
					(_nw4).ArraySet1(_118_v20, 14)
					(_nw4).ArraySet1(_118_v20, 15)
					(_nw4).ArraySet1(_118_v20, 16)
					(_nw4).ArraySet1(_118_v20, 17)
					(_nw4).ArraySet1(_118_v20, 18)
					(_nw4).ArraySet1(_118_v20, 19)
					(_nw4).ArraySet1((func() _dafny.Array {
						if p2 {
							return _118_v20
						}
						return _118_v20
					})(), 20)
					(_nw4).ArraySet1(_118_v20, 21)
					_121_v21 = _nw4
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_121_v21), 0))
					_ = _index4
					(_121_v21).ArraySet1(_118_v20, (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_121_v21), 0))
					_ = _index5
					(_121_v21).ArraySet1(_118_v20, (_index5).Int())
					var _122_v22 _dafny.Map
					_ = _122_v22
					_122_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
					_122_v22 = (_122_v22).Update(_dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int())), p2)
					var _arr0 _dafny.Array = _dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))
					_ = _arr0
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))), 0))
					_ = _index6
					(_arr0).ArraySet1(p2, (_index6).Int())
					var _arr1 _dafny.Array = _dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))
					_ = _arr1
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))), 0))
					_ = _index7
					var _rhs5 _dafny.Int = (_dafny.Zero).Minus(_86_v0)
					_ = _rhs5
					var _rhs6 _dafny.Int = (_dafny.IntOfInt64(-243)).Minus(_86_v0)
					_ = _rhs6
					var _rhs7 bool = true
					_ = _rhs7
					var _rhs8 bool = (_111_v13).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_86_v0), _dafny.IntOfUint32((_111_v13).Cardinality()))).Uint32()).(bool)
					_ = _rhs8
					var _lhs2 *GlobalState = globalState
					_ = _lhs2
					var _lhs3 _dafny.Array = _dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))
					_ = _lhs3
					var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(943), _dafny.ArrayLen((_dafny.ArrayCastTo((_117_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(689), _dafny.ArrayLen((_117_v19), 0))).Int()))), 0))
					_ = _lhs4
					_lhs2.F1 = _rhs5
					_86_v0 = _rhs6
					(_lhs3).ArraySet1(_rhs7, (_lhs4).Int())
					r1 = _rhs8
				} else {
					(globalState).F1 = (_dafny.Zero).Minus(_86_v0)
					var _123_v23 _dafny.Array
					_ = _123_v23
					var _len0_1 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_1
					var _nw5 _dafny.Array
					_ = _nw5
					if _len0_1.Cmp(_dafny.Zero) == 0 {
						_nw5 = _dafny.NewArray(_len0_1)
					} else {
						var _init1 func(_dafny.Int) _dafny.Sequence = (func(_124_p2 bool) func(_dafny.Int) _dafny.Sequence {
							return func(_125_i3 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_124_p2, _124_p2), _dafny.SeqOf(true, _124_p2))
							}
						})(p2)
						_ = _init1
						var _element0_1 = _init1(_dafny.Zero)
						_ = _element0_1
						_nw5 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
						(_nw5).ArraySet1(_element0_1, 0)
						var _nativeLen0_1 = (_len0_1).Int()
						_ = _nativeLen0_1
						for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
							(_nw5).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
						}
					}
					_123_v23 = _nw5
					var _126_v24 _dafny.Sequence
					_ = _126_v24
					_126_v24 = _dafny.SeqOf(p2, p2)
					var _127_v25 _dafny.Sequence
					_ = _127_v25
					_127_v25 = _dafny.SeqOf((_126_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_87_v1).Cardinality()), _dafny.IntOfUint32((_126_v24).Cardinality()))).Uint32()).(bool), p2)
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_123_v23), 0))
					_ = _index8
					(_123_v23).ArraySet1(_126_v24, (_index8).Int())
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_123_v23), 0))
					_ = _index9
					(_123_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_127_v25, _126_v24), (_index9).Int())
					var _128_v26 D14
					_ = _128_v26
					_128_v26 = Companion_D14_.Create_DC26_(true, _86_v0, p3)
					var _129_v27 *C7
					_ = _129_v27
					var _nw6 *C7 = New_C7_()
					_ = _nw6
					_nw6.Ctor__((_dafny.IntOfInt64(112)).Plus(_86_v0), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oabcxl")).Cardinality()), _86_v0), p3, (_128_v26).Dtor_cf32())
					_129_v27 = _nw6
					var _130_v28 _dafny.Set
					_ = _130_v28
					_130_v28 = _dafny.SetOf((_109_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))).Int()).(_dafny.MultiSet), _110_v12, (_109_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))).Int()).(_dafny.MultiSet))
					r1 = ((_130_v28).Union(_130_v28)).Equals((_130_v28).Union(_130_v28))
					var _131_v29 _dafny.Map
					_ = _131_v29
					_131_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-444), _129_v27.F10)
					var _132_v30 _dafny.Set
					_ = _132_v30
					_132_v30 = _dafny.SetOf(p2, !(p2))
					var _133_v31 _dafny.Sequence
					_ = _133_v31
					_133_v31 = _dafny.UnicodeSeqOfUtf8Bytes("ycbrppf")
					var _134_v32 _dafny.MultiSet
					_ = _134_v32
					_134_v32 = _dafny.MultiSetOf(_133_v31)
					var _135_v33 _dafny.Map
					_ = _135_v33
					_135_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(402))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}(func(_136_i4 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('j')
					})))
					var _137_v34 _dafny.Array
					_ = _137_v34
					var _nwElement0_2 _dafny.Int = _86_v0
					_ = _nwElement0_2
					var _nw7 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(18))
					_ = _nw7
					(_nw7).ArraySet1(_nwElement0_2, 0)
					(_nw7).ArraySet1(_86_v0, 1)
					(_nw7).ArraySet1(_129_v27.F10, 2)
					(_nw7).ArraySet1((_131_v29).Cardinality(), 3)
					(_nw7).ArraySet1((_129_v27).F9(), 4)
					(_nw7).ArraySet1(_86_v0, 5)
					(_nw7).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_129_v27.F10, (_132_v30).Cardinality(), (_134_v32).Cardinality(), (_135_v33).Cardinality())).Cardinality()), 6)
					(_nw7).ArraySet1(_129_v27.F10, 7)
					(_nw7).ArraySet1(_dafny.IntOfInt64(552), 8)
					(_nw7).ArraySet1(_86_v0, 9)
					(_nw7).ArraySet1((_129_v27).F9(), 10)
					(_nw7).ArraySet1(_86_v0, 11)
					(_nw7).ArraySet1(_86_v0, 12)
					(_nw7).ArraySet1(_dafny.IntOfInt64(60), 13)
					(_nw7).ArraySet1((_129_v27).F9(), 14)
					(_nw7).ArraySet1(((_109_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_109_v11), 0))).Int()).(_dafny.MultiSet)).Cardinality(), 15)
					(_nw7).ArraySet1((_129_v27).F9(), 16)
					(_nw7).ArraySet1(_129_v27.F10, 17)
					_137_v34 = _nw7
					var _138_v35 _dafny.Map
					_ = _138_v35
					_138_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_137_v34, (_129_v27).F9())
					var _139_v36 _dafny.Map
					_ = _139_v36
					_139_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
					_138_v35 = (_138_v35).Update(_137_v34, (_129_v27.F10).Minus((_139_v36).Cardinality()))
				}
				var _140_v37 _dafny.Sequence
				_ = _140_v37
				_140_v37 = _dafny.SeqOf(p2)
				var _141_v38 _dafny.Sequence
				_ = _141_v38
				_141_v38 = _dafny.SeqOf(p2, p2, (_140_v37).Select((Companion_Default___.SafeIndex(_86_v0, _dafny.IntOfUint32((_140_v37).Cardinality()))).Uint32()).(bool))
				_86_v0 = (func() _dafny.Int {
					if (func() bool {
						if p2 {
							return p2
						}
						return p2
					})() {
						return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_140_v37, _141_v38)).Cardinality())
					}
					return _86_v0
				})()
				if (Companion_Default___.SafeDivisionInt((_87_v1).Select((Companion_Default___.SafeIndex(_86_v0, _dafny.IntOfUint32((_87_v1).Cardinality()))).Uint32()).(_dafny.Int), _86_v0)).Cmp((_86_v0).Minus(_dafny.IntOfInt64(-634))) >= 0 {
					var _142_v39 _dafny.Array
					_ = _142_v39
					var _nwElement0_3 _dafny.CodePoint = p1
					_ = _nwElement0_3
					var _nw8 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(6))
					_ = _nw8
					(_nw8).ArraySet1CodePoint(_nwElement0_3, 0)
					(_nw8).ArraySet1CodePoint(p1, 1)
					(_nw8).ArraySet1CodePoint(p3, 2)
					(_nw8).ArraySet1CodePoint(p1, 3)
					(_nw8).ArraySet1CodePoint(p1, 4)
					(_nw8).ArraySet1CodePoint(p1, 5)
					_142_v39 = _nw8
					var _143_v40 _dafny.MultiSet
					_ = _143_v40
					_143_v40 = _dafny.MultiSetOf(_142_v39, _142_v39)
					var _144_v41 T3
					_ = _144_v41
					var _nw9 *C7 = New_C7_()
					_ = _nw9
					_nw9.Ctor__(_86_v0, _86_v0, p3, p2)
					_144_v41 = _nw9
					var _145_v42 _dafny.Map
					_ = _145_v42
					_145_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _144_v41)
					var _146_v43 _dafny.Map
					_ = _146_v43
					_146_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_143_v40, (_145_v42).Merge(_145_v42))
					var _147_v44 _dafny.Sequence
					_ = _147_v44
					_147_v44 = _dafny.SeqOf(_145_v42, _145_v42, _145_v42)
					_146_v43 = (_146_v43).Update((_143_v40).Union(_143_v40), (_147_v44).Select((Companion_Default___.SafeIndex(_86_v0, _dafny.IntOfUint32((_147_v44).Cardinality()))).Uint32()).(_dafny.Map))
					var _148_v45 _dafny.Array
					_ = _148_v45
					var _nw10 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
					_ = _nw10
					_148_v45 = _nw10
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index10
					(p0).ArraySet1((_dafny.IntOfInt64(865)).Cmp(_86_v0) != 0, (_index10).Int())
					var _149_v46 D26
					_ = _149_v46
					_149_v46 = Companion_D26_.Create_DC55_(_148_v45)
					var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index11
					var _rhs9 _dafny.Array = (_149_v46).Dtor_cf74()
					_ = _rhs9
					var _rhs10 bool = p2
					_ = _rhs10
					var _lhs5 _dafny.Array = p0
					_ = _lhs5
					var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _lhs6
					_148_v45 = _rhs9
					(_lhs5).ArraySet1(_rhs10, (_lhs6).Int())
					var _150_v47 _dafny.Map
					_ = _150_v47
					_150_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _86_v0)
					var _151_v48 _dafny.Set
					_ = _151_v48
					_151_v48 = _dafny.SetOf(_dafny.IntOfInt64(642))
					var _152_v49 _dafny.Set
					_ = _152_v49
					_152_v49 = _dafny.SetOf(_151_v48, _151_v48)
					var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index12
					var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index13
					var _rhs11 _dafny.Map = (_150_v47).Update(p3, (_dafny.Zero).Minus(_86_v0))
					_ = _rhs11
					var _rhs12 bool = (_144_v41).F6()
					_ = _rhs12
					var _rhs13 bool = (_152_v49).IsProperSubsetOf(_dafny.SetOf(_151_v48))
					_ = _rhs13
					var _lhs7 _dafny.Array = p0
					_ = _lhs7
					var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _lhs8
					var _lhs9 _dafny.Array = p0
					_ = _lhs9
					var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _lhs10
					_150_v47 = _rhs11
					(_lhs7).ArraySet1(_rhs12, (_lhs8).Int())
					(_lhs9).ArraySet1(_rhs13, (_lhs10).Int())
					var _153_v50 _dafny.Map
					_ = _153_v50
					_153_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_144_v41).F6(), _87_v1)
					var _154_v51 _dafny.MultiSet
					_ = _154_v51
					_154_v51 = _dafny.MultiSetOf(p2)
					var _155_v52 _dafny.Sequence
					_ = _155_v52
					_155_v52 = _dafny.UnicodeSeqOfUtf8Bytes("swp")
					_153_v50 = (_153_v50).Update(((_154_v51).Cardinality()).Cmp(_dafny.IntOfUint32((_155_v52).Cardinality())) > 0, _87_v1)
					var _156_v53 _dafny.MultiSet
					_ = _156_v53
					_156_v53 = _dafny.MultiSetOf(_155_v52)
					var _157_v54 _dafny.Array
					_ = _157_v54
					var _len0_2 _dafny.Int = _dafny.IntOfInt64(19)
					_ = _len0_2
					var _nw11 _dafny.Array
					_ = _nw11
					if _len0_2.Cmp(_dafny.Zero) == 0 {
						_nw11 = _dafny.NewArray(_len0_2)
					} else {
						var _init2 func(_dafny.Int) _dafny.Int = (func(_158_v41 T3, _159_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_160_i5 _dafny.Int) _dafny.Int {
								return (_160_i5).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_158_v41).F6(), _159_v0)).Cardinality())
							}
						})(_144_v41, _86_v0)
						_ = _init2
						var _element0_2 = _init2(_dafny.Zero)
						_ = _element0_2
						_nw11 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
						(_nw11).ArraySet1(_element0_2, 0)
						var _nativeLen0_2 = (_len0_2).Int()
						_ = _nativeLen0_2
						for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
							(_nw11).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
						}
					}
					_157_v54 = _nw11
					var _161_v55 _dafny.Map
					_ = _161_v55
					_161_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_157_v54, _156_v53)
					var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index14
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _index15
					var _rhs14 bool = !(!((_144_v41).F6()))
					_ = _rhs14
					var _rhs15 bool = !((func() _dafny.MultiSet {
						if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
							return _156_v53
						}
						return (func() _dafny.MultiSet {
							if (_161_v55).Contains(_157_v54) {
								return (_161_v55).Get(_157_v54).(_dafny.MultiSet)
							}
							return _dafny.MultiSetOf(_155_v52)
						})()
					})()).Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tyay"), _155_v52))
					_ = _rhs15
					var _rhs16 _dafny.Int = _86_v0
					_ = _rhs16
					var _rhs17 _dafny.Int = ((_87_v1).Select((Companion_Default___.SafeIndex(_86_v0, _dafny.IntOfUint32((_87_v1).Cardinality()))).Uint32()).(_dafny.Int)).Minus(_86_v0)
					_ = _rhs17
					var _rhs18 bool = Companion_Default___.Fm0(!(p2), globalState)
					_ = _rhs18
					var _lhs11 _dafny.Array = p0
					_ = _lhs11
					var _lhs12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _lhs12
					var _lhs13 *GlobalState = globalState
					_ = _lhs13
					var _lhs14 *GlobalState = globalState
					_ = _lhs14
					var _lhs15 _dafny.Array = p0
					_ = _lhs15
					var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((p0), 0))
					_ = _lhs16
					r1 = _rhs14
					(_lhs11).ArraySet1(_rhs15, (_lhs12).Int())
					_lhs13.F1 = _rhs16
					_lhs14.F1 = _rhs17
					(_lhs15).ArraySet1(_rhs18, (_lhs16).Int())
				} else {
					var _162_v56 _dafny.Sequence
					_ = _162_v56
					_162_v56 = _dafny.UnicodeSeqOfUtf8Bytes("xarvj")
					_162_v56 = _162_v56
					var _163_v57 *C4
					_ = _163_v57
					var _nw12 *C4 = New_C4_()
					_ = _nw12
					_nw12.Ctor__()
					_163_v57 = _nw12
					var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
					_ = _index16
					(p0).ArraySet1(p2, (_index16).Int())
					var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
					_ = _index17
					(p0).ArraySet1(true, (_index17).Int())
					(globalState).F1 = (func() _dafny.Int {
						if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
							return _86_v0
						}
						return _86_v0
					})()
					_86_v0 = (_dafny.Zero).Minus(_86_v0)
				}
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _164_v58 _dafny.Map
	_ = _164_v58
	_164_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
	var _165_v59 D0
	_ = _165_v59
	_165_v59 = Companion_D0_.Create_DC2_((_164_v58).Update(p0, p3), p1)
	r0 = _165_v59
	var _166_v60 _dafny.Set
	_ = _166_v60
	_166_v60 = _dafny.SetOf(_dafny.MultiSetFromSeq(_87_v1))
	r1 = ((_166_v60).Equals(_166_v60)) == (Companion_Default___.Fm0(p2, globalState))
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _167_v0 _dafny.Int
	_ = _167_v0
	_167_v0 = _dafny.IntOfInt64(647)
	var _168_v1 _dafny.Sequence
	_ = _168_v1
	_168_v1 = _dafny.SeqOf(_167_v0, _167_v0)
	var _169_v2 _dafny.Sequence
	_ = _169_v2
	_169_v2 = _dafny.SeqOf(_168_v1)
	var _170_v3 bool
	_ = _170_v3
	_170_v3 = true
	var _171_v4 D0
	_ = _171_v4
	_171_v4 = Companion_D0_.Create_DC0_(_170_v3)
	var _172_v5 _dafny.Sequence
	_ = _172_v5
	_172_v5 = _dafny.SeqOf(_170_v3, _170_v3, true, _170_v3, _170_v3)
	var _173_globalState *GlobalState
	_ = _173_globalState
	var _nw13 *GlobalState = New_GlobalState_()
	_ = _nw13
	_nw13.Ctor__(_dafny.CodePoint('i'), _dafny.IntOfInt64(10), _dafny.IntOfInt64(430), _169_v2, _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_170_v3, (_171_v4).Dtor_cf0()), _172_v5)))
	_173_globalState = _nw13
	var _174_i0 _dafny.Int
	_ = _174_i0
	_174_i0 = _dafny.Zero
	{
		for _170_v3 {
			{
				if (_174_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_174_i0 = (_174_i0).Plus(_dafny.One)
				var _175_v6 _dafny.Array
				_ = _175_v6
				var _len0_3 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_3
				var _nw14 _dafny.Array
				_ = _nw14
				if _len0_3.Cmp(_dafny.Zero) == 0 {
					_nw14 = _dafny.NewArray(_len0_3)
				} else {
					var _init3 func(_dafny.Int) bool = (func(_176_v3 bool) func(_dafny.Int) bool {
						return func(_177_i1 _dafny.Int) bool {
							return _176_v3
						}
					})(_170_v3)
					_ = _init3
					var _element0_3 = _init3(_dafny.Zero)
					_ = _element0_3
					_nw14 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
					(_nw14).ArraySet1(_element0_3, 0)
					var _nativeLen0_3 = (_len0_3).Int()
					_ = _nativeLen0_3
					for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
						(_nw14).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
					}
				}
				_175_v6 = _nw14
				var _178_v7 _dafny.CodePoint
				_ = _178_v7
				_178_v7 = _dafny.CodePoint('w')
				var _179_v8 D0
				_ = _179_v8
				var _180_v9 bool
				_ = _180_v9
				var _out0 D0
				_ = _out0
				var _out1 bool
				_ = _out1
				_out0, _out1 = Companion_Default___.M0(_175_v6, _178_v7, _170_v3, _178_v7, _173_globalState)
				_179_v8 = _out0
				_180_v9 = _out1
				var _181_v10 _dafny.Array
				_ = _181_v10
				var _nw15 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw15
				_181_v10 = _nw15
				var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_181_v10), 0))
				_ = _index18
				(_181_v10).ArraySet1(_167_v0, (_index18).Int())
				var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_181_v10), 0))
				_ = _index19
				(_181_v10).ArraySet1((_167_v0).Plus((func() _dafny.Int {
					if _180_v9 {
						return _167_v0
					}
					return _167_v0
				})()), (_index19).Int())
				var _182_v11 _dafny.Map
				_ = _182_v11
				_182_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _179_v8)
				var _183_v12 _dafny.Map
				_ = _183_v12
				_183_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(325), _167_v0)
				var _184_v13 _dafny.Map
				_ = _184_v13
				_184_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_183_v12, _182_v11)
				var _185_v14 _dafny.Map
				_ = _185_v14
				_185_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_181_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_181_v10), 0))).Int()).(_dafny.Int)).Times(_167_v0), _dafny.SetOf(_182_v11, (func() _dafny.Map {
					if (_184_v13).Contains(_183_v12) {
						return (_184_v13).Get(_183_v12).(_dafny.Map)
					}
					return _182_v11
				})(), _182_v11))
				var _186_v15 _dafny.Set
				_ = _186_v15
				_186_v15 = _dafny.SetOf(_182_v11)
				_185_v14 = (_185_v14).Update(Companion_Default___.SafeDivisionInt((_181_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_181_v10), 0))).Int()).(_dafny.Int), (_181_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen((_181_v10), 0))).Int()).(_dafny.Int)), _186_v15)
				var _187_v16 _dafny.MultiSet
				_ = _187_v16
				_187_v16 = _dafny.MultiSetOf(_167_v0, _dafny.IntOfInt64(493))
				var _188_v17 _dafny.Map
				_ = _188_v17
				_188_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_170_v3, _173_globalState), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(379), (_187_v16).Cardinality()))
				_188_v17 = (_188_v17).Update(_170_v3, _167_v0)
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _189_v18 D0
	_ = _189_v18
	_189_v18 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_172_v5).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_172_v5).Cardinality()))).Uint32()).(bool)))
	_189_v18 = _189_v18
	var _190_v19 _dafny.Array
	_ = _190_v19
	var _nwElement0_4 bool = _170_v3
	_ = _nwElement0_4
	var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(7))
	_ = _nw16
	(_nw16).ArraySet1(_nwElement0_4, 0)
	(_nw16).ArraySet1(true, 1)
	(_nw16).ArraySet1(_170_v3, 2)
	(_nw16).ArraySet1(_170_v3, 3)
	(_nw16).ArraySet1(_170_v3, 4)
	(_nw16).ArraySet1(_170_v3, 5)
	(_nw16).ArraySet1(_170_v3, 6)
	_190_v19 = _nw16
	var _191_v20 _dafny.CodePoint
	_ = _191_v20
	_191_v20 = _dafny.CodePoint('j')
	var _192_v21 D0
	_ = _192_v21
	var _193_v22 bool
	_ = _193_v22
	var _out2 D0
	_ = _out2
	var _out3 bool
	_ = _out3
	_out2, _out3 = Companion_Default___.M0(_190_v19, _191_v20, _170_v3, _191_v20, _173_globalState)
	_192_v21 = _out2
	_193_v22 = _out3
	var _194_v23 _dafny.Map
	_ = _194_v23
	_194_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_v20, _167_v0)
	_194_v23 = (_194_v23).Merge((_194_v23).Merge(_194_v23))
	_191_v20 = _191_v20
	var _195_v24 _dafny.Array
	_ = _195_v24
	var _nw17 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
	_ = _nw17
	_195_v24 = _nw17
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
	_ = _index20
	(_195_v24).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ebdgugv"), (_index20).Int())
	var _196_v25 _dafny.Sequence
	_ = _196_v25
	_196_v25 = _dafny.UnicodeSeqOfUtf8Bytes("ds")
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
	_ = _index21
	(_195_v24).ArraySet1(_196_v25, (_index21).Int())
	var _197_v26 _dafny.Map
	_ = _197_v26
	_197_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_190_v19, _191_v20)
	var _198_v27 _dafny.Array
	_ = _198_v27
	var _nwElement0_5 D0 = _192_v21
	_ = _nwElement0_5
	var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(4))
	_ = _nw18
	(_nw18).ArraySet1(_nwElement0_5, 0)
	(_nw18).ArraySet1(Companion_D0_.Create_DC2_(_197_v26, _191_v20), 1)
	(_nw18).ArraySet1(_192_v21, 2)
	(_nw18).ArraySet1(_192_v21, 3)
	_198_v27 = _nw18
	var _199_v28 _dafny.Sequence
	_ = _199_v28
	_199_v28 = _dafny.SeqOf(_198_v27)
	_198_v27 = (_199_v28).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_199_v28).Cardinality()))).Uint32()).(_dafny.Array)
	if _193_v22 {
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
		_ = _index22
		(_190_v19).ArraySet1(_170_v3, (_index22).Int())
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
		_ = _index23
		(_190_v19).ArraySet1(!(_dafny.Companion_Sequence_.Equal((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence), Companion_Default___.Fm1(_167_v0, (_172_v5).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_172_v5).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(606), _173_globalState))), (_index23).Int())
		var _200_v29 _dafny.Array
		_ = _200_v29
		var _nw19 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw19
		_200_v29 = _nw19
		_200_v29 = _200_v29
		var _201_v31 _dafny.Map
		_ = _201_v31
		_201_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm2(_170_v3, _193_v22, _173_globalState), _168_v1)
		var _202_v32 _dafny.Map
		_ = _202_v32
		_202_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v1, (func() _dafny.Sequence {
			if (_201_v31).Contains(_dafny.SeqOf((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool))) {
				return (_201_v31).Get(_dafny.SeqOf((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool))).(_dafny.Sequence)
			}
			return (_169_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_169_v2).Cardinality()))).Uint32()).(_dafny.Sequence)
		})())
		var _203_v33 _dafny.MultiSet
		_ = _203_v33
		_203_v33 = _dafny.MultiSetOf(_191_v20, _191_v20)
		var _204_v34 _dafny.Map
		_ = _204_v34
		_204_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), _203_v33)
		var _205_v35 D1
		_ = _205_v35
		_205_v35 = Companion_D1_.Create_DC4_((_204_v34).Cardinality())
		var _206_v36 D1
		_ = _206_v36
		_206_v36 = Companion_D1_.Create_DC5_(!((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool)))
		var _rhs19 _dafny.Int = _dafny.IntOfInt64(-242)
		_ = _rhs19
		var _rhs20 _dafny.Int = ((func() _dafny.Map {
			var _coll18 = _dafny.NewMapBuilder()
			_ = _coll18
			for _iter18 := _dafny.Iterate(((_202_v32).Update(_168_v1, _168_v1)).Keys().Elements()); ; {
				_compr_18, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _207_v30 _dafny.Sequence
				_207_v30 = interface{}(_compr_18).(_dafny.Sequence)
				if ((_202_v32).Update(_168_v1, _168_v1)).Contains(_207_v30) {
					_coll18.Add(_207_v30, _193_v22)
				}
			}
			return _coll18.ToMap()
		}()).Cardinality()).Plus((_205_v35).Dtor_cf4())
		_ = _rhs20
		var _rhs21 bool = (_206_v36).Dtor_cf5()
		_ = _rhs21
		var _lhs17 *GlobalState = _173_globalState
		_ = _lhs17
		_lhs17.F1 = _rhs19
		_167_v0 = _rhs20
		_170_v3 = _rhs21
		var _208_v37 _dafny.Set
		_ = _208_v37
		_208_v37 = _dafny.SetOf((_dafny.Zero).Minus((func() _dafny.Int {
			if (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool) {
				return _dafny.IntOfInt64(370)
			}
			return _167_v0
		})()), _167_v0, _167_v0, _167_v0)
		var _rhs22 _dafny.Int = ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_169_v2).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_169_v2).Cardinality()))).Uint32()).(_dafny.Sequence), _168_v1)).Cardinality()))).Plus(Companion_Default___.Fm3((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), _170_v3, _173_globalState))
		_ = _rhs22
		var _rhs23 _dafny.Set = (func() _dafny.Set {
			var _coll19 = _dafny.NewBuilder()
			_ = _coll19
			for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(270), _dafny.IntOfInt64(-534))); ; {
				_compr_19, _ok19 := _iter19()
				if !_ok19 {
					break
				}
				var _209_v38 _dafny.Int
				_209_v38 = interface{}(_compr_19).(_dafny.Int)
				if ((_dafny.IntOfInt64(270)).Cmp(_209_v38) <= 0) && ((_209_v38).Cmp(_dafny.IntOfInt64(-534)) < 0) {
					_coll19.Add((_209_v38).Times(_167_v0))
				}
			}
			return _coll19.ToSet()
		}()).Union(_208_v37)
		_ = _rhs23
		var _lhs18 *GlobalState = _173_globalState
		_ = _lhs18
		_lhs18.F1 = _rhs22
		_208_v37 = _rhs23
		var _210_v39 D2
		_ = _210_v39
		_210_v39 = Companion_D2_.Create_DC6_((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence))
		if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_196_v25, (_210_v39).Dtor_cf6()), _196_v25) {
			var _211_v40 _dafny.Map
			_ = _211_v40
			_211_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool))
			var _212_v41 _dafny.Map
			_ = _212_v41
			_212_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, (func() bool {
				if (_211_v40).Contains(_167_v0) {
					return (_211_v40).Get(_167_v0).(bool)
				}
				return _170_v3
			})())
			var _213_v42 *C7
			_ = _213_v42
			var _nw20 *C7 = New_C7_()
			_ = _nw20
			_nw20.Ctor__(_167_v0, ((_212_v41).Merge(_212_v41)).Cardinality(), _191_v20, (func() bool {
				if (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool) {
					return _170_v3
				}
				return (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool)
			})())
			_213_v42 = _nw20
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_200_v29), 0))
			_ = _index24
			(_200_v29).ArraySet1((_213_v42).F9(), (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_200_v29), 0))
			_ = _index25
			(_200_v29).ArraySet1(_167_v0, (_index25).Int())
			var _214_v43 _dafny.Array
			_ = _214_v43
			var _nw21 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
			_ = _nw21
			_214_v43 = _nw21
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_214_v43), 0))
			_ = _index26
			(_214_v43).ArraySet1(Companion_Default___.Fm2(_193_v22, !(_170_v3), _173_globalState), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_214_v43), 0))
			_ = _index27
			(_214_v43).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool)), _172_v5), (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _index28
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_200_v29), 0))
			_ = _index29
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _index30
			var _rhs24 bool = _193_v22
			_ = _rhs24
			var _rhs25 _dafny.Int = _dafny.IntOfInt64(-536)
			_ = _rhs25
			var _rhs26 _dafny.Int = (_213_v42).Fm8(!(((_213_v42).Fm7(_193_v22, _173_globalState)).Cmp(_167_v0) <= 0), (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), _193_v22, _173_globalState)
			_ = _rhs26
			var _rhs27 bool = _170_v3
			_ = _rhs27
			var _lhs19 _dafny.Array = _190_v19
			_ = _lhs19
			var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _lhs20
			var _lhs21 _dafny.Array = _200_v29
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(164), _dafny.ArrayLen((_200_v29), 0))
			_ = _lhs22
			var _lhs23 _dafny.Array = _190_v19
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _lhs24
			(_lhs19).ArraySet1(_rhs24, (_lhs20).Int())
			(_lhs21).ArraySet1(_rhs25, (_lhs22).Int())
			_167_v0 = _rhs26
			(_lhs23).ArraySet1(_rhs27, (_lhs24).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _index31
			(_190_v19).ArraySet1(_193_v22, (_index31).Int())
		} else {
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _index32
			(_190_v19).ArraySet1(_170_v3, (_index32).Int())
			(_173_globalState).F1 = Companion_Default___.Fm3(_193_v22, _193_v22, _173_globalState)
			var _215_v44 *C6
			_ = _215_v44
			var _nw22 *C6 = New_C6_()
			_ = _nw22
			_nw22.Ctor__()
			_215_v44 = _nw22
			var _216_v45 _dafny.Array
			_ = _216_v45
			var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw23
			_216_v45 = _nw23
			var _217_v46 _dafny.Sequence
			_ = _217_v46
			_217_v46 = _dafny.SeqOf(_216_v45)
			var _218_v47 _dafny.Array
			_ = _218_v47
			var _nwElement0_6 _dafny.Array = (func() _dafny.Array {
				if _193_v22 {
					return _216_v45
				}
				return _216_v45
			})()
			_ = _nwElement0_6
			var _nw24 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(7))
			_ = _nw24
			(_nw24).ArraySet1(_nwElement0_6, 0)
			(_nw24).ArraySet1(_216_v45, 1)
			(_nw24).ArraySet1((_217_v46).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_217_v46).Cardinality()))).Uint32()).(_dafny.Array), 2)
			(_nw24).ArraySet1(_216_v45, 3)
			(_nw24).ArraySet1(_216_v45, 4)
			(_nw24).ArraySet1(_216_v45, 5)
			(_nw24).ArraySet1(_216_v45, 6)
			_218_v47 = _nw24
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_218_v47), 0))
			_ = _index33
			(_218_v47).ArraySet1(_216_v45, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_218_v47), 0))
			_ = _index34
			(_218_v47).ArraySet1((_217_v46).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_167_v0), _dafny.IntOfUint32((_217_v46).Cardinality()))).Uint32()).(_dafny.Array), (_index34).Int())
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen((_190_v19), 0))
			_ = _index35
			(_190_v19).ArraySet1(!(!(_170_v3)), (_index35).Int())
		}
	} else {
		var _219_v48 _dafny.Set
		_ = _219_v48
		_219_v48 = _dafny.SetOf(_167_v0, _dafny.IntOfUint32(((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		_167_v0 = Companion_Default___.Fm3((false) && (_193_v22), (_219_v48).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(306), _167_v0)), _173_globalState)
		if (_170_v3) || (Companion_Default___.Fm0(_193_v22, _173_globalState)) {
			var _220_v49 _dafny.Array
			_ = _220_v49
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_4
			var _nw25 _dafny.Array
			_ = _nw25
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw25 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Map = (func(_221_v20 _dafny.CodePoint, _222_v0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_223_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_221_v20, _222_v0)
					}
				})(_191_v20, _167_v0)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw25 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw25).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw25).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_220_v49 = _nw25
			var _224_v50 _dafny.Sequence
			_ = _224_v50
			_224_v50 = _dafny.SeqOf(_220_v49, _220_v49, _220_v49, _220_v49, _220_v49)
			var _225_v51 _dafny.Map
			_ = _225_v51
			_225_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _dafny.IntOfUint32((_172_v5).Cardinality()))
			var _226_v52 _dafny.Array
			_ = _226_v52
			var _nwElement0_7 _dafny.Array = _220_v49
			_ = _nwElement0_7
			var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(5))
			_ = _nw26
			(_nw26).ArraySet1(_nwElement0_7, 0)
			(_nw26).ArraySet1(_220_v49, 1)
			(_nw26).ArraySet1((_224_v50).Select((Companion_Default___.SafeIndex((_225_v51).Cardinality(), _dafny.IntOfUint32((_224_v50).Cardinality()))).Uint32()).(_dafny.Array), 2)
			(_nw26).ArraySet1(_220_v49, 3)
			(_nw26).ArraySet1(_220_v49, 4)
			_226_v52 = _nw26
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_226_v52), 0))
			_ = _index36
			(_226_v52).ArraySet1(_220_v49, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_226_v52), 0))
			_ = _index37
			var _rhs28 _dafny.Int = Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus(_167_v0)).Times(_167_v0), _167_v0)
			_ = _rhs28
			var _rhs29 _dafny.Array = _220_v49
			_ = _rhs29
			var _lhs25 *GlobalState = _173_globalState
			_ = _lhs25
			var _lhs26 _dafny.Array = _226_v52
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_226_v52), 0))
			_ = _lhs27
			_lhs25.F1 = _rhs28
			(_lhs26).ArraySet1(_rhs29, (_lhs27).Int())
			var _227_v53 *C6
			_ = _227_v53
			var _nw27 *C6 = New_C6_()
			_ = _nw27
			_nw27.Ctor__()
			_227_v53 = _nw27
			var _228_v54 _dafny.Map
			_ = _228_v54
			_228_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v53, _193_v22)
			var _229_v55 _dafny.Map
			_ = _229_v55
			_229_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_228_v54).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_227_v53, _170_v3)), _dafny.IntOfInt64(719))
			(_173_globalState).F1 = (_229_v55).Cardinality()
			_167_v0 = (_dafny.Zero).Minus(_167_v0)
			_170_v3 = !(Companion_Default___.Fm0((_227_v53).Fm6(_dafny.SeqOf(_167_v0), _173_globalState), _173_globalState)) || (_170_v3)
			var _230_v56 _dafny.Int
			_ = _230_v56
			var _out4 _dafny.Int
			_ = _out4
			_out4 = (_227_v53).M3(_173_globalState)
			_230_v56 = _out4
		} else {
			_219_v48 = _219_v48
			var _231_v57 _dafny.Map
			_ = _231_v57
			_231_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v22, _167_v0)
			var _232_v58 _dafny.MultiSet
			_ = _232_v58
			_232_v58 = _dafny.MultiSetOf(_231_v57, _231_v57)
			var _233_v60 _dafny.MultiSet
			_ = _233_v60
			_233_v60 = _dafny.MultiSetOf((_231_v57).Cardinality())
			var _234_v61 _dafny.Map
			_ = _234_v61
			_234_v61 = func() _dafny.Map {
				var _coll20 = _dafny.NewMapBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate((_233_v60).Elements()); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _235_v59 _dafny.Int
					_235_v59 = interface{}(_compr_20).(_dafny.Int)
					if (_233_v60).Contains(_235_v59) {
						_coll20.Add((_235_v59).Plus(_167_v0), _193_v22)
					}
				}
				return _coll20.ToMap()
			}()
			var _236_v62 _dafny.Map
			_ = _236_v62
			_236_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v61, _170_v3)
			var _237_v63 _dafny.Sequence
			_ = _237_v63
			_237_v63 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, (_236_v62).Cardinality()), _231_v57)
			var _238_v64 _dafny.Sequence
			_ = _238_v64
			_238_v64 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_237_v63))
			var _239_v65 _dafny.Map
			_ = _239_v65
			_239_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _232_v58)
			var _rhs30 _dafny.MultiSet = (_232_v58).Union((func() _dafny.MultiSet {
				if _193_v22 {
					return (_238_v64).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_167_v0), _dafny.IntOfUint32((_238_v64).Cardinality()))).Uint32()).(_dafny.MultiSet)
				}
				return (func() _dafny.MultiSet {
					if (_239_v65).Contains(_167_v0) {
						return (_239_v65).Get(_167_v0).(_dafny.MultiSet)
					}
					return _232_v58
				})()
			})())
			_ = _rhs30
			var _rhs31 _dafny.Int = (_dafny.IntOfUint32((_172_v5).Cardinality())).Times(_167_v0)
			_ = _rhs31
			_232_v58 = _rhs30
			_167_v0 = _rhs31
			var _240_v66 _dafny.Array
			_ = _240_v66
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_5
			var _nw28 _dafny.Array
			_ = _nw28
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw28 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_241_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_242_i3 _dafny.Int) _dafny.CodePoint {
						return _241_v20
					}
				})(_191_v20)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw28 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw28).ArraySet1CodePoint(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw28).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_240_v66 = _nw28
			var _243_v67 _dafny.Map
			_ = _243_v67
			_243_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v66, _170_v3)
			_243_v67 = (_243_v67).Update(_240_v66, (_167_v0).Cmp((_233_v60).Cardinality()) >= 0)
			var _244_v68 *C6
			_ = _244_v68
			var _nw29 *C6 = New_C6_()
			_ = _nw29
			_nw29.Ctor__()
			_244_v68 = _nw29
			var _245_v69 _dafny.Map
			_ = _245_v69
			_245_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, _170_v3)
			var _246_v70 _dafny.MultiSet
			_ = _246_v70
			_246_v70 = _dafny.MultiSetOf((_245_v69).Merge(_245_v69), (_245_v69).Update(_170_v3, _170_v3))
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_190_v19), 0))
			_ = _index38
			(_190_v19).ArraySet1(!(_245_v69).Contains(_170_v3), (_index38).Int())
			var _247_v71 _dafny.Map
			_ = _247_v71
			_247_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, (_219_v48).Cardinality())
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_190_v19), 0))
			_ = _index39
			var _rhs32 _dafny.MultiSet = _dafny.MultiSetOf(_245_v69, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v22, _170_v3)).Update(_170_v3, true)).Update(_193_v22, _193_v22), _245_v69)
			_ = _rhs32
			var _rhs33 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_168_v1, _168_v1)
			_ = _rhs33
			var _rhs34 bool = (func() bool {
				if _193_v22 {
					return (_170_v3) == (_170_v3)
				}
				return _170_v3
			})()
			_ = _rhs34
			var _rhs35 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_167_v0, (func() _dafny.Int {
				if (_247_v71).Contains(_dafny.IntOfInt64(929)) {
					return (_247_v71).Get(_dafny.IntOfInt64(929)).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_167_v0)
			})()))
			_ = _rhs35
			var _lhs28 _dafny.Array = _190_v19
			_ = _lhs28
			var _lhs29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_190_v19), 0))
			_ = _lhs29
			var _lhs30 *GlobalState = _173_globalState
			_ = _lhs30
			_246_v70 = _rhs32
			_168_v1 = _rhs33
			(_lhs28).ArraySet1(_rhs34, (_lhs29).Int())
			_lhs30.F1 = _rhs35
		}
		var _248_v72 _dafny.Map
		_ = _248_v72
		_248_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _170_v3)
		var _249_v73 _dafny.MultiSet
		_ = _249_v73
		_249_v73 = _dafny.MultiSetOf((_248_v72).Cardinality(), _dafny.IntOfUint32((_168_v1).Cardinality()))
		var _250_v74 D0
		_ = _250_v74
		var _251_v75 bool
		_ = _251_v75
		var _out5 D0
		_ = _out5
		var _out6 bool
		_ = _out6
		_out5, _out6 = Companion_Default___.M0(_190_v19, _191_v20, (_249_v73).IsSubsetOf(_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32(((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)).Cardinality())), Companion_Default___.Fm3(_193_v22, _170_v3, _173_globalState))), _191_v20, _173_globalState)
		_250_v74 = _out5
		_251_v75 = _out6
		_248_v72 = (_248_v72).Update(_dafny.IntOfInt64(253), _251_v75)
		var _252_v76 _dafny.Map
		_ = _252_v76
		_252_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _dafny.IntOfInt64(-98))
		var _253_v77 D0
		_ = _253_v77
		var _254_v78 bool
		_ = _254_v78
		var _out7 D0
		_ = _out7
		var _out8 bool
		_ = _out8
		_out7, _out8 = Companion_Default___.M0(_190_v19, _191_v20, (_170_v3) && (_193_v22), Companion_Default___.Fm36(_167_v0, _167_v0, _167_v0, _252_v76, _173_globalState), _173_globalState)
		_253_v77 = _out7
		_254_v78 = _out8
	}
	_167_v0 = (func() _dafny.Int {
		if _170_v3 {
			return _167_v0
		}
		return _167_v0
	})()
	var _255_v79 _dafny.Map
	_ = _255_v79
	_255_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, _189_v18)
	var _256_v80 _dafny.Map
	_ = _256_v80
	_256_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_193_v22) || (_170_v3), _255_v79)
	var _257_v81 _dafny.Array
	_ = _257_v81
	var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
	_ = _nw30
	_257_v81 = _nw30
	var _258_v82 _dafny.Sequence
	_ = _258_v82
	_258_v82 = _dafny.SeqOf(_257_v81, _257_v81, _257_v81, _257_v81)
	var _259_v83 *C8
	_ = _259_v83
	var _nw31 *C8 = New_C8_()
	_ = _nw31
	_nw31.Ctor__(_190_v19, _258_v82, _191_v20, _193_v22)
	_259_v83 = _nw31
	var _260_v84 _dafny.MultiSet
	_ = _260_v84
	_260_v84 = _dafny.MultiSetOf(_170_v3)
	var _261_v85 D22
	_ = _261_v85
	_261_v85 = Companion_D22_.Create_DC48_(_259_v83)
	var _rhs36 _dafny.Map = (func() _dafny.Map {
		if _170_v3 {
			return _256_v80
		}
		return _256_v80
	})()
	_ = _rhs36
	var _rhs37 _dafny.Int = (((_260_v84).Cardinality()).Minus(_167_v0)).Plus(_167_v0)
	_ = _rhs37
	var _rhs38 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ibtmhviyp")
	_ = _rhs38
	var _rhs39 *C8 = (_261_v85).Dtor_cf65()
	_ = _rhs39
	_256_v80 = _rhs36
	_167_v0 = _rhs37
	_196_v25 = _rhs38
	_259_v83 = _rhs39
	var _262_v86 _dafny.Map
	_ = _262_v86
	_262_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_167_v0), _dafny.SeqOf(_dafny.IntOfInt64(598), _167_v0, _167_v0, _167_v0))
	var _263_v87 D19
	_ = _263_v87
	_263_v87 = Companion_D19_.Create_DC40_(_170_v3, (_259_v83).Fm8(_193_v22, _170_v3, _193_v22, _173_globalState), _262_v86)
	var _264_v88 _dafny.Map
	_ = _264_v88
	_264_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, _167_v0)
	var _265_v89 _dafny.MultiSet
	_ = _265_v89
	_265_v89 = _dafny.MultiSetOf(_264_v88)
	if ((_263_v87).Dtor_cf52()).Cmp((_265_v89).Cardinality()) >= 0 {
		var _266_v90 _dafny.Set
		_ = _266_v90
		_266_v90 = _dafny.SetOf(_257_v81, _257_v81)
		var _267_v91 D22
		_ = _267_v91
		_267_v91 = Companion_D22_.Create_DC49_(_266_v90)
		var _source7 D22 = _267_v91
		_ = _source7
		if _source7.Is_DC49() {
			var _268___mcc_h0 _dafny.Set = _source7.Get_().(D22_DC49).Cf66
			_ = _268___mcc_h0
			var _269_cf66 _dafny.Set = _268___mcc_h0
			_ = _269_cf66
			_170_v3 = _193_v22
			var _270_v92 _dafny.Set
			_ = _270_v92
			_270_v92 = _dafny.SetOf(_170_v3)
			var _271_v93 _dafny.Map
			_ = _271_v93
			_271_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v92, _dafny.IntOfUint32(((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)).Cardinality()))
			var _272_v94 bool
			_ = _272_v94
			var _273_v95 _dafny.Sequence
			_ = _273_v95
			var _out9 bool
			_ = _out9
			var _out10 _dafny.Sequence
			_ = _out10
			_out9, _out10 = (_259_v83).M2(_193_v22, ((_271_v93).Cardinality()).Plus(_167_v0), _196_v25, _173_globalState)
			_272_v94 = _out9
			_273_v95 = _out10
			(_173_globalState).F1 = (_167_v0).Plus(_167_v0)
			(_173_globalState).F1 = _167_v0
		} else {
			var _274___mcc_h1 *C8 = _source7.Get_().(D22_DC48).Cf65
			_ = _274___mcc_h1
			var _275_cf65 *C8 = _274___mcc_h1
			_ = _275_cf65
			var _276_v96 _dafny.Set
			_ = _276_v96
			_276_v96 = _dafny.SetOf(((_dafny.Zero).Minus(_167_v0)).Minus(_167_v0))
			var _277_v97 _dafny.MultiSet
			_ = _277_v97
			_277_v97 = _dafny.MultiSetOf((_dafny.SetOf(_167_v0)).Cardinality())
			_276_v96 = _dafny.SetOf((func() _dafny.Int {
				if _170_v3 {
					return (_277_v97).Cardinality()
				}
				return (func() _dafny.Int {
					if (_260_v84).Contains(_193_v22) {
						return (_260_v84).Multiplicity(_193_v22)
					}
					return (_259_v83).Fm8(_170_v3, _170_v3, _193_v22, _173_globalState)
				})()
			})(), _167_v0, _167_v0, (_275_cf65).Fm8(_170_v3, _170_v3, _170_v3, _173_globalState))
			(_275_cf65).M5(Companion_Default___.Fm1(_167_v0, _170_v3, _167_v0, _173_globalState), _167_v0, _173_globalState)
			(_173_globalState).F1 = (Companion_Default___.SafeDivisionInt(_167_v0, (_dafny.Zero).Minus(_167_v0))).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence), _196_v25)).Cardinality()))
			var _278_v98 *C0
			_ = _278_v98
			var _nw32 *C0 = New_C0_()
			_ = _nw32
			_nw32.Ctor__(_193_v22, _dafny.Companion_Sequence_.Concatenate(_169_v2, _169_v2))
			_278_v98 = _nw32
		}
		_168_v1 = _168_v1
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))
		_ = _index40
		(_257_v81).ArraySet1(_167_v0, (_index40).Int())
		var _279_v99 _dafny.Map
		_ = _279_v99
		_279_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _193_v22)
		var _280_v100 _dafny.Set
		_ = _280_v100
		_280_v100 = _dafny.SetOf(_193_v22)
		var _281_v101 *C5
		_ = _281_v101
		var _nw33 *C5 = New_C5_()
		_ = _nw33
		_nw33.Ctor__((_280_v100).Cardinality(), _191_v20, _193_v22)
		_281_v101 = _nw33
		var _282_v102 _dafny.Sequence
		_ = _282_v102
		_282_v102 = _dafny.SeqOf(_281_v101, _281_v101)
		var _283_v103 _dafny.Map
		_ = _283_v103
		_283_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_279_v99).Contains(_167_v0) {
				return (_279_v99).Get(_167_v0).(bool)
			}
			return !(_193_v22)
		})(), _dafny.Companion_Sequence_.Update(_282_v102, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_167_v0), _dafny.IntOfUint32((_282_v102).Cardinality()))).Uint32(), _281_v101))
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))
		_ = _index41
		var _rhs40 _dafny.Int = (_dafny.Zero).Minus(_167_v0)
		_ = _rhs40
		var _rhs41 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_168_v1, _dafny.Companion_Sequence_.Concatenate(_168_v1, _168_v1))
		_ = _rhs41
		var _rhs42 _dafny.Map = (_283_v103).Merge((_283_v103).Merge((_283_v103).Update(_193_v22, _282_v102)))
		_ = _rhs42
		var _rhs43 _dafny.Int = _167_v0
		_ = _rhs43
		var _lhs31 _dafny.Array = _257_v81
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))
		_ = _lhs32
		var _lhs33 *C5 = _281_v101
		_ = _lhs33
		(_lhs31).ArraySet1(_rhs40, (_lhs32).Int())
		_168_v1 = _rhs41
		_283_v103 = _rhs42
		_lhs33.F15 = _rhs43
		var _284_v104 *C0
		_ = _284_v104
		var _nw34 *C0 = New_C0_()
		_ = _nw34
		_nw34.Ctor__((_dafny.IntOfInt64(642)).Cmp((func() _dafny.Int {
			if (_264_v88).Contains(!(_193_v22)) {
				return (_264_v88).Get(!(_193_v22)).(_dafny.Int)
			}
			return _281_v101.F15
		})()) > 0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_169_v2, _169_v2), (Companion_Default___.SafeIndex(_281_v101.F15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_169_v2, _169_v2)).Cardinality()))).Uint32(), _168_v1))
		_284_v104 = _nw34
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
		_ = _index42
		var _rhs44 *C0 = _284_v104
		_ = _rhs44
		var _rhs45 _dafny.Int = (_dafny.Zero).Minus(_167_v0)
		_ = _rhs45
		var _rhs46 _dafny.Sequence = _196_v25
		_ = _rhs46
		var _lhs34 *GlobalState = _173_globalState
		_ = _lhs34
		var _lhs35 _dafny.Array = _195_v24
		_ = _lhs35
		var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
		_ = _lhs36
		_284_v104 = _rhs44
		_lhs34.F1 = _rhs45
		(_lhs35).ArraySet1(_rhs46, (_lhs36).Int())
		var _285_v105 _dafny.MultiSet
		_ = _285_v105
		_285_v105 = _dafny.MultiSetOf(_191_v20)
		var _286_v106 _dafny.Array
		_ = _286_v106
		var _nwElement0_8 _dafny.Int = _281_v101.F15
		_ = _nwElement0_8
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(29))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_8, 0)
		(_nw35).ArraySet1((_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), 1)
		(_nw35).ArraySet1((_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), 2)
		(_nw35).ArraySet1(_281_v101.F15, 3)
		(_nw35).ArraySet1(_167_v0, 4)
		(_nw35).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_168_v1).Cardinality())), 5)
		(_nw35).ArraySet1(_281_v101.F15, 6)
		(_nw35).ArraySet1(_281_v101.F15, 7)
		(_nw35).ArraySet1(_dafny.IntOfUint32(((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)).Cardinality()), 8)
		(_nw35).ArraySet1(_281_v101.F15, 9)
		(_nw35).ArraySet1(_167_v0, 10)
		(_nw35).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), 11)
		(_nw35).ArraySet1(_281_v101.F15, 12)
		(_nw35).ArraySet1((func() _dafny.Int {
			if (_285_v105).Contains(_191_v20) {
				return (_285_v105).Multiplicity(_191_v20)
			}
			return _dafny.IntOfInt64(73)
		})(), 13)
		(_nw35).ArraySet1(_281_v101.F15, 14)
		(_nw35).ArraySet1((_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), 15)
		(_nw35).ArraySet1(_281_v101.F15, 16)
		(_nw35).ArraySet1(_dafny.IntOfInt64(864), 17)
		(_nw35).ArraySet1(_dafny.IntOfInt64(933), 18)
		(_nw35).ArraySet1(_281_v101.F15, 19)
		(_nw35).ArraySet1((_279_v99).Cardinality(), 20)
		(_nw35).ArraySet1((_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), 21)
		(_nw35).ArraySet1((_dafny.SetOf(_dafny.IntOfUint32((_168_v1).Cardinality()))).Cardinality(), 22)
		(_nw35).ArraySet1(_281_v101.F15, 23)
		(_nw35).ArraySet1(_dafny.IntOfUint32((_168_v1).Cardinality()), 24)
		(_nw35).ArraySet1((_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), 25)
		(_nw35).ArraySet1(_281_v101.F15, 26)
		(_nw35).ArraySet1(_281_v101.F15, 27)
		(_nw35).ArraySet1((_281_v101).Fm8(_170_v3, _170_v3, !(_193_v22), _173_globalState), 28)
		_286_v106 = _nw35
		var _287_v107 D2
		_ = _287_v107
		var _288_v108 _dafny.MultiSet
		_ = _288_v108
		var _out11 D2
		_ = _out11
		var _out12 _dafny.MultiSet
		_ = _out12
		_out11, _out12 = (_259_v83).M4(_286_v106, _172_v5, (_284_v104).F11(), (_257_v81).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_257_v81), 0))).Int()).(_dafny.Int), _173_globalState)
		_287_v107 = _out11
		_288_v108 = _out12
	} else {
		var _289_v109 _dafny.Set
		_ = _289_v109
		_289_v109 = _dafny.SetOf(_259_v83.F7)
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))
		_ = _index43
		(_190_v19).ArraySet1(!(_289_v109).Contains(_259_v83.F7), (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))
		_ = _index44
		(_190_v19).ArraySet1(_170_v3, (_index44).Int())
		var _source8 D0 = _192_v21
		_ = _source8
		if _source8.Is_DC1() {
			(_173_globalState).F1 = (_167_v0).Plus((_dafny.Zero).Minus((_167_v0).Plus((_dafny.Zero).Minus(_167_v0))))
			var _290_v110 _dafny.Map
			_ = _290_v110
			_290_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v22, _191_v20)
			var _291_v111 _dafny.Set
			_ = _291_v111
			_291_v111 = _dafny.SetOf(_290_v110)
			(_259_v83).M5(_196_v25, (_291_v111).Cardinality(), _173_globalState)
			var _292_v112 _dafny.Map
			_ = _292_v112
			_292_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_170_v3, _193_v22)
			var _293_v113 _dafny.MultiSet
			_ = _293_v113
			_293_v113 = _dafny.MultiSetOf(_292_v112)
			(_173_globalState).F1 = ((_293_v113).Cardinality()).Times(_167_v0)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))
			_ = _index45
			(_190_v19).ArraySet1(false, (_index45).Int())
		} else if _source8.Is_DC2() {
			var _294___mcc_h2 _dafny.Map = _source8.Get_().(D0_DC2).Cf1
			_ = _294___mcc_h2
			var _295___mcc_h3 _dafny.CodePoint = _source8.Get_().(D0_DC2).Cf2
			_ = _295___mcc_h3
			var _296_cf2 _dafny.CodePoint = _295___mcc_h3
			_ = _296_cf2
			var _297_cf1 _dafny.Map = _294___mcc_h2
			_ = _297_cf1
			var _298_v114 *C3
			_ = _298_v114
			var _nw36 *C3 = New_C3_()
			_ = _nw36
			_nw36.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_196_v25, (_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)), true)
			_298_v114 = _nw36
			var _299_v115 _dafny.Set
			_ = _299_v115
			_299_v115 = _dafny.SetOf(_170_v3)
			var _300_v116 _dafny.Map
			_ = _300_v116
			_300_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _299_v115)
			var _301_v117 _dafny.Map
			_ = _301_v117
			_301_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_193_v22), _170_v3)
			var _rhs47 _dafny.Int = Companion_Default___.SafeModuloInt(_167_v0, Companion_Default___.SafeDivisionInt(_167_v0, _167_v0))
			_ = _rhs47
			var _rhs48 _dafny.CodePoint = Companion_Default___.Fm22((_299_v115).Difference(((func() _dafny.Set {
				if (_300_v116).Contains(_167_v0) {
					return (_300_v116).Get(_167_v0).(_dafny.Set)
				}
				return _299_v115
			})())), Companion_Default___.SafeDivisionInt((_259_v83).Fm7((_298_v114).F14(), _173_globalState), _167_v0), (_259_v83).Fm10(_260_v84, _301_v117, _167_v0, _173_globalState), (Companion_Default___.Fm48(_167_v0, !(true), _173_globalState)).IsProperSubsetOf(_dafny.MultiSetOf(true)), _173_globalState)
			_ = _rhs48
			var _rhs49 _dafny.Int = _167_v0
			_ = _rhs49
			var _lhs37 *GlobalState = _173_globalState
			_ = _lhs37
			var _lhs38 *GlobalState = _173_globalState
			_ = _lhs38
			_lhs37.F1 = _rhs47
			_296_cf2 = _rhs48
			_lhs38.F1 = _rhs49
			var _rhs50 bool = _170_v3
			_ = _rhs50
			var _rhs51 _dafny.Sequence = (_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence)
			_ = _rhs51
			var _rhs52 _dafny.Int = _167_v0
			_ = _rhs52
			var _lhs39 *GlobalState = _173_globalState
			_ = _lhs39
			_193_v22 = _rhs50
			_196_v25 = _rhs51
			_lhs39.F1 = _rhs52
			var _302_v118 _dafny.Array
			_ = _302_v118
			var _nw37 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
			_ = _nw37
			_302_v118 = _nw37
			_302_v118 = _302_v118
		} else if _source8.Is_DC0() {
			var _303___mcc_h4 bool = _source8.Get_().(D0_DC0).Cf0
			_ = _303___mcc_h4
			var _304_cf0 bool = _303___mcc_h4
			_ = _304_cf0
			var _305_v119 T1
			_ = _305_v119
			var _nw38 *C4 = New_C4_()
			_ = _nw38
			_nw38.Ctor__()
			_305_v119 = _nw38
			var _306_v120 _dafny.Sequence
			_ = _306_v120
			_306_v120 = _dafny.SeqOf(_305_v119)
			var _307_v121 _dafny.MultiSet
			_ = _307_v121
			_307_v121 = _dafny.MultiSetOf((_dafny.MultiSetFromSeq(_172_v5)).Cardinality(), _167_v0)
			var _308_v122 _dafny.Set
			_ = _308_v122
			_308_v122 = _dafny.SetOf(_167_v0, (_307_v121).Cardinality())
			var _309_v123 _dafny.Array
			_ = _309_v123
			var _nwElement0_9 T1 = _305_v119
			_ = _nwElement0_9
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(21))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_9, 0)
			(_nw39).ArraySet1(_305_v119, 1)
			(_nw39).ArraySet1(_305_v119, 2)
			(_nw39).ArraySet1(_305_v119, 3)
			(_nw39).ArraySet1(_305_v119, 4)
			(_nw39).ArraySet1(_305_v119, 5)
			(_nw39).ArraySet1(_305_v119, 6)
			(_nw39).ArraySet1((_306_v120).Select((Companion_Default___.SafeIndex((_308_v122).Cardinality(), _dafny.IntOfUint32((_306_v120).Cardinality()))).Uint32()).(T1), 7)
			(_nw39).ArraySet1(_305_v119, 8)
			(_nw39).ArraySet1(_305_v119, 9)
			(_nw39).ArraySet1(_305_v119, 10)
			(_nw39).ArraySet1(_305_v119, 11)
			(_nw39).ArraySet1(_305_v119, 12)
			(_nw39).ArraySet1(_305_v119, 13)
			(_nw39).ArraySet1(_305_v119, 14)
			(_nw39).ArraySet1(_305_v119, 15)
			(_nw39).ArraySet1(_305_v119, 16)
			(_nw39).ArraySet1(_305_v119, 17)
			(_nw39).ArraySet1(_305_v119, 18)
			(_nw39).ArraySet1(_305_v119, 19)
			(_nw39).ArraySet1(_305_v119, 20)
			_309_v123 = _nw39
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_309_v123), 0))
			_ = _index46
			(_309_v123).ArraySet1((func() T1 {
				if _170_v3 {
					return _305_v119
				}
				return _305_v119
			})(), (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_309_v123), 0))
			_ = _index47
			var _nw40 *C4 = New_C4_()
			_ = _nw40
			_nw40.Ctor__()
			(_309_v123).ArraySet1(_nw40, (_index47).Int())
			var _310_v124 _dafny.Sequence
			_ = _310_v124
			var _311_v125 bool
			_ = _311_v125
			var _312_v126 bool
			_ = _312_v126
			var _out13 _dafny.Sequence
			_ = _out13
			var _out14 bool
			_ = _out14
			var _out15 bool
			_ = _out15
			_out13, _out14, _out15 = (_259_v83).M6(true, (_259_v83).Fm10(_260_v84, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_193_v22), (Companion_D1_.Create_DC5_(_193_v22)).Dtor_cf5()), _167_v0, _173_globalState), _173_globalState)
			_310_v124 = _out13
			_311_v125 = _out14
			_312_v126 = _out15
			var _313_v127 _dafny.MultiSet
			_ = _313_v127
			_313_v127 = _dafny.MultiSetOf(_172_v5, _172_v5)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))
			_ = _index48
			(_190_v19).ArraySet1(((_313_v127).Update(_172_v5, Companion_Default___.Abs(_167_v0))).Contains(_172_v5), (_index48).Int())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_257_v81), 0))
			_ = _index49
			(_257_v81).ArraySet1(_167_v0, (_index49).Int())
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_257_v81), 0))
			_ = _index50
			(_257_v81).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if _312_v126 {
					return _dafny.IntOfInt64(110)
				}
				return _dafny.IntOfInt64(438)
			})(), _167_v0), (_index50).Int())
		} else {
			var _314___mcc_h5 D0 = _source8.Get_().(D0_DC3).Cf3
			_ = _314___mcc_h5
			var _315_cf3 D0 = _314___mcc_h5
			_ = _315_cf3
			var _316_v128 _dafny.Map
			_ = _316_v128
			_316_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_257_v81, _170_v3)
			var _317_v129 *C5
			_ = _317_v129
			var _nw41 *C5 = New_C5_()
			_ = _nw41
			_nw41.Ctor__(((func() _dafny.Map {
				if (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool) {
					return (_316_v128).Update(_257_v81, _193_v22)
				}
				return _316_v128
			})()).Cardinality(), (_196_v25).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_167_v0), _dafny.IntOfUint32((_196_v25).Cardinality()))).Uint32()).(_dafny.CodePoint), !(_170_v3))
			_317_v129 = _nw41
			var _318_v130 bool
			_ = _318_v130
			var _319_v131 _dafny.Sequence
			_ = _319_v131
			var _out16 bool
			_ = _out16
			var _out17 _dafny.Sequence
			_ = _out17
			_out16, _out17 = (_259_v83).M2(_193_v22, _317_v129.F15, _dafny.UnicodeSeqOfUtf8Bytes("ylqpcjtj"), _173_globalState)
			_318_v130 = _out16
			_319_v131 = _out17
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_190_v19), 0))
			_ = _index51
			(_190_v19).ArraySet1(_170_v3, (_index51).Int())
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_257_v81), 0))
			_ = _index52
			(_257_v81).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_172_v5, _172_v5)).Cardinality()), (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_257_v81), 0))
			_ = _index53
			(_257_v81).ArraySet1((_167_v0).Times((_317_v129.F15).Times(_167_v0)), (_index53).Int())
		}
		_191_v20 = _191_v20
		var _nw42 *C8 = New_C8_()
		_ = _nw42
		_nw42.Ctor__(_259_v83.F7, _258_v82, _191_v20, _170_v3)
		_259_v83 = _nw42
		_193_v22 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("asmrl"), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pjcwgjki")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("asmrl")).Cardinality()))).Uint32(), _191_v20), _dafny.CodePoint('f'))
	}
	var _hi0 _dafny.Int = (_167_v0).Minus(_167_v0)
	_ = _hi0
	for _320_i4 := _167_v0; _320_i4.Cmp(_hi0) < 0; _320_i4 = _320_i4.Plus(_dafny.One) {
		_167_v0 = _167_v0
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_190_v19), 0))
		_ = _index54
		(_190_v19).ArraySet1(_170_v3, (_index54).Int())
		var _321_v132 _dafny.MultiSet
		_ = _321_v132
		_321_v132 = _dafny.MultiSetOf((_320_i4).Plus((_260_v84).Cardinality()))
		var _322_v133 T3
		_ = _322_v133
		var _nw43 *C5 = New_C5_()
		_ = _nw43
		_nw43.Ctor__((_dafny.Zero).Minus(_167_v0), _191_v20, false)
		_322_v133 = _nw43
		var _323_v134 _dafny.Map
		_ = _323_v134
		_323_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_i4, _322_v133)
		var _324_v135 _dafny.Set
		_ = _324_v135
		_324_v135 = _dafny.SetOf(_322_v133, _322_v133, _322_v133, _322_v133, (func() T3 {
			if (_323_v134).Contains(_dafny.IntOfInt64(-341)) {
				return (_323_v134).Get(_dafny.IntOfInt64(-341)).(T3)
			}
			return _322_v133
		})())
		var _325_v136 _dafny.Sequence
		_ = _325_v136
		_325_v136 = _dafny.SeqOf(_321_v132, _321_v132)
		var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_190_v19), 0))
		_ = _index55
		var _rhs53 bool = _170_v3
		_ = _rhs53
		var _rhs54 _dafny.MultiSet = ((_dafny.MultiSetOf((_259_v83).Fm4(_172_v5, _320_i4, _320_i4, _dafny.MultiSetOf(_167_v0), _173_globalState), _320_i4, _167_v0, (_dafny.Zero).Minus(_167_v0), (_324_v135).Cardinality())).Union(_321_v132)).Union((_325_v136).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_196_v25).Cardinality()), _dafny.IntOfUint32((_325_v136).Cardinality()))).Uint32()).(_dafny.MultiSet))
		_ = _rhs54
		var _rhs55 _dafny.Int = (func() _dafny.Int {
			if (_322_v133).F6() {
				return (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(983), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(374))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_326_v133 T3) func(_dafny.Int) _dafny.CodePoint {
					return func(_327_i5 _dafny.Int) _dafny.CodePoint {
						return (_326_v133).F5()
					}
				})(_322_v133)))).Cardinality()))))
			}
			return (_322_v133).Fm7(_193_v22, _173_globalState)
		})()
		_ = _rhs55
		var _lhs40 _dafny.Array = _190_v19
		_ = _lhs40
		var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_190_v19), 0))
		_ = _lhs41
		var _lhs42 *GlobalState = _173_globalState
		_ = _lhs42
		(_lhs40).ArraySet1(_rhs53, (_lhs41).Int())
		_321_v132 = _rhs54
		_lhs42.F1 = _rhs55
		var _328_v137 _dafny.Array
		_ = _328_v137
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(27))
		_ = _nw44
		_328_v137 = _nw44
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_328_v137), 0))
		_ = _index56
		(_328_v137).ArraySet1(_321_v132, (_index56).Int())
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_328_v137), 0))
		_ = _index57
		var _rhs56 _dafny.MultiSet = (_321_v132).Difference(_321_v132)
		_ = _rhs56
		var _rhs57 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_dafny.IntOfInt64(347), (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), _167_v0, _173_globalState), _196_v25)).Cardinality())
		_ = _rhs57
		var _rhs58 _dafny.Int = (_259_v83).Fm4(_172_v5, _320_i4, _167_v0, _321_v132, _173_globalState)
		_ = _rhs58
		var _lhs43 _dafny.Array = _328_v137
		_ = _lhs43
		var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_328_v137), 0))
		_ = _lhs44
		var _lhs45 *GlobalState = _173_globalState
		_ = _lhs45
		(_lhs43).ArraySet1(_rhs56, (_lhs44).Int())
		_lhs45.F1 = _rhs57
		_167_v0 = _rhs58
		var _329_v138 _dafny.Map
		_ = _329_v138
		_329_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_i4, _320_i4)
		var _330_v139 _dafny.Map
		_ = _330_v139
		_330_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v138, _172_v5)
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_257_v81), 0))
		_ = _index58
		(_257_v81).ArraySet1(_dafny.IntOfInt64(782), (_index58).Int())
		var _331_v140 *C3
		_ = _331_v140
		var _nw45 *C3 = New_C3_()
		_ = _nw45
		_nw45.Ctor__(_193_v22, true)
		_331_v140 = _nw45
		var _332_v141 _dafny.Map
		_ = _332_v141
		_332_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_171_v4, _331_v140)
		var _333_v142 _dafny.Map
		_ = _333_v142
		_333_v142 = _330_v139
		var _334_v143 D20
		_ = _334_v143
		_334_v143 = Companion_D20_.Create_DC43_(_dafny.IntOfInt64(692), (_190_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_190_v19), 0))).Int()).(bool), _167_v0)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_257_v81), 0))
		_ = _index59
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
		_ = _index60
		var _rhs59 _dafny.Int = (_dafny.Zero).Minus(_320_i4)
		_ = _rhs59
		var _rhs60 bool = (_167_v0).Cmp((_332_v141).Cardinality()) != 0
		_ = _rhs60
		var _rhs61 _dafny.Map = (_333_v142)
		_ = _rhs61
		var _rhs62 _dafny.Int = (_334_v143).Dtor_cf56()
		_ = _rhs62
		var _rhs63 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}((func(_335_v133 T3) func(_dafny.Int) _dafny.CodePoint {
			return func(_336_i6 _dafny.Int) _dafny.CodePoint {
				return (_335_v133).F5()
			}
		})(_322_v133)))
		_ = _rhs63
		var _lhs46 *GlobalState = _173_globalState
		_ = _lhs46
		var _lhs47 _dafny.Array = _257_v81
		_ = _lhs47
		var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(439), _dafny.ArrayLen((_257_v81), 0))
		_ = _lhs48
		var _lhs49 _dafny.Array = _195_v24
		_ = _lhs49
		var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))
		_ = _lhs50
		_lhs46.F1 = _rhs59
		_193_v22 = _rhs60
		_330_v139 = _rhs61
		(_lhs47).ArraySet1(_rhs62, (_lhs48).Int())
		(_lhs49).ArraySet1(_rhs63, (_lhs50).Int())
	}
	if (_167_v0).Cmp(_167_v0) > 0 {
		var _337_v144 *C2
		_ = _337_v144
		var _nw46 *C2 = New_C2_()
		_ = _nw46
		_nw46.Ctor__(_191_v20, _170_v3)
		_337_v144 = _nw46
		_337_v144 = _337_v144
		(_173_globalState).F1 = (_dafny.IntOfUint32((_196_v25).Cardinality())).Minus(_167_v0)
		var _338_v145 _dafny.Sequence
		_ = _338_v145
		_338_v145 = _dafny.SeqOf((_195_v24).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(731), _dafny.ArrayLen((_195_v24), 0))).Int()).(_dafny.Sequence))
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_257_v81), 0))
		_ = _index61
		(_257_v81).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("b"), (_338_v145).Select((Companion_Default___.SafeIndex((_337_v144).Fm7(_193_v22, _173_globalState), _dafny.IntOfUint32((_338_v145).Cardinality()))).Uint32()).(_dafny.Sequence))).Cardinality()), (_index61).Int())
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_257_v81), 0))
		_ = _index62
		var _rhs64 _dafny.Int = (_dafny.Zero).Minus((_167_v0).Minus(_167_v0))
		_ = _rhs64
		var _rhs65 _dafny.Int = _167_v0
		_ = _rhs65
		var _rhs66 bool = (_337_v144).Fm9(_167_v0, _171_v4, false, _167_v0, _173_globalState)
		_ = _rhs66
		var _rhs67 _dafny.Int = (_167_v0).Minus((_259_v83).Fm8(_193_v22, true, _170_v3, _173_globalState))
		_ = _rhs67
		var _lhs51 *GlobalState = _173_globalState
		_ = _lhs51
		var _lhs52 _dafny.Array = _257_v81
		_ = _lhs52
		var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_257_v81), 0))
		_ = _lhs53
		_lhs51.F1 = _rhs64
		_167_v0 = _rhs65
		_193_v22 = _rhs66
		(_lhs52).ArraySet1(_rhs67, (_lhs53).Int())
		_167_v0 = _dafny.IntOfInt64(358)
		_193_v22 = true
	} else {
		(_173_globalState).F1 = _167_v0
		_170_v3 = _193_v22
		var _339_v146 *C3
		_ = _339_v146
		var _nw47 *C3 = New_C3_()
		_ = _nw47
		_nw47.Ctor__(!(_170_v3), _dafny.Companion_Sequence_.Equal(_172_v5, _172_v5))
		_339_v146 = _nw47
		_193_v22 = !((_339_v146).F14())
		_195_v24 = (func() _dafny.Array {
			if ((_339_v146).F13()) && (_193_v22) {
				return _195_v24
			}
			return _195_v24
		})()
	}
	var _340_v147 _dafny.Map
	_ = _340_v147
	_340_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, _167_v0)
	_170_v3 = (Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(151))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}((func(_341_v20 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_342_i7 _dafny.Int) _dafny.CodePoint {
			return _341_v20
		}
	})(_191_v20)))).Cardinality()), _167_v0)).Cmp((_dafny.Zero).Minus((_167_v0).Minus((_340_v147).Cardinality()))) <= 0
	var _343_v148 _dafny.Map
	_ = _343_v148
	_343_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_193_v22, _170_v3)
	var _344_v149 *C4
	_ = _344_v149
	var _nw48 *C4 = New_C4_()
	_ = _nw48
	_nw48.Ctor__()
	_344_v149 = _nw48
	var _345_v150 _dafny.Sequence
	_ = _345_v150
	_345_v150 = _dafny.SeqOf(_344_v149)
	var _346_v151 _dafny.MultiSet
	_ = _346_v151
	_346_v151 = _dafny.MultiSetOf(_dafny.IntOfInt64(829), _167_v0, _dafny.IntOfUint32((_345_v150).Cardinality()), _167_v0, _167_v0)
	var _347_i8 _dafny.Int
	_ = _347_i8
	_347_i8 = _dafny.Zero
	{
		for ((_259_v83).Fm4(_172_v5, (_343_v148).Cardinality(), (_dafny.Zero).Minus(_167_v0), _346_v151, _173_globalState)).Cmp(_167_v0) < 0 {
			{
				if (_347_i8).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_347_i8 = (_347_i8).Plus(_dafny.One)
				var _348_v152 _dafny.Array
				_ = _348_v152
				var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw49
				_348_v152 = _nw49
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_348_v152), 0))
				_ = _index63
				(_348_v152).ArraySet1(_190_v19, (_index63).Int())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(363), _dafny.ArrayLen((_348_v152), 0))
				_ = _index64
				(_348_v152).ArraySet1(_259_v83.F7, (_index64).Int())
				var _349_v153 _dafny.Set
				_ = _349_v153
				_349_v153 = _dafny.SetOf(_167_v0)
				var _350_v154 _dafny.Map
				_ = _350_v154
				_350_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_349_v153, (func() _dafny.Int {
					if (_340_v147).Contains(_167_v0) {
						return (_340_v147).Get(_167_v0).(_dafny.Int)
					}
					return _167_v0
				})())
				var _351_v155 _dafny.Set
				_ = _351_v155
				_351_v155 = _dafny.SetOf(_171_v4)
				(_173_globalState).F1 = (func() _dafny.Int {
					if (_260_v84).Contains(true) {
						return (_260_v84).Multiplicity(true)
					}
					return (_167_v0).Minus((func() _dafny.Int {
						if (_350_v154).Contains(_dafny.SetOf(_167_v0, _167_v0)) {
							return (_350_v154).Get(_dafny.SetOf(_167_v0, _167_v0)).(_dafny.Int)
						}
						return (_351_v155).Cardinality()
					})())
				})()
				var _352_v156 D2
				_ = _352_v156
				var _353_v157 _dafny.MultiSet
				_ = _353_v157
				var _out18 D2
				_ = _out18
				var _out19 _dafny.MultiSet
				_ = _out19
				_out18, _out19 = (_259_v83).M4(_257_v81, _dafny.SeqOf(_193_v22), _193_v22, (func() _dafny.Int {
					if (_260_v84).Contains(!(false)) {
						return (_260_v84).Multiplicity(!(false))
					}
					return (_259_v83).Fm4(_172_v5, _167_v0, _dafny.IntOfInt64(903), _346_v151, _173_globalState)
				})(), _173_globalState)
				_352_v156 = _out18
				_353_v157 = _out19
				var _rhs68 bool = (func() bool {
					if Companion_Default___.Fm0((_172_v5).Select((Companion_Default___.SafeIndex(_167_v0, _dafny.IntOfUint32((_172_v5).Cardinality()))).Uint32()).(bool), _173_globalState) {
						return _170_v3
					}
					return (_346_v151).IsSubsetOf(_346_v151)
				})()
				_ = _rhs68
				var _rhs69 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_168_v1)).Intersection(_346_v151)
				_ = _rhs69
				_170_v3 = _rhs68
				_346_v151 = _rhs69
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _354_i9 _dafny.Int
	_ = _354_i9
	_354_i9 = _dafny.Zero
	{
		for _193_v22 {
			{
				if (_354_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_354_i9 = (_354_i9).Plus(_dafny.One)
				_344_v149 = _344_v149
				(_173_globalState).F1 = (_346_v151).Cardinality()
				(_173_globalState).F1 = (_167_v0).Times(_167_v0)
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_257_v81), 0))
				_ = _index65
				(_257_v81).ArraySet1((Companion_Default___.Fm3(false, _170_v3, _173_globalState)).Minus((_dafny.Zero).Minus(_167_v0)), (_index65).Int())
				var _355_v158 _dafny.Array
				_ = _355_v158
				var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(13))
				_ = _nw50
				_355_v158 = _nw50
				var _356_v159 D16
				_ = _356_v159
				_356_v159 = Companion_D16_.Create_DC31_(_193_v22)
				var _357_v160 _dafny.MultiSet
				_ = _357_v160
				_357_v160 = _dafny.MultiSetOf(_356_v159)
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_355_v158), 0))
				_ = _index66
				(_355_v158).ArraySet1(_357_v160, (_index66).Int())
				var _358_v161 D25
				_ = _358_v161
				_358_v161 = Companion_D25_.Create_DC52_(_dafny.MultiSetOf(_356_v159))
				var _359_v162 _dafny.Map
				_ = _359_v162
				_359_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_167_v0, (_358_v161).Dtor_cf69())
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_257_v81), 0))
				_ = _index67
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_355_v158), 0))
				_ = _index68
				var _rhs70 _dafny.Int = _167_v0
				_ = _rhs70
				var _rhs71 _dafny.MultiSet = ((func() _dafny.MultiSet {
					if (_359_v162).Contains(_dafny.IntOfInt64(491)) {
						return (_359_v162).Get(_dafny.IntOfInt64(491)).(_dafny.MultiSet)
					}
					return _357_v160
				})()).Intersection((_357_v160).Update(_356_v159, Companion_Default___.Abs(_167_v0)))
				_ = _rhs71
				var _rhs72 _dafny.Array = _190_v19
				_ = _rhs72
				var _rhs73 _dafny.Int = _167_v0
				_ = _rhs73
				var _lhs54 _dafny.Array = _257_v81
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_257_v81), 0))
				_ = _lhs55
				var _lhs56 _dafny.Array = _355_v158
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(729), _dafny.ArrayLen((_355_v158), 0))
				_ = _lhs57
				var _lhs58 *C8 = _259_v83
				_ = _lhs58
				(_lhs54).ArraySet1(_rhs70, (_lhs55).Int())
				(_lhs56).ArraySet1(_rhs71, (_lhs57).Int())
				_lhs58.F7 = _rhs72
				_167_v0 = _rhs73
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	_dafny.Print(_167_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_168_v1, _dafny.SeqOf(_dafny.IntOfInt64(647), _dafny.IntOfInt64(647), _dafny.IntOfInt64(647), _dafny.IntOfInt64(647), _dafny.IntOfInt64(647), _dafny.IntOfInt64(647))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_169_v2, _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(647), _dafny.IntOfInt64(647)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_170_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_171_v4).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_172_v5, _dafny.SeqOf(true, true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_173_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_173_globalState).F3(), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(647), _dafny.IntOfInt64(647)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_173_globalState.F4).Equals(_dafny.MultiSetOf(true, true, true, true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_174_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_189_v18).Dtor_cf3()).Dtor_cf0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_190_v19).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_191_v20)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_192_v21).Dtor_cf1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_192_v21).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_193_v22)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v23).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _dafny.IntOfInt64(647))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v24).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_196_v25.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v26).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_198_v27).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v27).ArrayGet1((_dafny.Zero).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_198_v27).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v27).ArrayGet1((_dafny.One).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_198_v27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v27).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_198_v27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf1()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_198_v27).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(D0)).Dtor_cf2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_199_v28).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_255_v79).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_256_v80).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_257_v81).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_258_v82).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_259_v83.F7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32(((_259_v83).F8()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_260_v84).Equals(_dafny.MultiSetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65().F7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((((_261_v85).Dtor_cf65()).F8()).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65()).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_261_v85).Dtor_cf65()).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_262_v86).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-1), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.One, _dafny.One, _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v87).Dtor_cf51())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v87).Dtor_cf52())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_263_v87).Dtor_cf53()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-1), _dafny.SeqOf(_dafny.IntOfInt64(598), _dafny.One, _dafny.One, _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_264_v88).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_265_v89).Equals(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_340_v147).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_343_v148).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_345_v150).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_346_v151).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(829), _dafny.One, _dafny.One, _dafny.One, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_347_i8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_354_i9)
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf1 _dafny.Map
	Cf2 _dafny.CodePoint
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf1 _dafny.Map, Cf2 _dafny.CodePoint) D0 {
	return D0{D0_DC2{Cf1, Cf2}}
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
	Cf3 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf3 D0) D0 {
	return D0{D0_DC3{Cf3}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf1() _dafny.Map {
	return _this.Get_().(D0_DC2).Cf1
}

func (_this D0) Dtor_cf2() _dafny.CodePoint {
	return _this.Get_().(D0_DC2).Cf2
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf3() D0 {
	return _this.Get_().(D0_DC3).Cf3
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf3) + ")"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf1.Equals(data2.Cf1) && data1.Cf2 == data2.Cf2
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0 == data2.Cf0
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf3.Equals(data2.Cf3)
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
	Cf5 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf5 bool) D1 {
	return D1{D1_DC5{Cf5}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

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

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC5).Cf5
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf5) + ")"
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
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf5 == data2.Cf5
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0
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
	Cf7 _dafny.Int
	Cf8 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf7 _dafny.Int, Cf8 _dafny.Int) D2 {
	return D2{D2_DC7{Cf7, Cf8}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf9 bool
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf9 bool) D2 {
	return D2{D2_DC8{Cf9}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC6 struct {
	Cf6 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf6 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf6}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC9 struct {
	Cf10 D2
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf10 D2) D2 {
	return D2{D2_DC9{Cf10}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(_dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf6
}

func (_this D2) Dtor_cf10() D2 {
	return _this.Get_().(D2_DC9).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + data.Cf6.VerbatimString(true) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf9 == data2.Cf9
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf6.Equals(data2.Cf6)
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D3_DC10 struct {
	Cf11 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 _dafny.Map) D3 {
	return D3{D3_DC10{Cf11}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D3) Dtor_cf11() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf11.Equals(data2.Cf11)
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
	Cf12 _dafny.Map
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf12 _dafny.Map) D4 {
	return D4{D4_DC11{Cf12}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D4) Dtor_cf12() _dafny.Map {
	return _this.Get_().(D4_DC11).Cf12
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf12) + ")"
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
			return ok && data1.Cf12.Equals(data2.Cf12)
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

type D5_DC12 struct {
	Cf13 _dafny.Array
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf13 _dafny.Array) D5 {
	return D5{D5_DC12{Cf13}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D5) Dtor_cf13() _dafny.Array {
	return _this.Get_().(D5_DC12).Cf13
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf13) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
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

type D6_DC14 struct {
	Cf15 bool
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf15 bool) D6 {
	return D6{D6_DC14{Cf15}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC13 struct {
	Cf14 _dafny.Sequence
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf14 _dafny.Sequence) D6 {
	return D6{D6_DC13{Cf14}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(false)
}

func (_this D6) Dtor_cf15() bool {
	return _this.Get_().(D6_DC14).Cf15
}

func (_this D6) Dtor_cf14() _dafny.Sequence {
	return _this.Get_().(D6_DC13).Cf14
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf14) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
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
	Cf16 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf16 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf16}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

func (CompanionStruct_D7_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D7) Dtor_cf16() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf16
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D8_DC16 struct {
	Cf17 _dafny.Set
}

func (D8_DC16) isD8() {}

func (CompanionStruct_D8_) Create_DC16_(Cf17 _dafny.Set) D8 {
	return D8{D8_DC16{Cf17}}
}

func (_this D8) Is_DC16() bool {
	_, ok := _this.Get_().(D8_DC16)
	return ok
}

func (CompanionStruct_D8_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D8) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D8_DC16).Cf17
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC16:
		{
			return "D8.DC16" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC16:
		{
			data2, ok := other.Get_().(D8_DC16)
			return ok && data1.Cf17.Equals(data2.Cf17)
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

type D9_DC17 struct {
	Cf18 _dafny.Map
}

func (D9_DC17) isD9() {}

func (CompanionStruct_D9_) Create_DC17_(Cf18 _dafny.Map) D9 {
	return D9{D9_DC17{Cf18}}
}

func (_this D9) Is_DC17() bool {
	_, ok := _this.Get_().(D9_DC17)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D9) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D9_DC17).Cf18
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC17:
		{
			return "D9.DC17" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC17:
		{
			data2, ok := other.Get_().(D9_DC17)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D10_DC18 struct {
	Cf19 _dafny.Array
}

func (D10_DC18) isD10() {}

func (CompanionStruct_D10_) Create_DC18_(Cf19 _dafny.Array) D10 {
	return D10{D10_DC18{Cf19}}
}

func (_this D10) Is_DC18() bool {
	_, ok := _this.Get_().(D10_DC18)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D10) Dtor_cf19() _dafny.Array {
	return _this.Get_().(D10_DC18).Cf19
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC18:
		{
			return "D10.DC18" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC18:
		{
			data2, ok := other.Get_().(D10_DC18)
			return ok && data1.Cf19 == data2.Cf19
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

type D11_DC20 struct {
	Cf21 bool
	Cf22 _dafny.Int
}

func (D11_DC20) isD11() {}

func (CompanionStruct_D11_) Create_DC20_(Cf21 bool, Cf22 _dafny.Int) D11 {
	return D11{D11_DC20{Cf21, Cf22}}
}

func (_this D11) Is_DC20() bool {
	_, ok := _this.Get_().(D11_DC20)
	return ok
}

type D11_DC19 struct {
	Cf20 _dafny.CodePoint
}

func (D11_DC19) isD11() {}

func (CompanionStruct_D11_) Create_DC19_(Cf20 _dafny.CodePoint) D11 {
	return D11{D11_DC19{Cf20}}
}

func (_this D11) Is_DC19() bool {
	_, ok := _this.Get_().(D11_DC19)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC20_(false, _dafny.Zero)
}

func (_this D11) Dtor_cf21() bool {
	return _this.Get_().(D11_DC20).Cf21
}

func (_this D11) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D11_DC20).Cf22
}

func (_this D11) Dtor_cf20() _dafny.CodePoint {
	return _this.Get_().(D11_DC19).Cf20
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC20:
		{
			return "D11.DC20" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D11_DC19:
		{
			return "D11.DC19" + "(" + _dafny.String(data.Cf20) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC20:
		{
			data2, ok := other.Get_().(D11_DC20)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22.Cmp(data2.Cf22) == 0
		}
	case D11_DC19:
		{
			data2, ok := other.Get_().(D11_DC19)
			return ok && data1.Cf20 == data2.Cf20
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

type D12_DC22 struct {
	Cf24 _dafny.Map
	Cf25 _dafny.Sequence
	Cf26 _dafny.CodePoint
	Cf27 bool
	Cf28 _dafny.Sequence
}

func (D12_DC22) isD12() {}

func (CompanionStruct_D12_) Create_DC22_(Cf24 _dafny.Map, Cf25 _dafny.Sequence, Cf26 _dafny.CodePoint, Cf27 bool, Cf28 _dafny.Sequence) D12 {
	return D12{D12_DC22{Cf24, Cf25, Cf26, Cf27, Cf28}}
}

func (_this D12) Is_DC22() bool {
	_, ok := _this.Get_().(D12_DC22)
	return ok
}

type D12_DC23 struct {
	Cf29 _dafny.Array
}

func (D12_DC23) isD12() {}

func (CompanionStruct_D12_) Create_DC23_(Cf29 _dafny.Array) D12 {
	return D12{D12_DC23{Cf29}}
}

func (_this D12) Is_DC23() bool {
	_, ok := _this.Get_().(D12_DC23)
	return ok
}

type D12_DC21 struct {
	Cf23 _dafny.Map
}

func (D12_DC21) isD12() {}

func (CompanionStruct_D12_) Create_DC21_(Cf23 _dafny.Map) D12 {
	return D12{D12_DC21{Cf23}}
}

func (_this D12) Is_DC21() bool {
	_, ok := _this.Get_().(D12_DC21)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC22_(_dafny.EmptyMap, _dafny.EmptySeq, _dafny.CodePoint('D'), false, _dafny.EmptySeq)
}

func (_this D12) Dtor_cf24() _dafny.Map {
	return _this.Get_().(D12_DC22).Cf24
}

func (_this D12) Dtor_cf25() _dafny.Sequence {
	return _this.Get_().(D12_DC22).Cf25
}

func (_this D12) Dtor_cf26() _dafny.CodePoint {
	return _this.Get_().(D12_DC22).Cf26
}

func (_this D12) Dtor_cf27() bool {
	return _this.Get_().(D12_DC22).Cf27
}

func (_this D12) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D12_DC22).Cf28
}

func (_this D12) Dtor_cf29() _dafny.Array {
	return _this.Get_().(D12_DC23).Cf29
}

func (_this D12) Dtor_cf23() _dafny.Map {
	return _this.Get_().(D12_DC21).Cf23
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC22:
		{
			return "D12.DC22" + "(" + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + data.Cf28.VerbatimString(true) + ")"
		}
	case D12_DC23:
		{
			return "D12.DC23" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D12_DC21:
		{
			return "D12.DC21" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC22:
		{
			data2, ok := other.Get_().(D12_DC22)
			return ok && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Equals(data2.Cf25) && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28.Equals(data2.Cf28)
		}
	case D12_DC23:
		{
			data2, ok := other.Get_().(D12_DC23)
			return ok && data1.Cf29 == data2.Cf29
		}
	case D12_DC21:
		{
			data2, ok := other.Get_().(D12_DC21)
			return ok && data1.Cf23.Equals(data2.Cf23)
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

type D13_DC24 struct {
	Cf30 _dafny.Map
}

func (D13_DC24) isD13() {}

func (CompanionStruct_D13_) Create_DC24_(Cf30 _dafny.Map) D13 {
	return D13{D13_DC24{Cf30}}
}

func (_this D13) Is_DC24() bool {
	_, ok := _this.Get_().(D13_DC24)
	return ok
}

func (CompanionStruct_D13_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D13) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D13_DC24).Cf30
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC24:
		{
			return "D13.DC24" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC24:
		{
			data2, ok := other.Get_().(D13_DC24)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D14_DC26 struct {
	Cf32 bool
	Cf33 _dafny.Int
	Cf34 _dafny.CodePoint
}

func (D14_DC26) isD14() {}

func (CompanionStruct_D14_) Create_DC26_(Cf32 bool, Cf33 _dafny.Int, Cf34 _dafny.CodePoint) D14 {
	return D14{D14_DC26{Cf32, Cf33, Cf34}}
}

func (_this D14) Is_DC26() bool {
	_, ok := _this.Get_().(D14_DC26)
	return ok
}

type D14_DC25 struct {
	Cf31 *C4
}

func (D14_DC25) isD14() {}

func (CompanionStruct_D14_) Create_DC25_(Cf31 *C4) D14 {
	return D14{D14_DC25{Cf31}}
}

func (_this D14) Is_DC25() bool {
	_, ok := _this.Get_().(D14_DC25)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC26_(false, _dafny.Zero, _dafny.CodePoint('D'))
}

func (_this D14) Dtor_cf32() bool {
	return _this.Get_().(D14_DC26).Cf32
}

func (_this D14) Dtor_cf33() _dafny.Int {
	return _this.Get_().(D14_DC26).Cf33
}

func (_this D14) Dtor_cf34() _dafny.CodePoint {
	return _this.Get_().(D14_DC26).Cf34
}

func (_this D14) Dtor_cf31() *C4 {
	return _this.Get_().(D14_DC25).Cf31
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC26:
		{
			return "D14.DC26" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D14_DC25:
		{
			return "D14.DC25" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC26:
		{
			data2, ok := other.Get_().(D14_DC26)
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33.Cmp(data2.Cf33) == 0 && data1.Cf34 == data2.Cf34
		}
	case D14_DC25:
		{
			data2, ok := other.Get_().(D14_DC25)
			return ok && data1.Cf31 == data2.Cf31
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

type D15_DC28 struct {
	Cf36 bool
	Cf37 bool
}

func (D15_DC28) isD15() {}

func (CompanionStruct_D15_) Create_DC28_(Cf36 bool, Cf37 bool) D15 {
	return D15{D15_DC28{Cf36, Cf37}}
}

func (_this D15) Is_DC28() bool {
	_, ok := _this.Get_().(D15_DC28)
	return ok
}

type D15_DC27 struct {
	Cf35 _dafny.MultiSet
}

func (D15_DC27) isD15() {}

func (CompanionStruct_D15_) Create_DC27_(Cf35 _dafny.MultiSet) D15 {
	return D15{D15_DC27{Cf35}}
}

func (_this D15) Is_DC27() bool {
	_, ok := _this.Get_().(D15_DC27)
	return ok
}

type D15_DC29 struct {
	Cf38 D15
}

func (D15_DC29) isD15() {}

func (CompanionStruct_D15_) Create_DC29_(Cf38 D15) D15 {
	return D15{D15_DC29{Cf38}}
}

func (_this D15) Is_DC29() bool {
	_, ok := _this.Get_().(D15_DC29)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC28_(false, false)
}

func (_this D15) Dtor_cf36() bool {
	return _this.Get_().(D15_DC28).Cf36
}

func (_this D15) Dtor_cf37() bool {
	return _this.Get_().(D15_DC28).Cf37
}

func (_this D15) Dtor_cf35() _dafny.MultiSet {
	return _this.Get_().(D15_DC27).Cf35
}

func (_this D15) Dtor_cf38() D15 {
	return _this.Get_().(D15_DC29).Cf38
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC28:
		{
			return "D15.DC28" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ")"
		}
	case D15_DC27:
		{
			return "D15.DC27" + "(" + _dafny.String(data.Cf35) + ")"
		}
	case D15_DC29:
		{
			return "D15.DC29" + "(" + _dafny.String(data.Cf38) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC28:
		{
			data2, ok := other.Get_().(D15_DC28)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37
		}
	case D15_DC27:
		{
			data2, ok := other.Get_().(D15_DC27)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	case D15_DC29:
		{
			data2, ok := other.Get_().(D15_DC29)
			return ok && data1.Cf38.Equals(data2.Cf38)
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

type D16_DC31 struct {
	Cf40 bool
}

func (D16_DC31) isD16() {}

func (CompanionStruct_D16_) Create_DC31_(Cf40 bool) D16 {
	return D16{D16_DC31{Cf40}}
}

func (_this D16) Is_DC31() bool {
	_, ok := _this.Get_().(D16_DC31)
	return ok
}

type D16_DC32 struct {
	Cf41 bool
}

func (D16_DC32) isD16() {}

func (CompanionStruct_D16_) Create_DC32_(Cf41 bool) D16 {
	return D16{D16_DC32{Cf41}}
}

func (_this D16) Is_DC32() bool {
	_, ok := _this.Get_().(D16_DC32)
	return ok
}

type D16_DC30 struct {
	Cf39 _dafny.Sequence
}

func (D16_DC30) isD16() {}

func (CompanionStruct_D16_) Create_DC30_(Cf39 _dafny.Sequence) D16 {
	return D16{D16_DC30{Cf39}}
}

func (_this D16) Is_DC30() bool {
	_, ok := _this.Get_().(D16_DC30)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC31_(false)
}

func (_this D16) Dtor_cf40() bool {
	return _this.Get_().(D16_DC31).Cf40
}

func (_this D16) Dtor_cf41() bool {
	return _this.Get_().(D16_DC32).Cf41
}

func (_this D16) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D16_DC30).Cf39
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC31:
		{
			return "D16.DC31" + "(" + _dafny.String(data.Cf40) + ")"
		}
	case D16_DC32:
		{
			return "D16.DC32" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D16_DC30:
		{
			return "D16.DC30" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC31:
		{
			data2, ok := other.Get_().(D16_DC31)
			return ok && data1.Cf40 == data2.Cf40
		}
	case D16_DC32:
		{
			data2, ok := other.Get_().(D16_DC32)
			return ok && data1.Cf41 == data2.Cf41
		}
	case D16_DC30:
		{
			data2, ok := other.Get_().(D16_DC30)
			return ok && data1.Cf39.Equals(data2.Cf39)
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

type D17_DC34 struct {
	Cf43 _dafny.CodePoint
}

func (D17_DC34) isD17() {}

func (CompanionStruct_D17_) Create_DC34_(Cf43 _dafny.CodePoint) D17 {
	return D17{D17_DC34{Cf43}}
}

func (_this D17) Is_DC34() bool {
	_, ok := _this.Get_().(D17_DC34)
	return ok
}

type D17_DC33 struct {
	Cf42 _dafny.Map
}

func (D17_DC33) isD17() {}

func (CompanionStruct_D17_) Create_DC33_(Cf42 _dafny.Map) D17 {
	return D17{D17_DC33{Cf42}}
}

func (_this D17) Is_DC33() bool {
	_, ok := _this.Get_().(D17_DC33)
	return ok
}

type D17_DC35 struct {
	Cf44 D17
}

func (D17_DC35) isD17() {}

func (CompanionStruct_D17_) Create_DC35_(Cf44 D17) D17 {
	return D17{D17_DC35{Cf44}}
}

func (_this D17) Is_DC35() bool {
	_, ok := _this.Get_().(D17_DC35)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC34_(_dafny.CodePoint('D'))
}

func (_this D17) Dtor_cf43() _dafny.CodePoint {
	return _this.Get_().(D17_DC34).Cf43
}

func (_this D17) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D17_DC33).Cf42
}

func (_this D17) Dtor_cf44() D17 {
	return _this.Get_().(D17_DC35).Cf44
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC34:
		{
			return "D17.DC34" + "(" + _dafny.String(data.Cf43) + ")"
		}
	case D17_DC33:
		{
			return "D17.DC33" + "(" + _dafny.String(data.Cf42) + ")"
		}
	case D17_DC35:
		{
			return "D17.DC35" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC34:
		{
			data2, ok := other.Get_().(D17_DC34)
			return ok && data1.Cf43 == data2.Cf43
		}
	case D17_DC33:
		{
			data2, ok := other.Get_().(D17_DC33)
			return ok && data1.Cf42.Equals(data2.Cf42)
		}
	case D17_DC35:
		{
			data2, ok := other.Get_().(D17_DC35)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D18_DC37 struct {
	Cf46 _dafny.Int
	Cf47 _dafny.Int
	Cf48 _dafny.Int
}

func (D18_DC37) isD18() {}

func (CompanionStruct_D18_) Create_DC37_(Cf46 _dafny.Int, Cf47 _dafny.Int, Cf48 _dafny.Int) D18 {
	return D18{D18_DC37{Cf46, Cf47, Cf48}}
}

func (_this D18) Is_DC37() bool {
	_, ok := _this.Get_().(D18_DC37)
	return ok
}

type D18_DC36 struct {
	Cf45 _dafny.Array
}

func (D18_DC36) isD18() {}

func (CompanionStruct_D18_) Create_DC36_(Cf45 _dafny.Array) D18 {
	return D18{D18_DC36{Cf45}}
}

func (_this D18) Is_DC36() bool {
	_, ok := _this.Get_().(D18_DC36)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC37_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D18) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D18_DC37).Cf46
}

func (_this D18) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D18_DC37).Cf47
}

func (_this D18) Dtor_cf48() _dafny.Int {
	return _this.Get_().(D18_DC37).Cf48
}

func (_this D18) Dtor_cf45() _dafny.Array {
	return _this.Get_().(D18_DC36).Cf45
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC37:
		{
			return "D18.DC37" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D18_DC36:
		{
			return "D18.DC36" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC37:
		{
			data2, ok := other.Get_().(D18_DC37)
			return ok && data1.Cf46.Cmp(data2.Cf46) == 0 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Cmp(data2.Cf48) == 0
		}
	case D18_DC36:
		{
			data2, ok := other.Get_().(D18_DC36)
			return ok && data1.Cf45 == data2.Cf45
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

type D19_DC39 struct {
	Cf50 _dafny.Sequence
}

func (D19_DC39) isD19() {}

func (CompanionStruct_D19_) Create_DC39_(Cf50 _dafny.Sequence) D19 {
	return D19{D19_DC39{Cf50}}
}

func (_this D19) Is_DC39() bool {
	_, ok := _this.Get_().(D19_DC39)
	return ok
}

type D19_DC40 struct {
	Cf51 bool
	Cf52 _dafny.Int
	Cf53 _dafny.Map
}

func (D19_DC40) isD19() {}

func (CompanionStruct_D19_) Create_DC40_(Cf51 bool, Cf52 _dafny.Int, Cf53 _dafny.Map) D19 {
	return D19{D19_DC40{Cf51, Cf52, Cf53}}
}

func (_this D19) Is_DC40() bool {
	_, ok := _this.Get_().(D19_DC40)
	return ok
}

type D19_DC38 struct {
	Cf49 _dafny.Set
}

func (D19_DC38) isD19() {}

func (CompanionStruct_D19_) Create_DC38_(Cf49 _dafny.Set) D19 {
	return D19{D19_DC38{Cf49}}
}

func (_this D19) Is_DC38() bool {
	_, ok := _this.Get_().(D19_DC38)
	return ok
}

type D19_DC41 struct {
	Cf54 D19
}

func (D19_DC41) isD19() {}

func (CompanionStruct_D19_) Create_DC41_(Cf54 D19) D19 {
	return D19{D19_DC41{Cf54}}
}

func (_this D19) Is_DC41() bool {
	_, ok := _this.Get_().(D19_DC41)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC39_(_dafny.EmptySeq)
}

func (_this D19) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D19_DC39).Cf50
}

func (_this D19) Dtor_cf51() bool {
	return _this.Get_().(D19_DC40).Cf51
}

func (_this D19) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D19_DC40).Cf52
}

func (_this D19) Dtor_cf53() _dafny.Map {
	return _this.Get_().(D19_DC40).Cf53
}

func (_this D19) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D19_DC38).Cf49
}

func (_this D19) Dtor_cf54() D19 {
	return _this.Get_().(D19_DC41).Cf54
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC39:
		{
			return "D19.DC39" + "(" + _dafny.String(data.Cf50) + ")"
		}
	case D19_DC40:
		{
			return "D19.DC40" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D19_DC38:
		{
			return "D19.DC38" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D19_DC41:
		{
			return "D19.DC41" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC39:
		{
			data2, ok := other.Get_().(D19_DC39)
			return ok && data1.Cf50.Equals(data2.Cf50)
		}
	case D19_DC40:
		{
			data2, ok := other.Get_().(D19_DC40)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53.Equals(data2.Cf53)
		}
	case D19_DC38:
		{
			data2, ok := other.Get_().(D19_DC38)
			return ok && data1.Cf49.Equals(data2.Cf49)
		}
	case D19_DC41:
		{
			data2, ok := other.Get_().(D19_DC41)
			return ok && data1.Cf54.Equals(data2.Cf54)
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

type D20_DC43 struct {
	Cf56 _dafny.Int
	Cf57 bool
	Cf58 _dafny.Int
}

func (D20_DC43) isD20() {}

func (CompanionStruct_D20_) Create_DC43_(Cf56 _dafny.Int, Cf57 bool, Cf58 _dafny.Int) D20 {
	return D20{D20_DC43{Cf56, Cf57, Cf58}}
}

func (_this D20) Is_DC43() bool {
	_, ok := _this.Get_().(D20_DC43)
	return ok
}

type D20_DC42 struct {
	Cf55 *C6
}

func (D20_DC42) isD20() {}

func (CompanionStruct_D20_) Create_DC42_(Cf55 *C6) D20 {
	return D20{D20_DC42{Cf55}}
}

func (_this D20) Is_DC42() bool {
	_, ok := _this.Get_().(D20_DC42)
	return ok
}

type D20_DC44 struct {
	Cf59 D20
}

func (D20_DC44) isD20() {}

func (CompanionStruct_D20_) Create_DC44_(Cf59 D20) D20 {
	return D20{D20_DC44{Cf59}}
}

func (_this D20) Is_DC44() bool {
	_, ok := _this.Get_().(D20_DC44)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC43_(_dafny.Zero, false, _dafny.Zero)
}

func (_this D20) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D20_DC43).Cf56
}

func (_this D20) Dtor_cf57() bool {
	return _this.Get_().(D20_DC43).Cf57
}

func (_this D20) Dtor_cf58() _dafny.Int {
	return _this.Get_().(D20_DC43).Cf58
}

func (_this D20) Dtor_cf55() *C6 {
	return _this.Get_().(D20_DC42).Cf55
}

func (_this D20) Dtor_cf59() D20 {
	return _this.Get_().(D20_DC44).Cf59
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC43:
		{
			return "D20.DC43" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ")"
		}
	case D20_DC42:
		{
			return "D20.DC42" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D20_DC44:
		{
			return "D20.DC44" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC43:
		{
			data2, ok := other.Get_().(D20_DC43)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57 == data2.Cf57 && data1.Cf58.Cmp(data2.Cf58) == 0
		}
	case D20_DC42:
		{
			data2, ok := other.Get_().(D20_DC42)
			return ok && data1.Cf55 == data2.Cf55
		}
	case D20_DC44:
		{
			data2, ok := other.Get_().(D20_DC44)
			return ok && data1.Cf59.Equals(data2.Cf59)
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

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC46 struct {
	Cf61 bool
	Cf62 bool
	Cf63 _dafny.Int
}

func (D21_DC46) isD21() {}

func (CompanionStruct_D21_) Create_DC46_(Cf61 bool, Cf62 bool, Cf63 _dafny.Int) D21 {
	return D21{D21_DC46{Cf61, Cf62, Cf63}}
}

func (_this D21) Is_DC46() bool {
	_, ok := _this.Get_().(D21_DC46)
	return ok
}

type D21_DC47 struct {
	Cf64 bool
}

func (D21_DC47) isD21() {}

func (CompanionStruct_D21_) Create_DC47_(Cf64 bool) D21 {
	return D21{D21_DC47{Cf64}}
}

func (_this D21) Is_DC47() bool {
	_, ok := _this.Get_().(D21_DC47)
	return ok
}

type D21_DC45 struct {
	Cf60 _dafny.MultiSet
}

func (D21_DC45) isD21() {}

func (CompanionStruct_D21_) Create_DC45_(Cf60 _dafny.MultiSet) D21 {
	return D21{D21_DC45{Cf60}}
}

func (_this D21) Is_DC45() bool {
	_, ok := _this.Get_().(D21_DC45)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC46_(false, false, _dafny.Zero)
}

func (_this D21) Dtor_cf61() bool {
	return _this.Get_().(D21_DC46).Cf61
}

func (_this D21) Dtor_cf62() bool {
	return _this.Get_().(D21_DC46).Cf62
}

func (_this D21) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D21_DC46).Cf63
}

func (_this D21) Dtor_cf64() bool {
	return _this.Get_().(D21_DC47).Cf64
}

func (_this D21) Dtor_cf60() _dafny.MultiSet {
	return _this.Get_().(D21_DC45).Cf60
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC46:
		{
			return "D21.DC46" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D21_DC47:
		{
			return "D21.DC47" + "(" + _dafny.String(data.Cf64) + ")"
		}
	case D21_DC45:
		{
			return "D21.DC45" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC46:
		{
			data2, ok := other.Get_().(D21_DC46)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62 && data1.Cf63.Cmp(data2.Cf63) == 0
		}
	case D21_DC47:
		{
			data2, ok := other.Get_().(D21_DC47)
			return ok && data1.Cf64 == data2.Cf64
		}
	case D21_DC45:
		{
			data2, ok := other.Get_().(D21_DC45)
			return ok && data1.Cf60.Equals(data2.Cf60)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC49 struct {
	Cf66 _dafny.Set
}

func (D22_DC49) isD22() {}

func (CompanionStruct_D22_) Create_DC49_(Cf66 _dafny.Set) D22 {
	return D22{D22_DC49{Cf66}}
}

func (_this D22) Is_DC49() bool {
	_, ok := _this.Get_().(D22_DC49)
	return ok
}

type D22_DC48 struct {
	Cf65 *C8
}

func (D22_DC48) isD22() {}

func (CompanionStruct_D22_) Create_DC48_(Cf65 *C8) D22 {
	return D22{D22_DC48{Cf65}}
}

func (_this D22) Is_DC48() bool {
	_, ok := _this.Get_().(D22_DC48)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC49_(_dafny.EmptySet)
}

func (_this D22) Dtor_cf66() _dafny.Set {
	return _this.Get_().(D22_DC49).Cf66
}

func (_this D22) Dtor_cf65() *C8 {
	return _this.Get_().(D22_DC48).Cf65
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC49:
		{
			return "D22.DC49" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D22_DC48:
		{
			return "D22.DC48" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC49:
		{
			data2, ok := other.Get_().(D22_DC49)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	case D22_DC48:
		{
			data2, ok := other.Get_().(D22_DC48)
			return ok && data1.Cf65 == data2.Cf65
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of datatype D23
type D23 struct {
	Data_D23_
}

func (_this D23) Get_() Data_D23_ {
	return _this.Data_D23_
}

type Data_D23_ interface {
	isD23()
}

type CompanionStruct_D23_ struct {
}

var Companion_D23_ = CompanionStruct_D23_{}

type D23_DC50 struct {
	Cf67 _dafny.Set
}

func (D23_DC50) isD23() {}

func (CompanionStruct_D23_) Create_DC50_(Cf67 _dafny.Set) D23 {
	return D23{D23_DC50{Cf67}}
}

func (_this D23) Is_DC50() bool {
	_, ok := _this.Get_().(D23_DC50)
	return ok
}

func (CompanionStruct_D23_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D23) Dtor_cf67() _dafny.Set {
	return _this.Get_().(D23_DC50).Cf67
}

func (_this D23) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D23_DC50:
		{
			return "D23.DC50" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D23) Equals(other D23) bool {
	switch data1 := _this.Get_().(type) {
	case D23_DC50:
		{
			data2, ok := other.Get_().(D23_DC50)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D23) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D23)
	return ok && _this.Equals(typed)
}

func Type_D23_() _dafny.TypeDescriptor {
	return type_D23_{}
}

type type_D23_ struct {
}

func (_this type_D23_) Default() interface{} {
	return Companion_D23_.Default()
}

func (_this type_D23_) String() string {
	return "main.D23"
}
func (_this D23) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D23{}

// End of datatype D23

// Definition of datatype D24
type D24 struct {
	Data_D24_
}

func (_this D24) Get_() Data_D24_ {
	return _this.Data_D24_
}

type Data_D24_ interface {
	isD24()
}

type CompanionStruct_D24_ struct {
}

var Companion_D24_ = CompanionStruct_D24_{}

type D24_DC51 struct {
	Cf68 _dafny.Map
}

func (D24_DC51) isD24() {}

func (CompanionStruct_D24_) Create_DC51_(Cf68 _dafny.Map) D24 {
	return D24{D24_DC51{Cf68}}
}

func (_this D24) Is_DC51() bool {
	_, ok := _this.Get_().(D24_DC51)
	return ok
}

func (CompanionStruct_D24_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D24) Dtor_cf68() _dafny.Map {
	return _this.Get_().(D24_DC51).Cf68
}

func (_this D24) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D24_DC51:
		{
			return "D24.DC51" + "(" + _dafny.String(data.Cf68) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D24) Equals(other D24) bool {
	switch data1 := _this.Get_().(type) {
	case D24_DC51:
		{
			data2, ok := other.Get_().(D24_DC51)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D24) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D24)
	return ok && _this.Equals(typed)
}

func Type_D24_() _dafny.TypeDescriptor {
	return type_D24_{}
}

type type_D24_ struct {
}

func (_this type_D24_) Default() interface{} {
	return Companion_D24_.Default()
}

func (_this type_D24_) String() string {
	return "main.D24"
}
func (_this D24) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D24{}

// End of datatype D24

// Definition of datatype D25
type D25 struct {
	Data_D25_
}

func (_this D25) Get_() Data_D25_ {
	return _this.Data_D25_
}

type Data_D25_ interface {
	isD25()
}

type CompanionStruct_D25_ struct {
}

var Companion_D25_ = CompanionStruct_D25_{}

type D25_DC53 struct {
	Cf70 bool
}

func (D25_DC53) isD25() {}

func (CompanionStruct_D25_) Create_DC53_(Cf70 bool) D25 {
	return D25{D25_DC53{Cf70}}
}

func (_this D25) Is_DC53() bool {
	_, ok := _this.Get_().(D25_DC53)
	return ok
}

type D25_DC54 struct {
	Cf71 bool
	Cf72 bool
	Cf73 bool
}

func (D25_DC54) isD25() {}

func (CompanionStruct_D25_) Create_DC54_(Cf71 bool, Cf72 bool, Cf73 bool) D25 {
	return D25{D25_DC54{Cf71, Cf72, Cf73}}
}

func (_this D25) Is_DC54() bool {
	_, ok := _this.Get_().(D25_DC54)
	return ok
}

type D25_DC52 struct {
	Cf69 _dafny.MultiSet
}

func (D25_DC52) isD25() {}

func (CompanionStruct_D25_) Create_DC52_(Cf69 _dafny.MultiSet) D25 {
	return D25{D25_DC52{Cf69}}
}

func (_this D25) Is_DC52() bool {
	_, ok := _this.Get_().(D25_DC52)
	return ok
}

func (CompanionStruct_D25_) Default() D25 {
	return Companion_D25_.Create_DC53_(false)
}

func (_this D25) Dtor_cf70() bool {
	return _this.Get_().(D25_DC53).Cf70
}

func (_this D25) Dtor_cf71() bool {
	return _this.Get_().(D25_DC54).Cf71
}

func (_this D25) Dtor_cf72() bool {
	return _this.Get_().(D25_DC54).Cf72
}

func (_this D25) Dtor_cf73() bool {
	return _this.Get_().(D25_DC54).Cf73
}

func (_this D25) Dtor_cf69() _dafny.MultiSet {
	return _this.Get_().(D25_DC52).Cf69
}

func (_this D25) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D25_DC53:
		{
			return "D25.DC53" + "(" + _dafny.String(data.Cf70) + ")"
		}
	case D25_DC54:
		{
			return "D25.DC54" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ")"
		}
	case D25_DC52:
		{
			return "D25.DC52" + "(" + _dafny.String(data.Cf69) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D25) Equals(other D25) bool {
	switch data1 := _this.Get_().(type) {
	case D25_DC53:
		{
			data2, ok := other.Get_().(D25_DC53)
			return ok && data1.Cf70 == data2.Cf70
		}
	case D25_DC54:
		{
			data2, ok := other.Get_().(D25_DC54)
			return ok && data1.Cf71 == data2.Cf71 && data1.Cf72 == data2.Cf72 && data1.Cf73 == data2.Cf73
		}
	case D25_DC52:
		{
			data2, ok := other.Get_().(D25_DC52)
			return ok && data1.Cf69.Equals(data2.Cf69)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D25) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D25)
	return ok && _this.Equals(typed)
}

func Type_D25_() _dafny.TypeDescriptor {
	return type_D25_{}
}

type type_D25_ struct {
}

func (_this type_D25_) Default() interface{} {
	return Companion_D25_.Default()
}

func (_this type_D25_) String() string {
	return "main.D25"
}
func (_this D25) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D25{}

// End of datatype D25

// Definition of datatype D26
type D26 struct {
	Data_D26_
}

func (_this D26) Get_() Data_D26_ {
	return _this.Data_D26_
}

type Data_D26_ interface {
	isD26()
}

type CompanionStruct_D26_ struct {
}

var Companion_D26_ = CompanionStruct_D26_{}

type D26_DC56 struct {
	Cf75 _dafny.Sequence
	Cf76 bool
}

func (D26_DC56) isD26() {}

func (CompanionStruct_D26_) Create_DC56_(Cf75 _dafny.Sequence, Cf76 bool) D26 {
	return D26{D26_DC56{Cf75, Cf76}}
}

func (_this D26) Is_DC56() bool {
	_, ok := _this.Get_().(D26_DC56)
	return ok
}

type D26_DC57 struct {
	Cf77 bool
}

func (D26_DC57) isD26() {}

func (CompanionStruct_D26_) Create_DC57_(Cf77 bool) D26 {
	return D26{D26_DC57{Cf77}}
}

func (_this D26) Is_DC57() bool {
	_, ok := _this.Get_().(D26_DC57)
	return ok
}

type D26_DC58 struct {
	Cf78 _dafny.Sequence
	Cf79 _dafny.Int
}

func (D26_DC58) isD26() {}

func (CompanionStruct_D26_) Create_DC58_(Cf78 _dafny.Sequence, Cf79 _dafny.Int) D26 {
	return D26{D26_DC58{Cf78, Cf79}}
}

func (_this D26) Is_DC58() bool {
	_, ok := _this.Get_().(D26_DC58)
	return ok
}

type D26_DC55 struct {
	Cf74 _dafny.Array
}

func (D26_DC55) isD26() {}

func (CompanionStruct_D26_) Create_DC55_(Cf74 _dafny.Array) D26 {
	return D26{D26_DC55{Cf74}}
}

func (_this D26) Is_DC55() bool {
	_, ok := _this.Get_().(D26_DC55)
	return ok
}

type D26_DC59 struct {
	Cf80 D26
}

func (D26_DC59) isD26() {}

func (CompanionStruct_D26_) Create_DC59_(Cf80 D26) D26 {
	return D26{D26_DC59{Cf80}}
}

func (_this D26) Is_DC59() bool {
	_, ok := _this.Get_().(D26_DC59)
	return ok
}

func (CompanionStruct_D26_) Default() D26 {
	return Companion_D26_.Create_DC56_(_dafny.EmptySeq, false)
}

func (_this D26) Dtor_cf75() _dafny.Sequence {
	return _this.Get_().(D26_DC56).Cf75
}

func (_this D26) Dtor_cf76() bool {
	return _this.Get_().(D26_DC56).Cf76
}

func (_this D26) Dtor_cf77() bool {
	return _this.Get_().(D26_DC57).Cf77
}

func (_this D26) Dtor_cf78() _dafny.Sequence {
	return _this.Get_().(D26_DC58).Cf78
}

func (_this D26) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D26_DC58).Cf79
}

func (_this D26) Dtor_cf74() _dafny.Array {
	return _this.Get_().(D26_DC55).Cf74
}

func (_this D26) Dtor_cf80() D26 {
	return _this.Get_().(D26_DC59).Cf80
}

func (_this D26) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D26_DC56:
		{
			return "D26.DC56" + "(" + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D26_DC57:
		{
			return "D26.DC57" + "(" + _dafny.String(data.Cf77) + ")"
		}
	case D26_DC58:
		{
			return "D26.DC58" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D26_DC55:
		{
			return "D26.DC55" + "(" + _dafny.String(data.Cf74) + ")"
		}
	case D26_DC59:
		{
			return "D26.DC59" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D26) Equals(other D26) bool {
	switch data1 := _this.Get_().(type) {
	case D26_DC56:
		{
			data2, ok := other.Get_().(D26_DC56)
			return ok && data1.Cf75.Equals(data2.Cf75) && data1.Cf76 == data2.Cf76
		}
	case D26_DC57:
		{
			data2, ok := other.Get_().(D26_DC57)
			return ok && data1.Cf77 == data2.Cf77
		}
	case D26_DC58:
		{
			data2, ok := other.Get_().(D26_DC58)
			return ok && data1.Cf78.Equals(data2.Cf78) && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D26_DC55:
		{
			data2, ok := other.Get_().(D26_DC55)
			return ok && data1.Cf74 == data2.Cf74
		}
	case D26_DC59:
		{
			data2, ok := other.Get_().(D26_DC59)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D26) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D26)
	return ok && _this.Equals(typed)
}

func Type_D26_() _dafny.TypeDescriptor {
	return type_D26_{}
}

type type_D26_ struct {
}

func (_this type_D26_) Default() interface{} {
	return Companion_D26_.Default()
}

func (_this type_D26_) String() string {
	return "main.D26"
}
func (_this D26) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D26{}

// End of datatype D26

// Definition of trait T0
type T0 interface {
	String() string
	Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int
	Fm5(p0 D2, globalState *GlobalState) D1
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
	Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int
	Fm5(p0 D2, globalState *GlobalState) D1
	M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState)
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
	Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int
	Fm5(p0 D2, globalState *GlobalState) D1
	M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState)
	Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool
	Fm7(p0 bool, globalState *GlobalState) _dafny.Int
	M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence)
	M3(globalState *GlobalState) _dafny.Int
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

// Definition of trait T3
type T3 interface {
	String() string
	Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int
	Fm5(p0 D2, globalState *GlobalState) D1
	Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool
	Fm7(p0 bool, globalState *GlobalState) _dafny.Int
	M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState)
	M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence)
	M3(globalState *GlobalState) _dafny.Int
	Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int
	F5() _dafny.CodePoint
	F6() bool
}
type CompanionStruct_T3_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T3_ = CompanionStruct_T3_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T3_) CastTo_(x interface{}) T3 {
	var t T3
	t, _ = x.(T3)
	return t
}

// End of trait T3

// Definition of trait T4
type T4 interface {
	String() string
	Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int
	Fm5(p0 D2, globalState *GlobalState) D1
	Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool
	Fm7(p0 bool, globalState *GlobalState) _dafny.Int
	Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int
	M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState)
	M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence)
	M3(globalState *GlobalState) _dafny.Int
	F5() _dafny.CodePoint
	F6() bool
	Fm9(p0 _dafny.Int, p1 D0, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool
	Fm10(p0 _dafny.MultiSet, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool
	M4(p0 _dafny.Array, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) (D2, _dafny.MultiSet)
	M5(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState)
}
type CompanionStruct_T4_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T4_ = CompanionStruct_T4_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T4_) CastTo_(x interface{}) T4 {
	var t T4
	t, _ = x.(T4)
	return t
}

// End of trait T4

// Definition of class GlobalState
type GlobalState struct {
	F1  _dafny.Int
	F4  _dafny.MultiSet
	_f0 _dafny.CodePoint
	_f2 _dafny.Int
	_f3 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this.F4 = _dafny.EmptyMultiSet
	_this._f0 = _dafny.CodePoint('D')
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.CodePoint, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Sequence, f4 _dafny.MultiSet) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
	}
}
func (_this *GlobalState) F0() _dafny.CodePoint {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() _dafny.Int {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Sequence {
	{
		return _this._f3
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F12  _dafny.Sequence
	_f11 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F12 = _dafny.EmptySeq
	_this._f11 = false
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

func (_this *C0) Ctor__(f11 bool, f12 _dafny.Sequence) {
	{
		(_this)._f11 = f11
		(_this).F12 = f12
	}
}
func (_this *C0) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-765), (_this).F11())).Cardinality()
	}
}
func (_this *C0) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(31), _dafny.IntOfInt64(-303)))
	}
}
func (_this *C0) F11() bool {
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C1{}
var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		if !(!(!(false))) {
			return (_dafny.MultiSetOf(_dafny.IntOfInt64(-16), _dafny.IntOfInt64(436), _dafny.IntOfInt64(81), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(824))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_360_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(645)
			}))).Cardinality())), _dafny.IntOfInt64(-807))).IsSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(484), _dafny.IntOfInt64(447)))
		} else {
			return (true) == (!(!(false)))
		}
	}
}
func (_this *C1) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-399), (func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC5_(true))).Elements()); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _361_v0 D1
				_361_v0 = interface{}(_compr_21).(D1)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC5_(true)), _361_v0) {
					_coll21.Add(_361_v0)
				}
			}
			return _coll21.ToSet()
		}()).Cardinality())
	}
}
func (_this *C1) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.SeqOf(!((_dafny.SetOf((func() _dafny.Set {
			var _coll22 = _dafny.NewBuilder()
			_ = _coll22
			for _iter22 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(907), true)).Keys().Elements()); ; {
				_compr_22, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _362_v0 _dafny.Int
				_362_v0 = interface{}(_compr_22).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(907), true)).Contains(_362_v0) {
					_coll22.Add((_362_v0).Minus(_dafny.IntOfInt64(85)))
				}
			}
			return _coll22.ToSet()
		}()).Cardinality())).IsSubsetOf(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(46), !(true))).Cardinality(), _dafny.IntOfInt64(873)))))).Cardinality())
	}
}
func (_this *C1) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_(_dafny.IntOfInt64(-49))
	}
}
func (_this *C1) Fm14(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aam"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(143))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg23 _dafny.Int) interface{} {
				return coer23(arg23)
			}
		}(func(_363_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})), _dafny.UnicodeSeqOfUtf8Bytes("hteems")))).Cardinality())
	}
}
func (_this *C1) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		r0 = (_this).Fm6(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(42))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_364_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-615)
		})), globalState)
		var _365_v0 _dafny.Sequence
		_ = _365_v0
		_365_v0 = _dafny.SeqOf(p0, p0, p0)
		_365_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _365_v0)
		var _366_v1 _dafny.MultiSet
		_ = _366_v1
		_366_v1 = _dafny.MultiSetOf((_this).Fm7(p0, globalState))
		var _rhs74 bool = p0
		_ = _rhs74
		var _rhs75 _dafny.Int = (_this).Fm4(_dafny.SeqOf(p0), _dafny.IntOfInt64(325), Companion_Default___.SafeModuloInt(p1, p1), _366_v1, globalState)
		_ = _rhs75
		var _lhs59 *GlobalState = globalState
		_ = _lhs59
		r0 = _rhs74
		_lhs59.F1 = _rhs75
		var _367_v2 _dafny.Array
		_ = _367_v2
		var _nwElement0_10 bool = Companion_Default___.Fm0(p0, globalState)
		_ = _nwElement0_10
		var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(9))
		_ = _nw51
		(_nw51).ArraySet1(_nwElement0_10, 0)
		(_nw51).ArraySet1(!(p0), 1)
		(_nw51).ArraySet1(!(!((_this).Fm6(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(578))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_368_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_369_i1 _dafny.Int) _dafny.Int {
				return _368_p1
			}
		})(p1))), globalState))), 2)
		(_nw51).ArraySet1(p0, 3)
		(_nw51).ArraySet1(p0, 4)
		(_nw51).ArraySet1(p0, 5)
		(_nw51).ArraySet1(p0, 6)
		(_nw51).ArraySet1(p0, 7)
		(_nw51).ArraySet1(false, 8)
		_367_v2 = _nw51
		var _370_v3 _dafny.MultiSet
		_ = _370_v3
		_370_v3 = _dafny.MultiSetOf(_367_v2)
		var _371_v4 D0
		_ = _371_v4
		_371_v4 = Companion_D0_.Create_DC0_(p0)
		var _pat_let_tv0 = p2
		_ = _pat_let_tv0
		var _pat_let_tv1 = p2
		_ = _pat_let_tv1
		var _pat_let_tv2 = p2
		_ = _pat_let_tv2
		var _pat_let_tv3 = p2
		_ = _pat_let_tv3
		var _pat_let_tv4 = p2
		_ = _pat_let_tv4
		var _pat_let_tv5 = p2
		_ = _pat_let_tv5
		var _rhs76 _dafny.MultiSet = _dafny.MultiSetOf(_367_v2, _367_v2, _367_v2)
		_ = _rhs76
		var _rhs77 _dafny.Int = _dafny.IntOfUint32((func(_source9 D0) _dafny.Sequence {
			if _source9.Is_DC1() {
				return _pat_let_tv0
			} else if _source9.Is_DC2() {
				var _372___mcc_h0 _dafny.Map = _source9.Get_().(D0_DC2).Cf1
				_ = _372___mcc_h0
				var _373___mcc_h1 _dafny.CodePoint = _source9.Get_().(D0_DC2).Cf2
				_ = _373___mcc_h1
				var _374_cf2 _dafny.CodePoint = _373___mcc_h1
				_ = _374_cf2
				var _375_cf1 _dafny.Map = _372___mcc_h0
				_ = _375_cf1
				return _pat_let_tv1
			} else if _source9.Is_DC0() {
				var _376___mcc_h2 bool = _source9.Get_().(D0_DC0).Cf0
				_ = _376___mcc_h2
				var _377_cf0 bool = _376___mcc_h2
				_ = _377_cf0
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv2, _pat_let_tv3)
			} else {
				var _378___mcc_h3 D0 = _source9.Get_().(D0_DC3).Cf3
				_ = _378___mcc_h3
				var _379_cf3 D0 = _378___mcc_h3
				_ = _379_cf3
				return _dafny.Companion_Sequence_.Concatenate(_pat_let_tv4, _pat_let_tv5)
			}
		}(_371_v4)).Cardinality())
		_ = _rhs77
		var _lhs60 *GlobalState = globalState
		_ = _lhs60
		_370_v3 = _rhs76
		_lhs60.F1 = _rhs77
		var _380_v5 _dafny.Array
		_ = _380_v5
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(23))
		_ = _nw52
		_380_v5 = _nw52
		for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_380_v5), 0))); ; {
			_guard_loop_0, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _381_i2 _dafny.Int
			_381_i2 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_381_i2).Sign() != -1) && ((_381_i2).Cmp(_dafny.ArrayLen((_380_v5), 0)) < 0)) {
				(_380_v5).ArraySet1((func() _dafny.Map {
					if p0 {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(179), p1)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.SetOf(p0, p0)).Cardinality())
				})(), (_381_i2).Int())
			}
		}
		var _382_v6 _dafny.Array
		_ = _382_v6
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw53
		_382_v6 = _nw53
		var _383_v7 D2
		_ = _383_v7
		_383_v7 = Companion_D2_.Create_DC7_(p1, (_dafny.Zero).Minus((_this).Fm14(globalState)))
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.ArrayLen((_382_v6), 0))
		_ = _index69
		(_382_v6).ArraySet1((p1).Minus((_383_v7).Dtor_cf7()), (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(308), _dafny.ArrayLen((_382_v6), 0))
		_ = _index70
		(_382_v6).ArraySet1((_370_v3).Cardinality(), (_index70).Int())
		r0 = p0
		r1 = p2
		return r0, r1
	}
}
func (_this *C1) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _384_v0 bool
		_ = _384_v0
		_384_v0 = true
		var _385_v1 _dafny.Sequence
		_ = _385_v1
		_385_v1 = _dafny.UnicodeSeqOfUtf8Bytes("mof")
		var _386_v2 _dafny.Map
		_ = _386_v2
		_386_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((!(_384_v0)) == (_384_v0), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _385_v1))
		_386_v2 = _386_v2
		var _387_v3 _dafny.Int
		_ = _387_v3
		_387_v3 = _dafny.IntOfInt64(-319)
		var _388_v4 _dafny.Map
		_ = _388_v4
		_388_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_387_v3, _387_v3)
		var _389_v5 _dafny.Map
		_ = _389_v5
		_389_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v0, _388_v4)
		var _390_v6 _dafny.Map
		_ = _390_v6
		_390_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_389_v5).Merge(_389_v5), true)
		if (func() bool {
			if (_390_v6).Contains(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v0, _388_v4)) {
				return (_390_v6).Get(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v0, _388_v4)).(bool)
			}
			return !(_384_v0)
		})() {
			var _391_v7 _dafny.MultiSet
			_ = _391_v7
			_391_v7 = _dafny.MultiSetOf(_384_v0)
			_384_v0 = (func() bool {
				if (_391_v7).IsProperSubsetOf(_391_v7) {
					return !(_384_v0)
				}
				return _384_v0
			})()
			(globalState).F1 = _387_v3
			var _392_v8 _dafny.Array
			_ = _392_v8
			var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw54
			_392_v8 = _nw54
			var _rhs78 _dafny.Int = Companion_Default___.SafeModuloInt(_387_v3, _387_v3)
			_ = _rhs78
			var _rhs79 _dafny.Array = _392_v8
			_ = _rhs79
			_387_v3 = _rhs78
			_392_v8 = _rhs79
			var _393_v9 _dafny.Sequence
			_ = _393_v9
			_393_v9 = _dafny.SeqOf(_384_v0)
			if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_393_v9, _393_v9), _393_v9) {
				var _394_v10 _dafny.CodePoint
				_ = _394_v10
				_394_v10 = _dafny.CodePoint('k')
				_394_v10 = _394_v10
				var _395_v11 _dafny.Array
				_ = _395_v11
				var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw55
				_395_v11 = _nw55
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_395_v11), 0))
				_ = _index71
				(_395_v11).ArraySet1(_384_v0, (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_395_v11), 0))
				_ = _index72
				(_395_v11).ArraySet1((_393_v9).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_387_v3), _dafny.IntOfUint32((_393_v9).Cardinality()))).Uint32()).(bool), (_index72).Int())
				_394_v10 = _394_v10
				var _396_v12 _dafny.Map
				_ = _396_v12
				_396_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_395_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_395_v11), 0))).Int()).(bool), _384_v0)
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_395_v11), 0))
				_ = _index73
				(_395_v11).ArraySet1((func() bool {
					if (_396_v12).Contains(false) {
						return (_396_v12).Get(false).(bool)
					}
					return !(!((_387_v3).Cmp(_387_v3) < 0))
				})(), (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_392_v8), 0))
				_ = _index74
				(_392_v8).ArraySet1(_387_v3, (_index74).Int())
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_392_v8), 0))
				_ = _index75
				(_392_v8).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_387_v3)), (_index75).Int())
			} else {
				_384_v0 = !(_384_v0) || (_384_v0)
				var _397_v13 _dafny.Array
				_ = _397_v13
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(27))
				_ = _nw56
				_397_v13 = _nw56
				_397_v13 = _397_v13
				_384_v0 = !(_384_v0)
				_384_v0 = !((_384_v0) && (_384_v0))
				var _398_v14 _dafny.CodePoint
				_ = _398_v14
				_398_v14 = _dafny.CodePoint('k')
				_398_v14 = _398_v14
			}
			_384_v0 = true
		} else {
			_389_v5 = _389_v5
			(globalState).F1 = (_387_v3).Times(_387_v3)
			_385_v1 = _385_v1
			var _399_v15 _dafny.Map
			_ = _399_v15
			_399_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v0, _dafny.IntOfInt64(-385))
			var _400_v16 _dafny.Sequence
			_ = _400_v16
			_400_v16 = _dafny.SeqOf(_dafny.IntOfInt64(-967), (_399_v15).Cardinality())
			var _401_v17 _dafny.Map
			_ = _401_v17
			_401_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_400_v16, (Companion_Default___.SafeIndex(_387_v3, _dafny.IntOfUint32((_400_v16).Cardinality()))).Uint32(), _387_v3), _384_v0)
			var _402_v18 _dafny.Set
			_ = _402_v18
			_402_v18 = _dafny.SetOf(false, !(_384_v0))
			var _403_v19 _dafny.Map
			_ = _403_v19
			_403_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_401_v17).Merge(_401_v17), !((_402_v18).IsProperSubsetOf(_402_v18)))
			_403_v19 = (_403_v19).Update((_401_v17).Merge(Companion_Default___.Fm15((_dafny.Zero).Minus(_387_v3), globalState)), _384_v0)
			var _404_v20 _dafny.Array
			_ = _404_v20
			var _nw57 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw57
			_404_v20 = _nw57
			var _405_v21 _dafny.Sequence
			_ = _405_v21
			_405_v21 = _dafny.SeqOf(_400_v16, _400_v16, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-573))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_406_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_407_i0 _dafny.Int) _dafny.Int {
					return _406_v3
				}
			})(_387_v3))))
			var _408_v22 T0
			_ = _408_v22
			var _nw58 *C0 = New_C0_()
			_ = _nw58
			_nw58.Ctor__(_384_v0, _dafny.Companion_Sequence_.Concatenate(_405_v21, _405_v21))
			_408_v22 = _nw58
			var _409_v23 _dafny.Sequence
			_ = _409_v23
			_409_v23 = _dafny.SeqOf(_384_v0)
			var _rhs80 bool = _384_v0
			_ = _rhs80
			var _rhs81 _dafny.Array = _404_v20
			_ = _rhs81
			var _rhs82 bool = (_409_v23).Select((Companion_Default___.SafeIndex(_387_v3, _dafny.IntOfUint32((_409_v23).Cardinality()))).Uint32()).(bool)
			_ = _rhs82
			var _rhs83 T0 = _408_v22
			_ = _rhs83
			_384_v0 = _rhs80
			_404_v20 = _rhs81
			_384_v0 = _rhs82
			_408_v22 = _rhs83
		}
		var _410_v24 _dafny.Sequence
		_ = _410_v24
		_410_v24 = _dafny.SeqOf(_384_v0, _384_v0, (_387_v3).Cmp(_dafny.IntOfInt64(-789)) != 0, (_387_v3).Cmp(_387_v3) >= 0)
		if (_410_v24).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_387_v3), _dafny.IntOfUint32((_410_v24).Cardinality()))).Uint32()).(bool) {
			var _411_v25 _dafny.Sequence
			_ = _411_v25
			_411_v25 = _dafny.SeqOf((_388_v4).Cardinality(), _dafny.IntOfInt64(547), _dafny.IntOfInt64(858), _387_v3, _dafny.IntOfInt64(288))
			var _412_v26 _dafny.Sequence
			_ = _412_v26
			_412_v26 = _dafny.SeqOf(_411_v25, _411_v25)
			var _413_v27 *C0
			_ = _413_v27
			var _nw59 *C0 = New_C0_()
			_ = _nw59
			_nw59.Ctor__(_384_v0, _412_v26)
			_413_v27 = _nw59
			_413_v27 = _413_v27
			var _414_v28 _dafny.MultiSet
			_ = _414_v28
			_414_v28 = _dafny.MultiSetOf(_384_v0, _384_v0, !(_384_v0))
			(globalState).F4 = _414_v28
			var _415_v29 _dafny.Array
			_ = _415_v29
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw60
			_415_v29 = _nw60
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_415_v29), 0))
			_ = _index76
			(_415_v29).ArraySet1((func() bool {
				if _384_v0 {
					return (_413_v27).F11()
				}
				return _384_v0
			})(), (_index76).Int())
			var _416_v30 D1
			_ = _416_v30
			_416_v30 = Companion_D1_.Create_DC5_((_413_v27).F11())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_415_v29), 0))
			_ = _index77
			var _rhs84 bool = (_416_v30).Dtor_cf5()
			_ = _rhs84
			var _rhs85 bool = _384_v0
			_ = _rhs85
			var _rhs86 bool = (_413_v27).F11()
			_ = _rhs86
			var _lhs61 _dafny.Array = _415_v29
			_ = _lhs61
			var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_415_v29), 0))
			_ = _lhs62
			(_lhs61).ArraySet1(_rhs84, (_lhs62).Int())
			_384_v0 = _rhs85
			_384_v0 = _rhs86
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_415_v29), 0))
			_ = _index78
			(_415_v29).ArraySet1((_this).Fm6((_413_v27.F12).Select((Companion_Default___.SafeIndex(_387_v3, _dafny.IntOfUint32((_413_v27.F12).Cardinality()))).Uint32()).(_dafny.Sequence), globalState), (_index78).Int())
			_384_v0 = _384_v0
		} else {
			var _417_v31 _dafny.Sequence
			_ = _417_v31
			_417_v31 = _dafny.SeqOf(_410_v24)
			var _418_v32 _dafny.Sequence
			_ = _418_v32
			_418_v32 = _dafny.SeqOf(_387_v3)
			_388_v4 = (_388_v4).Update((_dafny.Zero).Minus(_dafny.IntOfUint32((_417_v31).Cardinality())), (_dafny.IntOfUint32((_dafny.SeqOf(_418_v32)).Cardinality())).Plus((_dafny.Zero).Minus(_387_v3)))
			var _419_v33 _dafny.Array
			_ = _419_v33
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_6
			var _nw61 _dafny.Array
			_ = _nw61
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw61 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) bool = (func(_420_v0 bool) func(_dafny.Int) bool {
					return func(_421_i1 _dafny.Int) bool {
						return _420_v0
					}
				})(_384_v0)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw61 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw61).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw61).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_419_v33 = _nw61
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))
			_ = _index79
			(_419_v33).ArraySet1(_384_v0, (_index79).Int())
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))
			_ = _index80
			var _rhs87 bool = !(!(_384_v0)) || (_384_v0)
			_ = _rhs87
			var _rhs88 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_422_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})), _385_v1)
			_ = _rhs88
			var _rhs89 _dafny.Sequence = _385_v1
			_ = _rhs89
			var _rhs90 _dafny.Int = _dafny.IntOfInt64(535)
			_ = _rhs90
			var _lhs63 _dafny.Array = _419_v33
			_ = _lhs63
			var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))
			_ = _lhs64
			(_lhs63).ArraySet1(_rhs87, (_lhs64).Int())
			_385_v1 = _rhs88
			_385_v1 = _rhs89
			_387_v3 = _rhs90
			var _423_v34 _dafny.Array
			_ = _423_v34
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw62
			_423_v34 = _nw62
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_423_v34), 0))
			_ = _index81
			(_423_v34).ArraySet1((_387_v3).Plus(_387_v3), (_index81).Int())
			var _424_v35 _dafny.Map
			_ = _424_v35
			_424_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_419_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))).Int()).(bool), (_dafny.Zero).Minus(_387_v3))
			var _425_v36 _dafny.MultiSet
			_ = _425_v36
			_425_v36 = _dafny.MultiSetOf(_388_v4, _388_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_387_v3, _dafny.IntOfInt64(-861)))
			var _426_v37 _dafny.CodePoint
			_ = _426_v37
			_426_v37 = _dafny.CodePoint('e')
			var _427_v38 _dafny.Map
			_ = _427_v38
			_427_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_v37, _387_v3)
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_423_v34), 0))
			_ = _index82
			(_423_v34).ArraySet1((func() _dafny.Int {
				if (_424_v35).Contains((_425_v36).IsDisjointFrom(_425_v36)) {
					return (_424_v35).Get((_425_v36).IsDisjointFrom(_425_v36)).(_dafny.Int)
				}
				return ((func() _dafny.Int {
					if (_427_v38).Contains(_426_v37) {
						return (_427_v38).Get(_426_v37).(_dafny.Int)
					}
					return _387_v3
				})()).Minus(_387_v3)
			})(), (_index82).Int())
			var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))
			_ = _index83
			(_419_v33).ArraySet1((_419_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))).Int()).(bool), (_index83).Int())
			var _428_v39 D1
			_ = _428_v39
			_428_v39 = Companion_D1_.Create_DC5_((_419_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(798), _dafny.ArrayLen((_419_v33), 0))).Int()).(bool))
			var _429_v40 _dafny.Map
			_ = _429_v40
			_429_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_387_v3, _428_v39)
			var _430_v41 _dafny.Map
			_ = _430_v41
			_430_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_423_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(456), _dafny.ArrayLen((_423_v34), 0))).Int()).(_dafny.Int), _385_v1)
			_429_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_430_v41).Cardinality(), _428_v39)
		}
		var _431_v42 _dafny.Sequence
		_ = _431_v42
		_431_v42 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(299))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_432_i3 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-622)
		})))
		var _433_v43 *C0
		_ = _433_v43
		var _nw63 *C0 = New_C0_()
		_ = _nw63
		_nw63.Ctor__(_384_v0, _dafny.Companion_Sequence_.Concatenate(_431_v42, _431_v42))
		_433_v43 = _nw63
		_384_v0 = false
		var _434_i4 _dafny.Int
		_ = _434_i4
		_434_i4 = _dafny.Zero
		{
			for !(Companion_Default___.Fm0(!((_433_v43).F11()), globalState)) {
				{
					if (_434_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_434_i4 = (_434_i4).Plus(_dafny.One)
					(globalState).F1 = _387_v3
					_384_v0 = _384_v0
					if (_433_v43).F11() {
						var _435_v44 _dafny.Array
						_ = _435_v44
						var _nw64 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
						_ = _nw64
						_435_v44 = _nw64
						var _436_v45 _dafny.MultiSet
						_ = _436_v45
						_436_v45 = _dafny.MultiSetOf(_435_v44, _435_v44, _435_v44)
						_384_v0 = ((_436_v45).Union(_436_v45)).IsSubsetOf((_dafny.MultiSetOf(_435_v44)).Difference(_436_v45))
						var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_435_v44), 0))
						_ = _index84
						(_435_v44).ArraySet1(false, (_index84).Int())
						var _437_v46 _dafny.Set
						_ = _437_v46
						_437_v46 = _dafny.SetOf((_433_v43).F11())
						var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_435_v44), 0))
						_ = _index85
						(_435_v44).ArraySet1((func() bool {
							if (_433_v43).F11() {
								return (_437_v46).IsDisjointFrom(_437_v46)
							}
							return _384_v0
						})(), (_index85).Int())
						var _rhs91 _dafny.Sequence = _385_v1
						_ = _rhs91
						var _rhs92 _dafny.Array = _435_v44
						_ = _rhs92
						var _rhs93 _dafny.Int = _dafny.IntOfInt64(-241)
						_ = _rhs93
						var _lhs65 *GlobalState = globalState
						_ = _lhs65
						_385_v1 = _rhs91
						_435_v44 = _rhs92
						_lhs65.F1 = _rhs93
						var _438_v47 *C0
						_ = _438_v47
						var _nw65 *C0 = New_C0_()
						_ = _nw65
						_nw65.Ctor__(_384_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(437))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg29 _dafny.Int) interface{} {
								return coer29(arg29)
							}
						}((func(_439_v3 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
							return func(_440_i5 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(_439_v3, _439_v3)
							}
						})(_387_v3))))
						_438_v47 = _nw65
						_387_v3 = _387_v3
					} else {
						var _441_v48 _dafny.CodePoint
						_ = _441_v48
						_441_v48 = _dafny.CodePoint('p')
						var _442_v49 _dafny.Map
						_ = _442_v49
						_442_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v48, _387_v3)
						(globalState).F1 = (func() _dafny.Int {
							if (_442_v49).Contains(_441_v48) {
								return (_442_v49).Get(_441_v48).(_dafny.Int)
							}
							return _387_v3
						})()
						_441_v48 = _441_v48
						r0 = (_this).Fm7((_433_v43).F11(), globalState)
						r0 = (_this).Fm14(globalState)
						var _443_v50 _dafny.Array
						_ = _443_v50
						var _nw66 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw66
						_443_v50 = _nw66
						var _444_v51 _dafny.Map
						_ = _444_v51
						_444_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_443_v50, _dafny.CodePoint('w'))
						var _445_v52 D0
						_ = _445_v52
						_445_v52 = Companion_D0_.Create_DC2_(_444_v51, _441_v48)
						_445_v52 = _445_v52
					}
					_433_v43 = _433_v43
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		r0 = _387_v3
		return r0
	}
}
func (_this *C1) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _446_v0 bool
		_ = _446_v0
		_446_v0 = true
		_446_v0 = _446_v0
		var _447_v1 _dafny.Sequence
		_ = _447_v1
		_447_v1 = _dafny.UnicodeSeqOfUtf8Bytes("pvlac")
		_447_v1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(444))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_448_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))
		if p2 {
			var _449_v2 _dafny.Int
			_ = _449_v2
			_449_v2 = _dafny.IntOfInt64(188)
			var _450_v3 _dafny.Set
			_ = _450_v3
			_450_v3 = _dafny.SetOf(_449_v2)
			var _451_v4 _dafny.Sequence
			_ = _451_v4
			_451_v4 = _dafny.SeqOf(p2, _446_v0, false)
			var _452_v5 _dafny.Map
			_ = _452_v5
			_452_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _449_v2)
			var _453_v6 _dafny.Array
			_ = _453_v6
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_7
			var _nw67 _dafny.Array
			_ = _nw67
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw67 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) bool = func(_454_i1 _dafny.Int) bool {
					return true
				}
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw67 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw67).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw67).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_453_v6 = _nw67
			var _455_v7 _dafny.Map
			_ = _455_v7
			_455_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _453_v6)
			var _456_v8 _dafny.Map
			_ = _456_v8
			_456_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_452_v5).Cardinality(), _455_v7)
			var _457_v9 _dafny.Map
			_ = _457_v9
			_457_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_449_v2, _449_v2)
			var _458_v10 _dafny.Sequence
			_ = _458_v10
			_458_v10 = _dafny.SeqOf(_451_v4)
			var _459_v11 _dafny.Array
			_ = _459_v11
			var _nwElement0_11 bool = p0
			_ = _nwElement0_11
			var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(22))
			_ = _nw68
			(_nw68).ArraySet1(_nwElement0_11, 0)
			(_nw68).ArraySet1(false, 1)
			(_nw68).ArraySet1(p0, 2)
			(_nw68).ArraySet1(p0, 3)
			(_nw68).ArraySet1(p0, 4)
			(_nw68).ArraySet1((_450_v3).Equals(_dafny.SetOf(_449_v2)), 5)
			(_nw68).ArraySet1(p0, 6)
			(_nw68).ArraySet1(p2, 7)
			(_nw68).ArraySet1(p0, 8)
			(_nw68).ArraySet1((Companion_Default___.Fm16(_dafny.IntOfInt64(71), globalState)).IsDisjointFrom(_450_v3), 9)
			(_nw68).ArraySet1(_446_v0, 10)
			(_nw68).ArraySet1(!((func() bool {
				if p2 {
					return _446_v0
				}
				return (_451_v4).Select((Companion_Default___.SafeIndex(_449_v2, _dafny.IntOfUint32((_451_v4).Cardinality()))).Uint32()).(bool)
			})()), 11)
			(_nw68).ArraySet1(p0, 12)
			(_nw68).ArraySet1(_446_v0, 13)
			(_nw68).ArraySet1((false) || (p2), 14)
			(_nw68).ArraySet1(!((func() _dafny.Map {
				if (_456_v8).Contains((func() _dafny.Int {
					if (_457_v9).Contains(_449_v2) {
						return (_457_v9).Get(_449_v2).(_dafny.Int)
					}
					return _449_v2
				})()) {
					return (_456_v8).Get((func() _dafny.Int {
						if (_457_v9).Contains(_449_v2) {
							return (_457_v9).Get(_449_v2).(_dafny.Int)
						}
						return _449_v2
					})()).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _453_v6)
			})()).Contains(p1), 15)
			(_nw68).ArraySet1(!(p0), 16)
			(_nw68).ArraySet1(_446_v0, 17)
			(_nw68).ArraySet1(Companion_Default___.Fm0(false, globalState), 18)
			(_nw68).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_451_v4, (_458_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(148), _dafny.IntOfUint32((_458_v10).Cardinality()))).Uint32()).(_dafny.Sequence)), 19)
			(_nw68).ArraySet1(p0, 20)
			(_nw68).ArraySet1(p0, 21)
			_459_v11 = _nw68
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_459_v11), 0))
			_ = _index86
			(_459_v11).ArraySet1(true, (_index86).Int())
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_459_v11), 0))
			_ = _index87
			(_459_v11).ArraySet1(_446_v0, (_index87).Int())
			var _460_v12 _dafny.Array
			_ = _460_v12
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw69
			_460_v12 = _nw69
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))
			_ = _index88
			(_460_v12).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cf")).Cardinality())).Times(_dafny.IntOfUint32((_dafny.SeqOf(_449_v2, _449_v2, _449_v2)).Cardinality())), (_index88).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))
			_ = _index89
			(_460_v12).ArraySet1(_449_v2, (_index89).Int())
			(globalState).F1 = (_460_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))).Int()).(_dafny.Int)
			var _461_v13 _dafny.CodePoint
			_ = _461_v13
			_461_v13 = _dafny.CodePoint('k')
			_447_v1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_447_v1, _447_v1), _447_v1), (Companion_Default___.SafeIndex(_449_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_447_v1, _447_v1), _447_v1)).Cardinality()))).Uint32(), _461_v13)
			var _462_v14 D1
			_ = _462_v14
			_462_v14 = Companion_D1_.Create_DC5_(p2)
			var _463_v15 _dafny.Map
			_ = _463_v15
			_463_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v14, p2)
			var _464_v16 _dafny.MultiSet
			_ = _464_v16
			_464_v16 = _dafny.MultiSetOf(_461_v13, _461_v13, _461_v13)
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))
			_ = _index90
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_459_v11), 0))
			_ = _index91
			var _rhs94 _dafny.Int = ((_464_v16).Cardinality()).Times((_460_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))).Int()).(_dafny.Int))
			_ = _rhs94
			var _rhs95 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_449_v2, ((_460_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))).Int()).(_dafny.Int)).Minus((_457_v9).Cardinality())))
			_ = _rhs95
			var _rhs96 _dafny.Map = (_463_v15).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v14, p2))
			_ = _rhs96
			var _rhs97 bool = _446_v0
			_ = _rhs97
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			var _lhs67 _dafny.Array = _460_v12
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_460_v12), 0))
			_ = _lhs68
			var _lhs69 _dafny.Array = _459_v11
			_ = _lhs69
			var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_459_v11), 0))
			_ = _lhs70
			_lhs66.F1 = _rhs94
			(_lhs67).ArraySet1(_rhs95, (_lhs68).Int())
			_463_v15 = _rhs96
			(_lhs69).ArraySet1(_rhs97, (_lhs70).Int())
		} else {
			var _465_v17 _dafny.MultiSet
			_ = _465_v17
			_465_v17 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-754))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_466_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('q')
			})))
			_446_v0 = (_465_v17).IsSubsetOf(_465_v17)
			var _467_v18 _dafny.Set
			_ = _467_v18
			_467_v18 = _dafny.SetOf(_446_v0)
			var _468_v19 _dafny.Array
			_ = _468_v19
			var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw70
			_468_v19 = _nw70
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _index92
			(_468_v19).ArraySet1(p0, (_index92).Int())
			var _469_v20 _dafny.Int
			_ = _469_v20
			_469_v20 = _dafny.IntOfInt64(192)
			var _470_v21 D2
			_ = _470_v21
			_470_v21 = Companion_D2_.Create_DC8_(p0)
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _index93
			var _rhs98 bool = !(_446_v0)
			_ = _rhs98
			var _rhs99 _dafny.Set = (_467_v18).Difference((_467_v18).Difference(Companion_Default___.Fm17(_469_v20, _446_v0, globalState)))
			_ = _rhs99
			var _rhs100 bool = !(_446_v0) || ((_469_v20).Cmp(_469_v20) < 0)
			_ = _rhs100
			var _rhs101 bool = (func() bool {
				if _446_v0 {
					return (func() bool {
						if (func(_pat_let0_0 D2) D2 {
							return func(_471_dt__update__tmp_h0 D2) D2 {
								return func(_pat_let1_0 bool) D2 {
									return func(_472_dt__update_hcf9_h0 bool) D2 {
										return Companion_D2_.Create_DC8_(_472_dt__update_hcf9_h0)
									}(_pat_let1_0)
								}(true)
							}(_pat_let0_0)
						}(_470_v21)).Dtor_cf9() {
							return p2
						}
						return p0
					})()
				}
				return p2
			})()
			_ = _rhs101
			var _rhs102 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.Fm3(_446_v0, p2, globalState))
			_ = _rhs102
			var _lhs71 _dafny.Array = _468_v19
			_ = _lhs71
			var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _lhs72
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			_446_v0 = _rhs98
			_467_v18 = _rhs99
			(_lhs71).ArraySet1(_rhs100, (_lhs72).Int())
			_446_v0 = _rhs101
			_lhs73.F1 = _rhs102
			var _473_v22 _dafny.Map
			_ = _473_v22
			_473_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_469_v20, _469_v20)
			var _474_v23 _dafny.CodePoint
			_ = _474_v23
			_474_v23 = _dafny.CodePoint('e')
			var _475_v24 _dafny.Map
			_ = _475_v24
			_475_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.IntOfUint32((_447_v1).Cardinality()), false, _469_v20, globalState), _dafny.Companion_Sequence_.Update(_447_v1, (Companion_Default___.SafeIndex(_469_v20, _dafny.IntOfUint32((_447_v1).Cardinality()))).Uint32(), _474_v23))
			_473_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_469_v20, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_475_v24).Contains(_447_v1) {
					return (_475_v24).Get(_447_v1).(_dafny.Sequence)
				}
				return _447_v1
			})()).Cardinality())), _469_v20)
			var _476_v25 _dafny.Map
			_ = _476_v25
			_476_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(211), p2)
			var _477_v26 _dafny.Map
			_ = _477_v26
			_477_v26 = _476_v25
			var _478_v27 _dafny.Array
			_ = _478_v27
			var _nwElement0_12 _dafny.Map = _477_v26
			_ = _nwElement0_12
			var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(28))
			_ = _nw71
			(_nw71).ArraySet1(_nwElement0_12, 0)
			(_nw71).ArraySet1(_476_v25, 1)
			(_nw71).ArraySet1(_477_v26, 2)
			(_nw71).ArraySet1(_476_v25, 3)
			(_nw71).ArraySet1(_476_v25, 4)
			(_nw71).ArraySet1(_477_v26, 5)
			(_nw71).ArraySet1(_477_v26, 6)
			(_nw71).ArraySet1(_477_v26, 7)
			(_nw71).ArraySet1(_477_v26, 8)
			(_nw71).ArraySet1(Companion_Default___.Fm18(_469_v20, globalState), 9)
			(_nw71).ArraySet1(_477_v26, 10)
			(_nw71).ArraySet1(Companion_Default___.Fm18(_469_v20, globalState), 11)
			(_nw71).ArraySet1(_476_v25, 12)
			(_nw71).ArraySet1(_476_v25, 13)
			(_nw71).ArraySet1(_477_v26, 14)
			(_nw71).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfUint32((_447_v1).Cardinality()), globalState), 15)
			(_nw71).ArraySet1(_476_v25, 16)
			(_nw71).ArraySet1(_477_v26, 17)
			(_nw71).ArraySet1(_477_v26, 18)
			(_nw71).ArraySet1(_477_v26, 19)
			(_nw71).ArraySet1(_477_v26, 20)
			(_nw71).ArraySet1(_477_v26, 21)
			(_nw71).ArraySet1(_477_v26, 22)
			(_nw71).ArraySet1(_477_v26, 23)
			(_nw71).ArraySet1(_477_v26, 24)
			(_nw71).ArraySet1(_477_v26, 25)
			(_nw71).ArraySet1((func() _dafny.Map {
				if _446_v0 {
					return _477_v26
				}
				return _477_v26
			})(), 26)
			(_nw71).ArraySet1(_477_v26, 27)
			_478_v27 = _nw71
			_478_v27 = _478_v27
			var _479_v28 _dafny.Sequence
			_ = _479_v28
			_479_v28 = _dafny.SeqOf(_469_v20, _469_v20, (_469_v20).Minus(Companion_Default___.Fm3(!(p2), _446_v0, globalState)))
			var _480_v29 _dafny.Sequence
			_ = _480_v29
			_480_v29 = _dafny.SeqOf(_479_v28, _479_v28)
			var _481_v30 T0
			_ = _481_v30
			var _nw72 *C0 = New_C0_()
			_ = _nw72
			_nw72.Ctor__((_this).Fm6(_479_v28, globalState), _480_v29)
			_481_v30 = _nw72
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _index94
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _index95
			var _rhs103 bool = _446_v0
			_ = _rhs103
			var _rhs104 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_479_v28, _479_v28), (Companion_Default___.SafeIndex(_469_v20, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_479_v28, _479_v28)).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_447_v1).Cardinality())))
			_ = _rhs104
			var _rhs105 T0 = _481_v30
			_ = _rhs105
			var _rhs106 bool = (p1).IsSubsetOf(p1)
			_ = _rhs106
			var _lhs74 _dafny.Array = _468_v19
			_ = _lhs74
			var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _lhs75
			var _lhs76 _dafny.Array = _468_v19
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_468_v19), 0))
			_ = _lhs77
			(_lhs74).ArraySet1(_rhs103, (_lhs75).Int())
			_479_v28 = _rhs104
			_481_v30 = _rhs105
			(_lhs76).ArraySet1(_rhs106, (_lhs77).Int())
		}
		var _482_v31 _dafny.Int
		_ = _482_v31
		_482_v31 = _dafny.IntOfInt64(174)
		var _483_v32 _dafny.Sequence
		_ = _483_v32
		_483_v32 = _dafny.SeqOf(_482_v31, (p1).Cardinality())
		var _484_v33 _dafny.Sequence
		_ = _484_v33
		_484_v33 = _dafny.SeqOf(_483_v32)
		var _485_v34 *C0
		_ = _485_v34
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(p0, _484_v33)
		_485_v34 = _nw73
		var _486_v35 _dafny.Array
		_ = _486_v35
		var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw74
		_486_v35 = _nw74
		var _487_v36 _dafny.Map
		_ = _487_v36
		_487_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_446_v0), _447_v1)
		var _488_v37 _dafny.Array
		_ = _488_v37
		var _nwElement0_13 _dafny.Sequence = _447_v1
		_ = _nwElement0_13
		var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(8))
		_ = _nw75
		(_nw75).ArraySet1(_nwElement0_13, 0)
		(_nw75).ArraySet1((func() _dafny.Sequence {
			if (_487_v36).Contains(p2) {
				return (_487_v36).Get(p2).(_dafny.Sequence)
			}
			return _447_v1
		})(), 1)
		(_nw75).ArraySet1(_447_v1, 2)
		(_nw75).ArraySet1(_447_v1, 3)
		(_nw75).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("cucmgtuti"), 4)
		(_nw75).ArraySet1(_447_v1, 5)
		(_nw75).ArraySet1((func() _dafny.Sequence {
			if true {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(717))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_489_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('q')
				}))
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("n")
		})(), 6)
		(_nw75).ArraySet1(Companion_Default___.Fm1(_482_v31, p2, _482_v31, globalState), 7)
		_488_v37 = _nw75
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_488_v37), 0))
		_ = _index96
		(_488_v37).ArraySet1(_447_v1, (_index96).Int())
		var _490_v38 _dafny.Map
		_ = _490_v38
		_490_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_446_v0, false)
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_488_v37), 0))
		_ = _index97
		var _rhs107 bool = !(!(_490_v38).Contains(!((_485_v34).F11()) || (_446_v0)))
		_ = _rhs107
		var _rhs108 _dafny.Array = _486_v35
		_ = _rhs108
		var _rhs109 _dafny.Sequence = _447_v1
		_ = _rhs109
		var _rhs110 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_482_v31, _482_v31))
		_ = _rhs110
		var _lhs78 _dafny.Array = _488_v37
		_ = _lhs78
		var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(549), _dafny.ArrayLen((_488_v37), 0))
		_ = _lhs79
		_446_v0 = _rhs107
		_486_v35 = _rhs108
		(_lhs78).ArraySet1(_rhs109, (_lhs79).Int())
		_482_v31 = _rhs110
		var _491_v39 _dafny.Array
		_ = _491_v39
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw76
		_491_v39 = _nw76
		var _492_v40 _dafny.Array
		_ = _492_v40
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_8
		var _nw77 _dafny.Array
		_ = _nw77
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw77 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) _dafny.Int = func(_493_i4 _dafny.Int) _dafny.Int {
				return (_493_i4).Minus(_dafny.IntOfInt64(410))
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
		_492_v40 = _nw77
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_491_v39), 0))
		_ = _index98
		(_491_v39).ArraySet1(_492_v40, (_index98).Int())
		var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.ArrayLen((_491_v39), 0))
		_ = _index99
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_9
		var _nw78 _dafny.Array
		_ = _nw78
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw78 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Int = (func(_494_v31 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_495_i5 _dafny.Int) _dafny.Int {
					return (_495_i5).Minus(_494_v31)
				}
			})(_482_v31)
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
		(_491_v39).ArraySet1(_nw78, (_index99).Int())
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f5 _dafny.CodePoint
	_f6 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
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
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C2{}
var _ T3 = &C2{}
var _ T2 = &C2{}
var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F5() _dafny.CodePoint {
	return _this._f5
}
func (_this *C2) F6() bool {
	return _this._f6
}
func (_this *C2) Ctor__(f5 _dafny.CodePoint, f6 bool) {
	{
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C2) Fm9(p0 _dafny.Int, p1 D0, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.UnicodeSeqOfUtf8Bytes("vlmbv"))).Cardinality())).Cmp((_dafny.IntOfInt64(619)).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(721))).Cardinality()))) != 0
	}
}
func (_this *C2) Fm10(p0 _dafny.MultiSet, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C2) Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(292)
	}
}
func (_this *C2) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C2) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(588)).Times(_dafny.IntOfInt64(737))
	}
}
func (_this *C2) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_dafny.IntOfInt64(212)).Minus(_dafny.IntOfInt64(663)))).Plus(((_dafny.SetOf((_this).F5())).Union(_dafny.SetOf((_this).F5()))).Cardinality())
	}
}
func (_this *C2) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		var _source10 _dafny.Map = (func() _dafny.Map {
			if true {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(63), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), (_this).F6())
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(136), (_this).F6())
		})()
		_ = _source10
		var _496___mcc_h0 _dafny.Map = _source10
		_ = _496___mcc_h0
		var _497_cf12 _dafny.Map = _496___mcc_h0
		_ = _497_cf12
		return Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6())).Cardinality()))
	}
}
func (_this *C2) Fm20(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	{
		return (_dafny.MultiSetOf((_this).F6())).Difference(_dafny.MultiSetOf((_this).F6()))
	}
}
func (_this *C2) Fm21(globalState *GlobalState) _dafny.Set {
	{
		return _dafny.SetOf(false)
	}
}
func (_this *C2) M4(p0 _dafny.Array, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) (D2, _dafny.MultiSet) {
	{
		var r0 D2 = Companion_D2_.Default()
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var _498_i0 _dafny.Int
		_ = _498_i0
		_498_i0 = _dafny.Zero
		{
			for !(p2) {
				{
					if (_498_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_498_i0 = (_498_i0).Plus(_dafny.One)
					if (_this).F6() {
						var _499_v0 _dafny.Sequence
						_ = _499_v0
						_499_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ljybtapti")
						var _500_v1 D2
						_ = _500_v1
						_500_v1 = Companion_D2_.Create_DC8_(p2)
						var _501_v2 _dafny.Set
						_ = _501_v2
						_501_v2 = _dafny.SetOf((_this).F6())
						var _502_v3 _dafny.Array
						_ = _502_v3
						var _nwElement0_14 bool = (_this).F6()
						_ = _nwElement0_14
						var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(9))
						_ = _nw79
						(_nw79).ArraySet1(_nwElement0_14, 0)
						(_nw79).ArraySet1(!_dafny.Companion_Sequence_.Contains(_499_v0, (_this).F5()), 1)
						(_nw79).ArraySet1(p2, 2)
						(_nw79).ArraySet1((_500_v1).Dtor_cf9(), 3)
						(_nw79).ArraySet1((_this).F6(), 4)
						(_nw79).ArraySet1((p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm2((_this).F6(), (_this).F6(), globalState)).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), 5)
						(_nw79).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("fllshammk"), Companion_Default___.Fm22(_501_v2, _dafny.IntOfInt64(102), (_this).F6(), (_this).F6(), globalState)), 6)
						(_nw79).ArraySet1(!((_this).F6()), 7)
						(_nw79).ArraySet1(p2, 8)
						_502_v3 = _nw79
						var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))
						_ = _index100
						(_502_v3).ArraySet1(p2, (_index100).Int())
						var _503_v4 _dafny.Map
						_ = _503_v4
						_503_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
						var _504_v5 _dafny.Sequence
						_ = _504_v5
						_504_v5 = _dafny.SeqOf(_503_v4)
						var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))
						_ = _index101
						(_502_v3).ArraySet1(!((_504_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dcgthjfh")).Cardinality())), _dafny.IntOfUint32((_504_v5).Cardinality()))).Uint32()).(_dafny.Map)).Contains((_this).F6()), (_index101).Int())
						(globalState).F1 = p3
						var _505_v6 _dafny.Map
						_ = _505_v6
						_505_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p0)
						var _506_v7 _dafny.Map
						_ = _506_v7
						_506_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
							if (_505_v6).Contains(p3) {
								return (_505_v6).Get(p3).(_dafny.Array)
							}
							return p0
						})(), (_this).F6())
						var _507_v8 _dafny.Map
						_ = _507_v8
						_507_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v2, (_502_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))).Int()).(bool))
						var _508_v9 _dafny.MultiSet
						_ = _508_v9
						_508_v9 = _dafny.MultiSetOf(p2, p2, (_this).F6(), (_this).Fm10(_dafny.MultiSetFromSeq(p1), (_507_v8).Update(_501_v2, p2), p3, globalState), (_502_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))).Int()).(bool))
						_506_v7 = (_506_v7).Update(p0, (_this).Fm10(_508_v9, _507_v8, p3, globalState))
						_499_v0 = _dafny.Companion_Sequence_.Concatenate(_499_v0, Companion_Default___.Fm1(_dafny.IntOfInt64(228), (_502_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))).Int()).(bool), _dafny.IntOfInt64(343), globalState))
						var _509_v10 _dafny.Sequence
						_ = _509_v10
						_509_v10 = _dafny.SeqOf(_502_v3, _502_v3)
						var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_502_v3), 0))
						_ = _index102
						(_502_v3).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_509_v10, _dafny.SeqOf(_502_v3, _502_v3)), (_index102).Int())
					} else {
						var _510_v11 _dafny.MultiSet
						_ = _510_v11
						_510_v11 = _dafny.MultiSetOf(p3, _dafny.IntOfInt64(53), p3)
						var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))
						_ = _index103
						(p0).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p2), (_510_v11).Cardinality()), (_this).F6())).Cardinality(), p3)), (_index103).Int())
						var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))
						_ = _index104
						(p0).ArraySet1(_dafny.IntOfInt64(325), (_index104).Int())
						var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))
						_ = _index105
						(p0).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_index105).Int())
						var _511_v12 _dafny.Sequence
						_ = _511_v12
						_511_v12 = _dafny.UnicodeSeqOfUtf8Bytes("pxu")
						var _512_v13 D0
						_ = _512_v13
						_512_v13 = Companion_D0_.Create_DC1_()
						var _513_v14 _dafny.Array
						_ = _513_v14
						var _nwElement0_15 D0 = _512_v13
						_ = _nwElement0_15
						var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(23))
						_ = _nw80
						(_nw80).ArraySet1(_nwElement0_15, 0)
						(_nw80).ArraySet1(_512_v13, 1)
						(_nw80).ArraySet1(_512_v13, 2)
						(_nw80).ArraySet1(_512_v13, 3)
						(_nw80).ArraySet1(_512_v13, 4)
						(_nw80).ArraySet1(_512_v13, 5)
						(_nw80).ArraySet1(_512_v13, 6)
						(_nw80).ArraySet1(_512_v13, 7)
						(_nw80).ArraySet1(Companion_Default___.Fm23(globalState), 8)
						(_nw80).ArraySet1(_512_v13, 9)
						(_nw80).ArraySet1(_512_v13, 10)
						(_nw80).ArraySet1((func() D0 {
							if !((_this).F6()) {
								return _512_v13
							}
							return _512_v13
						})(), 11)
						(_nw80).ArraySet1(_512_v13, 12)
						(_nw80).ArraySet1(_512_v13, 13)
						(_nw80).ArraySet1(_512_v13, 14)
						(_nw80).ArraySet1(Companion_Default___.Fm23(globalState), 15)
						(_nw80).ArraySet1(_512_v13, 16)
						(_nw80).ArraySet1((func() D0 {
							if p2 {
								return _512_v13
							}
							return _512_v13
						})(), 17)
						(_nw80).ArraySet1(Companion_D0_.Create_DC1_(), 18)
						(_nw80).ArraySet1(_512_v13, 19)
						(_nw80).ArraySet1(_512_v13, 20)
						(_nw80).ArraySet1(_512_v13, 21)
						(_nw80).ArraySet1(_512_v13, 22)
						_513_v14 = _nw80
						var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_513_v14), 0))
						_ = _index106
						(_513_v14).ArraySet1(_512_v13, (_index106).Int())
						var _514_v15 bool
						_ = _514_v15
						_514_v15 = true
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_513_v14), 0))
						_ = _index107
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))
						_ = _index108
						var _rhs111 _dafny.Sequence = _511_v12
						_ = _rhs111
						var _rhs112 _dafny.Int = (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)
						_ = _rhs112
						var _rhs113 D0 = _512_v13
						_ = _rhs113
						var _rhs114 bool = _514_v15
						_ = _rhs114
						var _rhs115 _dafny.Int = (p3).Minus(p3)
						_ = _rhs115
						var _lhs80 *GlobalState = globalState
						_ = _lhs80
						var _lhs81 _dafny.Array = _513_v14
						_ = _lhs81
						var _lhs82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_513_v14), 0))
						_ = _lhs82
						var _lhs83 _dafny.Array = p0
						_ = _lhs83
						var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.ArrayLen((p0), 0))
						_ = _lhs84
						_511_v12 = _rhs111
						_lhs80.F1 = _rhs112
						(_lhs81).ArraySet1(_rhs113, (_lhs82).Int())
						_514_v15 = _rhs114
						(_lhs83).ArraySet1(_rhs115, (_lhs84).Int())
						var _515_v16 _dafny.Sequence
						_ = _515_v16
						_515_v16 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ukntgl")).Cardinality()))
						var _516_v17 _dafny.Sequence
						_ = _516_v17
						_516_v17 = _dafny.SeqOf(_515_v16)
						var _517_v18 *C0
						_ = _517_v18
						var _nw81 *C0 = New_C0_()
						_ = _nw81
						_nw81.Ctor__(_514_v15, _516_v17)
						_517_v18 = _nw81
						_514_v15 = (_517_v18).F11()
					}
					var _518_v19 bool
					_ = _518_v19
					_518_v19 = false
					var _519_v20 _dafny.Map
					_ = _519_v20
					_519_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v19, !(_518_v19))
					var _520_v21 _dafny.Map
					_ = _520_v21
					_520_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_518_v19, p3)
					var _521_v22 _dafny.MultiSet
					_ = _521_v22
					_521_v22 = _dafny.MultiSetOf(p3)
					var _522_v23 _dafny.Sequence
					_ = _522_v23
					_522_v23 = _dafny.SeqOf(_521_v22)
					var _523_v24 _dafny.Sequence
					_ = _523_v24
					_523_v24 = _dafny.UnicodeSeqOfUtf8Bytes("pvirq")
					var _524_v25 D1
					_ = _524_v25
					_524_v25 = Companion_D1_.Create_DC5_(false)
					var _525_v26 _dafny.Sequence
					_ = _525_v26
					_525_v26 = _dafny.SeqOf(p3, p3, p3, p3)
					var _526_v27 _dafny.Sequence
					_ = _526_v27
					_526_v27 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(615))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}((func(_527_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_528_i1 _dafny.Int) _dafny.Int {
							return _527_p3
						}
					})(p3))), _525_v26)
					var _529_v28 *C0
					_ = _529_v28
					var _nw82 *C0 = New_C0_()
					_ = _nw82
					_nw82.Ctor__((_this).F6(), _526_v27)
					_529_v28 = _nw82
					var _530_v29 _dafny.Map
					_ = _530_v29
					_530_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _529_v28)
					var _531_v30 _dafny.MultiSet
					_ = _531_v30
					_531_v30 = _dafny.MultiSetOf(_530_v29)
					var _532_v31 _dafny.Array
					_ = _532_v31
					var _nw83 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
					_ = _nw83
					_532_v31 = _nw83
					var _533_v32 _dafny.MultiSet
					_ = _533_v32
					_533_v32 = _dafny.MultiSetOf(_532_v31, _532_v31)
					var _534_v33 _dafny.Array
					_ = _534_v33
					var _nwElement0_16 bool = _518_v19
					_ = _nwElement0_16
					var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(21))
					_ = _nw84
					(_nw84).ArraySet1(_nwElement0_16, 0)
					(_nw84).ArraySet1((func() bool {
						if (_519_v20).Contains(_518_v19) {
							return (_519_v20).Get(_518_v19).(bool)
						}
						return p2
					})(), 1)
					(_nw84).ArraySet1((func() bool {
						if (_519_v20).Contains((_this).F6()) {
							return (_519_v20).Get((_this).F6()).(bool)
						}
						return (_this).F6()
					})(), 2)
					(_nw84).ArraySet1(!(_518_v19) || (_518_v19), 3)
					(_nw84).ArraySet1((_dafny.IntOfUint32((p1).Cardinality())).Cmp(_dafny.IntOfInt64(-89)) >= 0, 4)
					(_nw84).ArraySet1(false, 5)
					(_nw84).ArraySet1(!(Companion_Default___.Fm0(p2, globalState)), 6)
					(_nw84).ArraySet1(((_this).Fm8(p2, (_this).F6(), p2, globalState)).Cmp(p3) > 0, 7)
					(_nw84).ArraySet1(false, 8)
					(_nw84).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)).Equals(_520_v21), 9)
					(_nw84).ArraySet1(p2, 10)
					(_nw84).ArraySet1((_521_v22).IsSubsetOf(((_522_v23).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_523_v24).Cardinality()), _dafny.IntOfUint32((_522_v23).Cardinality()))).Uint32()).(_dafny.MultiSet)).Update(p3, Companion_Default___.Abs(_dafny.IntOfInt64(108)))), 11)
					(_nw84).ArraySet1((_524_v25).Dtor_cf5(), 12)
					(_nw84).ArraySet1((func() bool {
						if !(_518_v19) {
							return false
						}
						return (_this).F6()
					})(), 13)
					(_nw84).ArraySet1((p3).Cmp(p3) <= 0, 14)
					(_nw84).ArraySet1(!(_531_v30).Contains(_530_v29), 15)
					(_nw84).ArraySet1(_518_v19, 16)
					(_nw84).ArraySet1(_518_v19, 17)
					(_nw84).ArraySet1((_529_v28).F11(), 18)
					(_nw84).ArraySet1(!(((_533_v32).Cardinality()).Cmp(p3) < 0), 19)
					(_nw84).ArraySet1(((_521_v22).Update(p3, Companion_Default___.Abs(p3))).IsDisjointFrom(_521_v22), 20)
					_534_v33 = _nw84
					var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))
					_ = _index109
					(_534_v33).ArraySet1((_529_v28).F11(), (_index109).Int())
					var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))
					_ = _index110
					var _rhs116 bool = _518_v19
					_ = _rhs116
					var _rhs117 bool = ((_518_v19) && ((_529_v28).F11())) || (_518_v19)
					_ = _rhs117
					var _rhs118 _dafny.Int = p3
					_ = _rhs118
					var _rhs119 bool = p2
					_ = _rhs119
					var _rhs120 _dafny.Int = (_dafny.IntOfUint32(((func() _dafny.Sequence {
						if (_this).F6() {
							return _523_v24
						}
						return _523_v24
					})()).Cardinality())).Plus(p3)
					_ = _rhs120
					var _lhs85 _dafny.Array = _534_v33
					_ = _lhs85
					var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))
					_ = _lhs86
					var _lhs87 *GlobalState = globalState
					_ = _lhs87
					var _lhs88 *GlobalState = globalState
					_ = _lhs88
					_518_v19 = _rhs116
					(_lhs85).ArraySet1(_rhs117, (_lhs86).Int())
					_lhs87.F1 = _rhs118
					_518_v19 = _rhs119
					_lhs88.F1 = _rhs120
					var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))
					_ = _index111
					(_534_v33).ArraySet1((_529_v28).F11(), (_index111).Int())
					if (_this).F6() {
						(globalState).F1 = p3
						var _535_v34 *C1
						_ = _535_v34
						var _nw85 *C1 = New_C1_()
						_ = _nw85
						_nw85.Ctor__()
						_535_v34 = _nw85
						(globalState).F1 = ((p3).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F6()), (Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Uint32(), (_this).F6())).Cardinality()))).Times(p3)
						var _536_v35 _dafny.Array
						_ = _536_v35
						var _len0_10 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_10
						var _nw86 _dafny.Array
						_ = _nw86
						if _len0_10.Cmp(_dafny.Zero) == 0 {
							_nw86 = _dafny.NewArray(_len0_10)
						} else {
							var _init10 func(_dafny.Int) _dafny.Sequence = (func(_537_v24 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_538_i2 _dafny.Int) _dafny.Sequence {
									return _537_v24
								}
							})(_523_v24)
							_ = _init10
							var _element0_10 = _init10(_dafny.Zero)
							_ = _element0_10
							_nw86 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
							(_nw86).ArraySet1(_element0_10, 0)
							var _nativeLen0_10 = (_len0_10).Int()
							_ = _nativeLen0_10
							for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
								(_nw86).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
							}
						}
						_536_v35 = _nw86
						var _539_v36 D0
						_ = _539_v36
						_539_v36 = Companion_D0_.Create_DC0_(_518_v19)
						var _540_v37 _dafny.Map
						_ = _540_v37
						_540_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_539_v36, _518_v19)
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_536_v35), 0))
						_ = _index112
						(_536_v35).ArraySet1((func() _dafny.Sequence {
							if (func() bool {
								if (_540_v37).Contains(_539_v36) {
									return (_540_v37).Get(_539_v36).(bool)
								}
								return (_534_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))).Int()).(bool)
							})() {
								return _523_v24
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("btwvxfgd")
						})(), (_index112).Int())
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(3), _dafny.ArrayLen((_536_v35), 0))
						_ = _index113
						(_536_v35).ArraySet1(_523_v24, (_index113).Int())
						var _541_v38 _dafny.Map
						_ = _541_v38
						_541_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
						_541_v38 = _541_v38
					} else {
						var _542_v39 _dafny.CodePoint
						_ = _542_v39
						_542_v39 = _dafny.CodePoint('u')
						_542_v39 = (_this).F5()
						var _543_v40 _dafny.Set
						_ = _543_v40
						_543_v40 = _dafny.SetOf((_534_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))).Int()).(bool))
						var _544_v41 D0
						_ = _544_v41
						var _545_v42 bool
						_ = _545_v42
						var _out20 D0
						_ = _out20
						var _out21 bool
						_ = _out21
						_out20, _out21 = Companion_Default___.M0(_534_v33, Companion_Default___.Fm22(_543_v40, p3, p2, (_529_v28).F11(), globalState), (_529_v28).F11(), (_this).F5(), globalState)
						_544_v41 = _out20
						_545_v42 = _out21
						var _546_v43 _dafny.MultiSet
						_ = _546_v43
						_546_v43 = _dafny.MultiSetOf((_529_v28).F11(), _545_v42)
						var _547_v44 _dafny.MultiSet
						_ = _547_v44
						_547_v44 = _dafny.MultiSetOf(((_546_v43).Update((_529_v28).F11(), Companion_Default___.Abs(p3))).Update((_529_v28).F11(), Companion_Default___.Abs((_this).Fm4(p1, (_546_v43).Cardinality(), p3, _521_v22, globalState))), (_this).Fm20(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-7))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg34 _dafny.Int) interface{} {
								return coer34(arg34)
							}
						}((func(_548_v39 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_549_i3 _dafny.Int) _dafny.CodePoint {
								return _548_v39
							}
						})(_542_v39))), false, globalState), _546_v43)
						_547_v44 = ((_547_v44).Difference(_547_v44)).Difference((_547_v44).Union(_547_v44))
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_534_v33), 0))
						_ = _index114
						(_534_v33).ArraySet1(true, (_index114).Int())
						var _550_v45 *C1
						_ = _550_v45
						var _nw87 *C1 = New_C1_()
						_ = _nw87
						_nw87.Ctor__()
						_550_v45 = _nw87
					}
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _551_v46 _dafny.Array
		_ = _551_v46
		var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(15))
		_ = _nw88
		_551_v46 = _nw88
		var _552_v47 _dafny.Array
		_ = _552_v47
		var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(21))
		_ = _nw89
		_552_v47 = _nw89
		var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_551_v46), 0))
		_ = _index115
		(_551_v46).ArraySet1(_dafny.MultiSetOf(_552_v47), (_index115).Int())
		var _553_v48 _dafny.MultiSet
		_ = _553_v48
		_553_v48 = _dafny.MultiSetOf(_552_v47, _552_v47, _552_v47)
		var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(794), _dafny.ArrayLen((_551_v46), 0))
		_ = _index116
		(_551_v46).ArraySet1(((_dafny.MultiSetOf(_552_v47)).Update(_552_v47, Companion_Default___.Abs(Companion_Default___.Fm3(p2, p2, globalState)))).Difference(_553_v48), (_index116).Int())
		var _554_v49 _dafny.Array
		_ = _554_v49
		var _nw90 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(6))
		_ = _nw90
		_554_v49 = _nw90
		_554_v49 = _554_v49
		var _555_v50 bool
		_ = _555_v50
		_555_v50 = false
		var _556_v51 _dafny.Sequence
		_ = _556_v51
		_556_v51 = _dafny.UnicodeSeqOfUtf8Bytes("rlqrprnjt")
		var _557_v52 _dafny.Map
		_ = _557_v52
		_557_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_555_v50, _555_v50)
		var _558_v53 _dafny.Set
		_ = _558_v53
		_558_v53 = _dafny.SetOf((_dafny.Zero).Minus(p3))
		var _rhs121 bool = !(_557_v52).Equals(_557_v52)
		_ = _rhs121
		var _rhs122 bool = ((_558_v53).Intersection(_dafny.SetOf(p3))).Contains((_this).Fm8((_this).F6(), p2, p2, globalState))
		_ = _rhs122
		var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_556_v51, _dafny.UnicodeSeqOfUtf8Bytes("hs"))
		_ = _rhs123
		var _rhs124 bool = ((func() _dafny.Int {
			if (_this).F6() {
				return p3
			}
			return p3
		})()).Cmp((_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).F6() {
				return _dafny.IntOfInt64(900)
			}
			return p3
		})())) <= 0
		_ = _rhs124
		var _rhs125 bool = (_this).F6()
		_ = _rhs125
		_555_v50 = _rhs121
		_555_v50 = _rhs122
		_556_v51 = _rhs123
		_555_v50 = _rhs124
		_555_v50 = _rhs125
		var _559_v54 _dafny.Map
		_ = _559_v54
		_559_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm20(_556_v51, _555_v50, globalState), (_this).F6())
		var _560_v55 _dafny.MultiSet
		_ = _560_v55
		_560_v55 = _dafny.MultiSetOf((_this).F6(), (_this).F6(), p2)
		var _561_v56 _dafny.Sequence
		_ = _561_v56
		_561_v56 = _dafny.SeqOf(((_560_v55).Update(Companion_Default___.Fm0(p2, globalState), Companion_Default___.Abs(p3))).Update(_555_v50, Companion_Default___.Abs(p3)), _560_v55)
		_559_v54 = (_559_v54).Update((_561_v56).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((_561_v56).Cardinality()))).Uint32()).(_dafny.MultiSet), true)
		_555_v50 = (p3).Cmp(_dafny.IntOfUint32((_556_v51).Cardinality())) > 0
		r0 = Companion_D2_.Create_DC8_((Companion_Default___.Fm24(globalState)).Dtor_cf0())
		r1 = (_560_v55).Union(_560_v55)
		return r0, r1
	}
}
func (_this *C2) M5(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _562_v0 _dafny.Sequence
		_ = _562_v0
		_562_v0 = _dafny.SeqOf(p1)
		var _hi1 _dafny.Int = (_562_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_562_v0).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi1
		for _563_i0 := (_dafny.Zero).Minus(p1); _563_i0.Cmp(_hi1) < 0; _563_i0 = _563_i0.Plus(_dafny.One) {
			if (_this).F6() {
				var _564_v1 _dafny.Map
				_ = _564_v1
				_564_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vqwmkiqrg"), _563_i0)
				_564_v1 = (_564_v1).Update(p0, _563_i0)
				var _565_v2 _dafny.Array
				_ = _565_v2
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_11
				var _nw91 _dafny.Array
				_ = _nw91
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw91 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) bool = func(_566_i1 _dafny.Int) bool {
						return (_this).F6()
					}
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw91 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw91).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw91).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_565_v2 = _nw91
				var _567_v3 D0
				_ = _567_v3
				var _568_v4 bool
				_ = _568_v4
				var _out22 D0
				_ = _out22
				var _out23 bool
				_ = _out23
				_out22, _out23 = Companion_Default___.M0(_565_v2, (_this).F5(), (_this).F6(), (_this).F5(), globalState)
				_567_v3 = _out22
				_568_v4 = _out23
				var _569_v5 _dafny.Map
				_ = _569_v5
				_569_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _563_i0)
				var _570_v6 D0
				_ = _570_v6
				_570_v6 = Companion_D0_.Create_DC0_(!(_568_v4))
				_568_v4 = (_this).Fm9((_563_i0).Times((_569_v5).Cardinality()), _570_v6, (_this).F6(), p1, globalState)
				var _571_v7 D0
				_ = _571_v7
				var _572_v8 bool
				_ = _572_v8
				var _out24 D0
				_ = _out24
				var _out25 bool
				_ = _out25
				_out24, _out25 = Companion_Default___.M0(_565_v2, _dafny.CodePoint('y'), ((_this).Fm20(p0, !((_this).F6()), globalState)).Contains(_568_v4), (_this).F5(), globalState)
				_571_v7 = _out24
				_572_v8 = _out25
				var _573_v9 _dafny.Map
				_ = _573_v9
				_573_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_i0, _dafny.UnicodeSeqOfUtf8Bytes("ycs"))
				var _574_v10 _dafny.Sequence
				_ = _574_v10
				_574_v10 = _dafny.SeqOf((_this).F6(), _572_v8, (_this).F6())
				var _575_v11 _dafny.Array
				_ = _575_v11
				var _nwElement0_17 _dafny.Sequence = _574_v10
				_ = _nwElement0_17
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(29))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_17, 0)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_572_v8), _574_v10), 1)
				(_nw92).ArraySet1(_574_v10, 2)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Update(_574_v10, (Companion_Default___.SafeIndex(_563_i0, _dafny.IntOfUint32((_574_v10).Cardinality()))).Uint32(), (_this).F6()), 3)
				(_nw92).ArraySet1(_574_v10, 4)
				(_nw92).ArraySet1(_574_v10, 5)
				(_nw92).ArraySet1(_dafny.SeqOf(_572_v8), 6)
				(_nw92).ArraySet1(_574_v10, 7)
				(_nw92).ArraySet1(_574_v10, 8)
				(_nw92).ArraySet1(_574_v10, 9)
				(_nw92).ArraySet1(_574_v10, 10)
				(_nw92).ArraySet1(_574_v10, 11)
				(_nw92).ArraySet1(_574_v10, 12)
				(_nw92).ArraySet1(_574_v10, 13)
				(_nw92).ArraySet1(_574_v10, 14)
				(_nw92).ArraySet1(_574_v10, 15)
				(_nw92).ArraySet1(_574_v10, 16)
				(_nw92).ArraySet1(_574_v10, 17)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_574_v10, _574_v10), 18)
				(_nw92).ArraySet1(_574_v10, 19)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_574_v10, _574_v10), 20)
				(_nw92).ArraySet1(_574_v10, 21)
				(_nw92).ArraySet1(_574_v10, 22)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_574_v10, _574_v10), 23)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_574_v10, _574_v10), 24)
				(_nw92).ArraySet1(_574_v10, 25)
				(_nw92).ArraySet1(_dafny.SeqOf((_this).F6()), 26)
				(_nw92).ArraySet1(_574_v10, 27)
				(_nw92).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_574_v10, Companion_Default___.Fm2(_568_v4, _568_v4, globalState)), 28)
				_575_v11 = _nw92
				var _576_v12 _dafny.Sequence
				_ = _576_v12
				_576_v12 = _dafny.SeqOf(_575_v11)
				var _rhs126 _dafny.Map = _573_v9
				_ = _rhs126
				var _rhs127 _dafny.Array = (_576_v12).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_i0, p1)).Cardinality(), _dafny.IntOfUint32((_576_v12).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs127
				var _rhs128 _dafny.Int = p1
				_ = _rhs128
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				_573_v9 = _rhs126
				_575_v11 = _rhs127
				_lhs89.F1 = _rhs128
			} else {
				var _577_v13 _dafny.Array
				_ = _577_v13
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw93
				_577_v13 = _nw93
				var _578_v14 D0
				_ = _578_v14
				var _579_v15 bool
				_ = _579_v15
				var _out26 D0
				_ = _out26
				var _out27 bool
				_ = _out27
				_out26, _out27 = Companion_Default___.M0(_577_v13, (_this).F5(), (_this).F6(), _dafny.CodePoint('x'), globalState)
				_578_v14 = _out26
				_579_v15 = _out27
				var _580_v16 _dafny.Sequence
				_ = _580_v16
				_580_v16 = _dafny.SeqOf(_577_v13, _577_v13)
				var _581_v17 _dafny.Map
				_ = _581_v17
				_581_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _577_v13)
				var _582_v18 _dafny.MultiSet
				_ = _582_v18
				_582_v18 = _dafny.MultiSetOf((_580_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_562_v0).Cardinality()), _dafny.IntOfUint32((_580_v16).Cardinality()))).Uint32()).(_dafny.Array), _577_v13, (func() _dafny.Array {
					if (_581_v17).Contains(_563_i0) {
						return (_581_v17).Get(_563_i0).(_dafny.Array)
					}
					return _577_v13
				})())
				var _583_v19 _dafny.Set
				_ = _583_v19
				_583_v19 = _dafny.SetOf(false)
				_582_v18 = (_582_v18).Update(_577_v13, Companion_Default___.Abs((_dafny.Zero).Minus((_583_v19).Cardinality())))
				_577_v13 = _577_v13
				var _584_v20 _dafny.Sequence
				_ = _584_v20
				_584_v20 = _dafny.SeqOf((_563_i0).Cmp(_563_i0) > 0, (_this).F6(), _579_v15, !((_this).F6()), (_this).F6())
				_584_v20 = _dafny.Companion_Sequence_.Concatenate(_584_v20, _584_v20)
				var _585_v21 _dafny.Map
				_ = _585_v21
				_585_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((false) == (true), _dafny.IntOfUint32((p0).Cardinality()))
				_585_v21 = _585_v21
			}
			var _586_v22 _dafny.CodePoint
			_ = _586_v22
			_586_v22 = _dafny.CodePoint('a')
			_586_v22 = _586_v22
			var _587_v23 _dafny.Map
			_ = _587_v23
			_587_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F6())
			var _588_v24 _dafny.Map
			_ = _588_v24
			_588_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_562_v0, _dafny.SeqOf((_587_v23).Cardinality(), _dafny.IntOfUint32((p0).Cardinality())))).Cardinality()), p0)
			_588_v24 = _588_v24
			var _589_v25 _dafny.Array
			_ = _589_v25
			var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw94
			_589_v25 = _nw94
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_589_v25), 0))
			_ = _index117
			(_589_v25).ArraySet1(_dafny.IntOfUint32((p0).Cardinality()), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_589_v25), 0))
			_ = _index118
			(_589_v25).ArraySet1((p1).Plus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-825))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_590_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_591_i2 _dafny.Int) _dafny.CodePoint {
					return _590_v22
				}
			})(_586_v22)))).Cardinality()))), (_index118).Int())
		}
		var _592_v26 bool
		_ = _592_v26
		_592_v26 = false
		_592_v26 = _592_v26
		var _593_v27 _dafny.Map
		_ = _593_v27
		_593_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F6()) || ((_this).F6()), (func() _dafny.Int {
			if (_this).F6() {
				return p1
			}
			return p1
		})())
		var _594_v28 _dafny.Set
		_ = _594_v28
		_594_v28 = _dafny.SetOf(p0)
		var _595_v29 _dafny.Array
		_ = _595_v29
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(14)
		_ = _len0_12
		var _nw95 _dafny.Array
		_ = _nw95
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw95 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Int = (func(_596_v26 bool) func(_dafny.Int) _dafny.Int {
				return func(_597_i3 _dafny.Int) _dafny.Int {
					return (_597_i3).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_596_v26, _596_v26, (_this).F6())).Cardinality()))
				}
			})(_592_v26)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw95 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw95).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw95).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_595_v29 = _nw95
		var _598_v30 _dafny.Array
		_ = _598_v30
		_598_v30 = _595_v29
		var _599_v31 _dafny.Map
		_ = _599_v31
		_599_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_598_v30, p1)
		_593_v27 = (_593_v27).Update((_594_v28).IsDisjointFrom(_594_v28), (_599_v31).Cardinality())
		var _600_v32 _dafny.MultiSet
		_ = _600_v32
		_600_v32 = _dafny.MultiSetOf(_595_v29, _595_v29, _595_v29)
		var _601_i4 _dafny.Int
		_ = _601_i4
		_601_i4 = _dafny.Zero
		{
			for ((_600_v32).Difference(_600_v32)).IsSubsetOf(_600_v32) {
				{
					if (_601_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_601_i4 = (_601_i4).Plus(_dafny.One)
					var _602_v33 _dafny.Set
					_ = _602_v33
					_602_v33 = _dafny.SetOf((_this).F6())
					var _603_v34 _dafny.Array
					_ = _603_v34
					var _nwElement0_18 _dafny.Set = (_602_v33).Intersection(_602_v33)
					_ = _nwElement0_18
					var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(5))
					_ = _nw96
					(_nw96).ArraySet1(_nwElement0_18, 0)
					(_nw96).ArraySet1((_602_v33).Intersection(_602_v33), 1)
					(_nw96).ArraySet1(_dafny.SetOf(false, _592_v26), 2)
					(_nw96).ArraySet1((_602_v33).Difference(_602_v33), 3)
					(_nw96).ArraySet1((_this).Fm21(globalState), 4)
					_603_v34 = _nw96
					var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_603_v34), 0))
					_ = _index119
					(_603_v34).ArraySet1((_602_v33).Difference(_602_v33), (_index119).Int())
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_603_v34), 0))
					_ = _index120
					var _rhs129 _dafny.Set = (_602_v33).Union(_602_v33)
					_ = _rhs129
					var _rhs130 _dafny.Int = p1
					_ = _rhs130
					var _lhs90 _dafny.Array = _603_v34
					_ = _lhs90
					var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_603_v34), 0))
					_ = _lhs91
					var _lhs92 *GlobalState = globalState
					_ = _lhs92
					(_lhs90).ArraySet1(_rhs129, (_lhs91).Int())
					_lhs92.F1 = _rhs130
					var _604_v35 D0
					_ = _604_v35
					_604_v35 = Companion_D0_.Create_DC0_(_592_v26)
					var _605_v36 _dafny.Map
					_ = _605_v36
					_605_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
					var _606_v37 _dafny.Map
					_ = _606_v37
					_606_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25((_this).F6(), p1, _604_v35, _605_v36, globalState), _592_v26)
					var _607_v38 _dafny.Map
					_ = _607_v38
					_607_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _562_v0)
					var _608_v39 _dafny.MultiSet
					_ = _608_v39
					_608_v39 = _dafny.MultiSetOf(_592_v26, (_this).F6())
					_606_v37 = (_606_v37).Update((_607_v38).Merge(_607_v38), (p1).Cmp((_608_v39).Cardinality()) >= 0)
					var _source11 _dafny.Array = _595_v29
					_ = _source11
					var _609___mcc_h0 _dafny.Array = _source11
					_ = _609___mcc_h0
					var _610_cf13 _dafny.Array = _609___mcc_h0
					_ = _610_cf13
					_592_v26 = _592_v26
					var _611_v40 _dafny.Map
					_ = _611_v40
					_611_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), _592_v26)
					(globalState).F1 = (_dafny.Zero).Minus((_dafny.Zero).Minus(((_611_v40).Merge(_611_v40)).Cardinality()))
					_593_v27 = (_593_v27).Update((_this).F6(), Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(_592_v26, Companion_Default___.Fm0(_592_v26, globalState), globalState), p1))
					var _612_v43 _dafny.Map
					_ = _612_v43
					_612_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), p1)
					_592_v26 = (_this).Fm10(_608_v39, func() _dafny.Map {
						var _coll23 = _dafny.NewMapBuilder()
						_ = _coll23
						for _iter24 := _dafny.Iterate((func() _dafny.Map {
							var _coll24 = _dafny.NewMapBuilder()
							_ = _coll24
							for _iter25 := _dafny.Iterate((_612_v43).Keys().Elements()); ; {
								_compr_24, _ok25 := _iter25()
								if !_ok25 {
									break
								}
								var _613_v42 _dafny.Set
								_613_v42 = interface{}(_compr_24).(_dafny.Set)
								if (_612_v43).Contains(_613_v42) {
									_coll24.Add(_613_v42, (_this).F6())
								}
							}
							return _coll24.ToMap()
						}()).Keys().Elements()); ; {
							_compr_23, _ok24 := _iter24()
							if !_ok24 {
								break
							}
							var _614_v41 _dafny.Set
							_614_v41 = interface{}(_compr_23).(_dafny.Set)
							if (func() _dafny.Map {
								var _coll25 = _dafny.NewMapBuilder()
								_ = _coll25
								for _iter26 := _dafny.Iterate((_612_v43).Keys().Elements()); ; {
									_compr_25, _ok26 := _iter26()
									if !_ok26 {
										break
									}
									var _613_v42 _dafny.Set
									_613_v42 = interface{}(_compr_25).(_dafny.Set)
									if (_612_v43).Contains(_613_v42) {
										_coll25.Add(_613_v42, (_this).F6())
									}
								}
								return _coll25.ToMap()
							}()).Contains(_614_v41) {
								_coll23.Add(_614_v41, (_this).F6())
							}
						}
						return _coll23.ToMap()
					}(), p1, globalState)
					var _615_v45 _dafny.Map
					_ = _615_v45
					_615_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Set {
						var _coll26 = _dafny.NewBuilder()
						_ = _coll26
						for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(572), _dafny.IntOfInt64(521))); ; {
							_compr_26, _ok27 := _iter27()
							if !_ok27 {
								break
							}
							var _616_v44 _dafny.Int
							_616_v44 = interface{}(_compr_26).(_dafny.Int)
							if ((_dafny.IntOfInt64(572)).Cmp(_616_v44) <= 0) && ((_616_v44).Cmp(_dafny.IntOfInt64(521)) < 0) {
								_coll26.Add((_616_v44).Plus(_dafny.IntOfInt64(611)))
							}
						}
						return _coll26.ToSet()
					}(), p1)
					var _617_v46 _dafny.Map
					_ = _617_v46
					_617_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_608_v39).Cardinality(), (_this).F6())
					var _618_v47 _dafny.Set
					_ = _618_v47
					_618_v47 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(976), _617_v46)).Cardinality())
					(globalState).F1 = (_dafny.Zero).Minus((func() _dafny.Int {
						if (_615_v45).Contains(_618_v47) {
							return (_615_v45).Get(_618_v47).(_dafny.Int)
						}
						return (p1).Plus(p1)
					})())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _source12 _dafny.Array = _598_v30
		_ = _source12
		var _619___mcc_h1 _dafny.Array = _source12
		_ = _619___mcc_h1
		var _620_cf13 _dafny.Array = _619___mcc_h1
		_ = _620_cf13
		var _621_v48 _dafny.Sequence
		_ = _621_v48
		_621_v48 = _dafny.SeqOf(_562_v0)
		var _622_v49 *C0
		_ = _622_v49
		var _nw97 *C0 = New_C0_()
		_ = _nw97
		_nw97.Ctor__(((_this).F6()) && (_592_v26), _621_v48)
		_622_v49 = _nw97
		var _623_v50 _dafny.Map
		_ = _623_v50
		_623_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F6())
		_623_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _592_v26)
		var _624_v51 _dafny.Set
		_ = _624_v51
		_624_v51 = _dafny.SetOf((_this).F6(), (_622_v49).F11())
		var _625_v52 _dafny.Array
		_ = _625_v52
		var _nwElement0_19 bool = _592_v26
		_ = _nwElement0_19
		var _nw98 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(6))
		_ = _nw98
		(_nw98).ArraySet1(_nwElement0_19, 0)
		(_nw98).ArraySet1(!(_592_v26) || (!((_622_v49).F11())), 1)
		(_nw98).ArraySet1((_622_v49).F11(), 2)
		(_nw98).ArraySet1(true, 3)
		(_nw98).ArraySet1(!((_dafny.SetOf((_622_v49).F11(), _592_v26, (_622_v49).F11(), (_622_v49).F11(), (_622_v49).F11())).IsProperSubsetOf(_624_v51)), 4)
		(_nw98).ArraySet1(((_622_v49).F11()) || (false), 5)
		_625_v52 = _nw98
		var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_625_v52), 0))
		_ = _index121
		(_625_v52).ArraySet1((_622_v49).F11(), (_index121).Int())
		var _626_v53 _dafny.Array
		_ = _626_v53
		var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(28))
		_ = _nw99
		_626_v53 = _nw99
		var _627_v54 _dafny.Sequence
		_ = _627_v54
		_627_v54 = _dafny.SeqOf(_592_v26, _592_v26, (_this).F6())
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_625_v52), 0))
		_ = _index122
		var _rhs131 _dafny.Int = p1
		_ = _rhs131
		var _rhs132 bool = (func() bool {
			if (_this).F6() {
				return _592_v26
			}
			return (p1).Cmp(_dafny.IntOfUint32((_627_v54).Cardinality())) >= 0
		})()
		_ = _rhs132
		var _rhs133 bool = (p1).Cmp((_dafny.Zero).Minus(p1)) < 0
		_ = _rhs133
		var _rhs134 _dafny.Int = Companion_Default___.Fm3(Companion_Default___.Fm0(_592_v26, globalState), (func() bool {
			if true {
				return _592_v26
			}
			return _592_v26
		})(), globalState)
		_ = _rhs134
		var _rhs135 _dafny.Array = _626_v53
		_ = _rhs135
		var _lhs93 *GlobalState = globalState
		_ = _lhs93
		var _lhs94 _dafny.Array = _625_v52
		_ = _lhs94
		var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_625_v52), 0))
		_ = _lhs95
		var _lhs96 *GlobalState = globalState
		_ = _lhs96
		_lhs93.F1 = _rhs131
		_592_v26 = _rhs132
		(_lhs94).ArraySet1(_rhs133, (_lhs95).Int())
		_lhs96.F1 = _rhs134
		_626_v53 = _rhs135
		var _628_v55 _dafny.Map
		_ = _628_v55
		_628_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v26, _595_v29)
		var _629_v56 _dafny.Array
		_ = _629_v56
		var _nwElement0_20 _dafny.Array = _595_v29
		_ = _nwElement0_20
		var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(20))
		_ = _nw100
		(_nw100).ArraySet1(_nwElement0_20, 0)
		(_nw100).ArraySet1((func() _dafny.Array {
			if (_628_v55).Contains((_this).F6()) {
				return (_628_v55).Get((_this).F6()).(_dafny.Array)
			}
			return _620_cf13
		})(), 1)
		(_nw100).ArraySet1(_620_cf13, 2)
		(_nw100).ArraySet1(_595_v29, 3)
		(_nw100).ArraySet1((func() _dafny.Array {
			if (_625_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_625_v52), 0))).Int()).(bool) {
				return _595_v29
			}
			return _595_v29
		})(), 4)
		(_nw100).ArraySet1((_598_v30), 5)
		(_nw100).ArraySet1(_620_cf13, 6)
		(_nw100).ArraySet1(_620_cf13, 7)
		(_nw100).ArraySet1(_595_v29, 8)
		(_nw100).ArraySet1(_620_cf13, 9)
		(_nw100).ArraySet1(_620_cf13, 10)
		(_nw100).ArraySet1(_620_cf13, 11)
		(_nw100).ArraySet1(_595_v29, 12)
		(_nw100).ArraySet1(_620_cf13, 13)
		(_nw100).ArraySet1(_595_v29, 14)
		(_nw100).ArraySet1(_595_v29, 15)
		(_nw100).ArraySet1(_620_cf13, 16)
		(_nw100).ArraySet1(_620_cf13, 17)
		(_nw100).ArraySet1(_620_cf13, 18)
		(_nw100).ArraySet1(_620_cf13, 19)
		_629_v56 = _nw100
		_629_v56 = _629_v56
		if _592_v26 {
			_592_v26 = (func() bool {
				if (_this).F6() {
					return _592_v26
				}
				return !(!(true)) || ((_this).F6())
			})()
			(globalState).F1 = p1
			var _630_v57 _dafny.Set
			_ = _630_v57
			_630_v57 = _dafny.SetOf(p1, p1)
			_592_v26 = (_592_v26) || ((_630_v57).IsSubsetOf(_630_v57))
			if (Companion_Default___.Fm26(globalState)).Dtor_cf5() {
				var _631_v58 D2
				_ = _631_v58
				_631_v58 = Companion_D2_.Create_DC7_(p1, p1)
				var _632_v59 _dafny.Map
				_ = _632_v59
				_632_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_631_v58, p0)
				_632_v59 = (_632_v59).Update(_631_v58, p0)
				var _633_v60 *C0
				_ = _633_v60
				var _nw101 *C0 = New_C0_()
				_ = _nw101
				_nw101.Ctor__((_this).F6(), _dafny.SeqOf(_562_v0))
				_633_v60 = _nw101
				_633_v60 = _633_v60
				var _634_v61 _dafny.Array
				_ = _634_v61
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_13
				var _nw102 _dafny.Array
				_ = _nw102
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw102 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) bool = (func(_635_v26 bool) func(_dafny.Int) bool {
						return func(_636_i5 _dafny.Int) bool {
							return _635_v26
						}
					})(_592_v26)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw102 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw102).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw102).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_634_v61 = _nw102
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw103
				_634_v61 = _nw103
				var _637_v62 _dafny.Sequence
				_ = _637_v62
				_637_v62 = _dafny.SeqOf(!((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)).Contains(p1)), (_this).F6(), _592_v26)
				var _rhs136 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_637_v62, _dafny.SeqOf((_633_v60).F11(), (_this).F6(), _592_v26))
				_ = _rhs136
				var _rhs137 _dafny.Int = p1
				_ = _rhs137
				var _rhs138 _dafny.Int = p1
				_ = _rhs138
				var _rhs139 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_562_v0, _dafny.Companion_Sequence_.Concatenate(_562_v0, _562_v0)), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_562_v0, _dafny.Companion_Sequence_.Concatenate(_562_v0, _562_v0))).Cardinality()))).Uint32(), p1)
				_ = _rhs139
				var _lhs97 *GlobalState = globalState
				_ = _lhs97
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				_637_v62 = _rhs136
				_lhs97.F1 = _rhs137
				_lhs98.F1 = _rhs138
				_562_v0 = _rhs139
				var _638_v63 _dafny.MultiSet
				_ = _638_v63
				_638_v63 = _dafny.MultiSetOf((p1).Plus(_dafny.IntOfInt64(-917)))
				_638_v63 = _638_v63
			} else {
				(globalState).F1 = (p1).Plus(p1)
				var _639_v64 _dafny.Array
				_ = _639_v64
				var _nw104 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw104
				_639_v64 = _nw104
				var _640_v65 _dafny.Sequence
				_ = _640_v65
				_640_v65 = _dafny.SeqOf(_593_v27)
				var _641_v66 _dafny.MultiSet
				_ = _641_v66
				_641_v66 = _dafny.MultiSetOf(_592_v26)
				var _642_v67 _dafny.Set
				_ = _642_v67
				_642_v67 = _dafny.SetOf(_592_v26, (_this).F6())
				var _643_v68 _dafny.Map
				_ = _643_v68
				_643_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_642_v67, _592_v26)
				var _rhs140 _dafny.Array = _639_v64
				_ = _rhs140
				var _rhs141 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((p0).Cardinality()), p1)).Plus(_dafny.IntOfInt64(349))
				_ = _rhs141
				var _rhs142 _dafny.Int = (p1).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(24))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_644_p1 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_645_i6 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _644_p1)
					}
				})(p1))), _640_v65)).Cardinality()))
				_ = _rhs142
				var _rhs143 bool = (_this).Fm10(_641_v66, _643_v68, (p1).Plus(p1), globalState)
				_ = _rhs143
				var _lhs99 *GlobalState = globalState
				_ = _lhs99
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				_639_v64 = _rhs140
				_lhs99.F1 = _rhs141
				_lhs100.F1 = _rhs142
				_592_v26 = _rhs143
				(globalState).F1 = (_dafny.Zero).Minus((_this).Fm7(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_562_v0, _562_v0)), globalState))
				_592_v26 = (_this).F6()
				var _646_v69 *C1
				_ = _646_v69
				var _nw105 *C1 = New_C1_()
				_ = _nw105
				_nw105.Ctor__()
				_646_v69 = _nw105
			}
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_595_v29), 0))
			_ = _index123
			(_595_v29).ArraySet1(p1, (_index123).Int())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen((_595_v29), 0))
			_ = _index124
			(_595_v29).ArraySet1((p1).Plus(p1), (_index124).Int())
		} else {
			var _647_v70 _dafny.Array
			_ = _647_v70
			var _nw106 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
			_ = _nw106
			_647_v70 = _nw106
			var _648_v71 D0
			_ = _648_v71
			_648_v71 = Companion_D0_.Create_DC0_(false)
			var _649_v72 D0
			_ = _649_v72
			_649_v72 = Companion_D0_.Create_DC3_(_648_v71)
			var _rhs144 _dafny.Array = _647_v70
			_ = _rhs144
			var _rhs145 bool = (p1).Cmp((_562_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_562_v0).Cardinality()))).Uint32()).(_dafny.Int)) >= 0
			_ = _rhs145
			var _rhs146 D0 = Companion_D0_.Create_DC3_(_648_v71)
			_ = _rhs146
			_647_v70 = _rhs144
			_592_v26 = _rhs145
			_649_v72 = _rhs146
			var _650_v73 _dafny.Array
			_ = _650_v73
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_14
			var _nw107 _dafny.Array
			_ = _nw107
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw107 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Sequence = (func(_651_v26 bool, _652_p0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_653_i7 _dafny.Int) _dafny.Sequence {
						return (func() _dafny.Sequence {
							if _651_v26 {
								return _652_p0
							}
							return _dafny.UnicodeSeqOfUtf8Bytes("cqdfgekt")
						})()
					}
				})(_592_v26, p0)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw107 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw107).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw107).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_650_v73 = _nw107
			var _rhs147 _dafny.Int = (_dafny.Zero).Minus(p1)
			_ = _rhs147
			var _rhs148 _dafny.Int = ((p1).Times(p1)).Minus((_dafny.IntOfUint32((p0).Cardinality())).Plus(p1))
			_ = _rhs148
			var _rhs149 _dafny.Int = ((p1).Times(p1)).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-35), p1))
			_ = _rhs149
			var _rhs150 _dafny.Array = (func() _dafny.Array {
				if !((_this).F6()) {
					return _650_v73
				}
				return (func() _dafny.Array {
					if (_this).F6() {
						return _650_v73
					}
					return _650_v73
				})()
			})()
			_ = _rhs150
			var _lhs101 *GlobalState = globalState
			_ = _lhs101
			var _lhs102 *GlobalState = globalState
			_ = _lhs102
			var _lhs103 *GlobalState = globalState
			_ = _lhs103
			_lhs101.F1 = _rhs147
			_lhs102.F1 = _rhs148
			_lhs103.F1 = _rhs149
			_650_v73 = _rhs150
			var _654_v74 _dafny.MultiSet
			_ = _654_v74
			_654_v74 = _dafny.MultiSetOf((_this).F6())
			var _655_v75 _dafny.Set
			_ = _655_v75
			_655_v75 = _dafny.SetOf(true, (_this).F6(), _592_v26, _592_v26)
			var _656_v76 _dafny.Map
			_ = _656_v76
			_656_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_655_v75, true)
			var _657_v77 _dafny.Sequence
			_ = _657_v77
			_657_v77 = _dafny.SeqOf((false) || ((_this).Fm10(_654_v74, _656_v76, p1, globalState)), true, _592_v26)
			if (_657_v77).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_562_v0).Cardinality())), _dafny.IntOfUint32((_657_v77).Cardinality()))).Uint32()).(bool) {
				var _658_v78 _dafny.Sequence
				_ = _658_v78
				_658_v78 = _dafny.UnicodeSeqOfUtf8Bytes("oomt")
				_658_v78 = p0
				var _659_v80 _dafny.Set
				_ = _659_v80
				_659_v80 = _dafny.SetOf(p1)
				_592_v26 = (_659_v80).IsSubsetOf(func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter28 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-802))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}(func(_660_i8 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(29)
					}))).Elements()); ; {
						_compr_27, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _661_v79 _dafny.Int
						_661_v79 = interface{}(_compr_27).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-802))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}(func(_660_i8 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(29)
						})), _661_v79) {
							_coll27.Add((_661_v79).Minus(_dafny.IntOfInt64(-659)))
						}
					}
					return _coll27.ToSet()
				}())
				var _662_v81 _dafny.CodePoint
				_ = _662_v81
				_662_v81 = _dafny.CodePoint('v')
				var _663_v82 D1
				_ = _663_v82
				_663_v82 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(683))
				var _rhs151 _dafny.Array = _647_v70
				_ = _rhs151
				var _rhs152 _dafny.Array = _647_v70
				_ = _rhs152
				var _rhs153 _dafny.CodePoint = _dafny.CodePoint('k')
				_ = _rhs153
				var _rhs154 _dafny.Int = (_663_v82).Dtor_cf4()
				_ = _rhs154
				var _lhs104 *GlobalState = globalState
				_ = _lhs104
				_647_v70 = _rhs151
				_647_v70 = _rhs152
				_662_v81 = _rhs153
				_lhs104.F1 = _rhs154
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_647_v70), 0))
				_ = _index125
				(_647_v70).ArraySet1((_592_v26) == (_592_v26), (_index125).Int())
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(633), _dafny.ArrayLen((_647_v70), 0))
				_ = _index126
				(_647_v70).ArraySet1(true, (_index126).Int())
				_592_v26 = ((_this).Fm21(globalState)).IsDisjointFrom(_dafny.SetOf((_this).F6(), (_this).F6()))
			} else {
				var _664_v83 *C1
				_ = _664_v83
				var _nw108 *C1 = New_C1_()
				_ = _nw108
				_nw108.Ctor__()
				_664_v83 = _nw108
				var _665_v84 _dafny.Map
				_ = _665_v84
				_665_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_592_v26, _664_v83)
				_665_v84 = _665_v84
				var _666_v85 _dafny.Array
				_ = _666_v85
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
				_ = _nw109
				_666_v85 = _nw109
				_666_v85 = _666_v85
				(globalState).F1 = _dafny.IntOfInt64(961)
				_592_v26 = (_this).F6()
				var _667_v86 *C1
				_ = _667_v86
				var _nw110 *C1 = New_C1_()
				_ = _nw110
				_nw110.Ctor__()
				_667_v86 = _nw110
			}
			(globalState).F1 = p1
			(globalState).F1 = p1
		}
	}
}
func (_this *C2) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		r0 = !_dafny.Companion_Sequence_.Equal(p2, p2)
		var _668_v0 _dafny.Array
		_ = _668_v0
		var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
		_ = _nw111
		_668_v0 = _nw111
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))
		_ = _index127
		(_668_v0).ArraySet1(p1, (_index127).Int())
		var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))
		_ = _index128
		(_668_v0).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), (_index128).Int())
		var _669_v1 _dafny.Sequence
		_ = _669_v1
		_669_v1 = _dafny.SeqOf(p0, (_this).F6())
		var _670_v2 D1
		_ = _670_v2
		_670_v2 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_669_v1).Cardinality()))
		var _pat_let_tv8 = _668_v0
		_ = _pat_let_tv8
		var _pat_let_tv9 = _668_v0
		_ = _pat_let_tv9
		var _hi2 _dafny.Int = p1
		_ = _hi2
		for _671_i0 := (func(_pat_let4_0 D1) D1 {
			return func(_680_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let5_0 _dafny.Int) D1 {
					return func(_681_dt__update_hcf4_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC4_(_681_dt__update_hcf4_h0)
					}(_pat_let5_0)
				}((_pat_let_tv9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_pat_let_tv8), 0))).Int()).(_dafny.Int))
			}(_pat_let4_0)
		}(_670_v2)).Dtor_cf4(); _671_i0.Cmp(_hi2) < 0; _671_i0 = _671_i0.Plus(_dafny.One) {
			var _672_v3 _dafny.Array
			_ = _672_v3
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
			_ = _nw112
			_672_v3 = _nw112
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _index129
			(_672_v3).ArraySet1((_this).F6(), (_index129).Int())
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _index130
			(_672_v3).ArraySet1((_this).F6(), (_index130).Int())
			r0 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ocyq"), p2)).Cardinality())).Cmp((p1).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_671_i0)))) >= 0
			var _673_v4 _dafny.Sequence
			_ = _673_v4
			_673_v4 = _dafny.SeqOf((_668_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))).Int()).(_dafny.Int))
			var _674_v5 D6
			_ = _674_v5
			_674_v5 = Companion_D6_.Create_DC13_(_673_v4)
			r0 = !_dafny.Companion_Sequence_.Equal((_674_v5).Dtor_cf14(), _673_v4)
			var _675_v6 _dafny.Array
			_ = _675_v6
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(4))
			_ = _nw113
			_675_v6 = _nw113
			var _676_v7 _dafny.Map
			_ = _676_v7
			_676_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_v3, (_this).F5())
			var _677_v8 D0
			_ = _677_v8
			_677_v8 = Companion_D0_.Create_DC2_(_676_v7, (_this).F5())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_675_v6), 0))
			_ = _index131
			(_675_v6).ArraySet1((func() D0 {
				if p0 {
					return _677_v8
				}
				return _677_v8
			})(), (_index131).Int())
			var _pat_let_tv6 = _676_v7
			_ = _pat_let_tv6
			var _pat_let_tv7 = _676_v7
			_ = _pat_let_tv7
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))
			_ = _index132
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _index133
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _index134
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_675_v6), 0))
			_ = _index135
			var _rhs155 _dafny.Int = (p1).Plus((_668_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))).Int()).(_dafny.Int))
			_ = _rhs155
			var _rhs156 bool = !((func() bool {
				if (_this).F6() {
					return !((_671_i0).Cmp(_dafny.IntOfUint32((_669_v1).Cardinality())) != 0)
				}
				return (func() bool {
					if (_this).F6() {
						return (_672_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))).Int()).(bool)
					}
					return (_this).F6()
				})()
			})())
			_ = _rhs156
			var _rhs157 bool = (_671_i0).Cmp(p1) == 0
			_ = _rhs157
			var _rhs158 D0 = func(_pat_let2_0 D0) D0 {
				return func(_678_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let3_0 _dafny.Map) D0 {
						return func(_679_dt__update_hcf1_h0 _dafny.Map) D0 {
							return Companion_D0_.Create_DC2_(_679_dt__update_hcf1_h0, (_678_dt__update__tmp_h1).Dtor_cf2())
						}(_pat_let3_0)
					}((_pat_let_tv6).Merge(_pat_let_tv7))
				}(_pat_let2_0)
			}(_677_v8)
			_ = _rhs158
			var _rhs159 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p1))
			_ = _rhs159
			var _lhs105 _dafny.Array = _668_v0
			_ = _lhs105
			var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))
			_ = _lhs106
			var _lhs107 _dafny.Array = _672_v3
			_ = _lhs107
			var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _lhs108
			var _lhs109 _dafny.Array = _672_v3
			_ = _lhs109
			var _lhs110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(660), _dafny.ArrayLen((_672_v3), 0))
			_ = _lhs110
			var _lhs111 _dafny.Array = _675_v6
			_ = _lhs111
			var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(868), _dafny.ArrayLen((_675_v6), 0))
			_ = _lhs112
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			(_lhs105).ArraySet1(_rhs155, (_lhs106).Int())
			(_lhs107).ArraySet1(_rhs156, (_lhs108).Int())
			(_lhs109).ArraySet1(_rhs157, (_lhs110).Int())
			(_lhs111).ArraySet1(_rhs158, (_lhs112).Int())
			_lhs113.F1 = _rhs159
		}
		var _682_v9 _dafny.MultiSet
		_ = _682_v9
		_682_v9 = _dafny.MultiSetOf(p0)
		var _683_v10 _dafny.Set
		_ = _683_v10
		_683_v10 = _dafny.SetOf(p0, (_this).F6())
		var _684_v11 _dafny.Map
		_ = _684_v11
		_684_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_683_v10, p0)
		var _685_v12 _dafny.Map
		_ = _685_v12
		_685_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p1)
		var _686_v13 _dafny.Map
		_ = _686_v13
		_686_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_685_v12).Cardinality())
		var _rhs160 bool = (_this).Fm10(_682_v9, (_684_v11).Merge(_684_v11), (func() _dafny.Int {
			if true {
				return (_686_v13).Cardinality()
			}
			return (_668_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))).Int()).(_dafny.Int)
		})(), globalState)
		_ = _rhs160
		var _rhs161 _dafny.Array = _668_v0
		_ = _rhs161
		var _rhs162 bool = (_this).F6()
		_ = _rhs162
		r0 = _rhs160
		_668_v0 = _rhs161
		r0 = _rhs162
		var _687_v14 _dafny.Sequence
		_ = _687_v14
		_687_v14 = _dafny.SeqOf((_668_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))).Int()).(_dafny.Int))
		var _688_v15 _dafny.Map
		_ = _688_v15
		_688_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_687_v14).Cardinality()), p0)
		var _689_v16 _dafny.Sequence
		_ = _689_v16
		_689_v16 = _dafny.SeqOf((_688_v15).Cardinality(), p1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-294))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg39 _dafny.Int) interface{} {
				return coer39(arg39)
			}
		}(func(_690_i1 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		}))).Cardinality()), p1)
		r0 = _dafny.Companion_Sequence_.Equal(_689_v16, _687_v14)
		var _hi3 _dafny.Int = p1
		_ = _hi3
		for _691_i2 := p1; _691_i2.Cmp(_hi3) < 0; _691_i2 = _691_i2.Plus(_dafny.One) {
			r0 = (_685_v12).Contains((_this).F6())
			var _692_v17 _dafny.Array
			_ = _692_v17
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_15
			var _nw114 _dafny.Array
			_ = _nw114
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw114 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) bool = func(_693_i3 _dafny.Int) bool {
					return !((_this).F6())
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw114 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw114).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw114).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_692_v17 = _nw114
			var _694_v18 D0
			_ = _694_v18
			_694_v18 = Companion_D0_.Create_DC0_((_this).F6())
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_692_v17), 0))
			_ = _index136
			(_692_v17).ArraySet1((_this).Fm9(_691_i2, _694_v18, (_this).F6(), _dafny.IntOfInt64(236), globalState), (_index136).Int())
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_692_v17), 0))
			_ = _index137
			(_692_v17).ArraySet1((_669_v1).Select((Companion_Default___.SafeIndex((_668_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_668_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_669_v1).Cardinality()))).Uint32()).(bool), (_index137).Int())
			var _695_v19 _dafny.Map
			_ = _695_v19
			_695_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _668_v0)
			_668_v0 = (func() _dafny.Array {
				if (_695_v19).Contains((_this).F5()) {
					return (_695_v19).Get((_this).F5()).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if !(p0) {
						return _668_v0
					}
					return _668_v0
				})()
			})()
			var _696_v20 _dafny.Array
			_ = _696_v20
			_696_v20 = _668_v0
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_692_v17), 0))
			_ = _index138
			var _rhs163 bool = (_682_v9).IsDisjointFrom(_682_v9)
			_ = _rhs163
			var _rhs164 _dafny.Array = _696_v20
			_ = _rhs164
			var _lhs114 _dafny.Array = _692_v17
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_692_v17), 0))
			_ = _lhs115
			(_lhs114).ArraySet1(_rhs163, (_lhs115).Int())
			_696_v20 = _rhs164
		}
		r0 = (_dafny.SetOf(p0, p0, p0)).IsSubsetOf((_683_v10).Intersection(_683_v10))
		r1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(425))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg40 _dafny.Int) interface{} {
				return coer40(arg40)
			}
		}(func(_697_i4 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})), _dafny.UnicodeSeqOfUtf8Bytes("jr"))
		return r0, r1
	}
}
func (_this *C2) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _698_v0 bool
		_ = _698_v0
		_698_v0 = false
		var _699_v1 _dafny.Array
		_ = _699_v1
		var _nw115 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(10))
		_ = _nw115
		_699_v1 = _nw115
		var _700_v2 _dafny.Set
		_ = _700_v2
		_700_v2 = _dafny.SetOf(_699_v1)
		_698_v0 = (_700_v2).IsProperSubsetOf(_700_v2)
		var _701_v3 _dafny.Int
		_ = _701_v3
		_701_v3 = _dafny.IntOfInt64(-976)
		var _702_v4 _dafny.Map
		_ = _702_v4
		_702_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(_698_v0, _698_v0, (_this).F6(), globalState), _701_v3)
		var _703_v5 _dafny.Map
		_ = _703_v5
		_703_v5 = _702_v4
		var _704_v6 _dafny.Sequence
		_ = _704_v6
		_704_v6 = _dafny.SeqOf((_this).F5())
		_703_v5 = Companion_Default___.Fm27(_704_v6, _698_v0, _698_v0, globalState)
		r0 = _701_v3
		var _705_v8 D2
		_ = _705_v8
		_705_v8 = Companion_D2_.Create_DC7_(_701_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(488))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}(func(_706_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		}))).Cardinality()))
		var _707_v9 _dafny.Set
		_ = _707_v9
		_707_v9 = _dafny.SetOf(_705_v8)
		var _708_v10 _dafny.Map
		_ = _708_v10
		_708_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), func() _dafny.Map {
			var _coll28 = _dafny.NewMapBuilder()
			_ = _coll28
			for _iter29 := _dafny.Iterate((_707_v9).Elements()); ; {
				_compr_28, _ok29 := _iter29()
				if !_ok29 {
					break
				}
				var _709_v7 D2
				_709_v7 = interface{}(_compr_28).(D2)
				if (_707_v9).Contains(_709_v7) {
					_coll28.Add(_709_v7, (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("tfqu"))).Cardinality())
				}
			}
			return _coll28.ToMap()
		}())
		var _710_v11 _dafny.Map
		_ = _710_v11
		_710_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_705_v8, _701_v3)
		_708_v10 = (_708_v10).Update((_this).F5(), _710_v11)
		var _711_v12 *C1
		_ = _711_v12
		var _nw116 *C1 = New_C1_()
		_ = _nw116
		_nw116.Ctor__()
		_711_v12 = _nw116
		var _712_v13 _dafny.Set
		_ = _712_v13
		_712_v13 = _dafny.SetOf(_711_v12)
		var _713_i1 _dafny.Int
		_ = _713_i1
		_713_i1 = _dafny.Zero
		{
			for (_712_v13).IsSubsetOf(_712_v13) {
				{
					if (_713_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_713_i1 = (_713_i1).Plus(_dafny.One)
					var _714_v14 _dafny.Map
					_ = _714_v14
					_714_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(308))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}(func(_715_i2 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					})), (_this).F6())
					_714_v14 = _714_v14
					r0 = _701_v3
					_698_v0 = (true) || (_698_v0)
					var _716_v15 _dafny.CodePoint
					_ = _716_v15
					_716_v15 = _dafny.CodePoint('w')
					_716_v15 = _716_v15
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _717_i3 _dafny.Int
		_ = _717_i3
		_717_i3 = _dafny.Zero
		{
			for _698_v0 {
				{
					if (_717_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_717_i3 = (_717_i3).Plus(_dafny.One)
					if (_this).F6() {
						var _718_v16 _dafny.Array
						_ = _718_v16
						var _len0_16 _dafny.Int = _dafny.IntOfInt64(15)
						_ = _len0_16
						var _nw117 _dafny.Array
						_ = _nw117
						if _len0_16.Cmp(_dafny.Zero) == 0 {
							_nw117 = _dafny.NewArray(_len0_16)
						} else {
							var _init16 func(_dafny.Int) bool = (func(_719_v0 bool) func(_dafny.Int) bool {
								return func(_720_i4 _dafny.Int) bool {
									return _719_v0
								}
							})(_698_v0)
							_ = _init16
							var _element0_16 = _init16(_dafny.Zero)
							_ = _element0_16
							_nw117 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
							(_nw117).ArraySet1(_element0_16, 0)
							var _nativeLen0_16 = (_len0_16).Int()
							_ = _nativeLen0_16
							for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
								(_nw117).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
							}
						}
						_718_v16 = _nw117
						var _721_v17 _dafny.Map
						_ = _721_v17
						_721_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _718_v16)
						(globalState).F1 = (_721_v17).Cardinality()
						var _722_v18 _dafny.Array
						_ = _722_v18
						var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
						_ = _nw118
						_722_v18 = _nw118
						_722_v18 = _722_v18
						var _723_v19 _dafny.MultiSet
						_ = _723_v19
						_723_v19 = _dafny.MultiSetOf(_718_v16, _718_v16)
						_698_v0 = ((_723_v19).Cardinality()).Cmp(_dafny.IntOfUint32((_704_v6).Cardinality())) >= 0
						var _724_v20 _dafny.Sequence
						_ = _724_v20
						_724_v20 = _dafny.SeqOf(_698_v0, (_this).F6())
						r0 = (Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_701_v3), _701_v3)).Minus(Companion_Default___.Fm3((_724_v20).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_724_v20).Cardinality()), _dafny.IntOfUint32((_724_v20).Cardinality()))).Uint32()).(bool), Companion_Default___.Fm0(false, globalState), globalState))
						var _725_v21 _dafny.MultiSet
						_ = _725_v21
						_725_v21 = _dafny.MultiSetOf(_711_v12)
						var _726_v22 _dafny.Sequence
						_ = _726_v22
						_726_v22 = _dafny.SeqOf(_701_v3)
						var _727_v23 _dafny.MultiSet
						_ = _727_v23
						_727_v23 = _dafny.MultiSetOf(_701_v3)
						var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_722_v18), 0))
						_ = _index139
						(_722_v18).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
							if (_725_v21).Contains(_711_v12) {
								return (_725_v21).Multiplicity(_711_v12)
							}
							return (_711_v12).Fm4(_dafny.Companion_Sequence_.Update(_724_v20, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_724_v20).Cardinality()), _dafny.IntOfUint32((_724_v20).Cardinality()))).Uint32(), (_711_v12).Fm6(_726_v22, globalState)), _701_v3, _701_v3, _727_v23, globalState)
						})()), (_index139).Int())
						var _728_v24 _dafny.Sequence
						_ = _728_v24
						_728_v24 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(761))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}(func(_729_i5 _dafny.Int) _dafny.CodePoint {
							return (_this).F5()
						})), _704_v6, _dafny.UnicodeSeqOfUtf8Bytes("ea"), _704_v6, _704_v6)
						var _730_v25 _dafny.Sequence
						_ = _730_v25
						_730_v25 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_728_v24, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(93))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg44 _dafny.Int) interface{} {
								return coer44(arg44)
							}
						}((func(_731_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_732_i6 _dafny.Int) _dafny.Sequence {
								return _731_v6
							}
						})(_704_v6)))), _728_v24)
						var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(972), _dafny.ArrayLen((_722_v18), 0))
						_ = _index140
						(_722_v18).ArraySet1(_dafny.IntOfUint32(((_730_v25).Select((Companion_Default___.SafeIndex(_701_v3, _dafny.IntOfUint32((_730_v25).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), (_index140).Int())
					} else {
						var _733_v26 _dafny.Array
						_ = _733_v26
						var _len0_17 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_17
						var _nw119 _dafny.Array
						_ = _nw119
						if _len0_17.Cmp(_dafny.Zero) == 0 {
							_nw119 = _dafny.NewArray(_len0_17)
						} else {
							var _init17 func(_dafny.Int) _dafny.Int = (func(_734_v6 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
								return func(_735_i7 _dafny.Int) _dafny.Int {
									return (_735_i7).Minus((_dafny.SetOf(_734_v6, _734_v6)).Cardinality())
								}
							})(_704_v6)
							_ = _init17
							var _element0_17 = _init17(_dafny.Zero)
							_ = _element0_17
							_nw119 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
							(_nw119).ArraySet1(_element0_17, 0)
							var _nativeLen0_17 = (_len0_17).Int()
							_ = _nativeLen0_17
							for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
								(_nw119).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
							}
						}
						_733_v26 = _nw119
						var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_733_v26), 0))
						_ = _index141
						(_733_v26).ArraySet1(_dafny.IntOfUint32((_704_v6).Cardinality()), (_index141).Int())
						var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_733_v26), 0))
						_ = _index142
						(_733_v26).ArraySet1((func() _dafny.Int {
							if (_701_v3).Cmp((_dafny.Zero).Minus(_701_v3)) >= 0 {
								return _701_v3
							}
							return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_704_v6).Cardinality()), (_702_v4).Cardinality())
						})(), (_index142).Int())
						_733_v26 = _733_v26
						var _736_v27 _dafny.Map
						_ = _736_v27
						_736_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_698_v0, (_this).F6())
						var _737_v28 _dafny.Sequence
						_ = _737_v28
						_737_v28 = _dafny.SeqOf((_this).Fm7(false, globalState))
						_736_v27 = (_736_v27).Update((_this).F6(), ((_this).Fm6(_737_v28, globalState)) || (_698_v0))
						var _738_v29 _dafny.Sequence
						_ = _738_v29
						_738_v29 = _dafny.SeqOf(_698_v0)
						var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_733_v26), 0))
						_ = _index143
						(_733_v26).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_738_v29, _738_v29), _dafny.Companion_Sequence_.Update(_738_v29, (Companion_Default___.SafeIndex(_701_v3, _dafny.IntOfUint32((_738_v29).Cardinality()))).Uint32(), (_this).F6()))).Cardinality()), (_index143).Int())
						var _739_v30 _dafny.Array
						_ = _739_v30
						var _nw120 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw120
						_739_v30 = _nw120
						var _740_v31 _dafny.CodePoint
						_ = _740_v31
						_740_v31 = _dafny.CodePoint('c')
						var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_733_v26), 0))
						_ = _index144
						var _rhs165 _dafny.Int = (_737_v28).Select((Companion_Default___.SafeIndex(_701_v3, _dafny.IntOfUint32((_737_v28).Cardinality()))).Uint32()).(_dafny.Int)
						_ = _rhs165
						var _rhs166 _dafny.Sequence = _704_v6
						_ = _rhs166
						var _rhs167 _dafny.Array = _739_v30
						_ = _rhs167
						var _rhs168 _dafny.CodePoint = (_this).F5()
						_ = _rhs168
						var _rhs169 _dafny.CodePoint = (_this).F5()
						_ = _rhs169
						var _lhs116 _dafny.Array = _733_v26
						_ = _lhs116
						var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.ArrayLen((_733_v26), 0))
						_ = _lhs117
						(_lhs116).ArraySet1(_rhs165, (_lhs117).Int())
						_704_v6 = _rhs166
						_739_v30 = _rhs167
						_740_v31 = _rhs168
						_740_v31 = _rhs169
					}
					var _741_v32 _dafny.Array
					_ = _741_v32
					var _nw121 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
					_ = _nw121
					_741_v32 = _nw121
					var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_741_v32), 0))
					_ = _index145
					(_741_v32).ArraySet1(_698_v0, (_index145).Int())
					var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_741_v32), 0))
					_ = _index146
					(_741_v32).ArraySet1(_698_v0, (_index146).Int())
					(globalState).F1 = (_dafny.Zero).Minus(_701_v3)
					var _742_v33 _dafny.Set
					_ = _742_v33
					_742_v33 = _dafny.SetOf(_698_v0)
					var _743_v34 _dafny.Array
					_ = _743_v34
					var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
					_ = _nw122
					_743_v34 = _nw122
					var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_743_v34), 0))
					_ = _index147
					(_743_v34).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_704_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(422))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_744_i8 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					})))).Cardinality()), (_index147).Int())
					var _745_v35 _dafny.Sequence
					_ = _745_v35
					_745_v35 = _dafny.SeqOf((_this).F6())
					var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_743_v34), 0))
					_ = _index148
					var _rhs170 _dafny.Set = (_742_v33).Difference(_dafny.SetOf((_741_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_741_v32), 0))).Int()).(bool), _698_v0))
					_ = _rhs170
					var _rhs171 _dafny.Int = Companion_Default___.SafeModuloInt((_701_v3).Plus(_dafny.IntOfUint32((_745_v35).Cardinality())), _701_v3)
					_ = _rhs171
					var _lhs118 _dafny.Array = _743_v34
					_ = _lhs118
					var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((_743_v34), 0))
					_ = _lhs119
					_742_v33 = _rhs170
					(_lhs118).ArraySet1(_rhs171, (_lhs119).Int())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		r0 = (_dafny.Zero).Minus(_701_v3)
		return r0
	}
}
func (_this *C2) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _746_v0 _dafny.Sequence
		_ = _746_v0
		_746_v0 = _dafny.UnicodeSeqOfUtf8Bytes("i")
		_746_v0 = _dafny.UnicodeSeqOfUtf8Bytes("e")
		var _747_v1 bool
		_ = _747_v1
		_747_v1 = false
		var _748_v2 _dafny.Set
		_ = _748_v2
		_748_v2 = _dafny.SetOf(p2, _747_v1)
		_747_v1 = (_748_v2).IsProperSubsetOf(_748_v2)
		var _749_v3 _dafny.Int
		_ = _749_v3
		_749_v3 = _dafny.IntOfInt64(153)
		var _750_v4 _dafny.Map
		_ = _750_v4
		_750_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('m'), _749_v3)
		var _751_v5 _dafny.Array
		_ = _751_v5
		var _nwElement0_21 _dafny.Int = _749_v3
		_ = _nwElement0_21
		var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(6))
		_ = _nw123
		(_nw123).ArraySet1(_nwElement0_21, 0)
		(_nw123).ArraySet1((_dafny.Zero).Minus(_749_v3), 1)
		(_nw123).ArraySet1(_749_v3, 2)
		(_nw123).ArraySet1(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_750_v4).Contains((_this).F5()) {
				return (_750_v4).Get((_this).F5()).(_dafny.Int)
			}
			return _749_v3
		})(), _749_v3), 3)
		(_nw123).ArraySet1((_dafny.Zero).Minus(_749_v3), 4)
		(_nw123).ArraySet1(_dafny.IntOfInt64(632), 5)
		_751_v5 = _nw123
		var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))
		_ = _index149
		(_751_v5).ArraySet1((_dafny.Zero).Minus(_749_v3), (_index149).Int())
		var _752_v6 _dafny.Set
		_ = _752_v6
		_752_v6 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-352))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_753_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})), _746_v0, _746_v0, _746_v0, _dafny.Companion_Sequence_.Concatenate(_746_v0, _746_v0))
		var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))
		_ = _index150
		var _rhs172 bool = (_747_v1) && (p0)
		_ = _rhs172
		var _rhs173 _dafny.Int = (_752_v6).Cardinality()
		_ = _rhs173
		var _lhs120 _dafny.Array = _751_v5
		_ = _lhs120
		var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))
		_ = _lhs121
		_747_v1 = _rhs172
		(_lhs120).ArraySet1(_rhs173, (_lhs121).Int())
		var _754_v7 _dafny.Array
		_ = _754_v7
		var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
		_ = _nw124
		_754_v7 = _nw124
		var _755_v8 _dafny.Array
		_ = _755_v8
		_755_v8 = _751_v5
		var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_754_v7), 0))
		_ = _index151
		(_754_v7).ArraySet1(_755_v8, (_index151).Int())
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_754_v7), 0))
		_ = _index152
		(_754_v7).ArraySet1(_751_v5, (_index152).Int())
		_749_v3 = _749_v3
		var _756_v9 _dafny.Sequence
		_ = _756_v9
		_756_v9 = _dafny.SeqOf(_747_v1)
		var _757_i1 _dafny.Int
		_ = _757_i1
		_757_i1 = _dafny.Zero
		{
			for ((_756_v9).Select((Companion_Default___.SafeIndex((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_756_v9).Cardinality()))).Uint32()).(bool)) == (p0) {
				{
					if (_757_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_757_i1 = (_757_i1).Plus(_dafny.One)
					var _758_v10 _dafny.Map
					_ = _758_v10
					_758_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(p2))
					var _759_v11 _dafny.Array
					_ = _759_v11
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_18
					var _nw125 _dafny.Array
					_ = _nw125
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw125 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) _dafny.MultiSet = (func(_760_v3 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
							return func(_761_i2 _dafny.Int) _dafny.MultiSet {
								return (_dafny.MultiSetOf(_760_v3, _dafny.IntOfInt64(4))).Union(_dafny.MultiSetOf(_760_v3))
							}
						})(_749_v3)
						_ = _init18
						var _element0_18 = _init18(_dafny.Zero)
						_ = _element0_18
						_nw125 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
						(_nw125).ArraySet1(_element0_18, 0)
						var _nativeLen0_18 = (_len0_18).Int()
						_ = _nativeLen0_18
						for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
							(_nw125).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
						}
					}
					_759_v11 = _nw125
					var _762_v12 _dafny.Sequence
					_ = _762_v12
					_762_v12 = _dafny.SeqOf(_749_v3)
					var _763_v13 _dafny.Map
					_ = _763_v13
					_763_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v3, _dafny.IntOfUint32((_746_v0).Cardinality()))
					var _764_v14 _dafny.Set
					_ = _764_v14
					_764_v14 = _dafny.SetOf((_763_v13).Cardinality())
					var _765_v15 _dafny.MultiSet
					_ = _765_v15
					_765_v15 = _dafny.MultiSetOf((_762_v12).Select((Companion_Default___.SafeIndex((_764_v14).Cardinality(), _dafny.IntOfUint32((_762_v12).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(_dafny.IntOfUint32((_762_v12).Cardinality())))
					var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_759_v11), 0))
					_ = _index153
					(_759_v11).ArraySet1(_765_v15, (_index153).Int())
					var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_759_v11), 0))
					_ = _index154
					var _rhs174 _dafny.Map = _758_v10
					_ = _rhs174
					var _rhs175 _dafny.Sequence = _746_v0
					_ = _rhs175
					var _rhs176 _dafny.Int = (_dafny.Zero).Minus(_749_v3)
					_ = _rhs176
					var _rhs177 _dafny.Int = Companion_Default___.SafeModuloInt((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int), _749_v3)
					_ = _rhs177
					var _rhs178 _dafny.MultiSet = (_765_v15).Update(_749_v3, Companion_Default___.Abs(_749_v3))
					_ = _rhs178
					var _lhs122 _dafny.Array = _759_v11
					_ = _lhs122
					var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(194), _dafny.ArrayLen((_759_v11), 0))
					_ = _lhs123
					_758_v10 = _rhs174
					_746_v0 = _rhs175
					_749_v3 = _rhs176
					_749_v3 = _rhs177
					(_lhs122).ArraySet1(_rhs178, (_lhs123).Int())
					(globalState).F1 = _749_v3
					var _766_v16 _dafny.Map
					_ = _766_v16
					_766_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_748_v2, (_this).F6())
					var _767_v17 _dafny.Map
					_ = _767_v17
					_767_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v3, p0)
					var _768_v18 _dafny.Map
					_ = _768_v18
					_768_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_762_v12, _dafny.IntOfInt64(316))
					var _769_v19 _dafny.Sequence
					_ = _769_v19
					_769_v19 = _dafny.SeqOf(_768_v18, (_768_v18).Update(_762_v12, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_747_v1, (_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int))).Cardinality()))
					var _770_v20 _dafny.Sequence
					_ = _770_v20
					_770_v20 = _dafny.SeqOf((_769_v19).Select((Companion_Default___.SafeIndex(_749_v3, _dafny.IntOfUint32((_769_v19).Cardinality()))).Uint32()).(_dafny.Map))
					var _771_v21 D0
					_ = _771_v21
					_771_v21 = Companion_D0_.Create_DC0_((_this).Fm10(_dafny.MultiSetFromSeq(_756_v9), _766_v16, (_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int), globalState))
					var _772_v22 _dafny.Array
					_ = _772_v22
					var _nwElement0_22 bool = (_this).Fm6(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(388))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg47 _dafny.Int) interface{} {
							return coer47(arg47)
						}
					}((func(_773_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_774_i3 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus((_773_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_773_v5), 0))).Int()).(_dafny.Int))
						}
					})(_751_v5))), globalState)
					_ = _nwElement0_22
					var _nw126 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
					_ = _nw126
					(_nw126).ArraySet1(_nwElement0_22, 0)
					(_nw126).ArraySet1(p0, 1)
					(_nw126).ArraySet1((_this).F6(), 2)
					(_nw126).ArraySet1(true, 3)
					(_nw126).ArraySet1(!((_756_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_746_v0).Cardinality()), _dafny.IntOfUint32((_756_v9).Cardinality()))).Uint32()).(bool)), 4)
					(_nw126).ArraySet1(((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(468)) == 0, 5)
					(_nw126).ArraySet1((_this).Fm10(_dafny.MultiSetOf(p2, _747_v1), _766_v16, _749_v3, globalState), 6)
					(_nw126).ArraySet1((_this).F6(), 7)
					(_nw126).ArraySet1(p0, 8)
					(_nw126).ArraySet1((_this).F6(), 9)
					(_nw126).ArraySet1(p2, 10)
					(_nw126).ArraySet1(!((func() bool {
						if (_767_v17).Contains((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int)) {
							return (_767_v17).Get((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int)).(bool)
						}
						return _747_v1
					})()) || (false), 11)
					(_nw126).ArraySet1(!(!(!(((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int)).Cmp((Companion_Default___.Fm28(_747_v1, p0, (_769_v19).Select((Companion_Default___.SafeIndex((_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_769_v19).Cardinality()))).Uint32()).(_dafny.Map), globalState)).Cardinality()) >= 0))), 12)
					(_nw126).ArraySet1((_771_v21).Dtor_cf0(), 13)
					(_nw126).ArraySet1(((_this).F6()) || (!((_this).F6())), 14)
					(_nw126).ArraySet1((_dafny.IntOfInt64(686)).Cmp(_dafny.IntOfInt64(53)) != 0, 15)
					_772_v22 = _nw126
					_772_v22 = _772_v22
					var _775_v23 D0
					_ = _775_v23
					_775_v23 = Companion_D0_.Create_DC0_(p2)
					var _776_v24 D0
					_ = _776_v24
					_776_v24 = Companion_D0_.Create_DC3_(_775_v23)
					var _777_v25 D2
					_ = _777_v25
					_777_v25 = Companion_D2_.Create_DC7_(_749_v3, (_751_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(784), _dafny.ArrayLen((_751_v5), 0))).Int()).(_dafny.Int))
					var _778_v26 _dafny.Map
					_ = _778_v26
					_778_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_776_v24, (_777_v25).Dtor_cf8())
					var _pat_let_tv10 = _775_v23
					_ = _pat_let_tv10
					_778_v26 = (_778_v26).Update(func(_pat_let6_0 D0) D0 {
						return func(_779_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let7_0 D0) D0 {
								return func(_780_dt__update_hcf3_h0 D0) D0 {
									return Companion_D0_.Create_DC3_(_780_dt__update_hcf3_h0)
								}(_pat_let7_0)
							}(_pat_let_tv10)
						}(_pat_let6_0)
					}(_776_v24), (_762_v12).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_749_v3)).Cardinality(), _dafny.IntOfUint32((_762_v12).Cardinality()))).Uint32()).(_dafny.Int))
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f13 bool
	_f14 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f13 = false
	_this._f14 = false
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

func (_this *C3) Ctor__(f13 bool, f14 bool) {
	{
		(_this)._f13 = f13
		(_this)._f14 = f14
	}
}
func (_this *C3) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(42), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_this).F14() {
				return _dafny.SeqOf(func() _dafny.Set {
					var _coll29 = _dafny.NewBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(403), _dafny.IntOfInt64(547))); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _781_v0 _dafny.Int
						_781_v0 = interface{}(_compr_29).(_dafny.Int)
						if ((_dafny.IntOfInt64(403)).Cmp(_781_v0) <= 0) && ((_781_v0).Cmp(_dafny.IntOfInt64(547)) < 0) {
							_coll29.Add((_781_v0).Plus(_dafny.IntOfInt64(-565)))
						}
					}
					return _coll29.ToSet()
				}(), _dafny.SetOf(_dafny.IntOfInt64(46), _dafny.IntOfInt64(880)), func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter31 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), (_this).F13())).Cardinality(), (_this).F14())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ueuscelno")).Cardinality()))).Elements()); ; {
						_compr_30, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _782_v1 _dafny.Int
						_782_v1 = interface{}(_compr_30).(_dafny.Int)
						if (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()), (_this).F13())).Cardinality(), (_this).F14())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ueuscelno")).Cardinality()))).Contains(_782_v1) {
							_coll30.Add((_782_v1).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg48 _dafny.Int) interface{} {
									return coer48(arg48)
								}
							}(func(_783_i0 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('y')
							}))).Cardinality())))
						}
					}
					return _coll30.ToSet()
				}())
			}
			return _dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetOf((_this).F13())).Cardinality()))
		})()).Cardinality())))
	}
}
func (_this *C3) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_((Companion_D2_.Create_DC7_(_dafny.IntOfInt64(135), _dafny.IntOfUint32((_dafny.SeqOf((_this).F13(), (_this).F13())).Cardinality()))).Dtor_cf7())
	}
}
func (_this *C3) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _784_v0 _dafny.Array
		_ = _784_v0
		var _len0_19 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_19
		var _nw127 _dafny.Array
		_ = _nw127
		if _len0_19.Cmp(_dafny.Zero) == 0 {
			_nw127 = _dafny.NewArray(_len0_19)
		} else {
			var _init19 func(_dafny.Int) _dafny.Int = (func(_785_p1 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
				return func(_786_i0 _dafny.Int) _dafny.Int {
					return (_786_i0).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-324)), (_785_p1).Cardinality())).Cardinality()))
				}
			})(p1)
			_ = _init19
			var _element0_19 = _init19(_dafny.Zero)
			_ = _element0_19
			_nw127 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
			(_nw127).ArraySet1(_element0_19, 0)
			var _nativeLen0_19 = (_len0_19).Int()
			_ = _nativeLen0_19
			for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
				(_nw127).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
			}
		}
		_784_v0 = _nw127
		var _787_v1 _dafny.Int
		_ = _787_v1
		_787_v1 = _dafny.IntOfInt64(-263)
		var _788_v2 _dafny.Sequence
		_ = _788_v2
		_788_v2 = _dafny.SeqOf(p0)
		var _789_v3 _dafny.Set
		_ = _789_v3
		_789_v3 = _dafny.SetOf(_dafny.IntOfInt64(114), _787_v1, _dafny.IntOfUint32((_788_v2).Cardinality()))
		var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))
		_ = _index155
		(_784_v0).ArraySet1((_789_v3).Cardinality(), (_index155).Int())
		var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))
		_ = _index156
		(_784_v0).ArraySet1(_787_v1, (_index156).Int())
		var _790_v4 bool
		_ = _790_v4
		_790_v4 = false
		_790_v4 = (_788_v2).Select((Companion_Default___.SafeIndex((_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_788_v2).Cardinality()))).Uint32()).(bool)
		var _791_v5 _dafny.Sequence
		_ = _791_v5
		_791_v5 = _dafny.SeqOf((_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int), _787_v1)
		var _792_v6 _dafny.Sequence
		_ = _792_v6
		_792_v6 = _dafny.UnicodeSeqOfUtf8Bytes("t")
		var _793_v7 _dafny.Sequence
		_ = _793_v7
		_793_v7 = _dafny.SeqOf(_792_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_794_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		})))
		var _795_v8 D2
		_ = _795_v8
		_795_v8 = Companion_D2_.Create_DC6_(Companion_Default___.Fm1(_dafny.IntOfInt64(897), true, (_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int), globalState))
		var _796_v9 D2
		_ = _796_v9
		_796_v9 = Companion_D2_.Create_DC9_(_795_v8)
		var _797_v10 D1
		_ = _797_v10
		_797_v10 = Companion_D1_.Create_DC4_(_dafny.IntOfUint32((_793_v7).Cardinality()))
		var _798_v11 _dafny.Array
		_ = _798_v11
		var _nw128 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw128
		_798_v11 = _nw128
		var _799_v12 _dafny.Array
		_ = _799_v12
		_799_v12 = _798_v11
		var _800_v13 _dafny.Map
		_ = _800_v13
		_800_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_790_v4, globalState), _799_v12)
		var _801_v14 _dafny.Sequence
		_ = _801_v14
		_801_v14 = _dafny.SeqOf(_800_v13)
		var _802_v15 _dafny.Map
		_ = _802_v15
		_802_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v1, (_801_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_792_v6).Cardinality()), _dafny.IntOfUint32((_801_v14).Cardinality()))).Uint32()).(_dafny.Map))
		var _803_v16 _dafny.Map
		_ = _803_v16
		_803_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v15, !(p2))
		var _rhs179 _dafny.Sequence = Companion_Default___.Fm37(!((_this).F13()), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_791_v5).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("g")).Cardinality())), _797_v10, _787_v1, globalState)
		_ = _rhs179
		var _rhs180 _dafny.Sequence = _793_v7
		_ = _rhs180
		var _rhs181 D2 = _796_v9
		_ = _rhs181
		var _rhs182 bool = (func() bool {
			if (_803_v16).Contains(_802_v15) {
				return (_803_v16).Get(_802_v15).(bool)
			}
			return (_this).F14()
		})()
		_ = _rhs182
		var _rhs183 _dafny.Int = _787_v1
		_ = _rhs183
		var _lhs124 *GlobalState = globalState
		_ = _lhs124
		_791_v5 = _rhs179
		_793_v7 = _rhs180
		_796_v9 = _rhs181
		_790_v4 = _rhs182
		_lhs124.F1 = _rhs183
		var _804_v17 _dafny.Map
		_ = _804_v17
		_804_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_787_v1, (_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int))
		var _805_v18 _dafny.Set
		_ = _805_v18
		_805_v18 = _dafny.SetOf(Companion_Default___.Fm0((_this).F14(), globalState), (_this).F13(), (_this).F14())
		(globalState).F1 = Companion_Default___.SafeModuloInt((_804_v17).Cardinality(), (_805_v18).Cardinality())
		if (_this).F13() {
			var _806_v19 D0
			_ = _806_v19
			_806_v19 = Companion_D0_.Create_DC0_(_790_v4)
			_790_v4 = (_806_v19).Dtor_cf0()
			var _807_v20 _dafny.CodePoint
			_ = _807_v20
			_807_v20 = _dafny.CodePoint('a')
			var _808_v21 _dafny.Map
			_ = _808_v21
			_808_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_798_v11, _807_v20)
			var _809_v22 D0
			_ = _809_v22
			_809_v22 = Companion_D0_.Create_DC2_(_808_v21, _807_v20)
			var _pat_let_tv11 = _798_v11
			_ = _pat_let_tv11
			var _pat_let_tv12 = _807_v20
			_ = _pat_let_tv12
			_807_v20 = (func(_pat_let8_0 D0) D0 {
				return func(_810_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let9_0 _dafny.Map) D0 {
						return func(_811_dt__update_hcf1_h0 _dafny.Map) D0 {
							return Companion_D0_.Create_DC2_(_811_dt__update_hcf1_h0, (_810_dt__update__tmp_h0).Dtor_cf2())
						}(_pat_let9_0)
					}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv11, _pat_let_tv12))
				}(_pat_let8_0)
			}(_809_v22)).Dtor_cf2()
			_790_v4 = (_dafny.IntOfInt64(58)).Cmp(_787_v1) < 0
			_790_v4 = (p2) == ((_dafny.IntOfInt64(243)).Cmp(_787_v1) == 0)
			var _812_v23 _dafny.Array
			_ = _812_v23
			var _nwElement0_23 _dafny.CodePoint = _807_v20
			_ = _nwElement0_23
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(23))
			_ = _nw129
			(_nw129).ArraySet1CodePoint(_nwElement0_23, 0)
			(_nw129).ArraySet1CodePoint(_807_v20, 1)
			(_nw129).ArraySet1CodePoint(_807_v20, 2)
			(_nw129).ArraySet1CodePoint(_807_v20, 3)
			(_nw129).ArraySet1CodePoint(_807_v20, 4)
			(_nw129).ArraySet1CodePoint(_807_v20, 5)
			(_nw129).ArraySet1CodePoint(_dafny.CodePoint('x'), 6)
			(_nw129).ArraySet1CodePoint(_807_v20, 7)
			(_nw129).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _807_v20
				}
				return _807_v20
			})(), 8)
			(_nw129).ArraySet1CodePoint(_807_v20, 9)
			(_nw129).ArraySet1CodePoint(_807_v20, 10)
			(_nw129).ArraySet1CodePoint(_807_v20, 11)
			(_nw129).ArraySet1CodePoint(_807_v20, 12)
			(_nw129).ArraySet1CodePoint(_807_v20, 13)
			(_nw129).ArraySet1CodePoint(_807_v20, 14)
			(_nw129).ArraySet1CodePoint(_807_v20, 15)
			(_nw129).ArraySet1CodePoint(_807_v20, 16)
			(_nw129).ArraySet1CodePoint(_807_v20, 17)
			(_nw129).ArraySet1CodePoint(_807_v20, 18)
			(_nw129).ArraySet1CodePoint(_807_v20, 19)
			(_nw129).ArraySet1CodePoint(_807_v20, 20)
			(_nw129).ArraySet1CodePoint(_807_v20, 21)
			(_nw129).ArraySet1CodePoint((_792_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_792_v6).Cardinality()), _dafny.IntOfUint32((_792_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), 22)
			_812_v23 = _nw129
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_812_v23), 0))
			_ = _index157
			(_812_v23).ArraySet1CodePoint(_dafny.CodePoint('d'), (_index157).Int())
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_812_v23), 0))
			_ = _index158
			(_812_v23).ArraySet1CodePoint(_807_v20, (_index158).Int())
		} else {
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))
			_ = _index159
			(_784_v0).ArraySet1((_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int), (_index159).Int())
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))
			_ = _index160
			var _rhs184 bool = !((_788_v2).Select((Companion_Default___.SafeIndex(_787_v1, _dafny.IntOfUint32((_788_v2).Cardinality()))).Uint32()).(bool))
			_ = _rhs184
			var _rhs185 _dafny.Int = (_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs185
			var _rhs186 _dafny.Sequence = _788_v2
			_ = _rhs186
			var _lhs125 _dafny.Array = _784_v0
			_ = _lhs125
			var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))
			_ = _lhs126
			_790_v4 = _rhs184
			(_lhs125).ArraySet1(_rhs185, (_lhs126).Int())
			_788_v2 = _rhs186
			var _813_v24 _dafny.MultiSet
			_ = _813_v24
			_813_v24 = _dafny.MultiSetOf((_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int))
			(globalState).F1 = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(872), (_this).Fm4(_788_v2, _787_v1, _dafny.IntOfInt64(176), _813_v24, globalState)), (_dafny.Zero).Minus((_dafny.IntOfInt64(781)).Minus(_787_v1)))
			var _814_v25 _dafny.CodePoint
			_ = _814_v25
			_814_v25 = _dafny.CodePoint('p')
			var _815_v26 D0
			_ = _815_v26
			var _816_v27 bool
			_ = _816_v27
			var _out28 D0
			_ = _out28
			var _out29 bool
			_ = _out29
			_out28, _out29 = Companion_Default___.M0(_798_v11, _814_v25, p0, _814_v25, globalState)
			_815_v26 = _out28
			_816_v27 = _out29
			_814_v25 = _dafny.CodePoint('x')
		}
		(globalState).F1 = (_dafny.Zero).Minus((_784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(248), _dafny.ArrayLen((_784_v0), 0))).Int()).(_dafny.Int))
	}
}
func (_this *C3) F13() bool {
	{
		return _this._f13
	}
}
func (_this *C3) F14() bool {
	{
		return _this._f14
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	dummy byte
}

func New_C4_() *C4 {
	_this := C4{}

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

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		if (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).IsSubsetOf(_dafny.MultiSetOf(true)) {
			return (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lsyykqbc")).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ol")).Cardinality())))
		} else {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.IntOfInt64(338)), false)).Cardinality()
		}
	}
}
func (_this *C4) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		if true {
			return Companion_D1_.Create_DC4_((func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter32 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qfa")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))).Elements()); ; {
					_compr_31, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _817_v0 _dafny.Int
					_817_v0 = interface{}(_compr_31).(_dafny.Int)
					if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qfa")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))).Contains(_817_v0) {
						_coll31.Add(Companion_Default___.SafeDivisionInt(_817_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.SeqOf(_dafny.IntOfInt64(-454)))).Cardinality()), _dafny.IntOfInt64(-177))
					}
				}
				return _coll31.ToMap()
			}()).Cardinality())
		} else {
			return Companion_D1_.Create_DC4_(_dafny.IntOfInt64(368))
		}
	}
}
func (_this *C4) Fm30(p0 _dafny.Map, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32(((func() _dafny.Sequence {
			if false {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-327))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}(func(_818_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('x')
				}))
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("wx")
		})()).Cardinality())
	}
}
func (_this *C4) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _819_v0 bool
		_ = _819_v0
		_819_v0 = true
		_819_v0 = false
		var _820_v1 *C0
		_ = _820_v1
		var _nw130 *C0 = New_C0_()
		_ = _nw130
		_nw130.Ctor__(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-580))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}(func(_821_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf(_dafny.IntOfInt64(822), (_dafny.MultiSetOf(_dafny.IntOfInt64(645), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-537)))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg52 _dafny.Int) interface{} {
					return coer52(arg52)
				}
			}(func(_822_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality()), _dafny.IntOfInt64(830), _dafny.IntOfInt64(979))
		})))
		_820_v1 = _nw130
		_819_v0 = _819_v0
		var _823_v2 _dafny.Sequence
		_ = _823_v2
		_823_v2 = _dafny.SeqOf(_819_v0)
		var _824_i2 _dafny.Int
		_ = _824_i2
		_824_i2 = _dafny.Zero
		{
			for !((_823_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-736), _dafny.IntOfUint32((_823_v2).Cardinality()))).Uint32()).(bool)) {
				{
					if (_824_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_824_i2 = (_824_i2).Plus(_dafny.One)
					var _825_v3 _dafny.Sequence
					_ = _825_v3
					_825_v3 = _dafny.UnicodeSeqOfUtf8Bytes("iygwbkys")
					var _826_v4 _dafny.Sequence
					_ = _826_v4
					_826_v4 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xkbpmdi"))
					_825_v3 = _dafny.Companion_Sequence_.Concatenate((_826_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-365), _dafny.IntOfUint32((_826_v4).Cardinality()))).Uint32()).(_dafny.Sequence), _825_v3)
					var _827_v5 _dafny.Int
					_ = _827_v5
					_827_v5 = _dafny.IntOfInt64(723)
					var _828_v6 D2
					_ = _828_v6
					_828_v6 = Companion_D2_.Create_DC6_(_825_v3)
					var _829_v7 _dafny.Set
					_ = _829_v7
					_829_v7 = _dafny.SetOf((_820_v1).F11())
					var _830_v8 _dafny.Map
					_ = _830_v8
					_830_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _829_v7)
					_819_v0 = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm31(p0, _827_v5, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_828_v6).Dtor_cf6()).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm31(p0, _827_v5, globalState)).Cardinality()))).Uint32(), _dafny.IntOfInt64(888)), (Companion_Default___.SafeIndex((_830_v8).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm31(p0, _827_v5, globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_828_v6).Dtor_cf6()).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm31(p0, _827_v5, globalState)).Cardinality()))).Uint32(), _dafny.IntOfInt64(888))).Cardinality()))).Uint32(), _dafny.IntOfUint32((_825_v3).Cardinality())))).IsDisjointFrom(Companion_Default___.Fm32(globalState))
					_827_v5 = _827_v5
					(globalState).F1 = _827_v5
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _831_v9 _dafny.Int
		_ = _831_v9
		_831_v9 = _dafny.IntOfInt64(140)
		var _832_v10 _dafny.Map
		_ = _832_v10
		_832_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_831_v9, p2)
		var _833_v11 _dafny.Map
		_ = _833_v11
		_833_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_831_v9).Times(_831_v9), (Companion_Default___.Fm33(globalState)).Merge(_832_v10))
		_832_v10 = (func() _dafny.Map {
			if (_833_v11).Contains(_831_v9) {
				return (_833_v11).Get(_831_v9).(_dafny.Map)
			}
			return _832_v10
		})()
		var _834_i3 _dafny.Int
		_ = _834_i3
		_834_i3 = _dafny.Zero
		{
			for (func() bool {
				if !(_819_v0) || (false) {
					return Companion_Default___.Fm0(_819_v0, globalState)
				}
				return _819_v0
			})() {
				{
					if (_834_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_834_i3 = (_834_i3).Plus(_dafny.One)
					var _835_v12 *C1
					_ = _835_v12
					var _nw131 *C1 = New_C1_()
					_ = _nw131
					_nw131.Ctor__()
					_835_v12 = _nw131
					var _836_v13 _dafny.Array
					_ = _836_v13
					var _nw132 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(18))
					_ = _nw132
					_836_v13 = _nw132
					var _837_v14 _dafny.Array
					_ = _837_v14
					_837_v14 = _836_v13
					var _838_v15 _dafny.Map
					_ = _838_v15
					_838_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _836_v13)
					var _839_v16 _dafny.Array
					_ = _839_v16
					var _nwElement0_24 _dafny.Array = _836_v13
					_ = _nwElement0_24
					var _nw133 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(13))
					_ = _nw133
					(_nw133).ArraySet1(_nwElement0_24, 0)
					(_nw133).ArraySet1(_836_v13, 1)
					(_nw133).ArraySet1(_836_v13, 2)
					(_nw133).ArraySet1(_836_v13, 3)
					(_nw133).ArraySet1(_836_v13, 4)
					(_nw133).ArraySet1(_836_v13, 5)
					(_nw133).ArraySet1((_837_v14), 6)
					(_nw133).ArraySet1(_836_v13, 7)
					(_nw133).ArraySet1(_836_v13, 8)
					(_nw133).ArraySet1((func() _dafny.Array {
						if (_838_v15).Contains((_820_v1).F11()) {
							return (_838_v15).Get((_820_v1).F11()).(_dafny.Array)
						}
						return _836_v13
					})(), 9)
					(_nw133).ArraySet1(_836_v13, 10)
					(_nw133).ArraySet1(_836_v13, 11)
					(_nw133).ArraySet1(_836_v13, 12)
					_839_v16 = _nw133
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_839_v16), 0))
					_ = _index161
					(_839_v16).ArraySet1(_836_v13, (_index161).Int())
					var _840_v17 _dafny.Sequence
					_ = _840_v17
					_840_v17 = _dafny.SeqOf(_836_v13)
					var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(610), _dafny.ArrayLen((_839_v16), 0))
					_ = _index162
					(_839_v16).ArraySet1((_840_v17).Select((Companion_Default___.SafeIndex(_831_v9, _dafny.IntOfUint32((_840_v17).Cardinality()))).Uint32()).(_dafny.Array), (_index162).Int())
					var _841_v18 D2
					_ = _841_v18
					_841_v18 = Companion_D2_.Create_DC8_(true)
					_819_v0 = (func() bool {
						if _819_v0 {
							return (_841_v18).Dtor_cf9()
						}
						return (_820_v1).F11()
					})()
					if (_831_v9).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(p0, false, _819_v0)).Cardinality()), _831_v9)).Cardinality()) == 0 {
						_819_v0 = (_823_v2).Select((Companion_Default___.SafeIndex((_831_v9).Plus(_831_v9), _dafny.IntOfUint32((_823_v2).Cardinality()))).Uint32()).(bool)
						var _842_v19 D2
						_ = _842_v19
						_842_v19 = Companion_D2_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("s"))
						_842_v19 = _842_v19
						var _843_v20 _dafny.Array
						_ = _843_v20
						var _nw134 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
						_ = _nw134
						_843_v20 = _nw134
						var _844_v21 _dafny.Map
						_ = _844_v21
						_844_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_v13, ((_820_v1).Fm5(_841_v18, globalState)).Dtor_cf4())
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_843_v20), 0))
						_ = _index163
						(_843_v20).ArraySet1(_844_v21, (_index163).Int())
						var _845_v22 _dafny.Sequence
						_ = _845_v22
						_845_v22 = _dafny.SeqOf(_844_v21, _844_v21, _844_v21, _844_v21, _844_v21)
						var _846_v23 _dafny.Map
						_ = _846_v23
						_846_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_820_v1).F11(), Companion_Default___.Fm34((_820_v1).F11(), globalState))
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_843_v20), 0))
						_ = _index164
						var _rhs187 _dafny.Map = (_845_v22).Select((Companion_Default___.SafeIndex(_831_v9, _dafny.IntOfUint32((_845_v22).Cardinality()))).Uint32()).(_dafny.Map)
						_ = _rhs187
						var _rhs188 bool = (((func() _dafny.MultiSet {
							if (_846_v23).Contains(p0) {
								return (_846_v23).Get(p0).(_dafny.MultiSet)
							}
							return p1
						})()).Equals(_dafny.MultiSetOf(p0))) && ((_820_v1).F11())
						_ = _rhs188
						var _rhs189 bool = (_820_v1).F11()
						_ = _rhs189
						var _lhs127 _dafny.Array = _843_v20
						_ = _lhs127
						var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_843_v20), 0))
						_ = _lhs128
						(_lhs127).ArraySet1(_rhs187, (_lhs128).Int())
						_819_v0 = _rhs188
						_819_v0 = _rhs189
						_819_v0 = true
						var _847_v24 _dafny.MultiSet
						_ = _847_v24
						_847_v24 = _dafny.MultiSetOf(_831_v9)
						var _848_v25 _dafny.CodePoint
						_ = _848_v25
						_848_v25 = _dafny.CodePoint('d')
						var _849_v26 _dafny.Map
						_ = _849_v26
						_849_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_848_v25, _831_v9)
						var _850_v27 _dafny.Map
						_ = _850_v27
						_850_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_847_v24).Cardinality()), (_849_v26).Cardinality())
						var _851_v28 _dafny.MultiSet
						_ = _851_v28
						_851_v28 = _dafny.MultiSetOf((_850_v27).Cardinality(), _831_v9, _831_v9, (_dafny.Zero).Minus(_831_v9), _dafny.IntOfInt64(741))
						var _852_v29 _dafny.Map
						_ = _852_v29
						_852_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_820_v1, _851_v28)
						var _853_v30 _dafny.Map
						_ = _853_v30
						_853_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_831_v9, (_this).Fm4(_823_v2, _dafny.IntOfInt64(574), _831_v9, (func() _dafny.MultiSet {
							if (_852_v29).Contains(_820_v1) {
								return (_852_v29).Get(_820_v1).(_dafny.MultiSet)
							}
							return _851_v28
						})(), globalState))
						var _854_v31 _dafny.Set
						_ = _854_v31
						_854_v31 = _dafny.SetOf(p0)
						_853_v30 = (_853_v30).Update(_831_v9, _dafny.IntOfUint32((Companion_Default___.Fm31(_819_v0, (_854_v31).Cardinality(), globalState)).Cardinality()))
					} else {
						var _855_v32 *C1
						_ = _855_v32
						var _nw135 *C1 = New_C1_()
						_ = _nw135
						_nw135.Ctor__()
						_855_v32 = _nw135
						_835_v12 = _835_v12
						(globalState).F1 = (_dafny.Zero).Minus(_831_v9)
						var _856_v34 _dafny.Map
						_ = _856_v34
						_856_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_820_v1).F11(), ((func() _dafny.Map {
							var _coll32 = _dafny.NewMapBuilder()
							_ = _coll32
							for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(599), _dafny.IntOfInt64(961))); ; {
								_compr_32, _ok33 := _iter33()
								if !_ok33 {
									break
								}
								var _857_v33 _dafny.Int
								_857_v33 = interface{}(_compr_32).(_dafny.Int)
								if ((_dafny.IntOfInt64(599)).Cmp(_857_v33) <= 0) && ((_857_v33).Cmp(_dafny.IntOfInt64(961)) < 0) {
									_coll32.Add((_857_v33).Plus(_831_v9), _dafny.CodePoint('q'))
								}
							}
							return _coll32.ToMap()
						}()).Cardinality()).Plus(_831_v9))
						_856_v34 = (_856_v34).Update(!(_819_v0) || ((_820_v1).F11()), _831_v9)
						_819_v0 = (Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_831_v9), (_dafny.Zero).Minus(_831_v9))).Cmp(_831_v9) == 0
					}
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
	}
}
func (_this *C4) M10(p0 _dafny.Array, globalState *GlobalState) (*C0, _dafny.Array, _dafny.Array) {
	{
		var r0 *C0 = (*C0)(nil)
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		if false {
			var _858_v0 bool
			_ = _858_v0
			_858_v0 = false
			var _859_v1 _dafny.Int
			_ = _859_v1
			_859_v1 = _dafny.IntOfInt64(479)
			var _860_v2 _dafny.Map
			_ = _860_v2
			_860_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_858_v0, _859_v1)
			var _861_v3 _dafny.Sequence
			_ = _861_v3
			_861_v3 = _dafny.SeqOf(_dafny.SeqOf(_859_v1), Companion_Default___.Fm31(_858_v0, _859_v1, globalState), _dafny.SeqOf((_860_v2).Cardinality(), _859_v1, _859_v1))
			var _862_v4 *C0
			_ = _862_v4
			var _nw136 *C0 = New_C0_()
			_ = _nw136
			_nw136.Ctor__(!(_858_v0) || (_858_v0), _861_v3)
			_862_v4 = _nw136
			var _863_v5 _dafny.CodePoint
			_ = _863_v5
			_863_v5 = _dafny.CodePoint('d')
			_863_v5 = _863_v5
			var _864_v6 _dafny.Sequence
			_ = _864_v6
			_864_v6 = _dafny.SeqOf((_862_v4).F11(), _858_v0, false)
			var _865_v7 D1
			_ = _865_v7
			_865_v7 = Companion_D1_.Create_DC4_(_859_v1)
			var _866_v8 _dafny.MultiSet
			_ = _866_v8
			_866_v8 = _dafny.MultiSetOf((_dafny.Zero).Minus(_859_v1))
			(globalState).F1 = (_862_v4).Fm4(_864_v6, _859_v1, (_865_v7).Dtor_cf4(), _866_v8, globalState)
			(globalState).F1 = ((func() _dafny.Int {
				if _858_v0 {
					return _859_v1
				}
				return _859_v1
			})()).Plus(_dafny.IntOfInt64(460))
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((p0), 0))
			_ = _index165
			(p0).ArraySet1((_862_v4).F11(), (_index165).Int())
			var _867_v9 _dafny.MultiSet
			_ = _867_v9
			_867_v9 = _dafny.MultiSetOf(_858_v0)
			var _868_v10 _dafny.Map
			_ = _868_v10
			_868_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_867_v9).Cardinality(), _858_v0)
			var _869_v11 _dafny.Map
			_ = _869_v11
			_869_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_862_v4, (func() bool {
				if (_868_v10).Contains((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_859_v1, (_862_v4).F11())).Cardinality())) {
					return (_868_v10).Get((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_859_v1, (_862_v4).F11())).Cardinality())).(bool)
				}
				return _858_v0
			})())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((p0), 0))
			_ = _index166
			(p0).ArraySet1((_869_v11).Contains(_862_v4), (_index166).Int())
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((p0), 0))
			_ = _index167
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((p0), 0))
			_ = _index168
			var _rhs190 bool = Companion_Default___.Fm0((_862_v4).F11(), globalState)
			_ = _rhs190
			var _rhs191 bool = _858_v0
			_ = _rhs191
			var _rhs192 bool = (func() bool {
				if _858_v0 {
					return (func() bool {
						if _858_v0 {
							return _858_v0
						}
						return (_862_v4).F11()
					})()
				}
				return _858_v0
			})()
			_ = _rhs192
			var _lhs129 _dafny.Array = p0
			_ = _lhs129
			var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((p0), 0))
			_ = _lhs130
			var _lhs131 _dafny.Array = p0
			_ = _lhs131
			var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((p0), 0))
			_ = _lhs132
			(_lhs129).ArraySet1(_rhs190, (_lhs130).Int())
			(_lhs131).ArraySet1(_rhs191, (_lhs132).Int())
			_858_v0 = _rhs192
		} else {
			var _870_v12 bool
			_ = _870_v12
			_870_v12 = true
			_870_v12 = Companion_Default___.Fm0(_870_v12, globalState)
			var _871_v13 _dafny.Sequence
			_ = _871_v13
			_871_v13 = _dafny.SeqOf(p0, (func() _dafny.Array {
				if _870_v12 {
					return p0
				}
				return p0
			})())
			var _872_v14 _dafny.Int
			_ = _872_v14
			_872_v14 = _dafny.IntOfInt64(497)
			_871_v13 = _dafny.Companion_Sequence_.Update(_871_v13, (Companion_Default___.SafeIndex(_872_v14, _dafny.IntOfUint32((_871_v13).Cardinality()))).Uint32(), p0)
			var _873_v15 _dafny.Sequence
			_ = _873_v15
			_873_v15 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_870_v12)).Cardinality()))))
			_873_v15 = _dafny.Companion_Sequence_.Concatenate(_873_v15, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-694))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_874_i0 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(389)
			})))
			var _875_v16 _dafny.Sequence
			_ = _875_v16
			_875_v16 = _dafny.SeqOf(_873_v15, _873_v15)
			var _876_v17 _dafny.Sequence
			_ = _876_v17
			_876_v17 = _dafny.SeqOf(_dafny.SeqOf(_873_v15), _875_v16)
			var _877_v18 *C0
			_ = _877_v18
			var _nw137 *C0 = New_C0_()
			_ = _nw137
			_nw137.Ctor__((func() bool {
				if false {
					return Companion_Default___.Fm0(_870_v12, globalState)
				}
				return !(_870_v12)
			})(), _dafny.Companion_Sequence_.Concatenate(_875_v16, (_876_v17).Select((Companion_Default___.SafeIndex(_872_v14, _dafny.IntOfUint32((_876_v17).Cardinality()))).Uint32()).(_dafny.Sequence)))
			_877_v18 = _nw137
			(globalState).F1 = _872_v14
		}
		var _878_v19 _dafny.Int
		_ = _878_v19
		_878_v19 = _dafny.IntOfInt64(470)
		var _879_v20 _dafny.Sequence
		_ = _879_v20
		_879_v20 = _dafny.SeqOf(_878_v19)
		var _hi4 _dafny.Int = (_879_v20).Select((Companion_Default___.SafeIndex(_878_v19, _dafny.IntOfUint32((_879_v20).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi4
		for _880_i1 := Companion_Default___.SafeModuloInt(_878_v19, _878_v19); _880_i1.Cmp(_hi4) < 0; _880_i1 = _880_i1.Plus(_dafny.One) {
			var _881_v21 bool
			_ = _881_v21
			_881_v21 = false
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
			_ = _index169
			(p0).ArraySet1(_881_v21, (_index169).Int())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
			_ = _index170
			(p0).ArraySet1(_881_v21, (_index170).Int())
			if (_880_i1).Cmp(_880_i1) <= 0 {
				var _882_v22 _dafny.Sequence
				_ = _882_v22
				_882_v22 = _dafny.SeqOf(_879_v20)
				var _883_v23 *C0
				_ = _883_v23
				var _nw138 *C0 = New_C0_()
				_ = _nw138
				_nw138.Ctor__(Companion_Default___.Fm0((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), (func() _dafny.Sequence {
					if (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool) {
						return _882_v22
					}
					return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg54 _dafny.Int) interface{} {
							return coer54(arg54)
						}
					}((func(_884_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_885_i2 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus(_884_i1)
						}
					})(_880_i1))), _879_v20)
				})())
				_883_v23 = _nw138
				var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _index171
				(p0).ArraySet1((_883_v23).F11(), (_index171).Int())
				var _886_v24 _dafny.Sequence
				_ = _886_v24
				_886_v24 = _dafny.SeqOf(!(_881_v21), Companion_Default___.Fm0(!((_883_v23).F11()), globalState))
				var _887_v25 _dafny.Set
				_ = _887_v25
				_887_v25 = _dafny.SetOf(_dafny.IntOfUint32((_886_v24).Cardinality()), _878_v19, _878_v19)
				var _888_v27 _dafny.Map
				_ = _888_v27
				_888_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(_880_i1, _878_v19, _880_i1, _880_i1, _878_v19))
				var _889_v28 _dafny.MultiSet
				_ = _889_v28
				_889_v28 = _dafny.MultiSetOf(_880_i1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tkfxh")).Cardinality()))
				var _890_v29 _dafny.Map
				_ = _890_v29
				_890_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_883_v23).F11(), (_883_v23).F11())
				var _891_v30 _dafny.Map
				_ = _891_v30
				_891_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-696), _878_v19)
				var _892_v31 _dafny.Array
				_ = _892_v31
				var _nwElement0_25 _dafny.Set = _887_v25
				_ = _nwElement0_25
				var _nw139 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(11))
				_ = _nw139
				(_nw139).ArraySet1(_nwElement0_25, 0)
				(_nw139).ArraySet1(_887_v25, 1)
				(_nw139).ArraySet1((_887_v25).Union(_887_v25), 2)
				(_nw139).ArraySet1((func() _dafny.Set {
					if (_883_v23).F11() {
						return _887_v25
					}
					return _887_v25
				})(), 3)
				(_nw139).ArraySet1((_887_v25).Union(_887_v25), 4)
				(_nw139).ArraySet1(_887_v25, 5)
				(_nw139).ArraySet1(func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(347), _dafny.IntOfInt64(914))); ; {
						_compr_33, _ok34 := _iter34()
						if !_ok34 {
							break
						}
						var _893_v26 _dafny.Int
						_893_v26 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(347)).Cmp(_893_v26) <= 0) && ((_893_v26).Cmp(_dafny.IntOfInt64(914)) < 0) {
							_coll33.Add(Companion_Default___.SafeModuloInt(_893_v26, _dafny.IntOfInt64(690)))
						}
					}
					return _coll33.ToSet()
				}(), 6)
				(_nw139).ArraySet1(_887_v25, 7)
				(_nw139).ArraySet1(_887_v25, 8)
				(_nw139).ArraySet1((func() _dafny.Set {
					if (_888_v27).Contains((_883_v23).F11()) {
						return (_888_v27).Get((_883_v23).F11()).(_dafny.Set)
					}
					return _887_v25
				})(), 9)
				(_nw139).ArraySet1(Companion_Default___.Fm35(Companion_Default___.Fm1(_880_i1, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), _880_i1, globalState), _889_v28, Companion_Default___.Fm36(_880_i1, (_890_v29).Cardinality(), _880_i1, _891_v30, globalState), globalState), 10)
				_892_v31 = _nw139
				_892_v31 = _892_v31
				_879_v20 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(784))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg55 _dafny.Int) interface{} {
						return coer55(arg55)
					}
				}((func(_894_i1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_895_i3 _dafny.Int) _dafny.Int {
						return _894_i1
					}
				})(_880_i1))), _879_v20)
				var _896_v32 _dafny.Map
				_ = _896_v32
				_896_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_886_v24, _886_v24))
				var _897_v33 _dafny.Sequence
				_ = _897_v33
				_897_v33 = _dafny.UnicodeSeqOfUtf8Bytes("iwoqi")
				var _898_v34 _dafny.CodePoint
				_ = _898_v34
				_898_v34 = _dafny.CodePoint('a')
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _index172
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _index173
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _index174
				var _rhs193 bool = !((_883_v23).F11())
				_ = _rhs193
				var _rhs194 bool = Companion_Default___.Fm0((_883_v23).F11(), globalState)
				_ = _rhs194
				var _rhs195 _dafny.Int = _878_v19
				_ = _rhs195
				var _rhs196 _dafny.Sequence = (func() _dafny.Sequence {
					if (_896_v32).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool)) {
						return (_896_v32).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool)).(_dafny.Sequence)
					}
					return _886_v24
				})()
				_ = _rhs196
				var _rhs197 bool = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_897_v33, _897_v33), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(211))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_899_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				})), (Companion_Default___.SafeIndex(_880_i1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(211))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}(func(_900_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))).Cardinality()))).Uint32(), _898_v34))
				_ = _rhs197
				var _lhs133 _dafny.Array = p0
				_ = _lhs133
				var _lhs134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _lhs134
				var _lhs135 _dafny.Array = p0
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _lhs136
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				var _lhs138 _dafny.Array = p0
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _lhs139
				(_lhs133).ArraySet1(_rhs193, (_lhs134).Int())
				(_lhs135).ArraySet1(_rhs194, (_lhs136).Int())
				_lhs137.F1 = _rhs195
				_886_v24 = _rhs196
				(_lhs138).ArraySet1(_rhs197, (_lhs139).Int())
			} else {
				var _901_v35 _dafny.Array
				_ = _901_v35
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw140
				_901_v35 = _nw140
				_901_v35 = _901_v35
				var _902_v36 T1
				_ = _902_v36
				var _nw141 *C3 = New_C3_()
				_ = _nw141
				_nw141.Ctor__(true, _881_v21)
				_902_v36 = _nw141
				var _903_v37 _dafny.Map
				_ = _903_v37
				_903_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm3(_881_v21, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), globalState), _878_v19)
				var _904_v38 _dafny.Map
				_ = _904_v38
				_904_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v37, _881_v21)
				var _905_v39 _dafny.Sequence
				_ = _905_v39
				_905_v39 = _dafny.SeqOf(_881_v21, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool))
				var _906_v40 _dafny.MultiSet
				_ = _906_v40
				_906_v40 = _dafny.MultiSetOf(_905_v39, _905_v39, _905_v39)
				var _907_v41 _dafny.Map
				_ = _907_v41
				_907_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm30(_904_v38, _906_v40, globalState), _902_v36)
				_902_v36 = (func() T1 {
					if (_907_v41).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg58 _dafny.Int) interface{} {
							return coer58(arg58)
						}
					}(func(_908_i5 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('u')
					}))).Cardinality()))) {
						return (_907_v41).Get((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(318))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg59 _dafny.Int) interface{} {
								return coer59(arg59)
							}
						}(func(_909_i5 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}))).Cardinality()))).(T1)
					}
					return _902_v36
				})()
				var _910_v42 _dafny.Array
				_ = _910_v42
				var _nw142 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw142
				_910_v42 = _nw142
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_910_v42), 0))
				_ = _index175
				(_910_v42).ArraySet1(_878_v19, (_index175).Int())
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.ArrayLen((_910_v42), 0))
				_ = _index176
				(_910_v42).ArraySet1(_878_v19, (_index176).Int())
				var _911_v43 *C3
				_ = _911_v43
				var _nw143 *C3 = New_C3_()
				_ = _nw143
				_nw143.Ctor__(false, true)
				_911_v43 = _nw143
				var _912_v44 *C2
				_ = _912_v44
				var _nw144 *C2 = New_C2_()
				_ = _nw144
				_nw144.Ctor__(_dafny.CodePoint('m'), _881_v21)
				_912_v44 = _nw144
			}
			if !(Companion_Default___.Fm0(true, globalState)) {
				var _913_v45 *C1
				_ = _913_v45
				var _nw145 *C1 = New_C1_()
				_ = _nw145
				_nw145.Ctor__()
				_913_v45 = _nw145
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_20
				var _nw146 _dafny.Array
				_ = _nw146
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw146 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = (func(_914_v21 bool, _915_p0 _dafny.Array, _916_i1 _dafny.Int, _917_v19 _dafny.Int) func(_dafny.Int) bool {
						return func(_918_i6 _dafny.Int) bool {
							return !(!(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_914_v21, (_915_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_915_p0), 0))).Int()).(bool))).Cardinality()), _916_i1, _dafny.IntOfInt64(266), _917_v19)).Equals(_dafny.MultiSetOf(_916_i1)))
						}
					})(_881_v21, p0, _880_i1, _878_v19)
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
				r2 = _nw146
				var _919_v46 _dafny.Sequence
				_ = _919_v46
				_919_v46 = _dafny.UnicodeSeqOfUtf8Bytes("xr")
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))
				_ = _index177
				(p0).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_919_v46, _919_v46), (_index177).Int())
				var _920_v47 D0
				_ = _920_v47
				_920_v47 = Companion_D0_.Create_DC1_()
				_920_v47 = _920_v47
				var _921_v48 _dafny.Array
				_ = _921_v48
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_21
				var _nw147 _dafny.Array
				_ = _nw147
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw147 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) bool = (func(_922_p0 _dafny.Array) func(_dafny.Int) bool {
						return func(_923_i7 _dafny.Int) bool {
							return (_922_p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((_922_p0), 0))).Int()).(bool)
						}
					})(p0)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw147 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw147).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw147).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_921_v48 = _nw147
				var _924_v49 _dafny.Sequence
				_ = _924_v49
				_924_v49 = _dafny.SeqOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), false)
				var _925_v50 _dafny.Sequence
				_ = _925_v50
				_925_v50 = _dafny.SeqOf(true, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), !((_924_v49).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_878_v19), _dafny.IntOfUint32((_924_v49).Cardinality()))).Uint32()).(bool)), !(false))
				var _926_v51 _dafny.Map
				_ = _926_v51
				_926_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_921_v48, (_924_v49).Select((Companion_Default___.SafeIndex(_880_i1, _dafny.IntOfUint32((_924_v49).Cardinality()))).Uint32()).(bool))
				_926_v51 = (_926_v51).Update(_921_v48, _881_v21)
			} else {
				var _927_v52 _dafny.CodePoint
				_ = _927_v52
				_927_v52 = _dafny.CodePoint('x')
				var _928_v53 _dafny.MultiSet
				_ = _928_v53
				_928_v53 = _dafny.MultiSetOf((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool))
				var _929_v54 *C2
				_ = _929_v54
				var _nw148 *C2 = New_C2_()
				_ = _nw148
				_nw148.Ctor__(_927_v52, (_928_v53).IsSubsetOf(_928_v53))
				_929_v54 = _nw148
				(globalState).F1 = _880_i1
				var _930_v55 *C2
				_ = _930_v55
				var _nw149 *C2 = New_C2_()
				_ = _nw149
				_nw149.Ctor__(_927_v52, (_878_v19).Cmp(_878_v19) < 0)
				_930_v55 = _nw149
				var _931_v56 D2
				_ = _931_v56
				_931_v56 = Companion_D2_.Create_DC7_(Companion_Default___.SafeDivisionInt(_880_i1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_879_v20, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.IntOfUint32((_879_v20).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_880_i1))).Cardinality())), _880_i1)
				var _932_v57 _dafny.Sequence
				_ = _932_v57
				_932_v57 = _dafny.UnicodeSeqOfUtf8Bytes("tecdwogxu")
				var _rhs198 _dafny.Int = _878_v19
				_ = _rhs198
				var _rhs199 D2 = _931_v56
				_ = _rhs199
				var _rhs200 _dafny.Sequence = _932_v57
				_ = _rhs200
				var _rhs201 bool = _881_v21
				_ = _rhs201
				var _rhs202 _dafny.CodePoint = (func() _dafny.CodePoint {
					if _881_v21 {
						return _927_v52
					}
					return _927_v52
				})()
				_ = _rhs202
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				_lhs140.F1 = _rhs198
				_931_v56 = _rhs199
				_932_v57 = _rhs200
				_881_v21 = _rhs201
				_927_v52 = _rhs202
				_927_v52 = _927_v52
			}
			var _933_v58 _dafny.Array
			_ = _933_v58
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_22
			var _nw150 _dafny.Array
			_ = _nw150
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw150 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) bool = (func(_934_v21 bool) func(_dafny.Int) bool {
					return func(_935_i8 _dafny.Int) bool {
						return _934_v21
					}
				})(_881_v21)
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw150 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw150).ArraySet1(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw150).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_933_v58 = _nw150
			var _936_v59 _dafny.Map
			_ = _936_v59
			_936_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_933_v58, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool))
			var _937_v60 *C3
			_ = _937_v60
			var _nw151 *C3 = New_C3_()
			_ = _nw151
			_nw151.Ctor__((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool), (func() bool {
				if (_936_v59).Contains(_933_v58) {
					return (_936_v59).Get(_933_v58).(bool)
				}
				return (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(600), _dafny.ArrayLen((p0), 0))).Int()).(bool)
			})())
			_937_v60 = _nw151
		}
		var _938_v61 bool
		_ = _938_v61
		_938_v61 = false
		var _939_v62 T1
		_ = _939_v62
		var _nw152 *C3 = New_C3_()
		_ = _nw152
		_nw152.Ctor__(_938_v61, _938_v61)
		_939_v62 = _nw152
		var _hi5 _dafny.Int = _878_v19
		_ = _hi5
		for _940_i9 := (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v62, _dafny.CodePoint('l'))).Cardinality(); _940_i9.Cmp(_hi5) < 0; _940_i9 = _940_i9.Plus(_dafny.One) {
			var _941_v63 _dafny.Sequence
			_ = _941_v63
			_941_v63 = _dafny.UnicodeSeqOfUtf8Bytes("amn")
			var _942_v64 _dafny.Map
			_ = _942_v64
			_942_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_940_i9, _dafny.IntOfUint32((_941_v63).Cardinality()))
			var _943_v65 _dafny.Map
			_ = _943_v65
			_943_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_941_v63, _940_i9)
			(globalState).F1 = ((_942_v64).Cardinality()).Plus(Companion_Default___.SafeModuloInt(_940_i9, (_943_v65).Cardinality()))
			_878_v19 = _878_v19
			var _944_v66 _dafny.Array
			_ = _944_v66
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_23
			var _nw153 _dafny.Array
			_ = _nw153
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw153 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Sequence = func(_945_i10 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg60 _dafny.Int) interface{} {
							return coer60(arg60)
						}
					}(func(_946_i11 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('o')
					}))
				}
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw153 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw153).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw153).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_944_v66 = _nw153
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_944_v66), 0))
			_ = _index178
			(_944_v66).ArraySet1(_941_v63, (_index178).Int())
			var _947_v67 _dafny.Sequence
			_ = _947_v67
			_947_v67 = _dafny.SeqOf(!(_938_v61))
			var _948_v68 _dafny.MultiSet
			_ = _948_v68
			_948_v68 = _dafny.MultiSetOf(_878_v19)
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(547), _dafny.ArrayLen((_944_v66), 0))
			_ = _index179
			(_944_v66).ArraySet1(Companion_Default___.Fm1((_dafny.IntOfUint32((_947_v67).Cardinality())).Plus((func() _dafny.Int {
				if (_948_v68).Contains(_940_i9) {
					return (_948_v68).Multiplicity(_940_i9)
				}
				return _940_i9
			})()), _938_v61, _940_i9, globalState), (_index179).Int())
			var _949_v69 _dafny.Map
			_ = _949_v69
			_949_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("buasua")).Cardinality()), false)
			var _950_v70 _dafny.Map
			_ = _950_v70
			_950_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v61, true)
			var _951_v71 _dafny.CodePoint
			_ = _951_v71
			_951_v71 = _dafny.CodePoint('t')
			var _952_v72 _dafny.Map
			_ = _952_v72
			_952_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v61, _951_v71)
			var _953_v73 _dafny.Array
			_ = _953_v73
			var _nwElement0_26 _dafny.Int = _940_i9
			_ = _nwElement0_26
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(19))
			_ = _nw154
			(_nw154).ArraySet1(_nwElement0_26, 0)
			(_nw154).ArraySet1(_940_i9, 1)
			(_nw154).ArraySet1(_878_v19, 2)
			(_nw154).ArraySet1(_878_v19, 3)
			(_nw154).ArraySet1(_dafny.IntOfUint32((_879_v20).Cardinality()), 4)
			(_nw154).ArraySet1((_dafny.Zero).Minus(_940_i9), 5)
			(_nw154).ArraySet1(_878_v19, 6)
			(_nw154).ArraySet1(_878_v19, 7)
			(_nw154).ArraySet1(_878_v19, 8)
			(_nw154).ArraySet1((_939_v62).Fm4(_dafny.SeqOf((func() bool {
				if (_949_v69).Contains((_950_v70).Cardinality()) {
					return (_949_v69).Get((_950_v70).Cardinality()).(bool)
				}
				return _938_v61
			})()), _940_i9, _dafny.IntOfInt64(301), _948_v68, globalState), 9)
			(_nw154).ArraySet1(_940_i9, 10)
			(_nw154).ArraySet1((_952_v72).Cardinality(), 11)
			(_nw154).ArraySet1(_878_v19, 12)
			(_nw154).ArraySet1(_878_v19, 13)
			(_nw154).ArraySet1(_940_i9, 14)
			(_nw154).ArraySet1((_dafny.Zero).Minus(_940_i9), 15)
			(_nw154).ArraySet1(_878_v19, 16)
			(_nw154).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(true, _938_v61, true, _938_v61)).Cardinality()), 17)
			(_nw154).ArraySet1(_940_i9, 18)
			_953_v73 = _nw154
			var _954_v74 _dafny.Sequence
			_ = _954_v74
			_954_v74 = _dafny.SeqOf((func() _dafny.Array {
				if !(_938_v61) {
					return _953_v73
				}
				return _953_v73
			})(), _953_v73, _953_v73, _953_v73)
			_954_v74 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_954_v74, _954_v74), _dafny.SeqOf(_953_v73, _953_v73, (_954_v74).Select((Companion_Default___.SafeIndex(_878_v19, _dafny.IntOfUint32((_954_v74).Cardinality()))).Uint32()).(_dafny.Array), _953_v73, _953_v73))
		}
		(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_878_v19, _878_v19))
		if (_938_v61) || (_938_v61) {
			var _955_v75 _dafny.Sequence
			_ = _955_v75
			_955_v75 = _dafny.UnicodeSeqOfUtf8Bytes("ig")
			_955_v75 = _955_v75
			(globalState).F1 = _878_v19
			_878_v19 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_878_v19, _878_v19), _878_v19)
			(globalState).F1 = _878_v19
			var _956_v76 _dafny.Set
			_ = _956_v76
			_956_v76 = _dafny.SetOf(_938_v61, ((_dafny.Zero).Minus(_878_v19)).Cmp(_878_v19) != 0)
			_956_v76 = _956_v76
		} else {
			_938_v61 = _938_v61
			var _957_v77 _dafny.Map
			_ = _957_v77
			_957_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_878_v19, _878_v19)
			var _958_v78 _dafny.Map
			_ = _958_v78
			_958_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_957_v77, _938_v61)
			var _959_v79 _dafny.Sequence
			_ = _959_v79
			_959_v79 = _dafny.SeqOf(false, _938_v61)
			var _960_v80 _dafny.MultiSet
			_ = _960_v80
			_960_v80 = _dafny.MultiSetOf(_959_v79)
			_878_v19 = ((_878_v19).Plus((_dafny.Zero).Minus(_878_v19))).Times((_this).Fm30(_958_v78, _960_v80, globalState))
			var _961_v81 _dafny.Sequence
			_ = _961_v81
			_961_v81 = _dafny.UnicodeSeqOfUtf8Bytes("xcfaexkop")
			var _962_v82 _dafny.Array
			_ = _962_v82
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_24
			var _nw155 _dafny.Array
			_ = _nw155
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw155 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Int = (func(_963_v19 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_964_i12 _dafny.Int) _dafny.Int {
						return (_964_i12).Plus(_963_v19)
					}
				})(_878_v19)
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw155 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw155).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw155).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_962_v82 = _nw155
			var _965_v83 _dafny.Array
			_ = _965_v83
			_965_v83 = _962_v82
			var _966_v84 _dafny.Map
			_ = _966_v84
			_966_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_938_v61, _965_v83)
			var _967_v85 _dafny.MultiSet
			_ = _967_v85
			_967_v85 = _dafny.MultiSetOf((_966_v84).Cardinality())
			var _968_v86 D2
			_ = _968_v86
			_968_v86 = Companion_D2_.Create_DC7_((_this).Fm4(_959_v79, _878_v19, _dafny.IntOfUint32((_959_v79).Cardinality()), (_967_v85).Update(_878_v19, Companion_Default___.Abs(_878_v19)), globalState), _878_v19)
			var _969_v87 _dafny.Map
			_ = _969_v87
			_969_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_961_v81, (_968_v86).Dtor_cf7())
			_969_v87 = (_969_v87).Update(_961_v81, _878_v19)
			var _970_v88 _dafny.Map
			_ = _970_v88
			_970_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_961_v81, Companion_Default___.Fm1(_878_v19, _938_v61, _dafny.IntOfInt64(726), globalState))).Cardinality()), true)
			_970_v88 = (_970_v88).Update(_878_v19, _938_v61)
			(globalState).F1 = _dafny.IntOfInt64(-63)
		}
		var _971_v89 _dafny.Array
		_ = _971_v89
		var _nw156 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(21))
		_ = _nw156
		_971_v89 = _nw156
		_971_v89 = _971_v89
		var _972_v90 _dafny.CodePoint
		_ = _972_v90
		_972_v90 = _dafny.CodePoint('a')
		var _973_v91 *C0
		_ = _973_v91
		var _nw157 *C0 = New_C0_()
		_ = _nw157
		_nw157.Ctor__(_938_v61, (func() _dafny.Sequence {
			if true {
				return _dafny.SeqOf(_879_v20)
			}
			return Companion_Default___.Fm38(Companion_Default___.Fm0(_938_v61, globalState), _938_v61, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bj")).Cardinality()), _972_v90, globalState)
		})())
		_973_v91 = _nw157
		r0 = _973_v91
		r1 = p0
		r2 = p0
		return r0, r1, r2
	}
}
func (_this *C4) M11(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) {
	{
		var _974_v0 _dafny.Sequence
		_ = _974_v0
		_974_v0 = _dafny.UnicodeSeqOfUtf8Bytes("latlvtxl")
		var _975_v1 D2
		_ = _975_v1
		_975_v1 = Companion_D2_.Create_DC6_(_974_v0)
		var _source13 D2 = _975_v1
		_ = _source13
		if _source13.Is_DC7() {
			var _976___mcc_h0 _dafny.Int = _source13.Get_().(D2_DC7).Cf7
			_ = _976___mcc_h0
			var _977___mcc_h1 _dafny.Int = _source13.Get_().(D2_DC7).Cf8
			_ = _977___mcc_h1
			var _978_cf8 _dafny.Int = _977___mcc_h1
			_ = _978_cf8
			var _979_cf7 _dafny.Int = _976___mcc_h0
			_ = _979_cf7
			var _980_v2 D2
			_ = _980_v2
			_980_v2 = Companion_D2_.Create_DC7_(p1, (_dafny.Zero).Minus(_978_cf8))
			(globalState).F1 = ((_980_v2).Dtor_cf8()).Times(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(305), _979_cf7))
			var _981_v3 _dafny.Array
			_ = _981_v3
			var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw158
			_981_v3 = _nw158
			_981_v3 = _981_v3
			var _982_v4 D2
			_ = _982_v4
			_982_v4 = Companion_D2_.Create_DC8_(p2)
			var _pat_let_tv13 = _978_cf8
			_ = _pat_let_tv13
			var _source14 D1 = func(_pat_let10_0 D1) D1 {
				return func(_983_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let11_0 _dafny.Int) D1 {
						return func(_984_dt__update_hcf4_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_(_984_dt__update_hcf4_h0)
						}(_pat_let11_0)
					}(_pat_let_tv13)
				}(_pat_let10_0)
			}((_this).Fm5(_982_v4, globalState))
			_ = _source14
			if _source14.Is_DC5() {
				var _985___mcc_h5 bool = _source14.Get_().(D1_DC5).Cf5
				_ = _985___mcc_h5
				var _986_cf5 bool = _985___mcc_h5
				_ = _986_cf5
				(globalState).F1 = Companion_Default___.Fm3(p2, p2, globalState)
				var _rhs203 bool = p2
				_ = _rhs203
				var _rhs204 _dafny.Int = Companion_Default___.SafeDivisionInt(_979_cf7, Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_979_cf7), (_dafny.Zero).Minus(_dafny.IntOfUint32((_974_v0).Cardinality()))))
				_ = _rhs204
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				_986_cf5 = _rhs203
				_lhs141.F1 = _rhs204
				var _987_v5 _dafny.CodePoint
				_ = _987_v5
				_987_v5 = _dafny.CodePoint('c')
				var _988_v6 _dafny.Map
				_ = _988_v6
				_988_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_987_v5, _978_cf8)
				_988_v6 = _988_v6
				(globalState).F1 = p1
			} else {
				var _989___mcc_h6 _dafny.Int = _source14.Get_().(D1_DC4).Cf4
				_ = _989___mcc_h6
				var _990_cf4 _dafny.Int = _989___mcc_h6
				_ = _990_cf4
				var _991_v7 bool
				_ = _991_v7
				_991_v7 = false
				_991_v7 = p2
				var _992_v8 D0
				_ = _992_v8
				_992_v8 = Companion_D0_.Create_DC0_(p2)
				var _993_v9 D0
				_ = _993_v9
				_993_v9 = Companion_D0_.Create_DC3_(_992_v8)
				var _994_v10 D0
				_ = _994_v10
				_994_v10 = Companion_D0_.Create_DC3_(_993_v9)
				var _995_v11 D0
				_ = _995_v11
				_995_v11 = Companion_D0_.Create_DC3_(_992_v8)
				var _996_v12 _dafny.Sequence
				_ = _996_v12
				_996_v12 = _dafny.SeqOf(_991_v7, true, false)
				var _997_v13 _dafny.Sequence
				_ = _997_v13
				_997_v13 = _dafny.SeqOf((Companion_D0_.Create_DC3_(_992_v8)).Dtor_cf3())
				var _998_v14 _dafny.Array
				_ = _998_v14
				var _len0_25 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_25
				var _nw159 _dafny.Array
				_ = _nw159
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw159 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) _dafny.Int = func(_999_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_999_i0, _dafny.IntOfInt64(864))
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw159 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw159).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw159).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_998_v14 = _nw159
				var _1000_v15 _dafny.Map
				_ = _1000_v15
				_1000_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v14, p1)
				var _pat_let_tv14 = _997_v13
				_ = _pat_let_tv14
				var _pat_let_tv15 = _990_cf4
				_ = _pat_let_tv15
				var _pat_let_tv16 = _997_v13
				_ = _pat_let_tv16
				var _rhs205 D0 = func(_pat_let12_0 D0) D0 {
					return func(_1001_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let13_0 D0) D0 {
							return func(_1002_dt__update_hcf3_h0 D0) D0 {
								return Companion_D0_.Create_DC3_(_1002_dt__update_hcf3_h0)
							}(_pat_let13_0)
						}((_pat_let_tv14).Select((Companion_Default___.SafeIndex(_pat_let_tv15, _dafny.IntOfUint32((_pat_let_tv16).Cardinality()))).Uint32()).(D0))
					}(_pat_let12_0)
				}(_995_v11)
				_ = _rhs205
				var _rhs206 _dafny.Int = ((_dafny.Zero).Minus((_1000_v15).Cardinality())).Plus(_990_cf4)
				_ = _rhs206
				var _rhs207 _dafny.Sequence = _996_v12
				_ = _rhs207
				_995_v11 = _rhs205
				_978_cf8 = _rhs206
				_996_v12 = _rhs207
				_991_v7 = p2
				_974_v0 = _974_v0
			}
			var _1003_v16 D0
			_ = _1003_v16
			_1003_v16 = Companion_D0_.Create_DC1_()
			var _1004_v17 _dafny.Array
			_ = _1004_v17
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_26
			var _nw160 _dafny.Array
			_ = _nw160
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw160 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Sequence = (func(_1005_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1006_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate((Companion_D2_.Create_DC6_(_dafny.UnicodeSeqOfUtf8Bytes("vgtgm"))).Dtor_cf6(), _1005_v0)
					}
				})(_974_v0)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw160 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw160).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw160).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1004_v17 = _nw160
			var _1007_v18 _dafny.CodePoint
			_ = _1007_v18
			_1007_v18 = _dafny.CodePoint('o')
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1004_v17), 0))
			_ = _index180
			(_1004_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm1(p1, p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-826))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg61 _dafny.Int) interface{} {
					return coer61(arg61)
				}
			}(func(_1008_i2 _dafny.Int) _dafny.CodePoint {
				return (Companion_D11_.Create_DC19_(_dafny.CodePoint('q'))).Dtor_cf20()
			}))).Cardinality()), globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((Companion_Default___.Fm1(p1, p0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-826))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}(func(_1009_i2 _dafny.Int) _dafny.CodePoint {
				return (Companion_D11_.Create_DC19_(_dafny.CodePoint('q'))).Dtor_cf20()
			}))).Cardinality()), globalState)).Cardinality()))).Uint32(), _1007_v18), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_1010_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1011_i3 _dafny.Int) _dafny.CodePoint {
					return _1010_v18
				}
			})(_1007_v18)))), (_index180).Int())
			var _1012_v19 _dafny.Map
			_ = _1012_v19
			_1012_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_cf8, p2)
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1004_v17), 0))
			_ = _index181
			var _rhs208 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(p1, _979_cf7), ((_1012_v19).Cardinality()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vfndmpit")).Cardinality()))))
			_ = _rhs208
			var _rhs209 D0 = _1003_v16
			_ = _rhs209
			var _rhs210 _dafny.Sequence = _974_v0
			_ = _rhs210
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			var _lhs143 _dafny.Array = _1004_v17
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1004_v17), 0))
			_ = _lhs144
			_lhs142.F1 = _rhs208
			_1003_v16 = _rhs209
			(_lhs143).ArraySet1(_rhs210, (_lhs144).Int())
		} else if _source13.Is_DC8() {
			var _1013___mcc_h2 bool = _source13.Get_().(D2_DC8).Cf9
			_ = _1013___mcc_h2
			var _1014_cf9 bool = _1013___mcc_h2
			_ = _1014_cf9
			var _1015_v20 _dafny.Map
			_ = _1015_v20
			_1015_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1014_cf9, p2)
			if (func() bool {
				if (_1015_v20).Contains(p2) {
					return (_1015_v20).Get(p2).(bool)
				}
				return p2
			})() {
				var _1016_v21 _dafny.CodePoint
				_ = _1016_v21
				_1016_v21 = _dafny.CodePoint('k')
				var _1017_v22 _dafny.Map
				_ = _1017_v22
				_1017_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_974_v0, _1016_v21)
				(globalState).F1 = (_1017_v22).Cardinality()
				var _1018_v23 _dafny.Map
				_ = _1018_v23
				_1018_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
				var _1019_v24 D11
				_ = _1019_v24
				_1019_v24 = Companion_D11_.Create_DC19_(_1016_v21)
				var _rhs211 _dafny.Map = (_1018_v23).Merge(_1018_v23)
				_ = _rhs211
				var _rhs212 D11 = Companion_D11_.Create_DC19_(_1016_v21)
				_ = _rhs212
				var _rhs213 _dafny.Sequence = _974_v0
				_ = _rhs213
				var _rhs214 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_974_v0, (Companion_D2_.Create_DC6_(_974_v0)).Dtor_cf6()), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_974_v0, (Companion_D2_.Create_DC6_(_974_v0)).Dtor_cf6())).Cardinality()))).Uint32(), _1016_v21)
				_ = _rhs214
				_1018_v23 = _rhs211
				_1019_v24 = _rhs212
				_974_v0 = _rhs213
				_974_v0 = _rhs214
				var _1020_v25 _dafny.Array
				_ = _1020_v25
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_27
				var _nw161 _dafny.Array
				_ = _nw161
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw161 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = (func(_1021_p0 bool) func(_dafny.Int) bool {
						return func(_1022_i4 _dafny.Int) bool {
							return _1021_p0
						}
					})(p0)
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw161 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw161).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw161).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1020_v25 = _nw161
				var _1023_v26 _dafny.MultiSet
				_ = _1023_v26
				_1023_v26 = _dafny.MultiSetOf(false, p0, p0)
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_1020_v25), 0))
				_ = _index182
				(_1020_v25).ArraySet1((_1023_v26).IsProperSubsetOf(Companion_Default___.Fm34(p2, globalState)), (_index182).Int())
				var _1024_v27 _dafny.Map
				_ = _1024_v27
				_1024_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_974_v0, ((_1015_v20).Cardinality()).Cmp(p1) != 0)
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(297), _dafny.ArrayLen((_1020_v25), 0))
				_ = _index183
				(_1020_v25).ArraySet1((func() bool {
					if (_1024_v27).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg64 _dafny.Int) interface{} {
							return coer64(arg64)
						}
					}((func(_1025_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1026_i5 _dafny.Int) _dafny.CodePoint {
							return _1025_v21
						}
					})(_1016_v21)))) {
						return (_1024_v27).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg65 _dafny.Int) interface{} {
								return coer65(arg65)
							}
						}((func(_1027_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1028_i5 _dafny.Int) _dafny.CodePoint {
								return _1027_v21
							}
						})(_1016_v21)))).(bool)
					}
					return _1014_cf9
				})(), (_index183).Int())
				(globalState).F1 = (p1).Times(p1)
				var _1029_v28 _dafny.Map
				_ = _1029_v28
				_1029_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1020_v25, _1016_v21)
				var _1030_v29 D0
				_ = _1030_v29
				_1030_v29 = Companion_D0_.Create_DC2_(_1029_v28, _1016_v21)
				var _1031_v30 _dafny.Map
				_ = _1031_v30
				_1031_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1030_v29, _1014_cf9)
				_1031_v30 = _1031_v30
			} else {
				(globalState).F1 = Companion_Default___.Fm3(p0, _1014_cf9, globalState)
				var _1032_v31 _dafny.Array
				_ = _1032_v31
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw162
				_1032_v31 = _nw162
				var _1033_v32 _dafny.CodePoint
				_ = _1033_v32
				_1033_v32 = _dafny.CodePoint('j')
				var _1034_v33 D0
				_ = _1034_v33
				var _1035_v34 bool
				_ = _1035_v34
				var _out30 D0
				_ = _out30
				var _out31 bool
				_ = _out31
				_out30, _out31 = Companion_Default___.M0(_1032_v31, _1033_v32, p0, _1033_v32, globalState)
				_1034_v33 = _out30
				_1035_v34 = _out31
				(globalState).F1 = (_dafny.Zero).Minus(p1)
				var _1036_v35 _dafny.Map
				_ = _1036_v35
				_1036_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_974_v0, p0)
				var _1037_v36 D0
				_ = _1037_v36
				var _1038_v37 bool
				_ = _1038_v37
				var _out32 D0
				_ = _out32
				var _out33 bool
				_ = _out33
				_out32, _out33 = Companion_Default___.M0(_1032_v31, _1033_v32, (func() bool {
					if (_1036_v35).Contains(_974_v0) {
						return (_1036_v35).Get(_974_v0).(bool)
					}
					return _1035_v34
				})(), _1033_v32, globalState)
				_1037_v36 = _out32
				_1038_v37 = _out33
				var _1039_v38 _dafny.Map
				_ = _1039_v38
				_1039_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfInt64(699))
				_1039_v38 = (_1039_v38).Update(_1014_cf9, _dafny.IntOfUint32((_974_v0).Cardinality()))
			}
			var _1040_v39 _dafny.Map
			_ = _1040_v39
			_1040_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _1041_v40 _dafny.Map
			_ = _1041_v40
			_1041_v40 = _1040_v39
			var _1042_v41 D6
			_ = _1042_v41
			_1042_v41 = Companion_D6_.Create_DC14_(false)
			var _1043_v42 _dafny.Map
			_ = _1043_v42
			_1043_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1041_v40), (_1042_v41).Dtor_cf15())
			var _1044_v43 _dafny.Sequence
			_ = _1044_v43
			_1044_v43 = _dafny.SeqOf(p0)
			var _1045_v44 _dafny.MultiSet
			_ = _1045_v44
			_1045_v44 = _dafny.MultiSetOf(_1044_v43)
			var _1046_v45 _dafny.MultiSet
			_ = _1046_v45
			_1046_v45 = _dafny.MultiSetOf((_this).Fm30(_1043_v42, _1045_v44, globalState), p1)
			var _1047_v46 _dafny.Array
			_ = _1047_v46
			var _nwElement0_27 _dafny.Int = p1
			_ = _nwElement0_27
			var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(14))
			_ = _nw163
			(_nw163).ArraySet1(_nwElement0_27, 0)
			(_nw163).ArraySet1(p1, 1)
			(_nw163).ArraySet1(p1, 2)
			(_nw163).ArraySet1(p1, 3)
			(_nw163).ArraySet1(p1, 4)
			(_nw163).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rodcjqa")).Cardinality()), 5)
			(_nw163).ArraySet1((_dafny.Zero).Minus(p1), 6)
			(_nw163).ArraySet1(_dafny.IntOfInt64(96), 7)
			(_nw163).ArraySet1(p1, 8)
			(_nw163).ArraySet1(p1, 9)
			(_nw163).ArraySet1(p1, 10)
			(_nw163).ArraySet1(_dafny.IntOfUint32((_974_v0).Cardinality()), 11)
			(_nw163).ArraySet1(p1, 12)
			(_nw163).ArraySet1(p1, 13)
			_1047_v46 = _nw163
			var _1048_v47 _dafny.Map
			_ = _1048_v47
			_1048_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1047_v46, false)
			_1046_v45 = _dafny.MultiSetOf((_1048_v47).Cardinality())
			var _1049_v48 _dafny.Sequence
			_ = _1049_v48
			_1049_v48 = _dafny.SeqOf(_974_v0, _dafny.UnicodeSeqOfUtf8Bytes("i"), _dafny.Companion_Sequence_.Concatenate(_974_v0, _974_v0))
			_1049_v48 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-320))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}((func(_1050_v0 _dafny.Sequence, _1051_p1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
				return func(_1052_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Update(_1050_v0, (Companion_Default___.SafeIndex(_1051_p1, _dafny.IntOfUint32((_1050_v0).Cardinality()))).Uint32(), _dafny.CodePoint('o'))
				}
			})(_974_v0, p1)))
			var _1053_v49 _dafny.Sequence
			_ = _1053_v49
			_1053_v49 = _dafny.SeqOf(p1)
			_1053_v49 = _1053_v49
		} else if _source13.Is_DC6() {
			var _1054___mcc_h3 _dafny.Sequence = _source13.Get_().(D2_DC6).Cf6
			_ = _1054___mcc_h3
			var _1055_cf6 _dafny.Sequence = _1054___mcc_h3
			_ = _1055_cf6
			var _1056_v50 bool
			_ = _1056_v50
			_1056_v50 = true
			_1056_v50 = _1056_v50
			var _1057_v51 _dafny.Map
			_ = _1057_v51
			_1057_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1056_v50, p0)
			var _1058_v52 D12
			_ = _1058_v52
			_1058_v52 = Companion_D12_.Create_DC21_(_1057_v51)
			var _1059_v53 _dafny.Set
			_ = _1059_v53
			_1059_v53 = _dafny.SetOf(_1057_v51, (_1058_v52).Dtor_cf23())
			(globalState).F1 = ((_1059_v53).Intersection(_1059_v53)).Cardinality()
			var _1060_v54 *C3
			_ = _1060_v54
			var _nw164 *C3 = New_C3_()
			_ = _nw164
			_nw164.Ctor__(p2, true)
			_1060_v54 = _nw164
			var _1061_v55 _dafny.Sequence
			_ = _1061_v55
			_1061_v55 = _dafny.SeqOf(p2, (_1060_v54).F13(), (_1060_v54).F14())
			_1061_v55 = _dafny.Companion_Sequence_.Concatenate(_1061_v55, _1061_v55)
		} else {
			var _1062___mcc_h4 D2 = _source13.Get_().(D2_DC9).Cf10
			_ = _1062___mcc_h4
			var _1063_cf10 D2 = _1062___mcc_h4
			_ = _1063_cf10
			var _1064_v56 _dafny.Sequence
			_ = _1064_v56
			_1064_v56 = _dafny.SeqOf(p0)
			_1064_v56 = _dafny.SeqOf(p2, p0, p0, p2)
			var _1065_v57 _dafny.CodePoint
			_ = _1065_v57
			_1065_v57 = _dafny.CodePoint('p')
			var _1066_v58 _dafny.Array
			_ = _1066_v58
			var _nwElement0_28 bool = true
			_ = _nwElement0_28
			var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(9))
			_ = _nw165
			(_nw165).ArraySet1(_nwElement0_28, 0)
			(_nw165).ArraySet1(p0, 1)
			(_nw165).ArraySet1(p2, 2)
			(_nw165).ArraySet1((p1).Cmp(p1) > 0, 3)
			(_nw165).ArraySet1(p0, 4)
			(_nw165).ArraySet1(!(true) || (p2), 5)
			(_nw165).ArraySet1(!(!(!(_dafny.Companion_Sequence_.Contains(_974_v0, _1065_v57)))), 6)
			(_nw165).ArraySet1(p2, 7)
			(_nw165).ArraySet1((func() bool {
				if p0 {
					return !(p0)
				}
				return p0
			})(), 8)
			_1066_v58 = _nw165
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1066_v58), 0))
			_ = _index184
			(_1066_v58).ArraySet1(p2, (_index184).Int())
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1066_v58), 0))
			_ = _index185
			var _rhs215 bool = p0
			_ = _rhs215
			var _rhs216 _dafny.CodePoint = _1065_v57
			_ = _rhs216
			var _lhs145 _dafny.Array = _1066_v58
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1066_v58), 0))
			_ = _lhs146
			(_lhs145).ArraySet1(_rhs215, (_lhs146).Int())
			_1065_v57 = _rhs216
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1066_v58), 0))
			_ = _index186
			(_1066_v58).ArraySet1((_1064_v56).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1064_v56).Cardinality()))).Uint32()).(bool), (_index186).Int())
			var _1067_v59 _dafny.Sequence
			_ = _1067_v59
			_1067_v59 = _dafny.SeqOf(p1)
			var _1068_v60 _dafny.Sequence
			_ = _1068_v60
			_1068_v60 = _dafny.SeqOf(_1067_v59, _dafny.SeqOf(_dafny.IntOfInt64(743)))
			var _1069_v61 *C0
			_ = _1069_v61
			var _nw166 *C0 = New_C0_()
			_ = _nw166
			_nw166.Ctor__((_1066_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1066_v58), 0))).Int()).(bool), _1068_v60)
			_1069_v61 = _nw166
		}
		var _1070_i7 _dafny.Int
		_ = _1070_i7
		_1070_i7 = _dafny.Zero
		{
			for (p2) == (p0) {
				{
					if (_1070_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1070_i7 = (_1070_i7).Plus(_dafny.One)
					var _1071_v62 bool
					_ = _1071_v62
					_1071_v62 = false
					_1071_v62 = p2
					var _1072_v63 _dafny.MultiSet
					_ = _1072_v63
					_1072_v63 = _dafny.MultiSetOf(_dafny.SetOf(_1071_v62, p2, p0))
					var _1073_v64 D11
					_ = _1073_v64
					_1073_v64 = Companion_D11_.Create_DC20_(_1071_v62, p1)
					var _1074_v65 _dafny.Set
					_ = _1074_v65
					_1074_v65 = _dafny.SetOf(p1, (_dafny.Zero).Minus((_1073_v64).Dtor_cf22()))
					var _1075_v66 _dafny.Map
					_ = _1075_v66
					_1075_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1072_v63).IsDisjointFrom(_1072_v63), (_1074_v65).Cardinality())
					_1075_v66 = (_1075_v66).Update(_1071_v62, p1)
					_975_v1 = Companion_D2_.Create_DC6_(_974_v0)
					(globalState).F1 = p1
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		if !(p0) || (p2) {
			var _1076_v67 _dafny.CodePoint
			_ = _1076_v67
			_1076_v67 = _dafny.CodePoint('f')
			var _1077_v68 *C2
			_ = _1077_v68
			var _nw167 *C2 = New_C2_()
			_ = _nw167
			_nw167.Ctor__(_1076_v67, true)
			_1077_v68 = _nw167
			var _1078_v69 D0
			_ = _1078_v69
			_1078_v69 = Companion_D0_.Create_DC0_(true)
			var _1079_v70 D0
			_ = _1079_v70
			_1079_v70 = Companion_D0_.Create_DC3_(_1078_v69)
			var _pat_let_tv17 = _1078_v69
			_ = _pat_let_tv17
			var _source15 D0 = func(_pat_let14_0 D0) D0 {
				return func(_1080_dt__update__tmp_h2 D0) D0 {
					return func(_pat_let15_0 D0) D0 {
						return func(_1081_dt__update_hcf3_h1 D0) D0 {
							return Companion_D0_.Create_DC3_(_1081_dt__update_hcf3_h1)
						}(_pat_let15_0)
					}(_pat_let_tv17)
				}(_pat_let14_0)
			}(_1079_v70)
			_ = _source15
			if _source15.Is_DC1() {
				var _1082_v71 _dafny.Map
				_ = _1082_v71
				_1082_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _1083_v72 D12
				_ = _1083_v72
				_1083_v72 = Companion_D12_.Create_DC21_(_1082_v71)
				_1083_v72 = _1083_v72
				var _1084_v73 _dafny.Map
				_ = _1084_v73
				_1084_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm35(_974_v0, _dafny.MultiSetOf(p1, (_dafny.Zero).Minus(_dafny.IntOfInt64(-114))), _1076_v67, globalState))
				var _1085_v74 _dafny.Set
				_ = _1085_v74
				_1085_v74 = _dafny.SetOf(p1, p1)
				_1084_v73 = (_1084_v73).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_974_v0, _974_v0)).Cardinality()), _1085_v74)
				var _1086_v75 bool
				_ = _1086_v75
				_1086_v75 = true
				_1086_v75 = ((func() bool {
					if p2 {
						return p2
					}
					return p0
				})()) || (p2)
				var _1087_v76 _dafny.Array
				_ = _1087_v76
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_28
				var _nw168 _dafny.Array
				_ = _nw168
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw168 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Int = (func(_1088_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1089_i8 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1089_i8, (_dafny.Zero).Minus(_1088_p1))
						}
					})(p1)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw168 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw168).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw168).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1087_v76 = _nw168
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1087_v76), 0))
				_ = _index187
				(_1087_v76).ArraySet1((_dafny.Zero).Minus(p1), (_index187).Int())
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_1087_v76), 0))
				_ = _index188
				(_1087_v76).ArraySet1(p1, (_index188).Int())
			} else if _source15.Is_DC2() {
				var _1090___mcc_h7 _dafny.Map = _source15.Get_().(D0_DC2).Cf1
				_ = _1090___mcc_h7
				var _1091___mcc_h8 _dafny.CodePoint = _source15.Get_().(D0_DC2).Cf2
				_ = _1091___mcc_h8
				var _1092_cf2 _dafny.CodePoint = _1091___mcc_h8
				_ = _1092_cf2
				var _1093_cf1 _dafny.Map = _1090___mcc_h7
				_ = _1093_cf1
				var _1094_v77 _dafny.Array
				_ = _1094_v77
				var _nwElement0_29 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(421))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}((func(_1095_cf2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1096_i9 _dafny.Int) _dafny.CodePoint {
						return _1095_cf2
					}
				})(_1092_cf2)))
				_ = _nwElement0_29
				var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(7))
				_ = _nw169
				(_nw169).ArraySet1(_nwElement0_29, 0)
				(_nw169).ArraySet1(_974_v0, 1)
				(_nw169).ArraySet1(_974_v0, 2)
				(_nw169).ArraySet1(_dafny.Companion_Sequence_.Update(_974_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-556), _dafny.IntOfUint32((_974_v0).Cardinality()))).Uint32(), _1092_cf2), 3)
				(_nw169).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_974_v0, _974_v0), 4)
				(_nw169).ArraySet1(_974_v0, 5)
				(_nw169).ArraySet1(_974_v0, 6)
				_1094_v77 = _nw169
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1094_v77), 0))
				_ = _index189
				(_1094_v77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-35))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}((func(_1097_v67 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1098_i10 _dafny.Int) _dafny.CodePoint {
						return _1097_v67
					}
				})(_1076_v67))), _974_v0), (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1094_v77), 0))
				_ = _index190
				(_1094_v77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("remd"), _974_v0), (_index190).Int())
				var _1099_v78 bool
				_ = _1099_v78
				_1099_v78 = false
				var _1100_v79 _dafny.Sequence
				_ = _1100_v79
				_1100_v79 = _dafny.SeqOf(p0)
				var _1101_v80 _dafny.Sequence
				_ = _1101_v80
				_1101_v80 = _dafny.SeqOf(!(Companion_Default___.Fm0(_1099_v78, globalState)), false, !(p0), _1099_v78)
				var _1102_v81 _dafny.Array
				_ = _1102_v81
				var _nwElement0_30 _dafny.Sequence = _1100_v79
				_ = _nwElement0_30
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(13))
				_ = _nw170
				(_nw170).ArraySet1(_nwElement0_30, 0)
				(_nw170).ArraySet1(_1100_v79, 1)
				(_nw170).ArraySet1(_1100_v79, 2)
				(_nw170).ArraySet1(_1100_v79, 3)
				(_nw170).ArraySet1(_1100_v79, 4)
				(_nw170).ArraySet1(_1100_v79, 5)
				(_nw170).ArraySet1(_1100_v79, 6)
				(_nw170).ArraySet1(_1101_v80, 7)
				(_nw170).ArraySet1(_1100_v79, 8)
				(_nw170).ArraySet1(Companion_Default___.Fm39(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xun")).Cardinality()), p0, p1, p1, globalState), 9)
				(_nw170).ArraySet1(_1100_v79, 10)
				(_nw170).ArraySet1(_1100_v79, 11)
				(_nw170).ArraySet1(_1100_v79, 12)
				_1102_v81 = _nw170
				var _1103_v82 _dafny.Set
				_ = _1103_v82
				_1103_v82 = _dafny.SetOf(_dafny.IntOfUint32(((_1094_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1094_v77), 0))).Int()).(_dafny.Sequence)).Cardinality()))
				var _rhs217 _dafny.Int = (_1077_v68).Fm8(true, _1099_v78, (p1).Cmp((_1103_v82).Cardinality()) == 0, globalState)
				_ = _rhs217
				var _rhs218 bool = (p1).Cmp((p1).Plus(p1)) != 0
				_ = _rhs218
				var _rhs219 _dafny.Array = _1102_v81
				_ = _rhs219
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				_lhs147.F1 = _rhs217
				_1099_v78 = _rhs218
				_1102_v81 = _rhs219
				var _1104_v83 _dafny.Array
				_ = _1104_v83
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw171
				_1104_v83 = _nw171
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))
				_ = _index191
				(_1104_v83).ArraySet1(p1, (_index191).Int())
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))
				_ = _index192
				(_1104_v83).ArraySet1((_dafny.IntOfUint32(((_1094_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1094_v77), 0))).Int()).(_dafny.Sequence)).Cardinality())).Times(p1), (_index192).Int())
				var _1105_v84 _dafny.Array
				_ = _1105_v84
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(4))
				_ = _nw172
				_1105_v84 = _nw172
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1105_v84), 0))
				_ = _index193
				(_1105_v84).ArraySet1(_1103_v82, (_index193).Int())
				var _1106_v86 _dafny.Map
				_ = _1106_v86
				_1106_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int), (_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int))
				var _1107_v87 _dafny.Map
				_ = _1107_v87
				_1107_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1106_v86, _1099_v78)
				var _1108_v88 _dafny.Sequence
				_ = _1108_v88
				_1108_v88 = _dafny.SeqOf(_1101_v80)
				var _1109_v89 _dafny.MultiSet
				_ = _1109_v89
				_1109_v89 = _dafny.MultiSetOf((_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int), p1, (_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int), p1, p1)
				var _1110_v90 _dafny.Sequence
				_ = _1110_v90
				_1110_v90 = _dafny.SeqOf((_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int), p1)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))
				_ = _index194
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1105_v84), 0))
				_ = _index195
				var _rhs220 _dafny.Int = Companion_Default___.SafeModuloInt(p1, ((_this).Fm30(func() _dafny.Map {
					var _coll34 = _dafny.NewMapBuilder()
					_ = _coll34
					for _iter35 := _dafny.Iterate((_1107_v87).Keys().Elements()); ; {
						_compr_34, _ok35 := _iter35()
						if !_ok35 {
							break
						}
						var _1111_v85 _dafny.Map
						_1111_v85 = interface{}(_compr_34).(_dafny.Map)
						if (_1107_v87).Contains(_1111_v85) {
							_coll34.Add(_1111_v85, _1099_v78)
						}
					}
					return _coll34.ToMap()
				}(), _dafny.MultiSetFromSeq(_1108_v88), globalState)).Times(_dafny.IntOfInt64(362)))
				_ = _rhs220
				var _rhs221 _dafny.Set = Companion_Default___.Fm35(_974_v0, _1109_v89, _1076_v67, globalState)
				_ = _rhs221
				var _rhs222 _dafny.Int = (func() _dafny.Int {
					if (_1109_v89).Contains((func() _dafny.Int {
						if _1099_v78 {
							return (_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int)
						}
						return (_1109_v89).Cardinality()
					})()) {
						return (_1109_v89).Multiplicity((func() _dafny.Int {
							if _1099_v78 {
								return (_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int)
							}
							return (_1109_v89).Cardinality()
						})())
					}
					return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(191), p1)
				})()
				_ = _rhs222
				var _rhs223 bool = !(((_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int)).Cmp((_1110_v90).Select((Companion_Default___.SafeIndex((_1104_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1110_v90).Cardinality()))).Uint32()).(_dafny.Int)) == 0)
				_ = _rhs223
				var _lhs148 _dafny.Array = _1104_v83
				_ = _lhs148
				var _lhs149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(827), _dafny.ArrayLen((_1104_v83), 0))
				_ = _lhs149
				var _lhs150 _dafny.Array = _1105_v84
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_1105_v84), 0))
				_ = _lhs151
				var _lhs152 *GlobalState = globalState
				_ = _lhs152
				(_lhs148).ArraySet1(_rhs220, (_lhs149).Int())
				(_lhs150).ArraySet1(_rhs221, (_lhs151).Int())
				_lhs152.F1 = _rhs222
				_1099_v78 = _rhs223
			} else if _source15.Is_DC0() {
				var _1112___mcc_h9 bool = _source15.Get_().(D0_DC0).Cf0
				_ = _1112___mcc_h9
				var _1113_cf0 bool = _1112___mcc_h9
				_ = _1113_cf0
				var _1114_v91 _dafny.Array
				_ = _1114_v91
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
				_ = _nw173
				_1114_v91 = _nw173
				var _1115_v92 _dafny.Array
				_ = _1115_v92
				var _len0_29 _dafny.Int = _dafny.One
				_ = _len0_29
				var _nw174 _dafny.Array
				_ = _nw174
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw174 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) bool = (func(_1116_p2 bool) func(_dafny.Int) bool {
						return func(_1117_i11 _dafny.Int) bool {
							return _1116_p2
						}
					})(p2)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw174 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw174).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw174).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1115_v92 = _nw174
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1114_v91), 0))
				_ = _index196
				(_1114_v91).ArraySet1(_1115_v92, (_index196).Int())
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1114_v91), 0))
				_ = _index197
				var _nw175 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw175
				(_1114_v91).ArraySet1(_nw175, (_index197).Int())
				var _1118_v93 D11
				_ = _1118_v93
				_1118_v93 = Companion_D11_.Create_DC19_(_1076_v67)
				var _1119_v94 _dafny.Map
				_ = _1119_v94
				_1119_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1118_v93, Companion_Default___.SafeModuloInt(p1, p1))
				var _1120_v95 _dafny.MultiSet
				_ = _1120_v95
				_1120_v95 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lrxmykw")).Cardinality()), p1)
				_1119_v94 = (_1119_v94).Update(_1118_v93, (func() _dafny.Int {
					if (_1120_v95).Contains((_this).Fm4(_dafny.SeqOf(_1113_cf0, p0), p1, p1, (_1120_v95).Update(p1, Companion_Default___.Abs(p1)), globalState)) {
						return (_1120_v95).Multiplicity((_this).Fm4(_dafny.SeqOf(_1113_cf0, p0), p1, p1, (_1120_v95).Update(p1, Companion_Default___.Abs(p1)), globalState))
					}
					return p1
				})())
				var _1121_v96 _dafny.Array
				_ = _1121_v96
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_30
				var _nw176 _dafny.Array
				_ = _nw176
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw176 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Int = (func(_1122_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1123_i12 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1123_i12, _1122_p1)
						}
					})(p1)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw176 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw176).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw176).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1121_v96 = _nw176
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1121_v96), 0))
				_ = _index198
				(_1121_v96).ArraySet1(p1, (_index198).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1121_v96), 0))
				_ = _index199
				var _rhs224 _dafny.Int = Companion_Default___.Fm3(p2, p0, globalState)
				_ = _rhs224
				var _rhs225 bool = (p1).Cmp(_dafny.IntOfInt64(970)) > 0
				_ = _rhs225
				var _rhs226 _dafny.Array = _1121_v96
				_ = _rhs226
				var _lhs153 _dafny.Array = _1121_v96
				_ = _lhs153
				var _lhs154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1121_v96), 0))
				_ = _lhs154
				(_lhs153).ArraySet1(_rhs224, (_lhs154).Int())
				_1113_cf0 = _rhs225
				_1121_v96 = _rhs226
				(globalState).F1 = (Companion_Default___.Fm40(globalState)).Cardinality()
			} else {
				var _1124___mcc_h10 D0 = _source15.Get_().(D0_DC3).Cf3
				_ = _1124___mcc_h10
				var _1125_cf3 D0 = _1124___mcc_h10
				_ = _1125_cf3
				var _1126_v97 _dafny.Set
				_ = _1126_v97
				_1126_v97 = _dafny.SetOf((_dafny.Zero).Minus(p1))
				var _1127_v98 _dafny.Array
				_ = _1127_v98
				var _nwElement0_31 _dafny.Int = _dafny.IntOfUint32((_974_v0).Cardinality())
				_ = _nwElement0_31
				var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(16))
				_ = _nw177
				(_nw177).ArraySet1(_nwElement0_31, 0)
				(_nw177).ArraySet1(p1, 1)
				(_nw177).ArraySet1(p1, 2)
				(_nw177).ArraySet1(((_1126_v97).Difference(_1126_v97)).Cardinality(), 3)
				(_nw177).ArraySet1(p1, 4)
				(_nw177).ArraySet1((p1).Times(p1), 5)
				(_nw177).ArraySet1(p1, 6)
				(_nw177).ArraySet1(p1, 7)
				(_nw177).ArraySet1(p1, 8)
				(_nw177).ArraySet1(p1, 9)
				(_nw177).ArraySet1(p1, 10)
				(_nw177).ArraySet1(p1, 11)
				(_nw177).ArraySet1((p1).Minus(_dafny.IntOfInt64(-580)), 12)
				(_nw177).ArraySet1(p1, 13)
				(_nw177).ArraySet1(p1, 14)
				(_nw177).ArraySet1(p1, 15)
				_1127_v98 = _nw177
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1127_v98), 0))
				_ = _index200
				(_1127_v98).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if !(true) {
						return (_dafny.Zero).Minus(p1)
					}
					return p1
				})())), (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1127_v98), 0))
				_ = _index201
				(_1127_v98).ArraySet1((p1).Minus(Companion_Default___.SafeDivisionInt(p1, p1)), (_index201).Int())
				var _1128_v99 bool
				_ = _1128_v99
				_1128_v99 = true
				var _1129_v100 _dafny.Sequence
				_ = _1129_v100
				_1129_v100 = _dafny.SeqOf(_1128_v99, p2)
				_1128_v99 = (_dafny.IntOfUint32((_1129_v100).Cardinality())).Cmp(_dafny.IntOfInt64(589)) < 0
				var _1130_v101 *C2
				_ = _1130_v101
				var _nw178 *C2 = New_C2_()
				_ = _nw178
				_nw178.Ctor__((_974_v0).Select((Companion_Default___.SafeIndex((_1127_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_1127_v98), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_974_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), !(Companion_Default___.Fm0(p0, globalState)))
				_1130_v101 = _nw178
				var _1131_v102 _dafny.Array
				_ = _1131_v102
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw179
				_1131_v102 = _nw179
				var _1132_v103 _dafny.Array
				_ = _1132_v103
				var _nwElement0_32 bool = p0
				_ = _nwElement0_32
				var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(8))
				_ = _nw180
				(_nw180).ArraySet1(_nwElement0_32, 0)
				(_nw180).ArraySet1(p2, 1)
				(_nw180).ArraySet1(p2, 2)
				(_nw180).ArraySet1(p2, 3)
				(_nw180).ArraySet1(p2, 4)
				(_nw180).ArraySet1(false, 5)
				(_nw180).ArraySet1(!(p0), 6)
				(_nw180).ArraySet1(_1128_v99, 7)
				_1132_v103 = _nw180
				var _1133_v104 _dafny.Map
				_ = _1133_v104
				_1133_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1132_v103)
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_1131_v102), 0))
				_ = _index202
				(_1131_v102).ArraySet1((_1133_v104).Merge(_1133_v104), (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(850), _dafny.ArrayLen((_1131_v102), 0))
				_ = _index203
				(_1131_v102).ArraySet1(_1133_v104, (_index203).Int())
			}
			var _1134_v105 *C0
			_ = _1134_v105
			var _nw181 *C0 = New_C0_()
			_ = _nw181
			_nw181.Ctor__(p2, _dafny.SeqOf(_dafny.SeqOf(p1)))
			_1134_v105 = _nw181
			var _1135_v106 _dafny.Array
			_ = _1135_v106
			var _nw182 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw182
			_1135_v106 = _nw182
			var _1136_v107 _dafny.Map
			_ = _1136_v107
			_1136_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1135_v106)
			var _1137_v108 _dafny.Sequence
			_ = _1137_v108
			_1137_v108 = _dafny.SeqOf(_1135_v106)
			var _1138_v109 _dafny.Array
			_ = _1138_v109
			_1138_v109 = (func() _dafny.Array {
				if (_1136_v107).Contains(p1) {
					return (_1136_v107).Get(p1).(_dafny.Array)
				}
				return (_1137_v108).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1137_v108).Cardinality()))).Uint32()).(_dafny.Array)
			})()
			var _1139_v110 _dafny.Map
			_ = _1139_v110
			_1139_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1138_v109, _1135_v106)
			var _1140_v111 _dafny.Map
			_ = _1140_v111
			_1140_v111 = _1139_v110
			_1139_v110 = (_1140_v111)
			var _1141_v112 _dafny.Array
			_ = _1141_v112
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_31
			var _nw183 _dafny.Array
			_ = _nw183
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw183 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Int = (func(_1142_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1143_i13 _dafny.Int) _dafny.Int {
						return (_1143_i13).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1142_p1)).Cardinality()))
					}
				})(p1)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw183 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw183).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw183).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1141_v112 = _nw183
			var _1144_v113 _dafny.Map
			_ = _1144_v113
			_1144_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1141_v112)
			_1141_v112 = (func() _dafny.Array {
				if (_1144_v113).Contains(!(p0)) {
					return (_1144_v113).Get(!(p0)).(_dafny.Array)
				}
				return _1141_v112
			})()
		} else {
			var _1145_v114 _dafny.MultiSet
			_ = _1145_v114
			_1145_v114 = _dafny.MultiSetOf(p2)
			(globalState).F4 = _1145_v114
			var _1146_v115 _dafny.MultiSet
			_ = _1146_v115
			_1146_v115 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(663))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1147_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('y')
			})))
			var _1148_v116 _dafny.Map
			_ = _1148_v116
			_1148_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1146_v115)
			_1148_v116 = (_1148_v116).Update(_dafny.Companion_Sequence_.Equal(_974_v0, _dafny.UnicodeSeqOfUtf8Bytes("bbc")), (_1146_v115).Union(_1146_v115))
			var _1149_v117 D0
			_ = _1149_v117
			_1149_v117 = Companion_D0_.Create_DC0_(p0)
			var _1150_v118 _dafny.Sequence
			_ = _1150_v118
			_1150_v118 = _dafny.SeqOf(_1149_v117)
			var _1151_v119 _dafny.Sequence
			_ = _1151_v119
			_1151_v119 = _dafny.SeqOf(p0)
			var _1152_v120 _dafny.Map
			_ = _1152_v120
			_1152_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v118, _dafny.MultiSetOf(_1151_v119))
			var _1153_v121 _dafny.MultiSet
			_ = _1153_v121
			_1153_v121 = _dafny.MultiSetOf(Companion_Default___.Fm2(p2, p0, globalState))
			_1152_v120 = (_1152_v120).Update(_1150_v118, (_dafny.MultiSetOf(_1151_v119, _1151_v119)).Intersection(_1153_v121))
			var _1154_v122 _dafny.Array
			_ = _1154_v122
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_32
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) bool = (func(_1155_p2 bool) func(_dafny.Int) bool {
					return func(_1156_i15 _dafny.Int) bool {
						return !(_1155_p2) || (_1155_p2)
					}
				})(p2)
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
			_1154_v122 = _nw184
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1154_v122), 0))
			_ = _index204
			(_1154_v122).ArraySet1((func() bool {
				if p0 {
					return p2
				}
				return p2
			})(), (_index204).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1154_v122), 0))
			_ = _index205
			(_1154_v122).ArraySet1(true, (_index205).Int())
			var _1157_v123 _dafny.Set
			_ = _1157_v123
			_1157_v123 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_974_v0, _974_v0))
			var _1158_v124 _dafny.CodePoint
			_ = _1158_v124
			_1158_v124 = _dafny.CodePoint('n')
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1154_v122), 0))
			_ = _index206
			var _rhs227 _dafny.Set = _1157_v123
			_ = _rhs227
			var _rhs228 bool = Companion_Default___.Fm0(((_dafny.Zero).Minus((_dafny.Zero).Minus(p1))).Cmp(_dafny.IntOfInt64(65)) >= 0, globalState)
			_ = _rhs228
			var _rhs229 _dafny.CodePoint = _1158_v124
			_ = _rhs229
			var _lhs155 _dafny.Array = _1154_v122
			_ = _lhs155
			var _lhs156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_1154_v122), 0))
			_ = _lhs156
			_1157_v123 = _rhs227
			(_lhs155).ArraySet1(_rhs228, (_lhs156).Int())
			_1158_v124 = _rhs229
		}
		var _1159_v125 _dafny.Array
		_ = _1159_v125
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_33
		var _nw185 _dafny.Array
		_ = _nw185
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw185 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.Int = (func(_1160_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1161_i16 _dafny.Int) _dafny.Int {
					return (_1161_i16).Minus((_dafny.Zero).Minus(_1160_p1))
				}
			})(p1)
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw185 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw185).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw185).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1159_v125 = _nw185
		var _1162_v126 _dafny.Map
		_ = _1162_v126
		_1162_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
		var _1163_v127 _dafny.Map
		_ = _1163_v127
		_1163_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(p2, globalState), (_1162_v126).Cardinality())
		var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))
		_ = _index207
		(_1159_v125).ArraySet1(((_1163_v127).Cardinality()).Minus(p1), (_index207).Int())
		var _1164_v128 _dafny.Map
		_ = _1164_v128
		_1164_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _974_v0)
		var _1165_v129 _dafny.Sequence
		_ = _1165_v129
		_1165_v129 = _dafny.SeqOf(p2)
		var _1166_v130 _dafny.Map
		_ = _1166_v130
		_1166_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1165_v129, _dafny.IntOfInt64(396))
		var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))
		_ = _index208
		(_1159_v125).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_974_v0, (func() _dafny.Sequence {
			if (_1164_v128).Contains((_1166_v130).Cardinality()) {
				return (_1164_v128).Get((_1166_v130).Cardinality()).(_dafny.Sequence)
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(367))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg70 _dafny.Int) interface{} {
					return coer70(arg70)
				}
			}(func(_1167_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			}))
		})()), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("mvaeqnmre"), _974_v0))).Cardinality()), (_index208).Int())
		var _1168_v131 _dafny.Sequence
		_ = _1168_v131
		_1168_v131 = _1165_v129
		var _source16 _dafny.Sequence = _1168_v131
		_ = _source16
		var _1169___mcc_h11 _dafny.Sequence = _source16
		_ = _1169___mcc_h11
		var _1170_cf16 _dafny.Sequence = _1169___mcc_h11
		_ = _1170_cf16
		var _1171_v132 bool
		_ = _1171_v132
		_1171_v132 = false
		_1171_v132 = p2
		var _1172_v133 D11
		_ = _1172_v133
		_1172_v133 = Companion_D11_.Create_DC20_(p2, _dafny.IntOfInt64(730))
		var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))
		_ = _index209
		(_1159_v125).ArraySet1(Companion_Default___.SafeModuloInt((_1172_v133).Dtor_cf22(), (_1159_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))).Int()).(_dafny.Int)), (_index209).Int())
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))
		_ = _index210
		(_1159_v125).ArraySet1((_1159_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))).Int()).(_dafny.Int), (_index210).Int())
		var _1173_v134 _dafny.Sequence
		_ = _1173_v134
		_1173_v134 = _dafny.SeqOf(p1)
		var _1174_v135 _dafny.Sequence
		_ = _1174_v135
		_1174_v135 = _dafny.SeqOf(_1173_v134, _1173_v134, _1173_v134)
		var _1175_v136 *C0
		_ = _1175_v136
		var _nw186 *C0 = New_C0_()
		_ = _nw186
		_nw186.Ctor__(p2, _1174_v135)
		_1175_v136 = _nw186
		var _1176_v137 _dafny.Sequence
		_ = _1176_v137
		_1176_v137 = _dafny.SeqOf(p1, _dafny.IntOfInt64(709), (_1159_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))).Int()).(_dafny.Int))
		var _1177_v138 _dafny.Sequence
		_ = _1177_v138
		_1177_v138 = _dafny.SeqOf(_1176_v137, _1176_v137)
		var _1178_v139 *C0
		_ = _1178_v139
		var _nw187 *C0 = New_C0_()
		_ = _nw187
		_nw187.Ctor__((_dafny.IntOfUint32((_1176_v137).Cardinality())).Cmp((_1159_v125).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1159_v125), 0))).Int()).(_dafny.Int)) < 0, _dafny.Companion_Sequence_.Concatenate(_1177_v138, _1177_v138))
		_1178_v139 = _nw187
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f5 _dafny.CodePoint
	_f6 bool
	F15 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
	_this.F15 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C5{}
var _ T2 = &C5{}
var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F5() _dafny.CodePoint {
	return _this._f5
}
func (_this *C5) F6() bool {
	return _this._f6
}
func (_this *C5) Ctor__(f15 _dafny.Int, f5 _dafny.CodePoint, f6 bool) {
	{
		(_this).F15 = f15
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C5) Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _this.F15
	}
}
func (_this *C5) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C5) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_this).F6() {
				return _dafny.UnicodeSeqOfUtf8Bytes("mbq")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("pt")
		})()).Cardinality())
	}
}
func (_this *C5) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return _this.F15
	}
}
func (_this *C5) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_this.F15))
	}
}
func (_this *C5) Fm42(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, func() _dafny.Set {
			var _coll35 = _dafny.NewBuilder()
			_ = _coll35
			for _iter36 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}(func(_1179_i0 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			}))).Cardinality()), _this.F15), (_this).F5())).Keys().Elements()); ; {
				_compr_35, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _1180_v0 _dafny.Map
				_1180_v0 = interface{}(_compr_35).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(385))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}(func(_1179_i0 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				}))).Cardinality()), _this.F15), (_this).F5())).Contains(_1180_v0) {
					_coll35.Add(_1180_v0)
				}
			}
			return _coll35.ToSet()
		}())).Merge(func() _dafny.Map {
			var _coll36 = _dafny.NewMapBuilder()
			_ = _coll36
			for _iter37 := _dafny.Iterate((_dafny.SeqOf((Companion_D1_.Create_DC4_(_this.F15)).Dtor_cf4())).Elements()); ; {
				_compr_36, _ok37 := _iter37()
				if !_ok37 {
					break
				}
				var _1181_v1 _dafny.Int
				_1181_v1 = interface{}(_compr_36).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((Companion_D1_.Create_DC4_(_this.F15)).Dtor_cf4()), _1181_v1) {
					_coll36.Add(Companion_Default___.SafeDivisionInt(_1181_v1, _this.F15), _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()))))
				}
			}
			return _coll36.ToMap()
		}())).Merge((func() _dafny.Map {
			var _coll37 = _dafny.NewMapBuilder()
			_ = _coll37
			for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(844), _dafny.IntOfInt64(86))); ; {
				_compr_37, _ok38 := _iter38()
				if !_ok38 {
					break
				}
				var _1182_v2 _dafny.Int
				_1182_v2 = interface{}(_compr_37).(_dafny.Int)
				if ((_dafny.IntOfInt64(844)).Cmp(_1182_v2) <= 0) && ((_1182_v2).Cmp(_dafny.IntOfInt64(86)) < 0) {
					_coll37.Add(Companion_Default___.SafeModuloInt(_1182_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jpe")).Cardinality())), func() _dafny.Set {
						var _coll38 = _dafny.NewBuilder()
						_ = _coll38
						for _iter39 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), _this.F15), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F15)).Cardinality(), _this.F15))).Elements()); ; {
							_compr_38, _ok39 := _iter39()
							if !_ok39 {
								break
							}
							var _1183_v3 _dafny.Map
							_1183_v3 = interface{}(_compr_38).(_dafny.Map)
							if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(415), _this.F15), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F15)).Cardinality(), _this.F15)), _1183_v3) {
								_coll38.Add(_1183_v3)
							}
						}
						return _coll38.ToSet()
					}())
				}
			}
			return _coll37.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, func() _dafny.Set {
			var _coll39 = _dafny.NewBuilder()
			_ = _coll39
			for _iter40 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15), _this.F15)).Keys().Elements()); ; {
				_compr_39, _ok40 := _iter40()
				if !_ok40 {
					break
				}
				var _1184_v4 _dafny.Map
				_1184_v4 = interface{}(_compr_39).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15), _this.F15)).Contains(_1184_v4) {
					_coll39.Add(_1184_v4)
				}
			}
			return _coll39.ToSet()
		}())))
	}
}
func (_this *C5) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _1185_v0 _dafny.Sequence
		_ = _1185_v0
		_1185_v0 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, true)).Cardinality())
		var _hi6 _dafny.Int = p1
		_ = _hi6
		for _1186_i0 := (_dafny.IntOfInt64(535)).Times(_dafny.IntOfUint32((_1185_v0).Cardinality())); _1186_i0.Cmp(_hi6) < 0; _1186_i0 = _1186_i0.Plus(_dafny.One) {
			r1 = (func() _dafny.Sequence {
				if !((_1186_i0).Cmp(_dafny.IntOfUint32((p2).Cardinality())) <= 0) {
					return p2
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("bpeabsjw")
			})()
			var _1187_v1 *C1
			_ = _1187_v1
			var _nw188 *C1 = New_C1_()
			_ = _nw188
			_nw188.Ctor__()
			_1187_v1 = _nw188
			var _1188_v2 _dafny.Array
			_ = _1188_v2
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_34
			var _nw189 _dafny.Array
			_ = _nw189
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw189 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) bool = func(_1189_i1 _dafny.Int) bool {
					return (_this).F6()
				}
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw189 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw189).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw189).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1188_v2 = _nw189
			var _1190_v3 _dafny.Map
			_ = _1190_v3
			_1190_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1187_v1, _1188_v2)
			var _source17 _dafny.Map = _1190_v3
			_ = _source17
			var _1191___mcc_h0 _dafny.Map = _source17
			_ = _1191___mcc_h0
			var _1192_cf18 _dafny.Map = _1191___mcc_h0
			_ = _1192_cf18
			r1 = _dafny.Companion_Sequence_.Concatenate(p2, p2)
			var _1193_v4 _dafny.MultiSet
			_ = _1193_v4
			_1193_v4 = _dafny.MultiSetOf(false)
			var _1194_v5 _dafny.Sequence
			_ = _1194_v5
			_1194_v5 = _dafny.SeqOf(Companion_Default___.Fm0((_this).F6(), globalState))
			(globalState).F1 = ((_1193_v4).Cardinality()).Times(_dafny.IntOfUint32((_1194_v5).Cardinality()))
			r0 = p0
			var _1195_v6 _dafny.MultiSet
			_ = _1195_v6
			_1195_v6 = _dafny.MultiSetOf(_this.F15)
			var _1196_v7 _dafny.Map
			_ = _1196_v7
			_1196_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1186_i0)
			var _1197_v8 _dafny.Map
			_ = _1197_v8
			_1197_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-800)), p1)
			var _1198_v9 _dafny.Map
			_ = _1198_v9
			_1198_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1197_v8, (_this).F6())
			var _1199_v10 _dafny.Map
			_ = _1199_v10
			_1199_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1193_v4).IsProperSubsetOf(_1193_v4), (_1195_v6).IsSubsetOf(((_1195_v6).Update(_dafny.IntOfInt64(-251), Companion_Default___.Abs((_1196_v7).Cardinality()))).Update(p1, Companion_Default___.Abs((_1198_v9).Cardinality()))))
			_1199_v10 = (_1199_v10).Update(p0, p0)
			var _1200_v11 D6
			_ = _1200_v11
			_1200_v11 = Companion_D6_.Create_DC14_((_dafny.IntOfInt64(332)).Cmp(_1186_i0) > 0)
			var _1201_v12 _dafny.Array
			_ = _1201_v12
			var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw190
			_1201_v12 = _nw190
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1201_v12), 0))
			_ = _index211
			(_1201_v12).ArraySet1(p2, (_index211).Int())
			var _1202_v13 _dafny.Array
			_ = _1202_v13
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
			_ = _nw191
			_1202_v13 = _nw191
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1202_v13), 0))
			_ = _index212
			(_1202_v13).ArraySet1(Companion_Default___.Fm43(!(p0), !((_this).F6()), globalState), (_index212).Int())
			var _1203_v14 _dafny.Set
			_ = _1203_v14
			_1203_v14 = _dafny.SetOf(p1, _this.F15, _dafny.IntOfInt64(439), _dafny.IntOfUint32((_1185_v0).Cardinality()), p1)
			var _1204_v15 _dafny.Sequence
			_ = _1204_v15
			_1204_v15 = _dafny.SeqOf((_dafny.SetOf(_this.F15)).Difference(_1203_v14))
			var _1205_v16 _dafny.MultiSet
			_ = _1205_v16
			_1205_v16 = _dafny.MultiSetOf((_this.F15).Times(_this.F15))
			var _1206_v17 _dafny.Map
			_ = _1206_v17
			_1206_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1187_v1).Fm7((_this).F6(), globalState), !((_this).F6()))
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1201_v12), 0))
			_ = _index213
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1202_v13), 0))
			_ = _index214
			var _rhs230 D6 = _1200_v11
			_ = _rhs230
			var _rhs231 _dafny.Int = (func() _dafny.Int {
				if (_1205_v16).Contains((p1).Times(_this.F15)) {
					return (_1205_v16).Multiplicity((p1).Times(_this.F15))
				}
				return _1186_i0
			})()
			_ = _rhs231
			var _rhs232 _dafny.Sequence = (func() _dafny.Sequence {
				if (_1186_i0).Cmp(_1186_i0) != 0 {
					return p2
				}
				return p2
			})()
			_ = _rhs232
			var _rhs233 _dafny.Set = (_1203_v14).Union((_1203_v14).Union(func() _dafny.Set {
				var _coll40 = _dafny.NewBuilder()
				_ = _coll40
				for _iter41 := _dafny.Iterate((_1206_v17).Keys().Elements()); ; {
					_compr_40, _ok41 := _iter41()
					if !_ok41 {
						break
					}
					var _1207_v18 _dafny.Int
					_1207_v18 = interface{}(_compr_40).(_dafny.Int)
					if (_1206_v17).Contains(_1207_v18) {
						_coll40.Add(Companion_Default___.SafeModuloInt(_1207_v18, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(905))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}(func(_1208_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('u')
						}))).Cardinality())))
					}
				}
				return _coll40.ToSet()
			}()))
			_ = _rhs233
			var _rhs234 _dafny.Sequence = _1204_v15
			_ = _rhs234
			var _lhs157 *GlobalState = globalState
			_ = _lhs157
			var _lhs158 _dafny.Array = _1201_v12
			_ = _lhs158
			var _lhs159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1201_v12), 0))
			_ = _lhs159
			var _lhs160 _dafny.Array = _1202_v13
			_ = _lhs160
			var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.ArrayLen((_1202_v13), 0))
			_ = _lhs161
			_1200_v11 = _rhs230
			_lhs157.F1 = _rhs231
			(_lhs158).ArraySet1(_rhs232, (_lhs159).Int())
			(_lhs160).ArraySet1(_rhs233, (_lhs161).Int())
			_1204_v15 = _rhs234
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1188_v2), 0))
			_ = _index215
			(_1188_v2).ArraySet1(p0, (_index215).Int())
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(480), _dafny.ArrayLen((_1188_v2), 0))
			_ = _index216
			(_1188_v2).ArraySet1((func() bool {
				if ((_this).F6()) || ((_this).F6()) {
					return !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-713))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg74 _dafny.Int) interface{} {
							return coer74(arg74)
						}
					}(func(_1209_i3 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					})), (_1201_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(153), _dafny.ArrayLen((_1201_v12), 0))).Int()).(_dafny.Sequence))
				}
				return !(!(p0) || (false))
			})(), (_index216).Int())
		}
		(_this).F15 = p1
		var _1210_v19 *C3
		_ = _1210_v19
		var _nw192 *C3 = New_C3_()
		_ = _nw192
		_nw192.Ctor__(p0, (_this).Fm6(_1185_v0, globalState))
		_1210_v19 = _nw192
		(_this).F15 = (_1185_v0).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1185_v0).Cardinality()))).Uint32()).(_dafny.Int)
		var _1211_v20 _dafny.MultiSet
		_ = _1211_v20
		_1211_v20 = _dafny.MultiSetOf((_1210_v19).F13(), false, p0, (_1210_v19).F14(), (_1210_v19).F14())
		r0 = (_dafny.MultiSetOf((_1210_v19).F13())).IsProperSubsetOf((func() _dafny.MultiSet {
			if (_this).F6() {
				return _1211_v20
			}
			return _1211_v20
		})())
		var _1212_v21 D0
		_ = _1212_v21
		_1212_v21 = Companion_D0_.Create_DC1_()
		var _1213_v22 D0
		_ = _1213_v22
		_1213_v22 = Companion_D0_.Create_DC3_(_1212_v21)
		var _1214_v23 _dafny.Map
		_ = _1214_v23
		_1214_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1213_v22, (_dafny.Zero).Minus(p1))
		var _1215_v24 _dafny.Sequence
		_ = _1215_v24
		_1215_v24 = _dafny.SeqOf(_1214_v23, _1214_v23, _1214_v23, _1214_v23)
		var _1216_v25 _dafny.Array
		_ = _1216_v25
		var _nwElement0_33 _dafny.Sequence = _1215_v24
		_ = _nwElement0_33
		var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(3))
		_ = _nw193
		(_nw193).ArraySet1(_nwElement0_33, 0)
		(_nw193).ArraySet1(_1215_v24, 1)
		(_nw193).ArraySet1(_1215_v24, 2)
		_1216_v25 = _nw193
		var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1216_v25), 0))
		_ = _index217
		(_1216_v25).ArraySet1(_dafny.SeqOf(_1214_v23), (_index217).Int())
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1216_v25), 0))
		_ = _index218
		(_1216_v25).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1215_v24, _1215_v24), (_index218).Int())
		r0 = (_this).F6()
		r1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p2, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F5()), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F5())).Cardinality()))).Uint32(), (_this).F5())), _dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), (_this).F5()))
		return r0, r1
	}
}
func (_this *C5) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _source18 D0 = Companion_D0_.Create_DC1_()
		_ = _source18
		if _source18.Is_DC1() {
			var _1217_v0 bool
			_ = _1217_v0
			_1217_v0 = false
			var _1218_v1 _dafny.Sequence
			_ = _1218_v1
			_1218_v1 = _dafny.SeqOf((_this).F6(), false)
			_1217_v0 = _dafny.Companion_Sequence_.IsPrefixOf(_1218_v1, _1218_v1)
			var _1219_v2 _dafny.Map
			_ = _1219_v2
			_1219_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this.F15).Cmp(_dafny.IntOfInt64(455)) < 0)
			_1219_v2 = (_1219_v2).Update((_this).F6(), (_this).F6())
			var _1220_v3 _dafny.Array
			_ = _1220_v3
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw194
			_1220_v3 = _nw194
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1220_v3), 0))
			_ = _index219
			(_1220_v3).ArraySet1((_1217_v0) || (!(false)), (_index219).Int())
			var _1221_v4 _dafny.MultiSet
			_ = _1221_v4
			_1221_v4 = _dafny.MultiSetOf(_1217_v0, _1217_v0)
			var _1222_v5 D15
			_ = _1222_v5
			_1222_v5 = Companion_D15_.Create_DC27_(_1221_v4)
			var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1220_v3), 0))
			_ = _index220
			var _rhs235 _dafny.MultiSet = ((_1221_v4).Union(_1221_v4)).Intersection((_1222_v5).Dtor_cf35())
			_ = _rhs235
			var _rhs236 bool = ((_dafny.Zero).Minus(_this.F15)).Cmp(_this.F15) != 0
			_ = _rhs236
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			var _lhs163 _dafny.Array = _1220_v3
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1220_v3), 0))
			_ = _lhs164
			_lhs162.F4 = _rhs235
			(_lhs163).ArraySet1(_rhs236, (_lhs164).Int())
			var _1223_v6 _dafny.MultiSet
			_ = _1223_v6
			_1223_v6 = _dafny.MultiSetOf(_this.F15, _this.F15)
			var _1224_v7 _dafny.Map
			_ = _1224_v7
			_1224_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wl")).Cardinality()), _1223_v6)
			var _1225_v8 _dafny.Sequence
			_ = _1225_v8
			_1225_v8 = _dafny.SeqOf(_dafny.IntOfInt64(758))
			var _1226_v9 _dafny.Sequence
			_ = _1226_v9
			_1226_v9 = _dafny.SeqOf(_1223_v6)
			var _1227_v10 _dafny.Array
			_ = _1227_v10
			var _nwElement0_34 _dafny.MultiSet = _dafny.MultiSetOf((Companion_Default___.Fm44(!((_1220_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_1220_v3), 0))).Int()).(bool)), globalState)).Cardinality())
			_ = _nwElement0_34
			var _nw195 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(14))
			_ = _nw195
			(_nw195).ArraySet1(_nwElement0_34, 0)
			(_nw195).ArraySet1(Companion_Default___.Fm45(globalState), 1)
			(_nw195).ArraySet1((func() _dafny.MultiSet {
				if (_1224_v7).Contains((Companion_Default___.Fm46((_this).Fm6(_1225_v8, globalState), globalState)).Cardinality()) {
					return (_1224_v7).Get((Companion_Default___.Fm46((_this).Fm6(_1225_v8, globalState), globalState)).Cardinality()).(_dafny.MultiSet)
				}
				return _dafny.MultiSetOf(_this.F15)
			})(), 2)
			(_nw195).ArraySet1((_1223_v6).Union((_1223_v6).Update(_this.F15, Companion_Default___.Abs(_this.F15))), 3)
			(_nw195).ArraySet1(_1223_v6, 4)
			(_nw195).ArraySet1(_1223_v6, 5)
			(_nw195).ArraySet1(_1223_v6, 6)
			(_nw195).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_1225_v8).Cardinality())), 7)
			(_nw195).ArraySet1(_1223_v6, 8)
			(_nw195).ArraySet1(_1223_v6, 9)
			(_nw195).ArraySet1(_1223_v6, 10)
			(_nw195).ArraySet1(_1223_v6, 11)
			(_nw195).ArraySet1((_1226_v9).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1226_v9).Cardinality()))).Uint32()).(_dafny.MultiSet), 12)
			(_nw195).ArraySet1(((_1223_v6).Update(_this.F15, Companion_Default___.Abs(_this.F15))).Intersection(_1223_v6), 13)
			_1227_v10 = _nw195
			var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1227_v10), 0))
			_ = _index221
			(_1227_v10).ArraySet1(_1223_v6, (_index221).Int())
			var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(9), _dafny.ArrayLen((_1227_v10), 0))
			_ = _index222
			(_1227_v10).ArraySet1(_1223_v6, (_index222).Int())
		} else if _source18.Is_DC2() {
			var _1228___mcc_h0 _dafny.Map = _source18.Get_().(D0_DC2).Cf1
			_ = _1228___mcc_h0
			var _1229___mcc_h1 _dafny.CodePoint = _source18.Get_().(D0_DC2).Cf2
			_ = _1229___mcc_h1
			var _1230_cf2 _dafny.CodePoint = _1229___mcc_h1
			_ = _1230_cf2
			var _1231_cf1 _dafny.Map = _1228___mcc_h0
			_ = _1231_cf1
			var _1232_v11 bool
			_ = _1232_v11
			_1232_v11 = false
			_1232_v11 = (_this).F6()
			r0 = (func() _dafny.Int {
				if (func() bool {
					if (_this).F6() {
						return (_this).F6()
					}
					return false
				})() {
					return _this.F15
				}
				return _this.F15
			})()
			var _1233_v12 _dafny.MultiSet
			_ = _1233_v12
			_1233_v12 = _dafny.MultiSetOf(_1232_v11)
			var _source19 D15 = Companion_D15_.Create_DC27_(_1233_v12)
			_ = _source19
			if _source19.Is_DC28() {
				var _1234___mcc_h4 bool = _source19.Get_().(D15_DC28).Cf36
				_ = _1234___mcc_h4
				var _1235___mcc_h5 bool = _source19.Get_().(D15_DC28).Cf37
				_ = _1235___mcc_h5
				var _1236_cf37 bool = _1235___mcc_h5
				_ = _1236_cf37
				var _1237_cf36 bool = _1234___mcc_h4
				_ = _1237_cf36
				(globalState).F1 = (_this).Fm7((_this).F6(), globalState)
				var _1238_v13 _dafny.Array
				_ = _1238_v13
				var _nw196 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw196
				_1238_v13 = _nw196
				var _1239_v14 _dafny.Sequence
				_ = _1239_v14
				_1239_v14 = _dafny.UnicodeSeqOfUtf8Bytes("g")
				var _1240_v15 _dafny.MultiSet
				_ = _1240_v15
				_1240_v15 = _dafny.MultiSetOf(_1239_v14, _1239_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(969))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1241_cf2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1242_i0 _dafny.Int) _dafny.CodePoint {
						return _1241_cf2
					}
				})(_1230_cf2))), _dafny.UnicodeSeqOfUtf8Bytes("bjrtxt"), _1239_v14)
				var _1243_v17 _dafny.Map
				_ = _1243_v17
				_1243_v17 = func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(717), _dafny.IntOfInt64(-196))); ; {
						_compr_41, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _1244_v16 _dafny.Int
						_1244_v16 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(717)).Cmp(_1244_v16) <= 0) && ((_1244_v16).Cmp(_dafny.IntOfInt64(-196)) < 0) {
							_coll41.Add(Companion_Default___.SafeDivisionInt(_1244_v16, (_dafny.MultiSetOf(_dafny.SeqOf(true, (_this).F6()))).Cardinality()), _dafny.IntOfInt64(868))
						}
					}
					return _coll41.ToMap()
				}()
				var _1245_v18 D12
				_ = _1245_v18
				_1245_v18 = Companion_D12_.Create_DC22_(_1243_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg76 _dafny.Int) interface{} {
						return coer76(arg76)
					}
				}((func(_1246_v11 bool, _1247_cf37 bool) func(_dafny.Int) _dafny.Sequence {
					return func(_1248_i1 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqOf(_1246_v11, _1247_cf37)
					}
				})(_1232_v11, _1236_cf37))), _1230_cf2, _1232_v11, _1239_v14)
				var _nwElement0_35 bool = (_this).F6()
				_ = _nwElement0_35
				var _nw197 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(19))
				_ = _nw197
				(_nw197).ArraySet1(_nwElement0_35, 0)
				(_nw197).ArraySet1((_1240_v15).IsSubsetOf(_1240_v15), 1)
				(_nw197).ArraySet1(_1236_cf37, 2)
				(_nw197).ArraySet1((func() bool {
					if _1232_v11 {
						return !(_1232_v11)
					}
					return _1237_cf36
				})(), 3)
				(_nw197).ArraySet1(_1236_cf37, 4)
				(_nw197).ArraySet1(false, 5)
				(_nw197).ArraySet1(_1237_cf36, 6)
				(_nw197).ArraySet1(_1237_cf36, 7)
				(_nw197).ArraySet1(_1236_cf37, 8)
				(_nw197).ArraySet1(_1232_v11, 9)
				(_nw197).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1239_v14, _dafny.Companion_Sequence_.Update(_1239_v14, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1239_v14).Cardinality()))).Uint32(), (_1245_v18).Dtor_cf26())), 10)
				(_nw197).ArraySet1((_dafny.IntOfInt64(440)).Cmp(_this.F15) > 0, 11)
				(_nw197).ArraySet1(_1232_v11, 12)
				(_nw197).ArraySet1(_1232_v11, 13)
				(_nw197).ArraySet1(_1236_cf37, 14)
				(_nw197).ArraySet1(_1237_cf36, 15)
				(_nw197).ArraySet1(_1237_cf36, 16)
				(_nw197).ArraySet1(!(!(_1236_cf37)), 17)
				(_nw197).ArraySet1(_1232_v11, 18)
				_1238_v13 = _nw197
				(globalState).F1 = (_this.F15).Times(_this.F15)
				var _1249_v19 D11
				_ = _1249_v19
				_1249_v19 = Companion_D11_.Create_DC20_((func() bool {
					if (_this).F6() {
						return _1236_cf37
					}
					return _1232_v11
				})(), (_this.F15).Minus(_this.F15))
				_1249_v19 = _1249_v19
			} else if _source19.Is_DC27() {
				var _1250___mcc_h6 _dafny.MultiSet = _source19.Get_().(D15_DC27).Cf35
				_ = _1250___mcc_h6
				var _1251_cf35 _dafny.MultiSet = _1250___mcc_h6
				_ = _1251_cf35
				var _1252_v20 _dafny.Map
				_ = _1252_v20
				_1252_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
				var _1253_v21 _dafny.Map
				_ = _1253_v21
				_1253_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1252_v20)
				var _1254_v22 _dafny.Map
				_ = _1254_v22
				_1254_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1232_v11, Companion_Default___.Fm47((_this).F6(), globalState))
				(globalState).F1 = ((_1253_v21).Merge((func() _dafny.Map {
					if (_1254_v22).Contains((_this).F6()) {
						return (_1254_v22).Get((_this).F6()).(_dafny.Map)
					}
					return _1253_v21
				})())).Cardinality()
				(_this).F15 = _this.F15
				var _1255_v23 _dafny.Array
				_ = _1255_v23
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw198
				_1255_v23 = _nw198
				var _1256_v24 _dafny.Array
				_ = _1256_v24
				_1256_v24 = _1255_v23
				var _1257_v25 _dafny.Array
				_ = _1257_v25
				var _nwElement0_36 _dafny.Array = _1255_v23
				_ = _nwElement0_36
				var _nw199 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(28))
				_ = _nw199
				(_nw199).ArraySet1(_nwElement0_36, 0)
				(_nw199).ArraySet1((_1256_v24), 1)
				(_nw199).ArraySet1((func() _dafny.Array {
					if _1232_v11 {
						return _1255_v23
					}
					return _1255_v23
				})(), 2)
				(_nw199).ArraySet1(_1255_v23, 3)
				(_nw199).ArraySet1(_1255_v23, 4)
				(_nw199).ArraySet1((_1256_v24), 5)
				(_nw199).ArraySet1(_1255_v23, 6)
				(_nw199).ArraySet1(_1255_v23, 7)
				(_nw199).ArraySet1(_1255_v23, 8)
				(_nw199).ArraySet1(_1255_v23, 9)
				(_nw199).ArraySet1(_1255_v23, 10)
				(_nw199).ArraySet1(_1255_v23, 11)
				(_nw199).ArraySet1(_1255_v23, 12)
				(_nw199).ArraySet1(_1255_v23, 13)
				(_nw199).ArraySet1(_1255_v23, 14)
				(_nw199).ArraySet1(_1255_v23, 15)
				(_nw199).ArraySet1(_1255_v23, 16)
				(_nw199).ArraySet1(_1255_v23, 17)
				(_nw199).ArraySet1(_1255_v23, 18)
				(_nw199).ArraySet1(_1255_v23, 19)
				(_nw199).ArraySet1(_1255_v23, 20)
				(_nw199).ArraySet1(_1255_v23, 21)
				(_nw199).ArraySet1(_1255_v23, 22)
				(_nw199).ArraySet1(_1255_v23, 23)
				(_nw199).ArraySet1(_1255_v23, 24)
				(_nw199).ArraySet1(_1255_v23, 25)
				(_nw199).ArraySet1(_1255_v23, 26)
				(_nw199).ArraySet1(_1255_v23, 27)
				_1257_v25 = _nw199
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1257_v25), 0))
				_ = _index223
				(_1257_v25).ArraySet1((func() _dafny.Array {
					if _1232_v11 {
						return _1255_v23
					}
					return _1255_v23
				})(), (_index223).Int())
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_1257_v25), 0))
				_ = _index224
				(_1257_v25).ArraySet1(_1255_v23, (_index224).Int())
				(globalState).F1 = (func() _dafny.Int {
					if (_this).F6() {
						return _this.F15
					}
					return _this.F15
				})()
			} else {
				var _1258___mcc_h7 D15 = _source19.Get_().(D15_DC29).Cf38
				_ = _1258___mcc_h7
				var _1259_cf38 D15 = _1258___mcc_h7
				_ = _1259_cf38
				_1232_v11 = (_this).F6()
				var _1260_v26 _dafny.Map
				_ = _1260_v26
				_1260_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_1233_v12).Cardinality())
				var _1261_v27 _dafny.Array
				_ = _1261_v27
				var _len0_35 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_35
				var _nw200 _dafny.Array
				_ = _nw200
				if _len0_35.Cmp(_dafny.Zero) == 0 {
					_nw200 = _dafny.NewArray(_len0_35)
				} else {
					var _init35 func(_dafny.Int) _dafny.Int = func(_1262_i2 _dafny.Int) _dafny.Int {
						return (_1262_i2).Plus(_this.F15)
					}
					_ = _init35
					var _element0_35 = _init35(_dafny.Zero)
					_ = _element0_35
					_nw200 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
					(_nw200).ArraySet1(_element0_35, 0)
					var _nativeLen0_35 = (_len0_35).Int()
					_ = _nativeLen0_35
					for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
						(_nw200).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
					}
				}
				_1261_v27 = _nw200
				var _1263_v28 _dafny.Map
				_ = _1263_v28
				_1263_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1232_v11, (_dafny.Zero).Minus(_this.F15))).Merge(_1260_v26), _1261_v27)
				var _1264_v29 _dafny.Set
				_ = _1264_v29
				_1264_v29 = _dafny.SetOf(_1260_v26, _1260_v26, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F15)).Merge(_1260_v26), _1260_v26, (_1260_v26).Merge(_1260_v26))
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1261_v27), 0))
				_ = _index225
				(_1261_v27).ArraySet1(_this.F15, (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1261_v27), 0))
				_ = _index226
				var _rhs237 _dafny.Int = _this.F15
				_ = _rhs237
				var _rhs238 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1260_v26, _1261_v27)).Merge(_1263_v28)
				_ = _rhs238
				var _rhs239 _dafny.Set = _1264_v29
				_ = _rhs239
				var _rhs240 _dafny.Int = _this.F15
				_ = _rhs240
				var _rhs241 _dafny.Int = _dafny.IntOfInt64(459)
				_ = _rhs241
				var _lhs165 *GlobalState = globalState
				_ = _lhs165
				var _lhs166 _dafny.Array = _1261_v27
				_ = _lhs166
				var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1261_v27), 0))
				_ = _lhs167
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				_lhs165.F1 = _rhs237
				_1263_v28 = _rhs238
				_1264_v29 = _rhs239
				(_lhs166).ArraySet1(_rhs240, (_lhs167).Int())
				_lhs168.F1 = _rhs241
				var _1265_v30 _dafny.Array
				_ = _1265_v30
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_36
				var _nw201 _dafny.Array
				_ = _nw201
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw201 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) bool = func(_1266_i3 _dafny.Int) bool {
						return (_this).F6()
					}
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw201 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw201).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw201).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1265_v30 = _nw201
				var _1267_v31 D0
				_ = _1267_v31
				var _1268_v32 bool
				_ = _1268_v32
				var _out34 D0
				_ = _out34
				var _out35 bool
				_ = _out35
				_out34, _out35 = Companion_Default___.M0(_1265_v30, (_this).F5(), !((_this).F6()), _1230_cf2, globalState)
				_1267_v31 = _out34
				_1268_v32 = _out35
				var _1269_v33 _dafny.Map
				_ = _1269_v33
				_1269_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((Companion_Default___.Fm48((_1261_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1261_v27), 0))).Int()).(_dafny.Int), true, globalState)).Cardinality()), _1268_v32)
				var _1270_v34 _dafny.Map
				_ = _1270_v34
				_1270_v34 = _1269_v33
				var _1271_v35 _dafny.Map
				_ = _1271_v35
				_1271_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v32, _1270_v34)
				_1271_v35 = (_1271_v35).Update((func() bool {
					if _1268_v32 {
						return (_this).F6()
					}
					return _1232_v11
				})(), _1270_v34)
			}
			var _1272_v36 _dafny.Sequence
			_ = _1272_v36
			_1272_v36 = _dafny.SeqOf(_this.F15)
			var _1273_v37 _dafny.Map
			_ = _1273_v37
			_1273_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D14_.Create_DC26_(_1232_v11, _dafny.IntOfUint32((_1272_v36).Cardinality()), _1230_cf2), (_this).F5())
			_1232_v11 = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1273_v37)).Equals(func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter43 := _dafny.Iterate((_1272_v36).Elements()); ; {
					_compr_42, _ok43 := _iter43()
					if !_ok43 {
						break
					}
					var _1274_v38 _dafny.Int
					_1274_v38 = interface{}(_compr_42).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1272_v36, _1274_v38) {
						_coll42.Add((_1274_v38).Plus(_this.F15), (func() _dafny.Map {
							var _coll43 = _dafny.NewMapBuilder()
							_ = _coll43
							for _iter44 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-187))).Uint32(), func(coer77 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
								return func(arg77 _dafny.Int) interface{} {
									return coer77(arg77)
								}
							}((func(_1275_v11 bool, _1276_cf2 _dafny.CodePoint) func(_dafny.Int) D14 {
								return func(_1277_i4 _dafny.Int) D14 {
									return Companion_D14_.Create_DC26_(_1275_v11, _this.F15, _1276_cf2)
								}
							})(_1232_v11, _1230_cf2)))).Elements()); ; {
								_compr_43, _ok44 := _iter44()
								if !_ok44 {
									break
								}
								var _1278_v39 D14
								_1278_v39 = interface{}(_compr_43).(D14)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-187))).Uint32(), func(coer78 func(_dafny.Int) D14) func(_dafny.Int) interface{} {
									return func(arg78 _dafny.Int) interface{} {
										return coer78(arg78)
									}
								}((func(_1279_v11 bool, _1280_cf2 _dafny.CodePoint) func(_dafny.Int) D14 {
									return func(_1277_i4 _dafny.Int) D14 {
										return Companion_D14_.Create_DC26_(_1279_v11, _this.F15, _1280_cf2)
									}
								})(_1232_v11, _1230_cf2))), _1278_v39) {
									_coll43.Add(_1278_v39, _dafny.CodePoint('q'))
								}
							}
							return _coll43.ToMap()
						}()).Update(Companion_D14_.Create_DC26_(_1232_v11, _this.F15, (_this).F5()), _1230_cf2))
					}
				}
				return _coll42.ToMap()
			}())
		} else if _source18.Is_DC0() {
			var _1281___mcc_h2 bool = _source18.Get_().(D0_DC0).Cf0
			_ = _1281___mcc_h2
			var _1282_cf0 bool = _1281___mcc_h2
			_ = _1282_cf0
			var _1283_v40 _dafny.Array
			_ = _1283_v40
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_37
			var _nw202 _dafny.Array
			_ = _nw202
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw202 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = func(_1284_i5 _dafny.Int) _dafny.Int {
					return (_1284_i5).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), (Companion_D6_.Create_DC14_((_this).F6())).Dtor_cf15())).Cardinality())
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw202 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw202).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw202).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1283_v40 = _nw202
			var _1285_v41 _dafny.Sequence
			_ = _1285_v41
			_1285_v41 = _dafny.UnicodeSeqOfUtf8Bytes("epvurt")
			var _1286_v42 _dafny.Sequence
			_ = _1286_v42
			_1286_v42 = _dafny.SeqOf(_dafny.IntOfUint32((_1285_v41).Cardinality()))
			var _1287_v43 _dafny.Sequence
			_ = _1287_v43
			_1287_v43 = _dafny.SeqOf(_dafny.IntOfUint32((_1286_v42).Cardinality()))
			var _1288_v44 _dafny.Array
			_ = _1288_v44
			var _nwElement0_37 bool = _1282_cf0
			_ = _nwElement0_37
			var _nw203 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(3))
			_ = _nw203
			(_nw203).ArraySet1(_nwElement0_37, 0)
			(_nw203).ArraySet1((_this).Fm6(_1287_v43, globalState), 1)
			(_nw203).ArraySet1(_1282_cf0, 2)
			_1288_v44 = _nw203
			var _1289_v45 _dafny.Set
			_ = _1289_v45
			_1289_v45 = _dafny.SetOf(_dafny.IntOfInt64(797), _dafny.IntOfInt64(526), (Companion_Default___.Fm49(globalState)).Cardinality())
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1288_v44), 0))
			_ = _index227
			(_1288_v44).ArraySet1((_1289_v45).Contains((_1289_v45).Cardinality()), (_index227).Int())
			var _1290_v46 _dafny.Sequence
			_ = _1290_v46
			_1290_v46 = _dafny.SeqOf(_1282_cf0, (_this).F6())
			var _1291_v47 _dafny.Sequence
			_ = _1291_v47
			_1291_v47 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_1290_v46).Cardinality()), _dafny.IntOfInt64(53), _this.F15, _this.F15))
			var _1292_v48 D6
			_ = _1292_v48
			_1292_v48 = Companion_D6_.Create_DC13_(_1286_v42)
			var _1293_v49 _dafny.Map
			_ = _1293_v49
			_1293_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1289_v45).Cardinality(), _1282_cf0)
			var _1294_v50 _dafny.Array
			_ = _1294_v50
			var _nwElement0_38 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1286_v42, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-560), _dafny.IntOfUint32((_1286_v42).Cardinality()))).Uint32(), _dafny.IntOfInt64(501)), _1286_v42)
			_ = _nwElement0_38
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(22))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_38, 0)
			(_nw204).ArraySet1(_1286_v42, 1)
			(_nw204).ArraySet1(_1287_v43, 2)
			(_nw204).ArraySet1((_1291_v47).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1291_v47).Cardinality()))).Uint32()).(_dafny.Sequence), 3)
			(_nw204).ArraySet1(_1287_v43, 4)
			(_nw204).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer79 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg79 _dafny.Int) interface{} {
					return coer79(arg79)
				}
			}(func(_1295_i6 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(189)
			})), 5)
			(_nw204).ArraySet1(_dafny.SeqOf(_this.F15, _dafny.IntOfUint32((_1285_v41).Cardinality()), _this.F15), 6)
			(_nw204).ArraySet1(_1286_v42, 7)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Update(_1287_v43, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(700), _dafny.IntOfUint32((_1287_v43).Cardinality()))).Uint32(), _this.F15), 8)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F15, _this.F15), (_1292_v48).Dtor_cf14()), 9)
			(_nw204).ArraySet1(_1286_v42, 10)
			(_nw204).ArraySet1(_1287_v43, 11)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1293_v49).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-2))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg80 _dafny.Int) interface{} {
					return coer80(arg80)
				}
			}(func(_1296_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(-85)
			}))), 12)
			(_nw204).ArraySet1(_1287_v43, 13)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1286_v42, _1286_v42), 14)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Update(_1286_v42, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F15), _dafny.IntOfUint32((_1286_v42).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_cf0, _this.F15)).Cardinality()), 15)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer81 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1297_i8 _dafny.Int) _dafny.Int {
				return _this.F15
			})), _1286_v42), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F15), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg82 _dafny.Int) interface{} {
					return coer82(arg82)
				}
			}(func(_1298_i8 _dafny.Int) _dafny.Int {
				return _this.F15
			})), _1286_v42)).Cardinality()))).Uint32(), _this.F15), 16)
			(_nw204).ArraySet1(_1287_v43, 17)
			(_nw204).ArraySet1(_1286_v42, 18)
			(_nw204).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1287_v43, _1287_v43), 19)
			(_nw204).ArraySet1(_1286_v42, 20)
			(_nw204).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(781))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg83 _dafny.Int) interface{} {
					return coer83(arg83)
				}
			}(func(_1299_i9 _dafny.Int) _dafny.Int {
				return _this.F15
			})), 21)
			_1294_v50 = _nw204
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1294_v50), 0))
			_ = _index228
			(_1294_v50).ArraySet1(_dafny.SeqOf(_this.F15, _this.F15), (_index228).Int())
			var _1300_v51 _dafny.MultiSet
			_ = _1300_v51
			_1300_v51 = _dafny.MultiSetOf((_this).F6())
			var _1301_v52 _dafny.Map
			_ = _1301_v52
			_1301_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1300_v51)
			var _1302_v53 _dafny.Sequence
			_ = _1302_v53
			_1302_v53 = _dafny.SeqOf(_1300_v51)
			var _1303_v54 _dafny.Array
			_ = _1303_v54
			var _nwElement0_39 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_1290_v46)).Union((_1300_v51).Update(true, Companion_Default___.Abs(_this.F15)))
			_ = _nwElement0_39
			var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(25))
			_ = _nw205
			(_nw205).ArraySet1(_nwElement0_39, 0)
			(_nw205).ArraySet1(_1300_v51, 1)
			(_nw205).ArraySet1((_1300_v51).Difference(_1300_v51), 2)
			(_nw205).ArraySet1((_1300_v51).Intersection(_1300_v51), 3)
			(_nw205).ArraySet1(_1300_v51, 4)
			(_nw205).ArraySet1(_1300_v51, 5)
			(_nw205).ArraySet1(_1300_v51, 6)
			(_nw205).ArraySet1(_1300_v51, 7)
			(_nw205).ArraySet1((_1300_v51).Difference(_1300_v51), 8)
			(_nw205).ArraySet1(_1300_v51, 9)
			(_nw205).ArraySet1((_1300_v51).Intersection(Companion_Default___.Fm48(_this.F15, (_this).F6(), globalState)), 10)
			(_nw205).ArraySet1((_1300_v51).Union(_1300_v51), 11)
			(_nw205).ArraySet1(_1300_v51, 12)
			(_nw205).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(!(true)))), 13)
			(_nw205).ArraySet1((_1300_v51).Union(_1300_v51), 14)
			(_nw205).ArraySet1(_1300_v51, 15)
			(_nw205).ArraySet1((func() _dafny.MultiSet {
				if (_1301_v52).Contains((_this).Fm7(_1282_cf0, globalState)) {
					return (_1301_v52).Get((_this).Fm7(_1282_cf0, globalState)).(_dafny.MultiSet)
				}
				return _1300_v51
			})(), 16)
			(_nw205).ArraySet1((_1302_v53).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1302_v53).Cardinality()))).Uint32()).(_dafny.MultiSet), 17)
			(_nw205).ArraySet1((_1300_v51).Difference(_1300_v51), 18)
			(_nw205).ArraySet1((_dafny.MultiSetFromSeq(_1290_v46)).Update((_1290_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.IntOfUint32((_1290_v46).Cardinality()))).Uint32()).(bool), Companion_Default___.Abs(_this.F15)), 19)
			(_nw205).ArraySet1((_dafny.MultiSetOf(_1282_cf0, _1282_cf0)).Union(_1300_v51), 20)
			(_nw205).ArraySet1(_dafny.MultiSetOf(_1282_cf0), 21)
			(_nw205).ArraySet1(_1300_v51, 22)
			(_nw205).ArraySet1(_1300_v51, 23)
			(_nw205).ArraySet1((_1300_v51).Difference(_1300_v51), 24)
			_1303_v54 = _nw205
			var _1304_v55 _dafny.Array
			_ = _1304_v55
			var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw206
			_1304_v55 = _nw206
			var _1305_v57 _dafny.MultiSet
			_ = _1305_v57
			_1305_v57 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1285_v41).Cardinality()))
			var _1306_v59 _dafny.Array
			_ = _1306_v59
			var _nwElement0_40 _dafny.Map = _1293_v49
			_ = _nwElement0_40
			var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(28))
			_ = _nw207
			(_nw207).ArraySet1(_nwElement0_40, 0)
			(_nw207).ArraySet1(_1293_v49, 1)
			(_nw207).ArraySet1(_1293_v49, 2)
			(_nw207).ArraySet1(_1293_v49, 3)
			(_nw207).ArraySet1(_1293_v49, 4)
			(_nw207).ArraySet1(_1293_v49, 5)
			(_nw207).ArraySet1(_1293_v49, 6)
			(_nw207).ArraySet1(_1293_v49, 7)
			(_nw207).ArraySet1(func() _dafny.Map {
				var _coll44 = _dafny.NewMapBuilder()
				_ = _coll44
				for _iter45 := _dafny.Iterate((_1305_v57).Elements()); ; {
					_compr_44, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _1307_v56 _dafny.Int
					_1307_v56 = interface{}(_compr_44).(_dafny.Int)
					if (_1305_v57).Contains(_1307_v56) {
						_coll44.Add((_1307_v56).Plus(_dafny.IntOfUint32((_1290_v46).Cardinality())), _1282_cf0)
					}
				}
				return _coll44.ToMap()
			}(), 8)
			(_nw207).ArraySet1(_1293_v49, 9)
			(_nw207).ArraySet1(_1293_v49, 10)
			(_nw207).ArraySet1(_1293_v49, 11)
			(_nw207).ArraySet1(_1293_v49, 12)
			(_nw207).ArraySet1(_1293_v49, 13)
			(_nw207).ArraySet1(_1293_v49, 14)
			(_nw207).ArraySet1(func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-991), _dafny.IntOfInt64(889))); ; {
					_compr_45, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _1308_v58 _dafny.Int
					_1308_v58 = interface{}(_compr_45).(_dafny.Int)
					if ((_dafny.IntOfInt64(-991)).Cmp(_1308_v58) <= 0) && ((_1308_v58).Cmp(_dafny.IntOfInt64(889)) < 0) {
						_coll45.Add((_1308_v58).Plus(_this.F15), (Companion_D14_.Create_DC26_(false, _this.F15, (_this).F5())).Dtor_cf32())
					}
				}
				return _coll45.ToMap()
			}(), 15)
			(_nw207).ArraySet1(_1293_v49, 16)
			(_nw207).ArraySet1((_1293_v49).Update(_this.F15, (_this).F6()), 17)
			(_nw207).ArraySet1(_1293_v49, 18)
			(_nw207).ArraySet1(_1293_v49, 19)
			(_nw207).ArraySet1(_1293_v49, 20)
			(_nw207).ArraySet1(_1293_v49, 21)
			(_nw207).ArraySet1(_1293_v49, 22)
			(_nw207).ArraySet1(_1293_v49, 23)
			(_nw207).ArraySet1(_1293_v49, 24)
			(_nw207).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (_this).F6()), 25)
			(_nw207).ArraySet1(_1293_v49, 26)
			(_nw207).ArraySet1(_1293_v49, 27)
			_1306_v59 = _nw207
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_1304_v55), 0))
			_ = _index229
			(_1304_v55).ArraySet1(_1306_v59, (_index229).Int())
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1288_v44), 0))
			_ = _index230
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1294_v50), 0))
			_ = _index231
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_1304_v55), 0))
			_ = _index232
			var _rhs242 _dafny.Array = _1283_v40
			_ = _rhs242
			var _rhs243 bool = !((_1289_v45).IsSubsetOf((func() _dafny.Set {
				if !(!(true)) {
					return _1289_v45
				}
				return _1289_v45
			})()))
			_ = _rhs243
			var _rhs244 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1287_v43, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1287_v43).Cardinality()))).Uint32(), _dafny.IntOfInt64(777))
			_ = _rhs244
			var _rhs245 _dafny.Array = _1303_v54
			_ = _rhs245
			var _rhs246 _dafny.Array = _1306_v59
			_ = _rhs246
			var _lhs169 _dafny.Array = _1288_v44
			_ = _lhs169
			var _lhs170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1288_v44), 0))
			_ = _lhs170
			var _lhs171 _dafny.Array = _1294_v50
			_ = _lhs171
			var _lhs172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_1294_v50), 0))
			_ = _lhs172
			var _lhs173 _dafny.Array = _1304_v55
			_ = _lhs173
			var _lhs174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(693), _dafny.ArrayLen((_1304_v55), 0))
			_ = _lhs174
			_1283_v40 = _rhs242
			(_lhs169).ArraySet1(_rhs243, (_lhs170).Int())
			(_lhs171).ArraySet1(_rhs244, (_lhs172).Int())
			_1303_v54 = _rhs245
			(_lhs173).ArraySet1(_rhs246, (_lhs174).Int())
			var _1309_v60 _dafny.Map
			_ = _1309_v60
			_1309_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("erhqhi"), _1285_v41)
			_1309_v60 = _1309_v60
			var _1310_v61 _dafny.Map
			_ = _1310_v61
			_1310_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1288_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1288_v44), 0))).Int()).(bool), _1283_v40)
			var _1311_v62 _dafny.Array
			_ = _1311_v62
			_1311_v62 = (func() _dafny.Array {
				if !(false) {
					return _1283_v40
				}
				return (func() _dafny.Array {
					if (_1310_v61).Contains((_1290_v46).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1290_v46).Cardinality()))).Uint32()).(bool)) {
						return (_1310_v61).Get((_1290_v46).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1290_v46).Cardinality()))).Uint32()).(bool)).(_dafny.Array)
					}
					return _1283_v40
				})()
			})()
			var _rhs247 bool = !(_1289_v45).Equals(_1289_v45)
			_ = _rhs247
			var _rhs248 _dafny.Set = _1289_v45
			_ = _rhs248
			var _rhs249 _dafny.Array = _1311_v62
			_ = _rhs249
			_1282_cf0 = _rhs247
			_1289_v45 = _rhs248
			_1311_v62 = _rhs249
			(_this).F15 = _dafny.IntOfInt64(75)
		} else {
			var _1312___mcc_h3 D0 = _source18.Get_().(D0_DC3).Cf3
			_ = _1312___mcc_h3
			var _1313_cf3 D0 = _1312___mcc_h3
			_ = _1313_cf3
			if (_this).F6() {
				var _1314_v63 bool
				_ = _1314_v63
				_1314_v63 = true
				var _1315_v64 _dafny.Map
				_ = _1315_v64
				_1315_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _this.F15)
				var _1316_v65 _dafny.Map
				_ = _1316_v65
				_1316_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v63, (_this).F5())
				_1314_v63 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (func() _dafny.Int {
					if (_1315_v64).Contains((func() _dafny.CodePoint {
						if (_1316_v65).Contains(_1314_v63) {
							return (_1316_v65).Get(_1314_v63).(_dafny.CodePoint)
						}
						return (_this).F5()
					})()) {
						return (_1315_v64).Get((func() _dafny.CodePoint {
							if (_1316_v65).Contains(_1314_v63) {
								return (_1316_v65).Get(_1314_v63).(_dafny.CodePoint)
							}
							return (_this).F5()
						})()).(_dafny.Int)
					}
					return _this.F15
				})())).Equals(func() _dafny.Map {
					var _coll46 = _dafny.NewMapBuilder()
					_ = _coll46
					for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(938), _dafny.IntOfInt64(755))); ; {
						_compr_46, _ok47 := _iter47()
						if !_ok47 {
							break
						}
						var _1317_v66 _dafny.Int
						_1317_v66 = interface{}(_compr_46).(_dafny.Int)
						if ((_dafny.IntOfInt64(938)).Cmp(_1317_v66) <= 0) && ((_1317_v66).Cmp(_dafny.IntOfInt64(755)) < 0) {
							_coll46.Add((_1317_v66).Times(_dafny.IntOfInt64(198)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(998), func() _dafny.Map {
								var _coll47 = _dafny.NewMapBuilder()
								_ = _coll47
								for _iter48 := _dafny.Iterate((_dafny.SetOf((_this).F5(), (_this).F5(), (_this).F5())).Elements()); ; {
									_compr_47, _ok48 := _iter48()
									if !_ok48 {
										break
									}
									var _1318_v67 _dafny.CodePoint
									_1318_v67 = interface{}(_compr_47).(_dafny.CodePoint)
									if (_dafny.SetOf((_this).F5(), (_this).F5(), (_this).F5())).Contains(_1318_v67) {
										_coll47.Add(_1318_v67, false)
									}
								}
								return _coll47.ToMap()
							}())).Cardinality())
						}
					}
					return _coll46.ToMap()
				}())
				_1314_v63 = _1314_v63
				var _1319_v68 _dafny.Array
				_ = _1319_v68
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_38
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) _dafny.Sequence = func(_1320_i10 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("tvpi")
					}
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw208 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw208).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw208).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1319_v68 = _nw208
				var _1321_v69 _dafny.Map
				_ = _1321_v69
				_1321_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v63, _dafny.UnicodeSeqOfUtf8Bytes("yfuflx"))
				var _1322_v70 _dafny.Sequence
				_ = _1322_v70
				_1322_v70 = _dafny.UnicodeSeqOfUtf8Bytes("bc")
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1319_v68), 0))
				_ = _index233
				(_1319_v68).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1321_v69).Contains((_this).F6()) {
						return (_1321_v69).Get((_this).F6()).(_dafny.Sequence)
					}
					return _1322_v70
				})(), _1322_v70), (_index233).Int())
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1319_v68), 0))
				_ = _index234
				(_1319_v68).ArraySet1(_1322_v70, (_index234).Int())
				var _1323_v71 _dafny.Array
				_ = _1323_v71
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
				_ = _nw209
				_1323_v71 = _nw209
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_1323_v71), 0))
				_ = _index235
				(_1323_v71).ArraySet1(!(_1314_v63), (_index235).Int())
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(452), _dafny.ArrayLen((_1323_v71), 0))
				_ = _index236
				(_1323_v71).ArraySet1(_1314_v63, (_index236).Int())
				var _1324_v72 _dafny.Map
				_ = _1324_v72
				_1324_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(871), (_this).F6())
				var _1325_v73 _dafny.Map
				_ = _1325_v73
				_1325_v73 = _1324_v72
				var _1326_v74 _dafny.Map
				_ = _1326_v74
				_1326_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_1325_v73), _this.F15)
				(globalState).F1 = (func() _dafny.Int {
					if (_1326_v74).Contains(Companion_Default___.Fm50(Companion_Default___.Fm0(_1314_v63, globalState), (_this).F6(), func() _dafny.Map {
						var _coll48 = _dafny.NewMapBuilder()
						_ = _coll48
						for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(477), _dafny.IntOfInt64(-150))); ; {
							_compr_48, _ok49 := _iter49()
							if !_ok49 {
								break
							}
							var _1327_v75 _dafny.Int
							_1327_v75 = interface{}(_compr_48).(_dafny.Int)
							if ((_dafny.IntOfInt64(477)).Cmp(_1327_v75) <= 0) && ((_1327_v75).Cmp(_dafny.IntOfInt64(-150)) < 0) {
								_coll48.Add((_1327_v75).Times(_dafny.IntOfUint32(((_1319_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1319_v68), 0))).Int()).(_dafny.Sequence)).Cardinality())), _this.F15)
							}
						}
						return _coll48.ToMap()
					}(), globalState)) {
						return (_1326_v74).Get(Companion_Default___.Fm50(Companion_Default___.Fm0(_1314_v63, globalState), (_this).F6(), func() _dafny.Map {
							var _coll49 = _dafny.NewMapBuilder()
							_ = _coll49
							for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(477), _dafny.IntOfInt64(-150))); ; {
								_compr_49, _ok50 := _iter50()
								if !_ok50 {
									break
								}
								var _1328_v75 _dafny.Int
								_1328_v75 = interface{}(_compr_49).(_dafny.Int)
								if ((_dafny.IntOfInt64(477)).Cmp(_1328_v75) <= 0) && ((_1328_v75).Cmp(_dafny.IntOfInt64(-150)) < 0) {
									_coll49.Add((_1328_v75).Times(_dafny.IntOfUint32(((_1319_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(88), _dafny.ArrayLen((_1319_v68), 0))).Int()).(_dafny.Sequence)).Cardinality())), _this.F15)
								}
							}
							return _coll49.ToMap()
						}(), globalState)).(_dafny.Int)
					}
					return _this.F15
				})()
			} else {
				var _1329_v76 bool
				_ = _1329_v76
				_1329_v76 = false
				var _1330_v77 _dafny.Map
				_ = _1330_v77
				_1330_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1329_v76)
				_1329_v76 = (((_1330_v77).Cardinality()).Minus(_this.F15)).Cmp(_this.F15) == 0
				(globalState).F1 = (Companion_Default___.Fm3((_this).F6(), true, globalState)).Plus(Companion_Default___.SafeModuloInt(_this.F15, _this.F15))
				var _1331_v78 _dafny.Sequence
				_ = _1331_v78
				_1331_v78 = _dafny.UnicodeSeqOfUtf8Bytes("paf")
				var _1332_v79 _dafny.Map
				_ = _1332_v79
				_1332_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F15)
				var _1333_v80 _dafny.Map
				_ = _1333_v80
				_1333_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1332_v79).Cardinality(), _this.F15)
				var _1334_v81 _dafny.Sequence
				_ = _1334_v81
				_1334_v81 = _dafny.SeqOf(_this.F15, _this.F15, (_1333_v80).Cardinality())
				var _1335_v82 _dafny.Array
				_ = _1335_v82
				var _nwElement0_41 _dafny.Int = (_dafny.Zero).Minus(_this.F15)
				_ = _nwElement0_41
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(25))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_41, 0)
				(_nw210).ArraySet1(_this.F15, 1)
				(_nw210).ArraySet1((_this.F15).Plus(_dafny.IntOfInt64(462)), 2)
				(_nw210).ArraySet1(_dafny.IntOfUint32((_1331_v78).Cardinality()), 3)
				(_nw210).ArraySet1(_this.F15, 4)
				(_nw210).ArraySet1((_this.F15).Plus(_dafny.IntOfUint32((_1331_v78).Cardinality())), 5)
				(_nw210).ArraySet1(_this.F15, 6)
				(_nw210).ArraySet1((func() _dafny.Int {
					if true {
						return _dafny.IntOfInt64(159)
					}
					return _dafny.IntOfUint32((_1331_v78).Cardinality())
				})(), 7)
				(_nw210).ArraySet1(_dafny.IntOfInt64(72), 8)
				(_nw210).ArraySet1((_1334_v81).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1334_v81).Cardinality()))).Uint32()).(_dafny.Int), 9)
				(_nw210).ArraySet1(_this.F15, 10)
				(_nw210).ArraySet1(_this.F15, 11)
				(_nw210).ArraySet1(_this.F15, 12)
				(_nw210).ArraySet1(_this.F15, 13)
				(_nw210).ArraySet1(_this.F15, 14)
				(_nw210).ArraySet1(_this.F15, 15)
				(_nw210).ArraySet1(_this.F15, 16)
				(_nw210).ArraySet1((_this.F15).Times(_this.F15), 17)
				(_nw210).ArraySet1((_dafny.Zero).Minus(_this.F15), 18)
				(_nw210).ArraySet1(Companion_Default___.SafeModuloInt(_this.F15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _this.F15)).Cardinality()), 19)
				(_nw210).ArraySet1(_this.F15, 20)
				(_nw210).ArraySet1((_1334_v81).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1334_v81).Cardinality()))).Uint32()).(_dafny.Int), 21)
				(_nw210).ArraySet1(_this.F15, 22)
				(_nw210).ArraySet1(_this.F15, 23)
				(_nw210).ArraySet1(_this.F15, 24)
				_1335_v82 = _nw210
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1335_v82), 0))
				_ = _index237
				(_1335_v82).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F15, (_this).Fm8(true, (_this).F6(), (_this).F6(), globalState)), (_index237).Int())
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1335_v82), 0))
				_ = _index238
				(_1335_v82).ArraySet1(_dafny.IntOfInt64(149), (_index238).Int())
				var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1335_v82), 0))
				_ = _index239
				(_1335_v82).ArraySet1((_1335_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1335_v82), 0))).Int()).(_dafny.Int), (_index239).Int())
				_1331_v78 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer84 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg84 _dafny.Int) interface{} {
						return coer84(arg84)
					}
				}(func(_1336_i11 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				})), _dafny.Companion_Sequence_.Update(_1331_v78, (Companion_Default___.SafeIndex((_1335_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1335_v82), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1331_v78).Cardinality()))).Uint32(), (_this).F5()))
			}
			if (_this).F6() {
				(globalState).F1 = (_dafny.IntOfInt64(355)).Times(_this.F15)
				var _1337_v83 bool
				_ = _1337_v83
				_1337_v83 = false
				_1337_v83 = _1337_v83
				var _1338_v84 _dafny.Sequence
				_ = _1338_v84
				_1338_v84 = _dafny.SeqOf(_this.F15, _this.F15, _dafny.IntOfInt64(904), _this.F15)
				var _1339_v85 _dafny.Sequence
				_ = _1339_v85
				_1339_v85 = _dafny.SeqOf(_1338_v84)
				var _1340_v86 *C0
				_ = _1340_v86
				var _nw211 *C0 = New_C0_()
				_ = _nw211
				_nw211.Ctor__(_1337_v83, _1339_v85)
				_1340_v86 = _nw211
				var _1341_v87 _dafny.Sequence
				_ = _1341_v87
				_1341_v87 = _dafny.SeqOf((Companion_D14_.Create_DC26_(_1337_v83, _this.F15, (_this).F5())).Dtor_cf32(), (_this).F6(), false)
				var _1342_v88 _dafny.Map
				_ = _1342_v88
				_1342_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_1340_v86).F11(), _1337_v83), (_this).F6())
				var _1343_v89 _dafny.Set
				_ = _1343_v89
				_1343_v89 = _dafny.SetOf((_this).Fm6(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(672))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg85 _dafny.Int) interface{} {
						return coer85(arg85)
					}
				}(func(_1344_i12 _dafny.Int) _dafny.Int {
					return _this.F15
				})), globalState))
				var _1345_v90 _dafny.Sequence
				_ = _1345_v90
				_1345_v90 = _dafny.UnicodeSeqOfUtf8Bytes("jdvhobhq")
				var _1346_v91 _dafny.Map
				_ = _1346_v91
				_1346_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8(_1337_v83, (_1340_v86).F11(), (_1340_v86).F11(), globalState), (_1340_v86).F11())
				var _1347_v92 _dafny.Map
				_ = _1347_v92
				_1347_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1338_v84).Cardinality()), _1345_v90)
				var _1348_v93 _dafny.Array
				_ = _1348_v93
				var _nwElement0_42 bool = (_this).F6()
				_ = _nwElement0_42
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(26))
				_ = _nw212
				(_nw212).ArraySet1(_nwElement0_42, 0)
				(_nw212).ArraySet1((_1340_v86).F11(), 1)
				(_nw212).ArraySet1(_1337_v83, 2)
				(_nw212).ArraySet1((_this.F15).Cmp(_this.F15) <= 0, 3)
				(_nw212).ArraySet1((_1340_v86).F11(), 4)
				(_nw212).ArraySet1((Companion_Default___.Fm0(_1337_v83, globalState)) || ((_1340_v86).F11()), 5)
				(_nw212).ArraySet1((_1340_v86).F11(), 6)
				(_nw212).ArraySet1((_1341_v87).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1341_v87).Cardinality()))).Uint32()).(bool), 7)
				(_nw212).ArraySet1(!((_this).F6()) || ((_1340_v86).F11()), 8)
				(_nw212).ArraySet1(false, 9)
				(_nw212).ArraySet1(!((_1340_v86).F11()), 10)
				(_nw212).ArraySet1((_this.F15).Cmp(_this.F15) < 0, 11)
				(_nw212).ArraySet1((_1340_v86).F11(), 12)
				(_nw212).ArraySet1((_1340_v86).F11(), 13)
				(_nw212).ArraySet1((func() bool {
					if (_1342_v88).Contains(_1343_v89) {
						return (_1342_v88).Get(_1343_v89).(bool)
					}
					return _1337_v83
				})(), 14)
				(_nw212).ArraySet1(_dafny.Companion_Sequence_.Contains(_1345_v90, (_this).F5()), 15)
				(_nw212).ArraySet1((_1340_v86).F11(), 16)
				(_nw212).ArraySet1((_1340_v86).F11(), 17)
				(_nw212).ArraySet1((_1340_v86).F11(), 18)
				(_nw212).ArraySet1((_this).F6(), 19)
				(_nw212).ArraySet1((func() bool {
					if (_1346_v91).Contains((_1347_v92).Cardinality()) {
						return (_1346_v91).Get((_1347_v92).Cardinality()).(bool)
					}
					return _1337_v83
				})(), 20)
				(_nw212).ArraySet1((_this).F6(), 21)
				(_nw212).ArraySet1(((_1340_v86).F11()) && (_1337_v83), 22)
				(_nw212).ArraySet1((_this).F6(), 23)
				(_nw212).ArraySet1((_1340_v86).F11(), 24)
				(_nw212).ArraySet1(true, 25)
				_1348_v93 = _nw212
				var _1349_v94 D2
				_ = _1349_v94
				_1349_v94 = Companion_D2_.Create_DC8_(_1337_v83)
				var _rhs250 _dafny.Int = _this.F15
				_ = _rhs250
				var _rhs251 _dafny.MultiSet = (((_dafny.MultiSetOf(_1337_v83)).Update((_1340_v86).F11(), Companion_Default___.Abs(_this.F15))).Update((_1340_v86).F11(), Companion_Default___.Abs(_this.F15))).Update(_1337_v83, Companion_Default___.Abs(_this.F15))
				_ = _rhs251
				var _rhs252 _dafny.Array = (func() _dafny.Array {
					if !((_this).F6()) {
						return _1348_v93
					}
					return _1348_v93
				})()
				_ = _rhs252
				var _rhs253 D2 = _1349_v94
				_ = _rhs253
				var _rhs254 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(539), Companion_Default___.SafeModuloInt((Companion_Default___.Fm46((_1340_v86).F11(), globalState)).Cardinality(), _this.F15))
				_ = _rhs254
				var _lhs175 *C5 = _this
				_ = _lhs175
				var _lhs176 *GlobalState = globalState
				_ = _lhs176
				_lhs175.F15 = _rhs250
				_lhs176.F4 = _rhs251
				_1348_v93 = _rhs252
				_1349_v94 = _rhs253
				r0 = _rhs254
				var _1350_v95 *C3
				_ = _1350_v95
				var _nw213 *C3 = New_C3_()
				_ = _nw213
				_nw213.Ctor__((Companion_Default___.Fm51(_1337_v83, (_1346_v91).Cardinality(), (_this).F6(), globalState)).IsSubsetOf(_dafny.SetOf((_this).F5())), (_this).F6())
				_1350_v95 = _nw213
			} else {
				var _1351_v96 _dafny.MultiSet
				_ = _1351_v96
				_1351_v96 = _dafny.MultiSetOf((_this).F6())
				var _1352_v97 *C3
				_ = _1352_v97
				var _nw214 *C3 = New_C3_()
				_ = _nw214
				_nw214.Ctor__((_1351_v96).IsSubsetOf(_1351_v96), !((_this).F6()))
				_1352_v97 = _nw214
				var _1353_v98 _dafny.Sequence
				_ = _1353_v98
				_1353_v98 = _dafny.SeqOf((_this).F5())
				var _1354_v99 _dafny.Map
				_ = _1354_v99
				_1354_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F5(), (_this).F5(), (_this).F5(), (_this).F5()), _1353_v98), _this.F15)
				_1354_v99 = (_1354_v99).Update(_1353_v98, _this.F15)
				(_this).F15 = (_this.F15).Plus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1353_v98).Cardinality()), (_1352_v97).F14())).Update(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aknri")).Cardinality()), (_1352_v97).F13())).Cardinality())
				var _1355_v100 _dafny.Array
				_ = _1355_v100
				var _len0_39 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_39
				var _nw215 _dafny.Array
				_ = _nw215
				if _len0_39.Cmp(_dafny.Zero) == 0 {
					_nw215 = _dafny.NewArray(_len0_39)
				} else {
					var _init39 func(_dafny.Int) _dafny.Int = func(_1356_i13 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1356_i13, _this.F15)
					}
					_ = _init39
					var _element0_39 = _init39(_dafny.Zero)
					_ = _element0_39
					_nw215 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
					(_nw215).ArraySet1(_element0_39, 0)
					var _nativeLen0_39 = (_len0_39).Int()
					_ = _nativeLen0_39
					for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
						(_nw215).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
					}
				}
				_1355_v100 = _nw215
				var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1355_v100), 0))
				_ = _index240
				(_1355_v100).ArraySet1(_this.F15, (_index240).Int())
				var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1355_v100), 0))
				_ = _index241
				(_1355_v100).ArraySet1(_this.F15, (_index241).Int())
				var _1357_v101 D15
				_ = _1357_v101
				_1357_v101 = Companion_D15_.Create_DC27_(_1351_v96)
				_1357_v101 = _1357_v101
			}
			var _1358_v102 _dafny.Map
			_ = _1358_v102
			var _1359_v103 _dafny.Sequence
			_ = _1359_v103
			var _1360_v104 _dafny.Set
			_ = _1360_v104
			var _1361_v105 bool
			_ = _1361_v105
			var _out36 _dafny.Map
			_ = _out36
			var _out37 _dafny.Sequence
			_ = _out37
			var _out38 _dafny.Set
			_ = _out38
			var _out39 bool
			_ = _out39
			_out36, _out37, _out38, _out39 = (_this).M12(_this.F15, (_this).F6(), _this.F15, globalState)
			_1358_v102 = _out36
			_1359_v103 = _out37
			_1360_v104 = _out38
			_1361_v105 = _out39
			var _1362_v106 _dafny.Array
			_ = _1362_v106
			var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw216
			_1362_v106 = _nw216
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1362_v106), 0))
			_ = _index242
			(_1362_v106).ArraySet1(_this.F15, (_index242).Int())
			var _1363_v107 _dafny.Sequence
			_ = _1363_v107
			_1363_v107 = _dafny.SeqOf((_this).F6(), (_this).F6(), _1361_v105, _1361_v105, (_this).F6())
			var _1364_v108 _dafny.MultiSet
			_ = _1364_v108
			_1364_v108 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1359_v103).Cardinality()), _this.F15)
			var _1365_v109 _dafny.MultiSet
			_ = _1365_v109
			_1365_v109 = _dafny.MultiSetOf((_dafny.Zero).Minus(_this.F15), (_this).Fm8(false, (_this).F6(), _1361_v105, globalState), (_this).Fm4(_1363_v107, _this.F15, _this.F15, _1364_v108, globalState), _this.F15, _this.F15)
			var _1366_v110 _dafny.Map
			_ = _1366_v110
			_1366_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F15)
			var _1367_v111 _dafny.Sequence
			_ = _1367_v111
			_1367_v111 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_this.F15, (func() _dafny.Int {
				if (_1364_v108).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15)).Cardinality()) {
					return (_1364_v108).Multiplicity((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15)).Cardinality())
				}
				return _this.F15
			})()), _1366_v110))
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(477), _dafny.ArrayLen((_1362_v106), 0))
			_ = _index243
			(_1362_v106).ArraySet1((_this.F15).Plus((((_1367_v111).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1367_v111).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1359_v103, _1366_v110))).Cardinality()), (_index243).Int())
		}
		(_this).F15 = ((_this).Fm7((_this).F6(), globalState)).Minus(_this.F15)
		var _1368_v112 _dafny.MultiSet
		_ = _1368_v112
		_1368_v112 = _dafny.MultiSetOf(_this.F15, _dafny.IntOfInt64(521))
		var _hi7 _dafny.Int = Companion_Default___.SafeModuloInt(_this.F15, _this.F15)
		_ = _hi7
		for _1369_i14 := (func() _dafny.Int {
			if (_1368_v112).Contains(_this.F15) {
				return (_1368_v112).Multiplicity(_this.F15)
			}
			return _this.F15
		})(); _1369_i14.Cmp(_hi7) < 0; _1369_i14 = _1369_i14.Plus(_dafny.One) {
			var _1370_v113 _dafny.Array
			_ = _1370_v113
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_40
			var _nw217 _dafny.Array
			_ = _nw217
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw217 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) _dafny.Int = func(_1371_i15 _dafny.Int) _dafny.Int {
					return (_1371_i15).Times(_this.F15)
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw217 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw217).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw217).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1370_v113 = _nw217
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1370_v113), 0))
			_ = _index244
			(_1370_v113).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jwrbku")).Cardinality()), (_index244).Int())
			var _1372_v114 _dafny.Set
			_ = _1372_v114
			_1372_v114 = _dafny.SetOf(_1369_i14)
			var _1373_v115 _dafny.Sequence
			_ = _1373_v115
			_1373_v115 = _dafny.SeqOf((_this).F6(), (_this).F6(), !((_this).F6()) || ((_this).F6()), ((_1372_v114).Cardinality()).Cmp(_this.F15) <= 0)
			var _1374_v116 _dafny.MultiSet
			_ = _1374_v116
			_1374_v116 = _dafny.MultiSetOf((_this).F6())
			var _1375_v117 D15
			_ = _1375_v117
			_1375_v117 = Companion_D15_.Create_DC27_(_1374_v116)
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1370_v113), 0))
			_ = _index245
			var _rhs255 _dafny.Int = (_dafny.Zero).Minus(_this.F15)
			_ = _rhs255
			var _rhs256 _dafny.Sequence = _1373_v115
			_ = _rhs256
			var _rhs257 _dafny.Int = _dafny.IntOfInt64(-813)
			_ = _rhs257
			var _rhs258 D15 = Companion_D15_.Create_DC27_(_1374_v116)
			_ = _rhs258
			var _lhs177 _dafny.Array = _1370_v113
			_ = _lhs177
			var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1370_v113), 0))
			_ = _lhs178
			var _lhs179 *C5 = _this
			_ = _lhs179
			(_lhs177).ArraySet1(_rhs255, (_lhs178).Int())
			_1373_v115 = _rhs256
			_lhs179.F15 = _rhs257
			_1375_v117 = _rhs258
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_1370_v113), 0))
			_ = _index246
			(_1370_v113).ArraySet1(_dafny.IntOfInt64(327), (_index246).Int())
			var _1376_v118 _dafny.Array
			_ = _1376_v118
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.One)
			_ = _nw218
			_1376_v118 = _nw218
			_1376_v118 = _1376_v118
			var _1377_v119 bool
			_ = _1377_v119
			_1377_v119 = false
			var _1378_v120 _dafny.MultiSet
			_ = _1378_v120
			_1378_v120 = _dafny.MultiSetOf(_dafny.CodePoint('o'), (_this).F5(), (_this).F5(), (_this).F5(), _dafny.CodePoint('c'))
			_1377_v119 = ((_dafny.Zero).Minus(_1369_i14)).Cmp((_1378_v120).Cardinality()) < 0
		}
		var _1379_v121 bool
		_ = _1379_v121
		_1379_v121 = false
		_1379_v121 = (_1379_v121) && (_1379_v121)
		var _1380_v122 _dafny.Sequence
		_ = _1380_v122
		_1380_v122 = _dafny.UnicodeSeqOfUtf8Bytes("wqjcniels")
		_1380_v122 = _1380_v122
		var _1381_v123 _dafny.Array
		_ = _1381_v123
		var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw219
		_1381_v123 = _nw219
		var _1382_v124 *C1
		_ = _1382_v124
		var _nw220 *C1 = New_C1_()
		_ = _nw220
		_nw220.Ctor__()
		_1382_v124 = _nw220
		var _1383_v125 _dafny.Array
		_ = _1383_v125
		var _nw221 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw221
		_1383_v125 = _nw221
		var _1384_v126 _dafny.Map
		_ = _1384_v126
		_1384_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1382_v124, _1383_v125)
		var _1385_v127 _dafny.Map
		_ = _1385_v127
		_1385_v127 = _1384_v126
		var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1381_v123), 0))
		_ = _index247
		(_1381_v123).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_1385_v127)).Cardinality()), (_index247).Int())
		var _1386_v128 _dafny.Sequence
		_ = _1386_v128
		_1386_v128 = _dafny.SeqOf(_this.F15, _this.F15)
		var _1387_v129 _dafny.Sequence
		_ = _1387_v129
		_1387_v129 = _dafny.SeqOf(_1386_v128)
		var _1388_v130 _dafny.Set
		_ = _1388_v130
		_1388_v130 = _dafny.SetOf(_1379_v121)
		var _1389_v131 _dafny.Map
		_ = _1389_v131
		_1389_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1387_v129, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1388_v130).Cardinality()), _dafny.IntOfUint32((_1387_v129).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(764))).Uint32(), func(coer86 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg86 _dafny.Int) interface{} {
				return coer86(arg86)
			}
		}(func(_1390_i16 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("utb")).Cardinality())
		}))), _1380_v122)
		var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1381_v123), 0))
		_ = _index248
		var _rhs259 _dafny.Int = _this.F15
		_ = _rhs259
		var _rhs260 _dafny.Sequence = (func() _dafny.Sequence {
			if (_1389_v131).Contains(_dafny.Companion_Sequence_.Concatenate(_1387_v129, Companion_Default___.Fm52((_this).F6(), false, globalState))) {
				return (_1389_v131).Get(_dafny.Companion_Sequence_.Concatenate(_1387_v129, Companion_Default___.Fm52((_this).F6(), false, globalState))).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("amowydn")
		})()
		_ = _rhs260
		var _lhs180 _dafny.Array = _1381_v123
		_ = _lhs180
		var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1381_v123), 0))
		_ = _lhs181
		(_lhs180).ArraySet1(_rhs259, (_lhs181).Int())
		_1380_v122 = _rhs260
		r0 = (_1386_v128).Select((Companion_Default___.SafeIndex((_1381_v123).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(712), _dafny.ArrayLen((_1381_v123), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1386_v128).Cardinality()))).Uint32()).(_dafny.Int)
		return r0
	}
}
func (_this *C5) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _1391_v0 _dafny.Map
		_ = _1391_v0
		_1391_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F15)
		var _hi8 _dafny.Int = (_this.F15).Times((_dafny.Zero).Minus(_this.F15))
		_ = _hi8
		for _1392_i0 := (func() _dafny.Int {
			if (_1391_v0).Contains(Companion_Default___.Fm0(p2, globalState)) {
				return (_1391_v0).Get(Companion_Default___.Fm0(p2, globalState)).(_dafny.Int)
			}
			return _this.F15
		})(); _1392_i0.Cmp(_hi8) < 0; _1392_i0 = _1392_i0.Plus(_dafny.One) {
			var _1393_v1 _dafny.Sequence
			_ = _1393_v1
			_1393_v1 = _dafny.UnicodeSeqOfUtf8Bytes("petsff")
			var _1394_v2 _dafny.Map
			_ = _1394_v2
			_1394_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1393_v1)
			var _1395_v3 _dafny.Array
			_ = _1395_v3
			var _nw222 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw222
			_1395_v3 = _nw222
			var _1396_v4 _dafny.Map
			_ = _1396_v4
			_1396_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_v3, _1392_i0)
			var _1397_v5 _dafny.Map
			_ = _1397_v5
			_1397_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !((_this).F6()))
			var _1398_v6 _dafny.Array
			_ = _1398_v6
			var _nwElement0_43 _dafny.Int = _1392_i0
			_ = _nwElement0_43
			var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(25))
			_ = _nw223
			(_nw223).ArraySet1(_nwElement0_43, 0)
			(_nw223).ArraySet1(_this.F15, 1)
			(_nw223).ArraySet1(((_this).Fm7(p2, globalState)).Plus((_dafny.Zero).Minus(_1392_i0)), 2)
			(_nw223).ArraySet1(_this.F15, 3)
			(_nw223).ArraySet1(_this.F15, 4)
			(_nw223).ArraySet1((_1392_i0).Times(_this.F15), 5)
			(_nw223).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_this).F6() {
					return _1393_v1
				}
				return _1393_v1
			})()).Cardinality())), 6)
			(_nw223).ArraySet1(_1392_i0, 7)
			(_nw223).ArraySet1((_dafny.Zero).Minus((_1394_v2).Cardinality()), 8)
			(_nw223).ArraySet1((_dafny.Zero).Minus(_1392_i0), 9)
			(_nw223).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if (_1396_v4).Contains(_1395_v3) {
					return (_1396_v4).Get(_1395_v3).(_dafny.Int)
				}
				return _1392_i0
			})())), 10)
			(_nw223).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_1392_i0)), 11)
			(_nw223).ArraySet1(_1392_i0, 12)
			(_nw223).ArraySet1(_this.F15, 13)
			(_nw223).ArraySet1((_dafny.IntOfInt64(317)).Times(_dafny.IntOfUint32((_1393_v1).Cardinality())), 14)
			(_nw223).ArraySet1(_dafny.IntOfUint32((_1393_v1).Cardinality()), 15)
			(_nw223).ArraySet1(_dafny.IntOfInt64(362), 16)
			(_nw223).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F15, _1392_i0)), 17)
			(_nw223).ArraySet1(_1392_i0, 18)
			(_nw223).ArraySet1(_this.F15, 19)
			(_nw223).ArraySet1(_this.F15, 20)
			(_nw223).ArraySet1(_this.F15, 21)
			(_nw223).ArraySet1((_dafny.Zero).Minus(_1392_i0), 22)
			(_nw223).ArraySet1(_this.F15, 23)
			(_nw223).ArraySet1((_this.F15).Times((_1397_v5).Cardinality()), 24)
			_1398_v6 = _nw223
			_1395_v3 = _1398_v6
			if Companion_Default___.Fm0(!((_this).F6()) || (p2), globalState) {
				(globalState).F1 = ((_dafny.Zero).Minus(_1392_i0)).Plus(_this.F15)
				var _1399_v7 bool
				_ = _1399_v7
				_1399_v7 = true
				_1399_v7 = p2
				var _1400_v8 _dafny.Array
				_ = _1400_v8
				var _nw224 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
				_ = _nw224
				_1400_v8 = _nw224
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1400_v8), 0))
				_ = _index249
				(_1400_v8).ArraySet1(_1399_v7, (_index249).Int())
				var _1401_v9 D11
				_ = _1401_v9
				_1401_v9 = Companion_D11_.Create_DC20_(_1399_v7, _dafny.IntOfInt64(77))
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1400_v8), 0))
				_ = _index250
				(_1400_v8).ArraySet1((_1401_v9).Dtor_cf21(), (_index250).Int())
				var _1402_v10 _dafny.Sequence
				_ = _1402_v10
				_1402_v10 = _dafny.SeqOf(_1400_v8, _1400_v8, _1400_v8)
				var _1403_v11 _dafny.Array
				_ = _1403_v11
				var _nwElement0_44 _dafny.Array = _1400_v8
				_ = _nwElement0_44
				var _nw225 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(16))
				_ = _nw225
				(_nw225).ArraySet1(_nwElement0_44, 0)
				(_nw225).ArraySet1(_1400_v8, 1)
				(_nw225).ArraySet1(_1400_v8, 2)
				(_nw225).ArraySet1((func() _dafny.Array {
					if (_this).F6() {
						return _1400_v8
					}
					return _1400_v8
				})(), 3)
				(_nw225).ArraySet1(_1400_v8, 4)
				(_nw225).ArraySet1(_1400_v8, 5)
				(_nw225).ArraySet1(_1400_v8, 6)
				(_nw225).ArraySet1(_1400_v8, 7)
				(_nw225).ArraySet1(_1400_v8, 8)
				(_nw225).ArraySet1(_1400_v8, 9)
				(_nw225).ArraySet1(_1400_v8, 10)
				(_nw225).ArraySet1(_1400_v8, 11)
				(_nw225).ArraySet1(_1400_v8, 12)
				(_nw225).ArraySet1(_1400_v8, 13)
				(_nw225).ArraySet1((_1402_v10).Select((Companion_Default___.SafeIndex(_1392_i0, _dafny.IntOfUint32((_1402_v10).Cardinality()))).Uint32()).(_dafny.Array), 14)
				(_nw225).ArraySet1(_1400_v8, 15)
				_1403_v11 = _nw225
				_1403_v11 = _1403_v11
				var _1404_v12 _dafny.Map
				_ = _1404_v12
				_1404_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1392_i0, _dafny.IntOfUint32((_dafny.SeqOf((_1400_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_1400_v8), 0))).Int()).(bool))).Cardinality()))
				var _1405_v13 _dafny.Map
				_ = _1405_v13
				_1405_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (_1404_v12).Cardinality())
				var _1406_v14 _dafny.MultiSet
				_ = _1406_v14
				_1406_v14 = _dafny.MultiSetOf((_this).F5())
				_1399_v7 = (_dafny.MultiSetOf((_this).F5(), Companion_Default___.Fm53(_1393_v1, (_dafny.Zero).Minus((_1405_v13).Cardinality()), globalState))).IsDisjointFrom((_1406_v14).Union(_dafny.MultiSetOf((_this).F5(), (_this).F5())))
			} else {
				var _1407_v15 bool
				_ = _1407_v15
				_1407_v15 = true
				_1407_v15 = (_1392_i0).Cmp(_1392_i0) < 0
				var _1408_v16 _dafny.Array
				_ = _1408_v16
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_41
				var _nw226 _dafny.Array
				_ = _nw226
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw226 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1409_p0 bool) func(_dafny.Int) bool {
						return func(_1410_i1 _dafny.Int) bool {
							return _1409_p0
						}
					})(p0)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw226 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw226).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw226).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1408_v16 = _nw226
				var _1411_v17 _dafny.CodePoint
				_ = _1411_v17
				_1411_v17 = _dafny.CodePoint('b')
				var _1412_v18 _dafny.Sequence
				_ = _1412_v18
				_1412_v18 = _dafny.SeqOf(_dafny.IntOfUint32((_1393_v1).Cardinality()), _1392_i0)
				var _rhs261 _dafny.Int = _this.F15
				_ = _rhs261
				var _rhs262 _dafny.Array = _1408_v16
				_ = _rhs262
				var _rhs263 _dafny.CodePoint = _1411_v17
				_ = _rhs263
				var _rhs264 bool = (_this).Fm6(_1412_v18, globalState)
				_ = _rhs264
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				_lhs182.F1 = _rhs261
				_1408_v16 = _rhs262
				_1411_v17 = _rhs263
				_1407_v15 = _rhs264
				(globalState).F1 = _this.F15
				(globalState).F4 = (p1).Intersection(_dafny.MultiSetOf(p0))
				_1411_v17 = _1411_v17
			}
			var _1413_v19 _dafny.Map
			_ = _1413_v19
			_1413_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, p2)
			_1413_v19 = (_1413_v19).Update(Companion_Default___.SafeDivisionInt(_this.F15, _this.F15), (_this).F6())
			var _1414_v20 *C3
			_ = _1414_v20
			var _nw227 *C3 = New_C3_()
			_ = _nw227
			_nw227.Ctor__((_this).F6(), p2)
			_1414_v20 = _nw227
		}
		var _1415_v21 _dafny.Sequence
		_ = _1415_v21
		_1415_v21 = _dafny.SeqOf((_this).F6())
		_1415_v21 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6()), _dafny.Companion_Sequence_.Update(_1415_v21, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1415_v21).Cardinality()))).Uint32(), (_this).F6()))
		var _1416_v22 _dafny.Map
		_ = _1416_v22
		_1416_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _this.F15)
		var _1417_v23 _dafny.Map
		_ = _1417_v23
		_1417_v23 = _1416_v22
		var _1418_v24 _dafny.Sequence
		_ = _1418_v24
		_1418_v24 = _dafny.SeqOf(_1415_v21)
		var _1419_v25 _dafny.Sequence
		_ = _1419_v25
		_1419_v25 = _dafny.UnicodeSeqOfUtf8Bytes("vinlx")
		var _1420_v26 D12
		_ = _1420_v26
		_1420_v26 = Companion_D12_.Create_DC22_(_1417_v23, _1418_v24, (_1419_v25).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32()).(_dafny.CodePoint), p0, _1419_v25)
		var _source20 D12 = _1420_v26
		_ = _source20
		if _source20.Is_DC22() {
			var _1421___mcc_h0 _dafny.Map = _source20.Get_().(D12_DC22).Cf24
			_ = _1421___mcc_h0
			var _1422___mcc_h1 _dafny.Sequence = _source20.Get_().(D12_DC22).Cf25
			_ = _1422___mcc_h1
			var _1423___mcc_h2 _dafny.CodePoint = _source20.Get_().(D12_DC22).Cf26
			_ = _1423___mcc_h2
			var _1424___mcc_h3 bool = _source20.Get_().(D12_DC22).Cf27
			_ = _1424___mcc_h3
			var _1425___mcc_h4 _dafny.Sequence = _source20.Get_().(D12_DC22).Cf28
			_ = _1425___mcc_h4
			var _1426_cf28 _dafny.Sequence = _1425___mcc_h4
			_ = _1426_cf28
			var _1427_cf27 bool = _1424___mcc_h3
			_ = _1427_cf27
			var _1428_cf26 _dafny.CodePoint = _1423___mcc_h2
			_ = _1428_cf26
			var _1429_cf25 _dafny.Sequence = _1422___mcc_h1
			_ = _1429_cf25
			var _1430_cf24 _dafny.Map = _1421___mcc_h0
			_ = _1430_cf24
			_1419_v25 = _dafny.UnicodeSeqOfUtf8Bytes("pnlpvun")
			if (_this).F6() {
				_1426_cf28 = _dafny.UnicodeSeqOfUtf8Bytes("stpkvw")
				var _1431_v27 _dafny.Array
				_ = _1431_v27
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
				_ = _nw228
				_1431_v27 = _nw228
				var _1432_v28 _dafny.Array
				_ = _1432_v28
				var _nw229 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw229
				_1432_v28 = _nw229
				var _1433_v29 _dafny.Map
				_ = _1433_v29
				_1433_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1432_v28, _this.F15)
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1431_v27), 0))
				_ = _index251
				(_1431_v27).ArraySet1(_1433_v29, (_index251).Int())
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(497), _dafny.ArrayLen((_1431_v27), 0))
				_ = _index252
				(_1431_v27).ArraySet1(_1433_v29, (_index252).Int())
				var _1434_v30 _dafny.Map
				_ = _1434_v30
				_1434_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1428_cf26, _this.F15)
				_1434_v30 = (_1434_v30).Update(Companion_Default___.Fm53(_1419_v25, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, !(_1427_cf27))).Cardinality(), globalState), _this.F15)
				var _1435_v31 _dafny.MultiSet
				_ = _1435_v31
				_1435_v31 = _dafny.MultiSetOf(_this.F15)
				_1427_cf27 = (_1435_v31).IsDisjointFrom(_1435_v31)
				(_this).F15 = _dafny.IntOfInt64(64)
			} else {
				var _1436_v32 _dafny.Sequence
				_ = _1436_v32
				_1436_v32 = _dafny.SeqOf(_this.F15)
				var _1437_v33 _dafny.Sequence
				_ = _1437_v33
				_1437_v33 = _dafny.SeqOf(_1436_v32)
				var _1438_v34 *C0
				_ = _1438_v34
				var _nw230 *C0 = New_C0_()
				_ = _nw230
				_nw230.Ctor__(!(!(_1427_cf27)), _dafny.Companion_Sequence_.Concatenate(_1437_v33, _1437_v33))
				_1438_v34 = _nw230
				var _1439_v35 _dafny.Array
				_ = _1439_v35
				var _nw231 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw231
				_1439_v35 = _nw231
				var _1440_v36 _dafny.Set
				_ = _1440_v36
				_1440_v36 = _dafny.SetOf(_this.F15)
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))
				_ = _index253
				(_1439_v35).ArraySet1(!(_1440_v36).Contains(_this.F15), (_index253).Int())
				var _1441_v37 _dafny.Map
				_ = _1441_v37
				_1441_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1439_v35, _dafny.SetOf(_dafny.IntOfInt64(563)))
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))
				_ = _index254
				var _rhs265 bool = !((p0) == (!_dafny.Companion_Sequence_.Equal(_1415_v21, _1415_v21)))
				_ = _rhs265
				var _rhs266 _dafny.MultiSet = ((p1).Intersection(_dafny.MultiSetOf(p2, p2, (_this).F6()))).Union(_dafny.MultiSetOf(p0))
				_ = _rhs266
				var _rhs267 _dafny.Map = (func() _dafny.Map {
					if (_this).Fm6(_1436_v32, globalState) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1439_v35, _1440_v36)
					}
					return _1441_v37
				})()
				_ = _rhs267
				var _lhs183 _dafny.Array = _1439_v35
				_ = _lhs183
				var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))
				_ = _lhs184
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				(_lhs183).ArraySet1(_rhs265, (_lhs184).Int())
				_lhs185.F4 = _rhs266
				_1441_v37 = _rhs267
				var _1442_v38 *C3
				_ = _1442_v38
				var _nw232 *C3 = New_C3_()
				_ = _nw232
				_nw232.Ctor__(!(p0), (_1439_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))).Int()).(bool))
				_1442_v38 = _nw232
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))
				_ = _index255
				var _rhs268 bool = (_1439_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))).Int()).(bool)
				_ = _rhs268
				var _rhs269 bool = (_1438_v34).F11()
				_ = _rhs269
				var _lhs186 _dafny.Array = _1439_v35
				_ = _lhs186
				var _lhs187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_1439_v35), 0))
				_ = _lhs187
				(_lhs186).ArraySet1(_rhs268, (_lhs187).Int())
				_1427_cf27 = _rhs269
				var _1443_v39 _dafny.Map
				_ = _1443_v39
				_1443_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1442_v38).F14(), p2)
				var _1444_v40 *C3
				_ = _1444_v40
				var _nw233 *C3 = New_C3_()
				_ = _nw233
				_nw233.Ctor__(!((_1443_v39).Equals((_1443_v39).Update((_this).F6(), p0))), p2)
				_1444_v40 = _nw233
			}
			var _1445_v41 _dafny.Array
			_ = _1445_v41
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(20))
			_ = _nw234
			_1445_v41 = _nw234
			var _1446_v42 _dafny.Sequence
			_ = _1446_v42
			_1446_v42 = _dafny.SeqOf(_this.F15)
			var _1447_v43 _dafny.Map
			_ = _1447_v43
			_1447_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).Fm6(_1446_v42, globalState))
			var _1448_v44 _dafny.Map
			_ = _1448_v44
			_1448_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1447_v43)
			var _1449_v45 _dafny.Set
			_ = _1449_v45
			_1449_v45 = _dafny.SetOf((func() _dafny.Map {
				if (_1448_v44).Contains(_this.F15) {
					return (_1448_v44).Get(_this.F15).(_dafny.Map)
				}
				return _1447_v43
			})())
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1445_v41), 0))
			_ = _index256
			(_1445_v41).ArraySet1(_1449_v45, (_index256).Int())
			var _1450_v46 _dafny.Map
			_ = _1450_v46
			_1450_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1449_v45)
			var _1451_v47 _dafny.Sequence
			_ = _1451_v47
			_1451_v47 = _dafny.SeqOf(_1447_v43)
			var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_1445_v41), 0))
			_ = _index257
			(_1445_v41).ArraySet1((func() _dafny.Set {
				if (_1450_v46).Contains((p0) && (p0)) {
					return (_1450_v46).Get((p0) && (p0)).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll50 = _dafny.NewBuilder()
					_ = _coll50
					for _iter51 := _dafny.Iterate((_1451_v47).Elements()); ; {
						_compr_50, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1452_v48 _dafny.Map
						_1452_v48 = interface{}(_compr_50).(_dafny.Map)
						if _dafny.Companion_Sequence_.Contains(_1451_v47, _1452_v48) {
							_coll50.Add(_1452_v48)
						}
					}
					return _coll50.ToSet()
				}()
			})(), (_index257).Int())
			var _1453_v49 _dafny.Array
			_ = _1453_v49
			var _nw235 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw235
			_1453_v49 = _nw235
			var _1454_v50 _dafny.Map
			_ = _1454_v50
			_1454_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1453_v49, Companion_Default___.SafeModuloInt(_this.F15, _this.F15))
			_1454_v50 = (_1454_v50).Update(_1453_v49, Companion_Default___.SafeDivisionInt(_this.F15, _this.F15))
		} else if _source20.Is_DC23() {
			var _1455___mcc_h5 _dafny.Array = _source20.Get_().(D12_DC23).Cf29
			_ = _1455___mcc_h5
			var _1456_cf29 _dafny.Array = _1455___mcc_h5
			_ = _1456_cf29
			var _1457_v51 D15
			_ = _1457_v51
			_1457_v51 = Companion_D15_.Create_DC28_((false) == (p0), ((Companion_D1_.Create_DC4_(_this.F15)).Dtor_cf4()).Cmp((_dafny.MultiSetOf(_dafny.CodePoint('n'), (_this).F5(), _dafny.CodePoint('g'), (_this).F5(), (_this).F5())).Cardinality()) <= 0)
			var _pat_let_tv18 = p0
			_ = _pat_let_tv18
			_1457_v51 = func(_pat_let16_0 D15) D15 {
				return func(_1458_dt__update__tmp_h0 D15) D15 {
					return func(_pat_let17_0 bool) D15 {
						return func(_1459_dt__update_hcf37_h0 bool) D15 {
							return Companion_D15_.Create_DC28_((_1458_dt__update__tmp_h0).Dtor_cf36(), _1459_dt__update_hcf37_h0)
						}(_pat_let17_0)
					}(_pat_let_tv18)
				}(_pat_let16_0)
			}(Companion_Default___.Fm54(globalState))
			(globalState).F1 = _this.F15
			(globalState).F1 = _this.F15
			var _1460_v52 _dafny.Map
			_ = _1460_v52
			_1460_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, p2)
			var _1461_v53 _dafny.Map
			_ = _1461_v53
			var _1462_v54 _dafny.Sequence
			_ = _1462_v54
			var _1463_v55 _dafny.Set
			_ = _1463_v55
			var _1464_v56 bool
			_ = _1464_v56
			var _out40 _dafny.Map
			_ = _out40
			var _out41 _dafny.Sequence
			_ = _out41
			var _out42 _dafny.Set
			_ = _out42
			var _out43 bool
			_ = _out43
			_out40, _out41, _out42, _out43 = (_this).M12((Companion_Default___.Fm55(_this.F15, _this.F15, globalState)).Cardinality(), ((_dafny.Zero).Minus((_1460_v52).Cardinality())).Cmp(_this.F15) >= 0, _this.F15, globalState)
			_1461_v53 = _out40
			_1462_v54 = _out41
			_1463_v55 = _out42
			_1464_v56 = _out43
		} else {
			var _1465___mcc_h6 _dafny.Map = _source20.Get_().(D12_DC21).Cf23
			_ = _1465___mcc_h6
			var _1466_cf23 _dafny.Map = _1465___mcc_h6
			_ = _1466_cf23
			var _1467_v57 _dafny.Array
			_ = _1467_v57
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_42
			var _nw236 _dafny.Array
			_ = _nw236
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw236 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.CodePoint = func(_1468_i2 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				}
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw236 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw236).ArraySet1CodePoint(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw236).ArraySet1CodePoint(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1467_v57 = _nw236
			_1467_v57 = _1467_v57
			(globalState).F1 = _dafny.IntOfInt64(229)
			var _1469_v58 _dafny.Array
			_ = _1469_v58
			var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(28))
			_ = _nw237
			_1469_v58 = _nw237
			var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1469_v58), 0))
			_ = _index258
			(_1469_v58).ArraySet1((_dafny.MultiSetOf((_this).F6())).Intersection(p1), (_index258).Int())
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1469_v58), 0))
			_ = _index259
			(_1469_v58).ArraySet1(p1, (_index259).Int())
			var _1470_v59 _dafny.Array
			_ = _1470_v59
			var _nw238 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw238
			_1470_v59 = _nw238
			var _1471_v60 _dafny.Map
			_ = _1471_v60
			_1471_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1470_v59, (func() bool {
				if (_this).F6() {
					return (_this).F6()
				}
				return p0
			})())
			if (func() bool {
				if (_1471_v60).Contains(_1470_v59) {
					return (_1471_v60).Get(_1470_v59).(bool)
				}
				return true
			})() {
				(_this).F15 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yl")).Cardinality())
				var _1472_v61 _dafny.Map
				_ = _1472_v61
				_1472_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, _1470_v59)
				_1472_v61 = (_1472_v61).Update(_this.F15, _1470_v59)
				var _1473_v62 bool
				_ = _1473_v62
				_1473_v62 = false
				_1473_v62 = p0
				(globalState).F1 = (_dafny.Zero).Minus((_this.F15).Plus(_this.F15))
				var _1474_v63 _dafny.Sequence
				_ = _1474_v63
				_1474_v63 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1419_v25).Cardinality()))))
				var _1475_v64 _dafny.Map
				_ = _1475_v64
				_1475_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1474_v63, _dafny.CodePoint('p'))
				_1475_v64 = (_1475_v64).Update(_dafny.Companion_Sequence_.Concatenate(_1474_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(4))).Uint32(), func(coer87 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}((func(_1476_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1477_i3 _dafny.Int) _dafny.Map {
						return _1476_v0
					}
				})(_1391_v0)))), (_this).F5())
			} else {
				var _1478_v65 bool
				_ = _1478_v65
				_1478_v65 = false
				var _1479_v66 _dafny.Array
				_ = _1479_v66
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_43
				var _nw239 _dafny.Array
				_ = _nw239
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw239 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) _dafny.Int = func(_1480_i4 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1480_i4, _this.F15)
					}
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw239 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw239).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw239).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1479_v66 = _nw239
				var _1481_v67 _dafny.Map
				_ = _1481_v67
				_1481_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1479_v66, p0)
				var _1482_v68 _dafny.Sequence
				_ = _1482_v68
				_1482_v68 = _dafny.SeqOf(_1479_v66, _1479_v66)
				var _1483_v69 D1
				_ = _1483_v69
				_1483_v69 = Companion_D1_.Create_DC5_(_1478_v65)
				_1478_v65 = !(!((func() bool {
					if (_1481_v67).Contains((_1482_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1415_v21).Cardinality()), _dafny.IntOfUint32((_1482_v68).Cardinality()))).Uint32()).(_dafny.Array)) {
						return (_1481_v67).Get((_1482_v68).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1415_v21).Cardinality()), _dafny.IntOfUint32((_1482_v68).Cardinality()))).Uint32()).(_dafny.Array)).(bool)
					}
					return (_1483_v69).Dtor_cf5()
				})()))
				(globalState).F1 = _this.F15
				var _1484_v70 _dafny.MultiSet
				_ = _1484_v70
				_1484_v70 = _dafny.MultiSetOf(_this.F15)
				var _1485_v71 _dafny.Sequence
				_ = _1485_v71
				_1485_v71 = _dafny.SeqOf(_dafny.IntOfUint32((_1419_v25).Cardinality()))
				var _1486_v72 _dafny.Sequence
				_ = _1486_v72
				_1486_v72 = _dafny.SeqOf((func() _dafny.Int {
					if (_1484_v70).Contains((func() _dafny.Int {
						if (p1).Contains(_1478_v65) {
							return (p1).Multiplicity(_1478_v65)
						}
						return _this.F15
					})()) {
						return (_1484_v70).Multiplicity((func() _dafny.Int {
							if (p1).Contains(_1478_v65) {
								return (p1).Multiplicity(_1478_v65)
							}
							return _this.F15
						})())
					}
					return _this.F15
				})(), _dafny.IntOfUint32((_1485_v71).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sf")).Cardinality()), _this.F15, _this.F15)
				var _1487_v73 _dafny.MultiSet
				_ = _1487_v73
				_1487_v73 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1485_v71).Cardinality()))
				var _1488_v74 _dafny.Map
				_ = _1488_v74
				_1488_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1484_v70, _this.F15)
				(_this).F15 = (func() _dafny.Int {
					if (func() bool {
						if p0 {
							return (_this).F6()
						}
						return p0
					})() {
						return (func() _dafny.Int {
							if (_1488_v74).Contains(_1487_v73) {
								return (_1488_v74).Get(_1487_v73).(_dafny.Int)
							}
							return _this.F15
						})()
					}
					return _this.F15
				})()
				var _1489_v75 _dafny.CodePoint
				_ = _1489_v75
				_1489_v75 = _dafny.CodePoint('k')
				var _1490_v76 D1
				_ = _1490_v76
				_1490_v76 = Companion_D1_.Create_DC4_(_this.F15)
				var _rhs270 bool = !(p0)
				_ = _rhs270
				var _rhs271 bool = (p2) && (!((p1).IsProperSubsetOf((_1469_v58).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(320), _dafny.ArrayLen((_1469_v58), 0))).Int()).(_dafny.MultiSet))))
				_ = _rhs271
				var _rhs272 bool = !((_this).F6()) || (!(Companion_Default___.Fm0((_this).F6(), globalState)))
				_ = _rhs272
				var _rhs273 _dafny.CodePoint = _1489_v75
				_ = _rhs273
				var _rhs274 D1 = _1490_v76
				_ = _rhs274
				_1478_v65 = _rhs270
				_1478_v65 = _rhs271
				_1478_v65 = _rhs272
				_1489_v75 = _rhs273
				_1490_v76 = _rhs274
				_1478_v65 = (_1478_v65) == (p2)
			}
		}
		var _1491_v77 D15
		_ = _1491_v77
		_1491_v77 = Companion_D15_.Create_DC28_(p2, true)
		var _1492_v78 D15
		_ = _1492_v78
		_1492_v78 = Companion_D15_.Create_DC29_(_1491_v77)
		var _source21 D15 = _1492_v78
		_ = _source21
		if _source21.Is_DC28() {
			var _1493___mcc_h7 bool = _source21.Get_().(D15_DC28).Cf36
			_ = _1493___mcc_h7
			var _1494___mcc_h8 bool = _source21.Get_().(D15_DC28).Cf37
			_ = _1494___mcc_h8
			var _1495_cf37 bool = _1494___mcc_h8
			_ = _1495_cf37
			var _1496_cf36 bool = _1493___mcc_h7
			_ = _1496_cf36
			var _1497_v79 _dafny.Array
			_ = _1497_v79
			var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw240
			_1497_v79 = _nw240
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1497_v79), 0))
			_ = _index260
			(_1497_v79).ArraySet1(_this.F15, (_index260).Int())
			var _1498_v80 _dafny.Sequence
			_ = _1498_v80
			_1498_v80 = _dafny.SeqOf(_dafny.IntOfInt64(-106))
			var _1499_v81 _dafny.MultiSet
			_ = _1499_v81
			_1499_v81 = _dafny.MultiSetOf((_this.F15).Times(_this.F15), (_1498_v80).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.IntOfUint32((_1498_v80).Cardinality()))).Uint32()).(_dafny.Int), _this.F15)
			var _1500_v82 _dafny.Map
			_ = _1500_v82
			_1500_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1496_cf36)
			var _1501_v83 D12
			_ = _1501_v83
			_1501_v83 = Companion_D12_.Create_DC21_(_1500_v82)
			var _1502_v84 D2
			_ = _1502_v84
			_1502_v84 = Companion_D2_.Create_DC7_(_this.F15, (_this).Fm7(Companion_Default___.Fm0(_1495_cf37, globalState), globalState))
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1497_v79), 0))
			_ = _index261
			var _rhs275 bool = p0
			_ = _rhs275
			var _rhs276 _dafny.Int = (_dafny.Zero).Minus(_this.F15)
			_ = _rhs276
			var _rhs277 _dafny.MultiSet = ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_1419_v25, _dafny.Companion_Sequence_.Update(_1419_v25, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(746), _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32(), (_this).F5()))).Cardinality()), _this.F15, _dafny.IntOfInt64(597))).Union(_1499_v81)).Union(_dafny.MultiSetOf(_this.F15, _dafny.IntOfInt64(-663), _dafny.IntOfInt64(804)))
			_ = _rhs277
			var _rhs278 _dafny.Int = (_1502_v84).Dtor_cf7()
			_ = _rhs278
			var _rhs279 D12 = _1501_v83
			_ = _rhs279
			var _lhs188 _dafny.Array = _1497_v79
			_ = _lhs188
			var _lhs189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1497_v79), 0))
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			_1496_cf36 = _rhs275
			(_lhs188).ArraySet1(_rhs276, (_lhs189).Int())
			_1499_v81 = _rhs277
			_lhs190.F1 = _rhs278
			_1501_v83 = _rhs279
			var _1503_v85 *C3
			_ = _1503_v85
			var _nw241 *C3 = New_C3_()
			_ = _nw241
			_nw241.Ctor__(_1495_cf37, p2)
			_1503_v85 = _nw241
			var _1504_v86 D11
			_ = _1504_v86
			_1504_v86 = Companion_D11_.Create_DC20_((_this).F6(), (_1497_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_1497_v79), 0))).Int()).(_dafny.Int))
			var _1505_v87 T4
			_ = _1505_v87
			var _nw242 *C2 = New_C2_()
			_ = _nw242
			_nw242.Ctor__((_this).F5(), (_1504_v86).Dtor_cf21())
			_1505_v87 = _nw242
			_1505_v87 = _1505_v87
			var _1506_v88 _dafny.Array
			_ = _1506_v88
			var _nw243 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw243
			_1506_v88 = _nw243
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1506_v88), 0))
			_ = _index262
			(_1506_v88).ArraySet1(_1496_cf36, (_index262).Int())
			var _1507_v89 _dafny.Array
			_ = _1507_v89
			var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw244
			_1507_v89 = _nw244
			var _1508_v90 _dafny.Array
			_ = _1508_v90
			var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw245
			_1508_v90 = _nw245
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_1507_v89), 0))
			_ = _index263
			(_1507_v89).ArraySet1(_1508_v90, (_index263).Int())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1506_v88), 0))
			_ = _index264
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_1507_v89), 0))
			_ = _index265
			var _rhs280 bool = !(_1496_cf36)
			_ = _rhs280
			var _rhs281 _dafny.Array = _1508_v90
			_ = _rhs281
			var _lhs191 _dafny.Array = _1506_v88
			_ = _lhs191
			var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(764), _dafny.ArrayLen((_1506_v88), 0))
			_ = _lhs192
			var _lhs193 _dafny.Array = _1507_v89
			_ = _lhs193
			var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_1507_v89), 0))
			_ = _lhs194
			(_lhs191).ArraySet1(_rhs280, (_lhs192).Int())
			(_lhs193).ArraySet1(_rhs281, (_lhs194).Int())
		} else if _source21.Is_DC27() {
			var _1509___mcc_h9 _dafny.MultiSet = _source21.Get_().(D15_DC27).Cf35
			_ = _1509___mcc_h9
			var _1510_cf35 _dafny.MultiSet = _1509___mcc_h9
			_ = _1510_cf35
			var _1511_v91 _dafny.Array
			_ = _1511_v91
			var _nw246 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw246
			_1511_v91 = _nw246
			var _1512_v92 _dafny.Array
			_ = _1512_v92
			_1512_v92 = _1511_v91
			var _1513_v93 _dafny.Sequence
			_ = _1513_v93
			_1513_v93 = _dafny.SeqOf(_1512_v92)
			_1513_v93 = _dafny.Companion_Sequence_.Concatenate(_1513_v93, _dafny.Companion_Sequence_.Concatenate(_1513_v93, _1513_v93))
			var _1514_v94 *C1
			_ = _1514_v94
			var _nw247 *C1 = New_C1_()
			_ = _nw247
			_nw247.Ctor__()
			_1514_v94 = _nw247
			var _1515_v95 D11
			_ = _1515_v95
			_1515_v95 = Companion_D11_.Create_DC20_(p0, _this.F15)
			var _1516_v96 _dafny.Sequence
			_ = _1516_v96
			_1516_v96 = _dafny.SeqOf(_this.F15)
			var _1517_v97 _dafny.MultiSet
			_ = _1517_v97
			_1517_v97 = _dafny.MultiSetOf((_this).Fm7(p0, globalState), _this.F15)
			var _1518_v98 _dafny.Array
			_ = _1518_v98
			var _nwElement0_45 _dafny.Int = _this.F15
			_ = _nwElement0_45
			var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(21))
			_ = _nw248
			(_nw248).ArraySet1(_nwElement0_45, 0)
			(_nw248).ArraySet1(_this.F15, 1)
			(_nw248).ArraySet1(_this.F15, 2)
			(_nw248).ArraySet1((_1515_v95).Dtor_cf22(), 3)
			(_nw248).ArraySet1((_1516_v96).Select((Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1516_v96).Cardinality()))).Uint32()).(_dafny.Int), 4)
			(_nw248).ArraySet1(_dafny.IntOfInt64(632), 5)
			(_nw248).ArraySet1((_dafny.Zero).Minus(_this.F15), 6)
			(_nw248).ArraySet1(_dafny.IntOfInt64(725), 7)
			(_nw248).ArraySet1(_dafny.IntOfInt64(223), 8)
			(_nw248).ArraySet1(_dafny.IntOfInt64(898), 9)
			(_nw248).ArraySet1(_this.F15, 10)
			(_nw248).ArraySet1(_this.F15, 11)
			(_nw248).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F15, _this.F15), 12)
			(_nw248).ArraySet1(_this.F15, 13)
			(_nw248).ArraySet1((_1514_v94).Fm14(globalState), 14)
			(_nw248).ArraySet1(((func() _dafny.Int {
				if (_1517_v97).Contains((_1514_v94).Fm4(_1415_v21, _dafny.IntOfInt64(832), _this.F15, _1517_v97, globalState)) {
					return (_1517_v97).Multiplicity((_1514_v94).Fm4(_1415_v21, _dafny.IntOfInt64(832), _this.F15, _1517_v97, globalState))
				}
				return (_this).Fm7(p2, globalState)
			})()).Minus(_this.F15), 15)
			(_nw248).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-5))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg88 _dafny.Int) interface{} {
					return coer88(arg88)
				}
			}(func(_1519_i5 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			})), _1419_v25)).Cardinality()), 16)
			(_nw248).ArraySet1(_this.F15, 17)
			(_nw248).ArraySet1((_this.F15).Times(_this.F15), 18)
			(_nw248).ArraySet1(_this.F15, 19)
			(_nw248).ArraySet1((_dafny.Zero).Minus((_1516_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(437), _dafny.IntOfUint32((_1516_v96).Cardinality()))).Uint32()).(_dafny.Int)), 20)
			_1518_v98 = _nw248
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _index266
			(_1518_v98).ArraySet1(_dafny.IntOfInt64(-44), (_index266).Int())
			var _1520_v99 bool
			_ = _1520_v99
			_1520_v99 = true
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _index267
			var _rhs282 _dafny.Int = ((_this).Fm8((_this).F6(), p2, p2, globalState)).Minus(_dafny.IntOfInt64(-696))
			_ = _rhs282
			var _rhs283 _dafny.Int = (_1514_v94).Fm7(p2, globalState)
			_ = _rhs283
			var _rhs284 bool = (Companion_Default___.Fm3(p2, (_this).F6(), globalState)).Cmp(_this.F15) != 0
			_ = _rhs284
			var _rhs285 _dafny.Int = (_dafny.Zero).Minus((_1516_v96).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1516_v96, _1516_v96), (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1516_v96, _1516_v96)).Cardinality()))).Uint32(), _this.F15)).Cardinality()), _dafny.IntOfUint32((_1516_v96).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs285
			var _lhs195 *GlobalState = globalState
			_ = _lhs195
			var _lhs196 _dafny.Array = _1518_v98
			_ = _lhs196
			var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _lhs197
			var _lhs198 *GlobalState = globalState
			_ = _lhs198
			_lhs195.F1 = _rhs282
			(_lhs196).ArraySet1(_rhs283, (_lhs197).Int())
			_1520_v99 = _rhs284
			_lhs198.F1 = _rhs285
			var _1521_v100 _dafny.MultiSet
			_ = _1521_v100
			_1521_v100 = _dafny.MultiSetOf(_1511_v91, _1511_v91, _1511_v91)
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _index268
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _index269
			var _rhs286 _dafny.Int = (func() _dafny.Int {
				if (_1521_v100).Contains((func() _dafny.Array {
					if (_this).F6() {
						return _1511_v91
					}
					return _1511_v91
				})()) {
					return (_1521_v100).Multiplicity((func() _dafny.Array {
						if (_this).F6() {
							return _1511_v91
						}
						return _1511_v91
					})())
				}
				return _this.F15
			})()
			_ = _rhs286
			var _rhs287 _dafny.Int = (_1518_v98).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))).Int()).(_dafny.Int)
			_ = _rhs287
			var _lhs199 _dafny.Array = _1518_v98
			_ = _lhs199
			var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _lhs200
			var _lhs201 _dafny.Array = _1518_v98
			_ = _lhs201
			var _lhs202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1518_v98), 0))
			_ = _lhs202
			(_lhs199).ArraySet1(_rhs286, (_lhs200).Int())
			(_lhs201).ArraySet1(_rhs287, (_lhs202).Int())
		} else {
			var _1522___mcc_h10 D15 = _source21.Get_().(D15_DC29).Cf38
			_ = _1522___mcc_h10
			var _1523_cf38 D15 = _1522___mcc_h10
			_ = _1523_cf38
			var _1524_v101 bool
			_ = _1524_v101
			_1524_v101 = true
			var _1525_v102 _dafny.Sequence
			_ = _1525_v102
			_1525_v102 = _dafny.SeqOf(_this.F15, _this.F15)
			var _rhs288 bool = (_this).F6()
			_ = _rhs288
			var _rhs289 _dafny.Int = _this.F15
			_ = _rhs289
			var _rhs290 _dafny.Int = (_dafny.IntOfUint32((_1525_v102).Cardinality())).Times(Companion_Default___.SafeDivisionInt(_this.F15, _this.F15))
			_ = _rhs290
			var _rhs291 _dafny.Int = _dafny.IntOfInt64(148)
			_ = _rhs291
			var _lhs203 *GlobalState = globalState
			_ = _lhs203
			var _lhs204 *GlobalState = globalState
			_ = _lhs204
			var _lhs205 *GlobalState = globalState
			_ = _lhs205
			_1524_v101 = _rhs288
			_lhs203.F1 = _rhs289
			_lhs204.F1 = _rhs290
			_lhs205.F1 = _rhs291
			var _1526_v103 _dafny.Set
			_ = _1526_v103
			_1526_v103 = _dafny.SetOf(p2)
			_1524_v101 = (_dafny.SetOf(p0, _1524_v101)).IsProperSubsetOf(_1526_v103)
			_1524_v101 = (_1415_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_this.F15))).Cardinality()), _dafny.IntOfUint32((_1415_v21).Cardinality()))).Uint32()).(bool)
			var _1527_v104 _dafny.MultiSet
			_ = _1527_v104
			_1527_v104 = _dafny.MultiSetOf(_this.F15, _this.F15)
			var _1528_v105 _dafny.Array
			_ = _1528_v105
			var _nwElement0_46 _dafny.MultiSet = _1527_v104
			_ = _nwElement0_46
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(20))
			_ = _nw249
			(_nw249).ArraySet1(_nwElement0_46, 0)
			(_nw249).ArraySet1(_1527_v104, 1)
			(_nw249).ArraySet1(_1527_v104, 2)
			(_nw249).ArraySet1((_1527_v104).Intersection(_1527_v104), 3)
			(_nw249).ArraySet1(_1527_v104, 4)
			(_nw249).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1525_v102, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(782))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg89 _dafny.Int) interface{} {
					return coer89(arg89)
				}
			}(func(_1529_i6 _dafny.Int) _dafny.Int {
				return _this.F15
			})))), 5)
			(_nw249).ArraySet1(_1527_v104, 6)
			(_nw249).ArraySet1(_1527_v104, 7)
			(_nw249).ArraySet1((Companion_Default___.Fm45(globalState)).Union(_1527_v104), 8)
			(_nw249).ArraySet1(_1527_v104, 9)
			(_nw249).ArraySet1(_1527_v104, 10)
			(_nw249).ArraySet1(_1527_v104, 11)
			(_nw249).ArraySet1(_1527_v104, 12)
			(_nw249).ArraySet1(_1527_v104, 13)
			(_nw249).ArraySet1(_1527_v104, 14)
			(_nw249).ArraySet1((_dafny.MultiSetOf(_this.F15, _this.F15)).Union(_1527_v104), 15)
			(_nw249).ArraySet1((Companion_Default___.Fm45(globalState)).Intersection(_1527_v104), 16)
			(_nw249).ArraySet1(_1527_v104, 17)
			(_nw249).ArraySet1(_1527_v104, 18)
			(_nw249).ArraySet1(_1527_v104, 19)
			_1528_v105 = _nw249
			var _rhs292 bool = false
			_ = _rhs292
			var _rhs293 _dafny.Int = _dafny.IntOfInt64(-677)
			_ = _rhs293
			var _rhs294 _dafny.Array = _1528_v105
			_ = _rhs294
			var _rhs295 bool = _1524_v101
			_ = _rhs295
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			_1524_v101 = _rhs292
			_lhs206.F1 = _rhs293
			_1528_v105 = _rhs294
			_1524_v101 = _rhs295
		}
		var _1530_v106 bool
		_ = _1530_v106
		_1530_v106 = false
		var _1531_v107 _dafny.Array
		_ = _1531_v107
		var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw250
		_1531_v107 = _nw250
		var _1532_v108 _dafny.Map
		_ = _1532_v108
		_1532_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1531_v107, (_this.F15).Plus(_this.F15))
		var _1533_v109 _dafny.Sequence
		_ = _1533_v109
		_1533_v109 = _dafny.SeqOf(_this.F15)
		var _1534_v110 D14
		_ = _1534_v110
		_1534_v110 = Companion_D14_.Create_DC26_((_this).Fm6(_1533_v109, globalState), _this.F15, (_this).F5())
		var _rhs296 _dafny.Int = (_1534_v110).Dtor_cf33()
		_ = _rhs296
		var _rhs297 bool = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1419_v25, _1419_v25), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1419_v25, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32(), (_this).F5()), _1419_v25), (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1419_v25, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32(), (_this).F5()), _1419_v25)).Cardinality()))).Uint32(), (_this).F5()), (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1419_v25, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32(), (_this).F5()), _1419_v25), (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1419_v25, (Companion_Default___.SafeIndex(_this.F15, _dafny.IntOfUint32((_1419_v25).Cardinality()))).Uint32(), (_this).F5()), _1419_v25)).Cardinality()))).Uint32(), (_this).F5())).Cardinality()))).Uint32(), (_this).F5()))
		_ = _rhs297
		var _rhs298 _dafny.Map = (_1532_v108).Merge(_1532_v108)
		_ = _rhs298
		var _lhs207 *C5 = _this
		_ = _lhs207
		_lhs207.F15 = _rhs296
		_1530_v106 = _rhs297
		_1532_v108 = _rhs298
		var _1535_v111 D6
		_ = _1535_v111
		_1535_v111 = Companion_D6_.Create_DC13_(_1533_v109)
		var _pat_let_tv19 = _1533_v109
		_ = _pat_let_tv19
		_1535_v111 = func(_pat_let18_0 D6) D6 {
			return func(_1536_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let19_0 _dafny.Sequence) D6 {
					return func(_1537_dt__update_hcf14_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC13_(_1537_dt__update_hcf14_h0)
					}(_pat_let19_0)
				}(_pat_let_tv19)
			}(_pat_let18_0)
		}(_1535_v111)
	}
}
func (_this *C5) M12(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Map, _dafny.Sequence, _dafny.Set, bool) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Set = _dafny.EmptySet
		_ = r2
		var r3 bool = false
		_ = r3
		var _1538_v0 _dafny.Array
		_ = _1538_v0
		var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
		_ = _nw251
		_1538_v0 = _nw251
		var _1539_v1 _dafny.Sequence
		_ = _1539_v1
		_1539_v1 = _dafny.UnicodeSeqOfUtf8Bytes("d")
		var _1540_v2 _dafny.Map
		_ = _1540_v2
		_1540_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_1539_v1).Cardinality()))
		var _1541_v3 _dafny.Map
		_ = _1541_v3
		_1541_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F15, (_this).F6())
		var _1542_v4 D2
		_ = _1542_v4
		_1542_v4 = Companion_D2_.Create_DC7_(_this.F15, (_1541_v3).Cardinality())
		var _1543_v5 _dafny.Sequence
		_ = _1543_v5
		_1543_v5 = _dafny.SeqOf(_this.F15)
		var _1544_v6 _dafny.MultiSet
		_ = _1544_v6
		_1544_v6 = _dafny.MultiSetOf((_this).F6())
		var _1545_v7 _dafny.Map
		_ = _1545_v7
		_1545_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), p1)
		var _1546_v8 _dafny.Array
		_ = _1546_v8
		var _nwElement0_47 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1538_v0, _1538_v0)).Cardinality()))
		_ = _nwElement0_47
		var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(24))
		_ = _nw252
		(_nw252).ArraySet1(_nwElement0_47, 0)
		(_nw252).ArraySet1(p2, 1)
		(_nw252).ArraySet1(Companion_Default___.SafeDivisionInt(_this.F15, p2), 2)
		(_nw252).ArraySet1((func() _dafny.Int {
			if p1 {
				return (func() _dafny.Int {
					if (_1540_v2).Contains((_this).F6()) {
						return (_1540_v2).Get((_this).F6()).(_dafny.Int)
					}
					return (_1542_v4).Dtor_cf8()
				})()
			}
			return p2
		})(), 3)
		(_nw252).ArraySet1(Companion_Default___.Fm3((_this).F6(), p1, globalState), 4)
		(_nw252).ArraySet1(_dafny.IntOfInt64(-81), 5)
		(_nw252).ArraySet1(p2, 6)
		(_nw252).ArraySet1((_1543_v5).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)).Cardinality(), _dafny.IntOfUint32((_1543_v5).Cardinality()))).Uint32()).(_dafny.Int), 7)
		(_nw252).ArraySet1(p2, 8)
		(_nw252).ArraySet1(_this.F15, 9)
		(_nw252).ArraySet1(((func() _dafny.Int {
			if (_1544_v6).Contains((_this).F6()) {
				return (_1544_v6).Multiplicity((_this).F6())
			}
			return _this.F15
		})()).Minus(p2), 10)
		(_nw252).ArraySet1((_1543_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-258), _dafny.IntOfUint32((_1543_v5).Cardinality()))).Uint32()).(_dafny.Int), 11)
		(_nw252).ArraySet1(p2, 12)
		(_nw252).ArraySet1((_this.F15).Times((_1545_v7).Cardinality()), 13)
		(_nw252).ArraySet1((func() _dafny.Int {
			if (_1544_v6).Contains(false) {
				return (_1544_v6).Multiplicity(false)
			}
			return _this.F15
		})(), 14)
		(_nw252).ArraySet1(_this.F15, 15)
		(_nw252).ArraySet1((func() _dafny.Int {
			if p1 {
				return p0
			}
			return p0
		})(), 16)
		(_nw252).ArraySet1(_dafny.IntOfInt64(886), 17)
		(_nw252).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm1(p0, (_this).F6(), _dafny.IntOfInt64(-623), globalState)).Cardinality()), 18)
		(_nw252).ArraySet1((p2).Minus(p2), 19)
		(_nw252).ArraySet1((_1540_v2).Cardinality(), 20)
		(_nw252).ArraySet1(_this.F15, 21)
		(_nw252).ArraySet1((_dafny.Zero).Minus(p0), 22)
		(_nw252).ArraySet1(p2, 23)
		_1546_v8 = _nw252
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))
		_ = _index270
		(_1546_v8).ArraySet1(Companion_Default___.SafeModuloInt(p0, _this.F15), (_index270).Int())
		var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))
		_ = _index271
		(_1546_v8).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(364))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}(func(_1547_i0 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		}))).Cardinality())).Minus(p2), (_index271).Int())
		var _1548_i1 _dafny.Int
		_ = _1548_i1
		_1548_i1 = _dafny.Zero
		{
			for p1 {
				{
					if (_1548_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_1548_i1 = (_1548_i1).Plus(_dafny.One)
					var _1549_v9 _dafny.Sequence
					_ = _1549_v9
					_1549_v9 = _dafny.SeqOf(_1543_v5)
					var _1550_v10 *C0
					_ = _1550_v10
					var _nw253 *C0 = New_C0_()
					_ = _nw253
					_nw253.Ctor__((_this).F6(), _1549_v9)
					_1550_v10 = _nw253
					r3 = true
					r3 = (_this).F6()
					var _1551_v11 _dafny.Array
					_ = _1551_v11
					var _nw254 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
					_ = _nw254
					_1551_v11 = _nw254
					var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_1551_v11), 0))
					_ = _index272
					(_1551_v11).ArraySet1(p1, (_index272).Int())
					var _1552_v13 _dafny.Map
					_ = _1552_v13
					_1552_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(723), _1539_v1)
					var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(272), _dafny.ArrayLen((_1551_v11), 0))
					_ = _index273
					(_1551_v11).ArraySet1((func() _dafny.Map {
						var _coll51 = _dafny.NewMapBuilder()
						_ = _coll51
						for _iter52 := _dafny.Iterate((Companion_Default___.Fm56(globalState)).Elements()); ; {
							_compr_51, _ok52 := _iter52()
							if !_ok52 {
								break
							}
							var _1553_v12 _dafny.Int
							_1553_v12 = interface{}(_compr_51).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm56(globalState), _1553_v12) {
								_coll51.Add(Companion_Default___.SafeModuloInt(_1553_v12, p0), _1539_v1)
							}
						}
						return _coll51.ToMap()
					}()).Equals(_1552_v13), (_index273).Int())
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))
		_ = _index274
		(_1546_v8).ArraySet1(Companion_Default___.SafeModuloInt(_this.F15, Companion_Default___.SafeModuloInt((func() _dafny.Set {
			var _coll52 = _dafny.NewBuilder()
			_ = _coll52
			for _iter53 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(804), _dafny.IntOfInt64(781))); ; {
				_compr_52, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _1554_v14 _dafny.Int
				_1554_v14 = interface{}(_compr_52).(_dafny.Int)
				if ((_dafny.IntOfInt64(804)).Cmp(_1554_v14) <= 0) && ((_1554_v14).Cmp(_dafny.IntOfInt64(781)) < 0) {
					_coll52.Add((_1554_v14).Times((_dafny.Zero).Minus(p2)))
				}
			}
			return _coll52.ToSet()
		}()).Cardinality(), p0)), (_index274).Int())
		var _1555_v15 _dafny.Array
		_ = _1555_v15
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_44
		var _nw255 _dafny.Array
		_ = _nw255
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw255 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) bool = func(_1556_i2 _dafny.Int) bool {
				return (_this).F6()
			}
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw255 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw255).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw255).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1555_v15 = _nw255
		var _1557_v16 _dafny.Array
		_ = _1557_v16
		var _nwElement0_48 _dafny.CodePoint = (_this).F5()
		_ = _nwElement0_48
		var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(17))
		_ = _nw256
		(_nw256).ArraySet1CodePoint(_nwElement0_48, 0)
		(_nw256).ArraySet1CodePoint((_this).F5(), 1)
		(_nw256).ArraySet1CodePoint((_this).F5(), 2)
		(_nw256).ArraySet1CodePoint(_dafny.CodePoint('u'), 3)
		(_nw256).ArraySet1CodePoint((_this).F5(), 4)
		(_nw256).ArraySet1CodePoint((_this).F5(), 5)
		(_nw256).ArraySet1CodePoint((_this).F5(), 6)
		(_nw256).ArraySet1CodePoint((_this).F5(), 7)
		(_nw256).ArraySet1CodePoint((_this).F5(), 8)
		(_nw256).ArraySet1CodePoint((_this).F5(), 9)
		(_nw256).ArraySet1CodePoint((_this).F5(), 10)
		(_nw256).ArraySet1CodePoint((_this).F5(), 11)
		(_nw256).ArraySet1CodePoint((_this).F5(), 12)
		(_nw256).ArraySet1CodePoint((_this).F5(), 13)
		(_nw256).ArraySet1CodePoint((_this).F5(), 14)
		(_nw256).ArraySet1CodePoint((_this).F5(), 15)
		(_nw256).ArraySet1CodePoint((func(_pat_let20_0 D0) D0 {
			return func(_1558_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let21_0 _dafny.CodePoint) D0 {
					return func(_1559_dt__update_hcf2_h0 _dafny.CodePoint) D0 {
						return Companion_D0_.Create_DC2_((_1558_dt__update__tmp_h0).Dtor_cf1(), _1559_dt__update_hcf2_h0)
					}(_pat_let21_0)
				}((_this).F5())
			}(_pat_let20_0)
		}(Companion_D0_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1555_v15, (_this).F5()), (_this).F5()))).Dtor_cf2(), 16)
		_1557_v16 = _nw256
		var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1557_v16), 0))
		_ = _index275
		(_1557_v16).ArraySet1CodePoint((_this).F5(), (_index275).Int())
		var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1557_v16), 0))
		_ = _index276
		(_1557_v16).ArraySet1CodePoint((_this).F5(), (_index276).Int())
		var _1560_i3 _dafny.Int
		_ = _1560_i3
		_1560_i3 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_1560_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L15
					}
					_1560_i3 = (_1560_i3).Plus(_dafny.One)
					var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1557_v16), 0))
					_ = _index277
					(_1557_v16).ArraySet1CodePoint(_dafny.CodePoint('p'), (_index277).Int())
					var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1555_v15), 0))
					_ = _index278
					(_1555_v15).ArraySet1((func() bool {
						if true {
							return (_this).F6()
						}
						return p1
					})(), (_index278).Int())
					var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1555_v15), 0))
					_ = _index279
					(_1555_v15).ArraySet1((_this).F6(), (_index279).Int())
					r3 = (_1555_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(519), _dafny.ArrayLen((_1555_v15), 0))).Int()).(bool)
					var _1561_v17 _dafny.Set
					_ = _1561_v17
					_1561_v17 = _dafny.SetOf((_1546_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))).Int()).(_dafny.Int), _this.F15)
					var _1562_v18 _dafny.Set
					_ = _1562_v18
					_1562_v18 = _dafny.SetOf((_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(857))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg91 _dafny.Int) interface{} {
							return coer91(arg91)
						}
					}(func(_1563_i4 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					}))).Cardinality())).Times((_1561_v17).Cardinality())), p0)
					_1561_v17 = _1562_v18
					goto C15
				}
			C15:
			}
			goto L15
		}
	L15:
		var _1564_v19 _dafny.Array
		_ = _1564_v19
		var _len0_45 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_45
		var _nw257 _dafny.Array
		_ = _nw257
		if _len0_45.Cmp(_dafny.Zero) == 0 {
			_nw257 = _dafny.NewArray(_len0_45)
		} else {
			var _init45 func(_dafny.Int) _dafny.Set = (func(_1565_v8 _dafny.Array) func(_dafny.Int) _dafny.Set {
				return func(_1566_i5 _dafny.Int) _dafny.Set {
					return (_dafny.SetOf((_1565_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1565_v8), 0))).Int()).(_dafny.Int))).Intersection(_dafny.SetOf(_this.F15, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())).Cardinality()))
				}
			})(_1546_v8)
			_ = _init45
			var _element0_45 = _init45(_dafny.Zero)
			_ = _element0_45
			_nw257 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
			(_nw257).ArraySet1(_element0_45, 0)
			var _nativeLen0_45 = (_len0_45).Int()
			_ = _nativeLen0_45
			for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
				(_nw257).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
			}
		}
		_1564_v19 = _nw257
		var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))
		_ = _index280
		var _rhs299 _dafny.Array = _1564_v19
		_ = _rhs299
		var _rhs300 bool = p1
		_ = _rhs300
		var _rhs301 bool = (_this).F6()
		_ = _rhs301
		var _rhs302 _dafny.Int = _this.F15
		_ = _rhs302
		var _rhs303 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-24))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg92 _dafny.Int) interface{} {
				return coer92(arg92)
			}
		}((func(_1567_v16 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_1568_i6 _dafny.Int) _dafny.CodePoint {
				return (_1567_v16).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1567_v16), 0))).Int())
			}
		})(_1557_v16))), (Companion_Default___.SafeIndex((_1546_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-24))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}((func(_1569_v16 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
			return func(_1570_i6 _dafny.Int) _dafny.CodePoint {
				return (_1569_v16).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(366), _dafny.ArrayLen((_1569_v16), 0))).Int())
			}
		})(_1557_v16)))).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
			if p1 {
				return (_this).F5()
			}
			return (_this).F5()
		})())).Cardinality())
		_ = _rhs303
		var _lhs208 _dafny.Array = _1546_v8
		_ = _lhs208
		var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))
		_ = _lhs209
		var _lhs210 *C5 = _this
		_ = _lhs210
		_1564_v19 = _rhs299
		r3 = _rhs300
		r3 = _rhs301
		(_lhs208).ArraySet1(_rhs302, (_lhs209).Int())
		_lhs210.F15 = _rhs303
		var _1571_v20 _dafny.Map
		_ = _1571_v20
		_1571_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_dafny.IntOfUint32((_1539_v1).Cardinality()), (_this).Fm6(_1543_v5, globalState), (_1546_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))).Int()).(_dafny.Int), globalState), p0)
		r0 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vudfnvin"), p0)).Merge(_1571_v20)
		r1 = (func() _dafny.Sequence {
			if p1 {
				return _dafny.Companion_Sequence_.Concatenate(_1543_v5, _1543_v5)
			}
			return _1543_v5
		})()
		var _1572_v21 _dafny.Set
		_ = _1572_v21
		_1572_v21 = _dafny.SetOf((_this).F6())
		r2 = (_1572_v21).Difference(Companion_Default___.Fm49(globalState))
		var _1573_v22 _dafny.Sequence
		_ = _1573_v22
		_1573_v22 = _dafny.SeqOf(_1538_v0)
		var _1574_v23 _dafny.Sequence
		_ = _1574_v23
		_1574_v23 = _dafny.SeqOf(_1539_v1)
		var _1575_v24 _dafny.Map
		_ = _1575_v24
		_1575_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1(_dafny.IntOfInt64(138), (_this).F6(), _dafny.IntOfUint32((_1573_v22).Cardinality()), globalState), (_1574_v23).Select((Companion_Default___.SafeIndex((_1546_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_1546_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1574_v23).Cardinality()))).Uint32()).(_dafny.Sequence)), (_this).F6())
		r3 = (func() bool {
			if (_1575_v24).Contains(_dafny.UnicodeSeqOfUtf8Bytes("rdv")) {
				return (_1575_v24).Get(_dafny.UnicodeSeqOfUtf8Bytes("rdv")).(bool)
			}
			return false
		})()
		return r0, r1, r2, r3
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	dummy byte
}

func New_C6_() *C6 {
	_this := C6{}

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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C6{}
var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__() {
	{
	}
}
func (_this *C6) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !(false) || (false)
	}
}
func (_this *C6) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.IntOfInt64(982)).Minus(_dafny.IntOfInt64(571))).Times(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((func() _dafny.Map {
			var _coll53 = _dafny.NewMapBuilder()
			_ = _coll53
			for _iter54 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('k'))).Elements()); ; {
				_compr_53, _ok54 := _iter54()
				if !_ok54 {
					break
				}
				var _1576_v0 _dafny.CodePoint
				_1576_v0 = interface{}(_compr_53).(_dafny.CodePoint)
				if (_dafny.SetOf(_dafny.CodePoint('k'))).Contains(_1576_v0) {
					_coll53.Add(_1576_v0, true)
				}
			}
			return _coll53.ToMap()
		}()).Cardinality()), _dafny.IntOfInt64(708)))
	}
}
func (_this *C6) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return (Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nyudstnyx")).Cardinality()), _dafny.IntOfInt64(832))).Cardinality(), false)).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(145))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}(func(_1577_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality())))).Plus(_dafny.IntOfInt64(111))
	}
}
func (_this *C6) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(276), _dafny.IntOfInt64(330)))
	}
}
func (_this *C6) Fm29(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus(((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality()).Times(_dafny.IntOfInt64(568)))).Minus(_dafny.IntOfInt64(215))
	}
}
func (_this *C6) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		r0 = !(!(p0)) || (p0)
		var _1578_v0 D0
		_ = _1578_v0
		_1578_v0 = Companion_D0_.Create_DC1_()
		var _1579_v1 _dafny.Set
		_ = _1579_v1
		_1579_v1 = _dafny.SetOf(_1578_v0, _1578_v0, _1578_v0)
		var _1580_v2 _dafny.Set
		_ = _1580_v2
		_1580_v2 = _1579_v1
		var _source22 _dafny.Set = _1580_v2
		_ = _source22
		var _1581___mcc_h0 _dafny.Set = _source22
		_ = _1581___mcc_h0
		var _1582_cf17 _dafny.Set = _1581___mcc_h0
		_ = _1582_cf17
		(globalState).F1 = p1
		r0 = (p1).Cmp(p1) <= 0
		var _1583_v3 *C1
		_ = _1583_v3
		var _nw258 *C1 = New_C1_()
		_ = _nw258
		_nw258.Ctor__()
		_1583_v3 = _nw258
		var _1584_v4 _dafny.Sequence
		_ = _1584_v4
		_1584_v4 = _dafny.SeqOf(p1)
		var _1585_v5 _dafny.Map
		_ = _1585_v5
		_1585_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1586_v6 _dafny.Array
		_ = _1586_v6
		var _nwElement0_49 bool = p0
		_ = _nwElement0_49
		var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(28))
		_ = _nw259
		(_nw259).ArraySet1(_nwElement0_49, 0)
		(_nw259).ArraySet1(p0, 1)
		(_nw259).ArraySet1(p0, 2)
		(_nw259).ArraySet1(p0, 3)
		(_nw259).ArraySet1(p0, 4)
		(_nw259).ArraySet1(p0, 5)
		(_nw259).ArraySet1(false, 6)
		(_nw259).ArraySet1(true, 7)
		(_nw259).ArraySet1(p0, 8)
		(_nw259).ArraySet1(p0, 9)
		(_nw259).ArraySet1(p0, 10)
		(_nw259).ArraySet1(Companion_Default___.Fm0(false, globalState), 11)
		(_nw259).ArraySet1(p0, 12)
		(_nw259).ArraySet1(p0, 13)
		(_nw259).ArraySet1(p0, 14)
		(_nw259).ArraySet1((_this).Fm6(_1584_v4, globalState), 15)
		(_nw259).ArraySet1(p0, 16)
		(_nw259).ArraySet1((func() bool {
			if (_1585_v5).Contains(p0) {
				return (_1585_v5).Get(p0).(bool)
			}
			return p0
		})(), 17)
		(_nw259).ArraySet1((_1583_v3).Fm6(_1584_v4, globalState), 18)
		(_nw259).ArraySet1(p0, 19)
		(_nw259).ArraySet1(p0, 20)
		(_nw259).ArraySet1(p0, 21)
		(_nw259).ArraySet1(p0, 22)
		(_nw259).ArraySet1(p0, 23)
		(_nw259).ArraySet1(p0, 24)
		(_nw259).ArraySet1(p0, 25)
		(_nw259).ArraySet1(p0, 26)
		(_nw259).ArraySet1(p0, 27)
		_1586_v6 = _nw259
		var _1587_v7 _dafny.Map
		_ = _1587_v7
		_1587_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1583_v3, _1586_v6)
		var _1588_v8 _dafny.Map
		_ = _1588_v8
		_1588_v8 = _1587_v7
		var _source23 _dafny.Map = _1588_v8
		_ = _source23
		var _1589___mcc_h1 _dafny.Map = _source23
		_ = _1589___mcc_h1
		var _1590_cf18 _dafny.Map = _1589___mcc_h1
		_ = _1590_cf18
		var _1591_v9 T1
		_ = _1591_v9
		var _nw260 *C4 = New_C4_()
		_ = _nw260
		_nw260.Ctor__()
		_1591_v9 = _nw260
		var _1592_v10 _dafny.MultiSet
		_ = _1592_v10
		_1592_v10 = _dafny.MultiSetOf(_1591_v9)
		var _1593_v11 _dafny.Map
		_ = _1593_v11
		_1593_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Merge(_1585_v5), _1592_v10)
		var _1594_v12 _dafny.Sequence
		_ = _1594_v12
		_1594_v12 = _dafny.SeqOf(_1578_v0, _1578_v0)
		var _1595_v13 _dafny.Sequence
		_ = _1595_v13
		_1595_v13 = _dafny.SeqOf(_1592_v10)
		var _rhs304 bool = !(p0)
		_ = _rhs304
		var _rhs305 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1585_v5, _1592_v10)).Update(_1585_v5, (_1595_v13).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1595_v13).Cardinality()))).Uint32()).(_dafny.MultiSet))
		_ = _rhs305
		var _rhs306 _dafny.Sequence = _1594_v12
		_ = _rhs306
		r0 = _rhs304
		_1593_v11 = _rhs305
		_1594_v12 = _rhs306
		var _1596_v14 _dafny.CodePoint
		_ = _1596_v14
		_1596_v14 = _dafny.CodePoint('q')
		var _1597_v15 *C3
		_ = _1597_v15
		var _nw261 *C3 = New_C3_()
		_ = _nw261
		_nw261.Ctor__(Companion_Default___.Fm0(p0, globalState), !_dafny.Companion_Sequence_.Contains(p2, _1596_v14))
		_1597_v15 = _nw261
		var _1598_v16 _dafny.Sequence
		_ = _1598_v16
		_1598_v16 = _dafny.SeqOf(!((_1597_v15).F14()))
		var _1599_v17 _dafny.Array
		_ = _1599_v17
		var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
		_ = _nw262
		_1599_v17 = _nw262
		var _1600_v18 _dafny.Map
		_ = _1600_v18
		_1600_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1598_v16, _1599_v17)
		var _1601_v19 _dafny.Map
		_ = _1601_v19
		_1601_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1600_v18, (_1597_v15).F14())
		_1601_v19 = (_1601_v19).Update(_1600_v18, p0)
		var _1602_v20 _dafny.Map
		_ = _1602_v20
		_1602_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1598_v16, (_1597_v15).F13())
		var _1603_v21 _dafny.Array
		_ = _1603_v21
		var _nwElement0_50 _dafny.Map = (_1585_v5).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v15).F14(), (_1597_v15).F14()))
		_ = _nwElement0_50
		var _nw263 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(15))
		_ = _nw263
		(_nw263).ArraySet1(_nwElement0_50, 0)
		(_nw263).ArraySet1(_1585_v5, 1)
		(_nw263).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v15).F14(), p0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(p0))), 2)
		(_nw263).ArraySet1((_1585_v5).Merge((Companion_Default___.Fm41(p0, true, p0, globalState)).Update((_1597_v15).F14(), (_1597_v15).F14())), 3)
		(_nw263).ArraySet1(_1585_v5, 4)
		(_nw263).ArraySet1(_1585_v5, 5)
		(_nw263).ArraySet1(_1585_v5, 6)
		(_nw263).ArraySet1((_1585_v5).Merge(_1585_v5), 7)
		(_nw263).ArraySet1((_1585_v5).Update((_1597_v15).F13(), false), 8)
		(_nw263).ArraySet1(_1585_v5, 9)
		(_nw263).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_1602_v20).Contains(_1598_v16) {
				return (_1602_v20).Get(_1598_v16).(bool)
			}
			return p0
		})(), (_1597_v15).F13()), 10)
		(_nw263).ArraySet1(_1585_v5, 11)
		(_nw263).ArraySet1(_1585_v5, 12)
		(_nw263).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v15).F13(), (_1597_v15).F14())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v15).F13(), p0)), 13)
		(_nw263).ArraySet1(_1585_v5, 14)
		_1603_v21 = _nw263
		var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1603_v21), 0))
		_ = _index281
		(_1603_v21).ArraySet1(_1585_v5, (_index281).Int())
		var _1604_v22 _dafny.Set
		_ = _1604_v22
		_1604_v22 = _dafny.SetOf(p1, (_dafny.Zero).Minus(p1))
		var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_1603_v21), 0))
		_ = _index282
		(_1603_v21).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1597_v15).F14(), (_1604_v22).IsSubsetOf(_1604_v22)), (_index282).Int())
		var _1605_v23 _dafny.CodePoint
		_ = _1605_v23
		_1605_v23 = _dafny.CodePoint('a')
		var _1606_v24 _dafny.Map
		_ = _1606_v24
		_1606_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _1607_v25 *C4
		_ = _1607_v25
		var _nw264 *C4 = New_C4_()
		_ = _nw264
		_nw264.Ctor__()
		_1607_v25 = _nw264
		var _1608_v26 D14
		_ = _1608_v26
		_1608_v26 = Companion_D14_.Create_DC25_(_1607_v25)
		var _1609_v27 _dafny.Map
		_ = _1609_v27
		_1609_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1608_v26).Dtor_cf31())
		var _1610_v28 _dafny.Sequence
		_ = _1610_v28
		_1610_v28 = _dafny.SeqOf(p0)
		var _1611_v29 _dafny.MultiSet
		_ = _1611_v29
		_1611_v29 = _dafny.MultiSetOf(_1605_v23)
		var _1612_v30 _dafny.Array
		_ = _1612_v30
		var _nwElement0_51 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p2, _dafny.Companion_Sequence_.Update(p2, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((p2).Cardinality()))).Uint32(), _1605_v23))).Cardinality()))
		_ = _nwElement0_51
		var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(10))
		_ = _nw265
		(_nw265).ArraySet1(_nwElement0_51, 0)
		(_nw265).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(339), p1), 1)
		(_nw265).ArraySet1(p1, 2)
		(_nw265).ArraySet1(p1, 3)
		(_nw265).ArraySet1((_1584_v4).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_1606_v24).Contains(p0) {
				return (_1606_v24).Get(p0).(_dafny.Int)
			}
			return (_1609_v27).Cardinality()
		})(), _dafny.IntOfUint32((_1584_v4).Cardinality()))).Uint32()).(_dafny.Int), 4)
		(_nw265).ArraySet1(p1, 5)
		(_nw265).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.MultiSetOf((_1585_v5).Cardinality(), p1, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1610_v28).Cardinality())))).Cardinality()).Times(p1))), 6)
		(_nw265).ArraySet1(p1, 7)
		(_nw265).ArraySet1((p1).Times(p1), 8)
		(_nw265).ArraySet1((func() _dafny.Int {
			if (_1611_v29).Contains(_1605_v23) {
				return (_1611_v29).Multiplicity(_1605_v23)
			}
			return p1
		})(), 9)
		_1612_v30 = _nw265
		var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1612_v30), 0))
		_ = _index283
		(_1612_v30).ArraySet1(p1, (_index283).Int())
		var _1613_v31 _dafny.Set
		_ = _1613_v31
		_1613_v31 = _dafny.SetOf(p1, (_dafny.Zero).Minus(p1))
		var _1614_v32 _dafny.Map
		_ = _1614_v32
		_1614_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1613_v31).Cardinality())
		var _1615_v33 _dafny.Map
		_ = _1615_v33
		_1615_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1614_v32, p0)
		var _1616_v34 _dafny.Sequence
		_ = _1616_v34
		_1616_v34 = _dafny.SeqOf(_1610_v28)
		var _1617_v35 _dafny.Map
		_ = _1617_v35
		_1617_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1610_v28)
		var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1612_v30), 0))
		_ = _index284
		(_1612_v30).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm1(((_1607_v25).Fm30(_1615_v33, _dafny.MultiSetFromSeq(_1616_v34), globalState)).Times(p1), p0, _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1617_v35).Contains(p0) {
				return (_1617_v35).Get(p0).(_dafny.Sequence)
			}
			return _1610_v28
		})()).Cardinality()), globalState)).Cardinality()), (_index284).Int())
		var _1618_v36 _dafny.Array
		_ = _1618_v36
		var _nw266 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
		_ = _nw266
		_1618_v36 = _nw266
		for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1618_v36), 0))); ; {
			_guard_loop_1, _ok55 := _iter55()
			if !_ok55 {
				break
			}
			var _1619_i0 _dafny.Int
			_1619_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_1619_i0).Sign() != -1) && ((_1619_i0).Cmp(_dafny.ArrayLen((_1618_v36), 0)) < 0)) {
				(_1618_v36).ArraySet1(((_dafny.MultiSetOf(p0, p0)).Intersection(_dafny.MultiSetOf(p0))).Union((_dafny.MultiSetOf(p0)).Intersection(_dafny.MultiSetOf(p0, p0, p0))), (_1619_i0).Int())
			}
		}
		var _hi9 _dafny.Int = _dafny.IntOfUint32((p2).Cardinality())
		_ = _hi9
		for _1620_i1 := p1; _1620_i1.Cmp(_hi9) < 0; _1620_i1 = _1620_i1.Plus(_dafny.One) {
			r0 = p0
			var _1621_v37 _dafny.Array
			_ = _1621_v37
			var _len0_46 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_46
			var _nw267 _dafny.Array
			_ = _nw267
			if _len0_46.Cmp(_dafny.Zero) == 0 {
				_nw267 = _dafny.NewArray(_len0_46)
			} else {
				var _init46 func(_dafny.Int) bool = (func(_1622_p2 _dafny.Sequence) func(_dafny.Int) bool {
					return func(_1623_i2 _dafny.Int) bool {
						return !_dafny.Companion_Sequence_.Contains(_1622_p2, _dafny.CodePoint('p'))
					}
				})(p2)
				_ = _init46
				var _element0_46 = _init46(_dafny.Zero)
				_ = _element0_46
				_nw267 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
				(_nw267).ArraySet1(_element0_46, 0)
				var _nativeLen0_46 = (_len0_46).Int()
				_ = _nativeLen0_46
				for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
					(_nw267).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
				}
			}
			_1621_v37 = _nw267
			var _1624_v38 _dafny.Sequence
			_ = _1624_v38
			_1624_v38 = _dafny.SeqOf(false, p0, p0, true)
			var _1625_v39 _dafny.Sequence
			_ = _1625_v39
			_1625_v39 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)).Cardinality())
			var _1626_v40 _dafny.Set
			_ = _1626_v40
			_1626_v40 = _dafny.SetOf(_1625_v39)
			var _1627_v41 _dafny.Array
			_ = _1627_v41
			var _nwElement0_52 _dafny.Int = p1
			_ = _nwElement0_52
			var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(5))
			_ = _nw268
			(_nw268).ArraySet1(_nwElement0_52, 0)
			(_nw268).ArraySet1((func() _dafny.Int {
				if p0 {
					return (_dafny.MultiSetOf(_1624_v38)).Cardinality()
				}
				return p1
			})(), 1)
			(_nw268).ArraySet1((_1620_i1).Times(_1620_i1), 2)
			(_nw268).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1624_v38, _dafny.Companion_Sequence_.Update(_1624_v38, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1624_v38).Cardinality()))).Uint32(), p0))).Cardinality()), 3)
			(_nw268).ArraySet1(Companion_Default___.SafeModuloInt((_1626_v40).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-942)))), 4)
			_1627_v41 = _nw268
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))
			_ = _index285
			(_1627_v41).ArraySet1(p1, (_index285).Int())
			var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))
			_ = _index286
			var _rhs307 _dafny.Array = _1621_v37
			_ = _rhs307
			var _rhs308 _dafny.Int = p1
			_ = _rhs308
			var _lhs211 _dafny.Array = _1627_v41
			_ = _lhs211
			var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))
			_ = _lhs212
			_1621_v37 = _rhs307
			(_lhs211).ArraySet1(_rhs308, (_lhs212).Int())
			var _1628_v42 _dafny.Map
			_ = _1628_v42
			_1628_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1620_i1, _dafny.IntOfUint32((_1625_v39).Cardinality()))
			r0 = !(!((p1).Cmp(Companion_Default___.Fm3(true, p0, globalState)) > 0) || (((func() _dafny.Int {
				if (_1628_v42).Contains(_1620_i1) {
					return (_1628_v42).Get(_1620_i1).(_dafny.Int)
				}
				return p1
			})()).Cmp((_1627_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))).Int()).(_dafny.Int)) < 0))
			var _1629_v43 _dafny.CodePoint
			_ = _1629_v43
			_1629_v43 = _dafny.CodePoint('f')
			var _1630_v44 _dafny.Map
			_ = _1630_v44
			_1630_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), !(p0))
			var _1631_v45 _dafny.Map
			_ = _1631_v45
			_1631_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1620_i1, (func() bool {
				if (_1630_v44).Contains((_1627_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))).Int()).(_dafny.Int)) {
					return (_1630_v44).Get((_1627_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(214), _dafny.ArrayLen((_1627_v41), 0))).Int()).(_dafny.Int)).(bool)
				}
				return p0
			})())
			var _1632_v46 D0
			_ = _1632_v46
			var _1633_v47 bool
			_ = _1633_v47
			var _out44 D0
			_ = _out44
			var _out45 bool
			_ = _out45
			_out44, _out45 = Companion_Default___.M0(_1621_v37, _1629_v43, p0, Companion_Default___.Fm36(_1620_i1, p1, (_1631_v45).Cardinality(), _1628_v42, globalState), globalState)
			_1632_v46 = _out44
			_1633_v47 = _out45
		}
		var _1634_v48 D2
		_ = _1634_v48
		_1634_v48 = Companion_D2_.Create_DC6_(p2)
		var _1635_v49 _dafny.Map
		_ = _1635_v49
		_1635_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _1636_v50 _dafny.Array
		_ = _1636_v50
		var _nwElement0_53 _dafny.Int = p1
		_ = _nwElement0_53
		var _nw269 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(22))
		_ = _nw269
		(_nw269).ArraySet1(_nwElement0_53, 0)
		(_nw269).ArraySet1(p1, 1)
		(_nw269).ArraySet1(p1, 2)
		(_nw269).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 3)
		(_nw269).ArraySet1(p1, 4)
		(_nw269).ArraySet1((_dafny.Zero).Minus(p1), 5)
		(_nw269).ArraySet1(p1, 6)
		(_nw269).ArraySet1(p1, 7)
		(_nw269).ArraySet1(p1, 8)
		(_nw269).ArraySet1(p1, 9)
		(_nw269).ArraySet1(p1, 10)
		(_nw269).ArraySet1(p1, 11)
		(_nw269).ArraySet1(p1, 12)
		(_nw269).ArraySet1((_1635_v49).Cardinality(), 13)
		(_nw269).ArraySet1(p1, 14)
		(_nw269).ArraySet1(_dafny.IntOfInt64(-618), 15)
		(_nw269).ArraySet1(p1, 16)
		(_nw269).ArraySet1(p1, 17)
		(_nw269).ArraySet1(p1, 18)
		(_nw269).ArraySet1(p1, 19)
		(_nw269).ArraySet1(p1, 20)
		(_nw269).ArraySet1(p1, 21)
		_1636_v50 = _nw269
		var _1637_v51 _dafny.MultiSet
		_ = _1637_v51
		_1637_v51 = _dafny.MultiSetOf(_1636_v50, _1636_v50)
		var _1638_v52 _dafny.Sequence
		_ = _1638_v52
		_1638_v52 = _dafny.SeqOf((func() _dafny.Int {
			if (_1637_v51).Contains(_1636_v50) {
				return (_1637_v51).Multiplicity(_1636_v50)
			}
			return p1
		})())
		var _hi10 _dafny.Int = (_1638_v52).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1638_v52).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi10
		for _1639_i3 := Companion_Default___.SafeDivisionInt((_this).Fm29(p1, (_1634_v48).Dtor_cf6(), globalState), p1); _1639_i3.Cmp(_hi10) < 0; _1639_i3 = _1639_i3.Plus(_dafny.One) {
			var _1640_v53 _dafny.CodePoint
			_ = _1640_v53
			_1640_v53 = _dafny.CodePoint('k')
			var _1641_v54 _dafny.Set
			_ = _1641_v54
			_1641_v54 = _dafny.SetOf(_1640_v53, _1640_v53)
			var _1642_v55 _dafny.Array
			_ = _1642_v55
			var _nw270 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(5))
			_ = _nw270
			_1642_v55 = _nw270
			var _1643_v56 *C3
			_ = _1643_v56
			var _nw271 *C3 = New_C3_()
			_ = _nw271
			_nw271.Ctor__(p0, p0)
			_1643_v56 = _nw271
			var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_1642_v55), 0))
			_ = _index287
			(_1642_v55).ArraySet1((func() *C3 {
				if p0 {
					return _1643_v56
				}
				return _1643_v56
			})(), (_index287).Int())
			var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1636_v50), 0))
			_ = _index288
			(_1636_v50).ArraySet1(p1, (_index288).Int())
			var _1644_v57 _dafny.Array
			_ = _1644_v57
			var _nwElement0_54 _dafny.CodePoint = _1640_v53
			_ = _nwElement0_54
			var _nw272 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(13))
			_ = _nw272
			(_nw272).ArraySet1CodePoint(_nwElement0_54, 0)
			(_nw272).ArraySet1CodePoint(_1640_v53, 1)
			(_nw272).ArraySet1CodePoint(_1640_v53, 2)
			(_nw272).ArraySet1CodePoint(_1640_v53, 3)
			(_nw272).ArraySet1CodePoint(_1640_v53, 4)
			(_nw272).ArraySet1CodePoint(_1640_v53, 5)
			(_nw272).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _1640_v53
				}
				return _1640_v53
			})(), 6)
			(_nw272).ArraySet1CodePoint(_dafny.CodePoint('w'), 7)
			(_nw272).ArraySet1CodePoint(_1640_v53, 8)
			(_nw272).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1643_v56).F13() {
					return _1640_v53
				}
				return _1640_v53
			})(), 9)
			(_nw272).ArraySet1CodePoint(_1640_v53, 10)
			(_nw272).ArraySet1CodePoint(_1640_v53, 11)
			(_nw272).ArraySet1CodePoint(_1640_v53, 12)
			_1644_v57 = _nw272
			var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1644_v57), 0))
			_ = _index289
			(_1644_v57).ArraySet1CodePoint(_1640_v53, (_index289).Int())
			var _1645_v58 _dafny.Array
			_ = _1645_v58
			var _len0_47 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_47
			var _nw273 _dafny.Array
			_ = _nw273
			if _len0_47.Cmp(_dafny.Zero) == 0 {
				_nw273 = _dafny.NewArray(_len0_47)
			} else {
				var _init47 func(_dafny.Int) bool = func(_1646_i4 _dafny.Int) bool {
					return false
				}
				_ = _init47
				var _element0_47 = _init47(_dafny.Zero)
				_ = _element0_47
				_nw273 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
				(_nw273).ArraySet1(_element0_47, 0)
				var _nativeLen0_47 = (_len0_47).Int()
				_ = _nativeLen0_47
				for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
					(_nw273).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
				}
			}
			_1645_v58 = _nw273
			var _1647_v59 _dafny.Map
			_ = _1647_v59
			_1647_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1645_v58)
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_1642_v55), 0))
			_ = _index290
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1636_v50), 0))
			_ = _index291
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1644_v57), 0))
			_ = _index292
			var _rhs309 _dafny.Set = _1641_v54
			_ = _rhs309
			var _rhs310 *C3 = _1643_v56
			_ = _rhs310
			var _rhs311 _dafny.Int = _1639_i3
			_ = _rhs311
			var _rhs312 _dafny.CodePoint = _1640_v53
			_ = _rhs312
			var _rhs313 _dafny.Map = _1647_v59
			_ = _rhs313
			var _lhs213 _dafny.Array = _1642_v55
			_ = _lhs213
			var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_1642_v55), 0))
			_ = _lhs214
			var _lhs215 _dafny.Array = _1636_v50
			_ = _lhs215
			var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1636_v50), 0))
			_ = _lhs216
			var _lhs217 _dafny.Array = _1644_v57
			_ = _lhs217
			var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(977), _dafny.ArrayLen((_1644_v57), 0))
			_ = _lhs218
			_1641_v54 = _rhs309
			(_lhs213).ArraySet1(_rhs310, (_lhs214).Int())
			(_lhs215).ArraySet1(_rhs311, (_lhs216).Int())
			(_lhs217).ArraySet1CodePoint(_rhs312, (_lhs218).Int())
			_1647_v59 = _rhs313
			var _1648_v60 *C2
			_ = _1648_v60
			var _nw274 *C2 = New_C2_()
			_ = _nw274
			_nw274.Ctor__(_1640_v53, false)
			_1648_v60 = _nw274
			var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_1636_v50), 0))
			_ = _index293
			(_1636_v50).ArraySet1(Companion_Default___.SafeDivisionInt(_1639_i3, Companion_Default___.SafeModuloInt(p1, _dafny.IntOfUint32((p2).Cardinality()))), (_index293).Int())
			r0 = false
		}
		var _1649_v61 _dafny.Array
		_ = _1649_v61
		var _nw275 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(14))
		_ = _nw275
		_1649_v61 = _nw275
		for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1649_v61), 0))); ; {
			_guard_loop_2, _ok56 := _iter56()
			if !_ok56 {
				break
			}
			var _1650_i5 _dafny.Int
			_1650_i5 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_1650_i5).Sign() != -1) && ((_1650_i5).Cmp(_dafny.ArrayLen((_1649_v61), 0)) < 0)) {
				(_1649_v61).ArraySet1((func() _dafny.MultiSet {
					if p0 {
						return (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality())))).Union(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(20))).Uint32(), func(coer95 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg95 _dafny.Int) interface{} {
								return coer95(arg95)
							}
						}((func(_1651_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1652_i6 _dafny.Int) _dafny.Int {
								return _1651_p1
							}
						})(p1)))))
					}
					return (_dafny.MultiSetOf(p1)).Intersection(_dafny.MultiSetOf(p1, _dafny.IntOfUint32((_1638_v52).Cardinality())))
				})(), (_1650_i5).Int())
			}
		}
		r0 = p0
		var _1653_v62 _dafny.Map
		_ = _1653_v62
		_1653_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(691), p0)
		var _1654_v63 _dafny.CodePoint
		_ = _1654_v63
		_1654_v63 = _dafny.CodePoint('x')
		r1 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p2, p2), _dafny.Companion_Sequence_.Concatenate(p2, p2)), (Companion_Default___.SafeIndex(((_1653_v62).Update(p1, p0)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p2, p2), _dafny.Companion_Sequence_.Concatenate(p2, p2))).Cardinality()))).Uint32(), _1654_v63)
		return r0, r1
	}
}
func (_this *C6) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1655_v0 _dafny.Sequence
		_ = _1655_v0
		_1655_v0 = _dafny.UnicodeSeqOfUtf8Bytes("ojxcdwwr")
		_1655_v0 = _1655_v0
		var _1656_v1 bool
		_ = _1656_v1
		_1656_v1 = true
		var _1657_i0 _dafny.Int
		_ = _1657_i0
		_1657_i0 = _dafny.Zero
		{
			for _1656_v1 {
				{
					if (_1657_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L16
					}
					_1657_i0 = (_1657_i0).Plus(_dafny.One)
					var _1658_v2 _dafny.Sequence
					_ = _1658_v2
					_1658_v2 = _dafny.SeqOf(_1656_v1)
					var _1659_v3 _dafny.Int
					_ = _1659_v3
					_1659_v3 = _dafny.IntOfInt64(715)
					var _1660_v4 D11
					_ = _1660_v4
					_1660_v4 = Companion_D11_.Create_DC20_(_1656_v1, (_dafny.Zero).Minus(_1659_v3))
					if (_1658_v2).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_1660_v4).Dtor_cf21() {
							return _1659_v3
						}
						return _1659_v3
					})(), _dafny.IntOfUint32((_1658_v2).Cardinality()))).Uint32()).(bool) {
						var _1661_v5 *C4
						_ = _1661_v5
						var _nw276 *C4 = New_C4_()
						_ = _nw276
						_nw276.Ctor__()
						_1661_v5 = _nw276
						var _1662_v6 _dafny.CodePoint
						_ = _1662_v6
						_1662_v6 = _dafny.CodePoint('o')
						_1662_v6 = _1662_v6
						var _1663_v7 _dafny.Array
						_ = _1663_v7
						var _nw277 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
						_ = _nw277
						_1663_v7 = _nw277
						var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1663_v7), 0))
						_ = _index294
						(_1663_v7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1655_v0, _1655_v0), (_index294).Int())
						var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1663_v7), 0))
						_ = _index295
						(_1663_v7).ArraySet1((func() _dafny.Sequence {
							if !(!(_1656_v1)) {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-456))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg96 _dafny.Int) interface{} {
										return coer96(arg96)
									}
								}((func(_1664_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_1665_i1 _dafny.Int) _dafny.CodePoint {
										return _1664_v6
									}
								})(_1662_v6))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-219))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg97 _dafny.Int) interface{} {
										return coer97(arg97)
									}
								}((func(_1666_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
									return func(_1667_i2 _dafny.Int) _dafny.CodePoint {
										return _1666_v6
									}
								})(_1662_v6))))
							}
							return _1655_v0
						})(), (_index295).Int())
						var _1668_v8 _dafny.Array
						_ = _1668_v8
						var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
						_ = _nw278
						_1668_v8 = _nw278
						var _1669_v9 _dafny.Map
						_ = _1669_v9
						_1669_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, (_dafny.Zero).Minus(_1659_v3))
						var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1668_v8), 0))
						_ = _index296
						(_1668_v8).ArraySet1(_1669_v9, (_index296).Int())
						var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1668_v8), 0))
						_ = _index297
						var _rhs314 _dafny.Int = _1659_v3
						_ = _rhs314
						var _rhs315 bool = !(((_dafny.Zero).Minus((_1659_v3).Times(_1659_v3))).Cmp(_1659_v3) != 0)
						_ = _rhs315
						var _rhs316 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qn"), _dafny.UnicodeSeqOfUtf8Bytes("syeqbdag")), _dafny.Companion_Sequence_.Concatenate(_1655_v0, (_1663_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1663_v7), 0))).Int()).(_dafny.Sequence)))
						_ = _rhs316
						var _rhs317 bool = _1656_v1
						_ = _rhs317
						var _rhs318 _dafny.Map = (_1669_v9).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1659_v3)).Merge(_1669_v9))
						_ = _rhs318
						var _lhs219 _dafny.Array = _1668_v8
						_ = _lhs219
						var _lhs220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(108), _dafny.ArrayLen((_1668_v8), 0))
						_ = _lhs220
						_1659_v3 = _rhs314
						_1656_v1 = _rhs315
						_1655_v0 = _rhs316
						_1656_v1 = _rhs317
						(_lhs219).ArraySet1(_rhs318, (_lhs220).Int())
						_1661_v5 = _1661_v5
					} else {
						var _1670_v10 _dafny.Array
						_ = _1670_v10
						var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
						_ = _nw279
						_1670_v10 = _nw279
						var _1671_v11 _dafny.Array
						_ = _1671_v11
						var _nwElement0_55 _dafny.Array = _1670_v10
						_ = _nwElement0_55
						var _nw280 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(13))
						_ = _nw280
						(_nw280).ArraySet1(_nwElement0_55, 0)
						(_nw280).ArraySet1(_1670_v10, 1)
						(_nw280).ArraySet1(_1670_v10, 2)
						(_nw280).ArraySet1(_1670_v10, 3)
						(_nw280).ArraySet1(_1670_v10, 4)
						(_nw280).ArraySet1(_1670_v10, 5)
						(_nw280).ArraySet1(_1670_v10, 6)
						(_nw280).ArraySet1(_1670_v10, 7)
						(_nw280).ArraySet1(_1670_v10, 8)
						(_nw280).ArraySet1(_1670_v10, 9)
						(_nw280).ArraySet1(_1670_v10, 10)
						(_nw280).ArraySet1(_1670_v10, 11)
						(_nw280).ArraySet1(_1670_v10, 12)
						_1671_v11 = _nw280
						var _1672_v12 _dafny.Array
						_ = _1672_v12
						var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
						_ = _nw281
						_1672_v12 = _nw281
						var _1673_v13 _dafny.Map
						_ = _1673_v13
						_1673_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1671_v11, _1672_v12)
						_1673_v13 = (_1673_v13).Update(_1671_v11, _1672_v12)
						var _1674_v14 _dafny.CodePoint
						_ = _1674_v14
						_1674_v14 = _dafny.CodePoint('r')
						var _1675_v15 D11
						_ = _1675_v15
						_1675_v15 = Companion_D11_.Create_DC19_(_1674_v14)
						_1675_v15 = _1675_v15
						var _1676_v16 *C4
						_ = _1676_v16
						var _nw282 *C4 = New_C4_()
						_ = _nw282
						_nw282.Ctor__()
						_1676_v16 = _nw282
						var _1677_v17 _dafny.Set
						_ = _1677_v17
						_1677_v17 = _dafny.SetOf(_1656_v1, _1656_v1)
						r0 = Companion_Default___.SafeDivisionInt(_1659_v3, Companion_Default___.SafeDivisionInt((_1677_v17).Cardinality(), _1659_v3))
						var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1672_v12), 0))
						_ = _index298
						(_1672_v12).ArraySet1(_1659_v3, (_index298).Int())
						var _1678_v18 _dafny.Map
						_ = _1678_v18
						_1678_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1659_v3, _1656_v1)
						var _1679_v19 _dafny.Map
						_ = _1679_v19
						_1679_v19 = _1678_v18
						var _1680_v20 _dafny.Array
						_ = _1680_v20
						var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(11))
						_ = _nw283
						_1680_v20 = _nw283
						var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1680_v20), 0))
						_ = _index299
						(_1680_v20).ArraySet1(_1658_v2, (_index299).Int())
						var _1681_v21 _dafny.Sequence
						_ = _1681_v21
						_1681_v21 = _dafny.SeqOf(_dafny.IntOfInt64(298), _1659_v3, _1659_v3)
						var _1682_v23 _dafny.Sequence
						_ = _1682_v23
						_1682_v23 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("yhq"), _1655_v0)
						var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1672_v12), 0))
						_ = _index300
						var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1680_v20), 0))
						_ = _index301
						var _rhs319 _dafny.Int = (func() _dafny.Int {
							if !((func() bool {
								if _1656_v1 {
									return _1656_v1
								}
								return false
							})()) {
								return _1659_v3
							}
							return _dafny.IntOfUint32((_1681_v21).Cardinality())
						})()
						_ = _rhs319
						var _rhs320 _dafny.Map = _1679_v19
						_ = _rhs320
						var _rhs321 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1658_v2, (Companion_Default___.SafeIndex((_1678_v18).Cardinality(), _dafny.IntOfUint32((_1658_v2).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex((func() _dafny.Map {
							var _coll54 = _dafny.NewMapBuilder()
							_ = _coll54
							for _iter57 := _dafny.Iterate((func() _dafny.Set {
								var _coll55 = _dafny.NewBuilder()
								_ = _coll55
								for _iter58 := _dafny.Iterate((_1682_v23).Elements()); ; {
									_compr_55, _ok58 := _iter58()
									if !_ok58 {
										break
									}
									var _1683_v24 _dafny.Sequence
									_1683_v24 = interface{}(_compr_55).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_1682_v23, _1683_v24) {
										_coll55.Add(_1683_v24)
									}
								}
								return _coll55.ToSet()
							}()).Elements()); ; {
								_compr_54, _ok57 := _iter57()
								if !_ok57 {
									break
								}
								var _1684_v22 _dafny.Sequence
								_1684_v22 = interface{}(_compr_54).(_dafny.Sequence)
								if (func() _dafny.Set {
									var _coll56 = _dafny.NewBuilder()
									_ = _coll56
									for _iter59 := _dafny.Iterate((_1682_v23).Elements()); ; {
										_compr_56, _ok59 := _iter59()
										if !_ok59 {
											break
										}
										var _1685_v24 _dafny.Sequence
										_1685_v24 = interface{}(_compr_56).(_dafny.Sequence)
										if _dafny.Companion_Sequence_.Contains(_1682_v23, _1685_v24) {
											_coll56.Add(_1685_v24)
										}
									}
									return _coll56.ToSet()
								}()).Contains(_1684_v22) {
									_coll54.Add(_1684_v22, _1656_v1)
								}
							}
							return _coll54.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1658_v2, (Companion_Default___.SafeIndex((_1678_v18).Cardinality(), _dafny.IntOfUint32((_1658_v2).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), _1656_v1), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1658_v2).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1658_v2, (Companion_Default___.SafeIndex((_1678_v18).Cardinality(), _dafny.IntOfUint32((_1658_v2).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex((func() _dafny.Map {
							var _coll57 = _dafny.NewMapBuilder()
							_ = _coll57
							for _iter60 := _dafny.Iterate((func() _dafny.Set {
								var _coll58 = _dafny.NewBuilder()
								_ = _coll58
								for _iter61 := _dafny.Iterate((_1682_v23).Elements()); ; {
									_compr_58, _ok61 := _iter61()
									if !_ok61 {
										break
									}
									var _1686_v24 _dafny.Sequence
									_1686_v24 = interface{}(_compr_58).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_1682_v23, _1686_v24) {
										_coll58.Add(_1686_v24)
									}
								}
								return _coll58.ToSet()
							}()).Elements()); ; {
								_compr_57, _ok60 := _iter60()
								if !_ok60 {
									break
								}
								var _1687_v22 _dafny.Sequence
								_1687_v22 = interface{}(_compr_57).(_dafny.Sequence)
								if (func() _dafny.Set {
									var _coll59 = _dafny.NewBuilder()
									_ = _coll59
									for _iter62 := _dafny.Iterate((_1682_v23).Elements()); ; {
										_compr_59, _ok62 := _iter62()
										if !_ok62 {
											break
										}
										var _1688_v24 _dafny.Sequence
										_1688_v24 = interface{}(_compr_59).(_dafny.Sequence)
										if _dafny.Companion_Sequence_.Contains(_1682_v23, _1688_v24) {
											_coll59.Add(_1688_v24)
										}
									}
									return _coll59.ToSet()
								}()).Contains(_1687_v22) {
									_coll57.Add(_1687_v22, _1656_v1)
								}
							}
							return _coll57.ToMap()
						}()).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1658_v2, (Companion_Default___.SafeIndex((_1678_v18).Cardinality(), _dafny.IntOfUint32((_1658_v2).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), _1656_v1)).Cardinality()))).Uint32(), _1656_v1), _1658_v2)
						_ = _rhs321
						var _lhs221 _dafny.Array = _1672_v12
						_ = _lhs221
						var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((_1672_v12), 0))
						_ = _lhs222
						var _lhs223 _dafny.Array = _1680_v20
						_ = _lhs223
						var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_1680_v20), 0))
						_ = _lhs224
						(_lhs221).ArraySet1(_rhs319, (_lhs222).Int())
						_1679_v19 = _rhs320
						(_lhs223).ArraySet1(_rhs321, (_lhs224).Int())
					}
					var _1689_v25 *C4
					_ = _1689_v25
					var _nw284 *C4 = New_C4_()
					_ = _nw284
					_nw284.Ctor__()
					_1689_v25 = _nw284
					var _1690_v26 D14
					_ = _1690_v26
					_1690_v26 = Companion_D14_.Create_DC25_(_1689_v25)
					var _pat_let_tv20 = _1689_v25
					_ = _pat_let_tv20
					var _1691_v27 _dafny.Array
					_ = _1691_v27
					var _nwElement0_56 D14 = _1690_v26
					_ = _nwElement0_56
					var _nw285 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(13))
					_ = _nw285
					(_nw285).ArraySet1(_nwElement0_56, 0)
					(_nw285).ArraySet1(_1690_v26, 1)
					(_nw285).ArraySet1(_1690_v26, 2)
					(_nw285).ArraySet1(_1690_v26, 3)
					(_nw285).ArraySet1(_1690_v26, 4)
					(_nw285).ArraySet1(_1690_v26, 5)
					(_nw285).ArraySet1(Companion_D14_.Create_DC25_(_1689_v25), 6)
					(_nw285).ArraySet1(_1690_v26, 7)
					(_nw285).ArraySet1(Companion_D14_.Create_DC25_(_1689_v25), 8)
					(_nw285).ArraySet1(Companion_D14_.Create_DC25_(_1689_v25), 9)
					(_nw285).ArraySet1(_1690_v26, 10)
					(_nw285).ArraySet1(_1690_v26, 11)
					(_nw285).ArraySet1(func(_pat_let22_0 D14) D14 {
						return func(_1692_dt__update__tmp_h0 D14) D14 {
							return func(_pat_let23_0 *C4) D14 {
								return func(_1693_dt__update_hcf31_h0 *C4) D14 {
									return Companion_D14_.Create_DC25_(_1693_dt__update_hcf31_h0)
								}(_pat_let23_0)
							}(_pat_let_tv20)
						}(_pat_let22_0)
					}(_1690_v26), 12)
					_1691_v27 = _nw285
					var _pat_let_tv21 = _1689_v25
					_ = _pat_let_tv21
					var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1691_v27), 0))
					_ = _index302
					(_1691_v27).ArraySet1(func(_pat_let24_0 D14) D14 {
						return func(_1694_dt__update__tmp_h1 D14) D14 {
							return func(_pat_let25_0 *C4) D14 {
								return func(_1695_dt__update_hcf31_h1 *C4) D14 {
									return Companion_D14_.Create_DC25_(_1695_dt__update_hcf31_h1)
								}(_pat_let25_0)
							}(_pat_let_tv21)
						}(_pat_let24_0)
					}(_1690_v26), (_index302).Int())
					var _1696_v28 _dafny.Array
					_ = _1696_v28
					var _nw286 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
					_ = _nw286
					_1696_v28 = _nw286
					var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1696_v28), 0))
					_ = _index303
					(_1696_v28).ArraySet1(_dafny.SeqOf(_1659_v3), (_index303).Int())
					var _1697_v29 _dafny.Array
					_ = _1697_v29
					var _len0_48 _dafny.Int = _dafny.IntOfInt64(10)
					_ = _len0_48
					var _nw287 _dafny.Array
					_ = _nw287
					if _len0_48.Cmp(_dafny.Zero) == 0 {
						_nw287 = _dafny.NewArray(_len0_48)
					} else {
						var _init48 func(_dafny.Int) _dafny.Map = (func(_1698_v2 _dafny.Sequence, _1699_v3 _dafny.Int) func(_dafny.Int) _dafny.Map {
							return func(_1700_i3 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1698_v2)).Cardinality(), _1699_v3)
							}
						})(_1658_v2, _1659_v3)
						_ = _init48
						var _element0_48 = _init48(_dafny.Zero)
						_ = _element0_48
						_nw287 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
						(_nw287).ArraySet1(_element0_48, 0)
						var _nativeLen0_48 = (_len0_48).Int()
						_ = _nativeLen0_48
						for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
							(_nw287).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
						}
					}
					_1697_v29 = _nw287
					var _1701_v30 _dafny.CodePoint
					_ = _1701_v30
					_1701_v30 = _dafny.CodePoint('b')
					var _1702_v31 D11
					_ = _1702_v31
					_1702_v31 = Companion_D11_.Create_DC19_(_1701_v30)
					var _1703_v32 _dafny.Map
					_ = _1703_v32
					_1703_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1656_v1), (_1702_v31).Dtor_cf20())).Update(_1656_v1, _1701_v30), _dafny.IntOfInt64(88))
					var _1704_v33 _dafny.Array
					_ = _1704_v33
					var _nw288 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
					_ = _nw288
					_1704_v33 = _nw288
					var _1705_v34 _dafny.Array
					_ = _1705_v34
					_1705_v34 = _1704_v33
					var _1706_v35 _dafny.Map
					_ = _1706_v35
					_1706_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1705_v34, _1656_v1)
					var _1707_v36 _dafny.Sequence
					_ = _1707_v36
					_1707_v36 = _dafny.SeqOf((_1703_v32).Cardinality(), ((func() _dafny.Map {
						if Companion_Default___.Fm0(Companion_Default___.Fm0(!(_1656_v1), globalState), globalState) {
							return _1706_v35
						}
						return _1706_v35
					})()).Cardinality(), _1659_v3, _1659_v3)
					var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1691_v27), 0))
					_ = _index304
					var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1696_v28), 0))
					_ = _index305
					var _rhs322 D14 = _1690_v26
					_ = _rhs322
					var _rhs323 _dafny.Sequence = _1707_v36
					_ = _rhs323
					var _rhs324 bool = !(_1656_v1) || (_1656_v1)
					_ = _rhs324
					var _rhs325 _dafny.Int = _dafny.IntOfInt64(342)
					_ = _rhs325
					var _rhs326 _dafny.Array = _1697_v29
					_ = _rhs326
					var _lhs225 _dafny.Array = _1691_v27
					_ = _lhs225
					var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(105), _dafny.ArrayLen((_1691_v27), 0))
					_ = _lhs226
					var _lhs227 _dafny.Array = _1696_v28
					_ = _lhs227
					var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1696_v28), 0))
					_ = _lhs228
					var _lhs229 *GlobalState = globalState
					_ = _lhs229
					(_lhs225).ArraySet1(_rhs322, (_lhs226).Int())
					(_lhs227).ArraySet1(_rhs323, (_lhs228).Int())
					_1656_v1 = _rhs324
					_lhs229.F1 = _rhs325
					_1697_v29 = _rhs326
					var _1708_v37 _dafny.Map
					_ = _1708_v37
					_1708_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1655_v0)
					var _1709_v38 _dafny.MultiSet
					_ = _1709_v38
					_1709_v38 = _dafny.MultiSetOf((func() _dafny.Sequence {
						if (_1708_v37).Contains(_1656_v1) {
							return (_1708_v37).Get(_1656_v1).(_dafny.Sequence)
						}
						return _1655_v0
					})())
					var _1710_v39 _dafny.Map
					_ = _1710_v39
					_1710_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1659_v3)
					var _1711_v40 _dafny.Map
					_ = _1711_v40
					_1711_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1710_v39)
					var _1712_v42 _dafny.MultiSet
					_ = _1712_v42
					_1712_v42 = _dafny.MultiSetOf(_1656_v1, _1656_v1, _1656_v1, _1656_v1)
					var _1713_v43 _dafny.MultiSet
					_ = _1713_v43
					_1713_v43 = _dafny.MultiSetOf(_1659_v3)
					var _1714_v44 _dafny.Array
					_ = _1714_v44
					var _nwElement0_57 _dafny.Int = (func() _dafny.Int {
						if _1656_v1 {
							return _dafny.IntOfInt64(-450)
						}
						return _1659_v3
					})()
					_ = _nwElement0_57
					var _nw289 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(17))
					_ = _nw289
					(_nw289).ArraySet1(_nwElement0_57, 0)
					(_nw289).ArraySet1((_dafny.Zero).Minus((_1711_v40).Cardinality()), 1)
					(_nw289).ArraySet1((_1659_v3).Plus(_dafny.IntOfUint32((_1655_v0).Cardinality())), 2)
					(_nw289).ArraySet1((Companion_Default___.Fm35(_1655_v0, _dafny.MultiSetFromSeq((_1696_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1696_v28), 0))).Int()).(_dafny.Sequence)), _1701_v30, globalState)).Cardinality(), 3)
					(_nw289).ArraySet1(_dafny.IntOfUint32((_1655_v0).Cardinality()), 4)
					(_nw289).ArraySet1((func() _dafny.Set {
						var _coll60 = _dafny.NewBuilder()
						_ = _coll60
						for _iter63 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(973), _dafny.IntOfInt64(8))); ; {
							_compr_60, _ok63 := _iter63()
							if !_ok63 {
								break
							}
							var _1715_v41 _dafny.Int
							_1715_v41 = interface{}(_compr_60).(_dafny.Int)
							if ((_dafny.IntOfInt64(973)).Cmp(_1715_v41) <= 0) && ((_1715_v41).Cmp(_dafny.IntOfInt64(8)) < 0) {
								_coll60.Add((_1715_v41).Minus(_1659_v3))
							}
						}
						return _coll60.ToSet()
					}()).Cardinality(), 5)
					(_nw289).ArraySet1((func() _dafny.Int {
						if _1656_v1 {
							return (_dafny.Zero).Minus(_dafny.IntOfUint32((_1655_v0).Cardinality()))
						}
						return _1659_v3
					})(), 6)
					(_nw289).ArraySet1(_1659_v3, 7)
					(_nw289).ArraySet1(_1659_v3, 8)
					(_nw289).ArraySet1(_dafny.IntOfInt64(927), 9)
					(_nw289).ArraySet1(_1659_v3, 10)
					(_nw289).ArraySet1(Companion_Default___.Fm3(_1656_v1, _1656_v1, globalState), 11)
					(_nw289).ArraySet1(_1659_v3, 12)
					(_nw289).ArraySet1(_1659_v3, 13)
					(_nw289).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1655_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(131))).Uint32(), func(coer98 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg98 _dafny.Int) interface{} {
							return coer98(arg98)
						}
					}((func(_1716_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1717_i4 _dafny.Int) _dafny.CodePoint {
							return _1716_v30
						}
					})(_1701_v30))))).Cardinality()), 14)
					(_nw289).ArraySet1((_1712_v42).Cardinality(), 15)
					(_nw289).ArraySet1((_1713_v43).Cardinality(), 16)
					_1714_v44 = _nw289
					var _1718_v45 _dafny.Array
					_ = _1718_v45
					_1718_v45 = _1714_v44
					var _1719_v46 D12
					_ = _1719_v46
					_1719_v46 = Companion_D12_.Create_DC23_(_1714_v44)
					var _1720_v47 _dafny.Sequence
					_ = _1720_v47
					_1720_v47 = _dafny.SeqOf(_1719_v46, _1719_v46)
					var _1721_v48 _dafny.Map
					_ = _1721_v48
					_1721_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(264)).Times(_dafny.IntOfUint32((_1655_v0).Cardinality())), _dafny.Companion_Sequence_.Update(_1720_v47, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1707_v36).Cardinality()), _dafny.IntOfUint32((_1720_v47).Cardinality()))).Uint32(), _1719_v46))
					var _1722_v49 _dafny.Sequence
					_ = _1722_v49
					_1722_v49 = _dafny.SeqOf(_1655_v0, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1655_v0, (Companion_Default___.SafeIndex(_1659_v3, _dafny.IntOfUint32((_1655_v0).Cardinality()))).Uint32(), _1701_v30), (Companion_Default___.SafeIndex(_1659_v3, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1655_v0, (Companion_Default___.SafeIndex(_1659_v3, _dafny.IntOfUint32((_1655_v0).Cardinality()))).Uint32(), _1701_v30)).Cardinality()))).Uint32(), _1701_v30))
					var _1723_v50 _dafny.Map
					_ = _1723_v50
					_1723_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(292), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1659_v3, _1720_v47))
					var _rhs327 _dafny.Int = (_1659_v3).Times(_1659_v3)
					_ = _rhs327
					var _rhs328 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1722_v49, _1722_v49))
					_ = _rhs328
					var _rhs329 _dafny.Int = ((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-808), _1659_v3))).Minus((_1659_v3).Minus(_dafny.IntOfInt64(393)))
					_ = _rhs329
					var _rhs330 _dafny.Array = (_1718_v45)
					_ = _rhs330
					var _rhs331 _dafny.Map = (_1721_v48).Merge((func() _dafny.Map {
						if (_1723_v50).Contains(_1659_v3) {
							return (_1723_v50).Get(_1659_v3).(_dafny.Map)
						}
						return _1721_v48
					})())
					_ = _rhs331
					var _lhs230 *GlobalState = globalState
					_ = _lhs230
					_1659_v3 = _rhs327
					_1709_v38 = _rhs328
					_lhs230.F1 = _rhs329
					_1714_v44 = _rhs330
					_1721_v48 = _rhs331
					(globalState).F1 = _1659_v3
					goto C16
				}
			C16:
			}
			goto L16
		}
	L16:
		if _1656_v1 {
			var _1724_v53 _dafny.Array
			_ = _1724_v53
			var _len0_49 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_49
			var _nw290 _dafny.Array
			_ = _nw290
			if _len0_49.Cmp(_dafny.Zero) == 0 {
				_nw290 = _dafny.NewArray(_len0_49)
			} else {
				var _init49 func(_dafny.Int) _dafny.CodePoint = (func(_1725_v0 _dafny.Sequence, _1726_v1 bool) func(_dafny.Int) _dafny.CodePoint {
					return func(_1727_i5 _dafny.Int) _dafny.CodePoint {
						return (_1725_v0).Select((Companion_Default___.SafeIndex((_dafny.SetOf((func() _dafny.Map {
							var _coll61 = _dafny.NewMapBuilder()
							_ = _coll61
							for _iter64 := _dafny.Iterate((func() _dafny.Set {
								var _coll62 = _dafny.NewBuilder()
								_ = _coll62
								for _iter65 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-862))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
									return func(arg99 _dafny.Int) interface{} {
										return coer99(arg99)
									}
								}(func(_1728_i6 _dafny.Int) _dafny.CodePoint {
									return _dafny.CodePoint('p')
								})), _1725_v0)).Elements()); ; {
									_compr_62, _ok65 := _iter65()
									if !_ok65 {
										break
									}
									var _1729_v52 _dafny.Sequence
									_1729_v52 = interface{}(_compr_62).(_dafny.Sequence)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-862))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg100 _dafny.Int) interface{} {
											return coer100(arg100)
										}
									}(func(_1728_i6 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('p')
									})), _1725_v0), _1729_v52) {
										_coll62.Add(_1729_v52)
									}
								}
								return _coll62.ToSet()
							}()).Elements()); ; {
								_compr_61, _ok64 := _iter64()
								if !_ok64 {
									break
								}
								var _1730_v51 _dafny.Sequence
								_1730_v51 = interface{}(_compr_61).(_dafny.Sequence)
								if (func() _dafny.Set {
									var _coll63 = _dafny.NewBuilder()
									_ = _coll63
									for _iter66 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-862))).Uint32(), func(coer101 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
										return func(arg101 _dafny.Int) interface{} {
											return coer101(arg101)
										}
									}(func(_1728_i6 _dafny.Int) _dafny.CodePoint {
										return _dafny.CodePoint('p')
									})), _1725_v0)).Elements()); ; {
										_compr_63, _ok66 := _iter66()
										if !_ok66 {
											break
										}
										var _1731_v52 _dafny.Sequence
										_1731_v52 = interface{}(_compr_63).(_dafny.Sequence)
										if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-862))).Uint32(), func(coer102 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
											return func(arg102 _dafny.Int) interface{} {
												return coer102(arg102)
											}
										}(func(_1728_i6 _dafny.Int) _dafny.CodePoint {
											return _dafny.CodePoint('p')
										})), _1725_v0), _1731_v52) {
											_coll63.Add(_1731_v52)
										}
									}
									return _coll63.ToSet()
								}()).Contains(_1730_v51) {
									_coll61.Add(_1730_v51, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_1726_v1)), _dafny.IntOfInt64(989)))
								}
							}
							return _coll61.ToMap()
						}()).Cardinality())).Cardinality(), _dafny.IntOfUint32((_1725_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_1655_v0, _1656_v1)
				_ = _init49
				var _element0_49 = _init49(_dafny.Zero)
				_ = _element0_49
				_nw290 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
				(_nw290).ArraySet1CodePoint(_element0_49, 0)
				var _nativeLen0_49 = (_len0_49).Int()
				_ = _nativeLen0_49
				for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
					(_nw290).ArraySet1CodePoint(_init49(_dafny.IntOf(_i0_49)), _i0_49)
				}
			}
			_1724_v53 = _nw290
			_1724_v53 = _1724_v53
			if !(_1656_v1) {
				var _1732_v54 _dafny.Int
				_ = _1732_v54
				_1732_v54 = _dafny.IntOfInt64(487)
				(globalState).F1 = (_1732_v54).Times(_1732_v54)
				var _1733_v55 _dafny.MultiSet
				_ = _1733_v55
				_1733_v55 = _dafny.MultiSetOf(!(_1656_v1))
				_1732_v54 = (func() _dafny.Int {
					if (_1733_v55).Contains(_1656_v1) {
						return (_1733_v55).Multiplicity(_1656_v1)
					}
					return _dafny.IntOfInt64(94)
				})()
				var _1734_v56 _dafny.Array
				_ = _1734_v56
				var _len0_50 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_50
				var _nw291 _dafny.Array
				_ = _nw291
				if _len0_50.Cmp(_dafny.Zero) == 0 {
					_nw291 = _dafny.NewArray(_len0_50)
				} else {
					var _init50 func(_dafny.Int) bool = (func(_1735_v1 bool) func(_dafny.Int) bool {
						return func(_1736_i7 _dafny.Int) bool {
							return _1735_v1
						}
					})(_1656_v1)
					_ = _init50
					var _element0_50 = _init50(_dafny.Zero)
					_ = _element0_50
					_nw291 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
					(_nw291).ArraySet1(_element0_50, 0)
					var _nativeLen0_50 = (_len0_50).Int()
					_ = _nativeLen0_50
					for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
						(_nw291).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
					}
				}
				_1734_v56 = _nw291
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_1734_v56), 0))
				_ = _index306
				(_1734_v56).ArraySet1(_1656_v1, (_index306).Int())
				var _1737_v57 _dafny.CodePoint
				_ = _1737_v57
				_1737_v57 = _dafny.CodePoint('r')
				var _1738_v58 _dafny.Map
				_ = _1738_v58
				_1738_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1734_v56, _dafny.CodePoint('c'))
				var _1739_v59 D0
				_ = _1739_v59
				_1739_v59 = Companion_D0_.Create_DC2_(_1738_v58, _1737_v57)
				var _1740_v60 _dafny.Sequence
				_ = _1740_v60
				_1740_v60 = _dafny.SeqOf(_1739_v59, Companion_D0_.Create_DC2_(_1738_v58, _1737_v57))
				var _1741_v61 _dafny.Sequence
				_ = _1741_v61
				_1741_v61 = _dafny.SeqOf(_1656_v1, _1656_v1)
				var _1742_v62 _dafny.Set
				_ = _1742_v62
				_1742_v62 = _dafny.SetOf(_1656_v1)
				var _1743_v63 _dafny.Sequence
				_ = _1743_v63
				_1743_v63 = _dafny.SeqOf(_1732_v54)
				var _1744_v64 _dafny.MultiSet
				_ = _1744_v64
				_1744_v64 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1743_v63).Cardinality()))
				var _1745_v65 _dafny.Map
				_ = _1745_v65
				_1745_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1732_v54, (_this).Fm4(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1741_v61, (Companion_Default___.SafeIndex((_1742_v62).Cardinality(), _dafny.IntOfUint32((_1741_v61).Cardinality()))).Uint32(), _1656_v1), (Companion_Default___.SafeIndex(_1732_v54, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1741_v61, (Companion_Default___.SafeIndex((_1742_v62).Cardinality(), _dafny.IntOfUint32((_1741_v61).Cardinality()))).Uint32(), _1656_v1)).Cardinality()))).Uint32(), !(_1656_v1)), _1732_v54, (_dafny.Zero).Minus(_1732_v54), _1744_v64, globalState))
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_1734_v56), 0))
				_ = _index307
				var _rhs332 bool = _1656_v1
				_ = _rhs332
				var _rhs333 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_1739_v59), _1740_v60)
				_ = _rhs333
				var _rhs334 _dafny.CodePoint = Companion_Default___.Fm36(_1732_v54, _1732_v54, (_1732_v54).Minus(_1732_v54), _1745_v65, globalState)
				_ = _rhs334
				var _lhs231 _dafny.Array = _1734_v56
				_ = _lhs231
				var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_1734_v56), 0))
				_ = _lhs232
				(_lhs231).ArraySet1(_rhs332, (_lhs232).Int())
				_1656_v1 = _rhs333
				_1737_v57 = _rhs334
				_1656_v1 = _1656_v1
				_1732_v54 = _1732_v54
			} else {
				var _1746_v66 _dafny.Sequence
				_ = _1746_v66
				_1746_v66 = _dafny.SeqOf(Companion_Default___.Fm3(_1656_v1, _1656_v1, globalState))
				var _1747_v67 _dafny.Int
				_ = _1747_v67
				_1747_v67 = _dafny.IntOfInt64(-506)
				var _1748_v68 _dafny.Array
				_ = _1748_v68
				var _nw292 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw292
				_1748_v68 = _nw292
				var _1749_v69 _dafny.Map
				_ = _1749_v69
				_1749_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1746_v66).Select((Companion_Default___.SafeIndex(_1747_v67, _dafny.IntOfUint32((_1746_v66).Cardinality()))).Uint32()).(_dafny.Int), _1748_v68)
				_1749_v69 = (_1749_v69).Update(_1747_v67, _1748_v68)
				var _1750_v70 _dafny.Array
				_ = _1750_v70
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_51
				var _nw293 _dafny.Array
				_ = _nw293
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw293 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.MultiSet = (func(_1751_v67 _dafny.Int) func(_dafny.Int) _dafny.MultiSet {
						return func(_1752_i8 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_1751_v67, _1751_v67)).Difference(_dafny.MultiSetOf(_1751_v67, _dafny.IntOfInt64(459)))
						}
					})(_1747_v67)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw293 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw293).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw293).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1750_v70 = _nw293
				var _1753_v71 _dafny.MultiSet
				_ = _1753_v71
				_1753_v71 = _dafny.MultiSetOf(_1747_v67)
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1750_v70), 0))
				_ = _index308
				(_1750_v70).ArraySet1((_1753_v71).Update(_1747_v67, Companion_Default___.Abs(_1747_v67)), (_index308).Int())
				var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(556), _dafny.ArrayLen((_1750_v70), 0))
				_ = _index309
				(_1750_v70).ArraySet1(_1753_v71, (_index309).Int())
				var _1754_v72 _dafny.CodePoint
				_ = _1754_v72
				_1754_v72 = _dafny.CodePoint('k')
				var _1755_v73 T3
				_ = _1755_v73
				var _nw294 *C5 = New_C5_()
				_ = _nw294
				_nw294.Ctor__((_dafny.Zero).Minus(_1747_v67), _1754_v72, _1656_v1)
				_1755_v73 = _nw294
				var _1756_v74 _dafny.Array
				_ = _1756_v74
				var _nwElement0_58 T3 = _1755_v73
				_ = _nwElement0_58
				var _nw295 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(2))
				_ = _nw295
				(_nw295).ArraySet1(_nwElement0_58, 0)
				(_nw295).ArraySet1(_1755_v73, 1)
				_1756_v74 = _nw295
				var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_1756_v74), 0))
				_ = _index310
				(_1756_v74).ArraySet1(_1755_v73, (_index310).Int())
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_1756_v74), 0))
				_ = _index311
				(_1756_v74).ArraySet1(_1755_v73, (_index311).Int())
				_1747_v67 = (_dafny.Zero).Minus((_dafny.IntOfInt64(-624)).Times(Companion_Default___.SafeDivisionInt(_1747_v67, _1747_v67)))
				_1656_v1 = (_1755_v73).F6()
			}
			if _1656_v1 {
				var _1757_v75 D2
				_ = _1757_v75
				_1757_v75 = Companion_D2_.Create_DC6_(_1655_v0)
				var _pat_let_tv22 = _1655_v0
				_ = _pat_let_tv22
				_1757_v75 = (func() D2 {
					if _1656_v1 {
						return (func() D2 {
							if _1656_v1 {
								return _1757_v75
							}
							return func(_pat_let26_0 D2) D2 {
								return func(_1758_dt__update__tmp_h3 D2) D2 {
									return func(_pat_let27_0 _dafny.Sequence) D2 {
										return func(_1759_dt__update_hcf6_h0 _dafny.Sequence) D2 {
											return Companion_D2_.Create_DC6_(_1759_dt__update_hcf6_h0)
										}(_pat_let27_0)
									}(_pat_let_tv22)
								}(_pat_let26_0)
							}(_1757_v75)
						})()
					}
					return _1757_v75
				})()
				_1656_v1 = _1656_v1
				var _1760_v76 _dafny.Array
				_ = _1760_v76
				var _nwElement0_59 bool = _1656_v1
				_ = _nwElement0_59
				var _nw296 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(8))
				_ = _nw296
				(_nw296).ArraySet1(_nwElement0_59, 0)
				(_nw296).ArraySet1(_1656_v1, 1)
				(_nw296).ArraySet1(_1656_v1, 2)
				(_nw296).ArraySet1(_1656_v1, 3)
				(_nw296).ArraySet1(_1656_v1, 4)
				(_nw296).ArraySet1(_1656_v1, 5)
				(_nw296).ArraySet1(false, 6)
				(_nw296).ArraySet1(_1656_v1, 7)
				_1760_v76 = _nw296
				var _1761_v77 _dafny.Array
				_ = _1761_v77
				_1761_v77 = _1760_v76
				var _1762_v78 _dafny.Map
				_ = _1762_v78
				_1762_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1761_v77, _1760_v76)
				var _1763_v79 _dafny.Map
				_ = _1763_v79
				_1763_v79 = _1762_v78
				var _1764_v80 _dafny.Array
				_ = _1764_v80
				var _nwElement0_60 _dafny.Map = _1763_v79
				_ = _nwElement0_60
				var _nw297 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(13))
				_ = _nw297
				(_nw297).ArraySet1(_nwElement0_60, 0)
				(_nw297).ArraySet1(_1763_v79, 1)
				(_nw297).ArraySet1(_1762_v78, 2)
				(_nw297).ArraySet1(_1762_v78, 3)
				(_nw297).ArraySet1((func() _dafny.Map {
					if _1656_v1 {
						return _1763_v79
					}
					return _1763_v79
				})(), 4)
				(_nw297).ArraySet1(_1763_v79, 5)
				(_nw297).ArraySet1(_1763_v79, 6)
				(_nw297).ArraySet1(_1762_v78, 7)
				(_nw297).ArraySet1(_1763_v79, 8)
				(_nw297).ArraySet1(_1763_v79, 9)
				(_nw297).ArraySet1(_1763_v79, 10)
				(_nw297).ArraySet1(_1762_v78, 11)
				(_nw297).ArraySet1(_1762_v78, 12)
				_1764_v80 = _nw297
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1764_v80), 0))
				_ = _index312
				(_1764_v80).ArraySet1(_1762_v78, (_index312).Int())
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(288), _dafny.ArrayLen((_1764_v80), 0))
				_ = _index313
				(_1764_v80).ArraySet1(_1763_v79, (_index313).Int())
				var _1765_v81 _dafny.Int
				_ = _1765_v81
				_1765_v81 = _dafny.IntOfInt64(958)
				var _1766_v82 _dafny.Map
				_ = _1766_v82
				_1766_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm29(_1765_v81, _1655_v0, globalState), _1760_v76)
				_1766_v82 = (_1766_v82).Update((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Set {
					var _coll64 = _dafny.NewBuilder()
					_ = _coll64
					for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(342), _dafny.IntOfInt64(17))); ; {
						_compr_64, _ok67 := _iter67()
						if !_ok67 {
							break
						}
						var _1767_v83 _dafny.Int
						_1767_v83 = interface{}(_compr_64).(_dafny.Int)
						if ((_dafny.IntOfInt64(342)).Cmp(_1767_v83) <= 0) && ((_1767_v83).Cmp(_dafny.IntOfInt64(17)) < 0) {
							_coll64.Add(Companion_Default___.SafeDivisionInt(_1767_v83, (_dafny.Zero).Minus(_1765_v81)))
						}
					}
					return _coll64.ToSet()
				}()).Cardinality(), _1765_v81)), _1760_v76)
				var _1768_v84 _dafny.Map
				_ = _1768_v84
				_1768_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1765_v81, _1656_v1)
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_1760_v76), 0))
				_ = _index314
				(_1760_v76).ArraySet1((func() bool {
					if (_1768_v84).Contains(_1765_v81) {
						return (_1768_v84).Get(_1765_v81).(bool)
					}
					return _1656_v1
				})(), (_index314).Int())
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_1760_v76), 0))
				_ = _index315
				(_1760_v76).ArraySet1(false, (_index315).Int())
			} else {
				var _1769_v85 _dafny.Int
				_ = _1769_v85
				_1769_v85 = _dafny.IntOfInt64(168)
				var _1770_v86 _dafny.Map
				_ = _1770_v86
				_1770_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(268), _1769_v85)
				var _1771_v87 _dafny.Map
				_ = _1771_v87
				_1771_v87 = _1770_v86
				var _1772_v88 _dafny.Sequence
				_ = _1772_v88
				_1772_v88 = _dafny.SeqOf(_1656_v1)
				var _1773_v89 _dafny.Sequence
				_ = _1773_v89
				_1773_v89 = _dafny.SeqOf(_1772_v88)
				var _1774_v90 _dafny.CodePoint
				_ = _1774_v90
				_1774_v90 = _dafny.CodePoint('k')
				var _1775_v91 D12
				_ = _1775_v91
				_1775_v91 = Companion_D12_.Create_DC22_(_1771_v87, _1773_v89, _1774_v90, _1656_v1, _1655_v0)
				var _1776_v92 _dafny.MultiSet
				_ = _1776_v92
				_1776_v92 = _dafny.MultiSetOf(_1775_v91)
				var _1777_v93 _dafny.Sequence
				_ = _1777_v93
				_1777_v93 = _dafny.SeqOf((_1769_v85).Cmp((_1776_v92).Cardinality()) >= 0)
				_1656_v1 = (_1777_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(755), _dafny.IntOfUint32((_1777_v93).Cardinality()))).Uint32()).(bool)
				var _1778_v94 _dafny.Array
				_ = _1778_v94
				var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw298
				_1778_v94 = _nw298
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1778_v94), 0))
				_ = _index316
				(_1778_v94).ArraySet1(_dafny.IntOfInt64(824), (_index316).Int())
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1778_v94), 0))
				_ = _index317
				(_1778_v94).ArraySet1(_dafny.IntOfUint32((_1655_v0).Cardinality()), (_index317).Int())
				var _1779_v95 _dafny.MultiSet
				_ = _1779_v95
				_1779_v95 = _dafny.MultiSetOf(_1656_v1)
				var _1780_v96 _dafny.Map
				_ = _1780_v96
				_1780_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1779_v95)
				var _1781_v97 *C2
				_ = _1781_v97
				var _nw299 *C2 = New_C2_()
				_ = _nw299
				_nw299.Ctor__(_1774_v90, _1656_v1)
				_1781_v97 = _nw299
				var _1782_v98 _dafny.Map
				_ = _1782_v98
				_1782_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.MultiSet {
					if (_1780_v96).Contains(_1656_v1) {
						return (_1780_v96).Get(_1656_v1).(_dafny.MultiSet)
					}
					return _1779_v95
				})()).IsProperSubsetOf(_1779_v95), _1781_v97)
				var _1783_v99 _dafny.Sequence
				_ = _1783_v99
				_1783_v99 = _dafny.SeqOf(_1769_v85)
				_1782_v98 = (_1782_v98).Update(!(((_1783_v99).Select((Companion_Default___.SafeIndex((_1778_v94).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1778_v94), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1783_v99).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_1769_v85) >= 0), _1781_v97)
				_1656_v1 = (!(!(!(_1656_v1)) || (_1656_v1))) && (_1656_v1)
				var _1784_v100 *C2
				_ = _1784_v100
				var _nw300 *C2 = New_C2_()
				_ = _nw300
				_nw300.Ctor__(_1774_v90, (_1656_v1) || (_1656_v1))
				_1784_v100 = _nw300
			}
			var _1785_v101 _dafny.Int
			_ = _1785_v101
			_1785_v101 = _dafny.IntOfInt64(904)
			var _1786_v102 _dafny.Map
			_ = _1786_v102
			_1786_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1785_v101), _1656_v1)
			var _1787_v103 _dafny.Sequence
			_ = _1787_v103
			_1787_v103 = _dafny.SeqOf(_1785_v101, (_1786_v102).Cardinality())
			var _1788_v104 _dafny.Sequence
			_ = _1788_v104
			_1788_v104 = _dafny.SeqOf(_1787_v103)
			var _1789_v105 *C0
			_ = _1789_v105
			var _nw301 *C0 = New_C0_()
			_ = _nw301
			_nw301.Ctor__(true, _1788_v104)
			_1789_v105 = _nw301
			var _1790_v106 _dafny.MultiSet
			_ = _1790_v106
			_1790_v106 = _dafny.MultiSetOf(_1789_v105, _1789_v105)
			var _rhs335 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1790_v106).Contains(_1789_v105) {
					return (_1790_v106).Multiplicity(_1789_v105)
				}
				return (_dafny.Zero).Minus(_1785_v101)
			})(), _1785_v101)
			_ = _rhs335
			var _rhs336 bool = (Companion_Default___.SafeDivisionInt(_1785_v101, _1785_v101)).Cmp(_1785_v101) == 0
			_ = _rhs336
			var _lhs233 *GlobalState = globalState
			_ = _lhs233
			_lhs233.F1 = _rhs335
			_1656_v1 = _rhs336
			var _1791_v107 _dafny.Sequence
			_ = _1791_v107
			_1791_v107 = _dafny.SeqOf(_1656_v1)
			var _1792_v108 _dafny.MultiSet
			_ = _1792_v108
			_1792_v108 = _dafny.MultiSetOf(_1791_v107)
			var _1793_v109 _dafny.Sequence
			_ = _1793_v109
			_1793_v109 = _dafny.SeqOf(_1791_v107)
			_1792_v108 = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1793_v109, _1793_v109))
		} else {
			if _1656_v1 {
				var _1794_v110 _dafny.Array
				_ = _1794_v110
				var _nwElement0_61 bool = _1656_v1
				_ = _nwElement0_61
				var _nw302 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(2))
				_ = _nw302
				(_nw302).ArraySet1(_nwElement0_61, 0)
				(_nw302).ArraySet1(_1656_v1, 1)
				_1794_v110 = _nw302
				var _1795_v111 _dafny.Map
				_ = _1795_v111
				_1795_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, _dafny.IntOfInt64(-376))
				var _1796_v112 _dafny.Sequence
				_ = _1796_v112
				_1796_v112 = _dafny.SeqOf(_1794_v110)
				var _1797_v113 _dafny.Int
				_ = _1797_v113
				_1797_v113 = _dafny.IntOfInt64(688)
				var _1798_v114 *C2
				_ = _1798_v114
				var _nw303 *C2 = New_C2_()
				_ = _nw303
				_nw303.Ctor__(_dafny.CodePoint('l'), _1656_v1)
				_1798_v114 = _nw303
				var _1799_v115 _dafny.MultiSet
				_ = _1799_v115
				_1799_v115 = _dafny.MultiSetOf(_1798_v114, _1798_v114)
				var _1800_v116 _dafny.Sequence
				_ = _1800_v116
				_1800_v116 = _dafny.SeqOf(_1795_v111, _1795_v111)
				var _1801_v117 _dafny.Map
				_ = _1801_v117
				_1801_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1797_v113, _1794_v110)
				var _1802_v118 _dafny.CodePoint
				_ = _1802_v118
				_1802_v118 = _dafny.CodePoint('h')
				var _1803_v119 _dafny.Sequence
				_ = _1803_v119
				_1803_v119 = _dafny.SeqOf(_1797_v113)
				var _1804_v120 _dafny.Map
				_ = _1804_v120
				_1804_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1803_v119)
				var _1805_v121 _dafny.Set
				_ = _1805_v121
				_1805_v121 = _dafny.SetOf(false)
				var _1806_v122 _dafny.Array
				_ = _1806_v122
				var _nwElement0_62 _dafny.Int = _1797_v113
				_ = _nwElement0_62
				var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(22))
				_ = _nw304
				(_nw304).ArraySet1(_nwElement0_62, 0)
				(_nw304).ArraySet1(_1797_v113, 1)
				(_nw304).ArraySet1(_1797_v113, 2)
				(_nw304).ArraySet1(_1797_v113, 3)
				(_nw304).ArraySet1(_1797_v113, 4)
				(_nw304).ArraySet1(_dafny.IntOfUint32((_1655_v0).Cardinality()), 5)
				(_nw304).ArraySet1(_1797_v113, 6)
				(_nw304).ArraySet1(_1797_v113, 7)
				(_nw304).ArraySet1(_1797_v113, 8)
				(_nw304).ArraySet1(_1797_v113, 9)
				(_nw304).ArraySet1(_1797_v113, 10)
				(_nw304).ArraySet1(_1797_v113, 11)
				(_nw304).ArraySet1(_1797_v113, 12)
				(_nw304).ArraySet1(_1797_v113, 13)
				(_nw304).ArraySet1(_1797_v113, 14)
				(_nw304).ArraySet1((_1804_v120).Cardinality(), 15)
				(_nw304).ArraySet1(_1797_v113, 16)
				(_nw304).ArraySet1(_dafny.IntOfInt64(105), 17)
				(_nw304).ArraySet1((_1805_v121).Cardinality(), 18)
				(_nw304).ArraySet1(_dafny.IntOfInt64(153), 19)
				(_nw304).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("toyypgf")).Cardinality()), 20)
				(_nw304).ArraySet1(_1797_v113, 21)
				_1806_v122 = _nw304
				var _1807_v123 _dafny.Map
				_ = _1807_v123
				_1807_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1802_v118, _1806_v122)
				var _1808_v124 _dafny.Array
				_ = _1808_v124
				var _nwElement0_63 _dafny.Map = (_1795_v111).Merge(_1795_v111)
				_ = _nwElement0_63
				var _nw305 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(24))
				_ = _nw305
				(_nw305).ArraySet1(_nwElement0_63, 0)
				(_nw305).ArraySet1(_1795_v111, 1)
				(_nw305).ArraySet1(_1795_v111, 2)
				(_nw305).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1796_v112).Select((Companion_Default___.SafeIndex(_1797_v113, _dafny.IntOfUint32((_1796_v112).Cardinality()))).Uint32()).(_dafny.Array), _1797_v113)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, (_1799_v115).Cardinality())), 3)
				(_nw305).ArraySet1(_1795_v111, 4)
				(_nw305).ArraySet1((_1795_v111).Update(_1794_v110, _dafny.IntOfInt64(948)), 5)
				(_nw305).ArraySet1(_1795_v111, 6)
				(_nw305).ArraySet1((_1800_v116).Select((Companion_Default___.SafeIndex(_1797_v113, _dafny.IntOfUint32((_1800_v116).Cardinality()))).Uint32()).(_dafny.Map), 7)
				(_nw305).ArraySet1(_1795_v111, 8)
				(_nw305).ArraySet1(_1795_v111, 9)
				(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if (_1801_v117).Contains(_dafny.IntOfUint32((_1655_v0).Cardinality())) {
						return (_1801_v117).Get(_dafny.IntOfUint32((_1655_v0).Cardinality())).(_dafny.Array)
					}
					return _1794_v110
				})(), _1797_v113), 10)
				(_nw305).ArraySet1(_1795_v111, 11)
				(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, _1797_v113), 12)
				(_nw305).ArraySet1(_1795_v111, 13)
				(_nw305).ArraySet1(_1795_v111, 14)
				(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, (_1807_v123).Cardinality()), 15)
				(_nw305).ArraySet1((_1795_v111).Merge(_1795_v111), 16)
				(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, _1797_v113), 17)
				(_nw305).ArraySet1((_1795_v111).Merge(_1795_v111), 18)
				(_nw305).ArraySet1(_1795_v111, 19)
				(_nw305).ArraySet1(_1795_v111, 20)
				(_nw305).ArraySet1((_1795_v111).Merge(_1795_v111), 21)
				(_nw305).ArraySet1(_1795_v111, 22)
				(_nw305).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1794_v110, _1797_v113), 23)
				_1808_v124 = _nw305
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_1808_v124), 0))
				_ = _index318
				(_1808_v124).ArraySet1(_1795_v111, (_index318).Int())
				var _1809_v125 _dafny.MultiSet
				_ = _1809_v125
				_1809_v125 = _dafny.MultiSetOf(_1797_v113, _1797_v113)
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_1808_v124), 0))
				_ = _index319
				(_1808_v124).ArraySet1(((_1795_v111).Update(_1794_v110, _1797_v113)).Update(_1794_v110, Companion_Default___.SafeModuloInt(_1797_v113, (_1809_v125).Cardinality())), (_index319).Int())
				var _rhs337 _dafny.CodePoint = _1802_v118
				_ = _rhs337
				var _rhs338 _dafny.Int = (func() _dafny.Int {
					if (func() bool {
						if _1656_v1 {
							return false
						}
						return _1656_v1
					})() {
						return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(811), _1797_v113)
					}
					return _1797_v113
				})()
				_ = _rhs338
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				_1802_v118 = _rhs337
				_lhs234.F1 = _rhs338
				var _1810_v126 D6
				_ = _1810_v126
				_1810_v126 = Companion_D6_.Create_DC13_(_1803_v119)
				_1810_v126 = _1810_v126
				var _1811_v127 _dafny.Map
				_ = _1811_v127
				_1811_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1797_v113)
				var _1812_v128 _dafny.Map
				_ = _1812_v128
				_1812_v128 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1797_v113, (_1811_v127).Cardinality())).Update(_1797_v113, _dafny.IntOfInt64(-464))
				var _1813_v129 _dafny.Sequence
				_ = _1813_v129
				_1813_v129 = _dafny.SeqOf(_1656_v1)
				var _1814_v130 _dafny.Sequence
				_ = _1814_v130
				_1814_v130 = _dafny.SeqOf(_1813_v129)
				var _1815_v131 D12
				_ = _1815_v131
				_1815_v131 = Companion_D12_.Create_DC22_(_1812_v128, _1814_v130, _1802_v118, _1656_v1, _1655_v0)
				var _1816_v132 _dafny.Set
				_ = _1816_v132
				_1816_v132 = _dafny.SetOf(_1794_v110)
				var _1817_v133 *C2
				_ = _1817_v133
				var _nw306 *C2 = New_C2_()
				_ = _nw306
				_nw306.Ctor__((func() _dafny.CodePoint {
					if _1656_v1 {
						return _1802_v118
					}
					return (_1815_v131).Dtor_cf26()
				})(), (_1816_v132).Contains(_1794_v110))
				_1817_v133 = _nw306
				var _1818_v134 _dafny.Set
				_ = _1818_v134
				_1818_v134 = _dafny.SetOf(Companion_D11_.Create_DC19_(_dafny.CodePoint('l')))
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_1794_v110), 0))
				_ = _index320
				(_1794_v110).ArraySet1((_1813_v129).Select((Companion_Default___.SafeIndex(_1797_v113, _dafny.IntOfUint32((_1813_v129).Cardinality()))).Uint32()).(bool), (_index320).Int())
				var _1819_v135 _dafny.Map
				_ = _1819_v135
				_1819_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1797_v113, (_1818_v134).Union(_1818_v134))
				var _1820_v136 _dafny.Map
				_ = _1820_v136
				_1820_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(137), _1656_v1)
				var _1821_v137 _dafny.Map
				_ = _1821_v137
				_1821_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1820_v136, _1797_v113)
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_1794_v110), 0))
				_ = _index321
				var _rhs339 bool = _1656_v1
				_ = _rhs339
				var _rhs340 _dafny.Set = (func() _dafny.Set {
					if (_1819_v135).Contains(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
						if (_1821_v137).Contains(_1820_v136) {
							return (_1821_v137).Get(_1820_v136).(_dafny.Int)
						}
						return (_dafny.MultiSetOf(_1794_v110)).Cardinality()
					})(), _1797_v113)) {
						return (_1819_v135).Get(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
							if (_1821_v137).Contains(_1820_v136) {
								return (_1821_v137).Get(_1820_v136).(_dafny.Int)
							}
							return (_dafny.MultiSetOf(_1794_v110)).Cardinality()
						})(), _1797_v113)).(_dafny.Set)
					}
					return _1818_v134
				})()
				_ = _rhs340
				var _rhs341 bool = _1656_v1
				_ = _rhs341
				var _lhs235 _dafny.Array = _1794_v110
				_ = _lhs235
				var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(447), _dafny.ArrayLen((_1794_v110), 0))
				_ = _lhs236
				_1656_v1 = _rhs339
				_1818_v134 = _rhs340
				(_lhs235).ArraySet1(_rhs341, (_lhs236).Int())
			} else {
				var _1822_v138 _dafny.Int
				_ = _1822_v138
				_1822_v138 = _dafny.IntOfInt64(-987)
				(globalState).F1 = (_dafny.Zero).Minus(_1822_v138)
				var _1823_v139 _dafny.Sequence
				_ = _1823_v139
				_1823_v139 = _dafny.SeqOf(_1822_v138, _1822_v138)
				_1656_v1 = (_this).Fm6(_1823_v139, globalState)
				var _1824_v140 _dafny.Array
				_ = _1824_v140
				var _nw307 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw307
				_1824_v140 = _nw307
				var _1825_v141 _dafny.Array
				_ = _1825_v141
				var _nwElement0_64 _dafny.Array = _1824_v140
				_ = _nwElement0_64
				var _nw308 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(13))
				_ = _nw308
				(_nw308).ArraySet1(_nwElement0_64, 0)
				(_nw308).ArraySet1(_1824_v140, 1)
				(_nw308).ArraySet1(_1824_v140, 2)
				(_nw308).ArraySet1(_1824_v140, 3)
				(_nw308).ArraySet1(_1824_v140, 4)
				(_nw308).ArraySet1(_1824_v140, 5)
				(_nw308).ArraySet1(_1824_v140, 6)
				(_nw308).ArraySet1(_1824_v140, 7)
				(_nw308).ArraySet1(_1824_v140, 8)
				(_nw308).ArraySet1(_1824_v140, 9)
				(_nw308).ArraySet1(_1824_v140, 10)
				(_nw308).ArraySet1(_1824_v140, 11)
				(_nw308).ArraySet1(_1824_v140, 12)
				_1825_v141 = _nw308
				var _rhs342 _dafny.Array = _1825_v141
				_ = _rhs342
				var _rhs343 bool = _1656_v1
				_ = _rhs343
				_1825_v141 = _rhs342
				_1656_v1 = _rhs343
				(globalState).F1 = _1822_v138
				(_this).M9(_dafny.UnicodeSeqOfUtf8Bytes("cdpjpt"), globalState)
			}
			var _1826_v142 _dafny.Int
			_ = _1826_v142
			_1826_v142 = _dafny.IntOfInt64(790)
			var _1827_v143 _dafny.Map
			_ = _1827_v143
			_1827_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(763), _1826_v142)
			(globalState).F1 = (func() _dafny.Int {
				if (_1827_v143).Contains(_1826_v142) {
					return (_1827_v143).Get(_1826_v142).(_dafny.Int)
				}
				return (_1826_v142).Minus(_1826_v142)
			})()
			var _1828_v144 _dafny.Sequence
			_ = _1828_v144
			_1828_v144 = _dafny.SeqOf(_1656_v1)
			var _1829_v145 _dafny.MultiSet
			_ = _1829_v145
			_1829_v145 = _dafny.MultiSetOf(_1828_v144)
			_1829_v145 = _dafny.MultiSetOf(_1828_v144)
			var _1830_v146 _dafny.Set
			_ = _1830_v146
			_1830_v146 = _dafny.SetOf(_1656_v1)
			if (_1830_v146).IsSubsetOf((func() _dafny.Set {
				if _1656_v1 {
					return _dafny.SetOf(_1656_v1, _1656_v1, _1656_v1, _1656_v1)
				}
				return _1830_v146
			})()) {
				var _1831_v147 _dafny.Array
				_ = _1831_v147
				var _nw309 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
				_ = _nw309
				_1831_v147 = _nw309
				var _1832_v148 _dafny.Map
				_ = _1832_v148
				_1832_v148 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1826_v142)
				var _1833_v149 D15
				_ = _1833_v149
				_1833_v149 = Companion_D15_.Create_DC28_(_1656_v1, _1656_v1)
				var _pat_let_tv23 = _1656_v1
				_ = _pat_let_tv23
				var _pat_let_tv24 = globalState
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1827_v143
				_ = _pat_let_tv25
				var _rhs344 bool = (func(_pat_let28_0 D15) D15 {
					return func(_1834_dt__update__tmp_h4 D15) D15 {
						return func(_pat_let29_0 bool) D15 {
							return func(_1835_dt__update_hcf36_h0 bool) D15 {
								return Companion_D15_.Create_DC28_(_1835_dt__update_hcf36_h0, (_1834_dt__update__tmp_h4).Dtor_cf37())
							}(_pat_let29_0)
						}((Companion_D11_.Create_DC20_(Companion_Default___.Fm0(_pat_let_tv23, _pat_let_tv24), (_pat_let_tv25).Cardinality())).Dtor_cf21())
					}(_pat_let28_0)
				}(_1833_v149)).Dtor_cf37()
				_ = _rhs344
				var _rhs345 _dafny.Int = _1826_v142
				_ = _rhs345
				var _rhs346 _dafny.Set = (_1830_v146).Intersection(_1830_v146)
				_ = _rhs346
				var _rhs347 _dafny.Array = _1831_v147
				_ = _rhs347
				var _rhs348 _dafny.Map = _1832_v148
				_ = _rhs348
				var _lhs237 *GlobalState = globalState
				_ = _lhs237
				_1656_v1 = _rhs344
				_lhs237.F1 = _rhs345
				_1830_v146 = _rhs346
				_1831_v147 = _rhs347
				_1832_v148 = _rhs348
				var _1836_v150 _dafny.MultiSet
				_ = _1836_v150
				_1836_v150 = _dafny.MultiSetOf(_1656_v1)
				var _1837_v151 _dafny.Sequence
				_ = _1837_v151
				_1837_v151 = _dafny.SeqOf(_1826_v142)
				var _1838_v152 _dafny.Array
				_ = _1838_v152
				var _nwElement0_65 _dafny.Int = _1826_v142
				_ = _nwElement0_65
				var _nw310 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(19))
				_ = _nw310
				(_nw310).ArraySet1(_nwElement0_65, 0)
				(_nw310).ArraySet1(_1826_v142, 1)
				(_nw310).ArraySet1((_this).Fm7(_1656_v1, globalState), 2)
				(_nw310).ArraySet1(_dafny.IntOfInt64(413), 3)
				(_nw310).ArraySet1(_1826_v142, 4)
				(_nw310).ArraySet1(_dafny.IntOfInt64(342), 5)
				(_nw310).ArraySet1(_1826_v142, 6)
				(_nw310).ArraySet1(_1826_v142, 7)
				(_nw310).ArraySet1((_1836_v150).Cardinality(), 8)
				(_nw310).ArraySet1(_1826_v142, 9)
				(_nw310).ArraySet1(_dafny.IntOfUint32((_1837_v151).Cardinality()), 10)
				(_nw310).ArraySet1(_1826_v142, 11)
				(_nw310).ArraySet1(_1826_v142, 12)
				(_nw310).ArraySet1(_1826_v142, 13)
				(_nw310).ArraySet1(_dafny.IntOfInt64(166), 14)
				(_nw310).ArraySet1(_1826_v142, 15)
				(_nw310).ArraySet1((func() _dafny.Int {
					if (_1832_v148).Contains(true) {
						return (_1832_v148).Get(true).(_dafny.Int)
					}
					return _1826_v142
				})(), 16)
				(_nw310).ArraySet1(_1826_v142, 17)
				(_nw310).ArraySet1((_dafny.Zero).Minus(_1826_v142), 18)
				_1838_v152 = _nw310
				var _1839_v153 _dafny.Sequence
				_ = _1839_v153
				_1839_v153 = _dafny.SeqOf(_1838_v152, _1838_v152, _1838_v152)
				var _1840_v154 _dafny.Map
				_ = _1840_v154
				_1840_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1826_v142), (_1839_v153).Select((Companion_Default___.SafeIndex(_1826_v142, _dafny.IntOfUint32((_1839_v153).Cardinality()))).Uint32()).(_dafny.Array))
				var _rhs349 _dafny.Int = (func() _dafny.Int {
					if _1656_v1 {
						return (func() _dafny.Int {
							if _1656_v1 {
								return _dafny.IntOfUint32((_1655_v0).Cardinality())
							}
							return _1826_v142
						})()
					}
					return (_1826_v142).Minus(_1826_v142)
				})()
				_ = _rhs349
				var _rhs350 _dafny.Map = _1840_v154
				_ = _rhs350
				var _lhs238 *GlobalState = globalState
				_ = _lhs238
				_lhs238.F1 = _rhs349
				_1840_v154 = _rhs350
				var _1841_v155 D0
				_ = _1841_v155
				_1841_v155 = Companion_D0_.Create_DC1_()
				var _1842_v156 D0
				_ = _1842_v156
				_1842_v156 = Companion_D0_.Create_DC3_(_1841_v155)
				var _1843_v157 _dafny.CodePoint
				_ = _1843_v157
				_1843_v157 = _dafny.CodePoint('u')
				var _1844_v158 *C2
				_ = _1844_v158
				var _nw311 *C2 = New_C2_()
				_ = _nw311
				_nw311.Ctor__(_1843_v157, Companion_Default___.Fm0(_1656_v1, globalState))
				_1844_v158 = _nw311
				var _1845_v159 _dafny.Map
				_ = _1845_v159
				_1845_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1842_v156, _1844_v158)
				_1845_v159 = (_1845_v159).Update(_1842_v156, _1844_v158)
				var _1846_v160 D11
				_ = _1846_v160
				_1846_v160 = Companion_D11_.Create_DC20_((_1844_v158).Fm6(_1837_v151, globalState), _1826_v142)
				_1832_v148 = (_1832_v148).Update((_1846_v160).Dtor_cf21(), _1826_v142)
				_1838_v152 = _1838_v152
			} else {
				_1656_v1 = !(_1656_v1) || (_1656_v1)
				var _1847_v161 _dafny.Map
				_ = _1847_v161
				_1847_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1655_v0, _1656_v1)
				_1847_v161 = (_1847_v161).Update(_dafny.Companion_Sequence_.Concatenate(_1655_v0, _1655_v0), !(_1656_v1))
				var _1848_v162 *C1
				_ = _1848_v162
				var _nw312 *C1 = New_C1_()
				_ = _nw312
				_nw312.Ctor__()
				_1848_v162 = _nw312
				var _1849_v163 _dafny.Map
				_ = _1849_v163
				_1849_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1656_v1)
				var _1850_v164 D12
				_ = _1850_v164
				_1850_v164 = Companion_D12_.Create_DC21_(_1849_v163)
				var _1851_v165 _dafny.Map
				_ = _1851_v165
				_1851_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v1, _1850_v164)
				_1851_v165 = _1851_v165
				_1656_v1 = _1656_v1
			}
			(globalState).F1 = Companion_Default___.SafeModuloInt(_1826_v142, _1826_v142)
		}
		var _1852_v166 _dafny.Int
		_ = _1852_v166
		_1852_v166 = _dafny.IntOfInt64(892)
		var _1853_v167 D1
		_ = _1853_v167
		_1853_v167 = Companion_D1_.Create_DC4_(_1852_v166)
		var _1854_v168 D2
		_ = _1854_v168
		_1854_v168 = Companion_D2_.Create_DC8_(_1656_v1)
		var _pat_let_tv26 = _1852_v166
		_ = _pat_let_tv26
		_1853_v167 = func(_pat_let30_0 D1) D1 {
			return func(_1855_dt__update__tmp_h5 D1) D1 {
				return func(_pat_let31_0 _dafny.Int) D1 {
					return func(_1856_dt__update_hcf4_h0 _dafny.Int) D1 {
						return Companion_D1_.Create_DC4_(_1856_dt__update_hcf4_h0)
					}(_pat_let31_0)
				}(_pat_let_tv26)
			}(_pat_let30_0)
		}((_this).Fm5(_1854_v168, globalState))
		var _1857_v169 _dafny.Sequence
		_ = _1857_v169
		_1857_v169 = _dafny.SeqOf(_dafny.IntOfInt64(-263))
		var _1858_v170 _dafny.Sequence
		_ = _1858_v170
		_1858_v170 = _dafny.SeqOf(_1857_v169, _1857_v169, _1857_v169)
		var _1859_v171 _dafny.Sequence
		_ = _1859_v171
		_1859_v171 = _dafny.SeqOf(_1857_v169, _1857_v169, _1857_v169, _1857_v169, (_1858_v170).Select((Companion_Default___.SafeIndex(_1852_v166, _dafny.IntOfUint32((_1858_v170).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1860_v172 *C0
		_ = _1860_v172
		var _nw313 *C0 = New_C0_()
		_ = _nw313
		_nw313.Ctor__((func() bool {
			if _1656_v1 {
				return _1656_v1
			}
			return !(_1656_v1)
		})(), _dafny.Companion_Sequence_.Concatenate(_1858_v170, _1858_v170))
		_1860_v172 = _nw313
		var _1861_v173 _dafny.Map
		_ = _1861_v173
		_1861_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1860_v172).F11()), _1852_v166)
		var _1862_v174 _dafny.MultiSet
		_ = _1862_v174
		_1862_v174 = _dafny.MultiSetOf(_1656_v1)
		var _1863_v175 _dafny.Sequence
		_ = _1863_v175
		_1863_v175 = _dafny.SeqOf((_1854_v168).Dtor_cf9(), _1656_v1, (_1860_v172).F11())
		var _1864_v176 _dafny.Map
		_ = _1864_v176
		_1864_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-512), _dafny.IntOfUint32((_dafny.SeqOf(_1656_v1, false)).Cardinality()))
		var _1865_v177 _dafny.Set
		_ = _1865_v177
		_1865_v177 = _dafny.SetOf((_1860_v172).F11(), (_1860_v172).F11())
		var _1866_v178 _dafny.Array
		_ = _1866_v178
		var _nwElement0_66 _dafny.Int = _1852_v166
		_ = _nwElement0_66
		var _nw314 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(27))
		_ = _nw314
		(_nw314).ArraySet1(_nwElement0_66, 0)
		(_nw314).ArraySet1(((_1861_v173).Merge(Companion_Default___.Fm57((_dafny.Zero).Minus(_1852_v166), (_1862_v174).Cardinality(), globalState))).Cardinality(), 1)
		(_nw314).ArraySet1(_1852_v166, 2)
		(_nw314).ArraySet1((_1852_v166).Plus(_1852_v166), 3)
		(_nw314).ArraySet1(_1852_v166, 4)
		(_nw314).ArraySet1((_dafny.Zero).Minus(_1852_v166), 5)
		(_nw314).ArraySet1(_1852_v166, 6)
		(_nw314).ArraySet1(_1852_v166, 7)
		(_nw314).ArraySet1(_1852_v166, 8)
		(_nw314).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pmtyrbao")).Cardinality()), _1852_v166), 9)
		(_nw314).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfUint32((_1857_v169).Cardinality())).Minus(_1852_v166)), 10)
		(_nw314).ArraySet1(_1852_v166, 11)
		(_nw314).ArraySet1(_dafny.IntOfUint32((_1863_v175).Cardinality()), 12)
		(_nw314).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_1852_v166), _1852_v166), 13)
		(_nw314).ArraySet1(_1852_v166, 14)
		(_nw314).ArraySet1((func() _dafny.Int {
			if (_1864_v176).Contains((func() _dafny.Int {
				if (_1864_v176).Contains(_1852_v166) {
					return (_1864_v176).Get(_1852_v166).(_dafny.Int)
				}
				return _1852_v166
			})()) {
				return (_1864_v176).Get((func() _dafny.Int {
					if (_1864_v176).Contains(_1852_v166) {
						return (_1864_v176).Get(_1852_v166).(_dafny.Int)
					}
					return _1852_v166
				})()).(_dafny.Int)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1852_v166, _1852_v166)).Cardinality()
		})(), 15)
		(_nw314).ArraySet1(_1852_v166, 16)
		(_nw314).ArraySet1(_dafny.IntOfUint32((_1857_v169).Cardinality()), 17)
		(_nw314).ArraySet1((func() _dafny.Int {
			if false {
				return _1852_v166
			}
			return _dafny.IntOfUint32((_1863_v175).Cardinality())
		})(), 18)
		(_nw314).ArraySet1(_1852_v166, 19)
		(_nw314).ArraySet1(_1852_v166, 20)
		(_nw314).ArraySet1(_dafny.IntOfUint32((_1655_v0).Cardinality()), 21)
		(_nw314).ArraySet1(Companion_Default___.SafeModuloInt(_1852_v166, _1852_v166), 22)
		(_nw314).ArraySet1(((func() _dafny.Set {
			if false {
				return _1865_v177
			}
			return _1865_v177
		})()).Cardinality(), 23)
		(_nw314).ArraySet1((_1852_v166).Plus(_1852_v166), 24)
		(_nw314).ArraySet1(_1852_v166, 25)
		(_nw314).ArraySet1(_1852_v166, 26)
		_1866_v178 = _nw314
		var _rhs351 _dafny.Array = _1866_v178
		_ = _rhs351
		var _rhs352 bool = _1656_v1
		_ = _rhs352
		_1866_v178 = _rhs351
		_1656_v1 = _rhs352
		r0 = _1852_v166
		return r0
	}
}
func (_this *C6) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _1867_v0 _dafny.Int
		_ = _1867_v0
		_1867_v0 = _dafny.IntOfInt64(442)
		var _1868_v1 _dafny.Sequence
		_ = _1868_v1
		_1868_v1 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(p0, Companion_Default___.Fm0(p0, globalState), globalState), _1867_v0))
		_1868_v1 = _1868_v1
		var _1869_v2 _dafny.Sequence
		_ = _1869_v2
		_1869_v2 = _dafny.UnicodeSeqOfUtf8Bytes("yqnx")
		_1867_v0 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1869_v2).Cardinality()), _1867_v0)
		var _1870_i0 _dafny.Int
		_ = _1870_i0
		_1870_i0 = _dafny.Zero
		{
			for p0 {
				{
					if (_1870_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L17
					}
					_1870_i0 = (_1870_i0).Plus(_dafny.One)
					var _1871_v3 bool
					_ = _1871_v3
					_1871_v3 = true
					var _1872_v4 _dafny.Array
					_ = _1872_v4
					var _nwElement0_67 bool = p0
					_ = _nwElement0_67
					var _nw315 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(5))
					_ = _nw315
					(_nw315).ArraySet1(_nwElement0_67, 0)
					(_nw315).ArraySet1((p0) && (p2), 1)
					(_nw315).ArraySet1(Companion_Default___.Fm0(p2, globalState), 2)
					(_nw315).ArraySet1(!(_1871_v3), 3)
					(_nw315).ArraySet1(p0, 4)
					_1872_v4 = _nw315
					var _1873_v5 _dafny.Sequence
					_ = _1873_v5
					_1873_v5 = _dafny.SeqOf(!(p0))
					var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1872_v4), 0))
					_ = _index322
					(_1872_v4).ArraySet1((_1873_v5).Select((Companion_Default___.SafeIndex(_1867_v0, _dafny.IntOfUint32((_1873_v5).Cardinality()))).Uint32()).(bool), (_index322).Int())
					var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1872_v4), 0))
					_ = _index323
					var _rhs353 bool = (_1873_v5).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_1867_v0, _1867_v0), _dafny.IntOfUint32((_1873_v5).Cardinality()))).Uint32()).(bool)
					_ = _rhs353
					var _rhs354 bool = !(p0) || (p2)
					_ = _rhs354
					var _lhs239 _dafny.Array = _1872_v4
					_ = _lhs239
					var _lhs240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(646), _dafny.ArrayLen((_1872_v4), 0))
					_ = _lhs240
					_1871_v3 = _rhs353
					(_lhs239).ArraySet1(_rhs354, (_lhs240).Int())
					var _1874_v6 _dafny.Set
					_ = _1874_v6
					_1874_v6 = _dafny.SetOf(_1871_v3, p2)
					var _1875_v7 _dafny.Sequence
					_ = _1875_v7
					_1875_v7 = _dafny.SeqOf(_1874_v6, _1874_v6, _1874_v6, Companion_Default___.Fm49(globalState))
					var _1876_v8 D16
					_ = _1876_v8
					_1876_v8 = Companion_D16_.Create_DC30_(_1875_v7)
					(globalState).F1 = _dafny.IntOfUint32(((_1876_v8).Dtor_cf39()).Cardinality())
					(globalState).F1 = (Companion_Default___.Fm58(globalState)).Dtor_cf22()
					var _1877_v9 D0
					_ = _1877_v9
					_1877_v9 = Companion_D0_.Create_DC1_()
					var _1878_v10 _dafny.Set
					_ = _1878_v10
					_1878_v10 = _dafny.SetOf(_1877_v9, _1877_v9, _1877_v9)
					var _1879_v11 _dafny.Set
					_ = _1879_v11
					_1879_v11 = _1878_v10
					_1879_v11 = _1879_v11
					goto C17
				}
			C17:
			}
			goto L17
		}
	L17:
		var _1880_v12 _dafny.Array
		_ = _1880_v12
		var _nw316 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
		_ = _nw316
		_1880_v12 = _nw316
		var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1880_v12), 0))
		_ = _index324
		(_1880_v12).ArraySet1(_1869_v2, (_index324).Int())
		var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1880_v12), 0))
		_ = _index325
		(_1880_v12).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1869_v2, _dafny.UnicodeSeqOfUtf8Bytes("q")), (_index325).Int())
		var _1881_v13 bool
		_ = _1881_v13
		_1881_v13 = false
		var _1882_v14 _dafny.Sequence
		_ = _1882_v14
		_1882_v14 = _dafny.SeqOf((_1880_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1880_v12), 0))).Int()).(_dafny.Sequence), (_1880_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1880_v12), 0))).Int()).(_dafny.Sequence))
		_1881_v13 = (func() bool {
			if p2 {
				return true
			}
			return !_dafny.Companion_Sequence_.Contains(_1882_v14, (_1880_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_1880_v12), 0))).Int()).(_dafny.Sequence))
		})()
		var _1883_v15 _dafny.Array
		_ = _1883_v15
		var _nwElement0_68 bool = p0
		_ = _nwElement0_68
		var _nw317 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(3))
		_ = _nw317
		(_nw317).ArraySet1(_nwElement0_68, 0)
		(_nw317).ArraySet1(_1881_v13, 1)
		(_nw317).ArraySet1(p0, 2)
		_1883_v15 = _nw317
		var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1883_v15), 0))
		_ = _index326
		(_1883_v15).ArraySet1(_1881_v13, (_index326).Int())
		var _1884_v16 _dafny.Map
		_ = _1884_v16
		_1884_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1867_v0, _1867_v0)
		var _1885_v17 _dafny.Map
		_ = _1885_v17
		_1885_v17 = _1884_v16
		var _pat_let_tv27 = p0
		_ = _pat_let_tv27
		var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(125), _dafny.ArrayLen((_1883_v15), 0))
		_ = _index327
		(_1883_v15).ArraySet1(!(func(_source24 _dafny.Map) bool {
			var _1886___mcc_h0 _dafny.Map = _source24
			_ = _1886___mcc_h0
			var _1887_cf11 _dafny.Map = _1886___mcc_h0
			_ = _1887_cf11
			return _pat_let_tv27
		}(_1884_v16)), (_index327).Int())
	}
}
func (_this *C6) M8(p0 bool, p1 bool, p2 _dafny.Array, globalState *GlobalState) (_dafny.Int, *C1) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 *C1 = (*C1)(nil)
		_ = r1
		var _1888_v0 _dafny.Array
		_ = _1888_v0
		var _nw318 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw318
		_1888_v0 = _nw318
		var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_1888_v0), 0))
		_ = _index328
		(_1888_v0).ArraySet1(Companion_Default___.Fm0(p1, globalState), (_index328).Int())
		var _1889_v1 _dafny.Int
		_ = _1889_v1
		_1889_v1 = _dafny.IntOfInt64(-792)
		var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))
		_ = _index329
		(p2).ArraySet1(_1889_v1, (_index329).Int())
		var _1890_v2 _dafny.Sequence
		_ = _1890_v2
		_1890_v2 = _dafny.UnicodeSeqOfUtf8Bytes("oyvrsduxe")
		var _1891_v3 _dafny.Sequence
		_ = _1891_v3
		_1891_v3 = _dafny.SeqOf(p0, p1)
		var _1892_v4 _dafny.MultiSet
		_ = _1892_v4
		_1892_v4 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thordvy")).Cardinality()))
		var _1893_v5 D1
		_ = _1893_v5
		_1893_v5 = Companion_D1_.Create_DC4_(_1889_v1)
		var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_1888_v0), 0))
		_ = _index330
		var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))
		_ = _index331
		var _rhs355 bool = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(941))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg103 _dafny.Int) interface{} {
				return coer103(arg103)
			}
		}(func(_1894_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality())).Cmp(_dafny.IntOfInt64(-462)) == 0
		_ = _rhs355
		var _rhs356 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1890_v2, _1890_v2)).Cardinality())
		_ = _rhs356
		var _rhs357 _dafny.Int = ((_dafny.Zero).Minus((_1889_v1).Minus(_1889_v1))).Times(Companion_Default___.SafeDivisionInt(_1889_v1, _1889_v1))
		_ = _rhs357
		var _rhs358 _dafny.Int = ((_this).Fm4(_1891_v3, _1889_v1, _1889_v1, _1892_v4, globalState)).Minus((_1893_v5).Dtor_cf4())
		_ = _rhs358
		var _lhs241 _dafny.Array = _1888_v0
		_ = _lhs241
		var _lhs242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(282), _dafny.ArrayLen((_1888_v0), 0))
		_ = _lhs242
		var _lhs243 _dafny.Array = p2
		_ = _lhs243
		var _lhs244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))
		_ = _lhs244
		var _lhs245 *GlobalState = globalState
		_ = _lhs245
		var _lhs246 *GlobalState = globalState
		_ = _lhs246
		(_lhs241).ArraySet1(_rhs355, (_lhs242).Int())
		(_lhs243).ArraySet1(_rhs356, (_lhs244).Int())
		_lhs245.F1 = _rhs357
		_lhs246.F1 = _rhs358
		_1890_v2 = _1890_v2
		var _1895_v6 _dafny.Map
		_ = _1895_v6
		_1895_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int))
		var _1896_v7 _dafny.CodePoint
		_ = _1896_v7
		_1896_v7 = _dafny.CodePoint('u')
		var _1897_v8 D0
		_ = _1897_v8
		var _1898_v9 bool
		_ = _1898_v9
		var _out46 D0
		_ = _out46
		var _out47 bool
		_ = _out47
		_out46, _out47 = Companion_Default___.M0(_1888_v0, Companion_Default___.Fm36(Companion_Default___.Fm3(p1, p1, globalState), (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int), _1889_v1, _1895_v6, globalState), _dafny.Companion_Sequence_.Equal(_1890_v2, _1890_v2), _1896_v7, globalState)
		_1897_v8 = _out46
		_1898_v9 = _out47
		_1896_v7 = _1896_v7
		var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))
		_ = _index332
		(p2).ArraySet1(Companion_Default___.SafeModuloInt((_this).Fm7(p0, globalState), _1889_v1), (_index332).Int())
		var _1899_v10 _dafny.Array
		_ = _1899_v10
		_1899_v10 = p2
		var _1900_v11 _dafny.Array
		_ = _1900_v11
		var _nwElement0_69 _dafny.Array = p2
		_ = _nwElement0_69
		var _nw319 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(9))
		_ = _nw319
		(_nw319).ArraySet1(_nwElement0_69, 0)
		(_nw319).ArraySet1(p2, 1)
		(_nw319).ArraySet1(p2, 2)
		(_nw319).ArraySet1(p2, 3)
		(_nw319).ArraySet1(p2, 4)
		(_nw319).ArraySet1(p2, 5)
		(_nw319).ArraySet1((_1899_v10), 6)
		(_nw319).ArraySet1(p2, 7)
		(_nw319).ArraySet1(p2, 8)
		_1900_v11 = _nw319
		var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1900_v11), 0))
		_ = _index333
		(_1900_v11).ArraySet1(p2, (_index333).Int())
		var _1901_v12 _dafny.Set
		_ = _1901_v12
		_1901_v12 = _dafny.SetOf(_1896_v7, _1896_v7)
		var _1902_v13 _dafny.Set
		_ = _1902_v13
		_1902_v13 = _dafny.SetOf((_dafny.SetOf(_1896_v7)).Union(_1901_v12))
		var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1900_v11), 0))
		_ = _index334
		var _rhs359 _dafny.Array = p2
		_ = _rhs359
		var _rhs360 _dafny.Set = (_1902_v13).Union(_1902_v13)
		_ = _rhs360
		var _lhs247 _dafny.Array = _1900_v11
		_ = _lhs247
		var _lhs248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_1900_v11), 0))
		_ = _lhs248
		(_lhs247).ArraySet1(_rhs359, (_lhs248).Int())
		_1902_v13 = _rhs360
		r0 = (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(112), _dafny.ArrayLen((p2), 0))).Int()).(_dafny.Int)
		var _1903_v14 *C1
		_ = _1903_v14
		var _nw320 *C1 = New_C1_()
		_ = _nw320
		_nw320.Ctor__()
		_1903_v14 = _nw320
		r1 = _1903_v14
		return r0, r1
	}
}
func (_this *C6) M9(p0 _dafny.Sequence, globalState *GlobalState) {
	{
		var _1904_v0 _dafny.Int
		_ = _1904_v0
		_1904_v0 = _dafny.IntOfInt64(-163)
		var _1905_v1 bool
		_ = _1905_v1
		_1905_v1 = true
		var _1906_v2 _dafny.Set
		_ = _1906_v2
		_1906_v2 = _dafny.SetOf(p0)
		var _1907_v3 _dafny.Sequence
		_ = _1907_v3
		_1907_v3 = _dafny.SeqOf(true)
		var _1908_v4 _dafny.MultiSet
		_ = _1908_v4
		_1908_v4 = _dafny.MultiSetOf(_1904_v0)
		var _1909_v5 _dafny.Map
		_ = _1909_v5
		_1909_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1904_v0, _1904_v0)
		var _1910_v6 _dafny.Map
		_ = _1910_v6
		_1910_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1904_v0, _1905_v1)
		var _1911_v7 D16
		_ = _1911_v7
		_1911_v7 = Companion_D16_.Create_DC32_((func() bool {
			if (_1910_v6).Contains(_dafny.IntOfInt64(230)) {
				return (_1910_v6).Get(_dafny.IntOfInt64(230)).(bool)
			}
			return _1905_v1
		})())
		var _1912_v8 _dafny.Sequence
		_ = _1912_v8
		_1912_v8 = _dafny.SeqOf(_1911_v7, Companion_D16_.Create_DC32_(!(_1905_v1)), _1911_v7)
		var _1913_v9 _dafny.Sequence
		_ = _1913_v9
		_1913_v9 = _dafny.SeqOf(_1904_v0)
		var _1914_v10 _dafny.Array
		_ = _1914_v10
		var _nwElement0_70 _dafny.Int = _dafny.IntOfInt64(167)
		_ = _nwElement0_70
		var _nw321 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(24))
		_ = _nw321
		(_nw321).ArraySet1(_nwElement0_70, 0)
		(_nw321).ArraySet1((_1904_v0).Plus(_1904_v0), 1)
		(_nw321).ArraySet1((func() _dafny.Int {
			if _1905_v1 {
				return _1904_v0
			}
			return _1904_v0
		})(), 2)
		(_nw321).ArraySet1(_1904_v0, 3)
		(_nw321).ArraySet1((_1906_v2).Cardinality(), 4)
		(_nw321).ArraySet1((_dafny.Zero).Minus((Companion_Default___.Fm59(globalState)).Cardinality()), 5)
		(_nw321).ArraySet1(_1904_v0, 6)
		(_nw321).ArraySet1(_1904_v0, 7)
		(_nw321).ArraySet1(_1904_v0, 8)
		(_nw321).ArraySet1(_dafny.IntOfInt64(158), 9)
		(_nw321).ArraySet1(_1904_v0, 10)
		(_nw321).ArraySet1(_1904_v0, 11)
		(_nw321).ArraySet1(_1904_v0, 12)
		(_nw321).ArraySet1(Companion_Default___.Fm3(_1905_v1, _1905_v1, globalState), 13)
		(_nw321).ArraySet1((_dafny.MultiSetFromSeq(_1907_v3)).Cardinality(), 14)
		(_nw321).ArraySet1(_dafny.IntOfInt64(-478), 15)
		(_nw321).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm3(_1905_v1, _1905_v1, globalState), (func() _dafny.Int {
			if (_1908_v4).Contains(_1904_v0) {
				return (_1908_v4).Multiplicity(_1904_v0)
			}
			return (_dafny.Zero).Minus(_1904_v0)
		})()), 16)
		(_nw321).ArraySet1(((_1908_v4).Union((_1908_v4).Update(_1904_v0, Companion_Default___.Abs(_1904_v0)))).Cardinality(), 17)
		(_nw321).ArraySet1(_1904_v0, 18)
		(_nw321).ArraySet1((Companion_Default___.Fm49(globalState)).Cardinality(), 19)
		(_nw321).ArraySet1((Companion_Default___.Fm57((_1909_v5).Cardinality(), (_dafny.Zero).Minus(_1904_v0), globalState)).Cardinality(), 20)
		(_nw321).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1912_v8).Cardinality())), 21)
		(_nw321).ArraySet1((_1904_v0).Minus(_dafny.IntOfUint32((_1913_v9).Cardinality())), 22)
		(_nw321).ArraySet1(_1904_v0, 23)
		_1914_v10 = _nw321
		for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1914_v10), 0))); ; {
			_guard_loop_3, _ok68 := _iter68()
			if !_ok68 {
				break
			}
			var _1915_i0 _dafny.Int
			_1915_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1915_i0).Sign() != -1) && ((_1915_i0).Cmp(_dafny.ArrayLen((_1914_v10), 0)) < 0)) {
				(_1914_v10).ArraySet1(Companion_Default___.SafeModuloInt(_1915_i0, (func() _dafny.Int {
					if (_1909_v5).Contains(_1904_v0) {
						return (_1909_v5).Get(_1904_v0).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.SeqOf(Companion_D15_.Create_DC29_(Companion_D15_.Create_DC27_(_dafny.MultiSetOf(_1905_v1, false))))).Cardinality())
				})()), (_1915_i0).Int())
			}
		}
		_1905_v1 = (_1904_v0).Cmp(_1904_v0) <= 0
		var _1916_i1 _dafny.Int
		_ = _1916_i1
		_1916_i1 = _dafny.Zero
		{
			for !(Companion_Default___.Fm0((_this).Fm6(_1913_v9, globalState), globalState)) {
				{
					if (_1916_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L18
					}
					_1916_i1 = (_1916_i1).Plus(_dafny.One)
					var _1917_v11 _dafny.Map
					_ = _1917_v11
					_1917_v11 = _1910_v6
					var _source25 _dafny.Map = _1917_v11
					_ = _source25
					var _1918___mcc_h0 _dafny.Map = _source25
					_ = _1918___mcc_h0
					var _1919_cf12 _dafny.Map = _1918___mcc_h0
					_ = _1919_cf12
					(globalState).F1 = _1904_v0
					var _1920_v12 _dafny.CodePoint
					_ = _1920_v12
					_1920_v12 = _dafny.CodePoint('c')
					var _1921_v13 D14
					_ = _1921_v13
					_1921_v13 = Companion_D14_.Create_DC26_(true, _1904_v0, _1920_v12)
					var _1922_v14 _dafny.MultiSet
					_ = _1922_v14
					_1922_v14 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1905_v1, _dafny.IntOfInt64(-334)))
					var _1923_v15 _dafny.Array
					_ = _1923_v15
					var _nwElement0_71 D14 = Companion_D14_.Create_DC26_(_1905_v1, _1904_v0, _1920_v12)
					_ = _nwElement0_71
					var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(28))
					_ = _nw322
					(_nw322).ArraySet1(_nwElement0_71, 0)
					(_nw322).ArraySet1(_1921_v13, 1)
					(_nw322).ArraySet1(_1921_v13, 2)
					(_nw322).ArraySet1(_1921_v13, 3)
					(_nw322).ArraySet1(_1921_v13, 4)
					(_nw322).ArraySet1(_1921_v13, 5)
					(_nw322).ArraySet1(_1921_v13, 6)
					(_nw322).ArraySet1(_1921_v13, 7)
					(_nw322).ArraySet1(_1921_v13, 8)
					(_nw322).ArraySet1(_1921_v13, 9)
					(_nw322).ArraySet1(_1921_v13, 10)
					(_nw322).ArraySet1(_1921_v13, 11)
					(_nw322).ArraySet1(_1921_v13, 12)
					(_nw322).ArraySet1(_1921_v13, 13)
					(_nw322).ArraySet1(_1921_v13, 14)
					(_nw322).ArraySet1(_1921_v13, 15)
					(_nw322).ArraySet1(Companion_D14_.Create_DC26_(_1905_v1, _dafny.IntOfInt64(538), _1920_v12), 16)
					(_nw322).ArraySet1(Companion_D14_.Create_DC26_((_1907_v3).Select((Companion_Default___.SafeIndex((_1922_v14).Cardinality(), _dafny.IntOfUint32((_1907_v3).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(245), _1920_v12), 17)
					(_nw322).ArraySet1(_1921_v13, 18)
					(_nw322).ArraySet1(_1921_v13, 19)
					(_nw322).ArraySet1(_1921_v13, 20)
					(_nw322).ArraySet1(_1921_v13, 21)
					(_nw322).ArraySet1(_1921_v13, 22)
					(_nw322).ArraySet1(_1921_v13, 23)
					(_nw322).ArraySet1(_1921_v13, 24)
					(_nw322).ArraySet1(Companion_Default___.Fm60(globalState), 25)
					(_nw322).ArraySet1((func() D14 {
						if _1905_v1 {
							return _1921_v13
						}
						return _1921_v13
					})(), 26)
					(_nw322).ArraySet1(Companion_Default___.Fm60(globalState), 27)
					_1923_v15 = _nw322
					_1923_v15 = _1923_v15
					var _1924_v16 _dafny.Array
					_ = _1924_v16
					var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(10))
					_ = _nw323
					_1924_v16 = _nw323
					var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1924_v16), 0))
					_ = _index335
					(_1924_v16).ArraySet1((_1909_v5).Merge(_1909_v5), (_index335).Int())
					var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_1924_v16), 0))
					_ = _index336
					(_1924_v16).ArraySet1(_1909_v5, (_index336).Int())
					var _1925_v17 T1
					_ = _1925_v17
					var _nw324 *C3 = New_C3_()
					_ = _nw324
					_nw324.Ctor__(_1905_v1, _1905_v1)
					_1925_v17 = _nw324
					var _1926_v18 _dafny.Map
					_ = _1926_v18
					_1926_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1925_v17, _1905_v1)
					_1905_v1 = (func() bool {
						if (_1926_v18).Contains(_1925_v17) {
							return (_1926_v18).Get(_1925_v17).(bool)
						}
						return _1905_v1
					})()
					_1905_v1 = _1905_v1
					var _1927_v19 _dafny.Sequence
					_ = _1927_v19
					_1927_v19 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-648))).Uint32(), func(coer104 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg104 _dafny.Int) interface{} {
							return coer104(arg104)
						}
					}((func(_1928_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1929_i2 _dafny.Int) _dafny.Int {
							return _1928_v0
						}
					})(_1904_v0))), _1913_v9)
					var _1930_v20 *C0
					_ = _1930_v20
					var _nw325 *C0 = New_C0_()
					_ = _nw325
					_nw325.Ctor__(_1905_v1, _1927_v19)
					_1930_v20 = _nw325
					var _1931_v21 *C1
					_ = _1931_v21
					var _nw326 *C1 = New_C1_()
					_ = _nw326
					_nw326.Ctor__()
					_1931_v21 = _nw326
					goto C18
				}
			C18:
			}
			goto L18
		}
	L18:
		var _1932_i3 _dafny.Int
		_ = _1932_i3
		_1932_i3 = _dafny.Zero
		{
			for (Companion_Default___.SafeModuloInt(_1904_v0, _1904_v0)).Cmp(_1904_v0) <= 0 {
				{
					if (_1932_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L19
					}
					_1932_i3 = (_1932_i3).Plus(_dafny.One)
					if _1905_v1 {
						var _1933_v22 _dafny.Map
						_ = _1933_v22
						_1933_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1905_v1, !(_1905_v1))
						_1933_v22 = (_1933_v22).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_1913_v9, _1913_v9), _1905_v1)
						(globalState).F1 = (_1904_v0).Minus((_1904_v0).Times(_1904_v0))
						var _1934_v23 _dafny.CodePoint
						_ = _1934_v23
						_1934_v23 = _dafny.CodePoint('p')
						var _rhs361 _dafny.Int = (_dafny.Zero).Minus(_1904_v0)
						_ = _rhs361
						var _rhs362 bool = _1905_v1
						_ = _rhs362
						var _rhs363 _dafny.Int = _1904_v0
						_ = _rhs363
						var _rhs364 _dafny.CodePoint = _1934_v23
						_ = _rhs364
						var _lhs249 *GlobalState = globalState
						_ = _lhs249
						var _lhs250 *GlobalState = globalState
						_ = _lhs250
						_lhs249.F1 = _rhs361
						_1905_v1 = _rhs362
						_lhs250.F1 = _rhs363
						_1934_v23 = _rhs364
						_1904_v0 = _1904_v0
						var _1935_v24 _dafny.Array
						_ = _1935_v24
						var _nwElement0_72 bool = _1905_v1
						_ = _nwElement0_72
						var _nw327 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.One)
						_ = _nw327
						(_nw327).ArraySet1(_nwElement0_72, 0)
						_1935_v24 = _nw327
						var _1936_v25 _dafny.Sequence
						_ = _1936_v25
						_1936_v25 = _dafny.SeqOf(_1935_v24)
						var _1937_v26 _dafny.Sequence
						_ = _1937_v26
						_1937_v26 = _dafny.SeqOf((_1936_v25).Select((Companion_Default___.SafeIndex(_1904_v0, _dafny.IntOfUint32((_1936_v25).Cardinality()))).Uint32()).(_dafny.Array), _1935_v24, _1935_v24)
						_1935_v24 = (_1937_v26).Select((Companion_Default___.SafeIndex(_1904_v0, _dafny.IntOfUint32((_1937_v26).Cardinality()))).Uint32()).(_dafny.Array)
					} else {
						(globalState).F1 = _1904_v0
						var _1938_v27 _dafny.Map
						_ = _1938_v27
						_1938_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(p0, p0), (_this).Fm29(_1904_v0, p0, globalState))
						_1938_v27 = (_1938_v27).Update(_dafny.UnicodeSeqOfUtf8Bytes("eqjlgq"), (func() _dafny.Int {
							if (_1909_v5).Contains(_1904_v0) {
								return (_1909_v5).Get(_1904_v0).(_dafny.Int)
							}
							return _1904_v0
						})())
						(globalState).F1 = _1904_v0
						var _1939_v28 *C1
						_ = _1939_v28
						var _nw328 *C1 = New_C1_()
						_ = _nw328
						_nw328.Ctor__()
						_1939_v28 = _nw328
						_1908_v4 = Companion_Default___.Fm32(globalState)
					}
					var _1940_v29 _dafny.Array
					_ = _1940_v29
					var _nw329 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
					_ = _nw329
					_1940_v29 = _nw329
					var _1941_v30 _dafny.Array
					_ = _1941_v30
					_1941_v30 = _1940_v29
					var _1942_v31 _dafny.Map
					_ = _1942_v31
					_1942_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1941_v30, !(_1905_v1))
					_1942_v31 = (_1942_v31).Update(_1940_v29, false)
					_1905_v1 = _1905_v1
					var _1943_v32 *C3
					_ = _1943_v32
					var _nw330 *C3 = New_C3_()
					_ = _nw330
					_nw330.Ctor__(_1905_v1, _1905_v1)
					_1943_v32 = _nw330
					goto C19
				}
			C19:
			}
			goto L19
		}
	L19:
		var _1944_v33 _dafny.CodePoint
		_ = _1944_v33
		_1944_v33 = _dafny.CodePoint('v')
		var _1945_v34 *C5
		_ = _1945_v34
		var _nw331 *C5 = New_C5_()
		_ = _nw331
		_nw331.Ctor__(_1904_v0, _1944_v33, !((_1905_v1) || (_1905_v1)))
		_1945_v34 = _nw331
		var _1946_v35 _dafny.Map
		_ = _1946_v35
		_1946_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1914_v10, _1907_v3)
		_1946_v35 = (_1946_v35).Update(_1914_v10, _dafny.Companion_Sequence_.Concatenate(_1907_v3, _1907_v3))
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	_f5 _dafny.CodePoint
	_f6 bool
	F10 _dafny.Int
	_f9 _dafny.Int
}

func New_C7_() *C7 {
	_this := C7{}

	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
	_this.F10 = _dafny.Zero
	_this._f9 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C7{}
var _ T2 = &C7{}
var _ T1 = &C7{}
var _ T0 = &C7{}
var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) F5() _dafny.CodePoint {
	return _this._f5
}
func (_this *C7) F6() bool {
	return _this._f6
}
func (_this *C7) Ctor__(f9 _dafny.Int, f10 _dafny.Int, f5 _dafny.CodePoint, f6 bool) {
	{
		(_this)._f9 = f9
		(_this).F10 = f10
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C7) Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt((_this).F9(), _this.F10)
	}
}
func (_this *C7) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((_this).F6())
	}
}
func (_this *C7) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F9()
	}
}
func (_this *C7) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return _this.F10
	}
}
func (_this *C7) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		if (_this).F6() {
			return Companion_D1_.Create_DC4_((_this).F9())
		} else {
			return Companion_D1_.Create_DC4_(_this.F10)
		}
	}
}
func (_this *C7) Fm13(globalState *GlobalState) bool {
	{
		return ((_this).F9()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mhlxx")).Cardinality())) == 0
	}
}
func (_this *C7) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _1947_v0 _dafny.Array
		_ = _1947_v0
		var _len0_52 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_52
		var _nw332 _dafny.Array
		_ = _nw332
		if _len0_52.Cmp(_dafny.Zero) == 0 {
			_nw332 = _dafny.NewArray(_len0_52)
		} else {
			var _init52 func(_dafny.Int) bool = (func(_1948_p2 _dafny.Sequence) func(_dafny.Int) bool {
				return func(_1949_i0 _dafny.Int) bool {
					return !_dafny.Companion_Sequence_.Contains(_1948_p2, (_this).F5())
				}
			})(p2)
			_ = _init52
			var _element0_52 = _init52(_dafny.Zero)
			_ = _element0_52
			_nw332 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
			(_nw332).ArraySet1(_element0_52, 0)
			var _nativeLen0_52 = (_len0_52).Int()
			_ = _nativeLen0_52
			for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
				(_nw332).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
			}
		}
		_1947_v0 = _nw332
		var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))
		_ = _index337
		(_1947_v0).ArraySet1((func() bool {
			if !((_this).F6()) {
				return !(p0)
			}
			return (_this).F6()
		})(), (_index337).Int())
		var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))
		_ = _index338
		(_1947_v0).ArraySet1(p0, (_index338).Int())
		var _1950_v1 _dafny.MultiSet
		_ = _1950_v1
		_1950_v1 = _dafny.MultiSetOf(p0)
		(_this).F10 = (_this.F10).Minus(Companion_Default___.SafeModuloInt((_1950_v1).Cardinality(), _this.F10))
		var _1951_v2 D2
		_ = _1951_v2
		_1951_v2 = Companion_D2_.Create_DC6_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(563))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg105 _dafny.Int) interface{} {
				return coer105(arg105)
			}
		}(func(_1952_i1 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})))
		var _source26 D2 = _1951_v2
		_ = _source26
		if _source26.Is_DC7() {
			var _1953___mcc_h0 _dafny.Int = _source26.Get_().(D2_DC7).Cf7
			_ = _1953___mcc_h0
			var _1954___mcc_h1 _dafny.Int = _source26.Get_().(D2_DC7).Cf8
			_ = _1954___mcc_h1
			var _1955_cf8 _dafny.Int = _1954___mcc_h1
			_ = _1955_cf8
			var _1956_cf7 _dafny.Int = _1953___mcc_h0
			_ = _1956_cf7
			var _1957_v3 _dafny.CodePoint
			_ = _1957_v3
			_1957_v3 = _dafny.CodePoint('l')
			_1957_v3 = (_this).F5()
			var _1958_v4 _dafny.Sequence
			_ = _1958_v4
			_1958_v4 = _dafny.SeqOf(_this.F10, _dafny.IntOfInt64(929), _1956_cf7, (_dafny.Zero).Minus((_this).F9()))
			(globalState).F1 = _dafny.IntOfUint32((_1958_v4).Cardinality())
			var _1959_v5 _dafny.Sequence
			_ = _1959_v5
			_1959_v5 = _dafny.SeqOf((_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), (_this).F6())
			_1959_v5 = Companion_Default___.Fm2((_1959_v5).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf(p1, _this.F10)).Cardinality())), _dafny.IntOfUint32((_1959_v5).Cardinality()))).Uint32()).(bool), (_1959_v5).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1959_v5).Cardinality()))).Uint32()).(bool), globalState)
			r1 = p2
		} else if _source26.Is_DC8() {
			var _1960___mcc_h2 bool = _source26.Get_().(D2_DC8).Cf9
			_ = _1960___mcc_h2
			var _1961_cf9 bool = _1960___mcc_h2
			_ = _1961_cf9
			var _1962_v6 _dafny.Set
			_ = _1962_v6
			_1962_v6 = _dafny.SetOf(true, (_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), p0, _1961_cf9)
			var _1963_v7 _dafny.Sequence
			_ = _1963_v7
			_1963_v7 = _dafny.SeqOf(false)
			var _1964_v8 _dafny.Map
			_ = _1964_v8
			_1964_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1963_v7, p1)
			var _1965_v9 _dafny.Sequence
			_ = _1965_v9
			_1965_v9 = _dafny.SeqOf(_1963_v7)
			var _1966_v10 _dafny.Map
			_ = _1966_v10
			_1966_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1961_cf9), true)
			var _1967_v11 _dafny.Sequence
			_ = _1967_v11
			_1967_v11 = _dafny.SeqOf(p1, _dafny.IntOfUint32((p2).Cardinality()), p1, p1)
			var _1968_v12 _dafny.Array
			_ = _1968_v12
			var _nwElement0_73 _dafny.Int = (_1962_v6).Cardinality()
			_ = _nwElement0_73
			var _nw333 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(12))
			_ = _nw333
			(_nw333).ArraySet1(_nwElement0_73, 0)
			(_nw333).ArraySet1((_this.F10).Plus((func() _dafny.Int {
				if (_1964_v8).Contains((_1965_v9).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_1965_v9).Cardinality()))).Uint32()).(_dafny.Sequence)) {
					return (_1964_v8).Get((_1965_v9).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_1965_v9).Cardinality()))).Uint32()).(_dafny.Sequence)).(_dafny.Int)
				}
				return _this.F10
			})()), 1)
			(_nw333).ArraySet1(_this.F10, 2)
			(_nw333).ArraySet1(_this.F10, 3)
			(_nw333).ArraySet1(_this.F10, 4)
			(_nw333).ArraySet1(_this.F10, 5)
			(_nw333).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 6)
			(_nw333).ArraySet1(_dafny.IntOfInt64(534), 7)
			(_nw333).ArraySet1(Companion_Default___.SafeDivisionInt((_1966_v10).Cardinality(), _dafny.IntOfUint32((_1967_v11).Cardinality())), 8)
			(_nw333).ArraySet1(p1, 9)
			(_nw333).ArraySet1(_dafny.IntOfInt64(599), 10)
			(_nw333).ArraySet1((_this).Fm7((_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), globalState), 11)
			_1968_v12 = _nw333
			var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1968_v12), 0))
			_ = _index339
			(_1968_v12).ArraySet1(p1, (_index339).Int())
			var _1969_v13 _dafny.Map
			_ = _1969_v13
			_1969_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Sequence {
				if (_this).F6() {
					return (_1951_v2).Dtor_cf6()
				}
				return p2
			})(), !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F6())).Contains(p1))
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1968_v12), 0))
			_ = _index340
			var _rhs365 _dafny.Int = (_this).F9()
			_ = _rhs365
			var _rhs366 _dafny.Map = _1969_v13
			_ = _rhs366
			var _rhs367 _dafny.Int = (Companion_Default___.SafeModuloInt((_this).F9(), (_dafny.Zero).Minus((_this).F9()))).Times((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(953), _dafny.IntOfInt64(-194))))
			_ = _rhs367
			var _lhs251 _dafny.Array = _1968_v12
			_ = _lhs251
			var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1968_v12), 0))
			_ = _lhs252
			var _lhs253 *C7 = _this
			_ = _lhs253
			(_lhs251).ArraySet1(_rhs365, (_lhs252).Int())
			_1969_v13 = _rhs366
			_lhs253.F10 = _rhs367
			_1963_v7 = _1963_v7
			var _1970_v14 *C0
			_ = _1970_v14
			var _nw334 *C0 = New_C0_()
			_ = _nw334
			_nw334.Ctor__((_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(794))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}((func(_1971_v11 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1972_i2 _dafny.Int) _dafny.Sequence {
					return _1971_v11
				}
			})(_1967_v11))))
			_1970_v14 = _nw334
			var _1973_v15 *C0
			_ = _1973_v15
			var _nw335 *C0 = New_C0_()
			_ = _nw335
			_nw335.Ctor__(p0, _1970_v14.F12)
			_1973_v15 = _nw335
		} else if _source26.Is_DC6() {
			var _1974___mcc_h3 _dafny.Sequence = _source26.Get_().(D2_DC6).Cf6
			_ = _1974___mcc_h3
			var _1975_cf6 _dafny.Sequence = _1974___mcc_h3
			_ = _1975_cf6
			(globalState).F1 = _this.F10
			(globalState).F1 = _this.F10
			(_this).F10 = (_this).F9()
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))
			_ = _index341
			(_1947_v0).ArraySet1(!(((_this).F9()).Cmp(Companion_Default___.Fm3((_this).F6(), p0, globalState)) < 0), (_index341).Int())
		} else {
			var _1976___mcc_h4 D2 = _source26.Get_().(D2_DC9).Cf10
			_ = _1976___mcc_h4
			var _1977_cf10 D2 = _1976___mcc_h4
			_ = _1977_cf10
			var _1978_v16 _dafny.Sequence
			_ = _1978_v16
			_1978_v16 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cofnqf"), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(932))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg107 _dafny.Int) interface{} {
					return coer107(arg107)
				}
			}(func(_1979_i3 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			})), p2))
			_1978_v16 = _1978_v16
			var _1980_v17 T0
			_ = _1980_v17
			var _nw336 *C0 = New_C0_()
			_ = _nw336
			_nw336.Ctor__((_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(748))).Uint32(), func(coer108 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}(func(_1981_i4 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqOf((_dafny.SetOf((_this).F6(), (_this).F6())).Cardinality())
			})))
			_1980_v17 = _nw336
			var _1982_v18 _dafny.Sequence
			_ = _1982_v18
			_1982_v18 = _dafny.SeqOf(p1)
			var _1983_v19 _dafny.Map
			_ = _1983_v19
			_1983_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1980_v17, _1982_v18)
			var _1984_v20 D2
			_ = _1984_v20
			_1984_v20 = Companion_D2_.Create_DC8_(p0)
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))
			_ = _index342
			var _rhs368 bool = ((_this.F10).Times(_dafny.IntOfInt64(611))).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6())).Cardinality())) <= 0
			_ = _rhs368
			var _rhs369 bool = ((_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool)) || ((_1984_v20).Dtor_cf9())
			_ = _rhs369
			var _rhs370 _dafny.Map = ((_1983_v19).Update(_1980_v17, _1982_v18)).Merge((_1983_v19).Update(_1980_v17, _dafny.SeqOf((_1950_v1).Cardinality())))
			_ = _rhs370
			var _lhs254 _dafny.Array = _1947_v0
			_ = _lhs254
			var _lhs255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))
			_ = _lhs255
			r0 = _rhs368
			(_lhs254).ArraySet1(_rhs369, (_lhs255).Int())
			_1983_v19 = _rhs370
			var _1985_v21 _dafny.Map
			_ = _1985_v21
			_1985_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _this.F10)
			var _1986_v22 _dafny.Map
			_ = _1986_v22
			_1986_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1985_v21, (func() _dafny.Int {
				if !(p0) {
					return p1
				}
				return _this.F10
			})())
			_1986_v22 = (_1986_v22).Update(_1985_v21, (_this.F10).Plus(_dafny.IntOfUint32((_1982_v18).Cardinality())))
			var _1987_v23 _dafny.Sequence
			_ = _1987_v23
			_1987_v23 = _dafny.SeqOf(p0)
			var _1988_v24 _dafny.Array
			_ = _1988_v24
			var _nw337 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw337
			_1988_v24 = _nw337
			var _1989_v25 _dafny.Array
			_ = _1989_v25
			_1989_v25 = _1988_v24
			var _rhs371 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1987_v23, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p0), _1987_v23))
			_ = _rhs371
			var _rhs372 D2 = (func() D2 {
				if false {
					return _1951_v2
				}
				return _1951_v2
			})()
			_ = _rhs372
			var _rhs373 _dafny.Array = ((func() _dafny.Array {
				if !(p0) {
					return _1989_v25
				}
				return _1989_v25
			})())
			_ = _rhs373
			_1987_v23 = _rhs371
			_1951_v2 = _rhs372
			_1988_v24 = _rhs373
		}
		var _1990_v26 _dafny.Sequence
		_ = _1990_v26
		_1990_v26 = _dafny.SeqOf(p0)
		var _1991_v27 _dafny.Sequence
		_ = _1991_v27
		_1991_v27 = _dafny.SeqOf(Companion_Default___.Fm3(p0, (_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool), globalState))
		var _1992_v28 _dafny.Map
		_ = _1992_v28
		_1992_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1990_v26, (_this).Fm6(_1991_v27, globalState))
		_1992_v28 = (_1992_v28).Update(_1990_v26, p0)
		var _1993_v29 _dafny.Sequence
		_ = _1993_v29
		_1993_v29 = _dafny.SeqOf(p2, p2)
		(_this).F10 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((_1993_v29).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_1993_v29).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), p1)
		var _1994_v30 *C1
		_ = _1994_v30
		var _nw338 *C1 = New_C1_()
		_ = _nw338
		_nw338.Ctor__()
		_1994_v30 = _nw338
		r0 = (func() bool {
			if false {
				return (_this).F6()
			}
			return (_1947_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(292), _dafny.ArrayLen((_1947_v0), 0))).Int()).(bool)
		})()
		r1 = p2
		return r0, r1
	}
}
func (_this *C7) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1995_v0 _dafny.Map
		_ = _1995_v0
		_1995_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), !((_this).F6()))
		var _1996_v1 _dafny.Map
		_ = _1996_v1
		_1996_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfInt64(-162))
		var _1997_v2 _dafny.Sequence
		_ = _1997_v2
		_1997_v2 = _dafny.SeqOf(_1995_v0, _1995_v0, _1995_v0)
		var _1998_v3 _dafny.Array
		_ = _1998_v3
		var _nwElement0_74 _dafny.Map = _1995_v0
		_ = _nwElement0_74
		var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(14))
		_ = _nw339
		(_nw339).ArraySet1(_nwElement0_74, 0)
		(_nw339).ArraySet1(_1995_v0, 1)
		(_nw339).ArraySet1((_1995_v0).Merge(_1995_v0), 2)
		(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6()), 3)
		(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_this).F6())), (_this).F6()), 4)
		(_nw339).ArraySet1((_1995_v0).Update((_this).F6(), (_this).F6()), 5)
		(_nw339).ArraySet1((_1995_v0).Merge(_1995_v0), 6)
		(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6()), 7)
		(_nw339).ArraySet1(Companion_Default___.Fm19(Companion_Default___.Fm0((_this).F6(), globalState), (func() _dafny.Int {
			if (_1996_v1).Contains((_this).F9()) {
				return (_1996_v1).Get((_this).F9()).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nud")).Cardinality())
		})(), globalState), 8)
		(_nw339).ArraySet1(_1995_v0, 9)
		(_nw339).ArraySet1(_1995_v0, 10)
		(_nw339).ArraySet1((_1995_v0).Merge(_1995_v0), 11)
		(_nw339).ArraySet1(_1995_v0, 12)
		(_nw339).ArraySet1((_1995_v0).Merge((_1997_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.IntOfUint32((_1997_v2).Cardinality()))).Uint32()).(_dafny.Map)), 13)
		_1998_v3 = _nw339
		for _iter69 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1998_v3), 0))); ; {
			_guard_loop_4, _ok69 := _iter69()
			if !_ok69 {
				break
			}
			var _1999_i0 _dafny.Int
			_1999_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1999_i0).Sign() != -1) && ((_1999_i0).Cmp(_dafny.ArrayLen((_1998_v3), 0)) < 0)) {
				(_1998_v3).ArraySet1(_1995_v0, (_1999_i0).Int())
			}
		}
		var _2000_i1 _dafny.Int
		_ = _2000_i1
		_2000_i1 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_2000_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L20
					}
					_2000_i1 = (_2000_i1).Plus(_dafny.One)
					var _2001_v4 bool
					_ = _2001_v4
					_2001_v4 = false
					_2001_v4 = (func() bool {
						if (_1995_v0).Contains((_this).F6()) {
							return (_1995_v0).Get((_this).F6()).(bool)
						}
						return ((_this).F6()) && (_2001_v4)
					})()
					(globalState).F1 = (_this).F9()
					var _2002_v5 _dafny.Map
					_ = _2002_v5
					_2002_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F5())
					var _2003_v6 _dafny.Sequence
					_ = _2003_v6
					_2003_v6 = _dafny.SeqOf(_2001_v4)
					var _2004_v7 _dafny.Set
					_ = _2004_v7
					_2004_v7 = _dafny.SetOf(_dafny.IntOfUint32((_2003_v6).Cardinality()), (_dafny.MultiSetOf(_2001_v4, (_this).F6())).Cardinality())
					var _2005_v8 _dafny.Map
					_ = _2005_v8
					_2005_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(133), Companion_D1_.Create_DC4_((_this).F9()))
					var _2006_v9 _dafny.MultiSet
					_ = _2006_v9
					_2006_v9 = _dafny.MultiSetOf(_2004_v7, _dafny.SetOf((_2005_v8).Cardinality()))
					var _2007_v10 _dafny.Array
					_ = _2007_v10
					var _nw340 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
					_ = _nw340
					_2007_v10 = _nw340
					var _2008_v11 _dafny.MultiSet
					_ = _2008_v11
					_2008_v11 = _dafny.MultiSetOf(_2007_v10, _2007_v10, _2007_v10, _2007_v10, _2007_v10)
					_2002_v5 = (_2002_v5).Update(((_2006_v9).Cardinality()).Cmp((_2008_v11).Cardinality()) <= 0, (_this).F5())
					var _2009_v12 _dafny.Sequence
					_ = _2009_v12
					_2009_v12 = _dafny.SeqOf(_dafny.IntOfInt64(445), _this.F10)
					var _2010_v13 _dafny.Sequence
					_ = _2010_v13
					_2010_v13 = _dafny.SeqOf(_2009_v12, _2009_v12)
					var _2011_v14 *C0
					_ = _2011_v14
					var _nw341 *C0 = New_C0_()
					_ = _nw341
					_nw341.Ctor__(Companion_Default___.Fm0((_this).F6(), globalState), _dafny.Companion_Sequence_.Concatenate(_2010_v13, _2010_v13))
					_2011_v14 = _nw341
					goto C20
				}
			C20:
			}
			goto L20
		}
	L20:
		if (_this).F6() {
			(globalState).F1 = _this.F10
			var _2012_v15 _dafny.Array
			_ = _2012_v15
			var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw342
			_2012_v15 = _nw342
			var _2013_v16 _dafny.Map
			_ = _2013_v16
			_2013_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _2012_v15)
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
			_ = _index343
			(_2012_v15).ArraySet1(((_this).F9()).Times((_this).F9()), (_index343).Int())
			var _2014_v17 bool
			_ = _2014_v17
			_2014_v17 = false
			var _2015_v18 _dafny.Map
			_ = _2015_v18
			_2015_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), false)
			var _2016_v19 _dafny.Map
			_ = _2016_v19
			_2016_v19 = _2015_v18
			var _2017_v20 _dafny.Map
			_ = _2017_v20
			_2017_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2016_v19, _dafny.CodePoint('j'))
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
			_ = _index344
			var _rhs374 _dafny.Map = _2013_v16
			_ = _rhs374
			var _rhs375 _dafny.Int = _this.F10
			_ = _rhs375
			var _rhs376 bool = false
			_ = _rhs376
			var _rhs377 _dafny.Int = (_this).Fm8(!(_2017_v20).Equals(_2017_v20), (_this).F6(), false, globalState)
			_ = _rhs377
			var _lhs256 _dafny.Array = _2012_v15
			_ = _lhs256
			var _lhs257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
			_ = _lhs257
			var _lhs258 *C7 = _this
			_ = _lhs258
			_2013_v16 = _rhs374
			(_lhs256).ArraySet1(_rhs375, (_lhs257).Int())
			_2014_v17 = _rhs376
			_lhs258.F10 = _rhs377
			var _2018_v21 _dafny.Sequence
			_ = _2018_v21
			_2018_v21 = _dafny.SeqOf((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int), _this.F10)
			var _2019_v22 *C0
			_ = _2019_v22
			var _nw343 *C0 = New_C0_()
			_ = _nw343
			_nw343.Ctor__(_2014_v17, _dafny.SeqOf(_2018_v21))
			_2019_v22 = _nw343
			_2019_v22 = _2019_v22
			if ((_this).F9()).Cmp(Companion_Default___.SafeDivisionInt((_this).F9(), (_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int))) < 0 {
				var _2020_v23 _dafny.CodePoint
				_ = _2020_v23
				_2020_v23 = _dafny.CodePoint('a')
				_2020_v23 = _2020_v23
				var _2021_v24 _dafny.Sequence
				_ = _2021_v24
				_2021_v24 = _dafny.UnicodeSeqOfUtf8Bytes("mcwnadjh")
				var _2022_v25 _dafny.Set
				_ = _2022_v25
				_2022_v25 = _dafny.SetOf(_this.F10)
				var _2023_v26 _dafny.Map
				_ = _2023_v26
				_2023_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int), (_this).F5())
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
				_ = _index345
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
				_ = _index346
				var _rhs378 bool = (_2019_v22).F11()
				_ = _rhs378
				var _rhs379 _dafny.Sequence = (func() _dafny.Sequence {
					if _2014_v17 {
						return _dafny.Companion_Sequence_.Concatenate(_2021_v24, _2021_v24)
					}
					return _2021_v24
				})()
				_ = _rhs379
				var _rhs380 bool = ((_2022_v25).Difference(_2022_v25)).IsSubsetOf(_2022_v25)
				_ = _rhs380
				var _rhs381 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1996_v1).Cardinality(), (_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)))
				_ = _rhs381
				var _rhs382 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2021_v24, _2021_v24), (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2021_v24, _2021_v24)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_2023_v26).Contains((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)) {
						return (_2023_v26).Get((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)).(_dafny.CodePoint)
					}
					return (_this).F5()
				})())).Cardinality())), (_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int))
				_ = _rhs382
				var _lhs259 _dafny.Array = _2012_v15
				_ = _lhs259
				var _lhs260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
				_ = _lhs260
				var _lhs261 _dafny.Array = _2012_v15
				_ = _lhs261
				var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))
				_ = _lhs262
				_2014_v17 = _rhs378
				_2021_v24 = _rhs379
				_2014_v17 = _rhs380
				(_lhs259).ArraySet1(_rhs381, (_lhs260).Int())
				(_lhs261).ArraySet1(_rhs382, (_lhs262).Int())
				var _2024_v27 D2
				_ = _2024_v27
				_2024_v27 = Companion_D2_.Create_DC6_(_2021_v24)
				_2024_v27 = Companion_D2_.Create_DC6_(_2021_v24)
				var _2025_v28 _dafny.Sequence
				_ = _2025_v28
				_2025_v28 = _dafny.SeqOf((_2019_v22).F11())
				var _2026_v29 _dafny.Map
				_ = _2026_v29
				_2026_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2025_v28, _2025_v28), _2014_v17)
				_2026_v29 = (_2026_v29).Merge(_2026_v29)
				var _2027_v30 _dafny.Set
				_ = _2027_v30
				_2027_v30 = _dafny.SetOf((_this).F6())
				(globalState).F1 = (_this.F10).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_2019_v22).F11())).Cardinality()).Minus((_2027_v30).Cardinality()))
			} else {
				r0 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((Companion_D1_.Create_DC4_((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int))).Dtor_cf4(), (_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_2019_v22.F12).Cardinality()))))
				(_this).F10 = (((func() _dafny.Map {
					if !((_this).F6()) {
						return _1995_v0
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2019_v22).F11(), (_2019_v22).F11())
				})()).Merge(_1995_v0)).Cardinality()
				var _2028_v31 _dafny.Sequence
				_ = _2028_v31
				_2028_v31 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int), !((_this).F6())))
				_2014_v17 = ((_this).F9()).Cmp(_dafny.IntOfUint32((_2028_v31).Cardinality())) == 0
				var _2029_v32 _dafny.Array
				_ = _2029_v32
				var _nw344 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw344
				_2029_v32 = _nw344
				var _2030_v33 _dafny.Map
				_ = _2030_v33
				_2030_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _2029_v32)
				var _2031_v34 _dafny.Map
				_ = _2031_v34
				_2031_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_this.F10), (_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int))
				var _rhs383 bool = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("latagnlxg"), _dafny.UnicodeSeqOfUtf8Bytes("avl"))
				_ = _rhs383
				var _rhs384 _dafny.Array = (func() _dafny.Array {
					if (_2030_v33).Contains((_2031_v34).Cardinality()) {
						return (_2030_v33).Get((_2031_v34).Cardinality()).(_dafny.Array)
					}
					return _2029_v32
				})()
				_ = _rhs384
				var _rhs385 _dafny.Int = (_dafny.Zero).Minus(_this.F10)
				_ = _rhs385
				_2014_v17 = _rhs383
				_2029_v32 = _rhs384
				r0 = _rhs385
				var _2032_v35 _dafny.Sequence
				_ = _2032_v35
				_2032_v35 = _dafny.UnicodeSeqOfUtf8Bytes("ywr")
				_2032_v35 = _dafny.Companion_Sequence_.Concatenate(_2032_v35, (func() _dafny.Sequence {
					if (_this).F6() {
						return _2032_v35
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(136))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg109 _dafny.Int) interface{} {
							return coer109(arg109)
						}
					}(func(_2033_i2 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					}))
				})())
			}
			var _2034_v36 D0
			_ = _2034_v36
			_2034_v36 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC1_())
			var _2035_v37 _dafny.Array
			_ = _2035_v37
			var _len0_53 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_53
			var _nw345 _dafny.Array
			_ = _nw345
			if _len0_53.Cmp(_dafny.Zero) == 0 {
				_nw345 = _dafny.NewArray(_len0_53)
			} else {
				var _init53 func(_dafny.Int) bool = func(_2036_i3 _dafny.Int) bool {
					return (_this).F6()
				}
				_ = _init53
				var _element0_53 = _init53(_dafny.Zero)
				_ = _element0_53
				_nw345 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
				(_nw345).ArraySet1(_element0_53, 0)
				var _nativeLen0_53 = (_len0_53).Int()
				_ = _nativeLen0_53
				for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
					(_nw345).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
				}
			}
			_2035_v37 = _nw345
			var _2037_v38 _dafny.Map
			_ = _2037_v38
			_2037_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2035_v37, (_this).F5())
			var _2038_v39 D0
			_ = _2038_v39
			_2038_v39 = Companion_D0_.Create_DC2_(_2037_v38, _dafny.CodePoint('e'))
			var _2039_v40 D0
			_ = _2039_v40
			_2039_v40 = Companion_D0_.Create_DC3_(_2038_v39)
			var _pat_let_tv28 = _2038_v39
			_ = _pat_let_tv28
			var _source27 D0 = func(_pat_let32_0 D0) D0 {
				return func(_2040_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let33_0 D0) D0 {
						return func(_2041_dt__update_hcf3_h0 D0) D0 {
							return Companion_D0_.Create_DC3_(_2041_dt__update_hcf3_h0)
						}(_pat_let33_0)
					}(_pat_let_tv28)
				}(_pat_let32_0)
			}(_2034_v36)
			_ = _source27
			if _source27.Is_DC1() {
				_2014_v17 = (_2019_v22).F11()
				var _2042_v41 D1
				_ = _2042_v41
				_2042_v41 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(-663))
				var _2043_v42 _dafny.Array
				_ = _2043_v42
				var _nwElement0_75 D1 = _2042_v41
				_ = _nwElement0_75
				var _nw346 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(7))
				_ = _nw346
				(_nw346).ArraySet1(_nwElement0_75, 0)
				(_nw346).ArraySet1(_2042_v41, 1)
				(_nw346).ArraySet1(Companion_D1_.Create_DC4_((_dafny.SetOf((_this).F6())).Cardinality()), 2)
				(_nw346).ArraySet1(_2042_v41, 3)
				(_nw346).ArraySet1(_2042_v41, 4)
				(_nw346).ArraySet1(_2042_v41, 5)
				(_nw346).ArraySet1(_2042_v41, 6)
				_2043_v42 = _nw346
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2043_v42), 0))
				_ = _index347
				(_2043_v42).ArraySet1(_2042_v41, (_index347).Int())
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2043_v42), 0))
				_ = _index348
				(_2043_v42).ArraySet1(_2042_v41, (_index348).Int())
				var _2044_v43 _dafny.Sequence
				_ = _2044_v43
				_2044_v43 = _dafny.SeqOf(_2012_v15, _2012_v15)
				var _2045_v44 _dafny.Sequence
				_ = _2045_v44
				_2045_v44 = _dafny.UnicodeSeqOfUtf8Bytes("qucb")
				var _rhs386 _dafny.Int = (_dafny.Zero).Minus((_2018_v21).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_2018_v21).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs386
				var _rhs387 _dafny.Int = ((_this).F9()).Minus((_dafny.Zero).Minus((_this.F10).Minus((_this).F9())))
				_ = _rhs387
				var _rhs388 _dafny.Sequence = _dafny.SeqOf(_2012_v15)
				_ = _rhs388
				var _rhs389 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("yhwkuxs")
				_ = _rhs389
				var _lhs263 *C7 = _this
				_ = _lhs263
				var _lhs264 *GlobalState = globalState
				_ = _lhs264
				_lhs263.F10 = _rhs386
				_lhs264.F1 = _rhs387
				_2044_v43 = _rhs388
				_2045_v44 = _rhs389
				var _2046_v45 _dafny.MultiSet
				_ = _2046_v45
				_2046_v45 = _dafny.MultiSetOf(Companion_Default___.Fm0((_2019_v22).F11(), globalState), (_2019_v22).F11())
				var _2047_v46 _dafny.Sequence
				_ = _2047_v46
				_2047_v46 = _dafny.SeqOf((_this).F6(), !(_2014_v17))
				var _2048_v47 _dafny.Array
				_ = _2048_v47
				var _nwElement0_76 _dafny.MultiSet = _dafny.MultiSetOf((_2019_v22).F11(), (_this).F6())
				_ = _nwElement0_76
				var _nw347 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(18))
				_ = _nw347
				(_nw347).ArraySet1(_nwElement0_76, 0)
				(_nw347).ArraySet1((_2046_v45).Union(_2046_v45), 1)
				(_nw347).ArraySet1((_dafny.MultiSetOf((_2019_v22).F11(), (_2019_v22).F11())).Difference(_2046_v45), 2)
				(_nw347).ArraySet1(_2046_v45, 3)
				(_nw347).ArraySet1((_2046_v45).Difference(_dafny.MultiSetFromSeq(_2047_v46)), 4)
				(_nw347).ArraySet1(_2046_v45, 5)
				(_nw347).ArraySet1(_dafny.MultiSetOf((_this).F6(), false, (_2019_v22).F11(), !((_this).F6()), (_2019_v22).F11()), 6)
				(_nw347).ArraySet1((_2046_v45).Union(_dafny.MultiSetFromSeq(_2047_v46)), 7)
				(_nw347).ArraySet1(_2046_v45, 8)
				(_nw347).ArraySet1(_2046_v45, 9)
				(_nw347).ArraySet1(_2046_v45, 10)
				(_nw347).ArraySet1(_2046_v45, 11)
				(_nw347).ArraySet1((_2046_v45).Union(_2046_v45), 12)
				(_nw347).ArraySet1(_2046_v45, 13)
				(_nw347).ArraySet1(_2046_v45, 14)
				(_nw347).ArraySet1(_2046_v45, 15)
				(_nw347).ArraySet1(_2046_v45, 16)
				(_nw347).ArraySet1((_dafny.MultiSetOf((_this).F6())).Union((_dafny.MultiSetOf((_this).F6(), (_2019_v22).F11(), (_this).F6())).Update((_2019_v22).F11(), Companion_Default___.Abs(((_2043_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_2043_v42), 0))).Int()).(D1)).Dtor_cf4()))), 17)
				_2048_v47 = _nw347
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_2048_v47), 0))
				_ = _index349
				(_2048_v47).ArraySet1((_2046_v45).Union(_2046_v45), (_index349).Int())
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_2048_v47), 0))
				_ = _index350
				(_2048_v47).ArraySet1(_2046_v45, (_index350).Int())
			} else if _source27.Is_DC2() {
				var _2049___mcc_h0 _dafny.Map = _source27.Get_().(D0_DC2).Cf1
				_ = _2049___mcc_h0
				var _2050___mcc_h1 _dafny.CodePoint = _source27.Get_().(D0_DC2).Cf2
				_ = _2050___mcc_h1
				var _2051_cf2 _dafny.CodePoint = _2050___mcc_h1
				_ = _2051_cf2
				var _2052_cf1 _dafny.Map = _2049___mcc_h0
				_ = _2052_cf1
				var _2053_v48 *C1
				_ = _2053_v48
				var _nw348 *C1 = New_C1_()
				_ = _nw348
				_nw348.Ctor__()
				_2053_v48 = _nw348
				_1995_v0 = (_1995_v0).Update(((_this).F9()).Cmp(_this.F10) > 0, (_2019_v22).F11())
				_2014_v17 = ((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeModuloInt((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int), (_this).F9())) != 0
				_2014_v17 = (_2019_v22).F11()
			} else if _source27.Is_DC0() {
				var _2054___mcc_h2 bool = _source27.Get_().(D0_DC0).Cf0
				_ = _2054___mcc_h2
				var _2055_cf0 bool = _2054___mcc_h2
				_ = _2055_cf0
				var _2056_v49 _dafny.Sequence
				_ = _2056_v49
				_2056_v49 = _dafny.UnicodeSeqOfUtf8Bytes("psooqryel")
				_2055_cf0 = _dafny.Companion_Sequence_.IsPrefixOf(_2056_v49, _dafny.UnicodeSeqOfUtf8Bytes("udq"))
				var _2057_v50 _dafny.Sequence
				_ = _2057_v50
				_2057_v50 = _dafny.SeqOf(_2012_v15)
				var _rhs390 _dafny.Sequence = _2056_v49
				_ = _rhs390
				var _rhs391 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-128), (_this).F9())
				_ = _rhs391
				var _rhs392 bool = true
				_ = _rhs392
				var _rhs393 bool = _2055_cf0
				_ = _rhs393
				var _rhs394 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_2057_v50, _2057_v50)
				_ = _rhs394
				var _lhs265 *GlobalState = globalState
				_ = _lhs265
				_2056_v49 = _rhs390
				_lhs265.F1 = _rhs391
				_2014_v17 = _rhs392
				_2014_v17 = _rhs393
				_2055_cf0 = _rhs394
				var _2058_v51 _dafny.Sequence
				_ = _2058_v51
				_2058_v51 = _dafny.SeqOf(_2014_v17)
				_2014_v17 = (true) || ((_2058_v51).Select((Companion_Default___.SafeIndex((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2058_v51).Cardinality()))).Uint32()).(bool))
				_2014_v17 = !(((_2019_v22).F11()) == (!(_dafny.MultiSetFromSeq(_2058_v51)).Contains((_2019_v22).F11())))
			} else {
				var _2059___mcc_h3 D0 = _source27.Get_().(D0_DC3).Cf3
				_ = _2059___mcc_h3
				var _2060_cf3 D0 = _2059___mcc_h3
				_ = _2060_cf3
				var _2061_v52 _dafny.Array
				_ = _2061_v52
				_2061_v52 = _2012_v15
				var _2062_v53 _dafny.Set
				_ = _2062_v53
				_2062_v53 = _dafny.SetOf(_2061_v52, _2061_v52, _2061_v52, _2061_v52)
				_2062_v53 = _2062_v53
				var _2063_v54 D0
				_ = _2063_v54
				_2063_v54 = Companion_D0_.Create_DC0_((_2019_v22).F11())
				var _2064_v55 _dafny.Set
				_ = _2064_v55
				_2064_v55 = _dafny.SetOf(_2063_v54, Companion_D0_.Create_DC0_(_2014_v17))
				_2014_v17 = !(_2064_v55).Equals(_2064_v55)
				var _2065_v56 D0
				_ = _2065_v56
				_2065_v56 = Companion_D0_.Create_DC2_(_2037_v38, (_this).F5())
				var _2066_v57 _dafny.Sequence
				_ = _2066_v57
				_2066_v57 = _dafny.SeqOf(_2065_v56)
				var _2067_v58 T4
				_ = _2067_v58
				var _nw349 *C2 = New_C2_()
				_ = _nw349
				_nw349.Ctor__((_this).F5(), (_this).F6())
				_2067_v58 = _nw349
				var _2068_v59 _dafny.MultiSet
				_ = _2068_v59
				_2068_v59 = _dafny.MultiSetOf(_2067_v58)
				var _2069_v60 _dafny.Map
				_ = _2069_v60
				_2069_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_2066_v57), _2068_v59)
				var _2070_v61 _dafny.Sequence
				_ = _2070_v61
				_2070_v61 = _dafny.SeqOf(_2066_v57, _2066_v57)
				_2069_v60 = (_2069_v60).Update(_dafny.Companion_Sequence_.Concatenate(_2070_v61, _2070_v61), (_2068_v59).Intersection(_2068_v59))
				var _2071_v62 _dafny.Sequence
				_ = _2071_v62
				_2071_v62 = _dafny.SeqOf(true)
				_1995_v0 = (_1995_v0).Update((_2071_v62).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2071_v62).Cardinality()))).Uint32()).(bool), ((_2012_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(106), _dafny.ArrayLen((_2012_v15), 0))).Int()).(_dafny.Int)).Cmp(_this.F10) >= 0)
			}
		} else {
			var _2072_v63 bool
			_ = _2072_v63
			_2072_v63 = true
			var _2073_v64 _dafny.Set
			_ = _2073_v64
			_2073_v64 = _dafny.SetOf(!(_2072_v63), _2072_v63)
			var _2074_v65 _dafny.Map
			_ = _2074_v65
			_2074_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_2073_v64).IsDisjointFrom(_2073_v64))
			_2072_v63 = (func() bool {
				if (_2074_v65).Contains((_this).F9()) {
					return (_2074_v65).Get((_this).F9()).(bool)
				}
				return (_this).F6()
			})()
			_2072_v63 = _2072_v63
			var _2075_v66 _dafny.Sequence
			_ = _2075_v66
			_2075_v66 = _dafny.SeqOf(false)
			var _2076_v67 _dafny.Sequence
			_ = _2076_v67
			_2076_v67 = _2075_v66
			var _2077_v68 _dafny.Sequence
			_ = _2077_v68
			_2077_v68 = _dafny.UnicodeSeqOfUtf8Bytes("ekfgutbiw")
			var _2078_v69 _dafny.Sequence
			_ = _2078_v69
			_2078_v69 = _dafny.SeqOf(_dafny.IntOfUint32((_2077_v68).Cardinality()))
			var _rhs395 _dafny.Sequence = (_2076_v67)
			_ = _rhs395
			var _rhs396 bool = ((_2073_v64).IsSubsetOf(_2073_v64)) == ((_2075_v66).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_this.F10), _dafny.IntOfUint32((_2075_v66).Cardinality()))).Uint32()).(bool))
			_ = _rhs396
			var _rhs397 bool = !_dafny.Companion_Sequence_.Contains(_2078_v69, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(652), (_dafny.Zero).Minus((_this).F9())))
			_ = _rhs397
			var _rhs398 bool = (_this).F6()
			_ = _rhs398
			_2075_v66 = _rhs395
			_2072_v63 = _rhs396
			_2072_v63 = _rhs397
			_2072_v63 = _rhs398
			var _2079_v70 _dafny.Array
			_ = _2079_v70
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_54
			var _nw350 _dafny.Array
			_ = _nw350
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw350 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) bool = (func(_2080_v65 _dafny.Map) func(_dafny.Int) bool {
					return func(_2081_i4 _dafny.Int) bool {
						return (func() bool {
							if (_2080_v65).Contains(_dafny.IntOfUint32((_dafny.SeqOf((_this).F5(), (_this).F5(), (_this).F5())).Cardinality())) {
								return (_2080_v65).Get(_dafny.IntOfUint32((_dafny.SeqOf((_this).F5(), (_this).F5(), (_this).F5())).Cardinality())).(bool)
							}
							return true
						})()
					}
				})(_2074_v65)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw350 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw350).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw350).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2079_v70 = _nw350
			_2079_v70 = _2079_v70
			var _2082_v71 _dafny.Array
			_ = _2082_v71
			var _nw351 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
			_ = _nw351
			_2082_v71 = _nw351
			var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_2082_v71), 0))
			_ = _index351
			(_2082_v71).ArraySet1(_2079_v70, (_index351).Int())
			var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_2082_v71), 0))
			_ = _index352
			var _rhs399 bool = (_dafny.IntOfUint32((_2077_v68).Cardinality())).Cmp(_this.F10) >= 0
			_ = _rhs399
			var _rhs400 _dafny.Int = _this.F10
			_ = _rhs400
			var _rhs401 _dafny.Int = (_this).F9()
			_ = _rhs401
			var _rhs402 _dafny.Array = _2079_v70
			_ = _rhs402
			var _rhs403 _dafny.Int = _dafny.IntOfInt64(848)
			_ = _rhs403
			var _lhs266 *C7 = _this
			_ = _lhs266
			var _lhs267 *GlobalState = globalState
			_ = _lhs267
			var _lhs268 _dafny.Array = _2082_v71
			_ = _lhs268
			var _lhs269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_2082_v71), 0))
			_ = _lhs269
			var _lhs270 *GlobalState = globalState
			_ = _lhs270
			_2072_v63 = _rhs399
			_lhs266.F10 = _rhs400
			_lhs267.F1 = _rhs401
			(_lhs268).ArraySet1(_rhs402, (_lhs269).Int())
			_lhs270.F1 = _rhs403
		}
		var _2083_v72 bool
		_ = _2083_v72
		_2083_v72 = false
		_2083_v72 = _2083_v72
		var _hi11 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_this.F10, (_this).F9()))
		_ = _hi11
		for _2084_i5 := _dafny.IntOfInt64(-364); _2084_i5.Cmp(_hi11) < 0; _2084_i5 = _2084_i5.Plus(_dafny.One) {
			var _2085_v73 _dafny.Array
			_ = _2085_v73
			var _nw352 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw352
			_2085_v73 = _nw352
			var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_2085_v73), 0))
			_ = _index353
			(_2085_v73).ArraySet1((_this).F6(), (_index353).Int())
			var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_2085_v73), 0))
			_ = _index354
			(_2085_v73).ArraySet1(!(_2083_v72), (_index354).Int())
			var _2086_v74 *C1
			_ = _2086_v74
			var _nw353 *C1 = New_C1_()
			_ = _nw353
			_nw353.Ctor__()
			_2086_v74 = _nw353
			var _2087_v75 _dafny.MultiSet
			_ = _2087_v75
			_2087_v75 = _dafny.MultiSetOf(_dafny.IntOfInt64(54))
			var _2088_v76 _dafny.Sequence
			_ = _2088_v76
			_2088_v76 = _dafny.SeqOf(true)
			var _2089_v77 _dafny.Map
			_ = _2089_v77
			_2089_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _2087_v75)
			var _2090_v78 _dafny.Sequence
			_ = _2090_v78
			_2090_v78 = _dafny.SeqOf((_this).F9(), _this.F10, (_this).F9())
			var _2091_v79 _dafny.Sequence
			_ = _2091_v79
			_2091_v79 = _dafny.SeqOf(_2090_v78, _2090_v78)
			var _2092_v80 *C0
			_ = _2092_v80
			var _nw354 *C0 = New_C0_()
			_ = _nw354
			_nw354.Ctor__((_2086_v74).Fm6(_dafny.Companion_Sequence_.Update(_2090_v78, (Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2090_v78).Cardinality()))).Uint32(), _2084_i5), globalState), _2091_v79)
			_2092_v80 = _nw354
			var _2093_v81 _dafny.MultiSet
			_ = _2093_v81
			_2093_v81 = _dafny.MultiSetOf(_2092_v80, _2092_v80)
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_2085_v73), 0))
			_ = _index355
			var _rhs404 _dafny.MultiSet = (func() _dafny.MultiSet {
				if (_2089_v77).Contains(_2084_i5) {
					return (_2089_v77).Get(_2084_i5).(_dafny.MultiSet)
				}
				return _2087_v75
			})()
			_ = _rhs404
			var _rhs405 _dafny.Int = (func() _dafny.Int {
				if (_2093_v81).Contains(_2092_v80) {
					return (_2093_v81).Multiplicity(_2092_v80)
				}
				return (_this).F9()
			})()
			_ = _rhs405
			var _rhs406 _dafny.Sequence = _2088_v76
			_ = _rhs406
			var _rhs407 bool = (_this).Fm6(_dafny.Companion_Sequence_.Concatenate(_2090_v78, _2090_v78), globalState)
			_ = _rhs407
			var _lhs271 _dafny.Array = _2085_v73
			_ = _lhs271
			var _lhs272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_2085_v73), 0))
			_ = _lhs272
			_2087_v75 = _rhs404
			r0 = _rhs405
			_2088_v76 = _rhs406
			(_lhs271).ArraySet1(_rhs407, (_lhs272).Int())
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_2085_v73), 0))
			_ = _index356
			(_2085_v73).ArraySet1(_2083_v72, (_index356).Int())
		}
		var _2094_i6 _dafny.Int
		_ = _2094_i6
		_2094_i6 = _dafny.Zero
		{
			for Companion_Default___.Fm0(!(_2083_v72), globalState) {
				{
					if (_2094_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L21
					}
					_2094_i6 = (_2094_i6).Plus(_dafny.One)
					if (_this).F6() {
						var _2095_v82 _dafny.Set
						_ = _2095_v82
						_2095_v82 = _dafny.SetOf((_this).F5(), (_this).F5(), (_this).F5(), (_this).F5())
						var _2096_v83 _dafny.Array
						_ = _2096_v83
						var _len0_55 _dafny.Int = _dafny.IntOfInt64(10)
						_ = _len0_55
						var _nw355 _dafny.Array
						_ = _nw355
						if _len0_55.Cmp(_dafny.Zero) == 0 {
							_nw355 = _dafny.NewArray(_len0_55)
						} else {
							var _init55 func(_dafny.Int) _dafny.Int = func(_2097_i7 _dafny.Int) _dafny.Int {
								return (_2097_i7).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ukbkpwm")).Cardinality()))
							}
							_ = _init55
							var _element0_55 = _init55(_dafny.Zero)
							_ = _element0_55
							_nw355 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
							(_nw355).ArraySet1(_element0_55, 0)
							var _nativeLen0_55 = (_len0_55).Int()
							_ = _nativeLen0_55
							for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
								(_nw355).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
							}
						}
						_2096_v83 = _nw355
						var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))
						_ = _index357
						(_2096_v83).ArraySet1(_this.F10, (_index357).Int())
						var _2098_v84 D2
						_ = _2098_v84
						_2098_v84 = Companion_D2_.Create_DC8_(_2083_v72)
						var _2099_v85 _dafny.MultiSet
						_ = _2099_v85
						_2099_v85 = _dafny.MultiSetOf((_this).F9())
						var _2100_v86 _dafny.MultiSet
						_ = _2100_v86
						_2100_v86 = _dafny.MultiSetOf((_this).F9(), (_this).F9(), _this.F10, (_2099_v85).Cardinality())
						var _2101_v87 _dafny.MultiSet
						_ = _2101_v87
						_2101_v87 = _dafny.MultiSetOf(_this.F10, (_2099_v85).Cardinality())
						var _2102_v88 D6
						_ = _2102_v88
						_2102_v88 = Companion_D6_.Create_DC14_(_2083_v72)
						var _2103_v89 _dafny.Sequence
						_ = _2103_v89
						_2103_v89 = _dafny.UnicodeSeqOfUtf8Bytes("caicnmft")
						var _2104_v90 D2
						_ = _2104_v90
						_2104_v90 = Companion_D2_.Create_DC7_(_this.F10, _dafny.IntOfUint32((_2103_v89).Cardinality()))
						var _2105_v91 _dafny.Sequence
						_ = _2105_v91
						_2105_v91 = _dafny.SeqOf(_dafny.SeqOf((_2104_v90).Dtor_cf7(), (_this).F9()))
						var _2106_v92 *C0
						_ = _2106_v92
						var _nw356 *C0 = New_C0_()
						_ = _nw356
						_nw356.Ctor__((_2102_v88).Dtor_cf15(), _2105_v91)
						_2106_v92 = _nw356
						var _2107_v93 _dafny.Map
						_ = _2107_v93
						_2107_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2106_v92, (_this).F9())
						var _2108_v94 _dafny.Sequence
						_ = _2108_v94
						_2108_v94 = _dafny.SeqOf(_2083_v72, (_this).F6(), false)
						var _2109_v95 _dafny.Array
						_ = _2109_v95
						var _nwElement0_77 bool = !((_this).F6()) || (false)
						_ = _nwElement0_77
						var _nw357 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(15))
						_ = _nw357
						(_nw357).ArraySet1(_nwElement0_77, 0)
						(_nw357).ArraySet1((_this).F6(), 1)
						(_nw357).ArraySet1((_2098_v84).Dtor_cf9(), 2)
						(_nw357).ArraySet1((_2099_v85).IsProperSubsetOf(_2100_v86), 3)
						(_nw357).ArraySet1(_2083_v72, 4)
						(_nw357).ArraySet1(_2083_v72, 5)
						(_nw357).ArraySet1((_this).F6(), 6)
						(_nw357).ArraySet1(((_this).F9()).Cmp((_2107_v93).Cardinality()) < 0, 7)
						(_nw357).ArraySet1((_2106_v92).F11(), 8)
						(_nw357).ArraySet1((_this).F6(), 9)
						(_nw357).ArraySet1(!((false) && (false)), 10)
						(_nw357).ArraySet1((_this).F6(), 11)
						(_nw357).ArraySet1((_2106_v92).F11(), 12)
						(_nw357).ArraySet1((_2108_v94).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2108_v94).Cardinality()))).Uint32()).(bool), 13)
						(_nw357).ArraySet1(false, 14)
						_2109_v95 = _nw357
						var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_2109_v95), 0))
						_ = _index358
						(_2109_v95).ArraySet1((_this).F6(), (_index358).Int())
						var _2110_v96 _dafny.Sequence
						_ = _2110_v96
						_2110_v96 = _dafny.SeqOf((_this).F9(), _dafny.IntOfUint32((_2103_v89).Cardinality()), _dafny.IntOfUint32((_2108_v94).Cardinality()))
						var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))
						_ = _index359
						var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_2109_v95), 0))
						_ = _index360
						var _rhs408 _dafny.Set = _2095_v82
						_ = _rhs408
						var _rhs409 _dafny.Int = (_this).F9()
						_ = _rhs409
						var _rhs410 bool = (_this.F10).Cmp((func() _dafny.Int {
							if (_2106_v92).F11() {
								return _this.F10
							}
							return _this.F10
						})()) <= 0
						_ = _rhs410
						var _rhs411 bool = _2083_v72
						_ = _rhs411
						var _rhs412 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2110_v96, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(975))).Uint32(), func(coer110 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg110 _dafny.Int) interface{} {
								return coer110(arg110)
							}
						}((func(_2111_v89 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
							return func(_2112_i8 _dafny.Int) _dafny.Int {
								return _dafny.IntOfUint32((_2111_v89).Cardinality())
							}
						})(_2103_v89))))
						_ = _rhs412
						var _lhs273 _dafny.Array = _2096_v83
						_ = _lhs273
						var _lhs274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))
						_ = _lhs274
						var _lhs275 _dafny.Array = _2109_v95
						_ = _lhs275
						var _lhs276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(875), _dafny.ArrayLen((_2109_v95), 0))
						_ = _lhs276
						_2095_v82 = _rhs408
						(_lhs273).ArraySet1(_rhs409, (_lhs274).Int())
						_2083_v72 = _rhs410
						(_lhs275).ArraySet1(_rhs411, (_lhs276).Int())
						_2110_v96 = _rhs412
						(globalState).F1 = ((_this).F9()).Minus((_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int))
						var _2113_v97 _dafny.Array
						_ = _2113_v97
						var _nw358 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
						_ = _nw358
						_2113_v97 = _nw358
						var _2114_v98 _dafny.Map
						_ = _2114_v98
						_2114_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _2113_v97)
						_2114_v98 = (_2114_v98).Update((_dafny.Zero).Minus((_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int)), _2113_v97)
						var _2115_v99 *C1
						_ = _2115_v99
						var _nw359 *C1 = New_C1_()
						_ = _nw359
						_nw359.Ctor__()
						_2115_v99 = _nw359
						var _2116_v100 _dafny.MultiSet
						_ = _2116_v100
						_2116_v100 = _dafny.MultiSetOf(_2109_v95, _2109_v95, _2109_v95, _2109_v95)
						var _2117_v101 _dafny.Map
						_ = _2117_v101
						_2117_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), _2116_v100)
						var _2118_v102 _dafny.Map
						_ = _2118_v102
						_2118_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2083_v72, _2109_v95)
						var _2119_v103 _dafny.Map
						_ = _2119_v103
						_2119_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int), _2116_v100)
						var _2120_v104 _dafny.Sequence
						_ = _2120_v104
						_2120_v104 = _dafny.SeqOf(_2116_v100)
						var _2121_v105 _dafny.Array
						_ = _2121_v105
						var _nwElement0_78 _dafny.MultiSet = _2116_v100
						_ = _nwElement0_78
						var _nw360 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_78, nil, _dafny.IntOfInt64(28))
						_ = _nw360
						(_nw360).ArraySet1(_nwElement0_78, 0)
						(_nw360).ArraySet1((_dafny.MultiSetOf(_2109_v95, _2109_v95)).Intersection(_2116_v100), 1)
						(_nw360).ArraySet1(_dafny.MultiSetOf(_2109_v95), 2)
						(_nw360).ArraySet1((func() _dafny.MultiSet {
							if (_2117_v101).Contains((_this).F5()) {
								return (_2117_v101).Get((_this).F5()).(_dafny.MultiSet)
							}
							return _2116_v100
						})(), 3)
						(_nw360).ArraySet1((_2116_v100).Update(_2109_v95, Companion_Default___.Abs((_this).F9())), 4)
						(_nw360).ArraySet1(_2116_v100, 5)
						(_nw360).ArraySet1(_dafny.MultiSetOf(_2109_v95, _2109_v95, _2109_v95, _2109_v95, _2109_v95), 6)
						(_nw360).ArraySet1(_2116_v100, 7)
						(_nw360).ArraySet1((_2116_v100).Intersection(_dafny.MultiSetOf(_2109_v95, _2109_v95, (func() _dafny.Array {
							if (_2118_v102).Contains(!(!((_this).F6()))) {
								return (_2118_v102).Get(!(!((_this).F6()))).(_dafny.Array)
							}
							return _2109_v95
						})(), _2109_v95)), 8)
						(_nw360).ArraySet1((func() _dafny.MultiSet {
							if (_2119_v103).Contains((_this).F9()) {
								return (_2119_v103).Get((_this).F9()).(_dafny.MultiSet)
							}
							return _2116_v100
						})(), 9)
						(_nw360).ArraySet1((_2120_v104).Select((Companion_Default___.SafeIndex((_2115_v99).Fm4(_2108_v94, _this.F10, (_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf(_dafny.IntOfInt64(-337)), globalState), _dafny.IntOfUint32((_2120_v104).Cardinality()))).Uint32()).(_dafny.MultiSet), 10)
						(_nw360).ArraySet1(_2116_v100, 11)
						(_nw360).ArraySet1((_2116_v100).Difference(_2116_v100), 12)
						(_nw360).ArraySet1((_2116_v100).Difference(_2116_v100), 13)
						(_nw360).ArraySet1(_2116_v100, 14)
						(_nw360).ArraySet1(_2116_v100, 15)
						(_nw360).ArraySet1((func() _dafny.MultiSet {
							if (_2119_v103).Contains((_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int)) {
								return (_2119_v103).Get((_2096_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(862), _dafny.ArrayLen((_2096_v83), 0))).Int()).(_dafny.Int)).(_dafny.MultiSet)
							}
							return _2116_v100
						})(), 16)
						(_nw360).ArraySet1(((func() _dafny.MultiSet {
							if (_2119_v103).Contains((_this).F9()) {
								return (_2119_v103).Get((_this).F9()).(_dafny.MultiSet)
							}
							return _2116_v100
						})()).Intersection(_dafny.MultiSetOf(_2109_v95)), 17)
						(_nw360).ArraySet1(_dafny.MultiSetOf(_2109_v95), 18)
						(_nw360).ArraySet1(_dafny.MultiSetOf(_2109_v95), 19)
						(_nw360).ArraySet1(_2116_v100, 20)
						(_nw360).ArraySet1(_2116_v100, 21)
						(_nw360).ArraySet1((_2116_v100).Difference(_dafny.MultiSetOf(_2109_v95, _2109_v95)), 22)
						(_nw360).ArraySet1(_dafny.MultiSetOf(_2109_v95, _2109_v95, _2109_v95, _2109_v95), 23)
						(_nw360).ArraySet1(_2116_v100, 24)
						(_nw360).ArraySet1(_2116_v100, 25)
						(_nw360).ArraySet1((_2116_v100).Union(_2116_v100), 26)
						(_nw360).ArraySet1(_2116_v100, 27)
						_2121_v105 = _nw360
						var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_2121_v105), 0))
						_ = _index361
						(_2121_v105).ArraySet1(_2116_v100, (_index361).Int())
						var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_2121_v105), 0))
						_ = _index362
						(_2121_v105).ArraySet1((_2116_v100).Update(_2109_v95, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(739))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg111 _dafny.Int) interface{} {
								return coer111(arg111)
							}
						}(func(_2122_i9 _dafny.Int) _dafny.CodePoint {
							return (_this).F5()
						})), _2103_v89)).Cardinality()))), (_index362).Int())
					} else {
						var _2123_v106 _dafny.Map
						_ = _2123_v106
						_2123_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F5())
						_2123_v106 = (_2123_v106).Update(true, _dafny.CodePoint('s'))
						_2083_v72 = (_this).F6()
						var _2124_v107 _dafny.CodePoint
						_ = _2124_v107
						_2124_v107 = _dafny.CodePoint('e')
						_2124_v107 = _dafny.CodePoint('t')
						(_this).F10 = (_this).F9()
						r0 = _this.F10
					}
					(globalState).F1 = _this.F10
					r0 = _dafny.IntOfInt64(-304)
					var _2125_v108 _dafny.Array
					_ = _2125_v108
					var _nw361 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
					_ = _nw361
					_2125_v108 = _nw361
					var _2126_v109 _dafny.Map
					_ = _2126_v109
					_2126_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(128), !(_2083_v72))
					var _2127_v110 T4
					_ = _2127_v110
					var _nw362 *C2 = New_C2_()
					_ = _nw362
					_nw362.Ctor__((_this).F5(), !((func() bool {
						if (_2126_v109).Contains((_this).F9()) {
							return (_2126_v109).Get((_this).F9()).(bool)
						}
						return (_this).F6()
					})()))
					_2127_v110 = _nw362
					var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_2125_v108), 0))
					_ = _index363
					(_2125_v108).ArraySet1(_2127_v110, (_index363).Int())
					var _2128_v111 _dafny.Map
					_ = _2128_v111
					_2128_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
						if !(_2083_v72) {
							return (_this).F9()
						}
						return _dafny.IntOfInt64(-367)
					})(), _2127_v110)
					var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_2125_v108), 0))
					_ = _index364
					(_2125_v108).ArraySet1((func() T4 {
						if (_2128_v111).Contains((_this).F9()) {
							return (_2128_v111).Get((_this).F9()).(T4)
						}
						return _2127_v110
					})(), (_index364).Int())
					goto C21
				}
			C21:
			}
			goto L21
		}
	L21:
		r0 = _this.F10
		return r0
	}
}
func (_this *C7) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _2129_v0 *C1
		_ = _2129_v0
		var _nw363 *C1 = New_C1_()
		_ = _nw363
		_nw363.Ctor__()
		_2129_v0 = _nw363
		var _2130_v1 _dafny.Sequence
		_ = _2130_v1
		_2130_v1 = _dafny.SeqOf(_this.F10)
		var _2131_v2 _dafny.Array
		_ = _2131_v2
		var _nwElement0_79 _dafny.Int = _this.F10
		_ = _nwElement0_79
		var _nw364 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_79, nil, _dafny.IntOfInt64(11))
		_ = _nw364
		(_nw364).ArraySet1(_nwElement0_79, 0)
		(_nw364).ArraySet1(((_this).F9()).Plus(_dafny.IntOfUint32((_2130_v1).Cardinality())), 1)
		(_nw364).ArraySet1((_this).F9(), 2)
		(_nw364).ArraySet1((_this).F9(), 3)
		(_nw364).ArraySet1((_dafny.Zero).Minus(_this.F10), 4)
		(_nw364).ArraySet1(_this.F10, 5)
		(_nw364).ArraySet1((_this.F10).Plus((_this).F9()), 6)
		(_nw364).ArraySet1(_dafny.IntOfInt64(37), 7)
		(_nw364).ArraySet1(_this.F10, 8)
		(_nw364).ArraySet1(_this.F10, 9)
		(_nw364).ArraySet1((_this).F9(), 10)
		_2131_v2 = _nw364
		var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
		_ = _index365
		(_2131_v2).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_dafny.Zero).Minus((p1).Cardinality())), _this.F10), (_index365).Int())
		var _2132_v3 bool
		_ = _2132_v3
		_2132_v3 = true
		var _2133_v4 _dafny.Array
		_ = _2133_v4
		var _nw365 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw365
		_2133_v4 = _nw365
		var _2134_v5 _dafny.CodePoint
		_ = _2134_v5
		_2134_v5 = _dafny.CodePoint('a')
		var _2135_v6 _dafny.Sequence
		_ = _2135_v6
		_2135_v6 = _dafny.SeqOf((_this).F6(), (_this).F6())
		var _2136_v7 _dafny.Sequence
		_ = _2136_v7
		_2136_v7 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, _dafny.IntOfUint32((_2135_v6).Cardinality())))
		var _2137_v8 _dafny.Map
		_ = _2137_v8
		_2137_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2129_v0).Fm14(globalState), (_this).F5())
		var _2138_v9 D1
		_ = _2138_v9
		_2138_v9 = Companion_D1_.Create_DC4_(_this.F10)
		var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
		_ = _index366
		var _rhs413 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2136_v7, _2136_v7)).Cardinality())
		_ = _rhs413
		var _rhs414 bool = true
		_ = _rhs414
		var _rhs415 _dafny.Array = _2133_v4
		_ = _rhs415
		var _rhs416 _dafny.CodePoint = (func() _dafny.CodePoint {
			if (_2137_v8).Contains((func(_pat_let34_0 D1) D1 {
				return func(_2139_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let35_0 _dafny.Int) D1 {
						return func(_2140_dt__update_hcf4_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_(_2140_dt__update_hcf4_h0)
						}(_pat_let35_0)
					}((_this).F9())
				}(_pat_let34_0)
			}(_2138_v9)).Dtor_cf4()) {
				return (_2137_v8).Get((func(_pat_let36_0 D1) D1 {
					return func(_2141_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let37_0 _dafny.Int) D1 {
							return func(_2142_dt__update_hcf4_h1 _dafny.Int) D1 {
								return Companion_D1_.Create_DC4_(_2142_dt__update_hcf4_h1)
							}(_pat_let37_0)
						}((_this).F9())
					}(_pat_let36_0)
				}(_2138_v9)).Dtor_cf4()).(_dafny.CodePoint)
			}
			return _dafny.CodePoint('b')
		})()
		_ = _rhs416
		var _lhs277 _dafny.Array = _2131_v2
		_ = _lhs277
		var _lhs278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
		_ = _lhs278
		(_lhs277).ArraySet1(_rhs413, (_lhs278).Int())
		_2132_v3 = _rhs414
		_2133_v4 = _rhs415
		_2134_v5 = _rhs416
		_2132_v3 = (((_this).F9()).Times((_this).F9())).Cmp(((_dafny.Zero).Minus((_this).F9())).Times(_dafny.IntOfInt64(580))) != 0
		var _2143_v10 _dafny.Sequence
		_ = _2143_v10
		_2143_v10 = _dafny.UnicodeSeqOfUtf8Bytes("jrsms")
		if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2143_v10, _dafny.UnicodeSeqOfUtf8Bytes("sqhcfc")), Companion_Default___.Fm1((_this).F9(), p2, _dafny.IntOfInt64(338), globalState)) {
			if false {
				var _2144_v11 _dafny.Array
				_ = _2144_v11
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_56
				var _nw366 _dafny.Array
				_ = _nw366
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw366 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) _dafny.Set = func(_2145_i0 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_(), Companion_D0_.Create_DC1_())
					}
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw366 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw366).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw366).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2144_v11 = _nw366
				var _2146_v12 D0
				_ = _2146_v12
				_2146_v12 = Companion_D0_.Create_DC1_()
				var _2147_v13 _dafny.Set
				_ = _2147_v13
				_2147_v13 = _dafny.SetOf(_2146_v12, _2146_v12)
				var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_2144_v11), 0))
				_ = _index367
				(_2144_v11).ArraySet1(_2147_v13, (_index367).Int())
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_2144_v11), 0))
				_ = _index368
				var _rhs417 _dafny.Set = (_2147_v13)
				_ = _rhs417
				var _rhs418 _dafny.Int = (_this).F9()
				_ = _rhs418
				var _rhs419 _dafny.Int = (_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int)
				_ = _rhs419
				var _lhs279 _dafny.Array = _2144_v11
				_ = _lhs279
				var _lhs280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_2144_v11), 0))
				_ = _lhs280
				var _lhs281 *GlobalState = globalState
				_ = _lhs281
				var _lhs282 *GlobalState = globalState
				_ = _lhs282
				(_lhs279).ArraySet1(_rhs417, (_lhs280).Int())
				_lhs281.F1 = _rhs418
				_lhs282.F1 = _rhs419
				(_this).F10 = (Companion_Default___.Fm3(false, (_this).F6(), globalState)).Times(_this.F10)
				var _2148_v14 _dafny.Map
				_ = _2148_v14
				_2148_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm27(_2143_v10, !(!(p2)), true, globalState), (_this).F9())
				var _2149_v15 _dafny.Map
				_ = _2149_v15
				_2149_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(_2131_v2)).Cardinality()))
				var _2150_v16 _dafny.Map
				_ = _2150_v16
				_2150_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(40))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_2151_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2152_i1 _dafny.Int) _dafny.CodePoint {
						return _2151_v5
					}
				})(_2134_v5)))).Cardinality()), _this.F10)
				var _2153_v17 D2
				_ = _2153_v17
				_2153_v17 = Companion_D2_.Create_DC8_(_2132_v3)
				_2148_v14 = (_2148_v14).Update(_2150_v16, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm1((_dafny.SetOf((_this).F6())).Cardinality(), (_2153_v17).Dtor_cf9(), (_this).F9(), globalState), _2143_v10)).Cardinality()))
				var _2154_v18 *C2
				_ = _2154_v18
				var _nw367 *C2 = New_C2_()
				_ = _nw367
				_nw367.Ctor__((_this).F5(), _2132_v3)
				_2154_v18 = _nw367
				var _2155_v19 _dafny.Array
				_ = _2155_v19
				var _nwElement0_80 *C2 = _2154_v18
				_ = _nwElement0_80
				var _nw368 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_80, nil, _dafny.IntOfInt64(29))
				_ = _nw368
				(_nw368).ArraySet1(_nwElement0_80, 0)
				(_nw368).ArraySet1(_2154_v18, 1)
				(_nw368).ArraySet1(_2154_v18, 2)
				(_nw368).ArraySet1(_2154_v18, 3)
				(_nw368).ArraySet1(_2154_v18, 4)
				(_nw368).ArraySet1(_2154_v18, 5)
				(_nw368).ArraySet1(_2154_v18, 6)
				(_nw368).ArraySet1(_2154_v18, 7)
				(_nw368).ArraySet1(_2154_v18, 8)
				(_nw368).ArraySet1(_2154_v18, 9)
				(_nw368).ArraySet1(_2154_v18, 10)
				(_nw368).ArraySet1(_2154_v18, 11)
				(_nw368).ArraySet1(_2154_v18, 12)
				(_nw368).ArraySet1(_2154_v18, 13)
				(_nw368).ArraySet1(_2154_v18, 14)
				(_nw368).ArraySet1(_2154_v18, 15)
				(_nw368).ArraySet1(_2154_v18, 16)
				(_nw368).ArraySet1(_2154_v18, 17)
				(_nw368).ArraySet1(_2154_v18, 18)
				(_nw368).ArraySet1(_2154_v18, 19)
				(_nw368).ArraySet1(_2154_v18, 20)
				(_nw368).ArraySet1(_2154_v18, 21)
				(_nw368).ArraySet1(_2154_v18, 22)
				(_nw368).ArraySet1(_2154_v18, 23)
				(_nw368).ArraySet1(_2154_v18, 24)
				(_nw368).ArraySet1(_2154_v18, 25)
				(_nw368).ArraySet1(_2154_v18, 26)
				(_nw368).ArraySet1(_2154_v18, 27)
				(_nw368).ArraySet1(_2154_v18, 28)
				_2155_v19 = _nw368
				var _2156_v20 _dafny.Sequence
				_ = _2156_v20
				_2156_v20 = _dafny.SeqOf(_2155_v19)
				var _2157_v21 _dafny.Map
				_ = _2157_v21
				_2157_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2133_v4, (_2156_v20).Select((Companion_Default___.SafeIndex((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2156_v20).Cardinality()))).Uint32()).(_dafny.Array))
				var _2158_v22 _dafny.Map
				_ = _2158_v22
				_2158_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2133_v4, _2155_v19)).Update(_2133_v4, _2155_v19)).Update(_2133_v4, _2155_v19))
				var _2159_v23 _dafny.Sequence
				_ = _2159_v23
				_2159_v23 = _2135_v6
				var _2160_v24 _dafny.Set
				_ = _2160_v24
				_2160_v24 = _dafny.SetOf(_2135_v6, _2135_v6, _2135_v6, _dafny.Companion_Sequence_.Update((_2159_v23), (Companion_Default___.SafeIndex((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2159_v23).Cardinality()))).Uint32(), (_this).F6()))
				_2157_v21 = (func() _dafny.Map {
					if (_2158_v22).Contains((_dafny.SetOf(_2135_v6)).IsSubsetOf(_2160_v24)) {
						return (_2158_v22).Get((_dafny.SetOf(_2135_v6)).IsSubsetOf(_2160_v24)).(_dafny.Map)
					}
					return _2157_v21
				})()
				var _2161_v25 _dafny.Map
				_ = _2161_v25
				_2161_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, !(_2132_v3))
				var _2162_v26 _dafny.Set
				_ = _2162_v26
				_2162_v26 = _dafny.SetOf(_2161_v25, _2161_v25, _2161_v25, _2161_v25)
				var _2163_v28 _dafny.Sequence
				_ = _2163_v28
				_2163_v28 = _dafny.SeqOf(_2162_v26, _2162_v26)
				_2162_v26 = (_2162_v26).Union((func() _dafny.Set {
					var _coll65 = _dafny.NewBuilder()
					_ = _coll65
					for _iter70 := _dafny.Iterate((_2162_v26).Elements()); ; {
						_compr_65, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _2164_v27 _dafny.Map
						_2164_v27 = interface{}(_compr_65).(_dafny.Map)
						if (_2162_v26).Contains(_2164_v27) {
							_coll65.Add(_2164_v27)
						}
					}
					return _coll65.ToSet()
				}()).Intersection((_2163_v28).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2163_v28).Cardinality()))).Uint32()).(_dafny.Set)))
			} else {
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
				_ = _index369
				(_2131_v2).ArraySet1(_dafny.IntOfUint32((_2135_v6).Cardinality()), (_index369).Int())
				_2132_v3 = p0
				_2132_v3 = (_this).F6()
				(globalState).F1 = ((_dafny.Zero).Minus((_this).Fm7(true, globalState))).Minus(_dafny.IntOfInt64(759))
				var _2165_v29 _dafny.Sequence
				_ = _2165_v29
				_2165_v29 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(294))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}(func(_2166_i2 _dafny.Int) _dafny.Int {
					return (_this).F9()
				})))
				var _2167_v30 T0
				_ = _2167_v30
				var _nw369 *C0 = New_C0_()
				_ = _nw369
				_nw369.Ctor__(p0, _2165_v29)
				_2167_v30 = _nw369
				var _2168_v31 _dafny.Sequence
				_ = _2168_v31
				_2168_v31 = _dafny.SeqOf((func() T0 {
					if p0 {
						return _2167_v30
					}
					return _2167_v30
				})(), _2167_v30, _2167_v30, _2167_v30)
				_2168_v31 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2168_v31, _dafny.SeqOf(_2167_v30)), _2168_v31)
			}
			var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
			_ = _index370
			(_2131_v2).ArraySet1((_dafny.Zero).Minus(_this.F10), (_index370).Int())
			_2132_v3 = true
			var _2169_v32 _dafny.Sequence
			_ = _2169_v32
			_2169_v32 = _dafny.SeqOf(_2130_v1)
			var _2170_v33 T0
			_ = _2170_v33
			var _nw370 *C0 = New_C0_()
			_ = _nw370
			_nw370.Ctor__(p2, _2169_v32)
			_2170_v33 = _nw370
			_2170_v33 = _2170_v33
			if ((_this).F9()).Cmp((_dafny.Zero).Minus((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int))) != 0 {
				var _2171_v34 _dafny.Set
				_ = _2171_v34
				_2171_v34 = _dafny.SetOf((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _this.F10)
				var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
				_ = _index371
				(_2131_v2).ArraySet1((_this).Fm8(p2, _2132_v3, (_2171_v34).IsDisjointFrom(_2171_v34), globalState), (_index371).Int())
				(globalState).F4 = p1
				var _2172_v35 _dafny.Array
				_ = _2172_v35
				var _nw371 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
				_ = _nw371
				_2172_v35 = _nw371
				var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_2172_v35), 0))
				_ = _index372
				(_2172_v35).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(448))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}(func(_2173_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				})), _2143_v10), (_index372).Int())
				var _2174_v36 _dafny.Map
				_ = _2174_v36
				_2174_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F9()), _2135_v6)
				var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(926), _dafny.ArrayLen((_2172_v35), 0))
				_ = _index373
				(_2172_v35).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}((func(_2175_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2176_i4 _dafny.Int) _dafny.CodePoint {
						return _2175_v5
					}
				})(_2134_v5))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-26))).Uint32(), func(coer116 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg116 _dafny.Int) interface{} {
						return coer116(arg116)
					}
				}(func(_2177_i5 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				})), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2143_v10, (Companion_Default___.SafeIndex((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2143_v10).Cardinality()))).Uint32(), (_this).F5()), (Companion_Default___.SafeIndex((_2174_v36).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2143_v10, (Companion_Default___.SafeIndex((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2143_v10).Cardinality()))).Uint32(), (_this).F5())).Cardinality()))).Uint32(), _dafny.CodePoint('r')))), (_index373).Int())
				var _rhs420 _dafny.Array = _2133_v4
				_ = _rhs420
				var _rhs421 bool = (_this).F6()
				_ = _rhs421
				_2133_v4 = _rhs420
				_2132_v3 = _rhs421
				var _2178_v37 _dafny.Sequence
				_ = _2178_v37
				_2178_v37 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-664))).Uint32(), func(coer117 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg117 _dafny.Int) interface{} {
						return coer117(arg117)
					}
				}((func(_2179_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2180_i6 _dafny.Int) _dafny.CodePoint {
						return _2179_v5
					}
				})(_2134_v5))), _dafny.UnicodeSeqOfUtf8Bytes("l")), _2143_v10)
				_2178_v37 = _2178_v37
			} else {
				var _2181_v38 _dafny.Map
				_ = _2181_v38
				_2181_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, (_this).F9())
				var _rhs422 bool = _2132_v3
				_ = _rhs422
				var _rhs423 bool = p2
				_ = _rhs423
				var _rhs424 _dafny.Map = _2181_v38
				_ = _rhs424
				var _rhs425 bool = !((_this.F10).Cmp(_dafny.IntOfUint32((_2143_v10).Cardinality())) == 0) || (_dafny.Companion_Sequence_.Equal(_2143_v10, _dafny.Companion_Sequence_.Update(_2143_v10, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.IntOfUint32((_2143_v10).Cardinality()))).Uint32(), _2134_v5)))
				_ = _rhs425
				_2132_v3 = _rhs422
				_2132_v3 = _rhs423
				_2181_v38 = _rhs424
				_2132_v3 = _rhs425
				var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
				_ = _index374
				(_2131_v2).ArraySet1((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), (_index374).Int())
				var _2182_v39 _dafny.Map
				_ = _2182_v39
				_2182_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2129_v0, _2133_v4)
				var _2183_v40 _dafny.Map
				_ = _2183_v40
				_2183_v40 = _2182_v39
				var _2184_v41 _dafny.Map
				_ = _2184_v41
				_2184_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2183_v40), _2143_v10)
				_2184_v41 = (_2184_v41).Update(((_2182_v39).Update(_2129_v0, _2133_v4)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2129_v0, _2133_v4)), _dafny.Companion_Sequence_.Concatenate(_2143_v10, _2143_v10))
				var _2185_v42 _dafny.Array
				_ = _2185_v42
				var _nw372 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw372
				_2185_v42 = _nw372
				var _2186_v43 T1
				_ = _2186_v43
				var _nw373 *C3 = New_C3_()
				_ = _nw373
				_nw373.Ctor__(p0, p2)
				_2186_v43 = _nw373
				var _2187_v44 _dafny.Map
				_ = _2187_v44
				_2187_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2138_v9, _2186_v43)
				var _2188_v45 _dafny.Sequence
				_ = _2188_v45
				_2188_v45 = _dafny.SeqOf(_2187_v44)
				var _2189_v46 _dafny.Map
				_ = _2189_v46
				_2189_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F10, (_dafny.MultiSetOf(p0)).Cardinality())
				var _2190_v47 _dafny.Set
				_ = _2190_v47
				_2190_v47 = _dafny.SetOf(_2187_v44, (_2188_v45).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_2189_v46).Contains(_dafny.IntOfUint32((_2143_v10).Cardinality())) {
						return (_2189_v46).Get(_dafny.IntOfUint32((_2143_v10).Cardinality())).(_dafny.Int)
					}
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_2135_v6).Select((Companion_Default___.SafeIndex((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2135_v6).Cardinality()))).Uint32()).(bool))).Cardinality()
				})(), _dafny.IntOfUint32((_2188_v45).Cardinality()))).Uint32()).(_dafny.Map))
				var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_2185_v42), 0))
				_ = _index375
				(_2185_v42).ArraySet1(_2190_v47, (_index375).Int())
				var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(878), _dafny.ArrayLen((_2185_v42), 0))
				_ = _index376
				(_2185_v42).ArraySet1(((_2190_v47).Intersection(_2190_v47)).Union(_2190_v47), (_index376).Int())
				var _2191_v48 *C3
				_ = _2191_v48
				var _nw374 *C3 = New_C3_()
				_ = _nw374
				_nw374.Ctor__((_this).F6(), false)
				_2191_v48 = _nw374
			}
		} else {
			var _2192_v49 *C6
			_ = _2192_v49
			var _nw375 *C6 = New_C6_()
			_ = _nw375
			_nw375.Ctor__()
			_2192_v49 = _nw375
			_2192_v49 = _2192_v49
			(_this).F10 = (_this).F9()
			if !(p0) || (p0) {
				_2132_v3 = true
				var _2193_v50 _dafny.Map
				_ = _2193_v50
				_2193_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2131_v2)
				_2193_v50 = (_2193_v50).Update(_2132_v3, _2131_v2)
				var _2194_v51 _dafny.Map
				_ = _2194_v51
				_2194_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_2143_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-377))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg118 _dafny.Int) interface{} {
						return coer118(arg118)
					}
				}((func(_2195_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2196_i7 _dafny.Int) _dafny.CodePoint {
						return _2195_v5
					}
				})(_2134_v5)))), (_this).F6())
				_2194_v51 = (_2194_v51).Update(_2143_v10, false)
				_2132_v3 = !(true)
				var _2197_v52 _dafny.Map
				_ = _2197_v52
				_2197_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), _2132_v3)
				var _2198_v53 _dafny.Map
				_ = _2198_v53
				_2198_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2135_v6).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(127), (_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), (_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int))).Cardinality()), _dafny.IntOfUint32((_2135_v6).Cardinality()))).Uint32()).(bool)), (_2197_v52).Update((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int), (_this).F6()))
				_2198_v53 = ((_2198_v53).Merge(_2198_v53)).Update(!(_dafny.Companion_Sequence_.IsPrefixOf(_2130_v1, _2130_v1)), (_2197_v52).Merge(_2197_v52))
			} else {
				var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
				_ = _index377
				var _rhs426 bool = ((_this).F6()) && (false)
				_ = _rhs426
				var _rhs427 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm3((_this).F6(), _2132_v3, globalState), (_this).F9()))
				_ = _rhs427
				var _lhs283 _dafny.Array = _2131_v2
				_ = _lhs283
				var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))
				_ = _lhs284
				_2132_v3 = _rhs426
				(_lhs283).ArraySet1(_rhs427, (_lhs284).Int())
				var _2199_v54 _dafny.Array
				_ = _2199_v54
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_57
				var _nw376 _dafny.Array
				_ = _nw376
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw376 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) D11 = func(_2200_i8 _dafny.Int) D11 {
						return Companion_D11_.Create_DC19_(_dafny.CodePoint('j'))
					}
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw376 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw376).ArraySet1(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw376).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2199_v54 = _nw376
				var _2201_v55 D11
				_ = _2201_v55
				_2201_v55 = Companion_D11_.Create_DC19_(_2134_v5)
				var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2199_v54), 0))
				_ = _index378
				(_2199_v54).ArraySet1(_2201_v55, (_index378).Int())
				var _2202_v56 _dafny.Sequence
				_ = _2202_v56
				_2202_v56 = _dafny.SeqOf(_2130_v1, _2130_v1)
				var _2203_v57 *C0
				_ = _2203_v57
				var _nw377 *C0 = New_C0_()
				_ = _nw377
				_nw377.Ctor__((_this).F6(), _2202_v56)
				_2203_v57 = _nw377
				var _2204_v58 _dafny.Sequence
				_ = _2204_v58
				_2204_v58 = _dafny.SeqOf(_2203_v57)
				var _2205_v59 _dafny.Set
				_ = _2205_v59
				_2205_v59 = _dafny.SetOf(_dafny.IntOfInt64(102))
				var _2206_v60 _dafny.Sequence
				_ = _2206_v60
				_2206_v60 = _dafny.SeqOf(_2205_v59)
				var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2199_v54), 0))
				_ = _index379
				var _rhs428 bool = _2132_v3
				_ = _rhs428
				var _rhs429 _dafny.Int = (_this).F9()
				_ = _rhs429
				var _rhs430 _dafny.Int = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _2132_v3 {
						return _2204_v58
					}
					return _2204_v58
				})(), _dafny.Companion_Sequence_.Concatenate(_2204_v58, _2204_v58)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F9(), _dafny.IntOfUint32((_2206_v60).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _2132_v3 {
						return _2204_v58
					}
					return _2204_v58
				})(), _dafny.Companion_Sequence_.Concatenate(_2204_v58, _2204_v58))).Cardinality()))).Uint32(), _2203_v57), (Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _2132_v3 {
						return _2204_v58
					}
					return _2204_v58
				})(), _dafny.Companion_Sequence_.Concatenate(_2204_v58, _2204_v58)), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_this).F9(), _dafny.IntOfUint32((_2206_v60).Cardinality())), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _2132_v3 {
						return _2204_v58
					}
					return _2204_v58
				})(), _dafny.Companion_Sequence_.Concatenate(_2204_v58, _2204_v58))).Cardinality()))).Uint32(), _2203_v57)).Cardinality()))).Uint32(), _2203_v57))).Cardinality()
				_ = _rhs430
				var _rhs431 bool = (_this).F6()
				_ = _rhs431
				var _rhs432 D11 = _2201_v55
				_ = _rhs432
				var _lhs285 *GlobalState = globalState
				_ = _lhs285
				var _lhs286 *GlobalState = globalState
				_ = _lhs286
				var _lhs287 _dafny.Array = _2199_v54
				_ = _lhs287
				var _lhs288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_2199_v54), 0))
				_ = _lhs288
				_2132_v3 = _rhs428
				_lhs285.F1 = _rhs429
				_lhs286.F1 = _rhs430
				_2132_v3 = _rhs431
				(_lhs287).ArraySet1(_rhs432, (_lhs288).Int())
				_2131_v2 = _2131_v2
				var _2207_v61 _dafny.Array
				_ = _2207_v61
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_58
				var _nw378 _dafny.Array
				_ = _nw378
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw378 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.Sequence = (func(_2208_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_2209_i9 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tbknyswhp"), _2208_v10)
						}
					})(_2143_v10)
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw378 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw378).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw378).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_2207_v61 = _nw378
				var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_2207_v61), 0))
				_ = _index380
				(_2207_v61).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("penmrr"), _2143_v10), (_index380).Int())
				var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(905), _dafny.ArrayLen((_2207_v61), 0))
				_ = _index381
				(_2207_v61).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kq"), _dafny.UnicodeSeqOfUtf8Bytes("onejtsx")), (_index381).Int())
				var _2210_v62 D2
				_ = _2210_v62
				_2210_v62 = Companion_D2_.Create_DC8_(_2132_v3)
				var _rhs433 _dafny.Int = (_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int)
				_ = _rhs433
				var _rhs434 D1 = (func() D1 {
					if ((_2131_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_2131_v2), 0))).Int()).(_dafny.Int)).Cmp(_this.F10) > 0 {
						return _2138_v9
					}
					return (_this).Fm5(_2210_v62, globalState)
				})()
				_ = _rhs434
				var _lhs289 *GlobalState = globalState
				_ = _lhs289
				_lhs289.F1 = _rhs433
				_2138_v9 = _rhs434
			}
			_2134_v5 = (func() _dafny.CodePoint {
				if p0 {
					return (func() _dafny.CodePoint {
						if p2 {
							return _2134_v5
						}
						return (_this).F5()
					})()
				}
				return (_this).F5()
			})()
			var _2211_v63 _dafny.Array
			_ = _2211_v63
			_2211_v63 = _2131_v2
			var _2212_v64 D12
			_ = _2212_v64
			_2212_v64 = Companion_D12_.Create_DC23_(_2211_v63)
			var _pat_let_tv29 = _2131_v2
			_ = _pat_let_tv29
			var _pat_let_tv30 = _2211_v63
			_ = _pat_let_tv30
			var _2213_v65 _dafny.Array
			_ = _2213_v65
			var _nwElement0_81 D12 = _2212_v64
			_ = _nwElement0_81
			var _nw379 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_81, nil, _dafny.IntOfInt64(13))
			_ = _nw379
			(_nw379).ArraySet1(_nwElement0_81, 0)
			(_nw379).ArraySet1(_2212_v64, 1)
			(_nw379).ArraySet1(_2212_v64, 2)
			(_nw379).ArraySet1(_2212_v64, 3)
			(_nw379).ArraySet1(_2212_v64, 4)
			(_nw379).ArraySet1(_2212_v64, 5)
			(_nw379).ArraySet1(_2212_v64, 6)
			(_nw379).ArraySet1(func(_pat_let38_0 D12) D12 {
				return func(_2216_dt__update__tmp_h4 D12) D12 {
					return func(_pat_let41_0 _dafny.Array) D12 {
						return func(_2217_dt__update_hcf29_h1 _dafny.Array) D12 {
							return Companion_D12_.Create_DC23_(_2217_dt__update_hcf29_h1)
						}(_pat_let41_0)
					}(_pat_let_tv30)
				}(_pat_let38_0)
			}(func(_pat_let39_0 D12) D12 {
				return func(_2214_dt__update__tmp_h3 D12) D12 {
					return func(_pat_let40_0 _dafny.Array) D12 {
						return func(_2215_dt__update_hcf29_h0 _dafny.Array) D12 {
							return Companion_D12_.Create_DC23_(_2215_dt__update_hcf29_h0)
						}(_pat_let40_0)
					}(_pat_let_tv29)
				}(_pat_let39_0)
			}(_2212_v64)), 7)
			(_nw379).ArraySet1(_2212_v64, 8)
			(_nw379).ArraySet1(_2212_v64, 9)
			(_nw379).ArraySet1(Companion_D12_.Create_DC23_(_2211_v63), 10)
			(_nw379).ArraySet1(Companion_D12_.Create_DC23_(_2211_v63), 11)
			(_nw379).ArraySet1(_2212_v64, 12)
			_2213_v65 = _nw379
			var _2218_v66 _dafny.Array
			_ = _2218_v66
			var _nwElement0_82 _dafny.Array = _2213_v65
			_ = _nwElement0_82
			var _nw380 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_82, nil, _dafny.IntOfInt64(2))
			_ = _nw380
			(_nw380).ArraySet1(_nwElement0_82, 0)
			(_nw380).ArraySet1(_2213_v65, 1)
			_2218_v66 = _nw380
			_2218_v66 = _2218_v66
		}
		var _2219_v67 _dafny.Set
		_ = _2219_v67
		_2219_v67 = _dafny.SetOf((_this).F9())
		var _2220_v68 _dafny.Map
		_ = _2220_v68
		_2220_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_dafny.MultiSetOf((_this).Fm6(_2130_v1, globalState))).Contains((Companion_D16_.Create_DC32_(p2)).Dtor_cf41()), _2219_v67)
		_2220_v68 = (_2220_v68).Update(false, _2219_v67)
		var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_2133_v4), 0))
		_ = _index382
		(_2133_v4).ArraySet1(p2, (_index382).Int())
		var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(806), _dafny.ArrayLen((_2133_v4), 0))
		_ = _index383
		(_2133_v4).ArraySet1((_this.F10).Cmp((_2130_v1).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2130_v1).Cardinality()))).Uint32()).(_dafny.Int)) > 0, (_index383).Int())
	}
}
func (_this *C7) M7(p0 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, _dafny.Sequence) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Sequence = _dafny.EmptySeq
		_ = r3
		var _2221_v0 D0
		_ = _2221_v0
		_2221_v0 = Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_((_this).F6()))
		_2221_v0 = (func() D0 {
			if (_this).F6() {
				return _2221_v0
			}
			return _2221_v0
		})()
		if (_this).F6() {
			var _2222_v1 _dafny.Map
			_ = _2222_v1
			_2222_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false)
			_2222_v1 = (_2222_v1).Merge(_2222_v1)
			var _2223_v2 *C3
			_ = _2223_v2
			var _nw381 *C3 = New_C3_()
			_ = _nw381
			_nw381.Ctor__((_this).F6(), (_this).F6())
			_2223_v2 = _nw381
			var _2224_v3 _dafny.Sequence
			_ = _2224_v3
			_2224_v3 = _dafny.SeqOf(p0, p0)
			var _2225_v4 _dafny.MultiSet
			_ = _2225_v4
			_2225_v4 = _dafny.MultiSetOf((_2224_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2224_v3).Cardinality()), _dafny.IntOfUint32((_2224_v3).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(-856))
			_2225_v4 = _2225_v4
			var _2226_v5 _dafny.CodePoint
			_ = _2226_v5
			_2226_v5 = _dafny.CodePoint('n')
			_2226_v5 = _2226_v5
			(globalState).F1 = (func() _dafny.Int {
				if !((_2223_v2).F13()) {
					return _this.F10
				}
				return _dafny.IntOfInt64(-730)
			})()
		} else {
			var _2227_v6 _dafny.Sequence
			_ = _2227_v6
			_2227_v6 = _dafny.UnicodeSeqOfUtf8Bytes("o")
			if !_dafny.Companion_Sequence_.Contains((Companion_D2_.Create_DC6_(_2227_v6)).Dtor_cf6(), (_this).F5()) {
				var _2228_v7 bool
				_ = _2228_v7
				_2228_v7 = false
				_2228_v7 = !(true)
				r0 = p0
				var _2229_v8 _dafny.Array
				_ = _2229_v8
				var _nw382 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(6))
				_ = _nw382
				_2229_v8 = _nw382
				var _2230_v9 _dafny.Array
				_ = _2230_v9
				var _nw383 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw383
				_2230_v9 = _nw383
				var _2231_v10 _dafny.Sequence
				_ = _2231_v10
				_2231_v10 = _dafny.SeqOf(_2230_v9, _2230_v9, _2230_v9)
				var _2232_v11 _dafny.Map
				_ = _2232_v11
				_2232_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2229_v8, (_2231_v10).Select((Companion_Default___.SafeIndex((_this).Fm8(_2228_v7, _2228_v7, (_this).F6(), globalState), _dafny.IntOfUint32((_2231_v10).Cardinality()))).Uint32()).(_dafny.Array))
				var _2233_v12 D17
				_ = _2233_v12
				_2233_v12 = Companion_D17_.Create_DC33_(_2232_v11)
				_2232_v11 = (_2233_v12).Dtor_cf42()
				var _2234_v13 *C6
				_ = _2234_v13
				var _nw384 *C6 = New_C6_()
				_ = _nw384
				_nw384.Ctor__()
				_2234_v13 = _nw384
				var _2235_v14 _dafny.Array
				_ = _2235_v14
				var _nw385 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw385
				_2235_v14 = _nw385
				var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_2235_v14), 0))
				_ = _index384
				(_2235_v14).ArraySet1(_2230_v9, (_index384).Int())
				var _2236_v15 _dafny.Map
				_ = _2236_v15
				_2236_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2227_v6)
				var _2237_v16 _dafny.Sequence
				_ = _2237_v16
				_2237_v16 = _dafny.SeqOf(_2227_v6, _2227_v6)
				var _2238_v17 _dafny.Map
				_ = _2238_v17
				_2238_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2237_v16).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_2237_v16).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _2239_v18 _dafny.Map
				_ = _2239_v18
				_2239_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _2240_v19 _dafny.Map
				_ = _2240_v19
				_2240_v19 = _2239_v18
				var _2241_v20 _dafny.Sequence
				_ = _2241_v20
				_2241_v20 = _dafny.SeqOf(_2228_v7)
				var _2242_v21 _dafny.Sequence
				_ = _2242_v21
				_2242_v21 = _dafny.SeqOf(_2241_v20, Companion_Default___.Fm2((_this).F6(), false, globalState), _2241_v20, _2241_v20)
				var _2243_v22 D12
				_ = _2243_v22
				_2243_v22 = Companion_D12_.Create_DC22_(_2240_v19, _2242_v21, (_this).F5(), _2228_v7, Companion_Default___.Fm1(_dafny.IntOfUint32((_2227_v6).Cardinality()), _2228_v7, p0, globalState))
				var _2244_v23 _dafny.Array
				_ = _2244_v23
				var _nwElement0_83 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_2236_v15).Contains(p0) {
						return (_2236_v15).Get(p0).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("v")
				})(), _dafny.UnicodeSeqOfUtf8Bytes("j"))
				_ = _nwElement0_83
				var _nw386 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_83, nil, _dafny.IntOfInt64(6))
				_ = _nw386
				(_nw386).ArraySet1(_nwElement0_83, 0)
				(_nw386).ArraySet1((func() _dafny.Sequence {
					if (_2238_v17).Contains(_2228_v7) {
						return (_2238_v17).Get(_2228_v7).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(3))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg119 _dafny.Int) interface{} {
							return coer119(arg119)
						}
					}(func(_2245_i0 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					}))
				})(), 1)
				(_nw386).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2227_v6, _2227_v6), 2)
				(_nw386).ArraySet1(_2227_v6, 3)
				(_nw386).ArraySet1(_2227_v6, 4)
				(_nw386).ArraySet1((_2243_v22).Dtor_cf28(), 5)
				_2244_v23 = _nw386
				var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2244_v23), 0))
				_ = _index385
				(_2244_v23).ArraySet1(_2227_v6, (_index385).Int())
				var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_2235_v14), 0))
				_ = _index386
				var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2244_v23), 0))
				_ = _index387
				var _rhs435 _dafny.Array = _2230_v9
				_ = _rhs435
				var _rhs436 _dafny.Sequence = _2227_v6
				_ = _rhs436
				var _lhs290 _dafny.Array = _2235_v14
				_ = _lhs290
				var _lhs291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(536), _dafny.ArrayLen((_2235_v14), 0))
				_ = _lhs291
				var _lhs292 _dafny.Array = _2244_v23
				_ = _lhs292
				var _lhs293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2244_v23), 0))
				_ = _lhs293
				(_lhs290).ArraySet1(_rhs435, (_lhs291).Int())
				(_lhs292).ArraySet1(_rhs436, (_lhs293).Int())
			} else {
				var _2246_v24 *C2
				_ = _2246_v24
				var _nw387 *C2 = New_C2_()
				_ = _nw387
				_nw387.Ctor__((_2227_v6).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_2227_v6).Cardinality()))).Uint32()).(_dafny.CodePoint), Companion_Default___.Fm0((_this).F6(), globalState))
				_2246_v24 = _nw387
				var _2247_v25 bool
				_ = _2247_v25
				_2247_v25 = true
				_2247_v25 = _2247_v25
				var _2248_v26 _dafny.MultiSet
				_ = _2248_v26
				_2248_v26 = _dafny.MultiSetOf((_this).F6())
				var _2249_v27 _dafny.Sequence
				_ = _2249_v27
				_2249_v27 = _dafny.SeqOf((_this).F6(), _2247_v25, _2247_v25)
				var _2250_v28 _dafny.Array
				_ = _2250_v28
				var _nwElement0_84 _dafny.MultiSet = _2248_v26
				_ = _nwElement0_84
				var _nw388 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_84, nil, _dafny.IntOfInt64(20))
				_ = _nw388
				(_nw388).ArraySet1(_nwElement0_84, 0)
				(_nw388).ArraySet1(_2248_v26, 1)
				(_nw388).ArraySet1(_2248_v26, 2)
				(_nw388).ArraySet1(_2248_v26, 3)
				(_nw388).ArraySet1(_2248_v26, 4)
				(_nw388).ArraySet1(_2248_v26, 5)
				(_nw388).ArraySet1(_2248_v26, 6)
				(_nw388).ArraySet1(_2248_v26, 7)
				(_nw388).ArraySet1(_dafny.MultiSetFromSeq(_2249_v27), 8)
				(_nw388).ArraySet1(_2248_v26, 9)
				(_nw388).ArraySet1(_2248_v26, 10)
				(_nw388).ArraySet1(_2248_v26, 11)
				(_nw388).ArraySet1(_2248_v26, 12)
				(_nw388).ArraySet1(_2248_v26, 13)
				(_nw388).ArraySet1(_2248_v26, 14)
				(_nw388).ArraySet1(_2248_v26, 15)
				(_nw388).ArraySet1(_2248_v26, 16)
				(_nw388).ArraySet1(_2248_v26, 17)
				(_nw388).ArraySet1(_2248_v26, 18)
				(_nw388).ArraySet1(_2248_v26, 19)
				_2250_v28 = _nw388
				var _2251_v29 D18
				_ = _2251_v29
				_2251_v29 = Companion_D18_.Create_DC36_(_2250_v28)
				var _2252_v30 _dafny.Array
				_ = _2252_v30
				var _nwElement0_85 _dafny.Array = (_2251_v29).Dtor_cf45()
				_ = _nwElement0_85
				var _nw389 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_85, nil, _dafny.IntOfInt64(20))
				_ = _nw389
				(_nw389).ArraySet1(_nwElement0_85, 0)
				(_nw389).ArraySet1(_2250_v28, 1)
				(_nw389).ArraySet1(_2250_v28, 2)
				(_nw389).ArraySet1(_2250_v28, 3)
				(_nw389).ArraySet1(_2250_v28, 4)
				(_nw389).ArraySet1(_2250_v28, 5)
				(_nw389).ArraySet1(_2250_v28, 6)
				(_nw389).ArraySet1(_2250_v28, 7)
				(_nw389).ArraySet1((func() _dafny.Array {
					if true {
						return _2250_v28
					}
					return _2250_v28
				})(), 8)
				(_nw389).ArraySet1(_2250_v28, 9)
				(_nw389).ArraySet1(_2250_v28, 10)
				(_nw389).ArraySet1(_2250_v28, 11)
				(_nw389).ArraySet1(_2250_v28, 12)
				(_nw389).ArraySet1(_2250_v28, 13)
				(_nw389).ArraySet1(_2250_v28, 14)
				(_nw389).ArraySet1(_2250_v28, 15)
				(_nw389).ArraySet1(_2250_v28, 16)
				(_nw389).ArraySet1(_2250_v28, 17)
				(_nw389).ArraySet1(_2250_v28, 18)
				(_nw389).ArraySet1(_2250_v28, 19)
				_2252_v30 = _nw389
				var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_2252_v30), 0))
				_ = _index388
				(_2252_v30).ArraySet1(_2250_v28, (_index388).Int())
				var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_2252_v30), 0))
				_ = _index389
				(_2252_v30).ArraySet1(_2250_v28, (_index389).Int())
				_2247_v25 = true
				var _rhs437 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_2247_v25)).Cardinality()), (_dafny.Zero).Minus(p0))
				_ = _rhs437
				var _rhs438 _dafny.Int = _dafny.IntOfInt64(287)
				_ = _rhs438
				var _lhs294 *C7 = _this
				_ = _lhs294
				r1 = _rhs437
				_lhs294.F10 = _rhs438
			}
			var _2253_v31 _dafny.Array
			_ = _2253_v31
			var _nwElement0_86 _dafny.Int = (_this).F9()
			_ = _nwElement0_86
			var _nw390 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_86, nil, _dafny.IntOfInt64(7))
			_ = _nw390
			(_nw390).ArraySet1(_nwElement0_86, 0)
			(_nw390).ArraySet1((_dafny.Zero).Minus((_this).F9()), 1)
			(_nw390).ArraySet1(_dafny.IntOfInt64(792), 2)
			(_nw390).ArraySet1(_dafny.IntOfUint32((_2227_v6).Cardinality()), 3)
			(_nw390).ArraySet1(_this.F10, 4)
			(_nw390).ArraySet1((_this).F9(), 5)
			(_nw390).ArraySet1(_this.F10, 6)
			_2253_v31 = _nw390
			var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_2253_v31), 0))
			_ = _index390
			(_2253_v31).ArraySet1((_this).Fm7((_this).F6(), globalState), (_index390).Int())
			var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(572), _dafny.ArrayLen((_2253_v31), 0))
			_ = _index391
			(_2253_v31).ArraySet1((_dafny.IntOfInt64(837)).Times((_this).F9()), (_index391).Int())
			var _2254_v32 *C5
			_ = _2254_v32
			var _nw391 *C5 = New_C5_()
			_ = _nw391
			_nw391.Ctor__(_dafny.IntOfInt64(109), (_this).F5(), (_this).F6())
			_2254_v32 = _nw391
			_2254_v32 = _2254_v32
			var _2255_v33 *C6
			_ = _2255_v33
			var _nw392 *C6 = New_C6_()
			_ = _nw392
			_nw392.Ctor__()
			_2255_v33 = _nw392
			var _2256_v34 _dafny.Array
			_ = _2256_v34
			var _nw393 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
			_ = _nw393
			_2256_v34 = _nw393
			_2256_v34 = _2256_v34
		}
		var _2257_v35 _dafny.Set
		_ = _2257_v35
		_2257_v35 = _dafny.SetOf(Companion_Default___.Fm0((_this).F6(), globalState))
		_2257_v35 = _2257_v35
		var _2258_v36 _dafny.Map
		_ = _2258_v36
		_2258_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(111))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg120 _dafny.Int) interface{} {
				return coer120(arg120)
			}
		}(func(_2259_i1 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})), (_this).F5())
		var _2260_v37 _dafny.Sequence
		_ = _2260_v37
		_2260_v37 = _dafny.UnicodeSeqOfUtf8Bytes("adh")
		var _2261_v38 _dafny.Sequence
		_ = _2261_v38
		_2261_v38 = _dafny.SeqOf(_2260_v37, _2260_v37)
		_2258_v36 = (_2258_v36).Update((_2261_v38).Select((Companion_Default___.SafeIndex(_this.F10, _dafny.IntOfUint32((_2261_v38).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.CodePoint('k'))
		r1 = p0
		var _2262_v39 _dafny.Array
		_ = _2262_v39
		var _nw394 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw394
		_2262_v39 = _nw394
		var _2263_v40 _dafny.Map
		_ = _2263_v40
		_2263_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
		var _2264_v41 D1
		_ = _2264_v41
		_2264_v41 = Companion_D1_.Create_DC4_((_2263_v40).Cardinality())
		var _2265_v42 _dafny.Map
		_ = _2265_v42
		_2265_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2264_v41, Companion_Default___.Fm0(false, globalState))
		var _2266_v43 D0
		_ = _2266_v43
		var _2267_v44 bool
		_ = _2267_v44
		var _out48 D0
		_ = _out48
		var _out49 bool
		_ = _out49
		_out48, _out49 = Companion_Default___.M0(_2262_v39, (_this).F5(), (func() bool {
			if (_2265_v42).Contains(_2264_v41) {
				return (_2265_v42).Get(_2264_v41).(bool)
			}
			return (_this).F6()
		})(), _dafny.CodePoint('i'), globalState)
		_2266_v43 = _out48
		_2267_v44 = _out49
		r0 = _this.F10
		r1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2260_v37).Cardinality()), _this.F10), Companion_Default___.Fm3((_this).F6(), _2267_v44, globalState))
		var _2268_v45 D6
		_ = _2268_v45
		_2268_v45 = Companion_D6_.Create_DC14_(_2267_v44)
		var _pat_let_tv31 = p0
		_ = _pat_let_tv31
		r2 = _dafny.IntOfUint32((func(_source28 D6) _dafny.Sequence {
			if _source28.Is_DC14() {
				var _2269___mcc_h0 bool = _source28.Get_().(D6_DC14).Cf15
				_ = _2269___mcc_h0
				var _2270_cf15 bool = _2269___mcc_h0
				_ = _2270_cf15
				if true {
					return _dafny.SeqOf((_this).F9(), _this.F10, (_this).F9(), _pat_let_tv31)
				} else {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(571))).Uint32(), func(coer121 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg121 _dafny.Int) interface{} {
							return coer121(arg121)
						}
					}(func(_2271_i2 _dafny.Int) _dafny.Int {
						return _this.F10
					}))
				}
			} else {
				var _2272___mcc_h1 _dafny.Sequence = _source28.Get_().(D6_DC13).Cf14
				_ = _2272___mcc_h1
				var _2273_cf14 _dafny.Sequence = _2272___mcc_h1
				_ = _2273_cf14
				return _2273_cf14
			}
		}(_2268_v45)).Cardinality())
		var _2274_v46 _dafny.Sequence
		_ = _2274_v46
		_2274_v46 = _dafny.SeqOf(_2262_v39)
		r3 = _2274_v46
		return r0, r1, r2, r3
	}
}
func (_this *C7) F9() _dafny.Int {
	{
		return _this._f9
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f5 _dafny.CodePoint
	_f6 bool
	F7  _dafny.Array
	_f8 _dafny.Sequence
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f5 = _dafny.CodePoint('D')
	_this._f6 = false
	_this.F7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C8_ struct {
}

var Companion_C8_ = CompanionStruct_C8_{}

func (_this *C8) Equals(other *C8) bool {
	return _this == other
}

func (_this *C8) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C8)
	return ok && _this.Equals(other)
}

func (*C8) String() string {
	return "_module.C8"
}

func Type_C8_() _dafny.TypeDescriptor {
	return type_C8_{}
}

type type_C8_ struct {
}

func (_this type_C8_) Default() interface{} {
	return (*C8)(nil)
}

func (_this type_C8_) String() string {
	return "main.C8"
}
func (_this *C8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T4_.TraitID_, Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T4 = &C8{}
var _ T3 = &C8{}
var _ T2 = &C8{}
var _ T1 = &C8{}
var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F5() _dafny.CodePoint {
	return _this._f5
}
func (_this *C8) F6() bool {
	return _this._f6
}
func (_this *C8) Ctor__(f7 _dafny.Array, f8 _dafny.Sequence, f5 _dafny.CodePoint, f6 bool) {
	{
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f5 = f5
		(_this)._f6 = f6
	}
}
func (_this *C8) Fm9(p0 _dafny.Int, p1 D0, p2 bool, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((func() _dafny.Set {
			var _coll66 = _dafny.NewBuilder()
			_ = _coll66
			for _iter71 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_((_this).F6()))).Elements()); ; {
				_compr_66, _ok71 := _iter71()
				if !_ok71 {
					break
				}
				var _2275_v0 D0
				_2275_v0 = interface{}(_compr_66).(D0)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC0_(false), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_((_this).F6())), _2275_v0) {
					_coll66.Add(_2275_v0)
				}
			}
			return _coll66.ToSet()
		}()).Union(_dafny.SetOf(Companion_D0_.Create_DC0_(true), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_((_this).F6()), Companion_D0_.Create_DC0_(!((_this).F6()))))).IsDisjointFrom(func() _dafny.Set {
			var _coll67 = _dafny.NewBuilder()
			_ = _coll67
			for _iter72 := _dafny.Iterate((func() _dafny.Set {
				var _coll68 = _dafny.NewBuilder()
				_ = _coll68
				for _iter73 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D0_.Create_DC0_(!((_this).F6())))).Elements()); ; {
					_compr_68, _ok73 := _iter73()
					if !_ok73 {
						break
					}
					var _2276_v1 D0
					_2276_v1 = interface{}(_compr_68).(D0)
					if (_dafny.MultiSetOf(Companion_D0_.Create_DC0_(!((_this).F6())))).Contains(_2276_v1) {
						_coll68.Add(_2276_v1)
					}
				}
				return _coll68.ToSet()
			}()).Elements()); ; {
				_compr_67, _ok72 := _iter72()
				if !_ok72 {
					break
				}
				var _2277_v2 D0
				_2277_v2 = interface{}(_compr_67).(D0)
				if (func() _dafny.Set {
					var _coll69 = _dafny.NewBuilder()
					_ = _coll69
					for _iter74 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D0_.Create_DC0_(!((_this).F6())))).Elements()); ; {
						_compr_69, _ok74 := _iter74()
						if !_ok74 {
							break
						}
						var _2278_v1 D0
						_2278_v1 = interface{}(_compr_69).(D0)
						if (_dafny.MultiSetOf(Companion_D0_.Create_DC0_(!((_this).F6())))).Contains(_2278_v1) {
							_coll69.Add(_2278_v1)
						}
					}
					return _coll69.ToSet()
				}()).Contains(_2277_v2) {
					_coll67.Add(_2277_v2)
				}
			}
			return _coll67.ToSet()
		}())
	}
}
func (_this *C8) Fm10(p0 _dafny.MultiSet, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-966), _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Cmp(((Companion_D2_.Create_DC7_(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-661))).Cardinality()), _dafny.IntOfInt64(-667))).Dtor_cf8()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_this).F6())), (_this).F6())).Cardinality())) != 0
	}
}
func (_this *C8) Fm8(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(434)
	}
}
func (_this *C8) Fm6(p0 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return !((_this).F6())
	}
}
func (_this *C8) Fm7(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var _source29 D2 = Companion_D2_.Create_DC7_(_dafny.IntOfInt64(899), (_dafny.Zero).Minus((_dafny.SetOf(_dafny.SetOf((_this).F6(), (_this).F6()))).Cardinality()))
		_ = _source29
		if _source29.Is_DC7() {
			var _2279___mcc_h0 _dafny.Int = _source29.Get_().(D2_DC7).Cf7
			_ = _2279___mcc_h0
			var _2280___mcc_h1 _dafny.Int = _source29.Get_().(D2_DC7).Cf8
			_ = _2280___mcc_h1
			var _2281_cf8 _dafny.Int = _2280___mcc_h1
			_ = _2281_cf8
			var _2282_cf7 _dafny.Int = _2279___mcc_h0
			_ = _2282_cf7
			return ((_dafny.SetOf((_this).F6())).Union(_dafny.SetOf(!((_this).F6())))).Cardinality()
		} else if _source29.Is_DC8() {
			var _2283___mcc_h2 bool = _source29.Get_().(D2_DC8).Cf9
			_ = _2283___mcc_h2
			var _2284_cf9 bool = _2283___mcc_h2
			_ = _2284_cf9
			return (func() _dafny.Map {
				var _coll70 = _dafny.NewMapBuilder()
				_ = _coll70
				for _iter75 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngprv")).Cardinality()), false)).Cardinality())).Elements()); ; {
					_compr_70, _ok75 := _iter75()
					if !_ok75 {
						break
					}
					var _2285_v0 _dafny.Int
					_2285_v0 = interface{}(_compr_70).(_dafny.Int)
					if (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ngprv")).Cardinality()), false)).Cardinality())).Contains(_2285_v0) {
						_coll70.Add(Companion_Default___.SafeDivisionInt(_2285_v0, _dafny.IntOfInt64(266)), _2284_cf9)
					}
				}
				return _coll70.ToMap()
			}()).Cardinality()
		} else if _source29.Is_DC6() {
			var _2286___mcc_h3 _dafny.Sequence = _source29.Get_().(D2_DC6).Cf6
			_ = _2286___mcc_h3
			var _2287_cf6 _dafny.Sequence = _2286___mcc_h3
			_ = _2287_cf6
			return (_dafny.IntOfInt64(-427)).Plus(_dafny.IntOfInt64(94))
		} else {
			var _2288___mcc_h4 D2 = _source29.Get_().(D2_DC9).Cf10
			_ = _2288___mcc_h4
			var _2289_cf10 D2 = _2288___mcc_h4
			_ = _2289_cf10
			return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(409), _dafny.IntOfInt64(250))
		}
	}
}
func (_this *C8) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!((_this).F6())), _dafny.SeqOf((_this).F6(), (_this).F6())), _dafny.SeqOf(false))).Cardinality())
	}
}
func (_this *C8) Fm5(p0 D2, globalState *GlobalState) D1 {
	{
		return Companion_D1_.Create_DC4_(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(113), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-300))).Cardinality()))).Cardinality()))
	}
}
func (_this *C8) Fm11(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		if (_dafny.SetOf((_this).F6())).IsSubsetOf(_dafny.SetOf((_this).F6(), (_this).F6(), (_this).F6())) {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(609), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-597))).Uint32(), func(coer122 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg122 _dafny.Int) interface{} {
					return coer122(arg122)
				}
			}(func(_2290_i0 _dafny.Int) _dafny.Int {
				return (_dafny.MultiSetOf(true, (_this).F6())).Cardinality()
			}))).Cardinality()))
		} else {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll71 = _dafny.NewBuilder()
				_ = _coll71
				for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(19), _dafny.IntOfInt64(9))); ; {
					_compr_71, _ok76 := _iter76()
					if !_ok76 {
						break
					}
					var _2291_v0 _dafny.Int
					_2291_v0 = interface{}(_compr_71).(_dafny.Int)
					if ((_dafny.IntOfInt64(19)).Cmp(_2291_v0) <= 0) && ((_2291_v0).Cmp(_dafny.IntOfInt64(9)) < 0) {
						_coll71.Add((_2291_v0).Times((_dafny.MultiSetOf(_dafny.IntOfInt64(585))).Cardinality()))
					}
				}
				return _coll71.ToSet()
			}()).Cardinality(), (_dafny.SetOf((_this).F6(), (_this).F6())).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(168), _dafny.IntOfInt64(385)))
		}
	}
}
func (_this *C8) Fm12(globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C8) M4(p0 _dafny.Array, p1 _dafny.Sequence, p2 bool, p3 _dafny.Int, globalState *GlobalState) (D2, _dafny.MultiSet) {
	{
		var r0 D2 = Companion_D2_.Default()
		_ = r0
		var r1 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r1
		var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))
		_ = _index392
		(p0).ArraySet1(p3, (_index392).Int())
		var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))
		_ = _index393
		(p0).ArraySet1(Companion_Default___.Fm3(!(p2), (func() bool {
			if p2 {
				return p2
			}
			return (_this).F6()
		})(), globalState), (_index393).Int())
		var _2292_v0 bool
		_ = _2292_v0
		_2292_v0 = true
		var _2293_v1 _dafny.Map
		_ = _2293_v1
		_2293_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_this).F5())
		var _2294_v2 D0
		_ = _2294_v2
		_2294_v2 = Companion_D0_.Create_DC2_(_2293_v1, (_this).F5())
		var _2295_v3 _dafny.Sequence
		_ = _2295_v3
		_2295_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ghkphwmku")
		var _2296_v4 _dafny.MultiSet
		_ = _2296_v4
		_2296_v4 = _dafny.MultiSetOf(p3, p3)
		var _2297_v5 _dafny.Sequence
		_ = _2297_v5
		_2297_v5 = _dafny.SeqOf(_2296_v4, _2296_v4)
		var _2298_v6 _dafny.MultiSet
		_ = _2298_v6
		_2298_v6 = _dafny.MultiSetOf((_this).F6(), p2, (_2296_v4).IsSubsetOf((_2297_v5).Select((Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2297_v5).Cardinality()))).Uint32()).(_dafny.MultiSet)), (_2292_v0) && ((_this).F6()))
		var _rhs439 bool = _2292_v0
		_ = _rhs439
		var _rhs440 D0 = _2294_v2
		_ = _rhs440
		var _rhs441 _dafny.Int = (_2298_v6).Cardinality()
		_ = _rhs441
		var _rhs442 _dafny.Sequence = _2295_v3
		_ = _rhs442
		var _lhs295 *GlobalState = globalState
		_ = _lhs295
		_2292_v0 = _rhs439
		_2294_v2 = _rhs440
		_lhs295.F1 = _rhs441
		_2295_v3 = _rhs442
		var _2299_v7 *C1
		_ = _2299_v7
		var _nw395 *C1 = New_C1_()
		_ = _nw395
		_nw395.Ctor__()
		_2299_v7 = _nw395
		var _2300_v8 _dafny.Sequence
		_ = _2300_v8
		_2300_v8 = _dafny.SeqOf(p3, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
		var _2301_v9 _dafny.Map
		_ = _2301_v9
		_2301_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _2300_v8)
		var _2302_v10 D6
		_ = _2302_v10
		_2302_v10 = Companion_D6_.Create_DC13_((func() _dafny.Sequence {
			if (_2301_v9).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)) {
				return (_2301_v9).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
			}
			return _2300_v8
		})())
		var _2303_v11 _dafny.Map
		_ = _2303_v11
		_2303_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2302_v10, p3)
		var _hi12 _dafny.Int = Companion_Default___.SafeModuloInt(p3, (_dafny.Zero).Minus((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)))
		_ = _hi12
		for _2304_i0 := (func() _dafny.Int {
			if (_2303_v11).Contains(_2302_v10) {
				return (_2303_v11).Get(_2302_v10).(_dafny.Int)
			}
			return _dafny.IntOfInt64(415)
		})(); _2304_i0.Cmp(_hi12) < 0; _2304_i0 = _2304_i0.Plus(_dafny.One) {
			var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))
			_ = _index394
			(p0).ArraySet1((_2300_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-641), _dafny.IntOfUint32((_2300_v8).Cardinality()))).Uint32()).(_dafny.Int), (_index394).Int())
			var _2305_v12 _dafny.Array
			_ = _2305_v12
			var _len0_59 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_59
			var _nw396 _dafny.Array
			_ = _nw396
			if _len0_59.Cmp(_dafny.Zero) == 0 {
				_nw396 = _dafny.NewArray(_len0_59)
			} else {
				var _init59 func(_dafny.Int) _dafny.CodePoint = func(_2306_i1 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				}
				_ = _init59
				var _element0_59 = _init59(_dafny.Zero)
				_ = _element0_59
				_nw396 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
				(_nw396).ArraySet1CodePoint(_element0_59, 0)
				var _nativeLen0_59 = (_len0_59).Int()
				_ = _nativeLen0_59
				for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
					(_nw396).ArraySet1CodePoint(_init59(_dafny.IntOf(_i0_59)), _i0_59)
				}
			}
			_2305_v12 = _nw396
			var _nw397 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
			_ = _nw397
			_2305_v12 = _nw397
			_2300_v8 = _2300_v8
			var _2307_v13 *C3
			_ = _2307_v13
			var _nw398 *C3 = New_C3_()
			_ = _nw398
			_nw398.Ctor__((_this).F6(), p2)
			_2307_v13 = _nw398
		}
		if _2292_v0 {
			if (p2) && (_2292_v0) {
				var _2308_v14 *C7
				_ = _2308_v14
				var _nw399 *C7 = New_C7_()
				_ = _nw399
				_nw399.Ctor__((_dafny.Zero).Minus(p3), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_this).F5(), _2292_v0)
				_2308_v14 = _nw399
				var _2309_v15 _dafny.CodePoint
				_ = _2309_v15
				_2309_v15 = _dafny.CodePoint('q')
				_2309_v15 = _2309_v15
				var _2310_v16 _dafny.Map
				_ = _2310_v16
				_2310_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_2308_v14).F9())
				var _2311_v17 _dafny.MultiSet
				_ = _2311_v17
				_2311_v17 = _dafny.MultiSetOf(_2295_v3)
				_2292_v0 = (_2311_v17).IsSubsetOf(Companion_Default___.Fm61(_2310_v16, globalState))
				var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))
				_ = _index395
				(p0).ArraySet1(_dafny.IntOfInt64(567), (_index395).Int())
				var _2312_v18 _dafny.Map
				_ = _2312_v18
				_2312_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm12(globalState), _2292_v0)
				_2312_v18 = (_2312_v18).Update(_2292_v0, _2292_v0)
			} else {
				var _2313_v19 _dafny.Array
				_ = _2313_v19
				var _nw400 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(20))
				_ = _nw400
				_2313_v19 = _nw400
				var _2314_v20 *C2
				_ = _2314_v20
				var _nw401 *C2 = New_C2_()
				_ = _nw401
				_nw401.Ctor__(Companion_Default___.Fm22(_dafny.SetOf((_this).F6()), (_dafny.MultiSetOf(true)).Cardinality(), p2, (p1).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), globalState), !(_2292_v0))
				_2314_v20 = _nw401
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2313_v19), 0))
				_ = _index396
				(_2313_v19).ArraySet1(_2314_v20, (_index396).Int())
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(435), _dafny.ArrayLen((_2313_v19), 0))
				_ = _index397
				(_2313_v19).ArraySet1(_2314_v20, (_index397).Int())
				var _2315_v21 _dafny.Array
				_ = _2315_v21
				var _nw402 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(21))
				_ = _nw402
				_2315_v21 = _nw402
				var _2316_v22 _dafny.Map
				_ = _2316_v22
				_2316_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.UnicodeSeqOfUtf8Bytes("lyknjnvny"))
				var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_2315_v21), 0))
				_ = _index398
				(_2315_v21).ArraySet1((_2316_v22).Update(p3, _2295_v3), (_index398).Int())
				var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_2315_v21), 0))
				_ = _index399
				(_2315_v21).ArraySet1((_2316_v22).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _2295_v3)), (_index399).Int())
				var _2317_v23 D11
				_ = _2317_v23
				_2317_v23 = Companion_D11_.Create_DC20_(_2292_v0, p3)
				var _rhs443 bool = p2
				_ = _rhs443
				var _rhs444 bool = (_2317_v23).Dtor_cf21()
				_ = _rhs444
				var _rhs445 _dafny.Int = p3
				_ = _rhs445
				var _lhs296 *GlobalState = globalState
				_ = _lhs296
				_2292_v0 = _rhs443
				_2292_v0 = _rhs444
				_lhs296.F1 = _rhs445
				_2292_v0 = ((_this).F6()) == (_2292_v0)
				var _2318_v24 _dafny.Map
				_ = _2318_v24
				_2318_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2292_v0, (_dafny.IntOfUint32((_2295_v3).Cardinality())).Plus(_dafny.IntOfInt64(-220)))
				_2318_v24 = (_2318_v24).Update((_this).F6(), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
			}
			var _2319_v25 _dafny.Map
			_ = _2319_v25
			_2319_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(p3)), p3)
			_2319_v25 = (_2319_v25).Update((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
			if (_2296_v4).IsProperSubsetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("smcp")).Cardinality()), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_2298_v6).Cardinality())).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F6())).Cardinality(), Companion_Default___.Abs(p3))) {
				_2295_v3 = _2295_v3
				_2292_v0 = false
				var _rhs446 _dafny.Int = (p3).Times(p3)
				_ = _rhs446
				var _rhs447 _dafny.Int = (p3).Plus(((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).Times((_dafny.Zero).Minus(p3)))
				_ = _rhs447
				var _lhs297 *GlobalState = globalState
				_ = _lhs297
				var _lhs298 *GlobalState = globalState
				_ = _lhs298
				_lhs297.F1 = _rhs446
				_lhs298.F1 = _rhs447
				var _2320_v26 _dafny.Map
				_ = _2320_v26
				_2320_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F6())
				_2320_v26 = _2320_v26
				(globalState).F1 = p3
			} else {
				var _2321_v27 _dafny.Array
				_ = _2321_v27
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_60
				var _nw403 _dafny.Array
				_ = _nw403
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw403 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) _dafny.Int = (func(_2322_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2323_i2 _dafny.Int) _dafny.Int {
							return (_2323_i2).Minus(_2322_p3)
						}
					})(p3)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw403 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw403).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw403).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_2321_v27 = _nw403
				var _rhs448 _dafny.Int = (p3).Times(_dafny.IntOfUint32(((Companion_D6_.Create_DC13_(_2300_v8)).Dtor_cf14()).Cardinality()))
				_ = _rhs448
				var _rhs449 _dafny.Array = _this.F7
				_ = _rhs449
				var _rhs450 bool = (_this).F6()
				_ = _rhs450
				var _rhs451 _dafny.Array = p0
				_ = _rhs451
				var _lhs299 *GlobalState = globalState
				_ = _lhs299
				var _lhs300 *C8 = _this
				_ = _lhs300
				_lhs299.F1 = _rhs448
				_lhs300.F7 = _rhs449
				_2292_v0 = _rhs450
				_2321_v27 = _rhs451
				_2292_v0 = (_this).F6()
				var _2324_v28 _dafny.Set
				_ = _2324_v28
				_2324_v28 = _dafny.SetOf(_2292_v0)
				var _2325_v29 _dafny.Map
				_ = _2325_v29
				_2325_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2324_v28).Union(_2324_v28)).Cardinality(), _2292_v0)
				_2325_v29 = (_2325_v29).Update((_dafny.IntOfInt64(-999)).Minus(p3), p2)
				var _2326_v30 _dafny.Set
				_ = _2326_v30
				_2326_v30 = _dafny.SetOf(_dafny.IntOfUint32((Companion_Default___.Fm1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), !(p2), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))
				_2326_v30 = _2326_v30
				_2292_v0 = p2
			}
			var _2327_v31 _dafny.CodePoint
			_ = _2327_v31
			_2327_v31 = _dafny.CodePoint('g')
			_2327_v31 = (_this).F5()
			_2292_v0 = _2292_v0
		} else {
			var _2328_v32 *C6
			_ = _2328_v32
			var _nw404 *C6 = New_C6_()
			_ = _nw404
			_nw404.Ctor__()
			_2328_v32 = _nw404
			var _2329_v33 D14
			_ = _2329_v33
			_2329_v33 = Companion_D14_.Create_DC26_((_this).F6(), p3, Companion_Default___.Fm53(_2295_v3, (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), globalState))
			(globalState).F1 = (_2300_v8).Select((Companion_Default___.SafeIndex((_2329_v33).Dtor_cf33(), _dafny.IntOfUint32((_2300_v8).Cardinality()))).Uint32()).(_dafny.Int)
			_2292_v0 = (_this).F6()
			var _2330_v34 _dafny.Set
			_ = _2330_v34
			_2330_v34 = _dafny.SetOf(p2)
			var _2331_v35 _dafny.Map
			_ = _2331_v35
			_2331_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, (_2330_v34).IsDisjointFrom(_2330_v34))
			var _2332_v36 _dafny.Sequence
			_ = _2332_v36
			_2332_v36 = _dafny.SeqOf(_2331_v35, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7, p2))
			var _rhs452 _dafny.Map = ((_2331_v35).Update(_this.F7, (_this).F6())).Merge((_2332_v36).Select((Companion_Default___.SafeIndex((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2332_v36).Cardinality()))).Uint32()).(_dafny.Map))
			_ = _rhs452
			var _rhs453 bool = false
			_ = _rhs453
			var _rhs454 bool = _2292_v0
			_ = _rhs454
			var _rhs455 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p2)).Update(Companion_Default___.Fm0((_this).F6(), globalState), (p1).Select((Companion_Default___.SafeIndex(p3, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool))).Cardinality()
			_ = _rhs455
			var _lhs301 *GlobalState = globalState
			_ = _lhs301
			_2331_v35 = _rhs452
			_2292_v0 = _rhs453
			_2292_v0 = _rhs454
			_lhs301.F1 = _rhs455
			var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))
			_ = _index400
			(p0).ArraySet1((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), (_index400).Int())
		}
		var _2333_v37 D11
		_ = _2333_v37
		_2333_v37 = Companion_D11_.Create_DC20_(!(true), (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
		var _2334_v38 _dafny.Map
		_ = _2334_v38
		_2334_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2300_v8, (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ymtkf")).Cardinality())).Cmp((_2333_v37).Dtor_cf22()) >= 0)
		(globalState).F1 = (_2334_v38).Cardinality()
		var _2335_v39 _dafny.Map
		_ = _2335_v39
		_2335_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
		var _2336_v40 _dafny.Map
		_ = _2336_v40
		_2336_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_2335_v39).Contains(!(p2)) {
				return (_2335_v39).Get(!(p2)).(_dafny.Int)
			}
			return (p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)
		})(), _2292_v0)
		r0 = Companion_Default___.Fm62(Companion_D15_.Create_DC27_(_dafny.MultiSetOf(false, p2)), (func() bool {
			if (_2336_v40).Contains(p3) {
				return (_2336_v40).Get(p3).(bool)
			}
			return _2292_v0
		})(), _2292_v0, _dafny.Companion_Sequence_.Concatenate(_2300_v8, _2300_v8), globalState)
		r1 = _2298_v6
		return r0, r1
	}
}
func (_this *C8) M5(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _2337_v0 _dafny.Sequence
		_ = _2337_v0
		_2337_v0 = _dafny.UnicodeSeqOfUtf8Bytes("duo")
		_2337_v0 = p0
		var _2338_v1 bool
		_ = _2338_v1
		_2338_v1 = false
		_2338_v1 = _2338_v1
		var _2339_v2 _dafny.Map
		_ = _2339_v2
		_2339_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2338_v1)
		var _source30 _dafny.Map = (_2339_v2).Merge(_2339_v2)
		_ = _source30
		var _2340___mcc_h0 _dafny.Map = _source30
		_ = _2340___mcc_h0
		var _2341_cf12 _dafny.Map = _2340___mcc_h0
		_ = _2341_cf12
		_2338_v1 = _2338_v1
		(globalState).F1 = p1
		var _2342_v3 _dafny.Map
		_ = _2342_v3
		_2342_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p1, (_dafny.SetOf(_2338_v1)).Cardinality()), p1)
		_2342_v3 = (_2342_v3).Update(_dafny.IntOfInt64(133), (_dafny.Zero).Minus(p1))
		var _2343_v4 _dafny.Array
		_ = _2343_v4
		_2343_v4 = _this.F7
		var _2344_v5 _dafny.Map
		_ = _2344_v5
		_2344_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2343_v4, _this.F7)
		var _2345_v6 _dafny.Map
		_ = _2345_v6
		_2345_v6 = _2344_v5
		var _source31 _dafny.Map = _2345_v6
		_ = _source31
		var _2346___mcc_h1 _dafny.Map = _source31
		_ = _2346___mcc_h1
		var _2347_cf30 _dafny.Map = _2346___mcc_h1
		_ = _2347_cf30
		var _rhs456 _dafny.Sequence = p0
		_ = _rhs456
		var _rhs457 _dafny.Map = (_2342_v3).Merge((_2342_v3).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfInt64(-500))).Update(p1, p1)))
		_ = _rhs457
		_2337_v0 = _rhs456
		_2342_v3 = _rhs457
		_2338_v1 = (_this).F6()
		var _2348_v7 _dafny.Array
		_ = _2348_v7
		var _len0_61 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_61
		var _nw405 _dafny.Array
		_ = _nw405
		if _len0_61.Cmp(_dafny.Zero) == 0 {
			_nw405 = _dafny.NewArray(_len0_61)
		} else {
			var _init61 func(_dafny.Int) _dafny.Int = (func(_2349_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2350_i0 _dafny.Int) _dafny.Int {
					return (_2350_i0).Minus(_2349_p1)
				}
			})(p1)
			_ = _init61
			var _element0_61 = _init61(_dafny.Zero)
			_ = _element0_61
			_nw405 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
			(_nw405).ArraySet1(_element0_61, 0)
			var _nativeLen0_61 = (_len0_61).Int()
			_ = _nativeLen0_61
			for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
				(_nw405).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
			}
		}
		_2348_v7 = _nw405
		var _2351_v8 _dafny.Sequence
		_ = _2351_v8
		_2351_v8 = _dafny.SeqOf((_this).F6(), !(_2338_v1), (_this).F6(), (_this).F6(), true)
		var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_2348_v7), 0))
		_ = _index401
		(_2348_v7).ArraySet1((_dafny.IntOfUint32((_2351_v8).Cardinality())).Times(p1), (_index401).Int())
		var _arr2 _dafny.Array = _this.F7
		_ = _arr2
		var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_this.F7), 0))
		_ = _index402
		(_arr2).ArraySet1(false, (_index402).Int())
		var _2352_v10 _dafny.Sequence
		_ = _2352_v10
		_2352_v10 = _dafny.SeqOf(p1)
		var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_2348_v7), 0))
		_ = _index403
		var _arr3 _dafny.Array = _this.F7
		_ = _arr3
		var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_this.F7), 0))
		_ = _index404
		var _rhs458 _dafny.Int = p1
		_ = _rhs458
		var _rhs459 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
		_ = _rhs459
		var _rhs460 _dafny.Int = (Companion_Default___.SafeDivisionInt(p1, (_dafny.Zero).Minus((func() _dafny.Map {
			var _coll72 = _dafny.NewMapBuilder()
			_ = _coll72
			for _iter77 := _dafny.Iterate((_2352_v10).Elements()); ; {
				_compr_72, _ok77 := _iter77()
				if !_ok77 {
					break
				}
				var _2353_v9 _dafny.Int
				_2353_v9 = interface{}(_compr_72).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_2352_v10, _2353_v9) {
					_coll72.Add(Companion_Default___.SafeDivisionInt(_2353_v9, p1), p1)
				}
			}
			return _coll72.ToMap()
		}()).Cardinality()))).Plus(p1)
		_ = _rhs460
		var _rhs461 bool = (_this).F6()
		_ = _rhs461
		var _lhs302 *GlobalState = globalState
		_ = _lhs302
		var _lhs303 _dafny.Array = _2348_v7
		_ = _lhs303
		var _lhs304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(221), _dafny.ArrayLen((_2348_v7), 0))
		_ = _lhs304
		var _lhs305 *GlobalState = globalState
		_ = _lhs305
		var _lhs306 _dafny.Array = _this.F7
		_ = _lhs306
		var _lhs307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_this.F7), 0))
		_ = _lhs307
		_lhs302.F1 = _rhs458
		(_lhs303).ArraySet1(_rhs459, (_lhs304).Int())
		_lhs305.F1 = _rhs460
		(_lhs306).ArraySet1(_rhs461, (_lhs307).Int())
		var _2354_v11 _dafny.Map
		_ = _2354_v11
		_2354_v11 = _2342_v3
		var _2355_v12 _dafny.Sequence
		_ = _2355_v12
		_2355_v12 = _dafny.SeqOf(_2351_v8, _2351_v8)
		var _2356_v13 D12
		_ = _2356_v13
		_2356_v13 = Companion_D12_.Create_DC22_(_2354_v11, _2355_v12, (_this).F5(), _2338_v1, _2337_v0)
		var _2357_v14 _dafny.Sequence
		_ = _2357_v14
		_2357_v14 = _dafny.SeqOf((_2356_v13).Dtor_cf28(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aldrxfvuv"), p0), _2337_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(461))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg123 _dafny.Int) interface{} {
				return coer123(arg123)
			}
		}(func(_2358_i1 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})), _dafny.UnicodeSeqOfUtf8Bytes("toj"))
		_2357_v14 = _2357_v14
		_2338_v1 = !(_2338_v1)
		(globalState).F1 = ((Companion_D19_.Create_DC38_(Companion_Default___.Fm63(p1, globalState))).Dtor_cf49()).Cardinality()
		var _2359_v15 _dafny.Array
		_ = _2359_v15
		var _len0_62 _dafny.Int = _dafny.IntOfInt64(3)
		_ = _len0_62
		var _nw406 _dafny.Array
		_ = _nw406
		if _len0_62.Cmp(_dafny.Zero) == 0 {
			_nw406 = _dafny.NewArray(_len0_62)
		} else {
			var _init62 func(_dafny.Int) _dafny.Int = func(_2360_i3 _dafny.Int) _dafny.Int {
				return (_2360_i3).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))
			}
			_ = _init62
			var _element0_62 = _init62(_dafny.Zero)
			_ = _element0_62
			_nw406 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
			(_nw406).ArraySet1(_element0_62, 0)
			var _nativeLen0_62 = (_len0_62).Int()
			_ = _nativeLen0_62
			for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
				(_nw406).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
			}
		}
		_2359_v15 = _nw406
		for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2359_v15), 0))); ; {
			_guard_loop_5, _ok78 := _iter78()
			if !_ok78 {
				break
			}
			var _2361_i2 _dafny.Int
			_2361_i2 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_2361_i2).Sign() != -1) && ((_2361_i2).Cmp(_dafny.ArrayLen((_2359_v15), 0)) < 0)) {
				(_2359_v15).ArraySet1((_2361_i2).Plus((Companion_D14_.Create_DC26_((_this).F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), false)).Cardinality(), (_this).F5())).Dtor_cf33()), (_2361_i2).Int())
			}
		}
	}
}
func (_this *C8) M2(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Sequence) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _2362_v0 _dafny.MultiSet
		_ = _2362_v0
		_2362_v0 = _dafny.MultiSetOf((_this).F6(), (_this).F6(), (_this).F6())
		var _hi13 _dafny.Int = p1
		_ = _hi13
		for _2363_i0 := (_2362_v0).Cardinality(); _2363_i0.Cmp(_hi13) < 0; _2363_i0 = _2363_i0.Plus(_dafny.One) {
			var _2364_v1 _dafny.Array
			_ = _2364_v1
			var _len0_63 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_63
			var _nw407 _dafny.Array
			_ = _nw407
			if _len0_63.Cmp(_dafny.Zero) == 0 {
				_nw407 = _dafny.NewArray(_len0_63)
			} else {
				var _init63 func(_dafny.Int) _dafny.Int = (func(_2365_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2366_i1 _dafny.Int) _dafny.Int {
						return (_2366_i1).Minus(_2365_p1)
					}
				})(p1)
				_ = _init63
				var _element0_63 = _init63(_dafny.Zero)
				_ = _element0_63
				_nw407 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
				(_nw407).ArraySet1(_element0_63, 0)
				var _nativeLen0_63 = (_len0_63).Int()
				_ = _nativeLen0_63
				for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
					(_nw407).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
				}
			}
			_2364_v1 = _nw407
			var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))
			_ = _index405
			(_2364_v1).ArraySet1(_2363_i0, (_index405).Int())
			var _2367_v2 _dafny.Map
			_ = _2367_v2
			_2367_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(815))
			var _2368_v3 D0
			_ = _2368_v3
			_2368_v3 = Companion_D0_.Create_DC0_((_this).F6())
			var _2369_v5 _dafny.Map
			_ = _2369_v5
			_2369_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_2363_i0)).Cardinality()), p1))
			var _2370_v6 _dafny.Sequence
			_ = _2370_v6
			_2370_v6 = _dafny.SeqOf(p1, _dafny.IntOfUint32((p2).Cardinality()))
			var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))
			_ = _index406
			var _rhs462 _dafny.Int = (p1).Plus((func() _dafny.Int {
				if (_2367_v2).Contains(p0) {
					return (_2367_v2).Get(p0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p1)
			})())
			_ = _rhs462
			var _rhs463 bool = Companion_Default___.Fm0((_this).Fm9(p1, _2368_v3, (_this).F6(), p1, globalState), globalState)
			_ = _rhs463
			var _rhs464 bool = ((func() _dafny.Int {
				if (_2362_v0).Contains(p0) {
					return (_2362_v0).Multiplicity(p0)
				}
				return Companion_Default___.Fm3((_this).F6(), (_this).F6(), globalState)
			})()).Cmp(_2363_i0) >= 0
			_ = _rhs464
			var _rhs465 _dafny.Int = Companion_Default___.SafeDivisionInt(((func() _dafny.Map {
				var _coll73 = _dafny.NewMapBuilder()
				_ = _coll73
				for _iter79 := _dafny.Iterate(((func() _dafny.Sequence {
					if (_2369_v5).Contains((_this).F6()) {
						return (_2369_v5).Get((_this).F6()).(_dafny.Sequence)
					}
					return _2370_v6
				})()).Elements()); ; {
					_compr_73, _ok79 := _iter79()
					if !_ok79 {
						break
					}
					var _2371_v4 _dafny.Int
					_2371_v4 = interface{}(_compr_73).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
						if (_2369_v5).Contains((_this).F6()) {
							return (_2369_v5).Get((_this).F6()).(_dafny.Sequence)
						}
						return _2370_v6
					})(), _2371_v4) {
						_coll73.Add((_2371_v4).Times(_2363_i0), (_2367_v2).Cardinality())
					}
				}
				return _coll73.ToMap()
			}()).Cardinality()).Plus(p1), (_2363_i0).Times(p1))
			_ = _rhs465
			var _lhs308 *GlobalState = globalState
			_ = _lhs308
			var _lhs309 _dafny.Array = _2364_v1
			_ = _lhs309
			var _lhs310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))
			_ = _lhs310
			_lhs308.F1 = _rhs462
			r0 = _rhs463
			r0 = _rhs464
			(_lhs309).ArraySet1(_rhs465, (_lhs310).Int())
			var _2372_v7 _dafny.Sequence
			_ = _2372_v7
			_2372_v7 = _dafny.SeqOf(p0)
			var _rhs466 bool = (_this).F6()
			_ = _rhs466
			var _rhs467 _dafny.MultiSet = _2362_v0
			_ = _rhs467
			var _rhs468 bool = (!((_this).F6())) == ((_2372_v7).Select((Companion_Default___.SafeIndex((_2364_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2372_v7).Cardinality()))).Uint32()).(bool))
			_ = _rhs468
			var _rhs469 _dafny.Array = _this.F7
			_ = _rhs469
			var _lhs311 *GlobalState = globalState
			_ = _lhs311
			var _lhs312 *C8 = _this
			_ = _lhs312
			r0 = _rhs466
			_lhs311.F4 = _rhs467
			r0 = _rhs468
			_lhs312.F7 = _rhs469
			r0 = (_this).F6()
			var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))
			_ = _index407
			(_2364_v1).ArraySet1((_2364_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_2364_v1), 0))).Int()).(_dafny.Int), (_index407).Int())
		}
		var _2373_i2 _dafny.Int
		_ = _2373_i2
		_2373_i2 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_2373_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L22
					}
					_2373_i2 = (_2373_i2).Plus(_dafny.One)
					(globalState).F1 = p1
					var _2374_v8 _dafny.Map
					_ = _2374_v8
					_2374_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F5())
					var _2375_v9 D11
					_ = _2375_v9
					_2375_v9 = Companion_D11_.Create_DC20_((_this).F6(), (_2374_v8).Cardinality())
					var _2376_v10 _dafny.Array
					_ = _2376_v10
					var _nwElement0_87 _dafny.Int = (_2375_v9).Dtor_cf22()
					_ = _nwElement0_87
					var _nw408 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_87, nil, _dafny.IntOfInt64(12))
					_ = _nw408
					(_nw408).ArraySet1(_nwElement0_87, 0)
					(_nw408).ArraySet1((_dafny.Zero).Minus(p1), 1)
					(_nw408).ArraySet1(p1, 2)
					(_nw408).ArraySet1((_2362_v0).Cardinality(), 3)
					(_nw408).ArraySet1(_dafny.IntOfInt64(544), 4)
					(_nw408).ArraySet1(_dafny.IntOfUint32((p2).Cardinality()), 5)
					(_nw408).ArraySet1(p1, 6)
					(_nw408).ArraySet1(p1, 7)
					(_nw408).ArraySet1(_dafny.IntOfInt64(669), 8)
					(_nw408).ArraySet1((_dafny.Zero).Minus(p1), 9)
					(_nw408).ArraySet1(p1, 10)
					(_nw408).ArraySet1(p1, 11)
					_2376_v10 = _nw408
					var _2377_v11 *C4
					_ = _2377_v11
					var _nw409 *C4 = New_C4_()
					_ = _nw409
					_nw409.Ctor__()
					_2377_v11 = _nw409
					var _2378_v12 _dafny.MultiSet
					_ = _2378_v12
					_2378_v12 = _dafny.MultiSetOf(_2377_v11, _2377_v11, _2377_v11, _2377_v11)
					var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2376_v10), 0))
					_ = _index408
					(_2376_v10).ArraySet1((func() _dafny.Int {
						if (_2378_v12).Contains(_2377_v11) {
							return (_2378_v12).Multiplicity(_2377_v11)
						}
						return _dafny.IntOfUint32((p2).Cardinality())
					})(), (_index408).Int())
					var _2379_v13 _dafny.Map
					_ = _2379_v13
					_2379_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p2)
					var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(913), _dafny.ArrayLen((_2376_v10), 0))
					_ = _index409
					(_2376_v10).ArraySet1(((_2379_v13).Cardinality()).Minus(p1), (_index409).Int())
					var _2380_v14 _dafny.Array
					_ = _2380_v14
					var _nw410 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
					_ = _nw410
					_2380_v14 = _nw410
					var _2381_v15 _dafny.Array
					_ = _2381_v15
					var _nw411 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(29))
					_ = _nw411
					_2381_v15 = _nw411
					var _2382_v16 _dafny.Array
					_ = _2382_v16
					var _nwElement0_88 _dafny.Array = _2381_v15
					_ = _nwElement0_88
					var _nw412 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_88, nil, _dafny.IntOfInt64(17))
					_ = _nw412
					(_nw412).ArraySet1(_nwElement0_88, 0)
					(_nw412).ArraySet1(_2381_v15, 1)
					(_nw412).ArraySet1(_2381_v15, 2)
					(_nw412).ArraySet1(_2381_v15, 3)
					(_nw412).ArraySet1(_2381_v15, 4)
					(_nw412).ArraySet1(_2381_v15, 5)
					(_nw412).ArraySet1(_2381_v15, 6)
					(_nw412).ArraySet1(_2381_v15, 7)
					(_nw412).ArraySet1(_2381_v15, 8)
					(_nw412).ArraySet1(_2381_v15, 9)
					(_nw412).ArraySet1(_2381_v15, 10)
					(_nw412).ArraySet1(_2381_v15, 11)
					(_nw412).ArraySet1(_2381_v15, 12)
					(_nw412).ArraySet1(_2381_v15, 13)
					(_nw412).ArraySet1(_2381_v15, 14)
					(_nw412).ArraySet1(_2381_v15, 15)
					(_nw412).ArraySet1(_2381_v15, 16)
					_2382_v16 = _nw412
					var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_2380_v14), 0))
					_ = _index410
					(_2380_v14).ArraySet1(_2382_v16, (_index410).Int())
					var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_2380_v14), 0))
					_ = _index411
					(_2380_v14).ArraySet1(_2382_v16, (_index411).Int())
					var _2383_v17 T1
					_ = _2383_v17
					var _nw413 *C6 = New_C6_()
					_ = _nw413
					_nw413.Ctor__()
					_2383_v17 = _nw413
					_2383_v17 = _2383_v17
					goto C22
				}
			C22:
			}
			goto L22
		}
	L22:
		var _2384_v18 _dafny.Sequence
		_ = _2384_v18
		_2384_v18 = _dafny.SeqOf(p2, _dafny.UnicodeSeqOfUtf8Bytes("s"))
		var _2385_i3 _dafny.Int
		_ = _2385_i3
		_2385_i3 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.IsProperPrefixOf(_2384_v18, _dafny.SeqOf(p2)) {
				{
					if (_2385_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L23
					}
					_2385_i3 = (_2385_i3).Plus(_dafny.One)
					var _2386_v19 *C6
					_ = _2386_v19
					var _nw414 *C6 = New_C6_()
					_ = _nw414
					_nw414.Ctor__()
					_2386_v19 = _nw414
					var _2387_v20 _dafny.Map
					_ = _2387_v20
					_2387_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2386_v19, p0)
					var _2388_v22 D14
					_ = _2388_v22
					_2388_v22 = Companion_D14_.Create_DC26_((_this).F6(), p1, Companion_Default___.Fm36(p1, p1, (_2387_v20).Cardinality(), func() _dafny.Map {
						var _coll74 = _dafny.NewMapBuilder()
						_ = _coll74
						for _iter80 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(137), _dafny.IntOfInt64(471))); ; {
							_compr_74, _ok80 := _iter80()
							if !_ok80 {
								break
							}
							var _2389_v21 _dafny.Int
							_2389_v21 = interface{}(_compr_74).(_dafny.Int)
							if ((_dafny.IntOfInt64(137)).Cmp(_2389_v21) <= 0) && ((_2389_v21).Cmp(_dafny.IntOfInt64(471)) < 0) {
								_coll74.Add(Companion_Default___.SafeModuloInt(_2389_v21, p1), p1)
							}
						}
						return _coll74.ToMap()
					}(), globalState))
					var _source32 D14 = _2388_v22
					_ = _source32
					if _source32.Is_DC26() {
						var _2390___mcc_h0 bool = _source32.Get_().(D14_DC26).Cf32
						_ = _2390___mcc_h0
						var _2391___mcc_h1 _dafny.Int = _source32.Get_().(D14_DC26).Cf33
						_ = _2391___mcc_h1
						var _2392___mcc_h2 _dafny.CodePoint = _source32.Get_().(D14_DC26).Cf34
						_ = _2392___mcc_h2
						var _2393_cf34 _dafny.CodePoint = _2392___mcc_h2
						_ = _2393_cf34
						var _2394_cf33 _dafny.Int = _2391___mcc_h1
						_ = _2394_cf33
						var _2395_cf32 bool = _2390___mcc_h0
						_ = _2395_cf32
						var _2396_v23 _dafny.Map
						_ = _2396_v23
						_2396_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
						_2396_v23 = (_2396_v23).Update((_2394_cf33).Cmp((func() _dafny.Int {
							if (_2362_v0).Contains(false) {
								return (_2362_v0).Multiplicity(false)
							}
							return p1
						})()) >= 0, p1)
						(globalState).F1 = _dafny.IntOfInt64(209)
						var _2397_v24 _dafny.Map
						_ = _2397_v24
						_2397_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2395_cf32, p2)
						_2394_cf33 = ((_2394_cf33).Times(p1)).Times(((_2397_v24).Cardinality()).Times(p1))
						var _2398_v25 _dafny.Map
						_ = _2398_v25
						_2398_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
						var _2399_v26 _dafny.Sequence
						_ = _2399_v26
						_2399_v26 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("inqarh")).Cardinality()), (_2398_v25).Cardinality())
						var _2400_v27 _dafny.MultiSet
						_ = _2400_v27
						_2400_v27 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_2399_v26, _2399_v26), _dafny.SeqOf(p1), _2399_v26, _2399_v26)
						_2400_v27 = (_2400_v27).Difference(_2400_v27)
					} else {
						var _2401___mcc_h3 *C4 = _source32.Get_().(D14_DC25).Cf31
						_ = _2401___mcc_h3
						var _2402_cf31 *C4 = _2401___mcc_h3
						_ = _2402_cf31
						var _2403_v28 _dafny.Array
						_ = _2403_v28
						var _nw415 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(6))
						_ = _nw415
						_2403_v28 = _nw415
						var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_2403_v28), 0))
						_ = _index412
						(_2403_v28).ArraySet1CodePoint(_dafny.CodePoint('q'), (_index412).Int())
						var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_2403_v28), 0))
						_ = _index413
						(_2403_v28).ArraySet1CodePoint(_dafny.CodePoint('t'), (_index413).Int())
						var _2404_v29 _dafny.Array
						_ = _2404_v29
						var _len0_64 _dafny.Int = _dafny.IntOfInt64(28)
						_ = _len0_64
						var _nw416 _dafny.Array
						_ = _nw416
						if _len0_64.Cmp(_dafny.Zero) == 0 {
							_nw416 = _dafny.NewArray(_len0_64)
						} else {
							var _init64 func(_dafny.Int) _dafny.Int = (func(_2405_p0 bool) func(_dafny.Int) _dafny.Int {
								return func(_2406_i4 _dafny.Int) _dafny.Int {
									return (_2406_i4).Times((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_2405_p0, (_this).F6())).Cardinality()), _dafny.IntOfInt64(218))).Cardinality())
								}
							})(p0)
							_ = _init64
							var _element0_64 = _init64(_dafny.Zero)
							_ = _element0_64
							_nw416 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
							(_nw416).ArraySet1(_element0_64, 0)
							var _nativeLen0_64 = (_len0_64).Int()
							_ = _nativeLen0_64
							for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
								(_nw416).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
							}
						}
						_2404_v29 = _nw416
						var _2407_v30 _dafny.Set
						_ = _2407_v30
						_2407_v30 = _dafny.SetOf((_this).F6(), (_this).F6())
						var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2404_v29), 0))
						_ = _index414
						(_2404_v29).ArraySet1(((_2407_v30).Cardinality()).Minus((_dafny.Zero).Minus(p1)), (_index414).Int())
						var _2408_v31 _dafny.Sequence
						_ = _2408_v31
						_2408_v31 = _dafny.SeqOf((_this).F6())
						var _2409_v32 _dafny.Map
						_ = _2409_v32
						_2409_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
						var _2410_v33 _dafny.Map
						_ = _2410_v33
						_2410_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
						var _2411_v34 _dafny.MultiSet
						_ = _2411_v34
						_2411_v34 = _dafny.MultiSetOf(p1, (_2410_v33).Cardinality())
						var _2412_v35 _dafny.Map
						_ = _2412_v35
						_2412_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((p2).Cardinality()), (_2386_v19).Fm4(_2408_v31, _dafny.IntOfInt64(306), (_2409_v32).Cardinality(), _2411_v34, globalState))
						var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_2404_v29), 0))
						_ = _index415
						(_2404_v29).ArraySet1(((func() _dafny.Int {
							if (_2412_v35).Contains(p1) {
								return (_2412_v35).Get(p1).(_dafny.Int)
							}
							return p1
						})()).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(30), p1)), (_index415).Int())
						var _2413_v36 *C2
						_ = _2413_v36
						var _nw417 *C2 = New_C2_()
						_ = _nw417
						_nw417.Ctor__((func() _dafny.CodePoint {
							if p0 {
								return _dafny.CodePoint('l')
							}
							return _dafny.CodePoint('c')
						})(), (_this).F6())
						_2413_v36 = _nw417
						var _2414_v37 D20
						_ = _2414_v37
						_2414_v37 = Companion_D20_.Create_DC42_(_2386_v19)
						var _2415_v38 _dafny.Map
						_ = _2415_v38
						_2415_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_2413_v36).Fm8(true, (_this).F6(), p0, globalState)).Plus(p1), (_2414_v37).Dtor_cf55())
						_2415_v38 = (_2415_v38).Update((_2386_v19).Fm4(_2408_v31, p1, p1, _2411_v34, globalState), _2386_v19)
					}
					var _2416_v39 _dafny.Map
					_ = _2416_v39
					_2416_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!(p0), p0, (_this).F6(), (_this).F6(), (_this).F6())).Cardinality(), p1)
					var _2417_v40 _dafny.Set
					_ = _2417_v40
					_2417_v40 = _dafny.SetOf(p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)).Cardinality())
					var _2418_v41 _dafny.Map
					_ = _2418_v41
					_2418_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(155), _dafny.SeqOf(p1))
					var _2419_v42 D19
					_ = _2419_v42
					_2419_v42 = Companion_D19_.Create_DC40_(p0, (_2417_v40).Cardinality(), _2418_v41)
					_2416_v39 = (_2416_v39).Update(p1, (_2419_v42).Dtor_cf52())
					var _2420_v43 _dafny.Array
					_ = _2420_v43
					var _nw418 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
					_ = _nw418
					_2420_v43 = _nw418
					var _2421_v44 _dafny.Array
					_ = _2421_v44
					var _nwElement0_89 _dafny.Int = p1
					_ = _nwElement0_89
					var _nw419 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_89, nil, _dafny.One)
					_ = _nw419
					(_nw419).ArraySet1(_nwElement0_89, 0)
					_2421_v44 = _nw419
					var _index416 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_2420_v43), 0))
					_ = _index416
					(_2420_v43).ArraySet1(_2421_v44, (_index416).Int())
					var _index417 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_2421_v44), 0))
					_ = _index417
					(_2421_v44).ArraySet1((_dafny.IntOfInt64(797)).Times(p1), (_index417).Int())
					var _2422_v45 _dafny.Sequence
					_ = _2422_v45
					_2422_v45 = _dafny.SeqOf(p0, (_this).F6(), p0)
					var _index418 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_2420_v43), 0))
					_ = _index418
					var _index419 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_2421_v44), 0))
					_ = _index419
					var _rhs470 _dafny.Array = _2421_v44
					_ = _rhs470
					var _rhs471 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(946))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg124 _dafny.Int) interface{} {
							return coer124(arg124)
						}
					}(func(_2423_i5 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					})), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vhbetwjr"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(443))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg125 _dafny.Int) interface{} {
							return coer125(arg125)
						}
					}(func(_2424_i6 _dafny.Int) _dafny.CodePoint {
						return (_this).F5()
					}))))
					_ = _rhs471
					var _rhs472 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p1), p1)
					_ = _rhs472
					var _rhs473 _dafny.Sequence = _2422_v45
					_ = _rhs473
					var _lhs313 _dafny.Array = _2420_v43
					_ = _lhs313
					var _lhs314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_2420_v43), 0))
					_ = _lhs314
					var _lhs315 _dafny.Array = _2421_v44
					_ = _lhs315
					var _lhs316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_2421_v44), 0))
					_ = _lhs316
					(_lhs313).ArraySet1(_rhs470, (_lhs314).Int())
					r0 = _rhs471
					(_lhs315).ArraySet1(_rhs472, (_lhs316).Int())
					_2422_v45 = _rhs473
					var _arr4 _dafny.Array = _this.F7
					_ = _arr4
					var _index420 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_this.F7), 0))
					_ = _index420
					(_arr4).ArraySet1((func() bool {
						if p0 {
							return p0
						}
						return (_2422_v45).Select((Companion_Default___.SafeIndex((_2421_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(16), _dafny.ArrayLen((_2421_v44), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2422_v45).Cardinality()))).Uint32()).(bool)
					})(), (_index420).Int())
					var _arr5 _dafny.Array = _this.F7
					_ = _arr5
					var _index421 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(369), _dafny.ArrayLen((_this.F7), 0))
					_ = _index421
					(_arr5).ArraySet1(p0, (_index421).Int())
					goto C23
				}
			C23:
			}
			goto L23
		}
	L23:
		var _2425_v46 _dafny.Sequence
		_ = _2425_v46
		var _2426_v47 bool
		_ = _2426_v47
		var _2427_v48 bool
		_ = _2427_v48
		var _out50 _dafny.Sequence
		_ = _out50
		var _out51 bool
		_ = _out51
		var _out52 bool
		_ = _out52
		_out50, _out51, _out52 = (_this).M6(((_this).F6()) == ((_this).F6()), (p1).Cmp(p1) < 0, globalState)
		_2425_v46 = _out50
		_2426_v47 = _out51
		_2427_v48 = _out52
		(globalState).F1 = p1
		r0 = _2426_v47
		r0 = _2427_v48
		r1 = p2
		return r0, r1
	}
}
func (_this *C8) M3(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _2428_v0 _dafny.Int
		_ = _2428_v0
		_2428_v0 = _dafny.IntOfInt64(210)
		var _2429_v1 _dafny.Map
		_ = _2429_v1
		_2429_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2428_v0, (_this).F6())
		var _2430_v2 _dafny.Sequence
		_ = _2430_v2
		_2430_v2 = _dafny.SeqOf(Companion_Default___.Fm33(globalState), _2429_v1)
		var _2431_v3 _dafny.Map
		_ = _2431_v3
		_2431_v3 = (_2430_v2).Select((Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_2430_v2).Cardinality()))).Uint32()).(_dafny.Map)
		var _source33 _dafny.Map = _2431_v3
		_ = _source33
		var _2432___mcc_h0 _dafny.Map = _source33
		_ = _2432___mcc_h0
		var _2433_cf12 _dafny.Map = _2432___mcc_h0
		_ = _2433_cf12
		var _2434_v4 _dafny.Array
		_ = _2434_v4
		var _nwElement0_90 _dafny.Int = _2428_v0
		_ = _nwElement0_90
		var _nw420 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_90, nil, _dafny.IntOfInt64(9))
		_ = _nw420
		(_nw420).ArraySet1(_nwElement0_90, 0)
		(_nw420).ArraySet1(_2428_v0, 1)
		(_nw420).ArraySet1(_2428_v0, 2)
		(_nw420).ArraySet1(_2428_v0, 3)
		(_nw420).ArraySet1(_2428_v0, 4)
		(_nw420).ArraySet1(_2428_v0, 5)
		(_nw420).ArraySet1(_2428_v0, 6)
		(_nw420).ArraySet1(_2428_v0, 7)
		(_nw420).ArraySet1(_2428_v0, 8)
		_2434_v4 = _nw420
		var _2435_v5 _dafny.Array
		_ = _2435_v5
		_2435_v5 = _2434_v4
		var _2436_v6 D12
		_ = _2436_v6
		_2436_v6 = Companion_D12_.Create_DC23_(_2435_v5)
		var _pat_let_tv32 = _2435_v5
		_ = _pat_let_tv32
		var _pat_let_tv33 = _2434_v4
		_ = _pat_let_tv33
		var _source34 D12 = func(_pat_let42_0 D12) D12 {
			return func(_2439_dt__update__tmp_h1 D12) D12 {
				return func(_pat_let45_0 _dafny.Array) D12 {
					return func(_2440_dt__update_hcf29_h1 _dafny.Array) D12 {
						return Companion_D12_.Create_DC23_(_2440_dt__update_hcf29_h1)
					}(_pat_let45_0)
				}(_pat_let_tv33)
			}(_pat_let42_0)
		}((func() D12 {
			if (_this).F6() {
				return func(_pat_let43_0 D12) D12 {
					return func(_2437_dt__update__tmp_h0 D12) D12 {
						return func(_pat_let44_0 _dafny.Array) D12 {
							return func(_2438_dt__update_hcf29_h0 _dafny.Array) D12 {
								return Companion_D12_.Create_DC23_(_2438_dt__update_hcf29_h0)
							}(_pat_let44_0)
						}(_pat_let_tv32)
					}(_pat_let43_0)
				}(_2436_v6)
			}
			return _2436_v6
		})())
		_ = _source34
		if _source34.Is_DC22() {
			var _2441___mcc_h1 _dafny.Map = _source34.Get_().(D12_DC22).Cf24
			_ = _2441___mcc_h1
			var _2442___mcc_h2 _dafny.Sequence = _source34.Get_().(D12_DC22).Cf25
			_ = _2442___mcc_h2
			var _2443___mcc_h3 _dafny.CodePoint = _source34.Get_().(D12_DC22).Cf26
			_ = _2443___mcc_h3
			var _2444___mcc_h4 bool = _source34.Get_().(D12_DC22).Cf27
			_ = _2444___mcc_h4
			var _2445___mcc_h5 _dafny.Sequence = _source34.Get_().(D12_DC22).Cf28
			_ = _2445___mcc_h5
			var _2446_cf28 _dafny.Sequence = _2445___mcc_h5
			_ = _2446_cf28
			var _2447_cf27 bool = _2444___mcc_h4
			_ = _2447_cf27
			var _2448_cf26 _dafny.CodePoint = _2443___mcc_h3
			_ = _2448_cf26
			var _2449_cf25 _dafny.Sequence = _2442___mcc_h2
			_ = _2449_cf25
			var _2450_cf24 _dafny.Map = _2441___mcc_h1
			_ = _2450_cf24
			var _2451_v7 _dafny.Map
			_ = _2451_v7
			_2451_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_2428_v0), _2428_v0)
			var _2452_v8 _dafny.Sequence
			_ = _2452_v8
			_2452_v8 = _dafny.SeqOf((_2451_v7).Cardinality())
			var _2453_v9 *C3
			_ = _2453_v9
			var _nw421 *C3 = New_C3_()
			_ = _nw421
			_nw421.Ctor__(true, _2447_cf27)
			_2453_v9 = _nw421
			var _2454_v10 _dafny.Set
			_ = _2454_v10
			_2454_v10 = _dafny.SetOf(_2453_v9, _2453_v9, _2453_v9)
			var _2455_v11 _dafny.MultiSet
			_ = _2455_v11
			_2455_v11 = _dafny.MultiSetOf(_2452_v8, _2452_v8)
			var _2456_v12 _dafny.Map
			_ = _2456_v12
			_2456_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2455_v11, _2428_v0)
			var _rhs474 _dafny.Sequence = (func() _dafny.Sequence {
				if (_2454_v10).IsSubsetOf(_2454_v10) {
					return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2452_v8, _dafny.Companion_Sequence_.Update(_2452_v8, (Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_2452_v8).Cardinality()))).Uint32(), _2428_v0)), (Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_2456_v12).Contains(_2455_v11) {
							return (_2456_v12).Get(_2455_v11).(_dafny.Int)
						}
						return _2428_v0
					})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2452_v8, _dafny.Companion_Sequence_.Update(_2452_v8, (Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_2452_v8).Cardinality()))).Uint32(), _2428_v0))).Cardinality()))).Uint32(), _dafny.IntOfInt64(672))
				}
				return _2452_v8
			})()
			_ = _rhs474
			var _rhs475 bool = _2447_cf27
			_ = _rhs475
			_2452_v8 = _rhs474
			_2447_cf27 = _rhs475
			var _2457_v13 _dafny.Array
			_ = _2457_v13
			var _nw422 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
			_ = _nw422
			_2457_v13 = _nw422
			var _2458_v14 D18
			_ = _2458_v14
			_2458_v14 = Companion_D18_.Create_DC36_(_2457_v13)
			var _2459_v15 _dafny.Map
			_ = _2459_v15
			_2459_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2458_v14, (_2428_v0).Minus(_dafny.IntOfInt64(-10)))
			_2459_v15 = (_2459_v15).Update(_2458_v14, Companion_Default___.SafeDivisionInt(_2428_v0, _2428_v0))
			_2428_v0 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2434_v4, _2428_v0)).Cardinality()).Plus(_2428_v0)
			var _index422 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_2434_v4), 0))
			_ = _index422
			(_2434_v4).ArraySet1(_2428_v0, (_index422).Int())
			var _index423 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(267), _dafny.ArrayLen((_2434_v4), 0))
			_ = _index423
			(_2434_v4).ArraySet1(Companion_Default___.SafeDivisionInt(_2428_v0, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2446_cf28).Cardinality()), _dafny.IntOfUint32((_2452_v8).Cardinality()))), (_index423).Int())
		} else if _source34.Is_DC23() {
			var _2460___mcc_h6 _dafny.Array = _source34.Get_().(D12_DC23).Cf29
			_ = _2460___mcc_h6
			var _2461_cf29 _dafny.Array = _2460___mcc_h6
			_ = _2461_cf29
			var _2462_v16 _dafny.CodePoint
			_ = _2462_v16
			_2462_v16 = _dafny.CodePoint('e')
			_2462_v16 = (_this).F5()
			var _2463_v17 _dafny.Set
			_ = _2463_v17
			_2463_v17 = _dafny.SetOf((_this).F6())
			_2463_v17 = _2463_v17
			var _2464_v18 _dafny.Sequence
			_ = _2464_v18
			_2464_v18 = _dafny.SeqOf(_2428_v0, _2428_v0)
			var _2465_v19 _dafny.MultiSet
			_ = _2465_v19
			_2465_v19 = _dafny.MultiSetOf((_2464_v18).Select((Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_2464_v18).Cardinality()))).Uint32()).(_dafny.Int))
			_2465_v19 = Companion_Default___.Fm32(globalState)
			var _2466_v20 bool
			_ = _2466_v20
			_2466_v20 = true
			var _2467_v21 _dafny.Map
			_ = _2467_v21
			_2467_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2428_v0, _2462_v16)
			var _2468_v22 _dafny.Sequence
			_ = _2468_v22
			_2468_v22 = _dafny.UnicodeSeqOfUtf8Bytes("ryhpyo")
			var _2469_v23 _dafny.Map
			_ = _2469_v23
			_2469_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), (func() _dafny.CodePoint {
				if (_2467_v21).Contains(_dafny.IntOfUint32((_2468_v22).Cardinality())) {
					return (_2467_v21).Get(_dafny.IntOfUint32((_2468_v22).Cardinality())).(_dafny.CodePoint)
				}
				return (_this).F5()
			})())
			var _2470_v24 _dafny.Set
			_ = _2470_v24
			_2470_v24 = _dafny.SetOf(_2469_v23)
			var _2471_v25 _dafny.Map
			_ = _2471_v25
			_2471_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2469_v23, (_this).F6())
			var _2472_v27 _dafny.MultiSet
			_ = _2472_v27
			_2472_v27 = _dafny.MultiSetOf(!((_this).F6()), (_this).F6())
			var _rhs476 bool = (_this).F6()
			_ = _rhs476
			var _rhs477 _dafny.Set = (func() _dafny.Set {
				var _coll75 = _dafny.NewBuilder()
				_ = _coll75
				for _iter81 := _dafny.Iterate((_2471_v25).Keys().Elements()); ; {
					_compr_75, _ok81 := _iter81()
					if !_ok81 {
						break
					}
					var _2473_v26 _dafny.Map
					_2473_v26 = interface{}(_compr_75).(_dafny.Map)
					if (_2471_v25).Contains(_2473_v26) {
						_coll75.Add(_2473_v26)
					}
				}
				return _coll75.ToSet()
			}()).Union(_2470_v24)
			_ = _rhs477
			var _rhs478 _dafny.MultiSet = (_2472_v27).Difference(_dafny.MultiSetOf(!(Companion_Default___.Fm0(_2466_v20, globalState))))
			_ = _rhs478
			var _lhs317 *GlobalState = globalState
			_ = _lhs317
			_2466_v20 = _rhs476
			_2470_v24 = _rhs477
			_lhs317.F4 = _rhs478
		} else {
			var _2474___mcc_h7 _dafny.Map = _source34.Get_().(D12_DC21).Cf23
			_ = _2474___mcc_h7
			var _2475_cf23 _dafny.Map = _2474___mcc_h7
			_ = _2475_cf23
			(globalState).F1 = _2428_v0
			var _2476_v28 *C5
			_ = _2476_v28
			var _nw423 *C5 = New_C5_()
			_ = _nw423
			_nw423.Ctor__(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6()))).Cardinality(), (_this).F5(), (_this).F6())
			_2476_v28 = _nw423
			var _2477_v29 D6
			_ = _2477_v29
			_2477_v29 = Companion_D6_.Create_DC14_((_this).F6())
			var _2478_v30 _dafny.Array
			_ = _2478_v30
			var _len0_65 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_65
			var _nw424 _dafny.Array
			_ = _nw424
			if _len0_65.Cmp(_dafny.Zero) == 0 {
				_nw424 = _dafny.NewArray(_len0_65)
			} else {
				var _init65 func(_dafny.Int) _dafny.MultiSet = func(_2479_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((_this).F6())
				}
				_ = _init65
				var _element0_65 = _init65(_dafny.Zero)
				_ = _element0_65
				_nw424 = _dafny.NewArrayFromExample(_element0_65, nil, _len0_65)
				(_nw424).ArraySet1(_element0_65, 0)
				var _nativeLen0_65 = (_len0_65).Int()
				_ = _nativeLen0_65
				for _i0_65 := 1; _i0_65 < _nativeLen0_65; _i0_65++ {
					(_nw424).ArraySet1(_init65(_dafny.IntOf(_i0_65)), _i0_65)
				}
			}
			_2478_v30 = _nw424
			var _2480_v31 D18
			_ = _2480_v31
			_2480_v31 = Companion_D18_.Create_DC36_(_2478_v30)
			var _2481_v32 _dafny.MultiSet
			_ = _2481_v32
			_2481_v32 = _dafny.MultiSetOf(_2480_v31, _2480_v31, Companion_D18_.Create_DC36_(_2478_v30), _2480_v31, (func() D18 {
				if (_this).F6() {
					return _2480_v31
				}
				return _2480_v31
			})())
			var _rhs479 D6 = Companion_Default___.Fm64((_this).F5(), globalState)
			_ = _rhs479
			var _rhs480 _dafny.MultiSet = (_dafny.MultiSetOf(_2480_v31)).Union(_2481_v32)
			_ = _rhs480
			_2477_v29 = _rhs479
			_2481_v32 = _rhs480
			var _2482_v33 _dafny.Set
			_ = _2482_v33
			_2482_v33 = _dafny.SetOf(_2428_v0, _2476_v28.F15)
			var _2483_v34 _dafny.Map
			_ = _2483_v34
			_2483_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2428_v0, _2482_v33)
			var _2484_v35 _dafny.Sequence
			_ = _2484_v35
			_2484_v35 = _dafny.UnicodeSeqOfUtf8Bytes("wae")
			_2483_v34 = (_2483_v34).Update(_2428_v0, (_2482_v33).Union(_dafny.SetOf(_dafny.IntOfUint32((_2484_v35).Cardinality()), _2476_v28.F15)))
		}
		var _2485_v36 bool
		_ = _2485_v36
		_2485_v36 = true
		var _2486_v37 _dafny.Sequence
		_ = _2486_v37
		_2486_v37 = _dafny.SeqOf(_2428_v0)
		var _2487_v38 _dafny.MultiSet
		_ = _2487_v38
		_2487_v38 = _dafny.MultiSetOf(true)
		_2485_v36 = (_this).Fm6(_dafny.Companion_Sequence_.Concatenate(_2486_v37, _dafny.SeqOf((_2487_v38).Cardinality(), _2428_v0)), globalState)
		var _arr6 _dafny.Array = _this.F7
		_ = _arr6
		var _index424 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_this.F7), 0))
		_ = _index424
		(_arr6).ArraySet1(_2485_v36, (_index424).Int())
		var _arr7 _dafny.Array = _this.F7
		_ = _arr7
		var _index425 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_this.F7), 0))
		_ = _index425
		(_arr7).ArraySet1((_this).F6(), (_index425).Int())
		var _arr8 _dafny.Array = _this.F7
		_ = _arr8
		var _index426 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_this.F7), 0))
		_ = _index426
		(_arr8).ArraySet1((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(173), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), (_index426).Int())
		var _2488_v39 _dafny.MultiSet
		_ = _2488_v39
		_2488_v39 = _dafny.MultiSetOf((_this).F6())
		var _2489_v40 _dafny.Sequence
		_ = _2489_v40
		_2489_v40 = _dafny.SeqOf((_this).F6(), (_this).F6(), (_this).F6())
		var _2490_v41 *C6
		_ = _2490_v41
		var _nw425 *C6 = New_C6_()
		_ = _nw425
		_nw425.Ctor__()
		_2490_v41 = _nw425
		var _2491_v42 _dafny.Map
		_ = _2491_v42
		_2491_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2490_v41, (_this).F6())
		var _2492_v43 _dafny.Map
		_ = _2492_v43
		_2492_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2491_v42).Cardinality(), _2428_v0)
		var _2493_v44 _dafny.Set
		_ = _2493_v44
		_2493_v44 = _dafny.SetOf(_2488_v39, _dafny.MultiSetFromSeq(_2489_v40), _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F6()), (Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Uint32(), (_this).F6()), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_2492_v43).Contains(_dafny.IntOfInt64(-249)) {
				return (_2492_v43).Get(_dafny.IntOfInt64(-249)).(_dafny.Int)
			}
			return _2428_v0
		})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F6()), (Companion_Default___.SafeIndex(_2428_v0, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()))).Uint32(), (_this).F6())).Cardinality()))).Uint32(), (_this).F6())))
		var _2494_v45 *C7
		_ = _2494_v45
		var _nw426 *C7 = New_C7_()
		_ = _nw426
		_nw426.Ctor__((_2493_v44).Cardinality(), Companion_Default___.SafeModuloInt(_2428_v0, _2428_v0), (_this).F5(), (_this).F6())
		_2494_v45 = _nw426
		if (_dafny.IntOfInt64(946)).Cmp(_dafny.IntOfInt64(800)) > 0 {
			var _2495_v46 bool
			_ = _2495_v46
			_2495_v46 = true
			_2495_v46 = ((_this).F6()) && ((_this).F6())
			var _arr9 _dafny.Array = _this.F7
			_ = _arr9
			var _index427 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))
			_ = _index427
			(_arr9).ArraySet1((_this).F6(), (_index427).Int())
			var _2496_v47 _dafny.Sequence
			_ = _2496_v47
			_2496_v47 = _dafny.UnicodeSeqOfUtf8Bytes("mkmwkypd")
			var _2497_v48 _dafny.Sequence
			_ = _2497_v48
			_2497_v48 = _dafny.SeqOf((_2494_v45).F9())
			var _2498_v49 _dafny.Map
			_ = _2498_v49
			_2498_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2428_v0, _2497_v48)
			var _2499_v50 _dafny.MultiSet
			_ = _2499_v50
			_2499_v50 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_2428_v0, _2494_v45.F10)).Cardinality()), (_2494_v45).F9())
			var _arr10 _dafny.Array = _this.F7
			_ = _arr10
			var _index428 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))
			_ = _index428
			var _rhs481 bool = (func() bool {
				if true {
					return _2495_v46
				}
				return (_2490_v41).Fm6(_2497_v48, globalState)
			})()
			_ = _rhs481
			var _rhs482 _dafny.Int = _dafny.IntOfInt64(111)
			_ = _rhs482
			var _rhs483 bool = (func() bool {
				if (func() bool {
					if !((Companion_D19_.Create_DC40_((_this).F6(), (_2494_v45).F9(), _2498_v49)).Dtor_cf51()) {
						return (_this).F6()
					}
					return (_this).F6()
				})() {
					return (_this).F6()
				}
				return !((_this).F6())
			})()
			_ = _rhs483
			var _rhs484 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2496_v47, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(113))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg126 _dafny.Int) interface{} {
					return coer126(arg126)
				}
			}(func(_2500_i1 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			})), (Companion_Default___.SafeIndex((_2494_v45).Fm4(_2489_v40, _2428_v0, _2494_v45.F10, _2499_v50, globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(113))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg127 _dafny.Int) interface{} {
					return coer127(arg127)
				}
			}(func(_2501_i1 _dafny.Int) _dafny.CodePoint {
				return (_this).F5()
			}))).Cardinality()))).Uint32(), (_this).F5())), _2496_v47)
			_ = _rhs484
			var _lhs318 _dafny.Array = _this.F7
			_ = _lhs318
			var _lhs319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))
			_ = _lhs319
			_2495_v46 = _rhs481
			_2428_v0 = _rhs482
			(_lhs318).ArraySet1(_rhs483, (_lhs319).Int())
			_2496_v47 = _rhs484
			var _2502_v51 _dafny.Set
			_ = _2502_v51
			_2502_v51 = _dafny.SetOf(_2428_v0, (_2494_v45).F9())
			_2495_v46 = (_2502_v51).Equals(_2502_v51)
			var _2503_v52 _dafny.Map
			_ = _2503_v52
			_2503_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_2495_v46, globalState), false)
			var _2504_v53 D21
			_ = _2504_v53
			_2504_v53 = Companion_D21_.Create_DC45_(_2499_v50)
			var _2505_v54 _dafny.Map
			_ = _2505_v54
			_2505_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2488_v39).Update((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), Companion_Default___.Abs((_2503_v52).Cardinality())), ((_2504_v53).Dtor_cf60()).IsDisjointFrom(_2499_v50))
			_2505_v54 = (_2505_v54).Update(_2488_v39, !(!((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool))))
			var _2506_v55 _dafny.Sequence
			_ = _2506_v55
			var _2507_v56 bool
			_ = _2507_v56
			var _2508_v57 bool
			_ = _2508_v57
			var _out53 _dafny.Sequence
			_ = _out53
			var _out54 bool
			_ = _out54
			var _out55 bool
			_ = _out55
			_out53, _out54, _out55 = (_this).M6((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), (_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), globalState)
			_2506_v55 = _out53
			_2507_v56 = _out54
			_2508_v57 = _out55
		} else {
			if true {
				var _2509_v58 _dafny.Map
				_ = _2509_v58
				_2509_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_2494_v45.F10, !((_this).F6()), _2494_v45.F10, globalState), _2428_v0)
				var _2510_v59 _dafny.Sequence
				_ = _2510_v59
				_2510_v59 = _dafny.UnicodeSeqOfUtf8Bytes("deayq")
				var _2511_v60 _dafny.Map
				_ = _2511_v60
				_2511_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _2494_v45.F10)
				_2509_v58 = (_2509_v58).Update(_dafny.Companion_Sequence_.Concatenate(_2510_v59, _2510_v59), (func() _dafny.Int {
					if (_2511_v60).Contains((_this).F6()) {
						return (_2511_v60).Get((_this).F6()).(_dafny.Int)
					}
					return _2428_v0
				})())
				var _2512_v61 *C1
				_ = _2512_v61
				var _nw427 *C1 = New_C1_()
				_ = _nw427
				_nw427.Ctor__()
				_2512_v61 = _nw427
				var _2513_v62 bool
				_ = _2513_v62
				_2513_v62 = true
				_2513_v62 = !((_this).F6()) || (!((_this).F6()))
				var _2514_v63 _dafny.Sequence
				_ = _2514_v63
				_2514_v63 = _dafny.SeqOf(_dafny.IntOfInt64(638))
				var _2515_v64 _dafny.Sequence
				_ = _2515_v64
				_2515_v64 = _dafny.SeqOf(_2514_v63)
				var _rhs485 bool = !_dafny.Companion_Sequence_.Contains(_2515_v64, (_2515_v64).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2514_v63).Cardinality()), _dafny.IntOfUint32((_2515_v64).Cardinality()))).Uint32()).(_dafny.Sequence))
				_ = _rhs485
				var _rhs486 _dafny.Int = _2428_v0
				_ = _rhs486
				var _rhs487 _dafny.Int = (_2494_v45).F9()
				_ = _rhs487
				var _lhs320 *C7 = _2494_v45
				_ = _lhs320
				var _lhs321 *GlobalState = globalState
				_ = _lhs321
				_2513_v62 = _rhs485
				_lhs320.F10 = _rhs486
				_lhs321.F1 = _rhs487
				var _2516_v65 _dafny.Set
				_ = _2516_v65
				_2516_v65 = _dafny.SetOf((_this).F6())
				var _2517_v66 _dafny.Map
				_ = _2517_v66
				_2517_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm6(_dafny.Companion_Sequence_.Update(_2514_v63, (Companion_Default___.SafeIndex(_2494_v45.F10, _dafny.IntOfUint32((_2514_v63).Cardinality()))).Uint32(), (_2494_v45).F9()), globalState), _2516_v65)
				var _2518_v67 _dafny.Sequence
				_ = _2518_v67
				_2518_v67 = _dafny.SeqOf(_2516_v65, (func() _dafny.Set {
					if (_2517_v66).Contains(true) {
						return (_2517_v66).Get(true).(_dafny.Set)
					}
					return _2516_v65
				})())
				var _2519_v68 _dafny.Map
				_ = _2519_v68
				_2519_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _this.F7)
				var _2520_v69 _dafny.Map
				_ = _2520_v69
				_2520_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2518_v67, ((_2519_v68).Merge(_2519_v68)).Cardinality())
				_2520_v69 = (_2520_v69).Update(_2518_v67, _2428_v0)
			} else {
				var _arr11 _dafny.Array = _this.F7
				_ = _arr11
				var _index429 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_this.F7), 0))
				_ = _index429
				(_arr11).ArraySet1(!(((_2494_v45).F9()).Cmp((_2490_v41).Fm7((_this).F6(), globalState)) >= 0), (_index429).Int())
				var _arr12 _dafny.Array = _this.F7
				_ = _arr12
				var _index430 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_this.F7), 0))
				_ = _index430
				(_arr12).ArraySet1(Companion_Default___.Fm0((_this).F6(), globalState), (_index430).Int())
				var _2521_v70 _dafny.Sequence
				_ = _2521_v70
				_2521_v70 = _dafny.UnicodeSeqOfUtf8Bytes("baymc")
				var _2522_v71 _dafny.Array
				_ = _2522_v71
				var _nw428 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw428
				_2522_v71 = _nw428
				var _2523_v72 _dafny.Map
				_ = _2523_v72
				_2523_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2494_v45.F10, _2522_v71)
				var _2524_v73 _dafny.Sequence
				_ = _2524_v73
				_2524_v73 = _dafny.SeqOf(_2494_v45.F10, _2494_v45.F10, (_dafny.Zero).Minus(_dafny.IntOfUint32((_2521_v70).Cardinality())))
				var _2525_v74 _dafny.MultiSet
				_ = _2525_v74
				_2525_v74 = _dafny.MultiSetOf((_2494_v45).F9())
				var _2526_v75 _dafny.Map
				_ = _2526_v75
				_2526_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2494_v45.F10, _dafny.SeqOf((func() _dafny.Int {
					if (_2525_v74).Contains(_2494_v45.F10) {
						return (_2525_v74).Multiplicity(_2494_v45.F10)
					}
					return (_2494_v45).F9()
				})()))
				var _2527_v76 _dafny.MultiSet
				_ = _2527_v76
				_2527_v76 = _dafny.MultiSetOf((_2526_v75).Cardinality(), _dafny.IntOfInt64(-162))
				var _2528_v77 _dafny.Array
				_ = _2528_v77
				var _nwElement0_91 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F6() {
						return (_2494_v45).F9()
					}
					return _2494_v45.F10
				})())
				_ = _nwElement0_91
				var _nw429 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_91, nil, _dafny.IntOfInt64(24))
				_ = _nw429
				(_nw429).ArraySet1(_nwElement0_91, 0)
				(_nw429).ArraySet1((_2494_v45).F9(), 1)
				(_nw429).ArraySet1((_2494_v45).F9(), 2)
				(_nw429).ArraySet1((_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2521_v70, (Companion_Default___.SafeIndex((_2494_v45).F9(), _dafny.IntOfUint32((_2521_v70).Cardinality()))).Uint32(), (_this).F5())).Cardinality())).Times((_2494_v45).F9()), 3)
				(_nw429).ArraySet1((((_2523_v72).Update(_2494_v45.F10, _2522_v71)).Merge(_2523_v72)).Cardinality(), 4)
				(_nw429).ArraySet1(((_2494_v45).F9()).Minus(_dafny.IntOfInt64(-572)), 5)
				(_nw429).ArraySet1(_dafny.IntOfUint32((_2524_v73).Cardinality()), 6)
				(_nw429).ArraySet1((_2494_v45).F9(), 7)
				(_nw429).ArraySet1((_2494_v45).F9(), 8)
				(_nw429).ArraySet1((_2494_v45).F9(), 9)
				(_nw429).ArraySet1(_2494_v45.F10, 10)
				(_nw429).ArraySet1(_2494_v45.F10, 11)
				(_nw429).ArraySet1((_2494_v45).Fm4(_2489_v40, (func() _dafny.Int {
					if (_2488_v39).Contains((_this).F6()) {
						return (_2488_v39).Multiplicity((_this).F6())
					}
					return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(162))).Uint32(), func(coer128 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg128 _dafny.Int) interface{} {
							return coer128(arg128)
						}
					}((func(_2529_v40 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_2530_i2 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer129 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
								return func(arg129 _dafny.Int) interface{} {
									return coer129(arg129)
								}
							}((func(_2531_v40 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_2532_i3 _dafny.Int) _dafny.Sequence {
									return _2531_v40
								}
							})(_2529_v40)))).Cardinality())
						}
					})(_2489_v40)))).Cardinality())
				})(), _2494_v45.F10, _2527_v76, globalState), 12)
				(_nw429).ArraySet1((_2494_v45).Fm7((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), globalState), 13)
				(_nw429).ArraySet1((_2525_v74).Cardinality(), 14)
				(_nw429).ArraySet1(_dafny.IntOfInt64(-739), 15)
				(_nw429).ArraySet1(Companion_Default___.Fm3((_this).F6(), (_this).F6(), globalState), 16)
				(_nw429).ArraySet1(_2494_v45.F10, 17)
				(_nw429).ArraySet1((_2494_v45).F9(), 18)
				(_nw429).ArraySet1(_dafny.IntOfInt64(221), 19)
				(_nw429).ArraySet1(((_2494_v45).F9()).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(585), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), _2494_v45.F10)).Cardinality()), 20)
				(_nw429).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(738))).Uint32(), func(coer130 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}(func(_2533_i4 _dafny.Int) _dafny.CodePoint {
					return (_this).F5()
				}))).Cardinality()), 21)
				(_nw429).ArraySet1((_2494_v45).F9(), 22)
				(_nw429).ArraySet1(_dafny.IntOfUint32((_2521_v70).Cardinality()), 23)
				_2528_v77 = _nw429
				var _2534_v78 _dafny.Set
				_ = _2534_v78
				_2534_v78 = _dafny.SetOf(_2494_v45.F10, _2494_v45.F10)
				var _2535_v79 _dafny.Sequence
				_ = _2535_v79
				_2535_v79 = _dafny.SeqOf(_2534_v78, _2534_v78)
				var _index431 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2528_v77), 0))
				_ = _index431
				(_2528_v77).ArraySet1(_dafny.IntOfUint32((_2535_v79).Cardinality()), (_index431).Int())
				var _index432 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2528_v77), 0))
				_ = _index432
				(_2528_v77).ArraySet1(Companion_Default___.SafeModuloInt((_2494_v45).F9(), _2494_v45.F10), (_index432).Int())
				_2428_v0 = _dafny.IntOfInt64(135)
				var _2536_v80 _dafny.Array
				_ = _2536_v80
				_2536_v80 = _2528_v77
				var _2537_v81 _dafny.Map
				_ = _2537_v81
				_2537_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2489_v40).Cardinality()), _2536_v80)
				(_2494_v45).F10 = Companion_Default___.SafeModuloInt(((_2528_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2528_v77), 0))).Int()).(_dafny.Int)).Minus((_2528_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2528_v77), 0))).Int()).(_dafny.Int)), (_2537_v81).Cardinality())
				var _index433 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_2528_v77), 0))
				_ = _index433
				(_2528_v77).ArraySet1(_2494_v45.F10, (_index433).Int())
			}
			var _2538_v82 _dafny.MultiSet
			_ = _2538_v82
			_2538_v82 = _dafny.MultiSetOf((_2494_v45).F9())
			var _2539_v83 _dafny.Map
			_ = _2539_v83
			_2539_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2538_v82)
			(globalState).F1 = ((_2539_v83).Merge((_2539_v83).Merge(_2539_v83))).Cardinality()
			var _arr13 _dafny.Array = _this.F7
			_ = _arr13
			var _index434 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_this.F7), 0))
			_ = _index434
			(_arr13).ArraySet1(((_this).F6()) && ((_2489_v40).Select((Companion_Default___.SafeIndex(_2494_v45.F10, _dafny.IntOfUint32((_2489_v40).Cardinality()))).Uint32()).(bool)), (_index434).Int())
			var _arr14 _dafny.Array = _this.F7
			_ = _arr14
			var _index435 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_this.F7), 0))
			_ = _index435
			(_arr14).ArraySet1(false, (_index435).Int())
			var _2540_v84 _dafny.Map
			_ = _2540_v84
			_2540_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_2494_v45).F9())
			var _2541_v85 D6
			_ = _2541_v85
			_2541_v85 = Companion_D6_.Create_DC13_(_dafny.SeqOf((_2494_v45).F9()))
			var _2542_v86 _dafny.Sequence
			_ = _2542_v86
			_2542_v86 = _dafny.UnicodeSeqOfUtf8Bytes("uuafj")
			var _2543_v87 _dafny.Map
			_ = _2543_v87
			_2543_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2541_v85).Dtor_cf14(), _dafny.IntOfUint32((_2542_v86).Cardinality()))
			var _rhs488 _dafny.Map = Companion_Default___.Fm28((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), (_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), _2543_v87, globalState)
			_ = _rhs488
			var _rhs489 _dafny.Int = (_2494_v45).F9()
			_ = _rhs489
			var _lhs322 *GlobalState = globalState
			_ = _lhs322
			_2540_v84 = _rhs488
			_lhs322.F1 = _rhs489
			var _2544_v88 _dafny.Array
			_ = _2544_v88
			var _len0_66 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_66
			var _nw430 _dafny.Array
			_ = _nw430
			if _len0_66.Cmp(_dafny.Zero) == 0 {
				_nw430 = _dafny.NewArray(_len0_66)
			} else {
				var _init66 func(_dafny.Int) _dafny.Int = (func(_2545_v45 *C7) func(_dafny.Int) _dafny.Int {
					return func(_2546_i5 _dafny.Int) _dafny.Int {
						return (_2546_i5).Times(_2545_v45.F10)
					}
				})(_2494_v45)
				_ = _init66
				var _element0_66 = _init66(_dafny.Zero)
				_ = _element0_66
				_nw430 = _dafny.NewArrayFromExample(_element0_66, nil, _len0_66)
				(_nw430).ArraySet1(_element0_66, 0)
				var _nativeLen0_66 = (_len0_66).Int()
				_ = _nativeLen0_66
				for _i0_66 := 1; _i0_66 < _nativeLen0_66; _i0_66++ {
					(_nw430).ArraySet1(_init66(_dafny.IntOf(_i0_66)), _i0_66)
				}
			}
			_2544_v88 = _nw430
			var _len0_67 _dafny.Int = _dafny.One
			_ = _len0_67
			var _nw431 _dafny.Array
			_ = _nw431
			if _len0_67.Cmp(_dafny.Zero) == 0 {
				_nw431 = _dafny.NewArray(_len0_67)
			} else {
				var _init67 func(_dafny.Int) _dafny.Int = (func(_2547_v45 *C7) func(_dafny.Int) _dafny.Int {
					return func(_2548_i6 _dafny.Int) _dafny.Int {
						return (_2548_i6).Times(_2547_v45.F10)
					}
				})(_2494_v45)
				_ = _init67
				var _element0_67 = _init67(_dafny.Zero)
				_ = _element0_67
				_nw431 = _dafny.NewArrayFromExample(_element0_67, nil, _len0_67)
				(_nw431).ArraySet1(_element0_67, 0)
				var _nativeLen0_67 = (_len0_67).Int()
				_ = _nativeLen0_67
				for _i0_67 := 1; _i0_67 < _nativeLen0_67; _i0_67++ {
					(_nw431).ArraySet1(_init67(_dafny.IntOf(_i0_67)), _i0_67)
				}
			}
			_2544_v88 = _nw431
		}
		var _2549_v89 bool
		_ = _2549_v89
		_2549_v89 = true
		_2549_v89 = _2549_v89
		_2549_v89 = _2549_v89
		var _2550_v90 _dafny.Map
		_ = _2550_v90
		_2550_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2549_v89, !(true))
		var _2551_v91 *C5
		_ = _2551_v91
		var _nw432 *C5 = New_C5_()
		_ = _nw432
		_nw432.Ctor__((_2488_v39).Cardinality(), (_this).F5(), !(!((_2492_v43).Update((_2494_v45).F9(), _dafny.IntOfUint32((Companion_Default___.Fm1((_2550_v90).Cardinality(), (_this).F6(), _2428_v0, globalState)).Cardinality()))).Contains(_2494_v45.F10)))
		_2551_v91 = _nw432
		r0 = (_2428_v0).Times(_2551_v91.F15)
		return r0
	}
}
func (_this *C8) M1(p0 bool, p1 _dafny.MultiSet, p2 bool, globalState *GlobalState) {
	{
		var _2552_v0 _dafny.CodePoint
		_ = _2552_v0
		_2552_v0 = _dafny.CodePoint('l')
		_2552_v0 = (_this).F5()
		var _2553_v1 _dafny.Sequence
		_ = _2553_v1
		_2553_v1 = _dafny.SeqOf(p0)
		var _2554_v2 D15
		_ = _2554_v2
		_2554_v2 = Companion_D15_.Create_DC28_((_this).F6(), (_2553_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_2553_v1).Cardinality()))), _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32()).(bool))
		var _source35 D15 = _2554_v2
		_ = _source35
		if _source35.Is_DC28() {
			var _2555___mcc_h0 bool = _source35.Get_().(D15_DC28).Cf36
			_ = _2555___mcc_h0
			var _2556___mcc_h1 bool = _source35.Get_().(D15_DC28).Cf37
			_ = _2556___mcc_h1
			var _2557_cf37 bool = _2556___mcc_h1
			_ = _2557_cf37
			var _2558_cf36 bool = _2555___mcc_h0
			_ = _2558_cf36
			var _2559_v3 _dafny.Int
			_ = _2559_v3
			_2559_v3 = _dafny.IntOfInt64(-712)
			var _2560_v4 _dafny.Sequence
			_ = _2560_v4
			_2560_v4 = _dafny.SeqOf(_2553_v1)
			var _rhs490 bool = (_2559_v3).Cmp((_2559_v3).Plus(_2559_v3)) <= 0
			_ = _rhs490
			var _rhs491 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2553_v1, _2553_v1), (_2560_v4).Select((Companion_Default___.SafeIndex(_2559_v3, _dafny.IntOfUint32((_2560_v4).Cardinality()))).Uint32()).(_dafny.Sequence))
			_ = _rhs491
			var _rhs492 bool = (_2558_cf36) && (_2557_cf37)
			_ = _rhs492
			_2558_cf36 = _rhs490
			_2553_v1 = _rhs491
			_2557_cf37 = _rhs492
			_2558_cf36 = _2557_cf37
			var _arr15 _dafny.Array = _this.F7
			_ = _arr15
			var _index436 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_this.F7), 0))
			_ = _index436
			(_arr15).ArraySet1(p0, (_index436).Int())
			var _2561_v5 D0
			_ = _2561_v5
			_2561_v5 = Companion_D0_.Create_DC0_(p2)
			var _2562_v6 _dafny.Map
			_ = _2562_v6
			_2562_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2559_v3, _dafny.MultiSetOf((_dafny.SetOf((_this).Fm9(_2559_v3, _2561_v5, _2558_cf36, _2559_v3, globalState))).Cardinality()))
			var _2563_v7 _dafny.MultiSet
			_ = _2563_v7
			_2563_v7 = _dafny.MultiSetOf(_2559_v3)
			var _2564_v8 _dafny.Sequence
			_ = _2564_v8
			_2564_v8 = _dafny.UnicodeSeqOfUtf8Bytes("rgibkeo")
			var _2565_v9 D6
			_ = _2565_v9
			_2565_v9 = Companion_D6_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer131 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg131 _dafny.Int) interface{} {
					return coer131(arg131)
				}
			}((func(_2566_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_2567_i0 _dafny.Int) _dafny.Int {
					return _2566_v3
				}
			})(_2559_v3))))
			var _2568_v10 _dafny.Map
			_ = _2568_v10
			_2568_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)
			var _2569_v11 D20
			_ = _2569_v11
			_2569_v11 = Companion_D20_.Create_DC43_(_dafny.IntOfUint32(((_2565_v9).Dtor_cf14()).Cardinality()), (func() bool {
				if (_2568_v10).Contains(p0) {
					return (_2568_v10).Get(p0).(bool)
				}
				return (_this).F6()
			})(), _dafny.IntOfInt64(88))
			var _2570_v12 D20
			_ = _2570_v12
			_2570_v12 = Companion_D20_.Create_DC44_(_2569_v11)
			var _2571_v13 _dafny.Set
			_ = _2571_v13
			_2571_v13 = _dafny.SetOf(_2570_v12)
			var _arr16 _dafny.Array = _this.F7
			_ = _arr16
			var _index437 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen((_this.F7), 0))
			_ = _index437
			(_arr16).ArraySet1(((func() _dafny.MultiSet {
				if (_2562_v6).Contains((_dafny.Zero).Minus(_2559_v3)) {
					return (_2562_v6).Get((_dafny.Zero).Minus(_2559_v3)).(_dafny.MultiSet)
				}
				return (_dafny.MultiSetOf(_dafny.IntOfUint32((_2564_v8).Cardinality()))).Update(_dafny.IntOfInt64(689), Companion_Default___.Abs((_2571_v13).Cardinality()))
			})()).IsSubsetOf((func() _dafny.MultiSet {
				if (_2562_v6).Contains(_2559_v3) {
					return (_2562_v6).Get(_2559_v3).(_dafny.MultiSet)
				}
				return _2563_v7
			})()), (_index437).Int())
			(globalState).F1 = _2559_v3
		} else if _source35.Is_DC27() {
			var _2572___mcc_h2 _dafny.MultiSet = _source35.Get_().(D15_DC27).Cf35
			_ = _2572___mcc_h2
			var _2573_cf35 _dafny.MultiSet = _2572___mcc_h2
			_ = _2573_cf35
			if p0 {
				var _2574_v14 _dafny.Sequence
				_ = _2574_v14
				_2574_v14 = _dafny.UnicodeSeqOfUtf8Bytes("giybh")
				var _2575_v15 _dafny.Int
				_ = _2575_v15
				_2575_v15 = _dafny.IntOfInt64(752)
				_2552_v0 = (func(_pat_let46_0 D17) D17 {
					return func(_2576_dt__update__tmp_h0 D17) D17 {
						return func(_pat_let47_0 _dafny.CodePoint) D17 {
							return func(_2577_dt__update_hcf43_h0 _dafny.CodePoint) D17 {
								return Companion_D17_.Create_DC34_(_2577_dt__update_hcf43_h0)
							}(_pat_let47_0)
						}((_this).F5())
					}(_pat_let46_0)
				}(Companion_Default___.Fm65((_this).F6(), _dafny.IntOfUint32((_2574_v14).Cardinality()), _2575_v15, globalState))).Dtor_cf43()
				var _2578_v16 *C5
				_ = _2578_v16
				var _nw433 *C5 = New_C5_()
				_ = _nw433
				_nw433.Ctor__(_2575_v15, (_this).F5(), (_dafny.IntOfInt64(-195)).Cmp((_dafny.Zero).Minus(_2575_v15)) < 0)
				_2578_v16 = _nw433
				var _arr17 _dafny.Array = _this.F7
				_ = _arr17
				var _index438 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_this.F7), 0))
				_ = _index438
				(_arr17).ArraySet1(!(p0) || (p0), (_index438).Int())
				var _arr18 _dafny.Array = _this.F7
				_ = _arr18
				var _index439 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_this.F7), 0))
				_ = _index439
				(_arr18).ArraySet1(!(false) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_2553_v1, _2553_v1)), (_index439).Int())
				var _2579_v17 _dafny.Sequence
				_ = _2579_v17
				_2579_v17 = _dafny.SeqOf(_2578_v16.F15, _dafny.IntOfUint32((_2574_v14).Cardinality()))
				var _2580_v18 _dafny.Sequence
				_ = _2580_v18
				_2580_v18 = _dafny.SeqOf(_2579_v17)
				var _2581_v19 *C0
				_ = _2581_v19
				var _nw434 *C0 = New_C0_()
				_ = _nw434
				_nw434.Ctor__(p0, _2580_v18)
				_2581_v19 = _nw434
				var _2582_v20 *C1
				_ = _2582_v20
				var _nw435 *C1 = New_C1_()
				_ = _nw435
				_nw435.Ctor__()
				_2582_v20 = _nw435
				var _2583_v21 _dafny.Map
				_ = _2583_v21
				_2583_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2574_v14).Cardinality()), (_2581_v19).F11())
				var _2584_v22 _dafny.Map
				_ = _2584_v22
				_2584_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2553_v1, _2583_v21)
				var _2585_v23 _dafny.Map
				_ = _2585_v23
				_2585_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2584_v22).Update(_2553_v1, _2583_v21), _2581_v19)
				var _arr19 _dafny.Array = _this.F7
				_ = _arr19
				var _index440 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_this.F7), 0))
				_ = _index440
				var _rhs493 *C0 = (func() *C0 {
					if (_2585_v23).Contains((_2584_v22).Merge(_2584_v22)) {
						return (_2585_v23).Get((_2584_v22).Merge(_2584_v22)).(*C0)
					}
					return _2581_v19
				})()
				_ = _rhs493
				var _rhs494 bool = (_2553_v1).Select((Companion_Default___.SafeIndex(_2578_v16.F15, _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32()).(bool)
				_ = _rhs494
				var _rhs495 *C1 = _2582_v20
				_ = _rhs495
				var _lhs323 _dafny.Array = _this.F7
				_ = _lhs323
				var _lhs324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_this.F7), 0))
				_ = _lhs324
				_2581_v19 = _rhs493
				(_lhs323).ArraySet1(_rhs494, (_lhs324).Int())
				_2582_v20 = _rhs495
				var _2586_v24 _dafny.Array
				_ = _2586_v24
				var _nw436 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw436
				_2586_v24 = _nw436
				var _2587_v25 _dafny.Array
				_ = _2587_v25
				var _nw437 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw437
				_2587_v25 = _nw437
				var _index441 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_2586_v24), 0))
				_ = _index441
				(_2586_v24).ArraySet1(_2587_v25, (_index441).Int())
				var _2588_v26 _dafny.Map
				_ = _2588_v26
				_2588_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(397), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), _dafny.IntOfInt64(-212))
				var _2589_v27 _dafny.Map
				_ = _2589_v27
				_2589_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), _2588_v26)
				var _index442 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_2586_v24), 0))
				_ = _index442
				var _rhs496 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2578_v16.F15), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2575_v15), _2579_v17)), (Companion_Default___.SafeIndex((_2589_v27).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2578_v16.F15), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2575_v15), _2579_v17))).Cardinality()))).Uint32(), (_2578_v16.F15).Minus(_dafny.IntOfUint32((_2579_v17).Cardinality())))
				_ = _rhs496
				var _rhs497 _dafny.Array = _2587_v25
				_ = _rhs497
				var _lhs325 _dafny.Array = _2586_v24
				_ = _lhs325
				var _lhs326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_2586_v24), 0))
				_ = _lhs326
				_2579_v17 = _rhs496
				(_lhs325).ArraySet1(_rhs497, (_lhs326).Int())
			} else {
				var _2590_v28 D21
				_ = _2590_v28
				_2590_v28 = Companion_D21_.Create_DC47_((_this).F6())
				var _2591_v29 bool
				_ = _2591_v29
				_2591_v29 = false
				var _2592_v30 _dafny.Array
				_ = _2592_v30
				var _nw438 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw438
				_2592_v30 = _nw438
				var _2593_v31 _dafny.Int
				_ = _2593_v31
				_2593_v31 = _dafny.IntOfInt64(893)
				var _index443 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))
				_ = _index443
				(_2592_v30).ArraySet1(Companion_Default___.SafeModuloInt(_2593_v31, _2593_v31), (_index443).Int())
				var _2594_v32 _dafny.Map
				_ = _2594_v32
				_2594_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).Fm7(p0, globalState)), (_this).F6())
				var _2595_v33 _dafny.Sequence
				_ = _2595_v33
				_2595_v33 = _dafny.SeqOf(_this.F7)
				var _2596_v34 _dafny.Sequence
				_ = _2596_v34
				_2596_v34 = _dafny.UnicodeSeqOfUtf8Bytes("eu")
				var _pat_let_tv34 = _2593_v31
				_ = _pat_let_tv34
				var _pat_let_tv35 = _2593_v31
				_ = _pat_let_tv35
				var _index444 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))
				_ = _index444
				var _rhs498 _dafny.Int = (func() _dafny.Int {
					if !((func() bool {
						if (_2594_v32).Contains(_2593_v31) {
							return (_2594_v32).Get(_2593_v31).(bool)
						}
						return _2591_v29
					})()) {
						return _2593_v31
					}
					return (p1).Cardinality()
				})()
				_ = _rhs498
				var _rhs499 D21 = func(_pat_let48_0 D21) D21 {
					return func(_2597_dt__update__tmp_h1 D21) D21 {
						return func(_pat_let49_0 bool) D21 {
							return func(_2598_dt__update_hcf64_h0 bool) D21 {
								return Companion_D21_.Create_DC47_(_2598_dt__update_hcf64_h0)
							}(_pat_let49_0)
						}((_pat_let_tv34).Cmp((_dafny.Zero).Minus(_pat_let_tv35)) <= 0)
					}(_pat_let48_0)
				}(_2590_v28)
				_ = _rhs499
				var _rhs500 bool = _2591_v29
				_ = _rhs500
				var _rhs501 _dafny.Array = (_2595_v33).Select((Companion_Default___.SafeIndex(_2593_v31, _dafny.IntOfUint32((_2595_v33).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs501
				var _rhs502 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(_2596_v34, _2596_v34, _2596_v34, _2596_v34, _2596_v34)).Cardinality()), (_this).Fm8(p2, _2591_v29, !(true), globalState))
				_ = _rhs502
				var _lhs327 *GlobalState = globalState
				_ = _lhs327
				var _lhs328 *C8 = _this
				_ = _lhs328
				var _lhs329 _dafny.Array = _2592_v30
				_ = _lhs329
				var _lhs330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))
				_ = _lhs330
				_lhs327.F1 = _rhs498
				_2590_v28 = _rhs499
				_2591_v29 = _rhs500
				_lhs328.F7 = _rhs501
				(_lhs329).ArraySet1(_rhs502, (_lhs330).Int())
				var _arr20 _dafny.Array = _this.F7
				_ = _arr20
				var _index445 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_this.F7), 0))
				_ = _index445
				(_arr20).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(419))).Uint32(), func(coer132 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg132 _dafny.Int) interface{} {
						return coer132(arg132)
					}
				}(func(_2599_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				})), _2596_v34), (_index445).Int())
				var _arr21 _dafny.Array = _this.F7
				_ = _arr21
				var _index446 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_this.F7), 0))
				_ = _index446
				(_arr21).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_2553_v1, _2553_v1), _2553_v1), (_index446).Int())
				_2593_v31 = _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ff"))).Cardinality())
				var _2600_v35 _dafny.Map
				_ = _2600_v35
				_2600_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2593_v31, _2596_v34)
				var _2601_v36 _dafny.Set
				_ = _2601_v36
				_2601_v36 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_2596_v34, (Companion_Default___.SafeIndex(_2593_v31, _dafny.IntOfUint32((_2596_v34).Cardinality()))).Uint32(), _dafny.CodePoint('f')))
				var _2602_v37 _dafny.Map
				_ = _2602_v37
				_2602_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2592_v30, (_2553_v1).Select((Companion_Default___.SafeIndex(_2593_v31, _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32()).(bool))
				_2596_v34 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_2600_v35).Contains((_2592_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))).Int()).(_dafny.Int)) {
						return (_2600_v35).Get((_2592_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
					}
					return _2596_v34
				})(), (Companion_Default___.SafeIndex((_2601_v36).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_2600_v35).Contains((_2592_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))).Int()).(_dafny.Int)) {
						return (_2600_v35).Get((_2592_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
					}
					return _2596_v34
				})()).Cardinality()))).Uint32(), _2552_v0), _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer133 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg133 _dafny.Int) interface{} {
						return coer133(arg133)
					}
				}((func(_2603_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2604_i2 _dafny.Int) _dafny.CodePoint {
						return _2603_v0
					}
				})(_2552_v0))), (Companion_Default___.SafeIndex((_2602_v37).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(628))).Uint32(), func(coer134 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg134 _dafny.Int) interface{} {
						return coer134(arg134)
					}
				}((func(_2605_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2606_i2 _dafny.Int) _dafny.CodePoint {
						return _2605_v0
					}
				})(_2552_v0)))).Cardinality()))).Uint32(), (Companion_Default___.Fm65((_this).F6(), _dafny.IntOfInt64(266), (_2592_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_2592_v30), 0))).Int()).(_dafny.Int), globalState)).Dtor_cf43()))
				var _arr22 _dafny.Array = _this.F7
				_ = _arr22
				var _index447 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_this.F7), 0))
				_ = _index447
				(_arr22).ArraySet1(_2591_v29, (_index447).Int())
			}
			if p2 {
				var _2607_v38 _dafny.Int
				_ = _2607_v38
				_2607_v38 = _dafny.IntOfInt64(788)
				(globalState).F1 = (_2607_v38).Minus((_this).Fm8(p2, p2, p2, globalState))
				var _2608_v39 _dafny.Map
				_ = _2608_v39
				_2608_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2607_v38)
				_2608_v39 = (_2608_v39).Update(false, _2607_v38)
				var _2609_v40 _dafny.Sequence
				_ = _2609_v40
				_2609_v40 = _dafny.SeqOf(_2607_v38)
				(globalState).F1 = (_this).Fm4(_2553_v1, _dafny.IntOfUint32((_2609_v40).Cardinality()), (_dafny.Zero).Minus(_2607_v38), _dafny.MultiSetOf(_2607_v38), globalState)
				var _2610_v41 *C3
				_ = _2610_v41
				var _nw439 *C3 = New_C3_()
				_ = _nw439
				_nw439.Ctor__(p2, p0)
				_2610_v41 = _nw439
				var _2611_v42 _dafny.MultiSet
				_ = _2611_v42
				_2611_v42 = _dafny.MultiSetOf((_dafny.Zero).Minus(_2607_v38))
				var _2612_v43 *C7
				_ = _2612_v43
				var _nw440 *C7 = New_C7_()
				_ = _nw440
				_nw440.Ctor__(_2607_v38, (_2611_v42).Cardinality(), (_this).F5(), p2)
				_2612_v43 = _nw440
			} else {
				_2552_v0 = (_this).F5()
				var _arr23 _dafny.Array = _this.F7
				_ = _arr23
				var _index448 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_this.F7), 0))
				_ = _index448
				(_arr23).ArraySet1((p0) == (false), (_index448).Int())
				var _2613_v44 _dafny.Map
				_ = _2613_v44
				_2613_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm49(globalState), p2)
				var _2614_v45 _dafny.Int
				_ = _2614_v45
				_2614_v45 = _dafny.IntOfInt64(120)
				var _2615_v46 _dafny.MultiSet
				_ = _2615_v46
				_2615_v46 = _dafny.MultiSetOf(_2614_v45)
				var _2616_v47 T3
				_ = _2616_v47
				var _nw441 *C5 = New_C5_()
				_ = _nw441
				_nw441.Ctor__((func() _dafny.Int {
					if (_this).Fm10(_2573_cf35, _2613_v44, _dafny.IntOfInt64(697), globalState) {
						return _dafny.IntOfInt64(345)
					}
					return (_2615_v46).Cardinality()
				})(), (_this).F5(), p2)
				_2616_v47 = _nw441
				var _2617_v48 _dafny.Sequence
				_ = _2617_v48
				_2617_v48 = _dafny.SeqOf(_dafny.IntOfInt64(-920), _2614_v45, _dafny.IntOfInt64(128), _2614_v45)
				var _arr24 _dafny.Array = _this.F7
				_ = _arr24
				var _index449 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_this.F7), 0))
				_ = _index449
				var _rhs503 bool = p0
				_ = _rhs503
				var _rhs504 _dafny.Int = _dafny.IntOfUint32((_2553_v1).Cardinality())
				_ = _rhs504
				var _rhs505 _dafny.Int = Companion_Default___.SafeModuloInt((_2614_v45).Minus((_2617_v48).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus(_2614_v45)), _dafny.IntOfUint32((_2617_v48).Cardinality()))).Uint32()).(_dafny.Int)), (func() _dafny.Int {
					if (_2573_cf35).Contains(p2) {
						return (_2573_cf35).Multiplicity(p2)
					}
					return _2614_v45
				})())
				_ = _rhs505
				var _rhs506 T3 = _2616_v47
				_ = _rhs506
				var _lhs331 _dafny.Array = _this.F7
				_ = _lhs331
				var _lhs332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(36), _dafny.ArrayLen((_this.F7), 0))
				_ = _lhs332
				var _lhs333 *GlobalState = globalState
				_ = _lhs333
				var _lhs334 *GlobalState = globalState
				_ = _lhs334
				(_lhs331).ArraySet1(_rhs503, (_lhs332).Int())
				_lhs333.F1 = _rhs504
				_lhs334.F1 = _rhs505
				_2616_v47 = _rhs506
				var _2618_v49 _dafny.Sequence
				_ = _2618_v49
				_2618_v49 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2617_v48, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-385))).Uint32(), func(coer135 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg135 _dafny.Int) interface{} {
						return coer135(arg135)
					}
				}(func(_2619_i3 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(205)
				}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm3((_2616_v47).F6(), p2, globalState), _dafny.IntOfInt64(-206)), _2617_v48))
				var _2620_v50 _dafny.Sequence
				_ = _2620_v50
				_2620_v50 = _dafny.UnicodeSeqOfUtf8Bytes("fdtgl")
				_2618_v49 = Companion_Default___.Fm38((_dafny.IntOfUint32((_2620_v50).Cardinality())).Cmp(_2614_v45) != 0, p0, _dafny.IntOfUint32((_2620_v50).Cardinality()), _2552_v0, globalState)
				_2614_v45 = _dafny.IntOfUint32((_2617_v48).Cardinality())
				(globalState).F1 = ((_dafny.Zero).Minus(_2614_v45)).Times(_2614_v45)
			}
			var _2621_v51 _dafny.Int
			_ = _2621_v51
			_2621_v51 = _dafny.IntOfInt64(710)
			var _2622_v52 _dafny.MultiSet
			_ = _2622_v52
			_2622_v52 = _dafny.MultiSetOf(_dafny.IntOfInt64(422))
			var _2623_v53 _dafny.Sequence
			_ = _2623_v53
			_2623_v53 = _dafny.UnicodeSeqOfUtf8Bytes("i")
			var _2624_v54 _dafny.Map
			_ = _2624_v54
			_2624_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4(_dafny.SeqOf(false), _2621_v51, _2621_v51, _2622_v52, globalState), _2623_v53)
			(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(644), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_2624_v54).Cardinality(), _2621_v51))))
			var _2625_v55 _dafny.Array
			_ = _2625_v55
			var _nw442 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw442
			_2625_v55 = _nw442
			var _2626_v56 _dafny.Map
			_ = _2626_v56
			_2626_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2621_v51, (_this).F6())
			var _2627_v58 _dafny.Set
			_ = _2627_v58
			_2627_v58 = _dafny.SetOf(_2621_v51, _2621_v51)
			var _2628_v59 _dafny.Array
			_ = _2628_v59
			var _nwElement0_92 _dafny.Map = _2626_v56
			_ = _nwElement0_92
			var _nw443 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_92, nil, _dafny.IntOfInt64(25))
			_ = _nw443
			(_nw443).ArraySet1(_nwElement0_92, 0)
			(_nw443).ArraySet1(_2626_v56, 1)
			(_nw443).ArraySet1(func() _dafny.Map {
				var _coll76 = _dafny.NewMapBuilder()
				_ = _coll76
				for _iter82 := _dafny.Iterate((_2627_v58).Elements()); ; {
					_compr_76, _ok82 := _iter82()
					if !_ok82 {
						break
					}
					var _2629_v57 _dafny.Int
					_2629_v57 = interface{}(_compr_76).(_dafny.Int)
					if (_2627_v58).Contains(_2629_v57) {
						_coll76.Add((_2629_v57).Plus((_dafny.Zero).Minus(_2621_v51)), p2)
					}
				}
				return _coll76.ToMap()
			}(), 2)
			(_nw443).ArraySet1(_2626_v56, 3)
			(_nw443).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_2621_v51)).Cardinality()), p0), 4)
			(_nw443).ArraySet1(_2626_v56, 5)
			(_nw443).ArraySet1(_2626_v56, 6)
			(_nw443).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2621_v51, (_this).F6())).Update(_2621_v51, (_this).F6()), 7)
			(_nw443).ArraySet1(_2626_v56, 8)
			(_nw443).ArraySet1(_2626_v56, 9)
			(_nw443).ArraySet1(_2626_v56, 10)
			(_nw443).ArraySet1(_2626_v56, 11)
			(_nw443).ArraySet1(_2626_v56, 12)
			(_nw443).ArraySet1(_2626_v56, 13)
			(_nw443).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2621_v51, p2), 14)
			(_nw443).ArraySet1(_2626_v56, 15)
			(_nw443).ArraySet1(_2626_v56, 16)
			(_nw443).ArraySet1(_2626_v56, 17)
			(_nw443).ArraySet1(_2626_v56, 18)
			(_nw443).ArraySet1(_2626_v56, 19)
			(_nw443).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2621_v51, p2), 20)
			(_nw443).ArraySet1(_2626_v56, 21)
			(_nw443).ArraySet1(_2626_v56, 22)
			(_nw443).ArraySet1((_2626_v56).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2621_v51, p2)).Cardinality(), p2), 23)
			(_nw443).ArraySet1(_2626_v56, 24)
			_2628_v59 = _nw443
			var _index450 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_2625_v55), 0))
			_ = _index450
			(_2625_v55).ArraySet1(_2628_v59, (_index450).Int())
			var _index451 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_2625_v55), 0))
			_ = _index451
			(_2625_v55).ArraySet1(_2628_v59, (_index451).Int())
		} else {
			var _2630___mcc_h3 D15 = _source35.Get_().(D15_DC29).Cf38
			_ = _2630___mcc_h3
			var _2631_cf38 D15 = _2630___mcc_h3
			_ = _2631_cf38
			var _2632_v60 _dafny.Map
			_ = _2632_v60
			_2632_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-135), _2553_v1)
			var _2633_v61 _dafny.Int
			_ = _2633_v61
			_2633_v61 = _dafny.IntOfInt64(202)
			(globalState).F1 = Companion_Default___.SafeModuloInt((_2632_v60).Cardinality(), _2633_v61)
			var _2634_v62 _dafny.Array
			_ = _2634_v62
			var _nw444 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw444
			_2634_v62 = _nw444
			var _index452 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))
			_ = _index452
			(_2634_v62).ArraySet1(_2633_v61, (_index452).Int())
			var _2635_v63 D20
			_ = _2635_v63
			_2635_v63 = Companion_D20_.Create_DC43_(_2633_v61, p2, _2633_v61)
			var _index453 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))
			_ = _index453
			var _rhs507 _dafny.Array = _this.F7
			_ = _rhs507
			var _rhs508 _dafny.Int = _2633_v61
			_ = _rhs508
			var _rhs509 _dafny.Int = (_2635_v63).Dtor_cf58()
			_ = _rhs509
			var _lhs335 *C8 = _this
			_ = _lhs335
			var _lhs336 _dafny.Array = _2634_v62
			_ = _lhs336
			var _lhs337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))
			_ = _lhs337
			_lhs335.F7 = _rhs507
			(_lhs336).ArraySet1(_rhs508, (_lhs337).Int())
			_2633_v61 = _rhs509
			_2633_v61 = (func() _dafny.Int {
				if p2 {
					return (_dafny.Zero).Minus((_this).Fm7(true, globalState))
				}
				return _dafny.IntOfInt64(956)
			})()
			var _2636_v64 _dafny.Map
			_ = _2636_v64
			_2636_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_2634_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))).Int()).(_dafny.Int)), (_2553_v1).Select((Companion_Default___.SafeIndex(_2633_v61, _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32()).(bool))
			if (p2) && ((func() bool {
				if (_2636_v64).Contains(_dafny.IntOfInt64(176)) {
					return (_2636_v64).Get(_dafny.IntOfInt64(176)).(bool)
				}
				return p2
			})()) {
				(_this).F7 = _this.F7
				_2553_v1 = _2553_v1
				var _2637_v65 bool
				_ = _2637_v65
				_2637_v65 = false
				_2637_v65 = !(!(!((_this).F6())))
				var _2638_v66 _dafny.Array
				_ = _2638_v66
				_2638_v66 = _this.F7
				_2638_v66 = _2638_v66
				_2637_v65 = p2
			} else {
				(globalState).F1 = _2633_v61
				var _arr25 _dafny.Array = _this.F7
				_ = _arr25
				var _index454 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_this.F7), 0))
				_ = _index454
				(_arr25).ArraySet1(p2, (_index454).Int())
				var _2639_v67 _dafny.Map
				_ = _2639_v67
				_2639_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(118))
				var _2640_v68 _dafny.MultiSet
				_ = _2640_v68
				_2640_v68 = _dafny.MultiSetOf((_this).F5(), (_this).F5())
				var _arr26 _dafny.Array = _this.F7
				_ = _arr26
				var _index455 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_this.F7), 0))
				_ = _index455
				var _rhs510 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_2553_v1, (Companion_Default___.SafeIndex(_2633_v61, _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32(), (_this).F6())
				_ = _rhs510
				var _rhs511 _dafny.CodePoint = _2552_v0
				_ = _rhs511
				var _rhs512 bool = p2
				_ = _rhs512
				var _rhs513 _dafny.Map = _2639_v67
				_ = _rhs513
				var _rhs514 _dafny.Int = ((_this).Fm11(_2633_v61, p0, (_2640_v68).Cardinality(), globalState)).Cardinality()
				_ = _rhs514
				var _lhs338 _dafny.Array = _this.F7
				_ = _lhs338
				var _lhs339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_this.F7), 0))
				_ = _lhs339
				var _lhs340 *GlobalState = globalState
				_ = _lhs340
				_2553_v1 = _rhs510
				_2552_v0 = _rhs511
				(_lhs338).ArraySet1(_rhs512, (_lhs339).Int())
				_2639_v67 = _rhs513
				_lhs340.F1 = _rhs514
				(globalState).F1 = (_2634_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))).Int()).(_dafny.Int)
				var _2641_v69 _dafny.Map
				_ = _2641_v69
				_2641_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), (_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(996), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool))
				_2641_v69 = _2641_v69
				(globalState).F1 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if p0 {
						return _dafny.IntOfUint32((_dafny.SeqOf((_2634_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(14), _dafny.ArrayLen((_2634_v62), 0))).Int()).(_dafny.Int))).Cardinality())
					}
					return _2633_v61
				})(), (p1).Cardinality())
			}
		}
		var _2642_v70 _dafny.Int
		_ = _2642_v70
		_2642_v70 = _dafny.IntOfInt64(712)
		(globalState).F1 = _2642_v70
		var _2643_v71 _dafny.Sequence
		_ = _2643_v71
		_2643_v71 = _dafny.UnicodeSeqOfUtf8Bytes("fadnxqgcp")
		var _2644_v72 _dafny.Sequence
		_ = _2644_v72
		_2644_v72 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_2643_v71, _2643_v71), _2643_v71)
		_2644_v72 = _dafny.Companion_Sequence_.Concatenate(_2644_v72, _2644_v72)
		var _2645_v73 _dafny.MultiSet
		_ = _2645_v73
		_2645_v73 = _dafny.MultiSetOf(_2643_v71)
		var _2646_v74 _dafny.Map
		_ = _2646_v74
		_2646_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()) || ((_2553_v1).Select((Companion_Default___.SafeIndex(_2642_v70, _dafny.IntOfUint32((_2553_v1).Cardinality()))).Uint32()).(bool)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(261))).Uint32(), func(coer136 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg136 _dafny.Int) interface{} {
				return coer136(arg136)
			}
		}(func(_2647_i4 _dafny.Int) _dafny.CodePoint {
			return (_this).F5()
		})))
		var _rhs515 _dafny.Int = _2642_v70
		_ = _rhs515
		var _rhs516 _dafny.Sequence = (func() _dafny.Sequence {
			if (_2646_v74).Contains((_this).F6()) {
				return (_2646_v74).Get((_this).F6()).(_dafny.Sequence)
			}
			return _2643_v71
		})()
		_ = _rhs516
		var _rhs517 _dafny.Int = (_2642_v70).Minus(_2642_v70)
		_ = _rhs517
		var _rhs518 _dafny.MultiSet = (_2645_v73).Intersection(((_dafny.MultiSetOf(_2643_v71, _2643_v71)).Update(_dafny.UnicodeSeqOfUtf8Bytes("anuxkuyu"), Companion_Default___.Abs(_dafny.IntOfInt64(-918)))).Union(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("xbkjrf"))))
		_ = _rhs518
		var _lhs341 *GlobalState = globalState
		_ = _lhs341
		_lhs341.F1 = _rhs515
		_2643_v71 = _rhs516
		_2642_v70 = _rhs517
		_2645_v73 = _rhs518
		var _2648_v75 _dafny.Set
		_ = _2648_v75
		_2648_v75 = _dafny.SetOf(_this.F7, _this.F7)
		_2648_v75 = (func() _dafny.Set {
			if p0 {
				return _2648_v75
			}
			return _dafny.SetOf(_this.F7)
		})()
	}
}
func (_this *C8) M6(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Sequence, bool, bool) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _2649_v0 _dafny.Sequence
		_ = _2649_v0
		_2649_v0 = _dafny.SeqOf(p0)
		r1 = !_dafny.Companion_Sequence_.Equal(_2649_v0, _2649_v0)
		var _2650_v1 _dafny.Int
		_ = _2650_v1
		_2650_v1 = _dafny.IntOfInt64(556)
		var _2651_v2 D21
		_ = _2651_v2
		_2651_v2 = Companion_D21_.Create_DC46_(p0, p1, _2650_v1)
		r2 = (_2651_v2).Dtor_cf61()
		var _2652_v3 _dafny.Array
		_ = _2652_v3
		var _nwElement0_93 _dafny.Int = _2650_v1
		_ = _nwElement0_93
		var _nw445 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_93, nil, _dafny.IntOfInt64(2))
		_ = _nw445
		(_nw445).ArraySet1(_nwElement0_93, 0)
		(_nw445).ArraySet1(_2650_v1, 1)
		_2652_v3 = _nw445
		var _index456 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))
		_ = _index456
		(_2652_v3).ArraySet1(_2650_v1, (_index456).Int())
		var _index457 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))
		_ = _index457
		(_2652_v3).ArraySet1(Companion_Default___.SafeModuloInt((_2650_v1).Plus(_2650_v1), _dafny.IntOfInt64(33)), (_index457).Int())
		r2 = (_this).F6()
		var _2653_v4 _dafny.Map
		_ = _2653_v4
		_2653_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _2654_v5 D15
		_ = _2654_v5
		_2654_v5 = Companion_D15_.Create_DC28_((func() bool {
			if (_2653_v4).Contains(Companion_Default___.Fm0(p1, globalState)) {
				return (_2653_v4).Get(Companion_Default___.Fm0(p1, globalState)).(bool)
			}
			return (_this).F6()
		})(), (_this).F6())
		var _2655_i0 _dafny.Int
		_ = _2655_i0
		_2655_i0 = _dafny.Zero
		{
			for (_2654_v5).Dtor_cf37() {
				{
					if (_2655_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L24
					}
					_2655_i0 = (_2655_i0).Plus(_dafny.One)
					var _2656_v6 D0
					_ = _2656_v6
					_2656_v6 = Companion_D0_.Create_DC0_(p1)
					var _2657_v7 *C3
					_ = _2657_v7
					var _nw446 *C3 = New_C3_()
					_ = _nw446
					_nw446.Ctor__(p1, (_this).Fm9((_dafny.Zero).Minus(Companion_Default___.Fm3(p1, p0, globalState)), _2656_v6, Companion_Default___.Fm0(p0, globalState), _2650_v1, globalState))
					_2657_v7 = _nw446
					var _2658_v8 _dafny.Sequence
					_ = _2658_v8
					_2658_v8 = _dafny.SeqOf(_2657_v7)
					var _2659_v9 _dafny.Sequence
					_ = _2659_v9
					_2659_v9 = _dafny.SeqOf((func() *C3 {
						if !(p0) {
							return _2657_v7
						}
						return _2657_v7
					})(), (_2658_v8).Select((Companion_Default___.SafeIndex((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2658_v8).Cardinality()))).Uint32()).(*C3), _2657_v7)
					var _2660_v10 _dafny.Sequence
					_ = _2660_v10
					_2660_v10 = _dafny.SeqOf((_dafny.Zero).Minus((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), _2650_v1, (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int))
					_2657_v7 = (_2659_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2660_v10).Cardinality()), _dafny.IntOfUint32((_2659_v9).Cardinality()))).Uint32()).(*C3)
					var _2661_v11 _dafny.MultiSet
					_ = _2661_v11
					_2661_v11 = _dafny.MultiSetOf(_2650_v1)
					var _2662_v12 _dafny.Set
					_ = _2662_v12
					_2662_v12 = _dafny.SetOf((_this).Fm4(_2649_v0, _dafny.IntOfInt64(914), _2650_v1, _2661_v11, globalState))
					_2660_v10 = _dafny.SeqOf(_2650_v1, _2650_v1, ((_2662_v12).Union(_2662_v12)).Cardinality(), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int))
					var _index458 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))
					_ = _index458
					(_2652_v3).ArraySet1(Companion_Default___.Fm3(!((_2657_v7).F14()), p0, globalState), (_index458).Int())
					r2 = (_2662_v12).IsDisjointFrom(_2662_v12)
					goto C24
				}
			C24:
			}
			goto L24
		}
	L24:
		var _2663_v13 D6
		_ = _2663_v13
		_2663_v13 = Companion_D6_.Create_DC14_((_this).F6())
		var _2664_i1 _dafny.Int
		_ = _2664_i1
		_2664_i1 = _dafny.Zero
		{
			for ((func() D6 {
				if (_this).F6() {
					return Companion_D6_.Create_DC14_(p1)
				}
				return _2663_v13
			})()).Dtor_cf15() {
				{
					if (_2664_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L25
					}
					_2664_i1 = (_2664_i1).Plus(_dafny.One)
					var _2665_v14 D0
					_ = _2665_v14
					_2665_v14 = Companion_D0_.Create_DC1_()
					var _2666_v15 _dafny.Map
					_ = _2666_v15
					_2666_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2665_v14, _dafny.IntOfInt64(25))
					_2666_v15 = (_2666_v15).Merge(_2666_v15)
					var _arr27 _dafny.Array = _this.F7
					_ = _arr27
					var _index459 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
					_ = _index459
					(_arr27).ArraySet1(p1, (_index459).Int())
					var _arr28 _dafny.Array = _this.F7
					_ = _arr28
					var _index460 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
					_ = _index460
					(_arr28).ArraySet1((func() bool {
						if !((_this).F6()) {
							return (_this).F6()
						}
						return p1
					})(), (_index460).Int())
					var _2667_v16 _dafny.Array
					_ = _2667_v16
					var _nw447 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
					_ = _nw447
					_2667_v16 = _nw447
					var _2668_v17 _dafny.Map
					_ = _2668_v17
					_2668_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2667_v16, !((_this).F6()))
					_2668_v17 = (_2668_v17).Update(_2667_v16, p0)
					var _2669_v18 _dafny.MultiSet
					_ = _2669_v18
					_2669_v18 = _dafny.MultiSetOf(true, p0)
					var _2670_v19 D15
					_ = _2670_v19
					_2670_v19 = Companion_D15_.Create_DC27_((_2669_v18).Union(_dafny.MultiSetFromSeq(_2649_v0)))
					var _source36 D15 = _2670_v19
					_ = _source36
					if _source36.Is_DC28() {
						var _2671___mcc_h0 bool = _source36.Get_().(D15_DC28).Cf36
						_ = _2671___mcc_h0
						var _2672___mcc_h1 bool = _source36.Get_().(D15_DC28).Cf37
						_ = _2672___mcc_h1
						var _2673_cf37 bool = _2672___mcc_h1
						_ = _2673_cf37
						var _2674_cf36 bool = _2671___mcc_h0
						_ = _2674_cf36
						var _2675_v20 _dafny.Map
						_ = _2675_v20
						_2675_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2674_cf36, _dafny.MultiSetOf(_2650_v1, _2650_v1, (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)))
						var _2676_v21 _dafny.MultiSet
						_ = _2676_v21
						_2676_v21 = _dafny.MultiSetOf((_2675_v20).Cardinality())
						var _2677_v22 _dafny.MultiSet
						_ = _2677_v22
						_2677_v22 = _dafny.MultiSetOf((_2676_v21).Cardinality())
						var _2678_v24 _dafny.Map
						_ = _2678_v24
						_2678_v24 = func() _dafny.Map {
							var _coll77 = _dafny.NewMapBuilder()
							_ = _coll77
							for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(138), _dafny.IntOfInt64(86))); ; {
								_compr_77, _ok83 := _iter83()
								if !_ok83 {
									break
								}
								var _2679_v23 _dafny.Int
								_2679_v23 = interface{}(_compr_77).(_dafny.Int)
								if ((_dafny.IntOfInt64(138)).Cmp(_2679_v23) <= 0) && ((_2679_v23).Cmp(_dafny.IntOfInt64(86)) < 0) {
									_coll77.Add((_2679_v23).Minus(_dafny.IntOfInt64(-264)), false)
								}
							}
							return _coll77.ToMap()
						}()
						var _2680_v25 _dafny.Map
						_ = _2680_v25
						_2680_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2676_v21).Contains(_2650_v1), _2678_v24)
						_2680_v25 = (_2680_v25).Update(!(true), (func() _dafny.Map {
							if p1 {
								return _2678_v24
							}
							return _2678_v24
						})())
						var _2681_v26 _dafny.Set
						_ = _2681_v26
						_2681_v26 = _dafny.SetOf(_2673_cf37)
						var _2682_v27 _dafny.Sequence
						_ = _2682_v27
						_2682_v27 = _dafny.SeqOf(_2681_v26)
						var _2683_v28 _dafny.Map
						_ = _2683_v28
						_2683_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2673_cf37, _dafny.IntOfUint32((_2649_v0).Cardinality()))
						var _2684_v29 _dafny.Array
						_ = _2684_v29
						var _nwElement0_94 _dafny.Set = _dafny.SetOf(true)
						_ = _nwElement0_94
						var _nw448 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_94, nil, _dafny.IntOfInt64(8))
						_ = _nw448
						(_nw448).ArraySet1(_nwElement0_94, 0)
						(_nw448).ArraySet1((_dafny.SetOf(p1)).Difference(_2681_v26), 1)
						(_nw448).ArraySet1(_2681_v26, 2)
						(_nw448).ArraySet1((_2682_v27).Select((Companion_Default___.SafeIndex((_2683_v28).Cardinality(), _dafny.IntOfUint32((_2682_v27).Cardinality()))).Uint32()).(_dafny.Set), 3)
						(_nw448).ArraySet1(Companion_Default___.Fm17(_2650_v1, p1, globalState), 4)
						(_nw448).ArraySet1(Companion_Default___.Fm17(_2650_v1, (_this).F6(), globalState), 5)
						(_nw448).ArraySet1((_2681_v26).Union(_2681_v26), 6)
						(_nw448).ArraySet1(Companion_Default___.Fm49(globalState), 7)
						_2684_v29 = _nw448
						var _2685_v30 _dafny.Array
						_ = _2685_v30
						var _nw449 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(7))
						_ = _nw449
						_2685_v30 = _nw449
						var _index461 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_2685_v30), 0))
						_ = _index461
						(_2685_v30).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-409))).Uint32(), func(coer137 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg137 _dafny.Int) interface{} {
								return coer137(arg137)
							}
						}(func(_2686_i2 _dafny.Int) _dafny.CodePoint {
							return (_this).F5()
						})), (_index461).Int())
						var _2687_v31 _dafny.Sequence
						_ = _2687_v31
						_2687_v31 = _dafny.UnicodeSeqOfUtf8Bytes("ayjjit")
						var _index462 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_2685_v30), 0))
						_ = _index462
						var _rhs519 _dafny.Array = _2684_v29
						_ = _rhs519
						var _rhs520 _dafny.Sequence = _2687_v31
						_ = _rhs520
						var _lhs342 _dafny.Array = _2685_v30
						_ = _lhs342
						var _lhs343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_2685_v30), 0))
						_ = _lhs343
						_2684_v29 = _rhs519
						(_lhs342).ArraySet1(_rhs520, (_lhs343).Int())
						(globalState).F1 = (func() _dafny.Int {
							if !((p0) == (_2673_cf37)) {
								return (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)
							}
							return Companion_Default___.SafeDivisionInt((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm1(_dafny.IntOfInt64(597), _2674_cf36, (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))
						})()
						var _index463 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_2685_v30), 0))
						_ = _index463
						(_2685_v30).ArraySet1(_2687_v31, (_index463).Int())
					} else if _source36.Is_DC27() {
						var _2688___mcc_h2 _dafny.MultiSet = _source36.Get_().(D15_DC27).Cf35
						_ = _2688___mcc_h2
						var _2689_cf35 _dafny.MultiSet = _2688___mcc_h2
						_ = _2689_cf35
						r1 = false
						(globalState).F1 = (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)
						var _2690_v32 _dafny.Sequence
						_ = _2690_v32
						_2690_v32 = _dafny.SeqOf((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int))
						var _2691_v33 _dafny.Array
						_ = _2691_v33
						var _nwElement0_95 bool = !((func() bool {
							if (_this).F6() {
								return p1
							}
							return p0
						})())
						_ = _nwElement0_95
						var _nw450 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_95, nil, _dafny.IntOfInt64(13))
						_ = _nw450
						(_nw450).ArraySet1(_nwElement0_95, 0)
						(_nw450).ArraySet1(p1, 1)
						(_nw450).ArraySet1((_2649_v0).Select((Companion_Default___.SafeIndex(_2650_v1, _dafny.IntOfUint32((_2649_v0).Cardinality()))).Uint32()).(bool), 2)
						(_nw450).ArraySet1(p1, 3)
						(_nw450).ArraySet1(p1, 4)
						(_nw450).ArraySet1((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), 5)
						(_nw450).ArraySet1(_dafny.Companion_Sequence_.Equal(_2690_v32, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer138 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg138 _dafny.Int) interface{} {
								return coer138(arg138)
							}
						}((func(_2692_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2693_i3 _dafny.Int) _dafny.Int {
								return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2692_v1, (_this).F5())).Cardinality()
							}
						})(_2650_v1)))), 6)
						(_nw450).ArraySet1((_this).F6(), 7)
						(_nw450).ArraySet1(!((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool)) || (p0), 8)
						(_nw450).ArraySet1((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), 9)
						(_nw450).ArraySet1((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), 10)
						(_nw450).ArraySet1(p0, 11)
						(_nw450).ArraySet1(!_dafny.Companion_Sequence_.Contains(_2690_v32, _2650_v1), 12)
						_2691_v33 = _nw450
						var _2694_v34 _dafny.Map
						_ = _2694_v34
						_2694_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), ((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)).Plus(_2650_v1))
						var _index464 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))
						_ = _index464
						var _arr29 _dafny.Array = _this.F7
						_ = _arr29
						var _index465 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
						_ = _index465
						var _rhs521 _dafny.Int = _2650_v1
						_ = _rhs521
						var _rhs522 _dafny.Array = _2691_v33
						_ = _rhs522
						var _rhs523 _dafny.Array = _2691_v33
						_ = _rhs523
						var _rhs524 _dafny.Map = _2694_v34
						_ = _rhs524
						var _rhs525 bool = (_this).F6()
						_ = _rhs525
						var _lhs344 _dafny.Array = _2652_v3
						_ = _lhs344
						var _lhs345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))
						_ = _lhs345
						var _lhs346 _dafny.Array = _this.F7
						_ = _lhs346
						var _lhs347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
						_ = _lhs347
						(_lhs344).ArraySet1(_rhs521, (_lhs345).Int())
						_2691_v33 = _rhs522
						_2691_v33 = _rhs523
						_2694_v34 = _rhs524
						(_lhs346).ArraySet1(_rhs525, (_lhs347).Int())
						var _2695_v35 _dafny.Set
						_ = _2695_v35
						_2695_v35 = _dafny.SetOf((_this).F6())
						var _2696_v36 _dafny.Map
						_ = _2696_v36
						_2696_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2695_v35, (_this).F6())
						var _arr30 _dafny.Array = _this.F7
						_ = _arr30
						var _index466 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
						_ = _index466
						(_arr30).ArraySet1(!((_this).Fm10(_2669_v18, _2696_v36, ((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)).Minus(_2650_v1), globalState)), (_index466).Int())
					} else {
						var _2697___mcc_h3 D15 = _source36.Get_().(D15_DC29).Cf38
						_ = _2697___mcc_h3
						var _2698_cf38 D15 = _2697___mcc_h3
						_ = _2698_cf38
						var _2699_v37 _dafny.Sequence
						_ = _2699_v37
						_2699_v37 = _dafny.SeqOf(_2650_v1)
						var _2700_v38 _dafny.Map
						_ = _2700_v38
						_2700_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(709), _dafny.Companion_Sequence_.Update(_2699_v37, (Companion_Default___.SafeIndex((_this).Fm8(true, (_this).F6(), (_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_2699_v37).Cardinality()))).Uint32(), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)))
						var _2701_v39 _dafny.Sequence
						_ = _2701_v39
						_2701_v39 = _dafny.SeqOf((func() _dafny.Sequence {
							if (_2700_v38).Contains((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)) {
								return (_2700_v38).Get((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
							}
							return _2699_v37
						})(), _2699_v37)
						var _2702_v40 _dafny.Map
						_ = _2702_v40
						_2702_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2701_v39).Cardinality()), true)
						_2702_v40 = (_2702_v40).Update((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int), !((_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool)))
						var _2703_v41 _dafny.Set
						_ = _2703_v41
						_2703_v41 = _dafny.SetOf(_2650_v1)
						var _2704_v42 _dafny.Sequence
						_ = _2704_v42
						_2704_v42 = _dafny.UnicodeSeqOfUtf8Bytes("gxqw")
						var _2705_v43 _dafny.Map
						_ = _2705_v43
						_2705_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2703_v41).Cardinality(), (_this).Fm8((_2649_v0).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_2704_v42).Cardinality()), _dafny.IntOfUint32((_2649_v0).Cardinality()))).Uint32()).(bool), p1, (_this.F7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))).Int()).(bool), globalState))
						_2705_v43 = (_2705_v43).Update(_2650_v1, (_this).Fm8((_this).F6(), false, true, globalState))
						var _arr31 _dafny.Array = _this.F7
						_ = _arr31
						var _index467 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_this.F7), 0))
						_ = _index467
						(_arr31).ArraySet1((_this).F6(), (_index467).Int())
						r1 = false
					}
					goto C25
				}
			C25:
			}
			goto L25
		}
	L25:
		var _2706_v44 _dafny.Sequence
		_ = _2706_v44
		_2706_v44 = _dafny.SeqOf((_dafny.Zero).Minus((_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)), _2650_v1)
		r0 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2706_v44, _2706_v44), _dafny.SeqOf((_dafny.Zero).Minus(_2650_v1), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int))), (Companion_Default___.SafeIndex((_2650_v1).Times(_2650_v1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_2706_v44, _2706_v44), _dafny.SeqOf((_dafny.Zero).Minus(_2650_v1), (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int)))).Cardinality()))).Uint32(), _2650_v1)
		var _2707_v45 D11
		_ = _2707_v45
		_2707_v45 = Companion_D11_.Create_DC20_(p0, (_2652_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(402), _dafny.ArrayLen((_2652_v3), 0))).Int()).(_dafny.Int))
		r1 = (_2707_v45).Dtor_cf21()
		r2 = !(p1)
		return r0, r1, r2
	}
}
func (_this *C8) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
