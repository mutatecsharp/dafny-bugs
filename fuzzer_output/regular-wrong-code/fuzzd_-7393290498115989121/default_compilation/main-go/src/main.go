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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.MultiSet, globalState *GlobalState) _dafny.Int {
	return ((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fbkunt")).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ich")).Cardinality()), _dafny.IntOfInt64(910))).Cardinality()), _dafny.IntOfInt64(208)))).Cardinality())).Cardinality()))).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cgadqenod")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Set, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(!(!(false)))).IsProperSubsetOf(_dafny.SetOf(true, true, !(!(false)))), _dafny.IntOfInt64(646))
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false), _dafny.SeqOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-866), _dafny.IntOfInt64(9))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-866)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(9)) < 0) {
				_coll0.Add((_0_v0).Minus((_dafny.MultiSetOf(Companion_D2_.Create_DC7_(_dafny.SeqOf(false)))).Cardinality()))
			}
		}
		return _coll0.ToSet()
	}()).Union(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("k"), _dafny.IntOfInt64(696))).Cardinality()))).Difference((_dafny.SetOf(_dafny.IntOfInt64(267))).Intersection(_dafny.SetOf(_dafny.IntOfInt64(305))))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf((Companion_D9_.Create_DC30_(_dafny.CodePoint('q'), false, _dafny.IntOfInt64(-248), _dafny.SetOf(_dafny.IntOfInt64(588), _dafny.IntOfInt64(628), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("slllsh")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality(), _dafny.IntOfInt64(-740)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pbe")).Cardinality()))).Dtor_cf54()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(_dafny.IntOfInt64(-855))))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.MultiSet, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(620)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vsrbg")).Cardinality()), _dafny.IntOfInt64(617), (_dafny.MultiSetOf(true)).Cardinality())))).Intersection(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(481))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(false, !(false))).Cardinality()))
	})), _dafny.SeqOf(_dafny.IntOfInt64(-578))))).Intersection(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), true)).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(-82), _dafny.IntOfInt64(371))), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-616))))))
}
func (_static *CompanionStruct_Default___) Fm14(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-188), _dafny.IntOfInt64(202))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(330), _dafny.IntOfInt64(681))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(619))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(524)
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.CodePoint('a'), _dafny.CodePoint('d'), _dafny.CodePoint('a'))).Difference(_dafny.SetOf(_dafny.CodePoint('g'), _dafny.CodePoint('s'), _dafny.CodePoint('m'), _dafny.CodePoint('d'), _dafny.CodePoint('h')))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D0 {
	var _source0 D7 = (func() D7 {
		if false {
			return Companion_D7_.Create_DC23_(false, _dafny.IntOfInt64(842), false)
		}
		return Companion_D7_.Create_DC23_(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(-164))).Cardinality()), !(true))
	})()
	_ = _source0
	if _source0.Is_DC22() {
		var _3___mcc_h0 _dafny.Array = _source0.Get_().(D7_DC22).Cf37
		_ = _3___mcc_h0
		var _4___mcc_h1 _dafny.Sequence = _source0.Get_().(D7_DC22).Cf38
		_ = _4___mcc_h1
		var _5___mcc_h2 _dafny.Sequence = _source0.Get_().(D7_DC22).Cf39
		_ = _5___mcc_h2
		var _6___mcc_h3 _dafny.Array = _source0.Get_().(D7_DC22).Cf40
		_ = _6___mcc_h3
		var _7___mcc_h4 D5 = _source0.Get_().(D7_DC22).Cf41
		_ = _7___mcc_h4
		var _8_cf41 D5 = _7___mcc_h4
		_ = _8_cf41
		var _9_cf40 _dafny.Array = _6___mcc_h3
		_ = _9_cf40
		var _10_cf39 _dafny.Sequence = _5___mcc_h2
		_ = _10_cf39
		var _11_cf38 _dafny.Sequence = _4___mcc_h1
		_ = _11_cf38
		var _12_cf37 _dafny.Array = _3___mcc_h0
		_ = _12_cf37
		return Companion_D0_.Create_DC0_(true)
	} else if _source0.Is_DC23() {
		var _13___mcc_h5 bool = _source0.Get_().(D7_DC23).Cf42
		_ = _13___mcc_h5
		var _14___mcc_h6 _dafny.Int = _source0.Get_().(D7_DC23).Cf43
		_ = _14___mcc_h6
		var _15___mcc_h7 bool = _source0.Get_().(D7_DC23).Cf44
		_ = _15___mcc_h7
		var _16_cf44 bool = _15___mcc_h7
		_ = _16_cf44
		var _17_cf43 _dafny.Int = _14___mcc_h6
		_ = _17_cf43
		var _18_cf42 bool = _13___mcc_h5
		_ = _18_cf42
		return Companion_D0_.Create_DC0_(_18_cf42)
	} else {
		var _19___mcc_h8 _dafny.Map = _source0.Get_().(D7_DC21).Cf36
		_ = _19___mcc_h8
		var _20_cf36 _dafny.Map = _19___mcc_h8
		_ = _20_cf36
		return Companion_D0_.Create_DC0_(true)
	}
}
func (_static *CompanionStruct_Default___) Fm21(p0 D1, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	var _source1 D4 = Companion_D4_.Create_DC14_(true, (Companion_D7_.Create_DC23_(true, _dafny.IntOfInt64(781), true)).Dtor_cf44())
	_ = _source1
	if _source1.Is_DC13() {
		var _21___mcc_h0 _dafny.Sequence = _source1.Get_().(D4_DC13).Cf22
		_ = _21___mcc_h0
		var _22___mcc_h1 _dafny.Int = _source1.Get_().(D4_DC13).Cf23
		_ = _22___mcc_h1
		var _23___mcc_h2 bool = _source1.Get_().(D4_DC13).Cf24
		_ = _23___mcc_h2
		var _24___mcc_h3 _dafny.Array = _source1.Get_().(D4_DC13).Cf25
		_ = _24___mcc_h3
		var _25_cf25 _dafny.Array = _24___mcc_h3
		_ = _25_cf25
		var _26_cf24 bool = _23___mcc_h2
		_ = _26_cf24
		var _27_cf23 _dafny.Int = _22___mcc_h1
		_ = _27_cf23
		var _28_cf22 _dafny.Sequence = _21___mcc_h0
		_ = _28_cf22
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(339))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg2 _dafny.Int) interface{} {
				return coer2(arg2)
			}
		}((func(_29_cf23 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_30_i0 _dafny.Int) _dafny.Int {
				return _29_cf23
			}
		})(_27_cf23)))
	} else if _source1.Is_DC14() {
		var _31___mcc_h4 bool = _source1.Get_().(D4_DC14).Cf26
		_ = _31___mcc_h4
		var _32___mcc_h5 bool = _source1.Get_().(D4_DC14).Cf27
		_ = _32___mcc_h5
		var _33_cf27 bool = _32___mcc_h5
		_ = _33_cf27
		var _34_cf26 bool = _31___mcc_h4
		_ = _34_cf26
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-807))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}((func(_35_cf26 bool) func(_dafny.Int) _dafny.Int {
			return func(_36_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_35_cf26)).Cardinality()), (func() _dafny.Set {
					var _coll1 = _dafny.NewBuilder()
					_ = _coll1
					for _iter1 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(752), _dafny.IntOfInt64(428))); ; {
						_compr_1, _ok1 := _iter1()
						if !_ok1 {
							break
						}
						var _37_v0 _dafny.Int
						_37_v0 = interface{}(_compr_1).(_dafny.Int)
						if ((_dafny.IntOfInt64(752)).Cmp(_37_v0) <= 0) && ((_37_v0).Cmp(_dafny.IntOfInt64(428)) < 0) {
							_coll1.Add((_37_v0).Minus(_dafny.IntOfInt64(708)))
						}
					}
					return _coll1.ToSet()
				}()).Cardinality())).Cardinality())).Cardinality(), _dafny.Zero)).Cardinality()))).Cardinality())
			}
		})(_34_cf26)))
	} else if _source1.Is_DC12() {
		var _38___mcc_h6 _dafny.Array = _source1.Get_().(D4_DC12).Cf21
		_ = _38___mcc_h6
		var _39_cf21 _dafny.Array = _38___mcc_h6
		_ = _39_cf21
		return _dafny.SeqOf(_dafny.IntOfInt64(106), (_dafny.SetOf(_dafny.IntOfInt64(-163), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(937), _dafny.IntOfInt64(667))).Cardinality(), _dafny.IntOfInt64(795), (_dafny.SetOf(true, !(true), !(true))).Cardinality())).Cardinality())
	} else {
		var _40___mcc_h7 D4 = _source1.Get_().(D4_DC15).Cf28
		_ = _40___mcc_h7
		var _41_cf28 D4 = _40___mcc_h7
		_ = _41_cf28
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fdqjqr")).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(141))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(!(true), false)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pvuwwxen")).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(false, !(true), true)).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 bool, globalState *GlobalState) _dafny.Set {
	if !(!(true) || (true)) {
		return (_dafny.SetOf((_dafny.Zero).Minus((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nuqceuqr")).Cardinality()), _dafny.IntOfInt64(-359))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(870))).Cardinality()))).Union(func() _dafny.Set {
			var _coll2 = _dafny.NewBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(768), _dafny.IntOfInt64(-704))); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _42_v0 _dafny.Int
				_42_v0 = interface{}(_compr_2).(_dafny.Int)
				if ((_dafny.IntOfInt64(768)).Cmp(_42_v0) <= 0) && ((_42_v0).Cmp(_dafny.IntOfInt64(-704)) < 0) {
					_coll2.Add(Companion_Default___.SafeDivisionInt(_42_v0, _dafny.IntOfInt64(236)))
				}
			}
			return _coll2.ToSet()
		}())
	} else {
		return _dafny.SetOf(_dafny.IntOfInt64(569), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, !(false)))).Cardinality(), (func() _dafny.Set {
			var _coll3 = _dafny.NewBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())))).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _43_v1 _dafny.Int
				_43_v1 = interface{}(_compr_3).(_dafny.Int)
				if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality())))).Contains(_43_v1) {
					_coll3.Add(Companion_Default___.SafeModuloInt(_43_v1, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(688), _dafny.IntOfInt64(355), _dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(412))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg4 _dafny.Int) interface{} {
							return coer4(arg4)
						}
					}(func(_44_i0 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(384)
					}))).Cardinality()))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dp")).Cardinality()), _dafny.IntOfInt64(-318), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality()), _dafny.IntOfInt64(813))).Cardinality())))).Cardinality())))
				}
			}
			return _coll3.ToSet()
		}()).Cardinality(), _dafny.IntOfInt64(-904))
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 bool, p2 _dafny.Map, p3 _dafny.Set, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), (_dafny.Zero).Minus(_dafny.IntOfInt64(-778)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), (Companion_D18_.Create_DC56_(true, true, (_dafny.SetOf(_dafny.IntOfInt64(754), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kk")).Cardinality()))).Cardinality())).Dtor_cf92())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('g')
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	var _source2 D12 = Companion_D12_.Create_DC38_((Companion_D9_.Create_DC30_(_dafny.CodePoint('y'), true, (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-170))).Cardinality()), _dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-48)), _dafny.IntOfInt64(262))).Cardinality()), (_dafny.SetOf(true)).Cardinality()), _dafny.IntOfInt64(-96))).Dtor_cf52())
	_ = _source2
	if _source2.Is_DC38() {
		var _45___mcc_h0 _dafny.Int = _source2.Get_().(D12_DC38).Cf64
		_ = _45___mcc_h0
		var _46_cf64 _dafny.Int = _45___mcc_h0
		_ = _46_cf64
		return _dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('b'), _dafny.CodePoint('x')))
	} else if _source2.Is_DC39() {
		var _47___mcc_h1 _dafny.Int = _source2.Get_().(D12_DC39).Cf65
		_ = _47___mcc_h1
		var _48_cf65 _dafny.Int = _47___mcc_h1
		_ = _48_cf65
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(881))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_49_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(277))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg6 _dafny.Int) interface{} {
					return coer6(arg6)
				}
			}(func(_50_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('r')
			}))
		}))
	} else if _source2.Is_DC37() {
		var _51___mcc_h2 _dafny.MultiSet = _source2.Get_().(D12_DC37).Cf63
		_ = _51___mcc_h2
		var _52_cf63 _dafny.MultiSet = _51___mcc_h2
		_ = _52_cf63
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(114))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}(func(_53_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(87))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}(func(_54_i3 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqOf((Companion_D9_.Create_DC30_(_dafny.CodePoint('j'), true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality(), func() _dafny.Set {
				var _coll4 = _dafny.NewBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(750), _dafny.IntOfInt64(208))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _55_v0 _dafny.Int
					_55_v0 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(750)).Cmp(_55_v0) <= 0) && ((_55_v0).Cmp(_dafny.IntOfInt64(208)) < 0) {
						_coll4.Add(Companion_Default___.SafeModuloInt(_55_v0, (func() _dafny.Set {
							var _coll5 = _dafny.NewBuilder()
							_ = _coll5
							for _iter5 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()), true)).Keys().Elements()); ; {
								_compr_5, _ok5 := _iter5()
								if !_ok5 {
									break
								}
								var _56_v1 _dafny.Int
								_56_v1 = interface{}(_compr_5).(_dafny.Int)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality())).Cardinality()), true)).Contains(_56_v1) {
									_coll5.Add(Companion_Default___.SafeModuloInt(_56_v1, _dafny.IntOfInt64(881)))
								}
							}
							return _coll5.ToSet()
						}()).Cardinality()))
					}
				}
				return _coll4.ToSet()
			}(), _dafny.IntOfInt64(-338))).Dtor_cf50())
		})))
	} else {
		var _57___mcc_h3 D12 = _source2.Get_().(D12_DC40).Cf66
		_ = _57___mcc_h3
		var _58_cf66 D12 = _57___mcc_h3
		_ = _58_cf66
		return _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(874))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_59_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		})))
	}
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vkgblbh"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(122))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_60_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	})))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('a'), false)).Cardinality()).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hyccnriii")).Cardinality())), (_dafny.IntOfInt64(-573)).Minus(_dafny.IntOfInt64(618)))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true, !(true))))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(_dafny.IntOfInt64(-690))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(-639), (_dafny.MultiSetOf(_dafny.SeqOf(false), _dafny.SeqOf(false))).Cardinality(), _dafny.IntOfInt64(-736), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-596))).Uint32(), func(coer11 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_61_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cmgn")).Cardinality()))).Cardinality())
	}))).Cardinality()), _dafny.IntOfInt64(465)))
}
func (_static *CompanionStruct_Default___) Fm31(p0 bool, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ymdokrv"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(957))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_62_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	})))).Cardinality())), _dafny.CodePoint('c'))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Sequence, p1 _dafny.MultiSet, globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.MultiSetFromSeq(_dafny.SeqOf(false, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.MultiSetOf(false)))).Keys().Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _63_v0 _dafny.CodePoint
			_63_v0 = interface{}(_compr_6).(_dafny.CodePoint)
			if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('b'), _dafny.MultiSetFromSeq(_dafny.SeqOf(false, true)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), _dafny.MultiSetOf(false)))).Contains(_63_v0) {
				_coll6.Add(_63_v0)
			}
		}
		return _coll6.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(!(false)), !(!(true))), _dafny.SeqOf(true, true)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(true, false, true, false))))
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(((func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_64_i0 _dafny.Int) _dafny.Sequence {
				return (Companion_D22_.Create_DC70_(true, !(false), _dafny.UnicodeSeqOfUtf8Bytes("rgiu"), (_dafny.MultiSetOf(_dafny.IntOfInt64(331), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ia")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-47), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}(func(_65_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))).Cardinality()))).Cardinality())).Dtor_cf114()
			}))).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _66_v0 _dafny.Sequence
				_66_v0 = interface{}(_compr_8).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_64_i0 _dafny.Int) _dafny.Sequence {
					return (Companion_D22_.Create_DC70_(true, !(false), _dafny.UnicodeSeqOfUtf8Bytes("rgiu"), (_dafny.MultiSetOf(_dafny.IntOfInt64(331), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ia")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-47), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg16 _dafny.Int) interface{} {
							return coer16(arg16)
						}
					}(func(_65_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality()))).Cardinality())).Dtor_cf114()
				})), _66_v0) {
					_coll8.Add(_66_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true, !(true)))).Cardinality()))
				}
			}
			return _coll8.ToMap()
		}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("elx"), _dafny.IntOfInt64(919)))).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _67_v1 _dafny.Sequence
			_67_v1 = interface{}(_compr_7).(_dafny.Sequence)
			if ((func() _dafny.Map {
				var _coll9 = _dafny.NewMapBuilder()
				_ = _coll9
				for _iter9 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg17 _dafny.Int) interface{} {
						return coer17(arg17)
					}
				}(func(_64_i0 _dafny.Int) _dafny.Sequence {
					return (Companion_D22_.Create_DC70_(true, !(false), _dafny.UnicodeSeqOfUtf8Bytes("rgiu"), (_dafny.MultiSetOf(_dafny.IntOfInt64(331), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ia")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-47), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}(func(_65_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					}))).Cardinality()))).Cardinality())).Dtor_cf114()
				}))).Elements()); ; {
					_compr_9, _ok9 := _iter9()
					if !_ok9 {
						break
					}
					var _66_v0 _dafny.Sequence
					_66_v0 = interface{}(_compr_9).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg19 _dafny.Int) interface{} {
							return coer19(arg19)
						}
					}(func(_64_i0 _dafny.Int) _dafny.Sequence {
						return (Companion_D22_.Create_DC70_(true, !(false), _dafny.UnicodeSeqOfUtf8Bytes("rgiu"), (_dafny.MultiSetOf(_dafny.IntOfInt64(331), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ia")).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-47), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg20 _dafny.Int) interface{} {
								return coer20(arg20)
							}
						}(func(_65_i1 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('l')
						}))).Cardinality()))).Cardinality())).Dtor_cf114()
					})), _66_v0) {
						_coll9.Add(_66_v0, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(true, !(true)))).Cardinality()))
					}
				}
				return _coll9.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("elx"), _dafny.IntOfInt64(919)))).Contains(_67_v1) {
				_coll7.Add(_67_v1)
			}
		}
		return _coll7.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D0_.Create_DC1_(false)).Dtor_cf1(), false)).Cardinality(), Companion_D5_.Create_DC18_(false))).Merge((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(527), _dafny.IntOfInt64(315))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _68_v0 _dafny.Int
			_68_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(527)).Cmp(_68_v0) <= 0) && ((_68_v0).Cmp(_dafny.IntOfInt64(315)) < 0) {
				_coll10.Add((_68_v0).Plus(_dafny.IntOfInt64(909)), Companion_D5_.Create_DC18_(!(false)))
			}
		}
		return _coll10.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("niflmco")).Cardinality()), Companion_D5_.Create_DC18_(true))))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.MultiSet, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(510))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false))).Cardinality(), _dafny.IntOfInt64(940))).Cardinality()))), _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xcail")).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(954), _dafny.IntOfInt64(68))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _69_v0 _dafny.Int
			_69_v0 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(954)).Cmp(_69_v0) <= 0) && ((_69_v0).Cmp(_dafny.IntOfInt64(68)) < 0) {
				_coll11.Add((_69_v0).Plus((_dafny.SetOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('f'))).Cardinality())).Cardinality()))
			}
		}
		return _coll11.ToMap()
	}()).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Set, p1 _dafny.Map, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC1_(!(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), true)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tvidhsf")).Cardinality())) < 0))
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) _dafny.Sequence {
	var _source3 D5 = Companion_D5_.Create_DC17_(_dafny.SeqOf(_dafny.SeqOf(true, false, false)), false, _dafny.MultiSetOf(_dafny.IntOfInt64(372), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(142))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg21 _dafny.Int) interface{} {
			return coer21(arg21)
		}
	}(func(_70_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	}))).Cardinality()), (_dafny.SetOf(_dafny.CodePoint('w'), _dafny.CodePoint('n'), _dafny.CodePoint('j'), _dafny.CodePoint('r'))).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.MultiSetOf(_dafny.IntOfInt64(839)))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(327), false)).Cardinality()))
	_ = _source3
	if _source3.Is_DC17() {
		var _71___mcc_h0 _dafny.Sequence = _source3.Get_().(D5_DC17).Cf30
		_ = _71___mcc_h0
		var _72___mcc_h1 bool = _source3.Get_().(D5_DC17).Cf31
		_ = _72___mcc_h1
		var _73___mcc_h2 _dafny.MultiSet = _source3.Get_().(D5_DC17).Cf32
		_ = _73___mcc_h2
		var _74_cf32 _dafny.MultiSet = _73___mcc_h2
		_ = _74_cf32
		var _75_cf31 bool = _72___mcc_h1
		_ = _75_cf31
		var _76_cf30 _dafny.Sequence = _71___mcc_h0
		_ = _76_cf30
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_75_cf31), _dafny.SeqOf(_75_cf31))
	} else if _source3.Is_DC18() {
		var _77___mcc_h3 bool = _source3.Get_().(D5_DC18).Cf33
		_ = _77___mcc_h3
		var _78_cf33 bool = _77___mcc_h3
		_ = _78_cf33
		return _dafny.SeqOf(_78_cf33, true)
	} else if _source3.Is_DC16() {
		var _79___mcc_h4 _dafny.Array = _source3.Get_().(D5_DC16).Cf29
		_ = _79___mcc_h4
		var _80_cf29 _dafny.Array = _79___mcc_h4
		_ = _80_cf29
		return _dafny.SeqOf(true, true)
	} else {
		var _81___mcc_h5 D5 = _source3.Get_().(D5_DC19).Cf34
		_ = _81___mcc_h5
		var _82_cf34 D5 = _81___mcc_h5
		_ = _82_cf34
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, true, !(false), true, true), _dafny.SeqOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC30_(_dafny.CodePoint('w'), (func() bool {
		if true {
			return false
		}
		return true
	})(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfdadptf")).Cardinality()), true)).Cardinality()).Minus(_dafny.IntOfInt64(844)), func() _dafny.Set {
		var _coll12 = _dafny.NewBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(626), _dafny.IntOfInt64(349))); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _83_v0 _dafny.Int
			_83_v0 = interface{}(_compr_12).(_dafny.Int)
			if ((_dafny.IntOfInt64(626)).Cmp(_83_v0) <= 0) && ((_83_v0).Cmp(_dafny.IntOfInt64(349)) < 0) {
				_coll12.Add(Companion_Default___.SafeModuloInt(_83_v0, _dafny.IntOfInt64(168)))
			}
		}
		return _coll12.ToSet()
	}(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((Companion_D7_.Create_DC23_(false, _dafny.IntOfInt64(232), !(false))).Dtor_cf43(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ixq")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-233), true)).Cardinality())).Cardinality(), true)).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_D2_.Create_DC7_(_dafny.SeqOf(false, true)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(false), !(true), !(true))).Intersection(_dafny.SetOf(!(true)))).Union((_dafny.SetOf(!(!(!(false))))).Union(_dafny.SetOf(false)))
}
func (_static *CompanionStruct_Default___) Fm42(p0 bool, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D8 {
	var _source4 D11 = Companion_D11_.Create_DC33_(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false))))
	_ = _source4
	if _source4.Is_DC34() {
		var _84___mcc_h0 _dafny.MultiSet = _source4.Get_().(D11_DC34).Cf58
		_ = _84___mcc_h0
		var _85___mcc_h1 _dafny.MultiSet = _source4.Get_().(D11_DC34).Cf59
		_ = _85___mcc_h1
		var _86___mcc_h2 _dafny.Array = _source4.Get_().(D11_DC34).Cf60
		_ = _86___mcc_h2
		var _87___mcc_h3 D1 = _source4.Get_().(D11_DC34).Cf61
		_ = _87___mcc_h3
		var _88_cf61 D1 = _87___mcc_h3
		_ = _88_cf61
		var _89_cf60 _dafny.Array = _86___mcc_h2
		_ = _89_cf60
		var _90_cf59 _dafny.MultiSet = _85___mcc_h1
		_ = _90_cf59
		var _91_cf58 _dafny.MultiSet = _84___mcc_h0
		_ = _91_cf58
		return Companion_D8_.Create_DC26_(true, func() _dafny.Map {
			var _coll13 = _dafny.NewMapBuilder()
			_ = _coll13
			for _iter13 := _dafny.Iterate((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("di")).Cardinality())))).Elements()); ; {
				_compr_13, _ok13 := _iter13()
				if !_ok13 {
					break
				}
				var _92_v0 _dafny.Int
				_92_v0 = interface{}(_compr_13).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains((_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("di")).Cardinality())))), _92_v0) {
					_coll13.Add((_92_v0).Minus(_dafny.IntOfInt64(-650)), _dafny.UnicodeSeqOfUtf8Bytes("pc"))
				}
			}
			return _coll13.ToMap()
		}())
	} else if _source4.Is_DC35() {
		return Companion_D8_.Create_DC26_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(66), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(108))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_93_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))))
	} else if _source4.Is_DC33() {
		var _94___mcc_h4 _dafny.Set = _source4.Get_().(D11_DC33).Cf57
		_ = _94___mcc_h4
		var _95_cf57 _dafny.Set = _94___mcc_h4
		_ = _95_cf57
		return Companion_D8_.Create_DC26_(!(true), func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(528), _dafny.IntOfInt64(791))); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _96_v1 _dafny.Int
				_96_v1 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(528)).Cmp(_96_v1) <= 0) && ((_96_v1).Cmp(_dafny.IntOfInt64(791)) < 0) {
					_coll14.Add((_96_v1).Minus(_dafny.IntOfInt64(586)), _dafny.UnicodeSeqOfUtf8Bytes("hsqevyqb"))
				}
			}
			return _coll14.ToMap()
		}())
	} else {
		var _97___mcc_h5 D11 = _source4.Get_().(D11_DC36).Cf62
		_ = _97___mcc_h5
		var _98_cf62 D11 = _97___mcc_h5
		_ = _98_cf62
		if !(true) {
			return Companion_D8_.Create_DC26_(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-119), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_99_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			}))))
		} else {
			return Companion_D8_.Create_DC26_(false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-953), _dafny.UnicodeSeqOfUtf8Bytes("m")))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm43(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf((Companion_D0_.Create_DC0_(true)).Dtor_cf0(), !(true), true, true)), _dafny.SeqOf(_dafny.MultiSetOf(false), _dafny.MultiSetFromSeq(_dafny.SeqOf(true, false)), _dafny.MultiSetOf(false), _dafny.MultiSetFromSeq(_dafny.SeqOf(true)), _dafny.MultiSetOf(!(false)))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(true)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-251))).Uint32(), func(coer24 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_100_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))
	}))))
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) D2 {
	var _source5 D13 = Companion_D13_.Create_DC41_(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(689))))
	_ = _source5
	if _source5.Is_DC42() {
		var _101___mcc_h0 _dafny.CodePoint = _source5.Get_().(D13_DC42).Cf68
		_ = _101___mcc_h0
		var _102___mcc_h1 bool = _source5.Get_().(D13_DC42).Cf69
		_ = _102___mcc_h1
		var _103_cf69 bool = _102___mcc_h1
		_ = _103_cf69
		var _104_cf68 _dafny.CodePoint = _101___mcc_h0
		_ = _104_cf68
		return Companion_D2_.Create_DC8_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(609))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}((func(_105_cf68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_106_i0 _dafny.Int) _dafny.CodePoint {
				return _105_cf68
			}
		})(_104_cf68))))).Cardinality())), _dafny.IntOfInt64(-147), _dafny.IntOfInt64(-214))
	} else if _source5.Is_DC41() {
		var _107___mcc_h2 _dafny.Set = _source5.Get_().(D13_DC41).Cf67
		_ = _107___mcc_h2
		var _108_cf67 _dafny.Set = _107___mcc_h2
		_ = _108_cf67
		return Companion_D2_.Create_DC8_(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfInt64(117))).Cardinality()), _dafny.IntOfInt64(293), _dafny.IntOfInt64(-852))
	} else {
		var _109___mcc_h3 D13 = _source5.Get_().(D13_DC43).Cf70
		_ = _109___mcc_h3
		var _110_cf70 D13 = _109___mcc_h3
		_ = _110_cf70
		return Companion_D2_.Create_DC8_((_dafny.SetOf(_dafny.IntOfInt64(188))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, false, true)).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm45(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) D3 {
	var _source6 D9 = Companion_D9_.Create_DC30_(_dafny.CodePoint('q'), false, (_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())), _dafny.Zero)).Cardinality(), (_dafny.MultiSetOf(false, true, false)).Cardinality())).Cardinality(), _dafny.SetOf(_dafny.IntOfInt64(335), _dafny.IntOfInt64(218), _dafny.IntOfInt64(-979)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-175), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality()))).Cardinality())
	_ = _source6
	if _source6.Is_DC30() {
		var _111___mcc_h0 _dafny.CodePoint = _source6.Get_().(D9_DC30).Cf50
		_ = _111___mcc_h0
		var _112___mcc_h1 bool = _source6.Get_().(D9_DC30).Cf51
		_ = _112___mcc_h1
		var _113___mcc_h2 _dafny.Int = _source6.Get_().(D9_DC30).Cf52
		_ = _113___mcc_h2
		var _114___mcc_h3 _dafny.Set = _source6.Get_().(D9_DC30).Cf53
		_ = _114___mcc_h3
		var _115___mcc_h4 _dafny.Int = _source6.Get_().(D9_DC30).Cf54
		_ = _115___mcc_h4
		var _116_cf54 _dafny.Int = _115___mcc_h4
		_ = _116_cf54
		var _117_cf53 _dafny.Set = _114___mcc_h3
		_ = _117_cf53
		var _118_cf52 _dafny.Int = _113___mcc_h2
		_ = _118_cf52
		var _119_cf51 bool = _112___mcc_h1
		_ = _119_cf51
		var _120_cf50 _dafny.CodePoint = _111___mcc_h0
		_ = _120_cf50
		return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_cf51, _119_cf51))
	} else if _source6.Is_DC29() {
		var _121___mcc_h5 _dafny.Map = _source6.Get_().(D9_DC29).Cf49
		_ = _121___mcc_h5
		var _122_cf49 _dafny.Map = _121___mcc_h5
		_ = _122_cf49
		return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	} else {
		var _123___mcc_h6 D9 = _source6.Get_().(D9_DC31).Cf55
		_ = _123___mcc_h6
		var _124_cf55 D9 = _123___mcc_h6
		_ = _124_cf55
		return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
	}
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('p'), !(!(true)))).Merge(func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('e'), _dafny.CodePoint('a'))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _125_v0 _dafny.CodePoint
			_125_v0 = interface{}(_compr_15).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('d'), _dafny.CodePoint('g'), _dafny.CodePoint('b'), _dafny.CodePoint('e'), _dafny.CodePoint('a')), _125_v0) {
				_coll15.Add(_125_v0, true)
			}
		}
		return _coll15.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm47(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(false))).Difference(_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true, false, false)), _dafny.MultiSetOf(!(false))))).Union(func() _dafny.Set {
		var _coll16 = _dafny.NewBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, false, true)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(false, !(false)), _dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), true, true, true)))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _126_v0 _dafny.MultiSet
			_126_v0 = interface{}(_compr_16).(_dafny.MultiSet)
			if (_dafny.SetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false, false, true)), _dafny.MultiSetOf(true), _dafny.MultiSetOf(false, !(false)), _dafny.MultiSetFromSeq(_dafny.SeqOf(!(true), true, true, true)))).Contains(_126_v0) {
				_coll16.Add(_126_v0)
			}
		}
		return _coll16.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm48(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(!(true))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)).Cardinality()), (_dafny.IntOfInt64(871)).Plus(_dafny.IntOfInt64(758)))
}
func (_static *CompanionStruct_Default___) Fm49(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqOf(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC30_(_dafny.CodePoint('h'), false, (func() _dafny.Map {
		var _coll17 = _dafny.NewMapBuilder()
		_ = _coll17
		for _iter17 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(386)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))).Keys().Elements()); ; {
			_compr_17, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _127_v0 _dafny.Sequence
			_127_v0 = interface{}(_compr_17).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(386)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true))).Contains(_127_v0) {
				_coll17.Add(_127_v0, false)
			}
		}
		return _coll17.ToMap()
	}()).Cardinality(), func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.IntOfInt64(954))).Elements()); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _128_v1 _dafny.Int
			_128_v1 = interface{}(_compr_18).(_dafny.Int)
			if (_dafny.MultiSetOf(_dafny.IntOfInt64(954))).Contains(_128_v1) {
				_coll18.Add(Companion_Default___.SafeDivisionInt(_128_v1, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-253), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(880))).Cardinality())).Cardinality(), Companion_D1_.Create_DC4_((_dafny.SetOf(_dafny.IntOfInt64(-15))).Cardinality(), (_dafny.MultiSetOf(_dafny.IntOfInt64(974), _dafny.IntOfUint32((_dafny.SeqOf(!(false), false)).Cardinality()), (_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(106), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xpse")).Cardinality()))).Cardinality()))).Cardinality(), _dafny.IntOfInt64(162), _dafny.IntOfUint32((_dafny.SeqOf(false, false, !(false))).Cardinality()), _dafny.IntOfInt64(929))).Cardinality())))
			}
		}
		return _coll18.ToSet()
	}(), _dafny.IntOfInt64(277)))), _dafny.SeqOf(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(func() _dafny.Map {
		var _coll19 = _dafny.NewMapBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(140), _dafny.IntOfInt64(644))); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _129_v2 _dafny.Int
			_129_v2 = interface{}(_compr_19).(_dafny.Int)
			if ((_dafny.IntOfInt64(140)).Cmp(_129_v2) <= 0) && ((_129_v2).Cmp(_dafny.IntOfInt64(644)) < 0) {
				_coll19.Add((_129_v2).Plus(_dafny.IntOfInt64(293)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("phlpopw")).Cardinality()))
			}
		}
		return _coll19.ToMap()
	}())))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-447))).Uint32(), func(coer26 func(_dafny.Int) D9) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_130_i0 _dafny.Int) D9 {
		return Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC30_(_dafny.CodePoint('i'), true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality())).Cardinality(), _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true, true)).Cardinality())), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-825))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_131_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))).Cardinality()))))
	})))).Intersection(_dafny.SetOf(_dafny.SeqOf(Companion_D9_.Create_DC31_((Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC30_(_dafny.CodePoint('w'), false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false)).Cardinality())).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(271), _dafny.IntOfInt64(444)), _dafny.IntOfInt64(168))))).Dtor_cf55()), Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(733), _dafny.IntOfInt64(-491)))))))).Intersection((func() _dafny.Set {
		var _coll20 = _dafny.NewBuilder()
		_ = _coll20
		for _iter20 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqOf(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(394), _dafny.IntOfInt64(-187))); ; {
				_compr_21, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _132_v3 _dafny.Int
				_132_v3 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(394)).Cmp(_132_v3) <= 0) && ((_132_v3).Cmp(_dafny.IntOfInt64(-187)) < 0) {
					_coll21.Add(Companion_Default___.SafeDivisionInt(_132_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg28 _dafny.Int) interface{} {
							return coer28(arg28)
						}
					}(func(_133_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('f')
					}))).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfInt64(-559), _dafny.IntOfInt64(990))).Cardinality())
				}
			}
			return _coll21.ToMap()
		}()))), Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_134_i3 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))).Cardinality()))))))).Elements()); ; {
			_compr_20, _ok20 := _iter20()
			if !_ok20 {
				break
			}
			var _135_v4 _dafny.Sequence
			_135_v4 = interface{}(_compr_20).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqOf(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(func() _dafny.Map {
				var _coll22 = _dafny.NewMapBuilder()
				_ = _coll22
				for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(394), _dafny.IntOfInt64(-187))); ; {
					_compr_22, _ok22 := _iter22()
					if !_ok22 {
						break
					}
					var _132_v3 _dafny.Int
					_132_v3 = interface{}(_compr_22).(_dafny.Int)
					if ((_dafny.IntOfInt64(394)).Cmp(_132_v3) <= 0) && ((_132_v3).Cmp(_dafny.IntOfInt64(-187)) < 0) {
						_coll22.Add(Companion_Default___.SafeDivisionInt(_132_v3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg30 _dafny.Int) interface{} {
								return coer30(arg30)
							}
						}(func(_133_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('f')
						}))).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfInt64(-559), _dafny.IntOfInt64(990))).Cardinality())
					}
				}
				return _coll22.ToMap()
			}()))), Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_134_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Cardinality())))))), _135_v4) {
				_coll20.Add(_135_v4)
			}
		}
		return _coll20.ToSet()
	}()).Difference(_dafny.SetOf(_dafny.SeqOf(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-463), _dafny.IntOfInt64(601)))))))), Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC31_(Companion_D9_.Create_DC29_(func() _dafny.Map {
		var _coll23 = _dafny.NewMapBuilder()
		_ = _coll23
		for _iter23 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, false)).Cardinality(), false)).Keys().Elements()); ; {
			_compr_23, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _136_v5 _dafny.Int
			_136_v5 = interface{}(_compr_23).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, false)).Cardinality(), false)).Contains(_136_v5) {
				_coll23.Add(Companion_Default___.SafeDivisionInt(_136_v5, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(126), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fjfiraa")).Cardinality()))).Cardinality())).Cardinality())).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("to")).Cardinality())))
			}
		}
		return _coll23.ToMap()
	}()))))))))
}
func (_static *CompanionStruct_Default___) Fm50(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.SeqOf((func() _dafny.Map {
		var _coll24 = _dafny.NewMapBuilder()
		_ = _coll24
		for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(614), _dafny.IntOfInt64(484))); ; {
			_compr_24, _ok24 := _iter24()
			if !_ok24 {
				break
			}
			var _137_v0 _dafny.Int
			_137_v0 = interface{}(_compr_24).(_dafny.Int)
			if ((_dafny.IntOfInt64(614)).Cmp(_137_v0) <= 0) && ((_137_v0).Cmp(_dafny.IntOfInt64(484)) < 0) {
				_coll24.Add((_137_v0).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-399))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_138_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				}))).Cardinality())), true)
			}
		}
		return _coll24.ToMap()
	}()).Cardinality()))).Intersection(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm51(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D13 {
	return Companion_D13_.Create_DC43_(Companion_D13_.Create_DC43_(Companion_D13_.Create_DC42_(_dafny.CodePoint('k'), true)))
}
func (_static *CompanionStruct_Default___) Fm52(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), _dafny.SeqOf((_dafny.SetOf((_dafny.MultiSetOf(true)).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-309))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_139_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jlkqqbehd")).Cardinality())
	}))).Cardinality()), _dafny.IntOfInt64(589)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(578), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-859))))))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-541), _dafny.SeqOf(_dafny.IntOfInt64(83))))
}
func (_static *CompanionStruct_Default___) Fm53(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-116), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf((_dafny.SetOf(!(false))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(50)))).Cardinality()), _dafny.IntOfInt64(945))).Cardinality())), _dafny.SetOf((func() _dafny.Map {
		var _coll25 = _dafny.NewMapBuilder()
		_ = _coll25
		for _iter25 := _dafny.Iterate(((Companion_D22_.Create_DC67_(_dafny.UnicodeSeqOfUtf8Bytes("hac"))).Dtor_cf108()).Elements()); ; {
			_compr_25, _ok25 := _iter25()
			if !_ok25 {
				break
			}
			var _140_v0 _dafny.CodePoint
			_140_v0 = interface{}(_compr_25).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains((Companion_D22_.Create_DC67_(_dafny.UnicodeSeqOfUtf8Bytes("hac"))).Dtor_cf108(), _140_v0) {
				_coll25.Add(_140_v0, (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(426))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg34 _dafny.Int) interface{} {
						return coer34(arg34)
					}
				}(func(_141_i0 _dafny.Int) _dafny.Int {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(911))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg35 _dafny.Int) interface{} {
							return coer35(arg35)
						}
					}(func(_142_i1 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('x')
					}))).Cardinality()))
				})))).Cardinality())
			}
		}
		return _coll25.ToMap()
	}()).Cardinality()), _dafny.SetOf(_dafny.IntOfInt64(106), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(427))).Cardinality())).Cardinality())))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf((_dafny.MultiSetOf(_dafny.CodePoint('q'))).Cardinality(), _dafny.IntOfInt64(326))), _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(887)))))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	if !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pgi")).Cardinality())).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(371))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg36 _dafny.Int) interface{} {
			return coer36(arg36)
		}
	}(func(_143_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))).Cardinality()))) >= 0) {
		return _dafny.SetOf((Companion_D8_.Create_DC24_(_dafny.SetOf((Companion_D1_.Create_DC5_(true, false)).Dtor_cf7()))).Dtor_cf45())
	} else {
		return (func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter26 := _dafny.Iterate((_dafny.SetOf(_dafny.SetOf(true, false))).Elements()); ; {
				_compr_26, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _144_v0 _dafny.Set
				_144_v0 = interface{}(_compr_26).(_dafny.Set)
				if (_dafny.SetOf(_dafny.SetOf(true, false))).Contains(_144_v0) {
					_coll26.Add(_144_v0)
				}
			}
			return _coll26.ToSet()
		}()).Intersection(_dafny.SetOf(_dafny.SetOf(true, false)))
	}
}
func (_static *CompanionStruct_Default___) Fm55(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(352))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg37 _dafny.Int) interface{} {
			return coer37(arg37)
		}
	}(func(_145_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wvjfiqb")).Cardinality()))
	})), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-532), _dafny.IntOfInt64(-470)))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(328))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg38 _dafny.Int) interface{} {
			return coer38(arg38)
		}
	}(func(_146_i1 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(false)).Cardinality()
	}))), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(615))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}(func(_147_i2 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-994))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg40 _dafny.Int) interface{} {
			return coer40(arg40)
		}
	}(func(_148_i3 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xoxncefj")).Cardinality())
	})), _dafny.SeqOf(_dafny.IntOfInt64(809)))))
}
func (_static *CompanionStruct_Default___) M11(p0 D13, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Map) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 _dafny.Map = _dafny.EmptyMap
	_ = r1
	var _149_v0 _dafny.Sequence
	_ = _149_v0
	_149_v0 = _dafny.UnicodeSeqOfUtf8Bytes("plwnbed")
	var _150_v1 _dafny.Array
	_ = _150_v1
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
	_ = _nw0
	_150_v1 = _nw0
	var _151_v2 D4
	_ = _151_v2
	_151_v2 = Companion_D4_.Create_DC13_(_149_v0, p1, false, _150_v1)
	var _source7 D4 = func(_pat_let0_0 D4) D4 {
		return func(_152_dt__update__tmp_h0 D4) D4 {
			return func(_pat_let1_0 _dafny.Sequence) D4 {
				return func(_153_dt__update_hcf22_h0 _dafny.Sequence) D4 {
					return Companion_D4_.Create_DC13_(_153_dt__update_hcf22_h0, (_152_dt__update__tmp_h0).Dtor_cf23(), (_152_dt__update__tmp_h0).Dtor_cf24(), (_152_dt__update__tmp_h0).Dtor_cf25())
				}(_pat_let1_0)
			}(_dafny.UnicodeSeqOfUtf8Bytes("lkgqgh"))
		}(_pat_let0_0)
	}((func() D4 {
		if false {
			return _151_v2
		}
		return _151_v2
	})())
	_ = _source7
	if _source7.Is_DC13() {
		var _154___mcc_h0 _dafny.Sequence = _source7.Get_().(D4_DC13).Cf22
		_ = _154___mcc_h0
		var _155___mcc_h1 _dafny.Int = _source7.Get_().(D4_DC13).Cf23
		_ = _155___mcc_h1
		var _156___mcc_h2 bool = _source7.Get_().(D4_DC13).Cf24
		_ = _156___mcc_h2
		var _157___mcc_h3 _dafny.Array = _source7.Get_().(D4_DC13).Cf25
		_ = _157___mcc_h3
		var _158_cf25 _dafny.Array = _157___mcc_h3
		_ = _158_cf25
		var _159_cf24 bool = _156___mcc_h2
		_ = _159_cf24
		var _160_cf23 _dafny.Int = _155___mcc_h1
		_ = _160_cf23
		var _161_cf22 _dafny.Sequence = _154___mcc_h0
		_ = _161_cf22
		_159_cf24 = (func() bool {
			if true {
				return p3
			}
			return p2
		})()
		(globalState).F9 = p2
		var _162_v3 _dafny.MultiSet
		_ = _162_v3
		_162_v3 = _dafny.MultiSetOf(p3)
		var _163_v4 T0
		_ = _163_v4
		var _nw1 *C6 = New_C6_()
		_ = _nw1
		_nw1.Ctor__(_159_cf24, (_162_v3).Union((_162_v3).Update(!(p3), Companion_Default___.Abs(_160_cf23))))
		_163_v4 = _nw1
		_163_v4 = _163_v4
		var _164_v5 _dafny.Array
		_ = _164_v5
		var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
		_ = _nw2
		_164_v5 = _nw2
		var _165_v6 _dafny.Map
		_ = _165_v6
		_165_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-374), _164_v5)
		var _166_v7 D18
		_ = _166_v7
		_166_v7 = Companion_D18_.Create_DC56_(p3, _163_v4.F14(), _160_cf23)
		var _167_v8 _dafny.Sequence
		_ = _167_v8
		_167_v8 = _dafny.SeqOf(p2)
		var _168_v9 *C0
		_ = _168_v9
		var _nw3 *C0 = New_C0_()
		_ = _nw3
		_nw3.Ctor__((_167_v8).Select((Companion_Default___.SafeIndex(_160_cf23, _dafny.IntOfUint32((_167_v8).Cardinality()))).Uint32()).(bool), p2)
		_168_v9 = _nw3
		var _169_v10 _dafny.Sequence
		_ = _169_v10
		_169_v10 = _dafny.SeqOf(_168_v9)
		var _170_v11 _dafny.Map
		_ = _170_v11
		_170_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_166_v7).Dtor_cf93(), _dafny.IntOfUint32((_169_v10).Cardinality()))
		var _171_v12 _dafny.Sequence
		_ = _171_v12
		_171_v12 = _dafny.SeqOf(_170_v11, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _160_cf23), (_170_v11).Update((_168_v9).F24(), p1))
		var _172_v13 *C8
		_ = _172_v13
		var _nw4 *C8 = New_C8_()
		_ = _nw4
		_nw4.Ctor__(_165_v6, _171_v12, p2, _162_v3)
		_172_v13 = _nw4
		var _173_v14 _dafny.Array
		_ = _173_v14
		var _nwElement0_0 *C8 = _172_v13
		_ = _nwElement0_0
		var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(5))
		_ = _nw5
		(_nw5).ArraySet1(_nwElement0_0, 0)
		(_nw5).ArraySet1(_172_v13, 1)
		(_nw5).ArraySet1(_172_v13, 2)
		(_nw5).ArraySet1(_172_v13, 3)
		(_nw5).ArraySet1(_172_v13, 4)
		_173_v14 = _nw5
		_173_v14 = (func() _dafny.Array {
			if _159_cf24 {
				return _173_v14
			}
			return _173_v14
		})()
	} else if _source7.Is_DC14() {
		var _174___mcc_h4 bool = _source7.Get_().(D4_DC14).Cf26
		_ = _174___mcc_h4
		var _175___mcc_h5 bool = _source7.Get_().(D4_DC14).Cf27
		_ = _175___mcc_h5
		var _176_cf27 bool = _175___mcc_h5
		_ = _176_cf27
		var _177_cf26 bool = _174___mcc_h4
		_ = _177_cf26
		var _178_v15 _dafny.Sequence
		_ = _178_v15
		_178_v15 = _dafny.SeqOf(p2)
		var _179_v16 *C6
		_ = _179_v16
		var _nw6 *C6 = New_C6_()
		_ = _nw6
		_nw6.Ctor__(true, _dafny.MultiSetFromSeq(_dafny.SeqOf(_177_cf26)))
		_179_v16 = _nw6
		var _180_v17 _dafny.Map
		_ = _180_v17
		_180_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _179_v16)
		var _rhs0 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if p3 {
				return _178_v15
			}
			return _178_v15
		})()).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(690), _177_cf26)).Cardinality()).Minus(((_180_v17).Update((_178_v15).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_178_v15).Cardinality()))).Uint32()).(bool), _179_v16)).Cardinality()))
		_ = _rhs0
		var _rhs1 bool = (_179_v16).Fm2(_dafny.Companion_Sequence_.IsPrefixOf(_178_v15, _178_v15), p2, globalState)
		_ = _rhs1
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		_lhs0.F3 = _rhs0
		_lhs1.F9 = _rhs1
		var _181_v18 _dafny.MultiSet
		_ = _181_v18
		_181_v18 = _dafny.MultiSetOf(_177_cf26, _176_cf27)
		var _182_v19 *C4
		_ = _182_v19
		var _nw7 *C4 = New_C4_()
		_ = _nw7
		_nw7.Ctor__(_176_cf27, _177_cf26, _181_v18)
		_182_v19 = _nw7
		var _183_v20 _dafny.Map
		_ = _183_v20
		_183_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_v19, p2)
		var _184_v21 *C5
		_ = _184_v21
		var _nw8 *C5 = New_C5_()
		_ = _nw8
		_nw8.Ctor__((p1).Plus((_183_v20).Cardinality()), p2, _dafny.MultiSetOf(_177_cf26, true))
		_184_v21 = _nw8
		_184_v21 = _184_v21
		var _185_v22 _dafny.Array
		_ = _185_v22
		var _nw9 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw9
		_185_v22 = _nw9
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_185_v22), 0))
		_ = _index0
		(_185_v22).ArraySet1(((_184_v21).F18()).Times((_dafny.Zero).Minus(p1)), (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_150_v1), 0))
		_ = _index1
		(_150_v1).ArraySet1((p1).Cmp(p1) != 0, (_index1).Int())
		var _186_v23 _dafny.Array
		_ = _186_v23
		var _nw10 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
		_ = _nw10
		_186_v23 = _nw10
		var _187_v24 _dafny.Set
		_ = _187_v24
		_187_v24 = _dafny.SetOf(_186_v23, _186_v23, _186_v23)
		var _188_v25 _dafny.Map
		_ = _188_v25
		_188_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_179_v16).Fm2(p2, (_182_v19).F19(), globalState), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p1), (_187_v24).Cardinality()))
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_185_v22), 0))
		_ = _index2
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_150_v1), 0))
		_ = _index3
		var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_149_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(918))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg41 _dafny.Int) interface{} {
				return coer41(arg41)
			}
		}(func(_189_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		})))
		_ = _rhs2
		var _rhs3 bool = _177_cf26
		_ = _rhs3
		var _rhs4 _dafny.Int = (_188_v25).Cardinality()
		_ = _rhs4
		var _rhs5 bool = p2
		_ = _rhs5
		var _lhs2 *GlobalState = globalState
		_ = _lhs2
		var _lhs3 _dafny.Array = _185_v22
		_ = _lhs3
		var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_185_v22), 0))
		_ = _lhs4
		var _lhs5 _dafny.Array = _150_v1
		_ = _lhs5
		var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_150_v1), 0))
		_ = _lhs6
		_149_v0 = _rhs2
		_lhs2.F9 = _rhs3
		(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
		(_lhs5).ArraySet1(_rhs5, (_lhs6).Int())
		var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(564), _dafny.ArrayLen((_150_v1), 0))
		_ = _index4
		(_150_v1).ArraySet1((_178_v15).Select((Companion_Default___.SafeIndex((_185_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_185_v22), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_178_v15).Cardinality()))).Uint32()).(bool), (_index4).Int())
	} else if _source7.Is_DC12() {
		var _190___mcc_h6 _dafny.Array = _source7.Get_().(D4_DC12).Cf21
		_ = _190___mcc_h6
		var _191_cf21 _dafny.Array = _190___mcc_h6
		_ = _191_cf21
		var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
		_ = _index5
		(_150_v1).ArraySet1((p1).Cmp(p1) < 0, (_index5).Int())
		var _192_v26 _dafny.Set
		_ = _192_v26
		_192_v26 = _dafny.SetOf(p3, p2, false)
		var _193_v27 _dafny.Set
		_ = _193_v27
		_193_v27 = _dafny.SetOf(_192_v26)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
		_ = _index6
		var _rhs6 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, p1))
		_ = _rhs6
		var _rhs7 bool = ((_193_v27).Union(_193_v27)).IsSubsetOf(Companion_Default___.Fm54(_dafny.IntOfInt64(850), _149_v0, globalState))
		_ = _rhs7
		var _lhs7 *GlobalState = globalState
		_ = _lhs7
		var _lhs8 _dafny.Array = _150_v1
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
		_ = _lhs9
		_lhs7.F3 = _rhs6
		(_lhs8).ArraySet1(_rhs7, (_lhs9).Int())
		if (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool) {
			(globalState).F3 = p1
			var _194_v28 _dafny.Sequence
			_ = _194_v28
			_194_v28 = _dafny.SeqOf(p1)
			var _195_v29 _dafny.MultiSet
			_ = _195_v29
			_195_v29 = _dafny.MultiSetOf(p3)
			var _196_v30 *C6
			_ = _196_v30
			var _nw11 *C6 = New_C6_()
			_ = _nw11
			_nw11.Ctor__((_dafny.IntOfUint32((_194_v28).Cardinality())).Cmp(p1) > 0, _195_v29)
			_196_v30 = _nw11
			(globalState).F3 = (p1).Plus(p1)
			var _197_v31 T0
			_ = _197_v31
			var _nw12 *C3 = New_C3_()
			_ = _nw12
			_nw12.Ctor__(p1, _149_v0, p3, _195_v29)
			_197_v31 = _nw12
			var _198_v32 _dafny.Sequence
			_ = _198_v32
			_198_v32 = _dafny.SeqOf(_197_v31)
			var _199_v33 _dafny.Array
			_ = _199_v33
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw13
			_199_v33 = _nw13
			var _200_v34 _dafny.Set
			_ = _200_v34
			_200_v34 = _dafny.SetOf(_199_v33, _199_v33, _199_v33)
			var _201_v35 _dafny.Set
			_ = _201_v35
			_201_v35 = _dafny.SetOf((_200_v34).Cardinality())
			(globalState).F9 = !(_201_v35).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_198_v32).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_198_v32).Cardinality()))).Uint32()).(T0), _197_v31, _197_v31, _197_v31, _197_v31), _198_v32)).Cardinality()))
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
			_ = _index7
			(_150_v1).ArraySet1(!(true), (_index7).Int())
		} else {
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
			_ = _index8
			(_150_v1).ArraySet1((p1).Cmp(p1) >= 0, (_index8).Int())
			var _202_v36 _dafny.Map
			_ = _202_v36
			_202_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool))
			var _203_v37 D1
			_ = _203_v37
			_203_v37 = Companion_D1_.Create_DC5_(p3, true)
			var _204_v38 _dafny.Map
			_ = _204_v38
			_204_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_202_v36).Contains(p1) {
					return (_202_v36).Get(p1).(bool)
				}
				return true
			})(), _203_v37)
			_204_v38 = _204_v38
			var _205_v39 _dafny.Set
			_ = _205_v39
			_205_v39 = _dafny.SetOf(p1, p1)
			var _206_v40 _dafny.Array
			_ = _206_v40
			var _nw14 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(28))
			_ = _nw14
			_206_v40 = _nw14
			var _207_v41 _dafny.Sequence
			_ = _207_v41
			_207_v41 = _dafny.SeqOf(_205_v39, Companion_Default___.Fm6(globalState))
			var _208_v42 _dafny.Sequence
			_ = _208_v42
			_208_v42 = _dafny.SeqOf(p2, p3)
			var _209_v43 _dafny.Map
			_ = _209_v43
			_209_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool))
			var _210_v44 _dafny.CodePoint
			_ = _210_v44
			_210_v44 = _dafny.CodePoint('c')
			var _211_v45 _dafny.Map
			_ = _211_v45
			_211_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_210_v44, p1)
			var _212_v46 _dafny.Sequence
			_ = _212_v46
			_212_v46 = _dafny.SeqOf(p1, p1, p1, p1, p1)
			var _213_v47 _dafny.Set
			_ = _213_v47
			_213_v47 = _dafny.SetOf(_212_v46, _212_v46, _212_v46, _212_v46)
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
			_ = _index9
			var _rhs8 _dafny.Set = ((_207_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_208_v42).Cardinality()), _dafny.IntOfUint32((_207_v41).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_205_v39)
			_ = _rhs8
			var _rhs9 bool = (func() bool {
				if (_209_v43).Contains(!((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool))) {
					return (_209_v43).Get(!((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool))).(bool)
				}
				return Companion_Default___.Fm23((_202_v36).Cardinality(), (func() bool {
					if (_202_v36).Contains(p1) {
						return (_202_v36).Get(p1).(bool)
					}
					return !(p3)
				})(), _211_v45, _213_v47, globalState)
			})()
			_ = _rhs9
			var _rhs10 _dafny.Array = _206_v40
			_ = _rhs10
			var _lhs10 _dafny.Array = _150_v1
			_ = _lhs10
			var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
			_ = _lhs11
			_205_v39 = _rhs8
			(_lhs10).ArraySet1(_rhs9, (_lhs11).Int())
			_206_v40 = _rhs10
			var _214_v48 _dafny.Array
			_ = _214_v48
			var _nwElement0_1 _dafny.Array = _191_cf21
			_ = _nwElement0_1
			var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
			_ = _nw15
			(_nw15).ArraySet1(_nwElement0_1, 0)
			(_nw15).ArraySet1(_191_cf21, 1)
			(_nw15).ArraySet1(_191_cf21, 2)
			(_nw15).ArraySet1(_191_cf21, 3)
			(_nw15).ArraySet1(_191_cf21, 4)
			(_nw15).ArraySet1(_191_cf21, 5)
			(_nw15).ArraySet1(_191_cf21, 6)
			(_nw15).ArraySet1(_191_cf21, 7)
			_214_v48 = _nw15
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_214_v48), 0))
			_ = _index10
			(_214_v48).ArraySet1(_191_cf21, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(246), _dafny.ArrayLen((_214_v48), 0))
			_ = _index11
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
			_ = _nw16
			(_214_v48).ArraySet1(_nw16, (_index11).Int())
			var _215_v49 _dafny.MultiSet
			_ = _215_v49
			_215_v49 = _dafny.MultiSetOf(p3, p2)
			var _216_v50 *C5
			_ = _216_v50
			var _nw17 *C5 = New_C5_()
			_ = _nw17
			_nw17.Ctor__(p1, p3, (_215_v49).Difference((_215_v49).Update(p3, Companion_Default___.Abs((_215_v49).Cardinality()))))
			_216_v50 = _nw17
		}
		var _217_v51 _dafny.Map
		_ = _217_v51
		_217_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
		var _218_v52 _dafny.MultiSet
		_ = _218_v52
		_218_v52 = _dafny.MultiSetOf((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), p2)
		var _219_v53 _dafny.MultiSet
		_ = _219_v53
		_219_v53 = _dafny.MultiSetOf(_dafny.IntOfInt64(-561), p1)
		var _220_v54 _dafny.Sequence
		_ = _220_v54
		_220_v54 = _dafny.SeqOf(p1)
		var _221_v55 _dafny.Sequence
		_ = _221_v55
		_221_v55 = _dafny.SeqOf(_220_v54)
		var _222_v56 _dafny.Map
		_ = _222_v56
		_222_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(622))
		var _223_v57 _dafny.Array
		_ = _223_v57
		var _nwElement0_2 _dafny.Int = Companion_Default___.SafeModuloInt(p1, p1)
		_ = _nwElement0_2
		var _nw18 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(23))
		_ = _nw18
		(_nw18).ArraySet1(_nwElement0_2, 0)
		(_nw18).ArraySet1(Companion_Default___.Fm0(p2, (_dafny.Zero).Minus(p1), p1, (_dafny.MultiSetOf(_dafny.IntOfUint32((_149_v0).Cardinality()), _dafny.IntOfInt64(-622), p1, p1, (_dafny.Zero).Minus(p1))).Update(p1, Companion_Default___.Abs((_dafny.Zero).Minus((_217_v51).Cardinality()))), globalState), 1)
		(_nw18).ArraySet1(p1, 2)
		(_nw18).ArraySet1((p1).Plus((_dafny.Zero).Minus((_218_v52).Cardinality())), 3)
		(_nw18).ArraySet1(p1, 4)
		(_nw18).ArraySet1(_dafny.IntOfInt64(978), 5)
		(_nw18).ArraySet1(Companion_Default___.Fm0(false, p1, p1, _219_v53, globalState), 6)
		(_nw18).ArraySet1(_dafny.IntOfUint32((_221_v55).Cardinality()), 7)
		(_nw18).ArraySet1(p1, 8)
		(_nw18).ArraySet1((func() _dafny.Int {
			if p2 {
				return p1
			}
			return p1
		})(), 9)
		(_nw18).ArraySet1((func() _dafny.Int {
			if (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool) {
				return (_220_v54).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_220_v54).Cardinality()))).Uint32()).(_dafny.Int)
			}
			return p1
		})(), 10)
		(_nw18).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p1), 11)
		(_nw18).ArraySet1(p1, 12)
		(_nw18).ArraySet1(_dafny.IntOfInt64(929), 13)
		(_nw18).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_220_v54).Cardinality())), 14)
		(_nw18).ArraySet1(p1, 15)
		(_nw18).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_220_v54).Cardinality())), 16)
		(_nw18).ArraySet1(p1, 17)
		(_nw18).ArraySet1((func() _dafny.Int {
			if (_222_v56).Contains((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool)) {
				return (_222_v56).Get((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool)).(_dafny.Int)
			}
			return _dafny.IntOfInt64(171)
		})(), 18)
		(_nw18).ArraySet1(p1, 19)
		(_nw18).ArraySet1(p1, 20)
		(_nw18).ArraySet1(p1, 21)
		(_nw18).ArraySet1(p1, 22)
		_223_v57 = _nw18
		_223_v57 = _223_v57
		if ((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, p1))).Cmp(p1) <= 0 {
			(globalState).F9 = !((func() bool {
				if true {
					return (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool)
				}
				return p2
			})())
			var _224_v58 _dafny.Map
			_ = _224_v58
			_224_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_cf21, (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool))
			var _225_v59 _dafny.CodePoint
			_ = _225_v59
			_225_v59 = _dafny.CodePoint('p')
			var _226_v60 _dafny.Map
			_ = _226_v60
			_226_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_225_v59, _dafny.IntOfUint32((_dafny.SeqOf(p1, (_dafny.SetOf(p1)).Cardinality())).Cardinality()))
			var _227_v61 _dafny.Set
			_ = _227_v61
			_227_v61 = _dafny.SetOf(_220_v54)
			_224_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_191_cf21, Companion_Default___.Fm23(p1, false, Companion_Default___.Fm24(_dafny.IntOfInt64(-773), p3, true, p1, globalState), func() _dafny.Set {
				var _coll27 = _dafny.NewBuilder()
				_ = _coll27
				for _iter27 := _dafny.Iterate((Companion_Default___.Fm55(Companion_Default___.Fm23(p1, p2, _226_v60, _227_v61, globalState), globalState)).Elements()); ; {
					_compr_27, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _228_v62 _dafny.Sequence
					_228_v62 = interface{}(_compr_27).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(Companion_Default___.Fm55(Companion_Default___.Fm23(p1, p2, _226_v60, _227_v61, globalState), globalState), _228_v62) {
						_coll27.Add(_228_v62)
					}
				}
				return _coll27.ToSet()
			}(), globalState))
			var _229_v64 _dafny.Map
			_ = _229_v64
			_229_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v54, p2)
			var _230_v66 _dafny.Array
			_ = _230_v66
			var _nw19 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
			_ = _nw19
			_230_v66 = _nw19
			var _231_v67 *C3
			_ = _231_v67
			var _nw20 *C3 = New_C3_()
			_ = _nw20
			_nw20.Ctor__(Companion_Default___.Fm0((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), p1, p1, _219_v53, globalState), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_149_v0, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(Companion_Default___.Fm23(p1, false, func() _dafny.Map {
				var _coll28 = _dafny.NewMapBuilder()
				_ = _coll28
				for _iter28 := _dafny.Iterate((_149_v0).Elements()); ; {
					_compr_28, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _232_v63 _dafny.CodePoint
					_232_v63 = interface{}(_compr_28).(_dafny.CodePoint)
					if _dafny.Companion_Sequence_.Contains(_149_v0, _232_v63) {
						_coll28.Add(_232_v63, p1)
					}
				}
				return _coll28.ToMap()
			}(), func() _dafny.Set {
				var _coll29 = _dafny.NewBuilder()
				_ = _coll29
				for _iter29 := _dafny.Iterate((_229_v64).Keys().Elements()); ; {
					_compr_29, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _233_v65 _dafny.Sequence
					_233_v65 = interface{}(_compr_29).(_dafny.Sequence)
					if (_229_v64).Contains(_233_v65) {
						_coll29.Add(_233_v65)
					}
				}
				return _coll29.ToSet()
			}(), globalState), p1, p1, _219_v53, globalState), _230_v66)).Cardinality(), _dafny.IntOfUint32((_149_v0).Cardinality()))).Uint32(), _225_v59), _dafny.Companion_Sequence_.Update(_149_v0, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_149_v0).Cardinality()))).Uint32(), _225_v59)), (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), _218_v52)
			_231_v67 = _nw20
			var _nw21 *C3 = New_C3_()
			_ = _nw21
			_nw21.Ctor__((_231_v67).F20(), (_231_v67).F21(), p3, _218_v52)
			_231_v67 = _nw21
			var _234_v68 _dafny.Set
			_ = _234_v68
			_234_v68 = _dafny.SetOf((_231_v67).F20(), (_231_v67).F20())
			var _235_v69 _dafny.Int
			_ = _235_v69
			var _236_v70 _dafny.Int
			_ = _236_v70
			var _237_v71 bool
			_ = _237_v71
			var _out0 _dafny.Int
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = (_231_v67).M1(_234_v68, (_231_v67).F20(), globalState)
			_235_v69 = _out0
			_236_v70 = _out1
			_237_v71 = _out2
			var _238_v72 D5
			_ = _238_v72
			_238_v72 = Companion_D5_.Create_DC18_(false)
			var _239_v73 _dafny.Map
			_ = _239_v73
			_239_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _238_v72)
			_239_v73 = (_239_v73).Update(p3, Companion_D5_.Create_DC18_((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool)))
		} else {
			var _240_v75 _dafny.CodePoint
			_ = _240_v75
			_240_v75 = _dafny.CodePoint('f')
			var _241_v77 _dafny.Sequence
			_ = _241_v77
			_241_v77 = _dafny.SeqOf(p2)
			var _242_v78 _dafny.Map
			_ = _242_v78
			_242_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_149_v0).Cardinality()), (func() _dafny.Map {
				var _coll30 = _dafny.NewMapBuilder()
				_ = _coll30
				for _iter30 := _dafny.Iterate((_dafny.MultiSetOf(_241_v77)).Elements()); ; {
					_compr_30, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _243_v76 _dafny.Sequence
					_243_v76 = interface{}(_compr_30).(_dafny.Sequence)
					if (_dafny.MultiSetOf(_241_v77)).Contains(_243_v76) {
						_coll30.Add(_243_v76, _dafny.SeqOf(p1))
					}
				}
				return _coll30.ToMap()
			}()).Cardinality())
			var _244_v79 _dafny.Map
			_ = _244_v79
			_244_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_240_v75, (_242_v78).Update(_dafny.IntOfUint32((_241_v77).Cardinality()), _dafny.IntOfInt64(165)))
			var _245_v80 _dafny.Set
			_ = _245_v80
			_245_v80 = _dafny.SetOf(p1, p1)
			var _246_v81 D9
			_ = _246_v81
			_246_v81 = Companion_D9_.Create_DC30_(Companion_Default___.Fm25(p1, p1, p3, Companion_Default___.Fm1(p2, p1, true, _245_v80, globalState), globalState), (_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool), p1, _245_v80, _dafny.IntOfInt64(961))
			var _247_v82 _dafny.Map
			_ = _247_v82
			_247_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_246_v81).Dtor_cf50(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_248_v75 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_249_i1 _dafny.Int) _dafny.CodePoint {
					return _248_v75
				}
			})(_240_v75)))).Cardinality())))
			var _250_v84 _dafny.Map
			_ = _250_v84
			_250_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(101)), p1)
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_150_v1), 0))
			_ = _index12
			(_150_v1).ArraySet1(Companion_Default___.Fm23(p1, false, (func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter31 := _dafny.Iterate((_244_v79).Keys().Elements()); ; {
					_compr_31, _ok31 := _iter31()
					if !_ok31 {
						break
					}
					var _251_v74 _dafny.CodePoint
					_251_v74 = interface{}(_compr_31).(_dafny.CodePoint)
					if (_244_v79).Contains(_251_v74) {
						_coll31.Add(_251_v74, _dafny.IntOfInt64(781))
					}
				}
				return _coll31.ToMap()
			}()).Merge(_247_v82), func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter32 := _dafny.Iterate((func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter33 := _dafny.Iterate((_250_v84).Keys().Elements()); ; {
						_compr_33, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _252_v83 _dafny.Sequence
						_252_v83 = interface{}(_compr_33).(_dafny.Sequence)
						if (_250_v84).Contains(_252_v83) {
							_coll33.Add(_252_v83, _dafny.IntOfInt64(540))
						}
					}
					return _coll33.ToMap()
				}()).Keys().Elements()); ; {
					_compr_32, _ok32 := _iter32()
					if !_ok32 {
						break
					}
					var _253_v85 _dafny.Sequence
					_253_v85 = interface{}(_compr_32).(_dafny.Sequence)
					if (func() _dafny.Map {
						var _coll34 = _dafny.NewMapBuilder()
						_ = _coll34
						for _iter34 := _dafny.Iterate((_250_v84).Keys().Elements()); ; {
							_compr_34, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _252_v83 _dafny.Sequence
							_252_v83 = interface{}(_compr_34).(_dafny.Sequence)
							if (_250_v84).Contains(_252_v83) {
								_coll34.Add(_252_v83, _dafny.IntOfInt64(540))
							}
						}
						return _coll34.ToMap()
					}()).Contains(_253_v85) {
						_coll32.Add(_253_v85)
					}
				}
				return _coll32.ToSet()
			}(), globalState), (_index12).Int())
			var _254_v86 _dafny.Array
			_ = _254_v86
			var _nwElement0_3 _dafny.Sequence = _149_v0
			_ = _nwElement0_3
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.One)
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_3, 0)
			_254_v86 = _nw22
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))
			_ = _index13
			(_254_v86).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_149_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(869))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}((func(_255_v75 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_256_i2 _dafny.Int) _dafny.CodePoint {
					return _255_v75
				}
			})(_240_v75)))), (_index13).Int())
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))
			_ = _index14
			(_254_v86).ArraySet1(_149_v0, (_index14).Int())
			var _rhs11 _dafny.Int = _dafny.IntOfInt64(667)
			_ = _rhs11
			var _rhs12 _dafny.Array = _191_cf21
			_ = _rhs12
			var _rhs13 _dafny.Int = (p1).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_254_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.Update((_254_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))).Int()).(_dafny.Sequence), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_254_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32(), ((_254_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_222_v56).Cardinality(), _dafny.IntOfUint32(((_254_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(664), _dafny.ArrayLen((_254_v86), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.CodePoint)))).Cardinality()))
			_ = _rhs13
			var _lhs12 *GlobalState = globalState
			_ = _lhs12
			_lhs12.F3 = _rhs11
			_191_cf21 = _rhs12
			r0 = _rhs13
			_245_v80 = ((_245_v80).Intersection(_245_v80)).Union((_dafny.SetOf(p1)).Difference(_245_v80))
			var _257_v87 D1
			_ = _257_v87
			_257_v87 = Companion_D1_.Create_DC3_(p1)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_223_v57), 0))
			_ = _index15
			(_223_v57).ArraySet1((_257_v87).Dtor_cf3(), (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_223_v57), 0))
			_ = _index16
			(_223_v57).ArraySet1((Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfInt64(474))).Minus(p1), (_index16).Int())
		}
	} else {
		var _258___mcc_h7 D4 = _source7.Get_().(D4_DC15).Cf28
		_ = _258___mcc_h7
		var _259_cf28 D4 = _258___mcc_h7
		_ = _259_cf28
		var _260_v88 _dafny.CodePoint
		_ = _260_v88
		_260_v88 = _dafny.CodePoint('i')
		var _261_v89 _dafny.Map
		_ = _261_v89
		_261_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_260_v88, p1)
		var _262_v90 _dafny.MultiSet
		_ = _262_v90
		_262_v90 = _dafny.MultiSetOf(p3, p3)
		var _263_v91 _dafny.Sequence
		_ = _263_v91
		_263_v91 = _dafny.SeqOf((_262_v90).Cardinality(), _dafny.IntOfInt64(555))
		var _264_v92 _dafny.Set
		_ = _264_v92
		_264_v92 = _dafny.SetOf(_263_v91, _dafny.SeqOf(p1))
		var _265_v93 _dafny.Sequence
		_ = _265_v93
		_265_v93 = _dafny.SeqOf(Companion_Default___.Fm23((_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-800)), p1)).Cardinality(), p2, _261_v89, _264_v92, globalState))
		var _266_v94 _dafny.Array
		_ = _266_v94
		var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.One)
		_ = _nw23
		_266_v94 = _nw23
		var _267_v95 _dafny.Map
		_ = _267_v95
		_267_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_265_v93)).Cardinality(), _266_v94)
		var _268_v96 _dafny.Map
		_ = _268_v96
		_268_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_dafny.Zero).Minus(_dafny.IntOfUint32((_149_v0).Cardinality())))
		var _269_v97 _dafny.Sequence
		_ = _269_v97
		_269_v97 = _dafny.SeqOf(_268_v96)
		var _270_v98 T0
		_ = _270_v98
		var _nw24 *C8 = New_C8_()
		_ = _nw24
		_nw24.Ctor__(_267_v95, _269_v97, (_265_v93).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_265_v93).Cardinality()), _dafny.IntOfUint32((_265_v93).Cardinality()))).Uint32()).(bool), (_262_v90).Update(p3, Companion_Default___.Abs((_dafny.Zero).Minus(p1))))
		_270_v98 = _nw24
		var _271_v99 D2
		_ = _271_v99
		_271_v99 = Companion_D2_.Create_DC9_(_270_v98.F14(), _dafny.IntOfInt64(19), (_149_v0).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_149_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))
		var _nw25 *C3 = New_C3_()
		_ = _nw25
		_nw25.Ctor__((_271_v99).Dtor_cf14(), Companion_Default___.Fm27(p1, _270_v98.F14(), globalState), _270_v98.F14(), ((_270_v98).F15()).Intersection(_dafny.MultiSetFromSeq(_265_v93)))
		_270_v98 = _nw25
		(_270_v98).F14_set_(false)
		var _272_v100 _dafny.Array
		_ = _272_v100
		var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
		_ = _nw26
		_272_v100 = _nw26
		var _273_v101 _dafny.Map
		_ = _273_v101
		_273_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_270_v98.F14(), p3)
		var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_272_v100), 0))
		_ = _index17
		(_272_v100).ArraySet1(_273_v101, (_index17).Int())
		var _274_v102 _dafny.MultiSet
		_ = _274_v102
		_274_v102 = _dafny.MultiSetOf(_dafny.IntOfUint32((_149_v0).Cardinality()))
		var _275_v103 _dafny.Array
		_ = _275_v103
		var _nwElement0_4 _dafny.Int = _dafny.IntOfInt64(-46)
		_ = _nwElement0_4
		var _nw27 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(3))
		_ = _nw27
		(_nw27).ArraySet1(_nwElement0_4, 0)
		(_nw27).ArraySet1(Companion_Default___.Fm0(_270_v98.F14(), p1, p1, _274_v102, globalState), 1)
		(_nw27).ArraySet1((_dafny.Zero).Minus(p1), 2)
		_275_v103 = _nw27
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
		_ = _index18
		(_275_v103).ArraySet1(p1, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_150_v1), 0))
		_ = _index19
		(_150_v1).ArraySet1(_270_v98.F14(), (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_266_v94), 0))
		_ = _index20
		(_266_v94).ArraySet1CodePoint(_260_v88, (_index20).Int())
		var _276_v104 _dafny.Sequence
		_ = _276_v104
		_276_v104 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _270_v98.F14()), _273_v101, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2))
		var _277_v105 *C8
		_ = _277_v105
		var _nw28 *C8 = New_C8_()
		_ = _nw28
		_nw28.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _266_v94), _269_v97, _270_v98.F14(), (_dafny.MultiSetOf(p2)).Update(_270_v98.F14(), Companion_Default___.Abs(p1)))
		_277_v105 = _nw28
		var _278_v106 _dafny.MultiSet
		_ = _278_v106
		_278_v106 = _dafny.MultiSetOf(_277_v105)
		var _279_v107 _dafny.Map
		_ = _279_v107
		_279_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _270_v98.F14())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_272_v100), 0))
		_ = _index21
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
		_ = _index22
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_150_v1), 0))
		_ = _index23
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_266_v94), 0))
		_ = _index24
		var _rhs14 _dafny.Map = (_276_v104).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_276_v104).Cardinality()))).Uint32()).(_dafny.Map)
		_ = _rhs14
		var _rhs15 _dafny.Int = (_dafny.Zero).Minus(((_dafny.Zero).Minus(((_dafny.MultiSetOf((_dafny.Zero).Minus((_278_v106).Cardinality()))).Cardinality()).Times(p1))).Plus(p1))
		_ = _rhs15
		var _rhs16 _dafny.Int = p1
		_ = _rhs16
		var _rhs17 bool = (func() bool {
			if (_279_v107).Contains(_dafny.IntOfUint32((Companion_Default___.Fm27(_dafny.IntOfInt64(467), _270_v98.F14(), globalState)).Cardinality())) {
				return (_279_v107).Get(_dafny.IntOfUint32((Companion_Default___.Fm27(_dafny.IntOfInt64(467), _270_v98.F14(), globalState)).Cardinality())).(bool)
			}
			return p2
		})()
		_ = _rhs17
		var _rhs18 _dafny.CodePoint = _260_v88
		_ = _rhs18
		var _lhs13 _dafny.Array = _272_v100
		_ = _lhs13
		var _lhs14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.ArrayLen((_272_v100), 0))
		_ = _lhs14
		var _lhs15 _dafny.Array = _275_v103
		_ = _lhs15
		var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
		_ = _lhs16
		var _lhs17 _dafny.Array = _150_v1
		_ = _lhs17
		var _lhs18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_150_v1), 0))
		_ = _lhs18
		var _lhs19 _dafny.Array = _266_v94
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_266_v94), 0))
		_ = _lhs20
		(_lhs13).ArraySet1(_rhs14, (_lhs14).Int())
		r0 = _rhs15
		(_lhs15).ArraySet1(_rhs16, (_lhs16).Int())
		(_lhs17).ArraySet1(_rhs17, (_lhs18).Int())
		(_lhs19).ArraySet1CodePoint(_rhs18, (_lhs20).Int())
		if !((_150_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_150_v1), 0))).Int()).(bool)) {
			var _280_v108 *C2
			_ = _280_v108
			var _nw29 *C2 = New_C2_()
			_ = _nw29
			_nw29.Ctor__(_279_v107)
			_280_v108 = _nw29
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
			_ = _index25
			(_275_v103).ArraySet1((_263_v91).Select((Companion_Default___.SafeIndex((_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_263_v91).Cardinality()))).Uint32()).(_dafny.Int), (_index25).Int())
			var _281_v109 _dafny.Array
			_ = _281_v109
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw30
			_281_v109 = _nw30
			var _282_v110 _dafny.Map
			_ = _282_v110
			_282_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int), _275_v103)
			var _283_v111 _dafny.Map
			_ = _283_v111
			_283_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _275_v103)
			var _284_v112 _dafny.Array
			_ = _284_v112
			var _nwElement0_5 _dafny.Array = _275_v103
			_ = _nwElement0_5
			var _nw31 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(25))
			_ = _nw31
			(_nw31).ArraySet1(_nwElement0_5, 0)
			(_nw31).ArraySet1(_275_v103, 1)
			(_nw31).ArraySet1(_275_v103, 2)
			(_nw31).ArraySet1(_275_v103, 3)
			(_nw31).ArraySet1(_275_v103, 4)
			(_nw31).ArraySet1(_275_v103, 5)
			(_nw31).ArraySet1(_275_v103, 6)
			(_nw31).ArraySet1(_275_v103, 7)
			(_nw31).ArraySet1(_275_v103, 8)
			(_nw31).ArraySet1(_275_v103, 9)
			(_nw31).ArraySet1(_275_v103, 10)
			(_nw31).ArraySet1(_275_v103, 11)
			(_nw31).ArraySet1(_275_v103, 12)
			(_nw31).ArraySet1(_275_v103, 13)
			(_nw31).ArraySet1(_275_v103, 14)
			(_nw31).ArraySet1(_275_v103, 15)
			(_nw31).ArraySet1((func() _dafny.Array {
				if (_282_v110).Contains(p1) {
					return (_282_v110).Get(p1).(_dafny.Array)
				}
				return _275_v103
			})(), 16)
			(_nw31).ArraySet1(_275_v103, 17)
			(_nw31).ArraySet1(_275_v103, 18)
			(_nw31).ArraySet1((func() _dafny.Array {
				if (_283_v111).Contains(_270_v98.F14()) {
					return (_283_v111).Get(_270_v98.F14()).(_dafny.Array)
				}
				return _275_v103
			})(), 19)
			(_nw31).ArraySet1(_275_v103, 20)
			(_nw31).ArraySet1(_275_v103, 21)
			(_nw31).ArraySet1(_275_v103, 22)
			(_nw31).ArraySet1(_275_v103, 23)
			(_nw31).ArraySet1(_275_v103, 24)
			_284_v112 = _nw31
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_281_v109), 0))
			_ = _index26
			(_281_v109).ArraySet1(_284_v112, (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_281_v109), 0))
			_ = _index27
			var _rhs19 _dafny.Array = _284_v112
			_ = _rhs19
			var _rhs20 _dafny.Int = (_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int)
			_ = _rhs20
			var _rhs21 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, (_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int))
			_ = _rhs21
			var _lhs21 _dafny.Array = _281_v109
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(822), _dafny.ArrayLen((_281_v109), 0))
			_ = _lhs22
			var _lhs23 *GlobalState = globalState
			_ = _lhs23
			var _lhs24 *GlobalState = globalState
			_ = _lhs24
			(_lhs21).ArraySet1(_rhs19, (_lhs22).Int())
			_lhs23.F3 = _rhs20
			_lhs24.F3 = _rhs21
			var _285_v113 *C1
			_ = _285_v113
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__()
			_285_v113 = _nw32
			var _286_v114 D0
			_ = _286_v114
			_286_v114 = Companion_D0_.Create_DC1_(p3)
			_286_v114 = _286_v114
		} else {
			var _287_v115 _dafny.Array
			_ = _287_v115
			var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw33
			_287_v115 = _nw33
			var _288_v116 _dafny.Map
			_ = _288_v116
			_288_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _287_v115)
			(globalState).F2 = (func() _dafny.Array {
				if (_288_v116).Contains((func() _dafny.Int {
					if _270_v98.F14() {
						return (_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(-369)
				})()) {
					return (_288_v116).Get((func() _dafny.Int {
						if _270_v98.F14() {
							return (_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int)
						}
						return _dafny.IntOfInt64(-369)
					})()).(_dafny.Array)
				}
				return _287_v115
			})()
			var _289_v117 _dafny.Map
			_ = _289_v117
			_289_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int), _270_v98.F14()), _270_v98.F14())
			var _290_v118 _dafny.Map
			_ = _290_v118
			_290_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_289_v117).Cardinality(), _149_v0)
			var _291_v119 D8
			_ = _291_v119
			_291_v119 = Companion_D8_.Create_DC26_(Companion_Default___.Fm23(p1, _270_v98.F14(), _261_v89, _264_v92, globalState), _290_v118)
			var _292_v120 D1
			_ = _292_v120
			_292_v120 = Companion_D1_.Create_DC3_(_dafny.IntOfInt64(808))
			var _293_v121 D1
			_ = _293_v121
			_293_v121 = Companion_D1_.Create_DC6_(_292_v120)
			var _294_v122 D1
			_ = _294_v122
			_294_v122 = Companion_D1_.Create_DC6_(_292_v120)
			var _295_v123 _dafny.Map
			_ = _295_v123
			_295_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_291_v119, _293_v121)
			var _296_v124 D1
			_ = _296_v124
			_296_v124 = Companion_D1_.Create_DC6_((func() D1 {
				if (_295_v123).Contains(_291_v119) {
					return (_295_v123).Get(_291_v119).(D1)
				}
				return _292_v120
			})())
			var _297_v125 _dafny.Array
			_ = _297_v125
			var _nwElement0_6 D1 = _296_v124
			_ = _nwElement0_6
			var _nw34 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.One)
			_ = _nw34
			(_nw34).ArraySet1(_nwElement0_6, 0)
			_297_v125 = _nw34
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_297_v125), 0))
			_ = _index28
			(_297_v125).ArraySet1(_296_v124, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_297_v125), 0))
			_ = _index29
			(_297_v125).ArraySet1(_296_v124, (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
			_ = _index30
			var _rhs22 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(667), Companion_Default___.Fm0(_270_v98.F14(), _dafny.IntOfInt64(173), p1, _274_v102, globalState))
			_ = _rhs22
			var _rhs23 bool = p3
			_ = _rhs23
			var _lhs25 _dafny.Array = _275_v103
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))
			_ = _lhs26
			var _lhs27 T0 = _270_v98
			_ = _lhs27
			(_lhs25).ArraySet1(_rhs22, (_lhs26).Int())
			_lhs27.F14_set_(_rhs23)
			(globalState).F2 = _287_v115
			var _298_v126 _dafny.Set
			_ = _298_v126
			_298_v126 = _dafny.SetOf((_268_v96).Cardinality())
			var _rhs24 T0 = _270_v98
			_ = _rhs24
			var _rhs25 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_279_v107).Cardinality()).Plus((_275_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.ArrayLen((_275_v103), 0))).Int()).(_dafny.Int)), p1))
			_ = _rhs25
			var _rhs26 _dafny.CodePoint = (_266_v94).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(614), _dafny.ArrayLen((_266_v94), 0))).Int())
			_ = _rhs26
			var _rhs27 _dafny.Set = (_298_v126).Union((Companion_Default___.Fm6(globalState)).Union(_298_v126))
			_ = _rhs27
			var _lhs28 *GlobalState = globalState
			_ = _lhs28
			_270_v98 = _rhs24
			_lhs28.F3 = _rhs25
			_260_v88 = _rhs26
			_298_v126 = _rhs27
		}
	}
	var _299_v127 *C7
	_ = _299_v127
	var _nw35 *C7 = New_C7_()
	_ = _nw35
	_nw35.Ctor__()
	_299_v127 = _nw35
	_299_v127 = _299_v127
	r0 = (_dafny.Zero).Minus(p1)
	_149_v0 = (_151_v2).Dtor_cf22()
	var _300_v128 _dafny.Map
	_ = _300_v128
	_300_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p1)
	var _301_v129 *C4
	_ = _301_v129
	var _nw36 *C4 = New_C4_()
	_ = _nw36
	_nw36.Ctor__(p3, p3, _dafny.MultiSetOf(p2))
	_301_v129 = _nw36
	var _302_v130 D20
	_ = _302_v130
	_302_v130 = Companion_D20_.Create_DC60_(_301_v129)
	var _303_v131 _dafny.MultiSet
	_ = _303_v131
	_303_v131 = _dafny.MultiSetOf(p1, (Companion_Default___.Fm7(globalState)).Cardinality(), p1, (_300_v128).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_302_v130).Dtor_cf99(), p3)).Cardinality())
	var _304_v132 _dafny.Map
	_ = _304_v132
	_304_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p1, (_dafny.Zero).Minus(p1))), (_299_v127).Fm10(_300_v128, _303_v131, (_301_v129).F19(), _149_v0, globalState))
	if (func() bool {
		if (_304_v132).Contains(Companion_Default___.SafeDivisionInt(p1, p1)) {
			return (_304_v132).Get(Companion_Default___.SafeDivisionInt(p1, p1)).(bool)
		}
		return p3
	})() {
		(globalState).F3 = (p1).Plus(p1)
		var _305_v133 _dafny.MultiSet
		_ = _305_v133
		_305_v133 = _dafny.MultiSetOf(p2, !(p2), (_301_v129).F19())
		var _306_v134 *C5
		_ = _306_v134
		var _nw37 *C5 = New_C5_()
		_ = _nw37
		_nw37.Ctor__(_dafny.IntOfInt64(709), true, _305_v133)
		_306_v134 = _nw37
		_306_v134 = (func() *C5 {
			if (_301_v129).F19() {
				return _306_v134
			}
			return _306_v134
		})()
		var _307_v135 *C0
		_ = _307_v135
		var _nw38 *C0 = New_C0_()
		_ = _nw38
		_nw38.Ctor__((_301_v129).F19(), p2)
		_307_v135 = _nw38
		(globalState).F9 = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("wbhurb"), (_151_v2).Dtor_cf22())) || ((_307_v135).F23())
		var _308_v136 D1
		_ = _308_v136
		_308_v136 = Companion_D1_.Create_DC5_((_301_v129).F19(), (_307_v135).F23())
		var _309_v137 D1
		_ = _309_v137
		_309_v137 = Companion_D1_.Create_DC6_(Companion_D1_.Create_DC6_(_308_v136))
		var _pat_let_tv0 = _308_v136
		_ = _pat_let_tv0
		var _310_v138 _dafny.Map
		_ = _310_v138
		_310_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let2_0 D1) D1 {
			return func(_311_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 D1) D1 {
					return func(_312_dt__update_hcf8_h0 D1) D1 {
						return Companion_D1_.Create_DC6_(_312_dt__update_hcf8_h0)
					}(_pat_let3_0)
				}(Companion_D1_.Create_DC6_(_pat_let_tv0))
			}(_pat_let2_0)
		}(_309_v137), p3)
		_310_v138 = (_310_v138).Update(_309_v137, p3)
	} else {
		var _313_v139 _dafny.Array
		_ = _313_v139
		var _len0_0 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_0
		var _nw39 _dafny.Array
		_ = _nw39
		if _len0_0.Cmp(_dafny.Zero) == 0 {
			_nw39 = _dafny.NewArray(_len0_0)
		} else {
			var _init0 func(_dafny.Int) _dafny.Int = (func(_314_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_315_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_315_i3, _314_p1)
				}
			})(p1)
			_ = _init0
			var _element0_0 = _init0(_dafny.Zero)
			_ = _element0_0
			_nw39 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
			(_nw39).ArraySet1(_element0_0, 0)
			var _nativeLen0_0 = (_len0_0).Int()
			_ = _nativeLen0_0
			for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
				(_nw39).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
			}
		}
		_313_v139 = _nw39
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_313_v139), 0))
		_ = _index31
		(_313_v139).ArraySet1(p1, (_index31).Int())
		var _316_v140 _dafny.Sequence
		_ = _316_v140
		_316_v140 = _dafny.SeqOf((_dafny.IntOfInt64(421)).Plus(p1))
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_313_v139), 0))
		_ = _index32
		(_313_v139).ArraySet1(_dafny.IntOfUint32((_316_v140).Cardinality()), (_index32).Int())
		var _317_v141 _dafny.Set
		_ = _317_v141
		_317_v141 = _dafny.SetOf((_299_v127).Fm10(_300_v128, _303_v131, (_301_v129).F19(), _149_v0, globalState), p2)
		(globalState).F3 = (Companion_Default___.SafeModuloInt(p1, p1)).Minus(Companion_Default___.SafeModuloInt((_313_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_313_v139), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_317_v141).Cardinality())))
		(globalState).F9 = ((func() bool {
			if (_301_v129).F19() {
				return (_301_v129).F19()
			}
			return true
		})()) || ((_301_v129).F19())
		var _318_v142 D0
		_ = _318_v142
		_318_v142 = Companion_D0_.Create_DC0_(p2)
		var _319_v143 _dafny.MultiSet
		_ = _319_v143
		_319_v143 = _dafny.MultiSetOf((_318_v142).Dtor_cf0(), (_301_v129).F19())
		var _320_v144 T0
		_ = _320_v144
		var _nw40 *C3 = New_C3_()
		_ = _nw40
		_nw40.Ctor__((_313_v139).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_313_v139), 0))).Int()).(_dafny.Int), _149_v0, (_301_v129).F19(), _319_v143)
		_320_v144 = _nw40
		var _321_v145 *C6
		_ = _321_v145
		var _nw41 *C6 = New_C6_()
		_ = _nw41
		_nw41.Ctor__(p3, (_320_v144).F15())
		_321_v145 = _nw41
		var _322_v146 D21
		_ = _322_v146
		_322_v146 = Companion_D21_.Create_DC63_(_321_v145)
		var _323_v147 _dafny.Map
		_ = _323_v147
		_323_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_320_v144, (_322_v146).Dtor_cf101())
		_323_v147 = (_323_v147).Update(_320_v144, _321_v145)
		var _324_v148 _dafny.CodePoint
		_ = _324_v148
		_324_v148 = _dafny.CodePoint('o')
		var _325_v149 _dafny.Set
		_ = _325_v149
		_325_v149 = _dafny.SetOf((_320_v144).F15())
		var _rhs28 bool = (_325_v149).Contains((_320_v144).F15())
		_ = _rhs28
		var _rhs29 _dafny.CodePoint = _324_v148
		_ = _rhs29
		var _lhs29 *GlobalState = globalState
		_ = _lhs29
		_lhs29.F9 = _rhs28
		_324_v148 = _rhs29
	}
	var _326_v150 D5
	_ = _326_v150
	_326_v150 = Companion_D5_.Create_DC18_((_299_v127).Fm10(_300_v128, _303_v131, p2, _149_v0, globalState))
	var _pat_let_tv1 = p3
	_ = _pat_let_tv1
	var _pat_let_tv2 = p3
	_ = _pat_let_tv2
	var _pat_let_tv3 = p1
	_ = _pat_let_tv3
	var _pat_let_tv4 = p1
	_ = _pat_let_tv4
	var _pat_let_tv5 = p1
	_ = _pat_let_tv5
	var _pat_let_tv6 = p2
	_ = _pat_let_tv6
	var _pat_let_tv7 = _301_v129
	_ = _pat_let_tv7
	var _pat_let_tv8 = _301_v129
	_ = _pat_let_tv8
	var _pat_let_tv9 = _301_v129
	_ = _pat_let_tv9
	var _pat_let_tv10 = p2
	_ = _pat_let_tv10
	(globalState).F9 = func(_source8 D5) bool {
		if _source8.Is_DC17() {
			var _327___mcc_h8 _dafny.Sequence = _source8.Get_().(D5_DC17).Cf30
			_ = _327___mcc_h8
			var _328___mcc_h9 bool = _source8.Get_().(D5_DC17).Cf31
			_ = _328___mcc_h9
			var _329___mcc_h10 _dafny.MultiSet = _source8.Get_().(D5_DC17).Cf32
			_ = _329___mcc_h10
			var _330_cf32 _dafny.MultiSet = _329___mcc_h10
			_ = _330_cf32
			var _331_cf31 bool = _328___mcc_h9
			_ = _331_cf31
			var _332_cf30 _dafny.Sequence = _327___mcc_h8
			_ = _332_cf30
			return _pat_let_tv1
		} else if _source8.Is_DC18() {
			var _333___mcc_h11 bool = _source8.Get_().(D5_DC18).Cf33
			_ = _333___mcc_h11
			var _334_cf33 bool = _333___mcc_h11
			_ = _334_cf33
			return _pat_let_tv2
		} else if _source8.Is_DC16() {
			var _335___mcc_h12 _dafny.Array = _source8.Get_().(D5_DC16).Cf29
			_ = _335___mcc_h12
			var _336_cf29 _dafny.Array = _335___mcc_h12
			_ = _336_cf29
			return (_dafny.IntOfInt64(459)).Cmp((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv3, (Companion_D1_.Create_DC4_(_pat_let_tv4, _pat_let_tv5)).Dtor_cf4())).Cardinality())) <= 0
		} else {
			var _337___mcc_h13 D5 = _source8.Get_().(D5_DC19).Cf34
			_ = _337___mcc_h13
			var _338_cf34 D5 = _337___mcc_h13
			_ = _338_cf34
			return !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_pat_let_tv6), _dafny.SeqOf((_pat_let_tv7).F19(), (_pat_let_tv8).F19(), (_pat_let_tv9).F19(), _pat_let_tv10))
		}
	}(_326_v150)
	var _339_v151 _dafny.Sequence
	_ = _339_v151
	_339_v151 = _dafny.SeqOf((_301_v129).F19())
	r0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_339_v151).Cardinality()), p1)).Times(p1)
	var _340_v152 _dafny.Array
	_ = _340_v152
	var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
	_ = _nw42
	_340_v152 = _nw42
	var _341_v153 _dafny.Map
	_ = _341_v153
	_341_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _340_v152)
	r1 = _341_v153
	return r0, r1
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _342_v0 _dafny.Int
	_ = _342_v0
	_342_v0 = _dafny.IntOfInt64(206)
	var _343_v1 _dafny.MultiSet
	_ = _343_v1
	_343_v1 = _dafny.MultiSetOf(_342_v0)
	var _344_v2 _dafny.Array
	_ = _344_v2
	var _nw43 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
	_ = _nw43
	_344_v2 = _nw43
	var _345_v3 _dafny.Sequence
	_ = _345_v3
	_345_v3 = _dafny.UnicodeSeqOfUtf8Bytes("bwiunjo")
	var _346_v4 bool
	_ = _346_v4
	_346_v4 = true
	var _347_v5 _dafny.Set
	_ = _347_v5
	_347_v5 = _dafny.SetOf(_346_v4, _346_v4)
	var _348_v6 _dafny.Array
	_ = _348_v6
	var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
	_ = _nw44
	_348_v6 = _nw44
	var _349_v7 _dafny.Array
	_ = _349_v7
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(11)
	_ = _len0_1
	var _nw45 _dafny.Array
	_ = _nw45
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw45 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Sequence = (func(_350_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_351_i0 _dafny.Int) _dafny.Sequence {
				return _350_v3
			}
		})(_345_v3)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw45 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw45).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw45).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_349_v7 = _nw45
	var _352_globalState *GlobalState
	_ = _352_globalState
	var _nw46 *GlobalState = New_GlobalState_()
	_ = _nw46
	_nw46.Ctor__(_dafny.IntOfInt64(946), (_343_v1).Difference(_343_v1), _344_v2, _dafny.IntOfInt64(664), true, _345_v3, _347_v5, false, _dafny.IntOfInt64(762), false, _dafny.IntOfInt64(-427), _348_v6, _349_v7, _dafny.SetOf(_346_v4, _346_v4, false))
	_352_globalState = _nw46
	var _353_v8 _dafny.Sequence
	_ = _353_v8
	_353_v8 = _dafny.SeqOf(_346_v4)
	_353_v8 = _dafny.Companion_Sequence_.Concatenate(_353_v8, _353_v8)
	var _354_v9 _dafny.Map
	_ = _354_v9
	_354_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_342_v0).Cmp(_342_v0) != 0, _347_v5)
	_354_v9 = (_354_v9).Update(_346_v4, _347_v5)
	var _355_v10 _dafny.Sequence
	_ = _355_v10
	_355_v10 = _dafny.SeqOf(_342_v0, _342_v0)
	var _hi0 _dafny.Int = _342_v0
	_ = _hi0
	for _356_i1 := Companion_Default___.Fm0(_346_v4, _342_v0, _342_v0, (_dafny.MultiSetFromSeq(_355_v10)).Update(_342_v0, Companion_Default___.Abs(_342_v0)), _352_globalState); _356_i1.Cmp(_hi0) < 0; _356_i1 = _356_i1.Plus(_dafny.One) {
		var _357_v11 _dafny.Map
		_ = _357_v11
		_357_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, _349_v7)
		_349_v7 = (func() _dafny.Array {
			if (_357_v11).Contains((_342_v0).Times(_dafny.IntOfInt64(-478))) {
				return (_357_v11).Get((_342_v0).Times(_dafny.IntOfInt64(-478))).(_dafny.Array)
			}
			return _349_v7
		})()
		var _358_v12 _dafny.Map
		_ = _358_v12
		_358_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, Companion_Default___.SafeModuloInt(_356_i1, _356_i1))
		var _359_v13 _dafny.Array
		_ = _359_v13
		var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
		_ = _nw47
		_359_v13 = _nw47
		var _360_v14 _dafny.MultiSet
		_ = _360_v14
		_360_v14 = _dafny.MultiSetOf(_359_v13)
		var _361_v15 _dafny.Map
		_ = _361_v15
		_361_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, _359_v13)
		var _362_v16 _dafny.Set
		_ = _362_v16
		_362_v16 = _dafny.SetOf(_dafny.IntOfUint32((_355_v10).Cardinality()), _342_v0, _342_v0)
		var _rhs30 _dafny.Int = Companion_Default___.SafeModuloInt((_356_i1).Minus(_356_i1), Companion_Default___.SafeDivisionInt(_342_v0, _342_v0))
		_ = _rhs30
		var _rhs31 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_345_v3, _dafny.Companion_Sequence_.Concatenate(_345_v3, _345_v3))).Cardinality())
		_ = _rhs31
		var _rhs32 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((((func() _dafny.MultiSet {
			if _346_v4 {
				return _dafny.MultiSetOf(_359_v13)
			}
			return _360_v14
		})()).Update((func() _dafny.Array {
			if (_361_v15).Contains(_356_i1) {
				return (_361_v15).Get(_356_i1).(_dafny.Array)
			}
			return _359_v13
		})(), Companion_Default___.Abs((_dafny.Zero).Minus(_342_v0)))).Cardinality()))
		_ = _rhs32
		var _rhs33 _dafny.Map = (_358_v12).Merge(Companion_Default___.Fm1(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_356_i1, _346_v4)).Cardinality(), _346_v4, _362_v16, _352_globalState))
		_ = _rhs33
		var _lhs30 *GlobalState = _352_globalState
		_ = _lhs30
		var _lhs31 *GlobalState = _352_globalState
		_ = _lhs31
		_342_v0 = _rhs30
		_lhs30.F3 = _rhs31
		_lhs31.F3 = _rhs32
		_358_v12 = _rhs33
		var _363_v17 _dafny.Map
		_ = _363_v17
		_363_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, _346_v4)
		if (func() bool {
			if (_363_v17).Contains(_342_v0) {
				return (_363_v17).Get(_342_v0).(bool)
			}
			return true
		})() {
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_359_v13), 0))
			_ = _index33
			(_359_v13).ArraySet1(_342_v0, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_359_v13), 0))
			_ = _index34
			(_359_v13).ArraySet1(_dafny.IntOfInt64(971), (_index34).Int())
			var _rhs34 _dafny.Sequence = _345_v3
			_ = _rhs34
			var _rhs35 bool = _346_v4
			_ = _rhs35
			var _lhs32 *GlobalState = _352_globalState
			_ = _lhs32
			_345_v3 = _rhs34
			_lhs32.F9 = _rhs35
			(_352_globalState).F3 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-824))
			_346_v4 = _346_v4
			var _364_v18 *C1
			_ = _364_v18
			var _nw48 *C1 = New_C1_()
			_ = _nw48
			_nw48.Ctor__()
			_364_v18 = _nw48
		} else {
			var _365_v19 _dafny.CodePoint
			_ = _365_v19
			_365_v19 = _dafny.CodePoint('r')
			var _366_v20 _dafny.Map
			_ = _366_v20
			_366_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_365_v19, _342_v0)
			_366_v20 = (_366_v20).Update(_365_v19, _342_v0)
			var _367_v21 D5
			_ = _367_v21
			_367_v21 = Companion_D5_.Create_DC16_(_359_v13)
			var _pat_let_tv11 = _359_v13
			_ = _pat_let_tv11
			var _pat_let_tv12 = _359_v13
			_ = _pat_let_tv12
			var _pat_let_tv13 = _359_v13
			_ = _pat_let_tv13
			var _368_v22 _dafny.Array
			_ = _368_v22
			var _nwElement0_7 D5 = _367_v21
			_ = _nwElement0_7
			var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(15))
			_ = _nw49
			(_nw49).ArraySet1(_nwElement0_7, 0)
			(_nw49).ArraySet1(_367_v21, 1)
			(_nw49).ArraySet1(_367_v21, 2)
			(_nw49).ArraySet1(_367_v21, 3)
			(_nw49).ArraySet1(_367_v21, 4)
			(_nw49).ArraySet1(func(_pat_let4_0 D5) D5 {
				return func(_369_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let5_0 _dafny.Array) D5 {
						return func(_370_dt__update_hcf29_h0 _dafny.Array) D5 {
							return Companion_D5_.Create_DC16_(_370_dt__update_hcf29_h0)
						}(_pat_let5_0)
					}(_pat_let_tv11)
				}(_pat_let4_0)
			}(_367_v21), 5)
			(_nw49).ArraySet1(func(_pat_let6_0 D5) D5 {
				return func(_373_dt__update__tmp_h2 D5) D5 {
					return func(_pat_let9_0 _dafny.Array) D5 {
						return func(_374_dt__update_hcf29_h2 _dafny.Array) D5 {
							return Companion_D5_.Create_DC16_(_374_dt__update_hcf29_h2)
						}(_pat_let9_0)
					}(_pat_let_tv13)
				}(_pat_let6_0)
			}(func(_pat_let7_0 D5) D5 {
				return func(_371_dt__update__tmp_h1 D5) D5 {
					return func(_pat_let8_0 _dafny.Array) D5 {
						return func(_372_dt__update_hcf29_h1 _dafny.Array) D5 {
							return Companion_D5_.Create_DC16_(_372_dt__update_hcf29_h1)
						}(_pat_let8_0)
					}(_pat_let_tv12)
				}(_pat_let7_0)
			}(_367_v21)), 6)
			(_nw49).ArraySet1(_367_v21, 7)
			(_nw49).ArraySet1(_367_v21, 8)
			(_nw49).ArraySet1(_367_v21, 9)
			(_nw49).ArraySet1(Companion_D5_.Create_DC16_(_359_v13), 10)
			(_nw49).ArraySet1(_367_v21, 11)
			(_nw49).ArraySet1(Companion_D5_.Create_DC16_(_359_v13), 12)
			(_nw49).ArraySet1(_367_v21, 13)
			(_nw49).ArraySet1(_367_v21, 14)
			_368_v22 = _nw49
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_368_v22), 0))
			_ = _index35
			(_368_v22).ArraySet1(_367_v21, (_index35).Int())
			var _375_v23 _dafny.MultiSet
			_ = _375_v23
			_375_v23 = _dafny.MultiSetOf((func() bool {
				if _346_v4 {
					return _346_v4
				}
				return false
			})(), !_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_356_i1), _342_v0), _346_v4, (_356_i1).Cmp(_342_v0) >= 0, _346_v4)
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_344_v2), 0))
			_ = _index36
			(_344_v2).ArraySet1(_346_v4, (_index36).Int())
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_368_v22), 0))
			_ = _index37
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_344_v2), 0))
			_ = _index38
			var _rhs36 _dafny.Int = (_dafny.Zero).Minus(_342_v0)
			_ = _rhs36
			var _rhs37 D5 = _367_v21
			_ = _rhs37
			var _rhs38 _dafny.MultiSet = ((func() _dafny.MultiSet {
				if _346_v4 {
					return _375_v23
				}
				return _375_v23
			})()).Intersection(_375_v23)
			_ = _rhs38
			var _rhs39 bool = _346_v4
			_ = _rhs39
			var _lhs33 *GlobalState = _352_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _368_v22
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_368_v22), 0))
			_ = _lhs35
			var _lhs36 _dafny.Array = _344_v2
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_344_v2), 0))
			_ = _lhs37
			_lhs33.F3 = _rhs36
			(_lhs34).ArraySet1(_rhs37, (_lhs35).Int())
			_375_v23 = _rhs38
			(_lhs36).ArraySet1(_rhs39, (_lhs37).Int())
			(_352_globalState).F3 = (func() _dafny.Int {
				if _346_v4 {
					return Companion_Default___.SafeModuloInt(_342_v0, _356_i1)
				}
				return _356_i1
			})()
			var _376_v24 _dafny.Map
			_ = _376_v24
			_376_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_356_i1), _342_v0)
			_376_v24 = (_376_v24).Update((_342_v0).Plus(_356_i1), _342_v0)
			_346_v4 = _346_v4
		}
		var _377_v25 T0
		_ = _377_v25
		var _nw50 *C5 = New_C5_()
		_ = _nw50
		_nw50.Ctor__(_dafny.IntOfUint32((_353_v8).Cardinality()), _346_v4, _dafny.MultiSetOf(false, _346_v4))
		_377_v25 = _nw50
		var _378_v26 _dafny.Sequence
		_ = _378_v26
		_378_v26 = _dafny.SeqOf(_377_v25, _377_v25)
		var _379_v27 _dafny.Array
		_ = _379_v27
		var _nwElement0_8 T0 = _377_v25
		_ = _nwElement0_8
		var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(24))
		_ = _nw51
		(_nw51).ArraySet1(_nwElement0_8, 0)
		(_nw51).ArraySet1(_377_v25, 1)
		(_nw51).ArraySet1(_377_v25, 2)
		(_nw51).ArraySet1(_377_v25, 3)
		(_nw51).ArraySet1((_378_v26).Select((Companion_Default___.SafeIndex(_356_i1, _dafny.IntOfUint32((_378_v26).Cardinality()))).Uint32()).(T0), 4)
		(_nw51).ArraySet1(_377_v25, 5)
		(_nw51).ArraySet1(_377_v25, 6)
		(_nw51).ArraySet1(_377_v25, 7)
		(_nw51).ArraySet1(_377_v25, 8)
		(_nw51).ArraySet1(_377_v25, 9)
		(_nw51).ArraySet1(_377_v25, 10)
		(_nw51).ArraySet1(_377_v25, 11)
		(_nw51).ArraySet1(_377_v25, 12)
		(_nw51).ArraySet1(_377_v25, 13)
		(_nw51).ArraySet1(_377_v25, 14)
		(_nw51).ArraySet1(_377_v25, 15)
		(_nw51).ArraySet1(_377_v25, 16)
		(_nw51).ArraySet1(_377_v25, 17)
		(_nw51).ArraySet1(_377_v25, 18)
		(_nw51).ArraySet1(_377_v25, 19)
		(_nw51).ArraySet1(_377_v25, 20)
		(_nw51).ArraySet1(_377_v25, 21)
		(_nw51).ArraySet1(_377_v25, 22)
		(_nw51).ArraySet1(_377_v25, 23)
		_379_v27 = _nw51
		var _380_v28 _dafny.MultiSet
		_ = _380_v28
		_380_v28 = _dafny.MultiSetOf(_379_v27, _379_v27, _379_v27)
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_344_v2), 0))
		_ = _index39
		(_344_v2).ArraySet1(!((_380_v28).IsSubsetOf(_380_v28)), (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(802), _dafny.ArrayLen((_344_v2), 0))
		_ = _index40
		(_344_v2).ArraySet1(_346_v4, (_index40).Int())
	}
	var _381_v29 D13
	_ = _381_v29
	_381_v29 = Companion_D13_.Create_DC43_(Companion_Default___.Fm51(_346_v4, _346_v4, _346_v4, _342_v0, _352_globalState))
	var _382_v30 _dafny.Map
	_ = _382_v30
	_382_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, _342_v0)
	var _383_v31 _dafny.CodePoint
	_ = _383_v31
	_383_v31 = _dafny.CodePoint('m')
	var _384_v32 _dafny.Int
	_ = _384_v32
	var _385_v33 _dafny.Map
	_ = _385_v33
	var _out3 _dafny.Int
	_ = _out3
	var _out4 _dafny.Map
	_ = _out4
	_out3, _out4 = Companion_Default___.M11(_381_v29, ((_382_v30).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_345_v3, (Companion_Default___.SafeIndex(_342_v0, _dafny.IntOfUint32((_345_v3).Cardinality()))).Uint32(), _383_v31)).Cardinality()), _342_v0)).Cardinality(), _346_v4, _346_v4, _352_globalState)
	_384_v32 = _out3
	_385_v33 = _out4
	if (func() bool {
		if !(!((true) || (_346_v4))) {
			return _346_v4
		}
		return (_384_v32).Cmp(_384_v32) <= 0
	})() {
		var _386_v34 _dafny.Array
		_ = _386_v34
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw52
		_386_v34 = _nw52
		var _387_v35 D16
		_ = _387_v35
		_387_v35 = Companion_D16_.Create_DC52_(_342_v0, true, _386_v34, _346_v4, _342_v0)
		var _388_v36 D2
		_ = _388_v36
		_388_v36 = Companion_D2_.Create_DC9_(_346_v4, _342_v0, _dafny.CodePoint('l'))
		var _389_v37 _dafny.Int
		_ = _389_v37
		var _390_v38 _dafny.Map
		_ = _390_v38
		var _out5 _dafny.Int
		_ = _out5
		var _out6 _dafny.Map
		_ = _out6
		_out5, _out6 = Companion_Default___.M11(_381_v29, (_384_v32).Times((_387_v35).Dtor_cf88()), _dafny.Companion_Sequence_.Contains(_345_v3, (_388_v36).Dtor_cf15()), _346_v4, _352_globalState)
		_389_v37 = _out5
		_390_v38 = _out6
		var _391_v39 _dafny.Array
		_ = _391_v39
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(15)
		_ = _len0_2
		var _nw53 _dafny.Array
		_ = _nw53
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw53 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_392_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_393_i2 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_393_i2, (Companion_D1_.Create_DC4_(_392_v0, _dafny.Zero)).Dtor_cf5())
				}
			})(_342_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw53 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw53).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw53).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_391_v39 = _nw53
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))
		_ = _index41
		(_391_v39).ArraySet1(_384_v32, (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))
		_ = _index42
		(_391_v39).ArraySet1((_dafny.Zero).Minus(_384_v32), (_index42).Int())
		if _346_v4 {
			var _394_v40 _dafny.Map
			_ = _394_v40
			_394_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, _353_v8)
			var _395_v41 _dafny.Map
			_ = _395_v41
			_395_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v31, (_343_v1).Cardinality())
			var _396_v42 _dafny.Set
			_ = _396_v42
			_396_v42 = _dafny.SetOf(_355_v10)
			var _397_v43 _dafny.Map
			_ = _397_v43
			_397_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("pshth"), _dafny.MultiSetFromSeq(_353_v8))
			var _398_v44 _dafny.MultiSet
			_ = _398_v44
			_398_v44 = _dafny.MultiSetOf(_346_v4)
			var _399_v45 _dafny.Map
			_ = _399_v45
			_399_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, (_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int))
			var _400_v46 _dafny.Map
			_ = _400_v46
			_400_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_399_v45).Contains(_346_v4) {
					return (_399_v45).Get(_346_v4).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_355_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-625))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}(func(_401_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				}))).Cardinality()), _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32(), _389_v37)).Cardinality())
			})(), _346_v4)
			var _402_v47 D1
			_ = _402_v47
			_402_v47 = Companion_D1_.Create_DC5_(_346_v4, false)
			var _pat_let_tv14 = _346_v4
			_ = _pat_let_tv14
			var _403_v48 D11
			_ = _403_v48
			_403_v48 = Companion_D11_.Create_DC34_(_398_v44, _343_v1, _344_v2, func(_pat_let10_0 D1) D1 {
				return func(_404_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let11_0 bool) D1 {
						return func(_405_dt__update_hcf7_h0 bool) D1 {
							return Companion_D1_.Create_DC5_((_404_dt__update__tmp_h3).Dtor_cf6(), _405_dt__update_hcf7_h0)
						}(_pat_let11_0)
					}(_pat_let_tv14)
				}(_pat_let10_0)
			}(_402_v47))
			var _pat_let_tv15 = _398_v44
			_ = _pat_let_tv15
			var _pat_let_tv16 = _344_v2
			_ = _pat_let_tv16
			var _406_v49 _dafny.Array
			_ = _406_v49
			var _nwElement0_9 _dafny.MultiSet = _dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if (_394_v40).Contains(Companion_Default___.Fm23(_dafny.IntOfUint32((_353_v8).Cardinality()), _346_v4, _395_v41, _396_v42, _352_globalState)) {
					return (_394_v40).Get(Companion_Default___.Fm23(_dafny.IntOfUint32((_353_v8).Cardinality()), _346_v4, _395_v41, _396_v42, _352_globalState)).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Update(_353_v8, (Companion_Default___.SafeIndex((_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_353_v8).Cardinality()))).Uint32(), Companion_Default___.Fm23(_dafny.IntOfUint32((_353_v8).Cardinality()), !(_346_v4), Companion_Default___.Fm24(_389_v37, _346_v4, _346_v4, (_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int), _352_globalState), _396_v42, _352_globalState))
			})())
			_ = _nwElement0_9
			var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(26))
			_ = _nw54
			(_nw54).ArraySet1(_nwElement0_9, 0)
			(_nw54).ArraySet1(((func() _dafny.MultiSet {
				if (_397_v43).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_407_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_408_i3 _dafny.Int) _dafny.CodePoint {
						return _407_v31
					}
				})(_383_v31)))) {
					return (_397_v43).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-948))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg46 _dafny.Int) interface{} {
							return coer46(arg46)
						}
					}((func(_409_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_410_i3 _dafny.Int) _dafny.CodePoint {
							return _409_v31
						}
					})(_383_v31)))).(_dafny.MultiSet)
				}
				return _398_v44
			})()).Intersection(_dafny.MultiSetOf(_346_v4)), 1)
			(_nw54).ArraySet1(_398_v44, 2)
			(_nw54).ArraySet1(_398_v44, 3)
			(_nw54).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(!(_346_v4), _346_v4)), 4)
			(_nw54).ArraySet1((Companion_Default___.Fm33(!(false), _352_globalState)).Union(_398_v44), 5)
			(_nw54).ArraySet1((_398_v44).Intersection(_398_v44), 6)
			(_nw54).ArraySet1(_dafny.MultiSetFromSeq(_353_v8), 7)
			(_nw54).ArraySet1(_398_v44, 8)
			(_nw54).ArraySet1(_dafny.MultiSetOf(false), 9)
			(_nw54).ArraySet1(_398_v44, 10)
			(_nw54).ArraySet1((_398_v44).Update(_346_v4, Companion_Default___.Abs(_384_v32)), 11)
			(_nw54).ArraySet1(_dafny.MultiSetOf(_346_v4, !(_346_v4), (func() bool {
				if (_400_v46).Contains(_389_v37) {
					return (_400_v46).Get(_389_v37).(bool)
				}
				return false
			})()), 12)
			(_nw54).ArraySet1(_398_v44, 13)
			(_nw54).ArraySet1(_dafny.MultiSetOf(true), 14)
			(_nw54).ArraySet1(((_398_v44).Update(_346_v4, Companion_Default___.Abs(_389_v37))).Update(_346_v4, Companion_Default___.Abs(_342_v0)), 15)
			(_nw54).ArraySet1((_398_v44).Union(_398_v44), 16)
			(_nw54).ArraySet1(_398_v44, 17)
			(_nw54).ArraySet1(_398_v44, 18)
			(_nw54).ArraySet1(_398_v44, 19)
			(_nw54).ArraySet1(_398_v44, 20)
			(_nw54).ArraySet1(_dafny.MultiSetOf(_346_v4, !(_346_v4), true, _346_v4), 21)
			(_nw54).ArraySet1(_398_v44, 22)
			(_nw54).ArraySet1((func(_pat_let12_0 D11) D11 {
				return func(_411_dt__update__tmp_h4 D11) D11 {
					return func(_pat_let13_0 _dafny.MultiSet) D11 {
						return func(_412_dt__update_hcf58_h0 _dafny.MultiSet) D11 {
							return func(_pat_let14_0 _dafny.Array) D11 {
								return func(_413_dt__update_hcf60_h0 _dafny.Array) D11 {
									return Companion_D11_.Create_DC34_(_412_dt__update_hcf58_h0, (_411_dt__update__tmp_h4).Dtor_cf59(), _413_dt__update_hcf60_h0, (_411_dt__update__tmp_h4).Dtor_cf61())
								}(_pat_let14_0)
							}(_pat_let_tv16)
						}(_pat_let13_0)
					}(_pat_let_tv15)
				}(_pat_let12_0)
			}(_403_v48)).Dtor_cf58(), 23)
			(_nw54).ArraySet1((_398_v44).Difference(_398_v44), 24)
			(_nw54).ArraySet1(_398_v44, 25)
			_406_v49 = _nw54
			var _414_v50 _dafny.Array
			_ = _414_v50
			var _nw55 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
			_ = _nw55
			_414_v50 = _nw55
			var _415_v51 _dafny.Map
			_ = _415_v51
			_415_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_414_v50, _389_v37)
			var _416_v52 _dafny.Array
			_ = _416_v52
			var _len0_3 _dafny.Int = _dafny.One
			_ = _len0_3
			var _nw56 _dafny.Array
			_ = _nw56
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw56 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_417_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_418_i5 _dafny.Int) _dafny.CodePoint {
						return _417_v31
					}
				})(_383_v31)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw56 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw56).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw56).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_416_v52 = _nw56
			var _419_v53 D16
			_ = _419_v53
			_419_v53 = Companion_D16_.Create_DC51_(_416_v52)
			var _rhs40 _dafny.Int = Companion_Default___.Fm0((_384_v32).Cmp((_415_v51).Cardinality()) >= 0, (_dafny.Zero).Minus(_384_v32), _384_v32, _343_v1, _352_globalState)
			_ = _rhs40
			var _rhs41 _dafny.Int = _389_v37
			_ = _rhs41
			var _rhs42 _dafny.CodePoint = _383_v31
			_ = _rhs42
			var _rhs43 _dafny.Array = _406_v49
			_ = _rhs43
			var _rhs44 bool = (_353_v8).Select((Companion_Default___.SafeIndex(((_dafny.MultiSetOf(_416_v52)).Update((_419_v53).Dtor_cf83(), Companion_Default___.Abs(_dafny.IntOfInt64(35)))).Cardinality(), _dafny.IntOfUint32((_353_v8).Cardinality()))).Uint32()).(bool)
			_ = _rhs44
			var _lhs38 *GlobalState = _352_globalState
			_ = _lhs38
			var _lhs39 *GlobalState = _352_globalState
			_ = _lhs39
			_lhs38.F3 = _rhs40
			_389_v37 = _rhs41
			_383_v31 = _rhs42
			_406_v49 = _rhs43
			_lhs39.F9 = _rhs44
			var _420_v54 _dafny.Array
			_ = _420_v54
			var _nw57 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(7))
			_ = _nw57
			_420_v54 = _nw57
			var _421_v55 D14
			_ = _421_v55
			_421_v55 = Companion_D14_.Create_DC44_(_349_v7)
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_420_v54), 0))
			_ = _index43
			(_420_v54).ArraySet1(_421_v55, (_index43).Int())
			var _pat_let_tv17 = _349_v7
			_ = _pat_let_tv17
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(715), _dafny.ArrayLen((_420_v54), 0))
			_ = _index44
			(_420_v54).ArraySet1(func(_pat_let15_0 D14) D14 {
				return func(_422_dt__update__tmp_h5 D14) D14 {
					return func(_pat_let16_0 _dafny.Array) D14 {
						return func(_423_dt__update_hcf71_h0 _dafny.Array) D14 {
							return Companion_D14_.Create_DC44_(_423_dt__update_hcf71_h0)
						}(_pat_let16_0)
					}(_pat_let_tv17)
				}(_pat_let15_0)
			}(_421_v55), (_index44).Int())
			var _424_v56 _dafny.Map
			_ = _424_v56
			_424_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int), (_398_v44).Update(_346_v4, Companion_Default___.Abs(_384_v32)))
			var _425_v57 *C5
			_ = _425_v57
			var _nw58 *C5 = New_C5_()
			_ = _nw58
			_nw58.Ctor__(_389_v37, _346_v4, ((func() _dafny.MultiSet {
				if (_424_v56).Contains(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_426_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_427_i6 _dafny.Int) _dafny.CodePoint {
						return _426_v31
					}
				})(_383_v31)))).Cardinality())) {
					return (_424_v56).Get(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}((func(_428_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_429_i6 _dafny.Int) _dafny.CodePoint {
							return _428_v31
						}
					})(_383_v31)))).Cardinality())).(_dafny.MultiSet)
				}
				return _398_v44
			})()).Union(_dafny.MultiSetFromSeq(_353_v8)))
			_425_v57 = _nw58
			_383_v31 = _383_v31
			var _430_v58 _dafny.Array
			_ = _430_v58
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_4
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Sequence = (func(_431_v4 bool, _432_v3 _dafny.Sequence, _433_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_434_i7 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_431_v4, _431_v4), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_432_v3).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_431_v4, _431_v4)).Cardinality()))).Uint32(), _431_v4), _433_v8)
					}
				})(_346_v4, _345_v3, _353_v8)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw59 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw59).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw59).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_430_v58 = _nw59
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_430_v58), 0))
			_ = _index45
			(_430_v58).ArraySet1(_353_v8, (_index45).Int())
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_430_v58), 0))
			_ = _index46
			(_430_v58).ArraySet1(_353_v8, (_index46).Int())
		} else {
			(_352_globalState).F9 = _346_v4
			_391_v39 = _391_v39
			(_352_globalState).F3 = Companion_Default___.Fm0(!(_346_v4), _389_v37, _342_v0, (_343_v1).Intersection(_343_v1), _352_globalState)
			var _435_v59 _dafny.Array
			_ = _435_v59
			var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw60
			_435_v59 = _nw60
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_435_v59), 0))
			_ = _index47
			(_435_v59).ArraySet1(_353_v8, (_index47).Int())
			var _436_v60 _dafny.Map
			_ = _436_v60
			_436_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_353_v8, _dafny.SeqOf(false)), (func() _dafny.Int {
				if _346_v4 {
					return (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm53(_346_v4, _346_v4, _352_globalState)).Cardinality()))
				}
				return (_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int)
			})())
			var _437_v61 *C4
			_ = _437_v61
			var _nw61 *C4 = New_C4_()
			_ = _nw61
			_nw61.Ctor__(_346_v4, _346_v4, (_dafny.MultiSetOf(_346_v4, _346_v4)).Update(_346_v4, Companion_Default___.Abs((_343_v1).Cardinality())))
			_437_v61 = _nw61
			var _438_v62 _dafny.Map
			_ = _438_v62
			_438_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_437_v61, _384_v32)
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_435_v59), 0))
			_ = _index48
			var _rhs45 _dafny.Sequence = Companion_Default___.Fm38(_352_globalState)
			_ = _rhs45
			var _rhs46 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_v8, (_438_v62).Cardinality())).Update(_353_v8, _dafny.IntOfInt64(267))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_353_v8, (Companion_Default___.SafeIndex(_342_v0, _dafny.IntOfUint32((_353_v8).Cardinality()))).Uint32(), (_437_v61).F19()), _384_v32))
			_ = _rhs46
			var _lhs40 _dafny.Array = _435_v59
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_435_v59), 0))
			_ = _lhs41
			(_lhs40).ArraySet1(_rhs45, (_lhs41).Int())
			_436_v60 = _rhs46
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))
			_ = _index49
			(_391_v39).ArraySet1(_dafny.IntOfInt64(-956), (_index49).Int())
		}
		var _439_v64 _dafny.Map
		_ = _439_v64
		_439_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_355_v10).Select((Companion_Default___.SafeIndex((_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32()).(_dafny.Int), _dafny.SetOf(_346_v4))
		var _440_v65 D0
		_ = _440_v65
		_440_v65 = Companion_D0_.Create_DC2_(Companion_Default___.Fm37(func() _dafny.Set {
			var _coll35 = _dafny.NewBuilder()
			_ = _coll35
			for _iter35 := _dafny.Iterate((_dafny.Companion_Sequence_.Update(_355_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_345_v3).Cardinality()), _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32(), (_343_v1).Cardinality())).Elements()); ; {
				_compr_35, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _441_v63 _dafny.Int
				_441_v63 = interface{}(_compr_35).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_355_v10, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_345_v3).Cardinality()), _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32(), (_343_v1).Cardinality()), _441_v63) {
					_coll35.Add(Companion_Default___.SafeDivisionInt(_441_v63, _dafny.IntOfInt64(240)))
				}
			}
			return _coll35.ToSet()
		}(), _439_v64, _346_v4, _346_v4, _352_globalState))
		var _source9 D0 = _440_v65
		_ = _source9
		if _source9.Is_DC1() {
			var _442___mcc_h0 bool = _source9.Get_().(D0_DC1).Cf1
			_ = _442___mcc_h0
			var _443_cf1 bool = _442___mcc_h0
			_ = _443_cf1
			(_352_globalState).F3 = _389_v37
			_384_v32 = (_342_v0).Plus(Companion_Default___.Fm0(!(false), _389_v37, (_347_v5).Cardinality(), _343_v1, _352_globalState))
			var _444_v66 _dafny.Int
			_ = _444_v66
			var _445_v67 _dafny.Map
			_ = _445_v67
			var _out7 _dafny.Int
			_ = _out7
			var _out8 _dafny.Map
			_ = _out8
			_out7, _out8 = Companion_Default___.M11(_381_v29, _389_v37, (_dafny.IntOfInt64(633)).Cmp((_dafny.Zero).Minus(_384_v32)) < 0, false, _352_globalState)
			_444_v66 = _out7
			_445_v67 = _out8
			_444_v66 = _384_v32
		} else if _source9.Is_DC0() {
			var _446___mcc_h1 bool = _source9.Get_().(D0_DC0).Cf0
			_ = _446___mcc_h1
			var _447_cf0 bool = _446___mcc_h1
			_ = _447_cf0
			_346_v4 = _447_cf0
			_389_v37 = _384_v32
			var _448_v68 _dafny.MultiSet
			_ = _448_v68
			_448_v68 = _dafny.MultiSetOf(_346_v4, _447_cf0)
			var _449_v69 _dafny.Set
			_ = _449_v69
			_449_v69 = _dafny.SetOf(_384_v32, (_dafny.Zero).Minus(_342_v0), (func() _dafny.Int {
				if (_448_v68).Contains(false) {
					return (_448_v68).Multiplicity(false)
				}
				return _389_v37
			})())
			(_352_globalState).F9 = (_449_v69).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(981), _342_v0))
			var _450_v70 _dafny.Map
			_ = _450_v70
			_450_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v32, _dafny.SeqOf(_346_v4, _447_cf0))
			(_352_globalState).F9 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if (_450_v70).Contains((_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int)) {
					return (_450_v70).Get((_391_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_391_v39), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
				}
				return _353_v8
			})(), _353_v8), _353_v8)
		} else {
			var _451___mcc_h2 D0 = _source9.Get_().(D0_DC2).Cf2
			_ = _451___mcc_h2
			var _452_cf2 D0 = _451___mcc_h2
			_ = _452_cf2
			var _453_v71 D13
			_ = _453_v71
			_453_v71 = Companion_D13_.Create_DC42_(_383_v31, !(true))
			var _454_v72 _dafny.Int
			_ = _454_v72
			var _455_v73 _dafny.Map
			_ = _455_v73
			var _out9 _dafny.Int
			_ = _out9
			var _out10 _dafny.Map
			_ = _out10
			_out9, _out10 = Companion_Default___.M11(Companion_D13_.Create_DC43_(_453_v71), _389_v37, _346_v4, _346_v4, _352_globalState)
			_454_v72 = _out9
			_455_v73 = _out10
			_346_v4 = (_343_v1).Equals(_343_v1)
			var _456_v74 D5
			_ = _456_v74
			_456_v74 = Companion_D5_.Create_DC16_(_391_v39)
			_391_v39 = (_456_v74).Dtor_cf29()
			(_352_globalState).F9 = _346_v4
		}
		_342_v0 = _342_v0
	} else {
		(_352_globalState).F9 = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_353_v8, _353_v8), (func() bool {
			if _346_v4 {
				return !(_346_v4)
			}
			return _346_v4
		})())
		var _457_v75 _dafny.Map
		_ = _457_v75
		_457_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v31, _342_v0)
		var _458_v77 _dafny.Int
		_ = _458_v77
		var _459_v78 _dafny.Map
		_ = _459_v78
		var _out11 _dafny.Int
		_ = _out11
		var _out12 _dafny.Map
		_ = _out12
		_out11, _out12 = Companion_Default___.M11(_381_v29, _342_v0, _346_v4, Companion_Default___.Fm23(_342_v0, _346_v4, _457_v75, func() _dafny.Set {
			var _coll36 = _dafny.NewBuilder()
			_ = _coll36
			for _iter36 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_342_v0), _342_v0)).Keys().Elements()); ; {
				_compr_36, _ok36 := _iter36()
				if !_ok36 {
					break
				}
				var _460_v76 _dafny.Sequence
				_460_v76 = interface{}(_compr_36).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_342_v0), _342_v0)).Contains(_460_v76) {
					_coll36.Add(_460_v76)
				}
			}
			return _coll36.ToSet()
		}(), _352_globalState), _352_globalState)
		_458_v77 = _out11
		_459_v78 = _out12
		var _461_v79 *C4
		_ = _461_v79
		var _nw62 *C4 = New_C4_()
		_ = _nw62
		_nw62.Ctor__(_346_v4, _346_v4, Companion_Default___.Fm33(_346_v4, _352_globalState))
		_461_v79 = _nw62
		var _462_v80 _dafny.MultiSet
		_ = _462_v80
		_462_v80 = _dafny.MultiSetOf(_461_v79)
		var _463_v81 _dafny.MultiSet
		_ = _463_v81
		_463_v81 = _dafny.MultiSetOf(_346_v4, _346_v4)
		var _464_v82 _dafny.Map
		_ = _464_v82
		_464_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_462_v80, (func() _dafny.Int {
			if (_463_v81).Contains((_461_v79).F19()) {
				return (_463_v81).Multiplicity((_461_v79).F19())
			}
			return _384_v32
		})())
		_464_v82 = (_464_v82).Update(_462_v80, Companion_Default___.SafeModuloInt(_384_v32, _458_v77))
		if (func() bool {
			if _346_v4 {
				return true
			}
			return (func() bool {
				if (_461_v79).Fm2((_461_v79).F19(), _346_v4, _352_globalState) {
					return (_461_v79).F19()
				}
				return _346_v4
			})()
		})() {
			var _465_v83 _dafny.Array
			_ = _465_v83
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_5
			var _nw63 _dafny.Array
			_ = _nw63
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw63 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_466_v1 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
					return func(_467_i8 _dafny.Int) _dafny.Int {
						return (_467_i8).Times((_466_v1).Cardinality())
					}
				})(_343_v1)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw63 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw63).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw63).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_465_v83 = _nw63
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_465_v83), 0))
			_ = _index50
			(_465_v83).ArraySet1(_458_v77, (_index50).Int())
			var _468_v84 D4
			_ = _468_v84
			_468_v84 = Companion_D4_.Create_DC12_(_344_v2)
			var _469_v85 D4
			_ = _469_v85
			_469_v85 = Companion_D4_.Create_DC15_(_468_v84)
			var _470_v86 *C6
			_ = _470_v86
			var _nw64 *C6 = New_C6_()
			_ = _nw64
			_nw64.Ctor__(_346_v4, _463_v81)
			_470_v86 = _nw64
			var _471_v87 _dafny.Sequence
			_ = _471_v87
			_471_v87 = _dafny.SeqOf(_470_v86)
			var _472_v88 _dafny.Map
			_ = _472_v88
			_472_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_471_v87, Companion_Default___.Fm6(_352_globalState))
			var _pat_let_tv18 = _468_v84
			_ = _pat_let_tv18
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_465_v83), 0))
			_ = _index51
			var _rhs47 _dafny.Int = (Companion_Default___.SafeDivisionInt(_342_v0, _dafny.IntOfInt64(632))).Times((_472_v88).Cardinality())
			_ = _rhs47
			var _rhs48 _dafny.Int = _384_v32
			_ = _rhs48
			var _rhs49 _dafny.Int = _384_v32
			_ = _rhs49
			var _rhs50 D4 = func(_pat_let17_0 D4) D4 {
				return func(_473_dt__update__tmp_h6 D4) D4 {
					return func(_pat_let18_0 D4) D4 {
						return func(_474_dt__update_hcf28_h0 D4) D4 {
							return Companion_D4_.Create_DC15_(_474_dt__update_hcf28_h0)
						}(_pat_let18_0)
					}(_pat_let_tv18)
				}(_pat_let17_0)
			}(_469_v85)
			_ = _rhs50
			var _lhs42 _dafny.Array = _465_v83
			_ = _lhs42
			var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_465_v83), 0))
			_ = _lhs43
			var _lhs44 *GlobalState = _352_globalState
			_ = _lhs44
			var _lhs45 *GlobalState = _352_globalState
			_ = _lhs45
			(_lhs42).ArraySet1(_rhs47, (_lhs43).Int())
			_lhs44.F3 = _rhs48
			_lhs45.F3 = _rhs49
			_469_v85 = _rhs50
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_465_v83), 0))
			_ = _index52
			(_465_v83).ArraySet1((_465_v83).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_465_v83), 0))).Int()).(_dafny.Int), (_index52).Int())
			var _475_v89 _dafny.Int
			_ = _475_v89
			var _476_v90 _dafny.Int
			_ = _476_v90
			var _477_v91 _dafny.MultiSet
			_ = _477_v91
			var _478_v92 bool
			_ = _478_v92
			var _out13 _dafny.Int
			_ = _out13
			var _out14 _dafny.Int
			_ = _out14
			var _out15 _dafny.MultiSet
			_ = _out15
			var _out16 bool
			_ = _out16
			_out13, _out14, _out15, _out16 = (_470_v86).M5(_352_globalState)
			_475_v89 = _out13
			_476_v90 = _out14
			_477_v91 = _out15
			_478_v92 = _out16
			var _479_v93 _dafny.Int
			_ = _479_v93
			var _480_v94 bool
			_ = _480_v94
			var _481_v95 _dafny.Map
			_ = _481_v95
			var _482_v96 bool
			_ = _482_v96
			var _out17 _dafny.Int
			_ = _out17
			var _out18 bool
			_ = _out18
			var _out19 _dafny.Map
			_ = _out19
			var _out20 bool
			_ = _out20
			_out17, _out18, _out19, _out20 = (_461_v79).M0(_352_globalState)
			_479_v93 = _out17
			_480_v94 = _out18
			_481_v95 = _out19
			_482_v96 = _out20
			var _483_v97 *C7
			_ = _483_v97
			var _nw65 *C7 = New_C7_()
			_ = _nw65
			_nw65.Ctor__()
			_483_v97 = _nw65
		} else {
			var _484_v98 _dafny.Array
			_ = _484_v98
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_6
			var _nw66 _dafny.Array
			_ = _nw66
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw66 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.CodePoint = (func(_485_v31 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_486_i9 _dafny.Int) _dafny.CodePoint {
						return _485_v31
					}
				})(_383_v31)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw66 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw66).ArraySet1CodePoint(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw66).ArraySet1CodePoint(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_484_v98 = _nw66
			var _487_v99 _dafny.Map
			_ = _487_v99
			_487_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _484_v98)
			var _488_v100 _dafny.Map
			_ = _488_v100
			_488_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_461_v79).F19(), _384_v32)
			var _489_v101 _dafny.Array
			_ = _489_v101
			var _nwElement0_10 _dafny.Array = _484_v98
			_ = _nwElement0_10
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(15))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_10, 0)
			(_nw67).ArraySet1(_484_v98, 1)
			(_nw67).ArraySet1(_484_v98, 2)
			(_nw67).ArraySet1(_484_v98, 3)
			(_nw67).ArraySet1((func() _dafny.Array {
				if (_487_v99).Contains(Companion_Default___.Fm25(_342_v0, _dafny.IntOfUint32((_355_v10).Cardinality()), (_461_v79).F19(), _488_v100, _352_globalState)) {
					return (_487_v99).Get(Companion_Default___.Fm25(_342_v0, _dafny.IntOfUint32((_355_v10).Cardinality()), (_461_v79).F19(), _488_v100, _352_globalState)).(_dafny.Array)
				}
				return _484_v98
			})(), 4)
			(_nw67).ArraySet1(_484_v98, 5)
			(_nw67).ArraySet1(_484_v98, 6)
			(_nw67).ArraySet1(_484_v98, 7)
			(_nw67).ArraySet1(_484_v98, 8)
			(_nw67).ArraySet1(_484_v98, 9)
			(_nw67).ArraySet1(_484_v98, 10)
			(_nw67).ArraySet1((func() _dafny.Array {
				if (_461_v79).F19() {
					return _484_v98
				}
				return _484_v98
			})(), 11)
			(_nw67).ArraySet1(_484_v98, 12)
			(_nw67).ArraySet1(_484_v98, 13)
			(_nw67).ArraySet1(_484_v98, 14)
			_489_v101 = _nw67
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_489_v101), 0))
			_ = _index53
			(_489_v101).ArraySet1(_484_v98, (_index53).Int())
			var _490_v102 D16
			_ = _490_v102
			_490_v102 = Companion_D16_.Create_DC51_(_484_v98)
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(670), _dafny.ArrayLen((_489_v101), 0))
			_ = _index54
			(_489_v101).ArraySet1((_490_v102).Dtor_cf83(), (_index54).Int())
			var _491_v103 _dafny.Array
			_ = _491_v103
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
			_ = _nw68
			_491_v103 = _nw68
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_491_v103), 0))
			_ = _index55
			(_491_v103).ArraySet1((func() _dafny.Int {
				if (_461_v79).F19() {
					return _458_v77
				}
				return _342_v0
			})(), (_index55).Int())
			var _492_v104 _dafny.Map
			_ = _492_v104
			_492_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_491_v103, (func() bool {
				if _346_v4 {
					return (_353_v8).Select((Companion_Default___.SafeIndex((_382_v30).Cardinality(), _dafny.IntOfUint32((_353_v8).Cardinality()))).Uint32()).(bool)
				}
				return _346_v4
			})())
			var _493_v105 D5
			_ = _493_v105
			_493_v105 = Companion_D5_.Create_DC18_((_461_v79).F19())
			var _494_v106 _dafny.Map
			_ = _494_v106
			_494_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, Companion_D11_.Create_DC35_())
			var _495_v107 _dafny.Set
			_ = _495_v107
			_495_v107 = _dafny.SetOf(_345_v3)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_491_v103), 0))
			_ = _index56
			var _rhs51 bool = (_461_v79).F19()
			_ = _rhs51
			var _rhs52 bool = (_493_v105).Dtor_cf33()
			_ = _rhs52
			var _rhs53 bool = (func() bool {
				if !(_494_v106).Equals(_494_v106) {
					return (_461_v79).F19()
				}
				return (_495_v107).IsDisjointFrom(_495_v107)
			})()
			_ = _rhs53
			var _rhs54 _dafny.Int = Companion_Default___.SafeModuloInt(_342_v0, _342_v0)
			_ = _rhs54
			var _rhs55 _dafny.Map = _492_v104
			_ = _rhs55
			var _lhs46 *GlobalState = _352_globalState
			_ = _lhs46
			var _lhs47 *GlobalState = _352_globalState
			_ = _lhs47
			var _lhs48 _dafny.Array = _491_v103
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_491_v103), 0))
			_ = _lhs49
			_lhs46.F9 = _rhs51
			_lhs47.F9 = _rhs52
			_346_v4 = _rhs53
			(_lhs48).ArraySet1(_rhs54, (_lhs49).Int())
			_492_v104 = _rhs55
			var _496_v108 *C3
			_ = _496_v108
			var _nw69 *C3 = New_C3_()
			_ = _nw69
			_nw69.Ctor__((_491_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_491_v103), 0))).Int()).(_dafny.Int), _345_v3, _346_v4, _463_v81)
			_496_v108 = _nw69
			var _497_v109 D18
			_ = _497_v109
			_497_v109 = Companion_D18_.Create_DC55_(_496_v108)
			var _498_v110 _dafny.MultiSet
			_ = _498_v110
			_498_v110 = _dafny.MultiSetOf((_497_v109).Dtor_cf91(), _496_v108)
			var _499_v111 _dafny.Sequence
			_ = _499_v111
			_499_v111 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_498_v110).Cardinality(), (_496_v108).F20())).Update(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_353_v8, (Companion_Default___.SafeIndex((_491_v103).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_491_v103), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_353_v8).Cardinality()))).Uint32(), (_461_v79).F19())).Cardinality()), (_dafny.Zero).Minus((_496_v108).F20())))
			_499_v111 = _499_v111
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_344_v2), 0))
			_ = _index57
			(_344_v2).ArraySet1(true, (_index57).Int())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_344_v2), 0))
			_ = _index58
			(_344_v2).ArraySet1((_461_v79).F19(), (_index58).Int())
			var _500_v112 _dafny.Map
			_ = _500_v112
			_500_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, _383_v31)
			var _501_v113 _dafny.MultiSet
			_ = _501_v113
			_501_v113 = _dafny.MultiSetOf((_500_v112).Update((_461_v79).F19(), _383_v31), _500_v112)
			(_352_globalState).F9 = (_501_v113).IsProperSubsetOf(_501_v113)
		}
		var _502_v114 _dafny.Map
		_ = _502_v114
		_502_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_344_v2, _383_v31)
		var _503_v115 D15
		_ = _503_v115
		_503_v115 = Companion_D15_.Create_DC49_((_502_v114).Merge(_502_v114), _dafny.IntOfInt64(-88))
		_503_v115 = Companion_D15_.Create_DC49_(_502_v114, _384_v32)
	}
	_342_v0 = _384_v32
	_346_v4 = _346_v4
	var _504_v116 _dafny.MultiSet
	_ = _504_v116
	_504_v116 = _dafny.MultiSetOf(_346_v4, _346_v4, _346_v4, _346_v4, _346_v4)
	var _505_v117 *C6
	_ = _505_v117
	var _nw70 *C6 = New_C6_()
	_ = _nw70
	_nw70.Ctor__(_346_v4, _504_v116)
	_505_v117 = _nw70
	if !_dafny.Companion_Sequence_.Equal(_355_v10, _355_v10) {
		_342_v0 = Companion_Default___.Fm0(_346_v4, ((_dafny.Zero).Minus(_384_v32)).Plus(_dafny.IntOfInt64(359)), _384_v32, _343_v1, _352_globalState)
		var _506_v118 D8
		_ = _506_v118
		_506_v118 = Companion_D8_.Create_DC24_(_dafny.SetOf(_346_v4))
		var _507_v119 _dafny.Map
		_ = _507_v119
		_507_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_384_v32, _347_v5)
		var _508_v120 _dafny.Map
		_ = _508_v120
		_508_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_342_v0, _346_v4)
		var _509_v121 _dafny.Sequence
		_ = _509_v121
		_509_v121 = _dafny.SeqOf(_347_v5, _dafny.SetOf(_346_v4, _346_v4, true), _dafny.SetOf(_346_v4))
		var _510_v122 _dafny.Array
		_ = _510_v122
		var _nwElement0_11 _dafny.Set = _347_v5
		_ = _nwElement0_11
		var _nw71 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(16))
		_ = _nw71
		(_nw71).ArraySet1(_nwElement0_11, 0)
		(_nw71).ArraySet1(_347_v5, 1)
		(_nw71).ArraySet1(_347_v5, 2)
		(_nw71).ArraySet1((_506_v118).Dtor_cf45(), 3)
		(_nw71).ArraySet1(_347_v5, 4)
		(_nw71).ArraySet1((func() _dafny.Set {
			if (_507_v119).Contains(_342_v0) {
				return (_507_v119).Get(_342_v0).(_dafny.Set)
			}
			return _347_v5
		})(), 5)
		(_nw71).ArraySet1((_347_v5).Union(Companion_Default___.Fm41(_384_v32, (_508_v120).Cardinality(), _352_globalState)), 6)
		(_nw71).ArraySet1(((_509_v121).Select((Companion_Default___.SafeIndex(_384_v32, _dafny.IntOfUint32((_509_v121).Cardinality()))).Uint32()).(_dafny.Set)).Intersection(_dafny.SetOf(_346_v4)), 7)
		(_nw71).ArraySet1((func() _dafny.Set {
			if _346_v4 {
				return _347_v5
			}
			return _347_v5
		})(), 8)
		(_nw71).ArraySet1(Companion_Default___.Fm41(_dafny.IntOfUint32((_355_v10).Cardinality()), _342_v0, _352_globalState), 9)
		(_nw71).ArraySet1(_dafny.SetOf(_346_v4, _346_v4), 10)
		(_nw71).ArraySet1(_347_v5, 11)
		(_nw71).ArraySet1((_347_v5).Difference(_dafny.SetOf(_346_v4, _346_v4, _346_v4, _346_v4, _346_v4)), 12)
		(_nw71).ArraySet1((_347_v5).Union(_347_v5), 13)
		(_nw71).ArraySet1(_347_v5, 14)
		(_nw71).ArraySet1(_347_v5, 15)
		_510_v122 = _nw71
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_510_v122), 0))
		_ = _index59
		(_510_v122).ArraySet1((_347_v5).Intersection(_347_v5), (_index59).Int())
		var _511_v123 _dafny.Array
		_ = _511_v123
		var _nwElement0_12 _dafny.Int = (_dafny.Zero).Minus(_342_v0)
		_ = _nwElement0_12
		var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(10))
		_ = _nw72
		(_nw72).ArraySet1(_nwElement0_12, 0)
		(_nw72).ArraySet1((_355_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_345_v3).Cardinality()), _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32()).(_dafny.Int), 1)
		(_nw72).ArraySet1(_dafny.IntOfUint32((_353_v8).Cardinality()), 2)
		(_nw72).ArraySet1(Companion_Default___.SafeDivisionInt(_384_v32, _342_v0), 3)
		(_nw72).ArraySet1(_342_v0, 4)
		(_nw72).ArraySet1((_dafny.Zero).Minus((_342_v0).Plus(_342_v0)), 5)
		(_nw72).ArraySet1(_384_v32, 6)
		(_nw72).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, _383_v31)).Update(_346_v4, _383_v31)).Cardinality(), 7)
		(_nw72).ArraySet1((_384_v32).Minus(_384_v32), 8)
		(_nw72).ArraySet1(_342_v0, 9)
		_511_v123 = _nw72
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_511_v123), 0))
		_ = _index60
		(_511_v123).ArraySet1(_342_v0, (_index60).Int())
		var _512_v124 D12
		_ = _512_v124
		_512_v124 = Companion_D12_.Create_DC39_(_384_v32)
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_510_v122), 0))
		_ = _index61
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_511_v123), 0))
		_ = _index62
		var _rhs56 _dafny.Set = Companion_Default___.Fm41(_dafny.IntOfUint32((_345_v3).Cardinality()), (_512_v124).Dtor_cf65(), _352_globalState)
		_ = _rhs56
		var _rhs57 _dafny.Set = (_347_v5).Intersection(_347_v5)
		_ = _rhs57
		var _rhs58 _dafny.Int = (_dafny.Zero).Minus((_342_v0).Plus(_dafny.IntOfUint32((_345_v3).Cardinality())))
		_ = _rhs58
		var _rhs59 _dafny.Int = Companion_Default___.Fm0(_346_v4, _342_v0, _dafny.IntOfInt64(-571), _343_v1, _352_globalState)
		_ = _rhs59
		var _lhs50 _dafny.Array = _510_v122
		_ = _lhs50
		var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen((_510_v122), 0))
		_ = _lhs51
		var _lhs52 *GlobalState = _352_globalState
		_ = _lhs52
		var _lhs53 _dafny.Array = _511_v123
		_ = _lhs53
		var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_511_v123), 0))
		_ = _lhs54
		(_lhs50).ArraySet1(_rhs56, (_lhs51).Int())
		_lhs52.F6 = _rhs57
		(_lhs53).ArraySet1(_rhs58, (_lhs54).Int())
		_384_v32 = _rhs59
		_346_v4 = !(_346_v4)
		_346_v4 = true
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_511_v123), 0))
		_ = _index63
		(_511_v123).ArraySet1(_342_v0, (_index63).Int())
	} else {
		if _346_v4 {
			_504_v116 = _504_v116
			var _513_v125 _dafny.Set
			_ = _513_v125
			_513_v125 = _dafny.SetOf(_dafny.IntOfInt64(78), _342_v0)
			var _514_v126 _dafny.Map
			_ = _514_v126
			_514_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_383_v31, _513_v125)
			var _515_v127 _dafny.Int
			_ = _515_v127
			var _516_v128 _dafny.Int
			_ = _516_v128
			var _517_v129 bool
			_ = _517_v129
			var _out21 _dafny.Int
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 bool
			_ = _out23
			_out21, _out22, _out23 = (_505_v117).M1(((func() _dafny.Set {
				if (_514_v126).Contains(_383_v31) {
					return (_514_v126).Get(_383_v31).(_dafny.Set)
				}
				return _513_v125
			})()).Union(_dafny.SetOf(_dafny.IntOfInt64(-222))), _dafny.IntOfInt64(30), _352_globalState)
			_515_v127 = _out21
			_516_v128 = _out22
			_517_v129 = _out23
			var _518_v130 *C3
			_ = _518_v130
			var _nw73 *C3 = New_C3_()
			_ = _nw73
			_nw73.Ctor__(_516_v128, _345_v3, true, _504_v116)
			_518_v130 = _nw73
			var _519_v131 _dafny.Map
			_ = _519_v131
			_519_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_516_v128, _516_v128), _518_v130)
			_519_v131 = (_519_v131).Update(_384_v32, _518_v130)
			var _520_v132 _dafny.Map
			_ = _520_v132
			_520_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_516_v128, _344_v2)
			var _521_v133 *C2
			_ = _521_v133
			var _nw74 *C2 = New_C2_()
			_ = _nw74
			_nw74.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_520_v132).Cardinality(), _346_v4)).Update(_384_v32, _346_v4))
			_521_v133 = _nw74
			var _522_v134 _dafny.Map
			_ = _522_v134
			_522_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_517_v129, _342_v0)
			_522_v134 = (_522_v134).Update(_517_v129, _342_v0)
		} else {
			var _523_v135 _dafny.Int
			_ = _523_v135
			var _524_v136 _dafny.Int
			_ = _524_v136
			var _525_v137 _dafny.MultiSet
			_ = _525_v137
			var _526_v138 bool
			_ = _526_v138
			var _out24 _dafny.Int
			_ = _out24
			var _out25 _dafny.Int
			_ = _out25
			var _out26 _dafny.MultiSet
			_ = _out26
			var _out27 bool
			_ = _out27
			_out24, _out25, _out26, _out27 = (_505_v117).M5(_352_globalState)
			_523_v135 = _out24
			_524_v136 = _out25
			_525_v137 = _out26
			_526_v138 = _out27
			var _527_v139 *C7
			_ = _527_v139
			var _nw75 *C7 = New_C7_()
			_ = _nw75
			_nw75.Ctor__()
			_527_v139 = _nw75
			var _nw76 *C7 = New_C7_()
			_ = _nw76
			_nw76.Ctor__()
			_527_v139 = _nw76
			var _528_v140 _dafny.Map
			_ = _528_v140
			_528_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, _346_v4)).Cardinality(), _526_v138)
			var _529_v141 _dafny.Array
			_ = _529_v141
			var _nw77 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw77
			_529_v141 = _nw77
			var _530_v142 D16
			_ = _530_v142
			_530_v142 = Companion_D16_.Create_DC52_(_342_v0, _526_v138, (Companion_D16_.Create_DC52_(_524_v136, _526_v138, _529_v141, true, _342_v0)).Dtor_cf86(), _346_v4, _524_v136)
			_528_v140 = (_528_v140).Update(((_530_v142).Dtor_cf88()).Times(_384_v32), _346_v4)
			var _531_v143 _dafny.Array
			_ = _531_v143
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_7
			var _nw78 _dafny.Array
			_ = _nw78
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw78 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Int = (func(_532_v136 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_533_i10 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_533_i10, _532_v136)
					}
				})(_524_v136)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw78 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw78).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw78).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_531_v143 = _nw78
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_531_v143), 0))
			_ = _index64
			(_531_v143).ArraySet1((_355_v10).Select((Companion_Default___.SafeIndex(_342_v0, _dafny.IntOfUint32((_355_v10).Cardinality()))).Uint32()).(_dafny.Int), (_index64).Int())
			var _534_v144 _dafny.Map
			_ = _534_v144
			_534_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_526_v138, _384_v32)
			var _535_v145 _dafny.Set
			_ = _535_v145
			_535_v145 = _dafny.SetOf(_384_v32)
			var _536_v147 T0
			_ = _536_v147
			var _nw79 *C5 = New_C5_()
			_ = _nw79
			_nw79.Ctor__(_524_v136, _526_v138, _504_v116)
			_536_v147 = _nw79
			var _537_v148 D19
			_ = _537_v148
			_537_v148 = Companion_D19_.Create_DC59_(_536_v147, _534_v144)
			var _538_v149 _dafny.Array
			_ = _538_v149
			var _nwElement0_13 _dafny.Map = (_534_v144).Update(_526_v138, _523_v135)
			_ = _nwElement0_13
			var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(29))
			_ = _nw80
			(_nw80).ArraySet1(_nwElement0_13, 0)
			(_nw80).ArraySet1(_534_v144, 1)
			(_nw80).ArraySet1((_534_v144).Merge(_534_v144), 2)
			(_nw80).ArraySet1(Companion_Default___.Fm1(_526_v138, _384_v32, _346_v4, _535_v145, _352_globalState), 3)
			(_nw80).ArraySet1(_534_v144, 4)
			(_nw80).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, _384_v32), 5)
			(_nw80).ArraySet1(_534_v144, 6)
			(_nw80).ArraySet1((func() _dafny.Map {
				if !(_346_v4) {
					return Companion_Default___.Fm1(_346_v4, _384_v32, !(_526_v138), func() _dafny.Set {
						var _coll37 = _dafny.NewBuilder()
						_ = _coll37
						for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(994), _dafny.IntOfInt64(775))); ; {
							_compr_37, _ok37 := _iter37()
							if !_ok37 {
								break
							}
							var _539_v146 _dafny.Int
							_539_v146 = interface{}(_compr_37).(_dafny.Int)
							if ((_dafny.IntOfInt64(994)).Cmp(_539_v146) <= 0) && ((_539_v146).Cmp(_dafny.IntOfInt64(775)) < 0) {
								_coll37.Add(Companion_Default___.SafeDivisionInt(_539_v146, _384_v32))
							}
						}
						return _coll37.ToSet()
					}(), _352_globalState)
				}
				return _534_v144
			})(), 7)
			(_nw80).ArraySet1(_534_v144, 8)
			(_nw80).ArraySet1(_534_v144, 9)
			(_nw80).ArraySet1(Companion_Default___.Fm1(_346_v4, _dafny.IntOfInt64(705), _346_v4, _535_v145, _352_globalState), 10)
			(_nw80).ArraySet1((_537_v148).Dtor_cf98(), 11)
			(_nw80).ArraySet1(_534_v144, 12)
			(_nw80).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_526_v138, _523_v135)).Merge(_534_v144), 13)
			(_nw80).ArraySet1(_534_v144, 14)
			(_nw80).ArraySet1(_534_v144, 15)
			(_nw80).ArraySet1((_534_v144).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_526_v138, _384_v32)), 16)
			(_nw80).ArraySet1((Companion_Default___.Fm1(!(true), _524_v136, _526_v138, _535_v145, _352_globalState)).Merge(_534_v144), 17)
			(_nw80).ArraySet1((_534_v144).Merge((_534_v144).Update(_346_v4, _384_v32)), 18)
			(_nw80).ArraySet1((Companion_D19_.Create_DC59_(_536_v147, _534_v144)).Dtor_cf98(), 19)
			(_nw80).ArraySet1(_534_v144, 20)
			(_nw80).ArraySet1((_534_v144).Merge(_534_v144), 21)
			(_nw80).ArraySet1(_534_v144, 22)
			(_nw80).ArraySet1(_534_v144, 23)
			(_nw80).ArraySet1(_534_v144, 24)
			(_nw80).ArraySet1(_534_v144, 25)
			(_nw80).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _342_v0), 26)
			(_nw80).ArraySet1((_534_v144).Merge(Companion_Default___.Fm1(_526_v138, _523_v135, _526_v138, _535_v145, _352_globalState)), 27)
			(_nw80).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_536_v147.F14(), _342_v0), 28)
			_538_v149 = _nw80
			var _540_v150 _dafny.Sequence
			_ = _540_v150
			_540_v150 = _dafny.SeqOf(_534_v144)
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_538_v149), 0))
			_ = _index65
			(_538_v149).ArraySet1(((_540_v150).Select((Companion_Default___.SafeIndex(_523_v135, _dafny.IntOfUint32((_540_v150).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_536_v147.F14(), _524_v136)), (_index65).Int())
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_531_v143), 0))
			_ = _index66
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_538_v149), 0))
			_ = _index67
			var _rhs60 _dafny.CodePoint = _383_v31
			_ = _rhs60
			var _rhs61 _dafny.Int = _384_v32
			_ = _rhs61
			var _rhs62 _dafny.Map = ((_534_v144).Merge(_534_v144)).Merge(((_534_v144).Update(_346_v4, _523_v135)).Merge(_534_v144))
			_ = _rhs62
			var _rhs63 _dafny.Int = _342_v0
			_ = _rhs63
			var _rhs64 _dafny.CodePoint = _383_v31
			_ = _rhs64
			var _lhs55 _dafny.Array = _531_v143
			_ = _lhs55
			var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_531_v143), 0))
			_ = _lhs56
			var _lhs57 _dafny.Array = _538_v149
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(514), _dafny.ArrayLen((_538_v149), 0))
			_ = _lhs58
			_383_v31 = _rhs60
			(_lhs55).ArraySet1(_rhs61, (_lhs56).Int())
			(_lhs57).ArraySet1(_rhs62, (_lhs58).Int())
			_524_v136 = _rhs63
			_383_v31 = _rhs64
			_342_v0 = (_524_v136).Minus(Companion_Default___.SafeDivisionInt(_523_v135, _dafny.IntOfUint32((_345_v3).Cardinality())))
		}
		var _541_v151 _dafny.Array
		_ = _541_v151
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
		_ = _nw81
		_541_v151 = _nw81
		_541_v151 = _541_v151
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_541_v151), 0))
		_ = _index68
		(_541_v151).ArraySet1(_dafny.IntOfInt64(815), (_index68).Int())
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_541_v151), 0))
		_ = _index69
		(_541_v151).ArraySet1((func() _dafny.Int {
			if true {
				return _384_v32
			}
			return _384_v32
		})(), (_index69).Int())
		_345_v3 = _345_v3
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_541_v151), 0))
		_ = _index70
		(_541_v151).ArraySet1((_541_v151).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_541_v151), 0))).Int()).(_dafny.Int), (_index70).Int())
	}
	var _542_i11 _dafny.Int
	_ = _542_i11
	_542_i11 = _dafny.Zero
	{
		for _346_v4 {
			{
				if (_542_i11).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_542_i11 = (_542_i11).Plus(_dafny.One)
				var _543_v152 _dafny.Array
				_ = _543_v152
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_8
				var _nw82 _dafny.Array
				_ = _nw82
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw82 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Int = (func(_544_v4 bool) func(_dafny.Int) _dafny.Int {
						return func(_545_i12 _dafny.Int) _dafny.Int {
							return (_545_i12).Times((_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v4, _544_v4))))).Cardinality())
						}
					})(_346_v4)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw82 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw82).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw82).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_543_v152 = _nw82
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_543_v152), 0))
				_ = _index71
				(_543_v152).ArraySet1(_384_v32, (_index71).Int())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.ArrayLen((_543_v152), 0))
				_ = _index72
				(_543_v152).ArraySet1(Companion_Default___.SafeModuloInt((_343_v1).Cardinality(), _384_v32), (_index72).Int())
				var _546_v153 _dafny.Array
				_ = _546_v153
				var _nw83 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
				_ = _nw83
				_546_v153 = _nw83
				var _547_v154 *C3
				_ = _547_v154
				var _nw84 *C3 = New_C3_()
				_ = _nw84
				_nw84.Ctor__(_342_v0, _dafny.UnicodeSeqOfUtf8Bytes("ihvjxhpxy"), _346_v4, _504_v116)
				_547_v154 = _nw84
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_546_v153), 0))
				_ = _index73
				(_546_v153).ArraySet1(_547_v154, (_index73).Int())
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(945), _dafny.ArrayLen((_546_v153), 0))
				_ = _index74
				var _nw85 *C3 = New_C3_()
				_ = _nw85
				_nw85.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(true, _346_v4, _346_v4, _346_v4)).Cardinality()), _dafny.Companion_Sequence_.Concatenate((_547_v154).F21(), (_547_v154).F21()), _346_v4, _504_v116)
				(_546_v153).ArraySet1(_nw85, (_index74).Int())
				var _548_v155 D2
				_ = _548_v155
				_548_v155 = Companion_D2_.Create_DC7_(_353_v8)
				var _549_v156 _dafny.Set
				_ = _549_v156
				_549_v156 = _dafny.SetOf(_548_v155, _548_v155, _548_v155)
				(_352_globalState).F9 = !(_549_v156).Contains(Companion_D2_.Create_DC7_(_353_v8))
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_344_v2), 0))
				_ = _index75
				(_344_v2).ArraySet1(_346_v4, (_index75).Int())
				var _550_v157 _dafny.Map
				_ = _550_v157
				_550_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_346_v4, ((_547_v154).F21()).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-660), _dafny.IntOfUint32(((_547_v154).F21()).Cardinality()))).Uint32()).(_dafny.CodePoint))
				var _551_v158 _dafny.Set
				_ = _551_v158
				_551_v158 = _dafny.SetOf((_dafny.Zero).Minus((_550_v157).Cardinality()))
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_344_v2), 0))
				_ = _index76
				(_344_v2).ArraySet1((Companion_Default___.Fm6(_352_globalState)).IsSubsetOf(_551_v158), (_index76).Int())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _552_v159 _dafny.Array
	_ = _552_v159
	var _len0_9 _dafny.Int = _dafny.IntOfInt64(13)
	_ = _len0_9
	var _nw86 _dafny.Array
	_ = _nw86
	if _len0_9.Cmp(_dafny.Zero) == 0 {
		_nw86 = _dafny.NewArray(_len0_9)
	} else {
		var _init9 func(_dafny.Int) _dafny.Int = (func(_553_v8 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
			return func(_554_i14 _dafny.Int) _dafny.Int {
				return (_554_i14).Minus(_dafny.IntOfUint32((_553_v8).Cardinality()))
			}
		})(_353_v8)
		_ = _init9
		var _element0_9 = _init9(_dafny.Zero)
		_ = _element0_9
		_nw86 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
		(_nw86).ArraySet1(_element0_9, 0)
		var _nativeLen0_9 = (_len0_9).Int()
		_ = _nativeLen0_9
		for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
			(_nw86).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
		}
	}
	_552_v159 = _nw86
	for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_552_v159), 0))); ; {
		_guard_loop_0, _ok38 := _iter38()
		if !_ok38 {
			break
		}
		var _555_i13 _dafny.Int
		_555_i13 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_555_i13).Sign() != -1) && ((_555_i13).Cmp(_dafny.ArrayLen((_552_v159), 0)) < 0)) {
			(_552_v159).ArraySet1((_555_i13).Minus(_dafny.IntOfInt64(525)), (_555_i13).Int())
		}
	}
	var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
	_ = _index77
	(_344_v2).ArraySet1(_346_v4, (_index77).Int())
	var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
	_ = _index78
	(_344_v2).ArraySet1(_346_v4, (_index78).Int())
	var _556_v160 _dafny.Int
	_ = _556_v160
	var _557_v161 bool
	_ = _557_v161
	var _558_v162 _dafny.Map
	_ = _558_v162
	var _559_v163 bool
	_ = _559_v163
	var _out28 _dafny.Int
	_ = _out28
	var _out29 bool
	_ = _out29
	var _out30 _dafny.Map
	_ = _out30
	var _out31 bool
	_ = _out31
	_out28, _out29, _out30, _out31 = (_505_v117).M0(_352_globalState)
	_556_v160 = _out28
	_557_v161 = _out29
	_558_v162 = _out30
	_559_v163 = _out31
	var _560_v164 _dafny.Map
	_ = _560_v164
	_560_v164 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v160, _346_v4)
	_560_v164 = (_560_v164).Update((_dafny.Zero).Minus(_342_v0), (_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool))
	var _561_i15 _dafny.Int
	_ = _561_i15
	_561_i15 = _dafny.Zero
	{
		for (_dafny.IntOfUint32((_355_v10).Cardinality())).Cmp(Companion_Default___.SafeDivisionInt(_342_v0, _342_v0)) <= 0 {
			{
				if (_561_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_561_i15 = (_561_i15).Plus(_dafny.One)
				var _562_v165 _dafny.MultiSet
				_ = _562_v165
				_562_v165 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(76))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg49 _dafny.Int) interface{} {
						return coer49(arg49)
					}
				}(func(_563_i16 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(188)
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(731))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_564_v32 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_565_i17 _dafny.Int) _dafny.Int {
						return _564_v32
					}
				})(_384_v32))), _355_v10)
				_562_v165 = _562_v165
				var _566_v166 D1
				_ = _566_v166
				_566_v166 = Companion_D1_.Create_DC5_(_346_v4, _557_v161)
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
				_ = _index79
				var _rhs65 bool = (func() bool {
					if (_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool) {
						return (_566_v166).Dtor_cf6()
					}
					return _557_v161
				})()
				_ = _rhs65
				var _rhs66 bool = _559_v163
				_ = _rhs66
				var _rhs67 _dafny.Sequence = (Companion_D4_.Create_DC13_(_345_v3, _556_v160, _557_v161, _344_v2)).Dtor_cf22()
				_ = _rhs67
				var _lhs59 _dafny.Array = _344_v2
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
				_ = _lhs60
				_559_v163 = _rhs65
				(_lhs59).ArraySet1(_rhs66, (_lhs60).Int())
				_345_v3 = _rhs67
				var _567_v167 _dafny.Int
				_ = _567_v167
				var _568_v168 _dafny.Map
				_ = _568_v168
				var _out32 _dafny.Int
				_ = _out32
				var _out33 _dafny.Map
				_ = _out33
				_out32, _out33 = Companion_Default___.M11(_381_v29, _dafny.IntOfUint32((_345_v3).Cardinality()), true, (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(307), _384_v32))).IsProperSubsetOf(_343_v1), _352_globalState)
				_567_v167 = _out32
				_568_v168 = _out33
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
				_ = _index80
				(_344_v2).ArraySet1(false, (_index80).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	if (_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool) {
		var _569_v169 _dafny.Set
		_ = _569_v169
		_569_v169 = _dafny.SetOf(_dafny.IntOfInt64(600))
		var _570_v170 _dafny.Int
		_ = _570_v170
		var _571_v171 _dafny.Int
		_ = _571_v171
		var _572_v172 bool
		_ = _572_v172
		var _out34 _dafny.Int
		_ = _out34
		var _out35 _dafny.Int
		_ = _out35
		var _out36 bool
		_ = _out36
		_out34, _out35, _out36 = (_505_v117).M1(_569_v169, (_560_v164).Cardinality(), _352_globalState)
		_570_v170 = _out34
		_571_v171 = _out35
		_572_v172 = _out36
		var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
		_ = _index81
		(_344_v2).ArraySet1(_559_v163, (_index81).Int())
		var _573_v173 *C1
		_ = _573_v173
		var _nw87 *C1 = New_C1_()
		_ = _nw87
		_nw87.Ctor__()
		_573_v173 = _nw87
		var _574_v174 T0
		_ = _574_v174
		var _nw88 *C5 = New_C5_()
		_ = _nw88
		_nw88.Ctor__(_556_v160, !(_557_v161) || (_557_v161), _dafny.MultiSetOf(_346_v4))
		_574_v174 = _nw88
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
		_ = _index82
		var _rhs68 bool = _572_v172
		_ = _rhs68
		var _rhs69 T0 = _574_v174
		_ = _rhs69
		var _lhs61 _dafny.Array = _344_v2
		_ = _lhs61
		var _lhs62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
		_ = _lhs62
		(_lhs61).ArraySet1(_rhs68, (_lhs62).Int())
		_574_v174 = _rhs69
		_559_v163 = _557_v161
	} else {
		var _575_v175 _dafny.Int
		_ = _575_v175
		var _576_v176 _dafny.Map
		_ = _576_v176
		var _out37 _dafny.Int
		_ = _out37
		var _out38 _dafny.Map
		_ = _out38
		_out37, _out38 = Companion_Default___.M11(_381_v29, Companion_Default___.Fm0((_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool), _342_v0, _342_v0, Companion_Default___.Fm30(_557_v161, _352_globalState), _352_globalState), (_dafny.MultiSetFromSeq(_353_v8)).IsSubsetOf(_504_v116), _559_v163, _352_globalState)
		_575_v175 = _out37
		_576_v176 = _out38
		var _577_v177 _dafny.Map
		_ = _577_v177
		_577_v177 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_560_v164).Cardinality(), _345_v3)
		var _578_v178 D8
		_ = _578_v178
		_578_v178 = Companion_D8_.Create_DC26_(_557_v161, _577_v177)
		var _579_v179 _dafny.Set
		_ = _579_v179
		_579_v179 = _dafny.SetOf(_578_v178, _578_v178, _578_v178, _578_v178)
		var _580_v180 _dafny.Map
		_ = _580_v180
		_580_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_559_v163, !(_557_v161))
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
		_ = _index83
		var _rhs70 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_559_v163), _dafny.SeqOf(!(_559_v163), (func() bool {
			if (_580_v180).Contains((_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool)) {
				return (_580_v180).Get((_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool)).(bool)
			}
			return _557_v161
		})()))
		_ = _rhs70
		var _rhs71 bool = _559_v163
		_ = _rhs71
		var _rhs72 _dafny.Set = _579_v179
		_ = _rhs72
		var _rhs73 _dafny.MultiSet = _343_v1
		_ = _rhs73
		var _rhs74 bool = (_344_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))).Int()).(bool)
		_ = _rhs74
		var _lhs63 _dafny.Array = _344_v2
		_ = _lhs63
		var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(302), _dafny.ArrayLen((_344_v2), 0))
		_ = _lhs64
		_353_v8 = _rhs70
		_559_v163 = _rhs71
		_579_v179 = _rhs72
		_343_v1 = _rhs73
		(_lhs63).ArraySet1(_rhs74, (_lhs64).Int())
		var _581_v181 _dafny.Set
		_ = _581_v181
		_581_v181 = _dafny.SetOf(_342_v0)
		var _582_v182 _dafny.Int
		_ = _582_v182
		var _583_v183 _dafny.Int
		_ = _583_v183
		var _584_v184 bool
		_ = _584_v184
		var _out39 _dafny.Int
		_ = _out39
		var _out40 _dafny.Int
		_ = _out40
		var _out41 bool
		_ = _out41
		_out39, _out40, _out41 = (_505_v117).M1((_dafny.SetOf((_dafny.Zero).Minus(_556_v160), _dafny.IntOfInt64(314))).Union(_581_v181), _384_v32, _352_globalState)
		_582_v182 = _out39
		_583_v183 = _out40
		_584_v184 = _out41
		(_352_globalState).F3 = _556_v160
		var _585_v185 *C0
		_ = _585_v185
		var _nw89 *C0 = New_C0_()
		_ = _nw89
		_nw89.Ctor__(!(_559_v163), !(false))
		_585_v185 = _nw89
	}
	_dafny.Print(_342_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_343_v1).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(206))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v2).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v2).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_344_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_345_v3.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_346_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_347_v5).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_349_v7).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F1()).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState.F2).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState.F2).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState.F2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_352_globalState.F3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState.F6).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_352_globalState.F9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_352_globalState).F12()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_352_globalState.F13).Equals(_dafny.SetOf(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_353_v8, _dafny.SeqOf(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_354_v9).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf(true)).UpdateUnsafe(true, _dafny.SetOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_355_v10, _dafny.SeqOf(_dafny.IntOfInt64(206), _dafny.IntOfInt64(206))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_381_v29).Dtor_cf70()).Dtor_cf70()).Dtor_cf70()).Dtor_cf68())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((((_381_v29).Dtor_cf70()).Dtor_cf70()).Dtor_cf70()).Dtor_cf69())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_382_v30).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_383_v31)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_384_v32)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_385_v33).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_504_v116).Equals(_dafny.MultiSetOf(true, true, true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_542_i11)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_552_v159).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_556_v160)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_557_v161)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_558_v162).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-153), true), _dafny.IntOfInt64(-804))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_559_v163)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_560_v164).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-153539075), true).UpdateUnsafe(_dafny.IntOfInt64(-2), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_561_i15)
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

func (_this D0) Dtor_cf0() bool {
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
			return ok && data1.Cf0 == data2.Cf0
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
	Cf5 _dafny.Int
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf4 _dafny.Int, Cf5 _dafny.Int) D1 {
	return D1{D1_DC4{Cf4, Cf5}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf6 bool
	Cf7 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf6 bool, Cf7 bool) D1 {
	return D1{D1_DC5{Cf6, Cf7}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf3 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf3 _dafny.Int) D1 {
	return D1{D1_DC3{Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC6 struct {
	Cf8 D1
}

func (D1_DC6) isD1() {}

func (CompanionStruct_D1_) Create_DC6_(Cf8 D1) D1 {
	return D1{D1_DC6{Cf8}}
}

func (_this D1) Is_DC6() bool {
	_, ok := _this.Get_().(D1_DC6)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf4
}

func (_this D1) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf5
}

func (_this D1) Dtor_cf6() bool {
	return _this.Get_().(D1_DC5).Cf6
}

func (_this D1) Dtor_cf7() bool {
	return _this.Get_().(D1_DC5).Cf7
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf8() D1 {
	return _this.Get_().(D1_DC6).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC6:
		{
			return "D1.DC6" + "(" + _dafny.String(data.Cf8) + ")"
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
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5.Cmp(data2.Cf5) == 0
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D1_DC6:
		{
			data2, ok := other.Get_().(D1_DC6)
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
	Cf10 _dafny.Int
	Cf11 _dafny.Int
	Cf12 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf10 _dafny.Int, Cf11 _dafny.Int, Cf12 _dafny.Int) D2 {
	return D2{D2_DC8{Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC9 struct {
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 _dafny.CodePoint
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf13 bool, Cf14 _dafny.Int, Cf15 _dafny.CodePoint) D2 {
	return D2{D2_DC9{Cf13, Cf14, Cf15}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

type D2_DC7 struct {
	Cf9 _dafny.Sequence
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf9 _dafny.Sequence) D2 {
	return D2{D2_DC7{Cf9}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC8_(_dafny.Zero, _dafny.Zero, _dafny.Zero)
}

func (_this D2) Dtor_cf10() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC9).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC9).Cf14
}

func (_this D2) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D2_DC9).Cf15
}

func (_this D2) Dtor_cf9() _dafny.Sequence {
	return _this.Get_().(D2_DC7).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf9) + ")"
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
			return ok && data1.Cf10.Cmp(data2.Cf10) == 0 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15 == data2.Cf15
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf9.Equals(data2.Cf9)
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
	Cf17 _dafny.Int
	Cf18 _dafny.Int
	Cf19 bool
	Cf20 _dafny.Int
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf17 _dafny.Int, Cf18 _dafny.Int, Cf19 bool, Cf20 _dafny.Int) D3 {
	return D3{D3_DC11{Cf17, Cf18, Cf19, Cf20}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf16 _dafny.Map
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 _dafny.Map) D3 {
	return D3{D3_DC10{Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(_dafny.Zero, _dafny.Zero, false, _dafny.Zero)
}

func (_this D3) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf17
}

func (_this D3) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf18
}

func (_this D3) Dtor_cf19() bool {
	return _this.Get_().(D3_DC11).Cf19
}

func (_this D3) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf17) + ", " + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ")"
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
			return ok && data1.Cf17.Cmp(data2.Cf17) == 0 && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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
	Cf22 _dafny.Sequence
	Cf23 _dafny.Int
	Cf24 bool
	Cf25 _dafny.Array
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf22 _dafny.Sequence, Cf23 _dafny.Int, Cf24 bool, Cf25 _dafny.Array) D4 {
	return D4{D4_DC13{Cf22, Cf23, Cf24, Cf25}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC14 struct {
	Cf26 bool
	Cf27 bool
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf26 bool, Cf27 bool) D4 {
	return D4{D4_DC14{Cf26, Cf27}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

type D4_DC12 struct {
	Cf21 _dafny.Array
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf21 _dafny.Array) D4 {
	return D4{D4_DC12{Cf21}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC15 struct {
	Cf28 D4
}

func (D4_DC15) isD4() {}

func (CompanionStruct_D4_) Create_DC15_(Cf28 D4) D4 {
	return D4{D4_DC15{Cf28}}
}

func (_this D4) Is_DC15() bool {
	_, ok := _this.Get_().(D4_DC15)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(_dafny.EmptySeq, _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D4) Dtor_cf22() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf22
}

func (_this D4) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D4_DC13).Cf23
}

func (_this D4) Dtor_cf24() bool {
	return _this.Get_().(D4_DC13).Cf24
}

func (_this D4) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D4_DC13).Cf25
}

func (_this D4) Dtor_cf26() bool {
	return _this.Get_().(D4_DC14).Cf26
}

func (_this D4) Dtor_cf27() bool {
	return _this.Get_().(D4_DC14).Cf27
}

func (_this D4) Dtor_cf21() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf21
}

func (_this D4) Dtor_cf28() D4 {
	return _this.Get_().(D4_DC15).Cf28
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + data.Cf22.VerbatimString(true) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D4_DC15:
		{
			return "D4.DC15" + "(" + _dafny.String(data.Cf28) + ")"
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
			return ok && data1.Cf22.Equals(data2.Cf22) && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24 == data2.Cf24 && data1.Cf25 == data2.Cf25
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf21 == data2.Cf21
		}
	case D4_DC15:
		{
			data2, ok := other.Get_().(D4_DC15)
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

type D5_DC17 struct {
	Cf30 _dafny.Sequence
	Cf31 bool
	Cf32 _dafny.MultiSet
}

func (D5_DC17) isD5() {}

func (CompanionStruct_D5_) Create_DC17_(Cf30 _dafny.Sequence, Cf31 bool, Cf32 _dafny.MultiSet) D5 {
	return D5{D5_DC17{Cf30, Cf31, Cf32}}
}

func (_this D5) Is_DC17() bool {
	_, ok := _this.Get_().(D5_DC17)
	return ok
}

type D5_DC18 struct {
	Cf33 bool
}

func (D5_DC18) isD5() {}

func (CompanionStruct_D5_) Create_DC18_(Cf33 bool) D5 {
	return D5{D5_DC18{Cf33}}
}

func (_this D5) Is_DC18() bool {
	_, ok := _this.Get_().(D5_DC18)
	return ok
}

type D5_DC16 struct {
	Cf29 _dafny.Array
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf29 _dafny.Array) D5 {
	return D5{D5_DC16{Cf29}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC19 struct {
	Cf34 D5
}

func (D5_DC19) isD5() {}

func (CompanionStruct_D5_) Create_DC19_(Cf34 D5) D5 {
	return D5{D5_DC19{Cf34}}
}

func (_this D5) Is_DC19() bool {
	_, ok := _this.Get_().(D5_DC19)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC17_(_dafny.EmptySeq, false, _dafny.EmptyMultiSet)
}

func (_this D5) Dtor_cf30() _dafny.Sequence {
	return _this.Get_().(D5_DC17).Cf30
}

func (_this D5) Dtor_cf31() bool {
	return _this.Get_().(D5_DC17).Cf31
}

func (_this D5) Dtor_cf32() _dafny.MultiSet {
	return _this.Get_().(D5_DC17).Cf32
}

func (_this D5) Dtor_cf33() bool {
	return _this.Get_().(D5_DC18).Cf33
}

func (_this D5) Dtor_cf29() _dafny.Array {
	return _this.Get_().(D5_DC16).Cf29
}

func (_this D5) Dtor_cf34() D5 {
	return _this.Get_().(D5_DC19).Cf34
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC17:
		{
			return "D5.DC17" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D5_DC18:
		{
			return "D5.DC18" + "(" + _dafny.String(data.Cf33) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D5_DC19:
		{
			return "D5.DC19" + "(" + _dafny.String(data.Cf34) + ")"
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
			return ok && data1.Cf30.Equals(data2.Cf30) && data1.Cf31 == data2.Cf31 && data1.Cf32.Equals(data2.Cf32)
		}
	case D5_DC18:
		{
			data2, ok := other.Get_().(D5_DC18)
			return ok && data1.Cf33 == data2.Cf33
		}
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf29 == data2.Cf29
		}
	case D5_DC19:
		{
			data2, ok := other.Get_().(D5_DC19)
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

type D6_DC20 struct {
	Cf35 _dafny.Map
}

func (D6_DC20) isD6() {}

func (CompanionStruct_D6_) Create_DC20_(Cf35 _dafny.Map) D6 {
	return D6{D6_DC20{Cf35}}
}

func (_this D6) Is_DC20() bool {
	_, ok := _this.Get_().(D6_DC20)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D6) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D6_DC20).Cf35
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC20:
		{
			return "D6.DC20" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf35.Equals(data2.Cf35)
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
	Cf37 _dafny.Array
	Cf38 _dafny.Sequence
	Cf39 _dafny.Sequence
	Cf40 _dafny.Array
	Cf41 D5
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf37 _dafny.Array, Cf38 _dafny.Sequence, Cf39 _dafny.Sequence, Cf40 _dafny.Array, Cf41 D5) D7 {
	return D7{D7_DC22{Cf37, Cf38, Cf39, Cf40, Cf41}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

type D7_DC23 struct {
	Cf42 bool
	Cf43 _dafny.Int
	Cf44 bool
}

func (D7_DC23) isD7() {}

func (CompanionStruct_D7_) Create_DC23_(Cf42 bool, Cf43 _dafny.Int, Cf44 bool) D7 {
	return D7{D7_DC23{Cf42, Cf43, Cf44}}
}

func (_this D7) Is_DC23() bool {
	_, ok := _this.Get_().(D7_DC23)
	return ok
}

type D7_DC21 struct {
	Cf36 _dafny.Map
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf36 _dafny.Map) D7 {
	return D7{D7_DC21{Cf36}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC22_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq, _dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), Companion_D5_.Default())
}

func (_this D7) Dtor_cf37() _dafny.Array {
	return _this.Get_().(D7_DC22).Cf37
}

func (_this D7) Dtor_cf38() _dafny.Sequence {
	return _this.Get_().(D7_DC22).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D7_DC22).Cf39
}

func (_this D7) Dtor_cf40() _dafny.Array {
	return _this.Get_().(D7_DC22).Cf40
}

func (_this D7) Dtor_cf41() D5 {
	return _this.Get_().(D7_DC22).Cf41
}

func (_this D7) Dtor_cf42() bool {
	return _this.Get_().(D7_DC23).Cf42
}

func (_this D7) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D7_DC23).Cf43
}

func (_this D7) Dtor_cf44() bool {
	return _this.Get_().(D7_DC23).Cf44
}

func (_this D7) Dtor_cf36() _dafny.Map {
	return _this.Get_().(D7_DC21).Cf36
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D7_DC23:
		{
			return "D7.DC23" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf36) + ")"
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
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38.Equals(data2.Cf38) && data1.Cf39.Equals(data2.Cf39) && data1.Cf40 == data2.Cf40 && data1.Cf41.Equals(data2.Cf41)
		}
	case D7_DC23:
		{
			data2, ok := other.Get_().(D7_DC23)
			return ok && data1.Cf42 == data2.Cf42 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44 == data2.Cf44
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf36.Equals(data2.Cf36)
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

type D8_DC25 struct {
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_() D8 {
	return D8{D8_DC25{}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC26 struct {
	Cf46 bool
	Cf47 _dafny.Map
}

func (D8_DC26) isD8() {}

func (CompanionStruct_D8_) Create_DC26_(Cf46 bool, Cf47 _dafny.Map) D8 {
	return D8{D8_DC26{Cf46, Cf47}}
}

func (_this D8) Is_DC26() bool {
	_, ok := _this.Get_().(D8_DC26)
	return ok
}

type D8_DC27 struct {
}

func (D8_DC27) isD8() {}

func (CompanionStruct_D8_) Create_DC27_() D8 {
	return D8{D8_DC27{}}
}

func (_this D8) Is_DC27() bool {
	_, ok := _this.Get_().(D8_DC27)
	return ok
}

type D8_DC24 struct {
	Cf45 _dafny.Set
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf45 _dafny.Set) D8 {
	return D8{D8_DC24{Cf45}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC28 struct {
	Cf48 D8
}

func (D8_DC28) isD8() {}

func (CompanionStruct_D8_) Create_DC28_(Cf48 D8) D8 {
	return D8{D8_DC28{Cf48}}
}

func (_this D8) Is_DC28() bool {
	_, ok := _this.Get_().(D8_DC28)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC25_()
}

func (_this D8) Dtor_cf46() bool {
	return _this.Get_().(D8_DC26).Cf46
}

func (_this D8) Dtor_cf47() _dafny.Map {
	return _this.Get_().(D8_DC26).Cf47
}

func (_this D8) Dtor_cf45() _dafny.Set {
	return _this.Get_().(D8_DC24).Cf45
}

func (_this D8) Dtor_cf48() D8 {
	return _this.Get_().(D8_DC28).Cf48
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC25:
		{
			return "D8.DC25"
		}
	case D8_DC26:
		{
			return "D8.DC26" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D8_DC27:
		{
			return "D8.DC27"
		}
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf45) + ")"
		}
	case D8_DC28:
		{
			return "D8.DC28" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC25:
		{
			_, ok := other.Get_().(D8_DC25)
			return ok
		}
	case D8_DC26:
		{
			data2, ok := other.Get_().(D8_DC26)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Equals(data2.Cf47)
		}
	case D8_DC27:
		{
			_, ok := other.Get_().(D8_DC27)
			return ok
		}
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf45.Equals(data2.Cf45)
		}
	case D8_DC28:
		{
			data2, ok := other.Get_().(D8_DC28)
			return ok && data1.Cf48.Equals(data2.Cf48)
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

type D9_DC30 struct {
	Cf50 _dafny.CodePoint
	Cf51 bool
	Cf52 _dafny.Int
	Cf53 _dafny.Set
	Cf54 _dafny.Int
}

func (D9_DC30) isD9() {}

func (CompanionStruct_D9_) Create_DC30_(Cf50 _dafny.CodePoint, Cf51 bool, Cf52 _dafny.Int, Cf53 _dafny.Set, Cf54 _dafny.Int) D9 {
	return D9{D9_DC30{Cf50, Cf51, Cf52, Cf53, Cf54}}
}

func (_this D9) Is_DC30() bool {
	_, ok := _this.Get_().(D9_DC30)
	return ok
}

type D9_DC29 struct {
	Cf49 _dafny.Map
}

func (D9_DC29) isD9() {}

func (CompanionStruct_D9_) Create_DC29_(Cf49 _dafny.Map) D9 {
	return D9{D9_DC29{Cf49}}
}

func (_this D9) Is_DC29() bool {
	_, ok := _this.Get_().(D9_DC29)
	return ok
}

type D9_DC31 struct {
	Cf55 D9
}

func (D9_DC31) isD9() {}

func (CompanionStruct_D9_) Create_DC31_(Cf55 D9) D9 {
	return D9{D9_DC31{Cf55}}
}

func (_this D9) Is_DC31() bool {
	_, ok := _this.Get_().(D9_DC31)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC30_(_dafny.CodePoint('D'), false, _dafny.Zero, _dafny.EmptySet, _dafny.Zero)
}

func (_this D9) Dtor_cf50() _dafny.CodePoint {
	return _this.Get_().(D9_DC30).Cf50
}

func (_this D9) Dtor_cf51() bool {
	return _this.Get_().(D9_DC30).Cf51
}

func (_this D9) Dtor_cf52() _dafny.Int {
	return _this.Get_().(D9_DC30).Cf52
}

func (_this D9) Dtor_cf53() _dafny.Set {
	return _this.Get_().(D9_DC30).Cf53
}

func (_this D9) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D9_DC30).Cf54
}

func (_this D9) Dtor_cf49() _dafny.Map {
	return _this.Get_().(D9_DC29).Cf49
}

func (_this D9) Dtor_cf55() D9 {
	return _this.Get_().(D9_DC31).Cf55
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC30:
		{
			return "D9.DC30" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D9_DC29:
		{
			return "D9.DC29" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D9_DC31:
		{
			return "D9.DC31" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC30:
		{
			data2, ok := other.Get_().(D9_DC30)
			return ok && data1.Cf50 == data2.Cf50 && data1.Cf51 == data2.Cf51 && data1.Cf52.Cmp(data2.Cf52) == 0 && data1.Cf53.Equals(data2.Cf53) && data1.Cf54.Cmp(data2.Cf54) == 0
		}
	case D9_DC29:
		{
			data2, ok := other.Get_().(D9_DC29)
			return ok && data1.Cf49.Equals(data2.Cf49)
		}
	case D9_DC31:
		{
			data2, ok := other.Get_().(D9_DC31)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D10_DC32 struct {
	Cf56 _dafny.Sequence
}

func (D10_DC32) isD10() {}

func (CompanionStruct_D10_) Create_DC32_(Cf56 _dafny.Sequence) D10 {
	return D10{D10_DC32{Cf56}}
}

func (_this D10) Is_DC32() bool {
	_, ok := _this.Get_().(D10_DC32)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D10) Dtor_cf56() _dafny.Sequence {
	return _this.Get_().(D10_DC32).Cf56
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC32:
		{
			return "D10.DC32" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC32:
		{
			data2, ok := other.Get_().(D10_DC32)
			return ok && data1.Cf56.Equals(data2.Cf56)
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

type D11_DC34 struct {
	Cf58 _dafny.MultiSet
	Cf59 _dafny.MultiSet
	Cf60 _dafny.Array
	Cf61 D1
}

func (D11_DC34) isD11() {}

func (CompanionStruct_D11_) Create_DC34_(Cf58 _dafny.MultiSet, Cf59 _dafny.MultiSet, Cf60 _dafny.Array, Cf61 D1) D11 {
	return D11{D11_DC34{Cf58, Cf59, Cf60, Cf61}}
}

func (_this D11) Is_DC34() bool {
	_, ok := _this.Get_().(D11_DC34)
	return ok
}

type D11_DC35 struct {
}

func (D11_DC35) isD11() {}

func (CompanionStruct_D11_) Create_DC35_() D11 {
	return D11{D11_DC35{}}
}

func (_this D11) Is_DC35() bool {
	_, ok := _this.Get_().(D11_DC35)
	return ok
}

type D11_DC33 struct {
	Cf57 _dafny.Set
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf57 _dafny.Set) D11 {
	return D11{D11_DC33{Cf57}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC36 struct {
	Cf62 D11
}

func (D11_DC36) isD11() {}

func (CompanionStruct_D11_) Create_DC36_(Cf62 D11) D11 {
	return D11{D11_DC36{Cf62}}
}

func (_this D11) Is_DC36() bool {
	_, ok := _this.Get_().(D11_DC36)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC34_(_dafny.EmptyMultiSet, _dafny.EmptyMultiSet, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), Companion_D1_.Default())
}

func (_this D11) Dtor_cf58() _dafny.MultiSet {
	return _this.Get_().(D11_DC34).Cf58
}

func (_this D11) Dtor_cf59() _dafny.MultiSet {
	return _this.Get_().(D11_DC34).Cf59
}

func (_this D11) Dtor_cf60() _dafny.Array {
	return _this.Get_().(D11_DC34).Cf60
}

func (_this D11) Dtor_cf61() D1 {
	return _this.Get_().(D11_DC34).Cf61
}

func (_this D11) Dtor_cf57() _dafny.Set {
	return _this.Get_().(D11_DC33).Cf57
}

func (_this D11) Dtor_cf62() D11 {
	return _this.Get_().(D11_DC36).Cf62
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC34:
		{
			return "D11.DC34" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ")"
		}
	case D11_DC35:
		{
			return "D11.DC35"
		}
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf57) + ")"
		}
	case D11_DC36:
		{
			return "D11.DC36" + "(" + _dafny.String(data.Cf62) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC34:
		{
			data2, ok := other.Get_().(D11_DC34)
			return ok && data1.Cf58.Equals(data2.Cf58) && data1.Cf59.Equals(data2.Cf59) && data1.Cf60 == data2.Cf60 && data1.Cf61.Equals(data2.Cf61)
		}
	case D11_DC35:
		{
			_, ok := other.Get_().(D11_DC35)
			return ok
		}
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf57.Equals(data2.Cf57)
		}
	case D11_DC36:
		{
			data2, ok := other.Get_().(D11_DC36)
			return ok && data1.Cf62.Equals(data2.Cf62)
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

type D12_DC38 struct {
	Cf64 _dafny.Int
}

func (D12_DC38) isD12() {}

func (CompanionStruct_D12_) Create_DC38_(Cf64 _dafny.Int) D12 {
	return D12{D12_DC38{Cf64}}
}

func (_this D12) Is_DC38() bool {
	_, ok := _this.Get_().(D12_DC38)
	return ok
}

type D12_DC39 struct {
	Cf65 _dafny.Int
}

func (D12_DC39) isD12() {}

func (CompanionStruct_D12_) Create_DC39_(Cf65 _dafny.Int) D12 {
	return D12{D12_DC39{Cf65}}
}

func (_this D12) Is_DC39() bool {
	_, ok := _this.Get_().(D12_DC39)
	return ok
}

type D12_DC37 struct {
	Cf63 _dafny.MultiSet
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf63 _dafny.MultiSet) D12 {
	return D12{D12_DC37{Cf63}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC40 struct {
	Cf66 D12
}

func (D12_DC40) isD12() {}

func (CompanionStruct_D12_) Create_DC40_(Cf66 D12) D12 {
	return D12{D12_DC40{Cf66}}
}

func (_this D12) Is_DC40() bool {
	_, ok := _this.Get_().(D12_DC40)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC38_(_dafny.Zero)
}

func (_this D12) Dtor_cf64() _dafny.Int {
	return _this.Get_().(D12_DC38).Cf64
}

func (_this D12) Dtor_cf65() _dafny.Int {
	return _this.Get_().(D12_DC39).Cf65
}

func (_this D12) Dtor_cf63() _dafny.MultiSet {
	return _this.Get_().(D12_DC37).Cf63
}

func (_this D12) Dtor_cf66() D12 {
	return _this.Get_().(D12_DC40).Cf66
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC38:
		{
			return "D12.DC38" + "(" + _dafny.String(data.Cf64) + ")"
		}
	case D12_DC39:
		{
			return "D12.DC39" + "(" + _dafny.String(data.Cf65) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf63) + ")"
		}
	case D12_DC40:
		{
			return "D12.DC40" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC38:
		{
			data2, ok := other.Get_().(D12_DC38)
			return ok && data1.Cf64.Cmp(data2.Cf64) == 0
		}
	case D12_DC39:
		{
			data2, ok := other.Get_().(D12_DC39)
			return ok && data1.Cf65.Cmp(data2.Cf65) == 0
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf63.Equals(data2.Cf63)
		}
	case D12_DC40:
		{
			data2, ok := other.Get_().(D12_DC40)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D13_DC42 struct {
	Cf68 _dafny.CodePoint
	Cf69 bool
}

func (D13_DC42) isD13() {}

func (CompanionStruct_D13_) Create_DC42_(Cf68 _dafny.CodePoint, Cf69 bool) D13 {
	return D13{D13_DC42{Cf68, Cf69}}
}

func (_this D13) Is_DC42() bool {
	_, ok := _this.Get_().(D13_DC42)
	return ok
}

type D13_DC41 struct {
	Cf67 _dafny.Set
}

func (D13_DC41) isD13() {}

func (CompanionStruct_D13_) Create_DC41_(Cf67 _dafny.Set) D13 {
	return D13{D13_DC41{Cf67}}
}

func (_this D13) Is_DC41() bool {
	_, ok := _this.Get_().(D13_DC41)
	return ok
}

type D13_DC43 struct {
	Cf70 D13
}

func (D13_DC43) isD13() {}

func (CompanionStruct_D13_) Create_DC43_(Cf70 D13) D13 {
	return D13{D13_DC43{Cf70}}
}

func (_this D13) Is_DC43() bool {
	_, ok := _this.Get_().(D13_DC43)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC42_(_dafny.CodePoint('D'), false)
}

func (_this D13) Dtor_cf68() _dafny.CodePoint {
	return _this.Get_().(D13_DC42).Cf68
}

func (_this D13) Dtor_cf69() bool {
	return _this.Get_().(D13_DC42).Cf69
}

func (_this D13) Dtor_cf67() _dafny.Set {
	return _this.Get_().(D13_DC41).Cf67
}

func (_this D13) Dtor_cf70() D13 {
	return _this.Get_().(D13_DC43).Cf70
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC42:
		{
			return "D13.DC42" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D13_DC41:
		{
			return "D13.DC41" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D13_DC43:
		{
			return "D13.DC43" + "(" + _dafny.String(data.Cf70) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC42:
		{
			data2, ok := other.Get_().(D13_DC42)
			return ok && data1.Cf68 == data2.Cf68 && data1.Cf69 == data2.Cf69
		}
	case D13_DC41:
		{
			data2, ok := other.Get_().(D13_DC41)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	case D13_DC43:
		{
			data2, ok := other.Get_().(D13_DC43)
			return ok && data1.Cf70.Equals(data2.Cf70)
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

type D14_DC45 struct {
	Cf72 _dafny.Map
	Cf73 *C0
	Cf74 _dafny.Int
}

func (D14_DC45) isD14() {}

func (CompanionStruct_D14_) Create_DC45_(Cf72 _dafny.Map, Cf73 *C0, Cf74 _dafny.Int) D14 {
	return D14{D14_DC45{Cf72, Cf73, Cf74}}
}

func (_this D14) Is_DC45() bool {
	_, ok := _this.Get_().(D14_DC45)
	return ok
}

type D14_DC44 struct {
	Cf71 _dafny.Array
}

func (D14_DC44) isD14() {}

func (CompanionStruct_D14_) Create_DC44_(Cf71 _dafny.Array) D14 {
	return D14{D14_DC44{Cf71}}
}

func (_this D14) Is_DC44() bool {
	_, ok := _this.Get_().(D14_DC44)
	return ok
}

type D14_DC46 struct {
	Cf75 D14
}

func (D14_DC46) isD14() {}

func (CompanionStruct_D14_) Create_DC46_(Cf75 D14) D14 {
	return D14{D14_DC46{Cf75}}
}

func (_this D14) Is_DC46() bool {
	_, ok := _this.Get_().(D14_DC46)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC45_(_dafny.EmptyMap, (*C0)(nil), _dafny.Zero)
}

func (_this D14) Dtor_cf72() _dafny.Map {
	return _this.Get_().(D14_DC45).Cf72
}

func (_this D14) Dtor_cf73() *C0 {
	return _this.Get_().(D14_DC45).Cf73
}

func (_this D14) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D14_DC45).Cf74
}

func (_this D14) Dtor_cf71() _dafny.Array {
	return _this.Get_().(D14_DC44).Cf71
}

func (_this D14) Dtor_cf75() D14 {
	return _this.Get_().(D14_DC46).Cf75
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC45:
		{
			return "D14.DC45" + "(" + _dafny.String(data.Cf72) + ", " + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ")"
		}
	case D14_DC44:
		{
			return "D14.DC44" + "(" + _dafny.String(data.Cf71) + ")"
		}
	case D14_DC46:
		{
			return "D14.DC46" + "(" + _dafny.String(data.Cf75) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC45:
		{
			data2, ok := other.Get_().(D14_DC45)
			return ok && data1.Cf72.Equals(data2.Cf72) && data1.Cf73 == data2.Cf73 && data1.Cf74.Cmp(data2.Cf74) == 0
		}
	case D14_DC44:
		{
			data2, ok := other.Get_().(D14_DC44)
			return ok && data1.Cf71 == data2.Cf71
		}
	case D14_DC46:
		{
			data2, ok := other.Get_().(D14_DC46)
			return ok && data1.Cf75.Equals(data2.Cf75)
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

type D15_DC48 struct {
	Cf77 _dafny.Int
	Cf78 *C5
	Cf79 _dafny.Int
}

func (D15_DC48) isD15() {}

func (CompanionStruct_D15_) Create_DC48_(Cf77 _dafny.Int, Cf78 *C5, Cf79 _dafny.Int) D15 {
	return D15{D15_DC48{Cf77, Cf78, Cf79}}
}

func (_this D15) Is_DC48() bool {
	_, ok := _this.Get_().(D15_DC48)
	return ok
}

type D15_DC49 struct {
	Cf80 _dafny.Map
	Cf81 _dafny.Int
}

func (D15_DC49) isD15() {}

func (CompanionStruct_D15_) Create_DC49_(Cf80 _dafny.Map, Cf81 _dafny.Int) D15 {
	return D15{D15_DC49{Cf80, Cf81}}
}

func (_this D15) Is_DC49() bool {
	_, ok := _this.Get_().(D15_DC49)
	return ok
}

type D15_DC47 struct {
	Cf76 _dafny.Map
}

func (D15_DC47) isD15() {}

func (CompanionStruct_D15_) Create_DC47_(Cf76 _dafny.Map) D15 {
	return D15{D15_DC47{Cf76}}
}

func (_this D15) Is_DC47() bool {
	_, ok := _this.Get_().(D15_DC47)
	return ok
}

type D15_DC50 struct {
	Cf82 D15
}

func (D15_DC50) isD15() {}

func (CompanionStruct_D15_) Create_DC50_(Cf82 D15) D15 {
	return D15{D15_DC50{Cf82}}
}

func (_this D15) Is_DC50() bool {
	_, ok := _this.Get_().(D15_DC50)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC48_(_dafny.Zero, (*C5)(nil), _dafny.Zero)
}

func (_this D15) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D15_DC48).Cf77
}

func (_this D15) Dtor_cf78() *C5 {
	return _this.Get_().(D15_DC48).Cf78
}

func (_this D15) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D15_DC48).Cf79
}

func (_this D15) Dtor_cf80() _dafny.Map {
	return _this.Get_().(D15_DC49).Cf80
}

func (_this D15) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D15_DC49).Cf81
}

func (_this D15) Dtor_cf76() _dafny.Map {
	return _this.Get_().(D15_DC47).Cf76
}

func (_this D15) Dtor_cf82() D15 {
	return _this.Get_().(D15_DC50).Cf82
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC48:
		{
			return "D15.DC48" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D15_DC49:
		{
			return "D15.DC49" + "(" + _dafny.String(data.Cf80) + ", " + _dafny.String(data.Cf81) + ")"
		}
	case D15_DC47:
		{
			return "D15.DC47" + "(" + _dafny.String(data.Cf76) + ")"
		}
	case D15_DC50:
		{
			return "D15.DC50" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC48:
		{
			data2, ok := other.Get_().(D15_DC48)
			return ok && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78 == data2.Cf78 && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D15_DC49:
		{
			data2, ok := other.Get_().(D15_DC49)
			return ok && data1.Cf80.Equals(data2.Cf80) && data1.Cf81.Cmp(data2.Cf81) == 0
		}
	case D15_DC47:
		{
			data2, ok := other.Get_().(D15_DC47)
			return ok && data1.Cf76.Equals(data2.Cf76)
		}
	case D15_DC50:
		{
			data2, ok := other.Get_().(D15_DC50)
			return ok && data1.Cf82.Equals(data2.Cf82)
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

type D16_DC52 struct {
	Cf84 _dafny.Int
	Cf85 bool
	Cf86 _dafny.Array
	Cf87 bool
	Cf88 _dafny.Int
}

func (D16_DC52) isD16() {}

func (CompanionStruct_D16_) Create_DC52_(Cf84 _dafny.Int, Cf85 bool, Cf86 _dafny.Array, Cf87 bool, Cf88 _dafny.Int) D16 {
	return D16{D16_DC52{Cf84, Cf85, Cf86, Cf87, Cf88}}
}

func (_this D16) Is_DC52() bool {
	_, ok := _this.Get_().(D16_DC52)
	return ok
}

type D16_DC51 struct {
	Cf83 _dafny.Array
}

func (D16_DC51) isD16() {}

func (CompanionStruct_D16_) Create_DC51_(Cf83 _dafny.Array) D16 {
	return D16{D16_DC51{Cf83}}
}

func (_this D16) Is_DC51() bool {
	_, ok := _this.Get_().(D16_DC51)
	return ok
}

type D16_DC53 struct {
	Cf89 D16
}

func (D16_DC53) isD16() {}

func (CompanionStruct_D16_) Create_DC53_(Cf89 D16) D16 {
	return D16{D16_DC53{Cf89}}
}

func (_this D16) Is_DC53() bool {
	_, ok := _this.Get_().(D16_DC53)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC52_(_dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero)
}

func (_this D16) Dtor_cf84() _dafny.Int {
	return _this.Get_().(D16_DC52).Cf84
}

func (_this D16) Dtor_cf85() bool {
	return _this.Get_().(D16_DC52).Cf85
}

func (_this D16) Dtor_cf86() _dafny.Array {
	return _this.Get_().(D16_DC52).Cf86
}

func (_this D16) Dtor_cf87() bool {
	return _this.Get_().(D16_DC52).Cf87
}

func (_this D16) Dtor_cf88() _dafny.Int {
	return _this.Get_().(D16_DC52).Cf88
}

func (_this D16) Dtor_cf83() _dafny.Array {
	return _this.Get_().(D16_DC51).Cf83
}

func (_this D16) Dtor_cf89() D16 {
	return _this.Get_().(D16_DC53).Cf89
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC52:
		{
			return "D16.DC52" + "(" + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ", " + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ", " + _dafny.String(data.Cf88) + ")"
		}
	case D16_DC51:
		{
			return "D16.DC51" + "(" + _dafny.String(data.Cf83) + ")"
		}
	case D16_DC53:
		{
			return "D16.DC53" + "(" + _dafny.String(data.Cf89) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC52:
		{
			data2, ok := other.Get_().(D16_DC52)
			return ok && data1.Cf84.Cmp(data2.Cf84) == 0 && data1.Cf85 == data2.Cf85 && data1.Cf86 == data2.Cf86 && data1.Cf87 == data2.Cf87 && data1.Cf88.Cmp(data2.Cf88) == 0
		}
	case D16_DC51:
		{
			data2, ok := other.Get_().(D16_DC51)
			return ok && data1.Cf83 == data2.Cf83
		}
	case D16_DC53:
		{
			data2, ok := other.Get_().(D16_DC53)
			return ok && data1.Cf89.Equals(data2.Cf89)
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

type D17_DC54 struct {
	Cf90 _dafny.Array
}

func (D17_DC54) isD17() {}

func (CompanionStruct_D17_) Create_DC54_(Cf90 _dafny.Array) D17 {
	return D17{D17_DC54{Cf90}}
}

func (_this D17) Is_DC54() bool {
	_, ok := _this.Get_().(D17_DC54)
	return ok
}

func (CompanionStruct_D17_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D17) Dtor_cf90() _dafny.Array {
	return _this.Get_().(D17_DC54).Cf90
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC54:
		{
			return "D17.DC54" + "(" + _dafny.String(data.Cf90) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC54:
		{
			data2, ok := other.Get_().(D17_DC54)
			return ok && data1.Cf90 == data2.Cf90
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

type D18_DC56 struct {
	Cf92 bool
	Cf93 bool
	Cf94 _dafny.Int
}

func (D18_DC56) isD18() {}

func (CompanionStruct_D18_) Create_DC56_(Cf92 bool, Cf93 bool, Cf94 _dafny.Int) D18 {
	return D18{D18_DC56{Cf92, Cf93, Cf94}}
}

func (_this D18) Is_DC56() bool {
	_, ok := _this.Get_().(D18_DC56)
	return ok
}

type D18_DC57 struct {
	Cf95 *C0
}

func (D18_DC57) isD18() {}

func (CompanionStruct_D18_) Create_DC57_(Cf95 *C0) D18 {
	return D18{D18_DC57{Cf95}}
}

func (_this D18) Is_DC57() bool {
	_, ok := _this.Get_().(D18_DC57)
	return ok
}

type D18_DC55 struct {
	Cf91 *C3
}

func (D18_DC55) isD18() {}

func (CompanionStruct_D18_) Create_DC55_(Cf91 *C3) D18 {
	return D18{D18_DC55{Cf91}}
}

func (_this D18) Is_DC55() bool {
	_, ok := _this.Get_().(D18_DC55)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC56_(false, false, _dafny.Zero)
}

func (_this D18) Dtor_cf92() bool {
	return _this.Get_().(D18_DC56).Cf92
}

func (_this D18) Dtor_cf93() bool {
	return _this.Get_().(D18_DC56).Cf93
}

func (_this D18) Dtor_cf94() _dafny.Int {
	return _this.Get_().(D18_DC56).Cf94
}

func (_this D18) Dtor_cf95() *C0 {
	return _this.Get_().(D18_DC57).Cf95
}

func (_this D18) Dtor_cf91() *C3 {
	return _this.Get_().(D18_DC55).Cf91
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC56:
		{
			return "D18.DC56" + "(" + _dafny.String(data.Cf92) + ", " + _dafny.String(data.Cf93) + ", " + _dafny.String(data.Cf94) + ")"
		}
	case D18_DC57:
		{
			return "D18.DC57" + "(" + _dafny.String(data.Cf95) + ")"
		}
	case D18_DC55:
		{
			return "D18.DC55" + "(" + _dafny.String(data.Cf91) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC56:
		{
			data2, ok := other.Get_().(D18_DC56)
			return ok && data1.Cf92 == data2.Cf92 && data1.Cf93 == data2.Cf93 && data1.Cf94.Cmp(data2.Cf94) == 0
		}
	case D18_DC57:
		{
			data2, ok := other.Get_().(D18_DC57)
			return ok && data1.Cf95 == data2.Cf95
		}
	case D18_DC55:
		{
			data2, ok := other.Get_().(D18_DC55)
			return ok && data1.Cf91 == data2.Cf91
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

type D19_DC59 struct {
	Cf97 T0
	Cf98 _dafny.Map
}

func (D19_DC59) isD19() {}

func (CompanionStruct_D19_) Create_DC59_(Cf97 T0, Cf98 _dafny.Map) D19 {
	return D19{D19_DC59{Cf97, Cf98}}
}

func (_this D19) Is_DC59() bool {
	_, ok := _this.Get_().(D19_DC59)
	return ok
}

type D19_DC58 struct {
	Cf96 _dafny.Map
}

func (D19_DC58) isD19() {}

func (CompanionStruct_D19_) Create_DC58_(Cf96 _dafny.Map) D19 {
	return D19{D19_DC58{Cf96}}
}

func (_this D19) Is_DC58() bool {
	_, ok := _this.Get_().(D19_DC58)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC59_((T0)(nil), _dafny.EmptyMap)
}

func (_this D19) Dtor_cf97() T0 {
	return _this.Get_().(D19_DC59).Cf97
}

func (_this D19) Dtor_cf98() _dafny.Map {
	return _this.Get_().(D19_DC59).Cf98
}

func (_this D19) Dtor_cf96() _dafny.Map {
	return _this.Get_().(D19_DC58).Cf96
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC59:
		{
			return "D19.DC59" + "(" + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ")"
		}
	case D19_DC58:
		{
			return "D19.DC58" + "(" + _dafny.String(data.Cf96) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC59:
		{
			data2, ok := other.Get_().(D19_DC59)
			return ok && _dafny.AreEqual(data1.Cf97, data2.Cf97) && data1.Cf98.Equals(data2.Cf98)
		}
	case D19_DC58:
		{
			data2, ok := other.Get_().(D19_DC58)
			return ok && data1.Cf96.Equals(data2.Cf96)
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

type D20_DC61 struct {
}

func (D20_DC61) isD20() {}

func (CompanionStruct_D20_) Create_DC61_() D20 {
	return D20{D20_DC61{}}
}

func (_this D20) Is_DC61() bool {
	_, ok := _this.Get_().(D20_DC61)
	return ok
}

type D20_DC60 struct {
	Cf99 *C4
}

func (D20_DC60) isD20() {}

func (CompanionStruct_D20_) Create_DC60_(Cf99 *C4) D20 {
	return D20{D20_DC60{Cf99}}
}

func (_this D20) Is_DC60() bool {
	_, ok := _this.Get_().(D20_DC60)
	return ok
}

type D20_DC62 struct {
	Cf100 D20
}

func (D20_DC62) isD20() {}

func (CompanionStruct_D20_) Create_DC62_(Cf100 D20) D20 {
	return D20{D20_DC62{Cf100}}
}

func (_this D20) Is_DC62() bool {
	_, ok := _this.Get_().(D20_DC62)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC61_()
}

func (_this D20) Dtor_cf99() *C4 {
	return _this.Get_().(D20_DC60).Cf99
}

func (_this D20) Dtor_cf100() D20 {
	return _this.Get_().(D20_DC62).Cf100
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC61:
		{
			return "D20.DC61"
		}
	case D20_DC60:
		{
			return "D20.DC60" + "(" + _dafny.String(data.Cf99) + ")"
		}
	case D20_DC62:
		{
			return "D20.DC62" + "(" + _dafny.String(data.Cf100) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC61:
		{
			_, ok := other.Get_().(D20_DC61)
			return ok
		}
	case D20_DC60:
		{
			data2, ok := other.Get_().(D20_DC60)
			return ok && data1.Cf99 == data2.Cf99
		}
	case D20_DC62:
		{
			data2, ok := other.Get_().(D20_DC62)
			return ok && data1.Cf100.Equals(data2.Cf100)
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

type D21_DC64 struct {
	Cf102 bool
	Cf103 *C3
}

func (D21_DC64) isD21() {}

func (CompanionStruct_D21_) Create_DC64_(Cf102 bool, Cf103 *C3) D21 {
	return D21{D21_DC64{Cf102, Cf103}}
}

func (_this D21) Is_DC64() bool {
	_, ok := _this.Get_().(D21_DC64)
	return ok
}

type D21_DC65 struct {
	Cf104 _dafny.Int
	Cf105 _dafny.Int
	Cf106 *C1
}

func (D21_DC65) isD21() {}

func (CompanionStruct_D21_) Create_DC65_(Cf104 _dafny.Int, Cf105 _dafny.Int, Cf106 *C1) D21 {
	return D21{D21_DC65{Cf104, Cf105, Cf106}}
}

func (_this D21) Is_DC65() bool {
	_, ok := _this.Get_().(D21_DC65)
	return ok
}

type D21_DC63 struct {
	Cf101 *C6
}

func (D21_DC63) isD21() {}

func (CompanionStruct_D21_) Create_DC63_(Cf101 *C6) D21 {
	return D21{D21_DC63{Cf101}}
}

func (_this D21) Is_DC63() bool {
	_, ok := _this.Get_().(D21_DC63)
	return ok
}

type D21_DC66 struct {
	Cf107 D21
}

func (D21_DC66) isD21() {}

func (CompanionStruct_D21_) Create_DC66_(Cf107 D21) D21 {
	return D21{D21_DC66{Cf107}}
}

func (_this D21) Is_DC66() bool {
	_, ok := _this.Get_().(D21_DC66)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC64_(false, (*C3)(nil))
}

func (_this D21) Dtor_cf102() bool {
	return _this.Get_().(D21_DC64).Cf102
}

func (_this D21) Dtor_cf103() *C3 {
	return _this.Get_().(D21_DC64).Cf103
}

func (_this D21) Dtor_cf104() _dafny.Int {
	return _this.Get_().(D21_DC65).Cf104
}

func (_this D21) Dtor_cf105() _dafny.Int {
	return _this.Get_().(D21_DC65).Cf105
}

func (_this D21) Dtor_cf106() *C1 {
	return _this.Get_().(D21_DC65).Cf106
}

func (_this D21) Dtor_cf101() *C6 {
	return _this.Get_().(D21_DC63).Cf101
}

func (_this D21) Dtor_cf107() D21 {
	return _this.Get_().(D21_DC66).Cf107
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC64:
		{
			return "D21.DC64" + "(" + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ")"
		}
	case D21_DC65:
		{
			return "D21.DC65" + "(" + _dafny.String(data.Cf104) + ", " + _dafny.String(data.Cf105) + ", " + _dafny.String(data.Cf106) + ")"
		}
	case D21_DC63:
		{
			return "D21.DC63" + "(" + _dafny.String(data.Cf101) + ")"
		}
	case D21_DC66:
		{
			return "D21.DC66" + "(" + _dafny.String(data.Cf107) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC64:
		{
			data2, ok := other.Get_().(D21_DC64)
			return ok && data1.Cf102 == data2.Cf102 && data1.Cf103 == data2.Cf103
		}
	case D21_DC65:
		{
			data2, ok := other.Get_().(D21_DC65)
			return ok && data1.Cf104.Cmp(data2.Cf104) == 0 && data1.Cf105.Cmp(data2.Cf105) == 0 && data1.Cf106 == data2.Cf106
		}
	case D21_DC63:
		{
			data2, ok := other.Get_().(D21_DC63)
			return ok && data1.Cf101 == data2.Cf101
		}
	case D21_DC66:
		{
			data2, ok := other.Get_().(D21_DC66)
			return ok && data1.Cf107.Equals(data2.Cf107)
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

type D22_DC68 struct {
	Cf109 bool
	Cf110 _dafny.Int
}

func (D22_DC68) isD22() {}

func (CompanionStruct_D22_) Create_DC68_(Cf109 bool, Cf110 _dafny.Int) D22 {
	return D22{D22_DC68{Cf109, Cf110}}
}

func (_this D22) Is_DC68() bool {
	_, ok := _this.Get_().(D22_DC68)
	return ok
}

type D22_DC69 struct {
	Cf111 bool
}

func (D22_DC69) isD22() {}

func (CompanionStruct_D22_) Create_DC69_(Cf111 bool) D22 {
	return D22{D22_DC69{Cf111}}
}

func (_this D22) Is_DC69() bool {
	_, ok := _this.Get_().(D22_DC69)
	return ok
}

type D22_DC70 struct {
	Cf112 bool
	Cf113 bool
	Cf114 _dafny.Sequence
	Cf115 _dafny.Int
	Cf116 _dafny.Int
}

func (D22_DC70) isD22() {}

func (CompanionStruct_D22_) Create_DC70_(Cf112 bool, Cf113 bool, Cf114 _dafny.Sequence, Cf115 _dafny.Int, Cf116 _dafny.Int) D22 {
	return D22{D22_DC70{Cf112, Cf113, Cf114, Cf115, Cf116}}
}

func (_this D22) Is_DC70() bool {
	_, ok := _this.Get_().(D22_DC70)
	return ok
}

type D22_DC67 struct {
	Cf108 _dafny.Sequence
}

func (D22_DC67) isD22() {}

func (CompanionStruct_D22_) Create_DC67_(Cf108 _dafny.Sequence) D22 {
	return D22{D22_DC67{Cf108}}
}

func (_this D22) Is_DC67() bool {
	_, ok := _this.Get_().(D22_DC67)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC68_(false, _dafny.Zero)
}

func (_this D22) Dtor_cf109() bool {
	return _this.Get_().(D22_DC68).Cf109
}

func (_this D22) Dtor_cf110() _dafny.Int {
	return _this.Get_().(D22_DC68).Cf110
}

func (_this D22) Dtor_cf111() bool {
	return _this.Get_().(D22_DC69).Cf111
}

func (_this D22) Dtor_cf112() bool {
	return _this.Get_().(D22_DC70).Cf112
}

func (_this D22) Dtor_cf113() bool {
	return _this.Get_().(D22_DC70).Cf113
}

func (_this D22) Dtor_cf114() _dafny.Sequence {
	return _this.Get_().(D22_DC70).Cf114
}

func (_this D22) Dtor_cf115() _dafny.Int {
	return _this.Get_().(D22_DC70).Cf115
}

func (_this D22) Dtor_cf116() _dafny.Int {
	return _this.Get_().(D22_DC70).Cf116
}

func (_this D22) Dtor_cf108() _dafny.Sequence {
	return _this.Get_().(D22_DC67).Cf108
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC68:
		{
			return "D22.DC68" + "(" + _dafny.String(data.Cf109) + ", " + _dafny.String(data.Cf110) + ")"
		}
	case D22_DC69:
		{
			return "D22.DC69" + "(" + _dafny.String(data.Cf111) + ")"
		}
	case D22_DC70:
		{
			return "D22.DC70" + "(" + _dafny.String(data.Cf112) + ", " + _dafny.String(data.Cf113) + ", " + data.Cf114.VerbatimString(true) + ", " + _dafny.String(data.Cf115) + ", " + _dafny.String(data.Cf116) + ")"
		}
	case D22_DC67:
		{
			return "D22.DC67" + "(" + data.Cf108.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC68:
		{
			data2, ok := other.Get_().(D22_DC68)
			return ok && data1.Cf109 == data2.Cf109 && data1.Cf110.Cmp(data2.Cf110) == 0
		}
	case D22_DC69:
		{
			data2, ok := other.Get_().(D22_DC69)
			return ok && data1.Cf111 == data2.Cf111
		}
	case D22_DC70:
		{
			data2, ok := other.Get_().(D22_DC70)
			return ok && data1.Cf112 == data2.Cf112 && data1.Cf113 == data2.Cf113 && data1.Cf114.Equals(data2.Cf114) && data1.Cf115.Cmp(data2.Cf115) == 0 && data1.Cf116.Cmp(data2.Cf116) == 0
		}
	case D22_DC67:
		{
			data2, ok := other.Get_().(D22_DC67)
			return ok && data1.Cf108.Equals(data2.Cf108)
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

// Definition of trait T0
type T0 interface {
	String() string
	F14() bool
	F14_set_(value bool)
	Fm2(p0 bool, p1 bool, globalState *GlobalState) bool
	Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence
	M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool)
	M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool)
	F15() _dafny.MultiSet
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
	F2   _dafny.Array
	F3   _dafny.Int
	F6   _dafny.Set
	F9   bool
	F11  _dafny.Array
	F13  _dafny.Set
	_f0  _dafny.Int
	_f1  _dafny.MultiSet
	_f4  bool
	_f5  _dafny.Sequence
	_f7  bool
	_f8  _dafny.Int
	_f10 _dafny.Int
	_f12 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F3 = _dafny.Zero
	_this.F6 = _dafny.EmptySet
	_this.F9 = false
	_this.F11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this.F13 = _dafny.EmptySet
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.EmptyMultiSet
	_this._f4 = false
	_this._f5 = _dafny.EmptySeq
	_this._f7 = false
	_this._f8 = _dafny.Zero
	_this._f10 = _dafny.Zero
	_this._f12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.MultiSet, f2 _dafny.Array, f3 _dafny.Int, f4 bool, f5 _dafny.Sequence, f6 _dafny.Set, f7 bool, f8 _dafny.Int, f9 bool, f10 _dafny.Int, f11 _dafny.Array, f12 _dafny.Array, f13 _dafny.Set) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this).F3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this).F9 = f9
		(_this)._f10 = f10
		(_this).F11 = f11
		(_this)._f12 = f12
		(_this).F13 = f13
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.MultiSet {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F4() bool {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() _dafny.Sequence {
	{
		return _this._f5
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
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F12() _dafny.Array {
	{
		return _this._f12
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f23 bool
	_f24 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f23 = false
	_this._f24 = false
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

func (_this *C0) Ctor__(f23 bool, f24 bool) {
	{
		(_this)._f23 = f23
		(_this)._f24 = f24
	}
}
func (_this *C0) F23() bool {
	{
		return _this._f23
	}
}
func (_this *C0) F24() bool {
	{
		return _this._f24
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) M10(p0 _dafny.Array, p1 _dafny.Array, p2 bool, globalState *GlobalState) (D1, _dafny.Int, bool, _dafny.Int) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		if p2 {
			var _586_v0 _dafny.Array
			_ = _586_v0
			var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw90
			_586_v0 = _nw90
			var _587_v1 D5
			_ = _587_v1
			_587_v1 = Companion_D5_.Create_DC16_(p1)
			_586_v0 = (_587_v1).Dtor_cf29()
			var _588_v2 _dafny.Sequence
			_ = _588_v2
			_588_v2 = _dafny.UnicodeSeqOfUtf8Bytes("dhcjdj")
			var _589_v3 _dafny.Sequence
			_ = _589_v3
			_589_v3 = _dafny.SeqOf(_588_v2)
			var _590_v4 _dafny.Int
			_ = _590_v4
			_590_v4 = _dafny.IntOfInt64(-738)
			var _591_v5 *C0
			_ = _591_v5
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf((_589_v3).Select((Companion_Default___.SafeIndex(_590_v4, _dafny.IntOfUint32((_589_v3).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.UnicodeSeqOfUtf8Bytes("rylyonx")), true)
			_591_v5 = _nw91
			var _592_v6 *C0
			_ = _592_v6
			var _nw92 *C0 = New_C0_()
			_ = _nw92
			_nw92.Ctor__(!((_591_v5).F23()), (true) || ((_591_v5).F23()))
			_592_v6 = _nw92
			var _593_v7 _dafny.Map
			_ = _593_v7
			_593_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _590_v4)
			r3 = ((_dafny.Zero).Minus(_590_v4)).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((func() _dafny.Int {
				if (_593_v7).Contains(p2) {
					return (_593_v7).Get(p2).(_dafny.Int)
				}
				return _590_v4
			})()), _590_v4))
			var _594_v8 _dafny.Sequence
			_ = _594_v8
			_594_v8 = _dafny.SeqOf(_590_v4)
			_594_v8 = _594_v8
		} else {
			var _595_v9 _dafny.Array
			_ = _595_v9
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw93
			_595_v9 = _nw93
			(globalState).F2 = _595_v9
			var _596_v10 _dafny.Int
			_ = _596_v10
			_596_v10 = _dafny.IntOfInt64(489)
			var _597_v11 _dafny.Sequence
			_ = _597_v11
			_597_v11 = _dafny.SeqOf(_596_v10)
			var _598_v12 _dafny.Map
			_ = _598_v12
			_598_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _599_v13 _dafny.MultiSet
			_ = _599_v13
			_599_v13 = _dafny.MultiSetOf(p2, p2)
			var _600_v14 _dafny.Array
			_ = _600_v14
			var _nwElement0_14 _dafny.Int = _dafny.IntOfUint32((_597_v11).Cardinality())
			_ = _nwElement0_14
			var _nw94 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(27))
			_ = _nw94
			(_nw94).ArraySet1(_nwElement0_14, 0)
			(_nw94).ArraySet1(_596_v10, 1)
			(_nw94).ArraySet1((_596_v10).Minus(_dafny.IntOfInt64(620)), 2)
			(_nw94).ArraySet1(_596_v10, 3)
			(_nw94).ArraySet1((_598_v12).Cardinality(), 4)
			(_nw94).ArraySet1(Companion_Default___.Fm0(p2, _596_v10, _596_v10, _dafny.MultiSetOf(_596_v10, _596_v10, _596_v10), globalState), 5)
			(_nw94).ArraySet1(_596_v10, 6)
			(_nw94).ArraySet1(_596_v10, 7)
			(_nw94).ArraySet1(_596_v10, 8)
			(_nw94).ArraySet1(_596_v10, 9)
			(_nw94).ArraySet1(_596_v10, 10)
			(_nw94).ArraySet1(_596_v10, 11)
			(_nw94).ArraySet1(_596_v10, 12)
			(_nw94).ArraySet1(_596_v10, 13)
			(_nw94).ArraySet1(Companion_Default___.Fm0(p2, _596_v10, _596_v10, _dafny.MultiSetOf(_dafny.IntOfInt64(550), _596_v10), globalState), 14)
			(_nw94).ArraySet1(_596_v10, 15)
			(_nw94).ArraySet1(Companion_Default___.SafeModuloInt(_596_v10, _596_v10), 16)
			(_nw94).ArraySet1(((_599_v13).Intersection(_599_v13)).Cardinality(), 17)
			(_nw94).ArraySet1(_596_v10, 18)
			(_nw94).ArraySet1(_596_v10, 19)
			(_nw94).ArraySet1(_596_v10, 20)
			(_nw94).ArraySet1(_dafny.IntOfUint32((_597_v11).Cardinality()), 21)
			(_nw94).ArraySet1(_596_v10, 22)
			(_nw94).ArraySet1(_596_v10, 23)
			(_nw94).ArraySet1(_dafny.IntOfInt64(589), 24)
			(_nw94).ArraySet1(_596_v10, 25)
			(_nw94).ArraySet1(_596_v10, 26)
			_600_v14 = _nw94
			_600_v14 = p1
			var _601_v15 _dafny.Sequence
			_ = _601_v15
			_601_v15 = _dafny.SeqOf(p2, p2)
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))
			_ = _index84
			(p1).ArraySet1(_dafny.IntOfUint32((_601_v15).Cardinality()), (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))
			_ = _index85
			(p1).ArraySet1(_596_v10, (_index85).Int())
			var _602_v16 _dafny.Set
			_ = _602_v16
			_602_v16 = _dafny.SetOf(p2)
			var _603_v17 _dafny.MultiSet
			_ = _603_v17
			_603_v17 = _dafny.MultiSetOf((_602_v16).Cardinality())
			var _604_v18 D2
			_ = _604_v18
			_604_v18 = Companion_D2_.Create_DC7_(_601_v15)
			var _605_v19 _dafny.Sequence
			_ = _605_v19
			_605_v19 = _dafny.SeqOf(_604_v18, _604_v18, _604_v18, _604_v18, _604_v18)
			_597_v11 = Companion_Default___.Fm21(Companion_D1_.Create_DC3_(Companion_Default___.Fm0(p2, _596_v10, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_603_v17).Update(_596_v10, Companion_Default___.Abs(_596_v10)), globalState)), (_601_v15).Select((Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_601_v15).Cardinality()))).Uint32()).(bool), _605_v19, globalState)
			var _606_v20 _dafny.Set
			_ = _606_v20
			_606_v20 = _dafny.SetOf(_597_v11, _597_v11, _597_v11, _597_v11)
			var _607_v21 _dafny.Map
			_ = _607_v21
			_607_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(p2)), (_606_v20).Cardinality())
			var _608_v22 _dafny.Sequence
			_ = _608_v22
			_608_v22 = _dafny.UnicodeSeqOfUtf8Bytes("ye")
			var _609_v23 _dafny.Set
			_ = _609_v23
			_609_v23 = _dafny.SetOf(_596_v10)
			var _610_v24 _dafny.Sequence
			_ = _610_v24
			_610_v24 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_603_v17).Contains(_dafny.IntOfInt64(651)) {
					return (_603_v17).Multiplicity(_dafny.IntOfInt64(651))
				}
				return (_609_v23).Cardinality()
			})(), p2))
			var _611_v25 _dafny.Array
			_ = _611_v25
			var _nwElement0_15 _dafny.Map = _607_v21
			_ = _nwElement0_15
			var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(15))
			_ = _nw95
			(_nw95).ArraySet1(_nwElement0_15, 0)
			(_nw95).ArraySet1((func() _dafny.Map {
				if p2 {
					return _607_v21
				}
				return _607_v21
			})(), 1)
			(_nw95).ArraySet1(_607_v21, 2)
			(_nw95).ArraySet1((_607_v21).Merge(_607_v21), 3)
			(_nw95).ArraySet1((Companion_Default___.Fm1((func() bool {
				if (_598_v12).Contains(p2) {
					return (_598_v12).Get(p2).(bool)
				}
				return p2
			})(), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), true, Companion_Default___.Fm22(_dafny.IntOfUint32((_608_v22).Cardinality()), _610_v24, p2, p2, globalState), globalState)).Update(p2, _596_v10), 4)
			(_nw95).ArraySet1((_607_v21).Update(false, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)), 5)
			(_nw95).ArraySet1((_607_v21).Merge(_607_v21), 6)
			(_nw95).ArraySet1(_607_v21, 7)
			(_nw95).ArraySet1(_607_v21, 8)
			(_nw95).ArraySet1(_607_v21, 9)
			(_nw95).ArraySet1(_607_v21, 10)
			(_nw95).ArraySet1(_607_v21, 11)
			(_nw95).ArraySet1(_607_v21, 12)
			(_nw95).ArraySet1(_607_v21, 13)
			(_nw95).ArraySet1((_607_v21).Merge(_607_v21), 14)
			_611_v25 = _nw95
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_611_v25), 0))
			_ = _index86
			(_611_v25).ArraySet1((_607_v21).Merge((_607_v21).Update(p2, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))), (_index86).Int())
			var _612_v26 D1
			_ = _612_v26
			_612_v26 = Companion_D1_.Create_DC5_(true, p2)
			var _613_v27 _dafny.Set
			_ = _613_v27
			_613_v27 = _dafny.SetOf(Companion_D1_.Create_DC5_(!(p2), p2), _612_v26, _612_v26, _612_v26, _612_v26)
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_611_v25), 0))
			_ = _index87
			var _rhs75 _dafny.Int = ((_613_v27).Union(_613_v27)).Cardinality()
			_ = _rhs75
			var _rhs76 bool = p2
			_ = _rhs76
			var _rhs77 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(69), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))).Merge(_607_v21)
			_ = _rhs77
			var _rhs78 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_597_v11, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_614_i0 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), _dafny.CodePoint('o'))).Cardinality()
			})))
			_ = _rhs78
			var _lhs65 *GlobalState = globalState
			_ = _lhs65
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			var _lhs67 _dafny.Array = _611_v25
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_611_v25), 0))
			_ = _lhs68
			_lhs65.F3 = _rhs75
			_lhs66.F9 = _rhs76
			(_lhs67).ArraySet1(_rhs77, (_lhs68).Int())
			_597_v11 = _rhs78
		}
		var _615_v28 _dafny.Int
		_ = _615_v28
		_615_v28 = _dafny.IntOfInt64(267)
		var _616_v29 _dafny.Sequence
		_ = _616_v29
		_616_v29 = _dafny.SeqOf(_615_v28, _615_v28)
		var _617_v30 _dafny.CodePoint
		_ = _617_v30
		_617_v30 = _dafny.CodePoint('j')
		var _618_v31 _dafny.Map
		_ = _618_v31
		_618_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_617_v30, _615_v28)
		var _619_v32 _dafny.Set
		_ = _619_v32
		_619_v32 = _dafny.SetOf(_616_v29)
		if Companion_Default___.Fm23(Companion_Default___.SafeDivisionInt(_615_v28, _615_v28), Companion_Default___.Fm23(_615_v28, p2, Companion_Default___.Fm24(_615_v28, p2, p2, _615_v28, globalState), _dafny.SetOf(_616_v29, _616_v29), globalState), (_618_v31).Merge((_618_v31).Update(_617_v30, _615_v28)), _619_v32, globalState) {
			var _620_v33 _dafny.Sequence
			_ = _620_v33
			_620_v33 = _dafny.SeqOf(p2, Companion_Default___.Fm23(_dafny.IntOfInt64(-497), false, _618_v31, _619_v32, globalState))
			var _621_v34 D2
			_ = _621_v34
			_621_v34 = Companion_D2_.Create_DC7_(_620_v33)
			_620_v33 = (_621_v34).Dtor_cf9()
			var _622_v35 _dafny.Sequence
			_ = _622_v35
			_622_v35 = _dafny.UnicodeSeqOfUtf8Bytes("yjf")
			var _623_v36 _dafny.Map
			_ = _623_v36
			_623_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v28, _dafny.IntOfUint32((_622_v35).Cardinality()))
			var _624_v37 D1
			_ = _624_v37
			_624_v37 = Companion_D1_.Create_DC4_((_623_v36).Cardinality(), _dafny.IntOfInt64(112))
			var _source10 D1 = _624_v37
			_ = _source10
			if _source10.Is_DC4() {
				var _625___mcc_h0 _dafny.Int = _source10.Get_().(D1_DC4).Cf4
				_ = _625___mcc_h0
				var _626___mcc_h1 _dafny.Int = _source10.Get_().(D1_DC4).Cf5
				_ = _626___mcc_h1
				var _627_cf5 _dafny.Int = _626___mcc_h1
				_ = _627_cf5
				var _628_cf4 _dafny.Int = _625___mcc_h0
				_ = _628_cf4
				var _629_v38 D1
				_ = _629_v38
				_629_v38 = Companion_D1_.Create_DC3_(_627_cf5)
				var _630_v39 _dafny.Array
				_ = _630_v39
				var _nwElement0_16 bool = p2
				_ = _nwElement0_16
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(29))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_16, 0)
				(_nw96).ArraySet1(!(!(p2)), 1)
				(_nw96).ArraySet1(p2, 2)
				(_nw96).ArraySet1(false, 3)
				(_nw96).ArraySet1(Companion_Default___.Fm23(_615_v28, true, _618_v31, _dafny.SetOf(Companion_Default___.Fm21(_629_v38, p2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-918))).Uint32(), func(coer52 func(_dafny.Int) D2) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_631_v34 D2) func(_dafny.Int) D2 {
					return func(_632_i1 _dafny.Int) D2 {
						return _631_v34
					}
				})(_621_v34))), globalState)), globalState), 4)
				(_nw96).ArraySet1(p2, 5)
				(_nw96).ArraySet1(p2, 6)
				(_nw96).ArraySet1(p2, 7)
				(_nw96).ArraySet1(p2, 8)
				(_nw96).ArraySet1(p2, 9)
				(_nw96).ArraySet1(false, 10)
				(_nw96).ArraySet1(p2, 11)
				(_nw96).ArraySet1(p2, 12)
				(_nw96).ArraySet1(p2, 13)
				(_nw96).ArraySet1(!((_620_v33).Select((Companion_Default___.SafeIndex(_627_cf5, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool)), 14)
				(_nw96).ArraySet1(false, 15)
				(_nw96).ArraySet1(p2, 16)
				(_nw96).ArraySet1(true, 17)
				(_nw96).ArraySet1(p2, 18)
				(_nw96).ArraySet1(true, 19)
				(_nw96).ArraySet1(p2, 20)
				(_nw96).ArraySet1(p2, 21)
				(_nw96).ArraySet1(p2, 22)
				(_nw96).ArraySet1((_620_v33).Select((Companion_Default___.SafeIndex(_628_cf4, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool), 23)
				(_nw96).ArraySet1(!(false), 24)
				(_nw96).ArraySet1(p2, 25)
				(_nw96).ArraySet1(p2, 26)
				(_nw96).ArraySet1(p2, 27)
				(_nw96).ArraySet1(true, 28)
				_630_v39 = _nw96
				var _633_v40 _dafny.Map
				_ = _633_v40
				_633_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_620_v33).Select((Companion_Default___.SafeIndex(_628_cf4, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool), _630_v39)
				var _634_v41 *C0
				_ = _634_v41
				var _nw97 *C0 = New_C0_()
				_ = _nw97
				_nw97.Ctor__(p2, !(p2))
				_634_v41 = _nw97
				var _635_v42 _dafny.Sequence
				_ = _635_v42
				_635_v42 = _dafny.SeqOf(_634_v41, _634_v41, _634_v41)
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_630_v39), 0))
				_ = _index88
				(_630_v39).ArraySet1(_dafny.Companion_Sequence_.Equal(_635_v42, _635_v42), (_index88).Int())
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_630_v39), 0))
				_ = _index89
				var _rhs79 _dafny.Sequence = _622_v35
				_ = _rhs79
				var _rhs80 _dafny.Map = _633_v40
				_ = _rhs80
				var _rhs81 bool = (func() bool {
					if (_628_cf4).Cmp(_628_cf4) < 0 {
						return (_634_v41).F23()
					}
					return p2
				})()
				_ = _rhs81
				var _lhs69 _dafny.Array = _630_v39
				_ = _lhs69
				var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_630_v39), 0))
				_ = _lhs70
				_622_v35 = _rhs79
				_633_v40 = _rhs80
				(_lhs69).ArraySet1(_rhs81, (_lhs70).Int())
				var _636_v43 _dafny.Array
				_ = _636_v43
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(28))
				_ = _nw98
				_636_v43 = _nw98
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_636_v43), 0))
				_ = _index90
				(_636_v43).ArraySet1(_621_v34, (_index90).Int())
				var _637_v44 _dafny.Map
				_ = _637_v44
				_637_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_630_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_630_v39), 0))).Int()).(bool), _dafny.IntOfInt64(284))
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_636_v43), 0))
				_ = _index91
				var _rhs82 _dafny.CodePoint = Companion_Default___.Fm25(_628_cf4, (_dafny.Zero).Minus((_628_cf4).Plus(_628_cf4)), (_634_v41).F23(), _637_v44, globalState)
				_ = _rhs82
				var _rhs83 *C0 = _634_v41
				_ = _rhs83
				var _rhs84 _dafny.CodePoint = _617_v30
				_ = _rhs84
				var _rhs85 _dafny.Int = _dafny.IntOfInt64(948)
				_ = _rhs85
				var _rhs86 D2 = _621_v34
				_ = _rhs86
				var _lhs71 _dafny.Array = _636_v43
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_636_v43), 0))
				_ = _lhs72
				_617_v30 = _rhs82
				_634_v41 = _rhs83
				_617_v30 = _rhs84
				_627_cf5 = _rhs85
				(_lhs71).ArraySet1(_rhs86, (_lhs72).Int())
				var _638_v45 _dafny.Map
				_ = _638_v45
				_638_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_634_v41).F23())
				_638_v45 = (_638_v45).Update(p1, (_634_v41).F23())
				var _639_v46 _dafny.MultiSet
				_ = _639_v46
				_639_v46 = _dafny.MultiSetOf((_dafny.Zero).Minus(_615_v28))
				var _640_v47 _dafny.Map
				_ = _640_v47
				_640_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SetOf((_630_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(170), _dafny.ArrayLen((_630_v39), 0))).Int()).(bool)))
				var _rhs87 _dafny.MultiSet = _639_v46
				_ = _rhs87
				var _rhs88 _dafny.Array = _630_v39
				_ = _rhs88
				var _rhs89 _dafny.Sequence = _622_v35
				_ = _rhs89
				var _rhs90 _dafny.Map = _640_v47
				_ = _rhs90
				var _rhs91 _dafny.Int = _dafny.IntOfInt64(79)
				_ = _rhs91
				var _lhs73 *GlobalState = globalState
				_ = _lhs73
				_639_v46 = _rhs87
				_lhs73.F2 = _rhs88
				_622_v35 = _rhs89
				_640_v47 = _rhs90
				r3 = _rhs91
			} else if _source10.Is_DC5() {
				var _641___mcc_h2 bool = _source10.Get_().(D1_DC5).Cf6
				_ = _641___mcc_h2
				var _642___mcc_h3 bool = _source10.Get_().(D1_DC5).Cf7
				_ = _642___mcc_h3
				var _643_cf7 bool = _642___mcc_h3
				_ = _643_cf7
				var _644_cf6 bool = _641___mcc_h2
				_ = _644_cf6
				var _645_v48 D4
				_ = _645_v48
				_645_v48 = Companion_D4_.Create_DC14_(_643_cf7, false)
				var _646_v49 _dafny.Map
				_ = _646_v49
				_646_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_645_v48, (func() bool {
					if _644_cf6 {
						return _643_cf7
					}
					return true
				})())
				_646_v49 = (_646_v49).Update(_645_v48, (_620_v33).Select((Companion_Default___.SafeIndex(_615_v28, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool))
				var _647_v50 _dafny.Array
				_ = _647_v50
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw99
				_647_v50 = _nw99
				var _648_v51 _dafny.Array
				_ = _648_v51
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw100
				_648_v51 = _nw100
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))
				_ = _index92
				(_647_v50).ArraySet1(_648_v51, (_index92).Int())
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))
				_ = _index93
				(_647_v50).ArraySet1(_648_v51, (_index93).Int())
				var _arr0 _dafny.Array = _dafny.ArrayCastTo((_647_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))).Int()))
				_ = _arr0
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_dafny.ArrayCastTo((_647_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))).Int()))), 0))
				_ = _index94
				(_arr0).ArraySet1(p2, (_index94).Int())
				var _arr1 _dafny.Array = _dafny.ArrayCastTo((_647_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))).Int()))
				_ = _arr1
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen((_dafny.ArrayCastTo((_647_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_647_v50), 0))).Int()))), 0))
				_ = _index95
				(_arr1).ArraySet1(p2, (_index95).Int())
				var _649_v52 _dafny.Sequence
				_ = _649_v52
				_649_v52 = _dafny.SeqOf(_622_v35, _dafny.UnicodeSeqOfUtf8Bytes("tbocxqw"), _622_v35)
				var _650_v53 _dafny.Map
				_ = _650_v53
				_650_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_616_v29, (_649_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_615_v28, _615_v28, _615_v28)).Cardinality()), _dafny.IntOfUint32((_649_v52).Cardinality()))).Uint32()).(_dafny.Sequence))
				(globalState).F9 = !(_650_v53).Contains(_616_v29)
			} else if _source10.Is_DC3() {
				var _651___mcc_h4 _dafny.Int = _source10.Get_().(D1_DC3).Cf3
				_ = _651___mcc_h4
				var _652_cf3 _dafny.Int = _651___mcc_h4
				_ = _652_cf3
				var _rhs92 _dafny.Int = _615_v28
				_ = _rhs92
				var _rhs93 _dafny.Int = (Companion_Default___.SafeModuloInt(_652_cf3, _dafny.IntOfInt64(324))).Times(_dafny.IntOfUint32((_622_v35).Cardinality()))
				_ = _rhs93
				r3 = _rhs92
				r3 = _rhs93
				var _653_v54 _dafny.Array
				_ = _653_v54
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_10
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) bool = (func(_654_p2 bool) func(_dafny.Int) bool {
						return func(_655_i2 _dafny.Int) bool {
							return _654_p2
						}
					})(p2)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw101 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw101).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw101).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_653_v54 = _nw101
				var _656_v55 D4
				_ = _656_v55
				_656_v55 = Companion_D4_.Create_DC13_(_622_v35, _dafny.IntOfInt64(930), p2, _653_v54)
				(globalState).F2 = (_656_v55).Dtor_cf25()
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_653_v54), 0))
				_ = _index96
				(_653_v54).ArraySet1(p2, (_index96).Int())
				var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_653_v54), 0))
				_ = _index97
				(_653_v54).ArraySet1(p2, (_index97).Int())
				_615_v28 = _615_v28
			} else {
				var _657___mcc_h5 D1 = _source10.Get_().(D1_DC6).Cf8
				_ = _657___mcc_h5
				var _658_cf8 D1 = _657___mcc_h5
				_ = _658_cf8
				var _659_v56 _dafny.Map
				_ = _659_v56
				_659_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, true)
				var _660_v57 _dafny.Array
				_ = _660_v57
				var _nwElement0_17 bool = p2
				_ = _nwElement0_17
				var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(27))
				_ = _nw102
				(_nw102).ArraySet1(_nwElement0_17, 0)
				(_nw102).ArraySet1(p2, 1)
				(_nw102).ArraySet1(p2, 2)
				(_nw102).ArraySet1(p2, 3)
				(_nw102).ArraySet1(p2, 4)
				(_nw102).ArraySet1(p2, 5)
				(_nw102).ArraySet1(true, 6)
				(_nw102).ArraySet1(p2, 7)
				(_nw102).ArraySet1(p2, 8)
				(_nw102).ArraySet1(p2, 9)
				(_nw102).ArraySet1((func() bool {
					if (_659_v56).Contains(p2) {
						return (_659_v56).Get(p2).(bool)
					}
					return p2
				})(), 10)
				(_nw102).ArraySet1(p2, 11)
				(_nw102).ArraySet1(p2, 12)
				(_nw102).ArraySet1(p2, 13)
				(_nw102).ArraySet1(p2, 14)
				(_nw102).ArraySet1(p2, 15)
				(_nw102).ArraySet1(p2, 16)
				(_nw102).ArraySet1((_620_v33).Select((Companion_Default___.SafeIndex(_615_v28, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool), 17)
				(_nw102).ArraySet1(p2, 18)
				(_nw102).ArraySet1(p2, 19)
				(_nw102).ArraySet1(!(false), 20)
				(_nw102).ArraySet1(p2, 21)
				(_nw102).ArraySet1((_620_v33).Select((Companion_Default___.SafeIndex(_615_v28, _dafny.IntOfUint32((_620_v33).Cardinality()))).Uint32()).(bool), 22)
				(_nw102).ArraySet1(true, 23)
				(_nw102).ArraySet1(p2, 24)
				(_nw102).ArraySet1(p2, 25)
				(_nw102).ArraySet1(false, 26)
				_660_v57 = _nw102
				var _661_v58 D4
				_ = _661_v58
				_661_v58 = Companion_D4_.Create_DC13_(_622_v35, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_622_v35, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.IntOfUint32((_622_v35).Cardinality()))).Uint32(), _617_v30)).Cardinality()), !(false), _660_v57)
				var _662_v59 D4
				_ = _662_v59
				_662_v59 = Companion_D4_.Create_DC15_(_661_v58)
				var _663_v60 _dafny.Sequence
				_ = _663_v60
				_663_v60 = _dafny.SeqOf(_662_v59, _662_v59)
				r1 = (_615_v28).Minus((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_663_v60, _663_v60))).Cardinality())
				r2 = (_dafny.IntOfInt64(434)).Cmp(_dafny.IntOfInt64(22)) != 0
				var _664_v61 _dafny.MultiSet
				_ = _664_v61
				_664_v61 = _dafny.MultiSetOf(_615_v28)
				var _665_v62 D4
				_ = _665_v62
				_665_v62 = Companion_D4_.Create_DC13_(_622_v35, (_664_v61).Cardinality(), true, _660_v57)
				(globalState).F3 = (_665_v62).Dtor_cf23()
				var _666_v63 _dafny.Map
				_ = _666_v63
				_666_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v28, p2)
				var _667_v64 _dafny.Sequence
				_ = _667_v64
				_667_v64 = _dafny.SeqOf(p1)
				var _668_v65 _dafny.Map
				_ = _668_v65
				_668_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _615_v28)
				var _669_v66 _dafny.Map
				_ = _669_v66
				_669_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_666_v63, (_dafny.IntOfUint32((_667_v64).Cardinality())).Cmp((func() _dafny.Int {
					if (_668_v65).Contains(p2) {
						return (_668_v65).Get(p2).(_dafny.Int)
					}
					return (_668_v65).Cardinality()
				})()) != 0)
				_669_v66 = _669_v66
			}
			var _670_v67 D1
			_ = _670_v67
			_670_v67 = Companion_D1_.Create_DC5_(p2, p2)
			_670_v67 = _670_v67
			(globalState).F3 = _615_v28
			var _671_v68 *C0
			_ = _671_v68
			var _nw103 *C0 = New_C0_()
			_ = _nw103
			_nw103.Ctor__(p2, p2)
			_671_v68 = _nw103
		} else {
			var _672_v69 _dafny.Array
			_ = _672_v69
			var _nw104 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
			_ = _nw104
			_672_v69 = _nw104
			var _673_v70 _dafny.MultiSet
			_ = _673_v70
			_673_v70 = _dafny.MultiSetOf(_672_v69, _672_v69, _672_v69, _672_v69, _672_v69)
			_673_v70 = _dafny.MultiSetOf(_672_v69)
			var _674_v71 _dafny.MultiSet
			_ = _674_v71
			_674_v71 = _dafny.MultiSetOf(_615_v28)
			var _675_v72 _dafny.Map
			_ = _675_v72
			_675_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}((func(_676_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_677_i3 _dafny.Int) _dafny.CodePoint {
					return _676_v30
				}
			})(_617_v30))), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(p2, _615_v28, _615_v28, _674_v71, globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}((func(_678_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_679_i3 _dafny.Int) _dafny.CodePoint {
					return _678_v30
				}
			})(_617_v30)))).Cardinality()))).Uint32(), _617_v30))
			var _rhs94 _dafny.Map = ((_675_v72).Merge(_675_v72)).Merge(_675_v72)
			_ = _rhs94
			var _rhs95 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(59))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}((func(_680_v28 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_681_i4 _dafny.Int) _dafny.Int {
					return (Companion_D1_.Create_DC4_(_dafny.IntOfInt64(-838), _680_v28)).Dtor_cf4()
				}
			})(_615_v28)))
			_ = _rhs95
			_675_v72 = _rhs94
			_616_v29 = _rhs95
			var _682_v73 _dafny.Array
			_ = _682_v73
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_11
			var _nw105 _dafny.Array
			_ = _nw105
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw105 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Sequence = (func(_683_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_684_i5 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.CodePoint('h'))), _dafny.SeqOf(_dafny.SeqOf(_683_v30), _dafny.SeqOf(_683_v30, _683_v30)))
					}
				})(_617_v30)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw105 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw105).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw105).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_682_v73 = _nw105
			var _685_v74 _dafny.Sequence
			_ = _685_v74
			_685_v74 = _dafny.SeqOf(_617_v30, _617_v30)
			var _686_v75 _dafny.Sequence
			_ = _686_v75
			_686_v75 = _dafny.SeqOf(_685_v74, _dafny.SeqOf(_617_v30, _617_v30), _685_v74)
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_682_v73), 0))
			_ = _index98
			(_682_v73).ArraySet1(_686_v75, (_index98).Int())
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_682_v73), 0))
			_ = _index99
			(_682_v73).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_686_v75, _dafny.Companion_Sequence_.Concatenate(_686_v75, Companion_Default___.Fm26(_615_v28, p2, _dafny.IntOfUint32((_685_v74).Cardinality()), globalState))), (_index99).Int())
			var _687_v76 _dafny.Array
			_ = _687_v76
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_12
			var _nw106 _dafny.Array
			_ = _nw106
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw106 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Sequence = func(_688_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("tulofkoow")
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw106 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw106).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw106).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_687_v76 = _nw106
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_687_v76), 0))
			_ = _index100
			(_687_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(600))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg56 _dafny.Int) interface{} {
					return coer56(arg56)
				}
			}((func(_689_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_690_i7 _dafny.Int) _dafny.CodePoint {
					return _689_v30
				}
			})(_617_v30))), _dafny.Companion_Sequence_.Update(_685_v74, (Companion_Default___.SafeIndex(_615_v28, _dafny.IntOfUint32((_685_v74).Cardinality()))).Uint32(), _617_v30)), (_index100).Int())
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(573), _dafny.ArrayLen((_687_v76), 0))
			_ = _index101
			(_687_v76).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm27((_dafny.Zero).Minus(_615_v28), p2, globalState), _685_v74), (_index101).Int())
			var _691_v77 _dafny.Sequence
			_ = _691_v77
			_691_v77 = _dafny.SeqOf(true, p2, p2, p2, p2)
			var _692_v78 _dafny.Map
			_ = _692_v78
			_692_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_691_v77).Cardinality()), _615_v28)
			var _693_v79 _dafny.Map
			_ = _693_v79
			_693_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_692_v78, _615_v28)
			(globalState).F3 = ((func() _dafny.Map {
				if p2 {
					return _693_v79
				}
				return _693_v79
			})()).Cardinality()
		}
		var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))
		_ = _index102
		(p1).ArraySet1(_dafny.IntOfUint32((_616_v29).Cardinality()), (_index102).Int())
		var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))
		_ = _index103
		(p1).ArraySet1(_615_v28, (_index103).Int())
		r3 = (_615_v28).Times((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))
		var _hi1 _dafny.Int = _615_v28
		_ = _hi1
		for _694_i8 := (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int); _694_i8.Cmp(_hi1) < 0; _694_i8 = _694_i8.Plus(_dafny.One) {
			var _695_v80 *C0
			_ = _695_v80
			var _nw107 *C0 = New_C0_()
			_ = _nw107
			_nw107.Ctor__((func() bool {
				if p2 {
					return p2
				}
				return false
			})(), (p2) && (false))
			_695_v80 = _nw107
			if (_615_v28).Cmp(_694_i8) <= 0 {
				r2 = (_695_v80).F24()
				var _696_v81 _dafny.Sequence
				_ = _696_v81
				_696_v81 = _dafny.UnicodeSeqOfUtf8Bytes("lpi")
				_696_v81 = _696_v81
				_696_v81 = _696_v81
				var _697_v82 _dafny.Map
				_ = _697_v82
				_697_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_695_v80, Companion_Default___.Fm27((_dafny.Zero).Minus(_694_i8), p2, globalState))
				var _698_v83 _dafny.Array
				_ = _698_v83
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw108
				_698_v83 = _nw108
				var _699_v84 D4
				_ = _699_v84
				_699_v84 = Companion_D4_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(103))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_700_v30 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_701_i9 _dafny.Int) _dafny.CodePoint {
						return _700_v30
					}
				})(_617_v30))), _dafny.IntOfInt64(12), (_695_v80).F23(), _698_v83)
				var _702_v85 _dafny.Sequence
				_ = _702_v85
				_702_v85 = _dafny.SeqOf(_696_v81, _696_v81)
				var _703_v86 _dafny.Array
				_ = _703_v86
				var _nwElement0_18 _dafny.Sequence = _696_v81
				_ = _nwElement0_18
				var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(25))
				_ = _nw109
				(_nw109).ArraySet1(_nwElement0_18, 0)
				(_nw109).ArraySet1((func() _dafny.Sequence {
					if (_697_v82).Contains(_695_v80) {
						return (_697_v82).Get(_695_v80).(_dafny.Sequence)
					}
					return _696_v81
				})(), 1)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("pmsifhi"), _696_v81), 2)
				(_nw109).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sta"), 3)
				(_nw109).ArraySet1(_696_v81, 4)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_696_v81, Companion_Default___.Fm27(_dafny.IntOfUint32((_696_v81).Cardinality()), (_695_v80).F24(), globalState)), 5)
				(_nw109).ArraySet1(_696_v81, 6)
				(_nw109).ArraySet1(Companion_Default___.Fm27(_615_v28, (_695_v80).F23(), globalState), 7)
				(_nw109).ArraySet1(_696_v81, 8)
				(_nw109).ArraySet1(_696_v81, 9)
				(_nw109).ArraySet1(_696_v81, 10)
				(_nw109).ArraySet1((_699_v84).Dtor_cf22(), 11)
				(_nw109).ArraySet1(_696_v81, 12)
				(_nw109).ArraySet1(_696_v81, 13)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_696_v81, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(462))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}(func(_704_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				}))), 14)
				(_nw109).ArraySet1((_702_v85).Select((Companion_Default___.SafeIndex(_615_v28, _dafny.IntOfUint32((_702_v85).Cardinality()))).Uint32()).(_dafny.Sequence), 15)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_696_v81, (Companion_Default___.SafeIndex(_694_i8, _dafny.IntOfUint32((_696_v81).Cardinality()))).Uint32(), _617_v30), (Companion_Default___.SafeIndex((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_696_v81, (Companion_Default___.SafeIndex(_694_i8, _dafny.IntOfUint32((_696_v81).Cardinality()))).Uint32(), _617_v30)).Cardinality()))).Uint32(), _617_v30), 16)
				(_nw109).ArraySet1(_696_v81, 17)
				(_nw109).ArraySet1(_696_v81, 18)
				(_nw109).ArraySet1((_699_v84).Dtor_cf22(), 19)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Update(_696_v81, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_696_v81).Cardinality()), _dafny.IntOfUint32((_696_v81).Cardinality()))).Uint32(), _617_v30), 20)
				(_nw109).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_696_v81, _696_v81), 21)
				(_nw109).ArraySet1(_696_v81, 22)
				(_nw109).ArraySet1((func() _dafny.Sequence {
					if p2 {
						return _696_v81
					}
					return _696_v81
				})(), 23)
				(_nw109).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yhvna"), 24)
				_703_v86 = _nw109
				_703_v86 = _703_v86
				var _705_v87 _dafny.Array
				_ = _705_v87
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_13
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Sequence = (func(_706_v80 *C0) func(_dafny.Int) _dafny.Sequence {
						return func(_707_i11 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf((_706_v80).F24(), true, !(!((_706_v80).F23())))
						}
					})(_695_v80)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw110 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw110).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw110).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_705_v87 = _nw110
				var _708_v88 _dafny.Sequence
				_ = _708_v88
				_708_v88 = _dafny.SeqOf((_695_v80).F24(), false)
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_705_v87), 0))
				_ = _index104
				(_705_v87).ArraySet1(_708_v88, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(157), _dafny.ArrayLen((_705_v87), 0))
				_ = _index105
				(_705_v87).ArraySet1(_708_v88, (_index105).Int())
			} else {
				var _709_v89 _dafny.Sequence
				_ = _709_v89
				_709_v89 = _dafny.SeqOf(_617_v30)
				(globalState).F3 = Companion_Default___.Fm0(Companion_Default___.Fm23(_694_i8, (_695_v80).F23(), _618_v31, _dafny.SetOf(_616_v29), globalState), (_dafny.IntOfUint32((_709_v89).Cardinality())).Plus(_615_v28), (_dafny.SetOf((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))).Cardinality(), _dafny.MultiSetOf(_694_i8), globalState)
				(globalState).F9 = false
				(globalState).F9 = (_615_v28).Cmp(_615_v28) > 0
				var _710_v90 *C0
				_ = _710_v90
				var _nw111 *C0 = New_C0_()
				_ = _nw111
				_nw111.Ctor__((_695_v80).F23(), (_695_v80).F24())
				_710_v90 = _nw111
				var _711_v91 _dafny.Array
				_ = _711_v91
				var _nw112 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
				_ = _nw112
				_711_v91 = _nw112
				var _712_v92 _dafny.MultiSet
				_ = _712_v92
				_712_v92 = _dafny.MultiSetOf(_711_v91, _711_v91)
				(globalState).F9 = !((((_712_v92).Update(_711_v91, Companion_Default___.Abs(_dafny.IntOfInt64(911)))).Difference(_712_v92)).Equals(_712_v92))
			}
			r3 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
			var _713_v93 D5
			_ = _713_v93
			_713_v93 = Companion_D5_.Create_DC18_((_695_v80).F24())
			var _714_v94 _dafny.Map
			_ = _714_v94
			_714_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_713_v93, _615_v28)
			(globalState).F9 = (_615_v28).Cmp((_615_v28).Times((_714_v94).Cardinality())) > 0
		}
		var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))
		_ = _index106
		(p1).ArraySet1((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (_index106).Int())
		r0 = Companion_Default___.Fm28(p2, globalState)
		r1 = (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int)
		var _715_v95 _dafny.Sequence
		_ = _715_v95
		_715_v95 = _dafny.SeqOf(!(p2))
		var _716_v96 _dafny.Sequence
		_ = _716_v96
		_716_v96 = _dafny.SeqOf(_715_v95, _dafny.Companion_Sequence_.Update(_715_v95, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), (p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int))).Cardinality(), _dafny.IntOfUint32((_715_v95).Cardinality()))).Uint32(), p2))
		var _717_v97 _dafny.MultiSet
		_ = _717_v97
		_717_v97 = _dafny.MultiSetOf((_dafny.Zero).Minus(_615_v28))
		var _pat_let_tv19 = p2
		_ = _pat_let_tv19
		var _718_v98 _dafny.Sequence
		_ = _718_v98
		_718_v98 = _dafny.SeqOf(func(_pat_let19_0 D5) D5 {
			return func(_719_dt__update__tmp_h0 D5) D5 {
				return func(_pat_let20_0 bool) D5 {
					return func(_720_dt__update_hcf31_h0 bool) D5 {
						return Companion_D5_.Create_DC17_((_719_dt__update__tmp_h0).Dtor_cf30(), _720_dt__update_hcf31_h0, (_719_dt__update__tmp_h0).Dtor_cf32())
					}(_pat_let20_0)
				}(_pat_let_tv19)
			}(_pat_let19_0)
		}(Companion_D5_.Create_DC17_(_716_v96, p2, _717_v97)))
		r2 = Companion_Default___.Fm23((p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(15), _dafny.ArrayLen((p1), 0))).Int()).(_dafny.Int), p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_617_v30, _dafny.IntOfUint32((_718_v98).Cardinality()))).Merge(_618_v31), _619_v32, globalState)
		r3 = _615_v28
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F22 _dafny.Map
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F22 = _dafny.EmptyMap
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f22 _dafny.Map) {
	{
		(_this).F22 = f22
	}
}
func (_this *C2) M8(p0 _dafny.Array, globalState *GlobalState) (_dafny.CodePoint, _dafny.Int, D3) {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 D3 = Companion_D3_.Default()
		_ = r2
		var _721_v0 bool
		_ = _721_v0
		_721_v0 = true
		var _722_i0 _dafny.Int
		_ = _722_i0
		_722_i0 = _dafny.Zero
		{
			for _721_v0 {
				{
					if (_722_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_722_i0 = (_722_i0).Plus(_dafny.One)
					(globalState).F9 = _721_v0
					var _723_v1 _dafny.CodePoint
					_ = _723_v1
					_723_v1 = _dafny.CodePoint('y')
					r0 = _723_v1
					var _724_v2 D0
					_ = _724_v2
					_724_v2 = Companion_D0_.Create_DC1_(_721_v0)
					var _source11 D0 = _724_v2
					_ = _source11
					if _source11.Is_DC1() {
						var _725___mcc_h0 bool = _source11.Get_().(D0_DC1).Cf1
						_ = _725___mcc_h0
						var _726_cf1 bool = _725___mcc_h0
						_ = _726_cf1
						var _727_v3 _dafny.Int
						_ = _727_v3
						_727_v3 = _dafny.IntOfInt64(8)
						(globalState).F3 = _727_v3
						var _728_v4 _dafny.Array
						_ = _728_v4
						var _len0_14 _dafny.Int = _dafny.IntOfInt64(24)
						_ = _len0_14
						var _nw113 _dafny.Array
						_ = _nw113
						if _len0_14.Cmp(_dafny.Zero) == 0 {
							_nw113 = _dafny.NewArray(_len0_14)
						} else {
							var _init14 func(_dafny.Int) _dafny.Sequence = (func(_729_v3 _dafny.Int, _730_cf1 bool, _731_v0 bool) func(_dafny.Int) _dafny.Sequence {
								return func(_732_i1 _dafny.Int) _dafny.Sequence {
									return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(981))).Uint32(), func(coer59 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
										return func(arg59 _dafny.Int) interface{} {
											return coer59(arg59)
										}
									}((func(_733_v3 _dafny.Int) func(_dafny.Int) D1 {
										return func(_734_i2 _dafny.Int) D1 {
											return Companion_D1_.Create_DC4_(_733_v3, _733_v3)
										}
									})(_729_v3))), _dafny.SeqOf(Companion_D1_.Create_DC4_((_this.F22).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_729_v3, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_730_cf1, _731_v0)).Cardinality())).Cardinality())))
								}
							})(_727_v3, _726_cf1, _721_v0)
							_ = _init14
							var _element0_14 = _init14(_dafny.Zero)
							_ = _element0_14
							_nw113 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
							(_nw113).ArraySet1(_element0_14, 0)
							var _nativeLen0_14 = (_len0_14).Int()
							_ = _nativeLen0_14
							for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
								(_nw113).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
							}
						}
						_728_v4 = _nw113
						var _735_v5 _dafny.Map
						_ = _735_v5
						_735_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_726_cf1, _726_cf1)
						var _736_v6 _dafny.Map
						_ = _736_v6
						_736_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v1, (_735_v5).Cardinality())
						var _737_v7 _dafny.Sequence
						_ = _737_v7
						_737_v7 = _dafny.UnicodeSeqOfUtf8Bytes("pr")
						var _738_v8 _dafny.Sequence
						_ = _738_v8
						_738_v8 = _dafny.SeqOf((func() _dafny.Int {
							if (_736_v6).Contains(_723_v1) {
								return (_736_v6).Get(_723_v1).(_dafny.Int)
							}
							return _727_v3
						})(), _727_v3, _dafny.IntOfUint32((_737_v7).Cardinality()))
						var _739_v9 _dafny.Array
						_ = _739_v9
						var _nw114 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
						_ = _nw114
						_739_v9 = _nw114
						var _740_v10 _dafny.Map
						_ = _740_v10
						_740_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v9, _723_v1)
						var _741_v11 _dafny.Sequence
						_ = _741_v11
						_741_v11 = _dafny.SeqOf(_721_v0, _726_cf1)
						var _742_v12 _dafny.Sequence
						_ = _742_v12
						_742_v12 = _dafny.SeqOf(_741_v11, _741_v11)
						var _743_v13 _dafny.Array
						_ = _743_v13
						var _nwElement0_19 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cudqbqnau")).Cardinality())
						_ = _nwElement0_19
						var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(8))
						_ = _nw115
						(_nw115).ArraySet1(_nwElement0_19, 0)
						(_nw115).ArraySet1(_dafny.IntOfInt64(-98), 1)
						(_nw115).ArraySet1((_727_v3).Plus(_727_v3), 2)
						(_nw115).ArraySet1(_727_v3, 3)
						(_nw115).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if false {
								return _dafny.SeqOf(_727_v3)
							}
							return _dafny.Companion_Sequence_.Update(_738_v8, (Companion_Default___.SafeIndex(_727_v3, _dafny.IntOfUint32((_738_v8).Cardinality()))).Uint32(), (_740_v10).Cardinality())
						})()).Cardinality()), 4)
						(_nw115).ArraySet1((_727_v3).Minus(_dafny.IntOfUint32((_738_v8).Cardinality())), 5)
						(_nw115).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_726_cf1, _727_v3, _dafny.IntOfInt64(441), _dafny.MultiSetFromSeq(_738_v8), globalState), (_dafny.Zero).Minus(_dafny.IntOfUint32(((_742_v12).Select((Companion_Default___.SafeIndex(_727_v3, _dafny.IntOfUint32((_742_v12).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))), 6)
						(_nw115).ArraySet1((_dafny.IntOfUint32((_738_v8).Cardinality())).Plus(_dafny.IntOfInt64(558)), 7)
						_743_v13 = _nw115
						var _744_v14 _dafny.Array
						_ = _744_v14
						var _nwElement0_20 D0 = _724_v2
						_ = _nwElement0_20
						var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(13))
						_ = _nw116
						(_nw116).ArraySet1(_nwElement0_20, 0)
						(_nw116).ArraySet1(_724_v2, 1)
						(_nw116).ArraySet1(_724_v2, 2)
						(_nw116).ArraySet1(Companion_D0_.Create_DC1_((func() bool {
							if (_this.F22).Contains(_727_v3) {
								return (_this.F22).Get(_727_v3).(bool)
							}
							return false
						})()), 3)
						(_nw116).ArraySet1(_724_v2, 4)
						(_nw116).ArraySet1(_724_v2, 5)
						(_nw116).ArraySet1(_724_v2, 6)
						(_nw116).ArraySet1(_724_v2, 7)
						(_nw116).ArraySet1(_724_v2, 8)
						(_nw116).ArraySet1(_724_v2, 9)
						(_nw116).ArraySet1(_724_v2, 10)
						(_nw116).ArraySet1(_724_v2, 11)
						(_nw116).ArraySet1(_724_v2, 12)
						_744_v14 = _nw116
						var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_744_v14), 0))
						_ = _index107
						(_744_v14).ArraySet1(_724_v2, (_index107).Int())
						var _745_v15 _dafny.Sequence
						_ = _745_v15
						_745_v15 = _dafny.SeqOf(_728_v4)
						var _746_v16 _dafny.Map
						_ = _746_v16
						_746_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_727_v3, p0)
						var _747_v17 _dafny.MultiSet
						_ = _747_v17
						_747_v17 = _dafny.MultiSetOf((_this.F22).Cardinality())
						var _pat_let_tv20 = _726_cf1
						_ = _pat_let_tv20
						var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_744_v14), 0))
						_ = _index108
						var _rhs96 _dafny.Array = (_745_v15).Select((Companion_Default___.SafeIndex((_727_v3).Times((_746_v16).Cardinality()), _dafny.IntOfUint32((_745_v15).Cardinality()))).Uint32()).(_dafny.Array)
						_ = _rhs96
						var _rhs97 _dafny.Int = Companion_Default___.Fm0(false, (_727_v3).Plus(_dafny.IntOfInt64(358)), _727_v3, _747_v17, globalState)
						_ = _rhs97
						var _rhs98 _dafny.Array = _743_v13
						_ = _rhs98
						var _rhs99 _dafny.Array = (Companion_D4_.Create_DC12_(_739_v9)).Dtor_cf21()
						_ = _rhs99
						var _rhs100 D0 = func(_pat_let21_0 D0) D0 {
							return func(_748_dt__update__tmp_h0 D0) D0 {
								return func(_pat_let22_0 bool) D0 {
									return func(_749_dt__update_hcf1_h0 bool) D0 {
										return Companion_D0_.Create_DC1_(_749_dt__update_hcf1_h0)
									}(_pat_let22_0)
								}(_pat_let_tv20)
							}(_pat_let21_0)
						}(_724_v2)
						_ = _rhs100
						var _lhs74 _dafny.Array = _744_v14
						_ = _lhs74
						var _lhs75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(203), _dafny.ArrayLen((_744_v14), 0))
						_ = _lhs75
						_728_v4 = _rhs96
						_727_v3 = _rhs97
						_743_v13 = _rhs98
						_739_v9 = _rhs99
						(_lhs74).ArraySet1(_rhs100, (_lhs75).Int())
						var _750_v18 *C1
						_ = _750_v18
						var _nw117 *C1 = New_C1_()
						_ = _nw117
						_nw117.Ctor__()
						_750_v18 = _nw117
						var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((p0), 0))
						_ = _index109
						(p0).ArraySet1(_727_v3, (_index109).Int())
						var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(418), _dafny.ArrayLen((p0), 0))
						_ = _index110
						(p0).ArraySet1(((_727_v3).Times(_727_v3)).Times((_727_v3).Minus(_727_v3)), (_index110).Int())
					} else if _source11.Is_DC0() {
						var _751___mcc_h1 bool = _source11.Get_().(D0_DC0).Cf0
						_ = _751___mcc_h1
						var _752_cf0 bool = _751___mcc_h1
						_ = _752_cf0
						var _753_v19 D0
						_ = _753_v19
						_753_v19 = Companion_D0_.Create_DC0_(_721_v0)
						_753_v19 = _753_v19
						var _754_v20 _dafny.Map
						_ = _754_v20
						_754_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_752_cf0) || (_721_v0), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(529))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg60 _dafny.Int) interface{} {
								return coer60(arg60)
							}
						}((func(_755_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_756_i3 _dafny.Int) _dafny.CodePoint {
								return _755_v1
							}
						})(_723_v1))))
						var _757_v21 _dafny.Sequence
						_ = _757_v21
						_757_v21 = _dafny.UnicodeSeqOfUtf8Bytes("lksdhhcbj")
						_754_v20 = (_754_v20).Update(_752_cf0, _757_v21)
						var _758_v23 _dafny.Int
						_ = _758_v23
						var _out42 _dafny.Int
						_ = _out42
						_out42 = (_this).M9((func() _dafny.Set {
							var _coll38 = _dafny.NewBuilder()
							_ = _coll38
							for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(502), _dafny.IntOfInt64(693))); ; {
								_compr_38, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _759_v22 _dafny.Int
								_759_v22 = interface{}(_compr_38).(_dafny.Int)
								if ((_dafny.IntOfInt64(502)).Cmp(_759_v22) <= 0) && ((_759_v22).Cmp(_dafny.IntOfInt64(693)) < 0) {
									_coll38.Add(Companion_Default___.SafeDivisionInt(_759_v22, _dafny.IntOfUint32((_757_v21).Cardinality())))
								}
							}
							return _coll38.ToSet()
						}()).Cardinality(), globalState)
						_758_v23 = _out42
						var _760_v24 _dafny.Array
						_ = _760_v24
						var _nwElement0_21 bool = _752_cf0
						_ = _nwElement0_21
						var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(3))
						_ = _nw118
						(_nw118).ArraySet1(_nwElement0_21, 0)
						(_nw118).ArraySet1(_752_cf0, 1)
						(_nw118).ArraySet1(_752_cf0, 2)
						_760_v24 = _nw118
						var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_760_v24), 0))
						_ = _index111
						(_760_v24).ArraySet1(_721_v0, (_index111).Int())
						var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen((_760_v24), 0))
						_ = _index112
						(_760_v24).ArraySet1(_721_v0, (_index112).Int())
					} else {
						var _761___mcc_h2 D0 = _source11.Get_().(D0_DC2).Cf2
						_ = _761___mcc_h2
						var _762_cf2 D0 = _761___mcc_h2
						_ = _762_cf2
						var _763_v25 _dafny.Int
						_ = _763_v25
						_763_v25 = _dafny.IntOfInt64(921)
						r1 = _763_v25
						var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((p0), 0))
						_ = _index113
						(p0).ArraySet1(_dafny.IntOfInt64(887), (_index113).Int())
						var _764_v26 _dafny.Sequence
						_ = _764_v26
						_764_v26 = _dafny.UnicodeSeqOfUtf8Bytes("wfyhouxd")
						var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((p0), 0))
						_ = _index114
						(p0).ArraySet1(Companion_Default___.SafeDivisionInt(_763_v25, (_dafny.IntOfInt64(-526)).Plus(_dafny.IntOfUint32((_764_v26).Cardinality()))), (_index114).Int())
						var _765_v27 _dafny.Array
						_ = _765_v27
						var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
						_ = _nw119
						_765_v27 = _nw119
						_765_v27 = _765_v27
						var _766_v28 _dafny.Array
						_ = _766_v28
						var _nw120 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(14))
						_ = _nw120
						_766_v28 = _nw120
						var _767_v29 _dafny.Map
						_ = _767_v29
						_767_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_763_v25, _766_v28)
						var _768_v30 _dafny.Array
						_ = _768_v30
						var _nwElement0_22 _dafny.Array = _766_v28
						_ = _nwElement0_22
						var _nw121 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(10))
						_ = _nw121
						(_nw121).ArraySet1(_nwElement0_22, 0)
						(_nw121).ArraySet1(_766_v28, 1)
						(_nw121).ArraySet1(_766_v28, 2)
						(_nw121).ArraySet1(_766_v28, 3)
						(_nw121).ArraySet1(_766_v28, 4)
						(_nw121).ArraySet1((func() _dafny.Array {
							if true {
								return _766_v28
							}
							return _766_v28
						})(), 5)
						(_nw121).ArraySet1(_766_v28, 6)
						(_nw121).ArraySet1(_766_v28, 7)
						(_nw121).ArraySet1((func() _dafny.Array {
							if (_767_v29).Contains((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)) {
								return (_767_v29).Get((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int)).(_dafny.Array)
							}
							return _766_v28
						})(), 8)
						(_nw121).ArraySet1(_766_v28, 9)
						_768_v30 = _nw121
						var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_768_v30), 0))
						_ = _index115
						(_768_v30).ArraySet1(_766_v28, (_index115).Int())
						var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(97), _dafny.ArrayLen((_768_v30), 0))
						_ = _index116
						(_768_v30).ArraySet1(_766_v28, (_index116).Int())
					}
					(globalState).F9 = (_721_v0) && (_721_v0)
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _769_v31 _dafny.Int
		_ = _769_v31
		_769_v31 = _dafny.IntOfInt64(86)
		var _770_v32 _dafny.Map
		_ = _770_v32
		_770_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_769_v31).Cmp(_dafny.IntOfInt64(-586)) != 0, true)
		var _771_v33 _dafny.CodePoint
		_ = _771_v33
		_771_v33 = _dafny.CodePoint('f')
		var _772_v34 _dafny.Map
		_ = _772_v34
		_772_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_771_v33, _769_v31)
		var _773_v35 _dafny.Set
		_ = _773_v35
		_773_v35 = _dafny.SetOf(_dafny.SeqOf(_769_v31, _769_v31))
		_770_v32 = (_770_v32).Update(_721_v0, Companion_Default___.Fm23(_769_v31, true, _772_v34, _773_v35, globalState))
		r0 = _771_v33
		var _774_v36 *C1
		_ = _774_v36
		var _nw122 *C1 = New_C1_()
		_ = _nw122
		_nw122.Ctor__()
		_774_v36 = _nw122
		var _775_i4 _dafny.Int
		_ = _775_i4
		_775_i4 = _dafny.Zero
		{
			for _721_v0 {
				{
					if (_775_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_775_i4 = (_775_i4).Plus(_dafny.One)
					var _776_v37 _dafny.Set
					_ = _776_v37
					_776_v37 = _dafny.SetOf(_dafny.IntOfInt64(365), _769_v31, _769_v31)
					var _777_v38 _dafny.Sequence
					_ = _777_v38
					_777_v38 = _dafny.SeqOf(_776_v37, _776_v37, _776_v37)
					var _778_v39 _dafny.Map
					_ = _778_v39
					_778_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_769_v31, (_777_v38).Select((Companion_Default___.SafeIndex(_769_v31, _dafny.IntOfUint32((_777_v38).Cardinality()))).Uint32()).(_dafny.Set))
					_778_v39 = (_778_v39).Update(_769_v31, _776_v37)
					var _779_v40 _dafny.Sequence
					_ = _779_v40
					_779_v40 = _dafny.UnicodeSeqOfUtf8Bytes("ncnudk")
					var _780_v41 _dafny.MultiSet
					_ = _780_v41
					_780_v41 = _dafny.MultiSetOf(_dafny.IntOfUint32((_779_v40).Cardinality()), _769_v31, _769_v31, _769_v31, _769_v31)
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((p0), 0))
					_ = _index117
					(p0).ArraySet1((_780_v41).Cardinality(), (_index117).Int())
					var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((p0), 0))
					_ = _index118
					(p0).ArraySet1((_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(802))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}((func(_781_v31 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_782_i5 _dafny.Int) _dafny.Int {
							return _781_v31
						}
					})(_769_v31)))).Cardinality())).Plus(Companion_Default___.Fm0(_721_v0, _769_v31, _769_v31, _780_v41, globalState)), (_index118).Int())
					var _783_v42 *C0
					_ = _783_v42
					var _nw123 *C0 = New_C0_()
					_ = _nw123
					_nw123.Ctor__(_721_v0, _721_v0)
					_783_v42 = _nw123
					var _784_v43 _dafny.Map
					_ = _784_v43
					_784_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int), _769_v31)
					r1 = (_dafny.IntOfUint32((_dafny.SeqOf((_784_v43).Cardinality())).Cardinality())).Plus((p0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(215), _dafny.ArrayLen((p0), 0))).Int()).(_dafny.Int))
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _785_v44 _dafny.MultiSet
		_ = _785_v44
		_785_v44 = _dafny.MultiSetOf(!(_721_v0))
		var _786_v45 _dafny.Sequence
		_ = _786_v45
		_786_v45 = _dafny.SeqOf(_769_v31, _769_v31, _769_v31, _769_v31, _dafny.IntOfInt64(494))
		(globalState).F3 = Companion_Default___.SafeModuloInt((_785_v44).Cardinality(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _721_v0 {
				return _786_v45
			}
			return _786_v45
		})()).Cardinality()))
		r0 = _771_v33
		r1 = _769_v31
		var _787_v46 _dafny.Sequence
		_ = _787_v46
		_787_v46 = _dafny.SeqOf(_772_v34, _772_v34, _772_v34, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.IntOfInt64(701)), _772_v34)
		var _788_v47 _dafny.Sequence
		_ = _788_v47
		_788_v47 = _dafny.UnicodeSeqOfUtf8Bytes("ymak")
		var _789_v48 D3
		_ = _789_v48
		_789_v48 = Companion_D3_.Create_DC10_(_770_v32)
		r2 = (func() D3 {
			if Companion_Default___.Fm23(_769_v31, _721_v0, (_787_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_788_v47).Cardinality()), _dafny.IntOfUint32((_787_v46).Cardinality()))).Uint32()).(_dafny.Map), _773_v35, globalState) {
				return _789_v48
			}
			return _789_v48
		})()
		return r0, r1, r2
	}
}
func (_this *C2) M9(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _790_v0 bool
		_ = _790_v0
		_790_v0 = true
		var _791_v1 _dafny.MultiSet
		_ = _791_v1
		_791_v1 = _dafny.MultiSetOf(p0, p0, _dafny.IntOfInt64(917), p0)
		var _792_v2 _dafny.Sequence
		_ = _792_v2
		_792_v2 = _dafny.SeqOf(_791_v1)
		var _793_v3 _dafny.Sequence
		_ = _793_v3
		_793_v3 = _dafny.UnicodeSeqOfUtf8Bytes("p")
		var _794_v4 D5
		_ = _794_v4
		_794_v4 = Companion_D5_.Create_DC19_(Companion_D5_.Create_DC17_(Companion_Default___.Fm29(globalState), _790_v0, (_792_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_793_v3).Cardinality())), _dafny.IntOfUint32((_792_v2).Cardinality()))).Uint32()).(_dafny.MultiSet)))
		var _795_v5 _dafny.Sequence
		_ = _795_v5
		_795_v5 = _dafny.SeqOf(_794_v4)
		var _796_v6 _dafny.Sequence
		_ = _796_v6
		_796_v6 = _dafny.SeqOf(_790_v0)
		var _797_v7 D2
		_ = _797_v7
		_797_v7 = Companion_D2_.Create_DC7_(_796_v6)
		var _798_v8 _dafny.Array
		_ = _798_v8
		var _nwElement0_23 _dafny.Int = p0
		_ = _nwElement0_23
		var _nw124 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(17))
		_ = _nw124
		(_nw124).ArraySet1(_nwElement0_23, 0)
		(_nw124).ArraySet1(p0, 1)
		(_nw124).ArraySet1((_dafny.Zero).Minus(p0), 2)
		(_nw124).ArraySet1(p0, 3)
		(_nw124).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_794_v4), _795_v5)).Cardinality()), 4)
		(_nw124).ArraySet1(p0, 5)
		(_nw124).ArraySet1(p0, 6)
		(_nw124).ArraySet1((_dafny.Zero).Minus(p0), 7)
		(_nw124).ArraySet1(p0, 8)
		(_nw124).ArraySet1(p0, 9)
		(_nw124).ArraySet1((p0).Plus(p0), 10)
		(_nw124).ArraySet1((_dafny.IntOfUint32(((_797_v7).Dtor_cf9()).Cardinality())).Plus(p0), 11)
		(_nw124).ArraySet1(p0, 12)
		(_nw124).ArraySet1(p0, 13)
		(_nw124).ArraySet1(p0, 14)
		(_nw124).ArraySet1(p0, 15)
		(_nw124).ArraySet1(p0, 16)
		_798_v8 = _nw124
		var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _index119
		(_798_v8).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), (_index119).Int())
		var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _index120
		(_798_v8).ArraySet1(_dafny.IntOfInt64(-989), (_index120).Int())
		var _799_v9 _dafny.MultiSet
		_ = _799_v9
		_799_v9 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("lpyalxhrf"), _793_v3)
		(globalState).F9 = (_796_v6).Select((Companion_Default___.SafeIndex(((_799_v9).Union(_799_v9)).Cardinality(), _dafny.IntOfUint32((_796_v6).Cardinality()))).Uint32()).(bool)
		var _800_v10 *C1
		_ = _800_v10
		var _nw125 *C1 = New_C1_()
		_ = _nw125
		_nw125.Ctor__()
		_800_v10 = _nw125
		var _801_v11 _dafny.Array
		_ = _801_v11
		var _len0_15 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_15
		var _nw126 _dafny.Array
		_ = _nw126
		if _len0_15.Cmp(_dafny.Zero) == 0 {
			_nw126 = _dafny.NewArray(_len0_15)
		} else {
			var _init15 func(_dafny.Int) bool = func(_802_i0 _dafny.Int) bool {
				return (Companion_D5_.Create_DC18_(true)).Dtor_cf33()
			}
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
		_801_v11 = _nw126
		var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
		_ = _index121
		(_801_v11).ArraySet1(_790_v0, (_index121).Int())
		var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_801_v11), 0))
		_ = _index122
		(_801_v11).ArraySet1(!(_790_v0), (_index122).Int())
		var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
		_ = _index123
		var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _index124
		var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_801_v11), 0))
		_ = _index125
		var _rhs101 bool = _790_v0
		_ = _rhs101
		var _rhs102 bool = !((Companion_Default___.SafeDivisionInt(p0, (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))).Cmp((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)) <= 0)
		_ = _rhs102
		var _rhs103 _dafny.Int = Companion_Default___.SafeModuloInt((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0(_790_v0, _dafny.IntOfInt64(177), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _791_v1, globalState))
		_ = _rhs103
		var _rhs104 bool = _790_v0
		_ = _rhs104
		var _lhs76 _dafny.Array = _801_v11
		_ = _lhs76
		var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
		_ = _lhs77
		var _lhs78 _dafny.Array = _798_v8
		_ = _lhs78
		var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _lhs79
		var _lhs80 _dafny.Array = _801_v11
		_ = _lhs80
		var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(512), _dafny.ArrayLen((_801_v11), 0))
		_ = _lhs81
		(_lhs76).ArraySet1(_rhs101, (_lhs77).Int())
		_790_v0 = _rhs102
		(_lhs78).ArraySet1(_rhs103, (_lhs79).Int())
		(_lhs80).ArraySet1(_rhs104, (_lhs81).Int())
		var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
		_ = _index126
		var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _index127
		var _rhs105 bool = (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)
		_ = _rhs105
		var _rhs106 _dafny.Int = _dafny.IntOfInt64(174)
		_ = _rhs106
		var _lhs82 _dafny.Array = _801_v11
		_ = _lhs82
		var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
		_ = _lhs83
		var _lhs84 _dafny.Array = _798_v8
		_ = _lhs84
		var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
		_ = _lhs85
		(_lhs82).ArraySet1(_rhs105, (_lhs83).Int())
		(_lhs84).ArraySet1(_rhs106, (_lhs85).Int())
		var _803_v12 _dafny.CodePoint
		_ = _803_v12
		_803_v12 = _dafny.CodePoint('q')
		var _804_v13 _dafny.Map
		_ = _804_v13
		_804_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_790_v0), (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool))
		var _805_v14 _dafny.Sequence
		_ = _805_v14
		_805_v14 = _dafny.SeqOf(p0, (_804_v13).Cardinality(), p0, (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))
		var _806_v15 _dafny.Set
		_ = _806_v15
		_806_v15 = _dafny.SetOf(_805_v14)
		var _807_v16 _dafny.Set
		_ = _807_v16
		_807_v16 = _dafny.SetOf(_790_v0, Companion_Default___.Fm23((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), false, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_803_v12, (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), _806_v15, globalState))
		if (_807_v16).IsProperSubsetOf((_807_v16).Intersection(_807_v16)) {
			if (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool) {
				var _808_v17 _dafny.Map
				_ = _808_v17
				_808_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_790_v0, p0)
				var _809_v18 _dafny.Map
				_ = _809_v18
				_809_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_793_v3, _808_v17)
				var _810_v19 D3
				_ = _810_v19
				_810_v19 = Companion_D3_.Create_DC11_(p0, (_804_v13).Cardinality(), _790_v0, (_805_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_805_v14).Cardinality()))).Uint32()).(_dafny.Int))
				_809_v18 = (_809_v18).Update((func() _dafny.Sequence {
					if _790_v0 {
						return _dafny.UnicodeSeqOfUtf8Bytes("byt")
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("e")
				})(), Companion_Default___.Fm1((_810_v19).Dtor_cf19(), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), !(_790_v0), Companion_Default___.Fm22(p0, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _790_v0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool))), _790_v0, !((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)), globalState), globalState))
				var _811_v20 _dafny.MultiSet
				_ = _811_v20
				_811_v20 = _dafny.MultiSetOf((_808_v17).Update((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_790_v0, p0))
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
				_ = _index128
				(_798_v8).ArraySet1((_811_v20).Cardinality(), (_index128).Int())
				(globalState).F3 = p0
				(globalState).F9 = !(_790_v0)
				(globalState).F3 = Companion_Default___.Fm0(true, _dafny.IntOfInt64(312), p0, _791_v1, globalState)
			} else {
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
				_ = _index129
				(_801_v11).ArraySet1((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), (_index129).Int())
				r0 = (_805_v14).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(726), _dafny.IntOfUint32((_805_v14).Cardinality()))).Uint32()).(_dafny.Int)
				var _812_v21 D1
				_ = _812_v21
				_812_v21 = Companion_D1_.Create_DC5_(!_dafny.Companion_Sequence_.Contains(_805_v14, p0), false)
				var _813_v22 _dafny.Sequence
				_ = _813_v22
				_813_v22 = _dafny.SeqOf(_796_v6)
				var _814_v23 D5
				_ = _814_v23
				_814_v23 = Companion_D5_.Create_DC17_(_813_v22, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _791_v1)
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
				_ = _index130
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
				_ = _index131
				var _rhs107 bool = (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)
				_ = _rhs107
				var _rhs108 D1 = _812_v21
				_ = _rhs108
				var _rhs109 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), p0)).Minus((_dafny.Zero).Minus(((_791_v1).Cardinality()).Times((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)))))
				_ = _rhs109
				var _rhs110 bool = true
				_ = _rhs110
				var _rhs111 bool = !(((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)) || (!((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)))) || ((_814_v23).Dtor_cf31())
				_ = _rhs111
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 _dafny.Array = _798_v8
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
				_ = _lhs88
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 _dafny.Array = _801_v11
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
				_ = _lhs91
				_lhs86.F9 = _rhs107
				_812_v21 = _rhs108
				(_lhs87).ArraySet1(_rhs109, (_lhs88).Int())
				_lhs89.F9 = _rhs110
				(_lhs90).ArraySet1(_rhs111, (_lhs91).Int())
				(globalState).F3 = _dafny.IntOfInt64(31)
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
				_ = _index132
				(_801_v11).ArraySet1((_dafny.IntOfInt64(882)).Cmp((p0).Times((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))) == 0, (_index132).Int())
			}
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
			_ = _index133
			(_798_v8).ArraySet1((p0).Times(((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)).Times(p0)), (_index133).Int())
			if (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool) {
				var _815_v24 _dafny.Map
				_ = _815_v24
				_815_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), p0)
				_815_v24 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))).Update((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(p0))).Merge(_815_v24)
				(globalState).F9 = (p0).Cmp(p0) >= 0
				var _816_v25 D4
				_ = _816_v25
				_816_v25 = Companion_D4_.Create_DC14_(_790_v0, (_dafny.IntOfInt64(446)).Cmp(p0) <= 0)
				_816_v25 = _816_v25
				_790_v0 = _790_v0
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
				_ = _index134
				(_798_v8).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool) {
						return (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)
					}
					return Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_793_v3).Cardinality()))
				})()), (_index134).Int())
			} else {
				(globalState).F9 = !(!(_790_v0))
				var _817_v26 _dafny.Array
				_ = _817_v26
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_16
				var _nw127 _dafny.Array
				_ = _nw127
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw127 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.CodePoint = (func(_818_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_819_i1 _dafny.Int) _dafny.CodePoint {
							return _818_v12
						}
					})(_803_v12)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw127 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw127).ArraySet1CodePoint(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw127).ArraySet1CodePoint(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_817_v26 = _nw127
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_817_v26), 0))
				_ = _index135
				(_817_v26).ArraySet1CodePoint(_803_v12, (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_817_v26), 0))
				_ = _index136
				(_817_v26).ArraySet1CodePoint(_803_v12, (_index136).Int())
				var _820_v27 _dafny.Map
				_ = _820_v27
				_820_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_817_v26).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_817_v26), 0))).Int()), _dafny.IntOfInt64(118))
				var _821_v28 D1
				_ = _821_v28
				_821_v28 = Companion_D1_.Create_DC5_(_790_v0, Companion_Default___.Fm23((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _820_v27, _806_v15, globalState))
				var _822_v29 D1
				_ = _822_v29
				_822_v29 = Companion_D1_.Create_DC6_(_821_v28)
				var _823_v30 D1
				_ = _823_v30
				_823_v30 = Companion_D1_.Create_DC6_(_821_v28)
				var _824_v31 _dafny.Map
				_ = _824_v31
				_824_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _823_v30)
				_824_v31 = (_824_v31).Update(p0, _823_v30)
				(globalState).F3 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25(Companion_Default___.Fm0(_790_v0, (_791_v1).Cardinality(), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _791_v1, globalState), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _790_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23((_791_v1).Cardinality(), _790_v0, _820_v27, _806_v15, globalState), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), globalState), _796_v6)).Cardinality()
				(globalState).F3 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-204)), Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)))
			}
			_790_v0 = true
			var _825_v32 D2
			_ = _825_v32
			_825_v32 = Companion_D2_.Create_DC9_((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), p0, _803_v12)
			var _826_v33 D4
			_ = _826_v33
			_826_v33 = Companion_D4_.Create_DC13_(_793_v3, (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _801_v11)
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
			_ = _index137
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
			_ = _index138
			var _rhs112 _dafny.Int = (_825_v32).Dtor_cf14()
			_ = _rhs112
			var _rhs113 _dafny.Sequence = (_826_v33).Dtor_cf22()
			_ = _rhs113
			var _rhs114 bool = false
			_ = _rhs114
			var _rhs115 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, p0)).Plus(Companion_Default___.Fm0((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), p0, p0, _791_v1, globalState))
			_ = _rhs115
			var _lhs92 _dafny.Array = _798_v8
			_ = _lhs92
			var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
			_ = _lhs93
			var _lhs94 _dafny.Array = _801_v11
			_ = _lhs94
			var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))
			_ = _lhs95
			var _lhs96 *GlobalState = globalState
			_ = _lhs96
			(_lhs92).ArraySet1(_rhs112, (_lhs93).Int())
			_793_v3 = _rhs113
			(_lhs94).ArraySet1(_rhs114, (_lhs95).Int())
			_lhs96.F3 = _rhs115
		} else {
			(globalState).F3 = _dafny.IntOfInt64(835)
			(globalState).F3 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))))
			var _827_v34 D4
			_ = _827_v34
			_827_v34 = Companion_D4_.Create_DC13_(_793_v3, _dafny.IntOfInt64(452), _790_v0, _801_v11)
			var _828_v35 _dafny.Sequence
			_ = _828_v35
			_828_v35 = _dafny.SeqOf(_827_v34, _827_v34, _827_v34)
			_790_v0 = _dafny.Companion_Sequence_.IsPrefixOf(_828_v35, _828_v35)
			var _829_v36 D2
			_ = _829_v36
			_829_v36 = Companion_D2_.Create_DC9_(_790_v0, (_dafny.Zero).Minus(p0), _803_v12)
			var _830_v37 _dafny.Set
			_ = _830_v37
			_830_v37 = _dafny.SetOf(_dafny.MultiSetOf(!(true)), _dafny.MultiSetOf((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)))
			var _831_v38 _dafny.Map
			_ = _831_v38
			_831_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if false {
					return (_829_v36).Dtor_cf15()
				}
				return _803_v12
			})(), ((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)).Minus((_830_v37).Cardinality()))
			var _832_v39 _dafny.MultiSet
			_ = _832_v39
			_832_v39 = _dafny.MultiSetOf(_798_v8)
			_831_v38 = (_831_v38).Update(_dafny.CodePoint('w'), (_832_v39).Cardinality())
			if !(((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)).Cmp((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)) == 0) {
				(globalState).F13 = _807_v16
				var _833_v40 _dafny.Map
				_ = _833_v40
				_833_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _803_v12)
				var _834_v43 _dafny.Map
				_ = _834_v43
				_834_v43 = _833_v40
				var _835_v44 _dafny.Array
				_ = _835_v44
				var _nwElement0_24 _dafny.Map = _833_v40
				_ = _nwElement0_24
				var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(28))
				_ = _nw128
				(_nw128).ArraySet1(_nwElement0_24, 0)
				(_nw128).ArraySet1((_833_v40).Update(p0, _dafny.CodePoint('o')), 1)
				(_nw128).ArraySet1(_833_v40, 2)
				(_nw128).ArraySet1(func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter40 := _dafny.Iterate((Companion_Default___.Fm30((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), globalState)).Elements()); ; {
						_compr_39, _ok40 := _iter40()
						if !_ok40 {
							break
						}
						var _836_v41 _dafny.Int
						_836_v41 = interface{}(_compr_39).(_dafny.Int)
						if (Companion_Default___.Fm30((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), globalState)).Contains(_836_v41) {
							_coll39.Add(Companion_Default___.SafeDivisionInt(_836_v41, _dafny.IntOfUint32((_796_v6).Cardinality())), _803_v12)
						}
					}
					return _coll39.ToMap()
				}(), 3)
				(_nw128).ArraySet1(_833_v40, 4)
				(_nw128).ArraySet1(func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-39), _dafny.IntOfInt64(-511))); ; {
						_compr_40, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _837_v42 _dafny.Int
						_837_v42 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(-39)).Cmp(_837_v42) <= 0) && ((_837_v42).Cmp(_dafny.IntOfInt64(-511)) < 0) {
							_coll40.Add(Companion_Default___.SafeModuloInt(_837_v42, (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), _dafny.CodePoint('v'))
						}
					}
					return _coll40.ToMap()
				}(), 5)
				(_nw128).ArraySet1(_833_v40, 6)
				(_nw128).ArraySet1(Companion_Default___.Fm31(_790_v0, _790_v0, globalState), 7)
				(_nw128).ArraySet1(_833_v40, 8)
				(_nw128).ArraySet1(_833_v40, 9)
				(_nw128).ArraySet1(_833_v40, 10)
				(_nw128).ArraySet1(_833_v40, 11)
				(_nw128).ArraySet1(_833_v40, 12)
				(_nw128).ArraySet1(_833_v40, 13)
				(_nw128).ArraySet1(_833_v40, 14)
				(_nw128).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _803_v12), 15)
				(_nw128).ArraySet1(_833_v40, 16)
				(_nw128).ArraySet1(_833_v40, 17)
				(_nw128).ArraySet1((_834_v43), 18)
				(_nw128).ArraySet1(_833_v40, 19)
				(_nw128).ArraySet1(_833_v40, 20)
				(_nw128).ArraySet1(Companion_Default___.Fm31(true, _790_v0, globalState), 21)
				(_nw128).ArraySet1(_833_v40, 22)
				(_nw128).ArraySet1(_833_v40, 23)
				(_nw128).ArraySet1(_833_v40, 24)
				(_nw128).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _803_v12), 25)
				(_nw128).ArraySet1(_833_v40, 26)
				(_nw128).ArraySet1(_833_v40, 27)
				_835_v44 = _nw128
				var _838_v45 D0
				_ = _838_v45
				_838_v45 = Companion_D0_.Create_DC1_(_790_v0)
				var _839_v46 D1
				_ = _839_v46
				var _840_v47 _dafny.Int
				_ = _840_v47
				var _841_v48 bool
				_ = _841_v48
				var _842_v49 _dafny.Int
				_ = _842_v49
				var _out43 D1
				_ = _out43
				var _out44 _dafny.Int
				_ = _out44
				var _out45 bool
				_ = _out45
				var _out46 _dafny.Int
				_ = _out46
				_out43, _out44, _out45, _out46 = (_800_v10).M10(_835_v44, _798_v8, ((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)) && ((_838_v45).Dtor_cf1()), globalState)
				_839_v46 = _out43
				_840_v47 = _out44
				_841_v48 = _out45
				_842_v49 = _out46
				var _843_v50 _dafny.MultiSet
				_ = _843_v50
				_843_v50 = _dafny.MultiSetOf(_803_v12, _803_v12, _803_v12, _803_v12, _803_v12)
				var _844_v51 D4
				_ = _844_v51
				_844_v51 = Companion_D4_.Create_DC14_((_dafny.IntOfUint32((Companion_Default___.Fm29(globalState)).Cardinality())).Cmp((_843_v50).Cardinality()) != 0, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool))
				var _845_v52 _dafny.MultiSet
				_ = _845_v52
				_845_v52 = _dafny.MultiSetOf(_841_v48, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _790_v0, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _841_v48)
				var _pat_let_tv21 = _845_v52
				_ = _pat_let_tv21
				var _pat_let_tv22 = _845_v52
				_ = _pat_let_tv22
				_844_v51 = func(_pat_let23_0 D4) D4 {
					return func(_846_dt__update__tmp_h0 D4) D4 {
						return func(_pat_let24_0 bool) D4 {
							return func(_847_dt__update_hcf26_h0 bool) D4 {
								return Companion_D4_.Create_DC14_(_847_dt__update_hcf26_h0, (_846_dt__update__tmp_h0).Dtor_cf27())
							}(_pat_let24_0)
						}((_pat_let_tv21).IsProperSubsetOf(_pat_let_tv22))
					}(_pat_let23_0)
				}(Companion_D4_.Create_DC14_((func() bool {
					if (_this.F22).Contains(_842_v49) {
						return (_this.F22).Get(_842_v49).(bool)
					}
					return _790_v0
				})(), _790_v0))
				(globalState).F3 = _840_v47
				var _848_v53 *C0
				_ = _848_v53
				var _nw129 *C0 = New_C0_()
				_ = _nw129
				_nw129.Ctor__(_790_v0, _790_v0)
				_848_v53 = _nw129
			} else {
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))
				_ = _index139
				(_798_v8).ArraySet1(_dafny.IntOfInt64(288), (_index139).Int())
				(globalState).F9 = (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)
				(globalState).F9 = _790_v0
				var _849_v54 _dafny.Map
				_ = _849_v54
				_849_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-811))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_850_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_851_i2 _dafny.Int) _dafny.CodePoint {
						return _850_v12
					}
				})(_803_v12))))
				_793_v3 = _dafny.Companion_Sequence_.Concatenate(_793_v3, (func() _dafny.Sequence {
					if (_849_v54).Contains((_dafny.Zero).Minus((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))) {
						return (_849_v54).Get((_dafny.Zero).Minus((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int))).(_dafny.Sequence)
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("tecdvlbj")
				})())
				var _852_v55 _dafny.Map
				_ = _852_v55
				_852_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(327), _796_v6)
				var _853_v56 _dafny.MultiSet
				_ = _853_v56
				_853_v56 = _dafny.MultiSetOf(_790_v0, _790_v0, _790_v0)
				var _rhs116 _dafny.Int = (_dafny.SetOf(((_852_v55).Cardinality()).Cmp(_dafny.IntOfInt64(970)) <= 0, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), _790_v0, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool))).Cardinality()
				_ = _rhs116
				var _rhs117 _dafny.Int = (func() _dafny.Int {
					if (func() bool {
						if (_804_v13).Contains((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)) {
							return (_804_v13).Get((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool)).(bool)
						}
						return false
					})() {
						return (_799_v9).Cardinality()
					}
					return (func() _dafny.Int {
						if (_853_v56).Contains(_790_v0) {
							return (_853_v56).Multiplicity(_790_v0)
						}
						return Companion_Default___.Fm0((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), p0, p0, (_791_v1).Update(_dafny.IntOfInt64(554), Companion_Default___.Abs(_dafny.IntOfUint32((_805_v14).Cardinality()))), globalState)
					})()
				})()
				_ = _rhs117
				r0 = _rhs116
				r0 = _rhs117
			}
		}
		var _854_v58 _dafny.Set
		_ = _854_v58
		_854_v58 = _dafny.SetOf(_803_v12)
		r0 = Companion_Default___.Fm0(!(_790_v0), Companion_Default___.Fm0((_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), (_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int), _dafny.MultiSetOf((_798_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_798_v8), 0))).Int()).(_dafny.Int)), globalState), (func() _dafny.Map {
			var _coll41 = _dafny.NewMapBuilder()
			_ = _coll41
			for _iter42 := _dafny.Iterate((_854_v58).Elements()); ; {
				_compr_41, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _855_v57 _dafny.CodePoint
				_855_v57 = interface{}(_compr_41).(_dafny.CodePoint)
				if (_854_v58).Contains(_855_v57) {
					_coll41.Add(_855_v57, (_801_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_801_v11), 0))).Int()).(bool))
				}
			}
			return _coll41.ToMap()
		}()).Cardinality(), (func() _dafny.MultiSet {
			if false {
				return _791_v1
			}
			return _791_v1
		})(), globalState)
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f14 bool
	_f15 _dafny.MultiSet
	_f20 _dafny.Int
	_f21 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f14 = false
	_this._f15 = _dafny.EmptyMultiSet
	_this._f20 = _dafny.Zero
	_this._f21 = _dafny.EmptySeq
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

func (_this *C3) F14() bool {
	return _this._f14
}
func (_this *C3) F14_set_(value bool) {
	_this._f14 = value
}
func (_this *C3) F15() _dafny.MultiSet {
	return _this._f15
}
func (_this *C3) Ctor__(f20 _dafny.Int, f21 _dafny.Sequence, f14 bool, f15 _dafny.MultiSet) {
	{
		(_this)._f20 = f20
		(_this)._f21 = f21
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C3) Fm2(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		if _this.F14() {
			return _this.F14()
		} else {
			return (_dafny.IntOfInt64(-49)).Cmp(_dafny.IntOfUint32(((_this).F21()).Cardinality())) != 0
		}
	}
}
func (_this *C3) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dowcmd"), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F21(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(952))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg63 _dafny.Int) interface{} {
				return coer63(arg63)
			}
		}(func(_856_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()), _dafny.IntOfUint32(((_this).F21()).Cardinality()))).Uint32(), _dafny.CodePoint('y')), (_this).F21())), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dowcmd"), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F21(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(952))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg64 _dafny.Int) interface{} {
				return coer64(arg64)
			}
		}(func(_857_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('j')
		}))).Cardinality()), _dafny.IntOfUint32(((_this).F21()).Cardinality()))).Uint32(), _dafny.CodePoint('y')), (_this).F21()))).Cardinality()))).Uint32(), _dafny.CodePoint('a'))
	}
}
func (_this *C3) M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _858_v0 _dafny.MultiSet
		_ = _858_v0
		_858_v0 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14()))
		var _859_v1 _dafny.Map
		_ = _859_v1
		_859_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14())
		var _860_v2 _dafny.Map
		_ = _860_v2
		_860_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_859_v1).Contains(!(_this.F14())) {
				return (_859_v1).Get(!(_this.F14())).(bool)
			}
			return _this.F14()
		})(), _this.F14())
		var _861_v3 _dafny.MultiSet
		_ = _861_v3
		_861_v3 = _dafny.MultiSetOf((_this).F20())
		_858_v0 = (_858_v0).Difference((_dafny.MultiSetOf((_860_v2).Update(false, _this.F14()), _860_v2)).Update(_859_v1, Companion_Default___.Abs(Companion_Default___.Fm0(_this.F14(), (_this).F20(), (_this).F20(), _861_v3, globalState))))
		var _862_v4 D2
		_ = _862_v4
		_862_v4 = Companion_D2_.Create_DC7_(_dafny.SeqOf(_this.F14()))
		var _863_v5 _dafny.MultiSet
		_ = _863_v5
		_863_v5 = _dafny.MultiSetOf(_862_v4)
		_863_v5 = _863_v5
		(globalState).F3 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(713), ((_this).F20()).Times(_dafny.IntOfInt64(115)))
		var _864_v6 _dafny.Sequence
		_ = _864_v6
		_864_v6 = _dafny.SeqOf(_this.F14(), _this.F14())
		var _hi2 _dafny.Int = _dafny.IntOfUint32((_864_v6).Cardinality())
		_ = _hi2
		for _865_i0 := (_this).F20(); _865_i0.Cmp(_hi2) < 0; _865_i0 = _865_i0.Plus(_dafny.One) {
			var _866_v7 _dafny.Array
			_ = _866_v7
			var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw130
			_866_v7 = _nw130
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _index140
			(_866_v7).ArraySet1(_865_i0, (_index140).Int())
			var _867_v8 _dafny.Array
			_ = _867_v8
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw131
			_867_v8 = _nw131
			var _868_v9 D1
			_ = _868_v9
			_868_v9 = Companion_D1_.Create_DC4_(_865_i0, _dafny.IntOfUint32(((_this).Fm3((_this).F20(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fo")).Cardinality()), _this.F14(), (_this).F20(), globalState)).Cardinality()))
			var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _index141
			var _rhs118 _dafny.Array = _867_v8
			_ = _rhs118
			var _rhs119 bool = !(_this.F14())
			_ = _rhs119
			var _rhs120 _dafny.Int = (func() _dafny.Int {
				if ((_this).F15()).Contains((_this).Fm2(_this.F14(), _this.F14(), globalState)) {
					return ((_this).F15()).Multiplicity((_this).Fm2(_this.F14(), _this.F14(), globalState))
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}(func(_869_i1 _dafny.Int) _dafny.Sequence {
					return (_this).F21()
				}))).Cardinality())
			})()
			_ = _rhs120
			var _rhs121 _dafny.Int = (_868_v9).Dtor_cf4()
			_ = _rhs121
			var _lhs97 *GlobalState = globalState
			_ = _lhs97
			var _lhs98 *GlobalState = globalState
			_ = _lhs98
			var _lhs99 _dafny.Array = _866_v7
			_ = _lhs99
			var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _lhs100
			var _lhs101 *GlobalState = globalState
			_ = _lhs101
			_lhs97.F2 = _rhs118
			_lhs98.F9 = _rhs119
			(_lhs99).ArraySet1(_rhs120, (_lhs100).Int())
			_lhs101.F3 = _rhs121
			var _870_v10 _dafny.Array
			_ = _870_v10
			var _nw132 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw132
			_870_v10 = _nw132
			var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_870_v10), 0))
			_ = _index142
			(_870_v10).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-993))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_871_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), (_index142).Int())
			var _872_v11 _dafny.Array
			_ = _872_v11
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_17
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) D1 = (func(_873_v9 D1, _874_v7 _dafny.Array) func(_dafny.Int) D1 {
					return func(_875_i3 _dafny.Int) D1 {
						return func(_pat_let25_0 D1) D1 {
							return func(_876_dt__update__tmp_h0 D1) D1 {
								return func(_pat_let26_0 _dafny.Int) D1 {
									return func(_877_dt__update_hcf4_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC4_(_877_dt__update_hcf4_h0, (_876_dt__update__tmp_h0).Dtor_cf5())
									}(_pat_let26_0)
								}((_dafny.Zero).Minus((_874_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_874_v7), 0))).Int()).(_dafny.Int)))
							}(_pat_let25_0)
						}(_873_v9)
					}
				})(_868_v9, _866_v7)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw133 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw133).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw133).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_872_v11 = _nw133
			var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_872_v11), 0))
			_ = _index143
			(_872_v11).ArraySet1(_868_v9, (_index143).Int())
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _index144
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_870_v10), 0))
			_ = _index145
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_872_v11), 0))
			_ = _index146
			var _rhs122 _dafny.Int = _865_i0
			_ = _rhs122
			var _rhs123 _dafny.Sequence = (_this).F21()
			_ = _rhs123
			var _rhs124 bool = ((func() bool {
				if true {
					return false
				}
				return _this.F14()
			})()) || (_this.F14())
			_ = _rhs124
			var _rhs125 _dafny.Int = _865_i0
			_ = _rhs125
			var _rhs126 D1 = _868_v9
			_ = _rhs126
			var _lhs102 _dafny.Array = _866_v7
			_ = _lhs102
			var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _lhs103
			var _lhs104 _dafny.Array = _870_v10
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_870_v10), 0))
			_ = _lhs105
			var _lhs106 *GlobalState = globalState
			_ = _lhs106
			var _lhs107 _dafny.Array = _872_v11
			_ = _lhs107
			var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(817), _dafny.ArrayLen((_872_v11), 0))
			_ = _lhs108
			(_lhs102).ArraySet1(_rhs122, (_lhs103).Int())
			(_lhs104).ArraySet1(_rhs123, (_lhs105).Int())
			r1 = _rhs124
			_lhs106.F3 = _rhs125
			(_lhs107).ArraySet1(_rhs126, (_lhs108).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
			_ = _index147
			(_866_v7).ArraySet1(_dafny.IntOfInt64(578), (_index147).Int())
			if true {
				var _878_v12 _dafny.Map
				_ = _878_v12
				_878_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F14())
				_878_v12 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_865_i0, true)).Merge((_878_v12).Merge(_878_v12))
				var _879_v13 D0
				_ = _879_v13
				_879_v13 = Companion_D0_.Create_DC0_(!(_this.F14()))
				var _880_v14 _dafny.Array
				_ = _880_v14
				var _nwElement0_25 D0 = _879_v13
				_ = _nwElement0_25
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(13))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_25, 0)
				(_nw134).ArraySet1(_879_v13, 1)
				(_nw134).ArraySet1(Companion_Default___.Fm20(_dafny.IntOfInt64(-575), (_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), (_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), globalState), 2)
				(_nw134).ArraySet1((func() D0 {
					if _this.F14() {
						return _879_v13
					}
					return _879_v13
				})(), 3)
				(_nw134).ArraySet1(Companion_Default___.Fm20(_dafny.IntOfInt64(600), (_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), (_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), globalState), 4)
				(_nw134).ArraySet1(_879_v13, 5)
				(_nw134).ArraySet1((func() D0 {
					if _this.F14() {
						return _879_v13
					}
					return _879_v13
				})(), 6)
				(_nw134).ArraySet1(_879_v13, 7)
				(_nw134).ArraySet1(_879_v13, 8)
				(_nw134).ArraySet1((func() D0 {
					if _this.F14() {
						return _879_v13
					}
					return _879_v13
				})(), 9)
				(_nw134).ArraySet1(_879_v13, 10)
				(_nw134).ArraySet1(_879_v13, 11)
				(_nw134).ArraySet1(_879_v13, 12)
				_880_v14 = _nw134
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_880_v14), 0))
				_ = _index148
				(_880_v14).ArraySet1(_879_v13, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(768), _dafny.ArrayLen((_880_v14), 0))
				_ = _index149
				(_880_v14).ArraySet1(_879_v13, (_index149).Int())
				var _881_v15 _dafny.CodePoint
				_ = _881_v15
				_881_v15 = _dafny.CodePoint('e')
				var _882_v16 _dafny.MultiSet
				_ = _882_v16
				_882_v16 = _dafny.MultiSetOf(_881_v15, _dafny.CodePoint('p'))
				var _883_v17 _dafny.Sequence
				_ = _883_v17
				_883_v17 = _dafny.SeqOf((_this).F20())
				var _rhs127 _dafny.Int = (_this).F20()
				_ = _rhs127
				var _rhs128 _dafny.MultiSet = _882_v16
				_ = _rhs128
				var _rhs129 _dafny.Int = _865_i0
				_ = _rhs129
				var _rhs130 _dafny.Int = (_this).F20()
				_ = _rhs130
				var _rhs131 _dafny.Int = (_883_v17).Select((Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32((_883_v17).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs131
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				var _lhs111 *GlobalState = globalState
				_ = _lhs111
				var _lhs112 *GlobalState = globalState
				_ = _lhs112
				_lhs109.F3 = _rhs127
				_882_v16 = _rhs128
				_lhs110.F3 = _rhs129
				_lhs111.F3 = _rhs130
				_lhs112.F3 = _rhs131
				var _884_v18 _dafny.Map
				_ = _884_v18
				_884_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.MultiSetOf(_this.F14(), _this.F14(), _this.F14(), _this.F14(), false))
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_867_v8), 0))
				_ = _index150
				(_867_v8).ArraySet1((((func() _dafny.MultiSet {
					if (_884_v18).Contains(!(_this.F14())) {
						return (_884_v18).Get(!(_this.F14())).(_dafny.MultiSet)
					}
					return (_this).F15()
				})()).Cardinality()).Cmp(_865_i0) != 0, (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(322), _dafny.ArrayLen((_867_v8), 0))
				_ = _index151
				(_867_v8).ArraySet1(true, (_index151).Int())
				(globalState).F3 = ((_this).F20()).Minus((_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int))
			} else {
				var _885_v19 *C2
				_ = _885_v19
				var _nw135 *C2 = New_C2_()
				_ = _nw135
				_nw135.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F14()))
				_885_v19 = _nw135
				var _886_v20 _dafny.Map
				_ = _886_v20
				_886_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if _this.F14() {
						return _this.F14()
					}
					return _this.F14()
				})(), (_this).F20())
				_886_v20 = (_886_v20).Update(_this.F14(), ((_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int)).Times(_865_i0))
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))
				_ = _index152
				(_866_v7).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(657), (_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int)), _dafny.IntOfInt64(631)), (_index152).Int())
				var _887_v21 _dafny.Map
				_ = _887_v21
				_887_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), _865_i0)
				var _888_v22 _dafny.Map
				_ = _888_v22
				_888_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21)
				var _889_v23 _dafny.Sequence
				_ = _889_v23
				_889_v23 = _dafny.SeqOf((_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int))
				var _890_v24 _dafny.Array
				_ = _890_v24
				var _nwElement0_26 _dafny.Map = _888_v22
				_ = _nwElement0_26
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(25))
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_26, 0)
				(_nw136).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21)).Merge(_888_v22), 1)
				(_nw136).ArraySet1(_888_v22, 2)
				(_nw136).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21), 3)
				(_nw136).ArraySet1(_888_v22, 4)
				(_nw136).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21), 5)
				(_nw136).ArraySet1(_888_v22, 6)
				(_nw136).ArraySet1((_888_v22).Merge(_888_v22), 7)
				(_nw136).ArraySet1(_888_v22, 8)
				(_nw136).ArraySet1(_888_v22, 9)
				(_nw136).ArraySet1(_888_v22, 10)
				(_nw136).ArraySet1(_888_v22, 11)
				(_nw136).ArraySet1(_888_v22, 12)
				(_nw136).ArraySet1(_888_v22, 13)
				(_nw136).ArraySet1(_888_v22, 14)
				(_nw136).ArraySet1(_888_v22, 15)
				(_nw136).ArraySet1(_888_v22, 16)
				(_nw136).ArraySet1((_888_v22).Merge(_888_v22), 17)
				(_nw136).ArraySet1((_888_v22).Update(_866_v7, _887_v21), 18)
				(_nw136).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21)).Merge(_888_v22), 19)
				(_nw136).ArraySet1((_888_v22).Merge(_888_v22), 20)
				(_nw136).ArraySet1((_888_v22).Merge(_888_v22), 21)
				(_nw136).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _887_v21), 22)
				(_nw136).ArraySet1((_888_v22).Merge(_888_v22), 23)
				(_nw136).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_866_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(135), Companion_Default___.Fm0(_this.F14(), _865_i0, (_dafny.Zero).Minus((_889_v23).Select((Companion_Default___.SafeIndex((_866_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_866_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_889_v23).Cardinality()))).Uint32()).(_dafny.Int)), _dafny.MultiSetFromSeq(_889_v23), globalState))), 24)
				_890_v24 = _nw136
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_890_v24), 0))
				_ = _index153
				(_890_v24).ArraySet1((_888_v22).Merge(_888_v22), (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_890_v24), 0))
				_ = _index154
				(_890_v24).ArraySet1((_888_v22).Update(_866_v7, _887_v21), (_index154).Int())
				(_this).F14_set_((func() bool {
					if (_864_v6).Select((Companion_Default___.SafeIndex(_865_i0, _dafny.IntOfUint32((_864_v6).Cardinality()))).Uint32()).(bool) {
						return _this.F14()
					}
					return _this.F14()
				})())
			}
		}
		if !(_this.F14()) {
			var _891_v25 _dafny.Sequence
			_ = _891_v25
			_891_v25 = _dafny.SeqOf(_dafny.IntOfInt64(301), ((_this).F20()).Times((_dafny.Zero).Minus((_this).F20())), (_this).F20())
			_891_v25 = _891_v25
			var _892_v26 D3
			_ = _892_v26
			_892_v26 = Companion_D3_.Create_DC11_((_this).F20(), ((_this).F20()).Times((_this).F20()), _this.F14(), (_dafny.IntOfInt64(326)).Plus((Companion_Default___.Fm32((_this).F21(), _861_v3, globalState)).Cardinality()))
			_892_v26 = _892_v26
			var _893_v27 _dafny.Array
			_ = _893_v27
			var _nwElement0_27 _dafny.Int = _dafny.IntOfInt64(898)
			_ = _nwElement0_27
			var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(8))
			_ = _nw137
			(_nw137).ArraySet1(_nwElement0_27, 0)
			(_nw137).ArraySet1((_this).F20(), 1)
			(_nw137).ArraySet1((_this).F20(), 2)
			(_nw137).ArraySet1((_this).F20(), 3)
			(_nw137).ArraySet1((_this).F20(), 4)
			(_nw137).ArraySet1((_this).F20(), 5)
			(_nw137).ArraySet1(_dafny.IntOfUint32((_864_v6).Cardinality()), 6)
			(_nw137).ArraySet1((_this).F20(), 7)
			_893_v27 = _nw137
			var _894_v28 _dafny.Sequence
			_ = _894_v28
			_894_v28 = _dafny.SeqOf(_893_v27)
			_894_v28 = _dafny.Companion_Sequence_.Concatenate(_894_v28, _894_v28)
			r3 = (((_this).F20()).Cmp((_dafny.Zero).Minus((_this).F20())) > 0) && (_this.F14())
			(globalState).F9 = _this.F14()
		} else {
			var _895_v29 _dafny.Array
			_ = _895_v29
			var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw138
			_895_v29 = _nw138
			var _896_v30 _dafny.Map
			_ = _896_v30
			_896_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wnwodjwai"), (_this).F21()), _895_v29)
			_896_v30 = (_896_v30).Update((_this).F21(), _895_v29)
			var _897_v31 *C0
			_ = _897_v31
			var _nw139 *C0 = New_C0_()
			_ = _nw139
			_nw139.Ctor__(_this.F14(), _this.F14())
			_897_v31 = _nw139
			var _898_v32 _dafny.Array
			_ = _898_v32
			var _nw140 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
			_ = _nw140
			_898_v32 = _nw140
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_898_v32), 0))
			_ = _index155
			(_898_v32).ArraySet1(_this.F14(), (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(829), _dafny.ArrayLen((_898_v32), 0))
			_ = _index156
			(_898_v32).ArraySet1((_897_v31).F23(), (_index156).Int())
			var _899_v33 _dafny.Array
			_ = _899_v33
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
			_ = _nw141
			_899_v33 = _nw141
			var _900_v34 _dafny.CodePoint
			_ = _900_v34
			_900_v34 = _dafny.CodePoint('p')
			var _901_v35 _dafny.Map
			_ = _901_v35
			_901_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_899_v33, _dafny.Companion_Sequence_.Update((_this).F21(), (Companion_Default___.SafeIndex((_this).F20(), _dafny.IntOfUint32(((_this).F21()).Cardinality()))).Uint32(), _900_v34))
			_901_v35 = _901_v35
			(globalState).F3 = _dafny.IntOfInt64(255)
		}
		(globalState).F3 = (_this).F20()
		r0 = (func() _dafny.Int {
			if (_861_v3).Contains((_this).F20()) {
				return (_861_v3).Multiplicity((_this).F20())
			}
			return (_this).F20()
		})()
		r1 = _this.F14()
		var _902_v36 _dafny.Map
		_ = _902_v36
		_902_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), _this.F14())
		var _903_v37 _dafny.Map
		_ = _903_v37
		_903_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_902_v36, (_this).F20())
		r2 = (Companion_D7_.Create_DC21_(_903_v37)).Dtor_cf36()
		r3 = true
		return r0, r1, r2, r3
	}
}
func (_this *C3) M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		(_this).F14_set_(_this.F14())
		var _904_v0 _dafny.Array
		_ = _904_v0
		var _nw142 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw142
		_904_v0 = _nw142
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))
		_ = _index157
		(_904_v0).ArraySet1(_this.F14(), (_index157).Int())
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))
		_ = _index158
		(_904_v0).ArraySet1(_this.F14(), (_index158).Int())
		var _905_v1 *C0
		_ = _905_v1
		var _nw143 *C0 = New_C0_()
		_ = _nw143
		_nw143.Ctor__(_this.F14(), (_this.F14()) == (_this.F14()))
		_905_v1 = _nw143
		var _906_v2 _dafny.Array
		_ = _906_v2
		var _nw144 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw144
		_906_v2 = _nw144
		var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_906_v2), 0))
		_ = _index159
		(_906_v2).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm27((p0).Cardinality(), (_905_v1).F24(), globalState)).Cardinality()), (_index159).Int())
		var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_906_v2), 0))
		_ = _index160
		(_906_v2).ArraySet1(Companion_Default___.SafeDivisionInt(p1, Companion_Default___.Fm0((_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool), (_this).F20(), p1, Companion_Default___.Fm30(false, globalState), globalState)), (_index160).Int())
		for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_906_v2), 0))); ; {
			_guard_loop_1, _ok43 := _iter43()
			if !_ok43 {
				break
			}
			var _907_i0 _dafny.Int
			_907_i0 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_907_i0).Sign() != -1) && ((_907_i0).Cmp(_dafny.ArrayLen((_906_v2), 0)) < 0)) {
				(_906_v2).ArraySet1((_907_i0).Plus((Companion_D3_.Create_DC11_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F20(), (_dafny.Zero).Minus((_this).F20()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_D2_.Create_DC9_(!((_905_v1).F24()), p1, _dafny.CodePoint('w'))))).Cardinality(), (_905_v1).F24(), (_this).F20())).Dtor_cf20()), (_907_i0).Int())
			}
		}
		if false {
			r2 = (func() bool {
				if (_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool) {
					return (p1).Cmp(_dafny.IntOfUint32(((_this).F21()).Cardinality())) >= 0
				}
				return _this.F14()
			})()
			var _908_v3 _dafny.Array
			_ = _908_v3
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw145
			_908_v3 = _nw145
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_908_v3), 0))
			_ = _index161
			(_908_v3).ArraySet1((_this).F21(), (_index161).Int())
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_908_v3), 0))
			_ = _index162
			(_908_v3).ArraySet1((_this).F21(), (_index162).Int())
			var _909_v4 _dafny.Map
			_ = _909_v4
			_909_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_905_v1).F23())
			var _910_v5 _dafny.Sequence
			_ = _910_v5
			_910_v5 = _dafny.SeqOf((_905_v1).F23())
			(globalState).F6 = _dafny.SetOf(true, (func() bool {
				if (_905_v1).F23() {
					return (func() bool {
						if (_909_v4).Contains((_908_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_908_v3), 0))).Int()).(_dafny.Sequence)) {
							return (_909_v4).Get((_908_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_908_v3), 0))).Int()).(_dafny.Sequence)).(bool)
						}
						return (_910_v5).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_910_v5).Cardinality()))).Uint32()).(bool)
					})()
				}
				return (_905_v1).F24()
			})(), !(!(false)))
			(_this).F14_set_(!((_this).Fm2(_this.F14(), !(((_905_v1).F23()) == ((_905_v1).F24())), globalState)))
			var _911_v6 _dafny.CodePoint
			_ = _911_v6
			_911_v6 = _dafny.CodePoint('f')
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))
			_ = _index163
			var _rhs132 bool = ((_905_v1).F24()) && ((_905_v1).F24())
			_ = _rhs132
			var _rhs133 _dafny.CodePoint = _911_v6
			_ = _rhs133
			var _rhs134 bool = (_905_v1).F24()
			_ = _rhs134
			var _lhs113 *C3 = _this
			_ = _lhs113
			var _lhs114 _dafny.Array = _904_v0
			_ = _lhs114
			var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))
			_ = _lhs115
			_lhs113.F14_set_(_rhs132)
			_911_v6 = _rhs133
			(_lhs114).ArraySet1(_rhs134, (_lhs115).Int())
		} else {
			var _912_v7 _dafny.Array
			_ = _912_v7
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw146
			_912_v7 = _nw146
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_912_v7), 0))
			_ = _index164
			(_912_v7).ArraySet1(_906_v2, (_index164).Int())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(283), _dafny.ArrayLen((_912_v7), 0))
			_ = _index165
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw147
			(_912_v7).ArraySet1(_nw147, (_index165).Int())
			(globalState).F9 = (_905_v1).F24()
			var _913_v8 _dafny.Map
			_ = _913_v8
			_913_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool), !((_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool)))
			var _914_v9 D3
			_ = _914_v9
			_914_v9 = Companion_D3_.Create_DC10_((_913_v8).Update((_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool), (_905_v1).F23()))
			var _915_v10 _dafny.Map
			_ = _915_v10
			_915_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_914_v9, ((_906_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_906_v2), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F14(), (_905_v1).F24())).Cardinality()))) < 0)
			_915_v10 = (_915_v10).Update(_914_v9, (_905_v1).F23())
			var _916_v11 _dafny.Set
			_ = _916_v11
			_916_v11 = _dafny.SetOf(p1)
			_916_v11 = (_916_v11).Difference((func() _dafny.Set {
				if (_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool) {
					return _916_v11
				}
				return _916_v11
			})())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))
			_ = _index166
			(_904_v0).ArraySet1((_905_v1).F23(), (_index166).Int())
		}
		var _917_v12 _dafny.Sequence
		_ = _917_v12
		_917_v12 = _dafny.SeqOf(_this.F14(), false)
		var _918_v13 _dafny.CodePoint
		_ = _918_v13
		_918_v13 = _dafny.CodePoint('h')
		var _919_v14 _dafny.Map
		_ = _919_v14
		_919_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_917_v12).Cardinality()), _918_v13)
		r0 = ((((_this).F15()).Difference((_this).F15())).Cardinality()).Minus((_919_v14).Cardinality())
		var _920_v15 _dafny.MultiSet
		_ = _920_v15
		_920_v15 = _dafny.MultiSetOf((_906_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_906_v2), 0))).Int()).(_dafny.Int), p1)
		r1 = (((_dafny.SetOf(p1)).Cardinality()).Times((_920_v15).Cardinality())).Minus(Companion_Default___.SafeDivisionInt((_this).F20(), _dafny.IntOfInt64(546)))
		r2 = !(!(!(false)) || ((p1).Cmp(_dafny.IntOfUint32((Companion_Default___.Fm27((_dafny.Zero).Minus(_dafny.IntOfUint32((_917_v12).Cardinality())), (_904_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(887), _dafny.ArrayLen((_904_v0), 0))).Int()).(bool), globalState)).Cardinality())) > 0))
		return r0, r1, r2
	}
}
func (_this *C3) F20() _dafny.Int {
	{
		return _this._f20
	}
}
func (_this *C3) F21() _dafny.Sequence {
	{
		return _this._f21
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f14 bool
	_f15 _dafny.MultiSet
	_f19 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f14 = false
	_this._f15 = _dafny.EmptyMultiSet
	_this._f19 = false
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

func (_this *C4) F14() bool {
	return _this._f14
}
func (_this *C4) F14_set_(value bool) {
	_this._f14 = value
}
func (_this *C4) F15() _dafny.MultiSet {
	return _this._f15
}
func (_this *C4) Ctor__(f19 bool, f14 bool, f15 _dafny.MultiSet) {
	{
		(_this)._f19 = f19
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C4) Fm2(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return (_this).F19()
	}
}
func (_this *C4) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-13))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg67 _dafny.Int) interface{} {
				return coer67(arg67)
			}
		}(func(_921_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		})), _dafny.UnicodeSeqOfUtf8Bytes("xlscb")), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fu"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(866))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg68 _dafny.Int) interface{} {
				return coer68(arg68)
			}
		}(func(_922_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		}))))
	}
}
func (_this *C4) Fm17(p0 _dafny.Int, p1 bool, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (((_this).F15()).Intersection((_this).F15())).Cardinality()
	}
}
func (_this *C4) Fm18(p0 bool, p1 _dafny.Int, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return ((Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_this).F19()))).Dtor_cf16()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _this.F14()))
	}
}
func (_this *C4) M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		if !((_this).F19()) {
			var _923_v0 _dafny.Array
			_ = _923_v0
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_18
			var _nw148 _dafny.Array
			_ = _nw148
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw148 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = func(_924_i0 _dafny.Int) bool {
					return false
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw148 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw148).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw148).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_923_v0 = _nw148
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_923_v0), 0))
			_ = _index167
			(_923_v0).ArraySet1(!((_this).F19()) || ((_this).F19()), (_index167).Int())
			var _925_v1 _dafny.Int
			_ = _925_v1
			_925_v1 = _dafny.IntOfInt64(-689)
			var _926_v2 _dafny.Sequence
			_ = _926_v2
			_926_v2 = _dafny.SeqOf(_this.F14())
			var _927_v3 _dafny.CodePoint
			_ = _927_v3
			_927_v3 = _dafny.CodePoint('v')
			var _928_v4 _dafny.Set
			_ = _928_v4
			_928_v4 = _dafny.SetOf(_927_v3)
			var _929_v5 _dafny.Map
			_ = _929_v5
			_929_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_927_v3, (_this).F19())
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_923_v0), 0))
			_ = _index168
			(_923_v0).ArraySet1(((Companion_Default___.Fm19(_925_v1, (_926_v2).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_925_v1), _dafny.IntOfUint32((_926_v2).Cardinality()))).Uint32()).(bool), globalState)).Difference(_928_v4)).IsDisjointFrom(func() _dafny.Set {
				var _coll42 = _dafny.NewBuilder()
				_ = _coll42
				for _iter44 := _dafny.Iterate((_929_v5).Keys().Elements()); ; {
					_compr_42, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _930_v6 _dafny.CodePoint
					_930_v6 = interface{}(_compr_42).(_dafny.CodePoint)
					if (_929_v5).Contains(_930_v6) {
						_coll42.Add(_930_v6)
					}
				}
				return _coll42.ToSet()
			}()), (_index168).Int())
			(globalState).F9 = _this.F14()
			var _931_v7 _dafny.Map
			_ = _931_v7
			_931_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_923_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_923_v0), 0))).Int()).(bool), _this.F14())
			var _932_v8 _dafny.Sequence
			_ = _932_v8
			_932_v8 = _dafny.SeqOf(_931_v7)
			(globalState).F9 = _dafny.Companion_Sequence_.Contains(_932_v8, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14()))
			var _933_v9 _dafny.Sequence
			_ = _933_v9
			_933_v9 = _dafny.SeqOf((_this).F15())
			var _934_v10 _dafny.Array
			_ = _934_v10
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_19
			var _nw149 _dafny.Array
			_ = _nw149
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw149 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = (func(_935_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_936_i1 _dafny.Int) _dafny.Int {
						return (_936_i1).Times(_935_v1)
					}
				})(_925_v1)
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw149 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw149).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw149).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_934_v10 = _nw149
			var _937_v11 _dafny.Set
			_ = _937_v11
			_937_v11 = _dafny.SetOf(_934_v10, _934_v10, _934_v10)
			var _938_v12 _dafny.MultiSet
			_ = _938_v12
			_938_v12 = _dafny.MultiSetOf(_925_v1)
			(_this).M7(_933_v9, (_dafny.SetOf(_934_v10)).IsDisjointFrom(_937_v11), (func() _dafny.Int {
				if (_this).F19() {
					return _dafny.IntOfInt64(772)
				}
				return _925_v1
			})(), (_925_v1).Cmp(Companion_Default___.Fm0(_this.F14(), _dafny.IntOfInt64(-955), _925_v1, _938_v12, globalState)) >= 0, globalState)
			var _939_v13 _dafny.Map
			_ = _939_v13
			_939_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_925_v1, _927_v3)
			var _940_v14 D3
			_ = _940_v14
			_940_v14 = Companion_D3_.Create_DC11_(_925_v1, _925_v1, (_this).F19(), (_939_v13).Cardinality())
			(globalState).F3 = Companion_Default___.Fm0((_925_v1).Cmp((((_this).F15()).Update((_this).F19(), Companion_Default___.Abs(_925_v1))).Cardinality()) >= 0, _925_v1, (_940_v14).Dtor_cf20(), _938_v12, globalState)
		} else {
			var _941_v15 D1
			_ = _941_v15
			_941_v15 = Companion_D1_.Create_DC5_(_this.F14(), _this.F14())
			var _942_v16 _dafny.Array
			_ = _942_v16
			var _nwElement0_28 D1 = _941_v15
			_ = _nwElement0_28
			var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(18))
			_ = _nw150
			(_nw150).ArraySet1(_nwElement0_28, 0)
			(_nw150).ArraySet1(_941_v15, 1)
			(_nw150).ArraySet1(_941_v15, 2)
			(_nw150).ArraySet1(_941_v15, 3)
			(_nw150).ArraySet1(_941_v15, 4)
			(_nw150).ArraySet1(_941_v15, 5)
			(_nw150).ArraySet1(_941_v15, 6)
			(_nw150).ArraySet1(_941_v15, 7)
			(_nw150).ArraySet1(Companion_D1_.Create_DC5_(false, true), 8)
			(_nw150).ArraySet1(_941_v15, 9)
			(_nw150).ArraySet1(_941_v15, 10)
			(_nw150).ArraySet1(_941_v15, 11)
			(_nw150).ArraySet1(_941_v15, 12)
			(_nw150).ArraySet1(_941_v15, 13)
			(_nw150).ArraySet1(_941_v15, 14)
			(_nw150).ArraySet1(_941_v15, 15)
			(_nw150).ArraySet1(Companion_D1_.Create_DC5_(_this.F14(), _this.F14()), 16)
			(_nw150).ArraySet1(_941_v15, 17)
			_942_v16 = _nw150
			var _943_v17 _dafny.Sequence
			_ = _943_v17
			_943_v17 = _dafny.SeqOf(_942_v16, _942_v16, _942_v16)
			var _944_v18 _dafny.Int
			_ = _944_v18
			_944_v18 = _dafny.IntOfInt64(-618)
			var _945_v19 _dafny.Array
			_ = _945_v19
			var _nwElement0_29 _dafny.Sequence = _943_v17
			_ = _nwElement0_29
			var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(12))
			_ = _nw151
			(_nw151).ArraySet1(_nwElement0_29, 0)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17), 1)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17), (Companion_Default___.SafeIndex(_944_v18, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17)).Cardinality()))).Uint32(), _942_v16), 2)
			(_nw151).ArraySet1(_943_v17, 3)
			(_nw151).ArraySet1(_943_v17, 4)
			(_nw151).ArraySet1(_943_v17, 5)
			(_nw151).ArraySet1(_943_v17, 6)
			(_nw151).ArraySet1(_dafny.SeqOf(_942_v16, _942_v16), 7)
			(_nw151).ArraySet1(_943_v17, 8)
			(_nw151).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17), 9)
			(_nw151).ArraySet1(_943_v17, 10)
			(_nw151).ArraySet1(_943_v17, 11)
			_945_v19 = _nw151
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_945_v19), 0))
			_ = _index169
			(_945_v19).ArraySet1(_943_v17, (_index169).Int())
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(930), _dafny.ArrayLen((_945_v19), 0))
			_ = _index170
			(_945_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17), _dafny.Companion_Sequence_.Concatenate(_943_v17, _943_v17)), (_index170).Int())
			(globalState).F9 = (_944_v18).Cmp((_dafny.Zero).Minus(_944_v18)) <= 0
			var _946_v20 _dafny.Array
			_ = _946_v20
			var _nw152 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw152
			_946_v20 = _nw152
			var _947_v21 _dafny.Array
			_ = _947_v21
			var _nwElement0_30 _dafny.Array = _946_v20
			_ = _nwElement0_30
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(8))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_30, 0)
			(_nw153).ArraySet1(_946_v20, 1)
			(_nw153).ArraySet1(_946_v20, 2)
			(_nw153).ArraySet1(_946_v20, 3)
			(_nw153).ArraySet1(_946_v20, 4)
			(_nw153).ArraySet1(_946_v20, 5)
			(_nw153).ArraySet1((func() _dafny.Array {
				if true {
					return _946_v20
				}
				return _946_v20
			})(), 6)
			(_nw153).ArraySet1(_946_v20, 7)
			_947_v21 = _nw153
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_947_v21), 0))
			_ = _index171
			(_947_v21).ArraySet1(_946_v20, (_index171).Int())
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_947_v21), 0))
			_ = _index172
			var _nw154 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
			_ = _nw154
			(_947_v21).ArraySet1(_nw154, (_index172).Int())
			var _948_v22 _dafny.Sequence
			_ = _948_v22
			_948_v22 = _dafny.UnicodeSeqOfUtf8Bytes("gkgn")
			var _949_v23 T0
			_ = _949_v23
			var _nw155 *C3 = New_C3_()
			_ = _nw155
			_nw155.Ctor__(_944_v18, _948_v22, !((_this).F19()), Companion_Default___.Fm33((_this).F19(), globalState))
			_949_v23 = _nw155
			var _950_v24 _dafny.Map
			_ = _950_v24
			_950_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_944_v18).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_944_v18)).Cardinality())), _949_v23)
			var _951_v25 _dafny.Map
			_ = _951_v25
			_951_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_944_v18, _dafny.IntOfInt64(-353))
			var _952_v26 _dafny.MultiSet
			_ = _952_v26
			_952_v26 = _dafny.MultiSetOf((_951_v25).Cardinality())
			var _953_v27 D1
			_ = _953_v27
			_953_v27 = Companion_D1_.Create_DC4_(_944_v18, Companion_Default___.Fm0(_949_v23.F14(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_948_v22).Cardinality())), _dafny.IntOfInt64(548), _952_v26, globalState))
			_950_v24 = (_950_v24).Update((_953_v27).Dtor_cf5(), _949_v23)
			_944_v18 = Companion_Default___.SafeModuloInt(_944_v18, (_944_v18).Times(_944_v18))
		}
		var _954_v28 _dafny.Int
		_ = _954_v28
		_954_v28 = _dafny.IntOfInt64(-847)
		var _955_v29 _dafny.Sequence
		_ = _955_v29
		_955_v29 = _dafny.SeqOf(_954_v28)
		var _956_v30 _dafny.MultiSet
		_ = _956_v30
		_956_v30 = _dafny.MultiSetOf((_955_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-609))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg69 _dafny.Int) interface{} {
				return coer69(arg69)
			}
		}(func(_957_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()), _dafny.IntOfUint32((_955_v29).Cardinality()))).Uint32()).(_dafny.Int))
		var _958_v31 _dafny.Sequence
		_ = _958_v31
		_958_v31 = _dafny.SeqOf(_954_v28, _954_v28, (_956_v30).Cardinality())
		_955_v29 = _958_v31
		var _959_v32 _dafny.Sequence
		_ = _959_v32
		_959_v32 = _dafny.UnicodeSeqOfUtf8Bytes("hsfh")
		(_this).F14_set_(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("eduiltxa"), _959_v32)) || (_this.F14()))
		var _960_v33 _dafny.Map
		_ = _960_v33
		_960_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_954_v28, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.IntOfUint32((_dafny.SeqOf(_this.F14(), false, _this.F14())).Cardinality())))
		var _961_v34 _dafny.Map
		_ = _961_v34
		_961_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v33, (_this).F19())
		var _962_v35 _dafny.Sequence
		_ = _962_v35
		_962_v35 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-201))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg70 _dafny.Int) interface{} {
				return coer70(arg70)
			}
		}(func(_963_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		})))
		var _964_v36 _dafny.Array
		_ = _964_v36
		var _nwElement0_31 bool = (func() bool {
			if (_961_v34).Contains(_960_v33) {
				return (_961_v34).Get(_960_v33).(bool)
			}
			return _this.F14()
		})()
		_ = _nwElement0_31
		var _nw156 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(22))
		_ = _nw156
		(_nw156).ArraySet1(_nwElement0_31, 0)
		(_nw156).ArraySet1((_dafny.IntOfInt64(377)).Cmp(_954_v28) <= 0, 1)
		(_nw156).ArraySet1((_this).F19(), 2)
		(_nw156).ArraySet1((_954_v28).Cmp((_dafny.Zero).Minus(_954_v28)) <= 0, 3)
		(_nw156).ArraySet1((_this).F19(), 4)
		(_nw156).ArraySet1((_this).Fm2((_this).F19(), _this.F14(), globalState), 5)
		(_nw156).ArraySet1(!((_this).F19()), 6)
		(_nw156).ArraySet1((_this).F19(), 7)
		(_nw156).ArraySet1(_this.F14(), 8)
		(_nw156).ArraySet1(!_dafny.Companion_Sequence_.Equal(_959_v32, (_962_v35).Select((Companion_Default___.SafeIndex(_954_v28, _dafny.IntOfUint32((_962_v35).Cardinality()))).Uint32()).(_dafny.Sequence)), 9)
		(_nw156).ArraySet1((_this).Fm2((_this).F19(), (_this).F19(), globalState), 10)
		(_nw156).ArraySet1(!((_this).F19()), 11)
		(_nw156).ArraySet1(((_this).F19()) == (_this.F14()), 12)
		(_nw156).ArraySet1(!(_this.F14()) || (_this.F14()), 13)
		(_nw156).ArraySet1(false, 14)
		(_nw156).ArraySet1((true) || (false), 15)
		(_nw156).ArraySet1((func() bool {
			if _this.F14() {
				return _this.F14()
			}
			return (_this).F19()
		})(), 16)
		(_nw156).ArraySet1(_this.F14(), 17)
		(_nw156).ArraySet1(false, 18)
		(_nw156).ArraySet1(_this.F14(), 19)
		(_nw156).ArraySet1((_dafny.MultiSetOf(_959_v32)).IsDisjointFrom(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("lpfkgsqvw"))), 20)
		(_nw156).ArraySet1(!(_this.F14()), 21)
		_964_v36 = _nw156
		for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_964_v36), 0))); ; {
			_guard_loop_2, _ok45 := _iter45()
			if !_ok45 {
				break
			}
			var _965_i3 _dafny.Int
			_965_i3 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_965_i3).Sign() != -1) && ((_965_i3).Cmp(_dafny.ArrayLen((_964_v36), 0)) < 0)) {
				(_964_v36).ArraySet1((_dafny.SetOf(_this.F14(), !(true), _this.F14(), _this.F14())).IsSubsetOf((_dafny.SetOf(false, _this.F14())).Difference(_dafny.SetOf(true, _this.F14()))), (_965_i3).Int())
			}
		}
		var _966_v37 *C1
		_ = _966_v37
		var _nw157 *C1 = New_C1_()
		_ = _nw157
		_nw157.Ctor__()
		_966_v37 = _nw157
		var _967_v38 _dafny.Map
		_ = _967_v38
		_967_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_this).F19())
		var _968_v39 _dafny.Set
		_ = _968_v39
		_968_v39 = _dafny.SetOf(_954_v28, _dafny.IntOfInt64(608), (_967_v38).Cardinality(), _954_v28)
		r1 = !((_968_v39).IsSubsetOf(_968_v39))
		r0 = _dafny.IntOfUint32((_959_v32).Cardinality())
		var _969_v40 _dafny.CodePoint
		_ = _969_v40
		_969_v40 = _dafny.CodePoint('s')
		var _970_v41 _dafny.Map
		_ = _970_v41
		_970_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_969_v40, _954_v28)
		var _971_v42 _dafny.Set
		_ = _971_v42
		_971_v42 = _dafny.SetOf(_955_v29, _958_v31)
		r1 = Companion_Default___.Fm23((_954_v28).Times(_954_v28), (_this).F19(), (_970_v41).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_969_v40, _954_v28)), _971_v42, globalState)
		var _972_v43 _dafny.Map
		_ = _972_v43
		_972_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_954_v28, _this.F14())
		r2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_972_v43).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_954_v28, (_this).Fm2(_this.F14(), _this.F14(), globalState))), _954_v28)
		r3 = (_954_v28).Cmp(_954_v28) == 0
		return r0, r1, r2, r3
	}
}
func (_this *C4) M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		r1 = p1
		if !(!(!((_this).F19()) || (_this.F14()))) || ((_this).F19()) {
			var _973_v0 _dafny.Array
			_ = _973_v0
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_20
			var _nw158 _dafny.Array
			_ = _nw158
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw158 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Sequence = func(_974_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("mpehub")
				}
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw158 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw158).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw158).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_973_v0 = _nw158
			var _975_v2 _dafny.Map
			_ = _975_v2
			_975_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			var _rhs135 _dafny.Int = (func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter46 := _dafny.Iterate((_975_v2).Keys().Elements()); ; {
					_compr_43, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _976_v1 _dafny.Int
					_976_v1 = interface{}(_compr_43).(_dafny.Int)
					if (_975_v2).Contains(_976_v1) {
						_coll43.Add((_976_v1).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf((Companion_D2_.Create_DC9_(_this.F14(), p1, _dafny.CodePoint('q'))).Dtor_cf13()))).Cardinality()), _this.F14())
					}
				}
				return _coll43.ToMap()
			}()).Cardinality()
			_ = _rhs135
			var _rhs136 _dafny.Array = _973_v0
			_ = _rhs136
			r1 = _rhs135
			_973_v0 = _rhs136
			r1 = p1
			r2 = true
			var _977_v3 _dafny.Sequence
			_ = _977_v3
			_977_v3 = _dafny.UnicodeSeqOfUtf8Bytes("iycqg")
			var _978_v4 T0
			_ = _978_v4
			var _nw159 *C3 = New_C3_()
			_ = _nw159
			_nw159.Ctor__(p1, _977_v3, _this.F14(), (_this).F15())
			_978_v4 = _nw159
			var _979_v5 _dafny.Array
			_ = _979_v5
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_21
			var _nw160 _dafny.Array
			_ = _nw160
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw160 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_980_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_981_i1 _dafny.Int) _dafny.Int {
						return (_981_i1).Times(_dafny.IntOfUint32((_dafny.SeqOf(_980_p1)).Cardinality()))
					}
				})(p1)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw160 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw160).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw160).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_979_v5 = _nw160
			var _982_v6 _dafny.Map
			_ = _982_v6
			_982_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_979_v5, _978_v4)
			if !(_dafny.SetOf(_978_v4, (func() T0 {
				if (_982_v6).Contains(_979_v5) {
					return (_982_v6).Get(_979_v5).(T0)
				}
				return _978_v4
			})(), _978_v4, _978_v4)).Contains(_978_v4) {
				var _983_v7 _dafny.CodePoint
				_ = _983_v7
				_983_v7 = _dafny.CodePoint('v')
				var _984_v8 *C3
				_ = _984_v8
				var _nw161 *C3 = New_C3_()
				_ = _nw161
				_nw161.Ctor__(p1, _dafny.Companion_Sequence_.Update(_977_v3, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_977_v3).Cardinality()))).Uint32(), _983_v7), _this.F14(), (_978_v4).F15())
				_984_v8 = _nw161
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_973_v0), 0))
				_ = _index173
				(_973_v0).ArraySet1((_984_v8).F21(), (_index173).Int())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_973_v0), 0))
				_ = _index174
				(_973_v0).ArraySet1((func() _dafny.Sequence {
					if !((_this).F19()) {
						return (_984_v8).F21()
					}
					return (_984_v8).F21()
				})(), (_index174).Int())
				var _985_v9 _dafny.Set
				_ = _985_v9
				_985_v9 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(749))).Uint32(), func(coer71 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_986_v8 *C3) func(_dafny.Int) _dafny.Int {
					return func(_987_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32(((_986_v8).F21()).Cardinality())
					}
				})(_984_v8))))
				var _988_v10 _dafny.Sequence
				_ = _988_v10
				_988_v10 = _dafny.SeqOf((_this).F19())
				var _989_v11 _dafny.Sequence
				_ = _989_v11
				_989_v11 = _dafny.SeqOf((_988_v10).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_988_v10).Cardinality()))).Uint32()).(bool))
				var _990_v12 _dafny.Set
				_ = _990_v12
				_990_v12 = _dafny.SetOf((_984_v8).F21())
				r2 = (func() bool {
					if !(Companion_Default___.Fm23((_984_v8).F20(), _978_v4.F14(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_983_v7, p1), _985_v9, globalState)) {
						return _dafny.Companion_Sequence_.Equal(_989_v11, _dafny.Companion_Sequence_.Update(_988_v10, (Companion_Default___.SafeIndex((_984_v8).F20(), _dafny.IntOfUint32((_988_v10).Cardinality()))).Uint32(), (_this).F19()))
					}
					return (_990_v12).IsProperSubsetOf(Companion_Default___.Fm34(globalState))
				})()
				var _991_v13 _dafny.Array
				_ = _991_v13
				var _nw162 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw162
				_991_v13 = _nw162
				(globalState).F2 = _991_v13
				var _992_v14 _dafny.Array
				_ = _992_v14
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_22
				var _nw163 _dafny.Array
				_ = _nw163
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw163 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_993_v4 T0) func(_dafny.Int) _dafny.Sequence {
						return func(_994_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_993_v4.F14()), (Companion_Default___.SafeIndex((_dafny.SetOf(_this.F14())).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_993_v4.F14())).Cardinality()))).Uint32(), true)
						}
					})(_978_v4)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw163 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw163).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw163).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_992_v14 = _nw163
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_992_v14), 0))
				_ = _index175
				(_992_v14).ArraySet1(_dafny.SeqOf(true, _978_v4.F14(), _978_v4.F14()), (_index175).Int())
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_979_v5), 0))
				_ = _index176
				(_979_v5).ArraySet1(((_984_v8).F20()).Times(p1), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_992_v14), 0))
				_ = _index177
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_979_v5), 0))
				_ = _index178
				var _rhs137 _dafny.Sequence = _988_v10
				_ = _rhs137
				var _rhs138 _dafny.CodePoint = _983_v7
				_ = _rhs138
				var _rhs139 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.IntOfUint32((_977_v3).Cardinality())).Plus(p1)), (((_978_v4).F15()).Difference((_978_v4).F15())).Cardinality())
				_ = _rhs139
				var _rhs140 bool = _978_v4.F14()
				_ = _rhs140
				var _lhs116 _dafny.Array = _992_v14
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_992_v14), 0))
				_ = _lhs117
				var _lhs118 _dafny.Array = _979_v5
				_ = _lhs118
				var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_979_v5), 0))
				_ = _lhs119
				var _lhs120 *C4 = _this
				_ = _lhs120
				(_lhs116).ArraySet1(_rhs137, (_lhs117).Int())
				_983_v7 = _rhs138
				(_lhs118).ArraySet1(_rhs139, (_lhs119).Int())
				_lhs120.F14_set_(_rhs140)
			} else {
				r0 = (p1).Times((_dafny.Zero).Minus(p1))
				r0 = (_dafny.Zero).Minus((_dafny.IntOfInt64(621)).Minus((func() _dafny.Int {
					if _978_v4.F14() {
						return p1
					}
					return (_dafny.Zero).Minus(p1)
				})()))
				var _995_v15 _dafny.MultiSet
				_ = _995_v15
				_995_v15 = _dafny.MultiSetOf((_this).F19(), (p1).Cmp(p1) != 0, false, false, true)
				_995_v15 = (((_dafny.MultiSetOf(_978_v4.F14())).Update(_978_v4.F14(), Companion_Default___.Abs(p1))).Difference(_dafny.MultiSetOf(_978_v4.F14(), _this.F14(), (_this).F19(), (_this).F19()))).Difference(((_978_v4).F15()).Union(_995_v15))
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_979_v5), 0))
				_ = _index179
				(_979_v5).ArraySet1(p1, (_index179).Int())
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_979_v5), 0))
				_ = _index180
				(_979_v5).ArraySet1(p1, (_index180).Int())
				r2 = _978_v4.F14()
			}
			var _996_v16 _dafny.Array
			_ = _996_v16
			var _nw164 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(21))
			_ = _nw164
			_996_v16 = _nw164
			var _997_v17 D5
			_ = _997_v17
			_997_v17 = Companion_D5_.Create_DC18_(_978_v4.F14())
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_996_v16), 0))
			_ = _index181
			(_996_v16).ArraySet1((_997_v17).Dtor_cf33(), (_index181).Int())
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(334), _dafny.ArrayLen((_996_v16), 0))
			_ = _index182
			(_996_v16).ArraySet1(!((!((_this).F19())) && (_this.F14())), (_index182).Int())
		} else {
			(globalState).F9 = (p1).Cmp(_dafny.IntOfInt64(819)) > 0
			var _998_v18 _dafny.CodePoint
			_ = _998_v18
			_998_v18 = _dafny.CodePoint('c')
			var _999_v19 _dafny.Map
			_ = _999_v19
			_999_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_998_v18, p1)
			var _1000_v20 _dafny.Sequence
			_ = _1000_v20
			_1000_v20 = _dafny.SeqOf(p1, p1)
			var _1001_v21 _dafny.MultiSet
			_ = _1001_v21
			_1001_v21 = _dafny.MultiSetOf(_dafny.IntOfInt64(-711))
			var _1002_v22 _dafny.Map
			_ = _1002_v22
			_1002_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
			if !(!(_1002_v22).Contains(Companion_Default___.Fm0(true, p1, Companion_Default___.Fm0((_this).F19(), _dafny.IntOfInt64(877), p1, _1001_v21, globalState), _1001_v21, globalState))) || (Companion_Default___.Fm23(p1, _this.F14(), _999_v19, _dafny.SetOf(_1000_v20), globalState)) {
				(globalState).F3 = p1
				var _1003_v23 *C1
				_ = _1003_v23
				var _nw165 *C1 = New_C1_()
				_ = _nw165
				_nw165.Ctor__()
				_1003_v23 = _nw165
				_1003_v23 = _1003_v23
				var _1004_v24 _dafny.Sequence
				_ = _1004_v24
				_1004_v24 = _dafny.SeqOf(_this.F14(), (_this).F19(), _this.F14(), _this.F14())
				var _1005_v25 _dafny.MultiSet
				_ = _1005_v25
				_1005_v25 = _dafny.MultiSetOf(_1004_v24)
				var _1006_v26 _dafny.Sequence
				_ = _1006_v26
				_1006_v26 = _dafny.SeqOf(_1001_v21)
				r1 = (_this).Fm17(p1, _this.F14(), (_dafny.MultiSetOf(_1004_v24)).Difference(_1005_v25), _dafny.IntOfUint32((_1006_v26).Cardinality()), globalState)
				(_this).M6(globalState)
				var _1007_v27 _dafny.Array
				_ = _1007_v27
				var _nw166 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(6))
				_ = _nw166
				_1007_v27 = _nw166
				var _1008_v28 _dafny.Sequence
				_ = _1008_v28
				_1008_v28 = _dafny.SeqOf(_1007_v27, _1007_v27, _1007_v27, _1007_v27)
				r1 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1008_v28).Cardinality()))
			} else {
				var _1009_v29 _dafny.Sequence
				_ = _1009_v29
				_1009_v29 = _dafny.UnicodeSeqOfUtf8Bytes("erx")
				var _1010_v30 _dafny.Sequence
				_ = _1010_v30
				_1010_v30 = _dafny.SeqOf((_this).F19(), _this.F14())
				var _1011_v31 *C3
				_ = _1011_v31
				var _nw167 *C3 = New_C3_()
				_ = _nw167
				_nw167.Ctor__(p1, _1009_v29, !(_this.F14()), _dafny.MultiSetFromSeq(_1010_v30))
				_1011_v31 = _nw167
				var _rhs141 bool = (_dafny.IntOfInt64(-958)).Cmp(p1) != 0
				_ = _rhs141
				var _rhs142 _dafny.Sequence = _dafny.SeqOf(p1)
				_ = _rhs142
				var _rhs143 _dafny.Int = Companion_Default___.Fm0((_1011_v31).Fm2(_this.F14(), (_this).F19(), globalState), p1, (Companion_Default___.Fm35(p1, p1, globalState)).Cardinality(), _1001_v21, globalState)
				_ = _rhs143
				var _rhs144 bool = (_this).F19()
				_ = _rhs144
				var _lhs121 *GlobalState = globalState
				_ = _lhs121
				var _lhs122 *GlobalState = globalState
				_ = _lhs122
				r2 = _rhs141
				_1000_v20 = _rhs142
				_lhs121.F3 = _rhs143
				_lhs122.F9 = _rhs144
				r0 = p1
				(globalState).F3 = (_1011_v31).F20()
				var _1012_v32 *C3
				_ = _1012_v32
				var _nw168 *C3 = New_C3_()
				_ = _nw168
				_nw168.Ctor__((_1011_v31).F20(), _dafny.UnicodeSeqOfUtf8Bytes("oek"), _this.F14(), (_this).F15())
				_1012_v32 = _nw168
			}
			_998_v18 = _998_v18
			var _1013_v33 D0
			_ = _1013_v33
			_1013_v33 = Companion_D0_.Create_DC0_((_this).F19())
			var _1014_v34 D0
			_ = _1014_v34
			_1014_v34 = Companion_D0_.Create_DC2_(_1013_v33)
			var _1015_v35 D0
			_ = _1015_v35
			_1015_v35 = Companion_D0_.Create_DC2_(_1014_v34)
			var _source12 D0 = _1015_v35
			_ = _source12
			if _source12.Is_DC1() {
				var _1016___mcc_h0 bool = _source12.Get_().(D0_DC1).Cf1
				_ = _1016___mcc_h0
				var _1017_cf1 bool = _1016___mcc_h0
				_ = _1017_cf1
				var _1018_v36 _dafny.Sequence
				_ = _1018_v36
				_1018_v36 = _dafny.UnicodeSeqOfUtf8Bytes("ukldksby")
				var _1019_v37 _dafny.Map
				_ = _1019_v37
				_1019_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), p1)
				_998_v18 = Companion_Default___.Fm25(p1, Companion_Default___.SafeDivisionInt(p1, p1), _dafny.Companion_Sequence_.Contains(_1018_v36, _998_v18), _1019_v37, globalState)
				(globalState).F3 = p1
				(globalState).F9 = _1017_cf1
				var _1020_v38 _dafny.Array
				_ = _1020_v38
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_23
				var _nw169 _dafny.Array
				_ = _nw169
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw169 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) D3 = (func(_1021_cf1 bool) func(_dafny.Int) D3 {
						return func(_1022_i4 _dafny.Int) D3 {
							return Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D1_.Create_DC5_(_this.F14(), (_this).F19())).Dtor_cf6(), _1021_cf1))
						}
					})(_1017_cf1)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw169 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw169).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw169).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_1020_v38 = _nw169
				_1020_v38 = _1020_v38
			} else if _source12.Is_DC0() {
				var _1023___mcc_h1 bool = _source12.Get_().(D0_DC0).Cf0
				_ = _1023___mcc_h1
				var _1024_cf0 bool = _1023___mcc_h1
				_ = _1024_cf0
				(globalState).F3 = Companion_Default___.SafeDivisionInt(p1, (func() _dafny.Int {
					if (_1002_v22).Contains(p1) {
						return (_1002_v22).Get(p1).(_dafny.Int)
					}
					return p1
				})())
				var _1025_v39 _dafny.Map
				_ = _1025_v39
				_1025_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1024_cf0, p1)
				var _1026_v40 _dafny.Sequence
				_ = _1026_v40
				_1026_v40 = _dafny.SeqOf(_1025_v39)
				var _1027_v41 _dafny.Sequence
				_ = _1027_v41
				_1027_v41 = _dafny.SeqOf(Companion_Default___.Fm36(_1001_v21, p1, _1024_cf0, (_this).F19(), globalState))
				_1026_v40 = _dafny.Companion_Sequence_.Concatenate((_1027_v41).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_1027_v41).Cardinality()))).Uint32()).(_dafny.Sequence), _1026_v40)
				var _1028_v42 _dafny.Array
				_ = _1028_v42
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw170
				_1028_v42 = _nw170
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1028_v42), 0))
				_ = _index183
				(_1028_v42).ArraySet1(_this.F14(), (_index183).Int())
				var _1029_v43 _dafny.Sequence
				_ = _1029_v43
				_1029_v43 = _dafny.SeqOf(_1024_cf0)
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1028_v42), 0))
				_ = _index184
				(_1028_v42).ArraySet1((!((_this).F19()) || ((_this).F19())) && ((_1029_v43).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1029_v43).Cardinality()))).Uint32()).(bool)), (_index184).Int())
				(_this).M6(globalState)
			} else {
				var _1030___mcc_h2 D0 = _source12.Get_().(D0_DC2).Cf2
				_ = _1030___mcc_h2
				var _1031_cf2 D0 = _1030___mcc_h2
				_ = _1031_cf2
				var _1032_v44 _dafny.Array
				_ = _1032_v44
				var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
				_ = _nw171
				_1032_v44 = _nw171
				var _1033_v45 _dafny.Array
				_ = _1033_v45
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_24
				var _nw172 _dafny.Array
				_ = _nw172
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw172 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.Int = func(_1034_i5 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1034_i5, _dafny.IntOfInt64(588))
					}
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw172 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw172).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw172).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_1033_v45 = _nw172
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))
				_ = _index185
				(_1032_v44).ArraySet1(_1033_v45, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))
				_ = _index186
				(_1032_v44).ArraySet1(_1033_v45, (_index186).Int())
				var _1035_v46 D7
				_ = _1035_v46
				_1035_v46 = Companion_D7_.Create_DC23_(_this.F14(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1000_v20, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_1000_v20).Cardinality()))).Uint32(), p1)).Cardinality()), _this.F14())
				var _1036_v47 _dafny.Map
				_ = _1036_v47
				_1036_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1035_v46, (_this).Fm2((_this).F19(), (_this).F19(), globalState))
				_1036_v47 = (_1036_v47).Update(_1035_v46, (_this).F19())
				var _1037_v48 _dafny.Sequence
				_ = _1037_v48
				_1037_v48 = _dafny.SeqOf(_this.F14())
				var _1038_v49 _dafny.Array
				_ = _1038_v49
				var _len0_25 _dafny.Int = _dafny.One
				_ = _len0_25
				var _nw173 _dafny.Array
				_ = _nw173
				if _len0_25.Cmp(_dafny.Zero) == 0 {
					_nw173 = _dafny.NewArray(_len0_25)
				} else {
					var _init25 func(_dafny.Int) bool = func(_1039_i6 _dafny.Int) bool {
						return (_this).F19()
					}
					_ = _init25
					var _element0_25 = _init25(_dafny.Zero)
					_ = _element0_25
					_nw173 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
					(_nw173).ArraySet1(_element0_25, 0)
					var _nativeLen0_25 = (_len0_25).Int()
					_ = _nativeLen0_25
					for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
						(_nw173).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
					}
				}
				_1038_v49 = _nw173
				var _1040_v50 _dafny.Set
				_ = _1040_v50
				_1040_v50 = _dafny.SetOf((_this).F19())
				var _1041_v51 D8
				_ = _1041_v51
				_1041_v51 = Companion_D8_.Create_DC24_(_1040_v50)
				var _1042_v52 *C1
				_ = _1042_v52
				var _nw174 *C1 = New_C1_()
				_ = _nw174
				_nw174.Ctor__()
				_1042_v52 = _nw174
				var _1043_v53 _dafny.Map
				_ = _1043_v53
				_1043_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1042_v52)
				var _1044_v54 _dafny.Map
				_ = _1044_v54
				_1044_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1043_v53, _dafny.SetOf((_this).Fm2(true, (_this).F19(), globalState), !(false)))
				var _1045_v55 _dafny.MultiSet
				_ = _1045_v55
				_1045_v55 = _dafny.MultiSetOf(_1037_v48, _1037_v48)
				var _1046_v56 _dafny.Array
				_ = _1046_v56
				var _nwElement0_32 bool = (_this).F19()
				_ = _nwElement0_32
				var _nw175 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(25))
				_ = _nw175
				(_nw175).ArraySet1(_nwElement0_32, 0)
				(_nw175).ArraySet1((_this).F19(), 1)
				(_nw175).ArraySet1(_this.F14(), 2)
				(_nw175).ArraySet1(!((func() bool {
					if (_this).F19() {
						return (_this).F19()
					}
					return (_this).F19()
				})()), 3)
				(_nw175).ArraySet1(_this.F14(), 4)
				(_nw175).ArraySet1(_this.F14(), 5)
				(_nw175).ArraySet1(_this.F14(), 6)
				(_nw175).ArraySet1(!((_this).F19()), 7)
				(_nw175).ArraySet1(true, 8)
				(_nw175).ArraySet1((p1).Cmp(_dafny.IntOfUint32((_1037_v48).Cardinality())) == 0, 9)
				(_nw175).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1038_v49, _998_v18)).Contains(_1038_v49), 10)
				(_nw175).ArraySet1(!((_this).F19()), 11)
				(_nw175).ArraySet1(!(_this.F14()), 12)
				(_nw175).ArraySet1(((func() _dafny.Int {
					if (_1002_v22).Contains(p1) {
						return (_1002_v22).Get(p1).(_dafny.Int)
					}
					return p1
				})()).Cmp((p0).Cardinality()) <= 0, 13)
				(_nw175).ArraySet1((func() bool {
					if (_this).F19() {
						return (_this).F19()
					}
					return (_this).F19()
				})(), 14)
				(_nw175).ArraySet1((_this).F19(), 15)
				(_nw175).ArraySet1((_this).F19(), 16)
				(_nw175).ArraySet1(_this.F14(), 17)
				(_nw175).ArraySet1(((_1041_v51).Dtor_cf45()).IsDisjointFrom((func() _dafny.Set {
					if (_1044_v54).Contains(_1043_v53) {
						return (_1044_v54).Get(_1043_v53).(_dafny.Set)
					}
					return _1040_v50
				})()), 18)
				(_nw175).ArraySet1(true, 19)
				(_nw175).ArraySet1(!((func() bool {
					if (_this).F19() {
						return (_this).F19()
					}
					return (_this).Fm2((_this).F19(), _this.F14(), globalState)
				})()), 20)
				(_nw175).ArraySet1(!((_this).F19()) || (_this.F14()), 21)
				(_nw175).ArraySet1((p1).Cmp((_this).Fm17(p1, (_this).F19(), _1045_v55, _dafny.IntOfInt64(29), globalState)) == 0, 22)
				(_nw175).ArraySet1(_this.F14(), 23)
				(_nw175).ArraySet1((_this).F19(), 24)
				_1046_v56 = _nw175
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1038_v49), 0))
				_ = _index187
				(_1038_v49).ArraySet1((_this).F19(), (_index187).Int())
				var _1047_v57 _dafny.Array
				_ = _1047_v57
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
				_ = _nw176
				_1047_v57 = _nw176
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_1047_v57), 0))
				_ = _index188
				(_1047_v57).ArraySet1(_1038_v49, (_index188).Int())
				var _1048_v58 _dafny.Map
				_ = _1048_v58
				_1048_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1038_v49)
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1038_v49), 0))
				_ = _index189
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_1047_v57), 0))
				_ = _index190
				var _rhs145 bool = !(!(!(_this.F14())))
				_ = _rhs145
				var _rhs146 _dafny.Array = (func() _dafny.Array {
					if (_1048_v58).Contains((_this).F19()) {
						return (_1048_v58).Get((_this).F19()).(_dafny.Array)
					}
					return _1046_v56
				})()
				_ = _rhs146
				var _lhs123 _dafny.Array = _1038_v49
				_ = _lhs123
				var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1038_v49), 0))
				_ = _lhs124
				var _lhs125 _dafny.Array = _1047_v57
				_ = _lhs125
				var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_1047_v57), 0))
				_ = _lhs126
				(_lhs123).ArraySet1(_rhs145, (_lhs124).Int())
				(_lhs125).ArraySet1(_rhs146, (_lhs126).Int())
				var _1049_v59 _dafny.Map
				_ = _1049_v59
				_1049_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-263), false)
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))
				_ = _arr2
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))), 0))
				_ = _index191
				(_arr2).ArraySet1((func() _dafny.Int {
					if (_1001_v21).Contains(p1) {
						return (_1001_v21).Multiplicity(p1)
					}
					return (_1049_v59).Cardinality()
				})(), (_index191).Int())
				var _1050_v60 _dafny.Map
				_ = _1050_v60
				_1050_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1038_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1038_v49), 0))).Int()).(bool), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(258))).Uint32(), func(coer72 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg72 _dafny.Int) interface{} {
						return coer72(arg72)
					}
				}((func(_1051_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1052_i7 _dafny.Int) _dafny.Int {
						return _1051_p1
					}
				})(p1)))).Cardinality()))
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))
				_ = _arr3
				var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))), 0))
				_ = _index192
				var _rhs147 _dafny.Int = (func() _dafny.Int {
					if (_1050_v60).Contains((_this).F19()) {
						return (_1050_v60).Get((_this).F19()).(_dafny.Int)
					}
					return p1
				})()
				_ = _rhs147
				var _rhs148 _dafny.Int = (p1).Minus(p1)
				_ = _rhs148
				var _lhs127 *GlobalState = globalState
				_ = _lhs127
				var _lhs128 _dafny.Array = _dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))
				_ = _lhs128
				var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(932), _dafny.ArrayLen((_dafny.ArrayCastTo((_1032_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen((_1032_v44), 0))).Int()))), 0))
				_ = _lhs129
				_lhs127.F3 = _rhs147
				(_lhs128).ArraySet1(_rhs148, (_lhs129).Int())
			}
			var _1053_v61 _dafny.Array
			_ = _1053_v61
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_26
			var _nw177 _dafny.Array
			_ = _nw177
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw177 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) _dafny.Int = (func(_1054_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1055_i8 _dafny.Int) _dafny.Int {
						return (_1055_i8).Plus(_1054_p1)
					}
				})(p1)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw177 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw177).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw177).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_1053_v61 = _nw177
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1053_v61), 0))
			_ = _index193
			(_1053_v61).ArraySet1(p1, (_index193).Int())
			var _1056_v62 _dafny.Sequence
			_ = _1056_v62
			_1056_v62 = _dafny.SeqOf((_this).Fm2(_this.F14(), false, globalState))
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1053_v61), 0))
			_ = _index194
			var _rhs149 _dafny.Int = p1
			_ = _rhs149
			var _rhs150 _dafny.Int = p1
			_ = _rhs150
			var _rhs151 _dafny.Int = (p1).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1056_v62, _1056_v62)).Cardinality()))
			_ = _rhs151
			var _rhs152 bool = (_this).Fm2(_this.F14(), (_this).F19(), globalState)
			_ = _rhs152
			var _rhs153 _dafny.Int = (func() _dafny.Int {
				if (_this).Fm2(_this.F14(), _this.F14(), globalState) {
					return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1056_v62, _1056_v62)).Cardinality())
				}
				return p1
			})()
			_ = _rhs153
			var _lhs130 _dafny.Array = _1053_v61
			_ = _lhs130
			var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1053_v61), 0))
			_ = _lhs131
			var _lhs132 *GlobalState = globalState
			_ = _lhs132
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			r0 = _rhs149
			(_lhs130).ArraySet1(_rhs150, (_lhs131).Int())
			_lhs132.F3 = _rhs151
			r2 = _rhs152
			_lhs133.F3 = _rhs153
		}
		var _1057_v63 _dafny.Sequence
		_ = _1057_v63
		_1057_v63 = _dafny.UnicodeSeqOfUtf8Bytes("iep")
		_1057_v63 = _1057_v63
		if (_this).F19() {
			(globalState).F9 = (_this).F19()
			var _1058_v64 D4
			_ = _1058_v64
			_1058_v64 = Companion_D4_.Create_DC14_(false, (_this).F19())
			var _1059_v65 _dafny.Array
			_ = _1059_v65
			var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
			_ = _nw178
			_1059_v65 = _nw178
			var _1060_v66 _dafny.Map
			_ = _1060_v66
			_1060_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1059_v65, !((_this).F19()))
			var _1061_v67 D5
			_ = _1061_v67
			_1061_v67 = Companion_D5_.Create_DC16_(_1059_v65)
			var _1062_v68 _dafny.Set
			_ = _1062_v68
			_1062_v68 = _dafny.SetOf(false)
			var _1063_v69 _dafny.Array
			_ = _1063_v69
			var _nwElement0_33 bool = !((_1058_v64).Dtor_cf26())
			_ = _nwElement0_33
			var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(22))
			_ = _nw179
			(_nw179).ArraySet1(_nwElement0_33, 0)
			(_nw179).ArraySet1((_this).F19(), 1)
			(_nw179).ArraySet1(false, 2)
			(_nw179).ArraySet1((_this).F19(), 3)
			(_nw179).ArraySet1((func() bool {
				if true {
					return _this.F14()
				}
				return (_this).F19()
			})(), 4)
			(_nw179).ArraySet1((_this).Fm2(_this.F14(), (_this).F19(), globalState), 5)
			(_nw179).ArraySet1(_this.F14(), 6)
			(_nw179).ArraySet1(_this.F14(), 7)
			(_nw179).ArraySet1((_this).F19(), 8)
			(_nw179).ArraySet1(_this.F14(), 9)
			(_nw179).ArraySet1((_this).Fm2(_this.F14(), _this.F14(), globalState), 10)
			(_nw179).ArraySet1((func() bool {
				if (_this).F19() {
					return (_this).F19()
				}
				return (_this).F19()
			})(), 11)
			(_nw179).ArraySet1(((_this).F19()) == ((func() bool {
				if (_1060_v66).Contains((_1061_v67).Dtor_cf29()) {
					return (_1060_v66).Get((_1061_v67).Dtor_cf29()).(bool)
				}
				return _this.F14()
			})()), 12)
			(_nw179).ArraySet1(!(_this.F14()), 13)
			(_nw179).ArraySet1((_this).F19(), 14)
			(_nw179).ArraySet1((_this).F19(), 15)
			(_nw179).ArraySet1((_dafny.SetOf(!((_this).F19()), _this.F14(), (_this).F19(), _this.F14())).IsSubsetOf(_1062_v68), 16)
			(_nw179).ArraySet1((_this).F19(), 17)
			(_nw179).ArraySet1((_this).Fm2(_this.F14(), (_this).F19(), globalState), 18)
			(_nw179).ArraySet1(_this.F14(), 19)
			(_nw179).ArraySet1((p1).Cmp(p1) == 0, 20)
			(_nw179).ArraySet1(!((_this).F19()), 21)
			_1063_v69 = _nw179
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1063_v69), 0))
			_ = _index195
			(_1063_v69).ArraySet1(_this.F14(), (_index195).Int())
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1063_v69), 0))
			_ = _index196
			(_1063_v69).ArraySet1(_this.F14(), (_index196).Int())
			var _1064_v70 _dafny.Map
			_ = _1064_v70
			_1064_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1063_v69)
			_1064_v70 = (_1064_v70).Update((_this).F19(), _1063_v69)
			var _1065_v71 _dafny.Sequence
			_ = _1065_v71
			_1065_v71 = _dafny.SeqOf((_this).F15())
			(_this).M7(_1065_v71, (_1063_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1063_v69), 0))).Int()).(bool), p1, (_1063_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1063_v69), 0))).Int()).(bool), globalState)
			r2 = (p1).Cmp(p1) == 0
		} else {
			(globalState).F3 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-458), _dafny.IntOfInt64(715))
			var _1066_v72 _dafny.Set
			_ = _1066_v72
			_1066_v72 = _dafny.SetOf(true)
			if !(_1066_v72).Contains(_this.F14()) {
				var _1067_v73 _dafny.Array
				_ = _1067_v73
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(8))
				_ = _nw180
				_1067_v73 = _nw180
				_1067_v73 = _1067_v73
				var _1068_v74 _dafny.Sequence
				_ = _1068_v74
				_1068_v74 = _dafny.SeqOf(_this.F14(), (_this).F19())
				(globalState).F9 = (_1068_v74).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1068_v74).Cardinality()))).Uint32()).(bool)
				var _1069_v75 D0
				_ = _1069_v75
				_1069_v75 = Companion_D0_.Create_DC1_(!(_this.F14()))
				var _1070_v76 _dafny.Map
				_ = _1070_v76
				_1070_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(500), _dafny.SetOf((_this).F19(), _this.F14()))
				var _1071_v77 _dafny.Array
				_ = _1071_v77
				var _nwElement0_34 D0 = _1069_v75
				_ = _nwElement0_34
				var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(25))
				_ = _nw181
				(_nw181).ArraySet1(_nwElement0_34, 0)
				(_nw181).ArraySet1(_1069_v75, 1)
				(_nw181).ArraySet1(_1069_v75, 2)
				(_nw181).ArraySet1(_1069_v75, 3)
				(_nw181).ArraySet1(func(_pat_let27_0 D0) D0 {
					return func(_1072_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let28_0 bool) D0 {
							return func(_1073_dt__update_hcf1_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_1073_dt__update_hcf1_h0)
							}(_pat_let28_0)
						}((_this).F19())
					}(_pat_let27_0)
				}(_1069_v75), 4)
				(_nw181).ArraySet1(_1069_v75, 5)
				(_nw181).ArraySet1(_1069_v75, 6)
				(_nw181).ArraySet1(_1069_v75, 7)
				(_nw181).ArraySet1(_1069_v75, 8)
				(_nw181).ArraySet1(Companion_Default___.Fm37(p0, _1070_v76, (_this).F19(), (Companion_D0_.Create_DC1_(!(_this.F14()))).Dtor_cf1(), globalState), 9)
				(_nw181).ArraySet1(_1069_v75, 10)
				(_nw181).ArraySet1(_1069_v75, 11)
				(_nw181).ArraySet1(func(_pat_let29_0 D0) D0 {
					return func(_1074_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let30_0 bool) D0 {
							return func(_1075_dt__update_hcf1_h1 bool) D0 {
								return Companion_D0_.Create_DC1_(_1075_dt__update_hcf1_h1)
							}(_pat_let30_0)
						}(_this.F14())
					}(_pat_let29_0)
				}(Companion_Default___.Fm37(p0, _1070_v76, _this.F14(), (_this).F19(), globalState)), 12)
				(_nw181).ArraySet1(_1069_v75, 13)
				(_nw181).ArraySet1(_1069_v75, 14)
				(_nw181).ArraySet1(Companion_D0_.Create_DC1_(_this.F14()), 15)
				(_nw181).ArraySet1(_1069_v75, 16)
				(_nw181).ArraySet1(_1069_v75, 17)
				(_nw181).ArraySet1(_1069_v75, 18)
				(_nw181).ArraySet1(_1069_v75, 19)
				(_nw181).ArraySet1(_1069_v75, 20)
				(_nw181).ArraySet1(_1069_v75, 21)
				(_nw181).ArraySet1(_1069_v75, 22)
				(_nw181).ArraySet1(_1069_v75, 23)
				(_nw181).ArraySet1(func(_pat_let31_0 D0) D0 {
					return func(_1076_dt__update__tmp_h2 D0) D0 {
						return func(_pat_let32_0 bool) D0 {
							return func(_1077_dt__update_hcf1_h2 bool) D0 {
								return Companion_D0_.Create_DC1_(_1077_dt__update_hcf1_h2)
							}(_pat_let32_0)
						}(_this.F14())
					}(_pat_let31_0)
				}(_1069_v75), 24)
				_1071_v77 = _nw181
				_1071_v77 = _1071_v77
				(globalState).F9 = false
				_1057_v63 = _1057_v63
			} else {
				(globalState).F9 = (func() bool {
					if _this.F14() {
						return _this.F14()
					}
					return (_this).F19()
				})()
				var _1078_v78 _dafny.Array
				_ = _1078_v78
				var _len0_27 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_27
				var _nw182 _dafny.Array
				_ = _nw182
				if _len0_27.Cmp(_dafny.Zero) == 0 {
					_nw182 = _dafny.NewArray(_len0_27)
				} else {
					var _init27 func(_dafny.Int) bool = func(_1079_i9 _dafny.Int) bool {
						return ((_this).F15()).IsDisjointFrom((_this).F15())
					}
					_ = _init27
					var _element0_27 = _init27(_dafny.Zero)
					_ = _element0_27
					_nw182 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
					(_nw182).ArraySet1(_element0_27, 0)
					var _nativeLen0_27 = (_len0_27).Int()
					_ = _nativeLen0_27
					for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
						(_nw182).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
					}
				}
				_1078_v78 = _nw182
				(globalState).F2 = _1078_v78
				r2 = _this.F14()
				var _1080_v79 _dafny.Map
				_ = _1080_v79
				_1080_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _this.F14())
				var _1081_v80 _dafny.Set
				_ = _1081_v80
				_1081_v80 = _dafny.SetOf(_1080_v79, _1080_v79, _1080_v79)
				var _1082_v81 _dafny.Map
				_ = _1082_v81
				_1082_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1081_v80, (_dafny.Zero).Minus((p1).Minus(_dafny.IntOfInt64(48))))
				_1082_v81 = (_1082_v81).Update(_1081_v80, p1)
				var _1083_v82 _dafny.Array
				_ = _1083_v82
				var _nw183 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw183
				_1083_v82 = _nw183
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_1083_v82), 0))
				_ = _index197
				(_1083_v82).ArraySet1(Companion_Default___.SafeModuloInt(p1, p1), (_index197).Int())
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.ArrayLen((_1083_v82), 0))
				_ = _index198
				(_1083_v82).ArraySet1(_dafny.IntOfInt64(-600), (_index198).Int())
			}
			var _1084_v83 _dafny.CodePoint
			_ = _1084_v83
			_1084_v83 = _dafny.CodePoint('c')
			var _1085_v84 _dafny.Map
			_ = _1085_v84
			_1085_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1084_v83, p1)
			var _1086_v85 _dafny.Sequence
			_ = _1086_v85
			_1086_v85 = _dafny.SeqOf(p1)
			var _1087_v86 _dafny.Map
			_ = _1087_v86
			_1087_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1086_v85, (_this).F19())
			var _1088_v88 D8
			_ = _1088_v88
			_1088_v88 = Companion_D8_.Create_DC24_(_dafny.SetOf(_this.F14(), Companion_Default___.Fm23(p1, _this.F14(), _1085_v84, func() _dafny.Set {
				var _coll44 = _dafny.NewBuilder()
				_ = _coll44
				for _iter47 := _dafny.Iterate((_1087_v86).Keys().Elements()); ; {
					_compr_44, _ok47 := _iter47()
					if !_ok47 {
						break
					}
					var _1089_v87 _dafny.Sequence
					_1089_v87 = interface{}(_compr_44).(_dafny.Sequence)
					if (_1087_v86).Contains(_1089_v87) {
						_coll44.Add(_1089_v87)
					}
				}
				return _coll44.ToSet()
			}(), globalState)))
			_1088_v88 = _1088_v88
			var _1090_v89 _dafny.Array
			_ = _1090_v89
			var _len0_28 _dafny.Int = _dafny.One
			_ = _len0_28
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) _dafny.Int = func(_1091_i10 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1091_i10, _dafny.IntOfInt64(380))
				}
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw184 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw184).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw184).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1090_v89 = _nw184
			var _1092_v90 _dafny.Map
			_ = _1092_v90
			_1092_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1057_v63)
			var _1093_v91 D8
			_ = _1093_v91
			_1093_v91 = Companion_D8_.Create_DC28_(Companion_D8_.Create_DC26_(!(_this.F14()), _1092_v90))
			var _1094_v92 _dafny.MultiSet
			_ = _1094_v92
			_1094_v92 = _dafny.MultiSetOf(_1093_v91)
			var _1095_v93 _dafny.Map
			_ = _1095_v93
			_1095_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1090_v89, p1)
			var _1096_v94 _dafny.Map
			_ = _1096_v94
			_1096_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1095_v93)
			var _1097_v95 _dafny.Sequence
			_ = _1097_v95
			_1097_v95 = _dafny.SeqOf(_this.F14(), !(_this.F14()))
			var _1098_v96 _dafny.Sequence
			_ = _1098_v96
			_1098_v96 = _dafny.SeqOf(_1090_v89, _1090_v89)
			var _1099_v97 D3
			_ = _1099_v97
			_1099_v97 = Companion_D3_.Create_DC11_((_1094_v92).Cardinality(), ((func() _dafny.Map {
				if (_1096_v94).Contains((_1097_v95).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1097_v95).Cardinality()))).Uint32()).(bool)) {
					return (_1096_v94).Get((_1097_v95).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1097_v95).Cardinality()))).Uint32()).(bool)).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1098_v96).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1098_v96).Cardinality()))).Uint32()).(_dafny.Array), p1)
			})()).Cardinality(), (_this).F19(), p1)
			var _rhs154 _dafny.Array = _1090_v89
			_ = _rhs154
			var _rhs155 bool = (_this).F19()
			_ = _rhs155
			var _rhs156 D3 = _1099_v97
			_ = _rhs156
			var _rhs157 bool = !((_this).F19())
			_ = _rhs157
			var _lhs134 *GlobalState = globalState
			_ = _lhs134
			var _lhs135 *C4 = _this
			_ = _lhs135
			_1090_v89 = _rhs154
			_lhs134.F9 = _rhs155
			_1099_v97 = _rhs156
			_lhs135.F14_set_(_rhs157)
			(globalState).F9 = (_1097_v95).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(525), _dafny.IntOfUint32((_1097_v95).Cardinality()))).Uint32()).(bool)
		}
		if ((_this).F19()) && (!(_this.F14())) {
			var _1100_v98 _dafny.Sequence
			_ = _1100_v98
			_1100_v98 = _dafny.SeqOf((_this).F19())
			r1 = (((_dafny.MultiSetOf(_this.F14())).Intersection((_this).F15())).Union((_dafny.MultiSetFromSeq(_1100_v98)).Union((_this).F15()))).Cardinality()
			if !((_this).Fm2((p1).Cmp(p1) > 0, (_this).F19(), globalState)) {
				(globalState).F9 = (_this).F19()
				_1057_v63 = _1057_v63
				r1 = (p1).Plus(p1)
				var _1101_v99 _dafny.Array
				_ = _1101_v99
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_29
				var _nw185 _dafny.Array
				_ = _nw185
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw185 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Int = (func(_1102_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1103_i11 _dafny.Int) _dafny.Int {
							return (_1103_i11).Plus((_dafny.Zero).Minus(_1102_p1))
						}
					})(p1)
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw185 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw185).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw185).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1101_v99 = _nw185
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))
				_ = _index199
				(_1101_v99).ArraySet1((p1).Plus(p1), (_index199).Int())
				var _1104_v100 _dafny.Sequence
				_ = _1104_v100
				_1104_v100 = _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F19(), _this.F14())).Cardinality()), p1)
				var _1105_v101 _dafny.MultiSet
				_ = _1105_v101
				_1105_v101 = _dafny.MultiSetOf(p1)
				var _1106_v102 _dafny.Sequence
				_ = _1106_v102
				_1106_v102 = _dafny.SeqOf((p1).Times(Companion_Default___.Fm0(_this.F14(), _dafny.IntOfUint32((_1104_v100).Cardinality()), p1, _1105_v101, globalState)))
				var _1107_v103 _dafny.Array
				_ = _1107_v103
				var _nwElement0_35 bool = !((_this).F19())
				_ = _nwElement0_35
				var _nw186 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(9))
				_ = _nw186
				(_nw186).ArraySet1(_nwElement0_35, 0)
				(_nw186).ArraySet1((_this).F19(), 1)
				(_nw186).ArraySet1(_this.F14(), 2)
				(_nw186).ArraySet1(_this.F14(), 3)
				(_nw186).ArraySet1((_this).F19(), 4)
				(_nw186).ArraySet1((_this).F19(), 5)
				(_nw186).ArraySet1(!(_this.F14()), 6)
				(_nw186).ArraySet1(_this.F14(), 7)
				(_nw186).ArraySet1((_this).F19(), 8)
				_1107_v103 = _nw186
				var _1108_v104 _dafny.CodePoint
				_ = _1108_v104
				_1108_v104 = _dafny.CodePoint('n')
				var _1109_v105 _dafny.Map
				_ = _1109_v105
				_1109_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_v104, _dafny.IntOfUint32((_1057_v63).Cardinality()))
				var _1110_v106 _dafny.Set
				_ = _1110_v106
				_1110_v106 = _dafny.SetOf(_1104_v100)
				var _1111_v107 D7
				_ = _1111_v107
				_1111_v107 = Companion_D7_.Create_DC23_(Companion_Default___.Fm23(_dafny.IntOfUint32((_dafny.SeqOf(_1107_v103)).Cardinality()), (_this).F19(), _1109_v105, _1110_v106, globalState), p1, _this.F14())
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))
				_ = _index200
				var _rhs158 _dafny.Int = (_1106_v102).Select((Companion_Default___.SafeIndex((_1111_v107).Dtor_cf43(), _dafny.IntOfUint32((_1106_v102).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs158
				var _rhs159 _dafny.Int = p1
				_ = _rhs159
				var _lhs136 _dafny.Array = _1101_v99
				_ = _lhs136
				var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))
				_ = _lhs137
				(_lhs136).ArraySet1(_rhs158, (_lhs137).Int())
				r0 = _rhs159
				var _1112_v108 _dafny.Map
				_ = _1112_v108
				_1112_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, false)
				var _1113_v109 _dafny.Map
				_ = _1113_v109
				_1113_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1112_v108, (_1101_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))).Int()).(_dafny.Int))
				var _1114_v110 D7
				_ = _1114_v110
				_1114_v110 = Companion_D7_.Create_DC21_(_1113_v109)
				var _1115_v111 _dafny.Set
				_ = _1115_v111
				_1115_v111 = _dafny.SetOf(_this.F14(), _this.F14(), (_this).F19(), _this.F14(), _this.F14())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))
				_ = _index201
				var _rhs160 _dafny.Int = (((_dafny.SetOf(_this.F14(), false, _this.F14(), _this.F14())).Intersection(_dafny.SetOf((_this).F19()))).Intersection(_1115_v111)).Cardinality()
				_ = _rhs160
				var _rhs161 _dafny.Int = (_1101_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))).Int()).(_dafny.Int)
				_ = _rhs161
				var _rhs162 _dafny.Int = (_1101_v99).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))).Int()).(_dafny.Int)
				_ = _rhs162
				var _rhs163 D7 = _1114_v110
				_ = _rhs163
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 _dafny.Array = _1101_v99
				_ = _lhs139
				var _lhs140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_1101_v99), 0))
				_ = _lhs140
				_lhs138.F3 = _rhs160
				r0 = _rhs161
				(_lhs139).ArraySet1(_rhs162, (_lhs140).Int())
				_1114_v110 = _rhs163
			} else {
				var _1116_v112 _dafny.MultiSet
				_ = _1116_v112
				_1116_v112 = _dafny.MultiSetOf(Companion_D8_.Create_DC25_())
				var _1117_v113 D2
				_ = _1117_v113
				_1117_v113 = Companion_D2_.Create_DC7_(_1100_v98)
				var _rhs164 _dafny.Int = _dafny.IntOfUint32(((_1117_v113).Dtor_cf9()).Cardinality())
				_ = _rhs164
				var _rhs165 _dafny.Int = Companion_Default___.SafeModuloInt((p1).Plus(p1), p1)
				_ = _rhs165
				var _rhs166 _dafny.MultiSet = (_1116_v112).Intersection(_1116_v112)
				_ = _rhs166
				var _rhs167 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("e")
				_ = _rhs167
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				_lhs141.F3 = _rhs164
				r1 = _rhs165
				_1116_v112 = _rhs166
				_1057_v63 = _rhs167
				(globalState).F3 = p1
				var _1118_v114 _dafny.Sequence
				_ = _1118_v114
				_1118_v114 = _dafny.SeqOf(p1)
				var _1119_v115 _dafny.MultiSet
				_ = _1119_v115
				_1119_v115 = _dafny.MultiSetOf(p1)
				var _1120_v116 _dafny.Map
				_ = _1120_v116
				_1120_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_dafny.MultiSetFromSeq(_1118_v114)).Intersection(_1119_v115))
				_1120_v116 = (_1120_v116).Update(_this.F14(), (_1119_v115).Intersection(_1119_v115))
				var _1121_v117 D3
				_ = _1121_v117
				_1121_v117 = Companion_D3_.Create_DC11_(p1, p1, (_this).F19(), p1)
				var _1122_v118 *C3
				_ = _1122_v118
				var _nw187 *C3 = New_C3_()
				_ = _nw187
				_nw187.Ctor__(p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}(func(_1123_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('b')
				})), (_this).F19(), (_this).F15())
				_1122_v118 = _nw187
				var _1124_v119 _dafny.MultiSet
				_ = _1124_v119
				_1124_v119 = _dafny.MultiSetOf(_1122_v118, _1122_v118)
				var _1125_v120 _dafny.CodePoint
				_ = _1125_v120
				_1125_v120 = _dafny.CodePoint('q')
				var _1126_v121 _dafny.Set
				_ = _1126_v121
				_1126_v121 = _dafny.SetOf(_1125_v120, _1125_v120)
				var _1127_v122 _dafny.MultiSet
				_ = _1127_v122
				_1127_v122 = _dafny.MultiSetOf(_1126_v121)
				var _1128_v123 _dafny.Map
				_ = _1128_v123
				_1128_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_1127_v122).Cardinality())
				var _1129_v124 D8
				_ = _1129_v124
				_1129_v124 = Companion_D8_.Create_DC24_(_dafny.SetOf((_this).F19(), false))
				var _1130_v125 _dafny.Map
				_ = _1130_v125
				_1130_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1129_v124, _this.F14())
				var _1131_v126 _dafny.Map
				_ = _1131_v126
				_1131_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_this).F19())
				var _pat_let_tv23 = _1130_v125
				_ = _pat_let_tv23
				var _pat_let_tv24 = p1
				_ = _pat_let_tv24
				var _pat_let_tv25 = _1128_v123
				_ = _pat_let_tv25
				var _pat_let_tv26 = _1122_v118
				_ = _pat_let_tv26
				var _pat_let_tv27 = _1119_v115
				_ = _pat_let_tv27
				var _pat_let_tv28 = globalState
				_ = _pat_let_tv28
				var _pat_let_tv29 = _1122_v118
				_ = _pat_let_tv29
				var _pat_let_tv30 = _1119_v115
				_ = _pat_let_tv30
				var _pat_let_tv31 = globalState
				_ = _pat_let_tv31
				var _pat_let_tv32 = _1122_v118
				_ = _pat_let_tv32
				var _1132_v127 _dafny.Array
				_ = _1132_v127
				var _nwElement0_36 D3 = _1121_v117
				_ = _nwElement0_36
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(20))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_36, 0)
				(_nw188).ArraySet1(_1121_v117, 1)
				(_nw188).ArraySet1((func() D3 {
					if (_this).F19() {
						return Companion_D3_.Create_DC11_(p1, p1, (_this).F19(), _dafny.IntOfInt64(967))
					}
					return func(_pat_let33_0 D3) D3 {
						return func(_1133_dt__update__tmp_h3 D3) D3 {
							return func(_pat_let34_0 bool) D3 {
								return func(_1134_dt__update_hcf19_h0 bool) D3 {
									return Companion_D3_.Create_DC11_((_1133_dt__update__tmp_h3).Dtor_cf17(), (_1133_dt__update__tmp_h3).Dtor_cf18(), _1134_dt__update_hcf19_h0, (_1133_dt__update__tmp_h3).Dtor_cf20())
								}(_pat_let34_0)
							}(_this.F14())
						}(_pat_let33_0)
					}(_1121_v117)
				})(), 2)
				(_nw188).ArraySet1(Companion_D3_.Create_DC11_(p1, p1, _this.F14(), _dafny.IntOfInt64(485)), 3)
				(_nw188).ArraySet1(_1121_v117, 4)
				(_nw188).ArraySet1(_1121_v117, 5)
				(_nw188).ArraySet1(_1121_v117, 6)
				(_nw188).ArraySet1(Companion_D3_.Create_DC11_(p1, p1, true, (_dafny.Zero).Minus((_dafny.Zero).Minus((_1124_v119).Cardinality()))), 7)
				(_nw188).ArraySet1(_1121_v117, 8)
				(_nw188).ArraySet1(_1121_v117, 9)
				(_nw188).ArraySet1(_1121_v117, 10)
				(_nw188).ArraySet1(_1121_v117, 11)
				(_nw188).ArraySet1(_1121_v117, 12)
				(_nw188).ArraySet1(_1121_v117, 13)
				(_nw188).ArraySet1(func(_pat_let35_0 D3) D3 {
					return func(_1135_dt__update__tmp_h4 D3) D3 {
						return func(_pat_let36_0 bool) D3 {
							return func(_1136_dt__update_hcf19_h1 bool) D3 {
								return func(_pat_let37_0 _dafny.Int) D3 {
									return func(_1137_dt__update_hcf18_h0 _dafny.Int) D3 {
										return func(_pat_let38_0 _dafny.Int) D3 {
											return func(_1138_dt__update_hcf20_h0 _dafny.Int) D3 {
												return Companion_D3_.Create_DC11_((_1135_dt__update__tmp_h4).Dtor_cf17(), _1137_dt__update_hcf18_h0, _1136_dt__update_hcf19_h1, _1138_dt__update_hcf20_h0)
											}(_pat_let38_0)
										}(Companion_Default___.Fm0((_this).F19(), Companion_Default___.Fm0(_this.F14(), (_dafny.SetOf(_pat_let_tv24, (_pat_let_tv25).Cardinality())).Cardinality(), (_pat_let_tv26).F20(), _pat_let_tv27, _pat_let_tv28), (_pat_let_tv29).F20(), _pat_let_tv30, _pat_let_tv31))
									}(_pat_let37_0)
								}((_pat_let_tv23).Cardinality())
							}(_pat_let36_0)
						}((_this).F19())
					}(_pat_let35_0)
				}(Companion_D3_.Create_DC11_(p1, Companion_Default___.Fm0(true, Companion_Default___.Fm0((_this).F19(), _dafny.IntOfInt64(535), p1, _1119_v115, globalState), (_1122_v118).F20(), _dafny.MultiSetFromSeq(_1118_v114), globalState), _this.F14(), p1)), 14)
				(_nw188).ArraySet1(_1121_v117, 15)
				(_nw188).ArraySet1(_1121_v117, 16)
				(_nw188).ArraySet1(Companion_D3_.Create_DC11_((_1128_v123).Cardinality(), ((_1131_v126).Update((_this).F19(), (_this).F19())).Cardinality(), _this.F14(), (_1122_v118).F20()), 17)
				(_nw188).ArraySet1((func() D3 {
					if (_this).F19() {
						return _1121_v117
					}
					return func(_pat_let39_0 D3) D3 {
						return func(_1139_dt__update__tmp_h5 D3) D3 {
							return func(_pat_let40_0 _dafny.Int) D3 {
								return func(_1140_dt__update_hcf20_h1 _dafny.Int) D3 {
									return Companion_D3_.Create_DC11_((_1139_dt__update__tmp_h5).Dtor_cf17(), (_1139_dt__update__tmp_h5).Dtor_cf18(), (_1139_dt__update__tmp_h5).Dtor_cf19(), _1140_dt__update_hcf20_h1)
								}(_pat_let40_0)
							}((_pat_let_tv32).F20())
						}(_pat_let39_0)
					}(_1121_v117)
				})(), 18)
				(_nw188).ArraySet1(_1121_v117, 19)
				_1132_v127 = _nw188
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1132_v127), 0))
				_ = _index202
				(_1132_v127).ArraySet1(_1121_v117, (_index202).Int())
				var _1141_v128 _dafny.Array
				_ = _1141_v128
				var _len0_30 _dafny.Int = _dafny.One
				_ = _len0_30
				var _nw189 _dafny.Array
				_ = _nw189
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw189 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.MultiSet = func(_1142_i13 _dafny.Int) _dafny.MultiSet {
						return (_this).F15()
					}
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw189 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw189).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw189).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1141_v128 = _nw189
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1141_v128), 0))
				_ = _index203
				(_1141_v128).ArraySet1((_this).F15(), (_index203).Int())
				var _1143_v129 _dafny.Map
				_ = _1143_v129
				_1143_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm25(p1, _dafny.IntOfUint32((_1100_v98).Cardinality()), _this.F14(), _1128_v123, globalState), p1)
				var _1144_v130 _dafny.Set
				_ = _1144_v130
				_1144_v130 = _dafny.SetOf(_1118_v114)
				var _1145_v133 _dafny.Map
				_ = _1145_v133
				_1145_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v98, false)
				var _pat_let_tv33 = _1143_v129
				_ = _pat_let_tv33
				var _pat_let_tv34 = _1125_v120
				_ = _pat_let_tv34
				var _pat_let_tv35 = p1
				_ = _pat_let_tv35
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1132_v127), 0))
				_ = _index204
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1141_v128), 0))
				_ = _index205
				var _rhs168 D3 = func(_pat_let41_0 D3) D3 {
					return func(_1146_dt__update__tmp_h6 D3) D3 {
						return func(_pat_let42_0 bool) D3 {
							return func(_1147_dt__update_hcf19_h2 bool) D3 {
								return func(_pat_let43_0 _dafny.Int) D3 {
									return func(_1148_dt__update_hcf18_h1 _dafny.Int) D3 {
										return Companion_D3_.Create_DC11_((_1146_dt__update__tmp_h6).Dtor_cf17(), _1148_dt__update_hcf18_h1, _1147_dt__update_hcf19_h2, (_1146_dt__update__tmp_h6).Dtor_cf20())
									}(_pat_let43_0)
								}(_pat_let_tv35)
							}(_pat_let42_0)
						}(!(_pat_let_tv33).Contains(_pat_let_tv34))
					}(_pat_let41_0)
				}(_1121_v117)
				_ = _rhs168
				var _rhs169 _dafny.MultiSet = (_this).F15()
				_ = _rhs169
				var _rhs170 bool = (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_Default___.Fm23(_dafny.IntOfUint32(((_1122_v118).F21()).Cardinality()), Companion_Default___.Fm23((_1122_v118).F20(), (_this).F19(), _1143_v129, _1144_v130, globalState), func() _dafny.Map {
					var _coll45 = _dafny.NewMapBuilder()
					_ = _coll45
					for _iter48 := _dafny.Iterate(((_1122_v118).F21()).Elements()); ; {
						_compr_45, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1149_v131 _dafny.CodePoint
						_1149_v131 = interface{}(_compr_45).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains((_1122_v118).F21(), _1149_v131) {
							_coll45.Add(_1149_v131, p1)
						}
					}
					return _coll45.ToMap()
				}(), func() _dafny.Set {
					var _coll46 = _dafny.NewBuilder()
					_ = _coll46
					for _iter49 := _dafny.Iterate((_dafny.MultiSetOf(_1118_v114, _dafny.SeqOf(_dafny.IntOfInt64(-546)))).Elements()); ; {
						_compr_46, _ok49 := _iter49()
						if !_ok49 {
							break
						}
						var _1150_v132 _dafny.Sequence
						_1150_v132 = interface{}(_compr_46).(_dafny.Sequence)
						if (_dafny.MultiSetOf(_1118_v114, _dafny.SeqOf(_dafny.IntOfInt64(-546)))).Contains(_1150_v132) {
							_coll46.Add(_1150_v132)
						}
					}
					return _coll46.ToSet()
				}(), globalState), _this.F14()), (_this).F19())).Update(_dafny.SeqOf((_this).F19()), (_this).F19())).Update(_1100_v98, (_this).F19())).Equals(_1145_v133)
				_ = _rhs170
				var _lhs142 _dafny.Array = _1132_v127
				_ = _lhs142
				var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1132_v127), 0))
				_ = _lhs143
				var _lhs144 _dafny.Array = _1141_v128
				_ = _lhs144
				var _lhs145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_1141_v128), 0))
				_ = _lhs145
				var _lhs146 *GlobalState = globalState
				_ = _lhs146
				(_lhs142).ArraySet1(_rhs168, (_lhs143).Int())
				(_lhs144).ArraySet1(_rhs169, (_lhs145).Int())
				_lhs146.F9 = _rhs170
				var _1151_v134 _dafny.Sequence
				_ = _1151_v134
				_1151_v134 = _dafny.SeqOf((_1131_v126).Merge(_1131_v126), _1131_v126, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false), _1131_v126)
				_1151_v134 = _dafny.Companion_Sequence_.Concatenate(_1151_v134, _1151_v134)
			}
			var _1152_v135 D7
			_ = _1152_v135
			_1152_v135 = Companion_D7_.Create_DC23_((_this).F19(), p1, _this.F14())
			var _1153_v136 _dafny.Sequence
			_ = _1153_v136
			_1153_v136 = _dafny.SeqOf(_1100_v98, Companion_Default___.Fm38(globalState), _1100_v98)
			(globalState).F3 = ((_1152_v135).Dtor_cf43()).Times(_dafny.IntOfUint32((_1153_v136).Cardinality()))
			(globalState).F3 = p1
			var _1154_v137 _dafny.Array
			_ = _1154_v137
			var _nw190 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw190
			_1154_v137 = _nw190
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1154_v137), 0))
			_ = _index206
			(_1154_v137).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1100_v98).Cardinality())), (_index206).Int())
			var _1155_v138 _dafny.Set
			_ = _1155_v138
			_1155_v138 = _dafny.SetOf((_this).F19(), _this.F14())
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1154_v137), 0))
			_ = _index207
			var _rhs171 _dafny.Int = (_dafny.Zero).Minus(p1)
			_ = _rhs171
			var _rhs172 bool = (_1155_v138).IsProperSubsetOf(_1155_v138)
			_ = _rhs172
			var _rhs173 _dafny.Int = Companion_Default___.SafeDivisionInt(p1, p1)
			_ = _rhs173
			var _lhs147 *GlobalState = globalState
			_ = _lhs147
			var _lhs148 *C4 = _this
			_ = _lhs148
			var _lhs149 _dafny.Array = _1154_v137
			_ = _lhs149
			var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1154_v137), 0))
			_ = _lhs150
			_lhs147.F3 = _rhs171
			_lhs148.F14_set_(_rhs172)
			(_lhs149).ArraySet1(_rhs173, (_lhs150).Int())
		} else {
			var _1156_v139 _dafny.MultiSet
			_ = _1156_v139
			_1156_v139 = _dafny.MultiSetOf(((_this).F19()) || ((_this).F19()))
			var _1157_v140 _dafny.Array
			_ = _1157_v140
			var _nw191 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw191
			_1157_v140 = _nw191
			var _1158_v141 _dafny.Sequence
			_ = _1158_v141
			_1158_v141 = _dafny.SeqOf((_this).F19())
			var _1159_v142 D0
			_ = _1159_v142
			_1159_v142 = Companion_D0_.Create_DC1_(_this.F14())
			var _1160_v143 _dafny.Array
			_ = _1160_v143
			var _nwElement0_37 bool = _this.F14()
			_ = _nwElement0_37
			var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(12))
			_ = _nw192
			(_nw192).ArraySet1(_nwElement0_37, 0)
			(_nw192).ArraySet1(_this.F14(), 1)
			(_nw192).ArraySet1(_this.F14(), 2)
			(_nw192).ArraySet1(_this.F14(), 3)
			(_nw192).ArraySet1(!(_this.F14()), 4)
			(_nw192).ArraySet1((_1158_v141).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1158_v141).Cardinality()))).Uint32()).(bool), 5)
			(_nw192).ArraySet1((_this).F19(), 6)
			(_nw192).ArraySet1(_this.F14(), 7)
			(_nw192).ArraySet1((_1159_v142).Dtor_cf1(), 8)
			(_nw192).ArraySet1((_this).F19(), 9)
			(_nw192).ArraySet1((_this).F19(), 10)
			(_nw192).ArraySet1(_this.F14(), 11)
			_1160_v143 = _nw192
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1157_v140), 0))
			_ = _index208
			(_1157_v140).ArraySet1(_1160_v143, (_index208).Int())
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1157_v140), 0))
			_ = _index209
			var _rhs174 _dafny.MultiSet = (_1156_v139).Union((_this).F15())
			_ = _rhs174
			var _rhs175 _dafny.Array = (func() _dafny.Array {
				if false {
					return _1160_v143
				}
				return _1160_v143
			})()
			_ = _rhs175
			var _lhs151 _dafny.Array = _1157_v140
			_ = _lhs151
			var _lhs152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1157_v140), 0))
			_ = _lhs152
			_1156_v139 = _rhs174
			(_lhs151).ArraySet1(_rhs175, (_lhs152).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))
			_ = _index210
			(_1160_v143).ArraySet1((_this).F19(), (_index210).Int())
			var _1161_v144 _dafny.Set
			_ = _1161_v144
			_1161_v144 = _dafny.SetOf((_this).F19(), (_this).F19(), false)
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))
			_ = _index211
			(_1160_v143).ArraySet1(!(((_dafny.Zero).Minus((_dafny.Zero).Minus((_1161_v144).Cardinality()))).Cmp(_dafny.IntOfInt64(70)) >= 0) || (!(_this.F14())), (_index211).Int())
			var _1162_v145 _dafny.Map
			_ = _1162_v145
			_1162_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(709), p1)
			var _1163_v147 _dafny.Map
			_ = _1163_v147
			_1163_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1057_v63).Cardinality()), (_this).F19())
			var _1164_v148 _dafny.Sequence
			_ = _1164_v148
			_1164_v148 = _dafny.SeqOf((func() _dafny.Map {
				var _coll47 = _dafny.NewMapBuilder()
				_ = _coll47
				for _iter50 := _dafny.Iterate((p0).Elements()); ; {
					_compr_47, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _1165_v146 _dafny.Int
					_1165_v146 = interface{}(_compr_47).(_dafny.Int)
					if (p0).Contains(_1165_v146) {
						_coll47.Add((_1165_v146).Plus(p1), p1)
					}
				}
				return _coll47.ToMap()
			}()).Update(_dafny.IntOfInt64(889), (_1163_v147).Cardinality()), _1162_v145)
			var _rhs176 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_1057_v63).Cardinality())), p1)
			_ = _rhs176
			var _rhs177 _dafny.Map = (_1162_v145).Merge((_1164_v148).Select((Companion_Default___.SafeIndex((func() _dafny.Map {
				var _coll48 = _dafny.NewMapBuilder()
				_ = _coll48
				for _iter51 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(581), _dafny.IntOfInt64(439))); ; {
					_compr_48, _ok51 := _iter51()
					if !_ok51 {
						break
					}
					var _1166_v149 _dafny.Int
					_1166_v149 = interface{}(_compr_48).(_dafny.Int)
					if ((_dafny.IntOfInt64(581)).Cmp(_1166_v149) <= 0) && ((_1166_v149).Cmp(_dafny.IntOfInt64(439)) < 0) {
						_coll48.Add(Companion_Default___.SafeDivisionInt(_1166_v149, (_1161_v144).Cardinality()), (_this).F19())
					}
				}
				return _coll48.ToMap()
			}()).Cardinality(), _dafny.IntOfUint32((_1164_v148).Cardinality()))).Uint32()).(_dafny.Map))
			_ = _rhs177
			var _rhs178 _dafny.Sequence = _1057_v63
			_ = _rhs178
			var _lhs153 *GlobalState = globalState
			_ = _lhs153
			_lhs153.F3 = _rhs176
			_1162_v145 = _rhs177
			_1057_v63 = _rhs178
			var _1167_v150 _dafny.MultiSet
			_ = _1167_v150
			_1167_v150 = _dafny.MultiSetOf(p1, p1)
			var _1168_v151 D1
			_ = _1168_v151
			_1168_v151 = Companion_D1_.Create_DC4_(Companion_Default___.Fm0((_1160_v143).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))).Int()).(bool), p1, p1, _1167_v150, globalState), p1)
			var _1169_v153 D3
			_ = _1169_v153
			_1169_v153 = Companion_D3_.Create_DC11_((_1168_v151).Dtor_cf5(), (func() _dafny.Map {
				var _coll49 = _dafny.NewMapBuilder()
				_ = _coll49
				for _iter52 := _dafny.Iterate((_1167_v150).Elements()); ; {
					_compr_49, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _1170_v152 _dafny.Int
					_1170_v152 = interface{}(_compr_49).(_dafny.Int)
					if (_1167_v150).Contains(_1170_v152) {
						_coll49.Add((_1170_v152).Times(p1), _1162_v145)
					}
				}
				return _coll49.ToMap()
			}()).Cardinality(), _this.F14(), p1)
			var _source13 D3 = _1169_v153
			_ = _source13
			if _source13.Is_DC11() {
				var _1171___mcc_h3 _dafny.Int = _source13.Get_().(D3_DC11).Cf17
				_ = _1171___mcc_h3
				var _1172___mcc_h4 _dafny.Int = _source13.Get_().(D3_DC11).Cf18
				_ = _1172___mcc_h4
				var _1173___mcc_h5 bool = _source13.Get_().(D3_DC11).Cf19
				_ = _1173___mcc_h5
				var _1174___mcc_h6 _dafny.Int = _source13.Get_().(D3_DC11).Cf20
				_ = _1174___mcc_h6
				var _1175_cf20 _dafny.Int = _1174___mcc_h6
				_ = _1175_cf20
				var _1176_cf19 bool = _1173___mcc_h5
				_ = _1176_cf19
				var _1177_cf18 _dafny.Int = _1172___mcc_h4
				_ = _1177_cf18
				var _1178_cf17 _dafny.Int = _1171___mcc_h3
				_ = _1178_cf17
				var _1179_v154 _dafny.Map
				_ = _1179_v154
				_1179_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_1057_v63, _1057_v63), _1162_v145)
				var _1180_v155 _dafny.CodePoint
				_ = _1180_v155
				_1180_v155 = _dafny.CodePoint('d')
				var _1181_v156 D9
				_ = _1181_v156
				_1181_v156 = Companion_D9_.Create_DC29_(_1162_v145)
				_1179_v154 = (_1179_v154).Update(_dafny.Companion_Sequence_.Update(_1057_v63, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(345))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg74 _dafny.Int) interface{} {
						return coer74(arg74)
					}
				}(func(_1182_i14 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()), _dafny.IntOfUint32((_1057_v63).Cardinality()))).Uint32(), _1180_v155), (_1181_v156).Dtor_cf49())
				r2 = (func() bool {
					if (_1163_v147).Contains((_1175_cf20).Times(_1178_cf17)) {
						return (_1163_v147).Get((_1175_cf20).Times(_1178_cf17)).(bool)
					}
					return _1176_cf19
				})()
				var _1183_v157 _dafny.Set
				_ = _1183_v157
				_1183_v157 = _dafny.SetOf(_1180_v155, _1180_v155)
				var _1184_v158 _dafny.Sequence
				_ = _1184_v158
				_1184_v158 = _dafny.SeqOf(_1183_v157, (_1183_v157).Difference(_1183_v157))
				_1184_v158 = _1184_v158
				(_this).F14_set_(!(_1176_cf19) || (!_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(137))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1185_v155 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1186_i15 _dafny.Int) _dafny.CodePoint {
						return _1185_v155
					}
				})(_1180_v155))), _1057_v63)))
			} else {
				var _1187___mcc_h7 _dafny.Map = _source13.Get_().(D3_DC10).Cf16
				_ = _1187___mcc_h7
				var _1188_cf16 _dafny.Map = _1187___mcc_h7
				_ = _1188_cf16
				(_this).M6(globalState)
				var _1189_v159 _dafny.Array
				_ = _1189_v159
				var _nwElement0_38 _dafny.Int = (p1).Minus(p1)
				_ = _nwElement0_38
				var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.One)
				_ = _nw193
				(_nw193).ArraySet1(_nwElement0_38, 0)
				_1189_v159 = _nw193
				var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1189_v159), 0))
				_ = _index212
				(_1189_v159).ArraySet1(_dafny.IntOfInt64(966), (_index212).Int())
				var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1189_v159), 0))
				_ = _index213
				(_1189_v159).ArraySet1(((p1).Plus(Companion_Default___.Fm0(_this.F14(), p1, p1, Companion_Default___.Fm30((_this).F19(), globalState), globalState))).Minus(p1), (_index213).Int())
				r2 = true
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))
				_ = _index214
				var _rhs179 bool = ((_1189_v159).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1189_v159), 0))).Int()).(_dafny.Int)).Cmp((_1168_v151).Dtor_cf4()) >= 0
				_ = _rhs179
				var _rhs180 _dafny.Array = _1160_v143
				_ = _rhs180
				var _rhs181 _dafny.Int = _dafny.IntOfInt64(-942)
				_ = _rhs181
				var _rhs182 _dafny.Int = (p1).Times((_1189_v159).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(918), _dafny.ArrayLen((_1189_v159), 0))).Int()).(_dafny.Int))
				_ = _rhs182
				var _lhs154 _dafny.Array = _1160_v143
				_ = _lhs154
				var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))
				_ = _lhs155
				var _lhs156 *GlobalState = globalState
				_ = _lhs156
				var _lhs157 *GlobalState = globalState
				_ = _lhs157
				(_lhs154).ArraySet1(_rhs179, (_lhs155).Int())
				_lhs156.F2 = _rhs180
				r0 = _rhs181
				_lhs157.F3 = _rhs182
			}
			var _1190_v160 _dafny.Map
			_ = _1190_v160
			_1190_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_1157_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1157_v140), 0))).Int())), p1)
			var _1191_v161 D4
			_ = _1191_v161
			_1191_v161 = Companion_D4_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(671))).Uint32(), func(coer76 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg76 _dafny.Int) interface{} {
					return coer76(arg76)
				}
			}(func(_1192_i16 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('p')
			})), _dafny.IntOfInt64(786), (_this).F19(), _dafny.ArrayCastTo((_1157_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(985), _dafny.ArrayLen((_1157_v140), 0))).Int())))
			var _1193_v162 _dafny.MultiSet
			_ = _1193_v162
			_1193_v162 = _dafny.MultiSetOf(_dafny.SeqOf(_this.F14(), (_this).F19()))
			var _rhs183 _dafny.Int = ((Companion_Default___.Fm30((_1160_v143).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_1160_v143), 0))).Int()).(bool), globalState)).Cardinality()).Plus(p1)
			_ = _rhs183
			var _rhs184 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if ((_dafny.Zero).Minus((_1190_v160).Cardinality())).Cmp(p1) != 0 {
					return Companion_Default___.SafeDivisionInt(p1, p1)
				}
				return p1
			})())
			_ = _rhs184
			var _rhs185 _dafny.Int = ((p1).Times(p1)).Minus((_dafny.Zero).Minus(p1))
			_ = _rhs185
			var _rhs186 _dafny.Int = (p1).Plus(((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_1191_v161).Dtor_cf22()).Cardinality())))).Times((_dafny.Zero).Minus(p1)))
			_ = _rhs186
			var _rhs187 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_this).Fm17(p1, _this.F14(), _1193_v162, p1, globalState), p1), p1)
			_ = _rhs187
			var _lhs158 *GlobalState = globalState
			_ = _lhs158
			r1 = _rhs183
			_lhs158.F3 = _rhs184
			r1 = _rhs185
			r0 = _rhs186
			r1 = _rhs187
		}
		var _1194_v163 _dafny.Array
		_ = _1194_v163
		var _nw194 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw194
		_1194_v163 = _nw194
		var _1195_v164 _dafny.Sequence
		_ = _1195_v164
		_1195_v164 = _dafny.SeqOf(_1194_v163)
		var _1196_v165 D4
		_ = _1196_v165
		_1196_v165 = Companion_D4_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("uoqpauedn"), p1, _this.F14(), (_1195_v164).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1195_v164).Cardinality()))).Uint32()).(_dafny.Array))
		var _source14 D4 = func(_pat_let44_0 D4) D4 {
			return func(_1197_dt__update__tmp_h7 D4) D4 {
				return func(_pat_let45_0 _dafny.Sequence) D4 {
					return func(_1199_dt__update_hcf22_h0 _dafny.Sequence) D4 {
						return Companion_D4_.Create_DC13_(_1199_dt__update_hcf22_h0, (_1197_dt__update__tmp_h7).Dtor_cf23(), (_1197_dt__update__tmp_h7).Dtor_cf24(), (_1197_dt__update__tmp_h7).Dtor_cf25())
					}(_pat_let45_0)
				}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(431))).Uint32(), func(coer77 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}(func(_1198_i17 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				})))
			}(_pat_let44_0)
		}(_1196_v165)
		_ = _source14
		if _source14.Is_DC13() {
			var _1200___mcc_h8 _dafny.Sequence = _source14.Get_().(D4_DC13).Cf22
			_ = _1200___mcc_h8
			var _1201___mcc_h9 _dafny.Int = _source14.Get_().(D4_DC13).Cf23
			_ = _1201___mcc_h9
			var _1202___mcc_h10 bool = _source14.Get_().(D4_DC13).Cf24
			_ = _1202___mcc_h10
			var _1203___mcc_h11 _dafny.Array = _source14.Get_().(D4_DC13).Cf25
			_ = _1203___mcc_h11
			var _1204_cf25 _dafny.Array = _1203___mcc_h11
			_ = _1204_cf25
			var _1205_cf24 bool = _1202___mcc_h10
			_ = _1205_cf24
			var _1206_cf23 _dafny.Int = _1201___mcc_h9
			_ = _1206_cf23
			var _1207_cf22 _dafny.Sequence = _1200___mcc_h8
			_ = _1207_cf22
			var _1208_v166 _dafny.Map
			_ = _1208_v166
			_1208_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1206_cf23, _1205_cf24)
			var _1209_v167 _dafny.Map
			_ = _1209_v167
			_1209_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1057_v63).Cardinality()), _1208_v166)
			_1209_v167 = (_1209_v167).Update(Companion_Default___.SafeDivisionInt(_1206_cf23, _1206_cf23), _1208_v166)
			var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))
			_ = _index215
			(_1204_cf25).ArraySet1(_1205_cf24, (_index215).Int())
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))
			_ = _index216
			(_1204_cf25).ArraySet1((_this).Fm2((_this).F19(), _this.F14(), globalState), (_index216).Int())
			if (_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool) {
				var _1210_v168 _dafny.Array
				_ = _1210_v168
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_31
				var _nw195 _dafny.Array
				_ = _nw195
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw195 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) _dafny.Int = (func(_1211_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1212_i18 _dafny.Int) _dafny.Int {
							return (_1212_i18).Plus(_1211_p1)
						}
					})(p1)
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw195 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw195).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw195).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1210_v168 = _nw195
				var _1213_v169 _dafny.Map
				_ = _1213_v169
				_1213_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1210_v168)
				(globalState).F3 = (_dafny.Zero).Minus((((_1213_v169).Merge(_1213_v169)).Merge(_1213_v169)).Cardinality())
				var _1214_v170 *C3
				_ = _1214_v170
				var _nw196 *C3 = New_C3_()
				_ = _nw196
				_nw196.Ctor__(p1, _dafny.Companion_Sequence_.Update(_1057_v63, (Companion_Default___.SafeIndex(_1206_cf23, _dafny.IntOfUint32((_1057_v63).Cardinality()))).Uint32(), _dafny.CodePoint('t')), false, (_this).F15())
				_1214_v170 = _nw196
				_1214_v170 = _1214_v170
				var _1215_v171 *C3
				_ = _1215_v171
				var _nw197 *C3 = New_C3_()
				_ = _nw197
				_nw197.Ctor__((p1).Plus((_1214_v170).F20()), _dafny.Companion_Sequence_.Concatenate(_1057_v63, _1207_cf22), false, (_this).F15())
				_1215_v171 = _nw197
				(globalState).F3 = ((_this).F15()).Cardinality()
				var _1216_v172 _dafny.CodePoint
				_ = _1216_v172
				_1216_v172 = _dafny.CodePoint('g')
				var _1217_v173 D9
				_ = _1217_v173
				_1217_v173 = Companion_D9_.Create_DC30_(_1216_v172, _1205_cf24, p1, p0, (_1215_v171).F20())
				var _1218_v174 _dafny.Sequence
				_ = _1218_v174
				_1218_v174 = _dafny.SeqOf(p1)
				var _1219_v175 D2
				_ = _1219_v175
				_1219_v175 = Companion_D2_.Create_DC9_((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool), (_1215_v171).F20(), _1216_v172)
				var _1220_v176 _dafny.Sequence
				_ = _1220_v176
				_1220_v176 = _dafny.SeqOf(_this.F14())
				var _1221_v177 _dafny.Set
				_ = _1221_v177
				_1221_v177 = _dafny.SetOf(_1057_v63)
				var _pat_let_tv36 = _1216_v172
				_ = _pat_let_tv36
				var _1222_v179 _dafny.Array
				_ = _1222_v179
				var _nwElement0_39 D9 = Companion_D9_.Create_DC30_(_1216_v172, (_this).F19(), _1206_cf23, p0, p1)
				_ = _nwElement0_39
				var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(21))
				_ = _nw198
				(_nw198).ArraySet1(_nwElement0_39, 0)
				(_nw198).ArraySet1(_1217_v173, 1)
				(_nw198).ArraySet1(_1217_v173, 2)
				(_nw198).ArraySet1(_1217_v173, 3)
				(_nw198).ArraySet1(func(_pat_let46_0 D9) D9 {
					return func(_1223_dt__update__tmp_h8 D9) D9 {
						return func(_pat_let47_0 _dafny.CodePoint) D9 {
							return func(_1224_dt__update_hcf50_h0 _dafny.CodePoint) D9 {
								return Companion_D9_.Create_DC30_(_1224_dt__update_hcf50_h0, (_1223_dt__update__tmp_h8).Dtor_cf51(), (_1223_dt__update__tmp_h8).Dtor_cf52(), (_1223_dt__update__tmp_h8).Dtor_cf53(), (_1223_dt__update__tmp_h8).Dtor_cf54())
							}(_pat_let47_0)
						}(_pat_let_tv36)
					}(_pat_let46_0)
				}(_1217_v173), 4)
				(_nw198).ArraySet1(_1217_v173, 5)
				(_nw198).ArraySet1(_1217_v173, 6)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_(_1216_v172, _1205_cf24, (_1214_v170).F20(), p0, _dafny.IntOfInt64(746)), 7)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_(_1216_v172, (_this).F19(), (_1215_v171).F20(), p0, p1), 8)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_(_1216_v172, false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_1218_v174, _1218_v174)).Cardinality())), _dafny.SetOf((_1215_v171).F20(), _dafny.IntOfInt64(5)), (_1215_v171).F20()), 9)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_((_1219_v175).Dtor_cf15(), (_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool), (_1214_v170).F20(), p0, (_1214_v170).F20()), 10)
				(_nw198).ArraySet1(_1217_v173, 11)
				(_nw198).ArraySet1(Companion_Default___.Fm39(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1220_v176, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1220_v176).Cardinality()))).Uint32(), (_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool))).Cardinality()), (_1214_v170).F20(), p1, (_1221_v177).Cardinality(), globalState), 12)
				(_nw198).ArraySet1(_1217_v173, 13)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_(_1216_v172, (_this).F19(), (_1214_v170).F20(), p0, p1), 14)
				(_nw198).ArraySet1(_1217_v173, 15)
				(_nw198).ArraySet1(Companion_D9_.Create_DC30_(_1216_v172, (_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool), (func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter53 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-580))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg78 _dafny.Int) interface{} {
							return coer78(arg78)
						}
					}((func(_1225_v172 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1226_i19 _dafny.Int) _dafny.CodePoint {
							return _1225_v172
						}
					})(_1216_v172)))).Elements()); ; {
						_compr_50, _ok53 := _iter53()
						if !_ok53 {
							break
						}
						var _1227_v178 _dafny.CodePoint
						_1227_v178 = interface{}(_compr_50).(_dafny.CodePoint)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-580))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg79 _dafny.Int) interface{} {
								return coer79(arg79)
							}
						}((func(_1228_v172 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1226_i19 _dafny.Int) _dafny.CodePoint {
								return _1228_v172
							}
						})(_1216_v172))), _1227_v178) {
							_coll50.Add(_1227_v178, false)
						}
					}
					return _coll50.ToMap()
				}()).Cardinality(), _dafny.SetOf(_dafny.IntOfUint32(((_1214_v170).F21()).Cardinality())), (_1215_v171).F20()), 16)
				(_nw198).ArraySet1(_1217_v173, 17)
				(_nw198).ArraySet1(_1217_v173, 18)
				(_nw198).ArraySet1(_1217_v173, 19)
				(_nw198).ArraySet1(_1217_v173, 20)
				_1222_v179 = _nw198
				var _1229_v180 _dafny.Map
				_ = _1229_v180
				_1229_v180 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1222_v179, (_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool))
				var _rhs188 bool = (func() bool {
					if (_1229_v180).Contains(_1222_v179) {
						return (_1229_v180).Get(_1222_v179).(bool)
					}
					return !_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(588))).Uint32(), func(coer80 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg80 _dafny.Int) interface{} {
							return coer80(arg80)
						}
					}((func(_1230_cf24 bool) func(_dafny.Int) _dafny.Int {
						return func(_1231_i20 _dafny.Int) _dafny.Int {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1230_cf24, _1230_cf24)).Cardinality()
						}
					})(_1205_cf24))), (Companion_D1_.Create_DC4_((_1214_v170).F20(), (_1214_v170).F20())).Dtor_cf4())
				})()
				_ = _rhs188
				var _rhs189 bool = (_1215_v171).Fm2(false, _1205_cf24, globalState)
				_ = _rhs189
				r2 = _rhs188
				_1205_cf24 = _rhs189
			} else {
				(_this).F14_set_(((p1).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qd")).Cardinality()))).Cmp(_1206_cf23) != 0)
				var _1232_v181 D4
				_ = _1232_v181
				_1232_v181 = Companion_D4_.Create_DC14_(_this.F14(), (_this).F19())
				var _1233_v182 D4
				_ = _1233_v182
				_1233_v182 = Companion_D4_.Create_DC15_(_1232_v181)
				var _1234_v183 _dafny.Map
				_ = _1234_v183
				_1234_v183 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v182, _1206_cf23)
				var _1235_v184 _dafny.Set
				_ = _1235_v184
				_1235_v184 = _dafny.SetOf(_this.F14(), _this.F14())
				var _1236_v185 _dafny.Map
				_ = _1236_v185
				_1236_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool), _1235_v184)
				var _1237_v186 _dafny.Map
				_ = _1237_v186
				_1237_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1234_v183).Update(Companion_D4_.Create_DC15_(_1232_v181), (_1236_v185).Cardinality()))
				_1237_v186 = (_1237_v186).Update(_1206_cf23, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1233_v182, p1))
				var _1238_v187 _dafny.Sequence
				_ = _1238_v187
				_1238_v187 = _dafny.SeqOf((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool))
				var _1239_v188 _dafny.MultiSet
				_ = _1239_v188
				_1239_v188 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1057_v63, (_this).Fm3(_dafny.IntOfUint32((_1238_v187).Cardinality()), _1206_cf23, !((_this).F19()), _1206_cf23, globalState)))
				_1239_v188 = _dafny.MultiSetOf(_1207_cf22)
				var _1240_v189 _dafny.Sequence
				_ = _1240_v189
				_1240_v189 = _dafny.SeqOf(_dafny.IntOfUint32((_1057_v63).Cardinality()), p1)
				var _1241_v190 _dafny.Map
				_ = _1241_v190
				_1241_v190 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1204_cf25, _dafny.IntOfUint32((_1240_v189).Cardinality()))
				var _1242_v191 _dafny.CodePoint
				_ = _1242_v191
				_1242_v191 = _dafny.CodePoint('y')
				var _1243_v192 _dafny.Map
				_ = _1243_v192
				_1243_v192 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1242_v191, _1206_cf23)
				var _1244_v193 _dafny.Set
				_ = _1244_v193
				_1244_v193 = _dafny.SetOf(_1240_v189)
				var _1245_v194 _dafny.Sequence
				_ = _1245_v194
				_1245_v194 = _dafny.SeqOf(_1244_v193)
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))
				_ = _index217
				(_1204_cf25).ArraySet1(!(Companion_Default___.Fm23(p1, (func() bool {
					if (_1208_v166).Contains((_1241_v190).Cardinality()) {
						return (_1208_v166).Get((_1241_v190).Cardinality()).(bool)
					}
					return _this.F14()
				})(), _1243_v192, (_1245_v194).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.IntOfUint32((_1245_v194).Cardinality()))).Uint32()).(_dafny.Set), globalState)) || ((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool)), (_index217).Int())
				var _1246_v195 *C2
				_ = _1246_v195
				var _nw199 *C2 = New_C2_()
				_ = _nw199
				_nw199.Ctor__((_1208_v166).Merge(_1208_v166))
				_1246_v195 = _nw199
			}
			var _1247_v196 _dafny.Map
			_ = _1247_v196
			_1247_v196 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_cf24, (_this).F15())
			var _1248_v197 *C3
			_ = _1248_v197
			var _nw200 *C3 = New_C3_()
			_ = _nw200
			_nw200.Ctor__(p1, _dafny.Companion_Sequence_.Concatenate(_1207_cf22, _1207_cf22), _1205_cf24, (func() _dafny.MultiSet {
				if (_1247_v196).Contains((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool)) {
					return (_1247_v196).Get((_1204_cf25).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1204_cf25), 0))).Int()).(bool)).(_dafny.MultiSet)
				}
				return (_this).F15()
			})())
			_1248_v197 = _nw200
		} else if _source14.Is_DC14() {
			var _1249___mcc_h12 bool = _source14.Get_().(D4_DC14).Cf26
			_ = _1249___mcc_h12
			var _1250___mcc_h13 bool = _source14.Get_().(D4_DC14).Cf27
			_ = _1250___mcc_h13
			var _1251_cf27 bool = _1250___mcc_h13
			_ = _1251_cf27
			var _1252_cf26 bool = _1249___mcc_h12
			_ = _1252_cf26
			var _1253_v198 _dafny.MultiSet
			_ = _1253_v198
			_1253_v198 = _dafny.MultiSetOf(Companion_Default___.Fm0(_this.F14(), p1, _dafny.IntOfUint32((_1057_v63).Cardinality()), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-852))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg81 _dafny.Int) interface{} {
					return coer81(arg81)
				}
			}(func(_1254_i21 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality())), globalState), _dafny.IntOfUint32((_1057_v63).Cardinality()), _dafny.IntOfInt64(869))
			var _1255_v199 _dafny.Sequence
			_ = _1255_v199
			_1255_v199 = _dafny.SeqOf(_1253_v198, _dafny.MultiSetOf(p1))
			r1 = Companion_Default___.Fm0(_1252_cf26, _dafny.IntOfUint32((_1057_v63).Cardinality()), p1, ((_1255_v199).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1255_v199).Cardinality()))).Uint32()).(_dafny.MultiSet)).Intersection(_1253_v198), globalState)
			var _1256_v200 _dafny.Sequence
			_ = _1256_v200
			_1256_v200 = _dafny.SeqOf((_this).F15())
			var _1257_v201 _dafny.Sequence
			_ = _1257_v201
			_1257_v201 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(p1, p1)).Cardinality())), p1)
			(_this).M7(_dafny.Companion_Sequence_.Concatenate(_1256_v200, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(770))).Uint32(), func(coer82 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg82 _dafny.Int) interface{} {
					return coer82(arg82)
				}
			}(func(_1258_i22 _dafny.Int) _dafny.MultiSet {
				return (_this).F15()
			}))), (p1).Cmp(_dafny.IntOfUint32((_1257_v201).Cardinality())) >= 0, _dafny.IntOfInt64(-452), _1251_cf27, globalState)
			(_this).F14_set_(_dafny.Companion_Sequence_.IsProperPrefixOf(_1057_v63, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vlgqwvqgo"), _1057_v63)))
			(globalState).F3 = p1
		} else if _source14.Is_DC12() {
			var _1259___mcc_h14 _dafny.Array = _source14.Get_().(D4_DC12).Cf21
			_ = _1259___mcc_h14
			var _1260_cf21 _dafny.Array = _1259___mcc_h14
			_ = _1260_cf21
			var _1261_v202 _dafny.Set
			_ = _1261_v202
			_1261_v202 = _dafny.SetOf(_this.F14())
			var _1262_v203 _dafny.Sequence
			_ = _1262_v203
			_1262_v203 = _dafny.SeqOf((_1261_v202).Cardinality(), p1)
			var _1263_v204 D1
			_ = _1263_v204
			_1263_v204 = Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_1057_v63).Cardinality()))
			var _1264_v205 _dafny.Sequence
			_ = _1264_v205
			_1264_v205 = _1262_v203
			var _1265_v206 _dafny.Map
			_ = _1265_v206
			_1265_v206 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F19())
			var _1266_v207 *C2
			_ = _1266_v207
			var _nw201 *C2 = New_C2_()
			_ = _nw201
			_nw201.Ctor__(_1265_v206)
			_1266_v207 = _nw201
			var _1267_v208 _dafny.Map
			_ = _1267_v208
			_1267_v208 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1266_v207, _1262_v203)
			var _1268_v209 _dafny.Sequence
			_ = _1268_v209
			_1268_v209 = _dafny.SeqOf((_this).F19(), _this.F14())
			var _1269_v210 _dafny.Array
			_ = _1269_v210
			var _nwElement0_40 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(955))).Uint32(), func(coer83 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg83 _dafny.Int) interface{} {
					return coer83(arg83)
				}
			}(func(_1270_i23 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(681)
			}))
			_ = _nwElement0_40
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(28))
			_ = _nw202
			(_nw202).ArraySet1(_nwElement0_40, 0)
			(_nw202).ArraySet1(_1262_v203, 1)
			(_nw202).ArraySet1(_1262_v203, 2)
			(_nw202).ArraySet1(_dafny.SeqOf(p1), 3)
			(_nw202).ArraySet1(Companion_Default___.Fm21(_1263_v204, _this.F14(), Companion_Default___.Fm40(globalState), globalState), 4)
			(_nw202).ArraySet1(_1262_v203, 5)
			(_nw202).ArraySet1(_1262_v203, 6)
			(_nw202).ArraySet1(_1262_v203, 7)
			(_nw202).ArraySet1(_1262_v203, 8)
			(_nw202).ArraySet1(_1262_v203, 9)
			(_nw202).ArraySet1(_1262_v203, 10)
			(_nw202).ArraySet1(_1262_v203, 11)
			(_nw202).ArraySet1(_1262_v203, 12)
			(_nw202).ArraySet1(_1262_v203, 13)
			(_nw202).ArraySet1(_1262_v203, 14)
			(_nw202).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(189)), 15)
			(_nw202).ArraySet1((_1264_v205), 16)
			(_nw202).ArraySet1((func() _dafny.Sequence {
				if (_1267_v208).Contains(_1266_v207) {
					return (_1267_v208).Get(_1266_v207).(_dafny.Sequence)
				}
				return _1262_v203
			})(), 17)
			(_nw202).ArraySet1((func() _dafny.Sequence {
				if (_this).F19() {
					return _1262_v203
				}
				return _1262_v203
			})(), 18)
			(_nw202).ArraySet1(_1262_v203, 19)
			(_nw202).ArraySet1(_1262_v203, 20)
			(_nw202).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1262_v203, _1262_v203), 21)
			(_nw202).ArraySet1(_dafny.SeqOf(p1, p1), 22)
			(_nw202).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-99))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg84 _dafny.Int) interface{} {
					return coer84(arg84)
				}
			}(func(_1271_i24 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqOf(!((_this).F19()), (_this).F19(), _this.F14(), (_this).F19(), (_this).F19())).Cardinality())
			})), 23)
			(_nw202).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1262_v203, _dafny.Companion_Sequence_.Update(_1262_v203, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1268_v209).Cardinality()), _dafny.IntOfUint32((_1262_v203).Cardinality()))).Uint32(), _dafny.IntOfInt64(790))), 24)
			(_nw202).ArraySet1(_1262_v203, 25)
			(_nw202).ArraySet1(_1262_v203, 26)
			(_nw202).ArraySet1(_1262_v203, 27)
			_1269_v210 = _nw202
			_1269_v210 = _1269_v210
			var _1272_v211 _dafny.Map
			_ = _1272_v211
			_1272_v211 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1261_v202)
			var _1273_v212 _dafny.Array
			_ = _1273_v212
			var _nwElement0_41 _dafny.Set = _1261_v202
			_ = _nwElement0_41
			var _nw203 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(9))
			_ = _nw203
			(_nw203).ArraySet1(_nwElement0_41, 0)
			(_nw203).ArraySet1(_1261_v202, 1)
			(_nw203).ArraySet1(Companion_Default___.Fm41(_dafny.IntOfUint32((_1262_v203).Cardinality()), p1, globalState), 2)
			(_nw203).ArraySet1((func() _dafny.Set {
				if (_1272_v211).Contains(_dafny.IntOfInt64(996)) {
					return (_1272_v211).Get(_dafny.IntOfInt64(996)).(_dafny.Set)
				}
				return _1261_v202
			})(), 3)
			(_nw203).ArraySet1((_1261_v202).Intersection(_1261_v202), 4)
			(_nw203).ArraySet1(_1261_v202, 5)
			(_nw203).ArraySet1((_1261_v202).Union(_dafny.SetOf(false, false)), 6)
			(_nw203).ArraySet1(_1261_v202, 7)
			(_nw203).ArraySet1(_1261_v202, 8)
			_1273_v212 = _nw203
			var _1274_v213 _dafny.MultiSet
			_ = _1274_v213
			_1274_v213 = _dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_1057_v63).Cardinality())), p1)
			var _1275_v215 _dafny.Set
			_ = _1275_v215
			_1275_v215 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(518))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg85 _dafny.Int) interface{} {
					return coer85(arg85)
				}
			}((func(_1276_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1277_i25 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(366), (func() _dafny.Map {
						var _coll51 = _dafny.NewMapBuilder()
						_ = _coll51
						for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(210), _dafny.IntOfInt64(972))); ; {
							_compr_51, _ok54 := _iter54()
							if !_ok54 {
								break
							}
							var _1278_v214 _dafny.Int
							_1278_v214 = interface{}(_compr_51).(_dafny.Int)
							if ((_dafny.IntOfInt64(210)).Cmp(_1278_v214) <= 0) && ((_1278_v214).Cmp(_dafny.IntOfInt64(972)) < 0) {
								_coll51.Add((_1278_v214).Times(_1276_p1), true)
							}
						}
						return _coll51.ToMap()
					}()).Cardinality())).Cardinality()
				}
			})(p1))))
			var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1273_v212), 0))
			_ = _index218
			(_1273_v212).ArraySet1(_dafny.SetOf(!(true), Companion_Default___.Fm23(Companion_Default___.Fm0(false, _dafny.IntOfInt64(279), p1, _dafny.MultiSetOf((func() _dafny.Int {
				if (_1274_v213).Contains(p1) {
					return (_1274_v213).Multiplicity(p1)
				}
				return (_dafny.Zero).Minus(p1)
			})(), _dafny.IntOfUint32((_1268_v209).Cardinality()), p1), globalState), _this.F14(), Companion_Default___.Fm24(p1, (_this).F19(), false, p1, globalState), _1275_v215, globalState), (_this).F19(), (_this).F19(), (_this).F19()), (_index218).Int())
			var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1273_v212), 0))
			_ = _index219
			var _rhs190 _dafny.Int = p1
			_ = _rhs190
			var _rhs191 bool = !(!(_dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1057_v63, _dafny.UnicodeSeqOfUtf8Bytes("nhxiqa")), _1057_v63)))
			_ = _rhs191
			var _rhs192 _dafny.Set = (func() _dafny.Set {
				if _this.F14() {
					return _dafny.SetOf(!(_this.F14()), _this.F14(), _this.F14(), (_this).F19(), (_this).F19())
				}
				return (_1261_v202).Difference(_1261_v202)
			})()
			_ = _rhs192
			var _lhs159 *GlobalState = globalState
			_ = _lhs159
			var _lhs160 _dafny.Array = _1273_v212
			_ = _lhs160
			var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(632), _dafny.ArrayLen((_1273_v212), 0))
			_ = _lhs161
			_lhs159.F3 = _rhs190
			r2 = _rhs191
			(_lhs160).ArraySet1(_rhs192, (_lhs161).Int())
			var _1279_v216 _dafny.Array
			_ = _1279_v216
			var _nw204 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw204
			_1279_v216 = _nw204
			var _1280_v217 *C3
			_ = _1280_v217
			var _nw205 *C3 = New_C3_()
			_ = _nw205
			_nw205.Ctor__(_dafny.IntOfInt64(6), _1057_v63, (_this).F19(), (_this).F15())
			_1280_v217 = _nw205
			var _1281_v218 _dafny.Map
			_ = _1281_v218
			_1281_v218 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1279_v216, _1280_v217)
			_1281_v218 = (_1281_v218).Update((func() _dafny.Array {
				if (_this).F19() {
					return _1279_v216
				}
				return _1279_v216
			})(), _1280_v217)
			var _1282_v219 _dafny.Map
			_ = _1282_v219
			_1282_v219 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_1280_v217).F20()), (_this).F19())
			var _1283_v220 _dafny.Array
			_ = _1283_v220
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_32
			var _nw206 _dafny.Array
			_ = _nw206
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw206 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) _dafny.Int = (func(_1284_p1 _dafny.Int, _1285_v217 *C3) func(_dafny.Int) _dafny.Int {
					return func(_1286_i26 _dafny.Int) _dafny.Int {
						return (_1286_i26).Minus((Companion_D1_.Create_DC4_((_dafny.Zero).Minus(_1284_p1), (_dafny.Zero).Minus((_dafny.Zero).Minus((_1285_v217).F20())))).Dtor_cf4())
					}
				})(p1, _1280_v217)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw206 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw206).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw206).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1283_v220 = _nw206
			var _1287_v221 _dafny.Map
			_ = _1287_v221
			_1287_v221 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_v219, _1283_v220)
			_1287_v221 = (_1287_v221).Update(_1282_v219, _1283_v220)
		} else {
			var _1288___mcc_h15 D4 = _source14.Get_().(D4_DC15).Cf28
			_ = _1288___mcc_h15
			var _1289_cf28 D4 = _1288___mcc_h15
			_ = _1289_cf28
			if (p1).Cmp(_dafny.IntOfInt64(595)) > 0 {
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))
				_ = _index220
				(_1194_v163).ArraySet1(!(_this.F14()) || ((_this).F19()), (_index220).Int())
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))
				_ = _index221
				(_1194_v163).ArraySet1((_this).F19(), (_index221).Int())
				var _1290_v222 _dafny.CodePoint
				_ = _1290_v222
				_1290_v222 = _dafny.CodePoint('r')
				var _1291_v223 _dafny.Map
				_ = _1291_v223
				_1291_v223 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1290_v222, (_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool))
				_1291_v223 = (_1291_v223).Update(_1290_v222, (_this).F19())
				var _1292_v224 _dafny.Map
				_ = _1292_v224
				_1292_v224 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!((_this).F19())), _dafny.IntOfUint32((_1195_v164).Cardinality()))
				var _1293_v226 _dafny.Sequence
				_ = _1293_v226
				_1293_v226 = _dafny.SeqOf((_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool))
				var _1294_v227 _dafny.Map
				_ = _1294_v227
				_1294_v227 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1057_v63)
				var _1295_v228 D8
				_ = _1295_v228
				_1295_v228 = Companion_D8_.Create_DC26_((_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool), _1294_v227)
				var _1296_v229 _dafny.MultiSet
				_ = _1296_v229
				_1296_v229 = _dafny.MultiSetOf(Companion_Default___.Fm42(false, !((_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool)), _dafny.IntOfInt64(265), (_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool), globalState), _1295_v228)
				var _1297_v230 D3
				_ = _1297_v230
				_1297_v230 = Companion_D3_.Create_DC11_(p1, p1, _this.F14(), _dafny.IntOfInt64(-278))
				var _1298_v231 _dafny.MultiSet
				_ = _1298_v231
				_1298_v231 = _dafny.MultiSetOf((_1297_v230).Dtor_cf18(), _dafny.IntOfUint32((_1057_v63).Cardinality()))
				var _1299_v232 _dafny.Set
				_ = _1299_v232
				_1299_v232 = _dafny.SetOf(_1298_v231, _1298_v231)
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))
				_ = _index222
				var _rhs193 _dafny.Map = func() _dafny.Map {
					var _coll52 = _dafny.NewMapBuilder()
					_ = _coll52
					for _iter55 := _dafny.Iterate((_dafny.SeqOf(_1293_v226)).Elements()); ; {
						_compr_52, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1300_v225 _dafny.Sequence
						_1300_v225 = interface{}(_compr_52).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_1293_v226), _1300_v225) {
							_coll52.Add(_1300_v225, p1)
						}
					}
					return _coll52.ToMap()
				}()
				_ = _rhs193
				var _rhs194 bool = !((_1296_v229).IsProperSubsetOf(_1296_v229))
				_ = _rhs194
				var _rhs195 bool = (_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool)
				_ = _rhs195
				var _rhs196 bool = (_1299_v232).IsDisjointFrom(_1299_v232)
				_ = _rhs196
				var _lhs162 *GlobalState = globalState
				_ = _lhs162
				var _lhs163 *C4 = _this
				_ = _lhs163
				var _lhs164 _dafny.Array = _1194_v163
				_ = _lhs164
				var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))
				_ = _lhs165
				_1292_v224 = _rhs193
				_lhs162.F9 = _rhs194
				_lhs163.F14_set_(_rhs195)
				(_lhs164).ArraySet1(_rhs196, (_lhs165).Int())
				var _1301_v233 _dafny.Array
				_ = _1301_v233
				var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw207
				_1301_v233 = _nw207
				var _1302_v234 _dafny.Array
				_ = _1302_v234
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_33
				var _nw208 _dafny.Array
				_ = _nw208
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw208 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) _dafny.Int = (func(_1303_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1304_i27 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1304_i27, _1303_p1)
						}
					})(p1)
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw208 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw208).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw208).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1302_v234 = _nw208
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1301_v233), 0))
				_ = _index223
				(_1301_v233).ArraySet1(_1302_v234, (_index223).Int())
				var _1305_v235 _dafny.Map
				_ = _1305_v235
				_1305_v235 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1290_v222)
				var _1306_v236 _dafny.Map
				_ = _1306_v236
				_1306_v236 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1290_v222)
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1301_v233), 0))
				_ = _index224
				var _rhs197 _dafny.Int = p1
				_ = _rhs197
				var _rhs198 _dafny.Int = p1
				_ = _rhs198
				var _rhs199 _dafny.Array = _1302_v234
				_ = _rhs199
				var _rhs200 _dafny.Sequence = _dafny.Companion_Sequence_.Update((_this).Fm3((_dafny.Zero).Minus(p1), p1, !((func() bool {
					if (_1291_v223).Contains(_1290_v222) {
						return (_1291_v223).Get(_1290_v222).(bool)
					}
					return _this.F14()
				})()), Companion_Default___.SafeModuloInt((_1305_v235).Cardinality(), p1), globalState), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_this).Fm3((_dafny.Zero).Minus(p1), p1, !((func() bool {
					if (_1291_v223).Contains(_1290_v222) {
						return (_1291_v223).Get(_1290_v222).(bool)
					}
					return _this.F14()
				})()), Companion_Default___.SafeModuloInt((_1305_v235).Cardinality(), p1), globalState)).Cardinality()))).Uint32(), (func() _dafny.CodePoint {
					if (_1306_v236).Contains(!(_this.F14())) {
						return (_1306_v236).Get(!(_this.F14())).(_dafny.CodePoint)
					}
					return _1290_v222
				})())
				_ = _rhs200
				var _lhs166 *GlobalState = globalState
				_ = _lhs166
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				var _lhs168 _dafny.Array = _1301_v233
				_ = _lhs168
				var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1301_v233), 0))
				_ = _lhs169
				_lhs166.F3 = _rhs197
				_lhs167.F3 = _rhs198
				(_lhs168).ArraySet1(_rhs199, (_lhs169).Int())
				_1057_v63 = _rhs200
				var _1307_v237 _dafny.Set
				_ = _1307_v237
				_1307_v237 = _dafny.SetOf((_dafny.SetOf((_this).F19(), (_this).F19(), !((_1194_v163).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(853), _dafny.ArrayLen((_1194_v163), 0))).Int()).(bool)))).Cardinality(), (_1297_v230).Dtor_cf18(), p1)
				_1307_v237 = _1307_v237
			} else {
				var _1308_v238 _dafny.Array
				_ = _1308_v238
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_34
				var _nw209 _dafny.Array
				_ = _nw209
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw209 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) _dafny.Sequence = (func(_1309_p1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1310_i28 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_1309_p1)
						}
					})(p1)
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw209 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw209).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw209).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1308_v238 = _nw209
				var _1311_v239 _dafny.Map
				_ = _1311_v239
				_1311_v239 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), p1)
				var _1312_v240 _dafny.Map
				_ = _1312_v240
				_1312_v240 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1311_v239)
				var _1313_v241 _dafny.Map
				_ = _1313_v241
				_1313_v241 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), (_this).Fm2((_this).F19(), _this.F14(), globalState))
				var _1314_v242 _dafny.Sequence
				_ = _1314_v242
				_1314_v242 = _dafny.SeqOf(p1, ((func() _dafny.Map {
					if (_1312_v240).Contains((_1313_v241).Cardinality()) {
						return (_1312_v240).Get((_1313_v241).Cardinality()).(_dafny.Map)
					}
					return _1311_v239
				})()).Cardinality(), p1, p1, p1)
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1308_v238), 0))
				_ = _index225
				(_1308_v238).ArraySet1((func() _dafny.Sequence {
					if _this.F14() {
						return _1314_v242
					}
					return _1314_v242
				})(), (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_1308_v238), 0))
				_ = _index226
				(_1308_v238).ArraySet1(_1314_v242, (_index226).Int())
				var _1315_v243 _dafny.Sequence
				_ = _1315_v243
				_1315_v243 = _dafny.SeqOf(_1057_v63)
				r1 = (func() _dafny.Int {
					if (_this).F19() {
						return Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfUint32((_1057_v63).Cardinality()))
					}
					return _dafny.IntOfUint32((_1315_v243).Cardinality())
				})()
				var _1316_v244 D4
				_ = _1316_v244
				_1316_v244 = Companion_D4_.Create_DC12_(_1194_v163)
				var _1317_v245 _dafny.Array
				_ = _1317_v245
				var _nwElement0_42 D4 = _1316_v244
				_ = _nwElement0_42
				var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(4))
				_ = _nw210
				(_nw210).ArraySet1(_nwElement0_42, 0)
				(_nw210).ArraySet1(_1316_v244, 1)
				(_nw210).ArraySet1(_1316_v244, 2)
				(_nw210).ArraySet1(_1316_v244, 3)
				_1317_v245 = _nw210
				var _1318_v246 _dafny.Map
				_ = _1318_v246
				_1318_v246 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1317_v245, true)
				(globalState).F9 = !(_1318_v246).Contains(_1317_v245)
				var _1319_v247 _dafny.Map
				_ = _1319_v247
				_1319_v247 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1057_v63).Cardinality()), _this.F14())
				(_this).F14_set_((((_1319_v247).Cardinality()).Cmp(p1) == 0) == ((_this).F19()))
				var _1320_v248 _dafny.Sequence
				_ = _1320_v248
				_1320_v248 = _dafny.SeqOf((_this).F19())
				(_this).F14_set_(!((_1320_v248).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.IntOfUint32((_1320_v248).Cardinality()))).Uint32()).(bool)) || ((_this).F19()))
			}
			if !(true) {
				var _1321_v249 _dafny.Sequence
				_ = _1321_v249
				_1321_v249 = _dafny.SeqOf(p1, p1)
				var _1322_v250 _dafny.MultiSet
				_ = _1322_v250
				_1322_v250 = _dafny.MultiSetOf(p1)
				var _1323_v251 *C3
				_ = _1323_v251
				var _nw211 *C3 = New_C3_()
				_ = _nw211
				_nw211.Ctor__(Companion_Default___.Fm0((_this).F19(), _dafny.IntOfUint32((_1321_v249).Cardinality()), p1, _1322_v250, globalState), _1057_v63, (p1).Cmp(p1) >= 0, (_this).F15())
				_1323_v251 = _nw211
				var _1324_v252 _dafny.Sequence
				_ = _1324_v252
				_1324_v252 = _dafny.SeqOf(_1322_v250, _1322_v250, _1322_v250, _1322_v250)
				(_this).M7(Companion_Default___.Fm43((_1323_v251).F20(), p1, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.Fm0(!((_this).F19()), (_1323_v251).F20(), p1, _1322_v250, globalState)))).Cardinality()), globalState), _this.F14(), ((_1322_v250).Intersection((_1324_v252).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1324_v252).Cardinality()))).Uint32()).(_dafny.MultiSet))).Cardinality(), _this.F14(), globalState)
				var _1325_v253 D2
				_ = _1325_v253
				_1325_v253 = Companion_D2_.Create_DC8_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-565))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg86 _dafny.Int) interface{} {
						return coer86(arg86)
					}
				}(func(_1326_i29 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}))).Cardinality()), Companion_Default___.SafeDivisionInt((_1323_v251).F20(), p1), p1)
				var _1327_v254 _dafny.Array
				_ = _1327_v254
				var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
				_ = _nw212
				_1327_v254 = _nw212
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1327_v254), 0))
				_ = _index227
				(_1327_v254).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1323_v251).F20(), p1)), (_index227).Int())
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1327_v254), 0))
				_ = _index228
				(_1327_v254).ArraySet1((_dafny.IntOfUint32((_1321_v249).Cardinality())).Minus(p1), (_index228).Int())
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1194_v163), 0))
				_ = _index229
				(_1194_v163).ArraySet1(true, (_index229).Int())
				var _1328_v255 _dafny.Set
				_ = _1328_v255
				_1328_v255 = _dafny.SetOf(_1327_v254)
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1327_v254), 0))
				_ = _index230
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1327_v254), 0))
				_ = _index231
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1194_v163), 0))
				_ = _index232
				var _rhs201 bool = (_this).F19()
				_ = _rhs201
				var _rhs202 D2 = Companion_Default___.Fm44((_this).F19(), (_1323_v251).F20(), (_1323_v251).F20(), globalState)
				_ = _rhs202
				var _rhs203 _dafny.Int = (func() _dafny.Int {
					if _this.F14() {
						return p1
					}
					return p1
				})()
				_ = _rhs203
				var _rhs204 _dafny.Int = (_1323_v251).F20()
				_ = _rhs204
				var _rhs205 bool = ((_1328_v255).Union(_1328_v255)).Equals(_dafny.SetOf(_1327_v254))
				_ = _rhs205
				var _lhs170 _dafny.Array = _1327_v254
				_ = _lhs170
				var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1327_v254), 0))
				_ = _lhs171
				var _lhs172 _dafny.Array = _1327_v254
				_ = _lhs172
				var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(605), _dafny.ArrayLen((_1327_v254), 0))
				_ = _lhs173
				var _lhs174 _dafny.Array = _1194_v163
				_ = _lhs174
				var _lhs175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_1194_v163), 0))
				_ = _lhs175
				r2 = _rhs201
				_1325_v253 = _rhs202
				(_lhs170).ArraySet1(_rhs203, (_lhs171).Int())
				(_lhs172).ArraySet1(_rhs204, (_lhs173).Int())
				(_lhs174).ArraySet1(_rhs205, (_lhs175).Int())
				var _1329_v256 _dafny.Map
				_ = _1329_v256
				_1329_v256 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(560))
				var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(298), _dafny.ArrayLen((_1327_v254), 0))
				_ = _index233
				(_1327_v254).ArraySet1(((_1323_v251).F20()).Times((func() _dafny.Int {
					if (_1329_v256).Contains(_this.F14()) {
						return (_1329_v256).Get(_this.F14()).(_dafny.Int)
					}
					return (_1323_v251).F20()
				})()), (_index233).Int())
				(globalState).F3 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("owt")).Cardinality())
			} else {
				var _1330_v257 _dafny.Array
				_ = _1330_v257
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw213
				_1330_v257 = _nw213
				_1330_v257 = _1330_v257
				var _1331_v258 D8
				_ = _1331_v258
				_1331_v258 = Companion_D8_.Create_DC26_((_this).F19(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1057_v63))
				var _1332_v259 _dafny.Array
				_ = _1332_v259
				var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw214
				_1332_v259 = _nw214
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))
				_ = _index234
				(_1332_v259).ArraySet1((func() _dafny.Int {
					if (_this).F19() {
						return p1
					}
					return p1
				})(), (_index234).Int())
				var _1333_v260 D4
				_ = _1333_v260
				_1333_v260 = Companion_D4_.Create_DC14_(_this.F14(), (_this).Fm2(_this.F14(), !(_this.F14()), globalState))
				var _1334_v261 _dafny.Map
				_ = _1334_v261
				_1334_v261 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(p1, p1, (_1333_v260).Dtor_cf26(), p1, globalState), _1332_v259)
				var _1335_v262 _dafny.Sequence
				_ = _1335_v262
				_1335_v262 = _dafny.SeqOf(_1057_v63, _1057_v63, _1057_v63)
				var _1336_v263 _dafny.Map
				_ = _1336_v263
				_1336_v263 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1335_v262).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1335_v262).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _pat_let_tv37 = _1336_v263
				_ = _pat_let_tv37
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))
				_ = _index235
				var _rhs206 _dafny.Int = p1
				_ = _rhs206
				var _rhs207 _dafny.Int = (_1334_v261).Cardinality()
				_ = _rhs207
				var _rhs208 D8 = func(_pat_let48_0 D8) D8 {
					return func(_1337_dt__update__tmp_h9 D8) D8 {
						return func(_pat_let49_0 _dafny.Map) D8 {
							return func(_1338_dt__update_hcf47_h0 _dafny.Map) D8 {
								return Companion_D8_.Create_DC26_((_1337_dt__update__tmp_h9).Dtor_cf46(), _1338_dt__update_hcf47_h0)
							}(_pat_let49_0)
						}(_pat_let_tv37)
					}(_pat_let48_0)
				}(_1331_v258)
				_ = _rhs208
				var _rhs209 _dafny.Int = Companion_Default___.SafeModuloInt((p1).Times(((_this).F15()).Cardinality()), p1)
				_ = _rhs209
				var _lhs176 *GlobalState = globalState
				_ = _lhs176
				var _lhs177 _dafny.Array = _1332_v259
				_ = _lhs177
				var _lhs178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))
				_ = _lhs178
				r0 = _rhs206
				_lhs176.F3 = _rhs207
				_1331_v258 = _rhs208
				(_lhs177).ArraySet1(_rhs209, (_lhs178).Int())
				var _1339_v264 D2
				_ = _1339_v264
				_1339_v264 = Companion_D2_.Create_DC9_((_this).F19(), (_1332_v259).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))).Int()).(_dafny.Int), _dafny.CodePoint('n'))
				r0 = (_1339_v264).Dtor_cf14()
				r2 = !(_this.F14())
				var _1340_v265 _dafny.Map
				_ = _1340_v265
				_1340_v265 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1332_v259).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))).Int()).(_dafny.Int), (p0).Cardinality())
				var _1341_v266 _dafny.Map
				_ = _1341_v266
				_1341_v266 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1340_v265, (_1332_v259).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))).Int()).(_dafny.Int))
				_1341_v266 = (_1341_v266).Update((_1340_v265).Merge(_1340_v265), (_1332_v259).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1332_v259), 0))).Int()).(_dafny.Int))
			}
			(_this).F14_set_(!(_this.F14()))
			if (_this).F19() {
				var _1342_v267 _dafny.Map
				_ = _1342_v267
				_1342_v267 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F14())
				var _1343_v268 *C2
				_ = _1343_v268
				var _nw215 *C2 = New_C2_()
				_ = _nw215
				_nw215.Ctor__((_1342_v267).Update(p1, (_this).F19()))
				_1343_v268 = _nw215
				var _1344_v269 _dafny.Array
				_ = _1344_v269
				var _nw216 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw216
				_1344_v269 = _nw216
				var _1345_v270 _dafny.CodePoint
				_ = _1345_v270
				_1345_v270 = _dafny.CodePoint('a')
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_1344_v269), 0))
				_ = _index236
				(_1344_v269).ArraySet1CodePoint(_1345_v270, (_index236).Int())
				var _1346_v271 _dafny.Set
				_ = _1346_v271
				_1346_v271 = _dafny.SetOf((_this).F15(), (_this).F15())
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_1344_v269), 0))
				_ = _index237
				var _rhs210 bool = (_dafny.IntOfInt64(241)).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(p1))) < 0
				_ = _rhs210
				var _rhs211 bool = (_1346_v271).IsSubsetOf(_1346_v271)
				_ = _rhs211
				var _rhs212 _dafny.CodePoint = (_1057_v63).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1057_v63).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs212
				var _lhs179 *GlobalState = globalState
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 _dafny.Array = _1344_v269
				_ = _lhs181
				var _lhs182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(236), _dafny.ArrayLen((_1344_v269), 0))
				_ = _lhs182
				_lhs179.F9 = _rhs210
				_lhs180.F9 = _rhs211
				(_lhs181).ArraySet1CodePoint(_rhs212, (_lhs182).Int())
				var _1347_v272 _dafny.MultiSet
				_ = _1347_v272
				_1347_v272 = _dafny.MultiSetOf(p1, p1)
				var _1348_v273 _dafny.Sequence
				_ = _1348_v273
				_1348_v273 = _dafny.SeqOf((_1347_v272).Cardinality())
				var _1349_v274 _dafny.Set
				_ = _1349_v274
				_1349_v274 = _dafny.SetOf(_1348_v273, _dafny.SeqOf(p1), _1348_v273)
				var _1350_v275 _dafny.Map
				_ = _1350_v275
				_1350_v275 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F19())
				var _pat_let_tv38 = _1350_v275
				_ = _pat_let_tv38
				var _1351_v276 _dafny.Map
				_ = _1351_v276
				_1351_v276 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wwmccp")).Cardinality())), (func() _dafny.Int {
					if _this.F14() {
						return (_dafny.SetOf(p1, p1)).Cardinality()
					}
					return ((func(_pat_let50_0 D3) D3 {
						return func(_1352_dt__update__tmp_h10 D3) D3 {
							return func(_pat_let51_0 _dafny.Map) D3 {
								return func(_1353_dt__update_hcf16_h0 _dafny.Map) D3 {
									return Companion_D3_.Create_DC10_(_1353_dt__update_hcf16_h0)
								}(_pat_let51_0)
							}(_pat_let_tv38)
						}(_pat_let50_0)
					}(Companion_Default___.Fm45(p1, Companion_Default___.Fm23(p1, true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1345_v270, p1), _1349_v274, globalState), p1, globalState))).Dtor_cf16()).Cardinality()
				})())
				_1351_v276 = (_1351_v276).Update(p1, _dafny.IntOfUint32((_1057_v63).Cardinality()))
				(globalState).F3 = p1
				(globalState).F9 = (_1347_v272).IsDisjointFrom((_1347_v272).Intersection(_1347_v272))
			} else {
				(globalState).F3 = p1
				var _1354_v277 _dafny.Map
				_ = _1354_v277
				_1354_v277 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (p1).Times(p1))
				_1354_v277 = _1354_v277
				r1 = p1
				_1057_v63 = _dafny.Companion_Sequence_.Concatenate(_1057_v63, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-916))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg87 _dafny.Int) interface{} {
						return coer87(arg87)
					}
				}(func(_1355_i30 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				})))
				var _1356_v278 _dafny.MultiSet
				_ = _1356_v278
				_1356_v278 = _dafny.MultiSetOf(p1, p1)
				(globalState).F3 = Companion_Default___.Fm0(!((_this).F19()) || (_this.F14()), p1, (func() _dafny.Int {
					if (_this).F19() {
						return p1
					}
					return p1
				})(), _1356_v278, globalState)
			}
		}
		r0 = p1
		r1 = Companion_Default___.SafeDivisionInt(p1, p1)
		r2 = _this.F14()
		return r0, r1, r2
	}
}
func (_this *C4) M6(globalState *GlobalState) {
	{
		if _this.F14() {
			var _1357_v0 _dafny.Array
			_ = _1357_v0
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_35
			var _nw217 _dafny.Array
			_ = _nw217
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw217 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = func(_1358_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1358_i0, _dafny.IntOfUint32((_dafny.SeqOf(_this.F14())).Cardinality()))
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw217 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw217).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw217).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1357_v0 = _nw217
			var _1359_v1 _dafny.Int
			_ = _1359_v1
			_1359_v1 = _dafny.IntOfInt64(835)
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1357_v0), 0))
			_ = _index238
			(_1357_v0).ArraySet1((_1359_v1).Times(_dafny.IntOfInt64(698)), (_index238).Int())
			var _1360_v2 _dafny.Sequence
			_ = _1360_v2
			_1360_v2 = _dafny.UnicodeSeqOfUtf8Bytes("dchmesu")
			var _1361_v3 _dafny.Sequence
			_ = _1361_v3
			_1361_v3 = _dafny.SeqOf(_1359_v1)
			var _1362_v4 _dafny.Sequence
			_ = _1362_v4
			_1362_v4 = _dafny.SeqOf(_this.F14())
			var _1363_v5 _dafny.Set
			_ = _1363_v5
			_1363_v5 = _dafny.SetOf(_1359_v1)
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1357_v0), 0))
			_ = _index239
			var _rhs213 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_1361_v3).Select((Companion_Default___.SafeIndex(_1359_v1, _dafny.IntOfUint32((_1361_v3).Cardinality()))).Uint32()).(_dafny.Int)))
			_ = _rhs213
			var _rhs214 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1360_v2, _1360_v2)
			_ = _rhs214
			var _rhs215 bool = (_1363_v5).IsProperSubsetOf(_dafny.SetOf(_1359_v1, (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_1359_v1))), _dafny.IntOfUint32((_1362_v4).Cardinality())))
			_ = _rhs215
			var _rhs216 _dafny.Int = _dafny.IntOfInt64(711)
			_ = _rhs216
			var _rhs217 _dafny.Int = Companion_Default___.SafeModuloInt(_1359_v1, _1359_v1)
			_ = _rhs217
			var _lhs183 _dafny.Array = _1357_v0
			_ = _lhs183
			var _lhs184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1357_v0), 0))
			_ = _lhs184
			var _lhs185 *GlobalState = globalState
			_ = _lhs185
			var _lhs186 *GlobalState = globalState
			_ = _lhs186
			var _lhs187 *GlobalState = globalState
			_ = _lhs187
			(_lhs183).ArraySet1(_rhs213, (_lhs184).Int())
			_1360_v2 = _rhs214
			_lhs185.F9 = _rhs215
			_lhs186.F3 = _rhs216
			_lhs187.F3 = _rhs217
			var _1364_v6 _dafny.Array
			_ = _1364_v6
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_36
			var _nw218 _dafny.Array
			_ = _nw218
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw218 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) bool = func(_1365_i1 _dafny.Int) bool {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14())).Contains((_this).F19())
				}
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw218 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw218).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw218).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1364_v6 = _nw218
			(globalState).F2 = _1364_v6
			var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1364_v6), 0))
			_ = _index240
			(_1364_v6).ArraySet1(_this.F14(), (_index240).Int())
			var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1364_v6), 0))
			_ = _index241
			(_1364_v6).ArraySet1(_this.F14(), (_index241).Int())
			var _1366_v7 _dafny.Map
			_ = _1366_v7
			_1366_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1357_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1357_v0), 0))).Int()).(_dafny.Int), false)
			_1363_v5 = (Companion_Default___.Fm22(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("epm")).Cardinality()), _dafny.SeqOf(_1366_v7), (_this).F19(), (_this).F19(), globalState)).Intersection((_1363_v5).Intersection(_1363_v5))
			var _1367_v8 _dafny.CodePoint
			_ = _1367_v8
			_1367_v8 = _dafny.CodePoint('h')
			var _rhs218 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1360_v2, (Companion_Default___.SafeIndex(_1359_v1, _dafny.IntOfUint32((_1360_v2).Cardinality()))).Uint32(), _1367_v8), _dafny.UnicodeSeqOfUtf8Bytes("ob"))).Cardinality()))
			_ = _rhs218
			var _rhs219 _dafny.Int = (func() _dafny.Int {
				if (_1364_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1364_v6), 0))).Int()).(bool) {
					return (_1357_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1357_v0), 0))).Int()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.SeqOf((_1364_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.ArrayLen((_1364_v6), 0))).Int()).(bool), (_this).F19(), _this.F14())).Cardinality())
			})()
			_ = _rhs219
			var _rhs220 _dafny.Array = _1357_v0
			_ = _rhs220
			var _rhs221 _dafny.Sequence = _1360_v2
			_ = _rhs221
			var _lhs188 *GlobalState = globalState
			_ = _lhs188
			_1359_v1 = _rhs218
			_lhs188.F3 = _rhs219
			_1357_v0 = _rhs220
			_1360_v2 = _rhs221
		} else {
			var _1368_v9 _dafny.Int
			_ = _1368_v9
			_1368_v9 = _dafny.IntOfInt64(819)
			(globalState).F3 = _1368_v9
			var _1369_v10 _dafny.Array
			_ = _1369_v10
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_37
			var _nw219 _dafny.Array
			_ = _nw219
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw219 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) bool = func(_1370_i2 _dafny.Int) bool {
					return (_this).F19()
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw219 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw219).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw219).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1369_v10 = _nw219
			var _1371_v11 _dafny.Sequence
			_ = _1371_v11
			_1371_v11 = _dafny.SeqOf(_1369_v10, _1369_v10, _1369_v10)
			var _1372_v12 _dafny.MultiSet
			_ = _1372_v12
			_1372_v12 = _dafny.MultiSetOf((_dafny.Zero).Minus(_1368_v9))
			(globalState).F2 = (_1371_v11).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm0(_this.F14(), (_dafny.Zero).Minus(_1368_v9), _1368_v9, _1372_v12, globalState), _dafny.IntOfUint32((_1371_v11).Cardinality()))).Uint32()).(_dafny.Array)
			var _1373_v13 _dafny.Array
			_ = _1373_v13
			var _nw220 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw220
			_1373_v13 = _nw220
			var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw221
			_1373_v13 = _nw221
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_1369_v10), 0))
			_ = _index242
			(_1369_v10).ArraySet1(_this.F14(), (_index242).Int())
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_1369_v10), 0))
			_ = _index243
			(_1369_v10).ArraySet1(((_this).F15()).IsProperSubsetOf((_this).F15()), (_index243).Int())
			(_this).F14_set_((_1369_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(630), _dafny.ArrayLen((_1369_v10), 0))).Int()).(bool))
		}
		var _1374_v14 _dafny.Int
		_ = _1374_v14
		_1374_v14 = _dafny.IntOfInt64(202)
		var _hi3 _dafny.Int = _1374_v14
		_ = _hi3
		for _1375_i3 := _1374_v14; _1375_i3.Cmp(_hi3) < 0; _1375_i3 = _1375_i3.Plus(_dafny.One) {
			var _1376_v15 _dafny.Sequence
			_ = _1376_v15
			_1376_v15 = _dafny.UnicodeSeqOfUtf8Bytes("gmbu")
			var _1377_v16 _dafny.MultiSet
			_ = _1377_v16
			_1377_v16 = _dafny.MultiSetOf((_this).F19())
			var _1378_v17 D8
			_ = _1378_v17
			_1378_v17 = Companion_D8_.Create_DC25_()
			var _rhs222 _dafny.Sequence = _1376_v15
			_ = _rhs222
			var _rhs223 _dafny.MultiSet = _1377_v16
			_ = _rhs223
			var _rhs224 bool = _this.F14()
			_ = _rhs224
			var _rhs225 D8 = _1378_v17
			_ = _rhs225
			var _lhs189 *GlobalState = globalState
			_ = _lhs189
			_1376_v15 = _rhs222
			_1377_v16 = _rhs223
			_lhs189.F9 = _rhs224
			_1378_v17 = _rhs225
			var _1379_v18 _dafny.Array
			_ = _1379_v18
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_38
			var _nw222 _dafny.Array
			_ = _nw222
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw222 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) bool = (func(_1380_i3 _dafny.Int, _1381_v14 _dafny.Int) func(_dafny.Int) bool {
					return func(_1382_i4 _dafny.Int) bool {
						return !((_1380_i3).Cmp(_1381_v14) != 0)
					}
				})(_1375_i3, _1374_v14)
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw222 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw222).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw222).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1379_v18 = _nw222
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1379_v18), 0))
			_ = _index244
			(_1379_v18).ArraySet1((_this).F19(), (_index244).Int())
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1379_v18), 0))
			_ = _index245
			(_1379_v18).ArraySet1(_this.F14(), (_index245).Int())
			var _1383_v19 _dafny.Sequence
			_ = _1383_v19
			_1383_v19 = _dafny.SeqOf((_1379_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1379_v18), 0))).Int()).(bool))
			var _1384_v20 _dafny.MultiSet
			_ = _1384_v20
			_1384_v20 = _dafny.MultiSetOf(Companion_Default___.Fm38(globalState))
			_1374_v14 = (_this).Fm17(_1375_i3, !((_1379_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(997), _dafny.ArrayLen((_1379_v18), 0))).Int()).(bool)), (_dafny.MultiSetOf(_1383_v19)).Intersection(_1384_v20), _dafny.IntOfInt64(-899), globalState)
			(globalState).F9 = (_1383_v19).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_1375_i3, _1375_i3), _dafny.IntOfUint32((_1383_v19).Cardinality()))).Uint32()).(bool)
		}
		if (_this).F19() {
			var _1385_v21 _dafny.Array
			_ = _1385_v21
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_39
			var _nw223 _dafny.Array
			_ = _nw223
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw223 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) _dafny.Set = func(_1386_i5 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_this.F14(), _this.F14())
				}
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw223 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw223).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw223).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1385_v21 = _nw223
			var _1387_v22 _dafny.Set
			_ = _1387_v22
			_1387_v22 = _dafny.SetOf(_this.F14())
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_1385_v21), 0))
			_ = _index246
			(_1385_v21).ArraySet1((_1387_v22).Difference(_1387_v22), (_index246).Int())
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(579), _dafny.ArrayLen((_1385_v21), 0))
			_ = _index247
			(_1385_v21).ArraySet1(_1387_v22, (_index247).Int())
			if _this.F14() {
				var _1388_v23 _dafny.Array
				_ = _1388_v23
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(2)
				_ = _len0_40
				var _nw224 _dafny.Array
				_ = _nw224
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw224 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Sequence = (func(_1389_v14 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1390_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf((_dafny.Zero).Minus(_1389_v14), _1389_v14, _1389_v14)
						}
					})(_1374_v14)
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw224 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw224).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw224).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1388_v23 = _nw224
				var _1391_v24 _dafny.Array
				_ = _1391_v24
				var _nw225 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(24))
				_ = _nw225
				_1391_v24 = _nw225
				var _1392_v25 D0
				_ = _1392_v25
				_1392_v25 = Companion_D0_.Create_DC1_((_this).Fm2(_this.F14(), true, globalState))
				var _1393_v26 D0
				_ = _1393_v26
				_1393_v26 = Companion_D0_.Create_DC2_(_1392_v25)
				var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1391_v24), 0))
				_ = _index248
				(_1391_v24).ArraySet1(_1393_v26, (_index248).Int())
				var _1394_v27 _dafny.Array
				_ = _1394_v27
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_41
				var _nw226 _dafny.Array
				_ = _nw226
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw226 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) D0 = func(_1395_i7 _dafny.Int) D0 {
						return Companion_D0_.Create_DC1_(_this.F14())
					}
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
				_1394_v27 = _nw226
				var _1396_v28 _dafny.Map
				_ = _1396_v28
				_1396_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v14, _dafny.SetOf(_this.F14(), (_this).F19()))
				var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1394_v27), 0))
				_ = _index249
				(_1394_v27).ArraySet1(Companion_Default___.Fm37(_dafny.SetOf(_1374_v14), _1396_v28, _this.F14(), (_this).F19(), globalState), (_index249).Int())
				var _1397_v29 _dafny.Array
				_ = _1397_v29
				var _nw227 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw227
				_1397_v29 = _nw227
				var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1397_v29), 0))
				_ = _index250
				(_1397_v29).ArraySet1(!((_this).F19()), (_index250).Int())
				var _1398_v30 _dafny.Set
				_ = _1398_v30
				_1398_v30 = _dafny.SetOf(_1374_v14)
				var _1399_v31 _dafny.Sequence
				_ = _1399_v31
				_1399_v31 = _dafny.UnicodeSeqOfUtf8Bytes("pw")
				var _1400_v32 _dafny.Map
				_ = _1400_v32
				_1400_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(343))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg88 _dafny.Int) interface{} {
						return coer88(arg88)
					}
				}(func(_1401_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()), _1399_v31)
				var _1402_v33 D8
				_ = _1402_v33
				_1402_v33 = Companion_D8_.Create_DC26_(_this.F14(), _1400_v32)
				var _pat_let_tv39 = _1402_v33
				_ = _pat_let_tv39
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1391_v24), 0))
				_ = _index251
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1394_v27), 0))
				_ = _index252
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1397_v29), 0))
				_ = _index253
				var _rhs226 _dafny.Array = _1388_v23
				_ = _rhs226
				var _rhs227 D0 = _1393_v26
				_ = _rhs227
				var _rhs228 bool = (_1398_v30).IsDisjointFrom(_1398_v30)
				_ = _rhs228
				var _rhs229 D0 = func(_pat_let52_0 D0) D0 {
					return func(_1405_dt__update__tmp_h1 D0) D0 {
						return func(_pat_let55_0 bool) D0 {
							return func(_1406_dt__update_hcf1_h1 bool) D0 {
								return Companion_D0_.Create_DC1_(_1406_dt__update_hcf1_h1)
							}(_pat_let55_0)
						}(!((_this).F19()))
					}(_pat_let52_0)
				}(func(_pat_let53_0 D0) D0 {
					return func(_1403_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let54_0 bool) D0 {
							return func(_1404_dt__update_hcf1_h0 bool) D0 {
								return Companion_D0_.Create_DC1_(_1404_dt__update_hcf1_h0)
							}(_pat_let54_0)
						}((_pat_let_tv39).Dtor_cf46())
					}(_pat_let53_0)
				}(Companion_D0_.Create_DC1_((_this).Fm2(!((_this).F19()), _this.F14(), globalState))))
				_ = _rhs229
				var _rhs230 bool = _this.F14()
				_ = _rhs230
				var _lhs190 _dafny.Array = _1391_v24
				_ = _lhs190
				var _lhs191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(618), _dafny.ArrayLen((_1391_v24), 0))
				_ = _lhs191
				var _lhs192 *GlobalState = globalState
				_ = _lhs192
				var _lhs193 _dafny.Array = _1394_v27
				_ = _lhs193
				var _lhs194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_1394_v27), 0))
				_ = _lhs194
				var _lhs195 _dafny.Array = _1397_v29
				_ = _lhs195
				var _lhs196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1397_v29), 0))
				_ = _lhs196
				_1388_v23 = _rhs226
				(_lhs190).ArraySet1(_rhs227, (_lhs191).Int())
				_lhs192.F9 = _rhs228
				(_lhs193).ArraySet1(_rhs229, (_lhs194).Int())
				(_lhs195).ArraySet1(_rhs230, (_lhs196).Int())
				(globalState).F2 = _1397_v29
				(_this).F14_set_(!((_1397_v29).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1397_v29), 0))).Int()).(bool)))
				var _1407_v34 *C1
				_ = _1407_v34
				var _nw228 *C1 = New_C1_()
				_ = _nw228
				_nw228.Ctor__()
				_1407_v34 = _nw228
				var _1408_v35 _dafny.MultiSet
				_ = _1408_v35
				_1408_v35 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uwcabk")).Cardinality()))
				var _1409_v36 _dafny.Sequence
				_ = _1409_v36
				_1409_v36 = _dafny.SeqOf(_1408_v35, _1408_v35)
				_1409_v36 = _1409_v36
			} else {
				var _1410_v37 _dafny.Sequence
				_ = _1410_v37
				_1410_v37 = _dafny.UnicodeSeqOfUtf8Bytes("nenovha")
				var _1411_v38 _dafny.Sequence
				_ = _1411_v38
				_1411_v38 = _dafny.SeqOf(_this.F14(), (_this).F19(), _this.F14(), _this.F14(), _this.F14())
				var _1412_v39 _dafny.CodePoint
				_ = _1412_v39
				_1412_v39 = _dafny.CodePoint('y')
				(globalState).F3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1410_v37, _1410_v37), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if (_1411_v38).Select((Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_1411_v38).Cardinality()))).Uint32()).(bool) {
						return _1410_v37
					}
					return _1410_v37
				})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1374_v14), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1411_v38).Select((Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_1411_v38).Cardinality()))).Uint32()).(bool) {
						return _1410_v37
					}
					return _1410_v37
				})()).Cardinality()))).Uint32(), _1412_v39))).Cardinality())
				var _1413_v40 _dafny.Sequence
				_ = _1413_v40
				_1413_v40 = _dafny.SeqOf(_dafny.IntOfInt64(571))
				var _1414_v41 _dafny.Sequence
				_ = _1414_v41
				_1414_v41 = _1413_v40
				var _1415_v42 _dafny.Map
				_ = _1415_v42
				_1415_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1414_v41, _this.F14())
				var _1416_v43 _dafny.Sequence
				_ = _1416_v43
				_1416_v43 = _dafny.SeqOf(_1410_v37, _1410_v37, _dafny.UnicodeSeqOfUtf8Bytes("qvqxuld"))
				var _1417_v44 _dafny.Array
				_ = _1417_v44
				var _nwElement0_43 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1410_v37, _1410_v37)
				_ = _nwElement0_43
				var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(10))
				_ = _nw229
				(_nw229).ArraySet1(_nwElement0_43, 0)
				(_nw229).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ebteocn"), 1)
				(_nw229).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lronodbd"), 2)
				(_nw229).ArraySet1(_1410_v37, 3)
				(_nw229).ArraySet1(_1410_v37, 4)
				(_nw229).ArraySet1(Companion_Default___.Fm27(_1374_v14, true, globalState), 5)
				(_nw229).ArraySet1((func() _dafny.Sequence {
					if (func() bool {
						if (_1415_v42).Contains(_1414_v41) {
							return (_1415_v42).Get(_1414_v41).(bool)
						}
						return (_this).F19()
					})() {
						return _1410_v37
					}
					return _1410_v37
				})(), 6)
				(_nw229).ArraySet1((_1416_v43).Select((Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_1416_v43).Cardinality()))).Uint32()).(_dafny.Sequence), 7)
				(_nw229).ArraySet1((func() _dafny.Sequence {
					if (_this).F19() {
						return _1410_v37
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("evmfvq")
				})(), 8)
				(_nw229).ArraySet1(_1410_v37, 9)
				_1417_v44 = _nw229
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1417_v44), 0))
				_ = _index254
				(_1417_v44).ArraySet1(_1410_v37, (_index254).Int())
				var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_1417_v44), 0))
				_ = _index255
				(_1417_v44).ArraySet1(_1410_v37, (_index255).Int())
				(globalState).F3 = _1374_v14
				var _1418_v45 _dafny.Sequence
				_ = _1418_v45
				_1418_v45 = _dafny.SeqOf((_this).F15())
				(_this).M7(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.MultiSetOf(!(_this.F14()))), _1418_v45), _dafny.Companion_Sequence_.Equal(_1413_v40, _1413_v40), _1374_v14, _this.F14(), globalState)
				var _1419_v46 _dafny.Array
				_ = _1419_v46
				var _nw230 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(24))
				_ = _nw230
				_1419_v46 = _nw230
				var _1420_v47 _dafny.MultiSet
				_ = _1420_v47
				_1420_v47 = _dafny.MultiSetOf(_1412_v39, _1412_v39, _1412_v39, _1412_v39, _1412_v39)
				var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_1419_v46), 0))
				_ = _index256
				(_1419_v46).ArraySet1((_1374_v14).Cmp((_1420_v47).Cardinality()) == 0, (_index256).Int())
				var _1421_v48 _dafny.MultiSet
				_ = _1421_v48
				_1421_v48 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1413_v40).Cardinality()), _dafny.IntOfInt64(180))
				var _1422_v49 _dafny.Array
				_ = _1422_v49
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_42
				var _nw231 _dafny.Array
				_ = _nw231
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw231 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Sequence = (func(_1423_v14 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_1424_i9 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(271))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
								return func(arg89 _dafny.Int) interface{} {
									return coer89(arg89)
								}
							}((func(_1425_v14 _dafny.Int) func(_dafny.Int) _dafny.Map {
								return func(_1426_i10 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1425_v14, false)
								}
							})(_1423_v14)))
						}
					})(_1374_v14)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw231 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw231).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw231).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1422_v49 = _nw231
				var _1427_v50 _dafny.Map
				_ = _1427_v50
				_1427_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v14, _this.F14())
				var _1428_v51 _dafny.Sequence
				_ = _1428_v51
				_1428_v51 = _dafny.SeqOf(_1427_v50)
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1422_v49), 0))
				_ = _index257
				(_1422_v49).ArraySet1(_1428_v51, (_index257).Int())
				var _1429_v52 _dafny.Set
				_ = _1429_v52
				_1429_v52 = _dafny.SetOf((_1421_v48).Cardinality(), _1374_v14)
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_1419_v46), 0))
				_ = _index258
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1422_v49), 0))
				_ = _index259
				var _rhs231 bool = _this.F14()
				_ = _rhs231
				var _rhs232 bool = ((func() _dafny.Set {
					var _coll53 = _dafny.NewBuilder()
					_ = _coll53
					for _iter56 := _dafny.Iterate((_1429_v52).Elements()); ; {
						_compr_53, _ok56 := _iter56()
						if !_ok56 {
							break
						}
						var _1430_v53 _dafny.Int
						_1430_v53 = interface{}(_compr_53).(_dafny.Int)
						if (_1429_v52).Contains(_1430_v53) {
							_coll53.Add((_1430_v53).Minus(_dafny.IntOfInt64(216)))
						}
					}
					return _coll53.ToSet()
				}()).Cardinality()).Cmp((_1374_v14).Minus(_1374_v14)) < 0
				_ = _rhs232
				var _rhs233 _dafny.MultiSet = _1421_v48
				_ = _rhs233
				var _rhs234 _dafny.Sequence = _dafny.SeqOf(_1427_v50, (_1427_v50).Merge(_1427_v50), _1427_v50, ((_1427_v50).Update(_dafny.IntOfInt64(298), (_this).F19())).Merge((_1427_v50).Update(_1374_v14, (_this).F19())))
				_ = _rhs234
				var _lhs197 _dafny.Array = _1419_v46
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(750), _dafny.ArrayLen((_1419_v46), 0))
				_ = _lhs198
				var _lhs199 *GlobalState = globalState
				_ = _lhs199
				var _lhs200 _dafny.Array = _1422_v49
				_ = _lhs200
				var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(598), _dafny.ArrayLen((_1422_v49), 0))
				_ = _lhs201
				(_lhs197).ArraySet1(_rhs231, (_lhs198).Int())
				_lhs199.F9 = _rhs232
				_1421_v48 = _rhs233
				(_lhs200).ArraySet1(_rhs234, (_lhs201).Int())
			}
			var _1431_v54 _dafny.Map
			_ = _1431_v54
			_1431_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1374_v14)
			var _1432_v58 _dafny.Sequence
			_ = _1432_v58
			_1432_v58 = _dafny.SeqOf(_1431_v54)
			var _1433_v59 _dafny.Sequence
			_ = _1433_v59
			_1433_v59 = _dafny.SeqOf(_1374_v14, _1374_v14)
			var _1434_v60 _dafny.Set
			_ = _1434_v60
			_1434_v60 = _dafny.SetOf(_1374_v14, (_1433_v59).Select((Companion_Default___.SafeIndex(((_this).F15()).Cardinality(), _dafny.IntOfUint32((_1433_v59).Cardinality()))).Uint32()).(_dafny.Int), _1374_v14)
			var _1435_v61 _dafny.Sequence
			_ = _1435_v61
			_1435_v61 = _dafny.UnicodeSeqOfUtf8Bytes("chvabiuu")
			var _1436_v62 _dafny.Array
			_ = _1436_v62
			var _nwElement0_44 _dafny.Map = (_1431_v54).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1374_v14))
			_ = _nwElement0_44
			var _nw232 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(23))
			_ = _nw232
			(_nw232).ArraySet1(_nwElement0_44, 0)
			(_nw232).ArraySet1(_1431_v54, 1)
			(_nw232).ArraySet1(_1431_v54, 2)
			(_nw232).ArraySet1((_1431_v54).Merge(Companion_Default___.Fm1(false, _1374_v14, _this.F14(), func() _dafny.Set {
				var _coll54 = _dafny.NewBuilder()
				_ = _coll54
				for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-621), _dafny.IntOfInt64(879))); ; {
					_compr_54, _ok57 := _iter57()
					if !_ok57 {
						break
					}
					var _1437_v55 _dafny.Int
					_1437_v55 = interface{}(_compr_54).(_dafny.Int)
					if ((_dafny.IntOfInt64(-621)).Cmp(_1437_v55) <= 0) && ((_1437_v55).Cmp(_dafny.IntOfInt64(879)) < 0) {
						_coll54.Add((_1437_v55).Plus((func() _dafny.Map {
							var _coll55 = _dafny.NewMapBuilder()
							_ = _coll55
							for _iter58 := _dafny.Iterate((func() _dafny.Map {
								var _coll56 = _dafny.NewMapBuilder()
								_ = _coll56
								for _iter59 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1374_v14), _this.F14())).Keys().Elements()); ; {
									_compr_56, _ok59 := _iter59()
									if !_ok59 {
										break
									}
									var _1438_v57 _dafny.Int
									_1438_v57 = interface{}(_compr_56).(_dafny.Int)
									if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1374_v14), _this.F14())).Contains(_1438_v57) {
										_coll56.Add(Companion_Default___.SafeDivisionInt(_1438_v57, _1374_v14), (_this).F19())
									}
								}
								return _coll56.ToMap()
							}()).Keys().Elements()); ; {
								_compr_55, _ok58 := _iter58()
								if !_ok58 {
									break
								}
								var _1439_v56 _dafny.Int
								_1439_v56 = interface{}(_compr_55).(_dafny.Int)
								if (func() _dafny.Map {
									var _coll57 = _dafny.NewMapBuilder()
									_ = _coll57
									for _iter60 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1374_v14), _this.F14())).Keys().Elements()); ; {
										_compr_57, _ok60 := _iter60()
										if !_ok60 {
											break
										}
										var _1438_v57 _dafny.Int
										_1438_v57 = interface{}(_compr_57).(_dafny.Int)
										if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1374_v14), _this.F14())).Contains(_1438_v57) {
											_coll57.Add(Companion_Default___.SafeDivisionInt(_1438_v57, _1374_v14), (_this).F19())
										}
									}
									return _coll57.ToMap()
								}()).Contains(_1439_v56) {
									_coll55.Add(Companion_Default___.SafeDivisionInt(_1439_v56, _dafny.IntOfInt64(403)), _1374_v14)
								}
							}
							return _coll55.ToMap()
						}()).Cardinality()))
					}
				}
				return _coll54.ToSet()
			}(), globalState)), 3)
			(_nw232).ArraySet1(((_1432_v58).Select((Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_1432_v58).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_1431_v54), 4)
			(_nw232).ArraySet1(_1431_v54, 5)
			(_nw232).ArraySet1((func() _dafny.Map {
				if (_this).F19() {
					return _1431_v54
				}
				return _1431_v54
			})(), 6)
			(_nw232).ArraySet1(_1431_v54, 7)
			(_nw232).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1374_v14), 8)
			(_nw232).ArraySet1(_1431_v54, 9)
			(_nw232).ArraySet1(_1431_v54, 10)
			(_nw232).ArraySet1((_1431_v54).Update(false, _1374_v14), 11)
			(_nw232).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1374_v14)).Merge(Companion_Default___.Fm1((_this).F19(), _1374_v14, _this.F14(), _1434_v60, globalState)), 12)
			(_nw232).ArraySet1(_1431_v54, 13)
			(_nw232).ArraySet1((_1431_v54).Merge(_1431_v54), 14)
			(_nw232).ArraySet1(_1431_v54, 15)
			(_nw232).ArraySet1(_1431_v54, 16)
			(_nw232).ArraySet1(_1431_v54, 17)
			(_nw232).ArraySet1(_1431_v54, 18)
			(_nw232).ArraySet1(_1431_v54, 19)
			(_nw232).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F19(), _1374_v14), 20)
			(_nw232).ArraySet1((_1431_v54).Merge(_1431_v54), 21)
			(_nw232).ArraySet1((_1431_v54).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F19()), _dafny.IntOfUint32((_1435_v61).Cardinality()))), 22)
			_1436_v62 = _nw232
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1436_v62), 0))
			_ = _index260
			(_1436_v62).ArraySet1(_1431_v54, (_index260).Int())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1436_v62), 0))
			_ = _index261
			(_1436_v62).ArraySet1(_1431_v54, (_index261).Int())
			_1374_v14 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F19() {
					return _1374_v14
				}
				return (_dafny.Zero).Minus(_1374_v14)
			})())
			var _1440_v63 *C1
			_ = _1440_v63
			var _nw233 *C1 = New_C1_()
			_ = _nw233
			_nw233.Ctor__()
			_1440_v63 = _nw233
		} else {
			(globalState).F9 = _this.F14()
			if (_this).F19() {
				var _1441_v64 _dafny.Sequence
				_ = _1441_v64
				_1441_v64 = _dafny.UnicodeSeqOfUtf8Bytes("frcf")
				_1441_v64 = _dafny.Companion_Sequence_.Concatenate(_1441_v64, _1441_v64)
				var _1442_v65 _dafny.Map
				_ = _1442_v65
				_1442_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), ((_this).F15()).Intersection((_this).F15()))
				_1442_v65 = (_1442_v65).Update(!((_this).F19()) || (!((_this).F19())), (_this).F15())
				var _1443_v66 _dafny.Array
				_ = _1443_v66
				var _nw234 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
				_ = _nw234
				_1443_v66 = _nw234
				var _1444_v67 D4
				_ = _1444_v67
				_1444_v67 = Companion_D4_.Create_DC13_(_1441_v64, _1374_v14, (_this).F19(), _1443_v66)
				var _1445_v68 _dafny.MultiSet
				_ = _1445_v68
				_1445_v68 = _dafny.MultiSetOf(_1374_v14, _1374_v14, _1374_v14)
				_1441_v64 = (_this).Fm3(Companion_Default___.Fm0(_this.F14(), _dafny.IntOfUint32(((_1444_v67).Dtor_cf22()).Cardinality()), _1374_v14, _1445_v68, globalState), (_dafny.SetOf(_1374_v14)).Cardinality(), _this.F14(), _1374_v14, globalState)
				var _1446_v69 _dafny.CodePoint
				_ = _1446_v69
				_1446_v69 = _dafny.CodePoint('m')
				var _1447_v70 _dafny.Map
				_ = _1447_v70
				_1447_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1446_v69, _this.F14())
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1443_v66), 0))
				_ = _index262
				(_1443_v66).ArraySet1(false, (_index262).Int())
				var _1448_v71 _dafny.Sequence
				_ = _1448_v71
				_1448_v71 = _dafny.SeqOf((_this).F19())
				var _1449_v72 _dafny.Map
				_ = _1449_v72
				_1449_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1441_v64).Cardinality()), true)
				var _1450_v73 _dafny.Map
				_ = _1450_v73
				_1450_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm22(_1374_v14, _dafny.SeqOf(_1449_v72), _this.F14(), (_this).F19(), globalState), _1443_v66)
				var _1451_v74 _dafny.Set
				_ = _1451_v74
				_1451_v74 = _dafny.SetOf(_1445_v68)
				var _1452_v75 _dafny.MultiSet
				_ = _1452_v75
				_1452_v75 = _dafny.MultiSetOf(_1448_v71, _1448_v71, _1448_v71, _1448_v71)
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1443_v66), 0))
				_ = _index263
				var _rhs235 bool = (_this).F19()
				_ = _rhs235
				var _rhs236 _dafny.Map = (func() _dafny.Map {
					if (_1448_v71).Select((Companion_Default___.SafeIndex((_1450_v73).Cardinality(), _dafny.IntOfUint32((_1448_v71).Cardinality()))).Uint32()).(bool) {
						return (_1447_v70).Merge(_1447_v70)
					}
					return (Companion_Default___.Fm46(_1374_v14, (_this).F19(), globalState)).Merge(_1447_v70)
				})()
				_ = _rhs236
				var _rhs237 bool = true
				_ = _rhs237
				var _rhs238 _dafny.Int = (_this).Fm17(_1374_v14, (_1451_v74).IsDisjointFrom(_1451_v74), _1452_v75, _1374_v14, globalState)
				_ = _rhs238
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				var _lhs203 _dafny.Array = _1443_v66
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.ArrayLen((_1443_v66), 0))
				_ = _lhs204
				_lhs202.F9 = _rhs235
				_1447_v70 = _rhs236
				(_lhs203).ArraySet1(_rhs237, (_lhs204).Int())
				_1374_v14 = _rhs238
				(globalState).F3 = Companion_Default___.SafeDivisionInt(_1374_v14, _1374_v14)
			} else {
				var _1453_v76 _dafny.Array
				_ = _1453_v76
				var _nw235 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw235
				_1453_v76 = _nw235
				var _1454_v77 _dafny.Set
				_ = _1454_v77
				_1454_v77 = _dafny.SetOf(_1453_v76, _1453_v76)
				(globalState).F3 = (_1374_v14).Times((_1454_v77).Cardinality())
				var _1455_v78 _dafny.Sequence
				_ = _1455_v78
				_1455_v78 = _dafny.UnicodeSeqOfUtf8Bytes("wkk")
				var _1456_v79 _dafny.Set
				_ = _1456_v79
				_1456_v79 = _dafny.SetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_1455_v78, _1455_v78))
				(globalState).F6 = _1456_v79
				(globalState).F3 = (_dafny.IntOfUint32((_1455_v78).Cardinality())).Times(_1374_v14)
				var _1457_v80 _dafny.Map
				_ = _1457_v80
				_1457_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _1374_v14)
				var _1458_v81 _dafny.Sequence
				_ = _1458_v81
				_1458_v81 = _dafny.SeqOf(_1374_v14, _dafny.IntOfUint32((_1455_v78).Cardinality()))
				var _1459_v82 _dafny.Set
				_ = _1459_v82
				_1459_v82 = _dafny.SetOf(_1458_v81)
				var _1460_v83 _dafny.Map
				_ = _1460_v83
				_1460_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm23(_1374_v14, _this.F14(), _1457_v80, _1459_v82, globalState)) || (_this.F14()), _1374_v14)
				_1460_v83 = (_1460_v83).Update(_this.F14(), _1374_v14)
				(globalState).F9 = (_this).F19()
			}
			var _1461_v84 _dafny.Array
			_ = _1461_v84
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_43
			var _nw236 _dafny.Array
			_ = _nw236
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw236 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Sequence = func(_1462_i11 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F19()), _dafny.SeqOf(!((_this).F19())))
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw236 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw236).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw236).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1461_v84 = _nw236
			var _1463_v85 _dafny.Sequence
			_ = _1463_v85
			_1463_v85 = _dafny.SeqOf(_this.F14(), _this.F14(), (_this).F19())
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1461_v84), 0))
			_ = _index264
			(_1461_v84).ArraySet1(_1463_v85, (_index264).Int())
			var _1464_v86 _dafny.Map
			_ = _1464_v86
			_1464_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_this.F14())), _1374_v14)
			var _1465_v87 _dafny.Sequence
			_ = _1465_v87
			_1465_v87 = _dafny.UnicodeSeqOfUtf8Bytes("pw")
			var _1466_v88 _dafny.Sequence
			_ = _1466_v88
			_1466_v88 = _dafny.SeqOf(_1463_v85)
			var _1467_v89 _dafny.Sequence
			_ = _1467_v89
			_1467_v89 = _dafny.SeqOf(_1374_v14, (_this).Fm17((func() _dafny.Int {
				if (_1464_v86).Contains(_this.F14()) {
					return (_1464_v86).Get(_this.F14()).(_dafny.Int)
				}
				return _1374_v14
			})(), _this.F14(), _dafny.MultiSetOf(_1463_v85, _dafny.Companion_Sequence_.Update(_1463_v85, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1465_v87).Cardinality()), _dafny.IntOfUint32((_1463_v85).Cardinality()))).Uint32(), (_this).F19()), _1463_v85, _1463_v85, _1463_v85), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1466_v88, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1374_v14), _dafny.IntOfUint32((_1466_v88).Cardinality()))).Uint32(), _dafny.SeqOf((_this).F19()))).Cardinality())), globalState))
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1461_v84), 0))
			_ = _index265
			var _rhs239 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1463_v85, (Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_1463_v85).Cardinality()))).Uint32(), (_this).F19())
			_ = _rhs239
			var _rhs240 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(400))).Uint32(), func(coer90 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg90 _dafny.Int) interface{} {
					return coer90(arg90)
				}
			}((func(_1468_v14 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1469_i12 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1468_v14, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ldnplt")).Cardinality()))
				}
			})(_1374_v14)))
			_ = _rhs240
			var _lhs205 _dafny.Array = _1461_v84
			_ = _lhs205
			var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_1461_v84), 0))
			_ = _lhs206
			(_lhs205).ArraySet1(_rhs239, (_lhs206).Int())
			_1467_v89 = _rhs240
			(globalState).F9 = !(!((_this.F14()) == (!(_this.F14()))))
			(globalState).F9 = _this.F14()
		}
		var _1470_v90 _dafny.CodePoint
		_ = _1470_v90
		_1470_v90 = _dafny.CodePoint('o')
		_1470_v90 = _1470_v90
		var _1471_v91 _dafny.Array
		_ = _1471_v91
		var _nw237 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
		_ = _nw237
		_1471_v91 = _nw237
		var _1472_v92 _dafny.Sequence
		_ = _1472_v92
		_1472_v92 = _dafny.SeqOf((_this).F19(), true, (_this).F19(), _this.F14(), (_this).F19())
		var _1473_v93 D1
		_ = _1473_v93
		_1473_v93 = Companion_D1_.Create_DC3_(_dafny.IntOfInt64(338))
		var _1474_v94 _dafny.Map
		_ = _1474_v94
		_1474_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v14, _1473_v93)
		var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1471_v91), 0))
		_ = _index266
		(_1471_v91).ArraySet1((_1472_v92).Select((Companion_Default___.SafeIndex((_1474_v94).Cardinality(), _dafny.IntOfUint32((_1472_v92).Cardinality()))).Uint32()).(bool), (_index266).Int())
		var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1471_v91), 0))
		_ = _index267
		(_1471_v91).ArraySet1((_this).F19(), (_index267).Int())
		var _1475_v95 _dafny.Map
		_ = _1475_v95
		_1475_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.IntOfInt64(101))
		var _1476_v96 _dafny.Map
		_ = _1476_v96
		_1476_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1374_v14, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("voqhvj"), (Companion_Default___.SafeIndex(_1374_v14, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("voqhvj")).Cardinality()))).Uint32(), _1470_v90))
		var _1477_v97 _dafny.MultiSet
		_ = _1477_v97
		_1477_v97 = _dafny.MultiSetOf((_1475_v95).Cardinality(), _1374_v14, (_1476_v96).Cardinality())
		_1470_v90 = Companion_Default___.Fm25(_1374_v14, ((_dafny.Zero).Minus(_1374_v14)).Minus((_1477_v97).Cardinality()), !((_1471_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1471_v91), 0))).Int()).(bool)) || ((_this).F19()), _1475_v95, globalState)
	}
}
func (_this *C4) M7(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) {
	{
		var _1478_v0 *C1
		_ = _1478_v0
		var _nw238 *C1 = New_C1_()
		_ = _nw238
		_nw238.Ctor__()
		_1478_v0 = _nw238
		var _1479_v1 _dafny.Array
		_ = _1479_v1
		var _nw239 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw239
		_1479_v1 = _nw239
		var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))
		_ = _index268
		(_1479_v1).ArraySet1(p1, (_index268).Int())
		var _1480_v2 _dafny.Set
		_ = _1480_v2
		_1480_v2 = _dafny.SetOf((_this).F15())
		var _1481_v3 _dafny.Array
		_ = _1481_v3
		var _nwElement0_45 _dafny.Set = _1480_v2
		_ = _nwElement0_45
		var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(7))
		_ = _nw240
		(_nw240).ArraySet1(_nwElement0_45, 0)
		(_nw240).ArraySet1((Companion_D11_.Create_DC33_(_1480_v2)).Dtor_cf57(), 1)
		(_nw240).ArraySet1(_1480_v2, 2)
		(_nw240).ArraySet1(_1480_v2, 3)
		(_nw240).ArraySet1((_1480_v2).Difference(_1480_v2), 4)
		(_nw240).ArraySet1(_1480_v2, 5)
		(_nw240).ArraySet1((_1480_v2).Difference(_dafny.SetOf((_this).F15(), (_this).F15())), 6)
		_1481_v3 = _nw240
		var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_1481_v3), 0))
		_ = _index269
		(_1481_v3).ArraySet1((Companion_Default___.Fm47(globalState)).Difference(_1480_v2), (_index269).Int())
		var _1482_v4 _dafny.Set
		_ = _1482_v4
		_1482_v4 = _dafny.SetOf(p2)
		var _1483_v5 _dafny.Sequence
		_ = _1483_v5
		_1483_v5 = _dafny.UnicodeSeqOfUtf8Bytes("ovck")
		var _1484_v6 _dafny.Sequence
		_ = _1484_v6
		_1484_v6 = _dafny.SeqOf((_dafny.Zero).Minus((_1482_v4).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-295)), p2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1483_v5).Cardinality()), _1478_v0)).Cardinality())
		var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))
		_ = _index270
		var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_1481_v3), 0))
		_ = _index271
		var _rhs241 bool = (!(p3) || (p3)) && (_dafny.Companion_Sequence_.IsPrefixOf(_1484_v6, _1484_v6))
		_ = _rhs241
		var _rhs242 _dafny.Set = _1480_v2
		_ = _rhs242
		var _rhs243 bool = true
		_ = _rhs243
		var _lhs207 _dafny.Array = _1479_v1
		_ = _lhs207
		var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))
		_ = _lhs208
		var _lhs209 _dafny.Array = _1481_v3
		_ = _lhs209
		var _lhs210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(43), _dafny.ArrayLen((_1481_v3), 0))
		_ = _lhs210
		var _lhs211 *GlobalState = globalState
		_ = _lhs211
		(_lhs207).ArraySet1(_rhs241, (_lhs208).Int())
		(_lhs209).ArraySet1(_rhs242, (_lhs210).Int())
		_lhs211.F9 = _rhs243
		var _rhs244 _dafny.Int = p2
		_ = _rhs244
		var _rhs245 _dafny.Int = _dafny.IntOfInt64(538)
		_ = _rhs245
		var _lhs212 *GlobalState = globalState
		_ = _lhs212
		var _lhs213 *GlobalState = globalState
		_ = _lhs213
		_lhs212.F3 = _rhs244
		_lhs213.F3 = _rhs245
		var _1485_v7 _dafny.CodePoint
		_ = _1485_v7
		_1485_v7 = _dafny.CodePoint('e')
		var _1486_v8 _dafny.Map
		_ = _1486_v8
		_1486_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1483_v5).Cardinality())))
		var _1487_v9 D4
		_ = _1487_v9
		_1487_v9 = Companion_D4_.Create_DC13_(_dafny.UnicodeSeqOfUtf8Bytes("sikdur"), p2, (_this).F19(), _1479_v1)
		var _1488_v11 _dafny.Map
		_ = _1488_v11
		_1488_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p3)
		var _1489_v12 _dafny.Sequence
		_ = _1489_v12
		_1489_v12 = _dafny.SeqOf(func() _dafny.Map {
			var _coll58 = _dafny.NewMapBuilder()
			_ = _coll58
			for _iter61 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-314), _dafny.IntOfInt64(-878))); ; {
				_compr_58, _ok61 := _iter61()
				if !_ok61 {
					break
				}
				var _1490_v10 _dafny.Int
				_1490_v10 = interface{}(_compr_58).(_dafny.Int)
				if ((_dafny.IntOfInt64(-314)).Cmp(_1490_v10) <= 0) && ((_1490_v10).Cmp(_dafny.IntOfInt64(-878)) < 0) {
					_coll58.Add((_1490_v10).Plus(p2), _this.F14())
				}
			}
			return _coll58.ToMap()
		}(), _1488_v11, _1488_v11, _1488_v11)
		var _1491_v13 D9
		_ = _1491_v13
		_1491_v13 = Companion_D9_.Create_DC30_(_1485_v7, p1, (_1486_v8).Cardinality(), Companion_Default___.Fm22((_1487_v9).Dtor_cf23(), _1489_v12, (_this).F19(), (_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool), globalState), p2)
		var _1492_v14 _dafny.Map
		_ = _1492_v14
		_1492_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1491_v13, p2)
		var _hi4 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(934))).Uint32(), func(coer91 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg91 _dafny.Int) interface{} {
				return coer91(arg91)
			}
		}((func(_1493_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1494_i1 _dafny.Int) _dafny.CodePoint {
				return _1493_v7
			}
		})(_1485_v7)))).Cardinality())).Plus(p2)
		_ = _hi4
		for _1495_i0 := ((_1492_v14).Update(_1491_v13, p2)).Cardinality(); _1495_i0.Cmp(_hi4) < 0; _1495_i0 = _1495_i0.Plus(_dafny.One) {
			var _1496_v15 _dafny.MultiSet
			_ = _1496_v15
			_1496_v15 = _dafny.MultiSetOf(p2)
			var _1497_v16 D2
			_ = _1497_v16
			_1497_v16 = Companion_D2_.Create_DC9_((_this).F19(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(41))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg92 _dafny.Int) interface{} {
					return coer92(arg92)
				}
			}(func(_1498_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			}))).Cardinality()), _1485_v7)
			var _1499_v17 _dafny.Array
			_ = _1499_v17
			var _nwElement0_46 D2 = _1497_v16
			_ = _nwElement0_46
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(6))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_46, 0)
			(_nw241).ArraySet1(_1497_v16, 1)
			(_nw241).ArraySet1(_1497_v16, 2)
			(_nw241).ArraySet1(_1497_v16, 3)
			(_nw241).ArraySet1(Companion_D2_.Create_DC9_(p3, p2, _1485_v7), 4)
			(_nw241).ArraySet1(Companion_D2_.Create_DC9_(p1, _1495_i0, _1485_v7), 5)
			_1499_v17 = _nw241
			var _1500_v18 _dafny.Sequence
			_ = _1500_v18
			_1500_v18 = _dafny.SeqOf(_1499_v17)
			var _1501_v19 _dafny.Array
			_ = _1501_v19
			var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw242
			_1501_v19 = _nw242
			var _1502_v20 D5
			_ = _1502_v20
			_1502_v20 = Companion_D5_.Create_DC16_(_1501_v19)
			var _1503_v21 D5
			_ = _1503_v21
			_1503_v21 = Companion_D5_.Create_DC19_(_1502_v20)
			var _1504_v22 _dafny.Map
			_ = _1504_v22
			_1504_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm2(_this.F14(), p3, globalState), _1502_v20)
			var _1505_v23 D5
			_ = _1505_v23
			_1505_v23 = Companion_D5_.Create_DC19_((func() D5 {
				if (_1504_v22).Contains((_this).F19()) {
					return (_1504_v22).Get((_this).F19()).(D5)
				}
				return _1503_v21
			})())
			var _1506_v24 D7
			_ = _1506_v24
			_1506_v24 = Companion_D7_.Create_DC22_(_1479_v1, _1500_v18, _dafny.SeqOf(_1479_v1, _1479_v1), _1479_v1, _1505_v23)
			var _1507_v25 _dafny.Sequence
			_ = _1507_v25
			_1507_v25 = _dafny.SeqOf(_1506_v24)
			var _1508_v26 _dafny.Array
			_ = _1508_v26
			var _nwElement0_47 _dafny.Int = p2
			_ = _nwElement0_47
			var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(11))
			_ = _nw243
			(_nw243).ArraySet1(_nwElement0_47, 0)
			(_nw243).ArraySet1((p2).Plus(p2), 1)
			(_nw243).ArraySet1(p2, 2)
			(_nw243).ArraySet1(_1495_i0, 3)
			(_nw243).ArraySet1(((func() _dafny.Int {
				if ((_this).F15()).Contains((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool)) {
					return ((_this).F15()).Multiplicity((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool))
				}
				return Companion_Default___.Fm0((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool), _dafny.IntOfInt64(24), _1495_i0, _1496_v15, globalState)
			})()).Times(p2), 4)
			(_nw243).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p2), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1495_i0, p2)).Cardinality()), 5)
			(_nw243).ArraySet1(_dafny.IntOfUint32((_1483_v5).Cardinality()), 6)
			(_nw243).ArraySet1(Companion_Default___.Fm0(_this.F14(), p2, _1495_i0, _1496_v15, globalState), 7)
			(_nw243).ArraySet1(p2, 8)
			(_nw243).ArraySet1((_1495_i0).Minus(_dafny.IntOfUint32((_1483_v5).Cardinality())), 9)
			(_nw243).ArraySet1(_dafny.IntOfUint32((_1507_v25).Cardinality()), 10)
			_1508_v26 = _nw243
			var _1509_v27 _dafny.Sequence
			_ = _1509_v27
			_1509_v27 = _dafny.SeqOf(_this.F14(), _this.F14(), _this.F14(), p3, _this.F14())
			var _1510_v28 _dafny.MultiSet
			_ = _1510_v28
			_1510_v28 = _dafny.MultiSetOf(_1509_v27)
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1501_v19), 0))
			_ = _index272
			(_1501_v19).ArraySet1((_this).Fm17(p2, !((_this).F19()), _1510_v28, (_dafny.Zero).Minus(Companion_Default___.Fm0((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool), _1495_i0, _1495_i0, _1496_v15, globalState)), globalState), (_index272).Int())
			var _1511_v29 D12
			_ = _1511_v29
			_1511_v29 = Companion_D12_.Create_DC37_(_1510_v28)
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1501_v19), 0))
			_ = _index273
			(_1501_v19).ArraySet1(Companion_Default___.SafeModuloInt((p2).Plus(_dafny.IntOfInt64(528)), (_this).Fm17(_1495_i0, (_1509_v27).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1509_v27).Cardinality()))).Uint32()).(bool), (_1511_v29).Dtor_cf63(), _1495_i0, globalState)), (_index273).Int())
			var _1512_v30 _dafny.Map
			_ = _1512_v30
			_1512_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_1501_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(23), _dafny.ArrayLen((_1501_v19), 0))).Int()).(_dafny.Int))
			(globalState).F9 = (_dafny.IntOfInt64(128)).Cmp((_1495_i0).Plus((func() _dafny.Int {
				if (_1512_v30).Contains((_this).F19()) {
					return (_1512_v30).Get((_this).F19()).(_dafny.Int)
				}
				return p2
			})())) > 0
			var _rhs246 _dafny.CodePoint = _1485_v7
			_ = _rhs246
			var _rhs247 _dafny.CodePoint = _1485_v7
			_ = _rhs247
			var _rhs248 _dafny.Array = _1501_v19
			_ = _rhs248
			var _rhs249 _dafny.Int = (func() _dafny.Int {
				if (_1512_v30).Contains(p1) {
					return (_1512_v30).Get(p1).(_dafny.Int)
				}
				return (p2).Minus(p2)
			})()
			_ = _rhs249
			var _lhs214 *GlobalState = globalState
			_ = _lhs214
			_1485_v7 = _rhs246
			_1485_v7 = _rhs247
			_1501_v19 = _rhs248
			_lhs214.F3 = _rhs249
			(globalState).F9 = _dafny.Companion_Sequence_.Equal(_1509_v27, _1509_v27)
		}
		var _1513_v31 *C2
		_ = _1513_v31
		var _nw244 *C2 = New_C2_()
		_ = _nw244
		_nw244.Ctor__(_1488_v11)
		_1513_v31 = _nw244
		_1513_v31 = _1513_v31
		var _1514_v32 _dafny.Sequence
		_ = _1514_v32
		_1514_v32 = _dafny.SeqOf(p1)
		var _1515_v33 _dafny.Map
		_ = _1515_v33
		_1515_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool), _dafny.Companion_Sequence_.Concatenate(_1514_v32, _1514_v32))
		var _rhs250 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_1515_v33).Contains((_this.F14()) && (p1)) {
				return (_1515_v33).Get((_this.F14()) && (p1)).(_dafny.Sequence)
			}
			return _1514_v32
		})()).Cardinality())
		_ = _rhs250
		var _rhs251 bool = !((_1479_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(886), _dafny.ArrayLen((_1479_v1), 0))).Int()).(bool)) || ((_this).F19())
		_ = _rhs251
		var _rhs252 bool = (_this).F19()
		_ = _rhs252
		var _lhs215 *GlobalState = globalState
		_ = _lhs215
		var _lhs216 *GlobalState = globalState
		_ = _lhs216
		var _lhs217 *GlobalState = globalState
		_ = _lhs217
		_lhs215.F3 = _rhs250
		_lhs216.F9 = _rhs251
		_lhs217.F9 = _rhs252
	}
}
func (_this *C4) F19() bool {
	{
		return _this._f19
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f14 bool
	_f15 _dafny.MultiSet
	_f18 _dafny.Int
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f14 = false
	_this._f15 = _dafny.EmptyMultiSet
	_this._f18 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F14() bool {
	return _this._f14
}
func (_this *C5) F14_set_(value bool) {
	_this._f14 = value
}
func (_this *C5) F15() _dafny.MultiSet {
	return _this._f15
}
func (_this *C5) Ctor__(f18 _dafny.Int, f14 bool, f15 _dafny.MultiSet) {
	{
		(_this)._f18 = f18
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C5) Fm2(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return _this.F14()
	}
}
func (_this *C5) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("h"), _dafny.UnicodeSeqOfUtf8Bytes("e"))
	}
}
func (_this *C5) Fm15(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("y"), _dafny.UnicodeSeqOfUtf8Bytes("kfx")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(654))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg93 _dafny.Int) interface{} {
				return coer93(arg93)
			}
		}(func(_1516_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(25))).Uint32(), func(coer94 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}(func(_1517_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('o')
		}))))
	}
}
func (_this *C5) Fm16(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) bool {
	{
		return !(_this.F14())
	}
}
func (_this *C5) M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _1518_v0 _dafny.Sequence
		_ = _1518_v0
		_1518_v0 = _dafny.SeqOf(_this.F14())
		var _1519_v1 D2
		_ = _1519_v1
		_1519_v1 = Companion_D2_.Create_DC7_(_1518_v0)
		var _1520_v2 _dafny.Map
		_ = _1520_v2
		_1520_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_1519_v1).Dtor_cf9())
		var _1521_v3 _dafny.Sequence
		_ = _1521_v3
		_1521_v3 = _dafny.SeqOf(_1518_v0, _1518_v0, _1518_v0, _1518_v0, _1518_v0)
		_1520_v2 = (_1520_v2).Update((_this).F18(), (_1521_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(840), _dafny.IntOfUint32((_1521_v3).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _1522_v4 _dafny.Sequence
		_ = _1522_v4
		_1522_v4 = _dafny.SeqOf(_dafny.SetOf(true, _this.F14(), _this.F14(), true, _this.F14()))
		var _1523_v5 *C4
		_ = _1523_v5
		var _nw245 *C4 = New_C4_()
		_ = _nw245
		_nw245.Ctor__(_this.F14(), (_dafny.IntOfUint32((_1521_v3).Cardinality())).Cmp(_dafny.IntOfUint32((_1522_v4).Cardinality())) > 0, (_this).F15())
		_1523_v5 = _nw245
		r0 = _dafny.IntOfUint32(((_1521_v3).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1521_v3).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
		r0 = (_this).F18()
		var _1524_v6 _dafny.Array
		_ = _1524_v6
		var _nw246 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw246
		_1524_v6 = _nw246
		var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))
		_ = _index274
		(_1524_v6).ArraySet1(_this.F14(), (_index274).Int())
		var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))
		_ = _index275
		(_1524_v6).ArraySet1((func() bool {
			if _this.F14() {
				return _this.F14()
			}
			return (_1523_v5).F19()
		})(), (_index275).Int())
		var _1525_v7 _dafny.Map
		_ = _1525_v7
		_1525_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F18())
		var _1526_v8 _dafny.Map
		_ = _1526_v8
		_1526_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1523_v5).F19(), (_1525_v7).Cardinality())
		var _1527_v9 _dafny.Sequence
		_ = _1527_v9
		_1527_v9 = _dafny.UnicodeSeqOfUtf8Bytes("lm")
		var _1528_v10 _dafny.Sequence
		_ = _1528_v10
		_1528_v10 = _dafny.SeqOf((_this).F18())
		var _1529_v11 _dafny.Array
		_ = _1529_v11
		var _nwElement0_48 _dafny.Int = (_this).F18()
		_ = _nwElement0_48
		var _nw247 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(26))
		_ = _nw247
		(_nw247).ArraySet1(_nwElement0_48, 0)
		(_nw247).ArraySet1((_this).F18(), 1)
		(_nw247).ArraySet1((_this).F18(), 2)
		(_nw247).ArraySet1((_this).F18(), 3)
		(_nw247).ArraySet1((_this).F18(), 4)
		(_nw247).ArraySet1((_this).F18(), 5)
		(_nw247).ArraySet1((_this).F18(), 6)
		(_nw247).ArraySet1((_this).F18(), 7)
		(_nw247).ArraySet1((_this).F18(), 8)
		(_nw247).ArraySet1((_1526_v8).Cardinality(), 9)
		(_nw247).ArraySet1(_dafny.IntOfUint32((_1527_v9).Cardinality()), 10)
		(_nw247).ArraySet1(_dafny.IntOfInt64(-604), 11)
		(_nw247).ArraySet1(_dafny.IntOfUint32((_1528_v10).Cardinality()), 12)
		(_nw247).ArraySet1(_dafny.IntOfInt64(817), 13)
		(_nw247).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool), (_1523_v5).F19())).Cardinality()), 14)
		(_nw247).ArraySet1((_this).F18(), 15)
		(_nw247).ArraySet1((_this).F18(), 16)
		(_nw247).ArraySet1((_this).F18(), 17)
		(_nw247).ArraySet1((_this).F18(), 18)
		(_nw247).ArraySet1(_dafny.IntOfInt64(975), 19)
		(_nw247).ArraySet1((_this).F18(), 20)
		(_nw247).ArraySet1(_dafny.IntOfInt64(-187), 21)
		(_nw247).ArraySet1((_this).F18(), 22)
		(_nw247).ArraySet1((_this).F18(), 23)
		(_nw247).ArraySet1((_1528_v10).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1528_v10).Cardinality()))).Uint32()).(_dafny.Int), 24)
		(_nw247).ArraySet1((_this).F18(), 25)
		_1529_v11 = _nw247
		var _1530_v12 D5
		_ = _1530_v12
		_1530_v12 = Companion_D5_.Create_DC16_(_1529_v11)
		var _1531_v13 _dafny.Map
		_ = _1531_v13
		_1531_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1527_v9, _1529_v11)
		var _1532_v14 _dafny.Map
		_ = _1532_v14
		_1532_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1523_v5).F19(), _1529_v11)
		var _pat_let_tv40 = _1529_v11
		_ = _pat_let_tv40
		var _1533_v15 _dafny.Array
		_ = _1533_v15
		var _nwElement0_49 _dafny.Array = (func(_pat_let56_0 D5) D5 {
			return func(_1534_dt__update__tmp_h0 D5) D5 {
				return func(_pat_let57_0 _dafny.Array) D5 {
					return func(_1535_dt__update_hcf29_h0 _dafny.Array) D5 {
						return Companion_D5_.Create_DC16_(_1535_dt__update_hcf29_h0)
					}(_pat_let57_0)
				}(_pat_let_tv40)
			}(_pat_let56_0)
		}(_1530_v12)).Dtor_cf29()
		_ = _nwElement0_49
		var _nw248 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(26))
		_ = _nw248
		(_nw248).ArraySet1(_nwElement0_49, 0)
		(_nw248).ArraySet1(_1529_v11, 1)
		(_nw248).ArraySet1(_1529_v11, 2)
		(_nw248).ArraySet1(_1529_v11, 3)
		(_nw248).ArraySet1(_1529_v11, 4)
		(_nw248).ArraySet1((func() _dafny.Array {
			if (_1523_v5).F19() {
				return _1529_v11
			}
			return _1529_v11
		})(), 5)
		(_nw248).ArraySet1(_1529_v11, 6)
		(_nw248).ArraySet1(_1529_v11, 7)
		(_nw248).ArraySet1(_1529_v11, 8)
		(_nw248).ArraySet1(_1529_v11, 9)
		(_nw248).ArraySet1(_1529_v11, 10)
		(_nw248).ArraySet1((func() _dafny.Array {
			if (_1531_v13).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg95 _dafny.Int) interface{} {
					return coer95(arg95)
				}
			}(func(_1536_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('m')
			}))) {
				return (_1531_v13).Get(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg96 _dafny.Int) interface{} {
						return coer96(arg96)
					}
				}(func(_1537_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))).(_dafny.Array)
			}
			return _1529_v11
		})(), 11)
		(_nw248).ArraySet1(_1529_v11, 12)
		(_nw248).ArraySet1(_1529_v11, 13)
		(_nw248).ArraySet1(_1529_v11, 14)
		(_nw248).ArraySet1(_1529_v11, 15)
		(_nw248).ArraySet1(_1529_v11, 16)
		(_nw248).ArraySet1(_1529_v11, 17)
		(_nw248).ArraySet1(_1529_v11, 18)
		(_nw248).ArraySet1(_1529_v11, 19)
		(_nw248).ArraySet1(_1529_v11, 20)
		(_nw248).ArraySet1(_1529_v11, 21)
		(_nw248).ArraySet1(_1529_v11, 22)
		(_nw248).ArraySet1(_1529_v11, 23)
		(_nw248).ArraySet1((func() _dafny.Array {
			if (_1532_v14).Contains((_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool)) {
				return (_1532_v14).Get((_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool)).(_dafny.Array)
			}
			return _1529_v11
		})(), 24)
		(_nw248).ArraySet1(_1529_v11, 25)
		_1533_v15 = _nw248
		var _nw249 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
		_ = _nw249
		_1533_v15 = _nw249
		r0 = (_this).F18()
		r1 = (_this.F14()) && (_this.F14())
		var _1538_v16 _dafny.Sequence
		_ = _1538_v16
		_1538_v16 = _dafny.SeqOf(_1524_v6)
		var _1539_v17 D4
		_ = _1539_v17
		_1539_v17 = Companion_D4_.Create_DC13_(_1527_v9, (_this).F18(), (_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool), (_1538_v16).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1538_v16).Cardinality()))).Uint32()).(_dafny.Array))
		var _1540_v18 _dafny.Map
		_ = _1540_v18
		_1540_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool))).Update((_this).F18(), (_1523_v5).F19())).Update((_1539_v17).Dtor_cf23(), _this.F14()), (_this).F18())
		r2 = _1540_v18
		r3 = (_1524_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1524_v6), 0))).Int()).(bool)
		return r0, r1, r2, r3
	}
}
func (_this *C5) M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1541_v0 _dafny.Array
		_ = _1541_v0
		var _len0_44 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_44
		var _nw250 _dafny.Array
		_ = _nw250
		if _len0_44.Cmp(_dafny.Zero) == 0 {
			_nw250 = _dafny.NewArray(_len0_44)
		} else {
			var _init44 func(_dafny.Int) _dafny.Int = func(_1542_i1 _dafny.Int) _dafny.Int {
				return (_1542_i1).Times((_dafny.MultiSetOf(_dafny.CodePoint('b'))).Cardinality())
			}
			_ = _init44
			var _element0_44 = _init44(_dafny.Zero)
			_ = _element0_44
			_nw250 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
			(_nw250).ArraySet1(_element0_44, 0)
			var _nativeLen0_44 = (_len0_44).Int()
			_ = _nativeLen0_44
			for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
				(_nw250).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
			}
		}
		_1541_v0 = _nw250
		for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1541_v0), 0))); ; {
			_guard_loop_3, _ok62 := _iter62()
			if !_ok62 {
				break
			}
			var _1543_i0 _dafny.Int
			_1543_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1543_i0).Sign() != -1) && ((_1543_i0).Cmp(_dafny.ArrayLen((_1541_v0), 0)) < 0)) {
				(_1541_v0).ArraySet1(Companion_Default___.SafeModuloInt(_1543_i0, p1), (_1543_i0).Int())
			}
		}
		var _1544_v1 _dafny.Map
		_ = _1544_v1
		_1544_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14())
		var _1545_v2 D3
		_ = _1545_v2
		_1545_v2 = Companion_D3_.Create_DC10_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14()))
		_1544_v1 = func(_source15 D3) _dafny.Map {
			if _source15.Is_DC11() {
				var _1546___mcc_h0 _dafny.Int = _source15.Get_().(D3_DC11).Cf17
				_ = _1546___mcc_h0
				var _1547___mcc_h1 _dafny.Int = _source15.Get_().(D3_DC11).Cf18
				_ = _1547___mcc_h1
				var _1548___mcc_h2 bool = _source15.Get_().(D3_DC11).Cf19
				_ = _1548___mcc_h2
				var _1549___mcc_h3 _dafny.Int = _source15.Get_().(D3_DC11).Cf20
				_ = _1549___mcc_h3
				var _1550_cf20 _dafny.Int = _1549___mcc_h3
				_ = _1550_cf20
				var _1551_cf19 bool = _1548___mcc_h2
				_ = _1551_cf19
				var _1552_cf18 _dafny.Int = _1547___mcc_h1
				_ = _1552_cf18
				var _1553_cf17 _dafny.Int = _1546___mcc_h0
				_ = _1553_cf17
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _1551_cf19)
			} else {
				var _1554___mcc_h4 _dafny.Map = _source15.Get_().(D3_DC10).Cf16
				_ = _1554___mcc_h4
				var _1555_cf16 _dafny.Map = _1554___mcc_h4
				_ = _1555_cf16
				return (_1555_cf16).Merge(_1555_cf16)
			}
		}(_1545_v2)
		var _1556_v3 _dafny.Sequence
		_ = _1556_v3
		_1556_v3 = _dafny.SeqOf((_this).F18())
		var _1557_v4 _dafny.Sequence
		_ = _1557_v4
		_1557_v4 = _dafny.SeqOf(_this.F14())
		var _1558_v5 *C4
		_ = _1558_v5
		var _nw251 *C4 = New_C4_()
		_ = _nw251
		_nw251.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_1556_v3, _1556_v3), _this.F14(), _dafny.MultiSetOf((_1557_v4).Select((Companion_Default___.SafeIndex((_this).F18(), _dafny.IntOfUint32((_1557_v4).Cardinality()))).Uint32()).(bool), _this.F14()))
		_1558_v5 = _nw251
		var _1559_v6 _dafny.CodePoint
		_ = _1559_v6
		_1559_v6 = _dafny.CodePoint('t')
		var _1560_v7 _dafny.Sequence
		_ = _1560_v7
		_1560_v7 = _dafny.UnicodeSeqOfUtf8Bytes("h")
		var _1561_v8 _dafny.Sequence
		_ = _1561_v8
		_1561_v8 = _dafny.SeqOf(_1560_v7, _1560_v7)
		var _1562_v9 _dafny.Array
		_ = _1562_v9
		var _nwElement0_50 bool = (_1558_v5).F19()
		_ = _nwElement0_50
		var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(21))
		_ = _nw252
		(_nw252).ArraySet1(_nwElement0_50, 0)
		(_nw252).ArraySet1((_1558_v5).F19(), 1)
		(_nw252).ArraySet1((_dafny.IntOfInt64(720)).Cmp((_this).F18()) >= 0, 2)
		(_nw252).ArraySet1(_this.F14(), 3)
		(_nw252).ArraySet1(_this.F14(), 4)
		(_nw252).ArraySet1(((_this).F15()).IsSubsetOf((_this).F15()), 5)
		(_nw252).ArraySet1((_1558_v5).F19(), 6)
		(_nw252).ArraySet1(_this.F14(), 7)
		(_nw252).ArraySet1(_this.F14(), 8)
		(_nw252).ArraySet1(_this.F14(), 9)
		(_nw252).ArraySet1(true, 10)
		(_nw252).ArraySet1(_this.F14(), 11)
		(_nw252).ArraySet1(_dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(815))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg97 _dafny.Int) interface{} {
				return coer97(arg97)
			}
		}((func(_1563_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1564_i2 _dafny.Int) _dafny.CodePoint {
				return _1563_v6
			}
		})(_1559_v6))), _1559_v6), 12)
		(_nw252).ArraySet1((p1).Cmp(_dafny.IntOfUint32((_1560_v7).Cardinality())) <= 0, 13)
		(_nw252).ArraySet1(((_this).F18()).Cmp(_dafny.IntOfUint32((_1561_v8).Cardinality())) < 0, 14)
		(_nw252).ArraySet1(!(p0).Equals(p0), 15)
		(_nw252).ArraySet1(!((_1558_v5).F19()), 16)
		(_nw252).ArraySet1((p0).IsSubsetOf(p0), 17)
		(_nw252).ArraySet1((_1558_v5).F19(), 18)
		(_nw252).ArraySet1((_this).Fm2((_1558_v5).F19(), (_1558_v5).F19(), globalState), 19)
		(_nw252).ArraySet1(_this.F14(), 20)
		_1562_v9 = _nw252
		var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
		_ = _index276
		(_1562_v9).ArraySet1(_this.F14(), (_index276).Int())
		var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
		_ = _index277
		(_1541_v0).ArraySet1(_dafny.IntOfInt64(-934), (_index277).Int())
		var _1565_v10 _dafny.Array
		_ = _1565_v10
		var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
		_ = _nw253
		_1565_v10 = _nw253
		var _1566_v11 _dafny.Map
		_ = _1566_v11
		_1566_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm3(p1, p1, _this.F14(), p1, globalState), _this.F14())
		var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1565_v10), 0))
		_ = _index278
		(_1565_v10).ArraySet1(_1566_v11, (_index278).Int())
		var _1567_v12 _dafny.MultiSet
		_ = _1567_v12
		_1567_v12 = _dafny.MultiSetOf(_dafny.IntOfInt64(-837))
		var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
		_ = _index279
		var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
		_ = _index280
		var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1565_v10), 0))
		_ = _index281
		var _rhs253 bool = (_1567_v12).IsDisjointFrom((_dafny.MultiSetOf(p1)).Update((_this).F18(), Companion_Default___.Abs((_this).F18())))
		_ = _rhs253
		var _rhs254 _dafny.Int = (_dafny.Zero).Minus((_this).F18())
		_ = _rhs254
		var _rhs255 bool = (_1558_v5).F19()
		_ = _rhs255
		var _rhs256 _dafny.Map = _1566_v11
		_ = _rhs256
		var _lhs218 _dafny.Array = _1562_v9
		_ = _lhs218
		var _lhs219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
		_ = _lhs219
		var _lhs220 _dafny.Array = _1541_v0
		_ = _lhs220
		var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
		_ = _lhs221
		var _lhs222 *C5 = _this
		_ = _lhs222
		var _lhs223 _dafny.Array = _1565_v10
		_ = _lhs223
		var _lhs224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(245), _dafny.ArrayLen((_1565_v10), 0))
		_ = _lhs224
		(_lhs218).ArraySet1(_rhs253, (_lhs219).Int())
		(_lhs220).ArraySet1(_rhs254, (_lhs221).Int())
		_lhs222.F14_set_(_rhs255)
		(_lhs223).ArraySet1(_rhs256, (_lhs224).Int())
		var _1568_v13 _dafny.Sequence
		_ = _1568_v13
		_1568_v13 = _dafny.SeqOf(_1567_v12, _dafny.MultiSetOf((_this).F18()))
		var _1569_v14 _dafny.Set
		_ = _1569_v14
		_1569_v14 = _dafny.SetOf((_1568_v13).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1568_v13).Cardinality()))).Uint32()).(_dafny.MultiSet))
		var _hi5 _dafny.Int = Companion_Default___.SafeDivisionInt((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int), (_1569_v14).Cardinality())
		_ = _hi5
		for _1570_i3 := (_dafny.Zero).Minus((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int)); _1570_i3.Cmp(_hi5) < 0; _1570_i3 = _1570_i3.Plus(_dafny.One) {
			var _1571_v15 _dafny.Set
			_ = _1571_v15
			_1571_v15 = _dafny.SetOf(_1557_v4)
			var _1572_v17 _dafny.Set
			_ = _1572_v17
			_1572_v17 = _dafny.SetOf((_1558_v5).Fm2((_1558_v5).F19(), _this.F14(), globalState))
			var _1573_v18 _dafny.Map
			_ = _1573_v18
			_1573_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool), _1572_v17)
			var _1574_v19 _dafny.Map
			_ = _1574_v19
			_1574_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1556_v3, ((func() _dafny.Set {
				if (_1573_v18).Contains((_1558_v5).F19()) {
					return (_1573_v18).Get((_1558_v5).F19()).(_dafny.Set)
				}
				return _1572_v17
			})()).Cardinality())
			var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
			_ = _index282
			var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
			_ = _index283
			var _rhs257 _dafny.Sequence = _1560_v7
			_ = _rhs257
			var _rhs258 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(p1)).Times(_1570_i3), ((func() _dafny.Set {
				var _coll59 = _dafny.NewBuilder()
				_ = _coll59
				for _iter63 := _dafny.Iterate((_1571_v15).Elements()); ; {
					_compr_59, _ok63 := _iter63()
					if !_ok63 {
						break
					}
					var _1575_v16 _dafny.Sequence
					_1575_v16 = interface{}(_compr_59).(_dafny.Sequence)
					if (_1571_v15).Contains(_1575_v16) {
						_coll59.Add(_1575_v16)
					}
				}
				return _coll59.ToSet()
			}()).Intersection(_1571_v15)).Cardinality()))
			_ = _rhs258
			var _rhs259 _dafny.Int = (_this).F18()
			_ = _rhs259
			var _rhs260 _dafny.Sequence = (_this).Fm15((_1558_v5).F19(), ((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int)).Cmp(p1) == 0, (func() _dafny.Int {
				if (_1574_v19).Contains(_1556_v3) {
					return (_1574_v19).Get(_1556_v3).(_dafny.Int)
				}
				return p1
			})(), globalState)
			_ = _rhs260
			var _rhs261 bool = (_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool)
			_ = _rhs261
			var _lhs225 _dafny.Array = _1541_v0
			_ = _lhs225
			var _lhs226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
			_ = _lhs226
			var _lhs227 _dafny.Array = _1541_v0
			_ = _lhs227
			var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
			_ = _lhs228
			var _lhs229 *GlobalState = globalState
			_ = _lhs229
			_1560_v7 = _rhs257
			(_lhs225).ArraySet1(_rhs258, (_lhs226).Int())
			(_lhs227).ArraySet1(_rhs259, (_lhs228).Int())
			_1560_v7 = _rhs260
			_lhs229.F9 = _rhs261
			var _1576_v20 _dafny.Array
			_ = _1576_v20
			var _len0_45 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_45
			var _nw254 _dafny.Array
			_ = _nw254
			if _len0_45.Cmp(_dafny.Zero) == 0 {
				_nw254 = _dafny.NewArray(_len0_45)
			} else {
				var _init45 func(_dafny.Int) _dafny.CodePoint = func(_1577_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				}
				_ = _init45
				var _element0_45 = _init45(_dafny.Zero)
				_ = _element0_45
				_nw254 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
				(_nw254).ArraySet1CodePoint(_element0_45, 0)
				var _nativeLen0_45 = (_len0_45).Int()
				_ = _nativeLen0_45
				for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
					(_nw254).ArraySet1CodePoint(_init45(_dafny.IntOf(_i0_45)), _i0_45)
				}
			}
			_1576_v20 = _nw254
			var _1578_v21 _dafny.Map
			_ = _1578_v21
			_1578_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool), _dafny.IntOfInt64(95))
			var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1576_v20), 0))
			_ = _index284
			(_1576_v20).ArraySet1CodePoint((func() _dafny.CodePoint {
				if (_1558_v5).F19() {
					return Companion_Default___.Fm25(p1, p1, false, _1578_v21, globalState)
				}
				return _1559_v6
			})(), (_index284).Int())
			var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1576_v20), 0))
			_ = _index285
			(_1576_v20).ArraySet1CodePoint(_dafny.CodePoint('w'), (_index285).Int())
			var _1579_v22 _dafny.Map
			_ = _1579_v22
			_1579_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1578_v21, (_this).Fm2((_1558_v5).F19(), (_1558_v5).F19(), globalState))
			var _1580_v24 _dafny.Set
			_ = _1580_v24
			_1580_v24 = _dafny.SetOf(_1578_v21)
			if (func() _dafny.Set {
				var _coll60 = _dafny.NewBuilder()
				_ = _coll60
				for _iter64 := _dafny.Iterate((_1579_v22).Keys().Elements()); ; {
					_compr_60, _ok64 := _iter64()
					if !_ok64 {
						break
					}
					var _1581_v23 _dafny.Map
					_1581_v23 = interface{}(_compr_60).(_dafny.Map)
					if (_1579_v22).Contains(_1581_v23) {
						_coll60.Add(_1581_v23)
					}
				}
				return _coll60.ToSet()
			}()).IsSubsetOf(_1580_v24) {
				r1 = _1570_i3
				var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1576_v20), 0))
				_ = _index286
				(_1576_v20).ArraySet1CodePoint(_dafny.CodePoint('h'), (_index286).Int())
				var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
				_ = _index287
				var _rhs262 bool = (_1558_v5).F19()
				_ = _rhs262
				var _rhs263 bool = !(((_1558_v5).F19()) == (_this.F14()))
				_ = _rhs263
				var _rhs264 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(720))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg98 _dafny.Int) interface{} {
						return coer98(arg98)
					}
				}((func(_1582_p0 _dafny.Set) func(_dafny.Int) _dafny.Int {
					return func(_1583_i5 _dafny.Int) _dafny.Int {
						return (_1582_p0).Cardinality()
					}
				})(p0)))).Cardinality())
				_ = _rhs264
				var _lhs230 *GlobalState = globalState
				_ = _lhs230
				var _lhs231 *GlobalState = globalState
				_ = _lhs231
				var _lhs232 _dafny.Array = _1541_v0
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))
				_ = _lhs233
				_lhs230.F9 = _rhs262
				_lhs231.F9 = _rhs263
				(_lhs232).ArraySet1(_rhs264, (_lhs233).Int())
				var _nw255 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw255
				(globalState).F2 = _nw255
				var _1584_v25 _dafny.Array
				_ = _1584_v25
				var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
				_ = _nw256
				_1584_v25 = _nw256
				var _1585_v26 D0
				_ = _1585_v26
				_1585_v26 = Companion_D0_.Create_DC1_((_1558_v5).F19())
				var _rhs265 _dafny.Int = (_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int)
				_ = _rhs265
				var _rhs266 bool = (func() bool {
					if (_1558_v5).F19() {
						return !(_1578_v21).Contains((_1585_v26).Dtor_cf1())
					}
					return ((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool)) || ((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool))
				})()
				_ = _rhs266
				var _rhs267 _dafny.Int = ((_dafny.IntOfInt64(23)).Times(_1570_i3)).Minus((_dafny.IntOfInt64(63)).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("v")).Cardinality())))
				_ = _rhs267
				var _rhs268 _dafny.Array = _1584_v25
				_ = _rhs268
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				var _lhs235 *GlobalState = globalState
				_ = _lhs235
				_lhs234.F3 = _rhs265
				r2 = _rhs266
				_lhs235.F3 = _rhs267
				_1584_v25 = _rhs268
			} else {
				var _1586_v27 _dafny.Array
				_ = _1586_v27
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_46
				var _nw257 _dafny.Array
				_ = _nw257
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw257 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Sequence = (func(_1587_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1588_i6 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_1587_v3, _1587_v3)
						}
					})(_1556_v3)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw257 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw257).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw257).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1586_v27 = _nw257
				_1586_v27 = _1586_v27
				var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
				_ = _index288
				(_1562_v9).ArraySet1(!(!(!(true))), (_index288).Int())
				(_this).F14_set_(!((p1).Cmp((_this).F18()) <= 0))
				(globalState).F3 = p1
				_1541_v0 = _1541_v0
			}
			var _1589_v28 _dafny.Sequence
			_ = _1589_v28
			_1589_v28 = _dafny.SeqOf(_1562_v9)
			var _1590_v29 D1
			_ = _1590_v29
			_1590_v29 = Companion_D1_.Create_DC5_((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool), (_this).Fm2(false, _this.F14(), globalState))
			var _1591_v30 D11
			_ = _1591_v30
			_1591_v30 = Companion_D11_.Create_DC34_((_this).F15(), _1567_v12, _1562_v9, _1590_v29)
			var _1592_v31 _dafny.Array
			_ = _1592_v31
			var _nwElement0_51 _dafny.Array = _1562_v9
			_ = _nwElement0_51
			var _nw258 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(18))
			_ = _nw258
			(_nw258).ArraySet1(_nwElement0_51, 0)
			(_nw258).ArraySet1(_1562_v9, 1)
			(_nw258).ArraySet1(_1562_v9, 2)
			(_nw258).ArraySet1(_1562_v9, 3)
			(_nw258).ArraySet1(_1562_v9, 4)
			(_nw258).ArraySet1((_1589_v28).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1589_v28).Cardinality()))).Uint32()).(_dafny.Array), 5)
			(_nw258).ArraySet1(_1562_v9, 6)
			(_nw258).ArraySet1(_1562_v9, 7)
			(_nw258).ArraySet1(_1562_v9, 8)
			(_nw258).ArraySet1(_1562_v9, 9)
			(_nw258).ArraySet1(_1562_v9, 10)
			(_nw258).ArraySet1(_1562_v9, 11)
			(_nw258).ArraySet1(_1562_v9, 12)
			(_nw258).ArraySet1(_1562_v9, 13)
			(_nw258).ArraySet1(_1562_v9, 14)
			(_nw258).ArraySet1((_1591_v30).Dtor_cf60(), 15)
			(_nw258).ArraySet1(_1562_v9, 16)
			(_nw258).ArraySet1(_1562_v9, 17)
			_1592_v31 = _nw258
			var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1592_v31), 0))
			_ = _index289
			(_1592_v31).ArraySet1(_1562_v9, (_index289).Int())
			var _1593_v32 _dafny.Map
			_ = _1593_v32
			_1593_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1576_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(812), _dafny.ArrayLen((_1576_v20), 0))).Int()), _1557_v4)
			var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1592_v31), 0))
			_ = _index290
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
			_ = _index291
			var _rhs269 bool = (_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool)
			_ = _rhs269
			var _rhs270 _dafny.Array = _1562_v9
			_ = _rhs270
			var _rhs271 bool = (((_this).F15()).Difference((_this).F15())).IsDisjointFrom((_this).F15())
			_ = _rhs271
			var _rhs272 _dafny.Map = _1593_v32
			_ = _rhs272
			var _lhs236 _dafny.Array = _1592_v31
			_ = _lhs236
			var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1592_v31), 0))
			_ = _lhs237
			var _lhs238 _dafny.Array = _1562_v9
			_ = _lhs238
			var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))
			_ = _lhs239
			r2 = _rhs269
			(_lhs236).ArraySet1(_rhs270, (_lhs237).Int())
			(_lhs238).ArraySet1(_rhs271, (_lhs239).Int())
			_1593_v32 = _rhs272
		}
		var _1594_v33 _dafny.Map
		_ = _1594_v33
		_1594_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int), _1562_v9)
		var _1595_v34 _dafny.MultiSet
		_ = _1595_v34
		_1595_v34 = _dafny.MultiSetOf(_1557_v4, _1557_v4)
		var _1596_v35 _dafny.Sequence
		_ = _1596_v35
		_1596_v35 = _dafny.SeqOf(_1562_v9)
		var _1597_v36 _dafny.Map
		_ = _1597_v36
		_1597_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1559_v6, _1562_v9)
		var _1598_v37 _dafny.Map
		_ = _1598_v37
		_1598_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1562_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1562_v9), 0))).Int()).(bool), (_this).F18())
		var _1599_v38 _dafny.Array
		_ = _1599_v38
		var _nwElement0_52 _dafny.Array = _1562_v9
		_ = _nwElement0_52
		var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(24))
		_ = _nw259
		(_nw259).ArraySet1(_nwElement0_52, 0)
		(_nw259).ArraySet1(_1562_v9, 1)
		(_nw259).ArraySet1(_1562_v9, 2)
		(_nw259).ArraySet1(_1562_v9, 3)
		(_nw259).ArraySet1(_1562_v9, 4)
		(_nw259).ArraySet1((func() _dafny.Array {
			if (_1594_v33).Contains((_1558_v5).Fm17((_this).F18(), (_1558_v5).F19(), _1595_v34, (_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int), globalState)) {
				return (_1594_v33).Get((_1558_v5).Fm17((_this).F18(), (_1558_v5).F19(), _1595_v34, (_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int), globalState)).(_dafny.Array)
			}
			return _1562_v9
		})(), 5)
		(_nw259).ArraySet1(_1562_v9, 6)
		(_nw259).ArraySet1(_1562_v9, 7)
		(_nw259).ArraySet1((_1596_v35).Select((Companion_Default___.SafeIndex((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1596_v35).Cardinality()))).Uint32()).(_dafny.Array), 8)
		(_nw259).ArraySet1(_1562_v9, 9)
		(_nw259).ArraySet1(_1562_v9, 10)
		(_nw259).ArraySet1(_1562_v9, 11)
		(_nw259).ArraySet1(_1562_v9, 12)
		(_nw259).ArraySet1(_1562_v9, 13)
		(_nw259).ArraySet1(_1562_v9, 14)
		(_nw259).ArraySet1(_1562_v9, 15)
		(_nw259).ArraySet1(_1562_v9, 16)
		(_nw259).ArraySet1(_1562_v9, 17)
		(_nw259).ArraySet1((func() _dafny.Array {
			if (_1597_v36).Contains(Companion_Default___.Fm25(p1, p1, (_1558_v5).F19(), _1598_v37, globalState)) {
				return (_1597_v36).Get(Companion_Default___.Fm25(p1, p1, (_1558_v5).F19(), _1598_v37, globalState)).(_dafny.Array)
			}
			return _1562_v9
		})(), 18)
		(_nw259).ArraySet1(_1562_v9, 19)
		(_nw259).ArraySet1(_1562_v9, 20)
		(_nw259).ArraySet1(_1562_v9, 21)
		(_nw259).ArraySet1(_1562_v9, 22)
		(_nw259).ArraySet1(_1562_v9, 23)
		_1599_v38 = _nw259
		var _1600_v39 D4
		_ = _1600_v39
		_1600_v39 = Companion_D4_.Create_DC12_(_1562_v9)
		var _1601_v40 _dafny.Map
		_ = _1601_v40
		_1601_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1600_v39, _1599_v38)
		var _1602_v41 _dafny.Sequence
		_ = _1602_v41
		_1602_v41 = _dafny.SeqOf(_1599_v38)
		_1599_v38 = (func() _dafny.Array {
			if (_1601_v40).Contains(_1600_v39) {
				return (_1601_v40).Get(_1600_v39).(_dafny.Array)
			}
			return (_1602_v41).Select((Companion_Default___.SafeIndex((_1556_v3).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1556_v3).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1602_v41).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		r0 = _dafny.IntOfInt64(561)
		r1 = p1
		r2 = ((_dafny.Zero).Minus(p1)).Cmp((_1541_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(46), _dafny.ArrayLen((_1541_v0), 0))).Int()).(_dafny.Int)) > 0
		return r0, r1, r2
	}
}
func (_this *C5) F18() _dafny.Int {
	{
		return _this._f18
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f14 bool
	_f15 _dafny.MultiSet
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f14 = false
	_this._f15 = _dafny.EmptyMultiSet
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

func (_this *C6) F14() bool {
	return _this._f14
}
func (_this *C6) F14_set_(value bool) {
	_this._f14 = value
}
func (_this *C6) F15() _dafny.MultiSet {
	return _this._f15
}
func (_this *C6) Ctor__(f14 bool, f15 _dafny.MultiSet) {
	{
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C6) Fm2(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return !(func() _dafny.Set {
			var _coll61 = _dafny.NewBuilder()
			_ = _coll61
			for _iter65 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(490), _dafny.IntOfInt64(573))); ; {
				_compr_61, _ok65 := _iter65()
				if !_ok65 {
					break
				}
				var _1603_v0 _dafny.Int
				_1603_v0 = interface{}(_compr_61).(_dafny.Int)
				if ((_dafny.IntOfInt64(490)).Cmp(_1603_v0) <= 0) && ((_1603_v0).Cmp(_dafny.IntOfInt64(573)) < 0) {
					_coll61.Add(Companion_Default___.SafeDivisionInt(_1603_v0, _dafny.IntOfInt64(621)))
				}
			}
			return _coll61.ToSet()
		}()).Equals((func() _dafny.Set {
			var _coll62 = _dafny.NewBuilder()
			_ = _coll62
			for _iter66 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(-224))).Elements()); ; {
				_compr_62, _ok66 := _iter66()
				if !_ok66 {
					break
				}
				var _1604_v1 _dafny.Int
				_1604_v1 = interface{}(_compr_62).(_dafny.Int)
				if (_dafny.SetOf(_dafny.IntOfInt64(-224))).Contains(_1604_v1) {
					_coll62.Add((_1604_v1).Times(_dafny.IntOfInt64(544)))
				}
			}
			return _coll62.ToSet()
		}()).Difference(func() _dafny.Set {
			var _coll63 = _dafny.NewBuilder()
			_ = _coll63
			for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(613), _dafny.IntOfInt64(145))); ; {
				_compr_63, _ok67 := _iter67()
				if !_ok67 {
					break
				}
				var _1605_v2 _dafny.Int
				_1605_v2 = interface{}(_compr_63).(_dafny.Int)
				if ((_dafny.IntOfInt64(613)).Cmp(_1605_v2) <= 0) && ((_1605_v2).Cmp(_dafny.IntOfInt64(145)) < 0) {
					_coll63.Add((_1605_v2).Minus(_dafny.IntOfInt64(943)))
				}
			}
			return _coll63.ToSet()
		}()))
	}
}
func (_this *C6) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.UnicodeSeqOfUtf8Bytes("evj")
	}
}
func (_this *C6) Fm11(p0 _dafny.Map, p1 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this.F14()) && ((func() _dafny.Set {
			var _coll64 = _dafny.NewBuilder()
			_ = _coll64
			for _iter68 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(857), (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bv"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-494))).Uint32(), func(coer99 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}(func(_1606_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			})))).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfInt64(339))).Elements()); ; {
				_compr_64, _ok68 := _iter68()
				if !_ok68 {
					break
				}
				var _1607_v0 _dafny.Int
				_1607_v0 = interface{}(_compr_64).(_dafny.Int)
				if (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(857), (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("bv"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-494))).Uint32(), func(coer100 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}(func(_1606_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				})))).Cardinality())).Cardinality())).Cardinality(), _dafny.IntOfInt64(339))).Contains(_1607_v0) {
					_coll64.Add((_1607_v0).Minus(_dafny.IntOfInt64(177)))
				}
			}
			return _coll64.ToSet()
		}()).IsSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(864), _dafny.IntOfInt64(476), _dafny.IntOfInt64(180), (Companion_D1_.Create_DC3_(_dafny.IntOfInt64(646))).Dtor_cf3(), _dafny.IntOfInt64(305))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("njb")).Cardinality()))))
	}
}
func (_this *C6) Fm12(p0 bool, p1 _dafny.Map, globalState *GlobalState) bool {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.SeqOf(_this.F14())).Cardinality())), (_this).F15())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC3_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aksalss")).Cardinality())), (_this).F15()))).Equals(func() _dafny.Map {
			var _coll65 = _dafny.NewMapBuilder()
			_ = _coll65
			for _iter69 := _dafny.Iterate((func() _dafny.Map {
				var _coll66 = _dafny.NewMapBuilder()
				_ = _coll66
				for _iter70 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-509)))).Elements()); ; {
					_compr_66, _ok70 := _iter70()
					if !_ok70 {
						break
					}
					var _1608_v1 D1
					_1608_v1 = interface{}(_compr_66).(D1)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-509))), _1608_v1) {
						_coll66.Add(_1608_v1, _dafny.IntOfInt64(121))
					}
				}
				return _coll66.ToMap()
			}()).Keys().Elements()); ; {
				_compr_65, _ok69 := _iter69()
				if !_ok69 {
					break
				}
				var _1609_v0 D1
				_1609_v0 = interface{}(_compr_65).(D1)
				if (func() _dafny.Map {
					var _coll67 = _dafny.NewMapBuilder()
					_ = _coll67
					for _iter71 := _dafny.Iterate((_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-509)))).Elements()); ; {
						_compr_67, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _1608_v1 D1
						_1608_v1 = interface{}(_compr_67).(D1)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-509))), _1608_v1) {
							_coll67.Add(_1608_v1, _dafny.IntOfInt64(121))
						}
					}
					return _coll67.ToMap()
				}()).Contains(_1609_v0) {
					_coll65.Add(_1609_v0, (_this).F15())
				}
			}
			return _coll65.ToMap()
		}())
	}
}
func (_this *C6) M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _1610_v0 _dafny.Int
		_ = _1610_v0
		_1610_v0 = _dafny.IntOfInt64(-804)
		var _hi6 _dafny.Int = _1610_v0
		_ = _hi6
		for _1611_i0 := _1610_v0; _1611_i0.Cmp(_hi6) < 0; _1611_i0 = _1611_i0.Plus(_dafny.One) {
			var _1612_v1 _dafny.Sequence
			_ = _1612_v1
			_1612_v1 = _dafny.UnicodeSeqOfUtf8Bytes("li")
			var _1613_v2 _dafny.Map
			_ = _1613_v2
			_1613_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg101 _dafny.Int) interface{} {
					return coer101(arg101)
				}
			}(func(_1614_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("pf")
			})), (Companion_Default___.SafeIndex(_1610_v0, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(878))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg102 _dafny.Int) interface{} {
					return coer102(arg102)
				}
			}(func(_1615_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("pf")
			}))).Cardinality()))).Uint32(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(343))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}(func(_1616_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})))).Cardinality()), _1612_v1)
			_1613_v2 = (_1613_v2).Update(_1610_v0, (func() _dafny.Sequence {
				if _this.F14() {
					return _1612_v1
				}
				return _1612_v1
			})())
			var _1617_v3 D0
			_ = _1617_v3
			_1617_v3 = Companion_D0_.Create_DC0_(_this.F14())
			if ((func() D0 {
				if (_1617_v3).Dtor_cf0() {
					return _1617_v3
				}
				return _1617_v3
			})()).Dtor_cf0() {
				var _1618_v4 _dafny.Sequence
				_ = _1618_v4
				_1618_v4 = _dafny.SeqOf(_dafny.IntOfUint32((_1612_v1).Cardinality()))
				var _1619_v5 _dafny.MultiSet
				_ = _1619_v5
				_1619_v5 = _dafny.MultiSetOf(_1611_i0)
				var _1620_v6 _dafny.MultiSet
				_ = _1620_v6
				_1620_v6 = _dafny.MultiSetOf((func() _dafny.Sequence {
					if _this.F14() {
						return _dafny.SeqOf(Companion_Default___.Fm0(_this.F14(), _1611_i0, _dafny.IntOfUint32((_1618_v4).Cardinality()), _1619_v5, globalState))
					}
					return _dafny.SeqOf(_1611_i0)
				})(), _1618_v4, _dafny.Companion_Sequence_.Update(_1618_v4, (Companion_Default___.SafeIndex(_1610_v0, _dafny.IntOfUint32((_1618_v4).Cardinality()))).Uint32(), _1610_v0))
				_1620_v6 = Companion_Default___.Fm13((_this).F15(), (_dafny.IntOfInt64(488)).Times(_1611_i0), globalState)
				var _1621_v7 _dafny.Array
				_ = _1621_v7
				var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
				_ = _nw260
				_1621_v7 = _nw260
				var _1622_v8 _dafny.Array
				_ = _1622_v8
				var _nw261 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw261
				_1622_v8 = _nw261
				var _1623_v9 _dafny.Sequence
				_ = _1623_v9
				_1623_v9 = _dafny.SeqOf(_1622_v8, _1622_v8, _1622_v8)
				var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1621_v7), 0))
				_ = _index292
				(_1621_v7).ArraySet1(_1623_v9, (_index292).Int())
				var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_1621_v7), 0))
				_ = _index293
				(_1621_v7).ArraySet1(_dafny.SeqOf(_1622_v8), (_index293).Int())
				var _1624_v10 _dafny.Map
				_ = _1624_v10
				_1624_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false)).Equals((_this).F15()), _1618_v4)
				_1624_v10 = (_1624_v10).Update(_this.F14(), _1618_v4)
				var _1625_v11 _dafny.Int
				_ = _1625_v11
				var _1626_v12 _dafny.Int
				_ = _1626_v12
				var _1627_v13 _dafny.MultiSet
				_ = _1627_v13
				var _1628_v14 bool
				_ = _1628_v14
				var _out47 _dafny.Int
				_ = _out47
				var _out48 _dafny.Int
				_ = _out48
				var _out49 _dafny.MultiSet
				_ = _out49
				var _out50 bool
				_ = _out50
				_out47, _out48, _out49, _out50 = (_this).M5(globalState)
				_1625_v11 = _out47
				_1626_v12 = _out48
				_1627_v13 = _out49
				_1628_v14 = _out50
				var _1629_v15 _dafny.Sequence
				_ = _1629_v15
				_1629_v15 = _dafny.SeqOf(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(512))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg104 _dafny.Int) interface{} {
						return coer104(arg104)
					}
				}(func(_1630_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('g')
				})), _1612_v1), true, _1628_v14, _1628_v14, _1628_v14)
				var _rhs273 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1629_v15, _1629_v15)
				_ = _rhs273
				var _rhs274 _dafny.Array = _1622_v8
				_ = _rhs274
				var _rhs275 bool = _this.F14()
				_ = _rhs275
				var _lhs240 *GlobalState = globalState
				_ = _lhs240
				_1629_v15 = _rhs273
				_1622_v8 = _rhs274
				_lhs240.F9 = _rhs275
			} else {
				var _1631_v16 _dafny.Set
				_ = _1631_v16
				_1631_v16 = _dafny.SetOf(_this.F14())
				var _1632_v17 _dafny.Sequence
				_ = _1632_v17
				_1632_v17 = _dafny.SeqOf(_1631_v16)
				var _1633_v18 _dafny.MultiSet
				_ = _1633_v18
				_1633_v18 = _dafny.MultiSetOf(_1611_i0)
				var _1634_v19 _dafny.MultiSet
				_ = _1634_v19
				_1634_v19 = _dafny.MultiSetOf(Companion_Default___.Fm0(_this.F14(), (_dafny.Zero).Minus((_1633_v18).Cardinality()), _1611_i0, _1633_v18, globalState), _dafny.IntOfInt64(-334), _1611_i0, (_dafny.Zero).Minus(_1610_v0), _1611_i0)
				var _1635_v20 _dafny.Map
				_ = _1635_v20
				_1635_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_1632_v17, _dafny.SeqOf(_1631_v16, _1631_v16)), (_1634_v19).IsProperSubsetOf(_1633_v18))
				var _1636_v21 _dafny.Map
				_ = _1636_v21
				_1636_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm14(globalState), _1610_v0)
				var _1637_v22 _dafny.Sequence
				_ = _1637_v22
				_1637_v22 = _dafny.SeqOf(_1636_v21)
				_1635_v20 = (_1635_v20).Update(!((_this).Fm12(_this.F14(), (_1637_v22).Select((Companion_Default___.SafeIndex(_1610_v0, _dafny.IntOfUint32((_1637_v22).Cardinality()))).Uint32()).(_dafny.Map), globalState)), _this.F14())
				var _1638_v23 _dafny.Map
				_ = _1638_v23
				_1638_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1612_v1, Companion_D0_.Create_DC0_(_this.F14()))
				var _1639_v24 _dafny.Sequence
				_ = _1639_v24
				_1639_v24 = _dafny.SeqOf(_this.F14(), (_this).Fm12(_this.F14(), Companion_Default___.Fm48(_1611_i0, _1610_v0, _this.F14(), globalState), globalState), _this.F14())
				var _1640_v25 T0
				_ = _1640_v25
				var _nw262 *C5 = New_C5_()
				_ = _nw262
				_nw262.Ctor__(_1610_v0, true, _dafny.MultiSetFromSeq(_1639_v24))
				_1640_v25 = _nw262
				var _1641_v26 _dafny.Map
				_ = _1641_v26
				_1641_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1640_v25, _1612_v1)
				var _1642_v27 _dafny.Array
				_ = _1642_v27
				var _nw263 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw263
				_1642_v27 = _nw263
				_1638_v23 = (_1638_v23).Update(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_1641_v26).Contains(_1640_v25) {
						return (_1641_v26).Get(_1640_v25).(_dafny.Sequence)
					}
					return (Companion_D4_.Create_DC13_(_1612_v1, _1610_v0, _this.F14(), _1642_v27)).Dtor_cf22()
				})(), _1612_v1), _1617_v3)
				var _1643_v28 _dafny.Map
				_ = _1643_v28
				_1643_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("jvdqndys"), _1611_i0)
				r3 = !(_1643_v28).Contains(_dafny.Companion_Sequence_.Concatenate(_1612_v1, _1612_v1))
				r0 = _1610_v0
				var _1644_v29 _dafny.Sequence
				_ = _1644_v29
				_1644_v29 = _dafny.SeqOf(_1635_v20)
				var _1645_v30 _dafny.MultiSet
				_ = _1645_v30
				_1645_v30 = _dafny.MultiSetOf(_1635_v20)
				var _1646_v31 _dafny.Sequence
				_ = _1646_v31
				_1646_v31 = _dafny.SeqOf(_1645_v30)
				var _1647_v32 _dafny.CodePoint
				_ = _1647_v32
				_1647_v32 = _dafny.CodePoint('b')
				var _1648_v33 _dafny.Map
				_ = _1648_v33
				_1648_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1611_i0, _1647_v32)
				var _1649_v34 _dafny.Array
				_ = _1649_v34
				var _nwElement0_53 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_1644_v29)).Union(_1645_v30)
				_ = _nwElement0_53
				var _nw264 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(19))
				_ = _nw264
				(_nw264).ArraySet1(_nwElement0_53, 0)
				(_nw264).ArraySet1((_1645_v30).Intersection(_1645_v30), 1)
				(_nw264).ArraySet1(_1645_v30, 2)
				(_nw264).ArraySet1(_1645_v30, 3)
				(_nw264).ArraySet1(_1645_v30, 4)
				(_nw264).ArraySet1((_1646_v31).Select((Companion_Default___.SafeIndex((_1648_v33).Cardinality(), _dafny.IntOfUint32((_1646_v31).Cardinality()))).Uint32()).(_dafny.MultiSet), 5)
				(_nw264).ArraySet1(_1645_v30, 6)
				(_nw264).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1644_v29, (Companion_Default___.SafeIndex((_1631_v16).Cardinality(), _dafny.IntOfUint32((_1644_v29).Cardinality()))).Uint32(), _1635_v20)), 7)
				(_nw264).ArraySet1(_1645_v30, 8)
				(_nw264).ArraySet1((_1645_v30).Difference(_1645_v30), 9)
				(_nw264).ArraySet1(_1645_v30, 10)
				(_nw264).ArraySet1(_1645_v30, 11)
				(_nw264).ArraySet1(_1645_v30, 12)
				(_nw264).ArraySet1(_1645_v30, 13)
				(_nw264).ArraySet1(_1645_v30, 14)
				(_nw264).ArraySet1(_1645_v30, 15)
				(_nw264).ArraySet1(_1645_v30, 16)
				(_nw264).ArraySet1(_dafny.MultiSetOf(_1635_v20, _1635_v20), 17)
				(_nw264).ArraySet1((_1645_v30).Update(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1640_v25.F14()), Companion_Default___.Abs(_dafny.IntOfUint32((_1639_v24).Cardinality()))), 18)
				_1649_v34 = _nw264
				var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1649_v34), 0))
				_ = _index294
				(_1649_v34).ArraySet1((_1645_v30).Union(_1645_v30), (_index294).Int())
				var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_1649_v34), 0))
				_ = _index295
				(_1649_v34).ArraySet1(_dafny.MultiSetFromSeq(_1644_v29), (_index295).Int())
			}
			if _this.F14() {
				var _1650_v35 _dafny.Array
				_ = _1650_v35
				var _nw265 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(27))
				_ = _nw265
				_1650_v35 = _nw265
				_1650_v35 = _1650_v35
				(globalState).F9 = _this.F14()
				var _1651_v36 _dafny.Sequence
				_ = _1651_v36
				_1651_v36 = _dafny.SeqOf(_1610_v0)
				(globalState).F3 = Companion_Default___.SafeModuloInt(_1611_i0, _dafny.IntOfUint32((_1651_v36).Cardinality()))
				var _1652_v37 _dafny.Set
				_ = _1652_v37
				_1652_v37 = _dafny.SetOf((_this).F15())
				var _1653_v38 _dafny.Map
				_ = _1653_v38
				_1653_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D11_.Create_DC33_(_1652_v37), _1610_v0)
				var _1654_v39 D11
				_ = _1654_v39
				_1654_v39 = Companion_D11_.Create_DC33_(_1652_v37)
				_1653_v38 = (_1653_v38).Update(_1654_v39, ((_dafny.Zero).Minus(_1611_i0)).Times(_1610_v0))
				var _1655_v40 *C1
				_ = _1655_v40
				var _nw266 *C1 = New_C1_()
				_ = _nw266
				_nw266.Ctor__()
				_1655_v40 = _nw266
			} else {
				var _1656_v41 _dafny.Array
				_ = _1656_v41
				var _nwElement0_54 _dafny.Int = _1611_i0
				_ = _nwElement0_54
				var _nw267 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(2))
				_ = _nw267
				(_nw267).ArraySet1(_nwElement0_54, 0)
				(_nw267).ArraySet1(_1610_v0, 1)
				_1656_v41 = _nw267
				var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1656_v41), 0))
				_ = _index296
				(_1656_v41).ArraySet1(_1610_v0, (_index296).Int())
				var _1657_v42 _dafny.Array
				_ = _1657_v42
				var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw268
				_1657_v42 = _nw268
				var _1658_v43 _dafny.CodePoint
				_ = _1658_v43
				_1658_v43 = _dafny.CodePoint('t')
				var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1657_v42), 0))
				_ = _index297
				(_1657_v42).ArraySet1CodePoint(_1658_v43, (_index297).Int())
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1656_v41), 0))
				_ = _index298
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1657_v42), 0))
				_ = _index299
				var _rhs276 _dafny.Int = (_1610_v0).Plus(_1611_i0)
				_ = _rhs276
				var _rhs277 _dafny.Int = _1611_i0
				_ = _rhs277
				var _rhs278 _dafny.CodePoint = _1658_v43
				_ = _rhs278
				var _rhs279 _dafny.Int = (_dafny.Zero).Minus(_1611_i0)
				_ = _rhs279
				var _lhs241 *GlobalState = globalState
				_ = _lhs241
				var _lhs242 _dafny.Array = _1656_v41
				_ = _lhs242
				var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1656_v41), 0))
				_ = _lhs243
				var _lhs244 _dafny.Array = _1657_v42
				_ = _lhs244
				var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1657_v42), 0))
				_ = _lhs245
				_lhs241.F3 = _rhs276
				(_lhs242).ArraySet1(_rhs277, (_lhs243).Int())
				(_lhs244).ArraySet1CodePoint(_rhs278, (_lhs245).Int())
				_1610_v0 = _rhs279
				var _1659_v45 _dafny.Array
				_ = _1659_v45
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_47
				var _nw269 _dafny.Array
				_ = _nw269
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw269 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) bool = (func(_1660_v0 _dafny.Int, _1661_i0 _dafny.Int) func(_dafny.Int) bool {
						return func(_1662_i4 _dafny.Int) bool {
							return (func() _dafny.Set {
								var _coll68 = _dafny.NewBuilder()
								_ = _coll68
								for _iter72 := _dafny.Iterate((_dafny.SetOf(_1661_i0)).Elements()); ; {
									_compr_68, _ok72 := _iter72()
									if !_ok72 {
										break
									}
									var _1663_v44 _dafny.Int
									_1663_v44 = interface{}(_compr_68).(_dafny.Int)
									if (_dafny.SetOf(_1661_i0)).Contains(_1663_v44) {
										_coll68.Add((_1663_v44).Minus(_dafny.IntOfInt64(274)))
									}
								}
								return _coll68.ToSet()
							}()).IsDisjointFrom(_dafny.SetOf(_1660_v0))
						}
					})(_1610_v0, _1611_i0)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw269 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw269).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw269).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1659_v45 = _nw269
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_1659_v45), 0))
				_ = _index300
				(_1659_v45).ArraySet1(true, (_index300).Int())
				var _1664_v46 _dafny.Sequence
				_ = _1664_v46
				_1664_v46 = _dafny.SeqOf(_this.F14(), (_this.F14()) || (true), _this.F14())
				var _1665_v47 _dafny.Set
				_ = _1665_v47
				_1665_v47 = _dafny.SetOf(_this.F14())
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_1659_v45), 0))
				_ = _index301
				var _rhs280 bool = _this.F14()
				_ = _rhs280
				var _rhs281 bool = (_1664_v46).Select((Companion_Default___.SafeIndex(((func() _dafny.Set {
					if _this.F14() {
						return _dafny.SetOf(_this.F14(), true)
					}
					return _1665_v47
				})()).Cardinality(), _dafny.IntOfUint32((_1664_v46).Cardinality()))).Uint32()).(bool)
				_ = _rhs281
				var _rhs282 bool = _this.F14()
				_ = _rhs282
				var _lhs246 _dafny.Array = _1659_v45
				_ = _lhs246
				var _lhs247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.ArrayLen((_1659_v45), 0))
				_ = _lhs247
				r3 = _rhs280
				r1 = _rhs281
				(_lhs246).ArraySet1(_rhs282, (_lhs247).Int())
				var _1666_v48 _dafny.Map
				_ = _1666_v48
				_1666_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1657_v42).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1657_v42), 0))).Int()), _dafny.IntOfInt64(891))
				var _1667_v49 _dafny.Sequence
				_ = _1667_v49
				_1667_v49 = _dafny.SeqOf(_1610_v0, _1611_i0)
				var _1668_v50 _dafny.Set
				_ = _1668_v50
				_1668_v50 = _dafny.SetOf(_1667_v49, _1667_v49, _dafny.Companion_Sequence_.Update(_1667_v49, (Companion_Default___.SafeIndex((_1656_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1656_v41), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1667_v49).Cardinality()))).Uint32(), (_1656_v41).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_1656_v41), 0))).Int()).(_dafny.Int)), _1667_v49)
				(globalState).F9 = !((func() bool {
					if true {
						return Companion_Default___.Fm23((_dafny.Zero).Minus(_1610_v0), _this.F14(), _1666_v48, _1668_v50, globalState)
					}
					return false
				})())
				var _1669_v51 _dafny.Int
				_ = _1669_v51
				var _1670_v52 _dafny.Int
				_ = _1670_v52
				var _1671_v53 _dafny.MultiSet
				_ = _1671_v53
				var _1672_v54 bool
				_ = _1672_v54
				var _out51 _dafny.Int
				_ = _out51
				var _out52 _dafny.Int
				_ = _out52
				var _out53 _dafny.MultiSet
				_ = _out53
				var _out54 bool
				_ = _out54
				_out51, _out52, _out53, _out54 = (_this).M5(globalState)
				_1669_v51 = _out51
				_1670_v52 = _out52
				_1671_v53 = _out53
				_1672_v54 = _out54
				_1612_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1658_v43), _1612_v1)
			}
			var _1673_v55 _dafny.Array
			_ = _1673_v55
			var _nw270 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
			_ = _nw270
			_1673_v55 = _nw270
			var _1674_v56 _dafny.Map
			_ = _1674_v56
			_1674_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1610_v0, _1610_v0)
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1673_v55), 0))
			_ = _index302
			(_1673_v55).ArraySet1(((_1674_v56).Merge((func() _dafny.Map {
				var _coll69 = _dafny.NewMapBuilder()
				_ = _coll69
				for _iter73 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(964), _dafny.IntOfInt64(-867))); ; {
					_compr_69, _ok73 := _iter73()
					if !_ok73 {
						break
					}
					var _1675_v57 _dafny.Int
					_1675_v57 = interface{}(_compr_69).(_dafny.Int)
					if ((_dafny.IntOfInt64(964)).Cmp(_1675_v57) <= 0) && ((_1675_v57).Cmp(_dafny.IntOfInt64(-867)) < 0) {
						_coll69.Add((_1675_v57).Minus(_dafny.IntOfUint32((_1612_v1).Cardinality())), _1611_i0)
					}
				}
				return _coll69.ToMap()
			}()).Update((_dafny.Zero).Minus(_1610_v0), _1611_i0))).Cardinality(), (_index302).Int())
			var _1676_v58 _dafny.Map
			_ = _1676_v58
			_1676_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1612_v1, (_dafny.Zero).Minus(_1610_v0))
			var _1677_v59 _dafny.MultiSet
			_ = _1677_v59
			_1677_v59 = _dafny.MultiSetOf(_dafny.IntOfUint32((_1612_v1).Cardinality()))
			var _1678_v60 _dafny.Sequence
			_ = _1678_v60
			_1678_v60 = _dafny.SeqOf(_1610_v0)
			var _1679_v61 _dafny.Set
			_ = _1679_v61
			_1679_v61 = _dafny.SetOf(_dafny.SeqOf(Companion_Default___.Fm0(_this.F14(), _1611_i0, _1611_i0, _1677_v59, globalState)), _1678_v60, _1678_v60)
			var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(518), _dafny.ArrayLen((_1673_v55), 0))
			_ = _index303
			(_1673_v55).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1676_v58).Contains(_1612_v1) {
					return (_1676_v58).Get(_1612_v1).(_dafny.Int)
				}
				return ((Companion_D13_.Create_DC41_(_1679_v61)).Dtor_cf67()).Cardinality()
			})(), _1611_i0), (_index303).Int())
		}
		var _1680_v62 _dafny.Array
		_ = _1680_v62
		var _nw271 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(19))
		_ = _nw271
		_1680_v62 = _nw271
		var _1681_v63 _dafny.Set
		_ = _1681_v63
		_1681_v63 = _dafny.SetOf(_1610_v0)
		var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1680_v62), 0))
		_ = _index304
		(_1680_v62).ArraySet1(_1681_v63, (_index304).Int())
		var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(906), _dafny.ArrayLen((_1680_v62), 0))
		_ = _index305
		(_1680_v62).ArraySet1(_1681_v63, (_index305).Int())
		var _1682_v64 _dafny.CodePoint
		_ = _1682_v64
		_1682_v64 = _dafny.CodePoint('h')
		var _1683_v65 _dafny.MultiSet
		_ = _1683_v65
		_1683_v65 = _dafny.MultiSetOf(_1682_v64)
		_1683_v65 = (_1683_v65).Intersection((_1683_v65).Difference(_1683_v65))
		var _1684_v66 _dafny.Array
		_ = _1684_v66
		var _nw272 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw272
		_1684_v66 = _nw272
		var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))
		_ = _index306
		(_1684_v66).ArraySet1(!(_this.F14()), (_index306).Int())
		var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))
		_ = _index307
		(_1684_v66).ArraySet1(_this.F14(), (_index307).Int())
		for _iter74 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1684_v66), 0))); ; {
			_guard_loop_4, _ok74 := _iter74()
			if !_ok74 {
				break
			}
			var _1685_i5 _dafny.Int
			_1685_i5 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1685_i5).Sign() != -1) && ((_1685_i5).Cmp(_dafny.ArrayLen((_1684_v66), 0)) < 0)) {
				(_1684_v66).ArraySet1(((_1610_v0).Plus(_1610_v0)).Cmp(_1610_v0) != 0, (_1685_i5).Int())
			}
		}
		var _1686_v67 _dafny.Int
		_ = _1686_v67
		var _1687_v68 _dafny.Int
		_ = _1687_v68
		var _1688_v69 _dafny.MultiSet
		_ = _1688_v69
		var _1689_v70 bool
		_ = _1689_v70
		var _out55 _dafny.Int
		_ = _out55
		var _out56 _dafny.Int
		_ = _out56
		var _out57 _dafny.MultiSet
		_ = _out57
		var _out58 bool
		_ = _out58
		_out55, _out56, _out57, _out58 = (_this).M5(globalState)
		_1686_v67 = _out55
		_1687_v68 = _out56
		_1688_v69 = _out57
		_1689_v70 = _out58
		r0 = ((_1610_v0).Times(_1687_v68)).Plus((func() _dafny.Map {
			var _coll70 = _dafny.NewMapBuilder()
			_ = _coll70
			for _iter75 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer105 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg105 _dafny.Int) interface{} {
					return coer105(arg105)
				}
			}((func(_1690_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1691_i6 _dafny.Int) _dafny.Int {
					return _1690_v0
				}
			})(_1610_v0)))).Elements()); ; {
				_compr_70, _ok75 := _iter75()
				if !_ok75 {
					break
				}
				var _1692_v71 _dafny.Int
				_1692_v71 = interface{}(_compr_70).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(363))).Uint32(), func(coer106 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg106 _dafny.Int) interface{} {
						return coer106(arg106)
					}
				}((func(_1693_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1691_i6 _dafny.Int) _dafny.Int {
						return _1693_v0
					}
				})(_1610_v0))), _1692_v71) {
					_coll70.Add(Companion_Default___.SafeDivisionInt(_1692_v71, _1687_v68), _1610_v0)
				}
			}
			return _coll70.ToMap()
		}()).Cardinality())
		var _1694_v72 D1
		_ = _1694_v72
		_1694_v72 = Companion_D1_.Create_DC3_((_dafny.Zero).Minus(_1610_v0))
		var _1695_v73 D1
		_ = _1695_v73
		_1695_v73 = Companion_D1_.Create_DC6_(_1694_v72)
		var _pat_let_tv41 = _1682_v64
		_ = _pat_let_tv41
		var _pat_let_tv42 = _1682_v64
		_ = _pat_let_tv42
		r1 = func(_source16 D1) bool {
			if _source16.Is_DC4() {
				var _1696___mcc_h0 _dafny.Int = _source16.Get_().(D1_DC4).Cf4
				_ = _1696___mcc_h0
				var _1697___mcc_h1 _dafny.Int = _source16.Get_().(D1_DC4).Cf5
				_ = _1697___mcc_h1
				var _1698_cf5 _dafny.Int = _1697___mcc_h1
				_ = _1698_cf5
				var _1699_cf4 _dafny.Int = _1696___mcc_h0
				_ = _1699_cf4
				return (Companion_D13_.Create_DC42_(_dafny.CodePoint('r'), _this.F14())).Dtor_cf69()
			} else if _source16.Is_DC5() {
				var _1700___mcc_h2 bool = _source16.Get_().(D1_DC5).Cf6
				_ = _1700___mcc_h2
				var _1701___mcc_h3 bool = _source16.Get_().(D1_DC5).Cf7
				_ = _1701___mcc_h3
				var _1702_cf7 bool = _1701___mcc_h3
				_ = _1702_cf7
				var _1703_cf6 bool = _1700___mcc_h2
				_ = _1703_cf6
				return _this.F14()
			} else if _source16.Is_DC3() {
				var _1704___mcc_h4 _dafny.Int = _source16.Get_().(D1_DC3).Cf3
				_ = _1704___mcc_h4
				var _1705_cf3 _dafny.Int = _1704___mcc_h4
				_ = _1705_cf3
				return _this.F14()
			} else {
				var _1706___mcc_h5 D1 = _source16.Get_().(D1_DC6).Cf8
				_ = _1706___mcc_h5
				var _1707_cf8 D1 = _1706___mcc_h5
				_ = _1707_cf8
				return (_dafny.SetOf(_pat_let_tv41)).IsSubsetOf(_dafny.SetOf(_pat_let_tv42))
			}
		}((func() D1 {
			if true {
				return _1695_v73
			}
			return _1695_v73
		})())
		var _1708_v74 _dafny.Map
		_ = _1708_v74
		_1708_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1684_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))).Int()).(bool), _1686_v67)
		var _1709_v75 _dafny.Map
		_ = _1709_v75
		_1709_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1708_v74).Contains((_1684_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))).Int()).(bool)) {
				return (_1708_v74).Get((_1684_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))).Int()).(bool)).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(_1687_v68)
		})(), (_1684_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(542), _dafny.ArrayLen((_1684_v66), 0))).Int()).(bool)), _1610_v0)
		r2 = _1709_v75
		r3 = _1689_v70
		return r0, r1, r2, r3
	}
}
func (_this *C6) M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1710_v0 _dafny.Array
		_ = _1710_v0
		var _len0_48 _dafny.Int = _dafny.IntOfInt64(5)
		_ = _len0_48
		var _nw273 _dafny.Array
		_ = _nw273
		if _len0_48.Cmp(_dafny.Zero) == 0 {
			_nw273 = _dafny.NewArray(_len0_48)
		} else {
			var _init48 func(_dafny.Int) bool = func(_1711_i0 _dafny.Int) bool {
				return _this.F14()
			}
			_ = _init48
			var _element0_48 = _init48(_dafny.Zero)
			_ = _element0_48
			_nw273 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
			(_nw273).ArraySet1(_element0_48, 0)
			var _nativeLen0_48 = (_len0_48).Int()
			_ = _nativeLen0_48
			for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
				(_nw273).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
			}
		}
		_1710_v0 = _nw273
		var _1712_v1 _dafny.Sequence
		_ = _1712_v1
		_1712_v1 = _dafny.SeqOf(_dafny.CodePoint('i'))
		var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
		_ = _index308
		(_1710_v0).ArraySet1((_dafny.IntOfUint32((_1712_v1).Cardinality())).Cmp(p1) > 0, (_index308).Int())
		var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
		_ = _index309
		(_1710_v0).ArraySet1((_this).Fm2(_this.F14(), _this.F14(), globalState), (_index309).Int())
		(globalState).F9 = (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool)
		var _1713_v2 _dafny.Array
		_ = _1713_v2
		var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
		_ = _nw274
		_1713_v2 = _nw274
		var _1714_v3 _dafny.Array
		_ = _1714_v3
		var _nwElement0_55 _dafny.Array = _1713_v2
		_ = _nwElement0_55
		var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(17))
		_ = _nw275
		(_nw275).ArraySet1(_nwElement0_55, 0)
		(_nw275).ArraySet1(_1713_v2, 1)
		(_nw275).ArraySet1(_1713_v2, 2)
		(_nw275).ArraySet1(_1713_v2, 3)
		(_nw275).ArraySet1(_1713_v2, 4)
		(_nw275).ArraySet1(_1713_v2, 5)
		(_nw275).ArraySet1(_1713_v2, 6)
		(_nw275).ArraySet1(_1713_v2, 7)
		(_nw275).ArraySet1(_1713_v2, 8)
		(_nw275).ArraySet1(_1713_v2, 9)
		(_nw275).ArraySet1(_1713_v2, 10)
		(_nw275).ArraySet1(_1713_v2, 11)
		(_nw275).ArraySet1(_1713_v2, 12)
		(_nw275).ArraySet1(_1713_v2, 13)
		(_nw275).ArraySet1(_1713_v2, 14)
		(_nw275).ArraySet1(_1713_v2, 15)
		(_nw275).ArraySet1((Companion_D14_.Create_DC44_(_1713_v2)).Dtor_cf71(), 16)
		_1714_v3 = _nw275
		var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1714_v3), 0))
		_ = _index310
		(_1714_v3).ArraySet1(_1713_v2, (_index310).Int())
		var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.ArrayLen((_1714_v3), 0))
		_ = _index311
		(_1714_v3).ArraySet1((func() _dafny.Array {
			if (Companion_D1_.Create_DC5_(_this.F14(), true)).Dtor_cf7() {
				return _1713_v2
			}
			return (func() _dafny.Array {
				if (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool) {
					return _1713_v2
				}
				return _1713_v2
			})()
		})(), (_index311).Int())
		var _1715_v4 _dafny.CodePoint
		_ = _1715_v4
		_1715_v4 = _dafny.CodePoint('y')
		var _1716_v5 _dafny.Set
		_ = _1716_v5
		_1716_v5 = _dafny.SetOf(_1715_v4)
		var _1717_v6 _dafny.MultiSet
		_ = _1717_v6
		_1717_v6 = _dafny.MultiSetOf(p1)
		_1716_v5 = Companion_Default___.Fm32(_1712_v1, _1717_v6, globalState)
		var _1718_v7 D4
		_ = _1718_v7
		_1718_v7 = Companion_D4_.Create_DC12_(_1710_v0)
		var _1719_v8 _dafny.Sequence
		_ = _1719_v8
		_1719_v8 = _dafny.SeqOf(_1718_v7)
		var _1720_v9 *C4
		_ = _1720_v9
		var _nw276 *C4 = New_C4_()
		_ = _nw276
		_nw276.Ctor__(!_dafny.Companion_Sequence_.Equal(_1719_v8, _1719_v8), (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), (_this).F15())
		_1720_v9 = _nw276
		if (_1720_v9).F19() {
			if (_1720_v9).F19() {
				_1710_v0 = _1710_v0
				var _1721_v10 _dafny.Sequence
				_ = _1721_v10
				_1721_v10 = _dafny.SeqOf(_1712_v1, _1712_v1, _dafny.Companion_Sequence_.Concatenate(_1712_v1, _1712_v1), _1712_v1)
				_1721_v10 = _1721_v10
				(globalState).F9 = (_1720_v9).F19()
				var _1722_v11 _dafny.Map
				_ = _1722_v11
				_1722_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_this).F15()).Cardinality())
				var _1723_v12 _dafny.Map
				_ = _1723_v12
				_1723_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1712_v1).Cardinality()), _1722_v11)
				var _1724_v13 _dafny.Map
				_ = _1724_v13
				_1724_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, !(!((_1720_v9).F19()))), p1)
				var _1725_v14 D7
				_ = _1725_v14
				_1725_v14 = Companion_D7_.Create_DC21_(_1724_v13)
				var _1726_v15 _dafny.Map
				_ = _1726_v15
				_1726_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1725_v14, _1716_v5)
				var _1727_v16 _dafny.Sequence
				_ = _1727_v16
				_1727_v16 = _dafny.SeqOf(_1723_v12, _1723_v12)
				var _rhs283 _dafny.Set = ((func() _dafny.Set {
					if (_1726_v15).Contains(_1725_v14) {
						return (_1726_v15).Get(_1725_v14).(_dafny.Set)
					}
					return _1716_v5
				})()).Intersection(_dafny.SetOf(_dafny.CodePoint('f')))
				_ = _rhs283
				var _rhs284 _dafny.Int = p1
				_ = _rhs284
				var _rhs285 bool = !((p1).Cmp(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.Fm0((_1720_v9).F19(), (_dafny.Zero).Minus(p1), p1, _1717_v6, globalState)), p1)).Cardinality())) <= 0)
				_ = _rhs285
				var _rhs286 _dafny.Map = (_1727_v16).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p1, p1), _dafny.IntOfUint32((_1727_v16).Cardinality()))).Uint32()).(_dafny.Map)
				_ = _rhs286
				var _lhs248 *GlobalState = globalState
				_ = _lhs248
				_1716_v5 = _rhs283
				_lhs248.F3 = _rhs284
				r2 = _rhs285
				_1723_v12 = _rhs286
				var _1728_v17 *C1
				_ = _1728_v17
				var _nw277 *C1 = New_C1_()
				_ = _nw277
				_nw277.Ctor__()
				_1728_v17 = _nw277
			} else {
				var _1729_v18 D4
				_ = _1729_v18
				_1729_v18 = Companion_D4_.Create_DC13_(_1712_v1, p1, (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), _1710_v0)
				var _rhs287 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1729_v18).Dtor_cf22(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(292))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg107 _dafny.Int) interface{} {
						return coer107(arg107)
					}
				}((func(_1730_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1731_i1 _dafny.Int) _dafny.CodePoint {
						return _1730_v4
					}
				})(_1715_v4))))
				_ = _rhs287
				var _rhs288 _dafny.Sequence = _1712_v1
				_ = _rhs288
				_1712_v1 = _rhs287
				_1712_v1 = _rhs288
				var _1732_v19 _dafny.Map
				_ = _1732_v19
				_1732_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1720_v9).F19(), p1)
				var _1733_v20 _dafny.Array
				_ = _1733_v20
				var _nwElement0_56 _dafny.CodePoint = _1715_v4
				_ = _nwElement0_56
				var _nw278 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(19))
				_ = _nw278
				(_nw278).ArraySet1CodePoint(_nwElement0_56, 0)
				(_nw278).ArraySet1CodePoint(_1715_v4, 1)
				(_nw278).ArraySet1CodePoint(_1715_v4, 2)
				(_nw278).ArraySet1CodePoint(_1715_v4, 3)
				(_nw278).ArraySet1CodePoint(_1715_v4, 4)
				(_nw278).ArraySet1CodePoint(_1715_v4, 5)
				(_nw278).ArraySet1CodePoint(_1715_v4, 6)
				(_nw278).ArraySet1CodePoint((_1712_v1).Select((Companion_Default___.SafeIndex((_1717_v6).Cardinality(), _dafny.IntOfUint32((_1712_v1).Cardinality()))).Uint32()).(_dafny.CodePoint), 7)
				(_nw278).ArraySet1CodePoint(_1715_v4, 8)
				(_nw278).ArraySet1CodePoint(_1715_v4, 9)
				(_nw278).ArraySet1CodePoint(Companion_Default___.Fm25(p1, _dafny.IntOfInt64(129), _this.F14(), _1732_v19, globalState), 10)
				(_nw278).ArraySet1CodePoint(_1715_v4, 11)
				(_nw278).ArraySet1CodePoint(_1715_v4, 12)
				(_nw278).ArraySet1CodePoint(_dafny.CodePoint('t'), 13)
				(_nw278).ArraySet1CodePoint(_dafny.CodePoint('g'), 14)
				(_nw278).ArraySet1CodePoint(_1715_v4, 15)
				(_nw278).ArraySet1CodePoint(_1715_v4, 16)
				(_nw278).ArraySet1CodePoint(_1715_v4, 17)
				(_nw278).ArraySet1CodePoint(_1715_v4, 18)
				_1733_v20 = _nw278
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1733_v20), 0))
				_ = _index312
				(_1733_v20).ArraySet1CodePoint(_1715_v4, (_index312).Int())
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1733_v20), 0))
				_ = _index313
				(_1733_v20).ArraySet1CodePoint(_1715_v4, (_index313).Int())
				var _1734_v21 _dafny.Array
				_ = _1734_v21
				var _nw279 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(21))
				_ = _nw279
				_1734_v21 = _nw279
				var _1735_v22 _dafny.Set
				_ = _1735_v22
				_1735_v22 = _dafny.SetOf(Companion_D8_.Create_DC25_())
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1734_v21), 0))
				_ = _index314
				(_1734_v21).ArraySet1(_1735_v22, (_index314).Int())
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_1734_v21), 0))
				_ = _index315
				(_1734_v21).ArraySet1(_1735_v22, (_index315).Int())
				var _1736_v24 _dafny.Map
				_ = _1736_v24
				_1736_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Contains(_1712_v1, (_1733_v20).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1733_v20), 0))).Int())), (p0).Intersection(func() _dafny.Set {
					var _coll71 = _dafny.NewBuilder()
					_ = _coll71
					for _iter76 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(636), _dafny.IntOfInt64(181))); ; {
						_compr_71, _ok76 := _iter76()
						if !_ok76 {
							break
						}
						var _1737_v23 _dafny.Int
						_1737_v23 = interface{}(_compr_71).(_dafny.Int)
						if ((_dafny.IntOfInt64(636)).Cmp(_1737_v23) <= 0) && ((_1737_v23).Cmp(_dafny.IntOfInt64(181)) < 0) {
							_coll71.Add(Companion_Default___.SafeDivisionInt(_1737_v23, p1))
						}
					}
					return _coll71.ToSet()
				}()))
				_1736_v24 = _1736_v24
				var _1738_v25 _dafny.Map
				_ = _1738_v25
				_1738_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1712_v1, ((_this).F15()).Cardinality())
				_1738_v25 = (_1738_v25).Update(_1712_v1, (_dafny.IntOfUint32((_1712_v1).Cardinality())).Minus((_dafny.Zero).Minus(p1)))
			}
			var _1739_v26 _dafny.Map
			_ = _1739_v26
			_1739_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1720_v9).F19(), !((p1).Cmp(p1) <= 0))
			_1739_v26 = (_1739_v26).Update(_this.F14(), (p1).Cmp(p1) < 0)
			_1712_v1 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-802))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg108 _dafny.Int) interface{} {
					return coer108(arg108)
				}
			}((func(_1740_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1741_i2 _dafny.Int) _dafny.CodePoint {
					return _1740_v4
				}
			})(_1715_v4)))
			var _1742_v27 _dafny.Array
			_ = _1742_v27
			var _nwElement0_57 _dafny.Set = p0
			_ = _nwElement0_57
			var _nw280 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(5))
			_ = _nw280
			(_nw280).ArraySet1(_nwElement0_57, 0)
			(_nw280).ArraySet1(p0, 1)
			(_nw280).ArraySet1(p0, 2)
			(_nw280).ArraySet1(p0, 3)
			(_nw280).ArraySet1((p0).Union(p0), 4)
			_1742_v27 = _nw280
			var _1743_v28 _dafny.Map
			_ = _1743_v28
			_1743_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1720_v9).F19())
			var _1744_v29 _dafny.Map
			_ = _1744_v29
			_1744_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1743_v28).Cardinality(), p1)
			var _1745_v30 _dafny.Sequence
			_ = _1745_v30
			_1745_v30 = _dafny.SeqOf(p1)
			var _1746_v31 _dafny.Sequence
			_ = _1746_v31
			_1746_v31 = _dafny.SeqOf((_1744_v29).Cardinality(), p1, Companion_Default___.Fm0(_this.F14(), p1, p1, _1717_v6, globalState), _dafny.IntOfUint32((_1745_v30).Cardinality()), p1)
			var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1742_v27), 0))
			_ = _index316
			(_1742_v27).ArraySet1((_dafny.SetOf(p1, p1)).Union(func() _dafny.Set {
				var _coll72 = _dafny.NewBuilder()
				_ = _coll72
				for _iter77 := _dafny.Iterate((_1745_v30).Elements()); ; {
					_compr_72, _ok77 := _iter77()
					if !_ok77 {
						break
					}
					var _1747_v32 _dafny.Int
					_1747_v32 = interface{}(_compr_72).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_1745_v30, _1747_v32) {
						_coll72.Add((_1747_v32).Plus(_dafny.IntOfInt64(876)))
					}
				}
				return _coll72.ToSet()
			}()), (_index316).Int())
			var _1748_v33 _dafny.Array
			_ = _1748_v33
			var _nw281 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw281
			_1748_v33 = _nw281
			var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))
			_ = _index317
			(_1748_v33).ArraySet1((func() _dafny.Int {
				if (_1744_v29).Contains(p1) {
					return (_1744_v29).Get(p1).(_dafny.Int)
				}
				return p1
			})(), (_index317).Int())
			var _1749_v34 _dafny.Set
			_ = _1749_v34
			_1749_v34 = _dafny.SetOf(_1712_v1)
			var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1742_v27), 0))
			_ = _index318
			var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))
			_ = _index319
			var _rhs289 _dafny.Set = _dafny.SetOf(p1)
			_ = _rhs289
			var _rhs290 _dafny.Int = p1
			_ = _rhs290
			var _rhs291 _dafny.Int = p1
			_ = _rhs291
			var _rhs292 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-551))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}((func(_1750_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1751_i3 _dafny.Int) _dafny.CodePoint {
					return _1750_v4
				}
			})(_1715_v4)))
			_ = _rhs292
			var _rhs293 _dafny.Set = _1749_v34
			_ = _rhs293
			var _lhs249 _dafny.Array = _1742_v27
			_ = _lhs249
			var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(338), _dafny.ArrayLen((_1742_v27), 0))
			_ = _lhs250
			var _lhs251 _dafny.Array = _1748_v33
			_ = _lhs251
			var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))
			_ = _lhs252
			(_lhs249).ArraySet1(_rhs289, (_lhs250).Int())
			r1 = _rhs290
			(_lhs251).ArraySet1(_rhs291, (_lhs252).Int())
			_1712_v1 = _rhs292
			_1749_v34 = _rhs293
			var _1752_v35 _dafny.Map
			_ = _1752_v35
			_1752_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), (_1748_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))).Int()).(_dafny.Int)), _1748_v33)
			var _1753_v36 _dafny.Map
			_ = _1753_v36
			_1753_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_1748_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))).Int()).(_dafny.Int))
			var _1754_v37 _dafny.Sequence
			_ = _1754_v37
			_1754_v37 = _dafny.SeqOf((func() _dafny.Array {
				if (_1752_v35).Contains(_1753_v36) {
					return (_1752_v35).Get(_1753_v36).(_dafny.Array)
				}
				return _1748_v33
			})(), _1748_v33)
			var _1755_v38 D5
			_ = _1755_v38
			_1755_v38 = Companion_D5_.Create_DC16_(_1748_v33)
			var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
			_ = _index320
			var _rhs294 bool = (Companion_D0_.Create_DC0_((_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool))).Dtor_cf0()
			_ = _rhs294
			var _rhs295 _dafny.Sequence = _dafny.SeqOf(_1748_v33, _1748_v33, _1748_v33, (_1755_v38).Dtor_cf29(), _1748_v33)
			_ = _rhs295
			var _rhs296 _dafny.Int = (_1748_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1748_v33), 0))).Int()).(_dafny.Int)
			_ = _rhs296
			var _rhs297 bool = !(true)
			_ = _rhs297
			var _lhs253 _dafny.Array = _1710_v0
			_ = _lhs253
			var _lhs254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
			_ = _lhs254
			var _lhs255 *GlobalState = globalState
			_ = _lhs255
			(_lhs253).ArraySet1(_rhs294, (_lhs254).Int())
			_1754_v37 = _rhs295
			r1 = _rhs296
			_lhs255.F9 = _rhs297
		} else {
			_1715_v4 = _1715_v4
			var _1756_v39 _dafny.Array
			_ = _1756_v39
			var _nw282 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
			_ = _nw282
			_1756_v39 = _nw282
			var _nw283 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(4))
			_ = _nw283
			_1756_v39 = _nw283
			r1 = (p1).Minus(p1)
			if (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool) {
				var _1757_v40 _dafny.Array
				_ = _1757_v40
				var _len0_49 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_49
				var _nw284 _dafny.Array
				_ = _nw284
				if _len0_49.Cmp(_dafny.Zero) == 0 {
					_nw284 = _dafny.NewArray(_len0_49)
				} else {
					var _init49 func(_dafny.Int) _dafny.Int = (func(_1758_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1759_i4 _dafny.Int) _dafny.Int {
							return (_1759_i4).Minus(_1758_p1)
						}
					})(p1)
					_ = _init49
					var _element0_49 = _init49(_dafny.Zero)
					_ = _element0_49
					_nw284 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
					(_nw284).ArraySet1(_element0_49, 0)
					var _nativeLen0_49 = (_len0_49).Int()
					_ = _nativeLen0_49
					for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
						(_nw284).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
					}
				}
				_1757_v40 = _nw284
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))
				_ = _index321
				(_1757_v40).ArraySet1(p1, (_index321).Int())
				var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))
				_ = _index322
				(_1757_v40).ArraySet1(p1, (_index322).Int())
				var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
				_ = _index323
				(_1710_v0).ArraySet1((_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), (_index323).Int())
				var _1760_v41 _dafny.Map
				_ = _1760_v41
				_1760_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(303))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg110 _dafny.Int) interface{} {
						return coer110(arg110)
					}
				}((func(_1761_v4 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1762_i5 _dafny.Int) _dafny.CodePoint {
						return _1761_v4
					}
				})(_1715_v4)))).Cardinality()), (func() _dafny.Sequence {
					if (_1720_v9).F19() {
						return _1712_v1
					}
					return _1712_v1
				})())
				_1760_v41 = (_1760_v41).Update((_dafny.IntOfUint32((_1712_v1).Cardinality())).Minus(_dafny.IntOfUint32((_1712_v1).Cardinality())), (func() _dafny.Sequence {
					if (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool) {
						return _dafny.UnicodeSeqOfUtf8Bytes("aupfiojbx")
					}
					return _1712_v1
				})())
				var _1763_v42 _dafny.Map
				_ = _1763_v42
				_1763_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), (_1757_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))).Int()).(_dafny.Int))
				var _1764_v43 _dafny.Map
				_ = _1764_v43
				_1764_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1715_v4, _dafny.IntOfInt64(720))
				var _1765_v44 _dafny.Sequence
				_ = _1765_v44
				_1765_v44 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1720_v9).F19(), _1715_v4)).Cardinality(), (_dafny.Zero).Minus(p1))
				var _1766_v45 _dafny.Set
				_ = _1766_v45
				_1766_v45 = _dafny.SetOf(_1765_v44, _1765_v44)
				var _1767_v46 D13
				_ = _1767_v46
				_1767_v46 = Companion_D13_.Create_DC41_(_1766_v45)
				var _1768_v47 _dafny.Map
				_ = _1768_v47
				_1768_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_1763_v42).Contains((_1720_v9).F19()) {
						return (_1763_v42).Get((_1720_v9).F19()).(_dafny.Int)
					}
					return (_1757_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))).Int()).(_dafny.Int)
				})(), Companion_Default___.Fm23(p1, (_1720_v9).F19(), _1764_v43, (_1767_v46).Dtor_cf67(), globalState))
				var _1769_v48 *C2
				_ = _1769_v48
				var _nw285 *C2 = New_C2_()
				_ = _nw285
				_nw285.Ctor__(_1768_v47)
				_1769_v48 = _nw285
				var _1770_v49 *C5
				_ = _1770_v49
				var _nw286 *C5 = New_C5_()
				_ = _nw286
				_nw286.Ctor__(((_1757_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))).Int()).(_dafny.Int)).Times(p1), ((_1757_v40).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(831), _dafny.ArrayLen((_1757_v40), 0))).Int()).(_dafny.Int)).Cmp(p1) > 0, ((_this).F15()).Union((_this).F15()))
				_1770_v49 = _nw286
			} else {
				(_1720_v9).M6(globalState)
				var _1771_v50 _dafny.Map
				_ = _1771_v50
				_1771_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1717_v6).Cardinality(), _this.F14())
				var _1772_v51 _dafny.Sequence
				_ = _1772_v51
				_1772_v51 = _dafny.SeqOf((_1720_v9).F19())
				var _1773_v52 *C2
				_ = _1773_v52
				var _nw287 *C2 = New_C2_()
				_ = _nw287
				_nw287.Ctor__((_1771_v50).Merge((_1771_v50).Update(_dafny.IntOfUint32((_dafny.SeqOf(_1772_v51)).Cardinality()), (_1720_v9).F19())))
				_1773_v52 = _nw287
				_1773_v52 = _1773_v52
				var _1774_v53 _dafny.Sequence
				_ = _1774_v53
				_1774_v53 = _dafny.SeqOf(p1, p1)
				var _1775_v54 *C0
				_ = _1775_v54
				var _nw288 *C0 = New_C0_()
				_ = _nw288
				_nw288.Ctor__(_this.F14(), (_1720_v9).F19())
				_1775_v54 = _nw288
				var _1776_v55 _dafny.Set
				_ = _1776_v55
				_1776_v55 = _dafny.SetOf(_1775_v54)
				var _1777_v56 _dafny.Sequence
				_ = _1777_v56
				_1777_v56 = _dafny.SeqOf(_1776_v55, _1776_v55)
				var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))
				_ = _index324
				(_1710_v0).ArraySet1((p1).Cmp(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_1774_v53).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1712_v1).Cardinality()), _dafny.IntOfUint32((_1774_v53).Cardinality()))).Uint32()).(_dafny.Int))), _dafny.IntOfUint32((_1777_v56).Cardinality()))) != 0, (_index324).Int())
				var _1778_v57 *C3
				_ = _1778_v57
				var _nw289 *C3 = New_C3_()
				_ = _nw289
				_nw289.Ctor__(_dafny.IntOfInt64(360), _1712_v1, (_1775_v54).F23(), (_this).F15())
				_1778_v57 = _nw289
				var _1779_v58 _dafny.Map
				_ = _1779_v58
				_1779_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1778_v57, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfija")).Cardinality()))
				var _1780_v59 D15
				_ = _1780_v59
				_1780_v59 = Companion_D15_.Create_DC47_(_1779_v58)
				r1 = (Companion_D3_.Create_DC11_((_dafny.Zero).Minus(p1), (_1720_v9).Fm17(((_1780_v59).Dtor_cf76()).Cardinality(), _this.F14(), _dafny.MultiSetOf(_1772_v51), (_1778_v57).F20(), globalState), (_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), (_1778_v57).F20())).Dtor_cf17()
				_1774_v53 = _1774_v53
			}
			_1715_v4 = _1715_v4
		}
		var _1781_v60 _dafny.Set
		_ = _1781_v60
		_1781_v60 = _dafny.SetOf((_1720_v9).F19(), true, (_1720_v9).F19(), true, true)
		r0 = (func() _dafny.Int {
			if (p1).Cmp((func() _dafny.Int {
				if (_1717_v6).Contains((_1781_v60).Cardinality()) {
					return (_1717_v6).Multiplicity((_1781_v60).Cardinality())
				}
				return p1
			})()) >= 0 {
				return p1
			}
			return p1
		})()
		r1 = ((p1).Plus(p1)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1712_v1, _1712_v1)).Cardinality()))
		var _1782_v61 *C0
		_ = _1782_v61
		var _nw290 *C0 = New_C0_()
		_ = _nw290
		_nw290.Ctor__((_1710_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_1710_v0), 0))).Int()).(bool), !(true))
		_1782_v61 = _nw290
		var _1783_v62 _dafny.Map
		_ = _1783_v62
		_1783_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1782_v61, _1781_v60)
		r2 = (_1783_v62).Equals(_1783_v62)
		return r0, r1, r2
	}
}
func (_this *C6) M5(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.MultiSet, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r2
		var r3 bool = false
		_ = r3
		var _1784_v0 _dafny.Array
		_ = _1784_v0
		var _len0_50 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_50
		var _nw291 _dafny.Array
		_ = _nw291
		if _len0_50.Cmp(_dafny.Zero) == 0 {
			_nw291 = _dafny.NewArray(_len0_50)
		} else {
			var _init50 func(_dafny.Int) bool = func(_1785_i1 _dafny.Int) bool {
				return _this.F14()
			}
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
		_1784_v0 = _nw291
		for _iter78 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1784_v0), 0))); ; {
			_guard_loop_5, _ok78 := _iter78()
			if !_ok78 {
				break
			}
			var _1786_i0 _dafny.Int
			_1786_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1786_i0).Sign() != -1) && ((_1786_i0).Cmp(_dafny.ArrayLen((_1784_v0), 0)) < 0)) {
				(_1784_v0).ArraySet1((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-320), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vuy")).Cardinality()))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-304))).Uint32(), func(coer111 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}(func(_1787_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				})), _dafny.UnicodeSeqOfUtf8Bytes("qdea"))).Cardinality())) != 0, (_1786_i0).Int())
			}
		}
		if _this.F14() {
			r3 = _this.F14()
			r3 = _this.F14()
			var _1788_v1 D0
			_ = _1788_v1
			_1788_v1 = Companion_D0_.Create_DC0_(_this.F14())
			var _source17 D0 = _1788_v1
			_ = _source17
			if _source17.Is_DC1() {
				var _1789___mcc_h0 bool = _source17.Get_().(D0_DC1).Cf1
				_ = _1789___mcc_h0
				var _1790_cf1 bool = _1789___mcc_h0
				_ = _1790_cf1
				var _1791_v2 _dafny.Int
				_ = _1791_v2
				_1791_v2 = _dafny.IntOfInt64(-981)
				r0 = _1791_v2
				var _1792_v3 _dafny.Sequence
				_ = _1792_v3
				_1792_v3 = _dafny.UnicodeSeqOfUtf8Bytes("ikpk")
				var _1793_v4 _dafny.Map
				_ = _1793_v4
				_1793_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1790_cf1, _1792_v3)
				var _1794_v5 _dafny.Sequence
				_ = _1794_v5
				_1794_v5 = _dafny.SeqOf(_1792_v3, _1792_v3, _1792_v3, _1792_v3, _1792_v3)
				var _1795_v6 _dafny.CodePoint
				_ = _1795_v6
				_1795_v6 = _dafny.CodePoint('f')
				var _1796_v7 _dafny.Set
				_ = _1796_v7
				_1796_v7 = _dafny.SetOf(_1791_v2)
				var _1797_v8 D9
				_ = _1797_v8
				_1797_v8 = Companion_D9_.Create_DC30_(_1795_v6, _1790_cf1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(875))).Uint32(), func(coer112 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_1798_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1799_i3 _dafny.Int) _dafny.CodePoint {
						return _1798_v6
					}
				})(_1795_v6)))).Cardinality()), _1796_v7, _1791_v2)
				var _1800_v9 *C3
				_ = _1800_v9
				var _nw292 *C3 = New_C3_()
				_ = _nw292
				_nw292.Ctor__(((_1793_v4).Merge(_1793_v4)).Cardinality(), (_1794_v5).Select((Companion_Default___.SafeIndex(_1791_v2, _dafny.IntOfUint32((_1794_v5).Cardinality()))).Uint32()).(_dafny.Sequence), (func() bool {
					if _1790_cf1 {
						return (_1797_v8).Dtor_cf51()
					}
					return _1790_cf1
				})(), (_this).F15())
				_1800_v9 = _nw292
				var _1801_v10 *C3
				_ = _1801_v10
				var _nw293 *C3 = New_C3_()
				_ = _nw293
				_nw293.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf((_1800_v9).F20(), (_1800_v9).F20(), (_1800_v9).F20())).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_1792_v3, _1792_v3), true, _dafny.MultiSetFromSeq(_dafny.SeqOf(_this.F14(), !(_1790_cf1), _this.F14())))
				_1801_v10 = _nw293
				var _1802_v11 _dafny.Sequence
				_ = _1802_v11
				_1802_v11 = _dafny.SeqOf((_1801_v10).F20(), _1791_v2)
				r0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1802_v11, _dafny.Companion_Sequence_.Concatenate(_1802_v11, _1802_v11))).Cardinality())
			} else if _source17.Is_DC0() {
				var _1803___mcc_h1 bool = _source17.Get_().(D0_DC0).Cf0
				_ = _1803___mcc_h1
				var _1804_cf0 bool = _1803___mcc_h1
				_ = _1804_cf0
				var _1805_v12 _dafny.Int
				_ = _1805_v12
				_1805_v12 = _dafny.IntOfInt64(83)
				var _1806_v13 *C5
				_ = _1806_v13
				var _nw294 *C5 = New_C5_()
				_ = _nw294
				_nw294.Ctor__((_1805_v12).Plus(_1805_v12), _this.F14(), Companion_Default___.Fm33(_this.F14(), globalState))
				_1806_v13 = _nw294
				var _1807_v14 _dafny.CodePoint
				_ = _1807_v14
				_1807_v14 = _dafny.CodePoint('r')
				var _1808_v15 _dafny.Map
				_ = _1808_v15
				_1808_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1807_v14, _dafny.IntOfInt64(66))
				var _1809_v16 _dafny.Sequence
				_ = _1809_v16
				_1809_v16 = _dafny.SeqOf((_1806_v13).F18())
				var _1810_v17 _dafny.Set
				_ = _1810_v17
				_1810_v17 = _dafny.SetOf(_1809_v16, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(293))).Uint32(), func(coer113 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_1811_v13 *C5) func(_dafny.Int) _dafny.Int {
					return func(_1812_i4 _dafny.Int) _dafny.Int {
						return (_1811_v13).F18()
					}
				})(_1806_v13))), _1809_v16, _1809_v16, _1809_v16)
				var _1813_v18 *C4
				_ = _1813_v18
				var _nw295 *C4 = New_C4_()
				_ = _nw295
				_nw295.Ctor__(_1804_cf0, !(_1804_cf0) || (!(Companion_Default___.Fm23(_1805_v12, _1804_cf0, _1808_v15, _1810_v17, globalState))), (_this).F15())
				_1813_v18 = _nw295
				var _1814_v19 _dafny.Map
				_ = _1814_v19
				_1814_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1805_v12, _this.F14())
				var _1815_v20 *C2
				_ = _1815_v20
				var _nw296 *C2 = New_C2_()
				_ = _nw296
				_nw296.Ctor__(_1814_v19)
				_1815_v20 = _nw296
				(globalState).F3 = _1805_v12
			} else {
				var _1816___mcc_h2 D0 = _source17.Get_().(D0_DC2).Cf2
				_ = _1816___mcc_h2
				var _1817_cf2 D0 = _1816___mcc_h2
				_ = _1817_cf2
				var _1818_v21 _dafny.Set
				_ = _1818_v21
				_1818_v21 = _dafny.SetOf(_this.F14())
				r1 = (_1818_v21).Cardinality()
				var _1819_v22 _dafny.Int
				_ = _1819_v22
				_1819_v22 = _dafny.IntOfInt64(999)
				var _1820_v23 D12
				_ = _1820_v23
				_1820_v23 = Companion_D12_.Create_DC38_(_1819_v22)
				_1820_v23 = Companion_D12_.Create_DC38_(_1819_v22)
				(globalState).F2 = _1784_v0
				var _1821_v24 _dafny.Array
				_ = _1821_v24
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_51
				var _nw297 _dafny.Array
				_ = _nw297
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw297 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) D8 = (func(_1822_v21 _dafny.Set) func(_dafny.Int) D8 {
						return func(_1823_i5 _dafny.Int) D8 {
							return Companion_D8_.Create_DC24_(_1822_v21)
						}
					})(_1818_v21)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw297 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw297).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw297).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_1821_v24 = _nw297
				var _1824_v25 D8
				_ = _1824_v25
				_1824_v25 = Companion_D8_.Create_DC24_(_1818_v21)
				var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1821_v24), 0))
				_ = _index325
				(_1821_v24).ArraySet1(func(_pat_let58_0 D8) D8 {
					return func(_1825_dt__update__tmp_h0 D8) D8 {
						return func(_pat_let59_0 _dafny.Set) D8 {
							return func(_1826_dt__update_hcf45_h0 _dafny.Set) D8 {
								return Companion_D8_.Create_DC24_(_1826_dt__update_hcf45_h0)
							}(_pat_let59_0)
						}(_dafny.SetOf(_this.F14()))
					}(_pat_let58_0)
				}(_1824_v25), (_index325).Int())
				var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(774), _dafny.ArrayLen((_1821_v24), 0))
				_ = _index326
				(_1821_v24).ArraySet1(_1824_v25, (_index326).Int())
			}
			var _1827_v26 _dafny.Array
			_ = _1827_v26
			var _len0_52 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_52
			var _nw298 _dafny.Array
			_ = _nw298
			if _len0_52.Cmp(_dafny.Zero) == 0 {
				_nw298 = _dafny.NewArray(_len0_52)
			} else {
				var _init52 func(_dafny.Int) _dafny.Sequence = func(_1828_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tqux"), _dafny.UnicodeSeqOfUtf8Bytes("mdsulisdp"))
				}
				_ = _init52
				var _element0_52 = _init52(_dafny.Zero)
				_ = _element0_52
				_nw298 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
				(_nw298).ArraySet1(_element0_52, 0)
				var _nativeLen0_52 = (_len0_52).Int()
				_ = _nativeLen0_52
				for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
					(_nw298).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
				}
			}
			_1827_v26 = _nw298
			var _1829_v27 _dafny.Set
			_ = _1829_v27
			_1829_v27 = _dafny.SetOf((_this).F15())
			var _1830_v28 D11
			_ = _1830_v28
			_1830_v28 = Companion_D11_.Create_DC33_(_1829_v27)
			var _pat_let_tv43 = _1829_v27
			_ = _pat_let_tv43
			var _rhs298 _dafny.Array = _1827_v26
			_ = _rhs298
			var _rhs299 D11 = func(_pat_let60_0 D11) D11 {
				return func(_1831_dt__update__tmp_h1 D11) D11 {
					return func(_pat_let61_0 _dafny.Set) D11 {
						return func(_1832_dt__update_hcf57_h0 _dafny.Set) D11 {
							return Companion_D11_.Create_DC33_(_1832_dt__update_hcf57_h0)
						}(_pat_let61_0)
					}(_pat_let_tv43)
				}(_pat_let60_0)
			}(Companion_D11_.Create_DC33_(_1829_v27))
			_ = _rhs299
			_1827_v26 = _rhs298
			_1830_v28 = _rhs299
			var _1833_v29 *C1
			_ = _1833_v29
			var _nw299 *C1 = New_C1_()
			_ = _nw299
			_nw299.Ctor__()
			_1833_v29 = _nw299
		} else {
			var _1834_v30 _dafny.Int
			_ = _1834_v30
			_1834_v30 = _dafny.IntOfInt64(140)
			var _1835_v31 _dafny.Map
			_ = _1835_v31
			_1835_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1834_v30, _dafny.IntOfInt64(41))
			_1835_v31 = (_1835_v31).Update(_1834_v30, (_1834_v30).Times(_1834_v30))
			var _1836_v32 _dafny.Sequence
			_ = _1836_v32
			_1836_v32 = _dafny.UnicodeSeqOfUtf8Bytes("shmrn")
			if !(_this.F14()) || (!(_dafny.Companion_Sequence_.IsPrefixOf(_1836_v32, _1836_v32))) {
				var _1837_v33 *C0
				_ = _1837_v33
				var _nw300 *C0 = New_C0_()
				_ = _nw300
				_nw300.Ctor__(_this.F14(), _this.F14())
				_1837_v33 = _nw300
				_1834_v30 = _dafny.IntOfInt64(3)
				var _1838_v34 _dafny.Set
				_ = _1838_v34
				_1838_v34 = _dafny.SetOf(Companion_Default___.Fm27(_1834_v30, _this.F14(), globalState))
				_1838_v34 = ((_1838_v34).Union(_1838_v34)).Difference((func() _dafny.Set {
					if (_1837_v33).F24() {
						return _1838_v34
					}
					return _1838_v34
				})())
				var _1839_v35 _dafny.CodePoint
				_ = _1839_v35
				_1839_v35 = _dafny.CodePoint('n')
				_1839_v35 = _1839_v35
				var _1840_v36 _dafny.MultiSet
				_ = _1840_v36
				_1840_v36 = _dafny.MultiSetOf(_1834_v30)
				var _1841_v37 _dafny.Sequence
				_ = _1841_v37
				_1841_v37 = _dafny.SeqOf((_1837_v33).F23())
				var _1842_v38 _dafny.Set
				_ = _1842_v38
				_1842_v38 = _dafny.SetOf((_dafny.MultiSetOf(_this.F14(), (_1837_v33).F24())).Update(_this.F14(), Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(624))).Uint32(), func(coer114 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg114 _dafny.Int) interface{} {
						return coer114(arg114)
					}
				}((func(_1843_v35 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1844_i7 _dafny.Int) _dafny.CodePoint {
						return _1843_v35
					}
				})(_1839_v35)))).Cardinality()))), (_this).F15(), _dafny.MultiSetOf((_1837_v33).F23()), _dafny.MultiSetFromSeq(_1841_v37))
				var _rhs300 _dafny.Int = _1834_v30
				_ = _rhs300
				var _rhs301 _dafny.Int = ((_dafny.IntOfUint32((_1836_v32).Cardinality())).Minus((_1840_v36).Cardinality())).Plus((_1834_v30).Minus(_dafny.IntOfUint32((_1836_v32).Cardinality())))
				_ = _rhs301
				var _rhs302 bool = (_1842_v38).IsDisjointFrom(_1842_v38)
				_ = _rhs302
				var _lhs256 *GlobalState = globalState
				_ = _lhs256
				var _lhs257 *GlobalState = globalState
				_ = _lhs257
				_lhs256.F3 = _rhs300
				r0 = _rhs301
				_lhs257.F9 = _rhs302
			} else {
				var _1845_v39 _dafny.Map
				_ = _1845_v39
				_1845_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1834_v30)
				var _1846_v40 _dafny.CodePoint
				_ = _1846_v40
				_1846_v40 = _dafny.CodePoint('w')
				var _1847_v41 _dafny.Set
				_ = _1847_v41
				_1847_v41 = _dafny.SetOf(Companion_Default___.Fm25(_1834_v30, _1834_v30, _this.F14(), _1845_v39, globalState), _1846_v40)
				var _1848_v42 _dafny.Sequence
				_ = _1848_v42
				_1848_v42 = _dafny.SeqOf(_1847_v41)
				var _1849_v44 _dafny.Map
				_ = _1849_v44
				_1849_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('e'), _this.F14())
				var _1850_v45 _dafny.Sequence
				_ = _1850_v45
				_1850_v45 = _dafny.SeqOf(_1834_v30, _1834_v30)
				var _1851_v46 _dafny.Set
				_ = _1851_v46
				_1851_v46 = _dafny.SetOf(_1850_v45)
				var _rhs303 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus((_1834_v30).Plus(_dafny.IntOfInt64(658))))
				_ = _rhs303
				var _rhs304 bool = !((Companion_Default___.Fm23(_1834_v30, _this.F14(), func() _dafny.Map {
					var _coll73 = _dafny.NewMapBuilder()
					_ = _coll73
					for _iter79 := _dafny.Iterate((_1849_v44).Keys().Elements()); ; {
						_compr_73, _ok79 := _iter79()
						if !_ok79 {
							break
						}
						var _1852_v43 _dafny.CodePoint
						_1852_v43 = interface{}(_compr_73).(_dafny.CodePoint)
						if (_1849_v44).Contains(_1852_v43) {
							_coll73.Add(_1852_v43, _dafny.IntOfUint32((_dafny.SeqOf(_1834_v30, _1834_v30)).Cardinality()))
						}
					}
					return _coll73.ToMap()
				}(), _1851_v46, globalState)) && (!(_this.F14()) || (_this.F14())))
				_ = _rhs304
				var _rhs305 _dafny.Int = (func() _dafny.Int {
					if (_1835_v31).Contains(_dafny.IntOfUint32((_1836_v32).Cardinality())) {
						return (_1835_v31).Get(_dafny.IntOfUint32((_1836_v32).Cardinality())).(_dafny.Int)
					}
					return _1834_v30
				})()
				_ = _rhs305
				var _rhs306 _dafny.Sequence = _1848_v42
				_ = _rhs306
				var _lhs258 *GlobalState = globalState
				_ = _lhs258
				var _lhs259 *GlobalState = globalState
				_ = _lhs259
				_lhs258.F3 = _rhs303
				_lhs259.F9 = _rhs304
				_1834_v30 = _rhs305
				_1848_v42 = _rhs306
				var _1853_v47 _dafny.Set
				_ = _1853_v47
				_1853_v47 = _dafny.SetOf(!(_this.F14()))
				var _1854_v48 _dafny.Sequence
				_ = _1854_v48
				_1854_v48 = _dafny.SeqOf(_this.F14(), _dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("dpidlr"), Companion_Default___.Fm25(_1834_v30, (_1853_v47).Cardinality(), _this.F14(), _1845_v39, globalState)), _this.F14(), _this.F14())
				_1854_v48 = _1854_v48
				var _1855_v49 D0
				_ = _1855_v49
				_1855_v49 = Companion_D0_.Create_DC1_(_this.F14())
				var _1856_v50 D0
				_ = _1856_v50
				_1856_v50 = Companion_D0_.Create_DC2_(_1855_v49)
				var _1857_v51 _dafny.Map
				_ = _1857_v51
				_1857_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1856_v50, _this.F14())
				var _1858_v52 _dafny.Map
				_ = _1858_v52
				_1858_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1834_v30, (_this).Fm11(_1857_v51, _1834_v30, globalState))
				var _1859_v53 _dafny.Sequence
				_ = _1859_v53
				_1859_v53 = _dafny.SeqOf(_1858_v52)
				var _1860_v54 *C2
				_ = _1860_v54
				var _nw301 *C2 = New_C2_()
				_ = _nw301
				_nw301.Ctor__((func() _dafny.Map {
					if true {
						return (_1859_v53).Select((Companion_Default___.SafeIndex(_1834_v30, _dafny.IntOfUint32((_1859_v53).Cardinality()))).Uint32()).(_dafny.Map)
					}
					return _1858_v52
				})())
				_1860_v54 = _nw301
				var _1861_v55 _dafny.MultiSet
				_ = _1861_v55
				_1861_v55 = _dafny.MultiSetOf(_1834_v30, _1834_v30)
				r0 = Companion_Default___.Fm0(_this.F14(), _1834_v30, _1834_v30, _1861_v55, globalState)
				var _1862_v56 _dafny.Array
				_ = _1862_v56
				var _nw302 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw302
				_1862_v56 = _nw302
				var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1862_v56), 0))
				_ = _index327
				(_1862_v56).ArraySet1(_1834_v30, (_index327).Int())
				var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(407), _dafny.ArrayLen((_1862_v56), 0))
				_ = _index328
				(_1862_v56).ArraySet1((_1834_v30).Plus(_1834_v30), (_index328).Int())
			}
			var _1863_v57 _dafny.Sequence
			_ = _1863_v57
			_1863_v57 = _dafny.SeqOf(_this.F14())
			var _1864_v58 _dafny.Sequence
			_ = _1864_v58
			_1864_v58 = _dafny.SeqOf(_1863_v57)
			if !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, true, _this.F14(), _this.F14()), (_1864_v58).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.IntOfUint32((_1864_v58).Cardinality()))).Uint32()).(_dafny.Sequence)), _dafny.Companion_Sequence_.Update(_1863_v57, (Companion_Default___.SafeIndex(_1834_v30, _dafny.IntOfUint32((_1863_v57).Cardinality()))).Uint32(), _this.F14())) {
				var _1865_v59 _dafny.MultiSet
				_ = _1865_v59
				_1865_v59 = _dafny.MultiSetOf(_1834_v30, _dafny.IntOfInt64(943), _1834_v30, _dafny.IntOfInt64(97))
				r1 = (func() _dafny.Int {
					if (_1865_v59).Contains(Companion_Default___.SafeModuloInt(_1834_v30, Companion_Default___.Fm0(_this.F14(), _dafny.IntOfUint32((_1836_v32).Cardinality()), _1834_v30, _1865_v59, globalState))) {
						return (_1865_v59).Multiplicity(Companion_Default___.SafeModuloInt(_1834_v30, Companion_Default___.Fm0(_this.F14(), _dafny.IntOfUint32((_1836_v32).Cardinality()), _1834_v30, _1865_v59, globalState)))
					}
					return (_dafny.Zero).Minus(_1834_v30)
				})()
				var _1866_v60 *C3
				_ = _1866_v60
				var _nw303 *C3 = New_C3_()
				_ = _nw303
				_nw303.Ctor__(Companion_Default___.SafeModuloInt(_1834_v30, _1834_v30), _1836_v32, !(_this.F14()) || (_this.F14()), ((_this).F15()).Update(_this.F14(), Companion_Default___.Abs(_1834_v30)))
				_1866_v60 = _nw303
				_1836_v32 = _dafny.UnicodeSeqOfUtf8Bytes("srr")
				var _1867_v61 _dafny.Map
				_ = _1867_v61
				_1867_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1866_v60).F20(), (func() bool {
					if _this.F14() {
						return _this.F14()
					}
					return _this.F14()
				})())
				_1867_v61 = (_1867_v61).Update((func() _dafny.Int {
					if _this.F14() {
						return _dafny.IntOfInt64(737)
					}
					return _1834_v30
				})(), true)
				var _1868_v62 D5
				_ = _1868_v62
				_1868_v62 = Companion_D5_.Create_DC18_(_this.F14())
				(globalState).F9 = (_1868_v62).Dtor_cf33()
			} else {
				(globalState).F9 = ((_dafny.Zero).Minus((_1834_v30).Minus(_dafny.IntOfInt64(-320)))).Cmp(_1834_v30) < 0
				var _1869_v63 D1
				_ = _1869_v63
				_1869_v63 = Companion_D1_.Create_DC5_(_this.F14(), _this.F14())
				var _1870_v64 *C5
				_ = _1870_v64
				var _nw304 *C5 = New_C5_()
				_ = _nw304
				_nw304.Ctor__(_dafny.IntOfUint32((_1863_v57).Cardinality()), (_1869_v63).Dtor_cf6(), (_this).F15())
				_1870_v64 = _nw304
				var _1871_v65 D15
				_ = _1871_v65
				_1871_v65 = Companion_D15_.Create_DC48_(_1834_v30, _1870_v64, _dafny.IntOfInt64(491))
				(globalState).F3 = (_1871_v65).Dtor_cf79()
				var _1872_v66 _dafny.Map
				_ = _1872_v66
				_1872_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1784_v0, _this.F14())
				var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1784_v0), 0))
				_ = _index329
				(_1784_v0).ArraySet1(!((func() bool {
					if (_1872_v66).Contains(_1784_v0) {
						return (_1872_v66).Get(_1784_v0).(bool)
					}
					return _this.F14()
				})()), (_index329).Int())
				var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1784_v0), 0))
				_ = _index330
				(_1784_v0).ArraySet1(_this.F14(), (_index330).Int())
				var _1873_v67 D0
				_ = _1873_v67
				_1873_v67 = Companion_D0_.Create_DC0_((_1784_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(935), _dafny.ArrayLen((_1784_v0), 0))).Int()).(bool))
				(globalState).F3 = (func() _dafny.Int {
					if (_1873_v67).Dtor_cf0() {
						return _1834_v30
					}
					return (_dafny.Zero).Minus((_1870_v64).F18())
				})()
				var _1874_v68 D7
				_ = _1874_v68
				_1874_v68 = Companion_D7_.Create_DC23_(_this.F14(), _1834_v30, _this.F14())
				_1874_v68 = _1874_v68
			}
			(globalState).F2 = _1784_v0
			var _1875_v69 _dafny.CodePoint
			_ = _1875_v69
			_1875_v69 = _dafny.CodePoint('o')
			var _1876_v70 D2
			_ = _1876_v70
			_1876_v70 = Companion_D2_.Create_DC9_(_this.F14(), _1834_v30, _dafny.CodePoint('v'))
			(globalState).F9 = ((func() D2 {
				if _this.F14() {
					return Companion_D2_.Create_DC9_(_this.F14(), _1834_v30, _1875_v69)
				}
				return _1876_v70
			})()).Dtor_cf13()
		}
		var _1877_v71 _dafny.Array
		_ = _1877_v71
		var _nw305 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(6))
		_ = _nw305
		_1877_v71 = _nw305
		var _1878_v72 _dafny.Sequence
		_ = _1878_v72
		_1878_v72 = _dafny.UnicodeSeqOfUtf8Bytes("hau")
		var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_1877_v71), 0))
		_ = _index331
		(_1877_v71).ArraySet1(_1878_v72, (_index331).Int())
		var _1879_v73 _dafny.Int
		_ = _1879_v73
		_1879_v73 = _dafny.IntOfInt64(437)
		var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_1877_v71), 0))
		_ = _index332
		(_1877_v71).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1878_v72, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg115 _dafny.Int) interface{} {
				return coer115(arg115)
			}
		}(func(_1880_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('y')
		}))), _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm27(_1879_v73, _this.F14(), globalState), _1878_v72)), (_index332).Int())
		var _1881_v74 _dafny.Set
		_ = _1881_v74
		_1881_v74 = _dafny.SetOf(_1879_v73)
		var _hi7 _dafny.Int = _1879_v73
		_ = _hi7
		for _1882_i9 := (_1881_v74).Cardinality(); _1882_i9.Cmp(_hi7) < 0; _1882_i9 = _1882_i9.Plus(_dafny.One) {
			var _1883_v75 D5
			_ = _1883_v75
			_1883_v75 = Companion_D5_.Create_DC18_(!(true))
			var _1884_v76 _dafny.Map
			_ = _1884_v76
			_1884_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_i9, _1883_v75)
			_1884_v76 = (_1884_v76).Update(((_1881_v74).Cardinality()).Minus(_1879_v73), _1883_v75)
			var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1784_v0), 0))
			_ = _index333
			(_1784_v0).ArraySet1((func() bool {
				if _this.F14() {
					return _this.F14()
				}
				return !(_this.F14())
			})(), (_index333).Int())
			var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1784_v0), 0))
			_ = _index334
			(_1784_v0).ArraySet1(_this.F14(), (_index334).Int())
			var _1885_v77 _dafny.CodePoint
			_ = _1885_v77
			_1885_v77 = _dafny.CodePoint('h')
			var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1784_v0), 0))
			_ = _index335
			var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1784_v0), 0))
			_ = _index336
			var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_1877_v71), 0))
			_ = _index337
			var _rhs307 bool = true
			_ = _rhs307
			var _rhs308 bool = _this.F14()
			_ = _rhs308
			var _rhs309 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1878_v72, (Companion_Default___.SafeIndex(_1879_v73, _dafny.IntOfUint32((_1878_v72).Cardinality()))).Uint32(), _1885_v77), (_1877_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_1877_v71), 0))).Int()).(_dafny.Sequence))
			_ = _rhs309
			var _lhs260 _dafny.Array = _1784_v0
			_ = _lhs260
			var _lhs261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(686), _dafny.ArrayLen((_1784_v0), 0))
			_ = _lhs261
			var _lhs262 _dafny.Array = _1784_v0
			_ = _lhs262
			var _lhs263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(129), _dafny.ArrayLen((_1784_v0), 0))
			_ = _lhs263
			var _lhs264 _dafny.Array = _1877_v71
			_ = _lhs264
			var _lhs265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(202), _dafny.ArrayLen((_1877_v71), 0))
			_ = _lhs265
			(_lhs260).ArraySet1(_rhs307, (_lhs261).Int())
			(_lhs262).ArraySet1(_rhs308, (_lhs263).Int())
			(_lhs264).ArraySet1(_rhs309, (_lhs265).Int())
			var _1886_v78 _dafny.Sequence
			_ = _1886_v78
			_1886_v78 = _dafny.SeqOf(true, _this.F14())
			(globalState).F3 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1886_v78, _1886_v78), _1886_v78)).Cardinality())
			var _1887_v79 _dafny.Map
			_ = _1887_v79
			_1887_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1882_i9, _1882_i9)
			var _1888_v80 D9
			_ = _1888_v80
			_1888_v80 = Companion_D9_.Create_DC29_(_1887_v79)
			var _1889_v81 D9
			_ = _1889_v81
			_1889_v81 = Companion_D9_.Create_DC31_(_1888_v80)
			var _1890_v82 D9
			_ = _1890_v82
			_1890_v82 = Companion_D9_.Create_DC31_(_1889_v81)
			var _1891_v83 _dafny.Sequence
			_ = _1891_v83
			_1891_v83 = _dafny.SeqOf(_1890_v82)
			var _1892_v84 _dafny.Set
			_ = _1892_v84
			_1892_v84 = _dafny.SetOf(_1891_v83)
			_1892_v84 = Companion_Default___.Fm49(_1879_v73, globalState)
		}
		var _1893_v85 _dafny.Array
		_ = _1893_v85
		var _len0_53 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_53
		var _nw306 _dafny.Array
		_ = _nw306
		if _len0_53.Cmp(_dafny.Zero) == 0 {
			_nw306 = _dafny.NewArray(_len0_53)
		} else {
			var _init53 func(_dafny.Int) _dafny.Int = (func(_1894_v73 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1895_i10 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1895_i10, _1894_v73)
				}
			})(_1879_v73)
			_ = _init53
			var _element0_53 = _init53(_dafny.Zero)
			_ = _element0_53
			_nw306 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
			(_nw306).ArraySet1(_element0_53, 0)
			var _nativeLen0_53 = (_len0_53).Int()
			_ = _nativeLen0_53
			for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
				(_nw306).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
			}
		}
		_1893_v85 = _nw306
		var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))
		_ = _index338
		(_1893_v85).ArraySet1(_1879_v73, (_index338).Int())
		var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))
		_ = _index339
		(_1893_v85).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if _this.F14() {
				return _1879_v73
			}
			return _1879_v73
		})(), (func() _dafny.Int {
			if _this.F14() {
				return _dafny.IntOfInt64(284)
			}
			return _dafny.IntOfInt64(269)
		})())), (_index339).Int())
		var _1896_v86 _dafny.Sequence
		_ = _1896_v86
		_1896_v86 = _dafny.SeqOf((_1893_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))).Int()).(_dafny.Int))
		var _1897_v87 _dafny.Map
		_ = _1897_v87
		_1897_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1896_v86, _dafny.SeqOf((_1893_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))).Int()).(_dafny.Int)))).Cardinality()), _this.F14())
		_1897_v87 = (_1897_v87).Update(_1879_v73, ((_1893_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(510)) <= 0)
		r0 = (_1893_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))).Int()).(_dafny.Int)
		r1 = (_1879_v73).Times(_1879_v73)
		var _1898_v88 _dafny.MultiSet
		_ = _1898_v88
		_1898_v88 = _dafny.MultiSetOf(_1879_v73, _dafny.IntOfUint32((_1878_v72).Cardinality()), _1879_v73, Companion_Default___.SafeModuloInt(_1879_v73, (_1893_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(239), _dafny.ArrayLen((_1893_v85), 0))).Int()).(_dafny.Int)))
		r2 = _1898_v88
		var _1899_v90 _dafny.Sequence
		_ = _1899_v90
		_1899_v90 = _dafny.SeqOf(_1897_v87)
		var _1900_v91 D7
		_ = _1900_v91
		_1900_v91 = Companion_D7_.Create_DC21_(func() _dafny.Map {
			var _coll74 = _dafny.NewMapBuilder()
			_ = _coll74
			for _iter80 := _dafny.Iterate((_1899_v90).Elements()); ; {
				_compr_74, _ok80 := _iter80()
				if !_ok80 {
					break
				}
				var _1901_v89 _dafny.Map
				_1901_v89 = interface{}(_compr_74).(_dafny.Map)
				if _dafny.Companion_Sequence_.Contains(_1899_v90, _1901_v89) {
					_coll74.Add(_1901_v89, _1879_v73)
				}
			}
			return _coll74.ToMap()
		}())
		r3 = func(_source18 D7) bool {
			if _source18.Is_DC22() {
				var _1902___mcc_h3 _dafny.Array = _source18.Get_().(D7_DC22).Cf37
				_ = _1902___mcc_h3
				var _1903___mcc_h4 _dafny.Sequence = _source18.Get_().(D7_DC22).Cf38
				_ = _1903___mcc_h4
				var _1904___mcc_h5 _dafny.Sequence = _source18.Get_().(D7_DC22).Cf39
				_ = _1904___mcc_h5
				var _1905___mcc_h6 _dafny.Array = _source18.Get_().(D7_DC22).Cf40
				_ = _1905___mcc_h6
				var _1906___mcc_h7 D5 = _source18.Get_().(D7_DC22).Cf41
				_ = _1906___mcc_h7
				var _1907_cf41 D5 = _1906___mcc_h7
				_ = _1907_cf41
				var _1908_cf40 _dafny.Array = _1905___mcc_h6
				_ = _1908_cf40
				var _1909_cf39 _dafny.Sequence = _1904___mcc_h5
				_ = _1909_cf39
				var _1910_cf38 _dafny.Sequence = _1903___mcc_h4
				_ = _1910_cf38
				var _1911_cf37 _dafny.Array = _1902___mcc_h3
				_ = _1911_cf37
				return !(_this.F14())
			} else if _source18.Is_DC23() {
				var _1912___mcc_h8 bool = _source18.Get_().(D7_DC23).Cf42
				_ = _1912___mcc_h8
				var _1913___mcc_h9 _dafny.Int = _source18.Get_().(D7_DC23).Cf43
				_ = _1913___mcc_h9
				var _1914___mcc_h10 bool = _source18.Get_().(D7_DC23).Cf44
				_ = _1914___mcc_h10
				var _1915_cf44 bool = _1914___mcc_h10
				_ = _1915_cf44
				var _1916_cf43 _dafny.Int = _1913___mcc_h9
				_ = _1916_cf43
				var _1917_cf42 bool = _1912___mcc_h8
				_ = _1917_cf42
				return _1915_cf44
			} else {
				var _1918___mcc_h11 _dafny.Map = _source18.Get_().(D7_DC21).Cf36
				_ = _1918___mcc_h11
				var _1919_cf36 _dafny.Map = _1918___mcc_h11
				_ = _1919_cf36
				return _this.F14()
			}
		}(_1900_v91)
		return r0, r1, r2, r3
	}
}

// End of class C6

// Definition of class C7
type C7 struct {
	dummy byte
}

func New_C7_() *C7 {
	_this := C7{}

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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C7{}

func (_this *C7) Ctor__() {
	{
	}
}
func (_this *C7) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kko"), _dafny.UnicodeSeqOfUtf8Bytes("ddx"))).Cardinality())
	}
}
func (_this *C7) Fm10(p0 _dafny.Map, p1 _dafny.MultiSet, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C7) M3(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _hi8 _dafny.Int = p0
		_ = _hi8
		for _1920_i0 := p0; _1920_i0.Cmp(_hi8) < 0; _1920_i0 = _1920_i0.Plus(_dafny.One) {
			var _1921_v0 D1
			_ = _1921_v0
			_1921_v0 = Companion_D1_.Create_DC3_(_dafny.IntOfInt64(-181))
			var _1922_v1 _dafny.Map
			_ = _1922_v1
			_1922_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1921_v0, _1920_i0)
			var _1923_v2 _dafny.Sequence
			_ = _1923_v2
			_1923_v2 = _dafny.SeqOf((_1922_v1).Cardinality(), p0)
			var _1924_v3 _dafny.CodePoint
			_ = _1924_v3
			_1924_v3 = _dafny.CodePoint('j')
			var _1925_v4 _dafny.Sequence
			_ = _1925_v4
			_1925_v4 = _dafny.SeqOf(_1924_v3)
			var _1926_v5 _dafny.Array
			_ = _1926_v5
			var _nwElement0_58 _dafny.Sequence = _1923_v2
			_ = _nwElement0_58
			var _nw307 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(15))
			_ = _nw307
			(_nw307).ArraySet1(_nwElement0_58, 0)
			(_nw307).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1923_v2, _1923_v2), 1)
			(_nw307).ArraySet1(_1923_v2, 2)
			(_nw307).ArraySet1(_1923_v2, 3)
			(_nw307).ArraySet1(_1923_v2, 4)
			(_nw307).ArraySet1(_dafny.Companion_Sequence_.Update(_1923_v2, (Companion_Default___.SafeIndex(_1920_i0, _dafny.IntOfUint32((_1923_v2).Cardinality()))).Uint32(), (_1921_v0).Dtor_cf3()), 5)
			(_nw307).ArraySet1(_dafny.SeqOf(_dafny.IntOfUint32((_1925_v4).Cardinality()), p0), 6)
			(_nw307).ArraySet1(_1923_v2, 7)
			(_nw307).ArraySet1(_1923_v2, 8)
			(_nw307).ArraySet1((func() _dafny.Sequence {
				if true {
					return _1923_v2
				}
				return _1923_v2
			})(), 9)
			(_nw307).ArraySet1(_dafny.Companion_Sequence_.Update(_1923_v2, (Companion_Default___.SafeIndex(_1920_i0, _dafny.IntOfUint32((_1923_v2).Cardinality()))).Uint32(), _1920_i0), 10)
			(_nw307).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1923_v2, _dafny.SeqOf(p0, p0, _dafny.IntOfUint32((_1925_v4).Cardinality()))), 11)
			(_nw307).ArraySet1(_1923_v2, 12)
			(_nw307).ArraySet1(_1923_v2, 13)
			(_nw307).ArraySet1(_1923_v2, 14)
			_1926_v5 = _nw307
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1926_v5), 0))
			_ = _index340
			(_1926_v5).ArraySet1(_dafny.Companion_Sequence_.Update(_1923_v2, (Companion_Default___.SafeIndex(_1920_i0, _dafny.IntOfUint32((_1923_v2).Cardinality()))).Uint32(), p0), (_index340).Int())
			var _1927_v6 _dafny.Array
			_ = _1927_v6
			var _nw308 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw308
			_1927_v6 = _nw308
			var _1928_v7 bool
			_ = _1928_v7
			_1928_v7 = true
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1927_v6), 0))
			_ = _index341
			(_1927_v6).ArraySet1(_1928_v7, (_index341).Int())
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1926_v5), 0))
			_ = _index342
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1927_v6), 0))
			_ = _index343
			var _rhs310 bool = _1928_v7
			_ = _rhs310
			var _rhs311 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1923_v2, _1923_v2), _dafny.Companion_Sequence_.Concatenate(_1923_v2, _1923_v2))
			_ = _rhs311
			var _rhs312 bool = !(_1928_v7)
			_ = _rhs312
			var _lhs266 *GlobalState = globalState
			_ = _lhs266
			var _lhs267 _dafny.Array = _1926_v5
			_ = _lhs267
			var _lhs268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1926_v5), 0))
			_ = _lhs268
			var _lhs269 _dafny.Array = _1927_v6
			_ = _lhs269
			var _lhs270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1927_v6), 0))
			_ = _lhs270
			_lhs266.F9 = _rhs310
			(_lhs267).ArraySet1(_rhs311, (_lhs268).Int())
			(_lhs269).ArraySet1(_rhs312, (_lhs270).Int())
			var _1929_v8 _dafny.Map
			_ = _1929_v8
			_1929_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1924_v3)
			var _1930_v9 _dafny.MultiSet
			_ = _1930_v9
			_1930_v9 = _dafny.MultiSetOf(_1929_v8)
			var _1931_v10 _dafny.Array
			_ = _1931_v10
			var _nw309 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw309
			_1931_v10 = _nw309
			var _1932_v11 _dafny.Map
			_ = _1932_v11
			_1932_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer116 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg116 _dafny.Int) interface{} {
					return coer116(arg116)
				}
			}((func(_1933_v6 _dafny.Array) func(_dafny.Int) _dafny.Int {
				return func(_1934_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf((_1933_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1933_v6), 0))).Int()).(bool))).Cardinality())
				}
			})(_1927_v6))), _1924_v3)
			var _1935_v12 _dafny.Map
			_ = _1935_v12
			_1935_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1928_v7, _1927_v6)
			var _rhs313 _dafny.MultiSet = (_1930_v9).Difference(_1930_v9)
			_ = _rhs313
			var _rhs314 bool = !(_1928_v7)
			_ = _rhs314
			var _rhs315 _dafny.Array = (func() _dafny.Array {
				if (_1935_v12).Contains((_1927_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1927_v6), 0))).Int()).(bool)) {
					return (_1935_v12).Get((_1927_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_1927_v6), 0))).Int()).(bool)).(_dafny.Array)
				}
				return _1927_v6
			})()
			_ = _rhs315
			var _rhs316 _dafny.Array = _1931_v10
			_ = _rhs316
			var _rhs317 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1926_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(789), _dafny.ArrayLen((_1926_v5), 0))).Int()).(_dafny.Sequence), _1924_v3)
			_ = _rhs317
			var _lhs271 *GlobalState = globalState
			_ = _lhs271
			_1930_v9 = _rhs313
			_1928_v7 = _rhs314
			_lhs271.F2 = _rhs315
			_1931_v10 = _rhs316
			_1932_v11 = _rhs317
			_1925_v4 = _dafny.Companion_Sequence_.Concatenate(_1925_v4, _dafny.UnicodeSeqOfUtf8Bytes("reyaln"))
			(globalState).F9 = (_1920_i0).Cmp(_1920_i0) != 0
		}
		var _hi9 _dafny.Int = p0
		_ = _hi9
		for _1936_i2 := _dafny.IntOfInt64(-126); _1936_i2.Cmp(_hi9) < 0; _1936_i2 = _1936_i2.Plus(_dafny.One) {
			var _1937_v13 _dafny.Array
			_ = _1937_v13
			var _nw310 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(2))
			_ = _nw310
			_1937_v13 = _nw310
			var _1938_v14 D1
			_ = _1938_v14
			_1938_v14 = Companion_D1_.Create_DC4_(_dafny.IntOfInt64(573), p0)
			var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1937_v13), 0))
			_ = _index344
			(_1937_v13).ArraySet1(func(_pat_let62_0 D1) D1 {
				return func(_1939_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let63_0 _dafny.Int) D1 {
						return func(_1940_dt__update_hcf4_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC4_(_1940_dt__update_hcf4_h0, (_1939_dt__update__tmp_h0).Dtor_cf5())
						}(_pat_let63_0)
					}(_1936_i2)
				}(_pat_let62_0)
			}(_1938_v14), (_index344).Int())
			var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_1937_v13), 0))
			_ = _index345
			(_1937_v13).ArraySet1(_1938_v14, (_index345).Int())
			var _1941_v15 bool
			_ = _1941_v15
			_1941_v15 = true
			var _1942_v16 *C0
			_ = _1942_v16
			var _nw311 *C0 = New_C0_()
			_ = _nw311
			_nw311.Ctor__(_1941_v15, _1941_v15)
			_1942_v16 = _nw311
			var _1943_v17 _dafny.Map
			_ = _1943_v17
			_1943_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-189), p0)
			(globalState).F9 = !(_1943_v17).Equals(Companion_Default___.Fm14(globalState))
			var _1944_v18 _dafny.CodePoint
			_ = _1944_v18
			_1944_v18 = _dafny.CodePoint('o')
			var _1945_v19 _dafny.Map
			_ = _1945_v19
			_1945_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1944_v18, (_dafny.Zero).Minus(_1936_i2))
			var _1946_v20 _dafny.Sequence
			_ = _1946_v20
			_1946_v20 = _dafny.SeqOf(_1936_i2, _1936_i2)
			var _1947_v21 _dafny.Set
			_ = _1947_v21
			_1947_v21 = _dafny.SetOf(_1946_v20, _1946_v20)
			var _1948_v22 _dafny.Array
			_ = _1948_v22
			var _len0_54 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_54
			var _nw312 _dafny.Array
			_ = _nw312
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw312 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) bool = (func(_1949_v16 *C0) func(_dafny.Int) bool {
					return func(_1950_i3 _dafny.Int) bool {
						return (_1949_v16).F24()
					}
				})(_1942_v16)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw312 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw312).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw312).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_1948_v22 = _nw312
			(globalState).F2 = (func() _dafny.Array {
				if !((func() bool {
					if Companion_Default___.Fm23(p0, _1941_v15, _1945_v19, _1947_v21, globalState) {
						return _1941_v15
					}
					return false
				})()) {
					return _1948_v22
				}
				return _1948_v22
			})()
		}
		var _1951_v23 _dafny.MultiSet
		_ = _1951_v23
		_1951_v23 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0), _dafny.IntOfInt64(704))
		(globalState).F9 = !(!(((_1951_v23).Intersection(_1951_v23)).IsDisjointFrom(_1951_v23)))
		var _1952_i4 _dafny.Int
		_ = _1952_i4
		_1952_i4 = _dafny.Zero
		{
			for (p0).Cmp(p0) != 0 {
				{
					if (_1952_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1952_i4 = (_1952_i4).Plus(_dafny.One)
					var _1953_v24 _dafny.Array
					_ = _1953_v24
					var _nw313 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
					_ = _nw313
					_1953_v24 = _nw313
					var _1954_v25 _dafny.Sequence
					_ = _1954_v25
					_1954_v25 = _dafny.SeqOf(_1953_v24, _1953_v24, _1953_v24, _1953_v24, _1953_v24)
					var _1955_v26 bool
					_ = _1955_v26
					_1955_v26 = false
					_1954_v25 = (func() _dafny.Sequence {
						if _1955_v26 {
							return _1954_v25
						}
						return _1954_v25
					})()
					_1955_v26 = _1955_v26
					var _1956_v27 _dafny.Sequence
					_ = _1956_v27
					_1956_v27 = _dafny.SeqOf(_1955_v26)
					var _1957_v28 _dafny.Sequence
					_ = _1957_v28
					_1957_v28 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1956_v27, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1956_v27).Cardinality()))).Uint32(), !(_1955_v26))).Cardinality()), _dafny.IntOfInt64(-573), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1955_v26, _1955_v26)).Cardinality(), (_dafny.Zero).Minus(p0))
					var _1958_v29 _dafny.Sequence
					_ = _1958_v29
					_1958_v29 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(915))).Uint32(), func(coer117 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg117 _dafny.Int) interface{} {
							return coer117(arg117)
						}
					}(func(_1959_i5 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(506)
					})), _1957_v28)
					var _source19 _dafny.Sequence = _1958_v29
					_ = _source19
					var _1960___mcc_h0 _dafny.Sequence = _source19
					_ = _1960___mcc_h0
					var _1961_cf56 _dafny.Sequence = _1960___mcc_h0
					_ = _1961_cf56
					var _1962_v30 _dafny.Array
					_ = _1962_v30
					var _nw314 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(8))
					_ = _nw314
					_1962_v30 = _nw314
					var _1963_v31 _dafny.Sequence
					_ = _1963_v31
					_1963_v31 = _dafny.SeqOf(_1962_v30)
					var _1964_v32 D16
					_ = _1964_v32
					_1964_v32 = Companion_D16_.Create_DC51_(_1962_v30)
					_1963_v31 = _dafny.SeqOf(_1962_v30, (_1964_v32).Dtor_cf83())
					var _1965_v33 _dafny.Map
					_ = _1965_v33
					_1965_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1958_v29, p0)
					var _1966_v34 _dafny.MultiSet
					_ = _1966_v34
					_1966_v34 = _dafny.MultiSetOf(_1965_v33)
					(globalState).F9 = (_1966_v34).Contains(_1965_v33)
					var _1967_v35 _dafny.Set
					_ = _1967_v35
					_1967_v35 = _dafny.SetOf(_1955_v26)
					(globalState).F6 = (_1967_v35).Intersection(Companion_Default___.Fm41(p0, p0, globalState))
					var _1968_v36 _dafny.Map
					_ = _1968_v36
					_1968_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(_1955_v26))
					var _1969_v37 _dafny.Sequence
					_ = _1969_v37
					_1969_v37 = _dafny.UnicodeSeqOfUtf8Bytes("o")
					var _1970_v38 _dafny.MultiSet
					_ = _1970_v38
					_1970_v38 = _dafny.MultiSetOf(_1955_v26)
					var _1971_v39 *C3
					_ = _1971_v39
					var _nw315 *C3 = New_C3_()
					_ = _nw315
					_nw315.Ctor__(p0, _1969_v37, _1955_v26, _1970_v38)
					_1971_v39 = _nw315
					var _1972_v40 _dafny.Sequence
					_ = _1972_v40
					_1972_v40 = _dafny.SeqOf(_1971_v39, _1971_v39, _1971_v39, _1971_v39, _1971_v39)
					var _1973_v41 _dafny.Map
					_ = _1973_v41
					_1973_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1972_v40, _1970_v38)
					var _1974_v42 _dafny.Map
					_ = _1974_v42
					_1974_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1955_v26, _1973_v41)
					var _1975_v43 _dafny.Map
					_ = _1975_v43
					_1975_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1956_v27).Cardinality()), _1973_v41)
					_1968_v36 = (_1968_v36).Update((_dafny.Zero).Minus((((func() _dafny.Map {
						if (_1974_v42).Contains(_1955_v26) {
							return (_1974_v42).Get(_1955_v26).(_dafny.Map)
						}
						return (func() _dafny.Map {
							if (_1975_v43).Contains(p0) {
								return (_1975_v43).Get(p0).(_dafny.Map)
							}
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1972_v40, _1970_v38)
						})()
					})()).Update(_1972_v40, _1970_v38)).Cardinality()), _1955_v26)
					var _1976_v44 _dafny.Set
					_ = _1976_v44
					_1976_v44 = _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(866), p0))
					var _1977_v45 D13
					_ = _1977_v45
					_1977_v45 = Companion_D13_.Create_DC41_(_1976_v44)
					var _source20 D13 = _1977_v45
					_ = _source20
					if _source20.Is_DC42() {
						var _1978___mcc_h1 _dafny.CodePoint = _source20.Get_().(D13_DC42).Cf68
						_ = _1978___mcc_h1
						var _1979___mcc_h2 bool = _source20.Get_().(D13_DC42).Cf69
						_ = _1979___mcc_h2
						var _1980_cf69 bool = _1979___mcc_h2
						_ = _1980_cf69
						var _1981_cf68 _dafny.CodePoint = _1978___mcc_h1
						_ = _1981_cf68
						var _1982_v46 _dafny.Set
						_ = _1982_v46
						_1982_v46 = _dafny.SetOf(_1980_cf69, _1980_cf69, _1980_cf69)
						(globalState).F6 = (func() _dafny.Set {
							if _1955_v26 {
								return _1982_v46
							}
							return _1982_v46
						})()
						var _1983_v47 _dafny.Array
						_ = _1983_v47
						var _nw316 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
						_ = _nw316
						_1983_v47 = _nw316
						var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1983_v47), 0))
						_ = _index346
						(_1983_v47).ArraySet1((_this).Fm9(p0, globalState), (_index346).Int())
						var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1983_v47), 0))
						_ = _index347
						(_1983_v47).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1954_v25, _dafny.Companion_Sequence_.Concatenate(_1954_v25, _1954_v25))).Cardinality()), (_index347).Int())
						var _1984_v48 _dafny.Map
						_ = _1984_v48
						_1984_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1980_cf69, _dafny.Companion_Sequence_.Concatenate(_1956_v27, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1955_v26), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(_1955_v26)).Cardinality()))).Uint32(), _1980_cf69)))
						var _1985_v49 _dafny.Map
						_ = _1985_v49
						_1985_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), (_1983_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1983_v47), 0))).Int()).(_dafny.Int))
						var _1986_v50 _dafny.Map
						_ = _1986_v50
						_1986_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1980_cf69, _dafny.IntOfInt64(202))
						var _1987_v51 _dafny.Sequence
						_ = _1987_v51
						_1987_v51 = _dafny.UnicodeSeqOfUtf8Bytes("jrpyg")
						_1984_v48 = (_1984_v48).Update(Companion_Default___.Fm23((_1983_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1983_v47), 0))).Int()).(_dafny.Int), false, _1985_v49, _dafny.SetOf(_1957_v28, _1957_v28, _1957_v28), globalState), (func() _dafny.Sequence {
							if (_1984_v48).Contains(_1955_v26) {
								return (_1984_v48).Get(_1955_v26).(_dafny.Sequence)
							}
							return _dafny.SeqOf(_1955_v26, _1980_cf69, !((_this).Fm10(_1986_v50, _1951_v23, false, _1987_v51, globalState)))
						})())
						var _1988_v52 _dafny.Sequence
						_ = _1988_v52
						_1988_v52 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("sct"))
						var _1989_v53 D4
						_ = _1989_v53
						_1989_v53 = Companion_D4_.Create_DC13_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(326))).Uint32(), func(coer118 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg118 _dafny.Int) interface{} {
								return coer118(arg118)
							}
						}((func(_1990_cf68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1991_i8 _dafny.Int) _dafny.CodePoint {
								return _1990_cf68
							}
						})(_1981_cf68))), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-776))).Uint32(), func(coer119 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg119 _dafny.Int) interface{} {
								return coer119(arg119)
							}
						}((func(_1992_cf68 _dafny.CodePoint, _1993_cf69 bool, _1994_p0 _dafny.Int, _1995_v47 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
							return func(_1996_i9 _dafny.Int) _dafny.CodePoint {
								return (Companion_D13_.Create_DC42_(_1992_cf68, (Companion_D9_.Create_DC30_(_1992_cf68, _1993_cf69, _1994_p0, _dafny.SetOf((_1995_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1995_v47), 0))).Int()).(_dafny.Int)), (_1995_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1995_v47), 0))).Int()).(_dafny.Int))).Dtor_cf51())).Dtor_cf68()
							}
						})(_1981_cf68, _1980_cf69, p0, _1983_v47)))).Cardinality()), _1955_v26, _1953_v24)
						var _1997_v54 _dafny.Array
						_ = _1997_v54
						var _nwElement0_59 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_1988_v52).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_1988_v52).Cardinality()))).Uint32()).(_dafny.Sequence), _1987_v51)
						_ = _nwElement0_59
						var _nw317 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(23))
						_ = _nw317
						(_nw317).ArraySet1(_nwElement0_59, 0)
						(_nw317).ArraySet1(_1987_v51, 1)
						(_nw317).ArraySet1(_1987_v51, 2)
						(_nw317).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1987_v51, _1987_v51), 3)
						(_nw317).ArraySet1(_1987_v51, 4)
						(_nw317).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1987_v51, _1987_v51), 5)
						(_nw317).ArraySet1(_1987_v51, 6)
						(_nw317).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1987_v51, _1987_v51), 7)
						(_nw317).ArraySet1(_dafny.Companion_Sequence_.Update(_1987_v51, (Companion_Default___.SafeIndex((_1983_v47).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_1983_v47), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1987_v51).Cardinality()))).Uint32(), _1981_cf68), 8)
						(_nw317).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(218))).Uint32(), func(coer120 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg120 _dafny.Int) interface{} {
								return coer120(arg120)
							}
						}((func(_1998_cf68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1999_i6 _dafny.Int) _dafny.CodePoint {
								return _1998_cf68
							}
						})(_1981_cf68))), 9)
						(_nw317).ArraySet1(_1987_v51, 10)
						(_nw317).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(804))).Uint32(), func(coer121 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg121 _dafny.Int) interface{} {
								return coer121(arg121)
							}
						}((func(_2000_cf68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2001_i7 _dafny.Int) _dafny.CodePoint {
								return _2000_cf68
							}
						})(_1981_cf68))), 11)
						(_nw317).ArraySet1(_1987_v51, 12)
						(_nw317).ArraySet1(_1987_v51, 13)
						(_nw317).ArraySet1((_1989_v53).Dtor_cf22(), 14)
						(_nw317).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(325))).Uint32(), func(coer122 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg122 _dafny.Int) interface{} {
								return coer122(arg122)
							}
						}((func(_2002_cf68 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_2003_i10 _dafny.Int) _dafny.CodePoint {
								return _2002_cf68
							}
						})(_1981_cf68))), 15)
						(_nw317).ArraySet1(_1987_v51, 16)
						(_nw317).ArraySet1(_1987_v51, 17)
						(_nw317).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("abhvu"), 18)
						(_nw317).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1987_v51, _dafny.UnicodeSeqOfUtf8Bytes("u")), 19)
						(_nw317).ArraySet1(_1987_v51, 20)
						(_nw317).ArraySet1((func() _dafny.Sequence {
							if _1980_cf69 {
								return _1987_v51
							}
							return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(390))).Uint32(), func(coer123 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg123 _dafny.Int) interface{} {
									return coer123(arg123)
								}
							}(func(_2004_i11 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('w')
							}))
						})(), 21)
						(_nw317).ArraySet1(_1987_v51, 22)
						_1997_v54 = _nw317
						var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_1997_v54), 0))
						_ = _index348
						(_1997_v54).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("i"), (_index348).Int())
						var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_1997_v54), 0))
						_ = _index349
						(_1997_v54).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1987_v51, _dafny.UnicodeSeqOfUtf8Bytes("k")), (_index349).Int())
					} else if _source20.Is_DC41() {
						var _2005___mcc_h3 _dafny.Set = _source20.Get_().(D13_DC41).Cf67
						_ = _2005___mcc_h3
						var _2006_cf67 _dafny.Set = _2005___mcc_h3
						_ = _2006_cf67
						var _2007_v55 _dafny.Sequence
						_ = _2007_v55
						_2007_v55 = _dafny.UnicodeSeqOfUtf8Bytes("gv")
						var _2008_v56 _dafny.Map
						_ = _2008_v56
						_2008_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_2007_v55).Cardinality()))
						(globalState).F3 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
							if _1955_v26 {
								return p0
							}
							return p0
						})(), ((_2008_v56).Merge((func() _dafny.Map {
							var _coll75 = _dafny.NewMapBuilder()
							_ = _coll75
							for _iter81 := _dafny.Iterate((_1951_v23).Elements()); ; {
								_compr_75, _ok81 := _iter81()
								if !_ok81 {
									break
								}
								var _2009_v57 _dafny.Int
								_2009_v57 = interface{}(_compr_75).(_dafny.Int)
								if (_1951_v23).Contains(_2009_v57) {
									_coll75.Add(Companion_Default___.SafeModuloInt(_2009_v57, _dafny.IntOfUint32((_2007_v55).Cardinality())), p0)
								}
							}
							return _coll75.ToMap()
						}()).Update(p0, p0))).Cardinality())
						var _2010_v58 _dafny.CodePoint
						_ = _2010_v58
						_2010_v58 = _dafny.CodePoint('f')
						var _2011_v59 _dafny.Array
						_ = _2011_v59
						var _nw318 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
						_ = _nw318
						_2011_v59 = _nw318
						var _2012_v60 _dafny.Map
						_ = _2012_v60
						_2012_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2011_v59, _1955_v26)
						var _2013_v61 _dafny.Set
						_ = _2013_v61
						_2013_v61 = _dafny.SetOf(Companion_Default___.Fm0(_1955_v26, _dafny.IntOfInt64(418), (_2012_v60).Cardinality(), _1951_v23, globalState), _dafny.IntOfUint32((_2007_v55).Cardinality()))
						var _2014_v62 _dafny.Set
						_ = _2014_v62
						_2014_v62 = _dafny.SetOf(Companion_D9_.Create_DC30_(_2010_v58, _1955_v26, p0, _2013_v61, p0))
						var _2015_v63 _dafny.Map
						_ = _2015_v63
						_2015_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2010_v58, p0)
						var _2016_v64 D0
						_ = _2016_v64
						_2016_v64 = Companion_D0_.Create_DC1_(_1955_v26)
						var _rhs318 bool = (_2014_v62).Contains(Companion_D9_.Create_DC30_(_2010_v58, _1955_v26, p0, _2013_v61, p0))
						_ = _rhs318
						var _rhs319 _dafny.Array = (func() _dafny.Array {
							if Companion_Default___.Fm23(p0, _1955_v26, _2015_v63, Companion_Default___.Fm50(p0, (_2016_v64).Dtor_cf1(), p0, p0, globalState), globalState) {
								return _1953_v24
							}
							return _1953_v24
						})()
						_ = _rhs319
						var _lhs272 *GlobalState = globalState
						_ = _lhs272
						_1955_v26 = _rhs318
						_lhs272.F2 = _rhs319
						var _2017_v65 _dafny.Array
						_ = _2017_v65
						var _nwElement0_60 _dafny.Array = _1953_v24
						_ = _nwElement0_60
						var _nw319 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(13))
						_ = _nw319
						(_nw319).ArraySet1(_nwElement0_60, 0)
						(_nw319).ArraySet1(_1953_v24, 1)
						(_nw319).ArraySet1(_1953_v24, 2)
						(_nw319).ArraySet1(_1953_v24, 3)
						(_nw319).ArraySet1(_1953_v24, 4)
						(_nw319).ArraySet1(_1953_v24, 5)
						(_nw319).ArraySet1(_1953_v24, 6)
						(_nw319).ArraySet1(_1953_v24, 7)
						(_nw319).ArraySet1(_1953_v24, 8)
						(_nw319).ArraySet1(_1953_v24, 9)
						(_nw319).ArraySet1(_1953_v24, 10)
						(_nw319).ArraySet1(_1953_v24, 11)
						(_nw319).ArraySet1(_1953_v24, 12)
						_2017_v65 = _nw319
						var _2018_v66 _dafny.Array
						_ = _2018_v66
						_2018_v66 = _2017_v65
						var _2019_v67 _dafny.Array
						_ = _2019_v67
						var _nwElement0_61 _dafny.Array = (func() _dafny.Array {
							if _1955_v26 {
								return _2017_v65
							}
							return _2017_v65
						})()
						_ = _nwElement0_61
						var _nw320 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(3))
						_ = _nw320
						(_nw320).ArraySet1(_nwElement0_61, 0)
						(_nw320).ArraySet1((_2018_v66), 1)
						(_nw320).ArraySet1(_2017_v65, 2)
						_2019_v67 = _nw320
						var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_2019_v67), 0))
						_ = _index350
						(_2019_v67).ArraySet1((func() _dafny.Array {
							if _1955_v26 {
								return _2017_v65
							}
							return _2017_v65
						})(), (_index350).Int())
						var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_2019_v67), 0))
						_ = _index351
						(_2019_v67).ArraySet1(_2017_v65, (_index351).Int())
						var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2011_v59), 0))
						_ = _index352
						(_2011_v59).ArraySet1(p0, (_index352).Int())
						var _2020_v68 _dafny.Map
						_ = _2020_v68
						_2020_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.UnicodeSeqOfUtf8Bytes("itn"))
						var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_2011_v59), 0))
						_ = _index353
						(_2011_v59).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-772), (_2020_v68).Cardinality()), (_index353).Int())
					} else {
						var _2021___mcc_h4 D13 = _source20.Get_().(D13_DC43).Cf70
						_ = _2021___mcc_h4
						var _2022_cf70 D13 = _2021___mcc_h4
						_ = _2022_cf70
						var _2023_v69 *C0
						_ = _2023_v69
						var _nw321 *C0 = New_C0_()
						_ = _nw321
						_nw321.Ctor__(_1955_v26, !(_1955_v26))
						_2023_v69 = _nw321
						(globalState).F9 = (_2023_v69).F24()
						(globalState).F9 = ((_dafny.IntOfInt64(53)).Minus(_dafny.IntOfInt64(175))).Cmp(p0) > 0
						var _2024_v70 _dafny.Array
						_ = _2024_v70
						var _nw322 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
						_ = _nw322
						_2024_v70 = _nw322
						var _2025_v71 _dafny.Array
						_ = _2025_v71
						var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(10))
						_ = _nw323
						_2025_v71 = _nw323
						var _2026_v72 _dafny.Map
						_ = _2026_v72
						_2026_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1955_v26, (_2023_v69).F23())
						var _rhs320 _dafny.Array = _2025_v71
						_ = _rhs320
						var _rhs321 _dafny.Array = (func() _dafny.Array {
							if (func() bool {
								if (_2026_v72).Contains(_1955_v26) {
									return (_2026_v72).Get(_1955_v26).(bool)
								}
								return _1955_v26
							})() {
								return (func() _dafny.Array {
									if false {
										return _2024_v70
									}
									return _2024_v70
								})()
							}
							return _2024_v70
						})()
						_ = _rhs321
						var _lhs273 *GlobalState = globalState
						_ = _lhs273
						_lhs273.F11 = _rhs320
						_2024_v70 = _rhs321
					}
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _2027_v73 _dafny.Sequence
		_ = _2027_v73
		_2027_v73 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		var _2028_v74 D0
		_ = _2028_v74
		_2028_v74 = Companion_D0_.Create_DC2_(Companion_Default___.Fm20(p0, _dafny.IntOfUint32((_2027_v73).Cardinality()), p0, globalState))
		var _2029_v75 D0
		_ = _2029_v75
		_2029_v75 = Companion_D0_.Create_DC2_(_2028_v74)
		var _pat_let_tv44 = p0
		_ = _pat_let_tv44
		if func(_source21 D0) bool {
			if _source21.Is_DC1() {
				var _2030___mcc_h5 bool = _source21.Get_().(D0_DC1).Cf1
				_ = _2030___mcc_h5
				var _2031_cf1 bool = _2030___mcc_h5
				_ = _2031_cf1
				return !(!(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(234))).Uint32(), func(coer124 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg124 _dafny.Int) interface{} {
						return coer124(arg124)
					}
				}(func(_2032_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()), _pat_let_tv44), _dafny.IntOfInt64(952))))
			} else if _source21.Is_DC0() {
				var _2033___mcc_h6 bool = _source21.Get_().(D0_DC0).Cf0
				_ = _2033___mcc_h6
				var _2034_cf0 bool = _2033___mcc_h6
				_ = _2034_cf0
				return _2034_cf0
			} else {
				var _2035___mcc_h7 D0 = _source21.Get_().(D0_DC2).Cf2
				_ = _2035___mcc_h7
				var _2036_cf2 D0 = _2035___mcc_h7
				_ = _2036_cf2
				return !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(false)))
			}
		}(_2029_v75) {
			var _2037_v76 bool
			_ = _2037_v76
			_2037_v76 = false
			var _2038_v77 _dafny.Map
			_ = _2038_v77
			_2038_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2027_v73, _2037_v76)
			_2038_v77 = _2038_v77
			(globalState).F9 = (_2037_v76) || (_2037_v76)
			var _2039_v78 _dafny.MultiSet
			_ = _2039_v78
			_2039_v78 = _dafny.MultiSetOf(_2037_v76, _2037_v76)
			var _2040_v79 *C5
			_ = _2040_v79
			var _nw324 *C5 = New_C5_()
			_ = _nw324
			_nw324.Ctor__((_dafny.Zero).Minus(p0), (_dafny.IntOfUint32((_2027_v73).Cardinality())).Cmp(p0) <= 0, _2039_v78)
			_2040_v79 = _nw324
			var _2041_v80 _dafny.Map
			_ = _2041_v80
			_2041_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _2039_v78)
			var _2042_v81 _dafny.Map
			_ = _2042_v81
			_2042_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
				if (_2041_v80).Contains(!(_2037_v76)) {
					return (_2041_v80).Get(!(_2037_v76)).(_dafny.MultiSet)
				}
				return Companion_Default___.Fm33(_2037_v76, globalState)
			})(), (func() bool {
				if _2037_v76 {
					return _2037_v76
				}
				return _2037_v76
			})())
			_2042_v81 = (_2042_v81).Update(_2039_v78, false)
			_2037_v76 = (Companion_Default___.SafeModuloInt((_2040_v79).F18(), p0)).Cmp(_dafny.IntOfUint32((_2027_v73).Cardinality())) < 0
		} else {
			(globalState).F3 = p0
			var _2043_v82 _dafny.CodePoint
			_ = _2043_v82
			_2043_v82 = _dafny.CodePoint('l')
			var _2044_v83 D2
			_ = _2044_v83
			_2044_v83 = Companion_D2_.Create_DC9_(true, p0, _2043_v82)
			(globalState).F9 = (_2044_v83).Dtor_cf13()
			var _2045_v84 bool
			_ = _2045_v84
			_2045_v84 = true
			var _2046_v85 _dafny.Map
			_ = _2046_v85
			_2046_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(409))).Uint32(), func(coer125 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg125 _dafny.Int) interface{} {
					return coer125(arg125)
				}
			}((func(_2047_v82 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2048_i13 _dafny.Int) _dafny.CodePoint {
					return _2047_v82
				}
			})(_2043_v82))), _2045_v84)
			_2046_v85 = ((_2046_v85).Merge(_2046_v85)).Merge(_2046_v85)
			var _2049_v86 _dafny.MultiSet
			_ = _2049_v86
			_2049_v86 = _dafny.MultiSetOf(_2045_v84, _2045_v84)
			var _2050_v87 *C4
			_ = _2050_v87
			var _nw325 *C4 = New_C4_()
			_ = _nw325
			_nw325.Ctor__(_2045_v84, !(_2045_v84), (_2049_v86).Difference(_2049_v86))
			_2050_v87 = _nw325
			(globalState).F9 = (_2050_v87).F19()
		}
		var _2051_v88 bool
		_ = _2051_v88
		_2051_v88 = true
		if _2051_v88 {
			var _2052_v89 _dafny.MultiSet
			_ = _2052_v89
			_2052_v89 = _dafny.MultiSetOf(false)
			var _2053_v90 _dafny.Map
			_ = _2053_v90
			_2053_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2051_v88, _2051_v88)
			var _2054_v91 _dafny.Array
			_ = _2054_v91
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_55
			var _nw326 _dafny.Array
			_ = _nw326
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw326 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) bool = (func(_2055_v88 bool) func(_dafny.Int) bool {
					return func(_2056_i14 _dafny.Int) bool {
						return _2055_v88
					}
				})(_2051_v88)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw326 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw326).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw326).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_2054_v91 = _nw326
			var _2057_v92 D11
			_ = _2057_v92
			_2057_v92 = Companion_D11_.Create_DC34_(_2052_v89, _dafny.MultiSetOf((_2053_v90).Cardinality()), _2054_v91, Companion_D1_.Create_DC5_(_2051_v88, false))
			var _2058_v93 _dafny.MultiSet
			_ = _2058_v93
			_2058_v93 = _dafny.MultiSetOf((_2057_v92).Dtor_cf60())
			var _2059_v94 _dafny.CodePoint
			_ = _2059_v94
			_2059_v94 = _dafny.CodePoint('x')
			var _2060_v95 _dafny.Map
			_ = _2060_v95
			_2060_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Minus((_2058_v93).Cardinality()), _2059_v94)
			var _2061_v96 _dafny.Sequence
			_ = _2061_v96
			_2061_v96 = _dafny.SeqOf(p0, (_2052_v89).Cardinality())
			_2060_v95 = (_2060_v95).Update(_dafny.IntOfUint32((_2061_v96).Cardinality()), _2059_v94)
			var _2062_v97 *C0
			_ = _2062_v97
			var _nw327 *C0 = New_C0_()
			_ = _nw327
			_nw327.Ctor__(_2051_v88, _2051_v88)
			_2062_v97 = _nw327
			var _2063_v98 _dafny.Map
			_ = _2063_v98
			_2063_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2062_v97)
			_2063_v98 = _2063_v98
			var _2064_v99 D14
			_ = _2064_v99
			_2064_v99 = Companion_D14_.Create_DC45_(_2053_v90, _2062_v97, p0)
			var _2065_v100 D14
			_ = _2065_v100
			_2065_v100 = Companion_D14_.Create_DC46_(_2064_v99)
			var _pat_let_tv45 = _2064_v99
			_ = _pat_let_tv45
			var _pat_let_tv46 = _2064_v99
			_ = _pat_let_tv46
			var _2066_v101 _dafny.Array
			_ = _2066_v101
			var _nwElement0_62 D14 = _2065_v100
			_ = _nwElement0_62
			var _nw328 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(26))
			_ = _nw328
			(_nw328).ArraySet1(_nwElement0_62, 0)
			(_nw328).ArraySet1(_2065_v100, 1)
			(_nw328).ArraySet1(func(_pat_let64_0 D14) D14 {
				return func(_2067_dt__update__tmp_h1 D14) D14 {
					return func(_pat_let65_0 D14) D14 {
						return func(_2068_dt__update_hcf75_h0 D14) D14 {
							return Companion_D14_.Create_DC46_(_2068_dt__update_hcf75_h0)
						}(_pat_let65_0)
					}(_pat_let_tv45)
				}(_pat_let64_0)
			}(_2065_v100), 2)
			(_nw328).ArraySet1(_2065_v100, 3)
			(_nw328).ArraySet1(_2065_v100, 4)
			(_nw328).ArraySet1(_2065_v100, 5)
			(_nw328).ArraySet1(_2065_v100, 6)
			(_nw328).ArraySet1(_2065_v100, 7)
			(_nw328).ArraySet1(_2065_v100, 8)
			(_nw328).ArraySet1(_2065_v100, 9)
			(_nw328).ArraySet1(_2065_v100, 10)
			(_nw328).ArraySet1(_2065_v100, 11)
			(_nw328).ArraySet1(_2065_v100, 12)
			(_nw328).ArraySet1(_2065_v100, 13)
			(_nw328).ArraySet1(_2065_v100, 14)
			(_nw328).ArraySet1(_2065_v100, 15)
			(_nw328).ArraySet1(_2065_v100, 16)
			(_nw328).ArraySet1(Companion_D14_.Create_DC46_(_2064_v99), 17)
			(_nw328).ArraySet1(_2065_v100, 18)
			(_nw328).ArraySet1(_2065_v100, 19)
			(_nw328).ArraySet1(_2065_v100, 20)
			(_nw328).ArraySet1((func() D14 {
				if _2051_v88 {
					return _2065_v100
				}
				return _2065_v100
			})(), 21)
			(_nw328).ArraySet1(_2065_v100, 22)
			(_nw328).ArraySet1(_2065_v100, 23)
			(_nw328).ArraySet1(func(_pat_let66_0 D14) D14 {
				return func(_2069_dt__update__tmp_h2 D14) D14 {
					return func(_pat_let67_0 D14) D14 {
						return func(_2070_dt__update_hcf75_h1 D14) D14 {
							return Companion_D14_.Create_DC46_(_2070_dt__update_hcf75_h1)
						}(_pat_let67_0)
					}(_pat_let_tv46)
				}(_pat_let66_0)
			}(_2065_v100), 24)
			(_nw328).ArraySet1(_2065_v100, 25)
			_2066_v101 = _nw328
			var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2066_v101), 0))
			_ = _index354
			(_2066_v101).ArraySet1(_2065_v100, (_index354).Int())
			var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2066_v101), 0))
			_ = _index355
			var _rhs322 bool = (_2062_v97).F23()
			_ = _rhs322
			var _rhs323 D14 = (func() D14 {
				if (_2062_v97).F23() {
					return _2065_v100
				}
				return _2065_v100
			})()
			_ = _rhs323
			var _rhs324 _dafny.Sequence = Companion_Default___.Fm27(p0, _2051_v88, globalState)
			_ = _rhs324
			var _rhs325 _dafny.Int = p0
			_ = _rhs325
			var _lhs274 *GlobalState = globalState
			_ = _lhs274
			var _lhs275 _dafny.Array = _2066_v101
			_ = _lhs275
			var _lhs276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_2066_v101), 0))
			_ = _lhs276
			var _lhs277 *GlobalState = globalState
			_ = _lhs277
			_lhs274.F9 = _rhs322
			(_lhs275).ArraySet1(_rhs323, (_lhs276).Int())
			_2027_v73 = _rhs324
			_lhs277.F3 = _rhs325
			var _2071_v102 _dafny.Sequence
			_ = _2071_v102
			_2071_v102 = _dafny.SeqOf((_2062_v97).F23())
			var _2072_v103 D2
			_ = _2072_v103
			_2072_v103 = Companion_D2_.Create_DC7_(_2071_v102)
			if _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true), _dafny.Companion_Sequence_.Update((_2072_v103).Dtor_cf9(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_2072_v103).Dtor_cf9()).Cardinality()))).Uint32(), (_2062_v97).F23())) {
				var _2073_v104 _dafny.Array
				_ = _2073_v104
				var _nw329 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
				_ = _nw329
				_2073_v104 = _nw329
				var _2074_v106 _dafny.Array
				_ = _2074_v106
				var _len0_56 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_56
				var _nw330 _dafny.Array
				_ = _nw330
				if _len0_56.Cmp(_dafny.Zero) == 0 {
					_nw330 = _dafny.NewArray(_len0_56)
				} else {
					var _init56 func(_dafny.Int) D11 = (func(_2075_v89 _dafny.MultiSet, _2076_v102 _dafny.Sequence) func(_dafny.Int) D11 {
						return func(_2077_i15 _dafny.Int) D11 {
							return Companion_D11_.Create_DC33_(func() _dafny.Set {
								var _coll76 = _dafny.NewBuilder()
								_ = _coll76
								for _iter82 := _dafny.Iterate((_dafny.SeqOf(_2075_v89, _2075_v89, _dafny.MultiSetFromSeq(_2076_v102))).Elements()); ; {
									_compr_76, _ok82 := _iter82()
									if !_ok82 {
										break
									}
									var _2078_v105 _dafny.MultiSet
									_2078_v105 = interface{}(_compr_76).(_dafny.MultiSet)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_2075_v89, _2075_v89, _dafny.MultiSetFromSeq(_2076_v102)), _2078_v105) {
										_coll76.Add(_2078_v105)
									}
								}
								return _coll76.ToSet()
							}())
						}
					})(_2052_v89, _2071_v102)
					_ = _init56
					var _element0_56 = _init56(_dafny.Zero)
					_ = _element0_56
					_nw330 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
					(_nw330).ArraySet1(_element0_56, 0)
					var _nativeLen0_56 = (_len0_56).Int()
					_ = _nativeLen0_56
					for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
						(_nw330).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
					}
				}
				_2074_v106 = _nw330
				var _2079_v107 _dafny.Map
				_ = _2079_v107
				_2079_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2059_v94, (_dafny.Zero).Minus(p0))
				var _2080_v108 _dafny.Set
				_ = _2080_v108
				_2080_v108 = _dafny.SetOf(_2061_v96)
				var _2081_v109 _dafny.Map
				_ = _2081_v109
				_2081_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2074_v106, Companion_Default___.Fm23(p0, (_2062_v97).F23(), _2079_v107, _2080_v108, globalState))
				var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2073_v104), 0))
				_ = _index356
				(_2073_v104).ArraySet1(_2081_v109, (_index356).Int())
				var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_2073_v104), 0))
				_ = _index357
				(_2073_v104).ArraySet1((_2081_v109).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2074_v106, (_2062_v97).F23())), (_index357).Int())
				(globalState).F9 = !((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yacasvvoi")).Cardinality())).Cmp(p0) > 0) || (_2051_v88)
				var _2082_v110 _dafny.Map
				_ = _2082_v110
				_2082_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(p0, p0), p0)
				_2082_v110 = (_2082_v110).Merge(_2082_v110)
				var _2083_v111 D13
				_ = _2083_v111
				_2083_v111 = Companion_D13_.Create_DC42_(_2059_v94, (_2062_v97).F24())
				var _2084_v112 D13
				_ = _2084_v112
				_2084_v112 = Companion_D13_.Create_DC43_(_2083_v111)
				var _2085_v113 D13
				_ = _2085_v113
				_2085_v113 = Companion_D13_.Create_DC43_(Companion_D13_.Create_DC43_(_2083_v111))
				var _2086_v114 D13
				_ = _2086_v114
				_2086_v114 = Companion_D13_.Create_DC43_(_2085_v113)
				var _2087_v115 D13
				_ = _2087_v115
				_2087_v115 = Companion_D13_.Create_DC43_(_2084_v112)
				var _2088_v116 _dafny.Set
				_ = _2088_v116
				_2088_v116 = _dafny.SetOf(p0)
				var _2089_v118 _dafny.Map
				_ = _2089_v118
				_2089_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf(p0))
				var _2090_v119 _dafny.Sequence
				_ = _2090_v119
				_2090_v119 = _dafny.SeqOf(_2027_v73)
				var _2091_v120 D14
				_ = _2091_v120
				_2091_v120 = Companion_D14_.Create_DC45_(_2053_v90, _2062_v97, _dafny.IntOfUint32(((_2090_v119).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2090_v119).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
				var _2092_v124 _dafny.Map
				_ = _2092_v124
				_2092_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2062_v97).F24())
				var _2093_v126 _dafny.Sequence
				_ = _2093_v126
				_2093_v126 = _dafny.SeqOf(_2092_v124, (func() _dafny.Map {
					var _coll77 = _dafny.NewMapBuilder()
					_ = _coll77
					for _iter83 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(206), _dafny.IntOfInt64(40))); ; {
						_compr_77, _ok83 := _iter83()
						if !_ok83 {
							break
						}
						var _2094_v125 _dafny.Int
						_2094_v125 = interface{}(_compr_77).(_dafny.Int)
						if ((_dafny.IntOfInt64(206)).Cmp(_2094_v125) <= 0) && ((_2094_v125).Cmp(_dafny.IntOfInt64(40)) < 0) {
							_coll77.Add(Companion_Default___.SafeDivisionInt(_2094_v125, _dafny.IntOfInt64(-866)), (_2062_v97).F23())
						}
					}
					return _coll77.ToMap()
				}()).Update(p0, true))
				var _2095_v128 _dafny.Array
				_ = _2095_v128
				var _nw331 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw331
				_2095_v128 = _nw331
				var _2096_v130 _dafny.Array
				_ = _2096_v130
				var _nwElement0_63 _dafny.Set = _2088_v116
				_ = _nwElement0_63
				var _nw332 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(19))
				_ = _nw332
				(_nw332).ArraySet1(_nwElement0_63, 0)
				(_nw332).ArraySet1(_2088_v116, 1)
				(_nw332).ArraySet1(func() _dafny.Set {
					var _coll78 = _dafny.NewBuilder()
					_ = _coll78
					for _iter84 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(778), _dafny.IntOfInt64(596))); ; {
						_compr_78, _ok84 := _iter84()
						if !_ok84 {
							break
						}
						var _2097_v117 _dafny.Int
						_2097_v117 = interface{}(_compr_78).(_dafny.Int)
						if ((_dafny.IntOfInt64(778)).Cmp(_2097_v117) <= 0) && ((_2097_v117).Cmp(_dafny.IntOfInt64(596)) < 0) {
							_coll78.Add((_2097_v117).Times(p0))
						}
					}
					return _coll78.ToSet()
				}(), 2)
				(_nw332).ArraySet1((func() _dafny.Set {
					if (_2062_v97).F23() {
						return (func() _dafny.Set {
							if (_2089_v118).Contains(((_2091_v120).Dtor_cf72()).Cardinality()) {
								return (_2089_v118).Get(((_2091_v120).Dtor_cf72()).Cardinality()).(_dafny.Set)
							}
							return _2088_v116
						})()
					}
					return _2088_v116
				})(), 3)
				(_nw332).ArraySet1(func() _dafny.Set {
					var _coll79 = _dafny.NewBuilder()
					_ = _coll79
					for _iter85 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(407), _dafny.IntOfInt64(-440))); ; {
						_compr_79, _ok85 := _iter85()
						if !_ok85 {
							break
						}
						var _2098_v121 _dafny.Int
						_2098_v121 = interface{}(_compr_79).(_dafny.Int)
						if ((_dafny.IntOfInt64(407)).Cmp(_2098_v121) <= 0) && ((_2098_v121).Cmp(_dafny.IntOfInt64(-440)) < 0) {
							_coll79.Add(Companion_Default___.SafeModuloInt(_2098_v121, p0))
						}
					}
					return _coll79.ToSet()
				}(), 4)
				(_nw332).ArraySet1(_2088_v116, 5)
				(_nw332).ArraySet1(_2088_v116, 6)
				(_nw332).ArraySet1(_2088_v116, 7)
				(_nw332).ArraySet1(func() _dafny.Set {
					var _coll80 = _dafny.NewBuilder()
					_ = _coll80
					for _iter86 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-243), _dafny.IntOfInt64(713))); ; {
						_compr_80, _ok86 := _iter86()
						if !_ok86 {
							break
						}
						var _2099_v122 _dafny.Int
						_2099_v122 = interface{}(_compr_80).(_dafny.Int)
						if ((_dafny.IntOfInt64(-243)).Cmp(_2099_v122) <= 0) && ((_2099_v122).Cmp(_dafny.IntOfInt64(713)) < 0) {
							_coll80.Add((_2099_v122).Times(p0))
						}
					}
					return _coll80.ToSet()
				}(), 8)
				(_nw332).ArraySet1((func() _dafny.Set {
					var _coll81 = _dafny.NewBuilder()
					_ = _coll81
					for _iter87 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(437), _dafny.IntOfInt64(-919))); ; {
						_compr_81, _ok87 := _iter87()
						if !_ok87 {
							break
						}
						var _2100_v123 _dafny.Int
						_2100_v123 = interface{}(_compr_81).(_dafny.Int)
						if ((_dafny.IntOfInt64(437)).Cmp(_2100_v123) <= 0) && ((_2100_v123).Cmp(_dafny.IntOfInt64(-919)) < 0) {
							_coll81.Add((_2100_v123).Minus((_dafny.SetOf((_2062_v97).F23())).Cardinality()))
						}
					}
					return _coll81.ToSet()
				}()).Difference(_2088_v116), 9)
				(_nw332).ArraySet1(Companion_Default___.Fm22(_dafny.IntOfUint32((_2027_v73).Cardinality()), _2093_v126, (_2062_v97).F23(), (_2062_v97).F24(), globalState), 10)
				(_nw332).ArraySet1(func() _dafny.Set {
					var _coll82 = _dafny.NewBuilder()
					_ = _coll82
					for _iter88 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-200), _dafny.IntOfInt64(104))); ; {
						_compr_82, _ok88 := _iter88()
						if !_ok88 {
							break
						}
						var _2101_v127 _dafny.Int
						_2101_v127 = interface{}(_compr_82).(_dafny.Int)
						if ((_dafny.IntOfInt64(-200)).Cmp(_2101_v127) <= 0) && ((_2101_v127).Cmp(_dafny.IntOfInt64(104)) < 0) {
							_coll82.Add((_2101_v127).Plus((_1951_v23).Cardinality()))
						}
					}
					return _coll82.ToSet()
				}(), 11)
				(_nw332).ArraySet1(_2088_v116, 12)
				(_nw332).ArraySet1(_2088_v116, 13)
				(_nw332).ArraySet1(_2088_v116, 14)
				(_nw332).ArraySet1(_dafny.SetOf(p0, _dafny.IntOfUint32((_2093_v126).Cardinality()), Companion_Default___.Fm0(_2051_v88, (_dafny.MultiSetOf(_2095_v128)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_2062_v97).F23(), (_2062_v97).F24())).Cardinality()), _1951_v23, globalState), p0, p0), 15)
				(_nw332).ArraySet1(func() _dafny.Set {
					var _coll83 = _dafny.NewBuilder()
					_ = _coll83
					for _iter89 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(311), _dafny.IntOfInt64(608))); ; {
						_compr_83, _ok89 := _iter89()
						if !_ok89 {
							break
						}
						var _2102_v129 _dafny.Int
						_2102_v129 = interface{}(_compr_83).(_dafny.Int)
						if ((_dafny.IntOfInt64(311)).Cmp(_2102_v129) <= 0) && ((_2102_v129).Cmp(_dafny.IntOfInt64(608)) < 0) {
							_coll83.Add((_2102_v129).Times(_dafny.IntOfUint32((_dafny.SeqOf((_2062_v97).F24(), (_2062_v97).F24())).Cardinality())))
						}
					}
					return _coll83.ToSet()
				}(), 16)
				(_nw332).ArraySet1(_2088_v116, 17)
				(_nw332).ArraySet1(_2088_v116, 18)
				_2096_v130 = _nw332
				var _pat_let_tv47 = _2086_v114
				_ = _pat_let_tv47
				var _rhs326 _dafny.Array = _2096_v130
				_ = _rhs326
				var _rhs327 bool = (_2062_v97).F23()
				_ = _rhs327
				var _rhs328 _dafny.Int = p0
				_ = _rhs328
				var _rhs329 D13 = func(_pat_let68_0 D13) D13 {
					return func(_2103_dt__update__tmp_h3 D13) D13 {
						return func(_pat_let69_0 D13) D13 {
							return func(_2104_dt__update_hcf70_h0 D13) D13 {
								return Companion_D13_.Create_DC43_(_2104_dt__update_hcf70_h0)
							}(_pat_let69_0)
						}(_pat_let_tv47)
					}(_pat_let68_0)
				}(Companion_Default___.Fm51(_2051_v88, (_2062_v97).F23(), _2051_v88, (_2092_v124).Cardinality(), globalState))
				_ = _rhs329
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				var _lhs280 *GlobalState = globalState
				_ = _lhs280
				_lhs278.F11 = _rhs326
				_lhs279.F9 = _rhs327
				_lhs280.F3 = _rhs328
				_2087_v115 = _rhs329
				_2027_v73 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-323))).Uint32(), func(coer126 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg126 _dafny.Int) interface{} {
						return coer126(arg126)
					}
				}((func(_2105_v73 _dafny.Sequence, _2106_v97 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_2107_i16 _dafny.Int) _dafny.CodePoint {
						return (_2105_v73).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_2106_v97).F24())).Cardinality()), _dafny.IntOfUint32((_2105_v73).Cardinality()))).Uint32()).(_dafny.CodePoint)
					}
				})(_2027_v73, _2062_v97)))
			} else {
				var _2108_v131 _dafny.Array
				_ = _2108_v131
				var _len0_57 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_57
				var _nw333 _dafny.Array
				_ = _nw333
				if _len0_57.Cmp(_dafny.Zero) == 0 {
					_nw333 = _dafny.NewArray(_len0_57)
				} else {
					var _init57 func(_dafny.Int) _dafny.Int = (func(_2109_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2110_i17 _dafny.Int) _dafny.Int {
							return (_2110_i17).Minus(_2109_p0)
						}
					})(p0)
					_ = _init57
					var _element0_57 = _init57(_dafny.Zero)
					_ = _element0_57
					_nw333 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
					(_nw333).ArraySet1(_element0_57, 0)
					var _nativeLen0_57 = (_len0_57).Int()
					_ = _nativeLen0_57
					for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
						(_nw333).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
					}
				}
				_2108_v131 = _nw333
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_2108_v131), 0))
				_ = _index358
				(_2108_v131).ArraySet1(p0, (_index358).Int())
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_2108_v131), 0))
				_ = _index359
				(_2108_v131).ArraySet1(p0, (_index359).Int())
				var _2111_v132 *C3
				_ = _2111_v132
				var _nw334 *C3 = New_C3_()
				_ = _nw334
				_nw334.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(845), p0), _dafny.Companion_Sequence_.Concatenate(_2027_v73, _2027_v73), false, (_2052_v89).Union(_2052_v89))
				_2111_v132 = _nw334
				var _2112_v134 D9
				_ = _2112_v134
				_2112_v134 = Companion_D9_.Create_DC30_(_2059_v94, (_2062_v97).F23(), _dafny.IntOfUint32((_2071_v102).Cardinality()), func() _dafny.Set {
					var _coll84 = _dafny.NewBuilder()
					_ = _coll84
					for _iter90 := _dafny.Iterate((_2061_v96).Elements()); ; {
						_compr_84, _ok90 := _iter90()
						if !_ok90 {
							break
						}
						var _2113_v133 _dafny.Int
						_2113_v133 = interface{}(_compr_84).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_2061_v96, _2113_v133) {
							_coll84.Add(Companion_Default___.SafeModuloInt(_2113_v133, _dafny.IntOfInt64(979)))
						}
					}
					return _coll84.ToSet()
				}(), (_2108_v131).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_2108_v131), 0))).Int()).(_dafny.Int))
				(globalState).F9 = (_2112_v134).Dtor_cf51()
				var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_2108_v131), 0))
				_ = _index360
				(_2108_v131).ArraySet1((_2108_v131).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(927), _dafny.ArrayLen((_2108_v131), 0))).Int()).(_dafny.Int), (_index360).Int())
				var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_2054_v91), 0))
				_ = _index361
				(_2054_v91).ArraySet1(((_2062_v97).F23()) && ((_2062_v97).F24()), (_index361).Int())
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_2054_v91), 0))
				_ = _index362
				(_2054_v91).ArraySet1((_2062_v97).F24(), (_index362).Int())
			}
			(globalState).F9 = (_2062_v97).F24()
		} else {
			(globalState).F9 = _2051_v88
			var _2114_v135 _dafny.Map
			_ = _2114_v135
			_2114_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), !(!(_2051_v88)))
			var _2115_v136 _dafny.Sequence
			_ = _2115_v136
			_2115_v136 = _dafny.SeqOf(_2114_v135, (_2114_v135).Merge(Companion_Default___.Fm46(_dafny.IntOfUint32((_2027_v73).Cardinality()), _2051_v88, globalState)), _2114_v135, _2114_v135, _2114_v135)
			(globalState).F3 = _dafny.IntOfUint32((_2115_v136).Cardinality())
			var _2116_v137 _dafny.Array
			_ = _2116_v137
			var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw335
			_2116_v137 = _nw335
			var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2116_v137), 0))
			_ = _index363
			(_2116_v137).ArraySet1(_2027_v73, (_index363).Int())
			var _2117_v138 _dafny.CodePoint
			_ = _2117_v138
			_2117_v138 = _dafny.CodePoint('e')
			var _2118_v139 _dafny.Array
			_ = _2118_v139
			var _len0_58 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_58
			var _nw336 _dafny.Array
			_ = _nw336
			if _len0_58.Cmp(_dafny.Zero) == 0 {
				_nw336 = _dafny.NewArray(_len0_58)
			} else {
				var _init58 func(_dafny.Int) bool = (func(_2119_v88 bool) func(_dafny.Int) bool {
					return func(_2120_i18 _dafny.Int) bool {
						return _2119_v88
					}
				})(_2051_v88)
				_ = _init58
				var _element0_58 = _init58(_dafny.Zero)
				_ = _element0_58
				_nw336 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
				(_nw336).ArraySet1(_element0_58, 0)
				var _nativeLen0_58 = (_len0_58).Int()
				_ = _nativeLen0_58
				for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
					(_nw336).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
				}
			}
			_2118_v139 = _nw336
			var _2121_v140 _dafny.Sequence
			_ = _2121_v140
			_2121_v140 = _dafny.SeqOf(_2118_v139)
			var _2122_v141 _dafny.MultiSet
			_ = _2122_v141
			_2122_v141 = _dafny.MultiSetOf((_2121_v140).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2121_v140).Cardinality()))).Uint32()).(_dafny.Array))
			var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2116_v137), 0))
			_ = _index364
			(_2116_v137).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(p0, _2051_v88, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm27(p0, _2051_v88, globalState)).Cardinality()))).Uint32(), _2117_v138), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-50), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(p0, _2051_v88, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm27(p0, _2051_v88, globalState)).Cardinality()))).Uint32(), _2117_v138)).Cardinality()))).Uint32(), _2117_v138), (Companion_Default___.SafeIndex((_2122_v141).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(p0, _2051_v88, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm27(p0, _2051_v88, globalState)).Cardinality()))).Uint32(), _2117_v138), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-50), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(p0, _2051_v88, globalState), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((Companion_Default___.Fm27(p0, _2051_v88, globalState)).Cardinality()))).Uint32(), _2117_v138)).Cardinality()))).Uint32(), _2117_v138)).Cardinality()))).Uint32(), _2117_v138), (_index364).Int())
			var _2123_v142 _dafny.Array
			_ = _2123_v142
			var _nw337 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
			_ = _nw337
			_2123_v142 = _nw337
			var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_2123_v142), 0))
			_ = _index365
			(_2123_v142).ArraySet1(p0, (_index365).Int())
			var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(4), _dafny.ArrayLen((_2123_v142), 0))
			_ = _index366
			(_2123_v142).ArraySet1(((_dafny.IntOfUint32(((_2116_v137).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2116_v137), 0))).Int()).(_dafny.Sequence)).Cardinality())).Plus(p0)).Minus(p0), (_index366).Int())
			_2117_v138 = _2117_v138
		}
	}
}
func (_this *C7) M4(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Set, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		(globalState).F9 = p0
		var _2124_v0 _dafny.Array
		_ = _2124_v0
		var _len0_59 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_59
		var _nw338 _dafny.Array
		_ = _nw338
		if _len0_59.Cmp(_dafny.Zero) == 0 {
			_nw338 = _dafny.NewArray(_len0_59)
		} else {
			var _init59 func(_dafny.Int) _dafny.Int = func(_2125_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_2125_i1, _dafny.IntOfInt64(150))
			}
			_ = _init59
			var _element0_59 = _init59(_dafny.Zero)
			_ = _element0_59
			_nw338 = _dafny.NewArrayFromExample(_element0_59, nil, _len0_59)
			(_nw338).ArraySet1(_element0_59, 0)
			var _nativeLen0_59 = (_len0_59).Int()
			_ = _nativeLen0_59
			for _i0_59 := 1; _i0_59 < _nativeLen0_59; _i0_59++ {
				(_nw338).ArraySet1(_init59(_dafny.IntOf(_i0_59)), _i0_59)
			}
		}
		_2124_v0 = _nw338
		for _iter91 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2124_v0), 0))); ; {
			_guard_loop_6, _ok91 := _iter91()
			if !_ok91 {
				break
			}
			var _2126_i0 _dafny.Int
			_2126_i0 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_2126_i0).Sign() != -1) && ((_2126_i0).Cmp(_dafny.ArrayLen((_2124_v0), 0)) < 0)) {
				(_2124_v0).ArraySet1((_2126_i0).Plus((Companion_D3_.Create_DC11_(_dafny.IntOfInt64(463), p2, p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)).Cardinality())).Dtor_cf18()), (_2126_i0).Int())
			}
		}
		var _2127_v1 _dafny.Set
		_ = _2127_v1
		_2127_v1 = _dafny.SetOf(p0, p0)
		var _2128_v2 _dafny.Map
		_ = _2128_v2
		_2128_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2127_v1).Union(Companion_Default___.Fm41(p2, p1, globalState)), p0)
		(globalState).F3 = (_2128_v2).Cardinality()
		var _2129_v3 _dafny.Array
		_ = _2129_v3
		var _nwElement0_64 bool = false
		_ = _nwElement0_64
		var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(2))
		_ = _nw339
		(_nw339).ArraySet1(_nwElement0_64, 0)
		(_nw339).ArraySet1(p0, 1)
		_2129_v3 = _nw339
		var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))
		_ = _index367
		(_2129_v3).ArraySet1(p0, (_index367).Int())
		var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))
		_ = _index368
		(_2129_v3).ArraySet1(p0, (_index368).Int())
		var _2130_v4 _dafny.Map
		_ = _2130_v4
		_2130_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_2129_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))).Int()).(bool))
		var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _index369
		(_2124_v0).ArraySet1((_dafny.Zero).Minus(p1), (_index369).Int())
		var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))
		_ = _index370
		var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _index371
		var _rhs330 _dafny.Map = _2130_v4
		_ = _rhs330
		var _rhs331 bool = (p1).Cmp((_dafny.Zero).Minus((p1).Minus(p2))) != 0
		_ = _rhs331
		var _rhs332 _dafny.Int = p1
		_ = _rhs332
		var _rhs333 _dafny.Int = p2
		_ = _rhs333
		var _lhs281 _dafny.Array = _2129_v3
		_ = _lhs281
		var _lhs282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))
		_ = _lhs282
		var _lhs283 _dafny.Array = _2124_v0
		_ = _lhs283
		var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _lhs284
		var _lhs285 *GlobalState = globalState
		_ = _lhs285
		_2130_v4 = _rhs330
		(_lhs281).ArraySet1(_rhs331, (_lhs282).Int())
		(_lhs283).ArraySet1(_rhs332, (_lhs284).Int())
		_lhs285.F3 = _rhs333
		var _2131_v5 _dafny.CodePoint
		_ = _2131_v5
		_2131_v5 = _dafny.CodePoint('c')
		var _2132_v6 _dafny.MultiSet
		_ = _2132_v6
		_2132_v6 = _dafny.MultiSetOf((_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(441))
		var _2133_v7 _dafny.MultiSet
		_ = _2133_v7
		_2133_v7 = _dafny.MultiSetOf(false)
		var _2134_v8 T0
		_ = _2134_v8
		var _nw340 *C5 = New_C5_()
		_ = _nw340
		_nw340.Ctor__((_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), p0, _2133_v7)
		_2134_v8 = _nw340
		var _2135_v9 _dafny.MultiSet
		_ = _2135_v9
		_2135_v9 = _dafny.MultiSetOf(_2134_v8, _2134_v8)
		var _2136_v10 _dafny.Sequence
		_ = _2136_v10
		_2136_v10 = _dafny.UnicodeSeqOfUtf8Bytes("nnbtloi")
		var _2137_v11 _dafny.Set
		_ = _2137_v11
		_2137_v11 = _dafny.SetOf((_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), Companion_Default___.Fm0((_2129_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))).Int()).(bool), (_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), (_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), (_2132_v6).Update((_2135_v9).Cardinality(), Companion_Default___.Abs(p1)), globalState), _dafny.IntOfUint32((_2136_v10).Cardinality()), p1)
		var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _index372
		var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _index373
		var _rhs334 bool = (_2129_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))).Int()).(bool)
		_ = _rhs334
		var _rhs335 _dafny.Int = Companion_Default___.SafeModuloInt((_2124_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))).Int()).(_dafny.Int), (_2137_v11).Cardinality())
		_ = _rhs335
		var _rhs336 _dafny.CodePoint = _2131_v5
		_ = _rhs336
		var _rhs337 _dafny.Int = (_dafny.Zero).Minus(p2)
		_ = _rhs337
		var _lhs286 *GlobalState = globalState
		_ = _lhs286
		var _lhs287 _dafny.Array = _2124_v0
		_ = _lhs287
		var _lhs288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _lhs288
		var _lhs289 _dafny.Array = _2124_v0
		_ = _lhs289
		var _lhs290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2124_v0), 0))
		_ = _lhs290
		_lhs286.F9 = _rhs334
		(_lhs287).ArraySet1(_rhs335, (_lhs288).Int())
		_2131_v5 = _rhs336
		(_lhs289).ArraySet1(_rhs337, (_lhs290).Int())
		var _2138_v12 _dafny.Sequence
		_ = _2138_v12
		_2138_v12 = _dafny.SeqOf(p0, true, true)
		r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf((_2129_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(262), _dafny.ArrayLen((_2129_v3), 0))).Int()).(bool)), _2138_v12)
		return r0
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	_f14 bool
	_f15 _dafny.MultiSet
	_f16 _dafny.Map
	_f17 _dafny.Sequence
}

func New_C8_() *C8 {
	_this := C8{}

	_this._f14 = false
	_this._f15 = _dafny.EmptyMultiSet
	_this._f16 = _dafny.EmptyMap
	_this._f17 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C8{}
var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) F14() bool {
	return _this._f14
}
func (_this *C8) F14_set_(value bool) {
	_this._f14 = value
}
func (_this *C8) F15() _dafny.MultiSet {
	return _this._f15
}
func (_this *C8) Ctor__(f16 _dafny.Map, f17 _dafny.Sequence, f14 bool, f15 _dafny.MultiSet) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C8) Fm2(p0 bool, p1 bool, globalState *GlobalState) bool {
	{
		return _this.F14()
	}
}
func (_this *C8) Fm3(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if _this.F14() {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-738))).Uint32(), func(coer127 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg127 _dafny.Int) interface{} {
						return coer127(arg127)
					}
				}(func(_2139_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				}))
			}
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(123))).Uint32(), func(coer128 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg128 _dafny.Int) interface{} {
					return coer128(arg128)
				}
			}(func(_2140_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('l')
			}))
		})(), _dafny.UnicodeSeqOfUtf8Bytes("jshowx"))
	}
}
func (_this *C8) Fm4(p0 _dafny.Map, p1 bool, globalState *GlobalState) bool {
	{
		return _this.F14()
	}
}
func (_this *C8) M0(globalState *GlobalState) (_dafny.Int, bool, _dafny.Map, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Map = _dafny.EmptyMap
		_ = r2
		var r3 bool = false
		_ = r3
		var _2141_v0 _dafny.Array
		_ = _2141_v0
		var _nwElement0_65 bool = !(_this.F14())
		_ = _nwElement0_65
		var _nw341 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(8))
		_ = _nw341
		(_nw341).ArraySet1(_nwElement0_65, 0)
		(_nw341).ArraySet1(_this.F14(), 1)
		(_nw341).ArraySet1(_this.F14(), 2)
		(_nw341).ArraySet1(_this.F14(), 3)
		(_nw341).ArraySet1(_this.F14(), 4)
		(_nw341).ArraySet1(_this.F14(), 5)
		(_nw341).ArraySet1(_this.F14(), 6)
		(_nw341).ArraySet1(_this.F14(), 7)
		_2141_v0 = _nw341
		var _2142_v1 _dafny.MultiSet
		_ = _2142_v1
		_2142_v1 = _dafny.MultiSetOf(_2141_v0, _2141_v0)
		if !(_2142_v1).Contains(_2141_v0) {
			var _2143_v2 _dafny.Set
			_ = _2143_v2
			_2143_v2 = _dafny.SetOf(!(_this.F14()), _this.F14(), _this.F14(), _this.F14())
			(globalState).F9 = !(_2143_v2).Equals(_2143_v2)
			if (func() bool {
				if _this.F14() {
					return true
				}
				return _this.F14()
			})() {
				r3 = !(_this.F14())
				var _2144_v3 _dafny.Sequence
				_ = _2144_v3
				_2144_v3 = _dafny.UnicodeSeqOfUtf8Bytes("okwamfa")
				r3 = (_this).Fm2(_this.F14(), _dafny.Companion_Sequence_.Equal(_2144_v3, _2144_v3), globalState)
				(globalState).F13 = _dafny.SetOf(_this.F14())
				var _2145_v4 _dafny.Int
				_ = _2145_v4
				_2145_v4 = _dafny.IntOfInt64(431)
				(globalState).F3 = (_dafny.Zero).Minus(_2145_v4)
				var _2146_v5 _dafny.Sequence
				_ = _2146_v5
				_2146_v5 = _dafny.SeqOf(_2145_v4)
				var _2147_v6 _dafny.Sequence
				_ = _2147_v6
				_2147_v6 = _dafny.SeqOf(_2146_v5, _2146_v5, _2146_v5)
				var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_2141_v0), 0))
				_ = _index374
				(_2141_v0).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("ahnxhipsh"), _2144_v3), (_index374).Int())
				var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_2141_v0), 0))
				_ = _index375
				var _rhs338 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if _this.F14() {
						return _2147_v6
					}
					return _2147_v6
				})(), _2147_v6)
				_ = _rhs338
				var _rhs339 bool = (func() bool {
					if ((_this).F15()).Equals((_this).F15()) {
						return _this.F14()
					}
					return _this.F14()
				})()
				_ = _rhs339
				var _lhs291 _dafny.Array = _2141_v0
				_ = _lhs291
				var _lhs292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_2141_v0), 0))
				_ = _lhs292
				_2147_v6 = _rhs338
				(_lhs291).ArraySet1(_rhs339, (_lhs292).Int())
			} else {
				var _2148_v7 _dafny.Array
				_ = _2148_v7
				var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
				_ = _nw342
				_2148_v7 = _nw342
				var _2149_v8 _dafny.Int
				_ = _2149_v8
				_2149_v8 = _dafny.IntOfInt64(726)
				var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))
				_ = _index376
				(_2148_v7).ArraySet1(Companion_Default___.Fm0(_this.F14(), _dafny.IntOfInt64(325), _2149_v8, _dafny.MultiSetOf(_2149_v8, _dafny.IntOfInt64(370), _2149_v8, _2149_v8, _2149_v8), globalState), (_index376).Int())
				var _2150_v9 _dafny.Sequence
				_ = _2150_v9
				_2150_v9 = _dafny.SeqOf(_dafny.IntOfInt64(-348))
				var _2151_v10 _dafny.Map
				_ = _2151_v10
				_2151_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2149_v8, _2149_v8)
				var _2152_v11 _dafny.Map
				_ = _2152_v11
				_2152_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2150_v9, _2151_v10)
				var _2153_v12 _dafny.MultiSet
				_ = _2153_v12
				_2153_v12 = _dafny.MultiSetOf(_2149_v8, _dafny.IntOfUint32((Companion_Default___.Fm5(_this.F14(), globalState)).Cardinality()), (_2152_v11).Cardinality())
				var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))
				_ = _index377
				(_2148_v7).ArraySet1(Companion_Default___.Fm0(_this.F14(), _2149_v8, _2149_v8, _2153_v12, globalState), (_index377).Int())
				var _2154_v13 D0
				_ = _2154_v13
				_2154_v13 = Companion_D0_.Create_DC1_(_this.F14())
				var _2155_v14 _dafny.Sequence
				_ = _2155_v14
				_2155_v14 = _dafny.SeqOf((_2154_v13).Dtor_cf1())
				var _2156_v15 _dafny.Map
				_ = _2156_v15
				_2156_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2150_v9, _dafny.IntOfUint32((_2155_v14).Cardinality()))
				var _2157_v16 _dafny.Map
				_ = _2157_v16
				_2157_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if _this.F14() {
						return (_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int)
					}
					return (func() _dafny.Int {
						if (_2156_v15).Contains(_2150_v9) {
							return (_2156_v15).Get(_2150_v9).(_dafny.Int)
						}
						return (_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int)
					})()
				})(), (_this.F14()) == (_this.F14()))
				var _2158_v17 D0
				_ = _2158_v17
				_2158_v17 = Companion_D0_.Create_DC0_(_this.F14())
				_2157_v16 = (_2157_v16).Update(Companion_Default___.SafeModuloInt(_2149_v8, _2149_v8), (_2158_v17).Dtor_cf0())
				var _2159_v18 _dafny.Map
				_ = _2159_v18
				_2159_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _2148_v7)
				var _2160_v19 _dafny.Set
				_ = _2160_v19
				_2160_v19 = _dafny.SetOf(_2149_v8, (_2159_v18).Cardinality(), (_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int), _2149_v8)
				var _2161_v20 _dafny.Sequence
				_ = _2161_v20
				_2161_v20 = _dafny.UnicodeSeqOfUtf8Bytes("h")
				var _2162_v21 _dafny.CodePoint
				_ = _2162_v21
				var _out59 _dafny.CodePoint
				_ = _out59
				_out59 = (_this).M2(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), (_2160_v19).Cardinality())).Cardinality(), _2149_v8), _dafny.Companion_Sequence_.Concatenate(_2161_v20, _dafny.UnicodeSeqOfUtf8Bytes("fmlv")), _this.F14(), !(_this.F14()) || (_this.F14()), globalState)
				_2162_v21 = _out59
				var _2163_v23 _dafny.Map
				_ = _2163_v23
				_2163_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("mkl"), true)
				var _2164_v24 _dafny.Map
				_ = _2164_v24
				_2164_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(false), ((func() _dafny.Map {
					var _coll85 = _dafny.NewMapBuilder()
					_ = _coll85
					for _iter92 := _dafny.Iterate((_2163_v23).Keys().Elements()); ; {
						_compr_85, _ok92 := _iter92()
						if !_ok92 {
							break
						}
						var _2165_v22 _dafny.Sequence
						_2165_v22 = interface{}(_compr_85).(_dafny.Sequence)
						if (_2163_v23).Contains(_2165_v22) {
							_coll85.Add(_2165_v22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(549))).Uint32(), func(coer129 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg129 _dafny.Int) interface{} {
									return coer129(arg129)
								}
							}((func(_2166_v21 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_2167_i0 _dafny.Int) _dafny.CodePoint {
									return _2166_v21
								}
							})(_2162_v21))))
						}
					}
					return _coll85.ToMap()
				}()).Cardinality()).Cmp(_dafny.IntOfUint32((_2161_v20).Cardinality())) > 0)
				(_this).F14_set_((func() bool {
					if (_2164_v24).Contains(_2158_v17) {
						return (_2164_v24).Get(_2158_v17).(bool)
					}
					return (_dafny.IntOfUint32((_2161_v20).Cardinality())).Cmp(_dafny.IntOfInt64(264)) > 0
				})())
				var _2168_v25 _dafny.Map
				_ = _2168_v25
				_2168_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.IntOfInt64(562))
				var _2169_v26 _dafny.Map
				_ = _2169_v26
				_2169_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int), _2168_v25)
				_2168_v25 = (func() _dafny.Map {
					if (_2169_v26).Contains((_dafny.Zero).Minus((_dafny.IntOfInt64(945)).Minus((_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int)))) {
						return (_2169_v26).Get((_dafny.Zero).Minus((_dafny.IntOfInt64(945)).Minus((_2148_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(68), _dafny.ArrayLen((_2148_v7), 0))).Int()).(_dafny.Int)))).(_dafny.Map)
					}
					return (_2168_v25).Update(_this.F14(), (_dafny.SetOf(_this.F14(), _this.F14())).Cardinality())
				})()
			}
			var _2170_v27 _dafny.Int
			_ = _2170_v27
			_2170_v27 = _dafny.IntOfInt64(-905)
			var _2171_v28 _dafny.Sequence
			_ = _2171_v28
			_2171_v28 = _dafny.SeqOf((func() _dafny.Int {
				if true {
					return _2170_v27
				}
				return _dafny.IntOfInt64(3)
			})(), _2170_v27)
			_2171_v28 = _2171_v28
			r0 = _2170_v27
			r3 = !((func() bool {
				if _this.F14() {
					return _this.F14()
				}
				return _this.F14()
			})())
		} else {
			var _2172_v29 _dafny.Sequence
			_ = _2172_v29
			_2172_v29 = _dafny.UnicodeSeqOfUtf8Bytes("vo")
			var _2173_v30 _dafny.Array
			_ = _2173_v30
			var _len0_60 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_60
			var _nw343 _dafny.Array
			_ = _nw343
			if _len0_60.Cmp(_dafny.Zero) == 0 {
				_nw343 = _dafny.NewArray(_len0_60)
			} else {
				var _init60 func(_dafny.Int) D0 = func(_2174_i1 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_this.F14())
				}
				_ = _init60
				var _element0_60 = _init60(_dafny.Zero)
				_ = _element0_60
				_nw343 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
				(_nw343).ArraySet1(_element0_60, 0)
				var _nativeLen0_60 = (_len0_60).Int()
				_ = _nativeLen0_60
				for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
					(_nw343).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
				}
			}
			_2173_v30 = _nw343
			var _2175_v31 _dafny.Array
			_ = _2175_v31
			var _nwElement0_66 _dafny.Array = _2173_v30
			_ = _nwElement0_66
			var _nw344 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(28))
			_ = _nw344
			(_nw344).ArraySet1(_nwElement0_66, 0)
			(_nw344).ArraySet1(_2173_v30, 1)
			(_nw344).ArraySet1(_2173_v30, 2)
			(_nw344).ArraySet1(_2173_v30, 3)
			(_nw344).ArraySet1(_2173_v30, 4)
			(_nw344).ArraySet1(_2173_v30, 5)
			(_nw344).ArraySet1(_2173_v30, 6)
			(_nw344).ArraySet1(_2173_v30, 7)
			(_nw344).ArraySet1(_2173_v30, 8)
			(_nw344).ArraySet1(_2173_v30, 9)
			(_nw344).ArraySet1(_2173_v30, 10)
			(_nw344).ArraySet1(_2173_v30, 11)
			(_nw344).ArraySet1(_2173_v30, 12)
			(_nw344).ArraySet1(_2173_v30, 13)
			(_nw344).ArraySet1(_2173_v30, 14)
			(_nw344).ArraySet1(_2173_v30, 15)
			(_nw344).ArraySet1(_2173_v30, 16)
			(_nw344).ArraySet1(_2173_v30, 17)
			(_nw344).ArraySet1(_2173_v30, 18)
			(_nw344).ArraySet1(_2173_v30, 19)
			(_nw344).ArraySet1(_2173_v30, 20)
			(_nw344).ArraySet1(_2173_v30, 21)
			(_nw344).ArraySet1(_2173_v30, 22)
			(_nw344).ArraySet1(_2173_v30, 23)
			(_nw344).ArraySet1(_2173_v30, 24)
			(_nw344).ArraySet1(_2173_v30, 25)
			(_nw344).ArraySet1(_2173_v30, 26)
			(_nw344).ArraySet1(_2173_v30, 27)
			_2175_v31 = _nw344
			var _2176_v32 _dafny.Map
			_ = _2176_v32
			_2176_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2172_v29, _2175_v31)
			_2176_v32 = (_2176_v32).Update(_2172_v29, _2175_v31)
			var _2177_v33 _dafny.Int
			_ = _2177_v33
			_2177_v33 = _dafny.IntOfInt64(411)
			var _2178_v34 _dafny.Sequence
			_ = _2178_v34
			_2178_v34 = _dafny.SeqOf(_2177_v33)
			var _2179_v35 _dafny.Set
			_ = _2179_v35
			_2179_v35 = _dafny.SetOf(_dafny.IntOfUint32((_2178_v34).Cardinality()), _2177_v33, _2177_v33)
			(globalState).F3 = ((Companion_Default___.Fm6(globalState)).Difference(_2179_v35)).Cardinality()
			r0 = _dafny.IntOfInt64(-607)
			var _2180_v36 _dafny.Array
			_ = _2180_v36
			var _nw345 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
			_ = _nw345
			_2180_v36 = _nw345
			var _2181_v37 _dafny.Map
			_ = _2181_v37
			_2181_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _this.F14())
			var _2182_v38 _dafny.Map
			_ = _2182_v38
			_2182_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if !(_this.F14()) {
					return Companion_Default___.Fm7(globalState)
				}
				return _2181_v37
			})(), _2177_v33)
			var _rhs340 _dafny.Array = _2180_v36
			_ = _rhs340
			var _rhs341 _dafny.Map = func() _dafny.Map {
				var _coll86 = _dafny.NewMapBuilder()
				_ = _coll86
				for _iter93 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(646))).Uint32(), func(coer130 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg130 _dafny.Int) interface{} {
						return coer130(arg130)
					}
				}((func(_2183_v37 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_2184_i2 _dafny.Int) _dafny.Map {
						return (_2183_v37).Update(false, _this.F14())
					}
				})(_2181_v37)))).Elements()); ; {
					_compr_86, _ok93 := _iter93()
					if !_ok93 {
						break
					}
					var _2185_v39 _dafny.Map
					_2185_v39 = interface{}(_compr_86).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(646))).Uint32(), func(coer131 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
						return func(arg131 _dafny.Int) interface{} {
							return coer131(arg131)
						}
					}((func(_2186_v37 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_2184_i2 _dafny.Int) _dafny.Map {
							return (_2186_v37).Update(false, _this.F14())
						}
					})(_2181_v37))), _2185_v39) {
						_coll86.Add(_2185_v39, _2177_v33)
					}
				}
				return _coll86.ToMap()
			}()
			_ = _rhs341
			var _rhs342 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2172_v29, _2172_v29)
			_ = _rhs342
			_2180_v36 = _rhs340
			_2182_v38 = _rhs341
			_2172_v29 = _rhs342
			var _2187_v40 _dafny.Array
			_ = _2187_v40
			var _nw346 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
			_ = _nw346
			_2187_v40 = _nw346
			var _2188_v41 _dafny.Map
			_ = _2188_v41
			_2188_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2177_v33, _2187_v40)
			_2188_v41 = (_2188_v41).Update(_2177_v33, _2187_v40)
		}
		var _2189_v42 _dafny.Set
		_ = _2189_v42
		_2189_v42 = _dafny.SetOf(_this.F14())
		var _2190_v43 _dafny.Int
		_ = _2190_v43
		_2190_v43 = _dafny.IntOfInt64(703)
		var _2191_v44 _dafny.CodePoint
		_ = _2191_v44
		_2191_v44 = _dafny.CodePoint('h')
		var _2192_v45 _dafny.Map
		_ = _2192_v45
		_2192_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2190_v43, _2191_v44)
		var _2193_v46 _dafny.Map
		_ = _2193_v46
		_2193_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_this.F14())).IsDisjointFrom(_2189_v42), _2192_v45)
		_2193_v46 = (_2193_v46).Update(_this.F14(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2190_v43, _2191_v44))
		var _2194_v47 _dafny.Sequence
		_ = _2194_v47
		_2194_v47 = _dafny.SeqOf(_this.F14(), _this.F14())
		var _2195_v48 _dafny.Map
		_ = _2195_v48
		_2195_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2194_v47, !(_this.F14()))
		_2195_v48 = (_2195_v48).Update(_2194_v47, (_this).Fm2(_this.F14(), _this.F14(), globalState))
		var _2196_v49 _dafny.Map
		_ = _2196_v49
		_2196_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), !(_this.F14()))
		var _2197_i3 _dafny.Int
		_ = _2197_i3
		_2197_i3 = _dafny.Zero
		{
			for (_this).Fm4(_2196_v49, _this.F14(), globalState) {
				{
					if (_2197_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_2197_i3 = (_2197_i3).Plus(_dafny.One)
					var _2198_v50 _dafny.Map
					_ = _2198_v50
					_2198_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _2190_v43)
					var _2199_v51 _dafny.MultiSet
					_ = _2199_v51
					_2199_v51 = _dafny.MultiSetOf(_dafny.IntOfInt64(501))
					(_this).F14_set_((_2199_v51).IsProperSubsetOf(_dafny.MultiSetOf((_2198_v50).Cardinality())))
					var _2200_v52 _dafny.Map
					_ = _2200_v52
					_2200_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2199_v51).Difference(_2199_v51), _this.F14())
					_2200_v52 = _2200_v52
					if (Companion_D0_.Create_DC0_((_2194_v47).Select((Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32((_2194_v47).Cardinality()))).Uint32()).(bool))).Dtor_cf0() {
						var _2201_v53 D1
						_ = _2201_v53
						_2201_v53 = Companion_D1_.Create_DC4_(_2190_v43, _2190_v43)
						(globalState).F9 = ((_dafny.Zero).Minus(_2190_v43)).Cmp((_2201_v53).Dtor_cf4()) == 0
						var _2202_v54 _dafny.Sequence
						_ = _2202_v54
						_2202_v54 = _dafny.UnicodeSeqOfUtf8Bytes("wfaglrkli")
						var _2203_v55 _dafny.Map
						_ = _2203_v55
						_2203_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.SeqOf(_2190_v43, _dafny.IntOfUint32((_2202_v54).Cardinality())))
						var _2204_v56 _dafny.Sequence
						_ = _2204_v56
						_2204_v56 = _dafny.SeqOf(_2190_v43, _2190_v43)
						var _2205_v57 _dafny.Map
						_ = _2205_v57
						_2205_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
							if (_2203_v55).Contains(_this.F14()) {
								return (_2203_v55).Get(_this.F14()).(_dafny.Sequence)
							}
							return _2204_v56
						})(), (Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32(((func() _dafny.Sequence {
							if (_2203_v55).Contains(_this.F14()) {
								return (_2203_v55).Get(_this.F14()).(_dafny.Sequence)
							}
							return _2204_v56
						})()).Cardinality()))).Uint32(), _2190_v43))
						_2203_v55 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _dafny.SeqOf(_2190_v43))).Merge(Companion_Default___.Fm8(_this.F14(), _this.F14(), (Companion_Default___.Fm6(globalState)).Cardinality(), _this.F14(), globalState))).Merge((func() _dafny.Map {
							if _this.F14() {
								return _2205_v57
							}
							return _2203_v55
						})())
						var _2206_v58 _dafny.Map
						_ = _2206_v58
						_2206_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_2196_v49).Contains(_this.F14()) {
								return (_2196_v49).Get(_this.F14()).(bool)
							}
							return _this.F14()
						})(), _2201_v53)
						_2206_v58 = (_2206_v58).Update((_2194_v47).Select((Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32((_2194_v47).Cardinality()))).Uint32()).(bool), Companion_D1_.Create_DC4_(_2190_v43, _2190_v43))
						(globalState).F9 = _this.F14()
						(globalState).F9 = false
					} else {
						var _2207_v59 _dafny.Array
						_ = _2207_v59
						var _nw347 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(15))
						_ = _nw347
						_2207_v59 = _nw347
						var _2208_v60 _dafny.Array
						_ = _2208_v60
						var _nwElement0_67 _dafny.Array = _2207_v59
						_ = _nwElement0_67
						var _nw348 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(28))
						_ = _nw348
						(_nw348).ArraySet1(_nwElement0_67, 0)
						(_nw348).ArraySet1(_2207_v59, 1)
						(_nw348).ArraySet1(_2207_v59, 2)
						(_nw348).ArraySet1(_2207_v59, 3)
						(_nw348).ArraySet1(_2207_v59, 4)
						(_nw348).ArraySet1(_2207_v59, 5)
						(_nw348).ArraySet1(_2207_v59, 6)
						(_nw348).ArraySet1(_2207_v59, 7)
						(_nw348).ArraySet1(_2207_v59, 8)
						(_nw348).ArraySet1(_2207_v59, 9)
						(_nw348).ArraySet1(_2207_v59, 10)
						(_nw348).ArraySet1((func() _dafny.Array {
							if _this.F14() {
								return _2207_v59
							}
							return _2207_v59
						})(), 11)
						(_nw348).ArraySet1(_2207_v59, 12)
						(_nw348).ArraySet1(_2207_v59, 13)
						(_nw348).ArraySet1(_2207_v59, 14)
						(_nw348).ArraySet1(_2207_v59, 15)
						(_nw348).ArraySet1(_2207_v59, 16)
						(_nw348).ArraySet1(_2207_v59, 17)
						(_nw348).ArraySet1(_2207_v59, 18)
						(_nw348).ArraySet1(_2207_v59, 19)
						(_nw348).ArraySet1(_2207_v59, 20)
						(_nw348).ArraySet1(_2207_v59, 21)
						(_nw348).ArraySet1(_2207_v59, 22)
						(_nw348).ArraySet1(_2207_v59, 23)
						(_nw348).ArraySet1(_2207_v59, 24)
						(_nw348).ArraySet1(_2207_v59, 25)
						(_nw348).ArraySet1(_2207_v59, 26)
						(_nw348).ArraySet1(_2207_v59, 27)
						_2208_v60 = _nw348
						var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2208_v60), 0))
						_ = _index378
						(_2208_v60).ArraySet1(_2207_v59, (_index378).Int())
						var _2209_v61 D1
						_ = _2209_v61
						_2209_v61 = Companion_D1_.Create_DC4_(_2190_v43, _2190_v43)
						var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2208_v60), 0))
						_ = _index379
						var _rhs343 _dafny.Array = (func() _dafny.Array {
							if (false) && (_this.F14()) {
								return _2207_v59
							}
							return _2207_v59
						})()
						_ = _rhs343
						var _rhs344 D1 = _2209_v61
						_ = _rhs344
						var _rhs345 _dafny.Int = _2190_v43
						_ = _rhs345
						var _lhs293 _dafny.Array = _2208_v60
						_ = _lhs293
						var _lhs294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_2208_v60), 0))
						_ = _lhs294
						(_lhs293).ArraySet1(_rhs343, (_lhs294).Int())
						_2209_v61 = _rhs344
						_2190_v43 = _rhs345
						var _2210_v62 _dafny.Sequence
						_ = _2210_v62
						_2210_v62 = _dafny.SeqOf(_2141_v0, _2141_v0, _2141_v0)
						_2210_v62 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_2210_v62, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-990), _dafny.IntOfUint32((_2210_v62).Cardinality()))).Uint32(), _2141_v0), (Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_2210_v62, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-990), _dafny.IntOfUint32((_2210_v62).Cardinality()))).Uint32(), _2141_v0)).Cardinality()))).Uint32(), _2141_v0)
						var _2211_v63 _dafny.Sequence
						_ = _2211_v63
						_2211_v63 = _dafny.SeqOf(_2194_v47)
						var _2212_v64 *C6
						_ = _2212_v64
						var _nw349 *C6 = New_C6_()
						_ = _nw349
						_nw349.Ctor__(_this.F14(), _dafny.MultiSetFromSeq((_2211_v63).Select((Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32((_2211_v63).Cardinality()))).Uint32()).(_dafny.Sequence)))
						_2212_v64 = _nw349
						var _2213_v65 _dafny.Sequence
						_ = _2213_v65
						_2213_v65 = _dafny.UnicodeSeqOfUtf8Bytes("ehllenung")
						_2213_v65 = _2213_v65
						var _2214_v66 *C7
						_ = _2214_v66
						var _nw350 *C7 = New_C7_()
						_ = _nw350
						_nw350.Ctor__()
						_2214_v66 = _nw350
						_2214_v66 = _2214_v66
					}
					var _2215_v67 _dafny.Array
					_ = _2215_v67
					var _nw351 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(22))
					_ = _nw351
					_2215_v67 = _nw351
					_2215_v67 = _2215_v67
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _2216_v69 _dafny.Array
		_ = _2216_v69
		var _len0_61 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_61
		var _nw352 _dafny.Array
		_ = _nw352
		if _len0_61.Cmp(_dafny.Zero) == 0 {
			_nw352 = _dafny.NewArray(_len0_61)
		} else {
			var _init61 func(_dafny.Int) _dafny.Map = (func(_2217_v44 _dafny.CodePoint, _2218_v43 _dafny.Int) func(_dafny.Int) _dafny.Map {
				return func(_2219_i4 _dafny.Int) _dafny.Map {
					return (func() _dafny.Map {
						if _this.F14() {
							return func() _dafny.Map {
								var _coll87 = _dafny.NewMapBuilder()
								_ = _coll87
								for _iter94 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-553), _dafny.IntOfInt64(967))); ; {
									_compr_87, _ok94 := _iter94()
									if !_ok94 {
										break
									}
									var _2220_v68 _dafny.Int
									_2220_v68 = interface{}(_compr_87).(_dafny.Int)
									if ((_dafny.IntOfInt64(-553)).Cmp(_2220_v68) <= 0) && ((_2220_v68).Cmp(_dafny.IntOfInt64(967)) < 0) {
										_coll87.Add((_2220_v68).Times(_2218_v43), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(636))).Uint32(), func(coer132 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
											return func(arg132 _dafny.Int) interface{} {
												return coer132(arg132)
											}
										}((func(_2221_v43 _dafny.Int) func(_dafny.Int) _dafny.Int {
											return func(_2222_i5 _dafny.Int) _dafny.Int {
												return _2221_v43
											}
										})(_2218_v43))))
									}
								}
								return _coll87.ToMap()
							}()
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_2217_v44, _dafny.CodePoint('v'))).Cardinality(), (_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer133 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg133 _dafny.Int) interface{} {
								return coer133(arg133)
							}
						}((func(_2223_v43 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2224_i6 _dafny.Int) _dafny.Int {
								return _2223_v43
							}
						})(_2218_v43)))))
					})()
				}
			})(_2191_v44, _2190_v43)
			_ = _init61
			var _element0_61 = _init61(_dafny.Zero)
			_ = _element0_61
			_nw352 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
			(_nw352).ArraySet1(_element0_61, 0)
			var _nativeLen0_61 = (_len0_61).Int()
			_ = _nativeLen0_61
			for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
				(_nw352).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
			}
		}
		_2216_v69 = _nw352
		var _2225_v70 *C1
		_ = _2225_v70
		var _nw353 *C1 = New_C1_()
		_ = _nw353
		_nw353.Ctor__()
		_2225_v70 = _nw353
		var _2226_v71 _dafny.Map
		_ = _2226_v71
		_2226_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F14(), _2225_v70)
		var _2227_v72 _dafny.Sequence
		_ = _2227_v72
		_2227_v72 = _dafny.SeqOf((_2226_v71).Cardinality())
		var _2228_v73 _dafny.Map
		_ = _2228_v73
		_2228_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2190_v43, _2227_v72)
		var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_2216_v69), 0))
		_ = _index380
		(_2216_v69).ArraySet1((_2228_v73).Merge(_2228_v73), (_index380).Int())
		var _2229_v74 _dafny.Set
		_ = _2229_v74
		_2229_v74 = _dafny.SetOf(_2190_v43)
		var _2230_v75 D9
		_ = _2230_v75
		_2230_v75 = Companion_D9_.Create_DC30_(_2191_v44, _this.F14(), _2190_v43, _2229_v74, _2190_v43)
		var _pat_let_tv48 = _2190_v43
		_ = _pat_let_tv48
		var _pat_let_tv49 = _2190_v43
		_ = _pat_let_tv49
		var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_2216_v69), 0))
		_ = _index381
		var _rhs346 bool = true
		_ = _rhs346
		var _rhs347 _dafny.CodePoint = (func(_pat_let70_0 D9) D9 {
			return func(_2231_dt__update__tmp_h0 D9) D9 {
				return func(_pat_let71_0 _dafny.Int) D9 {
					return func(_2232_dt__update_hcf54_h0 _dafny.Int) D9 {
						return func(_pat_let72_0 bool) D9 {
							return func(_2233_dt__update_hcf51_h0 bool) D9 {
								return func(_pat_let73_0 _dafny.Int) D9 {
									return func(_2234_dt__update_hcf52_h0 _dafny.Int) D9 {
										return Companion_D9_.Create_DC30_((_2231_dt__update__tmp_h0).Dtor_cf50(), _2233_dt__update_hcf51_h0, _2234_dt__update_hcf52_h0, (_2231_dt__update__tmp_h0).Dtor_cf53(), _2232_dt__update_hcf54_h0)
									}(_pat_let73_0)
								}(_pat_let_tv49)
							}(_pat_let72_0)
						}(_this.F14())
					}(_pat_let71_0)
				}(_pat_let_tv48)
			}(_pat_let70_0)
		}(_2230_v75)).Dtor_cf50()
		_ = _rhs347
		var _rhs348 _dafny.Map = Companion_Default___.Fm52(((_dafny.Zero).Minus(_2190_v43)).Times(_2190_v43), globalState)
		_ = _rhs348
		var _lhs295 *GlobalState = globalState
		_ = _lhs295
		var _lhs296 _dafny.Array = _2216_v69
		_ = _lhs296
		var _lhs297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(793), _dafny.ArrayLen((_2216_v69), 0))
		_ = _lhs297
		_lhs295.F9 = _rhs346
		_2191_v44 = _rhs347
		(_lhs296).ArraySet1(_rhs348, (_lhs297).Int())
		var _2235_v76 _dafny.Array
		_ = _2235_v76
		var _nw354 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw354
		_2235_v76 = _nw354
		var _2236_v77 _dafny.Sequence
		_ = _2236_v77
		_2236_v77 = _dafny.SeqOf(_2235_v76, _2235_v76)
		_2235_v76 = (func() _dafny.Array {
			if (func() bool {
				if _this.F14() {
					return true
				}
				return _this.F14()
			})() {
				return _2235_v76
			}
			return (_2236_v77).Select((Companion_Default___.SafeIndex(_2190_v43, _dafny.IntOfUint32((_2236_v77).Cardinality()))).Uint32()).(_dafny.Array)
		})()
		r0 = (_2229_v74).Cardinality()
		var _2237_v78 _dafny.Map
		_ = _2237_v78
		_2237_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(287), _this.F14())
		var _2238_v79 _dafny.Sequence
		_ = _2238_v79
		_2238_v79 = _dafny.SeqOf(_2237_v78)
		r1 = (Companion_D9_.Create_DC30_(_2191_v44, _this.F14(), _2190_v43, Companion_Default___.Fm22(_2190_v43, _2238_v79, _this.F14(), _this.F14(), globalState), _2190_v43)).Dtor_cf51()
		var _2239_v80 _dafny.Map
		_ = _2239_v80
		_2239_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2237_v78, _2190_v43)
		r2 = _2239_v80
		r3 = !(_this.F14())
		return r0, r1, r2, r3
	}
}
func (_this *C8) M1(p0 _dafny.Set, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _2240_v0 _dafny.Map
		_ = _2240_v0
		_2240_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
		_2240_v0 = _2240_v0
		var _2241_v1 _dafny.Array
		_ = _2241_v1
		var _nwElement0_68 bool = _this.F14()
		_ = _nwElement0_68
		var _nw355 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(10))
		_ = _nw355
		(_nw355).ArraySet1(_nwElement0_68, 0)
		(_nw355).ArraySet1(_this.F14(), 1)
		(_nw355).ArraySet1(_this.F14(), 2)
		(_nw355).ArraySet1(_this.F14(), 3)
		(_nw355).ArraySet1(_this.F14(), 4)
		(_nw355).ArraySet1(_this.F14(), 5)
		(_nw355).ArraySet1(false, 6)
		(_nw355).ArraySet1(_this.F14(), 7)
		(_nw355).ArraySet1(_this.F14(), 8)
		(_nw355).ArraySet1(_this.F14(), 9)
		_2241_v1 = _nw355
		var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2241_v1), 0))
		_ = _index382
		(_2241_v1).ArraySet1(_this.F14(), (_index382).Int())
		var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2241_v1), 0))
		_ = _index383
		(_2241_v1).ArraySet1(!(_this.F14()), (_index383).Int())
		var _2242_v2 *C1
		_ = _2242_v2
		var _nw356 *C1 = New_C1_()
		_ = _nw356
		_nw356.Ctor__()
		_2242_v2 = _nw356
		(globalState).F9 = !((p1).Cmp(p1) >= 0)
		var _2243_v3 _dafny.Array
		_ = _2243_v3
		var _nw357 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw357
		_2243_v3 = _nw357
		for _iter95 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2243_v3), 0))); ; {
			_guard_loop_7, _ok95 := _iter95()
			if !_ok95 {
				break
			}
			var _2244_i0 _dafny.Int
			_2244_i0 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_2244_i0).Sign() != -1) && ((_2244_i0).Cmp(_dafny.ArrayLen((_2243_v3), 0)) < 0)) {
				(_2243_v3).ArraySet1(Companion_Default___.SafeDivisionInt(_2244_i0, p1), (_2244_i0).Int())
			}
		}
		var _2245_v4 _dafny.MultiSet
		_ = _2245_v4
		_2245_v4 = _dafny.MultiSetOf(((_2240_v0).Cardinality()).Cmp(p1) == 0)
		_2245_v4 = (_this).F15()
		var _2246_v5 _dafny.Sequence
		_ = _2246_v5
		_2246_v5 = _dafny.UnicodeSeqOfUtf8Bytes("cxh")
		r0 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2246_v5, _dafny.UnicodeSeqOfUtf8Bytes("u"))).Cardinality())).Times(p1)
		r1 = p1
		r2 = (_2241_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(984), _dafny.ArrayLen((_2241_v1), 0))).Int()).(bool)
		return r0, r1, r2
	}
}
func (_this *C8) M2(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	{
		var r0 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r0
		var _2247_v0 D8
		_ = _2247_v0
		_2247_v0 = Companion_D8_.Create_DC27_()
		_2247_v0 = _2247_v0
		(globalState).F3 = p0
		var _2248_v1 _dafny.Array
		_ = _2248_v1
		var _nw358 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
		_ = _nw358
		_2248_v1 = _nw358
		for _iter96 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2248_v1), 0))); ; {
			_guard_loop_8, _ok96 := _iter96()
			if !_ok96 {
				break
			}
			var _2249_i0 _dafny.Int
			_2249_i0 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_2249_i0).Sign() != -1) && ((_2249_i0).Cmp(_dafny.ArrayLen((_2248_v1), 0)) < 0)) {
				(_2248_v1).ArraySet1(Companion_Default___.SafeDivisionInt(_2249_i0, p0), (_2249_i0).Int())
			}
		}
		var _2250_v2 _dafny.Sequence
		_ = _2250_v2
		_2250_v2 = _dafny.SeqOf(_this.F14(), p3, p3)
		var _2251_v3 *C3
		_ = _2251_v3
		var _nw359 *C3 = New_C3_()
		_ = _nw359
		_nw359.Ctor__(p0, p1, p3, (_this).F15())
		_2251_v3 = _nw359
		var _2252_v4 _dafny.MultiSet
		_ = _2252_v4
		_2252_v4 = _dafny.MultiSetOf(_2251_v3)
		var _2253_v5 _dafny.CodePoint
		_ = _2253_v5
		_2253_v5 = _dafny.CodePoint('x')
		var _2254_v6 _dafny.Map
		_ = _2254_v6
		_2254_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F14())
		var _2255_v7 _dafny.Map
		_ = _2255_v7
		_2255_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2253_v5, ((_2254_v6).Update(p0, _this.F14())).Cardinality())
		var _2256_v8 _dafny.Set
		_ = _2256_v8
		_2256_v8 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(14))).Uint32(), func(coer134 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg134 _dafny.Int) interface{} {
				return coer134(arg134)
			}
		}((func(_2257_v3 *C3) func(_dafny.Int) _dafny.Int {
			return func(_2258_i1 _dafny.Int) _dafny.Int {
				return (_2257_v3).F20()
			}
		})(_2251_v3))))
		var _2259_v9 _dafny.Map
		_ = _2259_v9
		_2259_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23(p0, (_2250_v2).Select((Companion_Default___.SafeIndex((_2252_v4).Cardinality(), _dafny.IntOfUint32((_2250_v2).Cardinality()))).Uint32()).(bool), _2255_v7, _2256_v8, globalState), (_2251_v3).F20())
		var _2260_v10 _dafny.MultiSet
		_ = _2260_v10
		_2260_v10 = _dafny.MultiSetOf(p0)
		(globalState).F3 = (func() _dafny.Int {
			if (_2259_v9).Contains(!((_this.F14()) && (_this.F14()))) {
				return (_2259_v9).Get(!((_this.F14()) && (_this.F14()))).(_dafny.Int)
			}
			return (_dafny.Zero).Minus(Companion_Default___.Fm0(_this.F14(), _dafny.IntOfUint32(((_2251_v3).F21()).Cardinality()), p0, _2260_v10, globalState))
		})()
		var _2261_v11 _dafny.Array
		_ = _2261_v11
		var _nw360 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw360
		_2261_v11 = _nw360
		(globalState).F2 = _2261_v11
		var _hi10 _dafny.Int = (_2251_v3).F20()
		_ = _hi10
		for _2262_i2 := p0; _2262_i2.Cmp(_hi10) < 0; _2262_i2 = _2262_i2.Plus(_dafny.One) {
			var _2263_v12 *C5
			_ = _2263_v12
			var _nw361 *C5 = New_C5_()
			_ = _nw361
			_nw361.Ctor__(p0, p3, ((_this).F15()).Intersection((_this).F15()))
			_2263_v12 = _nw361
			var _2264_v13 _dafny.Map
			_ = _2264_v13
			_2264_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(751), p1)
			var _2265_v14 D8
			_ = _2265_v14
			_2265_v14 = Companion_D8_.Create_DC26_(p3, _2264_v13)
			var _2266_v15 _dafny.Map
			_ = _2266_v15
			_2266_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_2263_v12).Fm2(p2, p3, globalState))
			var _2267_v16 _dafny.Map
			_ = _2267_v16
			_2267_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (func() bool {
				if (_2266_v15).Contains(p2) {
					return (_2266_v15).Get(p2).(bool)
				}
				return p3
			})())
			var _rhs349 *C5 = _2263_v12
			_ = _rhs349
			var _rhs350 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_2265_v14).Dtor_cf46() {
					return (_2263_v12).F18()
				}
				return Companion_Default___.SafeModuloInt((_2263_v12).F18(), (_dafny.Zero).Minus((_2267_v16).Cardinality()))
			})())
			_ = _rhs350
			var _lhs298 *GlobalState = globalState
			_ = _lhs298
			_2263_v12 = _rhs349
			_lhs298.F3 = _rhs350
			var _2268_v17 D4
			_ = _2268_v17
			_2268_v17 = Companion_D4_.Create_DC12_(_2261_v11)
			_2268_v17 = _2268_v17
			var _2269_v18 T0
			_ = _2269_v18
			var _nw362 *C5 = New_C5_()
			_ = _nw362
			_nw362.Ctor__(p0, p3, (_this).F15())
			_2269_v18 = _nw362
			var _2270_v19 _dafny.Map
			_ = _2270_v19
			_2270_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2269_v18, p2)
			_2270_v19 = _2270_v19
			var _2271_v20 _dafny.Array
			_ = _2271_v20
			var _nw363 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw363
			_2271_v20 = _nw363
			var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_2271_v20), 0))
			_ = _index384
			(_2271_v20).ArraySet1(_2248_v1, (_index384).Int())
			var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen((_2271_v20), 0))
			_ = _index385
			(_2271_v20).ArraySet1(_2248_v1, (_index385).Int())
		}
		r0 = _2253_v5
		return r0
	}
}
func (_this *C8) F16() _dafny.Map {
	{
		return _this._f16
	}
}
func (_this *C8) F17() _dafny.Sequence {
	{
		return _this._f17
	}
}

// End of class C8
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
