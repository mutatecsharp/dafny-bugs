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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Sequence, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	return ((func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(985), _dafny.IntOfInt64(907))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Int
			_0_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(985)).Cmp(_0_v0) <= 0) && ((_0_v0).Cmp(_dafny.IntOfInt64(907)) < 0) {
				_coll0.Add((_0_v0).Minus(_dafny.IntOfInt64(-431)), _dafny.CodePoint('k'))
			}
		}
		return _coll0.ToMap()
	}()).Cardinality()).Minus((_dafny.Zero).Minus((_dafny.IntOfInt64(-875)).Plus(_dafny.IntOfInt64(32))))
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) bool {
	return (_dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-909)), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()), _dafny.IntOfInt64(-14), _dafny.IntOfInt64(948), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()))), _dafny.SeqOf(_dafny.IntOfInt64(502), _dafny.IntOfInt64(581), _dafny.IntOfInt64(-87), _dafny.IntOfInt64(242)))) && (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), false)).Cardinality())).Cardinality()).Cmp(_dafny.IntOfInt64(-105)) == 0)
}
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('t')
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	var _source0 D7 = Companion_D7_.Create_DC16_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(453))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality()))
	_ = _source0
	if _source0.Is_DC16() {
		var _2___mcc_h0 _dafny.Int = _source0.Get_().(D7_DC16).Cf29
		_ = _2___mcc_h0
		var _3_cf29 _dafny.Int = _2___mcc_h0
		_ = _3_cf29
		if true {
			return _dafny.CodePoint('w')
		} else {
			return _dafny.CodePoint('r')
		}
	} else if _source0.Is_DC15() {
		var _4___mcc_h1 _dafny.Sequence = _source0.Get_().(D7_DC15).Cf28
		_ = _4___mcc_h1
		var _5_cf28 _dafny.Sequence = _4___mcc_h1
		_ = _5_cf28
		return _dafny.CodePoint('i')
	} else {
		var _6___mcc_h2 D7 = _source0.Get_().(D7_DC17).Cf30
		_ = _6___mcc_h2
		var _7_cf30 D7 = _6___mcc_h2
		_ = _7_cf30
		return _dafny.CodePoint('q')
	}
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 _dafny.Sequence, p2 bool, p3 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	if !(!((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("og"))).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sple"))))) {
		return _dafny.CodePoint('l')
	} else {
		return _dafny.CodePoint('t')
	}
}
func (_static *CompanionStruct_Default___) Fm19(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(8))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_8_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('f')
	}))).Cardinality()))), (_dafny.Zero).Minus(_dafny.IntOfInt64(-377)))).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(218), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aqrya")).Cardinality()))).Cardinality()))))).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("c")).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(-649), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mbapixuot")).Cardinality()), _dafny.IntOfInt64(394))).Cardinality(), _dafny.IntOfInt64(923), _dafny.IntOfInt64(695)))))
}
func (_static *CompanionStruct_Default___) Fm20(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ofvstaor"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rfjmdefgm"), _dafny.UnicodeSeqOfUtf8Bytes("r")))
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _dafny.UnicodeSeqOfUtf8Bytes("sflqu")))
}
func (_static *CompanionStruct_Default___) Fm22(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Set {
	var _source1 D6 = (func() D6 {
		if false {
			return Companion_D6_.Create_DC11_(_dafny.MultiSetOf(_dafny.IntOfInt64(344)))
		}
		return Companion_D6_.Create_DC11_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfInt64(-70))))
	})()
	_ = _source1
	if _source1.Is_DC12() {
		var _9___mcc_h0 bool = _source1.Get_().(D6_DC12).Cf18
		_ = _9___mcc_h0
		var _10___mcc_h1 _dafny.Sequence = _source1.Get_().(D6_DC12).Cf19
		_ = _10___mcc_h1
		var _11___mcc_h2 _dafny.Int = _source1.Get_().(D6_DC12).Cf20
		_ = _11___mcc_h2
		var _12___mcc_h3 bool = _source1.Get_().(D6_DC12).Cf21
		_ = _12___mcc_h3
		var _13_cf21 bool = _12___mcc_h3
		_ = _13_cf21
		var _14_cf20 _dafny.Int = _11___mcc_h2
		_ = _14_cf20
		var _15_cf19 _dafny.Sequence = _10___mcc_h1
		_ = _15_cf19
		var _16_cf18 bool = _9___mcc_h0
		_ = _16_cf18
		return _dafny.SetOf(_15_cf19)
	} else if _source1.Is_DC13() {
		var _17___mcc_h4 bool = _source1.Get_().(D6_DC13).Cf22
		_ = _17___mcc_h4
		var _18___mcc_h5 bool = _source1.Get_().(D6_DC13).Cf23
		_ = _18___mcc_h5
		var _19___mcc_h6 _dafny.Int = _source1.Get_().(D6_DC13).Cf24
		_ = _19___mcc_h6
		var _20___mcc_h7 _dafny.Int = _source1.Get_().(D6_DC13).Cf25
		_ = _20___mcc_h7
		var _21___mcc_h8 bool = _source1.Get_().(D6_DC13).Cf26
		_ = _21___mcc_h8
		var _22_cf26 bool = _21___mcc_h8
		_ = _22_cf26
		var _23_cf25 _dafny.Int = _20___mcc_h7
		_ = _23_cf25
		var _24_cf24 _dafny.Int = _19___mcc_h6
		_ = _24_cf24
		var _25_cf23 bool = _18___mcc_h5
		_ = _25_cf23
		var _26_cf22 bool = _17___mcc_h4
		_ = _26_cf22
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("yyjhaqfxv"))
	} else if _source1.Is_DC11() {
		var _27___mcc_h9 _dafny.MultiSet = _source1.Get_().(D6_DC11).Cf17
		_ = _27___mcc_h9
		var _28_cf17 _dafny.MultiSet = _27___mcc_h9
		_ = _28_cf17
		return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ycub"), _dafny.UnicodeSeqOfUtf8Bytes("dvtlsfkq"))
	} else {
		var _29___mcc_h10 D6 = _source1.Get_().(D6_DC14).Cf27
		_ = _29___mcc_h10
		var _30_cf27 D6 = _29___mcc_h10
		_ = _30_cf27
		return func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vecee"), true)).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _31_v0 _dafny.Sequence
				_31_v0 = interface{}(_compr_1).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vecee"), true)).Contains(_31_v0) {
					_coll1.Add(_31_v0)
				}
			}
			return _coll1.ToSet()
		}()
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(((Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.SetOf(true, false), _dafny.MultiSetOf(_dafny.IntOfInt64(296)), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lqbbxmxcd")))).Dtor_cf34()).Cardinality()))).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _32_v0 _dafny.Int
			_32_v0 = interface{}(_compr_2).(_dafny.Int)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(((Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), _dafny.SetOf(true, false), _dafny.MultiSetOf(_dafny.IntOfInt64(296)), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lqbbxmxcd")))).Dtor_cf34()).Cardinality()))).Contains(_32_v0) {
				_coll2.Add(Companion_Default___.SafeModuloInt(_32_v0, (_dafny.MultiSetOf(_dafny.IntOfInt64(576), _dafny.IntOfInt64(33))).Cardinality()), _dafny.CodePoint('f'))
			}
		}
		return _coll2.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-86), _dafny.CodePoint('f')))).Merge((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(137), _dafny.CodePoint('n'))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-424))).Cardinality(), _dafny.CodePoint('w'))
	})())
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(137), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gfpmde")).Cardinality()), true)).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(508))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_33_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qtnakf")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bug")).Cardinality())), _dafny.IntOfInt64(798)))
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-151))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_34_i0 _dafny.Int) _dafny.Map {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("puihdeaw")).Cardinality()), !(!(false)))
	}))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(757))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_35_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		})))).Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _36_v0 _dafny.Sequence
			_36_v0 = interface{}(_compr_3).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(757))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_35_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))), _36_v0) {
				_coll3.Add(_36_v0, !(true))
			}
		}
		return _coll3.ToMap()
	}()).Merge(func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tio"), _dafny.SetOf(true))).Keys().Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _37_v1 _dafny.Sequence
			_37_v1 = interface{}(_compr_4).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("tio"), _dafny.SetOf(true))).Contains(_37_v1) {
				_coll4.Add(_37_v1, false)
			}
		}
		return _coll4.ToMap()
	}())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D9_.Create_DC22_(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(false, true)).Cardinality(), true)).Cardinality(), !(true), _dafny.SetOf(!(true)), _dafny.UnicodeSeqOfUtf8Bytes("gc"))).Dtor_cf41(), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("vgp"), !(false))))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(538))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfInt64(-901)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-458))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_39_i1 _dafny.Int) _dafny.Int {
		return (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(((Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(104)), _dafny.SetOf(false), _dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfInt64(445))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dixgeiij")).Cardinality())), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("lqypklsa")))).Dtor_cf32()).Cardinality())).Cardinality(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true, true)).Cardinality())))).Cardinality())
	})))
}
func (_static *CompanionStruct_Default___) Fm28(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(false)).Intersection(_dafny.MultiSetOf(false))).Intersection((_dafny.MultiSetOf(false, false, true)).Union(_dafny.MultiSetOf((Companion_D6_.Create_DC12_(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(186))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_40_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	})), _dafny.IntOfInt64(574), true)).Dtor_cf18(), true, false, false, true)))
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(!(false)))
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-335))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_41_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("aavabqni")))).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _42_v0 _dafny.Sequence
			_42_v0 = interface{}(_compr_5).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-335))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_41_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('a')
			}))), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("aavabqni"))), _42_v0) {
				_coll5.Add(_42_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.CodePoint('n')))
			}
		}
		return _coll5.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm32(p0 D6, globalState *GlobalState) _dafny.Sequence {
	var _source2 D16 = Companion_D16_.Create_DC41_(true)
	_ = _source2
	if _source2.Is_DC41() {
		var _43___mcc_h0 bool = _source2.Get_().(D16_DC41).Cf84
		_ = _43___mcc_h0
		var _44_cf84 bool = _43___mcc_h0
		_ = _44_cf84
		return _dafny.UnicodeSeqOfUtf8Bytes("sqsbuy")
	} else if _source2.Is_DC42() {
		var _45___mcc_h1 _dafny.MultiSet = _source2.Get_().(D16_DC42).Cf85
		_ = _45___mcc_h1
		var _46_cf85 _dafny.MultiSet = _45___mcc_h1
		_ = _46_cf85
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(660))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_47_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))
	} else if _source2.Is_DC40() {
		var _48___mcc_h2 T0 = _source2.Get_().(D16_DC40).Cf83
		_ = _48___mcc_h2
		var _49_cf83 T0 = _48___mcc_h2
		_ = _49_cf83
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sidmalhfq"), _dafny.UnicodeSeqOfUtf8Bytes("nga"))
	} else {
		var _50___mcc_h3 D16 = _source2.Get_().(D16_DC43).Cf86
		_ = _50___mcc_h3
		var _51_cf86 D16 = _50___mcc_h3
		_ = _51_cf86
		return _dafny.UnicodeSeqOfUtf8Bytes("bnryd")
	}
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('o')
}
func (_static *CompanionStruct_Default___) Fm34(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pdggpy")).Cardinality())), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(859)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-621))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_52_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-560))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg13 _dafny.Int) interface{} {
				return coer13(arg13)
			}
		}(func(_53_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(497)
		}))).Cardinality())
	}))))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((true) == (false), _dafny.UnicodeSeqOfUtf8Bytes("coubrsv"))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 _dafny.CodePoint, p2 bool, globalState *GlobalState) _dafny.Set {
	if (_dafny.SetOf((func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(789), _dafny.IntOfInt64(442))); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _54_v1 _dafny.Int
			_54_v1 = interface{}(_compr_6).(_dafny.Int)
			if ((_dafny.IntOfInt64(789)).Cmp(_54_v1) <= 0) && ((_54_v1).Cmp(_dafny.IntOfInt64(442)) < 0) {
				_coll6.Add((_54_v1).Times(_dafny.IntOfInt64(-848)), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(true, true)).Cardinality())).Cardinality()), _dafny.IntOfInt64(12)))).Cardinality())
			}
		}
		return _coll6.ToMap()
	}()).Cardinality())).IsSubsetOf(func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hvfvypqh")).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("weervqcj"))).Keys().Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _55_v0 _dafny.Int
			_55_v0 = interface{}(_compr_7).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hvfvypqh")).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("weervqcj"))).Contains(_55_v0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_55_v0, (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lhu")).Cardinality()))).Cardinality()))
			}
		}
		return _coll7.ToSet()
	}()) {
		return (_dafny.SetOf(true, !(false), false)).Union(_dafny.SetOf(true))
	} else {
		return _dafny.SetOf(true)
	}
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D3_.Create_DC7_(), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ti")).Cardinality()), _dafny.IntOfInt64(332)))).Cardinality(), (_dafny.IntOfInt64(-570)).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wmlctwlx")).Cardinality())), false)
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf((func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(201), _dafny.IntOfInt64(-244))); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _56_v0 _dafny.Int
			_56_v0 = interface{}(_compr_8).(_dafny.Int)
			if ((_dafny.IntOfInt64(201)).Cmp(_56_v0) <= 0) && ((_56_v0).Cmp(_dafny.IntOfInt64(-244)) < 0) {
				_coll8.Add((_56_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-467), _dafny.IntOfInt64(716))).Cardinality()))))
			}
		}
		return _coll8.ToSet()
	}()).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(289), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-587), _dafny.CodePoint('q'))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(969))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_57_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('v')
	}))).Cardinality()), (func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(371), _dafny.IntOfInt64(582))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _58_v1 _dafny.Int
			_58_v1 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(371)).Cmp(_58_v1) <= 0) && ((_58_v1).Cmp(_dafny.IntOfInt64(582)) < 0) {
				_coll9.Add((_58_v1).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(704))).Cardinality()))
			}
		}
		return _coll9.ToSet()
	}()).Cardinality())).Cardinality(), false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm39(p0 bool, p1 bool, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC13_((_dafny.IntOfInt64(-93)).Cmp(_dafny.IntOfInt64(37)) == 0, false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(933))).Uint32(), func(coer15 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg15 _dafny.Int) interface{} {
			return coer15(arg15)
		}
	}(func(_59_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-741)
	}))).Cardinality())), _dafny.IntOfInt64(232), !(!(true) || (!(false))))
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true))).Union(_dafny.MultiSetOf(true, !(false)))
}
func (_static *CompanionStruct_Default___) Fm41(p0 _dafny.Int, p1 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((Companion_D21_.Create_DC54_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fdkkibdc")).Cardinality()))))).Dtor_cf105()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(958))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dg")).Cardinality())))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf((func() _dafny.Set {
		var _coll10 = _dafny.NewBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(272), _dafny.IntOfInt64(733))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _60_v0 _dafny.Int
			_60_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(272)).Cmp(_60_v0) <= 0) && ((_60_v0).Cmp(_dafny.IntOfInt64(733)) < 0) {
				_coll10.Add(Companion_Default___.SafeDivisionInt(_60_v0, _dafny.IntOfInt64(509)))
			}
		}
		return _coll10.ToSet()
	}()).Cardinality(), (_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("win"))).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm44(p0 bool, globalState *GlobalState) _dafny.Map {
	var _source3 D21 = (func() D21 {
		if true {
			return Companion_D21_.Create_DC56_(true)
		}
		return Companion_D21_.Create_DC56_(false)
	})()
	_ = _source3
	if _source3.Is_DC55() {
		var _61___mcc_h0 bool = _source3.Get_().(D21_DC55).Cf106
		_ = _61___mcc_h0
		var _62___mcc_h1 _dafny.Int = _source3.Get_().(D21_DC55).Cf107
		_ = _62___mcc_h1
		var _63___mcc_h2 _dafny.Int = _source3.Get_().(D21_DC55).Cf108
		_ = _63___mcc_h2
		var _64_cf108 _dafny.Int = _63___mcc_h2
		_ = _64_cf108
		var _65_cf107 _dafny.Int = _62___mcc_h1
		_ = _65_cf107
		var _66_cf106 bool = _61___mcc_h0
		_ = _66_cf106
		return (func() _dafny.Map {
			var _coll11 = _dafny.NewMapBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_66_cf106), _66_cf106)).Keys().Elements()); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _67_v1 _dafny.Set
					_67_v1 = interface{}(_compr_12).(_dafny.Set)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_66_cf106), _66_cf106)).Contains(_67_v1) {
						_coll12.Add(_67_v1)
					}
				}
				return _coll12.ToSet()
			}()).Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _68_v0 _dafny.Set
				_68_v0 = interface{}(_compr_11).(_dafny.Set)
				if (func() _dafny.Set {
					var _coll13 = _dafny.NewBuilder()
					_ = _coll13
					for _iter13 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_66_cf106), _66_cf106)).Keys().Elements()); ; {
						_compr_13, _ok13 := _iter13()
						if !_ok13 {
							break
						}
						var _69_v1 _dafny.Set
						_69_v1 = interface{}(_compr_13).(_dafny.Set)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_66_cf106), _66_cf106)).Contains(_69_v1) {
							_coll13.Add(_69_v1)
						}
					}
					return _coll13.ToSet()
				}()).Contains(_68_v0) {
					_coll11.Add(_68_v0, (_dafny.Zero).Minus(_64_cf108))
				}
			}
			return _coll11.ToMap()
		}()).Merge((Companion_D22_.Create_DC58_(func() _dafny.Map {
			var _coll14 = _dafny.NewMapBuilder()
			_ = _coll14
			for _iter14 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(!(true)))).Elements()); ; {
				_compr_14, _ok14 := _iter14()
				if !_ok14 {
					break
				}
				var _70_v2 _dafny.Set
				_70_v2 = interface{}(_compr_14).(_dafny.Set)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(!(true))), _70_v2) {
					_coll14.Add(_70_v2, _64_cf108)
				}
			}
			return _coll14.ToMap()
		}())).Dtor_cf112())
	} else if _source3.Is_DC56() {
		var _71___mcc_h3 bool = _source3.Get_().(D21_DC56).Cf109
		_ = _71___mcc_h3
		var _72_cf109 bool = _71___mcc_h3
		_ = _72_cf109
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_72_cf109, _72_cf109), (_dafny.MultiSetOf(true)).Cardinality())
	} else if _source3.Is_DC57() {
		var _73___mcc_h4 bool = _source3.Get_().(D21_DC57).Cf110
		_ = _73___mcc_h4
		var _74___mcc_h5 _dafny.CodePoint = _source3.Get_().(D21_DC57).Cf111
		_ = _74___mcc_h5
		var _75_cf111 _dafny.CodePoint = _74___mcc_h5
		_ = _75_cf111
		var _76_cf110 bool = _73___mcc_h4
		_ = _76_cf110
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_76_cf110), (_dafny.MultiSetOf(_76_cf110)).Cardinality())
	} else {
		var _77___mcc_h6 _dafny.Map = _source3.Get_().(D21_DC54).Cf105
		_ = _77___mcc_h6
		var _78_cf105 _dafny.Map = _77___mcc_h6
		_ = _78_cf105
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(true), _dafny.IntOfInt64(345))
	}
}
func (_static *CompanionStruct_Default___) Fm45(p0 D6, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("k"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(483))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_79_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(797))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}(func(_80_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})), _dafny.UnicodeSeqOfUtf8Bytes("wenrpms")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("shmvcrb")))
}
func (_static *CompanionStruct_Default___) Fm46(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC22_((func() bool {
		if false {
			return false
		}
		return true
	})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-266)), _dafny.SeqOf(_dafny.IntOfInt64(997), _dafny.IntOfInt64(737)))).Cardinality())), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uxklvq")).Cardinality()))).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(264))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg18 _dafny.Int) interface{} {
			return coer18(arg18)
		}
	}(func(_81_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality())) <= 0, _dafny.SetOf(false, false), _dafny.UnicodeSeqOfUtf8Bytes("vot"))
}
func (_static *CompanionStruct_Default___) Fm47(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(167), (_dafny.Zero).Minus(_dafny.IntOfInt64(-372)))).Intersection((_dafny.SetOf(_dafny.IntOfInt64(-541))).Difference(func() _dafny.Set {
		var _coll15 = _dafny.NewBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_dafny.SetOf((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-279))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg19 _dafny.Int) interface{} {
				return coer19(arg19)
			}
		}(func(_82_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		}))).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg20 _dafny.Int) interface{} {
				return coer20(arg20)
			}
		}(func(_83_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))).Cardinality()))).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _84_v0 _dafny.Int
			_84_v0 = interface{}(_compr_15).(_dafny.Int)
			if (_dafny.SetOf((_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-279))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_82_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))).Cardinality())))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_83_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('i')
			}))).Cardinality()))).Contains(_84_v0) {
				_coll15.Add((_84_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("thfmph")).Cardinality())))
			}
		}
		return _coll15.ToSet()
	}()))
}
func (_static *CompanionStruct_Default___) Fm48(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(727))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg23 _dafny.Int) interface{} {
			return coer23(arg23)
		}
	}(func(_85_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('w')
	})))).Intersection(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("utp")))).Intersection((func() _dafny.Set {
		var _coll16 = _dafny.NewBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("i")))).Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _86_v0 _dafny.Sequence
			_86_v0 = interface{}(_compr_16).(_dafny.Sequence)
			if (_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("i")))).Contains(_86_v0) {
				_coll16.Add(_86_v0)
			}
		}
		return _coll16.ToSet()
	}()).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ogjtqc"), _dafny.UnicodeSeqOfUtf8Bytes("hyaujykjq"))))
}
func (_static *CompanionStruct_Default___) Fm49(p0 D1, p1 _dafny.Set, globalState *GlobalState) D3 {
	return Companion_D3_.Create_DC8_(((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("idqyyo")).Cardinality()))).Union(_dafny.MultiSetOf((_dafny.MultiSetOf(true)).Cardinality()))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm50(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.MultiSetOf(_dafny.IntOfInt64(956))).Cardinality()).Plus(_dafny.IntOfInt64(415)), !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("v"), _dafny.CodePoint('s')))
}
func (_static *CompanionStruct_Default___) Fm51(globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC23_((func() _dafny.Int {
		if true {
			return _dafny.IntOfInt64(416)
		}
		return _dafny.IntOfInt64(-787)
	})(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(371), _dafny.IntOfInt64(820)), ((_dafny.MultiSetOf(true, false, !(true))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, !(!(true)), true, false)))).Cardinality(), (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-494))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg24 _dafny.Int) interface{} {
			return coer24(arg24)
		}
	}(func(_87_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality())).Times(_dafny.IntOfInt64(839)), _dafny.IntOfInt64(109))
}
func (_static *CompanionStruct_Default___) Fm54(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(!(!(false)))).Intersection(_dafny.SetOf(false, true))).Difference(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm55(p0 _dafny.Int, p1 _dafny.Set, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(true)
}
func (_static *CompanionStruct_Default___) Fm56(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf((_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bdebqr")).Cardinality()))).Cardinality(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hixovsg")).Cardinality()))).Cardinality()))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer25 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg25 _dafny.Int) interface{} {
			return coer25(arg25)
		}
	}(func(_88_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(-869)
	}))).Cardinality()), (_dafny.MultiSetOf(false, false, false)).Cardinality()))).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(941)))
}
func (_static *CompanionStruct_Default___) Fm57(globalState *GlobalState) D12 {
	var _source4 D21 = Companion_D21_.Create_DC56_(true)
	_ = _source4
	if _source4.Is_DC55() {
		var _89___mcc_h0 bool = _source4.Get_().(D21_DC55).Cf106
		_ = _89___mcc_h0
		var _90___mcc_h1 _dafny.Int = _source4.Get_().(D21_DC55).Cf107
		_ = _90___mcc_h1
		var _91___mcc_h2 _dafny.Int = _source4.Get_().(D21_DC55).Cf108
		_ = _91___mcc_h2
		var _92_cf108 _dafny.Int = _91___mcc_h2
		_ = _92_cf108
		var _93_cf107 _dafny.Int = _90___mcc_h1
		_ = _93_cf107
		var _94_cf106 bool = _89___mcc_h0
		_ = _94_cf106
		return Companion_D12_.Create_DC32_(_94_cf106, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("k")).Cardinality()), false, _94_cf106, _92_cf108)
	} else if _source4.Is_DC56() {
		var _95___mcc_h3 bool = _source4.Get_().(D21_DC56).Cf109
		_ = _95___mcc_h3
		var _96_cf109 bool = _95___mcc_h3
		_ = _96_cf109
		return Companion_D12_.Create_DC32_(_96_cf109, _dafny.IntOfInt64(649), _96_cf109, _96_cf109, (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_cf109, _96_cf109), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_cf109, _96_cf109), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_cf109, _96_cf109), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_cf109, _96_cf109), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_cf109, _96_cf109))).Cardinality())
	} else if _source4.Is_DC57() {
		var _97___mcc_h4 bool = _source4.Get_().(D21_DC57).Cf110
		_ = _97___mcc_h4
		var _98___mcc_h5 _dafny.CodePoint = _source4.Get_().(D21_DC57).Cf111
		_ = _98___mcc_h5
		var _99_cf111 _dafny.CodePoint = _98___mcc_h5
		_ = _99_cf111
		var _100_cf110 bool = _97___mcc_h4
		_ = _100_cf110
		return Companion_D12_.Create_DC32_(!(_100_cf110), _dafny.IntOfUint32(((Companion_D22_.Create_DC60_(_dafny.IntOfInt64(-656), _dafny.SeqOf(false, !(_100_cf110)), false)).Dtor_cf116()).Cardinality()), _100_cf110, _100_cf110, _dafny.IntOfInt64(559))
	} else {
		var _101___mcc_h6 _dafny.Map = _source4.Get_().(D21_DC54).Cf105
		_ = _101___mcc_h6
		var _102_cf105 _dafny.Map = _101___mcc_h6
		_ = _102_cf105
		return Companion_D12_.Create_DC32_(false, _dafny.IntOfInt64(-50), false, true, (_dafny.SetOf(_dafny.IntOfInt64(698), _dafny.IntOfInt64(259))).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm58(globalState *GlobalState) _dafny.CodePoint {
	if (_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(325))).Cardinality())).Cmp((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-639))).Uint32(), func(coer26 func(_dafny.Int) D22) func(_dafny.Int) interface{} {
		return func(arg26 _dafny.Int) interface{} {
			return coer26(arg26)
		}
	}(func(_103_i0 _dafny.Int) D22 {
		return Companion_D22_.Create_DC59_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(506))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}(func(_104_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))).Cardinality()), false)
	})), _dafny.SeqOf(_dafny.IntOfInt64(668)))).Cardinality())).Cardinality()) > 0 {
		return _dafny.CodePoint('i')
	} else if !(false) {
		return _dafny.CodePoint('f')
	} else {
		return _dafny.CodePoint('q')
	}
}
func (_static *CompanionStruct_Default___) Fm59(p0 bool, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), true)
}
func (_static *CompanionStruct_Default___) Fm60(p0 bool, p1 bool, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(182))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg28 _dafny.Int) interface{} {
			return coer28(arg28)
		}
	}(func(_105_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('h')
	}))
}
func (_static *CompanionStruct_Default___) Fm61(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) _dafny.Sequence {
	var _source5 D6 = Companion_D6_.Create_DC11_(_dafny.MultiSetOf(_dafny.IntOfInt64(-259)))
	_ = _source5
	if _source5.Is_DC12() {
		var _106___mcc_h0 bool = _source5.Get_().(D6_DC12).Cf18
		_ = _106___mcc_h0
		var _107___mcc_h1 _dafny.Sequence = _source5.Get_().(D6_DC12).Cf19
		_ = _107___mcc_h1
		var _108___mcc_h2 _dafny.Int = _source5.Get_().(D6_DC12).Cf20
		_ = _108___mcc_h2
		var _109___mcc_h3 bool = _source5.Get_().(D6_DC12).Cf21
		_ = _109___mcc_h3
		var _110_cf21 bool = _109___mcc_h3
		_ = _110_cf21
		var _111_cf20 _dafny.Int = _108___mcc_h2
		_ = _111_cf20
		var _112_cf19 _dafny.Sequence = _107___mcc_h1
		_ = _112_cf19
		var _113_cf18 bool = _106___mcc_h0
		_ = _113_cf18
		return _dafny.SeqOf(!(_110_cf21), true, _110_cf21, false, false)
	} else if _source5.Is_DC13() {
		var _114___mcc_h4 bool = _source5.Get_().(D6_DC13).Cf22
		_ = _114___mcc_h4
		var _115___mcc_h5 bool = _source5.Get_().(D6_DC13).Cf23
		_ = _115___mcc_h5
		var _116___mcc_h6 _dafny.Int = _source5.Get_().(D6_DC13).Cf24
		_ = _116___mcc_h6
		var _117___mcc_h7 _dafny.Int = _source5.Get_().(D6_DC13).Cf25
		_ = _117___mcc_h7
		var _118___mcc_h8 bool = _source5.Get_().(D6_DC13).Cf26
		_ = _118___mcc_h8
		var _119_cf26 bool = _118___mcc_h8
		_ = _119_cf26
		var _120_cf25 _dafny.Int = _117___mcc_h7
		_ = _120_cf25
		var _121_cf24 _dafny.Int = _116___mcc_h6
		_ = _121_cf24
		var _122_cf23 bool = _115___mcc_h5
		_ = _122_cf23
		var _123_cf22 bool = _114___mcc_h4
		_ = _123_cf22
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_123_cf22), _dafny.SeqOf(_119_cf26))
	} else if _source5.Is_DC11() {
		var _124___mcc_h9 _dafny.MultiSet = _source5.Get_().(D6_DC11).Cf17
		_ = _124___mcc_h9
		var _125_cf17 _dafny.MultiSet = _124___mcc_h9
		_ = _125_cf17
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(true))
	} else {
		var _126___mcc_h10 D6 = _source5.Get_().(D6_DC14).Cf27
		_ = _126___mcc_h10
		var _127_cf27 D6 = _126___mcc_h10
		_ = _127_cf27
		return _dafny.SeqOf(!(true))
	}
}
func (_static *CompanionStruct_Default___) Fm62(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cb"), _dafny.UnicodeSeqOfUtf8Bytes("rjfj")))
}
func (_static *CompanionStruct_Default___) Fm63(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	if !(false) {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(791))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_128_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(489)
		})), _dafny.SeqOf(_dafny.IntOfInt64(31), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(449), true)).Cardinality())))
	} else {
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(995))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg30 _dafny.Int) interface{} {
				return coer30(arg30)
			}
		}(func(_129_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gqn")).Cardinality())
		})), _dafny.SeqOf((_dafny.MultiSetOf(false)).Cardinality()))
	}
}
func (_static *CompanionStruct_Default___) Fm64(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(787))).Intersection((_dafny.SetOf(_dafny.IntOfInt64(350), _dafny.IntOfInt64(-91), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()))).Intersection((Companion_D19_.Create_DC49_(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(816))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg31 _dafny.Int) interface{} {
			return coer31(arg31)
		}
	}(func(_130_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))).Cardinality()))).Cardinality()))).Dtor_cf93()))
}
func (_static *CompanionStruct_Default___) Fm65(p0 bool, p1 _dafny.Sequence, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32(((Companion_D19_.Create_DC50_(true, (_dafny.MultiSetOf(_dafny.IntOfInt64(107))).Cardinality(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(885))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg32 _dafny.Int) interface{} {
			return coer32(arg32)
		}
	}(func(_131_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Set {
			var _coll17 = _dafny.NewBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(649))).Cardinality())).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _132_v0 _dafny.Int
				_132_v0 = interface{}(_compr_17).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(649))).Cardinality()), _132_v0) {
					_coll17.Add((_132_v0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xfi")).Cardinality())))
				}
			}
			return _coll17.ToSet()
		}()).Cardinality()), (_dafny.SetOf(_dafny.IntOfInt64(74))).Cardinality())).Cardinality()
	})))).Dtor_cf96()).Cardinality()), _dafny.IntOfInt64(719), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(false, true), _dafny.SeqOf(true))).Cardinality())), (Companion_D11_.Create_DC27_(true, _dafny.CodePoint('o'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ycrugglv")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nokerh")).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(453)))).Dtor_cf55(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ggmughh")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((Companion_D3_.Create_DC6_((_dafny.MultiSetOf(true, true)).Cardinality(), false)).Dtor_cf13(), true)).Cardinality()), (func() _dafny.Set {
		var _coll18 = _dafny.NewBuilder()
		_ = _coll18
		for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(210), _dafny.IntOfInt64(31))); ; {
			_compr_18, _ok18 := _iter18()
			if !_ok18 {
				break
			}
			var _133_v1 _dafny.Int
			_133_v1 = interface{}(_compr_18).(_dafny.Int)
			if ((_dafny.IntOfInt64(210)).Cmp(_133_v1) <= 0) && ((_133_v1).Cmp(_dafny.IntOfInt64(31)) < 0) {
				_coll18.Add((_133_v1).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jrgri")).Cardinality())))
			}
		}
		return _coll18.ToSet()
	}()).Cardinality(), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.MultiSetOf(true)).Cardinality())), (_dafny.SetOf(true, true, false)).Cardinality()), _dafny.SeqOf((_dafny.SetOf(false, false)).Cardinality())), _dafny.SeqOf((_dafny.MultiSetOf((func() _dafny.Set {
		var _coll19 = _dafny.NewBuilder()
		_ = _coll19
		for _iter19 := _dafny.Iterate((func() _dafny.Map {
			var _coll20 = _dafny.NewMapBuilder()
			_ = _coll20
			for _iter20 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(293)))).Elements()); ; {
				_compr_20, _ok20 := _iter20()
				if !_ok20 {
					break
				}
				var _134_v2 _dafny.Sequence
				_134_v2 = interface{}(_compr_20).(_dafny.Sequence)
				if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(293)))).Contains(_134_v2) {
					_coll20.Add(_134_v2, true)
				}
			}
			return _coll20.ToMap()
		}()).Keys().Elements()); ; {
			_compr_19, _ok19 := _iter19()
			if !_ok19 {
				break
			}
			var _135_v3 _dafny.Sequence
			_135_v3 = interface{}(_compr_19).(_dafny.Sequence)
			if (func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter21 := _dafny.Iterate((_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(293)))).Elements()); ; {
					_compr_21, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _134_v2 _dafny.Sequence
					_134_v2 = interface{}(_compr_21).(_dafny.Sequence)
					if (_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(293)))).Contains(_134_v2) {
						_coll21.Add(_134_v2, true)
					}
				}
				return _coll21.ToMap()
			}()).Contains(_135_v3) {
				_coll19.Add(_135_v3)
			}
		}
		return _coll19.ToSet()
	}()).Cardinality(), _dafny.IntOfInt64(-985), (_dafny.Zero).Minus(_dafny.IntOfInt64(-569)), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg33 _dafny.Int) interface{} {
			return coer33(arg33)
		}
	}(func(_136_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(106)
	}))).Cardinality()), (_dafny.SetOf(true)).Cardinality())).Cardinality()), (func() _dafny.Sequence {
		if false {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-157))).Uint32(), func(coer34 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_137_i2 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.CodePoint('m'))).Cardinality())).Cardinality())).Cardinality()
			}))
		}
		return _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(true)).Cardinality())).Cardinality(), (func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter22 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("glco"), _dafny.IntOfInt64(914))).Keys().Elements()); ; {
				_compr_22, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _138_v4 _dafny.Sequence
				_138_v4 = interface{}(_compr_22).(_dafny.Sequence)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("glco"), _dafny.IntOfInt64(914))).Contains(_138_v4) {
					_coll22.Add(_138_v4, false)
				}
			}
			return _coll22.ToMap()
		}()).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(751))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqoyap")).Cardinality()))).Cardinality()))
	})())
}
func (_static *CompanionStruct_Default___) Fm66(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("nnwdhunb"), _dafny.UnicodeSeqOfUtf8Bytes("dpcfmoy"))).Union(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("ehco"), _dafny.UnicodeSeqOfUtf8Bytes("eo"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(16))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg35 _dafny.Int) interface{} {
			return coer35(arg35)
		}
	}(func(_139_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(487))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg36 _dafny.Int) interface{} {
			return coer36(arg36)
		}
	}(func(_140_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('e')
	})), _dafny.UnicodeSeqOfUtf8Bytes("bj")))).Union(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("m"), _dafny.UnicodeSeqOfUtf8Bytes("bs"), _dafny.UnicodeSeqOfUtf8Bytes("rlwyon")))
}
func (_static *CompanionStruct_Default___) Fm67(p0 bool, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(693))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg37 _dafny.Int) interface{} {
			return coer37(arg37)
		}
	}(func(_141_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('t')
	}))).Cardinality()))), (_dafny.SetOf(true, true)).Difference(_dafny.SetOf(false, false)), (_dafny.MultiSetOf((_dafny.Zero).Minus(_dafny.IntOfInt64(-601)), _dafny.IntOfInt64(-799))).Union(_dafny.MultiSetOf(_dafny.IntOfInt64(-453))), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("n")))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.CodePoint, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
	_ = r0
	var _142_v0 _dafny.Array
	_ = _142_v0
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(5))
	_ = _nw0
	_142_v0 = _nw0
	for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_142_v0), 0))); ; {
		_guard_loop_0, _ok23 := _iter23()
		if !_ok23 {
			break
		}
		var _143_i0 _dafny.Int
		_143_i0 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_143_i0).Sign() != -1) && ((_143_i0).Cmp(_dafny.ArrayLen((_142_v0), 0)) < 0)) {
			(_142_v0).ArraySet1((_143_i0).Plus(_dafny.IntOfInt64(605)), (_143_i0).Int())
		}
	}
	var _144_i1 _dafny.Int
	_ = _144_i1
	_144_i1 = _dafny.Zero
	{
		for (p1) || (p1) {
			{
				if (_144_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_144_i1 = (_144_i1).Plus(_dafny.One)
				(globalState).F7 = _dafny.IntOfInt64(-8)
				_142_v0 = _142_v0
				var _145_v1 _dafny.Int
				_ = _145_v1
				_145_v1 = _dafny.IntOfInt64(859)
				var _146_v2 _dafny.Sequence
				_ = _146_v2
				_146_v2 = _dafny.SeqOf(_145_v1)
				var _147_v3 _dafny.Map
				_ = _147_v3
				_147_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _146_v2)
				var _148_v4 _dafny.Array
				_ = _148_v4
				var _nwElement0_0 _dafny.Map = _147_v3
				_ = _nwElement0_0
				var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(2))
				_ = _nw1
				(_nw1).ArraySet1(_nwElement0_0, 0)
				(_nw1).ArraySet1(_147_v3, 1)
				_148_v4 = _nw1
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_148_v4), 0))
				_ = _index0
				(_148_v4).ArraySet1(_147_v3, (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(937), _dafny.ArrayLen((_148_v4), 0))
				_ = _index1
				(_148_v4).ArraySet1(_147_v3, (_index1).Int())
				var _149_v5 _dafny.Map
				_ = _149_v5
				_149_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v1, p1)
				var _150_v6 _dafny.Array
				_ = _150_v6
				var _nwElement0_1 bool = p1
				_ = _nwElement0_1
				var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(8))
				_ = _nw2
				(_nw2).ArraySet1(_nwElement0_1, 0)
				(_nw2).ArraySet1((func() bool {
					if (_149_v5).Contains(_145_v1) {
						return (_149_v5).Get(_145_v1).(bool)
					}
					return p1
				})(), 1)
				(_nw2).ArraySet1(false, 2)
				(_nw2).ArraySet1(p1, 3)
				(_nw2).ArraySet1(p1, 4)
				(_nw2).ArraySet1(p1, 5)
				(_nw2).ArraySet1(p1, 6)
				(_nw2).ArraySet1(p1, 7)
				_150_v6 = _nw2
				var _151_v7 _dafny.Set
				_ = _151_v7
				_151_v7 = _dafny.SetOf(_145_v1)
				var _152_v8 D1
				_ = _152_v8
				_152_v8 = Companion_D1_.Create_DC2_(_145_v1, _145_v1, p1)
				var _153_v9 T2
				_ = _153_v9
				var _nw3 *C5 = New_C5_()
				_ = _nw3
				_nw3.Ctor__(p1, _150_v6, _151_v7, _152_v8, p1)
				_153_v9 = _nw3
				var _154_v10 _dafny.Sequence
				_ = _154_v10
				_154_v10 = _dafny.SeqOf(_153_v9)
				var _155_v11 _dafny.Map
				_ = _155_v11
				_155_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_v10, _145_v1)
				var _156_v12 _dafny.Map
				_ = _156_v12
				_156_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.IntOfInt64(499)).Cmp((_155_v11).Cardinality()) < 0)
				_156_v12 = (_156_v12).Update(p1, (func() bool {
					if (_156_v12).Contains(p1) {
						return (_156_v12).Get(p1).(bool)
					}
					return p1
				})())
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _157_v13 _dafny.Array
	_ = _157_v13
	var _nwElement0_2 bool = p1
	_ = _nwElement0_2
	var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(12))
	_ = _nw4
	(_nw4).ArraySet1(_nwElement0_2, 0)
	(_nw4).ArraySet1(p1, 1)
	(_nw4).ArraySet1(p1, 2)
	(_nw4).ArraySet1(true, 3)
	(_nw4).ArraySet1(p1, 4)
	(_nw4).ArraySet1(p1, 5)
	(_nw4).ArraySet1(p1, 6)
	(_nw4).ArraySet1(p1, 7)
	(_nw4).ArraySet1(p1, 8)
	(_nw4).ArraySet1(p1, 9)
	(_nw4).ArraySet1(p1, 10)
	(_nw4).ArraySet1(p1, 11)
	_157_v13 = _nw4
	var _158_v14 _dafny.Int
	_ = _158_v14
	_158_v14 = _dafny.IntOfInt64(-466)
	var _159_v15 _dafny.Set
	_ = _159_v15
	_159_v15 = _dafny.SetOf(_158_v14, _158_v14)
	var _160_v16 _dafny.Sequence
	_ = _160_v16
	_160_v16 = _dafny.SeqOf(p1, p1)
	var _161_v17 D1
	_ = _161_v17
	_161_v17 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_160_v16).Cardinality()), _158_v14, p1)
	var _162_v18 *C5
	_ = _162_v18
	var _nw5 *C5 = New_C5_()
	_ = _nw5
	_nw5.Ctor__(p1, _157_v13, _159_v15, _161_v17, p1)
	_162_v18 = _nw5
	var _163_v19 _dafny.Sequence
	_ = _163_v19
	_163_v19 = _dafny.SeqOf(_162_v18)
	var _164_v20 _dafny.Map
	_ = _164_v20
	_164_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (_163_v19).Select((Companion_Default___.SafeIndex(_158_v14, _dafny.IntOfUint32((_163_v19).Cardinality()))).Uint32()).(*C5))
	var _165_v21 T2
	_ = _165_v21
	var _nw6 *C5 = New_C5_()
	_ = _nw6
	_nw6.Ctor__(p1, _157_v13, _159_v15, _161_v17, p1)
	_165_v21 = _nw6
	var _166_v22 _dafny.Set
	_ = _166_v22
	_166_v22 = _dafny.SetOf(_165_v21)
	var _hi0 _dafny.Int = ((func() _dafny.Set {
		if _162_v18.F24 {
			return _166_v22
		}
		return _166_v22
	})()).Cardinality()
	_ = _hi0
	for _167_i2 := (_164_v20).Cardinality(); _167_i2.Cmp(_hi0) < 0; _167_i2 = _167_i2.Plus(_dafny.One) {
		var _168_v23 _dafny.Map
		_ = _168_v23
		_168_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_162_v18.F24, false)
		var _169_v24 _dafny.Sequence
		_ = _169_v24
		_169_v24 = _dafny.SeqOf(_168_v23)
		var _170_v25 _dafny.Map
		_ = _170_v25
		_170_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_165_v21).F13(), _dafny.SeqOf(_168_v23, _168_v23))
		var _171_v26 _dafny.Set
		_ = _171_v26
		_171_v26 = _dafny.SetOf(_160_v16, _dafny.SeqOf(false, (_165_v21).F13()), _160_v16)
		var _172_v27 _dafny.Set
		_ = _172_v27
		_172_v27 = _dafny.SetOf(_dafny.SeqOf((_162_v18).Fm52(_167_i2, _171_v26, globalState), _162_v18.F24))
		var _173_v28 *C0
		_ = _173_v28
		var _nw7 *C0 = New_C0_()
		_ = _nw7
		_nw7.Ctor__(!_dafny.Companion_Sequence_.Equal(_169_v24, (func() _dafny.Sequence {
			if (_170_v25).Contains(_162_v18.F24) {
				return (_170_v25).Get(_162_v18.F24).(_dafny.Sequence)
			}
			return _169_v24
		})()), _dafny.CodePoint('b'), _165_v21.F12(), (_162_v18).Fm52(_158_v14, _172_v27, globalState))
		_173_v28 = _nw7
		var _174_v29 _dafny.Sequence
		_ = _174_v29
		_174_v29 = _dafny.SeqOf(_173_v28)
		_173_v28 = (_174_v29).Select((Companion_Default___.SafeIndex((_173_v28).Fm16((_165_v21).F13(), _167_i2, globalState), _dafny.IntOfUint32((_174_v29).Cardinality()))).Uint32()).(*C0)
		var _175_v30 _dafny.Array
		_ = _175_v30
		var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw8
		_175_v30 = _nw8
		var _176_v31 _dafny.Array
		_ = _176_v31
		var _nwElement0_3 _dafny.Array = _175_v30
		_ = _nwElement0_3
		var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(24))
		_ = _nw9
		(_nw9).ArraySet1(_nwElement0_3, 0)
		(_nw9).ArraySet1(_175_v30, 1)
		(_nw9).ArraySet1(_175_v30, 2)
		(_nw9).ArraySet1(_175_v30, 3)
		(_nw9).ArraySet1(_175_v30, 4)
		(_nw9).ArraySet1(_175_v30, 5)
		(_nw9).ArraySet1(_175_v30, 6)
		(_nw9).ArraySet1(_175_v30, 7)
		(_nw9).ArraySet1(_175_v30, 8)
		(_nw9).ArraySet1(_175_v30, 9)
		(_nw9).ArraySet1(_175_v30, 10)
		(_nw9).ArraySet1(_175_v30, 11)
		(_nw9).ArraySet1(_175_v30, 12)
		(_nw9).ArraySet1(_175_v30, 13)
		(_nw9).ArraySet1(_175_v30, 14)
		(_nw9).ArraySet1(_175_v30, 15)
		(_nw9).ArraySet1(_175_v30, 16)
		(_nw9).ArraySet1(_175_v30, 17)
		(_nw9).ArraySet1(_175_v30, 18)
		(_nw9).ArraySet1(_175_v30, 19)
		(_nw9).ArraySet1(_175_v30, 20)
		(_nw9).ArraySet1(_175_v30, 21)
		(_nw9).ArraySet1(_175_v30, 22)
		(_nw9).ArraySet1(_175_v30, 23)
		_176_v31 = _nw9
		var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_176_v31), 0))
		_ = _index2
		(_176_v31).ArraySet1(_175_v30, (_index2).Int())
		var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(722), _dafny.ArrayLen((_176_v31), 0))
		_ = _index3
		(_176_v31).ArraySet1(_175_v30, (_index3).Int())
		var _177_v32 _dafny.Sequence
		_ = _177_v32
		_177_v32 = _dafny.UnicodeSeqOfUtf8Bytes("qrcl")
		_177_v32 = Companion_Default___.Fm20(globalState)
		(globalState).F7 = _158_v14
	}
	(_162_v18).F24 = (func() bool {
		if p1 {
			return !(_162_v18.F24)
		}
		return (p1) && (_162_v18.F24)
	})()
	var _178_v33 _dafny.Sequence
	_ = _178_v33
	_178_v33 = _dafny.UnicodeSeqOfUtf8Bytes("ulnyf")
	var _179_v34 _dafny.Map
	_ = _179_v34
	_179_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_178_v33).Cardinality()), _158_v14)
	(_162_v18).F24 = Companion_Default___.Fm30((_dafny.Zero).Minus(((func() _dafny.Map {
		if _162_v18.F24 {
			return _179_v34
		}
		return _179_v34
	})()).Cardinality()), globalState)
	var _180_v35 D1
	_ = _180_v35
	_180_v35 = Companion_D1_.Create_DC1_(_178_v33)
	var _181_v36 *C3
	_ = _181_v36
	var _nw10 *C3 = New_C3_()
	_ = _nw10
	_nw10.Ctor__((_165_v21).F13(), _180_v35, _162_v18.F24, _157_v13, _159_v15, _161_v17, (func() bool {
		if p1 {
			return _162_v18.F24
		}
		return p1
	})())
	_181_v36 = _nw10
	var _pat_let_tv0 = _181_v36
	_ = _pat_let_tv0
	var _pat_let_tv1 = _162_v18
	_ = _pat_let_tv1
	var _pat_let_tv2 = _165_v21
	_ = _pat_let_tv2
	var _pat_let_tv3 = _181_v36
	_ = _pat_let_tv3
	var _pat_let_tv4 = _162_v18
	_ = _pat_let_tv4
	var _pat_let_tv5 = _181_v36
	_ = _pat_let_tv5
	var _pat_let_tv6 = _160_v16
	_ = _pat_let_tv6
	var _pat_let_tv7 = _158_v14
	_ = _pat_let_tv7
	var _pat_let_tv8 = _160_v16
	_ = _pat_let_tv8
	r0 = func(_source6 D12) _dafny.MultiSet {
		if _source6.Is_DC30() {
			var _182___mcc_h0 _dafny.Array = _source6.Get_().(D12_DC30).Cf58
			_ = _182___mcc_h0
			var _183___mcc_h1 bool = _source6.Get_().(D12_DC30).Cf59
			_ = _183___mcc_h1
			var _184_cf59 bool = _183___mcc_h1
			_ = _184_cf59
			var _185_cf58 _dafny.Array = _182___mcc_h0
			_ = _185_cf58
			var _186_dt__update__tmp_h0 _dafny.MultiSet = _dafny.MultiSetOf((_pat_let_tv0).F18(), _pat_let_tv1.F24)
			_ = _186_dt__update__tmp_h0
			var _187_dt__update_hcf0_h0 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.SeqOf((_pat_let_tv2).F13()))
			_ = _187_dt__update_hcf0_h0
			return _187_dt__update_hcf0_h0
		} else if _source6.Is_DC31() {
			var _188___mcc_h2 bool = _source6.Get_().(D12_DC31).Cf60
			_ = _188___mcc_h2
			var _189___mcc_h3 _dafny.Array = _source6.Get_().(D12_DC31).Cf61
			_ = _189___mcc_h3
			var _190___mcc_h4 bool = _source6.Get_().(D12_DC31).Cf62
			_ = _190___mcc_h4
			var _191___mcc_h5 _dafny.CodePoint = _source6.Get_().(D12_DC31).Cf63
			_ = _191___mcc_h5
			var _192___mcc_h6 _dafny.MultiSet = _source6.Get_().(D12_DC31).Cf64
			_ = _192___mcc_h6
			var _193_cf64 _dafny.MultiSet = _192___mcc_h6
			_ = _193_cf64
			var _194_cf63 _dafny.CodePoint = _191___mcc_h5
			_ = _194_cf63
			var _195_cf62 bool = _190___mcc_h4
			_ = _195_cf62
			var _196_cf61 _dafny.Array = _189___mcc_h3
			_ = _196_cf61
			var _197_cf60 bool = _188___mcc_h2
			_ = _197_cf60
			return _dafny.MultiSetOf((_pat_let_tv3).F18(), _197_cf60)
		} else if _source6.Is_DC32() {
			var _198___mcc_h7 bool = _source6.Get_().(D12_DC32).Cf65
			_ = _198___mcc_h7
			var _199___mcc_h8 _dafny.Int = _source6.Get_().(D12_DC32).Cf66
			_ = _199___mcc_h8
			var _200___mcc_h9 bool = _source6.Get_().(D12_DC32).Cf67
			_ = _200___mcc_h9
			var _201___mcc_h10 bool = _source6.Get_().(D12_DC32).Cf68
			_ = _201___mcc_h10
			var _202___mcc_h11 _dafny.Int = _source6.Get_().(D12_DC32).Cf69
			_ = _202___mcc_h11
			var _203_cf69 _dafny.Int = _202___mcc_h11
			_ = _203_cf69
			var _204_cf68 bool = _201___mcc_h10
			_ = _204_cf68
			var _205_cf67 bool = _200___mcc_h9
			_ = _205_cf67
			var _206_cf66 _dafny.Int = _199___mcc_h8
			_ = _206_cf66
			var _207_cf65 bool = _198___mcc_h7
			_ = _207_cf65
			return _dafny.MultiSetOf(!(true), _205_cf67, _pat_let_tv4.F24)
		} else {
			var _208___mcc_h12 _dafny.Map = _source6.Get_().(D12_DC29).Cf57
			_ = _208___mcc_h12
			var _209_cf57 _dafny.Map = _208___mcc_h12
			_ = _209_cf57
			return _dafny.MultiSetOf((_pat_let_tv5).F18(), (_pat_let_tv6).Select((Companion_Default___.SafeIndex(_pat_let_tv7, _dafny.IntOfUint32((_pat_let_tv8).Cardinality()))).Uint32()).(bool))
		}
	}(Companion_D12_.Create_DC32_(p1, _dafny.IntOfUint32((_178_v33).Cardinality()), (_181_v36).F18(), _162_v18.F24, _158_v14))
	return r0
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _210_v1 _dafny.Array
	_ = _210_v1
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(2)
	_ = _len0_0
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.Map = func(_211_i0 _dafny.Int) _dafny.Map {
			return func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(633), (_dafny.MultiSetOf(false)).Cardinality())).Keys().Elements()); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _212_v0 _dafny.Int
					_212_v0 = interface{}(_compr_23).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(633), (_dafny.MultiSetOf(false)).Cardinality())).Contains(_212_v0) {
						_coll23.Add(Companion_Default___.SafeDivisionInt(_212_v0, _dafny.IntOfInt64(805)), false)
					}
				}
				return _coll23.ToMap()
			}()
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw11 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw11).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw11).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_210_v1 = _nw11
	var _213_v2 bool
	_ = _213_v2
	_213_v2 = true
	var _214_v3 _dafny.Sequence
	_ = _214_v3
	_214_v3 = _dafny.SeqOf(_213_v2)
	var _215_v4 _dafny.Sequence
	_ = _215_v4
	_215_v4 = _dafny.UnicodeSeqOfUtf8Bytes("dgxao")
	var _216_v5 _dafny.Array
	_ = _216_v5
	var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
	_ = _nw12
	_216_v5 = _nw12
	var _217_v6 _dafny.MultiSet
	_ = _217_v6
	_217_v6 = _dafny.MultiSetOf(_216_v5, _216_v5)
	var _218_v7 _dafny.Map
	_ = _218_v7
	_218_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v2, _215_v4)
	var _219_v8 _dafny.CodePoint
	_ = _219_v8
	_219_v8 = _dafny.CodePoint('g')
	var _220_v9 _dafny.Int
	_ = _220_v9
	_220_v9 = _dafny.IntOfInt64(336)
	var _221_v10 _dafny.Map
	_ = _221_v10
	_221_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v9, _215_v4)
	var _222_v11 D1
	_ = _222_v11
	_222_v11 = Companion_D1_.Create_DC1_(_215_v4)
	var _223_v12 _dafny.Array
	_ = _223_v12
	var _nwElement0_4 _dafny.Sequence = (func() _dafny.Sequence {
		if (_218_v7).Contains(_213_v2) {
			return (_218_v7).Get(_213_v2).(_dafny.Sequence)
		}
		return _215_v4
	})()
	_ = _nwElement0_4
	var _nw13 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(28))
	_ = _nw13
	(_nw13).ArraySet1(_nwElement0_4, 0)
	(_nw13).ArraySet1(_215_v4, 1)
	(_nw13).ArraySet1(_215_v4, 2)
	(_nw13).ArraySet1(_215_v4, 3)
	(_nw13).ArraySet1(_215_v4, 4)
	(_nw13).ArraySet1(_215_v4, 5)
	(_nw13).ArraySet1(_215_v4, 6)
	(_nw13).ArraySet1(_215_v4, 7)
	(_nw13).ArraySet1(_215_v4, 8)
	(_nw13).ArraySet1(_215_v4, 9)
	(_nw13).ArraySet1(_215_v4, 10)
	(_nw13).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("yauffnrew"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("yauffnrew")).Cardinality()))).Uint32(), _219_v8), 11)
	(_nw13).ArraySet1(_215_v4, 12)
	(_nw13).ArraySet1(_215_v4, 13)
	(_nw13).ArraySet1((func() _dafny.Sequence {
		if (_221_v10).Contains(_220_v9) {
			return (_221_v10).Get(_220_v9).(_dafny.Sequence)
		}
		return _215_v4
	})(), 14)
	(_nw13).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(51))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg38 _dafny.Int) interface{} {
			return coer38(arg38)
		}
	}((func(_224_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_225_i1 _dafny.Int) _dafny.CodePoint {
			return _224_v8
		}
	})(_219_v8))), 15)
	(_nw13).ArraySet1(_215_v4, 16)
	(_nw13).ArraySet1(_215_v4, 17)
	(_nw13).ArraySet1((_222_v11).Dtor_cf1(), 18)
	(_nw13).ArraySet1(_215_v4, 19)
	(_nw13).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("tqk"), 20)
	(_nw13).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dgjly"), 21)
	(_nw13).ArraySet1(_215_v4, 22)
	(_nw13).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(658))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg39 _dafny.Int) interface{} {
			return coer39(arg39)
		}
	}((func(_226_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_227_i2 _dafny.Int) _dafny.CodePoint {
			return _226_v8
		}
	})(_219_v8))), 23)
	(_nw13).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dffircy"), 24)
	(_nw13).ArraySet1(_215_v4, 25)
	(_nw13).ArraySet1(_215_v4, 26)
	(_nw13).ArraySet1(_215_v4, 27)
	_223_v12 = _nw13
	var _228_globalState *GlobalState
	_ = _228_globalState
	var _nw14 *GlobalState = New_GlobalState_()
	_ = _nw14
	_nw14.Ctor__(_210_v1, true, _214_v3, _215_v4, true, false, _dafny.IntOfInt64(142), _dafny.IntOfInt64(530), _dafny.Companion_Sequence_.Concatenate(_215_v4, _215_v4), _217_v6, _223_v12)
	_228_globalState = _nw14
	var _229_v13 _dafny.MultiSet
	_ = _229_v13
	var _out0 _dafny.MultiSet
	_ = _out0
	_out0 = Companion_Default___.M0(_dafny.CodePoint('g'), _213_v2, _228_globalState)
	_229_v13 = _out0
	var _230_v14 _dafny.Array
	_ = _230_v14
	var _nwElement0_5 bool = !(_213_v2)
	_ = _nwElement0_5
	var _nw15 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(6))
	_ = _nw15
	(_nw15).ArraySet1(_nwElement0_5, 0)
	(_nw15).ArraySet1(!(_213_v2), 1)
	(_nw15).ArraySet1(_213_v2, 2)
	(_nw15).ArraySet1(!(_213_v2), 3)
	(_nw15).ArraySet1(true, 4)
	(_nw15).ArraySet1(_213_v2, 5)
	_230_v14 = _nw15
	var _231_v15 D2
	_ = _231_v15
	_231_v15 = Companion_D2_.Create_DC3_(_230_v14)
	var _232_v16 _dafny.Sequence
	_ = _232_v16
	_232_v16 = _dafny.SeqOf(_230_v14)
	var _233_v17 _dafny.Set
	_ = _233_v17
	_233_v17 = _dafny.SetOf(_213_v2)
	var _234_v18 _dafny.MultiSet
	_ = _234_v18
	_234_v18 = _dafny.MultiSetOf((_233_v17).Cardinality())
	var _235_v19 _dafny.Map
	_ = _235_v19
	_235_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_234_v18, _213_v2)
	var _236_v20 D2
	_ = _236_v20
	_236_v20 = Companion_D2_.Create_DC4_(_216_v5, _213_v2, _220_v9, false, _230_v14)
	var _237_v21 _dafny.Array
	_ = _237_v21
	var _nwElement0_6 _dafny.Array = (func() _dafny.Array {
		if _213_v2 {
			return _230_v14
		}
		return _230_v14
	})()
	_ = _nwElement0_6
	var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(29))
	_ = _nw16
	(_nw16).ArraySet1(_nwElement0_6, 0)
	(_nw16).ArraySet1(_230_v14, 1)
	(_nw16).ArraySet1(_230_v14, 2)
	(_nw16).ArraySet1(_230_v14, 3)
	(_nw16).ArraySet1(_230_v14, 4)
	(_nw16).ArraySet1(_230_v14, 5)
	(_nw16).ArraySet1(_230_v14, 6)
	(_nw16).ArraySet1(_230_v14, 7)
	(_nw16).ArraySet1(_230_v14, 8)
	(_nw16).ArraySet1(_230_v14, 9)
	(_nw16).ArraySet1(_230_v14, 10)
	(_nw16).ArraySet1(_230_v14, 11)
	(_nw16).ArraySet1(_230_v14, 12)
	(_nw16).ArraySet1(_230_v14, 13)
	(_nw16).ArraySet1((func() _dafny.Array {
		if _213_v2 {
			return _230_v14
		}
		return _230_v14
	})(), 14)
	(_nw16).ArraySet1((_231_v15).Dtor_cf5(), 15)
	(_nw16).ArraySet1(_230_v14, 16)
	(_nw16).ArraySet1(_230_v14, 17)
	(_nw16).ArraySet1(_230_v14, 18)
	(_nw16).ArraySet1(_230_v14, 19)
	(_nw16).ArraySet1(_230_v14, 20)
	(_nw16).ArraySet1(_230_v14, 21)
	(_nw16).ArraySet1(_230_v14, 22)
	(_nw16).ArraySet1(_230_v14, 23)
	(_nw16).ArraySet1(_230_v14, 24)
	(_nw16).ArraySet1((_232_v16).Select((Companion_Default___.SafeIndex((_235_v19).Cardinality(), _dafny.IntOfUint32((_232_v16).Cardinality()))).Uint32()).(_dafny.Array), 25)
	(_nw16).ArraySet1(_230_v14, 26)
	(_nw16).ArraySet1(_230_v14, 27)
	(_nw16).ArraySet1((_236_v20).Dtor_cf10(), 28)
	_237_v21 = _nw16
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))
	_ = _index4
	(_237_v21).ArraySet1(_230_v14, (_index4).Int())
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))
	_ = _index5
	var _rhs0 _dafny.Array = _230_v14
	_ = _rhs0
	var _rhs1 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("egift")).Cardinality()), _dafny.IntOfInt64(294))
	_ = _rhs1
	var _lhs0 _dafny.Array = _237_v21
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))
	_ = _lhs1
	var _lhs2 *GlobalState = _228_globalState
	_ = _lhs2
	(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
	_lhs2.F7 = _rhs1
	var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
	_ = _index6
	(_216_v5).ArraySet1(_220_v9, (_index6).Int())
	var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
	_ = _index7
	(_216_v5).ArraySet1(_220_v9, (_index7).Int())
	var _238_v22 _dafny.Map
	_ = _238_v22
	_238_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v2, _214_v3)
	var _hi1 _dafny.Int = (_220_v9).Times(Companion_Default___.Fm0((func() _dafny.Sequence {
		if (_238_v22).Contains(_213_v2) {
			return (_238_v22).Get(_213_v2).(_dafny.Sequence)
		}
		return _214_v3
	})(), _213_v2, _213_v2, (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _228_globalState))
	_ = _hi1
	for _239_i3 := (_233_v17).Cardinality(); _239_i3.Cmp(_hi1) < 0; _239_i3 = _239_i3.Plus(_dafny.One) {
		_213_v2 = _213_v2
		var _source7 D1 = Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Concatenate(_215_v4, _215_v4))
		_ = _source7
		if _source7.Is_DC2() {
			var _240___mcc_h0 _dafny.Int = _source7.Get_().(D1_DC2).Cf2
			_ = _240___mcc_h0
			var _241___mcc_h1 _dafny.Int = _source7.Get_().(D1_DC2).Cf3
			_ = _241___mcc_h1
			var _242___mcc_h2 bool = _source7.Get_().(D1_DC2).Cf4
			_ = _242___mcc_h2
			var _243_cf4 bool = _242___mcc_h2
			_ = _243_cf4
			var _244_cf3 _dafny.Int = _241___mcc_h1
			_ = _244_cf3
			var _245_cf2 _dafny.Int = _240___mcc_h0
			_ = _245_cf2
			var _246_v23 _dafny.Map
			_ = _246_v23
			_246_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_215_v4).Cardinality()), _230_v14)
			var _247_v24 D1
			_ = _247_v24
			_247_v24 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(87), _dafny.IntOfUint32((_215_v4).Cardinality()), _243_cf4)
			var _248_v25 *C2
			_ = _248_v25
			var _nw17 *C2 = New_C2_()
			_ = _nw17
			_nw17.Ctor__(_244_cf3, !(!(_243_cf4)), _222_v11, _213_v2, (func() _dafny.Array {
				if _243_cf4 {
					return _230_v14
				}
				return (func() _dafny.Array {
					if (_246_v23).Contains(_245_cf2) {
						return (_246_v23).Get(_245_cf2).(_dafny.Array)
					}
					return _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
				})()
			})(), _dafny.SetOf((_dafny.SetOf(_243_cf4, _243_cf4)).Cardinality()), _247_v24, _243_cf4)
			_248_v25 = _nw17
			var _249_v26 _dafny.Map
			_ = _249_v26
			_249_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_243_cf4, _220_v9)
			_249_v26 = (_249_v26).Update(!((_248_v25).F22()), (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))
			_244_cf3 = (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)
			_216_v5 = _216_v5
		} else {
			var _250___mcc_h3 _dafny.Sequence = _source7.Get_().(D1_DC1).Cf1
			_ = _250___mcc_h3
			var _251_cf1 _dafny.Sequence = _250___mcc_h3
			_ = _251_cf1
			var _252_v27 *C5
			_ = _252_v27
			var _nw18 *C5 = New_C5_()
			_ = _nw18
			_nw18.Ctor__(_213_v2, _230_v14, Companion_Default___.Fm64(_228_globalState), Companion_Default___.Fm37((_dafny.Zero).Minus((_dafny.Zero).Minus((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))), _239_i3, (_dafny.Zero).Minus(_220_v9), (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _228_globalState), ((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)).Cmp(_239_i3) > 0)
			_252_v27 = _nw18
			var _253_v29 D1
			_ = _253_v29
			_253_v29 = Companion_D1_.Create_DC2_(_239_i3, _dafny.IntOfUint32((_215_v4).Cardinality()), _213_v2)
			var _254_v30 *C2
			_ = _254_v30
			var _nw19 *C2 = New_C2_()
			_ = _nw19
			_nw19.Ctor__(_239_i3, _252_v27.F24, Companion_D1_.Create_DC1_(_251_cf1), _213_v2, _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int())), func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(536), _dafny.IntOfInt64(571))); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _255_v28 _dafny.Int
					_255_v28 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(536)).Cmp(_255_v28) <= 0) && ((_255_v28).Cmp(_dafny.IntOfInt64(571)) < 0) {
						_coll24.Add((_255_v28).Plus(_220_v9))
					}
				}
				return _coll24.ToSet()
			}(), _253_v29, _252_v27.F24)
			_254_v30 = _nw19
			var _256_v31 _dafny.Map
			_ = _256_v31
			_256_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_252_v27.F24, _254_v30)
			var _257_v32 _dafny.Sequence
			_ = _257_v32
			_257_v32 = _dafny.SeqOf((_256_v31).Cardinality(), (_254_v30).F21())
			var _258_v33 _dafny.Array
			_ = _258_v33
			var _nwElement0_7 _dafny.CodePoint = (Companion_D11_.Create_DC27_(true, _219_v8, _220_v9, (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _257_v32)).Dtor_cf52()
			_ = _nwElement0_7
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(5))
			_ = _nw20
			(_nw20).ArraySet1CodePoint(_nwElement0_7, 0)
			(_nw20).ArraySet1CodePoint(_219_v8, 1)
			(_nw20).ArraySet1CodePoint((_215_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-56), _dafny.IntOfUint32((_215_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), 2)
			(_nw20).ArraySet1CodePoint(_219_v8, 3)
			(_nw20).ArraySet1CodePoint(_219_v8, 4)
			_258_v33 = _nw20
			_258_v33 = _258_v33
			var _259_v34 _dafny.Array
			_ = _259_v34
			var _len0_1 _dafny.Int = _dafny.One
			_ = _len0_1
			var _nw21 _dafny.Array
			_ = _nw21
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw21 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.MultiSet = (func(_260_v30 *C2, _261_v32 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
					return func(_262_i4 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(Companion_D19_.Create_DC50_((_260_v30).F22(), _dafny.IntOfInt64(-629), _261_v32), Companion_D19_.Create_DC50_((_260_v30).F22(), (_260_v30).F21(), _261_v32))
					}
				})(_254_v30, _257_v32)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw21 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw21).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw21).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_259_v34 = _nw21
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))
			_ = _index8
			var _rhs2 _dafny.Array = _216_v5
			_ = _rhs2
			var _rhs3 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
			_ = _rhs3
			var _rhs4 _dafny.Int = (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)
			_ = _rhs4
			var _rhs5 _dafny.Array = _259_v34
			_ = _rhs5
			var _lhs3 _dafny.Array = _237_v21
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))
			_ = _lhs4
			var _lhs5 *GlobalState = _228_globalState
			_ = _lhs5
			_216_v5 = _rhs2
			(_lhs3).ArraySet1(_rhs3, (_lhs4).Int())
			_lhs5.F7 = _rhs4
			_259_v34 = _rhs5
			var _263_v35 _dafny.Int
			_ = _263_v35
			var _out1 _dafny.Int
			_ = _out1
			_out1 = (_254_v30).M6((_254_v30).F22(), _228_globalState)
			_263_v35 = _out1
		}
		var _264_v36 _dafny.Set
		_ = _264_v36
		_264_v36 = _dafny.SetOf(_220_v9)
		var _265_v37 D19
		_ = _265_v37
		_265_v37 = Companion_D19_.Create_DC49_(_264_v36)
		var _pat_let_tv9 = _264_v36
		_ = _pat_let_tv9
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
		_ = _arr0
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
		_ = _index9
		(_arr0).ArraySet1(((func(_pat_let0_0 D19) D19 {
			return func(_266_dt__update__tmp_h0 D19) D19 {
				return func(_pat_let1_0 _dafny.Set) D19 {
					return func(_267_dt__update_hcf93_h0 _dafny.Set) D19 {
						return Companion_D19_.Create_DC49_(_267_dt__update_hcf93_h0)
					}(_pat_let1_0)
				}(_pat_let_tv9)
			}(_pat_let0_0)
		}(_265_v37)).Dtor_cf93()).IsDisjointFrom(_264_v36), (_index9).Int())
		var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _index10
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
		_ = _arr1
		var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
		_ = _index11
		var _rhs6 _dafny.Int = _220_v9
		_ = _rhs6
		var _rhs7 _dafny.Int = (func() _dafny.Int {
			if _213_v2 {
				return _220_v9
			}
			return _dafny.IntOfInt64(-487)
		})()
		_ = _rhs7
		var _rhs8 bool = !(true)
		_ = _rhs8
		var _rhs9 bool = false
		_ = _rhs9
		var _lhs6 *GlobalState = _228_globalState
		_ = _lhs6
		var _lhs7 _dafny.Array = _216_v5
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _lhs8
		var _lhs9 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
		_ = _lhs9
		var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
		_ = _lhs10
		_lhs6.F7 = _rhs6
		(_lhs7).ArraySet1(_rhs7, (_lhs8).Int())
		_213_v2 = _rhs8
		(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _index12
		(_216_v5).ArraySet1(_220_v9, (_index12).Int())
	}
	var _268_i5 _dafny.Int
	_ = _268_i5
	_268_i5 = _dafny.Zero
	{
		for false {
			{
				if (_268_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_268_i5 = (_268_i5).Plus(_dafny.One)
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
				_ = _arr2
				var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
				_ = _index13
				(_arr2).ArraySet1(_213_v2, (_index13).Int())
				var _arr3 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
				_ = _arr3
				var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
				_ = _index14
				(_arr3).ArraySet1(_213_v2, (_index14).Int())
				if _213_v2 {
					var _269_v38 _dafny.Map
					_ = _269_v38
					_269_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _220_v9)
					var _270_v40 D1
					_ = _270_v40
					_270_v40 = Companion_D1_.Create_DC2_(_220_v9, (_dafny.MultiSetOf(_214_v3)).Cardinality(), !((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool)))
					var _271_v41 *C6
					_ = _271_v41
					var _nw22 *C6 = New_C6_()
					_ = _nw22
					_nw22.Ctor__(_222_v11, Companion_Default___.Fm17(_228_globalState), _230_v14, func() _dafny.Set {
						var _coll25 = _dafny.NewBuilder()
						_ = _coll25
						for _iter26 := _dafny.Iterate((_269_v38).Keys().Elements()); ; {
							_compr_25, _ok26 := _iter26()
							if !_ok26 {
								break
							}
							var _272_v39 _dafny.Int
							_272_v39 = interface{}(_compr_25).(_dafny.Int)
							if (_269_v38).Contains(_272_v39) {
								_coll25.Add((_272_v39).Times((_dafny.Zero).Minus(_dafny.IntOfInt64(-879))))
							}
						}
						return _coll25.ToSet()
					}(), _270_v40, (_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool))
					_271_v41 = _nw22
					_220_v9 = (_233_v17).Cardinality()
					var _273_v42 _dafny.MultiSet
					_ = _273_v42
					_273_v42 = _dafny.MultiSetOf(_230_v14, _230_v14)
					var _arr4 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
					_ = _arr4
					var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
					_ = _index15
					(_arr4).ArraySet1(((_273_v42).Difference(_273_v42)).IsSubsetOf(_273_v42), (_index15).Int())
					var _274_v43 _dafny.Array
					_ = _274_v43
					var _nw23 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
					_ = _nw23
					_274_v43 = _nw23
					_274_v43 = _274_v43
					var _275_v44 _dafny.Map
					_ = _275_v44
					_275_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v8, _220_v9)
					var _276_v45 _dafny.Sequence
					_ = _276_v45
					_276_v45 = _dafny.SeqOf(_275_v44)
					_276_v45 = _276_v45
				} else {
					_215_v4 = _dafny.Companion_Sequence_.Concatenate(_215_v4, _215_v4)
					var _277_v46 *C4
					_ = _277_v46
					var _nw24 *C4 = New_C4_()
					_ = _nw24
					_nw24.Ctor__()
					_277_v46 = _nw24
					var _278_v47 _dafny.Map
					_ = _278_v47
					_278_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v9, !((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool)))
					var _279_v48 _dafny.Map
					_ = _279_v48
					_279_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool), _278_v47)
					var _280_v49 _dafny.Map
					_ = _280_v49
					_280_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v46, ((func() _dafny.Map {
						if (_279_v48).Contains(_213_v2) {
							return (_279_v48).Get(_213_v2).(_dafny.Map)
						}
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v9, (_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool))
					})()).Cardinality())
					var _281_v50 D1
					_ = _281_v50
					_281_v50 = Companion_D1_.Create_DC2_((_dafny.Zero).Minus((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)), (_dafny.MultiSetFromSeq(_dafny.SeqOf(_220_v9))).Cardinality(), _213_v2)
					var _282_v51 *C0
					_ = _282_v51
					var _nw25 *C0 = New_C0_()
					_ = _nw25
					_nw25.Ctor__((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool), (_215_v4).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_280_v49).Contains(_277_v46) {
							return (_280_v49).Get(_277_v46).(_dafny.Int)
						}
						return _dafny.IntOfInt64(86)
					})(), _dafny.IntOfUint32((_215_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _281_v50, ((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool)) == (true))
					_282_v51 = _nw25
					_282_v51 = _282_v51
					var _283_v52 _dafny.Map
					_ = _283_v52
					_283_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v9, _dafny.IntOfInt64(-643))
					var _284_v53 _dafny.Sequence
					_ = _284_v53
					_284_v53 = _dafny.SeqOf((_dafny.Zero).Minus((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)), (_283_v52).Cardinality())
					var _285_v54 _dafny.Set
					_ = _285_v54
					_285_v54 = _dafny.SetOf((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))
					var _286_v55 *C3
					_ = _286_v55
					var _nw26 *C3 = New_C3_()
					_ = _nw26
					_nw26.Ctor__(_282_v51.F19, _222_v11, Companion_Default___.Fm30(_dafny.IntOfUint32((_284_v53).Cardinality()), _228_globalState), _230_v14, _285_v54, _281_v50, (_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))).Int()).(bool))
					_286_v55 = _nw26
					var _287_v56 _dafny.Sequence
					_ = _287_v56
					_287_v56 = _dafny.SeqOf(_286_v55, _286_v55, _286_v55)
					_286_v55 = (_287_v56).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(830), _dafny.IntOfUint32((_287_v56).Cardinality()))).Uint32()).(*C3)
					var _288_v57 _dafny.MultiSet
					_ = _288_v57
					_288_v57 = _dafny.MultiSetOf(_213_v2)
					(_282_v51).F19 = ((_288_v57).Difference(_288_v57)).IsDisjointFrom((_288_v57).Intersection(_288_v57))
					(_228_globalState).F2 = _214_v3
				}
				var _arr5 _dafny.Array = _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))
				_ = _arr5
				var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int()))), 0))
				_ = _index16
				(_arr5).ArraySet1(!(!(Companion_Default___.Fm3((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _228_globalState))), (_index16).Int())
				_213_v2 = false
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _289_v58 _dafny.Map
	_ = _289_v58
	_289_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v2, _213_v2)
	var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
	_ = _index17
	(_230_v14).ArraySet1((func() bool {
		if (_289_v58).Contains(_213_v2) {
			return (_289_v58).Get(_213_v2).(bool)
		}
		return _213_v2
	})(), (_index17).Int())
	var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
	_ = _index18
	(_230_v14).ArraySet1(((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_213_v2)).Cardinality())) == 0, (_index18).Int())
	_289_v58 = (_289_v58).Update(_213_v2, ((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool)) == (_213_v2))
	var _hi2 _dafny.Int = _220_v9
	_ = _hi2
	for _290_i6 := _220_v9; _290_i6.Cmp(_hi2) < 0; _290_i6 = _290_i6.Plus(_dafny.One) {
		_213_v2 = _dafny.Companion_Sequence_.IsPrefixOf(_232_v16, _232_v16)
		var _291_v59 _dafny.MultiSet
		_ = _291_v59
		var _out2 _dafny.MultiSet
		_ = _out2
		_out2 = Companion_Default___.M0(_219_v8, Companion_Default___.Fm17(_228_globalState), _228_globalState)
		_291_v59 = _out2
		var _292_v60 _dafny.Array
		_ = _292_v60
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
		_ = _nw27
		_292_v60 = _nw27
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_292_v60), 0))
		_ = _index19
		(_292_v60).ArraySet1(_216_v5, (_index19).Int())
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_292_v60), 0))
		_ = _index20
		(_292_v60).ArraySet1(_216_v5, (_index20).Int())
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_292_v60), 0))
		_ = _index21
		(_292_v60).ArraySet1(_216_v5, (_index21).Int())
	}
	var _293_v61 D19
	_ = _293_v61
	_293_v61 = Companion_D19_.Create_DC50_((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _dafny.SeqOf((Companion_Default___.Fm24(_dafny.IntOfInt64(997), _228_globalState)).Cardinality(), _220_v9))
	var _294_v62 _dafny.MultiSet
	_ = _294_v62
	var _out3 _dafny.MultiSet
	_ = _out3
	_out3 = Companion_Default___.M0(_219_v8, (_293_v61).Dtor_cf94(), _228_globalState)
	_294_v62 = _out3
	var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
	_ = _index22
	var _rhs10 _dafny.Int = Companion_Default___.Fm0((func() _dafny.Sequence {
		if Companion_Default___.Fm17(_228_globalState) {
			return _214_v3
		}
		return _214_v3
	})(), (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), !((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool)), Companion_Default___.Fm0(_214_v3, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _dafny.IntOfInt64(-259), _228_globalState), _228_globalState)
	_ = _rhs10
	var _rhs11 bool = Companion_Default___.Fm3((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _228_globalState)
	_ = _rhs11
	var _rhs12 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _213_v2)).Cardinality()
	_ = _rhs12
	var _lhs11 *GlobalState = _228_globalState
	_ = _lhs11
	var _lhs12 _dafny.Array = _230_v14
	_ = _lhs12
	var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
	_ = _lhs13
	_lhs11.F7 = _rhs10
	(_lhs12).ArraySet1(_rhs11, (_lhs13).Int())
	_220_v9 = _rhs12
	var _295_v63 _dafny.Array
	_ = _295_v63
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_2
	var _nw28 _dafny.Array
	_ = _nw28
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw28 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Map = (func(_296_v2 bool, _297_v14 _dafny.Array) func(_dafny.Int) _dafny.Map {
			return func(_298_i8 _dafny.Int) _dafny.Map {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_296_v2)), (_297_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_297_v14), 0))).Int()).(bool))
			}
		})(_213_v2, _230_v14)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw28 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw28).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw28).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_295_v63 = _nw28
	var _ingredients0 = _dafny.NewBuilder()
	_ = _ingredients0
	for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_295_v63), 0))); ; {
		_guard_loop_1, _ok27 := _iter27()
		if !_ok27 {
			break
		}
		var _299_i7 _dafny.Int
		_299_i7 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_299_i7).Sign() != -1) && ((_299_i7).Cmp(_dafny.ArrayLen((_295_v63), 0)) < 0)) {
			_ingredients0.Add(_dafny.TupleOf(_295_v63, (_299_i7).Int(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_233_v17).Equals(_233_v17), (func() bool {
				if (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool) {
					return false
				}
				return _213_v2
			})())))
		}
	}
	for _iter28 := _dafny.Iterate(_ingredients0); ; {
		_tup0, _ok28 := _iter28()
		if !_ok28 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup0.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup0.(_dafny.Tuple)).IndexInt(2)), (*(_tup0.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	(_228_globalState).F7 = ((_dafny.Zero).Minus((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))).Times(((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)).Times((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)))
	var _ingredients1 = _dafny.NewBuilder()
	_ = _ingredients1
	for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_230_v14), 0))); ; {
		_guard_loop_2, _ok29 := _iter29()
		if !_ok29 {
			break
		}
		var _300_i9 _dafny.Int
		_300_i9 = interface{}(_guard_loop_2).(_dafny.Int)
		if (true) && (((_300_i9).Sign() != -1) && ((_300_i9).Cmp(_dafny.ArrayLen((_230_v14), 0)) < 0)) {
			_ingredients1.Add(_dafny.TupleOf(_230_v14, (_300_i9).Int(), ((func() _dafny.Set {
				var _coll26 = _dafny.NewBuilder()
				_ = _coll26
				for _iter30 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_301_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_302_i10 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus((_301_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_301_v5), 0))).Int()).(_dafny.Int))
					}
				})(_216_v5)))).Elements()); ; {
					_compr_26, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _303_v64 _dafny.Int
					_303_v64 = interface{}(_compr_26).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer41 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg41 _dafny.Int) interface{} {
							return coer41(arg41)
						}
					}((func(_304_v5 _dafny.Array) func(_dafny.Int) _dafny.Int {
						return func(_302_i10 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus((_304_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_304_v5), 0))).Int()).(_dafny.Int))
						}
					})(_216_v5))), _303_v64) {
						_coll26.Add((_303_v64).Plus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(51))).Cardinality())))
					}
				}
				return _coll26.ToSet()
			}()).Union(_dafny.SetOf(_dafny.IntOfInt64(-321)))).IsProperSubsetOf((_dafny.SetOf((_289_v58).Cardinality())).Union(_dafny.SetOf(_220_v9)))))
		}
	}
	for _iter31 := _dafny.Iterate(_ingredients1); ; {
		_tup1, _ok31 := _iter31()
		if !_ok31 {
			break
		}
		(_dafny.ArrayCastTo((*(_tup1.(_dafny.Tuple)).IndexInt(0)))).ArraySet1((*(_tup1.(_dafny.Tuple)).IndexInt(2)), (*(_tup1.(_dafny.Tuple)).IndexInt(1)).(int))
	}
	var _305_v65 _dafny.Map
	_ = _305_v65
	_305_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_213_v2, _220_v9)
	var _hi3 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))).Cardinality())
	_ = _hi3
	for _306_i11 := Companion_Default___.SafeModuloInt((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), (_305_v65).Cardinality()); _306_i11.Cmp(_hi3) < 0; _306_i11 = _306_i11.Plus(_dafny.One) {
		var _307_v67 _dafny.Sequence
		_ = _307_v67
		_307_v67 = _dafny.SeqOf(_220_v9, (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int))
		var _308_v68 D13
		_ = _308_v68
		_308_v68 = Companion_D13_.Create_DC34_(_220_v9, Companion_Default___.Fm65(_213_v2, _307_v67, _219_v8, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _228_globalState))
		var _309_v69 D1
		_ = _309_v69
		_309_v69 = Companion_D1_.Create_DC2_(_220_v9, (_308_v68).Dtor_cf71(), _213_v2)
		var _310_v70 T3
		_ = _310_v70
		var _nw29 *C6 = New_C6_()
		_ = _nw29
		_nw29.Ctor__(_222_v11, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _dafny.ArrayCastTo((_237_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_237_v21), 0))).Int())), func() _dafny.Set {
			var _coll27 = _dafny.NewBuilder()
			_ = _coll27
			for _iter32 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(163), _dafny.IntOfInt64(958))); ; {
				_compr_27, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _311_v66 _dafny.Int
				_311_v66 = interface{}(_compr_27).(_dafny.Int)
				if ((_dafny.IntOfInt64(163)).Cmp(_311_v66) <= 0) && ((_311_v66).Cmp(_dafny.IntOfInt64(958)) < 0) {
					_coll27.Add(Companion_Default___.SafeModuloInt(_311_v66, _306_i11))
				}
			}
			return _coll27.ToSet()
		}(), _309_v69, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool))
		_310_v70 = _nw29
		var _312_v71 _dafny.Map
		_ = _312_v71
		_312_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_306_i11, _310_v70)
		var _pat_let_tv10 = _215_v4
		_ = _pat_let_tv10
		var _313_v72 T1
		_ = _313_v72
		var _nw30 *C3 = New_C3_()
		_ = _nw30
		_nw30.Ctor__(!((_312_v71).Equals(_312_v71)), func(_pat_let2_0 D1) D1 {
			return func(_314_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let3_0 _dafny.Sequence) D1 {
					return func(_315_dt__update_hcf1_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC1_(_315_dt__update_hcf1_h0)
					}(_pat_let3_0)
				}(_pat_let_tv10)
			}(_pat_let2_0)
		}(Companion_D1_.Create_DC1_(_215_v4)), (_310_v70).F13(), (_310_v70).F14(), (_310_v70).F15(), _309_v69, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool))
		_313_v72 = _nw30
		_313_v72 = (func() T1 {
			if (_310_v70).F17() {
				return _313_v72
			}
			return _313_v72
		})()
		_234_v18 = _234_v18
		_213_v2 = (_220_v9).Cmp(_220_v9) > 0
		var _316_v73 _dafny.Sequence
		_ = _316_v73
		_316_v73 = _dafny.SeqOf(_305_v65)
		var _317_v74 *C8
		_ = _317_v74
		var _nw31 *C8 = New_C8_()
		_ = _nw31
		_nw31.Ctor__((_310_v70).F17())
		_317_v74 = _nw31
		var _318_v75 _dafny.Set
		_ = _318_v75
		_318_v75 = _dafny.SetOf(_317_v74)
		var _319_v76 _dafny.Array
		_ = _319_v76
		var _nwElement0_8 _dafny.Map = (_305_v65).Merge((_305_v65).Update((_313_v72).F13(), _306_i11))
		_ = _nwElement0_8
		var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(16))
		_ = _nw32
		(_nw32).ArraySet1(_nwElement0_8, 0)
		(_nw32).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _306_i11), 1)
		(_nw32).ArraySet1((_310_v70).Fm10((_310_v70).F13(), (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _dafny.SeqOf((_310_v70).F17()), _228_globalState), 2)
		(_nw32).ArraySet1(_305_v65, 3)
		(_nw32).ArraySet1(_305_v65, 4)
		(_nw32).ArraySet1(_305_v65, 5)
		(_nw32).ArraySet1(_305_v65, 6)
		(_nw32).ArraySet1((_305_v65).Merge(_305_v65), 7)
		(_nw32).ArraySet1((_305_v65).Merge((_316_v73).Select((Companion_Default___.SafeIndex((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_316_v73).Cardinality()))).Uint32()).(_dafny.Map)), 8)
		(_nw32).ArraySet1((_305_v65).Merge(_305_v65), 9)
		(_nw32).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_310_v70).F17(), (_318_v75).Cardinality())).Merge((Companion_Default___.Fm67((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _228_globalState)).Dtor_cf32()), 10)
		(_nw32).ArraySet1(_305_v65, 11)
		(_nw32).ArraySet1(_305_v65, 12)
		(_nw32).ArraySet1(_305_v65, 13)
		(_nw32).ArraySet1(_305_v65, 14)
		(_nw32).ArraySet1(_305_v65, 15)
		_319_v76 = _nw32
		_319_v76 = _319_v76
	}
	if _213_v2 {
		var _320_v77 _dafny.MultiSet
		_ = _320_v77
		var _out4 _dafny.MultiSet
		_ = _out4
		_out4 = Companion_Default___.M0((_215_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oxurkuj")).Cardinality()), _dafny.IntOfUint32((_215_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _228_globalState)
		_320_v77 = _out4
		(_228_globalState).F7 = (_220_v9).Times(_220_v9)
		var _321_v78 _dafny.MultiSet
		_ = _321_v78
		var _out5 _dafny.MultiSet
		_ = _out5
		_out5 = Companion_Default___.M0(_219_v8, _213_v2, _228_globalState)
		_321_v78 = _out5
		var _322_v79 D16
		_ = _322_v79
		_322_v79 = Companion_D16_.Create_DC41_(!(_213_v2))
		_213_v2 = (_322_v79).Dtor_cf84()
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
		_ = _index23
		(_230_v14).ArraySet1((func() bool {
			if ((_dafny.Zero).Minus(_220_v9)).Cmp(_220_v9) < 0 {
				return _213_v2
			}
			return !(!_dafny.Companion_Sequence_.Equal(_215_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-115))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg42 _dafny.Int) interface{} {
					return coer42(arg42)
				}
			}((func(_323_v8 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_324_i12 _dafny.Int) _dafny.CodePoint {
					return _323_v8
				}
			})(_219_v8)))))
		})(), (_index23).Int())
	} else {
		var _325_v80 _dafny.Map
		_ = _325_v80
		_325_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _233_v17), ((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool)) || ((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool)))
		var _326_v81 _dafny.Map
		_ = _326_v81
		_326_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), _233_v17)
		_325_v80 = (_325_v80).Update(_326_v81, _213_v2)
		_215_v4 = _dafny.Companion_Sequence_.Concatenate(_215_v4, _dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if _213_v2 {
				return _215_v4
			}
			return _215_v4
		})(), (Companion_Default___.SafeIndex((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if _213_v2 {
				return _215_v4
			}
			return _215_v4
		})()).Cardinality()))).Uint32(), _219_v8))
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
		_ = _index24
		var _rhs13 bool = false
		_ = _rhs13
		var _rhs14 bool = (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool)
		_ = _rhs14
		var _lhs14 _dafny.Array = _230_v14
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))
		_ = _lhs15
		_213_v2 = _rhs13
		(_lhs14).ArraySet1(_rhs14, (_lhs15).Int())
		var _327_v82 _dafny.Map
		_ = _327_v82
		_327_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_219_v8, (_234_v18).IsProperSubsetOf(_234_v18))
		_213_v2 = (func() bool {
			if (_327_v82).Contains(_dafny.CodePoint('x')) {
				return (_327_v82).Get(_dafny.CodePoint('x')).(bool)
			}
			return _213_v2
		})()
		var _328_v83 _dafny.Array
		_ = _328_v83
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
		_ = _nw33
		_328_v83 = _nw33
		var _329_v84 _dafny.MultiSet
		_ = _329_v84
		_329_v84 = _dafny.MultiSetOf(_215_v4, _215_v4)
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_328_v83), 0))
		_ = _index25
		(_328_v83).ArraySet1(_329_v84, (_index25).Int())
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(931), _dafny.ArrayLen((_328_v83), 0))
		_ = _index26
		(_328_v83).ArraySet1(_329_v84, (_index26).Int())
	}
	var _330_v85 _dafny.Map
	_ = _330_v85
	_330_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_220_v9, (_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool))
	var _331_v86 _dafny.Sequence
	_ = _331_v86
	_331_v86 = _dafny.SeqOf((_dafny.Zero).Minus(_220_v9), _220_v9, _220_v9, (_330_v85).Cardinality())
	var _hi4 _dafny.Int = (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)
	_ = _hi4
	for _332_i13 := (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_331_v86, _dafny.SeqOf((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)))).Cardinality())); _332_i13.Cmp(_hi4) < 0; _332_i13 = _332_i13.Plus(_dafny.One) {
		var _333_v87 _dafny.Array
		_ = _333_v87
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(4))
		_ = _nw34
		_333_v87 = _nw34
		var _334_v88 *C9
		_ = _334_v88
		var _nw35 *C9 = New_C9_()
		_ = _nw35
		_nw35.Ctor__()
		_334_v88 = _nw35
		var _335_v89 _dafny.Map
		_ = _335_v89
		_335_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _333_v87)
		var _336_v90 D1
		_ = _336_v90
		_336_v90 = Companion_D1_.Create_DC2_(_332_i13, (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int), _213_v2)
		var _337_v91 *C0
		_ = _337_v91
		var _nw36 *C0 = New_C0_()
		_ = _nw36
		_nw36.Ctor__((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool), Companion_Default___.Fm58(_228_globalState), _336_v90, _213_v2)
		_337_v91 = _nw36
		var _338_v92 _dafny.Set
		_ = _338_v92
		_338_v92 = _dafny.SetOf(_337_v91, _337_v91)
		var _339_v93 _dafny.Map
		_ = _339_v93
		_339_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_334_v88, (func() _dafny.Array {
			if (_335_v89).Contains(((Companion_D20_.Create_DC52_(_338_v92)).Dtor_cf99()).Cardinality()) {
				return (_335_v89).Get(((Companion_D20_.Create_DC52_(_338_v92)).Dtor_cf99()).Cardinality()).(_dafny.Array)
			}
			return _333_v87
		})())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _index27
		var _rhs15 _dafny.Int = (_220_v9).Plus(_220_v9)
		_ = _rhs15
		var _rhs16 bool = !((_230_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(919), _dafny.ArrayLen((_230_v14), 0))).Int()).(bool))
		_ = _rhs16
		var _rhs17 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_332_i13))
		_ = _rhs17
		var _rhs18 _dafny.Array = (func() _dafny.Array {
			if (_339_v93).Contains(_334_v88) {
				return (_339_v93).Get(_334_v88).(_dafny.Array)
			}
			return _333_v87
		})()
		_ = _rhs18
		var _lhs16 _dafny.Array = _216_v5
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _lhs17
		(_lhs16).ArraySet1(_rhs15, (_lhs17).Int())
		_213_v2 = _rhs16
		_220_v9 = _rhs17
		_333_v87 = _rhs18
		_331_v86 = _331_v86
		_220_v9 = (func() _dafny.Int {
			if Companion_Default___.Fm17(_228_globalState) {
				return _220_v9
			}
			return (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)
		})()
		var _340_v94 *C4
		_ = _340_v94
		var _nw37 *C4 = New_C4_()
		_ = _nw37
		_nw37.Ctor__()
		_340_v94 = _nw37
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _index28
		var _rhs19 _dafny.Int = (_216_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))).Int()).(_dafny.Int)
		_ = _rhs19
		var _rhs20 _dafny.Int = _332_i13
		_ = _rhs20
		var _rhs21 _dafny.Array = _216_v5
		_ = _rhs21
		var _rhs22 *C4 = (func() *C4 {
			if _337_v91.F19 {
				return _340_v94
			}
			return _340_v94
		})()
		_ = _rhs22
		var _lhs18 *GlobalState = _228_globalState
		_ = _lhs18
		var _lhs19 _dafny.Array = _216_v5
		_ = _lhs19
		var _lhs20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_216_v5), 0))
		_ = _lhs20
		_lhs18.F7 = _rhs19
		(_lhs19).ArraySet1(_rhs20, (_lhs20).Int())
		_216_v5 = _rhs21
		_340_v94 = _rhs22
	}
	_dafny.Print(((_210_v1).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_210_v1).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_213_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_214_v3, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_215_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_216_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_217_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_218_v7).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("dgxao"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_219_v8)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_220_v9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_221_v10).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(336), _dafny.UnicodeSeqOfUtf8Bytes("dgxao"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v11).Dtor_cf1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_223_v12).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_223_v12).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v12).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_228_globalState).F0()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((((_228_globalState).F0()).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_228_globalState.F2, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F3().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_228_globalState.F7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_228_globalState).F8().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F9()).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.Zero).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.One).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.Sequence), _dafny.SeqOf(_dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'), _dafny.CodePoint('g'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_228_globalState).F10()).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.Sequence).VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_229_v13).Equals(_dafny.MultiSetOf(false, false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_230_v14).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_231_v15).Dtor_cf5()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_232_v16).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_233_v17).Equals(_dafny.SetOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_234_v18).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_235_v19).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_dafny.One), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf6()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v20).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v20).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_236_v20).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_236_v20).Dtor_cf10()).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.Zero).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.One).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(2)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(3)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(4)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(5)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(6)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(7)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(8)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(9)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(10)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(11)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(12)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(13)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(14)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(15)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(16)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(17)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(18)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(19)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(20)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(21)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(22)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(23)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(24)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(25)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(26)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(27)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_dafny.ArrayCastTo((_237_v21).ArrayGet1((_dafny.IntOfInt64(28)).Int()))).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_238_v22).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(true))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_268_i5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_289_v58).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_293_v61).Dtor_cf94())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_293_v61).Dtor_cf95())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_293_v61).Dtor_cf96(), _dafny.SeqOf(_dafny.IntOfInt64(3), _dafny.IntOfInt64(336))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_294_v62).Equals(_dafny.MultiSetOf(false, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_295_v63).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_305_v65).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_330_v85).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_331_v86, _dafny.SeqOf(_dafny.IntOfInt64(-1), _dafny.One, _dafny.One, _dafny.One)))
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
	Cf0 _dafny.MultiSet
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.MultiSet) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D0) Dtor_cf0() _dafny.MultiSet {
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
			return ok && data1.Cf0.Equals(data2.Cf0)
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
	Cf2 _dafny.Int
	Cf3 _dafny.Int
	Cf4 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf2 _dafny.Int, Cf3 _dafny.Int, Cf4 bool) D1 {
	return D1{D1_DC2{Cf2, Cf3, Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

type D1_DC1 struct {
	Cf1 _dafny.Sequence
}

func (D1_DC1) isD1() {}

func (CompanionStruct_D1_) Create_DC1_(Cf1 _dafny.Sequence) D1 {
	return D1{D1_DC1{Cf1}}
}

func (_this D1) Is_DC1() bool {
	_, ok := _this.Get_().(D1_DC1)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC2_(_dafny.Zero, _dafny.Zero, false)
}

func (_this D1) Dtor_cf2() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC2).Cf3
}

func (_this D1) Dtor_cf4() bool {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) Dtor_cf1() _dafny.Sequence {
	return _this.Get_().(D1_DC1).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D1_DC1:
		{
			return "D1.DC1" + "(" + data.Cf1.VerbatimString(true) + ")"
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
			return ok && data1.Cf2.Cmp(data2.Cf2) == 0 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D1_DC1:
		{
			data2, ok := other.Get_().(D1_DC1)
			return ok && data1.Cf1.Equals(data2.Cf1)
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
	Cf6  _dafny.Array
	Cf7  bool
	Cf8  _dafny.Int
	Cf9  bool
	Cf10 _dafny.Array
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf6 _dafny.Array, Cf7 bool, Cf8 _dafny.Int, Cf9 bool, Cf10 _dafny.Array) D2 {
	return D2{D2_DC4{Cf6, Cf7, Cf8, Cf9, Cf10}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC3 struct {
	Cf5 _dafny.Array
}

func (D2_DC3) isD2() {}

func (CompanionStruct_D2_) Create_DC3_(Cf5 _dafny.Array) D2 {
	return D2{D2_DC3{Cf5}}
}

func (_this D2) Is_DC3() bool {
	_, ok := _this.Get_().(D2_DC3)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC4_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.Zero, false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D2) Dtor_cf6() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf6
}

func (_this D2) Dtor_cf7() bool {
	return _this.Get_().(D2_DC4).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC4).Cf8
}

func (_this D2) Dtor_cf9() bool {
	return _this.Get_().(D2_DC4).Cf9
}

func (_this D2) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D2_DC4).Cf10
}

func (_this D2) Dtor_cf5() _dafny.Array {
	return _this.Get_().(D2_DC3).Cf5
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ")"
		}
	case D2_DC3:
		{
			return "D2.DC3" + "(" + _dafny.String(data.Cf5) + ")"
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
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7 == data2.Cf7 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9 == data2.Cf9 && data1.Cf10 == data2.Cf10
		}
	case D2_DC3:
		{
			data2, ok := other.Get_().(D2_DC3)
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
	Cf12 _dafny.Int
	Cf13 bool
}

func (D3_DC6) isD3() {}

func (CompanionStruct_D3_) Create_DC6_(Cf12 _dafny.Int, Cf13 bool) D3 {
	return D3{D3_DC6{Cf12, Cf13}}
}

func (_this D3) Is_DC6() bool {
	_, ok := _this.Get_().(D3_DC6)
	return ok
}

type D3_DC7 struct {
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_() D3 {
	return D3{D3_DC7{}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC8 struct {
	Cf14 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 _dafny.Int) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC5 struct {
	Cf11 _dafny.CodePoint
}

func (D3_DC5) isD3() {}

func (CompanionStruct_D3_) Create_DC5_(Cf11 _dafny.CodePoint) D3 {
	return D3{D3_DC5{Cf11}}
}

func (_this D3) Is_DC5() bool {
	_, ok := _this.Get_().(D3_DC5)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC6_(_dafny.Zero, false)
}

func (_this D3) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D3_DC6).Cf12
}

func (_this D3) Dtor_cf13() bool {
	return _this.Get_().(D3_DC6).Cf13
}

func (_this D3) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf11() _dafny.CodePoint {
	return _this.Get_().(D3_DC5).Cf11
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC6:
		{
			return "D3.DC6" + "(" + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC5:
		{
			return "D3.DC5" + "(" + _dafny.String(data.Cf11) + ")"
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
			return ok && data1.Cf12.Cmp(data2.Cf12) == 0 && data1.Cf13 == data2.Cf13
		}
	case D3_DC7:
		{
			_, ok := other.Get_().(D3_DC7)
			return ok
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf14.Cmp(data2.Cf14) == 0
		}
	case D3_DC5:
		{
			data2, ok := other.Get_().(D3_DC5)
			return ok && data1.Cf11 == data2.Cf11
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

type D4_DC9 struct {
	Cf15 _dafny.Array
}

func (D4_DC9) isD4() {}

func (CompanionStruct_D4_) Create_DC9_(Cf15 _dafny.Array) D4 {
	return D4{D4_DC9{Cf15}}
}

func (_this D4) Is_DC9() bool {
	_, ok := _this.Get_().(D4_DC9)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D4) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D4_DC9).Cf15
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC9:
		{
			return "D4.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC9:
		{
			data2, ok := other.Get_().(D4_DC9)
			return ok && data1.Cf15 == data2.Cf15
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

type D5_DC10 struct {
	Cf16 _dafny.Set
}

func (D5_DC10) isD5() {}

func (CompanionStruct_D5_) Create_DC10_(Cf16 _dafny.Set) D5 {
	return D5{D5_DC10{Cf16}}
}

func (_this D5) Is_DC10() bool {
	_, ok := _this.Get_().(D5_DC10)
	return ok
}

func (CompanionStruct_D5_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D5) Dtor_cf16() _dafny.Set {
	return _this.Get_().(D5_DC10).Cf16
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC10:
		{
			return "D5.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC10:
		{
			data2, ok := other.Get_().(D5_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D6_DC12 struct {
	Cf18 bool
	Cf19 _dafny.Sequence
	Cf20 _dafny.Int
	Cf21 bool
}

func (D6_DC12) isD6() {}

func (CompanionStruct_D6_) Create_DC12_(Cf18 bool, Cf19 _dafny.Sequence, Cf20 _dafny.Int, Cf21 bool) D6 {
	return D6{D6_DC12{Cf18, Cf19, Cf20, Cf21}}
}

func (_this D6) Is_DC12() bool {
	_, ok := _this.Get_().(D6_DC12)
	return ok
}

type D6_DC13 struct {
	Cf22 bool
	Cf23 bool
	Cf24 _dafny.Int
	Cf25 _dafny.Int
	Cf26 bool
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf22 bool, Cf23 bool, Cf24 _dafny.Int, Cf25 _dafny.Int, Cf26 bool) D6 {
	return D6{D6_DC13{Cf22, Cf23, Cf24, Cf25, Cf26}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

type D6_DC11 struct {
	Cf17 _dafny.MultiSet
}

func (D6_DC11) isD6() {}

func (CompanionStruct_D6_) Create_DC11_(Cf17 _dafny.MultiSet) D6 {
	return D6{D6_DC11{Cf17}}
}

func (_this D6) Is_DC11() bool {
	_, ok := _this.Get_().(D6_DC11)
	return ok
}

type D6_DC14 struct {
	Cf27 D6
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf27 D6) D6 {
	return D6{D6_DC14{Cf27}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC12_(false, _dafny.EmptySeq, _dafny.Zero, false)
}

func (_this D6) Dtor_cf18() bool {
	return _this.Get_().(D6_DC12).Cf18
}

func (_this D6) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D6_DC12).Cf19
}

func (_this D6) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D6_DC12).Cf20
}

func (_this D6) Dtor_cf21() bool {
	return _this.Get_().(D6_DC12).Cf21
}

func (_this D6) Dtor_cf22() bool {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) Dtor_cf23() bool {
	return _this.Get_().(D6_DC13).Cf23
}

func (_this D6) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC13).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC13).Cf26
}

func (_this D6) Dtor_cf17() _dafny.MultiSet {
	return _this.Get_().(D6_DC11).Cf17
}

func (_this D6) Dtor_cf27() D6 {
	return _this.Get_().(D6_DC14).Cf27
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC12:
		{
			return "D6.DC12" + "(" + _dafny.String(data.Cf18) + ", " + data.Cf19.VerbatimString(true) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ")"
		}
	case D6_DC11:
		{
			return "D6.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf27) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC12:
		{
			data2, ok := other.Get_().(D6_DC12)
			return ok && data1.Cf18 == data2.Cf18 && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21 == data2.Cf21
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf22 == data2.Cf22 && data1.Cf23 == data2.Cf23 && data1.Cf24.Cmp(data2.Cf24) == 0 && data1.Cf25.Cmp(data2.Cf25) == 0 && data1.Cf26 == data2.Cf26
		}
	case D6_DC11:
		{
			data2, ok := other.Get_().(D6_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D6_DC14:
		{
			data2, ok := other.Get_().(D6_DC14)
			return ok && data1.Cf27.Equals(data2.Cf27)
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

type D7_DC16 struct {
	Cf29 _dafny.Int
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf29 _dafny.Int) D7 {
	return D7{D7_DC16{Cf29}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

type D7_DC15 struct {
	Cf28 _dafny.Sequence
}

func (D7_DC15) isD7() {}

func (CompanionStruct_D7_) Create_DC15_(Cf28 _dafny.Sequence) D7 {
	return D7{D7_DC15{Cf28}}
}

func (_this D7) Is_DC15() bool {
	_, ok := _this.Get_().(D7_DC15)
	return ok
}

type D7_DC17 struct {
	Cf30 D7
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf30 D7) D7 {
	return D7{D7_DC17{Cf30}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC16_(_dafny.Zero)
}

func (_this D7) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D7_DC16).Cf29
}

func (_this D7) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D7_DC15).Cf28
}

func (_this D7) Dtor_cf30() D7 {
	return _this.Get_().(D7_DC17).Cf30
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC15:
		{
			return "D7.DC15" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
			return ok && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D7_DC15:
		{
			data2, ok := other.Get_().(D7_DC15)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D8_DC19 struct {
	Cf32 _dafny.Map
	Cf33 _dafny.Set
	Cf34 _dafny.MultiSet
	Cf35 D1
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf32 _dafny.Map, Cf33 _dafny.Set, Cf34 _dafny.MultiSet, Cf35 D1) D8 {
	return D8{D8_DC19{Cf32, Cf33, Cf34, Cf35}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC18 struct {
	Cf31 _dafny.Array
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf31 _dafny.Array) D8 {
	return D8{D8_DC18{Cf31}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC19_(_dafny.EmptyMap, _dafny.EmptySet, _dafny.EmptyMultiSet, Companion_D1_.Default())
}

func (_this D8) Dtor_cf32() _dafny.Map {
	return _this.Get_().(D8_DC19).Cf32
}

func (_this D8) Dtor_cf33() _dafny.Set {
	return _this.Get_().(D8_DC19).Cf33
}

func (_this D8) Dtor_cf34() _dafny.MultiSet {
	return _this.Get_().(D8_DC19).Cf34
}

func (_this D8) Dtor_cf35() D1 {
	return _this.Get_().(D8_DC19).Cf35
}

func (_this D8) Dtor_cf31() _dafny.Array {
	return _this.Get_().(D8_DC18).Cf31
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ")"
		}
	case D8_DC18:
		{
			return "D8.DC18" + "(" + _dafny.String(data.Cf31) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf32.Equals(data2.Cf32) && data1.Cf33.Equals(data2.Cf33) && data1.Cf34.Equals(data2.Cf34) && data1.Cf35.Equals(data2.Cf35)
		}
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && data1.Cf31 == data2.Cf31
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
}

func (D9_DC21) isD9() {}

func (CompanionStruct_D9_) Create_DC21_() D9 {
	return D9{D9_DC21{}}
}

func (_this D9) Is_DC21() bool {
	_, ok := _this.Get_().(D9_DC21)
	return ok
}

type D9_DC22 struct {
	Cf37 bool
	Cf38 _dafny.Int
	Cf39 bool
	Cf40 _dafny.Set
	Cf41 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf37 bool, Cf38 _dafny.Int, Cf39 bool, Cf40 _dafny.Set, Cf41 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf37, Cf38, Cf39, Cf40, Cf41}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

type D9_DC23 struct {
	Cf42 _dafny.Int
	Cf43 _dafny.Int
	Cf44 _dafny.Int
	Cf45 _dafny.Int
	Cf46 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf42 _dafny.Int, Cf43 _dafny.Int, Cf44 _dafny.Int, Cf45 _dafny.Int, Cf46 _dafny.Int) D9 {
	return D9{D9_DC23{Cf42, Cf43, Cf44, Cf45, Cf46}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC20 struct {
	Cf36 _dafny.Set
}

func (D9_DC20) isD9() {}

func (CompanionStruct_D9_) Create_DC20_(Cf36 _dafny.Set) D9 {
	return D9{D9_DC20{Cf36}}
}

func (_this D9) Is_DC20() bool {
	_, ok := _this.Get_().(D9_DC20)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC21_()
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC22).Cf37
}

func (_this D9) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D9_DC22).Cf38
}

func (_this D9) Dtor_cf39() bool {
	return _this.Get_().(D9_DC22).Cf39
}

func (_this D9) Dtor_cf40() _dafny.Set {
	return _this.Get_().(D9_DC22).Cf40
}

func (_this D9) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf41
}

func (_this D9) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf42
}

func (_this D9) Dtor_cf43() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf43
}

func (_this D9) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf44
}

func (_this D9) Dtor_cf45() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf45
}

func (_this D9) Dtor_cf46() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf46
}

func (_this D9) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D9_DC20).Cf36
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC21:
		{
			return "D9.DC21"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + data.Cf41.VerbatimString(true) + ")"
		}
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ", " + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ")"
		}
	case D9_DC20:
		{
			return "D9.DC20" + "(" + _dafny.String(data.Cf36) + ")"
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
			_, ok := other.Get_().(D9_DC21)
			return ok
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0 && data1.Cf39 == data2.Cf39 && data1.Cf40.Equals(data2.Cf40) && data1.Cf41.Equals(data2.Cf41)
		}
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43.Cmp(data2.Cf43) == 0 && data1.Cf44.Cmp(data2.Cf44) == 0 && data1.Cf45.Cmp(data2.Cf45) == 0 && data1.Cf46.Cmp(data2.Cf46) == 0
		}
	case D9_DC20:
		{
			data2, ok := other.Get_().(D9_DC20)
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

type D10_DC24 struct {
	Cf47 _dafny.MultiSet
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf47 _dafny.MultiSet) D10 {
	return D10{D10_DC24{Cf47}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

func (CompanionStruct_D10_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D10) Dtor_cf47() _dafny.MultiSet {
	return _this.Get_().(D10_DC24).Cf47
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf47) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf47.Equals(data2.Cf47)
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

type D11_DC26 struct {
	Cf49 _dafny.Array
	Cf50 _dafny.Sequence
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf49 _dafny.Array, Cf50 _dafny.Sequence) D11 {
	return D11{D11_DC26{Cf49, Cf50}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC27 struct {
	Cf51 bool
	Cf52 _dafny.CodePoint
	Cf53 _dafny.Int
	Cf54 _dafny.Int
	Cf55 _dafny.Sequence
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf51 bool, Cf52 _dafny.CodePoint, Cf53 _dafny.Int, Cf54 _dafny.Int, Cf55 _dafny.Sequence) D11 {
	return D11{D11_DC27{Cf51, Cf52, Cf53, Cf54, Cf55}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC25 struct {
	Cf48 _dafny.Map
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf48 _dafny.Map) D11 {
	return D11{D11_DC25{Cf48}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC28 struct {
	Cf56 D11
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf56 D11) D11 {
	return D11{D11_DC28{Cf56}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC26_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq)
}

func (_this D11) Dtor_cf49() _dafny.Array {
	return _this.Get_().(D11_DC26).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Sequence {
	return _this.Get_().(D11_DC26).Cf50
}

func (_this D11) Dtor_cf51() bool {
	return _this.Get_().(D11_DC27).Cf51
}

func (_this D11) Dtor_cf52() _dafny.CodePoint {
	return _this.Get_().(D11_DC27).Cf52
}

func (_this D11) Dtor_cf53() _dafny.Int {
	return _this.Get_().(D11_DC27).Cf53
}

func (_this D11) Dtor_cf54() _dafny.Int {
	return _this.Get_().(D11_DC27).Cf54
}

func (_this D11) Dtor_cf55() _dafny.Sequence {
	return _this.Get_().(D11_DC27).Cf55
}

func (_this D11) Dtor_cf48() _dafny.Map {
	return _this.Get_().(D11_DC25).Cf48
}

func (_this D11) Dtor_cf56() D11 {
	return _this.Get_().(D11_DC28).Cf56
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf49) + ", " + data.Cf50.VerbatimString(true) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ", " + _dafny.String(data.Cf55) + ")"
		}
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf48) + ")"
		}
	case D11_DC28:
		{
			return "D11.DC28" + "(" + _dafny.String(data.Cf56) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && data1.Cf49 == data2.Cf49 && data1.Cf50.Equals(data2.Cf50)
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf51 == data2.Cf51 && data1.Cf52 == data2.Cf52 && data1.Cf53.Cmp(data2.Cf53) == 0 && data1.Cf54.Cmp(data2.Cf54) == 0 && data1.Cf55.Equals(data2.Cf55)
		}
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf48.Equals(data2.Cf48)
		}
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf56.Equals(data2.Cf56)
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
	Cf58 _dafny.Array
	Cf59 bool
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf58 _dafny.Array, Cf59 bool) D12 {
	return D12{D12_DC30{Cf58, Cf59}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

type D12_DC31 struct {
	Cf60 bool
	Cf61 _dafny.Array
	Cf62 bool
	Cf63 _dafny.CodePoint
	Cf64 _dafny.MultiSet
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf60 bool, Cf61 _dafny.Array, Cf62 bool, Cf63 _dafny.CodePoint, Cf64 _dafny.MultiSet) D12 {
	return D12{D12_DC31{Cf60, Cf61, Cf62, Cf63, Cf64}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC32 struct {
	Cf65 bool
	Cf66 _dafny.Int
	Cf67 bool
	Cf68 bool
	Cf69 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf65 bool, Cf66 _dafny.Int, Cf67 bool, Cf68 bool, Cf69 _dafny.Int) D12 {
	return D12{D12_DC32{Cf65, Cf66, Cf67, Cf68, Cf69}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC29 struct {
	Cf57 _dafny.Map
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf57 _dafny.Map) D12 {
	return D12{D12_DC29{Cf57}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC30_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D12) Dtor_cf58() _dafny.Array {
	return _this.Get_().(D12_DC30).Cf58
}

func (_this D12) Dtor_cf59() bool {
	return _this.Get_().(D12_DC30).Cf59
}

func (_this D12) Dtor_cf60() bool {
	return _this.Get_().(D12_DC31).Cf60
}

func (_this D12) Dtor_cf61() _dafny.Array {
	return _this.Get_().(D12_DC31).Cf61
}

func (_this D12) Dtor_cf62() bool {
	return _this.Get_().(D12_DC31).Cf62
}

func (_this D12) Dtor_cf63() _dafny.CodePoint {
	return _this.Get_().(D12_DC31).Cf63
}

func (_this D12) Dtor_cf64() _dafny.MultiSet {
	return _this.Get_().(D12_DC31).Cf64
}

func (_this D12) Dtor_cf65() bool {
	return _this.Get_().(D12_DC32).Cf65
}

func (_this D12) Dtor_cf66() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf66
}

func (_this D12) Dtor_cf67() bool {
	return _this.Get_().(D12_DC32).Cf67
}

func (_this D12) Dtor_cf68() bool {
	return _this.Get_().(D12_DC32).Cf68
}

func (_this D12) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf69
}

func (_this D12) Dtor_cf57() _dafny.Map {
	return _this.Get_().(D12_DC29).Cf57
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf65) + ", " + _dafny.String(data.Cf66) + ", " + _dafny.String(data.Cf67) + ", " + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ")"
		}
	case D12_DC29:
		{
			return "D12.DC29" + "(" + _dafny.String(data.Cf57) + ")"
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
			return ok && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf60 == data2.Cf60 && data1.Cf61 == data2.Cf61 && data1.Cf62 == data2.Cf62 && data1.Cf63 == data2.Cf63 && data1.Cf64.Equals(data2.Cf64)
		}
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf65 == data2.Cf65 && data1.Cf66.Cmp(data2.Cf66) == 0 && data1.Cf67 == data2.Cf67 && data1.Cf68 == data2.Cf68 && data1.Cf69.Cmp(data2.Cf69) == 0
		}
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf57.Equals(data2.Cf57)
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
	Cf71 _dafny.Int
	Cf72 _dafny.Sequence
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf71 _dafny.Int, Cf72 _dafny.Sequence) D13 {
	return D13{D13_DC34{Cf71, Cf72}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

type D13_DC33 struct {
	Cf70 _dafny.Array
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf70 _dafny.Array) D13 {
	return D13{D13_DC33{Cf70}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

type D13_DC35 struct {
	Cf73 D13
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf73 D13) D13 {
	return D13{D13_DC35{Cf73}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC34_(_dafny.Zero, _dafny.EmptySeq)
}

func (_this D13) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D13_DC34).Cf71
}

func (_this D13) Dtor_cf72() _dafny.Sequence {
	return _this.Get_().(D13_DC34).Cf72
}

func (_this D13) Dtor_cf70() _dafny.Array {
	return _this.Get_().(D13_DC33).Cf70
}

func (_this D13) Dtor_cf73() D13 {
	return _this.Get_().(D13_DC35).Cf73
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf71) + ", " + _dafny.String(data.Cf72) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf70) + ")"
		}
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf73) + ")"
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
			return ok && data1.Cf71.Cmp(data2.Cf71) == 0 && data1.Cf72.Equals(data2.Cf72)
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf70 == data2.Cf70
		}
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf73.Equals(data2.Cf73)
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
	Cf75 _dafny.Sequence
	Cf76 _dafny.Int
}

func (D14_DC37) isD14() {}

func (CompanionStruct_D14_) Create_DC37_(Cf75 _dafny.Sequence, Cf76 _dafny.Int) D14 {
	return D14{D14_DC37{Cf75, Cf76}}
}

func (_this D14) Is_DC37() bool {
	_, ok := _this.Get_().(D14_DC37)
	return ok
}

type D14_DC36 struct {
	Cf74 _dafny.Map
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf74 _dafny.Map) D14 {
	return D14{D14_DC36{Cf74}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC37_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D14) Dtor_cf75() _dafny.Sequence {
	return _this.Get_().(D14_DC37).Cf75
}

func (_this D14) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D14_DC37).Cf76
}

func (_this D14) Dtor_cf74() _dafny.Map {
	return _this.Get_().(D14_DC36).Cf74
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC37:
		{
			return "D14.DC37" + "(" + data.Cf75.VerbatimString(true) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf74) + ")"
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
			return ok && data1.Cf75.Equals(data2.Cf75) && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf74.Equals(data2.Cf74)
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
	Cf78 _dafny.Int
	Cf79 _dafny.CodePoint
	Cf80 _dafny.Sequence
	Cf81 _dafny.Int
	Cf82 _dafny.Int
}

func (D15_DC39) isD15() {}

func (CompanionStruct_D15_) Create_DC39_(Cf78 _dafny.Int, Cf79 _dafny.CodePoint, Cf80 _dafny.Sequence, Cf81 _dafny.Int, Cf82 _dafny.Int) D15 {
	return D15{D15_DC39{Cf78, Cf79, Cf80, Cf81, Cf82}}
}

func (_this D15) Is_DC39() bool {
	_, ok := _this.Get_().(D15_DC39)
	return ok
}

type D15_DC38 struct {
	Cf77 _dafny.Sequence
}

func (D15_DC38) isD15() {}

func (CompanionStruct_D15_) Create_DC38_(Cf77 _dafny.Sequence) D15 {
	return D15{D15_DC38{Cf77}}
}

func (_this D15) Is_DC38() bool {
	_, ok := _this.Get_().(D15_DC38)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC39_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D15) Dtor_cf78() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf78
}

func (_this D15) Dtor_cf79() _dafny.CodePoint {
	return _this.Get_().(D15_DC39).Cf79
}

func (_this D15) Dtor_cf80() _dafny.Sequence {
	return _this.Get_().(D15_DC39).Cf80
}

func (_this D15) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf81
}

func (_this D15) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D15_DC39).Cf82
}

func (_this D15) Dtor_cf77() _dafny.Sequence {
	return _this.Get_().(D15_DC38).Cf77
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC39:
		{
			return "D15.DC39" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ", " + data.Cf80.VerbatimString(true) + ", " + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ")"
		}
	case D15_DC38:
		{
			return "D15.DC38" + "(" + _dafny.String(data.Cf77) + ")"
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
			return ok && data1.Cf78.Cmp(data2.Cf78) == 0 && data1.Cf79 == data2.Cf79 && data1.Cf80.Equals(data2.Cf80) && data1.Cf81.Cmp(data2.Cf81) == 0 && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D15_DC38:
		{
			data2, ok := other.Get_().(D15_DC38)
			return ok && data1.Cf77.Equals(data2.Cf77)
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
	Cf84 bool
}

func (D16_DC41) isD16() {}

func (CompanionStruct_D16_) Create_DC41_(Cf84 bool) D16 {
	return D16{D16_DC41{Cf84}}
}

func (_this D16) Is_DC41() bool {
	_, ok := _this.Get_().(D16_DC41)
	return ok
}

type D16_DC42 struct {
	Cf85 _dafny.MultiSet
}

func (D16_DC42) isD16() {}

func (CompanionStruct_D16_) Create_DC42_(Cf85 _dafny.MultiSet) D16 {
	return D16{D16_DC42{Cf85}}
}

func (_this D16) Is_DC42() bool {
	_, ok := _this.Get_().(D16_DC42)
	return ok
}

type D16_DC40 struct {
	Cf83 T0
}

func (D16_DC40) isD16() {}

func (CompanionStruct_D16_) Create_DC40_(Cf83 T0) D16 {
	return D16{D16_DC40{Cf83}}
}

func (_this D16) Is_DC40() bool {
	_, ok := _this.Get_().(D16_DC40)
	return ok
}

type D16_DC43 struct {
	Cf86 D16
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf86 D16) D16 {
	return D16{D16_DC43{Cf86}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC41_(false)
}

func (_this D16) Dtor_cf84() bool {
	return _this.Get_().(D16_DC41).Cf84
}

func (_this D16) Dtor_cf85() _dafny.MultiSet {
	return _this.Get_().(D16_DC42).Cf85
}

func (_this D16) Dtor_cf83() T0 {
	return _this.Get_().(D16_DC40).Cf83
}

func (_this D16) Dtor_cf86() D16 {
	return _this.Get_().(D16_DC43).Cf86
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC41:
		{
			return "D16.DC41" + "(" + _dafny.String(data.Cf84) + ")"
		}
	case D16_DC42:
		{
			return "D16.DC42" + "(" + _dafny.String(data.Cf85) + ")"
		}
	case D16_DC40:
		{
			return "D16.DC40" + "(" + _dafny.String(data.Cf83) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf86) + ")"
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
			return ok && data1.Cf84 == data2.Cf84
		}
	case D16_DC42:
		{
			data2, ok := other.Get_().(D16_DC42)
			return ok && data1.Cf85.Equals(data2.Cf85)
		}
	case D16_DC40:
		{
			data2, ok := other.Get_().(D16_DC40)
			return ok && _dafny.AreEqual(data1.Cf83, data2.Cf83)
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf86.Equals(data2.Cf86)
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
	Cf88 bool
}

func (D17_DC45) isD17() {}

func (CompanionStruct_D17_) Create_DC45_(Cf88 bool) D17 {
	return D17{D17_DC45{Cf88}}
}

func (_this D17) Is_DC45() bool {
	_, ok := _this.Get_().(D17_DC45)
	return ok
}

type D17_DC44 struct {
	Cf87 _dafny.Array
}

func (D17_DC44) isD17() {}

func (CompanionStruct_D17_) Create_DC44_(Cf87 _dafny.Array) D17 {
	return D17{D17_DC44{Cf87}}
}

func (_this D17) Is_DC44() bool {
	_, ok := _this.Get_().(D17_DC44)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC45_(false)
}

func (_this D17) Dtor_cf88() bool {
	return _this.Get_().(D17_DC45).Cf88
}

func (_this D17) Dtor_cf87() _dafny.Array {
	return _this.Get_().(D17_DC44).Cf87
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC45:
		{
			return "D17.DC45" + "(" + _dafny.String(data.Cf88) + ")"
		}
	case D17_DC44:
		{
			return "D17.DC44" + "(" + _dafny.String(data.Cf87) + ")"
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
			return ok && data1.Cf88 == data2.Cf88
		}
	case D17_DC44:
		{
			data2, ok := other.Get_().(D17_DC44)
			return ok && data1.Cf87 == data2.Cf87
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

type D18_DC47 struct {
	Cf90 _dafny.Int
	Cf91 _dafny.Array
}

func (D18_DC47) isD18() {}

func (CompanionStruct_D18_) Create_DC47_(Cf90 _dafny.Int, Cf91 _dafny.Array) D18 {
	return D18{D18_DC47{Cf90, Cf91}}
}

func (_this D18) Is_DC47() bool {
	_, ok := _this.Get_().(D18_DC47)
	return ok
}

type D18_DC46 struct {
	Cf89 _dafny.MultiSet
}

func (D18_DC46) isD18() {}

func (CompanionStruct_D18_) Create_DC46_(Cf89 _dafny.MultiSet) D18 {
	return D18{D18_DC46{Cf89}}
}

func (_this D18) Is_DC46() bool {
	_, ok := _this.Get_().(D18_DC46)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC47_(_dafny.Zero, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)))
}

func (_this D18) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D18_DC47).Cf90
}

func (_this D18) Dtor_cf91() _dafny.Array {
	return _this.Get_().(D18_DC47).Cf91
}

func (_this D18) Dtor_cf89() _dafny.MultiSet {
	return _this.Get_().(D18_DC46).Cf89
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC47:
		{
			return "D18.DC47" + "(" + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D18_DC46:
		{
			return "D18.DC46" + "(" + _dafny.String(data.Cf89) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC47:
		{
			data2, ok := other.Get_().(D18_DC47)
			return ok && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91 == data2.Cf91
		}
	case D18_DC46:
		{
			data2, ok := other.Get_().(D18_DC46)
			return ok && data1.Cf89.Equals(data2.Cf89)
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

type D19_DC49 struct {
	Cf93 _dafny.Set
}

func (D19_DC49) isD19() {}

func (CompanionStruct_D19_) Create_DC49_(Cf93 _dafny.Set) D19 {
	return D19{D19_DC49{Cf93}}
}

func (_this D19) Is_DC49() bool {
	_, ok := _this.Get_().(D19_DC49)
	return ok
}

type D19_DC50 struct {
	Cf94 bool
	Cf95 _dafny.Int
	Cf96 _dafny.Sequence
}

func (D19_DC50) isD19() {}

func (CompanionStruct_D19_) Create_DC50_(Cf94 bool, Cf95 _dafny.Int, Cf96 _dafny.Sequence) D19 {
	return D19{D19_DC50{Cf94, Cf95, Cf96}}
}

func (_this D19) Is_DC50() bool {
	_, ok := _this.Get_().(D19_DC50)
	return ok
}

type D19_DC51 struct {
	Cf97 _dafny.Int
	Cf98 _dafny.Array
}

func (D19_DC51) isD19() {}

func (CompanionStruct_D19_) Create_DC51_(Cf97 _dafny.Int, Cf98 _dafny.Array) D19 {
	return D19{D19_DC51{Cf97, Cf98}}
}

func (_this D19) Is_DC51() bool {
	_, ok := _this.Get_().(D19_DC51)
	return ok
}

type D19_DC48 struct {
	Cf92 *C0
}

func (D19_DC48) isD19() {}

func (CompanionStruct_D19_) Create_DC48_(Cf92 *C0) D19 {
	return D19{D19_DC48{Cf92}}
}

func (_this D19) Is_DC48() bool {
	_, ok := _this.Get_().(D19_DC48)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC49_(_dafny.EmptySet)
}

func (_this D19) Dtor_cf93() _dafny.Set {
	return _this.Get_().(D19_DC49).Cf93
}

func (_this D19) Dtor_cf94() bool {
	return _this.Get_().(D19_DC50).Cf94
}

func (_this D19) Dtor_cf95() _dafny.Int {
	return _this.Get_().(D19_DC50).Cf95
}

func (_this D19) Dtor_cf96() _dafny.Sequence {
	return _this.Get_().(D19_DC50).Cf96
}

func (_this D19) Dtor_cf97() _dafny.Int {
	return _this.Get_().(D19_DC51).Cf97
}

func (_this D19) Dtor_cf98() _dafny.Array {
	return _this.Get_().(D19_DC51).Cf98
}

func (_this D19) Dtor_cf92() *C0 {
	return _this.Get_().(D19_DC48).Cf92
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC49:
		{
			return "D19.DC49" + "(" + _dafny.String(data.Cf93) + ")"
		}
	case D19_DC50:
		{
			return "D19.DC50" + "(" + _dafny.String(data.Cf94) + ", " + _dafny.String(data.Cf95) + ", " + _dafny.String(data.Cf96) + ")"
		}
	case D19_DC51:
		{
			return "D19.DC51" + "(" + _dafny.String(data.Cf97) + ", " + _dafny.String(data.Cf98) + ")"
		}
	case D19_DC48:
		{
			return "D19.DC48" + "(" + _dafny.String(data.Cf92) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC49:
		{
			data2, ok := other.Get_().(D19_DC49)
			return ok && data1.Cf93.Equals(data2.Cf93)
		}
	case D19_DC50:
		{
			data2, ok := other.Get_().(D19_DC50)
			return ok && data1.Cf94 == data2.Cf94 && data1.Cf95.Cmp(data2.Cf95) == 0 && data1.Cf96.Equals(data2.Cf96)
		}
	case D19_DC51:
		{
			data2, ok := other.Get_().(D19_DC51)
			return ok && data1.Cf97.Cmp(data2.Cf97) == 0 && data1.Cf98 == data2.Cf98
		}
	case D19_DC48:
		{
			data2, ok := other.Get_().(D19_DC48)
			return ok && data1.Cf92 == data2.Cf92
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

type D20_DC53 struct {
	Cf100 *C2
	Cf101 bool
	Cf102 _dafny.Int
	Cf103 bool
	Cf104 bool
}

func (D20_DC53) isD20() {}

func (CompanionStruct_D20_) Create_DC53_(Cf100 *C2, Cf101 bool, Cf102 _dafny.Int, Cf103 bool, Cf104 bool) D20 {
	return D20{D20_DC53{Cf100, Cf101, Cf102, Cf103, Cf104}}
}

func (_this D20) Is_DC53() bool {
	_, ok := _this.Get_().(D20_DC53)
	return ok
}

type D20_DC52 struct {
	Cf99 _dafny.Set
}

func (D20_DC52) isD20() {}

func (CompanionStruct_D20_) Create_DC52_(Cf99 _dafny.Set) D20 {
	return D20{D20_DC52{Cf99}}
}

func (_this D20) Is_DC52() bool {
	_, ok := _this.Get_().(D20_DC52)
	return ok
}

func (CompanionStruct_D20_) Default() D20 {
	return Companion_D20_.Create_DC53_((*C2)(nil), false, _dafny.Zero, false, false)
}

func (_this D20) Dtor_cf100() *C2 {
	return _this.Get_().(D20_DC53).Cf100
}

func (_this D20) Dtor_cf101() bool {
	return _this.Get_().(D20_DC53).Cf101
}

func (_this D20) Dtor_cf102() _dafny.Int {
	return _this.Get_().(D20_DC53).Cf102
}

func (_this D20) Dtor_cf103() bool {
	return _this.Get_().(D20_DC53).Cf103
}

func (_this D20) Dtor_cf104() bool {
	return _this.Get_().(D20_DC53).Cf104
}

func (_this D20) Dtor_cf99() _dafny.Set {
	return _this.Get_().(D20_DC52).Cf99
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC53:
		{
			return "D20.DC53" + "(" + _dafny.String(data.Cf100) + ", " + _dafny.String(data.Cf101) + ", " + _dafny.String(data.Cf102) + ", " + _dafny.String(data.Cf103) + ", " + _dafny.String(data.Cf104) + ")"
		}
	case D20_DC52:
		{
			return "D20.DC52" + "(" + _dafny.String(data.Cf99) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC53:
		{
			data2, ok := other.Get_().(D20_DC53)
			return ok && data1.Cf100 == data2.Cf100 && data1.Cf101 == data2.Cf101 && data1.Cf102.Cmp(data2.Cf102) == 0 && data1.Cf103 == data2.Cf103 && data1.Cf104 == data2.Cf104
		}
	case D20_DC52:
		{
			data2, ok := other.Get_().(D20_DC52)
			return ok && data1.Cf99.Equals(data2.Cf99)
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

type D21_DC55 struct {
	Cf106 bool
	Cf107 _dafny.Int
	Cf108 _dafny.Int
}

func (D21_DC55) isD21() {}

func (CompanionStruct_D21_) Create_DC55_(Cf106 bool, Cf107 _dafny.Int, Cf108 _dafny.Int) D21 {
	return D21{D21_DC55{Cf106, Cf107, Cf108}}
}

func (_this D21) Is_DC55() bool {
	_, ok := _this.Get_().(D21_DC55)
	return ok
}

type D21_DC56 struct {
	Cf109 bool
}

func (D21_DC56) isD21() {}

func (CompanionStruct_D21_) Create_DC56_(Cf109 bool) D21 {
	return D21{D21_DC56{Cf109}}
}

func (_this D21) Is_DC56() bool {
	_, ok := _this.Get_().(D21_DC56)
	return ok
}

type D21_DC57 struct {
	Cf110 bool
	Cf111 _dafny.CodePoint
}

func (D21_DC57) isD21() {}

func (CompanionStruct_D21_) Create_DC57_(Cf110 bool, Cf111 _dafny.CodePoint) D21 {
	return D21{D21_DC57{Cf110, Cf111}}
}

func (_this D21) Is_DC57() bool {
	_, ok := _this.Get_().(D21_DC57)
	return ok
}

type D21_DC54 struct {
	Cf105 _dafny.Map
}

func (D21_DC54) isD21() {}

func (CompanionStruct_D21_) Create_DC54_(Cf105 _dafny.Map) D21 {
	return D21{D21_DC54{Cf105}}
}

func (_this D21) Is_DC54() bool {
	_, ok := _this.Get_().(D21_DC54)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC55_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D21) Dtor_cf106() bool {
	return _this.Get_().(D21_DC55).Cf106
}

func (_this D21) Dtor_cf107() _dafny.Int {
	return _this.Get_().(D21_DC55).Cf107
}

func (_this D21) Dtor_cf108() _dafny.Int {
	return _this.Get_().(D21_DC55).Cf108
}

func (_this D21) Dtor_cf109() bool {
	return _this.Get_().(D21_DC56).Cf109
}

func (_this D21) Dtor_cf110() bool {
	return _this.Get_().(D21_DC57).Cf110
}

func (_this D21) Dtor_cf111() _dafny.CodePoint {
	return _this.Get_().(D21_DC57).Cf111
}

func (_this D21) Dtor_cf105() _dafny.Map {
	return _this.Get_().(D21_DC54).Cf105
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC55:
		{
			return "D21.DC55" + "(" + _dafny.String(data.Cf106) + ", " + _dafny.String(data.Cf107) + ", " + _dafny.String(data.Cf108) + ")"
		}
	case D21_DC56:
		{
			return "D21.DC56" + "(" + _dafny.String(data.Cf109) + ")"
		}
	case D21_DC57:
		{
			return "D21.DC57" + "(" + _dafny.String(data.Cf110) + ", " + _dafny.String(data.Cf111) + ")"
		}
	case D21_DC54:
		{
			return "D21.DC54" + "(" + _dafny.String(data.Cf105) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC55:
		{
			data2, ok := other.Get_().(D21_DC55)
			return ok && data1.Cf106 == data2.Cf106 && data1.Cf107.Cmp(data2.Cf107) == 0 && data1.Cf108.Cmp(data2.Cf108) == 0
		}
	case D21_DC56:
		{
			data2, ok := other.Get_().(D21_DC56)
			return ok && data1.Cf109 == data2.Cf109
		}
	case D21_DC57:
		{
			data2, ok := other.Get_().(D21_DC57)
			return ok && data1.Cf110 == data2.Cf110 && data1.Cf111 == data2.Cf111
		}
	case D21_DC54:
		{
			data2, ok := other.Get_().(D21_DC54)
			return ok && data1.Cf105.Equals(data2.Cf105)
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

type D22_DC59 struct {
	Cf113 _dafny.Int
	Cf114 bool
}

func (D22_DC59) isD22() {}

func (CompanionStruct_D22_) Create_DC59_(Cf113 _dafny.Int, Cf114 bool) D22 {
	return D22{D22_DC59{Cf113, Cf114}}
}

func (_this D22) Is_DC59() bool {
	_, ok := _this.Get_().(D22_DC59)
	return ok
}

type D22_DC60 struct {
	Cf115 _dafny.Int
	Cf116 _dafny.Sequence
	Cf117 bool
}

func (D22_DC60) isD22() {}

func (CompanionStruct_D22_) Create_DC60_(Cf115 _dafny.Int, Cf116 _dafny.Sequence, Cf117 bool) D22 {
	return D22{D22_DC60{Cf115, Cf116, Cf117}}
}

func (_this D22) Is_DC60() bool {
	_, ok := _this.Get_().(D22_DC60)
	return ok
}

type D22_DC58 struct {
	Cf112 _dafny.Map
}

func (D22_DC58) isD22() {}

func (CompanionStruct_D22_) Create_DC58_(Cf112 _dafny.Map) D22 {
	return D22{D22_DC58{Cf112}}
}

func (_this D22) Is_DC58() bool {
	_, ok := _this.Get_().(D22_DC58)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC59_(_dafny.Zero, false)
}

func (_this D22) Dtor_cf113() _dafny.Int {
	return _this.Get_().(D22_DC59).Cf113
}

func (_this D22) Dtor_cf114() bool {
	return _this.Get_().(D22_DC59).Cf114
}

func (_this D22) Dtor_cf115() _dafny.Int {
	return _this.Get_().(D22_DC60).Cf115
}

func (_this D22) Dtor_cf116() _dafny.Sequence {
	return _this.Get_().(D22_DC60).Cf116
}

func (_this D22) Dtor_cf117() bool {
	return _this.Get_().(D22_DC60).Cf117
}

func (_this D22) Dtor_cf112() _dafny.Map {
	return _this.Get_().(D22_DC58).Cf112
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC59:
		{
			return "D22.DC59" + "(" + _dafny.String(data.Cf113) + ", " + _dafny.String(data.Cf114) + ")"
		}
	case D22_DC60:
		{
			return "D22.DC60" + "(" + _dafny.String(data.Cf115) + ", " + _dafny.String(data.Cf116) + ", " + _dafny.String(data.Cf117) + ")"
		}
	case D22_DC58:
		{
			return "D22.DC58" + "(" + _dafny.String(data.Cf112) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC59:
		{
			data2, ok := other.Get_().(D22_DC59)
			return ok && data1.Cf113.Cmp(data2.Cf113) == 0 && data1.Cf114 == data2.Cf114
		}
	case D22_DC60:
		{
			data2, ok := other.Get_().(D22_DC60)
			return ok && data1.Cf115.Cmp(data2.Cf115) == 0 && data1.Cf116.Equals(data2.Cf116) && data1.Cf117 == data2.Cf117
		}
	case D22_DC58:
		{
			data2, ok := other.Get_().(D22_DC58)
			return ok && data1.Cf112.Equals(data2.Cf112)
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
	F12() D1
	F12_set_(value D1)
	F13() bool
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
	F12() D1
	F12_set_(value D1)
	F13() bool
	Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int
	Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
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
	F12() D1
	F12_set_(value D1)
	Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int
	Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	F13() bool
	Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map
	M5(p0 _dafny.Int, globalState *GlobalState) bool
	F14() _dafny.Array
	F15() _dafny.Set
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
	F12() D1
	F12_set_(value D1)
	Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map
	Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int
	Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int
	M5(p0 _dafny.Int, globalState *GlobalState) bool
	F14() _dafny.Array
	F15() _dafny.Set
	F13() bool
	F16() D1
	F16_set_(value D1)
	M6(p0 bool, globalState *GlobalState) _dafny.Int
	M7(globalState *GlobalState)
	F17() bool
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

// Definition of class GlobalState
type GlobalState struct {
	F2   _dafny.Sequence
	F7   _dafny.Int
	_f0  _dafny.Array
	_f1  bool
	_f3  _dafny.Sequence
	_f4  bool
	_f5  bool
	_f6  _dafny.Int
	_f8  _dafny.Sequence
	_f9  _dafny.MultiSet
	_f10 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F2 = _dafny.EmptySeq
	_this.F7 = _dafny.Zero
	_this._f0 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = false
	_this._f3 = _dafny.EmptySeq
	_this._f4 = false
	_this._f5 = false
	_this._f6 = _dafny.Zero
	_this._f8 = _dafny.EmptySeq
	_this._f9 = _dafny.EmptyMultiSet
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 _dafny.Array, f1 bool, f2 _dafny.Sequence, f3 _dafny.Sequence, f4 bool, f5 bool, f6 _dafny.Int, f7 _dafny.Int, f8 _dafny.Sequence, f9 _dafny.MultiSet, f10 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f6 = f6
		(_this).F7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
	}
}
func (_this *GlobalState) F0() _dafny.Array {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() bool {
	{
		return _this._f1
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
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}
func (_this *GlobalState) F6() _dafny.Int {
	{
		return _this._f6
	}
}
func (_this *GlobalState) F8() _dafny.Sequence {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() _dafny.MultiSet {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Array {
	{
		return _this._f10
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f12 D1
	_f13 bool
	F19  bool
	_f20 _dafny.CodePoint
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f12 = Companion_D1_.Default()
	_this._f13 = false
	_this.F19 = false
	_this._f20 = _dafny.CodePoint('D')
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

func (_this *C0) F12() D1 {
	return _this._f12
}
func (_this *C0) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C0) F13() bool {
	return _this._f13
}
func (_this *C0) Ctor__(f19 bool, f20 _dafny.CodePoint, f12 D1, f13 bool) {
	{
		(_this).F19 = f19
		(_this)._f20 = f20
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C0) Fm16(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((func() _dafny.Set {
			var _coll28 = _dafny.NewBuilder()
			_ = _coll28
			for _iter33 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bxjtdswbn"), _dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(560))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_341_i0 _dafny.Int) _dafny.CodePoint {
				return (_this).F20()
			})))).Elements()); ; {
				_compr_28, _ok33 := _iter33()
				if !_ok33 {
					break
				}
				var _342_v0 _dafny.Sequence
				_342_v0 = interface{}(_compr_28).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bxjtdswbn"), _dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(560))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}(func(_341_i0 _dafny.Int) _dafny.CodePoint {
					return (_this).F20()
				}))), _342_v0) {
					_coll28.Add(_342_v0)
				}
			}
			return _coll28.ToSet()
		}()).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("vlirfxd")))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ivnwymujb")).Cardinality())))
	}
}
func (_this *C0) F20() _dafny.CodePoint {
	{
		return _this._f20
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f12 D1
	_f13 bool
	F23  _dafny.CodePoint
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f12 = Companion_D1_.Default()
	_this._f13 = false
	_this.F23 = _dafny.CodePoint('D')
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F12() D1 {
	return _this._f12
}
func (_this *C1) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C1) F13() bool {
	return _this._f13
}
func (_this *C1) Ctor__(f23 _dafny.CodePoint, f12 D1, f13 bool) {
	{
		(_this).F23 = f23
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C1) Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt((_dafny.SetOf(_dafny.CodePoint('f'), _this.F23, _this.F23, _this.F23)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F13()), (_this).F13())).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cug")).Cardinality())))
	}
}
func (_this *C1) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F13(), (_this).F13(), (_this).F13())).Cardinality())), _dafny.IntOfInt64(454))).Cardinality()).Plus(_dafny.IntOfInt64(-762))
	}
}
func (_this *C1) Fm42(globalState *GlobalState) _dafny.MultiSet {
	{
		return _dafny.MultiSetOf(_dafny.IntOfInt64(97), _dafny.IntOfInt64(693), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((func() _dafny.Map {
			var _coll29 = _dafny.NewMapBuilder()
			_ = _coll29
			for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-83), _dafny.IntOfInt64(-824))); ; {
				_compr_29, _ok34 := _iter34()
				if !_ok34 {
					break
				}
				var _343_v0 _dafny.Int
				_343_v0 = interface{}(_compr_29).(_dafny.Int)
				if ((_dafny.IntOfInt64(-83)).Cmp(_343_v0) <= 0) && ((_343_v0).Cmp(_dafny.IntOfInt64(-824)) < 0) {
					_coll29.Add((_343_v0).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lx")).Cardinality()))), (_this).F13())
				}
			}
			return _coll29.ToMap()
		}()).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F13())).Cardinality()), (_dafny.MultiSetOf((_this).F13())).Cardinality())).Cardinality()))), Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter35 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(436), _dafny.IntOfInt64(724))).Elements()); ; {
				_compr_30, _ok35 := _iter35()
				if !_ok35 {
					break
				}
				var _344_v1 _dafny.Int
				_344_v1 = interface{}(_compr_30).(_dafny.Int)
				if (_dafny.SetOf(_dafny.IntOfInt64(436), _dafny.IntOfInt64(724))).Contains(_344_v1) {
					_coll30.Add((_344_v1).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(931), _dafny.IntOfInt64(668))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC15_(_dafny.SeqOf(_dafny.IntOfInt64(477))), (_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F13())).Cardinality()))).Cardinality())).Cardinality())
				}
			}
			return _coll30.ToMap()
		}())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iik")).Cardinality())), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-784), _dafny.IntOfInt64(-992)))
	}
}
func (_this *C1) Fm43(globalState *GlobalState) bool {
	{
		var _source8 D3 = Companion_D3_.Create_DC6_(_dafny.IntOfInt64(209), (_this).F13())
		_ = _source8
		if _source8.Is_DC6() {
			var _345___mcc_h0 _dafny.Int = _source8.Get_().(D3_DC6).Cf12
			_ = _345___mcc_h0
			var _346___mcc_h1 bool = _source8.Get_().(D3_DC6).Cf13
			_ = _346___mcc_h1
			var _347_cf13 bool = _346___mcc_h1
			_ = _347_cf13
			var _348_cf12 _dafny.Int = _345___mcc_h0
			_ = _348_cf12
			return _347_cf13
		} else if _source8.Is_DC7() {
			return (_this).F13()
		} else if _source8.Is_DC8() {
			var _349___mcc_h2 _dafny.Int = _source8.Get_().(D3_DC8).Cf14
			_ = _349___mcc_h2
			var _350_cf14 _dafny.Int = _349___mcc_h2
			_ = _350_cf14
			return (_this).F13()
		} else {
			var _351___mcc_h3 _dafny.CodePoint = _source8.Get_().(D3_DC5).Cf11
			_ = _351___mcc_h3
			var _352_cf11 _dafny.CodePoint = _351___mcc_h3
			_ = _352_cf11
			return true
		}
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f12 D1
	_f16 D1
	_f14 _dafny.Array
	_f15 _dafny.Set
	_f13 bool
	_f17 bool
	_f21 _dafny.Int
	_f22 bool
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f12 = Companion_D1_.Default()
	_this._f16 = Companion_D1_.Default()
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.EmptySet
	_this._f13 = false
	_this._f17 = false
	_this._f21 = _dafny.Zero
	_this._f22 = false
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C2{}
var _ T2 = &C2{}
var _ T1 = &C2{}
var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F12() D1 {
	return _this._f12
}
func (_this *C2) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C2) F16() D1 {
	return _this._f16
}
func (_this *C2) F16_set_(value D1) {
	_this._f16 = value
}
func (_this *C2) F14() _dafny.Array {
	return _this._f14
}
func (_this *C2) F15() _dafny.Set {
	return _this._f15
}
func (_this *C2) F13() bool {
	return _this._f13
}
func (_this *C2) F17() bool {
	return _this._f17
}
func (_this *C2) Ctor__(f21 _dafny.Int, f22 bool, f16 D1, f17 bool, f14 _dafny.Array, f15 _dafny.Set, f12 D1, f13 bool) {
	{
		(_this)._f21 = f21
		(_this)._f22 = f22
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C2) Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F22()), (_this).F21())).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_dafny.SetOf((_this).F17())).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_this).F21())))
	}
}
func (_this *C2) Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F21()
	}
}
func (_this *C2) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F21()
	}
}
func (_this *C2) M6(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _353_i0 _dafny.Int
		_ = _353_i0
		_353_i0 = _dafny.Zero
		{
			for (_this).F22() {
				{
					if (_353_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_353_i0 = (_353_i0).Plus(_dafny.One)
					var _354_v0 _dafny.Array
					_ = _354_v0
					var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
					_ = _nw38
					_354_v0 = _nw38
					var _355_v1 _dafny.Sequence
					_ = _355_v1
					_355_v1 = _dafny.SeqOf((_this).F17())
					var _356_v2 _dafny.CodePoint
					_ = _356_v2
					_356_v2 = _dafny.CodePoint('k')
					var _357_v3 _dafny.MultiSet
					_ = _357_v3
					_357_v3 = _dafny.MultiSetOf(_dafny.CodePoint('r'), _356_v2, _356_v2)
					var _source9 D2 = Companion_D2_.Create_DC4_(_354_v0, (_this).F22(), Companion_Default___.Fm0(_355_v1, false, (_this).F13(), (_357_v3).Cardinality(), globalState), (_this).F22(), (_this).F14())
					_ = _source9
					if _source9.Is_DC4() {
						var _358___mcc_h0 _dafny.Array = _source9.Get_().(D2_DC4).Cf6
						_ = _358___mcc_h0
						var _359___mcc_h1 bool = _source9.Get_().(D2_DC4).Cf7
						_ = _359___mcc_h1
						var _360___mcc_h2 _dafny.Int = _source9.Get_().(D2_DC4).Cf8
						_ = _360___mcc_h2
						var _361___mcc_h3 bool = _source9.Get_().(D2_DC4).Cf9
						_ = _361___mcc_h3
						var _362___mcc_h4 _dafny.Array = _source9.Get_().(D2_DC4).Cf10
						_ = _362___mcc_h4
						var _363_cf10 _dafny.Array = _362___mcc_h4
						_ = _363_cf10
						var _364_cf9 bool = _361___mcc_h3
						_ = _364_cf9
						var _365_cf8 _dafny.Int = _360___mcc_h2
						_ = _365_cf8
						var _366_cf7 bool = _359___mcc_h1
						_ = _366_cf7
						var _367_cf6 _dafny.Array = _358___mcc_h0
						_ = _367_cf6
						var _368_v4 *C0
						_ = _368_v4
						var _nw39 *C0 = New_C0_()
						_ = _nw39
						_nw39.Ctor__(Companion_Default___.Fm30((_this).F21(), globalState), _dafny.CodePoint('p'), _this.F12(), _364_cf9)
						_368_v4 = _nw39
						var _369_v5 _dafny.MultiSet
						_ = _369_v5
						_369_v5 = _dafny.MultiSetOf((_this).F21(), _dafny.IntOfUint32((_355_v1).Cardinality()))
						var _370_v6 _dafny.Map
						_ = _370_v6
						_370_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC11_(_369_v5), _356_v2)
						var _371_v7 D6
						_ = _371_v7
						_371_v7 = Companion_D6_.Create_DC11_(_369_v5)
						var _372_v8 _dafny.Sequence
						_ = _372_v8
						_372_v8 = _dafny.SeqOf((_this).F21())
						var _373_v9 _dafny.Sequence
						_ = _373_v9
						_373_v9 = _dafny.SeqOf(_372_v8)
						var _374_v10 _dafny.Map
						_ = _374_v10
						_374_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_372_v8, (_373_v9).Select((Companion_Default___.SafeIndex(_365_cf8, _dafny.IntOfUint32((_373_v9).Cardinality()))).Uint32()).(_dafny.Sequence)), _365_cf8)
						var _rhs23 bool = _364_cf9
						_ = _rhs23
						var _rhs24 _dafny.Int = _365_cf8
						_ = _rhs24
						var _rhs25 _dafny.Int = (_368_v4).Fm16(_366_cf7, ((_370_v6).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_371_v7, (_368_v4).F20()))).Cardinality(), globalState)
						_ = _rhs25
						var _rhs26 *C0 = _368_v4
						_ = _rhs26
						var _rhs27 _dafny.Int = (func() _dafny.Int {
							if (_374_v10).Contains(_dafny.Companion_Sequence_.Concatenate(_372_v8, _372_v8)) {
								return (_374_v10).Get(_dafny.Companion_Sequence_.Concatenate(_372_v8, _372_v8)).(_dafny.Int)
							}
							return (_this).F21()
						})()
						_ = _rhs27
						var _lhs21 *GlobalState = globalState
						_ = _lhs21
						var _lhs22 *GlobalState = globalState
						_ = _lhs22
						_366_cf7 = _rhs23
						_lhs21.F7 = _rhs24
						_lhs22.F7 = _rhs25
						_368_v4 = _rhs26
						r0 = _rhs27
						var _375_v11 D2
						_ = _375_v11
						_375_v11 = Companion_D2_.Create_DC3_(_363_cf10)
						var _376_v12 _dafny.Map
						_ = _376_v12
						_376_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
							if (_this).F22() {
								return true
							}
							return (_this).F22()
						})(), _375_v11)
						_376_v12 = (_376_v12).Update((_355_v1).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_355_v1).Cardinality()))).Uint32()).(bool), _375_v11)
						var _377_v13 _dafny.Map
						_ = _377_v13
						_377_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_372_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(677))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg45 _dafny.Int) interface{} {
								return coer45(arg45)
							}
						}(func(_378_i1 _dafny.Int) _dafny.Int {
							return (_dafny.Zero).Minus((_this).F21())
						}))).Cardinality()), _dafny.IntOfUint32((_372_v8).Cardinality()))).Uint32()).(_dafny.Int), _356_v2)
						var _379_v14 _dafny.Map
						_ = _379_v14
						_379_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), _356_v2)
						_377_v13 = (_377_v13).Update((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
							if _366_cf7 {
								return (_this).F21()
							}
							return _365_cf8
						})())), (func() _dafny.CodePoint {
							if (_379_v14).Contains((_this).F17()) {
								return (_379_v14).Get((_this).F17()).(_dafny.CodePoint)
							}
							return (_368_v4).F20()
						})())
						var _380_v15 _dafny.Sequence
						_ = _380_v15
						_380_v15 = _dafny.SeqOf((_368_v4).F20())
						_380_v15 = _380_v15
					} else {
						var _381___mcc_h5 _dafny.Array = _source9.Get_().(D2_DC3).Cf5
						_ = _381___mcc_h5
						var _382_cf5 _dafny.Array = _381___mcc_h5
						_ = _382_cf5
						var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_354_v0), 0))
						_ = _index29
						(_354_v0).ArraySet1((_this).F21(), (_index29).Int())
						var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_354_v0), 0))
						_ = _index30
						(_354_v0).ArraySet1((_this).F21(), (_index30).Int())
						var _383_v16 _dafny.Map
						_ = _383_v16
						_383_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
						var _rhs28 _dafny.Int = ((_354_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_354_v0), 0))).Int()).(_dafny.Int)).Plus(((_383_v16).Merge(_383_v16)).Cardinality())
						_ = _rhs28
						var _rhs29 _dafny.Int = (_354_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_354_v0), 0))).Int()).(_dafny.Int)
						_ = _rhs29
						var _lhs23 *GlobalState = globalState
						_ = _lhs23
						var _lhs24 *GlobalState = globalState
						_ = _lhs24
						_lhs23.F7 = _rhs28
						_lhs24.F7 = _rhs29
						var _384_v17 _dafny.Map
						_ = _384_v17
						_384_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _356_v2)
						var _385_v18 _dafny.Map
						_ = _385_v18
						_385_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F21()).Times(((_this).F15()).Cardinality()), _384_v17)
						_385_v18 = _385_v18
						var _386_v19 _dafny.MultiSet
						_ = _386_v19
						var _out6 _dafny.MultiSet
						_ = _out6
						_out6 = Companion_Default___.M0(_356_v2, !(false), globalState)
						_386_v19 = _out6
					}
					(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_355_v1, _355_v1), _355_v1)
					_356_v2 = _356_v2
					var _387_v20 *C0
					_ = _387_v20
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__((_this).F17(), _356_v2, _this.F12(), (_this).F17())
					_387_v20 = _nw40
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _388_v21 _dafny.Map
		_ = _388_v21
		_388_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(66), (_this).F13())
		var _389_v22 _dafny.Map
		_ = _389_v22
		_389_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_388_v21).Cardinality(), !(false))
		var _390_v23 D6
		_ = _390_v23
		_390_v23 = Companion_D6_.Create_DC12_((_this).F13(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(685))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_391_i2 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		})), (_this).F21(), p0)
		var _392_v24 _dafny.Sequence
		_ = _392_v24
		_392_v24 = _dafny.UnicodeSeqOfUtf8Bytes("mpspf")
		var _pat_let_tv11 = p0
		_ = _pat_let_tv11
		var _pat_let_tv12 = _392_v24
		_ = _pat_let_tv12
		var _pat_let_tv13 = p0
		_ = _pat_let_tv13
		var _pat_let_tv14 = _392_v24
		_ = _pat_let_tv14
		var _393_v25 _dafny.Array
		_ = _393_v25
		var _nwElement0_9 _dafny.Map = _388_v21
		_ = _nwElement0_9
		var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(2))
		_ = _nw41
		(_nw41).ArraySet1(_nwElement0_9, 0)
		(_nw41).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (func(_pat_let4_0 D6) D6 {
			return func(_398_dt__update__tmp_h1 D6) D6 {
				return func(_pat_let9_0 _dafny.Int) D6 {
					return func(_399_dt__update_hcf20_h1 _dafny.Int) D6 {
						return Companion_D6_.Create_DC12_((_398_dt__update__tmp_h1).Dtor_cf18(), (_398_dt__update__tmp_h1).Dtor_cf19(), _399_dt__update_hcf20_h1, (_398_dt__update__tmp_h1).Dtor_cf21())
					}(_pat_let9_0)
				}((_this).F21())
			}(_pat_let4_0)
		}(func(_pat_let5_0 D6) D6 {
			return func(_394_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let6_0 _dafny.Int) D6 {
					return func(_395_dt__update_hcf20_h0 _dafny.Int) D6 {
						return func(_pat_let7_0 bool) D6 {
							return func(_396_dt__update_hcf18_h0 bool) D6 {
								return func(_pat_let8_0 _dafny.Sequence) D6 {
									return func(_397_dt__update_hcf19_h0 _dafny.Sequence) D6 {
										return Companion_D6_.Create_DC12_(_396_dt__update_hcf18_h0, _397_dt__update_hcf19_h0, _395_dt__update_hcf20_h0, (_394_dt__update__tmp_h0).Dtor_cf21())
									}(_pat_let8_0)
								}(_pat_let_tv14)
							}(_pat_let7_0)
						}(false)
					}(_pat_let6_0)
				}((Companion_D6_.Create_DC12_(_pat_let_tv11, _pat_let_tv12, (_this).F21(), !(_pat_let_tv13))).Dtor_cf20())
			}(_pat_let5_0)
		}(_390_v23))).Dtor_cf21()), 1)
		_393_v25 = _nw41
		var _400_v26 _dafny.CodePoint
		_ = _400_v26
		_400_v26 = _dafny.CodePoint('u')
		var _401_v27 T0
		_ = _401_v27
		var _nw42 *C0 = New_C0_()
		_ = _nw42
		_nw42.Ctor__((_this).F17(), _400_v26, _this.F12(), p0)
		_401_v27 = _nw42
		var _402_v28 _dafny.Sequence
		_ = _402_v28
		_402_v28 = _dafny.SeqOf((_this).F15(), (_this).F15(), (_this).F15())
		var _403_v29 _dafny.Map
		_ = _403_v29
		_403_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), p0)
		var _404_v30 _dafny.MultiSet
		_ = _404_v30
		_404_v30 = _dafny.MultiSetOf((_this).F21(), (_403_v29).Cardinality())
		var _405_v31 D6
		_ = _405_v31
		_405_v31 = Companion_D6_.Create_DC11_(_404_v30)
		var _406_v32 _dafny.Map
		_ = _406_v32
		_406_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v27, (_dafny.Zero).Minus(((_402_v28).Select((Companion_Default___.SafeIndex(((_405_v31).Dtor_cf17()).Cardinality(), _dafny.IntOfUint32((_402_v28).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()))
		var _407_v33 _dafny.Sequence
		_ = _407_v33
		_407_v33 = _dafny.SeqOf(_406_v32)
		var _408_v34 _dafny.Sequence
		_ = _408_v34
		_408_v34 = _dafny.SeqOf((_this).F21(), (_this).F21())
		var _409_v35 _dafny.Sequence
		_ = _409_v35
		_409_v35 = _dafny.SeqOf(_dafny.IntOfUint32((_408_v34).Cardinality()), (_this).F21())
		var _410_v36 _dafny.Map
		_ = _410_v36
		_410_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_409_v35).Cardinality()), _392_v24)
		var _rhs30 _dafny.Array = _393_v25
		_ = _rhs30
		var _rhs31 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v27, (_this).F21())).Merge((_407_v33).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_407_v33).Cardinality()))).Uint32()).(_dafny.Map))).Merge((_406_v32).Merge(_406_v32))
		_ = _rhs31
		var _rhs32 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_410_v36).Contains((_this).F21()) {
				return (_410_v36).Get((_this).F21()).(_dafny.Sequence)
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("uxadk")
		})()).Cardinality())
		_ = _rhs32
		var _lhs25 *GlobalState = globalState
		_ = _lhs25
		_393_v25 = _rhs30
		_406_v32 = _rhs31
		_lhs25.F7 = _rhs32
		(globalState).F7 = (_this).F21()
		_389_v22 = (_389_v22).Update(Companion_Default___.SafeDivisionInt((_this).Fm9(_400_v26, globalState), (_this).F21()), (_this).F17())
		var _411_v37 *C0
		_ = _411_v37
		var _nw43 *C0 = New_C0_()
		_ = _nw43
		_nw43.Ctor__(!((_this).F17()), _400_v26, _this.F12(), (_401_v27).F13())
		_411_v37 = _nw43
		if (_this).F22() {
			var _412_v38 D3
			_ = _412_v38
			_412_v38 = Companion_D3_.Create_DC8_((_this).F21())
			var _413_v39 _dafny.Map
			_ = _413_v39
			_413_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_401_v27).F13(), (_this).F21())
			var _414_v40 _dafny.Map
			_ = _414_v40
			_414_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_401_v27).F13(), _400_v26)
			var _415_v41 _dafny.MultiSet
			_ = _415_v41
			_415_v41 = _dafny.MultiSetOf((_this).F22(), (_401_v27).F13())
			var _416_v43 _dafny.Map
			_ = _416_v43
			_416_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_411_v37).F20(), (_this).F21())
			var _417_v44 _dafny.Sequence
			_ = _417_v44
			_417_v44 = _dafny.SeqOf(func() _dafny.Map {
				var _coll31 = _dafny.NewMapBuilder()
				_ = _coll31
				for _iter36 := _dafny.Iterate((_416_v43).Keys().Elements()); ; {
					_compr_31, _ok36 := _iter36()
					if !_ok36 {
						break
					}
					var _418_v42 _dafny.CodePoint
					_418_v42 = interface{}(_compr_31).(_dafny.CodePoint)
					if (_416_v43).Contains(_418_v42) {
						_coll31.Add(_418_v42, (_this).F21())
					}
				}
				return _coll31.ToMap()
			}())
			var _419_v45 _dafny.Set
			_ = _419_v45
			_419_v45 = _dafny.SetOf(false, (_401_v27).F13())
			var _420_v46 _dafny.Map
			_ = _420_v46
			_420_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_403_v29, _419_v45)
			var _421_v47 _dafny.Sequence
			_ = _421_v47
			_421_v47 = _dafny.SeqOf(_411_v37.F19)
			var _422_v48 _dafny.Array
			_ = _422_v48
			var _nwElement0_10 _dafny.Int = (func(_pat_let10_0 D3) D3 {
				return func(_423_dt__update__tmp_h2 D3) D3 {
					return func(_pat_let11_0 _dafny.Int) D3 {
						return func(_424_dt__update_hcf14_h0 _dafny.Int) D3 {
							return Companion_D3_.Create_DC8_(_424_dt__update_hcf14_h0)
						}(_pat_let11_0)
					}(_dafny.IntOfInt64(381))
				}(_pat_let10_0)
			}(_412_v38)).Dtor_cf14()
			_ = _nwElement0_10
			var _nw44 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(23))
			_ = _nw44
			(_nw44).ArraySet1(_nwElement0_10, 0)
			(_nw44).ArraySet1(((_dafny.Zero).Minus((_this).F21())).Plus((_this).F21()), 1)
			(_nw44).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_408_v34, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_408_v34).Cardinality()))).Uint32(), (_this).F21()), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_408_v34, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_408_v34).Cardinality()))).Uint32(), (_this).F21())).Cardinality()))).Uint32(), (_this).F21())).Cardinality()), 2)
			(_nw44).ArraySet1((_this).F21(), 3)
			(_nw44).ArraySet1((_dafny.Zero).Minus((_this).F21()), 4)
			(_nw44).ArraySet1((_this).F21(), 5)
			(_nw44).ArraySet1(Companion_Default___.SafeModuloInt((_this).F21(), (_this).F21()), 6)
			(_nw44).ArraySet1((func() _dafny.Int {
				if (_413_v39).Contains((_401_v27).F13()) {
					return (_413_v39).Get((_401_v27).F13()).(_dafny.Int)
				}
				return (_this).F21()
			})(), 7)
			(_nw44).ArraySet1(_dafny.IntOfInt64(670), 8)
			(_nw44).ArraySet1(_dafny.IntOfUint32((_392_v24).Cardinality()), 9)
			(_nw44).ArraySet1((_this).F21(), 10)
			(_nw44).ArraySet1((_this).F21(), 11)
			(_nw44).ArraySet1((_this).F21(), 12)
			(_nw44).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v24, _414_v40)).Merge(Companion_Default___.Fm31(globalState))).Cardinality(), 13)
			(_nw44).ArraySet1((_this).F21(), 14)
			(_nw44).ArraySet1(((_415_v41).Cardinality()).Times(_dafny.IntOfUint32((_417_v44).Cardinality())), 15)
			(_nw44).ArraySet1(Companion_Default___.SafeModuloInt((_this).F21(), (_this).F21()), 16)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_392_v24).Cardinality()), (_dafny.Zero).Minus((_this).F21())), 17)
			(_nw44).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F21(), (_this).F21()), 18)
			(_nw44).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.Set {
				if (_420_v46).Contains(_403_v29) {
					return (_420_v46).Get(_403_v29).(_dafny.Set)
				}
				return _dafny.SetOf(_411_v37.F19)
			})())).Cardinality())).Times((_this).F21()), 19)
			(_nw44).ArraySet1(Companion_Default___.SafeModuloInt((Companion_D3_.Create_DC6_((_dafny.Zero).Minus((_this).F21()), (_401_v27).F13())).Dtor_cf12(), _dafny.IntOfInt64(703)), 20)
			(_nw44).ArraySet1(((_this).F15()).Cardinality(), 21)
			(_nw44).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_421_v47, _421_v47)).Cardinality()), 22)
			_422_v48 = _nw44
			var _425_v49 _dafny.Map
			_ = _425_v49
			_425_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_400_v26, _411_v37.F19)
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_422_v48), 0))
			_ = _index31
			(_422_v48).ArraySet1(Companion_Default___.SafeDivisionInt((_425_v49).Cardinality(), (_this).F21()), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_422_v48), 0))
			_ = _index32
			(_422_v48).ArraySet1((_dafny.Zero).Minus((_this).F21()), (_index32).Int())
			var _426_v50 *C0
			_ = _426_v50
			var _nw45 *C0 = New_C0_()
			_ = _nw45
			_nw45.Ctor__((_401_v27).F13(), (_411_v37).F20(), _401_v27.F12(), (_this).F17())
			_426_v50 = _nw45
			var _427_v51 _dafny.Array
			_ = _427_v51
			var _nw46 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
			_ = _nw46
			_427_v51 = _nw46
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_427_v51), 0))
			_ = _index33
			(_427_v51).ArraySet1(_392_v24, (_index33).Int())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_427_v51), 0))
			_ = _index34
			(_427_v51).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_392_v24, (func() _dafny.Sequence {
				if false {
					return (func() _dafny.Sequence {
						if (_410_v36).Contains((_422_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_422_v48), 0))).Int()).(_dafny.Int)) {
							return (_410_v36).Get((_422_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_422_v48), 0))).Int()).(_dafny.Int)).(_dafny.Sequence)
						}
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(611))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg47 _dafny.Int) interface{} {
								return coer47(arg47)
							}
						}((func(_428_v37 *C0) func(_dafny.Int) _dafny.CodePoint {
							return func(_429_i3 _dafny.Int) _dafny.CodePoint {
								return (_428_v37).F20()
							}
						})(_411_v37)))
					})()
				}
				return _392_v24
			})()), (_index34).Int())
			_422_v48 = _422_v48
			(globalState).F7 = (_422_v48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(442), _dafny.ArrayLen((_422_v48), 0))).Int()).(_dafny.Int)
		} else {
			var _430_v52 _dafny.MultiSet
			_ = _430_v52
			_430_v52 = _dafny.MultiSetOf((_402_v28).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_402_v28).Cardinality()))).Uint32()).(_dafny.Set))
			var _431_v53 _dafny.Sequence
			_ = _431_v53
			_431_v53 = _dafny.SeqOf((_401_v27).F13(), (_this).F22())
			var _432_v54 _dafny.Map
			_ = _432_v54
			_432_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_401_v27).F13(), _430_v52)
			_430_v52 = (_dafny.MultiSetOf(_dafny.SetOf(_dafny.IntOfUint32((_431_v53).Cardinality())))).Intersection((func() _dafny.MultiSet {
				if (_432_v54).Contains(!(p0)) {
					return (_432_v54).Get(!(p0)).(_dafny.MultiSet)
				}
				return _430_v52
			})())
			var _433_v56 _dafny.Map
			_ = _433_v56
			_433_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll32 = _dafny.NewBuilder()
				_ = _coll32
				for _iter37 := _dafny.Iterate((_dafny.MultiSetOf((_this).F21())).Elements()); ; {
					_compr_32, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _434_v55 _dafny.Int
					_434_v55 = interface{}(_compr_32).(_dafny.Int)
					if (_dafny.MultiSetOf((_this).F21())).Contains(_434_v55) {
						_coll32.Add(Companion_Default___.SafeModuloInt(_434_v55, _dafny.IntOfInt64(781)))
					}
				}
				return _coll32.ToSet()
			}()).Union((_this).F15()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v24, (_411_v37).F20()))
			var _435_v57 _dafny.Map
			_ = _435_v57
			_435_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v24, _400_v26)
			_433_v56 = (_433_v56).Update((_this).F15(), _435_v57)
			var _436_v58 _dafny.Array
			_ = _436_v58
			var _nw47 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
			_ = _nw47
			_436_v58 = _nw47
			var _437_v59 _dafny.Map
			_ = _437_v59
			_437_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_431_v53, (_this).F21())
			var _438_v60 D7
			_ = _438_v60
			_438_v60 = Companion_D7_.Create_DC15_(_dafny.SeqOf((func() _dafny.Int {
				if (_437_v59).Contains(_431_v53) {
					return (_437_v59).Get(_431_v53).(_dafny.Int)
				}
				return (_this).F21()
			})()))
			var _439_v61 _dafny.Array
			_ = _439_v61
			var _nwElement0_11 _dafny.Sequence = _408_v34
			_ = _nwElement0_11
			var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(16))
			_ = _nw48
			(_nw48).ArraySet1(_nwElement0_11, 0)
			(_nw48).ArraySet1(_408_v34, 1)
			(_nw48).ArraySet1(_409_v35, 2)
			(_nw48).ArraySet1(_409_v35, 3)
			(_nw48).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(389))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}((func(_440_v37 *C0) func(_dafny.Int) _dafny.Int {
				return func(_441_i4 _dafny.Int) _dafny.Int {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_440_v37).F20(), (_440_v37).F20()), true)).Cardinality()
				}
			})(_411_v37))), 4)
			(_nw48).ArraySet1(_408_v34, 5)
			(_nw48).ArraySet1(_409_v35, 6)
			(_nw48).ArraySet1((_438_v60).Dtor_cf28(), 7)
			(_nw48).ArraySet1(_408_v34, 8)
			(_nw48).ArraySet1(_409_v35, 9)
			(_nw48).ArraySet1(_409_v35, 10)
			(_nw48).ArraySet1(_409_v35, 11)
			(_nw48).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(954))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_442_i5 _dafny.Int) _dafny.Int {
				return (_this).F21()
			})), 12)
			(_nw48).ArraySet1(_408_v34, 13)
			(_nw48).ArraySet1(_408_v34, 14)
			(_nw48).ArraySet1(_408_v34, 15)
			_439_v61 = _nw48
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_436_v58), 0))
			_ = _index35
			(_436_v58).ArraySet1(_439_v61, (_index35).Int())
			var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_436_v58), 0))
			_ = _index36
			(_436_v58).ArraySet1(_439_v61, (_index36).Int())
			(_411_v37).F19 = ((func() _dafny.Int {
				if (_404_v30).Contains((_this).F21()) {
					return (_404_v30).Multiplicity((_this).F21())
				}
				return (_this).F21()
			})()).Cmp((func() _dafny.Int {
				if (_this).F17() {
					return _dafny.IntOfUint32((_392_v24).Cardinality())
				}
				return (_this).F21()
			})()) >= 0
			if !(_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
				if (_this.F12()).Dtor_cf4() {
					return _392_v24
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("agacmtia")
			})(), _392_v24)) {
				var _443_v62 _dafny.Sequence
				_ = _443_v62
				_443_v62 = _dafny.SeqOf((_this).F14())
				var _444_v63 _dafny.Set
				_ = _444_v63
				_444_v63 = _dafny.SetOf(false, _411_v37.F19, true)
				var _445_v64 _dafny.Map
				_ = _445_v64
				_445_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_443_v62).Select((Companion_Default___.SafeIndex((_444_v63).Cardinality(), _dafny.IntOfUint32((_443_v62).Cardinality()))).Uint32()).(_dafny.Array))
				var _446_v65 D6
				_ = _446_v65
				_446_v65 = Companion_D6_.Create_DC11_(_404_v30)
				var _447_v66 D6
				_ = _447_v66
				_447_v66 = Companion_D6_.Create_DC14_(_446_v65)
				var _448_v67 _dafny.Array
				_ = _448_v67
				var _nwElement0_12 _dafny.Array = (_this).F14()
				_ = _nwElement0_12
				var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(20))
				_ = _nw49
				(_nw49).ArraySet1(_nwElement0_12, 0)
				(_nw49).ArraySet1((_this).F14(), 1)
				(_nw49).ArraySet1((_this).F14(), 2)
				(_nw49).ArraySet1((_this).F14(), 3)
				(_nw49).ArraySet1((_this).F14(), 4)
				(_nw49).ArraySet1((_this).F14(), 5)
				(_nw49).ArraySet1((_this).F14(), 6)
				(_nw49).ArraySet1((_this).F14(), 7)
				(_nw49).ArraySet1((_this).F14(), 8)
				(_nw49).ArraySet1((func() _dafny.Array {
					if (_445_v64).Contains((_this).F21()) {
						return (_445_v64).Get((_this).F21()).(_dafny.Array)
					}
					return (_this).F14()
				})(), 9)
				(_nw49).ArraySet1((_this).F14(), 10)
				(_nw49).ArraySet1((_this).F14(), 11)
				(_nw49).ArraySet1((_this).F14(), 12)
				(_nw49).ArraySet1((_this).F14(), 13)
				(_nw49).ArraySet1((_this).F14(), 14)
				(_nw49).ArraySet1((_this).F14(), 15)
				(_nw49).ArraySet1((_this).F14(), 16)
				(_nw49).ArraySet1((_this).F14(), 17)
				(_nw49).ArraySet1((_this).F14(), 18)
				(_nw49).ArraySet1((_443_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm32(_447_v66, globalState)).Cardinality()), _dafny.IntOfUint32((_443_v62).Cardinality()))).Uint32()).(_dafny.Array), 19)
				_448_v67 = _nw49
				var _449_v68 _dafny.Array
				_ = _449_v68
				_449_v68 = _448_v67
				_449_v68 = _449_v68
				(_411_v37).F19 = false
				(globalState).F7 = ((_this).F21()).Plus((_this).F21())
				var _450_v69 _dafny.Array
				_ = _450_v69
				var _len0_3 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_3
				var _nw50 _dafny.Array
				_ = _nw50
				if _len0_3.Cmp(_dafny.Zero) == 0 {
					_nw50 = _dafny.NewArray(_len0_3)
				} else {
					var _init3 func(_dafny.Int) _dafny.Int = func(_451_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_451_i6, (_this).F21())
					}
					_ = _init3
					var _element0_3 = _init3(_dafny.Zero)
					_ = _element0_3
					_nw50 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
					(_nw50).ArraySet1(_element0_3, 0)
					var _nativeLen0_3 = (_len0_3).Int()
					_ = _nativeLen0_3
					for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
						(_nw50).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
					}
				}
				_450_v69 = _nw50
				_450_v69 = (func() _dafny.Array {
					if (_this).F22() {
						return _450_v69
					}
					return _450_v69
				})()
				var _452_v70 _dafny.Sequence
				_ = _452_v70
				_452_v70 = _dafny.SeqOf(_390_v23)
				var _453_v71 _dafny.Sequence
				_ = _453_v71
				_453_v71 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(541))).Uint32(), func(coer50 func(_dafny.Int) D6) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}((func(_454_v23 D6) func(_dafny.Int) D6 {
					return func(_455_i7 _dafny.Int) D6 {
						return _454_v23
					}
				})(_390_v23))), _452_v70)
				var _456_v72 _dafny.Map
				_ = _456_v72
				_456_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _411_v37.F19)).Cardinality(), (_453_v71).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_453_v71).Cardinality()))).Uint32()).(_dafny.Sequence))
				var _457_v73 *C0
				_ = _457_v73
				var _nw51 *C0 = New_C0_()
				_ = _nw51
				_nw51.Ctor__((func() bool {
					if (_this).F17() {
						return (_401_v27).F13()
					}
					return (_401_v27).F13()
				})(), _400_v26, _401_v27.F12(), !(_456_v72).Contains(_dafny.IntOfInt64(693)))
				_457_v73 = _nw51
			} else {
				var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index37
				((_this).F14()).ArraySet1((_401_v27).F13(), (_index37).Int())
				var _458_v74 _dafny.Map
				_ = _458_v74
				_458_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
				var _459_v75 _dafny.Set
				_ = _459_v75
				_459_v75 = _dafny.SetOf(_411_v37.F19, p0, (_this).F17(), false)
				var _460_v76 _dafny.Sequence
				_ = _460_v76
				_460_v76 = _dafny.SeqOf(_392_v24)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index38
				var _rhs33 bool = (_401_v27).F13()
				_ = _rhs33
				var _rhs34 bool = (_459_v75).Equals((_459_v75).Union(_459_v75))
				_ = _rhs34
				var _rhs35 _dafny.Map = _458_v74
				_ = _rhs35
				var _rhs36 bool = (func() bool {
					if (_this).F22() {
						return (_431_v53).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_431_v53).Cardinality()))).Uint32()).(bool)
					}
					return !(((_this).F21()).Cmp(_dafny.IntOfUint32((_460_v76).Cardinality())) >= 0)
				})()
				_ = _rhs36
				var _lhs26 *C0 = _411_v37
				_ = _lhs26
				var _lhs27 _dafny.Array = (_this).F14()
				_ = _lhs27
				var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(376), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs28
				var _lhs29 *C0 = _411_v37
				_ = _lhs29
				_lhs26.F19 = _rhs33
				(_lhs27).ArraySet1(_rhs34, (_lhs28).Int())
				_458_v74 = _rhs35
				_lhs29.F19 = _rhs36
				_392_v24 = _392_v24
				var _461_v77 _dafny.Array
				_ = _461_v77
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(21)
				_ = _len0_4
				var _nw52 _dafny.Array
				_ = _nw52
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw52 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) _dafny.Int = func(_462_i8 _dafny.Int) _dafny.Int {
						return (_462_i8).Minus((_this).F21())
					}
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw52 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw52).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw52).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_461_v77 = _nw52
				_461_v77 = _461_v77
				var _463_v78 _dafny.Array
				_ = _463_v78
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_5
				var _nw53 _dafny.Array
				_ = _nw53
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw53 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.CodePoint = (func(_464_v37 *C0) func(_dafny.Int) _dafny.CodePoint {
						return func(_465_i9 _dafny.Int) _dafny.CodePoint {
							return (_464_v37).F20()
						}
					})(_411_v37)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw53 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw53).ArraySet1CodePoint(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw53).ArraySet1CodePoint(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_463_v78 = _nw53
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_463_v78), 0))
				_ = _index39
				(_463_v78).ArraySet1CodePoint((_411_v37).F20(), (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_463_v78), 0))
				_ = _index40
				var _rhs37 _dafny.CodePoint = (_411_v37).F20()
				_ = _rhs37
				var _rhs38 _dafny.Sequence = _431_v53
				_ = _rhs38
				var _rhs39 _dafny.CodePoint = (_411_v37).F20()
				_ = _rhs39
				var _lhs30 _dafny.Array = _463_v78
				_ = _lhs30
				var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_463_v78), 0))
				_ = _lhs31
				var _lhs32 *GlobalState = globalState
				_ = _lhs32
				(_lhs30).ArraySet1CodePoint(_rhs37, (_lhs31).Int())
				_lhs32.F2 = _rhs38
				_400_v26 = _rhs39
				(globalState).F7 = (_dafny.Zero).Minus((_this).F21())
			}
		}
		r0 = _dafny.IntOfInt64(139)
		return r0
	}
}
func (_this *C2) M7(globalState *GlobalState) {
	{
		var _466_v0 _dafny.Sequence
		_ = _466_v0
		_466_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rxto")
		var _467_v1 _dafny.Array
		_ = _467_v1
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(9)
		_ = _len0_6
		var _nw54 _dafny.Array
		_ = _nw54
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw54 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = func(_468_i0 _dafny.Int) _dafny.Int {
				return (_468_i0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-952))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_469_i1 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))).Cardinality()))
			}
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw54 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw54).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw54).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_467_v1 = _nw54
		var _470_v2 _dafny.Sequence
		_ = _470_v2
		_470_v2 = _dafny.SeqOf(_467_v1, _467_v1, _467_v1)
		var _471_v3 D6
		_ = _471_v3
		_471_v3 = Companion_D6_.Create_DC12_((_this).F17(), _466_v0, _dafny.IntOfUint32((_470_v2).Cardinality()), (_this).F22())
		var _source10 D6 = _471_v3
		_ = _source10
		if _source10.Is_DC12() {
			var _472___mcc_h0 bool = _source10.Get_().(D6_DC12).Cf18
			_ = _472___mcc_h0
			var _473___mcc_h1 _dafny.Sequence = _source10.Get_().(D6_DC12).Cf19
			_ = _473___mcc_h1
			var _474___mcc_h2 _dafny.Int = _source10.Get_().(D6_DC12).Cf20
			_ = _474___mcc_h2
			var _475___mcc_h3 bool = _source10.Get_().(D6_DC12).Cf21
			_ = _475___mcc_h3
			var _476_cf21 bool = _475___mcc_h3
			_ = _476_cf21
			var _477_cf20 _dafny.Int = _474___mcc_h2
			_ = _477_cf20
			var _478_cf19 _dafny.Sequence = _473___mcc_h1
			_ = _478_cf19
			var _479_cf18 bool = _472___mcc_h0
			_ = _479_cf18
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))
			_ = _index41
			(_467_v1).ArraySet1((_this).F21(), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))
			_ = _index42
			(_467_v1).ArraySet1(_477_cf20, (_index42).Int())
			var _480_v4 _dafny.Map
			_ = _480_v4
			_480_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_477_cf20, (_this).F14())
			var _481_v5 _dafny.Map
			_ = _481_v5
			_481_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_476_cf21, (_480_v4).Cardinality())
			var _482_v6 _dafny.Map
			_ = _482_v6
			_482_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_481_v5, (_dafny.MultiSetOf(_dafny.IntOfInt64(877))).Cardinality())
			var _483_v7 _dafny.Map
			_ = _483_v7
			_483_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_479_cf18, false)
			var _484_v8 _dafny.Sequence
			_ = _484_v8
			_484_v8 = _dafny.SeqOf((_this).F21(), (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _477_cf20, (_this).F21(), _477_cf20)
			var _485_v9 _dafny.MultiSet
			_ = _485_v9
			_485_v9 = _dafny.MultiSetOf(_477_cf20)
			var _486_v10 _dafny.Array
			_ = _486_v10
			var _nwElement0_13 _dafny.Int = (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int)
			_ = _nwElement0_13
			var _nw55 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(18))
			_ = _nw55
			(_nw55).ArraySet1(_nwElement0_13, 0)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_478_cf19).Cardinality()), 1)
			(_nw55).ArraySet1((func() _dafny.Int {
				if true {
					return (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_482_v6).Contains(_481_v5) {
						return (_482_v6).Get(_481_v5).(_dafny.Int)
					}
					return (_this).F21()
				})()
			})(), 2)
			(_nw55).ArraySet1(((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int)).Plus((_483_v7).Cardinality()), 3)
			(_nw55).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 4)
			(_nw55).ArraySet1(_477_cf20, 5)
			(_nw55).ArraySet1((_477_cf20).Minus((_this).F21()), 6)
			(_nw55).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 7)
			(_nw55).ArraySet1((_this).F21(), 8)
			(_nw55).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_484_v8, (Companion_Default___.SafeIndex((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_484_v8).Cardinality()))).Uint32(), _477_cf20), _484_v8)).Cardinality()), 9)
			(_nw55).ArraySet1((_477_cf20).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_484_v8, (Companion_Default___.SafeIndex(_477_cf20, _dafny.IntOfUint32((_484_v8).Cardinality()))).Uint32(), _477_cf20)).Cardinality())), 10)
			(_nw55).ArraySet1(Companion_Default___.SafeModuloInt(_477_cf20, _477_cf20), 11)
			(_nw55).ArraySet1((func() _dafny.Int {
				if true {
					return _dafny.IntOfInt64(547)
				}
				return _477_cf20
			})(), 12)
			(_nw55).ArraySet1((_485_v9).Cardinality(), 13)
			(_nw55).ArraySet1((_484_v8).Select((Companion_Default___.SafeIndex(_477_cf20, _dafny.IntOfUint32((_484_v8).Cardinality()))).Uint32()).(_dafny.Int), 14)
			(_nw55).ArraySet1(_dafny.IntOfInt64(865), 15)
			(_nw55).ArraySet1(_477_cf20, 16)
			(_nw55).ArraySet1((_this).F21(), 17)
			_486_v10 = _nw55
			_486_v10 = (Companion_D2_.Create_DC4_(_486_v10, true, (_this).F21(), (_this).F22(), (_this).F14())).Dtor_cf6()
			(globalState).F7 = _dafny.IntOfInt64(527)
			var _487_v11 _dafny.Sequence
			_ = _487_v11
			_487_v11 = _dafny.SeqOf(!((_this).F13()), _479_cf18)
			var _488_v12 *C0
			_ = _488_v12
			var _nw56 *C0 = New_C0_()
			_ = _nw56
			_nw56.Ctor__((func() bool {
				if (_this).F22() {
					return (_487_v11).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F21()), _dafny.IntOfUint32((_487_v11).Cardinality()))).Uint32()).(bool)
				}
				return false
			})(), (_466_v0).Select((Companion_Default___.SafeIndex((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(186), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_466_v0).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F12(), _479_cf18)
			_488_v12 = _nw56
		} else if _source10.Is_DC13() {
			var _489___mcc_h4 bool = _source10.Get_().(D6_DC13).Cf22
			_ = _489___mcc_h4
			var _490___mcc_h5 bool = _source10.Get_().(D6_DC13).Cf23
			_ = _490___mcc_h5
			var _491___mcc_h6 _dafny.Int = _source10.Get_().(D6_DC13).Cf24
			_ = _491___mcc_h6
			var _492___mcc_h7 _dafny.Int = _source10.Get_().(D6_DC13).Cf25
			_ = _492___mcc_h7
			var _493___mcc_h8 bool = _source10.Get_().(D6_DC13).Cf26
			_ = _493___mcc_h8
			var _494_cf26 bool = _493___mcc_h8
			_ = _494_cf26
			var _495_cf25 _dafny.Int = _492___mcc_h7
			_ = _495_cf25
			var _496_cf24 _dafny.Int = _491___mcc_h6
			_ = _496_cf24
			var _497_cf23 bool = _490___mcc_h5
			_ = _497_cf23
			var _498_cf22 bool = _489___mcc_h4
			_ = _498_cf22
			_497_cf23 = (_this).F22()
			var _499_v13 _dafny.Set
			_ = _499_v13
			_499_v13 = _dafny.SetOf((_this).F17())
			var _500_v14 bool
			_ = _500_v14
			var _out7 bool
			_ = _out7
			_out7 = (_this).M13((_this).F13(), Companion_Default___.Fm33((_this).F13(), _495_cf25, globalState), _499_v13, (_dafny.Zero).Minus(_496_cf24), globalState)
			_500_v14 = _out7
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_467_v1), 0))
			_ = _index43
			(_467_v1).ArraySet1(_496_cf24, (_index43).Int())
			var _501_v15 _dafny.Sequence
			_ = _501_v15
			_501_v15 = _dafny.SeqOf((_this).F17())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.ArrayLen((_467_v1), 0))
			_ = _index44
			(_467_v1).ArraySet1((_dafny.Zero).Minus((_dafny.SetOf(!((func() bool {
				if (_501_v15).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_501_v15).Cardinality()))).Uint32()).(bool) {
					return (_this).F22()
				}
				return !(_497_cf23)
			})()), (_495_cf25).Cmp((_this).F21()) == 0, (_494_cf26) || (_500_v14))).Cardinality()), (_index44).Int())
			_500_v14 = (_497_cf23) == ((_this).F13())
		} else if _source10.Is_DC11() {
			var _502___mcc_h9 _dafny.MultiSet = _source10.Get_().(D6_DC11).Cf17
			_ = _502___mcc_h9
			var _503_cf17 _dafny.MultiSet = _502___mcc_h9
			_ = _503_cf17
			var _504_v16 _dafny.Map
			_ = _504_v16
			_504_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).F21())
			(globalState).F7 = (func() _dafny.Int {
				if (_504_v16).Contains((_this).F22()) {
					return (_504_v16).Get((_this).F22()).(_dafny.Int)
				}
				return (_this).F21()
			})()
			var _505_v17 _dafny.Map
			_ = _505_v17
			_505_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _this.F12())
			var _506_v18 _dafny.CodePoint
			_ = _506_v18
			_506_v18 = _dafny.CodePoint('c')
			var _507_v19 *C0
			_ = _507_v19
			var _nw57 *C0 = New_C0_()
			_ = _nw57
			_nw57.Ctor__(!(_505_v17).Equals(_505_v17), _506_v18, Companion_D1_.Create_DC2_((_this).F21(), (_this).F21(), true), (_this).F17())
			_507_v19 = _nw57
			var _508_v20 D6
			_ = _508_v20
			_508_v20 = Companion_D6_.Create_DC11_(_503_cf17)
			var _source11 D6 = _508_v20
			_ = _source11
			if _source11.Is_DC12() {
				var _509___mcc_h11 bool = _source11.Get_().(D6_DC12).Cf18
				_ = _509___mcc_h11
				var _510___mcc_h12 _dafny.Sequence = _source11.Get_().(D6_DC12).Cf19
				_ = _510___mcc_h12
				var _511___mcc_h13 _dafny.Int = _source11.Get_().(D6_DC12).Cf20
				_ = _511___mcc_h13
				var _512___mcc_h14 bool = _source11.Get_().(D6_DC12).Cf21
				_ = _512___mcc_h14
				var _513_cf21 bool = _512___mcc_h14
				_ = _513_cf21
				var _514_cf20 _dafny.Int = _511___mcc_h13
				_ = _514_cf20
				var _515_cf19 _dafny.Sequence = _510___mcc_h12
				_ = _515_cf19
				var _516_cf18 bool = _509___mcc_h11
				_ = _516_cf18
				var _517_v21 _dafny.Array
				_ = _517_v21
				var _len0_7 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_7
				var _nw58 _dafny.Array
				_ = _nw58
				if _len0_7.Cmp(_dafny.Zero) == 0 {
					_nw58 = _dafny.NewArray(_len0_7)
				} else {
					var _init7 func(_dafny.Int) _dafny.CodePoint = func(_518_i2 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('t')
					}
					_ = _init7
					var _element0_7 = _init7(_dafny.Zero)
					_ = _element0_7
					_nw58 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
					(_nw58).ArraySet1CodePoint(_element0_7, 0)
					var _nativeLen0_7 = (_len0_7).Int()
					_ = _nativeLen0_7
					for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
						(_nw58).ArraySet1CodePoint(_init7(_dafny.IntOf(_i0_7)), _i0_7)
					}
				}
				_517_v21 = _nw58
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_517_v21), 0))
				_ = _index45
				(_517_v21).ArraySet1CodePoint(_506_v18, (_index45).Int())
				var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_517_v21), 0))
				_ = _index46
				(_517_v21).ArraySet1CodePoint(_dafny.CodePoint('v'), (_index46).Int())
				var _519_v22 _dafny.MultiSet
				_ = _519_v22
				_519_v22 = _dafny.MultiSetOf((_this).F13(), _516_cf18)
				var _520_v23 _dafny.Sequence
				_ = _520_v23
				_520_v23 = _dafny.SeqOf(_516_cf18, _513_cf21)
				var _521_v24 _dafny.Sequence
				_ = _521_v24
				_521_v24 = _dafny.SeqOf(_514_cf20, _dafny.IntOfUint32((_466_v0).Cardinality()))
				var _522_v25 D6
				_ = _522_v25
				_522_v25 = Companion_D6_.Create_DC14_(Companion_D6_.Create_DC11_(_dafny.MultiSetFromSeq(_521_v24)))
				var _523_v26 D6
				_ = _523_v26
				_523_v26 = Companion_D6_.Create_DC14_(_522_v25)
				var _524_v27 _dafny.Map
				_ = _524_v27
				_524_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((func() _dafny.Int {
					if (_519_v22).Contains((_this).F17()) {
						return (_519_v22).Multiplicity((_this).F17())
					}
					return Companion_Default___.Fm0(_520_v23, (_this).F13(), (_this).F22(), _514_cf20, globalState)
				})()).Plus(_dafny.IntOfUint32((_521_v24).Cardinality())), (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm34(_514_cf20, globalState), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((Companion_Default___.Fm34(_514_cf20, globalState)).Cardinality()))).Uint32(), (_this).F21())).Cardinality())).Minus(_dafny.IntOfUint32((Companion_Default___.Fm32(_523_v26, globalState)).Cardinality())))
				_524_v27 = (_524_v27).Update(_514_cf20, _dafny.IntOfUint32((_515_cf19).Cardinality()))
				var _525_v28 _dafny.Array
				_ = _525_v28
				var _nw59 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(12))
				_ = _nw59
				_525_v28 = _nw59
				var _526_v29 D7
				_ = _526_v29
				_526_v29 = Companion_D7_.Create_DC16_(_514_cf20)
				var _527_v30 D7
				_ = _527_v30
				_527_v30 = Companion_D7_.Create_DC17_(_526_v29)
				var _528_v31 D7
				_ = _528_v31
				_528_v31 = Companion_D7_.Create_DC17_(_527_v30)
				var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_525_v28), 0))
				_ = _index47
				(_525_v28).ArraySet1(_528_v31, (_index47).Int())
				var _529_v32 _dafny.Map
				_ = _529_v32
				_529_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _515_cf19)
				var _530_v33 _dafny.Sequence
				_ = _530_v33
				_530_v33 = _dafny.SeqOf(_466_v0, _466_v0, _515_cf19, _dafny.Companion_Sequence_.Concatenate(_466_v0, _515_cf19), _466_v0)
				var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_525_v28), 0))
				_ = _index48
				var _rhs40 _dafny.Int = _514_cf20
				_ = _rhs40
				var _rhs41 _dafny.Int = _dafny.IntOfUint32(((_530_v33).Select((Companion_Default___.SafeIndex((_521_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(391), _dafny.IntOfUint32((_521_v24).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_530_v33).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())
				_ = _rhs41
				var _rhs42 D7 = _528_v31
				_ = _rhs42
				var _rhs43 _dafny.Map = Companion_Default___.Fm35((_471_v3).Dtor_cf20(), (func() _dafny.Int {
					if _516_cf18 {
						return _dafny.IntOfUint32((_520_v23).Cardinality())
					}
					return _514_cf20
				})(), (_507_v19).F20(), !(Companion_Default___.Fm30((_dafny.Zero).Minus((_this).F21()), globalState)), globalState)
				_ = _rhs43
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				var _lhs35 _dafny.Array = _525_v28
				_ = _lhs35
				var _lhs36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_525_v28), 0))
				_ = _lhs36
				_lhs33.F7 = _rhs40
				_lhs34.F7 = _rhs41
				(_lhs35).ArraySet1(_rhs42, (_lhs36).Int())
				_529_v32 = _rhs43
				_523_v26 = _523_v26
			} else if _source11.Is_DC13() {
				var _531___mcc_h15 bool = _source11.Get_().(D6_DC13).Cf22
				_ = _531___mcc_h15
				var _532___mcc_h16 bool = _source11.Get_().(D6_DC13).Cf23
				_ = _532___mcc_h16
				var _533___mcc_h17 _dafny.Int = _source11.Get_().(D6_DC13).Cf24
				_ = _533___mcc_h17
				var _534___mcc_h18 _dafny.Int = _source11.Get_().(D6_DC13).Cf25
				_ = _534___mcc_h18
				var _535___mcc_h19 bool = _source11.Get_().(D6_DC13).Cf26
				_ = _535___mcc_h19
				var _536_cf26 bool = _535___mcc_h19
				_ = _536_cf26
				var _537_cf25 _dafny.Int = _534___mcc_h18
				_ = _537_cf25
				var _538_cf24 _dafny.Int = _533___mcc_h17
				_ = _538_cf24
				var _539_cf23 bool = _532___mcc_h16
				_ = _539_cf23
				var _540_cf22 bool = _531___mcc_h15
				_ = _540_cf22
				var _541_v34 _dafny.Array
				_ = _541_v34
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_8
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Sequence = (func(_542_cf23 bool, _543_cf25 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_544_i3 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_542_cf23, true)).Cardinality()), _dafny.SeqOf((_this).F21(), (_this).F21(), _543_cf25))
						}
					})(_539_cf23, _537_cf25)
					_ = _init8
					var _element0_8 = _init8(_dafny.Zero)
					_ = _element0_8
					_nw60 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
					(_nw60).ArraySet1(_element0_8, 0)
					var _nativeLen0_8 = (_len0_8).Int()
					_ = _nativeLen0_8
					for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
						(_nw60).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
					}
				}
				_541_v34 = _nw60
				var _545_v35 _dafny.Sequence
				_ = _545_v35
				_545_v35 = _dafny.SeqOf((_this).F21())
				var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_541_v34), 0))
				_ = _index49
				(_541_v34).ArraySet1(_545_v35, (_index49).Int())
				var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_541_v34), 0))
				_ = _index50
				(_541_v34).ArraySet1(_545_v35, (_index50).Int())
				var _546_v36 _dafny.Map
				_ = _546_v36
				_546_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_537_cf25, !((_this).F17()))
				var _547_v37 _dafny.Sequence
				_ = _547_v37
				_547_v37 = _dafny.SeqOf((_this).F13(), (func() bool {
					if (_546_v36).Contains(_538_cf24) {
						return (_546_v36).Get(_538_cf24).(bool)
					}
					return _507_v19.F19
				})(), _536_cf26, (_this).F22(), _536_cf26)
				(_507_v19).F19 = (((_this).Fm10(_539_cf23, _536_cf26, _547_v37, globalState)).Cardinality()).Cmp(_538_cf24) < 0
				_539_cf23 = ((_this).F21()).Cmp(_537_cf25) == 0
				var _548_v38 _dafny.Array
				_ = _548_v38
				var _nwElement0_14 _dafny.Array = (_this).F14()
				_ = _nwElement0_14
				var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(29))
				_ = _nw61
				(_nw61).ArraySet1(_nwElement0_14, 0)
				(_nw61).ArraySet1((_this).F14(), 1)
				(_nw61).ArraySet1((_this).F14(), 2)
				(_nw61).ArraySet1((_this).F14(), 3)
				(_nw61).ArraySet1((_this).F14(), 4)
				(_nw61).ArraySet1((_this).F14(), 5)
				(_nw61).ArraySet1((_this).F14(), 6)
				(_nw61).ArraySet1((_this).F14(), 7)
				(_nw61).ArraySet1((_this).F14(), 8)
				(_nw61).ArraySet1((_this).F14(), 9)
				(_nw61).ArraySet1((_this).F14(), 10)
				(_nw61).ArraySet1((_this).F14(), 11)
				(_nw61).ArraySet1((Companion_D2_.Create_DC4_(_467_v1, (_this).F22(), _dafny.IntOfUint32((_466_v0).Cardinality()), _507_v19.F19, (_this).F14())).Dtor_cf10(), 12)
				(_nw61).ArraySet1((_this).F14(), 13)
				(_nw61).ArraySet1((_this).F14(), 14)
				(_nw61).ArraySet1((_this).F14(), 15)
				(_nw61).ArraySet1((_this).F14(), 16)
				(_nw61).ArraySet1((_this).F14(), 17)
				(_nw61).ArraySet1((_this).F14(), 18)
				(_nw61).ArraySet1((_this).F14(), 19)
				(_nw61).ArraySet1((_this).F14(), 20)
				(_nw61).ArraySet1((_this).F14(), 21)
				(_nw61).ArraySet1((_this).F14(), 22)
				(_nw61).ArraySet1((_this).F14(), 23)
				(_nw61).ArraySet1((_this).F14(), 24)
				(_nw61).ArraySet1((_this).F14(), 25)
				(_nw61).ArraySet1((_this).F14(), 26)
				(_nw61).ArraySet1((_this).F14(), 27)
				(_nw61).ArraySet1((_this).F14(), 28)
				_548_v38 = _nw61
				var _549_v39 _dafny.Array
				_ = _549_v39
				_549_v39 = _548_v38
				var _550_v40 _dafny.Array
				_ = _550_v40
				var _nwElement0_15 _dafny.Array = _549_v39
				_ = _nwElement0_15
				var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(27))
				_ = _nw62
				(_nw62).ArraySet1(_nwElement0_15, 0)
				(_nw62).ArraySet1(_549_v39, 1)
				(_nw62).ArraySet1(_548_v38, 2)
				(_nw62).ArraySet1(_549_v39, 3)
				(_nw62).ArraySet1(_549_v39, 4)
				(_nw62).ArraySet1(_549_v39, 5)
				(_nw62).ArraySet1(_548_v38, 6)
				(_nw62).ArraySet1(_549_v39, 7)
				(_nw62).ArraySet1(_549_v39, 8)
				(_nw62).ArraySet1(_549_v39, 9)
				(_nw62).ArraySet1(_549_v39, 10)
				(_nw62).ArraySet1(_549_v39, 11)
				(_nw62).ArraySet1(_549_v39, 12)
				(_nw62).ArraySet1(_549_v39, 13)
				(_nw62).ArraySet1(_549_v39, 14)
				(_nw62).ArraySet1(_549_v39, 15)
				(_nw62).ArraySet1(_549_v39, 16)
				(_nw62).ArraySet1(_549_v39, 17)
				(_nw62).ArraySet1(_549_v39, 18)
				(_nw62).ArraySet1(_548_v38, 19)
				(_nw62).ArraySet1(_549_v39, 20)
				(_nw62).ArraySet1(_549_v39, 21)
				(_nw62).ArraySet1(_549_v39, 22)
				(_nw62).ArraySet1(_549_v39, 23)
				(_nw62).ArraySet1(_549_v39, 24)
				(_nw62).ArraySet1(_549_v39, 25)
				(_nw62).ArraySet1(_549_v39, 26)
				_550_v40 = _nw62
				var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_550_v40), 0))
				_ = _index51
				(_550_v40).ArraySet1(_549_v39, (_index51).Int())
				var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(257), _dafny.ArrayLen((_550_v40), 0))
				_ = _index52
				(_550_v40).ArraySet1(_548_v38, (_index52).Int())
			} else if _source11.Is_DC11() {
				var _551___mcc_h20 _dafny.MultiSet = _source11.Get_().(D6_DC11).Cf17
				_ = _551___mcc_h20
				var _552_cf17 _dafny.MultiSet = _551___mcc_h20
				_ = _552_cf17
				var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index53
				((_this).F14()).ArraySet1((_this).F22(), (_index53).Int())
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index54
				((_this).F14()).ArraySet1(_507_v19.F19, (_index54).Int())
				(globalState).F7 = (_this).F21()
				var _553_v41 _dafny.MultiSet
				_ = _553_v41
				_553_v41 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_466_v0, _466_v0), _466_v0, _466_v0, _466_v0, _466_v0)
				var _554_v42 _dafny.Map
				_ = _554_v42
				_554_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(658)).Cmp((_this).F21()) <= 0, _467_v1)
				var _555_v43 _dafny.Set
				_ = _555_v43
				_555_v43 = _dafny.SetOf((_this).F17())
				var _556_v44 _dafny.Sequence
				_ = _556_v44
				_556_v44 = _dafny.SeqOf(_dafny.SetOf(false), _555_v43)
				var _557_v45 _dafny.MultiSet
				_ = _557_v45
				_557_v45 = _dafny.MultiSetOf(((_556_v44).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F21())).Cardinality()), _dafny.IntOfUint32((_556_v44).Cardinality()))).Uint32()).(_dafny.Set)).IsDisjointFrom(_555_v43), _507_v19.F19, ((_this).F15()).IsSubsetOf((_this).F15()), _507_v19.F19, Companion_Default___.Fm30((_this).F21(), globalState))
				var _558_v46 _dafny.Sequence
				_ = _558_v46
				_558_v46 = _dafny.SeqOf(_557_v45, _557_v45, (_557_v45).Difference(_557_v45))
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index55
				var _rhs44 bool = false
				_ = _rhs44
				var _rhs45 _dafny.MultiSet = ((func() _dafny.MultiSet {
					if _507_v19.F19 {
						return _553_v41
					}
					return (_553_v41).Update(_466_v0, Companion_Default___.Abs((_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F21()))).Cardinality()))
				})()).Update((_this.F16()).Dtor_cf1(), Companion_Default___.Abs((_this).F21()))
				_ = _rhs45
				var _rhs46 _dafny.Map = _554_v42
				_ = _rhs46
				var _rhs47 _dafny.MultiSet = (_558_v46).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_558_v46).Cardinality()))).Uint32()).(_dafny.MultiSet)
				_ = _rhs47
				var _rhs48 bool = !(!(func() _dafny.Map {
					var _coll33 = _dafny.NewMapBuilder()
					_ = _coll33
					for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(391), _dafny.IntOfInt64(894))); ; {
						_compr_33, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _559_v47 _dafny.Int
						_559_v47 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(391)).Cmp(_559_v47) <= 0) && ((_559_v47).Cmp(_dafny.IntOfInt64(894)) < 0) {
							_coll33.Add((_559_v47).Times((_this).F21()), _557_v45)
						}
					}
					return _coll33.ToMap()
				}()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _557_v45))) || ((_this).F13())
				_ = _rhs48
				var _lhs37 _dafny.Array = (_this).F14()
				_ = _lhs37
				var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs38
				var _lhs39 *C0 = _507_v19
				_ = _lhs39
				(_lhs37).ArraySet1(_rhs44, (_lhs38).Int())
				_553_v41 = _rhs45
				_554_v42 = _rhs46
				_557_v45 = _rhs47
				_lhs39.F19 = _rhs48
				var _560_v48 _dafny.Sequence
				_ = _560_v48
				_560_v48 = _dafny.SeqOf((_this).F17(), (_this).F13())
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(312), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index56
				((_this).F14()).ArraySet1(((Companion_Default___.Fm36((_560_v48).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(117), _dafny.IntOfUint32((_560_v48).Cardinality()))).Uint32()).(bool), (_507_v19).F20(), Companion_Default___.Fm30((_this).F21(), globalState), globalState)).Equals(_555_v43)) && (true), (_index56).Int())
			} else {
				var _561___mcc_h21 D6 = _source11.Get_().(D6_DC14).Cf27
				_ = _561___mcc_h21
				var _562_cf27 D6 = _561___mcc_h21
				_ = _562_cf27
				var _563_v49 _dafny.Array
				_ = _563_v49
				var _nw63 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(20))
				_ = _nw63
				_563_v49 = _nw63
				var _rhs49 _dafny.Array = _563_v49
				_ = _rhs49
				var _rhs50 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("mxwbndl")
				_ = _rhs50
				var _rhs51 bool = _507_v19.F19
				_ = _rhs51
				var _lhs40 *C0 = _507_v19
				_ = _lhs40
				_563_v49 = _rhs49
				_466_v0 = _rhs50
				_lhs40.F19 = _rhs51
				var _564_v50 _dafny.Sequence
				_ = _564_v50
				_564_v50 = _dafny.SeqOf((_this).F21())
				var _565_v51 _dafny.Map
				_ = _565_v51
				_565_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_466_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_564_v50, _564_v50)).Cardinality()))
				_565_v51 = (_565_v51).Update(_466_v0, (_this).F21())
				(_507_v19).F19 = _507_v19.F19
				var _566_v52 T0
				_ = _566_v52
				var _nw64 *C0 = New_C0_()
				_ = _nw64
				_nw64.Ctor__((_this).F22(), (_507_v19).F20(), Companion_Default___.Fm37((_this).F21(), (_this).F21(), (_this).F21(), (_507_v19).Fm16((_this).F22(), (_this).F21(), globalState), globalState), _507_v19.F19)
				_566_v52 = _nw64
				var _567_v53 _dafny.MultiSet
				_ = _567_v53
				_567_v53 = _dafny.MultiSetOf(_566_v52, _566_v52, _566_v52)
				_567_v53 = (_567_v53).Intersection(_567_v53)
			}
			var _568_v54 _dafny.Array
			_ = _568_v54
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_9
			var _nw65 _dafny.Array
			_ = _nw65
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw65 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.CodePoint = (func(_569_v18 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_570_i4 _dafny.Int) _dafny.CodePoint {
						return _569_v18
					}
				})(_506_v18)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw65 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw65).ArraySet1CodePoint(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw65).ArraySet1CodePoint(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_568_v54 = _nw65
			var _571_v55 _dafny.MultiSet
			_ = _571_v55
			_571_v55 = _dafny.MultiSetOf(_568_v54)
			var _572_v56 _dafny.Map
			_ = _572_v56
			_572_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _571_v55)
			var _573_v57 _dafny.MultiSet
			_ = _573_v57
			_573_v57 = _dafny.MultiSetOf((_this).F13(), !((_this).F13()))
			if !(_573_v57).Contains(((func() _dafny.MultiSet {
				if (_572_v56).Contains((_this).F14()) {
					return (_572_v56).Get((_this).F14()).(_dafny.MultiSet)
				}
				return _571_v55
			})()).Contains(_568_v54)) {
				var _574_v58 *C0
				_ = _574_v58
				var _nw66 *C0 = New_C0_()
				_ = _nw66
				_nw66.Ctor__(!(_507_v19.F19) || ((_this).F13()), (_507_v19).F20(), Companion_D1_.Create_DC2_((_this).F21(), _dafny.IntOfInt64(609), Companion_Default___.Fm30((_this).F21(), globalState)), _507_v19.F19)
				_574_v58 = _nw66
				_506_v18 = (_507_v19).F20()
				_466_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("avvykq"), _dafny.Companion_Sequence_.Concatenate(_466_v0, _466_v0))
				var _575_v59 _dafny.Sequence
				_ = _575_v59
				_575_v59 = _dafny.SeqOf((_this).F21())
				var _576_v60 _dafny.Map
				_ = _576_v60
				_576_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_466_v0, _575_v59)
				var _577_v62 _dafny.Map
				_ = _577_v62
				_577_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter39 := _dafny.Iterate((_576_v60).Keys().Elements()); ; {
						_compr_34, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _578_v61 _dafny.Sequence
						_578_v61 = interface{}(_compr_34).(_dafny.Sequence)
						if (_576_v60).Contains(_578_v61) {
							_coll34.Add(_578_v61)
						}
					}
					return _coll34.ToSet()
				}())
				var _579_v63 _dafny.Set
				_ = _579_v63
				_579_v63 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("mbugy"), _466_v0, _466_v0, _dafny.UnicodeSeqOfUtf8Bytes("fw"), _466_v0)
				_577_v62 = (_577_v62).Update((_this).F21(), _579_v63)
				var _rhs52 _dafny.Int = (_this).F21()
				_ = _rhs52
				var _rhs53 _dafny.Int = ((_this).F15()).Cardinality()
				_ = _rhs53
				var _lhs41 *GlobalState = globalState
				_ = _lhs41
				var _lhs42 *GlobalState = globalState
				_ = _lhs42
				_lhs41.F7 = _rhs52
				_lhs42.F7 = _rhs53
			} else {
				var _580_v64 _dafny.Map
				_ = _580_v64
				_580_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F16(), (_this).F21())
				var _581_v65 _dafny.Map
				_ = _581_v65
				_581_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_507_v19.F19, _580_v64)
				_581_v65 = _581_v65
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))
				_ = _index57
				(_467_v1).ArraySet1((_this).F21(), (_index57).Int())
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_467_v1), 0))
				_ = _index58
				(_467_v1).ArraySet1((_this).F21(), (_index58).Int())
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))
				_ = _index59
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_467_v1), 0))
				_ = _index60
				var _rhs54 _dafny.Int = (_this).F21()
				_ = _rhs54
				var _rhs55 _dafny.Int = (((_this).F21()).Minus((_this).F21())).Plus((_dafny.Zero).Minus((_this).F21()))
				_ = _rhs55
				var _lhs43 _dafny.Array = _467_v1
				_ = _lhs43
				var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))
				_ = _lhs44
				var _lhs45 _dafny.Array = _467_v1
				_ = _lhs45
				var _lhs46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(577), _dafny.ArrayLen((_467_v1), 0))
				_ = _lhs46
				(_lhs43).ArraySet1(_rhs54, (_lhs44).Int())
				(_lhs45).ArraySet1(_rhs55, (_lhs46).Int())
				var _582_v66 _dafny.Sequence
				_ = _582_v66
				_582_v66 = _dafny.SeqOf((_this).F17())
				var _583_v67 _dafny.Set
				_ = _583_v67
				_583_v67 = _dafny.SetOf((_582_v66).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_582_v66).Cardinality()))).Uint32()).(bool), true)
				_583_v67 = (_dafny.SetOf(_507_v19.F19)).Intersection(_583_v67)
				var _584_v68 _dafny.Map
				_ = _584_v68
				_584_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F15())
				var _585_v69 _dafny.Array
				_ = _585_v69
				var _nwElement0_16 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_466_v0, (Companion_Default___.SafeIndex((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_466_v0).Cardinality()))).Uint32(), (_507_v19).F20())).Cardinality())
				_ = _nwElement0_16
				var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(24))
				_ = _nw67
				(_nw67).ArraySet1(_nwElement0_16, 0)
				(_nw67).ArraySet1((_this).F21(), 1)
				(_nw67).ArraySet1((_this).F21(), 2)
				(_nw67).ArraySet1((_504_v16).Cardinality(), 3)
				(_nw67).ArraySet1((_dafny.Zero).Minus((_this).F21()), 4)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 5)
				(_nw67).ArraySet1(_dafny.IntOfInt64(256), 6)
				(_nw67).ArraySet1(_dafny.IntOfInt64(565), 7)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 8)
				(_nw67).ArraySet1(_dafny.IntOfInt64(780), 9)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 10)
				(_nw67).ArraySet1((_this).F21(), 11)
				(_nw67).ArraySet1(Companion_Default___.Fm0(_582_v66, _507_v19.F19, (_this).F22(), (_this).Fm8(Companion_Default___.Fm30(_dafny.IntOfUint32((_582_v66).Cardinality()), globalState), _this.F16(), globalState), globalState), 12)
				(_nw67).ArraySet1((_584_v68).Cardinality(), 13)
				(_nw67).ArraySet1((func() _dafny.Int {
					if (_504_v16).Contains(false) {
						return (_504_v16).Get(false).(_dafny.Int)
					}
					return (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int)
				})(), 14)
				(_nw67).ArraySet1((_this).F21(), 15)
				(_nw67).ArraySet1((_this).F21(), 16)
				(_nw67).ArraySet1((_this).F21(), 17)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 18)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 19)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 20)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 21)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 22)
				(_nw67).ArraySet1((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), 23)
				_585_v69 = _nw67
				var _586_v70 D2
				_ = _586_v70
				_586_v70 = Companion_D2_.Create_DC4_(_585_v69, ((_this).F21()).Cmp((_583_v67).Cardinality()) == 0, (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), Companion_Default___.Fm30((_this).F21(), globalState), (_this).F14())
				var _587_v71 _dafny.Sequence
				_ = _587_v71
				_587_v71 = _dafny.SeqOf((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int))
				var _588_v72 _dafny.Map
				_ = _588_v72
				_588_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("txsjiq"), (_this).F22())
				var _589_v73 _dafny.Sequence
				_ = _589_v73
				_589_v73 = _dafny.SeqOf(_588_v72)
				var _590_v74 _dafny.Sequence
				_ = _590_v74
				_590_v74 = _dafny.SeqOf(_588_v72, (_589_v73).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_589_v73).Cardinality()))).Uint32()).(_dafny.Map), _588_v72)
				var _591_v75 _dafny.Map
				_ = _591_v75
				_591_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_585_v69, (_503_cf17).IsDisjointFrom((_dafny.MultiSetFromSeq(_587_v71)).Update((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), Companion_Default___.Abs(((_590_v74).Select((Companion_Default___.SafeIndex((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_590_v74).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()))))
				var _rhs56 D2 = Companion_D2_.Create_DC4_(_585_v69, (_this).F22(), (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), (func() bool {
					if (_this).F22() {
						return (_582_v66).Select((Companion_Default___.SafeIndex((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(242), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_582_v66).Cardinality()))).Uint32()).(bool)
					}
					return _507_v19.F19
				})(), (_this).F14())
				_ = _rhs56
				var _rhs57 _dafny.Map = (func() _dafny.Map {
					if _dafny.Companion_Sequence_.Equal(Companion_Default___.Fm38(false, _507_v19.F19, _507_v19.F19, (_this).F22(), globalState), _582_v66) {
						return _591_v75
					}
					return (_591_v75).Merge(_591_v75)
				})()
				_ = _rhs57
				_586_v70 = _rhs56
				_591_v75 = _rhs57
				(globalState).F2 = Companion_Default___.Fm38((_this).F22(), (_this).F17(), false, (_this).F22(), globalState)
			}
		} else {
			var _592___mcc_h10 D6 = _source10.Get_().(D6_DC14).Cf27
			_ = _592___mcc_h10
			var _593_cf27 D6 = _592___mcc_h10
			_ = _593_cf27
			var _594_v76 _dafny.Sequence
			_ = _594_v76
			_594_v76 = _dafny.SeqOf(_this.F12())
			_594_v76 = _594_v76
			var _595_v77 _dafny.Array
			_ = _595_v77
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
			_ = _nw68
			_595_v77 = _nw68
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))
			_ = _index61
			(_595_v77).ArraySet1(_467_v1, (_index61).Int())
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))
			_ = _index62
			(_595_v77).ArraySet1(_467_v1, (_index62).Int())
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index63
			((_this).F14()).ArraySet1((_this).F22(), (_index63).Int())
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index64
			((_this).F14()).ArraySet1((_this).F17(), (_index64).Int())
			if !(!((_this).F22())) {
				var _596_v78 _dafny.Array
				_ = _596_v78
				var _nw69 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
				_ = _nw69
				_596_v78 = _nw69
				var _597_v79 _dafny.CodePoint
				_ = _597_v79
				_597_v79 = _dafny.CodePoint('b')
				var _598_v80 *C0
				_ = _598_v80
				var _nw70 *C0 = New_C0_()
				_ = _nw70
				_nw70.Ctor__(false, _597_v79, _this.F12(), (_this).F22())
				_598_v80 = _nw70
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_596_v78), 0))
				_ = _index65
				(_596_v78).ArraySet1(_598_v80, (_index65).Int())
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index66
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_596_v78), 0))
				_ = _index67
				var _rhs58 bool = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				_ = _rhs58
				var _rhs59 *C0 = _598_v80
				_ = _rhs59
				var _lhs47 _dafny.Array = (_this).F14()
				_ = _lhs47
				var _lhs48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs48
				var _lhs49 _dafny.Array = _596_v78
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_596_v78), 0))
				_ = _lhs50
				(_lhs47).ArraySet1(_rhs58, (_lhs48).Int())
				(_lhs49).ArraySet1(_rhs59, (_lhs50).Int())
				(_598_v80).F19 = _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_599_v80 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_600_i5 _dafny.Int) _dafny.CodePoint {
						return (_599_v80).F20()
					}
				})(_598_v80))), (_598_v80).F20())
				var _601_v81 _dafny.Set
				_ = _601_v81
				_601_v81 = _dafny.SetOf(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_this).F13())
				(_598_v80).F19 = ((_this).F21()).Cmp((func() _dafny.Int {
					if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
						return (_601_v81).Cardinality()
					}
					return _dafny.IntOfUint32((_466_v0).Cardinality())
				})()) > 0
				var _602_v82 _dafny.Sequence
				_ = _602_v82
				_602_v82 = _dafny.SeqOf(_dafny.IntOfUint32((_466_v0).Cardinality()))
				var _603_v83 _dafny.Map
				_ = _603_v83
				_603_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_602_v82, _598_v80.F19)
				var _604_v84 _dafny.Sequence
				_ = _604_v84
				_604_v84 = _dafny.SeqOf(_602_v82)
				_603_v83 = (_603_v83).Update((_604_v84).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_604_v84).Cardinality()))).Uint32()).(_dafny.Sequence), (_this).F22())
				var _605_v85 _dafny.Array
				_ = _605_v85
				var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(10))
				_ = _nw71
				_605_v85 = _nw71
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_605_v85), 0))
				_ = _index68
				(_605_v85).ArraySet1CodePoint((_598_v80).F20(), (_index68).Int())
				var _606_v86 _dafny.Set
				_ = _606_v86
				_606_v86 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_607_v79 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_608_i6 _dafny.Int) _dafny.CodePoint {
						return _607_v79
					}
				})(_597_v79))))
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_605_v85), 0))
				_ = _index69
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index70
				var _rhs60 _dafny.Int = (_this).Fm9(_dafny.CodePoint('y'), globalState)
				_ = _rhs60
				var _rhs61 _dafny.CodePoint = _597_v79
				_ = _rhs61
				var _rhs62 bool = false
				_ = _rhs62
				var _rhs63 bool = (_606_v86).IsSubsetOf(_dafny.SetOf(_466_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(726))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}(func(_609_i7 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('h')
				}))))
				_ = _rhs63
				var _rhs64 bool = !((_this).F22())
				_ = _rhs64
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 _dafny.Array = _605_v85
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(790), _dafny.ArrayLen((_605_v85), 0))
				_ = _lhs53
				var _lhs54 _dafny.Array = (_this).F14()
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs55
				var _lhs56 *C0 = _598_v80
				_ = _lhs56
				var _lhs57 *C0 = _598_v80
				_ = _lhs57
				_lhs51.F7 = _rhs60
				(_lhs52).ArraySet1CodePoint(_rhs61, (_lhs53).Int())
				(_lhs54).ArraySet1(_rhs62, (_lhs55).Int())
				_lhs56.F19 = _rhs63
				_lhs57.F19 = _rhs64
			} else {
				var _610_v87 _dafny.Sequence
				_ = _610_v87
				_610_v87 = _dafny.SeqOf((_this).F21())
				var _611_v88 D7
				_ = _611_v88
				_611_v88 = Companion_D7_.Create_DC15_(_610_v87)
				var _pat_let_tv15 = globalState
				_ = _pat_let_tv15
				_611_v88 = (func() D7 {
					if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
						return func(_pat_let12_0 D7) D7 {
							return func(_612_dt__update__tmp_h3 D7) D7 {
								return func(_pat_let13_0 _dafny.Sequence) D7 {
									return func(_613_dt__update_hcf28_h0 _dafny.Sequence) D7 {
										return Companion_D7_.Create_DC15_(_613_dt__update_hcf28_h0)
									}(_pat_let13_0)
								}(Companion_Default___.Fm34((_this).F21(), _pat_let_tv15))
							}(_pat_let12_0)
						}(_611_v88)
					}
					return _611_v88
				})()
				var _614_v89 _dafny.Map
				_ = _614_v89
				_614_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F13())
				var _615_v90 D6
				_ = _615_v90
				_615_v90 = Companion_D6_.Create_DC13_((_this).F17(), (_this).F22(), _dafny.IntOfUint32((_466_v0).Cardinality()), (_this).F21(), (func() bool {
					if (_614_v89).Contains((_this).F13()) {
						return (_614_v89).Get((_this).F13()).(bool)
					}
					return (_this).F17()
				})())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index71
				((_this).F14()).ArraySet1((_615_v90).Dtor_cf26(), (_index71).Int())
				var _616_v91 _dafny.Array
				_ = _616_v91
				var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
				_ = _nw72
				_616_v91 = _nw72
				var _617_v92 _dafny.Sequence
				_ = _617_v92
				_617_v92 = _dafny.SeqOf(true, (_this).F17(), true, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_616_v91), 0))
				_ = _index72
				(_616_v91).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_617_v92, _617_v92), (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_616_v91), 0))
				_ = _index73
				var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index74
				var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index75
				var _rhs65 _dafny.Sequence = _617_v92
				_ = _rhs65
				var _rhs66 bool = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				_ = _rhs66
				var _rhs67 _dafny.Sequence = (func() _dafny.Sequence {
					if Companion_Default___.Fm30((_this).F21(), globalState) {
						return _466_v0
					}
					return _466_v0
				})()
				_ = _rhs67
				var _rhs68 bool = !(!((_this).F13()))
				_ = _rhs68
				var _lhs58 _dafny.Array = _616_v91
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_616_v91), 0))
				_ = _lhs59
				var _lhs60 _dafny.Array = (_this).F14()
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs61
				var _lhs62 _dafny.Array = (_this).F14()
				_ = _lhs62
				var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs63
				(_lhs58).ArraySet1(_rhs65, (_lhs59).Int())
				(_lhs60).ArraySet1(_rhs66, (_lhs61).Int())
				_466_v0 = _rhs67
				(_lhs62).ArraySet1(_rhs68, (_lhs63).Int())
				var _618_v93 _dafny.Map
				_ = _618_v93
				_618_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_this).F21())
				var _arr6 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _arr6
				var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _index76
				(_arr6).ArraySet1((_618_v93).Cardinality(), (_index76).Int())
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_467_v1), 0))
				_ = _index77
				(_467_v1).ArraySet1((_610_v87).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(226), _dafny.IntOfUint32((_610_v87).Cardinality()))).Uint32()).(_dafny.Int), (_index77).Int())
				var _arr7 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _arr7
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _index78
				(_arr7).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F17())).Cardinality(), (_index78).Int())
				var _619_v94 _dafny.CodePoint
				_ = _619_v94
				_619_v94 = _dafny.CodePoint('x')
				var _620_v95 D6
				_ = _620_v95
				_620_v95 = Companion_D6_.Create_DC14_(Companion_Default___.Fm39(false, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), globalState))
				var _621_v96 _dafny.Array
				_ = _621_v96
				var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
				_ = _nw73
				_621_v96 = _nw73
				var _622_v97 _dafny.Sequence
				_ = _622_v97
				_622_v97 = _dafny.SeqOf(_621_v96, _621_v96, _621_v96)
				var _arr8 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _arr8
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _index79
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index80
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_467_v1), 0))
				_ = _index81
				var _arr9 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _arr9
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _index82
				var _rhs69 _dafny.Int = (_this).F21()
				_ = _rhs69
				var _rhs70 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_466_v0, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_466_v0).Cardinality()))).Uint32(), _619_v94), _466_v0, Companion_Default___.Fm32(_620_v95, globalState)), _466_v0)
				_ = _rhs70
				var _rhs71 _dafny.Int = (_dafny.IntOfUint32((_610_v87).Cardinality())).Plus(((_this).F21()).Plus((_this).F21()))
				_ = _rhs71
				var _rhs72 _dafny.Int = _dafny.IntOfUint32((_622_v97).Cardinality())
				_ = _rhs72
				var _lhs64 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _lhs64
				var _lhs65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _lhs65
				var _lhs66 _dafny.Array = (_this).F14()
				_ = _lhs66
				var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs67
				var _lhs68 _dafny.Array = _467_v1
				_ = _lhs68
				var _lhs69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_467_v1), 0))
				_ = _lhs69
				var _lhs70 _dafny.Array = _dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))
				_ = _lhs70
				var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(109), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))
				_ = _lhs71
				(_lhs64).ArraySet1(_rhs69, (_lhs65).Int())
				(_lhs66).ArraySet1(_rhs70, (_lhs67).Int())
				(_lhs68).ArraySet1(_rhs71, (_lhs69).Int())
				(_lhs70).ArraySet1(_rhs72, (_lhs71).Int())
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.ArrayLen((_467_v1), 0))
				_ = _index83
				(_467_v1).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-317), (_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen((_dafny.ArrayCastTo((_595_v77).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(98), _dafny.ArrayLen((_595_v77), 0))).Int()))), 0))).Int()).(_dafny.Int)), (_index83).Int())
			}
		}
		var _623_v98 bool
		_ = _623_v98
		_623_v98 = false
		var _624_v99 _dafny.Sequence
		_ = _624_v99
		_624_v99 = _dafny.SeqOf(false, (_this).F13(), _623_v98, _623_v98, (_this).F13())
		var _625_v100 _dafny.Map
		_ = _625_v100
		_625_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _624_v99)
		var _626_v101 _dafny.Map
		_ = _626_v101
		_626_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_624_v99).Cardinality()), (_this).F22())
		_623_v98 = _dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_625_v100).Contains((_626_v101).Cardinality()) {
				return (_625_v100).Get((_626_v101).Cardinality()).(_dafny.Sequence)
			}
			return _624_v99
		})(), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_625_v100).Contains((_626_v101).Cardinality()) {
				return (_625_v100).Get((_626_v101).Cardinality()).(_dafny.Sequence)
			}
			return _624_v99
		})()).Cardinality()))).Uint32(), (_this).F22()), _dafny.Companion_Sequence_.Concatenate(_624_v99, _624_v99))
		var _627_v102 _dafny.MultiSet
		_ = _627_v102
		_627_v102 = _dafny.MultiSetOf((_this).F17())
		var _628_v103 _dafny.MultiSet
		_ = _628_v103
		_628_v103 = _627_v102
		var _629_v104 _dafny.Array
		_ = _629_v104
		var _nwElement0_17 _dafny.MultiSet = _627_v102
		_ = _nwElement0_17
		var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(16))
		_ = _nw74
		(_nw74).ArraySet1(_nwElement0_17, 0)
		(_nw74).ArraySet1(_627_v102, 1)
		(_nw74).ArraySet1((_627_v102).Difference((_dafny.MultiSetOf((_this).F22(), true)).Update(true, Companion_Default___.Abs(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(124))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}(func(_630_i9 _dafny.Int) _dafny.Int {
			return (_this).F21()
		}))).Cardinality())))), 2)
		(_nw74).ArraySet1(_627_v102, 3)
		(_nw74).ArraySet1((_627_v102).Intersection(_627_v102), 4)
		(_nw74).ArraySet1((_628_v103), 5)
		(_nw74).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F13())), 6)
		(_nw74).ArraySet1(_627_v102, 7)
		(_nw74).ArraySet1(_dafny.MultiSetOf((_this).F22()), 8)
		(_nw74).ArraySet1((Companion_Default___.Fm40(globalState)).Intersection(_627_v102), 9)
		(_nw74).ArraySet1(_627_v102, 10)
		(_nw74).ArraySet1((_627_v102).Difference(_627_v102), 11)
		(_nw74).ArraySet1(_627_v102, 12)
		(_nw74).ArraySet1(_dafny.MultiSetFromSeq(_624_v99), 13)
		(_nw74).ArraySet1((_dafny.MultiSetOf((_this).F22())).Union(_627_v102), 14)
		(_nw74).ArraySet1((_627_v102).Difference(_dafny.MultiSetOf((_this).F13())), 15)
		_629_v104 = _nw74
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_629_v104), 0))); ; {
			_guard_loop_3, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _631_i8 _dafny.Int
			_631_i8 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_631_i8).Sign() != -1) && ((_631_i8).Cmp(_dafny.ArrayLen((_629_v104), 0)) < 0)) {
				(_629_v104).ArraySet1((_627_v102).Union(_627_v102), (_631_i8).Int())
			}
		}
		var _632_v105 _dafny.Array
		_ = _632_v105
		var _len0_10 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_10
		var _nw75 _dafny.Array
		_ = _nw75
		if _len0_10.Cmp(_dafny.Zero) == 0 {
			_nw75 = _dafny.NewArray(_len0_10)
		} else {
			var _init10 func(_dafny.Int) _dafny.Sequence = func(_633_i10 _dafny.Int) _dafny.Sequence {
				return (_this.F16()).Dtor_cf1()
			}
			_ = _init10
			var _element0_10 = _init10(_dafny.Zero)
			_ = _element0_10
			_nw75 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
			(_nw75).ArraySet1(_element0_10, 0)
			var _nativeLen0_10 = (_len0_10).Int()
			_ = _nativeLen0_10
			for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
				(_nw75).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
			}
		}
		_632_v105 = _nw75
		var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_632_v105), 0))
		_ = _index84
		(_632_v105).ArraySet1((func() _dafny.Sequence {
			if _623_v98 {
				return _466_v0
			}
			return _466_v0
		})(), (_index84).Int())
		var _634_v106 D6
		_ = _634_v106
		_634_v106 = Companion_D6_.Create_DC14_(Companion_D6_.Create_DC13_((_this).F17(), (_this).F13(), (_this).F21(), (_this).F21(), _623_v98))
		var _635_v107 D6
		_ = _635_v107
		_635_v107 = Companion_D6_.Create_DC14_(_634_v106)
		var _636_v108 D6
		_ = _636_v108
		_636_v108 = Companion_D6_.Create_DC14_(_634_v106)
		var _pat_let_tv16 = _466_v0
		_ = _pat_let_tv16
		var _pat_let_tv17 = _466_v0
		_ = _pat_let_tv17
		var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_632_v105), 0))
		_ = _index85
		(_632_v105).ArraySet1(func(_source12 D6) _dafny.Sequence {
			if _source12.Is_DC12() {
				var _637___mcc_h22 bool = _source12.Get_().(D6_DC12).Cf18
				_ = _637___mcc_h22
				var _638___mcc_h23 _dafny.Sequence = _source12.Get_().(D6_DC12).Cf19
				_ = _638___mcc_h23
				var _639___mcc_h24 _dafny.Int = _source12.Get_().(D6_DC12).Cf20
				_ = _639___mcc_h24
				var _640___mcc_h25 bool = _source12.Get_().(D6_DC12).Cf21
				_ = _640___mcc_h25
				var _641_cf21 bool = _640___mcc_h25
				_ = _641_cf21
				var _642_cf20 _dafny.Int = _639___mcc_h24
				_ = _642_cf20
				var _643_cf19 _dafny.Sequence = _638___mcc_h23
				_ = _643_cf19
				var _644_cf18 bool = _637___mcc_h22
				_ = _644_cf18
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(86))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_645_i11 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}))
			} else if _source12.Is_DC13() {
				var _646___mcc_h26 bool = _source12.Get_().(D6_DC13).Cf22
				_ = _646___mcc_h26
				var _647___mcc_h27 bool = _source12.Get_().(D6_DC13).Cf23
				_ = _647___mcc_h27
				var _648___mcc_h28 _dafny.Int = _source12.Get_().(D6_DC13).Cf24
				_ = _648___mcc_h28
				var _649___mcc_h29 _dafny.Int = _source12.Get_().(D6_DC13).Cf25
				_ = _649___mcc_h29
				var _650___mcc_h30 bool = _source12.Get_().(D6_DC13).Cf26
				_ = _650___mcc_h30
				var _651_cf26 bool = _650___mcc_h30
				_ = _651_cf26
				var _652_cf25 _dafny.Int = _649___mcc_h29
				_ = _652_cf25
				var _653_cf24 _dafny.Int = _648___mcc_h28
				_ = _653_cf24
				var _654_cf23 bool = _647___mcc_h27
				_ = _654_cf23
				var _655_cf22 bool = _646___mcc_h26
				_ = _655_cf22
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-521))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}(func(_656_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('u')
				}))
			} else if _source12.Is_DC11() {
				var _657___mcc_h31 _dafny.MultiSet = _source12.Get_().(D6_DC11).Cf17
				_ = _657___mcc_h31
				var _658_cf17 _dafny.MultiSet = _657___mcc_h31
				_ = _658_cf17
				return _pat_let_tv16
			} else {
				var _659___mcc_h32 D6 = _source12.Get_().(D6_DC14).Cf27
				_ = _659___mcc_h32
				var _660_cf27 D6 = _659___mcc_h32
				_ = _660_cf27
				return _pat_let_tv17
			}
		}(_636_v108), (_index85).Int())
		var _661_v109 _dafny.Array
		_ = _661_v109
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw76
		_661_v109 = _nw76
		var _662_v110 _dafny.CodePoint
		_ = _662_v110
		_662_v110 = _dafny.CodePoint('u')
		var _663_v111 _dafny.Map
		_ = _663_v111
		_663_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F22(), (_this).F15())
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_661_v109), 0))
		_ = _index86
		(_661_v109).ArraySet1((Companion_Default___.Fm41((_this).F21(), _662_v110, globalState)).Merge(_663_v111), (_index86).Int())
		var _664_v112 _dafny.Sequence
		_ = _664_v112
		_664_v112 = _dafny.SeqOf(_466_v0, _466_v0)
		var _665_v113 _dafny.Set
		_ = _665_v113
		_665_v113 = _dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}((func(_666_v110 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_667_i13 _dafny.Int) _dafny.CodePoint {
				return _666_v110
			}
		})(_662_v110))), (_664_v112).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_664_v112).Cardinality()))).Uint32()).(_dafny.Sequence))
		var _668_v114 _dafny.Map
		_ = _668_v114
		_668_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), true)
		var _669_v115 D3
		_ = _669_v115
		_669_v115 = Companion_D3_.Create_DC5_(Companion_Default___.Fm33((_this).F22(), (_668_v114).Cardinality(), globalState))
		var _pat_let_tv18 = _663_v111
		_ = _pat_let_tv18
		var _pat_let_tv19 = _663_v111
		_ = _pat_let_tv19
		var _pat_let_tv20 = _663_v111
		_ = _pat_let_tv20
		var _pat_let_tv21 = _663_v111
		_ = _pat_let_tv21
		var _pat_let_tv22 = _663_v111
		_ = _pat_let_tv22
		var _pat_let_tv23 = _663_v111
		_ = _pat_let_tv23
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_661_v109), 0))
		_ = _index87
		var _rhs73 bool = func(_source13 _dafny.Set) bool {
			var _670___mcc_h33 _dafny.Set = _source13
			_ = _670___mcc_h33
			var _671_cf16 _dafny.Set = _670___mcc_h33
			_ = _671_cf16
			return (_this).F22()
		}(_665_v113)
		_ = _rhs73
		var _rhs74 _dafny.Map = func(_source14 D3) _dafny.Map {
			if _source14.Is_DC6() {
				var _672___mcc_h34 _dafny.Int = _source14.Get_().(D3_DC6).Cf12
				_ = _672___mcc_h34
				var _673___mcc_h35 bool = _source14.Get_().(D3_DC6).Cf13
				_ = _673___mcc_h35
				var _674_cf13 bool = _673___mcc_h35
				_ = _674_cf13
				var _675_cf12 _dafny.Int = _672___mcc_h34
				_ = _675_cf12
				return (_pat_let_tv18).Merge(_pat_let_tv19)
			} else if _source14.Is_DC7() {
				return _pat_let_tv20
			} else if _source14.Is_DC8() {
				var _676___mcc_h36 _dafny.Int = _source14.Get_().(D3_DC8).Cf14
				_ = _676___mcc_h36
				var _677_cf14 _dafny.Int = _676___mcc_h36
				_ = _677_cf14
				return _pat_let_tv21
			} else {
				var _678___mcc_h37 _dafny.CodePoint = _source14.Get_().(D3_DC5).Cf11
				_ = _678___mcc_h37
				var _679_cf11 _dafny.CodePoint = _678___mcc_h37
				_ = _679_cf11
				if (_this).F17() {
					return _pat_let_tv22
				} else {
					return _pat_let_tv23
				}
			}
		}(_669_v115)
		_ = _rhs74
		var _rhs75 _dafny.Int = (_this).F21()
		_ = _rhs75
		var _rhs76 _dafny.Int = (_this).F21()
		_ = _rhs76
		var _lhs72 _dafny.Array = _661_v109
		_ = _lhs72
		var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_661_v109), 0))
		_ = _lhs73
		var _lhs74 *GlobalState = globalState
		_ = _lhs74
		var _lhs75 *GlobalState = globalState
		_ = _lhs75
		_623_v98 = _rhs73
		(_lhs72).ArraySet1(_rhs74, (_lhs73).Int())
		_lhs74.F7 = _rhs75
		_lhs75.F7 = _rhs76
		var _680_i14 _dafny.Int
		_ = _680_i14
		_680_i14 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_680_i14).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_680_i14 = (_680_i14).Plus(_dafny.One)
					var _681_v116 _dafny.Map
					_ = _681_v116
					_681_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
						if (_this).F13() {
							return (_this).F22()
						}
						return !(!((func() bool {
							if (_626_v101).Contains((_this).F21()) {
								return (_626_v101).Get((_this).F21()).(bool)
							}
							return _623_v98
						})()))
					})(), _dafny.UnicodeSeqOfUtf8Bytes("gr"))
					_681_v116 = (_681_v116).Update((_624_v99).Select((Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_624_v99).Cardinality()))).Uint32()).(bool), _dafny.Companion_Sequence_.Concatenate((_632_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_632_v105), 0))).Int()).(_dafny.Sequence), (_632_v105).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(719), _dafny.ArrayLen((_632_v105), 0))).Int()).(_dafny.Sequence)))
					var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))
					_ = _index88
					(_467_v1).ArraySet1((_this).F21(), (_index88).Int())
					var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))
					_ = _index89
					(_467_v1).ArraySet1((_this).F21(), (_index89).Int())
					_623_v98 = ((func() _dafny.Int {
						if _623_v98 {
							return (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int)
						}
						return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("av")).Cardinality())
					})()).Cmp(_dafny.IntOfInt64(540)) != 0
					if (_this).F17() {
						var _pat_let_tv24 = _662_v110
						_ = _pat_let_tv24
						var _pat_let_tv25 = _662_v110
						_ = _pat_let_tv25
						_669_v115 = func(_pat_let14_0 D3) D3 {
							return func(_684_dt__update__tmp_h5 D3) D3 {
								return func(_pat_let17_0 _dafny.CodePoint) D3 {
									return func(_685_dt__update_hcf11_h1 _dafny.CodePoint) D3 {
										return Companion_D3_.Create_DC5_(_685_dt__update_hcf11_h1)
									}(_pat_let17_0)
								}(_pat_let_tv25)
							}(_pat_let14_0)
						}(func(_pat_let15_0 D3) D3 {
							return func(_682_dt__update__tmp_h4 D3) D3 {
								return func(_pat_let16_0 _dafny.CodePoint) D3 {
									return func(_683_dt__update_hcf11_h0 _dafny.CodePoint) D3 {
										return Companion_D3_.Create_DC5_(_683_dt__update_hcf11_h0)
									}(_pat_let16_0)
								}(_pat_let_tv24)
							}(_pat_let15_0)
						}(_669_v115))
						var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index90
						((_this).F14()).ArraySet1(false, (_index90).Int())
						var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index91
						((_this).F14()).ArraySet1(!((_this).F13()), (_index91).Int())
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(879), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index92
						((_this).F14()).ArraySet1((_this).F17(), (_index92).Int())
						var _686_v117 _dafny.Sequence
						_ = _686_v117
						_686_v117 = _dafny.SeqOf((_this).F21())
						var _687_v118 _dafny.Map
						_ = _687_v118
						_687_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_686_v117).Cardinality()), _626_v101)
						_687_v118 = (_687_v118).Update((_this).F21(), (_626_v101).Merge(_626_v101))
						_623_v98 = !((_this).F13())
					} else {
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))
						_ = _index93
						(_467_v1).ArraySet1((_this).F21(), (_index93).Int())
						_626_v101 = (_626_v101).Update((_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), (_this).F22())
						var _688_v119 *C0
						_ = _688_v119
						var _nw77 *C0 = New_C0_()
						_ = _nw77
						_nw77.Ctor__(!((_this).F13()), _662_v110, _this.F12(), (func() bool {
							if !((_this).F22()) {
								return !(true)
							}
							return (_this).F22()
						})())
						_688_v119 = _nw77
						_623_v98 = (_this).F17()
						var _689_v120 _dafny.MultiSet
						_ = _689_v120
						_689_v120 = _dafny.MultiSetOf(_dafny.IntOfInt64(449), (_467_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(607), _dafny.ArrayLen((_467_v1), 0))).Int()).(_dafny.Int), ((_this).F15()).Cardinality(), (_this).F21())
						var _690_v121 _dafny.Array
						_ = _690_v121
						var _nwElement0_18 _dafny.CodePoint = _662_v110
						_ = _nwElement0_18
						var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(21))
						_ = _nw78
						(_nw78).ArraySet1CodePoint(_nwElement0_18, 0)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 1)
						(_nw78).ArraySet1CodePoint(_662_v110, 2)
						(_nw78).ArraySet1CodePoint(_662_v110, 3)
						(_nw78).ArraySet1CodePoint(_662_v110, 4)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 5)
						(_nw78).ArraySet1CodePoint(_dafny.CodePoint('u'), 6)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 7)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 8)
						(_nw78).ArraySet1CodePoint(_662_v110, 9)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 10)
						(_nw78).ArraySet1CodePoint(_662_v110, 11)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 12)
						(_nw78).ArraySet1CodePoint(_dafny.CodePoint('v'), 13)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 14)
						(_nw78).ArraySet1CodePoint(_662_v110, 15)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 16)
						(_nw78).ArraySet1CodePoint((_688_v119).F20(), 17)
						(_nw78).ArraySet1CodePoint(_662_v110, 18)
						(_nw78).ArraySet1CodePoint(_dafny.CodePoint('p'), 19)
						(_nw78).ArraySet1CodePoint(_dafny.CodePoint('b'), 20)
						_690_v121 = _nw78
						var _rhs77 _dafny.Sequence = (func() _dafny.Sequence {
							if _623_v98 {
								return _dafny.UnicodeSeqOfUtf8Bytes("pwfti")
							}
							return _dafny.Companion_Sequence_.Concatenate(_466_v0, _466_v0)
						})()
						_ = _rhs77
						var _rhs78 bool = !((_689_v120).IsProperSubsetOf((_689_v120).Update(_dafny.IntOfInt64(365), Companion_Default___.Abs((_dafny.MultiSetOf(_690_v121)).Cardinality()))))
						_ = _rhs78
						var _lhs76 *C0 = _688_v119
						_ = _lhs76
						_466_v0 = _rhs77
						_lhs76.F19 = _rhs78
					}
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
	}
}
func (_this *C2) M5(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _691_v0 _dafny.Array
		_ = _691_v0
		var _nwElement0_19 _dafny.Int = (_this).F21()
		_ = _nwElement0_19
		var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(3))
		_ = _nw79
		(_nw79).ArraySet1(_nwElement0_19, 0)
		(_nw79).ArraySet1((_this).F21(), 1)
		(_nw79).ArraySet1((_this).F21(), 2)
		_691_v0 = _nw79
		for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_691_v0), 0))); ; {
			_guard_loop_4, _ok41 := _iter41()
			if !_ok41 {
				break
			}
			var _692_i0 _dafny.Int
			_692_i0 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_692_i0).Sign() != -1) && ((_692_i0).Cmp(_dafny.ArrayLen((_691_v0), 0)) < 0)) {
				(_691_v0).ArraySet1((_692_i0).Times((p0).Plus(p0)), (_692_i0).Int())
			}
		}
		(globalState).F7 = ((_this).F21()).Plus((_dafny.Zero).Minus(p0))
		var _693_v1 _dafny.CodePoint
		_ = _693_v1
		_693_v1 = _dafny.CodePoint('m')
		var _694_v2 *C0
		_ = _694_v2
		var _nw80 *C0 = New_C0_()
		_ = _nw80
		_nw80.Ctor__((_this).F22(), _693_v1, (func() D1 {
			if (_this).F22() {
				return _this.F12()
			}
			return Companion_D1_.Create_DC2_((_this).F21(), (_this).Fm8((_this).F22(), _this.F16(), globalState), (_this).F13())
		})(), (_this).F13())
		_694_v2 = _nw80
		var _695_v3 _dafny.Array
		_ = _695_v3
		var _len0_11 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_11
		var _nw81 _dafny.Array
		_ = _nw81
		if _len0_11.Cmp(_dafny.Zero) == 0 {
			_nw81 = _dafny.NewArray(_len0_11)
		} else {
			var _init11 func(_dafny.Int) _dafny.Sequence = (func(_696_v2 *C0) func(_dafny.Int) _dafny.Sequence {
				return func(_697_i1 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_696_v2.F19, (_this).F17())
				}
			})(_694_v2)
			_ = _init11
			var _element0_11 = _init11(_dafny.Zero)
			_ = _element0_11
			_nw81 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
			(_nw81).ArraySet1(_element0_11, 0)
			var _nativeLen0_11 = (_len0_11).Int()
			_ = _nativeLen0_11
			for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
				(_nw81).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
			}
		}
		_695_v3 = _nw81
		var _698_v4 D8
		_ = _698_v4
		_698_v4 = Companion_D8_.Create_DC18_(_695_v3)
		var _699_v5 _dafny.Map
		_ = _699_v5
		_699_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_698_v4).Dtor_cf31(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(762))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg59 _dafny.Int) interface{} {
				return coer59(arg59)
			}
		}((func(_700_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_701_i2 _dafny.Int) _dafny.Int {
				return _700_p0
			}
		})(p0))))
		var _702_v6 _dafny.Sequence
		_ = _702_v6
		_702_v6 = _dafny.SeqOf(p0)
		var _703_v7 _dafny.Sequence
		_ = _703_v7
		_703_v7 = _dafny.SeqOf(p0, (_this).F21(), (_this).F21(), _dafny.IntOfUint32((_702_v6).Cardinality()), p0)
		_699_v5 = (_699_v5).Update(_695_v3, _702_v6)
		if (_this).F17() {
			var _704_v8 _dafny.Sequence
			_ = _704_v8
			_704_v8 = _dafny.SeqOf(_694_v2.F19, _694_v2.F19, (_this).F13())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_695_v3), 0))
			_ = _index94
			(_695_v3).ArraySet1(_704_v8, (_index94).Int())
			var _705_v9 _dafny.Array
			_ = _705_v9
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_12
			var _nw82 _dafny.Array
			_ = _nw82
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw82 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) _dafny.Map = (func(_706_v2 *C0, _707_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.Map {
					return func(_708_i3 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_706_v2).F20(), _dafny.SeqOf(_707_v1))
					}
				})(_694_v2, _693_v1)
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw82 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw82).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw82).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_705_v9 = _nw82
			var _709_v10 _dafny.Sequence
			_ = _709_v10
			_709_v10 = _dafny.SeqOf((_694_v2).F20(), Companion_Default___.Fm33(_694_v2.F19, p0, globalState))
			var _710_v11 _dafny.Map
			_ = _710_v11
			_710_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_693_v1, _709_v10)
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_705_v9), 0))
			_ = _index95
			(_705_v9).ArraySet1(_710_v11, (_index95).Int())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_695_v3), 0))
			_ = _index96
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_705_v9), 0))
			_ = _index97
			var _rhs79 _dafny.Int = Companion_Default___.Fm0(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_704_v8, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_704_v8).Cardinality()))).Uint32(), _694_v2.F19), (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_704_v8, (Companion_Default___.SafeIndex((_this).F21(), _dafny.IntOfUint32((_704_v8).Cardinality()))).Uint32(), _694_v2.F19)).Cardinality()))).Uint32(), (_this).F17()), !((func() bool {
				if !((_this).F13()) {
					return (_this).F17()
				}
				return _694_v2.F19
			})()), (func() bool {
				if _694_v2.F19 {
					return Companion_Default___.Fm30((_702_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_702_v6).Cardinality()))).Uint32()).(_dafny.Int), globalState)
				}
				return _694_v2.F19
			})(), Companion_Default___.SafeModuloInt(p0, p0), globalState)
			_ = _rhs79
			var _rhs80 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_704_v8, _704_v8)
			_ = _rhs80
			var _rhs81 _dafny.Map = _710_v11
			_ = _rhs81
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 _dafny.Array = _695_v3
			_ = _lhs78
			var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_695_v3), 0))
			_ = _lhs79
			var _lhs80 _dafny.Array = _705_v9
			_ = _lhs80
			var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.ArrayLen((_705_v9), 0))
			_ = _lhs81
			_lhs77.F7 = _rhs79
			(_lhs78).ArraySet1(_rhs80, (_lhs79).Int())
			(_lhs80).ArraySet1(_rhs81, (_lhs81).Int())
			(globalState).F7 = (_this).Fm9((_694_v2).F20(), globalState)
			var _711_v12 _dafny.MultiSet
			_ = _711_v12
			_711_v12 = _dafny.MultiSetOf((_this).F21(), (_this).F21())
			var _712_v13 _dafny.MultiSet
			_ = _712_v13
			_712_v13 = _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32(((_695_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_695_v3), 0))).Int()).(_dafny.Sequence)).Cardinality()), (_dafny.MultiSetFromSeq((_695_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_695_v3), 0))).Int()).(_dafny.Sequence))).Cardinality(), (_this).F21())).Intersection((_711_v12).Update(p0, Companion_Default___.Abs(_dafny.IntOfUint32((_709_v10).Cardinality())))))
			(globalState).F7 = (func() _dafny.Int {
				if (_712_v13).Contains(_711_v12) {
					return (_712_v13).Multiplicity(_711_v12)
				}
				return (_this).F21()
			})()
			var _713_v14 _dafny.Set
			_ = _713_v14
			_713_v14 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_709_v10, _709_v10), _dafny.UnicodeSeqOfUtf8Bytes("wk"), _709_v10, _dafny.Companion_Sequence_.Concatenate(_709_v10, _709_v10))
			_713_v14 = _713_v14
			(_694_v2).F19 = (_this).F22()
		} else {
			var _714_v15 _dafny.MultiSet
			_ = _714_v15
			_714_v15 = _dafny.MultiSetOf((_this).F22(), _694_v2.F19)
			var _715_v16 _dafny.Map
			_ = _715_v16
			_715_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_714_v15, _694_v2.F19)
			var _716_v17 _dafny.Map
			_ = _716_v17
			_716_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_714_v15, _694_v2.F19))
			_715_v16 = (func() _dafny.Map {
				if (_716_v17).Contains((_this).F21()) {
					return (_716_v17).Get((_this).F21()).(_dafny.Map)
				}
				return _715_v16
			})()
			if (_this).F13() {
				var _717_v18 _dafny.Sequence
				_ = _717_v18
				_717_v18 = _dafny.UnicodeSeqOfUtf8Bytes("b")
				var _718_v19 D6
				_ = _718_v19
				_718_v19 = Companion_D6_.Create_DC13_((_this).F17(), false, (_this).F21(), (_this).F21(), _694_v2.F19)
				var _719_v20 _dafny.Set
				_ = _719_v20
				_719_v20 = _dafny.SetOf(_694_v2.F19)
				var _720_v21 T1
				_ = _720_v21
				var _nw83 *C1 = New_C1_()
				_ = _nw83
				_nw83.Ctor__((_694_v2).F20(), _this.F12(), (_this).F17())
				_720_v21 = _nw83
				var _721_v22 _dafny.MultiSet
				_ = _721_v22
				_721_v22 = _dafny.MultiSetOf(_720_v21)
				var _722_v23 _dafny.Array
				_ = _722_v23
				var _nwElement0_20 bool = _694_v2.F19
				_ = _nwElement0_20
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(22))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_20, 0)
				(_nw84).ArraySet1((_this).F13(), 1)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_717_v18, Companion_Default___.Fm32(Companion_D6_.Create_DC14_(_718_v19), globalState)), 2)
				(_nw84).ArraySet1(_694_v2.F19, 3)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_694_v2.F19, (_this).F22()), Companion_Default___.Fm38((_this).F13(), (_this).F22(), _694_v2.F19, (_this).F17(), globalState)), 4)
				(_nw84).ArraySet1((_this).F22(), 5)
				(_nw84).ArraySet1(Companion_Default___.Fm30(_dafny.IntOfInt64(-545), globalState), 6)
				(_nw84).ArraySet1(Companion_Default___.Fm30(p0, globalState), 7)
				(_nw84).ArraySet1(_694_v2.F19, 8)
				(_nw84).ArraySet1((_this).F13(), 9)
				(_nw84).ArraySet1((func() bool {
					if (_this).F13() {
						return !(_694_v2.F19)
					}
					return true
				})(), 10)
				(_nw84).ArraySet1((p0).Cmp((_702_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.IntOfUint32((_702_v6).Cardinality()))).Uint32()).(_dafny.Int)) >= 0, 11)
				(_nw84).ArraySet1(((_this).F21()).Cmp(_dafny.IntOfInt64(274)) <= 0, 12)
				(_nw84).ArraySet1((_this).F17(), 13)
				(_nw84).ArraySet1((_719_v20).IsSubsetOf(_719_v20), 14)
				(_nw84).ArraySet1((_this).F13(), 15)
				(_nw84).ArraySet1((_this).F17(), 16)
				(_nw84).ArraySet1((_721_v22).IsSubsetOf(_721_v22), 17)
				(_nw84).ArraySet1(((_this).F21()).Cmp(p0) > 0, 18)
				(_nw84).ArraySet1((_dafny.IntOfUint32((_703_v7).Cardinality())).Cmp((_this).F21()) <= 0, 19)
				(_nw84).ArraySet1(_694_v2.F19, 20)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_717_v18, _717_v18), 21)
				_722_v23 = _nw84
				_722_v23 = _722_v23
				var _723_v24 _dafny.Array
				_ = _723_v24
				var _nw85 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw85
				_723_v24 = _nw85
				var _724_v25 _dafny.Array
				_ = _724_v25
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
				_ = _nw86
				_724_v25 = _nw86
				var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_723_v24), 0))
				_ = _index98
				(_723_v24).ArraySet1(_724_v25, (_index98).Int())
				var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(815), _dafny.ArrayLen((_723_v24), 0))
				_ = _index99
				(_723_v24).ArraySet1(_724_v25, (_index99).Int())
				(_694_v2).F19 = (func() bool {
					if _694_v2.F19 {
						return !((_720_v21).F13())
					}
					return Companion_Default___.Fm30(p0, globalState)
				})()
				var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_722_v23), 0))
				_ = _index100
				(_722_v23).ArraySet1(!((_this).F13()), (_index100).Int())
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen((_722_v23), 0))
				_ = _index101
				(_722_v23).ArraySet1((_this).F17(), (_index101).Int())
				var _725_v26 _dafny.Map
				_ = _725_v26
				_725_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_720_v21).F13())
				_702_v6 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_702_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-120))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_726_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_727_i4 _dafny.Int) _dafny.Int {
						return _726_p0
					}
				})(p0)))), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_702_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-120))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg61 _dafny.Int) interface{} {
						return coer61(arg61)
					}
				}((func(_728_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_729_i4 _dafny.Int) _dafny.Int {
						return _728_p0
					}
				})(p0))))).Cardinality()))).Uint32(), (_725_v26).Cardinality())
			} else {
				var _730_v27 _dafny.Sequence
				_ = _730_v27
				_730_v27 = _dafny.SeqOf(true, (_this).F22(), (_this).F22())
				(globalState).F7 = Companion_Default___.Fm0(_730_v27, false, (_this).F22(), (p0).Minus((_this).F21()), globalState)
				var _731_v28 _dafny.Array
				_ = _731_v28
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
				_ = _nw87
				_731_v28 = _nw87
				var _732_v29 _dafny.Set
				_ = _732_v29
				_732_v29 = _dafny.SetOf(_694_v2.F19)
				var _733_v30 D8
				_ = _733_v30
				_733_v30 = Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_694_v2.F19), (_this).F21()), _732_v29, _dafny.MultiSetOf(_dafny.IntOfInt64(680)), _this.F16())
				var _734_v31 _dafny.Sequence
				_ = _734_v31
				_734_v31 = _dafny.SeqOf(_732_v29, (_733_v30).Dtor_cf33(), _732_v29, _732_v29, _732_v29)
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_731_v28), 0))
				_ = _index102
				(_731_v28).ArraySet1(_734_v31, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_731_v28), 0))
				_ = _index103
				(_731_v28).ArraySet1(_734_v31, (_index103).Int())
				var _735_v32 _dafny.MultiSet
				_ = _735_v32
				var _out8 _dafny.MultiSet
				_ = _out8
				_out8 = Companion_Default___.M0((_694_v2).F20(), (_this).F22(), globalState)
				_735_v32 = _out8
				(globalState).F7 = (_dafny.Zero).Minus(((p0).Times(p0)).Times(p0))
				(globalState).F7 = (_dafny.Zero).Minus((func() _dafny.Int {
					if !((_this).F22()) || (_694_v2.F19) {
						return p0
					}
					return p0
				})())
			}
			var _736_v33 D6
			_ = _736_v33
			_736_v33 = Companion_D6_.Create_DC11_(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_702_v6, _702_v6)))
			_736_v33 = _736_v33
			r0 = _694_v2.F19
			(globalState).F7 = ((_this).F21()).Minus((_dafny.Zero).Minus(p0))
		}
		(globalState).F7 = (_dafny.IntOfInt64(-221)).Minus(p0)
		r0 = (_this).F22()
		return r0
	}
}
func (_this *C2) M13(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Set, p3 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _737_v0 _dafny.Array
		_ = _737_v0
		var _nw88 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
		_ = _nw88
		_737_v0 = _nw88
		_737_v0 = _737_v0
		for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_737_v0), 0))); ; {
			_guard_loop_5, _ok42 := _iter42()
			if !_ok42 {
				break
			}
			var _738_i0 _dafny.Int
			_738_i0 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_738_i0).Sign() != -1) && ((_738_i0).Cmp(_dafny.ArrayLen((_737_v0), 0)) < 0)) {
				(_737_v0).ArraySet1((_this).F17(), (_738_i0).Int())
			}
		}
		var _739_v1 _dafny.Map
		_ = _739_v1
		_739_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), (_this).F21())
		var _740_v2 _dafny.Sequence
		_ = _740_v2
		_740_v2 = _dafny.SeqOf((_739_v1).Cardinality(), (_this).F21(), (_this).F21())
		_740_v2 = _740_v2
		var _741_v3 _dafny.Map
		_ = _741_v3
		_741_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_dafny.IntOfInt64(908)).Cmp(p3) == 0), (_this).F22())
		_741_v3 = (_741_v3).Update(((_this).F17()) == (p0), (_this).F17())
		if (_this).F22() {
			(globalState).F7 = (_this).Fm8((_this).F17(), _this.F16(), globalState)
			var _742_v4 _dafny.Array
			_ = _742_v4
			var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(23))
			_ = _nw89
			_742_v4 = _nw89
			_742_v4 = _742_v4
			var _743_v5 D3
			_ = _743_v5
			_743_v5 = Companion_D3_.Create_DC5_(p1)
			var _744_v6 *C1
			_ = _744_v6
			var _nw90 *C1 = New_C1_()
			_ = _nw90
			_nw90.Ctor__((_743_v5).Dtor_cf11(), Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-728), p3, (_this).F17()), true)
			_744_v6 = _nw90
			var _745_v7 _dafny.MultiSet
			_ = _745_v7
			var _out9 _dafny.MultiSet
			_ = _out9
			_out9 = Companion_Default___.M0(_744_v6.F23, (_this).F22(), globalState)
			_745_v7 = _out9
			r0 = Companion_Default___.Fm30(p3, globalState)
		} else {
			var _746_v8 _dafny.Sequence
			_ = _746_v8
			_746_v8 = _dafny.SeqOf(p0)
			(globalState).F7 = ((_dafny.Zero).Minus((_this).F21())).Minus(Companion_Default___.Fm0(_746_v8, (_this).F17(), (_this).F13(), (_this).F21(), globalState))
			var _747_v9 _dafny.Sequence
			_ = _747_v9
			_747_v9 = _dafny.UnicodeSeqOfUtf8Bytes("rkdajmk")
			(globalState).F7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_747_v9, _dafny.UnicodeSeqOfUtf8Bytes("lvxjwypfe"))).Cardinality())
			var _748_v10 _dafny.Map
			_ = _748_v10
			_748_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F21(), _737_v0)
			_748_v10 = (_748_v10).Update((_this).F21(), (_this).F14())
			(globalState).F7 = p3
			var _749_v11 _dafny.Array
			_ = _749_v11
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw91
			_749_v11 = _nw91
			var _750_v12 _dafny.Map
			_ = _750_v12
			_750_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_v0, _749_v11)
			_750_v12 = (_750_v12).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_737_v0, _749_v11))
		}
		r0 = !((_this).F13())
		r0 = (_this).F22()
		return r0
	}
}
func (_this *C2) F21() _dafny.Int {
	{
		return _this._f21
	}
}
func (_this *C2) F22() bool {
	{
		return _this._f22
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f12 D1
	_f16 D1
	_f14 _dafny.Array
	_f15 _dafny.Set
	_f13 bool
	_f17 bool
	_f18 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f12 = Companion_D1_.Default()
	_this._f16 = Companion_D1_.Default()
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.EmptySet
	_this._f13 = false
	_this._f17 = false
	_this._f18 = false
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T0_.TraitID_, Companion_T1_.TraitID_, Companion_T2_.TraitID_}
}

var _ T3 = &C3{}
var _ T0 = &C3{}
var _ T1 = &C3{}
var _ T2 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F12() D1 {
	return _this._f12
}
func (_this *C3) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C3) F16() D1 {
	return _this._f16
}
func (_this *C3) F16_set_(value D1) {
	_this._f16 = value
}
func (_this *C3) F14() _dafny.Array {
	return _this._f14
}
func (_this *C3) F15() _dafny.Set {
	return _this._f15
}
func (_this *C3) F13() bool {
	return _this._f13
}
func (_this *C3) F17() bool {
	return _this._f17
}
func (_this *C3) Ctor__(f18 bool, f16 D1, f17 bool, f14 _dafny.Array, f15 _dafny.Set, f12 D1, f13 bool) {
	{
		(_this)._f18 = f18
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C3) Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F18()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfdhco")).Cardinality()))
	}
}
func (_this *C3) Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(144)).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality())), _dafny.IntOfInt64(150))
	}
}
func (_this *C3) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfInt64(874)).Plus((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F17())).Cardinality())).Cardinality())))).Times((func() _dafny.Int {
			if false {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F17())).Cardinality()
			}
			return _dafny.IntOfInt64(537)
		})())
	}
}
func (_this *C3) Fm15(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(981), _dafny.IntOfInt64(611))).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-191), _dafny.IntOfInt64(434))))
	}
}
func (_this *C3) M6(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		if (_this.F12()).Dtor_cf4() {
			var _751_v0 _dafny.Array
			_ = _751_v0
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_13
			var _nw92 _dafny.Array
			_ = _nw92
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw92 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Int = (func(_752_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_753_i0 _dafny.Int) _dafny.Int {
						return (_753_i0).Minus((_dafny.MultiSetOf((_this).F13(), _752_p0, (_this).F18(), (_this).F17())).Cardinality())
					}
				})(p0)
				_ = _init13
				var _element0_13 = _init13(_dafny.Zero)
				_ = _element0_13
				_nw92 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
				(_nw92).ArraySet1(_element0_13, 0)
				var _nativeLen0_13 = (_len0_13).Int()
				_ = _nativeLen0_13
				for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
					(_nw92).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
				}
			}
			_751_v0 = _nw92
			var _754_v1 _dafny.Int
			_ = _754_v1
			_754_v1 = _dafny.IntOfInt64(673)
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))
			_ = _index104
			(_751_v0).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_754_v1, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(31))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}((func(_755_p0 bool) func(_dafny.Int) _dafny.Map {
				return func(_756_i1 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.IntOfUint32((_dafny.SeqOf(_755_p0)).Cardinality()))
				}
			})(p0)))).Cardinality()))), (_index104).Int())
			var _757_v2 _dafny.Map
			_ = _757_v2
			_757_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true) || ((_this).F13()), !((_this).F17()))
			var _758_v3 _dafny.Sequence
			_ = _758_v3
			_758_v3 = _dafny.SeqOf(_dafny.IntOfInt64(111))
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))
			_ = _index105
			var _rhs82 _dafny.Int = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_758_v3).Cardinality()), _754_v1)).Plus((_758_v3).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_754_v1), _dafny.IntOfUint32((_758_v3).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs82
			var _rhs83 _dafny.Map = _757_v2
			_ = _rhs83
			var _rhs84 _dafny.Int = _754_v1
			_ = _rhs84
			var _lhs82 _dafny.Array = _751_v0
			_ = _lhs82
			var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))
			_ = _lhs83
			(_lhs82).ArraySet1(_rhs82, (_lhs83).Int())
			_757_v2 = _rhs83
			_754_v1 = _rhs84
			var _759_v4 bool
			_ = _759_v4
			_759_v4 = true
			var _760_v5 _dafny.Sequence
			_ = _760_v5
			_760_v5 = _dafny.UnicodeSeqOfUtf8Bytes("db")
			_759_v4 = _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("rgxktld"), _760_v5)
			var _761_v6 _dafny.Map
			_ = _761_v6
			_761_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_751_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))).Int()).(_dafny.Int), ((_751_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))).Int()).(_dafny.Int)).Minus(_754_v1))
			var _762_v7 _dafny.Sequence
			_ = _762_v7
			_762_v7 = _dafny.SeqOf((_this).F17(), (_this).F13())
			_761_v6 = (_761_v6).Update((_751_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))).Int()).(_dafny.Int), Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_762_v7, !(true), (_this).F17(), _754_v1, globalState), (_751_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))).Int()).(_dafny.Int)))
			var _763_v8 _dafny.CodePoint
			_ = _763_v8
			_763_v8 = _dafny.CodePoint('w')
			var _764_v9 T0
			_ = _764_v9
			var _nw93 *C0 = New_C0_()
			_ = _nw93
			_nw93.Ctor__((_this).F13(), _763_v8, _this.F12(), (_this).F18())
			_764_v9 = _nw93
			var _765_v10 _dafny.Sequence
			_ = _765_v10
			_765_v10 = _dafny.SeqOf(_764_v9, _764_v9)
			var _766_v11 _dafny.Map
			_ = _766_v11
			_766_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _dafny.IntOfUint32((_765_v10).Cardinality()))
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_751_v0), 0))
			_ = _index106
			(_751_v0).ArraySet1((_758_v3).Select((Companion_Default___.SafeIndex((_766_v11).Cardinality(), _dafny.IntOfUint32((_758_v3).Cardinality()))).Uint32()).(_dafny.Int), (_index106).Int())
			r0 = _754_v1
		} else {
			var _767_v12 _dafny.Int
			_ = _767_v12
			_767_v12 = _dafny.IntOfInt64(352)
			var _768_v13 _dafny.MultiSet
			_ = _768_v13
			_768_v13 = _dafny.MultiSetOf(p0)
			var _769_v14 _dafny.Sequence
			_ = _769_v14
			_769_v14 = _dafny.SeqOf(_768_v13)
			var _770_v15 _dafny.Map
			_ = _770_v15
			_770_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_767_v12, _dafny.IntOfUint32((_769_v14).Cardinality()))
			(globalState).F7 = Companion_Default___.SafeDivisionInt((_770_v15).Cardinality(), _dafny.IntOfInt64(6))
			r0 = _767_v12
			if (_this).F18() {
				var _rhs85 _dafny.Int = ((_this).F15()).Cardinality()
				_ = _rhs85
				var _rhs86 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-213))
				_ = _rhs86
				var _lhs84 *GlobalState = globalState
				_ = _lhs84
				_lhs84.F7 = _rhs85
				_767_v12 = _rhs86
				var _771_v16 _dafny.Sequence
				_ = _771_v16
				_771_v16 = _dafny.SeqOf((_this).F14())
				_767_v12 = _dafny.IntOfUint32((_771_v16).Cardinality())
				var _772_v17 bool
				_ = _772_v17
				_772_v17 = true
				_772_v17 = Companion_Default___.Fm17(globalState)
				var _773_v18 _dafny.Sequence
				_ = _773_v18
				_773_v18 = _dafny.SeqOf(_767_v12)
				var _774_v19 _dafny.MultiSet
				_ = _774_v19
				_774_v19 = _dafny.MultiSetOf(_dafny.IntOfUint32((_773_v18).Cardinality()))
				var _775_v20 _dafny.Sequence
				_ = _775_v20
				_775_v20 = _dafny.UnicodeSeqOfUtf8Bytes("j")
				var _776_v21 _dafny.Array
				_ = _776_v21
				var _nw94 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
				_ = _nw94
				_776_v21 = _nw94
				var _777_v22 D2
				_ = _777_v22
				_777_v22 = Companion_D2_.Create_DC4_(_776_v21, (_this).F18(), _dafny.IntOfInt64(945), Companion_Default___.Fm17(globalState), (_this).F14())
				var _778_v23 _dafny.Array
				_ = _778_v23
				var _nwElement0_21 _dafny.Int = _dafny.IntOfUint32((_773_v18).Cardinality())
				_ = _nwElement0_21
				var _nw95 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(17))
				_ = _nw95
				(_nw95).ArraySet1(_nwElement0_21, 0)
				(_nw95).ArraySet1((_dafny.Zero).Minus(_767_v12), 1)
				(_nw95).ArraySet1(_767_v12, 2)
				(_nw95).ArraySet1(_767_v12, 3)
				(_nw95).ArraySet1(Companion_Default___.SafeDivisionInt(_767_v12, _767_v12), 4)
				(_nw95).ArraySet1((_770_v15).Cardinality(), 5)
				(_nw95).ArraySet1(_767_v12, 6)
				(_nw95).ArraySet1(_767_v12, 7)
				(_nw95).ArraySet1(_767_v12, 8)
				(_nw95).ArraySet1(Companion_Default___.SafeDivisionInt(_767_v12, _767_v12), 9)
				(_nw95).ArraySet1((_774_v19).Cardinality(), 10)
				(_nw95).ArraySet1(Companion_Default___.SafeDivisionInt(_767_v12, _767_v12), 11)
				(_nw95).ArraySet1(_dafny.IntOfUint32((_775_v20).Cardinality()), 12)
				(_nw95).ArraySet1((_777_v22).Dtor_cf8(), 13)
				(_nw95).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_773_v18, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_775_v20).Cardinality()), _dafny.IntOfUint32((_773_v18).Cardinality()))).Uint32(), _dafny.IntOfInt64(263))).Cardinality()), 14)
				(_nw95).ArraySet1(_767_v12, 15)
				(_nw95).ArraySet1(_767_v12, 16)
				_778_v23 = _nw95
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_776_v21), 0))
				_ = _index107
				(_776_v21).ArraySet1(_dafny.IntOfInt64(680), (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_776_v21), 0))
				_ = _index108
				var _rhs87 _dafny.Int = _767_v12
				_ = _rhs87
				var _rhs88 _dafny.Int = _767_v12
				_ = _rhs88
				var _lhs85 _dafny.Array = _776_v21
				_ = _lhs85
				var _lhs86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_776_v21), 0))
				_ = _lhs86
				(_lhs85).ArraySet1(_rhs87, (_lhs86).Int())
				_767_v12 = _rhs88
				var _779_v24 _dafny.CodePoint
				_ = _779_v24
				_779_v24 = _dafny.CodePoint('l')
				var _780_v25 D3
				_ = _780_v25
				_780_v25 = Companion_D3_.Create_DC5_(_dafny.CodePoint('s'))
				_779_v24 = (_780_v25).Dtor_cf11()
			} else {
				var _781_v26 _dafny.CodePoint
				_ = _781_v26
				_781_v26 = _dafny.CodePoint('w')
				var _782_v27 _dafny.Map
				_ = _782_v27
				_782_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), _781_v26)
				var _783_v28 _dafny.Sequence
				_ = _783_v28
				_783_v28 = _dafny.UnicodeSeqOfUtf8Bytes("fshc")
				var _784_v29 _dafny.Sequence
				_ = _784_v29
				_784_v29 = _dafny.SeqOf((_this).F13())
				var _785_v30 _dafny.Array
				_ = _785_v30
				var _nwElement0_22 _dafny.CodePoint = _781_v26
				_ = _nwElement0_22
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(25))
				_ = _nw96
				(_nw96).ArraySet1CodePoint(_nwElement0_22, 0)
				(_nw96).ArraySet1CodePoint(_781_v26, 1)
				(_nw96).ArraySet1CodePoint(_781_v26, 2)
				(_nw96).ArraySet1CodePoint((func() _dafny.CodePoint {
					if !((_this).F13()) {
						return _781_v26
					}
					return (func() _dafny.CodePoint {
						if (_782_v27).Contains(p0) {
							return (_782_v27).Get(p0).(_dafny.CodePoint)
						}
						return _781_v26
					})()
				})(), 3)
				(_nw96).ArraySet1CodePoint(_781_v26, 4)
				(_nw96).ArraySet1CodePoint(_781_v26, 5)
				(_nw96).ArraySet1CodePoint(_781_v26, 6)
				(_nw96).ArraySet1CodePoint(_dafny.CodePoint('a'), 7)
				(_nw96).ArraySet1CodePoint(_781_v26, 8)
				(_nw96).ArraySet1CodePoint(_dafny.CodePoint('k'), 9)
				(_nw96).ArraySet1CodePoint(_781_v26, 10)
				(_nw96).ArraySet1CodePoint(_dafny.CodePoint('i'), 11)
				(_nw96).ArraySet1CodePoint(_781_v26, 12)
				(_nw96).ArraySet1CodePoint(_781_v26, 13)
				(_nw96).ArraySet1CodePoint(_781_v26, 14)
				(_nw96).ArraySet1CodePoint(_dafny.CodePoint('b'), 15)
				(_nw96).ArraySet1CodePoint(Companion_Default___.Fm18(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ukupbqlom")).Cardinality()), _783_v28, (_this).F17(), (_this).Fm10((_this).F13(), (_this).F13(), _784_v29, globalState), globalState), 16)
				(_nw96).ArraySet1CodePoint(_dafny.CodePoint('w'), 17)
				(_nw96).ArraySet1CodePoint((func() _dafny.CodePoint {
					if !((_this).F18()) {
						return _781_v26
					}
					return _781_v26
				})(), 18)
				(_nw96).ArraySet1CodePoint((_783_v28).Select((Companion_Default___.SafeIndex((_this).Fm9(_781_v26, globalState), _dafny.IntOfUint32((_783_v28).Cardinality()))).Uint32()).(_dafny.CodePoint), 19)
				(_nw96).ArraySet1CodePoint(_781_v26, 20)
				(_nw96).ArraySet1CodePoint(_781_v26, 21)
				(_nw96).ArraySet1CodePoint(_781_v26, 22)
				(_nw96).ArraySet1CodePoint(_781_v26, 23)
				(_nw96).ArraySet1CodePoint(_781_v26, 24)
				_785_v30 = _nw96
				_785_v30 = _785_v30
				var _786_v31 bool
				_ = _786_v31
				_786_v31 = false
				_786_v31 = ((_767_v12).Plus(_767_v12)).Cmp(_767_v12) <= 0
				_786_v31 = _786_v31
				var _787_v32 _dafny.Array
				_ = _787_v32
				var _nw97 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(10))
				_ = _nw97
				_787_v32 = _nw97
				var _788_v33 _dafny.MultiSet
				_ = _788_v33
				_788_v33 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lrwaorxc")).Cardinality()))
				var _789_v34 _dafny.Sequence
				_ = _789_v34
				_789_v34 = _dafny.SeqOf(_788_v33)
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_787_v32), 0))
				_ = _index109
				(_787_v32).ArraySet1((func() _dafny.Sequence {
					if (_this).F17() {
						return _789_v34
					}
					return _dafny.SeqOf(_788_v33, _788_v33)
				})(), (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_787_v32), 0))
				_ = _index110
				(_787_v32).ArraySet1(_789_v34, (_index110).Int())
				var _790_v35 _dafny.Array
				_ = _790_v35
				var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw98
				_790_v35 = _nw98
				_790_v35 = _790_v35
			}
			var _791_v36 _dafny.CodePoint
			_ = _791_v36
			_791_v36 = _dafny.CodePoint('g')
			var _792_v37 *C0
			_ = _792_v37
			var _nw99 *C0 = New_C0_()
			_ = _nw99
			_nw99.Ctor__((_767_v12).Cmp(_767_v12) > 0, _791_v36, Companion_D1_.Create_DC2_(_767_v12, _767_v12, p0), p0)
			_792_v37 = _nw99
			var _793_v38 _dafny.Sequence
			_ = _793_v38
			_793_v38 = _dafny.SeqOf((_this).F14())
			var _794_v39 _dafny.Sequence
			_ = _794_v39
			_794_v39 = _dafny.UnicodeSeqOfUtf8Bytes("b")
			var _795_v40 _dafny.Array
			_ = _795_v40
			var _nwElement0_23 _dafny.Array = (_this).F14()
			_ = _nwElement0_23
			var _nw100 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(14))
			_ = _nw100
			(_nw100).ArraySet1(_nwElement0_23, 0)
			(_nw100).ArraySet1((_this).F14(), 1)
			(_nw100).ArraySet1((_this).F14(), 2)
			(_nw100).ArraySet1((_this).F14(), 3)
			(_nw100).ArraySet1((_this).F14(), 4)
			(_nw100).ArraySet1((_this).F14(), 5)
			(_nw100).ArraySet1((_this).F14(), 6)
			(_nw100).ArraySet1((_this).F14(), 7)
			(_nw100).ArraySet1((_this).F14(), 8)
			(_nw100).ArraySet1((_this).F14(), 9)
			(_nw100).ArraySet1((_this).F14(), 10)
			(_nw100).ArraySet1((_this).F14(), 11)
			(_nw100).ArraySet1((_793_v38).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_794_v39).Cardinality()), _dafny.IntOfUint32((_793_v38).Cardinality()))).Uint32()).(_dafny.Array), 12)
			(_nw100).ArraySet1((_this).F14(), 13)
			_795_v40 = _nw100
			var _796_v41 _dafny.Array
			_ = _796_v41
			_796_v41 = _795_v40
			var _797_v42 _dafny.Sequence
			_ = _797_v42
			_797_v42 = _dafny.SeqOf(_795_v40, _795_v40, (_796_v41), _795_v40)
			var _798_v43 _dafny.Array
			_ = _798_v43
			var _nwElement0_24 _dafny.Array = _795_v40
			_ = _nwElement0_24
			var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(14))
			_ = _nw101
			(_nw101).ArraySet1(_nwElement0_24, 0)
			(_nw101).ArraySet1(_795_v40, 1)
			(_nw101).ArraySet1(_795_v40, 2)
			(_nw101).ArraySet1(_795_v40, 3)
			(_nw101).ArraySet1(_795_v40, 4)
			(_nw101).ArraySet1((_797_v42).Select((Companion_Default___.SafeIndex(_767_v12, _dafny.IntOfUint32((_797_v42).Cardinality()))).Uint32()).(_dafny.Array), 5)
			(_nw101).ArraySet1(_795_v40, 6)
			(_nw101).ArraySet1(_795_v40, 7)
			(_nw101).ArraySet1((_797_v42).Select((Companion_Default___.SafeIndex(_767_v12, _dafny.IntOfUint32((_797_v42).Cardinality()))).Uint32()).(_dafny.Array), 8)
			(_nw101).ArraySet1(_795_v40, 9)
			(_nw101).ArraySet1(_795_v40, 10)
			(_nw101).ArraySet1(_795_v40, 11)
			(_nw101).ArraySet1(_795_v40, 12)
			(_nw101).ArraySet1((_797_v42).Select((Companion_Default___.SafeIndex(_767_v12, _dafny.IntOfUint32((_797_v42).Cardinality()))).Uint32()).(_dafny.Array), 13)
			_798_v43 = _nw101
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_798_v43), 0))
			_ = _index111
			(_798_v43).ArraySet1(_795_v40, (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(976), _dafny.ArrayLen((_798_v43), 0))
			_ = _index112
			(_798_v43).ArraySet1(_795_v40, (_index112).Int())
		}
		var _799_v44 _dafny.Array
		_ = _799_v44
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_14
		var _nw102 _dafny.Array
		_ = _nw102
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw102 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Map = (func(_800_p0 bool) func(_dafny.Int) _dafny.Map {
				return func(_801_i2 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F18()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F13()))).Cardinality())), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_800_p0, _dafny.IntOfUint32((_dafny.SeqOf((_this).F18())).Cardinality()))).Cardinality()))).Cardinality(), !(_800_p0))
				}
			})(p0)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw102 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw102).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw102).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_799_v44 = _nw102
		var _802_v45 _dafny.Int
		_ = _802_v45
		_802_v45 = _dafny.IntOfInt64(10)
		var _803_v46 _dafny.Map
		_ = _803_v46
		_803_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_799_v44, _802_v45)
		_803_v46 = (_803_v46).Update(_799_v44, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(771), _dafny.IntOfInt64(190)))
		if Companion_Default___.Fm17(globalState) {
			var _804_v47 _dafny.Array
			_ = _804_v47
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(22))
			_ = _nw103
			_804_v47 = _nw103
			var _805_v48 _dafny.Array
			_ = _805_v48
			_805_v48 = _804_v47
			var _806_v49 _dafny.Array
			_ = _806_v49
			var _nwElement0_25 _dafny.Array = _804_v47
			_ = _nwElement0_25
			var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(4))
			_ = _nw104
			(_nw104).ArraySet1(_nwElement0_25, 0)
			(_nw104).ArraySet1(_805_v48, 1)
			(_nw104).ArraySet1(_805_v48, 2)
			(_nw104).ArraySet1(_804_v47, 3)
			_806_v49 = _nw104
			var _rhs89 _dafny.Array = _806_v49
			_ = _rhs89
			var _rhs90 _dafny.Int = _802_v45
			_ = _rhs90
			var _lhs87 *GlobalState = globalState
			_ = _lhs87
			_806_v49 = _rhs89
			_lhs87.F7 = _rhs90
			var _807_v50 _dafny.Map
			_ = _807_v50
			_807_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _802_v45)
			var _808_v51 _dafny.Sequence
			_ = _808_v51
			_808_v51 = _dafny.UnicodeSeqOfUtf8Bytes("pwgijra")
			var _809_v52 _dafny.MultiSet
			_ = _809_v52
			_809_v52 = _dafny.MultiSetOf((_this).F13())
			var _810_v53 _dafny.Sequence
			_ = _810_v53
			_810_v53 = _dafny.SeqOf(_809_v52, _809_v52, _809_v52)
			var _811_v54 _dafny.Array
			_ = _811_v54
			var _nw105 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
			_ = _nw105
			_811_v54 = _nw105
			var _812_v55 D3
			_ = _812_v55
			_812_v55 = Companion_D3_.Create_DC7_()
			var _813_v56 _dafny.Map
			_ = _813_v56
			_813_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_811_v54, _812_v55)
			var _814_v57 _dafny.Sequence
			_ = _814_v57
			_814_v57 = _dafny.SeqOf(_802_v45)
			var _815_v58 _dafny.CodePoint
			_ = _815_v58
			_815_v58 = _dafny.CodePoint('p')
			var _816_v59 _dafny.Sequence
			_ = _816_v59
			_816_v59 = _dafny.SeqOf(_811_v54)
			var _817_v60 D3
			_ = _817_v60
			_817_v60 = Companion_D3_.Create_DC8_(_dafny.IntOfInt64(698))
			var _818_v61 _dafny.Array
			_ = _818_v61
			var _nwElement0_26 _dafny.Int = _802_v45
			_ = _nwElement0_26
			var _nw106 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(26))
			_ = _nw106
			(_nw106).ArraySet1(_nwElement0_26, 0)
			(_nw106).ArraySet1(_dafny.IntOfInt64(268), 1)
			(_nw106).ArraySet1((_802_v45).Times(_802_v45), 2)
			(_nw106).ArraySet1(_802_v45, 3)
			(_nw106).ArraySet1(_802_v45, 4)
			(_nw106).ArraySet1(_802_v45, 5)
			(_nw106).ArraySet1(Companion_Default___.SafeModuloInt(_802_v45, _802_v45), 6)
			(_nw106).ArraySet1((func() _dafny.Int {
				if (_807_v50).Contains((_this).F17()) {
					return (_807_v50).Get((_this).F17()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_808_v51).Cardinality())
			})(), 7)
			(_nw106).ArraySet1(((_810_v53).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm19(globalState)).Cardinality(), _dafny.IntOfUint32((_810_v53).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), 8)
			(_nw106).ArraySet1((_813_v56).Cardinality(), 9)
			(_nw106).ArraySet1(_802_v45, 10)
			(_nw106).ArraySet1(_802_v45, 11)
			(_nw106).ArraySet1(_802_v45, 12)
			(_nw106).ArraySet1(_802_v45, 13)
			(_nw106).ArraySet1(_802_v45, 14)
			(_nw106).ArraySet1((_814_v57).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(Companion_Default___.Fm20(globalState), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(90))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}(func(_819_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm20(globalState)).Cardinality()))).Uint32(), _815_v58)).Cardinality()), _dafny.IntOfUint32((_814_v57).Cardinality()))).Uint32()).(_dafny.Int), 15)
			(_nw106).ArraySet1((_this).Fm8((_this).F18(), Companion_Default___.Fm21(globalState), globalState), 16)
			(_nw106).ArraySet1(_802_v45, 17)
			(_nw106).ArraySet1(_802_v45, 18)
			(_nw106).ArraySet1(_802_v45, 19)
			(_nw106).ArraySet1((_this).Fm8(Companion_Default___.Fm17(globalState), _this.F16(), globalState), 20)
			(_nw106).ArraySet1(_802_v45, 21)
			(_nw106).ArraySet1(((_dafny.MultiSetOf(_802_v45)).Union(_dafny.MultiSetOf(_802_v45))).Cardinality(), 22)
			(_nw106).ArraySet1(Companion_Default___.SafeDivisionInt(_802_v45, _dafny.IntOfUint32((_816_v59).Cardinality())), 23)
			(_nw106).ArraySet1(_802_v45, 24)
			(_nw106).ArraySet1((_817_v60).Dtor_cf14(), 25)
			_818_v61 = _nw106
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_811_v54), 0))
			_ = _index113
			(_811_v54).ArraySet1(((func() _dafny.Int {
				if (_807_v50).Contains((_this).F17()) {
					return (_807_v50).Get((_this).F17()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_814_v57).Cardinality())
			})()).Minus(_802_v45), (_index113).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_811_v54), 0))
			_ = _index114
			(_811_v54).ArraySet1((_802_v45).Minus((func() _dafny.Int {
				if p0 {
					return _802_v45
				}
				return _dafny.IntOfUint32((Companion_Default___.Fm20(globalState)).Cardinality())
			})()), (_index114).Int())
			(globalState).F7 = (_dafny.Zero).Minus((_811_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_811_v54), 0))).Int()).(_dafny.Int))
			(globalState).F7 = _802_v45
			var _820_v63 _dafny.Array
			_ = _820_v63
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_15
			var _nw107 _dafny.Array
			_ = _nw107
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw107 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Set = (func(_821_p0 bool, _822_v58 _dafny.CodePoint, _823_v51 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
					return func(_824_i4 _dafny.Int) _dafny.Set {
						return (func() _dafny.Set {
							if _821_p0 {
								return func() _dafny.Set {
									var _coll35 = _dafny.NewBuilder()
									_ = _coll35
									for _iter43 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_823_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.IntOfUint32((_823_v51).Cardinality()))).Uint32()).(_dafny.CodePoint), !((_this).F17()))).Keys().Elements()); ; {
										_compr_35, _ok43 := _iter43()
										if !_ok43 {
											break
										}
										var _825_v62 _dafny.CodePoint
										_825_v62 = interface{}(_compr_35).(_dafny.CodePoint)
										if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_823_v51).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.IntOfUint32((_823_v51).Cardinality()))).Uint32()).(_dafny.CodePoint), !((_this).F17()))).Contains(_825_v62) {
											_coll35.Add(_825_v62)
										}
									}
									return _coll35.ToSet()
								}()
							}
							return _dafny.SetOf(_822_v58, _822_v58)
						})()
					}
				})(p0, _815_v58, _808_v51)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw107 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw107).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw107).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_820_v63 = _nw107
			var _826_v64 _dafny.Map
			_ = _826_v64
			_826_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(494), _dafny.UnicodeSeqOfUtf8Bytes("uh"))
			var _827_v65 _dafny.Set
			_ = _827_v65
			_827_v65 = _dafny.SetOf((func() _dafny.Sequence {
				if (_826_v64).Contains(_802_v45) {
					return (_826_v64).Get(_802_v45).(_dafny.Sequence)
				}
				return _808_v51
			})(), _dafny.UnicodeSeqOfUtf8Bytes("dgxndwfh"), _808_v51)
			var _828_v66 _dafny.Set
			_ = _828_v66
			_828_v66 = _827_v65
			var _829_v67 _dafny.Set
			_ = _829_v67
			_829_v67 = _dafny.SetOf(!(Companion_Default___.Fm17(globalState)), true, (_this).F13(), (_this).F13())
			var _rhs91 _dafny.Array = _820_v63
			_ = _rhs91
			var _rhs92 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_808_v51, _dafny.Companion_Sequence_.Concatenate(_808_v51, _dafny.UnicodeSeqOfUtf8Bytes("wxpgj")))
			_ = _rhs92
			var _rhs93 _dafny.Set = (_828_v66)
			_ = _rhs93
			var _rhs94 _dafny.Array = _811_v54
			_ = _rhs94
			var _rhs95 _dafny.Int = Companion_Default___.SafeModuloInt((_802_v45).Minus((_829_v67).Cardinality()), (_811_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen((_811_v54), 0))).Int()).(_dafny.Int))
			_ = _rhs95
			_820_v63 = _rhs91
			_808_v51 = _rhs92
			_827_v65 = _rhs93
			_818_v61 = _rhs94
			_802_v45 = _rhs95
		} else {
			var _830_v69 _dafny.Sequence
			_ = _830_v69
			_830_v69 = _dafny.SeqOf(_802_v45, _802_v45)
			var _831_v70 _dafny.MultiSet
			_ = _831_v70
			_831_v70 = _dafny.MultiSetOf(_dafny.IntOfUint32((_830_v69).Cardinality()), (_dafny.Zero).Minus(_802_v45))
			var _832_v71 _dafny.Sequence
			_ = _832_v71
			_832_v71 = _dafny.SeqOf(_831_v70)
			var _833_v73 D6
			_ = _833_v73
			_833_v73 = Companion_D6_.Create_DC11_(_831_v70)
			var _834_v74 _dafny.MultiSet
			_ = _834_v74
			_834_v74 = _dafny.MultiSetOf((_833_v73).Dtor_cf17())
			var _835_v75 _dafny.Map
			_ = _835_v75
			_835_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll36 = _dafny.NewMapBuilder()
				_ = _coll36
				for _iter44 := _dafny.Iterate((_832_v71).Elements()); ; {
					_compr_36, _ok44 := _iter44()
					if !_ok44 {
						break
					}
					var _836_v68 _dafny.MultiSet
					_836_v68 = interface{}(_compr_36).(_dafny.MultiSet)
					if _dafny.Companion_Sequence_.Contains(_832_v71, _836_v68) {
						_coll36.Add(_836_v68, true)
					}
				}
				return _coll36.ToMap()
			}()).Merge(func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter45 := _dafny.Iterate((_834_v74).Elements()); ; {
					_compr_37, _ok45 := _iter45()
					if !_ok45 {
						break
					}
					var _837_v72 _dafny.MultiSet
					_837_v72 = interface{}(_compr_37).(_dafny.MultiSet)
					if (_834_v74).Contains(_837_v72) {
						_coll37.Add(_837_v72, (_this).F13())
					}
				}
				return _coll37.ToMap()
			}()), false)
			var _838_v76 _dafny.Map
			_ = _838_v76
			_838_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_831_v70, (_this).F13())
			_835_v75 = (_835_v75).Update(_838_v76, p0)
			var _839_v77 D2
			_ = _839_v77
			_839_v77 = Companion_D2_.Create_DC3_((_this).F14())
			var _source15 D2 = _839_v77
			_ = _source15
			if _source15.Is_DC4() {
				var _840___mcc_h0 _dafny.Array = _source15.Get_().(D2_DC4).Cf6
				_ = _840___mcc_h0
				var _841___mcc_h1 bool = _source15.Get_().(D2_DC4).Cf7
				_ = _841___mcc_h1
				var _842___mcc_h2 _dafny.Int = _source15.Get_().(D2_DC4).Cf8
				_ = _842___mcc_h2
				var _843___mcc_h3 bool = _source15.Get_().(D2_DC4).Cf9
				_ = _843___mcc_h3
				var _844___mcc_h4 _dafny.Array = _source15.Get_().(D2_DC4).Cf10
				_ = _844___mcc_h4
				var _845_cf10 _dafny.Array = _844___mcc_h4
				_ = _845_cf10
				var _846_cf9 bool = _843___mcc_h3
				_ = _846_cf9
				var _847_cf8 _dafny.Int = _842___mcc_h2
				_ = _847_cf8
				var _848_cf7 bool = _841___mcc_h1
				_ = _848_cf7
				var _849_cf6 _dafny.Array = _840___mcc_h0
				_ = _849_cf6
				var _850_v78 _dafny.Sequence
				_ = _850_v78
				_850_v78 = _dafny.UnicodeSeqOfUtf8Bytes("fkeo")
				var _851_v79 _dafny.Set
				_ = _851_v79
				_851_v79 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("gn"), _850_v78)
				var _852_v80 _dafny.Set
				_ = _852_v80
				_852_v80 = _851_v79
				var _853_v81 _dafny.Map
				_ = _853_v81
				_853_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_852_v80, _852_v80, Companion_Default___.Fm22(_846_cf9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-793))).Uint32(), func(coer64 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}(func(_854_i5 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(-328)
				})), globalState), _852_v80, _851_v79), p0)
				var _855_v82 _dafny.MultiSet
				_ = _855_v82
				_855_v82 = _dafny.MultiSetOf(_852_v80, _852_v80)
				var _856_v83 _dafny.Sequence
				_ = _856_v83
				_856_v83 = _dafny.SeqOf(_846_cf9)
				_853_v81 = (_853_v81).Update((_855_v82).Intersection(_855_v82), (_856_v83).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_847_cf8)).Cardinality(), _dafny.IntOfUint32((_856_v83).Cardinality()))).Uint32()).(bool))
				var _857_v84 _dafny.CodePoint
				_ = _857_v84
				_857_v84 = _dafny.CodePoint('b')
				var _858_v85 _dafny.Set
				_ = _858_v85
				_858_v85 = _dafny.SetOf(_849_cf6, _849_cf6)
				var _859_v86 *C0
				_ = _859_v86
				var _nw108 *C0 = New_C0_()
				_ = _nw108
				_nw108.Ctor__((_this).F18(), _857_v84, _this.F12(), !(_858_v85).Contains(_849_cf6))
				_859_v86 = _nw108
				var _860_v87 _dafny.Sequence
				_ = _860_v87
				_860_v87 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("h"), _dafny.UnicodeSeqOfUtf8Bytes("x")), _850_v78, _850_v78, _dafny.Companion_Sequence_.Concatenate((_this.F16()).Dtor_cf1(), _dafny.UnicodeSeqOfUtf8Bytes("te")), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer65 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg65 _dafny.Int) interface{} {
						return coer65(arg65)
					}
				}((func(_861_v86 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_862_i6 _dafny.Int) _dafny.CodePoint {
						return (Companion_D3_.Create_DC5_((_861_v86).F20())).Dtor_cf11()
					}
				})(_859_v86))))
				var _863_v88 _dafny.Map
				_ = _863_v88
				_863_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_847_cf8, _860_v87)
				_860_v87 = (func() _dafny.Sequence {
					if (_863_v88).Contains(_802_v45) {
						return (_863_v88).Get(_802_v45).(_dafny.Sequence)
					}
					return _860_v87
				})()
				_846_cf9 = _848_cf7
			} else {
				var _864___mcc_h5 _dafny.Array = _source15.Get_().(D2_DC3).Cf5
				_ = _864___mcc_h5
				var _865_cf5 _dafny.Array = _864___mcc_h5
				_ = _865_cf5
				var _rhs96 _dafny.Int = _802_v45
				_ = _rhs96
				var _rhs97 _dafny.Int = (_dafny.Zero).Minus(_802_v45)
				_ = _rhs97
				var _lhs88 *GlobalState = globalState
				_ = _lhs88
				_lhs88.F7 = _rhs96
				r0 = _rhs97
				var _866_v89 _dafny.Sequence
				_ = _866_v89
				_866_v89 = _dafny.SeqOf((_this).F13())
				(globalState).F2 = _866_v89
				var _867_v90 bool
				_ = _867_v90
				_867_v90 = true
				var _868_v91 _dafny.CodePoint
				_ = _868_v91
				_868_v91 = _dafny.CodePoint('x')
				var _869_v92 _dafny.Array
				_ = _869_v92
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_16
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) D1 = func(_870_i7 _dafny.Int) D1 {
						return _this.F12()
					}
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw109 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw109).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw109).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_869_v92 = _nw109
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_869_v92), 0))
				_ = _index115
				(_869_v92).ArraySet1(_this.F12(), (_index115).Int())
				var _871_v93 _dafny.Array
				_ = _871_v93
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(28)
				_ = _len0_17
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = (func(_872_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_873_i8 _dafny.Int) _dafny.Int {
							return (_873_i8).Plus(_872_v45)
						}
					})(_802_v45)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw110 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw110).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw110).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_871_v93 = _nw110
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_871_v93), 0))
				_ = _index116
				(_871_v93).ArraySet1(((_dafny.Zero).Minus(_802_v45)).Times(_802_v45), (_index116).Int())
				var _pat_let_tv26 = _802_v45
				_ = _pat_let_tv26
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_869_v92), 0))
				_ = _index117
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_871_v93), 0))
				_ = _index118
				var _rhs98 bool = _867_v90
				_ = _rhs98
				var _rhs99 _dafny.Int = _802_v45
				_ = _rhs99
				var _rhs100 _dafny.CodePoint = _868_v91
				_ = _rhs100
				var _rhs101 D1 = func(_pat_let18_0 D1) D1 {
					return func(_874_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let19_0 bool) D1 {
							return func(_875_dt__update_hcf4_h0 bool) D1 {
								return func(_pat_let20_0 _dafny.Int) D1 {
									return func(_876_dt__update_hcf2_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC2_(_876_dt__update_hcf2_h0, (_874_dt__update__tmp_h1).Dtor_cf3(), _875_dt__update_hcf4_h0)
									}(_pat_let20_0)
								}(_pat_let_tv26)
							}(_pat_let19_0)
						}((_this).F17())
					}(_pat_let18_0)
				}(_this.F12())
				_ = _rhs101
				var _rhs102 _dafny.Int = _802_v45
				_ = _rhs102
				var _lhs89 _dafny.Array = _869_v92
				_ = _lhs89
				var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_869_v92), 0))
				_ = _lhs90
				var _lhs91 _dafny.Array = _871_v93
				_ = _lhs91
				var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(644), _dafny.ArrayLen((_871_v93), 0))
				_ = _lhs92
				_867_v90 = _rhs98
				_802_v45 = _rhs99
				_868_v91 = _rhs100
				(_lhs89).ArraySet1(_rhs101, (_lhs90).Int())
				(_lhs91).ArraySet1(_rhs102, (_lhs92).Int())
				(globalState).F7 = _dafny.IntOfInt64(785)
			}
			var _877_v94 _dafny.Array
			_ = _877_v94
			var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw111
			_877_v94 = _nw111
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_877_v94), 0))
			_ = _index119
			(_877_v94).ArraySet1((_802_v45).Plus(_802_v45), (_index119).Int())
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(543), _dafny.ArrayLen((_877_v94), 0))
			_ = _index120
			(_877_v94).ArraySet1((_dafny.IntOfInt64(520)).Times(_802_v45), (_index120).Int())
			var _878_v95 _dafny.CodePoint
			_ = _878_v95
			_878_v95 = _dafny.CodePoint('t')
			var _879_v96 *C0
			_ = _879_v96
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__((_this).F13(), _878_v95, _this.F12(), Companion_Default___.Fm17(globalState))
			_879_v96 = _nw112
			_879_v96 = _879_v96
		}
		(_this).F12_set_(func(_pat_let21_0 D1) D1 {
			return func(_880_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let22_0 bool) D1 {
					return func(_881_dt__update_hcf4_h1 bool) D1 {
						return Companion_D1_.Create_DC2_((_880_dt__update__tmp_h2).Dtor_cf2(), (_880_dt__update__tmp_h2).Dtor_cf3(), _881_dt__update_hcf4_h1)
					}(_pat_let22_0)
				}((_this).F13())
			}(_pat_let21_0)
		}(Companion_D1_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), (_this).F13())).Cardinality(), _802_v45, (_this).F13())))
		if !((_802_v45).Cmp(_802_v45) < 0) {
			var _882_v97 _dafny.Map
			_ = _882_v97
			_882_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F18())
			var _883_v98 _dafny.Array
			_ = _883_v98
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
			_ = _nw113
			_883_v98 = _nw113
			var _884_v99 _dafny.MultiSet
			_ = _884_v99
			_884_v99 = _dafny.MultiSetOf(_883_v98)
			var _885_v100 _dafny.Map
			_ = _885_v100
			_885_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				if (_this).F17() {
					return _882_v97
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F18())
			})(), Companion_Default___.SafeModuloInt((_884_v99).Cardinality(), _dafny.IntOfInt64(527)))
			var _886_v101 _dafny.Map
			_ = _886_v101
			_886_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23((_this).F13(), globalState), _885_v100)
			var _887_v102 _dafny.CodePoint
			_ = _887_v102
			_887_v102 = _dafny.CodePoint('h')
			var _888_v103 _dafny.Map
			_ = _888_v103
			_888_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, _887_v102)
			var _889_v105 _dafny.Map
			_ = _889_v105
			_889_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, (_this).F17())
			_885_v100 = (func() _dafny.Map {
				if (_886_v101).Contains((_888_v103).Merge(func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter46 := _dafny.Iterate((_889_v105).Keys().Elements()); ; {
						_compr_38, _ok46 := _iter46()
						if !_ok46 {
							break
						}
						var _890_v104 _dafny.Int
						_890_v104 = interface{}(_compr_38).(_dafny.Int)
						if (_889_v105).Contains(_890_v104) {
							_coll38.Add((_890_v104).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F13(), (_this).F17())).Cardinality())), _887_v102)
						}
					}
					return _coll38.ToMap()
				}())) {
					return (_886_v101).Get((_888_v103).Merge(func() _dafny.Map {
						var _coll39 = _dafny.NewMapBuilder()
						_ = _coll39
						for _iter47 := _dafny.Iterate((_889_v105).Keys().Elements()); ; {
							_compr_39, _ok47 := _iter47()
							if !_ok47 {
								break
							}
							var _891_v104 _dafny.Int
							_891_v104 = interface{}(_compr_39).(_dafny.Int)
							if (_889_v105).Contains(_891_v104) {
								_coll39.Add((_891_v104).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F13(), (_this).F17())).Cardinality())), _887_v102)
							}
						}
						return _coll39.ToMap()
					}())).(_dafny.Map)
				}
				return func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter48 := _dafny.Iterate((_885_v100).Keys().Elements()); ; {
						_compr_40, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _892_v106 _dafny.Map
						_892_v106 = interface{}(_compr_40).(_dafny.Map)
						if (_885_v100).Contains(_892_v106) {
							_coll40.Add(_892_v106, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_887_v102, (_this).F13())).Cardinality())
						}
					}
					return _coll40.ToMap()
				}()
			})()
			var _893_v107 *C0
			_ = _893_v107
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__((_this).F18(), _887_v102, _this.F12(), (_this).F17())
			_893_v107 = _nw114
			var _894_v108 _dafny.Map
			_ = _894_v108
			_894_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, _802_v45)
			var _895_v109 _dafny.Sequence
			_ = _895_v109
			_895_v109 = _dafny.SeqOf((_this).F13(), _893_v107.F19)
			var _rhs103 bool = Companion_Default___.Fm17(globalState)
			_ = _rhs103
			var _rhs104 _dafny.Int = _802_v45
			_ = _rhs104
			var _rhs105 _dafny.Map = Companion_Default___.Fm24(_802_v45, globalState)
			_ = _rhs105
			var _rhs106 _dafny.Int = (Companion_Default___.Fm0(_895_v109, (_this).F17(), (_this).F18(), _802_v45, globalState)).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(892))).Uint32(), func(coer66 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_896_i9 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			})), _dafny.UnicodeSeqOfUtf8Bytes("e"))).Cardinality()))
			_ = _rhs106
			var _lhs93 *C0 = _893_v107
			_ = _lhs93
			_lhs93.F19 = _rhs103
			_802_v45 = _rhs104
			_894_v108 = _rhs105
			r0 = _rhs106
			(_893_v107).F19 = _893_v107.F19
			_889_v105 = (_889_v105).Update(_802_v45, (_this).F17())
		} else {
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index121
			((_this).F14()).ArraySet1(p0, (_index121).Int())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index122
			((_this).F14()).ArraySet1(p0, (_index122).Int())
			var _897_v110 _dafny.CodePoint
			_ = _897_v110
			_897_v110 = _dafny.CodePoint('w')
			var _898_v111 *C0
			_ = _898_v111
			var _nw115 *C0 = New_C0_()
			_ = _nw115
			_nw115.Ctor__((!(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(83), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))) || ((_this).F18()), _897_v110, _this.F12(), (_this).F18())
			_898_v111 = _nw115
			var _899_v112 _dafny.MultiSet
			_ = _899_v112
			_899_v112 = _dafny.MultiSetOf(true)
			var _900_v113 _dafny.Sequence
			_ = _900_v113
			_900_v113 = _dafny.SeqOf(_899_v112)
			_900_v113 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(368))).Uint32(), func(coer67 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg67 _dafny.Int) interface{} {
					return coer67(arg67)
				}
			}((func(_901_v112 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_902_i10 _dafny.Int) _dafny.MultiSet {
					return _901_v112
				}
			})(_899_v112)))
			var _903_v114 _dafny.MultiSet
			_ = _903_v114
			_903_v114 = _dafny.MultiSetOf(_802_v45, _802_v45)
			var _904_v115 _dafny.Sequence
			_ = _904_v115
			_904_v115 = _dafny.SeqOf((_this).F17(), (_this).F17())
			var _905_v116 _dafny.Map
			_ = _905_v116
			_905_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F17())
			var _906_v117 _dafny.Set
			_ = _906_v117
			_906_v117 = _dafny.SetOf(Companion_Default___.Fm17(globalState), (func() bool {
				if (_905_v116).Contains((_this).F17()) {
					return (_905_v116).Get((_this).F17()).(bool)
				}
				return true
			})())
			var _907_v118 _dafny.Map
			_ = _907_v118
			_907_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_898_v111.F19, _802_v45)
			var _908_v119 _dafny.Array
			_ = _908_v119
			var _nwElement0_27 _dafny.Int = _802_v45
			_ = _nwElement0_27
			var _nw116 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(26))
			_ = _nw116
			(_nw116).ArraySet1(_nwElement0_27, 0)
			(_nw116).ArraySet1((_dafny.Zero).Minus(_802_v45), 1)
			(_nw116).ArraySet1(_802_v45, 2)
			(_nw116).ArraySet1((func() _dafny.Int {
				if (_903_v114).Contains(_802_v45) {
					return (_903_v114).Multiplicity(_802_v45)
				}
				return (_dafny.MultiSetOf((_this).F13())).Cardinality()
			})(), 3)
			(_nw116).ArraySet1(_dafny.IntOfUint32((_904_v115).Cardinality()), 4)
			(_nw116).ArraySet1(_802_v45, 5)
			(_nw116).ArraySet1(_802_v45, 6)
			(_nw116).ArraySet1(_802_v45, 7)
			(_nw116).ArraySet1((_906_v117).Cardinality(), 8)
			(_nw116).ArraySet1(_802_v45, 9)
			(_nw116).ArraySet1(_802_v45, 10)
			(_nw116).ArraySet1(_802_v45, 11)
			(_nw116).ArraySet1(_dafny.IntOfInt64(-840), 12)
			(_nw116).ArraySet1(_802_v45, 13)
			(_nw116).ArraySet1((_907_v118).Cardinality(), 14)
			(_nw116).ArraySet1(_802_v45, 15)
			(_nw116).ArraySet1(_802_v45, 16)
			(_nw116).ArraySet1(_802_v45, 17)
			(_nw116).ArraySet1(_802_v45, 18)
			(_nw116).ArraySet1(_dafny.IntOfInt64(787), 19)
			(_nw116).ArraySet1(_802_v45, 20)
			(_nw116).ArraySet1(_802_v45, 21)
			(_nw116).ArraySet1(_802_v45, 22)
			(_nw116).ArraySet1(_802_v45, 23)
			(_nw116).ArraySet1(_802_v45, 24)
			(_nw116).ArraySet1(_802_v45, 25)
			_908_v119 = _nw116
			var _909_v120 _dafny.Array
			_ = _909_v120
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_18
			var _nw117 _dafny.Array
			_ = _nw117
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw117 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) bool = func(_910_i11 _dafny.Int) bool {
					return (_this).F13()
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw117 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw117).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw117).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_909_v120 = _nw117
			var _911_v121 D2
			_ = _911_v121
			_911_v121 = Companion_D2_.Create_DC4_(_908_v119, (_this).F17(), (_dafny.Zero).Minus(_802_v45), _898_v111.F19, _909_v120)
			(_898_v111).F19 = (func() bool {
				if (p0) && (_898_v111.F19) {
					return (_911_v121).Dtor_cf7()
				}
				return true
			})()
			_802_v45 = (_dafny.Zero).Minus((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_802_v45, _802_v45)))
		}
		if p0 {
			var _912_v122 _dafny.Map
			_ = _912_v122
			_912_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
			r0 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F18() {
					return _802_v45
				}
				return Companion_Default___.SafeDivisionInt((_912_v122).Cardinality(), _802_v45)
			})())
			var _913_v123 _dafny.Map
			_ = _913_v123
			_913_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, (_this).F17())
			var _914_v125 _dafny.Sequence
			_ = _914_v125
			_914_v125 = _dafny.SeqOf(_913_v123, (_913_v123).Merge(_913_v123), _913_v123, (func() _dafny.Map {
				var _coll41 = _dafny.NewMapBuilder()
				_ = _coll41
				for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(627), _dafny.IntOfInt64(879))); ; {
					_compr_41, _ok49 := _iter49()
					if !_ok49 {
						break
					}
					var _915_v124 _dafny.Int
					_915_v124 = interface{}(_compr_41).(_dafny.Int)
					if ((_dafny.IntOfInt64(627)).Cmp(_915_v124) <= 0) && ((_915_v124).Cmp(_dafny.IntOfInt64(879)) < 0) {
						_coll41.Add(Companion_Default___.SafeDivisionInt(_915_v124, _802_v45), !((_this).F13()))
					}
				}
				return _coll41.ToMap()
			}()).Update(_802_v45, (_this).F18()))
			var _916_v126 _dafny.Sequence
			_ = _916_v126
			_916_v126 = _dafny.UnicodeSeqOfUtf8Bytes("fonto")
			_914_v125 = Companion_Default___.Fm25(_916_v126, p0, (_dafny.Zero).Minus(_802_v45), globalState)
			var _917_v127 _dafny.CodePoint
			_ = _917_v127
			_917_v127 = _dafny.CodePoint('a')
			var _pat_let_tv27 = _802_v45
			_ = _pat_let_tv27
			var _pat_let_tv28 = _802_v45
			_ = _pat_let_tv28
			var _pat_let_tv29 = _802_v45
			_ = _pat_let_tv29
			var _918_v128 *C0
			_ = _918_v128
			var _nw118 *C0 = New_C0_()
			_ = _nw118
			_nw118.Ctor__(p0, _917_v127, func(_pat_let23_0 D1) D1 {
				return func(_921_dt__update__tmp_h4 D1) D1 {
					return func(_pat_let26_0 _dafny.Int) D1 {
						return func(_922_dt__update_hcf3_h0 _dafny.Int) D1 {
							return func(_pat_let27_0 _dafny.Int) D1 {
								return func(_923_dt__update_hcf2_h2 _dafny.Int) D1 {
									return Companion_D1_.Create_DC2_(_923_dt__update_hcf2_h2, _922_dt__update_hcf3_h0, (_921_dt__update__tmp_h4).Dtor_cf4())
								}(_pat_let27_0)
							}(_dafny.IntOfInt64(-76))
						}(_pat_let26_0)
					}((Companion_D1_.Create_DC2_(_pat_let_tv28, _pat_let_tv29, (_this).F13())).Dtor_cf3())
				}(_pat_let23_0)
			}(func(_pat_let24_0 D1) D1 {
				return func(_919_dt__update__tmp_h3 D1) D1 {
					return func(_pat_let25_0 _dafny.Int) D1 {
						return func(_920_dt__update_hcf2_h1 _dafny.Int) D1 {
							return Companion_D1_.Create_DC2_(_920_dt__update_hcf2_h1, (_919_dt__update__tmp_h3).Dtor_cf3(), (_919_dt__update__tmp_h3).Dtor_cf4())
						}(_pat_let25_0)
					}(_pat_let_tv27)
				}(_pat_let24_0)
			}(_this.F12())), (_this).F13())
			_918_v128 = _nw118
			var _924_v129 _dafny.Sequence
			_ = _924_v129
			_924_v129 = _dafny.SeqOf(!(p0))
			var _925_v130 _dafny.Array
			_ = _925_v130
			var _nwElement0_28 _dafny.Int = (func() _dafny.Int {
				if p0 {
					return _802_v45
				}
				return _dafny.IntOfUint32((_924_v129).Cardinality())
			})()
			_ = _nwElement0_28
			var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(18))
			_ = _nw119
			(_nw119).ArraySet1(_nwElement0_28, 0)
			(_nw119).ArraySet1(_dafny.IntOfInt64(923), 1)
			(_nw119).ArraySet1(_dafny.IntOfUint32((_916_v126).Cardinality()), 2)
			(_nw119).ArraySet1(_dafny.IntOfInt64(581), 3)
			(_nw119).ArraySet1(Companion_Default___.SafeModuloInt(_802_v45, _802_v45), 4)
			(_nw119).ArraySet1(_802_v45, 5)
			(_nw119).ArraySet1(_802_v45, 6)
			(_nw119).ArraySet1((Companion_Default___.Fm26(_802_v45, _802_v45, p0, globalState)).Cardinality(), 7)
			(_nw119).ArraySet1((func() _dafny.Int {
				if p0 {
					return _802_v45
				}
				return _802_v45
			})(), 8)
			(_nw119).ArraySet1(_802_v45, 9)
			(_nw119).ArraySet1(_802_v45, 10)
			(_nw119).ArraySet1(_802_v45, 11)
			(_nw119).ArraySet1(_802_v45, 12)
			(_nw119).ArraySet1(_802_v45, 13)
			(_nw119).ArraySet1(_802_v45, 14)
			(_nw119).ArraySet1(_802_v45, 15)
			(_nw119).ArraySet1((_802_v45).Times(_802_v45), 16)
			(_nw119).ArraySet1(_802_v45, 17)
			_925_v130 = _nw119
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_925_v130), 0))
			_ = _index123
			(_925_v130).ArraySet1(_802_v45, (_index123).Int())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_925_v130), 0))
			_ = _index124
			(_925_v130).ArraySet1(_802_v45, (_index124).Int())
			var _926_v131 _dafny.Array
			_ = _926_v131
			var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(24))
			_ = _nw120
			_926_v131 = _nw120
			var _927_v132 _dafny.Map
			_ = _927_v132
			_927_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, (_925_v130).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_925_v130), 0))).Int()).(_dafny.Int))
			var _928_v133 _dafny.Sequence
			_ = _928_v133
			_928_v133 = _dafny.SeqOf((_927_v132).Cardinality())
			var _929_v134 _dafny.Set
			_ = _929_v134
			_929_v134 = _dafny.SetOf(_928_v133)
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_926_v131), 0))
			_ = _index125
			(_926_v131).ArraySet1((_929_v134).Union(_929_v134), (_index125).Int())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(702), _dafny.ArrayLen((_926_v131), 0))
			_ = _index126
			(_926_v131).ArraySet1(_929_v134, (_index126).Int())
		} else {
			var _930_v135 _dafny.Array
			_ = _930_v135
			var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw121
			_930_v135 = _nw121
			var _931_v136 _dafny.Array
			_ = _931_v136
			var _nwElement0_29 D1 = _this.F12()
			_ = _nwElement0_29
			var _nw122 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.One)
			_ = _nw122
			(_nw122).ArraySet1(_nwElement0_29, 0)
			_931_v136 = _nw122
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_930_v135), 0))
			_ = _index127
			(_930_v135).ArraySet1(_931_v136, (_index127).Int())
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_930_v135), 0))
			_ = _index128
			(_930_v135).ArraySet1(_931_v136, (_index128).Int())
			var _932_v137 D6
			_ = _932_v137
			_932_v137 = Companion_D6_.Create_DC12_(p0, Companion_Default___.Fm20(globalState), _802_v45, (_this).F13())
			var _source16 D6 = _932_v137
			_ = _source16
			if _source16.Is_DC12() {
				var _933___mcc_h6 bool = _source16.Get_().(D6_DC12).Cf18
				_ = _933___mcc_h6
				var _934___mcc_h7 _dafny.Sequence = _source16.Get_().(D6_DC12).Cf19
				_ = _934___mcc_h7
				var _935___mcc_h8 _dafny.Int = _source16.Get_().(D6_DC12).Cf20
				_ = _935___mcc_h8
				var _936___mcc_h9 bool = _source16.Get_().(D6_DC12).Cf21
				_ = _936___mcc_h9
				var _937_cf21 bool = _936___mcc_h9
				_ = _937_cf21
				var _938_cf20 _dafny.Int = _935___mcc_h8
				_ = _938_cf20
				var _939_cf19 _dafny.Sequence = _934___mcc_h7
				_ = _939_cf19
				var _940_cf18 bool = _933___mcc_h6
				_ = _940_cf18
				var _941_v138 _dafny.Array
				_ = _941_v138
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_19
				var _nw123 _dafny.Array
				_ = _nw123
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw123 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Sequence = (func(_942_cf19 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_943_i12 _dafny.Int) _dafny.Sequence {
							return _942_cf19
						}
					})(_939_cf19)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw123 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw123).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw123).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_941_v138 = _nw123
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_941_v138), 0))
				_ = _index129
				(_941_v138).ArraySet1(_939_cf19, (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(804), _dafny.ArrayLen((_941_v138), 0))
				_ = _index130
				(_941_v138).ArraySet1(_939_cf19, (_index130).Int())
				(_this).M12(globalState)
				var _944_v139 _dafny.CodePoint
				_ = _944_v139
				_944_v139 = _dafny.CodePoint('b')
				var _945_v140 *C0
				_ = _945_v140
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__((_this).F18(), _944_v139, _this.F12(), (_802_v45).Cmp(_802_v45) < 0)
				_945_v140 = _nw124
				var _946_v141 *C0
				_ = _946_v141
				var _nw125 *C0 = New_C0_()
				_ = _nw125
				_nw125.Ctor__(_940_cf18, _944_v139, _this.F12(), _940_cf18)
				_946_v141 = _nw125
			} else if _source16.Is_DC13() {
				var _947___mcc_h10 bool = _source16.Get_().(D6_DC13).Cf22
				_ = _947___mcc_h10
				var _948___mcc_h11 bool = _source16.Get_().(D6_DC13).Cf23
				_ = _948___mcc_h11
				var _949___mcc_h12 _dafny.Int = _source16.Get_().(D6_DC13).Cf24
				_ = _949___mcc_h12
				var _950___mcc_h13 _dafny.Int = _source16.Get_().(D6_DC13).Cf25
				_ = _950___mcc_h13
				var _951___mcc_h14 bool = _source16.Get_().(D6_DC13).Cf26
				_ = _951___mcc_h14
				var _952_cf26 bool = _951___mcc_h14
				_ = _952_cf26
				var _953_cf25 _dafny.Int = _950___mcc_h13
				_ = _953_cf25
				var _954_cf24 _dafny.Int = _949___mcc_h12
				_ = _954_cf24
				var _955_cf23 bool = _948___mcc_h11
				_ = _955_cf23
				var _956_cf22 bool = _947___mcc_h10
				_ = _956_cf22
				var _957_v142 _dafny.Array
				_ = _957_v142
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
				_ = _nw126
				_957_v142 = _nw126
				_957_v142 = _957_v142
				var _958_v143 _dafny.Sequence
				_ = _958_v143
				_958_v143 = _dafny.SeqOf(_954_cf24)
				_958_v143 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_802_v45, _953_cf25), _958_v143)
				var _959_v144 _dafny.CodePoint
				_ = _959_v144
				_959_v144 = _dafny.CodePoint('t')
				var _960_v145 _dafny.Sequence
				_ = _960_v145
				_960_v145 = _dafny.UnicodeSeqOfUtf8Bytes("tpil")
				_952_cf26 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_960_v145, _960_v145), _959_v144)
				_952_cf26 = (func() bool {
					if !((_this).F17()) {
						return (_this).F18()
					}
					return (_802_v45).Cmp(_802_v45) <= 0
				})()
			} else if _source16.Is_DC11() {
				var _961___mcc_h15 _dafny.MultiSet = _source16.Get_().(D6_DC11).Cf17
				_ = _961___mcc_h15
				var _962_cf17 _dafny.MultiSet = _961___mcc_h15
				_ = _962_cf17
				var _963_v146 _dafny.Array
				_ = _963_v146
				var _nw127 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw127
				_963_v146 = _nw127
				var _964_v147 _dafny.Array
				_ = _964_v147
				var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw128
				_964_v147 = _nw128
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_963_v146), 0))
				_ = _index131
				(_963_v146).ArraySet1(_964_v147, (_index131).Int())
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_963_v146), 0))
				_ = _index132
				(_963_v146).ArraySet1(_964_v147, (_index132).Int())
				var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_964_v147), 0))
				_ = _index133
				(_964_v147).ArraySet1(_802_v45, (_index133).Int())
				var _965_v148 _dafny.Array
				_ = _965_v148
				var _nw129 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(20))
				_ = _nw129
				_965_v148 = _nw129
				var _966_v149 _dafny.Sequence
				_ = _966_v149
				_966_v149 = _dafny.SeqOf(_965_v148)
				var _967_v150 _dafny.Sequence
				_ = _967_v150
				_967_v150 = _dafny.SeqOf(_966_v149, _966_v149, _966_v149)
				var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_964_v147), 0))
				_ = _index134
				(_964_v147).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_802_v45, _dafny.IntOfUint32(((_967_v150).Select((Companion_Default___.SafeIndex(_802_v45, _dafny.IntOfUint32((_967_v150).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())), _802_v45), (_index134).Int())
				var _968_v151 _dafny.Map
				_ = _968_v151
				_968_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _969_v152 _dafny.Map
				_ = _969_v152
				_969_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_964_v147).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(933), _dafny.ArrayLen((_964_v147), 0))).Int()).(_dafny.Int))
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index135
				((_this).F14()).ArraySet1(((_dafny.Zero).Minus((_968_v151).Cardinality())).Cmp((func() _dafny.Int {
					if (_969_v152).Contains((_this).F18()) {
						return (_969_v152).Get((_this).F18()).(_dafny.Int)
					}
					return _802_v45
				})()) > 0, (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index136
				((_this).F14()).ArraySet1(p0, (_index136).Int())
				var _970_v153 _dafny.Sequence
				_ = _970_v153
				_970_v153 = _dafny.UnicodeSeqOfUtf8Bytes("x")
				var _971_v154 _dafny.Map
				_ = _971_v154
				_971_v154 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_970_v153, _964_v147)
				(globalState).F7 = ((_971_v154).Update(_dafny.Companion_Sequence_.Concatenate(_970_v153, _970_v153), _dafny.ArrayCastTo((_963_v146).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(90), _dafny.ArrayLen((_963_v146), 0))).Int())))).Cardinality()
			} else {
				var _972___mcc_h16 D6 = _source16.Get_().(D6_DC14).Cf27
				_ = _972___mcc_h16
				var _973_cf27 D6 = _972___mcc_h16
				_ = _973_cf27
				(globalState).F7 = _802_v45
				var _974_v155 _dafny.Sequence
				_ = _974_v155
				_974_v155 = _dafny.SeqOf((_this).F14(), (_this).F14())
				var _975_v156 _dafny.Sequence
				_ = _975_v156
				_975_v156 = _dafny.UnicodeSeqOfUtf8Bytes("xwoa")
				var _976_v157 _dafny.Map
				_ = _976_v157
				_976_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_974_v155).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_975_v156).Cardinality()), _dafny.IntOfUint32((_974_v155).Cardinality()))).Uint32()).(_dafny.Array), _dafny.SeqOf((_dafny.Zero).Minus(_802_v45), _802_v45))
				_976_v157 = (_976_v157).Update((_this).F14(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(315))).Uint32(), func(coer68 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_977_i13 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(447)
				})))
				var _978_v158 bool
				_ = _978_v158
				_978_v158 = true
				var _979_v159 _dafny.Array
				_ = _979_v159
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_20
				var _nw130 _dafny.Array
				_ = _nw130
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw130 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) _dafny.Sequence = (func(_980_v45 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_981_i14 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_980_v45)
						}
					})(_802_v45)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw130 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw130).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw130).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_979_v159 = _nw130
				var _982_v160 _dafny.Sequence
				_ = _982_v160
				_982_v160 = _dafny.SeqOf(_802_v45)
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_979_v159), 0))
				_ = _index137
				(_979_v159).ArraySet1(_982_v160, (_index137).Int())
				var _983_v161 _dafny.Sequence
				_ = _983_v161
				_983_v161 = _dafny.SeqOf(Companion_Default___.Fm17(globalState))
				var _984_v162 _dafny.Map
				_ = _984_v162
				_984_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, p0)
				var _985_v163 _dafny.MultiSet
				_ = _985_v163
				_985_v163 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_983_v161, (Companion_Default___.SafeIndex((_984_v162).Cardinality(), _dafny.IntOfUint32((_983_v161).Cardinality()))).Uint32(), true))
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_979_v159), 0))
				_ = _index138
				var _rhs107 bool = ((_this).F17()) && (!(p0))
				_ = _rhs107
				var _rhs108 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_802_v45, p0)).Cardinality(), (_985_v163).Cardinality(), _802_v45, _802_v45), _dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(globalState), (Companion_Default___.SafeIndex(_802_v45, _dafny.IntOfUint32((Companion_Default___.Fm27(globalState)).Cardinality()))).Uint32(), _802_v45))
				_ = _rhs108
				var _rhs109 _dafny.Int = ((_802_v45).Minus(_802_v45)).Minus(Companion_Default___.SafeDivisionInt(_802_v45, _802_v45))
				_ = _rhs109
				var _lhs94 _dafny.Array = _979_v159
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(565), _dafny.ArrayLen((_979_v159), 0))
				_ = _lhs95
				_978_v158 = _rhs107
				(_lhs94).ArraySet1(_rhs108, (_lhs95).Int())
				_802_v45 = _rhs109
				var _986_v164 *C0
				_ = _986_v164
				var _nw131 *C0 = New_C0_()
				_ = _nw131
				_nw131.Ctor__(p0, _dafny.CodePoint('n'), _this.F12(), _978_v158)
				_986_v164 = _nw131
				var _987_v165 _dafny.Map
				_ = _987_v165
				_987_v165 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_986_v164, _975_v156)
				_987_v165 = (_987_v165).Update(_986_v164, _dafny.Companion_Sequence_.Concatenate(_975_v156, _dafny.UnicodeSeqOfUtf8Bytes("e")))
			}
			var _988_v166 _dafny.MultiSet
			_ = _988_v166
			_988_v166 = _dafny.MultiSetOf((_this).F13(), (_this).F18())
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index139
			((_this).F14()).ArraySet1((_this).F17(), (_index139).Int())
			var _989_v167 bool
			_ = _989_v167
			_989_v167 = false
			var _990_v168 _dafny.CodePoint
			_ = _990_v168
			_990_v168 = _dafny.CodePoint('b')
			var _991_v169 _dafny.Map
			_ = _991_v169
			_991_v169 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_990_v168, (_this).F18())
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index140
			var _rhs110 _dafny.MultiSet = Companion_Default___.Fm28((func() bool {
				if (_991_v169).Contains(_990_v168) {
					return (_991_v169).Get(_990_v168).(bool)
				}
				return (_this).F18()
			})(), (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ybya")).Cardinality()), _802_v45)), !((_this).F17()) || (!(p0)), globalState)
			_ = _rhs110
			var _rhs111 bool = (_this).F13()
			_ = _rhs111
			var _rhs112 bool = Companion_Default___.Fm17(globalState)
			_ = _rhs112
			var _rhs113 bool = ((_802_v45).Plus(_802_v45)).Cmp(_802_v45) < 0
			_ = _rhs113
			var _rhs114 bool = (_this).F13()
			_ = _rhs114
			var _lhs96 _dafny.Array = (_this).F14()
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs97
			_988_v166 = _rhs110
			(_lhs96).ArraySet1(_rhs111, (_lhs97).Int())
			_989_v167 = _rhs112
			_989_v167 = _rhs113
			_989_v167 = _rhs114
			if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
				var _992_v170 _dafny.Array
				_ = _992_v170
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(25)
				_ = _len0_21
				var _nw132 _dafny.Array
				_ = _nw132
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw132 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_993_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_994_i15 _dafny.Int) _dafny.Int {
							return (_994_i15).Plus(_993_v45)
						}
					})(_802_v45)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw132 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw132).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw132).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_992_v170 = _nw132
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_992_v170), 0))
				_ = _index141
				(_992_v170).ArraySet1(((_this).Fm15(globalState)).Minus((_this).Fm9(_990_v168, globalState)), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_992_v170), 0))
				_ = _index142
				(_992_v170).ArraySet1(_802_v45, (_index142).Int())
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index143
				((_this).F14()).ArraySet1(_989_v167, (_index143).Int())
				var _995_v171 _dafny.Sequence
				_ = _995_v171
				_995_v171 = _dafny.UnicodeSeqOfUtf8Bytes("nsq")
				_995_v171 = _dafny.Companion_Sequence_.Update(_995_v171, (Companion_Default___.SafeIndex(_802_v45, _dafny.IntOfUint32((_995_v171).Cardinality()))).Uint32(), _990_v168)
				var _996_v172 D6
				_ = _996_v172
				_996_v172 = Companion_D6_.Create_DC13_(p0, p0, _802_v45, _802_v45, Companion_Default___.Fm17(globalState))
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index144
				var _rhs115 bool = (func() bool {
					if p0 {
						return _989_v167
					}
					return ((_992_v170).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(362), _dafny.ArrayLen((_992_v170), 0))).Int()).(_dafny.Int)).Cmp(_802_v45) < 0
				})()
				_ = _rhs115
				var _rhs116 _dafny.Int = ((_dafny.Zero).Minus((_dafny.Zero).Minus(((_dafny.SetOf(_992_v170, _992_v170)).Cardinality()).Times(_802_v45)))).Plus((_996_v172).Dtor_cf25())
				_ = _rhs116
				var _lhs98 _dafny.Array = (_this).F14()
				_ = _lhs98
				var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs99
				var _lhs100 *GlobalState = globalState
				_ = _lhs100
				(_lhs98).ArraySet1(_rhs115, (_lhs99).Int())
				_lhs100.F7 = _rhs116
				var _997_v173 _dafny.Array
				_ = _997_v173
				var _nw133 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(23))
				_ = _nw133
				_997_v173 = _nw133
				var _998_v174 _dafny.Map
				_ = _998_v174
				_998_v174 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v167, _997_v173)
				_998_v174 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _997_v173)).Merge(_998_v174)).Merge(_998_v174)
			} else {
				var _999_v175 _dafny.Sequence
				_ = _999_v175
				_999_v175 = _dafny.UnicodeSeqOfUtf8Bytes("cmyb")
				var _pat_let_tv30 = _999_v175
				_ = _pat_let_tv30
				var _1000_v176 _dafny.Map
				_ = _1000_v176
				_1000_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm29((func(_pat_let28_0 D6) D6 {
					return func(_1001_dt__update__tmp_h5 D6) D6 {
						return func(_pat_let29_0 _dafny.Sequence) D6 {
							return func(_1002_dt__update_hcf19_h0 _dafny.Sequence) D6 {
								return func(_pat_let30_0 bool) D6 {
									return func(_1003_dt__update_hcf21_h0 bool) D6 {
										return Companion_D6_.Create_DC12_((_1001_dt__update__tmp_h5).Dtor_cf18(), _1002_dt__update_hcf19_h0, (_1001_dt__update__tmp_h5).Dtor_cf20(), _1003_dt__update_hcf21_h0)
									}(_pat_let30_0)
								}((_this).F17())
							}(_pat_let29_0)
						}(_pat_let_tv30)
					}(_pat_let28_0)
				}(_932_v137)).Dtor_cf20(), globalState), _802_v45)
				var _1004_v177 _dafny.Sequence
				_ = _1004_v177
				_1004_v177 = _dafny.SeqOf(_989_v167, !(p0), true, !(p0))
				_1000_v176 = (_1000_v176).Update(_1004_v177, _802_v45)
				var _1005_v178 _dafny.Map
				_ = _1005_v178
				_1005_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_989_v167, _999_v175)
				var _1006_v179 _dafny.Array
				_ = _1006_v179
				var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_999_v175, _999_v175)
				_ = _nwElement0_30
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(13))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_30, 0)
				(_nw134).ArraySet1(_999_v175, 1)
				(_nw134).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_999_v175, _999_v175), 2)
				(_nw134).ArraySet1(_999_v175, 3)
				(_nw134).ArraySet1(_999_v175, 4)
				(_nw134).ArraySet1(_999_v175, 5)
				(_nw134).ArraySet1(_999_v175, 6)
				(_nw134).ArraySet1(_999_v175, 7)
				(_nw134).ArraySet1((func() _dafny.Sequence {
					if (_1005_v178).Contains(p0) {
						return (_1005_v178).Get(p0).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(44))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg69 _dafny.Int) interface{} {
							return coer69(arg69)
						}
					}((func(_1007_v168 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1008_i16 _dafny.Int) _dafny.CodePoint {
							return _1007_v168
						}
					})(_990_v168)))
				})(), 8)
				(_nw134).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_999_v175, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(104))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1009_v168 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1010_i17 _dafny.Int) _dafny.CodePoint {
						return _1009_v168
					}
				})(_990_v168)))), 9)
				(_nw134).ArraySet1((func() _dafny.Sequence {
					if (_this).F13() {
						return _999_v175
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("plvqc")
				})(), 10)
				(_nw134).ArraySet1(_999_v175, 11)
				(_nw134).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_999_v175, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-228))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg71 _dafny.Int) interface{} {
						return coer71(arg71)
					}
				}((func(_1011_v168 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1012_i18 _dafny.Int) _dafny.CodePoint {
						return _1011_v168
					}
				})(_990_v168)))), 12)
				_1006_v179 = _nw134
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1006_v179), 0))
				_ = _index145
				(_1006_v179).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yhhsck"), (_index145).Int())
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1006_v179), 0))
				_ = _index146
				var _rhs117 _dafny.CodePoint = _990_v168
				_ = _rhs117
				var _rhs118 _dafny.Sequence = _999_v175
				_ = _rhs118
				var _rhs119 _dafny.Sequence = _999_v175
				_ = _rhs119
				var _lhs101 _dafny.Array = _1006_v179
				_ = _lhs101
				var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1006_v179), 0))
				_ = _lhs102
				_990_v168 = _rhs117
				_999_v175 = _rhs118
				(_lhs101).ArraySet1(_rhs119, (_lhs102).Int())
				(globalState).F7 = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32(((_1006_v179).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(367), _dafny.ArrayLen((_1006_v179), 0))).Int()).(_dafny.Sequence)).Cardinality())).Times(_802_v45), _802_v45)
				_989_v167 = _989_v167
				_989_v167 = false
			}
			var _1013_v180 _dafny.MultiSet
			_ = _1013_v180
			_1013_v180 = _dafny.MultiSetOf(_802_v45, _dafny.IntOfInt64(882), _802_v45, _802_v45)
			var _1014_v181 _dafny.Sequence
			_ = _1014_v181
			_1014_v181 = _dafny.SeqOf(_802_v45, (func() _dafny.Int {
				if (_1013_v180).Contains(_802_v45) {
					return (_1013_v180).Multiplicity(_802_v45)
				}
				return _802_v45
			})(), _802_v45)
			_1014_v181 = _dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(globalState), (Companion_Default___.SafeIndex((_1014_v181).Select((Companion_Default___.SafeIndex(_802_v45, _dafny.IntOfUint32((_1014_v181).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((Companion_Default___.Fm27(globalState)).Cardinality()))).Uint32(), _802_v45)
		}
		var _1015_v182 _dafny.Sequence
		_ = _1015_v182
		_1015_v182 = _dafny.SeqOf((_this).F17())
		r0 = Companion_Default___.Fm0(_1015_v182, (_this).F13(), (_this).F13(), _802_v45, globalState)
		return r0
	}
}
func (_this *C3) M7(globalState *GlobalState) {
	{
		if (_this).F17() {
			var _1016_v0 _dafny.Int
			_ = _1016_v0
			_1016_v0 = _dafny.IntOfInt64(554)
			var _1017_v1 D6
			_ = _1017_v1
			_1017_v1 = Companion_D6_.Create_DC13_(!((_this).F17()), !((_dafny.IntOfInt64(272)).Cmp(_dafny.IntOfInt64(383)) >= 0), _1016_v0, _1016_v0, (_this).F17())
			_1017_v1 = _1017_v1
			var _1018_v2 _dafny.Sequence
			_ = _1018_v2
			_1018_v2 = _dafny.SeqOf(((_this).Fm10((_this.F12()).Dtor_cf4(), Companion_Default___.Fm17(globalState), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F13(), (_this).F13()), (Companion_Default___.SafeIndex(_1016_v0, _dafny.IntOfUint32((_dafny.SeqOf((_this).F13(), (_this).F13())).Cardinality()))).Uint32(), (_this).F13()), globalState)).Cardinality())
			var _1019_v3 _dafny.Sequence
			_ = _1019_v3
			_1019_v3 = _dafny.SeqOf((_this).F18())
			var _1020_v4 _dafny.Sequence
			_ = _1020_v4
			_1020_v4 = _dafny.UnicodeSeqOfUtf8Bytes("xdavfw")
			var _1021_v5 _dafny.CodePoint
			_ = _1021_v5
			_1021_v5 = _dafny.CodePoint('e')
			var _1022_v6 _dafny.Map
			_ = _1022_v6
			_1022_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1020_v4, _1021_v5)
			var _1023_v7 _dafny.Array
			_ = _1023_v7
			var _nwElement0_31 _dafny.Int = (func() _dafny.Int {
				if (_this).F17() {
					return _1016_v0
				}
				return _1016_v0
			})()
			_ = _nwElement0_31
			var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(17))
			_ = _nw135
			(_nw135).ArraySet1(_nwElement0_31, 0)
			(_nw135).ArraySet1(_1016_v0, 1)
			(_nw135).ArraySet1(_1016_v0, 2)
			(_nw135).ArraySet1(_1016_v0, 3)
			(_nw135).ArraySet1(_1016_v0, 4)
			(_nw135).ArraySet1(_dafny.IntOfUint32((_1018_v2).Cardinality()), 5)
			(_nw135).ArraySet1(_1016_v0, 6)
			(_nw135).ArraySet1((_1016_v0).Plus(_1016_v0), 7)
			(_nw135).ArraySet1(_dafny.IntOfUint32((_1019_v3).Cardinality()), 8)
			(_nw135).ArraySet1((_1018_v2).Select((Companion_Default___.SafeIndex(_1016_v0, _dafny.IntOfUint32((_1018_v2).Cardinality()))).Uint32()).(_dafny.Int), 9)
			(_nw135).ArraySet1(_1016_v0, 10)
			(_nw135).ArraySet1(_1016_v0, 11)
			(_nw135).ArraySet1(_dafny.IntOfInt64(-137), 12)
			(_nw135).ArraySet1((_1022_v6).Cardinality(), 13)
			(_nw135).ArraySet1(_1016_v0, 14)
			(_nw135).ArraySet1(_dafny.IntOfInt64(440), 15)
			(_nw135).ArraySet1((_1016_v0).Minus(_1016_v0), 16)
			_1023_v7 = _nw135
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1023_v7), 0))
			_ = _index147
			(_1023_v7).ArraySet1(_dafny.IntOfUint32(((Companion_D1_.Create_DC1_(_1020_v4)).Dtor_cf1()).Cardinality()), (_index147).Int())
			var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1023_v7), 0))
			_ = _index148
			(_1023_v7).ArraySet1(_1016_v0, (_index148).Int())
			var _1024_v8 _dafny.MultiSet
			_ = _1024_v8
			_1024_v8 = _dafny.MultiSetOf(_1021_v5, _1021_v5, _1021_v5)
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1023_v7), 0))
			_ = _index149
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1023_v7), 0))
			_ = _index150
			var _rhs120 _dafny.Int = _1016_v0
			_ = _rhs120
			var _rhs121 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1020_v4, _1020_v4)).Cardinality())).Plus(_1016_v0)
			_ = _rhs121
			var _rhs122 _dafny.MultiSet = (_dafny.MultiSetOf(_1021_v5)).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_1020_v4, _dafny.SeqOf(_1021_v5, _1021_v5, _1021_v5))))
			_ = _rhs122
			var _lhs103 _dafny.Array = _1023_v7
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1023_v7), 0))
			_ = _lhs104
			var _lhs105 _dafny.Array = _1023_v7
			_ = _lhs105
			var _lhs106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(594), _dafny.ArrayLen((_1023_v7), 0))
			_ = _lhs106
			(_lhs103).ArraySet1(_rhs120, (_lhs104).Int())
			(_lhs105).ArraySet1(_rhs121, (_lhs106).Int())
			_1024_v8 = _rhs122
			(globalState).F7 = (_1023_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1023_v7), 0))).Int()).(_dafny.Int)
			var _1025_v9 bool
			_ = _1025_v9
			_1025_v9 = true
			_1025_v9 = (_this).F13()
			var _1026_v10 _dafny.Map
			_ = _1026_v10
			_1026_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1023_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(674), _dafny.ArrayLen((_1023_v7), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1020_v4).Cardinality()))
			var _1027_v11 _dafny.Sequence
			_ = _1027_v11
			_1027_v11 = _dafny.SeqOf(_1020_v4, _1020_v4)
			_1026_v10 = (_1026_v10).Update(_dafny.IntOfInt64(-212), _dafny.IntOfUint32((_1027_v11).Cardinality()))
		} else {
			var _1028_v12 _dafny.Sequence
			_ = _1028_v12
			_1028_v12 = _dafny.SeqOf((_this).F13())
			var _1029_v13 _dafny.Array
			_ = _1029_v13
			var _nwElement0_32 bool = ((_this).F17()) == ((_this).F18())
			_ = _nwElement0_32
			var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(3))
			_ = _nw136
			(_nw136).ArraySet1(_nwElement0_32, 0)
			(_nw136).ArraySet1((_this).F13(), 1)
			(_nw136).ArraySet1((_1028_v12).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(562), _dafny.IntOfUint32((_1028_v12).Cardinality()))).Uint32()).(bool), 2)
			_1029_v13 = _nw136
			var _1030_v14 bool
			_ = _1030_v14
			_1030_v14 = false
			var _1031_v15 _dafny.Array
			_ = _1031_v15
			var _len0_22 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_22
			var _nw137 _dafny.Array
			_ = _nw137
			if _len0_22.Cmp(_dafny.Zero) == 0 {
				_nw137 = _dafny.NewArray(_len0_22)
			} else {
				var _init22 func(_dafny.Int) _dafny.CodePoint = func(_1032_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('w')
				}
				_ = _init22
				var _element0_22 = _init22(_dafny.Zero)
				_ = _element0_22
				_nw137 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
				(_nw137).ArraySet1CodePoint(_element0_22, 0)
				var _nativeLen0_22 = (_len0_22).Int()
				_ = _nativeLen0_22
				for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
					(_nw137).ArraySet1CodePoint(_init22(_dafny.IntOf(_i0_22)), _i0_22)
				}
			}
			_1031_v15 = _nw137
			var _1033_v16 _dafny.CodePoint
			_ = _1033_v16
			_1033_v16 = _dafny.CodePoint('c')
			var _1034_v17 _dafny.Sequence
			_ = _1034_v17
			_1034_v17 = _dafny.SeqOf(_1033_v16, _1033_v16)
			var _1035_v18 _dafny.Int
			_ = _1035_v18
			_1035_v18 = _dafny.IntOfInt64(677)
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_1031_v15), 0))
			_ = _index151
			(_1031_v15).ArraySet1CodePoint((_1034_v17).Select((Companion_Default___.SafeIndex(_1035_v18, _dafny.IntOfUint32((_1034_v17).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index151).Int())
			var _1036_v19 _dafny.Map
			_ = _1036_v19
			_1036_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(709), !((!(_1030_v14)) || ((_this).F18())))
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_1031_v15), 0))
			_ = _index152
			var _rhs123 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1030_v14), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1030_v14), Companion_Default___.Fm29(_1035_v18, globalState)))
			_ = _rhs123
			var _rhs124 _dafny.Array = _1029_v13
			_ = _rhs124
			var _rhs125 bool = (func() bool {
				if (_1036_v19).Contains(_1035_v18) {
					return (_1036_v19).Get(_1035_v18).(bool)
				}
				return (_this).F17()
			})()
			_ = _rhs125
			var _rhs126 _dafny.CodePoint = _1033_v16
			_ = _rhs126
			var _lhs107 *GlobalState = globalState
			_ = _lhs107
			var _lhs108 _dafny.Array = _1031_v15
			_ = _lhs108
			var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_1031_v15), 0))
			_ = _lhs109
			_lhs107.F2 = _rhs123
			_1029_v13 = _rhs124
			_1030_v14 = _rhs125
			(_lhs108).ArraySet1CodePoint(_rhs126, (_lhs109).Int())
			var _1037_v20 _dafny.Set
			_ = _1037_v20
			_1037_v20 = _dafny.SetOf((_this).F13())
			(globalState).F7 = (Companion_Default___.SafeModuloInt((_this).Fm9(_1033_v16, globalState), _dafny.IntOfInt64(757))).Times((_1037_v20).Cardinality())
			var _1038_v21 _dafny.MultiSet
			_ = _1038_v21
			_1038_v21 = _dafny.MultiSetOf(_1035_v18)
			var _1039_v22 _dafny.Sequence
			_ = _1039_v22
			_1039_v22 = _dafny.SeqOf(_1035_v18)
			var _1040_v23 _dafny.Array
			_ = _1040_v23
			var _nwElement0_33 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(733))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}((func(_1041_v16 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1042_i1 _dafny.Int) _dafny.CodePoint {
					return _1041_v16
				}
			})(_1033_v16)))).Cardinality())
			_ = _nwElement0_33
			var _nw138 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(6))
			_ = _nw138
			(_nw138).ArraySet1(_nwElement0_33, 0)
			(_nw138).ArraySet1(_dafny.IntOfUint32((_1034_v17).Cardinality()), 1)
			(_nw138).ArraySet1(_1035_v18, 2)
			(_nw138).ArraySet1(_1035_v18, 3)
			(_nw138).ArraySet1((_1038_v21).Cardinality(), 4)
			(_nw138).ArraySet1((_1039_v22).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1034_v17).Cardinality()), _dafny.IntOfUint32((_1039_v22).Cardinality()))).Uint32()).(_dafny.Int), 5)
			_1040_v23 = _nw138
			var _1043_v24 D2
			_ = _1043_v24
			_1043_v24 = Companion_D2_.Create_DC4_(_1040_v23, Companion_Default___.Fm17(globalState), (_1035_v18).Plus((_1038_v21).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1035_v18, _1035_v18)).Contains(_1035_v18), _1029_v13)
			var _source17 D2 = _1043_v24
			_ = _source17
			if _source17.Is_DC4() {
				var _1044___mcc_h0 _dafny.Array = _source17.Get_().(D2_DC4).Cf6
				_ = _1044___mcc_h0
				var _1045___mcc_h1 bool = _source17.Get_().(D2_DC4).Cf7
				_ = _1045___mcc_h1
				var _1046___mcc_h2 _dafny.Int = _source17.Get_().(D2_DC4).Cf8
				_ = _1046___mcc_h2
				var _1047___mcc_h3 bool = _source17.Get_().(D2_DC4).Cf9
				_ = _1047___mcc_h3
				var _1048___mcc_h4 _dafny.Array = _source17.Get_().(D2_DC4).Cf10
				_ = _1048___mcc_h4
				var _1049_cf10 _dafny.Array = _1048___mcc_h4
				_ = _1049_cf10
				var _1050_cf9 bool = _1047___mcc_h3
				_ = _1050_cf9
				var _1051_cf8 _dafny.Int = _1046___mcc_h2
				_ = _1051_cf8
				var _1052_cf7 bool = _1045___mcc_h1
				_ = _1052_cf7
				var _1053_cf6 _dafny.Array = _1044___mcc_h0
				_ = _1053_cf6
				(globalState).F7 = _1035_v18
				var _1054_v25 _dafny.Array
				_ = _1054_v25
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw139
				_1054_v25 = _nw139
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1054_v25), 0))
				_ = _index153
				(_1054_v25).ArraySet1(_1029_v13, (_index153).Int())
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1054_v25), 0))
				_ = _index154
				(_1054_v25).ArraySet1(_1049_cf10, (_index154).Int())
				var _1055_v26 _dafny.Sequence
				_ = _1055_v26
				_1055_v26 = _dafny.SeqOf(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer73 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg73 _dafny.Int) interface{} {
						return coer73(arg73)
					}
				}((func(_1056_v15 _dafny.Array) func(_dafny.Int) _dafny.CodePoint {
					return func(_1057_i2 _dafny.Int) _dafny.CodePoint {
						return (_1056_v15).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(921), _dafny.ArrayLen((_1056_v15), 0))).Int())
					}
				})(_1031_v15)))))
				var _1058_v27 D3
				_ = _1058_v27
				_1058_v27 = Companion_D3_.Create_DC7_()
				var _1059_v28 _dafny.Set
				_ = _1059_v28
				_1059_v28 = _dafny.SetOf(_1058_v27)
				var _rhs127 _dafny.Int = _1035_v18
				_ = _rhs127
				var _rhs128 _dafny.Int = ((_1035_v18).Plus(_dafny.IntOfUint32((_1055_v26).Cardinality()))).Plus(_1035_v18)
				_ = _rhs128
				var _rhs129 _dafny.Int = (_1059_v28).Cardinality()
				_ = _rhs129
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				_1051_cf8 = _rhs127
				_1051_cf8 = _rhs128
				_lhs110.F7 = _rhs129
				(globalState).F7 = _1035_v18
			} else {
				var _1060___mcc_h5 _dafny.Array = _source17.Get_().(D2_DC3).Cf5
				_ = _1060___mcc_h5
				var _1061_cf5 _dafny.Array = _1060___mcc_h5
				_ = _1061_cf5
				(_this).M11(_1029_v13, _1030_v14, (_dafny.IntOfUint32((_1028_v12).Cardinality())).Cmp(_1035_v18) >= 0, globalState)
				_1061_cf5 = (_this).F14()
				(globalState).F7 = (_1035_v18).Plus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_1035_v18), _1035_v18))
				_1061_cf5 = (_this).F14()
			}
			var _1062_v29 _dafny.Array
			_ = _1062_v29
			var _nw140 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw140
			_1062_v29 = _nw140
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_1040_v23), 0))
			_ = _index155
			(_1040_v23).ArraySet1(_1035_v18, (_index155).Int())
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_1040_v23), 0))
			_ = _index156
			var _rhs130 _dafny.Array = _1062_v29
			_ = _rhs130
			var _rhs131 _dafny.Array = (_this).F14()
			_ = _rhs131
			var _rhs132 _dafny.Int = _1035_v18
			_ = _rhs132
			var _rhs133 _dafny.Int = (_1039_v22).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_1035_v18, _1035_v18), _dafny.IntOfUint32((_1039_v22).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs133
			var _lhs111 _dafny.Array = _1040_v23
			_ = _lhs111
			var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_1040_v23), 0))
			_ = _lhs112
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			_1062_v29 = _rhs130
			_1029_v13 = _rhs131
			(_lhs111).ArraySet1(_rhs132, (_lhs112).Int())
			_lhs113.F7 = _rhs133
			_1030_v14 = (_this).F18()
		}
		var _1063_v30 bool
		_ = _1063_v30
		_1063_v30 = false
		_1063_v30 = true
		var _1064_v31 D6
		_ = _1064_v31
		_1064_v31 = Companion_D6_.Create_DC11_(Companion_Default___.Fm19(globalState))
		_1064_v31 = _1064_v31
		var _source18 D1 = _this.F12()
		_ = _source18
		if _source18.Is_DC2() {
			var _1065___mcc_h6 _dafny.Int = _source18.Get_().(D1_DC2).Cf2
			_ = _1065___mcc_h6
			var _1066___mcc_h7 _dafny.Int = _source18.Get_().(D1_DC2).Cf3
			_ = _1066___mcc_h7
			var _1067___mcc_h8 bool = _source18.Get_().(D1_DC2).Cf4
			_ = _1067___mcc_h8
			var _1068_cf4 bool = _1067___mcc_h8
			_ = _1068_cf4
			var _1069_cf3 _dafny.Int = _1066___mcc_h7
			_ = _1069_cf3
			var _1070_cf2 _dafny.Int = _1065___mcc_h6
			_ = _1070_cf2
			var _1071_v32 _dafny.CodePoint
			_ = _1071_v32
			_1071_v32 = _dafny.CodePoint('f')
			_1071_v32 = _1071_v32
			var _1072_v33 _dafny.Sequence
			_ = _1072_v33
			_1072_v33 = _dafny.SeqOf(_1069_cf3)
			_1063_v30 = !_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_1072_v33, _1072_v33), _1072_v33)
			var _1073_v34 T3
			_ = _1073_v34
			var _nw141 *C2 = New_C2_()
			_ = _nw141
			_nw141.Ctor__(_dafny.IntOfInt64(-855), _1068_cf4, _this.F16(), _1068_cf4, (_this).F14(), (_this).F15(), _this.F12(), _1063_v30)
			_1073_v34 = _nw141
			var _1074_v35 _dafny.Array
			_ = _1074_v35
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_23
			var _nw142 _dafny.Array
			_ = _nw142
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw142 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) _dafny.Int = (func(_1075_cf2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1076_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1076_i3, _1075_cf2)
					}
				})(_1070_cf2)
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw142 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw142).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw142).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1074_v35 = _nw142
			var _1077_v36 _dafny.Map
			_ = _1077_v36
			_1077_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T3 {
				if (_this).F17() {
					return _1073_v34
				}
				return _1073_v34
			})(), _1074_v35)
			(globalState).F7 = (_dafny.Zero).Minus((_1077_v36).Cardinality())
			var _1078_v37 _dafny.Sequence
			_ = _1078_v37
			_1078_v37 = _dafny.SeqOf((func() bool {
				if _1068_cf4 {
					return (_this).F13()
				}
				return _1068_cf4
			})())
			_1063_v30 = !(!((_1078_v37).Select((Companion_Default___.SafeIndex(_1070_cf2, _dafny.IntOfUint32((_1078_v37).Cardinality()))).Uint32()).(bool)))
		} else {
			var _1079___mcc_h9 _dafny.Sequence = _source18.Get_().(D1_DC1).Cf1
			_ = _1079___mcc_h9
			var _1080_cf1 _dafny.Sequence = _1079___mcc_h9
			_ = _1080_cf1
			_1063_v30 = (func() bool {
				if !(false) {
					return _1063_v30
				}
				return Companion_Default___.Fm30(_dafny.IntOfInt64(38), globalState)
			})()
			var _1081_v38 _dafny.CodePoint
			_ = _1081_v38
			_1081_v38 = _dafny.CodePoint('v')
			var _1082_v39 *C0
			_ = _1082_v39
			var _nw143 *C0 = New_C0_()
			_ = _nw143
			_nw143.Ctor__((_this).F17(), _1081_v38, _this.F12(), (func() bool {
				if (_this).F13() {
					return (_this).F13()
				}
				return _1063_v30
			})())
			_1082_v39 = _nw143
			var _1083_v40 _dafny.Int
			_ = _1083_v40
			_1083_v40 = _dafny.IntOfInt64(872)
			var _1084_v41 _dafny.Set
			_ = _1084_v41
			_1084_v41 = _dafny.SetOf((_this).F13(), !(!((_this).F13())))
			var _1085_v42 *C2
			_ = _1085_v42
			var _nw144 *C2 = New_C2_()
			_ = _nw144
			_nw144.Ctor__((_dafny.Zero).Minus(_1083_v40), (_dafny.SetOf((_this).F17(), !(false))).IsSubsetOf(_1084_v41), Companion_D1_.Create_DC1_(_1080_cf1), _dafny.Companion_Sequence_.IsPrefixOf(_1080_cf1, _1080_cf1), (_this).F14(), (_this).F15(), _this.F12(), _1063_v30)
			_1085_v42 = _nw144
			var _1086_v43 _dafny.Map
			_ = _1086_v43
			_1086_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1085_v42).F21(), _1063_v30)
			var _1087_v44 _dafny.Map
			_ = _1087_v44
			_1087_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1086_v43, (_1083_v40).Plus(_1083_v40))
			var _1088_v45 _dafny.Map
			_ = _1088_v45
			_1088_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F18(), (_this).F13())
			_1087_v44 = (_1087_v44).Update(_1086_v43, ((_1088_v45).Cardinality()).Minus(_1083_v40))
		}
		var _1089_v46 _dafny.Int
		_ = _1089_v46
		_1089_v46 = _dafny.IntOfInt64(-61)
		var _1090_v47 _dafny.Sequence
		_ = _1090_v47
		_1090_v47 = _dafny.UnicodeSeqOfUtf8Bytes("hu")
		var _1091_v48 _dafny.Map
		_ = _1091_v48
		_1091_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1089_v46, _1090_v47)
		var _hi5 _dafny.Int = Companion_Default___.SafeModuloInt((_1091_v48).Cardinality(), _1089_v46)
		_ = _hi5
		for _1092_i4 := (_dafny.Zero).Minus(_1089_v46); _1092_i4.Cmp(_hi5) < 0; _1092_i4 = _1092_i4.Plus(_dafny.One) {
			var _1093_v49 _dafny.MultiSet
			_ = _1093_v49
			_1093_v49 = _dafny.MultiSetOf((_this).F14())
			var _1094_v50 _dafny.Map
			_ = _1094_v50
			_1094_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1093_v49, (_this).F14())
			_1094_v50 = (_1094_v50).Update(_1093_v49, (_this).F14())
			var _1095_v51 _dafny.Array
			_ = _1095_v51
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_24
			var _nw145 _dafny.Array
			_ = _nw145
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw145 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Map = func(_1096_i5 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _dafny.IntOfInt64(323))
				}
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw145 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw145).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw145).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1095_v51 = _nw145
			var _1097_v53 _dafny.Map
			_ = _1097_v53
			_1097_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1095_v51, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter50 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F15()).Cardinality(), (_this).F17())).Keys().Elements()); ; {
					_compr_42, _ok50 := _iter50()
					if !_ok50 {
						break
					}
					var _1098_v52 _dafny.Int
					_1098_v52 = interface{}(_compr_42).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F15()).Cardinality(), (_this).F17())).Contains(_1098_v52) {
						_coll42.Add((_1098_v52).Times(_dafny.IntOfInt64(-276)), false)
					}
				}
				return _coll42.ToMap()
			}()).Update(_1089_v46, (_this).F13()), _1090_v47))
			var _1099_v54 _dafny.MultiSet
			_ = _1099_v54
			_1099_v54 = _dafny.MultiSetOf(_dafny.IntOfInt64(-748))
			var _1100_v55 _dafny.Map
			_ = _1100_v55
			_1100_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1099_v54).Cardinality()), _1063_v30)
			var _1101_v57 _dafny.Sequence
			_ = _1101_v57
			_1101_v57 = _dafny.SeqOf(_1063_v30)
			var _1102_v58 _dafny.Map
			_ = _1102_v58
			_1102_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v55, _1101_v57)
			_1097_v53 = (_1097_v53).Update(_1095_v51, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1100_v55, _1090_v47)).Merge((func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter51 := _dafny.Iterate((_1102_v58).Keys().Elements()); ; {
					_compr_43, _ok51 := _iter51()
					if !_ok51 {
						break
					}
					var _1103_v56 _dafny.Map
					_1103_v56 = interface{}(_compr_43).(_dafny.Map)
					if (_1102_v58).Contains(_1103_v56) {
						_coll43.Add(_1103_v56, _1090_v47)
					}
				}
				return _coll43.ToMap()
			}()).Update(_1100_v55, _1090_v47)))
			var _1104_v59 _dafny.Set
			_ = _1104_v59
			_1104_v59 = _dafny.SetOf(_1063_v30, (_this).F13())
			_1063_v30 = ((_dafny.Zero).Minus(_1089_v46)).Cmp((_1104_v59).Cardinality()) <= 0
			_1063_v30 = (_this).F17()
		}
		var _1105_v60 D6
		_ = _1105_v60
		_1105_v60 = Companion_D6_.Create_DC12_(false, _1090_v47, _dafny.IntOfInt64(629), _1063_v30)
		var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index157
		((_this).F14()).ArraySet1((_1105_v60).Dtor_cf21(), (_index157).Int())
		var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index158
		((_this).F14()).ArraySet1((_dafny.SetOf(_1089_v46)).IsDisjointFrom((Companion_D9_.Create_DC20_((_this).F15())).Dtor_cf36()), (_index158).Int())
	}
}
func (_this *C3) M5(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _source19 D1 = _this.F16()
		_ = _source19
		if _source19.Is_DC2() {
			var _1106___mcc_h0 _dafny.Int = _source19.Get_().(D1_DC2).Cf2
			_ = _1106___mcc_h0
			var _1107___mcc_h1 _dafny.Int = _source19.Get_().(D1_DC2).Cf3
			_ = _1107___mcc_h1
			var _1108___mcc_h2 bool = _source19.Get_().(D1_DC2).Cf4
			_ = _1108___mcc_h2
			var _1109_cf4 bool = _1108___mcc_h2
			_ = _1109_cf4
			var _1110_cf3 _dafny.Int = _1107___mcc_h1
			_ = _1110_cf3
			var _1111_cf2 _dafny.Int = _1106___mcc_h0
			_ = _1111_cf2
			var _1112_v0 _dafny.Sequence
			_ = _1112_v0
			_1112_v0 = _dafny.UnicodeSeqOfUtf8Bytes("havb")
			var _1113_v1 _dafny.MultiSet
			_ = _1113_v1
			_1113_v1 = _dafny.MultiSetOf(_1112_v0, _1112_v0)
			_1109_cf4 = !((_1113_v1).IsProperSubsetOf(_1113_v1))
			var _1114_v2 _dafny.Array
			_ = _1114_v2
			var _nw146 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw146
			_1114_v2 = _nw146
			_1114_v2 = _1114_v2
			var _1115_v3 _dafny.Set
			_ = _1115_v3
			_1115_v3 = _dafny.SetOf(_1114_v2)
			_1111_cf2 = (_1115_v3).Cardinality()
			var _1116_v4 _dafny.MultiSet
			_ = _1116_v4
			_1116_v4 = _dafny.MultiSetOf((_1111_cf2).Cmp((_this).Fm8((_this).F17(), _this.F16(), globalState)) == 0, (_this).F17())
			var _1117_v5 _dafny.Sequence
			_ = _1117_v5
			_1117_v5 = _dafny.SeqOf(_1109_cf4, _1109_cf4)
			var _1118_v6 _dafny.Sequence
			_ = _1118_v6
			_1118_v6 = _dafny.SeqOf(_1117_v5, _1117_v5)
			var _rhs134 _dafny.MultiSet = (func() _dafny.MultiSet {
				if true {
					return _1116_v4
				}
				return (_dafny.MultiSetFromSeq((_1118_v6).Select((Companion_Default___.SafeIndex(_1110_cf3, _dafny.IntOfUint32((_1118_v6).Cardinality()))).Uint32()).(_dafny.Sequence))).Difference(_1116_v4)
			})()
			_ = _rhs134
			var _rhs135 bool = (_this).F13()
			_ = _rhs135
			var _rhs136 bool = true
			_ = _rhs136
			_1116_v4 = _rhs134
			_1109_cf4 = _rhs135
			_1109_cf4 = _rhs136
		} else {
			var _1119___mcc_h3 _dafny.Sequence = _source19.Get_().(D1_DC1).Cf1
			_ = _1119___mcc_h3
			var _1120_cf1 _dafny.Sequence = _1119___mcc_h3
			_ = _1120_cf1
			var _1121_v7 _dafny.Array
			_ = _1121_v7
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_25
			var _nw147 _dafny.Array
			_ = _nw147
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw147 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = (func(_1122_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1123_i0 _dafny.Int) _dafny.Int {
						return (_1123_i0).Plus(_1122_p0)
					}
				})(p0)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw147 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw147).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw147).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_1121_v7 = _nw147
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1121_v7), 0))
			_ = _index159
			(_1121_v7).ArraySet1(p0, (_index159).Int())
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1121_v7), 0))
			_ = _index160
			(_1121_v7).ArraySet1(p0, (_index160).Int())
			var _1124_v8 _dafny.Set
			_ = _1124_v8
			_1124_v8 = _dafny.SetOf((_this).F17(), _dafny.Companion_Sequence_.Equal(_1120_cf1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(681))).Uint32(), func(coer74 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg74 _dafny.Int) interface{} {
					return coer74(arg74)
				}
			}(func(_1125_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			}))), (_this).F13(), (_this).F18())
			_1124_v8 = _1124_v8
			_1120_cf1 = _1120_cf1
			var _1126_v9 _dafny.CodePoint
			_ = _1126_v9
			_1126_v9 = _dafny.CodePoint('h')
			_1126_v9 = _1126_v9
		}
		var _1127_v10 _dafny.Array
		_ = _1127_v10
		var _nw148 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(23))
		_ = _nw148
		_1127_v10 = _nw148
		_1127_v10 = (func() _dafny.Array {
			if (p0).Cmp(p0) < 0 {
				return _1127_v10
			}
			return _1127_v10
		})()
		var _1128_v11 _dafny.Array
		_ = _1128_v11
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_26
		var _nw149 _dafny.Array
		_ = _nw149
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw149 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) _dafny.Int = (func(_1129_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1130_i3 _dafny.Int) _dafny.Int {
					return (_1130_i3).Times(_1129_p0)
				}
			})(p0)
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw149 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw149).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw149).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_1128_v11 = _nw149
		var _1131_v12 D2
		_ = _1131_v12
		_1131_v12 = Companion_D2_.Create_DC4_(_1128_v11, (_this).F17(), (_dafny.Zero).Minus(p0), true, (_this).F14())
		var _1132_v13 _dafny.MultiSet
		_ = _1132_v13
		_1132_v13 = _dafny.MultiSetOf((_this).F17())
		var _1133_v14 _dafny.MultiSet
		_ = _1133_v14
		_1133_v14 = _1132_v13
		var _1134_v15 _dafny.Map
		_ = _1134_v15
		_1134_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_this).F18()))
		var _1135_v16 D3
		_ = _1135_v16
		_1135_v16 = Companion_D3_.Create_DC6_((_dafny.Zero).Minus(p0), (_this.F12()).Dtor_cf4())
		var _1136_v17 _dafny.Array
		_ = _1136_v17
		var _nwElement0_34 D2 = _1131_v12
		_ = _nwElement0_34
		var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(14))
		_ = _nw150
		(_nw150).ArraySet1(_nwElement0_34, 0)
		(_nw150).ArraySet1(_1131_v12, 1)
		(_nw150).ArraySet1(Companion_D2_.Create_DC4_(_1128_v11, (_this).F17(), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1133_v14)).Cardinality()), (func() bool {
			if (_1134_v15).Contains(p0) {
				return (_1134_v15).Get(p0).(bool)
			}
			return false
		})(), (_this).F14()), 2)
		(_nw150).ArraySet1(Companion_D2_.Create_DC4_(_1128_v11, (_this).F18(), p0, (_this).F17(), (_this).F14()), 3)
		(_nw150).ArraySet1(Companion_D2_.Create_DC4_(_1128_v11, (_this).F18(), p0, (_1135_v16).Dtor_cf13(), (_this).F14()), 4)
		(_nw150).ArraySet1(_1131_v12, 5)
		(_nw150).ArraySet1(_1131_v12, 6)
		(_nw150).ArraySet1(_1131_v12, 7)
		(_nw150).ArraySet1(_1131_v12, 8)
		(_nw150).ArraySet1(_1131_v12, 9)
		(_nw150).ArraySet1(_1131_v12, 10)
		(_nw150).ArraySet1(_1131_v12, 11)
		(_nw150).ArraySet1(Companion_D2_.Create_DC4_(_1128_v11, (_this).F13(), p0, (_this).F13(), (_1131_v12).Dtor_cf10()), 12)
		(_nw150).ArraySet1(_1131_v12, 13)
		_1136_v17 = _nw150
		var _1137_v18 _dafny.Set
		_ = _1137_v18
		_1137_v18 = _dafny.SetOf(_1136_v17)
		var _1138_i2 _dafny.Int
		_ = _1138_i2
		_1138_i2 = _dafny.Zero
		{
			for (_dafny.SetOf(_1136_v17)).IsSubsetOf(_1137_v18) {
				{
					if (_1138_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1138_i2 = (_1138_i2).Plus(_dafny.One)
					_1134_v15 = (_1134_v15).Update(p0, (_this).F17())
					var _1139_v19 _dafny.Map
					_ = _1139_v19
					_1139_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _1140_v20 _dafny.Sequence
					_ = _1140_v20
					_1140_v20 = _dafny.SeqOf((_this).F17())
					_1139_v19 = (_1139_v19).Update(Companion_Default___.Fm0(_1140_v20, (_this).F18(), (_this).F13(), p0, globalState), _dafny.IntOfInt64(114))
					r0 = (_this).F18()
					var _1141_v21 _dafny.MultiSet
					_ = _1141_v21
					_1141_v21 = _dafny.MultiSetOf(p0, p0)
					r0 = (func() bool {
						if !(_1141_v21).Equals(_1141_v21) {
							return (_this).F18()
						}
						return ((_this).F15()).IsSubsetOf((_this).F15())
					})()
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1142_v22 _dafny.Sequence
		_ = _1142_v22
		_1142_v22 = _dafny.SeqOf(!(false))
		var _1143_v23 _dafny.Map
		_ = _1143_v23
		_1143_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-575), p0)
		var _1144_v24 _dafny.Map
		_ = _1144_v24
		_1144_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D9_.Create_DC23_(p0, p0, (_this).Fm15(globalState), Companion_Default___.Fm0(_1142_v22, (_this).F17(), false, (_1143_v23).Cardinality(), globalState), (_dafny.Zero).Minus(p0)), p0)
		var _1145_v25 _dafny.Sequence
		_ = _1145_v25
		_1145_v25 = _dafny.SeqOf((_1144_v24).Cardinality(), p0)
		var _1146_v26 _dafny.Map
		_ = _1146_v26
		_1146_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1145_v25, p0)
		var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_1128_v11), 0))
		_ = _index161
		(_1128_v11).ArraySet1((_1146_v26).Cardinality(), (_index161).Int())
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_1128_v11), 0))
		_ = _index162
		(_1128_v11).ArraySet1(p0, (_index162).Int())
		var _1147_v27 _dafny.Sequence
		_ = _1147_v27
		_1147_v27 = _dafny.UnicodeSeqOfUtf8Bytes("itq")
		var _1148_v28 D6
		_ = _1148_v28
		_1148_v28 = Companion_D6_.Create_DC12_((_this).F13(), _1147_v27, p0, (_this).F18())
		var _1149_v29 *C2
		_ = _1149_v29
		var _nw151 *C2 = New_C2_()
		_ = _nw151
		_nw151.Ctor__(Companion_Default___.SafeModuloInt(p0, (_1128_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_1128_v11), 0))).Int()).(_dafny.Int)), (_this).F13(), _this.F16(), (func() bool {
			if (_1148_v28).Dtor_cf21() {
				return (_this).F13()
			}
			return (_this).F13()
		})(), (_this).F14(), ((_this).F15()).Union((_this).F15()), _this.F12(), true)
		_1149_v29 = _nw151
		_1147_v27 = _dafny.Companion_Sequence_.Concatenate(_1147_v27, _1147_v27)
		r0 = (_this).F17()
		return r0
	}
}
func (_this *C3) M11(p0 _dafny.Array, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _1150_v0 _dafny.Int
		_ = _1150_v0
		_1150_v0 = _dafny.IntOfInt64(685)
		if Companion_Default___.Fm30(_1150_v0, globalState) {
			var _1151_v1 _dafny.Map
			_ = _1151_v1
			_1151_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(_1150_v0, globalState), p1)
			var _1152_v2 _dafny.Sequence
			_ = _1152_v2
			_1152_v2 = _dafny.SeqOf(true, false, p2, (_this).F18(), true)
			_1151_v1 = (_1151_v1).Update((_1152_v2).Select((Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1152_v2).Cardinality()))).Uint32()).(bool), (_this).F18())
			var _1153_v3 _dafny.Sequence
			_ = _1153_v3
			_1153_v3 = _dafny.SeqOf(_dafny.IntOfInt64(-44))
			var _1154_v4 _dafny.Map
			_ = _1154_v4
			_1154_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1150_v0)
			var _1155_v5 _dafny.Set
			_ = _1155_v5
			_1155_v5 = _dafny.SetOf(p2)
			var _1156_v6 _dafny.Map
			_ = _1156_v6
			_1156_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1155_v5, _1150_v0)
			var _1157_v7 _dafny.MultiSet
			_ = _1157_v7
			_1157_v7 = _dafny.MultiSetOf(false)
			var _1158_v8 _dafny.Sequence
			_ = _1158_v8
			_1158_v8 = _dafny.UnicodeSeqOfUtf8Bytes("jlbyvmu")
			var _1159_v9 _dafny.Array
			_ = _1159_v9
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(9)
			_ = _len0_27
			var _nw152 _dafny.Array
			_ = _nw152
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw152 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Int = (func(_1160_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1161_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1161_i0, _1160_v0)
					}
				})(_1150_v0)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw152 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw152).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw152).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_1159_v9 = _nw152
			var _1162_v10 _dafny.Map
			_ = _1162_v10
			_1162_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v0, _1150_v0)
			var _1163_v11 _dafny.Sequence
			_ = _1163_v11
			_1163_v11 = _dafny.SeqOf(_1154_v4)
			var _1164_v12 D9
			_ = _1164_v12
			_1164_v12 = Companion_D9_.Create_DC23_(_1150_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality()), _1150_v0, _dafny.IntOfUint32((_1163_v11).Cardinality()), _1150_v0)
			var _1165_v13 _dafny.Array
			_ = _1165_v13
			var _nwElement0_35 _dafny.Int = (func() _dafny.Int {
				if p2 {
					return (_1153_v3).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_1154_v4).Contains((_this).F17()) {
							return (_1154_v4).Get((_this).F17()).(_dafny.Int)
						}
						return _1150_v0
					})(), _dafny.IntOfUint32((_1153_v3).Cardinality()))).Uint32()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(272)
			})()
			_ = _nwElement0_35
			var _nw153 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(24))
			_ = _nw153
			(_nw153).ArraySet1(_nwElement0_35, 0)
			(_nw153).ArraySet1(((Companion_Default___.Fm44((_this).F17(), globalState)).Merge(_1156_v6)).Cardinality(), 1)
			(_nw153).ArraySet1(Companion_Default___.Fm0(_1152_v2, p1, false, _1150_v0, globalState), 2)
			(_nw153).ArraySet1(_1150_v0, 3)
			(_nw153).ArraySet1(_1150_v0, 4)
			(_nw153).ArraySet1(_1150_v0, 5)
			(_nw153).ArraySet1(_1150_v0, 6)
			(_nw153).ArraySet1((_1150_v0).Times(_1150_v0), 7)
			(_nw153).ArraySet1(Companion_Default___.SafeDivisionInt(_1150_v0, _1150_v0), 8)
			(_nw153).ArraySet1(_1150_v0, 9)
			(_nw153).ArraySet1((_dafny.Zero).Minus((_dafny.IntOfInt64(305)).Plus(((_this).F15()).Cardinality())), 10)
			(_nw153).ArraySet1(_dafny.IntOfInt64(770), 11)
			(_nw153).ArraySet1(Companion_Default___.SafeDivisionInt((_1157_v7).Cardinality(), (_dafny.Zero).Minus(_1150_v0)), 12)
			(_nw153).ArraySet1(Companion_Default___.Fm0(_1152_v2, (_this).F13(), false, Companion_Default___.Fm0(_1152_v2, (_this).F13(), Companion_Default___.Fm30(_dafny.IntOfUint32((_1158_v8).Cardinality()), globalState), _1150_v0, globalState), globalState), 13)
			(_nw153).ArraySet1(_1150_v0, 14)
			(_nw153).ArraySet1(_1150_v0, 15)
			(_nw153).ArraySet1((_dafny.SetOf((Companion_D2_.Create_DC4_(_1159_v9, !((_this).F17()), _1150_v0, (_this).F18(), (_this).F14())).Dtor_cf8())).Cardinality(), 16)
			(_nw153).ArraySet1(_1150_v0, 17)
			(_nw153).ArraySet1((_1150_v0).Minus(_dafny.IntOfInt64(998)), 18)
			(_nw153).ArraySet1((_dafny.MultiSetOf(_1150_v0, _1150_v0, (_1162_v10).Cardinality(), _dafny.IntOfInt64(-793), ((_this).F15()).Cardinality())).Cardinality(), 19)
			(_nw153).ArraySet1((_1150_v0).Times(_1150_v0), 20)
			(_nw153).ArraySet1((func() _dafny.Int {
				if (_1162_v10).Contains((_1164_v12).Dtor_cf45()) {
					return (_1162_v10).Get((_1164_v12).Dtor_cf45()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_1158_v8).Cardinality())
			})(), 21)
			(_nw153).ArraySet1(_1150_v0, 22)
			(_nw153).ArraySet1(_dafny.IntOfInt64(630), 23)
			_1165_v13 = _nw153
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))
			_ = _index163
			(_1165_v13).ArraySet1(_1150_v0, (_index163).Int())
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))
			_ = _index164
			(_1165_v13).ArraySet1(((_this).Fm15(globalState)).Times(_1150_v0), (_index164).Int())
			var _1166_v14 *C1
			_ = _1166_v14
			var _nw154 *C1 = New_C1_()
			_ = _nw154
			_nw154.Ctor__(_dafny.CodePoint('j'), _this.F12(), (_this).F17())
			_1166_v14 = _nw154
			var _1167_v15 _dafny.Array
			_ = _1167_v15
			var _nwElement0_36 *C1 = _1166_v14
			_ = _nwElement0_36
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(18))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_36, 0)
			(_nw155).ArraySet1(_1166_v14, 1)
			(_nw155).ArraySet1(_1166_v14, 2)
			(_nw155).ArraySet1(_1166_v14, 3)
			(_nw155).ArraySet1(_1166_v14, 4)
			(_nw155).ArraySet1(_1166_v14, 5)
			(_nw155).ArraySet1(_1166_v14, 6)
			(_nw155).ArraySet1(_1166_v14, 7)
			(_nw155).ArraySet1(_1166_v14, 8)
			(_nw155).ArraySet1(_1166_v14, 9)
			(_nw155).ArraySet1(_1166_v14, 10)
			(_nw155).ArraySet1(_1166_v14, 11)
			(_nw155).ArraySet1(_1166_v14, 12)
			(_nw155).ArraySet1(_1166_v14, 13)
			(_nw155).ArraySet1(_1166_v14, 14)
			(_nw155).ArraySet1(_1166_v14, 15)
			(_nw155).ArraySet1(_1166_v14, 16)
			(_nw155).ArraySet1(_1166_v14, 17)
			_1167_v15 = _nw155
			var _1168_v16 _dafny.MultiSet
			_ = _1168_v16
			_1168_v16 = _dafny.MultiSetOf(_1167_v15, _1167_v15)
			var _1169_v17 bool
			_ = _1169_v17
			_1169_v17 = false
			var _rhs137 _dafny.Int = (_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)
			_ = _rhs137
			var _rhs138 _dafny.MultiSet = _1168_v16
			_ = _rhs138
			var _rhs139 bool = ((_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)).Cmp((_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)) > 0
			_ = _rhs139
			var _rhs140 bool = ((_1157_v7).Cardinality()).Cmp((_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)) == 0
			_ = _rhs140
			var _rhs141 bool = p1
			_ = _rhs141
			var _lhs114 *GlobalState = globalState
			_ = _lhs114
			_lhs114.F7 = _rhs137
			_1168_v16 = _rhs138
			_1169_v17 = _rhs139
			_1169_v17 = _rhs140
			_1169_v17 = _rhs141
			if (_this).F13() {
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))
				_ = _index165
				(_1165_v13).ArraySet1(_dafny.IntOfInt64(-903), (_index165).Int())
				(globalState).F7 = (_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)
				_1169_v17 = p2
				var _1170_v18 _dafny.Map
				_ = _1170_v18
				_1170_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F15()).Cardinality(), p2)
				var _rhs142 _dafny.CodePoint = (func() _dafny.CodePoint {
					if (func() bool {
						if (_1170_v18).Contains((_1157_v7).Cardinality()) {
							return (_1170_v18).Get((_1157_v7).Cardinality()).(bool)
						}
						return (_this).F13()
					})() {
						return (func() _dafny.CodePoint {
							if p2 {
								return _1166_v14.F23
							}
							return _1166_v14.F23
						})()
					}
					return _1166_v14.F23
				})()
				_ = _rhs142
				var _rhs143 _dafny.CodePoint = _1166_v14.F23
				_ = _rhs143
				var _lhs115 *C1 = _1166_v14
				_ = _lhs115
				var _lhs116 *C1 = _1166_v14
				_ = _lhs116
				_lhs115.F23 = _rhs142
				_lhs116.F23 = _rhs143
				var _1171_v19 _dafny.Array
				_ = _1171_v19
				var _nw156 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
				_ = _nw156
				_1171_v19 = _nw156
				_1171_v19 = p0
			} else {
				var _1172_v20 _dafny.Map
				_ = _1172_v20
				_1172_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int), p0)
				_1172_v20 = (_1172_v20).Update(_dafny.IntOfUint32((_1152_v2).Cardinality()), (_this).F14())
				var _1173_v21 _dafny.Set
				_ = _1173_v21
				_1173_v21 = _dafny.SetOf(_1158_v8)
				var _1174_v22 _dafny.Set
				_ = _1174_v22
				_1174_v22 = _1173_v21
				var _pat_let_tv31 = _1165_v13
				_ = _pat_let_tv31
				var _pat_let_tv32 = _1165_v13
				_ = _pat_let_tv32
				_1174_v22 = (func() _dafny.Set {
					if (func(_pat_let31_0 D1) D1 {
						return func(_1175_dt__update__tmp_h0 D1) D1 {
							return func(_pat_let32_0 _dafny.Int) D1 {
								return func(_1176_dt__update_hcf2_h0 _dafny.Int) D1 {
									return Companion_D1_.Create_DC2_(_1176_dt__update_hcf2_h0, (_1175_dt__update__tmp_h0).Dtor_cf3(), (_1175_dt__update__tmp_h0).Dtor_cf4())
								}(_pat_let32_0)
							}((_pat_let_tv32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_pat_let_tv31), 0))).Int()).(_dafny.Int))
						}(_pat_let31_0)
					}(_this.F12())).Dtor_cf4() {
						return _1173_v21
					}
					return _1174_v22
				})()
				var _1177_v23 _dafny.Map
				_ = _1177_v23
				_1177_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1166_v14, _1150_v0)
				_1177_v23 = (_1177_v23).Update(_1166_v14, (_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int))
				var _1178_v24 _dafny.Array
				_ = _1178_v24
				var _nwElement0_37 _dafny.Array = _1159_v9
				_ = _nwElement0_37
				var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(20))
				_ = _nw157
				(_nw157).ArraySet1(_nwElement0_37, 0)
				(_nw157).ArraySet1(_1159_v9, 1)
				(_nw157).ArraySet1(_1159_v9, 2)
				(_nw157).ArraySet1(_1159_v9, 3)
				(_nw157).ArraySet1(_1159_v9, 4)
				(_nw157).ArraySet1(_1159_v9, 5)
				(_nw157).ArraySet1(_1159_v9, 6)
				(_nw157).ArraySet1(_1165_v13, 7)
				(_nw157).ArraySet1(_1165_v13, 8)
				(_nw157).ArraySet1(_1165_v13, 9)
				(_nw157).ArraySet1((Companion_D2_.Create_DC4_(_1159_v9, (_this).F18(), (_this).Fm8((_this).F13(), _this.F16(), globalState), !((_this).F18()), (_this).F14())).Dtor_cf6(), 10)
				(_nw157).ArraySet1(_1165_v13, 11)
				(_nw157).ArraySet1(_1159_v9, 12)
				(_nw157).ArraySet1(_1165_v13, 13)
				(_nw157).ArraySet1(_1159_v9, 14)
				(_nw157).ArraySet1(_1159_v9, 15)
				(_nw157).ArraySet1(_1165_v13, 16)
				(_nw157).ArraySet1(_1165_v13, 17)
				(_nw157).ArraySet1(_1159_v9, 18)
				(_nw157).ArraySet1(_1165_v13, 19)
				_1178_v24 = _nw157
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1178_v24), 0))
				_ = _index166
				(_1178_v24).ArraySet1(_1159_v9, (_index166).Int())
				var _1179_v25 _dafny.Map
				_ = _1179_v25
				_1179_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v0, _1165_v13)
				var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1178_v24), 0))
				_ = _index167
				(_1178_v24).ArraySet1((func() _dafny.Array {
					if (_1179_v25).Contains(_1150_v0) {
						return (_1179_v25).Get(_1150_v0).(_dafny.Array)
					}
					return (Companion_D2_.Create_DC4_(_1165_v13, (_this).F13(), _dafny.IntOfInt64(612), (_this).F13(), (_this).F14())).Dtor_cf6()
				})(), (_index167).Int())
				var _1180_v26 _dafny.Map
				_ = _1180_v26
				_1180_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.SetOf((_1165_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(465), _dafny.ArrayLen((_1165_v13), 0))).Int()).(_dafny.Int)))
				var _1181_v27 _dafny.Map
				_ = _1181_v27
				_1181_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1166_v14).Fm43(globalState), (func() _dafny.Set {
					if (_1180_v26).Contains((_this).F17()) {
						return (_1180_v26).Get((_this).F17()).(_dafny.Set)
					}
					return (_this).F15()
				})())
				_1180_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F15())
			}
			(_this).M12(globalState)
		} else {
			var _1182_v28 _dafny.Sequence
			_ = _1182_v28
			_1182_v28 = _dafny.UnicodeSeqOfUtf8Bytes("lxk")
			var _1183_v29 _dafny.Sequence
			_ = _1183_v29
			_1183_v29 = _dafny.SeqOf(_1150_v0)
			var _1184_v30 D3
			_ = _1184_v30
			_1184_v30 = Companion_D3_.Create_DC5_(_dafny.CodePoint('c'))
			var _1185_v31 _dafny.MultiSet
			_ = _1185_v31
			_1185_v31 = _dafny.MultiSetOf((_1184_v30).Dtor_cf11())
			var _1186_v32 _dafny.CodePoint
			_ = _1186_v32
			_1186_v32 = _dafny.CodePoint('q')
			var _1187_v33 _dafny.Map
			_ = _1187_v33
			_1187_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(997), Companion_D7_.Create_DC17_(Companion_D7_.Create_DC15_(_1183_v29)))
			var _1188_v34 _dafny.MultiSet
			_ = _1188_v34
			_1188_v34 = _dafny.MultiSetOf(_1150_v0, _1150_v0)
			var _1189_v35 D7
			_ = _1189_v35
			_1189_v35 = Companion_D7_.Create_DC16_(_1150_v0)
			var _1190_v36 D7
			_ = _1190_v36
			_1190_v36 = Companion_D7_.Create_DC17_(_1189_v35)
			var _1191_v37 _dafny.Array
			_ = _1191_v37
			var _nwElement0_38 _dafny.Int = _1150_v0
			_ = _nwElement0_38
			var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(20))
			_ = _nw158
			(_nw158).ArraySet1(_nwElement0_38, 0)
			(_nw158).ArraySet1(_1150_v0, 1)
			(_nw158).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(642), _1150_v0), 2)
			(_nw158).ArraySet1(Companion_Default___.SafeModuloInt(_1150_v0, _1150_v0), 3)
			(_nw158).ArraySet1((_dafny.Zero).Minus(_1150_v0), 4)
			(_nw158).ArraySet1(_1150_v0, 5)
			(_nw158).ArraySet1(_dafny.IntOfUint32((_1182_v28).Cardinality()), 6)
			(_nw158).ArraySet1((_dafny.IntOfUint32((_1183_v29).Cardinality())).Minus((_dafny.Zero).Minus(_1150_v0)), 7)
			(_nw158).ArraySet1(_1150_v0, 8)
			(_nw158).ArraySet1(_1150_v0, 9)
			(_nw158).ArraySet1((func() _dafny.Int {
				if p1 {
					return (func() _dafny.Int {
						if (_1185_v31).Contains(_1186_v32) {
							return (_1185_v31).Multiplicity(_1186_v32)
						}
						return _1150_v0
					})()
				}
				return (_this).Fm9(_1186_v32, globalState)
			})(), 10)
			(_nw158).ArraySet1(_1150_v0, 11)
			(_nw158).ArraySet1(((_this).F15()).Cardinality(), 12)
			(_nw158).ArraySet1(_1150_v0, 13)
			(_nw158).ArraySet1((_1150_v0).Plus(_1150_v0), 14)
			(_nw158).ArraySet1(_1150_v0, 15)
			(_nw158).ArraySet1(_dafny.IntOfInt64(-151), 16)
			(_nw158).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1182_v28, (Companion_Default___.SafeIndex(((_1187_v33).Update((_1188_v34).Cardinality(), _1190_v36)).Cardinality(), _dafny.IntOfUint32((_1182_v28).Cardinality()))).Uint32(), _1186_v32)).Cardinality()), 17)
			(_nw158).ArraySet1(_1150_v0, 18)
			(_nw158).ArraySet1((_1150_v0).Plus(_dafny.IntOfUint32((_1182_v28).Cardinality())), 19)
			_1191_v37 = _nw158
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))
			_ = _index168
			(_1191_v37).ArraySet1(_1150_v0, (_index168).Int())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))
			_ = _index169
			(_1191_v37).ArraySet1(_1150_v0, (_index169).Int())
			_1150_v0 = _dafny.IntOfUint32((Companion_Default___.Fm20(globalState)).Cardinality())
			var _1192_v38 _dafny.Array
			_ = _1192_v38
			var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(19))
			_ = _nw159
			_1192_v38 = _nw159
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1192_v38), 0))
			_ = _index170
			(_1192_v38).ArraySet1CodePoint(_1186_v32, (_index170).Int())
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(494), _dafny.ArrayLen((_1192_v38), 0))
			_ = _index171
			(_1192_v38).ArraySet1CodePoint(_1186_v32, (_index171).Int())
			var _1193_v39 _dafny.Array
			_ = _1193_v39
			var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
			_ = _nw160
			_1193_v39 = _nw160
			var _1194_v40 _dafny.Array
			_ = _1194_v40
			_1194_v40 = _1193_v39
			var _source20 _dafny.Array = _1194_v40
			_ = _source20
			var _1195___mcc_h0 _dafny.Array = _source20
			_ = _1195___mcc_h0
			var _1196_cf15 _dafny.Array = _1195___mcc_h0
			_ = _1196_cf15
			var _1197_v41 _dafny.Set
			_ = _1197_v41
			_1197_v41 = _dafny.SetOf((_this).F18())
			_1150_v0 = ((_1188_v34).Union((_dafny.MultiSetOf((_1197_v41).Cardinality())).Intersection(_dafny.MultiSetOf((_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int))))).Cardinality()
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((p0), 0))
			_ = _index172
			(p0).ArraySet1(!(p1), (_index172).Int())
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((p0), 0))
			_ = _index173
			(p0).ArraySet1(p1, (_index173).Int())
			var _1198_v42 _dafny.Array
			_ = _1198_v42
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_28
			var _nw161 _dafny.Array
			_ = _nw161
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw161 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) bool = func(_1199_i1 _dafny.Int) bool {
					return false
				}
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw161 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw161).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw161).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_1198_v42 = _nw161
			var _1200_v43 _dafny.Array
			_ = _1200_v43
			var _nwElement0_39 _dafny.Array = _1191_v37
			_ = _nwElement0_39
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.One)
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_39, 0)
			_1200_v43 = _nw162
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1200_v43), 0))
			_ = _index174
			(_1200_v43).ArraySet1(_1191_v37, (_index174).Int())
			var _1201_v44 _dafny.MultiSet
			_ = _1201_v44
			_1201_v44 = _dafny.MultiSetOf((_this).F14())
			var _1202_v45 _dafny.MultiSet
			_ = _1202_v45
			_1202_v45 = _dafny.MultiSetOf(_1198_v42)
			var _1203_v46 _dafny.MultiSet
			_ = _1203_v46
			_1203_v46 = _dafny.MultiSetOf(_1183_v29)
			var _1204_v47 _dafny.Array
			_ = _1204_v47
			var _nw163 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(5))
			_ = _nw163
			_1204_v47 = _nw163
			var _1205_v48 _dafny.Map
			_ = _1205_v48
			_1205_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1204_v47, p1)
			var _1206_v49 _dafny.Map
			_ = _1206_v49
			_1206_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1205_v48, _1198_v42)
			var _1207_v50 _dafny.Sequence
			_ = _1207_v50
			_1207_v50 = _dafny.SeqOf(_1191_v37, (func() _dafny.Array {
				if p1 {
					return _1191_v37
				}
				return _1191_v37
			})())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((p0), 0))
			_ = _index175
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1200_v43), 0))
			_ = _index176
			var _rhs144 _dafny.Int = Companion_Default___.SafeModuloInt((_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus((_1202_v45).Cardinality()))
			_ = _rhs144
			var _rhs145 bool = (_1203_v46).IsDisjointFrom(_1203_v46)
			_ = _rhs145
			var _rhs146 _dafny.Array = (func() _dafny.Array {
				if (_1206_v49).Contains(_1205_v48) {
					return (_1206_v49).Get(_1205_v48).(_dafny.Array)
				}
				return (_this).F14()
			})()
			_ = _rhs146
			var _rhs147 _dafny.Int = (_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int)
			_ = _rhs147
			var _rhs148 _dafny.Array = (_1207_v50).Select((Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1207_v50).Cardinality()))).Uint32()).(_dafny.Array)
			_ = _rhs148
			var _lhs117 _dafny.Array = p0
			_ = _lhs117
			var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((p0), 0))
			_ = _lhs118
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			var _lhs120 _dafny.Array = _1200_v43
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_1200_v43), 0))
			_ = _lhs121
			_1150_v0 = _rhs144
			(_lhs117).ArraySet1(_rhs145, (_lhs118).Int())
			_1198_v42 = _rhs146
			_lhs119.F7 = _rhs147
			(_lhs120).ArraySet1(_rhs148, (_lhs121).Int())
			var _1208_v51 _dafny.Sequence
			_ = _1208_v51
			_1208_v51 = _dafny.SeqOf(p1, (_this).F13())
			var _1209_v52 _dafny.Map
			_ = _1209_v52
			_1209_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_this).F18() {
					return _1198_v42
				}
				return _1198_v42
			})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_1208_v51).Cardinality())))
			_1209_v52 = (_1209_v52).Update(_1198_v42, Companion_Default___.SafeModuloInt(_1150_v0, (_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int)))
			if p1 {
				var _1210_v53 bool
				_ = _1210_v53
				_1210_v53 = false
				_1210_v53 = p2
				_1182_v28 = Companion_Default___.Fm20(globalState)
				_1186_v32 = _1186_v32
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))
				_ = _index177
				(_1191_v37).ArraySet1((_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(41), _1150_v0)).Times((_dafny.Zero).Minus(_1150_v0))), (_index177).Int())
				_1210_v53 = ((func() _dafny.MultiSet {
					if (_this).F18() {
						return _1188_v34
					}
					return _dafny.MultiSetOf(_dafny.IntOfInt64(670))
				})()).IsDisjointFrom(Companion_Default___.Fm19(globalState))
			} else {
				var _1211_v54 _dafny.Map
				_ = _1211_v54
				_1211_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC11_(_1188_v34), false)
				var _1212_v55 *C2
				_ = _1212_v55
				var _nw164 *C2 = New_C2_()
				_ = _nw164
				_nw164.Ctor__(_1150_v0, (_this).F18(), _this.F16(), (func() bool {
					if (_1211_v54).Contains(Companion_D6_.Create_DC11_(_1188_v34)) {
						return (_1211_v54).Get(Companion_D6_.Create_DC11_(_1188_v34)).(bool)
					}
					return p1
				})(), p0, ((_this).F15()).Difference((_this).F15()), _this.F12(), ((_dafny.Zero).Minus(_1150_v0)).Cmp((_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int)) == 0)
				_1212_v55 = _nw164
				var _1213_v56 _dafny.Map
				_ = _1213_v56
				_1213_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1191_v37, (_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int))
				_1213_v56 = (_1213_v56).Update(_1191_v37, Companion_Default___.SafeDivisionInt((_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int), _1150_v0))
				var _1214_v57 _dafny.Map
				_ = _1214_v57
				_1214_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int), (_1212_v55).F22())
				_1214_v57 = _1214_v57
				var _1215_v58 _dafny.Map
				_ = _1215_v58
				_1215_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1212_v55).F21())
				_1215_v58 = (_1215_v58).Update((_this).F13(), (_1212_v55).F21())
				(globalState).F7 = Companion_Default___.SafeDivisionInt(_1150_v0, (_1191_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(920), _dafny.ArrayLen((_1191_v37), 0))).Int()).(_dafny.Int))
			}
		}
		var _1216_v59 _dafny.Map
		_ = _1216_v59
		_1216_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v0, _dafny.CodePoint('h'))
		var _1217_v60 _dafny.Sequence
		_ = _1217_v60
		_1217_v60 = _dafny.SeqOf(_1150_v0, (_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F17(), (_this).F13(), (_this).F13(), (_this).F18()))).Cardinality())
		var _1218_v61 _dafny.Array
		_ = _1218_v61
		var _nwElement0_40 _dafny.Int = _1150_v0
		_ = _nwElement0_40
		var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(9))
		_ = _nw165
		(_nw165).ArraySet1(_nwElement0_40, 0)
		(_nw165).ArraySet1(((_1216_v59).Merge(_1216_v59)).Cardinality(), 1)
		(_nw165).ArraySet1((_dafny.Zero).Minus((_1217_v60).Select((Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1217_v60).Cardinality()))).Uint32()).(_dafny.Int)), 2)
		(_nw165).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus(_1150_v0))), 3)
		(_nw165).ArraySet1(Companion_Default___.Fm0(_dafny.SeqOf((_this).F13()), true, (_this).F17(), _dafny.IntOfInt64(384), globalState), 4)
		(_nw165).ArraySet1(_1150_v0, 5)
		(_nw165).ArraySet1(_1150_v0, 6)
		(_nw165).ArraySet1(_1150_v0, 7)
		(_nw165).ArraySet1(_1150_v0, 8)
		_1218_v61 = _nw165
		var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
		_ = _index178
		(_1218_v61).ArraySet1((_1150_v0).Plus(_1150_v0), (_index178).Int())
		var _1219_v62 _dafny.MultiSet
		_ = _1219_v62
		_1219_v62 = _dafny.MultiSetOf((_this).F14())
		var _1220_v63 bool
		_ = _1220_v63
		_1220_v63 = true
		var _1221_v64 _dafny.Sequence
		_ = _1221_v64
		_1221_v64 = _dafny.UnicodeSeqOfUtf8Bytes("paa")
		var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_1218_v61), 0))
		_ = _index179
		(_1218_v61).ArraySet1(_dafny.IntOfUint32((_1221_v64).Cardinality()), (_index179).Int())
		var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
		_ = _index180
		var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_1218_v61), 0))
		_ = _index181
		var _rhs149 _dafny.Int = _1150_v0
		_ = _rhs149
		var _rhs150 _dafny.MultiSet = _1219_v62
		_ = _rhs150
		var _rhs151 bool = p1
		_ = _rhs151
		var _rhs152 bool = false
		_ = _rhs152
		var _rhs153 _dafny.Int = _1150_v0
		_ = _rhs153
		var _lhs122 _dafny.Array = _1218_v61
		_ = _lhs122
		var _lhs123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
		_ = _lhs123
		var _lhs124 _dafny.Array = _1218_v61
		_ = _lhs124
		var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(70), _dafny.ArrayLen((_1218_v61), 0))
		_ = _lhs125
		(_lhs122).ArraySet1(_rhs149, (_lhs123).Int())
		_1219_v62 = _rhs150
		_1220_v63 = _rhs151
		_1220_v63 = _rhs152
		(_lhs124).ArraySet1(_rhs153, (_lhs125).Int())
		var _1222_v65 _dafny.Sequence
		_ = _1222_v65
		_1222_v65 = _dafny.SeqOf(_1220_v63)
		var _1223_v66 _dafny.Map
		_ = _1223_v66
		_1223_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), (_1222_v65).Select((Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1222_v65).Cardinality()))).Uint32()).(bool))
		if (func() bool {
			if (_1223_v66).Contains(_1150_v0) {
				return (_1223_v66).Get(_1150_v0).(bool)
			}
			return _1220_v63
		})() {
			(globalState).F7 = (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)
			var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
			_ = _index182
			(_1218_v61).ArraySet1(_dafny.IntOfUint32((_1217_v60).Cardinality()), (_index182).Int())
			var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
			_ = _index183
			(_1218_v61).ArraySet1(Companion_Default___.SafeDivisionInt((_this).Fm15(globalState), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)), (_index183).Int())
			var _1224_v67 _dafny.Map
			_ = _1224_v67
			_1224_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm21(globalState), _1220_v63)
			_1224_v67 = (_1224_v67).Update(_this.F16(), (_dafny.IntOfInt64(810)).Cmp(_1150_v0) < 0)
			(globalState).F7 = Companion_Default___.SafeModuloInt((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _1150_v0)
		} else {
			var _1225_v68 _dafny.Sequence
			_ = _1225_v68
			_1225_v68 = _dafny.SeqOf(_1221_v64)
			var _1226_v69 _dafny.Array
			_ = _1226_v69
			var _nwElement0_41 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1225_v68, _1225_v68)
			_ = _nwElement0_41
			var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(3))
			_ = _nw166
			(_nw166).ArraySet1(_nwElement0_41, 0)
			(_nw166).ArraySet1(_1225_v68, 1)
			(_nw166).ArraySet1(_dafny.SeqOf(_1221_v64), 2)
			_1226_v69 = _nw166
			var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_1226_v69), 0))
			_ = _index184
			(_1226_v69).ArraySet1(_1225_v68, (_index184).Int())
			var _1227_v70 D6
			_ = _1227_v70
			_1227_v70 = Companion_D6_.Create_DC11_(_dafny.MultiSetOf((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)))
			var _1228_v71 _dafny.Map
			_ = _1228_v71
			_1228_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1227_v70, _1221_v64)
			var _1229_v72 D6
			_ = _1229_v72
			_1229_v72 = Companion_D6_.Create_DC12_((_this).F13(), (func() _dafny.Sequence {
				if (_1228_v71).Contains(_1227_v70) {
					return (_1228_v71).Get(_1227_v70).(_dafny.Sequence)
				}
				return _1221_v64
			})(), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _1220_v63)
			var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_1226_v69), 0))
			_ = _index185
			(_1226_v69).ArraySet1(Companion_Default___.Fm45(_1229_v72, _1150_v0, globalState), (_index185).Int())
			var _1230_v73 _dafny.Map
			_ = _1230_v73
			_1230_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(332), _dafny.IntOfInt64(-981))
			var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index186
			((_this).F14()).ArraySet1(!(Companion_Default___.Fm30((_1230_v73).Cardinality(), globalState)), (_index186).Int())
			var _1231_v74 _dafny.CodePoint
			_ = _1231_v74
			_1231_v74 = _dafny.CodePoint('y')
			var _1232_v75 D9
			_ = _1232_v75
			_1232_v75 = Companion_D9_.Create_DC22_((_this).F17(), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), (_this).F18(), Companion_Default___.Fm36(p2, _dafny.CodePoint('k'), p1, globalState), _dafny.Companion_Sequence_.Update(_1221_v64, (Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1221_v64).Cardinality()))).Uint32(), _1231_v74))
			var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index187
			((_this).F14()).ArraySet1((_1232_v75).Dtor_cf37(), (_index187).Int())
			_1150_v0 = (_1150_v0).Plus(_dafny.IntOfInt64(398))
			var _1233_v76 _dafny.Map
			_ = _1233_v76
			_1233_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1220_v63, _1150_v0)
			var _1234_v77 _dafny.MultiSet
			_ = _1234_v77
			_1234_v77 = _dafny.MultiSetOf((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(32), (_1233_v76).Cardinality(), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int))
			var _1235_v78 _dafny.MultiSet
			_ = _1235_v78
			_1235_v78 = _dafny.MultiSetOf((_1234_v77).Cardinality())
			var _1236_v79 _dafny.Map
			_ = _1236_v79
			_1236_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(121), _1235_v78)
			var _1237_v80 _dafny.Sequence
			_ = _1237_v80
			_1237_v80 = _dafny.SeqOf(_1233_v76, _1233_v76, _1233_v76, _1233_v76)
			var _1238_v81 _dafny.MultiSet
			_ = _1238_v81
			_1238_v81 = _dafny.MultiSetOf(_1233_v76, _1233_v76, _1233_v76, (_1237_v80).Select((Companion_Default___.SafeIndex((_1229_v72).Dtor_cf20(), _dafny.IntOfUint32((_1237_v80).Cardinality()))).Uint32()).(_dafny.Map))
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
			_ = _index188
			var _rhs154 _dafny.Int = (_dafny.Zero).Minus(((_this).Fm8(_1220_v63, _this.F16(), globalState)).Minus(((_1236_v79).Merge((_1236_v79).Update((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _1235_v78))).Cardinality()))
			_ = _rhs154
			var _rhs155 _dafny.Int = ((Companion_Default___.Fm0(_1222_v65, (_this).F18(), (_this).F18(), (_1238_v81).Cardinality(), globalState)).Plus(_1150_v0)).Times((_dafny.Zero).Minus((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)))
			_ = _rhs155
			var _lhs126 *GlobalState = globalState
			_ = _lhs126
			var _lhs127 _dafny.Array = _1218_v61
			_ = _lhs127
			var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))
			_ = _lhs128
			_lhs126.F7 = _rhs154
			(_lhs127).ArraySet1(_rhs155, (_lhs128).Int())
			var _1239_v82 D6
			_ = _1239_v82
			_1239_v82 = Companion_D6_.Create_DC13_(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_this).F18(), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), p2)
			var _1240_v83 D6
			_ = _1240_v83
			_1240_v83 = Companion_D6_.Create_DC14_(_1239_v82)
			_1221_v64 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
					return Companion_Default___.Fm32(_1240_v83, globalState)
				}
				return _1221_v64
			})(), (_1225_v68).Select((Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((_1225_v68).Cardinality()))).Uint32()).(_dafny.Sequence))
		}
		var _1241_v84 _dafny.Map
		_ = _1241_v84
		_1241_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_this).Fm8(false, _this.F16(), globalState))
		var _1242_v85 _dafny.Sequence
		_ = _1242_v85
		_1242_v85 = _dafny.SeqOf((_1241_v84).Merge(_1241_v84))
		var _1243_v86 _dafny.Map
		_ = _1243_v86
		_1243_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _dafny.SeqOf(_1241_v84))
		_1242_v85 = _dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_1243_v86).Contains(_1150_v0) {
				return (_1243_v86).Get(_1150_v0).(_dafny.Sequence)
			}
			return _dafny.SeqOf((_1242_v85).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1242_v85).Cardinality()))).Uint32()).(_dafny.Map))
		})(), _1242_v85)
		var _1244_i2 _dafny.Int
		_ = _1244_i2
		_1244_i2 = _dafny.Zero
		{
			for p1 {
				{
					if (_1244_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1244_i2 = (_1244_i2).Plus(_dafny.One)
					var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((p0), 0))
					_ = _index189
					(p0).ArraySet1(false, (_index189).Int())
					var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((p0), 0))
					_ = _index190
					(p0).ArraySet1(p2, (_index190).Int())
					var _1245_v87 _dafny.MultiSet
					_ = _1245_v87
					_1245_v87 = _dafny.MultiSetOf((_this).F18(), true)
					var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((p0), 0))
					_ = _index191
					var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((p0), 0))
					_ = _index192
					var _rhs156 bool = (_this).F13()
					_ = _rhs156
					var _rhs157 bool = true
					_ = _rhs157
					var _rhs158 bool = (_this).F18()
					_ = _rhs158
					var _rhs159 bool = (_dafny.Companion_Sequence_.IsPrefixOf(_1217_v60, _dafny.Companion_Sequence_.Update(Companion_Default___.Fm34((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), globalState), (Companion_Default___.SafeIndex(_1150_v0, _dafny.IntOfUint32((Companion_Default___.Fm34((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), globalState)).Cardinality()))).Uint32(), (_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)))) && ((_dafny.IntOfUint32((_1221_v64).Cardinality())).Cmp((_1245_v87).Cardinality()) >= 0)
					_ = _rhs159
					var _rhs160 _dafny.Int = (_1217_v60).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
						if (_this).F13() {
							return _dafny.IntOfInt64(250)
						}
						return _1150_v0
					})(), _dafny.IntOfUint32((_1217_v60).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs160
					var _lhs129 _dafny.Array = p0
					_ = _lhs129
					var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((p0), 0))
					_ = _lhs130
					var _lhs131 _dafny.Array = p0
					_ = _lhs131
					var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(917), _dafny.ArrayLen((p0), 0))
					_ = _lhs132
					(_lhs129).ArraySet1(_rhs156, (_lhs130).Int())
					(_lhs131).ArraySet1(_rhs157, (_lhs132).Int())
					_1220_v63 = _rhs158
					_1220_v63 = _rhs159
					_1150_v0 = _rhs160
					var _1246_v88 _dafny.Array
					_ = _1246_v88
					var _len0_29 _dafny.Int = _dafny.IntOfInt64(6)
					_ = _len0_29
					var _nw167 _dafny.Array
					_ = _nw167
					if _len0_29.Cmp(_dafny.Zero) == 0 {
						_nw167 = _dafny.NewArray(_len0_29)
					} else {
						var _init29 func(_dafny.Int) D6 = (func(_1247_v0 _dafny.Int) func(_dafny.Int) D6 {
							return func(_1248_i3 _dafny.Int) D6 {
								return Companion_D6_.Create_DC11_(_dafny.MultiSetOf(_dafny.IntOfInt64(313), _1247_v0))
							}
						})(_1150_v0)
						_ = _init29
						var _element0_29 = _init29(_dafny.Zero)
						_ = _element0_29
						_nw167 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
						(_nw167).ArraySet1(_element0_29, 0)
						var _nativeLen0_29 = (_len0_29).Int()
						_ = _nativeLen0_29
						for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
							(_nw167).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
						}
					}
					_1246_v88 = _nw167
					_1246_v88 = _1246_v88
					(globalState).F2 = _1222_v65
					(globalState).F7 = (_dafny.IntOfInt64(-41)).Times(((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int)).Times(_1150_v0))
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1249_v89 _dafny.Map
		_ = _1249_v89
		_1249_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1150_v0, _1221_v64)
		_1249_v89 = (_1249_v89).Update((_1218_v61).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(210), _dafny.ArrayLen((_1218_v61), 0))).Int()).(_dafny.Int), _1221_v64)
	}
}
func (_this *C3) M12(globalState *GlobalState) {
	{
		var _1250_v0 _dafny.CodePoint
		_ = _1250_v0
		_1250_v0 = _dafny.CodePoint('j')
		var _1251_v1 *C0
		_ = _1251_v1
		var _nw168 *C0 = New_C0_()
		_ = _nw168
		_nw168.Ctor__(true, (func() _dafny.CodePoint {
			if (_this).F13() {
				return _dafny.CodePoint('p')
			}
			return _1250_v0
		})(), _this.F12(), (_this).F18())
		_1251_v1 = _nw168
		var _1252_v2 _dafny.Int
		_ = _1252_v2
		_1252_v2 = _dafny.IntOfInt64(385)
		(globalState).F7 = _1252_v2
		(_1251_v1).F19 = (_this).F18()
		var _1253_v3 _dafny.Sequence
		_ = _1253_v3
		_1253_v3 = _dafny.UnicodeSeqOfUtf8Bytes("j")
		var _1254_v4 _dafny.Map
		_ = _1254_v4
		_1254_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_v3, (_this).F18())
		_1254_v4 = _1254_v4
		_1253_v3 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer75 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg75 _dafny.Int) interface{} {
				return coer75(arg75)
			}
		}((func(_1255_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1256_i0 _dafny.Int) _dafny.CodePoint {
				return _1255_v0
			}
		})(_1250_v0)))
		var _hi6 _dafny.Int = _1252_v2
		_ = _hi6
		for _1257_i1 := _1252_v2; _1257_i1.Cmp(_hi6) < 0; _1257_i1 = _1257_i1.Plus(_dafny.One) {
			(globalState).F7 = Companion_Default___.SafeModuloInt(_1257_i1, _dafny.IntOfInt64(-145))
			var _1258_v5 _dafny.Array
			_ = _1258_v5
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw169
			_1258_v5 = _nw169
			var _1259_v6 _dafny.Set
			_ = _1259_v6
			_1259_v6 = _dafny.SetOf(_1251_v1.F19, _1251_v1.F19)
			var _1260_v7 _dafny.Sequence
			_ = _1260_v7
			_1260_v7 = _dafny.SeqOf((_1259_v6).Cardinality(), _1252_v2, _1252_v2, _1252_v2)
			var _1261_v8 _dafny.Map
			_ = _1261_v8
			_1261_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1258_v5, _1260_v7)
			var _1262_v9 _dafny.Sequence
			_ = _1262_v9
			_1262_v9 = _dafny.SeqOf(_1252_v2, (_1261_v8).Cardinality())
			var _1263_v10 D9
			_ = _1263_v10
			_1263_v10 = Companion_D9_.Create_DC22_(_1251_v1.F19, (_1260_v7).Select((Companion_Default___.SafeIndex(_1252_v2, _dafny.IntOfUint32((_1260_v7).Cardinality()))).Uint32()).(_dafny.Int), (_this).F13(), _1259_v6, _1253_v3)
			var _1264_v11 _dafny.Map
			_ = _1264_v11
			_1264_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1252_v2).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F18()), (Companion_Default___.SafeIndex(_1257_i1, _dafny.IntOfUint32((_dafny.SeqOf((_this).F18())).Cardinality()))).Uint32(), (_this).F13())).Cardinality())), (_1263_v10).Dtor_cf38())
			_1264_v11 = _1264_v11
			(globalState).F7 = _1257_i1
			if ((func() _dafny.Set {
				var _coll44 = _dafny.NewBuilder()
				_ = _coll44
				for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(892), _dafny.IntOfInt64(210))); ; {
					_compr_44, _ok52 := _iter52()
					if !_ok52 {
						break
					}
					var _1265_v12 _dafny.Int
					_1265_v12 = interface{}(_compr_44).(_dafny.Int)
					if ((_dafny.IntOfInt64(892)).Cmp(_1265_v12) <= 0) && ((_1265_v12).Cmp(_dafny.IntOfInt64(210)) < 0) {
						_coll44.Add((_1265_v12).Times(_1252_v2))
					}
				}
				return _coll44.ToSet()
			}()).Difference((_this).F15())).Equals(((_this).F15()).Difference((_this).F15())) {
				var _1266_v13 D3
				_ = _1266_v13
				_1266_v13 = Companion_D3_.Create_DC6_(_1257_i1, _1251_v1.F19)
				var _1267_v14 _dafny.Array
				_ = _1267_v14
				var _nwElement0_42 _dafny.CodePoint = (_1251_v1).F20()
				_ = _nwElement0_42
				var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(29))
				_ = _nw170
				(_nw170).ArraySet1CodePoint(_nwElement0_42, 0)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 1)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 2)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 3)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 4)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 5)
				(_nw170).ArraySet1CodePoint(_1250_v0, 6)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 7)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 8)
				(_nw170).ArraySet1CodePoint(_dafny.CodePoint('y'), 9)
				(_nw170).ArraySet1CodePoint(_1250_v0, 10)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 11)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 12)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 13)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 14)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 15)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 16)
				(_nw170).ArraySet1CodePoint(Companion_Default___.Fm18(_1252_v2, _1253_v3, (_this).F17(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_dafny.Zero).Minus(_1257_i1)), globalState), 17)
				(_nw170).ArraySet1CodePoint(Companion_Default___.Fm33((_1266_v13).Dtor_cf13(), _1257_i1, globalState), 18)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 19)
				(_nw170).ArraySet1CodePoint(_dafny.CodePoint('n'), 20)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 21)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 22)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 23)
				(_nw170).ArraySet1CodePoint(_1250_v0, 24)
				(_nw170).ArraySet1CodePoint((_1253_v3).Select((Companion_Default___.SafeIndex(_1257_i1, _dafny.IntOfUint32((_1253_v3).Cardinality()))).Uint32()).(_dafny.CodePoint), 25)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 26)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 27)
				(_nw170).ArraySet1CodePoint((_1251_v1).F20(), 28)
				_1267_v14 = _nw170
				var _1268_v15 D3
				_ = _1268_v15
				_1268_v15 = Companion_D3_.Create_DC5_((_1251_v1).F20())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1267_v14), 0))
				_ = _index193
				(_1267_v14).ArraySet1CodePoint((_1268_v15).Dtor_cf11(), (_index193).Int())
				var _1269_v16 _dafny.Sequence
				_ = _1269_v16
				_1269_v16 = _dafny.SeqOf(_1251_v1.F19, !((_this).F17()), Companion_Default___.Fm30(_1252_v2, globalState))
				var _1270_v17 _dafny.MultiSet
				_ = _1270_v17
				_1270_v17 = _dafny.MultiSetOf((_this).F17())
				var _1271_v18 _dafny.Map
				_ = _1271_v18
				_1271_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1251_v1.F19, (_1270_v17).Cardinality())
				var _1272_v19 _dafny.Sequence
				_ = _1272_v19
				_1272_v19 = _dafny.SeqOf((_this).F18(), (_this).F13(), !(_1251_v1.F19), (_this).F13(), (_1269_v16).Select((Companion_Default___.SafeIndex((_1271_v18).Cardinality(), _dafny.IntOfUint32((_1269_v16).Cardinality()))).Uint32()).(bool))
				var _1273_v20 _dafny.Sequence
				_ = _1273_v20
				_1273_v20 = _dafny.SeqOf(_1269_v16, _1272_v19)
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1267_v14), 0))
				_ = _index194
				var _rhs161 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1269_v16, (_1273_v20).Select((Companion_Default___.SafeIndex(_1257_i1, _dafny.IntOfUint32((_1273_v20).Cardinality()))).Uint32()).(_dafny.Sequence))
				_ = _rhs161
				var _rhs162 bool = (_1263_v10).Dtor_cf39()
				_ = _rhs162
				var _rhs163 _dafny.CodePoint = _1250_v0
				_ = _rhs163
				var _rhs164 _dafny.CodePoint = (_1251_v1).F20()
				_ = _rhs164
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				var _lhs134 *C0 = _1251_v1
				_ = _lhs134
				var _lhs135 _dafny.Array = _1267_v14
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(287), _dafny.ArrayLen((_1267_v14), 0))
				_ = _lhs136
				_lhs133.F2 = _rhs161
				_lhs134.F19 = _rhs162
				(_lhs135).ArraySet1CodePoint(_rhs163, (_lhs136).Int())
				_1250_v0 = _rhs164
				(_1251_v1).F19 = (_this).F17()
				var _1274_v21 _dafny.Map
				_ = _1274_v21
				_1274_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1257_i1, _1251_v1.F19)
				(_1251_v1).F19 = (func() bool {
					if (_1274_v21).Contains((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1272_v19).Cardinality()), _1257_i1)).Cardinality()) {
						return (_1274_v21).Get((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1272_v19).Cardinality()), _1257_i1)).Cardinality()).(bool)
					}
					return _dafny.Companion_Sequence_.Equal(_1260_v7, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(230))).Uint32(), func(coer76 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg76 _dafny.Int) interface{} {
							return coer76(arg76)
						}
					}((func(_1275_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
						return func(_1276_i2 _dafny.Int) _dafny.Int {
							return _dafny.IntOfUint32((_1275_v3).Cardinality())
						}
					})(_1253_v3))))
				})()
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index195
				((_this).F14()).ArraySet1(_1251_v1.F19, (_index195).Int())
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(78), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index196
				((_this).F14()).ArraySet1(!(!((_1252_v2).Cmp(_dafny.IntOfInt64(509)) <= 0)) || ((_1269_v16).Select((Companion_Default___.SafeIndex(_1252_v2, _dafny.IntOfUint32((_1269_v16).Cardinality()))).Uint32()).(bool)), (_index196).Int())
				(_1251_v1).F19 = ((_this).F18()) && (_1251_v1.F19)
			} else {
				(_1251_v1).F19 = (_this.F12()).Dtor_cf4()
				(globalState).F7 = ((_1257_i1).Plus(_dafny.IntOfUint32((_1253_v3).Cardinality()))).Times(_1252_v2)
				var _1277_v22 _dafny.Sequence
				_ = _1277_v22
				_1277_v22 = _dafny.SeqOf(_1253_v3)
				var _rhs165 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1253_v3, _1253_v3)
				_ = _rhs165
				var _rhs166 _dafny.Int = _1257_i1
				_ = _rhs166
				var _rhs167 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1277_v22, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-342))).Uint32(), func(coer77 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg77 _dafny.Int) interface{} {
						return coer77(arg77)
					}
				}((func(_1278_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1279_i3 _dafny.Int) _dafny.Sequence {
						return _1278_v3
					}
				})(_1253_v3))))
				_ = _rhs167
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				_1253_v3 = _rhs165
				_lhs137.F7 = _rhs166
				_1277_v22 = _rhs167
				(_1251_v1).F19 = _1251_v1.F19
				_1253_v3 = _1253_v3
			}
		}
	}
}
func (_this *C3) F18() bool {
	{
		return _this._f18
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) Ctor__() {
	{
	}
}
func (_this *C4) Fm14(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bq"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(963))).Uint32(), func(coer78 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg78 _dafny.Int) interface{} {
				return coer78(arg78)
			}
		}(func(_1280_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('k')
		})))
	}
}
func (_this *C4) M10(p0 bool, p1 _dafny.Sequence, p2 _dafny.CodePoint, globalState *GlobalState) (_dafny.Sequence, _dafny.Sequence) {
	{
		var r0 _dafny.Sequence = _dafny.EmptySeq
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var _1281_v0 _dafny.Sequence
		_ = _1281_v0
		_1281_v0 = _dafny.UnicodeSeqOfUtf8Bytes("gwx")
		_1281_v0 = _1281_v0
		var _1282_v1 _dafny.Array
		_ = _1282_v1
		var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
		_ = _nw171
		_1282_v1 = _nw171
		var _1283_v2 _dafny.Sequence
		_ = _1283_v2
		_1283_v2 = _dafny.SeqOf(_1282_v1)
		var _hi7 _dafny.Int = _dafny.IntOfUint32((_1283_v2).Cardinality())
		_ = _hi7
		for _1284_i0 := (func() _dafny.Int {
			if p0 {
				return _dafny.IntOfInt64(450)
			}
			return _dafny.IntOfInt64(-597)
		})(); _1284_i0.Cmp(_hi7) < 0; _1284_i0 = _1284_i0.Plus(_dafny.One) {
			var _1285_v3 _dafny.Sequence
			_ = _1285_v3
			_1285_v3 = _dafny.SeqOf(_1284_i0)
			if _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(-808)), _1285_v3) {
				var _1286_v4 bool
				_ = _1286_v4
				_1286_v4 = false
				_1286_v4 = _1286_v4
				var _1287_v5 _dafny.Array
				_ = _1287_v5
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
				_ = _nw172
				_1287_v5 = _nw172
				_1287_v5 = _1287_v5
				var _1288_v6 D1
				_ = _1288_v6
				_1288_v6 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(130), _1284_i0, p0)
				var _1289_v7 *C1
				_ = _1289_v7
				var _nw173 *C1 = New_C1_()
				_ = _nw173
				_nw173.Ctor__(p2, _1288_v6, p0)
				_1289_v7 = _nw173
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1287_v5), 0))
				_ = _index197
				(_1287_v5).ArraySet1(!(!(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(!(_1286_v4)), p0))), (_index197).Int())
				var _1290_v8 _dafny.Set
				_ = _1290_v8
				_1290_v8 = _dafny.SetOf(_1286_v4)
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1287_v5), 0))
				_ = _index198
				var _rhs168 bool = ((_1290_v8).Cardinality()).Cmp(_dafny.IntOfUint32((_1281_v0).Cardinality())) >= 0
				_ = _rhs168
				var _rhs169 _dafny.Sequence = _1281_v0
				_ = _rhs169
				var _rhs170 _dafny.Int = (_1285_v3).Select((Companion_Default___.SafeIndex(_1284_i0, _dafny.IntOfUint32((_1285_v3).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs170
				var _lhs138 _dafny.Array = _1287_v5
				_ = _lhs138
				var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1287_v5), 0))
				_ = _lhs139
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				(_lhs138).ArraySet1(_rhs168, (_lhs139).Int())
				_1281_v0 = _rhs169
				_lhs140.F7 = _rhs170
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(645), _dafny.ArrayLen((_1287_v5), 0))
				_ = _index199
				(_1287_v5).ArraySet1(!(p0), (_index199).Int())
			} else {
				var _1291_v9 _dafny.Sequence
				_ = _1291_v9
				_1291_v9 = _dafny.SeqOf(p0)
				var _1292_v10 D1
				_ = _1292_v10
				_1292_v10 = Companion_D1_.Create_DC1_(_1281_v0)
				var _1293_v11 _dafny.Array
				_ = _1293_v11
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
				_ = _nw174
				_1293_v11 = _nw174
				var _1294_v12 _dafny.Set
				_ = _1294_v12
				_1294_v12 = _dafny.SetOf(_1284_i0)
				var _pat_let_tv33 = _1281_v0
				_ = _pat_let_tv33
				var _1295_v13 *C2
				_ = _1295_v13
				var _nw175 *C2 = New_C2_()
				_ = _nw175
				_nw175.Ctor__(_1284_i0, (_1291_v9).Select((Companion_Default___.SafeIndex(_1284_i0, _dafny.IntOfUint32((_1291_v9).Cardinality()))).Uint32()).(bool), func(_pat_let33_0 D1) D1 {
					return func(_1296_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let34_0 _dafny.Sequence) D1 {
							return func(_1297_dt__update_hcf1_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC1_(_1297_dt__update_hcf1_h0)
							}(_pat_let34_0)
						}(_pat_let_tv33)
					}(_pat_let33_0)
				}(_1292_v10), p0, _1293_v11, _1294_v12, Companion_D1_.Create_DC2_(_1284_i0, _1284_i0, p0), !(p0) || (p0))
				_1295_v13 = _nw175
				var _1298_v14 _dafny.Set
				_ = _1298_v14
				_1298_v14 = _dafny.SetOf(true)
				var _1299_v15 D9
				_ = _1299_v15
				_1299_v15 = Companion_D9_.Create_DC22_(p0, (_1295_v13).F21(), p0, _1298_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer79 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg79 _dafny.Int) interface{} {
						return coer79(arg79)
					}
				}((func(_1300_p2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1301_i1 _dafny.Int) _dafny.CodePoint {
						return _1300_p2
					}
				})(p2))))
				var _1302_v16 _dafny.Set
				_ = _1302_v16
				_1302_v16 = _dafny.SetOf(_1299_v15, _1299_v15, Companion_D9_.Create_DC22_(p0, (_1295_v13).F21(), p0, _1298_v14, _dafny.UnicodeSeqOfUtf8Bytes("m")))
				var _1303_v17 D3
				_ = _1303_v17
				_1303_v17 = Companion_D3_.Create_DC5_(p2)
				var _1304_v18 _dafny.Array
				_ = _1304_v18
				var _nwElement0_43 _dafny.CodePoint = p2
				_ = _nwElement0_43
				var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(16))
				_ = _nw176
				(_nw176).ArraySet1CodePoint(_nwElement0_43, 0)
				(_nw176).ArraySet1CodePoint((_1303_v17).Dtor_cf11(), 1)
				(_nw176).ArraySet1CodePoint(p2, 2)
				(_nw176).ArraySet1CodePoint(p2, 3)
				(_nw176).ArraySet1CodePoint(p2, 4)
				(_nw176).ArraySet1CodePoint(p2, 5)
				(_nw176).ArraySet1CodePoint(p2, 6)
				(_nw176).ArraySet1CodePoint(p2, 7)
				(_nw176).ArraySet1CodePoint(p2, 8)
				(_nw176).ArraySet1CodePoint(_dafny.CodePoint('t'), 9)
				(_nw176).ArraySet1CodePoint(p2, 10)
				(_nw176).ArraySet1CodePoint(p2, 11)
				(_nw176).ArraySet1CodePoint(_dafny.CodePoint('g'), 12)
				(_nw176).ArraySet1CodePoint(p2, 13)
				(_nw176).ArraySet1CodePoint(p2, 14)
				(_nw176).ArraySet1CodePoint(p2, 15)
				_1304_v18 = _nw176
				var _1305_v19 _dafny.Map
				_ = _1305_v19
				_1305_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1304_v18, _dafny.IntOfUint32((_1285_v3).Cardinality()))
				var _1306_v20 _dafny.MultiSet
				_ = _1306_v20
				_1306_v20 = _dafny.MultiSetOf((_1302_v16).Cardinality(), (func() _dafny.Int {
					if (_1305_v19).Contains(_1304_v18) {
						return (_1305_v19).Get(_1304_v18).(_dafny.Int)
					}
					return _1284_i0
				})())
				var _1307_v21 D7
				_ = _1307_v21
				_1307_v21 = Companion_D7_.Create_DC16_((_1295_v13).F21())
				var _1308_v22 _dafny.Map
				_ = _1308_v22
				_1308_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1306_v20)
				var _1309_v23 D6
				_ = _1309_v23
				_1309_v23 = Companion_D6_.Create_DC11_(_dafny.MultiSetOf(_1284_i0, _dafny.IntOfUint32((_1281_v0).Cardinality()), (_1295_v13).F21()))
				var _1310_v24 _dafny.Array
				_ = _1310_v24
				var _nwElement0_44 _dafny.MultiSet = _1306_v20
				_ = _nwElement0_44
				var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(22))
				_ = _nw177
				(_nw177).ArraySet1(_nwElement0_44, 0)
				(_nw177).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1284_i0, (_1298_v14).Cardinality())), 1)
				(_nw177).ArraySet1(_1306_v20, 2)
				(_nw177).ArraySet1((_1306_v20).Intersection(_1306_v20), 3)
				(_nw177).ArraySet1((_1306_v20).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(981), _1284_i0, _dafny.IntOfUint32((_dafny.SeqOf(!((_1295_v13).F22()))).Cardinality()))).Cardinality()), (_1307_v21).Dtor_cf29())), 4)
				(_nw177).ArraySet1(_1306_v20, 5)
				(_nw177).ArraySet1(_1306_v20, 6)
				(_nw177).ArraySet1(_1306_v20, 7)
				(_nw177).ArraySet1(_1306_v20, 8)
				(_nw177).ArraySet1(_1306_v20, 9)
				(_nw177).ArraySet1(_1306_v20, 10)
				(_nw177).ArraySet1((_dafny.MultiSetOf(_1284_i0)).Difference(_1306_v20), 11)
				(_nw177).ArraySet1(_1306_v20, 12)
				(_nw177).ArraySet1((func() _dafny.MultiSet {
					if (_1308_v22).Contains((_1295_v13).F22()) {
						return (_1308_v22).Get((_1295_v13).F22()).(_dafny.MultiSet)
					}
					return _1306_v20
				})(), 13)
				(_nw177).ArraySet1(_1306_v20, 14)
				(_nw177).ArraySet1((_1306_v20).Union((_1306_v20).Update((_1306_v20).Cardinality(), Companion_Default___.Abs(_dafny.IntOfInt64(473)))), 15)
				(_nw177).ArraySet1(_1306_v20, 16)
				(_nw177).ArraySet1((_1306_v20).Intersection(_1306_v20), 17)
				(_nw177).ArraySet1((_1306_v20).Intersection(_1306_v20), 18)
				(_nw177).ArraySet1(_1306_v20, 19)
				(_nw177).ArraySet1(Companion_Default___.Fm19(globalState), 20)
				(_nw177).ArraySet1(((_1309_v23).Dtor_cf17()).Difference(_dafny.MultiSetFromSeq(_1285_v3)), 21)
				_1310_v24 = _nw177
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_1310_v24), 0))
				_ = _index200
				(_1310_v24).ArraySet1(_1306_v20, (_index200).Int())
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(506), _dafny.ArrayLen((_1310_v24), 0))
				_ = _index201
				(_1310_v24).ArraySet1(_1306_v20, (_index201).Int())
				var _1311_v25 bool
				_ = _1311_v25
				_1311_v25 = false
				var _1312_v26 _dafny.Array
				_ = _1312_v26
				var _nw178 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw178
				_1312_v26 = _nw178
				var _1313_v27 _dafny.Array
				_ = _1313_v27
				var _nw179 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(23))
				_ = _nw179
				_1313_v27 = _nw179
				var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_1312_v26), 0))
				_ = _index202
				(_1312_v26).ArraySet1(_1313_v27, (_index202).Int())
				var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _index203
				(_1293_v11).ArraySet1((_1295_v13).F22(), (_index203).Int())
				var _1314_v28 _dafny.Map
				_ = _1314_v28
				_1314_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1284_i0).Minus(_dafny.IntOfInt64(-417)), _1313_v27)
				var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_1312_v26), 0))
				_ = _index204
				var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _index205
				var _rhs171 bool = _dafny.Companion_Sequence_.Contains(_1281_v0, p2)
				_ = _rhs171
				var _rhs172 _dafny.Array = (func() _dafny.Array {
					if (_1314_v28).Contains(_dafny.IntOfInt64(183)) {
						return (_1314_v28).Get(_dafny.IntOfInt64(183)).(_dafny.Array)
					}
					return _1313_v27
				})()
				_ = _rhs172
				var _rhs173 bool = false
				_ = _rhs173
				var _rhs174 _dafny.Array = _1282_v1
				_ = _rhs174
				var _lhs141 _dafny.Array = _1312_v26
				_ = _lhs141
				var _lhs142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(460), _dafny.ArrayLen((_1312_v26), 0))
				_ = _lhs142
				var _lhs143 _dafny.Array = _1293_v11
				_ = _lhs143
				var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _lhs144
				_1311_v25 = _rhs171
				(_lhs141).ArraySet1(_rhs172, (_lhs142).Int())
				(_lhs143).ArraySet1(_rhs173, (_lhs144).Int())
				_1282_v1 = _rhs174
				var _1315_v29 _dafny.Map
				_ = _1315_v29
				_1315_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1295_v13).F22(), (_1293_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))).Int()).(bool))
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _index206
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _index207
				var _rhs175 bool = (_1293_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))).Int()).(bool)
				_ = _rhs175
				var _rhs176 _dafny.MultiSet = _1306_v20
				_ = _rhs176
				var _rhs177 bool = !(((_1295_v13).F21()).Cmp((_1295_v13).F21()) <= 0)
				_ = _rhs177
				var _rhs178 bool = !((func() bool {
					if (_1295_v13).F22() {
						return !_dafny.Companion_Sequence_.Contains(_1285_v3, _1284_i0)
					}
					return (func() bool {
						if (_1315_v29).Contains(_1311_v25) {
							return (_1315_v29).Get(_1311_v25).(bool)
						}
						return (_1293_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))).Int()).(bool)
					})()
				})())
				_ = _rhs178
				var _lhs145 _dafny.Array = _1293_v11
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _lhs146
				var _lhs147 _dafny.Array = _1293_v11
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))
				_ = _lhs148
				(_lhs145).ArraySet1(_rhs175, (_lhs146).Int())
				_1306_v20 = _rhs176
				(_lhs147).ArraySet1(_rhs177, (_lhs148).Int())
				_1311_v25 = _rhs178
				var _1316_v30 _dafny.MultiSet
				_ = _1316_v30
				_1316_v30 = _dafny.MultiSetOf((_1295_v13).F22())
				var _1317_v31 D1
				_ = _1317_v31
				_1317_v31 = Companion_D1_.Create_DC2_(_dafny.IntOfInt64(236), (_1295_v13).F21(), (_1295_v13).F22())
				var _1318_v32 _dafny.Set
				_ = _1318_v32
				_1318_v32 = _dafny.SetOf(_dafny.CodePoint('c'), p2, p2)
				var _pat_let_tv34 = _1281_v0
				_ = _pat_let_tv34
				var _1319_v33 *C3
				_ = _1319_v33
				var _nw180 *C3 = New_C3_()
				_ = _nw180
				_nw180.Ctor__((_1316_v30).IsDisjointFrom(_1316_v30), func(_pat_let35_0 D1) D1 {
					return func(_1320_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let36_0 _dafny.Sequence) D1 {
							return func(_1321_dt__update_hcf1_h1 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC1_(_1321_dt__update_hcf1_h1)
							}(_pat_let36_0)
						}(_pat_let_tv34)
					}(_pat_let35_0)
				}(_1292_v10), (_1295_v13).F22(), _1293_v11, _1294_v12, _1317_v31, (_dafny.SetOf(Companion_Default___.Fm33((_1293_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(743), _dafny.ArrayLen((_1293_v11), 0))).Int()).(bool), (_1295_v13).F21(), globalState))).IsSubsetOf(_1318_v32))
				_1319_v33 = _nw180
			}
			(globalState).F7 = (func() _dafny.Int {
				if p0 {
					return _1284_i0
				}
				return _1284_i0
			})()
			var _1322_v34 D6
			_ = _1322_v34
			_1322_v34 = Companion_D6_.Create_DC12_(p0, _1281_v0, _1284_i0, p0)
			var _1323_v35 _dafny.Set
			_ = _1323_v35
			_1323_v35 = _dafny.SetOf(_dafny.IntOfInt64(235), _1284_i0)
			var _1324_v36 _dafny.Array
			_ = _1324_v36
			var _nwElement0_45 bool = p0
			_ = _nwElement0_45
			var _nw181 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(4))
			_ = _nw181
			(_nw181).ArraySet1(_nwElement0_45, 0)
			(_nw181).ArraySet1((_1322_v34).Dtor_cf21(), 1)
			(_nw181).ArraySet1(p0, 2)
			(_nw181).ArraySet1((_1323_v35).IsProperSubsetOf(_1323_v35), 3)
			_1324_v36 = _nw181
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1324_v36), 0))
			_ = _index208
			(_1324_v36).ArraySet1(!(p0), (_index208).Int())
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1324_v36), 0))
			_ = _index209
			(_1324_v36).ArraySet1(!((p0) || (p0)), (_index209).Int())
			var _1325_v37 D1
			_ = _1325_v37
			_1325_v37 = Companion_D1_.Create_DC2_(_1284_i0, _1284_i0, (_1324_v36).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(230), _dafny.ArrayLen((_1324_v36), 0))).Int()).(bool))
			var _1326_v38 *C1
			_ = _1326_v38
			var _nw182 *C1 = New_C1_()
			_ = _nw182
			_nw182.Ctor__(p2, _1325_v37, p0)
			_1326_v38 = _nw182
		}
		var _1327_v39 _dafny.Set
		_ = _1327_v39
		_1327_v39 = _dafny.SetOf(p0)
		var _1328_v40 _dafny.Int
		_ = _1328_v40
		_1328_v40 = _dafny.IntOfInt64(706)
		var _rhs179 _dafny.Set = _1327_v39
		_ = _rhs179
		var _rhs180 _dafny.Int = _1328_v40
		_ = _rhs180
		var _rhs181 _dafny.Int = _1328_v40
		_ = _rhs181
		var _lhs149 *GlobalState = globalState
		_ = _lhs149
		var _lhs150 *GlobalState = globalState
		_ = _lhs150
		_1327_v39 = _rhs179
		_lhs149.F7 = _rhs180
		_lhs150.F7 = _rhs181
		var _1329_v41 D1
		_ = _1329_v41
		_1329_v41 = Companion_D1_.Create_DC1_(_1281_v0)
		var _1330_v42 _dafny.Array
		_ = _1330_v42
		var _len0_30 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_30
		var _nw183 _dafny.Array
		_ = _nw183
		if _len0_30.Cmp(_dafny.Zero) == 0 {
			_nw183 = _dafny.NewArray(_len0_30)
		} else {
			var _init30 func(_dafny.Int) bool = (func(_1331_p0 bool) func(_dafny.Int) bool {
				return func(_1332_i2 _dafny.Int) bool {
					return _1331_p0
				}
			})(p0)
			_ = _init30
			var _element0_30 = _init30(_dafny.Zero)
			_ = _element0_30
			_nw183 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
			(_nw183).ArraySet1(_element0_30, 0)
			var _nativeLen0_30 = (_len0_30).Int()
			_ = _nativeLen0_30
			for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
				(_nw183).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
			}
		}
		_1330_v42 = _nw183
		var _1333_v43 _dafny.Set
		_ = _1333_v43
		_1333_v43 = _dafny.SetOf(_1328_v40, _1328_v40)
		var _1334_v44 D1
		_ = _1334_v44
		_1334_v44 = Companion_D1_.Create_DC2_(_1328_v40, _1328_v40, !(p0))
		var _1335_v45 *C3
		_ = _1335_v45
		var _nw184 *C3 = New_C3_()
		_ = _nw184
		_nw184.Ctor__(!(_1327_v39).Contains(p0), _1329_v41, (func() bool {
			if p0 {
				return p0
			}
			return !(p0)
		})(), _1330_v42, _1333_v43, _1334_v44, true)
		_1335_v45 = _nw184
		var _1336_v46 bool
		_ = _1336_v46
		_1336_v46 = true
		_1336_v46 = p0
		var _hi8 _dafny.Int = _dafny.IntOfInt64(881)
		_ = _hi8
		for _1337_i3 := _1328_v40; _1337_i3.Cmp(_hi8) < 0; _1337_i3 = _1337_i3.Plus(_dafny.One) {
			(globalState).F7 = _1337_i3
			var _1338_v47 _dafny.CodePoint
			_ = _1338_v47
			_1338_v47 = _dafny.CodePoint('f')
			_1338_v47 = p2
			var _1339_v48 _dafny.Array
			_ = _1339_v48
			var _nw185 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(12))
			_ = _nw185
			_1339_v48 = _nw185
			var _rhs182 _dafny.Array = _1339_v48
			_ = _rhs182
			var _rhs183 bool = p0
			_ = _rhs183
			var _rhs184 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_1281_v0, _1281_v0)
			_ = _rhs184
			_1339_v48 = _rhs182
			_1336_v46 = _rhs183
			_1336_v46 = _rhs184
			if (_1335_v45).F18() {
				var _1340_v49 _dafny.Set
				_ = _1340_v49
				_1340_v49 = _dafny.SetOf(p2, p2)
				_1340_v49 = _1340_v49
				(globalState).F7 = _1328_v40
				_1336_v46 = _1336_v46
				var _1341_v50 _dafny.Sequence
				_ = _1341_v50
				_1341_v50 = _dafny.SeqOf(_1336_v46, false)
				var _1342_v51 _dafny.Sequence
				_ = _1342_v51
				_1342_v51 = _dafny.SeqOf((_1328_v40).Cmp(_dafny.IntOfUint32((_1341_v50).Cardinality())) != 0, (_1335_v45).F18(), false)
				(globalState).F2 = _1342_v51
				var _1343_v52 _dafny.MultiSet
				_ = _1343_v52
				_1343_v52 = _dafny.MultiSetOf(_1282_v1)
				var _1344_v53 *C2
				_ = _1344_v53
				var _nw186 *C2 = New_C2_()
				_ = _nw186
				_nw186.Ctor__(_1337_i3, !(!((_1335_v45).F18())) || ((_1335_v45).F18()), _1329_v41, (_1342_v51).Select((Companion_Default___.SafeIndex(_1337_i3, _dafny.IntOfUint32((_1342_v51).Cardinality()))).Uint32()).(bool), _1330_v42, _dafny.SetOf(_1328_v40), _1334_v44, (_1343_v52).IsProperSubsetOf(_1343_v52))
				_1344_v53 = _nw186
			} else {
				(globalState).F7 = _dafny.IntOfInt64(75)
				_1336_v46 = (_1335_v45).F18()
				var _1345_v54 *C0
				_ = _1345_v54
				var _nw187 *C0 = New_C0_()
				_ = _nw187
				_nw187.Ctor__((_1335_v45).F18(), p2, Companion_D1_.Create_DC2_(_1328_v40, _1337_i3, (_1335_v45).F18()), !(Companion_Default___.Fm17(globalState)))
				_1345_v54 = _nw187
				var _1346_v55 _dafny.Array
				_ = _1346_v55
				var _nw188 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
				_ = _nw188
				_1346_v55 = _nw188
				_1346_v55 = _1346_v55
				var _1347_v56 _dafny.Map
				_ = _1347_v56
				_1347_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1345_v54.F19, _1328_v40)
				var _1348_v57 _dafny.Map
				_ = _1348_v57
				_1348_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1347_v56, _dafny.IntOfInt64(-568))
				var _1349_v58 _dafny.Map
				_ = _1349_v58
				_1349_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1348_v57, _1336_v46)
				_1349_v58 = (_1349_v58).Update(_1348_v57, (_1335_v45).F18())
			}
		}
		r0 = p1
		r1 = _1283_v2
		return r0, r1
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f12 D1
	_f14 _dafny.Array
	_f15 _dafny.Set
	_f13 bool
	F24  bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f12 = Companion_D1_.Default()
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.EmptySet
	_this._f13 = false
	_this.F24 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C5{}
var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F12() D1 {
	return _this._f12
}
func (_this *C5) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C5) F14() _dafny.Array {
	return _this._f14
}
func (_this *C5) F15() _dafny.Set {
	return _this._f15
}
func (_this *C5) F13() bool {
	return _this._f13
}
func (_this *C5) Ctor__(f24 bool, f14 _dafny.Array, f15 _dafny.Set, f12 D1, f13 bool) {
	{
		(_this).F24 = f24
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C5) Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return (Companion_D8_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-193), _dafny.IntOfInt64(525))).Cardinality()), _dafny.SetOf(false, _this.F24, (Companion_D12_.Create_DC32_(_this.F24, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wbv")).Cardinality()), false, _this.F24, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.IntOfInt64(732))).Cardinality())).Dtor_cf67(), _this.F24, _this.F24), _dafny.MultiSetOf(_dafny.IntOfInt64(144), (_dafny.Zero).Minus((_dafny.MultiSetOf(_dafny.IntOfInt64(-496), _dafny.IntOfInt64(687))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.IntOfInt64(-869)))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(342), _this.F24)).Cardinality(), (_dafny.MultiSetOf((_this).F13(), _this.F24)).Cardinality()), Companion_D1_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(113))).Uint32(), func(coer80 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg80 _dafny.Int) interface{} {
				return coer80(arg80)
			}
		}(func(_1350_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		}))))).Dtor_cf32()
	}
}
func (_this *C5) Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int {
	{
		return (((_this).F15()).Difference((_this).F15())).Cardinality()
	}
}
func (_this *C5) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() D12 {
			if _this.F24 {
				return Companion_D12_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F13()), (_this).F13()))
			}
			return Companion_D12_.Create_DC29_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), _this.F24))
		})(), ((_dafny.MultiSetOf(_dafny.SeqOf(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC16_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(460))).Uint32(), func(coer81 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg81 _dafny.Int) interface{} {
				return coer81(arg81)
			}
		}(func(_1351_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('s')
		}))).Cardinality()))))), Companion_D7_.Create_DC17_(Companion_D7_.Create_DC15_(_dafny.SeqOf(_dafny.IntOfInt64(324))))))).Union(_dafny.MultiSetOf(_dafny.SeqOf(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC16_((_dafny.SetOf(_this.F24, (_this).F13())).Cardinality())), Companion_D7_.Create_DC17_(Companion_D7_.Create_DC15_(_dafny.SeqOf(_dafny.IntOfInt64(-703), _dafny.IntOfInt64(177)))), Companion_D7_.Create_DC17_(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-543))).Uint32(), func(coer82 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg82 _dafny.Int) interface{} {
				return coer82(arg82)
			}
		}(func(_1352_i1 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(387)
		})))))))), _dafny.SeqOf(Companion_D7_.Create_DC17_(Companion_D7_.Create_DC16_((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('x'))).Cardinality()))).Cardinality())))))))).Cardinality())).Cardinality()
	}
}
func (_this *C5) Fm52(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) bool {
	{
		return (_dafny.IntOfInt64(731)).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(459), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-683))).Cardinality()))) <= 0
	}
}
func (_this *C5) Fm53(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_dafny.IntOfInt64(347)).Plus((_dafny.IntOfInt64(111)).Plus(_dafny.IntOfInt64(19))))
	}
}
func (_this *C5) M5(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index210
		((_this).F14()).ArraySet1(false, (_index210).Int())
		var _1353_v0 _dafny.Sequence
		_ = _1353_v0
		_1353_v0 = _dafny.UnicodeSeqOfUtf8Bytes("av")
		var _1354_v1 _dafny.Sequence
		_ = _1354_v1
		_1354_v1 = _dafny.SeqOf(false, (_this).F13())
		var _1355_v2 _dafny.Map
		_ = _1355_v2
		_1355_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24, p0)
		var _1356_v3 _dafny.Array
		_ = _1356_v3
		var _nwElement0_46 _dafny.Int = (_dafny.Zero).Minus(p0)
		_ = _nwElement0_46
		var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(14))
		_ = _nw189
		(_nw189).ArraySet1(_nwElement0_46, 0)
		(_nw189).ArraySet1((_this).Fm8(_this.F24, Companion_D1_.Create_DC1_(_1353_v0), globalState), 1)
		(_nw189).ArraySet1(Companion_Default___.SafeDivisionInt(p0, p0), 2)
		(_nw189).ArraySet1((_dafny.IntOfUint32((_1354_v1).Cardinality())).Times((_1355_v2).Cardinality()), 3)
		(_nw189).ArraySet1((p0).Times(p0), 4)
		(_nw189).ArraySet1(p0, 5)
		(_nw189).ArraySet1(_dafny.IntOfInt64(370), 6)
		(_nw189).ArraySet1(p0, 7)
		(_nw189).ArraySet1(p0, 8)
		(_nw189).ArraySet1(p0, 9)
		(_nw189).ArraySet1(p0, 10)
		(_nw189).ArraySet1(p0, 11)
		(_nw189).ArraySet1(p0, 12)
		(_nw189).ArraySet1(p0, 13)
		_1356_v3 = _nw189
		var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _index211
		(_1356_v3).ArraySet1((func() _dafny.Int {
			if _this.F24 {
				return p0
			}
			return (_1355_v2).Cardinality()
		})(), (_index211).Int())
		var _1357_v4 _dafny.Array
		_ = _1357_v4
		var _nw190 _dafny.Array = _dafny.NewArrayWithValue(Companion_D12_.Default(), _dafny.IntOfInt64(12))
		_ = _nw190
		_1357_v4 = _nw190
		var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1357_v4), 0))
		_ = _index212
		(_1357_v4).ArraySet1(Companion_D12_.Create_DC30_((_this).F14(), (_this).F13()), (_index212).Int())
		var _1358_v5 _dafny.CodePoint
		_ = _1358_v5
		_1358_v5 = _dafny.CodePoint('f')
		var _1359_v6 _dafny.MultiSet
		_ = _1359_v6
		_1359_v6 = _dafny.MultiSetOf(_1358_v5)
		var _1360_v7 _dafny.MultiSet
		_ = _1360_v7
		_1360_v7 = _dafny.MultiSetOf(false, false, _this.F24, (Companion_D12_.Create_DC31_((_this).F13(), _1356_v3, (_this).F13(), _dafny.CodePoint('l'), _1359_v6)).Dtor_cf62())
		var _1361_v8 _dafny.MultiSet
		_ = _1361_v8
		_1361_v8 = _dafny.MultiSetOf(_1356_v3)
		var _1362_v9 _dafny.Sequence
		_ = _1362_v9
		_1362_v9 = _dafny.SeqOf(_1361_v8, _1361_v8)
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index213
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _index214
		var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1357_v4), 0))
		_ = _index215
		var _rhs185 _dafny.Int = p0
		_ = _rhs185
		var _rhs186 bool = !(!(!(false) || ((func() bool {
			if (_this).F13() {
				return _this.F24
			}
			return _this.F24
		})())))
		_ = _rhs186
		var _rhs187 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1354_v1).Cardinality()), ((func() _dafny.MultiSet {
			if _this.F24 {
				return _1360_v7
			}
			return _1360_v7
		})()).Cardinality())
		_ = _rhs187
		var _rhs188 D12 = Companion_D12_.Create_DC30_((func() _dafny.Array {
			if (_this).F13() {
				return (_this).F14()
			}
			return (_this).F14()
		})(), ((_1362_v9).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1362_v9).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_1361_v8))
		_ = _rhs188
		var _lhs151 *GlobalState = globalState
		_ = _lhs151
		var _lhs152 _dafny.Array = (_this).F14()
		_ = _lhs152
		var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(855), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _lhs153
		var _lhs154 _dafny.Array = _1356_v3
		_ = _lhs154
		var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _lhs155
		var _lhs156 _dafny.Array = _1357_v4
		_ = _lhs156
		var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(969), _dafny.ArrayLen((_1357_v4), 0))
		_ = _lhs157
		_lhs151.F7 = _rhs185
		(_lhs152).ArraySet1(_rhs186, (_lhs153).Int())
		(_lhs154).ArraySet1(_rhs187, (_lhs155).Int())
		(_lhs156).ArraySet1(_rhs188, (_lhs157).Int())
		var _1363_v10 _dafny.Sequence
		_ = _1363_v10
		_1363_v10 = _dafny.SeqOf(p0)
		var _1364_v11 D1
		_ = _1364_v11
		_1364_v11 = Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("rkrn"))
		var _1365_v12 _dafny.Array
		_ = _1365_v12
		var _nw191 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
		_ = _nw191
		_1365_v12 = _nw191
		var _pat_let_tv35 = _1358_v5
		_ = _pat_let_tv35
		var _1366_v13 *C2
		_ = _1366_v13
		var _nw192 *C2 = New_C2_()
		_ = _nw192
		_nw192.Ctor__((_dafny.IntOfUint32((_1363_v10).Cardinality())).Plus((_this.F12()).Dtor_cf3()), _this.F24, func(_pat_let37_0 D1) D1 {
			return func(_1367_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let38_0 _dafny.Sequence) D1 {
					return func(_1370_dt__update_hcf1_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC1_(_1370_dt__update_hcf1_h0)
					}(_pat_let38_0)
				}(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(629))).Uint32(), func(coer83 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg83 _dafny.Int) interface{} {
						return coer83(arg83)
					}
				}((func(_1368_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1369_i0 _dafny.Int) _dafny.CodePoint {
						return _1368_v5
					}
				})(_pat_let_tv35))))
			}(_pat_let37_0)
		}(_1364_v11), !(((_1356_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_1353_v0).Cardinality()))) < 0), _1365_v12, (_this).F15(), _this.F12(), (_this).F13())
		_1366_v13 = _nw192
		var _1371_v14 D9
		_ = _1371_v14
		_1371_v14 = Companion_D9_.Create_DC22_((_this).F13(), (_1356_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))).Int()).(_dafny.Int), (_this).F13(), Companion_Default___.Fm54(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(623))).Uint32(), func(coer84 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg84 _dafny.Int) interface{} {
				return coer84(arg84)
			}
		}((func(_1372_v0 _dafny.Sequence, _1373_v13 *C2) func(_dafny.Int) _dafny.Int {
			return func(_1374_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1372_v0, (Companion_Default___.SafeIndex((_1373_v13).F21(), _dafny.IntOfUint32((_1372_v0).Cardinality()))).Uint32(), (_1372_v0).Select((Companion_Default___.SafeIndex((_1373_v13).F21(), _dafny.IntOfUint32((_1372_v0).Cardinality()))).Uint32()).(_dafny.CodePoint))).Cardinality())
			}
		})(_1353_v0, _1366_v13)))).Cardinality()), globalState), _1353_v0)
		r0 = (_1371_v14).Dtor_cf37()
		var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _index216
		(_1356_v3).ArraySet1(p0, (_index216).Int())
		var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _index217
		(_1356_v3).ArraySet1((_1356_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))).Int()).(_dafny.Int), (_index217).Int())
		var _1375_v15 _dafny.Sequence
		_ = _1375_v15
		_1375_v15 = _dafny.SeqOf(_1365_v12, _1365_v12)
		var _1376_v16 _dafny.Set
		_ = _1376_v16
		_1376_v16 = _dafny.SetOf(_1365_v12, (_1375_v15).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(753), _dafny.IntOfUint32((_1375_v15).Cardinality()))).Uint32()).(_dafny.Array), _1365_v12, _1365_v12, _1365_v12)
		var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _index218
		var _rhs189 _dafny.Set = _1376_v16
		_ = _rhs189
		var _rhs190 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1354_v1).Cardinality()), (_1356_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))).Int()).(_dafny.Int))
		_ = _rhs190
		var _rhs191 _dafny.Int = (_1356_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))).Int()).(_dafny.Int)
		_ = _rhs191
		var _lhs158 *GlobalState = globalState
		_ = _lhs158
		var _lhs159 _dafny.Array = _1356_v3
		_ = _lhs159
		var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(773), _dafny.ArrayLen((_1356_v3), 0))
		_ = _lhs160
		_1376_v16 = _rhs189
		_lhs158.F7 = _rhs190
		(_lhs159).ArraySet1(_rhs191, (_lhs160).Int())
		r0 = (_1366_v13).F22()
		return r0
	}
}
func (_this *C5) M14(p0 bool, p1 _dafny.Sequence, p2 _dafny.Sequence, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1377_v0 _dafny.Int
		_ = _1377_v0
		var _out10 _dafny.Int
		_ = _out10
		_out10 = (_this).M15(globalState)
		_1377_v0 = _out10
		var _1378_i0 _dafny.Int
		_ = _1378_i0
		_1378_i0 = _dafny.Zero
		{
			for _this.F24 {
				{
					if (_1378_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1378_i0 = (_1378_i0).Plus(_dafny.One)
					var _1379_v1 _dafny.CodePoint
					_ = _1379_v1
					_1379_v1 = _dafny.CodePoint('e')
					var _1380_v2 D3
					_ = _1380_v2
					_1380_v2 = Companion_D3_.Create_DC5_(_1379_v1)
					var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index219
					((_this).F14()).ArraySet1(_this.F24, (_index219).Int())
					var _1381_v3 _dafny.Map
					_ = _1381_v3
					_1381_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1379_v1, _dafny.IntOfInt64(78))
					var _1382_v4 _dafny.Set
					_ = _1382_v4
					_1382_v4 = _dafny.SetOf(_1381_v3)
					var _1383_v5 _dafny.Array
					_ = _1383_v5
					var _nwElement0_47 _dafny.Int = (Companion_Default___.Fm55((_this).Fm9(_1379_v1, globalState), _1382_v4, p0, globalState)).Cardinality()
					_ = _nwElement0_47
					var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.One)
					_ = _nw193
					(_nw193).ArraySet1(_nwElement0_47, 0)
					_1383_v5 = _nw193
					var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1383_v5), 0))
					_ = _index220
					(_1383_v5).ArraySet1(_1377_v0, (_index220).Int())
					var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index221
					var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1383_v5), 0))
					_ = _index222
					var _rhs192 D3 = _1380_v2
					_ = _rhs192
					var _rhs193 bool = (func() bool {
						if p0 {
							return true
						}
						return (_this).F13()
					})()
					_ = _rhs193
					var _rhs194 _dafny.Int = _1377_v0
					_ = _rhs194
					var _lhs161 _dafny.Array = (_this).F14()
					_ = _lhs161
					var _lhs162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _lhs162
					var _lhs163 _dafny.Array = _1383_v5
					_ = _lhs163
					var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1383_v5), 0))
					_ = _lhs164
					_1380_v2 = _rhs192
					(_lhs161).ArraySet1(_rhs193, (_lhs162).Int())
					(_lhs163).ArraySet1(_rhs194, (_lhs164).Int())
					var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index223
					((_this).F14()).ArraySet1((_this).F13(), (_index223).Int())
					var _1384_v6 D11
					_ = _1384_v6
					_1384_v6 = Companion_D11_.Create_DC26_(_1383_v5, p3)
					_1383_v5 = ((func() D11 {
						if p0 {
							return _1384_v6
						}
						return _1384_v6
					})()).Dtor_cf49()
					r0 = (_1383_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(641), _dafny.ArrayLen((_1383_v5), 0))).Int()).(_dafny.Int)
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1385_v7 _dafny.Set
		_ = _1385_v7
		_1385_v7 = _dafny.SetOf(p2)
		var _1386_v8 _dafny.Set
		_ = _1386_v8
		_1386_v8 = (_1385_v7)
		var _source21 _dafny.Set = _1386_v8
		_ = _source21
		var _1387___mcc_h0 _dafny.Set = _source21
		_ = _1387___mcc_h0
		var _1388_cf16 _dafny.Set = _1387___mcc_h0
		_ = _1388_cf16
		var _1389_v9 _dafny.MultiSet
		_ = _1389_v9
		_1389_v9 = _dafny.MultiSetOf(_1377_v0, _1377_v0, _1377_v0)
		var _1390_v10 _dafny.Array
		_ = _1390_v10
		var _len0_31 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_31
		var _nw194 _dafny.Array
		_ = _nw194
		if _len0_31.Cmp(_dafny.Zero) == 0 {
			_nw194 = _dafny.NewArray(_len0_31)
		} else {
			var _init31 func(_dafny.Int) _dafny.Int = (func(_1391_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1392_i1 _dafny.Int) _dafny.Int {
					return (_1392_i1).Times(_1391_v0)
				}
			})(_1377_v0)
			_ = _init31
			var _element0_31 = _init31(_dafny.Zero)
			_ = _element0_31
			_nw194 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
			(_nw194).ArraySet1(_element0_31, 0)
			var _nativeLen0_31 = (_len0_31).Int()
			_ = _nativeLen0_31
			for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
				(_nw194).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
			}
		}
		_1390_v10 = _nw194
		var _1393_v11 D11
		_ = _1393_v11
		_1393_v11 = Companion_D11_.Create_DC26_(_1390_v10, p3)
		var _1394_v12 _dafny.Sequence
		_ = _1394_v12
		_1394_v12 = _dafny.SeqOf(_1393_v11, Companion_D11_.Create_DC26_(_1390_v10, p3))
		_1389_v9 = (func() _dafny.MultiSet {
			if _this.F24 {
				return _dafny.MultiSetOf(_1377_v0, _1377_v0)
			}
			return Companion_Default___.Fm56(_dafny.IntOfUint32((_1394_v12).Cardinality()), (_this).F15(), globalState)
		})()
		var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index224
		((_this).F14()).ArraySet1((_this).F13(), (_index224).Int())
		var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index225
		((_this).F14()).ArraySet1(p0, (_index225).Int())
		var _1395_v13 _dafny.Set
		_ = _1395_v13
		_1395_v13 = _dafny.SetOf(Companion_D1_.Create_DC1_(p3))
		if (_1395_v13).IsDisjointFrom(_1395_v13) {
			var _1396_v14 D3
			_ = _1396_v14
			_1396_v14 = Companion_D3_.Create_DC6_(_1377_v0, _this.F24)
			var _1397_v15 D7
			_ = _1397_v15
			_1397_v15 = Companion_D7_.Create_DC16_((_1396_v14).Dtor_cf12())
			(globalState).F7 = (_1397_v15).Dtor_cf29()
			var _1398_v16 _dafny.Set
			_ = _1398_v16
			_1398_v16 = _dafny.SetOf(false, _this.F24, p0)
			var _1399_v17 D9
			_ = _1399_v17
			_1399_v17 = Companion_D9_.Create_DC22_((_this).F13(), _1377_v0, false, _1398_v16, p2)
			var _1400_v18 _dafny.Map
			_ = _1400_v18
			_1400_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1377_v0).Cmp(_1377_v0) == 0, ((_this).F13()) && ((_1399_v17).Dtor_cf37()))
			_1400_v18 = (_1400_v18).Merge(_1400_v18)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index226
			((_this).F14()).ArraySet1((Companion_Default___.Fm57(globalState)).Dtor_cf65(), (_index226).Int())
			(_this).F24 = !(!(!(true) || ((func() bool {
				if p0 {
					return ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				}
				return (_this).F13()
			})())))
			(globalState).F7 = Companion_Default___.SafeModuloInt(_1377_v0, _1377_v0)
		} else {
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index227
			((_this).F14()).ArraySet1(false, (_index227).Int())
			var _1401_v19 _dafny.MultiSet
			_ = _1401_v19
			var _out11 _dafny.MultiSet
			_ = _out11
			_out11 = Companion_Default___.M0(_dafny.CodePoint('j'), p0, globalState)
			_1401_v19 = _out11
			var _1402_v20 _dafny.Map
			_ = _1402_v20
			_1402_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1377_v0, _1390_v10)
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))
			_ = _index228
			(_1390_v10).ArraySet1((_1402_v20).Cardinality(), (_index228).Int())
			var _1403_v21 _dafny.Map
			_ = _1403_v21
			_1403_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24, p0)
			var _1404_v22 _dafny.Sequence
			_ = _1404_v22
			_1404_v22 = _dafny.SeqOf((_1403_v21).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_1390_v10, _1390_v10)).Cardinality()))
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))
			_ = _index229
			(_1390_v10).ArraySet1((_dafny.IntOfUint32((_1404_v22).Cardinality())).Plus(_1377_v0), (_index229).Int())
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index230
			((_this).F14()).ArraySet1(!((_this).F13()), (_index230).Int())
			var _1405_v23 _dafny.Array
			_ = _1405_v23
			var _nw195 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(25))
			_ = _nw195
			_1405_v23 = _nw195
			var _1406_v24 _dafny.CodePoint
			_ = _1406_v24
			_1406_v24 = _dafny.CodePoint('o')
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1405_v23), 0))
			_ = _index231
			(_1405_v23).ArraySet1CodePoint(_1406_v24, (_index231).Int())
			var _1407_v25 D12
			_ = _1407_v25
			_1407_v25 = Companion_D12_.Create_DC32_(!(_this.F24) || (p0), (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int), ((_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus(_dafny.IntOfUint32((_1404_v22).Cardinality()))) < 0, (_this).F13(), (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int))
			var _1408_v26 _dafny.Map
			_ = _1408_v26
			_1408_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm58(globalState), (_this).F13())
			var _1409_v27 _dafny.Array
			_ = _1409_v27
			var _nwElement0_48 _dafny.Map = _1408_v26
			_ = _nwElement0_48
			var _nw196 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(6))
			_ = _nw196
			(_nw196).ArraySet1(_nwElement0_48, 0)
			(_nw196).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1406_v24, false), 1)
			(_nw196).ArraySet1((_1408_v26).Merge(_1408_v26), 2)
			(_nw196).ArraySet1((Companion_Default___.Fm59(_this.F24, p0, (_this).F13(), _1406_v24, globalState)).Merge(_1408_v26), 3)
			(_nw196).ArraySet1(_1408_v26, 4)
			(_nw196).ArraySet1(Companion_Default___.Fm59(p0, true, p0, _1406_v24, globalState), 5)
			_1409_v27 = _nw196
			var _1410_v28 _dafny.Sequence
			_ = _1410_v28
			_1410_v28 = _dafny.SeqOf((_this).F13(), (_this).F13(), (_this).F13(), false, (_this).F13())
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1405_v23), 0))
			_ = _index232
			var _rhs195 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_dafny.IntOfUint32((_1410_v28).Cardinality())).Cmp(_1377_v0) < 0 {
					return _1406_v24
				}
				return _1406_v24
			})()
			_ = _rhs195
			var _rhs196 _dafny.Sequence = _dafny.SeqOf((_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int), (_1377_v0).Minus((_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int)), Companion_Default___.SafeDivisionInt(_1377_v0, _1377_v0))
			_ = _rhs196
			var _rhs197 D12 = Companion_Default___.Fm57(globalState)
			_ = _rhs197
			var _rhs198 _dafny.Int = (_1390_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(721), _dafny.ArrayLen((_1390_v10), 0))).Int()).(_dafny.Int)
			_ = _rhs198
			var _rhs199 _dafny.Array = _1409_v27
			_ = _rhs199
			var _lhs165 _dafny.Array = _1405_v23
			_ = _lhs165
			var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(192), _dafny.ArrayLen((_1405_v23), 0))
			_ = _lhs166
			(_lhs165).ArraySet1CodePoint(_rhs195, (_lhs166).Int())
			_1404_v22 = _rhs196
			_1407_v25 = _rhs197
			r0 = _rhs198
			_1409_v27 = _rhs199
		}
		var _1411_v29 _dafny.Map
		_ = _1411_v29
		_1411_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1390_v10, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
		var _1412_v30 D14
		_ = _1412_v30
		_1412_v30 = Companion_D14_.Create_DC36_((_1411_v29).Update(_1390_v10, (_this).F13()))
		_1411_v29 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1390_v10, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(761), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))).Merge((_1412_v30).Dtor_cf74())
		(_this).F24 = (!(p0)) || (p0)
		var _1413_v31 _dafny.Array
		_ = _1413_v31
		var _len0_32 _dafny.Int = _dafny.One
		_ = _len0_32
		var _nw197 _dafny.Array
		_ = _nw197
		if _len0_32.Cmp(_dafny.Zero) == 0 {
			_nw197 = _dafny.NewArray(_len0_32)
		} else {
			var _init32 func(_dafny.Int) _dafny.CodePoint = func(_1414_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('o')
			}
			_ = _init32
			var _element0_32 = _init32(_dafny.Zero)
			_ = _element0_32
			_nw197 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
			(_nw197).ArraySet1CodePoint(_element0_32, 0)
			var _nativeLen0_32 = (_len0_32).Int()
			_ = _nativeLen0_32
			for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
				(_nw197).ArraySet1CodePoint(_init32(_dafny.IntOf(_i0_32)), _i0_32)
			}
		}
		_1413_v31 = _nw197
		var _1415_v32 _dafny.CodePoint
		_ = _1415_v32
		_1415_v32 = _dafny.CodePoint('b')
		var _1416_v33 _dafny.Sequence
		_ = _1416_v33
		_1416_v33 = _dafny.SeqOf(_1377_v0, _1377_v0, _1377_v0, _1377_v0, _1377_v0)
		var _1417_v34 D11
		_ = _1417_v34
		_1417_v34 = Companion_D11_.Create_DC27_(p0, _1415_v32, _1377_v0, _dafny.IntOfUint32((p1).Cardinality()), _1416_v33)
		var _nwElement0_49 _dafny.CodePoint = _1415_v32
		_ = _nwElement0_49
		var _nw198 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(28))
		_ = _nw198
		(_nw198).ArraySet1CodePoint(_nwElement0_49, 0)
		(_nw198).ArraySet1CodePoint(_1415_v32, 1)
		(_nw198).ArraySet1CodePoint(_1415_v32, 2)
		(_nw198).ArraySet1CodePoint(_1415_v32, 3)
		(_nw198).ArraySet1CodePoint(_1415_v32, 4)
		(_nw198).ArraySet1CodePoint(_1415_v32, 5)
		(_nw198).ArraySet1CodePoint((p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint), 6)
		(_nw198).ArraySet1CodePoint(_1415_v32, 7)
		(_nw198).ArraySet1CodePoint(_1415_v32, 8)
		(_nw198).ArraySet1CodePoint(_1415_v32, 9)
		(_nw198).ArraySet1CodePoint(_1415_v32, 10)
		(_nw198).ArraySet1CodePoint(_1415_v32, 11)
		(_nw198).ArraySet1CodePoint(_1415_v32, 12)
		(_nw198).ArraySet1CodePoint(_1415_v32, 13)
		(_nw198).ArraySet1CodePoint((_1417_v34).Dtor_cf52(), 14)
		(_nw198).ArraySet1CodePoint(_1415_v32, 15)
		(_nw198).ArraySet1CodePoint(_1415_v32, 16)
		(_nw198).ArraySet1CodePoint(_1415_v32, 17)
		(_nw198).ArraySet1CodePoint(_1415_v32, 18)
		(_nw198).ArraySet1CodePoint(_1415_v32, 19)
		(_nw198).ArraySet1CodePoint(_1415_v32, 20)
		(_nw198).ArraySet1CodePoint(_1415_v32, 21)
		(_nw198).ArraySet1CodePoint(_dafny.CodePoint('i'), 22)
		(_nw198).ArraySet1CodePoint(_1415_v32, 23)
		(_nw198).ArraySet1CodePoint(_1415_v32, 24)
		(_nw198).ArraySet1CodePoint(_1415_v32, 25)
		(_nw198).ArraySet1CodePoint(_1415_v32, 26)
		(_nw198).ArraySet1CodePoint(_1415_v32, 27)
		_1413_v31 = _nw198
		var _1418_v35 _dafny.Sequence
		_ = _1418_v35
		_1418_v35 = _dafny.SeqOf(p0, p0)
		(globalState).F2 = _1418_v35
		r0 = _1377_v0
		return r0
	}
}
func (_this *C5) M15(globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1419_v0 _dafny.MultiSet
		_ = _1419_v0
		_1419_v0 = _dafny.MultiSetOf((_this).F13())
		var _1420_v1 _dafny.Sequence
		_ = _1420_v1
		_1420_v1 = _dafny.SeqOf(_1419_v0)
		var _1421_v2 _dafny.Int
		_ = _1421_v2
		_1421_v2 = _dafny.IntOfInt64(-765)
		var _1422_v3 _dafny.MultiSet
		_ = _1422_v3
		_1422_v3 = (_1420_v1).Select((Companion_Default___.SafeIndex(_1421_v2, _dafny.IntOfUint32((_1420_v1).Cardinality()))).Uint32()).(_dafny.MultiSet)
		var _source22 _dafny.MultiSet = _1422_v3
		_ = _source22
		var _1423___mcc_h0 _dafny.MultiSet = _source22
		_ = _1423___mcc_h0
		var _1424_cf0 _dafny.MultiSet = _1423___mcc_h0
		_ = _1424_cf0
		var _1425_v4 _dafny.CodePoint
		_ = _1425_v4
		_1425_v4 = _dafny.CodePoint('u')
		var _1426_v5 D3
		_ = _1426_v5
		_1426_v5 = Companion_D3_.Create_DC5_(_1425_v4)
		var _source23 D3 = _1426_v5
		_ = _source23
		if _source23.Is_DC6() {
			var _1427___mcc_h1 _dafny.Int = _source23.Get_().(D3_DC6).Cf12
			_ = _1427___mcc_h1
			var _1428___mcc_h2 bool = _source23.Get_().(D3_DC6).Cf13
			_ = _1428___mcc_h2
			var _1429_cf13 bool = _1428___mcc_h2
			_ = _1429_cf13
			var _1430_cf12 _dafny.Int = _1427___mcc_h1
			_ = _1430_cf12
			var _1431_v7 _dafny.Sequence
			_ = _1431_v7
			_1431_v7 = _dafny.UnicodeSeqOfUtf8Bytes("cdytr")
			var _1432_v8 _dafny.Map
			_ = _1432_v8
			_1432_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1431_v7, _1425_v4)
			_1421_v2 = (_dafny.Zero).Minus((func() _dafny.Map {
				var _coll45 = _dafny.NewMapBuilder()
				_ = _coll45
				for _iter53 := _dafny.Iterate(((_1432_v8).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1431_v7, _1425_v4))).Keys().Elements()); ; {
					_compr_45, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _1433_v6 _dafny.Sequence
					_1433_v6 = interface{}(_compr_45).(_dafny.Sequence)
					if ((_1432_v8).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1431_v7, _1425_v4))).Contains(_1433_v6) {
						_coll45.Add(_1433_v6, (_dafny.IntOfUint32((_dafny.SeqOf(_this.F24, (_this).F13())).Cardinality())).Times(_1421_v2))
					}
				}
				return _coll45.ToMap()
			}()).Cardinality())
			var _1434_v9 _dafny.Array
			_ = _1434_v9
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_33
			var _nw199 _dafny.Array
			_ = _nw199
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw199 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) bool = func(_1435_i0 _dafny.Int) bool {
					return (_this).F13()
				}
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw199 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw199).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw199).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1434_v9 = _nw199
			_1434_v9 = (func() _dafny.Array {
				if _this.F24 {
					return _1434_v9
				}
				return (_this).F14()
			})()
			var _1436_v10 _dafny.Sequence
			_ = _1436_v10
			_1436_v10 = _dafny.SeqOf(_1429_cf13)
			var _1437_v11 _dafny.Set
			_ = _1437_v11
			_1437_v11 = _dafny.SetOf(_1436_v10)
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index233
			((_this).F14()).ArraySet1((_this).Fm52((_dafny.Zero).Minus((_this).Fm53(_1430_cf12, globalState)), _1437_v11, globalState), (_index233).Int())
			var _1438_v12 _dafny.Map
			_ = _1438_v12
			_1438_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F13())
			var _1439_v13 _dafny.MultiSet
			_ = _1439_v13
			_1439_v13 = _dafny.MultiSetOf(_1431_v7)
			var _1440_v14 _dafny.Map
			_ = _1440_v14
			_1440_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1430_cf12, _1429_cf13)
			var _1441_v16 _dafny.Set
			_ = _1441_v16
			_1441_v16 = _dafny.SetOf(func() _dafny.Set {
				var _coll46 = _dafny.NewBuilder()
				_ = _coll46
				for _iter54 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(996), _dafny.IntOfInt64(-82))); ; {
					_compr_46, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1442_v15 _dafny.Int
					_1442_v15 = interface{}(_compr_46).(_dafny.Int)
					if ((_dafny.IntOfInt64(996)).Cmp(_1442_v15) <= 0) && ((_1442_v15).Cmp(_dafny.IntOfInt64(-82)) < 0) {
						_coll46.Add((_1442_v15).Times(_1421_v2))
					}
				}
				return _coll46.ToSet()
			}())
			var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index234
			var _rhs200 bool = !((func() bool {
				if (_1438_v12).Contains((_this).Fm52(_1421_v2, _1437_v11, globalState)) {
					return (_1438_v12).Get((_this).Fm52(_1421_v2, _1437_v11, globalState)).(bool)
				}
				return (_this).F13()
			})())
			_ = _rhs200
			var _rhs201 bool = _dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kv"), _dafny.UnicodeSeqOfUtf8Bytes("nolgwiuyb")), _1425_v4)
			_ = _rhs201
			var _rhs202 _dafny.Int = (func() _dafny.Int {
				if (_1439_v13).Contains(_dafny.Companion_Sequence_.Concatenate(_1431_v7, Companion_Default___.Fm60(_this.F24, (_this).F13(), _1440_v14, (_1441_v16).Cardinality(), globalState))) {
					return (_1439_v13).Multiplicity(_dafny.Companion_Sequence_.Concatenate(_1431_v7, Companion_Default___.Fm60(_this.F24, (_this).F13(), _1440_v14, (_1441_v16).Cardinality(), globalState)))
				}
				return _1421_v2
			})()
			_ = _rhs202
			var _lhs167 _dafny.Array = (_this).F14()
			_ = _lhs167
			var _lhs168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(319), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs168
			var _lhs169 *GlobalState = globalState
			_ = _lhs169
			_1429_cf13 = _rhs200
			(_lhs167).ArraySet1(_rhs201, (_lhs168).Int())
			_lhs169.F7 = _rhs202
			var _1443_v17 *C4
			_ = _1443_v17
			var _nw200 *C4 = New_C4_()
			_ = _nw200
			_nw200.Ctor__()
			_1443_v17 = _nw200
			var _1444_v18 _dafny.Map
			_ = _1444_v18
			_1444_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1443_v17, _1421_v2)
			var _1445_v19 _dafny.Sequence
			_ = _1445_v19
			_1445_v19 = _dafny.SeqOf((_1444_v18).Cardinality())
			var _1446_v20 _dafny.Map
			_ = _1446_v20
			_1446_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24, (_1445_v19).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(988), _dafny.IntOfUint32((_1445_v19).Cardinality()))).Uint32()).(_dafny.Int))
			_1446_v20 = _1446_v20
		} else if _source23.Is_DC7() {
			var _1447_v21 _dafny.Map
			_ = _1447_v21
			_1447_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(937), (_dafny.Zero).Minus(_dafny.IntOfInt64(-550)))
			var _1448_v22 _dafny.Set
			_ = _1448_v22
			_1448_v22 = _dafny.SetOf(Companion_Default___.Fm61(_this.F24, (_1447_v21).Cardinality(), _dafny.MultiSetOf(_1425_v4, _1425_v4), globalState))
			var _1449_v23 _dafny.Map
			_ = _1449_v23
			_1449_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm52((_dafny.Zero).Minus(_1421_v2), _1448_v22, globalState), (_this).F14())
			var _1450_v24 *C2
			_ = _1450_v24
			var _nw201 *C2 = New_C2_()
			_ = _nw201
			_nw201.Ctor__(_1421_v2, (func() bool {
				if true {
					return _this.F24
				}
				return _this.F24
			})(), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("uy")), (_this).F13(), (func() _dafny.Array {
				if (_1449_v23).Contains((_this).F13()) {
					return (_1449_v23).Get((_this).F13()).(_dafny.Array)
				}
				return (_this).F14()
			})(), (_this).F15(), _this.F12(), _this.F24)
			_1450_v24 = _nw201
			var _1451_v25 _dafny.Array
			_ = _1451_v25
			var _nw202 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
			_ = _nw202
			_1451_v25 = _nw202
			var _1452_v26 D8
			_ = _1452_v26
			_1452_v26 = Companion_D8_.Create_DC18_(_1451_v25)
			_1452_v26 = _1452_v26
			var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index235
			((_this).F14()).ArraySet1((_this).F13(), (_index235).Int())
			var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index236
			var _rhs203 bool = !(_this.F24)
			_ = _rhs203
			var _rhs204 bool = _this.F24
			_ = _rhs204
			var _rhs205 bool = false
			_ = _rhs205
			var _lhs170 _dafny.Array = (_this).F14()
			_ = _lhs170
			var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs171
			var _lhs172 *C5 = _this
			_ = _lhs172
			var _lhs173 *C5 = _this
			_ = _lhs173
			(_lhs170).ArraySet1(_rhs203, (_lhs171).Int())
			_lhs172.F24 = _rhs204
			_lhs173.F24 = _rhs205
			var _1453_v27 _dafny.MultiSet
			_ = _1453_v27
			_1453_v27 = _dafny.MultiSetOf(_1421_v2, _1421_v2)
			(globalState).F7 = ((func() _dafny.Int {
				if (_1453_v27).Contains((_1450_v24).F21()) {
					return (_1453_v27).Multiplicity((_1450_v24).F21())
				}
				return _1421_v2
			})()).Times(_1421_v2)
		} else if _source23.Is_DC8() {
			var _1454___mcc_h3 _dafny.Int = _source23.Get_().(D3_DC8).Cf14
			_ = _1454___mcc_h3
			var _1455_cf14 _dafny.Int = _1454___mcc_h3
			_ = _1455_cf14
			var _1456_v28 _dafny.Sequence
			_ = _1456_v28
			_1456_v28 = _dafny.SeqOf(_this.F24, _this.F24)
			var _1457_v29 D15
			_ = _1457_v29
			_1457_v29 = Companion_D15_.Create_DC38_(_1456_v28)
			var _1458_v30 _dafny.Set
			_ = _1458_v30
			_1458_v30 = _dafny.SetOf((_1457_v29).Dtor_cf77(), _dafny.SeqOf(_this.F24))
			var _1459_v31 _dafny.Sequence
			_ = _1459_v31
			_1459_v31 = _dafny.SeqOf(_1458_v30, _dafny.SetOf(_1456_v28, _1456_v28, _1456_v28))
			var _1460_v32 _dafny.Set
			_ = _1460_v32
			_1460_v32 = _dafny.SetOf(_this.F24)
			var _1461_v33 _dafny.Sequence
			_ = _1461_v33
			_1461_v33 = _dafny.SeqOf(_1460_v32)
			(_this).F24 = !((_this).Fm52(_1421_v2, (_1459_v31).Select((Companion_Default___.SafeIndex(((_1461_v33).Select((Companion_Default___.SafeIndex(_1421_v2, _dafny.IntOfUint32((_1461_v33).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _dafny.IntOfUint32((_1459_v31).Cardinality()))).Uint32()).(_dafny.Set), globalState))
			var _1462_v34 _dafny.Sequence
			_ = _1462_v34
			_1462_v34 = _dafny.SeqOf(_1421_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, (_this).F13()), _1456_v28)).Cardinality()), _1455_cf14)
			(globalState).F7 = (_1462_v34).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1455_cf14), _dafny.IntOfUint32((_1462_v34).Cardinality()))).Uint32()).(_dafny.Int)
			_1421_v2 = _1455_cf14
			var _rhs206 bool = (((_this).F15()).IsProperSubsetOf((_this).F15())) || ((_this).F13())
			_ = _rhs206
			var _rhs207 _dafny.Int = Companion_Default___.SafeModuloInt(_1421_v2, _1455_cf14)
			_ = _rhs207
			var _rhs208 _dafny.Sequence = _1456_v28
			_ = _rhs208
			var _lhs174 *C5 = _this
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			var _lhs176 *GlobalState = globalState
			_ = _lhs176
			_lhs174.F24 = _rhs206
			_lhs175.F7 = _rhs207
			_lhs176.F2 = _rhs208
		} else {
			var _1463___mcc_h4 _dafny.CodePoint = _source23.Get_().(D3_DC5).Cf11
			_ = _1463___mcc_h4
			var _1464_cf11 _dafny.CodePoint = _1463___mcc_h4
			_ = _1464_cf11
			var _1465_v35 _dafny.Sequence
			_ = _1465_v35
			_1465_v35 = _dafny.UnicodeSeqOfUtf8Bytes("dffjrx")
			var _1466_v36 _dafny.Sequence
			_ = _1466_v36
			_1466_v36 = _dafny.SeqOf(_1421_v2)
			var _1467_v37 _dafny.Map
			_ = _1467_v37
			_1467_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_1466_v36)).Cardinality(), _1425_v4)
			var _1468_v38 _dafny.Map
			_ = _1468_v38
			_1468_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1467_v37).Cardinality(), (_this).F13())
			var _1469_v39 _dafny.Map
			_ = _1469_v39
			_1469_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24, _1421_v2)
			var _1470_v40 D6
			_ = _1470_v40
			_1470_v40 = Companion_D6_.Create_DC13_(_this.F24, _this.F24, _1421_v2, _dafny.IntOfUint32((_dafny.SeqOf((_this).F13())).Cardinality()), !((_this).F13()))
			var _1471_v41 T0
			_ = _1471_v41
			var _nw203 *C3 = New_C3_()
			_ = _nw203
			_nw203.Ctor__((_this).F13(), Companion_Default___.Fm62(_this.F24, _1465_v35, globalState), (_this).F13(), (_this).F14(), (_this).F15(), _this.F12(), _this.F24)
			_1471_v41 = _nw203
			var _1472_v42 _dafny.Map
			_ = _1472_v42
			_1472_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1471_v41, (_this).F13())
			var _1473_v43 _dafny.MultiSet
			_ = _1473_v43
			_1473_v43 = _dafny.MultiSetOf(_1421_v2)
			var _1474_v44 _dafny.Set
			_ = _1474_v44
			_1474_v44 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_1471_v41).F13(), true, (_this).F13(), (_this).F13()), (Companion_Default___.SafeIndex((_1473_v43).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_1471_v41).F13(), true, (_this).F13(), (_this).F13())).Cardinality()))).Uint32(), (_this).F13()))
			var _1475_v45 _dafny.Map
			_ = _1475_v45
			_1475_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).Fm52(_1421_v2, _1474_v44, globalState))
			var _1476_v46 _dafny.Array
			_ = _1476_v46
			var _nwElement0_50 _dafny.Int = (_dafny.IntOfUint32((_1465_v35).Cardinality())).Plus(_1421_v2)
			_ = _nwElement0_50
			var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(28))
			_ = _nw204
			(_nw204).ArraySet1(_nwElement0_50, 0)
			(_nw204).ArraySet1(((_1468_v38).Merge(_1468_v38)).Cardinality(), 1)
			(_nw204).ArraySet1(_1421_v2, 2)
			(_nw204).ArraySet1((func() _dafny.Int {
				if (_1469_v39).Contains(_this.F24) {
					return (_1469_v39).Get(_this.F24).(_dafny.Int)
				}
				return _1421_v2
			})(), 3)
			(_nw204).ArraySet1(_1421_v2, 4)
			(_nw204).ArraySet1(_1421_v2, 5)
			(_nw204).ArraySet1((_dafny.IntOfInt64(428)).Plus((_this).Fm8((_this).F13(), Companion_D1_.Create_DC1_(_1465_v35), globalState)), 6)
			(_nw204).ArraySet1(_1421_v2, 7)
			(_nw204).ArraySet1(_1421_v2, 8)
			(_nw204).ArraySet1(_1421_v2, 9)
			(_nw204).ArraySet1(((_dafny.Zero).Minus((_1470_v40).Dtor_cf25())).Plus(_1421_v2), 10)
			(_nw204).ArraySet1(_1421_v2, 11)
			(_nw204).ArraySet1((_dafny.Zero).Minus((_1421_v2).Plus(_1421_v2)), 12)
			(_nw204).ArraySet1(_1421_v2, 13)
			(_nw204).ArraySet1(_1421_v2, 14)
			(_nw204).ArraySet1(_1421_v2, 15)
			(_nw204).ArraySet1(_1421_v2, 16)
			(_nw204).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F24, _1421_v2)).Cardinality(), 17)
			(_nw204).ArraySet1(_1421_v2, 18)
			(_nw204).ArraySet1(_1421_v2, 19)
			(_nw204).ArraySet1((func() _dafny.Int {
				if (_this).F13() {
					return _1421_v2
				}
				return _1421_v2
			})(), 20)
			(_nw204).ArraySet1((_1472_v42).Cardinality(), 21)
			(_nw204).ArraySet1(((_1475_v45).Merge(_1475_v45)).Cardinality(), 22)
			(_nw204).ArraySet1(_1421_v2, 23)
			(_nw204).ArraySet1(_1421_v2, 24)
			(_nw204).ArraySet1((_dafny.IntOfInt64(105)).Minus(_1421_v2), 25)
			(_nw204).ArraySet1((_dafny.Zero).Minus(_1421_v2), 26)
			(_nw204).ArraySet1(_1421_v2, 27)
			_1476_v46 = _nw204
			var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1476_v46), 0))
			_ = _index237
			(_1476_v46).ArraySet1(_1421_v2, (_index237).Int())
			var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1476_v46), 0))
			_ = _index238
			(_1476_v46).ArraySet1(((_1424_cf0).Cardinality()).Plus(_1421_v2), (_index238).Int())
			_1473_v43 = _1473_v43
			var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1476_v46), 0))
			_ = _index239
			(_1476_v46).ArraySet1((_dafny.Zero).Minus((_1476_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1476_v46), 0))).Int()).(_dafny.Int)), (_index239).Int())
			var _1477_v47 _dafny.Map
			_ = _1477_v47
			_1477_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1476_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(642), _dafny.ArrayLen((_1476_v46), 0))).Int()).(_dafny.Int), (_dafny.Zero).Minus(_1421_v2))
			var _1478_v48 _dafny.Map
			_ = _1478_v48
			_1478_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1464_cf11, (_1466_v36).Select((Companion_Default___.SafeIndex((_1477_v47).Cardinality(), _dafny.IntOfUint32((_1466_v36).Cardinality()))).Uint32()).(_dafny.Int))
			_1478_v48 = (_1478_v48).Update(_1425_v4, _1421_v2)
		}
		var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index240
		((_this).F14()).ArraySet1(_this.F24, (_index240).Int())
		var _1479_v49 _dafny.Map
		_ = _1479_v49
		_1479_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1421_v2)
		var _1480_v50 D6
		_ = _1480_v50
		_1480_v50 = Companion_D6_.Create_DC13_((_this).F13(), (_this).F13(), _1421_v2, (_1479_v49).Cardinality(), !((_this).F13()))
		var _1481_v51 _dafny.Map
		_ = _1481_v51
		_1481_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F24)
		var _1482_v52 _dafny.Sequence
		_ = _1482_v52
		_1482_v52 = _dafny.UnicodeSeqOfUtf8Bytes("hmg")
		var _pat_let_tv36 = _1482_v52
		_ = _pat_let_tv36
		var _pat_let_tv37 = _1481_v51
		_ = _pat_let_tv37
		var _pat_let_tv38 = _1481_v51
		_ = _pat_let_tv38
		var _1483_v53 _dafny.Map
		_ = _1483_v53
		_1483_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1421_v2, (func(_pat_let39_0 D6) D6 {
			return func(_1484_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let40_0 _dafny.Int) D6 {
					return func(_1485_dt__update_hcf24_h0 _dafny.Int) D6 {
						return func(_pat_let41_0 bool) D6 {
							return func(_1486_dt__update_hcf22_h0 bool) D6 {
								return Companion_D6_.Create_DC13_(_1486_dt__update_hcf22_h0, (_1484_dt__update__tmp_h0).Dtor_cf23(), _1485_dt__update_hcf24_h0, (_1484_dt__update__tmp_h0).Dtor_cf25(), (_1484_dt__update__tmp_h0).Dtor_cf26())
							}(_pat_let41_0)
						}((func() bool {
							if (_pat_let_tv37).Contains(_this.F24) {
								return (_pat_let_tv38).Get(_this.F24).(bool)
							}
							return !(!((_this).F13()))
						})())
					}(_pat_let40_0)
				}(_dafny.IntOfUint32((_pat_let_tv36).Cardinality()))
			}(_pat_let39_0)
		}(_1480_v50)).Dtor_cf26())
		var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index241
		((_this).F14()).ArraySet1((func() bool {
			if (_1483_v53).Contains((_1421_v2).Minus(_1421_v2)) {
				return (_1483_v53).Get((_1421_v2).Minus(_1421_v2)).(bool)
			}
			return !(false)
		})(), (_index241).Int())
		(_this).F24 = (func() bool {
			if (_this).F13() {
				return !((_this).F13())
			}
			return _this.F24
		})()
		var _1487_v54 _dafny.Sequence
		_ = _1487_v54
		_1487_v54 = _dafny.SeqOf(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
		var _1488_v55 _dafny.MultiSet
		_ = _1488_v55
		_1488_v55 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1482_v52, (Companion_Default___.SafeIndex(_1421_v2, _dafny.IntOfUint32((_1482_v52).Cardinality()))).Uint32(), _1425_v4)).Cardinality()), _1421_v2)
		if (_1488_v55).IsProperSubsetOf((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(424))).Uint32(), func(coer85 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg85 _dafny.Int) interface{} {
				return coer85(arg85)
			}
		}((func(_1489_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1490_i1 _dafny.Int) _dafny.Int {
				return _1489_v2
			}
		})(_1421_v2))))).Intersection(Companion_Default___.Fm56(_dafny.IntOfUint32((_1487_v54).Cardinality()), (_this).F15(), globalState))) {
			var _1491_v56 _dafny.Sequence
			_ = _1491_v56
			_1491_v56 = _dafny.SeqOf(_1421_v2, _dafny.IntOfInt64(703), _1421_v2, _1421_v2, _1421_v2)
			_1491_v56 = _1491_v56
			var _1492_v57 _dafny.Array
			_ = _1492_v57
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_34
			var _nw205 _dafny.Array
			_ = _nw205
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw205 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Int = (func(_1493_v2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1494_i2 _dafny.Int) _dafny.Int {
						return (_1494_i2).Plus(_1493_v2)
					}
				})(_1421_v2)
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw205 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw205).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw205).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1492_v57 = _nw205
			var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1492_v57), 0))
			_ = _index242
			(_1492_v57).ArraySet1((_1421_v2).Minus((_1479_v49).Cardinality()), (_index242).Int())
			var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1492_v57), 0))
			_ = _index243
			(_1492_v57).ArraySet1((_1421_v2).Times(_1421_v2), (_index243).Int())
			(globalState).F7 = ((_this).F15()).Cardinality()
			_1482_v52 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(483))).Uint32(), func(coer86 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg86 _dafny.Int) interface{} {
					return coer86(arg86)
				}
			}(func(_1495_i3 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))
			var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1492_v57), 0))
			_ = _index244
			(_1492_v57).ArraySet1((_1492_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(548), _dafny.ArrayLen((_1492_v57), 0))).Int()).(_dafny.Int), (_index244).Int())
		} else {
			var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index245
			((_this).F14()).ArraySet1(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_index245).Int())
			var _1496_v58 _dafny.Sequence
			_ = _1496_v58
			_1496_v58 = _dafny.SeqOf(_1421_v2)
			var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(234), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index246
			((_this).F14()).ArraySet1(!((_dafny.IntOfUint32((_1496_v58).Cardinality())).Cmp(_1421_v2) <= 0), (_index246).Int())
			r0 = _1421_v2
			(globalState).F7 = _1421_v2
			var _1497_v59 D7
			_ = _1497_v59
			_1497_v59 = Companion_D7_.Create_DC16_(_1421_v2)
			_1497_v59 = _1497_v59
		}
		var _1498_v60 _dafny.Sequence
		_ = _1498_v60
		_1498_v60 = _dafny.UnicodeSeqOfUtf8Bytes("jnxuox")
		var _1499_v61 _dafny.Array
		_ = _1499_v61
		var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
		_ = _nw206
		_1499_v61 = _nw206
		var _1500_v62 _dafny.Map
		_ = _1500_v62
		_1500_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_1421_v2, _dafny.IntOfUint32((_1498_v60).Cardinality())), _1499_v61)
		var _1501_v63 _dafny.Map
		_ = _1501_v63
		_1501_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _1421_v2)
		var _1502_v64 D1
		_ = _1502_v64
		_1502_v64 = Companion_D1_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(734))).Uint32(), func(coer87 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg87 _dafny.Int) interface{} {
				return coer87(arg87)
			}
		}(func(_1503_i4 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('d')
		})))
		var _1504_v65 *C2
		_ = _1504_v65
		var _nw207 *C2 = New_C2_()
		_ = _nw207
		_nw207.Ctor__((_1501_v63).Cardinality(), _this.F24, _1502_v64, !(true), (_this).F14(), (_this).F15(), _this.F12(), _this.F24)
		_1504_v65 = _nw207
		var _1505_v66 _dafny.Map
		_ = _1505_v66
		_1505_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v65, _dafny.SeqOf((_this).F13()))
		var _1506_v67 _dafny.Sequence
		_ = _1506_v67
		_1506_v67 = _dafny.SeqOf(_this.F24)
		var _1507_v68 _dafny.Set
		_ = _1507_v68
		_1507_v68 = _dafny.SetOf((func() _dafny.Sequence {
			if (_1505_v66).Contains(_1504_v65) {
				return (_1505_v66).Get(_1504_v65).(_dafny.Sequence)
			}
			return _1506_v67
		})())
		var _1508_v69 D6
		_ = _1508_v69
		_1508_v69 = Companion_D6_.Create_DC14_(Companion_D6_.Create_DC13_((_this).F13(), (_this).Fm52(_1421_v2, _1507_v68, globalState), (_dafny.MultiSetFromSeq(_1506_v67)).Cardinality(), _1421_v2, (_1504_v65).F22()))
		var _1509_v70 _dafny.Map
		_ = _1509_v70
		_1509_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1421_v2, _1508_v69)
		_1500_v62 = (_1500_v62).Update((_1509_v70).Cardinality(), _1499_v61)
		var _hi9 _dafny.Int = _1421_v2
		_ = _hi9
		for _1510_i5 := _1421_v2; _1510_i5.Cmp(_hi9) < 0; _1510_i5 = _1510_i5.Plus(_dafny.One) {
			var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1499_v61), 0))
			_ = _index247
			(_1499_v61).ArraySet1(_1510_i5, (_index247).Int())
			var _1511_v71 _dafny.CodePoint
			_ = _1511_v71
			_1511_v71 = _dafny.CodePoint('t')
			var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1499_v61), 0))
			_ = _index248
			var _rhs209 bool = !(_dafny.Companion_Sequence_.IsProperPrefixOf(_1498_v60, _dafny.Companion_Sequence_.Update(_1498_v60, (Companion_Default___.SafeIndex((_1504_v65).F21(), _dafny.IntOfUint32((_1498_v60).Cardinality()))).Uint32(), _1511_v71))) || ((_this).Fm52(_1421_v2, _1507_v68, globalState))
			_ = _rhs209
			var _rhs210 _dafny.Array = _1499_v61
			_ = _rhs210
			var _rhs211 _dafny.Sequence = _1498_v60
			_ = _rhs211
			var _rhs212 bool = (_1504_v65).F22()
			_ = _rhs212
			var _rhs213 _dafny.Int = (_1504_v65).F21()
			_ = _rhs213
			var _lhs177 *C5 = _this
			_ = _lhs177
			var _lhs178 *C5 = _this
			_ = _lhs178
			var _lhs179 _dafny.Array = _1499_v61
			_ = _lhs179
			var _lhs180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_1499_v61), 0))
			_ = _lhs180
			_lhs177.F24 = _rhs209
			_1499_v61 = _rhs210
			_1498_v60 = _rhs211
			_lhs178.F24 = _rhs212
			(_lhs179).ArraySet1(_rhs213, (_lhs180).Int())
			var _1512_v72 _dafny.Array
			_ = _1512_v72
			var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
			_ = _nw208
			_1512_v72 = _nw208
			var _1513_v73 D11
			_ = _1513_v73
			_1513_v73 = Companion_D11_.Create_DC26_(_1512_v72, _1498_v60)
			var _1514_v74 _dafny.Array
			_ = _1514_v74
			var _nwElement0_51 _dafny.Array = _1512_v72
			_ = _nwElement0_51
			var _nw209 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(8))
			_ = _nw209
			(_nw209).ArraySet1(_nwElement0_51, 0)
			(_nw209).ArraySet1(_1512_v72, 1)
			(_nw209).ArraySet1(_1512_v72, 2)
			(_nw209).ArraySet1(_1512_v72, 3)
			(_nw209).ArraySet1(_1512_v72, 4)
			(_nw209).ArraySet1((_1513_v73).Dtor_cf49(), 5)
			(_nw209).ArraySet1(_1512_v72, 6)
			(_nw209).ArraySet1(_1512_v72, 7)
			_1514_v74 = _nw209
			var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1514_v74), 0))
			_ = _index249
			(_1514_v74).ArraySet1(_1512_v72, (_index249).Int())
			var _1515_v75 _dafny.Array
			_ = _1515_v75
			var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(9))
			_ = _nw210
			_1515_v75 = _nw210
			var _1516_v76 _dafny.Map
			_ = _1516_v76
			_1516_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1504_v65).F21(), _dafny.IntOfInt64(501))
			var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1514_v74), 0))
			_ = _index250
			var _rhs214 bool = _this.F24
			_ = _rhs214
			var _rhs215 _dafny.Array = _1512_v72
			_ = _rhs215
			var _rhs216 _dafny.Array = (func() _dafny.Array {
				if ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), true)).Cardinality()).Cmp((_dafny.Zero).Minus((_1516_v76).Cardinality())) > 0 {
					return _1515_v75
				}
				return _1515_v75
			})()
			_ = _rhs216
			var _lhs181 *C5 = _this
			_ = _lhs181
			var _lhs182 _dafny.Array = _1514_v74
			_ = _lhs182
			var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1514_v74), 0))
			_ = _lhs183
			_lhs181.F24 = _rhs214
			(_lhs182).ArraySet1(_rhs215, (_lhs183).Int())
			_1515_v75 = _rhs216
			var _1517_v77 D2
			_ = _1517_v77
			_1517_v77 = Companion_D2_.Create_DC4_(_dafny.ArrayCastTo((_1514_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(938), _dafny.ArrayLen((_1514_v74), 0))).Int())), (_this).F13(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("iqftrwd")).Cardinality()), true, (_this).F14())
			var _1518_v78 _dafny.Map
			_ = _1518_v78
			_1518_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if !(_this.F24) {
					return _1511_v71
				}
				return _1511_v71
			})(), (_1517_v77).Dtor_cf7())
			_1518_v78 = (_1518_v78).Update((func() _dafny.CodePoint {
				if (_this).F13() {
					return _1511_v71
				}
				return _1511_v71
			})(), (_1504_v65).F22())
			(globalState).F7 = _1510_i5
		}
		var _1519_i6 _dafny.Int
		_ = _1519_i6
		_1519_i6 = _dafny.Zero
		{
			for !((_1504_v65).F22()) {
				{
					if (_1519_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1519_i6 = (_1519_i6).Plus(_dafny.One)
					if true {
						var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index251
						((_this).F14()).ArraySet1((_1504_v65).F22(), (_index251).Int())
						var _1520_v79 _dafny.Set
						_ = _1520_v79
						_1520_v79 = _dafny.SetOf(_1499_v61, _1499_v61, _1499_v61)
						var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index252
						((_this).F14()).ArraySet1((_dafny.SetOf(_1499_v61)).Equals(_1520_v79), (_index252).Int())
						var _1521_v80 _dafny.CodePoint
						_ = _1521_v80
						_1521_v80 = _dafny.CodePoint('x')
						var _1522_v81 T0
						_ = _1522_v81
						var _nw211 *C0 = New_C0_()
						_ = _nw211
						_nw211.Ctor__(_this.F24, _1521_v80, _this.F12(), (_1504_v65).F22())
						_1522_v81 = _nw211
						var _1523_v82 D16
						_ = _1523_v82
						_1523_v82 = Companion_D16_.Create_DC40_(_1522_v81)
						_1522_v81 = (_1523_v82).Dtor_cf83()
						(globalState).F7 = (_1504_v65).F21()
						var _1524_v83 _dafny.Array
						_ = _1524_v83
						var _nw212 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
						_ = _nw212
						_1524_v83 = _nw212
						var _1525_v84 _dafny.Array
						_ = _1525_v84
						var _nw213 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
						_ = _nw213
						_1525_v84 = _nw213
						var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1524_v83), 0))
						_ = _index253
						(_1524_v83).ArraySet1(_1525_v84, (_index253).Int())
						var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(222), _dafny.ArrayLen((_1524_v83), 0))
						_ = _index254
						(_1524_v83).ArraySet1(_1525_v84, (_index254).Int())
						_1509_v70 = (_1509_v70).Update((_1504_v65).F21(), _1508_v69)
					} else {
						var _1526_v85 _dafny.MultiSet
						_ = _1526_v85
						_1526_v85 = _dafny.MultiSetOf((_this).F14(), (_this).F14())
						(globalState).F7 = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(822), (func() _dafny.Int {
							if (_1526_v85).Contains((_this).F14()) {
								return (_1526_v85).Multiplicity((_this).F14())
							}
							return (_1504_v65).F21()
						})())
						var _1527_v86 _dafny.Map
						_ = _1527_v86
						_1527_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1498_v60, (_this).F13())
						var _rhs217 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer88 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg88 _dafny.Int) interface{} {
								return coer88(arg88)
							}
						}(func(_1528_i7 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('i')
						}))).Cardinality())).Minus((_1527_v86).Cardinality())
						_ = _rhs217
						var _rhs218 _dafny.Int = (((_1504_v65).F21()).Plus(_1421_v2)).Minus((_dafny.Zero).Minus((_1504_v65).F21()))
						_ = _rhs218
						var _rhs219 bool = (_this).F13()
						_ = _rhs219
						var _rhs220 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1504_v65).F21(), _1421_v2))
						_ = _rhs220
						var _rhs221 bool = false
						_ = _rhs221
						var _lhs184 *GlobalState = globalState
						_ = _lhs184
						var _lhs185 *C5 = _this
						_ = _lhs185
						var _lhs186 *C5 = _this
						_ = _lhs186
						r0 = _rhs217
						_lhs184.F7 = _rhs218
						_lhs185.F24 = _rhs219
						r0 = _rhs220
						_lhs186.F24 = _rhs221
						(globalState).F7 = ((func() _dafny.Map {
							if !((_1504_v65).F22()) {
								return _1509_v70
							}
							return (func() _dafny.Map {
								if _this.F24 {
									return _1509_v70
								}
								return func() _dafny.Map {
									var _coll47 = _dafny.NewMapBuilder()
									_ = _coll47
									for _iter55 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(155), _dafny.IntOfInt64(-903))); ; {
										_compr_47, _ok55 := _iter55()
										if !_ok55 {
											break
										}
										var _1529_v87 _dafny.Int
										_1529_v87 = interface{}(_compr_47).(_dafny.Int)
										if ((_dafny.IntOfInt64(155)).Cmp(_1529_v87) <= 0) && ((_1529_v87).Cmp(_dafny.IntOfInt64(-903)) < 0) {
											_coll47.Add(Companion_Default___.SafeDivisionInt(_1529_v87, _1421_v2), _1508_v69)
										}
									}
									return _coll47.ToMap()
								}()
							})()
						})()).Cardinality()
						var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index255
						((_this).F14()).ArraySet1(_this.F24, (_index255).Int())
						var _1530_v88 _dafny.Sequence
						_ = _1530_v88
						_1530_v88 = _dafny.SeqOf(_1498_v60, _1498_v60, _1498_v60)
						var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index256
						var _rhs222 _dafny.Sequence = (_1530_v88).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_1421_v2, (_1504_v65).F21()), _dafny.IntOfUint32((_1530_v88).Cardinality()))).Uint32()).(_dafny.Sequence)
						_ = _rhs222
						var _rhs223 bool = ((_1504_v65).F21()).Cmp((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1504_v65).F21(), (_1504_v65).F21()))) >= 0
						_ = _rhs223
						var _lhs187 _dafny.Array = (_this).F14()
						_ = _lhs187
						var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _lhs188
						_1498_v60 = _rhs222
						(_lhs187).ArraySet1(_rhs223, (_lhs188).Int())
						(_this).F24 = ((_1504_v65).F21()).Cmp((_1504_v65).F21()) == 0
					}
					(_this).F24 = (_this).Fm52((_1504_v65).F21(), _1507_v68, globalState)
					var _1531_v89 _dafny.Sequence
					_ = _1531_v89
					_1531_v89 = _dafny.SeqOf(_1421_v2, _1421_v2, (_1504_v65).F21(), (_1504_v65).F21(), (_1504_v65).F21())
					var _1532_v90 _dafny.MultiSet
					_ = _1532_v90
					_1532_v90 = _dafny.MultiSetOf(_1531_v89)
					if ((func() _dafny.MultiSet {
						if _this.F24 {
							return _dafny.MultiSetOf(_1531_v89)
						}
						return _1532_v90
					})()).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(99))).Uint32(), func(coer89 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg89 _dafny.Int) interface{} {
							return coer89(arg89)
						}
					}((func(_1533_v65 *C2) func(_dafny.Int) _dafny.Int {
						return func(_1534_i8 _dafny.Int) _dafny.Int {
							return (_1533_v65).F21()
						}
					})(_1504_v65))), Companion_Default___.Fm63((_1504_v65).F22(), (_1504_v65).F22(), globalState))) {
						var _1535_v91 _dafny.Map
						_ = _1535_v91
						_1535_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1504_v65).F22(), (_this).F13())
						_1535_v91 = (_1535_v91).Update((_1504_v65).F22(), (_1504_v65).F22())
						var _1536_v92 _dafny.CodePoint
						_ = _1536_v92
						_1536_v92 = _dafny.CodePoint('t')
						var _1537_v93 *C1
						_ = _1537_v93
						var _nw214 *C1 = New_C1_()
						_ = _nw214
						_nw214.Ctor__(_1536_v92, Companion_D1_.Create_DC2_((_1504_v65).F21(), (_1504_v65).F21(), (_1504_v65).F22()), (_this).F13())
						_1537_v93 = _nw214
						var _1538_v94 _dafny.Map
						_ = _1538_v94
						_1538_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1537_v93, (_1504_v65).F22())
						var _1539_v95 _dafny.Map
						_ = _1539_v95
						_1539_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1538_v94).Merge(_1538_v94), (_1504_v65).F21())
						var _rhs224 bool = (_1537_v93).Fm43(globalState)
						_ = _rhs224
						var _rhs225 _dafny.Map = (_1539_v95).Merge(_1539_v95)
						_ = _rhs225
						var _lhs189 *C5 = _this
						_ = _lhs189
						_lhs189.F24 = _rhs224
						_1539_v95 = _rhs225
						(globalState).F7 = (_1504_v65).F21()
						(_this).F24 = (_this.F24) || ((_this).F13())
						var _1540_v96 _dafny.Map
						_ = _1540_v96
						_1540_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(502), (_1504_v65).F21())
						var _1541_v97 _dafny.Map
						_ = _1541_v97
						_1541_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1498_v60, _1540_v96)
						_1541_v97 = (_1541_v97).Update(_1498_v60, _1540_v96)
					} else {
						var _1542_v98 _dafny.MultiSet
						_ = _1542_v98
						_1542_v98 = _dafny.MultiSetOf(_1506_v67, _1506_v67)
						var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index257
						((_this).F14()).ArraySet1((_1542_v98).Equals(_dafny.MultiSetOf(_1506_v67)), (_index257).Int())
						var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen(((_this).F14()), 0))
						_ = _index258
						((_this).F14()).ArraySet1((_this.F24) && ((_this).Fm52((_dafny.MultiSetOf((_1504_v65).F22())).Cardinality(), _1507_v68, globalState)), (_index258).Int())
						r0 = _1421_v2
						(globalState).F7 = (_1504_v65).F21()
						var _1543_v99 _dafny.MultiSet
						_ = _1543_v99
						_1543_v99 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt((_1504_v65).F21(), (_1504_v65).F21()))
						_1421_v2 = (_1543_v99).Cardinality()
						_1498_v60 = (func() _dafny.Sequence {
							if (_1419_v0).IsProperSubsetOf(_dafny.MultiSetOf((_this).F13())) {
								return _1498_v60
							}
							return _1498_v60
						})()
					}
					var _1544_v100 _dafny.Map
					_ = _1544_v100
					_1544_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F24)
					_1544_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_1419_v0).IsDisjointFrom(_dafny.MultiSetOf((_this).F13())))
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		(globalState).F7 = (_dafny.IntOfUint32((Companion_Default___.Fm63(true, _this.F24, globalState)).Cardinality())).Plus(_1421_v2)
		var _1545_v101 D11
		_ = _1545_v101
		_1545_v101 = Companion_D11_.Create_DC27_(false, _dafny.CodePoint('q'), _dafny.IntOfInt64(890), _1421_v2, _dafny.SeqOf(_dafny.IntOfInt64(118)))
		var _1546_i9 _dafny.Int
		_ = _1546_i9
		_1546_i9 = _dafny.Zero
		{
			for (!(_1419_v0).Equals(_1419_v0)) && ((_1545_v101).Dtor_cf51()) {
				{
					if (_1546_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1546_i9 = (_1546_i9).Plus(_dafny.One)
					var _1547_v102 _dafny.Array
					_ = _1547_v102
					var _nw215 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
					_ = _nw215
					_1547_v102 = _nw215
					var _len0_35 _dafny.Int = _dafny.One
					_ = _len0_35
					var _nw216 _dafny.Array
					_ = _nw216
					if _len0_35.Cmp(_dafny.Zero) == 0 {
						_nw216 = _dafny.NewArray(_len0_35)
					} else {
						var _init35 func(_dafny.Int) bool = (func(_1548_v101 D11) func(_dafny.Int) bool {
							return func(_1549_i10 _dafny.Int) bool {
								return (_1548_v101).Dtor_cf51()
							}
						})(_1545_v101)
						_ = _init35
						var _element0_35 = _init35(_dafny.Zero)
						_ = _element0_35
						_nw216 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
						(_nw216).ArraySet1(_element0_35, 0)
						var _nativeLen0_35 = (_len0_35).Int()
						_ = _nativeLen0_35
						for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
							(_nw216).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
						}
					}
					_1547_v102 = _nw216
					_1498_v60 = (func() _dafny.Sequence {
						if (_1504_v65).F22() {
							return _1498_v60
						}
						return _1498_v60
					})()
					var _1550_v103 _dafny.CodePoint
					_ = _1550_v103
					_1550_v103 = _dafny.CodePoint('i')
					var _1551_v104 *C0
					_ = _1551_v104
					var _nw217 *C0 = New_C0_()
					_ = _nw217
					_nw217.Ctor__((_1504_v65).F22(), _1550_v103, _this.F12(), (_1506_v67).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-991), _dafny.IntOfUint32((_1506_v67).Cardinality()))).Uint32()).(bool))
					_1551_v104 = _nw217
					var _1552_v105 _dafny.Map
					_ = _1552_v105
					_1552_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1504_v65).F21(), (_dafny.SetOf((_this).Fm53(_1421_v2, globalState))).Union((_this).F15()))
					_1552_v105 = (_1552_v105).Update((_1504_v65).F21(), (Companion_Default___.Fm64(globalState)).Difference(_dafny.SetOf(_1421_v2)))
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1553_v106 _dafny.Map
		_ = _1553_v106
		_1553_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(134), (_1504_v65).F21())
		r0 = ((_1553_v106).Merge(_1553_v106)).Cardinality()
		return r0
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f12 D1
	_f16 D1
	_f14 _dafny.Array
	_f15 _dafny.Set
	_f13 bool
	_f17 bool
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f12 = Companion_D1_.Default()
	_this._f16 = Companion_D1_.Default()
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f15 = _dafny.EmptySet
	_this._f13 = false
	_this._f17 = false
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
	return [](*_dafny.TraitID){Companion_T3_.TraitID_, Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T3 = &C6{}
var _ T2 = &C6{}
var _ T1 = &C6{}
var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F12() D1 {
	return _this._f12
}
func (_this *C6) F12_set_(value D1) {
	_this._f12 = value
}
func (_this *C6) F16() D1 {
	return _this._f16
}
func (_this *C6) F16_set_(value D1) {
	_this._f16 = value
}
func (_this *C6) F14() _dafny.Array {
	return _this._f14
}
func (_this *C6) F15() _dafny.Set {
	return _this._f15
}
func (_this *C6) F13() bool {
	return _this._f13
}
func (_this *C6) F17() bool {
	return _this._f17
}
func (_this *C6) Ctor__(f16 D1, f17 bool, f14 _dafny.Array, f15 _dafny.Set, f12 D1, f13 bool) {
	{
		(_this)._f16 = f16
		(_this)._f17 = f17
		(_this)._f14 = f14
		(_this)._f15 = f15
		(_this)._f12 = f12
		(_this)._f13 = f13
	}
}
func (_this *C6) Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	{
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf((_this).F13(), (_this).F17())).IsSubsetOf(_dafny.MultiSetOf((_this).F17())), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(705), (_dafny.Zero).Minus(_dafny.IntOfInt64(-805)))).Cardinality()))
	}
}
func (_this *C6) Fm8(p0 bool, p1 D1, globalState *GlobalState) _dafny.Int {
	{
		if (_this).F17() {
			return _dafny.IntOfInt64(-683)
		} else {
			return (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _dafny.SeqOf((_this).F17(), (_this).F13()))).Cardinality())))).Cardinality())
		}
	}
}
func (_this *C6) Fm9(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(711))).Uint32(), func(coer90 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg90 _dafny.Int) interface{} {
				return coer90(arg90)
			}
		}(func(_1554_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('b')
		}))).Cardinality()))).Times((func() _dafny.Map {
			var _coll48 = _dafny.NewMapBuilder()
			_ = _coll48
			for _iter56 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-200), _dafny.IntOfInt64(231))); ; {
				_compr_48, _ok56 := _iter56()
				if !_ok56 {
					break
				}
				var _1555_v0 _dafny.Int
				_1555_v0 = interface{}(_compr_48).(_dafny.Int)
				if ((_dafny.IntOfInt64(-200)).Cmp(_1555_v0) <= 0) && ((_1555_v0).Cmp(_dafny.IntOfInt64(231)) < 0) {
					_coll48.Add(Companion_Default___.SafeDivisionInt(_1555_v0, _dafny.IntOfInt64(374)), _dafny.IntOfInt64(805))
				}
			}
			return _coll48.ToMap()
		}()).Cardinality()), ((_dafny.SetOf((_this).F13(), !((_this).F17()))).Intersection(_dafny.SetOf((_this).F13()))).Cardinality()))
	}
}
func (_this *C6) Fm11(globalState *GlobalState) bool {
	{
		return (_this).F13()
	}
}
func (_this *C6) Fm12(p0 _dafny.Map, p1 D1, globalState *GlobalState) bool {
	{
		return (_this).F17()
	}
}
func (_this *C6) M6(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var _1556_v0 _dafny.Int
		_ = _1556_v0
		_1556_v0 = _dafny.IntOfInt64(622)
		var _hi10 _dafny.Int = (_1556_v0).Minus(_1556_v0)
		_ = _hi10
		for _1557_i0 := (_1556_v0).Minus(_1556_v0); _1557_i0.Cmp(_hi10) < 0; _1557_i0 = _1557_i0.Plus(_dafny.One) {
			var _1558_v1 _dafny.Array
			_ = _1558_v1
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
			_ = _nw218
			_1558_v1 = _nw218
			var _1559_v2 _dafny.Set
			_ = _1559_v2
			_1559_v2 = _dafny.SetOf(p0, (_this).F17())
			var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_1558_v1), 0))
			_ = _index259
			(_1558_v1).ArraySet1((_1559_v2).Union(_1559_v2), (_index259).Int())
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_1558_v1), 0))
			_ = _index260
			(_1558_v1).ArraySet1(_1559_v2, (_index260).Int())
			var _1560_v3 bool
			_ = _1560_v3
			_1560_v3 = false
			var _1561_v4 _dafny.Array
			_ = _1561_v4
			var _nw219 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw219
			_1561_v4 = _nw219
			var _1562_v5 _dafny.MultiSet
			_ = _1562_v5
			_1562_v5 = _dafny.MultiSetOf((_this).F13())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))
			_ = _index261
			(_1561_v4).ArraySet1((_1562_v5).Cardinality(), (_index261).Int())
			var _1563_v6 _dafny.Sequence
			_ = _1563_v6
			_1563_v6 = _dafny.UnicodeSeqOfUtf8Bytes("pjvi")
			var _1564_v7 _dafny.Map
			_ = _1564_v7
			_1564_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1563_v6).Cardinality()), (_this).F13())
			var _1565_v8 _dafny.Sequence
			_ = _1565_v8
			_1565_v8 = _dafny.SeqOf(((_1558_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_1558_v1), 0))).Int()).(_dafny.Set)).Cardinality())
			var _1566_v9 _dafny.MultiSet
			_ = _1566_v9
			_1566_v9 = _dafny.MultiSetOf(_1557_i0)
			var _1567_v10 _dafny.Sequence
			_ = _1567_v10
			_1567_v10 = _dafny.SeqOf(!((func() bool {
				if (_1564_v7).Contains(_1556_v0) {
					return (_1564_v7).Get(_1556_v0).(bool)
				}
				return p0
			})()), (_1566_v9).IsProperSubsetOf(_dafny.MultiSetFromSeq(_1565_v8)))
			var _1568_v11 _dafny.Map
			_ = _1568_v11
			_1568_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm13(_1563_v6, globalState), (_this).F13())
			var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))
			_ = _index262
			var _rhs226 bool = (_1567_v10).Select((Companion_Default___.SafeIndex(_1557_i0, _dafny.IntOfUint32((_1567_v10).Cardinality()))).Uint32()).(bool)
			_ = _rhs226
			var _rhs227 _dafny.Int = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sn")).Cardinality())
			_ = _rhs227
			var _rhs228 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_1556_v0))
			_ = _rhs228
			var _rhs229 _dafny.Int = _1556_v0
			_ = _rhs229
			var _rhs230 bool = !(p0) || ((_this).Fm12(_1568_v11, _this.F12(), globalState))
			_ = _rhs230
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			var _lhs191 _dafny.Array = _1561_v4
			_ = _lhs191
			var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))
			_ = _lhs192
			_1560_v3 = _rhs226
			r0 = _rhs227
			_lhs190.F7 = _rhs228
			(_lhs191).ArraySet1(_rhs229, (_lhs192).Int())
			_1560_v3 = _rhs230
			var _1569_v12 _dafny.CodePoint
			_ = _1569_v12
			_1569_v12 = _dafny.CodePoint('v')
			var _1570_v13 _dafny.Map
			_ = _1570_v13
			_1570_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(181))).Uint32(), func(coer91 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg91 _dafny.Int) interface{} {
					return coer91(arg91)
				}
			}(func(_1571_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(486))).Uint32(), func(coer92 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg92 _dafny.Int) interface{} {
						return coer92(arg92)
					}
				}(func(_1572_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				}))).Cardinality())
			}))).Cardinality()), _1569_v12)
			var _1573_v14 _dafny.Map
			_ = _1573_v14
			_1573_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1561_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))).Int()).(_dafny.Int))
			var _1574_v15 _dafny.Set
			_ = _1574_v15
			_1574_v15 = _dafny.SetOf(_1561_v4)
			var _1575_v16 D2
			_ = _1575_v16
			_1575_v16 = Companion_D2_.Create_DC4_(_1561_v4, p0, _1557_i0, _1560_v3, (_this).F14())
			var _1576_v17 _dafny.Array
			_ = _1576_v17
			var _nwElement0_52 bool = _1560_v3
			_ = _nwElement0_52
			var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(17))
			_ = _nw220
			(_nw220).ArraySet1(_nwElement0_52, 0)
			(_nw220).ArraySet1(_1560_v3, 1)
			(_nw220).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_1563_v6, _1563_v6), 2)
			(_nw220).ArraySet1(p0, 3)
			(_nw220).ArraySet1((_1560_v3) || (_1560_v3), 4)
			(_nw220).ArraySet1(_1560_v3, 5)
			(_nw220).ArraySet1(_1560_v3, 6)
			(_nw220).ArraySet1(p0, 7)
			(_nw220).ArraySet1((false) && (_1560_v3), 8)
			(_nw220).ArraySet1((_this).F17(), 9)
			(_nw220).ArraySet1((_this).F13(), 10)
			(_nw220).ArraySet1((_dafny.MultiSetOf((_1570_v13).Cardinality(), (_dafny.Zero).Minus((_1565_v8).Select((Companion_Default___.SafeIndex((_1561_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1565_v8).Cardinality()))).Uint32()).(_dafny.Int)))).IsDisjointFrom(_dafny.MultiSetOf((_dafny.Zero).Minus(_1557_i0))), 11)
			(_nw220).ArraySet1(_1560_v3, 12)
			(_nw220).ArraySet1(((func() _dafny.Int {
				if (_1573_v14).Contains((_this).F17()) {
					return (_1573_v14).Get((_this).F17()).(_dafny.Int)
				}
				return Companion_Default___.Fm0(_dafny.Companion_Sequence_.Update(_1567_v10, (Companion_Default___.SafeIndex((_1561_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1567_v10).Cardinality()))).Uint32(), (_this).F17()), (_1567_v10).Select((Companion_Default___.SafeIndex((_1561_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1567_v10).Cardinality()))).Uint32()).(bool), (_this).F13(), (_1574_v15).Cardinality(), globalState)
			})()).Cmp(_1556_v0) >= 0, 13)
			(_nw220).ArraySet1((_this).F13(), 14)
			(_nw220).ArraySet1(!((_1567_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(980), _dafny.IntOfUint32((_1567_v10).Cardinality()))).Uint32()).(bool)), 15)
			(_nw220).ArraySet1((_1575_v16).Dtor_cf9(), 16)
			_1576_v17 = _nw220
			_1576_v17 = (_this).F14()
			var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))
			_ = _index263
			(_1561_v4).ArraySet1((_1561_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(344), _dafny.ArrayLen((_1561_v4), 0))).Int()).(_dafny.Int), (_index263).Int())
		}
		var _1577_v18 _dafny.Map
		_ = _1577_v18
		_1577_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
		var _1578_v19 _dafny.Sequence
		_ = _1578_v19
		_1578_v19 = _dafny.SeqOf((func() bool {
			if (_1577_v18).Contains(p0) {
				return (_1577_v18).Get(p0).(bool)
			}
			return !(false)
		})(), (_this).F13())
		if _dafny.Companion_Sequence_.IsPrefixOf(_1578_v19, _dafny.Companion_Sequence_.Concatenate(_1578_v19, _dafny.SeqOf((_this).F17()))) {
			var _1579_v20 _dafny.CodePoint
			_ = _1579_v20
			_1579_v20 = _dafny.CodePoint('b')
			var _1580_v21 _dafny.Sequence
			_ = _1580_v21
			_1580_v21 = _dafny.SeqOf(_1556_v0, _1556_v0, _1556_v0, _1556_v0)
			var _1581_v22 _dafny.Sequence
			_ = _1581_v22
			_1581_v22 = _dafny.SeqOf((_1580_v21).Select((Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1580_v21).Cardinality()))).Uint32()).(_dafny.Int), _1556_v0)
			var _1582_v23 *C1
			_ = _1582_v23
			var _nw221 *C1 = New_C1_()
			_ = _nw221
			_nw221.Ctor__(_1579_v20, _this.F12(), !_dafny.Companion_Sequence_.Equal(_1581_v22, _1580_v21))
			_1582_v23 = _nw221
			_1556_v0 = (_1556_v0).Plus(_1556_v0)
			var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index264
			((_this).F14()).ArraySet1((_this).F17(), (_index264).Int())
			var _1583_v24 _dafny.Sequence
			_ = _1583_v24
			_1583_v24 = _dafny.UnicodeSeqOfUtf8Bytes("mttwbblrp")
			var _1584_v25 _dafny.Map
			_ = _1584_v25
			_1584_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1556_v0)
			var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index265
			var _rhs231 _dafny.Int = Companion_Default___.Fm0(_dafny.Companion_Sequence_.Concatenate(_1578_v19, _1578_v19), !((_this).F17()), (_this).F13(), _1556_v0, globalState)
			_ = _rhs231
			var _rhs232 bool = (_this).F13()
			_ = _rhs232
			var _rhs233 _dafny.CodePoint = Companion_Default___.Fm18((_dafny.Zero).Minus((_dafny.Zero).Minus(_1556_v0)), _1583_v24, (_this).F17(), (func() _dafny.Map {
				if (_this).F13() {
					return _1584_v25
				}
				return (_this).Fm10((_this).F17(), (_this).F13(), _1578_v19, globalState)
			})(), globalState)
			_ = _rhs233
			var _rhs234 _dafny.Int = _1556_v0
			_ = _rhs234
			var _lhs193 *GlobalState = globalState
			_ = _lhs193
			var _lhs194 _dafny.Array = (_this).F14()
			_ = _lhs194
			var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs195
			var _lhs196 *C1 = _1582_v23
			_ = _lhs196
			_lhs193.F7 = _rhs231
			(_lhs194).ArraySet1(_rhs232, (_lhs195).Int())
			_lhs196.F23 = _rhs233
			r0 = _rhs234
			var _1585_v26 _dafny.Array
			_ = _1585_v26
			var _nwElement0_53 _dafny.Int = _1556_v0
			_ = _nwElement0_53
			var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(4))
			_ = _nw222
			(_nw222).ArraySet1(_nwElement0_53, 0)
			(_nw222).ArraySet1(_1556_v0, 1)
			(_nw222).ArraySet1(_1556_v0, 2)
			(_nw222).ArraySet1((_dafny.Zero).Minus(_1556_v0), 3)
			_1585_v26 = _nw222
			var _1586_v27 _dafny.Array
			_ = _1586_v27
			var _nw223 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw223
			_1586_v27 = _nw223
			var _1587_v28 D2
			_ = _1587_v28
			_1587_v28 = Companion_D2_.Create_DC4_(_1585_v26, (_1578_v19).Select((Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1578_v19).Cardinality()))).Uint32()).(bool), _1556_v0, (_this).F13(), _1586_v27)
			var _pat_let_tv39 = _1586_v27
			_ = _pat_let_tv39
			var _pat_let_tv40 = _1556_v0
			_ = _pat_let_tv40
			var _pat_let_tv41 = _1585_v26
			_ = _pat_let_tv41
			_1584_v25 = (_1584_v25).Update((func(_pat_let42_0 D2) D2 {
				return func(_1588_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let43_0 _dafny.Array) D2 {
						return func(_1589_dt__update_hcf10_h0 _dafny.Array) D2 {
							return func(_pat_let44_0 bool) D2 {
								return func(_1590_dt__update_hcf7_h0 bool) D2 {
									return func(_pat_let45_0 _dafny.Int) D2 {
										return func(_1591_dt__update_hcf8_h0 _dafny.Int) D2 {
											return func(_pat_let46_0 _dafny.Array) D2 {
												return func(_1592_dt__update_hcf6_h0 _dafny.Array) D2 {
													return Companion_D2_.Create_DC4_(_1592_dt__update_hcf6_h0, _1590_dt__update_hcf7_h0, _1591_dt__update_hcf8_h0, (_1588_dt__update__tmp_h0).Dtor_cf9(), _1589_dt__update_hcf10_h0)
												}(_pat_let46_0)
											}(_pat_let_tv41)
										}(_pat_let45_0)
									}(_pat_let_tv40)
								}(_pat_let44_0)
							}(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(608), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
						}(_pat_let43_0)
					}(_pat_let_tv39)
				}(_pat_let42_0)
			}(_1587_v28)).Dtor_cf9(), Companion_Default___.SafeDivisionInt(_1556_v0, _1556_v0))
			(globalState).F7 = (_1556_v0).Plus(_1556_v0)
		} else {
			var _1593_v29 *C4
			_ = _1593_v29
			var _nw224 *C4 = New_C4_()
			_ = _nw224
			_nw224.Ctor__()
			_1593_v29 = _nw224
			var _nw225 *C4 = New_C4_()
			_ = _nw225
			_nw225.Ctor__()
			_1593_v29 = _nw225
			var _1594_v30 _dafny.Sequence
			_ = _1594_v30
			_1594_v30 = _dafny.SeqOf((_1556_v0).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tfumdn")).Cardinality())), _1556_v0)
			_1594_v30 = _dafny.Companion_Sequence_.Concatenate(_1594_v30, _1594_v30)
			(globalState).F7 = _1556_v0
			r0 = (_1556_v0).Minus((_this.F12()).Dtor_cf3())
			var _1595_v31 D2
			_ = _1595_v31
			_1595_v31 = Companion_D2_.Create_DC3_((_this).F14())
			var _1596_v32 _dafny.Array
			_ = _1596_v32
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(10)
			_ = _len0_36
			var _nw226 _dafny.Array
			_ = _nw226
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw226 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = func(_1597_i3 _dafny.Int) _dafny.Int {
					return (_1597_i3).Plus(_dafny.IntOfInt64(118))
				}
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw226 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw226).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw226).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1596_v32 = _nw226
			var _1598_v33 D2
			_ = _1598_v33
			_1598_v33 = Companion_D2_.Create_DC4_(_1596_v32, (_this).F17(), _1556_v0, false, (_this).F14())
			var _1599_v34 _dafny.Map
			_ = _1599_v34
			_1599_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F17()), (_this).F14())
			var _1600_v35 _dafny.Array
			_ = _1600_v35
			var _nwElement0_54 _dafny.Array = (_this).F14()
			_ = _nwElement0_54
			var _nw227 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.IntOfInt64(17))
			_ = _nw227
			(_nw227).ArraySet1(_nwElement0_54, 0)
			(_nw227).ArraySet1((_1595_v31).Dtor_cf5(), 1)
			(_nw227).ArraySet1((_this).F14(), 2)
			(_nw227).ArraySet1((_1598_v33).Dtor_cf10(), 3)
			(_nw227).ArraySet1((_this).F14(), 4)
			(_nw227).ArraySet1((_this).F14(), 5)
			(_nw227).ArraySet1((_1598_v33).Dtor_cf10(), 6)
			(_nw227).ArraySet1((_this).F14(), 7)
			(_nw227).ArraySet1((_this).F14(), 8)
			(_nw227).ArraySet1((_this).F14(), 9)
			(_nw227).ArraySet1((_this).F14(), 10)
			(_nw227).ArraySet1((func() _dafny.Array {
				if (_1599_v34).Contains(true) {
					return (_1599_v34).Get(true).(_dafny.Array)
				}
				return (_this).F14()
			})(), 11)
			(_nw227).ArraySet1((_this).F14(), 12)
			(_nw227).ArraySet1((_this).F14(), 13)
			(_nw227).ArraySet1((_this).F14(), 14)
			(_nw227).ArraySet1((_this).F14(), 15)
			(_nw227).ArraySet1((_this).F14(), 16)
			_1600_v35 = _nw227
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_1600_v35), 0))
			_ = _index266
			(_1600_v35).ArraySet1((_this).F14(), (_index266).Int())
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_1600_v35), 0))
			_ = _index267
			(_1600_v35).ArraySet1((_this).F14(), (_index267).Int())
		}
		var _1601_v36 _dafny.Array
		_ = _1601_v36
		var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(26))
		_ = _nw228
		_1601_v36 = _nw228
		var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.One)
		_ = _nw229
		_1601_v36 = _nw229
		var _1602_v37 _dafny.Array
		_ = _1602_v37
		var _nw230 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw230
		_1602_v37 = _nw230
		for _iter57 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1602_v37), 0))); ; {
			_guard_loop_6, _ok57 := _iter57()
			if !_ok57 {
				break
			}
			var _1603_i4 _dafny.Int
			_1603_i4 = interface{}(_guard_loop_6).(_dafny.Int)
			if (true) && (((_1603_i4).Sign() != -1) && ((_1603_i4).Cmp(_dafny.ArrayLen((_1602_v37), 0)) < 0)) {
				(_1602_v37).ArraySet1(Companion_Default___.SafeModuloInt(_1603_i4, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_1556_v0)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dscsy")).Cardinality()))), (_1603_i4).Int())
			}
		}
		var _hi11 _dafny.Int = _1556_v0
		_ = _hi11
		for _1604_i5 := _1556_v0; _1604_i5.Cmp(_hi11) < 0; _1604_i5 = _1604_i5.Plus(_dafny.One) {
			var _1605_v38 bool
			_ = _1605_v38
			_1605_v38 = false
			_1605_v38 = !(p0)
			var _1606_v39 _dafny.MultiSet
			_ = _1606_v39
			_1606_v39 = _dafny.MultiSetOf(_1556_v0)
			var _1607_v40 _dafny.Map
			_ = _1607_v40
			_1607_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1606_v39)
			(globalState).F2 = _dafny.Companion_Sequence_.Update(_1578_v19, (Companion_Default___.SafeIndex(((func() _dafny.MultiSet {
				if (_1607_v40).Contains(!(p0)) {
					return (_1607_v40).Get(!(p0)).(_dafny.MultiSet)
				}
				return _1606_v39
			})()).Cardinality(), _dafny.IntOfUint32((_1578_v19).Cardinality()))).Uint32(), (func() bool {
				if true {
					return _1605_v38
				}
				return (_this).F13()
			})())
			var _1608_v41 _dafny.Sequence
			_ = _1608_v41
			_1608_v41 = _dafny.UnicodeSeqOfUtf8Bytes("fbqky")
			var _1609_v42 _dafny.Map
			_ = _1609_v42
			_1609_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-435), p0)
			var _rhs235 _dafny.Int = Companion_Default___.SafeModuloInt(_1556_v0, _1604_i5)
			_ = _rhs235
			var _rhs236 bool = !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-811))).Uint32(), func(coer93 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg93 _dafny.Int) interface{} {
					return coer93(arg93)
				}
			}(func(_1610_i6 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), _1608_v41)
			_ = _rhs236
			var _rhs237 _dafny.Map = (func() _dafny.Map {
				if (func() bool {
					if (_1609_v42).Contains(_1604_i5) {
						return (_1609_v42).Get(_1604_i5).(bool)
					}
					return _1605_v38
				})() {
					return _1577_v18
				}
				return (_1577_v18).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (_this).F17()))
			})()
			_ = _rhs237
			_1556_v0 = _rhs235
			_1605_v38 = _rhs236
			_1577_v18 = _rhs237
			var _1611_v43 D2
			_ = _1611_v43
			_1611_v43 = Companion_D2_.Create_DC3_((_this).F14())
			_1611_v43 = func(_pat_let47_0 D2) D2 {
				return func(_1612_dt__update__tmp_h1 D2) D2 {
					return func(_pat_let48_0 _dafny.Array) D2 {
						return func(_1613_dt__update_hcf5_h0 _dafny.Array) D2 {
							return Companion_D2_.Create_DC3_(_1613_dt__update_hcf5_h0)
						}(_pat_let48_0)
					}((_this).F14())
				}(_pat_let47_0)
			}(_1611_v43)
		}
		var _1614_v44 _dafny.Sequence
		_ = _1614_v44
		_1614_v44 = _dafny.UnicodeSeqOfUtf8Bytes("axb")
		if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.UnicodeSeqOfUtf8Bytes("gsllko"))).Cardinality())).Cmp((_1556_v0).Times(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(435))).Uint32(), func(coer94 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg94 _dafny.Int) interface{} {
				return coer94(arg94)
			}
		}((func(_1615_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1616_i7 _dafny.Int) _dafny.Int {
				return _1615_v0
			}
		})(_1556_v0)))).Cardinality()))) > 0 {
			var _1617_v45 _dafny.CodePoint
			_ = _1617_v45
			_1617_v45 = _dafny.CodePoint('h')
			_1617_v45 = _1617_v45
			var _1618_v46 _dafny.MultiSet
			_ = _1618_v46
			_1618_v46 = _dafny.MultiSetOf(_1602_v37)
			var _1619_v47 _dafny.Map
			_ = _1619_v47
			_1619_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1617_v45, (_this).F13())
			var _1620_v48 *C0
			_ = _1620_v48
			var _nw231 *C0 = New_C0_()
			_ = _nw231
			_nw231.Ctor__((_1618_v46).IsSubsetOf(_1618_v46), (func() _dafny.CodePoint {
				if (_this).F13() {
					return _1617_v45
				}
				return _1617_v45
			})(), _this.F12(), (_this).Fm12(_1619_v47, _this.F12(), globalState))
			_1620_v48 = _nw231
			var _1621_v49 _dafny.Array
			_ = _1621_v49
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_37
			var _nw232 _dafny.Array
			_ = _nw232
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw232 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) D7 = (func(_1622_v0 _dafny.Int) func(_dafny.Int) D7 {
					return func(_1623_i8 _dafny.Int) D7 {
						return Companion_D7_.Create_DC16_(_1622_v0)
					}
				})(_1556_v0)
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw232 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw232).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw232).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1621_v49 = _nw232
			var _1624_v50 _dafny.Map
			_ = _1624_v50
			_1624_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm46(_1556_v0, (func() bool {
				if (_1577_v18).Contains((_this).F17()) {
					return (_1577_v18).Get((_this).F17()).(bool)
				}
				return (_this.F12()).Dtor_cf4()
			})(), (_this).F17(), globalState), _1621_v49)
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))
			_ = _index268
			(_1602_v37).ArraySet1((_1624_v50).Cardinality(), (_index268).Int())
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))
			_ = _index269
			(_1602_v37).ArraySet1((Companion_Default___.Fm47(p0, globalState)).Cardinality(), (_index269).Int())
			var _1625_v51 _dafny.Map
			_ = _1625_v51
			_1625_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int), _1614_v44)
			var _1626_v52 _dafny.Array
			_ = _1626_v52
			var _nwElement0_55 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("wnydrhnob")
			_ = _nwElement0_55
			var _nw233 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(23))
			_ = _nw233
			(_nw233).ArraySet1(_nwElement0_55, 0)
			(_nw233).ArraySet1(_1614_v44, 1)
			(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-269))).Uint32(), func(coer95 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg95 _dafny.Int) interface{} {
					return coer95(arg95)
				}
			}((func(_1627_v48 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_1628_i9 _dafny.Int) _dafny.CodePoint {
					return (_1627_v48).F20()
				}
			})(_1620_v48))), 2)
			(_nw233).ArraySet1(_1614_v44, 3)
			(_nw233).ArraySet1(_1614_v44, 4)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fh"), _1614_v44), 5)
			(_nw233).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("orofsd"), 6)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1614_v44, _1614_v44), (Companion_Default___.SafeIndex((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1614_v44, _1614_v44)).Cardinality()))).Uint32(), _dafny.CodePoint('f')), 7)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ftv"), _1614_v44), 8)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(891))).Uint32(), func(coer96 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg96 _dafny.Int) interface{} {
					return coer96(arg96)
				}
			}(func(_1629_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			}))), 9)
			(_nw233).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("lgqso"), 10)
			(_nw233).ArraySet1(_1614_v44, 11)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.Companion_Sequence_.Update(_1614_v44, (Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1614_v44).Cardinality()))).Uint32(), _dafny.CodePoint('f'))), 12)
			(_nw233).ArraySet1(_1614_v44, 13)
			(_nw233).ArraySet1((func() _dafny.Sequence {
				if (_1625_v51).Contains((_dafny.Zero).Minus((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int))) {
					return (_1625_v51).Get((_dafny.Zero).Minus((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int))).(_dafny.Sequence)
				}
				return _1614_v44
			})(), 14)
			(_nw233).ArraySet1(_1614_v44, 15)
			(_nw233).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-534))).Uint32(), func(coer97 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg97 _dafny.Int) interface{} {
					return coer97(arg97)
				}
			}((func(_1630_v48 *C0) func(_dafny.Int) _dafny.CodePoint {
				return func(_1631_i11 _dafny.Int) _dafny.CodePoint {
					return (_1630_v48).F20()
				}
			})(_1620_v48))), 16)
			(_nw233).ArraySet1(_1614_v44, 17)
			(_nw233).ArraySet1(_1614_v44, 18)
			(_nw233).ArraySet1(_1614_v44, 19)
			(_nw233).ArraySet1(_1614_v44, 20)
			(_nw233).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.UnicodeSeqOfUtf8Bytes("xbclijk")), 21)
			(_nw233).ArraySet1(_1614_v44, 22)
			_1626_v52 = _nw233
			_1626_v52 = _1626_v52
			var _1632_v53 _dafny.Map
			_ = _1632_v53
			_1632_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), p0)
			var _1633_v54 _dafny.Array
			_ = _1633_v54
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
			_ = _nw234
			_1633_v54 = _nw234
			var _1634_v55 D2
			_ = _1634_v55
			_1634_v55 = Companion_D2_.Create_DC4_(_1633_v54, _1620_v48.F19, _1556_v0, _1620_v48.F19, (_this).F14())
			var _pat_let_tv42 = _1602_v37
			_ = _pat_let_tv42
			var _pat_let_tv43 = _1602_v37
			_ = _pat_let_tv43
			_1632_v53 = (_1632_v53).Update((_this).F14(), (func(_pat_let49_0 D2) D2 {
				return func(_1635_dt__update__tmp_h2 D2) D2 {
					return func(_pat_let50_0 _dafny.Int) D2 {
						return func(_1636_dt__update_hcf8_h1 _dafny.Int) D2 {
							return func(_pat_let51_0 bool) D2 {
								return func(_1637_dt__update_hcf9_h0 bool) D2 {
									return Companion_D2_.Create_DC4_((_1635_dt__update__tmp_h2).Dtor_cf6(), (_1635_dt__update__tmp_h2).Dtor_cf7(), _1636_dt__update_hcf8_h1, _1637_dt__update_hcf9_h0, (_1635_dt__update__tmp_h2).Dtor_cf10())
								}(_pat_let51_0)
							}((_this).F17())
						}(_pat_let50_0)
					}((_pat_let_tv43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(942), _dafny.ArrayLen((_pat_let_tv42), 0))).Int()).(_dafny.Int))
				}(_pat_let49_0)
			}(_1634_v55)).Dtor_cf9())
		} else {
			var _1638_v56 bool
			_ = _1638_v56
			_1638_v56 = false
			_1638_v56 = Companion_Default___.Fm17(globalState)
			r0 = _1556_v0
			var _1639_v57 _dafny.CodePoint
			_ = _1639_v57
			_1639_v57 = _dafny.CodePoint('q')
			var _rhs238 _dafny.Int = _dafny.IntOfInt64(486)
			_ = _rhs238
			var _rhs239 _dafny.CodePoint = _1639_v57
			_ = _rhs239
			r0 = _rhs238
			_1639_v57 = _rhs239
			var _1640_v58 _dafny.Sequence
			_ = _1640_v58
			_1640_v58 = _dafny.SeqOf(_1556_v0)
			if (func() _dafny.Set {
				var _coll49 = _dafny.NewBuilder()
				_ = _coll49
				for _iter58 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1640_v58).Select((Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1640_v58).Cardinality()))).Uint32()).(_dafny.Int), p0)).Keys().Elements()); ; {
					_compr_49, _ok58 := _iter58()
					if !_ok58 {
						break
					}
					var _1641_v59 _dafny.Int
					_1641_v59 = interface{}(_compr_49).(_dafny.Int)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1640_v58).Select((Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1640_v58).Cardinality()))).Uint32()).(_dafny.Int), p0)).Contains(_1641_v59) {
						_coll49.Add((_1641_v59).Plus(_dafny.IntOfInt64(-552)))
					}
				}
				return _coll49.ToSet()
			}()).IsDisjointFrom((_this).F15()) {
				var _1642_v60 _dafny.Map
				_ = _1642_v60
				_1642_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(309), _1556_v0)
				var _1643_v61 _dafny.MultiSet
				_ = _1643_v61
				_1643_v61 = _dafny.MultiSetOf((func() _dafny.Int {
					if (_1642_v60).Contains((_dafny.Zero).Minus(_1556_v0)) {
						return (_1642_v60).Get((_dafny.Zero).Minus(_1556_v0)).(_dafny.Int)
					}
					return _1556_v0
				})(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm27(globalState), _dafny.SeqOf(_dafny.IntOfInt64(126), _1556_v0))).Cardinality())))
				_1643_v61 = (_1643_v61).Union(_1643_v61)
				var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index270
				((_this).F14()).ArraySet1((_this).F13(), (_index270).Int())
				var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index271
				((_this).F14()).ArraySet1((_this).F13(), (_index271).Int())
				var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index272
				((_this).F14()).ArraySet1(!(!(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))), (_index272).Int())
				var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index273
				(_1602_v37).ArraySet1(_1556_v0, (_index273).Int())
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index274
				(_1602_v37).ArraySet1(_1556_v0, (_index274).Int())
				(globalState).F7 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.Companion_Sequence_.Update(_1614_v44, (Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1614_v44).Cardinality()))).Uint32(), _1639_v57)), (Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1614_v44, _dafny.Companion_Sequence_.Update(_1614_v44, (Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1614_v44).Cardinality()))).Uint32(), _1639_v57))).Cardinality()))).Uint32(), _1639_v57)).Cardinality())
			} else {
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index275
				((_this).F14()).ArraySet1(!(p0) || (p0), (_index275).Int())
				var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index276
				((_this).F14()).ArraySet1((_this).F13(), (_index276).Int())
				(globalState).F7 = (_this).Fm8((_this).F17(), _this.F16(), globalState)
				var _1644_v62 _dafny.Sequence
				_ = _1644_v62
				_1644_v62 = _dafny.SeqOf(_1614_v44)
				var _1645_v63 _dafny.Map
				_ = _1645_v63
				_1645_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm38((_this).F13(), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), false, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), globalState), !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Update(_1644_v62, (Companion_Default___.SafeIndex(_1556_v0, _dafny.IntOfUint32((_1644_v62).Cardinality()))).Uint32(), _1614_v44), _1614_v44))
				var _1646_v64 _dafny.Map
				_ = _1646_v64
				_1646_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1578_v19).Cardinality()), !(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)))
				var _1647_v65 _dafny.Map
				_ = _1647_v65
				_1647_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1646_v64).Update(_dafny.IntOfUint32((_1578_v19).Cardinality()), (_this).F13())).Cardinality(), (_this).F13())
				_1645_v63 = (_1645_v63).Update(_dafny.Companion_Sequence_.Concatenate(_1578_v19, _1578_v19), (func() bool {
					if (_this).F13() {
						return _1638_v56
					}
					return (func() bool {
						if (_1647_v65).Contains(_1556_v0) {
							return (_1647_v65).Get(_1556_v0).(bool)
						}
						return ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(411), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
					})()
				})())
				var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index277
				(_1602_v37).ArraySet1(((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm20(globalState)).Cardinality()))).Plus(_1556_v0), (_index277).Int())
				var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index278
				var _rhs240 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).Fm9(_dafny.CodePoint('g'), globalState), _1556_v0)
				_ = _rhs240
				var _rhs241 _dafny.CodePoint = _1639_v57
				_ = _rhs241
				var _rhs242 _dafny.Int = (_1556_v0).Minus(_dafny.IntOfInt64(-984))
				_ = _rhs242
				var _lhs197 _dafny.Array = _1602_v37
				_ = _lhs197
				var _lhs198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1602_v37), 0))
				_ = _lhs198
				(_lhs197).ArraySet1(_rhs240, (_lhs198).Int())
				_1639_v57 = _rhs241
				_1556_v0 = _rhs242
				var _1648_v66 _dafny.Set
				_ = _1648_v66
				_1648_v66 = _dafny.SetOf((_dafny.Zero).Minus((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(584), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int)))
				_1648_v66 = (_this).F15()
			}
			if _1638_v56 {
				var _1649_v67 D9
				_ = _1649_v67
				_1649_v67 = Companion_D9_.Create_DC20_((_this).F15())
				var _1650_v68 _dafny.MultiSet
				_ = _1650_v68
				_1650_v68 = _dafny.MultiSetOf(_1649_v67, _1649_v67)
				var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index279
				(_1602_v37).ArraySet1(_1556_v0, (_index279).Int())
				var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))
				_ = _index280
				var _rhs243 _dafny.MultiSet = ((_1650_v68).Intersection(_1650_v68)).Difference(_1650_v68)
				_ = _rhs243
				var _rhs244 _dafny.Int = _1556_v0
				_ = _rhs244
				var _rhs245 bool = true
				_ = _rhs245
				var _rhs246 _dafny.Sequence = _1578_v19
				_ = _rhs246
				var _lhs199 _dafny.Array = _1602_v37
				_ = _lhs199
				var _lhs200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))
				_ = _lhs200
				var _lhs201 *GlobalState = globalState
				_ = _lhs201
				_1650_v68 = _rhs243
				(_lhs199).ArraySet1(_rhs244, (_lhs200).Int())
				_1638_v56 = _rhs245
				_lhs201.F2 = _rhs246
				var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index281
				((_this).F14()).ArraySet1(true, (_index281).Int())
				var _1651_v69 *C3
				_ = _1651_v69
				var _nw235 *C3 = New_C3_()
				_ = _nw235
				_nw235.Ctor__((_this).F13(), _this.F16(), _1638_v56, (_this).F14(), (_this).F15(), Companion_D1_.Create_DC2_(_1556_v0, (_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int), (_this).F17()), _1638_v56)
				_1651_v69 = _nw235
				var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index282
				((_this).F14()).ArraySet1(Companion_Default___.Fm30(((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int)).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1651_v69)).Cardinality()), globalState), (_index282).Int())
				var _1652_v70 _dafny.Map
				_ = _1652_v70
				_1652_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int), (_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int))
				var _1653_v71 _dafny.Map
				_ = _1653_v71
				_1653_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1652_v70, Companion_Default___.Fm0(_1578_v19, !(true), p0, (_1602_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(863), _dafny.ArrayLen((_1602_v37), 0))).Int()).(_dafny.Int), globalState))
				var _1654_v72 _dafny.MultiSet
				_ = _1654_v72
				_1654_v72 = _dafny.MultiSetOf(p0, _1638_v56, p0)
				var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index283
				var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index284
				var _rhs247 _dafny.Sequence = Companion_Default___.Fm38(!(false) || ((_this).F13()), (func() bool {
					if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
						return (_this).F17()
					}
					return _1638_v56
				})(), _1638_v56, _1638_v56, globalState)
				_ = _rhs247
				var _rhs248 bool = (((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)) && ((_1651_v69).F18())
				_ = _rhs248
				var _rhs249 _dafny.Map = (Companion_D11_.Create_DC25_(_1653_v71)).Dtor_cf48()
				_ = _rhs249
				var _rhs250 _dafny.Map = (_1577_v18).Merge((_1577_v18).Update((_this).F13(), (_this).F17()))
				_ = _rhs250
				var _rhs251 bool = ((_1654_v72).Intersection(_dafny.MultiSetFromSeq(_1578_v19))).Equals(_1654_v72)
				_ = _rhs251
				var _lhs202 *GlobalState = globalState
				_ = _lhs202
				var _lhs203 _dafny.Array = (_this).F14()
				_ = _lhs203
				var _lhs204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs204
				var _lhs205 _dafny.Array = (_this).F14()
				_ = _lhs205
				var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(181), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs206
				_lhs202.F2 = _rhs247
				(_lhs203).ArraySet1(_rhs248, (_lhs204).Int())
				_1653_v71 = _rhs249
				_1577_v18 = _rhs250
				(_lhs205).ArraySet1(_rhs251, (_lhs206).Int())
				var _1655_v73 _dafny.Map
				_ = _1655_v73
				_1655_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((((_this).F15()).Union((_this).F15())).Cardinality(), ((_this).F13()) && ((_1651_v69).F18()))
				_1655_v73 = (_1655_v73).Merge(_1655_v73)
				_1655_v73 = (_1655_v73).Update(Companion_Default___.SafeModuloInt(_1556_v0, _dafny.IntOfInt64(671)), true)
			} else {
				var _1656_v74 D11
				_ = _1656_v74
				_1656_v74 = Companion_D11_.Create_DC26_(_1602_v37, _1614_v44)
				var _1657_v75 D11
				_ = _1657_v75
				_1657_v75 = Companion_D11_.Create_DC28_(_1656_v74)
				var _1658_v76 _dafny.Array
				_ = _1658_v76
				var _nwElement0_56 D11 = _1657_v75
				_ = _nwElement0_56
				var _nw236 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(9))
				_ = _nw236
				(_nw236).ArraySet1(_nwElement0_56, 0)
				(_nw236).ArraySet1(_1657_v75, 1)
				(_nw236).ArraySet1(_1657_v75, 2)
				(_nw236).ArraySet1(_1657_v75, 3)
				(_nw236).ArraySet1(_1657_v75, 4)
				(_nw236).ArraySet1(Companion_D11_.Create_DC28_(_1656_v74), 5)
				(_nw236).ArraySet1(_1657_v75, 6)
				(_nw236).ArraySet1(Companion_D11_.Create_DC28_(_1656_v74), 7)
				(_nw236).ArraySet1(Companion_D11_.Create_DC28_(_1656_v74), 8)
				_1658_v76 = _nw236
				var _1659_v77 _dafny.Sequence
				_ = _1659_v77
				_1659_v77 = _dafny.SeqOf(_1658_v76)
				_1659_v77 = _1659_v77
				_1638_v56 = ((_this).F15()).IsDisjointFrom(Companion_Default___.Fm47((_this).F13(), globalState))
				_1602_v37 = _1602_v37
				var _1660_v78 _dafny.Map
				_ = _1660_v78
				_1660_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1556_v0).Times(_1556_v0))
				_1660_v78 = (_1660_v78).Update((_this).F17(), _1556_v0)
				var _1661_v79 _dafny.Map
				_ = _1661_v79
				_1661_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1556_v0, _1639_v57)
				var _1662_v80 D3
				_ = _1662_v80
				_1662_v80 = Companion_D3_.Create_DC5_(_1639_v57)
				_1661_v79 = (_1661_v79).Update(_1556_v0, (_1662_v80).Dtor_cf11())
			}
		}
		r0 = _dafny.IntOfInt64(411)
		return r0
	}
}
func (_this *C6) M7(globalState *GlobalState) {
	{
		var _1663_i0 _dafny.Int
		_ = _1663_i0
		_1663_i0 = _dafny.Zero
		{
			for (_this).F13() {
				{
					if (_1663_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1663_i0 = (_1663_i0).Plus(_dafny.One)
					var _1664_v0 _dafny.CodePoint
					_ = _1664_v0
					_1664_v0 = _dafny.CodePoint('c')
					var _1665_v1 *C0
					_ = _1665_v1
					var _nw237 *C0 = New_C0_()
					_ = _nw237
					_nw237.Ctor__(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("erlwcj"), _1664_v0), _1664_v0, _this.F12(), (_this).F13())
					_1665_v1 = _nw237
					var _1666_v2 _dafny.Int
					_ = _1666_v2
					_1666_v2 = _dafny.IntOfInt64(697)
					var _1667_v3 *C1
					_ = _1667_v3
					var _nw238 *C1 = New_C1_()
					_ = _nw238
					_nw238.Ctor__(_1664_v0, Companion_D1_.Create_DC2_(_1666_v2, _1666_v2, (_this).F13()), (_this).F17())
					_1667_v3 = _nw238
					var _1668_v4 _dafny.Map
					_ = _1668_v4
					_1668_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1665_v1.F19), _1667_v3)
					var _1669_v5 _dafny.Map
					_ = _1669_v5
					_1669_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1664_v0, _1665_v1.F19)
					var _1670_v6 _dafny.Sequence
					_ = _1670_v6
					_1670_v6 = _dafny.SeqOf((_this).F17())
					_1668_v4 = (_1668_v4).Update((_this).Fm12(_1669_v5, Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_1670_v6).Cardinality()), _1666_v2, (_this).F17()), globalState), _1667_v3)
					var _1671_v7 D3
					_ = _1671_v7
					_1671_v7 = Companion_D3_.Create_DC5_(Companion_Default___.Fm33(_1665_v1.F19, _dafny.IntOfInt64(4), globalState))
					_1671_v7 = _1671_v7
					var _1672_v8 *C3
					_ = _1672_v8
					var _nw239 *C3 = New_C3_()
					_ = _nw239
					_nw239.Ctor__((_this).F17(), Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("yitlqi")), (_this).Fm12(_1669_v5, Companion_D1_.Create_DC2_(_1666_v2, _dafny.IntOfInt64(219), (_this).F17()), globalState), (_this).F14(), (_this).F15(), Companion_D1_.Create_DC2_(_1666_v2, _dafny.IntOfInt64(478), false), _1665_v1.F19)
					_1672_v8 = _nw239
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1673_v9 _dafny.Int
		_ = _1673_v9
		_1673_v9 = _dafny.IntOfInt64(993)
		(globalState).F7 = _1673_v9
		_1673_v9 = _1673_v9
		var _1674_v10 _dafny.Sequence
		_ = _1674_v10
		_1674_v10 = _dafny.SeqOf(true)
		var _1675_v11 _dafny.Sequence
		_ = _1675_v11
		_1675_v11 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_1674_v10, _1674_v10))
		(globalState).F2 = (_1675_v11).Select((Companion_Default___.SafeIndex(_1673_v9, _dafny.IntOfUint32((_1675_v11).Cardinality()))).Uint32()).(_dafny.Sequence)
		var _1676_v12 _dafny.Sequence
		_ = _1676_v12
		_1676_v12 = _dafny.UnicodeSeqOfUtf8Bytes("r")
		var _1677_v13 _dafny.Set
		_ = _1677_v13
		_1677_v13 = _dafny.SetOf(_1676_v12, _1676_v12)
		var _1678_v14 _dafny.Set
		_ = _1678_v14
		_1678_v14 = (func() _dafny.Set {
			if (_this).F13() {
				return Companion_Default___.Fm48(globalState)
			}
			return _1677_v13
		})()
		_1678_v14 = _1678_v14
		var _1679_v15 _dafny.Array
		_ = _1679_v15
		var _nw240 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
		_ = _nw240
		_1679_v15 = _nw240
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1679_v15), 0))); ; {
			_guard_loop_7, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1680_i1 _dafny.Int
			_1680_i1 = interface{}(_guard_loop_7).(_dafny.Int)
			if (true) && (((_1680_i1).Sign() != -1) && ((_1680_i1).Cmp(_dafny.ArrayLen((_1679_v15), 0)) < 0)) {
				(_1679_v15).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1673_v9)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-165)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), _1673_v9)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1673_v9))), (_1680_i1).Int())
			}
		}
	}
}
func (_this *C6) M5(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1681_i0 _dafny.Int
		_ = _1681_i0
		_1681_i0 = _dafny.Zero
		{
			for (_this).F17() {
				{
					if (_1681_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1681_i0 = (_1681_i0).Plus(_dafny.One)
					var _1682_v0 _dafny.Sequence
					_ = _1682_v0
					_1682_v0 = _dafny.SeqOf((_this).F17(), (_this).F13())
					var _1683_v1 _dafny.CodePoint
					_ = _1683_v1
					_1683_v1 = _dafny.CodePoint('o')
					var _1684_v2 D3
					_ = _1684_v2
					_1684_v2 = Companion_D3_.Create_DC5_(_1683_v1)
					var _1685_v3 _dafny.MultiSet
					_ = _1685_v3
					_1685_v3 = _dafny.MultiSetOf((_this).F17(), (_1682_v0).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1682_v0).Cardinality()))).Uint32()).(bool), (_this).Fm12(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1684_v2).Dtor_cf11(), (_this).F17()), Companion_D1_.Create_DC2_(p0, p0, (_this).F13()), globalState))
					var _1686_v4 _dafny.Map
					_ = _1686_v4
					_1686_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1683_v1, p0)
					var _1687_v5 _dafny.Array
					_ = _1687_v5
					var _len0_38 _dafny.Int = _dafny.IntOfInt64(20)
					_ = _len0_38
					var _nw241 _dafny.Array
					_ = _nw241
					if _len0_38.Cmp(_dafny.Zero) == 0 {
						_nw241 = _dafny.NewArray(_len0_38)
					} else {
						var _init38 func(_dafny.Int) _dafny.Int = func(_1688_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1688_i1, (_dafny.SetOf((_this).F17(), !((_this).F13()), (_this).F17(), (_this).F13())).Cardinality())
						}
						_ = _init38
						var _element0_38 = _init38(_dafny.Zero)
						_ = _element0_38
						_nw241 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
						(_nw241).ArraySet1(_element0_38, 0)
						var _nativeLen0_38 = (_len0_38).Int()
						_ = _nativeLen0_38
						for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
							(_nw241).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
						}
					}
					_1687_v5 = _nw241
					var _1689_v6 _dafny.Map
					_ = _1689_v6
					_1689_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1686_v4, _1687_v5)
					var _1690_v7 _dafny.Sequence
					_ = _1690_v7
					_1690_v7 = _dafny.SeqOf(p0, p0)
					var _1691_v8 _dafny.Set
					_ = _1691_v8
					_1691_v8 = _dafny.SetOf((_this).F17(), (_this).F17())
					var _1692_v9 _dafny.Map
					_ = _1692_v9
					_1692_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v8, _dafny.SetOf(p0))
					var _1693_v10 _dafny.Array
					_ = _1693_v10
					var _nwElement0_57 _dafny.Int = (_dafny.Zero).Minus((_1685_v3).Cardinality())
					_ = _nwElement0_57
					var _nw242 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(13))
					_ = _nw242
					(_nw242).ArraySet1(_nwElement0_57, 0)
					(_nw242).ArraySet1(((func() _dafny.Map {
						if false {
							return _1689_v6
						}
						return _1689_v6
					})()).Cardinality(), 1)
					(_nw242).ArraySet1(p0, 2)
					(_nw242).ArraySet1((p0).Times(p0), 3)
					(_nw242).ArraySet1((_dafny.SetOf((_1690_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1690_v7).Cardinality()))).Uint32()).(_dafny.Int), p0)).Cardinality(), 4)
					(_nw242).ArraySet1(((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(740))).Uint32(), func(coer98 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg98 _dafny.Int) interface{} {
							return coer98(arg98)
						}
					}(func(_1694_i2 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(811)
					})))).Cardinality()).Times(p0), 5)
					(_nw242).ArraySet1((_1691_v8).Cardinality(), 6)
					(_nw242).ArraySet1(p0, 7)
					(_nw242).ArraySet1(_dafny.IntOfInt64(-766), 8)
					(_nw242).ArraySet1(p0, 9)
					(_nw242).ArraySet1((_1690_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1690_v7).Cardinality()))).Uint32()).(_dafny.Int), 10)
					(_nw242).ArraySet1(p0, 11)
					(_nw242).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1691_v8, (_this).F15())).Merge(_1692_v9)).Cardinality(), 12)
					_1693_v10 = _nw242
					var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))
					_ = _index285
					(_1687_v5).ArraySet1(p0, (_index285).Int())
					var _1695_v11 _dafny.Sequence
					_ = _1695_v11
					_1695_v11 = _dafny.SeqOf(_1691_v8, _1691_v8)
					var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))
					_ = _index286
					(_1687_v5).ArraySet1((Companion_D9_.Create_DC22_((_this).F17(), p0, false, (_1695_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1695_v11).Cardinality()))).Uint32()).(_dafny.Set), _dafny.UnicodeSeqOfUtf8Bytes("apxjf"))).Dtor_cf38(), (_index286).Int())
					_1685_v3 = (_dafny.MultiSetFromSeq(Companion_Default___.Fm38((_this).F17(), (_this).F13(), (_this).F17(), Companion_Default___.Fm30(_dafny.IntOfInt64(596), globalState), globalState))).Update(false, Companion_Default___.Abs(p0))
					var _1696_v12 _dafny.Array
					_ = _1696_v12
					var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
					_ = _nw243
					_1696_v12 = _nw243
					var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))
					_ = _index287
					var _rhs252 _dafny.Int = p0
					_ = _rhs252
					var _rhs253 _dafny.Array = _1696_v12
					_ = _rhs253
					var _rhs254 bool = false
					_ = _rhs254
					var _lhs207 _dafny.Array = _1687_v5
					_ = _lhs207
					var _lhs208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))
					_ = _lhs208
					(_lhs207).ArraySet1(_rhs252, (_lhs208).Int())
					_1696_v12 = _rhs253
					r0 = _rhs254
					var _1697_v13 _dafny.Map
					_ = _1697_v13
					_1697_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(814))
					var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))
					_ = _index288
					(_1687_v5).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
						if (_1697_v13).Contains((_1687_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))).Int()).(_dafny.Int)) {
							return (_1697_v13).Get((_1687_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))).Int()).(_dafny.Int)).(_dafny.Int)
						}
						return (_1687_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(801), _dafny.ArrayLen((_1687_v5), 0))).Int()).(_dafny.Int)
					})(), (_1690_v7).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1690_v7).Cardinality()))).Uint32()).(_dafny.Int)), (_index288).Int())
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1698_v14 _dafny.Array
		_ = _1698_v14
		var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(26))
		_ = _nw244
		_1698_v14 = _nw244
		var _1699_v15 _dafny.Array
		_ = _1699_v15
		var _len0_39 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_39
		var _nw245 _dafny.Array
		_ = _nw245
		if _len0_39.Cmp(_dafny.Zero) == 0 {
			_nw245 = _dafny.NewArray(_len0_39)
		} else {
			var _init39 func(_dafny.Int) _dafny.Int = func(_1700_i3 _dafny.Int) _dafny.Int {
				return (_1700_i3).Plus(_dafny.IntOfInt64(683))
			}
			_ = _init39
			var _element0_39 = _init39(_dafny.Zero)
			_ = _element0_39
			_nw245 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
			(_nw245).ArraySet1(_element0_39, 0)
			var _nativeLen0_39 = (_len0_39).Int()
			_ = _nativeLen0_39
			for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
				(_nw245).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
			}
		}
		_1699_v15 = _nw245
		var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))
		_ = _index289
		(_1698_v14).ArraySet1(_1699_v15, (_index289).Int())
		var _1701_v16 _dafny.MultiSet
		_ = _1701_v16
		_1701_v16 = _dafny.MultiSetOf(_this.F12())
		var _1702_v18 _dafny.Map
		_ = _1702_v18
		_1702_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12(), true)
		var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))
		_ = _index290
		(_1698_v14).ArraySet1((func() _dafny.Array {
			if (func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter60 := _dafny.Iterate(((_1701_v16).Update(_this.F12(), Companion_Default___.Abs(p0))).Elements()); ; {
					_compr_50, _ok60 := _iter60()
					if !_ok60 {
						break
					}
					var _1703_v17 D1
					_1703_v17 = interface{}(_compr_50).(D1)
					if ((_1701_v16).Update(_this.F12(), Companion_Default___.Abs(p0))).Contains(_1703_v17) {
						_coll50.Add(_1703_v17)
					}
				}
				return _coll50.ToSet()
			}()).Equals(func() _dafny.Set {
				var _coll51 = _dafny.NewBuilder()
				_ = _coll51
				for _iter61 := _dafny.Iterate((_1702_v18).Keys().Elements()); ; {
					_compr_51, _ok61 := _iter61()
					if !_ok61 {
						break
					}
					var _1704_v19 D1
					_1704_v19 = interface{}(_compr_51).(D1)
					if (_1702_v18).Contains(_1704_v19) {
						_coll51.Add(_1704_v19)
					}
				}
				return _coll51.ToSet()
			}()) {
				return _1699_v15
			}
			return _1699_v15
		})(), (_index290).Int())
		var _1705_v20 D7
		_ = _1705_v20
		_1705_v20 = Companion_D7_.Create_DC16_(p0)
		var _1706_v21 D7
		_ = _1706_v21
		_1706_v21 = Companion_D7_.Create_DC17_(_1705_v20)
		var _pat_let_tv44 = _1705_v20
		_ = _pat_let_tv44
		var _source24 D7 = func(_pat_let52_0 D7) D7 {
			return func(_1707_dt__update__tmp_h0 D7) D7 {
				return func(_pat_let53_0 D7) D7 {
					return func(_1708_dt__update_hcf30_h0 D7) D7 {
						return Companion_D7_.Create_DC17_(_1708_dt__update_hcf30_h0)
					}(_pat_let53_0)
				}(_pat_let_tv44)
			}(_pat_let52_0)
		}(_1706_v21)
		_ = _source24
		if _source24.Is_DC16() {
			var _1709___mcc_h0 _dafny.Int = _source24.Get_().(D7_DC16).Cf29
			_ = _1709___mcc_h0
			var _1710_cf29 _dafny.Int = _1709___mcc_h0
			_ = _1710_cf29
			_1699_v15 = _1699_v15
			var _1711_v22 _dafny.Sequence
			_ = _1711_v22
			_1711_v22 = _dafny.SeqOf((_this).F13())
			var _1712_v23 _dafny.Map
			_ = _1712_v23
			_1712_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
			var _1713_v24 _dafny.Map
			_ = _1713_v24
			_1713_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_1712_v23).Contains(_1710_cf29) {
					return (_1712_v23).Get(_1710_cf29).(bool)
				}
				return (_this).F17()
			})(), p0)
			var _1714_v25 _dafny.Sequence
			_ = _1714_v25
			_1714_v25 = _dafny.SeqOf(_dafny.CodePoint('x'))
			var _1715_v26 _dafny.Map
			_ = _1715_v26
			_1715_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F17(), _1714_v25)
			var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))
			_ = _index291
			var _nwElement0_58 _dafny.Int = Companion_Default___.SafeModuloInt(_1710_cf29, (_dafny.Zero).Minus(Companion_Default___.Fm0(_1711_v22, !((_this).F13()), (_this).F17(), (func() _dafny.Int {
				if (_1713_v24).Contains((_this).F17()) {
					return (_1713_v24).Get((_this).F17()).(_dafny.Int)
				}
				return _1710_cf29
			})(), globalState)))
			_ = _nwElement0_58
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(7))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_58, 0)
			(_nw246).ArraySet1(_1710_cf29, 1)
			(_nw246).ArraySet1(((_1715_v26).Merge(_1715_v26)).Cardinality(), 2)
			(_nw246).ArraySet1(p0, 3)
			(_nw246).ArraySet1(p0, 4)
			(_nw246).ArraySet1((_dafny.Zero).Minus(((_this).F15()).Cardinality()), 5)
			(_nw246).ArraySet1(_1710_cf29, 6)
			(_1698_v14).ArraySet1(_nw246, (_index291).Int())
			var _arr10 _dafny.Array = _dafny.ArrayCastTo((_1698_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))).Int()))
			_ = _arr10
			var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_dafny.ArrayCastTo((_1698_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))).Int()))), 0))
			_ = _index292
			(_arr10).ArraySet1(_1710_cf29, (_index292).Int())
			var _arr11 _dafny.Array = _dafny.ArrayCastTo((_1698_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))).Int()))
			_ = _arr11
			var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(281), _dafny.ArrayLen((_dafny.ArrayCastTo((_1698_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(307), _dafny.ArrayLen((_1698_v14), 0))).Int()))), 0))
			_ = _index293
			(_arr11).ArraySet1(_dafny.IntOfInt64(943), (_index293).Int())
			_1714_v25 = _dafny.UnicodeSeqOfUtf8Bytes("g")
		} else if _source24.Is_DC15() {
			var _1716___mcc_h1 _dafny.Sequence = _source24.Get_().(D7_DC15).Cf28
			_ = _1716___mcc_h1
			var _1717_cf28 _dafny.Sequence = _1716___mcc_h1
			_ = _1717_cf28
			var _1718_v27 _dafny.Map
			_ = _1718_v27
			_1718_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), !((_this).F17()))
			_1718_v27 = (_1718_v27).Update((_this).F15(), (_this).F13())
			var _1719_v28 _dafny.Array
			_ = _1719_v28
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_40
			var _nw247 _dafny.Array
			_ = _nw247
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw247 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) bool = func(_1720_i4 _dafny.Int) bool {
					return (_this).F17()
				}
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw247 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw247).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw247).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1719_v28 = _nw247
			var _rhs255 bool = (_this).F17()
			_ = _rhs255
			var _rhs256 _dafny.Array = _1719_v28
			_ = _rhs256
			r0 = _rhs255
			_1719_v28 = _rhs256
			var _1721_v29 _dafny.CodePoint
			_ = _1721_v29
			_1721_v29 = _dafny.CodePoint('p')
			var _1722_v30 _dafny.Sequence
			_ = _1722_v30
			_1722_v30 = _dafny.SeqOf(_1721_v29, _dafny.CodePoint('k'), _1721_v29)
			var _1723_v31 *C0
			_ = _1723_v31
			var _nw248 *C0 = New_C0_()
			_ = _nw248
			_nw248.Ctor__((_this).F13(), (_1722_v30).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1722_v30).Cardinality()))).Uint32()).(_dafny.CodePoint), _this.F12(), true)
			_1723_v31 = _nw248
			var _1724_v32 _dafny.Sequence
			_ = _1724_v32
			_1724_v32 = _dafny.SeqOf(_1723_v31, _1723_v31)
			var _1725_v33 _dafny.Array
			_ = _1725_v33
			var _nwElement0_59 *C0 = (_1724_v32).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(2), _dafny.IntOfUint32((_1724_v32).Cardinality()))).Uint32()).(*C0)
			_ = _nwElement0_59
			var _nw249 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(14))
			_ = _nw249
			(_nw249).ArraySet1(_nwElement0_59, 0)
			(_nw249).ArraySet1(_1723_v31, 1)
			(_nw249).ArraySet1(_1723_v31, 2)
			(_nw249).ArraySet1(_1723_v31, 3)
			(_nw249).ArraySet1((func() *C0 {
				if (_this).F13() {
					return _1723_v31
				}
				return _1723_v31
			})(), 4)
			(_nw249).ArraySet1(_1723_v31, 5)
			(_nw249).ArraySet1(_1723_v31, 6)
			(_nw249).ArraySet1(_1723_v31, 7)
			(_nw249).ArraySet1(_1723_v31, 8)
			(_nw249).ArraySet1(_1723_v31, 9)
			(_nw249).ArraySet1(_1723_v31, 10)
			(_nw249).ArraySet1(_1723_v31, 11)
			(_nw249).ArraySet1(_1723_v31, 12)
			(_nw249).ArraySet1(_1723_v31, 13)
			_1725_v33 = _nw249
			var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1725_v33), 0))
			_ = _index294
			(_1725_v33).ArraySet1(_1723_v31, (_index294).Int())
			var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
			_ = _index295
			(_1699_v15).ArraySet1(p0, (_index295).Int())
			var _1726_v34 _dafny.Array
			_ = _1726_v34
			var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw250
			_1726_v34 = _nw250
			var _1727_v35 _dafny.Sequence
			_ = _1727_v35
			_1727_v35 = _dafny.SeqOf(_1726_v34)
			var _1728_v36 D8
			_ = _1728_v36
			_1728_v36 = Companion_D8_.Create_DC18_((_1727_v35).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1727_v35).Cardinality()))).Uint32()).(_dafny.Array))
			var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1725_v33), 0))
			_ = _index296
			var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
			_ = _index297
			var _rhs257 _dafny.Int = _dafny.IntOfInt64(883)
			_ = _rhs257
			var _rhs258 *C0 = _1723_v31
			_ = _rhs258
			var _rhs259 _dafny.Int = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F13() {
					return (p0).Plus(_dafny.IntOfInt64(751))
				}
				return (p0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ffp")).Cardinality()))
			})())
			_ = _rhs259
			var _rhs260 D8 = Companion_D8_.Create_DC18_(_1726_v34)
			_ = _rhs260
			var _rhs261 bool = _1723_v31.F19
			_ = _rhs261
			var _lhs209 *GlobalState = globalState
			_ = _lhs209
			var _lhs210 _dafny.Array = _1725_v33
			_ = _lhs210
			var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(528), _dafny.ArrayLen((_1725_v33), 0))
			_ = _lhs211
			var _lhs212 _dafny.Array = _1699_v15
			_ = _lhs212
			var _lhs213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
			_ = _lhs213
			var _lhs214 *C0 = _1723_v31
			_ = _lhs214
			_lhs209.F7 = _rhs257
			(_lhs210).ArraySet1(_rhs258, (_lhs211).Int())
			(_lhs212).ArraySet1(_rhs259, (_lhs213).Int())
			_1728_v36 = _rhs260
			_lhs214.F19 = _rhs261
			if _1723_v31.F19 {
				var _1729_v37 _dafny.Sequence
				_ = _1729_v37
				_1729_v37 = _dafny.SeqOf(_1723_v31.F19, _1723_v31.F19)
				var _1730_v38 D3
				_ = _1730_v38
				_1730_v38 = Companion_D3_.Create_DC8_((_dafny.Zero).Minus(Companion_Default___.Fm0(_1729_v37, _1723_v31.F19, !(_1723_v31.F19), (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), globalState)))
				var _1731_v39 _dafny.Set
				_ = _1731_v39
				_1731_v39 = _dafny.SetOf((_1723_v31).F20(), (_1723_v31).F20())
				var _1732_v40 _dafny.Map
				_ = _1732_v40
				_1732_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1721_v29, Companion_Default___.Fm17(globalState))
				var _pat_let_tv45 = _1699_v15
				_ = _pat_let_tv45
				var _pat_let_tv46 = _1699_v15
				_ = _pat_let_tv46
				var _pat_let_tv47 = _1699_v15
				_ = _pat_let_tv47
				var _pat_let_tv48 = _1699_v15
				_ = _pat_let_tv48
				var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index298
				var _rhs262 D3 = Companion_Default___.Fm49(_this.F16(), _1731_v39, globalState)
				_ = _rhs262
				var _rhs263 _dafny.Int = (_1723_v31).Fm16((_this).Fm12(_1732_v40, func(_pat_let54_0 D1) D1 {
					return func(_1733_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let55_0 _dafny.Int) D1 {
							return func(_1734_dt__update_hcf3_h0 _dafny.Int) D1 {
								return func(_pat_let56_0 _dafny.Int) D1 {
									return func(_1735_dt__update_hcf2_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC2_(_1735_dt__update_hcf2_h0, _1734_dt__update_hcf3_h0, (_1733_dt__update__tmp_h1).Dtor_cf4())
									}(_pat_let56_0)
								}((_pat_let_tv48).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_pat_let_tv47), 0))).Int()).(_dafny.Int))
							}(_pat_let55_0)
						}((_pat_let_tv46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_pat_let_tv45), 0))).Int()).(_dafny.Int))
					}(_pat_let54_0)
				}(_this.F12()), globalState), p0, globalState)
				_ = _rhs263
				var _rhs264 bool = _1723_v31.F19
				_ = _rhs264
				var _lhs215 _dafny.Array = _1699_v15
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
				_ = _lhs216
				var _lhs217 *C0 = _1723_v31
				_ = _lhs217
				_1730_v38 = _rhs262
				(_lhs215).ArraySet1(_rhs263, (_lhs216).Int())
				_lhs217.F19 = _rhs264
				var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index299
				(_1699_v15).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (func() bool {
						if _1723_v31.F19 {
							return Companion_Default___.Fm17(globalState)
						}
						return _1723_v31.F19
					})() {
						return (_1717_cf28).Select((Companion_Default___.SafeIndex((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1717_cf28).Cardinality()))).Uint32()).(_dafny.Int)
					}
					return Companion_Default___.Fm0(_dafny.SeqOf(_1723_v31.F19, false, (_1729_v37).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1729_v37).Cardinality()), _dafny.IntOfUint32((_1729_v37).Cardinality()))).Uint32()).(bool)), (_this).F17(), (_this).F17(), (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), globalState)
				})())), (_index299).Int())
				var _1736_v41 _dafny.Map
				_ = _1736_v41
				_1736_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1729_v37)
				_1736_v41 = (_1736_v41).Update((func() _dafny.Set {
					var _coll52 = _dafny.NewBuilder()
					_ = _coll52
					for _iter62 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(736), _dafny.IntOfInt64(816))); ; {
						_compr_52, _ok62 := _iter62()
						if !_ok62 {
							break
						}
						var _1737_v42 _dafny.Int
						_1737_v42 = interface{}(_compr_52).(_dafny.Int)
						if ((_dafny.IntOfInt64(736)).Cmp(_1737_v42) <= 0) && ((_1737_v42).Cmp(_dafny.IntOfInt64(816)) < 0) {
							_coll52.Add(Companion_Default___.SafeModuloInt(_1737_v42, _dafny.IntOfUint32((_dafny.SeqOf(_1723_v31.F19)).Cardinality())))
						}
					}
					return _coll52.ToSet()
				}()).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_1729_v37, _1729_v37))
				var _1738_v43 _dafny.MultiSet
				_ = _1738_v43
				_1738_v43 = _dafny.MultiSetOf((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int))
				var _1739_v44 D6
				_ = _1739_v44
				_1739_v44 = Companion_D6_.Create_DC11_(_1738_v43)
				var _1740_v45 D6
				_ = _1740_v45
				_1740_v45 = Companion_D6_.Create_DC14_(_1739_v44)
				var _1741_v46 _dafny.Array
				_ = _1741_v46
				var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(25))
				_ = _nw251
				_1741_v46 = _nw251
				var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1741_v46), 0))
				_ = _index300
				(_1741_v46).ArraySet1(Companion_Default___.Fm50(globalState), (_index300).Int())
				var _1742_v47 _dafny.MultiSet
				_ = _1742_v47
				_1742_v47 = _dafny.MultiSetOf((func() bool {
					if !(_1723_v31.F19) {
						return _1723_v31.F19
					}
					return (func() bool {
						if (_1732_v40).Contains(_1721_v29) {
							return (_1732_v40).Get(_1721_v29).(bool)
						}
						return (_this).F17()
					})()
				})(), false, false, !(((_this).F17()) && ((_this).F17())), _dafny.Companion_Sequence_.IsProperPrefixOf(_1722_v30, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ido"), (Companion_Default___.SafeIndex((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ido")).Cardinality()))).Uint32(), (_1723_v31).F20())))
				var _1743_v48 _dafny.Map
				_ = _1743_v48
				_1743_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
				var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1741_v46), 0))
				_ = _index301
				var _rhs265 _dafny.CodePoint = (_1723_v31).F20()
				_ = _rhs265
				var _rhs266 D6 = Companion_D6_.Create_DC14_(_1739_v44)
				_ = _rhs266
				var _rhs267 _dafny.Map = ((_1743_v48).Update(_dafny.IntOfInt64(376), _1723_v31.F19)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), _1723_v31.F19))
				_ = _rhs267
				var _rhs268 _dafny.MultiSet = (func() _dafny.MultiSet {
					if ((_this).F15()).IsProperSubsetOf((_this).F15()) {
						return _1742_v47
					}
					return _1742_v47
				})()
				_ = _rhs268
				var _lhs218 _dafny.Array = _1741_v46
				_ = _lhs218
				var _lhs219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(303), _dafny.ArrayLen((_1741_v46), 0))
				_ = _lhs219
				_1721_v29 = _rhs265
				_1740_v45 = _rhs266
				(_lhs218).ArraySet1(_rhs267, (_lhs219).Int())
				_1742_v47 = _rhs268
				var _rhs269 _dafny.Int = (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int)
				_ = _rhs269
				var _rhs270 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1721_v29, (_this).F13())
				_ = _rhs270
				var _rhs271 D6 = (func() D6 {
					if !(false) {
						return _1740_v45
					}
					return _1740_v45
				})()
				_ = _rhs271
				var _rhs272 bool = (_1729_v37).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_this).F13() {
						return (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int)
					}
					return (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_1729_v37).Cardinality()))).Uint32()).(bool)
				_ = _rhs272
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				_lhs220.F7 = _rhs269
				_1732_v40 = _rhs270
				_1740_v45 = _rhs271
				r0 = _rhs272
			} else {
				var _1744_v49 *C1
				_ = _1744_v49
				var _nw252 *C1 = New_C1_()
				_ = _nw252
				_nw252.Ctor__((_1723_v31).F20(), _this.F12(), !((_this).F13()))
				_1744_v49 = _nw252
				var _1745_v50 _dafny.Map
				_ = _1745_v50
				_1745_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(182), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int))
				_1745_v50 = _1745_v50
				(_1723_v31).F19 = _1723_v31.F19
				var _1746_v51 _dafny.Map
				_ = _1746_v51
				_1746_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1723_v31.F19, p0)
				_1746_v51 = (_1746_v51).Update((_this).F13(), p0)
				(_1723_v31).F19 = !((_this).F17()) || (_dafny.Companion_Sequence_.Equal(_1722_v30, _1722_v30))
			}
		} else {
			var _1747___mcc_h2 D7 = _source24.Get_().(D7_DC17).Cf30
			_ = _1747___mcc_h2
			var _1748_cf30 D7 = _1747___mcc_h2
			_ = _1748_cf30
			var _index302 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index302
			((_this).F14()).ArraySet1((_this).F17(), (_index302).Int())
			var _index303 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index303
			((_this).F14()).ArraySet1((_this).F17(), (_index303).Int())
			if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
				var _1749_v52 _dafny.Array
				_ = _1749_v52
				var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(9))
				_ = _nw253
				_1749_v52 = _nw253
				var _1750_v53 _dafny.CodePoint
				_ = _1750_v53
				_1750_v53 = _dafny.CodePoint('l')
				var _index304 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1749_v52), 0))
				_ = _index304
				(_1749_v52).ArraySet1CodePoint(_1750_v53, (_index304).Int())
				var _1751_v54 _dafny.Sequence
				_ = _1751_v54
				_1751_v54 = _dafny.SeqOf((_this).F17())
				var _index305 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1749_v52), 0))
				_ = _index305
				var _rhs273 bool = (p0).Cmp(Companion_Default___.Fm0(_1751_v54, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), p0, globalState)) < 0
				_ = _rhs273
				var _rhs274 _dafny.CodePoint = _1750_v53
				_ = _rhs274
				var _rhs275 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, p0)
				_ = _rhs275
				var _lhs221 _dafny.Array = _1749_v52
				_ = _lhs221
				var _lhs222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.ArrayLen((_1749_v52), 0))
				_ = _lhs222
				var _lhs223 *GlobalState = globalState
				_ = _lhs223
				r0 = _rhs273
				(_lhs221).ArraySet1CodePoint(_rhs274, (_lhs222).Int())
				_lhs223.F7 = _rhs275
				(globalState).F7 = (Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(832), p0)).Times(p0)
				var _1752_v55 _dafny.Map
				_ = _1752_v55
				_1752_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), (_this).F17())
				var _1753_v56 _dafny.Sequence
				_ = _1753_v56
				_1753_v56 = _dafny.UnicodeSeqOfUtf8Bytes("uwc")
				var _1754_v57 _dafny.Array
				_ = _1754_v57
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
				_ = _nw254
				_1754_v57 = _nw254
				var _pat_let_tv49 = _1753_v56
				_ = _pat_let_tv49
				var _1755_v58 *C2
				_ = _1755_v58
				var _nw255 *C2 = New_C2_()
				_ = _nw255
				_nw255.Ctor__(p0, (func() bool {
					if (_1752_v55).Contains(_dafny.IntOfInt64(-780)) {
						return (_1752_v55).Get(_dafny.IntOfInt64(-780)).(bool)
					}
					return (_this).F13()
				})(), func(_pat_let57_0 D1) D1 {
					return func(_1756_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let58_0 _dafny.Sequence) D1 {
							return func(_1757_dt__update_hcf1_h0 _dafny.Sequence) D1 {
								return Companion_D1_.Create_DC1_(_1757_dt__update_hcf1_h0)
							}(_pat_let58_0)
						}(_pat_let_tv49)
					}(_pat_let57_0)
				}(_this.F16()), (_this).F17(), _1754_v57, Companion_Default___.Fm47(!((_this).F13()), globalState), _this.F12(), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				_1755_v58 = _nw255
				var _1758_v59 *C3
				_ = _1758_v59
				var _nw256 *C3 = New_C3_()
				_ = _nw256
				_nw256.Ctor__((_this).F17(), (func() D1 {
					if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
						return _this.F16()
					}
					return Companion_D1_.Create_DC1_(_1753_v56)
				})(), (_this).F13(), _1754_v57, ((_this).F15()).Difference((_this).F15()), Companion_D1_.Create_DC2_(p0, (_1755_v58).F21(), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)), (_this).F13())
				_1758_v59 = _nw256
				var _index306 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index306
				((_this).F14()).ArraySet1((p0).Cmp(p0) != 0, (_index306).Int())
			} else {
				var _index307 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index307
				((_this).F14()).ArraySet1(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_index307).Int())
				var _1759_v60 _dafny.MultiSet
				_ = _1759_v60
				_1759_v60 = _dafny.MultiSetOf((_this).F13(), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				_1759_v60 = Companion_Default___.Fm40(globalState)
				var _1760_v61 _dafny.Map
				_ = _1760_v61
				_1760_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(5), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				var _1761_v62 _dafny.Sequence
				_ = _1761_v62
				_1761_v62 = _dafny.SeqOf((_this).F13())
				_1760_v61 = (_1760_v61).Update(p0, (_1761_v62).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1761_v62).Cardinality()))).Uint32()).(bool))
				var _1762_v63 _dafny.MultiSet
				_ = _1762_v63
				_1762_v63 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0), (_dafny.Zero).Minus(p0))
				var _1763_v64 D6
				_ = _1763_v64
				_1763_v64 = Companion_D6_.Create_DC11_(_1762_v63)
				var _1764_v65 D6
				_ = _1764_v65
				_1764_v65 = Companion_D6_.Create_DC14_(Companion_D6_.Create_DC11_(((_1763_v64).Dtor_cf17()).Update(p0, Companion_Default___.Abs(Companion_Default___.Fm0(_dafny.SeqOf(!((_this).F17())), (_this).F13(), false, p0, globalState)))))
				var _1765_v66 _dafny.MultiSet
				_ = _1765_v66
				var _out12 _dafny.MultiSet
				_ = _out12
				_out12 = Companion_Default___.M0(Companion_Default___.Fm13(Companion_Default___.Fm32(_1764_v65, globalState), globalState), (_1761_v62).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(156), _dafny.IntOfUint32((_1761_v62).Cardinality()))).Uint32()).(bool), globalState)
				_1765_v66 = _out12
				var _1766_v67 _dafny.Map
				_ = _1766_v67
				_1766_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1761_v62, !((_this).F13()))
				var _1767_v68 _dafny.Sequence
				_ = _1767_v68
				_1767_v68 = _dafny.UnicodeSeqOfUtf8Bytes("u")
				var _1768_v69 _dafny.CodePoint
				_ = _1768_v69
				_1768_v69 = _dafny.CodePoint('k')
				var _1769_v70 _dafny.Map
				_ = _1769_v70
				_1769_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1768_v69, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				var _1770_v71 _dafny.Map
				_ = _1770_v71
				_1770_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1767_v68).Cardinality()), _1768_v69)
				var _1771_v72 _dafny.MultiSet
				_ = _1771_v72
				_1771_v72 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Update(_1767_v68, (Companion_Default___.SafeIndex(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.Companion_Sequence_.Update(_1761_v62, (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_1761_v62, (_this).F17(), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), p0, globalState), _dafny.IntOfUint32((_1761_v62).Cardinality()))).Uint32(), (_this).F17()), (_this).F17(), (_this).F13(), p0, globalState), p0)).Update(p0, p0)).Cardinality(), _dafny.IntOfUint32((_1767_v68).Cardinality()))).Uint32(), Companion_Default___.Fm33((func() bool {
					if (_1769_v70).Contains((func() _dafny.CodePoint {
						if (_1770_v71).Contains(_dafny.IntOfUint32((_1767_v68).Cardinality())) {
							return (_1770_v71).Get(_dafny.IntOfUint32((_1767_v68).Cardinality())).(_dafny.CodePoint)
						}
						return _1768_v69
					})()) {
						return (_1769_v70).Get((func() _dafny.CodePoint {
							if (_1770_v71).Contains(_dafny.IntOfUint32((_1767_v68).Cardinality())) {
								return (_1770_v71).Get(_dafny.IntOfUint32((_1767_v68).Cardinality())).(_dafny.CodePoint)
							}
							return _1768_v69
						})()).(bool)
					}
					return Companion_Default___.Fm17(globalState)
				})(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jmgntc")).Cardinality()), globalState)), _1767_v68, _1767_v68)
				var _1772_v73 D12
				_ = _1772_v73
				_1772_v73 = Companion_D12_.Create_DC29_(_1766_v67)
				var _index308 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index308
				var _rhs276 bool = ((func() _dafny.Int {
					if false {
						return _dafny.IntOfInt64(148)
					}
					return p0
				})()).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_1771_v72).Contains(_dafny.UnicodeSeqOfUtf8Bytes("qxcd")) {
						return (_1771_v72).Multiplicity(_dafny.UnicodeSeqOfUtf8Bytes("qxcd"))
					}
					return (_dafny.Zero).Minus(p0)
				})()))) == 0
				_ = _rhs276
				var _rhs277 _dafny.Int = p0
				_ = _rhs277
				var _rhs278 _dafny.Map = (_1772_v73).Dtor_cf57()
				_ = _rhs278
				var _lhs224 _dafny.Array = (_this).F14()
				_ = _lhs224
				var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs225
				var _lhs226 *GlobalState = globalState
				_ = _lhs226
				(_lhs224).ArraySet1(_rhs276, (_lhs225).Int())
				_lhs226.F7 = _rhs277
				_1766_v67 = _rhs278
			}
			var _1773_v74 _dafny.Map
			_ = _1773_v74
			_1773_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
			var _1774_v75 _dafny.Sequence
			_ = _1774_v75
			_1774_v75 = _dafny.SeqOf(true)
			var _index309 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index309
			var _index310 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _index310
			var _rhs279 _dafny.Int = _dafny.IntOfInt64(243)
			_ = _rhs279
			var _rhs280 bool = true
			_ = _rhs280
			var _rhs281 bool = (func() bool {
				if (_1773_v74).Contains(p0) {
					return (_1773_v74).Get(p0).(bool)
				}
				return _dafny.Companion_Sequence_.IsPrefixOf(_1774_v75, _1774_v75)
			})()
			_ = _rhs281
			var _lhs227 *GlobalState = globalState
			_ = _lhs227
			var _lhs228 _dafny.Array = (_this).F14()
			_ = _lhs228
			var _lhs229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs229
			var _lhs230 _dafny.Array = (_this).F14()
			_ = _lhs230
			var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
			_ = _lhs231
			_lhs227.F7 = _rhs279
			(_lhs228).ArraySet1(_rhs280, (_lhs229).Int())
			(_lhs230).ArraySet1(_rhs281, (_lhs231).Int())
			var _source25 D9 = Companion_Default___.Fm51(globalState)
			_ = _source25
			if _source25.Is_DC21() {
				var _index311 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index311
				(_1699_v15).ArraySet1(p0, (_index311).Int())
				var _index312 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(384), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index312
				(_1699_v15).ArraySet1(p0, (_index312).Int())
				var _1775_v76 _dafny.CodePoint
				_ = _1775_v76
				_1775_v76 = _dafny.CodePoint('l')
				var _1776_v77 _dafny.MultiSet
				_ = _1776_v77
				var _out13 _dafny.MultiSet
				_ = _out13
				_out13 = Companion_Default___.M0(_1775_v76, false, globalState)
				_1776_v77 = _out13
				(globalState).F7 = (_dafny.Zero).Minus(p0)
				var _index313 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index313
				((_this).F14()).ArraySet1(!((_this).F17()), (_index313).Int())
			} else if _source25.Is_DC22() {
				var _1777___mcc_h3 bool = _source25.Get_().(D9_DC22).Cf37
				_ = _1777___mcc_h3
				var _1778___mcc_h4 _dafny.Int = _source25.Get_().(D9_DC22).Cf38
				_ = _1778___mcc_h4
				var _1779___mcc_h5 bool = _source25.Get_().(D9_DC22).Cf39
				_ = _1779___mcc_h5
				var _1780___mcc_h6 _dafny.Set = _source25.Get_().(D9_DC22).Cf40
				_ = _1780___mcc_h6
				var _1781___mcc_h7 _dafny.Sequence = _source25.Get_().(D9_DC22).Cf41
				_ = _1781___mcc_h7
				var _1782_cf41 _dafny.Sequence = _1781___mcc_h7
				_ = _1782_cf41
				var _1783_cf40 _dafny.Set = _1780___mcc_h6
				_ = _1783_cf40
				var _1784_cf39 bool = _1779___mcc_h5
				_ = _1784_cf39
				var _1785_cf38 _dafny.Int = _1778___mcc_h4
				_ = _1785_cf38
				var _1786_cf37 bool = _1777___mcc_h3
				_ = _1786_cf37
				var _index314 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index314
				(_1699_v15).ArraySet1((_dafny.Zero).Minus(_1785_cf38), (_index314).Int())
				var _index315 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index315
				(_1699_v15).ArraySet1(_dafny.IntOfInt64(879), (_index315).Int())
				(globalState).F2 = _dafny.Companion_Sequence_.Concatenate(_1774_v75, _dafny.SeqOf(_1786_cf37, (_this).F17(), Companion_Default___.Fm30((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(141), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), globalState)))
				var _1787_v78 _dafny.Array
				_ = _1787_v78
				var _nw257 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
				_ = _nw257
				_1787_v78 = _nw257
				var _rhs282 _dafny.Array = _1787_v78
				_ = _rhs282
				var _rhs283 _dafny.Int = p0
				_ = _rhs283
				var _lhs232 *GlobalState = globalState
				_ = _lhs232
				_1787_v78 = _rhs282
				_lhs232.F7 = _rhs283
				r0 = true
			} else if _source25.Is_DC23() {
				var _1788___mcc_h8 _dafny.Int = _source25.Get_().(D9_DC23).Cf42
				_ = _1788___mcc_h8
				var _1789___mcc_h9 _dafny.Int = _source25.Get_().(D9_DC23).Cf43
				_ = _1789___mcc_h9
				var _1790___mcc_h10 _dafny.Int = _source25.Get_().(D9_DC23).Cf44
				_ = _1790___mcc_h10
				var _1791___mcc_h11 _dafny.Int = _source25.Get_().(D9_DC23).Cf45
				_ = _1791___mcc_h11
				var _1792___mcc_h12 _dafny.Int = _source25.Get_().(D9_DC23).Cf46
				_ = _1792___mcc_h12
				var _1793_cf46 _dafny.Int = _1792___mcc_h12
				_ = _1793_cf46
				var _1794_cf45 _dafny.Int = _1791___mcc_h11
				_ = _1794_cf45
				var _1795_cf44 _dafny.Int = _1790___mcc_h10
				_ = _1795_cf44
				var _1796_cf43 _dafny.Int = _1789___mcc_h9
				_ = _1796_cf43
				var _1797_cf42 _dafny.Int = _1788___mcc_h8
				_ = _1797_cf42
				var _index316 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index316
				((_this).F14()).ArraySet1((func() bool {
					if (_this).F13() {
						return true
					}
					return (_this).F17()
				})(), (_index316).Int())
				var _1798_v79 D6
				_ = _1798_v79
				_1798_v79 = Companion_D6_.Create_DC13_(!((_this).F17()), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), _dafny.IntOfInt64(556), _1793_cf46, (_this).F17())
				_1798_v79 = _1798_v79
				var _index317 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index317
				((_this).F14()).ArraySet1(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_index317).Int())
				var _1799_v80 _dafny.Array
				_ = _1799_v80
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_41
				var _nw258 _dafny.Array
				_ = _nw258
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw258 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = func(_1800_i5 _dafny.Int) bool {
						return ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
					}
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw258 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw258).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw258).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1799_v80 = _nw258
				var _1801_v81 _dafny.MultiSet
				_ = _1801_v81
				_1801_v81 = _dafny.MultiSetOf(_1799_v80)
				var _1802_v82 _dafny.MultiSet
				_ = _1802_v82
				_1802_v82 = _1801_v81
				var _1803_v83 _dafny.Map
				_ = _1803_v83
				_1803_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1801_v81, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("f"), _dafny.UnicodeSeqOfUtf8Bytes("uothuadbe"))).Cardinality()))
				var _1804_v84 _dafny.CodePoint
				_ = _1804_v84
				_1804_v84 = _dafny.CodePoint('j')
				var _1805_v85 *C1
				_ = _1805_v85
				var _nw259 *C1 = New_C1_()
				_ = _nw259
				_nw259.Ctor__(_1804_v84, _this.F12(), (_this).F17())
				_1805_v85 = _nw259
				var _1806_v86 _dafny.Sequence
				_ = _1806_v86
				_1806_v86 = _dafny.SeqOf(_1805_v85)
				_1803_v83 = (_1803_v83).Update(_1802_v82, _dafny.IntOfUint32((_1806_v86).Cardinality()))
			} else {
				var _1807___mcc_h13 _dafny.Set = _source25.Get_().(D9_DC20).Cf36
				_ = _1807___mcc_h13
				var _1808_cf36 _dafny.Set = _1807___mcc_h13
				_ = _1808_cf36
				var _index318 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index318
				((_this).F14()).ArraySet1((((_this).F15()).Union(_dafny.SetOf(p0))).IsDisjointFrom((_this).F15()), (_index318).Int())
				var _index319 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index319
				(_1699_v15).ArraySet1(p0, (_index319).Int())
				var _index320 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index320
				(_1699_v15).ArraySet1(p0, (_index320).Int())
				var _1809_v87 _dafny.MultiSet
				_ = _1809_v87
				_1809_v87 = _dafny.MultiSetOf(p0)
				var _1810_v88 _dafny.MultiSet
				_ = _1810_v88
				_1810_v88 = _dafny.MultiSetOf((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if (_1809_v87).Contains((_dafny.Zero).Minus(p0)) {
						return (_1809_v87).Multiplicity((_dafny.Zero).Minus(p0))
					}
					return p0
				})())
				var _1811_v89 _dafny.Set
				_ = _1811_v89
				_1811_v89 = _dafny.SetOf((_1809_v87).IsProperSubsetOf(_1809_v87))
				var _1812_v90 _dafny.Map
				_ = _1812_v90
				_1812_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf((_this).F13()))
				var _1813_v91 _dafny.Sequence
				_ = _1813_v91
				_1813_v91 = _dafny.UnicodeSeqOfUtf8Bytes("kxckk")
				var _1814_v92 _dafny.CodePoint
				_ = _1814_v92
				_1814_v92 = _dafny.CodePoint('a')
				var _1815_v93 _dafny.Sequence
				_ = _1815_v93
				_1815_v93 = _dafny.SeqOf(p0)
				var _rhs284 bool = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				_ = _rhs284
				var _rhs285 _dafny.Set = (func() _dafny.Set {
					if (_1812_v90).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1813_v91, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1813_v91).Cardinality()))).Uint32(), _1814_v92)).Cardinality())) {
						return (_1812_v90).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1813_v91, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1813_v91).Cardinality()))).Uint32(), _1814_v92)).Cardinality())).(_dafny.Set)
					}
					return _dafny.SetOf(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
				})()
				_ = _rhs285
				var _rhs286 _dafny.MultiSet = ((_dafny.MultiSetFromSeq(_1815_v93)).Intersection((_1810_v88).Update(_dafny.IntOfInt64(901), Companion_Default___.Abs((_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int))))).Difference(_1809_v87)
				_ = _rhs286
				r0 = _rhs284
				_1811_v89 = _rhs285
				_1810_v88 = _rhs286
				var _index321 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))
				_ = _index321
				(_1699_v15).ArraySet1(Companion_Default___.Fm0(_dafny.SeqOf((_this).F13()), ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), true, (_1699_v15).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_1699_v15), 0))).Int()).(_dafny.Int), globalState), (_index321).Int())
			}
		}
		r0 = ((_this).F13()) || ((_this).F17())
		var _index322 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index322
		((_this).F14()).ArraySet1(!((_this).F17()), (_index322).Int())
		var _index323 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index323
		((_this).F14()).ArraySet1((_this).F17(), (_index323).Int())
		var _index324 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen(((_this).F14()), 0))
		_ = _index324
		((_this).F14()).ArraySet1(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (_index324).Int())
		r0 = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
		return r0
	}
}
func (_this *C6) M8(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1816_v0 _dafny.Sequence
		_ = _1816_v0
		_1816_v0 = _dafny.SeqOf((_this).F17())
		(globalState).F7 = Companion_Default___.Fm0(_1816_v0, (_this).F17(), (_this).F17(), p0, globalState)
		var _1817_i0 _dafny.Int
		_ = _1817_i0
		_1817_i0 = _dafny.Zero
		{
			for !((_this).F17()) {
				{
					if (_1817_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1817_i0 = (_1817_i0).Plus(_dafny.One)
					var _1818_v1 _dafny.MultiSet
					_ = _1818_v1
					_1818_v1 = _dafny.MultiSetOf((_this).F17())
					r0 = ((_1818_v1).Union(_1818_v1)).IsSubsetOf(_1818_v1)
					(globalState).F7 = p0
					var _1819_v2 _dafny.Array
					_ = _1819_v2
					var _nw260 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
					_ = _nw260
					_1819_v2 = _nw260
					var _1820_v3 _dafny.Array
					_ = _1820_v3
					var _nw261 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
					_ = _nw261
					_1820_v3 = _nw261
					var _index325 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1819_v2), 0))
					_ = _index325
					(_1819_v2).ArraySet1(_1820_v3, (_index325).Int())
					var _1821_v4 _dafny.Array
					_ = _1821_v4
					var _nw262 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(19))
					_ = _nw262
					_1821_v4 = _nw262
					var _1822_v5 _dafny.CodePoint
					_ = _1822_v5
					_1822_v5 = _dafny.CodePoint('h')
					var _index326 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1821_v4), 0))
					_ = _index326
					(_1821_v4).ArraySet1((_this).Fm9(_1822_v5, globalState), (_index326).Int())
					var _1823_v6 _dafny.Map
					_ = _1823_v6
					_1823_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
					var _1824_v7 _dafny.Map
					_ = _1824_v7
					_1824_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm17(globalState), (_dafny.MultiSetOf(_1823_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F13()))).IsSubsetOf(_dafny.MultiSetOf(_1823_v6)))
					var _1825_v8 D2
					_ = _1825_v8
					_1825_v8 = Companion_D2_.Create_DC4_(_1821_v4, (_this).F17(), p0, (_this).F13(), (_this).F14())
					var _index327 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1819_v2), 0))
					_ = _index327
					var _index328 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1821_v4), 0))
					_ = _index328
					var _rhs287 bool = (_this).F17()
					_ = _rhs287
					var _rhs288 bool = (func() bool {
						if (_1824_v7).Contains((_this).F17()) {
							return (_1824_v7).Get((_this).F17()).(bool)
						}
						return (_this).F13()
					})()
					_ = _rhs288
					var _rhs289 _dafny.Array = _1820_v3
					_ = _rhs289
					var _rhs290 _dafny.Int = (_1825_v8).Dtor_cf8()
					_ = _rhs290
					var _rhs291 _dafny.Int = (p0).Times(_dafny.IntOfInt64(2))
					_ = _rhs291
					var _lhs233 _dafny.Array = _1819_v2
					_ = _lhs233
					var _lhs234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_1819_v2), 0))
					_ = _lhs234
					var _lhs235 _dafny.Array = _1821_v4
					_ = _lhs235
					var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1821_v4), 0))
					_ = _lhs236
					var _lhs237 *GlobalState = globalState
					_ = _lhs237
					r0 = _rhs287
					r0 = _rhs288
					(_lhs233).ArraySet1(_rhs289, (_lhs234).Int())
					(_lhs235).ArraySet1(_rhs290, (_lhs236).Int())
					_lhs237.F7 = _rhs291
					var _index329 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1821_v4), 0))
					_ = _index329
					(_1821_v4).ArraySet1((_1821_v4).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(888), _dafny.ArrayLen((_1821_v4), 0))).Int()).(_dafny.Int), (_index329).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1826_v9 _dafny.CodePoint
		_ = _1826_v9
		_1826_v9 = _dafny.CodePoint('e')
		var _1827_v10 _dafny.Map
		_ = _1827_v10
		_1827_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1826_v9, _dafny.CodePoint('i'))
		_1827_v10 = (_1827_v10).Update(_1826_v9, _1826_v9)
		var _1828_v11 _dafny.Array
		_ = _1828_v11
		var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
		_ = _nw263
		_1828_v11 = _nw263
		var _1829_v12 _dafny.Array
		_ = _1829_v12
		_1829_v12 = _1828_v11
		var _1830_v13 _dafny.Map
		_ = _1830_v13
		_1830_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1829_v12), !(false))
		var _1831_v14 _dafny.Set
		_ = _1831_v14
		_1831_v14 = _dafny.SetOf((_this).F17())
		_1830_v13 = (_1830_v13).Update(_1828_v11, (_1831_v14).Equals(_1831_v14))
		(globalState).F7 = (_this).Fm9(_1826_v9, globalState)
		var _hi12 _dafny.Int = p0
		_ = _hi12
		for _1832_i1 := p0; _1832_i1.Cmp(_hi12) < 0; _1832_i1 = _1832_i1.Plus(_dafny.One) {
			var _1833_v15 _dafny.Map
			_ = _1833_v15
			_1833_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm11(globalState), _1832_i1)
			_1833_v15 = (_1833_v15).Update((_this).F13(), _1832_i1)
			(globalState).F7 = _1832_i1
			var _1834_v16 D7
			_ = _1834_v16
			_1834_v16 = Companion_D7_.Create_DC16_(_1832_i1)
			var _1835_v17 D9
			_ = _1835_v17
			_1835_v17 = Companion_D9_.Create_DC23_(Companion_Default___.SafeDivisionInt(p0, p0), p0, ((_1834_v16).Dtor_cf29()).Minus((_dafny.SetOf(p0)).Cardinality()), _1832_i1, Companion_Default___.SafeDivisionInt(Companion_Default___.Fm0(_1816_v0, (_this).F17(), !(false), p0, globalState), _1832_i1))
			var _1836_v18 _dafny.Array
			_ = _1836_v18
			var _nw264 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw264
			_1836_v18 = _nw264
			var _1837_v19 _dafny.Array
			_ = _1837_v19
			var _nwElement0_60 _dafny.CodePoint = _1826_v9
			_ = _nwElement0_60
			var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(27))
			_ = _nw265
			(_nw265).ArraySet1CodePoint(_nwElement0_60, 0)
			(_nw265).ArraySet1CodePoint(_1826_v9, 1)
			(_nw265).ArraySet1CodePoint(_1826_v9, 2)
			(_nw265).ArraySet1CodePoint(_dafny.CodePoint('f'), 3)
			(_nw265).ArraySet1CodePoint(_1826_v9, 4)
			(_nw265).ArraySet1CodePoint(_dafny.CodePoint('j'), 5)
			(_nw265).ArraySet1CodePoint(_1826_v9, 6)
			(_nw265).ArraySet1CodePoint(_1826_v9, 7)
			(_nw265).ArraySet1CodePoint(_1826_v9, 8)
			(_nw265).ArraySet1CodePoint(_1826_v9, 9)
			(_nw265).ArraySet1CodePoint(_1826_v9, 10)
			(_nw265).ArraySet1CodePoint(_1826_v9, 11)
			(_nw265).ArraySet1CodePoint(_1826_v9, 12)
			(_nw265).ArraySet1CodePoint(_1826_v9, 13)
			(_nw265).ArraySet1CodePoint(_1826_v9, 14)
			(_nw265).ArraySet1CodePoint(_1826_v9, 15)
			(_nw265).ArraySet1CodePoint(_1826_v9, 16)
			(_nw265).ArraySet1CodePoint(_1826_v9, 17)
			(_nw265).ArraySet1CodePoint(_1826_v9, 18)
			(_nw265).ArraySet1CodePoint(_1826_v9, 19)
			(_nw265).ArraySet1CodePoint(_dafny.CodePoint('b'), 20)
			(_nw265).ArraySet1CodePoint(_1826_v9, 21)
			(_nw265).ArraySet1CodePoint(_1826_v9, 22)
			(_nw265).ArraySet1CodePoint(_1826_v9, 23)
			(_nw265).ArraySet1CodePoint(_1826_v9, 24)
			(_nw265).ArraySet1CodePoint(_1826_v9, 25)
			(_nw265).ArraySet1CodePoint(_1826_v9, 26)
			_1837_v19 = _nw265
			var _index330 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1836_v18), 0))
			_ = _index330
			(_1836_v18).ArraySet1((Companion_D13_.Create_DC33_(_1837_v19)).Dtor_cf70(), (_index330).Int())
			var _1838_v20 _dafny.MultiSet
			_ = _1838_v20
			_1838_v20 = _dafny.MultiSetOf(true, (_this).F17(), (_this).F13())
			var _1839_v21 _dafny.Set
			_ = _1839_v21
			_1839_v21 = _dafny.SetOf(_1838_v20, _1838_v20)
			var _1840_v22 D6
			_ = _1840_v22
			_1840_v22 = Companion_D6_.Create_DC11_((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs(_1832_i1)))
			var _1841_v23 D6
			_ = _1841_v23
			_1841_v23 = Companion_D6_.Create_DC14_(_1840_v22)
			var _index331 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1836_v18), 0))
			_ = _index331
			var _rhs292 D9 = _1835_v17
			_ = _rhs292
			var _rhs293 _dafny.Array = _1837_v19
			_ = _rhs293
			var _rhs294 _dafny.CodePoint = _1826_v9
			_ = _rhs294
			var _rhs295 _dafny.Set = (_1839_v21).Difference(_1839_v21)
			_ = _rhs295
			var _rhs296 bool = Companion_Default___.Fm30((func() _dafny.Int {
				if (_this).F17() {
					return _dafny.IntOfUint32((Companion_Default___.Fm32(Companion_D6_.Create_DC14_(_1840_v22), globalState)).Cardinality())
				}
				return _1832_i1
			})(), globalState)
			_ = _rhs296
			var _lhs238 _dafny.Array = _1836_v18
			_ = _lhs238
			var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(698), _dafny.ArrayLen((_1836_v18), 0))
			_ = _lhs239
			_1835_v17 = _rhs292
			(_lhs238).ArraySet1(_rhs293, (_lhs239).Int())
			_1826_v9 = _rhs294
			_1839_v21 = _rhs295
			r0 = _rhs296
			var _1842_v24 *C4
			_ = _1842_v24
			var _nw266 *C4 = New_C4_()
			_ = _nw266
			_nw266.Ctor__()
			_1842_v24 = _nw266
		}
		r0 = ((_this).F15()).IsProperSubsetOf((_this).F15())
		return r0
	}
}
func (_this *C6) M9(p0 bool, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) (_dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var _1843_v0 _dafny.Sequence
		_ = _1843_v0
		_1843_v0 = _dafny.SeqOf((_this).F17())
		var _1844_v1 _dafny.Sequence
		_ = _1844_v1
		_1844_v1 = _dafny.SeqOf(p1, p1)
		var _1845_v2 _dafny.Map
		_ = _1845_v2
		_1845_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, _dafny.IntOfUint32((_1844_v1).Cardinality()))
		var _1846_v3 _dafny.Map
		_ = _1846_v3
		_1846_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_1845_v2).Update((_this).F13(), p1))
		var _1847_i0 _dafny.Int
		_ = _1847_i0
		_1847_i0 = _dafny.Zero
		{
			for (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1843_v0, _1843_v0)).Cardinality())).Cmp((_1846_v3).Cardinality()) == 0 {
				{
					if (_1847_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1847_i0 = (_1847_i0).Plus(_dafny.One)
					var _1848_v4 T2
					_ = _1848_v4
					var _nw267 *C5 = New_C5_()
					_ = _nw267
					_nw267.Ctor__(!((_this).F13()), (_this).F14(), (_this).F15(), Companion_D1_.Create_DC2_(p1, (_dafny.MultiSetOf(p3)).Cardinality(), (_this).F17()), p2)
					_1848_v4 = _nw267
					_1848_v4 = _1848_v4
					var _1849_v5 _dafny.Array
					_ = _1849_v5
					var _nw268 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
					_ = _nw268
					_1849_v5 = _nw268
					var _1850_v6 _dafny.Map
					_ = _1850_v6
					_1850_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1848_v4).F13(), _1849_v5)
					(globalState).F7 = (p1).Times((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1850_v6).Cardinality(), p1)))
					var _1851_v7 _dafny.Sequence
					_ = _1851_v7
					_1851_v7 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("thmbvav"))
					var _1852_v8 _dafny.Sequence
					_ = _1852_v8
					_1852_v8 = _dafny.UnicodeSeqOfUtf8Bytes("vhrtlm")
					_1851_v7 = _dafny.Companion_Sequence_.Concatenate(_1851_v7, _dafny.Companion_Sequence_.Update(_1851_v7, (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1851_v7).Cardinality()))).Uint32(), _1852_v8))
					r1 = (p1).Cmp(p1) == 0
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		if !(func() _dafny.Map {
			var _coll53 = _dafny.NewMapBuilder()
			_ = _coll53
			for _iter63 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-360))).Uint32(), func(coer99 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg99 _dafny.Int) interface{} {
					return coer99(arg99)
				}
			}((func(_1853_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_1854_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_1853_v0).Cardinality())
				}
			})(_1843_v0)))).Elements()); ; {
				_compr_53, _ok63 := _iter63()
				if !_ok63 {
					break
				}
				var _1855_v9 _dafny.Int
				_1855_v9 = interface{}(_compr_53).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-360))).Uint32(), func(coer100 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg100 _dafny.Int) interface{} {
						return coer100(arg100)
					}
				}((func(_1856_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
					return func(_1854_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfUint32((_1856_v0).Cardinality())
					}
				})(_1843_v0))), _1855_v9) {
					_coll53.Add(Companion_Default___.SafeDivisionInt(_1855_v9, p1), p0)
				}
			}
			return _coll53.ToMap()
		}()).Contains(p1) {
			r1 = false
			if !(p3) || ((_this).F17()) {
				var _1857_v10 _dafny.Array
				_ = _1857_v10
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_42
				var _nw269 _dafny.Array
				_ = _nw269
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw269 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Int = (func(_1858_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_1859_i2 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeDivisionInt(_1859_i2, _1858_p1)
						}
					})(p1)
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw269 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw269).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw269).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1857_v10 = _nw269
				var _index332 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1857_v10), 0))
				_ = _index332
				(_1857_v10).ArraySet1(p1, (_index332).Int())
				var _index333 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1857_v10), 0))
				_ = _index333
				(_1857_v10).ArraySet1((p1).Minus(Companion_Default___.SafeDivisionInt(p1, p1)), (_index333).Int())
				var _1860_v11 _dafny.CodePoint
				_ = _1860_v11
				_1860_v11 = _dafny.CodePoint('d')
				var _1861_v12 _dafny.MultiSet
				_ = _1861_v12
				_1861_v12 = _dafny.MultiSetOf(_dafny.CodePoint('l'), _1860_v11, _dafny.CodePoint('q'), _dafny.CodePoint('u'))
				_1861_v12 = _1861_v12
				r1 = false
				var _1862_v13 _dafny.Array
				_ = _1862_v13
				var _len0_43 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_43
				var _nw270 _dafny.Array
				_ = _nw270
				if _len0_43.Cmp(_dafny.Zero) == 0 {
					_nw270 = _dafny.NewArray(_len0_43)
				} else {
					var _init43 func(_dafny.Int) D7 = (func(_1863_p1 _dafny.Int) func(_dafny.Int) D7 {
						return func(_1864_i3 _dafny.Int) D7 {
							return Companion_D7_.Create_DC15_(_dafny.SeqOf(_1863_p1))
						}
					})(p1)
					_ = _init43
					var _element0_43 = _init43(_dafny.Zero)
					_ = _element0_43
					_nw270 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
					(_nw270).ArraySet1(_element0_43, 0)
					var _nativeLen0_43 = (_len0_43).Int()
					_ = _nativeLen0_43
					for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
						(_nw270).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
					}
				}
				_1862_v13 = _nw270
				var _1865_v14 _dafny.Set
				_ = _1865_v14
				_1865_v14 = _dafny.SetOf(_1862_v13, _1862_v13, _1862_v13)
				var _index334 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1857_v10), 0))
				_ = _index334
				var _rhs297 _dafny.Int = (_1857_v10).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1857_v10), 0))).Int()).(_dafny.Int)
				_ = _rhs297
				var _rhs298 bool = !(_1865_v14).Equals(_1865_v14)
				_ = _rhs298
				var _lhs240 _dafny.Array = _1857_v10
				_ = _lhs240
				var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(441), _dafny.ArrayLen((_1857_v10), 0))
				_ = _lhs241
				(_lhs240).ArraySet1(_rhs297, (_lhs241).Int())
				r1 = _rhs298
				var _1866_v15 _dafny.Array
				_ = _1866_v15
				var _len0_44 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_44
				var _nw271 _dafny.Array
				_ = _nw271
				if _len0_44.Cmp(_dafny.Zero) == 0 {
					_nw271 = _dafny.NewArray(_len0_44)
				} else {
					var _init44 func(_dafny.Int) bool = (func(_1867_p1 _dafny.Int) func(_dafny.Int) bool {
						return func(_1868_i4 _dafny.Int) bool {
							return (_dafny.IntOfInt64(545)).Cmp(_1867_p1) <= 0
						}
					})(p1)
					_ = _init44
					var _element0_44 = _init44(_dafny.Zero)
					_ = _element0_44
					_nw271 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
					(_nw271).ArraySet1(_element0_44, 0)
					var _nativeLen0_44 = (_len0_44).Int()
					_ = _nativeLen0_44
					for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
						(_nw271).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
					}
				}
				_1866_v15 = _nw271
				var _1869_v16 _dafny.Sequence
				_ = _1869_v16
				_1869_v16 = _dafny.UnicodeSeqOfUtf8Bytes("fxdirlqrk")
				var _1870_v17 _dafny.Array
				_ = _1870_v17
				var _nw272 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
				_ = _nw272
				_1870_v17 = _nw272
				var _index335 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1870_v17), 0))
				_ = _index335
				(_1870_v17).ArraySet1(_1866_v15, (_index335).Int())
				var _index336 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1870_v17), 0))
				_ = _index336
				var _rhs299 bool = (Companion_Default___.SafeDivisionInt(p1, p1)).Cmp(p1) == 0
				_ = _rhs299
				var _rhs300 _dafny.Array = _1857_v10
				_ = _rhs300
				var _rhs301 _dafny.Array = _1866_v15
				_ = _rhs301
				var _rhs302 _dafny.Sequence = _1869_v16
				_ = _rhs302
				var _rhs303 _dafny.Array = (_this).F14()
				_ = _rhs303
				var _lhs242 _dafny.Array = _1870_v17
				_ = _lhs242
				var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(495), _dafny.ArrayLen((_1870_v17), 0))
				_ = _lhs243
				r1 = _rhs299
				_1857_v10 = _rhs300
				_1866_v15 = _rhs301
				_1869_v16 = _rhs302
				(_lhs242).ArraySet1(_rhs303, (_lhs243).Int())
			} else {
				r0 = p1
				var _1871_v18 _dafny.Map
				_ = _1871_v18
				_1871_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p1, p1), (p1).Plus(p1))
				_1871_v18 = _1871_v18
				var _1872_v19 _dafny.MultiSet
				_ = _1872_v19
				_1872_v19 = _dafny.MultiSetOf(p2, (_this).F13())
				var _1873_v20 D16
				_ = _1873_v20
				_1873_v20 = Companion_D16_.Create_DC42_(_1872_v19)
				_1873_v20 = _1873_v20
				r1 = (_this).F13()
				var _1874_v21 _dafny.Sequence
				_ = _1874_v21
				_1874_v21 = _dafny.UnicodeSeqOfUtf8Bytes("olfi")
				var _1875_v22 _dafny.Sequence
				_ = _1875_v22
				_1875_v22 = _dafny.SeqOf(_1874_v21, _1874_v21)
				var _1876_v23 D6
				_ = _1876_v23
				_1876_v23 = Companion_D6_.Create_DC12_((_this).F13(), _1874_v21, p1, p3)
				var _1877_v24 _dafny.Array
				_ = _1877_v24
				var _nwElement0_61 _dafny.Sequence = _1875_v22
				_ = _nwElement0_61
				var _nw273 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(16))
				_ = _nw273
				(_nw273).ArraySet1(_nwElement0_61, 0)
				(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1874_v21), _1875_v22), 1)
				(_nw273).ArraySet1(_1875_v22, 2)
				(_nw273).ArraySet1(_1875_v22, 3)
				(_nw273).ArraySet1(Companion_Default___.Fm45(_1876_v23, p1, globalState), 4)
				(_nw273).ArraySet1(_1875_v22, 5)
				(_nw273).ArraySet1(_1875_v22, 6)
				(_nw273).ArraySet1(_1875_v22, 7)
				(_nw273).ArraySet1(_1875_v22, 8)
				(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1875_v22, _1875_v22), 9)
				(_nw273).ArraySet1(_1875_v22, 10)
				(_nw273).ArraySet1(_1875_v22, 11)
				(_nw273).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-12))).Uint32(), func(coer101 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg101 _dafny.Int) interface{} {
						return coer101(arg101)
					}
				}(func(_1878_i5 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("msjf")
				})), 12)
				(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1875_v22, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1874_v21).Cardinality()), _dafny.IntOfUint32((_1875_v22).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("vqwt")), _1875_v22), 13)
				(_nw273).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1874_v21), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(721))).Uint32(), func(coer102 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg102 _dafny.Int) interface{} {
						return coer102(arg102)
					}
				}(func(_1879_i6 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("qpiiluosn")
				}))), 14)
				(_nw273).ArraySet1(Companion_Default___.Fm45(_1876_v23, _dafny.IntOfUint32((_dafny.SeqOf(_1871_v18, Companion_Default___.Fm24(p1, globalState))).Cardinality()), globalState), 15)
				_1877_v24 = _nw273
				var _index337 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1877_v24), 0))
				_ = _index337
				(_1877_v24).ArraySet1(_1875_v22, (_index337).Int())
				var _index338 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(149), _dafny.ArrayLen((_1877_v24), 0))
				_ = _index338
				(_1877_v24).ArraySet1(_1875_v22, (_index338).Int())
			}
			var _1880_v25 _dafny.CodePoint
			_ = _1880_v25
			_1880_v25 = _dafny.CodePoint('k')
			var _1881_v26 *C1
			_ = _1881_v26
			var _nw274 *C1 = New_C1_()
			_ = _nw274
			_nw274.Ctor__(_1880_v25, _this.F12(), (p1).Cmp((_this).Fm9(_1880_v25, globalState)) >= 0)
			_1881_v26 = _nw274
			r1 = (_this).F17()
			var _1882_v27 _dafny.Map
			_ = _1882_v27
			_1882_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, Companion_Default___.Fm17(globalState))
			var _1883_v28 _dafny.Map
			_ = _1883_v28
			_1883_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F17())
			var _1884_v29 _dafny.Map
			_ = _1884_v29
			_1884_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1881_v26.F23, p3)
			var _1885_v30 D13
			_ = _1885_v30
			_1885_v30 = Companion_D13_.Create_DC34_(p1, Companion_Default___.Fm65(p3, _dafny.SeqOf((_1884_v29).Cardinality()), _1880_v25, !(p2), globalState))
			var _pat_let_tv50 = p1
			_ = _pat_let_tv50
			var _1886_v31 _dafny.Array
			_ = _1886_v31
			var _nwElement0_62 _dafny.Int = _dafny.IntOfInt64(105)
			_ = _nwElement0_62
			var _nw275 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(29))
			_ = _nw275
			(_nw275).ArraySet1(_nwElement0_62, 0)
			(_nw275).ArraySet1(p1, 1)
			(_nw275).ArraySet1(p1, 2)
			(_nw275).ArraySet1(p1, 3)
			(_nw275).ArraySet1(p1, 4)
			(_nw275).ArraySet1(p1, 5)
			(_nw275).ArraySet1(p1, 6)
			(_nw275).ArraySet1(p1, 7)
			(_nw275).ArraySet1(p1, 8)
			(_nw275).ArraySet1(((_1882_v27).Cardinality()).Times(p1), 9)
			(_nw275).ArraySet1((_dafny.Zero).Minus((p1).Plus(p1)), 10)
			(_nw275).ArraySet1(p1, 11)
			(_nw275).ArraySet1(_dafny.IntOfInt64(916), 12)
			(_nw275).ArraySet1(p1, 13)
			(_nw275).ArraySet1((_1883_v28).Cardinality(), 14)
			(_nw275).ArraySet1(p1, 15)
			(_nw275).ArraySet1(_dafny.IntOfInt64(99), 16)
			(_nw275).ArraySet1(p1, 17)
			(_nw275).ArraySet1(p1, 18)
			(_nw275).ArraySet1(p1, 19)
			(_nw275).ArraySet1(p1, 20)
			(_nw275).ArraySet1(_dafny.IntOfInt64(-859), 21)
			(_nw275).ArraySet1(p1, 22)
			(_nw275).ArraySet1(p1, 23)
			(_nw275).ArraySet1(_dafny.IntOfInt64(-842), 24)
			(_nw275).ArraySet1(Companion_Default___.Fm0(_1843_v0, p3, !(p3), p1, globalState), 25)
			(_nw275).ArraySet1(p1, 26)
			(_nw275).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(p1)), 27)
			(_nw275).ArraySet1((func(_pat_let59_0 D13) D13 {
				return func(_1887_dt__update__tmp_h0 D13) D13 {
					return func(_pat_let60_0 _dafny.Int) D13 {
						return func(_1888_dt__update_hcf71_h0 _dafny.Int) D13 {
							return Companion_D13_.Create_DC34_(_1888_dt__update_hcf71_h0, (_1887_dt__update__tmp_h0).Dtor_cf72())
						}(_pat_let60_0)
					}(_pat_let_tv50)
				}(_pat_let59_0)
			}(_1885_v30)).Dtor_cf71(), 28)
			_1886_v31 = _nw275
			_1886_v31 = _1886_v31
		} else {
			r1 = (_this).F17()
			r1 = (_this).F17()
			var _1889_v32 _dafny.CodePoint
			_ = _1889_v32
			_1889_v32 = _dafny.CodePoint('r')
			var _1890_v33 _dafny.Map
			_ = _1890_v33
			_1890_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm8((_this).F17(), _this.F16(), globalState), _1889_v32)
			var _1891_v34 _dafny.MultiSet
			_ = _1891_v34
			_1891_v34 = _dafny.MultiSetOf(p1, p1)
			_1890_v33 = (_1890_v33).Update((_1891_v34).Cardinality(), _1889_v32)
			var _1892_v35 _dafny.Sequence
			_ = _1892_v35
			_1892_v35 = _dafny.UnicodeSeqOfUtf8Bytes("vquk")
			var _1893_v36 _dafny.Sequence
			_ = _1893_v36
			_1893_v36 = _dafny.SeqOf(_1892_v35)
			var _1894_v37 _dafny.Map
			_ = _1894_v37
			_1894_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F17())
			var _1895_v38 _dafny.Array
			_ = _1895_v38
			var _nwElement0_63 _dafny.Int = _dafny.IntOfInt64(466)
			_ = _nwElement0_63
			var _nw276 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_63, nil, _dafny.IntOfInt64(8))
			_ = _nw276
			(_nw276).ArraySet1(_nwElement0_63, 0)
			(_nw276).ArraySet1((_1894_v37).Cardinality(), 1)
			(_nw276).ArraySet1(p1, 2)
			(_nw276).ArraySet1(p1, 3)
			(_nw276).ArraySet1(p1, 4)
			(_nw276).ArraySet1(p1, 5)
			(_nw276).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_1892_v35).Cardinality())), 6)
			(_nw276).ArraySet1(_dafny.IntOfInt64(-758), 7)
			_1895_v38 = _nw276
			var _1896_v39 _dafny.Map
			_ = _1896_v39
			_1896_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1893_v36).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1893_v36).Cardinality()))).Uint32()).(_dafny.Sequence), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1895_v38, p1))
			_1896_v39 = (_1896_v39).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(937))).Uint32(), func(coer103 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg103 _dafny.Int) interface{} {
					return coer103(arg103)
				}
			}((func(_1897_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1898_i7 _dafny.Int) _dafny.CodePoint {
					return _1897_v32
				}
			})(_1889_v32))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1895_v38, _dafny.IntOfUint32((_1892_v35).Cardinality())))
			var _pat_let_tv51 = p1
			_ = _pat_let_tv51
			var _1899_v40 *C2
			_ = _1899_v40
			var _nw277 *C2 = New_C2_()
			_ = _nw277
			_nw277.Ctor__(Companion_Default___.SafeDivisionInt(p1, _dafny.IntOfInt64(18)), (p3) || (!(p0)), _this.F16(), p2, (_this).F14(), ((_this).F15()).Intersection((_this).F15()), func(_pat_let61_0 D1) D1 {
				return func(_1900_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let62_0 _dafny.Int) D1 {
						return func(_1901_dt__update_hcf2_h0 _dafny.Int) D1 {
							return Companion_D1_.Create_DC2_(_1901_dt__update_hcf2_h0, (_1900_dt__update__tmp_h1).Dtor_cf3(), (_1900_dt__update__tmp_h1).Dtor_cf4())
						}(_pat_let62_0)
					}(_pat_let_tv51)
				}(_pat_let61_0)
			}(Companion_D1_.Create_DC2_(p1, p1, !((_this).F13()))), (_this).F13())
			_1899_v40 = _nw277
		}
		r1 = (_this).F17()
		var _1902_v41 _dafny.CodePoint
		_ = _1902_v41
		_1902_v41 = _dafny.CodePoint('a')
		var _1903_v42 _dafny.Set
		_ = _1903_v42
		_1903_v42 = _dafny.SetOf(Companion_D3_.Create_DC5_(_1902_v41))
		var _1904_v43 _dafny.MultiSet
		_ = _1904_v43
		_1904_v43 = _dafny.MultiSetOf(p1, p1)
		var _pat_let_tv52 = _1903_v42
		_ = _pat_let_tv52
		var _pat_let_tv53 = _1902_v41
		_ = _pat_let_tv53
		var _pat_let_tv54 = _1902_v41
		_ = _pat_let_tv54
		var _pat_let_tv55 = _1902_v41
		_ = _pat_let_tv55
		var _pat_let_tv56 = _1902_v41
		_ = _pat_let_tv56
		var _pat_let_tv57 = _1902_v41
		_ = _pat_let_tv57
		var _pat_let_tv58 = _1903_v42
		_ = _pat_let_tv58
		var _rhs304 _dafny.Int = p1
		_ = _rhs304
		var _rhs305 _dafny.Int = _dafny.IntOfInt64(209)
		_ = _rhs305
		var _rhs306 _dafny.Set = func(_source26 D6) _dafny.Set {
			if _source26.Is_DC12() {
				var _1905___mcc_h0 bool = _source26.Get_().(D6_DC12).Cf18
				_ = _1905___mcc_h0
				var _1906___mcc_h1 _dafny.Sequence = _source26.Get_().(D6_DC12).Cf19
				_ = _1906___mcc_h1
				var _1907___mcc_h2 _dafny.Int = _source26.Get_().(D6_DC12).Cf20
				_ = _1907___mcc_h2
				var _1908___mcc_h3 bool = _source26.Get_().(D6_DC12).Cf21
				_ = _1908___mcc_h3
				var _1909_cf21 bool = _1908___mcc_h3
				_ = _1909_cf21
				var _1910_cf20 _dafny.Int = _1907___mcc_h2
				_ = _1910_cf20
				var _1911_cf19 _dafny.Sequence = _1906___mcc_h1
				_ = _1911_cf19
				var _1912_cf18 bool = _1905___mcc_h0
				_ = _1912_cf18
				return _pat_let_tv52
			} else if _source26.Is_DC13() {
				var _1913___mcc_h4 bool = _source26.Get_().(D6_DC13).Cf22
				_ = _1913___mcc_h4
				var _1914___mcc_h5 bool = _source26.Get_().(D6_DC13).Cf23
				_ = _1914___mcc_h5
				var _1915___mcc_h6 _dafny.Int = _source26.Get_().(D6_DC13).Cf24
				_ = _1915___mcc_h6
				var _1916___mcc_h7 _dafny.Int = _source26.Get_().(D6_DC13).Cf25
				_ = _1916___mcc_h7
				var _1917___mcc_h8 bool = _source26.Get_().(D6_DC13).Cf26
				_ = _1917___mcc_h8
				var _1918_cf26 bool = _1917___mcc_h8
				_ = _1918_cf26
				var _1919_cf25 _dafny.Int = _1916___mcc_h7
				_ = _1919_cf25
				var _1920_cf24 _dafny.Int = _1915___mcc_h6
				_ = _1920_cf24
				var _1921_cf23 bool = _1914___mcc_h5
				_ = _1921_cf23
				var _1922_cf22 bool = _1913___mcc_h4
				_ = _1922_cf22
				return _dafny.SetOf(Companion_D3_.Create_DC5_(_pat_let_tv53), Companion_D3_.Create_DC5_(_pat_let_tv54))
			} else if _source26.Is_DC11() {
				var _1923___mcc_h9 _dafny.MultiSet = _source26.Get_().(D6_DC11).Cf17
				_ = _1923___mcc_h9
				var _1924_cf17 _dafny.MultiSet = _1923___mcc_h9
				_ = _1924_cf17
				return _dafny.SetOf(Companion_D3_.Create_DC5_(_pat_let_tv55), Companion_D3_.Create_DC5_(_pat_let_tv56), Companion_D3_.Create_DC5_(_pat_let_tv57))
			} else {
				var _1925___mcc_h10 D6 = _source26.Get_().(D6_DC14).Cf27
				_ = _1925___mcc_h10
				var _1926_cf27 D6 = _1925___mcc_h10
				_ = _1926_cf27
				return _pat_let_tv58
			}
		}(Companion_D6_.Create_DC11_(_1904_v43))
		_ = _rhs306
		var _rhs307 bool = p0
		_ = _rhs307
		r0 = _rhs304
		r0 = _rhs305
		_1903_v42 = _rhs306
		r1 = _rhs307
		r1 = p3
		if (p3) || (!((_this).F17())) {
			var _1927_v44 _dafny.Sequence
			_ = _1927_v44
			_1927_v44 = _dafny.UnicodeSeqOfUtf8Bytes("wypej")
			var _1928_v45 _dafny.Array
			_ = _1928_v45
			var _nw278 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw278
			_1928_v45 = _nw278
			var _index339 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))
			_ = _index339
			(_1928_v45).ArraySet1((_dafny.MultiSetOf((_this).F13(), p3)).Cardinality(), (_index339).Int())
			var _index340 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))
			_ = _index340
			var _rhs308 _dafny.Int = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_this).F17() {
					return p1
				}
				return _dafny.IntOfInt64(36)
			})(), p1)
			_ = _rhs308
			var _rhs309 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1927_v44, _1927_v44), _1927_v44)
			_ = _rhs309
			var _rhs310 _dafny.CodePoint = _1902_v41
			_ = _rhs310
			var _rhs311 bool = (p1).Cmp(p1) == 0
			_ = _rhs311
			var _rhs312 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _1843_v0), _1843_v0)).Cardinality())
			_ = _rhs312
			var _lhs244 _dafny.Array = _1928_v45
			_ = _lhs244
			var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))
			_ = _lhs245
			r0 = _rhs308
			_1927_v44 = _rhs309
			_1902_v41 = _rhs310
			r1 = _rhs311
			(_lhs244).ArraySet1(_rhs312, (_lhs245).Int())
			var _1929_v46 *C1
			_ = _1929_v46
			var _nw279 *C1 = New_C1_()
			_ = _nw279
			_nw279.Ctor__(_1902_v41, _this.F12(), p2)
			_1929_v46 = _nw279
			var _1930_v47 _dafny.Sequence
			_ = _1930_v47
			_1930_v47 = _dafny.SeqOf(_1929_v46)
			var _1931_v48 _dafny.Array
			_ = _1931_v48
			var _nwElement0_64 *C1 = _1929_v46
			_ = _nwElement0_64
			var _nw280 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_64, nil, _dafny.IntOfInt64(13))
			_ = _nw280
			(_nw280).ArraySet1(_nwElement0_64, 0)
			(_nw280).ArraySet1(_1929_v46, 1)
			(_nw280).ArraySet1(_1929_v46, 2)
			(_nw280).ArraySet1(_1929_v46, 3)
			(_nw280).ArraySet1(_1929_v46, 4)
			(_nw280).ArraySet1(_1929_v46, 5)
			(_nw280).ArraySet1(_1929_v46, 6)
			(_nw280).ArraySet1(_1929_v46, 7)
			(_nw280).ArraySet1(_1929_v46, 8)
			(_nw280).ArraySet1(_1929_v46, 9)
			(_nw280).ArraySet1((_1930_v47).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1930_v47).Cardinality()))).Uint32()).(*C1), 10)
			(_nw280).ArraySet1(_1929_v46, 11)
			(_nw280).ArraySet1(_1929_v46, 12)
			_1931_v48 = _nw280
			var _1932_v49 D17
			_ = _1932_v49
			_1932_v49 = Companion_D17_.Create_DC44_(_1931_v48)
			var _1933_v50 _dafny.Map
			_ = _1933_v50
			_1933_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_dafny.SeqOf(p3), (_this).F13(), p3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(527))).Uint32(), func(coer104 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg104 _dafny.Int) interface{} {
					return coer104(arg104)
				}
			}((func(_1934_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1935_i8 _dafny.Int) _dafny.CodePoint {
					return _1934_v41
				}
			})(_1902_v41)))).Cardinality()), globalState), _1931_v48)
			var _1936_v51 _dafny.Array
			_ = _1936_v51
			var _nwElement0_65 _dafny.Array = _1931_v48
			_ = _nwElement0_65
			var _nw281 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_65, nil, _dafny.IntOfInt64(14))
			_ = _nw281
			(_nw281).ArraySet1(_nwElement0_65, 0)
			(_nw281).ArraySet1(_1931_v48, 1)
			(_nw281).ArraySet1(_1931_v48, 2)
			(_nw281).ArraySet1(_1931_v48, 3)
			(_nw281).ArraySet1(_1931_v48, 4)
			(_nw281).ArraySet1((_1932_v49).Dtor_cf87(), 5)
			(_nw281).ArraySet1(_1931_v48, 6)
			(_nw281).ArraySet1(_1931_v48, 7)
			(_nw281).ArraySet1((func() _dafny.Array {
				if (_1933_v50).Contains((_1928_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))).Int()).(_dafny.Int)) {
					return (_1933_v50).Get((_1928_v45).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))).Int()).(_dafny.Int)).(_dafny.Array)
				}
				return _1931_v48
			})(), 8)
			(_nw281).ArraySet1(_1931_v48, 9)
			(_nw281).ArraySet1(_1931_v48, 10)
			(_nw281).ArraySet1(_1931_v48, 11)
			(_nw281).ArraySet1(_1931_v48, 12)
			(_nw281).ArraySet1(_1931_v48, 13)
			_1936_v51 = _nw281
			var _index341 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1936_v51), 0))
			_ = _index341
			(_1936_v51).ArraySet1(_1931_v48, (_index341).Int())
			var _1937_v52 _dafny.Sequence
			_ = _1937_v52
			_1937_v52 = _dafny.SeqOf((func() _dafny.Array {
				if false {
					return _1931_v48
				}
				return _1931_v48
			})(), _1931_v48, _1931_v48, _1931_v48, _1931_v48)
			var _index342 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_1936_v51), 0))
			_ = _index342
			(_1936_v51).ArraySet1((_1937_v52).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1937_v52).Cardinality()))).Uint32()).(_dafny.Array), (_index342).Int())
			var _index343 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(216), _dafny.ArrayLen((_1928_v45), 0))
			_ = _index343
			(_1928_v45).ArraySet1((_dafny.Zero).Minus(p1), (_index343).Int())
			r1 = (_this).F17()
			(globalState).F7 = p1
		} else {
			var _1938_v53 _dafny.Sequence
			_ = _1938_v53
			_1938_v53 = _dafny.UnicodeSeqOfUtf8Bytes("uvv")
			(globalState).F7 = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if p0 {
					return _1938_v53
				}
				return _1938_v53
			})()).Cardinality())
			var _1939_v54 _dafny.Map
			_ = _1939_v54
			_1939_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F17())
			var _1940_v55 *C2
			_ = _1940_v55
			var _nw282 *C2 = New_C2_()
			_ = _nw282
			_nw282.Ctor__(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("n"), _dafny.UnicodeSeqOfUtf8Bytes("bndxiw"))).Cardinality()), p0, _this.F16(), (func() bool {
				if (_1939_v54).Contains(p1) {
					return (_1939_v54).Get(p1).(bool)
				}
				return (_this).F13()
			})(), (_this).F14(), (_this).F15(), _this.F12(), (_this).F17())
			_1940_v55 = _nw282
			_1940_v55 = _1940_v55
			var _1941_v56 _dafny.Map
			_ = _1941_v56
			_1941_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1940_v55, (_this).F17())
			var _1942_v57 D6
			_ = _1942_v57
			_1942_v57 = Companion_D6_.Create_DC12_(((_1940_v55).F21()).Cmp((_1940_v55).F21()) != 0, _dafny.Companion_Sequence_.Concatenate(_1938_v53, _dafny.Companion_Sequence_.Update(_1938_v53, (Companion_Default___.SafeIndex((_1940_v55).F21(), _dafny.IntOfUint32((_1938_v53).Cardinality()))).Uint32(), _1902_v41)), ((_1941_v56).Merge(_1941_v56)).Cardinality(), p3)
			_1942_v57 = Companion_Default___.Fm39(!(p3) || (false), p3, globalState)
			if p0 {
				var _1943_v58 _dafny.Array
				_ = _1943_v58
				var _len0_45 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_45
				var _nw283 _dafny.Array
				_ = _nw283
				if _len0_45.Cmp(_dafny.Zero) == 0 {
					_nw283 = _dafny.NewArray(_len0_45)
				} else {
					var _init45 func(_dafny.Int) _dafny.CodePoint = (func(_1944_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1945_i9 _dafny.Int) _dafny.CodePoint {
							return _1944_v41
						}
					})(_1902_v41)
					_ = _init45
					var _element0_45 = _init45(_dafny.Zero)
					_ = _element0_45
					_nw283 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
					(_nw283).ArraySet1CodePoint(_element0_45, 0)
					var _nativeLen0_45 = (_len0_45).Int()
					_ = _nativeLen0_45
					for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
						(_nw283).ArraySet1CodePoint(_init45(_dafny.IntOf(_i0_45)), _i0_45)
					}
				}
				_1943_v58 = _nw283
				var _1946_v59 D13
				_ = _1946_v59
				_1946_v59 = Companion_D13_.Create_DC33_(_1943_v58)
				var _1947_v60 D13
				_ = _1947_v60
				_1947_v60 = Companion_D13_.Create_DC35_(_1946_v59)
				var _1948_v61 D13
				_ = _1948_v61
				_1948_v61 = Companion_D13_.Create_DC35_((_1947_v60).Dtor_cf73())
				_1948_v61 = _1948_v61
				var _index344 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index344
				((_this).F14()).ArraySet1((_this).F13(), (_index344).Int())
				var _1949_v62 _dafny.MultiSet
				_ = _1949_v62
				_1949_v62 = _dafny.MultiSetOf((_1940_v55).F22(), !((_this).F13()), !((_1940_v55).F22()))
				var _1950_v63 D6
				_ = _1950_v63
				_1950_v63 = Companion_D6_.Create_DC13_(p2, p2, (_1940_v55).F21(), (_1949_v62).Cardinality(), !(p0))
				var _index345 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index345
				var _rhs313 bool = (_1950_v63).Dtor_cf23()
				_ = _rhs313
				var _rhs314 *C2 = _1940_v55
				_ = _rhs314
				var _rhs315 bool = (_1940_v55).F22()
				_ = _rhs315
				var _lhs246 _dafny.Array = (_this).F14()
				_ = _lhs246
				var _lhs247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(354), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs247
				r1 = _rhs313
				_1940_v55 = _rhs314
				(_lhs246).ArraySet1(_rhs315, (_lhs247).Int())
				var _1951_v64 _dafny.Array
				_ = _1951_v64
				var _len0_46 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_46
				var _nw284 _dafny.Array
				_ = _nw284
				if _len0_46.Cmp(_dafny.Zero) == 0 {
					_nw284 = _dafny.NewArray(_len0_46)
				} else {
					var _init46 func(_dafny.Int) _dafny.Int = (func(_1952_v55 *C2) func(_dafny.Int) _dafny.Int {
						return func(_1953_i10 _dafny.Int) _dafny.Int {
							return (_1953_i10).Plus(_dafny.IntOfUint32((_dafny.SeqOf((_1952_v55).F22())).Cardinality()))
						}
					})(_1940_v55)
					_ = _init46
					var _element0_46 = _init46(_dafny.Zero)
					_ = _element0_46
					_nw284 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
					(_nw284).ArraySet1(_element0_46, 0)
					var _nativeLen0_46 = (_len0_46).Int()
					_ = _nativeLen0_46
					for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
						(_nw284).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
					}
				}
				_1951_v64 = _nw284
				_1951_v64 = _1951_v64
				var _1954_v65 _dafny.Map
				_ = _1954_v65
				_1954_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1940_v55).F21(), (_1940_v55).F21())
				var _1955_v66 D11
				_ = _1955_v66
				_1955_v66 = Companion_D11_.Create_DC27_((_this).F13(), _1902_v41, (func() _dafny.Int {
					if (_1954_v65).Contains((_1940_v55).F21()) {
						return (_1954_v65).Get((_1940_v55).F21()).(_dafny.Int)
					}
					return (_dafny.Zero).Minus((_1940_v55).F21())
				})(), (_1940_v55).F21(), _1844_v1)
				_1955_v66 = _1955_v66
				var _1956_v67 _dafny.Array
				_ = _1956_v67
				var _nw285 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw285
				_1956_v67 = _nw285
				_1956_v67 = _1956_v67
			} else {
				var _index346 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index346
				((_this).F14()).ArraySet1(false, (_index346).Int())
				var _1957_v68 _dafny.Array
				_ = _1957_v68
				var _nw286 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(7))
				_ = _nw286
				_1957_v68 = _nw286
				var _1958_v69 _dafny.Array
				_ = _1958_v69
				var _nwElement0_66 _dafny.MultiSet = _1904_v43
				_ = _nwElement0_66
				var _nw287 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_66, nil, _dafny.IntOfInt64(15))
				_ = _nw287
				(_nw287).ArraySet1(_nwElement0_66, 0)
				(_nw287).ArraySet1(_1904_v43, 1)
				(_nw287).ArraySet1((_dafny.MultiSetFromSeq(_1844_v1)).Intersection(_1904_v43), 2)
				(_nw287).ArraySet1(_1904_v43, 3)
				(_nw287).ArraySet1((_1904_v43).Intersection(Companion_Default___.Fm19(globalState)), 4)
				(_nw287).ArraySet1((_1904_v43).Union(_1904_v43), 5)
				(_nw287).ArraySet1(_1904_v43, 6)
				(_nw287).ArraySet1(_1904_v43, 7)
				(_nw287).ArraySet1(_1904_v43, 8)
				(_nw287).ArraySet1((_1904_v43).Union(_1904_v43), 9)
				(_nw287).ArraySet1(_1904_v43, 10)
				(_nw287).ArraySet1((_1904_v43).Union(_dafny.MultiSetFromSeq(Companion_Default___.Fm27(globalState))), 11)
				(_nw287).ArraySet1(_1904_v43, 12)
				(_nw287).ArraySet1((_1904_v43).Update((_1940_v55).F21(), Companion_Default___.Abs((_1940_v55).F21())), 13)
				(_nw287).ArraySet1((_1904_v43).Update((_dafny.Zero).Minus(p1), Companion_Default___.Abs((_dafny.Zero).Minus((_1940_v55).F21()))), 14)
				_1958_v69 = _nw287
				var _index347 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_1958_v69), 0))
				_ = _index347
				(_1958_v69).ArraySet1((Companion_Default___.Fm56(p1, (_this).F15(), globalState)).Intersection(_1904_v43), (_index347).Int())
				var _1959_v70 _dafny.Map
				_ = _1959_v70
				_1959_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1940_v55).F21(), (_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_1843_v0)).Cardinality()))
				var _index348 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index348
				var _index349 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_1958_v69), 0))
				_ = _index349
				var _rhs316 _dafny.Int = Companion_Default___.SafeDivisionInt(((_1940_v55).F21()).Times((_1940_v55).F21()), _dafny.IntOfInt64(721))
				_ = _rhs316
				var _rhs317 bool = Companion_Default___.Fm30(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1959_v70).Update((_1940_v55).F21(), _dafny.IntOfInt64(217))).Cardinality(), p1)).Cardinality()).Times((_1940_v55).F21()), globalState)
				_ = _rhs317
				var _rhs318 _dafny.Array = (_this).F14()
				_ = _rhs318
				var _rhs319 _dafny.MultiSet = (_1904_v43).Union(_1904_v43)
				_ = _rhs319
				var _rhs320 _dafny.CodePoint = _1902_v41
				_ = _rhs320
				var _lhs248 *GlobalState = globalState
				_ = _lhs248
				var _lhs249 _dafny.Array = (_this).F14()
				_ = _lhs249
				var _lhs250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs250
				var _lhs251 _dafny.Array = _1958_v69
				_ = _lhs251
				var _lhs252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(544), _dafny.ArrayLen((_1958_v69), 0))
				_ = _lhs252
				_lhs248.F7 = _rhs316
				(_lhs249).ArraySet1(_rhs317, (_lhs250).Int())
				_1957_v68 = _rhs318
				(_lhs251).ArraySet1(_rhs319, (_lhs252).Int())
				_1902_v41 = _rhs320
				r1 = (p1).Cmp(((_1940_v55).F21()).Times((_dafny.Zero).Minus((_1940_v55).Fm9(Companion_Default___.Fm18((_1940_v55).F21(), _1938_v53, p2, _1845_v2, globalState), globalState)))) > 0
				r0 = (func() _dafny.Int {
					if p2 {
						return (_dafny.Zero).Minus((_1940_v55).F21())
					}
					return (_1940_v55).F21()
				})()
				(globalState).F7 = (_1940_v55).F21()
				var _1960_v71 _dafny.Array
				_ = _1960_v71
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_47
				var _nw288 _dafny.Array
				_ = _nw288
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw288 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) _dafny.Int = (func(_1961_v55 *C2) func(_dafny.Int) _dafny.Int {
						return func(_1962_i11 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1962_i11, (_1961_v55).F21())
						}
					})(_1940_v55)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw288 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw288).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw288).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1960_v71 = _nw288
				var _index350 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1960_v71), 0))
				_ = _index350
				(_1960_v71).ArraySet1(p1, (_index350).Int())
				var _1963_v72 _dafny.Map
				_ = _1963_v72
				_1963_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)), _1845_v2)
				var _index351 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1960_v71), 0))
				_ = _index351
				var _rhs321 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).Fm9(_1902_v41, globalState), (_1940_v55).F21())), (p1).Times((_1940_v55).F21()))
				_ = _rhs321
				var _rhs322 _dafny.Array = _1957_v68
				_ = _rhs322
				var _rhs323 _dafny.Int = (_1940_v55).F21()
				_ = _rhs323
				var _rhs324 _dafny.Map = ((func() _dafny.Map {
					if (_this).F13() {
						return _1963_v72
					}
					return _1963_v72
				})()).Update((_1940_v55).F22(), (_1845_v2).Update((_this).F17(), _dafny.IntOfUint32((_1844_v1).Cardinality())))
				_ = _rhs324
				var _rhs325 bool = (_1940_v55).F22()
				_ = _rhs325
				var _lhs253 *GlobalState = globalState
				_ = _lhs253
				var _lhs254 _dafny.Array = _1960_v71
				_ = _lhs254
				var _lhs255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(602), _dafny.ArrayLen((_1960_v71), 0))
				_ = _lhs255
				_lhs253.F7 = _rhs321
				_1957_v68 = _rhs322
				(_lhs254).ArraySet1(_rhs323, (_lhs255).Int())
				_1963_v72 = _rhs324
				r1 = _rhs325
			}
			(globalState).F7 = (Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1845_v2).Contains(p2) {
					return (_1845_v2).Get(p2).(_dafny.Int)
				}
				return p1
			})(), (_1940_v55).F21())).Minus(p1)
		}
		r0 = p1
		r1 = ((Companion_Default___.Fm56(_dafny.IntOfInt64(414), (_this).F15(), globalState)).Cardinality()).Cmp(_dafny.IntOfInt64(756)) != 0
		return r0, r1
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
func (_this *C7) Fm6(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		if !(!(true)) {
			return (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('d'))).Cardinality()))).Cardinality()))).Times(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())))
		} else {
			return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pkq")).Cardinality()))).Cardinality()).Times(_dafny.IntOfInt64(959))
		}
	}
}
func (_this *C7) Fm7(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.Sequence, p3 _dafny.Set, globalState *GlobalState) bool {
	{
		return ((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-407))).Uint32(), func(coer105 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg105 _dafny.Int) interface{} {
				return coer105(arg105)
			}
		}(func(_1964_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()))).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("t"), _dafny.UnicodeSeqOfUtf8Bytes("sgoyyjn"))).Cardinality())) >= 0
	}
}
func (_this *C7) M4(p0 bool, p1 _dafny.Sequence, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1965_v0 _dafny.Array
		_ = _1965_v0
		var _len0_48 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_48
		var _nw289 _dafny.Array
		_ = _nw289
		if _len0_48.Cmp(_dafny.Zero) == 0 {
			_nw289 = _dafny.NewArray(_len0_48)
		} else {
			var _init48 func(_dafny.Int) _dafny.Int = (func(_1966_p3 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_1967_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1967_i0, _1966_p3)
				}
			})(p3)
			_ = _init48
			var _element0_48 = _init48(_dafny.Zero)
			_ = _element0_48
			_nw289 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
			(_nw289).ArraySet1(_element0_48, 0)
			var _nativeLen0_48 = (_len0_48).Int()
			_ = _nativeLen0_48
			for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
				(_nw289).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
			}
		}
		_1965_v0 = _nw289
		var _1968_v1 _dafny.Map
		_ = _1968_v1
		_1968_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1965_v0, p0)
		_1968_v1 = (_1968_v1).Update(_1965_v0, p0)
		var _1969_i1 _dafny.Int
		_ = _1969_i1
		_1969_i1 = _dafny.Zero
		{
			for !(p0) {
				{
					if (_1969_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1969_i1 = (_1969_i1).Plus(_dafny.One)
					var _1970_v2 _dafny.CodePoint
					_ = _1970_v2
					_1970_v2 = _dafny.CodePoint('j')
					var _1971_v3 _dafny.Sequence
					_ = _1971_v3
					_1971_v3 = _dafny.UnicodeSeqOfUtf8Bytes("yqxdb")
					_1970_v2 = (_1971_v3).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1971_v3).Cardinality()))).Uint32()).(_dafny.CodePoint)
					var _1972_v4 _dafny.Array
					_ = _1972_v4
					var _nw290 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
					_ = _nw290
					_1972_v4 = _nw290
					var _1973_v5 _dafny.Map
					_ = _1973_v5
					_1973_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
					var _1974_v6 _dafny.Set
					_ = _1974_v6
					_1974_v6 = _dafny.SetOf(p0, p0)
					var _index352 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1972_v4), 0))
					_ = _index352
					(_1972_v4).ArraySet1((_this).Fm7(p2, _1973_v5, _1971_v3, _1974_v6, globalState), (_index352).Int())
					var _1975_v7 D2
					_ = _1975_v7
					_1975_v7 = Companion_D2_.Create_DC4_(_1965_v0, p0, p3, p0, _1972_v4)
					var _index353 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(870), _dafny.ArrayLen((_1972_v4), 0))
					_ = _index353
					(_1972_v4).ArraySet1((_1975_v7).Dtor_cf9(), (_index353).Int())
					_1971_v3 = _1971_v3
					var _index354 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1965_v0), 0))
					_ = _index354
					(_1965_v0).ArraySet1(p3, (_index354).Int())
					var _index355 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(80), _dafny.ArrayLen((_1965_v0), 0))
					_ = _index355
					(_1965_v0).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(901), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vtl")).Cardinality())), p3), (_index355).Int())
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		var _1976_v8 _dafny.Array
		_ = _1976_v8
		var _nw291 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
		_ = _nw291
		_1976_v8 = _nw291
		var _1977_v9 D2
		_ = _1977_v9
		_1977_v9 = Companion_D2_.Create_DC4_(_1965_v0, !(p0), p3, p0, _1976_v8)
		var _source27 D2 = _1977_v9
		_ = _source27
		if _source27.Is_DC4() {
			var _1978___mcc_h0 _dafny.Array = _source27.Get_().(D2_DC4).Cf6
			_ = _1978___mcc_h0
			var _1979___mcc_h1 bool = _source27.Get_().(D2_DC4).Cf7
			_ = _1979___mcc_h1
			var _1980___mcc_h2 _dafny.Int = _source27.Get_().(D2_DC4).Cf8
			_ = _1980___mcc_h2
			var _1981___mcc_h3 bool = _source27.Get_().(D2_DC4).Cf9
			_ = _1981___mcc_h3
			var _1982___mcc_h4 _dafny.Array = _source27.Get_().(D2_DC4).Cf10
			_ = _1982___mcc_h4
			var _1983_cf10 _dafny.Array = _1982___mcc_h4
			_ = _1983_cf10
			var _1984_cf9 bool = _1981___mcc_h3
			_ = _1984_cf9
			var _1985_cf8 _dafny.Int = _1980___mcc_h2
			_ = _1985_cf8
			var _1986_cf7 bool = _1979___mcc_h1
			_ = _1986_cf7
			var _1987_cf6 _dafny.Array = _1978___mcc_h0
			_ = _1987_cf6
			_1985_cf8 = p3
			var _1988_v10 _dafny.CodePoint
			_ = _1988_v10
			_1988_v10 = _dafny.CodePoint('o')
			var _1989_v11 _dafny.MultiSet
			_ = _1989_v11
			var _out14 _dafny.MultiSet
			_ = _out14
			_out14 = Companion_Default___.M0(_1988_v10, p0, globalState)
			_1989_v11 = _out14
			var _1990_v12 _dafny.Sequence
			_ = _1990_v12
			_1990_v12 = _dafny.UnicodeSeqOfUtf8Bytes("c")
			var _1991_v13 D1
			_ = _1991_v13
			_1991_v13 = Companion_D1_.Create_DC1_(_1990_v12)
			var _1992_v14 _dafny.Set
			_ = _1992_v14
			_1992_v14 = _dafny.SetOf(p2)
			var _1993_v15 _dafny.Map
			_ = _1993_v15
			_1993_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1986_cf7, p3)
			var _1994_v16 D1
			_ = _1994_v16
			_1994_v16 = Companion_D1_.Create_DC2_((_1993_v15).Cardinality(), _1985_cf8, false)
			var _pat_let_tv59 = _1990_v12
			_ = _pat_let_tv59
			var _1995_v17 *C6
			_ = _1995_v17
			var _nw292 *C6 = New_C6_()
			_ = _nw292
			_nw292.Ctor__(func(_pat_let63_0 D1) D1 {
				return func(_1996_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let64_0 _dafny.Sequence) D1 {
						return func(_1997_dt__update_hcf1_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC1_(_1997_dt__update_hcf1_h0)
						}(_pat_let64_0)
					}(_pat_let_tv59)
				}(_pat_let63_0)
			}(_1991_v13), _1984_cf9, _1976_v8, _1992_v14, _1994_v16, true)
			_1995_v17 = _nw292
			var _index356 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1987_cf6), 0))
			_ = _index356
			(_1987_cf6).ArraySet1(_1985_cf8, (_index356).Int())
			var _1998_v18 _dafny.Sequence
			_ = _1998_v18
			_1998_v18 = _dafny.SeqOf(p1)
			var _index357 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(95), _dafny.ArrayLen((_1987_cf6), 0))
			_ = _index357
			(_1987_cf6).ArraySet1(Companion_Default___.Fm0((_1998_v18).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.IntOfUint32((_1998_v18).Cardinality()))).Uint32()).(_dafny.Sequence), _1984_cf9, _1986_cf7, (_dafny.Zero).Minus(p2), globalState), (_index357).Int())
		} else {
			var _1999___mcc_h5 _dafny.Array = _source27.Get_().(D2_DC3).Cf5
			_ = _1999___mcc_h5
			var _2000_cf5 _dafny.Array = _1999___mcc_h5
			_ = _2000_cf5
			var _2001_v19 _dafny.Sequence
			_ = _2001_v19
			_2001_v19 = _dafny.UnicodeSeqOfUtf8Bytes("pliy")
			var _2002_v20 _dafny.MultiSet
			_ = _2002_v20
			_2002_v20 = _dafny.MultiSetOf(_2001_v19, _2001_v19)
			var _2003_v21 _dafny.MultiSet
			_ = _2003_v21
			_2003_v21 = _dafny.MultiSetOf(p0)
			var _2004_v22 _dafny.Map
			_ = _2004_v22
			_2004_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2002_v20, _2003_v21)
			var _2005_v23 _dafny.CodePoint
			_ = _2005_v23
			_2005_v23 = _dafny.CodePoint('h')
			var _2006_v24 _dafny.Map
			_ = _2006_v24
			_2006_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2005_v23, true)
			var _2007_v25 _dafny.Map
			_ = _2007_v25
			_2007_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2006_v24).Cardinality(), p0)
			var _2008_v26 _dafny.Set
			_ = _2008_v26
			_2008_v26 = _dafny.SetOf((_2007_v25).Cardinality(), p3)
			var _2009_v27 T2
			_ = _2009_v27
			var _nw293 *C5 = New_C5_()
			_ = _nw293
			_nw293.Ctor__(p0, _2000_cf5, _2008_v26, Companion_D1_.Create_DC2_(_dafny.IntOfInt64(-499), (_dafny.SetOf(_2005_v23)).Cardinality(), p0), p0)
			_2009_v27 = _nw293
			var _2010_v28 _dafny.Map
			_ = _2010_v28
			_2010_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2009_v27)
			_2004_v22 = (_2004_v22).Update(Companion_Default___.Fm66(p3, globalState), (_2003_v21).Update(p0, Companion_Default___.Abs((_2010_v28).Cardinality())))
			var _2011_v29 _dafny.Array
			_ = _2011_v29
			var _nw294 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(13))
			_ = _nw294
			_2011_v29 = _nw294
			var _2012_v30 D17
			_ = _2012_v30
			_2012_v30 = Companion_D17_.Create_DC44_(_2011_v29)
			var _source28 D17 = _2012_v30
			_ = _source28
			if _source28.Is_DC45() {
				var _2013___mcc_h6 bool = _source28.Get_().(D17_DC45).Cf88
				_ = _2013___mcc_h6
				var _2014_cf88 bool = _2013___mcc_h6
				_ = _2014_cf88
				var _2015_v31 _dafny.Map
				_ = _2015_v31
				_2015_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2014_cf88, p3)
				var _2016_v32 _dafny.Sequence
				_ = _2016_v32
				_2016_v32 = _dafny.SeqOf(p3)
				var _2017_v33 _dafny.Set
				_ = _2017_v33
				_2017_v33 = _dafny.SetOf(_2016_v32)
				var _2018_v34 _dafny.Sequence
				_ = _2018_v34
				_2018_v34 = _dafny.SeqOf(_2017_v33, _2017_v33, _2017_v33)
				var _rhs326 _dafny.Int = (_this).Fm6(_2001_v19, globalState)
				_ = _rhs326
				var _rhs327 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (Companion_Default___.Fm0(p1, (_2009_v27).F13(), p0, (func() _dafny.Int {
						if (_2003_v21).Contains(!((_2009_v27).F13())) {
							return (_2003_v21).Multiplicity(!((_2009_v27).F13()))
						}
						return p2
					})(), globalState)).Cmp((_2015_v31).Cardinality()) == 0 {
						return _2001_v19
					}
					return _2001_v19
				})()).Cardinality())
				_ = _rhs327
				var _rhs328 bool = ((_2017_v33).Intersection(_2017_v33)).IsSubsetOf((func() _dafny.Set {
					if _2014_cf88 {
						return (_2018_v34).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2018_v34).Cardinality()))).Uint32()).(_dafny.Set)
					}
					return _2017_v33
				})())
				_ = _rhs328
				var _lhs256 *GlobalState = globalState
				_ = _lhs256
				var _lhs257 *GlobalState = globalState
				_ = _lhs257
				_lhs256.F7 = _rhs326
				_lhs257.F7 = _rhs327
				r2 = _rhs328
				r2 = p0
				var _index358 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_2009_v27).F14()), 0))
				_ = _index358
				((_2009_v27).F14()).ArraySet1((_2009_v27).F13(), (_index358).Int())
				var _index359 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(197), _dafny.ArrayLen(((_2009_v27).F14()), 0))
				_ = _index359
				((_2009_v27).F14()).ArraySet1(_2014_cf88, (_index359).Int())
				var _2019_v35 _dafny.Sequence
				_ = _2019_v35
				_2019_v35 = _dafny.SeqOf(_2001_v19)
				var _2020_v36 _dafny.Map
				_ = _2020_v36
				_2020_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC16_(p3), _2019_v35)
				_2020_v36 = (_2020_v36).Update(Companion_D7_.Create_DC16_(p3), _dafny.Companion_Sequence_.Concatenate(_2019_v35, _dafny.SeqOf(_2001_v19)))
			} else {
				var _2021___mcc_h7 _dafny.Array = _source28.Get_().(D17_DC44).Cf87
				_ = _2021___mcc_h7
				var _2022_cf87 _dafny.Array = _2021___mcc_h7
				_ = _2022_cf87
				var _2023_v37 _dafny.MultiSet
				_ = _2023_v37
				_2023_v37 = _dafny.MultiSetOf(_2000_cf5)
				var _2024_v38 _dafny.MultiSet
				_ = _2024_v38
				_2024_v38 = _2023_v37
				var _2025_v39 _dafny.Sequence
				_ = _2025_v39
				_2025_v39 = _dafny.SeqOf(_2024_v38, _2024_v38)
				var _2026_v40 _dafny.Set
				_ = _2026_v40
				_2026_v40 = _dafny.SetOf(_2001_v19)
				var _2027_v41 _dafny.Set
				_ = _2027_v41
				_2027_v41 = _dafny.SetOf(((_dafny.Zero).Minus(_dafny.IntOfUint32((_2025_v39).Cardinality()))).Cmp(_dafny.IntOfInt64(488)) < 0, ((_2026_v40).Cardinality()).Cmp(p2) > 0)
				_2027_v41 = _2027_v41
				var _2028_v42 _dafny.MultiSet
				_ = _2028_v42
				_2028_v42 = _dafny.MultiSetOf(p2)
				var _2029_v43 *C0
				_ = _2029_v43
				var _nw295 *C0 = New_C0_()
				_ = _nw295
				_nw295.Ctor__(true, _2005_v23, _2009_v27.F12(), (_2028_v42).IsProperSubsetOf(_2028_v42))
				_2029_v43 = _nw295
				var _2030_v44 _dafny.Sequence
				_ = _2030_v44
				_2030_v44 = _dafny.SeqOf(_1965_v0)
				_2030_v44 = _dafny.Companion_Sequence_.Concatenate(_2030_v44, _2030_v44)
				var _2031_v45 _dafny.Map
				_ = _2031_v45
				_2031_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2029_v43.F19, (_2009_v27).F13())
				_2031_v45 = (_2031_v45).Update(p0, (_2009_v27).F13())
			}
			r2 = (_2009_v27).F13()
			var _2032_v46 _dafny.Array
			_ = _2032_v46
			var _nw296 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw296
			_2032_v46 = _nw296
			var _index360 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2032_v46), 0))
			_ = _index360
			(_2032_v46).ArraySet1(_1976_v8, (_index360).Int())
			var _index361 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2032_v46), 0))
			_ = _index361
			var _rhs329 _dafny.CodePoint = _2005_v23
			_ = _rhs329
			var _rhs330 _dafny.Array = _1976_v8
			_ = _rhs330
			var _lhs258 _dafny.Array = _2032_v46
			_ = _lhs258
			var _lhs259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen((_2032_v46), 0))
			_ = _lhs259
			_2005_v23 = _rhs329
			(_lhs258).ArraySet1(_rhs330, (_lhs259).Int())
		}
		var _2033_v47 _dafny.Array
		_ = _2033_v47
		var _nw297 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(24))
		_ = _nw297
		_2033_v47 = _nw297
		for _iter64 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2033_v47), 0))); ; {
			_guard_loop_8, _ok64 := _iter64()
			if !_ok64 {
				break
			}
			var _2034_i2 _dafny.Int
			_2034_i2 = interface{}(_guard_loop_8).(_dafny.Int)
			if (true) && (((_2034_i2).Sign() != -1) && ((_2034_i2).Cmp(_dafny.ArrayLen((_2033_v47), 0)) < 0)) {
				(_2033_v47).ArraySet1(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0))), (_2034_i2).Int())
			}
		}
		var _2035_i3 _dafny.Int
		_ = _2035_i3
		_2035_i3 = _dafny.Zero
		{
			for (p3).Cmp(p3) <= 0 {
				{
					if (_2035_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_2035_i3 = (_2035_i3).Plus(_dafny.One)
					var _2036_v48 _dafny.MultiSet
					_ = _2036_v48
					_2036_v48 = _dafny.MultiSetOf(p3)
					_2036_v48 = _2036_v48
					var _2037_v49 _dafny.Sequence
					_ = _2037_v49
					_2037_v49 = _dafny.UnicodeSeqOfUtf8Bytes("lvqcjon")
					var _2038_v50 _dafny.CodePoint
					_ = _2038_v50
					_2038_v50 = _dafny.CodePoint('w')
					r1 = (_dafny.IntOfInt64(521)).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_2037_v49, (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_2037_v49).Cardinality()))).Uint32(), _2038_v50), _2037_v49)).Cardinality()))
					(globalState).F7 = (_dafny.Zero).Minus(p2)
					var _2039_v51 _dafny.Map
					_ = _2039_v51
					_2039_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p3, p3), !(true))
					r2 = (func() bool {
						if (_2039_v51).Contains(p2) {
							return (_2039_v51).Get(p2).(bool)
						}
						return p0
					})()
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _2040_v52 _dafny.Sequence
		_ = _2040_v52
		_2040_v52 = _dafny.UnicodeSeqOfUtf8Bytes("rxfde")
		var _hi13 _dafny.Int = _dafny.IntOfUint32((_2040_v52).Cardinality())
		_ = _hi13
		for _2041_i4 := (p3).Times(_dafny.IntOfInt64(325)); _2041_i4.Cmp(_hi13) < 0; _2041_i4 = _2041_i4.Plus(_dafny.One) {
			r0 = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(2))).Uint32(), func(coer106 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg106 _dafny.Int) interface{} {
					return coer106(arg106)
				}
			}(func(_2042_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('c')
			}))).Cardinality())
			var _2043_v53 D15
			_ = _2043_v53
			_2043_v53 = Companion_D15_.Create_DC39_(p3, _dafny.CodePoint('h'), _dafny.UnicodeSeqOfUtf8Bytes("uny"), p3, _dafny.IntOfInt64(80))
			var _2044_v54 _dafny.MultiSet
			_ = _2044_v54
			var _out15 _dafny.MultiSet
			_ = _out15
			_out15 = Companion_Default___.M0((_2043_v53).Dtor_cf79(), !(p0), globalState)
			_2044_v54 = _out15
			(globalState).F2 = p1
			r0 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(p2, _2041_i4))
		}
		var _2045_v55 _dafny.Set
		_ = _2045_v55
		_2045_v55 = _dafny.SetOf(p0)
		r0 = ((p3).Minus((_2045_v55).Cardinality())).Times(p3)
		r1 = p2
		r2 = p0
		return r0, r1, r2
	}
}

// End of class C7

// Definition of class C8
type C8 struct {
	F11 bool
}

func New_C8_() *C8 {
	_this := C8{}

	_this.F11 = false
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C8{}

func (_this *C8) Ctor__(f11 bool) {
	{
		(_this).F11 = f11
	}
}
func (_this *C8) Fm4(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) bool {
	{
		return _this.F11
	}
}
func (_this *C8) M2(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _2046_v0 _dafny.Sequence
		_ = _2046_v0
		_2046_v0 = _dafny.SeqOf(p2, _this.F11)
		var _hi14 _dafny.Int = p1
		_ = _hi14
		for _2047_i0 := (Companion_Default___.Fm0(_2046_v0, _this.F11, _this.F11, (_dafny.Zero).Minus(p0), globalState)).Times(_dafny.IntOfInt64(-964)); _2047_i0.Cmp(_hi14) < 0; _2047_i0 = _2047_i0.Plus(_dafny.One) {
			(_this).F11 = _this.F11
			var _2048_v1 _dafny.MultiSet
			_ = _2048_v1
			_2048_v1 = _dafny.MultiSetOf(p1)
			var _2049_v2 _dafny.Sequence
			_ = _2049_v2
			_2049_v2 = _dafny.UnicodeSeqOfUtf8Bytes("nmecnbdo")
			var _2050_v3 _dafny.Map
			_ = _2050_v3
			_2050_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).Fm4(_2046_v0, (_2048_v1).Cardinality(), _2049_v2, _this.F11, globalState), p2), p1)
			_2050_v3 = _2050_v3
			(globalState).F7 = p0
			if _this.F11 {
				(globalState).F7 = p1
				var _2051_v4 _dafny.Array
				_ = _2051_v4
				var _nw298 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw298
				_2051_v4 = _nw298
				_2051_v4 = _2051_v4
				var _2052_v5 _dafny.Map
				_ = _2052_v5
				_2052_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(_2049_v2, _2049_v2), Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_2046_v0).Cardinality()), (_dafny.Zero).Minus(_2047_i0)))
				_2052_v5 = (_2052_v5).Update(p2, (p1).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("npixpjow"), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("npixpjow")).Cardinality()))).Uint32(), Companion_Default___.Fm5(globalState))).Cardinality())))
				r0 = p2
				var _2053_v6 _dafny.Array
				_ = _2053_v6
				var _nw299 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw299
				_2053_v6 = _nw299
				var _2054_v7 _dafny.Sequence
				_ = _2054_v7
				_2054_v7 = _dafny.SeqOf(_2046_v0)
				var _index362 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_2053_v6), 0))
				_ = _index362
				(_2053_v6).ArraySet1(_2054_v7, (_index362).Int())
				var _2055_v8 _dafny.Sequence
				_ = _2055_v8
				_2055_v8 = _dafny.SeqOf(p0)
				var _index363 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_2053_v6), 0))
				_ = _index363
				(_2053_v6).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2054_v7, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2046_v0), _2054_v7), (Companion_Default___.SafeIndex((_2055_v8).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(279), _dafny.IntOfUint32((_2055_v8).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_2046_v0), _2054_v7)).Cardinality()))).Uint32(), _2046_v0)), (_index363).Int())
			} else {
				var _2056_v9 _dafny.Map
				_ = _2056_v9
				_2056_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _2047_i0)
				var _2057_v10 D6
				_ = _2057_v10
				_2057_v10 = Companion_D6_.Create_DC12_(_this.F11, _2049_v2, Companion_Default___.Fm0(_2046_v0, _this.F11, _this.F11, p0, globalState), _this.F11)
				var _2058_v11 D1
				_ = _2058_v11
				_2058_v11 = Companion_D1_.Create_DC1_(_2049_v2)
				var _2059_v12 _dafny.Array
				_ = _2059_v12
				var _nwElement0_67 bool = p2
				_ = _nwElement0_67
				var _nw300 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_67, nil, _dafny.IntOfInt64(17))
				_ = _nw300
				(_nw300).ArraySet1(_nwElement0_67, 0)
				(_nw300).ArraySet1(p2, 1)
				(_nw300).ArraySet1(true, 2)
				(_nw300).ArraySet1(_this.F11, 3)
				(_nw300).ArraySet1(_this.F11, 4)
				(_nw300).ArraySet1(_this.F11, 5)
				(_nw300).ArraySet1(_this.F11, 6)
				(_nw300).ArraySet1(!(_this.F11), 7)
				(_nw300).ArraySet1(p2, 8)
				(_nw300).ArraySet1(p2, 9)
				(_nw300).ArraySet1(_this.F11, 10)
				(_nw300).ArraySet1(p2, 11)
				(_nw300).ArraySet1(p2, 12)
				(_nw300).ArraySet1(_this.F11, 13)
				(_nw300).ArraySet1(p2, 14)
				(_nw300).ArraySet1(!(true), 15)
				(_nw300).ArraySet1(p2, 16)
				_2059_v12 = _nw300
				var _2060_v13 _dafny.Set
				_ = _2060_v13
				_2060_v13 = _dafny.SetOf(p0)
				var _pat_let_tv60 = p2
				_ = _pat_let_tv60
				var _2061_v14 *C2
				_ = _2061_v14
				var _nw301 *C2 = New_C2_()
				_ = _nw301
				_nw301.Ctor__((_dafny.Zero).Minus(((_dafny.SetOf(_2049_v2)).Cardinality()).Times((_2056_v9).Cardinality())), (func() bool {
					if (_2057_v10).Dtor_cf21() {
						return _this.F11
					}
					return _this.F11
				})(), _2058_v11, Companion_Default___.Fm30(p1, globalState), _2059_v12, (_2060_v13).Difference(_dafny.SetOf(p1)), func(_pat_let65_0 D1) D1 {
					return func(_2062_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let66_0 bool) D1 {
							return func(_2063_dt__update_hcf4_h0 bool) D1 {
								return func(_pat_let67_0 _dafny.Int) D1 {
									return func(_2064_dt__update_hcf3_h0 _dafny.Int) D1 {
										return Companion_D1_.Create_DC2_((_2062_dt__update__tmp_h0).Dtor_cf2(), _2064_dt__update_hcf3_h0, _2063_dt__update_hcf4_h0)
									}(_pat_let67_0)
								}(_2047_i0)
							}(_pat_let66_0)
						}(_pat_let_tv60)
					}(_pat_let65_0)
				}(Companion_Default___.Fm37(p1, p0, _dafny.IntOfUint32((_2049_v2).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("knl")).Cardinality()), globalState)), p2)
				_2061_v14 = _nw301
				(_this).F11 = !(_this.F11)
				_2059_v12 = _2059_v12
				(globalState).F7 = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(650), (p1).Plus((_2061_v14).F21()))
				var _2065_v15 _dafny.Set
				_ = _2065_v15
				_2065_v15 = _dafny.SetOf(p2)
				var _2066_v16 *C2
				_ = _2066_v16
				var _nw302 *C2 = New_C2_()
				_ = _nw302
				_nw302.Ctor__((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if p2 {
						return _2049_v2
					}
					return _dafny.UnicodeSeqOfUtf8Bytes("le")
				})()).Cardinality())), (_this).Fm4(_2046_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajkrilfy")).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("dli"), !((_2061_v14).F22()), globalState), Companion_D1_.Create_DC1_(_2049_v2), !(Companion_Default___.Fm17(globalState)), _2059_v12, (_2060_v13).Intersection(_2060_v13), Companion_D1_.Create_DC2_(Companion_Default___.Fm0(_dafny.SeqOf(_this.F11, (_2061_v14).F22()), p2, _this.F11, (_2065_v15).Cardinality(), globalState), (_2061_v14).F21(), (_2061_v14).F22()), (_2061_v14).F22())
				_2066_v16 = _nw302
			}
		}
		var _2067_v17 _dafny.MultiSet
		_ = _2067_v17
		_2067_v17 = _dafny.MultiSetOf(_dafny.MultiSetOf(_this.F11, p2, _this.F11), _dafny.MultiSetOf(_this.F11, _this.F11))
		var _2068_v18 _dafny.MultiSet
		_ = _2068_v18
		_2068_v18 = _dafny.MultiSetOf(p2)
		var _2069_v19 _dafny.Set
		_ = _2069_v19
		_2069_v19 = _dafny.SetOf(p1, p0)
		var _2070_v20 _dafny.Sequence
		_ = _2070_v20
		_2070_v20 = _dafny.SeqOf(_dafny.SeqOf((func() _dafny.Int {
			if (_2067_v17).Contains(_2068_v18) {
				return (_2067_v17).Multiplicity(_2068_v18)
			}
			return p0
		})(), Companion_Default___.Fm0(_2046_v0, _this.F11, p2, (_2069_v19).Cardinality(), globalState), p0, p1, p0))
		var _2071_v21 _dafny.Sequence
		_ = _2071_v21
		_2071_v21 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_2046_v0).Cardinality()))
		var _2072_v22 _dafny.CodePoint
		_ = _2072_v22
		_2072_v22 = _dafny.CodePoint('a')
		var _2073_v23 _dafny.Sequence
		_ = _2073_v23
		_2073_v23 = _dafny.UnicodeSeqOfUtf8Bytes("vrdiples")
		var _2074_v24 D1
		_ = _2074_v24
		_2074_v24 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_2073_v23).Cardinality()), p1, _this.F11)
		var _2075_v25 *C1
		_ = _2075_v25
		var _nw303 *C1 = New_C1_()
		_ = _nw303
		_nw303.Ctor__(_2072_v22, _2074_v24, _this.F11)
		_2075_v25 = _nw303
		var _2076_v26 _dafny.MultiSet
		_ = _2076_v26
		_2076_v26 = _dafny.MultiSetOf(_2075_v25, _2075_v25)
		var _2077_v27 _dafny.Array
		_ = _2077_v27
		var _nwElement0_68 _dafny.Int = _dafny.IntOfInt64(791)
		_ = _nwElement0_68
		var _nw304 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_68, nil, _dafny.IntOfInt64(25))
		_ = _nw304
		(_nw304).ArraySet1(_nwElement0_68, 0)
		(_nw304).ArraySet1(p1, 1)
		(_nw304).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("w")).Cardinality()), 2)
		(_nw304).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), 3)
		(_nw304).ArraySet1(_dafny.IntOfUint32((_2070_v20).Cardinality()), 4)
		(_nw304).ArraySet1(p1, 5)
		(_nw304).ArraySet1(p0, 6)
		(_nw304).ArraySet1((_dafny.Zero).Minus(p1), 7)
		(_nw304).ArraySet1(p1, 8)
		(_nw304).ArraySet1(_dafny.IntOfInt64(341), 9)
		(_nw304).ArraySet1(p0, 10)
		(_nw304).ArraySet1(p1, 11)
		(_nw304).ArraySet1(p1, 12)
		(_nw304).ArraySet1(p1, 13)
		(_nw304).ArraySet1(p0, 14)
		(_nw304).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus(p1), p1)).Cardinality()), 15)
		(_nw304).ArraySet1(Companion_Default___.Fm0(_2046_v0, p2, _this.F11, p0, globalState), 16)
		(_nw304).ArraySet1(p1, 17)
		(_nw304).ArraySet1(p0, 18)
		(_nw304).ArraySet1(p1, 19)
		(_nw304).ArraySet1(p1, 20)
		(_nw304).ArraySet1(p1, 21)
		(_nw304).ArraySet1((_2071_v21).Select((Companion_Default___.SafeIndex((_2076_v26).Cardinality(), _dafny.IntOfUint32((_2071_v21).Cardinality()))).Uint32()).(_dafny.Int), 22)
		(_nw304).ArraySet1(p1, 23)
		(_nw304).ArraySet1(p0, 24)
		_2077_v27 = _nw304
		var _2078_v28 _dafny.Map
		_ = _2078_v28
		_2078_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2077_v27)
		var _hi15 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2071_v21, _2071_v21)).Cardinality())
		_ = _hi15
		for _2079_i1 := ((func() _dafny.Map {
			if _this.F11 {
				return _2078_v28
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_2068_v18).Contains(p2) {
					return (_2068_v18).Multiplicity(p2)
				}
				return p1
			})(), _2077_v27)
		})()).Cardinality(); _2079_i1.Cmp(_hi15) < 0; _2079_i1 = _2079_i1.Plus(_dafny.One) {
			(globalState).F7 = _2079_i1
			(globalState).F7 = p1
			var _2080_v29 _dafny.MultiSet
			_ = _2080_v29
			_2080_v29 = _dafny.MultiSetOf(_2072_v22, _2072_v22, (func() _dafny.CodePoint {
				if _this.F11 {
					return _2075_v25.F23
				}
				return _2075_v25.F23
			})(), _dafny.CodePoint('f'))
			var _2081_v30 _dafny.Sequence
			_ = _2081_v30
			_2081_v30 = _dafny.SeqOf(_2080_v29)
			_2080_v29 = (_2081_v30).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.SetOf(_this.F11)).Cardinality())), _dafny.IntOfInt64(980))), _dafny.IntOfUint32((_2081_v30).Cardinality()))).Uint32()).(_dafny.MultiSet)
			var _2082_v31 _dafny.MultiSet
			_ = _2082_v31
			_2082_v31 = _dafny.MultiSetOf(_dafny.IntOfInt64(-531))
			(_this).F11 = (_this.F11) == ((_2082_v31).Equals(_2082_v31))
		}
		var _2083_v32 _dafny.Set
		_ = _2083_v32
		_2083_v32 = _dafny.SetOf(_this.F11, _this.F11)
		var _2084_v33 _dafny.Map
		_ = _2084_v33
		_2084_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2083_v32, _dafny.Companion_Sequence_.Concatenate(_2071_v21, _2071_v21))
		_2084_v33 = (_2084_v33).Update(_2083_v32, _dafny.SeqOf((_dafny.SetOf(p0)).Cardinality(), p0, _dafny.IntOfInt64(128), (_2071_v21).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(404), _dafny.IntOfUint32((_2071_v21).Cardinality()))).Uint32()).(_dafny.Int), p0))
		(_this).F11 = false
		var _2085_v34 _dafny.Map
		_ = _2085_v34
		_2085_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F11), p2)
		var _2086_v35 _dafny.Sequence
		_ = _2086_v35
		_2086_v35 = _dafny.SeqOf(_2085_v34, _2085_v34, _2085_v34)
		var _2087_v36 _dafny.Map
		_ = _2087_v36
		_2087_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _2086_v35)
		_2086_v35 = (func() _dafny.Sequence {
			if (_2087_v36).Contains(p1) {
				return (_2087_v36).Get(p1).(_dafny.Sequence)
			}
			return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_2085_v34, _2085_v34), (Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_dafny.SeqOf(_2085_v34, _2085_v34)).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F11))
		})()
		var _2088_v37 D1
		_ = _2088_v37
		_2088_v37 = Companion_D1_.Create_DC1_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(715))).Uint32(), func(coer107 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg107 _dafny.Int) interface{} {
				return coer107(arg107)
			}
		}((func(_2089_v22 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_2090_i2 _dafny.Int) _dafny.CodePoint {
				return _2089_v22
			}
		})(_2072_v22))))
		var _2091_v38 _dafny.Array
		_ = _2091_v38
		var _len0_49 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_49
		var _nw305 _dafny.Array
		_ = _nw305
		if _len0_49.Cmp(_dafny.Zero) == 0 {
			_nw305 = _dafny.NewArray(_len0_49)
		} else {
			var _init49 func(_dafny.Int) bool = (func(_2092_p2 bool) func(_dafny.Int) bool {
				return func(_2093_i3 _dafny.Int) bool {
					return _2092_p2
				}
			})(p2)
			_ = _init49
			var _element0_49 = _init49(_dafny.Zero)
			_ = _element0_49
			_nw305 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
			(_nw305).ArraySet1(_element0_49, 0)
			var _nativeLen0_49 = (_len0_49).Int()
			_ = _nativeLen0_49
			for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
				(_nw305).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
			}
		}
		_2091_v38 = _nw305
		var _2094_v39 *C3
		_ = _2094_v39
		var _nw306 *C3 = New_C3_()
		_ = _nw306
		_nw306.Ctor__(true, Companion_Default___.Fm62(!(false), _2073_v23, globalState), p2, _2091_v38, _2069_v19, _2074_v24, _this.F11)
		_2094_v39 = _nw306
		var _2095_v40 _dafny.MultiSet
		_ = _2095_v40
		_2095_v40 = _dafny.MultiSetOf(_2094_v39)
		var _2096_v41 _dafny.Sequence
		_ = _2096_v41
		_2096_v41 = _dafny.SeqOf(_2094_v39)
		var _2097_v42 _dafny.Set
		_ = _2097_v42
		_2097_v42 = _dafny.SetOf(_dafny.MultiSetFromSeq(_2096_v41), _dafny.MultiSetFromSeq(_dafny.SeqOf(_2094_v39)))
		var _2098_v43 *C2
		_ = _2098_v43
		var _nw307 *C2 = New_C2_()
		_ = _nw307
		_nw307.Ctor__(p1, p2, _2088_v37, (_2097_v42).Contains(_2095_v40), _2091_v38, _dafny.SetOf(p0), _2074_v24, (_2094_v39).F18())
		_2098_v43 = _nw307
		r0 = (_2098_v43).F22()
		return r0
	}
}
func (_this *C8) M3(p0 _dafny.Int, globalState *GlobalState) (_dafny.Map, D2) {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var r1 D2 = Companion_D2_.Default()
		_ = r1
		if true {
			(globalState).F7 = _dafny.IntOfInt64(56)
			var _2099_v0 _dafny.Sequence
			_ = _2099_v0
			_2099_v0 = _dafny.UnicodeSeqOfUtf8Bytes("yhesaykny")
			var _2100_v1 _dafny.Array
			_ = _2100_v1
			var _len0_50 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_50
			var _nw308 _dafny.Array
			_ = _nw308
			if _len0_50.Cmp(_dafny.Zero) == 0 {
				_nw308 = _dafny.NewArray(_len0_50)
			} else {
				var _init50 func(_dafny.Int) _dafny.Map = func(_2101_i0 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_this.F11, _this.F11, _this.F11, _this.F11), _this.F11)
				}
				_ = _init50
				var _element0_50 = _init50(_dafny.Zero)
				_ = _element0_50
				_nw308 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
				(_nw308).ArraySet1(_element0_50, 0)
				var _nativeLen0_50 = (_len0_50).Int()
				_ = _nativeLen0_50
				for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
					(_nw308).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
				}
			}
			_2100_v1 = _nw308
			var _2102_v2 _dafny.Set
			_ = _2102_v2
			_2102_v2 = _dafny.SetOf(_this.F11, !(_this.F11), !(_this.F11))
			var _2103_v3 _dafny.Map
			_ = _2103_v3
			_2103_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2102_v2, _this.F11)
			var _index364 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_2100_v1), 0))
			_ = _index364
			(_2100_v1).ArraySet1(_2103_v3, (_index364).Int())
			var _2104_v5 _dafny.Map
			_ = _2104_v5
			_2104_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2102_v2, _dafny.IntOfInt64(153))
			var _index365 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_2100_v1), 0))
			_ = _index365
			var _rhs331 bool = _this.F11
			_ = _rhs331
			var _rhs332 _dafny.Sequence = _2099_v0
			_ = _rhs332
			var _rhs333 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2102_v2, _this.F11)).Merge(func() _dafny.Map {
				var _coll54 = _dafny.NewMapBuilder()
				_ = _coll54
				for _iter65 := _dafny.Iterate((_2104_v5).Keys().Elements()); ; {
					_compr_54, _ok65 := _iter65()
					if !_ok65 {
						break
					}
					var _2105_v4 _dafny.Set
					_2105_v4 = interface{}(_compr_54).(_dafny.Set)
					if (_2104_v5).Contains(_2105_v4) {
						_coll54.Add(_2105_v4, !(_this.F11))
					}
				}
				return _coll54.ToMap()
			}())).Merge((_2103_v3).Update(_2102_v2, _this.F11))
			_ = _rhs333
			var _rhs334 bool = true
			_ = _rhs334
			var _rhs335 bool = false
			_ = _rhs335
			var _lhs260 *C8 = _this
			_ = _lhs260
			var _lhs261 _dafny.Array = _2100_v1
			_ = _lhs261
			var _lhs262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_2100_v1), 0))
			_ = _lhs262
			var _lhs263 *C8 = _this
			_ = _lhs263
			var _lhs264 *C8 = _this
			_ = _lhs264
			_lhs260.F11 = _rhs331
			_2099_v0 = _rhs332
			(_lhs261).ArraySet1(_rhs333, (_lhs262).Int())
			_lhs263.F11 = _rhs334
			_lhs264.F11 = _rhs335
			var _2106_v6 _dafny.Sequence
			_ = _2106_v6
			_2106_v6 = _dafny.SeqOf(_this.F11)
			var _2107_v7 _dafny.Map
			_ = _2107_v7
			_2107_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2106_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2106_v6).Cardinality()))).Uint32()).(bool), _this.F11)
			var _2108_v8 _dafny.Map
			_ = _2108_v8
			_2108_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, p0)
			var _2109_v9 _dafny.Map
			_ = _2109_v9
			_2109_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm18((_2107_v7).Cardinality(), _dafny.UnicodeSeqOfUtf8Bytes("narktrxe"), _this.F11, _2108_v8, globalState), _this.F11)
			var _2110_v10 _dafny.CodePoint
			_ = _2110_v10
			_2110_v10 = _dafny.CodePoint('k')
			_2109_v9 = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2110_v10, _this.F11)).Merge(_2109_v9)
			var _2111_v11 D6
			_ = _2111_v11
			_2111_v11 = Companion_D6_.Create_DC12_(_this.F11, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_2099_v0, _2099_v0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2099_v0, _2099_v0)).Cardinality()))).Uint32(), _2110_v10), p0, _this.F11)
			var _source29 D6 = _2111_v11
			_ = _source29
			if _source29.Is_DC12() {
				var _2112___mcc_h0 bool = _source29.Get_().(D6_DC12).Cf18
				_ = _2112___mcc_h0
				var _2113___mcc_h1 _dafny.Sequence = _source29.Get_().(D6_DC12).Cf19
				_ = _2113___mcc_h1
				var _2114___mcc_h2 _dafny.Int = _source29.Get_().(D6_DC12).Cf20
				_ = _2114___mcc_h2
				var _2115___mcc_h3 bool = _source29.Get_().(D6_DC12).Cf21
				_ = _2115___mcc_h3
				var _2116_cf21 bool = _2115___mcc_h3
				_ = _2116_cf21
				var _2117_cf20 _dafny.Int = _2114___mcc_h2
				_ = _2117_cf20
				var _2118_cf19 _dafny.Sequence = _2113___mcc_h1
				_ = _2118_cf19
				var _2119_cf18 bool = _2112___mcc_h0
				_ = _2119_cf18
				var _2120_v12 _dafny.Array
				_ = _2120_v12
				var _len0_51 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_51
				var _nw309 _dafny.Array
				_ = _nw309
				if _len0_51.Cmp(_dafny.Zero) == 0 {
					_nw309 = _dafny.NewArray(_len0_51)
				} else {
					var _init51 func(_dafny.Int) _dafny.Int = (func(_2121_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2122_i1 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_2122_i1, (_dafny.Zero).Minus(_2121_p0))
						}
					})(p0)
					_ = _init51
					var _element0_51 = _init51(_dafny.Zero)
					_ = _element0_51
					_nw309 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
					(_nw309).ArraySet1(_element0_51, 0)
					var _nativeLen0_51 = (_len0_51).Int()
					_ = _nativeLen0_51
					for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
						(_nw309).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
					}
				}
				_2120_v12 = _nw309
				_2120_v12 = _2120_v12
				var _index366 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_2120_v12), 0))
				_ = _index366
				(_2120_v12).ArraySet1(_2117_cf20, (_index366).Int())
				var _index367 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(613), _dafny.ArrayLen((_2120_v12), 0))
				_ = _index367
				(_2120_v12).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2099_v0).Cardinality()), _2117_cf20)), (_index367).Int())
				var _2123_v13 _dafny.Array
				_ = _2123_v13
				var _nw310 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(12))
				_ = _nw310
				_2123_v13 = _nw310
				var _nw311 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(14))
				_ = _nw311
				_2123_v13 = _nw311
				var _2124_v14 _dafny.Array
				_ = _2124_v14
				var _nw312 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
				_ = _nw312
				_2124_v14 = _nw312
				_2124_v14 = _2124_v14
			} else if _source29.Is_DC13() {
				var _2125___mcc_h4 bool = _source29.Get_().(D6_DC13).Cf22
				_ = _2125___mcc_h4
				var _2126___mcc_h5 bool = _source29.Get_().(D6_DC13).Cf23
				_ = _2126___mcc_h5
				var _2127___mcc_h6 _dafny.Int = _source29.Get_().(D6_DC13).Cf24
				_ = _2127___mcc_h6
				var _2128___mcc_h7 _dafny.Int = _source29.Get_().(D6_DC13).Cf25
				_ = _2128___mcc_h7
				var _2129___mcc_h8 bool = _source29.Get_().(D6_DC13).Cf26
				_ = _2129___mcc_h8
				var _2130_cf26 bool = _2129___mcc_h8
				_ = _2130_cf26
				var _2131_cf25 _dafny.Int = _2128___mcc_h7
				_ = _2131_cf25
				var _2132_cf24 _dafny.Int = _2127___mcc_h6
				_ = _2132_cf24
				var _2133_cf23 bool = _2126___mcc_h5
				_ = _2133_cf23
				var _2134_cf22 bool = _2125___mcc_h4
				_ = _2134_cf22
				var _2135_v15 _dafny.MultiSet
				_ = _2135_v15
				_2135_v15 = _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt(_2131_cf25, p0), ((_dafny.Zero).Minus(_2131_cf25)).Plus(_2131_cf25), _2132_cf24, (p0).Minus(p0))
				_2135_v15 = _2135_v15
				var _2136_v16 _dafny.Map
				_ = _2136_v16
				_2136_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2106_v6, _2134_cf22)
				var _2137_v17 D12
				_ = _2137_v17
				_2137_v17 = Companion_D12_.Create_DC29_(_2136_v16)
				var _2138_v18 _dafny.Map
				_ = _2138_v18
				_2138_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2137_v17, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("swnt"), (Companion_Default___.SafeIndex(_2131_cf25, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("swnt")).Cardinality()))).Uint32(), _2110_v10))
				_2138_v18 = (_2138_v18).Update(_2137_v17, (func() _dafny.Sequence {
					if _2130_cf26 {
						return _dafny.UnicodeSeqOfUtf8Bytes("dii")
					}
					return _2099_v0
				})())
				var _2139_v19 *C7
				_ = _2139_v19
				var _nw313 *C7 = New_C7_()
				_ = _nw313
				_nw313.Ctor__()
				_2139_v19 = _nw313
				var _2140_v20 _dafny.Map
				_ = _2140_v20
				_2140_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2132_cf24).Plus(_2131_cf25), Companion_Default___.SafeModuloInt(p0, p0))
				_2140_v20 = (_2140_v20).Update(p0, p0)
			} else if _source29.Is_DC11() {
				var _2141___mcc_h9 _dafny.MultiSet = _source29.Get_().(D6_DC11).Cf17
				_ = _2141___mcc_h9
				var _2142_cf17 _dafny.MultiSet = _2141___mcc_h9
				_ = _2142_cf17
				var _2143_v22 _dafny.Sequence
				_ = _2143_v22
				_2143_v22 = _dafny.SeqOf(p0)
				var _2144_v23 _dafny.MultiSet
				_ = _2144_v23
				_2144_v23 = _dafny.MultiSetOf(func() _dafny.Map {
					var _coll55 = _dafny.NewMapBuilder()
					_ = _coll55
					for _iter66 := _dafny.Iterate((_2143_v22).Elements()); ; {
						_compr_55, _ok66 := _iter66()
						if !_ok66 {
							break
						}
						var _2145_v21 _dafny.Int
						_2145_v21 = interface{}(_compr_55).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_2143_v22, _2145_v21) {
							_coll55.Add((_2145_v21).Plus(p0), _this.F11)
						}
					}
					return _coll55.ToMap()
				}())
				var _2146_v24 _dafny.Map
				_ = _2146_v24
				_2146_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F11)
				var _2147_v25 bool
				_ = _2147_v25
				var _out16 bool
				_ = _out16
				_out16 = (_this).M2(((_2144_v23).Difference(_dafny.MultiSetOf(_2146_v24))).Cardinality(), _dafny.IntOfUint32((_2099_v0).Cardinality()), !(_this.F11), globalState)
				_2147_v25 = _out16
				var _2148_v26 _dafny.Array
				_ = _2148_v26
				var _nw314 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
				_ = _nw314
				_2148_v26 = _nw314
				var _2149_v27 _dafny.Array
				_ = _2149_v27
				var _nwElement0_69 _dafny.Int = p0
				_ = _nwElement0_69
				var _nw315 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_69, nil, _dafny.IntOfInt64(3))
				_ = _nw315
				(_nw315).ArraySet1(_nwElement0_69, 0)
				(_nw315).ArraySet1(p0, 1)
				(_nw315).ArraySet1(p0, 2)
				_2149_v27 = _nw315
				var _index368 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2148_v26), 0))
				_ = _index368
				(_2148_v26).ArraySet1(_2149_v27, (_index368).Int())
				var _index369 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2148_v26), 0))
				_ = _index369
				var _rhs336 bool = !(_2147_v25)
				_ = _rhs336
				var _rhs337 _dafny.Array = _2149_v27
				_ = _rhs337
				var _lhs265 *C8 = _this
				_ = _lhs265
				var _lhs266 _dafny.Array = _2148_v26
				_ = _lhs266
				var _lhs267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2148_v26), 0))
				_ = _lhs267
				_lhs265.F11 = _rhs336
				(_lhs266).ArraySet1(_rhs337, (_lhs267).Int())
				_2143_v22 = _dafny.SeqOf(_dafny.IntOfInt64(533), _dafny.IntOfInt64(549))
				var _2150_v28 _dafny.Array
				_ = _2150_v28
				var _len0_52 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_52
				var _nw316 _dafny.Array
				_ = _nw316
				if _len0_52.Cmp(_dafny.Zero) == 0 {
					_nw316 = _dafny.NewArray(_len0_52)
				} else {
					var _init52 func(_dafny.Int) _dafny.Set = (func(_2151_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
						return func(_2152_i2 _dafny.Int) _dafny.Set {
							return _dafny.SetOf(_dafny.IntOfUint32((_2151_v0).Cardinality()))
						}
					})(_2099_v0)
					_ = _init52
					var _element0_52 = _init52(_dafny.Zero)
					_ = _element0_52
					_nw316 = _dafny.NewArrayFromExample(_element0_52, nil, _len0_52)
					(_nw316).ArraySet1(_element0_52, 0)
					var _nativeLen0_52 = (_len0_52).Int()
					_ = _nativeLen0_52
					for _i0_52 := 1; _i0_52 < _nativeLen0_52; _i0_52++ {
						(_nw316).ArraySet1(_init52(_dafny.IntOf(_i0_52)), _i0_52)
					}
				}
				_2150_v28 = _nw316
				_2150_v28 = _2150_v28
			} else {
				var _2153___mcc_h10 D6 = _source29.Get_().(D6_DC14).Cf27
				_ = _2153___mcc_h10
				var _2154_cf27 D6 = _2153___mcc_h10
				_ = _2154_cf27
				var _2155_v29 D1
				_ = _2155_v29
				_2155_v29 = Companion_D1_.Create_DC1_(_dafny.Companion_Sequence_.Update(_2099_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-526), _dafny.IntOfUint32((_2099_v0).Cardinality()))).Uint32(), _dafny.CodePoint('s')))
				var _2156_v30 _dafny.Array
				_ = _2156_v30
				var _len0_53 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_53
				var _nw317 _dafny.Array
				_ = _nw317
				if _len0_53.Cmp(_dafny.Zero) == 0 {
					_nw317 = _dafny.NewArray(_len0_53)
				} else {
					var _init53 func(_dafny.Int) bool = func(_2157_i3 _dafny.Int) bool {
						return !(_this.F11)
					}
					_ = _init53
					var _element0_53 = _init53(_dafny.Zero)
					_ = _element0_53
					_nw317 = _dafny.NewArrayFromExample(_element0_53, nil, _len0_53)
					(_nw317).ArraySet1(_element0_53, 0)
					var _nativeLen0_53 = (_len0_53).Int()
					_ = _nativeLen0_53
					for _i0_53 := 1; _i0_53 < _nativeLen0_53; _i0_53++ {
						(_nw317).ArraySet1(_init53(_dafny.IntOf(_i0_53)), _i0_53)
					}
				}
				_2156_v30 = _nw317
				var _2158_v31 _dafny.Set
				_ = _2158_v31
				_2158_v31 = _dafny.SetOf(p0)
				var _2159_v32 *C6
				_ = _2159_v32
				var _nw318 *C6 = New_C6_()
				_ = _nw318
				_nw318.Ctor__(_2155_v29, _this.F11, _2156_v30, _2158_v31, Companion_D1_.Create_DC2_(p0, p0, _this.F11), _this.F11)
				_2159_v32 = _nw318
				var _2160_v33 _dafny.Map
				_ = _2160_v33
				_2160_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2159_v32, _2107_v7)
				var _2161_v34 _dafny.MultiSet
				_ = _2161_v34
				_2161_v34 = _dafny.MultiSetOf(_this.F11, (_2159_v32).Fm11(globalState))
				var _2162_v35 _dafny.Map
				_ = _2162_v35
				_2162_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2161_v34).Cardinality(), p0)
				var _2163_v36 _dafny.Sequence
				_ = _2163_v36
				_2163_v36 = _dafny.SeqOf((_2162_v35).Cardinality(), p0, _dafny.IntOfInt64(352), p0)
				var _2164_v37 _dafny.Array
				_ = _2164_v37
				var _nwElement0_70 _dafny.Int = _dafny.IntOfUint32((_2099_v0).Cardinality())
				_ = _nwElement0_70
				var _nw319 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_70, nil, _dafny.IntOfInt64(13))
				_ = _nw319
				(_nw319).ArraySet1(_nwElement0_70, 0)
				(_nw319).ArraySet1((_2160_v33).Cardinality(), 1)
				(_nw319).ArraySet1(Companion_Default___.SafeModuloInt(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)).Cardinality()), 2)
				(_nw319).ArraySet1(p0, 3)
				(_nw319).ArraySet1(_dafny.IntOfUint32((_2163_v36).Cardinality()), 4)
				(_nw319).ArraySet1(p0, 5)
				(_nw319).ArraySet1((_dafny.IntOfUint32((_2099_v0).Cardinality())).Minus(_dafny.IntOfInt64(983)), 6)
				(_nw319).ArraySet1(_dafny.IntOfUint32((_2099_v0).Cardinality()), 7)
				(_nw319).ArraySet1(p0, 8)
				(_nw319).ArraySet1(p0, 9)
				(_nw319).ArraySet1(p0, 10)
				(_nw319).ArraySet1(((_2107_v7).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, true))).Cardinality(), 11)
				(_nw319).ArraySet1(p0, 12)
				_2164_v37 = _nw319
				var _2165_v38 D2
				_ = _2165_v38
				_2165_v38 = Companion_D2_.Create_DC4_(_2164_v37, _this.F11, p0, _this.F11, _2156_v30)
				var _index370 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_2164_v37), 0))
				_ = _index370
				(_2164_v37).ArraySet1((func() _dafny.Int {
					if _this.F11 {
						return (func() _dafny.Int {
							if (_2162_v35).Contains(p0) {
								return (_2162_v35).Get(p0).(_dafny.Int)
							}
							return p0
						})()
					}
					return (_2165_v38).Dtor_cf8()
				})(), (_index370).Int())
				var _index371 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_2164_v37), 0))
				_ = _index371
				(_2164_v37).ArraySet1(((p0).Times(p0)).Minus(p0), (_index371).Int())
				(globalState).F7 = (_2164_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_2164_v37), 0))).Int()).(_dafny.Int)
				(globalState).F7 = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_2099_v0).Cardinality()), (_2164_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_2164_v37), 0))).Int()).(_dafny.Int))
				var _index372 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_2164_v37), 0))
				_ = _index372
				(_2164_v37).ArraySet1(p0, (_index372).Int())
			}
			var _2166_v39 _dafny.Array
			_ = _2166_v39
			var _len0_54 _dafny.Int = _dafny.One
			_ = _len0_54
			var _nw320 _dafny.Array
			_ = _nw320
			if _len0_54.Cmp(_dafny.Zero) == 0 {
				_nw320 = _dafny.NewArray(_len0_54)
			} else {
				var _init54 func(_dafny.Int) _dafny.Sequence = (func(_2167_v0 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_2168_i4 _dafny.Int) _dafny.Sequence {
						return _2167_v0
					}
				})(_2099_v0)
				_ = _init54
				var _element0_54 = _init54(_dafny.Zero)
				_ = _element0_54
				_nw320 = _dafny.NewArrayFromExample(_element0_54, nil, _len0_54)
				(_nw320).ArraySet1(_element0_54, 0)
				var _nativeLen0_54 = (_len0_54).Int()
				_ = _nativeLen0_54
				for _i0_54 := 1; _i0_54 < _nativeLen0_54; _i0_54++ {
					(_nw320).ArraySet1(_init54(_dafny.IntOf(_i0_54)), _i0_54)
				}
			}
			_2166_v39 = _nw320
			_2166_v39 = _2166_v39
		} else {
			var _2169_v40 _dafny.Sequence
			_ = _2169_v40
			_2169_v40 = _dafny.UnicodeSeqOfUtf8Bytes("dpqp")
			var _2170_v41 _dafny.CodePoint
			_ = _2170_v41
			_2170_v41 = _dafny.CodePoint('y')
			var _2171_v42 _dafny.Array
			_ = _2171_v42
			var _nw321 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw321
			_2171_v42 = _nw321
			var _2172_v43 D12
			_ = _2172_v43
			_2172_v43 = Companion_D12_.Create_DC30_(_2171_v42, _this.F11)
			var _2173_v44 _dafny.Map
			_ = _2173_v44
			_2173_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _2174_v45 _dafny.Map
			_ = _2174_v45
			_2174_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _this.F11)
			var _2175_v46 _dafny.Array
			_ = _2175_v46
			var _nwElement0_71 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("sq")
			_ = _nwElement0_71
			var _nw322 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_71, nil, _dafny.IntOfInt64(27))
			_ = _nw322
			(_nw322).ArraySet1(_nwElement0_71, 0)
			(_nw322).ArraySet1(_2169_v40, 1)
			(_nw322).ArraySet1(_2169_v40, 2)
			(_nw322).ArraySet1(_2169_v40, 3)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2169_v40, _dafny.UnicodeSeqOfUtf8Bytes("ugei")), 4)
			(_nw322).ArraySet1(_2169_v40, 5)
			(_nw322).ArraySet1(_2169_v40, 6)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2169_v40, _dafny.UnicodeSeqOfUtf8Bytes("vjrdhd")), 7)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("hd"), _2169_v40), 8)
			(_nw322).ArraySet1(_2169_v40, 9)
			(_nw322).ArraySet1(_2169_v40, 10)
			(_nw322).ArraySet1(_2169_v40, 11)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Update(_2169_v40, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2169_v40).Cardinality()))).Uint32(), _2170_v41), 12)
			(_nw322).ArraySet1(_2169_v40, 13)
			(_nw322).ArraySet1(_2169_v40, 14)
			(_nw322).ArraySet1(_2169_v40, 15)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("scbj"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("scbj")).Cardinality()))).Uint32(), _2170_v41), 16)
			(_nw322).ArraySet1(_2169_v40, 17)
			(_nw322).ArraySet1(_2169_v40, 18)
			(_nw322).ArraySet1((func() _dafny.Sequence {
				if _this.F11 {
					return _dafny.UnicodeSeqOfUtf8Bytes("bbwe")
				}
				return _2169_v40
			})(), 19)
			(_nw322).ArraySet1(_2169_v40, 20)
			(_nw322).ArraySet1((func() _dafny.Sequence {
				if _this.F11 {
					return _dafny.Companion_Sequence_.Update(_2169_v40, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2169_v40).Cardinality()))).Uint32(), _2170_v41)
				}
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(34))).Uint32(), func(coer108 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg108 _dafny.Int) interface{} {
						return coer108(arg108)
					}
				}((func(_2176_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_2177_i5 _dafny.Int) _dafny.CodePoint {
						return _2176_v41
					}
				})(_2170_v41)))
			})(), 21)
			(_nw322).ArraySet1(Companion_Default___.Fm60(_this.F11, !(!((_2172_v43).Dtor_cf59())), _2173_v44, p0, globalState), 22)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("ihv"), (Companion_Default___.SafeIndex((_2174_v45).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ihv")).Cardinality()))).Uint32(), _dafny.CodePoint('e')), 23)
			(_nw322).ArraySet1(_2169_v40, 24)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2169_v40, _2169_v40), 25)
			(_nw322).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2169_v40, _dafny.UnicodeSeqOfUtf8Bytes("kf")), 26)
			_2175_v46 = _nw322
			var _index373 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_2175_v46), 0))
			_ = _index373
			(_2175_v46).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2169_v40, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(927))).Uint32(), func(coer109 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg109 _dafny.Int) interface{} {
					return coer109(arg109)
				}
			}((func(_2178_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2179_i6 _dafny.Int) _dafny.CodePoint {
					return _2178_v41
				}
			})(_2170_v41)))), (_index373).Int())
			var _2180_v47 _dafny.Array
			_ = _2180_v47
			var _nw323 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
			_ = _nw323
			_2180_v47 = _nw323
			var _2181_v48 _dafny.Map
			_ = _2181_v48
			_2181_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(828), _2180_v47)
			var _2182_v49 _dafny.MultiSet
			_ = _2182_v49
			_2182_v49 = _dafny.MultiSetOf(_2180_v47, _2180_v47, (func() _dafny.Array {
				if (_2181_v48).Contains(p0) {
					return (_2181_v48).Get(p0).(_dafny.Array)
				}
				return _2180_v47
			})())
			var _2183_v50 _dafny.MultiSet
			_ = _2183_v50
			_2183_v50 = _dafny.MultiSetOf(_dafny.IntOfInt64(773))
			var _index374 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_2175_v46), 0))
			_ = _index374
			var _rhs338 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(994))).Uint32(), func(coer110 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg110 _dafny.Int) interface{} {
					return coer110(arg110)
				}
			}((func(_2184_v41 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_2185_i7 _dafny.Int) _dafny.CodePoint {
					return _2184_v41
				}
			})(_2170_v41)))
			_ = _rhs338
			var _rhs339 _dafny.MultiSet = _2182_v49
			_ = _rhs339
			var _rhs340 bool = (func() bool {
				if (_2173_v44).Contains(p0) {
					return (_2173_v44).Get(p0).(bool)
				}
				return (_2183_v50).IsDisjointFrom(_2183_v50)
			})()
			_ = _rhs340
			var _rhs341 bool = _this.F11
			_ = _rhs341
			var _lhs268 _dafny.Array = _2175_v46
			_ = _lhs268
			var _lhs269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_2175_v46), 0))
			_ = _lhs269
			var _lhs270 *C8 = _this
			_ = _lhs270
			var _lhs271 *C8 = _this
			_ = _lhs271
			(_lhs268).ArraySet1(_rhs338, (_lhs269).Int())
			_2182_v49 = _rhs339
			_lhs270.F11 = _rhs340
			_lhs271.F11 = _rhs341
			var _index375 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))
			_ = _index375
			(_2171_v42).ArraySet1((func() bool {
				if _this.F11 {
					return _this.F11
				}
				return _this.F11
			})(), (_index375).Int())
			var _2186_v51 _dafny.Array
			_ = _2186_v51
			var _len0_55 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_55
			var _nw324 _dafny.Array
			_ = _nw324
			if _len0_55.Cmp(_dafny.Zero) == 0 {
				_nw324 = _dafny.NewArray(_len0_55)
			} else {
				var _init55 func(_dafny.Int) _dafny.Int = (func(_2187_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2188_i8 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_2188_i8, _2187_p0)
					}
				})(p0)
				_ = _init55
				var _element0_55 = _init55(_dafny.Zero)
				_ = _element0_55
				_nw324 = _dafny.NewArrayFromExample(_element0_55, nil, _len0_55)
				(_nw324).ArraySet1(_element0_55, 0)
				var _nativeLen0_55 = (_len0_55).Int()
				_ = _nativeLen0_55
				for _i0_55 := 1; _i0_55 < _nativeLen0_55; _i0_55++ {
					(_nw324).ArraySet1(_init55(_dafny.IntOf(_i0_55)), _i0_55)
				}
			}
			_2186_v51 = _nw324
			var _index376 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_2186_v51), 0))
			_ = _index376
			(_2186_v51).ArraySet1(p0, (_index376).Int())
			var _2189_v52 _dafny.Sequence
			_ = _2189_v52
			_2189_v52 = _dafny.SeqOf(true, false, true)
			var _2190_v53 _dafny.Map
			_ = _2190_v53
			_2190_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm30(_dafny.IntOfInt64(596), globalState), _2171_v42)
			var _index377 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))
			_ = _index377
			var _index378 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_2186_v51), 0))
			_ = _index378
			var _rhs342 bool = !(true) || ((_dafny.IntOfUint32((_2189_v52).Cardinality())).Cmp(p0) >= 0)
			_ = _rhs342
			var _rhs343 _dafny.Int = p0
			_ = _rhs343
			var _rhs344 _dafny.Int = (p0).Times(Companion_Default___.SafeModuloInt(Companion_Default___.Fm0(_2189_v52, _this.F11, _this.F11, p0, globalState), p0))
			_ = _rhs344
			var _rhs345 _dafny.Array = (func() _dafny.Array {
				if (_2190_v53).Contains(Companion_Default___.Fm30((_dafny.Zero).Minus(p0), globalState)) {
					return (_2190_v53).Get(Companion_Default___.Fm30((_dafny.Zero).Minus(p0), globalState)).(_dafny.Array)
				}
				return _2171_v42
			})()
			_ = _rhs345
			var _rhs346 _dafny.Int = Companion_Default___.Fm0(_2189_v52, (_this.F11) && (_this.F11), false, _dafny.IntOfInt64(953), globalState)
			_ = _rhs346
			var _lhs272 _dafny.Array = _2171_v42
			_ = _lhs272
			var _lhs273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))
			_ = _lhs273
			var _lhs274 *GlobalState = globalState
			_ = _lhs274
			var _lhs275 *GlobalState = globalState
			_ = _lhs275
			var _lhs276 _dafny.Array = _2186_v51
			_ = _lhs276
			var _lhs277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_2186_v51), 0))
			_ = _lhs277
			(_lhs272).ArraySet1(_rhs342, (_lhs273).Int())
			_lhs274.F7 = _rhs343
			_lhs275.F7 = _rhs344
			_2171_v42 = _rhs345
			(_lhs276).ArraySet1(_rhs346, (_lhs277).Int())
			(globalState).F7 = (func() _dafny.Int {
				if (_dafny.IntOfUint32((_2169_v40).Cardinality())).Cmp(p0) <= 0 {
					return (_2186_v51).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(314), _dafny.ArrayLen((_2186_v51), 0))).Int()).(_dafny.Int)
				}
				return Companion_Default___.Fm0(_2189_v52, _this.F11, (_2171_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))).Int()).(bool), p0, globalState)
			})()
			var _index379 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))
			_ = _index379
			(_2171_v42).ArraySet1(_this.F11, (_index379).Int())
			_2173_v44 = (_2173_v44).Update(p0, (_2171_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.ArrayLen((_2171_v42), 0))).Int()).(bool))
		}
		var _2191_v54 _dafny.Array
		_ = _2191_v54
		var _len0_56 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_56
		var _nw325 _dafny.Array
		_ = _nw325
		if _len0_56.Cmp(_dafny.Zero) == 0 {
			_nw325 = _dafny.NewArray(_len0_56)
		} else {
			var _init56 func(_dafny.Int) bool = func(_2192_i10 _dafny.Int) bool {
				return (_dafny.MultiSetOf(_this.F11)).IsProperSubsetOf(_dafny.MultiSetOf(_this.F11))
			}
			_ = _init56
			var _element0_56 = _init56(_dafny.Zero)
			_ = _element0_56
			_nw325 = _dafny.NewArrayFromExample(_element0_56, nil, _len0_56)
			(_nw325).ArraySet1(_element0_56, 0)
			var _nativeLen0_56 = (_len0_56).Int()
			_ = _nativeLen0_56
			for _i0_56 := 1; _i0_56 < _nativeLen0_56; _i0_56++ {
				(_nw325).ArraySet1(_init56(_dafny.IntOf(_i0_56)), _i0_56)
			}
		}
		_2191_v54 = _nw325
		for _iter67 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2191_v54), 0))); ; {
			_guard_loop_9, _ok67 := _iter67()
			if !_ok67 {
				break
			}
			var _2193_i9 _dafny.Int
			_2193_i9 = interface{}(_guard_loop_9).(_dafny.Int)
			if (true) && (((_2193_i9).Sign() != -1) && ((_2193_i9).Cmp(_dafny.ArrayLen((_2191_v54), 0)) < 0)) {
				(_2191_v54).ArraySet1(_this.F11, (_2193_i9).Int())
			}
		}
		var _2194_v55 _dafny.Array
		_ = _2194_v55
		var _len0_57 _dafny.Int = _dafny.IntOfInt64(17)
		_ = _len0_57
		var _nw326 _dafny.Array
		_ = _nw326
		if _len0_57.Cmp(_dafny.Zero) == 0 {
			_nw326 = _dafny.NewArray(_len0_57)
		} else {
			var _init57 func(_dafny.Int) _dafny.Set = (func(_2195_p0 _dafny.Int) func(_dafny.Int) _dafny.Set {
				return func(_2196_i12 _dafny.Int) _dafny.Set {
					return _dafny.SetOf(_2195_p0, _2195_p0)
				}
			})(p0)
			_ = _init57
			var _element0_57 = _init57(_dafny.Zero)
			_ = _element0_57
			_nw326 = _dafny.NewArrayFromExample(_element0_57, nil, _len0_57)
			(_nw326).ArraySet1(_element0_57, 0)
			var _nativeLen0_57 = (_len0_57).Int()
			_ = _nativeLen0_57
			for _i0_57 := 1; _i0_57 < _nativeLen0_57; _i0_57++ {
				(_nw326).ArraySet1(_init57(_dafny.IntOf(_i0_57)), _i0_57)
			}
		}
		_2194_v55 = _nw326
		for _iter68 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_2194_v55), 0))); ; {
			_guard_loop_10, _ok68 := _iter68()
			if !_ok68 {
				break
			}
			var _2197_i11 _dafny.Int
			_2197_i11 = interface{}(_guard_loop_10).(_dafny.Int)
			if (true) && (((_2197_i11).Sign() != -1) && ((_2197_i11).Cmp(_dafny.ArrayLen((_2194_v55), 0)) < 0)) {
				(_2194_v55).ArraySet1(_dafny.SetOf((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfInt64(235))), p0), (_2197_i11).Int())
			}
		}
		(_this).F11 = (p0).Cmp((_dafny.IntOfInt64(-217)).Times(p0)) != 0
		if true {
			var _2198_v56 _dafny.CodePoint
			_ = _2198_v56
			_2198_v56 = _dafny.CodePoint('n')
			var _2199_v57 _dafny.Map
			_ = _2199_v57
			_2199_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p0)
			var _2200_v58 D1
			_ = _2200_v58
			_2200_v58 = Companion_D1_.Create_DC2_((func() _dafny.Int {
				if (_2199_v57).Contains(_this.F11) {
					return (_2199_v57).Get(_this.F11).(_dafny.Int)
				}
				return p0
			})(), (_dafny.Zero).Minus(p0), _this.F11)
			var _2201_v59 *C0
			_ = _2201_v59
			var _nw327 *C0 = New_C0_()
			_ = _nw327
			_nw327.Ctor__(_this.F11, _2198_v56, _2200_v58, !(_this.F11))
			_2201_v59 = _nw327
			_2201_v59 = _2201_v59
			(globalState).F7 = p0
			if _2201_v59.F19 {
				var _2202_v60 _dafny.Sequence
				_ = _2202_v60
				_2202_v60 = _dafny.SeqOf(_2201_v59.F19, _2201_v59.F19, _this.F11)
				(globalState).F7 = Companion_Default___.Fm0(_2202_v60, _2201_v59.F19, _this.F11, (_dafny.IntOfUint32((_2202_v60).Cardinality())).Minus(p0), globalState)
				var _2203_v61 _dafny.Sequence
				_ = _2203_v61
				_2203_v61 = _dafny.UnicodeSeqOfUtf8Bytes("ped")
				var _rhs347 _dafny.Int = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm30((_dafny.SetOf(_2203_v61, _2203_v61)).Cardinality(), globalState))).Cardinality()).Plus(Companion_Default___.Fm0(_dafny.SeqOf(_2201_v59.F19, _2201_v59.F19), _this.F11, _this.F11, p0, globalState))
				_ = _rhs347
				var _rhs348 _dafny.Int = p0
				_ = _rhs348
				var _lhs278 *GlobalState = globalState
				_ = _lhs278
				var _lhs279 *GlobalState = globalState
				_ = _lhs279
				_lhs278.F7 = _rhs347
				_lhs279.F7 = _rhs348
				var _2204_v62 _dafny.Sequence
				_ = _2204_v62
				_2204_v62 = _dafny.SeqOf(_dafny.IntOfInt64(7))
				var _2205_v63 D7
				_ = _2205_v63
				_2205_v63 = Companion_D7_.Create_DC15_(_2204_v62)
				var _2206_v64 _dafny.Map
				_ = _2206_v64
				_2206_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2201_v59.F19, (_2201_v59).F20())
				var _2207_v65 *C4
				_ = _2207_v65
				var _nw328 *C4 = New_C4_()
				_ = _nw328
				_nw328.Ctor__()
				_2207_v65 = _nw328
				var _2208_v66 _dafny.Sequence
				_ = _2208_v66
				_2208_v66 = _dafny.SeqOf(_2207_v65)
				var _2209_v67 _dafny.Array
				_ = _2209_v67
				var _nwElement0_72 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_2204_v62, _dafny.SeqOf((_dafny.MultiSetOf(_this.F11, _2201_v59.F19)).Cardinality(), Companion_Default___.Fm0(_2202_v60, _this.F11, _2201_v59.F19, p0, globalState)))
				_ = _nwElement0_72
				var _nw329 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_72, nil, _dafny.IntOfInt64(20))
				_ = _nw329
				(_nw329).ArraySet1(_nwElement0_72, 0)
				(_nw329).ArraySet1(_2204_v62, 1)
				(_nw329).ArraySet1(_2204_v62, 2)
				(_nw329).ArraySet1(_2204_v62, 3)
				(_nw329).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_2205_v63).Dtor_cf28(), _2204_v62), 4)
				(_nw329).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(828))).Uint32(), func(coer111 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg111 _dafny.Int) interface{} {
						return coer111(arg111)
					}
				}((func(_2210_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2211_i13 _dafny.Int) _dafny.Int {
						return _2210_p0
					}
				})(p0))), 5)
				(_nw329).ArraySet1(_2204_v62, 6)
				(_nw329).ArraySet1(_2204_v62, 7)
				(_nw329).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2204_v62, _2204_v62), 8)
				(_nw329).ArraySet1(_2204_v62, 9)
				(_nw329).ArraySet1(_2204_v62, 10)
				(_nw329).ArraySet1(_2204_v62, 11)
				(_nw329).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(973))).Uint32(), func(coer112 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg112 _dafny.Int) interface{} {
						return coer112(arg112)
					}
				}((func(_2212_v59 *C0, _2213_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_2214_i14 _dafny.Int) _dafny.Int {
						return (_dafny.SetOf(_dafny.IntOfInt64(432), (_dafny.SetOf(_2212_v59.F19, _this.F11, _2212_v59.F19, _this.F11, _2212_v59.F19)).Cardinality(), _2213_p0)).Cardinality()
					}
				})(_2201_v59, p0))), _2204_v62), 12)
				(_nw329).ArraySet1(_2204_v62, 13)
				(_nw329).ArraySet1(_dafny.SeqOf(p0), 14)
				(_nw329).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_2204_v62, _dafny.Companion_Sequence_.Update(_2204_v62, (Companion_Default___.SafeIndex((_2206_v64).Cardinality(), _dafny.IntOfUint32((_2204_v62).Cardinality()))).Uint32(), p0)), 15)
				(_nw329).ArraySet1(_2204_v62, 16)
				(_nw329).ArraySet1(_2204_v62, 17)
				(_nw329).ArraySet1(_2204_v62, 18)
				(_nw329).ArraySet1(_dafny.SeqOf(p0, _dafny.IntOfUint32((_2208_v66).Cardinality())), 19)
				_2209_v67 = _nw329
				var _rhs349 _dafny.Int = (_dafny.Zero).Minus(p0)
				_ = _rhs349
				var _rhs350 _dafny.Array = _2209_v67
				_ = _rhs350
				var _lhs280 *GlobalState = globalState
				_ = _lhs280
				_lhs280.F7 = _rhs349
				_2209_v67 = _rhs350
				var _index380 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index380
				(_2191_v54).ArraySet1(_2201_v59.F19, (_index380).Int())
				var _2215_v68 _dafny.Map
				_ = _2215_v68
				_2215_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2201_v59.F19)
				var _index381 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index381
				(_2191_v54).ArraySet1(_dafny.Companion_Sequence_.Equal(_2204_v62, _dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, (_2215_v68).Cardinality(), p0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(p0, (_2215_v68).Cardinality(), p0)).Cardinality()))).Uint32(), _dafny.IntOfInt64(725))), (_index381).Int())
				var _2216_v69 _dafny.Set
				_ = _2216_v69
				_2216_v69 = _dafny.SetOf(_2201_v59.F19, true, _2201_v59.F19)
				var _2217_v70 _dafny.MultiSet
				_ = _2217_v70
				_2217_v70 = _dafny.MultiSetOf(_this.F11, _this.F11)
				var _index382 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index382
				var _index383 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index383
				var _rhs351 bool = !(_2216_v69).Contains((_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool))
				_ = _rhs351
				var _rhs352 bool = ((Companion_Default___.Fm40(globalState)).Update(_this.F11, Companion_Default___.Abs((_2216_v69).Cardinality()))).IsSubsetOf(_2217_v70)
				_ = _rhs352
				var _lhs281 _dafny.Array = _2191_v54
				_ = _lhs281
				var _lhs282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _lhs282
				var _lhs283 _dafny.Array = _2191_v54
				_ = _lhs283
				var _lhs284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_2191_v54), 0))
				_ = _lhs284
				(_lhs281).ArraySet1(_rhs351, (_lhs282).Int())
				(_lhs283).ArraySet1(_rhs352, (_lhs284).Int())
			} else {
				var _2218_v71 _dafny.Sequence
				_ = _2218_v71
				_2218_v71 = _dafny.UnicodeSeqOfUtf8Bytes("arsssdc")
				_2198_v56 = (_2218_v71).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2218_v71).Cardinality()))).Uint32()).(_dafny.CodePoint)
				var _2219_v72 _dafny.Set
				_ = _2219_v72
				_2219_v72 = _dafny.SetOf(_2218_v71)
				var _2220_v73 _dafny.Map
				_ = _2220_v73
				_2220_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2219_v72).Difference(_dafny.SetOf(_2218_v71, _2218_v71)), _2201_v59.F19)
				_2220_v73 = (_2220_v73).Update(_2219_v72, _2201_v59.F19)
				(_2201_v59).F19 = (func() bool {
					if _2201_v59.F19 {
						return false
					}
					return false
				})()
				var _index384 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index384
				(_2191_v54).ArraySet1(false, (_index384).Int())
				var _index385 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index385
				(_2191_v54).ArraySet1(_this.F11, (_index385).Int())
				var _2221_v74 _dafny.Set
				_ = _2221_v74
				_2221_v74 = _dafny.SetOf(p0)
				var _2222_v75 _dafny.Map
				_ = _2222_v75
				_2222_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, (_2221_v74).IsSubsetOf(_2221_v74))
				var _2223_v76 _dafny.Array
				_ = _2223_v76
				var _len0_58 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_58
				var _nw330 _dafny.Array
				_ = _nw330
				if _len0_58.Cmp(_dafny.Zero) == 0 {
					_nw330 = _dafny.NewArray(_len0_58)
				} else {
					var _init58 func(_dafny.Int) _dafny.Int = (func(_2224_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_2225_i15 _dafny.Int) _dafny.Int {
							return (_2225_i15).Minus(_2224_p0)
						}
					})(p0)
					_ = _init58
					var _element0_58 = _init58(_dafny.Zero)
					_ = _element0_58
					_nw330 = _dafny.NewArrayFromExample(_element0_58, nil, _len0_58)
					(_nw330).ArraySet1(_element0_58, 0)
					var _nativeLen0_58 = (_len0_58).Int()
					_ = _nativeLen0_58
					for _i0_58 := 1; _i0_58 < _nativeLen0_58; _i0_58++ {
						(_nw330).ArraySet1(_init58(_dafny.IntOf(_i0_58)), _i0_58)
					}
				}
				_2223_v76 = _nw330
				var _2226_v77 _dafny.MultiSet
				_ = _2226_v77
				_2226_v77 = _dafny.MultiSetOf((_2201_v59).F20())
				var _2227_v78 _dafny.Map
				_ = _2227_v78
				_2227_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _2226_v77)
				var _2228_v79 D12
				_ = _2228_v79
				_2228_v79 = Companion_D12_.Create_DC31_(_this.F11, _2223_v76, _2201_v59.F19, (_2201_v59).F20(), (func() _dafny.MultiSet {
					if (_2227_v78).Contains(_2201_v59.F19) {
						return (_2227_v78).Get(_2201_v59.F19).(_dafny.MultiSet)
					}
					return _dafny.MultiSetFromSeq(_dafny.SeqOf((_2201_v59).F20()))
				})())
				var _2229_v80 _dafny.Set
				_ = _2229_v80
				_2229_v80 = _dafny.SetOf(true, (_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool), (_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool), _2201_v59.F19, true)
				var _index386 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index386
				(_2191_v54).ArraySet1((func() bool {
					if (_2222_v75).Contains(!((_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool)) || ((_2228_v79).Dtor_cf62())) {
						return (_2222_v75).Get(!((_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(399), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool)) || ((_2228_v79).Dtor_cf62())).(bool)
					}
					return !(Companion_Default___.Fm30((_2229_v80).Cardinality(), globalState))
				})(), (_index386).Int())
			}
			var _2230_v81 _dafny.Sequence
			_ = _2230_v81
			_2230_v81 = _dafny.SeqOf(_2201_v59.F19, true)
			var _2231_v82 _dafny.MultiSet
			_ = _2231_v82
			_2231_v82 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2230_v81).Cardinality()))
			var _2232_v83 _dafny.Map
			_ = _2232_v83
			_2232_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2231_v82, _2191_v54)
			var _2233_v84 _dafny.Set
			_ = _2233_v84
			_2233_v84 = _dafny.SetOf(p0)
			var _2234_v85 *C5
			_ = _2234_v85
			var _nw331 *C5 = New_C5_()
			_ = _nw331
			_nw331.Ctor__(_2201_v59.F19, (func() _dafny.Array {
				if (_2232_v83).Contains(_2231_v82) {
					return (_2232_v83).Get(_2231_v82).(_dafny.Array)
				}
				return _2191_v54
			})(), (_2233_v84).Difference(_2233_v84), _2200_v58, _2201_v59.F19)
			_2234_v85 = _nw331
			if _this.F11 {
				var _2235_v86 _dafny.Sequence
				_ = _2235_v86
				_2235_v86 = _dafny.UnicodeSeqOfUtf8Bytes("hhqingxd")
				var _index387 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index387
				(_2191_v54).ArraySet1((_this).Fm4(_2230_v81, p0, _2235_v86, _2234_v85.F24, globalState), (_index387).Int())
				var _2236_v87 _dafny.Sequence
				_ = _2236_v87
				_2236_v87 = _dafny.SeqOf(_2235_v86, _2235_v86, _2235_v86, _2235_v86, _2235_v86)
				var _2237_v88 _dafny.Map
				_ = _2237_v88
				_2237_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2236_v87).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2236_v87).Cardinality()))).Uint32()).(_dafny.Sequence), _2199_v57)
				var _index388 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(6), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index388
				(_2191_v54).ArraySet1((_2237_v88).Contains(_2235_v86), (_index388).Int())
				var _2238_v89 _dafny.Sequence
				_ = _2238_v89
				_2238_v89 = _dafny.SeqOf(p0, (_dafny.Zero).Minus(p0))
				var _2239_v90 _dafny.Map
				_ = _2239_v90
				_2239_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !(_2234_v85.F24))
				var _2240_v91 _dafny.Set
				_ = _2240_v91
				_2240_v91 = _dafny.SetOf(false, _2234_v85.F24)
				var _2241_v92 _dafny.Map
				_ = _2241_v92
				_2241_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2240_v91).Cardinality(), p0)
				var _2242_v93 _dafny.Sequence
				_ = _2242_v93
				_2242_v93 = _dafny.SeqOf(_2231_v82, _2231_v82)
				var _2243_v94 _dafny.Map
				_ = _2243_v94
				_2243_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2240_v91, _dafny.IntOfUint32((_2230_v81).Cardinality()))
				var _2244_v95 _dafny.Array
				_ = _2244_v95
				var _nwElement0_73 _dafny.Int = (_2238_v89).Select((Companion_Default___.SafeIndex((_2234_v85).Fm9((_2201_v59).F20(), globalState), _dafny.IntOfUint32((_2238_v89).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _nwElement0_73
				var _nw332 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_73, nil, _dafny.IntOfInt64(16))
				_ = _nw332
				(_nw332).ArraySet1(_nwElement0_73, 0)
				(_nw332).ArraySet1(p0, 1)
				(_nw332).ArraySet1((_2239_v90).Cardinality(), 2)
				(_nw332).ArraySet1(p0, 3)
				(_nw332).ArraySet1((_dafny.Zero).Minus((_2234_v85).Fm9((_2201_v59).F20(), globalState)), 4)
				(_nw332).ArraySet1(p0, 5)
				(_nw332).ArraySet1((p0).Plus((_dafny.Zero).Minus((func() _dafny.Int {
					if (_2241_v92).Contains(p0) {
						return (_2241_v92).Get(p0).(_dafny.Int)
					}
					return p0
				})())), 6)
				(_nw332).ArraySet1(p0, 7)
				(_nw332).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, p0), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(p0, p0)).Cardinality()))).Uint32(), (_dafny.Zero).Minus(p0))).Cardinality())), 8)
				(_nw332).ArraySet1((p0).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("h")).Cardinality())), 9)
				(_nw332).ArraySet1(p0, 10)
				(_nw332).ArraySet1((_dafny.Zero).Minus(p0), 11)
				(_nw332).ArraySet1((_2239_v90).Cardinality(), 12)
				(_nw332).ArraySet1(_dafny.IntOfInt64(366), 13)
				(_nw332).ArraySet1(Companion_Default___.SafeDivisionInt(((_2242_v93).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2242_v93).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality(), (Companion_Default___.Fm64(globalState)).Cardinality()), 14)
				(_nw332).ArraySet1(((_2243_v94).Merge(_2243_v94)).Cardinality(), 15)
				_2244_v95 = _nw332
				_2244_v95 = _2244_v95
				(globalState).F7 = (p0).Times(p0)
				var _rhs353 bool = !(_this.F11)
				_ = _rhs353
				var _rhs354 _dafny.Int = p0
				_ = _rhs354
				var _lhs285 *C5 = _2234_v85
				_ = _lhs285
				var _lhs286 *GlobalState = globalState
				_ = _lhs286
				_lhs285.F24 = _rhs353
				_lhs286.F7 = _rhs354
				(_this).F11 = (_2201_v59.F19) == (_dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("sftaoljs"), (_2236_v87).Select((Companion_Default___.SafeIndex((_2240_v91).Cardinality(), _dafny.IntOfUint32((_2236_v87).Cardinality()))).Uint32()).(_dafny.Sequence)))
			} else {
				var _2245_v96 _dafny.Map
				_ = _2245_v96
				_2245_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F11)
				var _2246_v97 _dafny.MultiSet
				_ = _2246_v97
				_2246_v97 = _dafny.MultiSetOf((_2245_v96).Merge(_2245_v96))
				var _2247_v98 _dafny.Sequence
				_ = _2247_v98
				_2247_v98 = _dafny.SeqOf(_2245_v96, _2245_v96, _2245_v96, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2201_v59.F19, _2201_v59.F19), (_2245_v96).Update(_2234_v85.F24, _this.F11))
				_2246_v97 = (_dafny.MultiSetFromSeq(_2247_v98)).Difference(_2246_v97)
				var _2248_v99 _dafny.MultiSet
				_ = _2248_v99
				_2248_v99 = _dafny.MultiSetOf(_2191_v54, _2191_v54, _2191_v54, _2191_v54)
				_2248_v99 = _2248_v99
				var _2249_v100 _dafny.Sequence
				_ = _2249_v100
				_2249_v100 = _dafny.UnicodeSeqOfUtf8Bytes("bvpukkl")
				var _2250_v102 _dafny.Set
				_ = _2250_v102
				_2250_v102 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ipcpel"), _2249_v100)
				var _2251_v103 _dafny.MultiSet
				_ = _2251_v103
				_2251_v103 = _dafny.MultiSetOf(_2249_v100)
				var _2252_v105 _dafny.Sequence
				_ = _2252_v105
				_2252_v105 = _dafny.SeqOf(_dafny.SetOf(_2249_v100, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-402))).Uint32(), func(coer113 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg113 _dafny.Int) interface{} {
						return coer113(arg113)
					}
				}((func(_2253_v59 *C0) func(_dafny.Int) _dafny.CodePoint {
					return func(_2254_i16 _dafny.Int) _dafny.CodePoint {
						return (_2253_v59).F20()
					}
				})(_2201_v59))), _2249_v100), func() _dafny.Set {
					var _coll56 = _dafny.NewBuilder()
					_ = _coll56
					for _iter69 := _dafny.Iterate((_2251_v103).Elements()); ; {
						_compr_56, _ok69 := _iter69()
						if !_ok69 {
							break
						}
						var _2255_v104 _dafny.Sequence
						_2255_v104 = interface{}(_compr_56).(_dafny.Sequence)
						if (_2251_v103).Contains(_2255_v104) {
							_coll56.Add(_2255_v104)
						}
					}
					return _coll56.ToSet()
				}())
				var _2256_v106 _dafny.Sequence
				_ = _2256_v106
				_2256_v106 = _dafny.SeqOf((_dafny.Zero).Minus(p0), p0)
				var _2257_v108 _dafny.Map
				_ = _2257_v108
				_2257_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _2258_v109 *C7
				_ = _2258_v109
				var _nw333 *C7 = New_C7_()
				_ = _nw333
				_nw333.Ctor__()
				_2258_v109 = _nw333
				var _2259_v110 _dafny.Map
				_ = _2259_v110
				_2259_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2249_v100, _2258_v109)
				var _2260_v111 _dafny.Map
				_ = _2260_v111
				_2260_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Int {
					if (_2257_v108).Contains(p0) {
						return (_2257_v108).Get(p0).(_dafny.Int)
					}
					return (_2259_v110).Cardinality()
				})())
				var _2261_v112 _dafny.Array
				_ = _2261_v112
				var _nwElement0_74 _dafny.Set = func() _dafny.Set {
					var _coll57 = _dafny.NewBuilder()
					_ = _coll57
					for _iter70 := _dafny.Iterate((_dafny.SeqOf(_2249_v100)).Elements()); ; {
						_compr_57, _ok70 := _iter70()
						if !_ok70 {
							break
						}
						var _2262_v101 _dafny.Sequence
						_2262_v101 = interface{}(_compr_57).(_dafny.Sequence)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_2249_v100), _2262_v101) {
							_coll57.Add(_2262_v101)
						}
					}
					return _coll57.ToSet()
				}()
				_ = _nwElement0_74
				var _nw334 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_74, nil, _dafny.IntOfInt64(17))
				_ = _nw334
				(_nw334).ArraySet1(_nwElement0_74, 0)
				(_nw334).ArraySet1(Companion_Default___.Fm48(globalState), 1)
				(_nw334).ArraySet1((_2250_v102).Union(_2250_v102), 2)
				(_nw334).ArraySet1(_2250_v102, 3)
				(_nw334).ArraySet1(_2250_v102, 4)
				(_nw334).ArraySet1(((_2252_v105).Select((Companion_Default___.SafeIndex((_2256_v106).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2256_v106).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_2252_v105).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_2250_v102), 5)
				(_nw334).ArraySet1(func() _dafny.Set {
					var _coll58 = _dafny.NewBuilder()
					_ = _coll58
					for _iter71 := _dafny.Iterate((_2251_v103).Elements()); ; {
						_compr_58, _ok71 := _iter71()
						if !_ok71 {
							break
						}
						var _2263_v107 _dafny.Sequence
						_2263_v107 = interface{}(_compr_58).(_dafny.Sequence)
						if (_2251_v103).Contains(_2263_v107) {
							_coll58.Add(_2263_v107)
						}
					}
					return _coll58.ToSet()
				}(), 6)
				(_nw334).ArraySet1(_2250_v102, 7)
				(_nw334).ArraySet1(_2250_v102, 8)
				(_nw334).ArraySet1(_2250_v102, 9)
				(_nw334).ArraySet1(_2250_v102, 10)
				(_nw334).ArraySet1(_2250_v102, 11)
				(_nw334).ArraySet1((_2252_v105).Select((Companion_Default___.SafeIndex((_2257_v108).Cardinality(), _dafny.IntOfUint32((_2252_v105).Cardinality()))).Uint32()).(_dafny.Set), 12)
				(_nw334).ArraySet1(_2250_v102, 13)
				(_nw334).ArraySet1(_2250_v102, 14)
				(_nw334).ArraySet1(_2250_v102, 15)
				(_nw334).ArraySet1((_2250_v102).Intersection(Companion_Default___.Fm48(globalState)), 16)
				_2261_v112 = _nw334
				var _index389 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_2261_v112), 0))
				_ = _index389
				(_2261_v112).ArraySet1(_2250_v102, (_index389).Int())
				var _index390 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(813), _dafny.ArrayLen((_2261_v112), 0))
				_ = _index390
				(_2261_v112).ArraySet1((func() _dafny.Set {
					if (_this).Fm4(_2230_v81, p0, _2249_v100, _2201_v59.F19, globalState) {
						return _2250_v102
					}
					return (_dafny.SetOf(_2249_v100, _2249_v100)).Difference(_2250_v102)
				})(), (_index390).Int())
				var _2264_v113 _dafny.Array
				_ = _2264_v113
				var _nw335 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
				_ = _nw335
				_2264_v113 = _nw335
				var _2265_v114 _dafny.MultiSet
				_ = _2265_v114
				_2265_v114 = _dafny.MultiSetOf(_2264_v113)
				_2265_v114 = (_2265_v114).Difference(_2265_v114)
				(_this).F11 = !(false)
			}
		} else {
			var _2266_v115 _dafny.CodePoint
			_ = _2266_v115
			_2266_v115 = _dafny.CodePoint('i')
			var _2267_v116 _dafny.Map
			_ = _2267_v116
			_2267_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p0).Cmp(_dafny.IntOfInt64(680)) <= 0, _2266_v115)
			_2267_v116 = _2267_v116
			var _2268_v117 _dafny.Array
			_ = _2268_v117
			var _nw336 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw336
			_2268_v117 = _nw336
			var _2269_v118 _dafny.Set
			_ = _2269_v118
			_2269_v118 = _dafny.SetOf(_this.F11)
			var _index391 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))
			_ = _index391
			(_2268_v117).ArraySet1(((_2269_v118).Cardinality()).Minus(p0), (_index391).Int())
			var _2270_v119 _dafny.Map
			_ = _2270_v119
			_2270_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _2269_v118)
			var _2271_v120 _dafny.Sequence
			_ = _2271_v120
			_2271_v120 = _dafny.SeqOf(_this.F11, _this.F11)
			var _2272_v121 _dafny.Map
			_ = _2272_v121
			_2272_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _2269_v118)
			var _index392 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))
			_ = _index392
			(_2268_v117).ArraySet1(((func() _dafny.Set {
				if (_2270_v119).Contains((_dafny.Zero).Minus(Companion_Default___.Fm0(_2271_v120, _this.F11, _this.F11, p0, globalState))) {
					return (_2270_v119).Get((_dafny.Zero).Minus(Companion_Default___.Fm0(_2271_v120, _this.F11, _this.F11, p0, globalState))).(_dafny.Set)
				}
				return (func() _dafny.Set {
					if _this.F11 {
						return (func() _dafny.Set {
							if (_2272_v121).Contains(false) {
								return (_2272_v121).Get(false).(_dafny.Set)
							}
							return _2269_v118
						})()
					}
					return _2269_v118
				})()
			})()).Cardinality(), (_index392).Int())
			var _2273_v122 _dafny.Set
			_ = _2273_v122
			_2273_v122 = _dafny.SetOf((_dafny.Zero).Minus((_2269_v118).Cardinality()), (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), p0)
			var _rhs355 _dafny.Set = _2273_v122
			_ = _rhs355
			var _rhs356 _dafny.Int = ((p0).Plus(p0)).Minus(p0)
			_ = _rhs356
			var _rhs357 bool = _this.F11
			_ = _rhs357
			var _lhs287 *GlobalState = globalState
			_ = _lhs287
			var _lhs288 *C8 = _this
			_ = _lhs288
			_2273_v122 = _rhs355
			_lhs287.F7 = _rhs356
			_lhs288.F11 = _rhs357
			if _this.F11 {
				(_this).F11 = (func() bool {
					if _this.F11 {
						return (func() bool {
							if _this.F11 {
								return !(_this.F11)
							}
							return _this.F11
						})()
					}
					return (p0).Cmp(p0) != 0
				})()
				var _2274_v123 _dafny.Sequence
				_ = _2274_v123
				_2274_v123 = _dafny.UnicodeSeqOfUtf8Bytes("lf")
				var _index393 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))
				_ = _index393
				(_2268_v117).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((Companion_Default___.Fm55(_dafny.IntOfUint32((_2274_v123).Cardinality()), _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2266_v115, p0)), _this.F11, globalState)).Cardinality()), p0, (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int))).Cardinality()), (_index393).Int())
				var _2275_v124 _dafny.Array
				_ = _2275_v124
				var _nw337 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(15))
				_ = _nw337
				_2275_v124 = _nw337
				var _2276_v125 _dafny.Sequence
				_ = _2276_v125
				_2276_v125 = _dafny.SeqOf(_2269_v118)
				var _index394 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_2275_v124), 0))
				_ = _index394
				(_2275_v124).ArraySet1((_2276_v125).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2276_v125).Cardinality()))).Uint32()).(_dafny.Set), (_index394).Int())
				var _2277_v126 _dafny.Sequence
				_ = _2277_v126
				_2277_v126 = _dafny.SeqOf(_2191_v54)
				var _2278_v127 _dafny.Sequence
				_ = _2278_v127
				_2278_v127 = _dafny.SeqOf(_2274_v123, _2274_v123)
				var _2279_v128 _dafny.Map
				_ = _2279_v128
				_2279_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F11), _dafny.SeqOf(_this.F11, _this.F11))
				var _index395 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_2275_v124), 0))
				_ = _index395
				var _index396 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))
				_ = _index396
				var _rhs358 _dafny.Set = _2269_v118
				_ = _rhs358
				var _rhs359 _dafny.Array = (_2277_v126).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.IntOfUint32(((_2278_v127).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2278_v127).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality())).Minus(p0)), _dafny.IntOfUint32((_2277_v126).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs359
				var _rhs360 _dafny.Int = Companion_Default___.SafeModuloInt((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int))
				_ = _rhs360
				var _rhs361 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("e")
				_ = _rhs361
				var _rhs362 bool = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if (_2279_v128).Contains(_this.F11) {
						return (_2279_v128).Get(_this.F11).(_dafny.Sequence)
					}
					return _2271_v120
				})(), _2271_v120)).Cardinality())).Cmp((p0).Plus((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int))) > 0
				_ = _rhs362
				var _lhs289 _dafny.Array = _2275_v124
				_ = _lhs289
				var _lhs290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_2275_v124), 0))
				_ = _lhs290
				var _lhs291 _dafny.Array = _2268_v117
				_ = _lhs291
				var _lhs292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))
				_ = _lhs292
				var _lhs293 *C8 = _this
				_ = _lhs293
				(_lhs289).ArraySet1(_rhs358, (_lhs290).Int())
				_2191_v54 = _rhs359
				(_lhs291).ArraySet1(_rhs360, (_lhs292).Int())
				_2274_v123 = _rhs361
				_lhs293.F11 = _rhs362
				_2268_v117 = _2268_v117
				var _2280_v129 _dafny.Array
				_ = _2280_v129
				var _len0_59 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_59
				var _nw338 _dafny.Array
				_ = _nw338
				if _len0_59.Cmp(_dafny.Zero) == 0 {
					_nw338 = _dafny.NewArray(_len0_59)
				} else {
					var _init59 func(_dafny.Int) _dafny.MultiSet = (func(_2281_v117 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
						return func(_2282_i17 _dafny.Int) _dafny.MultiSet {
							return _dafny.MultiSetOf((_2281_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2281_v117), 0))).Int()).(_dafny.Int))
						}
					})(_2268_v117)
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
				_2280_v129 = _nw338
				var _2283_v130 _dafny.MultiSet
				_ = _2283_v130
				_2283_v130 = _dafny.MultiSetOf(false, _this.F11, _this.F11)
				var _2284_v131 _dafny.MultiSet
				_ = _2284_v131
				_2284_v131 = _dafny.MultiSetOf(_dafny.IntOfInt64(4), p0, (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
					if (_2283_v130).Contains(_this.F11) {
						return (_2283_v130).Multiplicity(_this.F11)
					}
					return (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int)
				})())
				var _index397 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_2280_v129), 0))
				_ = _index397
				(_2280_v129).ArraySet1(_2284_v131, (_index397).Int())
				var _index398 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_2280_v129), 0))
				_ = _index398
				(_2280_v129).ArraySet1(_2284_v131, (_index398).Int())
			} else {
				var _2285_v132 D7
				_ = _2285_v132
				_2285_v132 = Companion_D7_.Create_DC16_((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int))
				var _2286_v133 _dafny.Map
				_ = _2286_v133
				_2286_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC3_(_2191_v54), _this.F11)
				var _2287_v134 D2
				_ = _2287_v134
				_2287_v134 = Companion_D2_.Create_DC3_(_2191_v54)
				var _2288_v135 _dafny.Sequence
				_ = _2288_v135
				_2288_v135 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v134, true), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D2_.Create_DC3_(_2191_v54), true))
				var _2289_v136 _dafny.MultiSet
				_ = _2289_v136
				_2289_v136 = _dafny.MultiSetOf(_2266_v115)
				var _2290_v137 _dafny.Array
				_ = _2290_v137
				var _nwElement0_75 _dafny.Map = _2286_v133
				_ = _nwElement0_75
				var _nw339 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_75, nil, _dafny.IntOfInt64(24))
				_ = _nw339
				(_nw339).ArraySet1(_nwElement0_75, 0)
				(_nw339).ArraySet1(_2286_v133, 1)
				(_nw339).ArraySet1((_2288_v135).Select((Companion_Default___.SafeIndex((_2289_v136).Cardinality(), _dafny.IntOfUint32((_2288_v135).Cardinality()))).Uint32()).(_dafny.Map), 2)
				(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v134, _this.F11), 3)
				(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v134, false), 4)
				(_nw339).ArraySet1(_2286_v133, 5)
				(_nw339).ArraySet1(_2286_v133, 6)
				(_nw339).ArraySet1(_2286_v133, 7)
				(_nw339).ArraySet1((_2286_v133).Update(Companion_D2_.Create_DC3_(_2191_v54), _this.F11), 8)
				(_nw339).ArraySet1(_2286_v133, 9)
				(_nw339).ArraySet1(_2286_v133, 10)
				(_nw339).ArraySet1(_2286_v133, 11)
				(_nw339).ArraySet1(_2286_v133, 12)
				(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v134, _this.F11), 13)
				(_nw339).ArraySet1(_2286_v133, 14)
				(_nw339).ArraySet1(_2286_v133, 15)
				(_nw339).ArraySet1(_2286_v133, 16)
				(_nw339).ArraySet1(_2286_v133, 17)
				(_nw339).ArraySet1(_2286_v133, 18)
				(_nw339).ArraySet1(_2286_v133, 19)
				(_nw339).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2287_v134, _this.F11), 20)
				(_nw339).ArraySet1(_2286_v133, 21)
				(_nw339).ArraySet1(_2286_v133, 22)
				(_nw339).ArraySet1(_2286_v133, 23)
				_2290_v137 = _nw339
				var _2291_v138 _dafny.Map
				_ = _2291_v138
				_2291_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), (_2285_v132).Dtor_cf29()), _2290_v137)
				_2291_v138 = (_2291_v138).Update((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), _2290_v137)
				var _index399 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index399
				(_2191_v54).ArraySet1(!(_this.F11), (_index399).Int())
				var _2292_v139 _dafny.MultiSet
				_ = _2292_v139
				_2292_v139 = _dafny.MultiSetOf(true)
				var _2293_v140 _dafny.Map
				_ = _2293_v140
				_2293_v140 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _2292_v139)
				var _2294_v141 _dafny.MultiSet
				_ = _2294_v141
				_2294_v141 = _dafny.MultiSetOf(_2273_v122)
				var _2295_v142 D18
				_ = _2295_v142
				_2295_v142 = Companion_D18_.Create_DC46_(_2294_v141)
				var _2296_v143 _dafny.Sequence
				_ = _2296_v143
				_2296_v143 = _dafny.SeqOf(((_2295_v142).Dtor_cf89()).Cardinality())
				var _index400 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index400
				var _rhs363 _dafny.Int = (_dafny.Zero).Minus((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int))
				_ = _rhs363
				var _rhs364 _dafny.Set = _dafny.SetOf(((func() _dafny.MultiSet {
					if (_2293_v140).Contains(_this.F11) {
						return (_2293_v140).Get(_this.F11).(_dafny.MultiSet)
					}
					return _2292_v139
				})()).IsDisjointFrom(_dafny.MultiSetOf(false)), !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm63(false, (_2271_v120).Select((Companion_Default___.SafeIndex((_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_2271_v120).Cardinality()))).Uint32()).(bool), globalState), _2296_v143))
				_ = _rhs364
				var _rhs365 _dafny.Int = (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int)
				_ = _rhs365
				var _rhs366 bool = _this.F11
				_ = _rhs366
				var _rhs367 bool = _this.F11
				_ = _rhs367
				var _lhs294 *GlobalState = globalState
				_ = _lhs294
				var _lhs295 *GlobalState = globalState
				_ = _lhs295
				var _lhs296 _dafny.Array = _2191_v54
				_ = _lhs296
				var _lhs297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2191_v54), 0))
				_ = _lhs297
				var _lhs298 *C8 = _this
				_ = _lhs298
				_lhs294.F7 = _rhs363
				_2269_v118 = _rhs364
				_lhs295.F7 = _rhs365
				(_lhs296).ArraySet1(_rhs366, (_lhs297).Int())
				_lhs298.F11 = _rhs367
				var _2297_v144 _dafny.Array
				_ = _2297_v144
				var _len0_60 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_60
				var _nw340 _dafny.Array
				_ = _nw340
				if _len0_60.Cmp(_dafny.Zero) == 0 {
					_nw340 = _dafny.NewArray(_len0_60)
				} else {
					var _init60 func(_dafny.Int) bool = (func(_2298_v54 _dafny.Array) func(_dafny.Int) bool {
						return func(_2299_i18 _dafny.Int) bool {
							return (_2298_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2298_v54), 0))).Int()).(bool)
						}
					})(_2191_v54)
					_ = _init60
					var _element0_60 = _init60(_dafny.Zero)
					_ = _element0_60
					_nw340 = _dafny.NewArrayFromExample(_element0_60, nil, _len0_60)
					(_nw340).ArraySet1(_element0_60, 0)
					var _nativeLen0_60 = (_len0_60).Int()
					_ = _nativeLen0_60
					for _i0_60 := 1; _i0_60 < _nativeLen0_60; _i0_60++ {
						(_nw340).ArraySet1(_init60(_dafny.IntOf(_i0_60)), _i0_60)
					}
				}
				_2297_v144 = _nw340
				var _2300_v145 D2
				_ = _2300_v145
				_2300_v145 = Companion_D2_.Create_DC4_(_2268_v117, (_2191_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2191_v54), 0))).Int()).(bool), (_2268_v117).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(574), _dafny.ArrayLen((_2268_v117), 0))).Int()).(_dafny.Int), false, _2297_v144)
				var _2301_v146 _dafny.Array
				_ = _2301_v146
				var _nwElement0_76 _dafny.Array = _2268_v117
				_ = _nwElement0_76
				var _nw341 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_76, nil, _dafny.IntOfInt64(8))
				_ = _nw341
				(_nw341).ArraySet1(_nwElement0_76, 0)
				(_nw341).ArraySet1(_2268_v117, 1)
				(_nw341).ArraySet1(_2268_v117, 2)
				(_nw341).ArraySet1(_2268_v117, 3)
				(_nw341).ArraySet1(_2268_v117, 4)
				(_nw341).ArraySet1((_2300_v145).Dtor_cf6(), 5)
				(_nw341).ArraySet1(_2268_v117, 6)
				(_nw341).ArraySet1(_2268_v117, 7)
				_2301_v146 = _nw341
				var _index401 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_2301_v146), 0))
				_ = _index401
				(_2301_v146).ArraySet1(_2268_v117, (_index401).Int())
				var _index402 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(65), _dafny.ArrayLen((_2301_v146), 0))
				_ = _index402
				var _nw342 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
				_ = _nw342
				(_2301_v146).ArraySet1(_nw342, (_index402).Int())
				var _index403 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_2191_v54), 0))
				_ = _index403
				(_2191_v54).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_2271_v120, _2271_v120), (_index403).Int())
				var _2302_v147 _dafny.Array
				_ = _2302_v147
				var _nw343 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
				_ = _nw343
				_2302_v147 = _nw343
				var _2303_v148 _dafny.Sequence
				_ = _2303_v148
				_2303_v148 = _dafny.UnicodeSeqOfUtf8Bytes("psmsf")
				var _index404 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_2302_v147), 0))
				_ = _index404
				(_2302_v147).ArraySet1(_2303_v148, (_index404).Int())
				var _index405 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_2302_v147), 0))
				_ = _index405
				(_2302_v147).ArraySet1(_2303_v148, (_index405).Int())
			}
			var _2304_v149 _dafny.Array
			_ = _2304_v149
			var _len0_61 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_61
			var _nw344 _dafny.Array
			_ = _nw344
			if _len0_61.Cmp(_dafny.Zero) == 0 {
				_nw344 = _dafny.NewArray(_len0_61)
			} else {
				var _init61 func(_dafny.Int) _dafny.Map = func(_2305_i19 _dafny.Int) _dafny.Map {
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_this.F11), _this.F11)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_this.F11), _this.F11))
				}
				_ = _init61
				var _element0_61 = _init61(_dafny.Zero)
				_ = _element0_61
				_nw344 = _dafny.NewArrayFromExample(_element0_61, nil, _len0_61)
				(_nw344).ArraySet1(_element0_61, 0)
				var _nativeLen0_61 = (_len0_61).Int()
				_ = _nativeLen0_61
				for _i0_61 := 1; _i0_61 < _nativeLen0_61; _i0_61++ {
					(_nw344).ArraySet1(_init61(_dafny.IntOf(_i0_61)), _i0_61)
				}
			}
			_2304_v149 = _nw344
			var _2306_v150 _dafny.MultiSet
			_ = _2306_v150
			_2306_v150 = _dafny.MultiSetOf(_this.F11)
			var _2307_v151 _dafny.Map
			_ = _2307_v151
			_2307_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2306_v150, _this.F11)
			var _index406 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_2304_v149), 0))
			_ = _index406
			(_2304_v149).ArraySet1(_2307_v151, (_index406).Int())
			var _index407 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(414), _dafny.ArrayLen((_2304_v149), 0))
			_ = _index407
			(_2304_v149).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_this.F11, !(_this.F11))).Union(_2306_v150), _this.F11), (_index407).Int())
		}
		var _2308_v152 _dafny.CodePoint
		_ = _2308_v152
		_2308_v152 = _dafny.CodePoint('k')
		var _2309_v153 _dafny.Map
		_ = _2309_v153
		_2309_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _2308_v152)
		var _2310_v154 _dafny.Sequence
		_ = _2310_v154
		_2310_v154 = _dafny.SeqOf(p0, _dafny.IntOfUint32((_dafny.SeqOf(_2309_v153)).Cardinality()))
		var _2311_v155 D9
		_ = _2311_v155
		_2311_v155 = Companion_D9_.Create_DC23_((p0).Minus((_2310_v154).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(691), _dafny.IntOfUint32((_2310_v154).Cardinality()))).Uint32()).(_dafny.Int)), p0, p0, (p0).Plus(p0), _dafny.IntOfInt64(988))
		_2311_v155 = _2311_v155
		var _2312_v156 _dafny.Set
		_ = _2312_v156
		_2312_v156 = _dafny.SetOf((_dafny.Zero).Minus(p0))
		var _2313_v157 _dafny.MultiSet
		_ = _2313_v157
		_2313_v157 = _dafny.MultiSetOf(p0)
		var _2314_v158 _dafny.Sequence
		_ = _2314_v158
		_2314_v158 = _dafny.UnicodeSeqOfUtf8Bytes("gvrrji")
		var _2315_v159 _dafny.Map
		_ = _2315_v159
		_2315_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfUint32((_2310_v154).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _this.F11))
		var _2316_v160 _dafny.Map
		_ = _2316_v160
		_2316_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F11, _this.F11)
		var _2317_v161 _dafny.Map
		_ = _2317_v161
		_2317_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_2314_v158).Cardinality()), ((func() _dafny.Map {
			if (_2315_v159).Contains(_2310_v154) {
				return (_2315_v159).Get(_2310_v154).(_dafny.Map)
			}
			return _2316_v160
		})()).Cardinality())
		var _2318_v162 _dafny.Map
		_ = _2318_v162
		_2318_v162 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_D1_.Create_DC2_(p0, (_2312_v156).Cardinality(), _this.F11)), (_2313_v157).Update(_dafny.IntOfUint32((_2314_v158).Cardinality()), Companion_Default___.Abs((func() _dafny.Int {
			if (_2317_v161).Contains(p0) {
				return (_2317_v161).Get(p0).(_dafny.Int)
			}
			return p0
		})())))
		r0 = (_2318_v162).Merge((_2318_v162).Merge(_2318_v162))
		var _2319_v163 D2
		_ = _2319_v163
		_2319_v163 = Companion_D2_.Create_DC3_(_2191_v54)
		var _pat_let_tv61 = _2191_v54
		_ = _pat_let_tv61
		r1 = (func() D2 {
			if !(_this.F11) {
				return func(_pat_let68_0 D2) D2 {
					return func(_2320_dt__update__tmp_h0 D2) D2 {
						return func(_pat_let69_0 _dafny.Array) D2 {
							return func(_2321_dt__update_hcf5_h0 _dafny.Array) D2 {
								return Companion_D2_.Create_DC3_(_2321_dt__update_hcf5_h0)
							}(_pat_let69_0)
						}(_pat_let_tv61)
					}(_pat_let68_0)
				}(_2319_v163)
			}
			return _2319_v163
		})()
		return r0, r1
	}
}

// End of class C8

// Definition of class C9
type C9 struct {
	dummy byte
}

func New_C9_() *C9 {
	_this := C9{}

	return &_this
}

type CompanionStruct_C9_ struct {
}

var Companion_C9_ = CompanionStruct_C9_{}

func (_this *C9) Equals(other *C9) bool {
	return _this == other
}

func (_this *C9) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C9)
	return ok && _this.Equals(other)
}

func (*C9) String() string {
	return "_module.C9"
}

func Type_C9_() _dafny.TypeDescriptor {
	return type_C9_{}
}

type type_C9_ struct {
}

func (_this type_C9_) Default() interface{} {
	return (*C9)(nil)
}

func (_this type_C9_) String() string {
	return "main.C9"
}
func (_this *C9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C9{}

func (_this *C9) Ctor__() {
	{
	}
}
func (_this *C9) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	{
		var _source30 D1 = Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(260))).Uint32(), func(coer114 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg114 _dafny.Int) interface{} {
				return coer114(arg114)
			}
		}(func(_2322_i0 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)
		}))).Cardinality()), _dafny.IntOfInt64(-126), false)
		_ = _source30
		if _source30.Is_DC2() {
			var _2323___mcc_h0 _dafny.Int = _source30.Get_().(D1_DC2).Cf2
			_ = _2323___mcc_h0
			var _2324___mcc_h1 _dafny.Int = _source30.Get_().(D1_DC2).Cf3
			_ = _2324___mcc_h1
			var _2325___mcc_h2 bool = _source30.Get_().(D1_DC2).Cf4
			_ = _2325___mcc_h2
			var _2326_cf4 bool = _2325___mcc_h2
			_ = _2326_cf4
			var _2327_cf3 _dafny.Int = _2324___mcc_h1
			_ = _2327_cf3
			var _2328_cf2 _dafny.Int = _2323___mcc_h0
			_ = _2328_cf2
			return (_2328_cf2).Minus(_2328_cf2)
		} else {
			var _2329___mcc_h3 _dafny.Sequence = _source30.Get_().(D1_DC1).Cf1
			_ = _2329___mcc_h3
			var _2330_cf1 _dafny.Sequence = _2329___mcc_h3
			_ = _2330_cf1
			return _dafny.IntOfInt64(734)
		}
	}
}
func (_this *C9) Fm2(globalState *GlobalState) _dafny.MultiSet {
	{
		if ((_dafny.SetOf(true)).Cardinality()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vahhmfdxj")).Cardinality())) >= 0 {
			return _dafny.MultiSetOf(_dafny.CodePoint('h'))
		} else {
			return _dafny.MultiSetOf(_dafny.CodePoint('n'), _dafny.CodePoint('r'))
		}
	}
}
func (_this *C9) M1(p0 _dafny.Int, p1 _dafny.Array, globalState *GlobalState) (_dafny.Int, bool, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _2331_v0 bool
		_ = _2331_v0
		_2331_v0 = true
		var _2332_i0 _dafny.Int
		_ = _2332_i0
		_2332_i0 = _dafny.Zero
		{
			for _2331_v0 {
				{
					if (_2332_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L15
					}
					_2332_i0 = (_2332_i0).Plus(_dafny.One)
					var _index408 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((p1), 0))
					_ = _index408
					(p1).ArraySet1(_2331_v0, (_index408).Int())
					var _index409 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((p1), 0))
					_ = _index409
					(p1).ArraySet1((_dafny.IntOfInt64(51)).Cmp(p0) < 0, (_index409).Int())
					r3 = p0
					var _2333_v1 _dafny.Array
					_ = _2333_v1
					var _len0_62 _dafny.Int = _dafny.IntOfInt64(17)
					_ = _len0_62
					var _nw345 _dafny.Array
					_ = _nw345
					if _len0_62.Cmp(_dafny.Zero) == 0 {
						_nw345 = _dafny.NewArray(_len0_62)
					} else {
						var _init62 func(_dafny.Int) D1 = (func(_2334_p1 _dafny.Array) func(_dafny.Int) D1 {
							return func(_2335_i1 _dafny.Int) D1 {
								return (func() D1 {
									if (_2334_p1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(846), _dafny.ArrayLen((_2334_p1), 0))).Int()).(bool) {
										return Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("hrlebc"))
									}
									return Companion_D1_.Create_DC1_(_dafny.UnicodeSeqOfUtf8Bytes("houbdlof"))
								})()
							}
						})(p1)
						_ = _init62
						var _element0_62 = _init62(_dafny.Zero)
						_ = _element0_62
						_nw345 = _dafny.NewArrayFromExample(_element0_62, nil, _len0_62)
						(_nw345).ArraySet1(_element0_62, 0)
						var _nativeLen0_62 = (_len0_62).Int()
						_ = _nativeLen0_62
						for _i0_62 := 1; _i0_62 < _nativeLen0_62; _i0_62++ {
							(_nw345).ArraySet1(_init62(_dafny.IntOf(_i0_62)), _i0_62)
						}
					}
					_2333_v1 = _nw345
					var _2336_v2 _dafny.Sequence
					_ = _2336_v2
					_2336_v2 = _dafny.UnicodeSeqOfUtf8Bytes("rpfcdldk")
					var _2337_v3 D1
					_ = _2337_v3
					_2337_v3 = Companion_D1_.Create_DC1_(_2336_v2)
					var _index410 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_2333_v1), 0))
					_ = _index410
					(_2333_v1).ArraySet1(_2337_v3, (_index410).Int())
					var _2338_v4 _dafny.Map
					_ = _2338_v4
					_2338_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2331_v0, p0)
					var _2339_v5 _dafny.MultiSet
					_ = _2339_v5
					_2339_v5 = _dafny.MultiSetOf((_2338_v4).Update(Companion_Default___.Fm3(p0, globalState), p0))
					var _2340_v6 _dafny.Array
					_ = _2340_v6
					var _len0_63 _dafny.Int = _dafny.IntOfInt64(13)
					_ = _len0_63
					var _nw346 _dafny.Array
					_ = _nw346
					if _len0_63.Cmp(_dafny.Zero) == 0 {
						_nw346 = _dafny.NewArray(_len0_63)
					} else {
						var _init63 func(_dafny.Int) _dafny.Int = (func(_2341_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_2342_i2 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_2342_i2, _2341_p0)
							}
						})(p0)
						_ = _init63
						var _element0_63 = _init63(_dafny.Zero)
						_ = _element0_63
						_nw346 = _dafny.NewArrayFromExample(_element0_63, nil, _len0_63)
						(_nw346).ArraySet1(_element0_63, 0)
						var _nativeLen0_63 = (_len0_63).Int()
						_ = _nativeLen0_63
						for _i0_63 := 1; _i0_63 < _nativeLen0_63; _i0_63++ {
							(_nw346).ArraySet1(_init63(_dafny.IntOf(_i0_63)), _i0_63)
						}
					}
					_2340_v6 = _nw346
					var _2343_v7 _dafny.Sequence
					_ = _2343_v7
					_2343_v7 = _dafny.SeqOf(_2340_v6)
					var _index411 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_2333_v1), 0))
					_ = _index411
					var _rhs368 D1 = _2337_v3
					_ = _rhs368
					var _rhs369 _dafny.MultiSet = _2339_v5
					_ = _rhs369
					var _rhs370 _dafny.Sequence = (_2337_v3).Dtor_cf1()
					_ = _rhs370
					var _rhs371 bool = ((p0).Plus(_dafny.IntOfUint32((_2336_v2).Cardinality()))).Cmp(p0) >= 0
					_ = _rhs371
					var _rhs372 _dafny.Array = (_2343_v7).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p0), _dafny.IntOfUint32((_2343_v7).Cardinality()))).Uint32()).(_dafny.Array)
					_ = _rhs372
					var _lhs299 _dafny.Array = _2333_v1
					_ = _lhs299
					var _lhs300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_2333_v1), 0))
					_ = _lhs300
					(_lhs299).ArraySet1(_rhs368, (_lhs300).Int())
					_2339_v5 = _rhs369
					_2336_v2 = _rhs370
					r2 = _rhs371
					_2340_v6 = _rhs372
					r0 = p0
					goto C15
				}
			C15:
			}
			goto L15
		}
	L15:
		var _2344_v8 _dafny.Array
		_ = _2344_v8
		var _nw347 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
		_ = _nw347
		_2344_v8 = _nw347
		var _2345_v9 D2
		_ = _2345_v9
		_2345_v9 = Companion_D2_.Create_DC4_(_2344_v8, _2331_v0, p0, _2331_v0, p1)
		var _pat_let_tv62 = p0
		_ = _pat_let_tv62
		var _pat_let_tv63 = _2331_v0
		_ = _pat_let_tv63
		var _pat_let_tv64 = p1
		_ = _pat_let_tv64
		_2345_v9 = func(_pat_let70_0 D2) D2 {
			return func(_2346_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let71_0 _dafny.Int) D2 {
					return func(_2347_dt__update_hcf8_h0 _dafny.Int) D2 {
						return func(_pat_let72_0 bool) D2 {
							return func(_2348_dt__update_hcf9_h0 bool) D2 {
								return func(_pat_let73_0 _dafny.Array) D2 {
									return func(_2349_dt__update_hcf10_h0 _dafny.Array) D2 {
										return Companion_D2_.Create_DC4_((_2346_dt__update__tmp_h0).Dtor_cf6(), (_2346_dt__update__tmp_h0).Dtor_cf7(), _2347_dt__update_hcf8_h0, _2348_dt__update_hcf9_h0, _2349_dt__update_hcf10_h0)
									}(_pat_let73_0)
								}(_pat_let_tv64)
							}(_pat_let72_0)
						}(!(_pat_let_tv63))
					}(_pat_let71_0)
				}(_pat_let_tv62)
			}(_pat_let70_0)
		}(_2345_v9)
		var _2350_v10 _dafny.Map
		_ = _2350_v10
		_2350_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2331_v0, !(_2331_v0) || (_2331_v0))
		_2350_v10 = (_2350_v10).Update(_2331_v0, _2331_v0)
		r2 = _2331_v0
		var _2351_v11 _dafny.Sequence
		_ = _2351_v11
		_2351_v11 = _dafny.SeqOf(p0)
		var _hi16 _dafny.Int = (_2351_v11).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_2351_v11).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _hi16
		for _2352_i3 := p0; _2352_i3.Cmp(_hi16) < 0; _2352_i3 = _2352_i3.Plus(_dafny.One) {
			var _2353_v12 _dafny.Array
			_ = _2353_v12
			var _len0_64 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_64
			var _nw348 _dafny.Array
			_ = _nw348
			if _len0_64.Cmp(_dafny.Zero) == 0 {
				_nw348 = _dafny.NewArray(_len0_64)
			} else {
				var _init64 func(_dafny.Int) _dafny.Sequence = func(_2354_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("ouwtuvdwo")
				}
				_ = _init64
				var _element0_64 = _init64(_dafny.Zero)
				_ = _element0_64
				_nw348 = _dafny.NewArrayFromExample(_element0_64, nil, _len0_64)
				(_nw348).ArraySet1(_element0_64, 0)
				var _nativeLen0_64 = (_len0_64).Int()
				_ = _nativeLen0_64
				for _i0_64 := 1; _i0_64 < _nativeLen0_64; _i0_64++ {
					(_nw348).ArraySet1(_init64(_dafny.IntOf(_i0_64)), _i0_64)
				}
			}
			_2353_v12 = _nw348
			var _2355_v13 _dafny.Sequence
			_ = _2355_v13
			_2355_v13 = _dafny.UnicodeSeqOfUtf8Bytes("mm")
			var _2356_v14 _dafny.Map
			_ = _2356_v14
			_2356_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2353_v12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2355_v13, _2355_v13)).Cardinality()))
			_2356_v14 = (_2356_v14).Update(_2353_v12, p0)
			var _pat_let_tv65 = _2355_v13
			_ = _pat_let_tv65
			var _source31 D1 = func(_pat_let74_0 D1) D1 {
				return func(_2357_dt__update__tmp_h1 D1) D1 {
					return func(_pat_let75_0 _dafny.Sequence) D1 {
						return func(_2358_dt__update_hcf1_h0 _dafny.Sequence) D1 {
							return Companion_D1_.Create_DC1_(_2358_dt__update_hcf1_h0)
						}(_pat_let75_0)
					}(_pat_let_tv65)
				}(_pat_let74_0)
			}(Companion_D1_.Create_DC1_(_2355_v13))
			_ = _source31
			if _source31.Is_DC2() {
				var _2359___mcc_h0 _dafny.Int = _source31.Get_().(D1_DC2).Cf2
				_ = _2359___mcc_h0
				var _2360___mcc_h1 _dafny.Int = _source31.Get_().(D1_DC2).Cf3
				_ = _2360___mcc_h1
				var _2361___mcc_h2 bool = _source31.Get_().(D1_DC2).Cf4
				_ = _2361___mcc_h2
				var _2362_cf4 bool = _2361___mcc_h2
				_ = _2362_cf4
				var _2363_cf3 _dafny.Int = _2360___mcc_h1
				_ = _2363_cf3
				var _2364_cf2 _dafny.Int = _2359___mcc_h0
				_ = _2364_cf2
				r3 = _2364_cf2
				var _2365_v15 _dafny.MultiSet
				_ = _2365_v15
				_2365_v15 = _dafny.MultiSetOf(_dafny.IntOfInt64(931), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dhfqdpmbx")).Cardinality()))
				var _2366_v16 _dafny.Map
				_ = _2366_v16
				_2366_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_2365_v15).IsSubsetOf(_2365_v15), p0)
				_2366_v16 = (_2366_v16).Update(_2331_v0, _2352_i3)
				var _2367_v17 D1
				_ = _2367_v17
				_2367_v17 = Companion_D1_.Create_DC1_(_2355_v13)
				var _2368_v18 D1
				_ = _2368_v18
				_2368_v18 = Companion_D1_.Create_DC2_(_2364_cf2, _2363_cf3, _2331_v0)
				var _2369_v19 *C6
				_ = _2369_v19
				var _nw349 *C6 = New_C6_()
				_ = _nw349
				_nw349.Ctor__(_2367_v17, _2362_cf4, p1, _dafny.SetOf(_2363_cf3), _2368_v18, _2362_cf4)
				_2369_v19 = _nw349
				var _2370_v20 _dafny.Sequence
				_ = _2370_v20
				_2370_v20 = _dafny.SeqOf(_2369_v19, _2369_v19, _2369_v19, _2369_v19)
				var _2371_v21 _dafny.Sequence
				_ = _2371_v21
				_2371_v21 = _dafny.SeqOf(_2331_v0)
				var _2372_v22 _dafny.Set
				_ = _2372_v22
				_2372_v22 = _dafny.SetOf(p0, _2364_cf2, Companion_Default___.Fm0(_2371_v21, _2331_v0, true, _2352_i3, globalState))
				var _2373_v23 *C2
				_ = _2373_v23
				var _nw350 *C2 = New_C2_()
				_ = _nw350
				_nw350.Ctor__(_dafny.IntOfUint32((_2370_v20).Cardinality()), _2362_cf4, _2367_v17, _2362_cf4, p1, _2372_v22, _2368_v18, _2362_cf4)
				_2373_v23 = _nw350
				var _2374_v24 _dafny.Set
				_ = _2374_v24
				_2374_v24 = _dafny.SetOf(_2373_v23)
				var _2375_v25 _dafny.Sequence
				_ = _2375_v25
				_2375_v25 = _dafny.SeqOf(_2374_v24)
				var _2376_v26 _dafny.CodePoint
				_ = _2376_v26
				_2376_v26 = _dafny.CodePoint('l')
				var _2377_v27 *C0
				_ = _2377_v27
				var _nw351 *C0 = New_C0_()
				_ = _nw351
				_nw351.Ctor__((_dafny.SetOf(_2373_v23, _2373_v23)).IsProperSubsetOf((_2375_v25).Select((Companion_Default___.SafeIndex((_2373_v23).F21(), _dafny.IntOfUint32((_2375_v25).Cardinality()))).Uint32()).(_dafny.Set)), _2376_v26, _2368_v18, (_2373_v23).F22())
				_2377_v27 = _nw351
				var _2378_v28 _dafny.Sequence
				_ = _2378_v28
				_2378_v28 = _dafny.SeqOf(_2351_v11)
				var _2379_v29 D11
				_ = _2379_v29
				_2379_v29 = Companion_D11_.Create_DC26_(_2344_v8, _2355_v13)
				var _2380_v30 _dafny.Map
				_ = _2380_v30
				_2380_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_2352_i3).Cmp((Companion_D13_.Create_DC34_(_2352_i3, _2378_v28)).Dtor_cf71()) <= 0), _2379_v29)
				_2380_v30 = (_2380_v30).Update(_2331_v0, _2379_v29)
			} else {
				var _2381___mcc_h3 _dafny.Sequence = _source31.Get_().(D1_DC1).Cf1
				_ = _2381___mcc_h3
				var _2382_cf1 _dafny.Sequence = _2381___mcc_h3
				_ = _2382_cf1
				var _2383_v31 _dafny.Set
				_ = _2383_v31
				_2383_v31 = _dafny.SetOf(_2352_i3)
				var _2384_v32 D1
				_ = _2384_v32
				_2384_v32 = Companion_D1_.Create_DC2_(_2352_i3, _dafny.IntOfInt64(-87), _2331_v0)
				var _2385_v33 T3
				_ = _2385_v33
				var _nw352 *C2 = New_C2_()
				_ = _nw352
				_nw352.Ctor__(p0, _2331_v0, Companion_D1_.Create_DC1_(_2382_cf1), _2331_v0, p1, _2383_v31, _2384_v32, _2331_v0)
				_2385_v33 = _nw352
				var _2386_v34 _dafny.Map
				_ = _2386_v34
				_2386_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2352_i3, _2385_v33)
				_2386_v34 = _2386_v34
				var _2387_v35 *C4
				_ = _2387_v35
				var _nw353 *C4 = New_C4_()
				_ = _nw353
				_nw353.Ctor__()
				_2387_v35 = _nw353
				var _2388_v36 _dafny.Array
				_ = _2388_v36
				var _nw354 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(2))
				_ = _nw354
				_2388_v36 = _nw354
				var _2389_v37 D3
				_ = _2389_v37
				_2389_v37 = Companion_D3_.Create_DC8_(p0)
				var _index412 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_2388_v36), 0))
				_ = _index412
				(_2388_v36).ArraySet1(_2389_v37, (_index412).Int())
				var _index413 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.ArrayLen((_2388_v36), 0))
				_ = _index413
				(_2388_v36).ArraySet1(_2389_v37, (_index413).Int())
				var _2390_v38 _dafny.MultiSet
				_ = _2390_v38
				_2390_v38 = _dafny.MultiSetOf(_dafny.IntOfUint32((_2351_v11).Cardinality()), (_2383_v31).Cardinality(), _dafny.IntOfInt64(-180), _2352_i3)
				var _2391_v40 *C2
				_ = _2391_v40
				var _nw355 *C2 = New_C2_()
				_ = _nw355
				_nw355.Ctor__(_2352_i3, false, _2385_v33.F16(), !(_2331_v0), p1, _2383_v31, Companion_D1_.Create_DC2_(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer115 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg115 _dafny.Int) interface{} {
						return coer115(arg115)
					}
				}(func(_2392_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				}))).Cardinality()), (_2390_v38).Cardinality(), (_2385_v33).F13()), _dafny.Companion_Sequence_.Equal(_dafny.SeqOf((func() _dafny.Map {
					var _coll59 = _dafny.NewMapBuilder()
					_ = _coll59
					for _iter72 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(186), _dafny.IntOfInt64(82))); ; {
						_compr_59, _ok72 := _iter72()
						if !_ok72 {
							break
						}
						var _2393_v39 _dafny.Int
						_2393_v39 = interface{}(_compr_59).(_dafny.Int)
						if ((_dafny.IntOfInt64(186)).Cmp(_2393_v39) <= 0) && ((_2393_v39).Cmp(_dafny.IntOfInt64(82)) < 0) {
							_coll59.Add(Companion_Default___.SafeDivisionInt(_2393_v39, _2352_i3), (Companion_D15_.Create_DC39_(p0, _dafny.CodePoint('q'), _2355_v13, p0, _2352_i3)).Dtor_cf82())
						}
					}
					return _coll59.ToMap()
				}()).Cardinality()), _2351_v11))
				_2391_v40 = _nw355
				var _2394_v41 D9
				_ = _2394_v41
				_2394_v41 = Companion_D9_.Create_DC23_(_dafny.IntOfInt64(570), (_2352_i3).Plus(p0), (_2391_v40).F21(), _2352_i3, (_2391_v40).F21())
				var _rhs373 *C2 = _2391_v40
				_ = _rhs373
				var _rhs374 D9 = _2394_v41
				_ = _rhs374
				_2391_v40 = _rhs373
				_2394_v41 = _rhs374
			}
			var _index414 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_2344_v8), 0))
			_ = _index414
			(_2344_v8).ArraySet1((_dafny.Zero).Minus(_2352_i3), (_index414).Int())
			var _2395_v42 _dafny.Map
			_ = _2395_v42
			_2395_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2331_v0, _2355_v13)
			var _2396_v43 _dafny.Map
			_ = _2396_v43
			_2396_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Sequence {
				if (_2395_v42).Contains(!(_2331_v0)) {
					return (_2395_v42).Get(!(_2331_v0)).(_dafny.Sequence)
				}
				return _2355_v13
			})())
			var _index415 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(484), _dafny.ArrayLen((_2344_v8), 0))
			_ = _index415
			(_2344_v8).ArraySet1((Companion_Default___.SafeModuloInt(_2352_i3, (_2396_v43).Cardinality())).Plus(p0), (_index415).Int())
			var _2397_v44 _dafny.CodePoint
			_ = _2397_v44
			_2397_v44 = _dafny.CodePoint('a')
			var _2398_v45 D1
			_ = _2398_v45
			_2398_v45 = Companion_D1_.Create_DC2_(_2352_i3, p0, false)
			var _2399_v46 *C0
			_ = _2399_v46
			var _nw356 *C0 = New_C0_()
			_ = _nw356
			_nw356.Ctor__(_2331_v0, _2397_v44, _2398_v45, _2331_v0)
			_2399_v46 = _nw356
			var _2400_v47 _dafny.Map
			_ = _2400_v47
			_2400_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_2399_v46.F19, _2399_v46)
			var _2401_v48 D19
			_ = _2401_v48
			_2401_v48 = Companion_D19_.Create_DC48_(_2399_v46)
			var _2402_v49 _dafny.Array
			_ = _2402_v49
			var _nwElement0_77 *C0 = _2399_v46
			_ = _nwElement0_77
			var _nw357 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_77, nil, _dafny.IntOfInt64(25))
			_ = _nw357
			(_nw357).ArraySet1(_nwElement0_77, 0)
			(_nw357).ArraySet1(_2399_v46, 1)
			(_nw357).ArraySet1(_2399_v46, 2)
			(_nw357).ArraySet1(_2399_v46, 3)
			(_nw357).ArraySet1(_2399_v46, 4)
			(_nw357).ArraySet1(_2399_v46, 5)
			(_nw357).ArraySet1(_2399_v46, 6)
			(_nw357).ArraySet1(_2399_v46, 7)
			(_nw357).ArraySet1(_2399_v46, 8)
			(_nw357).ArraySet1(_2399_v46, 9)
			(_nw357).ArraySet1(_2399_v46, 10)
			(_nw357).ArraySet1(_2399_v46, 11)
			(_nw357).ArraySet1(_2399_v46, 12)
			(_nw357).ArraySet1(_2399_v46, 13)
			(_nw357).ArraySet1((func() *C0 {
				if (_2400_v47).Contains(true) {
					return (_2400_v47).Get(true).(*C0)
				}
				return _2399_v46
			})(), 14)
			(_nw357).ArraySet1(_2399_v46, 15)
			(_nw357).ArraySet1(_2399_v46, 16)
			(_nw357).ArraySet1((func() *C0 {
				if _2331_v0 {
					return _2399_v46
				}
				return _2399_v46
			})(), 17)
			(_nw357).ArraySet1(_2399_v46, 18)
			(_nw357).ArraySet1(_2399_v46, 19)
			(_nw357).ArraySet1(_2399_v46, 20)
			(_nw357).ArraySet1((_2401_v48).Dtor_cf92(), 21)
			(_nw357).ArraySet1(_2399_v46, 22)
			(_nw357).ArraySet1(_2399_v46, 23)
			(_nw357).ArraySet1(_2399_v46, 24)
			_2402_v49 = _nw357
			_2402_v49 = _2402_v49
		}
		var _2403_v50 *C4
		_ = _2403_v50
		var _nw358 *C4 = New_C4_()
		_ = _nw358
		_nw358.Ctor__()
		_2403_v50 = _nw358
		var _2404_v51 _dafny.MultiSet
		_ = _2404_v51
		_2404_v51 = _dafny.MultiSetOf(p0, p0, (_dafny.Zero).Minus(p0), (_dafny.Zero).Minus(p0), p0)
		r0 = (_2351_v11).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_2404_v51).Contains(_dafny.IntOfInt64(653)) {
				return (_2404_v51).Multiplicity(_dafny.IntOfInt64(653))
			}
			return p0
		})(), _dafny.IntOfUint32((_2351_v11).Cardinality()))).Uint32()).(_dafny.Int)
		r1 = _2331_v0
		r2 = (func() bool {
			if (_2350_v10).Contains(true) {
				return (_2350_v10).Get(true).(bool)
			}
			return (p0).Cmp(_dafny.IntOfInt64(940)) < 0
		})()
		var _2405_v52 _dafny.Sequence
		_ = _2405_v52
		_2405_v52 = _dafny.SeqOf(_2331_v0)
		var _2406_v53 _dafny.Sequence
		_ = _2406_v53
		_2406_v53 = _dafny.UnicodeSeqOfUtf8Bytes("sigql")
		var _2407_v54 _dafny.Map
		_ = _2407_v54
		_2407_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(_dafny.IntOfUint32((_2406_v53).Cardinality())))
		r3 = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_2405_v52, _2405_v52)).Cardinality())).Times((func() _dafny.Int {
			if _2331_v0 {
				return p0
			}
			return (_dafny.Zero).Minus((func() _dafny.Int {
				if (_2407_v54).Contains(p0) {
					return (_2407_v54).Get(p0).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p0)
			})())
		})())
		return r0, r1, r2, r3
	}
}

// End of class C9
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
