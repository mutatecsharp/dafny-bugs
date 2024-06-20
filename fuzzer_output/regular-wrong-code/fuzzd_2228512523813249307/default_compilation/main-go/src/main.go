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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Map, p2 _dafny.Int, globalState *GlobalState) bool {
	return false
}
func (_static *CompanionStruct_Default___) Fm1(p0 bool, globalState *GlobalState) _dafny.Int {
	return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jkbuevdy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(10))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_0_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	})))).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(895))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('m')
	}))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("irien"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(552))).Uint32(), func(coer2 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('x')
	}))), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nwfh"), _dafny.UnicodeSeqOfUtf8Bytes("hc")))
}
func (_static *CompanionStruct_Default___) Fm3(p0 bool, p1 _dafny.Int, p2 _dafny.Set, p3 _dafny.Map, globalState *GlobalState) D0 {
	if !(false) || (false) {
		return Companion_D0_.Create_DC0_(_dafny.SetOf(!(!(true)), false, !(true), true))
	} else {
		return Companion_D0_.Create_DC0_(_dafny.SetOf(true))
	}
}
func (_static *CompanionStruct_Default___) Fm5(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(-33), _dafny.IntOfInt64(87)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(67)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xjhywi")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm6(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(977)), false)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(88), _dafny.IntOfInt64(786)), false)))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC6_(Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(!(false))), !(false))), _dafny.UnicodeSeqOfUtf8Bytes("jm"), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(469))).Uint32(), func(coer3 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), true))
	})), _dafny.SeqOf(Companion_D1_.Create_DC2_(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(!(false))), false)).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _4_v0 D0
			_4_v0 = interface{}(_compr_0).(D0)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(!(false))), false)).Contains(_4_v0) {
				_coll0.Add(_4_v0, false)
			}
		}
		return _coll0.ToMap()
	}()), Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), false)), Companion_D1_.Create_DC2_(func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), Companion_D0_.Create_DC0_(_dafny.SetOf(true)))).Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _5_v1 D0
			_5_v1 = interface{}(_compr_1).(D0)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), Companion_D0_.Create_DC0_(_dafny.SetOf(true))), _5_v1) {
				_coll1.Add(_5_v1, true)
			}
		}
		return _coll1.ToMap()
	}()), Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(true, false, false, false)), !(true))), Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(true, true)), false))))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Sequence, p1 _dafny.CodePoint, p2 bool, p3 bool, globalState *GlobalState) _dafny.CodePoint {
	if (func() bool {
		if !(!(false)) {
			return false
		}
		return !(false)
	})() {
		return _dafny.CodePoint('f')
	} else {
		return _dafny.CodePoint('n')
	}
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(false)).Intersection(_dafny.SetOf(false, false, false, true))).Difference(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 bool, p2 _dafny.Sequence, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-398), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(Companion_D8_.Create_DC21_(Companion_D2_.Create_DC6_(Companion_D1_.Create_DC2_(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), !(false))).Keys().Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _6_v0 D0
			_6_v0 = interface{}(_compr_2).(D0)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), !(false))).Contains(_6_v0) {
				_coll2.Add(_6_v0, false)
			}
		}
		return _coll2.ToMap()
	}()), _dafny.UnicodeSeqOfUtf8Bytes("gjxpwm"), _dafny.IntOfInt64(833)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(689)), true)))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("e")).Cardinality()), (_dafny.SetOf(false, true, true)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(90), (_dafny.MultiSetFromSeq(_dafny.SeqOf(false, false))).Cardinality())).Cardinality())).Cardinality())), _dafny.IntOfInt64(505))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(793), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("uwa")).Cardinality()))))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)).Cardinality(), (_dafny.SetOf(false)).Cardinality())).Merge(func() _dafny.Map {
		var _coll3 = _dafny.NewMapBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(162), _dafny.IntOfInt64(62))); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _7_v1 _dafny.Int
			_7_v1 = interface{}(_compr_3).(_dafny.Int)
			if ((_dafny.IntOfInt64(162)).Cmp(_7_v1) <= 0) && ((_7_v1).Cmp(_dafny.IntOfInt64(62)) < 0) {
				_coll3.Add(Companion_Default___.SafeModuloInt(_7_v1, _dafny.IntOfInt64(859)), _dafny.IntOfInt64(979))
			}
		}
		return _coll3.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Difference(_dafny.MultiSetOf(true, !(!(false))))).Union((_dafny.MultiSetOf(false, true, false, !(!(true)))).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(true))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	if !(!_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("cdpuwaavw"), _dafny.UnicodeSeqOfUtf8Bytes("jyc"))) {
		return _dafny.CodePoint('o')
	} else {
		return _dafny.CodePoint('j')
	}
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.MultiSet, p1 bool, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-171), _dafny.IntOfInt64(923))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(451), _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(400))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg4 _dafny.Int) interface{} {
			return coer4(arg4)
		}
	}(func(_8_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(444)
	}))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(585))).Cardinality())).Cardinality())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm18(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(true, true, true)).Intersection(_dafny.SetOf(false))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(true))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Sequence, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('m')
}
func (_static *CompanionStruct_Default___) Fm21(globalState *GlobalState) _dafny.Sequence {
	var _source0 D9 = Companion_D9_.Create_DC22_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(210))).Uint32(), func(coer5 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg5 _dafny.Int) interface{} {
			return coer5(arg5)
		}
	}(func(_9_i0 _dafny.Int) _dafny.Int {
		return (_dafny.MultiSetOf(!(true))).Cardinality()
	})))
	_ = _source0
	if _source0.Is_DC23() {
		var _10___mcc_h0 bool = _source0.Get_().(D9_DC23).Cf46
		_ = _10___mcc_h0
		var _11___mcc_h1 _dafny.Int = _source0.Get_().(D9_DC23).Cf47
		_ = _11___mcc_h1
		var _12___mcc_h2 bool = _source0.Get_().(D9_DC23).Cf48
		_ = _12___mcc_h2
		var _13___mcc_h3 bool = _source0.Get_().(D9_DC23).Cf49
		_ = _13___mcc_h3
		var _14___mcc_h4 _dafny.Int = _source0.Get_().(D9_DC23).Cf50
		_ = _14___mcc_h4
		var _15_cf50 _dafny.Int = _14___mcc_h4
		_ = _15_cf50
		var _16_cf49 bool = _13___mcc_h3
		_ = _16_cf49
		var _17_cf48 bool = _12___mcc_h2
		_ = _17_cf48
		var _18_cf47 _dafny.Int = _11___mcc_h1
		_ = _18_cf47
		var _19_cf46 bool = _10___mcc_h0
		_ = _19_cf46
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_16_cf49, _17_cf48, !(_19_cf46)), _dafny.SeqOf(_19_cf46))
	} else {
		var _20___mcc_h5 _dafny.Sequence = _source0.Get_().(D9_DC22).Cf45
		_ = _20___mcc_h5
		var _21_cf45 _dafny.Sequence = _20___mcc_h5
		_ = _21_cf45
		return _dafny.SeqOf((Companion_D1_.Create_DC4_(true)).Dtor_cf8())
	}
}
func (_static *CompanionStruct_Default___) Fm22(p0 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(!(false), (_dafny.SetOf(_dafny.IntOfInt64(460))).IsSubsetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("wda")).Cardinality()))), !(((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cbv")).Cardinality()))).Cardinality()).Cmp(_dafny.IntOfInt64(-888)) > 0))
}
func (_static *CompanionStruct_Default___) Fm24(p0 bool, p1 D1, p2 _dafny.Int, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(true)), true))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) _dafny.CodePoint {
	var _source1 D7 = Companion_D7_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rrvjgmae")).Cardinality())))
	_ = _source1
	if _source1.Is_DC17() {
		var _22___mcc_h0 _dafny.Sequence = _source1.Get_().(D7_DC17).Cf31
		_ = _22___mcc_h0
		var _23___mcc_h1 _dafny.Set = _source1.Get_().(D7_DC17).Cf32
		_ = _23___mcc_h1
		var _24_cf32 _dafny.Set = _23___mcc_h1
		_ = _24_cf32
		var _25_cf31 _dafny.Sequence = _22___mcc_h0
		_ = _25_cf31
		return (_25_cf31).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.IntOfUint32((_25_cf31).Cardinality()))).Uint32()).(_dafny.CodePoint)
	} else {
		var _26___mcc_h2 _dafny.Map = _source1.Get_().(D7_DC16).Cf30
		_ = _26___mcc_h2
		var _27_cf30 _dafny.Map = _26___mcc_h2
		_ = _27_cf30
		return _dafny.CodePoint('a')
	}
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(-724))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(843))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_28_i0 _dafny.Int) _dafny.Set {
		return _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(true))).Cardinality())
	}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(298))).Uint32(), func(coer7 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_29_i1 _dafny.Int) _dafny.Set {
		return _dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cjmi")).Cardinality())), (_dafny.MultiSetOf(_dafny.IntOfInt64(53))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(756), !(true))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), _dafny.IntOfInt64(-95))
	})))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	var _source2 D12 = Companion_D12_.Create_DC30_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer8 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_30_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("ekdl")
	})))
	_ = _source2
	if _source2.Is_DC31() {
		var _31___mcc_h0 _dafny.Int = _source2.Get_().(D12_DC31).Cf60
		_ = _31___mcc_h0
		var _32_cf60 _dafny.Int = _31___mcc_h0
		_ = _32_cf60
		return _dafny.SeqOf(!(false))
	} else {
		var _33___mcc_h1 _dafny.Sequence = _source2.Get_().(D12_DC30).Cf59
		_ = _33___mcc_h1
		var _34_cf59 _dafny.Sequence = _33___mcc_h1
		_ = _34_cf59
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true)), _dafny.SeqOf(false, !(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Sequence, p1 _dafny.Int, globalState *GlobalState) D5 {
	return Companion_D5_.Create_DC12_(_dafny.IntOfInt64(806))
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll4 = _dafny.NewMapBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate((_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}(func(_35_i0 _dafny.Int) _dafny.Int {
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
		})))).Elements()); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _36_v0 _dafny.Sequence
			_36_v0 = interface{}(_compr_4).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(750))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg10 _dafny.Int) interface{} {
					return coer10(arg10)
				}
			}(func(_35_i0 _dafny.Int) _dafny.Int {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()
			}))), _36_v0) {
				_coll4.Add(_36_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cvfumpis")).Cardinality()))
			}
		}
		return _coll4.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(680), !(true))).Cardinality(), (_dafny.MultiSetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(false, true, false))).Cardinality())).Cardinality()), (func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(231), _dafny.IntOfInt64(-903))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _37_v1 _dafny.Int
			_37_v1 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(231)).Cmp(_37_v1) <= 0) && ((_37_v1).Cmp(_dafny.IntOfInt64(-903)) < 0) {
				_coll5.Add((_37_v1).Plus((_dafny.Zero).Minus((_dafny.SetOf(true)).Cardinality())), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality(), true)).Cardinality()))
			}
		}
		return _coll5.ToMap()
	}()).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_dafny.IntOfInt64(864)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('y'), true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-996))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_38_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('g')
	})), _dafny.UnicodeSeqOfUtf8Bytes("tsneoe")), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tmfqkanh")).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(true)).Difference(_dafny.SetOf((Companion_D10_.Create_DC26_(true)).Dtor_cf52()))).Intersection((func() _dafny.Set {
		if false {
			return _dafny.SetOf(!(false))
		}
		return _dafny.SetOf(false, false)
	})())
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) D9 {
	return Companion_D9_.Create_DC22_(_dafny.SeqOf(_dafny.IntOfInt64(998)))
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf(true, false, !(false), false, true)).Union(_dafny.MultiSetOf(false))
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.SetOf(_dafny.IntOfInt64(279), _dafny.IntOfInt64(407)), func() _dafny.Set {
		var _coll6 = _dafny.NewBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('c'))).Cardinality()), _dafny.IntOfInt64(-504))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _39_v0 _dafny.Int
			_39_v0 = interface{}(_compr_6).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('c'))).Cardinality()), _dafny.IntOfInt64(-504)), _39_v0) {
				_coll6.Add(Companion_Default___.SafeDivisionInt(_39_v0, _dafny.IntOfInt64(968)))
			}
		}
		return _coll6.ToSet()
	}())).Intersection(_dafny.SetOf(_dafny.SetOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-84))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg12 _dafny.Int) interface{} {
			return coer12(arg12)
		}
	}(func(_40_i0 _dafny.Int) _dafny.Int {
		return (_dafny.SetOf(false, true)).Cardinality()
	}))).Cardinality()), (func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-491), _dafny.IntOfInt64(-155))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _41_v1 _dafny.Int
			_41_v1 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(-491)).Cmp(_41_v1) <= 0) && ((_41_v1).Cmp(_dafny.IntOfInt64(-155)) < 0) {
				_coll7.Add(Companion_Default___.SafeModuloInt(_41_v1, _dafny.IntOfInt64(778)), _dafny.SeqOf(true, true))
			}
		}
		return _coll7.ToMap()
	}()).Cardinality())))).Difference(func() _dafny.Set {
		var _coll8 = _dafny.NewBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(731)))).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _42_v2 _dafny.Set
			_42_v2 = interface{}(_compr_8).(_dafny.Set)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.SetOf(_dafny.IntOfInt64(731))), _42_v2) {
				_coll8.Add(_42_v2)
			}
		}
		return _coll8.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfInt64(-469))), _dafny.IntOfInt64(655)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("emsbo")).Cardinality()), _dafny.IntOfInt64(330), _dafny.IntOfInt64(554), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))), _dafny.SeqOf((_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(132), true)).Cardinality(), (_dafny.MultiSetOf(!(true))).Cardinality())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(41), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("eqwwqd")).Cardinality()), _dafny.IntOfInt64(79), _dafny.IntOfInt64(710), _dafny.IntOfInt64(416))).Cardinality(), ((_dafny.SetOf(false, false, true)).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false, true, !(!(false)), false, !(!(false)))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Cardinality())), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(679), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(188))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}(func(_43_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('k')
	}))).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.Int, p1 D1, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(875))
}
func (_static *CompanionStruct_Default___) Fm38(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("xeomwdehc")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-764))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_44_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_45_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('g')
		}))
	})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer16 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_46_i2 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("unkma")
	}))))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) D10 {
	if (func() _dafny.Set {
		var _coll9 = _dafny.NewBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(935), _dafny.IntOfInt64(433))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _47_v0 _dafny.Int
				_47_v0 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(935)).Cmp(_47_v0) <= 0) && ((_47_v0).Cmp(_dafny.IntOfInt64(433)) < 0) {
					_coll10.Add(Companion_Default___.SafeDivisionInt(_47_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.One)).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('l')
					})), true)).Cardinality()), true)
				}
			}
			return _coll10.ToMap()
		}()).Cardinality(), false)).Keys().Elements()); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _49_v1 _dafny.Int
			_49_v1 = interface{}(_compr_9).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
				var _coll11 = _dafny.NewMapBuilder()
				_ = _coll11
				for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(935), _dafny.IntOfInt64(433))); ; {
					_compr_11, _ok11 := _iter11()
					if !_ok11 {
						break
					}
					var _47_v0 _dafny.Int
					_47_v0 = interface{}(_compr_11).(_dafny.Int)
					if ((_dafny.IntOfInt64(935)).Cmp(_47_v0) <= 0) && ((_47_v0).Cmp(_dafny.IntOfInt64(433)) < 0) {
						_coll11.Add(Companion_Default___.SafeDivisionInt(_47_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.One)).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg18 _dafny.Int) interface{} {
								return coer18(arg18)
							}
						}(func(_48_i0 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('l')
						})), true)).Cardinality()), true)
					}
				}
				return _coll11.ToMap()
			}()).Cardinality(), false)).Contains(_49_v1) {
				_coll9.Add((_49_v1).Times(_dafny.IntOfInt64(979)))
			}
		}
		return _coll9.ToSet()
	}()).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfInt64(737))) {
		return Companion_D10_.Create_DC24_(_dafny.CodePoint('o'))
	} else {
		return Companion_D10_.Create_DC24_(_dafny.CodePoint('l'))
	}
}
func (_static *CompanionStruct_Default___) Fm40(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), !(true)))).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(691), !(!(true))))
}
func (_static *CompanionStruct_Default___) M0(p0 D0, p1 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, _dafny.Array) {
	var r0 _dafny.Int = _dafny.Zero
	_ = r0
	var r1 bool = false
	_ = r1
	var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r2
	r1 = false
	var _50_v0 bool
	_ = _50_v0
	_50_v0 = false
	var _51_v1 _dafny.Map
	_ = _51_v1
	_51_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
	var _52_v2 D8
	_ = _52_v2
	_52_v2 = Companion_D8_.Create_DC20_(_50_v0, Companion_Default___.Fm0(_50_v0, (_51_v1).Update(p1, p1), p1, globalState), _50_v0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(_50_v0, _51_v1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("t")).Cardinality()), globalState), _dafny.IntOfInt64(515))).Contains(false))
	var _53_v3 _dafny.Array
	_ = _53_v3
	var _nw0 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(16))
	_ = _nw0
	_53_v3 = _nw0
	var _54_v4 _dafny.MultiSet
	_ = _54_v4
	_54_v4 = _dafny.MultiSetOf(_53_v3)
	_52_v2 = Companion_D8_.Create_DC20_(_50_v0, (_dafny.IntOfInt64(-148)).Cmp(p1) <= 0, true, (_54_v4).IsSubsetOf(_54_v4))
	var _55_v5 _dafny.CodePoint
	_ = _55_v5
	_55_v5 = _dafny.CodePoint('d')
	var _56_v6 T0
	_ = _56_v6
	var _nw1 *C1 = New_C1_()
	_ = _nw1
	_nw1.Ctor__(_55_v5, p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v0, _50_v0)).Cardinality())
	_56_v6 = _nw1
	var _57_v7 D8
	_ = _57_v7
	_57_v7 = Companion_D8_.Create_DC18_(_56_v6)
	var _pat_let_tv0 = _56_v6
	_ = _pat_let_tv0
	var _pat_let_tv1 = _56_v6
	_ = _pat_let_tv1
	var _source3 D8 = func(_pat_let0_0 D8) D8 {
		return func(_58_dt__update__tmp_h0 D8) D8 {
			return func(_pat_let1_0 T0) D8 {
				return func(_59_dt__update_hcf33_h0 T0) D8 {
					return Companion_D8_.Create_DC18_(_59_dt__update_hcf33_h0)
				}(_pat_let1_0)
			}((func() T0 {
				if false {
					return _pat_let_tv0
				}
				return _pat_let_tv1
			})())
		}(_pat_let0_0)
	}(_57_v7)
	_ = _source3
	if _source3.Is_DC19() {
		var _60___mcc_h0 bool = _source3.Get_().(D8_DC19).Cf34
		_ = _60___mcc_h0
		var _61___mcc_h1 bool = _source3.Get_().(D8_DC19).Cf35
		_ = _61___mcc_h1
		var _62___mcc_h2 _dafny.Set = _source3.Get_().(D8_DC19).Cf36
		_ = _62___mcc_h2
		var _63___mcc_h3 _dafny.Int = _source3.Get_().(D8_DC19).Cf37
		_ = _63___mcc_h3
		var _64___mcc_h4 _dafny.Int = _source3.Get_().(D8_DC19).Cf38
		_ = _64___mcc_h4
		var _65_cf38 _dafny.Int = _64___mcc_h4
		_ = _65_cf38
		var _66_cf37 _dafny.Int = _63___mcc_h3
		_ = _66_cf37
		var _67_cf36 _dafny.Set = _62___mcc_h2
		_ = _67_cf36
		var _68_cf35 bool = _61___mcc_h1
		_ = _68_cf35
		var _69_cf34 bool = _60___mcc_h0
		_ = _69_cf34
		var _70_v8 _dafny.Array
		_ = _70_v8
		var _nwElement0_0 bool = _69_cf34
		_ = _nwElement0_0
		var _nw2 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(22))
		_ = _nw2
		(_nw2).ArraySet1(_nwElement0_0, 0)
		(_nw2).ArraySet1(_68_cf35, 1)
		(_nw2).ArraySet1(_50_v0, 2)
		(_nw2).ArraySet1(_68_cf35, 3)
		(_nw2).ArraySet1(_68_cf35, 4)
		(_nw2).ArraySet1(_50_v0, 5)
		(_nw2).ArraySet1(_69_cf34, 6)
		(_nw2).ArraySet1(_69_cf34, 7)
		(_nw2).ArraySet1(_50_v0, 8)
		(_nw2).ArraySet1(_68_cf35, 9)
		(_nw2).ArraySet1(_69_cf34, 10)
		(_nw2).ArraySet1(false, 11)
		(_nw2).ArraySet1(_69_cf34, 12)
		(_nw2).ArraySet1(_50_v0, 13)
		(_nw2).ArraySet1(_69_cf34, 14)
		(_nw2).ArraySet1(_50_v0, 15)
		(_nw2).ArraySet1(_68_cf35, 16)
		(_nw2).ArraySet1(_69_cf34, 17)
		(_nw2).ArraySet1(_50_v0, 18)
		(_nw2).ArraySet1(_69_cf34, 19)
		(_nw2).ArraySet1(_50_v0, 20)
		(_nw2).ArraySet1(false, 21)
		_70_v8 = _nw2
		var _71_v9 _dafny.Map
		_ = _71_v9
		_71_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_cf34, _66_cf37)
		var _72_v10 _dafny.Map
		_ = _72_v10
		_72_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_71_v9).Cardinality(), _50_v0)
		var _73_v11 _dafny.Map
		_ = _73_v11
		_73_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v8, _72_v10)
		var _74_v12 *C4
		_ = _74_v12
		var _nw3 *C4 = New_C4_()
		_ = _nw3
		_nw3.Ctor__(_69_cf34, (_73_v11).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_70_v8, _72_v10)), (p1).Times(_56_v6.F7()), (_dafny.Zero).Minus((_56_v6).F6()))
		_74_v12 = _nw3
		(globalState).F1 = (p1).Minus((_56_v6).F6())
		(_56_v6).F7_set_((_56_v6).F6())
		if _50_v0 {
			var _75_v13 _dafny.Sequence
			_ = _75_v13
			_75_v13 = _dafny.SeqOf(_66_cf37)
			_71_v9 = (_71_v9).Update(_dafny.Companion_Sequence_.IsProperPrefixOf(_75_v13, _dafny.SeqOf(_dafny.IntOfInt64(93))), _65_cf38)
			_65_cf38 = ((_56_v6).F6()).Times(_56_v6.F7())
			var _76_v14 _dafny.MultiSet
			_ = _76_v14
			_76_v14 = _dafny.MultiSetOf((_56_v6).F6(), (_56_v6).F6())
			var _77_v15 _dafny.Sequence
			_ = _77_v15
			_77_v15 = _dafny.UnicodeSeqOfUtf8Bytes("ehcuiqm")
			var _78_v16 *C3
			_ = _78_v16
			var _nw4 *C3 = New_C3_()
			_ = _nw4
			_nw4.Ctor__(_56_v6.F7(), ((_76_v14).Cardinality()).Minus(_dafny.IntOfUint32((_77_v15).Cardinality())))
			_78_v16 = _nw4
			var _79_v17 _dafny.Sequence
			_ = _79_v17
			_79_v17 = _dafny.SeqOf((func() T0 {
				if _50_v0 {
					return _56_v6
				}
				return _56_v6
			})())
			var _rhs0 *C3 = _78_v16
			_ = _rhs0
			var _rhs1 _dafny.CodePoint = _55_v5
			_ = _rhs1
			var _rhs2 _dafny.Sequence = _79_v17
			_ = _rhs2
			var _rhs3 _dafny.Int = _56_v6.F7()
			_ = _rhs3
			var _lhs0 T0 = _56_v6
			_ = _lhs0
			_78_v16 = _rhs0
			_55_v5 = _rhs1
			_79_v17 = _rhs2
			_lhs0.F7_set_(_rhs3)
			var _80_v18 *C0
			_ = _80_v18
			var _nw5 *C0 = New_C0_()
			_ = _nw5
			_nw5.Ctor__((_56_v6.F7()).Cmp((_56_v6).F6()) < 0)
			_80_v18 = _nw5
			var _rhs4 _dafny.Sequence = _77_v15
			_ = _rhs4
			var _rhs5 _dafny.Int = _dafny.IntOfInt64(591)
			_ = _rhs5
			var _rhs6 *C0 = _80_v18
			_ = _rhs6
			var _rhs7 _dafny.Int = (_56_v6).F6()
			_ = _rhs7
			var _rhs8 T0 = _56_v6
			_ = _rhs8
			var _lhs1 T0 = _56_v6
			_ = _lhs1
			var _lhs2 *GlobalState = globalState
			_ = _lhs2
			_77_v15 = _rhs4
			_lhs1.F7_set_(_rhs5)
			_80_v18 = _rhs6
			_lhs2.F1 = _rhs7
			_56_v6 = _rhs8
			_55_v5 = _55_v5
		} else {
			var _81_v19 _dafny.Array
			_ = _81_v19
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
			_ = _nw6
			_81_v19 = _nw6
			_81_v19 = _81_v19
			var _82_v20 _dafny.Map
			_ = _82_v20
			_82_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _71_v9)
			_82_v20 = (_82_v20).Update(_55_v5, _71_v9)
			(globalState).F1 = _56_v6.F7()
			var _83_v21 D13
			_ = _83_v21
			_83_v21 = Companion_D13_.Create_DC32_(_51_v1)
			var _84_v22 _dafny.Set
			_ = _84_v22
			_84_v22 = _dafny.SetOf(false, false)
			var _85_v23 _dafny.Map
			_ = _85_v23
			_85_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0((_74_v12).F11(), (_83_v21).Dtor_cf61(), _dafny.IntOfInt64(764), globalState), _84_v22)
			var _86_v24 _dafny.Map
			_ = _86_v24
			_86_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_74_v12).F11(), (_74_v12).F11())
			_85_v23 = (_85_v23).Update((func() bool {
				if (_86_v24).Contains((func() bool {
					if (_72_v10).Contains(_56_v6.F7()) {
						return (_72_v10).Get(_56_v6.F7()).(bool)
					}
					return _69_cf34
				})()) {
					return (_86_v24).Get((func() bool {
						if (_72_v10).Contains(_56_v6.F7()) {
							return (_72_v10).Get(_56_v6.F7()).(bool)
						}
						return _69_cf34
					})()).(bool)
				}
				return _50_v0
			})(), _84_v22)
			_68_cf35 = _69_cf34
		}
	} else if _source3.Is_DC20() {
		var _87___mcc_h5 bool = _source3.Get_().(D8_DC20).Cf39
		_ = _87___mcc_h5
		var _88___mcc_h6 bool = _source3.Get_().(D8_DC20).Cf40
		_ = _88___mcc_h6
		var _89___mcc_h7 bool = _source3.Get_().(D8_DC20).Cf41
		_ = _89___mcc_h7
		var _90___mcc_h8 bool = _source3.Get_().(D8_DC20).Cf42
		_ = _90___mcc_h8
		var _91_cf42 bool = _90___mcc_h8
		_ = _91_cf42
		var _92_cf41 bool = _89___mcc_h7
		_ = _92_cf41
		var _93_cf40 bool = _88___mcc_h6
		_ = _93_cf40
		var _94_cf39 bool = _87___mcc_h5
		_ = _94_cf39
		var _95_v25 _dafny.Array
		_ = _95_v25
		var _nw7 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
		_ = _nw7
		_95_v25 = _nw7
		var _96_v26 *C0
		_ = _96_v26
		var _nw8 *C0 = New_C0_()
		_ = _nw8
		_nw8.Ctor__(_94_cf39)
		_96_v26 = _nw8
		var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_95_v25), 0))
		_ = _index0
		(_95_v25).ArraySet1(_96_v26, (_index0).Int())
		var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(5), _dafny.ArrayLen((_95_v25), 0))
		_ = _index1
		(_95_v25).ArraySet1(_96_v26, (_index1).Int())
		if ((_56_v6).F6()).Cmp((_dafny.Zero).Minus(_56_v6.F7())) <= 0 {
			var _97_v27 _dafny.Sequence
			_ = _97_v27
			_97_v27 = _dafny.UnicodeSeqOfUtf8Bytes("rcspe")
			(_96_v26).F10 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("bjfovcio"), (Companion_Default___.SafeIndex(_56_v6.F7(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjfovcio")).Cardinality()))).Uint32(), _55_v5), _dafny.Companion_Sequence_.Concatenate(_97_v27, _dafny.UnicodeSeqOfUtf8Bytes("xmouvrgh")))
			(_56_v6).F7_set_((_56_v6.F7()).Plus((_56_v6).F6()))
			var _98_v28 _dafny.MultiSet
			_ = _98_v28
			_98_v28 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(799))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_99_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_100_i0 _dafny.Int) _dafny.CodePoint {
					return _99_v5
				}
			})(_55_v5))))
			var _101_v29 _dafny.Sequence
			_ = _101_v29
			_101_v29 = _dafny.SeqOf(_91_cf42, _92_cf41)
			var _102_v30 _dafny.Sequence
			_ = _102_v30
			_102_v30 = _dafny.SeqOf((_56_v6).F6())
			var _rhs9 bool = _96_v26.F10
			_ = _rhs9
			var _rhs10 _dafny.Int = ((_56_v6).F6()).Times(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm35(_dafny.IntOfUint32((_101_v29).Cardinality()), (_56_v6).F6(), globalState), _102_v30)).Cardinality()))
			_ = _rhs10
			var _rhs11 _dafny.MultiSet = _98_v28
			_ = _rhs11
			var _rhs12 _dafny.Int = (_56_v6.F7()).Plus(((_56_v6).F6()).Plus(_56_v6.F7()))
			_ = _rhs12
			var _lhs3 *GlobalState = globalState
			_ = _lhs3
			_93_cf40 = _rhs9
			_lhs3.F1 = _rhs10
			_98_v28 = _rhs11
			r0 = _rhs12
			var _103_v31 _dafny.Array
			_ = _103_v31
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_0
			var _nw9 _dafny.Array
			_ = _nw9
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw9 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Set = (func(_104_v6 T0) func(_dafny.Int) _dafny.Set {
					return func(_105_i1 _dafny.Int) _dafny.Set {
						return _dafny.SetOf(_104_v6.F7(), (_104_v6).F6())
					}
				})(_56_v6)
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw9 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw9).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw9).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_103_v31 = _nw9
			var _106_v32 _dafny.Map
			_ = _106_v32
			_106_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (Companion_D11_.Create_DC28_(_97_v27, _103_v31, _93_cf40, !(_50_v0))).Dtor_cf54())
			_106_v32 = _106_v32
			var _107_v33 _dafny.Sequence
			_ = _107_v33
			_107_v33 = _dafny.SeqOf(_53_v3, _53_v3, _53_v3)
			_53_v3 = (_107_v33).Select((Companion_Default___.SafeIndex(_56_v6.F7(), _dafny.IntOfUint32((_107_v33).Cardinality()))).Uint32()).(_dafny.Array)
		} else {
			var _108_v34 _dafny.Sequence
			_ = _108_v34
			_108_v34 = _dafny.UnicodeSeqOfUtf8Bytes("tobbymm")
			var _109_v35 _dafny.Sequence
			_ = _109_v35
			_109_v35 = _dafny.SeqOf(_dafny.IntOfUint32((_108_v34).Cardinality()))
			var _110_v36 _dafny.Sequence
			_ = _110_v36
			_110_v36 = _dafny.SeqOf(_91_cf42, _50_v0, _50_v0, ((_109_v35).Select((Companion_Default___.SafeIndex(_56_v6.F7(), _dafny.IntOfUint32((_109_v35).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(_94_cf39)).Cardinality())) >= 0, !_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("wlngpj"), _55_v5))
			var _rhs13 T0 = (_57_v7).Dtor_cf33()
			_ = _rhs13
			var _rhs14 bool = !((_110_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.IntOfUint32((_110_v36).Cardinality()))).Uint32()).(bool))
			_ = _rhs14
			var _rhs15 _dafny.Sequence = _108_v34
			_ = _rhs15
			var _lhs4 *C0 = _96_v26
			_ = _lhs4
			_56_v6 = _rhs13
			_lhs4.F10 = _rhs14
			_108_v34 = _rhs15
			var _111_v37 _dafny.Array
			_ = _111_v37
			var _nw10 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw10
			_111_v37 = _nw10
			var _112_v38 _dafny.Map
			_ = _112_v38
			_112_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v6, _96_v26.F10)
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_111_v37), 0))
			_ = _index2
			(_111_v37).ArraySet1(Companion_Default___.Fm0(_50_v0, (_51_v1).Update((_112_v38).Cardinality(), _56_v6.F7()), (_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('f'), _92_cf41)).Cardinality()), globalState), (_index2).Int())
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(654), _dafny.ArrayLen((_111_v37), 0))
			_ = _index3
			(_111_v37).ArraySet1(_93_cf40, (_index3).Int())
			_108_v34 = _108_v34
			var _113_v39 _dafny.Map
			_ = _113_v39
			_113_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _91_cf42)
			var _114_v40 _dafny.Map
			_ = _114_v40
			_114_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v37, _113_v39)
			var _115_v41 *C4
			_ = _115_v41
			var _nw11 *C4 = New_C4_()
			_ = _nw11
			_nw11.Ctor__(_96_v26.F10, _114_v40, Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_108_v34).Cardinality()), (_dafny.Zero).Minus(_56_v6.F7())), Companion_Default___.Fm1(_91_cf42, globalState))
			_115_v41 = _nw11
			_115_v41 = _115_v41
			(_56_v6).F7_set_(_56_v6.F7())
		}
		var _116_v42 _dafny.Map
		_ = _116_v42
		_116_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_93_cf40)), _53_v3)
		var _117_v43 D1
		_ = _117_v43
		_117_v43 = Companion_D1_.Create_DC3_(_93_cf40, p1, (_116_v42).Cardinality())
		var _118_v44 _dafny.Map
		_ = _118_v44
		_118_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_cf41, (_56_v6).F6())
		var _119_v45 _dafny.Map
		_ = _119_v45
		_119_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v26.F10, _91_cf42)
		var _120_v46 _dafny.Map
		_ = _120_v46
		_120_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v43, (func() _dafny.Int {
			if (_118_v44).Contains((func() bool {
				if (_119_v45).Contains(_50_v0) {
					return (_119_v45).Get(_50_v0).(bool)
				}
				return _92_cf41
			})()) {
				return (_118_v44).Get((func() bool {
					if (_119_v45).Contains(_50_v0) {
						return (_119_v45).Get(_50_v0).(bool)
					}
					return _92_cf41
				})()).(_dafny.Int)
			}
			return _56_v6.F7()
		})())
		_120_v46 = (_120_v46).Update(_117_v43, (_56_v6.F7()).Times((_56_v6).F6()))
		var _121_v47 _dafny.Map
		_ = _121_v47
		_121_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v6, _53_v3)
		var _122_v48 _dafny.Array
		_ = _122_v48
		var _nwElement0_1 _dafny.Array = _53_v3
		_ = _nwElement0_1
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(26))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_1, 0)
		(_nw12).ArraySet1(_53_v3, 1)
		(_nw12).ArraySet1(_53_v3, 2)
		(_nw12).ArraySet1(_53_v3, 3)
		(_nw12).ArraySet1(_53_v3, 4)
		(_nw12).ArraySet1(_53_v3, 5)
		(_nw12).ArraySet1(_53_v3, 6)
		(_nw12).ArraySet1(_53_v3, 7)
		(_nw12).ArraySet1(_53_v3, 8)
		(_nw12).ArraySet1(_53_v3, 9)
		(_nw12).ArraySet1(_53_v3, 10)
		(_nw12).ArraySet1(_53_v3, 11)
		(_nw12).ArraySet1(_53_v3, 12)
		(_nw12).ArraySet1((func() _dafny.Array {
			if (_121_v47).Contains(_56_v6) {
				return (_121_v47).Get(_56_v6).(_dafny.Array)
			}
			return _53_v3
		})(), 13)
		(_nw12).ArraySet1(_53_v3, 14)
		(_nw12).ArraySet1(_53_v3, 15)
		(_nw12).ArraySet1(_53_v3, 16)
		(_nw12).ArraySet1(_53_v3, 17)
		(_nw12).ArraySet1(_53_v3, 18)
		(_nw12).ArraySet1(_53_v3, 19)
		(_nw12).ArraySet1(_53_v3, 20)
		(_nw12).ArraySet1(_53_v3, 21)
		(_nw12).ArraySet1(_53_v3, 22)
		(_nw12).ArraySet1(_53_v3, 23)
		(_nw12).ArraySet1(_53_v3, 24)
		(_nw12).ArraySet1(_53_v3, 25)
		_122_v48 = _nw12
		var _123_v49 _dafny.Sequence
		_ = _123_v49
		_123_v49 = _dafny.SeqOf(p1, p1, p1)
		var _124_v50 *C5
		_ = _124_v50
		var _nw13 *C5 = New_C5_()
		_ = _nw13
		_nw13.Ctor__(_122_v48, _123_v49, (_56_v6).F6(), (_56_v6).F6())
		_124_v50 = _nw13
	} else if _source3.Is_DC21() {
		var _125___mcc_h9 D2 = _source3.Get_().(D8_DC21).Cf43
		_ = _125___mcc_h9
		var _126___mcc_h10 _dafny.Map = _source3.Get_().(D8_DC21).Cf44
		_ = _126___mcc_h10
		var _127_cf44 _dafny.Map = _126___mcc_h10
		_ = _127_cf44
		var _128_cf43 D2 = _125___mcc_h9
		_ = _128_cf43
		var _129_v51 _dafny.MultiSet
		_ = _129_v51
		_129_v51 = _dafny.MultiSetOf(_50_v0)
		r1 = ((_129_v51).Intersection(_129_v51)).IsProperSubsetOf(_129_v51)
		if _50_v0 {
			var _130_v52 *C2
			_ = _130_v52
			var _nw14 *C2 = New_C2_()
			_ = _nw14
			_nw14.Ctor__((_56_v6).F6(), (_56_v6).F6())
			_130_v52 = _nw14
			var _131_v53 _dafny.Sequence
			_ = _131_v53
			_131_v53 = _dafny.SeqOf(p1)
			var _132_v54 D9
			_ = _132_v54
			_132_v54 = Companion_D9_.Create_DC22_(_131_v53)
			_132_v54 = _132_v54
			var _133_v55 _dafny.Map
			_ = _133_v55
			_133_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_50_v0, _50_v0, _50_v0), (_56_v6).F6())
			var _134_v56 _dafny.Set
			_ = _134_v56
			_134_v56 = _dafny.SetOf(_50_v0)
			_133_v55 = (_133_v55).Update(_134_v56, (_56_v6).F6())
			var _135_v57 *C3
			_ = _135_v57
			var _nw15 *C3 = New_C3_()
			_ = _nw15
			_nw15.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(_56_v6.F7())).Cardinality()), (_dafny.Zero).Minus((_56_v6).F6()))
			_135_v57 = _nw15
			var _136_v58 _dafny.Sequence
			_ = _136_v58
			_136_v58 = _dafny.SeqOf(_135_v57, _135_v57, _135_v57)
			var _137_v59 _dafny.Sequence
			_ = _137_v59
			_137_v59 = _dafny.UnicodeSeqOfUtf8Bytes("fxyo")
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_53_v3), 0))
			_ = _index4
			(_53_v3).ArraySet1(_dafny.IntOfUint32((_137_v59).Cardinality()), (_index4).Int())
			var _138_v60 _dafny.Array
			_ = _138_v60
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw16
			_138_v60 = _nw16
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_53_v3), 0))
			_ = _index5
			var _rhs16 _dafny.Sequence = _136_v58
			_ = _rhs16
			var _rhs17 _dafny.Int = Companion_Default___.SafeDivisionInt((_56_v6).F6(), _56_v6.F7())
			_ = _rhs17
			var _rhs18 _dafny.Array = (func() _dafny.Array {
				if (_56_v6.F7()).Cmp(_56_v6.F7()) <= 0 {
					return _138_v60
				}
				return _138_v60
			})()
			_ = _rhs18
			var _lhs5 _dafny.Array = _53_v3
			_ = _lhs5
			var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(952), _dafny.ArrayLen((_53_v3), 0))
			_ = _lhs6
			_136_v58 = _rhs16
			(_lhs5).ArraySet1(_rhs17, (_lhs6).Int())
			_138_v60 = _rhs18
			var _139_v61 _dafny.Map
			_ = _139_v61
			_139_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v0, _50_v0)
			var _140_v62 D0
			_ = _140_v62
			_140_v62 = Companion_D0_.Create_DC0_(_134_v56)
			var _141_v63 _dafny.Map
			_ = _141_v63
			_141_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v62, false)
			var _142_v64 D1
			_ = _142_v64
			_142_v64 = Companion_D1_.Create_DC2_(_141_v63)
			(_130_v52).M8(_139_v61, Companion_D2_.Create_DC6_(_142_v64, _137_v59, _dafny.IntOfInt64(-757)), _137_v59, globalState)
		} else {
			var _143_v65 _dafny.Array
			_ = _143_v65
			var _nw17 _dafny.Array = _dafny.NewArray(_dafny.One)
			_ = _nw17
			_143_v65 = _nw17
			var _nw18 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(9))
			_ = _nw18
			_143_v65 = _nw18
			var _144_v66 _dafny.Sequence
			_ = _144_v66
			_144_v66 = _dafny.SeqOf((_56_v6).F6())
			var _145_v67 _dafny.Map
			_ = _145_v67
			_145_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_51_v1, _144_v66)
			_144_v66 = (func() _dafny.Sequence {
				if (_145_v67).Contains((_51_v1).Merge(_51_v1)) {
					return (_145_v67).Get((_51_v1).Merge(_51_v1)).(_dafny.Sequence)
				}
				return _144_v66
			})()
			var _146_v68 _dafny.Sequence
			_ = _146_v68
			_146_v68 = _dafny.SeqOf(_50_v0, _50_v0, _50_v0)
			(_56_v6).F7_set_(_dafny.IntOfUint32((Companion_Default___.Fm2(_144_v66, _56_v6.F7(), _56_v6.F7(), (_dafny.MultiSetFromSeq(_146_v68)).IsProperSubsetOf(_129_v51), globalState)).Cardinality()))
			_50_v0 = false
			r1 = _50_v0
		}
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_53_v3), 0))
		_ = _index6
		(_53_v3).ArraySet1((_56_v6).F6(), (_index6).Int())
		var _147_v69 _dafny.Sequence
		_ = _147_v69
		_147_v69 = _dafny.UnicodeSeqOfUtf8Bytes("fummh")
		var _148_v70 _dafny.Sequence
		_ = _148_v70
		_148_v70 = _dafny.SeqOf((_56_v6).F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v6.F7(), _50_v0)).Cardinality())
		var _149_v71 _dafny.Sequence
		_ = _149_v71
		_149_v71 = _dafny.SeqOf(_dafny.IntOfUint32((_148_v70).Cardinality()))
		var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_53_v3), 0))
		_ = _index7
		var _rhs19 _dafny.Int = (_56_v6).F6()
		_ = _rhs19
		var _rhs20 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("pba")
		_ = _rhs20
		var _rhs21 _dafny.Array = _53_v3
		_ = _rhs21
		var _rhs22 _dafny.Sequence = _148_v70
		_ = _rhs22
		var _rhs23 bool = _50_v0
		_ = _rhs23
		var _lhs7 _dafny.Array = _53_v3
		_ = _lhs7
		var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(329), _dafny.ArrayLen((_53_v3), 0))
		_ = _lhs8
		(_lhs7).ArraySet1(_rhs19, (_lhs8).Int())
		_147_v69 = _rhs20
		_53_v3 = _rhs21
		_149_v71 = _rhs22
		r1 = _rhs23
		var _150_v72 _dafny.Array
		_ = _150_v72
		var _len0_1 _dafny.Int = _dafny.IntOfInt64(18)
		_ = _len0_1
		var _nw19 _dafny.Array
		_ = _nw19
		if _len0_1.Cmp(_dafny.Zero) == 0 {
			_nw19 = _dafny.NewArray(_len0_1)
		} else {
			var _init1 func(_dafny.Int) bool = (func(_151_v0 bool) func(_dafny.Int) bool {
				return func(_152_i2 _dafny.Int) bool {
					return _151_v0
				}
			})(_50_v0)
			_ = _init1
			var _element0_1 = _init1(_dafny.Zero)
			_ = _element0_1
			_nw19 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
			(_nw19).ArraySet1(_element0_1, 0)
			var _nativeLen0_1 = (_len0_1).Int()
			_ = _nativeLen0_1
			for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
				(_nw19).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
			}
		}
		_150_v72 = _nw19
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_150_v72), 0))
		_ = _index8
		(_150_v72).ArraySet1(!(_50_v0), (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(894), _dafny.ArrayLen((_150_v72), 0))
		_ = _index9
		(_150_v72).ArraySet1(_50_v0, (_index9).Int())
	} else {
		var _153___mcc_h11 T0 = _source3.Get_().(D8_DC18).Cf33
		_ = _153___mcc_h11
		var _154_cf33 T0 = _153___mcc_h11
		_ = _154_cf33
		var _155_v73 D10
		_ = _155_v73
		_155_v73 = Companion_D10_.Create_DC24_(_55_v5)
		var _pat_let_tv2 = _55_v5
		_ = _pat_let_tv2
		var _156_v74 _dafny.Map
		_ = _156_v74
		_156_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let2_0 D10) D10 {
			return func(_157_dt__update__tmp_h1 D10) D10 {
				return func(_pat_let3_0 _dafny.CodePoint) D10 {
					return func(_158_dt__update_hcf51_h0 _dafny.CodePoint) D10 {
						return Companion_D10_.Create_DC24_(_158_dt__update_hcf51_h0)
					}(_pat_let3_0)
				}(_pat_let_tv2)
			}(_pat_let2_0)
		}(_155_v73)).Dtor_cf51(), _154_cf33)
		var _159_v75 D8
		_ = _159_v75
		_159_v75 = Companion_D8_.Create_DC19_(true, !(_156_v74).Contains(_55_v5), _dafny.SetOf(_53_v3, _53_v3), (_154_cf33.F7()).Minus(p1), Companion_Default___.SafeDivisionInt(_56_v6.F7(), (_56_v6).F6()))
		var _source4 D8 = _159_v75
		_ = _source4
		if _source4.Is_DC19() {
			var _160___mcc_h12 bool = _source4.Get_().(D8_DC19).Cf34
			_ = _160___mcc_h12
			var _161___mcc_h13 bool = _source4.Get_().(D8_DC19).Cf35
			_ = _161___mcc_h13
			var _162___mcc_h14 _dafny.Set = _source4.Get_().(D8_DC19).Cf36
			_ = _162___mcc_h14
			var _163___mcc_h15 _dafny.Int = _source4.Get_().(D8_DC19).Cf37
			_ = _163___mcc_h15
			var _164___mcc_h16 _dafny.Int = _source4.Get_().(D8_DC19).Cf38
			_ = _164___mcc_h16
			var _165_cf38 _dafny.Int = _164___mcc_h16
			_ = _165_cf38
			var _166_cf37 _dafny.Int = _163___mcc_h15
			_ = _166_cf37
			var _167_cf36 _dafny.Set = _162___mcc_h14
			_ = _167_cf36
			var _168_cf35 bool = _161___mcc_h13
			_ = _168_cf35
			var _169_cf34 bool = _160___mcc_h12
			_ = _169_cf34
			var _170_v76 _dafny.Sequence
			_ = _170_v76
			_170_v76 = _dafny.UnicodeSeqOfUtf8Bytes("nvmlujfv")
			var _171_v77 _dafny.MultiSet
			_ = _171_v77
			_171_v77 = _dafny.MultiSetOf(_170_v76)
			_171_v77 = _171_v77
			var _172_v78 _dafny.Array
			_ = _172_v78
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_2
			var _nw20 _dafny.Array
			_ = _nw20
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw20 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_173_v0 bool) func(_dafny.Int) bool {
					return func(_174_i3 _dafny.Int) bool {
						return _173_v0
					}
				})(_50_v0)
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
			_172_v78 = _nw20
			var _175_v79 _dafny.Map
			_ = _175_v79
			_175_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_56_v6).F6(), !(_50_v0))
			var _176_v80 _dafny.Array
			_ = _176_v80
			var _nwElement0_2 _dafny.Array = _172_v78
			_ = _nwElement0_2
			var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(13))
			_ = _nw21
			(_nw21).ArraySet1(_nwElement0_2, 0)
			(_nw21).ArraySet1(_172_v78, 1)
			(_nw21).ArraySet1(_172_v78, 2)
			(_nw21).ArraySet1(_172_v78, 3)
			(_nw21).ArraySet1(_172_v78, 4)
			(_nw21).ArraySet1(_172_v78, 5)
			(_nw21).ArraySet1((func() _dafny.Array {
				if (func() bool {
					if (_175_v79).Contains((_154_cf33).F6()) {
						return (_175_v79).Get((_154_cf33).F6()).(bool)
					}
					return _169_cf34
				})() {
					return _172_v78
				}
				return _172_v78
			})(), 6)
			(_nw21).ArraySet1(_172_v78, 7)
			(_nw21).ArraySet1(_172_v78, 8)
			(_nw21).ArraySet1(_172_v78, 9)
			(_nw21).ArraySet1(_172_v78, 10)
			(_nw21).ArraySet1(_172_v78, 11)
			(_nw21).ArraySet1(_172_v78, 12)
			_176_v80 = _nw21
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_176_v80), 0))
			_ = _index10
			(_176_v80).ArraySet1(_172_v78, (_index10).Int())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(481), _dafny.ArrayLen((_176_v80), 0))
			_ = _index11
			(_176_v80).ArraySet1(_172_v78, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_53_v3), 0))
			_ = _index12
			(_53_v3).ArraySet1(Companion_Default___.SafeModuloInt(_56_v6.F7(), p1), (_index12).Int())
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_53_v3), 0))
			_ = _index13
			(_53_v3).ArraySet1(_56_v6.F7(), (_index13).Int())
			var _177_v81 _dafny.Map
			_ = _177_v81
			_177_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_55_v5, (_dafny.IntOfInt64(113)).Minus(_dafny.IntOfUint32((_170_v76).Cardinality())))
			(_154_cf33).F7_set_((_177_v81).Cardinality())
		} else if _source4.Is_DC20() {
			var _178___mcc_h17 bool = _source4.Get_().(D8_DC20).Cf39
			_ = _178___mcc_h17
			var _179___mcc_h18 bool = _source4.Get_().(D8_DC20).Cf40
			_ = _179___mcc_h18
			var _180___mcc_h19 bool = _source4.Get_().(D8_DC20).Cf41
			_ = _180___mcc_h19
			var _181___mcc_h20 bool = _source4.Get_().(D8_DC20).Cf42
			_ = _181___mcc_h20
			var _182_cf42 bool = _181___mcc_h20
			_ = _182_cf42
			var _183_cf41 bool = _180___mcc_h19
			_ = _183_cf41
			var _184_cf40 bool = _179___mcc_h18
			_ = _184_cf40
			var _185_cf39 bool = _178___mcc_h17
			_ = _185_cf39
			var _186_v82 _dafny.Map
			_ = _186_v82
			_186_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_182_cf42, (_56_v6).F6())
			var _187_v83 _dafny.Map
			_ = _187_v83
			_187_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_154_cf33.F7(), _185_cf39)
			_186_v82 = (_186_v82).Update((_56_v6.F7()).Cmp((_154_cf33).F6()) > 0, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((func() bool {
				if (_187_v83).Contains(_dafny.IntOfInt64(224)) {
					return (_187_v83).Get(_dafny.IntOfInt64(224)).(bool)
				}
				return false
			})(), _182_cf42, true)).Cardinality()), (_154_cf33).F6()))
			_51_v1 = (_51_v1).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_56_v6).F6(), _56_v6.F7()))
			(_56_v6).F7_set_((_56_v6).F6())
			var _188_v84 _dafny.Sequence
			_ = _188_v84
			_188_v84 = _dafny.UnicodeSeqOfUtf8Bytes("jiffwqiab")
			var _189_v85 _dafny.MultiSet
			_ = _189_v85
			_189_v85 = _dafny.MultiSetOf(((_51_v1).Update((_154_cf33).F6(), (_56_v6).F6())).Cardinality(), (_186_v82).Cardinality(), (_56_v6).F6())
			var _190_v86 _dafny.Map
			_ = _190_v86
			_190_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_189_v85, _188_v84)
			var _191_v87 _dafny.Array
			_ = _191_v87
			var _nwElement0_3 _dafny.Sequence = _188_v84
			_ = _nwElement0_3
			var _nw22 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(22))
			_ = _nw22
			(_nw22).ArraySet1(_nwElement0_3, 0)
			(_nw22).ArraySet1(_188_v84, 1)
			(_nw22).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_192_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_193_i4 _dafny.Int) _dafny.CodePoint {
					return _192_v5
				}
			})(_55_v5))), 2)
			(_nw22).ArraySet1(_188_v84, 3)
			(_nw22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("stfj"), 4)
			(_nw22).ArraySet1(_188_v84, 5)
			(_nw22).ArraySet1(_188_v84, 6)
			(_nw22).ArraySet1(_188_v84, 7)
			(_nw22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("iausxosc"), 8)
			(_nw22).ArraySet1(_188_v84, 9)
			(_nw22).ArraySet1(_188_v84, 10)
			(_nw22).ArraySet1(_188_v84, 11)
			(_nw22).ArraySet1(_188_v84, 12)
			(_nw22).ArraySet1(_188_v84, 13)
			(_nw22).ArraySet1(_188_v84, 14)
			(_nw22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("jjnked"), 15)
			(_nw22).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg21 _dafny.Int) interface{} {
					return coer21(arg21)
				}
			}(func(_194_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), (Companion_Default___.SafeIndex(_56_v6.F7(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(514))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_195_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			}))).Cardinality()))).Uint32(), _55_v5), 16)
			(_nw22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("laj"), 17)
			(_nw22).ArraySet1(_188_v84, 18)
			(_nw22).ArraySet1(_188_v84, 19)
			(_nw22).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("xt"), 20)
			(_nw22).ArraySet1((func() _dafny.Sequence {
				if (_190_v86).Contains(_189_v85) {
					return (_190_v86).Get(_189_v85).(_dafny.Sequence)
				}
				return _188_v84
			})(), 21)
			_191_v87 = _nw22
			var _196_v88 D13
			_ = _196_v88
			_196_v88 = Companion_D13_.Create_DC33_(_183_cf41, (_154_cf33).F6())
			var _197_v89 _dafny.Array
			_ = _197_v89
			var _nwElement0_4 _dafny.Array = _191_v87
			_ = _nwElement0_4
			var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(22))
			_ = _nw23
			(_nw23).ArraySet1(_nwElement0_4, 0)
			(_nw23).ArraySet1(_191_v87, 1)
			(_nw23).ArraySet1(_191_v87, 2)
			(_nw23).ArraySet1(_191_v87, 3)
			(_nw23).ArraySet1(_191_v87, 4)
			(_nw23).ArraySet1((func() _dafny.Array {
				if _182_cf42 {
					return _191_v87
				}
				return _191_v87
			})(), 5)
			(_nw23).ArraySet1(_191_v87, 6)
			(_nw23).ArraySet1(_191_v87, 7)
			(_nw23).ArraySet1(_191_v87, 8)
			(_nw23).ArraySet1(_191_v87, 9)
			(_nw23).ArraySet1(_191_v87, 10)
			(_nw23).ArraySet1(_191_v87, 11)
			(_nw23).ArraySet1(_191_v87, 12)
			(_nw23).ArraySet1(_191_v87, 13)
			(_nw23).ArraySet1(_191_v87, 14)
			(_nw23).ArraySet1(_191_v87, 15)
			(_nw23).ArraySet1(_191_v87, 16)
			(_nw23).ArraySet1((func() _dafny.Array {
				if (_196_v88).Dtor_cf62() {
					return _191_v87
				}
				return _191_v87
			})(), 17)
			(_nw23).ArraySet1(_191_v87, 18)
			(_nw23).ArraySet1(_191_v87, 19)
			(_nw23).ArraySet1((func() _dafny.Array {
				if _184_cf40 {
					return _191_v87
				}
				return _191_v87
			})(), 20)
			(_nw23).ArraySet1(_191_v87, 21)
			_197_v89 = _nw23
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_197_v89), 0))
			_ = _index14
			(_197_v89).ArraySet1(_191_v87, (_index14).Int())
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_197_v89), 0))
			_ = _index15
			(_197_v89).ArraySet1(_191_v87, (_index15).Int())
		} else if _source4.Is_DC21() {
			var _198___mcc_h21 D2 = _source4.Get_().(D8_DC21).Cf43
			_ = _198___mcc_h21
			var _199___mcc_h22 _dafny.Map = _source4.Get_().(D8_DC21).Cf44
			_ = _199___mcc_h22
			var _200_cf44 _dafny.Map = _199___mcc_h22
			_ = _200_cf44
			var _201_cf43 D2 = _198___mcc_h21
			_ = _201_cf43
			_55_v5 = _55_v5
			var _202_v90 D10
			_ = _202_v90
			_202_v90 = Companion_D10_.Create_DC25_()
			_202_v90 = (func() D10 {
				if _50_v0 {
					return _202_v90
				}
				return Companion_D10_.Create_DC25_()
			})()
			var _203_v91 _dafny.Array
			_ = _203_v91
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_3
			var _nw24 _dafny.Array
			_ = _nw24
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw24 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.CodePoint = (func(_204_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_205_i6 _dafny.Int) _dafny.CodePoint {
						return _204_v5
					}
				})(_55_v5)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw24 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw24).ArraySet1CodePoint(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw24).ArraySet1CodePoint(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_203_v91 = _nw24
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_203_v91), 0))
			_ = _index16
			(_203_v91).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _50_v0 {
					return _55_v5
				}
				return _55_v5
			})(), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(433), _dafny.ArrayLen((_203_v91), 0))
			_ = _index17
			(_203_v91).ArraySet1CodePoint(_55_v5, (_index17).Int())
			r1 = (_50_v0) || (false)
		} else {
			var _206___mcc_h23 T0 = _source4.Get_().(D8_DC18).Cf33
			_ = _206___mcc_h23
			var _207_cf33 T0 = _206___mcc_h23
			_ = _207_cf33
			var _208_v92 D10
			_ = _208_v92
			_208_v92 = Companion_D10_.Create_DC26_(_50_v0)
			_208_v92 = Companion_D10_.Create_DC26_(_50_v0)
			var _209_v93 _dafny.Sequence
			_ = _209_v93
			_209_v93 = _dafny.UnicodeSeqOfUtf8Bytes("bpmp")
			var _210_v94 _dafny.Set
			_ = _210_v94
			_210_v94 = _dafny.SetOf((_56_v6).F6(), (_207_cf33).F6(), (_56_v6).F6(), (_dafny.Zero).Minus(_56_v6.F7()))
			var _211_v95 _dafny.Array
			_ = _211_v95
			var _nwElement0_5 _dafny.Set = _210_v94
			_ = _nwElement0_5
			var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(4))
			_ = _nw25
			(_nw25).ArraySet1(_nwElement0_5, 0)
			(_nw25).ArraySet1(_210_v94, 1)
			(_nw25).ArraySet1(_210_v94, 2)
			(_nw25).ArraySet1(_210_v94, 3)
			_211_v95 = _nw25
			var _212_v96 D11
			_ = _212_v96
			_212_v96 = Companion_D11_.Create_DC28_(_209_v93, _211_v95, _50_v0, !(_50_v0))
			var _213_v97 _dafny.Set
			_ = _213_v97
			_213_v97 = _dafny.SetOf(_50_v0, (_212_v96).Dtor_cf56())
			_213_v97 = _dafny.SetOf(_50_v0)
			var _214_v98 _dafny.Array
			_ = _214_v98
			var _nw26 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(29))
			_ = _nw26
			_214_v98 = _nw26
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_214_v98), 0))
			_ = _index18
			(_214_v98).ArraySet1(_53_v3, (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_214_v98), 0))
			_ = _index19
			var _rhs24 _dafny.Array = _53_v3
			_ = _rhs24
			var _rhs25 _dafny.Sequence = _209_v93
			_ = _rhs25
			var _rhs26 bool = (!(_50_v0) || (_50_v0)) && (!(_50_v0) || (_50_v0))
			_ = _rhs26
			var _lhs9 _dafny.Array = _214_v98
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_214_v98), 0))
			_ = _lhs10
			(_lhs9).ArraySet1(_rhs24, (_lhs10).Int())
			_209_v93 = _rhs25
			_50_v0 = _rhs26
			var _215_v99 _dafny.Sequence
			_ = _215_v99
			_215_v99 = _dafny.SeqOf((_dafny.Zero).Minus((_207_cf33).F6()))
			var _216_v100 *C5
			_ = _216_v100
			var _nw27 *C5 = New_C5_()
			_ = _nw27
			_nw27.Ctor__(_214_v98, _215_v99, (_56_v6).F6(), (_56_v6).F6())
			_216_v100 = _nw27
			_216_v100 = _216_v100
		}
		var _217_v101 *C3
		_ = _217_v101
		var _nw28 *C3 = New_C3_()
		_ = _nw28
		_nw28.Ctor__(p1, (_154_cf33).F6())
		_217_v101 = _nw28
		var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_53_v3), 0))
		_ = _index20
		(_53_v3).ArraySet1((func() _dafny.Int {
			if false {
				return _56_v6.F7()
			}
			return (_56_v6).F6()
		})(), (_index20).Int())
		var _218_v102 _dafny.Array
		_ = _218_v102
		var _nwElement0_6 bool = _50_v0
		_ = _nwElement0_6
		var _nw29 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(13))
		_ = _nw29
		(_nw29).ArraySet1(_nwElement0_6, 0)
		(_nw29).ArraySet1(_50_v0, 1)
		(_nw29).ArraySet1(_50_v0, 2)
		(_nw29).ArraySet1(_50_v0, 3)
		(_nw29).ArraySet1(_50_v0, 4)
		(_nw29).ArraySet1((_159_v75).Dtor_cf35(), 5)
		(_nw29).ArraySet1(_50_v0, 6)
		(_nw29).ArraySet1(_50_v0, 7)
		(_nw29).ArraySet1(_50_v0, 8)
		(_nw29).ArraySet1(_50_v0, 9)
		(_nw29).ArraySet1(_50_v0, 10)
		(_nw29).ArraySet1(_50_v0, 11)
		(_nw29).ArraySet1(!((_217_v101).Fm16(globalState)), 12)
		_218_v102 = _nw29
		var _219_v103 _dafny.Set
		_ = _219_v103
		_219_v103 = _dafny.SetOf(_dafny.SeqOf(_217_v101))
		var _220_v104 _dafny.Map
		_ = _220_v104
		_220_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v103).Cardinality(), _50_v0)
		var _221_v105 _dafny.Map
		_ = _221_v105
		_221_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_218_v102, _220_v104)
		var _222_v106 *C4
		_ = _222_v106
		var _nw30 *C4 = New_C4_()
		_ = _nw30
		_nw30.Ctor__(_50_v0, _221_v105, Companion_Default___.Fm1(_50_v0, globalState), (_56_v6).F6())
		_222_v106 = _nw30
		var _223_v107 _dafny.Map
		_ = _223_v107
		_223_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_222_v106, true)
		var _224_v108 _dafny.Sequence
		_ = _224_v108
		_224_v108 = _dafny.SeqOf(_56_v6)
		var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_53_v3), 0))
		_ = _index21
		var _rhs27 bool = (func() bool {
			if _50_v0 {
				return (_56_v6.F7()).Cmp((_223_v107).Cardinality()) == 0
			}
			return (_222_v106).F11()
		})()
		_ = _rhs27
		var _rhs28 bool = (_222_v106).F11()
		_ = _rhs28
		var _rhs29 _dafny.Int = ((_154_cf33).F6()).Times(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_224_v108).Cardinality()), p1))
		_ = _rhs29
		var _rhs30 _dafny.Int = _154_cf33.F7()
		_ = _rhs30
		var _lhs11 T0 = _154_cf33
		_ = _lhs11
		var _lhs12 _dafny.Array = _53_v3
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_53_v3), 0))
		_ = _lhs13
		_50_v0 = _rhs27
		_50_v0 = _rhs28
		_lhs11.F7_set_(_rhs29)
		(_lhs12).ArraySet1(_rhs30, (_lhs13).Int())
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_218_v102), 0))
		_ = _index22
		(_218_v102).ArraySet1((_222_v106).F11(), (_index22).Int())
		var _225_v109 _dafny.Array
		_ = _225_v109
		var _nw31 _dafny.Array = _dafny.NewArray(_dafny.One)
		_ = _nw31
		_225_v109 = _nw31
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_225_v109), 0))
		_ = _index23
		(_225_v109).ArraySet1(_222_v106, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_218_v102), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_225_v109), 0))
		_ = _index25
		var _rhs31 bool = (_222_v106).F11()
		_ = _rhs31
		var _rhs32 *C4 = _222_v106
		_ = _rhs32
		var _rhs33 _dafny.Int = (_53_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen((_53_v3), 0))).Int()).(_dafny.Int)
		_ = _rhs33
		var _rhs34 _dafny.Int = Companion_Default___.SafeModuloInt(p1, (_dafny.Zero).Minus((_56_v6).F6()))
		_ = _rhs34
		var _lhs14 _dafny.Array = _218_v102
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(256), _dafny.ArrayLen((_218_v102), 0))
		_ = _lhs15
		var _lhs16 _dafny.Array = _225_v109
		_ = _lhs16
		var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_225_v109), 0))
		_ = _lhs17
		var _lhs18 T0 = _154_cf33
		_ = _lhs18
		var _lhs19 T0 = _56_v6
		_ = _lhs19
		(_lhs14).ArraySet1(_rhs31, (_lhs15).Int())
		(_lhs16).ArraySet1(_rhs32, (_lhs17).Int())
		_lhs18.F7_set_(_rhs33)
		_lhs19.F7_set_(_rhs34)
	}
	var _hi0 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p1, _56_v6.F7()))
	_ = _hi0
	for _226_i7 := p1; _226_i7.Cmp(_hi0) < 0; _226_i7 = _226_i7.Plus(_dafny.One) {
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
		_ = _index26
		(_53_v3).ArraySet1(_226_i7, (_index26).Int())
		var _227_v110 _dafny.Sequence
		_ = _227_v110
		_227_v110 = _dafny.SeqOf((_56_v6).F6())
		var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
		_ = _index27
		(_53_v3).ArraySet1((p1).Minus(_dafny.IntOfUint32((_227_v110).Cardinality())), (_index27).Int())
		var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
		_ = _index28
		(_53_v3).ArraySet1(_56_v6.F7(), (_index28).Int())
		var _228_v111 _dafny.Array
		_ = _228_v111
		var _nwElement0_7 T0 = _56_v6
		_ = _nwElement0_7
		var _nw32 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(13))
		_ = _nw32
		(_nw32).ArraySet1(_nwElement0_7, 0)
		(_nw32).ArraySet1(_56_v6, 1)
		(_nw32).ArraySet1(_56_v6, 2)
		(_nw32).ArraySet1(_56_v6, 3)
		(_nw32).ArraySet1(_56_v6, 4)
		(_nw32).ArraySet1(_56_v6, 5)
		(_nw32).ArraySet1(_56_v6, 6)
		(_nw32).ArraySet1((func() T0 {
			if _50_v0 {
				return _56_v6
			}
			return _56_v6
		})(), 7)
		(_nw32).ArraySet1(_56_v6, 8)
		(_nw32).ArraySet1(_56_v6, 9)
		(_nw32).ArraySet1(_56_v6, 10)
		(_nw32).ArraySet1(_56_v6, 11)
		(_nw32).ArraySet1(_56_v6, 12)
		_228_v111 = _nw32
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_228_v111), 0))
		_ = _index29
		(_228_v111).ArraySet1(_56_v6, (_index29).Int())
		var _229_v112 _dafny.Array
		_ = _229_v112
		var _nw33 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
		_ = _nw33
		_229_v112 = _nw33
		var _230_v113 _dafny.Map
		_ = _230_v113
		_230_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_229_v112, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_226_i7, !(!(_50_v0))))
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_228_v111), 0))
		_ = _index30
		var _nw34 *C4 = New_C4_()
		_ = _nw34
		_nw34.Ctor__(_50_v0, _230_v113, (p1).Times((_56_v6).F6()), Companion_Default___.SafeModuloInt((_56_v6).F6(), Companion_Default___.Fm1(_50_v0, globalState)))
		(_228_v111).ArraySet1(_nw34, (_index30).Int())
		if _50_v0 {
			var _231_v115 _dafny.Set
			_ = _231_v115
			_231_v115 = _dafny.SetOf((_56_v6).F6(), _56_v6.F7())
			var _232_v116 _dafny.Set
			_ = _232_v116
			_232_v116 = _dafny.SetOf(!((func() _dafny.Set {
				var _coll12 = _dafny.NewBuilder()
				_ = _coll12
				for _iter12 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(798), _dafny.IntOfInt64(348))); ; {
					_compr_12, _ok12 := _iter12()
					if !_ok12 {
						break
					}
					var _233_v114 _dafny.Int
					_233_v114 = interface{}(_compr_12).(_dafny.Int)
					if ((_dafny.IntOfInt64(798)).Cmp(_233_v114) <= 0) && ((_233_v114).Cmp(_dafny.IntOfInt64(348)) < 0) {
						_coll12.Add((_233_v114).Plus((_53_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))).Int()).(_dafny.Int)))
					}
				}
				return _coll12.ToSet()
			}()).IsProperSubsetOf(_231_v115)))
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
			_ = _index31
			(_53_v3).ArraySet1((_232_v116).Cardinality(), (_index31).Int())
			_50_v0 = (true) || (_50_v0)
			var _234_v117 _dafny.Sequence
			_ = _234_v117
			_234_v117 = _dafny.UnicodeSeqOfUtf8Bytes("nbbdvvot")
			var _235_v118 *C0
			_ = _235_v118
			var _nw35 *C0 = New_C0_()
			_ = _nw35
			_nw35.Ctor__(_50_v0)
			_235_v118 = _nw35
			var _236_v119 _dafny.Map
			_ = _236_v119
			_236_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_231_v115).Cardinality(), _235_v118)
			var _237_v120 *C1
			_ = _237_v120
			var _nw36 *C1 = New_C1_()
			_ = _nw36
			_nw36.Ctor__(Companion_Default___.Fm20(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(831))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}((func(_238_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_239_i8 _dafny.Int) _dafny.CodePoint {
					return _238_v5
				}
			})(_55_v5))), false, Companion_Default___.Fm1(_50_v0, globalState), globalState), (_dafny.IntOfUint32((_234_v117).Cardinality())).Plus((_236_v119).Cardinality()), (_56_v6).F6())
			_237_v120 = _nw36
			var _240_v121 _dafny.Array
			_ = _240_v121
			var _nwElement0_8 _dafny.CodePoint = (_237_v120).F13()
			_ = _nwElement0_8
			var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(4))
			_ = _nw37
			(_nw37).ArraySet1CodePoint(_nwElement0_8, 0)
			(_nw37).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _50_v0 {
					return _55_v5
				}
				return _55_v5
			})(), 1)
			(_nw37).ArraySet1CodePoint((_237_v120).F13(), 2)
			(_nw37).ArraySet1CodePoint((_237_v120).F13(), 3)
			_240_v121 = _nw37
			var _241_v122 _dafny.Array
			_ = _241_v122
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_4
			var _nw38 _dafny.Array
			_ = _nw38
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw38 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) _dafny.Int = (func(_242_v6 T0) func(_dafny.Int) _dafny.Int {
					return func(_243_i9 _dafny.Int) _dafny.Int {
						return (_243_i9).Times((_242_v6).F6())
					}
				})(_56_v6)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw38 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw38).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw38).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_241_v122 = _nw38
			var _244_v123 _dafny.Set
			_ = _244_v123
			_244_v123 = _dafny.SetOf(Companion_Default___.Fm14(_dafny.IntOfInt64(831), p1, _56_v6.F7(), globalState), (func() _dafny.CodePoint {
				if _235_v118.F10 {
					return _55_v5
				}
				return _55_v5
			})(), (_237_v120).F13(), (_237_v120).F13())
			var _245_v124 D2
			_ = _245_v124
			_245_v124 = Companion_D2_.Create_DC5_(_241_v122)
			var _rhs35 _dafny.Array = _240_v121
			_ = _rhs35
			var _rhs36 bool = _50_v0
			_ = _rhs36
			var _rhs37 _dafny.Array = (_245_v124).Dtor_cf9()
			_ = _rhs37
			var _rhs38 _dafny.Set = (_244_v123).Difference(_244_v123)
			_ = _rhs38
			_240_v121 = _rhs35
			_50_v0 = _rhs36
			_241_v122 = _rhs37
			_244_v123 = _rhs38
			var _246_v125 *C2
			_ = _246_v125
			var _nw39 *C2 = New_C2_()
			_ = _nw39
			_nw39.Ctor__(_56_v6.F7(), _dafny.IntOfUint32((_234_v117).Cardinality()))
			_246_v125 = _nw39
		} else {
			var _247_v126 *C0
			_ = _247_v126
			var _nw40 *C0 = New_C0_()
			_ = _nw40
			_nw40.Ctor__(_50_v0)
			_247_v126 = _nw40
			var _248_v127 _dafny.Array
			_ = _248_v127
			var _nwElement0_9 *C0 = _247_v126
			_ = _nwElement0_9
			var _nw41 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
			_ = _nw41
			(_nw41).ArraySet1(_nwElement0_9, 0)
			(_nw41).ArraySet1(_247_v126, 1)
			(_nw41).ArraySet1(_247_v126, 2)
			(_nw41).ArraySet1(_247_v126, 3)
			(_nw41).ArraySet1(_247_v126, 4)
			(_nw41).ArraySet1(_247_v126, 5)
			(_nw41).ArraySet1(_247_v126, 6)
			(_nw41).ArraySet1(_247_v126, 7)
			(_nw41).ArraySet1(_247_v126, 8)
			(_nw41).ArraySet1(_247_v126, 9)
			(_nw41).ArraySet1(_247_v126, 10)
			(_nw41).ArraySet1(_247_v126, 11)
			(_nw41).ArraySet1(_247_v126, 12)
			(_nw41).ArraySet1(_247_v126, 13)
			(_nw41).ArraySet1(_247_v126, 14)
			(_nw41).ArraySet1(_247_v126, 15)
			(_nw41).ArraySet1(_247_v126, 16)
			(_nw41).ArraySet1(_247_v126, 17)
			(_nw41).ArraySet1(_247_v126, 18)
			(_nw41).ArraySet1(_247_v126, 19)
			(_nw41).ArraySet1(_247_v126, 20)
			(_nw41).ArraySet1(_247_v126, 21)
			_248_v127 = _nw41
			var _249_v128 _dafny.Sequence
			_ = _249_v128
			_249_v128 = _dafny.SeqOf(_247_v126, _247_v126)
			var _250_v129 _dafny.Sequence
			_ = _250_v129
			_250_v129 = _dafny.UnicodeSeqOfUtf8Bytes("vrbqmsdid")
			var _251_v130 *C0
			_ = _251_v130
			_251_v130 = (_249_v128).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_250_v129).Cardinality()), _dafny.IntOfUint32((_249_v128).Cardinality()))).Uint32()).(*C0)
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_248_v127), 0))
			_ = _index32
			(_248_v127).ArraySet1((_251_v130), (_index32).Int())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_248_v127), 0))
			_ = _index33
			(_248_v127).ArraySet1(_247_v126, (_index33).Int())
			r1 = !(_50_v0)
			var _252_v131 _dafny.Array
			_ = _252_v131
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_5
			var _nw42 _dafny.Array
			_ = _nw42
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw42 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_253_v6 T0) func(_dafny.Int) _dafny.Int {
					return func(_254_i10 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_254_i10, (_253_v6).F6())
					}
				})(_56_v6)
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
			_252_v131 = _nw42
			var _255_v132 _dafny.Set
			_ = _255_v132
			_255_v132 = _dafny.SetOf(_252_v131)
			var _256_v133 D8
			_ = _256_v133
			_256_v133 = Companion_D8_.Create_DC19_(_50_v0, _50_v0, _255_v132, _56_v6.F7(), _dafny.IntOfInt64(130))
			var _257_v134 _dafny.Map
			_ = _257_v134
			_257_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v6, !((_256_v133).Dtor_cf34()))
			var _258_v135 D1
			_ = _258_v135
			_258_v135 = Companion_D1_.Create_DC3_(_50_v0, Companion_Default___.Fm1(_247_v126.F10, globalState), (_56_v6).F6())
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
			_ = _index34
			var _rhs39 _dafny.Int = (_227_v110).Select((Companion_Default___.SafeIndex((_53_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_227_v110).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs39
			var _rhs40 bool = _247_v126.F10
			_ = _rhs40
			var _rhs41 _dafny.Int = (func() _dafny.Int {
				if (func() bool {
					if (func() bool {
						if (_257_v134).Contains(_56_v6) {
							return (_257_v134).Get(_56_v6).(bool)
						}
						return _50_v0
					})() {
						return _50_v0
					}
					return _50_v0
				})() {
					return _dafny.IntOfUint32((_dafny.SeqOf((_258_v135).Dtor_cf5())).Cardinality())
				}
				return (_56_v6).F6()
			})()
			_ = _rhs41
			var _rhs42 _dafny.Int = Companion_Default___.Fm1(_247_v126.F10, globalState)
			_ = _rhs42
			var _lhs20 _dafny.Array = _53_v3
			_ = _lhs20
			var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_53_v3), 0))
			_ = _lhs21
			r0 = _rhs39
			r1 = _rhs40
			r0 = _rhs41
			(_lhs20).ArraySet1(_rhs42, (_lhs21).Int())
			var _nw43 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw43
			_252_v131 = _nw43
			(globalState).F1 = Companion_Default___.SafeModuloInt((_dafny.IntOfInt64(-750)).Minus((_56_v6).F6()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if _50_v0 {
					return _dafny.UnicodeSeqOfUtf8Bytes("wunqg")
				}
				return _250_v129
			})()).Cardinality()))
		}
	}
	var _hi1 _dafny.Int = _56_v6.F7()
	_ = _hi1
	for _259_i11 := (_56_v6.F7()).Minus(_56_v6.F7()); _259_i11.Cmp(_hi1) < 0; _259_i11 = _259_i11.Plus(_dafny.One) {
		var _260_v136 _dafny.Array
		_ = _260_v136
		var _nw44 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw44
		_260_v136 = _nw44
		var _261_v137 _dafny.Sequence
		_ = _261_v137
		_261_v137 = _dafny.SeqOf(_259_i11)
		var _262_v138 _dafny.Map
		_ = _262_v138
		_262_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_56_v6.F7(), _50_v0)
		var _263_v139 *C5
		_ = _263_v139
		var _nw45 *C5 = New_C5_()
		_ = _nw45
		_nw45.Ctor__(_260_v136, _261_v137, (func() _dafny.Int {
			if _50_v0 {
				return p1
			}
			return p1
		})(), (_262_v138).Cardinality())
		_263_v139 = _nw45
		_263_v139 = _263_v139
		var _264_v140 *C0
		_ = _264_v140
		var _nw46 *C0 = New_C0_()
		_ = _nw46
		_nw46.Ctor__(_50_v0)
		_264_v140 = _nw46
		var _265_v141 _dafny.Array
		_ = _265_v141
		var _nw47 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
		_ = _nw47
		_265_v141 = _nw47
		(_263_v139).M3(_265_v141, (_56_v6.F7()).Minus(p1), globalState)
		var _266_v142 _dafny.Array
		_ = _266_v142
		var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(18))
		_ = _nw48
		_266_v142 = _nw48
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_266_v142), 0))
		_ = _index35
		(_266_v142).ArraySet1((_dafny.SetOf(_264_v140.F10, _50_v0)).Intersection(_dafny.SetOf(false, (func() bool {
			if (_262_v138).Contains(_259_i11) {
				return (_262_v138).Get(_259_i11).(bool)
			}
			return false
		})())), (_index35).Int())
		var _267_v143 _dafny.Set
		_ = _267_v143
		_267_v143 = _dafny.SetOf(_50_v0, (_50_v0) || (true), (_264_v140.F10) && (_50_v0))
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(603), _dafny.ArrayLen((_266_v142), 0))
		_ = _index36
		(_266_v142).ArraySet1(_267_v143, (_index36).Int())
	}
	var _268_v144 _dafny.Sequence
	_ = _268_v144
	_268_v144 = _dafny.SeqOf(false, _50_v0, _50_v0)
	var _269_v145 _dafny.Map
	_ = _269_v145
	_269_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v144, !(_50_v0))
	var _270_v146 D5
	_ = _270_v146
	_270_v146 = Companion_D5_.Create_DC11_(_268_v144)
	_269_v145 = (_269_v145).Update((_270_v146).Dtor_cf20(), _50_v0)
	var _271_v147 _dafny.Set
	_ = _271_v147
	_271_v147 = _dafny.SetOf(_56_v6.F7(), _56_v6.F7())
	var _272_v148 _dafny.Sequence
	_ = _272_v148
	_272_v148 = _dafny.SeqOf((_271_v147).Cardinality(), _56_v6.F7())
	r0 = Companion_Default___.SafeDivisionInt(_56_v6.F7(), ((_272_v148).Select((Companion_Default___.SafeIndex((_56_v6).F6(), _dafny.IntOfUint32((_272_v148).Cardinality()))).Uint32()).(_dafny.Int)).Minus(p1))
	r1 = _50_v0
	var _273_v149 _dafny.Array
	_ = _273_v149
	var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(8))
	_ = _nw49
	_273_v149 = _nw49
	r2 = _273_v149
	return r0, r1, r2
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _274_globalState *GlobalState
	_ = _274_globalState
	var _nw50 *GlobalState = New_GlobalState_()
	_ = _nw50
	_nw50.Ctor__(true, _dafny.IntOfInt64(952), _dafny.IntOfInt64(374), _dafny.IntOfInt64(575), _dafny.IntOfInt64(555), false)
	_274_globalState = _nw50
	var _275_v0 _dafny.Int
	_ = _275_v0
	_275_v0 = _dafny.IntOfInt64(-379)
	(_274_globalState).F1 = Companion_Default___.SafeDivisionInt(_275_v0, (_275_v0).Minus(_275_v0))
	var _276_v1 bool
	_ = _276_v1
	_276_v1 = false
	(_274_globalState).F1 = (func() _dafny.Int {
		if _276_v1 {
			return (_275_v0).Times(_275_v0)
		}
		return _275_v0
	})()
	var _277_v2 _dafny.Array
	_ = _277_v2
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(17)
	_ = _len0_6
	var _nw51 _dafny.Array
	_ = _nw51
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw51 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = func(_278_i0 _dafny.Int) bool {
			return !(true)
		}
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw51 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw51).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw51).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_277_v2 = _nw51
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
	_ = _index37
	(_277_v2).ArraySet1(_276_v1, (_index37).Int())
	var _279_v3 _dafny.Set
	_ = _279_v3
	_279_v3 = _dafny.SetOf(_276_v1)
	var _280_v4 _dafny.Map
	_ = _280_v4
	_280_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_279_v3).Union(_279_v3), _276_v1)
	var _281_v5 _dafny.Array
	_ = _281_v5
	var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
	_ = _nw52
	_281_v5 = _nw52
	var _282_v6 _dafny.Set
	_ = _282_v6
	_282_v6 = _dafny.SetOf(_281_v5, _281_v5)
	var _283_v7 _dafny.Sequence
	_ = _283_v7
	_283_v7 = _dafny.UnicodeSeqOfUtf8Bytes("klhmhabdk")
	var _284_v8 _dafny.MultiSet
	_ = _284_v8
	_284_v8 = _dafny.MultiSetOf(_283_v7)
	var _285_v9 _dafny.Map
	_ = _285_v9
	_285_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_284_v8, _282_v6)
	var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
	_ = _index38
	var _rhs43 bool = _276_v1
	_ = _rhs43
	var _rhs44 bool = _276_v1
	_ = _rhs44
	var _rhs45 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_276_v1, _276_v1), _276_v1)
	_ = _rhs45
	var _rhs46 bool = (_282_v6).IsDisjointFrom((func() _dafny.Set {
		if (_285_v9).Contains(_dafny.MultiSetFromSeq(_dafny.SeqOf(_283_v7, _283_v7, _283_v7, _283_v7))) {
			return (_285_v9).Get(_dafny.MultiSetFromSeq(_dafny.SeqOf(_283_v7, _283_v7, _283_v7, _283_v7))).(_dafny.Set)
		}
		return _282_v6
	})())
	_ = _rhs46
	var _lhs22 _dafny.Array = _277_v2
	_ = _lhs22
	var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
	_ = _lhs23
	_276_v1 = _rhs43
	(_lhs22).ArraySet1(_rhs44, (_lhs23).Int())
	_280_v4 = _rhs45
	_276_v1 = _rhs46
	var _hi2 _dafny.Int = _275_v0
	_ = _hi2
	for _286_i1 := (_275_v0).Plus(_275_v0); _286_i1.Cmp(_hi2) < 0; _286_i1 = _286_i1.Plus(_dafny.One) {
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_281_v5), 0))
		_ = _index39
		(_281_v5).ArraySet1(_286_i1, (_index39).Int())
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_281_v5), 0))
		_ = _index40
		(_281_v5).ArraySet1(_286_i1, (_index40).Int())
		var _287_v10 _dafny.Map
		_ = _287_v10
		_287_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _286_i1)
		var _288_v11 _dafny.Map
		_ = _288_v11
		_288_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm0((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _287_v10, (_dafny.Zero).Minus(_286_i1), _274_globalState))
		_288_v11 = (_288_v11).Update(_276_v1, ((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)).Cmp(_286_i1) > 0)
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(990), _dafny.ArrayLen((_281_v5), 0))
		_ = _index41
		(_281_v5).ArraySet1(_275_v0, (_index41).Int())
		(_274_globalState).F1 = _275_v0
	}
	for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_281_v5), 0))); ; {
		_guard_loop_0, _ok13 := _iter13()
		if !_ok13 {
			break
		}
		var _289_i2 _dafny.Int
		_289_i2 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_289_i2).Sign() != -1) && ((_289_i2).Cmp(_dafny.ArrayLen((_281_v5), 0)) < 0)) {
			(_281_v5).ArraySet1((_289_i2).Minus(_275_v0), (_289_i2).Int())
		}
	}
	var _290_v12 _dafny.Map
	_ = _290_v12
	_290_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_275_v0, _275_v0)
	if (func() bool {
		if (func() bool {
			if _276_v1 {
				return _276_v1
			}
			return true
		})() {
			return (_276_v1) && (Companion_Default___.Fm0((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _290_v12, _275_v0, _274_globalState))
		}
		return (_275_v0).Cmp((_dafny.Zero).Minus(_275_v0)) == 0
	})() {
		var _291_v13 D0
		_ = _291_v13
		_291_v13 = Companion_D0_.Create_DC0_(_279_v3)
		var _292_v14 _dafny.MultiSet
		_ = _292_v14
		_292_v14 = _dafny.MultiSetOf(_279_v3, _279_v3, (_291_v13).Dtor_cf0())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_281_v5), 0))
		_ = _index42
		(_281_v5).ArraySet1(Companion_Default___.SafeDivisionInt(_275_v0, (func() _dafny.Int {
			if (_292_v14).Contains(_279_v3) {
				return (_292_v14).Multiplicity(_279_v3)
			}
			return _dafny.IntOfInt64(996)
		})()), (_index42).Int())
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_281_v5), 0))
		_ = _index43
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _index44
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _index45
		var _rhs47 _dafny.Int = (_275_v0).Minus(_275_v0)
		_ = _rhs47
		var _rhs48 _dafny.Int = (_275_v0).Plus(_275_v0)
		_ = _rhs48
		var _rhs49 bool = _276_v1
		_ = _rhs49
		var _rhs50 _dafny.Array = _277_v2
		_ = _rhs50
		var _rhs51 bool = _276_v1
		_ = _rhs51
		var _lhs24 _dafny.Array = _281_v5
		_ = _lhs24
		var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_281_v5), 0))
		_ = _lhs25
		var _lhs26 *GlobalState = _274_globalState
		_ = _lhs26
		var _lhs27 _dafny.Array = _277_v2
		_ = _lhs27
		var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _lhs28
		var _lhs29 _dafny.Array = _277_v2
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _lhs30
		(_lhs24).ArraySet1(_rhs47, (_lhs25).Int())
		_lhs26.F1 = _rhs48
		(_lhs27).ArraySet1(_rhs49, (_lhs28).Int())
		_277_v2 = _rhs50
		(_lhs29).ArraySet1(_rhs51, (_lhs30).Int())
		var _293_v15 _dafny.Array
		_ = _293_v15
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw53
		_293_v15 = _nw53
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_293_v15), 0))
		_ = _index46
		(_293_v15).ArraySet1(_dafny.SeqOf((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1(_276_v1, _274_globalState)), (_index46).Int())
		var _294_v16 _dafny.Sequence
		_ = _294_v16
		_294_v16 = _dafny.SeqOf(_275_v0)
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(659), _dafny.ArrayLen((_293_v15), 0))
		_ = _index47
		(_293_v15).ArraySet1(_294_v16, (_index47).Int())
		var _295_v17 D0
		_ = _295_v17
		_295_v17 = Companion_D0_.Create_DC1_(_276_v1, _277_v2, (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
		var _pat_let_tv3 = _277_v2
		_ = _pat_let_tv3
		var _pat_let_tv4 = _277_v2
		_ = _pat_let_tv4
		var _pat_let_tv5 = _277_v2
		_ = _pat_let_tv5
		var _296_v18 _dafny.Int
		_ = _296_v18
		var _297_v19 bool
		_ = _297_v19
		var _298_v20 _dafny.Array
		_ = _298_v20
		var _out0 _dafny.Int
		_ = _out0
		var _out1 bool
		_ = _out1
		var _out2 _dafny.Array
		_ = _out2
		_out0, _out1, _out2 = Companion_Default___.M0(func(_pat_let4_0 D0) D0 {
			return func(_302_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let8_0 bool) D0 {
					return func(_303_dt__update_hcf3_h1 bool) D0 {
						return Companion_D0_.Create_DC1_((_302_dt__update__tmp_h1).Dtor_cf1(), (_302_dt__update__tmp_h1).Dtor_cf2(), _303_dt__update_hcf3_h1)
					}(_pat_let8_0)
				}(false)
			}(_pat_let4_0)
		}(func(_pat_let5_0 D0) D0 {
			return func(_299_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let6_0 _dafny.Array) D0 {
					return func(_300_dt__update_hcf2_h0 _dafny.Array) D0 {
						return func(_pat_let7_0 bool) D0 {
							return func(_301_dt__update_hcf3_h0 bool) D0 {
								return Companion_D0_.Create_DC1_((_299_dt__update__tmp_h0).Dtor_cf1(), _300_dt__update_hcf2_h0, _301_dt__update_hcf3_h0)
							}(_pat_let7_0)
						}(!((_pat_let_tv5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_pat_let_tv4), 0))).Int()).(bool)))
					}(_pat_let6_0)
				}(_pat_let_tv3)
			}(_pat_let5_0)
		}(_295_v17)), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(30), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _274_globalState)
		_296_v18 = _out0
		_297_v19 = _out1
		_298_v20 = _out2
		_295_v17 = _295_v17
		var _304_v21 _dafny.Int
		_ = _304_v21
		var _305_v22 bool
		_ = _305_v22
		var _306_v23 _dafny.Array
		_ = _306_v23
		var _out3 _dafny.Int
		_ = _out3
		var _out4 bool
		_ = _out4
		var _out5 _dafny.Array
		_ = _out5
		_out3, _out4, _out5 = Companion_Default___.M0(_295_v17, _296_v18, _274_globalState)
		_304_v21 = _out3
		_305_v22 = _out4
		_306_v23 = _out5
	} else {
		_283_v7 = Companion_Default___.Fm2(_dafny.SeqOf(_275_v0, Companion_Default___.Fm1(Companion_Default___.Fm0((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _290_v12, (_dafny.Zero).Minus(_275_v0), _274_globalState), _274_globalState)), Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1(_276_v1, _274_globalState), _275_v0), _275_v0, !(_276_v1), _274_globalState)
		var _307_v24 _dafny.MultiSet
		_ = _307_v24
		_307_v24 = _dafny.MultiSetOf(_275_v0, _dafny.IntOfInt64(556))
		_276_v1 = ((_307_v24).IsProperSubsetOf(_307_v24)) || (_276_v1)
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _index48
		(_277_v2).ArraySet1(!((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool)) || ((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool)), (_index48).Int())
		(_274_globalState).F1 = _275_v0
		var _308_v25 _dafny.MultiSet
		_ = _308_v25
		_308_v25 = _dafny.MultiSetOf((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
		var _309_v26 _dafny.Map
		_ = _309_v26
		_309_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v1, (func() _dafny.Int {
			if (_308_v25).Contains(_276_v1) {
				return (_308_v25).Multiplicity(_276_v1)
			}
			return _275_v0
		})())
		var _310_v27 _dafny.Map
		_ = _310_v27
		_310_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_283_v7, (_309_v26).Cardinality())
		var _311_v28 D0
		_ = _311_v28
		_311_v28 = Companion_D0_.Create_DC0_(_dafny.SetOf(_276_v1))
		var _312_v29 _dafny.MultiSet
		_ = _312_v29
		_312_v29 = _dafny.MultiSetOf(Companion_Default___.Fm3(Companion_Default___.Fm0(_276_v1, _290_v12, _275_v0, _274_globalState), _275_v0, _279_v3, _310_v27, _274_globalState), _311_v28, Companion_D0_.Create_DC0_(_279_v3), _311_v28, _311_v28)
		(_274_globalState).F1 = (_312_v29).Cardinality()
	}
	for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_277_v2), 0))); ; {
		_guard_loop_1, _ok14 := _iter14()
		if !_ok14 {
			break
		}
		var _313_i3 _dafny.Int
		_313_i3 = interface{}(_guard_loop_1).(_dafny.Int)
		if (true) && (((_313_i3).Sign() != -1) && ((_313_i3).Cmp(_dafny.ArrayLen((_277_v2), 0)) < 0)) {
			(_277_v2).ArraySet1(false, (_313_i3).Int())
		}
	}
	var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))
	_ = _index49
	(_281_v5).ArraySet1(_dafny.IntOfUint32((_283_v7).Cardinality()), (_index49).Int())
	var _314_v30 _dafny.Set
	_ = _314_v30
	_314_v30 = _dafny.SetOf(_275_v0)
	var _315_v31 _dafny.Map
	_ = _315_v31
	_315_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_314_v30, _275_v0)
	var _316_v32 _dafny.Sequence
	_ = _316_v32
	_316_v32 = _dafny.SeqOf((_dafny.Zero).Minus(_275_v0), _dafny.IntOfInt64(405), _dafny.IntOfInt64(226), (_315_v31).Cardinality())
	var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))
	_ = _index50
	(_281_v5).ArraySet1(_dafny.IntOfUint32((_316_v32).Cardinality()), (_index50).Int())
	var _317_v33 _dafny.Map
	_ = _317_v33
	_317_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_275_v0, (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
	var _318_v34 _dafny.Map
	_ = _318_v34
	_318_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_277_v2, _317_v33)
	var _319_v35 *C4
	_ = _319_v35
	var _nw54 *C4 = New_C4_()
	_ = _nw54
	_nw54.Ctor__(_276_v1, _318_v34, _dafny.IntOfInt64(745), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
	_319_v35 = _nw54
	(_274_globalState).F1 = _dafny.IntOfInt64(-937)
	var _320_v37 D7
	_ = _320_v37
	_320_v37 = Companion_D7_.Create_DC16_(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter15 := _dafny.Iterate((_283_v7).Elements()); ; {
			_compr_13, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _321_v36 _dafny.CodePoint
			_321_v36 = interface{}(_compr_13).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_283_v7, _321_v36) {
				_coll13.Add(_321_v36, _275_v0)
			}
		}
		return _coll13.ToMap()
	}())
	var _322_v38 _dafny.Map
	_ = _322_v38
	_322_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _320_v37)
	_322_v38 = _322_v38
	var _323_v39 _dafny.MultiSet
	_ = _323_v39
	_323_v39 = _dafny.MultiSetOf((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _275_v0)
	var _hi3 _dafny.Int = (_323_v39).Cardinality()
	_ = _hi3
	for _324_i4 := (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int); _324_i4.Cmp(_hi3) < 0; _324_i4 = _324_i4.Plus(_dafny.One) {
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))
		_ = _index51
		(_281_v5).ArraySet1((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), (_index51).Int())
		var _325_v40 _dafny.Sequence
		_ = _325_v40
		_325_v40 = _dafny.SeqOf((_319_v35).F11(), (_319_v35).F11())
		var _326_v41 _dafny.CodePoint
		_ = _326_v41
		_326_v41 = _dafny.CodePoint('h')
		var _327_v42 _dafny.Array
		_ = _327_v42
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
		_ = _nw55
		_327_v42 = _nw55
		var _328_v43 _dafny.Map
		_ = _328_v43
		_328_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
		var _329_v44 *C5
		_ = _329_v44
		var _nw56 *C5 = New_C5_()
		_ = _nw56
		_nw56.Ctor__(_327_v42, _dafny.SeqOf((_328_v43).Cardinality(), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)), _324_i4, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
		_329_v44 = _nw56
		var _330_v45 _dafny.MultiSet
		_ = _330_v45
		_330_v45 = _dafny.MultiSetOf(_329_v44)
		var _331_v46 D5
		_ = _331_v46
		_331_v46 = Companion_D5_.Create_DC11_(_325_v40)
		var _332_v47 D8
		_ = _332_v47
		_332_v47 = Companion_D8_.Create_DC19_(_276_v1, (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _282_v6, _275_v0, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
		var _333_v48 _dafny.Array
		_ = _333_v48
		var _nwElement0_10 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_325_v40, _325_v40)
		_ = _nwElement0_10
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_10, 0)
		(_nw57).ArraySet1(_325_v40, 1)
		(_nw57).ArraySet1(_325_v40, 2)
		(_nw57).ArraySet1(_325_v40, 3)
		(_nw57).ArraySet1(_dafny.SeqOf((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool)), 4)
		(_nw57).ArraySet1(_325_v40, 5)
		(_nw57).ArraySet1((func() _dafny.Sequence {
			if _276_v1 {
				return _325_v40
			}
			return _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_319_v35).F11(), (_319_v35).F11()), (Companion_Default___.SafeIndex(_275_v0, _dafny.IntOfUint32((_dafny.SeqOf((_319_v35).F11(), (_319_v35).F11())).Cardinality()))).Uint32(), _276_v1)
		})(), 6)
		(_nw57).ArraySet1(_325_v40, 7)
		(_nw57).ArraySet1(_325_v40, 8)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Update(Companion_Default___.Fm27(_dafny.IntOfInt64(-187), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-901))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_334_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _326_v41, _274_globalState), (Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_330_v45).Contains(_329_v44) {
				return (_330_v45).Multiplicity(_329_v44)
			}
			return _dafny.IntOfInt64(538)
		})(), _dafny.IntOfUint32((Companion_Default___.Fm27(_dafny.IntOfInt64(-187), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-901))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_335_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		})), _326_v41, _274_globalState)).Cardinality()))).Uint32(), (_319_v35).F11()), 9)
		(_nw57).ArraySet1(_325_v40, 10)
		(_nw57).ArraySet1(_325_v40, 11)
		(_nw57).ArraySet1(_325_v40, 12)
		(_nw57).ArraySet1(_325_v40, 13)
		(_nw57).ArraySet1(_325_v40, 14)
		(_nw57).ArraySet1(_325_v40, 15)
		(_nw57).ArraySet1((_331_v46).Dtor_cf20(), 16)
		(_nw57).ArraySet1(_325_v40, 17)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_325_v40, _325_v40), 18)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_325_v40, _325_v40), 19)
		(_nw57).ArraySet1(_325_v40, 20)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_325_v40, _325_v40), 21)
		(_nw57).ArraySet1(Companion_Default___.Fm21(_274_globalState), 22)
		(_nw57).ArraySet1(_325_v40, 23)
		(_nw57).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_325_v40, _dafny.SeqOf((_332_v47).Dtor_cf34())), 24)
		(_nw57).ArraySet1(_dafny.SeqOf((_319_v35).F11()), 25)
		(_nw57).ArraySet1(_325_v40, 26)
		(_nw57).ArraySet1((func() _dafny.Sequence {
			if (_319_v35).F11() {
				return _325_v40
			}
			return _325_v40
		})(), 27)
		(_nw57).ArraySet1((_331_v46).Dtor_cf20(), 28)
		_333_v48 = _nw57
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_333_v48), 0))
		_ = _index52
		(_333_v48).ArraySet1(_325_v40, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_333_v48), 0))
		_ = _index53
		(_333_v48).ArraySet1(_325_v40, (_index53).Int())
		var _336_v49 D0
		_ = _336_v49
		_336_v49 = Companion_D0_.Create_DC0_(_279_v3)
		var _337_v50 _dafny.Map
		_ = _337_v50
		_337_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_336_v49, _276_v1)
		var _338_v51 D1
		_ = _338_v51
		_338_v51 = Companion_D1_.Create_DC2_(_337_v50)
		var _339_v52 D2
		_ = _339_v52
		_339_v52 = Companion_D2_.Create_DC6_(_338_v51, _dafny.UnicodeSeqOfUtf8Bytes("vku"), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
		var _340_v53 _dafny.Map
		_ = _340_v53
		_340_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v1, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
		var _341_v54 *C1
		_ = _341_v54
		var _nw58 *C1 = New_C1_()
		_ = _nw58
		_nw58.Ctor__(_326_v41, _324_i4, (func() _dafny.Int {
			if (_340_v53).Contains((_319_v35).F11()) {
				return (_340_v53).Get((_319_v35).F11()).(_dafny.Int)
			}
			return _dafny.IntOfUint32((_316_v32).Cardinality())
		})())
		_341_v54 = _nw58
		var _342_v55 _dafny.Map
		_ = _342_v55
		_342_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_341_v54, _283_v7)
		_283_v7 = _dafny.Companion_Sequence_.Concatenate((_339_v52).Dtor_cf11(), (func() _dafny.Sequence {
			if (_342_v55).Contains(_341_v54) {
				return (_342_v55).Get(_341_v54).(_dafny.Sequence)
			}
			return _283_v7
		})())
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_7
		var _nw59 _dafny.Array
		_ = _nw59
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw59 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) bool = (func(_343_v2 _dafny.Array) func(_dafny.Int) bool {
				return func(_344_i6 _dafny.Int) bool {
					return (_343_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_343_v2), 0))).Int()).(bool)
				}
			})(_277_v2)
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
		_277_v2 = _nw59
	}
	var _345_i7 _dafny.Int
	_ = _345_i7
	_345_i7 = _dafny.Zero
	{
		for (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool) {
			{
				if (_345_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_345_i7 = (_345_i7).Plus(_dafny.One)
				(_274_globalState).F1 = (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)
				var _346_v56 _dafny.Array
				_ = _346_v56
				var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_8
				var _nw60 _dafny.Array
				_ = _nw60
				if _len0_8.Cmp(_dafny.Zero) == 0 {
					_nw60 = _dafny.NewArray(_len0_8)
				} else {
					var _init8 func(_dafny.Int) _dafny.Sequence = (func(_347_v35 *C4) func(_dafny.Int) _dafny.Sequence {
						return func(_348_i8 _dafny.Int) _dafny.Sequence {
							return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf((_347_v35).F11(), (_347_v35).F11()))
						}
					})(_319_v35)
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
				_346_v56 = _nw60
				var _349_v57 _dafny.Sequence
				_ = _349_v57
				_349_v57 = _dafny.SeqOf(true, _276_v1)
				var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_346_v56), 0))
				_ = _index54
				(_346_v56).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_349_v57, _349_v57), (_index54).Int())
				var _350_v58 D1
				_ = _350_v58
				_350_v58 = Companion_D1_.Create_DC4_(_276_v1)
				var _351_v59 D2
				_ = _351_v59
				_351_v59 = Companion_D2_.Create_DC7_((_319_v35).F11(), (_279_v3).Cardinality(), _317_v33, _350_v58)
				var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_346_v56), 0))
				_ = _index55
				(_346_v56).ArraySet1(_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
					if !(Companion_Default___.Fm0((_351_v59).Dtor_cf13(), _290_v12, _dafny.IntOfUint32((_283_v7).Cardinality()), _274_globalState)) {
						return _349_v57
					}
					return _349_v57
				})(), _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_276_v1, (func() bool {
					if (_317_v33).Contains(_275_v0) {
						return (_317_v33).Get(_275_v0).(bool)
					}
					return (_319_v35).F11()
				})()), (Companion_Default___.SafeIndex(_275_v0, _dafny.IntOfUint32((_dafny.SeqOf(_276_v1, (func() bool {
					if (_317_v33).Contains(_275_v0) {
						return (_317_v33).Get(_275_v0).(bool)
					}
					return (_319_v35).F11()
				})())).Cardinality()))).Uint32(), true)), (_index55).Int())
				var _352_v60 *C0
				_ = _352_v60
				var _nw61 *C0 = New_C0_()
				_ = _nw61
				_nw61.Ctor__((_314_v30).IsSubsetOf(_dafny.SetOf((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))))
				_352_v60 = _nw61
				var _353_v61 T0
				_ = _353_v61
				var _nw62 *C4 = New_C4_()
				_ = _nw62
				_nw62.Ctor__((_319_v35).F11(), _319_v35.F12, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _274_globalState))
				_353_v61 = _nw62
				var _354_v62 _dafny.MultiSet
				_ = _354_v62
				_354_v62 = _dafny.MultiSetOf(_353_v61, _353_v61, _353_v61, _353_v61)
				var _355_v63 _dafny.Map
				_ = _355_v63
				_355_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_354_v62, false)
				var _356_v64 T0
				_ = _356_v64
				var _nw63 *C4 = New_C4_()
				_ = _nw63
				_nw63.Ctor__(Companion_Default___.Fm0((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _290_v12, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_349_v57, (Companion_Default___.SafeIndex(_275_v0, _dafny.IntOfUint32((_349_v57).Cardinality()))).Uint32(), (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))).Cardinality()), _274_globalState), _318_v34, _275_v0, ((_355_v63).Merge(_355_v63)).Cardinality())
				_356_v64 = _nw63
				var _357_v65 _dafny.CodePoint
				_ = _357_v65
				_357_v65 = _dafny.CodePoint('k')
				var _rhs52 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_283_v7, _dafny.SeqOf(_357_v65, Companion_Default___.Fm8(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(828))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg26 _dafny.Int) interface{} {
						return coer26(arg26)
					}
				}((func(_358_v65 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_359_i9 _dafny.Int) _dafny.CodePoint {
						return _358_v65
					}
				})(_357_v65))), _dafny.CodePoint('p'), false, (_319_v35).F11(), _274_globalState), _357_v65, _357_v65))
				_ = _rhs52
				var _rhs53 T0 = _353_v61
				_ = _rhs53
				_283_v7 = _rhs52
				_356_v64 = _rhs53
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
	_ = _index56
	(_277_v2).ArraySet1((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), (_index56).Int())
	var _360_v66 D0
	_ = _360_v66
	_360_v66 = Companion_D0_.Create_DC0_(_279_v3)
	var _361_v67 _dafny.Map
	_ = _361_v67
	_361_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_360_v66, (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
	var _362_v68 D1
	_ = _362_v68
	_362_v68 = Companion_D1_.Create_DC2_(_361_v67)
	var _source5 D1 = _362_v68
	_ = _source5
	if _source5.Is_DC3() {
		var _363___mcc_h0 bool = _source5.Get_().(D1_DC3).Cf5
		_ = _363___mcc_h0
		var _364___mcc_h1 _dafny.Int = _source5.Get_().(D1_DC3).Cf6
		_ = _364___mcc_h1
		var _365___mcc_h2 _dafny.Int = _source5.Get_().(D1_DC3).Cf7
		_ = _365___mcc_h2
		var _366_cf7 _dafny.Int = _365___mcc_h2
		_ = _366_cf7
		var _367_cf6 _dafny.Int = _364___mcc_h1
		_ = _367_cf6
		var _368_cf5 bool = _363___mcc_h0
		_ = _368_cf5
		var _369_v69 _dafny.Map
		_ = _369_v69
		_369_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_368_cf5, Companion_Default___.Fm1(false, _274_globalState))
		var _370_v70 bool
		_ = _370_v70
		var _371_v71 bool
		_ = _371_v71
		var _372_v72 _dafny.Int
		_ = _372_v72
		var _out6 bool
		_ = _out6
		var _out7 bool
		_ = _out7
		var _out8 _dafny.Int
		_ = _out8
		_out6, _out7, _out8 = (_319_v35).M1(!(!((_366_cf7).Cmp((func() _dafny.Int {
			if (_369_v69).Contains(false) {
				return (_369_v69).Get(false).(_dafny.Int)
			}
			return _367_cf6
		})()) <= 0)), _dafny.IntOfUint32((_283_v7).Cardinality()), _dafny.IntOfInt64(103), _274_globalState)
		_370_v70 = _out6
		_371_v71 = _out7
		_372_v72 = _out8
		var _373_v73 _dafny.Sequence
		_ = _373_v73
		_373_v73 = _dafny.SeqOf(Companion_Default___.Fm0(!(_368_cf5), _290_v12, _367_cf6, _274_globalState), _276_v1, _370_v70)
		var _374_v74 _dafny.Sequence
		_ = _374_v74
		_374_v74 = _dafny.SeqOf(_314_v30, _314_v30)
		var _375_v75 bool
		_ = _375_v75
		var _376_v76 bool
		_ = _376_v76
		var _377_v77 _dafny.Int
		_ = _377_v77
		var _out9 bool
		_ = _out9
		var _out10 bool
		_ = _out10
		var _out11 _dafny.Int
		_ = _out11
		_out9, _out10, _out11 = (_319_v35).M1((_373_v73).Select((Companion_Default___.SafeIndex(_275_v0, _dafny.IntOfUint32((_373_v73).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.IntOfUint32((_374_v74).Cardinality())).Minus(_367_cf6))), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_366_cf7, _367_cf6)).Cardinality(), _274_globalState)
		_375_v75 = _out9
		_376_v76 = _out10
		_377_v77 = _out11
		(_319_v35).M5((_dafny.Zero).Minus((func() _dafny.Int {
			if _368_cf5 {
				return (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)
			}
			return _275_v0
		})()), _274_globalState)
		var _378_v78 _dafny.CodePoint
		_ = _378_v78
		_378_v78 = _dafny.CodePoint('q')
		var _379_v79 _dafny.Set
		_ = _379_v79
		_379_v79 = _dafny.SetOf(_dafny.MultiSetOf(_378_v78))
		var _380_v80 _dafny.MultiSet
		_ = _380_v80
		_380_v80 = _dafny.MultiSetOf(_378_v78, _378_v78, _378_v78, _378_v78, _378_v78)
		_379_v79 = _dafny.SetOf(_380_v80, _380_v80)
	} else if _source5.Is_DC4() {
		var _381___mcc_h3 bool = _source5.Get_().(D1_DC4).Cf8
		_ = _381___mcc_h3
		var _382_cf8 bool = _381___mcc_h3
		_ = _382_cf8
		(_319_v35).M5(_275_v0, _274_globalState)
		var _383_v81 *C2
		_ = _383_v81
		var _nw64 *C2 = New_C2_()
		_ = _nw64
		_nw64.Ctor__((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _275_v0)
		_383_v81 = _nw64
		_277_v2 = _277_v2
		var _384_v82 _dafny.MultiSet
		_ = _384_v82
		_384_v82 = _dafny.MultiSetOf(!(false))
		if (((func() _dafny.MultiSet {
			if _382_cf8 {
				return (_384_v82).Update(_276_v1, Companion_Default___.Abs((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)))
			}
			return (_384_v82).Update((_319_v35).F11(), Companion_Default___.Abs(_275_v0))
		})()).Cardinality()).Cmp((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)) > 0 {
			(_319_v35).M5(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("kdrbwar"), _283_v7)).Cardinality()), _274_globalState)
			(_274_globalState).F1 = (func() _dafny.Int {
				if !(_382_cf8) {
					return _dafny.IntOfInt64(785)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tmolciku")).Cardinality())
			})()
			var _385_v83 _dafny.Sequence
			_ = _385_v83
			_385_v83 = _dafny.SeqOf(_277_v2)
			_277_v2 = (_385_v83).Select((Companion_Default___.SafeIndex((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_385_v83).Cardinality()))).Uint32()).(_dafny.Array)
			var _386_v84 _dafny.Array
			_ = _386_v84
			var _nw65 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(14))
			_ = _nw65
			_386_v84 = _nw65
			var _387_v85 _dafny.Sequence
			_ = _387_v85
			_387_v85 = _dafny.SeqOf(_386_v84, _386_v84, _386_v84)
			_386_v84 = (_387_v85).Select((Companion_Default___.SafeIndex((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_387_v85).Cardinality()))).Uint32()).(_dafny.Array)
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
			_ = _index57
			(_277_v2).ArraySet1((_319_v35).F11(), (_index57).Int())
		} else {
			var _388_v86 *C0
			_ = _388_v86
			var _nw66 *C0 = New_C0_()
			_ = _nw66
			_nw66.Ctor__((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool))
			_388_v86 = _nw66
			var _389_v87 *C0
			_ = _389_v87
			_389_v87 = _388_v86
			_389_v87 = _389_v87
			_382_cf8 = (_319_v35).F11()
			var _390_v88 _dafny.Sequence
			_ = _390_v88
			_390_v88 = _dafny.SeqOf((_382_cf8) || (!(_388_v86.F10)), _382_cf8)
			_390_v88 = _dafny.Companion_Sequence_.Concatenate(_390_v88, _390_v88)
			var _391_v89 D1
			_ = _391_v89
			_391_v89 = Companion_D1_.Create_DC3_((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), (_dafny.Zero).Minus(_275_v0), (_323_v39).Cardinality())
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
			_ = _index58
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))
			_ = _index59
			var _rhs54 bool = (func() bool {
				if _388_v86.F10 {
					return ((_314_v30).Cardinality()).Cmp(_dafny.IntOfInt64(-192)) <= 0
				}
				return (_319_v35).F11()
			})()
			_ = _rhs54
			var _rhs55 bool = (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool)
			_ = _rhs55
			var _rhs56 _dafny.Int = Companion_Default___.SafeDivisionInt((_275_v0).Minus(_dafny.IntOfUint32((_390_v88).Cardinality())), (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
			_ = _rhs56
			var _rhs57 bool = (((_317_v33).Cardinality()).Times(_275_v0)).Cmp(_275_v0) == 0
			_ = _rhs57
			var _rhs58 bool = (_391_v89).Dtor_cf5()
			_ = _rhs58
			var _lhs31 _dafny.Array = _277_v2
			_ = _lhs31
			var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
			_ = _lhs32
			var _lhs33 _dafny.Array = _281_v5
			_ = _lhs33
			var _lhs34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))
			_ = _lhs34
			_276_v1 = _rhs54
			(_lhs31).ArraySet1(_rhs55, (_lhs32).Int())
			(_lhs33).ArraySet1(_rhs56, (_lhs34).Int())
			_382_cf8 = _rhs57
			_382_cf8 = _rhs58
			(_274_globalState).F1 = Companion_Default___.SafeModuloInt(_275_v0, (func() _dafny.Int {
				if (_290_v12).Contains((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)) {
					return (_290_v12).Get((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)).(_dafny.Int)
				}
				return _275_v0
			})())
		}
	} else {
		var _392___mcc_h4 _dafny.Map = _source5.Get_().(D1_DC2).Cf4
		_ = _392___mcc_h4
		var _393_cf4 _dafny.Map = _392___mcc_h4
		_ = _393_cf4
		var _394_v90 _dafny.Sequence
		_ = _394_v90
		_394_v90 = _dafny.SeqOf(_277_v2, _277_v2)
		var _395_v91 _dafny.Array
		_ = _395_v91
		var _nwElement0_11 _dafny.Array = _277_v2
		_ = _nwElement0_11
		var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(12))
		_ = _nw67
		(_nw67).ArraySet1(_nwElement0_11, 0)
		(_nw67).ArraySet1((_394_v90).Select((Companion_Default___.SafeIndex((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_394_v90).Cardinality()))).Uint32()).(_dafny.Array), 1)
		(_nw67).ArraySet1(_277_v2, 2)
		(_nw67).ArraySet1(_277_v2, 3)
		(_nw67).ArraySet1((Companion_D0_.Create_DC1_((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _277_v2, (_319_v35).F11())).Dtor_cf2(), 4)
		(_nw67).ArraySet1(_277_v2, 5)
		(_nw67).ArraySet1(_277_v2, 6)
		(_nw67).ArraySet1(_277_v2, 7)
		(_nw67).ArraySet1(_277_v2, 8)
		(_nw67).ArraySet1(_277_v2, 9)
		(_nw67).ArraySet1(_277_v2, 10)
		(_nw67).ArraySet1(_277_v2, 11)
		_395_v91 = _nw67
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_395_v91), 0))
		_ = _index60
		(_395_v91).ArraySet1((func() _dafny.Array {
			if (_319_v35).F11() {
				return _277_v2
			}
			return _277_v2
		})(), (_index60).Int())
		var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_395_v91), 0))
		_ = _index61
		(_395_v91).ArraySet1(_277_v2, (_index61).Int())
		_290_v12 = (_290_v12).Update((_dafny.Zero).Minus((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)), (_dafny.IntOfUint32((_283_v7).Cardinality())).Plus((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)))
		var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _index62
		(_277_v2).ArraySet1(_276_v1, (_index62).Int())
		_275_v0 = (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)
	}
	if (_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool) {
		_319_v35 = _319_v35
		var _396_v92 _dafny.CodePoint
		_ = _396_v92
		_396_v92 = _dafny.CodePoint('p')
		_396_v92 = (_283_v7).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_275_v0, (_279_v3).Cardinality()), _dafny.IntOfUint32((_283_v7).Cardinality()))).Uint32()).(_dafny.CodePoint)
		_290_v12 = (_290_v12).Update(Companion_Default___.Fm1((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _274_globalState), Companion_Default___.Fm1((_277_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))).Int()).(bool), _274_globalState))
		_275_v0 = Companion_Default___.Fm1(_276_v1, _274_globalState)
		var _397_v94 _dafny.Array
		_ = _397_v94
		var _nwElement0_12 _dafny.Set = _314_v30
		_ = _nwElement0_12
		var _nw68 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(6))
		_ = _nw68
		(_nw68).ArraySet1(_nwElement0_12, 0)
		(_nw68).ArraySet1(_314_v30, 1)
		(_nw68).ArraySet1(_314_v30, 2)
		(_nw68).ArraySet1(_dafny.SetOf(_275_v0, _275_v0, _275_v0, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int)), 3)
		(_nw68).ArraySet1(func() _dafny.Set {
			var _coll14 = _dafny.NewBuilder()
			_ = _coll14
			for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(877), _dafny.IntOfInt64(434))); ; {
				_compr_14, _ok16 := _iter16()
				if !_ok16 {
					break
				}
				var _398_v93 _dafny.Int
				_398_v93 = interface{}(_compr_14).(_dafny.Int)
				if ((_dafny.IntOfInt64(877)).Cmp(_398_v93) <= 0) && ((_398_v93).Cmp(_dafny.IntOfInt64(434)) < 0) {
					_coll14.Add((_398_v93).Times(_275_v0))
				}
			}
			return _coll14.ToSet()
		}(), 4)
		(_nw68).ArraySet1(_dafny.SetOf(_275_v0, _275_v0, _275_v0, _275_v0, _275_v0), 5)
		_397_v94 = _nw68
		var _399_v95 D11
		_ = _399_v95
		_399_v95 = Companion_D11_.Create_DC28_(_283_v7, _397_v94, (_319_v35).F11(), _276_v1)
		var _400_v96 D11
		_ = _400_v96
		_400_v96 = Companion_D11_.Create_DC29_(_399_v95)
		_400_v96 = _400_v96
	} else {
		var _401_v97 _dafny.Set
		_ = _401_v97
		_401_v97 = _dafny.SetOf(_277_v2)
		var _402_v98 _dafny.Map
		_ = _402_v98
		_402_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_276_v1, (_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int))
		var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(673), _dafny.ArrayLen((_277_v2), 0))
		_ = _index63
		(_277_v2).ArraySet1(!(!(Companion_Default___.Fm0((_319_v35).F11(), (_290_v12).Update(((_323_v39).Update(_275_v0, Companion_Default___.Abs((_401_v97).Cardinality()))).Cardinality(), (_402_v98).Cardinality()), _275_v0, _274_globalState))), (_index63).Int())
		var _403_v99 *C0
		_ = _403_v99
		var _nw69 *C0 = New_C0_()
		_ = _nw69
		_nw69.Ctor__((_319_v35).F11())
		_403_v99 = _nw69
		var _404_v100 *C0
		_ = _404_v100
		_404_v100 = _403_v99
		_404_v100 = _404_v100
		(_403_v99).F10 = Companion_Default___.Fm0(Companion_Default___.Fm0((_319_v35).F11(), _290_v12, _275_v0, _274_globalState), _290_v12, (Companion_Default___.Fm40((_281_v5).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_281_v5), 0))).Int()).(_dafny.Int), _403_v99.F10, _274_globalState)).Cardinality(), _274_globalState)
		(_274_globalState).F1 = (_dafny.Zero).Minus(_275_v0)
		(_274_globalState).F1 = _dafny.IntOfInt64(-252)
	}
	_dafny.Print((_274_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_274_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_globalState).F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_274_globalState).F5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_275_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_276_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_277_v2).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_279_v3).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_280_v4).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(false), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_281_v5).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_282_v6).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_283_v7.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_284_v8).Equals(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("klhmhabdk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_285_v9).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_290_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-379), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.IntOfInt64(-4), _dafny.IntOfInt64(567))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_314_v30).Equals(_dafny.SetOf(_dafny.IntOfInt64(-379))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_315_v31).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfInt64(-379)), _dafny.IntOfInt64(-379))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_316_v32, _dafny.SeqOf(_dafny.IntOfInt64(379), _dafny.IntOfInt64(405), _dafny.IntOfInt64(226), _dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_317_v33).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-379), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_318_v34).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_319_v35).F11())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_319_v35.F12).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_320_v37).Dtor_cf30()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('r'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('w'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('h'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('c'), _dafny.IntOfInt64(-379))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_322_v38).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_D7_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('i'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('r'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('e'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('n'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('x'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('w'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('f'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('h'), _dafny.IntOfInt64(-379)).UpdateUnsafe(_dafny.CodePoint('c'), _dafny.IntOfInt64(-379))))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_323_v39).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(4), _dafny.IntOfInt64(-379))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_345_i7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_360_v66).Dtor_cf0()).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_361_v67).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_362_v68).Dtor_cf4()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf(false)), false)))
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
	Cf2 _dafny.Array
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 bool, Cf2 _dafny.Array, Cf3 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Set
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Set) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(false, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false)
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Array {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf0() _dafny.Set {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
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
			return ok && data1.Cf1 == data2.Cf1 && data1.Cf2 == data2.Cf2 && data1.Cf3 == data2.Cf3
		}
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

type D1_DC3 struct {
	Cf5 bool
	Cf6 _dafny.Int
	Cf7 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf5 bool, Cf6 _dafny.Int, Cf7 _dafny.Int) D1 {
	return D1{D1_DC3{Cf5, Cf6, Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC4 struct {
	Cf8 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 bool) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC2 struct {
	Cf4 _dafny.Map
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf4 _dafny.Map) D1 {
	return D1{D1_DC2{Cf4}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D1) Dtor_cf5() bool {
	return _this.Get_().(D1_DC3).Cf5
}

func (_this D1) Dtor_cf6() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf6
}

func (_this D1) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) Dtor_cf8() bool {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf4() _dafny.Map {
	return _this.Get_().(D1_DC2).Cf4
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf4) + ")"
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
			return ok && data1.Cf5 == data2.Cf5 && data1.Cf6.Cmp(data2.Cf6) == 0 && data1.Cf7.Cmp(data2.Cf7) == 0
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8 == data2.Cf8
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
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

type D2_DC6 struct {
	Cf10 D1
	Cf11 _dafny.Sequence
	Cf12 _dafny.Int
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 D1, Cf11 _dafny.Sequence, Cf12 _dafny.Int) D2 {
	return D2{D2_DC6{Cf10, Cf11, Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf13 bool
	Cf14 _dafny.Int
	Cf15 _dafny.Map
	Cf16 D1
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf13 bool, Cf14 _dafny.Int, Cf15 _dafny.Map, Cf16 D1) D2 {
	return D2{D2_DC7{Cf13, Cf14, Cf15, Cf16}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC5 struct {
	Cf9 _dafny.Array
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf9 _dafny.Array) D2 {
	return D2{D2_DC5{Cf9}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC8 struct {
	Cf17 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf17 D2) D2 {
	return D2{D2_DC8{Cf17}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC6_(Companion_D1_.Default(), _dafny.EmptySeq, _dafny.Zero)
}

func (_this D2) Dtor_cf10() D1 {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf11
}

func (_this D2) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf14() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf14
}

func (_this D2) Dtor_cf15() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() D1 {
	return _this.Get_().(D2_DC7).Cf16
}

func (_this D2) Dtor_cf9() _dafny.Array {
	return _this.Get_().(D2_DC5).Cf9
}

func (_this D2) Dtor_cf17() D2 {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ", " + data.Cf11.VerbatimString(true) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf13) + ", " + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf17) + ")"
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
			return ok && data1.Cf10.Equals(data2.Cf10) && data1.Cf11.Equals(data2.Cf11) && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf13 == data2.Cf13 && data1.Cf14.Cmp(data2.Cf14) == 0 && data1.Cf15.Equals(data2.Cf15) && data1.Cf16.Equals(data2.Cf16)
		}
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf9 == data2.Cf9
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
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

type D3_DC9 struct {
	Cf18 _dafny.MultiSet
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf18 _dafny.MultiSet) D3 {
	return D3{D3_DC9{Cf18}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D3) Dtor_cf18() _dafny.MultiSet {
	return _this.Get_().(D3_DC9).Cf18
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf18) + ")"
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

type D4_DC10 struct {
	Cf19 *C0
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf19 *C0) D4 {
	return D4{D4_DC10{Cf19}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() *C0 {
	return (*C0)(nil)
}

func (_this D4) Dtor_cf19() *C0 {
	return _this.Get_().(D4_DC10).Cf19
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf19) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf19 == data2.Cf19
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
	Cf21 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf21 _dafny.Int) D5 {
	return D5{D5_DC12{Cf21}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC11 struct {
	Cf20 _dafny.Sequence
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf20 _dafny.Sequence) D5 {
	return D5{D5_DC11{Cf20}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(_dafny.Zero)
}

func (_this D5) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf21
}

func (_this D5) Dtor_cf20() _dafny.Sequence {
	return _this.Get_().(D5_DC11).Cf20
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf21) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf20) + ")"
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
			return ok && data1.Cf21.Cmp(data2.Cf21) == 0
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf20.Equals(data2.Cf20)
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
	Cf23 _dafny.Int
	Cf24 _dafny.MultiSet
	Cf25 _dafny.Int
}

func (D6_DC14) isD6() {}

func (CompanionStruct_D6_) Create_DC14_(Cf23 _dafny.Int, Cf24 _dafny.MultiSet, Cf25 _dafny.Int) D6 {
	return D6{D6_DC14{Cf23, Cf24, Cf25}}
}

func (_this D6) Is_DC14() bool {
	_, ok := _this.Get_().(D6_DC14)
	return ok
}

type D6_DC15 struct {
	Cf26 bool
	Cf27 _dafny.Int
	Cf28 _dafny.Int
	Cf29 _dafny.Int
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf26 bool, Cf27 _dafny.Int, Cf28 _dafny.Int, Cf29 _dafny.Int) D6 {
	return D6{D6_DC15{Cf26, Cf27, Cf28, Cf29}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

type D6_DC13 struct {
	Cf22 _dafny.Array
}

func (D6_DC13) isD6() {}

func (CompanionStruct_D6_) Create_DC13_(Cf22 _dafny.Array) D6 {
	return D6{D6_DC13{Cf22}}
}

func (_this D6) Is_DC13() bool {
	_, ok := _this.Get_().(D6_DC13)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC14_(_dafny.Zero, _dafny.EmptyMultiSet, _dafny.Zero)
}

func (_this D6) Dtor_cf23() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf23
}

func (_this D6) Dtor_cf24() _dafny.MultiSet {
	return _this.Get_().(D6_DC14).Cf24
}

func (_this D6) Dtor_cf25() _dafny.Int {
	return _this.Get_().(D6_DC14).Cf25
}

func (_this D6) Dtor_cf26() bool {
	return _this.Get_().(D6_DC15).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf27
}

func (_this D6) Dtor_cf28() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf28
}

func (_this D6) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D6_DC15).Cf29
}

func (_this D6) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D6_DC13).Cf22
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC14:
		{
			return "D6.DC14" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ", " + _dafny.String(data.Cf25) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
		}
	case D6_DC13:
		{
			return "D6.DC13" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Cmp(data2.Cf23) == 0 && data1.Cf24.Equals(data2.Cf24) && data1.Cf25.Cmp(data2.Cf25) == 0
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27.Cmp(data2.Cf27) == 0 && data1.Cf28.Cmp(data2.Cf28) == 0 && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D6_DC13:
		{
			data2, ok := other.Get_().(D6_DC13)
			return ok && data1.Cf22 == data2.Cf22
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

type D7_DC17 struct {
	Cf31 _dafny.Sequence
	Cf32 _dafny.Set
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf31 _dafny.Sequence, Cf32 _dafny.Set) D7 {
	return D7{D7_DC17{Cf31, Cf32}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

type D7_DC16 struct {
	Cf30 _dafny.Map
}

func (D7_DC16) isD7() {}

func (CompanionStruct_D7_) Create_DC16_(Cf30 _dafny.Map) D7 {
	return D7{D7_DC16{Cf30}}
}

func (_this D7) Is_DC16() bool {
	_, ok := _this.Get_().(D7_DC16)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC17_(_dafny.EmptySeq, _dafny.EmptySet)
}

func (_this D7) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D7_DC17).Cf31
}

func (_this D7) Dtor_cf32() _dafny.Set {
	return _this.Get_().(D7_DC17).Cf32
}

func (_this D7) Dtor_cf30() _dafny.Map {
	return _this.Get_().(D7_DC16).Cf30
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC17:
		{
			return "D7.DC17" + "(" + data.Cf31.VerbatimString(true) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC16:
		{
			return "D7.DC16" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf31.Equals(data2.Cf31) && data1.Cf32.Equals(data2.Cf32)
		}
	case D7_DC16:
		{
			data2, ok := other.Get_().(D7_DC16)
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
	Cf34 bool
	Cf35 bool
	Cf36 _dafny.Set
	Cf37 _dafny.Int
	Cf38 _dafny.Int
}

func (D8_DC19) isD8() {}

func (CompanionStruct_D8_) Create_DC19_(Cf34 bool, Cf35 bool, Cf36 _dafny.Set, Cf37 _dafny.Int, Cf38 _dafny.Int) D8 {
	return D8{D8_DC19{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D8) Is_DC19() bool {
	_, ok := _this.Get_().(D8_DC19)
	return ok
}

type D8_DC20 struct {
	Cf39 bool
	Cf40 bool
	Cf41 bool
	Cf42 bool
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf39 bool, Cf40 bool, Cf41 bool, Cf42 bool) D8 {
	return D8{D8_DC20{Cf39, Cf40, Cf41, Cf42}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC21 struct {
	Cf43 D2
	Cf44 _dafny.Map
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf43 D2, Cf44 _dafny.Map) D8 {
	return D8{D8_DC21{Cf43, Cf44}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC18 struct {
	Cf33 T0
}

func (D8_DC18) isD8() {}

func (CompanionStruct_D8_) Create_DC18_(Cf33 T0) D8 {
	return D8{D8_DC18{Cf33}}
}

func (_this D8) Is_DC18() bool {
	_, ok := _this.Get_().(D8_DC18)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC19_(false, false, _dafny.EmptySet, _dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf34() bool {
	return _this.Get_().(D8_DC19).Cf34
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC19).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Set {
	return _this.Get_().(D8_DC19).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC19).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D8_DC19).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC20).Cf39
}

func (_this D8) Dtor_cf40() bool {
	return _this.Get_().(D8_DC20).Cf40
}

func (_this D8) Dtor_cf41() bool {
	return _this.Get_().(D8_DC20).Cf41
}

func (_this D8) Dtor_cf42() bool {
	return _this.Get_().(D8_DC20).Cf42
}

func (_this D8) Dtor_cf43() D2 {
	return _this.Get_().(D8_DC21).Cf43
}

func (_this D8) Dtor_cf44() _dafny.Map {
	return _this.Get_().(D8_DC21).Cf44
}

func (_this D8) Dtor_cf33() T0 {
	return _this.Get_().(D8_DC18).Cf33
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC19:
		{
			return "D8.DC19" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ", " + _dafny.String(data.Cf42) + ")"
		}
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
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
	case D8_DC19:
		{
			data2, ok := other.Get_().(D8_DC19)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40 && data1.Cf41 == data2.Cf41 && data1.Cf42 == data2.Cf42
		}
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf43.Equals(data2.Cf43) && data1.Cf44.Equals(data2.Cf44)
		}
	case D8_DC18:
		{
			data2, ok := other.Get_().(D8_DC18)
			return ok && _dafny.AreEqual(data1.Cf33, data2.Cf33)
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

type D9_DC23 struct {
	Cf46 bool
	Cf47 _dafny.Int
	Cf48 bool
	Cf49 bool
	Cf50 _dafny.Int
}

func (D9_DC23) isD9() {}

func (CompanionStruct_D9_) Create_DC23_(Cf46 bool, Cf47 _dafny.Int, Cf48 bool, Cf49 bool, Cf50 _dafny.Int) D9 {
	return D9{D9_DC23{Cf46, Cf47, Cf48, Cf49, Cf50}}
}

func (_this D9) Is_DC23() bool {
	_, ok := _this.Get_().(D9_DC23)
	return ok
}

type D9_DC22 struct {
	Cf45 _dafny.Sequence
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf45 _dafny.Sequence) D9 {
	return D9{D9_DC22{Cf45}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC23_(false, _dafny.Zero, false, false, _dafny.Zero)
}

func (_this D9) Dtor_cf46() bool {
	return _this.Get_().(D9_DC23).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf47
}

func (_this D9) Dtor_cf48() bool {
	return _this.Get_().(D9_DC23).Cf48
}

func (_this D9) Dtor_cf49() bool {
	return _this.Get_().(D9_DC23).Cf49
}

func (_this D9) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D9_DC23).Cf50
}

func (_this D9) Dtor_cf45() _dafny.Sequence {
	return _this.Get_().(D9_DC22).Cf45
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC23:
		{
			return "D9.DC23" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ", " + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC23:
		{
			data2, ok := other.Get_().(D9_DC23)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48 == data2.Cf48 && data1.Cf49 == data2.Cf49 && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
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

type D10_DC25 struct {
}

func (D10_DC25) isD10() {}

func (CompanionStruct_D10_) Create_DC25_() D10 {
	return D10{D10_DC25{}}
}

func (_this D10) Is_DC25() bool {
	_, ok := _this.Get_().(D10_DC25)
	return ok
}

type D10_DC26 struct {
	Cf52 bool
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf52 bool) D10 {
	return D10{D10_DC26{Cf52}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

type D10_DC24 struct {
	Cf51 _dafny.CodePoint
}

func (D10_DC24) isD10() {}

func (CompanionStruct_D10_) Create_DC24_(Cf51 _dafny.CodePoint) D10 {
	return D10{D10_DC24{Cf51}}
}

func (_this D10) Is_DC24() bool {
	_, ok := _this.Get_().(D10_DC24)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC25_()
}

func (_this D10) Dtor_cf52() bool {
	return _this.Get_().(D10_DC26).Cf52
}

func (_this D10) Dtor_cf51() _dafny.CodePoint {
	return _this.Get_().(D10_DC24).Cf51
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC25:
		{
			return "D10.DC25"
		}
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf52) + ")"
		}
	case D10_DC24:
		{
			return "D10.DC24" + "(" + _dafny.String(data.Cf51) + ")"
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
			_, ok := other.Get_().(D10_DC25)
			return ok
		}
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf52 == data2.Cf52
		}
	case D10_DC24:
		{
			data2, ok := other.Get_().(D10_DC24)
			return ok && data1.Cf51 == data2.Cf51
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

type D11_DC28 struct {
	Cf54 _dafny.Sequence
	Cf55 _dafny.Array
	Cf56 bool
	Cf57 bool
}

func (D11_DC28) isD11() {}

func (CompanionStruct_D11_) Create_DC28_(Cf54 _dafny.Sequence, Cf55 _dafny.Array, Cf56 bool, Cf57 bool) D11 {
	return D11{D11_DC28{Cf54, Cf55, Cf56, Cf57}}
}

func (_this D11) Is_DC28() bool {
	_, ok := _this.Get_().(D11_DC28)
	return ok
}

type D11_DC27 struct {
	Cf53 _dafny.Set
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf53 _dafny.Set) D11 {
	return D11{D11_DC27{Cf53}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC29 struct {
	Cf58 D11
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf58 D11) D11 {
	return D11{D11_DC29{Cf58}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC28_(_dafny.EmptySeq, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, false)
}

func (_this D11) Dtor_cf54() _dafny.Sequence {
	return _this.Get_().(D11_DC28).Cf54
}

func (_this D11) Dtor_cf55() _dafny.Array {
	return _this.Get_().(D11_DC28).Cf55
}

func (_this D11) Dtor_cf56() bool {
	return _this.Get_().(D11_DC28).Cf56
}

func (_this D11) Dtor_cf57() bool {
	return _this.Get_().(D11_DC28).Cf57
}

func (_this D11) Dtor_cf53() _dafny.Set {
	return _this.Get_().(D11_DC27).Cf53
}

func (_this D11) Dtor_cf58() D11 {
	return _this.Get_().(D11_DC29).Cf58
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC28:
		{
			return "D11.DC28" + "(" + data.Cf54.VerbatimString(true) + ", " + _dafny.String(data.Cf55) + ", " + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf53) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf58) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC28:
		{
			data2, ok := other.Get_().(D11_DC28)
			return ok && data1.Cf54.Equals(data2.Cf54) && data1.Cf55 == data2.Cf55 && data1.Cf56 == data2.Cf56 && data1.Cf57 == data2.Cf57
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf53.Equals(data2.Cf53)
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf58.Equals(data2.Cf58)
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

type D12_DC31 struct {
	Cf60 _dafny.Int
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf60 _dafny.Int) D12 {
	return D12{D12_DC31{Cf60}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

type D12_DC30 struct {
	Cf59 _dafny.Sequence
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf59 _dafny.Sequence) D12 {
	return D12{D12_DC30{Cf59}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC31_(_dafny.Zero)
}

func (_this D12) Dtor_cf60() _dafny.Int {
	return _this.Get_().(D12_DC31).Cf60
}

func (_this D12) Dtor_cf59() _dafny.Sequence {
	return _this.Get_().(D12_DC30).Cf59
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf60) + ")"
		}
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf60.Cmp(data2.Cf60) == 0
		}
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf59.Equals(data2.Cf59)
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

type D13_DC33 struct {
	Cf62 bool
	Cf63 _dafny.Int
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf62 bool, Cf63 _dafny.Int) D13 {
	return D13{D13_DC33{Cf62, Cf63}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

type D13_DC32 struct {
	Cf61 _dafny.Map
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf61 _dafny.Map) D13 {
	return D13{D13_DC32{Cf61}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC33_(false, _dafny.Zero)
}

func (_this D13) Dtor_cf62() bool {
	return _this.Get_().(D13_DC33).Cf62
}

func (_this D13) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D13_DC33).Cf63
}

func (_this D13) Dtor_cf61() _dafny.Map {
	return _this.Get_().(D13_DC32).Cf61
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ")"
		}
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf61) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf62 == data2.Cf62 && data1.Cf63.Cmp(data2.Cf63) == 0
		}
	case D13_DC32:
		{
			data2, ok := other.Get_().(D13_DC32)
			return ok && data1.Cf61.Equals(data2.Cf61)
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

// Definition of trait T0
type T0 interface {
	String() string
	F7() _dafny.Int
	F7_set_(value _dafny.Int)
	M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int)
	F6() _dafny.Int
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
	_f0 bool
	_f2 _dafny.Int
	_f3 _dafny.Int
	_f4 _dafny.Int
	_f5 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this._f0 = false
	_this._f2 = _dafny.Zero
	_this._f3 = _dafny.Zero
	_this._f4 = _dafny.Zero
	_this._f5 = false
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 _dafny.Int, f3 _dafny.Int, f4 _dafny.Int, f5 bool) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
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
func (_this *GlobalState) F4() _dafny.Int {
	{
		return _this._f4
	}
}
func (_this *GlobalState) F5() bool {
	{
		return _this._f5
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	F10 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this.F10 = false
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

func (_this *C0) Ctor__(f10 bool) {
	{
		(_this).F10 = f10
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f7  _dafny.Int
	_f6  _dafny.Int
	_f13 _dafny.CodePoint
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f7 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f13 = _dafny.CodePoint('D')
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

func (_this *C1) F7() _dafny.Int {
	return _this._f7
}
func (_this *C1) F7_set_(value _dafny.Int) {
	_this._f7 = value
}
func (_this *C1) F6() _dafny.Int {
	return _this._f6
}
func (_this *C1) Ctor__(f13 _dafny.CodePoint, f6 _dafny.Int, f7 _dafny.Int) {
	{
		(_this)._f13 = f13
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C1) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _405_v0 _dafny.Map
		_ = _405_v0
		_405_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
		var _406_v1 _dafny.Sequence
		_ = _406_v1
		_406_v1 = _dafny.UnicodeSeqOfUtf8Bytes("sm")
		var _407_v2 _dafny.Sequence
		_ = _407_v2
		_407_v2 = _dafny.SeqOf(Companion_Default___.Fm0(!(p0), ((_405_v0).Update(p1, _this.F7())).Update(_dafny.IntOfUint32((_406_v1).Cardinality()), (_this).F6()), (_this).F6(), globalState), p0, p0, true, p0)
		var _408_v3 _dafny.Array
		_ = _408_v3
		var _nwElement0_13 bool = p0
		_ = _nwElement0_13
		var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(9))
		_ = _nw70
		(_nw70).ArraySet1(_nwElement0_13, 0)
		(_nw70).ArraySet1(p0, 1)
		(_nw70).ArraySet1(p0, 2)
		(_nw70).ArraySet1(true, 3)
		(_nw70).ArraySet1(p0, 4)
		(_nw70).ArraySet1(p0, 5)
		(_nw70).ArraySet1(p0, 6)
		(_nw70).ArraySet1(p0, 7)
		(_nw70).ArraySet1(p0, 8)
		_408_v3 = _nw70
		var _409_v4 _dafny.Map
		_ = _409_v4
		_409_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v3, _407_v2)
		var _410_v5 _dafny.Set
		_ = _410_v5
		_410_v5 = _dafny.SetOf(_407_v2, (func() _dafny.Sequence {
			if (_409_v4).Contains(_408_v3) {
				return (_409_v4).Get(_408_v3).(_dafny.Sequence)
			}
			return _407_v2
		})(), _407_v2, _407_v2, _407_v2)
		var _rhs59 _dafny.Set = _410_v5
		_ = _rhs59
		var _rhs60 bool = p0
		_ = _rhs60
		var _rhs61 _dafny.Int = ((_this).F6()).Plus(p2)
		_ = _rhs61
		var _lhs35 *C1 = _this
		_ = _lhs35
		_410_v5 = _rhs59
		r1 = _rhs60
		_lhs35.F7_set_(_rhs61)
		var _411_v6 *C0
		_ = _411_v6
		var _nw71 *C0 = New_C0_()
		_ = _nw71
		_nw71.Ctor__(p0)
		_411_v6 = _nw71
		var _412_v7 _dafny.Map
		_ = _412_v7
		_412_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_406_v1, _411_v6)
		_412_v7 = (_412_v7).Update(_406_v1, _411_v6)
		var _413_v8 _dafny.Array
		_ = _413_v8
		var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
		_ = _nw72
		_413_v8 = _nw72
		var _414_v9 D2
		_ = _414_v9
		_414_v9 = Companion_D2_.Create_DC5_(_413_v8)
		var _source6 D2 = (func() D2 {
			if _411_v6.F10 {
				return _414_v9
			}
			return _414_v9
		})()
		_ = _source6
		if _source6.Is_DC6() {
			var _415___mcc_h0 D1 = _source6.Get_().(D2_DC6).Cf10
			_ = _415___mcc_h0
			var _416___mcc_h1 _dafny.Sequence = _source6.Get_().(D2_DC6).Cf11
			_ = _416___mcc_h1
			var _417___mcc_h2 _dafny.Int = _source6.Get_().(D2_DC6).Cf12
			_ = _417___mcc_h2
			var _418_cf12 _dafny.Int = _417___mcc_h2
			_ = _418_cf12
			var _419_cf11 _dafny.Sequence = _416___mcc_h1
			_ = _419_cf11
			var _420_cf10 D1 = _415___mcc_h0
			_ = _420_cf10
			_419_cf11 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("g"), _dafny.UnicodeSeqOfUtf8Bytes("yo"))
			(globalState).F1 = _this.F7()
			var _421_v10 D1
			_ = _421_v10
			_421_v10 = Companion_D1_.Create_DC4_(_411_v6.F10)
			(_411_v6).F10 = (_421_v10).Dtor_cf8()
			var _422_v11 _dafny.Sequence
			_ = _422_v11
			_422_v11 = _dafny.SeqOf(_411_v6)
			var _423_v12 _dafny.Map
			_ = _423_v12
			_423_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), _411_v6)
			var _424_v13 _dafny.Array
			_ = _424_v13
			var _nwElement0_14 *C0 = _411_v6
			_ = _nwElement0_14
			var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(11))
			_ = _nw73
			(_nw73).ArraySet1(_nwElement0_14, 0)
			(_nw73).ArraySet1(_411_v6, 1)
			(_nw73).ArraySet1(_411_v6, 2)
			(_nw73).ArraySet1(_411_v6, 3)
			(_nw73).ArraySet1(_411_v6, 4)
			(_nw73).ArraySet1(_411_v6, 5)
			(_nw73).ArraySet1(_411_v6, 6)
			(_nw73).ArraySet1(_411_v6, 7)
			(_nw73).ArraySet1((_422_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(582), _dafny.IntOfUint32((_422_v11).Cardinality()))).Uint32()).(*C0), 8)
			(_nw73).ArraySet1((func() *C0 {
				if (_423_v12).Contains(_this.F7()) {
					return (_423_v12).Get(_this.F7()).(*C0)
				}
				return _411_v6
			})(), 9)
			(_nw73).ArraySet1(_411_v6, 10)
			_424_v13 = _nw73
			var _rhs62 bool = _411_v6.F10
			_ = _rhs62
			var _rhs63 bool = _411_v6.F10
			_ = _rhs63
			var _rhs64 bool = _411_v6.F10
			_ = _rhs64
			var _rhs65 _dafny.Array = _424_v13
			_ = _rhs65
			var _rhs66 bool = p0
			_ = _rhs66
			var _lhs36 *C0 = _411_v6
			_ = _lhs36
			r1 = _rhs62
			_lhs36.F10 = _rhs63
			r0 = _rhs64
			_424_v13 = _rhs65
			r0 = _rhs66
		} else if _source6.Is_DC7() {
			var _425___mcc_h3 bool = _source6.Get_().(D2_DC7).Cf13
			_ = _425___mcc_h3
			var _426___mcc_h4 _dafny.Int = _source6.Get_().(D2_DC7).Cf14
			_ = _426___mcc_h4
			var _427___mcc_h5 _dafny.Map = _source6.Get_().(D2_DC7).Cf15
			_ = _427___mcc_h5
			var _428___mcc_h6 D1 = _source6.Get_().(D2_DC7).Cf16
			_ = _428___mcc_h6
			var _429_cf16 D1 = _428___mcc_h6
			_ = _429_cf16
			var _430_cf15 _dafny.Map = _427___mcc_h5
			_ = _430_cf15
			var _431_cf14 _dafny.Int = _426___mcc_h4
			_ = _431_cf14
			var _432_cf13 bool = _425___mcc_h3
			_ = _432_cf13
			var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_413_v8), 0))
			_ = _index64
			(_413_v8).ArraySet1(_dafny.IntOfInt64(341), (_index64).Int())
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_413_v8), 0))
			_ = _index65
			(_413_v8).ArraySet1(_dafny.IntOfInt64(-810), (_index65).Int())
			r2 = _431_cf14
			var _433_v14 _dafny.CodePoint
			_ = _433_v14
			_433_v14 = _dafny.CodePoint('t')
			var _434_v15 _dafny.MultiSet
			_ = _434_v15
			_434_v15 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Contains(_407_v2, _411_v6.F10), _411_v6.F10)
			var _435_v16 _dafny.Sequence
			_ = _435_v16
			_435_v16 = _dafny.SeqOf((_dafny.IntOfUint32((_407_v2).Cardinality())).Times(_this.F7()))
			var _436_v17 _dafny.Array
			_ = _436_v17
			var _nw74 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw74
			_436_v17 = _nw74
			var _437_v18 _dafny.Map
			_ = _437_v18
			_437_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v6.F10, _436_v17)
			var _438_v19 _dafny.Set
			_ = _438_v19
			_438_v19 = _dafny.SetOf(p0)
			var _439_v20 _dafny.Map
			_ = _439_v20
			_439_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_438_v19), _411_v6.F10)
			var _440_v21 D1
			_ = _440_v21
			_440_v21 = Companion_D1_.Create_DC2_(_439_v20)
			var _441_v22 _dafny.MultiSet
			_ = _441_v22
			_441_v22 = _dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(135), p2, (_this).F6(), (_this).F6())
			var _442_v23 D2
			_ = _442_v23
			_442_v23 = Companion_D2_.Create_DC6_(_440_v21, _406_v1, (func() _dafny.Int {
				if (_441_v22).Contains(_dafny.IntOfInt64(595)) {
					return (_441_v22).Multiplicity(_dafny.IntOfInt64(595))
				}
				return _this.F7()
			})())
			var _443_v24 _dafny.Map
			_ = _443_v24
			_443_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_437_v18).Contains(_411_v6.F10) {
					return (_437_v18).Get(_411_v6.F10).(_dafny.Array)
				}
				return _436_v17
			})(), _442_v23)
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_413_v8), 0))
			_ = _index66
			var _rhs67 _dafny.CodePoint = Companion_Default___.Fm25(globalState)
			_ = _rhs67
			var _rhs68 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_435_v16).Cardinality())))
			_ = _rhs68
			var _rhs69 _dafny.MultiSet = _dafny.MultiSetOf(_432_cf13, !(_443_v24).Equals(_443_v24))
			_ = _rhs69
			var _rhs70 _dafny.Int = _this.F7()
			_ = _rhs70
			var _lhs37 *GlobalState = globalState
			_ = _lhs37
			var _lhs38 _dafny.Array = _413_v8
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(845), _dafny.ArrayLen((_413_v8), 0))
			_ = _lhs39
			_433_v14 = _rhs67
			_lhs37.F1 = _rhs68
			_434_v15 = _rhs69
			(_lhs38).ArraySet1(_rhs70, (_lhs39).Int())
			var _444_v25 _dafny.Set
			_ = _444_v25
			_444_v25 = _dafny.SetOf((_this).F6(), p1, (_dafny.SetOf((_this).F6(), (_this).F6())).Cardinality())
			var _445_v26 _dafny.MultiSet
			_ = _445_v26
			_445_v26 = _dafny.MultiSetOf(_444_v25, _444_v25, _444_v25, _444_v25, _444_v25)
			r2 = ((_445_v26).Difference((_445_v26).Intersection(_445_v26))).Cardinality()
		} else if _source6.Is_DC5() {
			var _446___mcc_h7 _dafny.Array = _source6.Get_().(D2_DC5).Cf9
			_ = _446___mcc_h7
			var _447_cf9 _dafny.Array = _446___mcc_h7
			_ = _447_cf9
			var _448_v27 _dafny.CodePoint
			_ = _448_v27
			_448_v27 = _dafny.CodePoint('d')
			_448_v27 = _448_v27
			r1 = _411_v6.F10
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_408_v3), 0))
			_ = _index67
			(_408_v3).ArraySet1(true, (_index67).Int())
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_408_v3), 0))
			_ = _index68
			(_408_v3).ArraySet1(!(_411_v6.F10), (_index68).Int())
			r0 = (_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(963), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool)
		} else {
			var _449___mcc_h8 D2 = _source6.Get_().(D2_DC8).Cf17
			_ = _449___mcc_h8
			var _450_cf17 D2 = _449___mcc_h8
			_ = _450_cf17
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_408_v3), 0))
			_ = _index69
			(_408_v3).ArraySet1(p0, (_index69).Int())
			var _451_v28 _dafny.Sequence
			_ = _451_v28
			_451_v28 = _dafny.SeqOf(_406_v1, _406_v1, _406_v1)
			var _452_v29 D0
			_ = _452_v29
			_452_v29 = Companion_D0_.Create_DC0_(_dafny.SetOf(p0, p0))
			var _453_v30 _dafny.Map
			_ = _453_v30
			_453_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_452_v29, _411_v6.F10)
			var _454_v31 D1
			_ = _454_v31
			_454_v31 = Companion_D1_.Create_DC2_(_453_v30)
			var _455_v32 D2
			_ = _455_v32
			_455_v32 = Companion_D2_.Create_DC6_(_454_v31, _dafny.UnicodeSeqOfUtf8Bytes("ukxgv"), (_this).F6())
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_408_v3), 0))
			_ = _index70
			var _rhs71 _dafny.Sequence = (_451_v28).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_406_v1, _406_v1)).Cardinality()), _dafny.IntOfUint32((_451_v28).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs71
			var _rhs72 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_455_v32).Dtor_cf11(), _406_v1), _dafny.Companion_Sequence_.Concatenate(_406_v1, _406_v1))
			_ = _rhs72
			var _rhs73 bool = Companion_Default___.Fm0(p0, _405_v0, _this.F7(), globalState)
			_ = _rhs73
			var _lhs40 _dafny.Array = _408_v3
			_ = _lhs40
			var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_408_v3), 0))
			_ = _lhs41
			_406_v1 = _rhs71
			_406_v1 = _rhs72
			(_lhs40).ArraySet1(_rhs73, (_lhs41).Int())
			var _456_v33 _dafny.Map
			_ = _456_v33
			_456_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v6.F10, p2)
			var _457_v34 _dafny.MultiSet
			_ = _457_v34
			_457_v34 = _dafny.MultiSetOf(_413_v8)
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_408_v3), 0))
			_ = _index71
			var _rhs74 _dafny.Int = ((_dafny.Zero).Minus((func() _dafny.Int {
				if (_456_v33).Contains(_411_v6.F10) {
					return (_456_v33).Get(_411_v6.F10).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(p2)
			})())).Minus(((_457_v34).Update(_413_v8, Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32((_406_v1).Cardinality()))))).Cardinality())
			_ = _rhs74
			var _rhs75 bool = ((p2).Cmp((_this).F6()) < 0) || (_411_v6.F10)
			_ = _rhs75
			var _lhs42 *C1 = _this
			_ = _lhs42
			var _lhs43 _dafny.Array = _408_v3
			_ = _lhs43
			var _lhs44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(364), _dafny.ArrayLen((_408_v3), 0))
			_ = _lhs44
			_lhs42.F7_set_(_rhs74)
			(_lhs43).ArraySet1(_rhs75, (_lhs44).Int())
			var _458_v35 _dafny.Set
			_ = _458_v35
			_458_v35 = _dafny.SetOf(p2)
			var _459_v37 _dafny.Sequence
			_ = _459_v37
			_459_v37 = _dafny.SeqOf((_458_v35).Intersection(_458_v35), _458_v35, func() _dafny.Set {
				var _coll15 = _dafny.NewBuilder()
				_ = _coll15
				for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(118), _dafny.IntOfInt64(-912))); ; {
					_compr_15, _ok17 := _iter17()
					if !_ok17 {
						break
					}
					var _460_v36 _dafny.Int
					_460_v36 = interface{}(_compr_15).(_dafny.Int)
					if ((_dafny.IntOfInt64(118)).Cmp(_460_v36) <= 0) && ((_460_v36).Cmp(_dafny.IntOfInt64(-912)) < 0) {
						_coll15.Add((_460_v36).Plus(p1))
					}
				}
				return _coll15.ToSet()
			}())
			_459_v37 = Companion_Default___.Fm26((_dafny.Zero).Minus((_455_v32).Dtor_cf12()), p0, globalState)
			var _461_v39 *C0
			_ = _461_v39
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__((func() _dafny.Set {
				var _coll16 = _dafny.NewBuilder()
				_ = _coll16
				for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-209), _dafny.IntOfInt64(576))); ; {
					_compr_16, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _462_v38 _dafny.Int
					_462_v38 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(-209)).Cmp(_462_v38) <= 0) && ((_462_v38).Cmp(_dafny.IntOfInt64(576)) < 0) {
						_coll16.Add((_462_v38).Plus(_dafny.IntOfUint32((_407_v2).Cardinality())))
					}
				}
				return _coll16.ToSet()
			}()).IsProperSubsetOf(_458_v35))
			_461_v39 = _nw75
		}
		(_this).F7_set_(_dafny.IntOfUint32((_406_v1).Cardinality()))
		if false {
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_413_v8), 0))
			_ = _index72
			(_413_v8).ArraySet1(p2, (_index72).Int())
			var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_413_v8), 0))
			_ = _index73
			(_413_v8).ArraySet1(p2, (_index73).Int())
			(globalState).F1 = _this.F7()
			var _463_v40 _dafny.Array
			_ = _463_v40
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_9
			var _nw76 _dafny.Array
			_ = _nw76
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw76 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) _dafny.Sequence = (func(_464_v6 *C0, _465_v2 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_466_i0 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_464_v6.F10, _464_v6.F10), _465_v2)
					}
				})(_411_v6, _407_v2)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw76 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw76).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw76).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_463_v40 = _nw76
			var _467_v41 D0
			_ = _467_v41
			_467_v41 = Companion_D0_.Create_DC1_(_411_v6.F10, _408_v3, _411_v6.F10)
			var _468_v42 _dafny.Map
			_ = _468_v42
			_468_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v6.F10, !(_411_v6.F10))
			var _469_v43 D5
			_ = _469_v43
			_469_v43 = Companion_D5_.Create_DC11_(_407_v2)
			var _470_v44 _dafny.Sequence
			_ = _470_v44
			_470_v44 = _dafny.SeqOf(_407_v2, _407_v2)
			var _471_v45 _dafny.Set
			_ = _471_v45
			_471_v45 = _dafny.SetOf(_this.F7())
			var _472_v46 _dafny.Set
			_ = _472_v46
			_472_v46 = _dafny.SetOf((_471_v45).Cardinality())
			var _pat_let_tv6 = _411_v6
			_ = _pat_let_tv6
			var _nwElement0_15 _dafny.Sequence = _407_v2
			_ = _nwElement0_15
			var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(28))
			_ = _nw77
			(_nw77).ArraySet1(_nwElement0_15, 0)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_407_v2, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_406_v1, (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_406_v1).Cardinality()))).Uint32(), (_this).F13())).Cardinality()), _dafny.IntOfUint32((_407_v2).Cardinality()))).Uint32(), _411_v6.F10), _407_v2), 1)
			(_nw77).ArraySet1(_407_v2, 2)
			(_nw77).ArraySet1(_407_v2, 3)
			(_nw77).ArraySet1(_dafny.SeqOf(p0, (_467_v41).Dtor_cf3()), 4)
			(_nw77).ArraySet1(_dafny.SeqOf((func() bool {
				if (_468_v42).Contains(_411_v6.F10) {
					return (_468_v42).Get(_411_v6.F10).(bool)
				}
				return _411_v6.F10
			})()), 5)
			(_nw77).ArraySet1(_407_v2, 6)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_411_v6.F10, _411_v6.F10), _dafny.Companion_Sequence_.Update((Companion_D5_.Create_DC11_(_407_v2)).Dtor_cf20(), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((Companion_D5_.Create_DC11_(_407_v2)).Dtor_cf20()).Cardinality()))).Uint32(), _411_v6.F10)), 7)
			(_nw77).ArraySet1(_dafny.SeqOf(_411_v6.F10, _411_v6.F10), 8)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_407_v2, (Companion_D5_.Create_DC11_(_dafny.SeqOf(false))).Dtor_cf20()), 9)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_407_v2, _407_v2), 10)
			(_nw77).ArraySet1((_469_v43).Dtor_cf20(), 11)
			(_nw77).ArraySet1(_407_v2, 12)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false, p0, p0)), 13)
			(_nw77).ArraySet1((_470_v44).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_470_v44).Cardinality()))).Uint32()).(_dafny.Sequence), 14)
			(_nw77).ArraySet1(_407_v2, 15)
			(_nw77).ArraySet1(_dafny.SeqOf((_407_v2).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_407_v2).Cardinality()))).Uint32()).(bool), _411_v6.F10), 16)
			(_nw77).ArraySet1(_407_v2, 17)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_407_v2, _407_v2), 18)
			(_nw77).ArraySet1(_407_v2, 19)
			(_nw77).ArraySet1(_407_v2, 20)
			(_nw77).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_407_v2, _407_v2), 21)
			(_nw77).ArraySet1(_407_v2, 22)
			(_nw77).ArraySet1(_407_v2, 23)
			(_nw77).ArraySet1(_dafny.SeqOf(Companion_Default___.Fm0(_411_v6.F10, _405_v0, p1, globalState)), 24)
			(_nw77).ArraySet1(_407_v2, 25)
			(_nw77).ArraySet1(_dafny.SeqOf((func(_pat_let9_0 D0) D0 {
				return func(_473_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let10_0 bool) D0 {
						return func(_474_dt__update_hcf1_h0 bool) D0 {
							return Companion_D0_.Create_DC1_(_474_dt__update_hcf1_h0, (_473_dt__update__tmp_h0).Dtor_cf2(), (_473_dt__update__tmp_h0).Dtor_cf3())
						}(_pat_let10_0)
					}(_pat_let_tv6.F10)
				}(_pat_let9_0)
			}(_467_v41)).Dtor_cf1()), 26)
			(_nw77).ArraySet1(Companion_Default___.Fm27((_471_v45).Cardinality(), _406_v1, (_this).F13(), globalState), 27)
			_463_v40 = _nw77
			r0 = !((func() bool {
				if _411_v6.F10 {
					return _411_v6.F10
				}
				return ((_472_v46).Cardinality()).Cmp((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)) < 0
			})())
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_413_v8), 0))
			_ = _index74
			(_413_v8).ArraySet1(_this.F7(), (_index74).Int())
		} else {
			var _475_v47 _dafny.Map
			_ = _475_v47
			_475_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_413_v8, (_this).F13())
			_475_v47 = _475_v47
			var _476_v48 _dafny.Map
			_ = _476_v48
			_476_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _411_v6.F10)
			var _477_v49 _dafny.Sequence
			_ = _477_v49
			_477_v49 = _dafny.SeqOf(_476_v48)
			var _478_v50 D1
			_ = _478_v50
			_478_v50 = Companion_D1_.Create_DC3_(false, ((_477_v49).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_406_v1).Cardinality()), _dafny.IntOfUint32((_477_v49).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality(), p1)
			r0 = (func() bool {
				if (_478_v50).Dtor_cf5() {
					return _411_v6.F10
				}
				return true
			})()
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
			_ = _index75
			(_408_v3).ArraySet1((_411_v6.F10) || (_411_v6.F10), (_index75).Int())
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))
			_ = _index76
			(_413_v8).ArraySet1(p1, (_index76).Int())
			var _479_v51 _dafny.Set
			_ = _479_v51
			_479_v51 = _dafny.SetOf(_411_v6.F10, _411_v6.F10)
			var _480_v52 _dafny.Map
			_ = _480_v52
			_480_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_479_v51).Cardinality(), _479_v51)
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
			_ = _index77
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))
			_ = _index78
			var _rhs76 bool = _411_v6.F10
			_ = _rhs76
			var _rhs77 bool = _411_v6.F10
			_ = _rhs77
			var _rhs78 _dafny.Int = (_dafny.Zero).Minus(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _479_v51)).Merge((_480_v52).Merge(_480_v52))).Cardinality())
			_ = _rhs78
			var _rhs79 _dafny.Int = p1
			_ = _rhs79
			var _lhs45 *C0 = _411_v6
			_ = _lhs45
			var _lhs46 _dafny.Array = _408_v3
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
			_ = _lhs47
			var _lhs48 *GlobalState = globalState
			_ = _lhs48
			var _lhs49 _dafny.Array = _413_v8
			_ = _lhs49
			var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))
			_ = _lhs50
			_lhs45.F10 = _rhs76
			(_lhs46).ArraySet1(_rhs77, (_lhs47).Int())
			_lhs48.F1 = _rhs78
			(_lhs49).ArraySet1(_rhs79, (_lhs50).Int())
			var _source7 D1 = _478_v50
			_ = _source7
			if _source7.Is_DC3() {
				var _481___mcc_h9 bool = _source7.Get_().(D1_DC3).Cf5
				_ = _481___mcc_h9
				var _482___mcc_h10 _dafny.Int = _source7.Get_().(D1_DC3).Cf6
				_ = _482___mcc_h10
				var _483___mcc_h11 _dafny.Int = _source7.Get_().(D1_DC3).Cf7
				_ = _483___mcc_h11
				var _484_cf7 _dafny.Int = _483___mcc_h11
				_ = _484_cf7
				var _485_cf6 _dafny.Int = _482___mcc_h10
				_ = _485_cf6
				var _486_cf5 bool = _481___mcc_h9
				_ = _486_cf5
				var _487_v53 _dafny.CodePoint
				_ = _487_v53
				_487_v53 = _dafny.CodePoint('u')
				var _488_v54 _dafny.Set
				_ = _488_v54
				_488_v54 = _dafny.SetOf(_this.F7())
				_487_v53 = (func() _dafny.CodePoint {
					if (p1).Cmp((_488_v54).Cardinality()) > 0 {
						return (_this).F13()
					}
					return _487_v53
				})()
				var _489_v55 D5
				_ = _489_v55
				_489_v55 = Companion_D5_.Create_DC12_(_485_cf6)
				var _490_v56 _dafny.MultiSet
				_ = _490_v56
				_490_v56 = _dafny.MultiSetOf((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int))
				var _491_v57 _dafny.Map
				_ = _491_v57
				_491_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _490_v56)
				_489_v55 = Companion_Default___.Fm28((func() _dafny.Sequence {
					if false {
						return Companion_Default___.Fm27((_491_v57).Cardinality(), _406_v1, (_this).F13(), globalState)
					}
					return _407_v2
				})(), p2, globalState)
				_478_v50 = _478_v50
				var _492_v59 _dafny.Map
				_ = _492_v59
				_492_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_485_cf6, _487_v53)
				var _493_v60 _dafny.Set
				_ = _493_v60
				_493_v60 = _dafny.SetOf(_492_v59, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(532), (_this).F13()))
				_485_cf6 = ((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)).Minus((_dafny.Zero).Minus((func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter19 := _dafny.Iterate((_493_v60).Elements()); ; {
						_compr_17, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _494_v58 _dafny.Map
						_494_v58 = interface{}(_compr_17).(_dafny.Map)
						if (_493_v60).Contains(_494_v58) {
							_coll17.Add(_494_v58, _485_cf6)
						}
					}
					return _coll17.ToMap()
				}()).Cardinality()))
			} else if _source7.Is_DC4() {
				var _495___mcc_h12 bool = _source7.Get_().(D1_DC4).Cf8
				_ = _495___mcc_h12
				var _496_cf8 bool = _495___mcc_h12
				_ = _496_cf8
				var _497_v61 _dafny.Sequence
				_ = _497_v61
				_497_v61 = _dafny.SeqOf((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int), (_this).F6())
				var _498_v62 _dafny.Array
				_ = _498_v62
				var _nwElement0_16 _dafny.Sequence = _497_v61
				_ = _nwElement0_16
				var _nw78 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(5))
				_ = _nw78
				(_nw78).ArraySet1(_nwElement0_16, 0)
				(_nw78).ArraySet1(_497_v61, 1)
				(_nw78).ArraySet1(_497_v61, 2)
				(_nw78).ArraySet1(_497_v61, 3)
				(_nw78).ArraySet1(_497_v61, 4)
				_498_v62 = _nw78
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_498_v62), 0))
				_ = _index79
				(_498_v62).ArraySet1(_497_v61, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
				_ = _index80
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_498_v62), 0))
				_ = _index81
				var _rhs80 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(403))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg27 _dafny.Int) interface{} {
						return coer27(arg27)
					}
				}(func(_499_i1 _dafny.Int) _dafny.CodePoint {
					return (_this).F13()
				})), _dafny.UnicodeSeqOfUtf8Bytes("endmc")), _dafny.UnicodeSeqOfUtf8Bytes("gmwqy"))
				_ = _rhs80
				var _rhs81 bool = _411_v6.F10
				_ = _rhs81
				var _rhs82 _dafny.Sequence = _497_v61
				_ = _rhs82
				var _lhs51 _dafny.Array = _408_v3
				_ = _lhs51
				var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
				_ = _lhs52
				var _lhs53 _dafny.Array = _498_v62
				_ = _lhs53
				var _lhs54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(328), _dafny.ArrayLen((_498_v62), 0))
				_ = _lhs54
				_406_v1 = _rhs80
				(_lhs51).ArraySet1(_rhs81, (_lhs52).Int())
				(_lhs53).ArraySet1(_rhs82, (_lhs54).Int())
				var _500_v63 _dafny.Array
				_ = _500_v63
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw79
				_500_v63 = _nw79
				var _501_v64 _dafny.Sequence
				_ = _501_v64
				_501_v64 = _dafny.SeqOf(_500_v63)
				var _502_v65 _dafny.Map
				_ = _502_v65
				_502_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_406_v1, false)
				_500_v63 = (_501_v64).Select((Companion_Default___.SafeIndex(((_502_v65).Cardinality()).Times(p1), _dafny.IntOfUint32((_501_v64).Cardinality()))).Uint32()).(_dafny.Array)
				var _503_v66 _dafny.Map
				_ = _503_v66
				_503_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _496_cf8)
				var _504_v67 _dafny.Sequence
				_ = _504_v67
				_504_v67 = _dafny.SeqOf(_503_v66)
				_503_v66 = ((_504_v67).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_504_v67).Cardinality()))).Uint32()).(_dafny.Map)).Update(!(_496_cf8), (_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool))
				var _505_v68 _dafny.MultiSet
				_ = _505_v68
				_505_v68 = _dafny.MultiSetOf(_dafny.IntOfInt64(168), (_this).F6(), p1)
				var _506_v69 _dafny.Map
				_ = _506_v69
				_506_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_cf8, _407_v2)
				var _507_v70 _dafny.MultiSet
				_ = _507_v70
				_507_v70 = _dafny.MultiSetOf(false, _411_v6.F10)
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
				_ = _index82
				var _rhs83 bool = p0
				_ = _rhs83
				var _rhs84 bool = ((func() _dafny.Int {
					if (_405_v0).Contains((_506_v69).Cardinality()) {
						return (_405_v0).Get((_506_v69).Cardinality()).(_dafny.Int)
					}
					return Companion_Default___.Fm1(_411_v6.F10, globalState)
				})()).Cmp((_dafny.Zero).Minus(_this.F7())) <= 0
				_ = _rhs84
				var _rhs85 _dafny.MultiSet = _dafny.MultiSetOf(((_dafny.MultiSetFromSeq(_dafny.SeqOf(_411_v6.F10, _411_v6.F10))).Union(_507_v70)).Cardinality())
				_ = _rhs85
				var _rhs86 bool = !(_411_v6.F10)
				_ = _rhs86
				var _lhs55 _dafny.Array = _408_v3
				_ = _lhs55
				var _lhs56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
				_ = _lhs56
				var _lhs57 *C0 = _411_v6
				_ = _lhs57
				(_lhs55).ArraySet1(_rhs83, (_lhs56).Int())
				_lhs57.F10 = _rhs84
				_505_v68 = _rhs85
				r0 = _rhs86
			} else {
				var _508___mcc_h13 _dafny.Map = _source7.Get_().(D1_DC2).Cf4
				_ = _508___mcc_h13
				var _509_cf4 _dafny.Map = _508___mcc_h13
				_ = _509_cf4
				var _510_v71 *C0
				_ = _510_v71
				var _nw80 *C0 = New_C0_()
				_ = _nw80
				_nw80.Ctor__(!(_411_v6.F10))
				_510_v71 = _nw80
				var _511_v72 _dafny.Array
				_ = _511_v72
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_10
				var _nw81 _dafny.Array
				_ = _nw81
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw81 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = (func(_512_v71 *C0, _513_v6 *C0) func(_dafny.Int) _dafny.Int {
						return func(_514_i2 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_514_i2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_512_v71.F10, _dafny.SeqOf(_513_v6.F10))).Cardinality())
						}
					})(_510_v71, _411_v6)
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw81 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw81).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw81).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_511_v72 = _nw81
				_511_v72 = _511_v72
				var _515_v73 _dafny.Array
				_ = _515_v73
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_11
				var _nw82 _dafny.Array
				_ = _nw82
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw82 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Map = (func(_516_v0 _dafny.Map) func(_dafny.Int) _dafny.Map {
						return func(_517_i3 _dafny.Int) _dafny.Map {
							return _516_v0
						}
					})(_405_v0)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw82 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw82).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw82).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_515_v73 = _nw82
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_515_v73), 0))
				_ = _index83
				(_515_v73).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm1(Companion_Default___.Fm0(_411_v6.F10, _405_v0, p2, globalState), globalState)), (_index83).Int())
				var _518_v74 _dafny.Array
				_ = _518_v74
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_12
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) _dafny.Sequence = (func(_519_p1 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
						return func(_520_i4 _dafny.Int) _dafny.Sequence {
							return _dafny.SeqOf(_dafny.IntOfInt64(959), _519_p1)
						}
					})(p1)
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw83 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw83).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw83).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_518_v74 = _nw83
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_515_v73), 0))
				_ = _index84
				var _rhs87 _dafny.Int = _this.F7()
				_ = _rhs87
				var _rhs88 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(-575))), ((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)).Plus((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)))
				_ = _rhs88
				var _rhs89 _dafny.Map = (_405_v0).Merge(_405_v0)
				_ = _rhs89
				var _rhs90 bool = Companion_Default___.Fm0(!(true) || (p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-267), _this.F7()), (_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int), globalState)
				_ = _rhs90
				var _rhs91 _dafny.Array = _518_v74
				_ = _rhs91
				var _lhs58 *GlobalState = globalState
				_ = _lhs58
				var _lhs59 _dafny.Array = _515_v73
				_ = _lhs59
				var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(939), _dafny.ArrayLen((_515_v73), 0))
				_ = _lhs60
				var _lhs61 *C0 = _411_v6
				_ = _lhs61
				_lhs58.F1 = _rhs87
				r2 = _rhs88
				(_lhs59).ArraySet1(_rhs89, (_lhs60).Int())
				_lhs61.F10 = _rhs90
				_518_v74 = _rhs91
				(_510_v71).F10 = false
			}
			if p0 {
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))
				_ = _index85
				(_408_v3).ArraySet1(((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)).Cmp((p2).Plus((_dafny.Zero).Minus((_this).F6()))) == 0, (_index85).Int())
				(globalState).F1 = _dafny.IntOfInt64(846)
				var _521_v75 _dafny.Map
				_ = _521_v75
				_521_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_479_v51), !(_411_v6.F10))
				var _522_v76 D1
				_ = _522_v76
				_522_v76 = Companion_D1_.Create_DC2_(_521_v75)
				var _523_v78 D0
				_ = _523_v78
				_523_v78 = Companion_D0_.Create_DC0_(_dafny.SetOf((_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool)))
				var _pat_let_tv7 = _479_v51
				_ = _pat_let_tv7
				var _pat_let_tv8 = _479_v51
				_ = _pat_let_tv8
				var _pat_let_tv9 = _521_v75
				_ = _pat_let_tv9
				var _524_v79 _dafny.Array
				_ = _524_v79
				var _nwElement0_17 D1 = _522_v76
				_ = _nwElement0_17
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(23))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_17, 0)
				(_nw84).ArraySet1(_522_v76, 1)
				(_nw84).ArraySet1(_522_v76, 2)
				(_nw84).ArraySet1(_522_v76, 3)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(func() _dafny.Map {
					var _coll18 = _dafny.NewMapBuilder()
					_ = _coll18
					for _iter20 := _dafny.Iterate((_dafny.SetOf(_523_v78, func(_pat_let11_0 D0) D0 {
						return func(_525_dt__update__tmp_h1 D0) D0 {
							return func(_pat_let12_0 _dafny.Set) D0 {
								return func(_526_dt__update_hcf0_h0 _dafny.Set) D0 {
									return Companion_D0_.Create_DC0_(_526_dt__update_hcf0_h0)
								}(_pat_let12_0)
							}(_pat_let_tv7)
						}(_pat_let11_0)
					}(Companion_D0_.Create_DC0_(_479_v51)))).Elements()); ; {
						_compr_18, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _527_v77 D0
						_527_v77 = interface{}(_compr_18).(D0)
						if (_dafny.SetOf(_523_v78, func(_pat_let13_0 D0) D0 {
							return func(_528_dt__update__tmp_h1 D0) D0 {
								return func(_pat_let14_0 _dafny.Set) D0 {
									return func(_529_dt__update_hcf0_h0 _dafny.Set) D0 {
										return Companion_D0_.Create_DC0_(_529_dt__update_hcf0_h0)
									}(_pat_let14_0)
								}(_pat_let_tv8)
							}(_pat_let13_0)
						}(Companion_D0_.Create_DC0_(_479_v51)))).Contains(_527_v77) {
							_coll18.Add(_527_v77, true)
						}
					}
					return _coll18.ToMap()
				}()), 4)
				(_nw84).ArraySet1(_522_v76, 5)
				(_nw84).ArraySet1(_522_v76, 6)
				(_nw84).ArraySet1(_522_v76, 7)
				(_nw84).ArraySet1(_522_v76, 8)
				(_nw84).ArraySet1(_522_v76, 9)
				(_nw84).ArraySet1(func(_pat_let15_0 D1) D1 {
					return func(_530_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let16_0 _dafny.Map) D1 {
							return func(_531_dt__update_hcf4_h0 _dafny.Map) D1 {
								return Companion_D1_.Create_DC2_(_531_dt__update_hcf4_h0)
							}(_pat_let16_0)
						}(_pat_let_tv9)
					}(_pat_let15_0)
				}(Companion_D1_.Create_DC2_(_521_v75)), 10)
				(_nw84).ArraySet1(_522_v76, 11)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(_521_v75), 12)
				(_nw84).ArraySet1(_522_v76, 13)
				(_nw84).ArraySet1(_522_v76, 14)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(_521_v75), 15)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(_521_v75), 16)
				(_nw84).ArraySet1(_522_v76, 17)
				(_nw84).ArraySet1(_522_v76, 18)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(_521_v75), 19)
				(_nw84).ArraySet1(_522_v76, 20)
				(_nw84).ArraySet1(_522_v76, 21)
				(_nw84).ArraySet1(Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_v78, false)), 22)
				_524_v79 = _nw84
				var _532_v80 D6
				_ = _532_v80
				_532_v80 = Companion_D6_.Create_DC13_(_524_v79)
				var _533_v81 _dafny.Array
				_ = _533_v81
				var _nwElement0_18 _dafny.Array = _524_v79
				_ = _nwElement0_18
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(23))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_18, 0)
				(_nw85).ArraySet1(_524_v79, 1)
				(_nw85).ArraySet1(_524_v79, 2)
				(_nw85).ArraySet1(_524_v79, 3)
				(_nw85).ArraySet1(_524_v79, 4)
				(_nw85).ArraySet1((_532_v80).Dtor_cf22(), 5)
				(_nw85).ArraySet1(_524_v79, 6)
				(_nw85).ArraySet1(_524_v79, 7)
				(_nw85).ArraySet1(_524_v79, 8)
				(_nw85).ArraySet1(_524_v79, 9)
				(_nw85).ArraySet1(_524_v79, 10)
				(_nw85).ArraySet1(_524_v79, 11)
				(_nw85).ArraySet1(_524_v79, 12)
				(_nw85).ArraySet1(_524_v79, 13)
				(_nw85).ArraySet1(_524_v79, 14)
				(_nw85).ArraySet1(_524_v79, 15)
				(_nw85).ArraySet1(_524_v79, 16)
				(_nw85).ArraySet1(_524_v79, 17)
				(_nw85).ArraySet1(_524_v79, 18)
				(_nw85).ArraySet1(_524_v79, 19)
				(_nw85).ArraySet1(_524_v79, 20)
				(_nw85).ArraySet1(_524_v79, 21)
				(_nw85).ArraySet1(_524_v79, 22)
				_533_v81 = _nw85
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_533_v81), 0))
				_ = _index86
				(_533_v81).ArraySet1(_524_v79, (_index86).Int())
				var _534_v82 _dafny.MultiSet
				_ = _534_v82
				_534_v82 = _dafny.MultiSetOf(true, _411_v6.F10, _411_v6.F10, !(p0), (_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool))
				var _535_v83 _dafny.Map
				_ = _535_v83
				_535_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
					if Companion_Default___.Fm0(p0, _405_v0, _this.F7(), globalState) {
						return _411_v6.F10
					}
					return (_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool)
				})(), _524_v79)
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_533_v81), 0))
				_ = _index87
				var _rhs92 _dafny.Int = Companion_Default___.SafeModuloInt((_this.F7()).Plus((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)), (_534_v82).Cardinality())
				_ = _rhs92
				var _rhs93 _dafny.Array = (func() _dafny.Array {
					if (_535_v83).Contains(_411_v6.F10) {
						return (_535_v83).Get(_411_v6.F10).(_dafny.Array)
					}
					return _524_v79
				})()
				_ = _rhs93
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				var _lhs63 _dafny.Array = _533_v81
				_ = _lhs63
				var _lhs64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(940), _dafny.ArrayLen((_533_v81), 0))
				_ = _lhs64
				_lhs62.F1 = _rhs92
				(_lhs63).ArraySet1(_rhs93, (_lhs64).Int())
				var _536_v84 *C0
				_ = _536_v84
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__(p0)
				_536_v84 = _nw86
				var _537_v85 _dafny.Array
				_ = _537_v85
				var _nw87 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(9))
				_ = _nw87
				_537_v85 = _nw87
				var _538_v86 _dafny.Sequence
				_ = _538_v86
				_538_v86 = _dafny.SeqOf(_537_v85, (func() _dafny.Array {
					if p0 {
						return _537_v85
					}
					return _537_v85
				})())
				_537_v85 = (_538_v86).Select((Companion_Default___.SafeIndex((_this.F7()).Minus(_this.F7()), _dafny.IntOfUint32((_538_v86).Cardinality()))).Uint32()).(_dafny.Array)
			} else {
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))
				_ = _index88
				(_413_v8).ArraySet1((_dafny.Zero).Minus((_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int)), (_index88).Int())
				_406_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_406_v1, _406_v1), _dafny.Companion_Sequence_.Concatenate(_406_v1, _dafny.UnicodeSeqOfUtf8Bytes("lvmvk")))
				(_411_v6).F10 = _411_v6.F10
				var _539_v88 _dafny.Set
				_ = _539_v88
				_539_v88 = _dafny.SetOf(_406_v1)
				_479_v51 = (Companion_Default___.Fm3(p0, p1, _479_v51, func() _dafny.Map {
					var _coll19 = _dafny.NewMapBuilder()
					_ = _coll19
					for _iter21 := _dafny.Iterate((_539_v88).Elements()); ; {
						_compr_19, _ok21 := _iter21()
						if !_ok21 {
							break
						}
						var _540_v87 _dafny.Sequence
						_540_v87 = interface{}(_compr_19).(_dafny.Sequence)
						if (_539_v88).Contains(_540_v87) {
							_coll19.Add(_540_v87, (_413_v8).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(100), _dafny.ArrayLen((_413_v8), 0))).Int()).(_dafny.Int))
						}
					}
					return _coll19.ToMap()
				}(), globalState)).Dtor_cf0()
				(_411_v6).F10 = !(!((_408_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(342), _dafny.ArrayLen((_408_v3), 0))).Int()).(bool)) || (_411_v6.F10))
			}
		}
		r1 = Companion_Default___.Fm0(p0, (_405_v0).Merge(_405_v0), (_this).F6(), globalState)
		r0 = true
		r1 = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v3, _this.F7())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_408_v3, p2))
		r2 = Companion_Default___.Fm1(_411_v6.F10, globalState)
		return r0, r1, r2
	}
}
func (_this *C1) M10(p0 _dafny.Int, globalState *GlobalState) (_dafny.MultiSet, _dafny.Int, D1, bool) {
	{
		var r0 _dafny.MultiSet = _dafny.EmptyMultiSet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 D1 = Companion_D1_.Default()
		_ = r2
		var r3 bool = false
		_ = r3
		var _541_v0 bool
		_ = _541_v0
		_541_v0 = false
		var _542_v1 _dafny.Map
		_ = _542_v1
		_542_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _dafny.IntOfInt64(654))
		var _hi4 _dafny.Int = ((_542_v1).Merge(_542_v1)).Cardinality()
		_ = _hi4
		for _543_i0 := (Companion_Default___.Fm1(_541_v0, globalState)).Times(_dafny.IntOfInt64(-878)); _543_i0.Cmp(_hi4) < 0; _543_i0 = _543_i0.Plus(_dafny.One) {
			var _544_v2 _dafny.Set
			_ = _544_v2
			_544_v2 = _dafny.SetOf(_543_i0)
			var _545_v3 _dafny.Set
			_ = _545_v3
			_545_v3 = _dafny.SetOf(_dafny.IntOfInt64(612), (_544_v2).Cardinality(), (_this).F6())
			var _546_v4 _dafny.Sequence
			_ = _546_v4
			_546_v4 = _dafny.UnicodeSeqOfUtf8Bytes("twoo")
			var _547_v5 _dafny.Sequence
			_ = _547_v5
			_547_v5 = _dafny.SeqOf(_541_v0, _541_v0)
			var _548_v6 _dafny.Set
			_ = _548_v6
			_548_v6 = _dafny.SetOf(_541_v0)
			var _549_v7 _dafny.Map
			_ = _549_v7
			_549_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _546_v4)
			var _550_v8 _dafny.Array
			_ = _550_v8
			var _nwElement0_19 _dafny.Int = (_545_v3).Cardinality()
			_ = _nwElement0_19
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(17))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_19, 0)
			(_nw88).ArraySet1(p0, 1)
			(_nw88).ArraySet1((_this).F6(), 2)
			(_nw88).ArraySet1(p0, 3)
			(_nw88).ArraySet1((_this).F6(), 4)
			(_nw88).ArraySet1(_543_i0, 5)
			(_nw88).ArraySet1(_dafny.IntOfUint32((_546_v4).Cardinality()), 6)
			(_nw88).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetFromSeq(_547_v5)).Cardinality()), 7)
			(_nw88).ArraySet1((_548_v6).Cardinality(), 8)
			(_nw88).ArraySet1(_543_i0, 9)
			(_nw88).ArraySet1(_543_i0, 10)
			(_nw88).ArraySet1(p0, 11)
			(_nw88).ArraySet1(_dafny.IntOfInt64(502), 12)
			(_nw88).ArraySet1(_543_i0, 13)
			(_nw88).ArraySet1(_dafny.IntOfUint32((_547_v5).Cardinality()), 14)
			(_nw88).ArraySet1(_543_i0, 15)
			(_nw88).ArraySet1(_dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_549_v7).Contains(_541_v0) {
					return (_549_v7).Get(_541_v0).(_dafny.Sequence)
				}
				return _546_v4
			})()).Cardinality()), 16)
			_550_v8 = _nw88
			var _551_v9 _dafny.Sequence
			_ = _551_v9
			_551_v9 = _dafny.SeqOf(_543_i0, (_dafny.Zero).Minus(_543_i0))
			var _552_v10 _dafny.Map
			_ = _552_v10
			_552_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_551_v9, _543_i0)
			var _553_v11 _dafny.Map
			_ = _553_v11
			_553_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v8, (_552_v10).Cardinality())
			var _554_v12 _dafny.Map
			_ = _554_v12
			_554_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_i0, _543_i0)
			_553_v11 = (_553_v11).Update((func() _dafny.Array {
				if Companion_Default___.Fm0(!(_541_v0), _554_v12, _dafny.IntOfInt64(432), globalState) {
					return _550_v8
				}
				return _550_v8
			})(), (_this).F6())
			r3 = _541_v0
			_548_v6 = _548_v6
			var _555_v13 _dafny.Array
			_ = _555_v13
			var _nw89 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw89
			_555_v13 = _nw89
			var _556_v14 _dafny.Map
			_ = _556_v14
			_556_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), p0)
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_555_v13), 0))
			_ = _index89
			(_555_v13).ArraySet1(_556_v14, (_index89).Int())
			var _557_v15 D7
			_ = _557_v15
			_557_v15 = Companion_D7_.Create_DC16_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), p0))
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(34), _dafny.ArrayLen((_555_v13), 0))
			_ = _index90
			(_555_v13).ArraySet1((func() _dafny.Map {
				if _541_v0 {
					return _556_v14
				}
				return (_557_v15).Dtor_cf30()
			})(), (_index90).Int())
		}
		var _558_v16 _dafny.Array
		_ = _558_v16
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(12)
		_ = _len0_13
		var _nw90 _dafny.Array
		_ = _nw90
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw90 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.Int = func(_559_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeDivisionInt(_559_i1, (_this).F6())
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw90 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw90).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw90).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_558_v16 = _nw90
		var _560_v17 _dafny.Map
		_ = _560_v17
		_560_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F7())
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _index91
		(_558_v16).ArraySet1((func() _dafny.Int {
			if (_560_v17).Contains(_this.F7()) {
				return (_560_v17).Get(_this.F7()).(_dafny.Int)
			}
			return (_this).F6()
		})(), (_index91).Int())
		var _561_v18 _dafny.MultiSet
		_ = _561_v18
		_561_v18 = _dafny.MultiSetOf(_541_v0)
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _index92
		var _rhs94 _dafny.Int = p0
		_ = _rhs94
		var _rhs95 bool = (_561_v18).IsSubsetOf(_561_v18)
		_ = _rhs95
		var _lhs65 _dafny.Array = _558_v16
		_ = _lhs65
		var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _lhs66
		(_lhs65).ArraySet1(_rhs94, (_lhs66).Int())
		r3 = _rhs95
		var _562_v19 _dafny.Sequence
		_ = _562_v19
		_562_v19 = _dafny.SeqOf(((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int)).Minus((_this).F6()))
		var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _index93
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _index94
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _index95
		var _rhs96 _dafny.Int = (_this).F6()
		_ = _rhs96
		var _rhs97 _dafny.Int = ((func() _dafny.Int {
			if _541_v0 {
				return (_this).F6()
			}
			return (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int)
		})()).Minus((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_this.F7(), (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int))))
		_ = _rhs97
		var _rhs98 _dafny.Int = (_562_v19).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_this).F6(), p0), _dafny.IntOfUint32((_562_v19).Cardinality()))).Uint32()).(_dafny.Int)
		_ = _rhs98
		var _lhs67 _dafny.Array = _558_v16
		_ = _lhs67
		var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _lhs68
		var _lhs69 _dafny.Array = _558_v16
		_ = _lhs69
		var _lhs70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _lhs70
		var _lhs71 _dafny.Array = _558_v16
		_ = _lhs71
		var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
		_ = _lhs72
		(_lhs67).ArraySet1(_rhs96, (_lhs68).Int())
		(_lhs69).ArraySet1(_rhs97, (_lhs70).Int())
		(_lhs71).ArraySet1(_rhs98, (_lhs72).Int())
		var _hi5 _dafny.Int = p0
		_ = _hi5
		for _563_i2 := Companion_Default___.Fm1(_541_v0, globalState); _563_i2.Cmp(_hi5) < 0; _563_i2 = _563_i2.Plus(_dafny.One) {
			var _564_v20 _dafny.Set
			_ = _564_v20
			_564_v20 = _dafny.SetOf((_563_i2).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_563_i2))) != 0, (_541_v0) == (_541_v0), _541_v0, Companion_Default___.Fm0(_541_v0, _560_v17, (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), globalState), false)
			_564_v20 = _564_v20
			var _565_v21 _dafny.Map
			_ = _565_v21
			_565_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_562_v19, _this.F7())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _index96
			var _rhs99 _dafny.Map = (Companion_Default___.Fm29(globalState)).Merge((_565_v21).Merge(_565_v21))
			_ = _rhs99
			var _rhs100 _dafny.Int = (_dafny.Zero).Minus((p0).Minus((p0).Plus(_dafny.IntOfInt64(410))))
			_ = _rhs100
			var _rhs101 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_562_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-878))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg28 _dafny.Int) interface{} {
					return coer28(arg28)
				}
			}(func(_566_i3 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-628))
			})), _562_v19)), (Companion_Default___.SafeIndex((Companion_D1_.Create_DC3_(_541_v0, (_542_v1).Cardinality(), (_this).F6())).Dtor_cf6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_562_v19, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-878))).Uint32(), func(coer29 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}(func(_567_i3 _dafny.Int) _dafny.Int {
				return (_dafny.Zero).Minus(_dafny.IntOfInt64(-628))
			})), _562_v19))).Cardinality()))).Uint32(), (_this).F6())
			_ = _rhs101
			var _lhs73 _dafny.Array = _558_v16
			_ = _lhs73
			var _lhs74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _lhs74
			_565_v21 = _rhs99
			(_lhs73).ArraySet1(_rhs100, (_lhs74).Int())
			_562_v19 = _rhs101
			var _568_v22 _dafny.Array
			_ = _568_v22
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
			_ = _nw91
			_568_v22 = _nw91
			_568_v22 = _568_v22
			var _569_v23 _dafny.Array
			_ = _569_v23
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_14
			var _nw92 _dafny.Array
			_ = _nw92
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw92 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Map = (func(_570_v17 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_571_i4 _dafny.Int) _dafny.Map {
						return _570_v17
					}
				})(_560_v17)
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw92 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw92).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw92).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_569_v23 = _nw92
			var _572_v24 _dafny.Map
			_ = _572_v24
			_572_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_541_v0, globalState), _569_v23)
			_572_v24 = (_572_v24).Update(Companion_Default___.SafeModuloInt((_this).F6(), (_dafny.Zero).Minus(_563_i2)), _569_v23)
		}
		var _hi6 _dafny.Int = _this.F7()
		_ = _hi6
		for _573_i5 := Companion_Default___.SafeModuloInt(p0, (_560_v17).Cardinality()); _573_i5.Cmp(_hi6) < 0; _573_i5 = _573_i5.Plus(_dafny.One) {
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _index97
			(_558_v16).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-363), _573_i5), (_index97).Int())
			(globalState).F1 = p0
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _index98
			(_558_v16).ArraySet1(_this.F7(), (_index98).Int())
			var _574_v25 _dafny.Array
			_ = _574_v25
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
			_ = _nw93
			_574_v25 = _nw93
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_574_v25), 0))
			_ = _index99
			(_574_v25).ArraySet1CodePoint((_this).F13(), (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(627), _dafny.ArrayLen((_574_v25), 0))
			_ = _index100
			(_574_v25).ArraySet1CodePoint((_this).F13(), (_index100).Int())
		}
		var _575_v26 _dafny.Set
		_ = _575_v26
		_575_v26 = _dafny.SetOf(_this.F7(), (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int))
		if (_dafny.SetOf(_dafny.IntOfInt64(-627))).IsDisjointFrom(_dafny.SetOf(_dafny.IntOfUint32((_562_v19).Cardinality()), (_575_v26).Cardinality(), _dafny.IntOfInt64(208))) {
			var _576_v27 *C0
			_ = _576_v27
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__((func() bool {
				if true {
					return _541_v0
				}
				return _541_v0
			})())
			_576_v27 = _nw94
			_541_v0 = _541_v0
			var _577_v29 _dafny.Sequence
			_ = _577_v29
			_577_v29 = _dafny.SeqOf(_541_v0, _541_v0, _576_v27.F10)
			var _578_v30 _dafny.Sequence
			_ = _578_v30
			_578_v30 = _dafny.SeqOf(_577_v29, _577_v29)
			(_this).F7_set_((func() _dafny.Int {
				if !(_576_v27.F10) {
					return (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if _576_v27.F10 {
						return (func() _dafny.Map {
							var _coll20 = _dafny.NewMapBuilder()
							_ = _coll20
							for _iter22 := _dafny.Iterate((_578_v30).Elements()); ; {
								_compr_20, _ok22 := _iter22()
								if !_ok22 {
									break
								}
								var _579_v28 _dafny.Sequence
								_579_v28 = interface{}(_compr_20).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_578_v30, _579_v28) {
									_coll20.Add(_579_v28, p0)
								}
							}
							return _coll20.ToMap()
						}()).Cardinality()
					}
					return (_this).F6()
				})()
			})())
			(_576_v27).F10 = _541_v0
			var _580_v31 _dafny.Sequence
			_ = _580_v31
			_580_v31 = _dafny.UnicodeSeqOfUtf8Bytes("cykvhpkly")
			var _581_v32 _dafny.MultiSet
			_ = _581_v32
			_581_v32 = _dafny.MultiSetOf((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int))
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _index101
			var _rhs102 _dafny.Int = (func() _dafny.Int {
				if _576_v27.F10 {
					return (_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int)
				}
				return _dafny.IntOfInt64(200)
			})()
			_ = _rhs102
			var _rhs103 _dafny.Array = _558_v16
			_ = _rhs103
			var _rhs104 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.IntOfUint32((_580_v31).Cardinality())).Plus((_dafny.SetOf((_581_v32).Cardinality(), _dafny.IntOfInt64(-162), _dafny.IntOfInt64(-131), _dafny.IntOfInt64(355))).Cardinality()), ((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int)).Minus(_this.F7()))
			_ = _rhs104
			var _rhs105 _dafny.Int = (_this.F7()).Minus(_this.F7())
			_ = _rhs105
			var _rhs106 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_577_v29, _dafny.Companion_Sequence_.Concatenate(_577_v29, _577_v29))
			_ = _rhs106
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			var _lhs76 _dafny.Array = _558_v16
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			_lhs75.F1 = _rhs102
			_558_v16 = _rhs103
			(_lhs76).ArraySet1(_rhs104, (_lhs77).Int())
			_lhs78.F1 = _rhs105
			_577_v29 = _rhs106
		} else {
			if _541_v0 {
				var _582_v33 _dafny.CodePoint
				_ = _582_v33
				_582_v33 = _dafny.CodePoint('i')
				_582_v33 = _582_v33
				var _583_v34 *C0
				_ = _583_v34
				var _nw95 *C0 = New_C0_()
				_ = _nw95
				_nw95.Ctor__(_541_v0)
				_583_v34 = _nw95
				var _584_v35 _dafny.Set
				_ = _584_v35
				_584_v35 = _dafny.SetOf(_558_v16, _558_v16, _558_v16, _558_v16)
				var _rhs107 _dafny.Set = _584_v35
				_ = _rhs107
				var _rhs108 _dafny.Int = Companion_Default___.Fm1((_dafny.MultiSetFromSeq(_562_v19)).IsProperSubsetOf(_dafny.MultiSetFromSeq(_562_v19)), globalState)
				_ = _rhs108
				var _lhs79 *GlobalState = globalState
				_ = _lhs79
				_584_v35 = _rhs107
				_lhs79.F1 = _rhs108
				(_583_v34).F10 = _541_v0
				var _585_v36 _dafny.Map
				_ = _585_v36
				_585_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), _558_v16)
				_585_v36 = (_585_v36).Update((_this).F6(), _558_v16)
			} else {
				_541_v0 = !((_541_v0) == ((false) || (_541_v0)))
				_541_v0 = (func() bool {
					if _541_v0 {
						return _541_v0
					}
					return _541_v0
				})()
				var _586_v37 D2
				_ = _586_v37
				_586_v37 = Companion_D2_.Create_DC5_(_558_v16)
				var _587_v38 _dafny.Sequence
				_ = _587_v38
				_587_v38 = _dafny.SeqOf(_558_v16, _558_v16, (_586_v37).Dtor_cf9(), _558_v16, _558_v16)
				_558_v16 = (func() _dafny.Array {
					if _541_v0 {
						return _558_v16
					}
					return (_587_v38).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_587_v38).Cardinality()))).Uint32()).(_dafny.Array)
				})()
				_541_v0 = !(_541_v0)
				var _588_v39 _dafny.Array
				_ = _588_v39
				var _nwElement0_20 bool = _541_v0
				_ = _nwElement0_20
				var _nw96 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(9))
				_ = _nw96
				(_nw96).ArraySet1(_nwElement0_20, 0)
				(_nw96).ArraySet1(false, 1)
				(_nw96).ArraySet1(_541_v0, 2)
				(_nw96).ArraySet1(_541_v0, 3)
				(_nw96).ArraySet1(_541_v0, 4)
				(_nw96).ArraySet1(true, 5)
				(_nw96).ArraySet1(_541_v0, 6)
				(_nw96).ArraySet1(_541_v0, 7)
				(_nw96).ArraySet1(_541_v0, 8)
				_588_v39 = _nw96
				var _589_v40 D0
				_ = _589_v40
				_589_v40 = Companion_D0_.Create_DC1_(_541_v0, _588_v39, false)
				var _590_v41 _dafny.Sequence
				_ = _590_v41
				_590_v41 = _dafny.SeqOf(_589_v40, _589_v40, _589_v40)
				var _591_v42 _dafny.Map
				_ = _591_v42
				_591_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(true, globalState), !(!(_541_v0)))
				var _592_v43 _dafny.Sequence
				_ = _592_v43
				_592_v43 = _dafny.SeqOf((_591_v42).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), true)), _591_v42, _591_v42)
				var _rhs109 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_590_v41, _590_v41)
				_ = _rhs109
				var _rhs110 _dafny.Sequence = _592_v43
				_ = _rhs110
				var _rhs111 _dafny.Int = ((((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _541_v0)).Update(_this.F7(), _541_v0)).Cardinality()).Times(p0)).Times(_dafny.IntOfInt64(364))
				_ = _rhs111
				var _rhs112 _dafny.Sequence = _562_v19
				_ = _rhs112
				var _lhs80 *GlobalState = globalState
				_ = _lhs80
				_590_v41 = _rhs109
				_592_v43 = _rhs110
				_lhs80.F1 = _rhs111
				_562_v19 = _rhs112
			}
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))
			_ = _index102
			(_558_v16).ArraySet1((_this).F6(), (_index102).Int())
			(_this).F7_set_((func() _dafny.Int {
				if _541_v0 {
					return Companion_Default___.Fm1(_541_v0, globalState)
				}
				return (_dafny.Zero).Minus((_this).F6())
			})())
			var _593_v44 _dafny.Array
			_ = _593_v44
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_15
			var _nw97 _dafny.Array
			_ = _nw97
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw97 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) bool = (func(_594_v0 bool) func(_dafny.Int) bool {
					return func(_595_i6 _dafny.Int) bool {
						return _594_v0
					}
				})(_541_v0)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw97 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw97).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw97).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_593_v44 = _nw97
			var _596_v45 D0
			_ = _596_v45
			_596_v45 = Companion_D0_.Create_DC1_(_541_v0, _593_v44, _541_v0)
			var _597_v46 _dafny.Int
			_ = _597_v46
			var _598_v47 bool
			_ = _598_v47
			var _599_v48 _dafny.Array
			_ = _599_v48
			var _out12 _dafny.Int
			_ = _out12
			var _out13 bool
			_ = _out13
			var _out14 _dafny.Array
			_ = _out14
			_out12, _out13, _out14 = Companion_Default___.M0(_596_v45, (_dafny.MultiSetOf(_558_v16)).Cardinality(), globalState)
			_597_v46 = _out12
			_598_v47 = _out13
			_599_v48 = _out14
			var _600_v49 _dafny.Map
			_ = _600_v49
			_600_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_560_v17, _this.F7())
			var _601_v51 _dafny.Map
			_ = _601_v51
			_601_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_541_v0, _541_v0)
			var _602_v52 _dafny.Map
			_ = _602_v52
			_602_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _601_v51)
			(globalState).F1 = ((_600_v49).Merge((_600_v49).Update(func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter23 := _dafny.Iterate((_602_v52).Keys().Elements()); ; {
					_compr_21, _ok23 := _iter23()
					if !_ok23 {
						break
					}
					var _603_v50 _dafny.Int
					_603_v50 = interface{}(_compr_21).(_dafny.Int)
					if (_602_v52).Contains(_603_v50) {
						_coll21.Add((_603_v50).Times(p0), _this.F7())
					}
				}
				return _coll21.ToMap()
			}(), (_this).F6()))).Cardinality()
		}
		r0 = _561_v18
		r1 = p0
		var _604_v54 _dafny.MultiSet
		_ = _604_v54
		_604_v54 = _dafny.MultiSetOf(_575_v26)
		var _605_v55 _dafny.Sequence
		_ = _605_v55
		_605_v55 = _dafny.SeqOf(_604_v54)
		var _606_v56 D1
		_ = _606_v56
		_606_v56 = Companion_D1_.Create_DC4_((_this.F7()).Cmp((func() _dafny.Map {
			var _coll22 = _dafny.NewMapBuilder()
			_ = _coll22
			for _iter24 := _dafny.Iterate(((_605_v55).Select((Companion_Default___.SafeIndex((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_605_v55).Cardinality()))).Uint32()).(_dafny.MultiSet)).Elements()); ; {
				_compr_22, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _607_v53 _dafny.Set
				_607_v53 = interface{}(_compr_22).(_dafny.Set)
				if ((_605_v55).Select((Companion_Default___.SafeIndex((_558_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(139), _dafny.ArrayLen((_558_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_605_v55).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_607_v53) {
					_coll22.Add(_607_v53, _541_v0)
				}
			}
			return _coll22.ToMap()
		}()).Cardinality()) <= 0)
		r2 = _606_v56
		r3 = _541_v0
		return r0, r1, r2, r3
	}
}
func (_this *C1) F13() _dafny.CodePoint {
	{
		return _this._f13
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f7 _dafny.Int
	_f6 _dafny.Int
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f7 = _dafny.Zero
	_this._f6 = _dafny.Zero
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

func (_this *C2) F7() _dafny.Int {
	return _this._f7
}
func (_this *C2) F7_set_(value _dafny.Int) {
	_this._f7 = value
}
func (_this *C2) F6() _dafny.Int {
	return _this._f6
}
func (_this *C2) Ctor__(f6 _dafny.Int, f7 _dafny.Int) {
	{
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C2) Fm23(p0 _dafny.Int, p1 _dafny.Int, p2 D2, globalState *GlobalState) D1 {
	{
		if true {
			return Companion_D1_.Create_DC4_(true)
		} else {
			return Companion_D1_.Create_DC4_(true)
		}
	}
}
func (_this *C2) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _608_v0 D1
		_ = _608_v0
		_608_v0 = Companion_D1_.Create_DC3_(p0, p1, _this.F7())
		if ((func() D1 {
			if p0 {
				return _608_v0
			}
			return _608_v0
		})()).Dtor_cf5() {
			r1 = p0
			if p0 {
				var _609_v1 _dafny.Sequence
				_ = _609_v1
				_609_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ignf")
				var _610_v2 *C0
				_ = _610_v2
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__(p0)
				_610_v2 = _nw98
				var _611_v3 _dafny.Set
				_ = _611_v3
				_611_v3 = _dafny.SetOf(_610_v2, _610_v2, _610_v2)
				var _612_v4 _dafny.Map
				_ = _612_v4
				_612_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_611_v3).Cardinality(), _610_v2.F10)
				var _613_v5 D1
				_ = _613_v5
				_613_v5 = Companion_D1_.Create_DC4_(_610_v2.F10)
				var _614_v6 D2
				_ = _614_v6
				_614_v6 = Companion_D2_.Create_DC7_(p0, (_dafny.Zero).Minus(_this.F7()), _612_v4, _613_v5)
				var _615_v7 _dafny.CodePoint
				_ = _615_v7
				_615_v7 = _dafny.CodePoint('x')
				var _616_v8 _dafny.Map
				_ = _616_v8
				_616_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_615_v7, true)
				var _617_v9 _dafny.Set
				_ = _617_v9
				_617_v9 = _dafny.SetOf(_610_v2.F10, !((Companion_D2_.Create_DC7_(p0, _this.F7(), _612_v4, _613_v5)).Dtor_cf13()))
				var _618_v11 _dafny.Map
				_ = _618_v11
				_618_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v1, _615_v7)
				var _pat_let_tv10 = _617_v9
				_ = _pat_let_tv10
				var _619_v12 _dafny.Map
				_ = _619_v12
				_619_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let17_0 D0) D0 {
					return func(_621_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let18_0 _dafny.Set) D0 {
							return func(_622_dt__update_hcf0_h0 _dafny.Set) D0 {
								return Companion_D0_.Create_DC0_(_622_dt__update_hcf0_h0)
							}(_pat_let18_0)
						}(_pat_let_tv10)
					}(_pat_let17_0)
				}(Companion_Default___.Fm3((_614_v6).Dtor_cf13(), (_616_v8).Cardinality(), _617_v9, func() _dafny.Map {
					var _coll23 = _dafny.NewMapBuilder()
					_ = _coll23
					for _iter25 := _dafny.Iterate((_618_v11).Keys().Elements()); ; {
						_compr_23, _ok25 := _iter25()
						if !_ok25 {
							break
						}
						var _620_v10 _dafny.Sequence
						_620_v10 = interface{}(_compr_23).(_dafny.Sequence)
						if (_618_v11).Contains(_620_v10) {
							_coll23.Add(_620_v10, _this.F7())
						}
					}
					return _coll23.ToMap()
				}(), globalState)), _610_v2.F10)
				var _623_v13 D1
				_ = _623_v13
				_623_v13 = Companion_D1_.Create_DC2_((_619_v12).Merge(_619_v12))
				var _rhs113 _dafny.Int = Companion_Default___.Fm1(_610_v2.F10, globalState)
				_ = _rhs113
				var _rhs114 _dafny.Int = (_this).F6()
				_ = _rhs114
				var _rhs115 _dafny.Sequence = _609_v1
				_ = _rhs115
				var _rhs116 _dafny.Sequence = _609_v1
				_ = _rhs116
				var _rhs117 D1 = Companion_D1_.Create_DC2_(_619_v12)
				_ = _rhs117
				var _lhs81 *C2 = _this
				_ = _lhs81
				_lhs81.F7_set_(_rhs113)
				r2 = _rhs114
				_609_v1 = _rhs115
				_609_v1 = _rhs116
				_623_v13 = _rhs117
				var _624_v14 _dafny.Map
				_ = _624_v14
				_624_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(398), _610_v2)
				var _625_v15 _dafny.Set
				_ = _625_v15
				_625_v15 = _dafny.SetOf(_624_v14, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _610_v2))
				var _626_v16 *C0
				_ = _626_v16
				var _nw99 *C0 = New_C0_()
				_ = _nw99
				_nw99.Ctor__((_625_v15).IsSubsetOf(_dafny.SetOf(_624_v14)))
				_626_v16 = _nw99
				var _627_v17 _dafny.Sequence
				_ = _627_v17
				_627_v17 = _dafny.SeqOf(_626_v16.F10, _610_v2.F10)
				_627_v17 = _dafny.Companion_Sequence_.Concatenate(_627_v17, _627_v17)
				var _628_v18 *C0
				_ = _628_v18
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__(p0)
				_628_v18 = _nw100
				var _629_v19 _dafny.Array
				_ = _629_v19
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_16
				var _nw101 _dafny.Array
				_ = _nw101
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw101 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = func(_630_i0 _dafny.Int) _dafny.Int {
						return (_630_i0).Times(_dafny.IntOfInt64(923))
					}
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw101 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw101).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw101).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_629_v19 = _nw101
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_629_v19), 0))
				_ = _index103
				(_629_v19).ArraySet1(p2, (_index103).Int())
				var _631_v20 _dafny.Array
				_ = _631_v20
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_17
				var _nw102 _dafny.Array
				_ = _nw102
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw102 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.CodePoint = (func(_632_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_633_i1 _dafny.Int) _dafny.CodePoint {
							return _632_v7
						}
					})(_615_v7)
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw102 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw102).ArraySet1CodePoint(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw102).ArraySet1CodePoint(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_631_v20 = _nw102
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_631_v20), 0))
				_ = _index104
				(_631_v20).ArraySet1CodePoint(_615_v7, (_index104).Int())
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_629_v19), 0))
				_ = _index105
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_631_v20), 0))
				_ = _index106
				var _rhs118 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F6(), (func() _dafny.Int {
					if false {
						return _this.F7()
					}
					return _this.F7()
				})())
				_ = _rhs118
				var _rhs119 _dafny.Sequence = _609_v1
				_ = _rhs119
				var _rhs120 _dafny.CodePoint = _615_v7
				_ = _rhs120
				var _rhs121 bool = _610_v2.F10
				_ = _rhs121
				var _lhs82 _dafny.Array = _629_v19
				_ = _lhs82
				var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_629_v19), 0))
				_ = _lhs83
				var _lhs84 _dafny.Array = _631_v20
				_ = _lhs84
				var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(681), _dafny.ArrayLen((_631_v20), 0))
				_ = _lhs85
				var _lhs86 *C0 = _610_v2
				_ = _lhs86
				(_lhs82).ArraySet1(_rhs118, (_lhs83).Int())
				_609_v1 = _rhs119
				(_lhs84).ArraySet1CodePoint(_rhs120, (_lhs85).Int())
				_lhs86.F10 = _rhs121
			} else {
				var _634_v21 _dafny.Array
				_ = _634_v21
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
				_ = _nw103
				_634_v21 = _nw103
				var _635_v22 _dafny.Sequence
				_ = _635_v22
				_635_v22 = _dafny.SeqOf(p1)
				var _636_v23 _dafny.Map
				_ = _636_v23
				_636_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_635_v22, p0)
				var _rhs122 _dafny.Array = _634_v21
				_ = _rhs122
				var _rhs123 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(655))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}(func(_637_i2 _dafny.Int) _dafny.Int {
					return _this.F7()
				})), p0)
				_ = _rhs123
				var _rhs124 _dafny.Array = _634_v21
				_ = _rhs124
				_634_v21 = _rhs122
				_636_v23 = _rhs123
				_634_v21 = _rhs124
				var _638_v24 _dafny.Map
				_ = _638_v24
				_638_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _639_v25 _dafny.Sequence
				_ = _639_v25
				_639_v25 = _dafny.UnicodeSeqOfUtf8Bytes("yujxjkyll")
				var _640_v26 _dafny.Sequence
				_ = _640_v26
				_640_v26 = _dafny.SeqOf(false)
				(_this).M8(_638_v24, Companion_D2_.Create_DC6_(Companion_Default___.Fm24(p0, Companion_D1_.Create_DC4_(p0), _this.F7(), globalState), _639_v25, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_640_v26, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_640_v26).Cardinality()))).Uint32(), !(p0))).Cardinality())), _639_v25, globalState)
				(globalState).F1 = _dafny.IntOfInt64(-712)
				r0 = false
				_638_v24 = (_638_v24).Update(p0, p0)
			}
			var _641_v27 _dafny.Array
			_ = _641_v27
			var _nw104 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
			_ = _nw104
			_641_v27 = _nw104
			_641_v27 = _641_v27
			(globalState).F1 = (((_this).F6()).Plus(p1)).Times((p2).Minus((_this).F6()))
			var _642_v28 _dafny.Array
			_ = _642_v28
			var _len0_18 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_18
			var _nw105 _dafny.Array
			_ = _nw105
			if _len0_18.Cmp(_dafny.Zero) == 0 {
				_nw105 = _dafny.NewArray(_len0_18)
			} else {
				var _init18 func(_dafny.Int) _dafny.Int = func(_643_i3 _dafny.Int) _dafny.Int {
					return (_643_i3).Times((_this).F6())
				}
				_ = _init18
				var _element0_18 = _init18(_dafny.Zero)
				_ = _element0_18
				_nw105 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
				(_nw105).ArraySet1(_element0_18, 0)
				var _nativeLen0_18 = (_len0_18).Int()
				_ = _nativeLen0_18
				for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
					(_nw105).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
				}
			}
			_642_v28 = _nw105
			var _644_v29 _dafny.Set
			_ = _644_v29
			_644_v29 = _dafny.SetOf(p1, (_this).F6(), _this.F7())
			var _rhs125 _dafny.Array = _642_v28
			_ = _rhs125
			var _rhs126 _dafny.Array = _642_v28
			_ = _rhs126
			var _rhs127 _dafny.Array = _642_v28
			_ = _rhs127
			var _rhs128 _dafny.Set = _644_v29
			_ = _rhs128
			_642_v28 = _rhs125
			_642_v28 = _rhs126
			_642_v28 = _rhs127
			_644_v29 = _rhs128
		} else {
			r0 = p0
			var _645_v30 _dafny.Map
			_ = _645_v30
			_645_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.SetOf(p0, p0, p0, p0))
			var _646_v31 _dafny.Set
			_ = _646_v31
			_646_v31 = _dafny.SetOf(true)
			_645_v30 = (_645_v30).Update(_dafny.IntOfInt64(138), (_646_v31).Difference(_646_v31))
			var _647_v32 _dafny.Array
			_ = _647_v32
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_19
			var _nw106 _dafny.Array
			_ = _nw106
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw106 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Sequence = func(_648_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("yqtii")
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
			_647_v32 = _nw106
			_647_v32 = _647_v32
			var _649_v33 D1
			_ = _649_v33
			_649_v33 = Companion_D1_.Create_DC4_(!(!(p0) || (p0)))
			var _source8 D1 = _649_v33
			_ = _source8
			if _source8.Is_DC3() {
				var _650___mcc_h0 bool = _source8.Get_().(D1_DC3).Cf5
				_ = _650___mcc_h0
				var _651___mcc_h1 _dafny.Int = _source8.Get_().(D1_DC3).Cf6
				_ = _651___mcc_h1
				var _652___mcc_h2 _dafny.Int = _source8.Get_().(D1_DC3).Cf7
				_ = _652___mcc_h2
				var _653_cf7 _dafny.Int = _652___mcc_h2
				_ = _653_cf7
				var _654_cf6 _dafny.Int = _651___mcc_h1
				_ = _654_cf6
				var _655_cf5 bool = _650___mcc_h0
				_ = _655_cf5
				var _656_v34 _dafny.Sequence
				_ = _656_v34
				_656_v34 = _dafny.SeqOf(_655_cf5, p0)
				var _657_v35 _dafny.MultiSet
				_ = _657_v35
				_657_v35 = _dafny.MultiSetOf(_656_v34, _656_v34, _656_v34, _dafny.SeqOf(_655_cf5), _dafny.SeqOf(_655_cf5))
				var _658_v36 _dafny.Map
				_ = _658_v36
				_658_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(((_dafny.Zero).Minus(p2)).Cmp(_dafny.IntOfInt64(297)) > 0), Companion_Default___.SafeDivisionInt(_654_cf6, (func() _dafny.Int {
					if (_657_v35).Contains(_dafny.SeqOf(p0)) {
						return (_657_v35).Multiplicity(_dafny.SeqOf(p0))
					}
					return (_this).F6()
				})()))
				_658_v36 = (_658_v36).Update(false, (_dafny.Zero).Minus((_this.F7()).Minus(_this.F7())))
				var _659_v37 *C0
				_ = _659_v37
				var _nw107 *C0 = New_C0_()
				_ = _nw107
				_nw107.Ctor__(_655_cf5)
				_659_v37 = _nw107
				var _660_v38 _dafny.Sequence
				_ = _660_v38
				_660_v38 = _dafny.UnicodeSeqOfUtf8Bytes("nbw")
				var _rhs129 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("tttqannns")
				_ = _rhs129
				var _rhs130 _dafny.Sequence = _660_v38
				_ = _rhs130
				var _rhs131 _dafny.Sequence = _656_v34
				_ = _rhs131
				_660_v38 = _rhs129
				_660_v38 = _rhs130
				_656_v34 = _rhs131
				var _661_v39 _dafny.Array
				_ = _661_v39
				var _nw108 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(7))
				_ = _nw108
				_661_v39 = _nw108
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_661_v39), 0))
				_ = _index107
				(_661_v39).ArraySet1(_653_cf7, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(280), _dafny.ArrayLen((_661_v39), 0))
				_ = _index108
				(_661_v39).ArraySet1(p1, (_index108).Int())
			} else if _source8.Is_DC4() {
				var _662___mcc_h3 bool = _source8.Get_().(D1_DC4).Cf8
				_ = _662___mcc_h3
				var _663_cf8 bool = _662___mcc_h3
				_ = _663_cf8
				var _664_v40 _dafny.Sequence
				_ = _664_v40
				_664_v40 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt(_this.F7(), p2), _this.F7(), (_this).F6())
				_664_v40 = _664_v40
				var _665_v41 _dafny.Map
				_ = _665_v41
				_665_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(p0, globalState), (true) || (p0))
				var _666_v42 _dafny.Map
				_ = _666_v42
				_666_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), p2)
				_665_v41 = (_665_v41).Update((_this).F6(), Companion_Default___.Fm0(_663_cf8, _666_v42, (_this).F6(), globalState))
				var _667_v43 _dafny.Array
				_ = _667_v43
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(23)
				_ = _len0_20
				var _nw109 _dafny.Array
				_ = _nw109
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw109 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) bool = (func(_668_cf8 bool) func(_dafny.Int) bool {
						return func(_669_i5 _dafny.Int) bool {
							return _668_cf8
						}
					})(_663_cf8)
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
				_667_v43 = _nw109
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_667_v43), 0))
				_ = _index109
				(_667_v43).ArraySet1(p0, (_index109).Int())
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_667_v43), 0))
				_ = _index110
				(_667_v43).ArraySet1(p0, (_index110).Int())
				var _670_v44 _dafny.Map
				_ = _670_v44
				_670_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_667_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_667_v43), 0))).Int()).(bool), p0)
				var _671_v45 D0
				_ = _671_v45
				_671_v45 = Companion_D0_.Create_DC0_(_646_v31)
				var _672_v46 _dafny.Map
				_ = _672_v46
				_672_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_671_v45, (_667_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(709), _dafny.ArrayLen((_667_v43), 0))).Int()).(bool))
				var _673_v47 D1
				_ = _673_v47
				_673_v47 = Companion_D1_.Create_DC2_(_672_v46)
				var _674_v48 _dafny.Sequence
				_ = _674_v48
				_674_v48 = _dafny.UnicodeSeqOfUtf8Bytes("vpeiqsqiw")
				var _675_v49 D2
				_ = _675_v49
				_675_v49 = Companion_D2_.Create_DC6_(_673_v47, _674_v48, Companion_Default___.Fm1(Companion_Default___.Fm0(_663_cf8, _666_v42, _this.F7(), globalState), globalState))
				(_this).M8(_670_v44, _675_v49, _dafny.Companion_Sequence_.Concatenate(_674_v48, _dafny.UnicodeSeqOfUtf8Bytes("sibwdulqv")), globalState)
			} else {
				var _676___mcc_h4 _dafny.Map = _source8.Get_().(D1_DC2).Cf4
				_ = _676___mcc_h4
				var _677_cf4 _dafny.Map = _676___mcc_h4
				_ = _677_cf4
				var _678_v50 _dafny.Array
				_ = _678_v50
				var _nw110 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(19))
				_ = _nw110
				_678_v50 = _nw110
				var _679_v51 _dafny.Array
				_ = _679_v51
				var _nwElement0_21 bool = p0
				_ = _nwElement0_21
				var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.One)
				_ = _nw111
				(_nw111).ArraySet1(_nwElement0_21, 0)
				_679_v51 = _nw111
				var _680_v52 _dafny.Set
				_ = _680_v52
				_680_v52 = _dafny.SetOf(_679_v51)
				var _681_v53 _dafny.Array
				_ = _681_v53
				var _nwElement0_22 _dafny.Set = _680_v52
				_ = _nwElement0_22
				var _nw112 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(20))
				_ = _nw112
				(_nw112).ArraySet1(_nwElement0_22, 0)
				(_nw112).ArraySet1(_680_v52, 1)
				(_nw112).ArraySet1(_680_v52, 2)
				(_nw112).ArraySet1(_680_v52, 3)
				(_nw112).ArraySet1(_680_v52, 4)
				(_nw112).ArraySet1(_680_v52, 5)
				(_nw112).ArraySet1(_dafny.SetOf(_679_v51), 6)
				(_nw112).ArraySet1((_680_v52).Difference(_680_v52), 7)
				(_nw112).ArraySet1(_680_v52, 8)
				(_nw112).ArraySet1(_680_v52, 9)
				(_nw112).ArraySet1(_680_v52, 10)
				(_nw112).ArraySet1(_dafny.SetOf(_679_v51), 11)
				(_nw112).ArraySet1(_680_v52, 12)
				(_nw112).ArraySet1(_680_v52, 13)
				(_nw112).ArraySet1(_680_v52, 14)
				(_nw112).ArraySet1(_680_v52, 15)
				(_nw112).ArraySet1(_680_v52, 16)
				(_nw112).ArraySet1(_dafny.SetOf(_679_v51, _679_v51), 17)
				(_nw112).ArraySet1(_dafny.SetOf(_679_v51, _679_v51), 18)
				(_nw112).ArraySet1((_680_v52).Union(_680_v52), 19)
				_681_v53 = _nw112
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_681_v53), 0))
				_ = _index111
				(_681_v53).ArraySet1(_680_v52, (_index111).Int())
				var _682_v54 _dafny.Map
				_ = _682_v54
				_682_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_681_v53), 0))
				_ = _index112
				var _rhs132 _dafny.Array = _678_v50
				_ = _rhs132
				var _rhs133 _dafny.Set = _646_v31
				_ = _rhs133
				var _rhs134 bool = (((_682_v54).Update(p0, p0)).Merge(_682_v54)).Equals((_682_v54).Merge(_682_v54))
				_ = _rhs134
				var _rhs135 _dafny.Set = (func() _dafny.Set {
					if (true) || (p0) {
						return _680_v52
					}
					return _dafny.SetOf(_679_v51)
				})()
				_ = _rhs135
				var _lhs87 _dafny.Array = _681_v53
				_ = _lhs87
				var _lhs88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(763), _dafny.ArrayLen((_681_v53), 0))
				_ = _lhs88
				_678_v50 = _rhs132
				_646_v31 = _rhs133
				r1 = _rhs134
				(_lhs87).ArraySet1(_rhs135, (_lhs88).Int())
				r0 = p0
				var _683_v55 _dafny.Array
				_ = _683_v55
				var _len0_21 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_21
				var _nw113 _dafny.Array
				_ = _nw113
				if _len0_21.Cmp(_dafny.Zero) == 0 {
					_nw113 = _dafny.NewArray(_len0_21)
				} else {
					var _init21 func(_dafny.Int) _dafny.Int = (func(_684_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_685_i6 _dafny.Int) _dafny.Int {
							return (_685_i6).Plus(_684_p2)
						}
					})(p2)
					_ = _init21
					var _element0_21 = _init21(_dafny.Zero)
					_ = _element0_21
					_nw113 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
					(_nw113).ArraySet1(_element0_21, 0)
					var _nativeLen0_21 = (_len0_21).Int()
					_ = _nativeLen0_21
					for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
						(_nw113).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
					}
				}
				_683_v55 = _nw113
				var _686_v56 D2
				_ = _686_v56
				_686_v56 = Companion_D2_.Create_DC5_(_683_v55)
				var _687_v57 D2
				_ = _687_v57
				_687_v57 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_((_686_v56).Dtor_cf9()))
				var _688_v58 _dafny.Map
				_ = _688_v58
				_688_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_687_v57, Companion_Default___.Fm1(p0, globalState))
				_688_v58 = (_688_v58).Update(_687_v57, Companion_Default___.SafeDivisionInt(_this.F7(), (_this).F6()))
				(globalState).F1 = _this.F7()
			}
			var _689_v59 *C0
			_ = _689_v59
			var _nw114 *C0 = New_C0_()
			_ = _nw114
			_nw114.Ctor__((p1).Cmp((_this).F6()) == 0)
			_689_v59 = _nw114
		}
		r1 = (Companion_Default___.Fm1(true, globalState)).Cmp((_this).F6()) < 0
		var _690_v60 _dafny.Map
		_ = _690_v60
		_690_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _691_v61 D1
		_ = _691_v61
		_691_v61 = Companion_D1_.Create_DC4_(p0)
		var _692_v62 _dafny.Map
		_ = _692_v62
		_692_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_691_v61).Dtor_cf8())
		var _hi7 _dafny.Int = (Companion_D2_.Create_DC7_(p0, (_this).F6(), _692_v62, _691_v61)).Dtor_cf14()
		_ = _hi7
		for _693_i7 := (_dafny.Zero).Minus((_this.F7()).Plus((_690_v60).Cardinality())); _693_i7.Cmp(_hi7) < 0; _693_i7 = _693_i7.Plus(_dafny.One) {
			var _694_v63 _dafny.Sequence
			_ = _694_v63
			_694_v63 = _dafny.UnicodeSeqOfUtf8Bytes("xgk")
			var _695_v64 _dafny.Map
			_ = _695_v64
			_695_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _694_v63)
			var _696_v65 _dafny.Sequence
			_ = _696_v65
			_696_v65 = _dafny.SeqOf(_this.F7())
			var _697_v66 _dafny.Array
			_ = _697_v66
			var _nwElement0_23 _dafny.Int = (_this).F6()
			_ = _nwElement0_23
			var _nw115 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(19))
			_ = _nw115
			(_nw115).ArraySet1(_nwElement0_23, 0)
			(_nw115).ArraySet1((func() _dafny.Int {
				if (_608_v0).Dtor_cf5() {
					return (_608_v0).Dtor_cf6()
				}
				return (_this).F6()
			})(), 1)
			(_nw115).ArraySet1((_dafny.Zero).Minus(p2), 2)
			(_nw115).ArraySet1(_this.F7(), 3)
			(_nw115).ArraySet1(p1, 4)
			(_nw115).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_694_v63, (func() _dafny.Sequence {
				if (_695_v64).Contains(_693_i7) {
					return (_695_v64).Get(_693_i7).(_dafny.Sequence)
				}
				return _694_v63
			})())).Cardinality()), 5)
			(_nw115).ArraySet1(((_dafny.Zero).Minus((_this).F6())).Plus(p2), 6)
			(_nw115).ArraySet1(p2, 7)
			(_nw115).ArraySet1(p2, 8)
			(_nw115).ArraySet1(p1, 9)
			(_nw115).ArraySet1(_693_i7, 10)
			(_nw115).ArraySet1((_dafny.Zero).Minus(Companion_Default___.Fm1(p0, globalState)), 11)
			(_nw115).ArraySet1(p2, 12)
			(_nw115).ArraySet1(_693_i7, 13)
			(_nw115).ArraySet1((_dafny.Zero).Minus(_693_i7), 14)
			(_nw115).ArraySet1((_696_v65).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_696_v65).Cardinality()))).Uint32()).(_dafny.Int), 15)
			(_nw115).ArraySet1((_608_v0).Dtor_cf6(), 16)
			(_nw115).ArraySet1((_dafny.Zero).Minus(p2), 17)
			(_nw115).ArraySet1(Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1(p0, globalState), _this.F7()), 18)
			_697_v66 = _nw115
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_697_v66), 0))
			_ = _index113
			(_697_v66).ArraySet1((_this).F6(), (_index113).Int())
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_697_v66), 0))
			_ = _index114
			(_697_v66).ArraySet1(_dafny.IntOfInt64(242), (_index114).Int())
			var _698_v67 _dafny.Map
			_ = _698_v67
			_698_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.IntOfUint32((_694_v63).Cardinality())).Times(p1))
			_698_v67 = (_698_v67).Update(p0, p2)
			var _699_v68 _dafny.Map
			_ = _699_v68
			_699_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_696_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_696_v65).Cardinality()), _dafny.IntOfUint32((_696_v65).Cardinality()))).Uint32()).(_dafny.Int), p1)
			_699_v68 = (_699_v68).Update((_this).F6(), (func() _dafny.Int {
				if (_698_v67).Contains(p0) {
					return (_698_v67).Get(p0).(_dafny.Int)
				}
				return (_this).F6()
			})())
			if p0 {
				var _700_v69 _dafny.Array
				_ = _700_v69
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw116
				_700_v69 = _nw116
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_700_v69), 0))
				_ = _index115
				(_700_v69).ArraySet1(((_dafny.Zero).Minus(_this.F7())).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm0(false, _699_v68, (_this).F6(), globalState), (_697_v66).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_697_v66), 0))).Int()).(_dafny.Int))).Cardinality()) <= 0, (_index115).Int())
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_700_v69), 0))
				_ = _index116
				(_700_v69).ArraySet1(!(p0), (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_697_v66), 0))
				_ = _index117
				(_697_v66).ArraySet1(Companion_Default___.Fm1(p0, globalState), (_index117).Int())
				var _701_v70 _dafny.MultiSet
				_ = _701_v70
				_701_v70 = _dafny.MultiSetOf(_dafny.IntOfUint32((_694_v63).Cardinality()))
				var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_700_v69), 0))
				_ = _index118
				var _rhs136 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.IntOfInt64(145)).Times(p2)), p2)
				_ = _rhs136
				var _rhs137 bool = ((_dafny.IntOfInt64(604)).Cmp((_this).F6()) <= 0) == ((_701_v70).IsDisjointFrom(_701_v70))
				_ = _rhs137
				var _lhs89 *GlobalState = globalState
				_ = _lhs89
				var _lhs90 _dafny.Array = _700_v69
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_700_v69), 0))
				_ = _lhs91
				_lhs89.F1 = _rhs136
				(_lhs90).ArraySet1(_rhs137, (_lhs91).Int())
				var _702_v71 _dafny.CodePoint
				_ = _702_v71
				_702_v71 = _dafny.CodePoint('i')
				var _703_v72 _dafny.Sequence
				_ = _703_v72
				_703_v72 = _dafny.SeqOf(_694_v63)
				var _704_v73 _dafny.Sequence
				_ = _704_v73
				_704_v73 = _dafny.SeqOf((_700_v69).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_700_v69), 0))).Int()).(bool))
				var _705_v74 _dafny.Set
				_ = _705_v74
				_705_v74 = _dafny.SetOf(_694_v63, _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_706_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				})), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_707_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()))).Uint32(), _702_v71), (_703_v72).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_704_v73).Cardinality()), _dafny.IntOfUint32((_703_v72).Cardinality()))).Uint32()).(_dafny.Sequence))
				_705_v74 = _705_v74
				var _708_v75 _dafny.Set
				_ = _708_v75
				_708_v75 = _dafny.SetOf(_696_v65)
				var _rhs138 bool = (_708_v75).IsProperSubsetOf(_708_v75)
				_ = _rhs138
				var _rhs139 _dafny.Int = (_this).F6()
				_ = _rhs139
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				r0 = _rhs138
				_lhs92.F1 = _rhs139
			} else {
				var _709_v76 T0
				_ = _709_v76
				var _nw117 *C1 = New_C1_()
				_ = _nw117
				_nw117.Ctor__(_dafny.CodePoint('x'), _693_i7, ((_690_v60).Cardinality()).Minus(p1))
				_709_v76 = _nw117
				var _710_v77 D8
				_ = _710_v77
				_710_v77 = Companion_D8_.Create_DC18_(_709_v76)
				var _711_v78 D8
				_ = _711_v78
				_711_v78 = Companion_D8_.Create_DC18_((_710_v77).Dtor_cf33())
				_709_v76 = (_710_v77).Dtor_cf33()
				var _712_v79 *C0
				_ = _712_v79
				var _nw118 *C0 = New_C0_()
				_ = _nw118
				_nw118.Ctor__((p0) || (p0))
				_712_v79 = _nw118
				_697_v66 = _697_v66
				var _713_v80 _dafny.Set
				_ = _713_v80
				_713_v80 = _dafny.SetOf(p0)
				var _714_v81 D5
				_ = _714_v81
				_714_v81 = Companion_D5_.Create_DC12_((_713_v80).Cardinality())
				var _715_v82 _dafny.Sequence
				_ = _715_v82
				_715_v82 = _dafny.SeqOf(_712_v79.F10)
				_714_v81 = Companion_Default___.Fm28(_715_v82, (_dafny.Zero).Minus(_693_i7), globalState)
				var _716_v83 _dafny.Array
				_ = _716_v83
				var _nwElement0_24 bool = p0
				_ = _nwElement0_24
				var _nw119 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(12))
				_ = _nw119
				(_nw119).ArraySet1(_nwElement0_24, 0)
				(_nw119).ArraySet1(p0, 1)
				(_nw119).ArraySet1(_712_v79.F10, 2)
				(_nw119).ArraySet1(_712_v79.F10, 3)
				(_nw119).ArraySet1(p0, 4)
				(_nw119).ArraySet1(p0, 5)
				(_nw119).ArraySet1(false, 6)
				(_nw119).ArraySet1(p0, 7)
				(_nw119).ArraySet1(p0, 8)
				(_nw119).ArraySet1(p0, 9)
				(_nw119).ArraySet1(p0, 10)
				(_nw119).ArraySet1(_712_v79.F10, 11)
				_716_v83 = _nw119
				var _717_v84 _dafny.Sequence
				_ = _717_v84
				_717_v84 = _dafny.SeqOf(_716_v83)
				var _718_v85 _dafny.Array
				_ = _718_v85
				var _nwElement0_25 _dafny.Array = _716_v83
				_ = _nwElement0_25
				var _nw120 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(11))
				_ = _nw120
				(_nw120).ArraySet1(_nwElement0_25, 0)
				(_nw120).ArraySet1(_716_v83, 1)
				(_nw120).ArraySet1(_716_v83, 2)
				(_nw120).ArraySet1(_716_v83, 3)
				(_nw120).ArraySet1(_716_v83, 4)
				(_nw120).ArraySet1(_716_v83, 5)
				(_nw120).ArraySet1(_716_v83, 6)
				(_nw120).ArraySet1((_717_v84).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_717_v84).Cardinality()))).Uint32()).(_dafny.Array), 7)
				(_nw120).ArraySet1(_716_v83, 8)
				(_nw120).ArraySet1(_716_v83, 9)
				(_nw120).ArraySet1(_716_v83, 10)
				_718_v85 = _nw120
				var _719_v86 _dafny.Array
				_ = _719_v86
				var _nw121 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw121
				_719_v86 = _nw121
				var _720_v87 _dafny.CodePoint
				_ = _720_v87
				_720_v87 = _dafny.CodePoint('o')
				var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_719_v86), 0))
				_ = _index119
				(_719_v86).ArraySet1CodePoint(_720_v87, (_index119).Int())
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_719_v86), 0))
				_ = _index120
				var _rhs140 _dafny.Array = _718_v85
				_ = _rhs140
				var _rhs141 bool = (func() bool {
					if (_690_v60).Contains(p0) {
						return (_690_v60).Get(p0).(bool)
					}
					return p0
				})()
				_ = _rhs141
				var _rhs142 _dafny.CodePoint = _720_v87
				_ = _rhs142
				var _rhs143 T0 = _709_v76
				_ = _rhs143
				var _rhs144 _dafny.Int = (_709_v76).F6()
				_ = _rhs144
				var _lhs93 *C0 = _712_v79
				_ = _lhs93
				var _lhs94 _dafny.Array = _719_v86
				_ = _lhs94
				var _lhs95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(866), _dafny.ArrayLen((_719_v86), 0))
				_ = _lhs95
				var _lhs96 *GlobalState = globalState
				_ = _lhs96
				_718_v85 = _rhs140
				_lhs93.F10 = _rhs141
				(_lhs94).ArraySet1CodePoint(_rhs142, (_lhs95).Int())
				_709_v76 = _rhs143
				_lhs96.F1 = _rhs144
			}
		}
		var _hi8 _dafny.Int = p2
		_ = _hi8
		for _721_i9 := (_this).F6(); _721_i9.Cmp(_hi8) < 0; _721_i9 = _721_i9.Plus(_dafny.One) {
			r0 = !(p0)
			if !(p0) {
				var _722_v88 _dafny.Sequence
				_ = _722_v88
				_722_v88 = _dafny.SeqOf(_721_i9)
				var _723_v89 _dafny.Array
				_ = _723_v89
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(12)
				_ = _len0_22
				var _nw122 _dafny.Array
				_ = _nw122
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw122 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Int = func(_724_i10 _dafny.Int) _dafny.Int {
						return (_724_i10).Times(_dafny.IntOfInt64(397))
					}
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw122 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw122).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw122).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_723_v89 = _nw122
				var _725_v90 _dafny.Set
				_ = _725_v90
				_725_v90 = _dafny.SetOf(_722_v88)
				var _726_v91 _dafny.CodePoint
				_ = _726_v91
				_726_v91 = _dafny.CodePoint('i')
				var _727_v92 _dafny.MultiSet
				_ = _727_v92
				_727_v92 = _dafny.MultiSetOf(_726_v91, _726_v91)
				var _728_v93 _dafny.Sequence
				_ = _728_v93
				_728_v93 = _dafny.SeqOf(p0)
				var _729_v94 _dafny.Array
				_ = _729_v94
				var _nwElement0_26 _dafny.Int = p2
				_ = _nwElement0_26
				var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(23))
				_ = _nw123
				(_nw123).ArraySet1(_nwElement0_26, 0)
				(_nw123).ArraySet1(p1, 1)
				(_nw123).ArraySet1(p2, 2)
				(_nw123).ArraySet1(_dafny.IntOfInt64(687), 3)
				(_nw123).ArraySet1((_dafny.Zero).Minus(_721_i9), 4)
				(_nw123).ArraySet1(p1, 5)
				(_nw123).ArraySet1(Companion_Default___.Fm1(!(p0), globalState), 6)
				(_nw123).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_722_v88).Cardinality()))), 7)
				(_nw123).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_723_v89, _723_v89)).Cardinality()), 8)
				(_nw123).ArraySet1(_721_i9, 9)
				(_nw123).ArraySet1((_725_v90).Cardinality(), 10)
				(_nw123).ArraySet1(p2, 11)
				(_nw123).ArraySet1((_727_v92).Cardinality(), 12)
				(_nw123).ArraySet1(_dafny.IntOfUint32((_728_v93).Cardinality()), 13)
				(_nw123).ArraySet1((_this).F6(), 14)
				(_nw123).ArraySet1(_dafny.IntOfUint32((_722_v88).Cardinality()), 15)
				(_nw123).ArraySet1(_dafny.IntOfInt64(503), 16)
				(_nw123).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()), 17)
				(_nw123).ArraySet1(p2, 18)
				(_nw123).ArraySet1(_dafny.IntOfInt64(835), 19)
				(_nw123).ArraySet1(_this.F7(), 20)
				(_nw123).ArraySet1((_this).F6(), 21)
				(_nw123).ArraySet1(_dafny.IntOfInt64(610), 22)
				_729_v94 = _nw123
				var _730_v95 _dafny.Map
				_ = _730_v95
				_730_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_723_v89, p0)
				var _731_v96 D2
				_ = _731_v96
				_731_v96 = Companion_D2_.Create_DC5_(_723_v89)
				r0 = (func() bool {
					if (_730_v95).Contains((_731_v96).Dtor_cf9()) {
						return (_730_v95).Get((_731_v96).Dtor_cf9()).(bool)
					}
					return p0
				})()
				var _732_v97 *C0
				_ = _732_v97
				var _nw124 *C0 = New_C0_()
				_ = _nw124
				_nw124.Ctor__(p0)
				_732_v97 = _nw124
				var _733_v98 _dafny.MultiSet
				_ = _733_v98
				_733_v98 = _dafny.MultiSetOf((p1).Times(_721_i9))
				var _734_v99 D6
				_ = _734_v99
				_734_v99 = Companion_D6_.Create_DC14_(p1, _dafny.MultiSetOf(p0, (func() bool {
					if (_690_v60).Contains(_732_v97.F10) {
						return (_690_v60).Get(_732_v97.F10).(bool)
					}
					return _732_v97.F10
				})()), (_690_v60).Cardinality())
				var _735_v100 _dafny.Map
				_ = _735_v100
				_735_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_734_v99).Dtor_cf23(), p1), p0)
				var _rhs145 bool = (func() bool {
					if p0 {
						return false
					}
					return p0
				})()
				_ = _rhs145
				var _rhs146 _dafny.Int = (func() _dafny.Int {
					if (_733_v98).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_722_v88, _722_v88)).Cardinality())) {
						return (_733_v98).Multiplicity(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_722_v88, _722_v88)).Cardinality()))
					}
					return (_735_v100).Cardinality()
				})()
				_ = _rhs146
				var _rhs147 *C0 = _732_v97
				_ = _rhs147
				r0 = _rhs145
				r2 = _rhs146
				_732_v97 = _rhs147
				var _736_v101 _dafny.Map
				_ = _736_v101
				_736_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_732_v97.F10, _dafny.UnicodeSeqOfUtf8Bytes("ksw"))
				var _737_v102 _dafny.Sequence
				_ = _737_v102
				_737_v102 = _dafny.UnicodeSeqOfUtf8Bytes("mqxjodfld")
				_736_v101 = (_736_v101).Update(_732_v97.F10, _737_v102)
				var _738_v103 _dafny.Array
				_ = _738_v103
				var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
				_ = _nw125
				_738_v103 = _nw125
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_738_v103), 0))
				_ = _index121
				(_738_v103).ArraySet1CodePoint((func() _dafny.CodePoint {
					if p0 {
						return _726_v91
					}
					return _dafny.CodePoint('j')
				})(), (_index121).Int())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(523), _dafny.ArrayLen((_738_v103), 0))
				_ = _index122
				(_738_v103).ArraySet1CodePoint(_726_v91, (_index122).Int())
				var _739_v104 _dafny.Array
				_ = _739_v104
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(3)
				_ = _len0_23
				var _nw126 _dafny.Array
				_ = _nw126
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw126 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) bool = (func(_740_v97 *C0) func(_dafny.Int) bool {
						return func(_741_i11 _dafny.Int) bool {
							return _740_v97.F10
						}
					})(_732_v97)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw126 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw126).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw126).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_739_v104 = _nw126
				var _742_v105 _dafny.Map
				_ = _742_v105
				_742_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_739_v104, (Companion_Default___.Fm1(_732_v97.F10, globalState)).Cmp(_721_i9) > 0)
				_742_v105 = (_742_v105).Update(_739_v104, (_728_v93).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_728_v93).Cardinality()))).Uint32()).(bool))
			} else {
				r1 = p0
				var _743_v106 _dafny.Set
				_ = _743_v106
				_743_v106 = _dafny.SetOf(p0, true)
				var _744_v107 D0
				_ = _744_v107
				_744_v107 = Companion_D0_.Create_DC0_(_743_v106)
				var _745_v108 D1
				_ = _745_v108
				_745_v108 = Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_744_v107, p0))
				var _746_v109 _dafny.Sequence
				_ = _746_v109
				_746_v109 = _dafny.UnicodeSeqOfUtf8Bytes("fbkeac")
				var _747_v110 D2
				_ = _747_v110
				_747_v110 = Companion_D2_.Create_DC6_(_745_v108, _746_v109, p1)
				(_this).M8((_690_v60).Merge(_690_v60), _747_v110, _dafny.Companion_Sequence_.Concatenate(_746_v109, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}(func(_748_i12 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('i')
				}))), globalState)
				var _749_v111 _dafny.MultiSet
				_ = _749_v111
				_749_v111 = _dafny.MultiSetOf(p2, _721_i9)
				var _750_v112 _dafny.Map
				_ = _750_v112
				_750_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_749_v111).Cardinality(), _this.F7())
				r1 = Companion_Default___.Fm0(p0, _750_v112, _721_i9, globalState)
				var _751_v113 D9
				_ = _751_v113
				_751_v113 = Companion_D9_.Create_DC22_(_dafny.SeqOf(_this.F7()))
				var _752_v114 _dafny.Map
				_ = _752_v114
				_752_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_749_v111, (p1).Cmp(_dafny.IntOfUint32(((_751_v113).Dtor_cf45()).Cardinality())) != 0)
				_752_v114 = (_752_v114).Update(_749_v111, p0)
				var _753_v115 _dafny.Array
				_ = _753_v115
				var _len0_24 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_24
				var _nw127 _dafny.Array
				_ = _nw127
				if _len0_24.Cmp(_dafny.Zero) == 0 {
					_nw127 = _dafny.NewArray(_len0_24)
				} else {
					var _init24 func(_dafny.Int) _dafny.MultiSet = func(_754_i13 _dafny.Int) _dafny.MultiSet {
						return _dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.CodePoint('f')))
					}
					_ = _init24
					var _element0_24 = _init24(_dafny.Zero)
					_ = _element0_24
					_nw127 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
					(_nw127).ArraySet1(_element0_24, 0)
					var _nativeLen0_24 = (_len0_24).Int()
					_ = _nativeLen0_24
					for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
						(_nw127).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
					}
				}
				_753_v115 = _nw127
				var _755_v116 _dafny.CodePoint
				_ = _755_v116
				_755_v116 = _dafny.CodePoint('b')
				var _756_v117 _dafny.MultiSet
				_ = _756_v117
				_756_v117 = _dafny.MultiSetOf(_755_v116, _dafny.CodePoint('t'), _755_v116, _755_v116, _755_v116)
				var _757_v118 _dafny.MultiSet
				_ = _757_v118
				_757_v118 = _dafny.MultiSetOf(_756_v117, _dafny.MultiSetOf(_755_v116), _756_v117)
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_753_v115), 0))
				_ = _index123
				(_753_v115).ArraySet1((_757_v118).Union(_757_v118), (_index123).Int())
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(63), _dafny.ArrayLen((_753_v115), 0))
				_ = _index124
				(_753_v115).ArraySet1(_757_v118, (_index124).Int())
			}
			(_this).F7_set_(p1)
			var _758_v119 _dafny.Map
			_ = _758_v119
			_758_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
			var _759_v120 _dafny.Sequence
			_ = _759_v120
			_759_v120 = _dafny.SeqOf(p0)
			var _760_v121 _dafny.Sequence
			_ = _760_v121
			_760_v121 = _dafny.UnicodeSeqOfUtf8Bytes("kifqmnwow")
			var _761_v122 _dafny.Array
			_ = _761_v122
			var _nwElement0_27 _dafny.Int = p1
			_ = _nwElement0_27
			var _nw128 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(17))
			_ = _nw128
			(_nw128).ArraySet1(_nwElement0_27, 0)
			(_nw128).ArraySet1(p2, 1)
			(_nw128).ArraySet1(Companion_Default___.Fm1(p0, globalState), 2)
			(_nw128).ArraySet1(_this.F7(), 3)
			(_nw128).ArraySet1((_this).F6(), 4)
			(_nw128).ArraySet1((_dafny.Zero).Minus((_this).F6()), 5)
			(_nw128).ArraySet1(Companion_Default___.SafeDivisionInt(p1, p2), 6)
			(_nw128).ArraySet1(p2, 7)
			(_nw128).ArraySet1((func() _dafny.Int {
				if p0 {
					return (_dafny.Zero).Minus(p2)
				}
				return _721_i9
			})(), 8)
			(_nw128).ArraySet1((p1).Plus((func() _dafny.Int {
				if (_758_v119).Contains(p0) {
					return (_758_v119).Get(p0).(_dafny.Int)
				}
				return (Companion_Default___.Fm30((_759_v120).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_759_v120).Cardinality()))).Uint32()).(bool), p0, p1, p1, globalState)).Cardinality()
			})()), 9)
			(_nw128).ArraySet1((_721_i9).Times(p1), 10)
			(_nw128).ArraySet1(_721_i9, 11)
			(_nw128).ArraySet1((_721_i9).Plus(p2), 12)
			(_nw128).ArraySet1(_dafny.IntOfInt64(-181), 13)
			(_nw128).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(116))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg34 _dafny.Int) interface{} {
					return coer34(arg34)
				}
			}(func(_762_i14 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			})), _760_v121)).Cardinality()), 14)
			(_nw128).ArraySet1(_dafny.IntOfInt64(767), 15)
			(_nw128).ArraySet1(_dafny.IntOfInt64(277), 16)
			_761_v122 = _nw128
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_761_v122), 0))
			_ = _index125
			(_761_v122).ArraySet1(_this.F7(), (_index125).Int())
			var _763_v123 _dafny.Sequence
			_ = _763_v123
			_763_v123 = _dafny.SeqOf(_721_i9)
			var _764_v124 _dafny.Map
			_ = _764_v124
			_764_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_721_i9), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_763_v123).Select((Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), p0)).Cardinality(), _dafny.IntOfUint32((_763_v123).Cardinality()))).Uint32()).(_dafny.Int), true)).Cardinality())
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(704), _dafny.ArrayLen((_761_v122), 0))
			_ = _index126
			(_761_v122).ArraySet1(Companion_Default___.SafeDivisionInt((Companion_Default___.Fm1(p0, globalState)).Times((_690_v60).Cardinality()), ((_dafny.Zero).Minus((func() _dafny.Int {
				if (_758_v119).Contains(p0) {
					return (_758_v119).Get(p0).(_dafny.Int)
				}
				return (func() _dafny.Int {
					if (_764_v124).Contains(p2) {
						return (_764_v124).Get(p2).(_dafny.Int)
					}
					return Companion_Default___.Fm1(p0, globalState)
				})()
			})())).Times((_this).F6())), (_index126).Int())
		}
		var _765_v125 _dafny.Map
		_ = _765_v125
		_765_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F6())
		var _766_v126 _dafny.Sequence
		_ = _766_v126
		_766_v126 = _dafny.SeqOf((_765_v125).Cardinality())
		var _hi9 _dafny.Int = p1
		_ = _hi9
		for _767_i15 := ((_766_v126).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_766_v126).Cardinality()))).Uint32()).(_dafny.Int)).Times(p1); _767_i15.Cmp(_hi9) < 0; _767_i15 = _767_i15.Plus(_dafny.One) {
			if (p0) || ((func() bool {
				if (_608_v0).Dtor_cf5() {
					return p0
				}
				return false
			})()) {
				var _768_v127 _dafny.Sequence
				_ = _768_v127
				_768_v127 = _dafny.UnicodeSeqOfUtf8Bytes("vg")
				(_this).M9(p0, (func() _dafny.Int {
					if !(p0) {
						return _dafny.IntOfUint32((_768_v127).Cardinality())
					}
					return _767_i15
				})(), p1, _dafny.IntOfInt64(-200), globalState)
				var _769_v128 _dafny.Array
				_ = _769_v128
				var _nwElement0_28 _dafny.Int = _this.F7()
				_ = _nwElement0_28
				var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(6))
				_ = _nw129
				(_nw129).ArraySet1(_nwElement0_28, 0)
				(_nw129).ArraySet1(_767_i15, 1)
				(_nw129).ArraySet1(_767_i15, 2)
				(_nw129).ArraySet1((func() _dafny.Int {
					if p0 {
						return (_dafny.Zero).Minus(_dafny.IntOfInt64(-823))
					}
					return p1
				})(), 3)
				(_nw129).ArraySet1(p2, 4)
				(_nw129).ArraySet1(p1, 5)
				_769_v128 = _nw129
				_769_v128 = _769_v128
				var _770_v129 *C0
				_ = _770_v129
				var _nw130 *C0 = New_C0_()
				_ = _nw130
				_nw130.Ctor__((p0) == (p0))
				_770_v129 = _nw130
				(_770_v129).F10 = p0
				(_this).F7_set_((_dafny.Zero).Minus((func() _dafny.Int {
					if (func() bool {
						if (_692_v62).Contains(_767_i15) {
							return (_692_v62).Get(_767_i15).(bool)
						}
						return p0
					})() {
						return _767_i15
					}
					return (_this.F7()).Minus(_767_i15)
				})()))
			} else {
				var _771_v130 _dafny.Array
				_ = _771_v130
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
				_ = _nw131
				_771_v130 = _nw131
				var _772_v131 _dafny.Sequence
				_ = _772_v131
				_772_v131 = _dafny.UnicodeSeqOfUtf8Bytes("oiiryjckw")
				var _rhs148 bool = _dafny.Companion_Sequence_.IsPrefixOf(_772_v131, _772_v131)
				_ = _rhs148
				var _rhs149 _dafny.Array = _771_v130
				_ = _rhs149
				r1 = _rhs148
				_771_v130 = _rhs149
				var _773_v132 _dafny.CodePoint
				_ = _773_v132
				_773_v132 = _dafny.CodePoint('t')
				var _774_v133 D7
				_ = _774_v133
				_774_v133 = Companion_D7_.Create_DC17_(_772_v131, Companion_Default___.Fm31((_this).F6(), p1, _this.F7(), _773_v132, globalState))
				var _775_v134 _dafny.Sequence
				_ = _775_v134
				_775_v134 = _dafny.SeqOf(p0)
				var _776_v135 _dafny.MultiSet
				_ = _776_v135
				_776_v135 = _dafny.MultiSetOf(_dafny.IntOfUint32((_775_v134).Cardinality()), _dafny.IntOfInt64(885), _dafny.IntOfUint32((_775_v134).Cardinality()))
				var _777_v136 _dafny.MultiSet
				_ = _777_v136
				_777_v136 = _dafny.MultiSetOf((_776_v135).Cardinality())
				var _778_v137 _dafny.Map
				_ = _778_v137
				_778_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_766_v126).Cardinality()))
				var _779_v138 T0
				_ = _779_v138
				var _nw132 *C1 = New_C1_()
				_ = _nw132
				_nw132.Ctor__(_773_v132, _this.F7(), (func() _dafny.Int {
					if (_777_v136).Contains(_767_i15) {
						return (_777_v136).Multiplicity(_767_i15)
					}
					return (_778_v137).Cardinality()
				})())
				_779_v138 = _nw132
				var _780_v139 _dafny.Map
				_ = _780_v139
				_780_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_774_v133, _779_v138)
				var _781_v140 _dafny.Set
				_ = _781_v140
				_781_v140 = _dafny.SetOf(p0, p0)
				var _782_v141 D0
				_ = _782_v141
				_782_v141 = Companion_D0_.Create_DC0_(_781_v140)
				var _783_v142 _dafny.Map
				_ = _783_v142
				_783_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_782_v141, p0)
				var _784_v143 D1
				_ = _784_v143
				_784_v143 = Companion_D1_.Create_DC2_(_783_v142)
				var _785_v144 D2
				_ = _785_v144
				_785_v144 = Companion_D2_.Create_DC6_(_784_v143, _772_v131, (_779_v138).F6())
				var _786_v145 D2
				_ = _786_v145
				_786_v145 = Companion_D2_.Create_DC7_(!(p0), _this.F7(), _692_v62, (_this).Fm23(_dafny.IntOfInt64(260), p1, _785_v144, globalState))
				var _787_v146 _dafny.Set
				_ = _787_v146
				_787_v146 = _dafny.SetOf(_this.F7())
				var _rhs150 _dafny.Map = _780_v139
				_ = _rhs150
				var _rhs151 _dafny.Int = (_dafny.IntOfUint32((_775_v134).Cardinality())).Plus(((_this).F6()).Plus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, true)).Cardinality()))
				_ = _rhs151
				var _rhs152 bool = ((_786_v145).Dtor_cf14()).Cmp(Companion_Default___.SafeModuloInt((_787_v146).Cardinality(), _779_v138.F7())) <= 0
				_ = _rhs152
				var _rhs153 bool = p0
				_ = _rhs153
				var _lhs97 *C2 = _this
				_ = _lhs97
				_780_v139 = _rhs150
				_lhs97.F7_set_(_rhs151)
				r0 = _rhs152
				r0 = _rhs153
				r0 = (func() bool {
					if p0 {
						return p0
					}
					return ((_779_v138).F6()).Cmp(_dafny.IntOfUint32((_772_v131).Cardinality())) > 0
				})()
				r1 = p0
				r0 = p0
			}
			var _788_v147 _dafny.Array
			_ = _788_v147
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_25
			var _nw133 _dafny.Array
			_ = _nw133
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw133 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Int = (func(_789_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_790_i16 _dafny.Int) _dafny.Int {
						return (_790_i16).Plus(_789_p2)
					}
				})(p2)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw133 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw133).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw133).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_788_v147 = _nw133
			var _791_v148 _dafny.Sequence
			_ = _791_v148
			_791_v148 = _dafny.UnicodeSeqOfUtf8Bytes("jaegurgi")
			var _792_v149 _dafny.CodePoint
			_ = _792_v149
			_792_v149 = _dafny.CodePoint('s')
			var _793_v150 _dafny.Map
			_ = _793_v150
			_793_v150 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _this.F7())
			var _rhs154 _dafny.Array = _788_v147
			_ = _rhs154
			var _rhs155 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("amfxpo"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(295), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("amfxpo")).Cardinality()))).Uint32(), _792_v149)
			_ = _rhs155
			var _rhs156 _dafny.Array = _788_v147
			_ = _rhs156
			var _rhs157 bool = p0
			_ = _rhs157
			var _rhs158 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_766_v126, _766_v126), (Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_766_v126, _766_v126)).Cardinality()))).Uint32(), Companion_Default___.SafeDivisionInt((_this).F6(), (_793_v150).Cardinality()))
			_ = _rhs158
			_788_v147 = _rhs154
			_791_v148 = _rhs155
			_788_v147 = _rhs156
			r1 = _rhs157
			_766_v126 = _rhs158
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_788_v147), 0))
			_ = _index127
			(_788_v147).ArraySet1(((_this).F6()).Times(_767_i15), (_index127).Int())
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(488), _dafny.ArrayLen((_788_v147), 0))
			_ = _index128
			(_788_v147).ArraySet1(Companion_Default___.SafeModuloInt(p1, (_this).F6()), (_index128).Int())
			var _794_v151 *C0
			_ = _794_v151
			var _nw134 *C0 = New_C0_()
			_ = _nw134
			_nw134.Ctor__(p0)
			_794_v151 = _nw134
		}
		var _795_v152 _dafny.Sequence
		_ = _795_v152
		_795_v152 = _dafny.UnicodeSeqOfUtf8Bytes("akifd")
		var _796_v153 _dafny.Sequence
		_ = _796_v153
		_796_v153 = _dafny.SeqOf(_dafny.SetOf(_dafny.IntOfUint32((_795_v152).Cardinality())))
		if Companion_Default___.Fm0(!((_796_v153).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_796_v153).Cardinality()))).Uint32()).(_dafny.Set)).Contains((_this).F6()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F7())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_765_v125).Cardinality(), p2)), p2, globalState) {
			var _797_v154 _dafny.Sequence
			_ = _797_v154
			_797_v154 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer35 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg35 _dafny.Int) interface{} {
					return coer35(arg35)
				}
			}((func(_798_p1 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_799_i18 _dafny.Int) _dafny.Int {
					return _798_p1
				}
			})(p1))), _766_v126, _766_v126, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-931))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg36 _dafny.Int) interface{} {
					return coer36(arg36)
				}
			}(func(_800_i19 _dafny.Int) _dafny.Int {
				return _this.F7()
			})), _766_v126)
			var _801_v155 _dafny.Array
			_ = _801_v155
			var _nwElement0_29 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(259))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg37 _dafny.Int) interface{} {
					return coer37(arg37)
				}
			}(func(_802_i17 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('g')
			})), _795_v152)
			_ = _nwElement0_29
			var _nw135 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(3))
			_ = _nw135
			(_nw135).ArraySet1(_nwElement0_29, 0)
			(_nw135).ArraySet1(Companion_Default___.Fm2((_797_v154).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_797_v154).Cardinality()))).Uint32()).(_dafny.Sequence), _this.F7(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(873))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}(func(_803_i20 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(184)
			}))).Cardinality()), p0, globalState), 1)
			(_nw135).ArraySet1(_795_v152, 2)
			_801_v155 = _nw135
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_801_v155), 0))
			_ = _index129
			(_801_v155).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_795_v152, _795_v152), (_index129).Int())
			var _804_v156 _dafny.Array
			_ = _804_v156
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
			_ = _nw136
			_804_v156 = _nw136
			var _805_v157 _dafny.Array
			_ = _805_v157
			var _nw137 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
			_ = _nw137
			_805_v157 = _nw137
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_804_v156), 0))
			_ = _index130
			(_804_v156).ArraySet1(_805_v157, (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_801_v155), 0))
			_ = _index131
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_804_v156), 0))
			_ = _index132
			var _rhs159 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("hv")
			_ = _rhs159
			var _rhs160 _dafny.Array = _805_v157
			_ = _rhs160
			var _lhs98 _dafny.Array = _801_v155
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(403), _dafny.ArrayLen((_801_v155), 0))
			_ = _lhs99
			var _lhs100 _dafny.Array = _804_v156
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(578), _dafny.ArrayLen((_804_v156), 0))
			_ = _lhs101
			(_lhs98).ArraySet1(_rhs159, (_lhs99).Int())
			(_lhs100).ArraySet1(_rhs160, (_lhs101).Int())
			(globalState).F1 = (_this).F6()
			var _806_v158 _dafny.Array
			_ = _806_v158
			var _len0_26 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_26
			var _nw138 _dafny.Array
			_ = _nw138
			if _len0_26.Cmp(_dafny.Zero) == 0 {
				_nw138 = _dafny.NewArray(_len0_26)
			} else {
				var _init26 func(_dafny.Int) D9 = (func(_807_v126 _dafny.Sequence) func(_dafny.Int) D9 {
					return func(_808_i21 _dafny.Int) D9 {
						return Companion_D9_.Create_DC22_(_807_v126)
					}
				})(_766_v126)
				_ = _init26
				var _element0_26 = _init26(_dafny.Zero)
				_ = _element0_26
				_nw138 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
				(_nw138).ArraySet1(_element0_26, 0)
				var _nativeLen0_26 = (_len0_26).Int()
				_ = _nativeLen0_26
				for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
					(_nw138).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
				}
			}
			_806_v158 = _nw138
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_806_v158), 0))
			_ = _index133
			(_806_v158).ArraySet1(Companion_D9_.Create_DC22_(_766_v126), (_index133).Int())
			var _809_v159 _dafny.CodePoint
			_ = _809_v159
			_809_v159 = _dafny.CodePoint('u')
			var _810_v160 _dafny.Map
			_ = _810_v160
			_810_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
			var _811_v161 D5
			_ = _811_v161
			_811_v161 = Companion_D5_.Create_DC12_(_dafny.IntOfUint32((_dafny.SeqOf(_810_v160)).Cardinality()))
			var _812_v162 _dafny.Sequence
			_ = _812_v162
			_812_v162 = _dafny.SeqOf(_811_v161, _811_v161)
			var _813_v163 _dafny.Sequence
			_ = _813_v163
			_813_v163 = _dafny.SeqOf(p0, p0, p0, true, p0)
			var _814_v164 D6
			_ = _814_v164
			_814_v164 = Companion_D6_.Create_DC15_(p0, p2, _dafny.IntOfInt64(79), p2)
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_806_v158), 0))
			_ = _index134
			var _rhs161 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_812_v162).Cardinality()), p1)
			_ = _rhs161
			var _rhs162 D9 = Companion_Default___.Fm32(p0, !(!(_dafny.Companion_Sequence_.IsPrefixOf(_813_v163, _813_v163))), p0, p1, globalState)
			_ = _rhs162
			var _rhs163 _dafny.CodePoint = _809_v159
			_ = _rhs163
			var _rhs164 bool = (_814_v164).Dtor_cf26()
			_ = _rhs164
			var _lhs102 *GlobalState = globalState
			_ = _lhs102
			var _lhs103 _dafny.Array = _806_v158
			_ = _lhs103
			var _lhs104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_806_v158), 0))
			_ = _lhs104
			_lhs102.F1 = _rhs161
			(_lhs103).ArraySet1(_rhs162, (_lhs104).Int())
			_809_v159 = _rhs163
			r1 = _rhs164
			var _815_v165 _dafny.Array
			_ = _815_v165
			var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw139
			_815_v165 = _nw139
			_815_v165 = _815_v165
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_815_v165), 0))
			_ = _index135
			(_815_v165).ArraySet1(((_dafny.Zero).Minus((_766_v126).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_766_v126).Cardinality()))).Uint32()).(_dafny.Int))).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(316))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg39 _dafny.Int) interface{} {
					return coer39(arg39)
				}
			}(func(_816_i22 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality())
			}))).Cardinality())), (_index135).Int())
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(535), _dafny.ArrayLen((_815_v165), 0))
			_ = _index136
			(_815_v165).ArraySet1(p1, (_index136).Int())
		} else {
			_766_v126 = _766_v126
			var _817_v166 _dafny.MultiSet
			_ = _817_v166
			_817_v166 = _dafny.MultiSetOf(p0, p0, false, p0, (p0) || (p0))
			_817_v166 = _817_v166
			var _818_v167 D10
			_ = _818_v167
			_818_v167 = Companion_D10_.Create_DC24_(_dafny.CodePoint('y'))
			var _819_v168 *C1
			_ = _819_v168
			var _nw140 *C1 = New_C1_()
			_ = _nw140
			_nw140.Ctor__((_818_v167).Dtor_cf51(), p2, (_this).F6())
			_819_v168 = _nw140
			r1 = !(_dafny.Companion_Sequence_.Equal(_795_v152, _795_v152)) || (p0)
			var _820_v169 *C1
			_ = _820_v169
			var _nw141 *C1 = New_C1_()
			_ = _nw141
			_nw141.Ctor__((_819_v168).F13(), Companion_Default___.Fm1(p0, globalState), (_dafny.IntOfInt64(-480)).Plus(p1))
			_820_v169 = _nw141
		}
		r0 = p0
		var _821_v170 _dafny.MultiSet
		_ = _821_v170
		_821_v170 = _dafny.MultiSetOf(p0)
		r1 = ((func() _dafny.MultiSet {
			if p0 {
				return Companion_Default___.Fm33(globalState)
			}
			return _821_v170
		})()).Contains((func() bool {
			if !(p0) {
				return p0
			}
			return p0
		})())
		var _822_v171 _dafny.MultiSet
		_ = _822_v171
		_822_v171 = _dafny.MultiSetOf((_this).F6())
		r2 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), p0)).Cardinality(), _this.F7()), (_dafny.Zero).Minus((func() _dafny.Int {
			if (_822_v171).Contains(p2) {
				return (_822_v171).Multiplicity(p2)
			}
			return (_822_v171).Cardinality()
		})()))
		return r0, r1, r2
	}
}
func (_this *C2) M8(p0 _dafny.Map, p1 D2, p2 _dafny.Sequence, globalState *GlobalState) {
	{
		var _823_v0 bool
		_ = _823_v0
		_823_v0 = true
		var _rhs165 bool = _823_v0
		_ = _rhs165
		var _rhs166 _dafny.Int = _this.F7()
		_ = _rhs166
		var _lhs105 *GlobalState = globalState
		_ = _lhs105
		_823_v0 = _rhs165
		_lhs105.F1 = _rhs166
		var _824_v1 _dafny.Sequence
		_ = _824_v1
		_824_v1 = _dafny.SeqOf(p2)
		var _825_v2 _dafny.MultiSet
		_ = _825_v2
		_825_v2 = _dafny.MultiSetOf(_dafny.IntOfUint32((_824_v1).Cardinality()))
		var _826_v3 _dafny.MultiSet
		_ = _826_v3
		_826_v3 = _dafny.MultiSetOf((_825_v2).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_823_v0, _dafny.SeqOf((_this).F6(), _this.F7()))).Cardinality())
		var _827_v4 _dafny.Array
		_ = _827_v4
		var _nwElement0_30 _dafny.Int = _dafny.IntOfUint32((p2).Cardinality())
		_ = _nwElement0_30
		var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(6))
		_ = _nw142
		(_nw142).ArraySet1(_nwElement0_30, 0)
		(_nw142).ArraySet1((_this).F6(), 1)
		(_nw142).ArraySet1((func() _dafny.Int {
			if (_826_v3).Contains(_dafny.IntOfInt64(683)) {
				return (_826_v3).Multiplicity(_dafny.IntOfInt64(683))
			}
			return (_this).F6()
		})(), 2)
		(_nw142).ArraySet1(_this.F7(), 3)
		(_nw142).ArraySet1(_this.F7(), 4)
		(_nw142).ArraySet1((_this).F6(), 5)
		_827_v4 = _nw142
		var _828_v5 _dafny.Set
		_ = _828_v5
		_828_v5 = _dafny.SetOf(_827_v4, _827_v4, _827_v4, _827_v4, _827_v4)
		var _829_v6 _dafny.Set
		_ = _829_v6
		_829_v6 = _dafny.SetOf((_dafny.Zero).Minus((_828_v5).Cardinality()), (_this).F6(), _this.F7())
		_829_v6 = func() _dafny.Set {
			var _coll24 = _dafny.NewBuilder()
			_ = _coll24
			for _iter26 := _dafny.Iterate(((_829_v6).Difference(_829_v6)).Elements()); ; {
				_compr_24, _ok26 := _iter26()
				if !_ok26 {
					break
				}
				var _830_v7 _dafny.Int
				_830_v7 = interface{}(_compr_24).(_dafny.Int)
				if ((_829_v6).Difference(_829_v6)).Contains(_830_v7) {
					_coll24.Add(Companion_Default___.SafeModuloInt(_830_v7, _dafny.IntOfInt64(264)))
				}
			}
			return _coll24.ToSet()
		}()
		if !(_823_v0) {
			_823_v0 = _823_v0
			var _831_v8 _dafny.Map
			_ = _831_v8
			_831_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(556), (_this).F6())).Cardinality()).Cmp((_this).F6()) != 0)
			var _832_v9 _dafny.Map
			_ = _832_v9
			_832_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), (_this).F6())
			_823_v0 = (func() bool {
				if (_831_v8).Contains((_this).F6()) {
					return (_831_v8).Get((_this).F6()).(bool)
				}
				return Companion_Default___.Fm0(_823_v0, _832_v9, _dafny.IntOfInt64(472), globalState)
			})()
			var _833_v10 _dafny.CodePoint
			_ = _833_v10
			_833_v10 = _dafny.CodePoint('g')
			var _834_v11 *C1
			_ = _834_v11
			var _nw143 *C1 = New_C1_()
			_ = _nw143
			_nw143.Ctor__(_833_v10, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-990), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_this.F7()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((p2).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_this.F7())).Cardinality()))).Uint32(), (_dafny.SetOf(_dafny.IntOfInt64(151), (_this).F6())).Cardinality())).Cardinality())), (func() _dafny.Int {
				if _823_v0 {
					return _dafny.IntOfInt64(-72)
				}
				return _dafny.IntOfUint32((p2).Cardinality())
			})())
			_834_v11 = _nw143
			(_this).F7_set_(_this.F7())
			var _835_v12 _dafny.MultiSet
			_ = _835_v12
			_835_v12 = _dafny.MultiSetOf(_823_v0)
			var _836_v13 *C1
			_ = _836_v13
			var _nw144 *C1 = New_C1_()
			_ = _nw144
			_nw144.Ctor__((_834_v11).F13(), _this.F7(), Companion_Default___.SafeDivisionInt((_835_v12).Cardinality(), _this.F7()))
			_836_v13 = _nw144
		} else {
			var _837_v14 _dafny.Sequence
			_ = _837_v14
			_837_v14 = _dafny.UnicodeSeqOfUtf8Bytes("jx")
			var _838_v15 _dafny.Array
			_ = _838_v15
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(16))
			_ = _nw145
			_838_v15 = _nw145
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_838_v15), 0))
			_ = _index137
			(_838_v15).ArraySet1((_825_v2).Difference(_825_v2), (_index137).Int())
			var _839_v16 _dafny.Map
			_ = _839_v16
			_839_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _823_v0)
			var _840_v19 _dafny.Sequence
			_ = _840_v19
			_840_v19 = _dafny.SeqOf(_this.F7())
			var _841_v20 _dafny.Array
			_ = _841_v20
			var _nwElement0_31 _dafny.Map = _839_v16
			_ = _nwElement0_31
			var _nw146 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(5))
			_ = _nw146
			(_nw146).ArraySet1(_nwElement0_31, 0)
			(_nw146).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _823_v0), 1)
			(_nw146).ArraySet1(func() _dafny.Map {
				var _coll25 = _dafny.NewMapBuilder()
				_ = _coll25
				for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(856), _dafny.IntOfInt64(645))); ; {
					_compr_25, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _842_v17 _dafny.Int
					_842_v17 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(856)).Cmp(_842_v17) <= 0) && ((_842_v17).Cmp(_dafny.IntOfInt64(645)) < 0) {
						_coll25.Add((_842_v17).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xrpgkhh")).Cardinality())), _823_v0)
					}
				}
				return _coll25.ToMap()
			}(), 2)
			(_nw146).ArraySet1((_839_v16).Merge(func() _dafny.Map {
				var _coll26 = _dafny.NewMapBuilder()
				_ = _coll26
				for _iter28 := _dafny.Iterate((_840_v19).Elements()); ; {
					_compr_26, _ok28 := _iter28()
					if !_ok28 {
						break
					}
					var _843_v18 _dafny.Int
					_843_v18 = interface{}(_compr_26).(_dafny.Int)
					if _dafny.Companion_Sequence_.Contains(_840_v19, _843_v18) {
						_coll26.Add(Companion_Default___.SafeModuloInt(_843_v18, (_this).F6()), _823_v0)
					}
				}
				return _coll26.ToMap()
			}()), 3)
			(_nw146).ArraySet1(_839_v16, 4)
			_841_v20 = _nw146
			var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_838_v15), 0))
			_ = _index138
			var _rhs167 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(223))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}(func(_844_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})), p2), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(983))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}(func(_845_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('u')
			})))
			_ = _rhs167
			var _rhs168 _dafny.MultiSet = _dafny.MultiSetFromSeq(_840_v19)
			_ = _rhs168
			var _rhs169 bool = _823_v0
			_ = _rhs169
			var _rhs170 bool = !(((_this).F6()).Cmp((_this.F7()).Times((_this).F6())) != 0)
			_ = _rhs170
			var _rhs171 _dafny.Array = _841_v20
			_ = _rhs171
			var _lhs106 _dafny.Array = _838_v15
			_ = _lhs106
			var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_838_v15), 0))
			_ = _lhs107
			_837_v14 = _rhs167
			(_lhs106).ArraySet1(_rhs168, (_lhs107).Int())
			_823_v0 = _rhs169
			_823_v0 = _rhs170
			_841_v20 = _rhs171
			var _846_v22 _dafny.Array
			_ = _846_v22
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_27
			var _nw147 _dafny.Array
			_ = _nw147
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw147 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Set = (func(_847_v19 _dafny.Sequence) func(_dafny.Int) _dafny.Set {
					return func(_848_i2 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf(_dafny.Companion_Sequence_.Update(_847_v19, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_847_v19).Cardinality()))).Uint32(), (_this).F6()))).Union(func() _dafny.Set {
							var _coll27 = _dafny.NewBuilder()
							_ = _coll27
							for _iter29 := _dafny.Iterate((_dafny.SeqOf(_847_v19, _847_v19)).Elements()); ; {
								_compr_27, _ok29 := _iter29()
								if !_ok29 {
									break
								}
								var _849_v21 _dafny.Sequence
								_849_v21 = interface{}(_compr_27).(_dafny.Sequence)
								if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_847_v19, _847_v19), _849_v21) {
									_coll27.Add(_849_v21)
								}
							}
							return _coll27.ToSet()
						}())
					}
				})(_840_v19)
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw147 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw147).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw147).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_846_v22 = _nw147
			var _850_v23 _dafny.Set
			_ = _850_v23
			_850_v23 = _dafny.SetOf(_840_v19, _840_v19)
			var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_846_v22), 0))
			_ = _index139
			(_846_v22).ArraySet1(_850_v23, (_index139).Int())
			var _851_v24 _dafny.Map
			_ = _851_v24
			_851_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_840_v19, false)
			var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_846_v22), 0))
			_ = _index140
			(_846_v22).ArraySet1((func() _dafny.Set {
				if _823_v0 {
					return func() _dafny.Set {
						var _coll28 = _dafny.NewBuilder()
						_ = _coll28
						for _iter30 := _dafny.Iterate(((_851_v24).Update(_840_v19, !(false))).Keys().Elements()); ; {
							_compr_28, _ok30 := _iter30()
							if !_ok30 {
								break
							}
							var _852_v25 _dafny.Sequence
							_852_v25 = interface{}(_compr_28).(_dafny.Sequence)
							if ((_851_v24).Update(_840_v19, !(false))).Contains(_852_v25) {
								_coll28.Add(_852_v25)
							}
						}
						return _coll28.ToSet()
					}()
				}
				return _850_v23
			})(), (_index140).Int())
			if !(_823_v0) {
				(globalState).F1 = _this.F7()
				var _853_v26 *C0
				_ = _853_v26
				var _nw148 *C0 = New_C0_()
				_ = _nw148
				_nw148.Ctor__(_823_v0)
				_853_v26 = _nw148
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_827_v4), 0))
				_ = _index141
				(_827_v4).ArraySet1((_dafny.Zero).Minus(_this.F7()), (_index141).Int())
				var _854_v27 _dafny.Sequence
				_ = _854_v27
				_854_v27 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_826_v3).Cardinality()))
				var _855_v29 _dafny.CodePoint
				_ = _855_v29
				_855_v29 = _dafny.CodePoint('f')
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_827_v4), 0))
				_ = _index142
				var _rhs172 _dafny.Int = _this.F7()
				_ = _rhs172
				var _rhs173 _dafny.Int = (_dafny.Zero).Minus((((_854_v27).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_854_v27).Cardinality()))).Uint32()).(_dafny.Map)).Merge(func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter31 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg42 _dafny.Int) interface{} {
							return coer42(arg42)
						}
					}(func(_856_i3 _dafny.Int) _dafny.Int {
						return (_this).F6()
					}))).Elements()); ; {
						_compr_29, _ok31 := _iter31()
						if !_ok31 {
							break
						}
						var _857_v28 _dafny.Int
						_857_v28 = interface{}(_compr_29).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(284))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg43 _dafny.Int) interface{} {
								return coer43(arg43)
							}
						}(func(_856_i3 _dafny.Int) _dafny.Int {
							return (_this).F6()
						})), _857_v28) {
							_coll29.Add((_857_v28).Times(_this.F7()), _this.F7())
						}
					}
					return _coll29.ToMap()
				}())).Cardinality())
				_ = _rhs173
				var _rhs174 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_837_v14, (func() _dafny.Sequence {
					if _853_v26.F10 {
						return _dafny.Companion_Sequence_.Update(_837_v14, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_837_v14).Cardinality()))).Uint32(), _855_v29)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(950))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}((func(_858_v29 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_859_i4 _dafny.Int) _dafny.CodePoint {
							return _858_v29
						}
					})(_855_v29)))
				})())).Cardinality())
				_ = _rhs174
				var _rhs175 _dafny.Sequence = p2
				_ = _rhs175
				var _lhs108 *GlobalState = globalState
				_ = _lhs108
				var _lhs109 *GlobalState = globalState
				_ = _lhs109
				var _lhs110 _dafny.Array = _827_v4
				_ = _lhs110
				var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_827_v4), 0))
				_ = _lhs111
				_lhs108.F1 = _rhs172
				_lhs109.F1 = _rhs173
				(_lhs110).ArraySet1(_rhs174, (_lhs111).Int())
				_837_v14 = _rhs175
				var _860_v30 *C0
				_ = _860_v30
				var _nw149 *C0 = New_C0_()
				_ = _nw149
				_nw149.Ctor__(_823_v0)
				_860_v30 = _nw149
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(666), _dafny.ArrayLen((_827_v4), 0))
				_ = _index143
				(_827_v4).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(222), (_this.F7()).Minus((_this).F6())), (_index143).Int())
			} else {
				var _861_v31 _dafny.Map
				_ = _861_v31
				_861_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-201))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}(func(_862_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('p')
				}))).Cardinality()))
				var _rhs176 bool = true
				_ = _rhs176
				var _rhs177 bool = Companion_Default___.Fm0(_823_v0, _861_v31, Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(734), (_this).F6()), globalState)
				_ = _rhs177
				_823_v0 = _rhs176
				_823_v0 = _rhs177
				(globalState).F1 = Companion_Default___.SafeDivisionInt(_this.F7(), _this.F7())
				(globalState).F1 = _dafny.IntOfInt64(767)
				(_this).M9(_823_v0, _this.F7(), _this.F7(), (_this).F6(), globalState)
				(globalState).F1 = _dafny.IntOfUint32((p2).Cardinality())
			}
			var _863_v32 _dafny.Array
			_ = _863_v32
			var _len0_28 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_28
			var _nw150 _dafny.Array
			_ = _nw150
			if _len0_28.Cmp(_dafny.Zero) == 0 {
				_nw150 = _dafny.NewArray(_len0_28)
			} else {
				var _init28 func(_dafny.Int) bool = (func(_864_v0 bool) func(_dafny.Int) bool {
					return func(_865_i6 _dafny.Int) bool {
						return _864_v0
					}
				})(_823_v0)
				_ = _init28
				var _element0_28 = _init28(_dafny.Zero)
				_ = _element0_28
				_nw150 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
				(_nw150).ArraySet1(_element0_28, 0)
				var _nativeLen0_28 = (_len0_28).Int()
				_ = _nativeLen0_28
				for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
					(_nw150).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
				}
			}
			_863_v32 = _nw150
			var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_863_v32), 0))
			_ = _index144
			(_863_v32).ArraySet1(!(_823_v0) || (_823_v0), (_index144).Int())
			var _866_v33 D10
			_ = _866_v33
			_866_v33 = Companion_D10_.Create_DC26_(_823_v0)
			var _pat_let_tv11 = _823_v0
			_ = _pat_let_tv11
			var _867_v34 _dafny.Map
			_ = _867_v34
			_867_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let19_0 D10) D10 {
				return func(_868_dt__update__tmp_h0 D10) D10 {
					return func(_pat_let20_0 bool) D10 {
						return func(_869_dt__update_hcf52_h0 bool) D10 {
							return Companion_D10_.Create_DC26_(_869_dt__update_hcf52_h0)
						}(_pat_let20_0)
					}(_pat_let_tv11)
				}(_pat_let19_0)
			}(_866_v33), _823_v0)
			var _870_v35 _dafny.Map
			_ = _870_v35
			_870_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_867_v34).Cardinality(), (_this).F6())
			var _871_v36 _dafny.Set
			_ = _871_v36
			_871_v36 = _dafny.SetOf(_823_v0, Companion_Default___.Fm0(_823_v0, _870_v35, _this.F7(), globalState))
			var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_863_v32), 0))
			_ = _index145
			(_863_v32).ArraySet1((_871_v36).IsSubsetOf(_871_v36), (_index145).Int())
			var _872_v37 _dafny.Map
			_ = _872_v37
			_872_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_863_v32).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_863_v32), 0))).Int()).(bool), _dafny.IntOfUint32((p2).Cardinality()))
			(globalState).F1 = ((func() _dafny.Int {
				if (_872_v37).Contains(false) {
					return (_872_v37).Get(false).(_dafny.Int)
				}
				return _this.F7()
			})()).Minus(_this.F7())
		}
		var _873_v39 _dafny.Map
		_ = _873_v39
		_873_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(-142))
		var _874_v40 _dafny.Set
		_ = _874_v40
		_874_v40 = _dafny.SetOf(_823_v0, _823_v0, Companion_Default___.Fm0(_823_v0, _873_v39, _dafny.IntOfInt64(685), globalState))
		var _875_v41 _dafny.Set
		_ = _875_v41
		_875_v41 = _dafny.SetOf(_829_v6)
		var _876_v42 _dafny.Array
		_ = _876_v42
		var _nwElement0_32 bool = _823_v0
		_ = _nwElement0_32
		var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(24))
		_ = _nw151
		(_nw151).ArraySet1(_nwElement0_32, 0)
		(_nw151).ArraySet1((_823_v0) == (_823_v0), 1)
		(_nw151).ArraySet1(_823_v0, 2)
		(_nw151).ArraySet1(_823_v0, 3)
		(_nw151).ArraySet1((func() bool {
			if !(_823_v0) {
				return _823_v0
			}
			return _823_v0
		})(), 4)
		(_nw151).ArraySet1(!((_dafny.SetOf(_this.F7(), _this.F7(), (_this).F6(), _this.F7())).Contains((_this).F6())), 5)
		(_nw151).ArraySet1(_823_v0, 6)
		(_nw151).ArraySet1(true, 7)
		(_nw151).ArraySet1(!(_823_v0), 8)
		(_nw151).ArraySet1(_823_v0, 9)
		(_nw151).ArraySet1((_825_v2).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(85))), 10)
		(_nw151).ArraySet1((_dafny.IntOfInt64(779)).Cmp((func() _dafny.Map {
			var _coll30 = _dafny.NewMapBuilder()
			_ = _coll30
			for _iter32 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_829_v6).Cardinality(), _825_v2)).Keys().Elements()); ; {
				_compr_30, _ok32 := _iter32()
				if !_ok32 {
					break
				}
				var _877_v38 _dafny.Int
				_877_v38 = interface{}(_compr_30).(_dafny.Int)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_829_v6).Cardinality(), _825_v2)).Contains(_877_v38) {
					_coll30.Add(Companion_Default___.SafeDivisionInt(_877_v38, (_829_v6).Cardinality()), _823_v0)
				}
			}
			return _coll30.ToMap()
		}()).Cardinality()) == 0, 11)
		(_nw151).ArraySet1(_823_v0, 12)
		(_nw151).ArraySet1(false, 13)
		(_nw151).ArraySet1(!_dafny.Companion_Sequence_.Equal(p2, p2), 14)
		(_nw151).ArraySet1(!(_874_v40).Contains(_823_v0), 15)
		(_nw151).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(p2, p2), 16)
		(_nw151).ArraySet1(_823_v0, 17)
		(_nw151).ArraySet1(_823_v0, 18)
		(_nw151).ArraySet1(_823_v0, 19)
		(_nw151).ArraySet1(!(!(!((_875_v41).IsDisjointFrom(Companion_Default___.Fm34(globalState))))), 20)
		(_nw151).ArraySet1(_823_v0, 21)
		(_nw151).ArraySet1(_823_v0, 22)
		(_nw151).ArraySet1(_823_v0, 23)
		_876_v42 = _nw151
		for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_876_v42), 0))); ; {
			_guard_loop_2, _ok33 := _iter33()
			if !_ok33 {
				break
			}
			var _878_i7 _dafny.Int
			_878_i7 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_878_i7).Sign() != -1) && ((_878_i7).Cmp(_dafny.ArrayLen((_876_v42), 0)) < 0)) {
				(_876_v42).ArraySet1(((_this).F6()).Cmp(_this.F7()) > 0, (_878_i7).Int())
			}
		}
		_823_v0 = _823_v0
		var _879_v43 _dafny.CodePoint
		_ = _879_v43
		_879_v43 = _dafny.CodePoint('k')
		var _880_v44 *C1
		_ = _880_v44
		var _nw152 *C1 = New_C1_()
		_ = _nw152
		_nw152.Ctor__(_879_v43, _this.F7(), (_dafny.Zero).Minus((_this).F6()))
		_880_v44 = _nw152
	}
}
func (_this *C2) M9(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) {
	{
		var _881_v0 bool
		_ = _881_v0
		_881_v0 = true
		var _882_v1 _dafny.Sequence
		_ = _882_v1
		_882_v1 = _dafny.UnicodeSeqOfUtf8Bytes("wdfh")
		_881_v0 = (Companion_Default___.Fm0(p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_882_v1).Cardinality()), p2), _this.F7(), globalState)) || (p0)
		var _883_v2 _dafny.CodePoint
		_ = _883_v2
		_883_v2 = _dafny.CodePoint('a')
		var _884_v3 T0
		_ = _884_v3
		var _nw153 *C1 = New_C1_()
		_ = _nw153
		_nw153.Ctor__(_883_v2, _this.F7(), p2)
		_884_v3 = _nw153
		var _rhs178 _dafny.Int = (_884_v3).F6()
		_ = _rhs178
		var _rhs179 T0 = _884_v3
		_ = _rhs179
		var _lhs112 *C2 = _this
		_ = _lhs112
		_lhs112.F7_set_(_rhs178)
		_884_v3 = _rhs179
		(_884_v3).F7_set_(_884_v3.F7())
		if !(!(p0)) {
			var _885_v4 _dafny.Sequence
			_ = _885_v4
			_885_v4 = _dafny.SeqOf(_this.F7())
			(_884_v3).F7_set_(Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_885_v4).Cardinality()), p1), Companion_Default___.Fm1(_881_v0, globalState)))
			(globalState).F1 = (_dafny.Zero).Minus(p2)
			var _886_v5 _dafny.Sequence
			_ = _886_v5
			_886_v5 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tjotuokyi"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-105))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}((func(_887_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_888_i0 _dafny.Int) _dafny.CodePoint {
					return _887_v2
				}
			})(_883_v2))))
			(_884_v3).F7_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_886_v5, _886_v5), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(135))).Uint32(), func(coer47 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}((func(_889_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_890_i1 _dafny.Int) _dafny.Sequence {
					return _889_v1
				}
			})(_882_v1))))).Cardinality()))
			var _891_v6 D10
			_ = _891_v6
			_891_v6 = Companion_D10_.Create_DC24_(_883_v2)
			var _892_v7 _dafny.Array
			_ = _892_v7
			var _nwElement0_33 _dafny.CodePoint = _883_v2
			_ = _nwElement0_33
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(7))
			_ = _nw154
			(_nw154).ArraySet1CodePoint(_nwElement0_33, 0)
			(_nw154).ArraySet1CodePoint(_883_v2, 1)
			(_nw154).ArraySet1CodePoint(_883_v2, 2)
			(_nw154).ArraySet1CodePoint(_dafny.CodePoint('t'), 3)
			(_nw154).ArraySet1CodePoint(_883_v2, 4)
			(_nw154).ArraySet1CodePoint(_883_v2, 5)
			(_nw154).ArraySet1CodePoint((_891_v6).Dtor_cf51(), 6)
			_892_v7 = _nw154
			_892_v7 = _892_v7
			var _893_v8 _dafny.Sequence
			_ = _893_v8
			_893_v8 = _dafny.SeqOf(_881_v0)
			var _894_v9 D5
			_ = _894_v9
			_894_v9 = Companion_D5_.Create_DC11_(_893_v8)
			(_884_v3).F7_set_(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_894_v9).Dtor_cf20(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(353), _dafny.IntOfUint32(((_894_v9).Dtor_cf20()).Cardinality()))).Uint32(), p0)).Cardinality()))
		} else {
			var _895_v10 _dafny.Map
			_ = _895_v10
			_895_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v0, p2)
			var _896_v11 _dafny.Array
			_ = _896_v11
			var _nwElement0_34 _dafny.Map = _895_v10
			_ = _nwElement0_34
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(16))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_34, 0)
			(_nw155).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_884_v3).F6()), 1)
			(_nw155).ArraySet1(_895_v10, 2)
			(_nw155).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p3)).Merge(_895_v10), 3)
			(_nw155).ArraySet1(_895_v10, 4)
			(_nw155).ArraySet1((_895_v10).Update(!(p0), p1), 5)
			(_nw155).ArraySet1((_895_v10).Merge(Companion_Default___.Fm30(p0, false, _884_v3.F7(), p2, globalState)), 6)
			(_nw155).ArraySet1((Companion_Default___.Fm30(p0, p0, _dafny.IntOfInt64(-807), p2, globalState)).Merge(_895_v10), 7)
			(_nw155).ArraySet1((_895_v10).Merge(_895_v10), 8)
			(_nw155).ArraySet1(_895_v10, 9)
			(_nw155).ArraySet1((_895_v10).Merge(_895_v10), 10)
			(_nw155).ArraySet1((_895_v10).Update(_881_v0, _dafny.IntOfInt64(834)), 11)
			(_nw155).ArraySet1((_895_v10).Merge(_895_v10), 12)
			(_nw155).ArraySet1(_895_v10, 13)
			(_nw155).ArraySet1(_895_v10, 14)
			(_nw155).ArraySet1(((_895_v10).Update(p0, p3)).Update(p0, _dafny.IntOfUint32((Companion_Default___.Fm35((_884_v3).F6(), _884_v3.F7(), globalState)).Cardinality())), 15)
			_896_v11 = _nw155
			var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_896_v11), 0))
			_ = _index146
			(_896_v11).ArraySet1((_895_v10).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_881_v0, _884_v3.F7())), (_index146).Int())
			var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_896_v11), 0))
			_ = _index147
			(_896_v11).ArraySet1(Companion_Default___.Fm30(p0, p0, (p1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_884_v3.F7())).Cardinality())), (_dafny.Zero).Minus((_dafny.IntOfInt64(-943)).Times(_884_v3.F7())), globalState), (_index147).Int())
			var _897_v12 _dafny.Sequence
			_ = _897_v12
			_897_v12 = _dafny.SeqOf(_this.F7())
			var _898_v13 *C0
			_ = _898_v13
			var _nw156 *C0 = New_C0_()
			_ = _nw156
			_nw156.Ctor__(p0)
			_898_v13 = _nw156
			var _899_v14 _dafny.MultiSet
			_ = _899_v14
			_899_v14 = _dafny.MultiSetOf(_898_v13)
			if (_884_v3.F7()).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((Companion_Default___.Fm2(_dafny.Companion_Sequence_.Update(_897_v12, (Companion_Default___.SafeIndex((_899_v14).Cardinality(), _dafny.IntOfUint32((_897_v12).Cardinality()))).Uint32(), p3), _dafny.IntOfInt64(986), _dafny.IntOfInt64(475), p0, globalState)).Cardinality()), p2)) != 0 {
				var _900_v15 _dafny.Set
				_ = _900_v15
				_900_v15 = _dafny.SetOf(p0)
				var _901_v16 _dafny.Map
				_ = _901_v16
				_901_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_900_v15).IsSubsetOf(_dafny.SetOf(_898_v13.F10)))
				_901_v16 = _901_v16
				var _902_v17 D10
				_ = _902_v17
				_902_v17 = Companion_D10_.Create_DC24_(_883_v2)
				var _903_v18 _dafny.Array
				_ = _903_v18
				var _nwElement0_35 _dafny.CodePoint = _883_v2
				_ = _nwElement0_35
				var _nw157 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(22))
				_ = _nw157
				(_nw157).ArraySet1CodePoint(_nwElement0_35, 0)
				(_nw157).ArraySet1CodePoint(_883_v2, 1)
				(_nw157).ArraySet1CodePoint(_883_v2, 2)
				(_nw157).ArraySet1CodePoint(Companion_Default___.Fm25(globalState), 3)
				(_nw157).ArraySet1CodePoint(_883_v2, 4)
				(_nw157).ArraySet1CodePoint(_883_v2, 5)
				(_nw157).ArraySet1CodePoint((_902_v17).Dtor_cf51(), 6)
				(_nw157).ArraySet1CodePoint(_883_v2, 7)
				(_nw157).ArraySet1CodePoint(_883_v2, 8)
				(_nw157).ArraySet1CodePoint(_883_v2, 9)
				(_nw157).ArraySet1CodePoint(_dafny.CodePoint('p'), 10)
				(_nw157).ArraySet1CodePoint(_883_v2, 11)
				(_nw157).ArraySet1CodePoint(_883_v2, 12)
				(_nw157).ArraySet1CodePoint(Companion_Default___.Fm25(globalState), 13)
				(_nw157).ArraySet1CodePoint(_883_v2, 14)
				(_nw157).ArraySet1CodePoint(_883_v2, 15)
				(_nw157).ArraySet1CodePoint(_883_v2, 16)
				(_nw157).ArraySet1CodePoint(_883_v2, 17)
				(_nw157).ArraySet1CodePoint(_883_v2, 18)
				(_nw157).ArraySet1CodePoint(_883_v2, 19)
				(_nw157).ArraySet1CodePoint(_883_v2, 20)
				(_nw157).ArraySet1CodePoint(_883_v2, 21)
				_903_v18 = _nw157
				var _904_v19 _dafny.Sequence
				_ = _904_v19
				_904_v19 = _dafny.SeqOf(_898_v13, _898_v13, _898_v13)
				var _905_v20 *C0
				_ = _905_v20
				_905_v20 = (_904_v19).Select((Companion_Default___.SafeIndex(_884_v3.F7(), _dafny.IntOfUint32((_904_v19).Cardinality()))).Uint32()).(*C0)
				var _906_v21 _dafny.Map
				_ = _906_v21
				_906_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v18, _905_v20)
				(_884_v3).F7_set_((_906_v21).Cardinality())
				var _907_v22 _dafny.Map
				_ = _907_v22
				_907_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_903_v18, p2)
				_907_v22 = (_907_v22).Update(_903_v18, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hh")).Cardinality()))
				_900_v15 = (_900_v15).Intersection(_900_v15)
				var _908_v23 _dafny.Sequence
				_ = _908_v23
				_908_v23 = _dafny.SeqOf(_898_v13.F10)
				var _909_v24 D10
				_ = _909_v24
				_909_v24 = Companion_D10_.Create_DC26_(_898_v13.F10)
				var _910_v25 _dafny.Array
				_ = _910_v25
				var _nwElement0_36 _dafny.Sequence = _908_v23
				_ = _nwElement0_36
				var _nw158 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(24))
				_ = _nw158
				(_nw158).ArraySet1(_nwElement0_36, 0)
				(_nw158).ArraySet1(_908_v23, 1)
				(_nw158).ArraySet1(_908_v23, 2)
				(_nw158).ArraySet1(_908_v23, 3)
				(_nw158).ArraySet1(_908_v23, 4)
				(_nw158).ArraySet1(_908_v23, 5)
				(_nw158).ArraySet1(_908_v23, 6)
				(_nw158).ArraySet1(Companion_Default___.Fm27(p2, _882_v1, _dafny.CodePoint('b'), globalState), 7)
				(_nw158).ArraySet1((func() _dafny.Sequence {
					if _898_v13.F10 {
						return _908_v23
					}
					return _908_v23
				})(), 8)
				(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_908_v23, _908_v23), 9)
				(_nw158).ArraySet1(_908_v23, 10)
				(_nw158).ArraySet1(_908_v23, 11)
				(_nw158).ArraySet1(_908_v23, 12)
				(_nw158).ArraySet1((func() _dafny.Sequence {
					if _898_v13.F10 {
						return _908_v23
					}
					return _908_v23
				})(), 13)
				(_nw158).ArraySet1(_dafny.SeqOf(_898_v13.F10, _898_v13.F10, _898_v13.F10, _898_v13.F10, (_908_v23).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_908_v23).Cardinality()))).Uint32()).(bool)), 14)
				(_nw158).ArraySet1(_dafny.SeqOf((_909_v24).Dtor_cf52(), true), 15)
				(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_908_v23, _dafny.SeqOf(_898_v13.F10)), 16)
				(_nw158).ArraySet1(_dafny.SeqOf(_881_v0, _898_v13.F10), 17)
				(_nw158).ArraySet1(_908_v23, 18)
				(_nw158).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_908_v23, _908_v23), 19)
				(_nw158).ArraySet1(_dafny.SeqOf(true), 20)
				(_nw158).ArraySet1(_908_v23, 21)
				(_nw158).ArraySet1(_dafny.SeqOf(p0, _881_v0, _881_v0), 22)
				(_nw158).ArraySet1(_908_v23, 23)
				_910_v25 = _nw158
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_910_v25), 0))
				_ = _index148
				(_910_v25).ArraySet1(_908_v23, (_index148).Int())
				var _911_v26 _dafny.Map
				_ = _911_v26
				_911_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_898_v13.F10, _908_v23)
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(520), _dafny.ArrayLen((_910_v25), 0))
				_ = _index149
				(_910_v25).ArraySet1((func() _dafny.Sequence {
					if (_911_v26).Contains((func() bool {
						if _898_v13.F10 {
							return _898_v13.F10
						}
						return _881_v0
					})()) {
						return (_911_v26).Get((func() bool {
							if _898_v13.F10 {
								return _898_v13.F10
							}
							return _881_v0
						})()).(_dafny.Sequence)
					}
					return _908_v23
				})(), (_index149).Int())
			} else {
				var _912_v27 _dafny.Array
				_ = _912_v27
				var _nw159 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw159
				_912_v27 = _nw159
				var _913_v28 _dafny.Array
				_ = _913_v28
				var _nw160 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
				_ = _nw160
				_913_v28 = _nw160
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))
				_ = _index150
				(_913_v28).ArraySet1(_898_v13.F10, (_index150).Int())
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))
				_ = _index151
				var _rhs180 _dafny.Array = _912_v27
				_ = _rhs180
				var _rhs181 _dafny.Int = _this.F7()
				_ = _rhs181
				var _rhs182 *C0 = _898_v13
				_ = _rhs182
				var _rhs183 bool = p0
				_ = _rhs183
				var _lhs113 *GlobalState = globalState
				_ = _lhs113
				var _lhs114 _dafny.Array = _913_v28
				_ = _lhs114
				var _lhs115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))
				_ = _lhs115
				_912_v27 = _rhs180
				_lhs113.F1 = _rhs181
				_898_v13 = _rhs182
				(_lhs114).ArraySet1(_rhs183, (_lhs115).Int())
				var _914_v29 _dafny.MultiSet
				_ = _914_v29
				_914_v29 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer48 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_915_v0 bool) func(_dafny.Int) D0 {
					return func(_916_i2 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_dafny.SetOf(_915_v0))
					}
				})(_881_v0))))
				var _917_v30 _dafny.Map
				_ = _917_v30
				_917_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_884_v3).F6(), _898_v13.F10)
				var _918_v31 D1
				_ = _918_v31
				_918_v31 = Companion_D1_.Create_DC4_(_898_v13.F10)
				var _919_v32 D2
				_ = _919_v32
				_919_v32 = Companion_D2_.Create_DC7_((_913_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))).Int()).(bool), _dafny.IntOfInt64(257), _917_v30, _918_v31)
				var _920_v33 _dafny.Set
				_ = _920_v33
				_920_v33 = _dafny.SetOf(!((_913_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))).Int()).(bool)), (_919_v32).Dtor_cf13())
				var _921_v34 D0
				_ = _921_v34
				_921_v34 = Companion_D0_.Create_DC0_(_920_v33)
				var _922_v35 _dafny.Sequence
				_ = _922_v35
				_922_v35 = _dafny.SeqOf(_921_v34)
				var _923_v36 _dafny.Sequence
				_ = _923_v36
				_923_v36 = _dafny.SeqOf(p0)
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))
				_ = _index152
				var _rhs184 bool = (Companion_D10_.Create_DC26_((_913_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))).Int()).(bool))).Dtor_cf52()
				_ = _rhs184
				var _rhs185 _dafny.Int = (func() _dafny.Int {
					if (_914_v29).Contains(_922_v35) {
						return (_914_v29).Multiplicity(_922_v35)
					}
					return _884_v3.F7()
				})()
				_ = _rhs185
				var _rhs186 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_923_v36, _923_v36)
				_ = _rhs186
				var _lhs116 _dafny.Array = _913_v28
				_ = _lhs116
				var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))
				_ = _lhs117
				var _lhs118 T0 = _884_v3
				_ = _lhs118
				(_lhs116).ArraySet1(_rhs184, (_lhs117).Int())
				_lhs118.F7_set_(_rhs185)
				_881_v0 = _rhs186
				(_898_v13).F10 = (_913_v28).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(18), _dafny.ArrayLen((_913_v28), 0))).Int()).(bool)
				(_898_v13).F10 = !(!(!(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("oqk"), _883_v2))))
				var _924_v37 _dafny.MultiSet
				_ = _924_v37
				_924_v37 = _dafny.MultiSetOf(((_884_v3).F6()).Times(_884_v3.F7()))
				(_884_v3).F7_set_((func() _dafny.Int {
					if (_924_v37).Contains((_884_v3).F6()) {
						return (_924_v37).Multiplicity((_884_v3).F6())
					}
					return (_dafny.IntOfInt64(288)).Minus(_884_v3.F7())
				})())
			}
			var _925_v38 _dafny.Array
			_ = _925_v38
			var _nw161 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw161
			_925_v38 = _nw161
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_925_v38), 0))
			_ = _index153
			(_925_v38).ArraySet1(_884_v3.F7(), (_index153).Int())
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_925_v38), 0))
			_ = _index154
			var _rhs187 _dafny.Int = _this.F7()
			_ = _rhs187
			var _rhs188 _dafny.Int = Companion_Default___.SafeDivisionInt((_this.F7()).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-376))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_926_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_927_i3 _dafny.Int) _dafny.CodePoint {
					return _926_v2
				}
			})(_883_v2)))).Cardinality())), p3)
			_ = _rhs188
			var _lhs119 T0 = _884_v3
			_ = _lhs119
			var _lhs120 _dafny.Array = _925_v38
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_925_v38), 0))
			_ = _lhs121
			_lhs119.F7_set_(_rhs187)
			(_lhs120).ArraySet1(_rhs188, (_lhs121).Int())
			var _928_v39 D2
			_ = _928_v39
			_928_v39 = Companion_D2_.Create_DC5_(_925_v38)
			var _pat_let_tv12 = _925_v38
			_ = _pat_let_tv12
			_928_v39 = func(_pat_let21_0 D2) D2 {
				return func(_929_dt__update__tmp_h0 D2) D2 {
					return func(_pat_let22_0 _dafny.Array) D2 {
						return func(_930_dt__update_hcf9_h0 _dafny.Array) D2 {
							return Companion_D2_.Create_DC5_(_930_dt__update_hcf9_h0)
						}(_pat_let22_0)
					}(_pat_let_tv12)
				}(_pat_let21_0)
			}(_928_v39)
			var _931_v40 _dafny.Sequence
			_ = _931_v40
			_931_v40 = _dafny.SeqOf(_881_v0)
			var _932_v41 D9
			_ = _932_v41
			_932_v41 = Companion_D9_.Create_DC23_((_931_v40).Select((Companion_Default___.SafeIndex((_884_v3).F6(), _dafny.IntOfUint32((_931_v40).Cardinality()))).Uint32()).(bool), _dafny.IntOfInt64(20), _881_v0, _881_v0, _884_v3.F7())
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(902), _dafny.ArrayLen((_925_v38), 0))
			_ = _index155
			(_925_v38).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_932_v41).Dtor_cf47())), (_index155).Int())
		}
		var _933_v42 _dafny.Array
		_ = _933_v42
		var _len0_29 _dafny.Int = _dafny.IntOfInt64(28)
		_ = _len0_29
		var _nw162 _dafny.Array
		_ = _nw162
		if _len0_29.Cmp(_dafny.Zero) == 0 {
			_nw162 = _dafny.NewArray(_len0_29)
		} else {
			var _init29 func(_dafny.Int) bool = (func(_934_v0 bool) func(_dafny.Int) bool {
				return func(_935_i4 _dafny.Int) bool {
					return !(_934_v0)
				}
			})(_881_v0)
			_ = _init29
			var _element0_29 = _init29(_dafny.Zero)
			_ = _element0_29
			_nw162 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
			(_nw162).ArraySet1(_element0_29, 0)
			var _nativeLen0_29 = (_len0_29).Int()
			_ = _nativeLen0_29
			for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
				(_nw162).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
			}
		}
		_933_v42 = _nw162
		_933_v42 = _933_v42
		var _936_v43 _dafny.Map
		_ = _936_v43
		_936_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _881_v0)).Cardinality(), p3)).Cardinality()), _881_v0)
		var _937_v44 _dafny.Map
		_ = _937_v44
		_937_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_936_v43).Contains(p3) {
				return (_936_v43).Get(p3).(bool)
			}
			return true
		})(), p2)
		var _938_i5 _dafny.Int
		_ = _938_i5
		_938_i5 = _dafny.Zero
		{
			var _pat_let_tv13 = _881_v0
			_ = _pat_let_tv13
			var _pat_let_tv14 = _881_v0
			_ = _pat_let_tv14
			var _pat_let_tv15 = _881_v0
			_ = _pat_let_tv15
			var _pat_let_tv16 = p0
			_ = _pat_let_tv16
			for func(_source9 D5) bool {
				if _source9.Is_DC12() {
					var _944___mcc_h0 _dafny.Int = _source9.Get_().(D5_DC12).Cf21
					_ = _944___mcc_h0
					var _945_cf21 _dafny.Int = _944___mcc_h0
					_ = _945_cf21
					if _pat_let_tv13 {
						return !(_pat_let_tv14)
					} else {
						return _pat_let_tv15
					}
				} else {
					var _946___mcc_h1 _dafny.Sequence = _source9.Get_().(D5_DC11).Cf20
					_ = _946___mcc_h1
					var _947_cf20 _dafny.Sequence = _946___mcc_h1
					_ = _947_cf20
					return _pat_let_tv16
				}
			}(Companion_D5_.Create_DC12_((func() _dafny.Int {
				if (_937_v44).Contains(p0) {
					return (_937_v44).Get(p0).(_dafny.Int)
				}
				return p3
			})())) {
				{
					if (_938_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_938_i5 = (_938_i5).Plus(_dafny.One)
					var _939_v45 _dafny.Map
					_ = _939_v45
					_939_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(706))
					var _940_v46 _dafny.Set
					_ = _940_v46
					_940_v46 = _dafny.SetOf(_884_v3.F7(), _884_v3.F7())
					var _941_v47 _dafny.Set
					_ = _941_v47
					_941_v47 = _dafny.SetOf(p0)
					_939_v45 = (_939_v45).Update(_dafny.IntOfInt64(-226), Companion_Default___.SafeModuloInt((_940_v46).Cardinality(), (_941_v47).Cardinality()))
					_882_v1 = _882_v1
					var _942_v49 _dafny.MultiSet
					_ = _942_v49
					_942_v49 = _dafny.MultiSetOf(func() _dafny.Map {
						var _coll31 = _dafny.NewMapBuilder()
						_ = _coll31
						for _iter34 := _dafny.Iterate((_936_v43).Keys().Elements()); ; {
							_compr_31, _ok34 := _iter34()
							if !_ok34 {
								break
							}
							var _943_v48 _dafny.Int
							_943_v48 = interface{}(_compr_31).(_dafny.Int)
							if (_936_v43).Contains(_943_v48) {
								_coll31.Add(Companion_Default___.SafeModuloInt(_943_v48, _dafny.IntOfInt64(983)), p3)
							}
						}
						return _coll31.ToMap()
					}(), _939_v45, _939_v45)
					_942_v49 = _942_v49
					_881_v0 = _881_v0
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f7 _dafny.Int
	_f6 _dafny.Int
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f7 = _dafny.Zero
	_this._f6 = _dafny.Zero
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

func (_this *C3) F7() _dafny.Int {
	return _this._f7
}
func (_this *C3) F7_set_(value _dafny.Int) {
	_this._f7 = value
}
func (_this *C3) F6() _dafny.Int {
	return _this._f6
}
func (_this *C3) Ctor__(f6 _dafny.Int, f7 _dafny.Int) {
	{
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C3) Fm15(p0 _dafny.Int, p1 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-928)
	}
}
func (_this *C3) Fm16(globalState *GlobalState) bool {
	{
		return !(!(((_this).F6()).Cmp(_this.F7()) < 0))
	}
}
func (_this *C3) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _948_v0 _dafny.Set
		_ = _948_v0
		_948_v0 = _dafny.SetOf(p0)
		var _949_v1 _dafny.Map
		_ = _949_v1
		_949_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _950_v2 _dafny.Sequence
		_ = _950_v2
		_950_v2 = _dafny.UnicodeSeqOfUtf8Bytes("gqysp")
		var _951_v3 _dafny.Sequence
		_ = _951_v3
		_951_v3 = _dafny.SeqOf(_949_v1, _949_v1, (_949_v1).Update(p0, _dafny.IntOfUint32((_950_v2).Cardinality())), _949_v1, _949_v1)
		var _952_v4 _dafny.Sequence
		_ = _952_v4
		_952_v4 = _dafny.SeqOf(p0)
		var _953_v5 _dafny.Map
		_ = _953_v5
		_953_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), ((_948_v0).Cardinality()).Plus(((_951_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_952_v4).Cardinality()), _dafny.IntOfUint32((_951_v3).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()))
		var _954_v6 _dafny.CodePoint
		_ = _954_v6
		_954_v6 = _dafny.CodePoint('a')
		var _955_v7 _dafny.Sequence
		_ = _955_v7
		_955_v7 = _dafny.SeqOf(_950_v2, _950_v2)
		var _956_v8 _dafny.Sequence
		_ = _956_v8
		_956_v8 = _dafny.SeqOf(p1, _this.F7(), _this.F7())
		var _957_v9 _dafny.Sequence
		_ = _957_v9
		_957_v9 = _dafny.SeqOf(_956_v8)
		var _958_v10 _dafny.Map
		_ = _958_v10
		_958_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
		var _959_v11 D1
		_ = _959_v11
		_959_v11 = Companion_D1_.Create_DC4_(p0)
		var _960_v12 D2
		_ = _960_v12
		_960_v12 = Companion_D2_.Create_DC7_(p0, _this.F7(), _958_v10, _959_v11)
		var _pat_let_tv17 = p0
		_ = _pat_let_tv17
		var _pat_let_tv18 = p0
		_ = _pat_let_tv18
		var _pat_let_tv19 = p0
		_ = _pat_let_tv19
		var _pat_let_tv20 = p0
		_ = _pat_let_tv20
		var _pat_let_tv21 = _958_v10
		_ = _pat_let_tv21
		var _rhs189 _dafny.Map = Companion_Default___.Fm17(_dafny.MultiSetFromSeq((_955_v7).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_956_v8).Cardinality()), _dafny.IntOfUint32((_955_v7).Cardinality()))).Uint32()).(_dafny.Sequence)), !(p0) || (p0), _this.F7(), _957_v9, globalState)
		_ = _rhs189
		var _rhs190 bool = func(_source10 D2) bool {
			if _source10.Is_DC6() {
				var _961___mcc_h0 D1 = _source10.Get_().(D2_DC6).Cf10
				_ = _961___mcc_h0
				var _962___mcc_h1 _dafny.Sequence = _source10.Get_().(D2_DC6).Cf11
				_ = _962___mcc_h1
				var _963___mcc_h2 _dafny.Int = _source10.Get_().(D2_DC6).Cf12
				_ = _963___mcc_h2
				var _964_cf12 _dafny.Int = _963___mcc_h2
				_ = _964_cf12
				var _965_cf11 _dafny.Sequence = _962___mcc_h1
				_ = _965_cf11
				var _966_cf10 D1 = _961___mcc_h0
				_ = _966_cf10
				return _pat_let_tv17
			} else if _source10.Is_DC7() {
				var _967___mcc_h3 bool = _source10.Get_().(D2_DC7).Cf13
				_ = _967___mcc_h3
				var _968___mcc_h4 _dafny.Int = _source10.Get_().(D2_DC7).Cf14
				_ = _968___mcc_h4
				var _969___mcc_h5 _dafny.Map = _source10.Get_().(D2_DC7).Cf15
				_ = _969___mcc_h5
				var _970___mcc_h6 D1 = _source10.Get_().(D2_DC7).Cf16
				_ = _970___mcc_h6
				var _971_cf16 D1 = _970___mcc_h6
				_ = _971_cf16
				var _972_cf15 _dafny.Map = _969___mcc_h5
				_ = _972_cf15
				var _973_cf14 _dafny.Int = _968___mcc_h4
				_ = _973_cf14
				var _974_cf13 bool = _967___mcc_h3
				_ = _974_cf13
				return true
			} else if _source10.Is_DC5() {
				var _975___mcc_h7 _dafny.Array = _source10.Get_().(D2_DC5).Cf9
				_ = _975___mcc_h7
				var _976_cf9 _dafny.Array = _975___mcc_h7
				_ = _976_cf9
				return _pat_let_tv18
			} else {
				var _977___mcc_h8 D2 = _source10.Get_().(D2_DC8).Cf17
				_ = _977___mcc_h8
				var _978_cf17 D2 = _977___mcc_h8
				_ = _978_cf17
				return !(_pat_let_tv19) || (_pat_let_tv20)
			}
		}(func(_pat_let23_0 D2) D2 {
			return func(_979_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let24_0 _dafny.Map) D2 {
					return func(_980_dt__update_hcf15_h0 _dafny.Map) D2 {
						return Companion_D2_.Create_DC7_((_979_dt__update__tmp_h0).Dtor_cf13(), (_979_dt__update__tmp_h0).Dtor_cf14(), _980_dt__update_hcf15_h0, (_979_dt__update__tmp_h0).Dtor_cf16())
					}(_pat_let24_0)
				}(_pat_let_tv21)
			}(_pat_let23_0)
		}(_960_v12))
		_ = _rhs190
		var _rhs191 _dafny.CodePoint = _954_v6
		_ = _rhs191
		var _rhs192 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(306), (func() _dafny.Int {
			if p0 {
				return _this.F7()
			}
			return (Companion_Default___.Fm18(Companion_Default___.Fm0(p0, _953_v5, (_this).F6(), globalState), p1, globalState)).Cardinality()
		})())
		_ = _rhs192
		var _rhs193 bool = p0
		_ = _rhs193
		var _lhs122 *C3 = _this
		_ = _lhs122
		_953_v5 = _rhs189
		r1 = _rhs190
		_954_v6 = _rhs191
		_lhs122.F7_set_(_rhs192)
		r1 = _rhs193
		var _981_i0 _dafny.Int
		_ = _981_i0
		_981_i0 = _dafny.Zero
		{
			for true {
				{
					if (_981_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_981_i0 = (_981_i0).Plus(_dafny.One)
					var _982_v13 _dafny.Int
					_ = _982_v13
					var _983_v14 _dafny.Int
					_ = _983_v14
					var _984_v15 _dafny.Int
					_ = _984_v15
					var _out15 _dafny.Int
					_ = _out15
					var _out16 _dafny.Int
					_ = _out16
					var _out17 _dafny.Int
					_ = _out17
					_out15, _out16, _out17 = (_this).M7(p0, (func() bool {
						if (_958_v10).Contains(p1) {
							return (_958_v10).Get(p1).(bool)
						}
						return p0
					})(), p2, globalState)
					_982_v13 = _out15
					_983_v14 = _out16
					_984_v15 = _out17
					var _985_v16 _dafny.Array
					_ = _985_v16
					var _nw163 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
					_ = _nw163
					_985_v16 = _nw163
					var _986_v17 _dafny.Map
					_ = _986_v17
					_986_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_985_v16), 0))
					_ = _index156
					(_985_v16).ArraySet1(((_986_v17).Merge(_986_v17)).Cardinality(), (_index156).Int())
					var _987_v18 _dafny.MultiSet
					_ = _987_v18
					_987_v18 = _dafny.MultiSetOf(p0)
					var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_985_v16), 0))
					_ = _index157
					(_985_v16).ArraySet1((func() _dafny.Int {
						if (_987_v18).Contains(p0) {
							return (_987_v18).Multiplicity(p0)
						}
						return (_this).F6()
					})(), (_index157).Int())
					var _988_v19 _dafny.MultiSet
					_ = _988_v19
					_988_v19 = _dafny.MultiSetOf(Companion_Default___.Fm1(p0, globalState), ((_dafny.Zero).Minus((_dafny.Zero).Minus(p1))).Minus(_dafny.IntOfInt64(-411)), (_this.F7()).Minus(p1))
					var _989_v20 _dafny.MultiSet
					_ = _989_v20
					_989_v20 = _987_v18
					var _rhs194 bool = (p0) && (p0)
					_ = _rhs194
					var _rhs195 _dafny.MultiSet = _988_v19
					_ = _rhs195
					var _rhs196 _dafny.MultiSet = (func() _dafny.MultiSet {
						if !(!(p0)) {
							return Companion_Default___.Fm19(p0, globalState)
						}
						return (_987_v18).Update(true, Companion_Default___.Abs(p2))
					})()
					_ = _rhs196
					var _rhs197 _dafny.Int = _dafny.IntOfInt64(217)
					_ = _rhs197
					var _lhs123 *GlobalState = globalState
					_ = _lhs123
					r0 = _rhs194
					_988_v19 = _rhs195
					_989_v20 = _rhs196
					_lhs123.F1 = _rhs197
					r1 = p0
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		_954_v6 = _954_v6
		var _990_i1 _dafny.Int
		_ = _990_i1
		_990_i1 = _dafny.Zero
		{
			for false {
				{
					if (_990_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_990_i1 = (_990_i1).Plus(_dafny.One)
					var _991_v21 _dafny.Array
					_ = _991_v21
					var _nw164 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(13))
					_ = _nw164
					_991_v21 = _nw164
					var _992_v22 D0
					_ = _992_v22
					_992_v22 = Companion_D0_.Create_DC1_(p0, _991_v21, p0)
					var _993_v23 _dafny.Array
					_ = _993_v23
					var _nwElement0_37 D0 = _992_v22
					_ = _nwElement0_37
					var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.One)
					_ = _nw165
					(_nw165).ArraySet1(_nwElement0_37, 0)
					_993_v23 = _nw165
					var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_993_v23), 0))
					_ = _index158
					(_993_v23).ArraySet1(_992_v22, (_index158).Int())
					var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_993_v23), 0))
					_ = _index159
					(_993_v23).ArraySet1(Companion_D0_.Create_DC1_(p0, _991_v21, p0), (_index159).Int())
					var _994_v24 _dafny.Map
					_ = _994_v24
					_994_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), _954_v6)
					var _995_v25 _dafny.Map
					_ = _995_v25
					_995_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(_994_v24).Contains(p0)), p0)
					_995_v25 = (_995_v25).Update(p0, p0)
					if ((_dafny.Zero).Minus(p1)).Cmp(p2) > 0 {
						var _996_v26 _dafny.Array
						_ = _996_v26
						var _len0_30 _dafny.Int = _dafny.IntOfInt64(17)
						_ = _len0_30
						var _nw166 _dafny.Array
						_ = _nw166
						if _len0_30.Cmp(_dafny.Zero) == 0 {
							_nw166 = _dafny.NewArray(_len0_30)
						} else {
							var _init30 func(_dafny.Int) _dafny.Int = func(_997_i2 _dafny.Int) _dafny.Int {
								return (_997_i2).Plus((_this).F6())
							}
							_ = _init30
							var _element0_30 = _init30(_dafny.Zero)
							_ = _element0_30
							_nw166 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
							(_nw166).ArraySet1(_element0_30, 0)
							var _nativeLen0_30 = (_len0_30).Int()
							_ = _nativeLen0_30
							for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
								(_nw166).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
							}
						}
						_996_v26 = _nw166
						var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_996_v26), 0))
						_ = _index160
						(_996_v26).ArraySet1(_this.F7(), (_index160).Int())
						var _998_v27 _dafny.Set
						_ = _998_v27
						_998_v27 = _dafny.SetOf(p1, _dafny.IntOfInt64(792), p1)
						var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_996_v26), 0))
						_ = _index161
						var _rhs198 bool = p0
						_ = _rhs198
						var _rhs199 _dafny.Int = p2
						_ = _rhs199
						var _rhs200 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_956_v8, _956_v8), _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm2(_956_v8, (_998_v27).Cardinality(), _this.F7(), p0, globalState)).Cardinality())))).Cardinality()))
						_ = _rhs200
						var _rhs201 _dafny.Sequence = _950_v2
						_ = _rhs201
						var _rhs202 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_this.F7()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_956_v8, _956_v8)).Cardinality()))
						_ = _rhs202
						var _lhs124 *C3 = _this
						_ = _lhs124
						var _lhs125 *C3 = _this
						_ = _lhs125
						var _lhs126 _dafny.Array = _996_v26
						_ = _lhs126
						var _lhs127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(372), _dafny.ArrayLen((_996_v26), 0))
						_ = _lhs127
						r1 = _rhs198
						_lhs124.F7_set_(_rhs199)
						_lhs125.F7_set_(_rhs200)
						_950_v2 = _rhs201
						(_lhs126).ArraySet1(_rhs202, (_lhs127).Int())
						r1 = !(p0)
						_950_v2 = _950_v2
						var _999_v28 *C0
						_ = _999_v28
						var _nw167 *C0 = New_C0_()
						_ = _nw167
						_nw167.Ctor__(true)
						_999_v28 = _nw167
						var _1000_v29 _dafny.MultiSet
						_ = _1000_v29
						_1000_v29 = _dafny.MultiSetOf(_dafny.MultiSetOf(p0, p0, p0))
						var _1001_v30 _dafny.MultiSet
						_ = _1001_v30
						_1001_v30 = _dafny.MultiSetOf(p0, p0)
						var _1002_v31 _dafny.MultiSet
						_ = _1002_v31
						_1002_v31 = _1001_v30
						_1000_v29 = (_1000_v29).Union((_1000_v29).Union(_dafny.MultiSetFromSeq(_dafny.SeqOf(_1002_v31))))
					} else {
						var _1003_v32 _dafny.Set
						_ = _1003_v32
						_1003_v32 = _dafny.SetOf(p2)
						var _1004_v33 _dafny.Map
						_ = _1004_v33
						_1004_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1003_v32, _this.F7())
						_1004_v33 = (_1004_v33).Update(_1003_v32, p2)
						var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))
						_ = _index162
						(_991_v21).ArraySet1(((_958_v10).Cardinality()).Cmp(p2) != 0, (_index162).Int())
						var _1005_v34 _dafny.Array
						_ = _1005_v34
						var _len0_31 _dafny.Int = _dafny.IntOfInt64(6)
						_ = _len0_31
						var _nw168 _dafny.Array
						_ = _nw168
						if _len0_31.Cmp(_dafny.Zero) == 0 {
							_nw168 = _dafny.NewArray(_len0_31)
						} else {
							var _init31 func(_dafny.Int) _dafny.Sequence = (func(_1006_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
								return func(_1007_i3 _dafny.Int) _dafny.Sequence {
									return _1006_v4
								}
							})(_952_v4)
							_ = _init31
							var _element0_31 = _init31(_dafny.Zero)
							_ = _element0_31
							_nw168 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
							(_nw168).ArraySet1(_element0_31, 0)
							var _nativeLen0_31 = (_len0_31).Int()
							_ = _nativeLen0_31
							for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
								(_nw168).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
							}
						}
						_1005_v34 = _nw168
						var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))
						_ = _index163
						var _rhs203 _dafny.Int = _dafny.IntOfInt64(-542)
						_ = _rhs203
						var _rhs204 bool = p0
						_ = _rhs204
						var _rhs205 _dafny.Array = _1005_v34
						_ = _rhs205
						var _lhs128 *GlobalState = globalState
						_ = _lhs128
						var _lhs129 _dafny.Array = _991_v21
						_ = _lhs129
						var _lhs130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))
						_ = _lhs130
						_lhs128.F1 = _rhs203
						(_lhs129).ArraySet1(_rhs204, (_lhs130).Int())
						_1005_v34 = _rhs205
						var _1008_v35 *C0
						_ = _1008_v35
						var _nw169 *C0 = New_C0_()
						_ = _nw169
						_nw169.Ctor__(!(p0) || ((_991_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))).Int()).(bool)))
						_1008_v35 = _nw169
						var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))
						_ = _index164
						var _rhs206 bool = p0
						_ = _rhs206
						var _rhs207 bool = _1008_v35.F10
						_ = _rhs207
						var _rhs208 bool = Companion_Default___.Fm0((_this).Fm16(globalState), _953_v5, (_this).Fm15(p1, _dafny.IntOfInt64(285), globalState), globalState)
						_ = _rhs208
						var _rhs209 bool = _1008_v35.F10
						_ = _rhs209
						var _rhs210 _dafny.Int = (_this).Fm15(p2, p1, globalState)
						_ = _rhs210
						var _lhs131 _dafny.Array = _991_v21
						_ = _lhs131
						var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))
						_ = _lhs132
						var _lhs133 *C0 = _1008_v35
						_ = _lhs133
						var _lhs134 *C0 = _1008_v35
						_ = _lhs134
						var _lhs135 *GlobalState = globalState
						_ = _lhs135
						r1 = _rhs206
						(_lhs131).ArraySet1(_rhs207, (_lhs132).Int())
						_lhs133.F10 = _rhs208
						_lhs134.F10 = _rhs209
						_lhs135.F1 = _rhs210
						var _1009_v36 D0
						_ = _1009_v36
						_1009_v36 = Companion_D0_.Create_DC0_(_948_v0)
						var _1010_v37 _dafny.Map
						_ = _1010_v37
						_1010_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1009_v36, (_991_v21).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_991_v21), 0))).Int()).(bool))
						var _1011_v38 D1
						_ = _1011_v38
						_1011_v38 = Companion_D1_.Create_DC2_(_1010_v37)
						var _1012_v39 _dafny.Set
						_ = _1012_v39
						_1012_v39 = _dafny.SetOf(_1011_v38, _1011_v38, Companion_D1_.Create_DC2_(_1010_v37), _1011_v38)
						var _1013_v40 bool
						_ = _1013_v40
						var _out18 bool
						_ = _out18
						_out18 = (_this).M6((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_949_v1).Cardinality(), _1008_v35.F10)).Merge(_958_v10), p1, (_1012_v39).Union(_1012_v39), globalState)
						_1013_v40 = _out18
					}
					_959_v11 = _959_v11
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		if (_dafny.IntOfUint32((_950_v2).Cardinality())).Cmp(p1) != 0 {
			var _1014_v41 *C0
			_ = _1014_v41
			var _nw170 *C0 = New_C0_()
			_ = _nw170
			_nw170.Ctor__(p0)
			_1014_v41 = _nw170
			_954_v6 = Companion_Default___.Fm20(_950_v2, _1014_v41.F10, _this.F7(), globalState)
			var _1015_v42 _dafny.MultiSet
			_ = _1015_v42
			_1015_v42 = _dafny.MultiSetOf(_958_v10, _958_v10)
			r1 = !((_1015_v42).Equals(_1015_v42))
			var _1016_v43 _dafny.Array
			_ = _1016_v43
			var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(24))
			_ = _nw171
			_1016_v43 = _nw171
			var _1017_v44 _dafny.MultiSet
			_ = _1017_v44
			_1017_v44 = _dafny.MultiSetOf(p2, _dafny.IntOfUint32((_950_v2).Cardinality()), _dafny.IntOfInt64(556), p1, (_958_v10).Cardinality())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_1016_v43), 0))
			_ = _index165
			(_1016_v43).ArraySet1((_1017_v44).Difference(_dafny.MultiSetOf(_dafny.IntOfUint32((_950_v2).Cardinality()), (_this).Fm15((_this).F6(), _dafny.IntOfInt64(156), globalState), (_1017_v44).Cardinality())), (_index165).Int())
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_1016_v43), 0))
			_ = _index166
			(_1016_v43).ArraySet1(_dafny.MultiSetFromSeq(_956_v8), (_index166).Int())
			var _1018_v45 _dafny.Array
			_ = _1018_v45
			var _nw172 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw172
			_1018_v45 = _nw172
			var _1019_v46 _dafny.MultiSet
			_ = _1019_v46
			_1019_v46 = _dafny.MultiSetOf(_1018_v45)
			var _1020_v47 _dafny.Map
			_ = _1020_v47
			_1020_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.UnicodeSeqOfUtf8Bytes("mubemdm"))
			var _1021_v48 bool
			_ = _1021_v48
			var _out19 bool
			_ = _out19
			_out19 = (_this).M6((_958_v10).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_1019_v46).Update(_1018_v45, Companion_Default___.Abs((_1020_v47).Cardinality()))).Cardinality(), _1014_v41.F10)), p2, _dafny.SetOf(Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_948_v0), _1014_v41.F10))), globalState)
			_1021_v48 = _out19
		} else {
			var _1022_v49 _dafny.Array
			_ = _1022_v49
			var _nwElement0_38 _dafny.CodePoint = _954_v6
			_ = _nwElement0_38
			var _nw173 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(21))
			_ = _nw173
			(_nw173).ArraySet1CodePoint(_nwElement0_38, 0)
			(_nw173).ArraySet1CodePoint(_dafny.CodePoint('l'), 1)
			(_nw173).ArraySet1CodePoint(_954_v6, 2)
			(_nw173).ArraySet1CodePoint(_954_v6, 3)
			(_nw173).ArraySet1CodePoint(Companion_Default___.Fm20(_950_v2, p0, (_958_v10).Cardinality(), globalState), 4)
			(_nw173).ArraySet1CodePoint(_dafny.CodePoint('c'), 5)
			(_nw173).ArraySet1CodePoint(_954_v6, 6)
			(_nw173).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _954_v6
				}
				return _954_v6
			})(), 7)
			(_nw173).ArraySet1CodePoint(_954_v6, 8)
			(_nw173).ArraySet1CodePoint((_950_v2).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_950_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
			(_nw173).ArraySet1CodePoint(_954_v6, 10)
			(_nw173).ArraySet1CodePoint(_954_v6, 11)
			(_nw173).ArraySet1CodePoint(_dafny.CodePoint('o'), 12)
			(_nw173).ArraySet1CodePoint(_954_v6, 13)
			(_nw173).ArraySet1CodePoint(_954_v6, 14)
			(_nw173).ArraySet1CodePoint(_954_v6, 15)
			(_nw173).ArraySet1CodePoint(_dafny.CodePoint('u'), 16)
			(_nw173).ArraySet1CodePoint(_954_v6, 17)
			(_nw173).ArraySet1CodePoint(Companion_Default___.Fm20(_dafny.UnicodeSeqOfUtf8Bytes("kpvixhs"), p0, p2, globalState), 18)
			(_nw173).ArraySet1CodePoint(_954_v6, 19)
			(_nw173).ArraySet1CodePoint((func() _dafny.CodePoint {
				if p0 {
					return _954_v6
				}
				return _954_v6
			})(), 20)
			_1022_v49 = _nw173
			_1022_v49 = _1022_v49
			var _1023_v50 *C0
			_ = _1023_v50
			var _nw174 *C0 = New_C0_()
			_ = _nw174
			_nw174.Ctor__(p0)
			_1023_v50 = _nw174
			var _1024_v51 *C0
			_ = _1024_v51
			_1024_v51 = (func() *C0 {
				if p0 {
					return _1023_v50
				}
				return _1023_v50
			})()
			var _source11 *C0 = _1024_v51
			_ = _source11
			var _1025___mcc_h9 *C0 = _source11
			_ = _1025___mcc_h9
			var _1026_cf19 *C0 = _1025___mcc_h9
			_ = _1026_cf19
			var _1027_v52 _dafny.Set
			_ = _1027_v52
			_1027_v52 = _dafny.SetOf(p1)
			var _1028_v53 _dafny.Array
			_ = _1028_v53
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(21))
			_ = _nw175
			_1028_v53 = _nw175
			var _1029_v54 _dafny.Map
			_ = _1029_v54
			_1029_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1028_v53, (_953_v5).Cardinality())
			var _1030_v55 _dafny.Array
			_ = _1030_v55
			var _nwElement0_39 _dafny.Int = (_1027_v52).Cardinality()
			_ = _nwElement0_39
			var _nw176 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(14))
			_ = _nw176
			(_nw176).ArraySet1(_nwElement0_39, 0)
			(_nw176).ArraySet1(_this.F7(), 1)
			(_nw176).ArraySet1(p2, 2)
			(_nw176).ArraySet1((_dafny.IntOfInt64(278)).Plus(p2), 3)
			(_nw176).ArraySet1(p1, 4)
			(_nw176).ArraySet1(((_this).F6()).Minus(p1), 5)
			(_nw176).ArraySet1(_this.F7(), 6)
			(_nw176).ArraySet1((_dafny.Zero).Minus(p1), 7)
			(_nw176).ArraySet1((_956_v8).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(p1), _dafny.IntOfUint32((_956_v8).Cardinality()))).Uint32()).(_dafny.Int), 8)
			(_nw176).ArraySet1(((_this).F6()).Plus(p2), 9)
			(_nw176).ArraySet1((_this).F6(), 10)
			(_nw176).ArraySet1(((_1029_v54).Merge(_1029_v54)).Cardinality(), 11)
			(_nw176).ArraySet1(_this.F7(), 12)
			(_nw176).ArraySet1(((_this).F6()).Plus(p2), 13)
			_1030_v55 = _nw176
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1028_v53), 0))
			_ = _index167
			(_1028_v53).ArraySet1((_this).Fm15(p2, _this.F7(), globalState), (_index167).Int())
			var _1031_v57 _dafny.Array
			_ = _1031_v57
			var _nwElement0_40 bool = _1023_v50.F10
			_ = _nwElement0_40
			var _nw177 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(11))
			_ = _nw177
			(_nw177).ArraySet1(_nwElement0_40, 0)
			(_nw177).ArraySet1((_this).Fm16(globalState), 1)
			(_nw177).ArraySet1(_1023_v50.F10, 2)
			(_nw177).ArraySet1(!(_1026_cf19.F10), 3)
			(_nw177).ArraySet1((func() bool {
				if _1023_v50.F10 {
					return _1023_v50.F10
				}
				return p0
			})(), 4)
			(_nw177).ArraySet1(!(_1023_v50.F10) || (_1026_cf19.F10), 5)
			(_nw177).ArraySet1((p0) == (p0), 6)
			(_nw177).ArraySet1(!(_1023_v50.F10) || (Companion_Default___.Fm0(_1026_cf19.F10, func() _dafny.Map {
				var _coll32 = _dafny.NewMapBuilder()
				_ = _coll32
				for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(283), _dafny.IntOfInt64(-466))); ; {
					_compr_32, _ok35 := _iter35()
					if !_ok35 {
						break
					}
					var _1032_v56 _dafny.Int
					_1032_v56 = interface{}(_compr_32).(_dafny.Int)
					if ((_dafny.IntOfInt64(283)).Cmp(_1032_v56) <= 0) && ((_1032_v56).Cmp(_dafny.IntOfInt64(-466)) < 0) {
						_coll32.Add(Companion_Default___.SafeModuloInt(_1032_v56, _this.F7()), (_948_v0).Cardinality())
					}
				}
				return _coll32.ToMap()
			}(), (_this).F6(), globalState)), 7)
			(_nw177).ArraySet1(_1026_cf19.F10, 8)
			(_nw177).ArraySet1(_1026_cf19.F10, 9)
			(_nw177).ArraySet1(!(_1026_cf19.F10), 10)
			_1031_v57 = _nw177
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1031_v57), 0))
			_ = _index168
			(_1031_v57).ArraySet1(_1026_cf19.F10, (_index168).Int())
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1028_v53), 0))
			_ = _index169
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1031_v57), 0))
			_ = _index170
			var _rhs211 _dafny.Int = p2
			_ = _rhs211
			var _rhs212 bool = _1023_v50.F10
			_ = _rhs212
			var _rhs213 _dafny.Int = (p2).Minus((p2).Plus(p1))
			_ = _rhs213
			var _rhs214 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_1030_v55, _1030_v55, _1028_v53, _1030_v55)).Cardinality()), p2)
			_ = _rhs214
			var _lhs136 _dafny.Array = _1028_v53
			_ = _lhs136
			var _lhs137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(144), _dafny.ArrayLen((_1028_v53), 0))
			_ = _lhs137
			var _lhs138 _dafny.Array = _1031_v57
			_ = _lhs138
			var _lhs139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1031_v57), 0))
			_ = _lhs139
			var _lhs140 *GlobalState = globalState
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			(_lhs136).ArraySet1(_rhs211, (_lhs137).Int())
			(_lhs138).ArraySet1(_rhs212, (_lhs139).Int())
			_lhs140.F1 = _rhs213
			_lhs141.F1 = _rhs214
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1031_v57), 0))
			_ = _index171
			(_1031_v57).ArraySet1((false) && (Companion_Default___.Fm0((_1031_v57).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(701), _dafny.ArrayLen((_1031_v57), 0))).Int()).(bool), _953_v5, _this.F7(), globalState)), (_index171).Int())
			_1023_v50 = _1026_cf19
			var _1033_v58 *C0
			_ = _1033_v58
			var _nw178 *C0 = New_C0_()
			_ = _nw178
			_nw178.Ctor__(!_dafny.Companion_Sequence_.Equal(_950_v2, _950_v2))
			_1033_v58 = _nw178
			var _1034_v59 D0
			_ = _1034_v59
			_1034_v59 = Companion_D0_.Create_DC0_(_948_v0)
			var _1035_v60 _dafny.Map
			_ = _1035_v60
			_1035_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1034_v59, _1023_v50.F10)
			var _1036_v61 _dafny.Map
			_ = _1036_v61
			_1036_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1035_v60)
			var _1037_v62 D1
			_ = _1037_v62
			_1037_v62 = Companion_D1_.Create_DC2_((func() _dafny.Map {
				if (_1036_v61).Contains(_1023_v50.F10) {
					return (_1036_v61).Get(_1023_v50.F10).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_948_v0), false)
			})())
			_1037_v62 = _1037_v62
			var _1038_v63 _dafny.Map
			_ = _1038_v63
			_1038_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1023_v50.F10)
			var _1039_v64 _dafny.Sequence
			_ = _1039_v64
			_1039_v64 = _dafny.SeqOf(_1038_v63)
			var _1040_v65 _dafny.MultiSet
			_ = _1040_v65
			_1040_v65 = _dafny.MultiSetOf(_this.F7(), p1)
			var _1041_v66 _dafny.Array
			_ = _1041_v66
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_32
			var _nw179 _dafny.Array
			_ = _nw179
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw179 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) bool = (func(_1042_v50 *C0) func(_dafny.Int) bool {
					return func(_1043_i4 _dafny.Int) bool {
						return _1042_v50.F10
					}
				})(_1023_v50)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw179 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw179).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw179).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1041_v66 = _nw179
			var _1044_v67 D0
			_ = _1044_v67
			_1044_v67 = Companion_D0_.Create_DC1_(_1023_v50.F10, _1041_v66, _1023_v50.F10)
			var _1045_v68 _dafny.MultiSet
			_ = _1045_v68
			_1045_v68 = _dafny.MultiSetOf(_1023_v50.F10, true, p0)
			var _1046_v69 _dafny.Array
			_ = _1046_v69
			var _nwElement0_41 bool = (func() bool {
				if _1023_v50.F10 {
					return p0
				}
				return true
			})()
			_ = _nwElement0_41
			var _nw180 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(27))
			_ = _nw180
			(_nw180).ArraySet1(_nwElement0_41, 0)
			(_nw180).ArraySet1(p0, 1)
			(_nw180).ArraySet1(!((p1).Cmp((((_1039_v64).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1039_v64).Cardinality()))).Uint32()).(_dafny.Map)).Update(_1023_v50.F10, _1023_v50.F10)).Cardinality()) == 0), 2)
			(_nw180).ArraySet1((_960_v12).Dtor_cf13(), 3)
			(_nw180).ArraySet1(_1023_v50.F10, 4)
			(_nw180).ArraySet1(!(p0), 5)
			(_nw180).ArraySet1(_1023_v50.F10, 6)
			(_nw180).ArraySet1(_1023_v50.F10, 7)
			(_nw180).ArraySet1(false, 8)
			(_nw180).ArraySet1(_1023_v50.F10, 9)
			(_nw180).ArraySet1((_948_v0).IsSubsetOf(_948_v0), 10)
			(_nw180).ArraySet1(p0, 11)
			(_nw180).ArraySet1(_1023_v50.F10, 12)
			(_nw180).ArraySet1(p0, 13)
			(_nw180).ArraySet1(_1023_v50.F10, 14)
			(_nw180).ArraySet1((_1040_v65).IsProperSubsetOf(_dafny.MultiSetFromSeq(_956_v8)), 15)
			(_nw180).ArraySet1(_1023_v50.F10, 16)
			(_nw180).ArraySet1((p2).Cmp(p2) >= 0, 17)
			(_nw180).ArraySet1(p0, 18)
			(_nw180).ArraySet1((false) || (true), 19)
			(_nw180).ArraySet1(!((func() bool {
				if _1023_v50.F10 {
					return _1023_v50.F10
				}
				return (_1044_v67).Dtor_cf3()
			})()), 20)
			(_nw180).ArraySet1((_1045_v68).Equals(_dafny.MultiSetOf(Companion_Default___.Fm0(_1023_v50.F10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), p1, globalState))), 21)
			(_nw180).ArraySet1(Companion_Default___.Fm0(_1023_v50.F10, _953_v5, p1, globalState), 22)
			(_nw180).ArraySet1(p0, 23)
			(_nw180).ArraySet1(p0, 24)
			(_nw180).ArraySet1((_this).Fm16(globalState), 25)
			(_nw180).ArraySet1(_1023_v50.F10, 26)
			_1046_v69 = _nw180
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_1046_v69), 0))
			_ = _index172
			(_1046_v69).ArraySet1(p0, (_index172).Int())
			var _1047_v70 _dafny.Set
			_ = _1047_v70
			_1047_v70 = _dafny.SetOf((_this).F6())
			var _1048_v72 _dafny.Array
			_ = _1048_v72
			var _nw181 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(16))
			_ = _nw181
			_1048_v72 = _nw181
			var _1049_v73 _dafny.Set
			_ = _1049_v73
			_1049_v73 = _dafny.SetOf(_1048_v72, _1048_v72)
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_1046_v69), 0))
			_ = _index173
			var _rhs215 bool = (func() bool {
				if (_1047_v70).IsDisjointFrom(func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(288), _dafny.IntOfInt64(394))); ; {
						_compr_33, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _1050_v71 _dafny.Int
						_1050_v71 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(288)).Cmp(_1050_v71) <= 0) && ((_1050_v71).Cmp(_dafny.IntOfInt64(394)) < 0) {
							_coll33.Add(Companion_Default___.SafeDivisionInt(_1050_v71, p1))
						}
					}
					return _coll33.ToSet()
				}()) {
					return p0
				}
				return true
			})()
			_ = _rhs215
			var _rhs216 bool = !(((_1049_v73).Union(_1049_v73)).IsSubsetOf((_1049_v73).Intersection(_1049_v73)))
			_ = _rhs216
			var _rhs217 _dafny.Int = (_dafny.Zero).Minus((p2).Minus((_953_v5).Cardinality()))
			_ = _rhs217
			var _lhs142 _dafny.Array = _1046_v69
			_ = _lhs142
			var _lhs143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(401), _dafny.ArrayLen((_1046_v69), 0))
			_ = _lhs143
			var _lhs144 *C3 = _this
			_ = _lhs144
			(_lhs142).ArraySet1(_rhs215, (_lhs143).Int())
			r1 = _rhs216
			_lhs144.F7_set_(_rhs217)
			r0 = p0
		}
		if (_this).Fm16(globalState) {
			_949_v1 = (_949_v1).Update(p0, (_dafny.IntOfUint32((_950_v2).Cardinality())).Times((_dafny.Zero).Minus(p2)))
			var _rhs218 _dafny.Sequence = (func() _dafny.Sequence {
				if true {
					return _952_v4
				}
				return _952_v4
			})()
			_ = _rhs218
			var _rhs219 _dafny.Sequence = _952_v4
			_ = _rhs219
			_952_v4 = _rhs218
			_952_v4 = _rhs219
			r2 = p2
			var _1051_v74 _dafny.Map
			_ = _1051_v74
			_1051_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_954_v6, p2)
			r1 = (_952_v4).Select((Companion_Default___.SafeIndex(((func() _dafny.Int {
				if (_1051_v74).Contains(_954_v6) {
					return (_1051_v74).Get(_954_v6).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_956_v8).Cardinality())
			})()).Plus(p1), _dafny.IntOfUint32((_952_v4).Cardinality()))).Uint32()).(bool)
			var _1052_v75 _dafny.Map
			_ = _1052_v75
			_1052_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1053_v76 _dafny.Set
			_ = _1053_v76
			_1053_v76 = _dafny.SetOf(_952_v4)
			var _1054_v77 _dafny.Array
			_ = _1054_v77
			var _nwElement0_42 _dafny.Int = Companion_Default___.SafeModuloInt((_this).Fm15(p1, p1, globalState), (_dafny.Zero).Minus((_1052_v75).Cardinality()))
			_ = _nwElement0_42
			var _nw182 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(11))
			_ = _nw182
			(_nw182).ArraySet1(_nwElement0_42, 0)
			(_nw182).ArraySet1((_1053_v76).Cardinality(), 1)
			(_nw182).ArraySet1(_dafny.IntOfInt64(886), 2)
			(_nw182).ArraySet1(p1, 3)
			(_nw182).ArraySet1(p2, 4)
			(_nw182).ArraySet1(p1, 5)
			(_nw182).ArraySet1((p2).Times((_956_v8).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_956_v8).Cardinality()))).Uint32()).(_dafny.Int)), 6)
			(_nw182).ArraySet1(_this.F7(), 7)
			(_nw182).ArraySet1(p1, 8)
			(_nw182).ArraySet1((_this).F6(), 9)
			(_nw182).ArraySet1(p2, 10)
			_1054_v77 = _nw182
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1054_v77), 0))
			_ = _index174
			(_1054_v77).ArraySet1(p2, (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1054_v77), 0))
			_ = _index175
			var _rhs220 bool = false
			_ = _rhs220
			var _rhs221 _dafny.Int = ((func() _dafny.Int {
				if (_949_v1).Contains(p0) {
					return (_949_v1).Get(p0).(_dafny.Int)
				}
				return _this.F7()
			})()).Times(Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()))
			_ = _rhs221
			var _lhs145 _dafny.Array = _1054_v77
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(566), _dafny.ArrayLen((_1054_v77), 0))
			_ = _lhs146
			r1 = _rhs220
			(_lhs145).ArraySet1(_rhs221, (_lhs146).Int())
		} else {
			var _1055_v78 *C0
			_ = _1055_v78
			var _nw183 *C0 = New_C0_()
			_ = _nw183
			_nw183.Ctor__(p0)
			_1055_v78 = _nw183
			var _1056_v79 *C0
			_ = _1056_v79
			_1056_v79 = _1055_v78
			var _rhs222 *C0 = _1055_v78
			_ = _rhs222
			var _rhs223 bool = p0
			_ = _rhs223
			_1056_v79 = _rhs222
			r1 = _rhs223
			var _1057_v80 _dafny.Array
			_ = _1057_v80
			var _len0_33 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_33
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_33.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_33)
			} else {
				var _init33 func(_dafny.Int) _dafny.Sequence = (func(_1058_v4 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1059_i5 _dafny.Int) _dafny.Sequence {
						return _1058_v4
					}
				})(_952_v4)
				_ = _init33
				var _element0_33 = _init33(_dafny.Zero)
				_ = _element0_33
				_nw184 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
				(_nw184).ArraySet1(_element0_33, 0)
				var _nativeLen0_33 = (_len0_33).Int()
				_ = _nativeLen0_33
				for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
					(_nw184).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
				}
			}
			_1057_v80 = _nw184
			var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1057_v80), 0))
			_ = _index176
			(_1057_v80).ArraySet1(_952_v4, (_index176).Int())
			var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1057_v80), 0))
			_ = _index177
			(_1057_v80).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_952_v4, _952_v4), (_index177).Int())
			var _1060_v81 _dafny.Map
			_ = _1060_v81
			_1060_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, Companion_Default___.Fm21(globalState))
			_958_v10 = (_958_v10).Update(_this.F7(), _dafny.Companion_Sequence_.IsProperPrefixOf((func() _dafny.Sequence {
				if (_1060_v81).Contains(_this.F7()) {
					return (_1060_v81).Get(_this.F7()).(_dafny.Sequence)
				}
				return (_1057_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1057_v80), 0))).Int()).(_dafny.Sequence)
			})(), (_1057_v80).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1057_v80), 0))).Int()).(_dafny.Sequence)))
			var _1061_v82 _dafny.Array
			_ = _1061_v82
			var _nw185 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw185
			_1061_v82 = _nw185
			var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1061_v82), 0))
			_ = _index178
			(_1061_v82).ArraySet1(p0, (_index178).Int())
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1061_v82), 0))
			_ = _index179
			(_1061_v82).ArraySet1(((_this).F6()).Cmp((_dafny.Zero).Minus(Companion_Default___.Fm1(_1055_v78.F10, globalState))) != 0, (_index179).Int())
			var _1062_v83 _dafny.Array
			_ = _1062_v83
			var _nw186 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(2))
			_ = _nw186
			_1062_v83 = _nw186
			var _1063_v84 _dafny.MultiSet
			_ = _1063_v84
			_1063_v84 = _dafny.MultiSetOf(_954_v6, _954_v6, _954_v6, _954_v6, _954_v6)
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_1062_v83), 0))
			_ = _index180
			(_1062_v83).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v84, (_1061_v82).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1061_v82), 0))).Int()).(bool)), (_index180).Int())
			var _1064_v85 _dafny.Map
			_ = _1064_v85
			_1064_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v84, !(_1055_v78.F10))
			var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_1062_v83), 0))
			_ = _index181
			var _rhs224 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_1065_v6 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1066_i6 _dafny.Int) _dafny.CodePoint {
					return _1065_v6
				}
			})(_954_v6)))).Cardinality()))
			_ = _rhs224
			var _rhs225 bool = _1055_v78.F10
			_ = _rhs225
			var _rhs226 _dafny.CodePoint = _954_v6
			_ = _rhs226
			var _rhs227 _dafny.Map = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1063_v84, true)).Merge(_1064_v85)
			_ = _rhs227
			var _lhs147 *GlobalState = globalState
			_ = _lhs147
			var _lhs148 *C0 = _1055_v78
			_ = _lhs148
			var _lhs149 _dafny.Array = _1062_v83
			_ = _lhs149
			var _lhs150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(922), _dafny.ArrayLen((_1062_v83), 0))
			_ = _lhs150
			_lhs147.F1 = _rhs224
			_lhs148.F10 = _rhs225
			_954_v6 = _rhs226
			(_lhs149).ArraySet1(_rhs227, (_lhs150).Int())
		}
		r0 = p0
		var _1067_v86 _dafny.Set
		_ = _1067_v86
		_1067_v86 = _dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("oh"), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("oh")).Cardinality()))).Uint32(), _dafny.CodePoint('f'))).Cardinality()))
		r1 = (_1067_v86).IsDisjointFrom(_1067_v86)
		r2 = ((_this).F6()).Plus(_dafny.IntOfInt64(324))
		return r0, r1, r2
	}
}
func (_this *C3) M6(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Set, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1068_v0 bool
		_ = _1068_v0
		_1068_v0 = false
		var _1069_i0 _dafny.Int
		_ = _1069_i0
		_1069_i0 = _dafny.Zero
		{
			for !(_1068_v0) {
				{
					if (_1069_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_1069_i0 = (_1069_i0).Plus(_dafny.One)
					var _1070_v1 _dafny.Sequence
					_ = _1070_v1
					_1070_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ea")
					var _1071_v2 _dafny.Array
					_ = _1071_v2
					var _nw187 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
					_ = _nw187
					_1071_v2 = _nw187
					var _1072_v3 D2
					_ = _1072_v3
					_1072_v3 = Companion_D2_.Create_DC5_(_1071_v2)
					var _1073_v4 _dafny.Array
					_ = _1073_v4
					var _nwElement0_43 _dafny.Array = _1071_v2
					_ = _nwElement0_43
					var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(20))
					_ = _nw188
					(_nw188).ArraySet1(_nwElement0_43, 0)
					(_nw188).ArraySet1(_1071_v2, 1)
					(_nw188).ArraySet1(_1071_v2, 2)
					(_nw188).ArraySet1(_1071_v2, 3)
					(_nw188).ArraySet1(_1071_v2, 4)
					(_nw188).ArraySet1(_1071_v2, 5)
					(_nw188).ArraySet1(_1071_v2, 6)
					(_nw188).ArraySet1(_1071_v2, 7)
					(_nw188).ArraySet1(_1071_v2, 8)
					(_nw188).ArraySet1(_1071_v2, 9)
					(_nw188).ArraySet1(_1071_v2, 10)
					(_nw188).ArraySet1(_1071_v2, 11)
					(_nw188).ArraySet1((_1072_v3).Dtor_cf9(), 12)
					(_nw188).ArraySet1(_1071_v2, 13)
					(_nw188).ArraySet1(_1071_v2, 14)
					(_nw188).ArraySet1(_1071_v2, 15)
					(_nw188).ArraySet1((_1072_v3).Dtor_cf9(), 16)
					(_nw188).ArraySet1(_1071_v2, 17)
					(_nw188).ArraySet1(_1071_v2, 18)
					(_nw188).ArraySet1(_1071_v2, 19)
					_1073_v4 = _nw188
					var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1073_v4), 0))
					_ = _index182
					(_1073_v4).ArraySet1((func() _dafny.Array {
						if _1068_v0 {
							return _1071_v2
						}
						return _1071_v2
					})(), (_index182).Int())
					var _1074_v5 _dafny.CodePoint
					_ = _1074_v5
					_1074_v5 = _dafny.CodePoint('g')
					var _1075_v6 _dafny.Map
					_ = _1075_v6
					_1075_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm18(_1068_v0, (_this).F6(), globalState)).Cardinality(), _1074_v5)
					var _1076_v7 _dafny.Map
					_ = _1076_v7
					_1076_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_1075_v6).Cardinality())
					var _1077_v8 _dafny.Sequence
					_ = _1077_v8
					_1077_v8 = _dafny.SeqOf(_this.F7())
					var _1078_v9 _dafny.Sequence
					_ = _1078_v9
					_1078_v9 = _dafny.SeqOf(p1, _dafny.IntOfInt64(938), _dafny.IntOfUint32((_1077_v8).Cardinality()))
					var _1079_v10 _dafny.Sequence
					_ = _1079_v10
					_1079_v10 = _dafny.SeqOf(_1068_v0)
					var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1073_v4), 0))
					_ = _index183
					var _rhs228 _dafny.Int = (p1).Plus(((_1076_v7).Cardinality()).Plus((_this).F6()))
					_ = _rhs228
					var _rhs229 bool = !(!(!(!(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(Companion_Default___.Fm1(_1068_v0, globalState)), _dafny.SeqOf(_dafny.IntOfUint32((_1078_v9).Cardinality())))))))
					_ = _rhs229
					var _rhs230 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1070_v1, _1070_v1)
					_ = _rhs230
					var _rhs231 _dafny.Array = _1071_v2
					_ = _rhs231
					var _rhs232 bool = (_1079_v10).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1079_v10).Cardinality()))).Uint32()).(bool)
					_ = _rhs232
					var _lhs151 *GlobalState = globalState
					_ = _lhs151
					var _lhs152 _dafny.Array = _1073_v4
					_ = _lhs152
					var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(974), _dafny.ArrayLen((_1073_v4), 0))
					_ = _lhs153
					_lhs151.F1 = _rhs228
					r0 = _rhs229
					_1070_v1 = _rhs230
					(_lhs152).ArraySet1(_rhs231, (_lhs153).Int())
					r0 = _rhs232
					_1076_v7 = (_1076_v7).Update(_1068_v0, (_this).F6())
					_1068_v0 = _1068_v0
					var _1080_v11 *C0
					_ = _1080_v11
					var _nw189 *C0 = New_C0_()
					_ = _nw189
					_nw189.Ctor__(_1068_v0)
					_1080_v11 = _nw189
					_1080_v11 = _1080_v11
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _1081_v12 _dafny.Set
		_ = _1081_v12
		_1081_v12 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tayy")).Cardinality()))
		var _1082_v13 _dafny.Map
		_ = _1082_v13
		_1082_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1068_v0, _1068_v0)
		var _1083_v14 _dafny.Map
		_ = _1083_v14
		_1083_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_1082_v13).Contains(_1068_v0) {
				return (_1082_v13).Get(_1068_v0).(bool)
			}
			return (_this).Fm16(globalState)
		})(), _1081_v12)
		var _1084_v16 _dafny.Sequence
		_ = _1084_v16
		_1084_v16 = _dafny.SeqOf(!(!(_1068_v0)), (_this).Fm16(globalState))
		var _1085_v17 _dafny.Array
		_ = _1085_v17
		var _nw190 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
		_ = _nw190
		_1085_v17 = _nw190
		var _1086_v18 D0
		_ = _1086_v18
		_1086_v18 = Companion_D0_.Create_DC1_(_1068_v0, _1085_v17, !(_1068_v0))
		var _1087_v19 _dafny.Map
		_ = _1087_v19
		_1087_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
		var _1088_v20 _dafny.Array
		_ = _1088_v20
		var _len0_34 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_34
		var _nw191 _dafny.Array
		_ = _nw191
		if _len0_34.Cmp(_dafny.Zero) == 0 {
			_nw191 = _dafny.NewArray(_len0_34)
		} else {
			var _init34 func(_dafny.Int) _dafny.Int = func(_1089_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_1089_i1, (_this).F6())
			}
			_ = _init34
			var _element0_34 = _init34(_dafny.Zero)
			_ = _element0_34
			_nw191 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
			(_nw191).ArraySet1(_element0_34, 0)
			var _nativeLen0_34 = (_len0_34).Int()
			_ = _nativeLen0_34
			for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
				(_nw191).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
			}
		}
		_1088_v20 = _nw191
		var _1090_v21 _dafny.MultiSet
		_ = _1090_v21
		_1090_v21 = _dafny.MultiSetOf(_1088_v20, _1088_v20)
		var _1091_v22 _dafny.Array
		_ = _1091_v22
		var _nwElement0_44 bool = _1068_v0
		_ = _nwElement0_44
		var _nw192 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(21))
		_ = _nw192
		(_nw192).ArraySet1(_nwElement0_44, 0)
		(_nw192).ArraySet1((_this.F7()).Cmp((_this).F6()) > 0, 1)
		(_nw192).ArraySet1(_1068_v0, 2)
		(_nw192).ArraySet1(_1068_v0, 3)
		(_nw192).ArraySet1((_1081_v12).IsDisjointFrom((func() _dafny.Set {
			if (_1083_v14).Contains(_1068_v0) {
				return (_1083_v14).Get(_1068_v0).(_dafny.Set)
			}
			return func() _dafny.Set {
				var _coll34 = _dafny.NewBuilder()
				_ = _coll34
				for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-142), _dafny.IntOfInt64(168))); ; {
					_compr_34, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _1092_v15 _dafny.Int
					_1092_v15 = interface{}(_compr_34).(_dafny.Int)
					if ((_dafny.IntOfInt64(-142)).Cmp(_1092_v15) <= 0) && ((_1092_v15).Cmp(_dafny.IntOfInt64(168)) < 0) {
						_coll34.Add((_1092_v15).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qkk")).Cardinality())))
					}
				}
				return _coll34.ToSet()
			}()
		})()), 4)
		(_nw192).ArraySet1(!(_1068_v0), 5)
		(_nw192).ArraySet1((_1084_v16).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1084_v16).Cardinality()))).Uint32()).(bool), 6)
		(_nw192).ArraySet1(_1068_v0, 7)
		(_nw192).ArraySet1(_1068_v0, 8)
		(_nw192).ArraySet1(!(!((_1086_v18).Dtor_cf3())) || (!(_1068_v0)), 9)
		(_nw192).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_1084_v16, _1084_v16), 10)
		(_nw192).ArraySet1(_1068_v0, 11)
		(_nw192).ArraySet1(_1068_v0, 12)
		(_nw192).ArraySet1(!((func() _dafny.Map {
			if (_1087_v19).Contains(p1) {
				return (_1087_v19).Get(p1).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1068_v0)
		})()).Contains(p1), 13)
		(_nw192).ArraySet1(_1068_v0, 14)
		(_nw192).ArraySet1((func() bool {
			if (p0).Contains(p1) {
				return (p0).Get(p1).(bool)
			}
			return _1068_v0
		})(), 15)
		(_nw192).ArraySet1(_1068_v0, 16)
		(_nw192).ArraySet1(_1068_v0, 17)
		(_nw192).ArraySet1(_1068_v0, 18)
		(_nw192).ArraySet1((_1090_v21).IsSubsetOf(_1090_v21), 19)
		(_nw192).ArraySet1(_1068_v0, 20)
		_1091_v22 = _nw192
		var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
		_ = _index184
		(_1091_v22).ArraySet1(false, (_index184).Int())
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_1091_v22), 0))
		_ = _index185
		(_1091_v22).ArraySet1(_1068_v0, (_index185).Int())
		var _1093_v23 _dafny.Sequence
		_ = _1093_v23
		_1093_v23 = _dafny.UnicodeSeqOfUtf8Bytes("w")
		var _1094_v24 _dafny.MultiSet
		_ = _1094_v24
		_1094_v24 = _dafny.MultiSetOf(_1068_v0, _1068_v0, !(!(_1068_v0)))
		var _1095_v25 _dafny.Set
		_ = _1095_v25
		_1095_v25 = _dafny.SetOf(_1068_v0, (_1084_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.IntOfUint32((_1084_v16).Cardinality()))).Uint32()).(bool))
		var _1096_v26 D0
		_ = _1096_v26
		_1096_v26 = Companion_D0_.Create_DC0_(_1095_v25)
		var _pat_let_tv22 = _1082_v13
		_ = _pat_let_tv22
		var _pat_let_tv23 = _1068_v0
		_ = _pat_let_tv23
		var _pat_let_tv24 = _1082_v13
		_ = _pat_let_tv24
		var _pat_let_tv25 = _1068_v0
		_ = _pat_let_tv25
		var _pat_let_tv26 = _1068_v0
		_ = _pat_let_tv26
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
		_ = _index186
		var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_1091_v22), 0))
		_ = _index187
		var _rhs233 bool = !(((((_dafny.MultiSetFromSeq(_1084_v16)).Update(_1068_v0, Companion_Default___.Abs(_dafny.IntOfUint32((_1093_v23).Cardinality())))).Update(_1068_v0, Companion_Default___.Abs(p1))).Update(_1068_v0, Companion_Default___.Abs(_dafny.IntOfInt64(144)))).Equals((Companion_Default___.Fm22(p1, globalState)).Intersection(_1094_v24)))
		_ = _rhs233
		var _rhs234 bool = func(_source12 D0) bool {
			if _source12.Is_DC1() {
				var _1097___mcc_h0 bool = _source12.Get_().(D0_DC1).Cf1
				_ = _1097___mcc_h0
				var _1098___mcc_h1 _dafny.Array = _source12.Get_().(D0_DC1).Cf2
				_ = _1098___mcc_h1
				var _1099___mcc_h2 bool = _source12.Get_().(D0_DC1).Cf3
				_ = _1099___mcc_h2
				var _1100_cf3 bool = _1099___mcc_h2
				_ = _1100_cf3
				var _1101_cf2 _dafny.Array = _1098___mcc_h1
				_ = _1101_cf2
				var _1102_cf1 bool = _1097___mcc_h0
				_ = _1102_cf1
				if (_pat_let_tv22).Contains(_pat_let_tv23) {
					return (_pat_let_tv24).Get(_pat_let_tv25).(bool)
				} else {
					return _pat_let_tv26
				}
			} else {
				var _1103___mcc_h3 _dafny.Set = _source12.Get_().(D0_DC0).Cf0
				_ = _1103___mcc_h3
				var _1104_cf0 _dafny.Set = _1103___mcc_h3
				_ = _1104_cf0
				return false
			}
		}(_1096_v26)
		_ = _rhs234
		var _rhs235 _dafny.Int = (_this).F6()
		_ = _rhs235
		var _lhs154 _dafny.Array = _1091_v22
		_ = _lhs154
		var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
		_ = _lhs155
		var _lhs156 _dafny.Array = _1091_v22
		_ = _lhs156
		var _lhs157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(11), _dafny.ArrayLen((_1091_v22), 0))
		_ = _lhs157
		var _lhs158 *C3 = _this
		_ = _lhs158
		(_lhs154).ArraySet1(_rhs233, (_lhs155).Int())
		(_lhs156).ArraySet1(_rhs234, (_lhs157).Int())
		_lhs158.F7_set_(_rhs235)
		var _pat_let_tv27 = _1093_v23
		_ = _pat_let_tv27
		_1093_v23 = func(_source13 D0) _dafny.Sequence {
			if _source13.Is_DC1() {
				var _1105___mcc_h4 bool = _source13.Get_().(D0_DC1).Cf1
				_ = _1105___mcc_h4
				var _1106___mcc_h5 _dafny.Array = _source13.Get_().(D0_DC1).Cf2
				_ = _1106___mcc_h5
				var _1107___mcc_h6 bool = _source13.Get_().(D0_DC1).Cf3
				_ = _1107___mcc_h6
				var _1108_cf3 bool = _1107___mcc_h6
				_ = _1108_cf3
				var _1109_cf2 _dafny.Array = _1106___mcc_h5
				_ = _1109_cf2
				var _1110_cf1 bool = _1105___mcc_h4
				_ = _1110_cf1
				return _pat_let_tv27
			} else {
				var _1111___mcc_h7 _dafny.Set = _source13.Get_().(D0_DC0).Cf0
				_ = _1111___mcc_h7
				var _1112_cf0 _dafny.Set = _1111___mcc_h7
				_ = _1112_cf0
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(979))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg51 _dafny.Int) interface{} {
						return coer51(arg51)
					}
				}(func(_1113_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))
			}
		}(_1096_v26)
		var _hi10 _dafny.Int = p1
		_ = _hi10
		for _1114_i3 := Companion_Default___.SafeDivisionInt((_this).F6(), p1); _1114_i3.Cmp(_hi10) < 0; _1114_i3 = _1114_i3.Plus(_dafny.One) {
			var _1115_v27 _dafny.Map
			_ = _1115_v27
			_1115_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false)
			_1115_v27 = (_1115_v27).Update(((func() _dafny.Set {
				if true {
					return _1081_v12
				}
				return _1081_v12
			})()).Cardinality(), (_1091_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))).Int()).(bool))
			_1068_v0 = false
			_1068_v0 = (_1091_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))).Int()).(bool)
			var _1116_v28 _dafny.CodePoint
			_ = _1116_v28
			_1116_v28 = _dafny.CodePoint('g')
			var _1117_v29 _dafny.Sequence
			_ = _1117_v29
			_1117_v29 = _dafny.SeqOf(_1095_v25)
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _index188
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _index189
			var _rhs236 bool = !((_1091_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))).Int()).(bool))
			_ = _rhs236
			var _rhs237 bool = !(!_dafny.Companion_Sequence_.Contains(_1093_v23, _1116_v28))
			_ = _rhs237
			var _rhs238 _dafny.Int = (_this).F6()
			_ = _rhs238
			var _rhs239 _dafny.Int = Companion_Default___.SafeModuloInt(Companion_Default___.SafeModuloInt((_this).F6(), _this.F7()), _1114_i3)
			_ = _rhs239
			var _rhs240 bool = ((_1117_v29).Select((Companion_Default___.SafeIndex(_1114_i3, _dafny.IntOfUint32((_1117_v29).Cardinality()))).Uint32()).(_dafny.Set)).IsProperSubsetOf((_1095_v25).Difference(_1095_v25))
			_ = _rhs240
			var _lhs159 _dafny.Array = _1091_v22
			_ = _lhs159
			var _lhs160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _lhs160
			var _lhs161 *GlobalState = globalState
			_ = _lhs161
			var _lhs162 *GlobalState = globalState
			_ = _lhs162
			var _lhs163 _dafny.Array = _1091_v22
			_ = _lhs163
			var _lhs164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _lhs164
			r0 = _rhs236
			(_lhs159).ArraySet1(_rhs237, (_lhs160).Int())
			_lhs161.F1 = _rhs238
			_lhs162.F1 = _rhs239
			(_lhs163).ArraySet1(_rhs240, (_lhs164).Int())
		}
		var _hi11 _dafny.Int = p1
		_ = _hi11
		for _1118_i4 := p1; _1118_i4.Cmp(_hi11) < 0; _1118_i4 = _1118_i4.Plus(_dafny.One) {
			var _1119_v30 _dafny.CodePoint
			_ = _1119_v30
			_1119_v30 = _dafny.CodePoint('c')
			var _1120_v31 _dafny.Sequence
			_ = _1120_v31
			_1120_v31 = _dafny.SeqOf(_1095_v25, _1095_v25)
			var _1121_v32 T0
			_ = _1121_v32
			var _nw193 *C1 = New_C1_()
			_ = _nw193
			_nw193.Ctor__(_1119_v30, Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(95)), Companion_Default___.SafeModuloInt(((_1120_v31).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1120_v31).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), _1118_i4))
			_1121_v32 = _nw193
			var _1122_v33 _dafny.Map
			_ = _1122_v33
			_1122_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _1121_v32)
			_1121_v32 = (func() T0 {
				if (_1122_v33).Contains((_1121_v32.F7()).Cmp(p1) < 0) {
					return (_1122_v33).Get((_1121_v32.F7()).Cmp(p1) < 0).(T0)
				}
				return _1121_v32
			})()
			var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1088_v20), 0))
			_ = _index190
			(_1088_v20).ArraySet1((_dafny.Zero).Minus((_dafny.MultiSetOf(_1119_v30)).Cardinality()), (_index190).Int())
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1088_v20), 0))
			_ = _index191
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _index192
			var _rhs241 bool = _1068_v0
			_ = _rhs241
			var _rhs242 _dafny.Int = _1118_i4
			_ = _rhs242
			var _rhs243 _dafny.Array = _1088_v20
			_ = _rhs243
			var _rhs244 _dafny.Int = _this.F7()
			_ = _rhs244
			var _rhs245 bool = (_1091_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))).Int()).(bool)
			_ = _rhs245
			var _lhs165 *GlobalState = globalState
			_ = _lhs165
			var _lhs166 _dafny.Array = _1088_v20
			_ = _lhs166
			var _lhs167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(509), _dafny.ArrayLen((_1088_v20), 0))
			_ = _lhs167
			var _lhs168 _dafny.Array = _1091_v22
			_ = _lhs168
			var _lhs169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))
			_ = _lhs169
			_1068_v0 = _rhs241
			_lhs165.F1 = _rhs242
			_1088_v20 = _rhs243
			(_lhs166).ArraySet1(_rhs244, (_lhs167).Int())
			(_lhs168).ArraySet1(_rhs245, (_lhs169).Int())
			var _1123_v34 _dafny.Array
			_ = _1123_v34
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(21)
			_ = _len0_35
			var _nw194 _dafny.Array
			_ = _nw194
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw194 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Sequence = (func(_1124_v23 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1125_i5 _dafny.Int) _dafny.Sequence {
						return _1124_v23
					}
				})(_1093_v23)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw194 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw194).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw194).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1123_v34 = _nw194
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_1123_v34), 0))
			_ = _index193
			(_1123_v34).ArraySet1(_1093_v23, (_index193).Int())
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(561), _dafny.ArrayLen((_1123_v34), 0))
			_ = _index194
			(_1123_v34).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xr"), _1093_v23), _dafny.Companion_Sequence_.Concatenate(_1093_v23, _1093_v23)), (_index194).Int())
			var _1126_v35 *C2
			_ = _1126_v35
			var _nw195 *C2 = New_C2_()
			_ = _nw195
			_nw195.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(256), (_1121_v32).F6()), p1)
			_1126_v35 = _nw195
		}
		var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1088_v20), 0))
		_ = _index195
		(_1088_v20).ArraySet1((_this).F6(), (_index195).Int())
		var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(821), _dafny.ArrayLen((_1088_v20), 0))
		_ = _index196
		(_1088_v20).ArraySet1(p1, (_index196).Int())
		var _1127_v36 _dafny.Map
		_ = _1127_v36
		_1127_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-348), (_this).F6())
		var _1128_v37 _dafny.Map
		_ = _1128_v37
		_1128_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1127_v36, (_1091_v22).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(143), _dafny.ArrayLen((_1091_v22), 0))).Int()).(bool))
		r0 = (func() bool {
			if (_1128_v37).Contains(_1127_v36) {
				return (_1128_v37).Get(_1127_v36).(bool)
			}
			return Companion_Default___.Fm0(_1068_v0, _1127_v36, p1, globalState)
		})()
		return r0
	}
}
func (_this *C3) M7(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1129_v0 _dafny.Array
		_ = _1129_v0
		var _nw196 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw196
		_1129_v0 = _nw196
		for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1129_v0), 0))); ; {
			_guard_loop_3, _ok38 := _iter38()
			if !_ok38 {
				break
			}
			var _1130_i0 _dafny.Int
			_1130_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_1130_i0).Sign() != -1) && ((_1130_i0).Cmp(_dafny.ArrayLen((_1129_v0), 0)) < 0)) {
				(_1129_v0).ArraySet1(p1, (_1130_i0).Int())
			}
		}
		var _1131_i1 _dafny.Int
		_ = _1131_i1
		_1131_i1 = _dafny.Zero
		{
			for p1 {
				{
					if (_1131_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1131_i1 = (_1131_i1).Plus(_dafny.One)
					var _1132_v1 bool
					_ = _1132_v1
					_1132_v1 = false
					var _1133_v2 _dafny.Sequence
					_ = _1133_v2
					_1133_v2 = _dafny.SeqOf(_1132_v1)
					var _rhs246 _dafny.Array = _1129_v0
					_ = _rhs246
					var _rhs247 bool = p0
					_ = _rhs247
					var _rhs248 _dafny.Sequence = _dafny.SeqOf(!(p0) || (true), !(((_this).F6()).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vgkqtd")).Cardinality())) > 0), p0, p1)
					_ = _rhs248
					var _rhs249 bool = ((p2).Plus(p2)).Cmp((_this).F6()) == 0
					_ = _rhs249
					_1129_v0 = _rhs246
					_1132_v1 = _rhs247
					_1133_v2 = _rhs248
					_1132_v1 = _rhs249
					var _1134_v3 _dafny.Array
					_ = _1134_v3
					var _len0_36 _dafny.Int = _dafny.IntOfInt64(29)
					_ = _len0_36
					var _nw197 _dafny.Array
					_ = _nw197
					if _len0_36.Cmp(_dafny.Zero) == 0 {
						_nw197 = _dafny.NewArray(_len0_36)
					} else {
						var _init36 func(_dafny.Int) _dafny.Int = func(_1135_i2 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1135_i2, (_this).F6())
						}
						_ = _init36
						var _element0_36 = _init36(_dafny.Zero)
						_ = _element0_36
						_nw197 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
						(_nw197).ArraySet1(_element0_36, 0)
						var _nativeLen0_36 = (_len0_36).Int()
						_ = _nativeLen0_36
						for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
							(_nw197).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
						}
					}
					_1134_v3 = _nw197
					var _1136_v4 _dafny.Sequence
					_ = _1136_v4
					_1136_v4 = _dafny.SeqOf(_1134_v3)
					var _1137_v5 _dafny.Map
					_ = _1137_v5
					_1137_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1136_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-925), _dafny.IntOfUint32((_1136_v4).Cardinality()))).Uint32()).(_dafny.Array), _1132_v1)
					_1137_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1134_v3, _1132_v1)
					var _1138_v6 D0
					_ = _1138_v6
					_1138_v6 = Companion_D0_.Create_DC1_(true, _1129_v0, _1132_v1)
					var _1139_v7 _dafny.Map
					_ = _1139_v7
					_1139_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(273), _1132_v1)
					var _1140_v8 _dafny.Int
					_ = _1140_v8
					var _1141_v9 bool
					_ = _1141_v9
					var _1142_v10 _dafny.Array
					_ = _1142_v10
					var _out20 _dafny.Int
					_ = _out20
					var _out21 bool
					_ = _out21
					var _out22 _dafny.Array
					_ = _out22
					_out20, _out21, _out22 = Companion_Default___.M0(_1138_v6, Companion_Default___.Fm1(!((func() bool {
						if (_1139_v7).Contains((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("blo")).Cardinality())))) {
							return (_1139_v7).Get((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("blo")).Cardinality())))).(bool)
						}
						return false
					})()), globalState), globalState)
					_1140_v8 = _out20
					_1141_v9 = _out21
					_1142_v10 = _out22
					var _1143_v11 _dafny.MultiSet
					_ = _1143_v11
					_1143_v11 = _dafny.MultiSetOf(p0)
					var _1144_v12 _dafny.Sequence
					_ = _1144_v12
					_1144_v12 = _dafny.SeqOf((_1143_v11).Cardinality(), _1140_v8)
					var _1145_v13 _dafny.Map
					_ = _1145_v13
					_1145_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1144_v12).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1144_v12).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfInt64(144))
					var _1146_v14 _dafny.Sequence
					_ = _1146_v14
					_1146_v14 = _dafny.UnicodeSeqOfUtf8Bytes("aywp")
					if Companion_Default___.Fm0(_1132_v1, _1145_v13, _dafny.IntOfUint32((_1146_v14).Cardinality()), globalState) {
						var _1147_v15 _dafny.Set
						_ = _1147_v15
						_1147_v15 = _dafny.SetOf((_this).F6(), _dafny.IntOfInt64(-188))
						var _1148_v16 *C1
						_ = _1148_v16
						var _nw198 *C1 = New_C1_()
						_ = _nw198
						_nw198.Ctor__(_dafny.CodePoint('s'), (_1147_v15).Cardinality(), (Companion_Default___.Fm36(_1140_v8, true, (_this).F6(), _1141_v9, globalState)).Cardinality())
						_1148_v16 = _nw198
						(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1146_v14), (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_dafny.SeqOf(_1146_v14)).Cardinality()))).Uint32(), _1146_v14)).Cardinality())
						var _1149_v17 D9
						_ = _1149_v17
						_1149_v17 = Companion_D9_.Create_DC22_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-814))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg52 _dafny.Int) interface{} {
								return coer52(arg52)
							}
						}((func(_1150_v8 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_1151_i3 _dafny.Int) _dafny.Int {
								return _1150_v8
							}
						})(_1140_v8))), _1144_v12))
						_1149_v17 = _1149_v17
						var _1152_v18 *C1
						_ = _1152_v18
						var _nw199 *C1 = New_C1_()
						_ = _nw199
						_nw199.Ctor__(Companion_Default___.Fm25(globalState), Companion_Default___.Fm1(_1132_v1, globalState), (_this).F6())
						_1152_v18 = _nw199
						_1146_v14 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(440))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg53 _dafny.Int) interface{} {
								return coer53(arg53)
							}
						}((func(_1153_v16 *C1) func(_dafny.Int) _dafny.CodePoint {
							return func(_1154_i4 _dafny.Int) _dafny.CodePoint {
								return (_1153_v16).F13()
							}
						})(_1148_v16)))
					} else {
						var _1155_v19 _dafny.CodePoint
						_ = _1155_v19
						_1155_v19 = _dafny.CodePoint('i')
						var _1156_v20 *C1
						_ = _1156_v20
						var _nw200 *C1 = New_C1_()
						_ = _nw200
						_nw200.Ctor__(_1155_v19, _1140_v8, ((_this).F6()).Plus(p2))
						_1156_v20 = _nw200
						(globalState).F1 = (_this).F6()
						var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1129_v0), 0))
						_ = _index197
						(_1129_v0).ArraySet1(_dafny.Companion_Sequence_.Contains(_1146_v14, _1155_v19), (_index197).Int())
						var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))
						_ = _index198
						(_1134_v3).ArraySet1(_this.F7(), (_index198).Int())
						var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1129_v0), 0))
						_ = _index199
						var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))
						_ = _index200
						var _rhs250 _dafny.Array = _1129_v0
						_ = _rhs250
						var _rhs251 _dafny.Array = _1129_v0
						_ = _rhs251
						var _rhs252 bool = p0
						_ = _rhs252
						var _rhs253 _dafny.Int = _dafny.IntOfUint32((_1146_v14).Cardinality())
						_ = _rhs253
						var _rhs254 _dafny.Int = _dafny.IntOfInt64(-972)
						_ = _rhs254
						var _lhs170 _dafny.Array = _1129_v0
						_ = _lhs170
						var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1129_v0), 0))
						_ = _lhs171
						var _lhs172 _dafny.Array = _1134_v3
						_ = _lhs172
						var _lhs173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))
						_ = _lhs173
						_1129_v0 = _rhs250
						_1129_v0 = _rhs251
						(_lhs170).ArraySet1(_rhs252, (_lhs171).Int())
						r0 = _rhs253
						(_lhs172).ArraySet1(_rhs254, (_lhs173).Int())
						var _1157_v21 _dafny.Map
						_ = _1157_v21
						_1157_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1133_v2, _1133_v2), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1133_v2, _1133_v2)).Cardinality()))).Uint32(), p0), (func() bool {
							if p1 {
								return _1132_v1
							}
							return _1132_v1
						})())
						_1157_v21 = (_1157_v21).Update(Companion_Default___.Fm21(globalState), !((_1129_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(559), _dafny.ArrayLen((_1129_v0), 0))).Int()).(bool)))
						var _1158_v22 _dafny.Set
						_ = _1158_v22
						_1158_v22 = _dafny.SetOf(p0)
						var _1159_v23 _dafny.Map
						_ = _1159_v23
						_1159_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1140_v8)
						var _1160_v25 _dafny.MultiSet
						_ = _1160_v25
						_1160_v25 = _dafny.MultiSetOf((_this).F6())
						var _1161_v28 _dafny.Array
						_ = _1161_v28
						var _nwElement0_45 _dafny.Int = (Companion_Default___.Fm1(p0, globalState)).Minus((_this).F6())
						_ = _nwElement0_45
						var _nw201 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(17))
						_ = _nw201
						(_nw201).ArraySet1(_nwElement0_45, 0)
						(_nw201).ArraySet1(((_1134_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))).Int()).(_dafny.Int)).Minus(_1140_v8), 1)
						(_nw201).ArraySet1(Companion_Default___.Fm1(p1, globalState), 2)
						(_nw201).ArraySet1(_1140_v8, 3)
						(_nw201).ArraySet1((func() _dafny.Int {
							if p0 {
								return (_1158_v22).Cardinality()
							}
							return (_1159_v23).Cardinality()
						})(), 4)
						(_nw201).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-436), p2), 5)
						(_nw201).ArraySet1(_1140_v8, 6)
						(_nw201).ArraySet1(((_1134_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))).Int()).(_dafny.Int)).Plus((_dafny.MultiSetOf(_1140_v8)).Cardinality()), 7)
						(_nw201).ArraySet1(Companion_Default___.Fm1(_1141_v9, globalState), 8)
						(_nw201).ArraySet1(((func() _dafny.Set {
							var _coll35 = _dafny.NewBuilder()
							_ = _coll35
							for _iter39 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(219), _dafny.IntOfInt64(988))); ; {
								_compr_35, _ok39 := _iter39()
								if !_ok39 {
									break
								}
								var _1162_v24 _dafny.Int
								_1162_v24 = interface{}(_compr_35).(_dafny.Int)
								if ((_dafny.IntOfInt64(219)).Cmp(_1162_v24) <= 0) && ((_1162_v24).Cmp(_dafny.IntOfInt64(988)) < 0) {
									_coll35.Add((_1162_v24).Times(p2))
								}
							}
							return _coll35.ToSet()
						}()).Union(func() _dafny.Set {
							var _coll36 = _dafny.NewBuilder()
							_ = _coll36
							for _iter40 := _dafny.Iterate((_1160_v25).Elements()); ; {
								_compr_36, _ok40 := _iter40()
								if !_ok40 {
									break
								}
								var _1163_v26 _dafny.Int
								_1163_v26 = interface{}(_compr_36).(_dafny.Int)
								if (_1160_v25).Contains(_1163_v26) {
									_coll36.Add((_1163_v26).Plus(_dafny.IntOfInt64(803)))
								}
							}
							return _coll36.ToSet()
						}())).Cardinality(), 9)
						(_nw201).ArraySet1(((_1134_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(795), _dafny.ArrayLen((_1134_v3), 0))).Int()).(_dafny.Int)).Plus((func() _dafny.Map {
							var _coll37 = _dafny.NewMapBuilder()
							_ = _coll37
							for _iter41 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(785), _dafny.IntOfInt64(192))); ; {
								_compr_37, _ok41 := _iter41()
								if !_ok41 {
									break
								}
								var _1164_v27 _dafny.Int
								_1164_v27 = interface{}(_compr_37).(_dafny.Int)
								if ((_dafny.IntOfInt64(785)).Cmp(_1164_v27) <= 0) && ((_1164_v27).Cmp(_dafny.IntOfInt64(192)) < 0) {
									_coll37.Add(Companion_Default___.SafeModuloInt(_1164_v27, _1140_v8), p1)
								}
							}
							return _coll37.ToMap()
						}()).Cardinality()), 10)
						(_nw201).ArraySet1(p2, 11)
						(_nw201).ArraySet1(p2, 12)
						(_nw201).ArraySet1(p2, 13)
						(_nw201).ArraySet1((_this).Fm15(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg54 _dafny.Int) interface{} {
								return coer54(arg54)
							}
						}((func(_1165_v12 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1166_i5 _dafny.Int) _dafny.Sequence {
								return _1165_v12
							}
						})(_1144_v12))), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(77), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(543))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
							return func(arg55 _dafny.Int) interface{} {
								return coer55(arg55)
							}
						}((func(_1167_v12 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
							return func(_1168_i5 _dafny.Int) _dafny.Sequence {
								return _1167_v12
							}
						})(_1144_v12)))).Cardinality()))).Uint32(), _1144_v12)).Cardinality()), _this.F7(), globalState), 14)
						(_nw201).ArraySet1((_dafny.IntOfInt64(270)).Times(p2), 15)
						(_nw201).ArraySet1((_this).F6(), 16)
						_1161_v28 = _nw201
						_1161_v28 = _1161_v28
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _1169_v29 _dafny.MultiSet
		_ = _1169_v29
		_1169_v29 = _dafny.MultiSetOf(p2)
		var _1170_v30 _dafny.Sequence
		_ = _1170_v30
		_1170_v30 = _dafny.SeqOf(Companion_Default___.Fm35(_this.F7(), (_1169_v29).Cardinality(), globalState))
		var _1171_v31 _dafny.Sequence
		_ = _1171_v31
		_1171_v31 = _dafny.SeqOf(_dafny.IntOfUint32(((_1170_v30).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1170_v30).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))
		var _1172_v32 _dafny.Sequence
		_ = _1172_v32
		_1172_v32 = _dafny.SeqOf((_1171_v31).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1171_v31).Cardinality()))).Uint32()).(_dafny.Int), p2)
		var _1173_v33 _dafny.Sequence
		_ = _1173_v33
		_1173_v33 = _dafny.SeqOf(_dafny.IntOfUint32((_1172_v32).Cardinality()))
		r1 = (_dafny.Zero).Minus((_1171_v31).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1171_v31).Cardinality()))).Uint32()).(_dafny.Int))
		var _1174_v34 _dafny.CodePoint
		_ = _1174_v34
		_1174_v34 = _dafny.CodePoint('e')
		var _1175_v35 *C1
		_ = _1175_v35
		var _nw202 *C1 = New_C1_()
		_ = _nw202
		_nw202.Ctor__(_1174_v34, p2, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-485))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}(func(_1176_i6 _dafny.Int) _dafny.CodePoint {
			return (Companion_D10_.Create_DC24_(_dafny.CodePoint('e'))).Dtor_cf51()
		}))).Cardinality()))
		_1175_v35 = _nw202
		var _rhs255 _dafny.CodePoint = _1174_v34
		_ = _rhs255
		var _rhs256 *C1 = _1175_v35
		_ = _rhs256
		_1174_v34 = _rhs255
		_1175_v35 = _rhs256
		var _1177_v36 _dafny.Map
		_ = _1177_v36
		_1177_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Plus(p2), _1175_v35)
		_1177_v36 = (_1177_v36).Update(p2, _1175_v35)
		var _1178_v37 bool
		_ = _1178_v37
		_1178_v37 = true
		var _1179_v38 D10
		_ = _1179_v38
		_1179_v38 = Companion_D10_.Create_DC26_(_1178_v37)
		_1178_v37 = (_1179_v38).Dtor_cf52()
		var _1180_v39 _dafny.Sequence
		_ = _1180_v39
		_1180_v39 = _dafny.UnicodeSeqOfUtf8Bytes("vdmx")
		r0 = (_dafny.IntOfUint32((_1180_v39).Cardinality())).Times(p2)
		r1 = ((_this).F6()).Plus(Companion_Default___.SafeDivisionInt(p2, p2))
		r2 = p2
		return r0, r1, r2
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f7  _dafny.Int
	_f6  _dafny.Int
	F12  _dafny.Map
	_f11 bool
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f7 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this.F12 = _dafny.EmptyMap
	_this._f11 = false
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

func (_this *C4) F7() _dafny.Int {
	return _this._f7
}
func (_this *C4) F7_set_(value _dafny.Int) {
	_this._f7 = value
}
func (_this *C4) F6() _dafny.Int {
	return _this._f6
}
func (_this *C4) Ctor__(f11 bool, f12 _dafny.Map, f6 _dafny.Int, f7 _dafny.Int) {
	{
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C4) Fm11(p0 _dafny.CodePoint, globalState *GlobalState) D0 {
	{
		return Companion_D0_.Create_DC0_((_dafny.SetOf((_this).F11())).Intersection(_dafny.SetOf((_this).F11(), false)))
	}
}
func (_this *C4) Fm12(p0 bool, p1 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.CodePoint('t'), _dafny.CodePoint('g')), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-535))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg57 _dafny.Int) interface{} {
				return coer57(arg57)
			}
		}(func(_1181_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(511))).Uint32(), func(coer58 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg58 _dafny.Int) interface{} {
				return coer58(arg58)
			}
		}(func(_1182_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('q')
		})), _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('u'))))
	}
}
func (_this *C4) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _hi12 _dafny.Int = Companion_Default___.SafeDivisionInt(Companion_Default___.Fm1(p0, globalState), p1)
		_ = _hi12
		for _1183_i0 := (_this).F6(); _1183_i0.Cmp(_hi12) < 0; _1183_i0 = _1183_i0.Plus(_dafny.One) {
			var _1184_v0 _dafny.MultiSet
			_ = _1184_v0
			_1184_v0 = _dafny.MultiSetOf(_1183_i0)
			var _1185_v1 _dafny.Sequence
			_ = _1185_v1
			_1185_v1 = _dafny.SeqOf(_dafny.IntOfInt64(77))
			_1184_v0 = _dafny.MultiSetFromSeq(_1185_v1)
			var _1186_v2 _dafny.Map
			_ = _1186_v2
			_1186_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1183_i0, Companion_Default___.SafeDivisionInt(p2, _dafny.IntOfInt64(76)))
			(_this).F7_set_((func() _dafny.Int {
				if (_1186_v2).Contains((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1185_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1183_i0), _dafny.IntOfUint32((_1185_v1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F6()))) {
					return (_1186_v2).Get((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_1185_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_1183_i0), _dafny.IntOfUint32((_1185_v1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F6()))).(_dafny.Int)
				}
				return _this.F7()
			})())
			var _1187_v3 _dafny.Sequence
			_ = _1187_v3
			_1187_v3 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			_1187_v3 = _1187_v3
			(_this).F7_set_((p1).Minus(_dafny.IntOfInt64(228)))
		}
		var _1188_v4 _dafny.Sequence
		_ = _1188_v4
		_1188_v4 = _dafny.UnicodeSeqOfUtf8Bytes("uwyinhv")
		var _1189_v5 _dafny.Array
		_ = _1189_v5
		var _len0_37 _dafny.Int = _dafny.IntOfInt64(27)
		_ = _len0_37
		var _nw203 _dafny.Array
		_ = _nw203
		if _len0_37.Cmp(_dafny.Zero) == 0 {
			_nw203 = _dafny.NewArray(_len0_37)
		} else {
			var _init37 func(_dafny.Int) _dafny.Int = func(_1190_i1 _dafny.Int) _dafny.Int {
				return (_1190_i1).Times((_this).F6())
			}
			_ = _init37
			var _element0_37 = _init37(_dafny.Zero)
			_ = _element0_37
			_nw203 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
			(_nw203).ArraySet1(_element0_37, 0)
			var _nativeLen0_37 = (_len0_37).Int()
			_ = _nativeLen0_37
			for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
				(_nw203).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
			}
		}
		_1189_v5 = _nw203
		var _1191_v6 _dafny.Sequence
		_ = _1191_v6
		_1191_v6 = _dafny.SeqOf(_1189_v5, _1189_v5)
		var _1192_v7 _dafny.Array
		_ = _1192_v7
		var _nwElement0_46 _dafny.Int = p2
		_ = _nwElement0_46
		var _nw204 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(9))
		_ = _nw204
		(_nw204).ArraySet1(_nwElement0_46, 0)
		(_nw204).ArraySet1(p2, 1)
		(_nw204).ArraySet1((_dafny.Zero).Minus(p2), 2)
		(_nw204).ArraySet1(_dafny.IntOfUint32((_1188_v4).Cardinality()), 3)
		(_nw204).ArraySet1(p1, 4)
		(_nw204).ArraySet1(p1, 5)
		(_nw204).ArraySet1(_this.F7(), 6)
		(_nw204).ArraySet1(_dafny.IntOfUint32((_1191_v6).Cardinality()), 7)
		(_nw204).ArraySet1(_this.F7(), 8)
		_1192_v7 = _nw204
		var _1193_v8 D2
		_ = _1193_v8
		_1193_v8 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_(_1189_v5))
		_1193_v8 = _1193_v8
		var _1194_v9 *C0
		_ = _1194_v9
		var _nw205 *C0 = New_C0_()
		_ = _nw205
		_nw205.Ctor__(false)
		_1194_v9 = _nw205
		var _1195_v10 _dafny.MultiSet
		_ = _1195_v10
		_1195_v10 = _dafny.MultiSetOf(p0, _1194_v9.F10)
		var _1196_v11 _dafny.Map
		_ = _1196_v11
		_1196_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _1195_v10)
		var _1197_v12 _dafny.CodePoint
		_ = _1197_v12
		_1197_v12 = _dafny.CodePoint('n')
		var _1198_v13 _dafny.MultiSet
		_ = _1198_v13
		_1198_v13 = ((func() _dafny.MultiSet {
			if (_1196_v11).Contains(_1197_v12) {
				return (_1196_v11).Get(_1197_v12).(_dafny.MultiSet)
			}
			return _dafny.MultiSetOf(_1194_v9.F10)
		})()).Union(_1195_v10)
		var _source14 _dafny.MultiSet = _1198_v13
		_ = _source14
		var _1199___mcc_h0 _dafny.MultiSet = _source14
		_ = _1199___mcc_h0
		var _1200_cf18 _dafny.MultiSet = _1199___mcc_h0
		_ = _1200_cf18
		var _1201_v14 _dafny.Array
		_ = _1201_v14
		var _nw206 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
		_ = _nw206
		_1201_v14 = _nw206
		var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_1201_v14), 0))
		_ = _index201
		(_1201_v14).ArraySet1(_1194_v9.F10, (_index201).Int())
		var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_1201_v14), 0))
		_ = _index202
		(_1201_v14).ArraySet1(true, (_index202).Int())
		var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(130), _dafny.ArrayLen((_1201_v14), 0))
		_ = _index203
		(_1201_v14).ArraySet1((_this).F11(), (_index203).Int())
		var _1202_v15 _dafny.MultiSet
		_ = _1202_v15
		_1202_v15 = _dafny.MultiSetOf(Companion_Default___.Fm1(p0, globalState))
		var _1203_v16 _dafny.Map
		_ = _1203_v16
		_1203_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1202_v15, p0)
		var _1204_v17 _dafny.Map
		_ = _1204_v17
		_1204_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_1203_v16).Contains((_1202_v15).Update(_dafny.IntOfInt64(520), Companion_Default___.Abs((_this).F6()))) {
				return (_1203_v16).Get((_1202_v15).Update(_dafny.IntOfInt64(520), Companion_Default___.Abs((_this).F6()))).(bool)
			}
			return _1194_v9.F10
		})(), _dafny.IntOfInt64(266))
		_1204_v17 = (_1204_v17).Merge((_1204_v17).Merge(_1204_v17))
		var _1205_v18 _dafny.Map
		_ = _1205_v18
		_1205_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _1194_v9.F10)
		var _1206_v19 _dafny.Map
		_ = _1206_v19
		_1206_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("lbfloex"), _1197_v12)
		_1205_v18 = (_1205_v18).Update((_1206_v19).Cardinality(), _1194_v9.F10)
		var _1207_v20 _dafny.Map
		_ = _1207_v20
		_1207_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p1), !((_this).F11()))
		var _1208_v21 D1
		_ = _1208_v21
		_1208_v21 = Companion_D1_.Create_DC4_(p0)
		var _1209_v22 D2
		_ = _1209_v22
		_1209_v22 = Companion_D2_.Create_DC7_(p0, _dafny.IntOfInt64(79), _1207_v20, _1208_v21)
		var _source15 D2 = _1209_v22
		_ = _source15
		if _source15.Is_DC6() {
			var _1210___mcc_h1 D1 = _source15.Get_().(D2_DC6).Cf10
			_ = _1210___mcc_h1
			var _1211___mcc_h2 _dafny.Sequence = _source15.Get_().(D2_DC6).Cf11
			_ = _1211___mcc_h2
			var _1212___mcc_h3 _dafny.Int = _source15.Get_().(D2_DC6).Cf12
			_ = _1212___mcc_h3
			var _1213_cf12 _dafny.Int = _1212___mcc_h3
			_ = _1213_cf12
			var _1214_cf11 _dafny.Sequence = _1211___mcc_h2
			_ = _1214_cf11
			var _1215_cf10 D1 = _1210___mcc_h1
			_ = _1215_cf10
			var _1216_v23 _dafny.Sequence
			_ = _1216_v23
			_1216_v23 = _dafny.SeqOf(_1194_v9.F10, _1194_v9.F10)
			r0 = !(((Companion_Default___.Fm13(globalState)).Update(_1194_v9.F10, Companion_Default___.Abs(_1213_cf12))).IsSubsetOf((Companion_Default___.Fm13(globalState)).Update(p0, Companion_Default___.Abs((_dafny.Zero).Minus(_dafny.IntOfUint32((_1216_v23).Cardinality()))))))
			r1 = (_this).F11()
			var _1217_v24 _dafny.Array
			_ = _1217_v24
			var _nw207 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(27))
			_ = _nw207
			_1217_v24 = _nw207
			_1217_v24 = _1217_v24
			var _1218_v25 *C0
			_ = _1218_v25
			var _nw208 *C0 = New_C0_()
			_ = _nw208
			_nw208.Ctor__(p0)
			_1218_v25 = _nw208
		} else if _source15.Is_DC7() {
			var _1219___mcc_h4 bool = _source15.Get_().(D2_DC7).Cf13
			_ = _1219___mcc_h4
			var _1220___mcc_h5 _dafny.Int = _source15.Get_().(D2_DC7).Cf14
			_ = _1220___mcc_h5
			var _1221___mcc_h6 _dafny.Map = _source15.Get_().(D2_DC7).Cf15
			_ = _1221___mcc_h6
			var _1222___mcc_h7 D1 = _source15.Get_().(D2_DC7).Cf16
			_ = _1222___mcc_h7
			var _1223_cf16 D1 = _1222___mcc_h7
			_ = _1223_cf16
			var _1224_cf15 _dafny.Map = _1221___mcc_h6
			_ = _1224_cf15
			var _1225_cf14 _dafny.Int = _1220___mcc_h5
			_ = _1225_cf14
			var _1226_cf13 bool = _1219___mcc_h4
			_ = _1226_cf13
			var _1227_v26 _dafny.Array
			_ = _1227_v26
			var _len0_38 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_38
			var _nw209 _dafny.Array
			_ = _nw209
			if _len0_38.Cmp(_dafny.Zero) == 0 {
				_nw209 = _dafny.NewArray(_len0_38)
			} else {
				var _init38 func(_dafny.Int) bool = func(_1228_i2 _dafny.Int) bool {
					return (false) && (false)
				}
				_ = _init38
				var _element0_38 = _init38(_dafny.Zero)
				_ = _element0_38
				_nw209 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
				(_nw209).ArraySet1(_element0_38, 0)
				var _nativeLen0_38 = (_len0_38).Int()
				_ = _nativeLen0_38
				for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
					(_nw209).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
				}
			}
			_1227_v26 = _nw209
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1227_v26), 0))
			_ = _index204
			(_1227_v26).ArraySet1(_1194_v9.F10, (_index204).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.ArrayLen((_1227_v26), 0))
			_ = _index205
			(_1227_v26).ArraySet1(_1194_v9.F10, (_index205).Int())
			(_1194_v9).F10 = _1226_cf13
			_1188_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("nvthhe"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(504))).Uint32(), func(coer59 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg59 _dafny.Int) interface{} {
					return coer59(arg59)
				}
			}((func(_1229_v12 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1230_i3 _dafny.Int) _dafny.CodePoint {
					return _1229_v12
				}
			})(_1197_v12))))
			var _1231_v27 D0
			_ = _1231_v27
			_1231_v27 = Companion_D0_.Create_DC0_(_dafny.SetOf((_this).F11(), (_this).F11()))
			var _1232_v28 _dafny.Map
			_ = _1232_v28
			_1232_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1231_v27, _1225_cf14)
			_1232_v28 = (_1232_v28).Update((_this).Fm11(_1197_v12, globalState), p1)
		} else if _source15.Is_DC5() {
			var _1233___mcc_h8 _dafny.Array = _source15.Get_().(D2_DC5).Cf9
			_ = _1233___mcc_h8
			var _1234_cf9 _dafny.Array = _1233___mcc_h8
			_ = _1234_cf9
			(globalState).F1 = Companion_Default___.Fm1(Companion_Default___.Fm0(_1194_v9.F10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), p2), (_dafny.Zero).Minus(p1), globalState), globalState)
			if _1194_v9.F10 {
				var _1235_v29 _dafny.Map
				_ = _1235_v29
				_1235_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(639), p2)
				_1235_v29 = (_1235_v29).Update(((_this).F6()).Minus(_dafny.IntOfInt64(750)), _dafny.IntOfUint32((_1188_v4).Cardinality()))
				var _1236_v30 D2
				_ = _1236_v30
				_1236_v30 = Companion_D2_.Create_DC5_(_1189_v5)
				var _1237_v31 _dafny.Map
				_ = _1237_v31
				_1237_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1207_v20, _dafny.IntOfInt64(243))
				var _rhs257 _dafny.Int = (func() _dafny.Int {
					if p0 {
						return (_dafny.Zero).Minus(p2)
					}
					return (_1237_v31).Cardinality()
				})()
				_ = _rhs257
				var _rhs258 D2 = _1236_v30
				_ = _rhs258
				var _lhs174 *GlobalState = globalState
				_ = _lhs174
				_lhs174.F1 = _rhs257
				_1236_v30 = _rhs258
				(_1194_v9).F10 = true
				var _1238_v32 _dafny.Array
				_ = _1238_v32
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw210
				_1238_v32 = _nw210
				var _1239_v33 D0
				_ = _1239_v33
				_1239_v33 = Companion_D0_.Create_DC1_(false, _1238_v32, (_this).F11())
				var _1240_v34 _dafny.Int
				_ = _1240_v34
				var _1241_v35 bool
				_ = _1241_v35
				var _1242_v36 _dafny.Array
				_ = _1242_v36
				var _out23 _dafny.Int
				_ = _out23
				var _out24 bool
				_ = _out24
				var _out25 _dafny.Array
				_ = _out25
				_out23, _out24, _out25 = Companion_Default___.M0(_1239_v33, (_this).F6(), globalState)
				_1240_v34 = _out23
				_1241_v35 = _out24
				_1242_v36 = _out25
				_1235_v29 = (_1235_v29).Update(_this.F7(), (_this).F6())
			} else {
				(globalState).F1 = (_this).F6()
				var _1243_v37 _dafny.Array
				_ = _1243_v37
				var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
				_ = _nw211
				_1243_v37 = _nw211
				_1243_v37 = _1243_v37
				(_1194_v9).F10 = p0
				_1207_v20 = (_1207_v20).Update(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-709), _this.F7()), (_this).F11())
				var _1244_v38 _dafny.MultiSet
				_ = _1244_v38
				_1244_v38 = _dafny.MultiSetOf(p2)
				var _1245_v39 _dafny.Sequence
				_ = _1245_v39
				_1245_v39 = _dafny.SeqOf(_1244_v38)
				(globalState).F1 = _dafny.IntOfUint32((_1245_v39).Cardinality())
			}
			_1189_v5 = _1234_cf9
			var _1246_v40 _dafny.Map
			_ = _1246_v40
			_1246_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1194_v9.F10, (_this).F11())
			var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1234_cf9), 0))
			_ = _index206
			(_1234_cf9).ArraySet1(((_1246_v40).Merge(_1246_v40)).Cardinality(), (_index206).Int())
			var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_1234_cf9), 0))
			_ = _index207
			(_1234_cf9).ArraySet1((func() _dafny.Int {
				if false {
					return (_this).F6()
				}
				return p1
			})(), (_index207).Int())
		} else {
			var _1247___mcc_h9 D2 = _source15.Get_().(D2_DC8).Cf17
			_ = _1247___mcc_h9
			var _1248_cf17 D2 = _1247___mcc_h9
			_ = _1248_cf17
			r1 = _1194_v9.F10
			var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1189_v5), 0))
			_ = _index208
			(_1189_v5).ArraySet1(p2, (_index208).Int())
			var _1249_v41 _dafny.Array
			_ = _1249_v41
			var _len0_39 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_39
			var _nw212 _dafny.Array
			_ = _nw212
			if _len0_39.Cmp(_dafny.Zero) == 0 {
				_nw212 = _dafny.NewArray(_len0_39)
			} else {
				var _init39 func(_dafny.Int) bool = (func(_1250_v9 *C0) func(_dafny.Int) bool {
					return func(_1251_i4 _dafny.Int) bool {
						return _1250_v9.F10
					}
				})(_1194_v9)
				_ = _init39
				var _element0_39 = _init39(_dafny.Zero)
				_ = _element0_39
				_nw212 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
				(_nw212).ArraySet1(_element0_39, 0)
				var _nativeLen0_39 = (_len0_39).Int()
				_ = _nativeLen0_39
				for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
					(_nw212).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
				}
			}
			_1249_v41 = _nw212
			var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_1249_v41), 0))
			_ = _index209
			(_1249_v41).ArraySet1(!(p0), (_index209).Int())
			var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1192_v7), 0))
			_ = _index210
			(_1192_v7).ArraySet1(_this.F7(), (_index210).Int())
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1189_v5), 0))
			_ = _index211
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_1249_v41), 0))
			_ = _index212
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1192_v7), 0))
			_ = _index213
			var _rhs259 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("mdoyxqucy"), _1188_v4)
			_ = _rhs259
			var _rhs260 _dafny.Int = p2
			_ = _rhs260
			var _rhs261 bool = p0
			_ = _rhs261
			var _rhs262 _dafny.Int = _dafny.IntOfInt64(-836)
			_ = _rhs262
			var _rhs263 bool = (_this).F11()
			_ = _rhs263
			var _lhs175 *C0 = _1194_v9
			_ = _lhs175
			var _lhs176 _dafny.Array = _1189_v5
			_ = _lhs176
			var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_1189_v5), 0))
			_ = _lhs177
			var _lhs178 _dafny.Array = _1249_v41
			_ = _lhs178
			var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(158), _dafny.ArrayLen((_1249_v41), 0))
			_ = _lhs179
			var _lhs180 _dafny.Array = _1192_v7
			_ = _lhs180
			var _lhs181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(371), _dafny.ArrayLen((_1192_v7), 0))
			_ = _lhs181
			_lhs175.F10 = _rhs259
			(_lhs176).ArraySet1(_rhs260, (_lhs177).Int())
			(_lhs178).ArraySet1(_rhs261, (_lhs179).Int())
			(_lhs180).ArraySet1(_rhs262, (_lhs181).Int())
			r1 = _rhs263
			var _1252_v42 _dafny.Array
			_ = _1252_v42
			var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(13))
			_ = _nw213
			_1252_v42 = _nw213
			_1252_v42 = _1252_v42
			var _1253_v43 _dafny.Sequence
			_ = _1253_v43
			_1253_v43 = _dafny.SeqOf(_1249_v41)
			var _1254_v44 _dafny.MultiSet
			_ = _1254_v44
			_1254_v44 = _dafny.MultiSetOf((_1253_v43).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1253_v43).Cardinality()))).Uint32()).(_dafny.Array))
			_1254_v44 = (_1254_v44).Update(_1249_v41, Companion_Default___.Abs(p1))
		}
		var _1255_i5 _dafny.Int
		_ = _1255_i5
		_1255_i5 = _dafny.Zero
		{
			for (Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(536), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1194_v9.F10, (_this).F11())).Cardinality())).Cmp(_this.F7()) == 0 {
				{
					if (_1255_i5).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1255_i5 = (_1255_i5).Plus(_dafny.One)
					_1197_v12 = Companion_Default___.Fm14(_dafny.IntOfInt64(194), _dafny.IntOfUint32((_1188_v4).Cardinality()), _dafny.IntOfInt64(767), globalState)
					(_1194_v9).F10 = p0
					var _1256_v45 _dafny.MultiSet
					_ = _1256_v45
					_1256_v45 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_1188_v4, _1188_v4), _dafny.Companion_Sequence_.Concatenate(_1188_v4, _1188_v4), _dafny.Companion_Sequence_.Concatenate(_1188_v4, _dafny.UnicodeSeqOfUtf8Bytes("uohlhbu")))
					(globalState).F1 = (_1256_v45).Cardinality()
					_1197_v12 = _1197_v12
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1257_v46 _dafny.Sequence
		_ = _1257_v46
		_1257_v46 = _dafny.SeqOf(p1)
		r0 = _dafny.Companion_Sequence_.Contains(_1257_v46, p2)
		r1 = _dafny.Companion_Sequence_.Equal(_1188_v4, _1188_v4)
		r2 = _this.F7()
		return r0, r1, r2
	}
}
func (_this *C4) M4(globalState *GlobalState) {
	{
		var _1258_v0 _dafny.Map
		_ = _1258_v0
		_1258_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (false) == ((_this).F11()))
		var _1259_v1 _dafny.MultiSet
		_ = _1259_v1
		_1259_v1 = _dafny.MultiSetOf((_this).F11())
		_1258_v0 = (_1258_v0).Update((_1259_v1).IsDisjointFrom(_1259_v1), true)
		var _1260_v2 bool
		_ = _1260_v2
		_1260_v2 = true
		_1260_v2 = (_this).F11()
		var _hi13 _dafny.Int = _dafny.IntOfInt64(110)
		_ = _hi13
		for _1261_i0 := _this.F7(); _1261_i0.Cmp(_hi13) < 0; _1261_i0 = _1261_i0.Plus(_dafny.One) {
			if _1260_v2 {
				var _1262_v3 _dafny.MultiSet
				_ = _1262_v3
				_1262_v3 = _1259_v1
				var _1263_v4 _dafny.Map
				_ = _1263_v4
				_1263_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1262_v3)
				_1263_v4 = (_1263_v4).Update(!((_this).F11()) || ((_this).F11()), _1262_v3)
				var _1264_v5 _dafny.Set
				_ = _1264_v5
				_1264_v5 = _dafny.SetOf(_1260_v2, (_this).F11(), _1260_v2)
				var _1265_v6 D8
				_ = _1265_v6
				_1265_v6 = Companion_D8_.Create_DC20_((_this).F11(), (_this).F11(), _1260_v2, _1260_v2)
				var _1266_v7 _dafny.Set
				_ = _1266_v7
				_1266_v7 = _dafny.SetOf(Companion_Default___.Fm1((_1265_v6).Dtor_cf39(), globalState))
				var _1267_v8 T0
				_ = _1267_v8
				var _nw214 *C2 = New_C2_()
				_ = _nw214
				_nw214.Ctor__((_1264_v5).Cardinality(), (_1266_v7).Cardinality())
				_1267_v8 = _nw214
				var _1268_v9 _dafny.Sequence
				_ = _1268_v9
				_1268_v9 = _dafny.SeqOf(_1260_v2, (_this).F11(), false, _1260_v2, (_this).F11())
				var _1269_v10 D8
				_ = _1269_v10
				_1269_v10 = Companion_D8_.Create_DC18_(_1267_v8)
				var _1270_v11 _dafny.Array
				_ = _1270_v11
				var _nwElement0_47 T0 = _1267_v8
				_ = _nwElement0_47
				var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(29))
				_ = _nw215
				(_nw215).ArraySet1(_nwElement0_47, 0)
				(_nw215).ArraySet1(_1267_v8, 1)
				(_nw215).ArraySet1(_1267_v8, 2)
				(_nw215).ArraySet1(_1267_v8, 3)
				(_nw215).ArraySet1(_1267_v8, 4)
				(_nw215).ArraySet1(_1267_v8, 5)
				(_nw215).ArraySet1(_1267_v8, 6)
				(_nw215).ArraySet1(_1267_v8, 7)
				(_nw215).ArraySet1(_1267_v8, 8)
				(_nw215).ArraySet1((func() T0 {
					if (_this).F11() {
						return _1267_v8
					}
					return _1267_v8
				})(), 9)
				(_nw215).ArraySet1(_1267_v8, 10)
				(_nw215).ArraySet1(_1267_v8, 11)
				(_nw215).ArraySet1((func() T0 {
					if (_this).F11() {
						return _1267_v8
					}
					return _1267_v8
				})(), 12)
				(_nw215).ArraySet1(_1267_v8, 13)
				(_nw215).ArraySet1(_1267_v8, 14)
				(_nw215).ArraySet1(_1267_v8, 15)
				(_nw215).ArraySet1(_1267_v8, 16)
				(_nw215).ArraySet1(_1267_v8, 17)
				(_nw215).ArraySet1(_1267_v8, 18)
				(_nw215).ArraySet1((func() T0 {
					if !((_1268_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1260_v2, _1260_v2, _1260_v2)).Cardinality()), _dafny.IntOfUint32((_1268_v9).Cardinality()))).Uint32()).(bool)) {
						return (_1269_v10).Dtor_cf33()
					}
					return _1267_v8
				})(), 19)
				(_nw215).ArraySet1(_1267_v8, 20)
				(_nw215).ArraySet1(_1267_v8, 21)
				(_nw215).ArraySet1(_1267_v8, 22)
				(_nw215).ArraySet1(_1267_v8, 23)
				(_nw215).ArraySet1(_1267_v8, 24)
				(_nw215).ArraySet1(_1267_v8, 25)
				(_nw215).ArraySet1(_1267_v8, 26)
				(_nw215).ArraySet1(_1267_v8, 27)
				(_nw215).ArraySet1(_1267_v8, 28)
				_1270_v11 = _nw215
				var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1270_v11), 0))
				_ = _index214
				(_1270_v11).ArraySet1(_1267_v8, (_index214).Int())
				var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1270_v11), 0))
				_ = _index215
				(_1270_v11).ArraySet1((Companion_D8_.Create_DC18_(_1267_v8)).Dtor_cf33(), (_index215).Int())
				var _1271_v12 _dafny.Sequence
				_ = _1271_v12
				_1271_v12 = _dafny.SeqOf(_1267_v8.F7(), _dafny.IntOfInt64(598), (_1261_i0).Plus(_1261_i0))
				(_this).F7_set_((_dafny.Zero).Minus((_1271_v12).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm1(_1260_v2, globalState), _dafny.IntOfUint32((_1271_v12).Cardinality()))).Uint32()).(_dafny.Int)))
				var _1272_v13 _dafny.Array
				_ = _1272_v13
				var _len0_40 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_40
				var _nw216 _dafny.Array
				_ = _nw216
				if _len0_40.Cmp(_dafny.Zero) == 0 {
					_nw216 = _dafny.NewArray(_len0_40)
				} else {
					var _init40 func(_dafny.Int) _dafny.Int = func(_1273_i1 _dafny.Int) _dafny.Int {
						return (_1273_i1).Minus((_this).F6())
					}
					_ = _init40
					var _element0_40 = _init40(_dafny.Zero)
					_ = _element0_40
					_nw216 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
					(_nw216).ArraySet1(_element0_40, 0)
					var _nativeLen0_40 = (_len0_40).Int()
					_ = _nativeLen0_40
					for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
						(_nw216).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
					}
				}
				_1272_v13 = _nw216
				var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1272_v13), 0))
				_ = _index216
				(_1272_v13).ArraySet1((func() _dafny.Int {
					if _1260_v2 {
						return _dafny.IntOfInt64(276)
					}
					return _1267_v8.F7()
				})(), (_index216).Int())
				var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(819), _dafny.ArrayLen((_1272_v13), 0))
				_ = _index217
				(_1272_v13).ArraySet1((_1267_v8).F6(), (_index217).Int())
				var _1274_v14 _dafny.Map
				_ = _1274_v14
				_1274_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_T0_.CastTo_((_1270_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1270_v11), 0))).Int())), _1260_v2)
				_1274_v14 = (_1274_v14).Update(Companion_T0_.CastTo_((_1270_v11).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(661), _dafny.ArrayLen((_1270_v11), 0))).Int())), _1260_v2)
			} else {
				var _1275_v15 _dafny.Sequence
				_ = _1275_v15
				_1275_v15 = _dafny.SeqOf((_this).F11())
				var _1276_v16 _dafny.Map
				_ = _1276_v16
				_1276_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1((_this).F11(), globalState), _dafny.IntOfUint32((_1275_v15).Cardinality()))
				var _1277_v17 _dafny.Map
				_ = _1277_v17
				_1277_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1276_v16, (func() _dafny.Int {
					if (_this).F11() {
						return _1261_i0
					}
					return _1261_i0
				})())
				_1277_v17 = (_1277_v17).Update(_1276_v16, (_this.F7()).Plus(_this.F7()))
				var _1278_v18 _dafny.CodePoint
				_ = _1278_v18
				_1278_v18 = _dafny.CodePoint('e')
				var _1279_v19 _dafny.Array
				_ = _1279_v19
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(29)
				_ = _len0_41
				var _nw217 _dafny.Array
				_ = _nw217
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw217 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1280_v2 bool) func(_dafny.Int) bool {
						return func(_1281_i2 _dafny.Int) bool {
							return _1280_v2
						}
					})(_1260_v2)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw217 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw217).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw217).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1279_v19 = _nw217
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1279_v19), 0))
				_ = _index218
				(_1279_v19).ArraySet1(_1260_v2, (_index218).Int())
				var _1282_v20 _dafny.Array
				_ = _1282_v20
				var _len0_42 _dafny.Int = _dafny.IntOfInt64(24)
				_ = _len0_42
				var _nw218 _dafny.Array
				_ = _nw218
				if _len0_42.Cmp(_dafny.Zero) == 0 {
					_nw218 = _dafny.NewArray(_len0_42)
				} else {
					var _init42 func(_dafny.Int) _dafny.Int = func(_1283_i3 _dafny.Int) _dafny.Int {
						return (_1283_i3).Times(_dafny.IntOfInt64(-171))
					}
					_ = _init42
					var _element0_42 = _init42(_dafny.Zero)
					_ = _element0_42
					_nw218 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
					(_nw218).ArraySet1(_element0_42, 0)
					var _nativeLen0_42 = (_len0_42).Int()
					_ = _nativeLen0_42
					for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
						(_nw218).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
					}
				}
				_1282_v20 = _nw218
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _index219
				(_1282_v20).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F6()), (_this).F6()), (_index219).Int())
				var _1284_v21 _dafny.Sequence
				_ = _1284_v21
				_1284_v21 = _dafny.UnicodeSeqOfUtf8Bytes("mxkhh")
				var _1285_v22 _dafny.Map
				_ = _1285_v22
				_1285_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1260_v2, _1284_v21)
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1279_v19), 0))
				_ = _index220
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _index221
				var _rhs264 _dafny.CodePoint = _1278_v18
				_ = _rhs264
				var _rhs265 bool = (!(_1260_v2)) || ((_1259_v1).Contains(_1260_v2))
				_ = _rhs265
				var _rhs266 _dafny.Int = _this.F7()
				_ = _rhs266
				var _rhs267 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfInt64(-360)), _1261_i0))
				_ = _rhs267
				var _rhs268 _dafny.Map = (func() _dafny.Map {
					if Companion_Default___.Fm0((_this).F11(), func() _dafny.Map {
						var _coll38 = _dafny.NewMapBuilder()
						_ = _coll38
						for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(63), _dafny.IntOfInt64(280))); ; {
							_compr_38, _ok42 := _iter42()
							if !_ok42 {
								break
							}
							var _1286_v23 _dafny.Int
							_1286_v23 = interface{}(_compr_38).(_dafny.Int)
							if ((_dafny.IntOfInt64(63)).Cmp(_1286_v23) <= 0) && ((_1286_v23).Cmp(_dafny.IntOfInt64(280)) < 0) {
								_coll38.Add((_1286_v23).Minus(_this.F7()), _dafny.IntOfInt64(689))
							}
						}
						return _coll38.ToMap()
					}(), _this.F7(), globalState) {
						return (_1285_v22).Merge(_1285_v22)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F11()), (_this).Fm12(false, (_this).F11(), globalState))
				})()
				_ = _rhs268
				var _lhs182 _dafny.Array = _1279_v19
				_ = _lhs182
				var _lhs183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1279_v19), 0))
				_ = _lhs183
				var _lhs184 _dafny.Array = _1282_v20
				_ = _lhs184
				var _lhs185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _lhs185
				var _lhs186 *GlobalState = globalState
				_ = _lhs186
				_1278_v18 = _rhs264
				(_lhs182).ArraySet1(_rhs265, (_lhs183).Int())
				(_lhs184).ArraySet1(_rhs266, (_lhs185).Int())
				_lhs186.F1 = _rhs267
				_1285_v22 = _rhs268
				_1284_v21 = _dafny.Companion_Sequence_.Concatenate(_1284_v21, _1284_v21)
				var _1287_v24 _dafny.Map
				_ = _1287_v24
				_1287_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1282_v20, _1261_i0)
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _index222
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _index223
				var _rhs269 bool = ((_1287_v24).Cardinality()).Cmp((Companion_Default___.Fm36(_this.F7(), (_1279_v19).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1279_v19), 0))).Int()).(bool), (_this).F6(), (_this).F11(), globalState)).Cardinality()) >= 0
				_ = _rhs269
				var _rhs270 _dafny.Array = _1282_v20
				_ = _rhs270
				var _rhs271 _dafny.Int = (_this).F6()
				_ = _rhs271
				var _rhs272 _dafny.Int = (_this.F7()).Plus(_this.F7())
				_ = _rhs272
				var _lhs187 _dafny.Array = _1282_v20
				_ = _lhs187
				var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _lhs188
				var _lhs189 _dafny.Array = _1282_v20
				_ = _lhs189
				var _lhs190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(190), _dafny.ArrayLen((_1282_v20), 0))
				_ = _lhs190
				_1260_v2 = _rhs269
				_1282_v20 = _rhs270
				(_lhs187).ArraySet1(_rhs271, (_lhs188).Int())
				(_lhs189).ArraySet1(_rhs272, (_lhs190).Int())
				var _1288_v25 _dafny.Set
				_ = _1288_v25
				_1288_v25 = _dafny.SetOf(_1260_v2)
				_1260_v2 = !((_1288_v25).Contains(true))
			}
			var _1289_v26 _dafny.CodePoint
			_ = _1289_v26
			_1289_v26 = _dafny.CodePoint('j')
			var _1290_v27 _dafny.Map
			_ = _1290_v27
			_1290_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.CodePoint {
				if _1260_v2 {
					return _dafny.CodePoint('i')
				}
				return _1289_v26
			})(), (_this).F6())
			var _1291_v28 _dafny.Map
			_ = _1291_v28
			_1291_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F6())
			var _1292_v29 _dafny.Sequence
			_ = _1292_v29
			_1292_v29 = _dafny.SeqOf(_this.F7())
			var _1293_v30 _dafny.MultiSet
			_ = _1293_v30
			_1293_v30 = _dafny.MultiSetOf(_this.F7(), (_this).F6(), _1261_i0, (_1259_v1).Cardinality())
			var _1294_v31 _dafny.Set
			_ = _1294_v31
			_1294_v31 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(_this.F7())).Cardinality()), _dafny.IntOfInt64(774), (_1293_v30).Cardinality())
			var _1295_v32 _dafny.Array
			_ = _1295_v32
			var _nwElement0_48 _dafny.Int = _this.F7()
			_ = _nwElement0_48
			var _nw219 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(23))
			_ = _nw219
			(_nw219).ArraySet1(_nwElement0_48, 0)
			(_nw219).ArraySet1((_this).F6(), 1)
			(_nw219).ArraySet1((func() _dafny.Int {
				if (_1291_v28).Contains(_1260_v2) {
					return (_1291_v28).Get(_1260_v2).(_dafny.Int)
				}
				return _this.F7()
			})(), 2)
			(_nw219).ArraySet1((_this).F6(), 3)
			(_nw219).ArraySet1((_this).F6(), 4)
			(_nw219).ArraySet1(_1261_i0, 5)
			(_nw219).ArraySet1(_dafny.IntOfInt64(802), 6)
			(_nw219).ArraySet1(_dafny.IntOfUint32((_1292_v29).Cardinality()), 7)
			(_nw219).ArraySet1(_1261_i0, 8)
			(_nw219).ArraySet1(_this.F7(), 9)
			(_nw219).ArraySet1(_1261_i0, 10)
			(_nw219).ArraySet1((_this).F6(), 11)
			(_nw219).ArraySet1(_this.F7(), 12)
			(_nw219).ArraySet1((_this).F6(), 13)
			(_nw219).ArraySet1((_this).F6(), 14)
			(_nw219).ArraySet1(_this.F7(), 15)
			(_nw219).ArraySet1(_1261_i0, 16)
			(_nw219).ArraySet1(_this.F7(), 17)
			(_nw219).ArraySet1(_1261_i0, 18)
			(_nw219).ArraySet1(_1261_i0, 19)
			(_nw219).ArraySet1((_1294_v31).Cardinality(), 20)
			(_nw219).ArraySet1(_1261_i0, 21)
			(_nw219).ArraySet1((_this).F6(), 22)
			_1295_v32 = _nw219
			var _1296_v33 _dafny.MultiSet
			_ = _1296_v33
			_1296_v33 = _dafny.MultiSetOf(_1295_v32)
			var _1297_v34 _dafny.Sequence
			_ = _1297_v34
			_1297_v34 = _dafny.UnicodeSeqOfUtf8Bytes("jpevpx")
			var _1298_v35 _dafny.Array
			_ = _1298_v35
			var _nw220 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(9))
			_ = _nw220
			_1298_v35 = _nw220
			var _1299_v36 _dafny.Map
			_ = _1299_v36
			_1299_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1298_v35, _1261_i0)
			var _1300_v37 _dafny.Map
			_ = _1300_v37
			_1300_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(false, globalState), _this.F7())
			var _1301_v38 _dafny.Array
			_ = _1301_v38
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_43
			var _nw221 _dafny.Array
			_ = _nw221
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw221 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.CodePoint = func(_1302_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('f')
				}
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw221 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw221).ArraySet1CodePoint(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw221).ArraySet1CodePoint(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1301_v38 = _nw221
			var _1303_v39 _dafny.Map
			_ = _1303_v39
			_1303_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1301_v38)
			var _rhs273 _dafny.Map = Companion_Default___.Fm37(((_dafny.SetOf(_1260_v2)).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(714))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_1304_i4 _dafny.Int) _dafny.Int {
				return _this.F7()
			}))).Cardinality())), Companion_D1_.Create_DC3_(!(true), _1261_i0, _dafny.IntOfUint32((_1297_v34).Cardinality())), globalState)
			_ = _rhs273
			var _rhs274 _dafny.Int = (_dafny.Zero).Minus(_this.F7())
			_ = _rhs274
			var _rhs275 bool = (_1299_v36).Contains(_1298_v35)
			_ = _rhs275
			var _rhs276 bool = Companion_Default___.Fm0(!(!(_1260_v2)), (_1300_v37).Merge(_1300_v37), (_1303_v39).Cardinality(), globalState)
			_ = _rhs276
			var _rhs277 _dafny.MultiSet = _1296_v33
			_ = _rhs277
			var _lhs191 *C4 = _this
			_ = _lhs191
			_1290_v27 = _rhs273
			_lhs191.F7_set_(_rhs274)
			_1260_v2 = _rhs275
			_1260_v2 = _rhs276
			_1296_v33 = _rhs277
			_1260_v2 = !((_this).F11())
			var _1305_v40 _dafny.Array
			_ = _1305_v40
			var _len0_44 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_44
			var _nw222 _dafny.Array
			_ = _nw222
			if _len0_44.Cmp(_dafny.Zero) == 0 {
				_nw222 = _dafny.NewArray(_len0_44)
			} else {
				var _init44 func(_dafny.Int) D1 = func(_1306_i6 _dafny.Int) D1 {
					return Companion_D1_.Create_DC2_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_dafny.SetOf((_this).F11())), (_this).F11()))
				}
				_ = _init44
				var _element0_44 = _init44(_dafny.Zero)
				_ = _element0_44
				_nw222 = _dafny.NewArrayFromExample(_element0_44, nil, _len0_44)
				(_nw222).ArraySet1(_element0_44, 0)
				var _nativeLen0_44 = (_len0_44).Int()
				_ = _nativeLen0_44
				for _i0_44 := 1; _i0_44 < _nativeLen0_44; _i0_44++ {
					(_nw222).ArraySet1(_init44(_dafny.IntOf(_i0_44)), _i0_44)
				}
			}
			_1305_v40 = _nw222
			var _1307_v41 D6
			_ = _1307_v41
			_1307_v41 = Companion_D6_.Create_DC13_(_1305_v40)
			var _source16 D6 = _1307_v41
			_ = _source16
			if _source16.Is_DC14() {
				var _1308___mcc_h0 _dafny.Int = _source16.Get_().(D6_DC14).Cf23
				_ = _1308___mcc_h0
				var _1309___mcc_h1 _dafny.MultiSet = _source16.Get_().(D6_DC14).Cf24
				_ = _1309___mcc_h1
				var _1310___mcc_h2 _dafny.Int = _source16.Get_().(D6_DC14).Cf25
				_ = _1310___mcc_h2
				var _1311_cf25 _dafny.Int = _1310___mcc_h2
				_ = _1311_cf25
				var _1312_cf24 _dafny.MultiSet = _1309___mcc_h1
				_ = _1312_cf24
				var _1313_cf23 _dafny.Int = _1308___mcc_h0
				_ = _1313_cf23
				_1313_cf23 = Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus((_this).F6())).Plus((_1294_v31).Cardinality()), (_1313_cf23).Times(_this.F7()))
				var _1314_v42 _dafny.Sequence
				_ = _1314_v42
				_1314_v42 = _dafny.SeqOf(_1260_v2)
				var _1315_v43 _dafny.Sequence
				_ = _1315_v43
				_1315_v43 = _dafny.SeqOf((_1314_v42).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1314_v42).Cardinality()), _dafny.IntOfUint32((_1314_v42).Cardinality()))).Uint32()).(bool), true)
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1295_v32), 0))
				_ = _index224
				(_1295_v32).ArraySet1((func() _dafny.Int {
					if (_1291_v28).Contains(_1260_v2) {
						return (_1291_v28).Get(_1260_v2).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_1315_v43).Cardinality())
				})(), (_index224).Int())
				var _1316_v44 D2
				_ = _1316_v44
				_1316_v44 = Companion_D2_.Create_DC5_(_1295_v32)
				var _1317_v45 _dafny.Array
				_ = _1317_v45
				var _nwElement0_49 D2 = _1316_v44
				_ = _nwElement0_49
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(2))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_49, 0)
				(_nw223).ArraySet1(_1316_v44, 1)
				_1317_v45 = _nw223
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1298_v35), 0))
				_ = _index225
				(_1298_v35).ArraySet1(true, (_index225).Int())
				var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1298_v35), 0))
				_ = _index226
				(_1298_v35).ArraySet1(_1260_v2, (_index226).Int())
				var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1295_v32), 0))
				_ = _index227
				var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1298_v35), 0))
				_ = _index228
				var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1298_v35), 0))
				_ = _index229
				var _rhs278 _dafny.Int = Companion_Default___.Fm1((_this).F11(), globalState)
				_ = _rhs278
				var _rhs279 _dafny.Array = _1317_v45
				_ = _rhs279
				var _rhs280 _dafny.CodePoint = _1289_v26
				_ = _rhs280
				var _rhs281 bool = (_this).F11()
				_ = _rhs281
				var _rhs282 bool = (_dafny.IntOfInt64(455)).Cmp((func() _dafny.Int {
					if (_this).F11() {
						return (_dafny.SetOf(_1260_v2, (_this).F11(), _1260_v2)).Cardinality()
					}
					return _1261_i0
				})()) != 0
				_ = _rhs282
				var _lhs192 _dafny.Array = _1295_v32
				_ = _lhs192
				var _lhs193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(453), _dafny.ArrayLen((_1295_v32), 0))
				_ = _lhs193
				var _lhs194 _dafny.Array = _1298_v35
				_ = _lhs194
				var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1298_v35), 0))
				_ = _lhs195
				var _lhs196 _dafny.Array = _1298_v35
				_ = _lhs196
				var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_1298_v35), 0))
				_ = _lhs197
				(_lhs192).ArraySet1(_rhs278, (_lhs193).Int())
				_1317_v45 = _rhs279
				_1289_v26 = _rhs280
				(_lhs194).ArraySet1(_rhs281, (_lhs195).Int())
				(_lhs196).ArraySet1(_rhs282, (_lhs197).Int())
				_1260_v2 = (_1298_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(747), _dafny.ArrayLen((_1298_v35), 0))).Int()).(bool)
				var _1318_v46 _dafny.Set
				_ = _1318_v46
				_1318_v46 = _dafny.SetOf(_1260_v2, _1260_v2, false)
				var _1319_v47 D7
				_ = _1319_v47
				_1319_v47 = Companion_D7_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("dufjuu"), _1318_v46)
				var _1320_v48 T0
				_ = _1320_v48
				var _nw224 *C2 = New_C2_()
				_ = _nw224
				_nw224.Ctor__(_1311_cf25, _1311_cf25)
				_1320_v48 = _nw224
				var _1321_v49 _dafny.Map
				_ = _1321_v49
				_1321_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1320_v48, _1318_v46)
				var _pat_let_tv28 = _1321_v49
				_ = _pat_let_tv28
				var _pat_let_tv29 = _1320_v48
				_ = _pat_let_tv29
				var _pat_let_tv30 = _1321_v49
				_ = _pat_let_tv30
				var _pat_let_tv31 = _1320_v48
				_ = _pat_let_tv31
				var _pat_let_tv32 = _1318_v46
				_ = _pat_let_tv32
				var _1322_v50 _dafny.Map
				_ = _1322_v50
				_1322_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let25_0 D7) D7 {
					return func(_1323_dt__update__tmp_h0 D7) D7 {
						return func(_pat_let26_0 _dafny.Set) D7 {
							return func(_1324_dt__update_hcf32_h0 _dafny.Set) D7 {
								return Companion_D7_.Create_DC17_((_1323_dt__update__tmp_h0).Dtor_cf31(), _1324_dt__update_hcf32_h0)
							}(_pat_let26_0)
						}((func() _dafny.Set {
							if (_pat_let_tv28).Contains(_pat_let_tv29) {
								return (_pat_let_tv30).Get(_pat_let_tv31).(_dafny.Set)
							}
							return _pat_let_tv32
						})())
					}(_pat_let25_0)
				}(_1319_v47), _dafny.IntOfInt64(374))
				var _1325_v51 _dafny.Array
				_ = _1325_v51
				var _nw225 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw225
				_1325_v51 = _nw225
				(_this).F7_set_((func() _dafny.Int {
					if (_1322_v50).Contains(Companion_D7_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("txexlpnd"), _1318_v46)) {
						return (_1322_v50).Get(Companion_D7_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("txexlpnd"), _1318_v46)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.SeqOf(_1325_v51)).Cardinality())
				})())
			} else if _source16.Is_DC15() {
				var _1326___mcc_h3 bool = _source16.Get_().(D6_DC15).Cf26
				_ = _1326___mcc_h3
				var _1327___mcc_h4 _dafny.Int = _source16.Get_().(D6_DC15).Cf27
				_ = _1327___mcc_h4
				var _1328___mcc_h5 _dafny.Int = _source16.Get_().(D6_DC15).Cf28
				_ = _1328___mcc_h5
				var _1329___mcc_h6 _dafny.Int = _source16.Get_().(D6_DC15).Cf29
				_ = _1329___mcc_h6
				var _1330_cf29 _dafny.Int = _1329___mcc_h6
				_ = _1330_cf29
				var _1331_cf28 _dafny.Int = _1328___mcc_h5
				_ = _1331_cf28
				var _1332_cf27 _dafny.Int = _1327___mcc_h4
				_ = _1332_cf27
				var _1333_cf26 bool = _1326___mcc_h3
				_ = _1333_cf26
				_1333_cf26 = _1333_cf26
				var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1295_v32), 0))
				_ = _index230
				(_1295_v32).ArraySet1(_1330_cf29, (_index230).Int())
				var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1295_v32), 0))
				_ = _index231
				(_1295_v32).ArraySet1(_1330_cf29, (_index231).Int())
				var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(219), _dafny.ArrayLen((_1295_v32), 0))
				_ = _index232
				(_1295_v32).ArraySet1(_dafny.IntOfInt64(-383), (_index232).Int())
				_1333_cf26 = (_1332_cf27).Cmp(_dafny.IntOfInt64(-823)) >= 0
			} else {
				var _1334___mcc_h7 _dafny.Array = _source16.Get_().(D6_DC13).Cf22
				_ = _1334___mcc_h7
				var _1335_cf22 _dafny.Array = _1334___mcc_h7
				_ = _1335_cf22
				_1300_v37 = (_1300_v37).Update((_this).F6(), _this.F7())
				var _1336_v52 _dafny.Map
				_ = _1336_v52
				_1336_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_1260_v2, _1260_v2), _1301_v38)
				var _1337_v53 _dafny.Sequence
				_ = _1337_v53
				_1337_v53 = _dafny.SeqOf(_1260_v2)
				_1336_v52 = (_1336_v52).Update(_dafny.Companion_Sequence_.Concatenate(_1337_v53, _dafny.Companion_Sequence_.Update(_1337_v53, (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1337_v53).Cardinality()))).Uint32(), !((_this).F11()))), _1301_v38)
				_1300_v37 = (_1300_v37).Update((_dafny.Zero).Minus((_this).F6()), _1261_i0)
				_1291_v28 = (_1291_v28).Update(_1260_v2, (_this).F6())
			}
		}
		var _1338_i7 _dafny.Int
		_ = _1338_i7
		_1338_i7 = _dafny.Zero
		{
			for (_this).F11() {
				{
					if (_1338_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_1338_i7 = (_1338_i7).Plus(_dafny.One)
					(_this).F7_set_(_this.F7())
					var _1339_v54 _dafny.MultiSet
					_ = _1339_v54
					_1339_v54 = _dafny.MultiSetOf((_dafny.Zero).Minus((_1258_v0).Cardinality()))
					var _1340_v55 _dafny.Map
					_ = _1340_v55
					_1340_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), (_1339_v54).Cardinality())
					var _1341_v56 _dafny.Set
					_ = _1341_v56
					_1341_v56 = _dafny.SetOf(_1340_v55)
					if (_1341_v56).IsSubsetOf((_1341_v56).Difference(_1341_v56)) {
						var _1342_v57 _dafny.Set
						_ = _1342_v57
						_1342_v57 = _dafny.SetOf(_1260_v2)
						var _1343_v58 D0
						_ = _1343_v58
						_1343_v58 = Companion_D0_.Create_DC0_(_1342_v57)
						var _1344_v59 _dafny.Map
						_ = _1344_v59
						_1344_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1343_v58, (_this).F11())
						_1344_v59 = _1344_v59
						var _1345_v60 _dafny.Array
						_ = _1345_v60
						var _len0_45 _dafny.Int = _dafny.IntOfInt64(12)
						_ = _len0_45
						var _nw226 _dafny.Array
						_ = _nw226
						if _len0_45.Cmp(_dafny.Zero) == 0 {
							_nw226 = _dafny.NewArray(_len0_45)
						} else {
							var _init45 func(_dafny.Int) _dafny.Int = func(_1346_i8 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeModuloInt(_1346_i8, _this.F7())
							}
							_ = _init45
							var _element0_45 = _init45(_dafny.Zero)
							_ = _element0_45
							_nw226 = _dafny.NewArrayFromExample(_element0_45, nil, _len0_45)
							(_nw226).ArraySet1(_element0_45, 0)
							var _nativeLen0_45 = (_len0_45).Int()
							_ = _nativeLen0_45
							for _i0_45 := 1; _i0_45 < _nativeLen0_45; _i0_45++ {
								(_nw226).ArraySet1(_init45(_dafny.IntOf(_i0_45)), _i0_45)
							}
						}
						_1345_v60 = _nw226
						var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _index233
						(_1345_v60).ArraySet1((_dafny.Zero).Minus(_this.F7()), (_index233).Int())
						var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _index234
						(_1345_v60).ArraySet1((_this).F6(), (_index234).Int())
						_1260_v2 = (_this).F11()
						_1260_v2 = (true) == (_1260_v2)
						var _1347_v61 _dafny.Array
						_ = _1347_v61
						var _nw227 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
						_ = _nw227
						_1347_v61 = _nw227
						var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1347_v61), 0))
						_ = _index235
						(_1347_v61).ArraySet1(!(_1260_v2), (_index235).Int())
						var _1348_v62 T0
						_ = _1348_v62
						var _nw228 *C2 = New_C2_()
						_ = _nw228
						_nw228.Ctor__(_this.F7(), (_this).F6())
						_1348_v62 = _nw228
						var _1349_v63 _dafny.Map
						_ = _1349_v63
						_1349_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1348_v62, (_this).F6())
						var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _index236
						var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1347_v61), 0))
						_ = _index237
						var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _index238
						var _rhs283 bool = (_1259_v1).IsProperSubsetOf(_1259_v1)
						_ = _rhs283
						var _rhs284 _dafny.Int = (func() _dafny.Int {
							if (_1349_v63).Contains(_1348_v62) {
								return (_1349_v63).Get(_1348_v62).(_dafny.Int)
							}
							return (_dafny.Zero).Minus((_dafny.Zero).Minus(((_1348_v62).F6()).Times(_1348_v62.F7())))
						})()
						_ = _rhs284
						var _rhs285 bool = (_this).F11()
						_ = _rhs285
						var _rhs286 _dafny.Int = _this.F7()
						_ = _rhs286
						var _lhs198 _dafny.Array = _1345_v60
						_ = _lhs198
						var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _lhs199
						var _lhs200 _dafny.Array = _1347_v61
						_ = _lhs200
						var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(899), _dafny.ArrayLen((_1347_v61), 0))
						_ = _lhs201
						var _lhs202 _dafny.Array = _1345_v60
						_ = _lhs202
						var _lhs203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_1345_v60), 0))
						_ = _lhs203
						_1260_v2 = _rhs283
						(_lhs198).ArraySet1(_rhs284, (_lhs199).Int())
						(_lhs200).ArraySet1(_rhs285, (_lhs201).Int())
						(_lhs202).ArraySet1(_rhs286, (_lhs203).Int())
					} else {
						_1260_v2 = (_this).F11()
						var _1350_v64 _dafny.Array
						_ = _1350_v64
						var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
						_ = _nw229
						_1350_v64 = _nw229
						var _1351_v65 _dafny.Array
						_ = _1351_v65
						var _len0_46 _dafny.Int = _dafny.IntOfInt64(22)
						_ = _len0_46
						var _nw230 _dafny.Array
						_ = _nw230
						if _len0_46.Cmp(_dafny.Zero) == 0 {
							_nw230 = _dafny.NewArray(_len0_46)
						} else {
							var _init46 func(_dafny.Int) bool = (func(_1352_v2 bool) func(_dafny.Int) bool {
								return func(_1353_i9 _dafny.Int) bool {
									return _1352_v2
								}
							})(_1260_v2)
							_ = _init46
							var _element0_46 = _init46(_dafny.Zero)
							_ = _element0_46
							_nw230 = _dafny.NewArrayFromExample(_element0_46, nil, _len0_46)
							(_nw230).ArraySet1(_element0_46, 0)
							var _nativeLen0_46 = (_len0_46).Int()
							_ = _nativeLen0_46
							for _i0_46 := 1; _i0_46 < _nativeLen0_46; _i0_46++ {
								(_nw230).ArraySet1(_init46(_dafny.IntOf(_i0_46)), _i0_46)
							}
						}
						_1351_v65 = _nw230
						var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1350_v64), 0))
						_ = _index239
						(_1350_v64).ArraySet1(_dafny.SetOf(_1351_v65), (_index239).Int())
						var _1354_v66 _dafny.Set
						_ = _1354_v66
						_1354_v66 = _dafny.SetOf(_1351_v65, _1351_v65)
						var _1355_v67 D11
						_ = _1355_v67
						_1355_v67 = Companion_D11_.Create_DC27_(_1354_v66)
						var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1350_v64), 0))
						_ = _index240
						(_1350_v64).ArraySet1((_1355_v67).Dtor_cf53(), (_index240).Int())
						var _1356_v68 _dafny.Array
						_ = _1356_v68
						var _nw231 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
						_ = _nw231
						_1356_v68 = _nw231
						var _1357_v69 _dafny.Sequence
						_ = _1357_v69
						_1357_v69 = _dafny.SeqOf(_1258_v0)
						var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1356_v68), 0))
						_ = _index241
						(_1356_v68).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1357_v69, (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1357_v69).Cardinality()))).Uint32(), _1258_v0)).Cardinality()), (_index241).Int())
						var _1358_v70 _dafny.Sequence
						_ = _1358_v70
						_1358_v70 = _dafny.UnicodeSeqOfUtf8Bytes("um")
						var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(115), _dafny.ArrayLen((_1356_v68), 0))
						_ = _index242
						(_1356_v68).ArraySet1(Companion_Default___.Fm1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ut"), _1358_v70), globalState), (_index242).Int())
						var _1359_v71 _dafny.Sequence
						_ = _1359_v71
						_1359_v71 = _dafny.SeqOf(_this.F7())
						_1359_v71 = _1359_v71
						_1260_v2 = (_this).F11()
					}
					var _1360_v73 _dafny.Sequence
					_ = _1360_v73
					_1360_v73 = _dafny.SeqOf((_this).F6(), _dafny.IntOfInt64(-596))
					var _1361_v74 _dafny.Sequence
					_ = _1361_v74
					_1361_v74 = _dafny.UnicodeSeqOfUtf8Bytes("utk")
					var _1362_v75 *C0
					_ = _1362_v75
					var _nw232 *C0 = New_C0_()
					_ = _nw232
					_nw232.Ctor__(Companion_Default___.Fm0(Companion_Default___.Fm0(true, func() _dafny.Map {
						var _coll39 = _dafny.NewMapBuilder()
						_ = _coll39
						for _iter43 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-881), _dafny.IntOfInt64(-796))); ; {
							_compr_39, _ok43 := _iter43()
							if !_ok43 {
								break
							}
							var _1363_v72 _dafny.Int
							_1363_v72 = interface{}(_compr_39).(_dafny.Int)
							if ((_dafny.IntOfInt64(-881)).Cmp(_1363_v72) <= 0) && ((_1363_v72).Cmp(_dafny.IntOfInt64(-796)) < 0) {
								_coll39.Add((_1363_v72).Times((_1259_v1).Cardinality()), _this.F7())
							}
						}
						return _coll39.ToMap()
					}(), (_1360_v73).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1361_v74).Cardinality()), _dafny.IntOfUint32((_1360_v73).Cardinality()))).Uint32()).(_dafny.Int), globalState), _1340_v55, (_dafny.Zero).Minus((_this).F6()), globalState))
					_1362_v75 = _nw232
					(_this).F7_set_((_this).F6())
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _1364_i10 _dafny.Int
		_ = _1364_i10
		_1364_i10 = _dafny.Zero
		{
			for false {
				{
					if (_1364_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_1364_i10 = (_1364_i10).Plus(_dafny.One)
					var _1365_v76 _dafny.Array
					_ = _1365_v76
					var _nw233 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
					_ = _nw233
					_1365_v76 = _nw233
					var _1366_v77 _dafny.Sequence
					_ = _1366_v77
					_1366_v77 = _dafny.UnicodeSeqOfUtf8Bytes("kdmrd")
					var _1367_v78 _dafny.Sequence
					_ = _1367_v78
					_1367_v78 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("qpdetefa"))
					var _1368_v79 _dafny.CodePoint
					_ = _1368_v79
					_1368_v79 = _dafny.CodePoint('y')
					var _1369_v80 D12
					_ = _1369_v80
					_1369_v80 = Companion_D12_.Create_DC30_(_1367_v78)
					var _1370_v81 _dafny.Array
					_ = _1370_v81
					var _nw234 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(24))
					_ = _nw234
					_1370_v81 = _nw234
					var _1371_v82 _dafny.Sequence
					_ = _1371_v82
					_1371_v82 = _dafny.SeqOf(_1370_v81)
					var _1372_v83 _dafny.Sequence
					_ = _1372_v83
					_1372_v83 = _dafny.SeqOf(_this.F7(), _this.F7(), (_this).F6())
					var _1373_v84 _dafny.Sequence
					_ = _1373_v84
					_1373_v84 = _dafny.SeqOf(_1260_v2, (_this).F11())
					var _1374_v85 _dafny.Array
					_ = _1374_v85
					var _nwElement0_50 _dafny.Sequence = _dafny.SeqOf(_1366_v77)
					_ = _nwElement0_50
					var _nw235 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(23))
					_ = _nw235
					(_nw235).ArraySet1(_nwElement0_50, 0)
					(_nw235).ArraySet1(_1367_v78, 1)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Update(_1367_v78, (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F11())).Cardinality(), _dafny.IntOfUint32((_1367_v78).Cardinality()))).Uint32(), _1366_v77), 2)
					(_nw235).ArraySet1(_1367_v78, 3)
					(_nw235).ArraySet1(_dafny.SeqOf(_dafny.Companion_Sequence_.Update(_1366_v77, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1366_v77).Cardinality()))).Uint32(), _1368_v79), _1366_v77), 4)
					(_nw235).ArraySet1(_1367_v78, 5)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1366_v77), _dafny.Companion_Sequence_.Update(_1367_v78, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1367_v78).Cardinality()))).Uint32(), _1366_v77)), 6)
					(_nw235).ArraySet1(_1367_v78, 7)
					(_nw235).ArraySet1(_1367_v78, 8)
					(_nw235).ArraySet1(_dafny.SeqOf(_1366_v77, _dafny.UnicodeSeqOfUtf8Bytes("fdn"), _1366_v77, _1366_v77, _1366_v77), 9)
					(_nw235).ArraySet1(_1367_v78, 10)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1367_v78, _1367_v78), (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1367_v78, _1367_v78)).Cardinality()))).Uint32(), _dafny.UnicodeSeqOfUtf8Bytes("smsluxg")), 11)
					(_nw235).ArraySet1(_1367_v78, 12)
					(_nw235).ArraySet1(_1367_v78, 13)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(196))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg61 _dafny.Int) interface{} {
							return coer61(arg61)
						}
					}((func(_1375_v77 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1376_i11 _dafny.Int) _dafny.Sequence {
							return _1375_v77
						}
					})(_1366_v77))), _1367_v78), 14)
					(_nw235).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(148))).Uint32(), func(coer62 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg62 _dafny.Int) interface{} {
							return coer62(arg62)
						}
					}((func(_1377_v77 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1378_i12 _dafny.Int) _dafny.Sequence {
							return _1377_v77
						}
					})(_1366_v77))), 15)
					(_nw235).ArraySet1(_1367_v78, 16)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Update((_1369_v80).Dtor_cf59(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1371_v82).Cardinality()), _dafny.IntOfUint32(((_1369_v80).Dtor_cf59()).Cardinality()))).Uint32(), Companion_Default___.Fm2(_1372_v83, _this.F7(), _this.F7(), (_this).F11(), globalState)), 17)
					(_nw235).ArraySet1(_1367_v78, 18)
					(_nw235).ArraySet1(Companion_Default___.Fm38(globalState), 19)
					(_nw235).ArraySet1(_1367_v78, 20)
					(_nw235).ArraySet1(_dafny.Companion_Sequence_.Update(_1367_v78, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1373_v84).Cardinality()), _dafny.IntOfUint32((_1367_v78).Cardinality()))).Uint32(), _1366_v77), 21)
					(_nw235).ArraySet1(_1367_v78, 22)
					_1374_v85 = _nw235
					var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1374_v85), 0))
					_ = _index243
					(_1374_v85).ArraySet1(_1367_v78, (_index243).Int())
					var _1379_v86 _dafny.Sequence
					_ = _1379_v86
					_1379_v86 = _dafny.SeqOf(_dafny.SeqOf((Companion_D10_.Create_DC26_(_1260_v2)).Dtor_cf52()))
					var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1374_v85), 0))
					_ = _index244
					var _rhs287 bool = _1260_v2
					_ = _rhs287
					var _rhs288 _dafny.Int = (func() _dafny.Int {
						if _1260_v2 {
							return (_1372_v83).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1372_v83).Cardinality()))).Uint32()).(_dafny.Int)
						}
						return Companion_Default___.Fm1(true, globalState)
					})()
					_ = _rhs288
					var _rhs289 _dafny.Array = (func() _dafny.Array {
						if (_this.F7()).Cmp(_this.F7()) < 0 {
							return _1365_v76
						}
						return _1365_v76
					})()
					_ = _rhs289
					var _rhs290 bool = !(((_this.F7()).Cmp(_this.F7()) > 0) && (_dafny.Companion_Sequence_.IsPrefixOf(_1379_v86, _dafny.SeqOf(_1373_v84))))
					_ = _rhs290
					var _rhs291 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1367_v78, _dafny.Companion_Sequence_.Update(_1367_v78, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1367_v78).Cardinality()))).Uint32(), _1366_v77))
					_ = _rhs291
					var _lhs204 *GlobalState = globalState
					_ = _lhs204
					var _lhs205 _dafny.Array = _1374_v85
					_ = _lhs205
					var _lhs206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen((_1374_v85), 0))
					_ = _lhs206
					_1260_v2 = _rhs287
					_lhs204.F1 = _rhs288
					_1365_v76 = _rhs289
					_1260_v2 = _rhs290
					(_lhs205).ArraySet1(_rhs291, (_lhs206).Int())
					_1260_v2 = (_this.F7()).Cmp((_this).F6()) <= 0
					var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_1365_v76), 0))
					_ = _index245
					(_1365_v76).ArraySet1((_this).F11(), (_index245).Int())
					var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(187), _dafny.ArrayLen((_1365_v76), 0))
					_ = _index246
					(_1365_v76).ArraySet1(((_this).F6()).Cmp(Companion_Default___.SafeDivisionInt(_this.F7(), _dafny.IntOfUint32((_1372_v83).Cardinality()))) >= 0, (_index246).Int())
					(_this).F7_set_(_this.F7())
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _1380_v87 _dafny.CodePoint
		_ = _1380_v87
		_1380_v87 = _dafny.CodePoint('k')
		var _1381_v88 D10
		_ = _1381_v88
		_1381_v88 = Companion_D10_.Create_DC24_(_1380_v87)
		_1380_v87 = (_1381_v88).Dtor_cf51()
	}
}
func (_this *C4) M5(p0 _dafny.Int, globalState *GlobalState) {
	{
		var _1382_i0 _dafny.Int
		_ = _1382_i0
		_1382_i0 = _dafny.Zero
		{
			for ((_this).F11()) && ((_this).F11()) {
				{
					if (_1382_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_1382_i0 = (_1382_i0).Plus(_dafny.One)
					var _1383_v0 _dafny.Array
					_ = _1383_v0
					var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(15))
					_ = _nw236
					_1383_v0 = _nw236
					var _1384_v1 _dafny.CodePoint
					_ = _1384_v1
					_1384_v1 = _dafny.CodePoint('j')
					var _1385_v2 _dafny.Sequence
					_ = _1385_v2
					_1385_v2 = _dafny.SeqOf(_1384_v1)
					var _1386_v3 _dafny.Array
					_ = _1386_v3
					var _nw237 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
					_ = _nw237
					_1386_v3 = _nw237
					var _1387_v4 _dafny.Map
					_ = _1387_v4
					_1387_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1385_v2).Cardinality()), _1386_v3)
					var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_1383_v0), 0))
					_ = _index247
					(_1383_v0).ArraySet1(_1387_v4, (_index247).Int())
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(459), _dafny.ArrayLen((_1383_v0), 0))
					_ = _index248
					(_1383_v0).ArraySet1(_1387_v4, (_index248).Int())
					var _1388_v5 _dafny.Sequence
					_ = _1388_v5
					_1388_v5 = _dafny.SeqOf(_1385_v2)
					var _1389_v6 _dafny.Map
					_ = _1389_v6
					_1389_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F11())
					var _1390_v7 _dafny.Sequence
					_ = _1390_v7
					_1390_v7 = _dafny.SeqOf(_this.F7(), p0, (_1389_v6).Cardinality())
					var _1391_v8 _dafny.Sequence
					_ = _1391_v8
					_1391_v8 = _dafny.SeqOf(_dafny.IntOfUint32((_1390_v7).Cardinality()))
					var _1392_v9 _dafny.Map
					_ = _1392_v9
					_1392_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1388_v5, Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1391_v8).Cardinality()), Companion_Default___.Fm1((_this).F11(), globalState)))
					(_this).F7_set_((func() _dafny.Int {
						if (_1392_v9).Contains(_1388_v5) {
							return (_1392_v9).Get(_1388_v5).(_dafny.Int)
						}
						return Companion_Default___.Fm1((_this).F11(), globalState)
					})())
					if !((_this.F7()).Cmp(_this.F7()) != 0) || ((_this).F11()) {
						var _1393_v10 bool
						_ = _1393_v10
						_1393_v10 = true
						var _1394_v11 _dafny.Map
						_ = _1394_v11
						_1394_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1385_v2).Cardinality()), (_this).F6())
						var _1395_v12 _dafny.Map
						_ = _1395_v12
						_1395_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1384_v1, (func() _dafny.Int {
							if (_1394_v11).Contains(_this.F7()) {
								return (_1394_v11).Get(_this.F7()).(_dafny.Int)
							}
							return _this.F7()
						})())
						var _1396_v13 D7
						_ = _1396_v13
						_1396_v13 = Companion_D7_.Create_DC16_(_1395_v12)
						var _pat_let_tv33 = _1395_v12
						_ = _pat_let_tv33
						var _pat_let_tv34 = _1395_v12
						_ = _pat_let_tv34
						var _rhs292 _dafny.Sequence = _1385_v2
						_ = _rhs292
						var _rhs293 bool = _1393_v10
						_ = _rhs293
						var _rhs294 _dafny.Int = Companion_Default___.Fm1((_this).F11(), globalState)
						_ = _rhs294
						var _rhs295 D7 = func(_pat_let27_0 D7) D7 {
							return func(_1397_dt__update__tmp_h0 D7) D7 {
								return func(_pat_let28_0 _dafny.Map) D7 {
									return func(_1398_dt__update_hcf30_h0 _dafny.Map) D7 {
										return Companion_D7_.Create_DC16_(_1398_dt__update_hcf30_h0)
									}(_pat_let28_0)
								}((_pat_let_tv33).Merge(_pat_let_tv34))
							}(_pat_let27_0)
						}(_1396_v13)
						_ = _rhs295
						var _lhs207 *GlobalState = globalState
						_ = _lhs207
						_1385_v2 = _rhs292
						_1393_v10 = _rhs293
						_lhs207.F1 = _rhs294
						_1396_v13 = _rhs295
						(globalState).F1 = _this.F7()
						(globalState).F1 = (func() _dafny.Int {
							if true {
								return _dafny.IntOfInt64(-97)
							}
							return _this.F7()
						})()
						var _1399_v14 _dafny.Map
						_ = _1399_v14
						_1399_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F11())
						_1399_v14 = _1399_v14
						var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1386_v3), 0))
						_ = _index249
						(_1386_v3).ArraySet1(_1393_v10, (_index249).Int())
						var _1400_v15 _dafny.Set
						_ = _1400_v15
						_1400_v15 = _dafny.SetOf((func() bool {
							if (_1389_v6).Contains(true) {
								return (_1389_v6).Get(true).(bool)
							}
							return _1393_v10
						})(), (_this).F11(), _1393_v10)
						var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1386_v3), 0))
						_ = _index250
						var _rhs296 D7 = _1396_v13
						_ = _rhs296
						var _rhs297 bool = (_1400_v15).Contains(false)
						_ = _rhs297
						var _rhs298 _dafny.Int = ((_this).F6()).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1385_v2).Cardinality())))
						_ = _rhs298
						var _lhs208 _dafny.Array = _1386_v3
						_ = _lhs208
						var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(464), _dafny.ArrayLen((_1386_v3), 0))
						_ = _lhs209
						var _lhs210 *GlobalState = globalState
						_ = _lhs210
						_1396_v13 = _rhs296
						(_lhs208).ArraySet1(_rhs297, (_lhs209).Int())
						_lhs210.F1 = _rhs298
					} else {
						var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))
						_ = _index251
						(_1386_v3).ArraySet1((_this).F11(), (_index251).Int())
						var _1401_v16 D8
						_ = _1401_v16
						_1401_v16 = Companion_D8_.Create_DC20_((_this).F11(), (_this).F11(), (_this).F11(), (_this).F11())
						var _1402_v17 _dafny.Sequence
						_ = _1402_v17
						_1402_v17 = _dafny.SeqOf((func(_pat_let29_0 D8) D8 {
							return func(_1403_dt__update__tmp_h1 D8) D8 {
								return func(_pat_let30_0 bool) D8 {
									return func(_1404_dt__update_hcf39_h0 bool) D8 {
										return Companion_D8_.Create_DC20_(_1404_dt__update_hcf39_h0, (_1403_dt__update__tmp_h1).Dtor_cf40(), (_1403_dt__update__tmp_h1).Dtor_cf41(), (_1403_dt__update__tmp_h1).Dtor_cf42())
									}(_pat_let30_0)
								}((_this).F11())
							}(_pat_let29_0)
						}(_1401_v16)).Dtor_cf39(), (_this).F11())
						var _1405_v18 _dafny.Map
						_ = _1405_v18
						_1405_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1386_v3)
						var _1406_v19 _dafny.Map
						_ = _1406_v19
						_1406_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), Companion_Default___.Fm35(_dafny.IntOfUint32((_1402_v17).Cardinality()), (_1405_v18).Cardinality(), globalState)), (_this).F11())
						var _1407_v20 _dafny.Map
						_ = _1407_v20
						_1407_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), _1390_v7)
						var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))
						_ = _index252
						(_1386_v3).ArraySet1((func() bool {
							if (_1406_v19).Contains(_1407_v20) {
								return (_1406_v19).Get(_1407_v20).(bool)
							}
							return false
						})(), (_index252).Int())
						var _1408_v21 _dafny.Map
						_ = _1408_v21
						_1408_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ctwrgg"), _this.F7())
						var _1409_v22 _dafny.Map
						_ = _1409_v22
						_1409_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1408_v21).Cardinality(), (_1386_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))).Int()).(bool))
						_1409_v22 = (_1409_v22).Update(p0, (_1386_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))).Int()).(bool))
						var _1410_v23 _dafny.Array
						_ = _1410_v23
						var _nw238 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(19))
						_ = _nw238
						_1410_v23 = _nw238
						var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1410_v23), 0))
						_ = _index253
						(_1410_v23).ArraySet1(Companion_D6_.Create_DC15_((_this).F11(), (_this).F6(), (_this).F6(), p0), (_index253).Int())
						var _1411_v24 _dafny.Array
						_ = _1411_v24
						var _nw239 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
						_ = _nw239
						_1411_v24 = _nw239
						var _1412_v25 D6
						_ = _1412_v25
						_1412_v25 = Companion_D6_.Create_DC15_((_1386_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))).Int()).(bool), Companion_Default___.Fm1(!((_1386_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))).Int()).(bool)), globalState), _this.F7(), p0)
						var _pat_let_tv35 = p0
						_ = _pat_let_tv35
						var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1410_v23), 0))
						_ = _index254
						var _rhs299 D6 = func(_pat_let31_0 D6) D6 {
							return func(_1413_dt__update__tmp_h2 D6) D6 {
								return func(_pat_let32_0 _dafny.Int) D6 {
									return func(_1414_dt__update_hcf29_h0 _dafny.Int) D6 {
										return func(_pat_let33_0 _dafny.Int) D6 {
											return func(_1415_dt__update_hcf27_h0 _dafny.Int) D6 {
												return Companion_D6_.Create_DC15_((_1413_dt__update__tmp_h2).Dtor_cf26(), _1415_dt__update_hcf27_h0, (_1413_dt__update__tmp_h2).Dtor_cf28(), _1414_dt__update_hcf29_h0)
											}(_pat_let33_0)
										}(_pat_let_tv35)
									}(_pat_let32_0)
								}((_this).F6())
							}(_pat_let31_0)
						}(_1412_v25)
						_ = _rhs299
						var _rhs300 _dafny.Array = _1411_v24
						_ = _rhs300
						var _lhs211 _dafny.Array = _1410_v23
						_ = _lhs211
						var _lhs212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_1410_v23), 0))
						_ = _lhs212
						(_lhs211).ArraySet1(_rhs299, (_lhs212).Int())
						_1411_v24 = _rhs300
						var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1386_v3), 0))
						_ = _index255
						(_1386_v3).ArraySet1((_this.F7()).Cmp(p0) == 0, (_index255).Int())
						_1390_v7 = _dafny.Companion_Sequence_.Concatenate(_1391_v8, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(84))).Uint32(), func(coer63 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg63 _dafny.Int) interface{} {
								return coer63(arg63)
							}
						}(func(_1416_i1 _dafny.Int) _dafny.Int {
							return _dafny.IntOfInt64(912)
						})))
					}
					var _1417_v26 _dafny.MultiSet
					_ = _1417_v26
					_1417_v26 = _dafny.MultiSetOf(false, (_this).F11())
					var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1386_v3), 0))
					_ = _index256
					(_1386_v3).ArraySet1((_1417_v26).IsSubsetOf(_1417_v26), (_index256).Int())
					var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(897), _dafny.ArrayLen((_1386_v3), 0))
					_ = _index257
					(_1386_v3).ArraySet1((_this).F11(), (_index257).Int())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
		var _1418_v27 _dafny.Sequence
		_ = _1418_v27
		_1418_v27 = _dafny.UnicodeSeqOfUtf8Bytes("jnumj")
		var _1419_v28 _dafny.CodePoint
		_ = _1419_v28
		_1419_v28 = _dafny.CodePoint('r')
		var _1420_v29 _dafny.Sequence
		_ = _1420_v29
		_1420_v29 = _dafny.SeqOf(_this.F7(), _this.F7())
		var _1421_v30 _dafny.MultiSet
		_ = _1421_v30
		_1421_v30 = _dafny.MultiSetOf(_this.F7(), _this.F7())
		var _1422_v31 _dafny.Array
		_ = _1422_v31
		var _nwElement0_51 _dafny.Sequence = _1418_v27
		_ = _nwElement0_51
		var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(18))
		_ = _nw240
		(_nw240).ArraySet1(_nwElement0_51, 0)
		(_nw240).ArraySet1(_1418_v27, 1)
		(_nw240).ArraySet1((func() _dafny.Sequence {
			if (_this).F11() {
				return _dafny.UnicodeSeqOfUtf8Bytes("jeil")
			}
			return _1418_v27
		})(), 2)
		(_nw240).ArraySet1(_1418_v27, 3)
		(_nw240).ArraySet1(_dafny.Companion_Sequence_.Update(_1418_v27, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1418_v27).Cardinality()))).Uint32(), _1419_v28), 4)
		(_nw240).ArraySet1(_1418_v27, 5)
		(_nw240).ArraySet1(_1418_v27, 6)
		(_nw240).ArraySet1(_1418_v27, 7)
		(_nw240).ArraySet1(_1418_v27, 8)
		(_nw240).ArraySet1(_1418_v27, 9)
		(_nw240).ArraySet1(_1418_v27, 10)
		(_nw240).ArraySet1(_1418_v27, 11)
		(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_1420_v29, (func() _dafny.Int {
			if (_1421_v30).Contains(p0) {
				return (_1421_v30).Multiplicity(p0)
			}
			return (_this).F6()
		})(), (_this).F6(), (_this).F11(), globalState), _1418_v27), 12)
		(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1418_v27, _1418_v27), 13)
		(_nw240).ArraySet1(_1418_v27, 14)
		(_nw240).ArraySet1(_1418_v27, 15)
		(_nw240).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1418_v27, _1418_v27), 16)
		(_nw240).ArraySet1(_1418_v27, 17)
		_1422_v31 = _nw240
		var _1423_v32 _dafny.Map
		_ = _1423_v32
		_1423_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _dafny.IntOfUint32((_1418_v27).Cardinality()))
		var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1422_v31), 0))
		_ = _index258
		(_1422_v31).ArraySet1((_this).Fm12(Companion_Default___.Fm0((_this).F11(), _1423_v32, (_dafny.Zero).Minus(p0), globalState), (_this).F11(), globalState), (_index258).Int())
		var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1422_v31), 0))
		_ = _index259
		(_1422_v31).ArraySet1(_1418_v27, (_index259).Int())
		var _1424_v33 _dafny.Set
		_ = _1424_v33
		_1424_v33 = _dafny.SetOf(p0, (_1420_v29).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-56), _dafny.IntOfUint32((_1420_v29).Cardinality()))).Uint32()).(_dafny.Int))
		var _1425_v34 _dafny.Map
		_ = _1425_v34
		_1425_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), Companion_D12_.Create_DC31_(_dafny.IntOfInt64(282)))
		var _1426_v35 _dafny.MultiSet
		_ = _1426_v35
		_1426_v35 = _dafny.MultiSetOf(_1425_v34, _1425_v34)
		var _1427_v36 _dafny.Map
		_ = _1427_v36
		_1427_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1424_v33).Cardinality(), _1426_v35)
		_1427_v36 = (_1427_v36).Update((func() _dafny.Int {
			if (_this).F11() {
				return (_this).F6()
			}
			return (_this).F6()
		})(), _1426_v35)
		var _1428_v37 bool
		_ = _1428_v37
		_1428_v37 = true
		var _1429_v38 _dafny.Sequence
		_ = _1429_v38
		_1429_v38 = _dafny.SeqOf(true)
		var _1430_v39 _dafny.Map
		_ = _1430_v39
		_1430_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1428_v37, _dafny.IntOfUint32(((_1422_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1422_v31), 0))).Int()).(_dafny.Sequence)).Cardinality()))
		_1428_v37 = !((_dafny.SetOf(_this.F7(), (_this).F6(), _dafny.IntOfUint32((_1429_v38).Cardinality()), (_this).F6(), (_1430_v39).Cardinality())).Difference(_1424_v33)).Equals(_dafny.SetOf(p0, p0, (_dafny.Zero).Minus(_dafny.IntOfInt64(-177)), _dafny.IntOfInt64(-522)))
		var _1431_v40 D10
		_ = _1431_v40
		_1431_v40 = Companion_D10_.Create_DC24_(_1419_v28)
		var _source17 D10 = _1431_v40
		_ = _source17
		if _source17.Is_DC25() {
			var _1432_v41 _dafny.Set
			_ = _1432_v41
			_1432_v41 = _dafny.SetOf(_dafny.Companion_Sequence_.Equal(_1429_v38, _1429_v38), _1428_v37)
			_1432_v41 = _1432_v41
			var _1433_v42 *C3
			_ = _1433_v42
			var _nw241 *C3 = New_C3_()
			_ = _nw241
			_nw241.Ctor__(_dafny.IntOfUint32((_1418_v27).Cardinality()), _dafny.IntOfInt64(-326))
			_1433_v42 = _nw241
			var _1434_v43 _dafny.Array
			_ = _1434_v43
			var _nw242 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw242
			_1434_v43 = _nw242
			var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))
			_ = _index260
			(_1434_v43).ArraySet1((_dafny.Zero).Minus(p0), (_index260).Int())
			var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))
			_ = _index261
			(_1434_v43).ArraySet1(_dafny.IntOfUint32((_1420_v29).Cardinality()), (_index261).Int())
			if (p0).Cmp((p0).Minus(_dafny.IntOfInt64(-984))) >= 0 {
				var _1435_v44 _dafny.Array
				_ = _1435_v44
				var _nw243 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw243
				_1435_v44 = _nw243
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))
				_ = _index262
				(_1435_v44).ArraySet1(Companion_Default___.Fm0(Companion_Default___.Fm0(true, _1423_v32, (_this).F6(), globalState), _1423_v32, (_1432_v41).Cardinality(), globalState), (_index262).Int())
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))
				_ = _index263
				(_1435_v44).ArraySet1((_this).F11(), (_index263).Int())
				var _1436_v45 _dafny.Array
				_ = _1436_v45
				var _nw244 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(12))
				_ = _nw244
				_1436_v45 = _nw244
				var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_1436_v45), 0))
				_ = _index264
				(_1436_v45).ArraySet1CodePoint(_1419_v28, (_index264).Int())
				var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))
				_ = _index265
				var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_1436_v45), 0))
				_ = _index266
				var _rhs301 bool = (_1435_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))).Int()).(bool)
				_ = _rhs301
				var _rhs302 _dafny.Int = p0
				_ = _rhs302
				var _rhs303 bool = (_1435_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))).Int()).(bool)
				_ = _rhs303
				var _rhs304 _dafny.CodePoint = _1419_v28
				_ = _rhs304
				var _rhs305 bool = (_this).F11()
				_ = _rhs305
				var _lhs213 _dafny.Array = _1435_v44
				_ = _lhs213
				var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))
				_ = _lhs214
				var _lhs215 *GlobalState = globalState
				_ = _lhs215
				var _lhs216 _dafny.Array = _1436_v45
				_ = _lhs216
				var _lhs217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(652), _dafny.ArrayLen((_1436_v45), 0))
				_ = _lhs217
				(_lhs213).ArraySet1(_rhs301, (_lhs214).Int())
				_lhs215.F1 = _rhs302
				_1428_v37 = _rhs303
				(_lhs216).ArraySet1CodePoint(_rhs304, (_lhs217).Int())
				_1428_v37 = _rhs305
				var _1437_v46 _dafny.Map
				_ = _1437_v46
				_1437_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, true)
				var _1438_v47 D0
				_ = _1438_v47
				_1438_v47 = Companion_D0_.Create_DC0_(_1432_v41)
				var _1439_v48 _dafny.Map
				_ = _1439_v48
				_1439_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1438_v47, !((_1435_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))).Int()).(bool)))
				var _1440_v49 D1
				_ = _1440_v49
				_1440_v49 = Companion_D1_.Create_DC2_(_1439_v48)
				var _pat_let_tv36 = _1439_v48
				_ = _pat_let_tv36
				var _1441_v50 _dafny.Set
				_ = _1441_v50
				_1441_v50 = _dafny.SetOf(_1440_v49, func(_pat_let34_0 D1) D1 {
					return func(_1442_dt__update__tmp_h3 D1) D1 {
						return func(_pat_let35_0 _dafny.Map) D1 {
							return func(_1443_dt__update_hcf4_h0 _dafny.Map) D1 {
								return Companion_D1_.Create_DC2_(_1443_dt__update_hcf4_h0)
							}(_pat_let35_0)
						}(_pat_let_tv36)
					}(_pat_let34_0)
				}(Companion_D1_.Create_DC2_(_1439_v48)), _1440_v49, _1440_v49)
				var _1444_v51 bool
				_ = _1444_v51
				var _out26 bool
				_ = _out26
				_out26 = (_1433_v42).M6((_1437_v46).Merge(_1437_v46), _this.F7(), _1441_v50, globalState)
				_1444_v51 = _out26
				var _1445_v52 _dafny.Array
				_ = _1445_v52
				var _len0_47 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_47
				var _nw245 _dafny.Array
				_ = _nw245
				if _len0_47.Cmp(_dafny.Zero) == 0 {
					_nw245 = _dafny.NewArray(_len0_47)
				} else {
					var _init47 func(_dafny.Int) D6 = (func(_1446_v51 bool, _1447_v44 _dafny.Array) func(_dafny.Int) D6 {
						return func(_1448_i2 _dafny.Int) D6 {
							return Companion_D6_.Create_DC14_(_this.F7(), _dafny.MultiSetOf(_1446_v51, (_1447_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1447_v44), 0))).Int()).(bool)), (_this).F6())
						}
					})(_1444_v51, _1435_v44)
					_ = _init47
					var _element0_47 = _init47(_dafny.Zero)
					_ = _element0_47
					_nw245 = _dafny.NewArrayFromExample(_element0_47, nil, _len0_47)
					(_nw245).ArraySet1(_element0_47, 0)
					var _nativeLen0_47 = (_len0_47).Int()
					_ = _nativeLen0_47
					for _i0_47 := 1; _i0_47 < _nativeLen0_47; _i0_47++ {
						(_nw245).ArraySet1(_init47(_dafny.IntOf(_i0_47)), _i0_47)
					}
				}
				_1445_v52 = _nw245
				var _1449_v53 D6
				_ = _1449_v53
				_1449_v53 = Companion_D6_.Create_DC14_(_dafny.IntOfUint32((_dafny.SeqOf(!((_1435_v44).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(450), _dafny.ArrayLen((_1435_v44), 0))).Int()).(bool)), _1444_v51)).Cardinality()), _dafny.MultiSetOf(_1428_v37), _this.F7())
				var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1445_v52), 0))
				_ = _index267
				(_1445_v52).ArraySet1(_1449_v53, (_index267).Int())
				var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(706), _dafny.ArrayLen((_1445_v52), 0))
				_ = _index268
				(_1445_v52).ArraySet1(_1449_v53, (_index268).Int())
				var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))
				_ = _index269
				(_1434_v43).ArraySet1(Companion_Default___.Fm1(((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(135)) > 0, globalState), (_index269).Int())
			} else {
				var _1450_v54 *C0
				_ = _1450_v54
				var _nw246 *C0 = New_C0_()
				_ = _nw246
				_nw246.Ctor__(_dafny.Companion_Sequence_.Equal((_1422_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1422_v31), 0))).Int()).(_dafny.Sequence), _1418_v27))
				_1450_v54 = _nw246
				var _1451_v55 _dafny.MultiSet
				_ = _1451_v55
				_1451_v55 = _dafny.MultiSetOf((_this).F11(), (_this).F11(), (_this).F11())
				var _1452_v56 D6
				_ = _1452_v56
				_1452_v56 = Companion_D6_.Create_DC15_((_1428_v37) == (_1428_v37), p0, (_dafny.IntOfInt64(857)).Times((_1420_v29).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1420_v29).Cardinality()))).Uint32()).(_dafny.Int)), Companion_Default___.SafeDivisionInt((_dafny.MultiSetOf((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int))).Cardinality(), (_this).F6()))
				var _pat_let_tv37 = _1419_v28
				_ = _pat_let_tv37
				var _1453_v57 _dafny.Sequence
				_ = _1453_v57
				_1453_v57 = _dafny.SeqOf((func() D10 {
					if !(_1450_v54.F10) {
						return _1431_v40
					}
					return func(_pat_let36_0 D10) D10 {
						return func(_1454_dt__update__tmp_h4 D10) D10 {
							return func(_pat_let37_0 _dafny.CodePoint) D10 {
								return func(_1455_dt__update_hcf51_h0 _dafny.CodePoint) D10 {
									return Companion_D10_.Create_DC24_(_1455_dt__update_hcf51_h0)
								}(_pat_let37_0)
							}(_pat_let_tv37)
						}(_pat_let36_0)
					}(_1431_v40)
				})(), _1431_v40, _1431_v40, _1431_v40)
				var _1456_v58 _dafny.MultiSet
				_ = _1456_v58
				_1456_v58 = _dafny.MultiSetOf(_1450_v54.F10)
				var _rhs306 _dafny.Int = ((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int)).Minus((_this.F7()).Plus((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int)))
				_ = _rhs306
				var _rhs307 _dafny.MultiSet = _1456_v58
				_ = _rhs307
				var _rhs308 D6 = _1452_v56
				_ = _rhs308
				var _rhs309 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(Companion_Default___.Fm39(_dafny.IntOfUint32((_dafny.SeqOf(p0)).Cardinality()), _dafny.CodePoint('h'), p0, (_this).F6(), globalState), _1431_v40), _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1453_v57, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1453_v57).Cardinality()))).Uint32(), _1431_v40), _1453_v57))
				_ = _rhs309
				var _rhs310 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("jl")
				_ = _rhs310
				var _lhs218 *GlobalState = globalState
				_ = _lhs218
				_lhs218.F1 = _rhs306
				_1451_v55 = _rhs307
				_1452_v56 = _rhs308
				_1453_v57 = _rhs309
				_1418_v27 = _rhs310
				var _rhs311 _dafny.Int = _dafny.IntOfInt64(753)
				_ = _rhs311
				var _rhs312 _dafny.Int = p0
				_ = _rhs312
				var _rhs313 _dafny.Int = p0
				_ = _rhs313
				var _lhs219 *GlobalState = globalState
				_ = _lhs219
				var _lhs220 *GlobalState = globalState
				_ = _lhs220
				var _lhs221 *GlobalState = globalState
				_ = _lhs221
				_lhs219.F1 = _rhs311
				_lhs220.F1 = _rhs312
				_lhs221.F1 = _rhs313
				(globalState).F1 = ((_this.F7()).Plus((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int))).Minus((_1434_v43).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(968), _dafny.ArrayLen((_1434_v43), 0))).Int()).(_dafny.Int))
				_1423_v32 = (_1423_v32).Update((_dafny.IntOfInt64(356)).Minus((func() _dafny.Int {
					if (_1421_v30).Contains(p0) {
						return (_1421_v30).Multiplicity(p0)
					}
					return (_this).F6()
				})()), _this.F7())
			}
		} else if _source17.Is_DC26() {
			var _1457___mcc_h0 bool = _source17.Get_().(D10_DC26).Cf52
			_ = _1457___mcc_h0
			var _1458_cf52 bool = _1457___mcc_h0
			_ = _1458_cf52
			var _1459_v59 _dafny.Array
			_ = _1459_v59
			var _nw247 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw247
			_1459_v59 = _nw247
			var _1460_v60 _dafny.Map
			_ = _1460_v60
			_1460_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1458_cf52)
			var _index270 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1459_v59), 0))
			_ = _index270
			(_1459_v59).ArraySet1((func() bool {
				if (_1460_v60).Contains(Companion_Default___.Fm1(_1458_cf52, globalState)) {
					return (_1460_v60).Get(Companion_Default___.Fm1(_1458_cf52, globalState)).(bool)
				}
				return (_this).F11()
			})(), (_index270).Int())
			var _index271 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(505), _dafny.ArrayLen((_1459_v59), 0))
			_ = _index271
			(_1459_v59).ArraySet1((_this).F11(), (_index271).Int())
			_1459_v59 = _1459_v59
			_1428_v37 = (_this).F11()
			var _1461_v61 _dafny.Array
			_ = _1461_v61
			var _nw248 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(20))
			_ = _nw248
			_1461_v61 = _nw248
			var _index272 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1461_v61), 0))
			_ = _index272
			(_1461_v61).ArraySet1((func() _dafny.Int {
				if _1458_cf52 {
					return (_this).F6()
				}
				return (_dafny.Zero).Minus(_this.F7())
			})(), (_index272).Int())
			var _index273 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_1461_v61), 0))
			_ = _index273
			(_1461_v61).ArraySet1(((_dafny.Zero).Minus(_this.F7())).Minus((_this).F6()), (_index273).Int())
		} else {
			var _1462___mcc_h1 _dafny.CodePoint = _source17.Get_().(D10_DC24).Cf51
			_ = _1462___mcc_h1
			var _1463_cf51 _dafny.CodePoint = _1462___mcc_h1
			_ = _1463_cf51
			_1428_v37 = !(_1428_v37)
			var _1464_v62 D5
			_ = _1464_v62
			_1464_v62 = Companion_D5_.Create_DC11_(_1429_v38)
			var _1465_v63 _dafny.Map
			_ = _1465_v63
			_1465_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1464_v62, (_this).F6())
			var _rhs314 _dafny.Int = (_dafny.IntOfInt64(-995)).Times(Companion_Default___.SafeModuloInt((_this).F6(), _dafny.IntOfInt64(477)))
			_ = _rhs314
			var _rhs315 bool = Companion_Default___.Fm0(_1428_v37, _1423_v32, (_this.F7()).Minus((_this).F6()), globalState)
			_ = _rhs315
			var _rhs316 bool = (_1429_v38).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
				if (_1465_v63).Contains(_1464_v62) {
					return (_1465_v63).Get(_1464_v62).(_dafny.Int)
				}
				return (_this).F6()
			})(), _dafny.IntOfUint32((_1429_v38).Cardinality()))).Uint32()).(bool)
			_ = _rhs316
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			_lhs222.F1 = _rhs314
			_1428_v37 = _rhs315
			_1428_v37 = _rhs316
			var _1466_v64 _dafny.Set
			_ = _1466_v64
			_1466_v64 = _dafny.SetOf(_1428_v37)
			var _1467_v65 _dafny.MultiSet
			_ = _1467_v65
			_1467_v65 = _dafny.MultiSetOf(Companion_Default___.Fm0((_this).F11(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1466_v64).Cardinality(), _dafny.IntOfUint32((_1429_v38).Cardinality())), _dafny.IntOfUint32((_1420_v29).Cardinality()), globalState))
			if (_1467_v65).IsDisjointFrom(_1467_v65) {
				var _1468_v66 *C2
				_ = _1468_v66
				var _nw249 *C2 = New_C2_()
				_ = _nw249
				_nw249.Ctor__((_this).F6(), _dafny.IntOfInt64(383))
				_1468_v66 = _nw249
				var _1469_v69 D1
				_ = _1469_v69
				_1469_v69 = Companion_D1_.Create_DC4_((_this).F11())
				var _1470_v70 D2
				_ = _1470_v70
				_1470_v70 = Companion_D2_.Create_DC7_((_this).F11(), (_1420_v29).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1420_v29).Cardinality()))).Uint32()).(_dafny.Int), func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(377), _dafny.IntOfInt64(-191))); ; {
						_compr_40, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _1471_v68 _dafny.Int
						_1471_v68 = interface{}(_compr_40).(_dafny.Int)
						if ((_dafny.IntOfInt64(377)).Cmp(_1471_v68) <= 0) && ((_1471_v68).Cmp(_dafny.IntOfInt64(-191)) < 0) {
							_coll40.Add(Companion_Default___.SafeDivisionInt(_1471_v68, _this.F7()), _1428_v37)
						}
					}
					return _coll40.ToMap()
				}(), _1469_v69)
				var _1472_v71 D8
				_ = _1472_v71
				_1472_v71 = Companion_D8_.Create_DC20_(Companion_Default___.Fm0(_1428_v37, func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(135), _dafny.IntOfInt64(264))); ; {
						_compr_41, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _1473_v67 _dafny.Int
						_1473_v67 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(135)).Cmp(_1473_v67) <= 0) && ((_1473_v67).Cmp(_dafny.IntOfInt64(264)) < 0) {
							_coll41.Add((_1473_v67).Times(_dafny.IntOfUint32((_1418_v27).Cardinality())), _dafny.IntOfInt64(277))
						}
					}
					return _coll41.ToMap()
				}(), (_this).F6(), globalState), _1428_v37, (_1470_v70).Dtor_cf13(), false)
				var _1474_v72 _dafny.Map
				_ = _1474_v72
				_1474_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1423_v32).Cardinality(), _1472_v71)
				_1474_v72 = (_1474_v72).Update(p0, _1472_v71)
				_1428_v37 = (_1429_v38).Select((Companion_Default___.SafeIndex(((_dafny.Zero).Minus(_dafny.IntOfUint32((_1429_v38).Cardinality()))).Plus(_this.F7()), _dafny.IntOfUint32((_1429_v38).Cardinality()))).Uint32()).(bool)
				var _1475_v73 _dafny.Array
				_ = _1475_v73
				var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(16))
				_ = _nw250
				_1475_v73 = _nw250
				_1475_v73 = _1475_v73
				var _1476_v74 _dafny.Map
				_ = _1476_v74
				_1476_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1428_v37, _1428_v37)
				var _1477_v75 D0
				_ = _1477_v75
				_1477_v75 = Companion_D0_.Create_DC0_(_1466_v64)
				var _1478_v76 _dafny.Map
				_ = _1478_v76
				_1478_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1477_v75, _1428_v37)
				var _1479_v77 D1
				_ = _1479_v77
				_1479_v77 = Companion_D1_.Create_DC2_(_1478_v76)
				var _pat_let_tv38 = _1477_v75
				_ = _pat_let_tv38
				var _pat_let_tv39 = _1478_v76
				_ = _pat_let_tv39
				var _1480_v78 D2
				_ = _1480_v78
				_1480_v78 = Companion_D2_.Create_DC6_(func(_pat_let38_0 D1) D1 {
					return func(_1483_dt__update__tmp_h7 D1) D1 {
						return func(_pat_let41_0 _dafny.Map) D1 {
							return func(_1484_dt__update_hcf4_h2 _dafny.Map) D1 {
								return Companion_D1_.Create_DC2_(_1484_dt__update_hcf4_h2)
							}(_pat_let41_0)
						}(_pat_let_tv39)
					}(_pat_let38_0)
				}(func(_pat_let39_0 D1) D1 {
					return func(_1481_dt__update__tmp_h6 D1) D1 {
						return func(_pat_let40_0 _dafny.Map) D1 {
							return func(_1482_dt__update_hcf4_h1 _dafny.Map) D1 {
								return Companion_D1_.Create_DC2_(_1482_dt__update_hcf4_h1)
							}(_pat_let40_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv38, (_this).F11()))
					}(_pat_let39_0)
				}(_1479_v77)), (_1422_v31).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(56), _dafny.ArrayLen((_1422_v31), 0))).Int()).(_dafny.Sequence), (_this).F6())
				(_1468_v66).M8(_1476_v74, _1480_v78, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(119))).Uint32(), func(coer64 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg64 _dafny.Int) interface{} {
						return coer64(arg64)
					}
				}((func(_1485_cf51 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1486_i3 _dafny.Int) _dafny.CodePoint {
						return _1485_cf51
					}
				})(_1463_cf51))), globalState)
			} else {
				var _1487_v79 _dafny.Map
				_ = _1487_v79
				_1487_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1428_v37, (_1466_v64).Difference(_1466_v64))
				var _1488_v80 _dafny.Array
				_ = _1488_v80
				var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.One)
				_ = _nw251
				_1488_v80 = _nw251
				var _1489_v81 _dafny.Map
				_ = _1489_v81
				_1489_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(268), (_this).F11())
				var _rhs317 _dafny.Int = p0
				_ = _rhs317
				var _rhs318 _dafny.Map = _1487_v79
				_ = _rhs318
				var _rhs319 _dafny.Int = Companion_Default___.SafeDivisionInt(_this.F7(), ((_this).F6()).Plus((_1489_v81).Cardinality()))
				_ = _rhs319
				var _rhs320 _dafny.Array = (func() _dafny.Array {
					if Companion_Default___.Fm0(_1428_v37, _1423_v32, (_1467_v65).Cardinality(), globalState) {
						return _1488_v80
					}
					return _1488_v80
				})()
				_ = _rhs320
				var _lhs223 *GlobalState = globalState
				_ = _lhs223
				var _lhs224 *GlobalState = globalState
				_ = _lhs224
				_lhs223.F1 = _rhs317
				_1487_v79 = _rhs318
				_lhs224.F1 = _rhs319
				_1488_v80 = _rhs320
				var _1490_v82 _dafny.Array
				_ = _1490_v82
				var _nwElement0_52 bool = _1428_v37
				_ = _nwElement0_52
				var _nw252 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(20))
				_ = _nw252
				(_nw252).ArraySet1(_nwElement0_52, 0)
				(_nw252).ArraySet1((_this).F11(), 1)
				(_nw252).ArraySet1(_1428_v37, 2)
				(_nw252).ArraySet1(_1428_v37, 3)
				(_nw252).ArraySet1(!(_1428_v37), 4)
				(_nw252).ArraySet1((_this).F11(), 5)
				(_nw252).ArraySet1(false, 6)
				(_nw252).ArraySet1(!((_this).F11()), 7)
				(_nw252).ArraySet1((_this).F11(), 8)
				(_nw252).ArraySet1((_this).F11(), 9)
				(_nw252).ArraySet1((_this).F11(), 10)
				(_nw252).ArraySet1(_1428_v37, 11)
				(_nw252).ArraySet1(_1428_v37, 12)
				(_nw252).ArraySet1(_1428_v37, 13)
				(_nw252).ArraySet1((_this).F11(), 14)
				(_nw252).ArraySet1(_1428_v37, 15)
				(_nw252).ArraySet1(_1428_v37, 16)
				(_nw252).ArraySet1((_this).F11(), 17)
				(_nw252).ArraySet1(!((_this).F11()), 18)
				(_nw252).ArraySet1(_1428_v37, 19)
				_1490_v82 = _nw252
				var _1491_v83 _dafny.Array
				_ = _1491_v83
				var _nwElement0_53 _dafny.Array = _1490_v82
				_ = _nwElement0_53
				var _nw253 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_53, nil, _dafny.IntOfInt64(6))
				_ = _nw253
				(_nw253).ArraySet1(_nwElement0_53, 0)
				(_nw253).ArraySet1(_1490_v82, 1)
				(_nw253).ArraySet1(_1490_v82, 2)
				(_nw253).ArraySet1(_1490_v82, 3)
				(_nw253).ArraySet1(_1490_v82, 4)
				(_nw253).ArraySet1(_1490_v82, 5)
				_1491_v83 = _nw253
				var _index274 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_1491_v83), 0))
				_ = _index274
				(_1491_v83).ArraySet1(_1490_v82, (_index274).Int())
				var _index275 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(732), _dafny.ArrayLen((_1491_v83), 0))
				_ = _index275
				(_1491_v83).ArraySet1(_1490_v82, (_index275).Int())
				var _1492_v84 _dafny.Map
				_ = _1492_v84
				_1492_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _1490_v82)
				(globalState).F1 = ((_1492_v84).Merge(_1492_v84)).Cardinality()
				_1428_v37 = Companion_Default___.Fm0(((_this).F6()).Cmp((_this).F6()) <= 0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6()), (_dafny.Zero).Minus(_this.F7()), globalState)
				_1428_v37 = (_this).F11()
			}
			var _1493_v85 *C1
			_ = _1493_v85
			var _nw254 *C1 = New_C1_()
			_ = _nw254
			_nw254.Ctor__((func() _dafny.CodePoint {
				if false {
					return _dafny.CodePoint('b')
				}
				return _1419_v28
			})(), (_dafny.IntOfUint32((_1418_v27).Cardinality())).Times((_1430_v39).Cardinality()), _this.F7())
			_1493_v85 = _nw254
		}
		var _1494_v86 _dafny.Array
		_ = _1494_v86
		var _nwElement0_54 bool = (_this).F11()
		_ = _nwElement0_54
		var _nw255 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_54, nil, _dafny.One)
		_ = _nw255
		(_nw255).ArraySet1(_nwElement0_54, 0)
		_1494_v86 = _nw255
		_1494_v86 = _1494_v86
	}
}
func (_this *C4) F11() bool {
	{
		return _this._f11
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f7 _dafny.Int
	_f6 _dafny.Int
	_f8 _dafny.Array
	_f9 _dafny.Sequence
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f7 = _dafny.Zero
	_this._f6 = _dafny.Zero
	_this._f8 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f9 = _dafny.EmptySeq
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

func (_this *C5) F7() _dafny.Int {
	return _this._f7
}
func (_this *C5) F7_set_(value _dafny.Int) {
	_this._f7 = value
}
func (_this *C5) F6() _dafny.Int {
	return _this._f6
}
func (_this *C5) Ctor__(f8 _dafny.Array, f9 _dafny.Sequence, f6 _dafny.Int, f7 _dafny.Int) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f6 = f6
		(_this)._f7 = f7
	}
}
func (_this *C5) Fm4(globalState *GlobalState) bool {
	{
		return false
	}
}
func (_this *C5) M1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) (bool, bool, _dafny.Int) {
	{
		var r0 bool = false
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1495_v0 _dafny.CodePoint
		_ = _1495_v0
		_1495_v0 = _dafny.CodePoint('p')
		var _1496_v1 _dafny.Sequence
		_ = _1496_v1
		_1496_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ybpygmgw")
		_1495_v0 = (_1496_v1).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if p0 {
				return p2
			}
			return p2
		})(), _dafny.IntOfUint32((_1496_v1).Cardinality()))).Uint32()).(_dafny.CodePoint)
		var _1497_v2 _dafny.Array
		_ = _1497_v2
		var _nw256 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
		_ = _nw256
		_1497_v2 = _nw256
		var _1498_v3 _dafny.Array
		_ = _1498_v3
		var _nwElement0_55 bool = p0
		_ = _nwElement0_55
		var _nw257 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_55, nil, _dafny.IntOfInt64(26))
		_ = _nw257
		(_nw257).ArraySet1(_nwElement0_55, 0)
		(_nw257).ArraySet1(p0, 1)
		(_nw257).ArraySet1(p0, 2)
		(_nw257).ArraySet1(p0, 3)
		(_nw257).ArraySet1(p0, 4)
		(_nw257).ArraySet1(p0, 5)
		(_nw257).ArraySet1(p0, 6)
		(_nw257).ArraySet1(p0, 7)
		(_nw257).ArraySet1(p0, 8)
		(_nw257).ArraySet1(p0, 9)
		(_nw257).ArraySet1(p0, 10)
		(_nw257).ArraySet1(p0, 11)
		(_nw257).ArraySet1(p0, 12)
		(_nw257).ArraySet1(p0, 13)
		(_nw257).ArraySet1(p0, 14)
		(_nw257).ArraySet1(p0, 15)
		(_nw257).ArraySet1(p0, 16)
		(_nw257).ArraySet1(true, 17)
		(_nw257).ArraySet1(!(p0), 18)
		(_nw257).ArraySet1(p0, 19)
		(_nw257).ArraySet1(p0, 20)
		(_nw257).ArraySet1(p0, 21)
		(_nw257).ArraySet1(p0, 22)
		(_nw257).ArraySet1(p0, 23)
		(_nw257).ArraySet1(p0, 24)
		(_nw257).ArraySet1(true, 25)
		_1498_v3 = _nw257
		var _index276 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))
		_ = _index276
		(_1497_v2).ArraySet1(_1498_v3, (_index276).Int())
		var _index277 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))
		_ = _index277
		(_1497_v2).ArraySet1(_1498_v3, (_index277).Int())
		var _1499_v4 _dafny.Map
		_ = _1499_v4
		_1499_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
		var _1500_v5 _dafny.Map
		_ = _1500_v5
		_1500_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt((_1499_v4).Cardinality(), _dafny.IntOfInt64(387)), (_this).F6())
		r2 = (_1500_v5).Cardinality()
		var _hi14 _dafny.Int = p2
		_ = _hi14
		for _1501_i0 := (_dafny.Zero).Minus((_this).F6()); _1501_i0.Cmp(_hi14) < 0; _1501_i0 = _1501_i0.Plus(_dafny.One) {
			var _1502_v6 _dafny.Array
			_ = _1502_v6
			var _nw258 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(21))
			_ = _nw258
			_1502_v6 = _nw258
			var _1503_v7 _dafny.Set
			_ = _1503_v7
			_1503_v7 = _dafny.SetOf(p0, p0)
			var _1504_v8 D0
			_ = _1504_v8
			_1504_v8 = Companion_D0_.Create_DC0_(_1503_v7)
			var _1505_v9 _dafny.Map
			_ = _1505_v9
			_1505_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v8, p0)
			var _pat_let_tv40 = _1503_v7
			_ = _pat_let_tv40
			var _1506_v11 _dafny.Set
			_ = _1506_v11
			_1506_v11 = _dafny.SetOf(_1504_v8, func(_pat_let42_0 D0) D0 {
				return func(_1507_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let43_0 _dafny.Set) D0 {
						return func(_1508_dt__update_hcf0_h0 _dafny.Set) D0 {
							return Companion_D0_.Create_DC0_(_1508_dt__update_hcf0_h0)
						}(_pat_let43_0)
					}(_pat_let_tv40)
				}(_pat_let42_0)
			}(_1504_v8), Companion_D0_.Create_DC0_(_1503_v7))
			var _1509_v13 _dafny.Sequence
			_ = _1509_v13
			_1509_v13 = _dafny.SeqOf(_1504_v8)
			var _1510_v14 D1
			_ = _1510_v14
			_1510_v14 = Companion_D1_.Create_DC2_(_1505_v9)
			var _1511_v15 _dafny.Array
			_ = _1511_v15
			var _nwElement0_56 _dafny.Map = _1505_v9
			_ = _nwElement0_56
			var _nw259 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_56, nil, _dafny.IntOfInt64(21))
			_ = _nw259
			(_nw259).ArraySet1(_nwElement0_56, 0)
			(_nw259).ArraySet1(_1505_v9, 1)
			(_nw259).ArraySet1(_1505_v9, 2)
			(_nw259).ArraySet1(_1505_v9, 3)
			(_nw259).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1504_v8, true), 4)
			(_nw259).ArraySet1(_1505_v9, 5)
			(_nw259).ArraySet1(_1505_v9, 6)
			(_nw259).ArraySet1(_1505_v9, 7)
			(_nw259).ArraySet1(_1505_v9, 8)
			(_nw259).ArraySet1(_1505_v9, 9)
			(_nw259).ArraySet1(func() _dafny.Map {
				var _coll42 = _dafny.NewMapBuilder()
				_ = _coll42
				for _iter46 := _dafny.Iterate((_1506_v11).Elements()); ; {
					_compr_42, _ok46 := _iter46()
					if !_ok46 {
						break
					}
					var _1512_v10 D0
					_1512_v10 = interface{}(_compr_42).(D0)
					if (_1506_v11).Contains(_1512_v10) {
						_coll42.Add(_1512_v10, p0)
					}
				}
				return _coll42.ToMap()
			}(), 10)
			(_nw259).ArraySet1(_1505_v9, 11)
			(_nw259).ArraySet1(_1505_v9, 12)
			(_nw259).ArraySet1(func() _dafny.Map {
				var _coll43 = _dafny.NewMapBuilder()
				_ = _coll43
				for _iter47 := _dafny.Iterate((_1509_v13).Elements()); ; {
					_compr_43, _ok47 := _iter47()
					if !_ok47 {
						break
					}
					var _1513_v12 D0
					_1513_v12 = interface{}(_compr_43).(D0)
					if _dafny.Companion_Sequence_.Contains(_1509_v13, _1513_v12) {
						_coll43.Add(_1513_v12, p0)
					}
				}
				return _coll43.ToMap()
			}(), 13)
			(_nw259).ArraySet1(_1505_v9, 14)
			(_nw259).ArraySet1((_1505_v9).Update(_1504_v8, false), 15)
			(_nw259).ArraySet1(_1505_v9, 16)
			(_nw259).ArraySet1(_1505_v9, 17)
			(_nw259).ArraySet1(_1505_v9, 18)
			(_nw259).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(_1503_v7), p0), 19)
			(_nw259).ArraySet1((_1510_v14).Dtor_cf4(), 20)
			_1511_v15 = _nw259
			var _index278 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1502_v6), 0))
			_ = _index278
			(_1502_v6).ArraySet1((func() _dafny.Array {
				if p0 {
					return _1511_v15
				}
				return _1511_v15
			})(), (_index278).Int())
			var _1514_v16 _dafny.Sequence
			_ = _1514_v16
			_1514_v16 = _dafny.SeqOf(_1511_v15)
			var _index279 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(429), _dafny.ArrayLen((_1502_v6), 0))
			_ = _index279
			(_1502_v6).ArraySet1((_1514_v16).Select((Companion_Default___.SafeIndex((Companion_D1_.Create_DC3_(p0, (_dafny.Zero).Minus(p1), _this.F7())).Dtor_cf6(), _dafny.IntOfUint32((_1514_v16).Cardinality()))).Uint32()).(_dafny.Array), (_index279).Int())
			var _1515_v17 _dafny.Map
			_ = _1515_v17
			_1515_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), (_this).F9())
			var _1516_v18 _dafny.Sequence
			_ = _1516_v18
			_1516_v18 = _dafny.SeqOf(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _1501_i0)).Cardinality()))
			var _1517_v19 _dafny.Map
			_ = _1517_v19
			_1517_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, Companion_Default___.Fm5(p0, p2, p2, p0, globalState))
			var _1518_v20 _dafny.Map
			_ = _1518_v20
			_1518_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_this).F9()).Cardinality()), p0)
			var _1519_v21 _dafny.Sequence
			_ = _1519_v21
			_1519_v21 = _dafny.SeqOf(p0, p0)
			var _1520_v22 _dafny.Map
			_ = _1520_v22
			_1520_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1519_v21, _1501_i0)
			var _1521_v23 _dafny.Array
			_ = _1521_v23
			var _nwElement0_57 _dafny.Sequence = (_this).F9()
			_ = _nwElement0_57
			var _nw260 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_57, nil, _dafny.IntOfInt64(18))
			_ = _nw260
			(_nw260).ArraySet1(_nwElement0_57, 0)
			(_nw260).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()), 1)
			(_nw260).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()), 2)
			(_nw260).ArraySet1((func() _dafny.Sequence {
				if (_1515_v17).Contains(_1501_i0) {
					return (_1515_v17).Get(_1501_i0).(_dafny.Sequence)
				}
				return (_this).F9()
			})(), 3)
			(_nw260).ArraySet1((_this).F9(), 4)
			(_nw260).ArraySet1((_this).F9(), 5)
			(_nw260).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_1516_v18).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1516_v18).Cardinality()))).Uint32()).(_dafny.Sequence)), 6)
			(_nw260).ArraySet1((_this).F9(), 7)
			(_nw260).ArraySet1((_this).F9(), 8)
			(_nw260).ArraySet1((func() _dafny.Sequence {
				if (_1517_v19).Contains(p0) {
					return (_1517_v19).Get(p0).(_dafny.Sequence)
				}
				return (_this).F9()
			})(), 9)
			(_nw260).ArraySet1((_this).F9(), 10)
			(_nw260).ArraySet1((_this).F9(), 11)
			(_nw260).ArraySet1(_dafny.SeqOf(p1, _dafny.IntOfInt64(821), (_dafny.Zero).Minus((_1518_v20).Cardinality()), (_this).F6(), p1), 12)
			(_nw260).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-17))).Uint32(), func(coer65 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}(func(_1522_i1 _dafny.Int) _dafny.Int {
				return (_this).F6()
			})), 13)
			(_nw260).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()), 14)
			(_nw260).ArraySet1((func() _dafny.Sequence {
				if (_1515_v17).Contains((func() _dafny.Int {
					if (_1520_v22).Contains(_1519_v21) {
						return (_1520_v22).Get(_1519_v21).(_dafny.Int)
					}
					return _1501_i0
				})()) {
					return (_1515_v17).Get((func() _dafny.Int {
						if (_1520_v22).Contains(_1519_v21) {
							return (_1520_v22).Get(_1519_v21).(_dafny.Int)
						}
						return _1501_i0
					})()).(_dafny.Sequence)
				}
				return (_this).F9()
			})(), 15)
			(_nw260).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()), 16)
			(_nw260).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(919))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg66 _dafny.Int) interface{} {
					return coer66(arg66)
				}
			}(func(_1523_i2 _dafny.Int) _dafny.Int {
				return (_this).F6()
			})), 17)
			_1521_v23 = _nw260
			var _index280 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_1521_v23), 0))
			_ = _index280
			(_1521_v23).ArraySet1(_dafny.SeqOf((_this).F6()), (_index280).Int())
			var _index281 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_1521_v23), 0))
			_ = _index281
			(_1521_v23).ArraySet1((func() _dafny.Sequence {
				if p0 {
					return _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(p0, p1, p2, p0, globalState), (_1516_v18).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1516_v18).Cardinality()))).Uint32()).(_dafny.Sequence))
				}
				return (_this).F9()
			})(), (_index281).Int())
			var _1524_v24 _dafny.MultiSet
			_ = _1524_v24
			_1524_v24 = _dafny.MultiSetOf(p2, _this.F7(), (_this).F6(), (_this).F6())
			var _1525_v25 _dafny.Set
			_ = _1525_v25
			_1525_v25 = _dafny.SetOf((_1518_v20).Cardinality(), _dafny.IntOfInt64(-479), p2)
			var _1526_v26 _dafny.Array
			_ = _1526_v26
			var _len0_48 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_48
			var _nw261 _dafny.Array
			_ = _nw261
			if _len0_48.Cmp(_dafny.Zero) == 0 {
				_nw261 = _dafny.NewArray(_len0_48)
			} else {
				var _init48 func(_dafny.Int) _dafny.Int = (func(_1527_p0 bool) func(_dafny.Int) _dafny.Int {
					return func(_1528_i3 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1528_i3, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(924))).Uint32(), func(coer67 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg67 _dafny.Int) interface{} {
								return coer67(arg67)
							}
						}((func(_1529_p0 bool) func(_dafny.Int) _dafny.Int {
							return func(_1530_i4 _dafny.Int) _dafny.Int {
								return (_dafny.MultiSetOf(!(_1529_p0))).Cardinality()
							}
						})(_1527_p0)))).Cardinality()))
					}
				})(p0)
				_ = _init48
				var _element0_48 = _init48(_dafny.Zero)
				_ = _element0_48
				_nw261 = _dafny.NewArrayFromExample(_element0_48, nil, _len0_48)
				(_nw261).ArraySet1(_element0_48, 0)
				var _nativeLen0_48 = (_len0_48).Int()
				_ = _nativeLen0_48
				for _i0_48 := 1; _i0_48 < _nativeLen0_48; _i0_48++ {
					(_nw261).ArraySet1(_init48(_dafny.IntOf(_i0_48)), _i0_48)
				}
			}
			_1526_v26 = _nw261
			var _1531_v27 _dafny.Map
			_ = _1531_v27
			_1531_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1495_v0, _1526_v26)
			var _1532_v28 _dafny.Array
			_ = _1532_v28
			var _nwElement0_58 _dafny.Int = (_this).F6()
			_ = _nwElement0_58
			var _nw262 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_58, nil, _dafny.IntOfInt64(23))
			_ = _nw262
			(_nw262).ArraySet1(_nwElement0_58, 0)
			(_nw262).ArraySet1(p2, 1)
			(_nw262).ArraySet1(_dafny.IntOfUint32((_1496_v1).Cardinality()), 2)
			(_nw262).ArraySet1(p1, 3)
			(_nw262).ArraySet1((func() _dafny.Int {
				if !(!(p0)) {
					return _dafny.IntOfUint32(((_1521_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_1521_v23), 0))).Int()).(_dafny.Sequence)).Cardinality())
				}
				return _1501_i0
			})(), 4)
			(_nw262).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("f")).Cardinality()), 5)
			(_nw262).ArraySet1((_dafny.Zero).Minus(_1501_i0), 6)
			(_nw262).ArraySet1(_1501_i0, 7)
			(_nw262).ArraySet1(_1501_i0, 8)
			(_nw262).ArraySet1((_this).F6(), 9)
			(_nw262).ArraySet1((_this).F6(), 10)
			(_nw262).ArraySet1((_1501_i0).Minus((_this).F6()), 11)
			(_nw262).ArraySet1(((_1521_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_1521_v23), 0))).Int()).(_dafny.Sequence)).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((_1521_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(973), _dafny.ArrayLen((_1521_v23), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Uint32()).(_dafny.Int), 12)
			(_nw262).ArraySet1(_this.F7(), 13)
			(_nw262).ArraySet1(_1501_i0, 14)
			(_nw262).ArraySet1((_1517_v19).Cardinality(), 15)
			(_nw262).ArraySet1((_dafny.Zero).Minus(p2), 16)
			(_nw262).ArraySet1((func() _dafny.Int {
				if (_1524_v24).Contains((_1525_v25).Cardinality()) {
					return (_1524_v24).Multiplicity((_1525_v25).Cardinality())
				}
				return _1501_i0
			})(), 17)
			(_nw262).ArraySet1((_this).F6(), 18)
			(_nw262).ArraySet1(_this.F7(), 19)
			(_nw262).ArraySet1((_1531_v27).Cardinality(), 20)
			(_nw262).ArraySet1(Companion_Default___.SafeModuloInt(_this.F7(), _this.F7()), 21)
			(_nw262).ArraySet1(p1, 22)
			_1532_v28 = _nw262
			_1526_v26 = (Companion_D2_.Create_DC5_(_1532_v28)).Dtor_cf9()
			if p0 {
				var _1533_v29 _dafny.MultiSet
				_ = _1533_v29
				_1533_v29 = _dafny.MultiSetOf(p0)
				var _1534_v30 _dafny.Array
				_ = _1534_v30
				var _nw263 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(2))
				_ = _nw263
				_1534_v30 = _nw263
				var _rhs321 bool = (func() bool {
					if (_1499_v4).Contains(!(p0) || (p0)) {
						return (_1499_v4).Get(!(p0) || (p0)).(bool)
					}
					return (_dafny.IntOfInt64(-273)).Cmp(p1) < 0
				})()
				_ = _rhs321
				var _rhs322 _dafny.MultiSet = ((_1533_v29).Union(_dafny.MultiSetOf(p0))).Union(_1533_v29)
				_ = _rhs322
				var _rhs323 _dafny.Array = _1534_v30
				_ = _rhs323
				var _rhs324 bool = p0
				_ = _rhs324
				r0 = _rhs321
				_1533_v29 = _rhs322
				_1534_v30 = _rhs323
				r1 = _rhs324
				var _1535_v31 *C0
				_ = _1535_v31
				var _nw264 *C0 = New_C0_()
				_ = _nw264
				_nw264.Ctor__(!((_1525_v25).IsDisjointFrom(_1525_v25)))
				_1535_v31 = _nw264
				_1496_v1 = _1496_v1
				(globalState).F1 = ((_1501_i0).Minus(_dafny.IntOfInt64(255))).Plus(_this.F7())
				_1495_v0 = _1495_v0
			} else {
				var _1536_v32 _dafny.Map
				_ = _1536_v32
				_1536_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1496_v1, p0)
				_1536_v32 = (_1536_v32).Update(_dafny.UnicodeSeqOfUtf8Bytes("m"), p0)
				(_this).F7_set_(Companion_Default___.SafeDivisionInt(_this.F7(), Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(853), _dafny.IntOfUint32((_1496_v1).Cardinality()))))
				var _1537_v33 _dafny.Sequence
				_ = _1537_v33
				_1537_v33 = _dafny.SeqOf(_1498_v3, _dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int())), _1498_v3)
				_1498_v3 = (_1537_v33).Select((Companion_Default___.SafeIndex(_1501_i0, _dafny.IntOfUint32((_1537_v33).Cardinality()))).Uint32()).(_dafny.Array)
				r1 = p0
				var _1538_v34 _dafny.MultiSet
				_ = _1538_v34
				_1538_v34 = _dafny.MultiSetOf(true, p0, p0, p0, p0)
				var _1539_v35 _dafny.MultiSet
				_ = _1539_v35
				_1539_v35 = _1538_v34
				var _1540_v36 _dafny.Map
				_ = _1540_v36
				_1540_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1538_v34)
				var _1541_v37 _dafny.Array
				_ = _1541_v37
				var _nwElement0_59 _dafny.MultiSet = (func() _dafny.MultiSet {
					if p0 {
						return _1538_v34
					}
					return _dafny.MultiSetOf(p0, p0)
				})()
				_ = _nwElement0_59
				var _nw265 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_59, nil, _dafny.IntOfInt64(24))
				_ = _nw265
				(_nw265).ArraySet1(_nwElement0_59, 0)
				(_nw265).ArraySet1(_1538_v34, 1)
				(_nw265).ArraySet1(Companion_Default___.Fm6(globalState), 2)
				(_nw265).ArraySet1(_1538_v34, 3)
				(_nw265).ArraySet1((_1539_v35), 4)
				(_nw265).ArraySet1(_1538_v34, 5)
				(_nw265).ArraySet1(_dafny.MultiSetOf(!(p0), p0, p0), 6)
				(_nw265).ArraySet1(_1538_v34, 7)
				(_nw265).ArraySet1(_dafny.MultiSetFromSeq(_1519_v21), 8)
				(_nw265).ArraySet1(_1538_v34, 9)
				(_nw265).ArraySet1(_1538_v34, 10)
				(_nw265).ArraySet1((_dafny.MultiSetOf(p0)).Update(p0, Companion_Default___.Abs(p1)), 11)
				(_nw265).ArraySet1(_1538_v34, 12)
				(_nw265).ArraySet1((_1538_v34).Union(_dafny.MultiSetOf(p0)), 13)
				(_nw265).ArraySet1(_1538_v34, 14)
				(_nw265).ArraySet1(_1538_v34, 15)
				(_nw265).ArraySet1((_1538_v34).Intersection((func() _dafny.MultiSet {
					if (_1540_v36).Contains(p0) {
						return (_1540_v36).Get(p0).(_dafny.MultiSet)
					}
					return Companion_Default___.Fm6(globalState)
				})()), 16)
				(_nw265).ArraySet1((_dafny.MultiSetOf(p0)).Difference(_1538_v34), 17)
				(_nw265).ArraySet1((_1538_v34).Intersection(_1538_v34), 18)
				(_nw265).ArraySet1(_1538_v34, 19)
				(_nw265).ArraySet1(_1538_v34, 20)
				(_nw265).ArraySet1(_1538_v34, 21)
				(_nw265).ArraySet1(_dafny.MultiSetOf(!(p0), p0, p0), 22)
				(_nw265).ArraySet1((_1538_v34).Update(p0, Companion_Default___.Abs(p2)), 23)
				_1541_v37 = _nw265
				_1541_v37 = _1541_v37
			}
		}
		var _1542_v38 *C0
		_ = _1542_v38
		var _nw266 *C0 = New_C0_()
		_ = _nw266
		_nw266.Ctor__(p0)
		_1542_v38 = _nw266
		var _arr0 _dafny.Array = _dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))
		_ = _arr0
		var _index282 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))), 0))
		_ = _index282
		(_arr0).ArraySet1(_1542_v38.F10, (_index282).Int())
		var _arr1 _dafny.Array = _dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))
		_ = _arr1
		var _index283 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))), 0))
		_ = _index283
		(_arr1).ArraySet1(p0, (_index283).Int())
		var _1543_v39 _dafny.MultiSet
		_ = _1543_v39
		_1543_v39 = _dafny.MultiSetOf((_dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(667), _dafny.ArrayLen((_dafny.ArrayCastTo((_1497_v2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(377), _dafny.ArrayLen((_1497_v2), 0))).Int()))), 0))).Int()).(bool), p0)
		var _1544_v40 _dafny.MultiSet
		_ = _1544_v40
		_1544_v40 = _1543_v39
		var _pat_let_tv41 = _1542_v38
		_ = _pat_let_tv41
		r0 = func(_source18 _dafny.MultiSet) bool {
			var _1545___mcc_h0 _dafny.MultiSet = _source18
			_ = _1545___mcc_h0
			var _1546_cf18 _dafny.MultiSet = _1545___mcc_h0
			_ = _1546_cf18
			return _pat_let_tv41.F10
		}(_1544_v40)
		r1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_1496_v1, _dafny.Companion_Sequence_.Concatenate(_1496_v1, _dafny.UnicodeSeqOfUtf8Bytes("lrjbfvxeg")))
		r2 = p2
		return r0, r1, r2
	}
}
func (_this *C5) M2(p0 _dafny.MultiSet, p1 bool, p2 bool, globalState *GlobalState) (bool, _dafny.Array) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var _1547_v0 _dafny.Array
		_ = _1547_v0
		var _nw267 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
		_ = _nw267
		_1547_v0 = _nw267
		var _1548_v1 _dafny.CodePoint
		_ = _1548_v1
		_1548_v1 = _dafny.CodePoint('g')
		var _1549_v2 _dafny.Set
		_ = _1549_v2
		_1549_v2 = _dafny.SetOf(_1548_v1)
		var _index284 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
		_ = _index284
		(_1547_v0).ArraySet1((_1549_v2).Cardinality(), (_index284).Int())
		var _index285 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
		_ = _index285
		(_1547_v0).ArraySet1(_this.F7(), (_index285).Int())
		var _1550_v3 _dafny.Set
		_ = _1550_v3
		_1550_v3 = _dafny.SetOf((_this).F6(), _this.F7())
		_1550_v3 = (_1550_v3).Union(_1550_v3)
		var _1551_v4 _dafny.Sequence
		_ = _1551_v4
		_1551_v4 = _dafny.UnicodeSeqOfUtf8Bytes("wukhtqknx")
		var _1552_v5 _dafny.MultiSet
		_ = _1552_v5
		_1552_v5 = _dafny.MultiSetOf(Companion_Default___.Fm2((_this).F9(), _this.F7(), _this.F7(), !(p1), globalState), _1551_v4, _1551_v4, _1551_v4)
		_1552_v5 = _1552_v5
		r0 = (p2) && (p1)
		var _1553_i0 _dafny.Int
		_ = _1553_i0
		_1553_i0 = _dafny.Zero
		{
			for !(_dafny.Companion_Sequence_.IsPrefixOf(_1551_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(139))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}((func(_1611_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_1612_i1 _dafny.Int) _dafny.CodePoint {
					return _1611_v1
				}
			})(_1548_v1))))) {
				{
					if (_1553_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1553_i0 = (_1553_i0).Plus(_dafny.One)
					var _1554_v6 _dafny.Set
					_ = _1554_v6
					_1554_v6 = _dafny.SetOf(p1, p1)
					var _1555_v7 D0
					_ = _1555_v7
					_1555_v7 = Companion_D0_.Create_DC0_(_1554_v6)
					var _1556_v8 _dafny.Map
					_ = _1556_v8
					_1556_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1555_v7, p1)
					var _1557_v9 D1
					_ = _1557_v9
					_1557_v9 = Companion_D1_.Create_DC2_(_1556_v8)
					var _1558_v10 D2
					_ = _1558_v10
					_1558_v10 = Companion_D2_.Create_DC6_(_1557_v9, _dafny.UnicodeSeqOfUtf8Bytes("jgkjag"), (_1554_v6).Cardinality())
					var _1559_v11 _dafny.Sequence
					_ = _1559_v11
					_1559_v11 = _dafny.SeqOf(p1, false)
					var _1560_v12 D1
					_ = _1560_v12
					_1560_v12 = Companion_D1_.Create_DC3_(p2, _dafny.IntOfUint32((_1559_v11).Cardinality()), (_this).F6())
					var _pat_let_tv42 = p1
					_ = _pat_let_tv42
					var _pat_let_tv43 = _1556_v8
					_ = _pat_let_tv43
					var _pat_let_tv44 = _1557_v9
					_ = _pat_let_tv44
					var _pat_let_tv45 = _1551_v4
					_ = _pat_let_tv45
					var _1561_v13 _dafny.Array
					_ = _1561_v13
					var _nwElement0_60 D2 = _1558_v10
					_ = _nwElement0_60
					var _nw268 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_60, nil, _dafny.IntOfInt64(25))
					_ = _nw268
					(_nw268).ArraySet1(_nwElement0_60, 0)
					(_nw268).ArraySet1(_1558_v10, 1)
					(_nw268).ArraySet1(_1558_v10, 2)
					(_nw268).ArraySet1(Companion_D2_.Create_DC6_(_1557_v9, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(759))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg68 _dafny.Int) interface{} {
							return coer68(arg68)
						}
					}((func(_1562_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_1563_i2 _dafny.Int) _dafny.CodePoint {
							return _1562_v1
						}
					})(_1548_v1))), (func(_pat_let44_0 D1) D1 {
						return func(_1564_dt__update__tmp_h0 D1) D1 {
							return func(_pat_let45_0 bool) D1 {
								return func(_1565_dt__update_hcf5_h0 bool) D1 {
									return func(_pat_let46_0 _dafny.Int) D1 {
										return func(_1566_dt__update_hcf7_h0 _dafny.Int) D1 {
											return Companion_D1_.Create_DC3_(_1565_dt__update_hcf5_h0, (_1564_dt__update__tmp_h0).Dtor_cf6(), _1566_dt__update_hcf7_h0)
										}(_pat_let46_0)
									}(_this.F7())
								}(_pat_let45_0)
							}(_pat_let_tv42)
						}(_pat_let44_0)
					}(_1560_v12)).Dtor_cf7()), 3)
					(_nw268).ArraySet1(Companion_D2_.Create_DC6_(_1557_v9, _1551_v4, _this.F7()), 4)
					(_nw268).ArraySet1(_1558_v10, 5)
					(_nw268).ArraySet1(func(_pat_let47_0 D2) D2 {
						return func(_1567_dt__update__tmp_h1 D2) D2 {
							return func(_pat_let48_0 _dafny.Int) D2 {
								return func(_1568_dt__update_hcf12_h0 _dafny.Int) D2 {
									return func(_pat_let49_0 D1) D2 {
										return func(_1569_dt__update_hcf10_h0 D1) D2 {
											return Companion_D2_.Create_DC6_(_1569_dt__update_hcf10_h0, (_1567_dt__update__tmp_h1).Dtor_cf11(), _1568_dt__update_hcf12_h0)
										}(_pat_let49_0)
									}(Companion_D1_.Create_DC2_(_pat_let_tv43))
								}(_pat_let48_0)
							}(_this.F7())
						}(_pat_let47_0)
					}(_1558_v10), 6)
					(_nw268).ArraySet1(_1558_v10, 7)
					(_nw268).ArraySet1(_1558_v10, 8)
					(_nw268).ArraySet1(_1558_v10, 9)
					(_nw268).ArraySet1(_1558_v10, 10)
					(_nw268).ArraySet1((func() D2 {
						if p1 {
							return _1558_v10
						}
						return _1558_v10
					})(), 11)
					(_nw268).ArraySet1(_1558_v10, 12)
					(_nw268).ArraySet1(_1558_v10, 13)
					(_nw268).ArraySet1(_1558_v10, 14)
					(_nw268).ArraySet1(func(_pat_let50_0 D2) D2 {
						return func(_1570_dt__update__tmp_h2 D2) D2 {
							return func(_pat_let51_0 D1) D2 {
								return func(_1571_dt__update_hcf10_h1 D1) D2 {
									return func(_pat_let52_0 _dafny.Sequence) D2 {
										return func(_1572_dt__update_hcf11_h0 _dafny.Sequence) D2 {
											return Companion_D2_.Create_DC6_(_1571_dt__update_hcf10_h1, _1572_dt__update_hcf11_h0, (_1570_dt__update__tmp_h2).Dtor_cf12())
										}(_pat_let52_0)
									}(_pat_let_tv45)
								}(_pat_let51_0)
							}(_pat_let_tv44)
						}(_pat_let50_0)
					}(_1558_v10), 15)
					(_nw268).ArraySet1(_1558_v10, 16)
					(_nw268).ArraySet1(Companion_Default___.Fm7(globalState), 17)
					(_nw268).ArraySet1(_1558_v10, 18)
					(_nw268).ArraySet1(_1558_v10, 19)
					(_nw268).ArraySet1(Companion_Default___.Fm7(globalState), 20)
					(_nw268).ArraySet1(_1558_v10, 21)
					(_nw268).ArraySet1(_1558_v10, 22)
					(_nw268).ArraySet1(_1558_v10, 23)
					(_nw268).ArraySet1(_1558_v10, 24)
					_1561_v13 = _nw268
					_1561_v13 = _1561_v13
					var _rhs325 bool = !(false)
					_ = _rhs325
					var _rhs326 bool = p2
					_ = _rhs326
					r0 = _rhs325
					r0 = _rhs326
					var _1573_v14 *C0
					_ = _1573_v14
					var _nw269 *C0 = New_C0_()
					_ = _nw269
					_nw269.Ctor__(p1)
					_1573_v14 = _nw269
					_1573_v14 = _1573_v14
					var _1574_v15 D2
					_ = _1574_v15
					_1574_v15 = Companion_D2_.Create_DC5_(_1547_v0)
					var _source19 D2 = _1574_v15
					_ = _source19
					if _source19.Is_DC6() {
						var _1575___mcc_h0 D1 = _source19.Get_().(D2_DC6).Cf10
						_ = _1575___mcc_h0
						var _1576___mcc_h1 _dafny.Sequence = _source19.Get_().(D2_DC6).Cf11
						_ = _1576___mcc_h1
						var _1577___mcc_h2 _dafny.Int = _source19.Get_().(D2_DC6).Cf12
						_ = _1577___mcc_h2
						var _1578_cf12 _dafny.Int = _1577___mcc_h2
						_ = _1578_cf12
						var _1579_cf11 _dafny.Sequence = _1576___mcc_h1
						_ = _1579_cf11
						var _1580_cf10 D1 = _1575___mcc_h0
						_ = _1580_cf10
						(_1573_v14).F10 = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate((_this).F9(), (_this).F9()), _1578_cf12)
						_1551_v4 = _1551_v4
						var _index286 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _index286
						(_1547_v0).ArraySet1((_dafny.IntOfUint32((_dafny.SeqOf(_1573_v14, _1573_v14)).Cardinality())).Plus((_1578_cf12).Minus((_this).F6())), (_index286).Int())
						(globalState).F1 = (_this).F6()
					} else if _source19.Is_DC7() {
						var _1581___mcc_h3 bool = _source19.Get_().(D2_DC7).Cf13
						_ = _1581___mcc_h3
						var _1582___mcc_h4 _dafny.Int = _source19.Get_().(D2_DC7).Cf14
						_ = _1582___mcc_h4
						var _1583___mcc_h5 _dafny.Map = _source19.Get_().(D2_DC7).Cf15
						_ = _1583___mcc_h5
						var _1584___mcc_h6 D1 = _source19.Get_().(D2_DC7).Cf16
						_ = _1584___mcc_h6
						var _1585_cf16 D1 = _1584___mcc_h6
						_ = _1585_cf16
						var _1586_cf15 _dafny.Map = _1583___mcc_h5
						_ = _1586_cf15
						var _1587_cf14 _dafny.Int = _1582___mcc_h4
						_ = _1587_cf14
						var _1588_cf13 bool = _1581___mcc_h3
						_ = _1588_cf13
						_1547_v0 = _1547_v0
						var _1589_v16 _dafny.Map
						_ = _1589_v16
						_1589_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32((_1551_v4).Cardinality()))
						var _1590_v17 _dafny.Map
						_ = _1590_v17
						_1590_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), (_1589_v16).Merge(_1589_v16))
						var _1591_v18 _dafny.Map
						_ = _1591_v18
						_1591_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1573_v14.F10, !(false))
						var _index287 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _index287
						var _rhs327 _dafny.Int = ((_1591_v18).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1573_v14.F10, p2))).Cardinality()
						_ = _rhs327
						var _rhs328 bool = _1588_cf13
						_ = _rhs328
						var _rhs329 _dafny.Sequence = _1559_v11
						_ = _rhs329
						var _rhs330 _dafny.Int = _dafny.IntOfInt64(-312)
						_ = _rhs330
						var _rhs331 _dafny.Map = _1590_v17
						_ = _rhs331
						var _lhs225 *C5 = _this
						_ = _lhs225
						var _lhs226 *C0 = _1573_v14
						_ = _lhs226
						var _lhs227 _dafny.Array = _1547_v0
						_ = _lhs227
						var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _lhs228
						_lhs225.F7_set_(_rhs327)
						_lhs226.F10 = _rhs328
						_1559_v11 = _rhs329
						(_lhs227).ArraySet1(_rhs330, (_lhs228).Int())
						_1590_v17 = _rhs331
						var _rhs332 bool = p2
						_ = _rhs332
						var _rhs333 bool = p2
						_ = _rhs333
						var _lhs229 *C0 = _1573_v14
						_ = _lhs229
						var _lhs230 *C0 = _1573_v14
						_ = _lhs230
						_lhs229.F10 = _rhs332
						_lhs230.F10 = _rhs333
						var _1592_v20 _dafny.Set
						_ = _1592_v20
						_1592_v20 = _dafny.SetOf(_1551_v4)
						_1588_cf13 = !((_1592_v20).Contains(Companion_Default___.Fm2((_this).F9(), (func() _dafny.Set {
							var _coll44 = _dafny.NewBuilder()
							_ = _coll44
							for _iter48 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1558_v10, _1588_cf13)).Keys().Elements()); ; {
								_compr_44, _ok48 := _iter48()
								if !_ok48 {
									break
								}
								var _1593_v19 D2
								_1593_v19 = interface{}(_compr_44).(D2)
								if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1558_v10, _1588_cf13)).Contains(_1593_v19) {
									_coll44.Add(_1593_v19)
								}
							}
							return _coll44.ToSet()
						}()).Cardinality(), _this.F7(), _1573_v14.F10, globalState)))
					} else if _source19.Is_DC5() {
						var _1594___mcc_h7 _dafny.Array = _source19.Get_().(D2_DC5).Cf9
						_ = _1594___mcc_h7
						var _1595_cf9 _dafny.Array = _1594___mcc_h7
						_ = _1595_cf9
						var _1596_v21 _dafny.Map
						_ = _1596_v21
						_1596_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1554_v6).IsProperSubsetOf(_1554_v6), _1547_v0)
						_1596_v21 = (_1596_v21).Update(_1573_v14.F10, _1595_cf9)
						_1548_v1 = Companion_Default___.Fm8(_dafny.Companion_Sequence_.Concatenate(_1551_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(200))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg69 _dafny.Int) interface{} {
								return coer69(arg69)
							}
						}((func(_1597_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1598_i3 _dafny.Int) _dafny.CodePoint {
								return _1597_v1
							}
						})(_1548_v1)))), _1548_v1, _1573_v14.F10, (_this).Fm4(globalState), globalState)
						var _1599_v22 _dafny.Sequence
						_ = _1599_v22
						_1599_v22 = _dafny.SeqOf((_1558_v10).Dtor_cf11())
						var _1600_v23 _dafny.Array
						_ = _1600_v23
						var _nwElement0_61 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(81))).Uint32(), func(coer70 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg70 _dafny.Int) interface{} {
								return coer70(arg70)
							}
						}((func(_1601_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1602_i4 _dafny.Int) _dafny.CodePoint {
								return _1601_v1
							}
						})(_1548_v1)))
						_ = _nwElement0_61
						var _nw270 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_61, nil, _dafny.IntOfInt64(7))
						_ = _nw270
						(_nw270).ArraySet1(_nwElement0_61, 0)
						(_nw270).ArraySet1(_dafny.Companion_Sequence_.Update(_1551_v4, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9())).Cardinality()), _dafny.IntOfUint32((_1551_v4).Cardinality()))).Uint32(), _1548_v1), 1)
						(_nw270).ArraySet1(_1551_v4, 2)
						(_nw270).ArraySet1(_1551_v4, 3)
						(_nw270).ArraySet1(_1551_v4, 4)
						(_nw270).ArraySet1((_1599_v22).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1599_v22).Cardinality()))).Uint32()).(_dafny.Sequence), 5)
						(_nw270).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("sfyudb"), 6)
						_1600_v23 = _nw270
						var _1603_v24 _dafny.Map
						_ = _1603_v24
						_1603_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int), _1600_v23)
						_1603_v24 = (_1603_v24).Update((_this).F6(), _1600_v23)
						var _1604_v25 _dafny.Sequence
						_ = _1604_v25
						_1604_v25 = _dafny.SeqOf((func() *C0 {
							if _1573_v14.F10 {
								return _1573_v14
							}
							return _1573_v14
						})(), _1573_v14, _1573_v14)
						var _1605_v26 _dafny.MultiSet
						_ = _1605_v26
						_1605_v26 = _dafny.MultiSetOf(_1573_v14.F10, _1573_v14.F10)
						var _index288 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _index288
						var _rhs334 *C0 = (_1604_v25).Select((Companion_Default___.SafeIndex((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1604_v25).Cardinality()))).Uint32()).(*C0)
						_ = _rhs334
						var _rhs335 _dafny.Int = (Companion_Default___.SafeModuloInt((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int), (_this).F6())).Times((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int))
						_ = _rhs335
						var _rhs336 _dafny.Sequence = _dafny.SeqOf(_1551_v4, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(387))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg71 _dafny.Int) interface{} {
								return coer71(arg71)
							}
						}((func(_1606_v1 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_1607_i5 _dafny.Int) _dafny.CodePoint {
								return _1606_v1
							}
						})(_1548_v1))))
						_ = _rhs336
						var _rhs337 _dafny.CodePoint = _1548_v1
						_ = _rhs337
						var _rhs338 bool = ((_this).F6()).Cmp(((_1605_v26).Cardinality()).Times((_this).F6())) == 0
						_ = _rhs338
						var _lhs231 _dafny.Array = _1547_v0
						_ = _lhs231
						var _lhs232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _lhs232
						_1573_v14 = _rhs334
						(_lhs231).ArraySet1(_rhs335, (_lhs232).Int())
						_1599_v22 = _rhs336
						_1548_v1 = _rhs337
						r0 = _rhs338
					} else {
						var _1608___mcc_h8 D2 = _source19.Get_().(D2_DC8).Cf17
						_ = _1608___mcc_h8
						var _1609_cf17 D2 = _1608___mcc_h8
						_ = _1609_cf17
						var _index289 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))
						_ = _index289
						(_1547_v0).ArraySet1((_this).F6(), (_index289).Int())
						(globalState).F1 = _this.F7()
						var _1610_v27 _dafny.Set
						_ = _1610_v27
						_1610_v27 = _dafny.SetOf(Companion_D2_.Create_DC5_(_1547_v0), _1574_v15, _1574_v15, _1574_v15)
						_1610_v27 = (func() _dafny.Set {
							if _1573_v14.F10 {
								return _1610_v27
							}
							return _1610_v27
						})()
						_1559_v11 = _1559_v11
					}
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1613_v28 _dafny.Map
		_ = _1613_v28
		_1613_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1551_v4).Cardinality()), (_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int))
		var _1614_v29 _dafny.Set
		_ = _1614_v29
		_1614_v29 = _dafny.SetOf(p1)
		var _1615_v30 D0
		_ = _1615_v30
		_1615_v30 = Companion_D0_.Create_DC0_(_1614_v29)
		var _1616_v31 _dafny.Map
		_ = _1616_v31
		_1616_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1615_v30, p1)
		var _1617_v32 D1
		_ = _1617_v32
		_1617_v32 = Companion_D1_.Create_DC2_(_1616_v31)
		var _1618_v33 D2
		_ = _1618_v33
		_1618_v33 = Companion_D2_.Create_DC6_(_1617_v32, _1551_v4, _dafny.IntOfInt64(229))
		var _1619_v34 _dafny.Map
		_ = _1619_v34
		_1619_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
			if (_1613_v28).Contains((_this).F6()) {
				return (_1613_v28).Get((_this).F6()).(_dafny.Int)
			}
			return (Companion_Default___.Fm9(_dafny.IntOfUint32((_1551_v4).Cardinality()), globalState)).Cardinality()
		})(), Companion_D2_.Create_DC8_(_1618_v33))
		var _1620_v35 _dafny.MultiSet
		_ = _1620_v35
		_1620_v35 = _dafny.MultiSetOf(p1)
		var _1621_v36 _dafny.Sequence
		_ = _1621_v36
		_1621_v36 = _dafny.SeqOf(p2, p2)
		var _1622_v37 D2
		_ = _1622_v37
		_1622_v37 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC8_(_1618_v33)))
		_1619_v34 = (_1619_v34).Update((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((func() _dafny.Int {
			if (_1620_v35).Contains((_1621_v36).Select((Companion_Default___.SafeIndex((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1621_v36).Cardinality()))).Uint32()).(bool)) {
				return (_1620_v35).Multiplicity((_1621_v36).Select((Companion_Default___.SafeIndex((_1547_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_1547_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1621_v36).Cardinality()))).Uint32()).(bool))
			}
			return _this.F7()
		})(), _dafny.IntOfUint32((_1551_v4).Cardinality()))), _1622_v37)
		r0 = p1
		var _1623_v38 _dafny.Array
		_ = _1623_v38
		var _len0_49 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_49
		var _nw271 _dafny.Array
		_ = _nw271
		if _len0_49.Cmp(_dafny.Zero) == 0 {
			_nw271 = _dafny.NewArray(_len0_49)
		} else {
			var _init49 func(_dafny.Int) D0 = (func(_1624_v29 _dafny.Set) func(_dafny.Int) D0 {
				return func(_1625_i6 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_1624_v29)
				}
			})(_1614_v29)
			_ = _init49
			var _element0_49 = _init49(_dafny.Zero)
			_ = _element0_49
			_nw271 = _dafny.NewArrayFromExample(_element0_49, nil, _len0_49)
			(_nw271).ArraySet1(_element0_49, 0)
			var _nativeLen0_49 = (_len0_49).Int()
			_ = _nativeLen0_49
			for _i0_49 := 1; _i0_49 < _nativeLen0_49; _i0_49++ {
				(_nw271).ArraySet1(_init49(_dafny.IntOf(_i0_49)), _i0_49)
			}
		}
		_1623_v38 = _nw271
		r1 = _1623_v38
		return r0, r1
	}
}
func (_this *C5) M3(p0 _dafny.Array, p1 _dafny.Int, globalState *GlobalState) {
	{
		var _1626_v0 bool
		_ = _1626_v0
		_1626_v0 = true
		var _1627_v1 _dafny.Sequence
		_ = _1627_v1
		_1627_v1 = _dafny.SeqOf(_1626_v0)
		var _1628_v2 _dafny.Map
		_ = _1628_v2
		_1628_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_1627_v1, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1627_v1).Cardinality()))).Uint32(), _1626_v0), (_this).F6())
		var _1629_v3 _dafny.MultiSet
		_ = _1629_v3
		_1629_v3 = _dafny.MultiSetOf((_this).F6())
		(_this).F7_set_((func() _dafny.Int {
			if (_1628_v2).Contains(_dafny.Companion_Sequence_.Concatenate(_1627_v1, _1627_v1)) {
				return (_1628_v2).Get(_dafny.Companion_Sequence_.Concatenate(_1627_v1, _1627_v1)).(_dafny.Int)
			}
			return Companion_Default___.SafeDivisionInt(((_1629_v3).Update(_dafny.IntOfInt64(957), Companion_Default___.Abs((_this).F6()))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("adkmbscvj")).Cardinality()))
		})())
		var _1630_i0 _dafny.Int
		_ = _1630_i0
		_1630_i0 = _dafny.Zero
		{
			for _1626_v0 {
				{
					if (_1630_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1630_i0 = (_1630_i0).Plus(_dafny.One)
					var _1631_v4 _dafny.Sequence
					_ = _1631_v4
					_1631_v4 = _dafny.UnicodeSeqOfUtf8Bytes("lvswxolf")
					_1631_v4 = Companion_Default___.Fm2((_this).F9(), (func() _dafny.Int {
						if (_1627_v1).Select((Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32((_1627_v1).Cardinality()))).Uint32()).(bool) {
							return _this.F7()
						}
						return (_this).F6()
					})(), (_this).F6(), !(_1626_v0), globalState)
					var _1632_v5 *C0
					_ = _1632_v5
					var _nw272 *C0 = New_C0_()
					_ = _nw272
					_nw272.Ctor__(_1626_v0)
					_1632_v5 = _nw272
					var _1633_v6 D0
					_ = _1633_v6
					_1633_v6 = Companion_D0_.Create_DC1_(true, p0, (p1).Cmp(p1) == 0)
					var _source20 D0 = _1633_v6
					_ = _source20
					if _source20.Is_DC1() {
						var _1634___mcc_h0 bool = _source20.Get_().(D0_DC1).Cf1
						_ = _1634___mcc_h0
						var _1635___mcc_h1 _dafny.Array = _source20.Get_().(D0_DC1).Cf2
						_ = _1635___mcc_h1
						var _1636___mcc_h2 bool = _source20.Get_().(D0_DC1).Cf3
						_ = _1636___mcc_h2
						var _1637_cf3 bool = _1636___mcc_h2
						_ = _1637_cf3
						var _1638_cf2 _dafny.Array = _1635___mcc_h1
						_ = _1638_cf2
						var _1639_cf1 bool = _1634___mcc_h0
						_ = _1639_cf1
						var _1640_v7 _dafny.Array
						_ = _1640_v7
						var _nw273 _dafny.Array = _dafny.NewArray(_dafny.One)
						_ = _nw273
						_1640_v7 = _nw273
						var _index290 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1640_v7), 0))
						_ = _index290
						(_1640_v7).ArraySet1(_1632_v5, (_index290).Int())
						var _1641_v8 *C0
						_ = _1641_v8
						_1641_v8 = _1632_v5
						var _1642_v9 _dafny.Sequence
						_ = _1642_v9
						_1642_v9 = _dafny.SeqOf((_1641_v8), _1632_v5, _1632_v5, _1632_v5)
						var _index291 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(928), _dafny.ArrayLen((_1640_v7), 0))
						_ = _index291
						(_1640_v7).ArraySet1((_1642_v9).Select((Companion_Default___.SafeIndex((Companion_Default___.Fm10(_1632_v5.F10, _1626_v0, _1627_v1, true, globalState)).Cardinality(), _dafny.IntOfUint32((_1642_v9).Cardinality()))).Uint32()).(*C0), (_index291).Int())
						var _index292 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((p0), 0))
						_ = _index292
						(p0).ArraySet1(_1626_v0, (_index292).Int())
						var _index293 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(159), _dafny.ArrayLen((p0), 0))
						_ = _index293
						(p0).ArraySet1(_1637_cf3, (_index293).Int())
						var _rhs339 _dafny.Sequence = _1631_v4
						_ = _rhs339
						var _rhs340 bool = _1632_v5.F10
						_ = _rhs340
						var _rhs341 bool = !(_1629_v3).Equals((_1629_v3).Difference(_1629_v3))
						_ = _rhs341
						var _lhs233 *C0 = _1632_v5
						_ = _lhs233
						_1631_v4 = _rhs339
						_1639_cf1 = _rhs340
						_lhs233.F10 = _rhs341
						var _1643_v10 _dafny.Array
						_ = _1643_v10
						var _nw274 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
						_ = _nw274
						_1643_v10 = _nw274
						var _index294 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen(((_this).F8()), 0))
						_ = _index294
						((_this).F8()).ArraySet1(_1643_v10, (_index294).Int())
						var _index295 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen(((_this).F8()), 0))
						_ = _index295
						((_this).F8()).ArraySet1(_1643_v10, (_index295).Int())
					} else {
						var _1644___mcc_h3 _dafny.Set = _source20.Get_().(D0_DC0).Cf0
						_ = _1644___mcc_h3
						var _1645_cf0 _dafny.Set = _1644___mcc_h3
						_ = _1645_cf0
						(globalState).F1 = p1
						_1631_v4 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1631_v4, _1631_v4), _1631_v4)
						var _1646_v11 _dafny.Array
						_ = _1646_v11
						var _len0_50 _dafny.Int = _dafny.IntOfInt64(9)
						_ = _len0_50
						var _nw275 _dafny.Array
						_ = _nw275
						if _len0_50.Cmp(_dafny.Zero) == 0 {
							_nw275 = _dafny.NewArray(_len0_50)
						} else {
							var _init50 func(_dafny.Int) _dafny.Int = func(_1647_i1 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeDivisionInt(_1647_i1, _dafny.IntOfInt64(-444))
							}
							_ = _init50
							var _element0_50 = _init50(_dafny.Zero)
							_ = _element0_50
							_nw275 = _dafny.NewArrayFromExample(_element0_50, nil, _len0_50)
							(_nw275).ArraySet1(_element0_50, 0)
							var _nativeLen0_50 = (_len0_50).Int()
							_ = _nativeLen0_50
							for _i0_50 := 1; _i0_50 < _nativeLen0_50; _i0_50++ {
								(_nw275).ArraySet1(_init50(_dafny.IntOf(_i0_50)), _i0_50)
							}
						}
						_1646_v11 = _nw275
						var _1648_v12 T0
						_ = _1648_v12
						var _nw276 *C2 = New_C2_()
						_ = _nw276
						_nw276.Ctor__(_dafny.IntOfUint32((_1627_v1).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer73 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg73 _dafny.Int) interface{} {
								return coer73(arg73)
							}
						}(func(_1649_i2 _dafny.Int) _dafny.Int {
							return (_this).F6()
						}))).Cardinality()))
						_1648_v12 = _nw276
						var _1650_v13 _dafny.Map
						_ = _1650_v13
						_1650_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1648_v12, true)
						var _1651_v14 _dafny.Map
						_ = _1651_v14
						_1651_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1626_v0, (_1650_v13).Cardinality())
						var _index296 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1646_v11), 0))
						_ = _index296
						(_1646_v11).ArraySet1((_1651_v14).Cardinality(), (_index296).Int())
						var _index297 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
						_ = _index297
						(p0).ArraySet1(_1626_v0, (_index297).Int())
						var _index298 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1646_v11), 0))
						_ = _index298
						var _index299 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
						_ = _index299
						var _rhs342 _dafny.Int = _dafny.IntOfUint32((_1631_v4).Cardinality())
						_ = _rhs342
						var _rhs343 bool = _1632_v5.F10
						_ = _rhs343
						var _lhs234 _dafny.Array = _1646_v11
						_ = _lhs234
						var _lhs235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(176), _dafny.ArrayLen((_1646_v11), 0))
						_ = _lhs235
						var _lhs236 _dafny.Array = p0
						_ = _lhs236
						var _lhs237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(213), _dafny.ArrayLen((p0), 0))
						_ = _lhs237
						(_lhs234).ArraySet1(_rhs342, (_lhs235).Int())
						(_lhs236).ArraySet1(_rhs343, (_lhs237).Int())
						var _1652_v15 *C0
						_ = _1652_v15
						var _nw277 *C0 = New_C0_()
						_ = _nw277
						_nw277.Ctor__(!(_1632_v5.F10))
						_1652_v15 = _nw277
					}
					var _1653_v16 _dafny.Set
					_ = _1653_v16
					_1653_v16 = _dafny.SetOf(true, _1632_v5.F10, (_dafny.IntOfInt64(-182)).Cmp(_this.F7()) == 0)
					var _1654_v17 _dafny.Array
					_ = _1654_v17
					var _len0_51 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_51
					var _nw278 _dafny.Array
					_ = _nw278
					if _len0_51.Cmp(_dafny.Zero) == 0 {
						_nw278 = _dafny.NewArray(_len0_51)
					} else {
						var _init51 func(_dafny.Int) _dafny.Int = func(_1655_i3 _dafny.Int) _dafny.Int {
							return (_1655_i3).Plus((_this).F6())
						}
						_ = _init51
						var _element0_51 = _init51(_dafny.Zero)
						_ = _element0_51
						_nw278 = _dafny.NewArrayFromExample(_element0_51, nil, _len0_51)
						(_nw278).ArraySet1(_element0_51, 0)
						var _nativeLen0_51 = (_len0_51).Int()
						_ = _nativeLen0_51
						for _i0_51 := 1; _i0_51 < _nativeLen0_51; _i0_51++ {
							(_nw278).ArraySet1(_init51(_dafny.IntOf(_i0_51)), _i0_51)
						}
					}
					_1654_v17 = _nw278
					var _index300 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_1654_v17), 0))
					_ = _index300
					(_1654_v17).ArraySet1(_this.F7(), (_index300).Int())
					var _index301 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_1654_v17), 0))
					_ = _index301
					var _rhs344 _dafny.Set = _1653_v16
					_ = _rhs344
					var _rhs345 _dafny.Int = p1
					_ = _rhs345
					var _lhs238 _dafny.Array = _1654_v17
					_ = _lhs238
					var _lhs239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(118), _dafny.ArrayLen((_1654_v17), 0))
					_ = _lhs239
					_1653_v16 = _rhs344
					(_lhs238).ArraySet1(_rhs345, (_lhs239).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
		var _1656_v18 *C0
		_ = _1656_v18
		var _nw279 *C0 = New_C0_()
		_ = _nw279
		_nw279.Ctor__(false)
		_1656_v18 = _nw279
		var _1657_v19 _dafny.Map
		_ = _1657_v19
		_1657_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_1656_v18, _1656_v18)).Cardinality())))).Cardinality())).Cardinality()))
		var _1658_v20 _dafny.Sequence
		_ = _1658_v20
		_1658_v20 = _dafny.UnicodeSeqOfUtf8Bytes("qmtta")
		var _1659_v21 _dafny.Map
		_ = _1659_v21
		_1659_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1656_v18.F10, p1)
		var _1660_v22 _dafny.MultiSet
		_ = _1660_v22
		_1660_v22 = _dafny.MultiSetOf(_1626_v0)
		var _1661_v23 D6
		_ = _1661_v23
		_1661_v23 = Companion_D6_.Create_DC14_(_this.F7(), _1660_v22, (_dafny.MultiSetOf(_this.F7())).Cardinality())
		var _1662_v24 _dafny.Array
		_ = _1662_v24
		var _nwElement0_62 _dafny.Sequence = (_this).F9()
		_ = _nwElement0_62
		var _nw280 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_62, nil, _dafny.IntOfInt64(26))
		_ = _nw280
		(_nw280).ArraySet1(_nwElement0_62, 0)
		(_nw280).ArraySet1(_dafny.SeqOf((_this).F6(), ((_this).F9()).Select((Companion_Default___.SafeIndex((_1657_v19).Cardinality(), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32()).(_dafny.Int), (_this).F6()), 1)
		(_nw280).ArraySet1(_dafny.SeqOf(p1), 2)
		(_nw280).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(423))).Uint32(), func(coer74 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg74 _dafny.Int) interface{} {
				return coer74(arg74)
			}
		}(func(_1663_i5 _dafny.Int) _dafny.Int {
			return (_this).F6()
		})), 3)
		(_nw280).ArraySet1(_dafny.Companion_Sequence_.Update((_this).F9(), (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1658_v20).Cardinality())), 4)
		(_nw280).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F9(), _dafny.SeqOf(_this.F7(), _dafny.IntOfInt64(-262))), 5)
		(_nw280).ArraySet1((_this).F9(), 6)
		(_nw280).ArraySet1((_this).F9(), 7)
		(_nw280).ArraySet1((_this).F9(), 8)
		(_nw280).ArraySet1((_this).F9(), 9)
		(_nw280).ArraySet1(Companion_Default___.Fm5(_1656_v18.F10, (func() _dafny.Int {
			if (_1659_v21).Contains(_1626_v0) {
				return (_1659_v21).Get(_1626_v0).(_dafny.Int)
			}
			return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1626_v0, p1)).Cardinality()
		})(), Companion_Default___.Fm1(_1656_v18.F10, globalState), _1626_v0, globalState), 10)
		(_nw280).ArraySet1(_dafny.Companion_Sequence_.Update((_this).F9(), (Companion_Default___.SafeIndex(_this.F7(), _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32(), (_1661_v23).Dtor_cf23()), 11)
		(_nw280).ArraySet1((_this).F9(), 12)
		(_nw280).ArraySet1((_this).F9(), 13)
		(_nw280).ArraySet1(Companion_Default___.Fm5(true, _dafny.IntOfInt64(-647), p1, _1656_v18.F10, globalState), 14)
		(_nw280).ArraySet1((_this).F9(), 15)
		(_nw280).ArraySet1((_this).F9(), 16)
		(_nw280).ArraySet1((_this).F9(), 17)
		(_nw280).ArraySet1(_dafny.SeqOf((_dafny.Zero).Minus(Companion_Default___.Fm1(false, globalState)), _this.F7(), (_this).F6(), (_this).F6()), 18)
		(_nw280).ArraySet1((_this).F9(), 19)
		(_nw280).ArraySet1((_this).F9(), 20)
		(_nw280).ArraySet1(Companion_Default___.Fm35(_this.F7(), _this.F7(), globalState), 21)
		(_nw280).ArraySet1((_this).F9(), 22)
		(_nw280).ArraySet1((_this).F9(), 23)
		(_nw280).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm5(_1626_v0, p1, _this.F7(), _1626_v0, globalState), (_this).F9()), 24)
		(_nw280).ArraySet1((_this).F9(), 25)
		_1662_v24 = _nw280
		for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1662_v24), 0))); ; {
			_guard_loop_4, _ok49 := _iter49()
			if !_ok49 {
				break
			}
			var _1664_i4 _dafny.Int
			_1664_i4 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1664_i4).Sign() != -1) && ((_1664_i4).Cmp(_dafny.ArrayLen((_1662_v24), 0)) < 0)) {
				(_1662_v24).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(835))).Uint32(), func(coer75 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg75 _dafny.Int) interface{} {
						return coer75(arg75)
					}
				}((func(_1665_v21 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_1666_i6 _dafny.Int) _dafny.Int {
						return (_1665_v21).Cardinality()
					}
				})(_1659_v21))), (_this).F9()), (_1664_i4).Int())
			}
		}
		var _1667_v25 _dafny.CodePoint
		_ = _1667_v25
		_1667_v25 = _dafny.CodePoint('t')
		var _1668_v26 *C1
		_ = _1668_v26
		var _nw281 *C1 = New_C1_()
		_ = _nw281
		_nw281.Ctor__(Companion_Default___.Fm8(_1658_v20, _1667_v25, _1656_v18.F10, !(_1656_v18.F10), globalState), _dafny.IntOfUint32((_1658_v20).Cardinality()), _this.F7())
		_1668_v26 = _nw281
		_1668_v26 = _1668_v26
		var _1669_v27 *C3
		_ = _1669_v27
		var _nw282 *C3 = New_C3_()
		_ = _nw282
		_nw282.Ctor__(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((Companion_Default___.Fm2(_dafny.SeqOf(_dafny.IntOfInt64(24)), _this.F7(), (_dafny.Zero).Minus(_this.F7()), _1626_v0, globalState)).Cardinality()), _this.F7()), _dafny.IntOfInt64(876))
		_1669_v27 = _nw282
		var _1670_v28 *C3
		_ = _1670_v28
		var _nw283 *C3 = New_C3_()
		_ = _nw283
		_nw283.Ctor__(Companion_Default___.SafeDivisionInt((_1669_v27).Fm15(_this.F7(), p1, globalState), (_this).F6()), ((_this).F9()).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32(((_this).F9()).Cardinality()))).Uint32()).(_dafny.Int))
		_1670_v28 = _nw283
	}
}
func (_this *C5) F8() _dafny.Array {
	{
		return _this._f8
	}
}
func (_this *C5) F9() _dafny.Sequence {
	{
		return _this._f9
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
